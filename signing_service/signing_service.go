package signing_service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/signing_service/proto"
)

type AuxData struct {
	Msg     []byte
	Sig     []byte
	ChainID *big.Int
	Address *common.Address
}

func (a *AuxData) Pack() ([]byte, error) {
	return ethcoder.AbiCoder(
		[]string{"address", "uint256", "bytes", "bytes"},
		[]interface{}{a.Address, a.ChainID, &a.Msg, &a.Sig},
	)
}

func ContextWithAuxData(ctx context.Context, auxData *AuxData) context.Context {
	return context.WithValue(ctx, "auxData", auxData)
}

func AuxDataFromContext(ctx context.Context) (*AuxData, error) {
	auxData, ok := ctx.Value("auxData").(*AuxData)
	if !ok {
		return nil, fmt.Errorf("auxData not found in context")
	}
	return auxData, nil
}

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

func (r *SigningService) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, *sequence.Signature, error) {
	var chainId *big.Int
	if len(optChainID) > 0 {
		chainId = optChainID[0]
	}

	var auxDataPacked []byte
	if auxData, err := AuxDataFromContext(ctx); err == nil {
		auxDataPacked, err = auxData.Pack()
		if err != nil {
			return nil, nil, err
		}
	}

	rCtx := ctx
	if r.params.AuthToken != "" {
		rCtx, _ = proto.WithHTTPRequestHeaders(ctx, map[string][]string{
			"Authorization": {fmt.Sprintf("Bearer %s", r.params.AuthToken)},
		})
	}

	encSig, err := r.params.Client.Sign(rCtx, &proto.SignRequest{
		ChainId: chainId.Uint64(),
		Msg:     ethcoder.HexEncode(digest[:]),
		AuxData: ethcoder.HexEncode(auxDataPacked),
	})
	if err != nil {
		return nil, nil, err
	}

	sig := common.FromHex(encSig)

	var sigPartType uint8
	switch sig[len(sig)-1] {
	case sequence.SignatureTypeEthSign:
		sigPartType = sequence.SignaturePartTypeEOA
	case sequence.SignatureTypeEip1271:
		sigPartType = sequence.SignaturePartTypeDynamic
	default:
		return nil, nil, fmt.Errorf("invalid signature type: %d", encSig[65])
	}

	// add signature
	var signatures sequence.Signature
	signatures.Signers = append(signatures.Signers, &sequence.SignaturePart{
		Weight:  r.params.SignerWeight,
		Address: r.params.SignerAddress,
		Type:    sigPartType,
		Value:   sig,
	})

	return sig, &signatures, nil
}

var _ sequence.SignerDigestSigner = &SigningService{}
