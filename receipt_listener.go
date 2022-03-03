package sequence

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/0xsequence/ethkit/ethmonitor"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/lib/logadapter"
	"github.com/goware/breaker"
	"github.com/rs/zerolog"
)

type ReceiptListener struct {
	log      zerolog.Logger
	provider *ethrpc.Provider
	monitor  *ethmonitor.Monitor

	pastReceipts []BlockOfReceipts
	subscribers  []*subscriber

	mu sync.Mutex
}

type ReceiptResult struct {
	MetaTxnID  MetaTxnID
	Status     MetaTxnStatus
	TxnReceipt *types.Receipt
}

type BlockOfReceipts []*ReceiptResult

type subscriber struct {
	ch          chan *ReceiptResult
	done        chan struct{}
	unsubscribe func()
}

func NewReceiptListener(log zerolog.Logger, provider *ethrpc.Provider, monitor *ethmonitor.Monitor) (*ReceiptListener, error) {
	if !monitor.Options().WithLogs {
		return nil, fmt.Errorf("ReceiptListener needs a monitor with WithLogs enabled to function")
	}

	return &ReceiptListener{
		log:          log.With().Str("ps", "ReceiptListener").Logger(),
		provider:     provider,
		monitor:      monitor,
		pastReceipts: make([]BlockOfReceipts, 0),
		subscribers:  make([]*subscriber, 0),
	}, nil
}

func (l *ReceiptListener) Run(ctx context.Context) error {
	sub := l.monitor.Subscribe()
	defer sub.Unsubscribe()

	br := breaker.New(logadapter.Wrap(l.log), 1*time.Second, 2, 10)

	for {
		select {

		case <-ctx.Done():
			l.log.Debug().Msgf("parent signaled to cancel - receipt listener is quitting")
			return nil

		case <-sub.Done():
			l.log.Info().Msgf("receipt listener is stopped because monitor signaled its stopping")
			return nil

		case blocks := <-sub.Blocks():
			for _, block := range blocks {
				if block.Event != ethmonitor.Added {
					continue
				}

				err := br.Do(ctx, func() error {
					return l.handleBlock(ctx, block)
				})
				if err != nil {
					if errors.Is(err, breaker.ErrHitMaxRetries) {
						l.log.Err(err).Msgf("failed to handle block %d after many retries", block.NumberU64())
					} else {
						l.log.Err(err).Msgf("failed to handle block %d", block.NumberU64())
					}
				}
			}
		}
	}
}

