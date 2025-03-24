package sequence

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
)

func DeploySequenceWallet(sender *ethwallet.Wallet, walletConfig core.WalletConfig, walletContext WalletContext) (common.Address, *types.Transaction, ethtxn.WaitReceipt, error) {
	if sender.GetProvider() == nil {
		return common.Address{}, nil, nil, ErrProviderNotSet
	}

	provider := sender.GetProvider()
	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	walletAddress, _, deployData, err := EncodeWalletDeployment(walletConfig, walletContext)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	deployTx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:   &walletContext.FactoryAddress,
		Data: deployData,
		// TODO: Move this hardcoded gas limit to a configuration
		// or fix it with a contract patch
		GasLimit: 131072,
	})
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	signedDeployTx, err := sender.SignTx(deployTx, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tx, waitReceipt, err := sender.SendTransaction(context.Background(), signedDeployTx)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return walletAddress, tx, waitReceipt, nil
}

func EncodeWalletDeployment(walletConfig core.ImageHashable, walletContext WalletContext) (common.Address, common.Address, []byte, error) {
	imageHash := walletConfig.ImageHash()
	address, err := AddressFromImageHash(imageHash, walletContext)
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}

	deployData, err := contracts.V3.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, imageHash.Hash)
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}
	return address, walletContext.FactoryAddress, deployData, nil
}

func ParseHexOrDec(s string) (*big.Int, bool) {
	if len(s) > 2 && s[0:2] == "0x" {
		return new(big.Int).SetString(s[2:], 16)
	}
	return new(big.Int).SetString(s, 10)
}
