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
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"

	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateIntentCallsPayload_Valid(t *testing.T) {
	// Create an intent operation
	op := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           big.NewInt(0),
			Data:            []byte("transaction1"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	bundle, err := sequence.CreateIntentCallsPayload(op)
	require.NoError(t, err)
	require.NotNil(t, bundle)

	// spew.Dump(bundle)
}

// TestCreateIntentDigestTree_Valid creates a valid payload and computes the intent digest
func TestCreateIntentDigestTree_Valid(t *testing.T) {
	// Create valid intent operations with required fields
	op1 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction1"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	op2 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction2"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	op3 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction3"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	t.Run("One batch", func(t *testing.T) {
		// Use a single payload
		batches := []*sequence.IntentOperation{op1}

		bundle, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		// spew.Dump((*tree))

		// Type assert to the concrete type
		anyAddressLeaf, ok := (*tree).(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "tree should be a WalletConfigTreeAnyAddressSubdigestLeaf")

		digest := bundle.Digest()
		require.Equal(t, digest.Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Two batches", func(t *testing.T) {
		// Use two payloads
		batches := []*sequence.IntentOperation{op1, op2}

		bundle1, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentCallsPayload(op2)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		// spew.Dump((*tree))

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		// For a node with two leaves, we should check both digests
		leftLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		rightLeaf, ok := nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		digest1 := bundle1.Digest()
		digest2 := bundle2.Digest()
		require.Equal(t, digest1.Hash, leftLeaf.Digest.Hash, "left leaf digest does not match")
		require.Equal(t, digest2.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		// Use three payloads
		batches := []*sequence.IntentOperation{op1, op2, op3}

		bundle1, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentCallsPayload(op2)
		require.NoError(t, err)
		bundle3, err := sequence.CreateIntentCallsPayload(op3)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		// spew.Dump((*tree))

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

		digest1 := bundle1.Digest()
		digest2 := bundle2.Digest()
		digest3 := bundle3.Digest()
		require.Equal(t, digest1.Hash, leftLeftLeaf.Digest.Hash, "left left leaf digest does not match")
		require.Equal(t, digest2.Hash, leftRightLeaf.Digest.Hash, "left right leaf digest does not match")
		require.Equal(t, digest3.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})
}

// TestCreateIntentTree_Valid creates a valid payload and computes the intent digest
func TestCreateIntentTree_Valid(t *testing.T) {
	// Create valid intent operations with required fields
	op1 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction1"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	op2 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction2"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	op3 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction3"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	t.Run("One batch", func(t *testing.T) {
		// Use a single payload
		batches := []*sequence.IntentOperation{op1}

		tree, err := sequence.CreateIntentTree(common.Address{}, batches)
		require.NoError(t, err)
		require.NotNil(t, tree)

		// spew.Dump(tree)

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		addressLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAddressLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAddressLeaf")
		require.Equal(t, addressLeaf.Address, common.Address{}, "address leaf should be the main signer")

		anyAddressLeaf, ok := nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		// Get the digest of the leaf
		digest := anyAddressLeaf.Digest.Hash
		require.NotNil(t, digest)

		// Get the digest of the payload
		bundle, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)
		require.NotNil(t, bundle)
	})

	t.Run("Two batches", func(t *testing.T) {
		// Use two payloads
		batches := []*sequence.IntentOperation{op1, op2}

		tree, err := sequence.CreateIntentTree(common.Address{}, batches)
		require.NoError(t, err)
		require.NotNil(t, tree)

		// spew.Dump(tree)

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		addressLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAddressLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAddressLeaf")
		require.Equal(t, addressLeaf.Address, common.Address{}, "address leaf should be the main signer")

		nodeRight, ok := nodeTree.Right.(*v3.WalletConfigTreeNode)
		require.True(t, ok, "right node should be WalletConfigTreeNode")

		anyAddressLeaf, ok := nodeRight.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right node should be WalletConfigTreeAnyAddressSubdigestLeaf")

		anyAddressLeaf2, ok := nodeRight.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		// Get the digest of the leaf
		digest := anyAddressLeaf.Digest.Hash
		require.NotNil(t, digest)

		// Get the digest of the payload
		bundle, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)
		require.NotNil(t, bundle)

		bundle2, err := sequence.CreateIntentCallsPayload(op2)
		require.NoError(t, err)
		require.NotNil(t, bundle2)

		bundleDigest := bundle.Digest()
		bundle2Digest := bundle2.Digest()

		require.Equal(t, bundleDigest.Hash, anyAddressLeaf2.Digest.Hash, "digests do not match")
		require.Equal(t, bundle2Digest.Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		// Use three payloads
		batches := []*sequence.IntentOperation{op1, op2, op3}

		tree, err := sequence.CreateIntentTree(common.Address{}, batches)
		require.NoError(t, err)

		// spew.Dump(tree)

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		addressLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAddressLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAddressLeaf")
		require.Equal(t, addressLeaf.Address, common.Address{}, "address leaf should be the main signer")

		nodeRight, ok := nodeTree.Right.(*v3.WalletConfigTreeNode)
		require.True(t, ok, "right node should be WalletConfigTreeNode")

		anyAddressLeaf, ok := nodeRight.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right node should be WalletConfigTreeAnyAddressSubdigestLeaf")

		nodeLeft, ok := nodeRight.Left.(*v3.WalletConfigTreeNode)
		require.True(t, ok, "left node should be WalletConfigTreeNode")

		anyAddressLeaf2, ok := nodeLeft.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		anyAddressLeaf3, ok := nodeLeft.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		digest1 := anyAddressLeaf.Digest.Hash
		require.NotNil(t, digest1)

		digest2 := anyAddressLeaf2.Digest.Hash
		require.NotNil(t, digest2)

		digest3 := anyAddressLeaf3.Digest.Hash
		require.NotNil(t, digest3)

		bundle, err := sequence.CreateIntentCallsPayload(op1)
		require.NoError(t, err)
		require.NotNil(t, bundle)

		bundle2, err := sequence.CreateIntentCallsPayload(op2)
		require.NoError(t, err)
		require.NotNil(t, bundle2)

		bundle3, err := sequence.CreateIntentCallsPayload(op3)
		require.NoError(t, err)
		require.NotNil(t, bundle3)

		bundleDigest := bundle.Digest()
		bundle2Digest := bundle2.Digest()
		bundle3Digest := bundle3.Digest()

		require.Equal(t, bundleDigest.Hash, anyAddressLeaf3.Digest.Hash, "digests do not match")
		require.Equal(t, bundle2Digest.Hash, anyAddressLeaf2.Digest.Hash, "digests do not match")
		require.Equal(t, bundle3Digest.Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})
}

