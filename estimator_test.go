package sequence_test

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
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
	assert.NoError(t, err)

	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)

	assert.Less(t, receipt.GasUsed, gas.Uint64())
	assert.Less(t, gas.Uint64()-receipt.GasUsed, uint64(5000))
}

func TestEstimateSimpleSequenceTransactionAndSend(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wallet, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v2", func(t *testing.T) {
		wallet, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v3", func(t *testing.T) {
		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})
}

func TestEstimateSimpleSequenceTransactionNonDeployedWallet(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())

		wallet, err := testChain.V1DummySequenceWallet(rand.Uint64(), true)
		assert.NoError(t, err)

		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		assert.False(t, isDeployed)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		wallet2, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet2, callmockContract.Address, clearData)

		data := make([]byte, 32)
		rand.Read(data)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(771), data)
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)
		assert.NoError(t, err)

		err = testChain.DeploySequenceWallet(wallet)
		assert.NoError(t, err)
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)

		assert.NoError(t, err)
		assert.NotZero(t, estimated)

		signed, err := wallet.SignTransactions(context.Background(), txs)
		assert.NoError(t, err)

		_, _, wait, err := wallet.SendTransactions(context.Background(), signed)
		assert.NoError(t, err)

		receipt, err := wait(context.Background())
		assert.NoError(t, err)

		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "771", ret[0])

		assert.LessOrEqual(t, receipt.GasUsed, estimated)
		assert.Less(t, estimated-receipt.GasUsed, uint64(25000))
	})

	t.Run("v2", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())

		wallet, err := testChain.V2DummySequenceWallet(rand.Uint64(), true)
		assert.NoError(t, err)

		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		assert.False(t, isDeployed)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		wallet2, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet2, callmockContract.Address, clearData)

		data := make([]byte, 32)
		rand.Read(data)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(771), data)
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)
		assert.NoError(t, err)

		err = testChain.DeploySequenceWallet(wallet)
		assert.NoError(t, err)
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)

		assert.NoError(t, err)
		assert.NotZero(t, estimated)

		signed, err := wallet.SignTransactions(context.Background(), txs)
		assert.NoError(t, err)

		_, _, wait, err := wallet.SendTransactions(context.Background(), signed)
		assert.NoError(t, err)

		receipt, err := wait(context.Background())
		assert.NoError(t, err)

		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "771", ret[0])

		assert.LessOrEqual(t, receipt.GasUsed, estimated)
		assert.Less(t, estimated-receipt.GasUsed, uint64(25000))
	})

	t.Run("v3", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())

		wallet, err := testChain.V3DummySequenceWallet(rand.Uint64(), true)
		assert.NoError(t, err)

		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		assert.False(t, isDeployed)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		wallet2, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet2, callmockContract.Address, clearData)

		data := make([]byte, 32)
		rand.Read(data)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(771), data)
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)
		assert.NoError(t, err)

		err = testChain.DeploySequenceWallet(wallet)
		assert.NoError(t, err)
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)

		assert.NoError(t, err)
		assert.NotZero(t, estimated)

		signed, err := wallet.SignTransactions(context.Background(), txs)
		assert.NoError(t, err)

		_, _, wait, err := wallet.SendTransactions(context.Background(), signed)
		assert.NoError(t, err)

		receipt, err := wait(context.Background())
		assert.NoError(t, err)

		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "771", ret[0])

		assert.LessOrEqual(t, receipt.GasUsed, estimated)
		assert.Less(t, estimated-receipt.GasUsed, uint64(25000))
	})
}

