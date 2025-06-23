package sequence

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
)

var (
	AnypayLiFiSapientSignerAddress      = common.HexToAddress("0xd7571bd1e3af468c3a49966c9a92a2e907cdfa52")
	AnypayLifiSapientSignerLiteAddress  = common.HexToAddress("0xaA3f6B332237aFb83789d3F5FBaD817EF3102648")
	AnypayRelaySapientSignerAddress     = common.HexToAddress("0xA93EC61Df992ED3F3c30Ca107Cb09DCB0aCFe838")
	AnypayRelaySapientSignerLiteAddress = common.HexToAddress("0xA93EC61Df992ED3F3c30Ca107Cb09DCB0aCFe838")
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

// AnypayExecutionInfo represents the information for a Lifi bridge or swap.
type AnypayExecutionInfo struct {
	OriginToken        common.Address `abi:"originToken"`
	Amount             *big.Int       `abi:"amount"`
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

// GetAnypayExecutionInfoHash computes the Keccak256 hash of ABI-encoded AnypayExecutionInfo array and an attestation address.
func GetAnypayExecutionInfoHash(lifiInfos []AnypayExecutionInfo, attestationAddress common.Address) ([32]byte, error) {
	// Define ABI type components for the AnypayExecutionInfo struct
	AnypayExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}

	// Define ABI type for a list of AnypayExecutionInfo structs (AnypayExecutionInfo[])
	AnypayExecutionInfoListType, err := abi.NewType("tuple[]", "AnypayExecutionInfo[]", AnypayExecutionInfoComponents)
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
		{Name: "lifiInfos", Type: AnypayExecutionInfoListType},
		{Name: "attestationAddress", Type: addressType},
	}

	// ABI encode the arguments
	encodedData, err := arguments.Pack(lifiInfos, attestationAddress)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to ABI pack arguments for AnypayExecutionInfo hash: %w", err)
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

// `CreateAnypayExecutionInfoSapientSignerTree` creates a tree from a list of AnypayExecutionInfo and a main signer address.
func CreateAnypayExecutionInfoSapientSignerTree(attestationSigner common.Address, anypayExecutionInfos []AnypayExecutionInfo, sapientType string) (v3.WalletConfigTree, error) {
	// Get the image hash for the main signer.
	sapientImageHash, err := GetAnypayExecutionInfoHash(anypayExecutionInfos, attestationSigner)
	if err != nil {
		return nil, fmt.Errorf("failed to get image hash for main signer: %w", err)
	}
	fmt.Printf("sapientImageHash: %s\n", common.Bytes2Hex(sapientImageHash[:]))

	var sapientSignerAddress common.Address
	switch sapientType {
	case "lifi":
		sapientSignerAddress = AnypayLifiSapientSignerLiteAddress
	case "relay":
		sapientSignerAddress = AnypayRelaySapientSignerAddress
	}

	// Create the lifi info leaf.
	sapientSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Address:    sapientSignerAddress,
		Weight:     1,
		ImageHash_: core.ImageHash{Hash: common.BytesToHash(sapientImageHash[:])},
	}
	fmt.Printf("sapientSignerLeaf.ImageHash(): %s\n", sapientSignerLeaf.ImageHash().Hash.Hex())

	return sapientSignerLeaf, nil
}

