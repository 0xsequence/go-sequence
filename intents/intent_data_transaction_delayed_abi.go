package intents

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// NOTE: this is deprecated
type delayedEncodeType struct {
	Abi  string          `json:"abi"`
	Func string          `json:"func"`
	Args json.RawMessage `json:"args"`
}

// NOTE: this is deprecated, use EncodeContractCall instead
func EncodeDelayedABI(data *delayedEncodeType) (string, error) {
	// Get the method from the abi
	method, order, err := getMethodFromAbi(data.Abi, data.Func)
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
