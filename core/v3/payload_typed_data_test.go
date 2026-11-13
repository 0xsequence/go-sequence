package v3_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

type typedDataPayload interface {
	core.Payload
	TypedData() *ethcoder.TypedData
}

// TestPayloadTypedDataRoundTrip asserts that the JSON encoding of TypedData()
// reparses (as the Guard service does) to the digest Digest() computes.
func TestPayloadTypedDataRoundTrip(t *testing.T) {
	address := common.HexToAddress("0x468E8e29F6cfb0F6b7ff10ec6A1AB516ec849c04")
	calls := []v3.Call{
		{
			To:              common.HexToAddress("0x2222222222222222222222222222222222222222"),
			Value:           new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil),
			Data:            common.Hex2Bytes("deadbeef"),
			GasLimit:        big.NewInt(100000),
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
		{
			To:              common.HexToAddress("0x3333333333333333333333333333333333333333"),
			DelegateCall:    true,
			OnlyFallback:    true,
			BehaviorOnError: v3.BehaviorOnErrorAbort,
		},
	}
	imageHash := core.ImageHash{Hash: common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")}

	for _, chainID := range []*big.Int{big.NewInt(0), big.NewInt(42161)} {
		for _, wallets := range [][]common.Address{
			nil,
			{common.HexToAddress("0x4444444444444444444444444444444444444444"), common.HexToAddress("0x5555555555555555555555555555555555555555")},
		} {
			payloads := map[string]typedDataPayload{
				"calls":        v3.NewCallsPayload(address, chainID, calls, big.NewInt(1), big.NewInt(7), wallets),
				"message":      v3.NewMessagePayload(address, chainID, []byte("hello world"), wallets),
				"configUpdate": v3.NewConfigUpdatePayload(address, chainID, imageHash.Hash, wallets),
			}
			for kind, p := range payloads {
				t.Run(fmt.Sprintf("%s/chainID=%v/wallets=%d", kind, chainID, len(wallets)), func(t *testing.T) {
					data, err := json.Marshal(p.TypedData())
					require.NoError(t, err)

					parsed, err := ethcoder.TypedDataFromJSON(string(data))
					require.NoError(t, err)

					digest, err := parsed.EncodeDigest()
					require.NoError(t, err)
					require.Equal(t, p.Digest().Hash, common.Hash(digest))

					td := p.TypedData()
					td.Domain.ChainID.SetInt64(99)
					for _, args := range td.Types {
						args[0].Name = "scrambled"
					}
					scramble(td.Message)
					require.Equal(t, common.Hash(digest), p.Digest().Hash, "TypedData result must not alias the payload")
				})
			}
		}
	}
}

func scramble(v any) {
	switch v := v.(type) {
	case *big.Int:
		v.Add(v, big.NewInt(1))
	case []byte:
		if len(v) > 0 {
			v[0] ^= 0xff
		}
	case map[string]any:
		for _, x := range v {
			scramble(x)
		}
	case []any:
		for _, x := range v {
			scramble(x)
		}
	}
}
