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

	logs := receipt.Logs()
done:
	for i := len(logs) - 1; i >= 0; i-- {
		log := logs[i]
		switch {
		case V1IsTxExecutedEvent(log, metaTxnHash):
			result.Status = MetaTxnExecuted
			break done
		case V2IsTxExecutedEvent(log, metaTxnHash):
			result.Status = MetaTxnExecuted
			break done
		case V3IsCallSuccessEvent(log, metaTxnHash):
			result.Status = MetaTxnExecuted
			break done
		case V1IsTxFailedEvent(log, metaTxnHash):
			result.Status = MetaTxnFailed
			_, result.Reason, _ = V1DecodeTxFailedEvent(log)
			break done
		case V2IsTxFailedEvent(log, metaTxnHash):
			result.Status = MetaTxnFailed
			_, result.Reason, _, _ = V2DecodeTxFailedEvent(log)
			break done
		case V3IsCallFailedEvent(log, metaTxnHash):
			result.Status = MetaTxnFailed
			_, _, result.Reason, _ = V3DecodeCallFailedEvent(log)
			break done
		case V3IsCallAbortedEvent(log, metaTxnHash):
			result.Status = MetaTxnFailed
			_, _, result.Reason, _ = V3DecodeCallAbortedEvent(log)
			break done
		}
	}

	return result, receipt, waitFinality, nil
}

func FilterMetaTransactionID(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			switch {
			case V1IsTxExecutedEvent(log, metaTxnID):
				fallthrough
			case V1IsTxFailedEvent(log, metaTxnID):
				fallthrough
			case V2IsTxExecutedEvent(log, metaTxnID):
				fallthrough
			case V2IsTxFailedEvent(log, metaTxnID):
				fallthrough
			case V3IsCallSuccessEvent(log, metaTxnID):
				fallthrough
			case V3IsCallFailedEvent(log, metaTxnID):
				fallthrough
			case V3IsCallAbortedEvent(log, metaTxnID):
				fallthrough
			case V3IsCallSkippedEvent(log, metaTxnID):
				return true
			}
		}
		return false
	})
}

// Find any Sequence meta txns
func FilterMetaTransactionAny() ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			if len(log.Topics) == 1 && log.Topics[0] == NonceChangeEventSig {
				return true
			}
		}
		return false
	})
}
