package hydrate

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

//go:embed abis/HydrateProxy.abi.json
var hydrateProxyABIJSON []byte

// ABI is the parsed HydrateProxy contract ABI (from trails-contracts artifact).
var ABI abi.ABI

func init() {
	var err error
	ABI, err = abi.JSON(bytes.NewReader(hydrateProxyABIJSON))
	if err != nil {
		panic(fmt.Sprintf("hydrate: parse embedded HydrateProxy ABI: %v", err))
	}
}

// PackHydrateExecute ABI-encodes a call to HydrateProxy.hydrateExecute.
func PackHydrateExecute(packedPayload, hydratePayload []byte) ([]byte, error) {
	return ABI.Pack("hydrateExecute", packedPayload, hydratePayload)
}

// PackHydrateExecuteAndSweep ABI-encodes hydrateExecuteAndSweep.
func PackHydrateExecuteAndSweep(packedPayload, hydratePayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) ([]byte, error) {
	return ABI.Pack("hydrateExecuteAndSweep", packedPayload, hydratePayload, sweepTarget, tokensToSweep, sweepNative)
}
