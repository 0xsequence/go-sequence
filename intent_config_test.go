package sequence_test

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"

	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestCreateIntentBundle_Valid(t *testing.T) {
	// Create a valid payload
	payload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	bundle, err := sequence.CreateIntentBundle(payload, false)
	require.NoError(t, err)
	require.NotNil(t, bundle)

	spew.Dump(bundle)
}

// TestCreateIntentDigestLeaves_Valid creates a valid payload and computes the intent digest
func TestCreateIntentDigestLeaves_Valid(t *testing.T) {
	// Create valid payloads with required fields
	payload1 := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	payload2 := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	payload3 := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		batches := []v3.CallsPayload{payload1}

		bundle, err := sequence.CreateIntentBundle(payload1, false)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches, testChain.ChainID(), false)
		require.NoError(t, err)
		require.NotNil(t, tree, "expected a tree")

		// Dump the tree
		spew.Dump((*tree))

		// Type assert to the concrete type
		anyAddressLeaf, ok := (*tree).(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "tree should be a WalletConfigTreeAnyAddressSubdigestLeaf")

		digest := bundle.Digest()
		require.Equal(t, digest.Hash, anyAddressLeaf.Digest.Hash, "digests do not match")
	})

	t.Run("Two batches", func(t *testing.T) {
		// Use two payloads
		batches := []v3.CallsPayload{payload1, payload2}

		bundle1, err := sequence.CreateIntentBundle(payload1, false)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(payload2, false)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches, testChain.ChainID(), false)
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

		digest1 := bundle1.Digest()
		digest2 := bundle2.Digest()
		require.Equal(t, digest1.Hash, leftLeaf.Digest.Hash, "left leaf digest does not match")
		require.Equal(t, digest2.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		// Use three payloads
		batches := []v3.CallsPayload{payload1, payload2, payload3}

		bundle1, err := sequence.CreateIntentBundle(payload1, false)
		require.NoError(t, err)
		bundle2, err := sequence.CreateIntentBundle(payload2, false)
		require.NoError(t, err)
		bundle3, err := sequence.CreateIntentBundle(payload3, false)
		require.NoError(t, err)

		tree, err := sequence.CreateIntentDigestTree(batches, testChain.ChainID(), false)
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

		digest1 := bundle1.Digest()
		digest2 := bundle2.Digest()
		digest3 := bundle3.Digest()
		require.Equal(t, digest1.Hash, leftLeftLeaf.Digest.Hash, "left left leaf digest does not match")
		require.Equal(t, digest2.Hash, leftRightLeaf.Digest.Hash, "left right leaf digest does not match")
		require.Equal(t, digest3.Hash, rightLeaf.Digest.Hash, "right leaf digest does not match")
	})
}

func TestCreateIntentConfiguration_Valid(t *testing.T) {
	// Create a valid payload
	payload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	batches := []v3.CallsPayload{payload}

	// Use a valid main signer address.
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	config, err := sequence.CreateIntentConfiguration(mainSigner, batches, testChain.ChainID(), false)
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

	payload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		batches := []v3.CallsPayload{payload}

		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), batches, testChain.ChainID(), false)
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), batches, testChain.ChainID(), false)
		require.NoError(t, err)

		fmt.Println("==> signature", common.Bytes2Hex(signature))

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
		spew.Dump(recoveredConfig)

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

		payload1 := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

		payload2 := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []v3.CallsPayload{payload1}, testChain.ChainID(), false)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []v3.CallsPayload{payload2}, testChain.ChainID(), true)
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Use the payload directly
		batches := []v3.CallsPayload{payload}

		// Create the same payload signature twice
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), batches, testChain.ChainID(), true)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), batches, testChain.ChainID(), true)
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
	multiCallPayload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	batches := []v3.CallsPayload{multiCallPayload}

	// Create a signature
	sig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), batches, testChain.ChainID(), true)
	require.NoError(t, err)

	// Convert the full signature into a hex string.
	sigHex := common.Bytes2Hex(sig)

	// Create the bundle from the payload
	bundle, err := sequence.CreateIntentBundle(multiCallPayload, false)
	require.NoError(t, err)

	// Compute the digest of the bundle
	bundleDigest := bundle.Digest()
	require.NoError(t, err)

	// Expect that the signature (in hex) contains the substrings of the bundle's digest.
	bundleDigestHex := bundleDigest.Hash.Hex()
	require.Contains(t, sigHex, bundleDigestHex[2:], "signature should contain transaction bundle digest")
}

