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
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"
	"github.com/davecgh/go-spew/spew"

	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/receipts"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func findSapientSignerLeaf(tree v3.WalletConfigTree, address common.Address) *v3.WalletConfigTreeSapientSignerLeaf {
	switch node := tree.(type) {
	case *v3.WalletConfigTreeSapientSignerLeaf:
		if node.Address == address {
			return node
		}
	case *v3.WalletConfigTreeNode:
		if leaf := findSapientSignerLeaf(node.Left, address); leaf != nil {
			return leaf
		}
		return findSapientSignerLeaf(node.Right, address)
	case *v3.WalletConfigTreeNestedLeaf:
		return findSapientSignerLeaf(node.Tree, address)
	}

	return nil
}

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
		require.Equal(t, digest.Hash, anyAddressLeaf.Digest, "digests do not match")
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
		require.Equal(t, digest1.Hash, leftLeaf.Digest, "left leaf digest does not match")
		require.Equal(t, digest2.Hash, rightLeaf.Digest, "right leaf digest does not match")
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
		require.Equal(t, digest1.Hash, leftLeftLeaf.Digest, "left left leaf digest does not match")
		require.Equal(t, digest2.Hash, leftRightLeaf.Digest, "left right leaf digest does not match")
		require.Equal(t, digest3.Hash, rightLeaf.Digest, "right leaf digest does not match")
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
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1})
		require.NoError(t, err)
		require.NotNil(t, tree)

		// spew.Dump(tree)

		// Type assert to the concrete type
		nodeTree, ok := (*tree).(*v3.WalletConfigTreeNode)
		require.True(t, ok, "tree should be a WalletConfigTreeNode")

		addressLeaf, ok := nodeTree.Left.(*v3.WalletConfigTreeAddressLeaf)
		require.True(t, ok, "left leaf should be WalletConfigTreeAddressLeaf")
		require.Equal(t, addressLeaf.Address, common.Address{}, "address leaf should be the main signer")

		_, ok = nodeTree.Right.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
		require.True(t, ok, "right leaf should be WalletConfigTreeAnyAddressSubdigestLeaf")

		// Get the digest of the payload
		bundle, err := sequence.CreateAnyAddressSubdigestTree([]*v3.CallsPayload{&payload1})
		require.NoError(t, err)
		require.NotNil(t, bundle)
	})

	t.Run("Two batches", func(t *testing.T) {
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1, &payload2})
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

		require.Equal(t, payload1.Digest().Hash, anyAddressLeaf2.Digest, "digests do not match")
		require.Equal(t, payload2.Digest().Hash, anyAddressLeaf.Digest, "digests do not match")
	})

	t.Run("Three batches", func(t *testing.T) {
		tree, err := sequence.CreateIntentTree(common.Address{}, []*v3.CallsPayload{&payload1, &payload2, &payload3})
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

		require.Equal(t, payload1.Digest().Hash, anyAddressLeaf3.Digest, "digests do not match")
		require.Equal(t, payload2.Digest().Hash, anyAddressLeaf2.Digest, "digests do not match")
		require.Equal(t, payload3.Digest().Hash, anyAddressLeaf.Digest, "digests do not match")
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

	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0)
	require.NoError(t, err)
	require.NotNil(t, config)
}

func TestCreateIntentConfigurationNilOptionsIgnored(t *testing.T) {
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

	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")

	// Legacy callers passed nil for the removed positional leaf params; those nils now
	// arrive as nil options and must mean "no option", not panic.
	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0, nil, nil)
	require.NoError(t, err)

	plainConfig, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0)
	require.NoError(t, err)
	require.Equal(t, plainConfig.ImageHash().Hash, config.ImageHash().Hash)
}

