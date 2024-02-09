package compressor

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
)

type Compressor struct {
	ctx context.Context

	Encoder   *Encoder
	CostModel *CostModel
	Contract  common.Address

	UseStorage bool

	UpdateInterval uint
	TrailBlocks    uint
	BatchSize      uint

	AddressesHeight uint
	Bytes32Height   uint

	LastIndexUpdate uint
	Provider        *ethrpc.Provider

	onLoadedIndexes func(uint, uint)

	mu *sync.RWMutex
}

func NewCompressor(
	ctx context.Context,
	provider *ethrpc.Provider,
	contract common.Address,
	costModel *CostModel,
	userStorage bool,
	updateInterval uint,
	trailBlocks uint,
	batchSize uint,
) (*Compressor, error) {
	if userStorage {
		if updateInterval == 0 {
			return nil, fmt.Errorf("update interval must be greater than 0")
		}

		if batchSize == 0 {
			return nil, fmt.Errorf("batch size must be greater than 0")
		}
	}

	// Check if the contract exists
	code, err := provider.CodeAt(ctx, contract, nil)
	if err != nil {
		return nil, err
	}

	if len(code) == 0 {
		return nil, fmt.Errorf("contract %s does not exist", contract.Hex())
	}

	c := &Compressor{
		ctx:             ctx,
		Encoder:         NewEncoder(),
		CostModel:       costModel,
		Contract:        contract,
		LastIndexUpdate: 0,
		Provider:        provider,
		UseStorage:      userStorage,
		UpdateInterval:  updateInterval,
		TrailBlocks:     trailBlocks,
		BatchSize:       batchSize,
		mu:              &sync.RWMutex{},
	}

	c.StartIndexUpdater()

	return c, nil
}

func (cm *Compressor) SetOnLoadedIndexes(f func(uint, uint)) {
	cm.onLoadedIndexes = f
}

func LoadIndexes(ci *Compressor) error {
	ci.mu.RLock()
	lenAddrs := ci.AddressesHeight
	lenBytes32s := ci.Bytes32Height
	ci.mu.RUnlock()

	// Don't lock while we read the state, it can take a while

	nah, addrs, nbh, bytes32s, err := LoadState(
		ci.ctx,
		ci.Provider,
		ci.Contract,
		ci.BatchSize,
		lenAddrs,
		lenBytes32s,
		ci.TrailBlocks,
	)

	if err != nil {
		return err
	}

	// Lock it only for the time it takes to update the indexes
	ci.mu.Lock()
	defer ci.mu.Unlock()

	for k, v := range addrs {
		ci.Encoder.AddressIndexes[k] = v
	}

	for k, v := range bytes32s {
		ci.Encoder.Bytes32Indexes[k] = v
	}

	ci.AddressesHeight = nah
	ci.Bytes32Height = nbh

	if ci.onLoadedIndexes != nil {
		ci.onLoadedIndexes(nah-lenAddrs, nbh-lenBytes32s)
	}

	return nil
}

func (cm *Compressor) StartIndexUpdater() {
	// No need to start the updater if we don't use storage
	if !cm.UseStorage {
		return
	}

	go func() {
		ticker := time.NewTicker(time.Duration(cm.UpdateInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := LoadIndexes(cm)
				if err != nil {
					fmt.Printf("Error updating indexes for chain %d: %v\n", cm.Contract, err)
				}
			case <-cm.ctx.Done():
				// Context cancelled, stop the ticker
				return
			}
		}
	}()
}

func (cm *Compressor) performCompress(ci *Compressor, wallet []byte, transaction *sequence.Transaction) ([]byte, EncodeType, error) {
	ci.mu.RLock()
	defer ci.mu.RUnlock()

	cbuffer := NewCBuffer(ci.UseStorage)
	et, err := ci.Encoder.WriteExecute(cbuffer, wallet, transaction)
	if err != nil {
		return nil, et, err
	}

	return cbuffer.Data(), et, nil
}

func (cm *Compressor) IsSaneCompression(
	input []byte,
	entrypoint common.Address,
	transaction *sequence.Transaction,
	decompressedEntrypoint common.Address,
	decompressed []byte,
) error {
	// The decompressed entrypoint should match the input entrypoint
	if !bytes.Equal(entrypoint.Bytes(), decompressedEntrypoint.Bytes()) {
		return fmt.Errorf("decompressed entrypoin t does not match input")
	}

	ed1, err := transaction.Execdata()
	if err != nil {
		return err
	}

	// We need to normalize the signature before comparing the exec data
	err = NormalizeTransactionSignature(transaction)
	if err != nil {
		return err
	}

	// Now we can re-compute the exec data and compare it with the decompressed data
	ed2, err := transaction.Execdata()
	if err != nil {
		return err
	}

	// One of the two ways of representing a Sequence signature in v2 is the dynamic format
	// should match the input data
	if !bytes.Equal(input, ed1) && !bytes.Equal(input, ed2) {
		return fmt.Errorf("exec data does not match input")
	}

	if !bytes.Equal(decompressed, ed2) && !bytes.Equal(decompressed, ed1) {
		return fmt.Errorf("decompressed exec data does not match input")
	}

	return nil
}

// We return two errors, the first one is the inner compression error, we returns it for traceability
// but the real error is the second one, which should only appear on a non-recoverable situation
func (cm *Compressor) TryCompress(
	input []byte,
	entrypoint common.Address,
	transaction *sequence.Transaction,
) (common.Address, []byte, EncodeType, error) {
	ci := cm
	compressed, et, err := cm.performCompress(ci, entrypoint.Bytes(), transaction)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	// We do a dry run of decompression to make sure it's sane
	decompressedEntrypoint, decompressed, err := DecompressTransaction(ci.ctx, ci.Provider, ci.Contract, compressed)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	// We need to validate that the compresseded data decompresses to what we expect
	err = cm.IsSaneCompression(input, entrypoint, transaction, decompressedEntrypoint, decompressed)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	return ci.Contract, compressed, et, nil
}
