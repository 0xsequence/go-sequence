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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061108b8061001f6000396000f3fe608060405234801561001057600080fd5b50600061001e600036610039565b9050600061002b82610457565b905061003782826104d2565b005b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c6001808216810361009d57600060608401526100ae565b84810135606090811c908401526014015b6007600183901c1680156101015760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b6000836010166010036101165750600161013f565b836020166020036101325750600282019186013560f01c61013f565b50600182019186013560f81c5b8067ffffffffffffffff81111561015857610158610c72565b6040519080825280602002602001820160405280156101e357816020015b6101d06040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b8152602001906001900390816101765790505b50604086015260005b8181101561044c5760018085019489013560f81c90808216900361024b57308760400151838151811061022157610221610ca1565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052610295565b8489013560601c601486018860400151848151811061026c5761026c610ca1565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036102d3578489013560208601886040015184815181106102be576102be610ca1565b60200260200101516020018197508281525050505b6004808216900361036b57600385019489013560e81c89868a6102f68483610cd0565b9261030393929190610d0a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061034e5761034e610ca1565b6020908102919091010151604001526103678187610cd0565b9550505b600880821690036103a95784890135602086018860400151848151811061039457610394610ca1565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106103c9576103c9610ca1565b602002602001015160800190151590811515815250508060201660ff16602014876040015183815181106103ff576103ff610ca1565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061043457610434610ca1565b602090810291909101015160c00152506001016101ec565b505050505092915050565b6000806104688360200151306107a5565b9050600061047584610872565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604082015151600090815b8181101561079e576000856040015182815181106104fd576104fd610ca1565b602002602001015190508060a001518015610516575083155b1561055e576040805186815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150610796565b606081015180158015906105715750805a105b156105b75786835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105ae93929190610f6e565b60405180910390fd5b8160800151156105f6576040517f230d1ccc000000000000000000000000000000000000000000000000000000008152600481018490526024016105ae565b8151602083015160009161061c9184156106105784610612565b5a5b8660400151610ade565b9050806107595760c083015161067a57600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d878561065b610af6565b60405161066a93929190610f93565b60405180910390a1505050610796565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016106e45787846106af610af6565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105ae93929190610fb2565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610759577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b878561073a610af6565b60405161074993929190610f93565b60405180910390a150505061079e565b60408051888152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b6001016104dd565b5050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856108155746610818565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0016040516020818303038152906040528051906020012090505b92915050565b60008082610100015160405160200161088b9190610fe7565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120835190915060ff166109345760006108dc8460400151610b15565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016104b3565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016109c45760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610a345760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016109a6565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01610aa45760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082015290810191909152606081018290526080016109a6565b82516040517f0481832000000000000000000000000000000000000000000000000000000000815260ff90911660048201526024016105ae565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6000606060005b8351811015610b97576000610b49858381518110610b3c57610b3c610ca1565b6020026020010151610ba6565b90508281604051602001610b5e929190611033565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101610b1c565b50805160209091012092915050565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001610c5598979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b604051602081830303815290604052805190602001209050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8082018082111561086c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008085851115610d1a57600080fd5b83861115610d2757600080fd5b5050820193919092039150565b60005b83811015610d4f578181015183820152602001610d37565b50506000910152565b60008151808452610d70816020860160208601610d34565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015610e72577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152610e2e60e0860182610d58565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101610dc0565b50909695505050505050565b600081518084526020840193506020830160005b82811015610ec657815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101610e92565b5093949350505050565b805160ff16825260006020820151610eec602085018215159052565b5060408201516101206040850152610f08610120850182610da2565b9050606083015160608501526080830151608085015260a083015184820360a0860152610f358282610d58565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152610f658282610e7e565b95945050505050565b606081526000610f816060830186610ed0565b60208301949094525060400152919050565b838152826020820152606060408201526000610f656060830184610d58565b606081526000610fc56060830186610ed0565b8460208401528281036040840152610fdd8185610d58565b9695505050505050565b8151600090829060208501835b8281101561102857815173ffffffffffffffffffffffffffffffffffffffff16845260209384019390910190600101610ff4565b509195945050505050565b60008351611045818460208801610d34565b919091019182525060200191905056fea26469706673582212204a0b2edf5a9600cbec74da59252163cf3e9be18af0ef0f08bbb33871d4055cb964736f6c634300081b0033",
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

// WalletGuestCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletGuest contract.
type WalletGuestCallAbortedIterator struct {
	Event *WalletGuestCallAborted // Event containing the contract specifics and raw log

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
func (it *WalletGuestCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestCallAborted)
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
		it.Event = new(WalletGuestCallAborted)
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
func (it *WalletGuestCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestCallAborted represents a CallAborted event raised by the WalletGuest contract.
type WalletGuestCallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletGuestCallAbortedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletGuestCallAbortedIterator{contract: _WalletGuest.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletGuestCallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestCallAborted)
				if err := _WalletGuest.contract.UnpackLog(event, "CallAborted", log); err != nil {
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

// ParseCallAborted is a log parse operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) ParseCallAborted(log types.Log) (*WalletGuestCallAborted, error) {
	event := new(WalletGuestCallAborted)
	if err := _WalletGuest.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletGuest contract.
type WalletGuestCallFailedIterator struct {
	Event *WalletGuestCallFailed // Event containing the contract specifics and raw log

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
func (it *WalletGuestCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestCallFailed)
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
		it.Event = new(WalletGuestCallFailed)
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
func (it *WalletGuestCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestCallFailed represents a CallFailed event raised by the WalletGuest contract.
type WalletGuestCallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletGuestCallFailedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletGuestCallFailedIterator{contract: _WalletGuest.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletGuestCallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestCallFailed)
				if err := _WalletGuest.contract.UnpackLog(event, "CallFailed", log); err != nil {
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

// ParseCallFailed is a log parse operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletGuest *WalletGuestFilterer) ParseCallFailed(log types.Log) (*WalletGuestCallFailed, error) {
	event := new(WalletGuestCallFailed)
	if err := _WalletGuest.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletGuest contract.
type WalletGuestCallSkippedIterator struct {
	Event *WalletGuestCallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletGuestCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestCallSkipped)
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
		it.Event = new(WalletGuestCallSkipped)
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
func (it *WalletGuestCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestCallSkipped represents a CallSkipped event raised by the WalletGuest contract.
type WalletGuestCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletGuestCallSkippedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletGuestCallSkippedIterator{contract: _WalletGuest.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletGuestCallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestCallSkipped)
				if err := _WalletGuest.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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

// ParseCallSkipped is a log parse operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseCallSkipped(log types.Log) (*WalletGuestCallSkipped, error) {
	event := new(WalletGuestCallSkipped)
	if err := _WalletGuest.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletGuest contract.
type WalletGuestCallSucceededIterator struct {
	Event *WalletGuestCallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletGuestCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestCallSucceeded)
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
		it.Event = new(WalletGuestCallSucceeded)
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
func (it *WalletGuestCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestCallSucceeded represents a CallSucceeded event raised by the WalletGuest contract.
type WalletGuestCallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterCallSucceeded(opts *bind.FilterOpts) (*WalletGuestCallSucceededIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &WalletGuestCallSucceededIterator{contract: _WalletGuest.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletGuestCallSucceeded) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestCallSucceeded)
				if err := _WalletGuest.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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

// ParseCallSucceeded is a log parse operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseCallSucceeded(log types.Log) (*WalletGuestCallSucceeded, error) {
	event := new(WalletGuestCallSucceeded)
	if err := _WalletGuest.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
