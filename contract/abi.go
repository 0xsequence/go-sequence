package contract

import (
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

type ContractABI struct {
	*abi.ABI
	Name string
	Bin  []byte
}

func (c *ContractABI) Encode(method string, args ...interface{}) ([]byte, error) {
	m, ok := c.ABI.Methods[method]
	if !ok {
		return nil, fmt.Errorf("contract method %s not found", method)
	}
	input, err := m.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	input = append(m.ID, input...)
	return input, nil
}

func NewContractRegistry() *ContractRegistry {
	return &ContractRegistry{
		contracts: map[string]*ContractABI{},
		names:     []string{},
	}
}

type ContractRegistry struct {
	contracts map[string]*ContractABI
	names     []string // index of contract names in the map
}

func (c *ContractRegistry) Register(contractABI *ContractABI) {
	if c.contracts == nil {
		c.contracts = map[string]*ContractABI{}
	}
	c.contracts[contractABI.Name] = contractABI
	c.names = append(c.names, contractABI.Name)
	sort.Sort(sort.StringSlice(c.names))
}

func (s *ContractRegistry) RegisterJSON(name string, abiJSON string) {
	s.Register(&ContractABI{
		Name: name,
		ABI:  MustParseABI(abiJSON),
	})
}

func (c *ContractRegistry) ContractNames() []string {
	return c.names
}

func (c *ContractRegistry) GetContractABI(name string) *ContractABI {
	return c.contracts[name]
}

func (c *ContractRegistry) Encode(contractName, method string, args ...interface{}) ([]byte, error) {
	if c.contracts == nil {
		return nil, fmt.Errorf("contract registry cannot find contract %s", contractName)
	}
	contractABI, _ := c.contracts[contractName]
	if contractABI == nil {
		return nil, fmt.Errorf("contract registry cannot find contract %s", contractName)
	}
	return contractABI.Encode(method, args...)
}

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
