package sequence

import (
	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

// abiTransactionsType represents abi coder of []Transaction
var abiTransactionsType = ethcoder.MustNewArrayTypeTuple([]abi.ArgumentMarshaling{
	{Name: "delegateCall", Type: "bool"},
	{Name: "revertOnError", Type: "bool"},
	{Name: "gasLimit", Type: "uint256"},
	{Name: "target", Type: "address"},
	{Name: "value", Type: "uint256"},
	{Name: "data", Type: "bytes"},
})

// abiTransactionsDigestType represents abi coder of []Transaction digest pre-image
var abiTransactionsDigestType = abi.Arguments{
	abi.Argument{Type: ethcoder.MustNewType("uint256")},
	abi.Argument{Type: abiTransactionsType},
}
