package sequence_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"

	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateIntentSubdigestLeaves_Valid(t *testing.T) {
	// Create a valid transaction (stx) with required fields.
	stx := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		Nonce:         big.NewInt(0),
		// other required fields for stx (if any) are assumed to be set.
	}

	// Use the valid transaction in a slice (representing one batch).
	txns := []*sequence.Transaction{stx}
	batches := [][]*sequence.Transaction{txns}

	bundle, err := sequence.CreateIntentBundle(txns)
	require.NoError(t, err)

	leaves, err := sequence.CreateIntentSubdigestLeaves(batches)
	require.NoError(t, err)
	require.Len(t, leaves, 1, "expected one subdigest leaf")

	// Verify that the leaf's digest matches the transaction's digest.
	digest, err := bundle.Digest()
	require.NoError(t, err)
	require.Equal(t, digest, leaves[0].Subdigest.Hash, "digests do not match")
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

		// Print the eoa1's address
		fmt.Println("eoa1 address:", eoa1.Address())

		// Dump the signature
		spew.Dump(signature)

		// Verify signature format
		require.Equal(t, byte(0x02), signature[0], "signature should start with version byte 0x02 (NoChainID)")

		// Get the subdigest from the config's tree
		var subdigestLeaf *v3.WalletConfigTreeSubdigestLeaf
		if node, ok := config.Tree.(*v3.WalletConfigTreeNode); ok {
			if rightNode, ok := node.Right.(*v3.WalletConfigTreeSubdigestLeaf); ok {
				subdigestLeaf = rightNode
				fmt.Println("decoded subdigest leaf:", subdigestLeaf)
			}
		}
		require.NotNil(t, subdigestLeaf, "config should contain a subdigest leaf")

		// Verify the signature can be decoded
		sig, err := v3.Core.DecodeSignature(signature)
		require.NoError(t, err, "signature should be decodable")

		// Dump the sig
		fmt.Println("DUMPING SIG")
		spew.Dump(sig)

		// Verify signature type by checking the first byte
		require.Equal(t, byte(0x02), signature[0], "signature should be a NoChainID signature type")

		// Print the image hash of the subdigest
		fmt.Println(subdigestLeaf.ImageHash())

		// Print the full signature in hex
		fmt.Println("full signature:", sig)

		// Get the full signature in string
		sigDataStr, err := sig.Data()
		require.NoError(t, err)
		fmt.Println("full signature hex:", common.Bytes2Hex(sigDataStr))

		subdigestStr := subdigestLeaf.Subdigest.Hash.Hex()
		fmt.Println("subdigest hex:", subdigestStr)

		// Verify the signature contains the subdigest
		require.Contains(t, common.Bytes2Hex(sigDataStr), subdigestStr[2:], "signature should contain the subdigest")
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

	// Now, we want to transfer 50 tokens from the wallet to a recipient.
	recipient := common.HexToAddress("0x3333333333333333333333333333333333333333")
	transferCalldata, err := erc20.Encode("transfer", recipient, big.NewInt(50))
	require.NoError(t, err)

	// Create a transaction that calls the ERC20 transfer.
	// (Since the wallet holds the tokens, a normal transfer works to "spend" tokens.)
	tx := &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		To:            erc20.Address,
		Data:          transferCalldata,
	}

	// Wrap the transaction in a single batch.
	txns := []*sequence.Transaction{tx}
	batches := [][]*sequence.Transaction{txns}

	// Generate a configuration signature for the batch.
	// (The configuration signature "attests" that the wallet owner authorizes
	// this bundle of transactions.)
	configSig, err := sequence.CreateIntentConfigurationSignature(mainSigner.Address(), batches)
	require.NoError(t, err)
	// For a NoChainID signature the version byte is expected to be 0x02.
	assert.Equal(t, byte(0x02), configSig[0], "configuration signature should start with 0x02")

	// Use a v2 dummy Sequence wallet
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(1, batches)
	require.NoError(t, err)
	require.NotNil(t, wallet)

	// Mint 100 tokens to the wallet.
	mintCalldata, err := erc20.Encode("mockMint", wallet.Address(), big.NewInt(100))
	require.NoError(t, err)
	err = testutil.SignAndSend(t, wallet, erc20.Address, mintCalldata)
	require.NoError(t, err)

	// Verify that the wallet received 100 tokens.
	balances, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{wallet.Address().Hex()})
	require.NoError(t, err)
	require.Len(t, balances, 1)
	require.Equal(t, "100", balances[0])

	// Get the signed transactions
	signedTxns, err := wallet.GetSignedIntentTransactions(context.Background(), txns, configSig)
	require.NoError(t, err)

	// Send the transaction bundle – note that we do not use wallet.SignTransaction here
	// since we are testing that a pre-generated configuration signature works.
	metaTxnID, sentTx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTxns)
	require.NoError(t, err)
	require.NotEmpty(t, metaTxnID)
	require.NotNil(t, sentTx)

	receipt, err := waitReceipt(context.Background())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "meta transaction should execute successfully")

	// Verify that the transfer took place:
	// the wallet's token balance should have dropped from 100 to 50
	// and the recipient should have received 50 tokens.
	walletBalance, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{wallet.Address().Hex()})
	require.NoError(t, err)
	require.Len(t, walletBalance, 1)
	require.Equal(t, "50", walletBalance[0], "wallet balance should be reduced to 50")

	recipientBalance, err := testutil.ContractQuery(testChain.Provider, erc20.Address, "balanceOf(address)", "uint256", []string{recipient.Hex()})
	require.NoError(t, err)
	require.Len(t, recipientBalance, 1)
	require.Equal(t, "50", recipientBalance[0], "recipient should receive 50 tokens")
}