func TestCreateIntentConfigurationWithTimedRefundSapient(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	timedRefundSigner := common.HexToAddress("0x3333333333333333333333333333333333333333")
	destination := common.HexToAddress("0x4444444444444444444444444444444444444444")

	timedRefundImageHash, err := sequence.TimedRefundSapientImageHash(destination, 1_750_000_000)
	require.NoError(t, err)
	timedRefundLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    timedRefundSigner,
		ImageHash_: timedRefundImageHash,
	}

	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0, sequence.WithSapientSigner(timedRefundLeaf))
	require.NoError(t, err)
	require.NotNil(t, config)

	// Solidity reference (trails-contracts/src/autoRecovery/TimedRefundSapient.sol):
	// keccak256(abi.encode("timed-refund", destination, uint256(1750000000)))
	expectedSapientImageHash := common.HexToHash("0x577e11f2280512fff4541fc08cc7eb98357bdcff482db5634db7327e3c97ba58")
	sapientLeaf := findSapientSignerLeaf(config.Tree, timedRefundSigner)
	require.NotNil(t, sapientLeaf)
	require.Equal(t, expectedSapientImageHash, sapientLeaf.ImageHash_.Hash)
	preimage, ok := sapientLeaf.ImageHash_.Preimage.(*sequence.TimedRefundSapientImageHashPreimage)
	require.True(t, ok)
	require.Equal(t, destination, preimage.Destination)
	require.Equal(t, uint64(1_750_000_000), preimage.UnlockTimestamp)
	require.Equal(t, expectedSapientImageHash, preimage.ImageHash().Hash)

	plainConfig, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0)
	require.NoError(t, err)
	require.NotEqual(t, plainConfig.ImageHash().Hash, config.ImageHash().Hash)

	signature, err := sequence.BuildIntentConfigurationSignature(config, nil)
	require.NoError(t, err)
	require.NotEmpty(t, signature)

	sig, err := v3.Core.DecodeSignature(signature)
	require.NoError(t, err)

	recoveredConfig, _, err := sig.Recover(context.Background(), payload, nil)
	require.NoError(t, err)
	require.Equal(t, config.ImageHash().Hash, recoveredConfig.ImageHash().Hash)

	plainSignature, err := sequence.GetIntentConfigurationSignature(mainSigner, []*v3.CallsPayload{&payload}, 0, nil)
	require.NoError(t, err)
	require.NotEqual(t, plainSignature, signature)
}

// With gateLeaf nil (the default/legacy case), the tree must keep the exact flat
// shape it had before this parameter existed: Node(mainSignerLeaf, Node(subdigestLeaf,
// additionalLeaf)) — no extra nesting — so already-derived counterfactual addresses do
// not change.
func TestCreateIntentConfigurationGateLeafNilUnchanged(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	sapientSignerLeafNode := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    common.HexToAddress("0x3333333333333333333333333333333333333333"),
		ImageHash_: core.ImageHash{Hash: common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111")},
	}

	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0, sequence.WithSapientSigner(sapientSignerLeafNode))
	require.NoError(t, err)

	top, ok := config.Tree.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	ownerLeaf, ok := top.Left.(*v3.WalletConfigTreeAddressLeaf)
	require.True(t, ok)
	require.Equal(t, mainSigner, ownerLeaf.Address)

	rest, ok := top.Right.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	_, subdigestOk := rest.Left.(*v3.WalletConfigTreeAnyAddressSubdigestLeaf)
	require.True(t, subdigestOk)
	require.Same(t, sapientSignerLeafNode, rest.Right)
}

func TestCreateIntentConfigurationWithGateLeaf(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	peerSigner := common.HexToAddress("0x72030E1dbf0a847196ae62EA3ee84BD7ce99D6c1")
	peerSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    peerSigner,
		ImageHash_: core.ImageHash{Hash: common.BigToHash(big.NewInt(1))},
	}

	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0, sequence.WithGate(peerSignerLeaf))
	require.NoError(t, err)
	require.NotNil(t, config)

	sapientLeaf := findSapientSignerLeaf(config.Tree, peerSigner)
	require.NotNil(t, sapientLeaf)

	// The gate must change the counterfactual address relative to an ungated config.
	plainConfig, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0)
	require.NoError(t, err)
	require.NotEqual(t, plainConfig.ImageHash().Hash, config.ImageHash().Hash)

	// A signature that does not include the peer signer's co-signature must not meet the
	// gate's threshold, even though the any-address-subdigest leaf matches the payload
	// (which alone would satisfy an ungated config's threshold).
	signatureWithoutPeerSig, err := sequence.BuildIntentConfigurationSignature(config, nil)
	require.NoError(t, err)

	sigWithoutPeerSig, err := v3.Core.DecodeSignature(signatureWithoutPeerSig)
	require.NoError(t, err)

	recoveredConfig, weight, err := sigWithoutPeerSig.Recover(context.Background(), payload, nil)
	require.NoError(t, err)
	require.Equal(t, config.ImageHash().Hash, recoveredConfig.ImageHash().Hash)
	require.Truef(t, weight.Cmp(big.NewInt(int64(config.Threshold()))) < 0,
		"recovered weight %v must not meet threshold %v without the peer signer's co-signature", weight, config.Threshold())

	// Including the peer signer's co-signature must produce a different signature
	// encoding than omitting it, proving the leaf is actually wired into the built signature.
	signatureWithPeerSig, err := sequence.BuildIntentConfigurationSignature(config, []*core.SignerSignature{
		{
			Signer:    core.SapientSigner(peerSigner, peerSignerLeaf.ImageHash_.Hash),
			Signature: []byte{},
			Type:      core.SignerSignatureTypeSapientCompact,
		},
	})
	require.NoError(t, err)
	require.NotEqual(t, signatureWithoutPeerSig, signatureWithPeerSig)
}

