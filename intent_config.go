package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
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

var maxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
var maxUint64 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 64), big.NewInt(1))

// wrapGate requires gateLeaf's co-signature alongside any one of
// gateableLeaves. An inner threshold-1 nest OR's the groups and contributes weight 1 at
// most, so satisfying many groups still cannot clear the outer threshold without the gate
// leaf. The outer threshold is gateLeaf's weight + 1, so any signer-leaf gate
// weight is safe by construction (the gate alone cannot meet it).
func wrapGate(gateLeaf v3.WalletConfigTree, gateableLeaves ...v3.WalletConfigTree) (v3.WalletConfigTree, error) {
	gateSigner, gateWeight, err := signerLeaf(gateLeaf)
	if err != nil {
		return nil, fmt.Errorf("invalid gateLeafNode: %w", err)
	}
	if gateWeight.Sign() <= 0 {
		return nil, fmt.Errorf("invalid gateLeafNode: weight must be > 0")
	}
	if gateWeight.Cmp(maxUint64) > 0 {
		return nil, fmt.Errorf("invalid gateLeafNode: weight is too large")
	}
	// A gated signer leaf sharing the gate's identity satisfies both sides of the outer
	// threshold with one signature, letting the gate authorize alone
	for _, leaf := range gateableLeaves {
		if signer, _, err := signerLeaf(leaf); err == nil && signer == gateSigner {
			return nil, fmt.Errorf("invalid gateLeafNode: gate signer must not appear among gated leaves")
		}
	}
	gateableTree := &v3.WalletConfigTreeNestedLeaf{
		Weight:    1,
		Threshold: 1,
		Tree:      v3.WalletConfigTreeNodes(gateableLeaves...),
	}
	return &v3.WalletConfigTreeNestedLeaf{
		Weight:    1,
		Threshold: uint16(gateWeight.Uint64()) + uint16(gateableTree.Weight),
		Tree:      v3.WalletConfigTreeNodes(gateLeaf, gateableTree),
	}, nil
}

// signerLeaf returns the signer identity and contribution weight of a single terminal leaf
func signerLeaf(tree v3.WalletConfigTree) (core.Signer, *big.Int, error) {
	if tree == nil {
		return core.Signer{}, nil, fmt.Errorf("nil leaf")
	}
	switch t := tree.(type) {
	case *v3.WalletConfigTreeAddressLeaf:
		if t == nil {
			return core.Signer{}, nil, fmt.Errorf("nil leaf")
		}
		return core.Signer{Address: t.Address}, big.NewInt(int64(t.Weight)), nil
	case *v3.WalletConfigTreeSapientSignerLeaf:
		if t == nil {
			return core.Signer{}, nil, fmt.Errorf("nil leaf")
		}
		return core.SapientSigner(t.Address, t.ImageHash_.Hash), big.NewInt(int64(t.Weight)), nil
	case *v3.WalletConfigTreeSubdigestLeaf, v3.WalletConfigTreeSubdigestLeaf,
		*v3.WalletConfigTreeAnyAddressSubdigestLeaf, v3.WalletConfigTreeAnyAddressSubdigestLeaf:
		// Payload-matching leaves report a signerless identity and maxUint256 weight.
		return core.Signer{}, new(big.Int).Set(maxUint256), nil
	default:
		return core.Signer{}, nil, fmt.Errorf("unsupported leaf type %T", tree)
	}
}

// IntentConfigOption configures the optional leaves of an intent configuration tree.
// A nil IntentConfigOption is ignored.
type IntentConfigOption func(*intentConfigOptions)

type intentConfigOptions struct {
	gateLeafNode          v3.WalletConfigTree
	sapientSignerLeafNode v3.WalletConfigTree
}

// WithGate gates the calls and the sapient signer leaf behind leaf's co-signature:
// either group, plus leaf's signature, authorizes the wallet (see wrapGate). The
// main signer is never gated.
func WithGate(leaf v3.WalletConfigTree) IntentConfigOption {
	return func(o *intentConfigOptions) { o.gateLeafNode = leaf }
}

// WithSapientSigner adds leaf (e.g. a timed-refund or gasless-deposit signer) as an
// authorizer alongside the calls' subdigest leaves.
func WithSapientSigner(leaf v3.WalletConfigTree) IntentConfigOption {
	return func(o *intentConfigOptions) { o.sapientSignerLeafNode = leaf }
}

func applyIntentConfigOptions(opts []IntentConfigOption) intentConfigOptions {
	var options intentConfigOptions
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	return options
}

