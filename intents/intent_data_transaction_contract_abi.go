package intents

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/davecgh/go-spew/spew"
)

type contractCallType struct {
	Abi  string `json:"abi"`
	Func string `json:"func"`
	Args []any  `json:"args"`
}

func EncodeContractCall(data *contractCallType) (string, error) {
	data.Abi = strings.TrimSpace(data.Abi)

	// Get the method from the abi
	method, _, argTypes, err := getMethodFromAbi(data.Abi, data.Func)
	if err != nil {
		return "", err
	}

	// Prepare the arguments, which may be nested
	argStringValues, err := prepareContractCallArgs(data.Args)
	if err != nil {
		return "", err
	}

	spew.Dump("method", method)
	spew.Dump("enc", argStringValues)
	spew.Dump("argTypes", argTypes)

	// Encode the method call
	// TODO: pass argTypes so we dont have to decode again... we can just copy what we do inside of `ethcoder.AbiEncodeMethodCalldataFromStringValuesAny`
	// res, err := ethcoder.AbiEncodeMethodCalldataFromStringValuesAny(method, enc)
	// if err != nil {
	// 	return "", err
	// }

	// TODO: we need to fix up and update ethcoder.AbiEncodeMethodCalldata to upgrade the method parsing functions... etc...

	// TODO: we are calling ParseEventDef multiple times..

	argValues, err := ethcoder.ABIUnmarshalStringValuesAny(argTypes, argStringValues)
	if err != nil {
		return "", err
	}
	// return AbiEncodeMethodCalldata(methodExpr, argValues)

	var mabi abi.ABI
	var methodName string

	if len(data.Abi) > 0 && strings.Contains(data.Abi, "(") && data.Abi[len(data.Abi)-1] == ')' {
		abiSig, err := ethcoder.ParseABISignature(data.Abi)
		if err != nil {
			return "", err
		}
		mabi, err = ethcoder.EventDefToABI(abiSig, true)
		if err != nil {
			return "", err
		}
		methodName = abiSig.Name
		spew.Dump(mabi)
		spew.Dump(argValues)
	} else {
		mabi, err = abi.JSON(strings.NewReader(data.Abi))
		if err != nil {
			return "", err
		}
		methodName = data.Func
	}

	args, err := packableArgValues(mabi, methodName, argValues)
	if err != nil {
		return "", err
	}

	packed, err := mabi.Pack(methodName, args...)
	if err != nil {
		return "", err
	}

	return "0x" + common.Bytes2Hex(packed), nil
}

func packableArgValues(mabi abi.ABI, method string, argValues []any) ([]any, error) {
	m, ok := mabi.Methods[method]
	if !ok {
		return nil, errors.New("method not found in abi")
	}

	if len(m.Inputs) != len(argValues) {
		return nil, errors.New("method inputs length does not match arg values length")
	}

	fmt.Println("$$$$$$$$$$$$$$$$$$$")
	spew.Dump(m.Inputs)

	out := make([]any, len(argValues))

	for i, input := range m.Inputs {
		isTuple := false
		typ := input.Type.String()
		if len(typ) >= 2 && typ[0] == '(' && typ[len(typ)-1] == ')' {
			isTuple = true
		}

		if !isTuple {
			out[i] = argValues[i]
		} else {
			// build struct for the tuple, as that is what the geth abi encoder expects
			// NOTE: in future we could fork or modify it if we want to avoid the need for this,
			// as it means decoding tuples will be more intensive the necessary.

			spew.Dump(input)

			fields := []reflect.StructField{}

			v, ok := argValues[i].([]any)
			if !ok {
				vv, ok := argValues[i].([]string)
				if !ok {
					return nil, errors.New("tuple arg values must be an array")
				}
				v = make([]any, len(vv))
				for j, x := range vv {
					v[j] = x
				}
			}

			for j, vv := range v {
				fields = append(fields, reflect.StructField{
					Name: fmt.Sprintf("Name%d", j),
					Type: reflect.TypeOf(vv),
				})
			}

			structType := reflect.StructOf(fields)
			instance := reflect.New(structType).Elem()

			for j, vv := range v {
				instance.Field(j).Set(reflect.ValueOf(vv))
			}

			spew.Dump(instance.Interface())

			out[i] = instance.Interface()
		}
	}

	return out, nil
}

func prepareContractCallArgs(args []any) ([]any, error) {
	var err error
	out := make([]any, len(args))

	for i, arg := range args {
		switch arg := arg.(type) {
		case string, []string, []any:
			out[i] = arg

		case map[string]interface{}:
			nst := arg

			var funcName string
			if v, ok := nst["func"].(string); ok {
				funcName = v
			}

			args, ok := nst["args"].([]interface{})
			if !ok {
				return nil, fmt.Errorf("nested encode expects the 'args' field to be an array")
			}

			abi, ok := nst["abi"].(string)
			if !ok {
				return nil, fmt.Errorf("nested encode expects an 'abi' field")
			}

			out[i], err = EncodeContractCall(&contractCallType{
				Abi:  abi,
				Func: funcName,
				Args: args,
			})
			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("abi encoding fail due to invalid arg type, '%T'", arg)
		}
	}

	return out, nil
}

// The abi may be a:
// - already encoded method abi: transferFrom(address,address,uint256)
// - already encoded named method: transferFrom(address from,address to,uint256 val)
// - an array of function abis: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"
// - or a single function abi:  "{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fees\",\"type\":\"address\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}"
// And it must always return it encoded, like this:
// - transferFrom(address,address,uint256)
// making sure that the method matches the returned one
func getMethodFromAbi(abi string, method string) (string, []string, []string, error) {
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
		abiSig, err := ethcoder.ParseABISignature(abi)
		if err != nil {
			return "", nil, nil, err
		}
		return abiSig.Signature, abiSig.ArgNames, abiSig.ArgTypes, nil
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
			return "", nil, nil, err
		}
	} else {
		var singleAbi FunctionAbi
		if err := json.Unmarshal([]byte(abi), &singleAbi); err != nil {
			return "", nil, nil, err
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
			return method + "(" + strings.Join(paramTypes, ",") + ")", order, paramTypes, nil
		}
	}

	return "", nil, nil, fmt.Errorf("method not found in abi")
}
