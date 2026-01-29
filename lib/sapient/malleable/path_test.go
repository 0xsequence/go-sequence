package malleable

import (
	"math/big"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

func TestPath_ResolveNestedPackedCalls(t *testing.T) {
	trailsABIJSON := `[{"name":"hydrateExecute","type":"function","inputs":[{"name":"payload","type":"bytes"},{"name":"hydrateData","type":"bytes"}]}]`
	erc2612ABIJSON := `[{"name":"permit","type":"function","inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"},{"name":"value","type":"uint256"},{"name":"deadline","type":"uint256"},{"name":"v","type":"uint8"},{"name":"r","type":"bytes32"},{"name":"s","type":"bytes32"}]}]`
	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}]}]`

	trailsABI, err := abi.JSON(strings.NewReader(trailsABIJSON))
	require.NoError(t, err)
	erc2612ABI, err := abi.JSON(strings.NewReader(erc2612ABIJSON))
	require.NoError(t, err)
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	permitCalldata, err := erc2612ABI.Pack(
		"permit",
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		big.NewInt(7),
		big.NewInt(8),
		uint8(27),
		common.HexToHash("0x01"),
		common.HexToHash("0x02"),
	)
	require.NoError(t, err)

	transferCalldata, err := erc20ABI.Pack(
		"transferFrom",
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		big.NewInt(7),
	)
	require.NoError(t, err)

	encodeAddr := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	innerPayload := v3.NewCallsPayload(encodeAddr, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"), Data: permitCalldata},
		{To: common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc"), Data: transferCalldata},
	}, big.NewInt(0), big.NewInt(0))

	packed := innerPayload.Encode(encodeAddr)
	hydrateData := []byte{0x55}
	outerCallData, err := trailsABI.Pack("hydrateExecute", packed, hydrateData)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0xdddddddddddddddddddddddddddddddddddddddd"), Data: outerCallData},
	}, big.NewInt(0), big.NewInt(0))

	permitValueSel := NewPath().
		CallData(0).
		ABI(&trailsABI, "hydrateExecute").
		ArgBytesData("payload").
		EncodedCallsPayload().
		EncodedCallData(0).
		ABI(&erc2612ABI, "permit").
		ArgSlot("value").
		AsSelector()
	t.Logf("permitValueSel: %s", permitValueSel.String())

	transferValueSel := NewPath().
		CallData(0).
		ABI(&trailsABI, "hydrateExecute").
		ArgBytesData("payload").
		EncodedCallsPayload().
		EncodedCallData(1).
		ABI(&erc20ABI, "transferFrom").
		ArgSlot("value").
		AsSelector()
	t.Logf("transferValueSel: %s", transferValueSel.String())

	permitRanges, err := permitValueSel.Resolve(&payload)
	require.NoError(t, err)
	require.Len(t, permitRanges, 1)

	transferRanges, err := transferValueSel.Resolve(&payload)
	require.NoError(t, err)
	require.Len(t, transferRanges, 1)

	permitSlice, err := permitRanges[0].Slice(&payload)
	require.NoError(t, err)
	transferSlice, err := transferRanges[0].Slice(&payload)
	require.NoError(t, err)

	require.Equal(t, permitCalldata[4+32*2:4+32*3], permitSlice)
	require.Equal(t, transferCalldata[4+32*2:4+32*3], transferSlice)
}

func TestPath_ResolveDirectTransferFromValue(t *testing.T) {
	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	transferCalldata, err := erc20ABI.Pack(
		"transferFrom",
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		big.NewInt(7),
	)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0x5555555555555555555555555555555555555555"), Data: transferCalldata},
	}, big.NewInt(0), big.NewInt(0))

	valueSel := NewPath().
		CallData(0).
		ABI(&erc20ABI, "transferFrom").
		ArgSlot("_value").
		AsSelector()
	t.Logf("valueSel: %s", valueSel.String())

	ranges, err := valueSel.Resolve(&payload)
	require.NoError(t, err)
	require.Len(t, ranges, 1)

	valueSlice, err := ranges[0].Slice(&payload)
	require.NoError(t, err)
	require.Equal(t, transferCalldata[4+32*2:4+32*3], valueSlice)
}

func TestPath_ResolveDirectTransferFromValueByIndex(t *testing.T) {
	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"","type":"address"},{"name":"","type":"address"},{"name":"","type":"uint256"}]}]`
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	transferCalldata, err := erc20ABI.Pack(
		"transferFrom",
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		big.NewInt(7),
	)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0x5555555555555555555555555555555555555555"), Data: transferCalldata},
	}, big.NewInt(0), big.NewInt(0))

	valueSel := NewPath().
		CallData(0).
		ABI(&erc20ABI, "transferFrom").
		ArgSlotIndex(2).
		AsSelector()
	t.Logf("valueSelByIndex: %s", valueSel.String())

	ranges, err := valueSel.Resolve(&payload)
	require.NoError(t, err)
	require.Len(t, ranges, 1)

	valueSlice, err := ranges[0].Slice(&payload)
	require.NoError(t, err)
	require.Equal(t, transferCalldata[4+32*2:4+32*3], valueSlice)
}

