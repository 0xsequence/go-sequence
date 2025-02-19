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
	Bin: "0x608060405234801561001057600080fd5b50610d6b806100206000396000f3fe6080604052600436106101295760003560e01c8063b7a72531116100a5578063d1db390711610074578063e90f13e711610059578063e90f13e7146102b4578063f209883a146102e6578063ffd7d741146102f957600080fd5b8063d1db3907146102b4578063d5b5337f146102c757600080fd5b8063b7a7253114610222578063c272d5c314610255578063c39f2d5c14610268578063c66764e11461028757600080fd5b80637f29d538116100fc57806398f9fbc4116100e157806398f9fbc41461020f578063aeea5fb514610222578063b472f0a21461023557600080fd5b80637f29d538146101b9578063984395bc146101db57600080fd5b80630fdecfac1461012e57806343d9c9351461015057806348acd29f14610165578063543196eb1461019a575b600080fd5b34801561013a57600080fd5b50465b6040519081526020015b60405180910390f35b34801561015c57600080fd5b5061013d61031a565b34801561017157600080fd5b5061013d610180366004610842565b73ffffffffffffffffffffffffffffffffffffffff163190565b3480156101a657600080fd5b5061013d6101b5366004610842565b3f90565b3480156101c557600080fd5b506101d96101d4366004610864565b610322565b005b3480156101e757600080fd5b50325b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610147565b34801561021b57600080fd5b50416101ea565b34801561022e57600080fd5b504461013d565b34801561024157600080fd5b506101d961025036600461087d565b6103b9565b34801561026157600080fd5b503a61013d565b34801561027457600080fd5b5061013d610283366004610842565b3b90565b34801561029357600080fd5b506102a76102a2366004610842565b610500565b6040516101479190610915565b3480156102c057600080fd5b504561013d565b3480156102d357600080fd5b5061013d6102e2366004610864565b4090565b3480156102f257600080fd5b504261013d565b61030c6103073660046109df565b610545565b604051610147929190610b99565b60005a905090565b8042106103b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f526571756972655574696c7323726571756972654e6f6e457870697265643a2060448201527f455850495245440000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b50565b600080606083901c6bffffffffffffffffffffffff84166040517f8c3f556300000000000000000000000000000000000000000000000000000000815260048101839052919350915060009073ffffffffffffffffffffffffffffffffffffffff861690638c3f556390602401602060405180830381865afa158015610443573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104679190610c51565b9050818110156104f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f526571756972655574696c7323726571756972654d696e4e6f6e63653a204e4f60448201527f4e43455f42454c4f575f5245515549524544000000000000000000000000000060648201526084016103ad565b5050505050565b60408051603f833b9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092528181529080600060208401853c50919050565b606080825167ffffffffffffffff81111561056257610562610928565b60405190808252806020026020018201604052801561058b578160200160208202803683370190505b509150825167ffffffffffffffff8111156105a8576105a8610928565b6040519080825280602002602001820160405280156105db57816020015b60608152602001906001900390816105c65790505b50905060005b83518110156108135760008482815181106105fe576105fe610c6a565b60200260200101519050806000015115610647576040517f230d1ccc000000000000000000000000000000000000000000000000000000008152600481018390526024016103ad565b80604001515a101561069d578181604001515a6040517f2bb3e3ba0000000000000000000000000000000000000000000000000000000081526004810193909352602483019190915260448201526064016103ad565b806060015173ffffffffffffffffffffffffffffffffffffffff16816080015182604001516000146106d35782604001516106d5565b5a5b908360a001516040516106e89190610c99565b600060405180830381858888f193505050503d8060008114610726576040519150601f19603f3d011682016040523d82523d6000602084013e61072b565b606091505b5085848151811061073e5761073e610c6a565b6020026020010185858151811061075757610757610c6a565b602002602001018290528215151515815250505083828151811061077d5761077d610c6a565b60200260200101511580156107ac575084828151811061079f5761079f610c6a565b6020026020010151602001515b1561080057818383815181106107c4576107c4610c6a565b60200260200101516040517f3b4c7a5f0000000000000000000000000000000000000000000000000000000081526004016103ad929190610cb5565b508061080b81610cd6565b9150506105e1565b50915091565b803573ffffffffffffffffffffffffffffffffffffffff8116811461083d57600080fd5b919050565b60006020828403121561085457600080fd5b61085d82610819565b9392505050565b60006020828403121561087657600080fd5b5035919050565b6000806040838503121561089057600080fd5b61089983610819565b946020939093013593505050565b60005b838110156108c25781810151838201526020016108aa565b50506000910152565b600081518084526108e38160208601602086016108a7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061085d60208301846108cb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff8111828210171561097a5761097a610928565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109c7576109c7610928565b604052919050565b8035801515811461083d57600080fd5b600060208083850312156109f257600080fd5b823567ffffffffffffffff80821115610a0a57600080fd5b818501915085601f830112610a1e57600080fd5b813581811115610a3057610a30610928565b8060051b610a3f858201610980565b9182528381018501918581019089841115610a5957600080fd5b86860192505b83831015610b8c57823585811115610a775760008081fd5b860160c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0828d038101821315610aae5760008081fd5b610ab6610957565b610ac18b85016109cf565b81526040610ad08186016109cf565b8c830152606080860135828401526080610aeb818801610819565b8285015260a091508187013581850152508486013594508a851115610b105760008081fd5b84860195508f603f870112610b2757600094508485fd5b8c86013594508a851115610b3d57610b3d610928565b610b4d8d85601f88011601610980565b93508484528f82868801011115610b645760008081fd5b848287018e86013760009484018d019490945250918201528352509186019190860190610a5f565b9998505050505050505050565b604080825283519082018190526000906020906060840190828701845b82811015610bd4578151151584529284019290840190600101610bb6565b50505083810382850152845180825282820190600581901b8301840187850160005b83811015610c42577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0868403018552610c308383516108cb565b94870194925090860190600101610bf6565b50909998505050505050505050565b600060208284031215610c6357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008251610cab8184602087016108a7565b9190910192915050565b828152604060208201526000610cce60408301846108cb565b949350505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d2e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea2646970667358221220b1427d933eb383c444aa6f4e6d4ee70e4f65b9a227f016eecfe265a983f9c42064736f6c63430008120033",
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
