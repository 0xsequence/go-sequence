package sequence

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/0xsequence/ethkit/ethmonitor"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"

	"github.com/rs/zerolog"
)

type MetaTxnListener struct {
	log      zerolog.Logger
	provider *ethrpc.Provider
	monitor  *ethmonitor.Monitor

	subscribers  []*subscriber
	pastReceipts []*BlockOfReceipts

	mu sync.Mutex
}

type MetaTxnReceiptResult struct {
	metaTxnId MetaTxnID
	status    MetaTxnStatus
	receipt   *types.Receipt
}

type BlockOfReceipts = []*MetaTxnReceiptResult

type subscriber struct {
	ch          chan *MetaTxnReceiptResult
	unsubscribe func()
}

func NewMetaTxnListener(log zerolog.Logger, provider *ethrpc.Provider, monitor *ethmonitor.Monitor) *MetaTxnListener {
	return &MetaTxnListener{
		log:         log,
		provider:    provider,
		monitor:     monitor,
		subscribers: make([]*subscriber, 0),
	}
}

func (l *MetaTxnListener) Subscribe() *subscriber {
	l.mu.Lock()
	defer l.mu.Unlock()

	subscriber := &subscriber{
		ch: make(chan *MetaTxnReceiptResult, 1024),
	}

	subscriber.unsubscribe = func() {
		l.mu.Lock()
		defer l.mu.Unlock()
		for i, sub := range l.subscribers {
			if sub == subscriber {
				l.subscribers = append(l.subscribers[:i], l.subscribers[i+1:]...)
				close(subscriber.ch)
				return
			}
		}
	}

	l.subscribers = append(l.subscribers, subscriber)

	return subscriber
}

func (l *MetaTxnListener) HandleBlock(ctx context.Context, block *types.Block) error {
	// Create slice for all receipts found in block
	blockOfReceipts := make(BlockOfReceipts, 0)

	nonceChangedTopics := [][]common.Hash{{NonceChangeEventSig}}
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).Sub(block.Number(), common.Big1),
		ToBlock:   block.Number(),
		Topics:    nonceChangedTopics,
	}

	// Find all nonce change events
	logs, err := l.provider.FilterLogs(ctx, query)
	if err != nil {
		return err
	}

	l.log.Debug().
		Uint64("block", block.NumberU64()).
		Int("logs", len(logs)).
		Msgf("Found logs")

	for _, log := range logs {
		// We need to find the metaTxnIds
		tx, err := l.provider.TransactionReceipt(ctx, log.TxHash)
		if err != nil {
			l.log.Warn().
				Uint64("block", block.NumberU64()).
				Str("tx", log.TxHash.Hex()).
				Err(err).
				Msgf("Error retrieving tx receipt")

			return err
		}

		// We could see multiple metaTxns on the same transaction
		for _, txLog := range tx.Logs {
			var status MetaTxnStatus
			var metaTxnId MetaTxnID

			// Success transactions have no topics and the metaTxId is the data
			// we can't really know if this is a metaTxn or not, but we assume it is
			// if it isn't is just going to get ignored
			if len(txLog.Topics) == 0 && len(txLog.Data) == 32 {
				status = MetaTxnExecuted
				metaTxnId = MetaTxnID(common.Bytes2Hex(txLog.Data))

				l.log.Debug().
					Str("tx", tx.TxHash.Hex()).
					Str("meta-tx", string(metaTxnId)).
					Msgf("Found succeed meta-tx")

				// Failed transactions have the TxFailed topic and the data begins with the metaTxInd
			} else if len(txLog.Topics) == 1 && txLog.Topics[0] == TxFailedEventSig && len(txLog.Data) >= 32 {
				status = MetaTxnExecuted
				metaTxnId = MetaTxnID(common.Bytes2Hex(txLog.Data[:32]))

				l.log.Debug().
					Str("tx", tx.TxHash.Hex()).
					Str("meta-tx", string(metaTxnId)).
					Msgf("Found failed meta-tx")
			}

			result := &MetaTxnReceiptResult{
				metaTxnId: metaTxnId,
				status:    status,
				receipt:   tx,
			}

			// Add found result to block of receipts
			blockOfReceipts = append(blockOfReceipts, result)
		}
	}

	// Store block of receipts
	// use a deque so it doesn't grow forever
	// TODO: We may as well be using the indexer database at this point
	l.mu.Lock()
	defer l.mu.Unlock()

	// Tell suscribers
	// non-blocking send
	for _, result := range blockOfReceipts {
		for _, sub := range l.subscribers {
			select {
			case sub.ch <- result:
			default:
			}
		}
	}

	l.log.Debug().
		Int("past-block-entries", len(l.pastReceipts)).
		Int("new-entries", len(blockOfReceipts)).
		Msgf("Push into past receipts")

	if len(l.pastReceipts) < 1024 {
		// Append at the end of slice
		l.pastReceipts = append(l.pastReceipts, &blockOfReceipts)
	} else {
		// Append value but also pop the queue
		l.pastReceipts = append(l.pastReceipts[1:], &blockOfReceipts)
	}

	return nil
}

func (l *MetaTxnListener) Run(ctx context.Context) error {
	oplog := l.log

	suscription := l.monitor.Subscribe()
	defer suscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			oplog.Debug().Msgf("parent signaled to cancel - listener is quitting")
			return nil
		case blocks := <-suscription.Blocks():
			block := blocks.LatestBlock().Block

			err := l.HandleBlock(ctx, block)
			errCount := 0

			for err != nil && errCount < 15 {
				l.log.Warn().
					Uint64("block", block.NumberU64()).
					Str("block-hash", block.Hash().Hex()).
					Uint("retries", uint(errCount)).
					Err(err).
					Msgf("Error processing block, schedule retry")

				err = l.HandleBlock(ctx, block)
				errCount++
				time.Sleep(5 * time.Second)
			}

			if err != nil {
				l.log.Err(err).
					Uint64("block", block.NumberU64()).
					Str("block-hash", block.Hash().Hex()).
					Uint("retries", uint(errCount)).
					Msgf("Error processing block, block dumped")
			}
		}
	}
}

func (l *MetaTxnListener) WaitForMetaTxn(ctx context.Context, metaTxnID MetaTxnID, optTimeout ...time.Duration) (MetaTxnStatus, *types.Receipt, error) {
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

	// Listen for new receipts
	sus := l.Subscribe()
	defer sus.unsubscribe()

	// See if metaTxn has been seen in past blocks
	tatalInspected := 0
	for _, bol := range l.pastReceipts {
		for _, receipt := range *bol {
			tatalInspected++
			if receipt.metaTxnId == metaTxnID {
				l.log.Debug().
					Int("inspected", tatalInspected).
					Str("meta-tx", string(metaTxnID)).
					Msgf("Found receipt among past receipts")

				return receipt.status, receipt.receipt, nil
			}
		}
	}

	l.log.Debug().
		Int("inspected", tatalInspected).
		Str("meta-tx", string(metaTxnID)).
		Msgf("Receipt not found among past receipts")

	// Wait for receipt or
	// context deadline
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.DeadlineExceeded) {
				return 0, nil, fmt.Errorf("waiting for meta transaction timeout for %v: %w", metaTxnID, err)
			} else if err != nil {
				return 0, nil, fmt.Errorf("failed waiting for meta transaction for %v: %w", metaTxnID, err)
			} else {
				return 0, nil, nil
			}
		case receipt := <-sus.ch:
			if receipt.metaTxnId == metaTxnID {
				return receipt.status, receipt.receipt, nil
			}
		}
	}
}
