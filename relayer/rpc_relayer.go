package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
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

func (r *RpcRelayer) EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, txns sequence.Transactions) (sequence.Transactions, error) {
	walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return nil, err
	}

	requestData, err := txns.EncodeRaw()
	if err != nil {
		return nil, err
	}

	response, err := r.Service.UpdateMetaTxnGasLimits(ctx, walletAddress.Hex(), walletConfig, hexutil.Encode(requestData))
	if err != nil {
		return nil, err
	}

	responseData, err := hexutil.Decode(response)
	if err != nil {
		return nil, err
	}

	return sequence.DecodeRawTransactions(responseData)
}

// NOTE: nonce space is 160 bits wide
func (r *RpcRelayer) GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
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

func (r *RpcRelayer) Simulate(ctx context.Context, txs *sequence.SignedTransactions) ([]*sequence.RelayerSimulateResult, error) {
	to, execdata, err := sequence.EncodeTransactionsForRelaying(
		r,
		txs.WalletAddress,
		txs.WalletConfig,
		txs.WalletContext,
		txs.Transactions,
		txs.Nonce,
		txs.Signature,
	)

	if err != nil {
		return nil, err
	}

	res, err := r.Service.Simulate(ctx, to.String(), "0x"+common.Bytes2Hex(execdata))
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
func (r *RpcRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions, quote ...*sequence.RelayerFeeQuote) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	walletAddress := signedTxs.WalletAddress
	var err error

	if walletAddress == (common.Address{}) {
		walletAddress, err = sequence.AddressFromWalletConfig(signedTxs.WalletConfig, signedTxs.WalletContext)
		if err != nil {
			return "", nil, nil, err
		}
	}

	var to common.Address
	var execdata []byte
	switch signedTxs.WalletConfig.(type) {
	case *v1.WalletConfig, *v2.WalletConfig:
		to, execdata, err = sequence.EncodeTransactionsForRelaying(
			r,
			signedTxs.WalletAddress,
			signedTxs.WalletConfig,
			signedTxs.WalletContext,
			signedTxs.Transactions,
			signedTxs.Nonce,
			signedTxs.Signature,
		)
	case *v3.WalletConfig:
		to, execdata, err = sequence.EncodeTransactionsForRelayingV3(
			r,
			signedTxs.WalletAddress,
			signedTxs.ChainID,
			signedTxs.WalletConfig,
			signedTxs.WalletContext,
			signedTxs.Transactions,
			signedTxs.Space,
			signedTxs.Nonce,
			signedTxs.Signature,
		)
	}
	if err != nil {
		return "", nil, nil, err
	}

	if r.IsDeployTransaction(signedTxs) {
		to = signedTxs.WalletContext.GuestModuleAddress
	}

	call := &proto.MetaTxn{
		Contract:      to.Hex(),
		Input:         hexutil.Encode(execdata),
		WalletAddress: walletAddress.Hex(),
	}

	var txQuote *string
	if len(quote) > 0 {
		txQuote = (*string)(quote[0])
	}

	// TODO: check contents of Contract and input, if empty, lets not even bother asking the server..

	ok, metaTxnID, err := r.Service.SendMetaTxn(ctx, call, txQuote, nil)
	if err != nil {
		return sequence.MetaTxnID(metaTxnID), nil, nil, err
	}
	if !ok {
		return sequence.MetaTxnID(metaTxnID), nil, nil, fmt.Errorf("failed to relay meta transaction: unknown reason")
	}
	if metaTxnID == "" {
		return "", nil, nil, fmt.Errorf("failed to relay meta transaction: server returned empty metaTxnID")
	}

	waitReceipt := func(ctx context.Context) (*types.Receipt, error) {
		// NOTE: to timeout the request, pass a ctx from context.WithTimeout
		_, receipt, err := r.Wait(ctx, sequence.MetaTxnID(metaTxnID))
		return receipt, err
	}

	// TODO: 2nd argument will be nil, we may even want to remove it from here...
	return sequence.MetaTxnID(metaTxnID), nil, waitReceipt, nil
}

func (r *RpcRelayer) FeeOptions(ctx context.Context, signedTxs *sequence.SignedTransactions) ([]*sequence.RelayerFeeOption, *sequence.RelayerFeeQuote, error) {
	data, err := signedTxs.Execdata()
	if err != nil {
		return nil, nil, err
	}

	options, _, quote, err := r.Service.FeeOptions(
		ctx,
		signedTxs.WalletAddress.String(),
		signedTxs.WalletAddress.String(),
		"0x"+common.Bytes2Hex(data),
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
func (r *RpcRelayer) Wait(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	// Fetch the meta transaction receipt from the relayer service
	if r.receiptListener == nil {
		return r.waitMetaTxnReceipt(ctx, metaTxnID, optTimeout...)
	}

	// Fetch the meta transaction receipt from the receipt listener
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

func (r *RpcRelayer) waitMetaTxnReceipt(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	// TODO: in future GetMetaTxnReceipt() will be renamed to WaitTransactionReceipt()

	var clear context.CancelFunc
	if len(optTimeout) > 0 {
		ctx, clear = context.WithTimeout(ctx, optTimeout[0])
		defer clear()
	}

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

		metaTxnReceipt, err := r.Service.GetMetaTxnReceipt(ctx, metaTxnID.String())
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

func (r *RpcRelayer) IsDeployTransaction(signedTxs *sequence.SignedTransactions) bool {
	for _, txn := range signedTxs.Transactions {
		if txn.To == signedTxs.WalletContext.FactoryAddress {
			return true
		}
	}
	return false
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
