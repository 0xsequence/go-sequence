// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale_erc1155

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

// IERC1155SaleFunctionsSaleDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC1155SaleFunctionsSaleDetails struct {
	Cost       *big.Int
	SupplyCap  *big.Int
	StartTime  uint64
	EndTime    uint64
	MerkleRoot [32]byte
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"globalSaleDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC1155SaleFunctions.SaleDetails\",\"components\":[{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxTotal\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"paymentToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenSaleDetails\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC1155SaleFunctions.SaleDetails\",\"components\":[{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GlobalSaleDetailsUpdated\",\"inputs\":[{\"name\":\"cost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSaleDetailsUpdated\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"supplyCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"GlobalSaleInactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPayment\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSupply\",\"inputs\":[{\"name\":\"currentSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requestedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSaleDetails\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenIds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaleInactive\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// SaleABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleMetaData.ABI instead.
var SaleABI = SaleMetaData.ABI

// Sale is an auto generated Go binding around an Ethereum contract.
type Sale struct {
	SaleCaller     // Read-only binding to the contract
	SaleTransactor // Write-only binding to the contract
	SaleFilterer   // Log filterer for contract events
}

// SaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleSession struct {
	Contract     *Sale             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleCallerSession struct {
	Contract *SaleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleTransactorSession struct {
	Contract     *SaleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleRaw struct {
	Contract *Sale // Generic contract binding to access the raw methods on
}

// SaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleCallerRaw struct {
	Contract *SaleCaller // Generic read-only contract binding to access the raw methods on
}

// SaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleTransactorRaw struct {
	Contract *SaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSale creates a new instance of Sale, bound to a specific deployed contract.
func NewSale(address common.Address, backend bind.ContractBackend) (*Sale, error) {
	contract, err := bindSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// NewSaleCaller creates a new read-only instance of Sale, bound to a specific deployed contract.
func NewSaleCaller(address common.Address, caller bind.ContractCaller) (*SaleCaller, error) {
	contract, err := bindSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleCaller{contract: contract}, nil
}

// NewSaleTransactor creates a new write-only instance of Sale, bound to a specific deployed contract.
func NewSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleTransactor, error) {
	contract, err := bindSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleTransactor{contract: contract}, nil
}

// NewSaleFilterer creates a new log filterer instance of Sale, bound to a specific deployed contract.
func NewSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleFilterer, error) {
	contract, err := bindSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleFilterer{contract: contract}, nil
}

// bindSale binds a generic wrapper to an already deployed contract.
func bindSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.SaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transact(opts, method, params...)
}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCaller) GlobalSaleDetails(opts *bind.CallOpts) (IERC1155SaleFunctionsSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "globalSaleDetails")

	if err != nil {
		return *new(IERC1155SaleFunctionsSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC1155SaleFunctionsSaleDetails)).(*IERC1155SaleFunctionsSaleDetails)

	return out0, err

}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleSession) GlobalSaleDetails() (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.GlobalSaleDetails(&_Sale.CallOpts)
}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCallerSession) GlobalSaleDetails() (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.GlobalSaleDetails(&_Sale.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleSession) PaymentToken() (common.Address, error) {
	return _Sale.Contract.PaymentToken(&_Sale.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleCallerSession) PaymentToken() (common.Address, error) {
	return _Sale.Contract.PaymentToken(&_Sale.CallOpts)
}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCaller) TokenSaleDetails(opts *bind.CallOpts, tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenSaleDetails", tokenId)

	if err != nil {
		return *new(IERC1155SaleFunctionsSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC1155SaleFunctionsSaleDetails)).(*IERC1155SaleFunctionsSaleDetails)

	return out0, err

}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleSession) TokenSaleDetails(tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetails(&_Sale.CallOpts, tokenId)
}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCallerSession) TokenSaleDetails(tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetails(&_Sale.CallOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "mint", to, tokenIds, amounts, data, paymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleSession) Mint(to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, tokenIds, amounts, data, paymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactorSession) Mint(to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, tokenIds, amounts, data, paymentToken, maxTotal, proof)
}

