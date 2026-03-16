package eip6492

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsEIP6492Signature(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		assert.False(t, IsEIP6492Signature(nil))
	})
	t.Run("empty", func(t *testing.T) {
		assert.False(t, IsEIP6492Signature([]byte{}))
	})
	t.Run("shorter than magic", func(t *testing.T) {
		assert.False(t, IsEIP6492Signature([]byte{0x64, 0x92}))
	})
	t.Run("plain signature no suffix", func(t *testing.T) {
		sig := make([]byte, 65)
		sig[64] = 27
		assert.False(t, IsEIP6492Signature(sig))
	})
	t.Run("ends with magic bytes", func(t *testing.T) {
		payload := []byte("some payload")
		sig := append(append([]byte{}, payload...), EIP6492MagicBytes...)
		assert.True(t, IsEIP6492Signature(sig))
	})
	t.Run("exactly magic bytes", func(t *testing.T) {
		assert.True(t, IsEIP6492Signature(EIP6492MagicBytes))
	})
	t.Run("suffix but wrong content", func(t *testing.T) {
		// 31 bytes + one wrong byte at end (magic is 32 bytes)
		bad := make([]byte, 32)
		copy(bad, EIP6492MagicBytes[:31])
		bad[31] = 0xff
		assert.False(t, IsEIP6492Signature(bad))
	})
}

func TestDecodeEIP6492Signature(t *testing.T) {
	t.Run("not eip6492", func(t *testing.T) {
		_, _, _, err := DecodeEIP6492Signature([]byte("not a 6492 sig"))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not an eip6492 signature")
	})
	t.Run("only magic bytes", func(t *testing.T) {
		_, _, _, err := DecodeEIP6492Signature(EIP6492MagicBytes)
		require.Error(t, err)
		// ABI decode will fail on empty/short payload
		assert.Error(t, err)
	})
	t.Run("valid payload roundtrip", func(t *testing.T) {
		factory := common.HexToAddress("0x00000000000000000000000000000000deadbeef")
		factoryCalldata := []byte{0xde, 0xad, 0xbe, 0xef}
		innerSig := make([]byte, 65)
		innerSig[64] = 28

		encoded, err := ethcoder.AbiCoder(
			[]string{"address", "bytes", "bytes"},
			[]interface{}{factory, factoryCalldata, innerSig},
		)
		require.NoError(t, err)
		signature := append(encoded, EIP6492MagicBytes...)

		decodedFactory, decodedCalldata, decodedSig, err := DecodeEIP6492Signature(signature)
		require.NoError(t, err)
		assert.Equal(t, factory, decodedFactory)
		assert.Equal(t, factoryCalldata, decodedCalldata)
		assert.Equal(t, innerSig, decodedSig)
	})
	t.Run("decode then is eip6492", func(t *testing.T) {
		factory := common.HexToAddress("0x1111111111111111111111111111111111111111")
		calldata := []byte("deploy()")
		innerSig := bytesFromHex("0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef00")
		if len(innerSig) != 65 {
			innerSig = make([]byte, 65)
			innerSig[64] = 27
		}

		encoded, err := ethcoder.AbiCoder(
			[]string{"address", "bytes", "bytes"},
			[]interface{}{factory, calldata, innerSig},
		)
		require.NoError(t, err)
		signature := append(encoded, EIP6492MagicBytes...)

		assert.True(t, IsEIP6492Signature(signature))
		gotFactory, gotCalldata, gotSig, err := DecodeEIP6492Signature(signature)
		require.NoError(t, err)
		assert.Equal(t, factory, gotFactory)
		assert.Equal(t, calldata, gotCalldata)
		assert.Equal(t, innerSig, gotSig)
	})
}

func bytesFromHex(s string) []byte {
	b, _ := hexutil.Decode(s)
	return b
}

func TestValidateEIP6492Offchain_RequiresRPC(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping ValidateEIP6492Offchain in short mode")
	}
	// Connect to local testchain; skip if not running
	provider, err := ethrpc.NewProvider("http://localhost:8545")
	if err != nil {
		t.Skipf("provider not available: %v", err)
	}
	ctx := context.Background()
	_, err = provider.ChainID(ctx)
	if err != nil {
		t.Skipf("testchain not running: %v", err)
	}

	// Use a zero address and invalid payload: we only check that the offchain path
	// runs and returns false (invalid sig). The chain may return (false, nil) or
	// (false, err) when the validator reverts.
	signer := common.HexToAddress("0x0000000000000000000000000000000000000001")
	hash := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
	signature := append([]byte{0x00}, EIP6492MagicBytes...)

	valid, err := ValidateEIP6492Offchain(ctx, provider, signer, hash, signature, nil)
	assert.False(t, valid, "invalid payload must not validate")
	// err may be nil (contract returned false) or non-nil (execution reverted etc.)
	_ = err
}
