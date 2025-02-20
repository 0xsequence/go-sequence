// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletutils

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

// IModuleCallsTransaction is an auto generated low-level Go binding around an user-defined struct.
type IModuleCallsTransaction struct {
	DelegateCall  bool
	RevertOnError bool
	GasLimit      *big.Int
	Target        common.Address
	Value         *big.Int
	Data          []byte
}

// WalletUtilsMetaData contains all meta data concerning the WalletUtils contract.
var WalletUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CallReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_available\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"callBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"}],\"name\":\"callBlockhash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"callCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"code\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"callCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"callCodeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callGasLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callOrigin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callPrevrandao\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"}],\"name\":\"multiCall\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_successes\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"requireMinNonce\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"}],\"name\":\"requireNonExpired\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d7e806100206000396000f3fe6080604052600436106101295760003560e01c8063b7a72531116100a5578063d1db390711610074578063e90f13e711610059578063e90f13e7146102e6578063f209883a146102f9578063ffd7d7411461030c57600080fd5b8063d1db3907146102b4578063d5b5337f146102c757600080fd5b8063b7a7253114610222578063c272d5c314610255578063c39f2d5c14610268578063c66764e11461028757600080fd5b80637f29d538116100fc57806398f9fbc4116100e157806398f9fbc41461020f578063aeea5fb514610222578063b472f0a21461023557600080fd5b80637f29d538146101b9578063984395bc146101db57600080fd5b80630fdecfac1461012e57806343d9c9351461015057806348acd29f14610165578063543196eb1461019a575b600080fd5b34801561013a57600080fd5b50465b6040519081526020015b60405180910390f35b34801561015c57600080fd5b5061013d61032d565b34801561017157600080fd5b5061013d610180366004610855565b73ffffffffffffffffffffffffffffffffffffffff163190565b3480156101a657600080fd5b5061013d6101b5366004610855565b3f90565b3480156101c557600080fd5b506101d96101d4366004610877565b610335565b005b3480156101e757600080fd5b50325b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610147565b34801561021b57600080fd5b50416101ea565b34801561022e57600080fd5b504461013d565b34801561024157600080fd5b506101d9610250366004610890565b6103cc565b34801561026157600080fd5b503a61013d565b34801561027457600080fd5b5061013d610283366004610855565b3b90565b34801561029357600080fd5b506102a76102a2366004610855565b610513565b6040516101479190610928565b3480156102c057600080fd5b504361013d565b3480156102d357600080fd5b5061013d6102e2366004610877565b4090565b3480156102f257600080fd5b504561013d565b34801561030557600080fd5b504261013d565b61031f61031a3660046109f2565b610558565b604051610147929190610bac565b60005a905090565b8042106103c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f526571756972655574696c7323726571756972654e6f6e457870697265643a2060448201527f455850495245440000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b50565b600080606083901c6bffffffffffffffffffffffff84166040517f8c3f556300000000000000000000000000000000000000000000000000000000815260048101839052919350915060009073ffffffffffffffffffffffffffffffffffffffff861690638c3f556390602401602060405180830381865afa158015610456573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047a9190610c64565b90508181101561050c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f526571756972655574696c7323726571756972654d696e4e6f6e63653a204e4f60448201527f4e43455f42454c4f575f5245515549524544000000000000000000000000000060648201526084016103c0565b5050505050565b60408051603f833b9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092528181529080600060208401853c50919050565b606080825167ffffffffffffffff8111156105755761057561093b565b60405190808252806020026020018201604052801561059e578160200160208202803683370190505b509150825167ffffffffffffffff8111156105bb576105bb61093b565b6040519080825280602002602001820160405280156105ee57816020015b60608152602001906001900390816105d95790505b50905060005b835181101561082657600084828151811061061157610611610c7d565b6020026020010151905080600001511561065a576040517f230d1ccc000000000000000000000000000000000000000000000000000000008152600481018390526024016103c0565b80604001515a10156106b0578181604001515a6040517f2bb3e3ba0000000000000000000000000000000000000000000000000000000081526004810193909352602483019190915260448201526064016103c0565b806060015173ffffffffffffffffffffffffffffffffffffffff16816080015182604001516000146106e65782604001516106e8565b5a5b908360a001516040516106fb9190610cac565b600060405180830381858888f193505050503d8060008114610739576040519150601f19603f3d011682016040523d82523d6000602084013e61073e565b606091505b5085848151811061075157610751610c7d565b6020026020010185858151811061076a5761076a610c7d565b602002602001018290528215151515815250505083828151811061079057610790610c7d565b60200260200101511580156107bf57508482815181106107b2576107b2610c7d565b6020026020010151602001515b1561081357818383815181106107d7576107d7610c7d565b60200260200101516040517f3b4c7a5f0000000000000000000000000000000000000000000000000000000081526004016103c0929190610cc8565b508061081e81610ce9565b9150506105f4565b50915091565b803573ffffffffffffffffffffffffffffffffffffffff8116811461085057600080fd5b919050565b60006020828403121561086757600080fd5b6108708261082c565b9392505050565b60006020828403121561088957600080fd5b5035919050565b600080604083850312156108a357600080fd5b6108ac8361082c565b946020939093013593505050565b60005b838110156108d55781810151838201526020016108bd565b50506000910152565b600081518084526108f68160208601602086016108ba565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061087060208301846108de565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff8111828210171561098d5761098d61093b565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109da576109da61093b565b604052919050565b8035801515811461085057600080fd5b60006020808385031215610a0557600080fd5b823567ffffffffffffffff80821115610a1d57600080fd5b818501915085601f830112610a3157600080fd5b813581811115610a4357610a4361093b565b8060051b610a52858201610993565b9182528381018501918581019089841115610a6c57600080fd5b86860192505b83831015610b9f57823585811115610a8a5760008081fd5b860160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d038101821315610ac15760008081fd5b610ac961096a565b610ad48b85016109e2565b81526040610ae38186016109e2565b8c830152606080860135828401526080610afe81880161082c565b8285015260a091508187013581850152508486013594508a851115610b235760008081fd5b84860195508f603f870112610b3a57600094508485fd5b8c86013594508a851115610b5057610b5061093b565b610b608d85601f88011601610993565b93508484528f82868801011115610b775760008081fd5b848287018e86013760009484018d019490945250918201528352509186019190860190610a72565b9998505050505050505050565b604080825283519082018190526000906020906060840190828701845b82811015610be7578151151584529284019290840190600101610bc9565b50505083810382850152845180825282820190600581901b8301840187850160005b83811015610c55577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018552610c438383516108de565b94870194925090860190600101610c09565b50909998505050505050505050565b600060208284031215610c7657600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251610cbe8184602087016108ba565b9190910192915050565b828152604060208201526000610ce160408301846108de565b949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d41577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea264697066735822122049ef042991a3ec1b25eb7d0bd2d2faaceec6986d39c10052cc3d76bbe0bb4ad664736f6c63430008120033",
}

// WalletUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletUtilsMetaData.ABI instead.
var WalletUtilsABI = WalletUtilsMetaData.ABI

// WalletUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletUtilsMetaData.Bin instead.
var WalletUtilsBin = WalletUtilsMetaData.Bin

// DeployWalletUtils deploys a new Ethereum contract, binding an instance of WalletUtils to it.
func DeployWalletUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletUtils, error) {
	parsed, err := WalletUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletUtils{WalletUtilsCaller: WalletUtilsCaller{contract: contract}, WalletUtilsTransactor: WalletUtilsTransactor{contract: contract}, WalletUtilsFilterer: WalletUtilsFilterer{contract: contract}}, nil
}

// WalletUtils is an auto generated Go binding around an Ethereum contract.
type WalletUtils struct {
	WalletUtilsCaller     // Read-only binding to the contract
	WalletUtilsTransactor // Write-only binding to the contract
	WalletUtilsFilterer   // Log filterer for contract events
}

// WalletUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletUtilsSession struct {
	Contract     *WalletUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletUtilsCallerSession struct {
	Contract *WalletUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WalletUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletUtilsTransactorSession struct {
	Contract     *WalletUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletUtilsRaw struct {
	Contract *WalletUtils // Generic contract binding to access the raw methods on
}

// WalletUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletUtilsCallerRaw struct {
	Contract *WalletUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// WalletUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletUtilsTransactorRaw struct {
	Contract *WalletUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletUtils creates a new instance of WalletUtils, bound to a specific deployed contract.
func NewWalletUtils(address common.Address, backend bind.ContractBackend) (*WalletUtils, error) {
	contract, err := bindWalletUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletUtils{WalletUtilsCaller: WalletUtilsCaller{contract: contract}, WalletUtilsTransactor: WalletUtilsTransactor{contract: contract}, WalletUtilsFilterer: WalletUtilsFilterer{contract: contract}}, nil
}

// NewWalletUtilsCaller creates a new read-only instance of WalletUtils, bound to a specific deployed contract.
func NewWalletUtilsCaller(address common.Address, caller bind.ContractCaller) (*WalletUtilsCaller, error) {
	contract, err := bindWalletUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletUtilsCaller{contract: contract}, nil
}

// NewWalletUtilsTransactor creates a new write-only instance of WalletUtils, bound to a specific deployed contract.
func NewWalletUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletUtilsTransactor, error) {
	contract, err := bindWalletUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletUtilsTransactor{contract: contract}, nil
}

// NewWalletUtilsFilterer creates a new log filterer instance of WalletUtils, bound to a specific deployed contract.
func NewWalletUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletUtilsFilterer, error) {
	contract, err := bindWalletUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletUtilsFilterer{contract: contract}, nil
}

