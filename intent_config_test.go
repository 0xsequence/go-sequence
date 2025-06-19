package sequence_test

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"
	"github.com/davecgh/go-spew/spew"

	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateIntentCallsPayload_Valid(t *testing.T) {
	// Create a calls payload
	calls := []v3.Call{
		{
			To:              common.Address{},
			Value:           big.NewInt(0),
			Data:            []byte("transaction1"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}

	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), calls, big.NewInt(0), big.NewInt(0))

	require.NotNil(t, payload)

	spew.Dump(payload)
}

// TestCreateIntentDigestTree_Valid creates a valid payload and computes the intent digest
func TestCreateIntentDigestTree_Valid(t *testing.T) {
	// Create valid calls payloads
	calls1 := v3.Call{
		To:              common.Address{},
		Value:           nil,
		Data:            []byte("transaction1"),
		GasLimit:        big.NewInt(0),
		DelegateCall:    false,
		OnlyFallback:    false,
		BehaviorOnError: v3.BehaviorOnErrorRevert,
	}

	payload1 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{calls1}, big.NewInt(0), big.NewInt(0))

	calls2 := v3.Call{
		To:              common.Address{},
		Value:           nil,
		Data:            []byte("transaction2"),
		GasLimit:        big.NewInt(0),
		DelegateCall:    false,
		OnlyFallback:    false,
		BehaviorOnError: v3.BehaviorOnErrorRevert,
	}
	payload2 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{calls2}, big.NewInt(0), big.NewInt(0))

	calls3 := v3.Call{
		To:              common.Address{},
		Value:           nil,
		Data:            []byte("transaction3"),
		GasLimit:        big.NewInt(0),
		DelegateCall:    false,
		OnlyFallback:    false,
		BehaviorOnError: v3.BehaviorOnErrorRevert,
	}
	payload3 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{calls3}, big.NewInt(0), big.NewInt(0))

	t.Run("One batch", func(t *testing.T) {
		leaves, err := sequence.CreateAnyAddressSubdigestTree([]*v3.CallsPayload{&payload1})
		require.NoError(t, err)

		// Create a tree from the subdigest leaves.
		tree := v3.WalletConfigTreeNodes(leaves...)
		require.NotNil(t, tree, "expected a tree")

		// Type assert to the concrete type
		anyAddressLeaf, ok := tree.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "tree should be a WalletConfigTreeAnyAddressSubdigestLeaf")

		digest := payload1.Digest()
		require.Equal(t, digest.Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Two batches", func(t *testing.T) {
		leaves, err := sequence.CreateAnyAddressSubdigestTree([]*v3.CallsPayload{&payload1, &payload2})
		require.NoError(t, err)

		// Create a tree from the subdigest leaves.
		tree := v3.WalletConfigTreeNodes(leaves...)
		require.NotNil(t, tree, "expected a tree")

		// Type assert to the concrete type
		nodeTree, ok := (tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		// For a node with two leaves, we should check both digests
		leftLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		rightLeaf, ok := nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		digest1 := payload1.Digest()
		digest2 := payload2.Digest()
		require.Equal(t, digest1.Hash, leftLeaf.Digest.Hash, "left leaf digest does not match")
		require.Equal(t, digest2.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		leaves, err := sequence.CreateAnyAddressSubdigestTree([]*v3.CallsPayload{&payload1, &payload2, &payload3})
		require.NoError(t, err)

		// Create a tree from the subdigest leaves.
		tree := v3.WalletConfigTreeNodes(leaves...)
		require.NotNil(t, tree, "expected a tree")

		// Type assert to the concrete type
		nodeTree, ok := tree.(*v3.WalletConfigTreeNode)
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

		digest1 := payload1.Digest()
		digest2 := payload2.Digest()
		digest3 := payload3.Digest()
		require.Equal(t, digest1.Hash, leftLeftLeaf.Digest.Hash, "left left leaf digest does not match")
		require.Equal(t, digest2.Hash, leftRightLeaf.Digest.Hash, "left right leaf digest does not match")
		require.Equal(t, digest3.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})
}

// TestCreateIntentTree_Valid creates a valid payload and computes the intent digest
func TestCreateIntentTree_Valid(t *testing.T) {
	// Create valid intent operations with required fields
	payload1 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	payload2 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	payload3 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1}, nil)
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
		bundle, err := sequence.CreateAnyAddressSubdigestTree([]*v3.CallsPayload{&payload1})
		require.NoError(t, err)
		require.NotNil(t, bundle)
	})

	t.Run("Two batches", func(t *testing.T) {
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1, &payload2}, nil)
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

		require.Equal(t, payload1.Digest().Hash, anyAddressLeaf2.Digest.Hash, "digests do not match")
		require.Equal(t, payload2.Digest().Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1, &payload2, &payload3}, nil)
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

		require.Equal(t, payload1.Digest().Hash, anyAddressLeaf3.Digest.Hash, "digests do not match")
		require.Equal(t, payload2.Digest().Hash, anyAddressLeaf2.Digest.Hash, "digests do not match")
		require.Equal(t, payload3.Digest().Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})
}

