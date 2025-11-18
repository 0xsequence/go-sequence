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

var passkeySigners = map[common.Address]struct{}{
	common.HexToAddress("0x0000000000005204F3711851EAD52CC9c241499a"): {}, // rc4
	common.HexToAddress("0x0000000000dc2d96870dc108c5E15570B715DFD2"): {}, // rc3
	common.HexToAddress("0x4491845806B757D67BE05BbD877Cab101B9bee5C"): {}, // dev2
	common.HexToAddress("0x8f26281dB84C18aAeEa8a53F94c835393229d296"): {}, // dev1
}

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

	data, err := contracts.ISapient.Encode("recoverSapientSignature", payload_, signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignature call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature: %w", err)
	}

	var imageHash common.Hash
	err = contracts.ISapient.Decode(&imageHash, "recoverSapientSignature", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignature response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}

func RecoverSapientSignatureCompact(ctx context.Context, signer common.Address, payload core.Payload, signature []byte, provider *ethrpc.Provider) (core.ImageHash, error) {
	if _, ok := passkeySigners[signer]; ok {
		signature_, err := core.DecodePasskeySignature(signature)
		if err != nil {
			return core.ImageHash{}, fmt.Errorf("unable to decode passkey signature: %w", err)
		}

		if !signature_.IsValid(payload.Digest().Hash) {
			return core.ImageHash{}, fmt.Errorf("invalid passkey signature")
		}

		return signature_.ImageHash(), nil
	}

	if provider == nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact without provider")
	}

	data, err := contracts.ISapientCompact.Encode("recoverSapientSignatureCompact", payload.Digest().Hash, signature)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to encode recoverSapientSignatureCompact call: %w", err)
	}

	response, err := provider.CallContract(ctx, ethereum.CallMsg{To: &signer, Data: data}, nil)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to recover sapient signature compact: %w", err)
	}

	var imageHash common.Hash
	err = contracts.ISapientCompact.Decode(&imageHash, "recoverSapientSignatureCompact", response)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("unable to decode recoverSapientSignatureCompact response: %w", err)
	}

	return core.ImageHash{Hash: imageHash}, nil
}
