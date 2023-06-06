// Sequence core primitives
//
// DecodeSignature takes raw signature data and returns a Signature.
// A Signature can Recover the WalletConfig it represents.
// A WalletConfig describes the configuration of signers that control a wallet.
package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

var cores []any

func RegisterCore[C WalletConfig, S Signature[C]](core Core[C, S]) {
	cores = append(cores, core)
}

func GetCoreForWalletConfig[C WalletConfig]() (Core[C, Signature[C]], error) {
	for _, core := range cores {
		if core, ok := core.(Core[C, Signature[C]]); ok {
			return core, nil
		}
	}
	return nil, fmt.Errorf("sequence: core not found")
}

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

	// Reduce returns an equivalent optimized signature.
	Reduce(subdigest Subdigest) Signature[C]

	// Join joins two signatures into one.
	Join(subdigest Subdigest, other Signature[C]) (Signature[C], error)

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
	Signer    common.Address
	Subdigest Subdigest
	Type      SignerSignatureType
	Signature []byte
}

// ErrSigningFunctionNotReady is returned when a signing function is not ready to sign and should be retried later.
var ErrSigningFunctionNotReady = fmt.Errorf("signing function not ready")

type SigningFunction func(ctx context.Context, signer common.Address, signatures []SignerSignature) (SignerSignatureType, []byte, error)

func SigningOrchestrator(ctx context.Context, signers map[common.Address]uint16, sign SigningFunction) chan SignerSignature {
	signaturesChan := make(chan SignerSignature)
	go func() {
		defer close(signaturesChan)

		wg := sync.WaitGroup{}

		var cond = sync.NewCond(&sync.Mutex{})
		var signatures = make([]SignerSignature, 0, len(signers))
		var signaturesExpected = len(signers)
		for signer := range signers {
			wg.Add(1)
			go func(signer common.Address) {
				defer wg.Done()

				go func() {
					select {
					case <-ctx.Done():
						cond.Broadcast()
					}
				}()

				retries := 0
				for {
					select {
					case <-ctx.Done():
						return
					default:
					}

					cond.L.Lock()
					signaturesArg := signatures[:]
					cond.L.Unlock()

					signatureType, signature, err := sign(ctx, signer, signaturesArg)
					if err != nil {
						if errors.Is(err, ErrSigningFunctionNotReady) && retries < 1 {
							cond.L.Lock()
							signaturesLeft := signaturesExpected - len(signatures)
							if signaturesLeft > 1 {
								cond.Wait()
							}
							cond.L.Unlock()

							retries++
							continue
						}
					}

					signerSignature := SignerSignature{
						Type:      signatureType,
						Signer:    signer,
						Signature: signature,
					}

					cond.L.Lock()
					signatures = append(signatures, signerSignature)
					cond.L.Unlock()

					signaturesChan <- signerSignature
					cond.Broadcast()
					break
				}
			}(signer)
		}

		wg.Wait()
	}()

	return signaturesChan
}

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
