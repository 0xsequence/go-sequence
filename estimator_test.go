package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/go-sequence"
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
