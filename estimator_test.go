package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestEstimateSimpleTransaction(t *testing.T) {
	provider := testChain.Provider
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	calldata, err := callmockContract.Encode("testCall", big.NewInt(11), ethcoder.MustHexDecode("0x99884422"))
	assert.NoError(t, err)

	gas, err := sequence.NewEstimator().EstimateCall(context.Background(), provider, &sequence.EstimateTransaction{
		To:   callmockContract.Address,
		Data: calldata,
	}, nil, "")

	assert.NoError(t, err)
	assert.NotZero(t, gas.Int64())

	wallet := testChain.MustWallet(2)
	tx, err := wallet.NewTransaction(context.Background(), &ethtxn.TransactionRequest{To: &callmockContract.Address, Data: calldata})
	assert.NoError(t, err)

	stx, err := wallet.SignTx(tx, testChain.ChainID())
	assert.NoError(t, err)

	_, waitReceipt, err := wallet.SendTransaction(context.Background(), stx)

	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)

	assert.Less(t, receipt.GasUsed, gas.Uint64())
	assert.Less(t, gas.Uint64()-receipt.GasUsed, uint64(5000))
}

func TestEstimateSimpleSequenceTransaction(t *testing.T) {
	wallet, err := testChain.DummySequenceWallet(1)
	assert.NoError(t, err)

	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
	assert.NoError(t, err)
	testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

	calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
	assert.NoError(t, err)

	txs := sequence.Transactions{
		&sequence.Transaction{
			To:    callmockContract.Address,
			Data:  calldata,
			Nonce: testChain.RandomNonce(),
		},
	}

	estimator := sequence.NewEstimator()
	estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

	assert.NoError(t, err)
	assert.NotZero(t, estimated)

	signed, err := wallet.SignTransactions(context.Background(), txs)
	assert.NoError(t, err)

	_, _, wait, err := wallet.SendTransactions(context.Background(), signed)
	assert.NoError(t, err)

	receipt, err := wait(context.Background())
	assert.NoError(t, err)

	assert.LessOrEqual(t, receipt.GasUsed, estimated)
	assert.Less(t, estimated-receipt.GasUsed, uint64(25000))

	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "3", ret[0])
}
