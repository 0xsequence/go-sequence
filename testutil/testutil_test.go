package testutil_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/stretchr/testify/assert"
)

// yes, we even have to test the testutil

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

func TestTestutil(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(1337))

	// V1DeploySequenceContext
	sequenceContext, err := testChain.V1DeploySequenceContext()
	assert.NoError(t, err)

	// Move to next block
	testChain.Provider.Do(context.Background(), ethrpc.NewCall("evm_mine", nil))

	sequenceContextV2, err := testChain.V2DeploySequenceContext()
	assert.NoError(t, err)

	testChain.Provider.Do(context.Background(), ethrpc.NewCall("evm_mine", nil))

	sequenceContextV3, err := testChain.V3DeploySequenceContext()
	assert.NoError(t, err)

	// Compare against "expexcted" testutil.V1SequenceContext
	expectedContext := testutil.V1SequenceContext()

	expectedContextV2 := testutil.V2SequenceContext()

	expectedContextV3 := testutil.V3SequenceContext()

	assert.Equal(t, expectedContext.FactoryAddress, sequenceContext.FactoryAddress)
	assert.Equal(t, expectedContext.MainModuleAddress, sequenceContext.MainModuleAddress)
	assert.Equal(t, expectedContext.MainModuleUpgradableAddress, sequenceContext.MainModuleUpgradableAddress)
	assert.Equal(t, expectedContext.GuestModuleAddress, sequenceContext.GuestModuleAddress)
	assert.Equal(t, expectedContext.UtilsAddress, sequenceContext.UtilsAddress)

	assert.Equal(t, expectedContextV2.FactoryAddress, sequenceContextV2.FactoryAddress)
	assert.Equal(t, expectedContextV2.MainModuleAddress, sequenceContextV2.MainModuleAddress)
	assert.Equal(t, expectedContextV2.MainModuleUpgradableAddress, sequenceContextV2.MainModuleUpgradableAddress)
	assert.Equal(t, expectedContextV2.GuestModuleAddress, sequenceContextV2.GuestModuleAddress)
	assert.Equal(t, expectedContextV2.UtilsAddress, sequenceContextV2.UtilsAddress)

	assert.Equal(t, expectedContextV3.FactoryAddress, sequenceContextV3.FactoryAddress)
	assert.Equal(t, expectedContextV3.MainModuleAddress, sequenceContextV3.MainModuleAddress)
	assert.Equal(t, expectedContextV3.MainModuleUpgradableAddress, sequenceContextV3.MainModuleUpgradableAddress)
	assert.Equal(t, expectedContextV3.GuestModuleAddress, sequenceContextV3.GuestModuleAddress)
	assert.Equal(t, expectedContextV3.UtilsAddress, sequenceContextV3.UtilsAddress)
}

func TestContractHelpers(t *testing.T) {
	callmockContract := testChain.UniDeploy(t, "WALLET_CALL_RECV_MOCK", 0)

	// Update contract value on CallReceiver by calling 'testCall' contract function
	receipt, err := testutil.ContractTransact(
		testChain.MustWallet(2),
		callmockContract.Address, callmockContract.ABI,
		"testCall", big.NewInt(143), ethcoder.MustHexDecode("0x112233"),
	)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	// Query the value ensuring its been updated on-chain
	ret, err := testutil.ContractQuery(testChain.Provider, callmockContract.Address, "lastValA()", "uint256", nil)
	assert.NoError(t, err)
	assert.Equal(t, []string{"143"}, ret)

	// Query the value using different method, where we unpack the value
	var result *big.Int
	_, err = testutil.ContractCall(testChain.Provider, callmockContract.Address, callmockContract.ABI, &result, "lastValA")
	assert.NoError(t, err)
	assert.Equal(t, uint64(143), result.Uint64())
}
