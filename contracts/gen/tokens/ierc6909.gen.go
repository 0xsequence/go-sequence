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

// IERC6909MetaData contains all meta data concerning the IERC6909 contract.
var IERC6909MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC6909ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC6909MetaData.ABI instead.
var IERC6909ABI = IERC6909MetaData.ABI

// IERC6909 is an auto generated Go binding around an Ethereum contract.
type IERC6909 struct {
	IERC6909Caller     // Read-only binding to the contract
	IERC6909Transactor // Write-only binding to the contract
	IERC6909Filterer   // Log filterer for contract events
}

// IERC6909Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC6909Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC6909Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC6909Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC6909Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC6909Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC6909Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC6909Session struct {
	Contract     *IERC6909         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC6909CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC6909CallerSession struct {
	Contract *IERC6909Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC6909TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC6909TransactorSession struct {
	Contract     *IERC6909Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC6909Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC6909Raw struct {
	Contract *IERC6909 // Generic contract binding to access the raw methods on
}

// IERC6909CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC6909CallerRaw struct {
	Contract *IERC6909Caller // Generic read-only contract binding to access the raw methods on
}

// IERC6909TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC6909TransactorRaw struct {
	Contract *IERC6909Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC6909 creates a new instance of IERC6909, bound to a specific deployed contract.
func NewIERC6909(address common.Address, backend bind.ContractBackend) (*IERC6909, error) {
	contract, err := bindIERC6909(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC6909{IERC6909Caller: IERC6909Caller{contract: contract}, IERC6909Transactor: IERC6909Transactor{contract: contract}, IERC6909Filterer: IERC6909Filterer{contract: contract}}, nil
}

// NewIERC6909Caller creates a new read-only instance of IERC6909, bound to a specific deployed contract.
func NewIERC6909Caller(address common.Address, caller bind.ContractCaller) (*IERC6909Caller, error) {
	contract, err := bindIERC6909(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC6909Caller{contract: contract}, nil
}

// NewIERC6909Transactor creates a new write-only instance of IERC6909, bound to a specific deployed contract.
func NewIERC6909Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC6909Transactor, error) {
	contract, err := bindIERC6909(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC6909Transactor{contract: contract}, nil
}

// NewIERC6909Filterer creates a new log filterer instance of IERC6909, bound to a specific deployed contract.
func NewIERC6909Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC6909Filterer, error) {
	contract, err := bindIERC6909(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC6909Filterer{contract: contract}, nil
}

// bindIERC6909 binds a generic wrapper to an already deployed contract.
func bindIERC6909(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC6909MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC6909 *IERC6909Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC6909.Contract.IERC6909Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC6909 *IERC6909Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC6909.Contract.IERC6909Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC6909 *IERC6909Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC6909.Contract.IERC6909Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC6909 *IERC6909CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC6909.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC6909 *IERC6909TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC6909.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC6909 *IERC6909TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC6909.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC6909.contract.Call(opts, &out, "allowance", owner, spender, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909Session) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _IERC6909.Contract.Allowance(&_IERC6909.CallOpts, owner, spender, id)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909CallerSession) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _IERC6909.Contract.Allowance(&_IERC6909.CallOpts, owner, spender, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909Caller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC6909.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909Session) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _IERC6909.Contract.BalanceOf(&_IERC6909.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_IERC6909 *IERC6909CallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _IERC6909.Contract.BalanceOf(&_IERC6909.CallOpts, owner, id)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool)
func (_IERC6909 *IERC6909Caller) IsOperator(opts *bind.CallOpts, owner common.Address, spender common.Address) (bool, error) {
	var out []interface{}
	err := _IERC6909.contract.Call(opts, &out, "isOperator", owner, spender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool)
func (_IERC6909 *IERC6909Session) IsOperator(owner common.Address, spender common.Address) (bool, error) {
	return _IERC6909.Contract.IsOperator(&_IERC6909.CallOpts, owner, spender)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address spender) view returns(bool)
func (_IERC6909 *IERC6909CallerSession) IsOperator(owner common.Address, spender common.Address) (bool, error) {
	return _IERC6909.Contract.IsOperator(&_IERC6909.CallOpts, owner, spender)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC6909 *IERC6909Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC6909.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC6909 *IERC6909Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC6909.Contract.SupportsInterface(&_IERC6909.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC6909 *IERC6909CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC6909.Contract.SupportsInterface(&_IERC6909.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Transactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.contract.Transact(opts, "approve", spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Session) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.Approve(&_IERC6909.TransactOpts, spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909TransactorSession) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.Approve(&_IERC6909.TransactOpts, spender, id, amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address spender, bool approved) returns(bool)
func (_IERC6909 *IERC6909Transactor) SetOperator(opts *bind.TransactOpts, spender common.Address, approved bool) (*types.Transaction, error) {
	return _IERC6909.contract.Transact(opts, "setOperator", spender, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address spender, bool approved) returns(bool)
func (_IERC6909 *IERC6909Session) SetOperator(spender common.Address, approved bool) (*types.Transaction, error) {
	return _IERC6909.Contract.SetOperator(&_IERC6909.TransactOpts, spender, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address spender, bool approved) returns(bool)
func (_IERC6909 *IERC6909TransactorSession) SetOperator(spender common.Address, approved bool) (*types.Transaction, error) {
	return _IERC6909.Contract.SetOperator(&_IERC6909.TransactOpts, spender, approved)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.contract.Transact(opts, "transfer", receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Session) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.Transfer(&_IERC6909.TransactOpts, receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909TransactorSession) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.Transfer(&_IERC6909.TransactOpts, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.contract.Transact(opts, "transferFrom", sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909Session) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.TransferFrom(&_IERC6909.TransactOpts, sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_IERC6909 *IERC6909TransactorSession) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _IERC6909.Contract.TransferFrom(&_IERC6909.TransactOpts, sender, receiver, id, amount)
}

// IERC6909ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC6909 contract.
type IERC6909ApprovalIterator struct {
	Event *IERC6909Approval // Event containing the contract specifics and raw log

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
func (it *IERC6909ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC6909Approval)
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
		it.Event = new(IERC6909Approval)
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
func (it *IERC6909ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC6909ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC6909Approval represents a Approval event raised by the IERC6909 contract.
type IERC6909Approval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*IERC6909ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC6909.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IERC6909ApprovalIterator{contract: _IERC6909.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC6909Approval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC6909.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC6909Approval)
				if err := _IERC6909.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) ParseApproval(log types.Log) (*IERC6909Approval, error) {
	event := new(IERC6909Approval)
	if err := _IERC6909.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC6909OperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the IERC6909 contract.
type IERC6909OperatorSetIterator struct {
	Event *IERC6909OperatorSet // Event containing the contract specifics and raw log

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
func (it *IERC6909OperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC6909OperatorSet)
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
		it.Event = new(IERC6909OperatorSet)
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
func (it *IERC6909OperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC6909OperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC6909OperatorSet represents a OperatorSet event raised by the IERC6909 contract.
type IERC6909OperatorSet struct {
	Owner    common.Address
	Spender  common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed spender, bool approved)
func (_IERC6909 *IERC6909Filterer) FilterOperatorSet(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC6909OperatorSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC6909.contract.FilterLogs(opts, "OperatorSet", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC6909OperatorSetIterator{contract: _IERC6909.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed spender, bool approved)
func (_IERC6909 *IERC6909Filterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *IERC6909OperatorSet, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC6909.contract.WatchLogs(opts, "OperatorSet", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC6909OperatorSet)
				if err := _IERC6909.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed spender, bool approved)
func (_IERC6909 *IERC6909Filterer) ParseOperatorSet(log types.Log) (*IERC6909OperatorSet, error) {
	event := new(IERC6909OperatorSet)
	if err := _IERC6909.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC6909TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC6909 contract.
type IERC6909TransferIterator struct {
	Event *IERC6909Transfer // Event containing the contract specifics and raw log

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
func (it *IERC6909TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC6909Transfer)
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
		it.Event = new(IERC6909Transfer)
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
func (it *IERC6909TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC6909TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC6909Transfer represents a Transfer event raised by the IERC6909 contract.
type IERC6909Transfer struct {
	Caller   common.Address
	Sender   common.Address
	Receiver common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed sender, address indexed receiver, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, id []*big.Int) (*IERC6909TransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC6909.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule, idRule)
	if err != nil {
		return nil, err
	}
	return &IERC6909TransferIterator{contract: _IERC6909.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed sender, address indexed receiver, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC6909Transfer, sender []common.Address, receiver []common.Address, id []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC6909.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC6909Transfer)
				if err := _IERC6909.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed sender, address indexed receiver, uint256 indexed id, uint256 amount)
func (_IERC6909 *IERC6909Filterer) ParseTransfer(log types.Log) (*IERC6909Transfer, error) {
	event := new(IERC6909Transfer)
	if err := _IERC6909.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
