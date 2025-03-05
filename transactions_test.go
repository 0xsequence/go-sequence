package sequence_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "55", ret[0])
	})

	t.Run("v2", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "55", ret[0])
	})

	t.Run("v3", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(55), ethcoder.MustHexDecode("0x112255"))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "55", ret[0])
	})
}

func TestTransactionVerbose(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// This test is similar to TestTransaction but includes a few extra assertions

		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		stx := &sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}

		// Sign and send the transaction
		signedTx, err := wallet.SignTransaction(context.Background(), stx)
		assert.NoError(t, err)

		assert.NotEmpty(t, signedTx.Digest)
		assert.NotEmpty(t, signedTx.Signature)

		// Recover walletconfig + address from the signed transaction digest + signature
		//txSubDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), signedTx.Digest)
		//assert.NoError(t, err)

		walletConfig, weight, err := sequence.GenericRecoverWalletConfigFromDigest[*v1.WalletConfig](signedTx.Digest.Bytes(), signedTx.Signature, wallet.Address(), testutil.V1SequenceContext(), testChain.ChainID(), testChain.Provider)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, weight.Cmp(big.NewInt(int64(walletConfig.Threshold()))), 0)

		walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, testutil.V1SequenceContext())
		assert.NoError(t, err)

		assert.Equal(t, wallet.Address(), walletAddress)

		expectedMetaTxnID, _, err := sequence.ComputeMetaTxnIDFromDigest(testChain.ChainID(), walletAddress, signedTx.Digest)
		assert.NoError(t, err)

		// Send the transaction
		metaTxnID, tx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)
		assert.NotEmpty(t, metaTxnID)
		assert.NotNil(t, tx)
		assert.Equal(t, expectedMetaTxnID, metaTxnID)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "65", ret[0])
	})

	t.Run("v2", func(t *testing.T) {
		// This test is similar to TestTransaction but includes a few extra assertions

		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		stx := &sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}

		// Sign and send the transaction
		signedTx, err := wallet.SignTransaction(context.Background(), stx)
		assert.NoError(t, err)

		assert.NotEmpty(t, signedTx.Digest)
		assert.NotEmpty(t, signedTx.Signature)

		// Recover walletconfig + address from the signed transaction digest + signature
		//txSubDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), signedTx.Digest)
		//assert.NoError(t, err)

		walletConfig, weight, err := sequence.GenericRecoverWalletConfigFromDigest[*v2.WalletConfig](signedTx.Digest.Bytes(), signedTx.Signature, wallet.Address(), testutil.V1SequenceContext(), testChain.ChainID(), testChain.Provider)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, weight.Cmp(big.NewInt(int64(walletConfig.Threshold()))), 0)

		walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, testutil.V2SequenceContext())
		assert.NoError(t, err)

		assert.Equal(t, wallet.Address(), walletAddress)

		expectedMetaTxnID, _, err := sequence.ComputeMetaTxnIDFromDigest(testChain.ChainID(), walletAddress, signedTx.Digest)
		assert.NoError(t, err)

		// Send the transaction
		metaTxnID, tx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)
		assert.NotEmpty(t, metaTxnID)
		assert.NotNil(t, tx)
		assert.Equal(t, expectedMetaTxnID, metaTxnID)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "65", ret[0])
	})

	t.Run("v3", func(t *testing.T) {
		// This test is similar to TestTransaction but includes a few extra assertions

		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		stx := &sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}

		// Sign and send the transaction
		signedTx, err := wallet.SignTransaction(context.Background(), stx)
		assert.NoError(t, err)

		assert.NotEmpty(t, signedTx.Digest)
		assert.NotEmpty(t, signedTx.Signature)

		// Recover walletconfig + address from the signed transaction digest + signature
		//txSubDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), signedTx.Digest)
		//assert.NoError(t, err)

		walletConfig, weight, err := sequence.GenericRecoverWalletConfigFromDigest[*v3.WalletConfig](signedTx.Digest.Bytes(), signedTx.Signature, wallet.Address(), testutil.V3SequenceContext(), testChain.ChainID(), testChain.Provider)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, weight.Cmp(big.NewInt(int64(walletConfig.Threshold()))), 0)

		walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, testutil.V3SequenceContext())
		assert.NoError(t, err)

		assert.Equal(t, wallet.Address(), walletAddress)

		expectedMetaTxnID, _, err := sequence.ComputeMetaTxnIDFromDigest(testChain.ChainID(), walletAddress, signedTx.Digest)
		assert.NoError(t, err)

		// Send the transaction
		metaTxnID, tx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)
		assert.NotEmpty(t, metaTxnID)
		assert.NotNil(t, tx)
		assert.Equal(t, expectedMetaTxnID, metaTxnID)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "65", ret[0])
	})
}

