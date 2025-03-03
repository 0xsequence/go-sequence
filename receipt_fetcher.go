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

	var isV2, isV3 bool

	for _, log := range receipt.Logs() {
		isTxExecutedV1 := V1IsTxExecutedEvent(log, metaTxnHash)
		isTxFailedV1 := V1IsTxFailedEvent(log, metaTxnHash)
		isTxExecutedV2 := IsTxExecutedEvent(log, metaTxnHash)
		isTxFailedV2 := V2IsTxFailedEvent(log, metaTxnHash)
		isTxExecutedV3 := V3IsTxExecutedEvent(log, metaTxnHash)
		isTxFailedV3 := V3IsTxFailedEvent(log, metaTxnHash)

		if isTxExecutedV1 || isTxFailedV1 {
			isV2 = false
			isV3 = false
			break
		} else if isTxExecutedV2 || isTxFailedV2 {
			isV2 = true
			isV3 = false
			break
		} else if isTxExecutedV3 || isTxFailedV3 {
			isV2 = false
			isV3 = true
			break
		}
	}

	for _, log := range receipt.Logs() {
		var isTxExecuted, isTxFailed bool
		var reason string

		if isV2 {
			isTxExecuted = IsTxExecutedEvent(log, metaTxnHash)
			isTxFailed = V2IsTxFailedEvent(log, metaTxnHash)

			if isTxFailed {
				_, reason, _, _ = V2DecodeTxFailedEvent(log)
			}
		} else if isV3 {
			isTxExecuted = V3IsTxExecutedEvent(log, metaTxnHash)
			isTxFailed = V3IsTxFailedEvent(log, metaTxnHash)

			if isTxFailed {
				_, reason, _, _ = V3DecodeTxFailedEvent(log)
			}
		} else {
			isTxExecuted = V1IsTxExecutedEvent(log, metaTxnHash)
			isTxFailed = V1IsTxFailedEvent(log, metaTxnHash)

			if isTxFailed {
				_, reason, _ = V1DecodeTxFailedEvent(log)
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
			isTxExecutedV1 := V1IsTxExecutedEvent(log, metaTxnID)
			isTxFailedV1 := V1IsTxFailedEvent(log, metaTxnID)
			isTxExecutedV2 := IsTxExecutedEvent(log, metaTxnID)
			isTxFailedV2 := V2IsTxFailedEvent(log, metaTxnID)
			isTxExecutedV3 := V3IsTxExecutedEvent(log, metaTxnID)
			isTxFailedV3 := V3IsTxFailedEvent(log, metaTxnID)

			if isTxExecutedV1 || isTxFailedV1 || isTxExecutedV2 || isTxFailedV2 || isTxExecutedV3 || isTxFailedV3 {
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
			if len(log.Topics) == 1 && (log.Topics[0] == V1TxFailedEventSig || log.Topics[0] == V2TxFailedEventSig || log.Topics[0] == V3CallFailed) {
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