func TestCreateIntentConfiguration_Valid(t *testing.T) {
	// Create a valid payload
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	// Use a valid main signer address.
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, nil)
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

	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), []*v3.CallsPayload{&payload}, nil)
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload}, eoa1, nil, "", nil)
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

		payload1 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

		payload2 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload1}, eoa1, nil, "", nil)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload2}, eoa1, nil, "", nil)
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Use the payload directly
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload}, eoa1, nil, "", nil)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload}, eoa1, nil, "", nil)
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
	payload1 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	// Create a signature
	sig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), common.Address{}, []*v3.CallsPayload{&payload1}, eoa1, nil, "", nil)
	require.NoError(t, err)

	// Convert the full signature into a hex string.
	sigHex := common.Bytes2Hex(sig)

	// Expect that the signature (in hex) contains the substrings of the bundle's digest.
	require.Contains(t, sigHex, payload1.Digest().Hash.Hex()[2:], "signature should contain transaction bundle digest")
}

func TestIntentTransactionToGuestModuleDeployAndCall(t *testing.T) {
	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata1, err := callmockContract.Encode("setRevertFlag", false)
	assert.NoError(t, err)
	calldata2, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
	assert.NoError(t, err)

	_, err = ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	// Ensure dummy sequence wallet from the intent operation
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(testutil.RandomSeed(), []*v3.CallsPayload{&payload}, true)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Get the payload w/ the operation (The on-chain payload is different since it's the digest of the wallet's address vs. any address subdigest is the address zero)
	opPayload := v3.NewCallsPayload(wallet.Address(), testChain.ChainID(), payload.Calls, big.NewInt(0), big.NewInt(0))

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
	intentConfigSig, err := sequence.GetIntentConfigurationSignature(mainSigner, common.Address{}, []*v3.CallsPayload{&payload}, nil, nil, "", nil)
	require.NoError(t, err)

	// fmt.Println("==> bundle.Digest", bundle.Digest().Hash)

	signedExecdata, err := contracts.V3.WalletStage1Module.Encode("execute", payload.Encode(common.Address{}), intentConfigSig)
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
	execdata := v3.NewCallsPayload(guestAddress, testChain.ChainID(), guestBundle, big.NewInt(0), big.NewInt(0)).Encode(guestAddress)

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
	result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, sequence.MetaTxnID(opPayload.Digest().Hash.Hex()[2:]))
	assert.NoError(t, err)
	assert.True(t, result.Status == sequence.MetaTxnExecuted)

	// Wallet should be deployed now
	isDeployed, err = wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, isDeployed)
}

