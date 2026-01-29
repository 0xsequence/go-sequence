package malleable

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

func TestParsePackedCalls_SingleCallWithData(t *testing.T) {
	encodeAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	callData := []byte{0x01, 0x02, 0x03, 0x04}
	payload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), []v3.Call{
		{
			To:       common.HexToAddress("0x2222222222222222222222222222222222222222"),
			Value:    big.NewInt(5),
			Data:     callData,
			GasLimit: big.NewInt(7),
		},
	}, big.NewInt(0), big.NewInt(0))

	packed := payload.Encode(encodeAddr)
	layout, err := ParsePackedCalls(packed)
	require.NoError(t, err)
	require.Equal(t, 1, layout.NumCalls)
	require.Equal(t, 1, len(layout.CallData))

	span := layout.CallData[0]
	require.True(t, span.Start >= 0)
	require.Equal(t, callData, packed[span.Start:span.Start+span.Len])
}

func TestParsePackedCalls_MultipleCalls_MixedData(t *testing.T) {
	encodeAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	payload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), []v3.Call{
		{
			To:   common.HexToAddress("0x4444444444444444444444444444444444444444"),
			Data: []byte{0xaa, 0xbb},
		},
		{
			To:   common.HexToAddress("0x5555555555555555555555555555555555555555"),
			Data: nil,
		},
	}, big.NewInt(0), big.NewInt(1))

	packed := payload.Encode(encodeAddr)
	layout, err := ParsePackedCalls(packed)
	require.NoError(t, err)
	require.Equal(t, 2, layout.NumCalls)

	span0 := layout.CallData[0]
	require.True(t, span0.Start >= 0)
	require.Equal(t, []byte{0xaa, 0xbb}, packed[span0.Start:span0.Start+span0.Len])

	span1 := layout.CallData[1]
	require.Equal(t, -1, span1.Start)
	require.Equal(t, 0, span1.Len)
}

func TestParsePackedCalls_TruncatedPacked(t *testing.T) {
	encodeAddr := common.HexToAddress("0x6666666666666666666666666666666666666666")
	payload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), []v3.Call{
		{
			To:   common.HexToAddress("0x7777777777777777777777777777777777777777"),
			Data: []byte{0xaa, 0xbb, 0xcc},
		},
	}, big.NewInt(0), big.NewInt(0))

	packed := payload.Encode(encodeAddr)
	truncated := packed[:len(packed)-1]
	_, err := ParsePackedCalls(truncated)
	if err != nil {
		t.Logf("parse error: %v", err)
	}
	require.Error(t, err)
}

func TestParsePackedCalls_LargeCallCount(t *testing.T) {
	encodeAddr := common.HexToAddress("0x8888888888888888888888888888888888888888")
	const callCount = 300
	calls := make([]v3.Call, callCount)
	for i := 0; i < callCount; i++ {
		calls[i] = v3.Call{
			To: common.HexToAddress("0x9999999999999999999999999999999999999999"),
		}
	}
	payload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), calls, big.NewInt(0), big.NewInt(0))

	packed := payload.Encode(encodeAddr)
	layout, err := ParsePackedCalls(packed)
	require.NoError(t, err)
	require.Equal(t, callCount, layout.NumCalls)
	require.Len(t, layout.CallData, callCount)
	for _, span := range layout.CallData {
		require.Equal(t, -1, span.Start)
		require.Equal(t, 0, span.Len)
	}
}

func TestParsePackedCalls_WithNonceAndSpace(t *testing.T) {
	encodeAddr := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	callData := []byte{0x01}
	payload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), []v3.Call{
		{
			To:   common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"),
			Data: callData,
		},
	}, big.NewInt(0x1234), big.NewInt(0x42))

	packed := payload.Encode(encodeAddr)
	layout, err := ParsePackedCalls(packed)
	require.NoError(t, err)
	require.Equal(t, 1, layout.NumCalls)
	span := layout.CallData[0]
	require.True(t, span.Start >= 0)
	require.Equal(t, callData, packed[span.Start:span.Start+span.Len])
}
