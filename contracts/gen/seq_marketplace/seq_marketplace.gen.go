// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seq_marketplace

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

// ISequenceMarketStorageRequest is an auto generated low-level Go binding around an user-defined struct.
type ISequenceMarketStorageRequest struct {
	Creator       common.Address
	IsListing     bool
	IsERC1155     bool
	TokenContract common.Address
	TokenId       *big.Int
	Quantity      *big.Int
	Expiry        *big.Int
	Currency      common.Address
	PricePerToken *big.Int
}

// ISequenceMarketStorageRequestParams is an auto generated low-level Go binding around an user-defined struct.
type ISequenceMarketStorageRequestParams struct {
	IsListing     bool
	IsERC1155     bool
	TokenContract common.Address
	TokenId       *big.Int
	Quantity      *big.Int
	Expiry        *big.Int
	Currency      common.Address
	PricePerToken *big.Int
}

// SequenceMarketplaceMetaData contains all meta data concerning the SequenceMarketplace contract.
var SequenceMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"additionalFees\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"additionalFeeRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptRequestBatch\",\"inputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"quantities\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"additionalFees\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"additionalFeeRecipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelRequestBatch\",\"inputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRequest\",\"inputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structISequenceMarketStorage.RequestParams\",\"components\":[{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRequestBatch\",\"inputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structISequenceMarketStorage.RequestParams[]\",\"components\":[{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structISequenceMarketStorage.Request\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestBatch\",\"inputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structISequenceMarketStorage.Request[]\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoyaltyInfo\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cost\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"royalty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"invalidateRequests\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"invalidateRequests\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isRequestValid\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"request\",\"type\":\"tuple\",\"internalType\":\"structISequenceMarketStorage.Request\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRequestValidBatch\",\"inputs\":[{\"name\":\"requestIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"quantities\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"requests\",\"type\":\"tuple[]\",\"internalType\":\"structISequenceMarketStorage.Request[]\",\"components\":[{\"name\":\"creator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isListing\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isERC1155\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CustomRoyaltyChanged\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestAccepted\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"quantityRemaining\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestCancelled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestCreated\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isListing\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currency\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"pricePerToken\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestsInvalidated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"invalidatedBefore\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestsInvalidated\",\"inputs\":[{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"invalidatedBefore\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAdditionalFees\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBatchRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurrency\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCurrencyApproval\",\"inputs\":[{\"name\":\"currency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidExpiry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidQuantity\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRequestId\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidRoyalty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenApproval\",\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"quantity\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Invalidated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedContractInterface\",\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]}]",
}

// SequenceMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use SequenceMarketplaceMetaData.ABI instead.
var SequenceMarketplaceABI = SequenceMarketplaceMetaData.ABI

// SequenceMarketplace is an auto generated Go binding around an Ethereum contract.
type SequenceMarketplace struct {
	SequenceMarketplaceCaller     // Read-only binding to the contract
	SequenceMarketplaceTransactor // Write-only binding to the contract
	SequenceMarketplaceFilterer   // Log filterer for contract events
}

// SequenceMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequenceMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequenceMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequenceMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequenceMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequenceMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequenceMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequenceMarketplaceSession struct {
	Contract     *SequenceMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequenceMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequenceMarketplaceCallerSession struct {
	Contract *SequenceMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SequenceMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequenceMarketplaceTransactorSession struct {
	Contract     *SequenceMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SequenceMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequenceMarketplaceRaw struct {
	Contract *SequenceMarketplace // Generic contract binding to access the raw methods on
}

// SequenceMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequenceMarketplaceCallerRaw struct {
	Contract *SequenceMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// SequenceMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequenceMarketplaceTransactorRaw struct {
	Contract *SequenceMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequenceMarketplace creates a new instance of SequenceMarketplace, bound to a specific deployed contract.
func NewSequenceMarketplace(address common.Address, backend bind.ContractBackend) (*SequenceMarketplace, error) {
	contract, err := bindSequenceMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplace{SequenceMarketplaceCaller: SequenceMarketplaceCaller{contract: contract}, SequenceMarketplaceTransactor: SequenceMarketplaceTransactor{contract: contract}, SequenceMarketplaceFilterer: SequenceMarketplaceFilterer{contract: contract}}, nil
}

// NewSequenceMarketplaceCaller creates a new read-only instance of SequenceMarketplace, bound to a specific deployed contract.
func NewSequenceMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*SequenceMarketplaceCaller, error) {
	contract, err := bindSequenceMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceCaller{contract: contract}, nil
}

// NewSequenceMarketplaceTransactor creates a new write-only instance of SequenceMarketplace, bound to a specific deployed contract.
func NewSequenceMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*SequenceMarketplaceTransactor, error) {
	contract, err := bindSequenceMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceTransactor{contract: contract}, nil
}

