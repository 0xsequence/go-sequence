// Sequence core primitives
//
// DecodeSignature takes raw signature data and returns a Signature.
// A Signature can Recover the WalletConfig it represents.
// A WalletConfig describes the configuration of signers that control a wallet.
package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

type Core[C WalletConfig, S Signature[C]] interface {
	// DecodeSignature takes raw signature data and returns a Signature that can Recover a WalletConfig.
	DecodeSignature(data []byte) (S, error)

	// DecodeWalletConfig takes a decoded JSON object and returns a WalletConfig.
	DecodeWalletConfig(object any) (C, error)
}

// A Signature is a decoded signature that can Recover a WalletConfig.
type Signature[C WalletConfig] interface {
	// Threshold is the minimum signing weight required for a signature to be valid.
	Threshold() uint16

	// Checkpoint is the nonce of the wallet configuration that the signature applies to.
	Checkpoint() uint32

	// Recover derives the wallet configuration that the signature applies to.
	// Also returns the signature's weight.
	// If chainID is not provided, provider must be provided.
	// If provider is not provided, EIP-1271 signatures are assumed to be valid.
	// If signerSignatures is provided, it will be populated with the valid signer signatures of this signature.
	Recover(ctx context.Context, digest Digest, wallet common.Address, chainID *big.Int, provider *ethrpc.Provider, signerSignatures ...SignerSignatures) (C, *big.Int, error)

	// Data is the raw signature data.
	Data() ([]byte, error)

	// Write writes the raw signature data to a Writer.
	Write(writer io.Writer) error
}

// A WalletConfig is a configuration of signers that control a wallet.
type WalletConfig interface {
	ImageHashable

	// Threshold is the minimum signing weight required for a signature to be valid.
	Threshold() uint16

	// Checkpoint is the nonce of the wallet configuration.
	Checkpoint() uint32

	// Signers is the set of signers in the wallet configuration.
	Signers() map[common.Address]uint16

	// IsUsable checks if it's possible to construct signatures that meet threshold.
	IsUsable() error
}

// A SignerSignatures object stores signer signatures indexed by signer and subdigest.
type SignerSignatures map[common.Address]map[common.Hash]SignerSignature

func (s SignerSignatures) Insert(signer common.Address, signature SignerSignature) {
	if s == nil {
		return
	}

	signerSignatures := s[signer]
	if signerSignatures == nil {
		signerSignatures = map[common.Hash]SignerSignature{}
		s[signer] = signerSignatures
	}

	signerSignatures[signature.Subdigest.Hash] = signature
}

type SignerSignatureType int

const (
	SignerSignatureTypeEIP712 SignerSignatureType = iota
	SignerSignatureTypeEthSign
	SignerSignatureTypeEIP1271
)

type SignerSignature struct {
	Subdigest Subdigest
	Type      SignerSignatureType
	Signature []byte
}

type SigningFunction func(ctx context.Context, signer common.Address) (SignerSignatureType, []byte, error)

// An ImageHashable is an object with an ImageHash.
type ImageHashable interface {
	// ImageHash is the digest of the object.
	ImageHash() ImageHash
}

// An ImageHash is a digest of an ImageHashable.
// Used for type safety and preimage recovery.
type ImageHash struct {
	common.Hash

	// Preimage is the ImageHashable with this ImageHash, nil if unknown.
	Preimage ImageHashable
}

var imageHashApprovalSalt = crypto.Keccak256Hash([]byte("SetImageHash(bytes32 imageHash)"))

// Approval derives the digest that must be signed to approve the ImageHash for subsequent signatures.
func (h ImageHash) Approval() Digest {
	return NewDigest(imageHashApprovalSalt.Bytes(), h.Bytes())
}

// A Digest is a hash signed by a Sequence wallet.
// Used for type safety and preimage recovery.
type Digest struct {
	common.Hash

	// Preimage is the preimage of the digest, nil if unknown.
	Preimage []byte
}

// NewDigest creates a Digest from a list of messages.
func NewDigest(messages ...[]byte) Digest {
	preimage := bytes.Join(messages, nil)

	return Digest{
		Hash:     crypto.Keccak256Hash(preimage),
		Preimage: preimage,
	}
}

// ApprovedImageHash recovers the ImageHash that the Digest approves for subsequent signatures, if known.
func (d Digest) ApprovedImageHash() (ImageHash, error) {
	if !bytes.HasPrefix(d.Preimage, imageHashApprovalSalt.Bytes()) || len(d.Preimage) != len(imageHashApprovalSalt.Bytes())+common.HashLength {
		return ImageHash{}, fmt.Errorf(`preimage %v of %v is not an image hash approval`, hexutil.Encode(d.Preimage), d.Hash)
	}

	return ImageHash{Hash: common.BytesToHash(d.Preimage[len(imageHashApprovalSalt.Bytes()):])}, nil
}

// Subdigest derives the hash to be signed by the Sequence wallet's signers to validate the digest.
func (d Digest) Subdigest(wallet common.Address, chainID ...*big.Int) Subdigest {
	if len(chainID) == 0 {
		chainID = []*big.Int{nil}
	}
	if chainID[0] == nil {
		chainID[0] = new(big.Int)
	}

	return Subdigest{
		Hash:    crypto.Keccak256Hash([]byte{0x19, 0x01}, common.BigToHash(chainID[0]).Bytes(), wallet.Bytes(), d.Bytes()),
		Digest:  &d,
		Wallet:  &wallet,
		ChainID: chainID[0],
	}
}

// A Subdigest is a hash signed by a Sequence wallet's signers.
// Used for type safety and preimage recovery.
type Subdigest struct {
	common.Hash

	// Digest is the preimage of the subdigest, nil if unknown.
	Digest *Digest

	// Wallet is the target wallet of the subdigest, nil if unknown.
	Wallet *common.Address

	// ChainID is the target chain ID of the subdigest, nil if unknown.
	ChainID *big.Int

	// EthSignPreimage is the preimage of the eth_sign subdigest, nil if unknown.
	EthSignPreimage *Subdigest
}

// EthSignSubdigest derives the eth_sign subdigest of a subdigest.
func (s Subdigest) EthSignSubdigest() Subdigest {
	return Subdigest{
		Hash:            crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n"), []byte(fmt.Sprintf("%v", len(s.Bytes()))), s.Bytes()),
		EthSignPreimage: &s,
	}
}
