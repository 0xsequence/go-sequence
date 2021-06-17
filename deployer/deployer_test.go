package deployer_test

import (
	"context"
	"testing"

	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/deployer"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

var (
	testChain *testutil.TestChain
)

func init() {
	var err error
	testChain, err = testutil.NewTestChain()
	if err != nil {
		panic(err)
	}
}

func TestDeployer(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(4337))

	testWallet := testChain.MustWallet(5)

	ud, err := deployer.NewUniversalDeployer(testWallet)
	assert.NoError(t, err)

	testSequenceContext := testutil.SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	walletFactoryAddress, err := ud.Deploy(context.Background(), contracts.WalletFactory.ABI, contracts.WalletFactory.Bin, 0, nil)
	assert.NoError(t, err)
	assert.Equal(t, testSequenceContext.FactoryAddress, walletFactoryAddress)
}
