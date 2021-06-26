package relayer

import (
	"context"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
)

// TODO ....
type RpcRelayer struct {
	// Service ...
}

var _ sequence.Relayer = &RpcRelayer{}

func (r *RpcRelayer) GetProvider() *ethrpc.Provider {
	return nil
}

// ..
func (r *RpcRelayer) EstimateGasLimits(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, transactions sequence.Transactions) (sequence.Transactions, error) {
	return nil, nil
}

// NOTE: nonce space is 160 bits wide
func (r *RpcRelayer) GetNonce(ctx context.Context, walletConfig sequence.WalletConfig, walletContext sequence.WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	return nil, nil
}

// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
func (r *RpcRelayer) Relay(ctx context.Context, signedTxs *sequence.SignedTransactions) (sequence.MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	return "", nil, nil, nil
}

// ..
func (r *RpcRelayer) Wait(ctx context.Context, metaTxID sequence.MetaTxnID, timeout time.Duration) (sequence.MetaTxnStatus, *types.Receipt, error) {
	return 0, nil, nil
}
