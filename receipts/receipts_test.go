package receipts_test

import (
	"bytes"
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/receipts"
	"github.com/stretchr/testify/assert"
)

func TestReceiptsV1(t *testing.T) {
	test(t, Test{
		Network:     "polygon",
		Transaction: "0x14d32d6cc53c84422eb9346ce22e1476c6ebebf6cd2cc204c342c9d679791732",
		Receipts: []Receipt{
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 0,
				LogEnd:   7,
			},
		},
	})

	test(t, Test{
		Network:     "polygon",
		Transaction: "0xd0587e41a6441260c0f0871cc5ba6e69834f2a76c9566907695c7320c4616ae5",
		Receipts: []Receipt{
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 1,
				LogEnd:   2,
			},
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 3,
				LogEnd:   4,
			},
		},
	})
}

func TestReceiptsV2(t *testing.T) {
	test(t, Test{
		Network:     "polygon",
		Transaction: "0x3e955ef200f611bb57566e05988b2ee929dfda05fb75a9b4b669b219cf33ad98",
		Receipts: []Receipt{
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 1,
				LogEnd:   2,
			},
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 3,
				LogEnd:   4,
			},
		},
	})
}

func TestReceiptsV3(t *testing.T) {
	test(t, Test{
		Network:     "arbitrum",
		Transaction: "0xa4d715c6239dd7c9ee6e488db9d2c7cb908eb3dfe10d9d359f0e5b3359b895fd",
		Receipts: []Receipt{
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 0,
				LogEnd:   0,
			},
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 1,
				LogEnd:   7,
			},
		},
	})
}

func TestReceiptsZKSync(t *testing.T) {
	test(t, Test{
		Network:     "sandbox-testnet",
		Transaction: "0x6ddb43ddf29ed48f302c6b7b6dc9bb47b33dd5099540fb0793f239edb406fa5d",
		Receipts: []Receipt{
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 0,
				LogEnd:   3,
			},
			{
				Status:   receipts.StatusSucceeded,
				LogStart: 4,
				LogEnd:   8,
			},
		},
	})
}

type Test struct {
	Network     string
	Transaction string
	Receipts    []Receipt
}

type Receipt struct {
	Status           receipts.Status
	Error            string
	LogStart, LogEnd int
}

func test(t *testing.T, test Test) {
	ctx := context.Background()

	provider, err := ethrpc.NewProvider(fmt.Sprintf("https://nodes.sequence.app/%v", test.Network))
	assert.NoError(t, err)

	transaction := common.HexToHash(test.Transaction)

	receipt, err := provider.TransactionReceipt(ctx, transaction)
	assert.NoError(t, err)

	receipts, err := receipts.TransactionReceipts(ctx, transaction, provider)
	assert.NoError(t, err)
	assert.Len(t, receipts.Receipts, len(test.Receipts))

	for i, receipt_ := range receipts.Receipts {
		assert.NoError(t, isEqual(receipt.Logs, receipt_, test.Receipts[i]))
	}
}

func isEqual(logs []*types.Log, actual receipts.Receipt, expected Receipt) error {
	if expected.LogStart < 0 {
		panic(fmt.Sprintf("log start %v < 0", expected.LogStart))
	}
	if expected.LogEnd < expected.LogStart {
		panic(fmt.Sprintf("log end %v < log start %v", expected.LogEnd, expected.LogStart))
	}

	if actual.Status != expected.Status {
		return fmt.Errorf("receipt status %v, expected %v", actual.Status, expected.Status)
	}

	if expected.Error == "" {
		if actual.Error != nil {
			return fmt.Errorf("receipt error %v, expected nil", actual.Error)
		}
	} else {
		if actual.Error == nil {
			return fmt.Errorf("receipt error nil, expected %v", expected.Error)
		} else if actual.Error.Error() != expected.Error {
			return fmt.Errorf("receipt error %v, expected %v", actual.Error, expected.Error)
		}
	}

	if expected.LogEnd > len(logs) {
		return fmt.Errorf("%v transaction logs, expected at least %v", len(logs), expected.LogEnd)
	}
	if len(actual.Logs) != expected.LogEnd-expected.LogStart {
		return fmt.Errorf("%v receipt logs, expected %v", actual.Logs, expected.LogEnd-expected.LogStart)
	}

	for i, log := range actual.Logs {
		if err := isEqualLog(log, logs[expected.LogStart+i]); err != nil {
			return fmt.Errorf("log %v: %w", i, err)
		}
	}

	return nil
}

func isEqualLog(actual, expected *types.Log) error {
	if actual == expected {
		return nil
	}
	if actual.Address != expected.Address {
		return fmt.Errorf("log address %v, expected %v", actual.Address, expected.Address)
	}
	if !slices.Equal(actual.Topics, expected.Topics) {
		return fmt.Errorf("log topics %v, expected %v", actual.Topics, expected.Topics)
	}
	if !bytes.Equal(actual.Data, expected.Data) {
		return fmt.Errorf("log data %v, expected %v", hexutil.Encode(actual.Data), hexutil.Encode(expected.Data))
	}
	if actual.BlockNumber != expected.BlockNumber {
		return fmt.Errorf("log block number %v, expected %v", actual.BlockNumber, expected.BlockNumber)
	}
	if actual.TxHash != expected.TxHash {
		return fmt.Errorf("log transaction hash %v, expected %v", actual.TxHash, expected.TxHash)
	}
	if actual.TxIndex != expected.TxIndex {
		return fmt.Errorf("log transaction index %v, expected %v", actual.TxIndex, expected.TxIndex)
	}
	if actual.BlockHash != expected.BlockHash {
		return fmt.Errorf("log block hash %v, expected %v", actual.BlockHash, expected.BlockHash)
	}
	if actual.Index != expected.Index {
		return fmt.Errorf("log index %v, expected %v", actual.Index, expected.Index)
	}
	return nil
}