func TestIntentTransactionToGuestModuleDeployAndCallMultiplePayloads(t *testing.T) {
	// Create normal txn of: callmockContract.testCall(55, 0x112255) for first chain
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata1, err := callmockContract.Encode("setRevertFlag", false)
	assert.NoError(t, err)
	calldata2, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
	assert.NoError(t, err)
	calldata3, err := callmockContract.Encode("testCall", big.NewInt(6655), ethcoder.MustHexDecode("0x332266"))
	assert.NoError(t, err)

	_, err = ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	// Create multiple payloads
	payload1 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata1,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata2,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
	}, big.NewInt(0), big.NewInt(0))

	payload2 := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              callmockContract.Address,
			Value:           nil,
			Data:            calldata3,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
	}, big.NewInt(1), big.NewInt(0))

	payloads := []*v3.CallsPayload{&payload1, &payload2}

	// Ensure dummy sequence wallet from the intent operation
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(testutil.RandomSeed(), payloads, true)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Log payload digests used for signature leaves
	fmt.Println("--- Payload Digests (for signature) ---")
	for i, p := range payloads {
		fmt.Printf("Payload %d Digest: %s\n", i+1, p.Digest().Hash.Hex())
	}

	// Get the payloads w/ the operation (The on-chain payload is different since it's the digest of the wallet's address vs. any address subdigest is the address zero)
	opPayload1 := v3.NewCallsPayload(wallet.Address(), testChain.ChainID(), payload1.Calls, big.NewInt(0), big.NewInt(0))
	opPayload2 := v3.NewCallsPayload(wallet.Address(), testChain.ChainID(), payload2.Calls, big.NewInt(1), big.NewInt(0))

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

	// Generate a configuration signature for both batches
	intentConfigSig, err := sequence.GetIntentConfigurationSignature(mainSigner, common.Address{}, payloads, nil, nil, "", nil)
	require.NoError(t, err)
	fmt.Printf("--- Intent Config Signature (for all payloads) ---\n%s\n", common.Bytes2Hex(intentConfigSig))

	// Create and execute guest bundles for each payload separately
	for i, payload := range payloads {
		signedExecdata, err := contracts.V3.WalletStage1Module.Encode("execute", payload.Encode(common.Address{}), intentConfigSig)
		assert.NoError(t, err)

		// Log encoded payload sent to contract
		fmt.Printf("--- Executing Payload %d ---\n", i+1)
		fmt.Printf("Encoded Payload (to execute): %s\n", common.Bytes2Hex(payload.Encode(common.Address{})))

		guestAddress := testChain.V3SequenceContext().GuestModuleAddress

		var guestBundle []v3.Call
		if i == 0 {
			guestBundle = []v3.Call{
				{
					To:   walletFactoryAddress,
					Data: walletDeployData,
				},
				{
					To:   wallet.Address(),
					Data: signedExecdata,
				},
			}
		} else {
			guestBundle = []v3.Call{
				{
					To:   wallet.Address(),
					Data: signedExecdata,
				},
			}
		}

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
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Log receipt status
		// spew.Dump(receipt)
		fmt.Printf("Receipt Status Payload %d: %d\n", i+1, receipt.Status)

		// Check the value for each contract
		var expectedValue string
		var contractAddress common.Address
		if i == 0 {
			expectedValue = "2255"
			contractAddress = callmockContract.Address
		} else {
			expectedValue = "6655"
			contractAddress = callmockContract.Address
		}

		ret, err := testutil.ContractQuery(testChain.Provider, contractAddress, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, expectedValue, ret[0])

		// Assert sequence.WaitForMetaTxn is able to find the metaTxnID for each operation
		var opPayload *v3.CallsPayload
		if i == 0 {
			opPayload = &opPayload1
		} else {
			opPayload = &opPayload2
		}

		// Log expected MetaTxnID
		expectedMetaTxnID := sequence.MetaTxnID(opPayload.Digest().Hash.Hex()[2:])
		fmt.Printf("Expected MetaTxnID Payload %d: %s\n", i+1, expectedMetaTxnID)

		result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, expectedMetaTxnID)
		assert.NoError(t, err)
		assert.True(t, result.Status == sequence.MetaTxnExecuted)
	}

	// Wallet should be deployed after executing both transactions
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
		payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
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
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, nil)
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

		assert.Equal(t, address, common.HexToAddress("0x8577dFb93fE58cc8EE90DEA522555Fdf01Fd7429"))
	})

	t.Run("multiple operations", func(t *testing.T) {
		// Create multiple operations matching TypeScript test
		payload1 := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
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
		payload2 := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
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
		)

		// Create intent configuration
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload1, &payload2}, nil)
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

		assert.Equal(t, address, common.HexToAddress("0xBd820eD5b1E969eD6509E8EdE687DfC4c714438F"))
	})
}

