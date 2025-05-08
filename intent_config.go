package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// IntentParamsV2 is a new version of intent parameters that uses CallsPayload for destination calls.
type IntentParams struct {
	UserAddress        common.Address
	OriginChainId      uint64
	OriginTokenAddress common.Address
	DestinationCalls   []*v3.CallsPayload
}

// HashIntentParams generates a unique bytes32 hash from the IntentParams struct.
func HashIntentParams(params *IntentParams) ([32]byte, error) {
	if params == nil {
		return [32]byte{}, fmt.Errorf("params is nil")
	}
	if params.UserAddress == (common.Address{}) {
		return [32]byte{}, fmt.Errorf("UserAddress is zero")
	}
	if params.OriginChainId == 0 {
		return [32]byte{}, fmt.Errorf("OriginChainId is zero")
	}
	if params.OriginTokenAddress == (common.Address{}) {
		return [32]byte{}, fmt.Errorf("OriginTokenAddress is zero")
	}
	if len(params.DestinationCalls) == 0 {
		return [32]byte{}, fmt.Errorf("DestinationCalls is empty")
	}
	for i, call := range params.DestinationCalls {
		if call == nil {
			return [32]byte{}, fmt.Errorf("DestinationCalls[%d] is nil", i)
		}
	}

	// ABI encode the fields in a deterministic order
	var (
		callsEncoded []byte
	)
	for _, call := range params.DestinationCalls {
		// Use the Encode method to get the byte encoding of each CallsPayload
		callsEncoded = append(callsEncoded, call.Encode(common.Address{})...) // Use AnyAddressSubdigestLeaf as the digest hash
	}

	encoded, err := ethcoder.ABIPackArguments(
		[]string{
			"address", // UserAddress
			"uint256", // OriginChainId
			"address", // OriginTokenAddress
			"bytes",   // DestinationCalls
		},
		[]interface{}{
			params.UserAddress,
			big.NewInt(int64(params.OriginChainId)),
			params.OriginTokenAddress,
			callsEncoded,
		},
	)
	if err != nil {
		return [32]byte{}, err
	}

	hash := ethcoder.Keccak256(encoded)
	var hash32 [32]byte
	copy(hash32[:], hash)
	return hash32, nil
}

// `CreateIntentDigestTree` iterates over each batch of payloads,
// validates that each call in the payload meets the following criteria:
//   - GasLimit must be 0,
//
// For each valid batch, it creates a bundle, computes its digest,
// and creates a new WalletConfigTreeSubdigestLeaf.
func CreateIntentDigestTree(calls []*v3.CallsPayload) (*v3.WalletConfigTree, error) {
	var leaves []*v3.WalletConfigTreeAnyAddressSubdigestLeaf

	for batchIndex, call := range calls {
		// Validate each call in the payload
		for j, call := range call.Calls {
			if call.GasLimit != nil && call.GasLimit.Cmp(big.NewInt(0)) != 0 {
				return nil, fmt.Errorf("batch %d, call %d: GasLimit must be 0", batchIndex, j)
			}
		}

		digest := call.Digest()

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
func CreateIntentTree(mainSigner common.Address, calls []*v3.CallsPayload) (*v3.WalletConfigTree, error) {
	digestTree, err := CreateIntentDigestTree(calls)
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
func CreateIntentConfiguration(mainSigner common.Address, calls []*v3.CallsPayload) (*v3.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	tree, err := CreateIntentTree(mainSigner, calls)
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
func GetIntentConfigurationSignature(mainSigner common.Address, calls []*v3.CallsPayload) ([]byte, error) {
	// Create the intent configuration using the batched transactions.
	config, err := CreateIntentConfiguration(mainSigner, calls)
	if err != nil {
		return nil, err
	}

	// Default to building the regular signature
	sig, err := config.BuildSubdigestSignature(false)
	if err != nil {
		return nil, fmt.Errorf("failed to build subdigest signature: %w", err)
	}

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	return data, nil
}
