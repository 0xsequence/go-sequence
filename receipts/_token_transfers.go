package receipts

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
)

// FetchTransactionReceiptTokenTransfers fetches the transaction receipt for the given transaction hash
// and decodes any token transfer events (ERC20) that occurred within that transaction. TODOXXX: we
// currently only support ERC20 token transfers, but we can extend this to support ERC721 and ERC1155 as well.
func FetchTransactionReceiptTokenTransfers(ctx context.Context, provider *ethrpc.Provider, transactionHash common.Hash) (*types.Receipt, TokenTransfers, error) {
	receipt, err := provider.TransactionReceipt(ctx, transactionHash)
	if err != nil {
		return nil, nil, err
	}
	if receipt == nil {
		return nil, nil, nil
	}

	transferTopic := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	polLogTransferTopic := common.HexToHash("0xe6497e3ee548a3372136af2fcb0696db31fc6cf20260707645068bd3fe97f3c4")

	var decoded []*TokenTransfer

	for _, log := range receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}
		if log.Topics[0] != transferTopic && log.Topics[0] != polLogTransferTopic {
			continue
		}

		if log.Topics[0] == transferTopic {
			filterer, err := tokens.NewIERC20Filterer(log.Address, provider)
			if err == nil {
				if ev, err := filterer.ParseTransfer(*log); err == nil && ev != nil {
					decoded = append(decoded, &TokenTransfer{From: ev.From, To: ev.To, Value: ev.Value, Raw: *log})
					continue
				}
			}
		}

		// TODO: need to try all of the various versions of this.. and we may as well support ERC721 and ERC1155 too
		// note: "indexed" args, etc.

		if len(log.Topics) >= 3 {
			from := common.BytesToAddress(log.Topics[1].Bytes())
			to := common.BytesToAddress(log.Topics[2].Bytes())
			value := new(big.Int).SetBytes(log.Data)
			decoded = append(decoded, &TokenTransfer{From: from, To: to, Value: value, Raw: *log})
		}
	}

	return receipt, decoded, nil
}

type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

type TokenTransfers []*TokenTransfer

func (t TokenTransfers) FilterTokenTransfersByContractAddress(ctx context.Context, contract common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.Raw.Address == contract {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) FilterTokenTransfersByFromAddress(ctx context.Context, from common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.From == from {
			out = append(out, transfer)
		}
	}
	return out
}

func (t TokenTransfers) FilterTokenTransfersByToAddress(ctx context.Context, to common.Address) TokenTransfers {
	var out TokenTransfers
	for _, transfer := range t {
		if transfer.To == to {
			out = append(out, transfer)
		}
	}
	return out
}
