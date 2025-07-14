package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
)

var (
	TrailsLiFiSapientSignerAddress   = common.HexToAddress("0xd7571bd1e3af468c3a49966c9a92a2e907cdfa52")
	TrailsRelaySapientSignerAddress  = common.HexToAddress("0x9a013E7D186611aF36A918eF23D81886E8C256F8")
	TrailsCCTPV2SapientSignerAddress = common.HexToAddress("0xc1A9B197eBb31Fc2B613C59dAC3f3E5698A429D0")
)

// AddressOverrides provides configurable address overrides for skewness protection
type AddressOverrides struct {
	TrailsLiFiSapientSignerAddress   *common.Address
	TrailsRelaySapientSignerAddress  *common.Address
	TrailsCCTPV2SapientSignerAddress *common.Address
}

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

// TrailsExecutionInfo represents the information for a Lifi bridge or swap.
type TrailsExecutionInfo struct {
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
		callPayload, _ := v3.PayloadWithAddress(callPayload, common.Address{})

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

// GetTrailsExecutionInfoHash computes the Keccak256 hash of ABI-encoded TrailsExecutionInfo array and an attestation address.
func GetTrailsExecutionInfoHash(executionInfos []TrailsExecutionInfo, attestationAddress common.Address) ([32]byte, error) {
	if len(executionInfos) == 0 {
		return [32]byte{}, fmt.Errorf("executionInfos is empty")
	}

	spew.Dump(executionInfos)
	spew.Dump(attestationAddress)

	// Define ABI type components for the TrailsExecutionInfo struct
	TrailsExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}

	// Define ABI type for a list of TrailsExecutionInfo structs (TrailsExecutionInfo[])
	TrailsExecutionInfoListType, err := abi.NewType("tuple[]", "TrailsExecutionInfo[]", TrailsExecutionInfoComponents)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create trails lifi info list ABI type: %w", err)
	}

	// Define ABI type for address
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	// Define the arguments for ABI encoding
	arguments := abi.Arguments{
		{Name: "executionInfos", Type: TrailsExecutionInfoListType},
		{Name: "attestationAddress", Type: addressType},
	}

	// ABI encode the arguments
	encodedData, err := arguments.Pack(executionInfos, attestationAddress)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to ABI pack arguments for TrailsExecutionInfo hash: %w", err)
	}

	// Compute Keccak256 hash
	hash := ethcoder.Keccak256(encodedData)

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