func TestIntentConfigurationAddress_RealWorldExample(t *testing.T) {
	// Main signer matching TypeScript test
	mainSigner := common.HexToAddress("0x8456195dd0793c621c7f9245edF0fEf85b1B879C")

	// Context for the counterfactual address
	context := sequence.V3SequenceContext()
	context.FactoryAddress = common.HexToAddress("0xBd0F8abD58B4449B39C57Ac9D5C67433239aC447")
	context.MainModuleAddress = common.HexToAddress("0x53bA242E7C2501839DF2972c75075dc693176Cd0")

	// Create first operation on Arbitrum (chainId: 42161)
	payload1 := v3.NewCallsPayload(common.Address{}, big.NewInt(42161), []v3.Call{
		{
			To:              common.HexToAddress("0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae"),
			Value:           big.NewInt(16618237),
			Data:            common.FromHex("0xa6010a660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000005801784da62343d604885e1181a759647447f13330cc2b8c925cda864b1ac1ce8fc000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e58310000000000000000000000008456195dd0793c621c7f9245edf0fef85b1b879c00000000000000000000000000000000000000000000000000000000000073ac000000000000000000000000000000000000000000000000000000000000210500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d737461726761746556324275730000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000086c6966692d61706900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000111111125421ca6dc452d289314280a0f8842a65000000000000000000000000111111125421ca6dc452d289314280a0f8842a650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e583100000000000000000000000000000000000000000000000000000f42af44fef500000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000020807ed2379000000000000000000000000de9e4fe32b049f821c7f3e9802381aa470ffca73000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e5831000000000000000000000000de9e4fe32b049f821c7f3e9802381aa470ffca730000000000000000000000001231deb6f5749ef6ce6943a275a1d3e7486f4eae00000000000000000000000000000000000000000000000000000f42af44fef500000000000000000000000000000000000000000000000000000000000075320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000ba00000000000000000000000000000000000000000000000000000000009c4160de632c3a214d5f14c1d8ddf0b92f8bcd188fee4500242668dfaa000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007532000000000000000000000000111111125421ca6dc452d289314280a0f8842a650000000000002a94d114000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000077368def90800000000000000000000000000000000000000000000000000000000000000000000000000000000000000008456195dd0793c621c7f9245edf0fef85b1b879c00000000000000000000000000000000000000000000000000000000000075e80000000000000000000000008456195dd0793c621c7f9245edf0fef85b1b879c0000000000000000000000000000000000000000000000000000000000007532000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
	}, big.NewInt(0), big.NewInt(0))

	// Create second operation on Base (chainId: 8453)
	payload2 := v3.NewCallsPayload(common.Address{}, big.NewInt(8453), []v3.Call{
		{
			To:              common.HexToAddress("0x833589fcd6edb6e08f4c7c32d4f71b54bda02913"),
			Value:           big.NewInt(0),
			Data:            common.FromHex("0xa9059cbb000000000000000000000000d8da6bf26964af9d7eed9e03e53415d37aa960450000000000000000000000000000000000000000000000000000000000007530"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
	}, big.NewInt(0), big.NewInt(0))

	// Create intent configuration
	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload1, &payload2}, nil)
	require.NoError(t, err)

	// Calculate image hash
	imageHash := config.ImageHash()

	// Calculate counterfactual address
	address, err := sequence.AddressFromImageHash(imageHash, context)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("\nReal World Example Test:\n")
	fmt.Printf("Address: %s\n", address.Hex())

	// The address should be deterministic based on the configuration
	assert.Equal(t, address, common.HexToAddress("0x5bD7F0269F4AA805F5a13B3104D596c151d8eC76"))
}

func TestCreateIntentCallsPayloadDigest(t *testing.T) {
	// Create an intent operation with a single call
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x0000000000000000000000000000000000000000"),
			Value:           big.NewInt(0),
			Data:            common.FromHex("0x1234"),
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorIgnore,
		},
	}, big.NewInt(0), big.NewInt(0))

	// Get the digest
	digest := payload.Digest()
	require.NotNil(t, digest)
	require.NotNil(t, digest.Hash)

	// Verify the digest hash is not empty
	require.NotEqual(t, common.Hash{}, digest.Hash)

	// Print the digest hash for debugging
	fmt.Printf("Digest Hash: %s\n", digest.Hash.Hex())
}

