package core

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

type PasskeySignature struct {
	requireUserVerification bool
	x, y                    common.Hash
	metadata                common.Hash
}

func DecodePasskeySignature(signature []byte) (PasskeySignature, error) {
}

func (s PasskeySignature) Encode() []byte {
}

func (s PasskeySignature) IsValid(digest common.Hash) bool {
}

func (s PasskeySignature) ImageHash() ImageHash {
	var requireUserVerification common.Hash
	if s.requireUserVerification {
		requireUserVerification[len(requireUserVerification)-1] = 1
	}

	return ImageHash{
		Hash: crypto.Keccak256Hash(
			crypto.Keccak256(s.x.Bytes(), s.y.Bytes()),
			crypto.Keccak256(requireUserVerification.Bytes(), s.metadata.Bytes()),
		),
		Preimage: s,
	}
}