func TestCreateIntentConfiguration_Valid(t *testing.T) {
	// Create a valid payload
	op := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            nil,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	batches := []*sequence.IntentOperation{op}

	// Use a valid main signer address.
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	config, err := sequence.CreateIntentConfiguration(mainSigner, batches)
	require.NoError(t, err)
	require.NotNil(t, config)
}

func TestGetIntentConfigurationSignature(t *testing.T) {
	// Create test wallets
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Create a mock transaction
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
	require.NoError(t, err)

	op := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	t.Run("signature matches subdigest", func(t *testing.T) {
		// Use the payload directly
		batches := []*sequence.IntentOperation{op}

		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), batches)
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), batches)
		require.NoError(t, err)

		// fmt.Println("==> signature", common.Bytes2Hex(signature))

		// Verify signature format
		// Root: TreeNode
		// Left: AddressLeaf
		// Right: AnyAddressSubdigestLeaf
		require.Equal(t, byte(0x04), signature[0], "signature should start with version byte 0x04 (Branch)")

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

		// Get the config from the signature
		recoveredConfig, _, err := sig.RecoverSubdigest(context.Background(), anyAddressSubdigestLeaf.Digest, nil)
		// spew.Dump(recoveredConfig)

		require.NoError(t, err)
		require.NotNil(t, recoveredConfig, "recovered config should not be nil")

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

		op1 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
			{
				To:              callmockContract.Address,
				Value:           nil,
				Data:            calldata1,
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.BehaviorOnErrorRevert,
			},
		}, big.NewInt(0), big.NewInt(0))

		op2 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
			{
				To:              callmockContract.Address,
				Value:           nil,
				Data:            calldata2,
				GasLimit:        big.NewInt(0),
				DelegateCall:    false,
				OnlyFallback:    false,
				BehaviorOnError: v3.BehaviorOnErrorRevert,
			},
		}, big.NewInt(0), big.NewInt(0))

		// Create signatures for each payload as separate batches
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*sequence.IntentOperation{op1})
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*sequence.IntentOperation{op2})
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Use the payload directly
		ops := []*sequence.IntentOperation{op}

		// Create the same payload signature twice
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), ops)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), ops)
		require.NoError(t, err)

		// Verify signatures are the same
		require.Equal(t, sig1, sig2, "same transactions should produce same signatures")
	})
}

