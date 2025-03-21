package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type TransactionStatus int

const (
	TransactionStatusSkipped TransactionStatus = iota
	TransactionStatusSucceeded
	TransactionStatusFailed
	TransactionStatusAborted
)

type Receipt struct {
	Status TransactionStatus
	Result []byte
	Logs   []*types.Log
}

func (r Receipt) Error() (string, error) {
	return abi.UnpackRevert(r.Result)
}

func (r Receipt) Receipts() (v3.PayloadDigest, common.Address, *big.Int, *big.Int, []Receipt, error) {
	return DecodeLogs(r.Logs)
}

func DecodeLogs(logs []*types.Log) (v3.PayloadDigest, common.Address, *big.Int, *big.Int, []Receipt, error) {
	var address common.Address
	var digest v3.PayloadDigest
	for i := len(logs) - 1; i >= 0; i-- {
		log := logs[i]

		var err error
		digest, _, err = DecodeCallSucceeded(log)
		if err == nil {
			address = log.Address
			break
		}

		digest, _, _, err = DecodeCallAborted(log)
		if err == nil {
			address = log.Address
			break
		}

		digest, _, _, err = DecodeCallFailed(log)
		if err == nil {
			address = log.Address
			break
		}

		digest, _, err = DecodeCallSkipped(log)
		if err == nil {
			address = log.Address
			break
		}
	}
	if digest.Hash == (common.Hash{}) {
		return v3.PayloadDigest{}, common.Address{}, nil, nil, nil, fmt.Errorf("no transaction logs")
	}

	var space, nonce *big.Int
	var receipts []Receipt
	var start int
	for i, log := range logs {
		space_, newNonce, err := DecodeNonceChange(log)
		if err == nil {
			if len(receipts) == 0 && space == nil && nonce == nil && log.Address == address {
				space, nonce = space_, new(big.Int).Sub(newNonce, common.Big1)
				start = i + 1
			}
			continue
		}

		digest_, index, err := DecodeCallSucceeded(log)
		if err == nil {
			if log.Address == address && digest_.Hash == digest.Hash {
				if index.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return v3.PayloadDigest{}, common.Address{}, nil, nil, nil, fmt.Errorf("logged index %v, expected %v", index, len(receipts))
				}
				receipts = append(receipts, Receipt{Status: TransactionStatusSucceeded, Logs: logs[start:i]})
				start = i + 1
			}
			continue
		}

		digest_, index, result, err := DecodeCallFailed(log)
		if err == nil {
			if log.Address == address && digest_.Hash == digest.Hash {
				if index.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return v3.PayloadDigest{}, common.Address{}, nil, nil, nil, fmt.Errorf("logged index %v, expected %v", index, len(receipts))
				}
				receipts = append(receipts, Receipt{Status: TransactionStatusFailed, Result: result, Logs: logs[start:i]})
				start = i + 1
			}
			continue
		}

		digest_, index, result, err = DecodeCallAborted(log)
		if err == nil {
			if log.Address == address && digest_.Hash == digest.Hash {
				if index.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return v3.PayloadDigest{}, common.Address{}, nil, nil, nil, fmt.Errorf("logged index %v, expected %v", index, len(receipts))
				}
				receipts = append(receipts, Receipt{Status: TransactionStatusAborted, Result: result, Logs: logs[start:i]})
				start = i + 1
			}
			continue
		}

		digest_, index, err = DecodeCallSkipped(log)
		if err == nil {
			if log.Address == address && digest_.Hash == digest.Hash {
				if index.Cmp(big.NewInt(int64(len(receipts)))) != 0 {
					return v3.PayloadDigest{}, common.Address{}, nil, nil, nil, fmt.Errorf("logged index %v, expected %v", index, len(receipts))
				}
				receipts = append(receipts, Receipt{Status: TransactionStatusSkipped, Logs: logs[start:i]})
				start = i + 1
			}
			continue
		}
	}

	return digest, address, space, nonce, receipts, nil
}

