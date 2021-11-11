// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package niftyswap

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
)

// WrapAndNiftyswapMetaData contains all meta data concerning the WrapAndNiftyswap contract.
var WrapAndNiftyswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_tokenWrapper\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc1155\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"erc1155\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_niftyswapOrder\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWrapper\",\"outputs\":[{\"internalType\":\"contractIERC20Wrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_niftyswapOrder\",\"type\":\"bytes\"}],\"name\":\"wrapAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200181f3803806200181f8339810160408190526200003591620001ec565b6001600160a01b038416158015906200005657506001600160a01b03831615155b80156200006b57506001600160a01b03821615155b80156200008057506001600160a01b03811615155b620000a85760405162461bcd60e51b81526004016200009f90620002c2565b60405180910390fd5b6001600160601b0319606085811b821660805284811b821660a05283811b821660c05282901b1660e05260405163095ea7b360e01b81526001600160a01b0383169063095ea7b3906200010490879060001990600401620002a9565b602060405180830381600087803b1580156200011f57600080fd5b505af115801562000134573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200015a919062000253565b506040516318fe01c760e21b81526001600160a01b038516906363f8071c906200018990859060040162000295565b60206040518083038186803b158015620001a257600080fd5b505afa158015620001b7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001dd91906200027c565b61010052506200031292505050565b6000806000806080858703121562000202578384fd5b84516200020f81620002f9565b60208601519094506200022281620002f9565b60408601519093506200023581620002f9565b60608601519092506200024881620002f9565b939692955090935050565b60006020828403121562000265578081fd5b8151801515811462000275578182fd5b9392505050565b6000602082840312156200028e578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6020808252601c908201527f494e56414c494420434f4e5354525543544f5220415247554d454e5400000000604082015260600190565b6001600160a01b03811681146200030f57600080fd5b50565b60805160601c60a05160601c60c05160601c60e05160601c61010051611462620003bd600039806103875280610450528061084e5250806105c652806106aa5280610ac052508061010252806101b2528061029b528061052d528061092b52806109b7525080610365528061077f5280610a9c52508061026e5280610336528061042152806105005280610651528061081f52806108fe5280610a785280610aef52506114626000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063c5e3dfd81161005b578063c5e3dfd8146100d5578063d2f7265a146100dd578063d56022d7146100e5578063f23a6e61146100ed5761007d565b8063785e9e8614610082578063a874d5b6146100a0578063bc197c81146100b5575b600080fd5b61008a610100565b60405161009791906110d3565b60405180910390f35b6100b36100ae366004610fca565b610124565b005b6100c86100c3366004610d17565b610639565b6040516100979190611219565b61008a610a76565b61008a610a9a565b61008a610abe565b6100c86100fb366004610dce565b610ae2565b7f000000000000000000000000000000000000000000000000000000000000000081565b61012c610b58565b61013882840184610e6b565b80519092506001600160a01b03161590508061015d575080516001600160a01b031630145b6101825760405162461bcd60e51b8152600401610179906113a9565b60405180910390fd5b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906101eb90339030908a906004016110e7565b602060405180830381600087803b15801561020557600080fd5b505af1158015610219573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023d9190610e44565b506040517f8340f5490000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638340f549906102c7907f00000000000000000000000000000000000000000000000000000000000000009030908a906004016110e7565b600060405180830381600087803b1580156102e157600080fd5b505af11580156102f5573d6000803e3d6000fd5b50506000805460ff1916600117905550506040517ff242432a0000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063f242432a906103b59030907f0000000000000000000000000000000000000000000000000000000000000000907f0000000000000000000000000000000000000000000000000000000000000000908b908a908a906004016111c7565b600060405180830381600087803b1580156103cf57600080fd5b505af11580156103e3573d6000803e3d6000fd5b50506000805460ff191681556040517efdd58e0000000000000000000000000000000000000000000000000000000081529092506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915062fdd58e906104789030907f000000000000000000000000000000000000000000000000000000000000000090600401611200565b60206040518083038186803b15801561049057600080fd5b505afa1580156104a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c89190610fb2565b9050801561058c576040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d9caed1290610559907f000000000000000000000000000000000000000000000000000000000000000090899086906004016110e7565b600060405180830381600087803b15801561057357600080fd5b505af1158015610587573d6000803e3d6000fd5b505050505b602082015160408084015190517f2eb2c2d60000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001692632eb2c2d6926105ff9230928b92909160040161116f565b600060405180830381600087803b15801561061957600080fd5b505af115801561062d573d6000803e3d6000fd5b50505050505050505050565b6000805460ff16806106735750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b1561069f57507fbc197c8100000000000000000000000000000000000000000000000000000000610a6a565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146106e75760405162461bcd60e51b815260040161017990611326565b6106ef610b89565b6106fb83850185610f35565b80519092506001600160a01b031615905080610720575080516001600160a01b031630145b61073c5760405162461bcd60e51b815260040161017990611246565b6000805460ff191660011790556040517f2eb2c2d60000000000000000000000000000000000000000000000000000000081523390632eb2c2d6906107b39030907f0000000000000000000000000000000000000000000000000000000000000000908d908d908d908d908d908d9060040161110b565b600060405180830381600087803b1580156107cd57600080fd5b505af11580156107e1573d6000803e3d6000fd5b50506000805460ff191681556040517efdd58e0000000000000000000000000000000000000000000000000000000081529092506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016915062fdd58e906108769030907f000000000000000000000000000000000000000000000000000000000000000090600401611200565b60206040518083038186803b15801561088e57600080fd5b505afa1580156108a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c69190610fb2565b90508015610a44576040517fd9caed120000000000000000000000000000000000000000000000000000000081526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063d9caed1290610957907f000000000000000000000000000000000000000000000000000000000000000090309086906004016110e7565b600060405180830381600087803b15801561097157600080fd5b505af1158015610985573d6000803e3d6000fd5b50506040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925063a9059cbb91506109f0908d908590600401611200565b602060405180830381600087803b158015610a0a57600080fd5b505af1158015610a1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a429190610e44565b505b507fbc197c81000000000000000000000000000000000000000000000000000000009150505b98975050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b2c5760405162461bcd60e51b8152600401610179906112c9565b507ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b604051806080016040528060006001600160a01b031681526020016060815260200160608152602001600081525090565b604051806060016040528060006001600160a01b0316815260200160008152602001600081525090565b80356001600160a01b0381168114610bca57600080fd5b919050565b60008083601f840112610be0578182fd5b50813567ffffffffffffffff811115610bf7578182fd5b6020830191508360208083028501011115610c1157600080fd5b9250929050565b600082601f830112610c28578081fd5b813567ffffffffffffffff80821115610c3d57fe5b602080830260405182828201018181108582111715610c5857fe5b604052848152945081850192508582018187018301881015610c7957600080fd5b600091505b84821015610c9c578035845292820192600191909101908201610c7e565b505050505092915050565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610bca57600080fd5b60008083601f840112610ce8578182fd5b50813567ffffffffffffffff811115610cff578182fd5b602083019150836020828501011115610c1157600080fd5b60008060008060008060008060a0898b031215610d32578384fd5b610d3b89610bb3565b9750610d4960208a01610bb3565b9650604089013567ffffffffffffffff80821115610d65578586fd5b610d718c838d01610bcf565b909850965060608b0135915080821115610d89578586fd5b610d958c838d01610bcf565b909650945060808b0135915080821115610dad578384fd5b50610dba8b828c01610cd7565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215610de6578182fd5b610def87610bb3565b9550610dfd60208801610bb3565b94506040870135935060608701359250608087013567ffffffffffffffff811115610e26578283fd5b610e3289828a01610cd7565b979a9699509497509295939492505050565b600060208284031215610e55578081fd5b81518015158114610e64578182fd5b9392505050565b60008060408385031215610e7d578182fd5b610e8683610ca7565b9150602083013567ffffffffffffffff80821115610ea2578283fd5b9084019060808287031215610eb5578283fd5b604051608081018181108382111715610eca57fe5b604052610ed683610bb3565b8152602083013582811115610ee9578485fd5b610ef588828601610c18565b602083015250604083013582811115610f0c578485fd5b610f1888828601610c18565b604083015250606083013560608201528093505050509250929050565b6000808284036080811215610f48578283fd5b610f5184610ca7565b92506060601f1982011215610f64578182fd5b506040516060810181811067ffffffffffffffff82111715610f8257fe5b604052610f9160208501610bb3565b81526040840135602082015260608401356040820152809150509250929050565b600060208284031215610fc3578081fd5b5051919050565b60008060008060608587031215610fdf578384fd5b84359350610fef60208601610bb3565b9250604085013567ffffffffffffffff81111561100a578283fd5b61101687828801610cd7565b95989497509550505050565b60008284527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611053578081fd5b6020830280836020870137939093016020019283525090919050565b6000815180845260208085019450808401835b8381101561109e57815187529582019590820190600101611082565b509495945050505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b60006001600160a01b03808b168352808a1660208401525060a0604083015261113860a08301888a611022565b828103606084015261114b818789611022565b905082810360808401526111608185876110a9565b9b9a5050505050505050505050565b60006001600160a01b03808716835280861660208401525060a0604083015261119b60a083018561106f565b82810360608401526111ad818561106f565b838103608090940193909352508152602001949350505050565b60006001600160a01b03808916835280881660208401525085604083015284606083015260a06080830152610a6a60a0830184866110a9565b6001600160a01b03929092168252602082015260400190565b7fffffffff0000000000000000000000000000000000000000000000000000000091909116815260200190565b6020808252604e908201527f57726170416e644e6966747973776170236f6e4552433131353542617463685260408201527f656365697665643a204f5244455220524543495049454e54204d55535420424560608201527f205448495320434f4e5452414354000000000000000000000000000000000000608082015260a00190565b6020808252603c908201527f57726170416e644e6966747973776170236f6e4552433131353552656365697660408201527f65643a20494e56414c49445f455243313135355f524543454956454400000000606082015260800190565b60208082526041908201527f57726170416e644e6966747973776170236f6e4552433131353542617463685260408201527f656365697665643a20494e56414c49445f455243313135355f5245434549564560608201527f4400000000000000000000000000000000000000000000000000000000000000608082015260a00190565b60208082526043908201527f57726170416e644e69667479737761702377726170416e64537761703a204f5260408201527f44455220524543495049454e54204d555354204245205448495320434f4e545260608201527f4143540000000000000000000000000000000000000000000000000000000000608082015260a0019056fea2646970667358221220ef4c3b3f7ae822cb7fdd18c740da1e395f098d7094f9a22dedd0f200d5b48a1964736f6c63430007040033",
}

