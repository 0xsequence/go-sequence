package v1v2_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/core/v1v2"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestGetReceiptOfTransaction(t *testing.T) {
	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.V1DummySequenceWallet(1)
	assert.NoError(t, err)
	assert.NotNil(t, wallet)

	// Create normal txn of: callmockContract.testCall(55, 0x112255)
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
	calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &v1v2.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	// Get transactions digest
	metaTxnID, _, err := v1v2.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := v1v2.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := v1v2.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, v1v2.MetaTxnExecuted, result.Status)
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

	err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Call callmock, this should revert and fail the transaction
	calldata, err = callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &v1v2.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	// Get transactions digest
	metaTxnID, _, err := v1v2.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := v1v2.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := v1v2.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, v1v2.MetaTxnFailed, result.Status)
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
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	calldata, err := callmockContract.Encode("setRevertFlag", true)
	assert.NoError(t, err)

	err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		stx := &v1v2.Transaction{
			To:            callmockContract.Address,
			Data:          calldata,
			GasLimit:      big.NewInt(190000),
			RevertOnError: false,
		}
		err = testutil.SignAndSendRawTransaction(t, wallet, stx)
		assert.NoError(t, err)
	}

	calldata, err = callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &v1v2.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		stx := &v1v2.Transaction{
			To:            callmockContract.Address,
			Data:          calldata,
			GasLimit:      big.NewInt(190000),
			RevertOnError: false,
		}
		err = testutil.SignAndSendRawTransaction(t, wallet, stx)
		assert.NoError(t, err)
	}

	// Get transactions digest
	metaTxnID, _, err := v1v2.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, v1v2.MetaTxnWalletExec)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := v1v2.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := v1v2.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status()) // native txn was successful
	assert.Equal(t, v1v2.MetaTxnFailed, result.Status)               // meta txn failed
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
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
	assert.NoError(t, err)

	nonce, err := wallet.GetNonce()
	assert.NoError(t, err)

	stx := &v1v2.Transaction{
		To:            callmockContract.Address,
		Data:          calldata,
		Value:         big.NewInt(0),
		GasLimit:      big.NewInt(190000),
		DelegateCall:  false,
		RevertOnError: false,
		Nonce:         nonce,
	}

	err = testutil.SignAndSendRawTransaction(t, wallet, stx)
	assert.NoError(t, err)

	for i := 1; i <= 3; i++ {
		calldata, err = callmockContract.Encode("testCall", big.NewInt(int64(i)), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)
	}

	// Get transactions digest
	metaTxnID, _, err := v1v2.ComputeMetaTxnID(testChain.ChainID(), wallet.Address(), stx.Bundle(), nonce, v1v2.MetaTxnWalletExec)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	// Find receipt
	// status, receipt, err := v1v2.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnId)
	result, receipt, _, err := v1v2.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, v1v2.MetaTxnExecuted, result.Status)
}
