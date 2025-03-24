package relayer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/relayer/proto"
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
		return sequence.Transactions{}, err
	}

	provider := r.GetProvider()

	isWalletDeployed, err := sequence.IsWalletDeployed(provider, walletAddress)
	if err != nil {
		return sequence.Transactions{}, err
	}

	defaultGasLimit := int64(800_000)

	for i := range txns.Calls {
		txn := &txns.Calls[i]

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

	return txns, nil
}

func (r *LocalRelayer) Nonce(ctx context.Context, wallet common.Address, space *big.Int) (*big.Int, error) {
	return sequence.Nonce(ctx, wallet, space, r.GetProvider())
}

func (r *LocalRelayer) Simulate(ctx context.Context, txs *sequence.SignedTransactions) ([]*sequence.RelayerSimulateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *LocalRelayer) Relay(ctx context.Context, transactions *sequence.SignedTransactions, deployConfig core.ImageHashable, quote ...*sequence.RelayerFeeQuote) (*types.Transaction, ethtxn.WaitReceipt, error) {
	// NOTE: this implementation assumes the wallet is deployed and does not do automatic bundle creation (aka prepending / bundling
	// a wallet creation call)

	// TODO: lets update LocalRelayer so it'll do auto-bundle creation.. to prepend, and send to guestModule, etc..
	// its more consistent, and easier for tests..

	sender := r.Sender
	chainID, err := sender.GetProvider().ChainID(ctx)
	if err != nil {
		return nil, nil, err
	}
	context := sequence.V3SequenceContext()

	to := transactions.Address()
	data, err := transactions.Data()
	if err != nil {
		return nil, nil, err
	}

	isDeployed, err := sequence.IsWalletDeployed(r.GetProvider(), to)
	if err != nil {
		return nil, nil, err
	}

	if !isDeployed {
		_, factoryAddress, deployData, err := sequence.EncodeWalletDeployment(deployConfig, context)
		if err != nil {
			return nil, nil, err
		}

		txns := []v3.Call{
			{
				To:              factoryAddress,
				Data:            deployData,
				BehaviorOnError: v3.BehaviorOnErrorRevert,
			},
			{
				To:              to,
				Data:            data,
				BehaviorOnError: v3.BehaviorOnErrorRevert,
			},
		}

		to = context.GuestModuleAddress
		data, err = sequence.SignedTransactions{CallsPayload: v3.ConstructCallsPayload(context.GuestModuleAddress, chainID, txns, nil, nil)}.Data()
	}

	ntx, err := sender.NewTransaction(ctx, &ethtxn.TransactionRequest{
		To: &to, Data: data,
	})
	if err != nil {
		return nil, nil, err
	}

	signedTx, err := sender.SignTx(ntx, chainID)
	if err != nil {
		return nil, nil, err
	}

	ntx, waitReceipt, err := sender.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, nil, err
	}

	return ntx, waitReceipt, nil
}

func (r *LocalRelayer) Wait(ctx context.Context, payload v3.PayloadDigestable) (sequence.MetaTxnStatus, *types.Receipt, error) {
	if r.receiptListener == nil {
		return 0, nil, fmt.Errorf("relayer: failed to wait for metaTxnID as receiptListener is not set")
	}
	receipt, status, _, _, err := sequence.FetchReceipt(ctx, r.receiptListener, payload)
	if err != nil {
		return 0, nil, err
	}
	return status, receipt.Receipt(), nil
}

func (r *LocalRelayer) FeeOptions(ctx context.Context, signedTxs *sequence.SignedTransactions) ([]*sequence.RelayerFeeOption, *sequence.RelayerFeeQuote, error) {
	return nil, nil, nil
}

func (r *LocalRelayer) Client() proto.Relayer {
	// no relayer used
	return nil
}
