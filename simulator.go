package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/v3/walletsimulator"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type SimulatorResult struct {
	Status  SimulatorResultStatus
	Result  []byte
	GasUsed *big.Int
}

type SimulatorResultStatus int

const (
	SimulatorResultStatusSkipped SimulatorResultStatus = iota
	SimulatorResultStatusSucceeded
	SimulatorResultStatusFailed
	SimulatorResultStatusAborted
	SimulatorResultStatusReverted
	SimulatorResultStatusNotEnoughGas
)

func Simulate(ctx context.Context, wallet common.Address, transactions Transactions, provider *ethrpc.Provider) ([]SimulatorResult, error) {
	calls := make([]walletsimulator.PayloadCall, 0, len(transactions))
	for _, transaction := range transactions {
		behaviorOnError := new(big.Int).SetInt64(int64(v3.BehaviorOnErrorIgnore))
		if transaction.RevertOnError {
			behaviorOnError.SetInt64(int64(v3.BehaviorOnErrorRevert))
		}
		calls = append(calls, walletsimulator.PayloadCall{
			To:              transaction.To,
			Value:           transaction.Value,
			Data:            transaction.Data,
			GasLimit:        transaction.GasLimit,
			DelegateCall:    transaction.DelegateCall,
			BehaviorOnError: behaviorOnError,
		})
	}

	data, err := contracts.V3.WalletSimulator.Encode("simulate", calls)
	if err != nil {
		return nil, fmt.Errorf("unable to encode calls: %w", err)
	}

	response, err := provider.Do(ctx, ethrpc.NewCall("eth_call", map[string]any{
		"to":    wallet,
		"input": data,
	}, map[string]any{
		wallet.String(): map[string]any{"code": contracts.V3.WalletSimulator.DeployedBin},
	}))
	if err != nil {
		return nil, fmt.Errorf("unable to simulate: %w", err)
	}

	var results []walletsimulator.SimulatorResult
	if err := contracts.V3.WalletSimulator.Decode(&results, "eth_call", response); err != nil {
		return nil, fmt.Errorf("unable to decode results: %w", err)
	}

	results_ := make([]SimulatorResult, 0, len(results))
	for _, result := range results {
		results_ = append(results_, SimulatorResult{
			Status:  SimulatorResultStatus(result.Status),
			Result:  result.Result,
			GasUsed: result.GasUsed,
		})
	}

	return results_, nil
}
