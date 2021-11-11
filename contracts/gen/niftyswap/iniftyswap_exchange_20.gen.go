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

// INiftyswapExchange20MetaData contains all meta data concerning the INiftyswapExchange20 contract.
var INiftyswapExchange20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensSoldIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensSoldAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyBoughtAmounts\",\"type\":\"uint256[]\"}],\"name\":\"CurrencyPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyAmounts\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencyAmounts\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyFee\",\"type\":\"uint256\"}],\"name\":\"RoyaltyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensBoughtIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokensBoughtAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"currencySoldAmounts\",\"type\":\"uint256[]\"}],\"name\":\"TokensPurchase\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokensBoughtAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxCurrency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_extraFeeRecipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_extraFeeAmounts\",\"type\":\"uint256[]\"}],\"name\":\"buyTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetBoughtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getBuyPriceWithRoyalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrencyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"getCurrencyReserves\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalRoyaltyFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalRoyaltyRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokensBought\",\"type\":\"uint256[]\"}],\"name\":\"getPrice_currencyToToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokensSold\",\"type\":\"uint256[]\"}],\"name\":\"getPrice_tokenToCurrency\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"}],\"name\":\"getRoyalties\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assetSoldAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getSellPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetSoldReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assetBoughtReserve\",\"type\":\"uint256\"}],\"name\":\"getSellPriceWithRoyalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"}],\"name\":\"sendRoyalties\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// INiftyswapExchange20ABI is the input ABI used to generate the binding from.
// Deprecated: Use INiftyswapExchange20MetaData.ABI instead.
var INiftyswapExchange20ABI = INiftyswapExchange20MetaData.ABI

// INiftyswapExchange20 is an auto generated Go binding around an Ethereum contract.
type INiftyswapExchange20 struct {
	INiftyswapExchange20Caller     // Read-only binding to the contract
	INiftyswapExchange20Transactor // Write-only binding to the contract
	INiftyswapExchange20Filterer   // Log filterer for contract events
}

// INiftyswapExchange20Caller is an auto generated read-only Go binding around an Ethereum contract.
type INiftyswapExchange20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchange20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type INiftyswapExchange20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchange20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INiftyswapExchange20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INiftyswapExchange20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INiftyswapExchange20Session struct {
	Contract     *INiftyswapExchange20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// INiftyswapExchange20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INiftyswapExchange20CallerSession struct {
	Contract *INiftyswapExchange20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// INiftyswapExchange20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INiftyswapExchange20TransactorSession struct {
	Contract     *INiftyswapExchange20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// INiftyswapExchange20Raw is an auto generated low-level Go binding around an Ethereum contract.
type INiftyswapExchange20Raw struct {
	Contract *INiftyswapExchange20 // Generic contract binding to access the raw methods on
}

// INiftyswapExchange20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INiftyswapExchange20CallerRaw struct {
	Contract *INiftyswapExchange20Caller // Generic read-only contract binding to access the raw methods on
}

// INiftyswapExchange20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INiftyswapExchange20TransactorRaw struct {
	Contract *INiftyswapExchange20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewINiftyswapExchange20 creates a new instance of INiftyswapExchange20, bound to a specific deployed contract.
func NewINiftyswapExchange20(address common.Address, backend bind.ContractBackend) (*INiftyswapExchange20, error) {
	contract, err := bindINiftyswapExchange20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20{INiftyswapExchange20Caller: INiftyswapExchange20Caller{contract: contract}, INiftyswapExchange20Transactor: INiftyswapExchange20Transactor{contract: contract}, INiftyswapExchange20Filterer: INiftyswapExchange20Filterer{contract: contract}}, nil
}

// NewINiftyswapExchange20Caller creates a new read-only instance of INiftyswapExchange20, bound to a specific deployed contract.
func NewINiftyswapExchange20Caller(address common.Address, caller bind.ContractCaller) (*INiftyswapExchange20Caller, error) {
	contract, err := bindINiftyswapExchange20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20Caller{contract: contract}, nil
}