func TestTransactionBundling(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// Ensure three dummy sequence wallets are deployed
		wallets, err := testChain.V1DummySequenceWallets(3, 1)
		assert.NoError(t, err)
		assert.NotNil(t, wallets)

		// Create normal txn of: callmockContract.mockMint(wallets[0], 100)
		callmockContract, _ := testChain.Deploy(t, "ERC20Mock")
		calldata, err := callmockContract.Encode("mockMint", wallets[0].Address(), big.NewInt(100))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallets[0], callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "100", ret[0])

		// Create normal txn of: callmockContract.transfer(wallets[1], 10)
		calldata1, err := callmockContract.Encode("transfer", wallets[1].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.transfer(wallets[2], 10)
		calldata2, err := callmockContract.Encode("transfer", wallets[2].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// These are wallet[0]'s intended transactions
		bundle := sequence.Transactions{
			{
				To:   callmockContract.Address,
				Data: calldata1,
			}, // balances: 90, 10, 0
			{
				To:   callmockContract.Address,
				Data: calldata2,
			}, // balances: 80, 10, 10
			{
				To: wallets[0].Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata1,
					}, // balances: 70, 20, 10
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 60, 20, 20
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 50, 20, 30
				},
			},
		}

		// wallet[0] must sign its bundle
		signedBundle, err := wallets[0].SignTransactions(context.Background(), bundle)
		assert.NoError(t, err)

		_, _, waitReceipt, err := wallets[0].SendTransaction(context.Background(), signedBundle)
		assert.NoError(t, err)

		waitReceipt(context.Background())

		// Check the value of wallet 1
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "50", ret[0])

		// Check the value of wallet 2
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[1].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "20", ret[0])

		// Check the value of wallet 3
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[2].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "30", ret[0])
	})

	t.Run("v2", func(t *testing.T) {
		// Ensure three dummy sequence wallets are deployed
		wallets, err := testChain.V2DummySequenceWallets(3, 1)
		assert.NoError(t, err)
		assert.NotNil(t, wallets)

		// Create normal txn of: callmockContract.mockMint(wallets[0], 100)
		callmockContract, _ := testChain.Deploy(t, "ERC20Mock")
		calldata, err := callmockContract.Encode("mockMint", wallets[0].Address(), big.NewInt(100))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallets[0], callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "100", ret[0])

		// Create normal txn of: callmockContract.transfer(wallets[1], 10)
		calldata1, err := callmockContract.Encode("transfer", wallets[1].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.transfer(wallets[2], 10)
		calldata2, err := callmockContract.Encode("transfer", wallets[2].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// These are wallet[0]'s intended transactions
		bundle := sequence.Transactions{
			{
				To:   callmockContract.Address,
				Data: calldata1,
			}, // balances: 90, 10, 0
			{
				To:   callmockContract.Address,
				Data: calldata2,
			}, // balances: 80, 10, 10
			{
				To: wallets[0].Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata1,
					}, // balances: 70, 20, 10
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 60, 20, 20
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 50, 20, 30
				},
			},
		}

		// wallet[0] must sign its bundle
		signedBundle, err := wallets[0].SignTransactions(context.Background(), bundle)
		assert.NoError(t, err)

		_, _, waitReceipt, err := wallets[0].SendTransaction(context.Background(), signedBundle)
		assert.NoError(t, err)

		waitReceipt(context.Background())

		// Check the value of wallet 1
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "50", ret[0])

		// Check the value of wallet 2
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[1].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "20", ret[0])

		// Check the value of wallet 3
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[2].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "30", ret[0])
	})

	t.Run("v3", func(t *testing.T) {
		// Ensure three dummy sequence wallets are deployed
		wallets, err := testChain.V3DummySequenceWallets(3, 1)
		assert.NoError(t, err)
		assert.NotNil(t, wallets)

		// Create normal txn of: callmockContract.mockMint(wallets[0], 100)
		callmockContract, _ := testChain.Deploy(t, "ERC20Mock")
		calldata, err := callmockContract.Encode("mockMint", wallets[0].Address(), big.NewInt(100))
		assert.NoError(t, err)

		// Sign and send the transaction
		err = testutil.SignAndSend(t, wallets[0], callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "100", ret[0])

		// Create normal txn of: callmockContract.transfer(wallets[1], 10)
		calldata1, err := callmockContract.Encode("transfer", wallets[1].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.transfer(wallets[2], 10)
		calldata2, err := callmockContract.Encode("transfer", wallets[2].Address(), big.NewInt(10))
		assert.NoError(t, err)

		// These are wallet[0]'s intended transactions
		bundle := sequence.Transactions{
			{
				To:   callmockContract.Address,
				Data: calldata1,
			}, // balances: 90, 10, 0
			{
				To:   callmockContract.Address,
				Data: calldata2,
			}, // balances: 80, 10, 10
			{
				To: wallets[0].Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata1,
					}, // balances: 70, 20, 10
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 60, 20, 20
					{
						To:   callmockContract.Address,
						Data: calldata2,
					}, // balances: 50, 20, 30
				},
			},
		}

		// wallet[0] must sign its bundle
		signedBundle, err := wallets[0].SignTransactions(context.Background(), bundle)
		assert.NoError(t, err)

		_, _, waitReceipt, err := wallets[0].SendTransaction(context.Background(), signedBundle)
		assert.NoError(t, err)

		waitReceipt(context.Background())

		// Check the value of wallet 1
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "50", ret[0])

		// Check the value of wallet 2
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[1].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "20", ret[0])

		// Check the value of wallet 3
		ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[2].Address().Hex()})
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "30", ret[0])
	})
}

