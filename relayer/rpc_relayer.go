package relayer

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type RpcRelayer struct {
	provider *ethrpc.Provider
	Service  proto.Relayer
}

var _ sequence.Relayer = &RpcRelayer{}

func NewRpcRelayer(provider *ethrpc.Provider, rpcRelayerURL string, httpClient proto.HTTPClient) (*RpcRelayer, error) {
	_, err := url.Parse(rpcRelayerURL)
	if err != nil {
		return nil, fmt.Errorf("rpcRelayerURL is invalid: %w", err)
	}

	service := proto.NewRelayerClient(rpcRelayerURL, httpClient)

	return &RpcRelayer{
		provider: provider,
		Service:  service,
	}, nil
}

func (r *RpcRelayer) GetProvider() *ethrpc.Provider {
	return r.provider
}

// ..
func (r *RpcRelayer) EstimateGasLimits(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, txns sequence.Transactions) (sequence.Transactions, error) {
	// NOTE: this is a stub
	// TODO: replace this impl with the new impl or call to the server, etc.

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

	for _, txn := range txns {
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
			return nil, fmt.Errorf("ethtxn: %w", err)
		}
		txn.GasLimit = big.NewInt(0).SetUint64(gasLimit)
	}

	return txns, nil
}

// NOTE: nonce space is 160 bits wide
func (r *RpcRelayer) GetNonce(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	if blockNum != nil {
		return sequence.GetWalletNonce(r.GetProvider(), walletConfig, walletContext, space, blockNum)
	}

	walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return nil, err
	}

	var encodedSpace *string
	if space != nil {
		encodedSpace = new(string)
		*encodedSpace = fmt.Sprintf("0x%x", space)
	}

	encodedNonce, err := r.Service.GetMetaTxnNonce(ctx, walletAddress.Hex(), encodedSpace)
	if err != nil {
		return nil, err
	}

	var nonce big.Int
	if _, ok := nonce.SetString(encodedNonce, 0); !ok {
		return nil, fmt.Errorf("%v is not a big.Int", encodedNonce)
	}

	return &nonce, nil
}

// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
func (r *RpcRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	walletAddress, err := sequence.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
	if err != nil {
		return "", nil, nil, err
	}

	to, execdata, err := sequence.EncodeTransactionsForRelaying(
		r,
		signedTxs.WalletConfig,
		signedTxs.WalletContext,
		signedTxs.Transactions,
		signedTxs.Nonce,
		signedTxs.Signature,
	)
	if err != nil {
		return "", nil, nil, err
	}

	call := &proto.MetaTxn{
		Contract:      to.Hex(),
		Input:         hexutil.Encode(execdata),
		WalletAddress: walletAddress.Hex(),
	}

	// TODO: check contents of Contract and input, if empty, lets not even bother asking the server..

	ok, metaTxnID, err := r.Service.SendMetaTxn(ctx, call)
	if err != nil {
		return sequence.MetaTxnID(metaTxnID), nil, nil, err
	}
	if !ok {
		return sequence.MetaTxnID(metaTxnID), nil, nil, proto.Failf("failed to relay meta transaction: unknown reason")
	}
	if metaTxnID == "" {
		return "", nil, nil, proto.Failf("failed to relay meta transaction: server returned empty metaTxnID")
	}

	waitReceipt := func(ctx context.Context) (*types.Receipt, error) {
		// NOTE: to timeout the request, pass a ctx from context.WithTimeout
		_, receipt, err := r.Wait(ctx, sequence.MetaTxnID(metaTxnID), nil)
		return receipt, err
	}

	// TODO: 2nd argument will be nil, we may even want to remove it from here...
	return sequence.MetaTxnID(metaTxnID), nil, waitReceipt, nil
}

// ..
func (r *RpcRelayer) Wait(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout *time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	return sequence.WaitForMetaTxn(ctx, r.GetProvider(), metaTxnID, optTimeout)
}