// NewSequenceMarketplaceFilterer creates a new log filterer instance of SequenceMarketplace, bound to a specific deployed contract.
func NewSequenceMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*SequenceMarketplaceFilterer, error) {
	contract, err := bindSequenceMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceFilterer{contract: contract}, nil
}

// bindSequenceMarketplace binds a generic wrapper to an already deployed contract.
func bindSequenceMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequenceMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequenceMarketplace *SequenceMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequenceMarketplace.Contract.SequenceMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequenceMarketplace *SequenceMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.SequenceMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequenceMarketplace *SequenceMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.SequenceMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequenceMarketplace *SequenceMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequenceMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequenceMarketplace *SequenceMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequenceMarketplace *SequenceMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.contract.Transact(opts, method, params...)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceCaller) GetRequest(opts *bind.CallOpts, requestId *big.Int) (ISequenceMarketStorageRequest, error) {
	var out []interface{}
	err := _SequenceMarketplace.contract.Call(opts, &out, "getRequest", requestId)

	if err != nil {
		return *new(ISequenceMarketStorageRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(ISequenceMarketStorageRequest)).(*ISequenceMarketStorageRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceSession) GetRequest(requestId *big.Int) (ISequenceMarketStorageRequest, error) {
	return _SequenceMarketplace.Contract.GetRequest(&_SequenceMarketplace.CallOpts, requestId)
}

// GetRequest is a free data retrieval call binding the contract method 0xc58343ef.
//
// Solidity: function getRequest(uint256 requestId) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceCallerSession) GetRequest(requestId *big.Int) (ISequenceMarketStorageRequest, error) {
	return _SequenceMarketplace.Contract.GetRequest(&_SequenceMarketplace.CallOpts, requestId)
}

// GetRequestBatch is a free data retrieval call binding the contract method 0x956463e5.
//
// Solidity: function getRequestBatch(uint256[] requestIds) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceCaller) GetRequestBatch(opts *bind.CallOpts, requestIds []*big.Int) ([]ISequenceMarketStorageRequest, error) {
	var out []interface{}
	err := _SequenceMarketplace.contract.Call(opts, &out, "getRequestBatch", requestIds)

	if err != nil {
		return *new([]ISequenceMarketStorageRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISequenceMarketStorageRequest)).(*[]ISequenceMarketStorageRequest)

	return out0, err

}

// GetRequestBatch is a free data retrieval call binding the contract method 0x956463e5.
//
// Solidity: function getRequestBatch(uint256[] requestIds) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceSession) GetRequestBatch(requestIds []*big.Int) ([]ISequenceMarketStorageRequest, error) {
	return _SequenceMarketplace.Contract.GetRequestBatch(&_SequenceMarketplace.CallOpts, requestIds)
}

// GetRequestBatch is a free data retrieval call binding the contract method 0x956463e5.
//
// Solidity: function getRequestBatch(uint256[] requestIds) view returns((address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceCallerSession) GetRequestBatch(requestIds []*big.Int) ([]ISequenceMarketStorageRequest, error) {
	return _SequenceMarketplace.Contract.GetRequestBatch(&_SequenceMarketplace.CallOpts, requestIds)
}

