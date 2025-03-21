package sequence

import (
	"context"

	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

func FetchReceipt(ctx context.Context, listener *ethreceipts.ReceiptsListener, payload v3.PayloadDigestable) (*ethreceipts.Receipt, MetaTxnStatus, string, ethreceipts.WaitReceiptFinalityFunc, error) {
	digest := payload.Digest()
	receipt, waitFinality, err := listener.FetchTransactionReceiptWithFilter(ctx, PayloadFilter(digest).LimitOne(true).SearchCache(true))
	if err != nil {
		return nil, 0, "", nil, err
	}

	for _, log := range receipt.Logs() {
		digest_, _, revert, err := DecodeCallFailed(log)
		if err != nil {
			digest_, _, revert, err = DecodeCallAborted(log)
		}
		if err == nil {
			if digest_.Hash == digest.Hash {
				reason, err := abi.UnpackRevert(revert)
				if err != nil {
					reason = hexutil.Encode(revert)
				}
				return receipt, MetaTxnFailed, reason, waitFinality, nil
			}
			continue
		}
	}

	return receipt, MetaTxnExecuted, "", waitFinality, nil
}

func PayloadFilter(payload v3.PayloadDigestable) ethreceipts.FilterQuery {
	return ethreceipts.FilterLogs(func(logs []*types.Log) bool {
		for _, log := range logs {
			_, _, err := DecodeCallSucceeded(log)
			if err == nil {
				return true
			}

			_, _, _, err = DecodeCallFailed(log)
			if err == nil {
				return true
			}

			_, _, _, err = DecodeCallAborted(log)
			if err == nil {
				return true
			}

			_, _, err = DecodeCallSkipped(log)
			if err == nil {
				return true
			}
		}

		return false
	})
}
