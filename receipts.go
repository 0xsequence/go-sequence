package sequence

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
)

type Receipt struct {
	*types.Receipt

	MetaTxnID MetaTxnID
	Status    MetaTxnStatus
	Reason    string
	Logs      []*types.Log

	Index       uint
	Transaction *Transaction
	Receipts    []*Receipt
}

func (r *Receipt) Find(metaTxnID MetaTxnID) *Receipt {
	if r.MetaTxnID == metaTxnID {
		return nil // the parent of this receipt is the receipt we want
	}

	for _, receipt := range r.Receipts {
		if receipt.MetaTxnID == metaTxnID {
			return r
		}

		found := receipt.Find(metaTxnID)
		if found != nil {
			return found
		}
	}

	return nil
}

func (r *Receipt) setNativeReceipt(receipt *types.Receipt) {
	r.Receipt = receipt

	for _, child := range r.Receipts {
		child.setNativeReceipt(receipt)
	}
}

// This method is duplicated code from: `compressor/contract.go`
// can't be used directly, because it would create a circular dependency
func DecompressCalldata(
	ctx context.Context,
	provider *ethrpc.Provider,
	toAddress common.Address,
	calldata []byte,
) (common.Address, []byte, error) {
	data := calldata

	if len(data) == 0 {
		return common.Address{}, nil, fmt.Errorf("empty transaction data")
	}

	if data[0] != byte(0x00) {
		return common.Address{}, nil, fmt.Errorf("not compressed data")
	}

	// Replace first byte with `METHOD_DECOMPRESS_1` (0x06)
	// and call `.to`
	c2 := make([]byte, len(data))
	copy(c2, data)
	c2[0] = byte(0x06)

	res, err := provider.CallContract(ctx, ethereum.CallMsg{
		To:   &toAddress,
		Data: c2,
	}, nil)

	if err != nil {
		return common.Address{}, nil, err
	}

	if len(res) < 32 {
		return common.Address{}, nil, fmt.Errorf("decompressed data too short")
	}

	addr := common.BytesToAddress(res[len(res)-32:])
	return addr, res[:len(res)-32], nil
}

func TryDecodeCalldata(
	ctx context.Context,
	provider *ethrpc.Provider,
	toAddress common.Address,
	calldata []byte,
	decompressCzip bool,
) (common.Address, Transactions, *big.Int, []byte, error) {
	decodedTransactions, decodedNonce, decodedSignature, err := DecodeExecdata(calldata)
	if err == nil {
		return toAddress, decodedTransactions, decodedNonce, decodedSignature, nil
	}

	var addr common.Address
	var decompressed []byte
	var err2 error

	// Try decoding it decompressed
	if decompressCzip {
		addr, decompressed, err2 = DecompressCalldata(ctx, provider, toAddress, calldata)
		if err2 != nil {
			// Don't bubble up the decompression error, as it might not be a decompression error
			return common.Address{}, nil, nil, nil, err
		}
	}

	decodedTransactions, decodedNonce, decodedSignature, err = DecodeExecdata(decompressed)
	if err != nil {
		return common.Address{}, nil, nil, nil, err
	}

	return addr, decodedTransactions, decodedNonce, decodedSignature, nil
}

func DecodeReceipt(
	ctx context.Context,
	receipt *types.Receipt,
	provider *ethrpc.Provider,
	calldata []byte,
	decompressCzip bool,
) ([]*Receipt, []*types.Log, error) {
	if calldata == nil {
		transaction, _, err := provider.TransactionByHash(ctx, receipt.TxHash)
		if err != nil {
			return nil, nil, err
		}
		calldata = transaction.Data()
	}

	chainID, err := provider.ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}

	wallet, decodedTransactions, decodedNonce, decodedSignature, err := TryDecodeCalldata(ctx, provider, receipt.To, calldata, decompressCzip)
	if err != nil {
		return nil, nil, err
	}

	isGuestExecute := decodedNonce != nil && len(decodedSignature) == 0
	logs, receipts, err := decodeReceipt(receipt.Logs, decodedTransactions, decodedNonce, wallet, chainID, isGuestExecute)
	if err != nil {
		return nil, nil, err
	}

	for _, child := range receipts {
		child.setNativeReceipt(receipt)
	}

	return receipts, logs, nil
}

func V1IsTxExecutedEvent(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 0 &&
		len(log.Data) == 32 &&
		bytes.Equal(log.Data, hash[:])
}

func V1IsTxFailedEvent(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 1 &&
		log.Topics[0] == V1TxFailedEventSig &&
		bytes.HasPrefix(log.Data, hash[:])
}

func V2IsTxExecutedEvent(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 2 &&
		log.Topics[0] == V2TxExecutedEventSig &&
		bytes.Equal(log.Topics[1][:], hash[:])
}

func IsTxExecutedEvent(log *types.Log, hash common.Hash) bool {
	return V2IsTxExecutedEvent(log, hash)
}

func V2IsTxFailedEvent(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 2 &&
		log.Topics[0] == V2TxFailedEventSig &&
		bytes.Equal(log.Topics[1][:], hash[:])
}

