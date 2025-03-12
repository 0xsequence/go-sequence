package compressor

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

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

func GetTotals(ctx context.Context, provider *ethrpc.Provider, contract common.Address, skipBlocks uint) (uint, uint, error) {
	// Get the last block
	block, err := provider.BlockNumber(ctx)
	if err != nil {
		return 0, 0, err
	}

	block -= uint64(skipBlocks)

	res, err := provider.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: []byte{byte(METHOD_READ_SIZES)},
	}, big.NewInt(int64(block)))

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

func DecompressTransaction(ctx context.Context, provider *ethrpc.Provider, contract common.Address, compressed []byte) (common.Address, []byte, error) {
	// Replace the first byte with `METHOD_DECOMPRESS_1`
	if len(compressed) == 0 {
		return common.Address{}, nil, fmt.Errorf("empty compressed data")
	}

	c2 := make([]byte, len(compressed))
	copy(c2, compressed)
	c2[0] = byte(METHOD_DECOMPRESS_1)
	res, err := provider.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: c2,
	}, nil)

	if err != nil {
		return common.Address{}, nil, err
	}

	if len(res) < 32 {
		return common.Address{}, nil, fmt.Errorf("decompressed data too short")
	}

	// The last 32 bytes are the address
	addr := common.BytesToAddress(res[len(res)-32:])

	// The rest is the transaction, encoded as a function call to execute
	return addr, res[:len(res)-32], nil
}
