package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"
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
	"github.com/0xsequence/go-sequence/lib/simulator"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type Options struct {
	// JWTAuthToken is an optional JWT token to authenticate with the relayer service.
	JWTAuthToken string

	// Receipts listener is an optional ethreceipts.ReceiptsListener instance to use for fetching
	// transaction receipts. If unspecified, the relayer client will poll the relayer service for
	// transaction receipts instead.
	ReceiptsListener *ethreceipts.ReceiptsListener

	// HTTPClient is an optional custom HTTP client to use for communicating with the
	// relayer service.
	HTTPClient proto.HTTPClient
}

type Client struct {
	proto.RelayerClient
	provider        *ethrpc.Provider
	receiptListener *ethreceipts.ReceiptsListener
}

var _ sequence.Relayer = &Client{}

// NewClient creates a new Sequence Relayer client instance for a specific chain.
// Please see https://sequence.build to get a `projectAccessKey`, which is your
// project's access key used to communicate with Sequence services.
//
// NOTE: the `projectAccessKey` may be optional if you're using a JWT auth token
// passed in via the `clientOptions`.
//
// The `relayerURL` is the URL of the relayer service to connect to, for example:
// https://mainnet-relayer.sequence.app for Ethereum mainnet and https://polygon-relayer.sequence.app
// for Polygon mainnet. See https://docs.sequence.xyz for a complete list of relayer urls.
//
// The `rpcProvider` is an instance of ethrpc.Provider that is used to communicate with
// the underlying blockchain network. This is required for certain operations, such as
// fetching the chain ID and wallet nonces. You may pass in `nil` if you don't need
// those methods.
//
// Finally, you may pass in optional `clientOptions` to configure the relayer client
// with jwt-based authentication, a receipts listener, and a custom HTTP client.
func NewClient(relayerURL string, projectAccessKey string, rpcProvider *ethrpc.Provider, clientOptions ...Options) (*Client, error) {
	// TODO: move receiptListener to Options, and if unspecified, use the Sequence Indexer
	// for the receipts listener instead. Or, we can also have the receipts filter method
	// on the relayer service too.

	_, err := url.Parse(relayerURL)
	if err != nil {
		return nil, fmt.Errorf("relayerURL is invalid: %w", err)
	}

	opts := Options{}
	if len(clientOptions) > 0 {
		opts = clientOptions[0]
	}

	c := &httpClient{
		client:           opts.HTTPClient,
		projectAccessKey: projectAccessKey,
	}
	if opts.HTTPClient == nil {
		c.client = http.DefaultClient
	}
	if opts.JWTAuthToken != "" {
		c.jwtAuthHeader = fmt.Sprintf("BEARER %s", opts.JWTAuthToken)
	}

	relayerClient := proto.NewRelayerClient(strings.TrimSuffix(relayerURL, "/"), c)

	return &Client{
		RelayerClient:   relayerClient,
		provider:        rpcProvider,
		receiptListener: opts.ReceiptsListener,
	}, nil
}

func (r *Client) GetProvider() *ethrpc.Provider {
	return r.provider
}

func (r *Client) GetChainID(ctx context.Context) (*big.Int, error) {
	if r.provider == nil {
		return nil, sequence.ErrProviderNotSet
	}
	return r.provider.ChainID(ctx)
}

// NOTE: nonce space is 160 bits wide
func (r *Client) GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	if blockNum != nil {
		if r.provider == nil {
			return nil, sequence.ErrProviderNotSet
		}
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

	encodedNonce, err := r.GetMetaTxnNonce(ctx, walletAddress.Hex(), encodedSpace)
	if err != nil {
		return nil, err
	}

	var nonce big.Int
	if _, ok := nonce.SetString(encodedNonce, 0); !ok {
		return nil, fmt.Errorf("%v is not a big.Int", encodedNonce)
	}

	return &nonce, nil
}

func (r *Client) Simulate(ctx context.Context, wallet common.Address, transactions sequence.Transactions) ([]*sequence.SimulateResult, error) {
	payload, err := transactions.Payload(wallet, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	results, err := r.SimulateV3(ctx, wallet.String(), hexutil.Encode(payload.Encode(wallet)))
	if err != nil {
		return nil, err
	}

	var results_ []*sequence.SimulateResult
	for _, result := range results {
		var result_ []byte
		if result.Result != nil {
			var err error
			result_, err = hexutil.Decode(*result.Result)
			if err != nil {
				return nil, fmt.Errorf("invalid /SimulateV3 result: %w", err)
			}
		}

		var err error
		if result.Error != nil {
			err = fmt.Errorf("%s", *result.Error)
		}

		results_ = append(results_, &sequence.SimulateResult{
			Result: simulator.Result{
				Status:  simulator.Status(result.Status),
				Result:  result_,
				Error:   err,
				GasUsed: result.GasUsed,
			},
			GasLimit: result.GasLimit,
		})
	}
	return results_, nil
}

// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
func (r *Client) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions, quote ...*sequence.RelayerFeeQuote) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
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
	default:
		return "", nil, nil, fmt.Errorf("unknown wallet config type: %T", signedTxs.WalletConfig)
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

	ok, metaTxnID, err := r.SendMetaTxn(ctx, call, txQuote, nil, nil)
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

func (r *Client) FeeOptions(ctx context.Context, signedTxs *sequence.SignedTransactions) ([]*sequence.RelayerFeeOption, *sequence.RelayerFeeQuote, error) {
	data, err := signedTxs.Execdata()
	if err != nil {
		return nil, nil, err
	}

	options, _, quote, err := r.RelayerClient.FeeOptions(
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
func (r *Client) Wait(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
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

func (r *Client) waitMetaTxnReceipt(ctx context.Context, metaTxnID sequence.MetaTxnID, optTimeout ...time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
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

		metaTxnReceipt, err := r.GetMetaTxnReceipt(ctx, metaTxnID.String())
		if metaTxnReceipt == nil && err == nil {
			// currently we assume that if the receipt is nil, and error is nil, then
			// we're still searching for the transaction. This is a hack, and we should
			// use proper error codes when we upgrade webrpc version
			continue
		}
		if err != nil {
			return sequence.MetaTxnStatusUnknown, nil, err
		}

		var receipt *types.Receipt
		// Only unmarshal TxnReceipt when the field is non-empty, since attempting to decode an empty JSON string would result in an error.
		if metaTxnReceipt.TxnReceipt != "" {
			if err = json.Unmarshal([]byte(metaTxnReceipt.TxnReceipt), &receipt); err != nil {
				return 0, nil, fmt.Errorf("decode txn receipt data: %w", err)
			}
		}

		return MetaTxnStatusFromString(metaTxnReceipt.Status), receipt, nil
	}
}

func (r *Client) IsDeployTransaction(signedTxs *sequence.SignedTransactions) bool {
	for _, txn := range signedTxs.Transactions {
		if txn.To == signedTxs.WalletContext.FactoryAddress {
			return true
		}
	}
	return false
}

func (r *Client) Client() proto.RelayerClient {
	return r.RelayerClient
}
