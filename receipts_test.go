package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	v3 "github.com/0xsequence/go-sequence/core/v3"
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
		Calls: []v3.Call{
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
				Data: sequence.Transactions{Calls: []v3.Call{
					{
						To:   callmockContract.Address,
						Data: send[2],
					}, // logs: 7, balance: 6, executed
					{
						To:   callmockContract.Address,
						Data: send[10],
					}, // logs: 8, balance: -4, failed
				}}.MustSelfExecute(wallet.Address()),
			}, // logs: 9, balance: 6, executed
			{
				To: wallet.Address(),
				Data: sequence.Transactions{Calls: []v3.Call{
					{
						To:              wallet.Address(),
						BehaviorOnError: v3.BehaviorOnErrorRevert,
						Data: sequence.Transactions{ // subBundle
							Calls: []v3.Call{
								{
									To:              callmockContract.Address,
									Data:            send[3],
									BehaviorOnError: v3.BehaviorOnErrorRevert,
								}, // logs: 11, balance: 3, reverted
								{
									To:              callmockContract.Address,
									Data:            send[10],
									BehaviorOnError: v3.BehaviorOnErrorRevert,
								}, // logs: 12, balance: -7, reverted
							},
						}.MustSelfExecute(wallet.Address()),
					}, // logs: 10, balance: 6, reverted
				}}.MustSelfExecute(wallet.Address()),
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
		},
	}

	signedBundle, err := wallet.SignTransactions(context.Background(), bundle)
	assert.NoError(t, err)

	_, waitReceipt, err := wallet.SendTransaction(context.Background(), signedBundle)
	assert.NoError(t, err)
	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)
	logs := receipt.Logs
	assert.Len(t, logs, 16)

	digest, address, space, nonce, receipts, err := sequence.DecodeLogs(logs)
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), address)
	assert.Equal(t, 0, space.Cmp(signedBundle.Space))
	assert.Equal(t, 0, nonce.Cmp(signedBundle.Nonce))
	assert.Len(t, receipts, len(bundle.Calls))
	assert.True(t, isIsomorphic(receipts, bundle.Calls))

	// check that all statuses are correct
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[0].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[1].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, subreceipts(receipts[2])[0].Status)
	assert.Equal(t, sequence.MetaTxnFailed, subreceipts(receipts[2])[1].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[2].Status)
	assert.Equal(t, sequence.MetaTxnReverted, subreceipts(subreceipts(receipts[3])[0])[0].Status)
	assert.Equal(t, sequence.MetaTxnReverted, subreceipts(subreceipts(receipts[3])[0])[1].Status)
	assert.Equal(t, sequence.MetaTxnReverted, subreceipts(receipts[3])[0].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[3].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[4].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[5].Status)
	assert.Equal(t, sequence.MetaTxnFailed, receipts[6].Status)
	assert.Equal(t, sequence.MetaTxnExecuted, receipts[7].Status)

	// check that all top-level logs are correct
	_, _, err = sequence.DecodeNonceChange(logs[0])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[0].Address)

	digest_, index, err := sequence.DecodeCallSucceeded(logs[2])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[2].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(0)))

	digest_, index, err = sequence.DecodeCallSucceeded(logs[4])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[4].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(1)))

	digest_, index, err = sequence.DecodeCallSucceeded(logs[9])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[9].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(2)))

	digest_, index, revert, err := sequence.DecodeCallFailed(logs[10])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[10].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(3)))
	reason, err := receipts[3].Error()
	assert.NoError(t, err)
	reason_, err := abi.UnpackRevert(revert)
	assert.NoError(t, err)
	assert.Equal(t, reason, reason_)

	digest_, index, revert, err = sequence.DecodeCallFailed(logs[11])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[11].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(4)))
	reason, err = receipts[4].Error()
	assert.NoError(t, err)
	reason_, err = abi.UnpackRevert(revert)
	assert.NoError(t, err)
	assert.Equal(t, reason, reason_)

	digest_, index, err = sequence.DecodeCallSucceeded(logs[13])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[13].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(5)))

	digest_, index, revert, err = sequence.DecodeCallFailed(logs[14])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[14].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(6)))
	reason, err = receipts[6].Error()
	assert.NoError(t, err)
	reason_, err = abi.UnpackRevert(revert)
	assert.NoError(t, err)
	assert.Equal(t, reason, reason_)

	digest_, index, err = sequence.DecodeCallSucceeded(logs[16])
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), logs[16].Address)
	assert.Equal(t, digest.Hash, digest_.Hash)
	assert.Equal(t, 0, index.Cmp(big.NewInt(7)))

	// check that all internal logs are correct
	assert.Len(t, receipts[0].Logs, 1)
	assert.Len(t, receipts[1].Logs, 1)
	assert.Len(t, subreceipts(receipts[2])[0].Logs, 1)
	assert.Len(t, subreceipts(receipts[2])[1].Logs, 0)
	assert.Len(t, receipts[2].Logs, 2)
	assert.Len(t, subreceipts(subreceipts(receipts[3])[0])[0].Logs, 0)
	assert.Len(t, subreceipts(subreceipts(receipts[3])[0])[1].Logs, 0)
	assert.Len(t, subreceipts(receipts[3])[0].Logs, 0)
	assert.Len(t, receipts[3].Logs, 0)
	assert.Len(t, receipts[4].Logs, 0)
	assert.Len(t, receipts[5].Logs, 1)
	assert.Len(t, receipts[6].Logs, 0)
	assert.Len(t, receipts[7].Logs, 1)
}

func isIsomorphic(receipts []sequence.Receipt, calls []v3.Call) bool {
	if len(receipts) != len(calls) {
		return false
	}

	// TODO

	return true
}

func subreceipts(receipt sequence.Receipt) []sequence.Receipt {
	_, _, _, _, receipts, _ := receipt.Receipts()
	return receipts
}