func TestGetIntentConfigurationSignature_MultipleTransactions(t *testing.T) {
	// Create test wallets
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Create a payload with multiple calls
	op1 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction1"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
		{
			To:              common.Address{},
			Value:           nil,
			Data:            []byte("transaction2"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	ops := []*sequence.IntentOperation{op1}

	// Create a signature
	sig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), ops)
	require.NoError(t, err)

	// Convert the full signature into a hex string.
	sigHex := common.Bytes2Hex(sig)

	// Create the bundle from the payload
	bundle, err := sequence.CreateIntentCallsPayload(op1)
	require.NoError(t, err)

	// Compute the digest of the bundle
	bundleDigest := bundle.Digest()
	require.NoError(t, err)

	// Expect that the signature (in hex) contains the substrings of the bundle's digest.
	bundleDigestHex := bundleDigest.Hash.Hex()
	require.Contains(t, sigHex, bundleDigestHex[2:], "signature should contain transaction bundle digest")
}

func TestIntentTransactionToGuestModuleDeployAndCall(t *testing.T) {
	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata1, err := callmockContract.Encode("setRevertFlag", false)
	assert.NoError(t, err)
	calldata2, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
	assert.NoError(t, err)

	op1 := sequence.NewIntentOperation(testChain.ChainID(), []v3.Call{
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata1,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata2,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	ops := []*sequence.IntentOperation{op1}

	// Ensure dummy sequence wallet from the intent operation
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(testutil.RandomSeed(), ops, true)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Assert the wallet is undeployed -- this is desired so we relayer the txn to the guest module
	isDeployed, err := wallet.IsDeployed()
	assert.NoError(t, err)
	if isDeployed {
		t.Fatal("expecting wallet to be undeployed")
	}

	// Wallet deployment data
	_, walletFactoryAddress, walletDeployData, err := sequence.EncodeWalletDeployment(wallet.GetWalletConfig(), wallet.GetWalletContext())
	assert.NoError(t, err)

	// Get the main signer
	signers := wallet.GetWalletConfig().Signers()
	var mainSigner common.Address
	for addr := range signers {
		mainSigner = addr
		break
	}
	require.NotNil(t, mainSigner)

	// Generate a configuration signature for the batch.
	intentConfigSig, err := sequence.GetIntentConfigurationSignature(mainSigner, ops)
	require.NoError(t, err)

	bundle, err := sequence.CreateIntentCallsPayload(op1)
	assert.NoError(t, err)

	// Since the opHash is generated w/ the wallet addr, create a payload w/ the valid wallet address
	opHashBundle, err := sequence.CreateIntentCallsPayload(op1, wallet.Address())
	assert.NoError(t, err)

	// fmt.Println("==> bundle.Digest", bundle.Digest().Hash)

	signedExecdata, err := contracts.V3.WalletStage1Module.Encode("execute", bundle.Encode(common.Address{}), intentConfigSig)
	assert.NoError(t, err)

	// fmt.Println("==> signedPayloadHash", bundle.Digest().Hash.Hex())
	// fmt.Println("==> payload", common.Bytes2Hex(bundle.Encode(wallet.Address())))
	// fmt.Println("==> signature", common.Bytes2Hex(intentConfigSig))
	// fmt.Println("==> signedExecData", common.Bytes2Hex(signedExecdata))

	guestBundle := []v3.Call{
		{
			To:   walletFactoryAddress,
			Data: walletDeployData,
		},
		{
			To:   wallet.Address(),
			Data: signedExecdata,
		},
	}

	guestAddress := testChain.V3SequenceContext().GuestModuleAddress
	execdata := v3.NewCallsPayload(guestAddress, testChain.ChainID(), guestBundle, nil, nil).Encode(guestAddress)

	// Relay the txn manually, directly to the guest module
	sender := testChain.GetRelayerWallet()
	ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:       &guestAddress,
		Data:     execdata,
		GasLimit: 1000000, // TODO: compute gas limit
	})
	assert.NoError(t, err)

	signedTx, err := sender.SignTx(ntx, testChain.ChainID())
	assert.NoError(t, err)

	_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err)

	receipt, err := waitReceipt(context.Background())
	// spew.Dump(receipt)
	assert.NoError(t, err)
	assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

	// fmt.Println("==> metaTxnId", opHashBundle.Digest().String())

	// for _, logs := range receipt.Logs {
	// 	for _, logTopics := range logs.Topics {
	// 		fmt.Println("==> logs.Topic", logTopics)
	// 	}
	// 	fmt.Println("==> logs.Data", common.Bytes2Hex(logs.Data))
	// }

	// Check the value
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "2255", ret[0])

	// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
	result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, sequence.MetaTxnID(opHashBundle.Digest().String()[2:]))
	assert.NoError(t, err)
	assert.True(t, result.Status == sequence.MetaTxnExecuted)

	// Wallet should be deployed now
	isDeployed, err = wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, isDeployed)
}

