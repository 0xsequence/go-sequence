package sequence_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
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
	err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Check the value
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "55", ret[0])
}

func TestTransactionVerbose(t *testing.T) {
	// This test is similar to TestTransaction but includes a few extra assertions

	// Ensure dummy sequence wallet from seed 1 is deployed
	wallet, err := testChain.DummySequenceWallet(1)
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
	txSubDigest, err := sequence.SubDigest(wallet.Address(), wallet.GetChainID(), signedTx.Digest)
	assert.NoError(t, err)

	walletConfig, err := sequence.RecoverWalletConfigFromDigest(txSubDigest, signedTx.Signature, testutil.SequenceContext(), testChain.ChainID(), testChain.Provider)
	assert.NoError(t, err)

	walletAddress, err := sequence.AddressFromWalletConfig(walletConfig, testutil.SequenceContext())
	assert.NoError(t, err)

	assert.Equal(t, wallet.Address(), walletAddress)

	expectedMetaTxnID, err := sequence.ComputeMetaTxnIDFromTransactionsDigest(walletAddress, testChain.ChainID(), signedTx.Digest, signedTx.Nonce)
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
}

func TestTransactionBundling(t *testing.T) {
	// Ensure three dummy sequence wallets are deployed
	wallets, err := testChain.DummySequenceWallets(3, 1)
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
}

func TestTransactionERC20Transfer(t *testing.T) {
	// Ensure three dummy sequence wallets are deployed
	wallets, err := testChain.DummySequenceWallets(3, 1)
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

	// Create normal txn of: callmockContract.transfer(wallets[1], 20)
	calldata, err = callmockContract.Encode("transfer", wallets[1].Address(), big.NewInt(20))
	assert.NoError(t, err)

	// Sign and send the transaction
	err = testutil.SignAndSend(t, wallets[0], callmockContract.Address, calldata)
	assert.NoError(t, err)

	// Check the value of wallet 1
	ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "80", ret[0])

	// Check the value of wallet 2
	ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[1].Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "20", ret[0])

	// Create normal txns to:
	// 1. callmockContract.transfer(wallets[1], 10)
	// 2. callmockContract.transfer(wallets[2], 30)
	var calldatas [][]byte
	calldata1, err := callmockContract.Encode("transfer", wallets[1].Address(), big.NewInt(10))
	calldatas = append(calldatas, calldata1)
	calldata2, err := callmockContract.Encode("transfer", wallets[2].Address(), big.NewInt(30))
	calldatas = append(calldatas, calldata2)

	// Sign and send the transaction
	err = testutil.BatchSignAndSend(t, wallets[0], callmockContract.Address, calldatas)
	assert.NoError(t, err)

	// Check the value of wallet 1
	ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[0].Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "40", ret[0])

	// Check the value of wallet 2
	ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[1].Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "30", ret[0])

	// Check the value of wallet 3
	ret, err = testutil.ContractQuery(testChain.Provider, callmockContract.Address, "balanceOf(address)", "uint256", []string{wallets[2].Address().Hex()})
	assert.NoError(t, err)
	assert.Len(t, ret, 1)
	assert.Equal(t, "30", ret[0])
}

func TestDecodeExecute(t *testing.T) {
	data, err := hex.DecodeString("7a9a162800000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000070000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000345458cfd2f0c808455342cd0a6e07a09f893292000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000005647a9a16280000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000002e00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000087920000000000000000000000006b175474e89094c44da98b954eedeac495271d0f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb0000000000000000000000008e2158d3fe77a98ba319b763690ea5836461b71e0000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000005208000000000000000000000000b21d86e40c25e8b7041b91d0d01123526c25373f0000000000000000000000000000000000000000000000000009a64bbc29a60000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000025d0005010216c83a218ef8913cb0b11b9121134cb4844ffaf601021d0fca76fe6f638d566e7edfaf54311c760758790102222fda4d1f352ff6985ca4095eb830a181de16d401023bcb633dcf7d35ddd2644d2d6b39badf801fdb3001024495a96369ae15ec93a0a2541eeacce62ec6f48201024580acd41fa3afda624daa55adc249b882a63858010256fa1f76812deb864af2bea60284cd66a5167b97010259361de312447bdc5d9c781f23d0f28516cafc6f0203596af90cecdbf9a768886e771178fd5561dd27ab005d0001000163e44499c8095e94435d9ec8249dadf0090e503ebe01bf0deb4ac4ec5ecb9daf3e3cde592b8df48828a10f3dff4388300b94b87a9a9ad9db6645c5cb8962c6021b020101c50adeadb7fe15bee45dcb820610cdedcd314eb00301025e92183c08a9de426f5bb2d82c4513c1b2d9f34b01026f1e50ccf59d1209b7c89d18d180083fa07ea7510002fa54f98f9c2fefe0a5700f5ee295a2899276ff154790eac8951ecdbd0de9d6970fe115a7594c9805b84c261e75279811aaa27fdb80295f60eba3872f659360d91c0201028a31e1dfba8a0be0e40db39dd2f182589d24603a01029a8acd6c7e88c5761b5d62e7e3af2ccb36aaf1690102b08bd63fec0bb973aa56f14300f2d2ee52f640600102b25e32d18dd1b99433eebade95d48d9a1a71244d0102b2f52d61b4860a1c04ff21aee9a97f81f8524a7b0103bebb0dc9a85e2e828dda0a060cac89dc600c396e0102cc91701809a617cd66517d11f9a3f384f2ed3fe70102db013f06a694c980df22189358d653c10ea24ebc0102e8af73c041cb4105397d424b802e0e1909185b52000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	transactions, nonce, signature, err := sequence.DecodeExecdata(data)
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
	assert.Len(t, transactions[0].Transactions, 2)
	assert.Equal(t, 0, nonce.Cmp(big.NewInt(0)))
	assert.NotNil(t, signature)
	assert.Len(t, signature, 0)

	transaction := sequence.Transaction{
		Transactions: transactions,
		Nonce:        nonce,
		Signature:    signature,
	}

	data2, err := transaction.Execdata()
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(data, data2))
}

func TestDecodeSelfExecute(t *testing.T) {
	data, err := hex.DecodeString("61c2926c000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a804304f8e7c9ab5932ac3a4cc1468e40183d199000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000044a9059cbb00000000000000000000000011111111111111111111111111111111111111110000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000")

	transactions, nonce, signature, err := sequence.DecodeExecdata(data)
	assert.NoError(t, err)
	assert.Len(t, transactions, 1)
	assert.Nil(t, nonce)
	assert.Nil(t, signature)

	transaction := sequence.Transaction{
		Transactions: transactions,
		Nonce:        nonce,
		Signature:    signature,
	}

	data2, err := transaction.Execdata()
	assert.NoError(t, err)
	assert.True(t, bytes.Equal(data, data2))
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
