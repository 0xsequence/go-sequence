package malleable

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
	require.Equal(t, 32, length)
	require.Equal(t, calldata[4+32*2:4+32*3], calldata[start:start+length])
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
	// Arg 0 = pair (uint256[2]): 2 words at offset 4, length 64
	start0, length0, err := CalldataStaticWord(method, 0)
	require.NoError(t, err)
	require.Equal(t, expectedPair, calldata[start0:start0+length0], "arg 0 should be encoded pair (111, 222)")
	// Arg 1 = single (uint256): 1 word at offset 4+64=68
	start1, length1, err := CalldataStaticWord(method, 1)
	require.NoError(t, err)
	require.Equal(t, expectedSingle, calldata[start1:start1+length1], "arg 1 should be encoded single (333)")
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