func TestIntentConfigurationAddress(t *testing.T) {
	// Create context matching TypeScript test
	context := sequence.V3SequenceContext()
	context.FactoryAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	context.MainModuleAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// Main signer matching TypeScript test
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	t.Run("single operation", func(t *testing.T) {
		// Create a single operation matching TypeScript test
		operation := sequence.NewIntentOperation(
			big.NewInt(1),
			[]v3.Call{
				{
					To:              common.HexToAddress("0x0000000000000000000000000000000000000000"),
					Value:           big.NewInt(0),
					Data:            common.FromHex("0x1234"),
					GasLimit:        big.NewInt(0),
					DelegateCall:    false,
					OnlyFallback:    false,
					BehaviorOnError: v3.BehaviorOnErrorRevert,
				},
			},
			big.NewInt(0),
			big.NewInt(0),
		)

		// Create intent configuration
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*sequence.IntentOperation{operation})
		require.NoError(t, err)

		// Calculate image hash
		imageHash := config.ImageHash()

		// Calculate counterfactual address
		address, err := sequence.AddressFromImageHash(imageHash, context)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("Single Operation Test:\n")
		fmt.Printf("Address: %s\n", address.Hex())
	})

	t.Run("multiple operations", func(t *testing.T) {
		// Create multiple operations matching TypeScript test
		operations := []*sequence.IntentOperation{
			sequence.NewIntentOperation(
				big.NewInt(1),
				[]v3.Call{
					{
						To:              common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Value:           big.NewInt(0),
						Data:            common.FromHex("0x1234"),
						GasLimit:        big.NewInt(0),
						DelegateCall:    false,
						OnlyFallback:    false,
						BehaviorOnError: v3.BehaviorOnErrorRevert,
					},
				},
				big.NewInt(0),
				big.NewInt(0),
			),
			sequence.NewIntentOperation(
				big.NewInt(1),
				[]v3.Call{
					{
						To:              common.HexToAddress("0x0000000000000000000000000000000000000000"),
						Value:           big.NewInt(0),
						Data:            common.FromHex("0x5678"),
						GasLimit:        big.NewInt(0),
						DelegateCall:    false,
						OnlyFallback:    false,
						BehaviorOnError: v3.BehaviorOnErrorRevert,
					},
				},
				big.NewInt(0),
				big.NewInt(0),
			),
		}

		// Create intent configuration
		config, err := sequence.CreateIntentConfiguration(mainSigner, operations)
		require.NoError(t, err)

		// Calculate image hash
		imageHash := config.ImageHash()

		// Calculate counterfactual address
		address, err := sequence.AddressFromImageHash(imageHash, context)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("\nMultiple Operations Test:\n")
		fmt.Printf("Address: %s\n", address.Hex())
	})
}
