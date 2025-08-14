package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
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

func v3Calldata(to common.Address, chainID *big.Int, transaction *Transaction) []byte {
	if len(transaction.Data) != 0 {
		return transaction.Data
	}

	if len(transaction.Transactions) != 0 {
		calls := make([]v3.Call, 0, len(transaction.Transactions))
		for _, transaction_ := range transaction.Transactions {
			var behaviorOnError v3.BehaviorOnError
			if transaction_.RevertOnError {
				behaviorOnError = v3.BehaviorOnErrorRevert
			}

			calls = append(calls, v3.Call{
				To:              transaction_.To,
				Value:           transaction_.Value,
				Data:            v3Calldata(transaction.To, chainID, transaction_),
				GasLimit:        transaction_.GasLimit,
				DelegateCall:    transaction_.DelegateCall,
				BehaviorOnError: behaviorOnError,
			})
		}

		packed := v3.NewCallsPayload(to, chainID, calls, big.NewInt(123456789), nil, nil).Encode(to)

		if len(transaction.Signature) != 0 {
			calldata, _ := contracts.V3.WalletStage1Module.Encode("execute", packed, transaction.Signature)
			return calldata
		} else {
			calldata, _ := contracts.V3.WalletStage1Module.Encode("selfExecute", packed)
			return calldata
		}
	}

	return nil
}

func Simulate(ctx context.Context, wallet common.Address, transactions Transactions, provider *ethrpc.Provider) ([]SimulatorResult, error) {
	chainID, err := provider.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get chain id: %v", err)
	}

	calls := make([]walletsimulator.PayloadCall, 0, len(transactions))
	for _, transaction := range transactions {
		behaviorOnError := new(big.Int).SetInt64(int64(v3.BehaviorOnErrorIgnore))
		if transaction.RevertOnError {
			behaviorOnError.SetInt64(int64(v3.BehaviorOnErrorRevert))
		}
		calls = append(calls, walletsimulator.PayloadCall{
			To:              transaction.To,
			Value:           transaction.Value,
			Data:            v3Calldata(wallet, chainID, transaction),
			GasLimit:        transaction.GasLimit,
			DelegateCall:    transaction.DelegateCall,
			BehaviorOnError: behaviorOnError,
		})
	}

	data, err := contracts.V3.WalletSimulator.Encode("simulate", calls)
	if err != nil {
		return nil, fmt.Errorf("unable to encode calls: %w", err)
	}

	params := struct {
		To   common.Address `json:"to"`
		Data string         `json:"input"`
	}{
		To:   wallet,
		Data: hexutil.Encode(data),
	}

	overrides := map[string]any{
		wallet.String(): map[string]any{
			"code": hexutil.Encode(contracts.V3.WalletSimulator.DeployedBin),
		},
	}

	var response string
	call := ethrpc.NewCallBuilder[string]("eth_call", nil, params, "latest", overrides)
	_, err = provider.Do(ctx, call.Into(&response))
	if err != nil {
		return nil, fmt.Errorf("unable to simulate: %w", err)
	}

	decoded, err := hexutil.Decode(response)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	var results []walletsimulator.SimulatorResult
	if err := contracts.V3.WalletSimulator.Decode(&results, "simulate", decoded); err != nil {
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
