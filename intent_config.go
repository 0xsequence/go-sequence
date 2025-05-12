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

// `CreateInitialIntentConfiguration` creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateInitialIntentConfiguration(mainSigner common.Address, authSigner common.Address) (*v3.WalletConfig, error) {
	// Create a 1/2 call payload with the main signer and auth signer addresses.

	// Create the main signer leaf with weight 1.
	mainSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: mainSigner,
	}

	// Create the auth signer leaf with weight 1.
	authSignerLeaf := &v3.WalletConfigTreeAddressLeaf{
		Weight:  1,
		Address: authSigner,
	}

	// Construct the new wallet config using:
	fullTree := v3.WalletConfigTreeNodes(mainSignerLeaf, authSignerLeaf)

	return &v3.WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree:        fullTree,
	}, nil
}

// `CreateIntentConfiguration` creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.

// `CreateRawIntentConfiguration` creates a wallet configuration where the intent's transaction batches are grouped into the initial subdigest.
func CreateRawIntentConfiguration(mainSigner common.Address, calls []*v3.CallsPayload) (*v3.WalletConfig, error) {
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

//

func GetIntentConfigurationUpdateSignature(mainSigner common.Address, calls []*v3.CallsPayload, authSigner common.Address, authSignerWallet ethwallet.Wallet) (v3.Payload, []byte, error) {
	// Get the auth signer addr
	authWalletSigner := authSignerWallet.Address()

	// Check if the wallet's authSigner is the same as the provided authSigner
	if authWalletSigner != authSigner {
		return nil, nil, fmt.Errorf("auth signer wallet's auth signer is not the same as the provided auth signer")
	}

	// Create the initial intent configuration using the main signer and auth signer addresses.
	initialConfig, err := CreateInitialIntentConfiguration(mainSigner, authSigner)
	if err != nil {
		return nil, nil, err
	}

	// Wallet address
	address, _, _, err := EncodeWalletDeployment(initialConfig, V3SequenceContext())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode wallet deployment: %w", err)
	}

	// Create a config update payload
	intentConfigUpdatePayload := v3.NewConfigUpdatePayload(address, nil, initialConfig.Tree.ImageHash().Hash)

	// Create a new auth signer
	authorizationSigner := NewAuthorizationSigner(&authSignerWallet)

	// Create a new wallet using the intentConfig v3.WalletConfig.
	wallet, err := V3NewWallet(WalletOptions[*v3.WalletConfig]{
		Config: initialConfig,
		Context: func() *WalletContext {
			ctx := V3SequenceContext()
			return &ctx
		}(),
	}, authorizationSigner)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	// Sign the intent config update payload
	initialConfigUpdateSig, _, err := wallet.SignV3Payload(context.Background(), intentConfigUpdatePayload)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign intent config update payload: %w", err)
	}

	return intentConfigUpdatePayload, initialConfigUpdateSig, nil
}

// `GetIntentConfigurationSignature` creates a signature for the intent configuration that can be used to bypass chain ID validation. The signature is based on the transaction bundle digests only.
func GetIntentConfigurationSignature(mainSigner common.Address, calls []*v3.CallsPayload, opts ...interface{}) ([]byte, error) {

	// Create the intent configuration using the batched transactions.
	config, err := CreateRawIntentConfiguration(mainSigner, calls)
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

type AuthorizationSigner struct {
	Wallet *ethwallet.Wallet
}

func NewAuthorizationSigner(wallet *ethwallet.Wallet) *AuthorizationSigner {
	return &AuthorizationSigner{Wallet: wallet}
}

func (n *AuthorizationSigner) Address() common.Address {
	return n.Wallet.Address()
}

func (n *AuthorizationSigner) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, error) {
	sig, err := n.Wallet.SignMessage(digest.Bytes())
	if err != nil {
		return nil, err
	}

	return append(sig, 1), nil
}