func TestPath_ResolveFailsWithSelectorMismatch(t *testing.T) {
	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	approveABIJSON := `[{"name":"approve","type":"function","inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}]}]`
	approveABI, err := abi.JSON(strings.NewReader(approveABIJSON))
	require.NoError(t, err)

	approveCalldata, err := approveABI.Pack(
		"approve",
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		big.NewInt(7),
	)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0x5555555555555555555555555555555555555555"), Data: approveCalldata},
	}, big.NewInt(0), big.NewInt(0))

	valueSel := NewPath().
		CallData(0).
		ABI(&erc20ABI, "transferFrom").
		ArgSlot("_value").
		AsSelector()
	t.Logf("mismatchedSelector: %s", valueSel.String())

	_, err = valueSel.Resolve(&payload)
	if err != nil {
		t.Logf("resolve error: %v", err)
	}
	require.Error(t, err)
}

func TestPath_ResolveFailsWithOutOfRangeCallIndex(t *testing.T) {
	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	transferCalldata, err := erc20ABI.Pack(
		"transferFrom",
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		big.NewInt(7),
	)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0x5555555555555555555555555555555555555555"), Data: transferCalldata},
	}, big.NewInt(0), big.NewInt(0))

	valueSel := NewPath().
		CallData(1).
		ABI(&erc20ABI, "transferFrom").
		ArgSlot("_value").
		AsSelector()
	t.Logf("outOfRangeCallIndex: %s", valueSel.String())

	_, err = valueSel.Resolve(&payload)
	if err != nil {
		t.Logf("resolve error: %v", err)
	}
	require.Error(t, err)
}

func TestPath_ResolveFailsWithInvalidEncodedCallsPayload(t *testing.T) {
	trailsABIJSON := `[{"name":"hydrateExecute","type":"function","inputs":[{"name":"payload","type":"bytes"},{"name":"hydrateData","type":"bytes"}]}]`
	trailsABI, err := abi.JSON(strings.NewReader(trailsABIJSON))
	require.NoError(t, err)

	erc20ABIJSON := `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`
	erc20ABI, err := abi.JSON(strings.NewReader(erc20ABIJSON))
	require.NoError(t, err)

	invalidPacked := []byte{}
	hydrateData := []byte{0x55}
	outerCallData, err := trailsABI.Pack("hydrateExecute", invalidPacked, hydrateData)
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{To: common.HexToAddress("0xdddddddddddddddddddddddddddddddddddddddd"), Data: outerCallData},
	}, big.NewInt(0), big.NewInt(0))

	valueSel := NewPath().
		CallData(0).
		ABI(&trailsABI, "hydrateExecute").
		ArgBytesData("payload").
		EncodedCallsPayload().
		EncodedCallData(0).
		ABI(&erc20ABI, "transferFrom").
		ArgSlot("_value").
		AsSelector()
	t.Logf("invalidEncodedCallsPayload: %s", valueSel.String())

	_, err = valueSel.Resolve(&payload)
	if err != nil {
		t.Logf("resolve error: %v", err)
	}
	require.Error(t, err)
}
