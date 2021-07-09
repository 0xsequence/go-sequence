package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
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

	finalOverwrites := map[common.Address]*Overwrite{
		call.From: {
			Code: contracts.WalletGasEstimator.Bin,
		},
	}

	for key, value := range overwrites {
		if key == *&call.From {
			return fmt.Errorf("can't overwride address from")
		}

		finalOverwrites[key] = value
	}

	var res interface{}
	err := provider.RPC.Call(res, "eth_call", call, blockTag, finalOverwrites)
	if err != nil {
		return err
	}

	fmt.Println("response", res)

	return nil
}
