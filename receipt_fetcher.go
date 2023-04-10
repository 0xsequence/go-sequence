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
	receipt, waitFinality, err := receiptListener.FetchTransactionReceiptWithFilter(ctx, FilterMetaTransactionID(metaTxnHash).LimitOne(true).SearchCache(true))
	if err != nil {
		return nil, nil, nil, err
	}

	result = &MetaTxnResult{
		MetaTxnID: metaTxnID,
	}

	var isV2 bool

	for _, log := range receipt.Logs() {
		isTxExecutedV1 := IsTxExecutedEventV1(log, metaTxnHash)
		isTxFailedV1 := IsTxFailedEventV1(log, metaTxnHash)
		isTxExecutedV2 := IsTxExecutedEventV2(log, metaTxnHash)
		isTxFailedV2 := IsTxFailedEventV2(log, metaTxnHash)

		if isTxExecutedV1 || isTxFailedV1 {
			isV2 = false
			break
		} else if isTxExecutedV2 || isTxFailedV2 {
			isV2 = true
			break
		}
	}

	for _, log := range receipt.Logs() {
		var isTxExecuted, isTxFailed bool
		var reason string

		if isV2 {
			isTxExecuted = IsTxExecutedEventV2(log, metaTxnHash)
			isTxFailed = IsTxFailedEventV2(log, metaTxnHash)

			if isTxFailed {
				_, reason, _, _ = DecodeTxFailedEventV2(log)
			}
		} else {
			isTxExecuted = IsTxExecutedEventV1(log, metaTxnHash)
			isTxFailed = IsTxFailedEventV1(log, metaTxnHash)

			if isTxFailed {
				_, reason, _ = DecodeTxFailedEventV1(log)
			}
		}

		if isTxExecuted {
			result.Status = MetaTxnExecuted
			break
		} else if isTxFailed {
			result.Status = MetaTxnFailed
			result.Reason = reason
		}
	}

	return result, receipt, waitFinality, nil
}

func FilterMetaTransactionID(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			isTxExecutedV1 := IsTxExecutedEventV1(log, metaTxnID)
			isTxFailedV1 := IsTxFailedEventV1(log, metaTxnID)
			isTxExecutedV2 := IsTxExecutedEventV2(log, metaTxnID)
			isTxFailedV2 := IsTxFailedEventV2(log, metaTxnID)

			if isTxExecutedV1 || isTxFailedV1 || isTxExecutedV2 || isTxFailedV2 {
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
