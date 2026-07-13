package v3_test

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// The expected digests below are reference vectors produced by
// SessionSig.hashPayloadCallIdx in 0xsequence/wallet-contracts-v3 (via a forge
// test constructing the identical payloads), so this test asserts parity with
// the on-chain verification.
func TestHashPayloadCallIdx(t *testing.T) {
	payload := v3.NewCallsPayload(
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		big.NewInt(42161),
		[]v3.Call{
			{
				To:              common.HexToAddress("0x2222222222222222222222222222222222222222"),
				Value:           new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil), // 1 ether
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
		},
		big.NewInt(0),
		big.NewInt(7),
	)

	expectedPayloadHash := common.HexToHash("0x9c891ce70c80739f54b60eb7f8a0d0e80e7f90f2c9aea51a0c4f34ec1cee067e")
	if digest := payload.Digest().Hash; digest != expectedPayloadHash {
		t.Errorf("payload digest mismatch: got %v, expected %v", digest, expectedPayloadHash)
	}

	expectedCallHashes := []common.Hash{
		common.HexToHash("0x1957bab0f26824823a1f40bb132e9a6ab186db88c0adb7fb806351f6766ba66c"),
		common.HexToHash("0xcb5346174462b1c9a0279b5acbdfe2458ba0825cb10d35eb912fa491c1ee5cdd"),
	}
	for i, expected := range expectedCallHashes {
		hash, err := v3.HashPayloadCallIdx(payload, i)
		if err != nil {
			t.Fatalf("HashPayloadCallIdx(%v): %v", i, err)
		}
		if hash != expected {
			t.Errorf("call %v hash mismatch: got %v, expected %v", i, hash, expected)
		}
	}

	payload2 := v3.NewCallsPayload(
		common.HexToAddress("0x5555555555555555555555555555555555555555"),
		big.NewInt(1),
		[]v3.Call{
			{
				To:              common.HexToAddress("0x4444444444444444444444444444444444444444"),
				Data:            []byte{0x00},
				BehaviorOnError: v3.BehaviorOnErrorIgnore,
			},
		},
		big.NewInt(12345),
		big.NewInt(0),
	)

	expected2 := common.HexToHash("0xde2ebba8ab9a581d22dbb5d066086ae1ab3d884f9becc493448b2875e7f3c6d4")
	hash2, err := v3.HashPayloadCallIdx(payload2, 0)
	if err != nil {
		t.Fatalf("HashPayloadCallIdx: %v", err)
	}
	if hash2 != expected2 {
		t.Errorf("vector 2 hash mismatch: got %v, expected %v", hash2, expected2)
	}

	if _, err := v3.HashPayloadCallIdx(payload2, 1); err == nil {
		t.Errorf("expected out-of-range error for call index 1")
	}
	if _, err := v3.HashPayloadCallIdx(payload2, -1); err == nil {
		t.Errorf("expected out-of-range error for call index -1")
	}
}
