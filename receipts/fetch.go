package receipts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	sequence "github.com/0xsequence/go-sequence"
)

// FetchReceipts looks up the transaction that emitted Call* events for the given
// digest and returns all decoded Sequence receipts along with the native receipt.
func FetchReceipts(ctx context.Context, digest common.Hash, provider *ethrpc.Provider, fromBlock, toBlock *big.Int) (Receipts, *types.Receipt, error) {
	if provider == nil {
		return Receipts{}, nil, fmt.Errorf("no provider")
	}

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics: [][]common.Hash{
			{sequence.V3CallSucceeded, sequence.V3CallFailed, sequence.V3CallAborted, sequence.V3CallSkipped},
			{digest},
		},
	}

	logs, err := provider.FilterLogs(ctx, query)
	if err != nil {
		return Receipts{}, nil, fmt.Errorf("unable to filter logs: %w", err)
	}
	if len(logs) == 0 {
		// Fallback for legacy events where the digest is not indexed.
		query.Topics = [][]common.Hash{{sequence.V3CallSucceeded, sequence.V3CallFailed, sequence.V3CallAborted, sequence.V3CallSkipped}}
		logs, err = provider.FilterLogs(ctx, query)
		if err != nil {
			return Receipts{}, nil, fmt.Errorf("unable to filter logs without digest topic: %w", err)
		}
	}

	log, err := findDigestLog(logs, digest)
	if err != nil {
		return Receipts{}, nil, err
	}

	receipt, err := provider.TransactionReceipt(ctx, log.TxHash)
	if err != nil {
		return Receipts{}, nil, fmt.Errorf("unable to get transaction receipt %v: %w", log.TxHash, err)
	}

	decoded, err := TransactionReceiptsForReceipt(ctx, receipt, provider)
	if err != nil {
		return Receipts{}, receipt, fmt.Errorf("unable to decode transaction receipt %v: %w", receipt.TxHash, err)
	}

	receipts := decoded.Find(digest)
	if receipts == nil {
		return Receipts{}, receipt, fmt.Errorf("decoded receipts do not include digest %v", digest)
	}

	return *receipts, receipt, nil
}

func findDigestLog(logs []types.Log, digest common.Hash) (*types.Log, error) {
	var selected *types.Log

	for i := range logs {
		log := &logs[i]

		if !matchesDigest(log, digest) {
			continue
		}

		if selected == nil || isNewerLog(log, selected) {
			selected = log
		}
	}

	if selected == nil {
		return nil, fmt.Errorf("no Call* events found for digest %v", digest)
	}

	return selected, nil
}

func matchesDigest(log *types.Log, digest common.Hash) bool {
	if hash, _, err := sequence.V3DecodeCallSucceededEvent(log); err == nil && hash == digest {
		return true
	}
	if hash, _, _, err := sequence.V3DecodeCallFailedEvent(log); err == nil && hash == digest {
		return true
	}
	if hash, _, _, err := sequence.V3DecodeCallAbortedEvent(log); err == nil && hash == digest {
		return true
	}
	if hash, _, err := sequence.V3DecodeCallSkippedEvent(log); err == nil && hash == digest {
		return true
	}
	return false
}

func isNewerLog(a, b *types.Log) bool {
	if a.BlockNumber != b.BlockNumber {
		return a.BlockNumber > b.BlockNumber
	}
	if a.TxIndex != b.TxIndex {
		return a.TxIndex > b.TxIndex
	}
	return a.Index > b.Index
}
