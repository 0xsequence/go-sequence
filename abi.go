package sequence

import (
	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

// TODO: duplicate in ethkit.. / move..?
func mustNewArrayTypeTuple(components []abi.ArgumentMarshaling) abi.Type {
	typ, err := abi.NewType("tuple[]", "", components)
	if err != nil {
		panic(err)
	}
	return typ
}

// TODO: rename
var ArrayOfMetaTxnType = mustNewArrayTypeTuple([]abi.ArgumentMarshaling{
	{
		Name: "delegateCall", Type: "bool",
	},
	{
		Name: "revertOnError", Type: "bool",
	},
	{
		Name: "gasLimit", Type: "uint256",
	},
	{
		Name: "target", Type: "address",
	},
	{
		Name: "value", Type: "uint256",
	},
	{
		Name: "data", Type: "bytes",
	},
})

// TODO: rename..
var nonceAndMetaTxns = abi.Arguments{
	abi.Argument{
		Type: ethcoder.MustNewType("uint256"),
	},
	abi.Argument{
		Type: ArrayOfMetaTxnType,
	},
}
