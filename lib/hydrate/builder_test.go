package hydrate

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/lib/abicalldata"
	"github.com/stretchr/testify/require"
)

// Golden bytes match the hand-built hydratePayload in trails-api
// lib/intentmachine/v1_5/intent_payloads.go (buildOriginPermitPayload).
func TestBuilder_MatchesIntentPayloadsGolden(t *testing.T) {
	owner := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	originToken := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")

	pl := v3.NewCallsPayload(
		common.Address{},
		big.NewInt(0),
		[]v3.Call{
			{To: originToken, Data: make([]byte, 128)},
			{To: originToken, Data: make([]byte, 128)},
			{To: originToken, Data: make([]byte, 128)},
			{To: originToken, Data: make([]byte, 128)},
		},
		big.NewInt(0),
		big.NewInt(0),
	)

	b := NewBuilder(&pl)
	require.NoError(t, b.ForCall(0).DataAddress(
		abicalldata.NewRangeSelector(0, 0x30, 20),
		SourceSelf(),
	))
	require.NoError(t, b.ForCall(1).DataAddress(
		abicalldata.NewRangeSelector(1, 0x30, 20),
		SourceSelf(),
	))
	require.NoError(t, b.ForCall(3).DataAddress(
		abicalldata.NewRangeSelector(3, 0x30, 20),
		SourceSelf(),
	))
	require.NoError(t, b.ForCall(3).DataERC20Allowance(
		abicalldata.NewRangeSelector(3, 0x44, 32),
		SourceAddress(owner),
		originToken,
		SourceSelf(),
	))

	got, err := b.Build()
	require.NoError(t, err)

	want := []byte{
		0x00, 0x10, 0x00, 0x30, 0x00,
		0x01, 0x10, 0x00, 0x30, 0x00,
		0x03, 0x10, 0x00, 0x30,
		0x43,
	}
	want = append(want, owner.Bytes()...)
	want = append(want, 0x00, 0x44)
	want = append(want, originToken.Bytes()...)
	want = append(want, 0x00, 0x00)
	require.Equal(t, want, got)
}

func TestPackHydrateExecute_RoundTripSelector(t *testing.T) {
	packed := []byte{0x01, 0x02, 0x03}
	hydratePayload := []byte{0x00, 0x10, 0x00, 0x04, 0x00}
	data, err := PackHydrateExecute(packed, hydratePayload)
	require.NoError(t, err)
	method, ok := ABI.Methods["hydrateExecute"]
	require.True(t, ok)
	require.GreaterOrEqual(t, len(data), 4)
	require.Equal(t, method.ID, data[:4])
}

// Simulate a transfer call with tx.origin and msg.sender balances hydrated.
// erc20.transfer(<tx.origin>, <erc20balance<msg.sender>>)
func TestBuilder_SingleCallMultipleReplacements_TransferTxOriginAndMsgSenderBalance(t *testing.T) {
	erc20ABI, err := abi.JSON(strings.NewReader(`[{"name":"transfer","type":"function","inputs":[{"name":"to","type":"address"},{"name":"amount","type":"uint256"}],"outputs":[{"name":"","type":"bool"}]}]`))
	require.NoError(t, err)

	token := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	placeholderTo := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	calldata, err := erc20ABI.Pack("transfer", placeholderTo, big.NewInt(0))
	require.NoError(t, err)

	method := erc20ABI.Methods["transfer"]
	toWordStart, _, err := abicalldata.CalldataStaticWord(method, 0)
	require.NoError(t, err)
	amountStart, amountLen, err := abicalldata.CalldataStaticWord(method, 1)
	require.NoError(t, err)
	require.Equal(t, 32, amountLen)

	pl := v3.NewCallsPayload(
		common.Address{},
		big.NewInt(0),
		[]v3.Call{{To: token, Data: calldata}},
		big.NewInt(0),
		big.NewInt(0),
	)

	b := NewBuilder(&pl)
	require.NoError(t, b.ForCall(0).DataAddress(
		abicalldata.NewRangeSelector(0, toWordStart+12, 20),
		SourceTxOrigin(),
	))
	require.NoError(t, b.ForCall(0).DataERC20Balance(
		abicalldata.NewRangeSelector(0, amountStart, 32),
		token,
		SourceMsgSender(),
	))

	got, err := b.Build()
	require.NoError(t, err)

	want := []byte{
		0x00,             // tindex
		0x12, 0x00, 0x10, // DATA_ADDRESS + TX_ORIGIN, cindex=0x0010 ("to")
		0x31, 0x00, 0x24, // DATA_ERC20_BALANCE + MSG_SENDER, cindex=0x0024 ("amount")
	}
	want = append(want, token.Bytes()...)
	want = append(want, 0x00) // signalNextHydrate

	require.Equal(t, want, got)
}

func TestBuilder_DataAddress_ArgSlotUsesRightAlignedAddressBytes(t *testing.T) {
	ownerABI, err := abi.JSON(strings.NewReader(`[{"name":"setOwner","type":"function","inputs":[{"name":"owner","type":"address"}],"outputs":[]}]`))
	require.NoError(t, err)

	placeholderOwner := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	calldata, err := ownerABI.Pack("setOwner", placeholderOwner)
	require.NoError(t, err)

	matches := make([]int, 0, 2)
	for i := 0; i+len(placeholderOwner.Bytes()) <= len(calldata); i++ {
		if bytes.Equal(calldata[i:i+len(placeholderOwner.Bytes())], placeholderOwner.Bytes()) {
			matches = append(matches, i)
		}
	}
	require.Len(t, matches, 1, "expected exactly one encoded address occurrence in calldata")
	addressDataOffset := matches[0]
	require.GreaterOrEqual(t, addressDataOffset, 12)
	require.Equal(t, make([]byte, 12), calldata[addressDataOffset-12:addressDataOffset], "address should be right-aligned in ABI word")

	pl := v3.NewCallsPayload(
		common.Address{},
		big.NewInt(0),
		[]v3.Call{{To: common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"), Data: calldata}},
		big.NewInt(0),
		big.NewInt(0),
	)

	ownerSelector := abicalldata.NewPath().
		CallData(0).
		ABI(&ownerABI, "setOwner").
		ArgSlot("owner").
		AsSelector()

	b := NewBuilder(&pl)
	require.NoError(t, b.ForCall(0).DataAddress(ownerSelector, SourceSelf()))

	got, err := b.Build()
	require.NoError(t, err)
	require.Len(t, got, 5)

	cindex := int(got[2])<<8 | int(got[3])
	require.Equal(t, addressDataOffset, cindex, "hydrate patch index must point to encoded address bytes, not slot start")
}
