package sequence_test

import (
	"context"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/lib/eip6492"
	"github.com/stretchr/testify/require"
)

// These tests require a live testchain (make start-testchain or make start-testchain-fork).
// They create a Sequence wallet, sign a message, wrap in EIP-6492, and validate via
// ValidateEIP6492Offchain or ValidateEIP6492Onchain. Onchain tests skip if the validator
// is not deployed at EIP_6492_ADDRESS on the chain.

func TestEIP6492Live_ValidateSequenceWalletMessage_Offchain(t *testing.T) {
	// V2 counterfactual wallet on testchain
	wallet, err := testChain.V2DummySequenceWallet(10, true)
	require.NoError(t, err)

	message := []byte("hello world!")
	_, eip191Message := accounts.TextAndHash(message)

	sig, err := wallet.SignMessage([]byte(eip191Message))
	require.NoError(t, err)

	wrapped, err := sequence.EIP6492Signature(sig, wallet.GetWalletConfig())
	require.NoError(t, err)

	digest := common.BytesToHash(accounts.TextHash(message))
	ctx := context.Background()

	valid, err := eip6492.ValidateEIP6492Offchain(ctx, testChain.Provider, wallet.Address(), digest, wrapped, nil)
	require.NoError(t, err)
	require.True(t, valid, "EIP-6492 offchain validation should succeed for Sequence wallet message signature")
}

func TestEIP6492Live_ValidateSequenceWalletMessage_ViaIsValidMessageSignature(t *testing.T) {
	// Same flow as above but via the public API to ensure full path works
	wallet, err := testChain.V2DummySequenceWallet(11, true)
	require.NoError(t, err)

	message := []byte("eip6492 live test")
	_, eip191Message := accounts.TextAndHash(message)

	sig, err := wallet.SignMessage([]byte(eip191Message))
	require.NoError(t, err)

	wrapped, err := sequence.EIP6492Signature(sig, wallet.GetWalletConfig())
	require.NoError(t, err)

	isValid, err := sequence.IsValidMessageSignature(
		wallet.Address(),
		message,
		wrapped,
		testChain.ChainID(),
		testChain.Provider,
		nil,
	)
	require.NoError(t, err)
	require.True(t, isValid, "IsValidMessageSignature should succeed for EIP-6492 wrapped Sequence wallet signature")
}

func TestEIP6492Live_ValidateDeployedSequenceWallet_Offchain(t *testing.T) {
	// Deploy the wallet first, then validate (tests both counterfactual deploy path and already-deployed path)
	wallet, err := testChain.V2DummySequenceWallet(12, false)
	require.NoError(t, err)

	message := []byte("deployed wallet sign")
	_, eip191Message := accounts.TextAndHash(message)

	sig, err := wallet.SignMessage([]byte(eip191Message))
	require.NoError(t, err)

	wrapped, err := sequence.EIP6492Signature(sig, wallet.GetWalletConfig())
	require.NoError(t, err)

	digest := common.BytesToHash(accounts.TextHash(message))
	ctx := context.Background()

	valid, err := eip6492.ValidateEIP6492Offchain(ctx, testChain.Provider, wallet.Address(), digest, wrapped, nil)
	require.NoError(t, err)
	require.True(t, valid, "EIP-6492 offchain validation should succeed for deployed Sequence wallet")
}

func TestEIP6492Live_ValidateSequenceWalletMessage_Onchain(t *testing.T) {
	// Same as offchain test but via pre-deployed validator at EIP_6492_ADDRESS.
	// Skips if the validator is not deployed on this chain.
	wallet, err := testChain.V2DummySequenceWallet(13, true)
	require.NoError(t, err)

	message := []byte("hello world onchain!")
	_, eip191Message := accounts.TextAndHash(message)

	sig, err := wallet.SignMessage([]byte(eip191Message))
	require.NoError(t, err)

	wrapped, err := sequence.EIP6492Signature(sig, wallet.GetWalletConfig())
	require.NoError(t, err)

	digest := common.BytesToHash(accounts.TextHash(message))
	ctx := context.Background()

	valid, err := eip6492.ValidateEIP6492Onchain(ctx, testChain.Provider, wallet.Address(), digest, wrapped, nil)
	if err != nil && strings.Contains(err.Error(), "returned no data") {
		t.Skipf("EIP-6492 validator not deployed on this chain: %v", err)
	}
	require.NoError(t, err)
	require.True(t, valid, "EIP-6492 onchain validation should succeed for Sequence wallet message signature")
}

func TestEIP6492Live_ValidateDeployedSequenceWallet_Onchain(t *testing.T) {
	// Deployed wallet, validated via onchain contract. Skips if validator not deployed.
	wallet, err := testChain.V2DummySequenceWallet(14, false)
	require.NoError(t, err)

	message := []byte("deployed wallet onchain sign")
	_, eip191Message := accounts.TextAndHash(message)

	sig, err := wallet.SignMessage([]byte(eip191Message))
	require.NoError(t, err)

	wrapped, err := sequence.EIP6492Signature(sig, wallet.GetWalletConfig())
	require.NoError(t, err)

	digest := common.BytesToHash(accounts.TextHash(message))
	ctx := context.Background()

	valid, err := eip6492.ValidateEIP6492Onchain(ctx, testChain.Provider, wallet.Address(), digest, wrapped, nil)
	if err != nil && strings.Contains(err.Error(), "returned no data") {
		t.Skipf("EIP-6492 validator not deployed on this chain: %v", err)
	}
	require.NoError(t, err)
	require.True(t, valid, "EIP-6492 onchain validation should succeed for deployed Sequence wallet")
}