func TestTransactionToGuestModuleBasic(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(1239), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		txns := sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}
		bundle := txns.Bundle()

		encodedTxns, err := bundle.EncodedTransactions()
		assert.NoError(t, err)

		execdata, err := contracts.V1.WalletGuestModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
		assert.NoError(t, err)

		metaTxnID, _, err := sequence.ComputeMetaTxnID(
			testChain.ChainID(),
			testChain.V1SequenceContext().GuestModuleAddress,
			bundle, nil, sequence.MetaTxnGuestExec,
		)
		assert.NoError(t, err)

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
		guestAddress := testChain.V1SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To: &guestAddress, Data: execdata,
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)

		// fmt.Println("==> metaTxnID", metaTxnID)
		// fmt.Println("==> txnHash", receipt.TxHash.Hex())

		// NOTE: we cannot use sequence.WaitForMetaTxn and listen for the metaTxnID, as there is no NonceChange
		// event in the above transaction as its just calling the mock receiver

		// Assert that first log in the receipt computes to the guest subdigest / id
		assert.True(t, fmt.Sprintf("0x%s", metaTxnID) == common.BytesToHash(receipt.Logs[0].Data).Hex())

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "1239", ret[0])
	})

	t.Run("v2", func(t *testing.T) {
		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(1239), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		txns := sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}
		bundle := txns.Bundle()

		encodedTxns, err := bundle.EncodedTransactions()
		assert.NoError(t, err)

		execdata, err := contracts.V2.WalletGuestModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
		assert.NoError(t, err)

		metaTxnID, _, err := sequence.ComputeMetaTxnID(
			testChain.ChainID(),
			testChain.V2SequenceContext().GuestModuleAddress,
			bundle, nil, sequence.MetaTxnGuestExec,
		)
		assert.NoError(t, err)

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
		guestAddress := testChain.V2SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To: &guestAddress, Data: execdata,
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)

		// fmt.Println("==> metaTxnID", metaTxnID)
		// fmt.Println("==> txnHash", receipt.TxHash.Hex())

		// NOTE: we cannot use sequence.WaitForMetaTxn and listen for the metaTxnID, as there is no NonceChange
		// event in the above transaction as its just calling the mock receiver

		// Assert that first log in the receipt computes to the guest subdigest / id
		assert.True(t, fmt.Sprintf("0x%s", metaTxnID) == receipt.Logs[0].Topics[1].Hex())

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "1239", ret[0])
	})

	t.Run("v3", func(t *testing.T) {
		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(1239), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		txns := sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}
		bundle := txns.Bundle()

		payload := sequence.ConvertTransactionsToV3Payload(bundle, big.NewInt(0), big.NewInt(0))
		encoded, err := v3.Encode(payload, nil)
		assert.NoError(t, err)

		execdata, err := contracts.V3.WalletStage1Module.Encode("execute", encoded, []byte{})
		assert.NoError(t, err)

		sender := testChain.GetRelayerWallet()

		metaTxnBytes, err := v3.HashPayload(sender.Address(), testChain.ChainID(), payload)
		assert.NoError(t, err)

		metaTxnHash := common.BytesToHash(metaTxnBytes[:])
		metaTxnID := sequence.MetaTxnID(metaTxnHash.Hex()[2:])

		// Relay the txn manually, directly to the guest module
		guestAddress := testChain.V3SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To:       &guestAddress,
			Data:     execdata,
			GasLimit: 1000000, // TODO: compute gas limit
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Assert that first log in the receipt computes to the guest subdigest / id
		assert.True(t, strings.Contains(common.BytesToHash(receipt.Logs[0].Data).Hex(), fmt.Sprintf("0x%s", metaTxnID)))

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "1239", ret[0])

	})
}

