package sequence

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
)

// `CreateIntentBundle` creates a bundle of transactions with the gas limit 0 and the initial nonce 0
func CreateIntentBundle(payload v3.DecodedPayload) (v3.DecodedPayload, error) {
	bundle := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls:     payload.Calls,
	}

	// Set consistent gas limit to 0 for all calls
	for i := range bundle.Calls {
		bundle.Calls[i].GasLimit = big.NewInt(0)
	}

	return bundle, nil
}

// `CreateIntentDigestTree` iterates over each batch of payloads,
// validates that each call in the payload meets the following criteria:
//   - GasLimit must be 0,
//
// For each valid batch, it creates a bundle, computes its digest,
// and creates a new WalletConfigTreeSubdigestLeaf.
func CreateIntentDigestTree(batchPayloads []v3.DecodedPayload) (*v3.WalletConfigTree, error) {
	var leaves []*v3.WalletConfigTreeAnyAddressSubdigestLeaf

	for batchIndex, payload := range batchPayloads {
		// Validate each call in the payload
		for j, call := range payload.Calls {
			if call.GasLimit != nil && call.GasLimit.Cmp(big.NewInt(0)) != 0 {
				return nil, fmt.Errorf("batch %d, call %d: GasLimit must be 0", batchIndex, j)
			}
		}

		// Create the intent bundle for this batch.
		bundle, err := CreateIntentBundle(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to create intent bundle for batch %d: %w", batchIndex, err)
		}

		// Use a zero address and no chain ID since we're constructing a `noChainID` and `AnyAddressSubdigestLeaf` payload which hashes w/ zero address
		noChainID := big.NewInt(0)
		zeroAddress := common.Address{}

		// Make sure all required fields for the payload are set
		if bundle.Space == nil {
			bundle.Space = big.NewInt(0)
		}
		if bundle.Nonce == nil {
			bundle.Nonce = big.NewInt(0)
		}
		if bundle.ParentWallets == nil {
			bundle.ParentWallets = []common.Address{}
		}

		digest, err := v3.HashPayload(zeroAddress, noChainID, bundle)
		if err != nil {
			return nil, fmt.Errorf("failed to compute digest for batch %d: %w", batchIndex, err)
		}

		log.Println("CreateIntentDigestTree digest:", common.Bytes2Hex(digest[:]))

		// Create a subdigest leaf with the computed digest.
		leaf := &v3.WalletConfigTreeAnyAddressSubdigestLeaf{
			Digest: core.Subdigest{Hash: common.BytesToHash(digest[:])},
		}
		leaves = append(leaves, leaf)
	}

	// If the length of the leaves is 1, return the leaf as the tree.
	if len(leaves) == 1 {
		tree := v3.WalletConfigTree(leaves[0])
		return &tree, nil
	}

	// Convert the slice of subdigest leaves into a slice of v3.WalletConfigTree.
	var treeNodes []v3.WalletConfigTree
	for _, leaf := range leaves {
		treeNodes = append(treeNodes, leaf)
	}

	// Create a tree from the subdigest leaves.
	tree := v3.WalletConfigTreeNodes(treeNodes...)

	return &tree, nil
}

// CreateIntentConfiguration creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateIntentConfiguration(mainSigner common.Address, batches []v3.DecodedPayload) (*v3.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	tree, err := CreateIntentDigestTree(batches)
	if err != nil {
		return nil, err
	}

	// Create the main signer leaf (with weight 1).
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Construct the new wallet config using:
	config := &v3.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: v3.WalletConfigTreeNodes(
			mainSignerLeaf,
			*tree,
		),
	}

	return config, nil
}

// `CreateIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation. The signature is based on the transaction bundle digests only.
func CreateIntentConfigurationSignature(mainSigner common.Address, batches []v3.DecodedPayload) ([]byte, error) {
	// Create the intent configuration using the batched transactions.
	config, err := CreateIntentConfiguration(mainSigner, batches)
	if err != nil {
		return nil, err
	}

	// Use BuildNoChainIDSignature with an empty signing function
	// Since we want the address leaves without signatures, we'll just return ErrSigningNoSigner
	// This causes the signature tree builder to include the address leaf without a signature
	sig, err := config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		// Return error to indicate we don't want to sign
		return 0, nil, core.ErrSigningNoSigner
	}, false)
	if err != nil {
		return nil, fmt.Errorf("failed to build no chain ID signature: %w", err)
	}

	spew.Dump(sig)

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	return data, nil
}
