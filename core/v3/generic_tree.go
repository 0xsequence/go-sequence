package v3

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
)

type Tree = core.ImageHashable

type TreeBranch []Tree

var _ Tree = TreeBranch{}

func (b TreeBranch) ImageHash() core.ImageHash {
	var imageHash common.Hash

	for i, tree := range b {
		if i == 0 {
			imageHash = tree.ImageHash().Hash
		} else {
			imageHash = crypto.Keccak256Hash(imageHash.Bytes(), tree.ImageHash().Bytes())
		}
	}

	return core.ImageHash{Hash: imageHash, Preimage: b}
}

type TreeLeaf []byte

var _ Tree = TreeLeaf{}

func (l TreeLeaf) ImageHash() core.ImageHash {
	return core.ImageHash{Hash: crypto.Keccak256Hash(l), Preimage: l}
}