// GetRoyaltyInfo is a free data retrieval call binding the contract method 0x36de9742.
//
// Solidity: function getRoyaltyInfo(address tokenContract, uint256 tokenId, uint256 cost) view returns(address recipient, uint256 royalty)
func (_SequenceMarketplace *SequenceMarketplaceCaller) GetRoyaltyInfo(opts *bind.CallOpts, tokenContract common.Address, tokenId *big.Int, cost *big.Int) (struct {
	Recipient common.Address
	Royalty   *big.Int
}, error) {
	var out []interface{}
	err := _SequenceMarketplace.contract.Call(opts, &out, "getRoyaltyInfo", tokenContract, tokenId, cost)

	outstruct := new(struct {
		Recipient common.Address
		Royalty   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Royalty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoyaltyInfo is a free data retrieval call binding the contract method 0x36de9742.
//
// Solidity: function getRoyaltyInfo(address tokenContract, uint256 tokenId, uint256 cost) view returns(address recipient, uint256 royalty)
func (_SequenceMarketplace *SequenceMarketplaceSession) GetRoyaltyInfo(tokenContract common.Address, tokenId *big.Int, cost *big.Int) (struct {
	Recipient common.Address
	Royalty   *big.Int
}, error) {
	return _SequenceMarketplace.Contract.GetRoyaltyInfo(&_SequenceMarketplace.CallOpts, tokenContract, tokenId, cost)
}

// GetRoyaltyInfo is a free data retrieval call binding the contract method 0x36de9742.
//
// Solidity: function getRoyaltyInfo(address tokenContract, uint256 tokenId, uint256 cost) view returns(address recipient, uint256 royalty)
func (_SequenceMarketplace *SequenceMarketplaceCallerSession) GetRoyaltyInfo(tokenContract common.Address, tokenId *big.Int, cost *big.Int) (struct {
	Recipient common.Address
	Royalty   *big.Int
}, error) {
	return _SequenceMarketplace.Contract.GetRoyaltyInfo(&_SequenceMarketplace.CallOpts, tokenContract, tokenId, cost)
}

// IsRequestValid is a free data retrieval call binding the contract method 0x24dc6bd0.
//
// Solidity: function isRequestValid(uint256 requestId, uint256 quantity) view returns(bool valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceCaller) IsRequestValid(opts *bind.CallOpts, requestId *big.Int, quantity *big.Int) (struct {
	Valid   bool
	Request ISequenceMarketStorageRequest
}, error) {
	var out []interface{}
	err := _SequenceMarketplace.contract.Call(opts, &out, "isRequestValid", requestId, quantity)

	outstruct := new(struct {
		Valid   bool
		Request ISequenceMarketStorageRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Request = *abi.ConvertType(out[1], new(ISequenceMarketStorageRequest)).(*ISequenceMarketStorageRequest)

	return *outstruct, err

}

// IsRequestValid is a free data retrieval call binding the contract method 0x24dc6bd0.
//
// Solidity: function isRequestValid(uint256 requestId, uint256 quantity) view returns(bool valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceSession) IsRequestValid(requestId *big.Int, quantity *big.Int) (struct {
	Valid   bool
	Request ISequenceMarketStorageRequest
}, error) {
	return _SequenceMarketplace.Contract.IsRequestValid(&_SequenceMarketplace.CallOpts, requestId, quantity)
}

// IsRequestValid is a free data retrieval call binding the contract method 0x24dc6bd0.
//
// Solidity: function isRequestValid(uint256 requestId, uint256 quantity) view returns(bool valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256) request)
func (_SequenceMarketplace *SequenceMarketplaceCallerSession) IsRequestValid(requestId *big.Int, quantity *big.Int) (struct {
	Valid   bool
	Request ISequenceMarketStorageRequest
}, error) {
	return _SequenceMarketplace.Contract.IsRequestValid(&_SequenceMarketplace.CallOpts, requestId, quantity)
}

// IsRequestValidBatch is a free data retrieval call binding the contract method 0x449768ad.
//
// Solidity: function isRequestValidBatch(uint256[] requestIds, uint256[] quantities) view returns(bool[] valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceCaller) IsRequestValidBatch(opts *bind.CallOpts, requestIds []*big.Int, quantities []*big.Int) (struct {
	Valid    []bool
	Requests []ISequenceMarketStorageRequest
}, error) {
	var out []interface{}
	err := _SequenceMarketplace.contract.Call(opts, &out, "isRequestValidBatch", requestIds, quantities)

	outstruct := new(struct {
		Valid    []bool
		Requests []ISequenceMarketStorageRequest
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new([]bool)).(*[]bool)
	outstruct.Requests = *abi.ConvertType(out[1], new([]ISequenceMarketStorageRequest)).(*[]ISequenceMarketStorageRequest)

	return *outstruct, err

}

// IsRequestValidBatch is a free data retrieval call binding the contract method 0x449768ad.
//
// Solidity: function isRequestValidBatch(uint256[] requestIds, uint256[] quantities) view returns(bool[] valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceSession) IsRequestValidBatch(requestIds []*big.Int, quantities []*big.Int) (struct {
	Valid    []bool
	Requests []ISequenceMarketStorageRequest
}, error) {
	return _SequenceMarketplace.Contract.IsRequestValidBatch(&_SequenceMarketplace.CallOpts, requestIds, quantities)
}

// IsRequestValidBatch is a free data retrieval call binding the contract method 0x449768ad.
//
// Solidity: function isRequestValidBatch(uint256[] requestIds, uint256[] quantities) view returns(bool[] valid, (address,bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests)
func (_SequenceMarketplace *SequenceMarketplaceCallerSession) IsRequestValidBatch(requestIds []*big.Int, quantities []*big.Int) (struct {
	Valid    []bool
	Requests []ISequenceMarketStorageRequest
}, error) {
	return _SequenceMarketplace.Contract.IsRequestValidBatch(&_SequenceMarketplace.CallOpts, requestIds, quantities)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xe6bd720e.
//
// Solidity: function acceptRequest(uint256 requestId, uint256 quantity, address recipient, uint256[] additionalFees, address[] additionalFeeRecipients) payable returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) AcceptRequest(opts *bind.TransactOpts, requestId *big.Int, quantity *big.Int, recipient common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "acceptRequest", requestId, quantity, recipient, additionalFees, additionalFeeRecipients)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xe6bd720e.
//
// Solidity: function acceptRequest(uint256 requestId, uint256 quantity, address recipient, uint256[] additionalFees, address[] additionalFeeRecipients) payable returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) AcceptRequest(requestId *big.Int, quantity *big.Int, recipient common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.AcceptRequest(&_SequenceMarketplace.TransactOpts, requestId, quantity, recipient, additionalFees, additionalFeeRecipients)
}

