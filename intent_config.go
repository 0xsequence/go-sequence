package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// Token represents a token with an address and chain ID. Zero addresses represent ETH, or other native tokens.
type OriginToken struct {
	Address common.Address `abi:"address"`
	ChainId *big.Int       `abi:"chainId"`
}

type DestinationToken struct {
	Address common.Address `abi:"address"`
	ChainId *big.Int       `abi:"chainId"`
	Amount  *big.Int       `abi:"amount"`
}

// MalleableSapientConfig contains the configuration for a MalleableSapient signer for a single call payload.
// Each call payload can have its own sapient address, signature, and provider (for different chains).
type MalleableSapientConfig struct {
	Address   common.Address   // The MalleableSapient contract address
	Signature []byte           // Signature describing which parts are static vs malleable
	Provider  *ethrpc.Provider // Provider for the chain this call is on
}

// IntentParams is a new version of intent parameters that uses CallsPayload for destination calls.
type IntentParams struct {
	UserAddress       common.Address
	Nonce             *big.Int
	OriginTokens      []OriginToken
	DestinationTokens []DestinationToken
	DestinationCalls  []*v3.CallsPayload
}

// HashIntentParams generates a unique bytes32 hash from the IntentParams struct.
func HashIntentParams(params *IntentParams) ([32]byte, error) {
	if params == nil {
		return [32]byte{}, fmt.Errorf("params is nil")
	}
	if params.UserAddress == (common.Address{}) {
		return [32]byte{}, fmt.Errorf("UserAddress is zero")
	}
	if params.Nonce == nil {
		return [32]byte{}, fmt.Errorf("Nonce is nil")
	}
	if len(params.OriginTokens) == 0 {
		return [32]byte{}, fmt.Errorf("OriginTokens is empty")
	}
	if len(params.DestinationCalls) == 0 {
		return [32]byte{}, fmt.Errorf("DestinationCalls is empty")
	}
	if len(params.DestinationTokens) == 0 {
		return [32]byte{}, fmt.Errorf("DestinationTokens is empty")
	}
	for i, call := range params.DestinationCalls {
		if call == nil {
			return [32]byte{}, fmt.Errorf("DestinationCalls[%d] is nil", i)
		}
	}

	// Calculate cumulativeCallsHash
	var cumulativeCallsHash [32]byte
	for i, callPayload := range params.DestinationCalls {
		_ = i

		// Change the address to address(0)
		callPayload, _ := v3.PayloadWithAddress(callPayload, common.Address{})

		individualPayloadDigest := callPayload.Digest()

		packedForKeccak := append(cumulativeCallsHash[:], individualPayloadDigest.Hash[:]...)
		// fmt.Printf("    packedForKeccak: 0x%s\n", common.Bytes2Hex(packedForKeccak))

		cumulativeCallsHash = ethcoder.Keccak256Hash(packedForKeccak)
		// fmt.Printf("    cumulativeCallsHash after callPayload[%d]: 0x%s\n", i, common.Bytes2Hex(cumulativeCallsHash[:]))
	}
	// fmt.Printf("  final cumulativeCallsHash: 0x%s\n", common.Bytes2Hex(cumulativeCallsHash[:]))
	// fmt.Printf("  destinationTokens: %v\n", params.DestinationTokens)

	// Prepare OriginTokens data
	originArgs := []OriginToken{}
	for _, originToken := range params.OriginTokens {
		originArgs = append(originArgs, OriginToken{
			Address: originToken.Address,
			ChainId: originToken.ChainId,
		})
	}

	// Prepare DestinationTokens data
	destinationArgs := []DestinationToken{}
	for _, destinationToken := range params.DestinationTokens {
		destinationArgs = append(destinationArgs, DestinationToken{
			Address: destinationToken.Address,
			ChainId: destinationToken.ChainId,
			Amount:  destinationToken.Amount,
		})
	}

	// Define ABI types for the combined packing
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	nonceType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create nonce ABI type: %w", err)
	}

	originTokensComponents := []abi.ArgumentMarshaling{
		{Name: "address", Type: "address"},
		{Name: "chainId", Type: "uint256"},
	}
	originTokensListType, err := abi.NewType("tuple[]", "", originTokensComponents)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create origin tokens list ABI type: %w", err)
	}

	destinationTokensComponents := []abi.ArgumentMarshaling{
		{Name: "address", Type: "address"},
		{Name: "chainId", Type: "uint256"},
		{Name: "amount", Type: "uint256"},
	}
	destinationTokensListType, err := abi.NewType("tuple[]", "", destinationTokensComponents)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create destination tokens list ABI type: %w", err)
	}

	bytes32Type, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create bytes32 ABI type: %w", err)
	}

	fullArguments := abi.Arguments{
		{Name: "userAddress", Type: addressType},
		{Name: "nonce", Type: nonceType},
		{Name: "originTokens", Type: originTokensListType},
		{Name: "destinationTokens", Type: destinationTokensListType},
		{Name: "cumulativeCallsHash", Type: bytes32Type},
	}

	// ABI encode all fields in one go
	encoded, err := fullArguments.Pack(
		params.UserAddress,
		params.Nonce,
		originArgs,
		destinationArgs,
		cumulativeCallsHash,
	)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to ABI pack combined arguments: %w", err)
	}
	// fmt.Printf("    encoded: 0x%s\n", common.Bytes2Hex(encoded))

	hash := ethcoder.Keccak256(encoded)

	var hash32 [32]byte
	copy(hash32[:], hash)
	return hash32, nil
}

