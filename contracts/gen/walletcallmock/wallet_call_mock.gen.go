// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletcallmock

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CallReceiverMockABI is the input ABI used to generate the binding from.
const CallReceiverMockABI = "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"lastValA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValB\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_revertFlag\",\"type\":\"bool\"}],\"name\":\"setRevertFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_valA\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_valB\",\"type\":\"bytes\"}],\"name\":\"testCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CallReceiverMockBin is the compiled bytecode used for deploying new contracts.
const CallReceiverMockBin = "0x608060405261040a806100136000396000f3fe60806040526004361061003f5760003560e01c8063381ba14014610044578063ad387c8a14610072578063c0aec4d3146100e9578063ebd35e4714610110575b600080fd5b34801561005057600080fd5b506100706004803603602081101561006757600080fd5b5035151561019a565b005b6100706004803603604081101561008857600080fd5b813591908101906040810160208201356401000000008111156100aa57600080fd5b8201836020820111156100bc57600080fd5b803590602001918460018302840111640100000000831117156100de57600080fd5b5090925090506101cb565b3480156100f557600080fd5b506100fe61023e565b60408051918252519081900360200190f35b34801561011c57600080fd5b50610125610244565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015f578181015183820152602001610147565b50505050905090810190601f16801561018c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600280547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055565b60025460ff1615610227576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806103af6026913960400191505060405180910390fd5b6000839055610238600183836102ef565b50505050565b60005481565b60018054604080516020600284861615610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190941693909304601f810184900484028201840190925281815292918301828280156102e75780601f106102bc576101008083540402835291602001916102e7565b820191906000526020600020905b8154815290600101906020018083116102ca57829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826103255760008555610389565b82601f1061035c578280017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00823516178555610389565b82800160010185558215610389579182015b8281111561038957823582559160200191906001019061036e565b50610395929150610399565b5090565b5b80821115610395576000815560010161039a56fe43616c6c52656365697665724d6f636b237465737443616c6c3a205245564552545f464c4147a2646970667358221220b0a2a89a0c39b5c99762b2b511726513017a3442bd84964eaae31cb48b091c5764736f6c63430007060033"

// DeployCallReceiverMock deploys a new Ethereum contract, binding an instance of CallReceiverMock to it.
func DeployCallReceiverMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallReceiverMock, error) {
	parsed, err := abi.JSON(strings.NewReader(CallReceiverMockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallReceiverMockBin), backend)
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
