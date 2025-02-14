package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v2 "github.com/0xsequence/go-sequence/core/v2"
)

// CreateIntentConfiguration creates a wallet configuration where the intent's transactions are grouped into the initial subdigest
func CreateIntentConfiguration(mainSigner common.Address, txns []*Transaction) (*v2.WalletConfig, error) {
	// Create the main signer leaf (with weight 1)
	mainSignerLeaf := &v2.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Create a bundle of transactions with the initial nonce.
	bundle := Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Transactions:  txns,
		Nonce:         big.NewInt(0),
	}

	// Compute the digest for the bundle. This digest represents the intent subdigest.
	digest, err := bundle.Digest()
	if err != nil {
		return nil, fmt.Errorf("failed to compute bundle digest: %w", err)
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