// WrapAndNiftyswapABI is the input ABI used to generate the binding from.
// Deprecated: Use WrapAndNiftyswapMetaData.ABI instead.
var WrapAndNiftyswapABI = WrapAndNiftyswapMetaData.ABI

// WrapAndNiftyswapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WrapAndNiftyswapMetaData.Bin instead.
var WrapAndNiftyswapBin = WrapAndNiftyswapMetaData.Bin

// DeployWrapAndNiftyswap deploys a new Ethereum contract, binding an instance of WrapAndNiftyswap to it.
func DeployWrapAndNiftyswap(auth *bind.TransactOpts, backend bind.ContractBackend, _tokenWrapper common.Address, _exchange common.Address, _erc20 common.Address, _erc1155 common.Address) (common.Address, *types.Transaction, *WrapAndNiftyswap, error) {
	parsed, err := WrapAndNiftyswapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WrapAndNiftyswapBin), backend, _tokenWrapper, _exchange, _erc20, _erc1155)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrapAndNiftyswap{WrapAndNiftyswapCaller: WrapAndNiftyswapCaller{contract: contract}, WrapAndNiftyswapTransactor: WrapAndNiftyswapTransactor{contract: contract}, WrapAndNiftyswapFilterer: WrapAndNiftyswapFilterer{contract: contract}}, nil
}

