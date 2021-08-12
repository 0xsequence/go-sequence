// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IWrapAndNiftyswapABI is the input ABI used to generate the binding from.
const IWrapAndNiftyswapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_niftyswapOrder\",\"type\":\"bytes\"}],\"name\":\"wrapAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWrapAndNiftyswap is an auto generated Go binding around an Ethereum contract.
type IWrapAndNiftyswap struct {
	IWrapAndNiftyswapCaller     // Read-only binding to the contract
	IWrapAndNiftyswapTransactor // Write-only binding to the contract
	IWrapAndNiftyswapFilterer   // Log filterer for contract events
}

// IWrapAndNiftyswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWrapAndNiftyswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapAndNiftyswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWrapAndNiftyswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapAndNiftyswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWrapAndNiftyswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWrapAndNiftyswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWrapAndNiftyswapSession struct {
	Contract     *IWrapAndNiftyswap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IWrapAndNiftyswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWrapAndNiftyswapCallerSession struct {
	Contract *IWrapAndNiftyswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IWrapAndNiftyswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWrapAndNiftyswapTransactorSession struct {
	Contract     *IWrapAndNiftyswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IWrapAndNiftyswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWrapAndNiftyswapRaw struct {
	Contract *IWrapAndNiftyswap // Generic contract binding to access the raw methods on
}

// IWrapAndNiftyswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWrapAndNiftyswapCallerRaw struct {
	Contract *IWrapAndNiftyswapCaller // Generic read-only contract binding to access the raw methods on
}

// IWrapAndNiftyswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWrapAndNiftyswapTransactorRaw struct {
	Contract *IWrapAndNiftyswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWrapAndNiftyswap creates a new instance of IWrapAndNiftyswap, bound to a specific deployed contract.
func NewIWrapAndNiftyswap(address common.Address, backend bind.ContractBackend) (*IWrapAndNiftyswap, error) {
	contract, err := bindIWrapAndNiftyswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWrapAndNiftyswap{IWrapAndNiftyswapCaller: IWrapAndNiftyswapCaller{contract: contract}, IWrapAndNiftyswapTransactor: IWrapAndNiftyswapTransactor{contract: contract}, IWrapAndNiftyswapFilterer: IWrapAndNiftyswapFilterer{contract: contract}}, nil
}

// NewIWrapAndNiftyswapCaller creates a new read-only instance of IWrapAndNiftyswap, bound to a specific deployed contract.
func NewIWrapAndNiftyswapCaller(address common.Address, caller bind.ContractCaller) (*IWrapAndNiftyswapCaller, error) {
	contract, err := bindIWrapAndNiftyswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWrapAndNiftyswapCaller{contract: contract}, nil
}

// NewIWrapAndNiftyswapTransactor creates a new write-only instance of IWrapAndNiftyswap, bound to a specific deployed contract.
func NewIWrapAndNiftyswapTransactor(address common.Address, transactor bind.ContractTransactor) (*IWrapAndNiftyswapTransactor, error) {
	contract, err := bindIWrapAndNiftyswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWrapAndNiftyswapTransactor{contract: contract}, nil
}

// NewIWrapAndNiftyswapFilterer creates a new log filterer instance of IWrapAndNiftyswap, bound to a specific deployed contract.
func NewIWrapAndNiftyswapFilterer(address common.Address, filterer bind.ContractFilterer) (*IWrapAndNiftyswapFilterer, error) {
	contract, err := bindIWrapAndNiftyswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWrapAndNiftyswapFilterer{contract: contract}, nil
}

// bindIWrapAndNiftyswap binds a generic wrapper to an already deployed contract.
func bindIWrapAndNiftyswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWrapAndNiftyswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrapAndNiftyswap.Contract.IWrapAndNiftyswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.IWrapAndNiftyswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.IWrapAndNiftyswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWrapAndNiftyswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.contract.Transact(opts, "onERC1155BatchReceived", arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.OnERC1155BatchReceived(&_IWrapAndNiftyswap.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactorSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.OnERC1155BatchReceived(&_IWrapAndNiftyswap.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactor) OnERC1155Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.contract.Transact(opts, "onERC1155Received", _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.OnERC1155Received(&_IWrapAndNiftyswap.TransactOpts, _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactorSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.OnERC1155Received(&_IWrapAndNiftyswap.TransactOpts, _operator, _from, _id, _amount, _data)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactor) WrapAndSwap(opts *bind.TransactOpts, _maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.contract.Transact(opts, "wrapAndSwap", _maxAmount, _recipient, _niftyswapOrder)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_IWrapAndNiftyswap *IWrapAndNiftyswapSession) WrapAndSwap(_maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.WrapAndSwap(&_IWrapAndNiftyswap.TransactOpts, _maxAmount, _recipient, _niftyswapOrder)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_IWrapAndNiftyswap *IWrapAndNiftyswapTransactorSession) WrapAndSwap(_maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _IWrapAndNiftyswap.Contract.WrapAndSwap(&_IWrapAndNiftyswap.TransactOpts, _maxAmount, _recipient, _niftyswapOrder)
}
