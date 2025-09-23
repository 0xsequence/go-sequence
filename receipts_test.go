package sequence_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestReceiptDecoding(t *testing.T) {
	// Ensure a dummy sequence wallet is deployed
	wallets, err := testChain.V1DummySequenceWallets(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, wallets)
	assert.Len(t, wallets, 1)
	wallet := wallets[0]

	// Create normal txn of: callmockContract.mockMint(wallet, 10)
	callmockContract, _ := testChain.Deploy(t, "ERC20Mock")
	calldata, err := callmockContract.Encode("mockMint", wallet.Address(), big.NewInt(10))
	assert.NoError(t, err)

	// Sign and send the transaction
	err = testutil.SignAndSend(t, wallet, callmockContract.Address, calldata)
	assert.NoError(t, err)

	var send [][]byte
	for i := 0; i <= 10; i++ {
		data, err := callmockContract.Encode("transfer", common.Address{1}, big.NewInt(int64(i)))
		assert.NoError(t, err)
		send = append(send, data)
	}

	// These are wallet's intended transactions
	bundle := sequence.Transactions{
		{
			To:   callmockContract.Address,
			Data: send[1],
		}, // logs: 3 (1 from NonceChange), balance: 9, executed
		{
			To:   callmockContract.Address,
			Data: send[1],
		}, // logs: 5, balance: 8, executed
		{
			To: wallet.Address(),
			Transactions: sequence.Transactions{
				{
					To:   callmockContract.Address,
					Data: send[2],
				}, // logs: 7, balance: 6, executed
				{
					To:   callmockContract.Address,
					Data: send[10],
				}, // logs: 8, balance: -4, failed
			},
		}, // logs: 9, balance: 6, executed
		{
			To: wallet.Address(),
			Transactions: sequence.Transactions{
				{
					To:            wallet.Address(),
					RevertOnError: true,
					Transactions: sequence.Transactions{ // subBundle
						{
							To:            callmockContract.Address,
							Data:          send[3],
							RevertOnError: true,
						}, // logs: 11, balance: 3, reverted
						{
							To:            callmockContract.Address,
							Data:          send[10],
							RevertOnError: true,
						}, // logs: 12, balance: -7, reverted
					},
				}, // logs: 10, balance: 6, reverted
			},
		}, // logs: 10, balance: 6, failed
		{
			To:   callmockContract.Address,
			Data: send[10],
		}, // logs: 11, balance: -4, failed
		{
			To:   callmockContract.Address,
			Data: send[4],
		}, // logs: 13, balance: 2, executed
		{
			To:   callmockContract.Address,
			Data: send[4],
		}, // logs: 14, balance: -2, failed
		{
			To:   callmockContract.Address,
			Data: send[2],
		}, // logs: 16, balance: 0, executed
	}

	signedBundle, err := wallet.SignTransactions(context.Background(), bundle)
	assert.NoError(t, err)

	_, _, waitReceipt, err := wallet.SendTransaction(context.Background(), signedBundle)
	assert.NoError(t, err)
	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)
	assert.Len(t, receipt.Logs, 16)

	receipts, logs, err := sequence.DecodeReceipt(context.Background(), receipt, wallet.GetProvider())
	assert.NoError(t, err)
	assert.Len(t, receipts, len(bundle))
	assert.Len(t, logs, 1+len(bundle)) // 1 NonceChange + n TxExecuted/TxFailed
	assert.True(t, haveSameMetaTxnID(receipts))

	for _, child := range receipts {
		assert.True(t, hasNativeReceipt(child, receipt))
		assert.NoError(t, areIsomorphic(child, child.Transaction))
	}

	// check that all statuses are correct
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[0].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[1].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[2].Receipts[0].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[2].Receipts[1].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[2].Status)
	assert.Equal(t, sequence.MetaTxnReverted, receipts[3].Receipts[0].Receipts[0].Status)
	assert.Equal(t, sequence.MetaTxnReverted, receipts[3].Receipts[0].Receipts[1].Status)
	assert.Equal(t, sequence.MetaTxnReverted, receipts[3].Receipts[0].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[3].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[4].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[5].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[6].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[7].Status)

	// check that all top-level logs are correct
	metaTxnID := string(receipts[0].MetaTxnID)
	_, _, err = sequence.DecodeNonceChangeEvent(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, metaTxnID, hex.EncodeToString(logs[1].Data))
	assert.Equal(t, metaTxnID, hex.EncodeToString(logs[2].Data))
	assert.Equal(t, metaTxnID, hex.EncodeToString(logs[3].Data))
	hash, _, err := sequence.V1DecodeTxFailedEvent(logs[4])
	assert.NoError(t, err)
	assert.Equal(t, metaTxnID, hex.EncodeToString(hash[:]))
	hash, _, err = sequence.V1DecodeTxFailedEvent(logs[5])
	assert.NoError(t, err)
	assert.Equal(t, metaTxnID, hex.EncodeToString(hash[:]))
	assert.Equal(t, metaTxnID, hex.EncodeToString(logs[6].Data))
	hash, _, err = sequence.V1DecodeTxFailedEvent(logs[7])
	assert.NoError(t, err)
	assert.Equal(t, metaTxnID, hex.EncodeToString(hash[:]))
	assert.Equal(t, metaTxnID, hex.EncodeToString(logs[8].Data))

	// check that all internal logs are correct
	assert.Len(t, receipts[0].Logs, 1)
	assert.Len(t, receipts[1].Logs, 1)
	assert.Len(t, receipts[2].Receipts[0].Logs, 1)
	assert.Len(t, receipts[2].Receipts[1].Logs, 0)
	assert.Len(t, receipts[2].Logs, 2)
	assert.Len(t, receipts[3].Receipts[0].Receipts[0].Logs, 0)
	assert.Len(t, receipts[3].Receipts[0].Receipts[1].Logs, 0)
	assert.Len(t, receipts[3].Receipts[0].Logs, 0)
	assert.Len(t, receipts[3].Logs, 0)
	assert.Len(t, receipts[4].Logs, 0)
	assert.Len(t, receipts[5].Logs, 1)
	assert.Len(t, receipts[6].Logs, 0)
	assert.Len(t, receipts[7].Logs, 1)

	// check that sequence.DecodeTxFailedEvent can decode TxFailed events
	_, reason, err := sequence.V1DecodeTxFailedEvent(receipts[2].Logs[1])
	assert.NoError(t, err)
	assert.Equal(t, reason, receipts[2].Receipts[1].Reason)

	// check that Receipt.Find() can find our sub-bundle
	subBundle := sequence.Transactions{
		{
			To:            callmockContract.Address,
			Data:          send[3],
			RevertOnError: true,
		},
		{
			To:            callmockContract.Address,
			Data:          send[10],
			RevertOnError: true,
		},
	}
	subBundleID, _, err := sequence.ComputeMetaTxnID(wallet.GetChainID(), wallet.Address(), subBundle, nil, sequence.MetaTxnSelfExec)
	assert.NoError(t, err)
	assert.Equal(t, receipts[3].Find(subBundleID), receipts[3].Receipts[0])
}

