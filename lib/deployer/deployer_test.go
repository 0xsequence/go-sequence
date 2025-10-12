package deployer_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/lib/deployer"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

var (
	testChain *testutil.TestChain
)

func init() {
	var err error
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	testChain, err = testutil.NewTestChain(logger)
	if err != nil {
		panic(err)
	}
	if err := testChain.Connect(); err != nil {
		panic(err)
	}
}

func TestDeployer(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(1337))

	testWallet := testChain.MustWallet(5)

	ud, err := deployer.NewUniversalDeployer(testWallet)
	assert.NoError(t, err)

	testV1SequenceContext := testutil.V1SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V1WalletFactoryAddress, err := ud.Deploy(context.Background(), contracts.V1.WalletFactory.ABI, contracts.V1.WalletFactory.Bin, 0, nil, 1000000)
	assert.NoError(t, err)
	assert.Equal(t, testV1SequenceContext.FactoryAddress, V1WalletFactoryAddress)

	testV2SequenceContext := testutil.V2SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V2WalletFactoryAddress, err := ud.Deploy(context.Background(), contracts.V2.WalletFactory.ABI, contracts.V2.WalletFactory.Bin, 0, nil, 1000000)
	assert.NoError(t, err)
	// v2 is deployed w/ the EIP2470 deployer, not the universal deployer
	assert.NotEqual(t, testV2SequenceContext.FactoryAddress, V2WalletFactoryAddress)

	testV3SequenceContext := testutil.V3SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V3WalletFactoryAddress, err := ud.Deploy(context.Background(), contracts.V3.WalletFactory.ABI, contracts.V3.WalletFactory.Bin, 0, nil, 1000000)
	assert.NoError(t, err)
	// v3 is deployed w/ the EIP2470 deployer, not the universal deployer
	assert.NotEqual(t, testV3SequenceContext.FactoryAddress, V3WalletFactoryAddress)
}
