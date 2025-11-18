// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package passkeys

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

// PasskeysMetaData contains all meta data concerning the Passkeys contract.
var PasskeysMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"recoverSapientSignatureCompact\",\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidPasskeySignature\",\"inputs\":[{\"name\":\"_webAuthnAuth\",\"type\":\"tuple\",\"internalType\":\"structWebAuthn.WebAuthnAuth\",\"components\":[{\"name\":\"authenticatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"clientDataJSON\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"challengeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_requireUserVerification\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureLength\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557610b81908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63898bd9211461002757600080fd5b346100bb5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb5760043560243567ffffffffffffffff81116100bb57366023820112156100bb57806004013567ffffffffffffffff81116100bb5736602482840101116100bb576100b79260246100a793019061022c565b6040519081529081906020820190565b0390f35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013057604052565b6100c0565b919082519283825260005b84811061017f5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610140565b90926102236060939695946080845260a06101f26101bf835160c06080890152610140880190610135565b60208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808883030184890152610135565b604083015160c08701528683015160e087015260808301516101008701529101516101208501529615156020840152565b60408201520152565b9291610236610641565b9160009261026d6102478484610685565b357fff000000000000000000000000000000000000000000000000000000000000001690565b957f200000000000000000000000000000000000000000000000000000000000000087166105f1577f0100000000000000000000000000000000000000000000000000000000000000871615159660f881901c6102de6102d860f984901c6001165b60010160ff1690565b60ff1690565b6102f16102d86001600285901c166102cf565b906103176102d86102cf601061030c83600160038a901c1684565b961660041c600f1690565b937f4000000000000000000000000000000000000000000000000000000000000000600191166105e2575b6104e4946104526104d6958b61040e6103bb8d6104036104c19a8f6103bb906103c26103ae6104ac9e6104979e88929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9390840180948b896106f8565b3691610710565b905283929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9081019485926106f8565b60208c01528b929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9060408a015289929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90606088015287602090929192830135920190565b90608087015286602090929192830135920190565b9060a086015285602090929192830135920190565b909481013591602090910190565b949094036105b8576105428484848a61051161053d610546975b6040519283916020830160209181520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826100ef565b61085b565b1590565b61057c57506105799394929190600052602052604060002091600052602052604060002090600052602052604060002090565b90565b6105b48387936040519485947f12a693e600000000000000000000000000000000000000000000000000000000865260048601610194565b0390fd5b7f4be6321b0000000000000000000000000000000000000000000000000000000060005260046000fd5b50600187013598506021610342565b926105119650610546945061053d91506106128161061a92610542956106bd565b81019061079d565b97939592909687929687929b8c919788936104fe565b6040519061063f60c0836100ef565b565b6040519060c0820182811067ffffffffffffffff82111761013057604052600060a08360608152606060208201528260408201528260608201528260808201520152565b901561068e5790565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90929192836001116100bb5783116100bb57600101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b909392938483116100bb5784116100bb578101920390565b92919267ffffffffffffffff82116101305760405191610758601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001846100ef565b8294818452818301116100bb578281602093846000960137010152565b9080601f830112156100bb5781602061057993359101610710565b359081151582036100bb57565b9060a0828203126100bb57813567ffffffffffffffff81116100bb5782019060c0828203126100bb576107ce610630565b91803567ffffffffffffffff81116100bb57826107ec918301610775565b8352602081013567ffffffffffffffff81116100bb5760a092610810918301610775565b6020840152604081013560408401526060810135606084015260808101356080840152013560a08201529161084760208301610790565b916040810135916080606083013592013590565b909493929161086b6000926109a6565b9560208201519081519060408401516060850151908a51918b600d84017f226368616c6c656e6765223a220000000000000000000000000000000000000060981c82528388019060208160138a600d898b0101106022602d8b8801015160001a141695012092012014169185826014011090821760801c109060207f2274797065223a22776562617574686e2e67657422000000000000000000000060581c918801015160581c1416169952835198895191151560021b6001178060218c01511614906020831116169889938461096c575b5050505061094c575b50505050565b61096294955060a0608082015191015191610ab3565b9038808080610946565b60208080959850846001959750840101809882808084519a019601940160025afa831b5afa5192523d156109a457853880808061093d565bfe5b805160609291816109b5575050565b9092506003600284010460021b604051937f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566601f527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f603f5260208501928286019160208301946020828401019060046003835195600085525b0191603f8351818160121c16516000538181600c1c1651600153818160061c165160025316516003536000518152019087821015610a7257600490600390610a2f565b5095506000936003936040925201604052066002048093037f3d3d000000000000000000000000000000000000000000000000000000000000815252038252565b9290939193604051938452602084015283604084015260608301526080820152600080526020600060a0836101005afa503d15610b19575b507f7fffffff800000007fffffffffffffffde737d56d38bcf4279dce5617e3192a860005160011491111090565b60209060a03d916dd01ea45f9efd5c54f037fa57ea1a5afa503d15610b3e5738610aeb565b63d0d5039b3d526004601cfdfea2646970667358221220af52cf218ed3d2bdd1b48fb73cc08c6da906d9535ee4874695357081c29c9ba864736f6c634300081c0033",
}

