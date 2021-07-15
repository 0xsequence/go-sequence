package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/relayer"
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

func TestEstimateBatchSequenceTransaction(t *testing.T) {
	wallet, err := testChain.DummySequenceWallet(1)
	assert.NoError(t, err)

	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
	assert.NoError(t, err)
	testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

	setRevertFlagData, err := callmockContract.Encode("setRevertFlag", true)
	assert.NoError(t, err)
	testutil.SignAndSend(t, wallet, callmockContract.Address, setRevertFlagData)

	calldata1, err := callmockContract.Encode("setRevertFlag", false)
	assert.NoError(t, err)

	calldata2, err := callmockContract.Encode("testCall", big.NewInt(6), ethcoder.MustHexDecode("0x223344"))
	assert.NoError(t, err)

	nonce := testChain.RandomNonce()

	txs := sequence.Transactions{
		&sequence.Transaction{
			To:    callmockContract.Address,
			Data:  calldata1,
			Nonce: nonce,
		},
		&sequence.Transaction{
			To:    callmockContract.Address,
			Data:  calldata2,
			Nonce: nonce,
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
	assert.Less(t, estimated-receipt.GasUsed, uint64(30000))

	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "6", ret[0])
}

func TestEstimateSequenceMultipleSigners(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := sequence.WalletConfig{
		Threshold: 3,
		Signers: sequence.WalletConfigSigners{
			{Weight: 2, Address: eoa1.Address()},
			{Weight: 5, Address: eoa2.Address()},
			{Weight: 1, Address: eoa3.Address()},
		},
	}

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config: walletConfig,
	}, eoa1, eoa3)
	assert.NoError(t, err)

	// Set provider on sequence wallet
	err = wallet.SetProvider(testChain.Provider)
	assert.NoError(t, err)

	// Set relayer on sequence wallet, which is used when the wallet sends transactions
	localRelayer, err := relayer.NewLocalRelayer(testChain.GetRelayerWallet())
	assert.NoError(t, err)
	err = wallet.SetRelayer(localRelayer)
	assert.NoError(t, err)

	ok, err := wallet.IsDeployed()
	assert.NoError(t, err)
	assert.False(t, ok)

	sender := testChain.GetRelayerWallet()
	_, _, waitReceipt, err := sequence.DeploySequenceWallet(sender, wallet.GetWalletConfig(), wallet.GetWalletContext())
	assert.NoError(t, err)
	_, err = waitReceipt(context.Background())
	assert.NoError(t, err)

	// Ensure deployment worked
	ok, err = wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, ok)

	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
	assert.NoError(t, err)
	testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

	calldata, err := callmockContract.Encode("testCall", big.NewInt(3333), ethcoder.MustHexDecode("0x11223344"))
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
	assert.Equal(t, "3333", ret[0])
}

func TestEstimateSequenceNestedSigners(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	wal2, err := testChain.DummySequenceWallet(1)
	assert.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := sequence.WalletConfig{
		Threshold: 3,
		Signers: sequence.WalletConfigSigners{
			{Weight: 2, Address: eoa1.Address()},
			{Weight: 1, Address: wal2.Address()},
			{Weight: 2, Address: eoa3.Address()},
		},
	}

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config: walletConfig,
	}, eoa1, eoa3)
	assert.NoError(t, err)

	// Set provider on sequence wallet
	err = wallet.SetProvider(testChain.Provider)
	assert.NoError(t, err)

	// Set relayer on sequence wallet, which is used when the wallet sends transactions
	localRelayer, err := relayer.NewLocalRelayer(testChain.GetRelayerWallet())
	assert.NoError(t, err)
	err = wallet.SetRelayer(localRelayer)
	assert.NoError(t, err)

	ok, err := wallet.IsDeployed()
	assert.NoError(t, err)
	assert.False(t, ok)

	sender := testChain.GetRelayerWallet()
	_, _, waitReceipt, err := sequence.DeploySequenceWallet(sender, wallet.GetWalletConfig(), wallet.GetWalletContext())
	assert.NoError(t, err)
	_, err = waitReceipt(context.Background())
	assert.NoError(t, err)

	// Ensure deployment worked
	ok, err = wallet.IsDeployed()
	assert.NoError(t, err)
	assert.True(t, ok)

	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
	assert.NoError(t, err)
	testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

	calldata, err := callmockContract.Encode("testCall", big.NewInt(3335), ethcoder.MustHexDecode("0x11223344"))
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
	assert.Less(t, estimated-receipt.GasUsed, uint64(45000))

	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "3335", ret[0])
}

func TestPickLowestWeightForEstimation(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa4, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa5, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	config := sequence.WalletConfig{
		Threshold: 4,
		Signers: []sequence.WalletConfigSigner{
			{
				Weight:  2,
				Address: eoa2.Address(),
			},
			{
				Weight:  1,
				Address: eoa1.Address(),
			},
			{
				Weight:  5,
				Address: eoa5.Address(),
			},
			{
				Weight:  4,
				Address: eoa4.Address(),
			},
			{
				Weight:  3,
				Address: eoa3.Address(),
			},
		},
	}

	pick, err := sequence.NewEstimator().PickSigners(context.Background(), config, []bool{true, true, true, true, true})
	assert.NoError(t, err)

	assert.Equal(t, len(pick), 5)
	assert.Equal(t, pick[0], true)
	assert.Equal(t, pick[1], true)
	assert.Equal(t, pick[2], false)
	assert.Equal(t, pick[3], false)
	assert.Equal(t, pick[4], true)
}

func TestPickNonEOAsForEstimation(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa4, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa5, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	config := sequence.WalletConfig{
		Threshold: 5,
		Signers: []sequence.WalletConfigSigner{
			{
				Weight:  2,
				Address: eoa2.Address(),
			},
			{
				Weight:  1,
				Address: eoa1.Address(),
			},
			{
				Weight:  5,
				Address: eoa5.Address(),
			},
			{
				Weight:  4,
				Address: eoa4.Address(),
			},
			{
				Weight:  3,
				Address: eoa3.Address(),
			},
		},
	}

	pick, err := sequence.NewEstimator().PickSigners(context.Background(), config, []bool{true, true, true, false, true})
	assert.NoError(t, err)

	assert.Equal(t, len(pick), 5)
	assert.Equal(t, pick[0], false)
	assert.Equal(t, pick[1], true)
	assert.Equal(t, pick[2], false)
	assert.Equal(t, pick[3], true)
	assert.Equal(t, pick[4], false)
}

func TestPickNonEOAsLowestWeightForEstimation(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa4, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa5, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	config := sequence.WalletConfig{
		Threshold: 5,
		Signers: []sequence.WalletConfigSigner{
			{
				Weight:  2,
				Address: eoa2.Address(),
			},
			{
				Weight:  1,
				Address: eoa1.Address(),
			},
			{
				Weight:  5,
				Address: eoa5.Address(),
			},
			{
				Weight:  4,
				Address: eoa4.Address(),
			},
			{
				Weight:  3,
				Address: eoa3.Address(),
			},
		},
	}

	pick, err := sequence.NewEstimator().PickSigners(context.Background(), config, []bool{false, false, false, false, false})
	assert.NoError(t, err)

	assert.Equal(t, len(pick), 5)
	assert.Equal(t, pick[0], true)
	assert.Equal(t, pick[1], true)
	assert.Equal(t, pick[2], false)
	assert.Equal(t, pick[3], false)
	assert.Equal(t, pick[4], true)
}
