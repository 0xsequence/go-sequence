package relayer

import (
	"context"
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
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
)

// LocalRelayer is a simple implementation of a relayer which will dispatch
// meta transactions locally. This should only be used for testing / debugging,
// and never seriously in a real app.
type LocalRelayer struct {
	Sender          *ethwallet.Wallet
	receiptListener *ethreceipts.ReceiptsListener
}

var _ sequence.Relayer = &LocalRelayer{}

func NewLocalRelayer(sender *ethwallet.Wallet, receiptListener *ethreceipts.ReceiptsListener) (*LocalRelayer, error) {
	if sender.GetProvider() == nil {
		return nil, sequence.ErrProviderNotSet
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

func (r *LocalRelayer) EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, txns sequence.Transactions) (sequence.Transactions, error) {
	walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return nil, err
	}

	provider := r.GetProvider()

	isWalletDeployed, err := sequence.IsWalletDeployed(provider, walletAddress)
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

func (r *LocalRelayer) GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	return sequence.GetWalletNonce(r.GetProvider(), walletConfig, walletContext, space, blockNum)
}

func (r *LocalRelayer) Simulate(ctx context.Context, txs *sequence.SignedTransactions) ([]*sequence.RelayerSimulateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *LocalRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	// NOTE: this implementation assumes the wallet is deployed and does not do automatic bundle creation (aka prepending / bundling
	// a wallet creation call)

	// TODO: lets update LocalRelayer so it'll do auto-bundle creation.. to prepend, and send to guestModule, etc..
	// its more consistent, and easier for tests..

	sender := r.Sender

	to, execdata, err := sequence.EncodeTransactionsForRelaying(
		r,
		signedTxs.WalletAddress,
		signedTxs.WalletConfig,
		signedTxs.WalletContext,
		signedTxs.Transactions,
		signedTxs.Nonce,
		signedTxs.Signature,
	)
	if err != nil {
		return "", nil, nil, err
	}

	if r.IsDeployTransaction(signedTxs) {
		to = signedTxs.WalletContext.GuestModuleAddress
	} else {
		isDeployed, err := sequence.IsWalletDeployed(r.GetProvider(), to)
		if err != nil {
			return "", nil, nil, err
		}

		if !isDeployed {
			_, factoryAddress, deployData, err := sequence.EncodeWalletDeployment(signedTxs.WalletConfig, signedTxs.WalletContext)
			if err != nil {
				return "", nil, nil, err
			}

			txns := sequence.Transactions{
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

			encodedTxns, err := txns.EncodedTransactions()
			if err != nil {
				return "", nil, nil, err
			}

			execdata, err = contracts.WalletMainModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
			if err != nil {
				return "", nil, nil, err
			}

			to = signedTxs.WalletContext.GuestModuleAddress
		}
	}

	walletAddress, err := sequence.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
	if err != nil {
		return "", nil, nil, err
	}

	metaTxnID, _, err := sequence.ComputeMetaTxnID(signedTxs.ChainID, walletAddress, signedTxs.Transactions, signedTxs.Nonce, sequence.MetaTxnWalletExec)
	if err != nil {
		return "", nil, nil, err
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

func (r *LocalRelayer) Wait(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	if r.receiptListener == nil {
		return 0, nil, fmt.Errorf("relayer: failed to wait for metaTxnID as receiptListener is not set")
	}
	result, receipt, _, err := sequence.FetchMetaTransactionReceipt(ctx, r.receiptListener, metaTxnID, optTimeout...)
	if err != nil {
		return 0, nil, err
	}
	var status sequence.MetaTxnStatus
	if result != nil {
		status = result.Status
	}
	return status, receipt.Receipt(), nil
}

func (r *LocalRelayer) IsDeployTransaction(signedTxs *sequence.SignedTransactions) bool {
	for _, txn := range signedTxs.Transactions {
		if txn.To == signedTxs.WalletContext.FactoryAddress {
			return true
		}
	}
	return false
}
