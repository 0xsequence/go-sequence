package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v2 "github.com/0xsequence/go-sequence/core/v2"
)

// createIntentBundle creates a bundle of transactions with the initial nonce 0
func createIntentBundle(txns []*Transaction) (common.Hash, error) {
	bundle := Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Transactions:  txns,
		Nonce:         big.NewInt(0),
	}

	// Compute the digest for the bundle
	digest, err := bundle.Digest()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to compute bundle digest: %w", err)
	}

	return digest, nil
}

// CreateIntentConfiguration creates a wallet configuration where the intent's transactions are grouped into the initial subdigest
func CreateIntentConfiguration(mainSigner common.Address, txns []*Transaction) (*v2.WalletConfig, error) {
	// Create the bundle digest
	digest, err := createIntentBundle(txns)
	if err != nil {
		return nil, err
	}

	// Create the main signer leaf (with weight 1)
	mainSignerLeaf := &v2.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Create a leaf node from the computed subdigest.
	subdigestLeaf := &v2.WalletConfigTreeSubdigestLeaf{
		Subdigest: core.Subdigest{Hash: digest},
	}

	// Construct the new wallet config using the main signer (as the first leaf) and a tree node for the subdigest.
	config := &v2.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: v2.WalletConfigTreeNodes(
			mainSignerLeaf,
			v2.WalletConfigTreeNodes(subdigestLeaf),
		),
	}

	return config, nil
}

// CreateIntentConfigurationSignature creates a signature for the intent configuration that can be used
// to bypass chain ID validation. The signature is based on the transaction digest only.
func CreateIntentConfigurationSignature(mainSigner common.Address, txns []*Transaction) ([]byte, error) {
	// Create the bundle digest
	digest, err := createIntentBundle(txns)
	if err != nil {
		return nil, err
	}

	// Create a temporary wallet config for signing
	config := &v2.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &v2.WalletConfigTreeSubdigestLeaf{
			Subdigest: core.Subdigest{Hash: digest},
		},
	}

	// Build a no chain ID signature
	sig, err := config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		return core.SignerSignatureTypeEIP712, nil, nil
	}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to build no chain ID signature: %w", err)
	}

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	return data, nil
}
