package sequence

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
)

type StateOverwrite struct {
	Key   string
	Value string
}

type Overwrite struct {
	Code      string
	Balance   *big.Int
	Nonce     *big.Int
	StateDiff []*StateOverwrite
	State     []*StateOverwrite
}

type EstimateTransaction struct {
	From common.Address
	To   common.Address
	Data []byte
}

type Estimator struct {
	BaseCost     uint64
	DataOneCost  uint64
	DataZeroCost uint64
}

var defaultEstimator = &Estimator{
	BaseCost:     21000,
	DataOneCost:  16,
	DataZeroCost: 4,
}

func NewEstimator() *Estimator {
	return &Estimator{
		BaseCost:     defaultEstimator.BaseCost,
		DataZeroCost: defaultEstimator.DataZeroCost,
		DataOneCost:  defaultEstimator.DataOneCost,
	}
}

func (e *Estimator) CalldataCost(data []byte) uint64 {
	cost := e.BaseCost

	for _, b := range data {
		if b == 0 {
			cost += e.DataZeroCost
		} else {
			cost += e.DataOneCost
		}
	}

	return cost
}

func (e *Estimator) EstimateCall(ctx context.Context, provider *ethrpc.Provider, call *EstimateTransaction, overwrites map[common.Address]*Overwrite, blockTag string) (*big.Int, error) {
	if blockTag == "" {
		blockTag = "latest"
	}

	from := call.From
	if from.Hash().Big().Cmp(common.Big0) == 0 {
		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		if err != nil {
			return nil, err
		}

		from = wallet.Address()
	}

	finalOverwrites := map[common.Address]*Overwrite{
		from: {
			Code: contracts.GasEstimatorDeployedBytecode,
		},
	}

	if overwrites != nil {
		for key, value := range overwrites {
			if key == call.From {
				return nil, fmt.Errorf("can't overwride address from")
			}

			finalOverwrites[key] = value
		}
	}

	estimator := ethcontract.NewContractCaller(from, contracts.GasEstimator.ABI, provider)
	callData, err := estimator.Encode("estimate", call.To, call.Data)
	if err != nil {
		return nil, err
	}

	type Call struct {
		To   common.Address `json:"to"`
		Data string         `json:"data"`
	}

	estimateCall := &Call{
		To:   from,
		Data: "0x" + common.Bytes2Hex(callData),
	}

	var res string
	err = provider.RPC.Call(&res, "eth_call", estimateCall, blockTag, finalOverwrites)
	if err != nil {
		return nil, err
	}

	resBytes := common.Hex2Bytes(strings.Replace(res, "0x", "", 1))

	var success bool
	var result []byte
	var gas *big.Int

	if err := ethcoder.AbiDecoder([]string{"bool", "bytes", "uint256"}, resBytes, []interface{}{&success, &result, &gas}); err != nil {
		return nil, err
	}

	gas.Add(gas, big.NewInt(int64(e.CalldataCost(call.Data))))

	return gas, nil
}
