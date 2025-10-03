package receipts

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/v2/walletmain"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Status int

const (
	StatusNotExecuted Status = iota
	StatusSucceeded
	StatusFailed
	StatusAborted
)

type Receipt struct {
	Status Status
	Error  error
	Logs   []*types.Log
}

func TransactionReceipts(ctx context.Context, transaction common.Hash, provider *ethrpc.Provider, chainID ...*big.Int) ([]Receipt, error) {
	if provider == nil {
		return nil, fmt.Errorf("no provider")
	}

	if len(chainID) == 0 {
		chainID_, err := provider.ChainID(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to get chain id: %w", err)
		}

		chainID = append(chainID, chainID_)
	}

	receipt, err := provider.TransactionReceipt(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("unable to get transaction %v receipt: %w", transaction, err)
	}

	transaction_, _, err := provider.TransactionByHash(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("unable to get transaction %v: %w", transaction, err)
	}

	to := transaction_.To()
	if to == nil || *to == (common.Address{}) {
		return nil, fmt.Errorf("no to address for transaction %v", transaction)
	}

	return decodeLogs(ctx, *to, transaction_.Data(), receipt.Logs, chainID[0], provider)
}

func decodeLogs(ctx context.Context, to common.Address, data []byte, logs []*types.Log, chainID *big.Int, provider *ethrpc.Provider) ([]Receipt, error) {
	if execute := contracts.V3.WalletStage1Module.ABI.Methods["execute"]; bytes.HasPrefix(data, execute.ID) {
		args, err := execute.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("unable to decode v3 execute: %w", err)
		} else if len(args) != 2 {
			return nil, fmt.Errorf("%v v3 execute arguments, expected two", len(args))
		}

		payload, ok := args[0].([]byte)
		if !ok {
			return nil, fmt.Errorf("v3 execute payload argument is %T, expected []byte", args[0])
		}

		payload_, err := v3.DecodeCalls(to, chainID, payload)
		if err != nil {
			return nil, fmt.Errorf("unable to decode v3 execute payload: %w", err)
		}

		return decodeLogsV3(ctx, payload_, logs, provider)
	} else if selfExecute := contracts.V3.WalletStage1Module.ABI.Methods["selfExecute"]; bytes.HasPrefix(data, selfExecute.ID) {
		args, err := selfExecute.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("unable to decode v3 selfExecute: %w", err)
		} else if len(args) != 1 {
			return nil, fmt.Errorf("%v v3 selfExecute arguments, expected one", len(args))
		}

		payload, ok := args[0].([]byte)
		if !ok {
			return nil, fmt.Errorf("v3 selfExecute payload argument is %T, expected []byte", args[0])
		}

		payload_, err := v3.DecodeCalls(to, chainID, payload)
		if err != nil {
			return nil, fmt.Errorf("unable to decode v3 selfExecute payload: %w", err)
		}

		return decodeLogsV3(ctx, payload_, logs, provider)
	} else if execute := contracts.V2.WalletMainModule.ABI.Methods["execute"]; bytes.HasPrefix(data, execute.ID) {
		args, err := execute.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("unable to decode v2 execute: %w", err)
		} else if len(args) != 3 {
			return nil, fmt.Errorf("%v v2 execute arguments, expected three", len(args))
		}

		transactions, ok := args[0].([]walletmain.IModuleCallsTransaction)
		if !ok {
			return nil, fmt.Errorf("v2 execute transactions argument is %T, expected []walletmain.IModuleCallsTransaction", args[0])
		}

		nonce, ok := args[1].(*big.Int)
		if !ok {
			return nil, fmt.Errorf("v2 execute nonce argument is %T, expected *big.Int", args[1])
		}

	} else if selfExecute := contracts.V2.WalletMainModule.ABI.Methods["selfExecute"]; bytes.HasPrefix(data, selfExecute.ID) {
		args, err := selfExecute.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, fmt.Errorf("unable to decode v2 selfExecute: %w", err)
		} else if len(args) != 1 {
			return nil, fmt.Errorf("%v v2 selfExecute arguments, expected one", len(args))
		}

		transactions, ok := args[0].([]walletmain.IModuleCallsTransaction)
		if !ok {
			return nil, fmt.Errorf("v2 selfExecute transactions argument is %T, expected []walletmain.IModuleCallsTransaction", args[0])
		}

	} else if payload, err := v3.DecodeCalls(to, chainID, data); err == nil {
		return decodeLogsV3(ctx, payload, logs, provider)
	} else if to_, data_, err := sequence.DecompressCalldata(ctx, &to, data, provider); err == nil {
		return decodeLogs(ctx, to_, data_, logs, chainID, provider)
	} else {
		return nil, fmt.Errorf("not sequence transaction data: %v", hexutil.Encode(data))
	}
}

