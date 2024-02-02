package compressor

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// The compressor contract has no standard ABI
// instead, it has some custom 1 byte functions
// - 0x00: Execute 1 execute
// - 0x01: Execute N executes
// - 0x02: Read address (32 bytes for the index)
// - 0x03: Read bytes32 (32 bytes for the index)
// - 0x04: Read sizes
// - 0x05: Read storage slots (32 bytes for each index)
// - 0x06: Decompress 1 execute
// - 0x07: Decompress N executes

func GetTotals(ctx context.Context, provider *ethrpc.Provider, contract common.Address) (uint, uint, error) {
	res, err := provider.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: common.Hex2Bytes("04"),
	}, nil)

	if err != nil {
		return 0, 0, err
	}

	// First 16 bytes are the total number of addresses
	// Next 16 bytes are the total number of bytes32

	// Read only an uint64, since there will be no more than 2^64 addresses
	asize := uint(binary.BigEndian.Uint64(res[8:16])) + 1
	bsize := uint(binary.BigEndian.Uint64(res[24:32])) + 1

	return asize, bsize, nil
}

func AddressIndex(i uint) []byte {
	padded32 := make([]byte, 32)
	binary.BigEndian.PutUint64(padded32[24:32], uint64(i+1))
	return padded32
}

func Bytes32Index(i uint) []byte {
	padded32 := make([]byte, 32)
	binary.BigEndian.PutUint64(padded32[8:16], uint64(i))
	return padded32
}

func GenBatch(from uint, to uint, max uint, itemplate func(uint) []byte) []byte {
	var end uint

	if to < max {
		end = to
	} else {
		end = max
	}

	indexes := make([]byte, end*32)

	for j := uint(0); j < end; j++ {
		copy(indexes[j*32:j*32+32], itemplate(from+j))
	}

	return indexes
}

func ParseBatchResult(to map[string]uint, res []byte, offset uint) error {
	if len(res)%32 != 0 {
		return fmt.Errorf("invalid result length")
	}

	for j := uint(0); j < uint(len(res)/32); j++ {
		// Ignore results that all 0s
		r := res[j*32 : j*32+32]
		allZero := true
		for i := 31; i >= 0; i-- {
			if r[i] != 0 {
				allZero = false
				break
			}
		}
		if !allZero {
			to[string(res[j*32:j*32+32])] = j + offset
		}
	}

	return nil
}

func LoadAddresses(ctx context.Context, provider *ethrpc.Provider, contract common.Address, skip uint) (map[string]uint, error) {
	// Load total number of addresses
	asize, _, err := GetTotals(ctx, provider, contract)
	if err != nil {
		return nil, err
	}

	return LoadStorage(ctx, provider, contract, skip, asize, AddressIndex)
}

func LoadBytes32(ctx context.Context, provider *ethrpc.Provider, contract common.Address, skip uint) (map[string]uint, error) {
	// Load total number of bytes32
	_, bsize, err := GetTotals(ctx, provider, contract)
	if err != nil {
		return nil, err
	}

	return LoadStorage(ctx, provider, contract, skip, bsize, Bytes32Index)
}

func LoadStorage(ctx context.Context, provider *ethrpc.Provider, contract common.Address, skip uint, total uint, itemplate func(uint) []byte) (map[string]uint, error) {
	// Load addresses using chunks of 64 addresses
	out := make(map[string]uint)

	for i := skip; i < total; i += 64 {
		batch := GenBatch(i, total, 64, itemplate)

		res, err := provider.CallContract(ctx, ethereum.CallMsg{
			To:   &contract,
			Data: append([]byte{0x05}, batch...),
		}, nil)

		if err != nil {
			return nil, err
		}

		err = ParseBatchResult(out, res, i)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