// bindWalletUtils binds a generic wrapper to an already deployed contract.
func bindWalletUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletUtils *WalletUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletUtils.Contract.WalletUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletUtils *WalletUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletUtils.Contract.WalletUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletUtils *WalletUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletUtils.Contract.WalletUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletUtils *WalletUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletUtils *WalletUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletUtils *WalletUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletUtils.Contract.contract.Transact(opts, method, params...)
}

// CallBalanceOf is a free data retrieval call binding the contract method 0x48acd29f.
//
// Solidity: function callBalanceOf(address _addr) view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallBalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callBalanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallBalanceOf is a free data retrieval call binding the contract method 0x48acd29f.
//
// Solidity: function callBalanceOf(address _addr) view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallBalanceOf(_addr common.Address) (*big.Int, error) {
	return _WalletUtils.Contract.CallBalanceOf(&_WalletUtils.CallOpts, _addr)
}

// CallBalanceOf is a free data retrieval call binding the contract method 0x48acd29f.
//
// Solidity: function callBalanceOf(address _addr) view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallBalanceOf(_addr common.Address) (*big.Int, error) {
	return _WalletUtils.Contract.CallBalanceOf(&_WalletUtils.CallOpts, _addr)
}

// CallBlockNumber is a free data retrieval call binding the contract method 0xd1db3907.
//
// Solidity: function callBlockNumber() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallBlockNumber is a free data retrieval call binding the contract method 0xd1db3907.
//
// Solidity: function callBlockNumber() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallBlockNumber() (*big.Int, error) {
	return _WalletUtils.Contract.CallBlockNumber(&_WalletUtils.CallOpts)
}

// CallBlockNumber is a free data retrieval call binding the contract method 0xd1db3907.
//
// Solidity: function callBlockNumber() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallBlockNumber() (*big.Int, error) {
	return _WalletUtils.Contract.CallBlockNumber(&_WalletUtils.CallOpts)
}

