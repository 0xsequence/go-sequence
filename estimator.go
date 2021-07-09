package sequence

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type StateOverwrite struct {
	Key   []byte
	Value []byte
}

type Overwrite struct {
	Code      []byte
	Balance   *big.Int
	Nonce     *big.Int
	StateDiff []*StateOverwrite
	State     []*StateOverwrite
}

func EstimateCall(ctx context.Context, provider *ethrpc.Provider, call *ethtxn.TransactionRequest, overwrites map[common.Address]*Overwrite, blockTag string) error {
	if blockTag == "" {
		blockTag = "latest"
	}

	estimateCall := &ethtxn.TransactionRequest{
		From:     call.From,
		Data:     call.Data,
		Nonce:    call.Nonce,
		ETHValue: call.ETHValue,
	}

	if call.Data != nil {
		estimateCall.Data = call.Data
	} else {
		estimateCall.Data = []byte{}
	}

	var res interface{}

	finalOverwrites := map[common.Address]*Overwrite{
		call.From: &Overwrite{
			Code: contracts.Wallet,
		},
	}

	// provider.RPC.Call(res, "eth_call", call, blockTag, interface{}{ a: 24123 })

	return nil
}
