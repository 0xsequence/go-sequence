package sequence

import (
	"context"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Signer interface {
	Address() common.Address
	Sign(ctx context.Context, payload v3.Payload) (core.SignerSignatureType, []byte, error)
}

func EOA(signer *ethwallet.Wallet) Signer {
	return eoa{signer}
}

type eoa struct{ *ethwallet.Wallet }

func (e eoa) Sign(ctx context.Context, payload v3.Payload) (core.SignerSignatureType, []byte, error) {
	signature, err := e.SignData(payload.Digest().Bytes())
	return core.SignerSignatureTypeEIP712, signature, err
}
