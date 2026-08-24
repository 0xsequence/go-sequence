package v3

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// GraftConfigTreeNode returns tree with every anonymous node-hash leaf whose
// image hash equals replacement.ImageHash() swapped for replacement, leaving
// the tree's image hash unchanged. ok reports whether the hash was found,
// either as a node leaf (replaced) or already as a full subtree (returned
// unchanged).
func GraftConfigTreeNode(tree WalletConfigTree, replacement WalletConfigTree) (grafted WalletConfigTree, ok bool) {
	if tree == nil || replacement == nil {
		return tree, false
	}
	return graftConfigTreeNode(tree, replacement, replacement.ImageHash().Hash)
}

func graftConfigTreeNode(tree WalletConfigTree, replacement WalletConfigTree, hash common.Hash) (WalletConfigTree, bool) {
	if tree == nil {
		return tree, false
	}

	if tree.ImageHash().Hash == hash {
		switch tree.(type) {
		case WalletConfigTreeNodeLeaf, *WalletConfigTreeNodeLeaf:
			return replacement, true
		}
		return tree, true
	}

	switch n := tree.(type) {
	case *WalletConfigTreeNode:
		left, leftOk := graftConfigTreeNode(n.Left, replacement, hash)
		right, rightOk := graftConfigTreeNode(n.Right, replacement, hash)
		if leftOk || rightOk {
			return &WalletConfigTreeNode{Left: left, Right: right}, true
		}
	case *WalletConfigTreeNestedLeaf:
		if inner, ok := graftConfigTreeNode(n.Tree, replacement, hash); ok {
			return &WalletConfigTreeNestedLeaf{Weight: n.Weight, Threshold: n.Threshold, Tree: inner}, true
		}
	}

	return tree, false
}