// CallBlockhash is a free data retrieval call binding the contract method 0xd5b5337f.
//
// Solidity: function callBlockhash(uint256 _i) view returns(bytes32)
func (_WalletUtils *WalletUtilsCaller) CallBlockhash(opts *bind.CallOpts, _i *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callBlockhash", _i)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CallBlockhash is a free data retrieval call binding the contract method 0xd5b5337f.
//
// Solidity: function callBlockhash(uint256 _i) view returns(bytes32)
func (_WalletUtils *WalletUtilsSession) CallBlockhash(_i *big.Int) ([32]byte, error) {
	return _WalletUtils.Contract.CallBlockhash(&_WalletUtils.CallOpts, _i)
}

// CallBlockhash is a free data retrieval call binding the contract method 0xd5b5337f.
//
// Solidity: function callBlockhash(uint256 _i) view returns(bytes32)
func (_WalletUtils *WalletUtilsCallerSession) CallBlockhash(_i *big.Int) ([32]byte, error) {
	return _WalletUtils.Contract.CallBlockhash(&_WalletUtils.CallOpts, _i)
}

// CallChainId is a free data retrieval call binding the contract method 0x0fdecfac.
//
// Solidity: function callChainId() view returns(uint256 id)
func (_WalletUtils *WalletUtilsCaller) CallChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallChainId is a free data retrieval call binding the contract method 0x0fdecfac.
//
// Solidity: function callChainId() view returns(uint256 id)
func (_WalletUtils *WalletUtilsSession) CallChainId() (*big.Int, error) {
	return _WalletUtils.Contract.CallChainId(&_WalletUtils.CallOpts)
}

// CallChainId is a free data retrieval call binding the contract method 0x0fdecfac.
//
// Solidity: function callChainId() view returns(uint256 id)
func (_WalletUtils *WalletUtilsCallerSession) CallChainId() (*big.Int, error) {
	return _WalletUtils.Contract.CallChainId(&_WalletUtils.CallOpts)
}

// CallCode is a free data retrieval call binding the contract method 0xc66764e1.
//
// Solidity: function callCode(address _addr) view returns(bytes code)
func (_WalletUtils *WalletUtilsCaller) CallCode(opts *bind.CallOpts, _addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callCode", _addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CallCode is a free data retrieval call binding the contract method 0xc66764e1.
//
// Solidity: function callCode(address _addr) view returns(bytes code)
func (_WalletUtils *WalletUtilsSession) CallCode(_addr common.Address) ([]byte, error) {
	return _WalletUtils.Contract.CallCode(&_WalletUtils.CallOpts, _addr)
}

// CallCode is a free data retrieval call binding the contract method 0xc66764e1.
//
// Solidity: function callCode(address _addr) view returns(bytes code)
func (_WalletUtils *WalletUtilsCallerSession) CallCode(_addr common.Address) ([]byte, error) {
	return _WalletUtils.Contract.CallCode(&_WalletUtils.CallOpts, _addr)
}

// CallCodeHash is a free data retrieval call binding the contract method 0x543196eb.
//
// Solidity: function callCodeHash(address _addr) view returns(bytes32 codeHash)
func (_WalletUtils *WalletUtilsCaller) CallCodeHash(opts *bind.CallOpts, _addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callCodeHash", _addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CallCodeHash is a free data retrieval call binding the contract method 0x543196eb.
//
// Solidity: function callCodeHash(address _addr) view returns(bytes32 codeHash)
func (_WalletUtils *WalletUtilsSession) CallCodeHash(_addr common.Address) ([32]byte, error) {
	return _WalletUtils.Contract.CallCodeHash(&_WalletUtils.CallOpts, _addr)
}

// CallCodeHash is a free data retrieval call binding the contract method 0x543196eb.
//
// Solidity: function callCodeHash(address _addr) view returns(bytes32 codeHash)
func (_WalletUtils *WalletUtilsCallerSession) CallCodeHash(_addr common.Address) ([32]byte, error) {
	return _WalletUtils.Contract.CallCodeHash(&_WalletUtils.CallOpts, _addr)
}

// CallCodeSize is a free data retrieval call binding the contract method 0xc39f2d5c.
//
// Solidity: function callCodeSize(address _addr) view returns(uint256 size)
func (_WalletUtils *WalletUtilsCaller) CallCodeSize(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callCodeSize", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallCodeSize is a free data retrieval call binding the contract method 0xc39f2d5c.
//
// Solidity: function callCodeSize(address _addr) view returns(uint256 size)
func (_WalletUtils *WalletUtilsSession) CallCodeSize(_addr common.Address) (*big.Int, error) {
	return _WalletUtils.Contract.CallCodeSize(&_WalletUtils.CallOpts, _addr)
}

// CallCodeSize is a free data retrieval call binding the contract method 0xc39f2d5c.
//
// Solidity: function callCodeSize(address _addr) view returns(uint256 size)
func (_WalletUtils *WalletUtilsCallerSession) CallCodeSize(_addr common.Address) (*big.Int, error) {
	return _WalletUtils.Contract.CallCodeSize(&_WalletUtils.CallOpts, _addr)
}

// CallCoinbase is a free data retrieval call binding the contract method 0x98f9fbc4.
//
// Solidity: function callCoinbase() view returns(address)
func (_WalletUtils *WalletUtilsCaller) CallCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallCoinbase is a free data retrieval call binding the contract method 0x98f9fbc4.
//
// Solidity: function callCoinbase() view returns(address)
func (_WalletUtils *WalletUtilsSession) CallCoinbase() (common.Address, error) {
	return _WalletUtils.Contract.CallCoinbase(&_WalletUtils.CallOpts)
}

// CallCoinbase is a free data retrieval call binding the contract method 0x98f9fbc4.
//
// Solidity: function callCoinbase() view returns(address)
func (_WalletUtils *WalletUtilsCallerSession) CallCoinbase() (common.Address, error) {
	return _WalletUtils.Contract.CallCoinbase(&_WalletUtils.CallOpts)
}

// CallDifficulty is a free data retrieval call binding the contract method 0xaeea5fb5.
//
// Solidity: function callDifficulty() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallDifficulty is a free data retrieval call binding the contract method 0xaeea5fb5.
//
// Solidity: function callDifficulty() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallDifficulty() (*big.Int, error) {
	return _WalletUtils.Contract.CallDifficulty(&_WalletUtils.CallOpts)
}

// CallDifficulty is a free data retrieval call binding the contract method 0xaeea5fb5.
//
// Solidity: function callDifficulty() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallDifficulty() (*big.Int, error) {
	return _WalletUtils.Contract.CallDifficulty(&_WalletUtils.CallOpts)
}

// CallGasLeft is a free data retrieval call binding the contract method 0x43d9c935.
//
// Solidity: function callGasLeft() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallGasLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callGasLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallGasLeft is a free data retrieval call binding the contract method 0x43d9c935.
//
// Solidity: function callGasLeft() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallGasLeft() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasLeft(&_WalletUtils.CallOpts)
}

// CallGasLeft is a free data retrieval call binding the contract method 0x43d9c935.
//
// Solidity: function callGasLeft() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallGasLeft() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasLeft(&_WalletUtils.CallOpts)
}