const (
	nonceChange   = "NonceChange(uint256 _space, uint256 _newNonce)"
	callSucceeded = "CallSucceeded(bytes32 _opHash, uint256 _index)"
	callFailed    = "CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)"
	callAborted   = "CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)"
	callSkipped   = "CallSkipped(bytes32 _opHash, uint256 _index)"
)

func DecodeNonceChange(log *types.Log) (*big.Int, *big.Int, error) {
	_, args, _, err := ethcoder.DecodeTransactionLogByEventSig(*log, nonceChange)
	if err != nil {
		return nil, nil, fmt.Errorf("not a NonceChange log: %w", err)
	}

	if len(args) != 2 {
		return nil, nil, fmt.Errorf("NonceChange with %v arguments, expected 2", len(args))
	}

	space, ok := args[0].(*big.Int)
	if !ok {
		return nil, nil, fmt.Errorf("NonceChange space is %T, expected *big.Int", args[0])
	}

	newNonce, ok := args[1].(*big.Int)
	if !ok {
		return nil, nil, fmt.Errorf("NonceChange new nonce is %T, expected *big.Int", args[1])
	}

	return space, newNonce, nil
}

func DecodeCallSucceeded(log *types.Log) (v3.PayloadDigest, *big.Int, error) {
	_, args, _, err := ethcoder.DecodeTransactionLogByEventSig(*log, callSucceeded)
	if err != nil {
		return v3.PayloadDigest{}, nil, fmt.Errorf("not a CallSucceeded log: %w", err)
	}

	if len(args) != 2 {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSucceeded with %v arguments, expected 2", len(args))
	}

	digest, ok := args[0].([32]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSucceeded digest is %T, expected [32]byte", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSucceeded index is %T, expected *big.Int", args[1])
	}

	return v3.PayloadDigest{Hash: digest}, index, nil
}

func DecodeCallFailed(log *types.Log) (v3.PayloadDigest, *big.Int, []byte, error) {
	_, args, _, err := ethcoder.DecodeTransactionLogByEventSig(*log, callFailed)
	if err != nil {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("not a CallFailed log: %w", err)
	}

	if len(args) != 3 {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallFailed with %v arguments, expected 3", len(args))
	}

	digest, ok := args[0].([32]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallFailed digest is %T, expected [32]byte", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallFailed index is %T, expected *big.Int", args[1])
	}

	result, ok := args[2].([]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallFailed result is %T, expected []byte", args[2])
	}

	return v3.PayloadDigest{Hash: digest}, index, result, nil
}

func DecodeCallAborted(log *types.Log) (v3.PayloadDigest, *big.Int, []byte, error) {
	_, args, _, err := ethcoder.DecodeTransactionLogByEventSig(*log, callAborted)
	if err != nil {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("not a CallAborted log: %w", err)
	}

	if len(args) != 3 {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallAborted with %v arguments, expected 3", len(args))
	}

	digest, ok := args[0].([32]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallAborted digest is %T, expected [32]byte", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallAborted index is %T, expected *big.Int", args[1])
	}

	result, ok := args[2].([]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, nil, fmt.Errorf("CallAborted result is %T, expected []byte", args[2])
	}

	return v3.PayloadDigest{Hash: digest}, index, result, nil
}

func DecodeCallSkipped(log *types.Log) (v3.PayloadDigest, *big.Int, error) {
	_, args, _, err := ethcoder.DecodeTransactionLogByEventSig(*log, callSkipped)
	if err != nil {
		return v3.PayloadDigest{}, nil, fmt.Errorf("not a CallSkipped log: %w", err)
	}

	if len(args) != 2 {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSkipped with %v arguments, expected 2", len(args))
	}

	digest, ok := args[0].([32]byte)
	if !ok {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSkipped digest is %T, expected [32]byte", args[0])
	}

	index, ok := args[1].(*big.Int)
	if !ok {
		return v3.PayloadDigest{}, nil, fmt.Errorf("CallSkipped index is %T, expected *big.Int", args[1])
	}

	return v3.PayloadDigest{Hash: digest}, index, nil
}
