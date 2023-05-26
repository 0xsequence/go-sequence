package sequence

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
)

type MessageSigner interface {
	SignMessage(msg []byte) ([]byte, error)
}

type DigestSigner interface {
	SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, core.Signature[core.WalletConfig], error)
}

type GuardSigner interface {
	IsGuard() bool
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

type SignerGuardSigner interface {
	Signer
	DigestSigner
	GuardSigner
}