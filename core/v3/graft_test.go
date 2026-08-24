package v3_test

import (
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

func graftTestSapientLeaf() *v3.WalletConfigTreeSapientSignerLeaf {
	return &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     1,
		Address:    common.HexToAddress("0x9999999999999999999999999999999999999999"),
		ImageHash_: core.ImageHash{Hash: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000042")},
	}
}

// anonymousNodeLeaf mimics a decoded node-hash leaf: hash only, no preimage.
func anonymousNodeLeaf(subtree v3.WalletConfigTree) core.ImageHash {
	return core.ImageHash{Hash: subtree.ImageHash().Hash}
}

func TestGraftConfigTreeNodeValueLeaf(t *testing.T) {
	replacement := graftTestSapientLeaf()

	tree := &v3.WalletConfigTreeNode{
		Left: &v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("0x1111111111111111111111111111111111111111")},
		Right: &v3.WalletConfigTreeNode{
			Left:  v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
			Right: &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x2222222222222222222222222222222222222222")},
		},
	}

	grafted, ok := v3.GraftConfigTreeNode(tree, replacement)
	require.True(t, ok)
	require.Equal(t, tree.ImageHash().Hash, grafted.ImageHash().Hash)

	node, ok := grafted.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	inner, ok := node.Right.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	require.Same(t, replacement, inner.Left)

	// The input tree is not mutated.
	require.IsType(t, v3.WalletConfigTreeNodeLeaf{}, tree.Right.(*v3.WalletConfigTreeNode).Left)
}

func TestGraftConfigTreeNodePointerLeaf(t *testing.T) {
	replacement := graftTestSapientLeaf()

	tree := &v3.WalletConfigTreeNode{
		Left:  &v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
		Right: &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x2222222222222222222222222222222222222222")},
	}

	grafted, ok := v3.GraftConfigTreeNode(tree, replacement)
	require.True(t, ok)
	require.Equal(t, tree.ImageHash().Hash, grafted.ImageHash().Hash)

	node, ok := grafted.(*v3.WalletConfigTreeNode)
	require.True(t, ok)
	require.Same(t, replacement, node.Left)
}

func TestGraftConfigTreeNodeNestedLeaf(t *testing.T) {
	// Non-sapient replacement: a plain subtree of address leaves.
	replacement := &v3.WalletConfigTreeNode{
		Left:  &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x3333333333333333333333333333333333333333")},
		Right: &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x4444444444444444444444444444444444444444")},
	}

	tree := &v3.WalletConfigTreeNestedLeaf{
		Weight:    3,
		Threshold: 2,
		Tree:      v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
	}

	grafted, ok := v3.GraftConfigTreeNode(tree, replacement)
	require.True(t, ok)
	require.Equal(t, tree.ImageHash().Hash, grafted.ImageHash().Hash)

	nested, ok := grafted.(*v3.WalletConfigTreeNestedLeaf)
	require.True(t, ok)
	require.Equal(t, uint8(3), nested.Weight)
	require.Equal(t, uint16(2), nested.Threshold)
	require.Same(t, replacement, nested.Tree)
}

func TestGraftConfigTreeNodeAlreadyPresent(t *testing.T) {
	present := graftTestSapientLeaf()

	tree := &v3.WalletConfigTreeNode{
		Left:  &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x1111111111111111111111111111111111111111")},
		Right: present,
	}

	grafted, ok := v3.GraftConfigTreeNode(tree, graftTestSapientLeaf())
	require.True(t, ok)
	require.Equal(t, tree, grafted)
	require.Same(t, present, grafted.(*v3.WalletConfigTreeNode).Right)

	// The whole tree matching the replacement is also left unchanged.
	grafted, ok = v3.GraftConfigTreeNode(present, graftTestSapientLeaf())
	require.True(t, ok)
	require.Same(t, present, grafted)
}

// TestGraftConfigTreeNodeAllOccurrences covers the duplicate-hash case: an
// already-materialized occurrence must not stop the traversal from grafting a
// later anonymous occurrence, and multiple anonymous occurrences are all
// replaced.
func TestGraftConfigTreeNodeAllOccurrences(t *testing.T) {
	replacement := graftTestSapientLeaf()

	// Materialized on the left, anonymous on the right.
	tree := &v3.WalletConfigTreeNode{
		Left:  graftTestSapientLeaf(),
		Right: v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
	}
	grafted, ok := v3.GraftConfigTreeNode(tree, replacement)
	require.True(t, ok)
	require.Equal(t, tree.ImageHash().Hash, grafted.ImageHash().Hash)
	node, isNode := grafted.(*v3.WalletConfigTreeNode)
	require.True(t, isNode)
	require.Same(t, replacement, node.Right)

	// Two anonymous occurrences: both replaced.
	tree = &v3.WalletConfigTreeNode{
		Left:  &v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
		Right: v3.WalletConfigTreeNodeLeaf{Node: anonymousNodeLeaf(replacement)},
	}
	grafted, ok = v3.GraftConfigTreeNode(tree, replacement)
	require.True(t, ok)
	require.Equal(t, tree.ImageHash().Hash, grafted.ImageHash().Hash)
	node, isNode = grafted.(*v3.WalletConfigTreeNode)
	require.True(t, isNode)
	require.Same(t, replacement, node.Left)
	require.Same(t, replacement, node.Right)
}

func TestGraftConfigTreeNodeNotFound(t *testing.T) {
	tree := &v3.WalletConfigTreeNode{
		Left:  &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x1111111111111111111111111111111111111111")},
		Right: v3.WalletConfigTreeNodeLeaf{Node: core.ImageHash{Hash: common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000ff")}},
	}

	grafted, ok := v3.GraftConfigTreeNode(tree, graftTestSapientLeaf())
	require.False(t, ok)
	require.Same(t, tree, grafted)
}

func TestGraftConfigTreeNodeNil(t *testing.T) {
	tree := &v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0x1111111111111111111111111111111111111111")}

	grafted, ok := v3.GraftConfigTreeNode(nil, graftTestSapientLeaf())
	require.False(t, ok)
	require.Nil(t, grafted)

	grafted, ok = v3.GraftConfigTreeNode(tree, nil)
	require.False(t, ok)
	require.Same(t, tree, grafted)
}
