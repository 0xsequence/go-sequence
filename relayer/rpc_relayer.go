package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"time"

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
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type RpcRelayer struct {
	provider        *ethrpc.Provider
	receiptListener *ethreceipts.ReceiptsListener
	Service         proto.Relayer
}

var _ sequence.Relayer = &RpcRelayer{}

func NewRpcRelayer(provider *ethrpc.Provider, receiptListener *ethreceipts.ReceiptsListener, rpcRelayerURL string, httpClient proto.HTTPClient) (*RpcRelayer, error) {
	_, err := url.Parse(rpcRelayerURL)
	if err != nil {
		return nil, fmt.Errorf("rpcRelayerURL is invalid: %w", err)
	}

	service := proto.NewRelayerClient(rpcRelayerURL, httpClient)

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

	config, err := r.protoConfig(ctx, walletConfig, walletAddress)
	if err != nil {
		return nil, err
	}

	response, err := r.Service.UpdateMetaTxnGasLimits(ctx, walletAddress.Hex(), config, hexutil.Encode(requestData))
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

	// send to guest module if factory address is the recipient (in order to deploy wallet)
	// todo: split wallet create and other transactions
	for _, txn := range signedTxs.Transactions {
		if txn.To == signedTxs.WalletContext.FactoryAddress {
			to = signedTxs.WalletContext.GuestModuleAddress
			break
		}
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
		_, receipt, err := r.Wait(ctx, sequence.MetaTxnID(metaTxnID))
		return receipt, err
	}

	// TODO: 2nd argument will be nil, we may even want to remove it from here...
	return sequence.MetaTxnID(metaTxnID), nil, waitReceipt, nil
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
		return proto.MetaTxnStatusFromString(metaTxnReceipt.Status), receipt, nil
	}
}

func (r *RpcRelayer) protoConfig(ctx context.Context, config core.WalletConfig, walletAddress common.Address) (*proto.WalletConfig, error) {
	if walletConfigV1 := config.(*v1.WalletConfig); walletConfigV1 != nil {
		var signers []*proto.WalletSigner
		for _, signer := range walletConfigV1.Signers_ {
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
			Threshold: walletConfigV1.Threshold_,
			ChainId:   &chainID,
		}, nil
	} else if walletConfigV2 := config.(*v2.WalletConfig); walletConfigV2 != nil {
		// todo: implement v2
	}

	return nil, fmt.Errorf("relayer: unknown wallet config version")
}
