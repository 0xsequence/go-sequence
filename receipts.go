package sequence

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
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

func DecodeReceipt(ctx context.Context, receipt *types.Receipt, provider *ethrpc.Provider) ([]*Receipt, []*types.Log, error) {
	transaction, _, err := provider.TransactionByHash(ctx, receipt.TxHash)
	if err != nil {
		return nil, nil, err
	}

	decodedTransactions, decodedNonce, decodedSignature, err := DecodeExecdata(transaction.Data())
	if err != nil {
		return nil, nil, err
	}

	isGuestExecute := decodedNonce != nil && len(decodedSignature) == 0
	logs, receipts, err := decodeReceipt(receipt.Logs, decodedTransactions, decodedNonce, *transaction.To(), transaction.ChainId(), isGuestExecute)
	if err != nil {
		return nil, nil, err
	}

	for _, child := range receipts {
		child.setNativeReceipt(receipt)
	}

	return receipts, logs, nil
}

func IsTxExecutedEventV1(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 0 &&
		len(log.Data) == 32 &&
		bytes.Equal(log.Data, hash[:])
}

func IsTxFailedEventV1(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 1 &&
		log.Topics[0] == TxFailedEventSigV1 &&
		bytes.HasPrefix(log.Data, hash[:])
}

func IsTxExecutedEventV2(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 2 &&
		log.Topics[0] == TxExecutedEventSigV2 &&
		bytes.Equal(log.Topics[1][:], hash[:])
}

func IsTxFailedEventV2(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 2 &&
		log.Topics[0] == TxFailedEventSigV2 &&
		bytes.Equal(log.Topics[1][:], hash[:])
}

func DecodeTxFailedEventV1(log *types.Log) (common.Hash, string, error) {
	if len(log.Topics) != 1 || log.Topics[0] != TxFailedEventSigV1 {
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

func DecodeTxFailedEventV2(log *types.Log) (common.Hash, string, uint, error) {
	if len(log.Topics) != 2 || log.Topics[0] != TxFailedEventSigV2 {
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

			isTxExecuted := IsTxExecutedEventV1(log, hash) || IsTxExecutedEventV2(log, hash)
			failedHash, failedReason, err := DecodeTxFailedEventV1(log)
			if err != nil {
				failedHash, failedReason, _, err = DecodeTxFailedEventV2(log)
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
