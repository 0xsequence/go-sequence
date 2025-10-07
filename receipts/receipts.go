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
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Status int

const (
	StatusNotExecuted Status = iota
	StatusSucceeded
	StatusFailed
	StatusAborted
)

type Receipts struct {
	Digest   common.Hash
	Receipts []Receipt
}

func (r *Receipts) Find(digest common.Hash) *Receipts {
	if digest == r.Digest {
		return r
	}

	for i := range r.Receipts {
		receipt := &r.Receipts[i]
		if receipt.Receipts != nil {
			receipts := receipt.Receipts.Find(digest)
			if receipts != nil {
				return receipts
			}
		}
	}

	return nil
}

type Receipt struct {
	Status   Status
	Error    error
	Logs     []*types.Log
	Receipts *Receipts
}

func TransactionReceipts(ctx context.Context, transaction common.Hash, provider *ethrpc.Provider, chainID ...*big.Int) (Receipts, error) {
	if provider == nil {
		return Receipts{}, fmt.Errorf("no provider")
	}

	receipt, err := provider.TransactionReceipt(ctx, transaction)
	if err != nil {
		return Receipts{}, fmt.Errorf("unable to get transaction %v receipt: %w", transaction, err)
	}

	return TransactionReceiptsForReceipt(ctx, receipt, provider, chainID...)
}

func TransactionReceiptsForReceipt(ctx context.Context, receipt *types.Receipt, provider *ethrpc.Provider, chainID ...*big.Int) (Receipts, error) {
	if provider == nil {
		return Receipts{}, fmt.Errorf("no provider")
	}

	if len(chainID) == 0 {
		chainID_, err := provider.ChainID(ctx)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to get chain id: %w", err)
		}

		chainID = append(chainID, chainID_)
	}

	transaction_, _, err := provider.TransactionByHash(ctx, receipt.TxHash)
	if err != nil {
		return Receipts{}, fmt.Errorf("unable to get transaction %v: %w", receipt.TxHash, err)
	}

	to := transaction_.To()
	if to == nil || *to == (common.Address{}) {
		return Receipts{}, fmt.Errorf("no to address for transaction %v", receipt.TxHash)
	}

	return decodeLogs(ctx, *to, transaction_.Data(), receipt.Logs, chainID[0], provider)
}