func hasNativeReceipt(receipt *sequence.Receipt, native *types.Receipt) bool {
	if receipt.Receipt != native {
		return false
	}

	for _, child := range receipt.Receipts {
		if !hasNativeReceipt(child, native) {
			return false
		}
	}

	return true
}

func haveSameMetaTxnID(receipts []*sequence.Receipt) bool {
	if len(receipts) <= 1 {
		return true
	}

	if len(receipts[0].MetaTxnID) == 0 {
		return false
	}

	for _, child := range receipts {
		if child.MetaTxnID != receipts[0].MetaTxnID {
			return false
		}

		if !haveSameMetaTxnID(child.Receipts) {
			return false
		}
	}

	return true
}

func areIsomorphic(receipt *sequence.Receipt, transaction *sequence.Transaction) error {
	if receipt.Transaction != transaction {
		return fmt.Errorf("receipt.Transaction != transaction")
	}

	if len(receipt.Receipts) != len(transaction.Transactions) {
		return fmt.Errorf("have %v receipts, expected %v", len(receipt.Receipts), len(transaction.Transactions))
	}

	for i := 0; i < len(receipt.Receipts); i++ {
		if err := areIsomorphic(receipt.Receipts[i], transaction.Transactions[i]); err != nil {
			return err
		}
	}

	return nil
}

func TestZKSyncReceipts(t *testing.T) {
	ctx := context.Background()

	provider, err := ethrpc.NewProvider("https://nodes.sequence.app/sandbox-testnet")
	assert.NoError(t, err)

	hash := common.HexToHash("0x6ddb43ddf29ed48f302c6b7b6dc9bb47b33dd5099540fb0793f239edb406fa5d")

	receipt, err := provider.TransactionReceipt(ctx, hash)
	assert.NoError(t, err)

	_, _, err = sequence.DecodeReceipt(ctx, receipt, provider)
	assert.NoError(t, err)
}
