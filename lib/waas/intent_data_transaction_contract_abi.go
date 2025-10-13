package intents

import (
	"github.com/0xsequence/ethkit/ethcoder"
)

type contractCallType struct {
	Abi  string `json:"abi"`
	Func string `json:"func"`
	Args []any  `json:"args"`
}

// EncodeContractCall encodes a contract call as a hex encoded calldata.
// NOTE: see ethcoder.EncodeContractCall for more details.
func EncodeContractCall(data *contractCallType) (string, error) {
	callDef := ethcoder.ContractCallDef{
		ABI:  data.Abi,
		Func: data.Func,
		Args: data.Args,
	}
	return ethcoder.EncodeContractCall(callDef)
}