func decodeLogs(ctx context.Context, to common.Address, data []byte, logs []*types.Log, chainID *big.Int, provider *ethrpc.Provider) (Receipts, error) {
	if execute := contracts.V3.WalletStage1Module.ABI.Methods["execute"]; bytes.HasPrefix(data, execute.ID) {
		args, err := execute.Inputs.Unpack(data[4:])
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v3 execute: %w", err)
		} else if len(args) != 2 {
			return Receipts{}, fmt.Errorf("%v v3 execute arguments, expected two", len(args))
		}

		payload, ok := args[0].([]byte)
		if !ok {
			return Receipts{}, fmt.Errorf("v3 execute payload argument is %T, expected []byte", args[0])
		}

		payload_, err := v3.DecodeCalls(to, chainID, payload)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v3 execute payload: %w", err)
		}

		digest, receipts, err := decodeLogsV3(payload_, trimNonceChange(logs, to))
		if err != nil {
			return Receipts{}, err
		}
		if len(receipts) != len(payload_.Calls) {
			return Receipts{}, fmt.Errorf("%v receipts for %v calls", len(receipts), len(payload_.Calls))
		}

		for i, call := range payload_.Calls {
			receipts_, err := decodeLogs(ctx, call.To, call.Data, receipts[i].Logs, chainID, provider)
			if err == nil {
				receipts[i].Receipts = &receipts_
			}
		}

		return Receipts{Digest: digest.Hash, Receipts: receipts}, nil
	} else if selfExecute := contracts.V3.WalletStage1Module.ABI.Methods["selfExecute"]; bytes.HasPrefix(data, selfExecute.ID) {
		args, err := selfExecute.Inputs.Unpack(data[4:])
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v3 selfExecute: %w", err)
		} else if len(args) != 1 {
			return Receipts{}, fmt.Errorf("%v v3 selfExecute arguments, expected one", len(args))
		}

		payload, ok := args[0].([]byte)
		if !ok {
			return Receipts{}, fmt.Errorf("v3 selfExecute payload argument is %T, expected []byte", args[0])
		}

		payload_, err := v3.DecodeCalls(to, chainID, payload)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v3 selfExecute payload: %w", err)
		}

		digest, receipts, err := decodeLogsV3(payload_, logs)
		if err != nil {
			return Receipts{}, err
		}
		if len(receipts) != len(payload_.Calls) {
			return Receipts{}, fmt.Errorf("%v receipts for %v calls", len(receipts), len(payload_.Calls))
		}

		for i, call := range payload_.Calls {
			receipts_, err := decodeLogs(ctx, call.To, call.Data, receipts[i].Logs, chainID, provider)
			if err == nil {
				receipts[i].Receipts = &receipts_
			}
		}

		return Receipts{Digest: digest.Hash, Receipts: receipts}, nil
	} else if execute := contracts.V2.WalletMainModule.ABI.Methods["execute"]; bytes.HasPrefix(data, execute.ID) {
		args, err := execute.Inputs.Unpack(data[4:])
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v2 execute: %w", err)
		} else if len(args) != 3 {
			return Receipts{}, fmt.Errorf("%v v2 execute arguments, expected three", len(args))
		}

		transactions, ok := args[0].([]struct {
			DelegateCall  bool           `json:"delegateCall"`
			RevertOnError bool           `json:"revertOnError"`
			GasLimit      *big.Int       `json:"gasLimit"`
			Target        common.Address `json:"target"`
			Value         *big.Int       `json:"value"`
			Data          []uint8        `json:"data"`
		})
		if !ok {
			return Receipts{}, fmt.Errorf("v2 execute transactions argument is %T", args[0])
		}

		nonce, ok := args[1].(*big.Int)
		if !ok {
			return Receipts{}, fmt.Errorf("v2 execute nonce argument is %T, expected *big.Int", args[1])
		}

		transactions_ := make(sequence.Transactions, 0, len(transactions))
		for _, transaction := range transactions {
			transactions_ = append(transactions_, &sequence.Transaction{
				DelegateCall:  transaction.DelegateCall,
				RevertOnError: transaction.RevertOnError,
				GasLimit:      transaction.GasLimit,
				To:            transaction.Target,
				Value:         transaction.Value,
				Data:          transaction.Data,
			})
		}

		guestDigest, err := sequence.ComputeGuestExecDigest(transactions_)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute guest digest: %w", err)
		}

		guestSubdigest, err := sequence.SubDigest(chainID, to, guestDigest)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute guest subdigest: %w", err)
		}
		guestSubdigest_ := common.BytesToHash(guestSubdigest)

		walletDigest, err := sequence.ComputeWalletExecDigest(nonce, transactions_)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute wallet digest: %w", err)
		}

		walletSubdigest, err := sequence.SubDigest(chainID, to, walletDigest)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute wallet subdigest: %w", err)
		}
		walletSubdigest_ := common.BytesToHash(walletSubdigest)

		subdigest, receipts, err := decodeLogsLegacy([]common.Hash{guestSubdigest_, walletSubdigest_}, trimNonceChange(logs, to))
		if err != nil {
			return Receipts{}, err
		}
		if len(receipts) != len(transactions_) {
			return Receipts{}, fmt.Errorf("%v receipts for %v transactions", len(receipts), len(transactions_))
		}

		for i, transaction := range transactions_ {
			receipts_, err := decodeLogs(ctx, transaction.To, transaction.Data, receipts[i].Logs, chainID, provider)
			if err == nil {
				receipts[i].Receipts = &receipts_
			}
		}

		return Receipts{Digest: subdigest, Receipts: receipts}, nil
	} else if selfExecute := contracts.V2.WalletMainModule.ABI.Methods["selfExecute"]; bytes.HasPrefix(data, selfExecute.ID) {
		args, err := selfExecute.Inputs.Unpack(data[4:])
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to decode v2 selfExecute: %w", err)
		} else if len(args) != 1 {
			return Receipts{}, fmt.Errorf("%v v2 selfExecute arguments, expected one", len(args))
		}

		transactions, ok := args[0].([]struct {
			DelegateCall  bool           `json:"delegateCall"`
			RevertOnError bool           `json:"revertOnError"`
			GasLimit      *big.Int       `json:"gasLimit"`
			Target        common.Address `json:"target"`
			Value         *big.Int       `json:"value"`
			Data          []uint8        `json:"data"`
		})
		if !ok {
			return Receipts{}, fmt.Errorf("v2 selfExecute transactions argument is %T", args[0])
		}

		transactions_ := make(sequence.Transactions, 0, len(transactions))
		for _, transaction := range transactions {
			transactions_ = append(transactions_, &sequence.Transaction{
				DelegateCall:  transaction.DelegateCall,
				RevertOnError: transaction.RevertOnError,
				GasLimit:      transaction.GasLimit,
				To:            transaction.Target,
				Value:         transaction.Value,
				Data:          transaction.Data,
			})
		}

		selfDigest, err := sequence.ComputeSelfExecDigest(transactions_)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute self digest: %w", err)
		}

		selfSubdigest, err := sequence.SubDigest(chainID, to, selfDigest)
		if err != nil {
			return Receipts{}, fmt.Errorf("unable to compute self subdigest: %w", err)
		}
		selfSubdigest_ := common.BytesToHash(selfSubdigest)

		subdigest, receipts, err := decodeLogsLegacy([]common.Hash{selfSubdigest_}, logs)
		if err != nil {
			return Receipts{}, err
		}
		if len(receipts) != len(transactions_) {
			return Receipts{}, fmt.Errorf("%v receipts for %v transactions", len(receipts), len(transactions_))
		}

		for i, transaction := range transactions_ {
			receipts_, err := decodeLogs(ctx, transaction.To, transaction.Data, receipts[i].Logs, chainID, provider)
			if err == nil {
				receipts[i].Receipts = &receipts_
			}
		}

		return Receipts{Digest: subdigest, Receipts: receipts}, nil
	} else if payload, err := v3.DecodeCalls(to, chainID, data); err == nil {
		digest, receipts, err := decodeLogsV3(payload, logs)
		if err != nil {
			return Receipts{}, err
		}
		if len(receipts) != len(payload.Calls) {
			return Receipts{}, fmt.Errorf("%v receipts for %v calls", len(receipts), len(payload.Calls))
		}

		for i, call := range payload.Calls {
			receipts_, err := decodeLogs(ctx, call.To, call.Data, receipts[i].Logs, chainID, provider)
			if err == nil {
				receipts[i].Receipts = &receipts_
			}
		}

		return Receipts{Digest: digest.Hash, Receipts: receipts}, nil
	} else if to_, data_, err := sequence.DecompressCalldata(ctx, &to, data, provider); err == nil {
		return decodeLogs(ctx, to_, data_, logs, chainID, provider)
	} else {
		return Receipts{}, fmt.Errorf("not sequence transaction data: %v", hexutil.Encode(data))
	}
}