func TestTransactionToGuestModuleDeployAndCall(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V1DummySequenceWallet(testutil.RandomSeed(), true)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Assert the wallet is undeployed -- this is desired so we relayer the txn to the guest module
		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		if isDeployed {
			t.Fatal("expecting wallet to be undeployed")
		}

		// Wallet deployment data
		_, walletFactoryAddress, walletDeployData, err := sequence.EncodeWalletDeployment(wallet.GetWalletConfig(), wallet.GetWalletContext())
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		// Bundle of transactions: 1.) deploy new wallet 2.) send txn to the wallet
		// Then we send it to the GuestModule

		walletBundle := sequence.Transactions{
			{
				To: wallet.Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata,
					},
				},
			},
		}

		signedWalletBundle, err := wallet.SignTransactions(context.Background(), walletBundle)
		assert.NoError(t, err)

		signedWalletData, err := signedWalletBundle.Execdata()
		assert.NoError(t, err)

		guestBundle := sequence.Transactions{
			{
				To:   walletFactoryAddress,
				Data: walletDeployData,
			},
			{
				To:   wallet.Address(),
				Data: signedWalletData,
			},
		}

		encodedTxns, err := guestBundle.EncodedTransactions()
		assert.NoError(t, err)

		execdata, err := contracts.V1.WalletGuestModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
		assert.NoError(t, err)

		metaTxnID, _, err := sequence.ComputeMetaTxnID(
			testChain.ChainID(),
			testChain.V1SequenceContext().GuestModuleAddress,
			guestBundle, nil, sequence.MetaTxnGuestExec,
		)
		assert.NoError(t, err)

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
		guestAddress := testChain.V1SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To:       &guestAddress,
			Data:     execdata,
			GasLimit: 1000000, // TODO: compute gas limit
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// spew.Dump(sequence.DecodeRevertReason(receipt.Logs))

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "2255", ret[0])

		// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
		result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
		// metaStatus, _, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnID)
		assert.NoError(t, err)
		assert.True(t, result.Status == sequence.MetaTxnExecuted)

		// Wallet should be deployed now
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})

	t.Run("v2", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V2DummySequenceWallet(testutil.RandomSeed(), true)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Assert the wallet is undeployed -- this is desired so we relayer the txn to the guest module
		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		if isDeployed {
			t.Fatal("expecting wallet to be undeployed")
		}

		// Wallet deployment data
		_, walletFactoryAddress, walletDeployData, err := sequence.EncodeWalletDeployment(wallet.GetWalletConfig(), wallet.GetWalletContext())
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		// Bundle of transactions: 1.) deploy new wallet 2.) send txn to the wallet
		// Then we send it to the GuestModule

		walletBundle := sequence.Transactions{
			{
				To: wallet.Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata,
					},
				},
			},
		}

		signedWalletBundle, err := wallet.SignTransactions(context.Background(), walletBundle)
		assert.NoError(t, err)

		signedWalletData, err := signedWalletBundle.Execdata()
		assert.NoError(t, err)

		guestBundle := sequence.Transactions{
			{
				To:   walletFactoryAddress,
				Data: walletDeployData,
			},
			{
				To:   wallet.Address(),
				Data: signedWalletData,
			},
		}

		encodedTxns, err := guestBundle.EncodedTransactions()
		assert.NoError(t, err)

		execdata, err := contracts.V2.WalletGuestModule.Encode("execute", encodedTxns, big.NewInt(0), []byte{})
		assert.NoError(t, err)

		metaTxnID, _, err := sequence.ComputeMetaTxnID(
			testChain.ChainID(),
			testChain.V2SequenceContext().GuestModuleAddress,
			guestBundle, nil, sequence.MetaTxnGuestExec,
		)
		assert.NoError(t, err)

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
		guestAddress := testChain.V2SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To:       &guestAddress,
			Data:     execdata,
			GasLimit: 1000000, // TODO: compute gas limit
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// spew.Dump(sequence.DecodeRevertReason(receipt.Logs))

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "2255", ret[0])

		// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
		result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
		// metaStatus, _, err := sequence.WaitForMetaTxn(context.Background(), testChain.Provider, metaTxnID)
		assert.NoError(t, err)
		assert.True(t, result.Status == sequence.MetaTxnExecuted)

		// Wallet should be deployed now
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})

	t.Run("v3", func(t *testing.T) {
		// Ensure dummy sequence wallet from seed 1 is deployed
		wallet, err := testChain.V3DummySequenceWallet(testutil.RandomSeed(), true)
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Assert the wallet is undeployed -- this is desired so we relayer the txn to the guest module
		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		if isDeployed {
			t.Fatal("expecting wallet to be undeployed")
		}

		// Wallet deployment data
		_, walletFactoryAddress, walletDeployData, err := sequence.EncodeWalletDeployment(wallet.GetWalletConfig(), wallet.GetWalletContext())
		assert.NoError(t, err)

		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(2255), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		// Bundle of transactions: 1.) deploy new wallet 2.) send txn to the wallet
		// Then we send it to the GuestModule

		walletBundle := sequence.Transactions{
			{
				To: wallet.Address(),
				Transactions: sequence.Transactions{
					{
						To:   callmockContract.Address,
						Data: calldata,
					},
				},
			},
		}

		signedWalletBundle, err := wallet.SignTransactions(context.Background(), walletBundle)
		assert.NoError(t, err)

		signedWalletData, err := signedWalletBundle.Execdata()
		assert.NoError(t, err)

		guestBundle := sequence.Transactions{
			{
				To:   walletFactoryAddress,
				Data: walletDeployData,
			},
			{
				To:   wallet.Address(),
				Data: signedWalletData,
			},
		}

		payload := sequence.ConvertTransactionsToV3Payload(guestBundle, signedWalletBundle.Space, signedWalletBundle.Nonce)
		encoded, err := v3.Encode(payload, nil)
		assert.NoError(t, err)

		execdata, err := contracts.V3.WalletStage1Module.Encode("execute", encoded, []byte{})
		assert.NoError(t, err)

		metaTxnBytes, err := v3.HashPayload(wallet.Address(), testChain.ChainID(), payload)
		assert.NoError(t, err)

		metaTxnHash := common.BytesToHash(metaTxnBytes[:])
		metaTxnID := sequence.MetaTxnID(metaTxnHash.Hex()[2:])

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
		guestAddress := testChain.V3SequenceContext().GuestModuleAddress
		ntx, err := sender.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
			To:       &guestAddress,
			Data:     execdata,
			GasLimit: 1000000, // TODO: compute gas limit
		})
		assert.NoError(t, err)

		signedTx, err := sender.SignTx(ntx, testChain.ChainID())
		assert.NoError(t, err)

		_, waitReceipt, err := sender.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		log.Println(callmockContract.Address)

		log.Println("logs:")
		spew.Dump(receipt.Logs)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "2255", ret[0])

		// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
		result, _, _, err := sequence.FetchMetaTransactionReceipt(context.Background(), testChain.ReceiptsListener, metaTxnID)
		assert.NoError(t, err)
		assert.True(t, result.Status == sequence.MetaTxnExecuted)

		// Wallet should be deployed now
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})
}

