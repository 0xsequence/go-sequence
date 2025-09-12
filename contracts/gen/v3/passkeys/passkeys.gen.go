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
	ABI: "[{\"type\":\"function\",\"name\":\"recoverSapientSignatureCompact\",\"inputs\":[{\"name\":\"_digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidPasskeySignature\",\"inputs\":[{\"name\":\"_webAuthnAuth\",\"type\":\"tuple\",\"internalType\":\"structWebAuthn.WebAuthnAuth\",\"components\":[{\"name\":\"authenticatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"clientDataJSON\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"challengeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"typeIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_requireUserVerification\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_x\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60808060405234601557610b7d908161001b8239f35b600080fdfe6080604052600436101561001257600080fd5b60003560e01c63898bd9211461002757600080fd5b346100bb5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100bb5760043560243567ffffffffffffffff81116100bb57366023820112156100bb57806004013567ffffffffffffffff81116100bb5736602482840101116100bb576100b79260246100a793019061022c565b6040519081529081906020820190565b0390f35b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761013057604052565b6100c0565b919082519283825260005b84811061017f5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201610140565b90926102236060939695946080845260a06101f26101bf835160c06080890152610140880190610135565b60208401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808883030184890152610135565b604083015160c08701528683015160e087015260808301516101008701529101516101208501529615156020840152565b60408201520152565b9092919261023861063d565b60009161026e6102488783610681565b357fff000000000000000000000000000000000000000000000000000000000000001690565b957f200000000000000000000000000000000000000000000000000000000000000087166105ed577f01000000000000000000000000000000000000000000000000000000000000008716151596879160f882901c906102e26102dc60f985901c6001165b60010160ff1690565b60ff1690565b6102f56102dc6001600286901c166102d3565b9061031b6102dc6102d3601061031083600160038b901c1684565b971660041c600f1690565b947f4000000000000000000000000000000000000000000000000000000000000000600191166105bd575b6104f06104db6104c66104b1610547996105429e9f9961046c906104fe9f9e9a8f9a61054b9f9a8d61041d610428938f6103d5906105169f6103c86103d5976103dc9288929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9390840180948b896106f4565b369161070c565b905283929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9081019485926106f4565b60208c01528b929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9060408a015289929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90606088015287602090929192830135920190565b90608087015286602090929192830135920190565b9060a086015285602090929192830135920190565b90998a958201359160200190565b50998a955b6040519283916020830160209181520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826100ef565b610857565b1590565b610581575061057e9394929190600052602052604060002091600052602052604060002090600052602052604060002090565b90565b6105b98387936040519485947f12a693e600000000000000000000000000000000000000000000000000000000865260048601610194565b0390fd5b50989296959397508391948691826105d9856001013590602190565b9c909b975092979498505091979399610346565b610516965061054b949350610547925061060e8161061692610542946106b9565b810190610799565b97939592909687929687929b8c91978893610503565b6040519061063b60c0836100ef565b565b6040519060c0820182811067ffffffffffffffff82111761013057604052600060a08360608152606060208201528260408201528260608201528260808201520152565b901561068a5790565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90929192836001116100bb5783116100bb57600101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b909392938483116100bb5784116100bb578101920390565b92919267ffffffffffffffff82116101305760405191610754601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016602001846100ef565b8294818452818301116100bb578281602093846000960137010152565b9080601f830112156100bb5781602061057e9335910161070c565b359081151582036100bb57565b9060a0828203126100bb57813567ffffffffffffffff81116100bb5782019060c0828203126100bb576107ca61062c565b91803567ffffffffffffffff81116100bb57826107e8918301610771565b8352602081013567ffffffffffffffff81116100bb5760a09261080c918301610771565b6020840152604081013560408401526060810135606084015260808101356080840152013560a0820152916108436020830161078c565b916040810135916080606083013592013590565b90949392916108676000926109a2565b9560208201519081519060408401516060850151908a51918b600d84017f226368616c6c656e6765223a220000000000000000000000000000000000000060981c82528388019060208160138a600d898b0101106022602d8b8801015160001a141695012092012014169185826014011090821760801c109060207f2274797065223a22776562617574686e2e67657422000000000000000000000060581c918801015160581c1416169952835198895191151560021b6001178060218c015116149060208311161698899384610968575b50505050610948575b50505050565b61095e94955060a0608082015191015191610aaf565b9038808080610942565b60208080959850846001959750840101809882808084519a019601940160025afa831b5afa5192523d156109a0578538808080610939565bfe5b805160609291816109b1575050565b9092506003600284010460021b604051937f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566601f527f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f603f5260208501928286019160208301946020828401019060046003835195600085525b0191603f8351818160121c16516000538181600c1c1651600153818160061c165160025316516003536000518152019087821015610a6e57600490600390610a2b565b5095506000936003936040925201604052066002048093037f3d3d000000000000000000000000000000000000000000000000000000000000815252038252565b9290939193604051938452602084015283604084015260608301526080820152600080526020600060a0836101005afa503d15610b15575b507f7fffffff800000007fffffffffffffffde737d56d38bcf4279dce5617e3192a860005160011491111090565b60209060a03d916dd01ea45f9efd5c54f037fa57ea1a5afa503d15610b3a5738610ae7565b63d0d5039b3d526004601cfdfea26469706673582212209181eabfef66c6e29c38fd10cd6a76c0062b366560f3faf088a8f3240b7c513f64736f6c634300081c0033",
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