// A typed-nil gate leaf passes the interface nil check, so signerLeaf must reject it
// before dereferencing the concrete pointer.
func TestCreateIntentConfigurationGateTypedNilRejected(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))
	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")

	var gate *v3.WalletConfigTreeSapientSignerLeaf
	_, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
		sequence.WithGate(gate))
	require.ErrorContains(t, err, "nil leaf")
}

// A gated signer leaf sharing the gate's identity would satisfy both sides of the outer
// threshold with one signature, so the config must be rejected at construction.
func TestCreateIntentConfigurationGateDuplicateSapientRejected(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	gateSigner := common.HexToAddress("0x72030E1dbf0a847196ae62EA3ee84BD7ce99D6c1")
	gateImageHash := core.ImageHash{Hash: common.BigToHash(big.NewInt(1))}
	gateLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    gateSigner,
		ImageHash_: gateImageHash,
	}

	t.Run("identical leaf in both roles is rejected", func(t *testing.T) {
		_, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
			sequence.WithGate(gateLeaf), sequence.WithSapientSigner(gateLeaf))
		require.ErrorContains(t, err, "gate signer must not appear among gated leaves")
	})

	t.Run("same signer with different weight is rejected", func(t *testing.T) {
		heavierLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     2,
			Address:    gateSigner,
			ImageHash_: gateImageHash,
		}
		_, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
			sequence.WithGate(gateLeaf), sequence.WithSapientSigner(heavierLeaf))
		require.ErrorContains(t, err, "gate signer must not appear among gated leaves")
	})

	t.Run("payload-matching leaf as gate is rejected", func(t *testing.T) {
		subdigestGate := &v3.WalletConfigTreeAnyAddressSubdigestLeaf{
			Digest: common.BigToHash(big.NewInt(3)),
		}
		_, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
			sequence.WithGate(subdigestGate))
		require.ErrorContains(t, err, "weight is too large")
	})

	t.Run("gate weight above 1 is allowed", func(t *testing.T) {
		heavyGate := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     2,
			Address:    gateSigner,
			ImageHash_: gateImageHash,
		}
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
			sequence.WithGate(heavyGate))
		require.NoError(t, err)
		require.NotNil(t, config)
	})

	t.Run("same address with different image hash is allowed", func(t *testing.T) {
		otherImageHashLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     1,
			Address:    gateSigner,
			ImageHash_: core.ImageHash{Hash: common.BigToHash(big.NewInt(2))},
		}
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0,
			sequence.WithGate(gateLeaf), sequence.WithSapientSigner(otherImageHashLeaf))
		require.NoError(t, err)
		require.NotNil(t, config)
	})
}

// A sapient-only config (calls is empty) must not build a broken calls group: with no
// subdigest leaves, only the sapient leaf is gated behind gateLeaf. ImageHash must
// still succeed (a nil inner Tree would panic on traversal).
func TestCreateIntentConfigurationWithGateLeaf_EmptyCallsOmitsCallsGate(t *testing.T) {
	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	peerSigner := common.HexToAddress("0x72030E1dbf0a847196ae62EA3ee84BD7ce99D6c1")
	peerSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    peerSigner,
		ImageHash_: core.ImageHash{Hash: common.BigToHash(big.NewInt(1))},
	}
	timedRefundSigner := common.HexToAddress("0x4444444444444444444444444444444444444444")
	destination := common.HexToAddress("0x5555555555555555555555555555555555555555")
	timedRefundImageHash, err := sequence.TimedRefundSapientImageHash(destination, 1_750_000_000)
	require.NoError(t, err)
	timedRefundLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    timedRefundSigner,
		ImageHash_: timedRefundImageHash,
	}

	config, err := sequence.CreateIntentConfiguration(mainSigner, nil, 0, sequence.WithGate(peerSignerLeaf), sequence.WithSapientSigner(timedRefundLeaf))
	require.NoError(t, err)
	require.NotNil(t, config)

	// Must not panic computing the image hash — this is exactly what a nil inner Tree breaks.
	require.NotEqual(t, common.Hash{}, config.ImageHash().Hash)

	require.NotNil(t, findSapientSignerLeaf(config.Tree, timedRefundSigner))
	require.NotNil(t, findSapientSignerLeaf(config.Tree, peerSigner))
}

