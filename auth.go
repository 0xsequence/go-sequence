package sequence

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-ethauth"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	"github.com/0xsequence/go-sequence/eip6492"
	"github.com/goware/logger"
)

// Utility functions to use with ethauth, in order to validate Sequence Wallet signatures, encoded
// as ethauth proofs and verifable on a Go backend.

func GenericValidateSequenceAccountProof[C core.WalletConfig]() ethauth.ValidatorFunc {
	var walletConfig C
	return GenericValidateSequenceAccountProofWith[C](SequenceContextForWalletConfig(walletConfig))
}

func V1ValidateSequenceAccountProof() ethauth.ValidatorFunc {
	return GenericValidateSequenceAccountProofWith[*v1.WalletConfig](V1SequenceContext())
}

func V2ValidateSequenceAccountProof() ethauth.ValidatorFunc {
	return GenericValidateSequenceAccountProofWith[*v2.WalletConfig](V2SequenceContext())
}

func ValidateSequenceAccountProof(log logger.Logger) ethauth.ValidatorFunc {
	eip6492Validator := EIP6492ValidateSignature()
	legacyValidator := V2ValidateSequenceAccountProof()

	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
		// NOTICE: We only need the legacy validator so we can deploy these API updates
		// without having to pause the wallet, otherwise we need to coordinate both.
		// as soon as we have EIP-6492 running on prod we can remove this.
		eip6492Valid, _, _ := eip6492Validator(ctx, provider, chainID, proof)
		if eip6492Valid {
			return true, proof.Address, nil
		}

		legacyIsValid, addr, err := legacyValidator(ctx, provider, chainID, proof)
		if err != nil {
			return false, addr, err
		}

		if legacyIsValid {
			log.Errorf("WARNING: Legacy signature validation used, please upgrade to EIP-6492: %s, %s", proof.Address, proof.Signature)
		}

		return legacyIsValid, addr, nil
	}
}

func GeneralValidateSequenceAccountProof() ethauth.ValidatorFunc {
	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
		validatorFuncV1 := GenericValidateSequenceAccountProof[*v1.WalletConfig]()
		validatorFuncV2 := GenericValidateSequenceAccountProof[*v2.WalletConfig]()

		valid1, address1, err1 := validatorFuncV2(ctx, provider, chainID, proof)
		if valid1 && err1 == nil {
			return valid1, address1, nil
		}

		valid2, address2, err2 := validatorFuncV1(ctx, provider, chainID, proof)
		if valid2 && err2 == nil {
			return valid2, address2, nil
		}

		return false, "", fmt.Errorf("ValidateSequenceAccountProof failed. %v, %v", err1, err2)
	}
}

func EIP6492ValidateSignature() ethauth.ValidatorFunc {
	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
		if provider == nil {
			return false, "", fmt.Errorf("EIP6492ValidateSignature failed. provider is nil")
		}

		messageDigest, err := proof.MessageDigest()
		if err != nil {
			return false, "", fmt.Errorf("EIP6492ValidateSignature failed. Unable to compute token message digest, because %w", err)
		}

		if !common.IsHexAddress(proof.Address) {
			return false, "", fmt.Errorf(`"%v" is not a valid address`, proof.Address)
		}

		hash := common.BytesToHash(messageDigest)
		signer := common.HexToAddress(proof.Address)
		sig, err := ethcoder.HexDecode(proof.Signature)
		if err != nil {
			return false, "", fmt.Errorf("sig is invalid: %w", err)
		}

		isValid, err := eip6492.ValidateEIP6492Offchain(provider, signer, hash, sig)
		if err != nil {
			return false, "", fmt.Errorf("failed to validate: %w", err)
		}

		return isValid, proof.Address, nil
	}
}

func GenericValidateSequenceAccountProofWith[C core.WalletConfig](walletContexts WalletContext) ethauth.ValidatorFunc {
	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
		if provider == nil {
			return false, "", fmt.Errorf("ValidateContractAccountToken failed. provider is nil")
		}
		if chainID == nil {
			return false, "", fmt.Errorf("ValidateContractAccountToken failed. chainID is nil")
		}

		// Compute eip712 message digest from the token claims
		messageDigest, err := proof.MessageDigest()
		if err != nil {
			return false, "", fmt.Errorf("ValidateEOAToken failed. Unable to compute token message digest, because %w", err)
		}

		sig, err := ethcoder.HexDecode(proof.Signature)
		if err != nil {
			return false, "", fmt.Errorf("sig is invalid")
		}

		// Auto-retry validation a number of times as it might take a node to sync with the latest state
		var valid bool
		numAttempts := 4

		for i := 1; i <= numAttempts; i++ {
			valid, _ = GenericIsValidSignature[C](
				common.HexToAddress(proof.Address),
				common.BytesToHash(messageDigest),
				sig,
				walletContexts,
				chainID,
				provider,
			)
			if valid {
				break
			}
			time.Sleep(time.Duration(i) * 1500 * time.Millisecond)
		}

		if !valid {
			return false, "", fmt.Errorf("failed to validate")
		}

		return true, proof.Address, nil
	}
}

func V1ValidateSequenceAccountProofWith(walletContext WalletContext) ethauth.ValidatorFunc {
	return GenericValidateSequenceAccountProofWith[*v1.WalletConfig](walletContext)
}

func V2ValidateSequenceAccountProofWith(walletContext WalletContext) ethauth.ValidatorFunc {
	return GenericValidateSequenceAccountProofWith[*v2.WalletConfig](walletContext)
}

func ValidateSequenceAccountProofWith(walletContext WalletContext) ethauth.ValidatorFunc {
	return V2ValidateSequenceAccountProofWith(walletContext)
}

func GeneralValidateSequenceAccountProofWith(walletContexts WalletContexts) ethauth.ValidatorFunc {
	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
		validatorFuncV1 := GenericValidateSequenceAccountProofWith[*v1.WalletConfig](walletContexts[1])
		validatorFuncV2 := GenericValidateSequenceAccountProofWith[*v2.WalletConfig](walletContexts[2])

		valid1, address1, err1 := validatorFuncV2(ctx, provider, chainID, proof)
		if valid1 && err1 == nil {
			return valid1, address1, nil
		}

		valid2, address2, err2 := validatorFuncV1(ctx, provider, chainID, proof)
		if valid2 && err2 == nil {
			return valid2, address2, nil
		}

		return false, "", fmt.Errorf("ValidateSequenceAccountProof failed. %v, %v", err1, err2)
	}
}
