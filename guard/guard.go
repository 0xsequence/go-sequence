package guard

import (
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

func (g *GuardSigningService) IsGuard() bool {
	return true
}
