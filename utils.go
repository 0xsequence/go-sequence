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

	signedDeployTx, err := sender.SignTx(deployTx, chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	tx, waitReceipt, err := sender.SendTransaction(context.Background(), signedDeployTx)

	return walletAddress, tx, waitReceipt, nil
}

func EncodeWalletDeployment(walletConfig WalletConfig, walletContext WalletContext) (common.Address, common.Address, []byte, error) {
	walletImageHash, err := ImageHashOfWalletConfig(walletConfig)
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}

	walletAddress, err := AddressFromImageHash(walletImageHash, walletContext)
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}

	deployData, err := contracts.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, common.HexToHash(walletImageHash))
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}

	return walletAddress, walletContext.FactoryAddress, deployData, nil
}

func DecodeRevertReason(logs []*types.Log) []string {
	reasons := []string{}
	for _, log := range logs {
		_, reason, _ := DecodeTxFailedEvent(log)
		reasons = append(reasons, reason)
	}
	return reasons
}
