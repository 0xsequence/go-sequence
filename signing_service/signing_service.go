package signing_service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/signing_service/proto"
)

type SigningServiceParams struct {
	SignerAddress common.Address
	SignerWeight  uint8

	Client    proto.SigningService
	AuthToken string
}

type SigningService struct {
	params SigningServiceParams
}

func NewSigningService(params SigningServiceParams) *SigningService {
	return &SigningService{
		params: params,
	}
}

func (r *SigningService) Address() common.Address {
	return r.params.SignerAddress
}

func (r *SigningService) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, core.Signature[core.WalletConfig], error) {
	var chainId *big.Int
	if len(optChainID) > 0 {
		chainId = optChainID[0]
	}

	rCtx := ctx
	if r.params.AuthToken != "" {
		rCtx, _ = proto.WithHTTPRequestHeaders(ctx, map[string][]string{
			"Authorization": {fmt.Sprintf("Bearer %s", r.params.AuthToken)},
		})
	}

	encSig, err := r.params.Client.SignWith(rCtx, r.params.SignerAddress.Hex(), &proto.SignRequest{
		ChainId:     chainId.Uint64(),
		Msg:         ethcoder.HexEncode(digest[:]),
		SignContext: SignContextFromContext(ctx),
	})
	if err != nil {
		return nil, nil, err
	}

	sig := common.FromHex(encSig)

	return sig, nil, nil
}

const SignContextKey = "signing_service_sign_context"

func ContextWithSignContext(ctx context.Context, signer *proto.SignContext) context.Context {
	return context.WithValue(ctx, SignContextKey, signer)
}

func SignContextFromContext(ctx context.Context) *proto.SignContext {
	signCtx, ok := ctx.Value(SignContextKey).(*proto.SignContext)
	if !ok {
		return nil
	}
	return signCtx
}