// CallGasLimit is a free data retrieval call binding the contract method 0xe90f13e7.
//
// Solidity: function callGasLimit() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallGasLimit is a free data retrieval call binding the contract method 0xe90f13e7.
//
// Solidity: function callGasLimit() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallGasLimit() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasLimit(&_WalletUtils.CallOpts)
}

// CallGasLimit is a free data retrieval call binding the contract method 0xe90f13e7.
//
// Solidity: function callGasLimit() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallGasLimit() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasLimit(&_WalletUtils.CallOpts)
}

// CallGasPrice is a free data retrieval call binding the contract method 0xc272d5c3.
//
// Solidity: function callGasPrice() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallGasPrice is a free data retrieval call binding the contract method 0xc272d5c3.
//
// Solidity: function callGasPrice() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallGasPrice() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasPrice(&_WalletUtils.CallOpts)
}

// CallGasPrice is a free data retrieval call binding the contract method 0xc272d5c3.
//
// Solidity: function callGasPrice() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallGasPrice() (*big.Int, error) {
	return _WalletUtils.Contract.CallGasPrice(&_WalletUtils.CallOpts)
}

// CallOrigin is a free data retrieval call binding the contract method 0x984395bc.
//
// Solidity: function callOrigin() view returns(address)
func (_WalletUtils *WalletUtilsCaller) CallOrigin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callOrigin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallOrigin is a free data retrieval call binding the contract method 0x984395bc.
//
// Solidity: function callOrigin() view returns(address)
func (_WalletUtils *WalletUtilsSession) CallOrigin() (common.Address, error) {
	return _WalletUtils.Contract.CallOrigin(&_WalletUtils.CallOpts)
}

// CallOrigin is a free data retrieval call binding the contract method 0x984395bc.
//
// Solidity: function callOrigin() view returns(address)
func (_WalletUtils *WalletUtilsCallerSession) CallOrigin() (common.Address, error) {
	return _WalletUtils.Contract.CallOrigin(&_WalletUtils.CallOpts)
}

// CallPrevrandao is a free data retrieval call binding the contract method 0xb7a72531.
//
// Solidity: function callPrevrandao() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallPrevrandao(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callPrevrandao")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallPrevrandao is a free data retrieval call binding the contract method 0xb7a72531.
//
// Solidity: function callPrevrandao() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallPrevrandao() (*big.Int, error) {
	return _WalletUtils.Contract.CallPrevrandao(&_WalletUtils.CallOpts)
}

// CallPrevrandao is a free data retrieval call binding the contract method 0xb7a72531.
//
// Solidity: function callPrevrandao() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallPrevrandao() (*big.Int, error) {
	return _WalletUtils.Contract.CallPrevrandao(&_WalletUtils.CallOpts)
}

