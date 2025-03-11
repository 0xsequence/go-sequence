package v3

import (
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

type Leaf struct {
	Type  string `json:"type"`
	Value []byte `json:"value"`
}

type Node []byte

type Branch []Tree

type Tree interface{}

func IsBranch(tree Tree) bool {
	_, ok := tree.(Branch)
	return ok
}

func IsLeaf(tree Tree) bool {
	if leaf, ok := tree.(Leaf); ok {
		return leaf.Type == "leaf" && len(leaf.Value) > 0
	}
	return false
}

func IsNode(tree Tree) bool {
	node, ok := tree.(Node)
	return ok && len(node) == 32
}

func IsTree(tree interface{}) bool {
	return IsBranch(tree) || IsLeaf(tree) || IsNode(tree)
}

// Hash computes the Keccak-256 hash of a tree
func Hash(tree Tree) []byte {
	if IsBranch(tree) {
		branch := tree.(Branch)
		if len(branch) == 0 {
			return []byte{}
		}

		hashedChildren := make([][]byte, len(branch))
		for i, child := range branch {
			hashedChildren[i] = Hash(child)
		}

		chash := hashedChildren[0]
		for i := 1; i < len(hashedChildren); i++ {
			combined := append(chash, hashedChildren[i]...)
			chash = crypto.Keccak256(combined)
		}
		return chash
	}

	if IsNode(tree) {
		return tree.(Node)
	}

	if IsLeaf(tree) {
		leaf := tree.(Leaf)
		return crypto.Keccak256(leaf.Value)
	}

	return []byte{}
}

// DeepCopy creates a deep copy of a tree
func DeepCopy(tree Tree) Tree {
	if IsBranch(tree) {
		branch := tree.(Branch)
		newBranch := make(Branch, len(branch))
		for i, child := range branch {
			newBranch[i] = DeepCopy(child)
		}
		return newBranch
	}

	if IsLeaf(tree) {
		leaf := tree.(Leaf)
		newLeaf := Leaf{
			Type:  leaf.Type,
			Value: make([]byte, len(leaf.Value)),
		}
		copy(newLeaf.Value, leaf.Value)
		return newLeaf
	}

	if IsNode(tree) {
		node := tree.(Node)
		newNode := make(Node, len(node))
		copy(newNode, node)
		return newNode
	}

	return nil
}

// ToConfigurationTree converts a tree to a ConfigurationTree
func ToConfigurationTree(tree Tree) *ConfigurationTree {
	nodeHash := Hash(tree)
	var imageHash [32]byte
	copy(imageHash[:], nodeHash)

	nodes := extractNodes(tree)

	return &ConfigurationTree{
		ImageHash: imageHash,
		Nodes:     nodes,
	}
}

// Helper function to extract SessionNodes from a tree
func extractNodes(tree Tree) []SessionNode {
	if IsBranch(tree) {
		branch := tree.(Branch)
		var nodes []SessionNode

		for _, child := range branch {
			childNodes := extractNodes(child)
			nodes = append(nodes, childNodes...)
		}

		return nodes
	}

	return []SessionNode{}
}
