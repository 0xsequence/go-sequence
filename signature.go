package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts/gen/ierc1271"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	"github.com/0xsequence/go-sequence/core/v2"
)

func Sign[C core.WalletConfig](wallet *Wallet[C], input common.Hash) ([]byte, core.Signature[C], error) {
	return wallet.SignDigest(context.Background(), input)
}

// Solves the knapsack problem
func knapsack(capacity int, weights []int, values []int) (int, []int) {
	// m[i][j] = maximum value achievable using only first i items up to capacity j
	m := [][]int{make([]int, capacity+1)}
	for range weights {
		m = append(m, make([]int, capacity+1))
	}

	for i, w := range weights {
		for j := 0; j <= capacity; j++ {
			m[i+1][j] = m[i][j]
			if j >= w && m[i+1][j] < m[i][j-w]+values[i] {
				m[i+1][j] = m[i][j-w] + values[i]
			}
		}
	}

	var s []int
	j := capacity
	for i := len(weights) - 1; i >= 0; i-- {
		if m[i+1][j] != m[i][j] {
			s = append(s, i)
			j -= weights[i]
		}
	}

	return m[len(weights)][capacity], s
}

// DecodeSignature sequence into individual parts
func DecodeSignature[C core.WalletConfig](sequenceSignature []byte) (core.Signature[C], error) {
	sig := sequenceSignature

	s := len(sig)
	if s < 2 {
		return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds threshold")
	}

	var typeOfConfig C
	var generalTypeOfConfig core.WalletConfig = typeOfConfig
	if _, ok := generalTypeOfConfig.(*v2.WalletConfig); ok {
		decodedSignature, err := v2.Core.DecodeSignature(sig)
		if err != nil {
			return nil, fmt.Errorf("sequence: invalid signature, %v", err)
		}
		return decodedSignature.(core.Signature[C]), nil
	} else if _, ok := generalTypeOfConfig.(*v1.WalletConfig); ok {
		decodedSignature, err := v1.Core.DecodeSignature(sig)
		if err != nil {
			return nil, fmt.Errorf("sequence: invalid signature, %v", err)
		}
		return decodedSignature.(core.Signature[C]), nil
	} else {
		return nil, fmt.Errorf("sequence: unknown config type")
	}
}

func RecoverWalletConfigFromDigest[C core.WalletConfig](digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (C, *big.Int, error) {
	var (
		wc     C
		weight *big.Int
		err    error
	)

	decoded, err := DecodeSignature[C](seqSig)
	if err != nil {
		return wc, weight, err
	}

	wc, weight, err = decoded.Recover(context.Background(), core.Digest{Hash: common.BytesToHash(digest)}, walletAddress, chainID, provider)
	if err != nil {
		return wc, weight, err
	}

	return wc, weight, nil
}

func IsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContexts WalletContexts, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	// Try to do it first with ethereum sign signature format
	ok, err := ethwallet.IsValid191Signature(walletAddress, digest[:], seqSig)
	if err == nil {
		return ok, nil
	}

	var code []byte
	if provider != nil {
		code, err = provider.CodeAt(context.Background(), walletAddress, nil)
		if err != nil {
			return false, err
		}
	}

	if len(code) == 0 {
		// It may be a signature from a non-deployed sequence wallet, check and attempt to validate
		// first we try for a v2 wallet, assuming that the context is a valid v2 wallet context
		// if it's not then we can safely assume that it's a v1 wallet context (or invalid)
		res2, err2 := IsValidV2UndeployedSignature(walletAddress, digest, seqSig, walletContexts[2], chainID, provider)
		if err2 == nil && res2 {
			return true, nil
		}

		subDigest, err := SubDigest(chainID, walletAddress, digest)
		if err != nil {
			return false, err
		}

		res1, err1 := IsValidV1UndeployedSignature(walletAddress, subDigest, seqSig, walletContexts[1], chainID, provider)
		if err1 == nil && res1 {
			return true, nil
		}

		return false, fmt.Errorf("failed to validate: %v, %v", err1, err2)

	} else {
		// Smart wallet is deployed, query erc1271 method to check signature
		erc1271, err := ierc1271.NewIERC1271(walletAddress, provider)
		if err != nil {
			return false, err
		}

		// NOTE: we expect digest to be ready for ERC1271 call, so digest must be fully encoded as expected
		res, err := erc1271.IsValidSignature(&bind.CallOpts{From: common.Address{0x1}}, digest, seqSig)
		if err != nil {
			return false, err
		}

		if ierc1271.IsValidSignatureBytes32_MagicReturnValue != hexutil.Encode(res[:]) {
			return false, fmt.Errorf("failed to validate")
		}
	}

	return true, nil
}

// TODO: Use core methods
func IsValidV1UndeployedSignature(walletAddress common.Address, subDigest []byte, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	recoveredWalletConfig, weight, err := RecoverWalletConfigFromDigest[*v1.WalletConfig](subDigest, seqSig, walletAddress, walletContext, chainID, provider)
	if err != nil {
		return false, err
	}

	recoveredAddress, err := AddressFromWalletConfig(recoveredWalletConfig, walletContext)
	if err != nil {
		return false, err
	}

	if recoveredAddress != walletAddress || weight.Cmp(big.NewInt(int64(recoveredWalletConfig.Threshold()))) < 0 {
		return false, fmt.Errorf("failed to validate")
	}

	return true, nil
}

func IsValidV2UndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	decoded, err := v2.Core.DecodeSignature(seqSig)
	if err != nil {
		return false, err
	}

	recovered, weight, err := decoded.Recover(context.Background(), core.Digest{Hash: digest}, walletAddress, chainID, provider)
	if err != nil {
		return false, err
	}

	imageHash := recovered.ImageHash()
	address, err := AddressFromImageHash(imageHash.Hex(), walletContext)
	if err != nil {
		return false, err
	}

	if address != walletAddress || weight.Uint64() < uint64(recovered.Threshold_) {
		return false, fmt.Errorf("failed to validate")
	}

	return true, nil
}

func MessageDigest(message []byte) common.Hash {
	return common.BytesToHash(ethcoder.Keccak256(message))
}

func MustEncodeSig(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}
