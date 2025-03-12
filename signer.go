package sequence

import (
	"context"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
)

type Signer interface {
	Address() common.Address
	Sign(ctx context.Context, subdigest core.Subdigest) (core.SignerSignatureType, []byte, error)
}
