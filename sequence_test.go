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
	assert.Equal(t, testChain.ChainID().Uint64(), uint64(1337))
}
