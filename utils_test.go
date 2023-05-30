package sequence_test

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

func TestDeploySequenceWallet(t *testing.T) {

	t.Run("v1", func(t *testing.T) {
		// Create new single owner smart wallet (initially undeployed, of course)
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa, testutil.V1SequenceContext())
		assert.NoError(t, err)

		wallet.SetProvider(testChain.Provider)
		chainID := wallet.GetChainID()
		assert.Equal(t, uint64(1337), chainID.Uint64())

		// Confirm the wallet is not deployed
		isDeployed, err := wallet.IsDeployed()
		if err != nil {
			t.Fatalf("wallet is deployed, but expecting it to be undeployed")
		}
		assert.False(t, isDeployed)

		relayWallet := testChain.GetRelayerWallet()

		walletAddress, tx, waitReceipt, err := sequence.DeploySequenceWallet(relayWallet, wallet.GetWalletConfig(), wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.NotNil(t, tx)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Confirm wallet is now deployed
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)

		assert.Equal(t, wallet.Address(), walletAddress)
	})

	t.Run("v2", func(t *testing.T) {
		// Create new single owner smart wallet (initially undeployed, of course)
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa, testutil.V2SequenceContext())
		assert.NoError(t, err)

		wallet.SetProvider(testChain.Provider)
		chainID := wallet.GetChainID()
		assert.Equal(t, uint64(1337), chainID.Uint64())

		// Confirm the wallet is not deployed
		isDeployed, err := wallet.IsDeployed()
		if err != nil {
			t.Fatalf("wallet is deployed, but expecting it to be undeployed")
		}
		assert.False(t, isDeployed)

		relayWallet := testChain.GetRelayerWallet()

		walletAddress, tx, waitReceipt, err := sequence.DeploySequenceWallet(relayWallet, wallet.GetWalletConfig(), wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.NotNil(t, tx)

		receipt, err := waitReceipt(context.Background())
		assert.NoError(t, err)
		assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

		// Confirm wallet is now deployed
		isDeployed, err = wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)

		assert.Equal(t, wallet.Address(), walletAddress)
	})

}
