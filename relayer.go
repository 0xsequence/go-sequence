package sequence

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/davecgh/go-spew/spew"
)

type MetaTxnID string // TODO: can this be a common.Hash ..? I think not.. I think has to be just a string..

// TODO: RelayerOptions struct {
// 	 BundleCreation bool
//   CreationGasLimit *big.Int
//   Provider *ethrpc.Provider
//   URL string
// }

type Relayer interface {
	GetProvider() *ethrpc.Provider

	EstimateGasLimits(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, transactions Transactions) (Transactions, error)

	// TODO: how many bits for the space..?
	GetNonce(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error)

	// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
	// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
	// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
	//
	// TODO: ethwallet.WaitReceipt needs to move to ethtxn or ethrpc ..
	Relay(ctx context.Context, signedTxs *SignedTransactions) (MetaTxnID, *types.Transaction, ethwallet.WaitReceipt, error)

	// TODO: what "type" should metaTxID be..? it is a hash..
	// but, its kinda different cuz we dont require 0x so may not be able to use common.Hash
	// alternatively we could make metaTxID a different length of hash to differentiate it, but its a breaking change on client-side
	Wait(ctx context.Context, metaTxID MetaTxnID, timeout time.Duration) (*types.Receipt, error)

	// GasRefundOptions() // TODO, in future when needed..
}

// LocalRelayer is a simple implementation of a relayer which will dispatch
// meta transactions locally. This should only be used for testing / debugging,
// and never seriously in a real app.
type LocalRelayer struct {
	Sender *ethwallet.Wallet
}

var _ Relayer = &LocalRelayer{}

func NewLocalRelayer(sender *ethwallet.Wallet) (*LocalRelayer, error) {
	if sender.GetProvider() == nil {
		return nil, ErrProviderNotSet
	}
	return &LocalRelayer{Sender: sender}, nil
}

func (r *LocalRelayer) GetProvider() *ethrpc.Provider {
	if r.Sender == nil || r.Sender.GetProvider() == nil {
		return nil
	}
	return r.Sender.GetProvider()
}

func (r *LocalRelayer) EstimateGasLimits(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, transactions Transactions) (Transactions, error) {
	return nil, nil
}

func (r *LocalRelayer) GetNonce(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	return GetWalletNonce(r.GetProvider(), walletConfig, walletContext, space, blockNum)
}

func (r *LocalRelayer) Relay(ctx context.Context, signedTxs *SignedTransactions) (MetaTxnID, *types.Transaction, ethwallet.WaitReceipt, error) {
	sender := r.Sender

	to, execdata, err := encodeTransactionsForRelaying(
		r,
		signedTxs.WalletConfig,
		signedTxs.WalletContext,
		signedTxs.Signature,
		signedTxs.Transactions,
	)
	if err != nil {
		return "", nil, nil, err
	}

	ntx, _, err := sender.NewTransaction(ctx, &ethwallet.TransactionRequest{
		To: &to, Data: execdata,
	})
	if err != nil {
		return "", nil, nil, err
	}

	signedTx, err := sender.SignTx(ntx, signedTxs.ChainID)
	if err != nil {
		return "", nil, nil, err
	}

	// TODO: compute the metaTxID .. we can prob do this sooner too....
	metaTxnID := ""

	ntx, waitReceipt, err := sender.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", nil, nil, nil
	}

	return MetaTxnID(metaTxnID), ntx, waitReceipt, nil
}

func (r *LocalRelayer) Wait(ctx context.Context, metaTxID MetaTxnID, timeout time.Duration) (*types.Receipt, error) {
	return nil, nil
}

// returns `to` address (either guest or wallet) and `data` of signed-metatx-calldata, aka execdata
func encodeTransactionsForRelaying(relayer Relayer, walletConfig WalletConfig, walletContext WalletContext, metaSig []byte, txns Transactions) (common.Address, []byte, error) {
	// TODO/NOTE: first version, we assume the wallet is deployed, then we can add bundlecreation after.
	// .....

	if len(txns) == 0 {
		return common.Address{}, nil, fmt.Errorf("cannot encode empty transactions")
	}

	// Encode transaction to be sent to a deployed wallet
	walletAddress, err := AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return common.Address{}, nil, err
	}

	seqNonce, err := txns.Nonce()
	if err != nil {
		return common.Address{}, nil, err
	}

	execdata, err := contracts.WalletMainModule.ABI.Pack("execute", transactionsArrayOfValues(txns), seqNonce, metaSig)
	if err != nil {
		fmt.Println("here..")
		spew.Dump(txns)
		return common.Address{}, nil, err
	}

	return walletAddress, execdata, nil
}
