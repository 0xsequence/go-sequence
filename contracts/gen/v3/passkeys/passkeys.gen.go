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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"authenticatorData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"clientDataJSON\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"typeIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWebAuthn.WebAuthnAuth\",\"name\":\"_webAuthnAuth\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_requireUserVerification\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_y\",\"type\":\"bytes32\"}],\"name\":\"InvalidPasskeySignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignatureCompact\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610be88061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063898bd92114610030575b600080fd5b61004361003e3660046107bb565b610055565b60405190815260200160405180910390f35b6000806000806000806100688888610100565b945094509450945094506100a08960405160200161008891815260200190565b60405160208183030381529060405285878686610493565b6100e757848484846040517f12a693e60000000000000000000000000000000000000000000000000000000081526004016100de949392919061089d565b60405180910390fd5b6100f3848484846105cd565b9998505050505050505050565b6040805160c081018252606080825260208201819052600092820183905281018290526080810182905260a0810182905290808080808787828161014657610146610943565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f200000000000000000000000000000000000000000000000000000000000000016600003610461577f010000000000000000000000000000000000000000000000000000000000000081161515945060ff600160f983901c8116810182169160fa84901c8216820181169160fb85901c8116810182169160fc86901c8216820116907f4000000000000000000000000000000000000000000000000000000000000000861615610225578c81013596506020015b60006102738e8e848992810135600884026101008190039190911c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600190921b9190910116939201919050565b8093508192505050600081830190508e8e8490839261029493929190610972565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d60000181905250809250505060006103328e8e848892810135600884026101008190039190911c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600190921b9190910116939201919050565b8093508192505050600081830190508e8e8490839261035393929190610972565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060208e015291505060016008840290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e830135610100929092039190911c1683820160408d0191909152905060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e830135610100929092039190911c1682820160608d019190915290508c8101356020820160808d019190915290508c8101356020820160a08d019190915290508c8101356020820190995090508c8101359750610488945050505050565b61046e876001818b610972565b81019061047b9190610abd565b9399509197509550935091505b509295509295909350565b60008060006104a488600180610614565b905060208601518051602082019150604088015160608901518451600d81017f226368616c6c656e6765223a220000000000000000000000000000000000000060981c87528484820110602282868901015160001a14168160138901208286890120141685846014011085851760801c107f2274797065223a22776562617574686e2e67657422000000000000000000000060581c8589015160581c14161698505080865250505087515189151560021b600117808160218c51015116146020831188161696505085156105a157602089510181810180516020600160208601856020868a8c60025afa60011b5afa51915295503d90506105a157fe5b50505082156105c2576105bf8287608001518860a001518888610725565b92505b505095945050505050565b60008381526020839052604081208560006105f2828660009182526020526040902090565b9050610608838260009182526020526040902090565b98975050505050505050565b60608351801561071d576003600282010460021b60405192507f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566601f526106708515027f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392d5f18603f52602083018181018388602001018051600082525b60038a0199508951603f8160121c1651600053603f81600c1c1651600153603f8160061c1651600253603f8116516003535060005184526004840193508284106106905790526020016040527f3d3d000000000000000000000000000000000000000000000000000000000000600384066002048083039190915260008615159091029182900352900382525b509392505050565b6000604051868152856020820152846040820152836060820152826080820152600080526020600060a0836101005afa503d6107855760203d60a0836dd01ea45f9efd5c54f037fa57ea1a5afa503d6107855763d0d5039b3d526004601cfd5b506000516001147f7fffffff800000007fffffffffffffffde737d56d38bcf4279dce5617e3192a8851110905095945050505050565b6000806000604084860312156107d057600080fd5b83359250602084013567ffffffffffffffff8111156107ee57600080fd5b8401601f810186136107ff57600080fd5b803567ffffffffffffffff81111561081657600080fd5b86602082840101111561082857600080fd5b939660209190910195509293505050565b6000815180845260005b8181101561085f57602081850181015186830182015201610843565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b608081526000855160c060808401526108ba610140840182610839565b905060208701517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808483030160a08501526108f58282610839565b604089015160c0860152606089015160e0860152608089015161010086015260a089015161012086015287151560208601529250610931915050565b60408201939093526060015292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561098257600080fd5b8386111561098f57600080fd5b5050820193919092039150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160c0810167ffffffffffffffff811182821017156109ee576109ee61099c565b60405290565b600082601f830112610a0557600080fd5b81356020830160008067ffffffffffffffff841115610a2657610a2661099c565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff82111715610a7357610a7361099c565b604052838152905080828401871015610a8b57600080fd5b838360208301376000602085830101528094505050505092915050565b80358015158114610ab857600080fd5b919050565b600080600080600060a08688031215610ad557600080fd5b853567ffffffffffffffff811115610aec57600080fd5b860160c08189031215610afe57600080fd5b610b066109cb565b813567ffffffffffffffff811115610b1d57600080fd5b610b298a8285016109f4565b825250602082013567ffffffffffffffff811115610b4657600080fd5b610b528a8285016109f4565b60208381019190915260408481013590840152606080850135908401526080808501359084015260a09384013593830193909352509550610b94908701610aa8565b9497949650505050604083013592606081013592608090910135915056fea2646970667358221220471344837b183327d94afc9a34221c950ebfa5dae4c775d285923ae8457a8e9a64736f6c634300081b0033",
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
