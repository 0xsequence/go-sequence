package abicalldata

import (
	"fmt"
	"math"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

// AbiHeadWords returns the number of 32-byte words this type occupies in the
// ABI calldata head. Dynamic types (bytes, string, slice, and tuples/arrays
// that contain them) occupy 1 word (offset); static types use their encoded size.
func AbiHeadWords(t abi.Type) int {
	switch t.T {
	case abi.BytesTy, abi.StringTy, abi.SliceTy:
		return 1
	case abi.ArrayTy:
		if t.Size == 0 || isDynamicType(t) {
			return 1
		}
		return t.Size * AbiHeadWords(*t.Elem)
	case abi.TupleTy:
		if isDynamicType(t) {
			return 1
		}
		n := 0
		for _, e := range t.TupleElems {
			n += AbiHeadWords(*e)
		}
		return n
	default:
		return 1
	}
}

// calldataArgHeadOffset returns the byte offset (from start of calldata, so
// selector is 0..3) of the first word of the argument at argIndex. It is
// 4 + sum of ABI-encoded sizes of all preceding arguments.
func calldataArgHeadOffset(method abi.Method, argIndex int) (int, error) {
	if argIndex < 0 || argIndex >= len(method.Inputs) {
		return 0, fmt.Errorf("argIndex out of range")
	}
	offset := 4
	for j := 0; j < argIndex; j++ {
		offset += AbiHeadWords(method.Inputs[j].Type) * 32
	}
	return offset, nil
}

func CalldataStaticWord(method abi.Method, argIndex int) (start, length int, err error) {
	start, err = calldataArgHeadOffset(method, argIndex)
	if err != nil {
		return 0, 0, err
	}
	words := AbiHeadWords(method.Inputs[argIndex].Type)
	return start, words * 32, nil
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
	head, err := calldataArgHeadOffset(method, argIndex)
	if err != nil {
		return 0, 0, err
	}
	if head+32 > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for head word")
	}

	off := new(big.Int).SetBytes(calldata[head : head+32])
	if !off.IsInt64() {
		return 0, 0, fmt.Errorf("dynamic offset too large")
	}
	off64 := off.Int64()
	if off64 < 0 || off64 > math.MaxInt-4 {
		return 0, 0, fmt.Errorf("dynamic offset out of range")
	}
	tailStart = 4 + int(off64)
	if tailStart+32 > len(calldata) {
		return 0, 0, fmt.Errorf("calldata too short for tail length word")
	}

	l := new(big.Int).SetBytes(calldata[tailStart : tailStart+32])
	if !l.IsInt64() {
		return 0, 0, fmt.Errorf("bytes length too large")
	}
	dataLen = int(l.Int64())
	if dataLen > math.MaxInt-tailStart-32 {
		return 0, 0, fmt.Errorf("bytes length too large")
	}
	return tailStart, dataLen, nil
}

func isDynamicType(t abi.Type) bool {
	switch t.T {
	case abi.BytesTy, abi.StringTy, abi.SliceTy:
		return true
	case abi.ArrayTy:
		return t.Size == 0 || isDynamicType(*t.Elem)
	case abi.TupleTy:
		for _, elem := range t.TupleElems {
			if isDynamicType(*elem) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
