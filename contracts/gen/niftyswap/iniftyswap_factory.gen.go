// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package niftyswap

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

// INiftyswapFactoryMetaData contains all meta data concerning the INiftyswapFactory contract.
var INiftyswapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currencyID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"}],\"name\":\"NewExchange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_currencyID\",\"type\":\"uint256\"}],\"name\":\"createExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_currencyID\",\"type\":\"uint256\"}],\"name\":\"tokensToExchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// INiftyswapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use INiftyswapFactoryMetaData.ABI instead.
var INiftyswapFactoryABI = INiftyswapFactoryMetaData.ABI

// INiftyswapFactory is an auto generated Go binding around an Ethereum contract.
type INiftyswapFactory struct {
	INiftyswapFactoryCaller     // Read-only binding to the contract
	INiftyswapFactoryTransactor // Write-only binding to the contract
	INiftyswapFactoryFilterer   // Log filterer for contract events
}

// INiftyswapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type INiftyswapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INiftyswapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INiftyswapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INiftyswapFactorySession struct {
	Contract     *INiftyswapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// INiftyswapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INiftyswapFactoryCallerSession struct {
	Contract *INiftyswapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// INiftyswapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INiftyswapFactoryTransactorSession struct {
	Contract     *INiftyswapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// INiftyswapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type INiftyswapFactoryRaw struct {
	Contract *INiftyswapFactory // Generic contract binding to access the raw methods on
}

// INiftyswapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INiftyswapFactoryCallerRaw struct {
	Contract *INiftyswapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// INiftyswapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INiftyswapFactoryTransactorRaw struct {
	Contract *INiftyswapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINiftyswapFactory creates a new instance of INiftyswapFactory, bound to a specific deployed contract.
func NewINiftyswapFactory(address common.Address, backend bind.ContractBackend) (*INiftyswapFactory, error) {
	contract, err := bindINiftyswapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory{INiftyswapFactoryCaller: INiftyswapFactoryCaller{contract: contract}, INiftyswapFactoryTransactor: INiftyswapFactoryTransactor{contract: contract}, INiftyswapFactoryFilterer: INiftyswapFactoryFilterer{contract: contract}}, nil
}

// NewINiftyswapFactoryCaller creates a new read-only instance of INiftyswapFactory, bound to a specific deployed contract.
func NewINiftyswapFactoryCaller(address common.Address, caller bind.ContractCaller) (*INiftyswapFactoryCaller, error) {
	contract, err := bindINiftyswapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactoryCaller{contract: contract}, nil
}

// NewINiftyswapFactoryTransactor creates a new write-only instance of INiftyswapFactory, bound to a specific deployed contract.
func NewINiftyswapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*INiftyswapFactoryTransactor, error) {
	contract, err := bindINiftyswapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactoryTransactor{contract: contract}, nil
}

// NewINiftyswapFactoryFilterer creates a new log filterer instance of INiftyswapFactory, bound to a specific deployed contract.
func NewINiftyswapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*INiftyswapFactoryFilterer, error) {
	contract, err := bindINiftyswapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactoryFilterer{contract: contract}, nil
}

// bindINiftyswapFactory binds a generic wrapper to an already deployed contract.
func bindINiftyswapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INiftyswapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapFactory *INiftyswapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapFactory.Contract.INiftyswapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapFactory *INiftyswapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.INiftyswapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapFactory *INiftyswapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.INiftyswapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapFactory *INiftyswapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapFactory *INiftyswapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapFactory *INiftyswapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.contract.Transact(opts, method, params...)
}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _currencyID) view returns(address)
func (_INiftyswapFactory *INiftyswapFactoryCaller) TokensToExchange(opts *bind.CallOpts, _token common.Address, _currency common.Address, _currencyID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapFactory.contract.Call(opts, &out, "tokensToExchange", _token, _currency, _currencyID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _currencyID) view returns(address)
func (_INiftyswapFactory *INiftyswapFactorySession) TokensToExchange(_token common.Address, _currency common.Address, _currencyID *big.Int) (common.Address, error) {
	return _INiftyswapFactory.Contract.TokensToExchange(&_INiftyswapFactory.CallOpts, _token, _currency, _currencyID)
}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _currencyID) view returns(address)
func (_INiftyswapFactory *INiftyswapFactoryCallerSession) TokensToExchange(_token common.Address, _currency common.Address, _currencyID *big.Int) (common.Address, error) {
	return _INiftyswapFactory.Contract.TokensToExchange(&_INiftyswapFactory.CallOpts, _token, _currency, _currencyID)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _currencyID) returns()