// `CreateTrailsExecutionInfoSapientSignerTree` creates a tree from a list of TrailsExecutionInfo and a main signer address.
func CreateTrailsExecutionInfoSapientSignerTree(attestationSigner common.Address, trailsExecutionInfos []TrailsExecutionInfo, sapientType string, overrides ...*AddressOverrides) (v3.WalletConfigTree, error) {
	// Get the image hash for the main signer.
	sapientImageHash, err := GetTrailsExecutionInfoHash(trailsExecutionInfos, attestationSigner)
	if err != nil {
		return nil, fmt.Errorf("failed to get image hash for main signer: %w", err)
	}
	fmt.Printf("sapientImageHash: %s\n", common.Bytes2Hex(sapientImageHash[:]))

	var override *AddressOverrides
	if len(overrides) > 0 && overrides[0] != nil {
		override = overrides[0]
	}

	var sapientSignerAddress common.Address
	switch sapientType {
	case "lifi":
		sapientSignerAddress = TrailsLiFiSapientSignerAddress
		if override != nil && override.TrailsLiFiSapientSignerAddress != nil {
			sapientSignerAddress = *override.TrailsLiFiSapientSignerAddress
		}
	case "relay":
		sapientSignerAddress = TrailsRelaySapientSignerAddress
		if override != nil && override.TrailsRelaySapientSignerAddress != nil {
			sapientSignerAddress = *override.TrailsRelaySapientSignerAddress
		}
	case "cctpv2":
		sapientSignerAddress = TrailsCCTPV2SapientSignerAddress
		if override != nil && override.TrailsCCTPV2SapientSignerAddress != nil {
			sapientSignerAddress = *override.TrailsCCTPV2SapientSignerAddress
		}
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
func CreateLifiIntentConfiguration(mainSigner, attestationSigner common.Address, calls []*v3.CallsPayload, trailsExecutionInfos []TrailsExecutionInfo, overrides ...*AddressOverrides) (*v3.WalletConfig, error) {
	var sapientSignerLeafNode v3.WalletConfigTree
	var err error

	if attestationSigner != (common.Address{}) && len(trailsExecutionInfos) > 0 {
		sapientSignerLeafNode, err = CreateTrailsExecutionInfoSapientSignerTree(attestationSigner, trailsExecutionInfos, "lifi", overrides...)
		if err != nil {
			return nil, fmt.Errorf("failed to create lifi info leaf: %w", err)
		}
	}

	return CreateIntentConfiguration(mainSigner, calls, sapientSignerLeafNode)
}

// `CreateRelayIntentConfiguration` is a helper function to create a relay intent configuration.
func CreateRelayIntentConfiguration(mainSigner, attestationSigner common.Address, calls []*v3.CallsPayload, trailsExecutionInfos []TrailsExecutionInfo, overrides ...*AddressOverrides) (*v3.WalletConfig, error) {
	var sapientSignerLeafNode v3.WalletConfigTree
	var err error

	if attestationSigner != (common.Address{}) && len(trailsExecutionInfos) > 0 {
		sapientSignerLeafNode, err = CreateTrailsExecutionInfoSapientSignerTree(attestationSigner, trailsExecutionInfos, "relay", overrides...)
		if err != nil {
			return nil, fmt.Errorf("failed to create relay info leaf: %w", err)
		}
	}

	return CreateIntentConfiguration(mainSigner, calls, sapientSignerLeafNode)
}

// `CreateCCTPV2IntentConfiguration` is a helper function to create a CCTP V2 intent configuration.
func CreateCCTPV2IntentConfiguration(mainSigner, attestationSigner common.Address, calls []*v3.CallsPayload, trailsExecutionInfos []TrailsExecutionInfo, overrides ...*AddressOverrides) (*v3.WalletConfig, error) {
	var sapientSignerLeafNode v3.WalletConfigTree
	var err error

	if attestationSigner != (common.Address{}) && len(trailsExecutionInfos) > 0 {
		sapientSignerLeafNode, err = CreateTrailsExecutionInfoSapientSignerTree(attestationSigner, trailsExecutionInfos, "cctpv2", overrides...)
		if err != nil {
			return nil, fmt.Errorf("failed to create cctp info leaf: %w", err)
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
	sapientType string, // "lifi", "relay" or "cctp"
	trailsExecutionInfos []TrailsExecutionInfo,
	decodingStrategy *uint8,
	overrides ...*AddressOverrides,
) ([]byte, error) {
	var config *v3.WalletConfig
	var err error

	if sapientType == "lifi" {
		if decodingStrategy == nil {
			return nil, fmt.Errorf("decodingStrategy is required for lifi sapient type")
		}
	} else if decodingStrategy != nil {
		return nil, fmt.Errorf("decodingStrategy is only supported for lifi sapient type, but got %s", sapientType)
	}

	switch sapientType {
	case "lifi":
		config, err = CreateLifiIntentConfiguration(mainSigner, attestationSigner, calls, trailsExecutionInfos, overrides...)
		if err != nil {
			return nil, err
		}
	case "relay":
		config, err = CreateRelayIntentConfiguration(mainSigner, attestationSigner, calls, trailsExecutionInfos, overrides...)
		if err != nil {
			return nil, err
		}
	case "cctp":
		config, err = CreateCCTPV2IntentConfiguration(mainSigner, attestationSigner, calls, trailsExecutionInfos, overrides...)
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

	var override *AddressOverrides
	if len(overrides) > 0 && overrides[0] != nil {
		override = overrides[0]
	}

	trailsLiFiSapientSignerAddress := TrailsLiFiSapientSignerAddress
	if override != nil && override.TrailsLiFiSapientSignerAddress != nil {
		trailsLiFiSapientSignerAddress = *override.TrailsLiFiSapientSignerAddress
	}

	trailsRelaySapientSignerAddress := TrailsRelaySapientSignerAddress
	if override != nil && override.TrailsRelaySapientSignerAddress != nil {
		trailsRelaySapientSignerAddress = *override.TrailsRelaySapientSignerAddress
	}

	trailsCCTPV2SapientSignerAddress := TrailsCCTPV2SapientSignerAddress
	if override != nil && override.TrailsCCTPV2SapientSignerAddress != nil {
		trailsCCTPV2SapientSignerAddress = *override.TrailsCCTPV2SapientSignerAddress
	}

	signingFunc := func(ctx context.Context, signer common.Address, _ []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		fmt.Printf("signingFunc: signer: %s\n", signer.Hex())

		if signer == trailsLiFiSapientSignerAddress && len(trailsExecutionInfos) > 0 && targetPayload != nil {
			fmt.Printf("matched TrailsLifiSapientSignerAddress\n")
			fmt.Printf("signingFunc: trailsExecutionInfos: %v\n", trailsExecutionInfos)
			var attestationBytes []byte
			attestationBytes, err = CreateTrailsLifiAttestation(attestationSignerWallet, targetPayload, trailsExecutionInfos, *decodingStrategy)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to create attestation: %w", err)
			}
			return core.SignerSignatureTypeSapient, attestationBytes, nil
		}

		if signer == trailsRelaySapientSignerAddress && len(trailsExecutionInfos) > 0 && targetPayload != nil {
			fmt.Printf("matched TrailsRelaySapientSignerAddress\n")
			var attestationBytes []byte
			attestationBytes, err = CreateTrailsRelayAttestation(attestationSignerWallet, targetPayload, trailsExecutionInfos)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to create relay attestation: %w", err)
			}
			return core.SignerSignatureTypeSapient, attestationBytes, nil
		}

		if signer == trailsCCTPV2SapientSignerAddress && len(trailsExecutionInfos) > 0 && targetPayload != nil {
			fmt.Printf("matched TrailsCCTPV2SapientSignerAddress\n")
			var attestationBytes []byte
			attestationBytes, err = CreateTrailsCCTPAttestation(attestationSignerWallet, targetPayload, trailsExecutionInfos)
			if err != nil {
				return 0, nil, fmt.Errorf("failed to create cctp attestation: %w", err)
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

// CreateTrailsLifiAttestation generates the ABI-encoded signature required by the TrailsLifiSapientSigner contract.
func CreateTrailsLifiAttestation(
	attestationSignerWallet *ethwallet.Wallet,
	payload *v3.CallsPayload,
	executionInfos []TrailsExecutionInfo,
	decodingStrategy uint8,
) ([]byte, error) {
	if attestationSignerWallet == nil {
		return nil, fmt.Errorf("attestationSignerWallet is nil")
	}

	if payload == nil {
		return nil, fmt.Errorf("payload is nil for attestation")
	}

	if len(executionInfos) == 0 {
		return nil, fmt.Errorf("executionInfos is empty")
	}

	digestToSign := payload.Digest()
	fmt.Printf("CreateTrailsLifiAttestation: digestToSign: %s\n", common.Bytes2Hex(digestToSign.Hash[:]))
	rawSignature, err := attestationSignerWallet.SignMessage(digestToSign.Hash[:])
	fmt.Printf("CreateTrailsLifiAttestation: rawSignature: %s\n", common.Bytes2Hex(rawSignature))

	if err != nil {
		return nil, fmt.Errorf("failed to sign payload digest: %w", err)
	}

	// The signature V value is already adjusted by ethwallet.SignData.
	// We just need to ensure the signature length is correct.
	if len(rawSignature) != 65 {
		return nil, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(rawSignature))
	}
	eoaSignatureBytes := rawSignature

	// 4. Define ABI types for abi.encode(TrailsExecutionInfo[] memory, bytes memory)
	TrailsExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	lifiInfoArrayType, err := abi.NewType("tuple[]", "TrailsExecutionInfo[]", TrailsExecutionInfoComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create TrailsExecutionInfo[] ABI type: %w", err)
	}
	uint8Type, err := abi.NewType("uint8", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint8 ABI type: %w", err)
	}
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	fmt.Printf("CreateTrailsLifiAttestation: attestationSignerWallet.Address(): %s\n", attestationSignerWallet.Address().Hex())

	// 5. Pack executionInfos and eoaSignatureBytes
	encodedAttestation, err := abi.Arguments{{Type: lifiInfoArrayType}, {Type: uint8Type}, {Type: bytesType}, {Type: addressType}}.Pack(executionInfos, decodingStrategy, eoaSignatureBytes, attestationSignerWallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to ABI pack TrailsExecutionInfo[] and eoaSignature: %w", err)
	}

	return encodedAttestation, nil
}

// CreateTrailsRelayAttestation generates the ABI-encoded signature required by the TrailsRelaySapientSigner contract.
func CreateTrailsRelayAttestation(
	attestationSignerWallet *ethwallet.Wallet,
	payload *v3.CallsPayload,
	trailsExecutionInfos []TrailsExecutionInfo,
) ([]byte, error) {
	if attestationSignerWallet == nil {
		return nil, fmt.Errorf("attestationSignerWallet is nil")
	}
	if payload == nil {
		return nil, fmt.Errorf("payload is nil for attestation")
	}
	if len(trailsExecutionInfos) == 0 {
		return nil, fmt.Errorf("trailsExecutionInfos is empty")
	}

	digestToSign := payload.Digest()
	fmt.Printf("CreateTrailsRelayAttestation: digestToSign: %s\n", common.Bytes2Hex(digestToSign.Hash[:]))
	rawSignature, err := attestationSignerWallet.SignMessage(digestToSign.Hash[:])
	fmt.Printf("CreateTrailsRelayAttestation: rawSignature: %s\n", common.Bytes2Hex(rawSignature))

	if err != nil {
		return nil, fmt.Errorf("failed to sign payload digest: %w", err)
	}

	if len(rawSignature) != 65 {
		return nil, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(rawSignature))
	}
	eoaSignatureBytes := rawSignature

	// Define ABI types for abi.encode(TrailsExecutionInfo[] memory, bytes memory, address)
	TrailsExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	trailsInfoArrayType, err := abi.NewType("tuple[]", "TrailsExecutionInfo[]", TrailsExecutionInfoComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create TrailsExecutionInfo[] ABI type: %w", err)
	}
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	fmt.Printf("CreateTrailsRelayAttestation: attestationSignerWallet.Address(): %s\n", attestationSignerWallet.Address().Hex())

	// Pack trailsExecutionInfos, eoaSignatureBytes, and the signer's address
	encodedAttestation, err := abi.Arguments{{Type: trailsInfoArrayType}, {Type: bytesType}, {Type: addressType}}.Pack(trailsExecutionInfos, eoaSignatureBytes, attestationSignerWallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to ABI pack TrailsExecutionInfo[], signature and address: %w", err)
	}

	return encodedAttestation, nil
}

// CreateTrailsCCTPAttestation generates the ABI-encoded signature required by the TrailsCCTPV2SapientSigner contract.
func CreateTrailsCCTPAttestation(
	attestationSignerWallet *ethwallet.Wallet,
	payload *v3.CallsPayload,
	trailsExecutionInfos []TrailsExecutionInfo,
) ([]byte, error) {
	if attestationSignerWallet == nil {
		return nil, fmt.Errorf("attestationSignerWallet is nil")
	}
	if payload == nil {
		return nil, fmt.Errorf("payload is nil for attestation")
	}
	if len(trailsExecutionInfos) == 0 {
		return nil, fmt.Errorf("trailsExecutionInfos is empty")
	}

	digestToSign := payload.Digest()
	fmt.Printf("CreateTrailsCCTPAttestation: digestToSign: %s\n", common.Bytes2Hex(digestToSign.Hash[:]))
	rawSignature, err := attestationSignerWallet.SignMessage(digestToSign.Hash[:])
	fmt.Printf("CreateTrailsCCTPAttestation: rawSignature: %s\n", common.Bytes2Hex(rawSignature))

	if err != nil {
		return nil, fmt.Errorf("failed to sign payload digest: %w", err)
	}

	if len(rawSignature) != 65 {
		return nil, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(rawSignature))
	}
	eoaSignatureBytes := rawSignature

	// Define ABI types for abi.encode(TrailsExecutionInfo[] memory, bytes memory, address)
	TrailsExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	trailsInfoArrayType, err := abi.NewType("tuple[]", "TrailsExecutionInfo[]", TrailsExecutionInfoComponents)
	if err != nil {
		return nil, fmt.Errorf("failed to create TrailsExecutionInfo[] ABI type: %w", err)
	}
	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}
	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create address ABI type: %w", err)
	}

	fmt.Printf("CreateTrailsCCTPAttestation: attestationSignerWallet.Address(): %s\n", attestationSignerWallet.Address().Hex())

	// Pack trailsExecutionInfos, eoaSignatureBytes, and the signer's address
	encodedAttestation, err := abi.Arguments{{Type: trailsInfoArrayType}, {Type: bytesType}, {Type: addressType}}.Pack(trailsExecutionInfos, eoaSignatureBytes, attestationSignerWallet.Address())
	if err != nil {
		return nil, fmt.Errorf("failed to ABI pack TrailsExecutionInfo[], signature and address: %w", err)
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