// CallTimestamp is a free data retrieval call binding the contract method 0xf209883a.
//
// Solidity: function callTimestamp() view returns(uint256)
func (_WalletUtils *WalletUtilsCaller) CallTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "callTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallTimestamp is a free data retrieval call binding the contract method 0xf209883a.
//
// Solidity: function callTimestamp() view returns(uint256)
func (_WalletUtils *WalletUtilsSession) CallTimestamp() (*big.Int, error) {
	return _WalletUtils.Contract.CallTimestamp(&_WalletUtils.CallOpts)
}

// CallTimestamp is a free data retrieval call binding the contract method 0xf209883a.
//
// Solidity: function callTimestamp() view returns(uint256)
func (_WalletUtils *WalletUtilsCallerSession) CallTimestamp() (*big.Int, error) {
	return _WalletUtils.Contract.CallTimestamp(&_WalletUtils.CallOpts)
}

// RequireMinNonce is a free data retrieval call binding the contract method 0xb472f0a2.
//
// Solidity: function requireMinNonce(address _wallet, uint256 _nonce) view returns()
func (_WalletUtils *WalletUtilsCaller) RequireMinNonce(opts *bind.CallOpts, _wallet common.Address, _nonce *big.Int) error {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "requireMinNonce", _wallet, _nonce)

	if err != nil {
		return err
	}

	return err

}

// RequireMinNonce is a free data retrieval call binding the contract method 0xb472f0a2.
//
// Solidity: function requireMinNonce(address _wallet, uint256 _nonce) view returns()
func (_WalletUtils *WalletUtilsSession) RequireMinNonce(_wallet common.Address, _nonce *big.Int) error {
	return _WalletUtils.Contract.RequireMinNonce(&_WalletUtils.CallOpts, _wallet, _nonce)
}

// RequireMinNonce is a free data retrieval call binding the contract method 0xb472f0a2.
//
// Solidity: function requireMinNonce(address _wallet, uint256 _nonce) view returns()
func (_WalletUtils *WalletUtilsCallerSession) RequireMinNonce(_wallet common.Address, _nonce *big.Int) error {
	return _WalletUtils.Contract.RequireMinNonce(&_WalletUtils.CallOpts, _wallet, _nonce)
}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 _expiration) view returns()
func (_WalletUtils *WalletUtilsCaller) RequireNonExpired(opts *bind.CallOpts, _expiration *big.Int) error {
	var out []interface{}
	err := _WalletUtils.contract.Call(opts, &out, "requireNonExpired", _expiration)

	if err != nil {
		return err
	}

	return err

}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 _expiration) view returns()
func (_WalletUtils *WalletUtilsSession) RequireNonExpired(_expiration *big.Int) error {
	return _WalletUtils.Contract.RequireNonExpired(&_WalletUtils.CallOpts, _expiration)
}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 _expiration) view returns()
func (_WalletUtils *WalletUtilsCallerSession) RequireNonExpired(_expiration *big.Int) error {
	return _WalletUtils.Contract.RequireNonExpired(&_WalletUtils.CallOpts, _expiration)
}

// MultiCall is a paid mutator transaction binding the contract method 0xffd7d741.
//
// Solidity: function multiCall((bool,bool,uint256,address,uint256,bytes)[] _txs) payable returns(bool[] _successes, bytes[] _results)
func (_WalletUtils *WalletUtilsTransactor) MultiCall(opts *bind.TransactOpts, _txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletUtils.contract.Transact(opts, "multiCall", _txs)
}

// MultiCall is a paid mutator transaction binding the contract method 0xffd7d741.
//
// Solidity: function multiCall((bool,bool,uint256,address,uint256,bytes)[] _txs) payable returns(bool[] _successes, bytes[] _results)
func (_WalletUtils *WalletUtilsSession) MultiCall(_txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletUtils.Contract.MultiCall(&_WalletUtils.TransactOpts, _txs)
}

// MultiCall is a paid mutator transaction binding the contract method 0xffd7d741.
//
// Solidity: function multiCall((bool,bool,uint256,address,uint256,bytes)[] _txs) payable returns(bool[] _successes, bytes[] _results)
func (_WalletUtils *WalletUtilsTransactorSession) MultiCall(_txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletUtils.Contract.MultiCall(&_WalletUtils.TransactOpts, _txs)
}
