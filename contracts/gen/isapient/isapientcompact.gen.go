// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package isapient

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

// ISapientCompactMetaData contains all meta data concerning the ISapientCompact contract.
var ISapientCompactMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"recoverSapientSignatureCompact\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
}

// ISapientCompactABI is the input ABI used to generate the binding from.
// Deprecated: Use ISapientCompactMetaData.ABI instead.
var ISapientCompactABI = ISapientCompactMetaData.ABI

// ISapientCompact is an auto generated Go binding around an Ethereum contract.
type ISapientCompact struct {
	ISapientCompactCaller     // Read-only binding to the contract
	ISapientCompactTransactor // Write-only binding to the contract
	ISapientCompactFilterer   // Log filterer for contract events
}

// ISapientCompactCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISapientCompactCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientCompactTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISapientCompactTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientCompactFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISapientCompactFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientCompactSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISapientCompactSession struct {
	Contract     *ISapientCompact  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISapientCompactCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISapientCompactCallerSession struct {
	Contract *ISapientCompactCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISapientCompactTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISapientCompactTransactorSession struct {
	Contract     *ISapientCompactTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISapientCompactRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISapientCompactRaw struct {
	Contract *ISapientCompact // Generic contract binding to access the raw methods on
}

// ISapientCompactCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISapientCompactCallerRaw struct {
	Contract *ISapientCompactCaller // Generic read-only contract binding to access the raw methods on
}

// ISapientCompactTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISapientCompactTransactorRaw struct {
	Contract *ISapientCompactTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISapientCompact creates a new instance of ISapientCompact, bound to a specific deployed contract.
func NewISapientCompact(address common.Address, backend bind.ContractBackend) (*ISapientCompact, error) {
	contract, err := bindISapientCompact(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISapientCompact{ISapientCompactCaller: ISapientCompactCaller{contract: contract}, ISapientCompactTransactor: ISapientCompactTransactor{contract: contract}, ISapientCompactFilterer: ISapientCompactFilterer{contract: contract}}, nil
}

// NewISapientCompactCaller creates a new read-only instance of ISapientCompact, bound to a specific deployed contract.
func NewISapientCompactCaller(address common.Address, caller bind.ContractCaller) (*ISapientCompactCaller, error) {
	contract, err := bindISapientCompact(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISapientCompactCaller{contract: contract}, nil
}

// NewISapientCompactTransactor creates a new write-only instance of ISapientCompact, bound to a specific deployed contract.
func NewISapientCompactTransactor(address common.Address, transactor bind.ContractTransactor) (*ISapientCompactTransactor, error) {
	contract, err := bindISapientCompact(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISapientCompactTransactor{contract: contract}, nil
}

// NewISapientCompactFilterer creates a new log filterer instance of ISapientCompact, bound to a specific deployed contract.
func NewISapientCompactFilterer(address common.Address, filterer bind.ContractFilterer) (*ISapientCompactFilterer, error) {
	contract, err := bindISapientCompact(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISapientCompactFilterer{contract: contract}, nil
}

// bindISapientCompact binds a generic wrapper to an already deployed contract.
func bindISapientCompact(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISapientCompactMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISapientCompact *ISapientCompactRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISapientCompact.Contract.ISapientCompactCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISapientCompact *ISapientCompactRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISapientCompact.Contract.ISapientCompactTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISapientCompact *ISapientCompactRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISapientCompact.Contract.ISapientCompactTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISapientCompact *ISapientCompactCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISapientCompact.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISapientCompact *ISapientCompactTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISapientCompact.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISapientCompact *ISapientCompactTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISapientCompact.Contract.contract.Transact(opts, method, params...)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 digest, bytes signature) view returns(bytes32 imageHash)
func (_ISapientCompact *ISapientCompactCaller) RecoverSapientSignatureCompact(opts *bind.CallOpts, digest [32]byte, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _ISapientCompact.contract.Call(opts, &out, "recoverSapientSignatureCompact", digest, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 digest, bytes signature) view returns(bytes32 imageHash)
func (_ISapientCompact *ISapientCompactSession) RecoverSapientSignatureCompact(digest [32]byte, signature []byte) ([32]byte, error) {
	return _ISapientCompact.Contract.RecoverSapientSignatureCompact(&_ISapientCompact.CallOpts, digest, signature)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 digest, bytes signature) view returns(bytes32 imageHash)
func (_ISapientCompact *ISapientCompactCallerSession) RecoverSapientSignatureCompact(digest [32]byte, signature []byte) ([32]byte, error) {
	return _ISapientCompact.Contract.RecoverSapientSignatureCompact(&_ISapientCompact.CallOpts, digest, signature)
}
