package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v2 "github.com/0xsequence/go-sequence/core/v2"
)

// `CreateIntentBundle` creates a bundle of transactions with the initial nonce 0
func CreateIntentBundle(txns []*Transaction) (Transaction, error) {
	bundle := Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Transactions:  txns,
		Nonce:         big.NewInt(0),
	}

	return bundle, nil
}

// CreateIntentSubdigestLeaves iterates over the provided transactions,
// validates that each transaction meets the following criteria:
//   - DelegateCall must be false,
//   - RevertOnError must be true,
//   - Nonce must be non-nil and equal to 0.
//
// For each valid transaction, it computes the digest and creates a new WalletConfigTreeSubdigestLeaf.
func CreateIntentSubdigestLeaves(txns []*Transaction) ([]*v2.WalletConfigTreeSubdigestLeaf, error) {
	var leaves []*v2.WalletConfigTreeSubdigestLeaf

	for i, tx := range txns {
		// Validate the transaction fields:
		if tx.DelegateCall {
			return nil, fmt.Errorf("transaction at index %d: DelegateCall must be false", i)
		}
		if !tx.RevertOnError {
			return nil, fmt.Errorf("transaction at index %d: RevertOnError must be true", i)
		}
		if tx.Nonce == nil {
			return nil, fmt.Errorf("transaction at index %d: Nonce is nil", i)
		}
		if tx.Nonce.Cmp(big.NewInt(0)) != 0 {
			return nil, fmt.Errorf("transaction at index %d: Nonce must be 0", i)
		}

		// Create the intent bundle.
		bundle, err := CreateIntentBundle(txns)
		if err != nil {
			return nil, fmt.Errorf("failed to create intent bundle: %w", err)
		}

		// Compute digest of the transaction.
		digest, err := bundle.Digest()
		if err != nil {
			return nil, fmt.Errorf("failed to compute digest for transaction at index %d: %w", i, err)
		}

		// Create a subdigest leaf with the computed digest.
		leaf := &v2.WalletConfigTreeSubdigestLeaf{
			Subdigest: core.Subdigest{Hash: digest},
		}
		leaves = append(leaves, leaf)
	}

	return leaves, nil
}

// CreateIntentConfiguration creates a wallet configuration where the intent's transactions
// are grouped into the initial subdigest. It uses the subdigest leaves constructed from the txns.
func CreateIntentConfiguration(mainSigner common.Address, txns []*Transaction) (*v2.WalletConfig, error) {
	// Create the subdigest leaves from the transactions.
	leaves, err := CreateIntentSubdigestLeaves(txns)
	if err != nil {
		return nil, err
	}

	// Convert the slice of subdigest leaves into a slice of v2.WalletConfigTree.
	var treeNodes []v2.WalletConfigTree
	for _, leaf := range leaves {
		treeNodes = append(treeNodes, leaf)
	}

	// Create the main signer leaf (with weight 1).
	mainSignerLeaf := &v2.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Construct the new wallet config using:
	//  • the main signer leaf as the left branch, and
	//  • a nested tree node containing the subdigest leaves as the right branch.
	config := &v2.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: v2.WalletConfigTreeNodes(
			mainSignerLeaf,
			v2.WalletConfigTreeNodes(treeNodes...),
		),
	}

	return config, nil
}

// `CreateIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation. The signature is based on the transaction digest only.
func CreateIntentConfigurationSignature(mainSigner common.Address, txns []*Transaction) ([]byte, error) {
	// Create the intent configuration
	config, err := CreateIntentConfiguration(mainSigner, txns)
	if err != nil {
		return nil, err
	}

	// Build a no chain ID signature with an empty callback function
	sig, err := config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		// For all signers, return a zero value, since intent signatures are not signed by the main signer
		return 0, nil, nil
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
