package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type RpcRelayer struct {
	provider        *ethrpc.Provider
	receiptListener *ethreceipts.ReceiptsListener
	Service         proto.Relayer
}

var _ sequence.Relayer = &RpcRelayer{}

// NewRpcRelayer creates a new Sequence Relayer client instance. See https://docs.sequence.xyz for a list of
// relayer urls, and please see https://sequence.build to get a `projectAccessKey`.
func NewRpcRelayer(relayerURL string, projectAccessKey string, provider *ethrpc.Provider, receiptListener *ethreceipts.ReceiptsListener, options ...Options) (*RpcRelayer, error) {
	// TODO: move receiptListener to Options, and if unspecified, use the Sequence Indexer
	// for the receipts listener instead. Or, we can also have the receipts filter method
	// on the relayer service too.

	_, err := url.Parse(relayerURL)
	if err != nil {
		return nil, fmt.Errorf("relayerURL is invalid: %w", err)
	}

	service := NewRelayer(relayerURL, projectAccessKey, options...)

	return &RpcRelayer{
		provider:        provider,
		receiptListener: receiptListener,
		Service:         service,
	}, nil
}

func (r *RpcRelayer) GetProvider() *ethrpc.Provider {
	return r.provider
}

func (r *RpcRelayer) EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, transactions sequence.Transactions) (sequence.Transactions, error) {
	chainID, err := r.provider.ChainID(ctx)
	if err != nil {
		return sequence.Transactions{}, fmt.Errorf("unable to get chain id: %w", err)
	}

	to, err := sequence.AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return sequence.Transactions{}, fmt.Errorf("unable to derive address: %w", err)
	}

	data, err := sequence.SignedTransactions{CallsPayload: transactions.Payload(to, chainID)}.Data()
	if err != nil {
		return sequence.Transactions{}, fmt.Errorf("unable to encode calldata: %w", err)
	}

	results, err := r.Service.Simulate(ctx, to.String(), hexutil.Encode(data))
	if err != nil {
		return sequence.Transactions{}, fmt.Errorf("unable to simulate: %w", err)
	}
	if len(results) < len(transactions.Calls) {
		return sequence.Transactions{}, fmt.Errorf("%v simulation results, need %v", len(results), len(transactions.Calls))
	}

	transactions_ := sequence.Transactions{
		Calls: make([]v3.Call, 0, len(transactions.Calls)),
		Space: new(big.Int).Set(transactions.Space),
		Nonce: new(big.Int).Set(transactions.Nonce),
	}
	for i, call := range transactions.Calls {
		call.GasLimit = new(big.Int).SetUint64(uint64(results[i].GasLimit))
		transactions_.Calls = append(transactions_.Calls, call)
	}

	return transactions_, nil
}

// NOTE: nonce space is 160 bits wide
func (r *RpcRelayer) Nonce(ctx context.Context, wallet common.Address, space *big.Int) (*big.Int, error) {
	var encodedSpace *string
	if space != nil {
		encodedSpace = new(string)
		*encodedSpace = fmt.Sprintf("0x%x", space)
	}

	encodedNonce, err := r.Service.GetMetaTxnNonce(ctx, wallet.Hex(), encodedSpace)
	if err != nil {
		return nil, err
	}

	var nonce big.Int
	if _, ok := nonce.SetString(encodedNonce, 0); !ok {
		return nil, fmt.Errorf("%v is not a big.Int", encodedNonce)
	}

	return &nonce, nil
}

func (r *RpcRelayer) Simulate(ctx context.Context, txs *sequence.SignedTransactions) ([]*sequence.RelayerSimulateResult, error) {
	to := txs.Address()
	data, err := txs.Data()
	if err != nil {
		return nil, err
	}

	res, err := r.Service.Simulate(ctx, to.String(), hexutil.Encode(data))
	if err != nil {
		return nil, err
	}

	var results []*sequence.RelayerSimulateResult
	for _, r := range res {
		results = append(results, &sequence.RelayerSimulateResult{
			Executed:  r.Executed,
			Succeeded: r.Succeeded,
			Result:    r.Result,
			Reason:    r.Reason,
			GasUsed:   r.GasUsed,
			GasLimit:  r.GasLimit,
		})
	}
	return results, nil
}

// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
func (r *RpcRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions, deployConfig core.ImageHashable, quote ...*sequence.RelayerFeeQuote) (*types.Transaction, ethtxn.WaitReceipt, error) {
	to := signedTxs.Address()
	data, err := signedTxs.Data()
	if err != nil {
		return nil, nil, err
	}

	call := &proto.MetaTxn{
		Contract:      to.Hex(),
		Input:         hexutil.Encode(data),
		WalletAddress: to.Hex(),
	}

	var txQuote *string
	if len(quote) > 0 {
		txQuote = (*string)(quote[0])
	}

	// TODO: check contents of Contract and input, if empty, lets not even bother asking the server..

	ok, metaTxnID, err := r.Service.SendMetaTxn(ctx, call, txQuote, nil)
	if err != nil {
		return nil, nil, err
	}
	if !ok {
		return nil, nil, fmt.Errorf("failed to relay meta transaction: unknown reason")
	}
	if metaTxnID == "" {
		return nil, nil, fmt.Errorf("failed to relay meta transaction: server returned empty metaTxnID")
	}

	waitReceipt := func(ctx context.Context) (*types.Receipt, error) {
		// NOTE: to timeout the request, pass a ctx from context.WithTimeout
		_, receipt, err := r.Wait(ctx, signedTxs)
		return receipt, err
	}

	// TODO: 2nd argument will be nil, we may even want to remove it from here...
	return nil, waitReceipt, nil
}

func (r *RpcRelayer) FeeOptions(ctx context.Context, signedTxs *sequence.SignedTransactions) ([]*sequence.RelayerFeeOption, *sequence.RelayerFeeQuote, error) {
	data, err := signedTxs.Data()
	if err != nil {
		return nil, nil, err
	}

	options, _, quote, err := r.Service.FeeOptions(
		ctx,
		signedTxs.Address().String(),
		signedTxs.Address().String(),
		hexutil.Encode(data),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	var retQuote *sequence.RelayerFeeQuote
	if quote != nil {
		retQuote = ethkit.ToPtr(sequence.RelayerFeeQuote(*quote))
	}

	return convFeeOptionsToRelayerFeeOptions(options), retQuote, nil
}

// ....
func (r *RpcRelayer) Wait(ctx context.Context, payload v3.PayloadDigestable) (sequence.MetaTxnStatus, *types.Receipt, error) {
	// Fetch the meta transaction receipt from the relayer service
	if r.receiptListener == nil {
		return r.waitMetaTxnReceipt(ctx, payload)
	}

	// Fetch the meta transaction receipt from the receipt listener
	receipt, status, _, _, err := sequence.FetchReceipt(ctx, r.receiptListener, payload)
	if err != nil {
		return 0, nil, err
	}
	return status, receipt.Receipt(), nil
}

func (r *RpcRelayer) waitMetaTxnReceipt(ctx context.Context, payload v3.PayloadDigestable) (sequence.MetaTxnStatus, *types.Receipt, error) {
	// TODO: in future GetMetaTxnReceipt() will be renamed to WaitTransactionReceipt()
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				return 0, nil, err
			}
			return 0, nil, nil
		default:
		}

		metaTxnReceipt, err := r.Service.GetMetaTxnReceipt(ctx, payload.Digest().String())
		if metaTxnReceipt == nil && err == nil {
			// currently we assume that if the receipt is nil, and error is nil, then
			// we're still searching for the transaction. This is a hack, and we should
			// use proper error codes when we upgrade webrpc version
			continue
		}
		if err != nil {
			return sequence.MetaTxnStatusUnknown, nil, err
		}
		txnReceipt := metaTxnReceipt.TxnReceipt
		var receipt *types.Receipt
		err = json.Unmarshal([]byte(txnReceipt), &receipt)
		if err != nil {
			return 0, nil, fmt.Errorf("failed to decode txn receipt data: %w", err)
		}
		return MetaTxnStatusFromString(metaTxnReceipt.Status), receipt, nil
	}
}

func (r *RpcRelayer) Client() proto.Relayer {
	return r.Service
}

func convFeeTokenToRelayerFeeToken(token *proto.FeeToken) sequence.RelayerFeeToken {
	var contractAddress *common.Address
	if token.ContractAddress != nil {
		contractAddress = ethkit.ToPtr(common.HexToAddress(*token.ContractAddress))
	}

	return sequence.RelayerFeeToken{
		ChainID:         big.NewInt(0).SetUint64(token.ChainId),
		Name:            token.Name,
		Symbol:          token.Symbol,
		Type:            sequence.RelayerFeeTokenType(token.Type),
		Decimals:        token.Decimals,
		LogoURL:         token.LogoURL,
		ContractAddress: contractAddress,
	}
}

func convFeeOptionToRelayerFeeOption(option *proto.FeeOption) *sequence.RelayerFeeOption {
	value, _ := big.NewInt(0).SetString(option.Value, 10)
	return &sequence.RelayerFeeOption{
		Token:    convFeeTokenToRelayerFeeToken(option.Token),
		To:       common.HexToAddress(option.To),
		Value:    value,
		GasLimit: big.NewInt(0).SetUint64(uint64(option.GasLimit)),
	}
}

func convFeeOptionsToRelayerFeeOptions(options []*proto.FeeOption) []*sequence.RelayerFeeOption {
	var feeOptions []*sequence.RelayerFeeOption
	for _, option := range options {
		feeOptions = append(feeOptions, convFeeOptionToRelayerFeeOption(option))
	}
	return feeOptions
}
