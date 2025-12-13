package receipts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/receipts"
	"github.com/stretchr/testify/require"
)

func TestFetchReceiptTokenTransfers(t *testing.T) {
	// Case 1: Simple ERC20 Transfer via EOA, which shows just a simple transfer event
	// https://arbiscan.io/tx/0x6753a97203d159702e594662729b06608cf3b9c99c0cce177b9d7b66e6456150
	t.Run("Case 1: Simple ERC20 Transfer via EOA", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x6753a97203d159702e594662729b06608cf3b9c99c0cce177b9d7b66e6456150")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)

		// spew.Dump(transfers)
		require.Equal(t, 1, len(transfers))
		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), transfers[0].Token) // USDC
		require.Equal(t, common.HexToAddress("0x1D17C0F90A0b3dFb5124C2FF56B33a0D2E202e1d"), transfers[0].From)
		require.Equal(t, common.HexToAddress("0x5646E2424A7b7d43740EF14bc5b4f1e00Bf9B6Ba"), transfers[0].To)
		require.Equal(t, big.NewInt(184840), transfers[0].Value)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs()
		require.NotNil(t, balances)
		require.Equal(t, 2, len(balances))
		// spew.Dump(balances)

		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), balances[0].Token) // USDC
		require.Equal(t, common.HexToAddress("0x1D17C0F90A0b3dFb5124C2FF56B33a0D2E202e1d"), balances[0].Account)
		require.Equal(t, big.NewInt(-184840), balances[0].Balance)

		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), balances[1].Token) // USDC
		require.Equal(t, common.HexToAddress("0x5646E2424A7b7d43740EF14bc5b4f1e00Bf9B6Ba"), balances[1].Account)
		require.Equal(t, big.NewInt(184840), balances[1].Balance)
	})

	// Case 2: a txn with a bunch of different erc20 transfers inside of it, ie. batch send
	// this is a sequence v2 send of usdc and magic on arbitrum.. batch send
	// https://arbiscan.io/tx/0x65c70290232207a21ef6805ae50622def8d52b7a8f381e1c3eac24d5423e0657
	t.Run("Case 2: Batch of ERC20 Transfers via SCW", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x65c70290232207a21ef6805ae50622def8d52b7a8f381e1c3eac24d5423e0657")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)

		// spew.Dump(transfers)
		require.Equal(t, 2, len(transfers))

		// USDC
		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), transfers[0].Token)
		require.Equal(t, common.HexToAddress("0x8e3E38fe7367dd3b52D1e281E4e8400447C8d8B9"), transfers[0].From)
		require.Equal(t, common.HexToAddress("0x9b1A542f3C455E8d6057C3478EB945B48D8e17fF"), transfers[0].To)
		require.Equal(t, big.NewInt(100000), transfers[0].Value)

		// MAGIC
		require.Equal(t, common.HexToAddress("0x539bdE0d7Dbd336b79148AA742883198BBF60342"), transfers[1].Token)
		require.Equal(t, common.HexToAddress("0x8e3E38fe7367dd3b52D1e281E4e8400447C8d8B9"), transfers[1].From)
		require.Equal(t, common.HexToAddress("0x9b1A542f3C455E8d6057C3478EB945B48D8e17fF"), transfers[1].To)
		require.Equal(t, big.NewInt(200000000000000000), transfers[1].Value)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs()
		require.NotNil(t, balances)
		require.Equal(t, 4, len(balances))
		// spew.Dump(balances)

		require.Equal(t, common.HexToAddress("0x539bdE0d7Dbd336b79148AA742883198BBF60342"), balances[0].Token) // MAGIC
		require.Equal(t, common.HexToAddress("0x8e3E38fe7367dd3b52D1e281E4e8400447C8d8B9"), balances[0].Account)
		require.Equal(t, big.NewInt(-200000000000000000), balances[0].Balance)

		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), balances[1].Token) // USDC
		require.Equal(t, common.HexToAddress("0x8e3E38fe7367dd3b52D1e281E4e8400447C8d8B9"), balances[1].Account)
		require.Equal(t, big.NewInt(-100000), balances[1].Balance)

		require.Equal(t, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"), balances[2].Token) // USDC
		require.Equal(t, common.HexToAddress("0x9b1A542f3C455E8d6057C3478EB945B48D8e17fF"), balances[2].Account)
		require.Equal(t, big.NewInt(100000), balances[2].Balance)

		require.Equal(t, common.HexToAddress("0x539bdE0d7Dbd336b79148AA742883198BBF60342"), balances[3].Token) // MAGIC
		require.Equal(t, common.HexToAddress("0x9b1A542f3C455E8d6057C3478EB945B48D8e17fF"), balances[3].Account)
		require.Equal(t, big.NewInt(200000000000000000), balances[3].Balance)
	})

	// Case 3: a trails intent call, with bunch of other actions inside of the txn, including erc20 transfers
	// https://arbiscan.io/tx/0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9
	t.Run("Case 3: Trails intent origin call", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0xb88cc2fea7cd26c88e169f6244fea76f590fc0797ba4c424669d1b74643f1dc9")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)
		require.Equal(t, 13, len(receipt.Logs))

		// Trails intent
		usdc := common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831")
		owner := common.HexToAddress("0x9DAB7A98C207f01A35DF00257949a27609b93ad7")
		trailsRouter := common.HexToAddress("0xF8A739B9F24E297a98b7aba7A9cdFDBD457F6fF8")
		bridge := common.HexToAddress("0xf70da97812CB96acDF810712Aa562db8dfA3dbEF")
		collector := common.HexToAddress("0x76008498f26789dd8b691Bebe24C889A3dd1A2fc")

		// spew.Dump(transfers)
		require.Equal(t, 3, len(transfers))

		// Log 1, USDC owner to trailsRouter
		require.Equal(t, usdc, transfers[0].Token)
		require.Equal(t, owner, transfers[0].From)
		require.Equal(t, trailsRouter, transfers[0].To)
		require.Equal(t, big.NewInt(175353), transfers[0].Value)

		// Log 2, USDC trailsRouter from bridge
		require.Equal(t, usdc, transfers[1].Token)
		require.Equal(t, trailsRouter, transfers[1].From)
		require.Equal(t, bridge, transfers[1].To)
		require.Equal(t, big.NewInt(175353), transfers[1].Value)

		// Log 3, USDC owner to collector
		require.Equal(t, usdc, transfers[2].Token)
		require.Equal(t, owner, transfers[2].From)
		require.Equal(t, collector, transfers[2].To)
		require.Equal(t, big.NewInt(9979), transfers[2].Value)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs().OmitZeroBalances()
		require.NotNil(t, balances)
		require.Equal(t, 3, len(balances))
		// spew.Dump(balances)

		require.Equal(t, usdc, balances[0].Token)
		require.Equal(t, owner, balances[0].Account)
		require.Equal(t, big.NewInt(-185332), balances[0].Balance)

		require.Equal(t, usdc, balances[1].Token)
		require.Equal(t, collector, balances[1].Account)
		require.Equal(t, big.NewInt(9979), balances[1].Balance)

		require.Equal(t, usdc, balances[2].Token)
		require.Equal(t, bridge, balances[2].Account)
		require.Equal(t, big.NewInt(175353), balances[2].Balance)
	})

	// Case 4: a trails cross-chain swap where we use 0x + cctp to swap from MAGIC to USDC then bridge
	// over CCTP. This includes many calls with USDC and MAGIC.
	// https://arbiscan.io/tx/0xa5c17e51443c8a8ce60cdcbe84b89fd2570f073bbb3b9ec8cdc9361aa1ca984f
	t.Run("Case 4: Trails swap intent call", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/arbitrum")
		require.NoError(t, err)

		txnHash := common.HexToHash("0xa5c17e51443c8a8ce60cdcbe84b89fd2570f073bbb3b9ec8cdc9361aa1ca984f")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)
		require.Equal(t, 31, len(receipt.Logs))

		// Trails intent
		require.Equal(t, 10, len(transfers))
		// spew.Dump(transfers)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs().OmitZeroBalances()
		require.NotNil(t, balances)
		require.Equal(t, 9, len(balances))
		// spew.Dump(balances)

		usdc := common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831")
		magic := common.HexToAddress("0x539bdE0d7Dbd336b79148AA742883198BBF60342")
		owner := common.HexToAddress("0x8f2951E6b9Bd9F8cf3522A7Fa98A0F9bC767b155")
		trailsRouter := common.HexToAddress("0xF8A739B9F24E297a98b7aba7A9cdFDBD457F6fF8")
		collector := common.HexToAddress("0x76008498f26789dd8b691Bebe24C889A3dd1A2fc")

		// intermediary token used via uniswap
		weth := common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
		uniswap := common.HexToAddress("0xB7E50106A5bd3Cf21AF210A755F9C8740890A8c9")

		// some rando token used in the swap
		liqBook := common.HexToAddress("0xb7236B927e03542AC3bE0A054F2bEa8868AF9508")

		// NOTE: these balances are not in order of operations
		// it is the outputs, and sorted by smallest to highest
		// in terms of numeric value (not USD value obviously)
		// as decimals are not factored in here per token.

		// owner sending magic
		require.Equal(t, magic, balances[0].Token)
		require.Equal(t, owner, balances[0].Account)
		require.Equal(t, makeBigInt(t, "-10686807000000000000"), balances[0].Balance)

		// uniswap sending out weth
		require.Equal(t, weth, balances[1].Token)
		require.Equal(t, uniswap, balances[1].Account)
		require.Equal(t, makeBigInt(t, "-383769729558864"), balances[1].Balance)

		// liqbook sending usdc
		require.Equal(t, usdc, balances[2].Token)
		require.Equal(t, liqBook, balances[2].Account)
		require.Equal(t, makeBigInt(t, "-1191946"), balances[2].Balance)

		// balance of some 0x related wallet, probably a fee collector for 0x
		require.Equal(t, usdc, balances[3].Token)
		require.Equal(t, common.HexToAddress("0xaD01C20d5886137e056775af56915de824c8fCe5"), balances[3].Account)
		require.Equal(t, makeBigInt(t, "1787"), balances[3].Balance)

		// trailsRouter receiving usdc
		// TODO: this must be a bug in trails router or calls, as there shouldn't be any
		// dust left in the router.
		require.Equal(t, usdc, balances[4].Token)
		require.Equal(t, trailsRouter, balances[4].Account)
		require.Equal(t, makeBigInt(t, "36299"), balances[4].Balance)

		// usdc burn for cctp
		require.Equal(t, usdc, balances[5].Token)
		require.Equal(t, common.HexToAddress("0x0000000000000000000000000000000000000000"), balances[5].Account)
		require.Equal(t, makeBigInt(t, "1153860"), balances[5].Balance)

		// liqBook got the weth from the swap flow
		require.Equal(t, weth, balances[6].Token)
		require.Equal(t, liqBook, balances[6].Account)
		require.Equal(t, makeBigInt(t, "383769729558864"), balances[6].Balance)

		// fee collector paid in magic
		require.Equal(t, magic, balances[7].Token)
		require.Equal(t, collector, balances[7].Account)
		require.Equal(t, makeBigInt(t, "109223258035414205"), balances[7].Balance)

		// uniswap ending up with the magic from swap in
		require.Equal(t, magic, balances[8].Token)
		require.Equal(t, uniswap, balances[8].Account)
		require.Equal(t, makeBigInt(t, "10577583741964585795"), balances[8].Balance)

		// Get balance of just the cctp burn address
		cctpBurnAddress := balances.FilterByAccount(common.HexToAddress("0x0000000000000000000000000000000000000000"), usdc)
		require.Equal(t, 1, len(cctpBurnAddress))
		require.Equal(t, usdc, cctpBurnAddress[0].Token)
		require.Equal(t, makeBigInt(t, "1153860"), cctpBurnAddress[0].Balance)

		// Get balance of uniswap only
		uniswapBalances := balances.FilterByAccount(uniswap)
		require.Equal(t, 2, len(uniswapBalances))
	})

	// Case 5: vault bridge USDC .. lets check the token transfer event, prob just erc20 too
	// https://katanascan.com/tx/0x7bcd0068a5c3352cf4e1d75c7c4f78d99f02b8b2f5f96b2c407972f43e724f52
	t.Run("Case 5: Vault bridge USDC transfer", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/katana")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x7bcd0068a5c3352cf4e1d75c7c4f78d99f02b8b2f5f96b2c407972f43e724f52")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)
		require.Equal(t, 1, len(receipt.Logs))

		require.Equal(t, 1, len(transfers))
		// spew.Dump(transfers)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs()
		require.NotNil(t, balances)
		require.Equal(t, 2, len(balances))
		// spew.Dump(balances)

		require.Equal(t, common.HexToAddress("0x203A662b0BD271A6ed5a60EdFbd04bFce608FD36"), balances[0].Token)
		require.Equal(t, common.HexToAddress("0x1D17C0F90A0b3dFb5124C2FF56B33a0D2E202e1d"), balances[0].Account)
		require.Equal(t, makeBigInt(t, "-177353"), balances[0].Balance)

		require.Equal(t, common.HexToAddress("0x203A662b0BD271A6ed5a60EdFbd04bFce608FD36"), balances[1].Token)
		require.Equal(t, common.HexToAddress("0xE766bc22e31097940FFF8A73B240055EFB2424C8"), balances[1].Account)
		require.Equal(t, makeBigInt(t, "177353"), balances[1].Balance)
	})

	// Case 6: polygon POL LogTransfer event
	// https://polygonscan.com/tx/0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5
	t.Run("Case 6: POL with LogTransfer", func(t *testing.T) {
		provider, err := ethrpc.NewProvider("https://nodes.sequence.app/polygon")
		require.NoError(t, err)

		txnHash := common.HexToHash("0x252419983224542bfb07dab75808fa57186a7a269d0d267ae655eb7ef037fdd5")
		receipt, transfers, err := receipts.FetchReceiptTokenTransfers(context.Background(), provider, txnHash)
		require.NoError(t, err)
		require.NotNil(t, receipt)
		require.Greater(t, len(transfers), 0)
		require.Equal(t, 25, len(receipt.Logs)) // actually a ton of stuff in here

		// Trails intent
		transfers = transfers.FilterByContractAddress(common.HexToAddress("0x0000000000000000000000000000000000001010")) // POL token
		require.Equal(t, 2, len(transfers))
		// spew.Dump(transfers)

		// Get the balance outputs from the transfer logs
		balances := transfers.ComputeBalanceOutputs(true) // omit zero balances
		require.NotNil(t, balances)
		require.Equal(t, 2, len(balances))
		// spew.Dump(balances)

		polPsuedoToken := common.HexToAddress("0x0000000000000000000000000000000000001010")

		require.Equal(t, polPsuedoToken, balances[0].Token)
		require.Equal(t, common.HexToAddress("0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"), balances[0].Account)
		require.Equal(t, makeBigInt(t, "-3965683724100320759"), balances[0].Balance)

		require.Equal(t, polPsuedoToken, balances[1].Token)
		require.Equal(t, common.HexToAddress("0x1D17C0F90A0b3dFb5124C2FF56B33a0D2E202e1d"), balances[1].Account)
		require.Equal(t, makeBigInt(t, "3965683724100320759"), balances[1].Balance)
	})

	// Case 7: bunch of logs for the same erc20 token, and we need to sum it up, ie. a big uniswap call
	// https://etherscan.io/tx/0xb11ff491495e145b07a1d3cc304f7d04b235b80af51b50da9a54095a6882fca4
	t.Run("Case 7: Random txn with swap and many tokens", func(t *testing.T) {
		// NOTE, skippig writing this for now as its pretty similar
		// to case 4
	})
}

func makeBigInt(t *testing.T, s string) *big.Int {
	bi, ok := new(big.Int).SetString(s, 10)
	require.True(t, ok)
	return bi
}
