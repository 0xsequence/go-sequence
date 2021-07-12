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

func (r *Receipt) setNativeReceipt(receipt *types.Receipt) {
	r.Receipt = receipt

	for _, child := range r.Receipts {
		child.setNativeReceipt(receipt)
	}
}

func DecodeReceipt(ctx context.Context, receipt *types.Receipt, provider *ethrpc.Provider) ([]*Receipt, error) {
	transaction, isPending, err := provider.TransactionByHash(ctx, receipt.TxHash)
	if err != nil {
		return nil, err
	} else if isPending {
		return nil, fmt.Errorf("transaction %v is pending", receipt.TxHash.Hex())
	}

	decodedTransactions, decodedNonce, decodedSignature, err := DecodeExecdata(transaction.Data())
	if err != nil {
		return nil, err
	}

	transactions, err := prepareTransactionsForEncoding(decodedTransactions)
	if err != nil {
		return nil, err
	}

	isGuestExecute := decodedNonce != nil && len(decodedSignature) == 0
	_, receipts, err := decodeReceipt(receipt.Logs, transactions, decodedNonce, *transaction.To(), transaction.ChainId(), isGuestExecute, "")
	if err != nil {
		return nil, err
	}

	for _, child := range receipts {
		child.setNativeReceipt(receipt)
	}

	return receipts, nil
}

func IsTxExecutedEvent(log *types.Log, hash common.Hash) bool {
	return len(log.Topics) == 0 && bytes.Equal(log.Data, hash[:])
}

func DecodeTxFailedEvent(log *types.Log) (common.Hash, string, error) {
	if len(log.Topics) != 1 || log.Topics[0] != TxFailedEventSig {
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

func decodeReceipt(logs []*types.Log, transactions Transactions, nonce *big.Int, address common.Address, chainID *big.Int, isGuestExecute bool, metaTxnID MetaTxnID) ([]*types.Log, []*Receipt, error) {
	isSelfExecute := nonce == nil

	// compute the logged transaction hash
	var hash common.Hash
	if !isGuestExecute {
		bundle := Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		digest, err := bundle.Digest()
		if err != nil {
			return nil, nil, err
		}
		subDigest, err := SubDigest(address, chainID, digest)
		if err != nil {
			return nil, nil, err
		}
		hash = common.BytesToHash(subDigest)
	} else {
		bundle := Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		digest, err := bundle.GuestDigest()
		if err != nil {
			return nil, nil, err
		}
		subDigest, err := SubDigest(address, chainID, digest)
		if err != nil {
			return nil, nil, err
		}
		hash = common.BytesToHash(subDigest)
	}

	// compute the meta-transaction ID
	if !isSelfExecute && !isGuestExecute {
		var err error
		metaTxnID, err = ComputeMetaTxnID(address, chainID, transactions, nonce)
		if err != nil {
			return nil, nil, err
		}
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

			isTxExecuted := IsTxExecutedEvent(log, hash)
			failedHash, failedReason, err := DecodeTxFailedEvent(log)
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
			receipt.Logs, receipt.Receipts, err = decodeReceipt(receipt.Logs, transaction.Transactions, transaction.Nonce, transaction.To, chainID, isGuestExecuteTransaction(transaction), metaTxnID)
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
