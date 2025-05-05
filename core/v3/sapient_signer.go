package v3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/eip6492"
)

func RecoverSapientSignature(ctx context.Context, signer common.Address, payload Payload, signature []byte, provider *ethrpc.Provider) (core.ImageHash, error) {
	if provider == nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature without provider")
	}

	data, err := contracts.V3.WalletStage1Module.Encode("recoverSapientSignature", payload.ABIEncode(), signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignature call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature: %w", err)
	}

	var imageHash common.Hash
	err = contracts.V3.WalletStage1Module.Decode(&imageHash, "recoverSapientSignature", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignature response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}

func RecoverSapientSignatureCompact(ctx context.Context, signer common.Address, payload PayloadDigestable, signature []byte, provider *ethrpc.Provider) (core.ImageHash, error) {
	if provider == nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact without provider")
	}

	data, err := contracts.V3.WalletStage1Module.Encode("recoverSapientSignatureCompact", payload.Digest().Hash, signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignatureCompact call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact: %w", err)
	}

	var imageHash common.Hash
	err = contracts.V3.WalletStage1Module.Decode(&imageHash, "recoverSapientSignatureCompact", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignatureCompact response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}

type ValidationResult struct {
	IsValid         bool
	ImageHash       common.Hash
	EffectiveWeight uint8
	Error           error
}

func ValidateSapientSignature(
	ctx context.Context,
	provider *ethrpc.Provider,
	signer common.Address,
	subdigest core.Subdigest,
	signature []byte,
	weight uint8,
) ValidationResult {
	if provider == nil {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("provider required for Sapient signature validation"),
		}
	}

	isValid, err := eip6492.ValidateEIP6492Offchain(ctx, provider, signer, subdigest.Hash, signature, nil)
	if err != nil {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("failed to validate Sapient signature: %w", err),
		}
	}

	if !isValid {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("invalid Sapient signature for %v", signer),
		}
	}

	sigToValidate := signature
	if eip6492.IsEIP6492Signature(signature) {
		_, _, sigToValidate, err = eip6492.DecodeEIP6492Signature(signature)
		if err != nil {
			return ValidationResult{
				IsValid:         false,
				ImageHash:       common.Hash{},
				EffectiveWeight: 0,
				Error:           fmt.Errorf("failed to decode EIP-6492 signature: %w", err),
			}
		}
	}

	imageHash := crypto.Keccak256Hash(
		[]byte("Sequence sapient config:\n"),
		signer.Bytes(),
		new(big.Int).SetUint64(uint64(weight)).Bytes(),
		crypto.Keccak256(sigToValidate),
	)

	return ValidationResult{
		IsValid:         true,
		ImageHash:       imageHash,
		EffectiveWeight: weight,
		Error:           nil,
	}
}

func ValidateSapientCompactSignature(
	ctx context.Context,
	provider *ethrpc.Provider,
	signer common.Address,
	subdigest core.Subdigest,
	signature []byte,
	weight uint8,
) ValidationResult {
	if provider == nil {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("provider required for Sapient Compact signature validation"),
		}
	}

	isValid, err := eip6492.ValidateEIP6492Offchain(ctx, provider, signer, subdigest.Hash, signature, nil)
	if err != nil {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("failed to validate Sapient Compact signature: %w", err),
		}
	}

	if !isValid {
		return ValidationResult{
			IsValid:         false,
			ImageHash:       common.Hash{},
			EffectiveWeight: 0,
			Error:           fmt.Errorf("invalid Sapient Compact signature for %v", signer),
		}
	}

	sigToValidate := signature
	if eip6492.IsEIP6492Signature(signature) {
		_, _, sigToValidate, err = eip6492.DecodeEIP6492Signature(signature)
		if err != nil {
			return ValidationResult{
				IsValid:         false,
				ImageHash:       common.Hash{},
				EffectiveWeight: 0,
				Error:           fmt.Errorf("failed to decode EIP-6492 signature: %w", err),
			}
		}
	}

	imageHash := crypto.Keccak256Hash(
		[]byte("Sequence sapient config:\n"),
		signer.Bytes(),
		new(big.Int).SetUint64(uint64(weight)).Bytes(),
		crypto.Keccak256(sigToValidate),
	)

	return ValidationResult{
		IsValid:         true,
		ImageHash:       imageHash,
		EffectiveWeight: weight,
		Error:           nil,
	}
}
