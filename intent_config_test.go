package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIntentMachineTransactionVerbose(t *testing.T) {
	t.Run("v2", func(t *testing.T) {
		// Create normal txn of: callmockContract.testCall(55, 0x112255)
		callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

		calldata, err := callmockContract.Encode("testCall", big.NewInt(65), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		calldata2, err := callmockContract.Encode("testCall", big.NewInt(66), ethcoder.MustHexDecode("0x332255"))
		assert.NoError(t, err)

		stx := &sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata,
		}

		stx2 := &sequence.Transaction{
			To:   callmockContract.Address,
			Data: calldata2,
		}

		// Ensure dummy sequence wallet from seed 1 is deployed w/ the intent config
		wallet, err := testChain.V2DummySequenceWalletWithIntentConfig(1, []*sequence.Transaction{stx, stx2})
		assert.NoError(t, err)
		assert.NotNil(t, wallet)

		// Sign and send the transaction
		signedTx, err := wallet.SignTransaction(context.Background(), stx)
		assert.NoError(t, err)

		assert.NotEmpty(t, signedTx.Digest)
		assert.NotEmpty(t, signedTx.Signature)

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
}
