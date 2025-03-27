package sequence

import (
	"fmt"
	"log"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
)

// `IntentOperation` represents a single operation in an intent.
// It contains the chain ID + the calls to be executed in the operation.
// The space and nonce are optional, and if not provided, the default values will be used.
type IntentOperation struct {
	chainId *big.Int // The chain ID of the operation, required
	space   *big.Int // The space of the operation, optional
	nonce   *big.Int // The nonce of the operation, optional

	calls []v3.Call
}

// `NewIntentOperation` creates a new `IntentOperation` with the given chain ID and calls.
func NewIntentOperation(chainId *big.Int, calls []v3.Call, space *big.Int, nonce *big.Int) *IntentOperation {
	return &IntentOperation{
		chainId: chainId,
		calls:   calls,
		space:   space,
		nonce:   nonce,
	}
}

// `CreateIntentCallsPayload` creates a bundle of transactions with the gas limit 0 and the initial nonce 0
func CreateIntentCallsPayload(op *IntentOperation, target ...common.Address) (v3.CallsPayload, error) {
	// If the chainId is not provided, throw an error
	if op.chainId == nil {
		return v3.CallsPayload{}, fmt.Errorf("chainId is required")
	}

	// If the space is not provided, use the default value
	if op.space == nil {
		op.space = big.NewInt(0)
	}

	// If the nonce is not provided, use the default value
	if op.nonce == nil {
		op.nonce = big.NewInt(0)
	}

	// If the target is specified, get the first element
	var addr common.Address
	if target != nil {
		addr = target[0]
	}

	// Construct the payload with the correct address, chain ID, and calls.
	bundle := v3.NewCallsPayload(addr, op.chainId, op.calls, op.space, op.nonce)

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
func CreateIntentDigestTree(ops []*IntentOperation) (*v3.WalletConfigTree, error) {
	var leaves []*v3.WalletConfigTreeAnyAddressSubdigestLeaf

	for batchIndex, op := range ops {
		// Validate each call in the payload
		for j, call := range op.calls {
			if call.GasLimit != nil && call.GasLimit.Cmp(big.NewInt(0)) != 0 {
				return nil, fmt.Errorf("batch %d, call %d: GasLimit must be 0", batchIndex, j)
			}
		}

		// Create the intent bundle for this batch.
		bundle, err := CreateIntentCallsPayload(op)
		if err != nil {
			return nil, fmt.Errorf("failed to create intent bundle for batch %d: %w", batchIndex, err)
		}
		spew.Dump(bundle)

		digest := bundle.Digest()

		log.Println("CreateIntentDigestTree digest:", digest.Hash)

		// Create a subdigest leaf with the computed digest.
		leaf := &v3.WalletConfigTreeAnyAddressSubdigestLeaf{
			Digest: core.Subdigest{Hash: digest.Hash},
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

// CreateIntentTree creates a tree from a list of intent operations and a main signer address.
func CreateIntentTree(mainSigner common.Address, ops []*IntentOperation) (*v3.WalletConfigTree, error) {
	digestTree, err := CreateIntentDigestTree(ops)
	if err != nil {
		return nil, err
	}

	// Create the main signer leaf (with weight 1).
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Construct the new wallet config using:
	fullTree := v3.WalletConfigTreeNodes(mainSignerLeaf, *digestTree)

	return &fullTree, nil
}

// CreateIntentConfiguration creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateIntentConfiguration(mainSigner common.Address, ops []*IntentOperation) (*v3.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	tree, err := CreateIntentTree(mainSigner, ops)
	if err != nil {
		return nil, err
	}

	// Construct the new wallet config using:
	config := &v3.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree:        *tree,
	}

	return config, nil
}

// `GetIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation. The signature is based on the transaction bundle digests only.
func GetIntentConfigurationSignature(mainSigner common.Address, ops []*IntentOperation) ([]byte, error) {
	// Create the intent configuration using the batched transactions.
	config, err := CreateIntentConfiguration(mainSigner, ops)
	if err != nil {
		return nil, err
	}

	// Default to building the regular signature
	sig, err := config.BuildSubdigestSignature(false)
	if err != nil {
		return nil, fmt.Errorf("failed to build subdigest signature: %w", err)
	}

	spew.Dump(config.Tree)

	log.Printf("sig: %s", sig)

	spew.Dump(sig)

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	return data, nil
}