// WrapAndNiftyswap is an auto generated Go binding around an Ethereum contract.
type WrapAndNiftyswap struct {
	WrapAndNiftyswapCaller     // Read-only binding to the contract
	WrapAndNiftyswapTransactor // Write-only binding to the contract
	WrapAndNiftyswapFilterer   // Log filterer for contract events
}

// WrapAndNiftyswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapAndNiftyswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapAndNiftyswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapAndNiftyswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapAndNiftyswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrapAndNiftyswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapAndNiftyswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapAndNiftyswapSession struct {
	Contract     *WrapAndNiftyswap // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapAndNiftyswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapAndNiftyswapCallerSession struct {
	Contract *WrapAndNiftyswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// WrapAndNiftyswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapAndNiftyswapTransactorSession struct {
	Contract     *WrapAndNiftyswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// WrapAndNiftyswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapAndNiftyswapRaw struct {
	Contract *WrapAndNiftyswap // Generic contract binding to access the raw methods on
}

// WrapAndNiftyswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapAndNiftyswapCallerRaw struct {
	Contract *WrapAndNiftyswapCaller // Generic read-only contract binding to access the raw methods on
}

// WrapAndNiftyswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapAndNiftyswapTransactorRaw struct {
	Contract *WrapAndNiftyswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapAndNiftyswap creates a new instance of WrapAndNiftyswap, bound to a specific deployed contract.