func TestGetAnypayExecutionInfoHash(t *testing.T) {
	t.Run("SingleInfo", func(t *testing.T) {
		originTokenAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
		amountVal := big.NewInt(100)
		originChainIdVal := big.NewInt(1)
		destinationChainIdVal := big.NewInt(10)
		attestationAddrVal := common.HexToAddress("0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa")

		lifiInfos := []sequence.AnypayExecutionInfo{
			{
				OriginToken:        originTokenAddr,
				Amount:             amountVal,
				OriginChainId:      originChainIdVal,
				DestinationChainId: destinationChainIdVal,
			},
		}

		expectedHash := common.HexToHash("0x21872bd6b64711c4a5aecba95829c612f0b50c63f1a26991c2f76cf4a754aede")
		actualHashBytes, err := sequence.GetAnypayExecutionInfoHash(lifiInfos, attestationAddrVal)
		require.NoError(t, err)

		assert.Equal(t, expectedHash, common.Hash(actualHashBytes), "SingleInfo hash mismatch")
	})

	t.Run("MultipleInfo", func(t *testing.T) {
		lifiInfos := []sequence.AnypayExecutionInfo{
			{
				OriginToken:        common.HexToAddress("0x1111111111111111111111111111111111111111"),
				Amount:             big.NewInt(100),
				OriginChainId:      big.NewInt(1),
				DestinationChainId: big.NewInt(10),
			},
			{
				OriginToken:        common.HexToAddress("0x2222222222222222222222222222222222222222"),
				Amount:             big.NewInt(200),
				OriginChainId:      big.NewInt(137),
				DestinationChainId: big.NewInt(42161),
			},
		}
		attestationAddrVal := common.HexToAddress("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB")

		expectedHash := common.HexToHash("0xd18e54455db64ba31b9f9a447e181f83977cb70b136228d64ac85d64a6aefe71")
		actualHashBytes, err := sequence.GetAnypayExecutionInfoHash(lifiInfos, attestationAddrVal)
		require.NoError(t, err)

		assert.Equal(t, expectedHash, common.Hash(actualHashBytes), "MultipleInfo hash mismatch")
	})
}

