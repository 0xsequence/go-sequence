package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// var zeroAddress = common.Address{}

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

func EncodeWalletDeployment(walletConfig core.WalletConfig, walletContext WalletContext) (common.Address, common.Address, []byte, error) {
	walletImageHash := walletConfig.ImageHash().Hex()
	walletAddress, err := AddressFromImageHash(walletImageHash, walletContext)
	if err != nil {
		return common.Address{}, common.Address{}, nil, err
	}

	if _, ok := walletConfig.(*v1.WalletConfig); ok {
		deployData, err := contracts.V1.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, common.HexToHash(walletImageHash))
		if err != nil {
			return common.Address{}, common.Address{}, nil, err
		}

		return walletAddress, walletContext.FactoryAddress, deployData, nil
	} else if _, ok := walletConfig.(*v2.WalletConfig); ok {
		deployData, err := contracts.V2.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, common.HexToHash(walletImageHash))
		if err != nil {
			return common.Address{}, common.Address{}, nil, err
		}
		return walletAddress, walletContext.FactoryAddress, deployData, nil
	} else if _, ok := walletConfig.(*v3.WalletConfig); ok {
		deployData, err := contracts.V3.WalletFactory.ABI.Pack("deploy", walletContext.MainModuleAddress, common.HexToHash(walletImageHash))
		if err != nil {
			return common.Address{}, common.Address{}, nil, err
		}
		return walletAddress, walletContext.FactoryAddress, deployData, nil
	}
	return common.Address{}, common.Address{}, nil, fmt.Errorf("unsupported wallet config version")
}

func DecodeRevertReason(logs []*types.Log) []string {
	reasons := []string{}
	for _, log := range logs {
		_, reason, err := V1DecodeTxFailedEvent(log)
		if err != nil {
			_, reason, _, _ = V2DecodeTxFailedEvent(log)
		}
		if err != nil {
			_, reason, _, _ = V3DecodeTxFailedEvent(log)
		}

		reasons = append(reasons, reason)
	}
	return reasons
}

func ParseHexOrDec(s string) (*big.Int, bool) {
	if len(s) > 2 && s[0:2] == "0x" {
		return new(big.Int).SetString(s[2:], 16)
	}
	return new(big.Int).SetString(s, 10)
}