func TestEstimateSimpleSequenceTransactionWithStubConfig(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		stubConfig := &v1.WalletConfig{
			Threshold_: 2,
			Signers_: v1.WalletConfigSigners{
				{
					Address: common.HexToAddress("0x6d3A40AAA98DD6cF67a1e6C85807fCc1363935D5"),
					Weight:  1,
				},
				{
					Address: common.HexToAddress("0xE809672D8768fd124196C75e36202C9C0A82740A"),
					Weight:  1,
				},
			},
		}

		wallet, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), stubConfig, wallet.GetWalletContext(), txs)

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
	})

	t.Run("v2", func(t *testing.T) {
		stubConfig := &v2.WalletConfig{
			Threshold_: 2,
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{
					Address: common.HexToAddress("0x6d3A40AAA98DD6cF67a1e6C85807fCc1363935D5"),
					Weight:  1,
				},
				&v2.WalletConfigTreeAddressLeaf{
					Address: common.HexToAddress("0xE809672D8768fd124196C75e36202C9C0A82740A"),
					Weight:  1,
				},
			),
		}

		wallet, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), stubConfig, wallet.GetWalletContext(), txs)

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
	})

	t.Run("v3", func(t *testing.T) {
		stubConfig := &v3.WalletConfig{
			Threshold_: 2,
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{
					Address: common.HexToAddress("0x6d3A40AAA98DD6cF67a1e6C85807fCc1363935D5"),
					Weight:  1,
				},
				&v3.WalletConfigTreeAddressLeaf{
					Address: common.HexToAddress("0xE809672D8768fd124196C75e36202C9C0A82740A"),
					Weight:  1,
				},
			),
		}

		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), stubConfig, wallet.GetWalletContext(), txs)

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
	})
}

func TestEstimateSimpleSequenceTransactionWithBadNonce(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wallet, err := testChain.V1DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		badTxs := sequence.Transactions{
			&sequence.Transaction{
				To:    callmockContract.Address,
				Data:  calldata,
				Nonce: big.NewInt(999),
			},
		}

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), badTxs)

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
	})

	t.Run("v2", func(t *testing.T) {
		wallet, err := testChain.V2DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		badTxs := sequence.Transactions{
			&sequence.Transaction{
				To:    callmockContract.Address,
				Data:  calldata,
				Nonce: big.NewInt(999),
			},
		}

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), badTxs)

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
	})

	t.Run("v3", func(t *testing.T) {
		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		badTxs := sequence.Transactions{
			&sequence.Transaction{
				To:    callmockContract.Address,
				Data:  calldata,
				Nonce: big.NewInt(999),
			},
		}

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), badTxs)

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
	})
}

func TestEstimateBatchSequenceTransaction(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wallet, err := testChain.V1DummySequenceWallet(1)
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

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata1,
			},
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata2,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v2", func(t *testing.T) {
		wallet, err := testChain.V2DummySequenceWallet(1)
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

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata1,
			},
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata2,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v3", func(t *testing.T) {
		wallet, err := testChain.V3DummySequenceWallet(1)
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

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata1,
			},
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata2,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})
}