func decodeLogsV3(payload v3.CallsPayload, logs []*types.Log) (core.PayloadDigest, []Receipt, error) {
	digest := payload.Digest()
	receipts := make([]Receipt, len(payload.Calls))

	var index, start int
	for i, log := range logs {
		if digest_, index_, err := decodeCallSucceeded(log); err == nil {
			if digest_ == digest.Hash {
				if index_.Cmp(big.NewInt(int64(index))) != 0 {
					return core.PayloadDigest{}, nil, fmt.Errorf("index is %v, expected %v", index_, index)
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
					return core.PayloadDigest{}, nil, fmt.Errorf("index is %v, expected %v", index_, index)
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
					return core.PayloadDigest{}, nil, fmt.Errorf("index is %v, expected %v", index_, index)
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
					return core.PayloadDigest{}, nil, fmt.Errorf("index is %v, expected %v", index_, index)
				}
				if i != start {
					return core.PayloadDigest{}, nil, fmt.Errorf("skipped call %v has logs", index)
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

	return digest, receipts, nil
}

func decodeLogsLegacy(subdigests []common.Hash, logs []*types.Log) (common.Hash, []Receipt, error) {
	var subdigest common.Hash
	var receipts []Receipt

	subdigests_ := make(map[common.Hash]struct{}, len(subdigests))
	for _, subdigest_ := range subdigests {
		subdigests_[subdigest_] = struct{}{}
	}

	var start int
	for i, log := range logs {
		if subdigest_, err := decodeTxExecutedV1(log); err == nil {
			if _, ok := subdigests_[subdigest_]; ok {
				subdigest = subdigest_
				receipts = append(receipts, Receipt{
					Status: StatusSucceeded,
					Logs:   logs[start:i],
				})
				start = i + 1
				continue
			}
		}

		if subdigest_, index_, err := decodeTxExecutedV2(log); err == nil {
			if _, ok := subdigests_[subdigest_]; ok {
				if index_.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return common.Hash{}, nil, fmt.Errorf("index is %v, expected %v", index_, len(receipts))
				}
				subdigest = subdigest_
				receipts = append(receipts, Receipt{
					Status: StatusSucceeded,
					Logs:   logs[start:i],
				})
				start = i + 1
			}
		} else if subdigest_, index_, reason, err := decodeTxFailedV2(log); err == nil {
			if _, ok := subdigests_[subdigest_]; ok {
				if index_.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return common.Hash{}, nil, fmt.Errorf("index is %v, expected %v", index_, len(receipts))
				}
				subdigest = subdigest_
				receipts = append(receipts, Receipt{
					Status: StatusFailed,
					Error:  reason,
					Logs:   logs[start:i],
				})
				start = i + 1
			}
		} else if subdigest_, reason, err := decodeTxFailedV1(log); err == nil {
			if _, ok := subdigests_[subdigest_]; ok {
				subdigest = subdigest_
				receipts = append(receipts, Receipt{
					Status: StatusFailed,
					Error:  reason,
					Logs:   logs[start:i],
				})
				start = i + 1
			}
		}
	}

	return subdigest, receipts, nil
}

func trimNonceChange(logs []*types.Log, wallet common.Address) []*types.Log {
	for i, log := range logs {
		if log.Address == wallet {
			if _, _, err := decodeNonceChange(log); err == nil {
				return logs[i+1:]
			}
		}
	}

	return logs
}

func decodeNonceChange(log *types.Log) (*big.Int, *big.Int, error) {
	if len(log.Topics) != 1 {
		return nil, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V3.WalletStage1Module.ABI.Events["NonceChange"]

	if log.Topics[0] != event.ID {
		return nil, nil, fmt.Errorf("not NonceChange")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to decode NonceChange: %w", err)
	}
	if len(args) != 2 {
		return nil, nil, fmt.Errorf("%v NonceChange arguments, expected two", len(args))
	}

	space, ok := args[0].(*big.Int)
	if !ok {
		return nil, nil, fmt.Errorf("NonceChange space is %T, expected *big.Int", args[0])
	}

	nonce, ok := args[1].(*big.Int)
	if !ok {
		return nil, nil, fmt.Errorf("NonceChange nonce is %T, expected *big.Int", args[1])
	}

	return space, nonce, nil
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
		return common.Hash{}, nil, fmt.Errorf("%v CallSucceeded arguments, expected two", len(args))
	}

	digest, ok := args[0].([common.HashLength]byte)
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
		return common.Hash{}, nil, nil, fmt.Errorf("%v CallFailed arguments, expected three", len(args))
	}

	digest, ok := args[0].([common.HashLength]byte)
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
		return common.Hash{}, nil, nil, fmt.Errorf("%v CallAborted arguments, expected three", len(args))
	}

	digest, ok := args[0].([common.HashLength]byte)
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
		return common.Hash{}, nil, fmt.Errorf("%v CallSkipped arguments, expected two", len(args))
	}

	digest, ok := args[0].([common.HashLength]byte)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSkipped digest is %T, expected common.Hash", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("CallSkipped index is %T, expected *big.Int", args[1])
	}

	return digest, index, nil
}

