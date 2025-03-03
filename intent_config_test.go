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

func TestCreateIntentDigestLeaves_Valid(t *testing.T) {
	// Create a valid transaction (stx) with required fields.
	stx1 := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
		Data:          []byte("transaction1"),
	}
	stx2 := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(1),
		Data:          []byte("transaction2"),
	}
	stx3 := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(2),
		Data:          []byte("transaction3"),
	}

	t.Run("One batch", func(t *testing.T) {
		// Use the valid transaction in a slice (representing one batch).
		txns := []*sequence.Transaction{stx1}
		batches := [][]*sequence.Transaction{txns}

		bundle, err := sequence.CreateIntentBundle(txns)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		spew.Dump((*tree))

		// Type assert to the concrete type
		anyAddressLeaf, ok := (*tree).(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "tree should be a WalletConfigTreeAnyAddressSubdigestLeaf")

		// Verify that the leaf's digest matches the transaction's digest.
		digest, err := bundle.Digest()
		require.NoError(t, err)
		require.Equal(t, digest, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Two batches", func(t *testing.T) {
		// Use the valid transaction in a slice (representing one batch).
		txns1 := []*sequence.Transaction{stx1}
		txns2 := []*sequence.Transaction{stx2}
		batches := [][]*sequence.Transaction{txns1, txns2}

		bundle1, err := sequence.CreateIntentBundle(txns1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(txns2)
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

		// Verify that both leaves' digests match the transaction's digest.
		digest1, err := bundle1.Digest()
		require.NoError(t, err)
		digest2, err := bundle2.Digest()
		require.NoError(t, err)
		require.Equal(t, digest1, leftLeaf.Digest.Hash, "left leaf digest does not match")
		require.Equal(t, digest2, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		// Use the valid transaction in a slice (representing one batch).
		txns1 := []*sequence.Transaction{stx1}
		txns2 := []*sequence.Transaction{stx2}
		txns3 := []*sequence.Transaction{stx3}
		batches := [][]*sequence.Transaction{txns1, txns2, txns3}

		bundle1, err := sequence.CreateIntentBundle(txns1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(txns2)
		require.NoError(t, err)
		bundle3, err := sequence.CreateIntentBundle(txns3)
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

		// Verify that all leaves' digests match the transaction's digest.
		digest1, err := bundle1.Digest()
		require.NoError(t, err)
		digest2, err := bundle2.Digest()
		require.NoError(t, err)
		digest3, err := bundle3.Digest()
		require.NoError(t, err)
		require.Equal(t, digest1, leftLeftLeaf.Digest.Hash, "left left leaf digest does not match")
		require.Equal(t, digest2, leftRightLeaf.Digest.Hash, "left right leaf digest does not match")
		require.Equal(t, digest3, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})
}

func TestCreateIntentConfiguration_Valid(t *testing.T) {
	// Create a valid transaction.
	stx := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
	}
	txns := []*sequence.Transaction{stx}
	batches := [][]*sequence.Transaction{txns}

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

	stx := &sequence.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
	}

	t.Run("signature matches subdigest", func(t *testing.T) {
		// Wrap the transaction in a batch.
		batches := [][]*sequence.Transaction{{stx}}

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
		// Create two different transactions
		calldata1, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		require.NoError(t, err)

		calldata2, err := callmockContract.Encode("testCall", big.NewInt(66), ethcoder.MustHexDecode("0x332255"))
		require.NoError(t, err)

		tx1 := &sequence.Transaction{
			To:            callmockContract.Address,
			Data:          calldata1,
			RevertOnError: true,
			Nonce:         big.NewInt(0),
		}

		tx2 := &sequence.Transaction{
			To:            callmockContract.Address,
			Data:          calldata2,
			RevertOnError: true,
			Nonce:         big.NewInt(0),
		}

		// Create signatures for each transaction as separate batches.
		sig1, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), [][]*sequence.Transaction{{tx1}})
		require.NoError(t, err)

		sig2, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), [][]*sequence.Transaction{{tx2}})
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Wrap the transaction in a batch.
		batches := [][]*sequence.Transaction{{stx}}

		// Create the same transaction twice
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

	// Create two valid transactions with different Data fields so their digests differ.
	stx1 := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
		Data:          []byte("transaction1"),
	}
	stx2 := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
		Data:          []byte("transaction2"),
	}
	txns := []*sequence.Transaction{stx1, stx2}
	batches := [][]*sequence.Transaction{txns}

	// Create a signature
	sig, err := sequence.CreateIntentConfigurationSignature(eoa1.Address(), batches)
	require.NoError(t, err)

	// Convert the full signature into a hex string.
	sigHex := common.Bytes2Hex(sig)

	// Create the bundle from the transactions
	bundle, err := sequence.CreateIntentBundle(txns)
	require.NoError(t, err)

	// Compute the digest of the bundle
	bundleDigest, err := bundle.Digest()
	require.NoError(t, err)

	// Expect that the signature (in hex) contains the substrings of the bundle's digest.
	assert.Contains(t, sigHex, bundleDigest.Hex()[2:], "signature should contain transaction bundle digest")
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

	// Create a transaction that calls the ERC20 transfer.
	tx := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		To:            erc20.Address,
		Data:          transferCalldata,
		Nonce:         big.NewInt(0),
	}

	// Wrap the transaction in a single batch.
	txns := []*sequence.Transaction{tx}
	batches := [][]*sequence.Transaction{txns}

	require.NoError(t, err)

	// Generate a configuration signature for the batch.
	configSig, err := sequence.CreateIntentConfigurationSignature(mainSigner.Address(), batches)
	require.NoError(t, err)

	// For a Nested signature the version byte is expected to be 0x06.
	assert.Equal(t, byte(0x06), configSig[0], "configuration signature should start with 0x06")

	// Use a v3 dummy Sequence wallet
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(1, batches)
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

	// Get the signed transactions
	signedTxns, err := wallet.GetSignedIntentTransactions(context.Background(), txns, configSig)
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