// AcceptRequest is a paid mutator transaction binding the contract method 0xe6bd720e.
//
// Solidity: function acceptRequest(uint256 requestId, uint256 quantity, address recipient, uint256[] additionalFees, address[] additionalFeeRecipients) payable returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) AcceptRequest(requestId *big.Int, quantity *big.Int, recipient common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.AcceptRequest(&_SequenceMarketplace.TransactOpts, requestId, quantity, recipient, additionalFees, additionalFeeRecipients)
}

// AcceptRequestBatch is a paid mutator transaction binding the contract method 0xa93c5317.
//
// Solidity: function acceptRequestBatch(uint256[] requestIds, uint256[] quantities, address[] recipients, uint256[] additionalFees, address[] additionalFeeRecipients) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) AcceptRequestBatch(opts *bind.TransactOpts, requestIds []*big.Int, quantities []*big.Int, recipients []common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "acceptRequestBatch", requestIds, quantities, recipients, additionalFees, additionalFeeRecipients)
}

// AcceptRequestBatch is a paid mutator transaction binding the contract method 0xa93c5317.
//
// Solidity: function acceptRequestBatch(uint256[] requestIds, uint256[] quantities, address[] recipients, uint256[] additionalFees, address[] additionalFeeRecipients) returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) AcceptRequestBatch(requestIds []*big.Int, quantities []*big.Int, recipients []common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.AcceptRequestBatch(&_SequenceMarketplace.TransactOpts, requestIds, quantities, recipients, additionalFees, additionalFeeRecipients)
}

// AcceptRequestBatch is a paid mutator transaction binding the contract method 0xa93c5317.
//
// Solidity: function acceptRequestBatch(uint256[] requestIds, uint256[] quantities, address[] recipients, uint256[] additionalFees, address[] additionalFeeRecipients) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) AcceptRequestBatch(requestIds []*big.Int, quantities []*big.Int, recipients []common.Address, additionalFees []*big.Int, additionalFeeRecipients []common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.AcceptRequestBatch(&_SequenceMarketplace.TransactOpts, requestIds, quantities, recipients, additionalFees, additionalFeeRecipients)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 requestId) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) CancelRequest(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "cancelRequest", requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 requestId) returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) CancelRequest(requestId *big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CancelRequest(&_SequenceMarketplace.TransactOpts, requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x3015394c.
//
// Solidity: function cancelRequest(uint256 requestId) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) CancelRequest(requestId *big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CancelRequest(&_SequenceMarketplace.TransactOpts, requestId)
}

// CancelRequestBatch is a paid mutator transaction binding the contract method 0xf1f5176c.
//
// Solidity: function cancelRequestBatch(uint256[] requestIds) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) CancelRequestBatch(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "cancelRequestBatch", requestIds)
}

// CancelRequestBatch is a paid mutator transaction binding the contract method 0xf1f5176c.
//
// Solidity: function cancelRequestBatch(uint256[] requestIds) returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) CancelRequestBatch(requestIds []*big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CancelRequestBatch(&_SequenceMarketplace.TransactOpts, requestIds)
}

// CancelRequestBatch is a paid mutator transaction binding the contract method 0xf1f5176c.
//
// Solidity: function cancelRequestBatch(uint256[] requestIds) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) CancelRequestBatch(requestIds []*big.Int) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CancelRequestBatch(&_SequenceMarketplace.TransactOpts, requestIds)
}

// CreateRequest is a paid mutator transaction binding the contract method 0xdd0de6ec.
//
// Solidity: function createRequest((bool,bool,address,uint256,uint256,uint96,address,uint256) request) returns(uint256 requestId)
func (_SequenceMarketplace *SequenceMarketplaceTransactor) CreateRequest(opts *bind.TransactOpts, request ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "createRequest", request)
}

// CreateRequest is a paid mutator transaction binding the contract method 0xdd0de6ec.
//
// Solidity: function createRequest((bool,bool,address,uint256,uint256,uint96,address,uint256) request) returns(uint256 requestId)
func (_SequenceMarketplace *SequenceMarketplaceSession) CreateRequest(request ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CreateRequest(&_SequenceMarketplace.TransactOpts, request)
}

