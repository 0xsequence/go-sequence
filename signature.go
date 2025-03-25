package sequence

import (
	"context"
	"fmt"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/eip6492"
)

func IsValidTypedDataSignature(ctx context.Context, signature []byte, signer common.Address, message *ethcoder.TypedData, provider *ethrpc.Provider) (bool, error) {
	digest, err := message.EncodeDigest()
	if err != nil {
		return false, fmt.Errorf("unable to hash typed data: %w", err)
	}

	return IsValidDigestSignature(ctx, signature, signer, common.Hash(digest), provider)
}

func IsValidMessageSignature(ctx context.Context, signature []byte, signer common.Address, message []byte, provider *ethrpc.Provider) (bool, error) {
	chainID, err := provider.ChainID(ctx)
	if err != nil {
		return false, fmt.Errorf("unable to get chain id: %w", err)
	}

	return IsValidSignature(ctx, signature, v3.ConstructMessagePayload(signer, chainID, message), provider)
}

func IsValidDigestSignature(ctx context.Context, signature []byte, signer common.Address, digest common.Hash, provider *ethrpc.Provider) (bool, error) {
	chainID, err := provider.ChainID(ctx)
	if err != nil {
		return false, fmt.Errorf("unable to get chain id: %w", err)
	}

	return IsValidSignature(ctx, signature, v3.ConstructDigestPayload(signer, chainID, digest), provider)
}

func IsValidSignature(ctx context.Context, signature []byte, payload v3.Payload, provider *ethrpc.Provider) (bool, error) {
	return eip6492.ValidateEIP6492Offchain(ctx, provider, payload.Address(), payload.Digest().Hash, signature, nil)
}