func IsTxFailedEvent(log *types.Log, hash common.Hash) bool {
	return V2IsTxFailedEvent(log, hash)
}

func V1DecodeTxFailedEvent(log *types.Log) (common.Hash, string, error) {
	if len(log.Topics) != 1 || log.Topics[0] != V1TxFailedEventSig {
		return common.Hash{}, "", fmt.Errorf("not a TxFailed event")
	}

	var hash common.Hash
	var revert []byte
	if err := ethcoder.AbiDecoder([]string{"bytes32", "bytes"}, log.Data, []interface{}{&hash, &revert}); err != nil {
		return common.Hash{}, "", err
	}

	reason, err := abi.UnpackRevert(revert)
	if err != nil {
		return common.Hash{}, "", err
	}

	return hash, reason, nil
}

func V2DecodeTxFailedEvent(log *types.Log) (common.Hash, string, uint, error) {
	if len(log.Topics) != 2 || log.Topics[0] != V2TxFailedEventSig {
		return common.Hash{}, "", 0, fmt.Errorf("not a TxFailed event")
	}

	hash := common.BytesToHash(log.Topics[1][:])

	var index uint
	var revert []byte
	if err := ethcoder.AbiDecoder([]string{"uint256", "bytes"}, log.Data, []interface{}{&index, &revert}); err != nil {
		return common.Hash{}, "", 0, err
	}

	reason, err := abi.UnpackRevert(revert)
	if err != nil {
		return common.Hash{}, "", 0, err
	}

	return hash, reason, index, nil
}

func DecodeTxFailedEvent(log *types.Log) (common.Hash, string, uint, error) {
	return V2DecodeTxFailedEvent(log)
}

func DecodeNonceChangeEvent(log *types.Log) (*big.Int, *big.Int, error) {
	if len(log.Topics) != 1 || log.Topics[0] != NonceChangeEventSig {
		return nil, nil, fmt.Errorf("not a NonceChange event")
	}

	var space, nonce *big.Int
	if err := ethcoder.AbiDecoder([]string{"uint256", "uint256"}, log.Data, []interface{}{&space, &nonce}); err != nil {
		return nil, nil, err
	}

	return space, nonce, nil
}

func decodeReceipt(logs []*types.Log, transactions Transactions, nonce *big.Int, address common.Address, chainID *big.Int, isGuestExecute bool) ([]*types.Log, []*Receipt, error) {
	isSelfExecute := nonce == nil

	digestType := MetaTxnWalletExec
	if isSelfExecute {
		digestType = MetaTxnSelfExec
	} else if isGuestExecute {
		digestType = MetaTxnGuestExec
	}

	// compute the logged transaction hash
	metaTxnID, hash, err := ComputeMetaTxnID(chainID, address, transactions, nonce, digestType)
	if err != nil {
		return nil, nil, err
	}

	var topLevelLogs []*types.Log
	var receipts []*Receipt

	// check if we should expect a NonceChange event
	if !isSelfExecute && !isGuestExecute {
		if len(logs) > 0 {
			if _, _, err := DecodeNonceChangeEvent(logs[0]); err != nil {
				return nil, nil, fmt.Errorf("expected NonceChange event")
			}

			// consume the NonceChange event
			topLevelLogs = append(topLevelLogs, logs[0])
			logs = logs[1:]
		}
	}

	// create a receipt for each transaction
	for i, transaction := range transactions {
		receipt := Receipt{
			MetaTxnID:   metaTxnID,
			Status:      MetaTxnReverted,
			Index:       uint(i),
			Transaction: transaction,
		}

		for len(logs) > 0 {
			var log *types.Log
			log, logs = logs[0], logs[1:]

			isTxExecuted := V1IsTxExecutedEvent(log, hash) || IsTxExecutedEvent(log, hash)
			failedHash, failedReason, err := V1DecodeTxFailedEvent(log)
			if err != nil {
				failedHash, failedReason, _, err = V2DecodeTxFailedEvent(log)
			}

			isTxFailed := err == nil && failedHash == hash

			if isTxExecuted || isTxFailed {
				if isTxExecuted {
					receipt.Status = MetaTxnExecuted
				} else if isTxFailed {
					receipt.Status = MetaTxnFailed
					receipt.Reason = failedReason
				}

				topLevelLogs = append(topLevelLogs, log)
				break
			} else {
				receipt.Logs = append(receipt.Logs, log)
			}
		}

		if transaction.Transactions != nil {
			var err error
			receipt.Logs, receipt.Receipts, err = decodeReceipt(receipt.Logs, transaction.Transactions, transaction.Nonce, transaction.To, chainID, isGuestExecuteTransaction(transaction))
			if err != nil {
				return nil, nil, err
			}
		}

		receipts = append(receipts, &receipt)
	}

	if len(logs) > 0 {
		return nil, nil, fmt.Errorf("%v unexpected events", len(logs))
	}

	return topLevelLogs, receipts, nil
}

func isGuestExecuteTransaction(transaction *Transaction) bool {
	return transaction.Nonce != nil && len(transaction.Signature) == 0
}
