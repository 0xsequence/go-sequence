package receipts

import (
	"context"
	"math/big"
	"sort"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
)

// FetchReceiptTokenTransfers fetches the transaction receipt for the given transaction hash
// and decodes any token transfer events (ERC20) that occurred within that transaction. TODOXXX: we
// currently only support ERC20 token transfers, but we can extend this to support ERC721 and ERC1155 as well.
func FetchReceiptTokenTransfers(ctx context.Context, provider *ethrpc.Provider, transactionHash common.Hash) (*types.Receipt, TokenTransfers, error) {
	receipt, err := provider.TransactionReceipt(ctx, transactionHash)
	if err != nil {
		return nil, nil, err
	}
	if receipt == nil {
		return nil, nil, nil
	}
	transfers, err := DecodeTokenTransfersFromLogs(ctx, receipt.Logs)
	if err != nil {
		return receipt, nil, err
	}
	return receipt, transfers, nil
}

var tokenTransferTopicHashes = []common.Hash{
	common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"), // ERC20 Transfer
	common.HexToHash("0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4"), // Polygon POL LogTransfer (custom)
}

func DecodeTokenTransfersFromLogs(ctx context.Context, logs []*types.Log) (TokenTransfers, error) {
	transferTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	polLogTransferTopic := common.HexToHash("0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4")

	var decoded []*TokenTransfer

	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0] != transferTopic && log.Topics[0] != polLogTransferTopic {
			continue
		}

		tokenAddress := log.Address

		if log.Topics[0] == transferTopic {
			filterer, err := tokens.NewIERC20Filterer(log.Address, nil)
			if err == nil {
				if ev, err := filterer.ParseTransfer(*log); err == nil && ev != nil {
					decoded = append(decoded, &TokenTransfer{Token: tokenAddress, From: ev.From, To: ev.To, Value: ev.Value, Raw: *log})
					continue
				}
			}
		}

		// TODO: need to try all of the various versions of this.. and we may as well support ERC721 and ERC1155 too
		// note: "indexed" args, etc.

		// TODO: this is wrong, etc.
		if len(log.Topics) >= 3 {
			from := common.BytesToAddress(log.Topics[1].Bytes())
			to := common.BytesToAddress(log.Topics[2].Bytes())
			value := new(big.Int).SetBytes(log.Data)
			decoded = append(decoded, &TokenTransfer{Token: tokenAddress, From: from, To: to, Value: value, Raw: *log})
		}
	}

	return decoded, nil
}

type TokenTransfer struct {
	Token common.Address
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

type TokenTransfers []*TokenTransfer

type TokenBalance struct {
	Token   common.Address
	Account common.Address
	Balance *big.Int
}

type TokenBalances []*TokenBalance

func (t TokenTransfers) FilterByContractAddress(ctx context.Context, contract common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.Raw.Address == contract {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) FilterByAccountAddress(ctx context.Context, account common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.From == account || transfer.To == account {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) FilterByFromAddress(ctx context.Context, from common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.From == from {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) FilterByToAddress(ctx context.Context, to common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.To == to {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) Delta() TokenTransfers {
	out := TokenTransfers{}
	return out
}

// ComputeBalanceOutputs aggregates net balance changes per token per account from the transfers.
// For each transfer, it subtracts `Value` from `From` and adds `Value` to `To`.
// Accounts with a resulting zero balance change for a given token are omitted.
func (s TokenTransfers) ComputeBalanceOutputs() TokenBalances {
	// key: token address + account address
	type key struct {
		token   common.Address
		account common.Address
	}

	balances := make(map[key]*big.Int)

	for _, tr := range s {
		if tr == nil || tr.Value == nil {
			continue
		}

		// From: subtract value
		kFrom := key{token: tr.Token, account: tr.From}
		if _, ok := balances[kFrom]; !ok {
			balances[kFrom] = new(big.Int)
		}
		balances[kFrom].Sub(balances[kFrom], tr.Value)

		// To: add value
		kTo := key{token: tr.Token, account: tr.To}
		if _, ok := balances[kTo]; !ok {
			balances[kTo] = new(big.Int)
		}
		balances[kTo].Add(balances[kTo], tr.Value)
	}

	// Convert to slice, excluding zero balances
	out := TokenBalances{}
	zero := big.NewInt(0)

	for k, v := range balances {
		if v == nil || v.Cmp(zero) == 0 {
			continue
		}
		out = append(out, &TokenBalance{
			Token:   k.token,
			Account: k.account,
			Balance: new(big.Int).Set(v),
		})
	}

	sort.Slice(out, func(i, j int) bool {
		bi := out[i].Balance
		bj := out[j].Balance
		// ascending by numeric value (negative first)
		cmp := bi.Cmp(bj)
		if cmp != 0 {
			return cmp < 0
		}
		// account lexicographic
		ai := out[i].Account.Hex()
		aj := out[j].Account.Hex()
		if ai != aj {
			return ai < aj
		}
		// token lexicographic
		ti := out[i].Token.Hex()
		tj := out[j].Token.Hex()
		return ti < tj
	})

	return out
}
