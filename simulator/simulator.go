package simulator

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/ethclient/gethclient"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/v2/walletgasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/v3/walletsimulator"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Status int

const (
	StatusSkipped Status = iota
	StatusSucceeded
	StatusFailed
	StatusAborted
	StatusReverted
	StatusNotEnoughGas
)

type Result struct {
	Status  Status
	Result  []byte
	Error   error
	GasUsed uint64
}

func SimulateV2(ctx context.Context, wallet common.Address, transactions []walletgasestimator.IModuleCallsTransaction, provider *ethrpc.Provider) ([]Result, error) {
	overrides := map[common.Address]gethclient.OverrideAccount{
		wallet: {Code: contracts.V2.WalletGasEstimator.DeployedBin},
	}

	input, err := contracts.V2.WalletGasEstimator.Encode("simulateExecute", transactions)
	if err != nil {
		return nil, fmt.Errorf("unable to encode simulateExecute call: %w", err)
	}

	var encoded string
	_, err = provider.Do(ctx, ethrpc.NewCallBuilder[string]("eth_call", nil, map[string]any{"to": wallet, "input": hexutil.Encode(input)}, "latest", overrides).Into(&encoded))
	if err != nil {
		return nil, fmt.Errorf("unable to simulateExecute: %w", err)
	}

	decoded, err := hexutil.Decode(encoded)
	if err != nil {
		return nil, fmt.Errorf("unable to decode simulateExecute result: %w", err)
	}

	var results []walletgasestimator.MainModuleGasEstimationSimulateResult
	err = contracts.V2.WalletGasEstimator.Decode(&results, "simulateExecute", decoded)
	if err != nil {
		return nil, fmt.Errorf("unable to decode simulateExecute result: %w", err)
	}

	if len(results) != len(transactions) {
		return nil, fmt.Errorf("%v simulateExecute results for %v transactions", len(results), len(transactions))
	}

	results_ := make([]Result, 0, len(results))
	for i, result := range results {
		var result_ Result

		if result.Executed {
			if result.Succeeded {
				result_.Status = StatusSucceeded
			} else {
				if transactions[i].RevertOnError {
					result_.Status = StatusReverted
				} else {
					result_.Status = StatusFailed
				}

				revert, err := abi.UnpackRevert(result.Result)
				if err == nil {
					result_.Error = fmt.Errorf("%v", revert)
				} else {
					result_.Error = fmt.Errorf("%v", hexutil.Encode(result.Result))
				}
			}
		} else {
			result_.Status = StatusSkipped
		}

		result_.Result = result.Result
		result_.GasUsed = result.GasUsed.Uint64()

		results_ = append(results_, result_)
	}

	return results_, nil
}

func SimulateV3(ctx context.Context, wallet common.Address, calls []v3.Call, provider *ethrpc.Provider) ([]Result, error) {
	overrides := map[common.Address]gethclient.OverrideAccount{
		wallet: {Code: contracts.V3.WalletSimulator.DeployedBin},
	}

	calls_ := make([]walletsimulator.PayloadCall, 0, len(calls))
	for _, call := range calls {
		calls_ = append(calls_, walletsimulator.PayloadCall{
			To:              call.To,
			Value:           call.Value,
			Data:            call.Data,
			GasLimit:        call.GasLimit,
			DelegateCall:    call.DelegateCall,
			OnlyFallback:    call.OnlyFallback,
			BehaviorOnError: big.NewInt(int64(call.BehaviorOnError)),
		})
	}

	input, err := contracts.V3.WalletSimulator.Encode("simulate", calls_)
	if err != nil {
		return nil, fmt.Errorf("unable to encode simulate call: %w", err)
	}

	var encoded string
	_, err = provider.Do(ctx, ethrpc.NewCallBuilder[string]("eth_call", nil, map[string]any{"to": wallet, "input": hexutil.Encode(input)}, "latest", overrides).Into(&encoded))
	if err != nil {
		return nil, fmt.Errorf("unable to simulate: %w", err)
	}

	decoded, err := hexutil.Decode(encoded)
	if err != nil {
		return nil, fmt.Errorf("unable to decode simulate result: %w", err)
	}

	var results []walletsimulator.SimulatorResult
	err = contracts.V3.WalletSimulator.Decode(&results, "simulate", decoded)
	if err != nil {
		return nil, fmt.Errorf("unable to decode simulate result: %w", err)
	}

	if len(results) != len(calls) {
		return nil, fmt.Errorf("%v simulate results for %v calls", len(results), len(calls))
	}

	results_ := make([]Result, 0, len(results))
	for _, result := range results {
		var result_ Result

		result_.Status = Status(result.Status)
		result_.Result = result.Result
		result_.GasUsed = result.GasUsed.Uint64()

		switch result_.Status {
		case StatusFailed, StatusAborted, StatusReverted:
			revert, err := abi.UnpackRevert(result.Result)
			if err == nil {
				result_.Error = fmt.Errorf("%v", revert)
			} else {
				result_.Error = fmt.Errorf("%v", hexutil.Encode(result.Result))
			}
		}

		results_ = append(results_, result_)
	}

	return results_, nil
}
