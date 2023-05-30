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
	v2 "github.com/0xsequence/go-sequence/core/v2"
)

func Sign[C core.WalletConfig](wallet *Wallet[C], input common.Hash) ([]byte, core.Signature[C], error) {
	return wallet.SignDigest(context.Background(), input)
}

// GenericDecodeSignature sequence into individual parts
func GenericDecodeSignature[C core.WalletConfig](sequenceSignature []byte) (core.Signature[C], error) {
	sig := sequenceSignature

	s := len(sig)
	if s < 2 {
		return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds threshold")
	}

	c, err := core.GetCoreForWalletConfig[C]()
	if err != nil {
		return nil, fmt.Errorf("sequence: failed to get core implementation, %v", err)
	}

	decodedSignature, err := c.DecodeSignature(sig)
	if err != nil {
		return nil, fmt.Errorf("sequence: invalid signature, %v", err)
	}
	return decodedSignature, nil
}

func V1DecodeSignature(sequenceSignature []byte) (core.Signature[*v1.WalletConfig], error) {
	return GenericDecodeSignature[*v1.WalletConfig](sequenceSignature)
}

func V2DecodeSignature(sequenceSignature []byte) (core.Signature[*v2.WalletConfig], error) {
	return GenericDecodeSignature[*v2.WalletConfig](sequenceSignature)
}

func DecodeSignature(sequenceSignature []byte) (core.Signature[*v2.WalletConfig], error) {
	return V2DecodeSignature(sequenceSignature)
}

func GenericRecoverWalletConfigFromDigest[C core.WalletConfig](digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (C, *big.Int, error) {
	var (
		wc     C
		weight *big.Int
		err    error
	)

	decoded, err := GenericDecodeSignature[C](seqSig)
	if err != nil {
		return wc, weight, err
	}

	wc, weight, err = decoded.Recover(context.Background(), core.Digest{Hash: common.BytesToHash(digest)}, walletAddress, chainID, provider)
	if err != nil {
		return wc, weight, err
	}

	return wc, weight, nil
}

func V1RecoverWalletConfigFromDigest(digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (*v1.WalletConfig, *big.Int, error) {
	return GenericRecoverWalletConfigFromDigest[*v1.WalletConfig](digest, seqSig, walletAddress, walletContext, chainID, provider)
}

func V2RecoverWalletConfigFromDigest(digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (*v2.WalletConfig, *big.Int, error) {
	return GenericRecoverWalletConfigFromDigest[*v2.WalletConfig](digest, seqSig, walletAddress, walletContext, chainID, provider)
}

func RecoverWalletConfigFromDigest(digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (*v2.WalletConfig, *big.Int, error) {
	return V2RecoverWalletConfigFromDigest(digest, seqSig, walletAddress, walletContext, chainID, provider)
}

func GenericIsValidSignature[C core.WalletConfig](walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
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
		res, err := GenericIsValidUndeployedSignature[C](walletAddress, digest, seqSig, walletContext, chainID, provider)
		if err == nil && res {
			return true, nil
		}

		return false, fmt.Errorf("failed to validate: %v", err)
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

func V1IsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidSignature[*v1.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func V2IsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidSignature[*v2.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func IsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return V2IsValidSignature(walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func GenericIsValidUndeployedSignature[C core.WalletConfig](walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	c, err := core.GetCoreForWalletConfig[C]()
	if err != nil {
		return false, fmt.Errorf("sequence: failed to get core implementation, %v", err)
	}

	decoded, err := c.DecodeSignature(seqSig)
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

	if address != walletAddress || weight.Uint64() < uint64(recovered.Threshold()) {
		return false, fmt.Errorf("failed to validate")
	}

	return true, nil
}

func V1IsValidUndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidUndeployedSignature[*v1.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func V2IsValidUndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidUndeployedSignature[*v2.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func IsValidUndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return V2IsValidUndeployedSignature(walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func MessageDigest(message []byte) common.Hash {
	return common.BytesToHash(ethcoder.Keccak256(message))
}

func MustEncodeSig(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}
