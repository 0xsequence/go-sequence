// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sale

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

// IERC721SaleFunctionsSaleDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC721SaleFunctionsSaleDetails struct {
	SupplyCap    *big.Int
	Cost         *big.Int
	PaymentToken common.Address
	StartTime    uint64
	EndTime      uint64
	MerkleRoot   [32]byte
}

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxTotal\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"saleDetails\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIERC721SaleFunctions.SaleDetails\",\"components\":[{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setSaleDetails\",\"inputs\":[{\"name\":\"supplyCap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SaleDetailsUpdated\",\"inputs\":[{\"name\":\"supplyCap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"endTime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InsufficientPayment\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSupply\",\"inputs\":[{\"name\":\"currentSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSaleDetails\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaleInactive\",\"inputs\":[]}]",
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

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((uint256,uint256,address,uint64,uint64,bytes32))
func (_Sale *SaleCaller) SaleDetails(opts *bind.CallOpts) (IERC721SaleFunctionsSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "saleDetails")

	if err != nil {
		return *new(IERC721SaleFunctionsSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721SaleFunctionsSaleDetails)).(*IERC721SaleFunctionsSaleDetails)

	return out0, err

}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((uint256,uint256,address,uint64,uint64,bytes32))
func (_Sale *SaleSession) SaleDetails() (IERC721SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.SaleDetails(&_Sale.CallOpts)
}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((uint256,uint256,address,uint64,uint64,bytes32))
func (_Sale *SaleCallerSession) SaleDetails() (IERC721SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.SaleDetails(&_Sale.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x0668d0bb.
//
// Solidity: function mint(address to, uint256 amount, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "mint", to, amount, paymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x0668d0bb.
//
// Solidity: function mint(address to, uint256 amount, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleSession) Mint(to common.Address, amount *big.Int, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, amount, paymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x0668d0bb.
//
// Solidity: function mint(address to, uint256 amount, address paymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactorSession) Mint(to common.Address, amount *big.Int, paymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, amount, paymentToken, maxTotal, proof)
}

// SetSaleDetails is a paid mutator transaction binding the contract method 0x8c17030f.
//
// Solidity: function setSaleDetails(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactor) SetSaleDetails(opts *bind.TransactOpts, supplyCap *big.Int, cost *big.Int, paymentToken common.Address, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setSaleDetails", supplyCap, cost, paymentToken, startTime, endTime, merkleRoot)
}

// SetSaleDetails is a paid mutator transaction binding the contract method 0x8c17030f.
//
// Solidity: function setSaleDetails(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleSession) SetSaleDetails(supplyCap *big.Int, cost *big.Int, paymentToken common.Address, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetSaleDetails(&_Sale.TransactOpts, supplyCap, cost, paymentToken, startTime, endTime, merkleRoot)
}

// SetSaleDetails is a paid mutator transaction binding the contract method 0x8c17030f.
//
// Solidity: function setSaleDetails(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactorSession) SetSaleDetails(supplyCap *big.Int, cost *big.Int, paymentToken common.Address, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetSaleDetails(&_Sale.TransactOpts, supplyCap, cost, paymentToken, startTime, endTime, merkleRoot)
}

// SaleSaleDetailsUpdatedIterator is returned from FilterSaleDetailsUpdated and is used to iterate over the raw logs and unpacked data for SaleDetailsUpdated events raised by the Sale contract.
type SaleSaleDetailsUpdatedIterator struct {
	Event *SaleSaleDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *SaleSaleDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleSaleDetailsUpdated)
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
		it.Event = new(SaleSaleDetailsUpdated)
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
func (it *SaleSaleDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleSaleDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleSaleDetailsUpdated represents a SaleDetailsUpdated event raised by the Sale contract.
type SaleSaleDetailsUpdated struct {
	SupplyCap    *big.Int
	Cost         *big.Int
	PaymentToken common.Address
	StartTime    uint64
	EndTime      uint64
	MerkleRoot   [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSaleDetailsUpdated is a free log retrieval operation binding the contract event 0xabec13ca1773eed55d54d2f64593c33fa520ee45cac73a162f13928a2ebee233.
//
// Solidity: event SaleDetailsUpdated(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) FilterSaleDetailsUpdated(opts *bind.FilterOpts) (*SaleSaleDetailsUpdatedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "SaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SaleSaleDetailsUpdatedIterator{contract: _Sale.contract, event: "SaleDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchSaleDetailsUpdated is a free log subscription operation binding the contract event 0xabec13ca1773eed55d54d2f64593c33fa520ee45cac73a162f13928a2ebee233.
//
// Solidity: event SaleDetailsUpdated(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) WatchSaleDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SaleSaleDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "SaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleSaleDetailsUpdated)
				if err := _Sale.contract.UnpackLog(event, "SaleDetailsUpdated", log); err != nil {
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

// ParseSaleDetailsUpdated is a log parse operation binding the contract event 0xabec13ca1773eed55d54d2f64593c33fa520ee45cac73a162f13928a2ebee233.
//
// Solidity: event SaleDetailsUpdated(uint256 supplyCap, uint256 cost, address paymentToken, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) ParseSaleDetailsUpdated(log types.Log) (*SaleSaleDetailsUpdated, error) {
	event := new(SaleSaleDetailsUpdated)
	if err := _Sale.contract.UnpackLog(event, "SaleDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
