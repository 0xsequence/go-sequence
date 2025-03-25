package erc6492

import (
	"bytes"
	"fmt"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

var MagicBytes = hexutil.MustDecode("0x6492649264926492649264926492649264926492649264926492649264926492")

func Signature(signature []byte, deployConfig core.WalletConfig) ([]byte, error) {
	context := sequence.SequenceContextForWalletConfig(deployConfig)
	data, err := contracts.V3.WalletFactory.Encode("deploy", context.MainModuleAddress, deployConfig.ImageHash().Hash)
	if err != nil {
		return nil, fmt.Errorf("unable to encode erc-6492 deploy call: %w", err)
	}

	wrapped, err := ethcoder.ABIPackArguments([]string{"address", "bytes", "bytes"}, []any{context.FactoryAddress, data, signature})
	if err != nil {
		return nil, fmt.Errorf("unable to wrap erc-6492 signature: %w", err)
	}

	return bytes.Join([][]byte{wrapped, MagicBytes}, nil), nil
}

func MultiWalletSignature(signature []byte, deployConfig ...core.WalletConfig) ([]byte, error) {
	contexts := sequence.SequenceContexts()

	var calls []v3.Call
	for i, deployConfig := range deployConfig {
		var version uint16
		switch deployConfig.(type) {
		case *v1.WalletConfig:
			version = 1
		case *v2.WalletConfig:
			version = 2
		case *v3.WalletConfig:
			version = 3
		}

		data, err := contracts.V3.WalletFactory.Encode("deploy", contexts[version].MainModuleAddress, deployConfig.ImageHash().Hash)
		if err != nil {
			return nil, fmt.Errorf("unable to encode erc-6492 deploy call for config %v: %w", i, err)
		}

		calls = append(calls, v3.Call{To: contexts[version].FactoryAddress, Data: data})
	}

	data := v3.ConstructCallsPayload(contexts[3].GuestModuleAddress, nil, calls, nil, nil).Encode(contexts[3].GuestModuleAddress)

	wrapped, err := ethcoder.ABIPackArguments([]string{"address", "bytes", "bytes"}, []any{contexts[3].GuestModuleAddress, data, signature})
	if err != nil {
		return nil, fmt.Errorf("unable to wrap erc-6492 signature: %w", err)
	}

	return bytes.Join([][]byte{wrapped, MagicBytes}, nil), nil
}
