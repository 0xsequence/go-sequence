// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletsimulator

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

// PayloadCall is an auto generated low-level Go binding around an user-defined struct.
type PayloadCall struct {
	To              common.Address
	Value           *big.Int
	Data            []byte
	GasLimit        *big.Int
	DelegateCall    bool
	OnlyFallback    bool
	BehaviorOnError *big.Int
}

// SimulatorResult is an auto generated low-level Go binding around an user-defined struct.
type SimulatorResult struct {
	Status  uint8
	Result  []byte
	GasUsed *big.Int
}

// WalletSimulatorMetaData contains all meta data concerning the WalletSimulator contract.
var WalletSimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"simulate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSimulator.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"internalType\":\"structSimulator.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610ac38061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a27c392214610030575b600080fd5b61004361003e3660046105e5565b610059565b60405161005091906106ef565b60405180910390f35b606060005a90506000838067ffffffffffffffff81111561007c5761007c6107cd565b6040519080825280602002602001820160405280156100d257816020015b6100bf6040805160608101909152806000815260200160608152602001600081525090565b81526020019060019003908161009a5790505b50935060005b8181101561058d5760008787838181106100f4576100f46107fc565b9050602002810190610106919061082b565b61010f90610978565b905083156101205760009350610130565b8060a00151156101305750610585565b606081015180158015906101435750805a105b156101dd57600587848151811061015c5761015c6107fc565b60200260200101516000019060058111156101795761017961065c565b9081600581111561018c5761018c61065c565b9052505a6040516020016101a291815260200190565b6040516020818303038152906040528784815181106101c3576101c36107fc565b602002602001015160200181905250505050505050610592565b60008260800151156102e35760005a84519091506102ae9084156102015784610203565b5a5b634c4e814c60e01b60008c8a8c60008c6040015160405160240161022c96959493929190610a10565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610598565b91505a6102bb9082610a53565b8986815181106102cd576102cd6107fc565b602002602001015160400181815250505061033d565b60005a8451602086015191925061030c9185156103005785610302565b5a5b87604001516105ae565b91505a6103199082610a53565b89868151811061032b5761032b6107fc565b60200260200101516040018181525050505b806105115760c08301516103c757600195506002888581518110610363576103636107fc565b60200260200101516000019060058111156103805761038061065c565b908160058111156103935761039361065c565b90525061039e6105c6565b8885815181106103b0576103b06107fc565b602002602001015160200181905250505050610585565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161046e576004888581518110610406576104066107fc565b60200260200101516000019060058111156104235761042361065c565b908160058111156104365761043661065c565b9052506104416105c6565b888581518110610453576104536107fc565b60200260200101516020018190525050505050505050610592565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016105115760038885815181106104ad576104ad6107fc565b60200260200101516000019060058111156104ca576104ca61065c565b908160058111156104dd576104dd61065c565b9052506104e86105c6565b8885815181106104fa576104fa6107fc565b60200260200101516020018190525050505061058d565b6001888581518110610525576105256107fc565b60200260200101516000019060058111156105425761054261065c565b908160058111156105555761055561065c565b9052506105606105c6565b888581518110610572576105726107fc565b6020026020010151602001819052505050505b6001016100d8565b505050505b92915050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b600080602083850312156105f857600080fd5b823567ffffffffffffffff81111561060f57600080fd5b8301601f8101851361062057600080fd5b803567ffffffffffffffff81111561063757600080fd5b8560208260051b840101111561064c57600080fd5b6020919091019590945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000815180845260005b818110156106b157602081850181015186830182015201610695565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156107c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110610781577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261079e606088018261068b565b604092830151979092019690965294506020938401939190910190600101610717565b50929695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261085f57600080fd5b9190910192915050565b60405160e0810167ffffffffffffffff8111828210171561088c5761088c6107cd565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff811681146108b657600080fd5b919050565b600082601f8301126108cc57600080fd5b813567ffffffffffffffff8111156108e6576108e66107cd565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff81118282101715610933576109336107cd565b60405281815283820160200185101561094b57600080fd5b816020850160208301376000918101602001919091529392505050565b803580151581146108b657600080fd5b600060e0823603121561098a57600080fd5b610992610869565b61099b83610892565b815260208381013590820152604083013567ffffffffffffffff8111156109c157600080fd5b6109cd368286016108bb565b604083015250606083810135908201526109e960808401610968565b60808201526109fa60a08401610968565b60a082015260c092830135928101929092525090565b60ff8716815285602082015284604082015283606082015260ff8316608082015260c060a08201526000610a4760c083018461068b565b98975050505050505050565b81810381811115610592577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122083969f488afc3adf75c7597aa9b8946b020f56157b1cd22b72ec18db012cd62e64736f6c634300081b0033",
}

// WalletSimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletSimulatorMetaData.ABI instead.
var WalletSimulatorABI = WalletSimulatorMetaData.ABI

// WalletSimulatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletSimulatorMetaData.Bin instead.
var WalletSimulatorBin = WalletSimulatorMetaData.Bin

// DeployWalletSimulator deploys a new Ethereum contract, binding an instance of WalletSimulator to it.
func DeployWalletSimulator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletSimulator, error) {
	parsed, err := WalletSimulatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletSimulatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletSimulator{WalletSimulatorCaller: WalletSimulatorCaller{contract: contract}, WalletSimulatorTransactor: WalletSimulatorTransactor{contract: contract}, WalletSimulatorFilterer: WalletSimulatorFilterer{contract: contract}}, nil
}

// WalletSimulatorDeployedBin is the resulting bytecode of the created contract
const WalletSimulatorDeployedBin = "0x608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a27c392214610030575b600080fd5b61004361003e3660046105e5565b610059565b60405161005091906106ef565b60405180910390f35b606060005a90506000838067ffffffffffffffff81111561007c5761007c6107cd565b6040519080825280602002602001820160405280156100d257816020015b6100bf6040805160608101909152806000815260200160608152602001600081525090565b81526020019060019003908161009a5790505b50935060005b8181101561058d5760008787838181106100f4576100f46107fc565b9050602002810190610106919061082b565b61010f90610978565b905083156101205760009350610130565b8060a00151156101305750610585565b606081015180158015906101435750805a105b156101dd57600587848151811061015c5761015c6107fc565b60200260200101516000019060058111156101795761017961065c565b9081600581111561018c5761018c61065c565b9052505a6040516020016101a291815260200190565b6040516020818303038152906040528784815181106101c3576101c36107fc565b602002602001015160200181905250505050505050610592565b60008260800151156102e35760005a84519091506102ae9084156102015784610203565b5a5b634c4e814c60e01b60008c8a8c60008c6040015160405160240161022c96959493929190610a10565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610598565b91505a6102bb9082610a53565b8986815181106102cd576102cd6107fc565b602002602001015160400181815250505061033d565b60005a8451602086015191925061030c9185156103005785610302565b5a5b87604001516105ae565b91505a6103199082610a53565b89868151811061032b5761032b6107fc565b60200260200101516040018181525050505b806105115760c08301516103c757600195506002888581518110610363576103636107fc565b60200260200101516000019060058111156103805761038061065c565b908160058111156103935761039361065c565b90525061039e6105c6565b8885815181106103b0576103b06107fc565b602002602001015160200181905250505050610585565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161046e576004888581518110610406576104066107fc565b60200260200101516000019060058111156104235761042361065c565b908160058111156104365761043661065c565b9052506104416105c6565b888581518110610453576104536107fc565b60200260200101516020018190525050505050505050610592565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016105115760038885815181106104ad576104ad6107fc565b60200260200101516000019060058111156104ca576104ca61065c565b908160058111156104dd576104dd61065c565b9052506104e86105c6565b8885815181106104fa576104fa6107fc565b60200260200101516020018190525050505061058d565b6001888581518110610525576105256107fc565b60200260200101516000019060058111156105425761054261065c565b908160058111156105555761055561065c565b9052506105606105c6565b888581518110610572576105726107fc565b6020026020010151602001819052505050505b6001016100d8565b505050505b92915050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b600080602083850312156105f857600080fd5b823567ffffffffffffffff81111561060f57600080fd5b8301601f8101851361062057600080fd5b803567ffffffffffffffff81111561063757600080fd5b8560208260051b840101111561064c57600080fd5b6020919091019590945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000815180845260005b818110156106b157602081850181015186830182015201610695565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b828110156107c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110610781577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261079e606088018261068b565b604092830151979092019690965294506020938401939190910190600101610717565b50929695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2183360301811261085f57600080fd5b9190910192915050565b60405160e0810167ffffffffffffffff8111828210171561088c5761088c6107cd565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff811681146108b657600080fd5b919050565b600082601f8301126108cc57600080fd5b813567ffffffffffffffff8111156108e6576108e66107cd565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff81118282101715610933576109336107cd565b60405281815283820160200185101561094b57600080fd5b816020850160208301376000918101602001919091529392505050565b803580151581146108b657600080fd5b600060e0823603121561098a57600080fd5b610992610869565b61099b83610892565b815260208381013590820152604083013567ffffffffffffffff8111156109c157600080fd5b6109cd368286016108bb565b604083015250606083810135908201526109e960808401610968565b60808201526109fa60a08401610968565b60a082015260c092830135928101929092525090565b60ff8716815285602082015284604082015283606082015260ff8316608082015260c060a08201526000610a4760c083018461068b565b98975050505050505050565b81810381811115610592577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122083969f488afc3adf75c7597aa9b8946b020f56157b1cd22b72ec18db012cd62e64736f6c634300081b0033"