// NewINiftyswapExchange20Transactor creates a new write-only instance of INiftyswapExchange20, bound to a specific deployed contract.
func NewINiftyswapExchange20Transactor(address common.Address, transactor bind.ContractTransactor) (*INiftyswapExchange20Transactor, error) {
	contract, err := bindINiftyswapExchange20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20Transactor{contract: contract}, nil
}

// NewINiftyswapExchange20Filterer creates a new log filterer instance of INiftyswapExchange20, bound to a specific deployed contract.
func NewINiftyswapExchange20Filterer(address common.Address, filterer bind.ContractFilterer) (*INiftyswapExchange20Filterer, error) {
	contract, err := bindINiftyswapExchange20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20Filterer{contract: contract}, nil
}

// bindINiftyswapExchange20 binds a generic wrapper to an already deployed contract.
func bindINiftyswapExchange20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INiftyswapExchange20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapExchange20 *INiftyswapExchange20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapExchange20.Contract.INiftyswapExchange20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapExchange20 *INiftyswapExchange20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.INiftyswapExchange20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapExchange20 *INiftyswapExchange20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.INiftyswapExchange20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INiftyswapExchange20 *INiftyswapExchange20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INiftyswapExchange20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.contract.Transact(opts, method, params...)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetBuyPrice(opts *bind.CallOpts, _assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getBuyPrice", _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetBuyPrice(_assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetBuyPrice(&_INiftyswapExchange20.CallOpts, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0xfca16c3b.
//
// Solidity: function getBuyPrice(uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetBuyPrice(_assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetBuyPrice(&_INiftyswapExchange20.CallOpts, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetBuyPriceWithRoyalty is a free data retrieval call binding the contract method 0xaeaad208.
//
// Solidity: function getBuyPriceWithRoyalty(uint256 _tokenId, uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetBuyPriceWithRoyalty(opts *bind.CallOpts, _tokenId *big.Int, _assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getBuyPriceWithRoyalty", _tokenId, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPriceWithRoyalty is a free data retrieval call binding the contract method 0xaeaad208.
//
// Solidity: function getBuyPriceWithRoyalty(uint256 _tokenId, uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetBuyPriceWithRoyalty(_tokenId *big.Int, _assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetBuyPriceWithRoyalty(&_INiftyswapExchange20.CallOpts, _tokenId, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetBuyPriceWithRoyalty is a free data retrieval call binding the contract method 0xaeaad208.
//
// Solidity: function getBuyPriceWithRoyalty(uint256 _tokenId, uint256 _assetBoughtAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetBuyPriceWithRoyalty(_tokenId *big.Int, _assetBoughtAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetBuyPriceWithRoyalty(&_INiftyswapExchange20.CallOpts, _tokenId, _assetBoughtAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetCurrencyInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getCurrencyInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetCurrencyInfo() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetCurrencyInfo(&_INiftyswapExchange20.CallOpts)
}

// GetCurrencyInfo is a free data retrieval call binding the contract method 0x46adf5ca.
//
// Solidity: function getCurrencyInfo() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetCurrencyInfo() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetCurrencyInfo(&_INiftyswapExchange20.CallOpts)
}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetCurrencyReserves(opts *bind.CallOpts, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getCurrencyReserves", _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetCurrencyReserves(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetCurrencyReserves(&_INiftyswapExchange20.CallOpts, _ids)
}

// GetCurrencyReserves is a free data retrieval call binding the contract method 0x209b96c5.
//
// Solidity: function getCurrencyReserves(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetCurrencyReserves(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetCurrencyReserves(&_INiftyswapExchange20.CallOpts, _ids)
}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetFactoryAddress() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetFactoryAddress(&_INiftyswapExchange20.CallOpts)
}

// GetFactoryAddress is a free data retrieval call binding the contract method 0xa9c2e36c.
//
// Solidity: function getFactoryAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetFactoryAddress() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetFactoryAddress(&_INiftyswapExchange20.CallOpts)
}

// GetGlobalRoyaltyFee is a free data retrieval call binding the contract method 0x008e09d8.
//
// Solidity: function getGlobalRoyaltyFee() view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetGlobalRoyaltyFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getGlobalRoyaltyFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalRoyaltyFee is a free data retrieval call binding the contract method 0x008e09d8.
//
// Solidity: function getGlobalRoyaltyFee() view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetGlobalRoyaltyFee() (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetGlobalRoyaltyFee(&_INiftyswapExchange20.CallOpts)
}

