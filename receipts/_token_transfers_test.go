package receipts_test

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/receipts"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

func TestFetchTransactionReceiptTokenTransfers(t *testing.T) {
	// txnHash := https://polygonscan.com/tx/0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5
	provider, err := ethrpc.NewProvider("https://nodes.sequence.app/polygon")
	require.NoError(t, err)

	txnHash := common.HexToHash("0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5")

	receipt, transfers, err := receipts.FetchTransactionReceiptTokenTransfers(context.Background(), provider, txnHash)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Greater(t, len(transfers), 0)

	spew.Dump(transfers)
}
