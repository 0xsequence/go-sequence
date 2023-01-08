package sequence

import (
	"context"
	"time"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
)

func WaitForMetaTxn2(ctx context.Context, metaTxnID MetaTxnID, optTimeout ...time.Duration) ([]*MetaTxnResult, *types.Receipt, error) {
	// Use optional timeout if passed, otherwise use deadline on the provided ctx, or finally,
	// set a default timeout of 120 seconds.
	var cancel context.CancelFunc
	if len(optTimeout) > 0 {
		ctx, cancel = context.WithTimeout(ctx, optTimeout[0])
		defer cancel()
	} else {
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
			defer cancel()
		}
	}

	return nil, nil, nil
}

func FilterMetaTransactionID(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			if len(log.Data) != 32 {
				continue
			}
			if common.BytesToHash(log.Data) != metaTxnID {
				continue
			}
			isTxExecuted := IsTxExecutedEvent(log, metaTxnID)
			isTxFailed := IsTxFailedEvent(log, metaTxnID)
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
			if len(log.Topics) == 1 && log.Topics[0] == TxFailedEventSig {
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
