package sequence

import (
	"context"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Signer interface {
	Address() common.Address
	Sign(ctx context.Context, payload v3.Payload) (core.SignerSignatureType, []byte, error)
}