// A sapient signer leaf (e.g. a timed-refund or gasless-deposit leaf) passed as
// sapientSignerLeafNode shares gateLeaf's gate with the calls leaves: either group
// alone, plus gateLeaf's co-signature, is sufficient. mainSignerLeaf is the only leaf
// never gated, so the owner can always act (e.g. recover funds) regardless of the gate's
// paused state.
func TestCreateIntentConfigurationWithGateLeaf_SapientLeafGated(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, testChain.ChainID(), []v3.Call{
		{
			To:              common.HexToAddress("0x1111111111111111111111111111111111111111"),
			Value:           nil,
			Data:            []byte{0x12, 0x34},
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	mainSigner := common.HexToAddress("0x2222222222222222222222222222222222222222")
	peerSigner := common.HexToAddress("0x72030E1dbf0a847196ae62EA3ee84BD7ce99D6c1")
	peerSignerLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    peerSigner,
		ImageHash_: core.ImageHash{Hash: common.BigToHash(big.NewInt(1))},
	}
	timedRefundSigner := common.HexToAddress("0x4444444444444444444444444444444444444444")
	destination := common.HexToAddress("0x5555555555555555555555555555555555555555")

	timedRefundImageHash, err := sequence.TimedRefundSapientImageHash(destination, 1_750_000_000)
	require.NoError(t, err)
	timedRefundLeaf := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    timedRefundSigner,
		ImageHash_: timedRefundImageHash,
	}

	config, err := sequence.CreateIntentConfiguration(
		mainSigner,
		[]*v3.CallsPayload{&payload},
		0,
		sequence.WithGate(peerSignerLeaf),
		sequence.WithSapientSigner(timedRefundLeaf),
	)
	require.NoError(t, err)

	// mainSignerLeaf stays a plain sibling, never wrapped by any gate.
	top, ok := config.Tree.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	ownerLeaf, ok := top.Left.(*v3.WalletConfigTreeAddressLeaf)
	require.True(t, ok)
	require.Equal(t, mainSigner, ownerLeaf.Address)

	// The calls leaves and timedRefundLeaf share a single merged gate
	// (threshold = gateWeight+1), not two separate gates.
	mergedGate, ok := top.Right.(*v3.WalletConfigTreeNestedLeaf)
	require.True(t, ok, "calls and timedRefundLeaf must share one gate")
	require.Equal(t, uint16(peerSignerLeaf.Weight)+1, mergedGate.Threshold, "outer gate is gateWeight+1")

	require.NotNil(t, findSapientSignerLeaf(config.Tree, peerSigner))
	require.NotNil(t, findSapientSignerLeaf(config.Tree, timedRefundSigner))

	// Absent any signatures, neither gate has anything to recover: weight must be 0. Real
	// providers are only invoked to check an actually-embedded sapient signature, so this
	// stays a pure offline check.
	signatureNoSigs, err := sequence.BuildIntentConfigurationSignature(config, nil)
	require.NoError(t, err)
	decodedNoSigs, err := v3.Core.DecodeSignature(signatureNoSigs)
	require.NoError(t, err)
	_, weightNoSigs, err := decodedNoSigs.Recover(context.Background(), payload, nil)
	require.NoError(t, err)
	require.Truef(t, weightNoSigs.Cmp(big.NewInt(int64(config.Threshold()))) < 0,
		"recovered weight %v must not meet threshold %v with no signatures at all", weightNoSigs, config.Threshold())

	// Each leaf's signature must change the encoding on its own, proving it's wired into
	// its gated position.
	peerSignature := &core.SignerSignature{
		Signer:    core.SapientSigner(peerSigner, peerSignerLeaf.ImageHash_.Hash),
		Signature: []byte{},
		Type:      core.SignerSignatureTypeSapientCompact,
	}
	signatureWithGate, err := sequence.BuildIntentConfigurationSignature(config, []*core.SignerSignature{peerSignature})
	require.NoError(t, err)
	require.NotEqual(t, signatureNoSigs, signatureWithGate)

	timedRefundSignature := &core.SignerSignature{
		Signer:    core.SapientSigner(timedRefundSigner, timedRefundLeaf.ImageHash_.Hash),
		Signature: []byte{},
		Type:      core.SignerSignatureTypeSapientCompact,
	}
	signatureWithTimedRefund, err := sequence.BuildIntentConfigurationSignature(config, []*core.SignerSignature{timedRefundSignature})
	require.NoError(t, err)
	require.NotEqual(t, signatureNoSigs, signatureWithTimedRefund)

	// Supplying both signatures must embed both, deterministically. The gate leaf alone
	// meets the config threshold via the calls gate's payload-independent subdigest leaf,
	// so an early-cancelling builder could drop the other signature depending on goroutine
	// scheduling; every build must include both signatures and produce identical bytes.
	both := []*core.SignerSignature{peerSignature, timedRefundSignature}
	signatureCombined, err := sequence.BuildIntentConfigurationSignature(config, both)
	require.NoError(t, err)
	require.NotEqual(t, signatureNoSigs, signatureCombined)
	require.NotEqual(t, signatureWithGate, signatureCombined, "combined signature must also embed the timed refund signature")
	require.NotEqual(t, signatureWithTimedRefund, signatureCombined, "combined signature must also embed the gate signature")
	for range 50 {
		rebuilt, err := sequence.BuildIntentConfigurationSignature(config, both)
		require.NoError(t, err)
		require.Equal(t, signatureCombined, rebuilt, "combined signature must be deterministic")
	}
}

