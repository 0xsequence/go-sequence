package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
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

	decoded, err := DecodeTransaction(transaction.Data())
	if err != nil {
		return nil, err
	}

	transactions, err := prepareTransactionsForEncoding(decoded.Transactions)
	if err != nil {
		return nil, err
	}

	_, receipts, err := decodeReceipt(receipt.Logs, transactions, decoded.Nonce, *transaction.To(), transaction.ChainId(), isGuestExecuteTransaction(decoded), "")
	if err != nil {
		return nil, err
	}

	for _, child := range receipts {
		child.setNativeReceipt(receipt)
	}

	return receipts, nil
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
