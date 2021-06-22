package sequence

import (
	"context"

	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
)

var zeroAddress = common.Address{}

func DeploySequenceWallet(sender *ethwallet.Wallet, walletConfig WalletConfig, walletContext WalletContext) (common.Address, *types.Transaction, ethtxn.WaitReceipt, error) {
	if sender.GetProvider() == nil {
		return common.Address{}, nil, nil, ErrProviderNotSet
	}

	provider := sender.GetProvider()
	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	walletImageHash, err := ImageHashOfWalletConfig(walletConfig)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	walletAddress, err := AddressFromImageHash(walletImageHash, walletContext)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	deployData, err := contracts.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, common.HexToHash(walletImageHash))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	deployTx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:   &walletContext.FactoryAddress,
		Data: deployData,
	})

	signedDeployTx, err := sender.SignTx(deployTx, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tx, waitReceipt, err := sender.SendTransaction(context.Background(), signedDeployTx)

	return walletAddress, tx, waitReceipt, nil
}
