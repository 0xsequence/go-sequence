package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	// Sign and send the transaction
	err = signAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Check the value
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "55", ret[0])
}

func TestERC20Transfer(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Create normal txn of: callmockContract.mockMint(wallet, 100)
	callmockContract := testChain.UniDeploy(t, "ERC20Mock", 0)
	calldata, err := callmockContract.Encode("mockMint", wallet.Address(), big.NewInt(100))
	assert.NoError(t, err)

	// Sign and send the transaction
	err = signAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Check the value
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallet.Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "100", ret[0])
}

func signAndSend(t *testing.T, wallet *sequence.Wallet, to common.Address, data []byte) error {
	stx := &sequence.Transaction{
		// DelegateCall:  false,
		// RevertOnError: false,
		// GasLimit: big.NewInt(800000),
		// Value:         big.NewInt(0),
		To:   to,
		Data: data,
	}

	// Now, we must sign the meta txn
	signedTx, err := wallet.SignTransaction(context.Background(), stx)
	assert.NoError(t, err)

	metaTxnID, tx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)
	assert.NotNil(t, tx)

	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)
	assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

	// TODO: decode the receipt, and lets confirm we have the metaTxnID event in there..
	// NOTE: if you start test chain with `make start-test-chain-verbose`, you will see the metaTxnID above
	// correctly logged..

	return err
}
