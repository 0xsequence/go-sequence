// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sapientmock

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

// SapientMockMetaData contains all meta data concerning the SapientMock contract.
var SapientMockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"recoverSapientSignatureCompact\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]}]",
	Bin: "0x608080604052346015576101f9908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806313792a4a146100995763898bd9211461003257600080fd5b346100945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100945760243567ffffffffffffffff81116100945761008c6100866020923690600401610128565b90610156565b604051908152f35b600080fd5b346100945760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100945760043567ffffffffffffffff8111610094577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc61012091360301126100945760243567ffffffffffffffff81116100945761008c61008660209236906004015b9181601f840112156100945782359167ffffffffffffffff8311610094576020838186019501011161009457565b602082036101995735906020811061016c575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b7f4be6321b0000000000000000000000000000000000000000000000000000000060005260046000fdfea2646970667358221220eee95114dd28e7a8bc9444b8b58c282669afd11f4b40776539458a6435bcbd4b64736f6c634300081c0033",
}

// SapientMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SapientMockMetaData.ABI instead.
var SapientMockABI = SapientMockMetaData.ABI

// SapientMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SapientMockMetaData.Bin instead.
var SapientMockBin = SapientMockMetaData.Bin

// DeploySapientMock deploys a new Ethereum contract, binding an instance of SapientMock to it.
func DeploySapientMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SapientMock, error) {
	parsed, err := SapientMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SapientMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SapientMock{SapientMockCaller: SapientMockCaller{contract: contract}, SapientMockTransactor: SapientMockTransactor{contract: contract}, SapientMockFilterer: SapientMockFilterer{contract: contract}}, nil
}

// SapientMock is an auto generated Go binding around an Ethereum contract.
type SapientMock struct {
	SapientMockCaller     // Read-only binding to the contract
	SapientMockTransactor // Write-only binding to the contract
	SapientMockFilterer   // Log filterer for contract events
}

// SapientMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SapientMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SapientMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SapientMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SapientMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SapientMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SapientMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SapientMockSession struct {
	Contract     *SapientMock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SapientMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SapientMockCallerSession struct {
	Contract *SapientMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SapientMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SapientMockTransactorSession struct {
	Contract     *SapientMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SapientMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SapientMockRaw struct {
	Contract *SapientMock // Generic contract binding to access the raw methods on
}

// SapientMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SapientMockCallerRaw struct {
	Contract *SapientMockCaller // Generic read-only contract binding to access the raw methods on
}

// SapientMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SapientMockTransactorRaw struct {
	Contract *SapientMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSapientMock creates a new instance of SapientMock, bound to a specific deployed contract.
func NewSapientMock(address common.Address, backend bind.ContractBackend) (*SapientMock, error) {
	contract, err := bindSapientMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SapientMock{SapientMockCaller: SapientMockCaller{contract: contract}, SapientMockTransactor: SapientMockTransactor{contract: contract}, SapientMockFilterer: SapientMockFilterer{contract: contract}}, nil
}

// NewSapientMockCaller creates a new read-only instance of SapientMock, bound to a specific deployed contract.
func NewSapientMockCaller(address common.Address, caller bind.ContractCaller) (*SapientMockCaller, error) {
	contract, err := bindSapientMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SapientMockCaller{contract: contract}, nil
}

// NewSapientMockTransactor creates a new write-only instance of SapientMock, bound to a specific deployed contract.
func NewSapientMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SapientMockTransactor, error) {
	contract, err := bindSapientMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SapientMockTransactor{contract: contract}, nil
}

// NewSapientMockFilterer creates a new log filterer instance of SapientMock, bound to a specific deployed contract.
func NewSapientMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SapientMockFilterer, error) {
	contract, err := bindSapientMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SapientMockFilterer{contract: contract}, nil
}

// bindSapientMock binds a generic wrapper to an already deployed contract.
func bindSapientMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SapientMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SapientMock *SapientMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SapientMock.Contract.SapientMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SapientMock *SapientMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SapientMock.Contract.SapientMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SapientMock *SapientMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SapientMock.Contract.SapientMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SapientMock *SapientMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SapientMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SapientMock *SapientMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SapientMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SapientMock *SapientMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SapientMock.Contract.contract.Transact(opts, method, params...)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockCaller) RecoverSapientSignature(opts *bind.CallOpts, arg0 PayloadDecoded, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _SapientMock.contract.Call(opts, &out, "recoverSapientSignature", arg0, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockSession) RecoverSapientSignature(arg0 PayloadDecoded, signature []byte) ([32]byte, error) {
	return _SapientMock.Contract.RecoverSapientSignature(&_SapientMock.CallOpts, arg0, signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockCallerSession) RecoverSapientSignature(arg0 PayloadDecoded, signature []byte) ([32]byte, error) {
	return _SapientMock.Contract.RecoverSapientSignature(&_SapientMock.CallOpts, arg0, signature)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockCaller) RecoverSapientSignatureCompact(opts *bind.CallOpts, arg0 [32]byte, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _SapientMock.contract.Call(opts, &out, "recoverSapientSignatureCompact", arg0, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockSession) RecoverSapientSignatureCompact(arg0 [32]byte, signature []byte) ([32]byte, error) {
	return _SapientMock.Contract.RecoverSapientSignatureCompact(&_SapientMock.CallOpts, arg0, signature)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 , bytes signature) pure returns(bytes32)
func (_SapientMock *SapientMockCallerSession) RecoverSapientSignatureCompact(arg0 [32]byte, signature []byte) ([32]byte, error) {
	return _SapientMock.Contract.RecoverSapientSignatureCompact(&_SapientMock.CallOpts, arg0, signature)
}
