package intents

import (
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMethodFromABI(t *testing.T) {
	// From ABI, alone, in array
	res, argNames, argTypes, err := getMethodFromAbi(`[{"name":"transfer","type":"function","inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`, "transfer")
	assert.Nil(t, err)

	require.Equal(t, "transfer(address,uint256)", res)
	require.Equal(t, []string{"_to", "_value"}, argNames)
	require.Equal(t, []string{"address", "uint256"}, argTypes)

	// From ABI, alone, as object
	res, order, _, err := getMethodFromAbi(`{"name":"transfer","type":"function","inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}`, "transfer")
	assert.Nil(t, err)

	assert.Equal(t, "transfer(address,uint256)", res)
	assert.Equal(t, []string{"_to", "_value"}, order)

	// From ABI, with many args
	res, order, _, err = getMethodFromAbi(`[{"inputs":[{"internalType":"bytes32","name":"_orderId","type":"bytes32"},{"internalType":"uint256","name":"_maxCost","type":"uint256"},{"internalType":"address[]","name":"_fees","type":"address[]"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"fillOrKillOrder","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_val","type":"uint256"},{"internalType":"string","name":"_data","type":"string"}],"name":"notExpired","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[],"name":"otherMethods","outputs":[],"stateMutability":"nonpayable","type":"function"}]`, "fillOrKillOrder")
	assert.Nil(t, err)

	assert.Equal(t, "fillOrKillOrder(bytes32,uint256,address[],bytes)", res)
	assert.Equal(t, []string{"_orderId", "_maxCost", "_fees", "_data"}, order)

	res, order, _, err = getMethodFromAbi(`[{"inputs":[{"internalType":"bytes32","name":"_orderId","type":"bytes32"},{"internalType":"uint256","name":"_maxCost","type":"uint256"},{"internalType":"address[]","name":"_fees","type":"address[]"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"fillOrKillOrder","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_val","type":"uint256"},{"internalType":"string","name":"_data","type":"string"}],"name":"notExpired","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[],"name":"otherMethods","outputs":[],"stateMutability":"nonpayable","type":"function"}]`, "notExpired")
	assert.Nil(t, err)

	assert.Equal(t, "notExpired(uint256,string)", res)
	assert.Equal(t, []string{"_val", "_data"}, order)

	res, order, _, err = getMethodFromAbi(`[{"inputs":[{"internalType":"bytes32","name":"_orderId","type":"bytes32"},{"internalType":"uint256","name":"_maxCost","type":"uint256"},{"internalType":"address[]","name":"_fees","type":"address[]"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"fillOrKillOrder","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_val","type":"uint256"},{"internalType":"string","name":"_data","type":"string"}],"name":"notExpired","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[],"name":"otherMethods","outputs":[],"stateMutability":"nonpayable","type":"function"}]`, "otherMethods")
	assert.Nil(t, err)

	assert.Equal(t, "otherMethods()", res)
	assert.Equal(t, []string{}, order)

	// From plain method, without named args
	res, order, _, err = getMethodFromAbi(`transfer(address,uint256)`, "transfer")
	assert.Nil(t, err)

	assert.Equal(t, "transfer(address,uint256)", res)
	assert.Equal(t, []string{"arg1", "arg2"}, order)

	// From plain method, with named args
	res, order, _, err = getMethodFromAbi(`transfer(address _to,uint256 _value, bytes _mas)`, "transfer")
	assert.Nil(t, err)

	assert.Equal(t, "transfer(address,uint256,bytes)", res)
	assert.Equal(t, []string{"_to", "_value", "_mas"}, order)

	// Mixed plain method should return nil order
	res, order, _, err = getMethodFromAbi(`transfer(address _to,uint256, bytes _mas)`, "transfer")

	assert.Nil(t, err)
	assert.Equal(t, "transfer(address,uint256,bytes)", res)
	assert.Equal(t, []string{"_to", "arg2", "_mas"}, order)
}

func TestEncodeContractCall(t *testing.T) {
	// Encode simple transferFrom, not named
	res, err := EncodeContractCall(&contractCallType{
		Abi:  `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`,
		Func: "transferFrom",
		Args: []any{"0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "100"},
	})
	require.Nil(t, err)
	require.Equal(t, "0x23b872dd0000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000000000000000000000000000000000000000064", res)

	// Encode simple transferFrom, selector
	res, err = EncodeContractCall(&contractCallType{
		Abi:  `transferFrom(address,address,uint256)`,
		Args: []any{"0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "100"},
	})
	require.Nil(t, err)
	require.Equal(t, "0x23b872dd0000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000000000000000000000000000000000000000064", res)

	// Encode simple transferFrom, named
	// res, err = EncodeContractCall(&contractCallType{
	// 	Abi:  `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`,
	// 	Func: "transferFrom",
	// 	Args: json.RawMessage(`{"_from": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "_value": "100", "_to": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425"}`),
	// })
	// require.Nil(t, err)
	// require.Equal(t, res, "0x23b872dd0000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000000000000000000000000000000000000000064")

	// Encode simple transferFrom, not named, passed as function
	res, err = EncodeContractCall(&contractCallType{
		Abi:  `transferFrom(address,address,uint256)`,
		Args: []any{"0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462", "5192381927398174182391237"},
	})
	require.Nil(t, err)
	require.Equal(t, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d6462000000000000000000000000000000000000000000044b87969b06250e50bdc5", res)

	// Encode simple transferFrom, named, passed as function
	// res, err = EncodeContractCall(&contractCallType{
	// 	Abi:  `transferFrom(address _from,address _to,uint256 _value)`,
	// 	Func: "transferFrom",
	// 	Args: json.RawMessage(`{"_from": "0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "_value": "5192381927398174182391237", "_to": "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462"}`),
	// })
	// require.Nil(t, err)
	// require.Equal(t, res, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d6462000000000000000000000000000000000000000000044b87969b06250e50bdc5")

	// // Encode nested bytes, passed as function
	// nestedEncodeType1 := &contractCallType{
	// 	Abi:  `transferFrom(uint256)`,
	// 	Args: []any{"481923749816926378123"},
	// }

	// nestedEncodeType2 := &contractCallType{
	// 	Abi:  `hola(string)`,
	// 	Args: []any{"mundo"},
	// }

	// net1jsn, err := json.Marshal(nestedEncodeType1)
	// require.Nil(t, err)

	// net2jsn, err := json.Marshal(nestedEncodeType2)
	// require.Nil(t, err)

	arg2, err := ethcoder.AbiEncodeMethodCalldataFromStringValues("transferFrom(uint256)", []string{"481923749816926378123"})
	require.NoError(t, err)

	arg3, err := ethcoder.AbiEncodeMethodCalldataFromStringValues("hola(string)", []string{"mundo"})
	require.NoError(t, err)

	arg2Hex := "0x" + common.Bytes2Hex(arg2)
	arg3Hex := "0x" + common.Bytes2Hex(arg3)

	res, err = EncodeContractCall(&contractCallType{
		Abi:  "caller(address,bytes,bytes)",
		Func: "caller",
		Args: []any{"0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", arg2Hex, arg3Hex},
	})
	require.Nil(t, err)
	require.Equal(t, "0x8b6701df00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf25000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000002477a11f7e00000000000000000000000000000000000000000000001a2009191df61e988b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000646ce8ea55000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000056d756e646f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", res)

	// Fail passing named args to non-named abi
	// res, err = EncodeContractCall(&contractCallType{
	// 	Abi:  `transferFrom(address,uint256)`,
	// 	Func: "transferFrom",
	// 	Args: json.RawMessage(`{"_from": "0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "_value": "5192381927398174182391237", "_to": "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462"}`),
	// })
	// assert.NotNil(t, err)

	// Accept passing ordened args to named abi
	res, err = EncodeContractCall(&contractCallType{
		Abi:  `transferFrom(address _from,address _to,uint256 _value)`,
		Func: "transferFrom",
		Args: []any{"0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462", "9"},
	})
	require.Nil(t, err)
	require.Equal(t, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d64620000000000000000000000000000000000000000000000000000000000000009", res)

	// ...
	res, err = EncodeContractCall(&contractCallType{
		Abi: `fillOrder(uint256 orderId, uint256 maxCost, address[] fees, (uint256 a, address b) extra)`,
		Args: []any{
			"48774435471364917511246724398022004900255301025912680232738918790354204737320",
			"1000000000000000000",
			[]string{"0x8541D65829f98f7D71A4655cCD7B2bB8494673bF"},
			[]string{"123456789", "0x1231f65f29f98e7D71A4655cCD7B2bc441211feb"},
		},
	})
	require.Nil(t, err)
	require.Equal(t, "0x326a62086bd55a2877890bd58871eefe886770a7734077a74981910a75d7b1f044b5bf280000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000075bcd150000000000000000000000001231f65f29f98e7d71a4655ccd7b2bc441211feb00000000000000000000000000000000000000000000000000000000000000010000000000000000000000008541d65829f98f7d71a4655ccd7b2bb8494673bf", res)
}
