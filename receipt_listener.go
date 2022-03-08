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

const (
	maxConcurrentFetchReceipts = 30
	pastReceiptsBufSize        = 4096
)

type ReceiptListener struct {
	log      zerolog.Logger
	provider *ethrpc.Provider
	monitor  *ethmonitor.Monitor
	br       *breaker.Breaker

	receiptsSem chan struct{}

	pastReceipts   []BlockOfReceipts
	muPastReceipts sync.Mutex

	subscribers   []*subscriber
	muSubscribers sync.Mutex
}

type ReceiptResult struct {
	MetaTxnID  MetaTxnID
	Results    []*MetaTxnResult
	TxnReceipt *types.Receipt
}

type MetaTxnResult struct {
	Status MetaTxnStatus
	Reason string
}

type BlockOfReceipts []ReceiptResult

type subscriber struct {
	ch          <-chan ReceiptResult
	sendCh      chan<- ReceiptResult
	done        chan struct{}
	unsubscribe func()
}

func NewReceiptListener(log zerolog.Logger, provider *ethrpc.Provider, monitor *ethmonitor.Monitor) (*ReceiptListener, error) {
	if !monitor.Options().WithLogs {
		return nil, fmt.Errorf("ReceiptListener needs a monitor with WithLogs enabled to function")
	}

	log = log.With().Str("ps", "ReceiptListener").Logger()

	return &ReceiptListener{
		log:          log,
		provider:     provider,
		monitor:      monitor,
		br:           breaker.New(logadapter.Wrap(log), time.Second, 2, 10),
		receiptsSem:  make(chan struct{}, maxConcurrentFetchReceipts),
		pastReceipts: make([]BlockOfReceipts, 0),
		subscribers:  make([]*subscriber, 0),
	}, nil
}

func (l *ReceiptListener) Run(ctx context.Context) error {
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
			for _, block := range blocks {
				l.handleBlock(ctx, block)
			}
		}
	}
}

func (l *ReceiptListener) handleBlock(ctx context.Context, block *ethmonitor.Block) {
	if block.Event != ethmonitor.Added {
		return
	}

	// txHashes: the set of native transactions with at least one NonceChange event
	// txLogs: block's logs indexed by native transaction hash

	txHashes := map[common.Hash]struct{}{}
	txLogs := map[common.Hash][]*types.Log{}

	for i := range block.Logs {
		log := &block.Logs[i]

		txLogs[log.TxHash] = append(txLogs[log.TxHash], log)

		if len(log.Topics) == 1 && log.Topics[0] == NonceChangeEventSig {
			txHashes[log.TxHash] = struct{}{}
		}
	}

	// receipts: meta transaction receipts indexed by native transaction hash
	// receipts will have their TxnReceipt fields populated in a different goroutine

	receipts := make(map[common.Hash][]ReceiptResult, len(txHashes))

	for txHash := range txHashes {
		logs := txLogs[txHash]

		// metaTxnReceipts: meta transaction receipts indexed by meta transaction ID
		// receipt: gets the current receipt for a given meta transaction ID

		// these get mutated and must be *ReceiptResult
		metaTxnReceipts := map[common.Hash]*ReceiptResult{}

		receipt := func(metaTxnID common.Hash) *ReceiptResult {
			result := metaTxnReceipts[metaTxnID]
			if result == nil {
				result = &ReceiptResult{MetaTxnID: MetaTxnID(metaTxnID.Hex()[2:])}
				metaTxnReceipts[metaTxnID] = result
			}
			return result
		}

		// assuming block.Logs is sorted
		for _, log := range logs {
			if len(log.Topics) == 0 && len(log.Data) == 32 {
				// possible TxExecuted event

				r := receipt(common.BytesToHash(log.Data))
				r.Results = append(r.Results, &MetaTxnResult{
					Status: MetaTxnExecuted,
				})
			} else if len(log.Topics) == 1 && log.Topics[0] == TxFailedEventSig {
				// definite TxFailed event

				metaTxnID, reason, err := DecodeTxFailedEvent(log)
				if err != nil {
					l.log.Err(err).Msgf("unable to decode TxFailed event: topics=%v data=%v", log.Topics, log.Data)
					continue
				}

				r := receipt(metaTxnID)
				r.Results = append(r.Results, &MetaTxnResult{
					Status: MetaTxnFailed,
					Reason: reason,
				})
			}
		}

		receipts[txHash] = make([]ReceiptResult, 0, len(metaTxnReceipts))
		for _, receipt := range metaTxnReceipts {
			receipts[txHash] = append(receipts[txHash], *receipt)
		}
	}

	for txHash, txReceipts := range receipts {
		l.handleReceipts(ctx, txHash, txReceipts)
	}
}

