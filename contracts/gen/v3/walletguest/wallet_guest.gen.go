// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletguest

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

// WalletGuestMetaData contains all meta data concerning the WalletGuest contract.
var WalletGuestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Failed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Skipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Success\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506110fc8061001f6000396000f3fe608060405234801561001057600080fd5b503660008036600081811061002757610027610d16565b909101357fff0000000000000000000000000000000000000000000000000000000000000016600003905061006d57610064366001816000610d45565b91509150610076565b50606490506024355b6000610082838361009d565b9050600061008f826104bb565b905061009b8282610536565b005b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c600180821681036101015760006060840152610112565b84810135606090811c908401526014015b6007600183901c1680156101655760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b60008360101660100361017a575060016101a3565b836020166020036101965750600282019186013560f01c6101a3565b50600182019186013560f81c5b8067ffffffffffffffff8111156101bc576101bc610d6f565b60405190808252806020026020018201604052801561024757816020015b6102346040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b8152602001906001900390816101da5790505b50604086015260005b818110156104b05760018085019489013560f81c9080821690036102af57308760400151838151811061028557610285610d16565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526102f9565b8489013560601c60148601886040015184815181106102d0576102d0610d16565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036103375784890135602086018860400151848151811061032257610322610d16565b60200260200101516020018197508281525050505b600480821690036103cf57600385019489013560e81c89868a61035a8483610d9e565b9261036793929190610d45565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106103b2576103b2610d16565b6020908102919091010151604001526103cb8187610d9e565b9550505b6008808216900361040d578489013560208601886040015184815181106103f8576103f8610d16565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061042d5761042d610d16565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061046357610463610d16565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061049857610498610d16565b602090810291909101015160c0015250600101610250565b505050505092915050565b6000806104cc8360200151306107eb565b905060006104d9846108b8565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604082015151600090815b818110156107e45760008560400151828151811061056157610561610d16565b602002602001015190508060a00151801561057a575083155b156105c2576040805186815260208101849052600095507fa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7910160405180910390a1506107dc565b606081015180158015906105d55750805a105b1561061b5786835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161061293929190611012565b60405180910390fd5b81608001511561065a576040517f230d1ccc00000000000000000000000000000000000000000000000000000000815260048101849052602401610612565b600061067483600001518460200151848660400151610b0e565b90508061079f5760c08301516106c9576040805188815260208101869052600197507fe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643910160405180910390a15050506107dc565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016107335787846106fe610b26565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161061293929190611037565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161079f5760408051888152602081018690527f6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8910160405180910390a15050506107e4565b60408051888152602081018690527f2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3910160405180910390a15050505b600101610541565b5050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de8561085b574661085e565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0016040516020818303038152906040528051906020012090505b92915050565b6000806108c9836101000151610b45565b835190915060ff1661093c5760006108e48460400151610bc8565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001610517565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016109cc5760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610a3c5760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016109ae565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01610aac5760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e2624602082015290810191909152606081018290526080016109ae565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e64000000000000000000000000000000006044820152606401610612565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6000606060005b8351811015610bb95781848281518110610b6857610b68610d16565b6020026020010151604051602001610b8192919061106c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101610b4c565b50805160209091012092915050565b6000606060005b8351811015610bb9576000610bfc858381518110610bef57610bef610d16565b6020026020010151610c4a565b90508281604051602001610c119291906110a4565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101610bcf565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001610cf998979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b604051602081830303815290604052805190602001209050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008085851115610d5557600080fd5b83861115610d6257600080fd5b5050820193919092039150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b808201808211156108b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60005b83811015610df3578181015183820152602001610ddb565b50506000910152565b60008151808452610e14816020860160208601610dd8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015610f16577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152610ed260e0860182610dfc565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101610e64565b50909695505050505050565b600081518084526020840193506020830160005b82811015610f6a57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101610f36565b5093949350505050565b805160ff16825260006020820151610f90602085018215159052565b5060408201516101206040850152610fac610120850182610e46565b9050606083015160608501526080830151608085015260a083015184820360a0860152610fd98282610dfc565b91505060c083015160c085015260e083015160e08501526101008301518482036101008601526110098282610f22565b95945050505050565b6060815260006110256060830186610f74565b60208301949094525060400152919050565b60608152600061104a6060830186610f74565b84602084015282810360408401526110628185610dfc565b9695505050505050565b60408152600061107f6040830185610dfc565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600083516110b6818460208801610dd8565b919091019182525060200191905056fea2646970667358221220216f5accfdf8d241c3de7b985c06aadc7bc5391e677e3e9e36b40467fd06da7764736f6c634300081b0033",
}

// WalletGuestABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletGuestMetaData.ABI instead.
var WalletGuestABI = WalletGuestMetaData.ABI

// WalletGuestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletGuestMetaData.Bin instead.
var WalletGuestBin = WalletGuestMetaData.Bin

// DeployWalletGuest deploys a new Ethereum contract, binding an instance of WalletGuest to it.
func DeployWalletGuest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletGuest, error) {
	parsed, err := WalletGuestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletGuestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletGuest{WalletGuestCaller: WalletGuestCaller{contract: contract}, WalletGuestTransactor: WalletGuestTransactor{contract: contract}, WalletGuestFilterer: WalletGuestFilterer{contract: contract}}, nil
}

// WalletGuest is an auto generated Go binding around an Ethereum contract.
type WalletGuest struct {
	WalletGuestCaller     // Read-only binding to the contract
	WalletGuestTransactor // Write-only binding to the contract
	WalletGuestFilterer   // Log filterer for contract events
}

// WalletGuestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletGuestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletGuestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletGuestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletGuestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletGuestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletGuestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletGuestSession struct {
	Contract     *WalletGuest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletGuestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletGuestCallerSession struct {
	Contract *WalletGuestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WalletGuestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletGuestTransactorSession struct {
	Contract     *WalletGuestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletGuestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletGuestRaw struct {
	Contract *WalletGuest // Generic contract binding to access the raw methods on
}

// WalletGuestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletGuestCallerRaw struct {
	Contract *WalletGuestCaller // Generic read-only contract binding to access the raw methods on
}

// WalletGuestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletGuestTransactorRaw struct {
	Contract *WalletGuestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletGuest creates a new instance of WalletGuest, bound to a specific deployed contract.
func NewWalletGuest(address common.Address, backend bind.ContractBackend) (*WalletGuest, error) {
	contract, err := bindWalletGuest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletGuest{WalletGuestCaller: WalletGuestCaller{contract: contract}, WalletGuestTransactor: WalletGuestTransactor{contract: contract}, WalletGuestFilterer: WalletGuestFilterer{contract: contract}}, nil
}

// NewWalletGuestCaller creates a new read-only instance of WalletGuest, bound to a specific deployed contract.
func NewWalletGuestCaller(address common.Address, caller bind.ContractCaller) (*WalletGuestCaller, error) {
	contract, err := bindWalletGuest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletGuestCaller{contract: contract}, nil
}

// NewWalletGuestTransactor creates a new write-only instance of WalletGuest, bound to a specific deployed contract.
func NewWalletGuestTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletGuestTransactor, error) {
	contract, err := bindWalletGuest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletGuestTransactor{contract: contract}, nil
}

// NewWalletGuestFilterer creates a new log filterer instance of WalletGuest, bound to a specific deployed contract.
func NewWalletGuestFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletGuestFilterer, error) {
	contract, err := bindWalletGuest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletGuestFilterer{contract: contract}, nil
}

// bindWalletGuest binds a generic wrapper to an already deployed contract.
func bindWalletGuest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletGuestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletGuest *WalletGuestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletGuest.Contract.WalletGuestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletGuest *WalletGuestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletGuest.Contract.WalletGuestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletGuest *WalletGuestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletGuest.Contract.WalletGuestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletGuest *WalletGuestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletGuest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletGuest *WalletGuestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletGuest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletGuest *WalletGuestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletGuest.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_WalletGuest *WalletGuestTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_WalletGuest *WalletGuestSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Fallback(&_WalletGuest.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_WalletGuest *WalletGuestTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Fallback(&_WalletGuest.TransactOpts, calldata)
}

// WalletGuestAbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the WalletGuest contract.
type WalletGuestAbortedIterator struct {
	Event *WalletGuestAborted // Event containing the contract specifics and raw log

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
func (it *WalletGuestAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestAborted)
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
		it.Event = new(WalletGuestAborted)
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
func (it *WalletGuestAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestAborted represents a Aborted event raised by the WalletGuest contract.
type WalletGuestAborted struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterAborted(opts *bind.FilterOpts) (*WalletGuestAbortedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &WalletGuestAbortedIterator{contract: _WalletGuest.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *WalletGuestAborted) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestAborted)
				if err := _WalletGuest.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// ParseAborted is a log parse operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseAborted(log types.Log) (*WalletGuestAborted, error) {
	event := new(WalletGuestAborted)
	if err := _WalletGuest.contract.UnpackLog(event, "Aborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestFailedIterator is returned from FilterFailed and is used to iterate over the raw logs and unpacked data for Failed events raised by the WalletGuest contract.
type WalletGuestFailedIterator struct {
	Event *WalletGuestFailed // Event containing the contract specifics and raw log

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
func (it *WalletGuestFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestFailed)
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
		it.Event = new(WalletGuestFailed)
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
func (it *WalletGuestFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestFailed represents a Failed event raised by the WalletGuest contract.
type WalletGuestFailed struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailed is a free log retrieval operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterFailed(opts *bind.FilterOpts) (*WalletGuestFailedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return &WalletGuestFailedIterator{contract: _WalletGuest.contract, event: "Failed", logs: logs, sub: sub}, nil
}

// WatchFailed is a free log subscription operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchFailed(opts *bind.WatchOpts, sink chan<- *WalletGuestFailed) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestFailed)
				if err := _WalletGuest.contract.UnpackLog(event, "Failed", log); err != nil {
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

// ParseFailed is a log parse operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseFailed(log types.Log) (*WalletGuestFailed, error) {
	event := new(WalletGuestFailed)
	if err := _WalletGuest.contract.UnpackLog(event, "Failed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestSkippedIterator is returned from FilterSkipped and is used to iterate over the raw logs and unpacked data for Skipped events raised by the WalletGuest contract.
type WalletGuestSkippedIterator struct {
	Event *WalletGuestSkipped // Event containing the contract specifics and raw log

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
func (it *WalletGuestSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestSkipped)
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
		it.Event = new(WalletGuestSkipped)
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
func (it *WalletGuestSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestSkipped represents a Skipped event raised by the WalletGuest contract.
type WalletGuestSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkipped is a free log retrieval operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterSkipped(opts *bind.FilterOpts) (*WalletGuestSkippedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return &WalletGuestSkippedIterator{contract: _WalletGuest.contract, event: "Skipped", logs: logs, sub: sub}, nil
}

// WatchSkipped is a free log subscription operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchSkipped(opts *bind.WatchOpts, sink chan<- *WalletGuestSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestSkipped)
				if err := _WalletGuest.contract.UnpackLog(event, "Skipped", log); err != nil {
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

// ParseSkipped is a log parse operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseSkipped(log types.Log) (*WalletGuestSkipped, error) {
	event := new(WalletGuestSkipped)
	if err := _WalletGuest.contract.UnpackLog(event, "Skipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the WalletGuest contract.
type WalletGuestSuccessIterator struct {
	Event *WalletGuestSuccess // Event containing the contract specifics and raw log

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
func (it *WalletGuestSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestSuccess)
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
		it.Event = new(WalletGuestSuccess)
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
func (it *WalletGuestSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestSuccess represents a Success event raised by the WalletGuest contract.
type WalletGuestSuccess struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterSuccess(opts *bind.FilterOpts) (*WalletGuestSuccessIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return &WalletGuestSuccessIterator{contract: _WalletGuest.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *WalletGuestSuccess) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestSuccess)
				if err := _WalletGuest.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseSuccess(log types.Log) (*WalletGuestSuccess, error) {
	event := new(WalletGuestSuccess)
	if err := _WalletGuest.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
