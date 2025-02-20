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

// INiftyswapExchangeMetaData contains all meta data concerning the INiftyswapExchange contract.
var INiftyswapExchangeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensSoldIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensSoldAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyBoughtAmounts\",\"type\":\"uint256[]\"}],\"name\":\"CurrencyPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyAmounts\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyAmounts\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensBoughtIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensBoughtAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencySoldAmounts\",\"type\":\"uint256[]\"}],\"name\":\"TokensPurchase\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetBoughtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrencyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"getCurrencyReserves\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokensBought\",\"type\":\"uint256[]\"}],\"name\":\"getPrice_currencyToToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokensSold\",\"type\":\"uint256[]\"}],\"name\":\"getPrice_tokenToCurrency\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetSoldAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getSellPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INiftyswapExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use INiftyswapExchangeMetaData.ABI instead.
var INiftyswapExchangeABI = INiftyswapExchangeMetaData.ABI

// INiftyswapExchange is an auto generated Go binding around an Ethereum contract.
type INiftyswapExchange struct {
	INiftyswapExchangeCaller     // Read-only binding to the contract
	INiftyswapExchangeTransactor // Write-only binding to the contract
	INiftyswapExchangeFilterer   // Log filterer for contract events
}

// INiftyswapExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INiftyswapExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INiftyswapExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INiftyswapExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INiftyswapExchangeSession struct {
	Contract     *INiftyswapExchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// INiftyswapExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INiftyswapExchangeCallerSession struct {
	Contract *INiftyswapExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// INiftyswapExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INiftyswapExchangeTransactorSession struct {
	Contract     *INiftyswapExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// INiftyswapExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INiftyswapExchangeRaw struct {
	Contract *INiftyswapExchange // Generic contract binding to access the raw methods on
}

// INiftyswapExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INiftyswapExchangeCallerRaw struct {
	Contract *INiftyswapExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// INiftyswapExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INiftyswapExchangeTransactorRaw struct {
	Contract *INiftyswapExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINiftyswapExchange creates a new instance of INiftyswapExchange, bound to a specific deployed contract.
func NewINiftyswapExchange(address common.Address, backend bind.ContractBackend) (*INiftyswapExchange, error) {
	contract, err := bindINiftyswapExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange{INiftyswapExchangeCaller: INiftyswapExchangeCaller{contract: contract}, INiftyswapExchangeTransactor: INiftyswapExchangeTransactor{contract: contract}, INiftyswapExchangeFilterer: INiftyswapExchangeFilterer{contract: contract}}, nil
}

// NewINiftyswapExchangeCaller creates a new read-only instance of INiftyswapExchange, bound to a specific deployed contract.
func NewINiftyswapExchangeCaller(address common.Address, caller bind.ContractCaller) (*INiftyswapExchangeCaller, error) {
	contract, err := bindINiftyswapExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeCaller{contract: contract}, nil
}

// NewINiftyswapExchangeTransactor creates a new write-only instance of INiftyswapExchange, bound to a specific deployed contract.
func NewINiftyswapExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*INiftyswapExchangeTransactor, error) {
	contract, err := bindINiftyswapExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeTransactor{contract: contract}, nil
}

// NewINiftyswapExchangeFilterer creates a new log filterer instance of INiftyswapExchange, bound to a specific deployed contract.
func NewINiftyswapExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*INiftyswapExchangeFilterer, error) {
	contract, err := bindINiftyswapExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeFilterer{contract: contract}, nil
}

