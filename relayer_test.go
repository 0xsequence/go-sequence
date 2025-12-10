package sequence_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/receipts"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetReceiptOfTransaction(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V3DummySequenceWallet(1)
	require.NoError(t, err)
	require.NotNil(t, wallet)

	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	require.NoError(t, err)

	nonce, err := wallet.GetNonce()
	require.NoError(t, err)

	stx := &sequence.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	origReceipt, err := testutil.SignAndSendRawTransaction(t, wallet, stx)
	require.NoError(t, err)
	require.NotNil(t, origReceipt)
	t.Logf("original receipt txn hash: %v", origReceipt.TxHash.Hex())

	// Get transactions digest
	// Build V3 CallsPayload to compute the exact V3 metaTxnID
	// payload := v3.NewCallsPayload(
	// 	wallet.Address(),
	// 	testChain.ChainID(),
	// 	[]v3.Call{{
	// 		To:              callmockContract.Address,
	// 		Value:           big.NewInt(0),
	// 		Data:            calldata,
	// 		GasLimit:        big.NewInt(190000),
	// 		DelegateCall:    false,
	// 		OnlyFallback:    false,
	// 		BehaviorOnError: v3.BehaviorOnErrorIgnore,
	// 	}},
	// 	big.NewInt(0), // space
	// 	nonce,         // v3 nonce
	// )
	payload, err := stx.Bundle().Payload(wallet.Address(), testChain.ChainID(), big.NewInt(0), nonce)
	require.NoError(t, err)

	metaTxnID, _, err := sequence.ComputeMetaTxnIDFromCallsPayload(&payload)
	require.NoError(t, err)
	require.NotEmpty(t, metaTxnID)
	t.Logf("metaTxnID: 0x%s", string(metaTxnID))

	// NOTE: we can delay here by uncommenting, which will make it so we don't find the
	// receipt in the cache, and we depend on the QueryOnChain function to find it.
	time.Sleep(3 * time.Second)

	// Find receipt
	// status, receipt, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	require.NoError(t, err)
	require.NotNil(t, receipt)
	require.Equal(t, sequence.MetaTxnExecuted, result.Status)
	t.Logf("listener found receipt of txn hash: %v", receipt.TransactionHash().Hex())

	// Find receipt again via eth_getLogs method directly
	latestBlock := testChain.ReceiptsListener.LatestBlockNum()
	fromBlock := new(big.Int).Sub(latestBlock, big.NewInt(5000))
	if fromBlock.Cmp(big.NewInt(0)) < 0 {
		fromBlock = big.NewInt(0)
	}
	var toBlock *big.Int = nil

	metaTxnHash := common.HexToHash(string(metaTxnID))
	_, receipt2, err := receipts.FetchMetaTransactionReceiptByETHGetLogs(context.Background(), metaTxnHash, testChain.ReceiptsListener.RPCProvider(), fromBlock, toBlock)
	require.NoError(t, err)
	require.NotNil(t, receipt2)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt2.Status)
	require.Equal(t, receipt.Status(), receipt2.Status)
	require.Equal(t, receipt.TransactionHash(), receipt2.TxHash)
	t.Logf("fetch via eth_getLogs found receipt of txn hash: %v", receipt2.TxHash.Hex())
}

func TestGetReceiptOfErrorTransaction(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V1DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Turn on revert flag on callmock
	callmockContract, _ := testChain.Deploy(t, "WALLET_CALL_RECV_MOCK")
	calldata, err := callmockContract.Encode("setRevertFlag", true)
	assert.NoError(t, err)

	_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Call callmock, this should revert and fail the transaction
	calldata, err = callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &sequence.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	_, err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	// Get transactions digest
	metaTxnID, _, err := sequence.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, sequence.MetaTxnFailed, result.Status)
}

func TestGetReceiptOfFailedTransactionBetweenTransactions(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V1DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	isDeployed, err := wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, isDeployed)

	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	// Turn on revert flag on callmock
	callmockContract, _ := testChain.Deploy(t, "WALLET_CALL_RECV_MOCK")

	for i := 1; i <= 3; i++ {
		calldata, err := callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	calldata, err := callmockContract.Encode("setRevertFlag", true)
	assert.NoError(t, err)

	_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		stx := &sequence.Transaction{
			To:            callmockContract.Address,
			Data:          calldata,
			GasLimit:      big.NewInt(190000),
			RevertOnError: false,
		}
		_, err = testutil.SignAndSendRawTransaction(t, wallet, stx)
		assert.NoError(t, err)
	}

	calldata, err = callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &sequence.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	_, err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		stx := &sequence.Transaction{
			To:            callmockContract.Address,
			Data:          calldata,
			GasLimit:      big.NewInt(190000),
			RevertOnError: false,
		}
		_, err = testutil.SignAndSendRawTransaction(t, wallet, stx)
		assert.NoError(t, err)
	}

	// Get transactions digest
	metaTxnID, _, err := sequence.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, sequence.MetaTxnWalletExec)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status()) // native txn was successful
	assert.Equal(t, sequence.MetaTxnFailed, result.Status)           // meta txn failed
}

func TestGetReceiptOfTransactionBetweenTransactions(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V1DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	isDeployed, err := wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, isDeployed)

	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	for i := 1; i <= 3; i++ {
		calldata, err := callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &sequence.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	_, err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	// Get transactions digest
	metaTxnID, _, err := sequence.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, sequence.MetaTxnWalletExec)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := receipts.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, sequence.MetaTxnExecuted, result.Status)
}
