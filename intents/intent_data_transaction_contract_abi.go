package intents

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type contractCallType struct {
	Abi  string `json:"abi"`
	Func string `json:"func"`
	Args []any  `json:"args"`
}

func EncodeContractCall(data *contractCallType) (string, error) {
	// Get the method from the abi
	method, _, err := getMethodFromAbi(data.Abi, data.Func)
	if err != nil {
		return "", err
	}

	enc := make([]string, len(data.Args))

	// String args can be used right away, but any nested
	// `contractCallType` must be handled recursively
	for i, arg := range data.Args {
		switch arg := arg.(type) {
		case string:
			enc[i] = arg

		case map[string]interface{}:
			nst := arg

			var funcName string
			if v, ok := nst["func"].(string); ok {
				funcName = v
			}

			args, ok := nst["args"].([]interface{})
			if !ok {
				return "", fmt.Errorf("nested args expected to be an array")
			}

			abi, _ := nst["abi"].(string)

			enc[i], err = EncodeContractCall(&contractCallType{
				Abi:  abi,
				Func: funcName,
				Args: args,
			})
			if err != nil {
				return "", err
			}

		default:
			return "", fmt.Errorf("invalid arg type")
		}
	}

	// Encode the method call
	res, err := ethcoder.AbiEncodeMethodCalldataFromStringValues(method, enc)
	if err != nil {
		return "", err
	}

	return "0x" + common.Bytes2Hex(res), nil
}

// The abi may be a:
// - already encoded method abi: transferFrom(address,address,uint256)
// - already encoded named method: transferFrom(address from,address to,uint256 val)
// - an array of function abis: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"
// - or a single function abi:  "{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}"
// And it must always return it encoded, like this:
// - transferFrom(address,address,uint256)
// making sure that the method matches the returned one
func getMethodFromAbi(abi string, method string) (string, []string, error) {
	//
	// First attempt to parse `abi` string as a plain method abi
	// ie. transferFrom(address,address,uint256)
	//

	// Handle the case for already encoded method abi.
	// NOTE: we do not need the know the `method` argument here.
	abi = strings.TrimSpace(abi)
	if len(abi) > 0 && strings.Contains(abi, "(") && abi[len(abi)-1] == ')' {
		// NOTE: even though the ethcoder function is `ParseEventDef`, designed for event type parsing
		// the abi format for a single function structure is the same, so it works. Perhaps we will rename
		// `ParseEventDef` in the future, or just add another method with a different name.
		eventDef, err := ethcoder.ParseEventDef(abi)
		if err != nil {
			return "", nil, err
		}
		return eventDef.Sig, eventDef.ArgNames, nil
	}

	//
	// If above didn't work, attempt to parse `abi` string as
	// a JSON object of the full abi definition
	//

	type FunctionAbi struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Inputs []struct {
			InternalType string `json:"internalType"`
			Name         string `json:"name"`
			Type         string `json:"type"`
		} `json:"inputs"`
	}

	// Handle array of function abis and single function abi
	var abis []FunctionAbi
	if strings.HasPrefix(abi, "[") {
		if err := json.Unmarshal([]byte(abi), &abis); err != nil {
			return "", nil, err
		}
	} else {
		var singleAbi FunctionAbi
		if err := json.Unmarshal([]byte(abi), &singleAbi); err != nil {
			return "", nil, err
		}
		abis = append(abis, singleAbi)
	}

	// Find the correct method and encode it
	for _, fnAbi := range abis {
		if fnAbi.Name == method {
			var paramTypes []string
			order := make([]string, len(fnAbi.Inputs))
			for i, input := range fnAbi.Inputs {
				paramTypes = append(paramTypes, input.Type)
				order[i] = input.Name
			}
			return method + "(" + strings.Join(paramTypes, ",") + ")", order, nil
		}
	}

	return "", nil, errors.New("Method not found in ABI")
}
