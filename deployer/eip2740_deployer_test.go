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

func TestEIP2740Deployer(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(1337))

	testWallet := testChain.MustWallet(5)

	eip2740Deployer, err := deployer.NewEIP2740Deployer(testWallet)
	require.NoError(t, err)

	testSequenceContext := testutil.V2SequenceContext()

	// Deploy sequence wallet-contract factory and ensure it equals the expected value in testutil
	walletFactoryAddress, err := eip2740Deployer.Deploy(context.Background(), contracts.V2.WalletFactory.ABI, contracts.V2.WalletFactory.Bin, 0, nil, 1000000)
	assert.NoError(t, err)
	assert.Equal(t, testSequenceContext.FactoryAddress, walletFactoryAddress)
}
