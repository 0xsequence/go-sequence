package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-ethauth"
)

var zeroAddress = common.Address{}

func ValidateSequenceAccountProof() ethauth.ValidatorFunc {
	return ValidateSequenceAccountProofWith(sequenceContext.FactoryAddress, sequenceContext.MainModuleAddress)
}

func ValidateSequenceAccountProofWith(factory, mainModule common.Address) ethauth.ValidatorFunc {
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

		// TODO: do we need a subdigest of this message?... or shall we remove below..?
		// confirm with a test..
		// must hash the message as first argument to isValidSignature
		// messageHash := ethcoder.Keccak256(messageDigest)

		valid, err := IsValidSignature(
			common.HexToAddress(proof.Address),
			messageDigest,
			ethcoder.MustHexDecode(proof.Signature),
			WalletContext{FactoryAddress: factory, MainModuleAddress: mainModule},
			chainID,
			provider,
		)
		if err != nil {
			return false, "", err
		}
		if !valid {
			return false, "", nil
		}

		return true, proof.Address, nil
	}
}
