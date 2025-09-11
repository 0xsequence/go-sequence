package v3

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
)

type Tree = core.ImageHashable

type TreeNode []Tree

var _ Tree = TreeNode{}

func (n TreeNode) ImageHash() core.ImageHash {
	if len(n) == 0 {
		return core.ImageHash{Hash: common.Hash{}, Preimage: n}
	}

	imageHash := n[0].ImageHash().Hash
	for _, tree := range n[1:] {
		imageHash = crypto.Keccak256Hash(imageHash.Bytes(), tree.ImageHash().Bytes())
	}

	return core.ImageHash{Hash: imageHash, Preimage: n}
}

type TreeLeaf []byte

var _ Tree = TreeLeaf{}

func (l TreeLeaf) ImageHash() core.ImageHash {
	return core.ImageHash{Hash: crypto.Keccak256Hash(l), Preimage: l}
}

func MergeSubtrees(tree Tree, subtrees map[common.Hash]Tree) (Tree, bool) {
	switch tree := tree.(type) {
	case core.ImageHash:
		subtree, ok := subtrees[tree.Hash]
		if ok {
			return subtree, true
		} else if tree.Preimage != nil && tree.Preimage != tree {
			subtree_, _ := MergeSubtrees(tree.Preimage, subtrees)
			return subtree_, true
		} else {
			return tree, false
		}

	case *core.ImageHash:
		subtree, ok := subtrees[tree.Hash]
		if ok {
			return subtree, true
		} else if tree.Preimage != nil && tree.Preimage != tree {
			subtree_, _ := MergeSubtrees(tree.Preimage, subtrees)
			return subtree_, true
		} else {
			return tree, false
		}

	case TreeNode:
		updated := false

		node := make(TreeNode, 0, len(tree))
		for _, subtree := range tree {
			subtree_, updated_ := MergeSubtrees(subtree, subtrees)
			if updated_ {
				updated = true
			}
			node = append(node, subtree_)
		}

		if updated {
			return node, true
		} else {
			return tree, false
		}

	default:
		return tree, false
	}
}

func Subtrees(tree Tree) map[common.Hash]Tree {
	subtrees := map[common.Hash]Tree{}
	readSubtrees(subtrees, tree)
	return subtrees
}

func readSubtrees(subtrees map[common.Hash]Tree, tree Tree) {
	switch tree := tree.(type) {
	case core.ImageHash:
		if tree.Preimage != nil && tree.Preimage != tree {
			readSubtrees(subtrees, tree.Preimage)
		}

	case *core.ImageHash:
		if tree.Preimage != nil && tree.Preimage != tree {
			readSubtrees(subtrees, tree.Preimage)
		}

	case TreeNode:
		for _, subtree := range tree {
			readSubtrees(subtrees, subtree)
		}

		subtrees[tree.ImageHash().Hash] = tree

	default:
		subtrees[tree.ImageHash().Hash] = tree
	}
}

func IsCompleteTree(tree Tree) bool {
	switch tree := tree.(type) {
	case core.ImageHash, *core.ImageHash:
		return false

	case TreeNode:
		for _, subtree := range tree {
			if !IsCompleteTree(subtree) {
				return false
			}
		}

		return true

	default:
		return true
	}
}
