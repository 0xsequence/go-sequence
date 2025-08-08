// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package valueforwarder

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ValueForwarderMetaData contains all meta data concerning the ValueForwarder contract.
var ValueForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"forwardValue\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
	Bin: "0x6080806040523460155761014f908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c6398f850f11461002757600080fd5b6040366003190112610114576004356001600160a01b038116810361011457600080808093602435905af13d1561010f573d67ffffffffffffffff81116100f95760405190601f8101601f19908116603f0116820167ffffffffffffffff8111838210176100f9576040528152600060203d92013e5b156100a457005b60405162461bcd60e51b815260206004820152602760248201527f56616c7565466f727761726465723a204661696c656420746f20666f72776172604482015266642076616c756560c81b6064820152608490fd5b634e487b7160e01b600052604160045260246000fd5b61009d565b600080fdfea26469706673582212202706a11f313e042ccf1cd42b40b5f01308364a2290592fe4e14dce33175da1bf64736f6c634300081c0033",
}

// ValueForwarderABI is the input ABI used to generate the binding from.
// Deprecated: Use ValueForwarderMetaData.ABI instead.
var ValueForwarderABI = ValueForwarderMetaData.ABI

// ValueForwarderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValueForwarderMetaData.Bin instead.
var ValueForwarderBin = ValueForwarderMetaData.Bin

// DeployValueForwarder deploys a new Ethereum contract, binding an instance of ValueForwarder to it.
func DeployValueForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValueForwarder, error) {
	parsed, err := ValueForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValueForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValueForwarder{ValueForwarderCaller: ValueForwarderCaller{contract: contract}, ValueForwarderTransactor: ValueForwarderTransactor{contract: contract}, ValueForwarderFilterer: ValueForwarderFilterer{contract: contract}}, nil
}

// ValueForwarder is an auto generated Go binding around an Ethereum contract.
type ValueForwarder struct {
	ValueForwarderCaller     // Read-only binding to the contract
	ValueForwarderTransactor // Write-only binding to the contract
	ValueForwarderFilterer   // Log filterer for contract events
}

// ValueForwarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueForwarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueForwarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueForwarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueForwarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueForwarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueForwarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueForwarderSession struct {
	Contract     *ValueForwarder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueForwarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueForwarderCallerSession struct {
	Contract *ValueForwarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValueForwarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueForwarderTransactorSession struct {
	Contract     *ValueForwarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValueForwarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueForwarderRaw struct {
	Contract *ValueForwarder // Generic contract binding to access the raw methods on
}

// ValueForwarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueForwarderCallerRaw struct {
	Contract *ValueForwarderCaller // Generic read-only contract binding to access the raw methods on
}

// ValueForwarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueForwarderTransactorRaw struct {
	Contract *ValueForwarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValueForwarder creates a new instance of ValueForwarder, bound to a specific deployed contract.
func NewValueForwarder(address common.Address, backend bind.ContractBackend) (*ValueForwarder, error) {
	contract, err := bindValueForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValueForwarder{ValueForwarderCaller: ValueForwarderCaller{contract: contract}, ValueForwarderTransactor: ValueForwarderTransactor{contract: contract}, ValueForwarderFilterer: ValueForwarderFilterer{contract: contract}}, nil
}

// NewValueForwarderCaller creates a new read-only instance of ValueForwarder, bound to a specific deployed contract.
func NewValueForwarderCaller(address common.Address, caller bind.ContractCaller) (*ValueForwarderCaller, error) {
	contract, err := bindValueForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueForwarderCaller{contract: contract}, nil
}

// NewValueForwarderTransactor creates a new write-only instance of ValueForwarder, bound to a specific deployed contract.
func NewValueForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueForwarderTransactor, error) {
	contract, err := bindValueForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueForwarderTransactor{contract: contract}, nil
}

// NewValueForwarderFilterer creates a new log filterer instance of ValueForwarder, bound to a specific deployed contract.
func NewValueForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueForwarderFilterer, error) {
	contract, err := bindValueForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueForwarderFilterer{contract: contract}, nil
}

// bindValueForwarder binds a generic wrapper to an already deployed contract.
func bindValueForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValueForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueForwarder *ValueForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueForwarder.Contract.ValueForwarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueForwarder *ValueForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueForwarder.Contract.ValueForwarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueForwarder *ValueForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueForwarder.Contract.ValueForwarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValueForwarder *ValueForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValueForwarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValueForwarder *ValueForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValueForwarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValueForwarder *ValueForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValueForwarder.Contract.contract.Transact(opts, method, params...)
}

// ForwardValue is a paid mutator transaction binding the contract method 0x98f850f1.
//
// Solidity: function forwardValue(address to, uint256 value) payable returns()
func (_ValueForwarder *ValueForwarderTransactor) ForwardValue(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ValueForwarder.contract.Transact(opts, "forwardValue", to, value)
}

// ForwardValue is a paid mutator transaction binding the contract method 0x98f850f1.
//
// Solidity: function forwardValue(address to, uint256 value) payable returns()
func (_ValueForwarder *ValueForwarderSession) ForwardValue(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ValueForwarder.Contract.ForwardValue(&_ValueForwarder.TransactOpts, to, value)
}

// ForwardValue is a paid mutator transaction binding the contract method 0x98f850f1.
//
// Solidity: function forwardValue(address to, uint256 value) payable returns()
func (_ValueForwarder *ValueForwarderTransactorSession) ForwardValue(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ValueForwarder.Contract.ForwardValue(&_ValueForwarder.TransactOpts, to, value)
}
