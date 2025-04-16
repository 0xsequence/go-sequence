package deployer_test

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
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
	// v2 is deployed through the eip2470 factory address
	assert.NotEqual(t, testV2SequenceContext.FactoryAddress, V2WalletFactoryAddress)
	assert.Equal(t, common.HexToAddress("0xAA4A319D9f7b43e197555d1c219F9195B80F0C8c"), V2WalletFactoryAddress)

	testV3SequenceContext := testutil.V3SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V3WalletFactoryAddress, err := ud.Deploy(context.Background(), contracts.V3.WalletFactory.ABI, contracts.V3.WalletFactory.Bin, 0, nil, 1000000)
	assert.NoError(t, err)
	// v3 is deployed through the eip2470 factory address
	assert.NotEqual(t, testV3SequenceContext.FactoryAddress, V3WalletFactoryAddress)
	assert.Equal(t, common.HexToAddress("0x0DA1aA1c34FB9a70De8aC275c49abbC8b2420Ec0"), V3WalletFactoryAddress)
}
