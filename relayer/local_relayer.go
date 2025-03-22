package relayer

import (
	"context"
	"encoding/hex"
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
	"github.com/0xsequence/go-sequence/contracts"
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

func (r *LocalRelayer) GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	if space == nil {
		space = common.Big0
	}

	data, err := contracts.V3.WalletStage1Module.Encode("readNonce", space)
	if err != nil {
		return nil, fmt.Errorf("unable to read nonce: %w", err)
	}

	var nonce *big.Int
	err = contracts.V3.WalletStage1Module.Decode(&nonce, "readNonce", data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode nonce: %w", err)
	}

	return nonce, nil
}

func (r *LocalRelayer) Simulate(ctx context.Context, txs *sequence.SignedTransactions) ([]*sequence.RelayerSimulateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *LocalRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions, quote ...*sequence.RelayerFeeQuote) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	// NOTE: this implementation assumes the wallet is deployed and does not do automatic bundle creation (aka prepending / bundling
	// a wallet creation call)

	// TODO: lets update LocalRelayer so it'll do auto-bundle creation.. to prepend, and send to guestModule, etc..
	// its more consistent, and easier for tests..

	sender := r.Sender

	to := signedTxs.Address()
	data, err := signedTxs.Data()
	if err != nil {
		return "", nil, nil, err
	}

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
				Data:          data,
			},
		}

		switch version {
		case 1, 2:
			encodedTxns, err := txns.EncodedTransactions()
			if err != nil {
				return "", nil, nil, err
			}

			data, err = contracts.V1.WalletMainModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
			if err != nil {
				return "", nil, nil, err
			}

		case 3:
			_, data, err = sequence.EncodeTransactionsForRelayingV3(
				r,
				signedTxs.WalletAddress,
				signedTxs.ChainID,
				signedTxs.WalletConfig,
				signedTxs.WalletContext,
				txns,
				signedTxs.Space,
				signedTxs.Nonce,
				signedTxs.Signature,
			)
			if err != nil {
				return "", nil, nil, err
			}

			data, err = contracts.V3.WalletStage1Module.Encode("execute", data, []byte{})
			if err != nil {
				return "", nil, nil, err
			}
		}

		to = signedTxs.WalletContext.GuestModuleAddress
	}

	walletAddress, err := sequence.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
	if err != nil {
		return "", nil, nil, err
	}

	var metaTxnID sequence.MetaTxnID
	switch version {
	case 1, 2:
		metaTxnID, _, err = sequence.ComputeMetaTxnID(signedTxs.ChainID, walletAddress, signedTxs.Transactions, signedTxs.Nonce, sequence.MetaTxnWalletExec)
		if err != nil {
			return "", nil, nil, err
		}
	case 3:
		metaTxnID = sequence.MetaTxnID(hex.EncodeToString(signedTxs.Digest.Bytes()))
	}

	ntx, err := sender.NewTransaction(ctx, &ethtxn.TransactionRequest{
		To: &to, Data: data,
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
