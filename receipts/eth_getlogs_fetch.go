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

const MaxFilterBlockRange = 7500

// FetchMetaTransactionReceiptByETHGetLogs looks up the transaction that emitted Call*
// events for the given digest and returns all decoded Sequence receipts along with
// the native receipt.
//
// The `opHash` is also known as the "MetaTxnID" but please make sure
// it has the "0x" prefix when passing as a common.Hash, even though
// sometimes we represent a MetaTxnID without the 0x prefix as a string.
//
// NOTE: toBlock can also be nil, in which case the latest block is used.
//
// Finally, please note that this method will not find meta-transactions if there is a
// native transaction which fails. In that case, the Call* events are not emitted and
// thus cannot be found. In such cases, you will need to look up the native transaction
// receipt directly. However, this is not to be confused with where a "Call" inside
// of the native transaction fails, but the native transaction itself succeeds which
// is more common and works fine.
func FetchMetaTransactionReceiptByETHGetLogs(ctx context.Context, opHash common.Hash, provider ethrpc.Interface, fromBlock, toBlock *big.Int) (Receipts, *types.Receipt, error) {
	if provider == nil {
		return Receipts{}, nil, fmt.Errorf("receipts: no provider")
	}

	fromBlock_ := fromBlock
	if fromBlock_ == nil {
		fromBlock_ = new(big.Int)
	}

	toBlock_ := toBlock
	if toBlock_ == nil {
		latest, err := provider.BlockNumber(ctx)
		if err != nil {
			return Receipts{}, nil, fmt.Errorf("unable to get latest block: %w", err)
		}
		toBlock_ = new(big.Int).SetUint64(latest)
	}

	if new(big.Int).Sub(toBlock_, fromBlock_).Cmp(big.NewInt(MaxFilterBlockRange)) > 0 {
		return Receipts{}, nil, fmt.Errorf("block range %v to %v exceeds %v, reduce block range", fromBlock_, toBlock_, MaxFilterBlockRange)
	}

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topics: [][]common.Hash{
			{sequence.V3CallSucceeded, sequence.V3CallFailed}, //, sequence.V3CallAborted, sequence.V3CallSkipped},
			{opHash},
		},
	}

	// TODO: what if there is a node failure here, we should retry a few times
	// and also check the type of error from the node to see how we should behave/handle.
	logs, err := provider.FilterLogs(ctx, query)
	if err != nil {
		return Receipts{}, nil, fmt.Errorf("unable to filter logs: %w", err)
	}
	if len(logs) == 0 {
		return Receipts{}, nil, ethereum.NotFound
	}

	log, err := findV3CallsDigestLog(logs, opHash)
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

	receipts := decoded.Find(opHash)
	if receipts == nil {
		return Receipts{}, receipt, fmt.Errorf("decoded receipts do not include digest %v", opHash)
	}
	return *receipts, receipt, nil
}

func findV3CallsDigestLog(logs []types.Log, digest common.Hash) (*types.Log, error) {
	var selected *types.Log
	for i := range logs {
		log := &logs[i]
		if !matchesV3CallsDigest(log, digest) {
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

func matchesV3CallsDigest(log *types.Log, digest common.Hash) bool {
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
