// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletcallmock

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
)

// CallReceiverMockMetaData contains all meta data concerning the CallReceiverMock contract.
var CallReceiverMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"lastValA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValB\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_revertFlag\",\"type\":\"bool\"}],\"name\":\"setRevertFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_valA\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_valB\",\"type\":\"bytes\"}],\"name\":\"testCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052610560806100136000396000f3fe60806040526004361061003f5760003560e01c8063381ba14014610044578063ad387c8a14610092578063c0aec4d3146100a5578063ebd35e47146100ce575b600080fd5b34801561005057600080fd5b5061009061005f36600461022d565b600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b005b6100906100a0366004610256565b6100f0565b3480156100b157600080fd5b506100bb60005481565b6040519081526020015b60405180910390f35b3480156100da57600080fd5b506100e361019f565b6040516100c591906102d2565b60025460ff1615610187576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f43616c6c52656365697665724d6f636b237465737443616c6c3a20524556455260448201527f545f464c41470000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b6000839055600161019982848361040f565b50505050565b600180546101ac9061036d565b80601f01602080910402602001604051908101604052809291908181526020018280546101d89061036d565b80156102255780601f106101fa57610100808354040283529160200191610225565b820191906000526020600020905b81548152906001019060200180831161020857829003601f168201915b505050505081565b60006020828403121561023f57600080fd5b8135801515811461024f57600080fd5b9392505050565b60008060006040848603121561026b57600080fd5b83359250602084013567ffffffffffffffff8082111561028a57600080fd5b818601915086601f83011261029e57600080fd5b8135818111156102ad57600080fd5b8760208285010111156102bf57600080fd5b6020830194508093505050509250925092565b600060208083528351808285015260005b818110156102ff578581018301518582016040015282016102e3565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c9082168061038157607f821691505b6020821081036103ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561040a57600081815260208120601f850160051c810160208610156103e75750805b601f850160051c820191505b81811015610406578281556001016103f3565b5050505b505050565b67ffffffffffffffff8311156104275761042761033e565b61043b83610435835461036d565b836103c0565b6000601f84116001811461048d57600085156104575750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355610523565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b828110156104dc57868501358255602094850194600190920191016104bc565b5086821015610517577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b505050505056fea2646970667358221220957421523c943545e31f979350312047b130f797407f055fb4eff4c262a692dd64736f6c63430008120033",
}

// CallReceiverMockABI is the input ABI used to generate the binding from.
// Deprecated: Use CallReceiverMockMetaData.ABI instead.
var CallReceiverMockABI = CallReceiverMockMetaData.ABI

// CallReceiverMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CallReceiverMockMetaData.Bin instead.
var CallReceiverMockBin = CallReceiverMockMetaData.Bin

// DeployCallReceiverMock deploys a new Ethereum contract, binding an instance of CallReceiverMock to it.
func DeployCallReceiverMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallReceiverMock, error) {
	parsed, err := CallReceiverMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CallReceiverMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallReceiverMock{CallReceiverMockCaller: CallReceiverMockCaller{contract: contract}, CallReceiverMockTransactor: CallReceiverMockTransactor{contract: contract}, CallReceiverMockFilterer: CallReceiverMockFilterer{contract: contract}}, nil
}

// CallReceiverMock is an auto generated Go binding around an Ethereum contract.
type CallReceiverMock struct {
	CallReceiverMockCaller     // Read-only binding to the contract
	CallReceiverMockTransactor // Write-only binding to the contract
	CallReceiverMockFilterer   // Log filterer for contract events
}

// CallReceiverMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallReceiverMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallReceiverMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallReceiverMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallReceiverMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallReceiverMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallReceiverMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallReceiverMockSession struct {
	Contract     *CallReceiverMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallReceiverMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallReceiverMockCallerSession struct {
	Contract *CallReceiverMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CallReceiverMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallReceiverMockTransactorSession struct {
	Contract     *CallReceiverMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CallReceiverMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallReceiverMockRaw struct {
	Contract *CallReceiverMock // Generic contract binding to access the raw methods on
}

// CallReceiverMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallReceiverMockCallerRaw struct {
	Contract *CallReceiverMockCaller // Generic read-only contract binding to access the raw methods on
}

// CallReceiverMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallReceiverMockTransactorRaw struct {
	Contract *CallReceiverMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallReceiverMock creates a new instance of CallReceiverMock, bound to a specific deployed contract.
func NewCallReceiverMock(address common.Address, backend bind.ContractBackend) (*CallReceiverMock, error) {
	contract, err := bindCallReceiverMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallReceiverMock{CallReceiverMockCaller: CallReceiverMockCaller{contract: contract}, CallReceiverMockTransactor: CallReceiverMockTransactor{contract: contract}, CallReceiverMockFilterer: CallReceiverMockFilterer{contract: contract}}, nil
}

// NewCallReceiverMockCaller creates a new read-only instance of CallReceiverMock, bound to a specific deployed contract.
func NewCallReceiverMockCaller(address common.Address, caller bind.ContractCaller) (*CallReceiverMockCaller, error) {
	contract, err := bindCallReceiverMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallReceiverMockCaller{contract: contract}, nil
}

// NewCallReceiverMockTransactor creates a new write-only instance of CallReceiverMock, bound to a specific deployed contract.
func NewCallReceiverMockTransactor(address common.Address, transactor bind.ContractTransactor) (*CallReceiverMockTransactor, error) {
	contract, err := bindCallReceiverMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallReceiverMockTransactor{contract: contract}, nil
}

// NewCallReceiverMockFilterer creates a new log filterer instance of CallReceiverMock, bound to a specific deployed contract.
func NewCallReceiverMockFilterer(address common.Address, filterer bind.ContractFilterer) (*CallReceiverMockFilterer, error) {
	contract, err := bindCallReceiverMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallReceiverMockFilterer{contract: contract}, nil
}

// bindCallReceiverMock binds a generic wrapper to an already deployed contract.
func bindCallReceiverMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallReceiverMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallReceiverMock *CallReceiverMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallReceiverMock.Contract.CallReceiverMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallReceiverMock *CallReceiverMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.CallReceiverMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallReceiverMock *CallReceiverMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.CallReceiverMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallReceiverMock *CallReceiverMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallReceiverMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallReceiverMock *CallReceiverMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallReceiverMock *CallReceiverMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.contract.Transact(opts, method, params...)
}

