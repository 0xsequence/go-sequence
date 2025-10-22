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

// PayloadCall is an auto generated low-level Go binding around an user-defined struct.
type PayloadCall struct {
	To              common.Address
	Value           *big.Int
	Data            []byte
	GasLimit        *big.Int
	DelegateCall    bool
	OnlyFallback    bool
	BehaviorOnError *big.Int
}

// PayloadDecoded is an auto generated low-level Go binding around an user-defined struct.
type PayloadDecoded struct {
	Kind          uint8
	NoChainId     bool
	Calls         []PayloadCall
	Space         *big.Int
	Nonce         *big.Int
	Message       []byte
	ImageHash     [32]byte
	Digest        [32]byte
	ParentWallets []common.Address
}

// ISapientMetaData contains all meta data concerning the ISapient contract.
var ISapientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"}]",
}

// ISapientABI is the input ABI used to generate the binding from.
// Deprecated: Use ISapientMetaData.ABI instead.
var ISapientABI = ISapientMetaData.ABI

// ISapient is an auto generated Go binding around an Ethereum contract.
type ISapient struct {
	ISapientCaller     // Read-only binding to the contract
	ISapientTransactor // Write-only binding to the contract
	ISapientFilterer   // Log filterer for contract events
}

// ISapientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISapientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISapientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISapientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISapientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISapientSession struct {
	Contract     *ISapient         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISapientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISapientCallerSession struct {
	Contract *ISapientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ISapientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISapientTransactorSession struct {
	Contract     *ISapientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ISapientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISapientRaw struct {
	Contract *ISapient // Generic contract binding to access the raw methods on
}

// ISapientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISapientCallerRaw struct {
	Contract *ISapientCaller // Generic read-only contract binding to access the raw methods on
}

// ISapientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISapientTransactorRaw struct {
	Contract *ISapientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISapient creates a new instance of ISapient, bound to a specific deployed contract.
func NewISapient(address common.Address, backend bind.ContractBackend) (*ISapient, error) {
	contract, err := bindISapient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISapient{ISapientCaller: ISapientCaller{contract: contract}, ISapientTransactor: ISapientTransactor{contract: contract}, ISapientFilterer: ISapientFilterer{contract: contract}}, nil
}

// NewISapientCaller creates a new read-only instance of ISapient, bound to a specific deployed contract.
func NewISapientCaller(address common.Address, caller bind.ContractCaller) (*ISapientCaller, error) {
	contract, err := bindISapient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISapientCaller{contract: contract}, nil
}

// NewISapientTransactor creates a new write-only instance of ISapient, bound to a specific deployed contract.
func NewISapientTransactor(address common.Address, transactor bind.ContractTransactor) (*ISapientTransactor, error) {
	contract, err := bindISapient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISapientTransactor{contract: contract}, nil
}

// NewISapientFilterer creates a new log filterer instance of ISapient, bound to a specific deployed contract.
func NewISapientFilterer(address common.Address, filterer bind.ContractFilterer) (*ISapientFilterer, error) {
	contract, err := bindISapient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISapientFilterer{contract: contract}, nil
}

// bindISapient binds a generic wrapper to an already deployed contract.
func bindISapient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISapientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISapient *ISapientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISapient.Contract.ISapientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISapient *ISapientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISapient.Contract.ISapientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISapient *ISapientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISapient.Contract.ISapientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISapient *ISapientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISapient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISapient *ISapientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISapient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISapient *ISapientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISapient.Contract.contract.Transact(opts, method, params...)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_ISapient *ISapientCaller) RecoverSapientSignature(opts *bind.CallOpts, payload PayloadDecoded, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _ISapient.contract.Call(opts, &out, "recoverSapientSignature", payload, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_ISapient *ISapientSession) RecoverSapientSignature(payload PayloadDecoded, signature []byte) ([32]byte, error) {
	return _ISapient.Contract.RecoverSapientSignature(&_ISapient.CallOpts, payload, signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_ISapient *ISapientCallerSession) RecoverSapientSignature(payload PayloadDecoded, signature []byte) ([32]byte, error) {
	return _ISapient.Contract.RecoverSapientSignature(&_ISapient.CallOpts, payload, signature)
}