// GetGlobalRoyaltyFee is a free data retrieval call binding the contract method 0x008e09d8.
//
// Solidity: function getGlobalRoyaltyFee() view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetGlobalRoyaltyFee() (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetGlobalRoyaltyFee(&_INiftyswapExchange20.CallOpts)
}

// GetGlobalRoyaltyRecipient is a free data retrieval call binding the contract method 0xe523d3fc.
//
// Solidity: function getGlobalRoyaltyRecipient() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetGlobalRoyaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getGlobalRoyaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGlobalRoyaltyRecipient is a free data retrieval call binding the contract method 0xe523d3fc.
//
// Solidity: function getGlobalRoyaltyRecipient() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetGlobalRoyaltyRecipient() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetGlobalRoyaltyRecipient(&_INiftyswapExchange20.CallOpts)
}

// GetGlobalRoyaltyRecipient is a free data retrieval call binding the contract method 0xe523d3fc.
//
// Solidity: function getGlobalRoyaltyRecipient() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetGlobalRoyaltyRecipient() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetGlobalRoyaltyRecipient(&_INiftyswapExchange20.CallOpts)
}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetPriceCurrencyToToken(opts *bind.CallOpts, _ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getPrice_currencyToToken", _ids, _tokensBought)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetPriceCurrencyToToken(_ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetPriceCurrencyToToken(&_INiftyswapExchange20.CallOpts, _ids, _tokensBought)
}

// GetPriceCurrencyToToken is a free data retrieval call binding the contract method 0xbe571468.
//
// Solidity: function getPrice_currencyToToken(uint256[] _ids, uint256[] _tokensBought) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetPriceCurrencyToToken(_ids []*big.Int, _tokensBought []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetPriceCurrencyToToken(&_INiftyswapExchange20.CallOpts, _ids, _tokensBought)
}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetPriceTokenToCurrency(opts *bind.CallOpts, _ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getPrice_tokenToCurrency", _ids, _tokensSold)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetPriceTokenToCurrency(_ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetPriceTokenToCurrency(&_INiftyswapExchange20.CallOpts, _ids, _tokensSold)
}

// GetPriceTokenToCurrency is a free data retrieval call binding the contract method 0x863ed300.
//
// Solidity: function getPrice_tokenToCurrency(uint256[] _ids, uint256[] _tokensSold) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetPriceTokenToCurrency(_ids []*big.Int, _tokensSold []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetPriceTokenToCurrency(&_INiftyswapExchange20.CallOpts, _ids, _tokensSold)
}

// GetRoyalties is a free data retrieval call binding the contract method 0x14556a56.
//
// Solidity: function getRoyalties(address _royaltyRecipient) view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetRoyalties(opts *bind.CallOpts, _royaltyRecipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getRoyalties", _royaltyRecipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoyalties is a free data retrieval call binding the contract method 0x14556a56.
//
// Solidity: function getRoyalties(address _royaltyRecipient) view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetRoyalties(_royaltyRecipient common.Address) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetRoyalties(&_INiftyswapExchange20.CallOpts, _royaltyRecipient)
}