func TestEstimateSequenceMultipleSigners(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa1, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa2, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa3, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		walletConfig := &v1.WalletConfig{
			Threshold_: 3,
			Signers_: v1.WalletConfigSigners{
				{Weight: 2, Address: eoa1.Address()},
				{Weight: 5, Address: eoa2.Address()},
				{Weight: 1, Address: eoa3.Address()},
			},
		}

		wallet, err := sequence.GenericNewWallet[*v1.WalletConfig](sequence.WalletOptions[*v1.WalletConfig]{
			Config: walletConfig,
		}, eoa1, eoa3)
		assert.NoError(t, err)

		// Set provider on sequence wallet
		err = wallet.SetProvider(testChain.Provider)
		assert.NoError(t, err)

		// Set relayer on sequence wallet, which is used when the wallet sends transactions
		localRelayer, err := relayer.NewLocalRelayer(testChain.GetRelayerWallet(), testChain.ReceiptsListener)
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
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v2", func(t *testing.T) {
		eoa1, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa2, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa3, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		walletConfig := &v2.WalletConfig{
			Threshold_: 3,
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa3.Address(),
				},
			),
		}

		wallet, err := sequence.GenericNewWallet[*v2.WalletConfig](sequence.WalletOptions[*v2.WalletConfig]{
			Config: walletConfig,
		}, eoa1, eoa3)
		assert.NoError(t, err)

		// Set provider on sequence wallet
		err = wallet.SetProvider(testChain.Provider)
		assert.NoError(t, err)

		// Set relayer on sequence wallet, which is used when the wallet sends transactions
		localRelayer, err := relayer.NewLocalRelayer(testChain.GetRelayerWallet(), testChain.ReceiptsListener)
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
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

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
	})

	t.Run("v3", func(t *testing.T) {
		eoa1, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa3, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wal2, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		walletConfig := &v3.WalletConfig{
			Threshold_: 3,
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: eoa1.Address()},
				&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: wal2.Address()},
				&v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: eoa3.Address()},
			),
		}

		wallet, err := sequence.GenericNewWallet[*v3.WalletConfig](sequence.WalletOptions[*v3.WalletConfig]{
			Config: walletConfig,
		}, eoa1, eoa3)
		assert.NoError(t, err)

		// Set provider on sequence wallet
		err = wallet.SetProvider(testChain.Provider)
		assert.NoError(t, err)

		// Set relayer on sequence wallet, which is used when the wallet sends transactions
		localRelayer, err := relayer.NewLocalRelayer(testChain.GetRelayerWallet(), testChain.ReceiptsListener)
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
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		estimator := sequence.NewEstimator()
		estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), wallet.GetWalletConfig(), wallet.GetWalletContext(), txs)

		assert.NoError(t, err)
		assert.NotZero(t, estimated)

		signed, err := wallet.SignTransactions(context.Background(), txs)
		assert.NoError(t, err)

		_, _, wait, err := wallet.SendTransactions(context.Background(), signed)
		assert.NoError(t, err)

		receipt, err := wait(context.Background())
		assert.NoError(t, err)

		assert.LessOrEqual(t, receipt.GasUsed, estimated)
		assert.Less(t, estimated-receipt.GasUsed, uint64(40000))

		ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
		assert.NoError(t, err)
		assert.Len(t, ret, 1)
		assert.Equal(t, "3333", ret[0])
	})
}

func TestPickLowestWeightForEstimation(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
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

		config := &v1.WalletConfig{
			Threshold_: 4,
			Signers_: v1.WalletConfigSigners{
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

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: true,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v2", func(t *testing.T) {
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

		config := &v2.WalletConfig{
			Threshold_: 5,
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: true,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v3", func(t *testing.T) {
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

		config := &v3.WalletConfig{
			Threshold_: 5,
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: true,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})
}

func TestPickNonEOAsForEstimation(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
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

		config := &v1.WalletConfig{
			Threshold_: 5,
			Signers_: v1.WalletConfigSigners{
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

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}] || pick[core.Signer{Address: eoa1.Address()}] || pick[core.Signer{Address: eoa5.Address()}] || pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v2", func(t *testing.T) {
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

		config := &v2.WalletConfig{
			Threshold_: 5,
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}] || pick[core.Signer{Address: eoa1.Address()}] || pick[core.Signer{Address: eoa5.Address()}] || pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v3", func(t *testing.T) {
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

		config := &v3.WalletConfig{
			Threshold_: 5,
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: true,
				{Address: eoa1.Address()}: true,
				{Address: eoa5.Address()}: true,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: true,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}] || pick[core.Signer{Address: eoa1.Address()}] || pick[core.Signer{Address: eoa5.Address()}] || pick[core.Signer{Address: eoa3.Address()}], true)
	})
}

func TestPickNonEOAsLowestWeightForEstimation(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
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

		config := &v1.WalletConfig{
			Threshold_: 5,
			Signers_: v1.WalletConfigSigners{
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

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: false,
				{Address: eoa1.Address()}: false,
				{Address: eoa5.Address()}: false,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: false,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v2", func(t *testing.T) {
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

		config := &v2.WalletConfig{
			Threshold_: 5,
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v2.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: false,
				{Address: eoa1.Address()}: false,
				{Address: eoa5.Address()}: false,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: false,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})

	t.Run("v3", func(t *testing.T) {
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

		config := &v3.WalletConfig{
			Threshold_: 5,
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  2,
					Address: eoa2.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  1,
					Address: eoa1.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  5,
					Address: eoa5.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  4,
					Address: eoa4.Address(),
				},
				&v3.WalletConfigTreeAddressLeaf{
					Weight:  3,
					Address: eoa3.Address(),
				},
			),
		}

		pick, err := sequence.NewEstimator().PickSigners(context.Background(), config,
			map[core.Signer]bool{
				{Address: eoa2.Address()}: false,
				{Address: eoa1.Address()}: false,
				{Address: eoa5.Address()}: false,
				{Address: eoa4.Address()}: false,
				{Address: eoa3.Address()}: false,
			})
		assert.NoError(t, err)

		assert.Equal(t, len(pick), 5)
		assert.Equal(t, pick[core.Signer{Address: eoa2.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa1.Address()}], true)
		assert.Equal(t, pick[core.Signer{Address: eoa5.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa4.Address()}], false)
		assert.Equal(t, pick[core.Signer{Address: eoa3.Address()}], true)
	})
}