func createIntentTree(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	gateLeafNode v3.WalletConfigTree,
	sapientSignerLeafNode v3.WalletConfigTree,
) (*v3.WalletConfigTree, error) {
	var leaves []v3.WalletConfigTree

	// Create the subdigest leaves from the batched transactions.
	gateableLeaves, err := CreateAnyAddressSubdigestTree(calls)
	if err != nil {
		return nil, err
	}

	// Add the sapient signer leaf to the gateable leaves.
	if sapientSignerLeafNode != nil {
		gateableLeaves = append(gateableLeaves, sapientSignerLeafNode)
	}

	// If there are any gateable leaves, wrap them in a gate if a gate leaf is provided.
	if len(gateableLeaves) > 0 {
		if gateLeafNode == nil {
			// No gate: preserve flat structure so counterfactual addresses stay stable.
			leaves = append(leaves, gateableLeaves...)
		} else {
			// Calls and sapient share one gate; either needs gateLeaf's co-signature.
			gate, err := wrapGate(gateLeafNode, gateableLeaves...)
			if err != nil {
				return nil, err
			}
			leaves = append(leaves, gate)
		}
	}

	// Add the main signer leaf to the leaves (ungated).
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// If the length of the leaves is 1.
	if len(leaves) == 1 {
		tree := v3.WalletConfigTreeNodes(mainSignerLeaf, leaves[0])
		return &tree, nil
	}

	// Create a tree from the (gated) leaves.
	tree := v3.WalletConfigTreeNodes(leaves...)

	// Construct the new wallet config.
	fullTree := v3.WalletConfigTreeNodes(mainSignerLeaf, tree)

	return &fullTree, nil
}

// `CreateIntentTree` creates a tree from a list of intent operations and a main signer
// address. See WithGate and WithSapientSigner for the optional leaves; with no
// options the legacy tree shape is preserved.
func CreateIntentTree(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	opts ...IntentConfigOption,
) (*v3.WalletConfigTree, error) {
	options := applyIntentConfigOptions(opts)
	return createIntentTree(mainSigner, calls, options.gateLeafNode, options.sapientSignerLeafNode)
}

func createIntentConfiguration(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	checkpoint uint64,
	gateLeafNode v3.WalletConfigTree,
	sapientSignerLeafNode v3.WalletConfigTree,
) (*v3.WalletConfig, error) {
	tree, err := createIntentTree(mainSigner, calls, gateLeafNode, sapientSignerLeafNode)
	if err != nil {
		return nil, err
	}

	return &v3.WalletConfig{
		Threshold_:  1,
		Checkpoint_: checkpoint,
		Tree:        *tree,
	}, nil
}

// `CreateIntentConfiguration` creates a wallet configuration where the intent's transaction
// batches are grouped into the initial subdigest. See WithGate and WithSapientSigner
// for the optional leaves.
func CreateIntentConfiguration(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	checkpoint uint64,
	opts ...IntentConfigOption,
) (*v3.WalletConfig, error) {
	options := applyIntentConfigOptions(opts)
	return createIntentConfiguration(mainSigner, calls, checkpoint, options.gateLeafNode, options.sapientSignerLeafNode)
}

// `BuildIntentConfigurationSignature` creates a signature for an already-built intent configuration
// that can be used to bypass chain ID validation. All supplied signer signatures are
// embedded deterministically; signers without a supplied signature are encoded as their
// image hash.
func BuildIntentConfigurationSignature(config *v3.WalletConfig, signerSignatures []*core.SignerSignature) ([]byte, error) {
	if config == nil {
		return nil, fmt.Errorf("intent configuration is nil")
	}

	signatures := make(map[core.Signer]core.SignerSignature, len(signerSignatures))
	for _, signerSignature := range signerSignatures {
		if signerSignature != nil {
			signatures[signerSignature.Signer] = *signerSignature
		}
	}

	sig := config.BuildRegularSignatureFromSignatures(signatures)

	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	// Print the signature data
	// fmt.Printf("signature data: %s\n", common.Bytes2Hex(data))

	return data, nil
}

// `GetIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation.
// The signature is based on the transaction bundle digests only.
func GetIntentConfigurationSignature(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	checkpoint uint64,
	signerSignatures []*core.SignerSignature,
	opts ...IntentConfigOption,
) ([]byte, error) {
	options := applyIntentConfigOptions(opts)
	config, err := createIntentConfiguration(mainSigner, calls, checkpoint, options.gateLeafNode, options.sapientSignerLeafNode)
	if err != nil {
		return nil, err
	}

	return BuildIntentConfigurationSignature(config, signerSignatures)
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
