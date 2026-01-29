package malleable

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts/gen/trailsutils"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

const (
	trailsUtils = "0x0000000066c426Fe13962e276f894F11Aa6ebbF2"
	baseRPC     = "https://nodes.sequence.app/base"
	baseChainID = 8453
)

func TestComputeImageHash_DifferentSignatures(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{Data: []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
	}, big.NewInt(0), big.NewInt(0))

	hash1, err := ComputeImageHash(&payload, []byte{}, big.NewInt(1))
	require.NoError(t, err)

	signature := []byte{0x00, 0x00, 0x00, 0x00, 0x03}
	hash2, err := ComputeImageHash(&payload, signature, big.NewInt(1))
	require.NoError(t, err)

	require.NotEqual(t, hash1, hash2)
}

func TestValidateSignature_RepeatMismatch(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{Data: []byte{0x01, 0x02, 0x03}},
		{Data: []byte{0x01, 0xff, 0x03}},
	}, big.NewInt(0), big.NewInt(0))

	repeat := RepeatSection{TIndex: 0, CIndex: 0, Size: 2, TIndex2: 1, CIndex2: 0}
	sig, err := EncodeSignature(nil, []RepeatSection{repeat})
	require.NoError(t, err)

	err = ValidateSignature(&payload, sig)
	require.Error(t, err)
}

func TestComputeImageHash_RecoverSapientSignature(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(baseChainID), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           big.NewInt(0),
			Data:            []byte{0x01, 0x02},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	signature := []byte{}
	expected, err := ComputeImageHash(&payload, signature, big.NewInt(baseChainID))
	require.NoError(t, err)

	provider, err := ethrpc.NewProvider(baseRPC)
	require.NoError(t, err)

	contract, err := trailsutils.NewTrailsUtilsCaller(common.HexToAddress(trailsUtils), provider)
	require.NoError(t, err)

	callOpts := &bind.CallOpts{Context: context.Background()}
	onChainHash, err := contract.RecoverSapientSignature(callOpts, toTrailsUtilsPayload(payload), signature)
	require.NoError(t, err)

	require.Equal(t, expected, common.BytesToHash(onChainHash[:]))
}

func toTrailsUtilsPayload(payload v3.CallsPayload) trailsutils.PayloadDecoded {
	calls := make([]trailsutils.PayloadCall, len(payload.Calls))
	for i, call := range payload.Calls {
		value := call.Value
		if value == nil {
			value = big.NewInt(0)
		}
		gasLimit := call.GasLimit
		if gasLimit == nil {
			gasLimit = big.NewInt(0)
		}

		calls[i] = trailsutils.PayloadCall{
			To:              call.To,
			Value:           value,
			Data:            call.Data,
			GasLimit:        gasLimit,
			DelegateCall:    call.DelegateCall,
			OnlyFallback:    call.OnlyFallback,
			BehaviorOnError: big.NewInt(int64(call.BehaviorOnError)),
		}
	}

	noChainID := payload.ChainID().Sign() == 0

	return trailsutils.PayloadDecoded{
		Kind:          v3.KindTransactions,
		NoChainId:     noChainID,
		Calls:         calls,
		Space:         payload.Space,
		Nonce:         payload.Nonce,
		Message:       []byte{},
		ImageHash:     [32]byte{},
		Digest:        [32]byte{},
		ParentWallets: []common.Address{},
	}
}