// GetRoyalties is a free data retrieval call binding the contract method 0x14556a56.
//
// Solidity: function getRoyalties(address _royaltyRecipient) view returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetRoyalties(_royaltyRecipient common.Address) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetRoyalties(&_INiftyswapExchange20.CallOpts, _royaltyRecipient)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetSellPrice(opts *bind.CallOpts, _assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getSellPrice", _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetSellPrice(_assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetSellPrice(&_INiftyswapExchange20.CallOpts, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetSellPrice is a free data retrieval call binding the contract method 0x6ee8e134.
//
// Solidity: function getSellPrice(uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) pure returns(uint256)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetSellPrice(_assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetSellPrice(&_INiftyswapExchange20.CallOpts, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetSellPriceWithRoyalty is a free data retrieval call binding the contract method 0xa7380f6e.
//
// Solidity: function getSellPriceWithRoyalty(uint256 _tokenId, uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetSellPriceWithRoyalty(opts *bind.CallOpts, _tokenId *big.Int, _assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getSellPriceWithRoyalty", _tokenId, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellPriceWithRoyalty is a free data retrieval call binding the contract method 0xa7380f6e.
//
// Solidity: function getSellPriceWithRoyalty(uint256 _tokenId, uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetSellPriceWithRoyalty(_tokenId *big.Int, _assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetSellPriceWithRoyalty(&_INiftyswapExchange20.CallOpts, _tokenId, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetSellPriceWithRoyalty is a free data retrieval call binding the contract method 0xa7380f6e.
//
// Solidity: function getSellPriceWithRoyalty(uint256 _tokenId, uint256 _assetSoldAmount, uint256 _assetSoldReserve, uint256 _assetBoughtReserve) view returns(uint256 price)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetSellPriceWithRoyalty(_tokenId *big.Int, _assetSoldAmount *big.Int, _assetSoldReserve *big.Int, _assetBoughtReserve *big.Int) (*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetSellPriceWithRoyalty(&_INiftyswapExchange20.CallOpts, _tokenId, _assetSoldAmount, _assetSoldReserve, _assetBoughtReserve)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetTokenAddress() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetTokenAddress(&_INiftyswapExchange20.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetTokenAddress() (common.Address, error) {
	return _INiftyswapExchange20.Contract.GetTokenAddress(&_INiftyswapExchange20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Caller) GetTotalSupply(opts *bind.CallOpts, _ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _INiftyswapExchange20.contract.Call(opts, &out, "getTotalSupply", _ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Session) GetTotalSupply(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetTotalSupply(&_INiftyswapExchange20.CallOpts, _ids)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0x2bef5e38.
//
// Solidity: function getTotalSupply(uint256[] _ids) view returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20CallerSession) GetTotalSupply(_ids []*big.Int) ([]*big.Int, error) {
	return _INiftyswapExchange20.Contract.GetTotalSupply(&_INiftyswapExchange20.CallOpts, _ids)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd93e8aaa.
//
// Solidity: function buyTokens(uint256[] _tokenIds, uint256[] _tokensBoughtAmounts, uint256 _maxCurrency, uint256 _deadline, address _recipient, address[] _extraFeeRecipients, uint256[] _extraFeeAmounts) returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Transactor) BuyTokens(opts *bind.TransactOpts, _tokenIds []*big.Int, _tokensBoughtAmounts []*big.Int, _maxCurrency *big.Int, _deadline *big.Int, _recipient common.Address, _extraFeeRecipients []common.Address, _extraFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _INiftyswapExchange20.contract.Transact(opts, "buyTokens", _tokenIds, _tokensBoughtAmounts, _maxCurrency, _deadline, _recipient, _extraFeeRecipients, _extraFeeAmounts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd93e8aaa.
//
// Solidity: function buyTokens(uint256[] _tokenIds, uint256[] _tokensBoughtAmounts, uint256 _maxCurrency, uint256 _deadline, address _recipient, address[] _extraFeeRecipients, uint256[] _extraFeeAmounts) returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20Session) BuyTokens(_tokenIds []*big.Int, _tokensBoughtAmounts []*big.Int, _maxCurrency *big.Int, _deadline *big.Int, _recipient common.Address, _extraFeeRecipients []common.Address, _extraFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.BuyTokens(&_INiftyswapExchange20.TransactOpts, _tokenIds, _tokensBoughtAmounts, _maxCurrency, _deadline, _recipient, _extraFeeRecipients, _extraFeeAmounts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xd93e8aaa.
//
// Solidity: function buyTokens(uint256[] _tokenIds, uint256[] _tokensBoughtAmounts, uint256 _maxCurrency, uint256 _deadline, address _recipient, address[] _extraFeeRecipients, uint256[] _extraFeeAmounts) returns(uint256[])
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorSession) BuyTokens(_tokenIds []*big.Int, _tokensBoughtAmounts []*big.Int, _maxCurrency *big.Int, _deadline *big.Int, _recipient common.Address, _extraFeeRecipients []common.Address, _extraFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.BuyTokens(&_INiftyswapExchange20.TransactOpts, _tokenIds, _tokensBoughtAmounts, _maxCurrency, _deadline, _recipient, _extraFeeRecipients, _extraFeeAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20Transactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.contract.Transact(opts, "onERC1155BatchReceived", arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.OnERC1155BatchReceived(&_INiftyswapExchange20.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.OnERC1155BatchReceived(&_INiftyswapExchange20.TransactOpts, arg0, _from, _ids, _amounts, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20Transactor) OnERC1155Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.contract.Transact(opts, "onERC1155Received", _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20Session) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.OnERC1155Received(&_INiftyswapExchange20.TransactOpts, _operator, _from, _id, _amount, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address _operator, address _from, uint256 _id, uint256 _amount, bytes _data) returns(bytes4)
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorSession) OnERC1155Received(_operator common.Address, _from common.Address, _id *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.OnERC1155Received(&_INiftyswapExchange20.TransactOpts, _operator, _from, _id, _amount, _data)
}

// SendRoyalties is a paid mutator transaction binding the contract method 0xa6313875.
//
// Solidity: function sendRoyalties(address _royaltyRecipient) returns()
func (_INiftyswapExchange20 *INiftyswapExchange20Transactor) SendRoyalties(opts *bind.TransactOpts, _royaltyRecipient common.Address) (*types.Transaction, error) {
	return _INiftyswapExchange20.contract.Transact(opts, "sendRoyalties", _royaltyRecipient)
}

// SendRoyalties is a paid mutator transaction binding the contract method 0xa6313875.
//
// Solidity: function sendRoyalties(address _royaltyRecipient) returns()
func (_INiftyswapExchange20 *INiftyswapExchange20Session) SendRoyalties(_royaltyRecipient common.Address) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.SendRoyalties(&_INiftyswapExchange20.TransactOpts, _royaltyRecipient)
}

// SendRoyalties is a paid mutator transaction binding the contract method 0xa6313875.
//
// Solidity: function sendRoyalties(address _royaltyRecipient) returns()
func (_INiftyswapExchange20 *INiftyswapExchange20TransactorSession) SendRoyalties(_royaltyRecipient common.Address) (*types.Transaction, error) {
	return _INiftyswapExchange20.Contract.SendRoyalties(&_INiftyswapExchange20.TransactOpts, _royaltyRecipient)
}

// INiftyswapExchange20CurrencyPurchaseIterator is returned from FilterCurrencyPurchase and is used to iterate over the raw logs and unpacked data for CurrencyPurchase events raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20CurrencyPurchaseIterator struct {
	Event *INiftyswapExchange20CurrencyPurchase // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchange20CurrencyPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchange20CurrencyPurchase)
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
		it.Event = new(INiftyswapExchange20CurrencyPurchase)
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
func (it *INiftyswapExchange20CurrencyPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchange20CurrencyPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchange20CurrencyPurchase represents a CurrencyPurchase event raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20CurrencyPurchase struct {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) FilterCurrencyPurchase(opts *bind.FilterOpts, buyer []common.Address, recipient []common.Address) (*INiftyswapExchange20CurrencyPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.FilterLogs(opts, "CurrencyPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20CurrencyPurchaseIterator{contract: _INiftyswapExchange20.contract, event: "CurrencyPurchase", logs: logs, sub: sub}, nil
}

// WatchCurrencyPurchase is a free log subscription operation binding the contract event 0x89e4dbdd48f69e7920342e9ad9691b9a7150f254e6a0af177ccfd2556aab8bcd.
//
// Solidity: event CurrencyPurchase(address indexed buyer, address indexed recipient, uint256[] tokensSoldIds, uint256[] tokensSoldAmounts, uint256[] currencyBoughtAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) WatchCurrencyPurchase(opts *bind.WatchOpts, sink chan<- *INiftyswapExchange20CurrencyPurchase, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.WatchLogs(opts, "CurrencyPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchange20CurrencyPurchase)
				if err := _INiftyswapExchange20.contract.UnpackLog(event, "CurrencyPurchase", log); err != nil {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) ParseCurrencyPurchase(log types.Log) (*INiftyswapExchange20CurrencyPurchase, error) {
	event := new(INiftyswapExchange20CurrencyPurchase)
	if err := _INiftyswapExchange20.contract.UnpackLog(event, "CurrencyPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchange20LiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20LiquidityAddedIterator struct {
	Event *INiftyswapExchange20LiquidityAdded // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchange20LiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchange20LiquidityAdded)
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
		it.Event = new(INiftyswapExchange20LiquidityAdded)
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
func (it *INiftyswapExchange20LiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchange20LiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchange20LiquidityAdded represents a LiquidityAdded event raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20LiquidityAdded struct {
	Provider        common.Address
	TokenIds        []*big.Int
	TokenAmounts    []*big.Int
	CurrencyAmounts []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x403f9dc4582dae52d3eeb4a22d37540ffb13c32d964c92ec5ac0d3d5628da316.
//
// Solidity: event LiquidityAdded(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) FilterLiquidityAdded(opts *bind.FilterOpts, provider []common.Address) (*INiftyswapExchange20LiquidityAddedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.FilterLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20LiquidityAddedIterator{contract: _INiftyswapExchange20.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x403f9dc4582dae52d3eeb4a22d37540ffb13c32d964c92ec5ac0d3d5628da316.
//
// Solidity: event LiquidityAdded(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *INiftyswapExchange20LiquidityAdded, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.WatchLogs(opts, "LiquidityAdded", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchange20LiquidityAdded)
				if err := _INiftyswapExchange20.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) ParseLiquidityAdded(log types.Log) (*INiftyswapExchange20LiquidityAdded, error) {
	event := new(INiftyswapExchange20LiquidityAdded)
	if err := _INiftyswapExchange20.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchange20LiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20LiquidityRemovedIterator struct {
	Event *INiftyswapExchange20LiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchange20LiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchange20LiquidityRemoved)
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
		it.Event = new(INiftyswapExchange20LiquidityRemoved)
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
func (it *INiftyswapExchange20LiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchange20LiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchange20LiquidityRemoved represents a LiquidityRemoved event raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20LiquidityRemoved struct {
	Provider        common.Address
	TokenIds        []*big.Int
	TokenAmounts    []*big.Int
	CurrencyAmounts []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x711e9bcb94b4cf7bc99c1cb938edc75ac7e85a136838e90abf6ee1f5adebd423.
//
// Solidity: event LiquidityRemoved(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) FilterLiquidityRemoved(opts *bind.FilterOpts, provider []common.Address) (*INiftyswapExchange20LiquidityRemovedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.FilterLogs(opts, "LiquidityRemoved", providerRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20LiquidityRemovedIterator{contract: _INiftyswapExchange20.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x711e9bcb94b4cf7bc99c1cb938edc75ac7e85a136838e90abf6ee1f5adebd423.
//
// Solidity: event LiquidityRemoved(address indexed provider, uint256[] tokenIds, uint256[] tokenAmounts, uint256[] currencyAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *INiftyswapExchange20LiquidityRemoved, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.WatchLogs(opts, "LiquidityRemoved", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchange20LiquidityRemoved)
				if err := _INiftyswapExchange20.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) ParseLiquidityRemoved(log types.Log) (*INiftyswapExchange20LiquidityRemoved, error) {
	event := new(INiftyswapExchange20LiquidityRemoved)
	if err := _INiftyswapExchange20.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchange20RoyaltyChangedIterator is returned from FilterRoyaltyChanged and is used to iterate over the raw logs and unpacked data for RoyaltyChanged events raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20RoyaltyChangedIterator struct {
	Event *INiftyswapExchange20RoyaltyChanged // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchange20RoyaltyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchange20RoyaltyChanged)
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
		it.Event = new(INiftyswapExchange20RoyaltyChanged)
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
func (it *INiftyswapExchange20RoyaltyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchange20RoyaltyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchange20RoyaltyChanged represents a RoyaltyChanged event raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20RoyaltyChanged struct {
	RoyaltyRecipient common.Address
	RoyaltyFee       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyChanged is a free log retrieval operation binding the contract event 0x02365318429bf1d603e8383b62068288a077545c5c9e709201d563b3f56ce2b3.
//
// Solidity: event RoyaltyChanged(address indexed royaltyRecipient, uint256 royaltyFee)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) FilterRoyaltyChanged(opts *bind.FilterOpts, royaltyRecipient []common.Address) (*INiftyswapExchange20RoyaltyChangedIterator, error) {

	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.FilterLogs(opts, "RoyaltyChanged", royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20RoyaltyChangedIterator{contract: _INiftyswapExchange20.contract, event: "RoyaltyChanged", logs: logs, sub: sub}, nil
}

// WatchRoyaltyChanged is a free log subscription operation binding the contract event 0x02365318429bf1d603e8383b62068288a077545c5c9e709201d563b3f56ce2b3.
//
// Solidity: event RoyaltyChanged(address indexed royaltyRecipient, uint256 royaltyFee)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) WatchRoyaltyChanged(opts *bind.WatchOpts, sink chan<- *INiftyswapExchange20RoyaltyChanged, royaltyRecipient []common.Address) (event.Subscription, error) {

	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.WatchLogs(opts, "RoyaltyChanged", royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchange20RoyaltyChanged)
				if err := _INiftyswapExchange20.contract.UnpackLog(event, "RoyaltyChanged", log); err != nil {
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

// ParseRoyaltyChanged is a log parse operation binding the contract event 0x02365318429bf1d603e8383b62068288a077545c5c9e709201d563b3f56ce2b3.
//
// Solidity: event RoyaltyChanged(address indexed royaltyRecipient, uint256 royaltyFee)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) ParseRoyaltyChanged(log types.Log) (*INiftyswapExchange20RoyaltyChanged, error) {
	event := new(INiftyswapExchange20RoyaltyChanged)
	if err := _INiftyswapExchange20.contract.UnpackLog(event, "RoyaltyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INiftyswapExchange20TokensPurchaseIterator is returned from FilterTokensPurchase and is used to iterate over the raw logs and unpacked data for TokensPurchase events raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20TokensPurchaseIterator struct {
	Event *INiftyswapExchange20TokensPurchase // Event containing the contract specifics and raw log

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
func (it *INiftyswapExchange20TokensPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INiftyswapExchange20TokensPurchase)
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
		it.Event = new(INiftyswapExchange20TokensPurchase)
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
func (it *INiftyswapExchange20TokensPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INiftyswapExchange20TokensPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INiftyswapExchange20TokensPurchase represents a TokensPurchase event raised by the INiftyswapExchange20 contract.
type INiftyswapExchange20TokensPurchase struct {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) FilterTokensPurchase(opts *bind.FilterOpts, buyer []common.Address, recipient []common.Address) (*INiftyswapExchange20TokensPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.FilterLogs(opts, "TokensPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INiftyswapExchange20TokensPurchaseIterator{contract: _INiftyswapExchange20.contract, event: "TokensPurchase", logs: logs, sub: sub}, nil
}

// WatchTokensPurchase is a free log subscription operation binding the contract event 0xd38bc77e62e239476b3e25620d73f29a4a188e808aad79f4a81aaf44871a7a30.
//
// Solidity: event TokensPurchase(address indexed buyer, address indexed recipient, uint256[] tokensBoughtIds, uint256[] tokensBoughtAmounts, uint256[] currencySoldAmounts)
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) WatchTokensPurchase(opts *bind.WatchOpts, sink chan<- *INiftyswapExchange20TokensPurchase, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INiftyswapExchange20.contract.WatchLogs(opts, "TokensPurchase", buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INiftyswapExchange20TokensPurchase)
				if err := _INiftyswapExchange20.contract.UnpackLog(event, "TokensPurchase", log); err != nil {
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
func (_INiftyswapExchange20 *INiftyswapExchange20Filterer) ParseTokensPurchase(log types.Log) (*INiftyswapExchange20TokensPurchase, error) {
	event := new(INiftyswapExchange20TokensPurchase)
	if err := _INiftyswapExchange20.contract.UnpackLog(event, "TokensPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
