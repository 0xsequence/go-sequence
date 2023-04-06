package sequence

import (
	"context"
	"time"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
)

type MetaTxnResult struct {
	MetaTxnID MetaTxnID
	Status    MetaTxnStatus
	Reason    string
}

func FetchMetaTransactionReceipt(ctx context.Context, receiptListener *ethreceipts.ReceiptsListener, metaTxnID MetaTxnID, optTimeout ...time.Duration) (*MetaTxnResult, *ethreceipts.Receipt, ethreceipts.WaitReceiptFinalityFunc, error) {
	// Use optional timeout if passed, otherwise use deadline on the provided ctx, or finally,
	// set a default timeout of 200 seconds.
	var cancel context.CancelFunc
	if len(optTimeout) > 0 {
		ctx, cancel = context.WithTimeout(ctx, optTimeout[0])
		defer cancel()
	} else {
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, 200*time.Second)
			defer cancel()
		}
	}

	var result *MetaTxnResult

	metaTxnHash := common.HexToHash(string(metaTxnID))
	receipt, waitFinality, err := receiptListener.FetchTransactionReceiptWithFilter(ctx, FilterMetaTransactionIDv2(metaTxnHash).LimitOne(true).SearchCache(true))
	if err == nil {
		// found v2, decode v2
		result = &MetaTxnResult{
			MetaTxnID: metaTxnID,
		}

		for _, log := range receipt.Logs() {
			isTxExecuted := IsTxExecutedEventV2(log, metaTxnHash)
			isTxFailed := IsTxFailedEventV2(log, metaTxnHash)

			if isTxExecuted {
				result.Status = MetaTxnExecuted
				break
			} else if isTxFailed {
				result.Status = MetaTxnFailed
				_, reason, _, _ := DecodeTxFailedEventV2(log)
				result.Reason = reason
			}
		}

	} else {
		// fallback to v1
		receipt, waitFinality, err = receiptListener.FetchTransactionReceiptWithFilter(ctx, FilterMetaTransactionIDv1(metaTxnHash).LimitOne(true).SearchCache(true))
		if err != nil {
			return nil, nil, nil, err
		}

		result = &MetaTxnResult{
			MetaTxnID: metaTxnID,
		}

		for _, log := range receipt.Logs() {
			isTxExecuted := IsTxExecutedEventV1(log, metaTxnHash)
			isTxFailed := IsTxFailedEventV1(log, metaTxnHash)
			if isTxExecuted {
				result.Status = MetaTxnExecuted
				break
			} else if isTxFailed {
				result.Status = MetaTxnFailed
				_, reason, _ := DecodeTxFailedEventV1(log)
				result.Reason = reason
			}
		}
	}

	return result, receipt, waitFinality, nil
}

func FilterMetaTransactionIDv1(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			isTxExecuted := IsTxExecutedEventV1(log, metaTxnID)
			isTxFailed := IsTxFailedEventV1(log, metaTxnID)
			if isTxExecuted || isTxFailed {
				// found the sequence meta txn
				return true
			}
		}
		return false
	})
}

func FilterMetaTransactionIDv2(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			isTxExecuted := IsTxExecutedEventV2(log, metaTxnID)
			isTxFailed := IsTxFailedEventV2(log, metaTxnID)
			if isTxExecuted || isTxFailed {
				// found the sequence meta txn
				return true
			}
		}
		return false
	})
}

// Find any Sequence meta txns
func FilterMetaTransactionAny() ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		foundNonceEvent := false
		for _, log := range logs {
			if len(log.Topics) > 0 && log.Topics[0] == NonceChangeEventSig {
				foundNonceEvent = true
				break
			}
		}
		if !foundNonceEvent {
			return false
		}

		for _, log := range logs {
			if len(log.Topics) == 1 && log.Topics[0] == TxFailedEventSigV1 {
				// failed sequence txn
				return true
			} else if len(log.Topics) == 0 && len(log.Data) == 32 {
				// possibly a successful sequence txn -- but not for certain
				return true
			}
		}

		return false
	})
}
