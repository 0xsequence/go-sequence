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

// AnypayLifiInfo represents the information for a Lifi bridge or swap.
type AnypayLifiInfo struct {
	OriginToken        common.Address `abi:"originToken"`
	MinAmount          *big.Int       `abi:"minAmount"`
	OriginChainId      *big.Int       `abi:"originChainId"`
	DestinationChainId *big.Int       `abi:"destinationChainId"`
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
		// Change the address to address(0)
		callPayload.AddressZero()

		individualPayloadDigest := callPayload.Digest()

		packedForKeccak := append(cumulativeCallsHash[:], individualPayloadDigest.Hash[:]...)
		fmt.Printf("    packedForKeccak: 0x%s\n", common.Bytes2Hex(packedForKeccak))

		cumulativeCallsHash = ethcoder.Keccak256Hash(packedForKeccak)
		fmt.Printf("    cumulativeCallsHash after callPayload[%d]: 0x%s\n", i, common.Bytes2Hex(cumulativeCallsHash[:]))
	}
	fmt.Printf("  final cumulativeCallsHash: 0x%s\n", common.Bytes2Hex(cumulativeCallsHash[:]))
	fmt.Printf("  destinationTokens: %v\n", params.DestinationTokens)

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
	fmt.Printf("    encoded: 0x%s\n", common.Bytes2Hex(encoded))

	hash := ethcoder.Keccak256(encoded)

	var hash32 [32]byte
	copy(hash32[:], hash)
	return hash32, nil
}

// GetAnypayLifiInfoHash computes the Keccak256 hash of ABI-encoded AnypayLifiInfo array and an attestation address.
func GetAnypayLifiInfoHash(lifiInfos []AnypayLifiInfo, attestationAddress common.Address) ([32]byte, error) {
	// Define ABI type components for the AnypayLifiInfo struct
	anypayLifiInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "minAmount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}

	// Define ABI type for a list of AnypayLifiInfo structs (AnypayLifiInfo[])
	anypayLifiInfoListType, err := abi.NewType("tuple[]", "AnypayLifiInfo[]", anypayLifiInfoComponents)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create anypay lifi info list ABI type: %w", err)
	}

	// Define ABI type for address
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	// Define the arguments for ABI encoding
	arguments := abi.Arguments{
		{Name: "lifiInfos", Type: anypayLifiInfoListType},
		{Name: "attestationAddress", Type: addressType},
	}

	// ABI encode the arguments
	encodedData, err := arguments.Pack(lifiInfos, attestationAddress)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to ABI pack arguments for AnypayLifiInfo hash: %w", err)
	}

	// Compute Keccak256 hash
	hash := ethcoder.Keccak256(encodedData)

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
		leaf := &v3.WalletConfigTreeAnyAddressSubdigestLeaf{
			Digest: core.Subdigest{Hash: digest.Hash},
		}
		leaves = append(leaves, leaf)
	}

	return leaves, nil
}

// `CreateAnypaySapientSignerTree` creates a tree from a list of AnypayLifiInfo and a main signer address.
func CreateAnypaySapientSignerTree(mainSigner common.Address, lifiInfos []AnypayLifiInfo) (v3.WalletConfigTree, error) {
	// Get the image hash for the main signer.
	sapientImageHash, err := GetAnypayLifiInfoHash(lifiInfos, mainSigner)
	if err != nil {
		return nil, fmt.Errorf("failed to get image hash for main signer: %w", err)
	}

	// Create the lifi info leaf.
	sapientSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		ImageHash_: core.ImageHash{Hash: common.BytesToHash(sapientImageHash[:])},
	}

	return sapientSignerLeaf, nil
}

// `CreateIntentTree` creates a tree from a list of intent operations and a main signer address.
func CreateIntentTree(mainSigner common.Address, calls []*v3.CallsPayload, lifiInfos ...AnypayLifiInfo) (*v3.WalletConfigTree, error) {
	var sapientSignerLeaf v3.WalletConfigTree
	var err error

	if len(lifiInfos) > 0 {
		// Create the lifi info leaf.
		sapientSignerLeaf, err = CreateAnypaySapientSignerTree(mainSigner, lifiInfos)
		if err != nil {
			return nil, fmt.Errorf("failed to create lifi info leaf: %w", err)
		}
	}

	// Create the subdigest leaves from the batched transactions.
	leaves, err := CreateAnyAddressSubdigestTree(calls)
	if err != nil {
		return nil, err
	}

	// If the sapient signer leaf is not nil, add it to the leaves.
	if sapientSignerLeaf != nil {
		leaves = append(leaves, sapientSignerLeaf)
	}

	// If the length of the leaves is 1
	if len(leaves) == 1 {
		tree := v3.WalletConfigTree(leaves[0])
		return &tree, nil
	}

	// Create a tree from the subdigest leaves.
	tree := v3.WalletConfigTreeNodes(leaves...)

	// Create the main signer leaf (with weight 1).
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Construct the new wallet config using:
	fullTree := v3.WalletConfigTreeNodes(mainSignerLeaf, tree)

	return &fullTree, nil
}

// `CreateIntentConfiguration` creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateIntentConfiguration(mainSigner common.Address, calls []*v3.CallsPayload, lifiInfos ...AnypayLifiInfo) (*v3.WalletConfig, error) {
	// Create the subdigest leaves from the batched transactions.
	tree, err := CreateIntentTree(mainSigner, calls, lifiInfos...)
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
func GetIntentConfigurationSignature(mainSigner common.Address, calls []*v3.CallsPayload, lifiInfos ...AnypayLifiInfo) ([]byte, error) {
	// Create the intent configuration using the batched transactions.
	config, err := CreateIntentConfiguration(mainSigner, calls, lifiInfos...)
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