// `CreateIntentTree` creates a tree from a list of intent operations and a main signer address.
func CreateIntentTree(mainSigner common.Address, calls []*v3.CallsPayload, sapientSignerLeafNode v3.WalletConfigTree) (*v3.WalletConfigTree, error) {
	// Create the subdigest leaves from the batched transactions.
	leaves, err := CreateAnyAddressSubdigestTree(calls)
	if err != nil {
		return nil, err
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

// `CreateLifiIntentConfiguration` is a helper function to create a LiFi intent configuration.
func CreateLifiIntentConfiguration(mainSigner, attestationSigner common.Address, calls []*v3.CallsPayload, anypayExecutionInfos []AnypayExecutionInfo) (*v3.WalletConfig, error) {
	var sapientSignerLeafNode v3.WalletConfigTree
	var err error

	if attestationSigner != (common.Address{}) && len(anypayExecutionInfos) > 0 {
		sapientSignerLeafNode, err = CreateAnypayExecutionInfoSapientSignerTree(attestationSigner, anypayExecutionInfos, "lifi")
		if err != nil {
			return nil, fmt.Errorf("failed to create lifi info leaf: %w", err)
		}
	}

	return CreateIntentConfiguration(mainSigner, calls, sapientSignerLeafNode)
}

// `CreateRelayIntentConfiguration` is a helper function to create a relay intent configuration.
func CreateRelayIntentConfiguration(mainSigner, attestationSigner common.Address, calls []*v3.CallsPayload, anypayExecutionInfos []AnypayExecutionInfo) (*v3.WalletConfig, error) {
	var sapientSignerLeafNode v3.WalletConfigTree
	var err error

	if attestationSigner != (common.Address{}) && len(anypayExecutionInfos) > 0 {
		sapientSignerLeafNode, err = CreateAnypayExecutionInfoSapientSignerTree(attestationSigner, anypayExecutionInfos, "relay")
		if err != nil {
			return nil, fmt.Errorf("failed to create relay info leaf: %w", err)
		}
	}

	return CreateIntentConfiguration(mainSigner, calls, sapientSignerLeafNode)
}

// `GetIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation.
// The signature is based on the transaction bundle digests only.
func GetIntentConfigurationSignature(
	mainSigner common.Address,
	attestationSigner common.Address,
	calls []*v3.CallsPayload,
	attestationSignerWallet *ethwallet.Wallet,
	targetPayload *v3.CallsPayload,
	sapientType string, // "lifi" or "relay"
	anypayExecutionInfos []AnypayExecutionInfo,
) ([]byte, error) {
	var config *v3.WalletConfig
	var err error

	switch sapientType {
	case "lifi":
		config, err = CreateLifiIntentConfiguration(mainSigner, attestationSigner, calls, anypayExecutionInfos)
		if err != nil {
			return nil, err
		}
	case "relay":
		config, err = CreateRelayIntentConfiguration(mainSigner, attestationSigner, calls, anypayExecutionInfos)
		if err != nil {
			return nil, err
		}
	default:
		// Default case without any sapient signer
		config, err = CreateIntentConfiguration(mainSigner, calls, nil)
		if err != nil {
			return nil, err
		}
	}

	// If the target payload is nil, we need to replace the sapient signer with a node in the config tree.
	// This is because the sapient signer is not included in the target payload, so we need to replace it with a node, since we cannot call the sapient signer directly, but rather the node.
	if targetPayload == nil {
		config.Tree = replaceSapientSignerWithNodeInConfigTree(config.Tree)
	}

	signingFunc := func(ctx context.Context, signer common.Address, _ []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		fmt.Printf("signingFunc: signer: %s\n", signer.Hex())

		if signer == AnypayLifiSapientSignerLiteAddress && len(anypayExecutionInfos) > 0 && targetPayload != nil {
			fmt.Printf("matched AnypayLifiSapientSignerLiteAddress\n")
			fmt.Printf("signingFunc: anypayExecutionInfos: %v\n", anypayExecutionInfos)
			var attestationBytes []byte
			attestationBytes, err = CreateAnypayExecutionInfoAttestationLite(anypayExecutionInfos)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to create attestation: %w", err)
			}
			return core.SignerSignatureTypeSapient, attestationBytes, nil
		}

		if signer == AnypayRelaySapientSignerAddress && len(anypayExecutionInfos) > 0 && targetPayload != nil {
			fmt.Printf("matched AnypayRelaySapientSignerAddress\n")
			var attestationBytes []byte
			attestationBytes, err = CreateAnypayExecutionInfoAttestationLite(anypayExecutionInfos)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to create relay attestation: %w", err)
			}
			return core.SignerSignatureTypeSapient, attestationBytes, nil
		}

		fmt.Printf("signingFunc: returning nil signature for signer: %s\n", signer.Hex())
		// For mainSigner or other signers, we don't provide a signature here.
		// This will result in an AddressLeaf or NodeLeaf in the signature tree.
		return 0, nil, nil
	}

	spew.Dump(config)
	spew.Dump(config.Tree)

	// Build the signature using BuildNoChainIDSignature, which allows us to inject custom signatures via SigningFunction.
	// Set validateSigningPower to false, as we are not necessarily providing signatures for all parts of the config.
	sig, err := config.BuildRegularSignature(context.Background(), signingFunc, false)
	if err != nil {
		return nil, fmt.Errorf("failed to build regular signature: %w", err)
	}

	spew.Dump(sig)

	if regularSig, ok := sig.(*v3.RegularSignature); ok {
		if regularSig.Signature != nil {
			signatureTree := regularSig.Signature.Tree
			fmt.Println("Accessing sig.Signature.Tree:")
			spew.Dump(signatureTree)
		} else {
			fmt.Println("sig.Signature is nil")
		}
	} else if noChainIdSig, ok := sig.(*v3.NoChainIDSignature); ok {
		if noChainIdSig.Signature != nil {
			signatureTree := noChainIdSig.Signature.Tree
			fmt.Println("Accessing sig.Signature.Tree (NoChainID):")
			spew.Dump(signatureTree)
		} else {
			fmt.Println("sig.Signature is nil for NoChainIDSignature")
		}
	} else {
		fmt.Printf("sig is not of type *v3.RegularSignature or *v3.NoChainIDSignature, it is %T\n", sig)
	}

	// Get the signature data
	data, err := sig.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to get signature data: %w", err)
	}

	// Print the signature data
	fmt.Printf("signature data: %s\n", common.Bytes2Hex(data))

	return data, nil
}

