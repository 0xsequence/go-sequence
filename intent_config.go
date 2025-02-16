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

// `CreateIntentSubdigestLeaves` iterates over each batch of transactions,
// validates that each transaction in the batch meets the following criteria:
//   - DelegateCall must be false,
//   - RevertOnError must be true,
//   - Nonce must be non-nil and equal to 0.
//
// For each valid batch, it creates a bundle of transactions, computes its digest,
// and creates a new WalletConfigTreeSubdigestLeaf.
func CreateIntentSubdigestLeaves(batchTxns [][]*Transaction) ([]*v2.WalletConfigTreeSubdigestLeaf, error) {
	var leaves []*v2.WalletConfigTreeSubdigestLeaf

	for batchIndex, batch := range batchTxns {
		// Optional: Ensure the batch is not empty.
		if len(batch) == 0 {
			return nil, fmt.Errorf("batch at index %d is empty", batchIndex)
		}

		// Validate each transaction in the batch.
		for j, tx := range batch {
			if tx.DelegateCall {
				return nil, fmt.Errorf("batch %d, transaction %d: DelegateCall must be false", batchIndex, j)
			}
			if !tx.RevertOnError {
				return nil, fmt.Errorf("batch %d, transaction %d: RevertOnError must be true", batchIndex, j)
			}
			if tx.Nonce == nil {
				return nil, fmt.Errorf("batch %d, transaction %d: Nonce is nil", batchIndex, j)
			}
			if tx.Nonce.Cmp(big.NewInt(0)) != 0 {
				return nil, fmt.Errorf("batch %d, transaction %d: Nonce must be 0", batchIndex, j)
			}
		}

		// Create the intent bundle for this batch.
		bundle, err := CreateIntentBundle(batch)
		if err != nil {
			return nil, fmt.Errorf("failed to create intent bundle for batch %d: %w", batchIndex, err)
		}

		// Compute digest of the bundle.
		digest, err := bundle.Digest()
		if err != nil {
			return nil, fmt.Errorf("failed to compute digest for batch %d: %w", batchIndex, err)
		}

		// Create a subdigest leaf with the computed digest.
		leaf := &v2.WalletConfigTreeSubdigestLeaf{
			Subdigest: core.Subdigest{Hash: digest},
		}
		leaves = append(leaves, leaf)
	}

	return leaves, nil
}

// CreateIntentConfiguration creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateIntentConfiguration(mainSigner common.Address, batches [][]*Transaction) (*v2.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	leaves, err := CreateIntentSubdigestLeaves(batches)
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

// `CreateIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation. The signature is based on the transaction bundle digests only.
func CreateIntentConfigurationSignature(mainSigner common.Address, batches [][]*Transaction) ([]byte, error) {
	// Create the intent configuration using the batched transactions.
	config, err := CreateIntentConfiguration(mainSigner, batches)
	if err != nil {
		return nil, err
	}

	// Build a no chain ID signature with an empty callback function.
	sig, err := config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		// For all signers, return a zero value, since intent signatures are not signed by the main signer.
		return 0, nil, nil
	}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to build no chain ID signature: %w", err)
	}

	// Get the signature data.
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	return data, nil
}
