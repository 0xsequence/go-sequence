package sequence_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
)

const rpcURL = "https://nodes.sequence.app/mainnet"

func TestIsValidMessageSignatureEOA(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	signature, err := eoa.SignMessage([]byte(message))
	assert.NoError(t, err)
	assert.Len(t, signature, crypto.SignatureLength)

	provider, err := ethrpc.NewProvider(rpcURL)
	assert.NoError(t, err)

	isValid, err := sequence.IsValidMessageSignature(eoa.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValidMessageSignatureSequence(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)
	wallet, err := sequence.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURL)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	signature, _, err := wallet.SignMessage(context.Background(), accounts.TextHash([]byte(message)))
	assert.NoError(t, err)

	signature, err = sequence.EIP6492Signature(signature, wallet.GetWalletConfig())
	assert.NoError(t, err)

	isValid, err := sequence.IsValidMessageSignature(wallet.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}
