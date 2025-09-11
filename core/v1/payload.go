package v1

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
)

func Approval(imageHash core.ImageHashable, address common.Address) ApprovalPayload {
	return ApprovalPayload{ImageHashable: imageHash, address: address}
}

func Digest(digest common.Hash, address common.Address, chainID ...*big.Int) DigestPayload {
	if len(chainID) == 0 {
		chainID = []*big.Int{nil}
	}
	if chainID[0] == nil {
		chainID[0] = new(big.Int)
	}

	return DigestPayload{Hash: digest, address: address, chainID: chainID[0]}
}

type ApprovalPayload struct {
	core.ImageHashable

	address common.Address
}

func (p ApprovalPayload) Address() common.Address {
	return p.address
}

func (p ApprovalPayload) ChainID() *big.Int {
	return new(big.Int)
}

type DigestPayload struct {
	common.Hash

	address common.Address
	chainID *big.Int
}

func (p DigestPayload) Address() common.Address {
	return p.address
}

func (p DigestPayload) ChainID() *big.Int {
	if p.chainID != nil {
		return new(big.Int).Set(p.chainID)
	} else {
		return new(big.Int)
	}
}

var approvalPrefix = crypto.Keccak256Hash([]byte("SetImageHash(bytes32 imageHash)"))

func ApprovalDigest(imageHash core.ImageHashable) common.Hash {
	return crypto.Keccak256Hash(approvalPrefix.Bytes(), imageHash.ImageHash().Bytes())
}

func (p ApprovalPayload) Digest() core.PayloadDigest {
	digest := Digest(ApprovalDigest(p.ImageHashable), p.address).Digest()
	digest.Payload = p
	return digest
}

func (p DigestPayload) Digest() core.PayloadDigest {
	chainID := new(big.Int)
	if p.chainID != nil {
		chainID.Set(p.chainID)
	}

	return core.PayloadDigest{
		Hash: crypto.Keccak256Hash(
			[]byte{0x19, 0x01},
			common.BigToHash(chainID).Bytes(),
			p.address.Bytes(),
			p.Bytes(),
		),
		Address_: p.address,
		ChainID_: chainID,
		Payload:  p,
	}
}
