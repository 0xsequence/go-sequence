package relayer

import (
	"context"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
)

// LocalRelayer is a simple implementation of a relayer which will dispatch
// meta transactions locally. This should only be used for testing / debugging,
// and never seriously in a real app.
type LocalRelayer struct {
	Sender *ethwallet.Wallet
}

var _ sequence.Relayer = &LocalRelayer{}

func NewLocalRelayer(sender *ethwallet.Wallet) (*LocalRelayer, error) {
	if sender.GetProvider() == nil {
		return nil, sequence.ErrProviderNotSet
	}
	return &LocalRelayer{Sender: sender}, nil
}

func (r *LocalRelayer) GetProvider() *ethrpc.Provider {
	if r.Sender == nil || r.Sender.GetProvider() == nil {
		return nil
	}
	return r.Sender.GetProvider()
}

func (r *LocalRelayer) EstimateGasLimits(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, transactions sequence.Transactions) (sequence.Transactions, error) {
	// TODO: ..... lets go......
	return nil, nil
}

func (r *LocalRelayer) GetNonce(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	return sequence.GetWalletNonce(r.GetProvider(), walletConfig, walletContext, space, blockNum)
}

func (r *LocalRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	// NOTE: this implementation assumes the wallet is deployed and does not do automatic bundle creation (aka prepending / bundling
	// a wallet creation call)

	sender := r.Sender

	to, execdata, err := sequence.EncodeTransactionsForRelaying(
		r,
		signedTxs.WalletConfig,
		signedTxs.WalletContext,
		signedTxs.Signature,
		signedTxs.Transactions,
	)
	if err != nil {
		return "", nil, nil, err
	}

	walletAddress, err := sequence.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
	if err != nil {
		return "", nil, nil, err
	}

	metaTxnID, err := sequence.ComputeMetaTxnID(walletAddress, signedTxs.ChainID, signedTxs.Transactions)
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

func (r *LocalRelayer) Wait(ctx context.Context, metaTxID sequence.MetaTxnID, timeout time.Duration) (*types.Receipt, error) {
	return nil, nil
}
