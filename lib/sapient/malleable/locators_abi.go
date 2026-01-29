package malleable

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

func CalldataStaticWord(method abi.Method, argIndex int) (start, length int, err error) {
	if argIndex < 0 || argIndex >= len(method.Inputs) {
		return 0, 0, fmt.Errorf("argIndex out of range")
	}
	return 4 + 32*argIndex, 32, nil
}

// CalldataBytesContent returns the raw bytes/string content (excludes length word and padding).
func CalldataBytesContent(calldata []byte, method abi.Method, argIndex int) (start, length int, err error) {
	tailStart, dataLen, err := calldataBytesTail(calldata, method, argIndex)
	if err != nil {
		return 0, 0, err
	}
	contentStart := tailStart + 32
	if contentStart+dataLen > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for bytes content")
	}
	return contentStart, dataLen, nil
}

// CalldataBytesEncoded returns the full ABI-encoded tail: length word + data + padding.
func CalldataBytesEncoded(calldata []byte, method abi.Method, argIndex int) (start, length int, err error) {
	tailStart, dataLen, err := calldataBytesTail(calldata, method, argIndex)
	if err != nil {
		return 0, 0, err
	}
	padded := ((dataLen + 31) / 32) * 32
	total := 32 + padded
	if tailStart+total > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for bytes encoded")
	}
	return tailStart, total, nil
}

func calldataBytesTail(calldata []byte, method abi.Method, argIndex int) (tailStart int, dataLen int, err error) {
	if argIndex < 0 || argIndex >= len(method.Inputs) {
		return 0, 0, fmt.Errorf("argIndex out of range")
	}
	t := method.Inputs[argIndex].Type
	if t.T != abi.BytesTy && t.T != abi.StringTy {
		return 0, 0, fmt.Errorf("arg %d is not bytes/string (got %s)", argIndex, t.String())
	}
	head := 4 + 32*argIndex
	if head+32 > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for head word")
	}

	off := new(big.Int).SetBytes(calldata[head : head+32])
	if !off.IsInt64() {
		return 0, 0, fmt.Errorf("dynamic offset too large")
	}
	tailStart = 4 + int(off.Int64())
	if tailStart+32 > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for tail length word")
	}

	l := new(big.Int).SetBytes(calldata[tailStart : tailStart+32])
	if !l.IsInt64() {
		return 0, 0, fmt.Errorf("bytes length too large")
	}
	dataLen = int(l.Int64())
	return tailStart, dataLen, nil
}
