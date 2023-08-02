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
)

// INiftyswapFactory20MetaData contains all meta data concerning the INiftyswapFactory20 contract.
var INiftyswapFactory20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"}],\"name\":\"NewExchange\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_instance\",\"type\":\"uint256\"}],\"name\":\"createExchange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"}],\"name\":\"getPairExchanges\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_instance\",\"type\":\"uint256\"}],\"name\":\"tokensToExchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// INiftyswapFactory20ABI is the input ABI used to generate the binding from.
// Deprecated: Use INiftyswapFactory20MetaData.ABI instead.
var INiftyswapFactory20ABI = INiftyswapFactory20MetaData.ABI

// INiftyswapFactory20 is an auto generated Go binding around an Ethereum contract.
type INiftyswapFactory20 struct {
	INiftyswapFactory20Caller     // Read-only binding to the contract
	INiftyswapFactory20Transactor // Write-only binding to the contract
	INiftyswapFactory20Filterer   // Log filterer for contract events
}

// INiftyswapFactory20Caller is an auto generated read-only Go binding around an Ethereum contract.
type INiftyswapFactory20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactory20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type INiftyswapFactory20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactory20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INiftyswapFactory20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapFactory20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INiftyswapFactory20Session struct {
	Contract     *INiftyswapFactory20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// INiftyswapFactory20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INiftyswapFactory20CallerSession struct {
	Contract *INiftyswapFactory20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// INiftyswapFactory20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INiftyswapFactory20TransactorSession struct {
	Contract     *INiftyswapFactory20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// INiftyswapFactory20Raw is an auto generated low-level Go binding around an Ethereum contract.
type INiftyswapFactory20Raw struct {
	Contract *INiftyswapFactory20 // Generic contract binding to access the raw methods on
}

// INiftyswapFactory20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INiftyswapFactory20CallerRaw struct {
	Contract *INiftyswapFactory20Caller // Generic read-only contract binding to access the raw methods on
}

// INiftyswapFactory20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INiftyswapFactory20TransactorRaw struct {
	Contract *INiftyswapFactory20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewINiftyswapFactory20 creates a new instance of INiftyswapFactory20, bound to a specific deployed contract.
func NewINiftyswapFactory20(address common.Address, backend bind.ContractBackend) (*INiftyswapFactory20, error) {
	contract, err := bindINiftyswapFactory20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory20{INiftyswapFactory20Caller: INiftyswapFactory20Caller{contract: contract}, INiftyswapFactory20Transactor: INiftyswapFactory20Transactor{contract: contract}, INiftyswapFactory20Filterer: INiftyswapFactory20Filterer{contract: contract}}, nil
}

// NewINiftyswapFactory20Caller creates a new read-only instance of INiftyswapFactory20, bound to a specific deployed contract.
func NewINiftyswapFactory20Caller(address common.Address, caller bind.ContractCaller) (*INiftyswapFactory20Caller, error) {
	contract, err := bindINiftyswapFactory20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory20Caller{contract: contract}, nil
}

// NewINiftyswapFactory20Transactor creates a new write-only instance of INiftyswapFactory20, bound to a specific deployed contract.
func NewINiftyswapFactory20Transactor(address common.Address, transactor bind.ContractTransactor) (*INiftyswapFactory20Transactor, error) {
	contract, err := bindINiftyswapFactory20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory20Transactor{contract: contract}, nil
}

// NewINiftyswapFactory20Filterer creates a new log filterer instance of INiftyswapFactory20, bound to a specific deployed contract.
func NewINiftyswapFactory20Filterer(address common.Address, filterer bind.ContractFilterer) (*INiftyswapFactory20Filterer, error) {
	contract, err := bindINiftyswapFactory20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory20Filterer{contract: contract}, nil
}

// bindINiftyswapFactory20 binds a generic wrapper to an already deployed contract.
func bindINiftyswapFactory20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INiftyswapFactory20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapFactory20 *INiftyswapFactory20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapFactory20.Contract.INiftyswapFactory20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapFactory20 *INiftyswapFactory20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.INiftyswapFactory20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapFactory20 *INiftyswapFactory20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.INiftyswapFactory20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapFactory20 *INiftyswapFactory20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapFactory20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapFactory20 *INiftyswapFactory20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapFactory20 *INiftyswapFactory20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.contract.Transact(opts, method, params...)
}

