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

	proto "github.com/0xsequence/go-sequence/lib/sessions"
)

// Utility functions to use with ethauth, in order to validate Sequence Wallet signatures, encoded
// as ethauth proofs and verifable on a Go backend.

func ValidateSequenceAccountProof(configTracker proto.Sessions) ethauth.ValidatorFunc {
	return ValidateSequenceAccountProofWith(sequenceContext.FactoryAddress, sequenceContext.MainModuleAddress, configTracker)
}

func ValidateSequenceAccountProofWith(factory, mainModule common.Address, configTracker proto.Sessions) ethauth.ValidatorFunc {
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
			valid, _ = IsValidSignature(
				common.HexToAddress(proof.Address),
				common.BytesToHash(messageDigest),
				sig,
				WalletContext{FactoryAddress: factory, MainModuleAddress: mainModule},
				chainID,
				provider,
				configTracker,
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
