package sequence_test

import (
	"context"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-ethauth"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/testutil"
	"github.com/goware/logger"
	"github.com/stretchr/testify/assert"
)

func TestEthAuthEIP6492(t *testing.T) {
	signer, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	wallet, err := sequence.NewWalletSingleOwner(signer)
	assert.NoError(t, err)

	log := logger.NewLogger(logger.LogLevel_INFO)
	ethAuth, err := ethauth.New(sequence.ValidateSequenceAccountProof(log))
	assert.NoError(t, err)
	err = ethAuth.ConfigJsonRpcProvider(testutil.DefaultTestChainOptions.NodeURL)
	assert.NoError(t, err)

	proof := ethauth.NewProof()
	proof.Address = wallet.Address().String()
	proof.Claims.App = "TestEthAuthEIP6492"
	proof.Claims.IssuedAt = time.Now().Unix()
	proof.Claims.ExpiresAt = proof.Claims.IssuedAt + 3600

	var digest common.Hash
	digest_, err := proof.Claims.MessageDigest()
	assert.NoError(t, err)
	copy(digest[:], digest_)

	signature, _, err := wallet.SignDigest(context.Background(), digest, testChain.ChainID())
	assert.NoError(t, err)

	signature, err = sequence.EIP6492Signature(signature, wallet.GetWalletConfig())
	assert.NoError(t, err)

	proof.Signature = hexutil.Encode(signature)
	_, err = ethAuth.EncodeProof(proof)
	assert.NoError(t, err)

	isValid, err := ethAuth.ValidateProof(proof)
	assert.NoError(t, err)
	assert.True(t, isValid)
}
