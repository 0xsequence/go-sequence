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
	ABI: "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DelegateCallNotAllowed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
	Bin: "0x60808060405234601557610f6d908161001b8239f35b600080fdfe608060405261000c610330565b60008082523560f81c6001908160018083160361028a5750600060608401525b600761003b60ff831660011c90565b168061023d575b5060108181160361020c57506001905b61005b826103a3565b90604084019182526000905b838210610083576100818561007b81610540565b906108e0565b005b803560f81c906001019091816001808516036101ec57506100c5306100a9838751610441565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808416146101ce575b600480841614610183575b60088084161461014e575b61013761013160c0856101116010806001989916146080610108888c51610441565b51019015159052565b61012760208083161460a0610108888c51610441565b1660061c60031690565b60ff1690565b60c0610144838751610441565b5101520190610067565b61013761013160c06101666001959060208235920190565b96906060610175878b51610441565b5101529594505050506100e6565b906101c890803560e81c906003016101c16101a86101a18484610484565b83366104c0565b919060406101b7888b51610441565b51019236916104db565b9052610484565b906100db565b90803590602001919060206101e4838751610441565b5101526100d0565b6102079250803560601c9060140192906100a9838751610441565b6100c5565b60209081160361022c5761ffff90803560f01c906002015b911690610052565b60ff90803560f81c90600101610224565b61027d919291908060031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b0190843590610100031c16920190565b9190608084015238610042565b8035606090811c90850152601401915061002c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176102ea57604052565b61029f565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ea57604052565b60405190610120820182811067ffffffffffffffff8211176102ea576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b67ffffffffffffffff81116102ea5760051b60200190565b906103ad8261038b565b6103ba60405191826102ef565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06103e8829461038b565b019060005b8281106103f957505050565b602090604051610408816102ce565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c0820152828285010152016103ed565b80518210156104555760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9190820180921161049157565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291928382116104d65783116104d65780920390565b600080fd5b92919267ffffffffffffffff82116102ea576040519161052360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601846102ef565b8294818452818301116104d6578281602093846000960137010152565b6020810151156106645761063261065e6105f060005b60405160208101917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f83527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408301527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606083015260808201523060a082015260a081526105e760c0826102ef565b51902093610bc6565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826102ef565b51902090565b61063261065e6105f046610556565b919082519283825260005b8481106106bd5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161067e565b9080602083519182815201916020808360051b8301019401926000915b8383106106fe57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080610771604085015160e0604086015260e0850190610673565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906106ef565b906020808351928381520192019060005b8181106107c65750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016107b9565b805160ff1682526108769160208281015115159082015261010061085061082a604085015161012060408601526101208501906106d2565b606085015160608501526080850151608085015260a085015184820360a0860152610673565b9260c081015160c084015260e081015160e08401520151906101008184039101526107a8565b90565b610891604092959493956060835260608301906107f2565b9460208201520152565b6108769392606092825260208201528160408201520190610673565b6108cd61087694926060835260608301906107f2565b9260208201526040818403910152610673565b90600091604081019283515193815b8581106108ff575b505050505050565b61090a818351610441565b519261091960a0850151151590565b80610b7a575b610b3a575060009260608101518015801580610b31575b610af9576080830151610acc5761098a9161098691610969855173ffffffffffffffffffffffffffffffffffffffff1690565b91602086015191600014610ac657505a905b604086015192610da1565b1590565b6109cc575b506040805186815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016108ef565b60c001805115610a7d576001815114610a3a57516002146109ed573861098f565b93925050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250610a2c610a20610db3565b6040519384938461089b565b0390a13880808080806108f7565b5083610a79610a47610db3565b6040519384937f7f6b0bb1000000000000000000000000000000000000000000000000000000008552600485016108b7565b0390fd5b509150600180927f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d610abe610ab0610db3565b604051918291858b8461089b565b0390a16109c6565b9061097b565b7f230d1ccc0000000000000000000000000000000000000000000000000000000086526004849052602486fd5b8387610a795a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501610879565b50815a10610936565b6040805187815260208101849052919450600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101610abe565b50801561091f565b805160209091019060005b818110610b9a5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101610b8d565b610100810151604051610be181610632602082018095610b82565b51902090610bf0815160ff1690565b60ff811680610c695750509061065e610c0c6040840151610dfc565b9261063260806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103610cc757505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4669381019384529081019190915260608101929092529061065e8160808101610632565b60028103610d1d57505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e46020820190815291810192909252606082019290925261065e8160808101610632565b600303610d71575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4666020820190815291810192909252606082019290925261065e8160808101610632565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b9160009391849360208451940192f190565b3d90604051916020818401016040528083526000602084013e565b805160209091019060005b818110610de65750505090565b8251845260209384019390920191600101610dd9565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610e42610e2c8361038b565b92610e3a60405194856102ef565b80845261038b565b0136602083013760005b8351811015610f1e5780610e6260019286610441565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152610f0a610120826102ef565b519020610f178285610441565b5201610e4c565b5090915060405161065e81610632602082018095610dce56fea2646970667358221220baf48a03c446c575132553e62013c6320b5d90a1ba275593a31ee1589046bd1464736f6c634300081c0033",
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
// Solidity: fallback() payable returns()
func (_WalletGuest *WalletGuestTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletGuest *WalletGuestSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Fallback(&_WalletGuest.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
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