func decodeLogsV3(ctx context.Context, payload v3.CallsPayload, logs []*types.Log, provider *ethrpc.Provider) ([]Receipt, error) {
	receipts := make([]Receipt, len(payload.Calls))

	digest := payload.Digest()

	var index, start int
	for i, log := range logs {
		if digest_, index_, err := decodeCallSucceeded(log); err == nil {
			if digest_ == digest.Hash {
				if index_.Cmp(big.NewInt(int64(index))) != 0 {
					return nil, fmt.Errorf("index is %v, expected %v", index_, index)
				}
				receipts[index] = Receipt{
					Status: StatusSucceeded,
					Logs:   logs[start:i],
				}
				index++
				start = i + 1
			}
		} else if digest_, index_, reason, err := decodeCallFailed(log); err == nil {
			if digest_ == digest.Hash {
				if index_.Cmp(big.NewInt(int64(index))) != 0 {
					return nil, fmt.Errorf("index is %v, expected %v", index_, index)
				}
				receipts[index] = Receipt{
					Status: StatusFailed,
					Error:  reason,
					Logs:   logs[start:i],
				}
				index++
				start = i + 1
			}
		} else if digest_, index_, reason, err := decodeCallAborted(log); err == nil {
			if digest_ == digest.Hash {
				if index_.Cmp(big.NewInt(int64(index))) != 0 {
					return nil, fmt.Errorf("index is %v, expected %v", index_, index)
				}
				receipts[index] = Receipt{
					Status: StatusAborted,
					Error:  reason,
					Logs:   logs[start:i],
				}
				index++
				start = i + 1
			}
		} else if digest_, index_, err := decodeCallSkipped(log); err == nil {
			if digest_ == digest.Hash {
				if index_.Cmp(big.NewInt(int64(index))) != 0 {
					return nil, fmt.Errorf("index is %v, expected %v", index_, index)
				}
				if i != start {
					return nil, fmt.Errorf("skipped call %v has logs", index)
				}
				receipts[index] = Receipt{
					Status: StatusNotExecuted,
					Logs:   logs[start:i],
				}
				index++
				start = i + 1
			}
		}
	}

	return receipts, nil
}

func decodeCallSucceeded(log *types.Log) (common.Hash, *big.Int, error) {
	if len(log.Topics) != 1 {
		return common.Hash{}, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V3.WalletStage1Module.ABI.Events["CallSucceeded"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, fmt.Errorf("not CallSucceeded")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("unable to decode CallSucceeded: %w", err)
	}
	if len(args) != 2 {
		return common.Hash{}, nil, fmt.Errorf("%v CallSucceeded arguments, expected two: %w", len(args), err)
	}

	digest, ok := args[0].(common.Hash)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSucceeded digest is %T, expected common.Hash", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSucceeded index is %T, expected *big.Int", args[1])
	}

	return digest, index, nil
}

func decodeCallFailed(log *types.Log) (common.Hash, *big.Int, error, error) {
	if len(log.Topics) != 1 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V3.WalletStage1Module.ABI.Events["CallFailed"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, nil, fmt.Errorf("not CallFailed")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, nil, fmt.Errorf("unable to decode CallFailed: %w", err)
	}
	if len(args) != 3 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v CallFailed arguments, expected three: %w", len(args), err)
	}

	digest, ok := args[0].(common.Hash)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallFailed digest is %T, expected common.Hash", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallFailed index is %T, expected *big.Int", args[1])
	}

	reason, ok := args[2].([]byte)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallFailed reason is %T, expected []byte", args[2])
	}

	var reason_ error
	if reason__, err := abi.UnpackRevert(reason); err == nil {
		reason_ = fmt.Errorf("%v", reason__)
	} else {
		reason_ = fmt.Errorf("%v", hexutil.Encode(reason))
	}

	return digest, index, reason_, nil
}

func decodeCallAborted(log *types.Log) (common.Hash, *big.Int, error, error) {
	if len(log.Topics) != 1 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V3.WalletStage1Module.ABI.Events["CallAborted"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, nil, fmt.Errorf("not CallAborted")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, nil, fmt.Errorf("unable to decode CallAborted: %w", err)
	}
	if len(args) != 3 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v CallAborted arguments, expected three: %w", len(args), err)
	}

	digest, ok := args[0].(common.Hash)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallAborted digest is %T, expected common.Hash", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallAborted index is %T, expected *big.Int", args[1])
	}

	reason, ok := args[2].([]byte)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("CallAborted reason is %T, expected []byte", args[2])
	}

	var reason_ error
	if reason__, err := abi.UnpackRevert(reason); err == nil {
		reason_ = fmt.Errorf("%v", reason__)
	} else {
		reason_ = fmt.Errorf("%v", hexutil.Encode(reason))
	}

	return digest, index, reason_, nil
}

func decodeCallSkipped(log *types.Log) (common.Hash, *big.Int, error) {
	if len(log.Topics) != 1 {
		return common.Hash{}, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V3.WalletStage1Module.ABI.Events["CallSkipped"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, fmt.Errorf("not CallSkipped")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("unable to decode CallSkipped: %w", err)
	}
	if len(args) != 2 {
		return common.Hash{}, nil, fmt.Errorf("%v CallSkipped arguments, expected two: %w", len(args), err)
	}

	digest, ok := args[0].(common.Hash)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSkipped digest is %T, expected common.Hash", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSkipped index is %T, expected *big.Int", args[1])
	}

	return digest, index, nil
}
