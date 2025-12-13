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

func TestFetchReceiptTokenTransfers(t *testing.T) {

	t.Run("ERC20 Transfer on Arbitrum", func(t *testing.T) {
		// https://arbiscan.io/tx/0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9")

		// TODO: lets find a very simple metamask erc20 transfer, and check it in the test
		// TODO2: find a batch of different erc20 transfers to test against.. for example
		// its possible there can be multiple erc20 transfers in a single tx that come and go from an individual
		// so what we want to see is the delta, the diff, etc. .. aka, the "Result" .. aka... "TokenTransferResult"
		// and not just all of the TokenTransferLogs ..
		// TODO3: vault bridge USDC .. lets check the token transfer event, prob just erc20 too

		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)

		spew.Dump(transfers)

	})

	t.Run("Polygon POL LogTransfer", func(t *testing.T) {
		t.Skip("POL")

		// txnHash := https://polygonscan.com/tx/0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/polygon")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5")

		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)

		spew.Dump(transfers)
	})
}
