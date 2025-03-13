package sequence_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"

	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateIntentBundle_Valid(t *testing.T) {
	// Create a valid payload
	payload := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           big.NewInt(0),
				Data:            []byte("transaction1"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	bundle, err := sequence.CreateIntentBundle(payload)
	require.NoError(t, err)
	require.NotNil(t, bundle)

	spew.Dump(bundle)
}

// TestCreateIntentDigestLeaves_Valid creates a valid payload and computes the intent digest
func TestCreateIntentDigestLeaves_Valid(t *testing.T) {
	// Create valid payloads with required fields
	payload1 := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           nil,
				Data:            []byte("transaction1"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	payload2 := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(1),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           nil,
				Data:            []byte("transaction2"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	payload3 := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(2),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           nil,
				Data:            []byte("transaction3"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	t.Run("One batch", func(t *testing.T) {
		// Use a single payload
		batches := []v3.DecodedPayload{payload1}

		bundle, err := sequence.CreateIntentBundle(payload1)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		spew.Dump((*tree))

		// Type assert to the concrete type
		anyAddressLeaf, ok := (*tree).(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "tree should be a WalletConfigTreeAnyAddressSubdigestLeaf")

		// Verify that the leaf's digest matches the payload's digest
		dummyWallet := common.Address{}
		dummyChainID := big.NewInt(0)

		// Add missing fields that might be required for EIP712
		bundle.ParentWallets = []common.Address{}

		digest, err := v3.HashPayload(dummyWallet, dummyChainID, bundle)
		require.NoError(t, err)
		require.Equal(t, common.BytesToHash(digest[:]), anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Two batches", func(t *testing.T) {
		// Use two payloads
		batches := []v3.DecodedPayload{payload1, payload2}

		bundle1, err := sequence.CreateIntentBundle(payload1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(payload2)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		spew.Dump((*tree))

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		// For a node with two leaves, we should check both digests
		leftLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		rightLeaf, ok := nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		// Verify that both leaves' digests match the payload's digest.
		dummyWallet := common.Address{}
		dummyChainID := big.NewInt(0)

		// Add missing fields that might be required for EIP712
		bundle1.ParentWallets = []common.Address{}
		bundle2.ParentWallets = []common.Address{}

		digest1, err := v3.HashPayload(dummyWallet, dummyChainID, bundle1)
		require.NoError(t, err)
		digest2, err := v3.HashPayload(dummyWallet, dummyChainID, bundle2)
		require.NoError(t, err)
		require.Equal(t, common.BytesToHash(digest1[:]), leftLeaf.Digest.Hash, "left leaf digest does not match")
		require.Equal(t, common.BytesToHash(digest2[:]), rightLeaf.Digest.Hash, "right leaf digest does not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		// Use three payloads
		batches := []v3.DecodedPayload{payload1, payload2, payload3}

		bundle1, err := sequence.CreateIntentBundle(payload1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(payload2)
		require.NoError(t, err)
		bundle3, err := sequence.CreateIntentBundle(payload3)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		spew.Dump((*tree))

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		// For a node with three leaves, we should check all three digests
		leftLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeNode)
		require.True(t, ok, "left leaf should be WalletConfigTreeNode")

		rightLeaf, ok := nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		leftLeftLeaf, ok := leftLeaf.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		leftRightLeaf, ok := leftLeaf.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		// Verify that all leaves' digests match the payload's digest.
		dummyWallet := common.Address{}
		dummyChainID := big.NewInt(0)

		// Add missing fields that might be required for EIP712
		bundle1.ParentWallets = []common.Address{}
		bundle2.ParentWallets = []common.Address{}
		bundle3.ParentWallets = []common.Address{}

		digest1, err := v3.HashPayload(dummyWallet, dummyChainID, bundle1)
		require.NoError(t, err)
		digest2, err := v3.HashPayload(dummyWallet, dummyChainID, bundle2)
		require.NoError(t, err)
		digest3, err := v3.HashPayload(dummyWallet, dummyChainID, bundle3)
		require.NoError(t, err)
		require.Equal(t, common.BytesToHash(digest1[:]), leftLeftLeaf.Digest.Hash, "left left leaf digest does not match")
		require.Equal(t, common.BytesToHash(digest2[:]), leftRightLeaf.Digest.Hash, "left right leaf digest does not match")
		require.Equal(t, common.BytesToHash(digest3[:]), rightLeaf.Digest.Hash, "right leaf digest does not match")
	})
}

func TestCreateIntentConfiguration_Valid(t *testing.T) {
	// Create a valid payload
	payload := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           nil,
				Data:            nil,
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	batches := []v3.DecodedPayload{payload}

	// Use a valid main signer address.
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	config, err := sequence.CreateIntentConfiguration(mainSigner, batches)
	require.NoError(t, err)
	require.NotNil(t, config)
}

func TestCreateIntentConfigurationSignature(t *testing.T) {
	// Create test wallets
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Create a mock transaction
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
	assert.NoError(t, err)

	payload := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              callmockContract.Address,
				Value:           nil,
				Data:            calldata,
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	t.Run("signature matches subdigest", func(t *testing.T) {
		// Use the payload directly
		batches := []v3.DecodedPayload{payload}

		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), batches)
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), batches)
		require.NoError(t, err)

		// Verify signature format
		// Root: TreeNode
		// Left: AddressLeaf
		// Right: AnyAddressSubdigestLeaf
		require.Equal(t, byte(0x06), signature[0], "signature should start with version byte 0x06 (Node)")

		// Get the subdigest from the config's tree
		var anyAddressSubdigestLeaf *v3.WalletConfigTreeAnyAddressSubdigestLeaf
		if node, ok := config.Tree.(*v3.WalletConfigTreeNode); ok {
			if rightNode, ok := node.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf); ok {
				anyAddressSubdigestLeaf = rightNode
				fmt.Println("decoded any address subdigest leaf:", anyAddressSubdigestLeaf)
			}
		}
		require.NotNil(t, anyAddressSubdigestLeaf, "config should contain a any address subdigest leaf")

		// Verify the signature can be decoded
		sig, err := v3.Core.DecodeSignature(signature)
		require.NoError(t, err, "signature should be decodable")

		// Get the full signature in string
		sigDataStr, err := sig.Data()
		require.NoError(t, err)

		anyAddressSubdigestStr := anyAddressSubdigestLeaf.Digest.Hash.Hex()

		// Verify the signature contains the any address digest
		require.Contains(t, common.Bytes2Hex(sigDataStr), anyAddressSubdigestStr[2:], "signature should contain the any address digest")
	})

	t.Run("different transactions produce different signatures", func(t *testing.T) {
		// Create two different payloads
		calldata1, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		require.NoError(t, err)

		calldata2, err := callmockContract.Encode("testCall", big.NewInt(66), ethcoder.MustHexDecode("0x332255"))
		require.NoError(t, err)

		payload1 := v3.DecodedPayload{
			Kind:      v3.KindTransactions,
			NoChainId: true,
			Nonce:     big.NewInt(0),
			Space:     big.NewInt(0),
			Calls: []v3.Call{
				{
					To:              callmockContract.Address,
					Value:           nil,
					Data:            calldata1,
					GasLimit:        big.NewInt(0),
					DelegateCall:    false,
					OnlyFallback:    false,
					BehaviorOnError: v3.Revert,
				},
			},
		}

		payload2 := v3.DecodedPayload{
			Kind:      v3.KindTransactions,
			NoChainId: true,
			Nonce:     big.NewInt(0),
			Space:     big.NewInt(0),
			Calls: []v3.Call{
				{
					To:              callmockContract.Address,
					Value:           nil,
					Data:            calldata2,
					GasLimit:        big.NewInt(0),
					DelegateCall:    false,
					OnlyFallback:    false,
					BehaviorOnError: v3.Revert,
				},
			},
		}

		// Create signatures for each payload as separate batches
		sig1, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), []v3.DecodedPayload{payload1})
		require.NoError(t, err)

		sig2, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), []v3.DecodedPayload{payload2})
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Use the payload directly
		batches := []v3.DecodedPayload{payload}

		// Create the same payload signature twice
		sig1, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), batches)
		require.NoError(t, err)

		sig2, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), batches)
		require.NoError(t, err)

		// Verify signatures are the same
		require.Equal(t, sig1, sig2, "same transactions should produce same signatures")
	})
}

