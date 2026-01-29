package malleable

import (
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
