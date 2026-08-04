package v3_test

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

// A subdigest leaf reports max signersWeight for any payload, so the config threshold
// looks met as soon as the first signature is collected. Early cancellation on that
// estimate nondeterministically drops the other signer's signature even though recovery
// of a payload not matching the subdigest still needs it. Every build must embed both
// signatures and produce identical bytes.
func TestBuildRegularSignatureCollectsAllSignersDespiteSubdigestLeaf(t *testing.T) {
	signerA := common.HexToAddress("0x1111111111111111111111111111111111111111")
	signerB := common.HexToAddress("0x2222222222222222222222222222222222222222")

	dummySignature := func(fill byte) []byte {
		sig := make([]byte, 65)
		for i := range 64 {
			sig[i] = fill
		}
		sig[64] = 27
		return sig
	}
	signatureA := dummySignature(0xaa)
	signatureB := dummySignature(0xbb)

	config := &v3.WalletConfig{
		Threshold_: 2,
		Tree: v3.WalletConfigTreeNodes(
			v3.WalletConfigTreeSubdigestLeaf{Subdigest: common.BigToHash(big.NewInt(1))},
			&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: signerA},
			&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: signerB},
		),
	}

	signatures := map[common.Address][]byte{
		signerA: signatureA,
		signerB: signatureB,
	}
	signingFunc := func(ctx context.Context, signer core.Signer, _ []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		if sig, ok := signatures[signer.Address]; ok {
			return core.SignerSignatureTypeEthSign, sig, nil
		}
		return 0, nil, nil
	}

	var first []byte
	for range 100 {
		sig, err := config.BuildRegularSignature(context.Background(), signingFunc, true)
		require.NoError(t, err)
		data, err := sig.Data()
		require.NoError(t, err)
		require.True(t, bytes.Contains(data, signatureA[:64]), "signer A's signature must be embedded on every build")
		require.True(t, bytes.Contains(data, signatureB[:64]), "signer B's signature must be embedded on every build")
		if first == nil {
			first = data
		}
		require.Equal(t, first, data, "signature encoding must be deterministic")
	}
}

// A subdigest leaf reports max signersWeight even for an empty signer set, so
// signing-power validation must still fail when no signature is collected at all.
func TestBuildSignatureValidationRejectsEmptySignerSetDespiteSubdigestLeaf(t *testing.T) {
	config := &v3.WalletConfig{
		Threshold_: 2,
		Tree: v3.WalletConfigTreeNodes(
			v3.WalletConfigTreeSubdigestLeaf{Subdigest: common.BigToHash(big.NewInt(1))},
			&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x1111111111111111111111111111111111111111")},
		),
	}

	signingFunc := func(ctx context.Context, signer core.Signer, _ []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		return 0, nil, nil
	}

	_, err := config.BuildRegularSignature(context.Background(), signingFunc, true)
	require.ErrorContains(t, err, "not enough signers")

	_, err = config.BuildNoChainIDSignature(context.Background(), signingFunc, true)
	require.ErrorContains(t, err, "not enough signers")
}
