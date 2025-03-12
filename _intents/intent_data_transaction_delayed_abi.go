package intents

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// Deprecated: use EncodeContractCall instead
type delayedEncodeType struct {
	Abi  string          `json:"abi"`
	Func string          `json:"func"`
	Args json.RawMessage `json:"args"`
}

// Deprecated: use EncodeContractCall instead
func EncodeDelayedABI(data *delayedEncodeType) (string, error) {
	// Get the method from the abi
	_, method, order, _, err := getMethodFromAbi(data.Abi, data.Func)
	if err != nil {
		return "", err
	}

	// Try decode args as array
	var args1 []interface{}
	err = json.Unmarshal(data.Args, &args1)
	if err == nil {
		enc := make([]string, len(args1))
		// String args can be used right away, but any nested
		// `delayedEncodeType` must be handled recursively
		for i, arg := range args1 {
			switch arg := arg.(type) {
			case string:
				enc[i] = arg

			case map[string]interface{}:
				nst := arg

				rjsn, err := json.Marshal(nst["args"])
				if err != nil {
					return "", err
				}

				enc[i], err = EncodeDelayedABI(&delayedEncodeType{
					Abi:  nst["abi"].(string),
					Func: nst["func"].(string),
					Args: json.RawMessage(rjsn),
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

	// Try decode args as object
	var args2 map[string]interface{}
	err = json.Unmarshal(data.Args, &args2)
	if err == nil {
		// Convert args to array using the order from the abi
		// call this method again, for simplicity
		args := make([]interface{}, len(order))
		for i, argName := range order {
			// If argName is not found, then fail
			if _, ok := args2[argName]; !ok {
				return "", fmt.Errorf("arg '%s' not found", argName)
			}

			args[i] = args2[argName]
		}

		jsn, err := json.Marshal(args)
		if err != nil {
			return "", err
		}

		return EncodeDelayedABI(&delayedEncodeType{
			Abi:  data.Abi,
			Func: data.Func,
			Args: jsn,
		})
	}

	return "", err
}

// The abi may be a:
// - already encoded method abi: transferFrom(address,address,uint256)
// - already encoded named method: transferFrom(address from,address to,uint256 val)
// - an array of function abis: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"
// - or a single function abi:  "{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}"
// And it must always return it encoded, like this:
// - transferFrom(address,address,uint256)
// making sure that the method matches the returned one
func getMethodFromAbi(abiSigOrJSON string, method string) (*ethcoder.ABISignature, string, []string, []string, error) {
	//
	// First attempt to parse `abiSigOrJSON` string as a plain method abi
	// ie. transferFrom(address,address,uint256)
	//

	// Handle the case for already encoded method abi.
	// NOTE: we do not need the know the `method` argument here.
	abiSigOrJSON = strings.TrimSpace(abiSigOrJSON)
	if len(abiSigOrJSON) > 0 && strings.Contains(abiSigOrJSON, "(") && abiSigOrJSON[len(abiSigOrJSON)-1] == ')' {
		// NOTE: even though the ethcoder function is `ParseEventDef`, designed for event type parsing
		// the abi format for a single function structure is the same, so it works. Perhaps we will rename
		// `ParseEventDef` in the future, or just add another method with a different name.
		abiSig, err := ethcoder.ParseABISignature(abiSigOrJSON)
		if err != nil {
			return nil, "", nil, nil, err
		}

		// NOTE: we only return the abiSig if the `abiSigOrJSON` is an abi signature
		return &abiSig, abiSig.Signature, abiSig.ArgNames, abiSig.ArgTypes, nil
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
	if strings.HasPrefix(abiSigOrJSON, "[") {
		if err := json.Unmarshal([]byte(abiSigOrJSON), &abis); err != nil {
			return nil, "", nil, nil, err
		}
	} else {
		var singleAbi FunctionAbi
		if err := json.Unmarshal([]byte(abiSigOrJSON), &singleAbi); err != nil {
			return nil, "", nil, nil, err
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
			return nil, method + "(" + strings.Join(paramTypes, ",") + ")", order, paramTypes, nil
		}
	}

	return nil, "", nil, nil, fmt.Errorf("method not found in abi")
}
