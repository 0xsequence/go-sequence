package abicalldata

import (
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestCalldataBytesContent(t *testing.T) {
	abiDef := `[{"name":"hydrateExecute","type":"function","inputs":[{"name":"payload","type":"bytes"},{"name":"hydrateData","type":"bytes"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	payloadBytes := []byte{0x11, 0x22, 0x33, 0x44}
	hydrateData := []byte{0xaa}
	calldata, err := parsedABI.Pack("hydrateExecute", payloadBytes, hydrateData)
	require.NoError(t, err)

	method := parsedABI.Methods["hydrateExecute"]
	start, length, err := CalldataBytesContent(calldata, method, 0)
	require.NoError(t, err)
	require.Equal(t, payloadBytes, calldata[start:start+length])
}

func TestCalldataBytesEncoded(t *testing.T) {
	abiDef := `[{"name":"hydrateExecute","type":"function","inputs":[{"name":"payload","type":"bytes"},{"name":"hydrateData","type":"bytes"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	payloadBytes := []byte{0x11, 0x22, 0x33, 0x44}
	hydrateData := []byte{0xaa}
	calldata, err := parsedABI.Pack("hydrateExecute", payloadBytes, hydrateData)
	require.NoError(t, err)

	lengthWord := common.BigToHash(big.NewInt(int64(len(payloadBytes)))).Bytes()
	paddedLen := ((len(payloadBytes) + 31) / 32) * 32
	padding := make([]byte, paddedLen-len(payloadBytes))
	encodedPayloadBytes := append(append(lengthWord, payloadBytes...), padding...)

	method := parsedABI.Methods["hydrateExecute"]
	start, length, err := CalldataBytesEncoded(calldata, method, 0)
	require.NoError(t, err)
	require.Equal(t, encodedPayloadBytes, calldata[start:start+length])
}

func TestCalldataStaticWord(t *testing.T) {
	abiDef := `[{"name":"permit","type":"function","inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"deadline","type":"uint256"},{"name":"v","type":"uint8"},{"name":"r","type":"bytes32"},{"name":"s","type":"bytes32"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack(
		"permit",
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		big.NewInt(123),
		big.NewInt(456),
		uint8(27),
		common.HexToHash("0x01"),
		common.HexToHash("0x02"),
	)
	require.NoError(t, err)

	method := parsedABI.Methods["permit"]
	start, length, err := CalldataStaticWord(method, 2)
	require.NoError(t, err)
	expected := common.BigToHash(big.NewInt(123)).Bytes() // value is arg index 2
	require.Equal(t, expected, calldata[start:start+length])
}

// TestCalldataStaticWord_CompositeStaticArg verifies head offset is computed from
// cumulative ABI sizes, not 32*argIndex. Method has uint256[2] then uint256; arg 0
// is 2 words, arg 1 is 1 word. With 32*argIndex, arg 1 would be read at offset 36
// (wrong); correct offset is 4+64=68. We assert the slices at the computed offsets
// contain the intended ABI-encoded values (111, 222 and 333).
func TestCalldataStaticWord_CompositeStaticArg(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"pair","type":"uint256[2]"},{"name":"single","type":"uint256"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack("f", [2]*big.Int{big.NewInt(111), big.NewInt(222)}, big.NewInt(333))
	require.NoError(t, err)

	// Expected ABI-encoded values (32-byte big-endian each)
	expectedPair := append(
		common.BigToHash(big.NewInt(111)).Bytes(),
		common.BigToHash(big.NewInt(222)).Bytes()...,
	)
	expectedSingle := common.BigToHash(big.NewInt(333)).Bytes()

	method := parsedABI.Methods["f"]
	start0, length0, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expectedPair, calldata[start0:start0+length0], "arg 0 should be encoded pair (111, 222)")
	start1, length1, err := CalldataStaticWord(method, 1)
	require.NoError(t, err)
	require.Equal(t, expectedSingle, calldata[start1:start1+length1], "arg 1 should be encoded single (333)")
}

// TestCalldataStaticWord_DynamicTupleBeforeArg verifies a dynamic tuple (e.g. (uint256,bytes))
// is counted as 1 head word, not the sum of its elements. Otherwise the following arg's
// head offset would be wrong and ArgSlot/ArgBytesData would read incorrect bytes.
func TestCalldataStaticWord_DynamicTupleBeforeArg(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"pair","type":"tuple","components":[{"name":"n","type":"uint256"},{"name":"data","type":"bytes"}]},{"name":"value","type":"uint256"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	// Pack: tuple (big.NewInt(42), []byte("x")) and value 999
	calldata, err := parsedABI.Pack("f", struct {
		N    *big.Int
		Data []byte
	}{big.NewInt(42), []byte("x")}, big.NewInt(999))
	require.NoError(t, err)

	method := parsedABI.Methods["f"]
	// Expected arg 0 head: the offset word the encoder wrote (first head word after selector)
	expectedArg0 := calldata[4 : 4+32]
	start0, length0, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expectedArg0, calldata[start0:start0+length0], "arg 0 should be offset word")
	expectedArg1 := common.BigToHash(big.NewInt(999)).Bytes()
	start1, length1, err := CalldataStaticWord(method, 1)
	require.NoError(t, err)
	require.Equal(t, expectedArg1, calldata[start1:start1+length1], "arg 1 should be encoded value (999)")
}

// TestCalldataStaticWord_StaticTuple verifies a fully static tuple (uint256,uint256)
// is counted as 2 head words; the slice covers both words.
func TestCalldataStaticWord_StaticTuple(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"pair","type":"tuple","components":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}]}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack("f", struct {
		A *big.Int
		B *big.Int
	}{big.NewInt(11), big.NewInt(22)})
	require.NoError(t, err)

	method := parsedABI.Methods["f"]
	expected := append(
		common.BigToHash(big.NewInt(11)).Bytes(),
		common.BigToHash(big.NewInt(22)).Bytes()...,
	)
	start, length, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expected, calldata[start:start+length])
}

// TestCalldataStaticWord_NestedStaticTuple verifies a nested static tuple
// ((uint256,uint256), uint256) is 3 head words: inner 2 + outer 1.
func TestCalldataStaticWord_NestedStaticTuple(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"outer","type":"tuple","components":[{"name":"inner","type":"tuple","components":[{"name":"a","type":"uint256"},{"name":"b","type":"uint256"}]},{"name":"c","type":"uint256"}]}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack("f", struct {
		Inner struct {
			A *big.Int
			B *big.Int
		}
		C *big.Int
	}{
		Inner: struct {
			A *big.Int
			B *big.Int
		}{big.NewInt(1), big.NewInt(2)},
		C: big.NewInt(3),
	})
	require.NoError(t, err)

	method := parsedABI.Methods["f"]
	expected := append(
		common.BigToHash(big.NewInt(1)).Bytes(),
		common.BigToHash(big.NewInt(2)).Bytes()...,
	)
	expected = append(expected, common.BigToHash(big.NewInt(3)).Bytes()...)
	start, length, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expected, calldata[start:start+length])
}

// TestCalldataStaticWord_MixedStaticAndDynamicArgs verifies head offsets with
// uint256, then dynamic tuple, then uint256: 1 + 1 + 1 = 3 words.
func TestCalldataStaticWord_MixedStaticAndDynamicArgs(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"a","type":"uint256"},{"name":"t","type":"tuple","components":[{"name":"n","type":"uint256"},{"name":"data","type":"bytes"}]},{"name":"b","type":"uint256"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack("f",
		big.NewInt(100),
		struct {
			N    *big.Int
			Data []byte
		}{big.NewInt(42), []byte("x")},
		big.NewInt(200))
	require.NoError(t, err)

	method := parsedABI.Methods["f"]
	expectedArg0 := common.BigToHash(big.NewInt(100)).Bytes()
	start0, len0, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expectedArg0, calldata[start0:start0+len0])
	// Expected arg 1 head: the offset word the encoder wrote (second head word)
	expectedArg1 := calldata[4+32 : 4+64]
	start1, len1, err := CalldataStaticWord(method, 1)
	require.NoError(t, err)
	require.Equal(t, expectedArg1, calldata[start1:start1+len1])
	expectedArg2 := common.BigToHash(big.NewInt(200)).Bytes()
	start2, len2, err := CalldataStaticWord(method, 2)
	require.NoError(t, err)
	require.Equal(t, expectedArg2, calldata[start2:start2+len2])
}

// TestCalldataStaticWord_NestedDynamicTuple verifies an outer tuple that contains
// a dynamic element (uint256,(uint256,bytes)) is still 1 head word.
func TestCalldataStaticWord_NestedDynamicTuple(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"outer","type":"tuple","components":[{"name":"n","type":"uint256"},{"name":"inner","type":"tuple","components":[{"name":"a","type":"uint256"},{"name":"data","type":"bytes"}]}]}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)

	calldata, err := parsedABI.Pack("f", struct {
		N     *big.Int
		Inner struct {
			A    *big.Int
			Data []byte
		}
	}{big.NewInt(10), struct {
		A    *big.Int
		Data []byte
	}{big.NewInt(20), []byte("y")}})
	require.NoError(t, err)

	method := parsedABI.Methods["f"]
	// Expected head: the offset word the encoder wrote (first head word after selector)
	expected := calldata[4 : 4+32]
	start, length, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expected, calldata[start:start+length])
}

// TestCalldataBytesTail_RejectsHugeDynamicOffset ensures that a malformed
// calldata word with a dynamic offset near math.MaxInt64 does not overflow
// (4+offset can wrap tailStart negative) and panic on slice; we must return an error.
func TestCalldataBytesTail_RejectsHugeDynamicOffset(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"payload","type":"bytes"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)
	method := parsedABI.Methods["f"]

	// Build calldata: selector (4) + one 32-byte word. Word = huge offset so 4+offset overflows int.
	calldata := make([]byte, 4+32)
	copy(calldata[:4], parsedABI.Methods["f"].ID)
	offBytes := common.BigToHash(big.NewInt(math.MaxInt64 - 2)).Bytes()
	copy(calldata[4:4+32], offBytes)

	_, _, err = CalldataBytesContent(calldata, method, 0)
	require.Error(t, err)
	require.Contains(t, err.Error(), "dynamic offset")
}

// TestCalldataBytesTail_RejectsHugeBytesLength ensures a malformed length word
// (e.g. near math.MaxInt64) is rejected so contentStart+dataLen and padding math
// cannot overflow and cause panics downstream.
func TestCalldataBytesTail_RejectsHugeBytesLength(t *testing.T) {
	abiDef := `[{"name":"f","type":"function","inputs":[{"name":"payload","type":"bytes"}]}]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	require.NoError(t, err)
	method := parsedABI.Methods["f"]

	// Calldata: selector (4) + offset word 32 (points to 4+32=36) + at 36: length word.
	// Use length = math.MaxInt64 so (dataLen+31) and contentStart+dataLen would overflow.
	calldata := make([]byte, 4+32+32) // head + offset 32 + length word
	copy(calldata[:4], method.ID)
	copy(calldata[4:4+32], common.BigToHash(big.NewInt(32)).Bytes())
	lenBytes := common.BigToHash(big.NewInt(math.MaxInt64)).Bytes()
	copy(calldata[36:36+32], lenBytes)

	_, _, err = CalldataBytesContent(calldata, method, 0)
	require.Error(t, err)
	require.Contains(t, err.Error(), "length")
}
