package relayer

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/core/v1v2"
	"github.com/0xsequence/go-sequence/core/v1v2/relayer/proto"
	v1 "github.com/0xsequence/go-sequence/core/v1v2/v1"
	v2 "github.com/0xsequence/go-sequence/core/v1v2/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// LocalRelayer is a simple implementation of a relayer which will dispatch
// meta transactions locally. This should only be used for testing / debugging,
// and never seriously in a real app.
type LocalRelayer struct {
	Sender          *ethwallet.Wallet
	receiptListener *ethreceipts.ReceiptsListener
}

var _ v1v2.Relayer = &LocalRelayer{}

func NewLocalRelayer(sender *ethwallet.Wallet, receiptListener *ethreceipts.ReceiptsListener) (*LocalRelayer, error) {
	if sender.GetProvider() == nil {
		return nil, v1v2.ErrProviderNotSet
	}
	return &LocalRelayer{
		Sender:          sender,
		receiptListener: receiptListener,
	}, nil
}

func (r *LocalRelayer) GetProvider() *ethrpc.Provider {
	if r.Sender == nil || r.Sender.GetProvider() == nil {
		return nil
	}
	return r.Sender.GetProvider()
}

func (r *LocalRelayer) EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext v1v2.WalletContext, txns v1v2.Transactions) (v1v2.Transactions, error) {
	walletAddress, err := v1v2.AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return nil, err
	}

	provider := r.GetProvider()

	isWalletDeployed, err := v1v2.IsWalletDeployed(provider, walletAddress)
	if err != nil {
		return nil, err
	}

	defaultGasLimit := int64(800_000)

	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	for i := range encodedTxns {
		txn := &encodedTxns[i]

		// Respect gasLimit request of the transaction (as long as its not 0)
		if txn.GasLimit != nil && txn.GasLimit.Cmp(big.NewInt(0)) > 0 {
			continue
		}

		// Fee can't be estimated locally for delegateCalls
		if txn.DelegateCall {
			txn.GasLimit = big.NewInt(int64(defaultGasLimit))
			continue
		}

		// Fee can't be estimated for self-called if wallet hasn't been deployed
		if txn.To == walletAddress && !isWalletDeployed {
			txn.GasLimit = big.NewInt(int64(defaultGasLimit))
			continue
		}

		// Estimate with eth_estimate call
		callMsg := ethereum.CallMsg{
			From:  walletAddress,
			Gas:   0, // estimating this value
			Value: txn.Value,
			Data:  txn.Data,
		}
		zeroAddress := common.Address{}
		if txn.To != zeroAddress {
			callMsg.To = &txn.To
		}

		gasLimit, err := provider.EstimateGas(ctx, callMsg)
		if err != nil {
			txn.GasLimit = big.NewInt(int64(defaultGasLimit))
			continue
		}
		txn.GasLimit = big.NewInt(0).SetUint64(gasLimit)
	}

	// update gasLimit on original transactions
	for i, txn := range txns {
		txn.GasLimit = encodedTxns[i].GasLimit
	}

	return txns, nil
}

func (r *LocalRelayer) GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext v1v2.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	return v1v2.GetWalletNonce(r.GetProvider(), walletConfig, walletContext, space, blockNum)
}

