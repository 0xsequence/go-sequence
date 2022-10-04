package relayer

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
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

func (r *RpcRelayer) Simulate(ctx context.Context, walletAddress common.Address, txns sequence.Transactions) ([]sequence.SimulateResult, error) {
	requestData, err := txns.EncodeRaw()
	if err != nil {
		return nil, err
	}

	protoResults, err := r.Service.Simulate(ctx, walletAddress.Hex(), hexutil.Encode(requestData))
	if err != nil {
		return nil, err
	}

	results := make([]sequence.SimulateResult, len(protoResults))
	for i := range results {
		results[i] = sequence.SimulateResult(*protoResults[i])
	}

	return results, nil
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

	ok, metaTxnID, err := r.Service.SendMetaTxn(ctx, call, nil)
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
		_, receipt, err := r.Wait(ctx, sequence.MetaTxnID(metaTxnID))
		return receipt, err
	}

	// TODO: 2nd argument will be nil, we may even want to remove it from here...
	return sequence.MetaTxnID(metaTxnID), nil, waitReceipt, nil
}

// ..
func (r *RpcRelayer) Wait(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	return sequence.WaitForMetaTxn(ctx, r.GetProvider(), metaTxnID, optTimeout...)
}

func (r *RpcRelayer) protoConfig(ctx context.Context, config *sequence.WalletConfig, walletAddress common.Address) (*proto.WalletConfig, error) {
	var signers []*proto.WalletSigner
	for _, signer := range config.Signers {
		signers = append(signers, &proto.WalletSigner{
			Address: signer.Address.Hex(),
			Weight:  signer.Weight,
		})
	}

	result, err := r.provider.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	chainID := result.Uint64()

	return &proto.WalletConfig{
		Address:   walletAddress.Hex(),
		Signers:   signers,
		Threshold: config.Threshold,
		ChainId:   &chainID,
	}, nil
}