func decodeTxExecutedV2(log *types.Log) (common.Hash, *big.Int, error) {
	if len(log.Topics) != 2 {
		return common.Hash{}, nil, fmt.Errorf("%v topics, expected two", len(log.Topics))
	}

	event := contracts.V2.WalletMainModule.ABI.Events["TxExecuted"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, fmt.Errorf("not v2 TxExecuted")
	}

	subdigest := log.Topics[1]

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("unable to decode v2 TxExecuted: %w", err)
	}
	if len(args) != 1 {
		return common.Hash{}, nil, fmt.Errorf("%v v2 TxExecuted arguments, expected one", len(args))
	}

	index, ok := args[0].(*big.Int)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("v2 TxExecuted index is %T, expected *big.Int", args[0])
	}

	return subdigest, index, nil
}

func decodeTxFailedV2(log *types.Log) (common.Hash, *big.Int, error, error) {
	if len(log.Topics) != 2 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v topics, expected two", len(log.Topics))
	}

	event := contracts.V2.WalletMainModule.ABI.Events["TxFailed"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, nil, fmt.Errorf("not v2 TxFailed")
	}

	subdigest := log.Topics[1]

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, nil, fmt.Errorf("unable to decode v2 TxFailed: %w", err)
	}
	if len(args) != 2 {
		return common.Hash{}, nil, nil, fmt.Errorf("%v v2 TxFailed arguments, expected two", len(args))
	}

	index, ok := args[0].(*big.Int)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("v2 TxFailed index is %T, expected *big.Int", args[0])
	}

	reason, ok := args[1].([]byte)
	if !ok {
		return common.Hash{}, nil, nil, fmt.Errorf("v2 TxFailed reason is %T, expected []byte", args[1])
	}

	var reason_ error
	if reason__, err := abi.UnpackRevert(reason); err == nil {
		reason_ = fmt.Errorf("%v", reason__)
	} else {
		reason_ = fmt.Errorf("%v", hexutil.Encode(reason))
	}

	return subdigest, index, reason_, nil
}