func (_INiftyswapFactory *INiftyswapFactoryTransactor) CreateExchange(opts *bind.TransactOpts, _token common.Address, _currency common.Address, _currencyID *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory.contract.Transact(opts, "createExchange", _token, _currency, _currencyID)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _currencyID) returns()
func (_INiftyswapFactory *INiftyswapFactorySession) CreateExchange(_token common.Address, _currency common.Address, _currencyID *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.CreateExchange(&_INiftyswapFactory.TransactOpts, _token, _currency, _currencyID)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _currencyID) returns()
func (_INiftyswapFactory *INiftyswapFactoryTransactorSession) CreateExchange(_token common.Address, _currency common.Address, _currencyID *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory.Contract.CreateExchange(&_INiftyswapFactory.TransactOpts, _token, _currency, _currencyID)
}

// INiftyswapFactoryNewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the INiftyswapFactory contract.
type INiftyswapFactoryNewExchangeIterator struct {
	Event *INiftyswapFactoryNewExchange // Event containing the contract specifics and raw log

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
func (it *INiftyswapFactoryNewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapFactoryNewExchange)
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
		it.Event = new(INiftyswapFactoryNewExchange)
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
func (it *INiftyswapFactoryNewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapFactoryNewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapFactoryNewExchange represents a NewExchange event raised by the INiftyswapFactory contract.
type INiftyswapFactoryNewExchange struct {
	Token      common.Address
	Currency   common.Address
	CurrencyID *big.Int
	Exchange   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x23658fa6d505b3e3034045b3937d4239cbdaa345bfb0c4a2d6637ade8b85457c.
//
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed currencyID, address exchange)
func (_INiftyswapFactory *INiftyswapFactoryFilterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, currency []common.Address, currencyID []*big.Int) (*INiftyswapFactoryNewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var currencyIDRule []interface{}
	for _, currencyIDItem := range currencyID {
		currencyIDRule = append(currencyIDRule, currencyIDItem)
	}

	logs, sub, err := _INiftyswapFactory.contract.FilterLogs(opts, "NewExchange", tokenRule, currencyRule, currencyIDRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactoryNewExchangeIterator{contract: _INiftyswapFactory.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x23658fa6d505b3e3034045b3937d4239cbdaa345bfb0c4a2d6637ade8b85457c.
//
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed currencyID, address exchange)
func (_INiftyswapFactory *INiftyswapFactoryFilterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *INiftyswapFactoryNewExchange, token []common.Address, currency []common.Address, currencyID []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var currencyIDRule []interface{}
	for _, currencyIDItem := range currencyID {
		currencyIDRule = append(currencyIDRule, currencyIDItem)
	}

	logs, sub, err := _INiftyswapFactory.contract.WatchLogs(opts, "NewExchange", tokenRule, currencyRule, currencyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapFactoryNewExchange)
				if err := _INiftyswapFactory.contract.UnpackLog(event, "NewExchange", log); err != nil {
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

// ParseNewExchange is a log parse operation binding the contract event 0x23658fa6d505b3e3034045b3937d4239cbdaa345bfb0c4a2d6637ade8b85457c.
//
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed currencyID, address exchange)
func (_INiftyswapFactory *INiftyswapFactoryFilterer) ParseNewExchange(log types.Log) (*INiftyswapFactoryNewExchange, error) {
	event := new(INiftyswapFactoryNewExchange)
	if err := _INiftyswapFactory.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
