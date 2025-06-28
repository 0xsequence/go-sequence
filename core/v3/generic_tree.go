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
