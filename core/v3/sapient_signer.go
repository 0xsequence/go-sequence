package v3

import (
	"context"
	"fmt"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/v3/walletstage1"
	"github.com/0xsequence/go-sequence/core"
)

func RecoverSapientSignature(ctx context.Context, signer common.Address, payload core.Payload, signature []byte, provider *ethrpc.Provider) (core.ImageHash, error) {
	if provider == nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature without provider")
	}

	var payload_ walletstage1.PayloadDecoded
	switch payload := payload.(type) {
	case CallsPayload:
		payload_ = payload.ABIEncode()
	case MessagePayload:
		payload_ = payload.ABIEncode()
	case ConfigUpdatePayload:
		payload_ = payload.ABIEncode()
	case DigestPayload:
		payload_ = payload.ABIEncode()
	default:
		return core.ImageHash{}, fmt.Errorf("unable to abi encode %T payload", payload)
	}

	data, err := contracts.V3.WalletStage1Module.Encode("recoverSapientSignature", payload_, signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignature call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature: %w", err)
	}

	var imageHash common.Hash
	err = contracts.V3.WalletStage1Module.Decode(&imageHash, "recoverSapientSignature", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignature response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}

func RecoverSapientSignatureCompact(ctx context.Context, signer common.Address, payload core.Payload, signature []byte, provider *ethrpc.Provider) (core.ImageHash, error) {
	if provider == nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact without provider")
	}

	data, err := contracts.V3.WalletStage1Module.Encode("recoverSapientSignatureCompact", payload.Digest().Hash, signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignatureCompact call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact: %w", err)
	}

	var imageHash common.Hash
	err = contracts.V3.WalletStage1Module.Decode(&imageHash, "recoverSapientSignatureCompact", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignatureCompact response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}
