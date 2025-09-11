package deployer_test

import (
	"context"
	"testing"

	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/deployer"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEIP2470Deployer(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(1337))

	testWallet := testChain.MustWallet(5)

	eip2470Deployer, err := deployer.NewEIP2470Deployer(testWallet)
	require.NoError(t, err)

	testV1SequenceContext := testutil.V1SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V1WalletFactoryAddress, err := eip2470Deployer.Deploy(context.Background(), contracts.V1.WalletFactory.ABI, contracts.V1.WalletFactory.Bin, nil, nil, 1000000)
	assert.NoError(t, err)
	// v1 is deployed w/ the universal deployer, not the EIP2470 deployer
	assert.NotEqual(t, testV1SequenceContext.FactoryAddress, V1WalletFactoryAddress)

	testV2SequenceContext := testutil.V2SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V2WalletFactoryAddress, err := eip2470Deployer.Deploy(context.Background(), contracts.V2.WalletFactory.ABI, contracts.V2.WalletFactory.Bin, nil, nil, 1000000)
	assert.NoError(t, err)
	assert.Equal(t, testV2SequenceContext.FactoryAddress, V2WalletFactoryAddress)

	testV3SequenceContext := testutil.V3SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	V3WalletFactoryAddress, err := eip2470Deployer.Deploy(context.Background(), contracts.V3.WalletFactory.ABI, contracts.V3.WalletFactory.Bin, nil, nil, 1000000)
	assert.NoError(t, err)
	assert.Equal(t, testV3SequenceContext.FactoryAddress, V3WalletFactoryAddress)
}