// CreateRequest is a paid mutator transaction binding the contract method 0xdd0de6ec.
//
// Solidity: function createRequest((bool,bool,address,uint256,uint256,uint96,address,uint256) request) returns(uint256 requestId)
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) CreateRequest(request ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CreateRequest(&_SequenceMarketplace.TransactOpts, request)
}

// CreateRequestBatch is a paid mutator transaction binding the contract method 0xb056d4ae.
//
// Solidity: function createRequestBatch((bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests) returns(uint256[] requestIds)
func (_SequenceMarketplace *SequenceMarketplaceTransactor) CreateRequestBatch(opts *bind.TransactOpts, requests []ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "createRequestBatch", requests)
}

// CreateRequestBatch is a paid mutator transaction binding the contract method 0xb056d4ae.
//
// Solidity: function createRequestBatch((bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests) returns(uint256[] requestIds)
func (_SequenceMarketplace *SequenceMarketplaceSession) CreateRequestBatch(requests []ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CreateRequestBatch(&_SequenceMarketplace.TransactOpts, requests)
}

// CreateRequestBatch is a paid mutator transaction binding the contract method 0xb056d4ae.
//
// Solidity: function createRequestBatch((bool,bool,address,uint256,uint256,uint96,address,uint256)[] requests) returns(uint256[] requestIds)
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) CreateRequestBatch(requests []ISequenceMarketStorageRequestParams) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.CreateRequestBatch(&_SequenceMarketplace.TransactOpts, requests)
}

// InvalidateRequests is a paid mutator transaction binding the contract method 0x3ed97342.
//
// Solidity: function invalidateRequests() returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) InvalidateRequests(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "invalidateRequests")
}

// InvalidateRequests is a paid mutator transaction binding the contract method 0x3ed97342.
//
// Solidity: function invalidateRequests() returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) InvalidateRequests() (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.InvalidateRequests(&_SequenceMarketplace.TransactOpts)
}

// InvalidateRequests is a paid mutator transaction binding the contract method 0x3ed97342.
//
// Solidity: function invalidateRequests() returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) InvalidateRequests() (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.InvalidateRequests(&_SequenceMarketplace.TransactOpts)
}

// InvalidateRequests0 is a paid mutator transaction binding the contract method 0xffaf1d8d.
//
// Solidity: function invalidateRequests(address tokenContract) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactor) InvalidateRequests0(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.contract.Transact(opts, "invalidateRequests0", tokenContract)
}

// InvalidateRequests0 is a paid mutator transaction binding the contract method 0xffaf1d8d.
//
// Solidity: function invalidateRequests(address tokenContract) returns()
func (_SequenceMarketplace *SequenceMarketplaceSession) InvalidateRequests0(tokenContract common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.InvalidateRequests0(&_SequenceMarketplace.TransactOpts, tokenContract)
}

// InvalidateRequests0 is a paid mutator transaction binding the contract method 0xffaf1d8d.
//
// Solidity: function invalidateRequests(address tokenContract) returns()
func (_SequenceMarketplace *SequenceMarketplaceTransactorSession) InvalidateRequests0(tokenContract common.Address) (*types.Transaction, error) {
	return _SequenceMarketplace.Contract.InvalidateRequests0(&_SequenceMarketplace.TransactOpts, tokenContract)
}