func decodeTxExecutedV1(log *types.Log) (common.Hash, error) {
	if len(log.Topics) != 0 {
		return common.Hash{}, fmt.Errorf("%v topics, expected zero", len(log.Topics))
	}

	if len(log.Data) != common.HashLength {
		return common.Hash{}, fmt.Errorf("not v1 TxExecuted")
	}

	return common.BytesToHash(log.Data), nil
}

func decodeTxFailedV1(log *types.Log) (common.Hash, error, error) {
	if len(log.Topics) != 1 {
		return common.Hash{}, nil, fmt.Errorf("%v topics, expected one", len(log.Topics))
	}

	event := contracts.V1.WalletMainModule.ABI.Events["TxFailed"]

	if log.Topics[0] != event.ID {
		return common.Hash{}, nil, fmt.Errorf("not v1 TxFailed")
	}

	args, err := event.Inputs.NonIndexed().Unpack(log.Data)
	if err != nil {
		return common.Hash{}, nil, fmt.Errorf("unable to decode v1 TxFailed: %w", err)
	}
	if len(args) != 2 {
		return common.Hash{}, nil, fmt.Errorf("%v v1 TxFailed arguments, expected two", len(args))
	}

	subdigest, ok := args[0].([common.HashLength]byte)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("v1 TxFailed subdigest is %T, expected common.Hash", args[0])
	}

	reason, ok := args[1].([]byte)
	if !ok {
		return common.Hash{}, nil, fmt.Errorf("v1 TxFailed reason is %T, expected []byte", args[1])
	}

	var reason_ error
	if reason__, err := abi.UnpackRevert(reason); err == nil {
		reason_ = fmt.Errorf("%v", reason__)
	} else {
		reason_ = fmt.Errorf("%v", hexutil.Encode(reason))
	}

	return subdigest, reason_, nil
}
