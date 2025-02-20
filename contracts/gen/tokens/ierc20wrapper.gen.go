// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokens

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

// IERC20WrapperMetaData contains all meta data concerning the IERC20Wrapper contract.
var IERC20WrapperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getIdAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNTokens\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// IERC20WrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20WrapperMetaData.ABI instead.
var IERC20WrapperABI = IERC20WrapperMetaData.ABI

// IERC20Wrapper is an auto generated Go binding around an Ethereum contract.
type IERC20Wrapper struct {
	IERC20WrapperCaller     // Read-only binding to the contract
	IERC20WrapperTransactor // Write-only binding to the contract
	IERC20WrapperFilterer   // Log filterer for contract events
}

// IERC20WrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20WrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20WrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20WrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20WrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20WrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20WrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20WrapperSession struct {
	Contract     *IERC20Wrapper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20WrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20WrapperCallerSession struct {
	Contract *IERC20WrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IERC20WrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20WrapperTransactorSession struct {
	Contract     *IERC20WrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IERC20WrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20WrapperRaw struct {
	Contract *IERC20Wrapper // Generic contract binding to access the raw methods on
}

// IERC20WrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20WrapperCallerRaw struct {
	Contract *IERC20WrapperCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20WrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20WrapperTransactorRaw struct {
	Contract *IERC20WrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Wrapper creates a new instance of IERC20Wrapper, bound to a specific deployed contract.
func NewIERC20Wrapper(address common.Address, backend bind.ContractBackend) (*IERC20Wrapper, error) {
	contract, err := bindIERC20Wrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Wrapper{IERC20WrapperCaller: IERC20WrapperCaller{contract: contract}, IERC20WrapperTransactor: IERC20WrapperTransactor{contract: contract}, IERC20WrapperFilterer: IERC20WrapperFilterer{contract: contract}}, nil
}

// NewIERC20WrapperCaller creates a new read-only instance of IERC20Wrapper, bound to a specific deployed contract.
func NewIERC20WrapperCaller(address common.Address, caller bind.ContractCaller) (*IERC20WrapperCaller, error) {
	contract, err := bindIERC20Wrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperCaller{contract: contract}, nil
}

// NewIERC20WrapperTransactor creates a new write-only instance of IERC20Wrapper, bound to a specific deployed contract.
func NewIERC20WrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20WrapperTransactor, error) {
	contract, err := bindIERC20Wrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperTransactor{contract: contract}, nil
}

// NewIERC20WrapperFilterer creates a new log filterer instance of IERC20Wrapper, bound to a specific deployed contract.
func NewIERC20WrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20WrapperFilterer, error) {
	contract, err := bindIERC20Wrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperFilterer{contract: contract}, nil
}

// bindIERC20Wrapper binds a generic wrapper to an already deployed contract.
func bindIERC20Wrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20WrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Wrapper *IERC20WrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Wrapper.Contract.IERC20WrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Wrapper *IERC20WrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.IERC20WrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Wrapper *IERC20WrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.IERC20WrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Wrapper *IERC20WrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Wrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Wrapper *IERC20WrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Wrapper *IERC20WrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_IERC20Wrapper *IERC20WrapperCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "balanceOf", _owner, _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_IERC20Wrapper *IERC20WrapperSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _IERC20Wrapper.Contract.BalanceOf(&_IERC20Wrapper.CallOpts, _owner, _id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address _owner, uint256 _id) view returns(uint256)
func (_IERC20Wrapper *IERC20WrapperCallerSession) BalanceOf(_owner common.Address, _id *big.Int) (*big.Int, error) {
	return _IERC20Wrapper.Contract.BalanceOf(&_IERC20Wrapper.CallOpts, _owner, _id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_IERC20Wrapper *IERC20WrapperCaller) BalanceOfBatch(opts *bind.CallOpts, _owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "balanceOfBatch", _owners, _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_IERC20Wrapper *IERC20WrapperSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _IERC20Wrapper.Contract.BalanceOfBatch(&_IERC20Wrapper.CallOpts, _owners, _ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] _owners, uint256[] _ids) view returns(uint256[])
func (_IERC20Wrapper *IERC20WrapperCallerSession) BalanceOfBatch(_owners []common.Address, _ids []*big.Int) ([]*big.Int, error) {
	return _IERC20Wrapper.Contract.BalanceOfBatch(&_IERC20Wrapper.CallOpts, _owners, _ids)
}

// GetIdAddress is a free data retrieval call binding the contract method 0x7358e9a5.
//
// Solidity: function getIdAddress(uint256 _id) view returns(address token)
func (_IERC20Wrapper *IERC20WrapperCaller) GetIdAddress(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "getIdAddress", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdAddress is a free data retrieval call binding the contract method 0x7358e9a5.
//
// Solidity: function getIdAddress(uint256 _id) view returns(address token)
func (_IERC20Wrapper *IERC20WrapperSession) GetIdAddress(_id *big.Int) (common.Address, error) {
	return _IERC20Wrapper.Contract.GetIdAddress(&_IERC20Wrapper.CallOpts, _id)
}

// GetIdAddress is a free data retrieval call binding the contract method 0x7358e9a5.
//
// Solidity: function getIdAddress(uint256 _id) view returns(address token)
func (_IERC20Wrapper *IERC20WrapperCallerSession) GetIdAddress(_id *big.Int) (common.Address, error) {
	return _IERC20Wrapper.Contract.GetIdAddress(&_IERC20Wrapper.CallOpts, _id)
}

// GetNTokens is a free data retrieval call binding the contract method 0x9040a949.
//
// Solidity: function getNTokens() view returns()
func (_IERC20Wrapper *IERC20WrapperCaller) GetNTokens(opts *bind.CallOpts) error {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "getNTokens")

	if err != nil {
		return err
	}

	return err

}

// GetNTokens is a free data retrieval call binding the contract method 0x9040a949.
//
// Solidity: function getNTokens() view returns()
func (_IERC20Wrapper *IERC20WrapperSession) GetNTokens() error {
	return _IERC20Wrapper.Contract.GetNTokens(&_IERC20Wrapper.CallOpts)
}

// GetNTokens is a free data retrieval call binding the contract method 0x9040a949.
//
// Solidity: function getNTokens() view returns()
func (_IERC20Wrapper *IERC20WrapperCallerSession) GetNTokens() error {
	return _IERC20Wrapper.Contract.GetNTokens(&_IERC20Wrapper.CallOpts)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address _token) view returns(uint256 tokenID)
func (_IERC20Wrapper *IERC20WrapperCaller) GetTokenID(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "getTokenID", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address _token) view returns(uint256 tokenID)
func (_IERC20Wrapper *IERC20WrapperSession) GetTokenID(_token common.Address) (*big.Int, error) {
	return _IERC20Wrapper.Contract.GetTokenID(&_IERC20Wrapper.CallOpts, _token)
}

// GetTokenID is a free data retrieval call binding the contract method 0x63f8071c.
//
// Solidity: function getTokenID(address _token) view returns(uint256 tokenID)
func (_IERC20Wrapper *IERC20WrapperCallerSession) GetTokenID(_token common.Address) (*big.Int, error) {
	return _IERC20Wrapper.Contract.GetTokenID(&_IERC20Wrapper.CallOpts, _token)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_IERC20Wrapper *IERC20WrapperCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC20Wrapper.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_IERC20Wrapper *IERC20WrapperSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _IERC20Wrapper.Contract.IsApprovedForAll(&_IERC20Wrapper.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_IERC20Wrapper *IERC20WrapperCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _IERC20Wrapper.Contract.IsApprovedForAll(&_IERC20Wrapper.CallOpts, _owner, _operator)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _token, address _recipient, uint256 _value) payable returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "deposit", _token, _recipient, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _token, address _recipient, uint256 _value) payable returns()
func (_IERC20Wrapper *IERC20WrapperSession) Deposit(_token common.Address, _recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Deposit(&_IERC20Wrapper.TransactOpts, _token, _recipient, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _token, address _recipient, uint256 _value) payable returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) Deposit(_token common.Address, _recipient common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Deposit(&_IERC20Wrapper.TransactOpts, _token, _recipient, _value)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address _from, uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "onERC1155BatchReceived", _operator, _from, _ids, _values, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address _from, uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperSession) OnERC1155BatchReceived(_operator common.Address, _from common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.OnERC1155BatchReceived(&_IERC20Wrapper.TransactOpts, _operator, _from, _ids, _values, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address _from, uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperTransactorSession) OnERC1155BatchReceived(_operator common.Address, _from common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.OnERC1155BatchReceived(&_IERC20Wrapper.TransactOpts, _operator, _from, _ids, _values, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperTransactor) OnERC1155Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "onERC1155Received", _operator, _from, _id, _value, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.OnERC1155Received(&_IERC20Wrapper.TransactOpts, _operator, _from, _id, _value, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _value, bytes _data) returns(bytes4)
func (_IERC20Wrapper *IERC20WrapperTransactorSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.OnERC1155Received(&_IERC20Wrapper.TransactOpts, _operator, _from, _id, _value, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SafeBatchTransferFrom(&_IERC20Wrapper.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address _from, address _to, uint256[] _ids, uint256[] _amounts, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SafeBatchTransferFrom(&_IERC20Wrapper.TransactOpts, _from, _to, _ids, _amounts, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "safeTransferFrom", _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SafeTransferFrom(&_IERC20Wrapper.TransactOpts, _from, _to, _id, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _id, uint256 _amount, bytes _data) returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SafeTransferFrom(&_IERC20Wrapper.TransactOpts, _from, _to, _id, _amount, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC20Wrapper *IERC20WrapperSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SetApprovalForAll(&_IERC20Wrapper.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.SetApprovalForAll(&_IERC20Wrapper.TransactOpts, _operator, _approved)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _value) returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) Withdraw(opts *bind.TransactOpts, _token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.Transact(opts, "withdraw", _token, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _value) returns()
func (_IERC20Wrapper *IERC20WrapperSession) Withdraw(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Withdraw(&_IERC20Wrapper.TransactOpts, _token, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _token, address _to, uint256 _value) returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) Withdraw(_token common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Withdraw(&_IERC20Wrapper.TransactOpts, _token, _to, _value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IERC20Wrapper *IERC20WrapperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Wrapper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IERC20Wrapper *IERC20WrapperSession) Receive() (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Receive(&_IERC20Wrapper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IERC20Wrapper *IERC20WrapperTransactorSession) Receive() (*types.Transaction, error) {
	return _IERC20Wrapper.Contract.Receive(&_IERC20Wrapper.TransactOpts)
}

// IERC20WrapperApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC20Wrapper contract.
type IERC20WrapperApprovalForAllIterator struct {
	Event *IERC20WrapperApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20WrapperApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20WrapperApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20WrapperApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20WrapperApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20WrapperApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20WrapperApprovalForAll represents a ApprovalForAll event raised by the IERC20Wrapper contract.
type IERC20WrapperApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC20Wrapper *IERC20WrapperFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*IERC20WrapperApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperApprovalForAllIterator{contract: _IERC20Wrapper.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC20Wrapper *IERC20WrapperFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC20WrapperApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20WrapperApprovalForAll)
				if err := _IERC20Wrapper.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_IERC20Wrapper *IERC20WrapperFilterer) ParseApprovalForAll(log types.Log) (*IERC20WrapperApprovalForAll, error) {
	event := new(IERC20WrapperApprovalForAll)
	if err := _IERC20Wrapper.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20WrapperTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the IERC20Wrapper contract.
type IERC20WrapperTransferBatchIterator struct {
	Event *IERC20WrapperTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20WrapperTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20WrapperTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20WrapperTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20WrapperTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20WrapperTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20WrapperTransferBatch represents a TransferBatch event raised by the IERC20Wrapper contract.
type IERC20WrapperTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_IERC20Wrapper *IERC20WrapperFilterer) FilterTransferBatch(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*IERC20WrapperTransferBatchIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.FilterLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperTransferBatchIterator{contract: _IERC20Wrapper.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_IERC20Wrapper *IERC20WrapperFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *IERC20WrapperTransferBatch, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.WatchLogs(opts, "TransferBatch", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20WrapperTransferBatch)
				if err := _IERC20Wrapper.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed _operator, address indexed _from, address indexed _to, uint256[] _ids, uint256[] _amounts)
func (_IERC20Wrapper *IERC20WrapperFilterer) ParseTransferBatch(log types.Log) (*IERC20WrapperTransferBatch, error) {
	event := new(IERC20WrapperTransferBatch)
	if err := _IERC20Wrapper.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20WrapperTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the IERC20Wrapper contract.
type IERC20WrapperTransferSingleIterator struct {
	Event *IERC20WrapperTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20WrapperTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20WrapperTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20WrapperTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20WrapperTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20WrapperTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20WrapperTransferSingle represents a TransferSingle event raised by the IERC20Wrapper contract.
type IERC20WrapperTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_IERC20Wrapper *IERC20WrapperFilterer) FilterTransferSingle(opts *bind.FilterOpts, _operator []common.Address, _from []common.Address, _to []common.Address) (*IERC20WrapperTransferSingleIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.FilterLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20WrapperTransferSingleIterator{contract: _IERC20Wrapper.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_IERC20Wrapper *IERC20WrapperFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *IERC20WrapperTransferSingle, _operator []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _IERC20Wrapper.contract.WatchLogs(opts, "TransferSingle", _operatorRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20WrapperTransferSingle)
				if err := _IERC20Wrapper.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed _operator, address indexed _from, address indexed _to, uint256 _id, uint256 _amount)
func (_IERC20Wrapper *IERC20WrapperFilterer) ParseTransferSingle(log types.Log) (*IERC20WrapperTransferSingle, error) {
	event := new(IERC20WrapperTransferSingle)
	if err := _IERC20Wrapper.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
