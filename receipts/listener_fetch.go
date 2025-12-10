package receipts

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	sequence "github.com/0xsequence/go-sequence"
)

type MetaTxnResult struct {
	MetaTxnID sequence.MetaTxnID
	Status    sequence.MetaTxnStatus
	Reason    string
}

func FetchMetaTransactionReceipt(ctx context.Context, receiptListener *ethreceipts.ReceiptsListener, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (*MetaTxnResult, *ethreceipts.Receipt, ethreceipts.WaitReceiptFinalityFunc, error) {
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

	// Build the filter request for the receipt listener. Once the request is processing
	// it will immediately try the "QueryOnChain" function to see if the receipt is already
	// available on-chain before waiting for new blocks to be mined.
	filter := FilterMetaTransactionID(metaTxnHash).
		LimitOne(true).
		MaxWait(0).
		SearchCache(true).
		QueryOnChainTxnHash(false).
		QueryOnChain(func(ctx context.Context) (*types.Receipt, error) {
			// QueryOnChain function is called at the beginning of the subscriber filter
			// process, so we can quickly query the chain for the receipt without waiting
			// to check historical blocks up front.
			toBlock := receiptListener.LatestBlockNum()
			if toBlock == nil || toBlock.Cmp(big.NewInt(0)) == 0 {
				return nil, fmt.Errorf("receipts: no latest block number available from monitor")
			}
			fromBlock := new(big.Int).Sub(toBlock, big.NewInt(MaxFilterBlockRange))
			if fromBlock.Cmp(big.NewInt(0)) < 0 {
				fromBlock = big.NewInt(0)
			}
			_, receipt, err := FetchMetaTransactionReceiptByETHGetLogs(ctx, metaTxnHash, receiptListener.RPCProvider(), fromBlock, toBlock)
			return receipt, err
		})

	// Query the receipt listener
	receipt, waitFinality, err := receiptListener.FetchTransactionReceiptWithFilter(ctx, filter)
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
		case sequence.V3IsCallSucceededEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnExecuted
			break done
		case sequence.V3IsCallFailedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnFailed
			_, _, reason, _ := sequence.V3DecodeCallFailedEvent(log)
			result.Reason = fmt.Sprint(reason)
			break done
		case sequence.V3IsCallAbortedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnFailed
			_, _, reason, _ := sequence.V3DecodeCallAbortedEvent(log)
			result.Reason = fmt.Sprint(reason)
			break done
		case sequence.V2IsTxExecutedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnExecuted
			break done
		case sequence.V2IsTxFailedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnFailed
			_, result.Reason, _, _ = sequence.V2DecodeTxFailedEvent(log)
			break done
		case sequence.V1IsTxExecutedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnExecuted
			break done
		case sequence.V1IsTxFailedEvent(log, metaTxnHash):
			result.Status = sequence.MetaTxnFailed
			_, result.Reason, _ = sequence.V1DecodeTxFailedEvent(log)
			break done
		}
	}

	return result, receipt, waitFinality, nil
}

func FilterMetaTransactionID(metaTxnID ethkit.Hash) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			switch {
			case sequence.V3IsCallSucceededEvent(log, metaTxnID):
				return true
			case sequence.V3IsCallFailedEvent(log, metaTxnID):
				return true
			case sequence.V3IsCallAbortedEvent(log, metaTxnID):
				return true
			case sequence.V3IsCallSkippedEvent(log, metaTxnID):
				return true
			case sequence.V2IsTxExecutedEvent(log, metaTxnID):
				return true
			case sequence.V2IsTxFailedEvent(log, metaTxnID):
				return true
			case sequence.V1IsTxExecutedEvent(log, metaTxnID):
				return true
			case sequence.V1IsTxFailedEvent(log, metaTxnID):
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
			if len(log.Topics) == 1 && log.Topics[0] == sequence.NonceChangeEventSig {
				return true
			}
		}
		return false
	})
}