func TestV3IntentConfigWalletDeployment(t *testing.T) {
	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata1, err := callmockContract.Encode("setRevertFlag", false)
	require.NoError(t, err)

	calldata2, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
	require.NoError(t, err)

	log.Println("==> callmockContract", callmockContract.Address.Hex())
	log.Println("==> calldata1", common.Bytes2Hex(calldata1))
	log.Println("==> calldata2", common.Bytes2Hex(calldata2))

	// Create a v3 payload for mock contract call
	mockPayloadCall := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
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

	// Create batches from the mockPayloadCall
	// This is needed since the `CreateIntentBundle` function expects an array of `CallsPayload`.
	mockPayloadInitialBatch := []v3.CallsPayload{mockPayloadCall}
	mockBundles := []v3.CallsPayload{}
	metaTxnIds := []string{}

	// Create the processed bundles for the initial batch
	for _, batch := range mockPayloadInitialBatch {
		// Create the processed bundle for the payload
		mockBundle, err := sequence.CreateIntentBundle(batch, false)
		require.NoError(t, err)
		mockBundles = append(mockBundles, mockBundle)
	}

	// Create the digest for each call in the payload
	for _, bundle := range mockBundles {
		for _, call := range bundle.Calls {
			metaTxnIds = append(metaTxnIds, v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{call}, big.NewInt(0), big.NewInt(0)).Digest().Hash.Hex())
		}
	}

	spew.Dump(mockBundles)
	spew.Dump(metaTxnIds)

	for _, id := range metaTxnIds {
		fmt.Println("==> metaTxnId", id)
	}

	// Decode the payload
	// CallsPayload, err := v3.Decode(encodedPayload)
	// require.NoError(t, err)

	// spew.Dump(CallsPayload)

	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V3DummySequenceWalletWithIntentConfig(1, mockBundles)
	require.NoError(t, err)
	require.NotNil(t, wallet)

	// Get the first payload hash
	firstPayloadHash := mockBundles[0].Digest()
	require.NoError(t, err)

	log.Println("==> firstPayloadHash", firstPayloadHash.Hash.Hex())

	// Get the main signer
	signers := wallet.GetWalletConfig().Signers()
	var mainSigner common.Address
	for addr := range signers {
		mainSigner = addr
		break
	}
	require.NotNil(t, mainSigner)

	// Wallet deployment data
	_, walletFactoryAddress, walletDeployData, err := sequence.EncodeWalletDeployment(wallet.GetWalletConfig(), wallet.GetWalletContext())
	require.NoError(t, err)

	// Generate a configuration signature for the batch.
	configSig, err := sequence.GetIntentConfigurationSignature(mainSigner, mockBundles, testChain.ChainID(), false)
	require.NoError(t, err)

	// Print the signature in hex
	fmt.Println("==> configSig", common.Bytes2Hex(configSig))

	// Verify the signature contains the any address digest
	require.Contains(t, common.Bytes2Hex(configSig), firstPayloadHash.Hash.Hex()[2:], "signature should contain the any address digest")

	// Get the mock payload
	// Create the intent configuration using the batched transactions.
	config, err := sequence.CreateIntentConfiguration(mainSigner, mockBundles, testChain.ChainID(), false)
	require.NoError(t, err)

	spew.Dump(config)

	// Get the image of the config
	configImageHash := config.ImageHash()
	fmt.Println("==> configImageHash", configImageHash.Hex())

	// Get the hash of the first payload
	signedPayloadHash := mockBundles[0].Digest()
	require.NoError(t, err)

	spew.Dump(mockBundles[0])

	fmt.Println("==> signedPayloadHash", signedPayloadHash.Hash.Hex())

	encodedData := mockBundles[0].Encode(common.Address{})
	require.NoError(t, err)

	signedExecdata, err := contracts.V3.WalletStage1Module.Encode("execute", encodedData, configSig)
	require.NoError(t, err)

	fmt.Println("==> signedExecdata", common.Bytes2Hex(signedExecdata))

	spew.Dump(signedExecdata)

	guestBundle := []v3.Call{
		{
			To:   walletFactoryAddress,
			Data: walletDeployData,
		},
	}

	guestBundle2 := []v3.Call{

		{
			To:   wallet.Address(),
			Data: signedExecdata,
		},
	}

	guestPayload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), guestBundle, big.NewInt(0), big.NewInt(0))

	guestAddress := testChain.V3SequenceContext().GuestModuleAddress
	execdata := guestPayload.Encode(guestAddress)
	require.NoError(t, err)

	// Relay the txn manually, directly to the guest module
	sender := testChain.GetRelayerWallet()
	ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:       &guestAddress,
		Data:     execdata,
		GasLimit: 1000000, // TODO: compute gas limit
	})
	require.NoError(t, err)

	signedTx, err := sender.SignTx(ntx, testChain.ChainID())
	require.NoError(t, err)

	_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)

	receipt, err := waitReceipt(context.Background())
	require.NoError(t, err)
	spew.Dump(receipt)
	require.True(t, receipt.Status == types.ReceiptStatusSuccessful)

	guestBundle2Payload := v3.ConstructCallsPayload(common.Address{}, testChain.ChainID(), guestBundle2, big.NewInt(0), big.NewInt(0))

	spew.Dump(guestBundle2Payload)

	guestBundle2PayloadEncoded := guestBundle2Payload.Encode(guestAddress)
	require.NoError(t, err)

	log.Println("==> guestBundle2PayloadEncoded", common.Bytes2Hex(guestBundle2PayloadEncoded))

	// Relay the second txn manually, directly to the guest module
	ntx2, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:       &guestAddress,
		Data:     guestBundle2PayloadEncoded,
		GasLimit: 1000000, // TODO: compute gas limit
	})
	require.NoError(t, err)

	signedTx2, err := sender.SignTx(ntx2, testChain.ChainID())
	require.NoError(t, err)

	_, waitReceipt2, err := sender.SendTransaction(context.Background(), signedTx2)
	require.NoError(t, err)

	// After obtaining receipt2
	callFailedTopic := crypto.Keccak256Hash([]byte("CallFailed(bytes32,uint256,bytes)"))
	callSuccessTopic := crypto.Keccak256Hash([]byte("CallSuccess(bytes32,uint256)"))

	receipt2, err := waitReceipt2(context.Background())
	require.NoError(t, err)
	spew.Dump(receipt2)
	require.True(t, receipt2.Status == types.ReceiptStatusSuccessful)

	for _, log := range receipt2.Logs {
		if len(log.Topics) > 0 {
			switch log.Topics[0] {
			case callFailedTopic:
				fmt.Println("CallFailed event found with return data:", common.Bytes2Hex(log.Data))
			case callSuccessTopic:
				if len(log.Data) < 32 {
					fmt.Println("Log data too short to contain opHash")
					continue
				}
				opHash := common.BytesToHash(log.Data[0:32])
				fmt.Println("==> CallSuccess data", common.Bytes2Hex(log.Data))
				fmt.Println("CallSuccess event found with opHash:", opHash.Hex())
			}
		}
	}

	metaTxnID := sequence.MetaTxnID(metaTxnIds[0])

	// Check the value
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	require.NoError(t, err)
	require.Len(t, ret, 1)
	require.Equal(t, "2255", ret[0])

	// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
	result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	require.NoError(t, err)
	require.True(t, result.Status == sequence.MetaTxnExecuted)

	// Wallet should be deployed now
	isDeployed, err := wallet.IsDeployed()
	require.NoError(t, err)
	require.True(t, isDeployed)
}
