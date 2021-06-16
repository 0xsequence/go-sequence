package sequence

import (
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

// TODO: contract helpers to encode, etc.....?

// .. like artifacts, we have Name, ABI and Bin..
// we can encode calls here to slam into a txn.. etc..

// NewContractTransaction() .. uses NewTransaction() below..

// NewTransaction() -- more generic, pass your own data

// NewContractCall() -- lets us read data, ..

type ContractABI struct {
	Name string
	ABI  *abi.ABI
	Bin  []byte
}

func (c *ContractABI) Call(method string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

// type Contract

func MustNewType(str string) abi.Type {
	typ, err := abi.NewType(str, "", nil)
	if err != nil {
		panic(err)
	}
	return typ
}

func MustNewArrayTypeTuple(components []abi.ArgumentMarshaling) abi.Type {
	typ, err := abi.NewType("tuple[]", "", components)
	if err != nil {
		panic(err)
	}
	return typ
}

func MustParseABI(abiJSON string) *abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic("failed to parse abi json")
	}
	return &parsed
}
