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
	"github.com/0xsequence/go-sequence/lib/logadapter"
	"github.com/goware/breaker"
	"github.com/rs/zerolog"
)

type MetaTxnListener struct {
	log      zerolog.Logger
	provider *ethrpc.Provider
	monitor  *ethmonitor.Monitor

	pastReceipts []BlockOfReceipts
	subscribers  []*subscriber

	mu sync.RWMutex
}

type MetaTxnReceiptResult struct {
	MetaTxnID  MetaTxnID
	Status     MetaTxnStatus
	TxnReceipt *types.Receipt
}

type BlockOfReceipts []*MetaTxnReceiptResult

type subscriber struct {
	ch          chan *MetaTxnReceiptResult
	unsubscribe func()
}

func NewMetaTxnListener(log zerolog.Logger, provider *ethrpc.Provider, monitor *ethmonitor.Monitor) (*MetaTxnListener, error) {
	return &MetaTxnListener{
		log:          log.With().Str("ps", "ReceiptListener").Logger(),
		provider:     provider,
		monitor:      monitor,
		pastReceipts: make([]BlockOfReceipts, 0),
		subscribers:  make([]*subscriber, 0),
	}, nil
}

func (l *MetaTxnListener) Run(ctx context.Context) error {
	sub := l.monitor.Subscribe()
	defer sub.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			l.log.Debug().Msgf("parent signaled to cancel - receipt listener is quitting")
			return nil
		case <-sub.Done():
			l.log.Info().Msgf("receipt listener is stopped because monitor signaled its stopping")
			return nil
		case blocks := <-sub.Blocks():
			block := blocks.LatestBlock().Block
			if block == nil {
				l.log.Warn().Msgf("monitor return latestblock of nil, unexpected but skipping..")
				continue
			}

			err := breaker.ExpBackoffRetry(ctx, func(ctx context.Context) error {
				return l.handleBlock(ctx, block)
			}, logadapter.Wrap(l.log), 1*time.Second, 2, 10)

			if err != nil {
				if errors.Is(err, breaker.ErrHitMaxRetries) {
					l.log.Err(err).Msgf("failed to handle block %d after many retries", block.NumberU64())
					continue
				} else {
					l.log.Err(err).Msgf("failed to handle block %d", block.NumberU64())
					continue
				}
			}

		}
	}
}

func (l *MetaTxnListener) handleBlock(ctx context.Context, block *types.Block) error {
	blockOfReceipts := BlockOfReceipts{}

	nonceChangedTopics := [][]common.Hash{{NonceChangeEventSig}}
	query := ethereum.FilterQuery{
		FromBlock: block.Number(),
		ToBlock:   new(big.Int).Add(block.Number(), common.Big1),
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
			} else {
				continue // unknown, skip
			}

			result := &MetaTxnReceiptResult{
				MetaTxnID:  metaTxnId,
				Status:     status,
				TxnReceipt: tx,
			}

			// Add found result to block of receipts
			blockOfReceipts = append(blockOfReceipts, result)
		}
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	// Tell subscribers
	for _, result := range blockOfReceipts {
		for _, sub := range l.subscribers {
			sub.ch <- result
			// TODO: can we do a timeout thing here..?
		}
	}

	l.log.Debug().
		Int("past-block-entries", len(l.pastReceipts)).
		Int("new-entries", len(blockOfReceipts)).
		Msgf("Push into past receipts")

	if len(l.pastReceipts) < 1024 {
		// Append at the end of slice
		l.pastReceipts = append(l.pastReceipts, blockOfReceipts)
	} else {
		// Append value but also pop the queue
		l.pastReceipts = append(l.pastReceipts[1:], blockOfReceipts)
	}

	return nil
}

func (l *MetaTxnListener) subscribe() *subscriber {
	l.mu.Lock()
	defer l.mu.Unlock()

	subscriber := &subscriber{
		ch: make(chan *MetaTxnReceiptResult),
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
	sub := l.subscribe()
	defer sub.unsubscribe()

	// See if metaTxn has been seen in past blocks
	l.mu.RLock()
	totalInspected := 0
	for _, bol := range l.pastReceipts {
		for _, receipt := range bol {
			totalInspected++
			if receipt.MetaTxnID == metaTxnID {
				l.log.Debug().
					Int("inspected", totalInspected).
					Str("meta-tx", string(metaTxnID)).
					Msgf("Found receipt among past receipts")

				l.mu.RUnlock()
				return receipt.Status, receipt.TxnReceipt, nil
			}
		}
	}
	l.mu.RUnlock()

	l.log.Debug().
		Int("inspected", totalInspected).
		Str("meta-tx", string(metaTxnID)).
		Msgf("Receipt not found among past receipts. Now listening..")

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
		case receipt := <-sub.ch:
			if receipt.MetaTxnID == metaTxnID {
				return receipt.Status, receipt.TxnReceipt, nil
			}
		}
	}
}
