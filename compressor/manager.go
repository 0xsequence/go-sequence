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

type CompressorInstance struct {
	Context context.Context

	Compressor *Compressor
	CostModel  *CostModel
	Contract   common.Address

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

type CompressorManager struct {
	instance *CompressorInstance
}

func NewCompressorManager(
	context context.Context,
	provider *ethrpc.Provider,
	contract common.Address,
	costModel *CostModel,
	userStorage bool,
	updateInterval uint,
	trailBlocks uint,
	batchSize uint,
) *CompressorManager {
	c := CompressorManager{
		instance: &CompressorInstance{
			Context:         context,
			Compressor:      NewCompressor(),
			CostModel:       costModel,
			Contract:        contract,
			LastIndexUpdate: 0,
			Provider:        provider,
			UseStorage:      userStorage,
			UpdateInterval:  updateInterval,
			TrailBlocks:     trailBlocks,
			BatchSize:       batchSize,
			mu:              &sync.RWMutex{},
		},
	}

	c.StartIndexUpdater()

	return &c
}

func (cm *CompressorManager) SetOnLoadedIndexes(f func(uint, uint)) {
	cm.instance.onLoadedIndexes = f
}

func LoadIndexes(ci *CompressorInstance) error {
	ci.mu.RLock()
	lenAddrs := ci.AddressesHeight
	lenBytes32s := ci.Bytes32Height
	ci.mu.RUnlock()

	// Don't lock while we read the state, it can take a while

	nah, addrs, nbh, bytes32s, err := LoadState(
		ci.Context,
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
		ci.Compressor.AddressIndexes[k] = v
	}

	for k, v := range bytes32s {
		ci.Compressor.Bytes32Indexes[k] = v
	}

	ci.AddressesHeight = nah
	ci.Bytes32Height = nbh

	if ci.onLoadedIndexes != nil {
		ci.onLoadedIndexes(nah-lenAddrs, nbh-lenBytes32s)
	}

	return nil
}

func (cm *CompressorManager) StartIndexUpdater() {
	go func() {
		ticker := time.NewTicker(time.Duration(cm.instance.UpdateInterval) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := LoadIndexes(cm.instance)
				if err != nil {
					fmt.Printf("Error updating indexes for chain %d: %v\n", cm.instance.Contract, err)
				}
			case <-cm.instance.Context.Done():
				// Context cancelled, stop the ticker
				return
			}
		}
	}()
}

func (cm *CompressorManager) performCompress(ci *CompressorInstance, wallet []byte, transaction *sequence.Transaction) ([]byte, EncodeType, error) {
	ci.mu.RLock()
	defer ci.mu.RUnlock()

	cbuffer := NewCBuffer(ci.UseStorage)
	et, err := ci.Compressor.WriteExecute(cbuffer, wallet, transaction)
	if err != nil {
		return nil, et, err
	}

	return cbuffer.Data(), et, nil
}

func (cm *CompressorManager) IsSaneCompression(
	input []byte,
	entrypoint common.Address,
	transaction *sequence.Transaction,
	decompressedEntrypoint common.Address,
	decompressed []byte,
) error {
	// The decompressed entrypoint should match the input entrypoint
	if !bytes.Equal(entrypoint.Bytes(), decompressedEntrypoint.Bytes()) {
		return fmt.Errorf("decompressed entrypoint does not match input")
	}

	ed1, err := transaction.Execdata()
	if err != nil {
		return err
	}

	// One thing that happens is that there are two ways of representing a Sequence signature in v2
	// lagacy and dynamic, decompressor will ALWAYS decompress to the dynamic format, so if
	// the input is in the legacy format, we need to convert it to the dynamic format.
	// This is easy, because the legacy format starts with 0x00, if that's the case we just
	// add 0x01 at the beginning of the signature
	if len(transaction.Signature) == 0 {
		return fmt.Errorf("empty signature")
	}

	if transaction.Signature[0] == 0 {
		transaction.Signature = append([]byte{1}, transaction.Signature...)
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

	if !bytes.Equal(decompressed, ed2) {
		return fmt.Errorf("exec data does not match input")
	}

	return nil
}

// We return two errors, the first one is the inner compression error, we returns it for traceability
// but the real error is the second one, which should only appear on a non-recoverable situation
func (cm *CompressorManager) TryCompress(
	input []byte,
	entrypoint common.Address,
	transaction *sequence.Transaction,
) (common.Address, []byte, EncodeType, error) {
	ci := cm.instance
	compressed, et, err := cm.performCompress(ci, entrypoint.Bytes(), transaction)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	// We do a dry run of decompression to make sure it's sane
	decompressedEntrypoint, decompressed, err := DecompressTransaction(ci.Context, ci.Provider, ci.Contract, compressed)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	// We need to validate that the compresseded data decompresses to what we expect
	err = cm.IsSaneCompression(input, entrypoint, transaction, decompressedEntrypoint, decompressed)
	if err != nil {
		return common.Address{}, nil, 0, err
	}

	// Now, for chains that don't use storage, we should check if the compressed data is really
	// smaller than the input, if it's not, we need to return an error
	if !ci.UseStorage {
		// TODO: implement this
	}

	return ci.Contract, compressed, et, nil
}
