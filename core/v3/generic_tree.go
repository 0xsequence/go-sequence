package v3

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

// Leaf represents a leaf node.
// Its Value field contains the raw data which will be hashed when needed.
type Leaf struct {
	Value []byte `json:"value"`
}

// Node represents a pre-hashed node.
// Its byte slice MUST be exactly 32 bytes.
type Node []byte

// Branch represents a branch in the tree.
// It is a slice of Tree values (with at least 2 elements).
type Branch []Tree

// Tree is a union type that may be a Branch, a Leaf, or a Node.
type Tree interface{}

// IsBranch returns true if the given tree is a Branch (an array with at least two elements)
// and every child is a valid Tree.
func IsBranch(tree Tree) bool {
	b, ok := tree.(Branch)
	if !ok {
		return false
	}
	if len(b) < 2 {
		return false
	}
	for _, child := range b {
		if !IsTree(child) {
			return false
		}
	}
	return true
}

// IsLeaf returns true if the given tree is a Leaf.
func IsLeaf(tree Tree) bool {
	_, ok := tree.(Leaf)
	return ok
}

// IsNode returns true if the given tree is a Node (i.e. a byte slice of length 32).
func IsNode(tree Tree) bool {
	n, ok := tree.(Node)
	return ok && len(n) == 32
}

// IsTree returns true if the provided value is either a Branch, a Leaf, or a Node.
func IsTree(tree Tree) bool {
	return IsBranch(tree) || IsLeaf(tree) || IsNode(tree)
}

// HashTree computes the hash of a Tree as follows:
//   - If the tree is a Branch, it recursively hashes each child and then sequentially
//     combines them by concatenation and keccak256 hashing.
//   - If the tree is a Node, it is already hashed, so it is returned as is.
//   - If the tree is a Leaf, its Value is hashed using keccak256.
func HashTree(tree Tree) ([]byte, error) {
	if IsBranch(tree) {
		b := tree.(Branch)
		if len(b) == 0 {
			return nil, fmt.Errorf("empty branch")
		}
		// Start with the hash of the first child.
		h, err := HashTree(b[0])
		if err != nil {
			return nil, err
		}
		// Sequentially hash with each subsequent child's hash.
		for i := 1; i < len(b); i++ {
			childHash, err := HashTree(b[i])
			if err != nil {
				return nil, err
			}
			// Concatenate the current hash and the child hash, then compute keccak256.
			combined := append(h, childHash...)
			h = crypto.Keccak256(combined)
		}
		return PadLeft(h, 32), nil
	}

	if IsNode(tree) {
		return PadLeft(tree.(Node), 32), nil
	}

	if IsLeaf(tree) {
		leaf := tree.(Leaf)
		h := crypto.Keccak256(leaf.Value)
		return PadLeft(h, 32), nil
	}

	return nil, fmt.Errorf("invalid tree structure")
}