// SequenceMarketplaceCustomRoyaltyChangedIterator is returned from FilterCustomRoyaltyChanged and is used to iterate over the raw logs and unpacked data for CustomRoyaltyChanged events raised by the SequenceMarketplace contract.
type SequenceMarketplaceCustomRoyaltyChangedIterator struct {
	Event *SequenceMarketplaceCustomRoyaltyChanged // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceCustomRoyaltyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceCustomRoyaltyChanged)
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
		it.Event = new(SequenceMarketplaceCustomRoyaltyChanged)
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
func (it *SequenceMarketplaceCustomRoyaltyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceCustomRoyaltyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceCustomRoyaltyChanged represents a CustomRoyaltyChanged event raised by the SequenceMarketplace contract.
type SequenceMarketplaceCustomRoyaltyChanged struct {
	TokenContract common.Address
	Recipient     common.Address
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCustomRoyaltyChanged is a free log retrieval operation binding the contract event 0x60567f9d30ab22ef3cd7557f56b897c677c80a85c8673a4a5c26eb9349ef8c60.
//
// Solidity: event CustomRoyaltyChanged(address indexed tokenContract, address recipient, uint96 fee)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterCustomRoyaltyChanged(opts *bind.FilterOpts, tokenContract []common.Address) (*SequenceMarketplaceCustomRoyaltyChangedIterator, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "CustomRoyaltyChanged", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceCustomRoyaltyChangedIterator{contract: _SequenceMarketplace.contract, event: "CustomRoyaltyChanged", logs: logs, sub: sub}, nil
}

// WatchCustomRoyaltyChanged is a free log subscription operation binding the contract event 0x60567f9d30ab22ef3cd7557f56b897c677c80a85c8673a4a5c26eb9349ef8c60.
//
// Solidity: event CustomRoyaltyChanged(address indexed tokenContract, address recipient, uint96 fee)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchCustomRoyaltyChanged(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceCustomRoyaltyChanged, tokenContract []common.Address) (event.Subscription, error) {

	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "CustomRoyaltyChanged", tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceCustomRoyaltyChanged)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "CustomRoyaltyChanged", log); err != nil {
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

// ParseCustomRoyaltyChanged is a log parse operation binding the contract event 0x60567f9d30ab22ef3cd7557f56b897c677c80a85c8673a4a5c26eb9349ef8c60.
//
// Solidity: event CustomRoyaltyChanged(address indexed tokenContract, address recipient, uint96 fee)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseCustomRoyaltyChanged(log types.Log) (*SequenceMarketplaceCustomRoyaltyChanged, error) {
	event := new(SequenceMarketplaceCustomRoyaltyChanged)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "CustomRoyaltyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequenceMarketplaceRequestAcceptedIterator is returned from FilterRequestAccepted and is used to iterate over the raw logs and unpacked data for RequestAccepted events raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestAcceptedIterator struct {
	Event *SequenceMarketplaceRequestAccepted // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceRequestAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceRequestAccepted)
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
		it.Event = new(SequenceMarketplaceRequestAccepted)
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
func (it *SequenceMarketplaceRequestAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceRequestAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceRequestAccepted represents a RequestAccepted event raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestAccepted struct {
	RequestId         *big.Int
	Buyer             common.Address
	TokenContract     common.Address
	Recipient         common.Address
	Quantity          *big.Int
	QuantityRemaining *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRequestAccepted is a free log retrieval operation binding the contract event 0x146c7c3a67244cf406ca437a82b2bf587bc02b628b63659c73b4295b8b00e76f.
//
// Solidity: event RequestAccepted(uint256 indexed requestId, address indexed buyer, address indexed tokenContract, address recipient, uint256 quantity, uint256 quantityRemaining)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterRequestAccepted(opts *bind.FilterOpts, requestId []*big.Int, buyer []common.Address, tokenContract []common.Address) (*SequenceMarketplaceRequestAcceptedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "RequestAccepted", requestIdRule, buyerRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceRequestAcceptedIterator{contract: _SequenceMarketplace.contract, event: "RequestAccepted", logs: logs, sub: sub}, nil
}

// WatchRequestAccepted is a free log subscription operation binding the contract event 0x146c7c3a67244cf406ca437a82b2bf587bc02b628b63659c73b4295b8b00e76f.
//
// Solidity: event RequestAccepted(uint256 indexed requestId, address indexed buyer, address indexed tokenContract, address recipient, uint256 quantity, uint256 quantityRemaining)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchRequestAccepted(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceRequestAccepted, requestId []*big.Int, buyer []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "RequestAccepted", requestIdRule, buyerRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceRequestAccepted)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
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

// ParseRequestAccepted is a log parse operation binding the contract event 0x146c7c3a67244cf406ca437a82b2bf587bc02b628b63659c73b4295b8b00e76f.
//
// Solidity: event RequestAccepted(uint256 indexed requestId, address indexed buyer, address indexed tokenContract, address recipient, uint256 quantity, uint256 quantityRemaining)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseRequestAccepted(log types.Log) (*SequenceMarketplaceRequestAccepted, error) {
	event := new(SequenceMarketplaceRequestAccepted)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequenceMarketplaceRequestCancelledIterator is returned from FilterRequestCancelled and is used to iterate over the raw logs and unpacked data for RequestCancelled events raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestCancelledIterator struct {
	Event *SequenceMarketplaceRequestCancelled // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceRequestCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceRequestCancelled)
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
		it.Event = new(SequenceMarketplaceRequestCancelled)
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
func (it *SequenceMarketplaceRequestCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceRequestCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceRequestCancelled represents a RequestCancelled event raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestCancelled struct {
	RequestId     *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCancelled is a free log retrieval operation binding the contract event 0xe0d7665e06e7db1fc500d66d4e3898d1d4a5533d7efe54b352fcdaa177c22783.
//
// Solidity: event RequestCancelled(uint256 indexed requestId, address indexed tokenContract)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterRequestCancelled(opts *bind.FilterOpts, requestId []*big.Int, tokenContract []common.Address) (*SequenceMarketplaceRequestCancelledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "RequestCancelled", requestIdRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceRequestCancelledIterator{contract: _SequenceMarketplace.contract, event: "RequestCancelled", logs: logs, sub: sub}, nil
}

// WatchRequestCancelled is a free log subscription operation binding the contract event 0xe0d7665e06e7db1fc500d66d4e3898d1d4a5533d7efe54b352fcdaa177c22783.
//
// Solidity: event RequestCancelled(uint256 indexed requestId, address indexed tokenContract)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchRequestCancelled(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceRequestCancelled, requestId []*big.Int, tokenContract []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "RequestCancelled", requestIdRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceRequestCancelled)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestCancelled", log); err != nil {
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

// ParseRequestCancelled is a log parse operation binding the contract event 0xe0d7665e06e7db1fc500d66d4e3898d1d4a5533d7efe54b352fcdaa177c22783.
//
// Solidity: event RequestCancelled(uint256 indexed requestId, address indexed tokenContract)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseRequestCancelled(log types.Log) (*SequenceMarketplaceRequestCancelled, error) {
	event := new(SequenceMarketplaceRequestCancelled)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequenceMarketplaceRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestCreatedIterator struct {
	Event *SequenceMarketplaceRequestCreated // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceRequestCreated)
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
		it.Event = new(SequenceMarketplaceRequestCreated)
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
func (it *SequenceMarketplaceRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceRequestCreated represents a RequestCreated event raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestCreated struct {
	RequestId     *big.Int
	Creator       common.Address
	TokenContract common.Address
	TokenId       *big.Int
	IsListing     bool
	Quantity      *big.Int
	Currency      common.Address
	PricePerToken *big.Int
	Expiry        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0xb67ee55b059861d68b3b5640ee51bde6a5d9d849c6ecdb610663cbd4290bdfd5.
//
// Solidity: event RequestCreated(uint256 indexed requestId, address indexed creator, address indexed tokenContract, uint256 tokenId, bool isListing, uint256 quantity, address currency, uint256 pricePerToken, uint256 expiry)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterRequestCreated(opts *bind.FilterOpts, requestId []*big.Int, creator []common.Address, tokenContract []common.Address) (*SequenceMarketplaceRequestCreatedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "RequestCreated", requestIdRule, creatorRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceRequestCreatedIterator{contract: _SequenceMarketplace.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0xb67ee55b059861d68b3b5640ee51bde6a5d9d849c6ecdb610663cbd4290bdfd5.
//
// Solidity: event RequestCreated(uint256 indexed requestId, address indexed creator, address indexed tokenContract, uint256 tokenId, bool isListing, uint256 quantity, address currency, uint256 pricePerToken, uint256 expiry)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceRequestCreated, requestId []*big.Int, creator []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "RequestCreated", requestIdRule, creatorRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceRequestCreated)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0xb67ee55b059861d68b3b5640ee51bde6a5d9d849c6ecdb610663cbd4290bdfd5.
//
// Solidity: event RequestCreated(uint256 indexed requestId, address indexed creator, address indexed tokenContract, uint256 tokenId, bool isListing, uint256 quantity, address currency, uint256 pricePerToken, uint256 expiry)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseRequestCreated(log types.Log) (*SequenceMarketplaceRequestCreated, error) {
	event := new(SequenceMarketplaceRequestCreated)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequenceMarketplaceRequestsInvalidatedIterator is returned from FilterRequestsInvalidated and is used to iterate over the raw logs and unpacked data for RequestsInvalidated events raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestsInvalidatedIterator struct {
	Event *SequenceMarketplaceRequestsInvalidated // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceRequestsInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceRequestsInvalidated)
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
		it.Event = new(SequenceMarketplaceRequestsInvalidated)
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
func (it *SequenceMarketplaceRequestsInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceRequestsInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceRequestsInvalidated represents a RequestsInvalidated event raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestsInvalidated struct {
	Creator           common.Address
	InvalidatedBefore *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRequestsInvalidated is a free log retrieval operation binding the contract event 0xb9fd665eb33107da2a1f0113e84d01c5915024b696fa667e66f80505bf3b2977.
//
// Solidity: event RequestsInvalidated(address indexed creator, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterRequestsInvalidated(opts *bind.FilterOpts, creator []common.Address, invalidatedBefore []*big.Int) (*SequenceMarketplaceRequestsInvalidatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var invalidatedBeforeRule []interface{}
	for _, invalidatedBeforeItem := range invalidatedBefore {
		invalidatedBeforeRule = append(invalidatedBeforeRule, invalidatedBeforeItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "RequestsInvalidated", creatorRule, invalidatedBeforeRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceRequestsInvalidatedIterator{contract: _SequenceMarketplace.contract, event: "RequestsInvalidated", logs: logs, sub: sub}, nil
}

// WatchRequestsInvalidated is a free log subscription operation binding the contract event 0xb9fd665eb33107da2a1f0113e84d01c5915024b696fa667e66f80505bf3b2977.
//
// Solidity: event RequestsInvalidated(address indexed creator, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchRequestsInvalidated(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceRequestsInvalidated, creator []common.Address, invalidatedBefore []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var invalidatedBeforeRule []interface{}
	for _, invalidatedBeforeItem := range invalidatedBefore {
		invalidatedBeforeRule = append(invalidatedBeforeRule, invalidatedBeforeItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "RequestsInvalidated", creatorRule, invalidatedBeforeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceRequestsInvalidated)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestsInvalidated", log); err != nil {
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

// ParseRequestsInvalidated is a log parse operation binding the contract event 0xb9fd665eb33107da2a1f0113e84d01c5915024b696fa667e66f80505bf3b2977.
//
// Solidity: event RequestsInvalidated(address indexed creator, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseRequestsInvalidated(log types.Log) (*SequenceMarketplaceRequestsInvalidated, error) {
	event := new(SequenceMarketplaceRequestsInvalidated)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestsInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequenceMarketplaceRequestsInvalidated0Iterator is returned from FilterRequestsInvalidated0 and is used to iterate over the raw logs and unpacked data for RequestsInvalidated0 events raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestsInvalidated0Iterator struct {
	Event *SequenceMarketplaceRequestsInvalidated0 // Event containing the contract specifics and raw log

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
func (it *SequenceMarketplaceRequestsInvalidated0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequenceMarketplaceRequestsInvalidated0)
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
		it.Event = new(SequenceMarketplaceRequestsInvalidated0)
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
func (it *SequenceMarketplaceRequestsInvalidated0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequenceMarketplaceRequestsInvalidated0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequenceMarketplaceRequestsInvalidated0 represents a RequestsInvalidated0 event raised by the SequenceMarketplace contract.
type SequenceMarketplaceRequestsInvalidated0 struct {
	Creator           common.Address
	TokenContract     common.Address
	InvalidatedBefore *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRequestsInvalidated0 is a free log retrieval operation binding the contract event 0x359d2c7e56147963a26a08fd48199b16e28d5fff314390b6c79bb822362340ff.
//
// Solidity: event RequestsInvalidated(address indexed creator, address indexed tokenContract, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) FilterRequestsInvalidated0(opts *bind.FilterOpts, creator []common.Address, tokenContract []common.Address, invalidatedBefore []*big.Int) (*SequenceMarketplaceRequestsInvalidated0Iterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var invalidatedBeforeRule []interface{}
	for _, invalidatedBeforeItem := range invalidatedBefore {
		invalidatedBeforeRule = append(invalidatedBeforeRule, invalidatedBeforeItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.FilterLogs(opts, "RequestsInvalidated0", creatorRule, tokenContractRule, invalidatedBeforeRule)
	if err != nil {
		return nil, err
	}
	return &SequenceMarketplaceRequestsInvalidated0Iterator{contract: _SequenceMarketplace.contract, event: "RequestsInvalidated0", logs: logs, sub: sub}, nil
}

// WatchRequestsInvalidated0 is a free log subscription operation binding the contract event 0x359d2c7e56147963a26a08fd48199b16e28d5fff314390b6c79bb822362340ff.
//
// Solidity: event RequestsInvalidated(address indexed creator, address indexed tokenContract, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) WatchRequestsInvalidated0(opts *bind.WatchOpts, sink chan<- *SequenceMarketplaceRequestsInvalidated0, creator []common.Address, tokenContract []common.Address, invalidatedBefore []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var invalidatedBeforeRule []interface{}
	for _, invalidatedBeforeItem := range invalidatedBefore {
		invalidatedBeforeRule = append(invalidatedBeforeRule, invalidatedBeforeItem)
	}

	logs, sub, err := _SequenceMarketplace.contract.WatchLogs(opts, "RequestsInvalidated0", creatorRule, tokenContractRule, invalidatedBeforeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequenceMarketplaceRequestsInvalidated0)
				if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestsInvalidated0", log); err != nil {
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

// ParseRequestsInvalidated0 is a log parse operation binding the contract event 0x359d2c7e56147963a26a08fd48199b16e28d5fff314390b6c79bb822362340ff.
//
// Solidity: event RequestsInvalidated(address indexed creator, address indexed tokenContract, uint256 indexed invalidatedBefore)
func (_SequenceMarketplace *SequenceMarketplaceFilterer) ParseRequestsInvalidated0(log types.Log) (*SequenceMarketplaceRequestsInvalidated0, error) {
	event := new(SequenceMarketplaceRequestsInvalidated0)
	if err := _SequenceMarketplace.contract.UnpackLog(event, "RequestsInvalidated0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
