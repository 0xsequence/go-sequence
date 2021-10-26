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
	"github.com/goware/breaker"
)

// Utility functions to use with ethauth, in order to validate Sequence Wallet signatures, encoded
// as ethauth proofs and verifable on a Go backend.

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

		// Auto-retry validation a number of times as it make take a node to sync with the latest state
		var valid bool

		err = breaker.Do(ctx, func() error {
			valid, err = IsValidSignature(
				common.HexToAddress(proof.Address),
				common.BytesToHash(messageDigest),
				ethcoder.MustHexDecode(proof.Signature),
				WalletContext{FactoryAddress: factory, MainModuleAddress: mainModule},
				chainID,
				provider,
			)
			return err
		}, nil, 1*time.Second, 1.5, 4)

		if !valid || err != nil {
			return false, "", fmt.Errorf("failed to validate")
		}

		return true, proof.Address, nil
	}
}
