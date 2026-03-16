package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

var timedRefundSapientImageHashArguments = mustTimedRefundSapientImageHashArguments()

// TimedRefundIntentConfigurationSigner represents the dedicated timed-refund sapient signer
// attached to an intent configuration. Weight must be greater than zero.
type TimedRefundIntentConfigurationSigner struct {
	Address         common.Address
	Destination     common.Address
	UnlockTimestamp uint64
	Weight          uint8
}

// CreateIntentConfigurationWithTimedRefundSapient creates an intent configuration that includes
// a timed-refund sapient signer leaf in addition to the default any-address subdigests.
func CreateIntentConfigurationWithTimedRefundSapient(
	mainSigner common.Address,
	calls []*v3.CallsPayload,
	timedRefundSigner TimedRefundIntentConfigurationSigner,
) (*v3.WalletConfig, error) {
	timedRefundLeaf, err := createTimedRefundSapientSignerLeaf(timedRefundSigner)
	if err != nil {
		return nil, err
	}

	return createIntentConfiguration(mainSigner, calls, timedRefundLeaf)
}

func createTimedRefundSapientSignerLeaf(signer TimedRefundIntentConfigurationSigner) (*v3.WalletConfigTreeSapientSignerLeaf, error) {
	if signer.Address == (common.Address{}) {
		return nil, fmt.Errorf("timed refund sapient signer address is zero")
	}
	if signer.Destination == (common.Address{}) {
		return nil, fmt.Errorf("timed refund destination is zero")
	}
	if signer.UnlockTimestamp == 0 {
		return nil, fmt.Errorf("timed refund unlock timestamp is zero")
	}
	if signer.Weight == 0 {
		return nil, fmt.Errorf("timed refund sapient signer weight is zero")
	}

	imageHash, err := timedRefundSapientImageHash(signer.Destination, signer.UnlockTimestamp)
	if err != nil {
		return nil, err
	}

	return &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     signer.Weight,
		Address:    signer.Address,
		ImageHash_: imageHash,
	}, nil
}

func mustTimedRefundSapientImageHashArguments() abi.Arguments {
	stringType, err := abi.NewType("string", "", nil)
	if err != nil {
		panic(fmt.Errorf("create timed refund string ABI type: %w", err))
	}

	addressType, err := abi.NewType("address", "", nil)
	if err != nil {
		panic(fmt.Errorf("create timed refund address ABI type: %w", err))
	}

	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		panic(fmt.Errorf("create timed refund uint256 ABI type: %w", err))
	}

	return abi.Arguments{
		{Type: stringType},
		{Type: addressType},
		{Type: uint256Type},
	}
}

func timedRefundSapientImageHash(destination common.Address, unlockTimestamp uint64) (core.ImageHash, error) {
	encoded, err := timedRefundSapientImageHashArguments.Pack(
		"timed-refund",
		destination,
		new(big.Int).SetUint64(unlockTimestamp),
	)
	if err != nil {
		return core.ImageHash{}, fmt.Errorf("failed to ABI pack timed refund sapient image hash: %w", err)
	}

	return core.ImageHash{Hash: crypto.Keccak256Hash(encoded)}, nil
}