func TestTimedRefundSapientImageHash(t *testing.T) {
	destination := common.HexToAddress("0x4444444444444444444444444444444444444444")

	imageHash, err := sequence.TimedRefundSapientImageHash(destination, 1_750_000_000)
	require.NoError(t, err)
	require.Equal(t, common.HexToHash("0x577e11f2280512fff4541fc08cc7eb98357bdcff482db5634db7327e3c97ba58"), imageHash.Hash)

	preimage, ok := imageHash.Preimage.(*sequence.TimedRefundSapientImageHashPreimage)
	require.True(t, ok)
	require.Equal(t, destination, preimage.Destination)
	require.Equal(t, uint64(1_750_000_000), preimage.UnlockTimestamp)
	require.Equal(t, imageHash.Hash, preimage.ImageHash().Hash)
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
			Value:           big.NewInt(0),
			Data:            calldata,
			GasLimit:        big.NewInt(0),
			DelegateCall:    false,
			OnlyFallback:    false,
			BehaviorOnError: v3.BehaviorOnErrorRevert,
		},
	}, big.NewInt(0), big.NewInt(0))

	sapientContract := testChain.UniDeploy(t, "MOCK_SAPIENT", 0)
	sapientSignerAddress := sapientContract.Address
	sapientImageHash := common.HexToHash("0x1234567890ABCDEF01234567890ABCDEF01234567890ABCDEF1234567890ABCD")
	sapientSignerLeafNode := &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    sapientSignerAddress,
		ImageHash_: core.ImageHash{Hash: sapientImageHash},
	}

	sapientSignerSignature := sapientImageHash.Hex()
	signerSignature := &core.SignerSignature{
		Signer: core.Signer{
			Address:   sapientSignerAddress,
			IsSapient: true,
			ImageHash: sapientImageHash,
		},
		Signature: common.HexToHash(sapientSignerSignature).Bytes(),
		Type:      core.SignerSignatureTypeSapient,
	}

	t.Run("signature matches subdigest", func(t *testing.T) {
		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), []*v3.CallsPayload{&payload}, 0)
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, nil)
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
		recoveredConfig, _, err := sig.Recover(context.Background(), payload, nil)
		// spew.Dump(recoveredConfig)

		require.NoError(t, err)
		require.NotNil(t, recoveredConfig, "recovered config should not be nil")

		// Get the full signature in string
		sigDataStr, err := sig.Data()
		require.NoError(t, err)

		anyAddressSubdigestStr := anyAddressSubdigestLeaf.Digest.Hex()

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
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload1}, 0, nil)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload2}, 0, nil)
		require.NoError(t, err)

		// Verify signatures are different
		require.NotEqual(t, sig1, sig2, "different transactions should produce different signatures")
	})

	t.Run("same transactions produce same signatures", func(t *testing.T) {
		// Use the payload directly
		sig1, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, nil)
		require.NoError(t, err)

		sig2, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, nil)
		require.NoError(t, err)

		// Verify signatures are the same
		require.Equal(t, sig1, sig2, "same transactions should produce same signatures")
	})

	t.Run("signer signature included in the signature tree", func(t *testing.T) {
		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, []*core.SignerSignature{signerSignature}, sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)

		sapientLeaf := findSapientSignerLeaf(config.Tree, sapientSignerAddress)
		require.NotNil(t, sapientLeaf)
		require.Equal(t, sapientImageHash, sapientLeaf.ImageHash_.Hash)

		// Verify the signature can be decoded
		sig, err := v3.Core.DecodeSignature(signature)
		require.NoError(t, err, "signature should be decodable")

		// Get the config from the signature
		recoveredSignerSignatures := map[core.Signer]core.SignerSignature{}
		recoveredConfig, _, err := sig.Recover(context.Background(), payload, testChain.Provider, recoveredSignerSignatures)
		require.NoError(t, err)
		require.NotNil(t, recoveredConfig, "recovered config should not be nil")
		require.Len(t, recoveredSignerSignatures, 1, "expected exactly one recovered sapient signer signature")
		var recoveredSapientSig core.SignerSignature
		for signer, sig := range recoveredSignerSignatures {
			if signer.Address == sapientSignerAddress {
				recoveredSapientSig = sig
				break
			}
		}
		require.NotNil(t, recoveredSapientSig.Signature, "sapient signer signature should be recovered")
		require.Equal(t, signerSignature.Signature, recoveredSapientSig.Signature, "recovered sapient signature should match")

		// Get the full signature in string
		sigDataStr, err := sig.Data()
		require.NoError(t, err)

		sapientImageHash := sapientLeaf.ImageHash().Hash.Hex()
		require.Equal(t, sapientImageHash, sapientSignerLeafNode.ImageHash().Hash.Hex())

		// Verify the signature contains the sapient signature
		require.Contains(t, common.Bytes2Hex(sigDataStr), sapientSignerSignature[2:], "signature should contain the sapient signer signature")
	})

	t.Run("gate signature included in the signature tree", func(t *testing.T) {
		gateContract := testChain.UniDeploy(t, "MOCK_SAPIENT", 1)
		gateSignerAddress := gateContract.Address
		gateImageHash := common.HexToHash("0xABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF123456789A")
		gateLeafNode := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     1,
			Address:    gateSignerAddress,
			ImageHash_: core.ImageHash{Hash: gateImageHash},
		}
		gateSignature := &core.SignerSignature{
			Signer: core.Signer{
				Address:   gateSignerAddress,
				IsSapient: true,
				ImageHash: gateImageHash,
			},
			Signature: gateImageHash.Bytes(),
			Type:      core.SignerSignatureTypeSapient,
		}

		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, sequence.WithGate(gateLeafNode))
		require.NoError(t, err)

		// Create the signature
		signature, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, []*core.SignerSignature{gateSignature}, sequence.WithGate(gateLeafNode))
		require.NoError(t, err)

		gateLeaf := findSapientSignerLeaf(config.Tree, gateSignerAddress)
		require.NotNil(t, gateLeaf)
		require.Equal(t, gateImageHash, gateLeaf.ImageHash_.Hash)

		// Verify the signature can be decoded
		sig, err := v3.Core.DecodeSignature(signature)
		require.NoError(t, err, "signature should be decodable")

		// Get the config from the signature
		recoveredSignerSignatures := map[core.Signer]core.SignerSignature{}
		recoveredConfig, _, err := sig.Recover(context.Background(), payload, testChain.Provider, recoveredSignerSignatures)
		require.NoError(t, err)
		require.NotNil(t, recoveredConfig, "recovered config should not be nil")
		require.Len(t, recoveredSignerSignatures, 1, "expected exactly one recovered gate signer signature")
		var recoveredGateSig core.SignerSignature
		for signer, sig := range recoveredSignerSignatures {
			if signer.Address == gateSignerAddress {
				recoveredGateSig = sig
				break
			}
		}
		require.NotNil(t, recoveredGateSig.Signature, "gate signer signature should be recovered")
		require.Equal(t, gateSignature.Signature, recoveredGateSig.Signature, "recovered gate signature should match")

		// Get the full signature in string
		sigDataStr, err := sig.Data()
		require.NoError(t, err)

		// Verify the signature contains the gate signature
		require.Contains(t, common.Bytes2Hex(sigDataStr), gateImageHash.Hex()[2:], "signature should contain the gate signer signature")
	})

	t.Run("gate and sapient signer signatures included in the signature tree", func(t *testing.T) {
		gateContract := testChain.UniDeploy(t, "MOCK_SAPIENT", 2)
		gateSignerAddress := gateContract.Address
		gateImageHash := common.HexToHash("0xFEDCBA0987654321FEDCBA0987654321FEDCBA0987654321FEDCBA098765432")
		gateLeafNode := &v3.WalletConfigTreeSapientSignerLeaf{
			Weight:     1,
			Address:    gateSignerAddress,
			ImageHash_: core.ImageHash{Hash: gateImageHash},
		}
		gateSignature := &core.SignerSignature{
			Signer: core.Signer{
				Address:   gateSignerAddress,
				IsSapient: true,
				ImageHash: gateImageHash,
			},
			Signature: gateImageHash.Bytes(),
			Type:      core.SignerSignatureTypeSapient,
		}

		// Create the intent configuration
		config, err := sequence.CreateIntentConfiguration(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, sequence.WithGate(gateLeafNode), sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)

		sapientLeaf := findSapientSignerLeaf(config.Tree, sapientSignerAddress)
		require.NotNil(t, sapientLeaf)
		require.Equal(t, sapientImageHash, sapientLeaf.ImageHash_.Hash)

		gateLeaf := findSapientSignerLeaf(config.Tree, gateSignerAddress)
		require.NotNil(t, gateLeaf)
		require.Equal(t, gateImageHash, gateLeaf.ImageHash_.Hash)

		// Each signature is checked independently first to prove its leaf's wiring, then
		// combined to prove a single build embeds both.
		signatureNoSigs, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, nil, sequence.WithGate(gateLeafNode), sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)

		signatureWithGateSig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, []*core.SignerSignature{gateSignature}, sequence.WithGate(gateLeafNode), sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)
		require.NotEqual(t, signatureNoSigs, signatureWithGateSig, "including the gate signature must change the encoding")

		sigWithGateSig, err := v3.Core.DecodeSignature(signatureWithGateSig)
		require.NoError(t, err, "signature should be decodable")
		gateRecoveredSignatures := map[core.Signer]core.SignerSignature{}
		_, _, err = sigWithGateSig.Recover(context.Background(), payload, testChain.Provider, gateRecoveredSignatures)
		require.NoError(t, err)
		require.Len(t, gateRecoveredSignatures, 1, "expected exactly one recovered gate signer signature")
		for signer, sig := range gateRecoveredSignatures {
			require.Equal(t, gateSignerAddress, signer.Address)
			require.Equal(t, gateSignature.Signature, sig.Signature, "recovered gate signature should match")
		}

		signatureWithSapientSig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, []*core.SignerSignature{signerSignature}, sequence.WithGate(gateLeafNode), sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)
		require.NotEqual(t, signatureNoSigs, signatureWithSapientSig, "including the sapient signature must change the encoding")

		sigWithSapientSig, err := v3.Core.DecodeSignature(signatureWithSapientSig)
		require.NoError(t, err, "signature should be decodable")
		sapientRecoveredSignatures := map[core.Signer]core.SignerSignature{}
		_, _, err = sigWithSapientSig.Recover(context.Background(), payload, testChain.Provider, sapientRecoveredSignatures)
		require.NoError(t, err)
		require.Len(t, sapientRecoveredSignatures, 1, "expected exactly one recovered sapient signer signature")
		for signer, sig := range sapientRecoveredSignatures {
			require.Equal(t, sapientSignerAddress, signer.Address)
			require.Equal(t, signerSignature.Signature, sig.Signature, "recovered sapient signature should match")
		}

		// Supplying both signatures in one call must embed both: the gate leaf alone meets
		// the config threshold via the calls gate's payload-independent subdigest leaf, so
		// an early-cancelling builder could nondeterministically drop the sapient signature.
		signatureCombined, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload}, 0, []*core.SignerSignature{gateSignature, signerSignature}, sequence.WithGate(gateLeafNode), sequence.WithSapientSigner(sapientSignerLeafNode))
		require.NoError(t, err)
		require.NotEqual(t, signatureWithGateSig, signatureCombined, "combined signature must also embed the sapient signature")
		require.NotEqual(t, signatureWithSapientSig, signatureCombined, "combined signature must also embed the gate signature")

		sigCombined, err := v3.Core.DecodeSignature(signatureCombined)
		require.NoError(t, err, "signature should be decodable")
		combinedRecoveredSignatures := map[core.Signer]core.SignerSignature{}
		_, _, err = sigCombined.Recover(context.Background(), payload, testChain.Provider, combinedRecoveredSignatures)
		require.NoError(t, err)
		require.Len(t, combinedRecoveredSignatures, 2, "expected both the gate and sapient signer signatures recovered")
		for signer, sig := range combinedRecoveredSignatures {
			switch signer.Address {
			case gateSignerAddress:
				require.Equal(t, gateSignature.Signature, sig.Signature, "recovered gate signature should match")
			case sapientSignerAddress:
				require.Equal(t, signerSignature.Signature, sig.Signature, "recovered sapient signature should match")
			default:
				require.Failf(t, "unexpected recovered signer", "address %s", signer.Address)
			}
		}
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
	sig, err := sequence.GetIntentConfigurationSignature(eoa1.Address(), []*v3.CallsPayload{&payload1}, 0, nil)
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
	for signer := range signers {
		mainSigner = signer.Address
		break
	}
	require.NotZero(t, mainSigner)

	// Generate a configuration signature for the batch.
	intentConfigSig, err := sequence.GetIntentConfigurationSignature(mainSigner, []*v3.CallsPayload{&payload}, 0, nil)
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
	result, _, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, sequence.MetaTxnID(opPayload.Digest().Hash.Hex()[2:]))
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
	for signer := range signers {
		mainSigner = signer.Address
		break
	}
	require.NotZero(t, mainSigner)

	// Generate a configuration signature for both batches
	intentConfigSig, err := sequence.GetIntentConfigurationSignature(mainSigner, payloads, 0, nil)
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

		result, _, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, expectedMetaTxnID)
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
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, 0)
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

		assert.Equal(t, common.HexToAddress("0x3857F10693Aa54b39B640292CD5723ec97BEd285"), address)
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
		config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload1, &payload2}, 0)
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

		assert.Equal(t, common.HexToAddress("0x5784cF2b86eE8C930ee26279e7666241aC7e78B7"), address)
	})
}

