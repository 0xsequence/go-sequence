// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletfactory

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

// WalletFactoryABI is the input ABI used to generate the binding from.
const WalletFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainModule\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// WalletFactoryBin is the compiled bytecode used for deploying new contracts.
const WalletFactoryBin = "0x608060405234801561001057600080fd5b506101c8806100206000396000f3fe60806040526004361061001e5760003560e01c806332c02a1414610023575b600080fd5b61005c6004803603604081101561003957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610085565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60008060405180606001604052806028815260200161016b602891398473ffffffffffffffffffffffffffffffffffffffff166040516020018083805190602001908083835b6020831061010857805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016100cb565b51815160209384036101000a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0180199092169116179052920193845250604080518085038152938201905282519294508693508401905034f594935050505056fe603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3a26469706673582212209b0bce93afab3297b9ebf4e58fa642ef123d74bcbd3bdb4e48b662eb12b430ca64736f6c63430007060033"

// DeployWalletFactory deploys a new Ethereum contract, binding an instance of WalletFactory to it.
func DeployWalletFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletFactoryBin), backend)
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
	parsed, err := abi.JSON(strings.NewReader(WalletFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
