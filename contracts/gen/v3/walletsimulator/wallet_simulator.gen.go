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
	Bin: "0x6080604052348015600f57600080fd5b506109c58061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a27c392214610030575b600080fd5b61004361003e36600461053b565b610059565b60405161005091906105e1565b60405180910390f35b60606000828067ffffffffffffffff81111561007757610077610712565b6040519080825280602002602001820160405280156100cd57816020015b6100ba6040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816100955790505b50925060005b818110156104e45760008686838181106100ef576100ef610741565b90506020028101906101019190610770565b61010a906108bd565b90508060a00151801561011b575083155b1561012a5760009350506104dc565b6060810151801580159061013d5750805a105b156101d657600586848151811061015657610156610741565b6020026020010151600001906005811115610173576101736105b2565b90816005811115610186576101866105b2565b9052505a60405160200161019c91815260200190565b6040516020818303038152906040528684815181106101bd576101bd610741565b60200260200101516020018190525050505050506104e8565b600082608001511561023b5760005a84519091506102069084156101fa57846101fc565b5a5b86604001516104ee565b91505a6102139082610955565b88868151811061022557610225610741565b6020026020010151604001818152505050610295565b60005a84516020860151919250610264918515610258578561025a565b5a5b8760400151610504565b91505a6102719082610955565b88868151811061028357610283610741565b60200260200101516040018181525050505b806104685760c083015161031f576001955060028785815181106102bb576102bb610741565b60200260200101516000019060058111156102d8576102d86105b2565b908160058111156102eb576102eb6105b2565b9052506102f661051c565b87858151811061030857610308610741565b6020026020010151602001819052505050506104dc565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016103c557600487858151811061035e5761035e610741565b602002602001015160000190600581111561037b5761037b6105b2565b9081600581111561038e5761038e6105b2565b90525061039961051c565b8785815181106103ab576103ab610741565b6020026020010151602001819052505050505050506104e8565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161046857600387858151811061040457610404610741565b6020026020010151600001906005811115610421576104216105b2565b90816005811115610434576104346105b2565b90525061043f61051c565b87858151811061045157610451610741565b6020026020010151602001819052505050506104e4565b600187858151811061047c5761047c610741565b6020026020010151600001906005811115610499576104996105b2565b908160058111156104ac576104ac6105b2565b9052506104b761051c565b8785815181106104c9576104c9610741565b6020026020010151602001819052505050505b6001016100d3565b5050505b92915050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6000806020838503121561054e57600080fd5b823567ffffffffffffffff81111561056557600080fd5b8301601f8101851361057657600080fd5b803567ffffffffffffffff81111561058d57600080fd5b8560208260051b84010111156105a257600080fd5b6020919091019590945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015610706577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110610673577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b80875250602081015160606020880152805180606089015260005b818110156106ab57602081840181015160808b840101520161068e565b5060006080828a0101526040830151604089015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168901019750505050602082019150602084019350600181019050610609565b50929695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218336030181126107a457600080fd5b9190910192915050565b60405160e0810167ffffffffffffffff811182821017156107d1576107d1610712565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff811681146107fb57600080fd5b919050565b600082601f83011261081157600080fd5b813567ffffffffffffffff81111561082b5761082b610712565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff8111828210171561087857610878610712565b60405281815283820160200185101561089057600080fd5b816020850160208301376000918101602001919091529392505050565b803580151581146107fb57600080fd5b600060e082360312156108cf57600080fd5b6108d76107ae565b6108e0836107d7565b815260208381013590820152604083013567ffffffffffffffff81111561090657600080fd5b61091236828601610800565b6040830152506060838101359082015261092e608084016108ad565b608082015261093f60a084016108ad565b60a082015260c092830135928101929092525090565b818103818111156104e8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220106955f5b2841e769d6a8c3be19dc487a535e263ec5208429e2a82b92fe1dc2664736f6c634300081b0033",
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
const WalletSimulatorDeployedBin = "0x608060405234801561001057600080fd5b506004361061002b5760003560e01c8063a27c392214610030575b600080fd5b61004361003e36600461053b565b610059565b60405161005091906105e1565b60405180910390f35b60606000828067ffffffffffffffff81111561007757610077610712565b6040519080825280602002602001820160405280156100cd57816020015b6100ba6040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816100955790505b50925060005b818110156104e45760008686838181106100ef576100ef610741565b90506020028101906101019190610770565b61010a906108bd565b90508060a00151801561011b575083155b1561012a5760009350506104dc565b6060810151801580159061013d5750805a105b156101d657600586848151811061015657610156610741565b6020026020010151600001906005811115610173576101736105b2565b90816005811115610186576101866105b2565b9052505a60405160200161019c91815260200190565b6040516020818303038152906040528684815181106101bd576101bd610741565b60200260200101516020018190525050505050506104e8565b600082608001511561023b5760005a84519091506102069084156101fa57846101fc565b5a5b86604001516104ee565b91505a6102139082610955565b88868151811061022557610225610741565b6020026020010151604001818152505050610295565b60005a84516020860151919250610264918515610258578561025a565b5a5b8760400151610504565b91505a6102719082610955565b88868151811061028357610283610741565b60200260200101516040018181525050505b806104685760c083015161031f576001955060028785815181106102bb576102bb610741565b60200260200101516000019060058111156102d8576102d86105b2565b908160058111156102eb576102eb6105b2565b9052506102f661051c565b87858151811061030857610308610741565b6020026020010151602001819052505050506104dc565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016103c557600487858151811061035e5761035e610741565b602002602001015160000190600581111561037b5761037b6105b2565b9081600581111561038e5761038e6105b2565b90525061039961051c565b8785815181106103ab576103ab610741565b6020026020010151602001819052505050505050506104e8565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161046857600387858151811061040457610404610741565b6020026020010151600001906005811115610421576104216105b2565b90816005811115610434576104346105b2565b90525061043f61051c565b87858151811061045157610451610741565b6020026020010151602001819052505050506104e4565b600187858151811061047c5761047c610741565b6020026020010151600001906005811115610499576104996105b2565b908160058111156104ac576104ac6105b2565b9052506104b761051c565b8785815181106104c9576104c9610741565b6020026020010151602001819052505050505b6001016100d3565b5050505b92915050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6000806020838503121561054e57600080fd5b823567ffffffffffffffff81111561056557600080fd5b8301601f8101851361057657600080fd5b803567ffffffffffffffff81111561058d57600080fd5b8560208260051b84010111156105a257600080fd5b6020919091019590945092505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015610706577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110610673577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b80875250602081015160606020880152805180606089015260005b818110156106ab57602081840181015160808b840101520161068e565b5060006080828a0101526040830151604089015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168901019750505050602082019150602084019350600181019050610609565b50929695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218336030181126107a457600080fd5b9190910192915050565b60405160e0810167ffffffffffffffff811182821017156107d1576107d1610712565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff811681146107fb57600080fd5b919050565b600082601f83011261081157600080fd5b813567ffffffffffffffff81111561082b5761082b610712565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff8111828210171561087857610878610712565b60405281815283820160200185101561089057600080fd5b816020850160208301376000918101602001919091529392505050565b803580151581146107fb57600080fd5b600060e082360312156108cf57600080fd5b6108d76107ae565b6108e0836107d7565b815260208381013590820152604083013567ffffffffffffffff81111561090657600080fd5b61091236828601610800565b6040830152506060838101359082015261092e608084016108ad565b608082015261093f60a084016108ad565b60a082015260c092830135928101929092525090565b818103818111156104e8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220106955f5b2841e769d6a8c3be19dc487a535e263ec5208429e2a82b92fe1dc2664736f6c634300081b0033"

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
