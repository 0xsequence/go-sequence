package sequence

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

var zeroAddress = common.Address{}

func mustEncodeSig(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}

// func ValidateSequenceAccountProof(factory, mainModule common.Address) ethauth.ValidatorFunc {
// 	return func(ctx context.Context, provider *ethrpc.Provider, chainID *big.Int, proof *ethauth.Proof) (bool, string, error) {
// 		if provider == nil {
// 			return false, "", fmt.Errorf("ValidateContractAccountToken failed. provider is nil")
// 		}
// 		if chainID == nil {
// 			return false, "", fmt.Errorf("ValidateContractAccountToken failed. chainID is nil")
// 		}

// 		// Compute eip712 message digest from the token claims
// 		messageDigest, err := proof.MessageDigest()
// 		if err != nil {
// 			return false, "", fmt.Errorf("ValidateEOAToken failed. Unable to compute token message digest, because %w", err)
// 		}

// 		// must hash the message as first argument to isValidSignature
// 		// messageHash := ethcoder.Keccak256(messageDigest)

// 		valid, err := IsValidSignature2(
// 			common.HexToAddress(proof.Address),
// 			messageDigest,
// 			ethcoder.MustHexDecode(proof.Signature),
// 			chainID,
// 			&factory,
// 			&mainModule,
// 			provider,
// 		)
// 		if err != nil {
// 			return false, "", err
// 		}
// 		if !valid {
// 			return false, "", nil
// 		}

// 		return true, proof.Address, nil
// 	}
// }