func (l *ReceiptListener) handleReceipts(ctx context.Context, txHash common.Hash, txReceipts []ReceiptResult) {
	if len(txReceipts) == 0 {
		return
	}

	// handle receipts in an independent goroutine so that node failures don't stall the block handler
	go func() {
		err := l.br.Do(ctx, func() error {
			l.receiptsSem <- struct{}{}
			defer func() {
				<-l.receiptsSem
			}()

			receipt, err := l.provider.TransactionReceipt(ctx, txHash)
			if err != nil {
				return fmt.Errorf("unable to fetch receipt for %v: %w", txHash.Hex(), err)
			} else if receipt == nil {
				return fmt.Errorf("unable to fetch receipt for %v", txHash.Hex())
			}

			// populate TxnReceipt field
			for _, txReceipt := range txReceipts {
				txReceipt.TxnReceipt = receipt
			}

			return nil
		})
		if err != nil {
			l.log.Warn().Err(err).Msgf("failed to fetch receipt after several tries")
		}

		l.pushReceipts(txReceipts)

		l.muSubscribers.Lock()
		defer l.muSubscribers.Unlock()

		// broadcast receipts
		for _, txReceipt := range txReceipts {
			for _, sub := range l.subscribers {
				select {
				case <-sub.done:
				case sub.sendCh <- txReceipt:
				}
			}
		}
	}()
}

func (l *ReceiptListener) pushReceipts(txReceipts []ReceiptResult) {
	l.muPastReceipts.Lock()
	defer l.muPastReceipts.Unlock()

	if len(l.pastReceipts) < pastReceiptsBufSize {
		l.pastReceipts = append(l.pastReceipts, txReceipts)
	} else {
		l.pastReceipts = append(l.pastReceipts[1:], txReceipts)
	}
}

func (l *ReceiptListener) subscribe() *subscriber {
	l.muSubscribers.Lock()
	defer l.muSubscribers.Unlock()

	ch := make(chan ReceiptResult)
	subscriber := &subscriber{
		ch:     ch,
		sendCh: makeUnboundedBuffered(ch, logadapter.Wrap(l.log), 100),
		done:   make(chan struct{}),
	}

	subscriber.unsubscribe = func() {
		close(subscriber.done)
		l.muSubscribers.Lock()
		defer l.muSubscribers.Unlock()
		close(subscriber.sendCh)
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

// converts a blocking unbuffered send channel into a non-blocking unbounded buffered one
// inspired by https://medium.com/capital-one-tech/building-an-unbounded-channel-in-go-789e175cd2cd
func makeUnboundedBuffered(sendCh chan<- ReceiptResult, log logadapter.Logger, bufferLimitWarning int) chan<- ReceiptResult {
	ch := make(chan ReceiptResult)

	go func() {
		var buffer []ReceiptResult

		for {
			if len(buffer) == 0 {
				if message, ok := <-ch; ok {
					buffer = append(buffer, message)
					if len(buffer) > bufferLimitWarning {
						log.Warnf("channel buffer holds %v > %v messages", len(buffer), bufferLimitWarning)
					}
				} else {
					close(sendCh)
					break
				}
			} else {
				select {
				case sendCh <- buffer[0]:
					buffer = buffer[1:]

				case message, ok := <-ch:
					if ok {
						buffer = append(buffer, message)
						if len(buffer) > bufferLimitWarning {
							log.Warnf("channel buffer holds %v > %v messages", len(buffer), bufferLimitWarning)
						}
					}
				}
			}
		}
	}()

	return ch
}

func (l *ReceiptListener) WaitForMetaTxn(ctx context.Context, metaTxnID MetaTxnID, optTimeout ...time.Duration) ([]*MetaTxnResult, *types.Receipt, error) {
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
	receipt := func() *ReceiptResult {
		l.muPastReceipts.Lock()
		defer l.muPastReceipts.Unlock()

		totalInspected := 0
		for _, bol := range l.pastReceipts {
			for _, receipt := range bol {
				totalInspected++
				if receipt.MetaTxnID == metaTxnID {
					l.log.Debug().
						Int("inspected", totalInspected).
						Str("meta-tx", string(metaTxnID)).
						Msgf("Found receipt among past receipts")

					return &receipt
				}
			}
		}

		l.log.Debug().
			Int("inspected", totalInspected).
			Str("meta-tx", string(metaTxnID)).
			Msgf("Receipt not found among past receipts. Now listening..")

		return nil
	}()
	if receipt != nil {
		return receipt.Results, receipt.TxnReceipt, nil
	}

	// Wait for receipt or context deadline
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

		case r := <-sub.ch:
			if r.MetaTxnID == metaTxnID {
				receipt = &r
				done = true
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}
	if receipt != nil {
		return receipt.Results, receipt.TxnReceipt, nil
	}
	return nil, nil, nil
}
