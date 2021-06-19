package sequence_test

import (
	"testing"

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
	testChain.MustDeploySequenceContext()
}

func TestChainID(t *testing.T) {
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(4337))
}