func TestGetAnypayRelayInfoHash(t *testing.T) {
	t.Run("SingleInfo", func(t *testing.T) {
		relayInfos := []sequence.AnypayRelayInfo{
			{
				RequestId:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
				Signature:          common.FromHex("abcd"),
				NonEVMReceiver:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				ReceivingAssetId:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
				SendingAssetId:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
				Receiver:           common.HexToAddress("0x6666666666666666666666666666666666666666"),
				DestinationChainId: big.NewInt(137),
				MinAmount:          big.NewInt(1000),
				Target:             common.HexToAddress("0x7777777777777777777777777777777777777777"),
			},
		}
		attestationAddrVal := common.HexToAddress("0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa")

		expectedHash := common.HexToHash("0x34b1669f0dccfb1e185ee9012c92a17c8548dc504d7a3dc0fedf08522c8c5a63")
		actualHashBytes, err := sequence.GetAnypayRelayInfoHash(relayInfos, attestationAddrVal)
		require.NoError(t, err)

		assert.Equal(t, expectedHash, common.Hash(actualHashBytes), "SingleInfo hash mismatch")
	})

	t.Run("MultipleInfo", func(t *testing.T) {
		relayInfos := []sequence.AnypayRelayInfo{
			{
				RequestId:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
				Signature:          common.FromHex("abcd"),
				NonEVMReceiver:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				ReceivingAssetId:   common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
				SendingAssetId:     common.HexToAddress("0x5555555555555555555555555555555555555555"),
				Receiver:           common.HexToAddress("0x6666666666666666666666666666666666666666"),
				DestinationChainId: big.NewInt(137),
				MinAmount:          big.NewInt(1000),
				Target:             common.HexToAddress("0x7777777777777777777777777777777777777777"),
			},
			{
				RequestId:          common.HexToHash("0x1000000000000000000000000000000000000000000000000000000000000001"),
				Signature:          common.FromHex("dcba"),
				NonEVMReceiver:     common.HexToHash("0x1000000000000000000000000000000000000000000000000000000000000002"),
				ReceivingAssetId:   common.HexToHash("0x1000000000000000000000000000000000000000000000000000000000000003"),
				SendingAssetId:     common.HexToAddress("0x8888888888888888888888888888888888888888"),
				Receiver:           common.HexToAddress("0x9999999999999999999999999999999999999999"),
				DestinationChainId: big.NewInt(42161),
				MinAmount:          big.NewInt(2000),
				Target:             common.HexToAddress("0x7777777777777777777777777777777777777777"),
			},
		}
		attestationAddrVal := common.HexToAddress("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB")

		expectedHash := common.HexToHash("0xe36f2474265f43cea2e68e83112443c9dc35b844f7039d96450adcf1acd6a7e8")
		actualHashBytes, err := sequence.GetAnypayRelayInfoHash(relayInfos, attestationAddrVal)
		require.NoError(t, err)

		assert.Equal(t, expectedHash, common.Hash(actualHashBytes), "MultipleInfo hash mismatch")
	})
}

