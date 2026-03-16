package sequence_test

import (
	"log/slog"
	"os"
	"testing"

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
	testChain.MustDeploySequenceContext()
}

func TestChainID(t *testing.T) {
	// Default testchain (hardhat/anvil) is 1337; polygon-zkevm fork is 1101
	assert.NotZero(t, testChain.ChainID().Uint64(), "testchain should report a non-zero chain ID")
	assert.Contains(t, []uint64{1337, 1101}, testChain.ChainID().Uint64(), "expected default testchain (1337) or polygon-zkevm fork (1101)")
}