// `CreateAnyAddressSubdigestTree` iterates over each batch of payloads,
// validates that each call in the payload meets the following criteria:
//   - GasLimit must be 0,
//
// For each valid batch, it creates a bundle, computes its digest,
// and creates a new WalletConfigTreeSubdigestLeaf.
func CreateAnyAddressSubdigestTree(calls []*v3.CallsPayload) ([]v3.WalletConfigTree, error) {
	var leaves []v3.WalletConfigTree

	for batchIndex, call := range calls {
		// Validate each call in the payload
		for j, call := range call.Calls {
			if call.GasLimit != nil && call.GasLimit.Cmp(big.NewInt(0)) != 0 {
				return nil, fmt.Errorf("batch %d, call %d: GasLimit must be 0", batchIndex, j)
			}
		}

		digest := call.Digest()

		// Create a subdigest leaf with the computed digest.
		leaves = append(leaves, &v3.WalletConfigTreeAnyAddressSubdigestLeaf{Digest: digest.Hash})
	}

	return leaves, nil
}

// `CreateIntentTree` creates a tree from a list of intent operations and a main signer address.
func CreateIntentTree(mainSigner common.Address, calls []*v3.CallsPayload, sapientSignerLeafNode v3.WalletConfigTree) (*v3.WalletConfigTree, error) {
	var leaves []v3.WalletConfigTree

	if len(calls) > 0 {
		// Create the subdigest leaves from the batched transactions.
		subdigestLeaves, err := CreateAnyAddressSubdigestTree(calls)
		if err != nil {
			return nil, err
		}
		leaves = append(leaves, subdigestLeaves...)
	}

	// If the sapient signer leaf is not nil, add it to the leaves.
	if sapientSignerLeafNode != nil {
		leaves = append(leaves, sapientSignerLeafNode)
	}

	// Create the main signer leaf (with weight 1).
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	if len(leaves) == 0 {
		return nil, fmt.Errorf("no leaves to create tree from")
	}

	// If the length of the leaves is 1
	if len(leaves) == 1 {
		tree := v3.WalletConfigTreeNodes(mainSignerLeaf, leaves[0])
		return &tree, nil
	}

	// Create a tree from the subdigest leaves.
	tree := v3.WalletConfigTreeNodes(leaves...)

	// Construct the new wallet config using:
	fullTree := v3.WalletConfigTreeNodes(mainSignerLeaf, tree)

	return &fullTree, nil
}

// `CreateIntentConfiguration` creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateIntentConfiguration(mainSigner common.Address, calls []*v3.CallsPayload, sapientSignerLeafNode v3.WalletConfigTree) (*v3.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	tree, err := CreateIntentTree(mainSigner, calls, sapientSignerLeafNode)
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

