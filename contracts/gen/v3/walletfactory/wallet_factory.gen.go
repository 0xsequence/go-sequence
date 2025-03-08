// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletfactory

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

// WalletFactoryMetaData contains all meta data concerning the WalletFactory contract.
var WalletFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainModule\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"DeployFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainModule\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061020f8061001f6000396000f3fe60806040526004361061001e5760003560e01c806332c02a1414610023575b600080fd5b610036610031366004610136565b61005f565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6000806040518060600160405280602c81526020016101ae602c91398473ffffffffffffffffffffffffffffffffffffffff166040516020016100a392919061017b565b60405160208183030381529060405290508281516020830134f5915073ffffffffffffffffffffffffffffffffffffffff821661012f576040517f8caac80500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810184905260440160405180910390fd5b5092915050565b6000806040838503121561014957600080fd5b823573ffffffffffffffffffffffffffffffffffffffff8116811461016d57600080fd5b946020939093013593505050565b6000835160005b8181101561019c5760208187018101518583015201610182565b50919091019182525060200191905056fe603e600e3d39601e805130553df33d3d34601c57363d3d373d363d30545af43d82803e903d91601c57fd5bf3a26469706673582212200b013d50119518f1f732edb74583f8c2759568433b739459e61a09392636eaca64736f6c634300081b0033",
}

// WalletFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletFactoryMetaData.ABI instead.
var WalletFactoryABI = WalletFactoryMetaData.ABI

// WalletFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletFactoryMetaData.Bin instead.
var WalletFactoryBin = WalletFactoryMetaData.Bin

// DeployWalletFactory deploys a new Ethereum contract, binding an instance of WalletFactory to it.
func DeployWalletFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletFactory, error) {
	parsed, err := WalletFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletFactory{WalletFactoryCaller: WalletFactoryCaller{contract: contract}, WalletFactoryTransactor: WalletFactoryTransactor{contract: contract}, WalletFactoryFilterer: WalletFactoryFilterer{contract: contract}}, nil
}

// WalletFactory is an auto generated Go binding around an Ethereum contract.
type WalletFactory struct {
	WalletFactoryCaller     // Read-only binding to the contract
	WalletFactoryTransactor // Write-only binding to the contract
	WalletFactoryFilterer   // Log filterer for contract events
}

// WalletFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletFactorySession struct {
	Contract     *WalletFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletFactoryCallerSession struct {
	Contract *WalletFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletFactoryTransactorSession struct {
	Contract     *WalletFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletFactoryRaw struct {
	Contract *WalletFactory // Generic contract binding to access the raw methods on
}

// WalletFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletFactoryCallerRaw struct {
	Contract *WalletFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WalletFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletFactoryTransactorRaw struct {
	Contract *WalletFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletFactory creates a new instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactory(address common.Address, backend bind.ContractBackend) (*WalletFactory, error) {
	contract, err := bindWalletFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletFactory{WalletFactoryCaller: WalletFactoryCaller{contract: contract}, WalletFactoryTransactor: WalletFactoryTransactor{contract: contract}, WalletFactoryFilterer: WalletFactoryFilterer{contract: contract}}, nil
}

// NewWalletFactoryCaller creates a new read-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryCaller(address common.Address, caller bind.ContractCaller) (*WalletFactoryCaller, error) {
	contract, err := bindWalletFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryCaller{contract: contract}, nil
}

// NewWalletFactoryTransactor creates a new write-only instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletFactoryTransactor, error) {
	contract, err := bindWalletFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryTransactor{contract: contract}, nil
}

// NewWalletFactoryFilterer creates a new log filterer instance of WalletFactory, bound to a specific deployed contract.
func NewWalletFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFactoryFilterer, error) {
	contract, err := bindWalletFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFactoryFilterer{contract: contract}, nil
}

// bindWalletFactory binds a generic wrapper to an already deployed contract.
func bindWalletFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.WalletFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.WalletFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletFactory *WalletFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletFactory *WalletFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletFactory *WalletFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletFactory.Contract.contract.Transact(opts, method, params...)
}

// Deploy is a paid mutator transaction binding the contract method 0x32c02a14.
//
// Solidity: function deploy(address _mainModule, bytes32 _salt) payable returns(address _contract)
func (_WalletFactory *WalletFactoryTransactor) Deploy(opts *bind.TransactOpts, _mainModule common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _WalletFactory.contract.Transact(opts, "deploy", _mainModule, _salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x32c02a14.
//
// Solidity: function deploy(address _mainModule, bytes32 _salt) payable returns(address _contract)
func (_WalletFactory *WalletFactorySession) Deploy(_mainModule common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _WalletFactory.Contract.Deploy(&_WalletFactory.TransactOpts, _mainModule, _salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x32c02a14.
//
// Solidity: function deploy(address _mainModule, bytes32 _salt) payable returns(address _contract)
func (_WalletFactory *WalletFactoryTransactorSession) Deploy(_mainModule common.Address, _salt [32]byte) (*types.Transaction, error) {
	return _WalletFactory.Contract.Deploy(&_WalletFactory.TransactOpts, _mainModule, _salt)
}