// GetPairExchanges is a free data retrieval call binding the contract method 0x9b307388.
//
// Solidity: function getPairExchanges(address _token, address _currency) view returns(address[])
func (_INiftyswapFactory20 *INiftyswapFactory20Caller) GetPairExchanges(opts *bind.CallOpts, _token common.Address, _currency common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _INiftyswapFactory20.contract.Call(opts, &out, "getPairExchanges", _token, _currency)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPairExchanges is a free data retrieval call binding the contract method 0x9b307388.
//
// Solidity: function getPairExchanges(address _token, address _currency) view returns(address[])
func (_INiftyswapFactory20 *INiftyswapFactory20Session) GetPairExchanges(_token common.Address, _currency common.Address) ([]common.Address, error) {
	return _INiftyswapFactory20.Contract.GetPairExchanges(&_INiftyswapFactory20.CallOpts, _token, _currency)
}

// GetPairExchanges is a free data retrieval call binding the contract method 0x9b307388.
//
// Solidity: function getPairExchanges(address _token, address _currency) view returns(address[])
func (_INiftyswapFactory20 *INiftyswapFactory20CallerSession) GetPairExchanges(_token common.Address, _currency common.Address) ([]common.Address, error) {
	return _INiftyswapFactory20.Contract.GetPairExchanges(&_INiftyswapFactory20.CallOpts, _token, _currency)
}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _instance) view returns(address)
func (_INiftyswapFactory20 *INiftyswapFactory20Caller) TokensToExchange(opts *bind.CallOpts, _token common.Address, _currency common.Address, _instance *big.Int) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapFactory20.contract.Call(opts, &out, "tokensToExchange", _token, _currency, _instance)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _instance) view returns(address)
func (_INiftyswapFactory20 *INiftyswapFactory20Session) TokensToExchange(_token common.Address, _currency common.Address, _instance *big.Int) (common.Address, error) {
	return _INiftyswapFactory20.Contract.TokensToExchange(&_INiftyswapFactory20.CallOpts, _token, _currency, _instance)
}

// TokensToExchange is a free data retrieval call binding the contract method 0x1427474c.
//
// Solidity: function tokensToExchange(address _token, address _currency, uint256 _instance) view returns(address)
func (_INiftyswapFactory20 *INiftyswapFactory20CallerSession) TokensToExchange(_token common.Address, _currency common.Address, _instance *big.Int) (common.Address, error) {
	return _INiftyswapFactory20.Contract.TokensToExchange(&_INiftyswapFactory20.CallOpts, _token, _currency, _instance)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _instance) returns()
func (_INiftyswapFactory20 *INiftyswapFactory20Transactor) CreateExchange(opts *bind.TransactOpts, _token common.Address, _currency common.Address, _instance *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory20.contract.Transact(opts, "createExchange", _token, _currency, _instance)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _instance) returns()
func (_INiftyswapFactory20 *INiftyswapFactory20Session) CreateExchange(_token common.Address, _currency common.Address, _instance *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.CreateExchange(&_INiftyswapFactory20.TransactOpts, _token, _currency, _instance)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x8359289c.
//
// Solidity: function createExchange(address _token, address _currency, uint256 _instance) returns()
func (_INiftyswapFactory20 *INiftyswapFactory20TransactorSession) CreateExchange(_token common.Address, _currency common.Address, _instance *big.Int) (*types.Transaction, error) {
	return _INiftyswapFactory20.Contract.CreateExchange(&_INiftyswapFactory20.TransactOpts, _token, _currency, _instance)
}

// INiftyswapFactory20NewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the INiftyswapFactory20 contract.
type INiftyswapFactory20NewExchangeIterator struct {
	Event *INiftyswapFactory20NewExchange // Event containing the contract specifics and raw log

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
func (it *INiftyswapFactory20NewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapFactory20NewExchange)
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
		it.Event = new(INiftyswapFactory20NewExchange)
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
func (it *INiftyswapFactory20NewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapFactory20NewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapFactory20NewExchange represents a NewExchange event raised by the INiftyswapFactory20 contract.
type INiftyswapFactory20NewExchange struct {
	Token    common.Address
	Currency common.Address
	Salt     *big.Int
	Exchange common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x23658fa6d505b3e3034045b3937d4239cbdaa345bfb0c4a2d6637ade8b85457c.
//
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed salt, address exchange)
func (_INiftyswapFactory20 *INiftyswapFactory20Filterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, currency []common.Address, salt []*big.Int) (*INiftyswapFactory20NewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _INiftyswapFactory20.contract.FilterLogs(opts, "NewExchange", tokenRule, currencyRule, saltRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapFactory20NewExchangeIterator{contract: _INiftyswapFactory20.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x23658fa6d505b3e3034045b3937d4239cbdaa345bfb0c4a2d6637ade8b85457c.
//
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed salt, address exchange)
func (_INiftyswapFactory20 *INiftyswapFactory20Filterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *INiftyswapFactory20NewExchange, token []common.Address, currency []common.Address, salt []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var saltRule []interface{}
	for _, saltItem := range salt {
		saltRule = append(saltRule, saltItem)
	}

	logs, sub, err := _INiftyswapFactory20.contract.WatchLogs(opts, "NewExchange", tokenRule, currencyRule, saltRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapFactory20NewExchange)
				if err := _INiftyswapFactory20.contract.UnpackLog(event, "NewExchange", log); err != nil {
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
// Solidity: event NewExchange(address indexed token, address indexed currency, uint256 indexed salt, address exchange)
func (_INiftyswapFactory20 *INiftyswapFactory20Filterer) ParseNewExchange(log types.Log) (*INiftyswapFactory20NewExchange, error) {
	event := new(INiftyswapFactory20NewExchange)
	if err := _INiftyswapFactory20.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