func TestTransactionSignatureReduction(t *testing.T) {
	data, err := hex.DecodeString("7a9a1628000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002540000000000000000000000000000000000000000000000000000000000000340000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006d08000000000000000000000000a804304f8e7c9ab5932ac3a4cc1468e40183d199000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000450cb9fbb2d44d166aaca1f6cdb1dbd9ff168e4c000000000000000000000000000000000000000000000000002386f26fc1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000be41c134fc3517cb0ec94b6eeafb66cf9998782f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000909693cf965c0cf89c70dbfc2933e91eb10c0faf000000000000000000000000000000000000000000000000c188ce12363a605400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014100050102049c7c6590acb024dd32a543428c940e8d32ea490003b26d43cbbb07ced966e0ee69293b38b39ceb1b0dc1a150ca73af746339ff9e27335c532dfcd49d5d938ca7cce61f3b0776ab4c404d7f75009cb1e8512da766091b0200020e670ff1e43851fec836095827539558ab811cc71f0ceca027e9f033f62a732800f4afce5f72158d97bb5d7af294fe882ed837e5751beaf97479cfe21f4d2ba21b020203596af90cecdbf9a768886e771178fd5561dd27ab005d00010001a2b8a2403d4ae97c103e88438c868b23ccf5221d19e96ae9eee27c0c3a520c7748a29584d57c8a5af0c49cd0749c4c0307b7773232459624327c4737316589db1c020101c50adeadb7fe15bee45dcb820610cdedcd314eb0030102a9273c0f04e71b15c7cd177c6910b4738bb448c30102d5a5934cf965ff79374e8a1a571bd20f54694e63000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)

	data2, err := sequence.ReduceExecdataSignatures(big.NewInt(137), data)
	assert.NoError(t, err)
	assert.Less(t, len(data2), len(data))
}

func TestNonceOverflowChecks(t *testing.T) {
	// 2^160-1 should be a valid nonce space
	// 2^96-1 should be a valid nonce
	_, err := sequence.EncodeNonce(
		new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), big.NewInt(1)),
		new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil), big.NewInt(1)),
	)
	assert.NoError(t, err)

	// 2^160 exceeds the maximum nonce space
	_, err = sequence.EncodeNonce(
		new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil),
		new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil), big.NewInt(1)),
	)
	assert.Error(t, err)

	// 2^96 exceeds the maximum nonce
	_, err = sequence.EncodeNonce(
		new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil), big.NewInt(1)),
		new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil),
	)
	assert.Error(t, err)
}

func TestRawTransactionsEncoding(t *testing.T) {
	data, err := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000003e000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d4c4000000000000000000000000831de831a64405af965c67d6e0de2f9876fa2d9900000000000000000000000000000000000000000000000000005af3107a400000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c33f000000000000000000000000a804304f8e7c9ab5932ac3a4cc1468e40183d199000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000831de831a64405af965c67d6e0de2f9876fa2d990000000000000000000000000000000000000000000000000000b5e620f480000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c33f000000000000000000000000be41c134fc3517cb0ec94b6eeafb66cf9998782f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000831de831a64405af965c67d6e0de2f9876fa2d99000000000000000000000000000000000000000000000000000110d9316ec0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c33f0000000000000000000000009400294c78f7dd244f4b5af6dfff6c971d25ad86000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb000000000000000000000000831de831a64405af965c67d6e0de2f9876fa2d9900000000000000000000000000000000000000000000000000016bcc41e9000000000000000000000000000000000000000000000000000000000000")
	assert.NoError(t, err)

	transactions, err := sequence.DecodeRawTransactions(data)
	assert.NoError(t, err)
	assert.Len(t, transactions, 4)

	data2, err := transactions.EncodeRaw()
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(data, data2))
}
