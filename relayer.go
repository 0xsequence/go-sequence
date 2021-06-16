package sequence

import (
	"context"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type Relayer interface {
	EstimateGasLimits(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, transactions []interface{}) ([]interface{}, error)

	GetNonce(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, space *uint64, blockNum *big.Int) (*big.Int, error)

	Relay(ctx context.Context, signedTxs interface{}) (interface{}, error)

	// TODO: what "type" should metaTxID be..? it is a hash..
	// but, its kinda different cuz we dont require 0x so may not be able to use common.Hash
	// alternatively we could make metaTxID a different length of hash to differentiate it, but its a breaking change on client-side
	Wait(ctx context.Context, metaTxID common.Hash, timeout time.Duration) (interface{}, error)

	// GasRefundOptions() // TODO, in future when needed..
}

// TODO: a simple / bare-bones relayer, etc.....
type LocalRelayer struct {
}

// TODO: lets make some Transaction XXX types..

// TODO: let's implement LocalRelayer

// TODO: lets add wallet deploy method, on wallet itself..  via Relayer..
// we can prob add a NewWalletDeployTransaction(walletConfig, walletContext) method on sequence.XX too...
// ..
