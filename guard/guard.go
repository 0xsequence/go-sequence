package guard

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/signing_service"
)

type GuardSigningServiceParams struct {
	signing_service.SigningServiceParams
}

type GuardSigningService struct {
	*signing_service.SigningService
}

func NewGuardSigningService(params GuardSigningServiceParams) *GuardSigningService {
	return &GuardSigningService{
		SigningService: signing_service.NewSigningService(params.SigningServiceParams),
	}
}

func (g *GuardSigningService) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, core.Signature[core.WalletConfig], error) {
	auxData, err := sequence.AuxDataFromContext(ctx)
	if err != nil {
		return nil, nil, err
	}

	if len(auxData.Sig) == 0 {
		return nil, nil, core.ErrSigningFunctionNotReady
	}

	return g.SigningService.SignDigest(ctx, digest, optChainID...)
}