// bindINiftyswapExchange binds a generic wrapper to an already deployed contract.
func bindINiftyswapExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INiftyswapExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapExchange *INiftyswapExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapExchange.Contract.INiftyswapExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapExchange *INiftyswapExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.INiftyswapExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapExchange *INiftyswapExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.INiftyswapExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapExchange *INiftyswapExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapExchange *INiftyswapExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapExchange *INiftyswapExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.contract.Transact(opts, method, params...)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetBuyPrice(opts *bind.CallOpts, _assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getBuyPrice", _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeSession) GetBuyPrice(_assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange.Contract.GetBuyPrice(&_INiftyswapExchange.CallOpts, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetBuyPrice(_assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange.Contract.GetBuyPrice(&_INiftyswapExchange.CallOpts, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address, uint256)
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetCurrencyInfo(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getCurrencyInfo")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address, uint256)
func (_INiftyswapExchange *INiftyswapExchangeSession) GetCurrencyInfo() (common.Address, *big.Int, error) {
	return _INiftyswapExchange.Contract.GetCurrencyInfo(&_INiftyswapExchange.CallOpts)
}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address, uint256)
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetCurrencyInfo() (common.Address, *big.Int, error) {
	return _INiftyswapExchange.Contract.GetCurrencyInfo(&_INiftyswapExchange.CallOpts)
}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetCurrencyReserves(opts *bind.CallOpts, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getCurrencyReserves", _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeSession) GetCurrencyReserves(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetCurrencyReserves(&_INiftyswapExchange.CallOpts, _ids)
}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetCurrencyReserves(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetCurrencyReserves(&_INiftyswapExchange.CallOpts, _ids)
}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeSession) GetFactoryAddress() (common.Address, error) {
	return _INiftyswapExchange.Contract.GetFactoryAddress(&_INiftyswapExchange.CallOpts)
}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetFactoryAddress() (common.Address, error) {
	return _INiftyswapExchange.Contract.GetFactoryAddress(&_INiftyswapExchange.CallOpts)
}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetPriceCurrencyToToken(opts *bind.CallOpts, _ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getPrice_currencyToToken", _ids, _tokensBought)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeSession) GetPriceCurrencyToToken(_ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetPriceCurrencyToToken(&_INiftyswapExchange.CallOpts, _ids, _tokensBought)
}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetPriceCurrencyToToken(_ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetPriceCurrencyToToken(&_INiftyswapExchange.CallOpts, _ids, _tokensBought)
}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetPriceTokenToCurrency(opts *bind.CallOpts, _ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getPrice_tokenToCurrency", _ids, _tokensSold)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeSession) GetPriceTokenToCurrency(_ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetPriceTokenToCurrency(&_INiftyswapExchange.CallOpts, _ids, _tokensSold)
}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetPriceTokenToCurrency(_ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetPriceTokenToCurrency(&_INiftyswapExchange.CallOpts, _ids, _tokensSold)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetSellPrice(opts *bind.CallOpts, _assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getSellPrice", _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeSession) GetSellPrice(_assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange.Contract.GetSellPrice(&_INiftyswapExchange.CallOpts, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetSellPrice(_assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange.Contract.GetSellPrice(&_INiftyswapExchange.CallOpts, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeSession) GetTokenAddress() (common.Address, error) {
	return _INiftyswapExchange.Contract.GetTokenAddress(&_INiftyswapExchange.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetTokenAddress() (common.Address, error) {
	return _INiftyswapExchange.Contract.GetTokenAddress(&_INiftyswapExchange.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCaller) GetTotalSupply(opts *bind.CallOpts, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange.contract.Call(opts, &out, "getTotalSupply", _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeSession) GetTotalSupply(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetTotalSupply(&_INiftyswapExchange.CallOpts, _ids)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange *INiftyswapExchangeCallerSession) GetTotalSupply(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange.Contract.GetTotalSupply(&_INiftyswapExchange.CallOpts, _ids)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.contract.Transact(opts, "onERC1155BatchReceived", arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.OnERC1155BatchReceived(&_INiftyswapExchange.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.OnERC1155BatchReceived(&_INiftyswapExchange.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeTransactor) OnERC1155Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.contract.Transact(opts, "onERC1155Received", _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.OnERC1155Received(&_INiftyswapExchange.TransactOpts, _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange *INiftyswapExchangeTransactorSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange.Contract.OnERC1155Received(&_INiftyswapExchange.TransactOpts, _operator, _from, _id, _amount, _data)
}

// INiftyswapExchangeCurrencyPurchaseIterator is returned from FilterCurrencyPurchase and is used to iterate over the raw logs and unpacked data for CurrencyPurchase events raised by the INiftyswapExchange contract.
type INiftyswapExchangeCurrencyPurchaseIterator struct {
	Event *INiftyswapExchangeCurrencyPurchase // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchangeCurrencyPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchangeCurrencyPurchase)
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
		it.Event = new(INiftyswapExchangeCurrencyPurchase)
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
func (it *INiftyswapExchangeCurrencyPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchangeCurrencyPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchangeCurrencyPurchase represents a CurrencyPurchase event raised by the INiftyswapExchange contract.
type INiftyswapExchangeCurrencyPurchase struct {
	Buyer                 common.Address
	Recipient             common.Address
	TokensSoldIds         []*big.Int
	TokensSoldAmounts     []*big.Int
	CurrencyBoughtAmounts []*big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCurrencyPurchase is a free log retrieval operation binding the contract event 0x89e4dbdd48f69e7920342e9ad9691b9a7150f254e6a0af177ccfd2556aab8bcd.
//
// Solidity: event CurrencyPurchase(address indexed buyer, address indexed recipient, uint256[] tokensSoldIds, uint256[] tokensSoldAmounts, uint256[] currencyBoughtAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) FilterCurrencyPurchase(opts *bind.FilterOpts, buyer []common.Address, recipient []common.Address) (*INiftyswapExchangeCurrencyPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.FilterLogs(opts, "CurrencyPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeCurrencyPurchaseIterator{contract: _INiftyswapExchange.contract, event: "CurrencyPurchase", logs: logs, sub: sub}, nil
}

// WatchCurrencyPurchase is a free log subscription operation binding the contract event 0x89e4dbdd48f69e7920342e9ad9691b9a7150f254e6a0af177ccfd2556aab8bcd.
//
// Solidity: event CurrencyPurchase(address indexed buyer, address indexed recipient, uint256[] tokensSoldIds, uint256[] tokensSoldAmounts, uint256[] currencyBoughtAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) WatchCurrencyPurchase(opts *bind.WatchOpts, sink chan<- *INiftyswapExchangeCurrencyPurchase, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.WatchLogs(opts, "CurrencyPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchangeCurrencyPurchase)
				if err := _INiftyswapExchange.contract.UnpackLog(event, "CurrencyPurchase", log); err != nil {
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

// ParseCurrencyPurchase is a log parse operation binding the contract event 0x89e4dbdd48f69e7920342e9ad9691b9a7150f254e6a0af177ccfd2556aab8bcd.
//
// Solidity: event CurrencyPurchase(address indexed buyer, address indexed recipient, uint256[] tokensSoldIds, uint256[] tokensSoldAmounts, uint256[] currencyBoughtAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) ParseCurrencyPurchase(log types.Log) (*INiftyswapExchangeCurrencyPurchase, error) {
	event := new(INiftyswapExchangeCurrencyPurchase)
	if err := _INiftyswapExchange.contract.UnpackLog(event, "CurrencyPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchangeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the INiftyswapExchange contract.
type INiftyswapExchangeLiquidityAddedIterator struct {
	Event *INiftyswapExchangeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchangeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchangeLiquidityAdded)
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
		it.Event = new(INiftyswapExchangeLiquidityAdded)
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
func (it *INiftyswapExchangeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchangeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchangeLiquidityAdded represents a LiquidityAdded event raised by the INiftyswapExchange contract.
type INiftyswapExchangeLiquidityAdded struct {
	Provider        common.Address
	TokenIds        []*big.Int
	TokenAmounts    []*big.Int
	CurrencyAmounts []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x403f9dc4582dae52d3eeb4a22d37540ffb13c32d964c92ec5ac0d3d5628da316.
//
// Solidity: event LiquidityAdded(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*INiftyswapExchangeLiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeLiquidityAddedIterator{contract: _INiftyswapExchange.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x403f9dc4582dae52d3eeb4a22d37540ffb13c32d964c92ec5ac0d3d5628da316.
//
// Solidity: event LiquidityAdded(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *INiftyswapExchangeLiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchangeLiquidityAdded)
				if err := _INiftyswapExchange.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x403f9dc4582dae52d3eeb4a22d37540ffb13c32d964c92ec5ac0d3d5628da316.
//
// Solidity: event LiquidityAdded(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) ParseLiquidityAdded(log types.Log) (*INiftyswapExchangeLiquidityAdded, error) {
	event := new(INiftyswapExchangeLiquidityAdded)
	if err := _INiftyswapExchange.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchangeLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the INiftyswapExchange contract.
type INiftyswapExchangeLiquidityRemovedIterator struct {
	Event *INiftyswapExchangeLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchangeLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchangeLiquidityRemoved)
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
		it.Event = new(INiftyswapExchangeLiquidityRemoved)
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
func (it *INiftyswapExchangeLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchangeLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchangeLiquidityRemoved represents a LiquidityRemoved event raised by the INiftyswapExchange contract.
type INiftyswapExchangeLiquidityRemoved struct {
	Provider        common.Address
	TokenIds        []*big.Int
	TokenAmounts    []*big.Int
	CurrencyAmounts []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x711e9bcb94b4cf7bc99c1cb938edc75ac7e85a136838e90abf6ee1f5adebd423.
//
// Solidity: event LiquidityRemoved(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address) (*INiftyswapExchangeLiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.FilterLogs(opts, "LiquidityRemoved", providerRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeLiquidityRemovedIterator{contract: _INiftyswapExchange.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x711e9bcb94b4cf7bc99c1cb938edc75ac7e85a136838e90abf6ee1f5adebd423.
//
// Solidity: event LiquidityRemoved(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *INiftyswapExchangeLiquidityRemoved, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.WatchLogs(opts, "LiquidityRemoved", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchangeLiquidityRemoved)
				if err := _INiftyswapExchange.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0x711e9bcb94b4cf7bc99c1cb938edc75ac7e85a136838e90abf6ee1f5adebd423.
//
// Solidity: event LiquidityRemoved(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) ParseLiquidityRemoved(log types.Log) (*INiftyswapExchangeLiquidityRemoved, error) {
	event := new(INiftyswapExchangeLiquidityRemoved)
	if err := _INiftyswapExchange.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchangeTokensPurchaseIterator is returned from FilterTokensPurchase and is used to iterate over the raw logs and unpacked data for TokensPurchase events raised by the INiftyswapExchange contract.
type INiftyswapExchangeTokensPurchaseIterator struct {
	Event *INiftyswapExchangeTokensPurchase // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchangeTokensPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchangeTokensPurchase)
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
		it.Event = new(INiftyswapExchangeTokensPurchase)
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
func (it *INiftyswapExchangeTokensPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchangeTokensPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchangeTokensPurchase represents a TokensPurchase event raised by the INiftyswapExchange contract.
type INiftyswapExchangeTokensPurchase struct {
	Buyer               common.Address
	Recipient           common.Address
	TokensBoughtIds     []*big.Int
	TokensBoughtAmounts []*big.Int
	CurrencySoldAmounts []*big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchase is a free log retrieval operation binding the contract event 0xd38bc77e62e239476b3e25620d73f29a4a188e808aad79f4a81aaf44871a7a30.
//
// Solidity: event TokensPurchase(address indexed buyer, address indexed recipient, uint256[] tokensBoughtIds, uint256[] tokensBoughtAmounts, uint256[] currencySoldAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) FilterTokensPurchase(opts *bind.FilterOpts, buyer []common.Address, recipient []common.Address) (*INiftyswapExchangeTokensPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.FilterLogs(opts, "TokensPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchangeTokensPurchaseIterator{contract: _INiftyswapExchange.contract, event: "TokensPurchase", logs: logs, sub: sub}, nil
}

// WatchTokensPurchase is a free log subscription operation binding the contract event 0xd38bc77e62e239476b3e25620d73f29a4a188e808aad79f4a81aaf44871a7a30.
//
// Solidity: event TokensPurchase(address indexed buyer, address indexed recipient, uint256[] tokensBoughtIds, uint256[] tokensBoughtAmounts, uint256[] currencySoldAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) WatchTokensPurchase(opts *bind.WatchOpts, sink chan<- *INiftyswapExchangeTokensPurchase, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange.contract.WatchLogs(opts, "TokensPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchangeTokensPurchase)
				if err := _INiftyswapExchange.contract.UnpackLog(event, "TokensPurchase", log); err != nil {
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

// ParseTokensPurchase is a log parse operation binding the contract event 0xd38bc77e62e239476b3e25620d73f29a4a188e808aad79f4a81aaf44871a7a30.
//
// Solidity: event TokensPurchase(address indexed buyer, address indexed recipient, uint256[] tokensBoughtIds, uint256[] tokensBoughtAmounts, uint256[] currencySoldAmounts)
func (_INiftyswapExchange *INiftyswapExchangeFilterer) ParseTokensPurchase(log types.Log) (*INiftyswapExchangeTokensPurchase, error) {
	event := new(INiftyswapExchangeTokensPurchase)
	if err := _INiftyswapExchange.contract.UnpackLog(event, "TokensPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
