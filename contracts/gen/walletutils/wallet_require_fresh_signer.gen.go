// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletutils

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

// WalletRequireFreshSignerMetaData contains all meta data concerning the WalletRequireFreshSigner contract.
var WalletRequireFreshSignerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRequireUtils\",\"name\":\"_requireUtils\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"REQUIRE_UTILS\",\"outputs\":[{\"internalType\":\"contractRequireUtils\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"requireFreshSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516102aa3803806102aa8339818101604052602081101561003357600080fd5b5051606081901b6001600160601b0319166080526001600160a01b031661023f61006b6000398060a352806101af525061023f6000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630df0c4191461003b578063cfc63a4914610070575b600080fd5b61006e6004803603602081101561005157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100a1565b005b6100786101ad565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631cd05dc4826040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561012857600080fd5b505afa15801561013c573d6000803e3d6000fd5b505050506040513d602081101561015257600080fd5b5051156101aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260388152602001806101d26038913960400191505060405180910390fd5b50565b7f00000000000000000000000000000000000000000000000000000000000000008156fe5265717569726546726573685369676e6572237265717569726546726573685369676e65723a204455504c4943415445445f5349474e4552a2646970667358221220e659f104be18ad6f797d1c0a2939307f47f1ba93ae17f15e4cdb46b8942173f964736f6c63430007060033",
}

// WalletRequireFreshSignerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletRequireFreshSignerMetaData.ABI instead.
var WalletRequireFreshSignerABI = WalletRequireFreshSignerMetaData.ABI

// WalletRequireFreshSignerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletRequireFreshSignerMetaData.Bin instead.
var WalletRequireFreshSignerBin = WalletRequireFreshSignerMetaData.Bin

// DeployWalletRequireFreshSigner deploys a new Ethereum contract, binding an instance of WalletRequireFreshSigner to it.
func DeployWalletRequireFreshSigner(auth *bind.TransactOpts, backend bind.ContractBackend, _requireUtils common.Address) (common.Address, *types.Transaction, *WalletRequireFreshSigner, error) {
	parsed, err := WalletRequireFreshSignerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletRequireFreshSignerBin), backend, _requireUtils)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletRequireFreshSigner{WalletRequireFreshSignerCaller: WalletRequireFreshSignerCaller{contract: contract}, WalletRequireFreshSignerTransactor: WalletRequireFreshSignerTransactor{contract: contract}, WalletRequireFreshSignerFilterer: WalletRequireFreshSignerFilterer{contract: contract}}, nil
}

// WalletRequireFreshSigner is an auto generated Go binding around an Ethereum contract.
type WalletRequireFreshSigner struct {
	WalletRequireFreshSignerCaller     // Read-only binding to the contract
	WalletRequireFreshSignerTransactor // Write-only binding to the contract
	WalletRequireFreshSignerFilterer   // Log filterer for contract events
}

// WalletRequireFreshSignerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletRequireFreshSignerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletRequireFreshSignerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletRequireFreshSignerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletRequireFreshSignerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletRequireFreshSignerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletRequireFreshSignerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletRequireFreshSignerSession struct {
	Contract     *WalletRequireFreshSigner // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletRequireFreshSignerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletRequireFreshSignerCallerSession struct {
	Contract *WalletRequireFreshSignerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// WalletRequireFreshSignerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletRequireFreshSignerTransactorSession struct {
	Contract     *WalletRequireFreshSignerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// WalletRequireFreshSignerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRequireFreshSignerRaw struct {
	Contract *WalletRequireFreshSigner // Generic contract binding to access the raw methods on
}

// WalletRequireFreshSignerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletRequireFreshSignerCallerRaw struct {
	Contract *WalletRequireFreshSignerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletRequireFreshSignerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletRequireFreshSignerTransactorRaw struct {
	Contract *WalletRequireFreshSignerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletRequireFreshSigner creates a new instance of WalletRequireFreshSigner, bound to a specific deployed contract.
func NewWalletRequireFreshSigner(address common.Address, backend bind.ContractBackend) (*WalletRequireFreshSigner, error) {
	contract, err := bindWalletRequireFreshSigner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletRequireFreshSigner{WalletRequireFreshSignerCaller: WalletRequireFreshSignerCaller{contract: contract}, WalletRequireFreshSignerTransactor: WalletRequireFreshSignerTransactor{contract: contract}, WalletRequireFreshSignerFilterer: WalletRequireFreshSignerFilterer{contract: contract}}, nil
}

// NewWalletRequireFreshSignerCaller creates a new read-only instance of WalletRequireFreshSigner, bound to a specific deployed contract.
func NewWalletRequireFreshSignerCaller(address common.Address, caller bind.ContractCaller) (*WalletRequireFreshSignerCaller, error) {
	contract, err := bindWalletRequireFreshSigner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletRequireFreshSignerCaller{contract: contract}, nil
}

// NewWalletRequireFreshSignerTransactor creates a new write-only instance of WalletRequireFreshSigner, bound to a specific deployed contract.
func NewWalletRequireFreshSignerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletRequireFreshSignerTransactor, error) {
	contract, err := bindWalletRequireFreshSigner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletRequireFreshSignerTransactor{contract: contract}, nil
}

// NewWalletRequireFreshSignerFilterer creates a new log filterer instance of WalletRequireFreshSigner, bound to a specific deployed contract.
func NewWalletRequireFreshSignerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletRequireFreshSignerFilterer, error) {
	contract, err := bindWalletRequireFreshSigner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletRequireFreshSignerFilterer{contract: contract}, nil
}

// bindWalletRequireFreshSigner binds a generic wrapper to an already deployed contract.
func bindWalletRequireFreshSigner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletRequireFreshSignerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletRequireFreshSigner.Contract.WalletRequireFreshSignerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.WalletRequireFreshSignerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.WalletRequireFreshSignerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletRequireFreshSigner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletRequireFreshSigner *WalletRequireFreshSignerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.contract.Transact(opts, method, params...)
}

// REQUIREUTILS is a free data retrieval call binding the contract method 0xcfc63a49.
//
// Solidity: function REQUIRE_UTILS() view returns(address)
func (_WalletRequireFreshSigner *WalletRequireFreshSignerCaller) REQUIREUTILS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletRequireFreshSigner.contract.Call(opts, &out, "REQUIRE_UTILS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REQUIREUTILS is a free data retrieval call binding the contract method 0xcfc63a49.
//
// Solidity: function REQUIRE_UTILS() view returns(address)
func (_WalletRequireFreshSigner *WalletRequireFreshSignerSession) REQUIREUTILS() (common.Address, error) {
	return _WalletRequireFreshSigner.Contract.REQUIREUTILS(&_WalletRequireFreshSigner.CallOpts)
}

// REQUIREUTILS is a free data retrieval call binding the contract method 0xcfc63a49.
//
// Solidity: function REQUIRE_UTILS() view returns(address)
func (_WalletRequireFreshSigner *WalletRequireFreshSignerCallerSession) REQUIREUTILS() (common.Address, error) {
	return _WalletRequireFreshSigner.Contract.REQUIREUTILS(&_WalletRequireFreshSigner.CallOpts)
}

// RequireFreshSigner is a paid mutator transaction binding the contract method 0x0df0c419.
//
// Solidity: function requireFreshSigner(address _signer) returns()
func (_WalletRequireFreshSigner *WalletRequireFreshSignerTransactor) RequireFreshSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.contract.Transact(opts, "requireFreshSigner", _signer)
}

// RequireFreshSigner is a paid mutator transaction binding the contract method 0x0df0c419.
//
// Solidity: function requireFreshSigner(address _signer) returns()
func (_WalletRequireFreshSigner *WalletRequireFreshSignerSession) RequireFreshSigner(_signer common.Address) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.RequireFreshSigner(&_WalletRequireFreshSigner.TransactOpts, _signer)
}

// RequireFreshSigner is a paid mutator transaction binding the contract method 0x0df0c419.
//
// Solidity: function requireFreshSigner(address _signer) returns()
func (_WalletRequireFreshSigner *WalletRequireFreshSignerTransactorSession) RequireFreshSigner(_signer common.Address) (*types.Transaction, error) {
	return _WalletRequireFreshSigner.Contract.RequireFreshSigner(&_WalletRequireFreshSigner.TransactOpts, _signer)
}