// CreateAnypayLifiAttestation generates the ABI-encoded signature required by the AnypayLifiSapientSigner contract.
func CreateAnypayLifiAttestation(
	attestationSignerWallet *ethwallet.Wallet,
	payload *v3.CallsPayload,
	lifiInfos []AnypayExecutionInfo,
) ([]byte, error) {
	if attestationSignerWallet == nil {
		return nil, fmt.Errorf("attestationSignerWallet is nil")
	}

	if payload == nil {
		return nil, fmt.Errorf("payload is nil for attestation")
	}

	if len(lifiInfos) == 0 {
		return nil, fmt.Errorf("lifiInfos is empty")
	}

	digestToSign := payload.Digest()
	fmt.Printf("digestToSign: %s\n", common.Bytes2Hex(digestToSign.Hash[:]))

	rawSignature, err := attestationSignerWallet.SignData(digestToSign.Hash[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign payload digest: %w", err)
	}

	// Adjust V: Ethereum expects V to be 27 or 28. crypto.Sign returns 0 or 1.
	// The last byte of the 65-byte signature is V.
	if len(rawSignature) == 65 {
		rawSignature[64] = rawSignature[64] + 27
	} else {
		return nil, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(rawSignature))
	}
	eoaSignatureBytes := rawSignature

	// 4. Define ABI types for abi.encode(AnypayExecutionInfo[] memory, bytes memory)
	AnypayExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	lifiInfoArrayType, err := abi.NewType("tuple[]", "AnypayExecutionInfo[]", AnypayExecutionInfoComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create AnypayExecutionInfo[] ABI type: %w", err)
	}
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}

	// 5. Pack lifiInfos and eoaSignatureBytes
	encodedAttestation, err := abi.Arguments{{Type: lifiInfoArrayType}, {Type: bytesType}}.Pack(lifiInfos, eoaSignatureBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to ABI pack AnypayExecutionInfo[] and eoaSignature: %w", err)
	}

	return encodedAttestation, nil
}

func CreateAnypayExecutionInfoAttestationLite(
	lifiInfos []AnypayExecutionInfo,
) ([]byte, error) {
	// 4. Define ABI types for abi.encode(AnypayExecutionInfo[] memory, bytes memory, address)
	AnypayExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	lifiInfoArrayType, err := abi.NewType("tuple[]", "AnypayExecutionInfo[]", AnypayExecutionInfoComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create AnypayExecutionInfo[] ABI type: %w", err)
	}
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	// Create a random digest to sign
	randomDigest := make([]byte, 32)
	if _, err := rand.Read(randomDigest); err != nil {
		return nil, fmt.Errorf("failed to generate random digest: %w", err)
	}

	// Create a random EOA with a random private key
	randomPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate random private key: %w", err)
	}
	randomEoa := crypto.PubkeyToAddress(randomPrivateKey.PublicKey)

	// Generate a random signature for the EOA
	randomSignature, err := crypto.Sign(randomDigest, randomPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign payload digest: %w", err)
	}

	// 5. Pack lifiInfos and eoaSignatureBytes
	encodedAttestation, err := abi.Arguments{{Type: lifiInfoArrayType}, {Type: bytesType}, {Type: addressType}}.Pack(lifiInfos, randomSignature, randomEoa)
	if err != nil {
		return nil, fmt.Errorf("failed to ABI pack AnypayExecutionInfo[] and eoaSignature: %w", err)
	}

	return encodedAttestation, nil
}

// replaceSapientSignerWithNodeInConfigTree recursively traverses the WalletConfigTree.
func replaceSapientSignerWithNodeInConfigTree(tree v3.WalletConfigTree) v3.WalletConfigTree {
	if tree == nil {
		return nil
	}

	switch node := tree.(type) {
	case *v3.WalletConfigTreeNode:
		// Recursively call on left and right children
		left := replaceSapientSignerWithNodeInConfigTree(node.Left)
		right := replaceSapientSignerWithNodeInConfigTree(node.Right)

		if left == node.Left && right == node.Right {
			return node
		}
		return &v3.WalletConfigTreeNode{Left: left, Right: right}

	case *v3.WalletConfigTreeNestedLeaf:
		// Recursively call on the inner tree
		innerTree := replaceSapientSignerWithNodeInConfigTree(node.Tree)

		if innerTree == node.Tree { // Check for pointer equality
			return node // No change, return original
		}
		return &v3.WalletConfigTreeNestedLeaf{
			Weight:    node.Weight,
			Threshold: node.Threshold,
			Tree:      innerTree,
		}

	case *v3.WalletConfigTreeSapientSignerLeaf:
		// This is the target node type to replace
		return &v3.WalletConfigTreeNodeLeaf{Node: node.ImageHash()}

	default:
		return tree
	}
}