func TestCreateIntentConfigurationSignature_MultipleTransactions(t *testing.T) {
	// Create test wallets
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Create a payload with multiple calls
	multiCallPayload := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              common.Address{},
				Value:           nil,
				Data:            []byte("transaction1"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
			{
				To:              common.Address{},
				Value:           nil,
				Data:            []byte("transaction2"),
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	batches := []v3.DecodedPayload{multiCallPayload}

	// Create a signature
	sig, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), batches)
	require.NoError(t, err)

	// Convert the full signature into a hex string.
	sigHex := common.Bytes2Hex(sig)

	// Create the bundle from the payload
	bundle, err := sequence.CreateIntentBundle(multiCallPayload)
	require.NoError(t, err)

	// Compute the digest of the bundle
	dummyWallet := common.Address{}
	dummyChainID := big.NewInt(0)
	bundleDigest, err := v3.HashPayload(dummyWallet, dummyChainID, bundle)
	require.NoError(t, err)

	// Expect that the signature (in hex) contains the substrings of the bundle's digest.
	bundleDigestHex := common.BytesToHash(bundleDigest[:]).Hex()
	assert.Contains(t, sigHex, bundleDigestHex[2:], "signature should contain transaction bundle digest")
}

func TestConfigurationSignatureERC20Transfer(t *testing.T) {
	// Create the test main signer
	mainSigner, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Deploy the ERC20 mock contract.
	erc20, _ := testChain.Deploy(t, "ERC20Mock")
	fmt.Println("ERC20Mock address:", erc20.Address.Hex())

	// Now, we want to transfer 50 tokens from the wallet to a recipient.
	recipient := common.HexToAddress("0x3333333333333333333333333333333333333333")
	transferCalldata, err := erc20.Encode("transfer", recipient, big.NewInt(50))
	require.NoError(t, err)

	// Create a payload for ERC20 transfer
	transferPayload := v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: true,
		Nonce:     big.NewInt(0),
		Space:     big.NewInt(0),
		Calls: []v3.Call{
			{
				To:              erc20.Address,
				Value:           nil,
				Data:            transferCalldata,
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.Revert,
			},
		},
	}

	batches := []v3.DecodedPayload{transferPayload}
	require.NoError(t, err)

	spew.Dump(batches)

	// Generate a configuration signature for the batch.
	configSig, err := sequence.CreateIntentConfigurationSignature(mainSigner.Address(), batches)
	require.NoError(t, err)

	// For a Nested signature the version byte is expected to be 0x06.
	assert.Equal(t, byte(0x06), configSig[0], "configuration signature should start with 0x06")

	// Use a v3 dummy Sequence wallet
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(1, []v3.DecodedPayload{transferPayload})
	require.NoError(t, err)
	require.NotNil(t, wallet)

	// Mint 100 tokens to the wallet.
	mintCalldata, err := erc20.Encode("mockMint", wallet.Address(), big.NewInt(100))
	require.NoError(t, err)

	// Use the deploy wallet to mint tokens
	deployWallet := testChain.GetDeployWallet()
	signedTx, err := deployWallet.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:   &erc20.Address,
		Data: mintCalldata,
	})
	require.NoError(t, err)
	_, wait, err := deployWallet.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
	mintReceipt, err := wait(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(1), mintReceipt.Status, "mint transaction should succeed")

	// Verify that the wallet received 100 tokens.
	balances, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{wallet.Address().Hex()})
	require.NoError(t, err)
	require.Len(t, balances, 1)
	require.Equal(t, "100", balances[0])

	// Get the signed transactions using the DecodedPayload directly
	signedTxns, err := wallet.GetSignedIntentTransactionsWithDecodedPayloads(context.Background(), []v3.DecodedPayload{transferPayload}, configSig)
	require.NoError(t, err)

	// Send the transaction bundle
	metaTxnID, sentTx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTxns)
	require.NoError(t, err)
	require.NotEmpty(t, metaTxnID)
	require.NotNil(t, sentTx)

	spew.Dump(sentTx)

	receipt, err := waitReceipt(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "meta transaction should execute successfully")

	// Verify that the transfer took place
	walletBalance, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{wallet.Address().Hex()})
	require.NoError(t, err)

	recipientBalance, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{recipient.Hex()})
	require.NoError(t, err)

	require.Equal(t, "50", walletBalance[0], "wallet balance should be reduced to 50")
	require.Equal(t, "50", recipientBalance[0], "recipient should receive 50 tokens")
}