// `CreateIntentConfigurationWithMalleableSapient` creates a wallet configuration with optional MalleableSapient configuration
func CreateIntentConfigurationWithMalleableSapient(
	ctx context.Context,
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	configs []MalleableSapientConfig,
) (*v3.WalletConfig, error) {
	if len(calls) == 0 {
		return nil, fmt.Errorf("calls cannot be empty")
	}

	if len(configs) != len(calls) {
		return nil, fmt.Errorf("configs length (%d) must match calls length (%d)", len(configs), len(calls))
	}

	var callsWithoutMalleableSapient []*v3.CallsPayload
	var sapientSignerLeaves []v3.WalletConfigTree
	for i, callPayload := range calls {
		config := configs[i]

		// Skip if address is zero address. Will be handled by CreateIntentConfiguration.
		if config.Address == (common.Address{}) {
			callsWithoutMalleableSapient = append(callsWithoutMalleableSapient, callPayload)
			continue
		}

		if config.Provider == nil {
			return nil, fmt.Errorf("provider is required for configs[%d]", i)
		}

		// Recover the image hash for the payload and configuration
		imageHash, err := v3.RecoverSapientSignature(ctx, config.Address, *callPayload, config.Signature, config.Provider)
		if err != nil {
			return nil, fmt.Errorf("failed to recover sapient signature for configs[%d]: %w", i, err)
		}

		// Create a SapientSignerLeaf for this payload
		sapientSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     1,
			Address:    config.Address,
			ImageHash_: imageHash,
		}
		sapientSignerLeaves = append(sapientSignerLeaves, sapientSignerLeaf)
	}

	// Create a tree from all the SapientSignerLeaves
	var sapientSignerTree v3.WalletConfigTree
	if len(sapientSignerLeaves) > 0 {
		sapientSignerTree = v3.WalletConfigTreeNodes(sapientSignerLeaves...)
	}

	// Remaining calls go to CreateIntentConfiguration
	return CreateIntentConfiguration(mainSigner, callsWithoutMalleableSapient, sapientSignerTree)
}

// `GetIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation.
// The signature is based on the transaction bundle digests only.
func GetIntentConfigurationSignature(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
) ([]byte, error) {
	// Default case without any sapient signer
	config, err := CreateIntentConfiguration(mainSigner, calls, nil)
	if err != nil {
		return nil, err
	}

	// spew.Dump(config)
	// spew.Dump(config.Tree)

	signingFunc := func(ctx context.Context, signer core.Signer, _ []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		// For mainSigner or other signers, we don't provide a signature here.
		// This will result in an AddressLeaf or NodeLeaf in the signature tree.
		return 0, nil, nil
	}

	// Build the signature using BuildNoChainIDSignature, which allows us to inject custom signatures via SigningFunction.
	// Set validateSigningPower to false, as we are not necessarily providing signatures for all parts of the config.
	sig, err := config.BuildRegularSignature(context.Background(), signingFunc, false)
	if err != nil {
		return nil, fmt.Errorf("failed to build regular signature: %w", err)
	}

	// spew.Dump(sig)

	if regularSig, ok := sig.(*v3.RegularSignature); ok {
		if regularSig.Signature != nil {
			signatureTree := regularSig.Signature.Tree
			_ = signatureTree
			// fmt.Println("Accessing sig.Signature.Tree:")
			// spew.Dump(signatureTree)
		} else {
			// fmt.Println("sig.Signature is nil")
		}
	} else if noChainIdSig, ok := sig.(*v3.NoChainIDSignature); ok {
		if noChainIdSig.Signature != nil {
			signatureTree := noChainIdSig.Signature.Tree
			_ = signatureTree
			// fmt.Println("Accessing sig.Signature.Tree (NoChainID):")
			// spew.Dump(signatureTree)
		} else {
			// fmt.Println("sig.Signature is nil for NoChainIDSignature")
		}
	} else {
		// fmt.Printf("sig is not of type *v3.RegularSignature or *v3.NoChainIDSignature, it is %T\n", sig)
	}

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	// Print the signature data
	// fmt.Printf("signature data: %s\n", common.Bytes2Hex(data))

	return data, nil
}

// // replaceSapientSignerWithNodeInConfigTree recursively traverses the WalletConfigTree.
// func replaceSapientSignerWithNodeInConfigTree(tree v3.WalletConfigTree) v3.WalletConfigTree {
// 	if tree == nil {
// 		return nil
// 	}

// 	switch node := tree.(type) {
// 	case *v3.WalletConfigTreeNode:
// 		// Recursively call on left and right children
// 		left := replaceSapientSignerWithNodeInConfigTree(node.Left)
// 		right := replaceSapientSignerWithNodeInConfigTree(node.Right)

// 		if left == node.Left && right == node.Right {
// 			return node
// 		}
// 		return &v3.WalletConfigTreeNode{Left: left, Right: right}

// 	case *v3.WalletConfigTreeNestedLeaf:
// 		// Recursively call on the inner tree
// 		innerTree := replaceSapientSignerWithNodeInConfigTree(node.Tree)

// 		if innerTree == node.Tree { // Check for pointer equality
// 			return node // No change, return original
// 		}
// 		return &v3.WalletConfigTreeNestedLeaf{
// 			Weight:    node.Weight,
// 			Threshold: node.Threshold,
// 			Tree:      innerTree,
// 		}

// 	case *v3.WalletConfigTreeSapientSignerLeaf:
// 		// This is the target node type to replace
// 		return &v3.WalletConfigTreeNodeLeaf{Node: node.ImageHash()}

// 	default:
// 		return tree
// 	}
// }