// ecrecoverForTest is a helper function to recover an address from a message hash and signature.
// It assumes the input signature's V value has been incremented by 27 twice (e.g., V = original_V_from_crypto.Sign + 54)
// and normalizes it to original_V (0 or 1) for crypto.SigToPub.
func ecrecoverForTest(hash []byte, signature []byte) (common.Address, error) {
	if len(signature) != crypto.SignatureLength { // crypto.SignatureLength is 65
		return common.Address{}, fmt.Errorf("invalid signature length %d, expected %d", len(signature), crypto.SignatureLength)
	}

	// Create a mutable copy of the signature to adjust V.
	// crypto.SigToPub expects V to be 0 or 1.
	// This function assumes the input signature's V has been incremented by 27 twice (total +54).
	var sigToPub [crypto.SignatureLength]byte // Use array as in user's example for fixed-size operations
	copy(sigToPub[:], signature)
	sigToPub[crypto.SignatureLength-1] -= (2 * 27) // Adjust V from (original_V + 54) down to original_V (0 or 1)

	pubkey, err := crypto.SigToPub(hash, sigToPub[:]) // Pass as slice
	if err != nil {
		return common.Address{}, fmt.Errorf("unable to recover public key: %w", err)
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}

func TestCreateAnypayLifiAttestation(t *testing.T) {
	// 1. Setup
	attestationSignerWallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err, "Failed to create attestation signer wallet")

	// Create a simple CallsPayload
	call := v3.Call{
		To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Value:           big.NewInt(0),
		Data:            []byte("test_payload_data_for_lifi_attestation"),
		GasLimit:        big.NewInt(100000), // Example gas limit
		DelegateCall:    false,
		OnlyFallback:    false,
		BehaviorOnError: v3.BehaviorOnErrorRevert,
	}
	payloadAddress := common.HexToAddress("0xPayloadAddressForTest123456789012345")
	// Using chain ID 1 for simplicity in this test context, can be any valid chain ID
	payload := v3.NewCallsPayload(payloadAddress, big.NewInt(1), []v3.Call{call}, big.NewInt(0), big.NewInt(0))

	lifiInfos := []sequence.AnypayExecutionInfo{
		{
			OriginToken:        common.HexToAddress("0xOriginToken0000000000000000000000000000"),
			Amount:             big.NewInt(1000),
			OriginChainId:      big.NewInt(1),  // Ethereum Mainnet
			DestinationChainId: big.NewInt(10), // Optimism
		},
		{
			OriginToken:        common.HexToAddress("0xOriginToken2222222222222222222222222222"),
			Amount:             big.NewInt(500),
			OriginChainId:      big.NewInt(137),   // Polygon
			DestinationChainId: big.NewInt(42161), // Arbitrum
		},
	}

	// 2. Call CreateAnypayLifiAttestation
	encodedAttestation, err := sequence.CreateAnypayLifiAttestation(attestationSignerWallet, &payload, lifiInfos)
	require.NoError(t, err, "CreateAnypayLifiAttestation should not return an error for valid inputs")
	require.NotNil(t, encodedAttestation, "Encoded attestation should not be nil")
	require.NotEmpty(t, encodedAttestation, "Encoded attestation should not be empty")

	// 3. Decode the result (abi.encode(AnypayExecutionInfo[] memory, bytes memory))
	// Define ABI type components for the AnypayExecutionInfo struct
	AnypayExecutionInfoComponents := []abi.ArgumentMarshaling{
		{Name: "originToken", Type: "address"},
		{Name: "amount", Type: "uint256"},
		{Name: "originChainId", Type: "uint256"},
		{Name: "destinationChainId", Type: "uint256"},
	}
	// Define ABI type for a list of AnypayExecutionInfo structs (AnypayExecutionInfo[])
	lifiInfoArrayType, err := abi.NewType("tuple[]", "AnypayExecutionInfo[]", AnypayExecutionInfoComponents)
	require.NoError(t, err, "Failed to create AnypayExecutionInfo[] ABI type")
	bytesType, err := abi.NewType("bytes", "", nil)
	require.NoError(t, err, "Failed to create bytes ABI type")

	// Define the arguments for ABI decoding
	argumentsToUnpack := abi.Arguments{{Type: lifiInfoArrayType}, {Type: bytesType}}
	unpackedData, err := argumentsToUnpack.Unpack(encodedAttestation)
	require.NoError(t, err, "Failed to unpack encoded attestation")
	require.Len(t, unpackedData, 2, "Unpacked data should contain two elements: lifiInfos and signature")

	// 4. Verify decoded signature
	decodedSignature, ok := unpackedData[1].([]byte)
	require.True(t, ok, "Failed to assert unpackedData[1] as []byte for the signature")
	require.NotEmpty(t, decodedSignature, "Decoded signature should not be empty")
	require.Len(t, decodedSignature, 65, "Decoded signature should be 65 bytes long (R + S + V)")

	// Verify the ECDSA signature
	// 1. Determine the hash that was actually signed by crypto.Sign inside SignData.
	originalPayloadDigestHash := payload.Digest().Hash
	hashSignedByCryptoSign := crypto.Keccak256Hash(originalPayloadDigestHash[:])

	// 2. Use the ecrecoverForTest helper for signature verification.
	recoveredAddress, err := ecrecoverForTest(hashSignedByCryptoSign.Bytes(), decodedSignature)
	require.NoError(t, err, "Failed to recover address using ecrecoverForTest")
	assert.Equal(t, attestationSignerWallet.Address(), recoveredAddress, "Recovered address does not match the attestation signer's address")
}