// PasskeysABI is the input ABI used to generate the binding from.
// Deprecated: Use PasskeysMetaData.ABI instead.
var PasskeysABI = PasskeysMetaData.ABI

// PasskeysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PasskeysMetaData.Bin instead.
var PasskeysBin = PasskeysMetaData.Bin

// DeployPasskeys deploys a new Ethereum contract, binding an instance of Passkeys to it.
func DeployPasskeys(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Passkeys, error) {
	parsed, err := PasskeysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PasskeysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Passkeys{PasskeysCaller: PasskeysCaller{contract: contract}, PasskeysTransactor: PasskeysTransactor{contract: contract}, PasskeysFilterer: PasskeysFilterer{contract: contract}}, nil
}

// Passkeys is an auto generated Go binding around an Ethereum contract.
type Passkeys struct {
	PasskeysCaller     // Read-only binding to the contract
	PasskeysTransactor // Write-only binding to the contract
	PasskeysFilterer   // Log filterer for contract events
}

// PasskeysCaller is an auto generated read-only Go binding around an Ethereum contract.
type PasskeysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PasskeysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PasskeysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PasskeysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PasskeysSession struct {
	Contract     *Passkeys         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PasskeysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PasskeysCallerSession struct {
	Contract *PasskeysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PasskeysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PasskeysTransactorSession struct {
	Contract     *PasskeysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PasskeysRaw is an auto generated low-level Go binding around an Ethereum contract.
type PasskeysRaw struct {
	Contract *Passkeys // Generic contract binding to access the raw methods on
}

// PasskeysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PasskeysCallerRaw struct {
	Contract *PasskeysCaller // Generic read-only contract binding to access the raw methods on
}

// PasskeysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PasskeysTransactorRaw struct {
	Contract *PasskeysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPasskeys creates a new instance of Passkeys, bound to a specific deployed contract.
func NewPasskeys(address common.Address, backend bind.ContractBackend) (*Passkeys, error) {
	contract, err := bindPasskeys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Passkeys{PasskeysCaller: PasskeysCaller{contract: contract}, PasskeysTransactor: PasskeysTransactor{contract: contract}, PasskeysFilterer: PasskeysFilterer{contract: contract}}, nil
}

// NewPasskeysCaller creates a new read-only instance of Passkeys, bound to a specific deployed contract.
func NewPasskeysCaller(address common.Address, caller bind.ContractCaller) (*PasskeysCaller, error) {
	contract, err := bindPasskeys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PasskeysCaller{contract: contract}, nil
}

// NewPasskeysTransactor creates a new write-only instance of Passkeys, bound to a specific deployed contract.
func NewPasskeysTransactor(address common.Address, transactor bind.ContractTransactor) (*PasskeysTransactor, error) {
	contract, err := bindPasskeys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PasskeysTransactor{contract: contract}, nil
}

// NewPasskeysFilterer creates a new log filterer instance of Passkeys, bound to a specific deployed contract.
func NewPasskeysFilterer(address common.Address, filterer bind.ContractFilterer) (*PasskeysFilterer, error) {
	contract, err := bindPasskeys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PasskeysFilterer{contract: contract}, nil
}

// bindPasskeys binds a generic wrapper to an already deployed contract.
func bindPasskeys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PasskeysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Passkeys *PasskeysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Passkeys.Contract.PasskeysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Passkeys *PasskeysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Passkeys.Contract.PasskeysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Passkeys *PasskeysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Passkeys.Contract.PasskeysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Passkeys *PasskeysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Passkeys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Passkeys *PasskeysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Passkeys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Passkeys *PasskeysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Passkeys.Contract.contract.Transact(opts, method, params...)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 _digest, bytes _signature) view returns(bytes32)
func (_Passkeys *PasskeysCaller) RecoverSapientSignatureCompact(opts *bind.CallOpts, _digest [32]byte, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _Passkeys.contract.Call(opts, &out, "recoverSapientSignatureCompact", _digest, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 _digest, bytes _signature) view returns(bytes32)
func (_Passkeys *PasskeysSession) RecoverSapientSignatureCompact(_digest [32]byte, _signature []byte) ([32]byte, error) {
	return _Passkeys.Contract.RecoverSapientSignatureCompact(&_Passkeys.CallOpts, _digest, _signature)
}

// RecoverSapientSignatureCompact is a free data retrieval call binding the contract method 0x898bd921.
//
// Solidity: function recoverSapientSignatureCompact(bytes32 _digest, bytes _signature) view returns(bytes32)
func (_Passkeys *PasskeysCallerSession) RecoverSapientSignatureCompact(_digest [32]byte, _signature []byte) ([32]byte, error) {
	return _Passkeys.Contract.RecoverSapientSignatureCompact(&_Passkeys.CallOpts, _digest, _signature)
}
