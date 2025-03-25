package sequence_test

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
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
		_, err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
		assert.NoError(t, err)

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "55", ret[0])
	})
}

func TestTransactionVerbose(t *testing.T) {
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

		stx := v3.Call{
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

		isValid, err := sequence.IsValidSignature(context.Background(), signedTx.Signature, signedTx, testChain.Provider)
		assert.NoError(t, err)
		assert.True(t, isValid)

		// Send the transaction
		tx, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
		assert.NoError(t, err)
		assert.NotNil(t, tx)

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
		_, err = testutil.SignAndSend(t, wallets[0], callmockContract.Address, calldata)
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
			Calls: []v3.Call{
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
					Data: sequence.Transactions{
						Calls: []v3.Call{
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
					}.MustSelfExecute(wallets[0].Address()),
				},
			},
		}

		// wallet[0] must sign its bundle
		signedBundle, err := wallets[0].SignTransactions(context.Background(), bundle)
		assert.NoError(t, err)

		_, waitReceipt, err := wallets[0].SendTransaction(context.Background(), signedBundle)
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
	t.Run("v3", func(t *testing.T) {
		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)
		calldata, err := callmockContract.Encode("testCall", big.NewInt(1239), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		guestAddress := testChain.V3SequenceContext().GuestModuleAddress
		transactions := sequence.Transactions{
			Calls: []v3.Call{{
				To:   callmockContract.Address,
				Data: calldata,
			}},
		}
		payload := transactions.Payload(guestAddress, testChain.ChainID())

		execdata, err := contracts.V3.WalletStage1Module.Encode("execute", payload.Encode(guestAddress), []byte{})
		assert.NoError(t, err)

		sender := testChain.GetRelayerWallet()

		// Relay the txn manually, directly to the guest module
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
		assert.True(t, bytes.Contains(receipt.Logs[0].Data, payload.Digest().Bytes()))

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "1239", ret[0])

	})
}

func TestTransactionToGuestModuleDeployAndCall(t *testing.T) {
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
			Calls: []v3.Call{
				{
					To: wallet.Address(),
					Data: sequence.Transactions{
						Calls: []v3.Call{
							{
								To:   callmockContract.Address,
								Data: calldata,
							},
						},
					}.MustSelfExecute(wallet.Address()),
				},
			},
		}

		signedWalletBundle, err := wallet.SignTransactions(context.Background(), walletBundle)
		assert.NoError(t, err)

		payload := walletBundle.Payload(wallet.Address(), testChain.ChainID())
		signedExecdata, err := contracts.V3.WalletStage1Module.Encode("execute", payload.Encode(wallet.Address()), signedWalletBundle.Signature)
		assert.NoError(t, err)

		guestBundle := []v3.Call{
			{
				To:   walletFactoryAddress,
				Data: walletDeployData,
			},
			{
				To:   wallet.Address(),
				Data: signedExecdata,
			},
		}

		guestAddress := testChain.V3SequenceContext().GuestModuleAddress
		execdata := v3.ConstructCallsPayload(guestAddress, testChain.ChainID(), guestBundle, nil, nil).Encode(guestAddress)

		// Relay the txn manually, directly to the guest module
		sender := testChain.GetRelayerWallet()
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

		// Check the value
		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "2255", ret[0])

		// Assert sequence.WaitForMetaTxn is able to find the metaTxnID
		_, status, _, _, err := sequence.FetchReceipt(context.Background(), testChain.ReceiptsListener, payload)
		assert.NoError(t, err)
		assert.True(t, status == sequence.MetaTxnExecuted)

		// Wallet should be deployed now
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})
}