func TestIntentConfigurationAddress_WithCheckpoint(t *testing.T) {
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
	mainSigner := common.HexToAddress("0x1111111111111111111111111111111111111111")
	context := sequence.V3SequenceContext()

	checkpoint1 := uint64(1)
	checkpoint2 := uint64(2)

	// Create intent configuration
	config1, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, checkpoint1)
	require.NoError(t, err)
	config2, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload}, checkpoint2)
	require.NoError(t, err)

	// Checkpoints should be set correctly
	assert.Equal(t, checkpoint1, config1.Checkpoint())
	assert.Equal(t, checkpoint2, config2.Checkpoint())

	// Configurations should resolve to different image hashes and addresses
	imageHash1 := config1.ImageHash()
	imageHash2 := config2.ImageHash()
	assert.NotEqual(t, imageHash1.Hash, imageHash2.Hash)

	// Calculate counterfactual address
	address1, err := sequence.AddressFromImageHash(imageHash1, context)
	require.NoError(t, err)
	address2, err := sequence.AddressFromImageHash(imageHash2, context)
	require.NoError(t, err)
	assert.NotEqual(t, address1, address2)
}

func TestIntentConfigurationAddress_RealWorldExample(t *testing.T) {
	// Main signer matching TypeScript test
	mainSigner := common.HexToAddress("0x8456195dd0793c621c7f9245edF0fEf85b1B879C")

	// Context for the counterfactual address
	context := sequence.V3SequenceContext()

	// Create first operation on Arbitrum (chainId: 42161)
	payload1 := v3.NewCallsPayload(common.Address{}, big.NewInt(42161), []v3.Call{
		{
			To:              common.HexToAddress("0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae"),
			Value:           big.NewInt(16618237),
			Data:            common.FromHex("0xa6010a660000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000005801784da62343d604885e1181a759647447f13330cc2b8c925cda864b1ac1ce8fc00000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d737461727461746556324275730000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000086c6966692d61706900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000111111125421ca6dc452d289314280a0f8842a65000000000000000000000000111111125421ca6dc452d289314280a0f8842a650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e583100000000000000000000000000000000000000000000000000000f42af44fef500000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000020807ed2379000000000000000000000000de9e4fe32b049f821c7f3e9802381aa470ffca73000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000af88d065e77c8cc2239327c5edb3a432268e5831000000000000000000000000de9e4fe32b049f821c7f3e9802381aa470ffca730000000000000000000000001231deb6f5749ef6ce6943a275a1d3e7486f4eae00000000000000000000000000000000000000000000000000000f42af44fef500000000000000000000000000000000000000000000000000000000000075320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000ba00000000000000000000000000000000000000000000000000000000009c4160de632c3a214d5f14c1d8ddf0b92f8bcd188fee4500242668dfaa000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007532000000000000000000000000111111125421ca6dc452d289314280a0f8842a650000000000002a94d114000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000077368def90800000000000000000000000000000000000000000000000000000000000000000000000000000000000000008456195dd0793c621c7f9245edf0fef85b1b879c00000000000000000000000000000000000000000000000000000000000075e80000000000000000000000008456195dd0793c621c7f9245edf0fef85b1b879c0000000000000000000000000000000000000000000000000000000000007532000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000"),
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
	config, err := sequence.CreateIntentConfiguration(mainSigner, []*v3.CallsPayload{&payload1, &payload2}, 0)
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
	assert.Equal(t, common.HexToAddress("0x2c32Ffc5FbD28F34aDBaEEa7b493fbae1e44F7EC"), address)
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

func ecrecoverForTest(hash, sig []byte) (common.Address, error) {
	sigCopy := make([]byte, len(sig))
	copy(sigCopy, sig)
	if len(sigCopy) == 65 && sigCopy[64] >= 27 {
		sigCopy[64] -= 27
	}

	pubKeyBytes, err := crypto.Ecrecover(hash, sigCopy)
	if err != nil {
		return common.Address{}, err
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*pubKey), nil
}