func NewWrapAndNiftyswap(address common.Address, backend bind.ContractBackend) (*WrapAndNiftyswap, error) {
	contract, err := bindWrapAndNiftyswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapAndNiftyswap{WrapAndNiftyswapCaller: WrapAndNiftyswapCaller{contract: contract}, WrapAndNiftyswapTransactor: WrapAndNiftyswapTransactor{contract: contract}, WrapAndNiftyswapFilterer: WrapAndNiftyswapFilterer{contract: contract}}, nil
}

// NewWrapAndNiftyswapCaller creates a new read-only instance of WrapAndNiftyswap, bound to a specific deployed contract.
func NewWrapAndNiftyswapCaller(address common.Address, caller bind.ContractCaller) (*WrapAndNiftyswapCaller, error) {
	contract, err := bindWrapAndNiftyswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapAndNiftyswapCaller{contract: contract}, nil
}

// NewWrapAndNiftyswapTransactor creates a new write-only instance of WrapAndNiftyswap, bound to a specific deployed contract.
func NewWrapAndNiftyswapTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapAndNiftyswapTransactor, error) {
	contract, err := bindWrapAndNiftyswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapAndNiftyswapTransactor{contract: contract}, nil
}

// NewWrapAndNiftyswapFilterer creates a new log filterer instance of WrapAndNiftyswap, bound to a specific deployed contract.
func NewWrapAndNiftyswapFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapAndNiftyswapFilterer, error) {
	contract, err := bindWrapAndNiftyswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapAndNiftyswapFilterer{contract: contract}, nil
}

// bindWrapAndNiftyswap binds a generic wrapper to an already deployed contract.
func bindWrapAndNiftyswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrapAndNiftyswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapAndNiftyswap *WrapAndNiftyswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapAndNiftyswap.Contract.WrapAndNiftyswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapAndNiftyswap *WrapAndNiftyswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.WrapAndNiftyswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapAndNiftyswap *WrapAndNiftyswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.WrapAndNiftyswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapAndNiftyswap *WrapAndNiftyswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapAndNiftyswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.contract.Transact(opts, method, params...)
}