// LastValA is a free data retrieval call binding the contract method 0xc0aec4d3.
//
// Solidity: function lastValA() view returns(uint256)
func (_CallReceiverMock *CallReceiverMockCaller) LastValA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CallReceiverMock.contract.Call(opts, &out, "lastValA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValA is a free data retrieval call binding the contract method 0xc0aec4d3.
//
// Solidity: function lastValA() view returns(uint256)
func (_CallReceiverMock *CallReceiverMockSession) LastValA() (*big.Int, error) {
	return _CallReceiverMock.Contract.LastValA(&_CallReceiverMock.CallOpts)
}

// LastValA is a free data retrieval call binding the contract method 0xc0aec4d3.
//
// Solidity: function lastValA() view returns(uint256)
func (_CallReceiverMock *CallReceiverMockCallerSession) LastValA() (*big.Int, error) {
	return _CallReceiverMock.Contract.LastValA(&_CallReceiverMock.CallOpts)
}

// LastValB is a free data retrieval call binding the contract method 0xebd35e47.
//
// Solidity: function lastValB() view returns(bytes)
func (_CallReceiverMock *CallReceiverMockCaller) LastValB(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _CallReceiverMock.contract.Call(opts, &out, "lastValB")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LastValB is a free data retrieval call binding the contract method 0xebd35e47.
//
// Solidity: function lastValB() view returns(bytes)
func (_CallReceiverMock *CallReceiverMockSession) LastValB() ([]byte, error) {
	return _CallReceiverMock.Contract.LastValB(&_CallReceiverMock.CallOpts)
}

// LastValB is a free data retrieval call binding the contract method 0xebd35e47.
//
// Solidity: function lastValB() view returns(bytes)
func (_CallReceiverMock *CallReceiverMockCallerSession) LastValB() ([]byte, error) {
	return _CallReceiverMock.Contract.LastValB(&_CallReceiverMock.CallOpts)
}

// SetRevertFlag is a paid mutator transaction binding the contract method 0x381ba140.
//
// Solidity: function setRevertFlag(bool _revertFlag) returns()
func (_CallReceiverMock *CallReceiverMockTransactor) SetRevertFlag(opts *bind.TransactOpts, _revertFlag bool) (*types.Transaction, error) {
	return _CallReceiverMock.contract.Transact(opts, "setRevertFlag", _revertFlag)
}

// SetRevertFlag is a paid mutator transaction binding the contract method 0x381ba140.
//
// Solidity: function setRevertFlag(bool _revertFlag) returns()
func (_CallReceiverMock *CallReceiverMockSession) SetRevertFlag(_revertFlag bool) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.SetRevertFlag(&_CallReceiverMock.TransactOpts, _revertFlag)
}

// SetRevertFlag is a paid mutator transaction binding the contract method 0x381ba140.
//
// Solidity: function setRevertFlag(bool _revertFlag) returns()
func (_CallReceiverMock *CallReceiverMockTransactorSession) SetRevertFlag(_revertFlag bool) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.SetRevertFlag(&_CallReceiverMock.TransactOpts, _revertFlag)
}

// TestCall is a paid mutator transaction binding the contract method 0xad387c8a.
//
// Solidity: function testCall(uint256 _valA, bytes _valB) payable returns()
func (_CallReceiverMock *CallReceiverMockTransactor) TestCall(opts *bind.TransactOpts, _valA *big.Int, _valB []byte) (*types.Transaction, error) {
	return _CallReceiverMock.contract.Transact(opts, "testCall", _valA, _valB)
}

// TestCall is a paid mutator transaction binding the contract method 0xad387c8a.
//
// Solidity: function testCall(uint256 _valA, bytes _valB) payable returns()
func (_CallReceiverMock *CallReceiverMockSession) TestCall(_valA *big.Int, _valB []byte) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.TestCall(&_CallReceiverMock.TransactOpts, _valA, _valB)
}

// TestCall is a paid mutator transaction binding the contract method 0xad387c8a.
//
// Solidity: function testCall(uint256 _valA, bytes _valB) payable returns()
func (_CallReceiverMock *CallReceiverMockTransactorSession) TestCall(_valA *big.Int, _valB []byte) (*types.Transaction, error) {
	return _CallReceiverMock.Contract.TestCall(&_CallReceiverMock.TransactOpts, _valA, _valB)
}
