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

	// TODO: lets find a very simple metamask erc20 transfer, and check it in the test
	// https://arbiscan.io/tx/0x6753a97203d159702e594662729b06608cf3b9c99c0cce177b9d7b66e6456150

	// TODO2: a txn with a bunch of different erc20 transfers inside of it, ie. batch send
	// this is a sequence v2 send of usdc and magic on arbitrum.. batch send
	// https://arbiscan.io/tx/0x65c70290232207a21ef6805ae50622def8d52b7a8f381e1c3eac24d5423e0657

	// TODO3: a trails intent call, with bunch of other actions inside of the txn, including erc20 transfers
	// https://arbiscan.io/tx/0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9
	// .. lets get another one using zerox + cctp for example

	// TODO4: vault bridge USDC .. lets check the token transfer event, prob just erc20 too
	// https://katanascan.com/tx/0x7bcd0068a5c3352cf4e1d75c7c4f78d99f02b8b2f5f96b2c407972f43e724f52

	// TODO5: polygon POL LogTransfer event
	// https://polygonscan.com/tx/0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5

	// TODO6: bunch of logs for the same erc20 token, and we need to sum it up, ie. a big uniswap call
	// and we have to do a delta/diff, and return the "result" maybe "TokenTransferResult" ?
	// https://etherscan.io/tx/0xb11ff491495e145b07a1d3cc304f7d04b235b80af51b50da9a54095a6882fca4

	// Test 1: Simple ERC20 Transfer via EOA, which shows just a simple transfer event
	// https://arbiscan.io/tx/0x6753a97203d159702e594662729b06608cf3b9c99c0cce177b9d7b66e6456150
	t.Run("Simple ERC20 Transfer via EOA", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x6753a97203d159702e594662729b06608cf3b9c99c0cce177b9d7b66e6456150")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)

		spew.Dump(transfers)
	})

	//--

	// t.Run("Simple ERC20 Transfer", func(t *testing.T) {
	// 	// https://arbiscan.io/tx/0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9
	// 	provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
	// 	require.NoError(t, err)

	// 	txnHash := common.HexToHash("0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9")

	// 	receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
	// 	require.NoError(t, err)
	// 	require.NotNil(t, receipt)
	// 	require.Greater(t, len(transfers), 0)

	// 	spew.Dump(transfers)

	// })

	// t.Run("Polygon POL LogTransfer", func(t *testing.T) {
	// 	t.Skip("POL")

	// 	// txnHash := https://polygonscan.com/tx/0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5
	// 	provider, err := ethrpc.NewProvider("https://nodes.sequence.app/polygon")
	// 	require.NoError(t, err)

	// 	txnHash := common.HexToHash("0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5")

	// 	receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
	// 	require.NoError(t, err)
	// 	require.NotNil(t, receipt)
	// 	require.Greater(t, len(transfers), 0)

	// 	spew.Dump(transfers)
	// })
}