// WalletSimulator is an auto generated Go binding around an Ethereum contract.
type WalletSimulator struct {
	WalletSimulatorCaller     // Read-only binding to the contract
	WalletSimulatorTransactor // Write-only binding to the contract
	WalletSimulatorFilterer   // Log filterer for contract events
}

// WalletSimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletSimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletSimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletSimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSimulatorSession struct {
	Contract     *WalletSimulator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletSimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletSimulatorCallerSession struct {
	Contract *WalletSimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WalletSimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletSimulatorTransactorSession struct {
	Contract     *WalletSimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WalletSimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletSimulatorRaw struct {
	Contract *WalletSimulator // Generic contract binding to access the raw methods on
}

// WalletSimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletSimulatorCallerRaw struct {
	Contract *WalletSimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// WalletSimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletSimulatorTransactorRaw struct {
	Contract *WalletSimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletSimulator creates a new instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulator(address common.Address, backend bind.ContractBackend) (*WalletSimulator, error) {
	contract, err := bindWalletSimulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletSimulator{WalletSimulatorCaller: WalletSimulatorCaller{contract: contract}, WalletSimulatorTransactor: WalletSimulatorTransactor{contract: contract}, WalletSimulatorFilterer: WalletSimulatorFilterer{contract: contract}}, nil
}

// NewWalletSimulatorCaller creates a new read-only instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorCaller(address common.Address, caller bind.ContractCaller) (*WalletSimulatorCaller, error) {
	contract, err := bindWalletSimulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCaller{contract: contract}, nil
}

// NewWalletSimulatorTransactor creates a new write-only instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletSimulatorTransactor, error) {
	contract, err := bindWalletSimulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorTransactor{contract: contract}, nil
}

// NewWalletSimulatorFilterer creates a new log filterer instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletSimulatorFilterer, error) {
	contract, err := bindWalletSimulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorFilterer{contract: contract}, nil
}

// bindWalletSimulator binds a generic wrapper to an already deployed contract.
func bindWalletSimulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletSimulatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletSimulator *WalletSimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletSimulator.Contract.WalletSimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletSimulator *WalletSimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletSimulator.Contract.WalletSimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletSimulator *WalletSimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletSimulator.Contract.WalletSimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletSimulator *WalletSimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletSimulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletSimulator *WalletSimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletSimulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletSimulator *WalletSimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletSimulator.Contract.contract.Transact(opts, method, params...)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorTransactor) Simulate(opts *bind.TransactOpts, _calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "simulate", _calls)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorSession) Simulate(_calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Simulate(&_WalletSimulator.TransactOpts, _calls)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorTransactorSession) Simulate(_calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Simulate(&_WalletSimulator.TransactOpts, _calls)
}