func (r *LocalRelayer) Simulate(ctx context.Context, txs *v1v2.SignedTransactions) ([]*v1v2.RelayerSimulateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *LocalRelayer) Relay(ctx context.Context, signedTxs *v1v2.SignedTransactions, quote ...*v1v2.RelayerFeeQuote) (v1v2.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	// NOTE: this implementation assumes the wallet is deployed and does not do automatic bundle creation (aka prepending / bundling
	// a wallet creation call)

	// TODO: lets update LocalRelayer so it'll do auto-bundle creation.. to prepend, and send to guestModule, etc..
	// its more consistent, and easier for tests..

	var version int
	switch signedTxs.WalletConfig.(type) {
	case *v1.WalletConfig:
		version = 1
	case *v2.WalletConfig:
		version = 2
	case *v3.WalletConfig:
		version = 3
	default:
		return "", nil, nil, fmt.Errorf(`unknown version %T`, signedTxs.WalletConfig)
	}

	sender := r.Sender

	var to common.Address
	var execdata []byte
	var err error
	switch version {
	case 1, 2:
		to, execdata, err = v1v2.EncodeTransactionsForRelaying(
			r,
			signedTxs.WalletAddress,
			signedTxs.WalletConfig,
			signedTxs.WalletContext,
			signedTxs.Transactions,
			signedTxs.Nonce,
			signedTxs.Signature,
		)
	}
	if err != nil {
		return "", nil, nil, err
	}

	if r.IsDeployTransaction(signedTxs) {
		to = signedTxs.WalletContext.GuestModuleAddress
	} else {
		isDeployed, err := v1v2.IsWalletDeployed(r.GetProvider(), to)
		if err != nil {
			return "", nil, nil, err
		}

		if !isDeployed {
			_, factoryAddress, deployData, err := v1v2.EncodeWalletDeployment(signedTxs.WalletConfig, signedTxs.WalletContext)
			if err != nil {
				return "", nil, nil, err
			}

			txns := v1v2.Transactions{
				{
					RevertOnError: true,
					To:            factoryAddress,
					Data:          deployData,
				},
				{
					RevertOnError: true,
					To:            to,
					Data:          execdata,
				},
			}

			switch version {
			case 1, 2:
				encodedTxns, err := txns.EncodedTransactions()
				if err != nil {
					return "", nil, nil, err
				}

				execdata, err = contracts.V1.WalletMainModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
				if err != nil {
					return "", nil, nil, err
				}

			}

			to = signedTxs.WalletContext.GuestModuleAddress
		}
	}

	walletAddress, err := v1v2.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
	if err != nil {
		return "", nil, nil, err
	}

	var metaTxnID v1v2.MetaTxnID
	switch version {
	case 1, 2:
		metaTxnID, _, err = v1v2.ComputeMetaTxnID(signedTxs.ChainID, walletAddress, signedTxs.Transactions, signedTxs.Nonce, v1v2.MetaTxnWalletExec)
		if err != nil {
			return "", nil, nil, err
		}
	case 3:
		metaTxnID = v1v2.MetaTxnID(hex.EncodeToString(signedTxs.Digest.Bytes()))
	}

	ntx, err := sender.NewTransaction(ctx, &ethtxn.TransactionRequest{
		To: &to, Data: execdata,
	})
	if err != nil {
		return metaTxnID, nil, nil, err
	}

	signedTx, err := sender.SignTx(ntx, signedTxs.ChainID)
	if err != nil {
		return metaTxnID, nil, nil, err
	}

	ntx, waitReceipt, err := sender.SendTransaction(ctx, signedTx)
	if err != nil {
		return metaTxnID, nil, nil, err
	}

	return metaTxnID, ntx, waitReceipt, nil
}

func (r *LocalRelayer) Wait(ctx context.Context, metaTxnID v1v2.MetaTxnID, optTimeout ...time.Duration) (v1v2.MetaTxnStatus, *types.Receipt, error) {
	if r.receiptListener == nil {
		return 0, nil, fmt.Errorf("relayer: failed to wait for metaTxnID as receiptListener is not set")
	}
	result, receipt, _, err := v1v2.FetchMetaTransactionReceipt(ctx, r.receiptListener, metaTxnID, optTimeout...)
	if err != nil {
		return 0, nil, err
	}
	var status v1v2.MetaTxnStatus
	if result != nil {
		status = result.Status
	}
	return status, receipt.Receipt(), nil
}

func (r *LocalRelayer) FeeOptions(ctx context.Context, signedTxs *v1v2.SignedTransactions) ([]*v1v2.RelayerFeeOption, *v1v2.RelayerFeeQuote, error) {
	return nil, nil, nil
}

func (r *LocalRelayer) IsDeployTransaction(signedTxs *v1v2.SignedTransactions) bool {
	for _, txn := range signedTxs.Transactions {
		if txn.To == signedTxs.WalletContext.FactoryAddress {
			return true
		}
	}
	return false
}

func (r *LocalRelayer) Client() proto.Relayer {
	// no relayer used
	return nil
}
