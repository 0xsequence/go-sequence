// Sequence core primitives
//
// DecodeSignature takes raw signature data and returns a Signature.
// A Signature can Recover the WalletConfig it represents.
// A WalletConfig describes the configuration of signers that control a wallet.
package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
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
	Checkpoint() uint64

	// Recover derives the wallet configuration that the signature applies to.
	// Also returns the signature's weight.
	// If provider is not provided, EIP-1271 signatures are assumed to be NOT valid and ignored.
	// If provider is not provided and the signature contains sapient signer signatures, recovery will fail.
	// If payload is a digest without preimage and the signature contains non-compact sapient signer signatures, recovery will fail.
	// If signerSignatures is provided, it will be populated with the valid signer signatures of this signature.
	Recover(ctx context.Context, payload Payload, provider *ethrpc.Provider, signerSignatures ...map[Signer]SignerSignature) (C, *big.Int, error)

	// Reduce returns an equivalent optimized signature.
	Reduce(payload Payload) Signature[C]

	// Join joins two signatures into one.
	Join(payload Payload, other Signature[C]) (Signature[C], error)

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
	Checkpoint() uint64

	// Signers is the set of signers in the wallet configuration.
	Signers() map[Signer]uint16

	// SignersWeight is the total weight of the signers passed to the function according to the wallet configuration.
	SignersWeight(signers map[Signer]uint16) uint16

	// IsComplete checks if the wallet configuration doesn't have pruned subtrees.
	IsComplete() bool

	// IsUsable checks if it's possible to construct signatures that meet threshold.
	IsUsable() error
}

type Signer struct {
	Address   common.Address
	IsSapient bool
	ImageHash common.Hash
}

func SapientSigner(address common.Address, imageHash common.Hash) Signer {
	return Signer{Address: address, IsSapient: true, ImageHash: imageHash}
}

func SignerFromString(s string) (Signer, error) {
	signerLength := 2 + 2*common.AddressLength
	sapientSignerLength := signerLength + 1 + 2 + 2*common.HashLength

	switch len(s) {
	case signerLength:
		if !common.IsHexAddress(s) {
			return Signer{}, fmt.Errorf(`"%v" is not a signer`, s)
		}
		return Signer{Address: common.HexToAddress(s)}, nil

	case sapientSignerLength:
		if !common.IsHexAddress(s[:signerLength]) {
			return Signer{}, fmt.Errorf(`"%v" is not a signer`, s)
		}
		if s[signerLength] != ':' {
			return Signer{}, fmt.Errorf(`"%v" is not a signer`, s)
		}
		imageHash, err := hexutil.Decode(s[signerLength+1:])
		if err != nil {
			return Signer{}, fmt.Errorf(`"%v" is not a signer`, s)
		}
		return SapientSigner(common.HexToAddress(s[:signerLength]), common.BytesToHash(imageHash)), nil

	default:
		return Signer{}, fmt.Errorf(`"%v" is not a signer`, s)
	}
}

func (s Signer) String() string {
	if s.IsSapient {
		return fmt.Sprintf("%v:%v", s.Address, s.ImageHash)
	} else {
		return s.Address.String()
	}
}

type SignerSignatureType int

const (
	SignerSignatureTypeEIP712 SignerSignatureType = iota + 1
	SignerSignatureTypeEthSign
	SignerSignatureTypeEIP1271
	SignerSignatureTypeSapient
	SignerSignatureTypeSapientCompact
)

type SignerSignature struct {
	Signer    Signer
	Type      SignerSignatureType
	Signature []byte
	Error     error
}

// ErrSigningFunctionNotReady is returned when a signing function is not ready to sign and should be retried later.
var ErrSigningFunctionNotReady = fmt.Errorf("signing function not ready")
var ErrSigningNoSigner = fmt.Errorf("no signer")

type SigningFunction func(ctx context.Context, signer Signer, signatures []SignerSignature) (SignerSignatureType, []byte, error)

func SigningOrchestrator(ctx context.Context, signers map[Signer]uint16, sign SigningFunction) chan SignerSignature {
	signaturesChan := make(chan SignerSignature)
	go func() {
		defer close(signaturesChan)

		wg := sync.WaitGroup{}

		var cond = sync.NewCond(&sync.Mutex{})
		var signatures = make([]SignerSignature, 0, len(signers))
		var signaturesExpected = len(signers)
		for signer := range signers {
			wg.Add(1)
			go func(signer Signer) {
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
						} else if errors.Is(err, ErrSigningNoSigner) {
							err = nil
						}
					}

					signerSignature := SignerSignature{
						Signer:    signer,
						Type:      signatureType,
						Signature: signature,
						Error:     err,
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

func (h ImageHash) ImageHash() ImageHash {
	return h
}

var imageHashApprovalSalt = crypto.Keccak256Hash([]byte("SetImageHash(bytes32 imageHash)"))

// Approval derives the digest that must be signed to approve the ImageHash for subsequent signatures.
func (h ImageHash) Approval() common.Hash {
	return crypto.Keccak256Hash(imageHashApprovalSalt.Bytes(), h.Bytes())
}

func EthereumSignedMessage(message []byte) common.Hash {
	return common.Hash(accounts.TextHash(message))
}

func Ecrecover(digest common.Hash, signature []byte) (common.Address, error) {
	if length := len(signature); length != crypto.SignatureLength {
		return common.Address{}, fmt.Errorf("invalid ecdsa signature length %v, expected length %v", length, crypto.SignatureLength)
	}

	signature_ := [crypto.SignatureLength]byte(signature)
	signature_[len(signature_)-1] -= 27

	pubkey, err := crypto.SigToPub(digest.Bytes(), signature_[:])
	if err != nil {
		return common.Address{}, fmt.Errorf("unable to recover ecdsa signature: %w", err)
	}

	return crypto.PubkeyToAddress(*pubkey), nil
}
