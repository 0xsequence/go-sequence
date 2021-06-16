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
	// but, its kinda different cuz we dont require 0x right..?
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

//////---------

// STEPS.....

// 1. read sequence.js code and how we work with txns.... pretty much, we just encode the data field, and send.. okay..

// 2. read old sdk code how we do it.....

// 3. read skyweaver-api code how we're doing parallel txns etc.......
