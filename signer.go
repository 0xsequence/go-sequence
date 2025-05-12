package sequence

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type MessageSigner interface {
	SignMessage(msg []byte) ([]byte, error)
}

type DigestSigner interface {
	SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, error)
}

type Signer interface {
	Address() common.Address
}

type SignerMessageSigner interface {
	Signer
	MessageSigner
}

type SignerDigestSigner interface {
	Signer
	DigestSigner
}

// NewSigner adapts an ethwallet.Wallet to the Signer and MessageSigner interfaces.
type walletSignerAdapter struct {
	wallet *ethwallet.Wallet
}

func (w *walletSignerAdapter) Address() common.Address {
	return w.wallet.Address()
}

func (w *walletSignerAdapter) SignMessage(msg []byte) ([]byte, error) {
	return w.wallet.SignMessage(msg)
}

func NewSigner(wallet *ethwallet.Wallet) Signer {
	return &walletSignerAdapter{wallet: wallet}
}
