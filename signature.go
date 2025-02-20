package sequence

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/ierc1271"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/eip6492"
	"github.com/goware/logger"
)

func Sign[C core.WalletConfig](wallet *Wallet[C], input common.Hash) ([]byte, error) {
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

func V3DecodeSignature(sequenceSignature []byte) (core.Signature[*v3.WalletConfig], error) {
	return GenericDecodeSignature[*v3.WalletConfig](sequenceSignature)
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

func V3RecoverWalletConfigFromDigest(digest, seqSig []byte, walletAddress common.Address, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (*v3.WalletConfig, *big.Int, error) {
	return GenericRecoverWalletConfigFromDigest[*v3.WalletConfig](digest, seqSig, walletAddress, walletContext, chainID, provider)
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

	if provider == nil {
		return false, fmt.Errorf("failed to validate: provider is required")
	}

	code, err := provider.CodeAt(context.TODO(), walletAddress, nil)
	if err != nil {
		return false, fmt.Errorf("unable to get code at %v: %w", walletAddress, err)
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

func V3IsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidSignature[*v3.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func GeneralIsValidSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContexts WalletContexts, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	isValid, err3 := V3IsValidSignature(walletAddress, digest, seqSig, walletContexts[3], chainID, provider)
	if err3 == nil {
		return isValid, nil
	}

	isValid, err2 := V2IsValidSignature(walletAddress, digest, seqSig, walletContexts[2], chainID, provider)
	if err2 == nil {
		return isValid, nil
	}

	isValid, err := V1IsValidSignature(walletAddress, digest, seqSig, walletContexts[1], chainID, provider)
	if err != nil {
		return false, fmt.Errorf("failed to validate: %v, %v", err, err2)
	}

	return isValid, nil
}

func IsValidMessageSignature(address common.Address, message []byte, signature []byte, chainID *big.Int, provider *ethrpc.Provider, optLogger *logger.Logger) (bool, error) {
	log := logger.Nop()
	if optLogger != nil {
		log = *optLogger
	}

	isValid, err := ethwallet.IsValid191Signature(address, message, signature)
	if err == nil && isValid {
		return true, nil
	}

	return IsValidSignature(log, address, common.BytesToHash(accounts.TextHash(message)), signature, SequenceContexts(), chainID, provider)
}

func IsValidTypedDataSignature(address common.Address, encodedTypedData []byte, signature []byte, chainID *big.Int, provider *ethrpc.Provider, optLogger *logger.Logger) (bool, error) {
	log := logger.Nop()
	if optLogger != nil {
		log = *optLogger
	}

	isValid, err := ethwallet.IsValid191Signature(address, encodedTypedData, signature)
	if err == nil && isValid {
		return true, nil
	}

	typedDataDigest := common.BytesToHash(ethcoder.Keccak256(encodedTypedData))
	return IsValidSignature(log, address, typedDataDigest, signature, SequenceContexts(), chainID, provider)
}

func IsValidSignature(log logger.Logger, walletAddress common.Address, digest common.Hash, seqSig []byte, walletContexts WalletContexts, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	eip6492isValid, _ := eip6492.ValidateEIP6492Offchain(context.Background(), provider, walletAddress, digest, seqSig, nil)
	if eip6492isValid {
		return true, nil
	}

	// NOTICE: This is legacy code, we only need it while we deploy EIP-6492
	// on prod, otherwise we need to coordinate both.
	// as soon as we have EIP-6492 running on prod we can remove this.

	var err error
	seqSig, err = UnwrapEIP6492Signature(seqSig)
	if err != nil {
		return false, err
	}

	generalIsValid, err := GeneralIsValidSignature(walletAddress, digest, seqSig, walletContexts, chainID, provider)
	if err != nil {
		return false, err
	}

	if generalIsValid {
		log.Errorf("WARNING: Legacy signature validation used, please upgrade to EIP-6492: %s, %s, %s", walletAddress, digest, common.Bytes2Hex(seqSig))
	}

	return generalIsValid, nil
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

func V3IsValidUndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	return GenericIsValidUndeployedSignature[*v3.WalletConfig](walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func IsValidUndeployedSignature(walletAddress common.Address, digest common.Hash, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	isValid, err := V3IsValidUndeployedSignature(walletAddress, digest, seqSig, walletContext, chainID, provider)
	if err == nil {
		return isValid, nil
	}

	return V2IsValidUndeployedSignature(walletAddress, digest, seqSig, walletContext, chainID, provider)
}

func MessageDigest(message []byte) common.Hash {
	return common.BytesToHash(ethcoder.Keccak256(message))
}

func MustEncodeSig(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}

func EIP6492Signature(signature []byte, config core.WalletConfig) ([]byte, error) {
	context := SequenceContextForWalletConfig(config)
	factory := context.FactoryAddress
	if factory == (common.Address{}) {
		return nil, fmt.Errorf("unknown factory address for wallet config type %T", config)
	}

	mainModule := context.MainModuleAddress
	imageHash := config.ImageHash().Hash
	deploy, err := contracts.V2.WalletFactory.ABI.Pack("deploy", mainModule, imageHash)
	if err != nil {
		return nil, fmt.Errorf("unable to encode deploy call: %w", err)
	}

	signature, err = ethcoder.AbiCoder([]string{"address", "bytes", "bytes"}, []interface{}{factory, deploy, signature})
	if err != nil {
		return nil, fmt.Errorf("unable to encode eip-6492 signature: %w", err)
	}
	signature = append(signature, eip6492.EIP6492MagicBytes...)
	return signature, nil
}

func EIP6492SignatureWithMultipleDeployments(signature []byte, configs []core.WalletConfig) ([]byte, error) {
	if len(configs) == 1 {
		return EIP6492Signature(signature, configs[0])
	}

	var txns Transactions
	for _, config := range configs {
		context := SequenceContextForWalletConfig(config)
		factory := context.FactoryAddress
		if factory == (common.Address{}) {
			return nil, fmt.Errorf("unknown factory address for wallet config type %T", config)
		}

		mainModule := context.MainModuleAddress
		imageHash := config.ImageHash().Hash
		deploy, err := contracts.V2.WalletFactory.ABI.Pack("deploy", mainModule, imageHash)
		if err != nil {
			return nil, fmt.Errorf("unable to encode deploy call: %w", err)
		}

		txns = append(txns, &Transaction{
			RevertOnError: true,
			To:            factory,
			Data:          deploy,
		})
	}

	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	execdata, err := contracts.V2.WalletMainModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
	if err != nil {
		return nil, err
	}

	signature, err = ethcoder.AbiCoder([]string{"address", "bytes", "bytes"}, []interface{}{sequenceContextV2.GuestModuleAddress, execdata, signature})
	if err != nil {
		return nil, fmt.Errorf("unable to encode eip-6492 signature: %w", err)
	}
	signature = append(signature, eip6492.EIP6492MagicBytes...)
	return signature, nil
}

func UnwrapEIP6492Signature(signature []byte) ([]byte, error) {
	if !bytes.HasSuffix(signature, eip6492.EIP6492MagicBytes) {
		return signature, nil
	}
	signature = bytes.TrimSuffix(signature, eip6492.EIP6492MagicBytes)

	var (
		to         common.Address
		data       []byte
		signature_ []byte
	)

	err := ethcoder.AbiDecoder([]string{"address", "bytes", "bytes"}, signature, []any{&to, &data, &signature_})
	if err != nil {
		return nil, fmt.Errorf("unable to decode eip-6492 signature: %w", err)
	}

	return signature_, nil
}
