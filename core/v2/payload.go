package v2

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
)

func Approval(imageHash core.ImageHashable, address common.Address) ApprovalPayload {
	return ApprovalPayload{ImageHashable: imageHash, Address: address}
}

func Digest(digest common.Hash, address common.Address, chainID ...*big.Int) DigestPayload {
	if len(chainID) == 0 {
		chainID = []*big.Int{nil}
	}
	if chainID[0] == nil {
		chainID[0] = new(big.Int)
	}

	return DigestPayload{Hash: digest, Address: address, ChainID: chainID[0]}
}

type ApprovalPayload struct {
	core.ImageHashable

	Address common.Address
}

type DigestPayload struct {
	common.Hash

	Address common.Address
	ChainID *big.Int
}

var approvalPrefix = crypto.Keccak256Hash([]byte("SetImageHash(bytes32 imageHash)"))

func ApprovalDigest(imageHash core.ImageHashable) common.Hash {
	return crypto.Keccak256Hash(approvalPrefix.Bytes(), imageHash.ImageHash().Bytes())
}

func (p ApprovalPayload) Digest() core.PayloadDigest {
	digest := Digest(ApprovalDigest(p.ImageHashable), p.Address).Digest()
	digest.Payload = p
	return digest
}

func (p DigestPayload) Digest() core.PayloadDigest {
	chainID := new(big.Int)
	if p.ChainID != nil {
		chainID.Set(p.ChainID)
	}

	return core.PayloadDigest{
		Hash: crypto.Keccak256Hash(
			[]byte{0x19, 0x01},
			common.BigToHash(chainID).Bytes(),
			p.Address.Bytes(),
			p.Bytes(),
		),
		Address: p.Address,
		ChainID: chainID,
		Payload: p,
	}
}