// Erc1155 is a free data retrieval call binding the contract method 0xd56022d7.
//
// Solidity: function erc1155() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCaller) Erc1155(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapAndNiftyswap.contract.Call(opts, &out, "erc1155")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc1155 is a free data retrieval call binding the contract method 0xd56022d7.
//
// Solidity: function erc1155() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) Erc1155() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Erc1155(&_WrapAndNiftyswap.CallOpts)
}

// Erc1155 is a free data retrieval call binding the contract method 0xd56022d7.
//
// Solidity: function erc1155() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCallerSession) Erc1155() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Erc1155(&_WrapAndNiftyswap.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapAndNiftyswap.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) Erc20() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Erc20(&_WrapAndNiftyswap.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCallerSession) Erc20() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Erc20(&_WrapAndNiftyswap.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCaller) Exchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapAndNiftyswap.contract.Call(opts, &out, "exchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) Exchange() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Exchange(&_WrapAndNiftyswap.CallOpts)
}

// Exchange is a free data retrieval call binding the contract method 0xd2f7265a.
//
// Solidity: function exchange() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCallerSession) Exchange() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.Exchange(&_WrapAndNiftyswap.CallOpts)
}

// TokenWrapper is a free data retrieval call binding the contract method 0xc5e3dfd8.
//
// Solidity: function tokenWrapper() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCaller) TokenWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapAndNiftyswap.contract.Call(opts, &out, "tokenWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenWrapper is a free data retrieval call binding the contract method 0xc5e3dfd8.
//
// Solidity: function tokenWrapper() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) TokenWrapper() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.TokenWrapper(&_WrapAndNiftyswap.CallOpts)
}

// TokenWrapper is a free data retrieval call binding the contract method 0xc5e3dfd8.
//
// Solidity: function tokenWrapper() view returns(address)
func (_WrapAndNiftyswap *WrapAndNiftyswapCallerSession) TokenWrapper() (common.Address, error) {
	return _WrapAndNiftyswap.Contract.TokenWrapper(&_WrapAndNiftyswap.CallOpts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _niftyswapOrder) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.contract.Transact(opts, "onERC1155BatchReceived", arg0, _from, _ids, _amounts, _niftyswapOrder)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _niftyswapOrder) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.OnERC1155BatchReceived(&_WrapAndNiftyswap.TransactOpts, arg0, _from, _ids, _amounts, _niftyswapOrder)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address _from, uint256[] _ids, uint256[] _amounts, bytes _niftyswapOrder) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactorSession) OnERC1155BatchReceived(arg0 common.Address, _from common.Address, _ids []*big.Int, _amounts []*big.Int, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.OnERC1155BatchReceived(&_WrapAndNiftyswap.TransactOpts, arg0, _from, _ids, _amounts, _niftyswapOrder)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.OnERC1155Received(&_WrapAndNiftyswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.OnERC1155Received(&_WrapAndNiftyswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactor) WrapAndSwap(opts *bind.TransactOpts, _maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.contract.Transact(opts, "wrapAndSwap", _maxAmount, _recipient, _niftyswapOrder)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_WrapAndNiftyswap *WrapAndNiftyswapSession) WrapAndSwap(_maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.WrapAndSwap(&_WrapAndNiftyswap.TransactOpts, _maxAmount, _recipient, _niftyswapOrder)
}

// WrapAndSwap is a paid mutator transaction binding the contract method 0xa874d5b6.
//
// Solidity: function wrapAndSwap(uint256 _maxAmount, address _recipient, bytes _niftyswapOrder) returns()
func (_WrapAndNiftyswap *WrapAndNiftyswapTransactorSession) WrapAndSwap(_maxAmount *big.Int, _recipient common.Address, _niftyswapOrder []byte) (*types.Transaction, error) {
	return _WrapAndNiftyswap.Contract.WrapAndSwap(&_WrapAndNiftyswap.TransactOpts, _maxAmount, _recipient, _niftyswapOrder)
}
