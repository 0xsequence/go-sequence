package testutil_test

import (
	"testing"

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
}

func TestTestutil(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(4337))

	// DeploySequenceContext
	sequenceContext, err := testChain.DeploySequenceContext()
	assert.NoError(t, err)

	// Compare against "expexcted" testutil.SequenceContext
	expectedContext := testutil.SequenceContext()

	assert.Equal(t, expectedContext.FactoryAddress, sequenceContext.FactoryAddress)
	assert.Equal(t, expectedContext.MainModuleAddress, sequenceContext.MainModuleAddress)
	assert.Equal(t, expectedContext.MainModuleUpgradableAddress, sequenceContext.MainModuleUpgradableAddress)
	assert.Equal(t, expectedContext.GuestModuleAddress, sequenceContext.GuestModuleAddress)
	assert.Equal(t, expectedContext.UtilsAddress, sequenceContext.UtilsAddress)
}