// SaleGlobalSaleDetailsUpdatedIterator is returned from FilterGlobalSaleDetailsUpdated and is used to iterate over the raw logs and unpacked data for GlobalSaleDetailsUpdated events raised by the Sale contract.
type SaleGlobalSaleDetailsUpdatedIterator struct {
	Event *SaleGlobalSaleDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *SaleGlobalSaleDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleGlobalSaleDetailsUpdated)
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
		it.Event = new(SaleGlobalSaleDetailsUpdated)
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
func (it *SaleGlobalSaleDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleGlobalSaleDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleGlobalSaleDetailsUpdated represents a GlobalSaleDetailsUpdated event raised by the Sale contract.
type SaleGlobalSaleDetailsUpdated struct {
	Cost       *big.Int
	SupplyCap  *big.Int
	StartTime  uint64
	EndTime    uint64
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGlobalSaleDetailsUpdated is a free log retrieval operation binding the contract event 0x8fd3ac39fbb3d5e9c906dd9ec439dc6e584b8fa3ce02d5b67d589b22b22152a9.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) FilterGlobalSaleDetailsUpdated(opts *bind.FilterOpts) (*SaleGlobalSaleDetailsUpdatedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "GlobalSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SaleGlobalSaleDetailsUpdatedIterator{contract: _Sale.contract, event: "GlobalSaleDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalSaleDetailsUpdated is a free log subscription operation binding the contract event 0x8fd3ac39fbb3d5e9c906dd9ec439dc6e584b8fa3ce02d5b67d589b22b22152a9.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) WatchGlobalSaleDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SaleGlobalSaleDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "GlobalSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleGlobalSaleDetailsUpdated)
				if err := _Sale.contract.UnpackLog(event, "GlobalSaleDetailsUpdated", log); err != nil {
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

// ParseGlobalSaleDetailsUpdated is a log parse operation binding the contract event 0x8fd3ac39fbb3d5e9c906dd9ec439dc6e584b8fa3ce02d5b67d589b22b22152a9.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) ParseGlobalSaleDetailsUpdated(log types.Log) (*SaleGlobalSaleDetailsUpdated, error) {
	event := new(SaleGlobalSaleDetailsUpdated)
	if err := _Sale.contract.UnpackLog(event, "GlobalSaleDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleTokenSaleDetailsUpdatedIterator is returned from FilterTokenSaleDetailsUpdated and is used to iterate over the raw logs and unpacked data for TokenSaleDetailsUpdated events raised by the Sale contract.
type SaleTokenSaleDetailsUpdatedIterator struct {
	Event *SaleTokenSaleDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *SaleTokenSaleDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleTokenSaleDetailsUpdated)
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
		it.Event = new(SaleTokenSaleDetailsUpdated)
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
func (it *SaleTokenSaleDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleTokenSaleDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleTokenSaleDetailsUpdated represents a TokenSaleDetailsUpdated event raised by the Sale contract.
type SaleTokenSaleDetailsUpdated struct {
	TokenId    *big.Int
	Cost       *big.Int
	SupplyCap  *big.Int
	StartTime  uint64
	EndTime    uint64
	MerkleRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenSaleDetailsUpdated is a free log retrieval operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) FilterTokenSaleDetailsUpdated(opts *bind.FilterOpts) (*SaleTokenSaleDetailsUpdatedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "TokenSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SaleTokenSaleDetailsUpdatedIterator{contract: _Sale.contract, event: "TokenSaleDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenSaleDetailsUpdated is a free log subscription operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) WatchTokenSaleDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SaleTokenSaleDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "TokenSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleTokenSaleDetailsUpdated)
				if err := _Sale.contract.UnpackLog(event, "TokenSaleDetailsUpdated", log); err != nil {
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

// ParseTokenSaleDetailsUpdated is a log parse operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 supplyCap, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) ParseTokenSaleDetailsUpdated(log types.Log) (*SaleTokenSaleDetailsUpdated, error) {
	event := new(SaleTokenSaleDetailsUpdated)
	if err := _Sale.contract.UnpackLog(event, "TokenSaleDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
