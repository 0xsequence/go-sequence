package intents

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDelayedABI(t *testing.T) {
	// Encode simple transferFrom, not named
	res, err := EncodeDelayedABI(&delayedEncodeType{
		Abi:  `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`,
		Func: "transferFrom",
		Args: json.RawMessage(`["0x0dc9603d4da53841C1C83f3B550C6143e60e0425","0x0dc9603d4da53841C1C83f3B550C6143e60e0425","100"]`),
	})

	require.Nil(t, err)
	require.Equal(t, "0x23b872dd0000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000000000000000000000000000000000000000064", res)

	// Encode simple transferFrom, named
	res, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `[{"name":"transferFrom","type":"function","inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}]}]`,
		Func: "transferFrom",
		Args: json.RawMessage(`{"_from": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425", "_value": "100", "_to": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425"}`),
	})

	require.Nil(t, err)
	require.Equal(t, "0x23b872dd0000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000dc9603d4da53841c1c83f3b550c6143e60e04250000000000000000000000000000000000000000000000000000000000000064", res)

	// Encode simple transferFrom, not named, passed as function
	res, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `transferFrom(address,address,uint256)`,
		Func: "transferFrom",
		Args: json.RawMessage(`["0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25","0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462","5192381927398174182391237"]`),
	})

	require.Nil(t, err)
	require.Equal(t, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d6462000000000000000000000000000000000000000000044b87969b06250e50bdc5", res)

	// Encode simple transferFrom, named, passed as function
	res, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `transferFrom(address _from,address _to,uint256 _value)`,
		Func: "transferFrom",
		Args: json.RawMessage(`{"_from": "0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "_value": "5192381927398174182391237", "_to": "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462"}`),
	})

	require.Nil(t, err)
	require.Equal(t, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d6462000000000000000000000000000000000000000000044b87969b06250e50bdc5", res)

	// Encode nested bytes, passed as function

	// TODO/NOTE: this makes no sense, we are making up an abi type of transferFrom just to encode a uint256
	nestedEncodeType1 := &delayedEncodeType{
		Abi:  `transferFrom(uint256)`,
		Func: "transferFrom",
		Args: json.RawMessage(`["481923749816926378123"]`),
	}

	// TODO/NOTE: this makes no sense, we are making up a random abi just so we can encode a string..
	nestedEncodeType2 := &delayedEncodeType{
		Abi:  `hola(string)`,
		Func: "hola",
		Args: json.RawMessage(`["mundo"]`),
	}

	net1jsn, err := json.Marshal(nestedEncodeType1)
	require.Nil(t, err)

	net2jsn, err := json.Marshal(nestedEncodeType2)
	require.Nil(t, err)

	res, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `[{"inputs":[{"internalType":"address","name":"addr","type":"address"},{"internalType":"bytes","name":"_arg1","type":"bytes"},{"internalType":"bytes","name":"_arg2","type":"bytes"}],"name":"caller","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"_orderId","type":"bytes32"},{"internalType":"uint256","name":"_maxCost","type":"uint256"},{"internalType":"address[]","name":"_fees","type":"address[]"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"fillOrKillOrder","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"_val","type":"uint256"},{"internalType":"string","name":"_data","type":"string"}],"name":"notExpired","outputs":[],"stateMutability":"view","type":"function"},{"inputs":[],"name":"s","outputs":[],"stateMutability":"nonpayable","type":"function"}]`,
		Func: "caller",
		Args: json.RawMessage(`{"addr": "0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "_arg1": ` + string(net1jsn) + `, "_arg2": ` + string(net2jsn) + `}`),
	})
	require.Nil(t, err)
	require.Equal(t, res, "0x8b6701df00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf25000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000002477a11f7e00000000000000000000000000000000000000000000001a2009191df61e988b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000646ce8ea55000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000056d756e646f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	// Fail passing named args to non-named abi
	_, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `transferFrom(address,uint256)`,
		Func: "transferFrom",
		Args: json.RawMessage(`{"_from": "0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "_value": "5192381927398174182391237", "_to": "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462"}`),
	})
	require.NotNil(t, err) // error is expected

	// Accept passing ordened args to named abi
	res, err = EncodeDelayedABI(&delayedEncodeType{
		Abi:  `transferFrom(address _from,address _to,uint256 _value)`,
		Func: "transferFrom",
		Args: json.RawMessage(`["0x13915b1ea28Fd2E8197c88ff9D2422182E83bf25", "0x4Ad47F1611c78C824Ff3892c4aE1CC04637D6462", "9"]`),
	})
	require.Nil(t, err)
	require.Equal(t, "0x23b872dd00000000000000000000000013915b1ea28fd2e8197c88ff9d2422182e83bf250000000000000000000000004ad47f1611c78c824ff3892c4ae1cc04637d64620000000000000000000000000000000000000000000000000000000000000009", res)
}