func TestEstimateIssue5367MultipleSignersTime(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wallets := v1.WalletConfigSigners{}

		for i := 0; i < 80; i++ {
			randomWallet, err := ethwallet.NewWalletFromRandomEntropy()
			assert.NoError(t, err)

			wallets = append(wallets, &v1.WalletConfigSigner{
				Weight:  1,
				Address: randomWallet.Address(),
			})
		}

		walletConfig := &v1.WalletConfig{
			Threshold_: 2,
			Signers_:   wallets,
		}

		wallet, err := testChain.V1DummySequenceWallet(1)
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

		// run the estimator several times (results should be cached and reused so this loop should be pretty fast)
		for i := 0; i < 100; i++ {
			estimator := sequence.NewEstimator()
			estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), walletConfig, wallet.GetWalletContext(), txs)

			assert.NoError(t, err)
			assert.NotZero(t, estimated)
		}
	})

	t.Run("v2", func(t *testing.T) {
		wallets := []v2.WalletConfigTree{}

		for i := 0; i < 80; i++ {
			randomWallet, err := ethwallet.NewWalletFromRandomEntropy()
			assert.NoError(t, err)

			wallets = append(wallets, &v2.WalletConfigTreeAddressLeaf{
				Weight:  1,
				Address: randomWallet.Address(),
			})
		}

		walletConfig := &v2.WalletConfig{
			Threshold_: 2,
			Tree:       v2.WalletConfigTreeNodes(wallets...),
		}

		wallet, err := testChain.V2DummySequenceWallet(1)
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

		// run the estimator several times (results should be cached and reused so this loop should be pretty fast)
		for i := 0; i < 100; i++ {
			estimator := sequence.NewEstimator()
			estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), walletConfig, wallet.GetWalletContext(), txs)

			assert.NoError(t, err)
			assert.NotZero(t, estimated)
		}
	})

	t.Run("v3", func(t *testing.T) {
		wallets := []v3.WalletConfigTree{}

		for i := 0; i < 80; i++ {
			randomWallet, err := ethwallet.NewWalletFromRandomEntropy()
			assert.NoError(t, err)

			wallets = append(wallets, &v3.WalletConfigTreeAddressLeaf{
				Weight:  1,
				Address: randomWallet.Address(),
			})
		}

		walletConfig := &v3.WalletConfig{
			Threshold_: 2,
			Tree:       v3.WalletConfigTreeNodes(wallets...),
		}

		wallet, err := testChain.V3DummySequenceWallet(1)
		assert.NoError(t, err)

		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		clearData, err := callmockContract.Encode("testCall", big.NewInt(0), ethcoder.MustHexDecode("0x"))
		assert.NoError(t, err)
		testutil.SignAndSend(t, wallet, callmockContract.Address, clearData)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(3), ethcoder.MustHexDecode("0x11223344"))
		assert.NoError(t, err)

		txs := sequence.Transactions{
			&sequence.Transaction{
				To:   callmockContract.Address,
				Data: calldata,
			},
		}

		// run the estimator several times (results should be cached and reused so this loop should be pretty fast)
		for i := 0; i < 100; i++ {
			estimator := sequence.NewEstimator()
			estimated, err := estimator.Estimate(context.Background(), testChain.Provider, wallet.Address(), walletConfig, wallet.GetWalletContext(), txs)

			assert.NoError(t, err)
			assert.NotZero(t, estimated)
		}
	})
}