func (l *ReceiptListener) handleBlock(ctx context.Context, block *ethmonitor.Block) error {
	blockOfReceipts := BlockOfReceipts{}
	transactionLogs := map[common.Hash][]*types.Log{}
	var nonceChangeTransactionHashes []common.Hash

	// categorize logs by transaction hash
	// also note transactions with NonceChange events
	for i := range block.Logs {
		log := &block.Logs[i]
		transactionLogs[log.TxHash] = append(transactionLogs[log.TxHash], log)

		if len(log.Topics) == 1 && log.Topics[0] == NonceChangeEventSig {
			nonceChangeTransactionHashes = append(nonceChangeTransactionHashes, log.TxHash)
		}
	}

	l.log.Debug().
		Uint64("block", block.NumberU64()).
		Int("logs", len(block.Logs)).
		Msgf("Found logs")

	for _, txHash := range nonceChangeTransactionHashes {
		tx, err := l.provider.TransactionReceipt(ctx, txHash)
		if tx == nil || err != nil {
			l.log.Warn().
				Uint64("block", block.NumberU64()).
				Str("tx", txHash.Hex()).
				Err(err).
				Msgf("Error retrieving tx receipt")

			return err
		}

		// We could see multiple metaTxns on the same transaction
		for _, txLog := range transactionLogs[txHash] {
			var status MetaTxnStatus
			var metaTxnID MetaTxnID

			// Success transactions have no topics and the metaTxId is the data
			// we can't really know if this is a metaTxn or not, but we assume it is
			// if it isn't is just going to get ignored
			if len(txLog.Topics) == 0 && len(txLog.Data) == 32 {
				status = MetaTxnExecuted
				metaTxnID = MetaTxnID(common.Bytes2Hex(txLog.Data))

				l.log.Debug().
					Str("tx", txHash.Hex()).
					Str("meta-tx", string(metaTxnID)).
					Msgf("Found succeed meta-tx")

				// Failed transactions have the TxFailed topic and the data begins with the metaTxInd
			} else if len(txLog.Topics) == 1 && txLog.Topics[0] == TxFailedEventSig && len(txLog.Data) >= 32 {
				status = MetaTxnFailed
				metaTxnID = MetaTxnID(common.Bytes2Hex(txLog.Data[:32]))

				l.log.Debug().
					Str("tx", txHash.Hex()).
					Str("meta-tx", string(metaTxnID)).
					Msgf("Found failed meta-tx")
			} else {
				continue // unknown, skip
			}

			result := &ReceiptResult{
				MetaTxnID:  metaTxnID,
				Status:     status,
				TxnReceipt: tx,
			}

			// Add found result to block of receipts
			blockOfReceipts = append(blockOfReceipts, result)
		}
	}

	// Nothing to record, skipping.
	if len(blockOfReceipts) == 0 {
		return nil
	}

	// Publish to subscribers
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, result := range blockOfReceipts {
		for _, sub := range l.subscribers {
			select {
			case <-sub.done:
			case sub.ch <- result:
			case <-time.After(2 * time.Second):
				l.log.Warn().Msgf("channel publisher is blocked by a slow subscriber")
			}
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

func (l *ReceiptListener) subscribe() *subscriber {
	l.mu.Lock()
	defer l.mu.Unlock()

	subscriber := &subscriber{
		ch:   make(chan *ReceiptResult, 1),
		done: make(chan struct{}),
	}

	subscriber.unsubscribe = func() {
		close(subscriber.done)
		l.mu.Lock()
		defer l.mu.Unlock()
		close(subscriber.ch)
		for i, sub := range l.subscribers {
			if sub == subscriber {
				l.subscribers = append(l.subscribers[:i], l.subscribers[i+1:]...)
				return
			}
		}
	}

	l.subscribers = append(l.subscribers, subscriber)

	return subscriber
}

func (l *ReceiptListener) WaitForMetaTxn(ctx context.Context, metaTxnID MetaTxnID, optTimeout ...time.Duration) (MetaTxnStatus, *types.Receipt, error) {
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
	totalInspected := 0
	for _, bol := range l.pastReceipts {
		for _, receipt := range bol {
			totalInspected++
			if receipt.MetaTxnID == metaTxnID {
				l.log.Debug().
					Int("inspected", totalInspected).
					Str("meta-tx", string(metaTxnID)).
					Msgf("Found receipt among past receipts")

				return receipt.Status, receipt.TxnReceipt, nil
			}
		}
	}

	l.log.Debug().
		Int("inspected", totalInspected).
		Str("meta-tx", string(metaTxnID)).
		Msgf("Receipt not found among past receipts. Now listening..")

	// Wait for receipt or context deadline
	var receipt *ReceiptResult
	var err error

	for done := false; !done; {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				err = fmt.Errorf("waiting for meta transaction timeout for %v: %w", metaTxnID, ctx.Err())
			} else if ctx.Err() != nil {
				err = fmt.Errorf("failed waiting for meta transaction for %v: %w", metaTxnID, ctx.Err())
			}
			done = true

		case <-sub.done:
			done = true

		case receipt = <-sub.ch:
			if receipt.MetaTxnID == metaTxnID {
				done = true
			}
		}
	}

	if err != nil {
		return 0, nil, err
	}
	if receipt != nil {
		return receipt.Status, receipt.TxnReceipt, nil
	}
	return 0, nil, nil
}
