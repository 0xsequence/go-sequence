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

// WalletGuestMetaData contains all meta data concerning the WalletGuest contract.
var WalletGuestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"InvalidSValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_v\",\"type\":\"uint256\"}],\"name\":\"InvalidVValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_available\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDelegatecall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"OnlySelfAuth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"SignerIsAddress0\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_recoverMode\",\"type\":\"bool\"}],\"name\":\"UnsupportedSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"CreatedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"TxExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"TxFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SET_IMAGE_HASH_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"signatureRecovery\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b503060805260805161213761002d600039600050506121376000f3fe6080604052600436106100bc5760003560e01c806361c2926c116100745780638c3f55631161004e5780638c3f55631461025357806390042baf14610273578063affed0e0146102ab57600080fd5b806361c2926c146101cb5780637a9a1628146101eb578063853c50681461020b57600080fd5b806320c13b0b116100a557806320c13b0b14610147578063295614261461016757806357c56d6b1461018957600080fd5b806301ffc9a7146100c15780631626ba7e146100f6575b600080fd5b3480156100cd57600080fd5b506100e16100dc3660046117f8565b6102c0565b60405190151581526020015b60405180910390f35b34801561010257600080fd5b5061011661011136600461185e565b6102d1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016100ed565b34801561015357600080fd5b506101166101623660046118aa565b61031e565b34801561017357600080fd5b50610187610182366004611916565b610383565b005b34801561019557600080fd5b506101bd7f8713a7c4465f6fbee2b6e9d6646d1d9f83fec929edfc4baf661f3c865bdd04d181565b6040519081526020016100ed565b3480156101d757600080fd5b506101876101e6366004611974565b6103d5565b3480156101f757600080fd5b506101876102063660046119b6565b61041a565b34801561021757600080fd5b5061022b61022636600461185e565b610447565b604080519586526020860194909452928401919091526060830152608082015260a0016100ed565b34801561025f57600080fd5b506101bd61026e366004611916565b61060f565b610286610281366004611a5f565b61063b565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b3480156102b757600080fd5b506101bd6106d7565b60006102cb826106e8565b92915050565b6000806102df858585610744565b509050801561031157507f1626ba7e000000000000000000000000000000000000000000000000000000009050610317565b50600090505b9392505050565b6000806103438686604051610334929190611b2e565b60405180910390208585610744565b509050801561037557507f20c13b0b00000000000000000000000000000000000000000000000000000000905061037b565b50600090505b949350505050565b3330146103c9576040517fe12588940000000000000000000000000000000000000000000000000000000081523360048201523060248201526044015b60405180910390fd5b6103d28161077c565b50565b600061040883836040516020016103ed929190611d0c565b604051602081830303815290604052805190602001206107ae565b9050610415818484610833565b505050565b600061043286866040516020016103ed929190611d54565b905061043f818787610833565b505050505050565b6000806000806000808787600081811061046357610463611d9c565b909101357fff000000000000000000000000000000000000000000000000000000000000001691508190506104b95761049b896107ae565b92506104a88389896109c0565b929850909650945091506106049050565b7fff00000000000000000000000000000000000000000000000000000000000000818116016104f8576104eb896107ae565b92506104a8838989610a11565b7ffe000000000000000000000000000000000000000000000000000000000000007fff0000000000000000000000000000000000000000000000000000000000000082160161054a576104eb89610a3d565b7ffd000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000008216016105ae5761059e898989610aaa565b9550955095509550955050610604565b6040517f6085cd820000000000000000000000000000000000000000000000000000000081527fff00000000000000000000000000000000000000000000000000000000000000821660048201526024016103c0565b939792965093509350565b60006102cb7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c27565b600033301461067e576040517fe12588940000000000000000000000000000000000000000000000000000000081523360048201523060248201526044016103c0565b81516020830134f060405173ffffffffffffffffffffffffffffffffffffffff821681529091507fa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c9060200160405180910390a1919050565b60006106e3600061060f565b905090565b60007f6ffbd451000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083160161073b57506001919050565b6102cb82610c85565b6000806000806000610757888888610447565b5096509194509250905082821080159061076f575060015b9450505050935093915050565b6040517fa038794000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f190100000000000000000000000000000000000000000000000000000000000060208201524660228201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003060601b166042820152605681018290526000906076015b604051602081830303815290604052805190602001209050919050565b8060005b818110156109b9573684848381811061085257610852611d9c565b90506020028101906108649190611dcb565b90506108736020820182611e09565b156108ad576040517f230d1ccc000000000000000000000000000000000000000000000000000000008152600481018390526024016103c0565b6040810135805a10156109005782815a6040517f2bb3e3ba0000000000000000000000000000000000000000000000000000000081526004810193909352602483019190915260448201526064016103c0565b600061093a6109156080850160608601611e24565b608085013584156109265784610928565b5a5b61093560a0880188611e3f565b610ce1565b9050801561098157877f5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae78560405161097491815260200190565b60405180910390a26109a3565b6109a36109946040850160208601611e09565b898661099e610cfe565b610d1d565b50505080806109b190611ed3565b915050610837565b5050505050565b60008080806109db876109d6876006818b611f0b565b610d6b565b6000908152873560f01c6020818152604080842084526002909a013560e01c908190529890912090999198509695509350505050565b6000808080610a2c87610a27876001818b611f0b565b6109c0565b935093509350935093509350935093565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526000602282018190527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003060601b1660428301526056820183905290607601610816565b6000808080806004600188013560e81c82610ac58383611f35565b9050610ad78b61022683868d8f611f0b565b939b5091995097509550935087871015610b2f57610af781848b8d611f0b565b89896040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016103c09493929190611f48565b8092505b88831015610c195760038301928a013560e81c9150610b528383611f35565b90506000610b74610b6288611201565b8c8c8790869261022693929190611f0b565b939c50919a5098509091505088881015610bcc57610b9482858c8e611f0b565b8a8a6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016103c09493929190611f48565b848110610c0f576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016103c0565b9350915081610b33565b505050939792965093509350565b6000808383604051602001610c46929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b60007fe4a77bbc000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610cd857506001919050565b6102cb82611235565b6000604051828482376000808483898b8af1979650505050505050565b60603d604051915060208201818101604052818352816000823e505090565b8315610d2b57805160208201fd5b827fab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b4198383604051610d5d929190611f6f565b60405180910390a250505050565b60008060005b838110156111f857600181019085013560f81c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101610e1257601582019186013560f881901c9060581c73ffffffffffffffffffffffffffffffffffffffff81169074ff000000000000000000000000000000000000000016811785610df85780610e07565b60008681526020829052604090205b955050505050610d71565b80610ea85760018201918681013560f81c906043016000610e3e8a610e3984888c8e611f0b565b61131f565b60ff841697909701969194508491905060a083901b74ff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff82161786610e8d5780610e9c565b60008781526020829052604090205b96505050505050610d71565b60028103610fd0576000808784013560f881901c9060581c73ffffffffffffffffffffffffffffffffffffffff16601586019550909250905060008885013560e81c600386018162ffffff169150809650819250505060008186019050610f218b848c8c8a908692610f1c93929190611f0b565b6115e2565b610f69578a83610f3383898d8f611f0b565b6040517f9a9462320000000000000000000000000000000000000000000000000000000081526004016103c09493929190611fe3565b60ff8416979097019694508460a084901b74ff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161787610fb45780610fc3565b60008881526020829052604090205b9750505050505050610d71565b6003810361100357602082019186013583610feb5780610ffa565b60008481526020829052604090205b93505050610d71565b6004810361104f576003808301928781013560e81c91908201016000806110308b6109d685898d8f611f0b565b60009889526020526040909720969097019650909350610d7192505050565b600681036111575760008287013560f81c60018401935060ff16905060008784013560f01c60028501945061ffff16905060008885013560e81c600386018162ffffff1691508096508192505050600081860190506000806110bd8d8d8d8b9087926109d693929190611f0b565b939850889390925090508482106110d357988501985b604080517f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000602080830191909152603882018490526058820188905260788083018a90528351808403909101815260989092019092528051910120896111395780611148565b60008a81526020829052604090205b99505050505050505050610d71565b600581036111c3576020820191860135878103611192577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff94505b600061119d8261178f565b9050846111aa57806111b9565b60008581526020829052604090205b9450505050610d71565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016103c0565b50935093915050565b7f8713a7c4465f6fbee2b6e9d6646d1d9f83fec929edfc4baf661f3c865bdd04d160009081526020829052604081206102cb565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fac6a444e0000000000000000000000000000000000000000000000000000000014806112c857507fffffffff0000000000000000000000000000000000000000000000000000000082167f36e7817500000000000000000000000000000000000000000000000000000000145b156112d557506001919050565b7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146102cb565b60006042821461135f5782826040517f2ee17a3d0000000000000000000000000000000000000000000000000000000081526004016103c0929190612023565b600061137861136f600185612037565b85013560f81c90565b60ff169050604084013560f81c843560208601357f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08111156113ec578686826040517fad4aac760000000000000000000000000000000000000000000000000000000081526004016103c09392919061204a565b8260ff16601b1415801561140457508260ff16601c14155b15611441578686846040517fe578897e0000000000000000000000000000000000000000000000000000000081526004016103c09392919061206e565b600184036114ae576040805160008152602081018083528a905260ff851691810191909152606081018390526080810182905260019060a0015b6020604051602081039080840390855afa15801561149d573d6000803e3d6000fd5b505050602060405103519450611586565b6002840361154b576040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101899052600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff861690820152606081018490526080810183905260a00161147b565b86868560016040517f9dfba8520000000000000000000000000000000000000000000000000000000081526004016103c09493929190612095565b73ffffffffffffffffffffffffffffffffffffffff85166115d75786866040517f6c1719d20000000000000000000000000000000000000000000000000000000081526004016103c0929190612023565b505050509392505050565b60008083836115f2600182612037565b81811061160157611601611d9c565b919091013560f81c915050600181148061161b5750600281145b15611660578473ffffffffffffffffffffffffffffffffffffffff1661164287868661131f565b73ffffffffffffffffffffffffffffffffffffffff16149150611786565b6003810361174b5773ffffffffffffffffffffffffffffffffffffffff8516631626ba7e8786600087611694600182612037565b926116a193929190611f0b565b6040518463ffffffff1660e01b81526004016116bf939291906120c1565b602060405180830381865afa1580156116dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170091906120e4565b7fffffffff00000000000000000000000000000000000000000000000000000000167f1626ba7e00000000000000000000000000000000000000000000000000000000149150611786565b83838260006040517f9dfba8520000000000000000000000000000000000000000000000000000000081526004016103c09493929190612095565b50949350505050565b6040517f53657175656e636520737461746963206469676573743a0a0000000000000000602082015260388101829052600090605801610816565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146103d257600080fd5b60006020828403121561180a57600080fd5b8135610317816117ca565b60008083601f84011261182757600080fd5b50813567ffffffffffffffff81111561183f57600080fd5b60208301915083602082850101111561185757600080fd5b9250929050565b60008060006040848603121561187357600080fd5b83359250602084013567ffffffffffffffff81111561189157600080fd5b61189d86828701611815565b9497909650939450505050565b600080600080604085870312156118c057600080fd5b843567ffffffffffffffff808211156118d857600080fd5b6118e488838901611815565b909650945060208701359150808211156118fd57600080fd5b5061190a87828801611815565b95989497509550505050565b60006020828403121561192857600080fd5b5035919050565b60008083601f84011261194157600080fd5b50813567ffffffffffffffff81111561195957600080fd5b6020830191508360208260051b850101111561185757600080fd5b6000806020838503121561198757600080fd5b823567ffffffffffffffff81111561199e57600080fd5b6119aa8582860161192f565b90969095509350505050565b6000806000806000606086880312156119ce57600080fd5b853567ffffffffffffffff808211156119e657600080fd5b6119f289838a0161192f565b9097509550602088013594506040880135915080821115611a1257600080fd5b50611a1f88828901611815565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215611a7157600080fd5b813567ffffffffffffffff80821115611a8957600080fd5b818401915084601f830112611a9d57600080fd5b813581811115611aaf57611aaf611a30565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611af557611af5611a30565b81604052828152876020848701011115611b0e57600080fd5b826020860160208301376000928101602001929092525095945050505050565b8183823760009101908152919050565b80358015158114611b4e57600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611b4e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b818352600060208085019450848460051b86018460005b87811015611cff57838303895281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41883603018112611c1657600080fd5b870160c0611c2382611b3e565b15158552611c32878301611b3e565b15158588015260408281013590860152606073ffffffffffffffffffffffffffffffffffffffff611c64828501611b53565b16908601526080828101359086015260a080830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611caa57600080fd5b90920187810192903567ffffffffffffffff811115611cc857600080fd5b803603841315611cd757600080fd5b8282880152611ce98388018286611b77565b9c89019c96505050928601925050600101611bd7565b5090979650505050505050565b60408152600560408201527f73656c663a000000000000000000000000000000000000000000000000000000606082015260806020820152600061037b608083018486611bc0565b60408152600660408201527f67756573743a0000000000000000000000000000000000000000000000000000606082015260806020820152600061037b608083018486611bc0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41833603018112611dff57600080fd5b9190910192915050565b600060208284031215611e1b57600080fd5b61031782611b3e565b600060208284031215611e3657600080fd5b61031782611b53565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611e7457600080fd5b83018035915067ffffffffffffffff821115611e8f57600080fd5b60200191503681900382131561185757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f0457611f04611ea4565b5060010190565b60008085851115611f1b57600080fd5b83861115611f2857600080fd5b5050820193919092039150565b808201808211156102cb576102cb611ea4565b606081526000611f5c606083018688611b77565b6020830194909452506040015292915050565b82815260006020604081840152835180604085015260005b81811015611fa357858101830151858201606001528201611f87565b5060006060828601015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509392505050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000612019606083018486611b77565b9695505050505050565b60208152600061037b602083018486611b77565b818103818111156102cb576102cb611ea4565b60408152600061205e604083018587611b77565b9050826020830152949350505050565b604081526000612082604083018587611b77565b905060ff83166020830152949350505050565b6060815260006120a9606083018688611b77565b60208301949094525090151560409091015292915050565b8381526040602082015260006120db604083018486611b77565b95945050505050565b6000602082840312156120f657600080fd5b8151610317816117ca56fea2646970667358221220bee73eb30cfa01fe71819f2f5092b01a4c803b4fb92e0a4e700d47ff596892ab64736f6c63430008120033",
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
	parsed, err := abi.JSON(strings.NewReader(WalletGuestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// SETIMAGEHASHTYPEHASH is a free data retrieval call binding the contract method 0x57c56d6b.
//
// Solidity: function SET_IMAGE_HASH_TYPE_HASH() view returns(bytes32)
func (_WalletGuest *WalletGuestCaller) SETIMAGEHASHTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "SET_IMAGE_HASH_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETIMAGEHASHTYPEHASH is a free data retrieval call binding the contract method 0x57c56d6b.
//
// Solidity: function SET_IMAGE_HASH_TYPE_HASH() view returns(bytes32)
func (_WalletGuest *WalletGuestSession) SETIMAGEHASHTYPEHASH() ([32]byte, error) {
	return _WalletGuest.Contract.SETIMAGEHASHTYPEHASH(&_WalletGuest.CallOpts)
}

// SETIMAGEHASHTYPEHASH is a free data retrieval call binding the contract method 0x57c56d6b.
//
// Solidity: function SET_IMAGE_HASH_TYPE_HASH() view returns(bytes32)
func (_WalletGuest *WalletGuestCallerSession) SETIMAGEHASHTYPEHASH() ([32]byte, error) {
	return _WalletGuest.Contract.SETIMAGEHASHTYPEHASH(&_WalletGuest.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signatures []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "isValidSignature", _hash, _signatures)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestSession) IsValidSignature(_hash [32]byte, _signatures []byte) ([4]byte, error) {
	return _WalletGuest.Contract.IsValidSignature(&_WalletGuest.CallOpts, _hash, _signatures)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestCallerSession) IsValidSignature(_hash [32]byte, _signatures []byte) ([4]byte, error) {
	return _WalletGuest.Contract.IsValidSignature(&_WalletGuest.CallOpts, _hash, _signatures)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestCaller) IsValidSignature0(opts *bind.CallOpts, _data []byte, _signatures []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "isValidSignature0", _data, _signatures)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestSession) IsValidSignature0(_data []byte, _signatures []byte) ([4]byte, error) {
	return _WalletGuest.Contract.IsValidSignature0(&_WalletGuest.CallOpts, _data, _signatures)
}

// IsValidSignature0 is a free data retrieval call binding the contract method 0x20c13b0b.
//
// Solidity: function isValidSignature(bytes _data, bytes _signatures) view returns(bytes4)
func (_WalletGuest *WalletGuestCallerSession) IsValidSignature0(_data []byte, _signatures []byte) ([4]byte, error) {
	return _WalletGuest.Contract.IsValidSignature0(&_WalletGuest.CallOpts, _data, _signatures)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_WalletGuest *WalletGuestCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_WalletGuest *WalletGuestSession) Nonce() (*big.Int, error) {
	return _WalletGuest.Contract.Nonce(&_WalletGuest.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_WalletGuest *WalletGuestCallerSession) Nonce() (*big.Int, error) {
	return _WalletGuest.Contract.Nonce(&_WalletGuest.CallOpts)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletGuest *WalletGuestCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletGuest *WalletGuestSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletGuest.Contract.ReadNonce(&_WalletGuest.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletGuest *WalletGuestCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletGuest.Contract.ReadNonce(&_WalletGuest.CallOpts, _space)
}

// SignatureRecovery is a free data retrieval call binding the contract method 0x853c5068.
//
// Solidity: function signatureRecovery(bytes32 _digest, bytes _signature) view returns(uint256 threshold, uint256 weight, bytes32 imageHash, bytes32 subdigest, uint256 checkpoint)
func (_WalletGuest *WalletGuestCaller) SignatureRecovery(opts *bind.CallOpts, _digest [32]byte, _signature []byte) (struct {
	Threshold  *big.Int
	Weight     *big.Int
	ImageHash  [32]byte
	Subdigest  [32]byte
	Checkpoint *big.Int
}, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "signatureRecovery", _digest, _signature)

	outstruct := new(struct {
		Threshold  *big.Int
		Weight     *big.Int
		ImageHash  [32]byte
		Subdigest  [32]byte
		Checkpoint *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Threshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ImageHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Subdigest = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Checkpoint = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SignatureRecovery is a free data retrieval call binding the contract method 0x853c5068.
//
// Solidity: function signatureRecovery(bytes32 _digest, bytes _signature) view returns(uint256 threshold, uint256 weight, bytes32 imageHash, bytes32 subdigest, uint256 checkpoint)
func (_WalletGuest *WalletGuestSession) SignatureRecovery(_digest [32]byte, _signature []byte) (struct {
	Threshold  *big.Int
	Weight     *big.Int
	ImageHash  [32]byte
	Subdigest  [32]byte
	Checkpoint *big.Int
}, error) {
	return _WalletGuest.Contract.SignatureRecovery(&_WalletGuest.CallOpts, _digest, _signature)
}

// SignatureRecovery is a free data retrieval call binding the contract method 0x853c5068.
//
// Solidity: function signatureRecovery(bytes32 _digest, bytes _signature) view returns(uint256 threshold, uint256 weight, bytes32 imageHash, bytes32 subdigest, uint256 checkpoint)
func (_WalletGuest *WalletGuestCallerSession) SignatureRecovery(_digest [32]byte, _signature []byte) (struct {
	Threshold  *big.Int
	Weight     *big.Int
	ImageHash  [32]byte
	Subdigest  [32]byte
	Checkpoint *big.Int
}, error) {
	return _WalletGuest.Contract.SignatureRecovery(&_WalletGuest.CallOpts, _digest, _signature)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_WalletGuest *WalletGuestCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _WalletGuest.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_WalletGuest *WalletGuestSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _WalletGuest.Contract.SupportsInterface(&_WalletGuest.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) pure returns(bool)
func (_WalletGuest *WalletGuestCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _WalletGuest.Contract.SupportsInterface(&_WalletGuest.CallOpts, _interfaceID)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletGuest *WalletGuestTransactor) CreateContract(opts *bind.TransactOpts, _code []byte) (*types.Transaction, error) {
	return _WalletGuest.contract.Transact(opts, "createContract", _code)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletGuest *WalletGuestSession) CreateContract(_code []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.CreateContract(&_WalletGuest.TransactOpts, _code)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletGuest *WalletGuestTransactorSession) CreateContract(_code []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.CreateContract(&_WalletGuest.TransactOpts, _code)
}

// Execute is a paid mutator transaction binding the contract method 0x7a9a1628.
//
// Solidity: function execute((bool,bool,uint256,address,uint256,bytes)[] _txs, uint256 , bytes ) returns()
func (_WalletGuest *WalletGuestTransactor) Execute(opts *bind.TransactOpts, _txs []IModuleCallsTransaction, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletGuest.contract.Transact(opts, "execute", _txs, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x7a9a1628.
//
// Solidity: function execute((bool,bool,uint256,address,uint256,bytes)[] _txs, uint256 , bytes ) returns()
func (_WalletGuest *WalletGuestSession) Execute(_txs []IModuleCallsTransaction, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Execute(&_WalletGuest.TransactOpts, _txs, arg1, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0x7a9a1628.
//
// Solidity: function execute((bool,bool,uint256,address,uint256,bytes)[] _txs, uint256 , bytes ) returns()
func (_WalletGuest *WalletGuestTransactorSession) Execute(_txs []IModuleCallsTransaction, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Execute(&_WalletGuest.TransactOpts, _txs, arg1, arg2)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x61c2926c.
//
// Solidity: function selfExecute((bool,bool,uint256,address,uint256,bytes)[] _txs) returns()
func (_WalletGuest *WalletGuestTransactor) SelfExecute(opts *bind.TransactOpts, _txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletGuest.contract.Transact(opts, "selfExecute", _txs)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x61c2926c.
//
// Solidity: function selfExecute((bool,bool,uint256,address,uint256,bytes)[] _txs) returns()
func (_WalletGuest *WalletGuestSession) SelfExecute(_txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletGuest.Contract.SelfExecute(&_WalletGuest.TransactOpts, _txs)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x61c2926c.
//
// Solidity: function selfExecute((bool,bool,uint256,address,uint256,bytes)[] _txs) returns()
func (_WalletGuest *WalletGuestTransactorSession) SelfExecute(_txs []IModuleCallsTransaction) (*types.Transaction, error) {
	return _WalletGuest.Contract.SelfExecute(&_WalletGuest.TransactOpts, _txs)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletGuest *WalletGuestTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletGuest.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletGuest *WalletGuestSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.UpdateImageHash(&_WalletGuest.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletGuest *WalletGuestTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.UpdateImageHash(&_WalletGuest.TransactOpts, _imageHash)
}

// WalletGuestCreatedContractIterator is returned from FilterCreatedContract and is used to iterate over the raw logs and unpacked data for CreatedContract events raised by the WalletGuest contract.
type WalletGuestCreatedContractIterator struct {
	Event *WalletGuestCreatedContract // Event containing the contract specifics and raw log

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
func (it *WalletGuestCreatedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestCreatedContract)
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
		it.Event = new(WalletGuestCreatedContract)
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
func (it *WalletGuestCreatedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestCreatedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestCreatedContract represents a CreatedContract event raised by the WalletGuest contract.
type WalletGuestCreatedContract struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatedContract is a free log retrieval operation binding the contract event 0xa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c.
//
// Solidity: event CreatedContract(address _contract)
func (_WalletGuest *WalletGuestFilterer) FilterCreatedContract(opts *bind.FilterOpts) (*WalletGuestCreatedContractIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "CreatedContract")
	if err != nil {
		return nil, err
	}
	return &WalletGuestCreatedContractIterator{contract: _WalletGuest.contract, event: "CreatedContract", logs: logs, sub: sub}, nil
}

// WatchCreatedContract is a free log subscription operation binding the contract event 0xa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c.
//
// Solidity: event CreatedContract(address _contract)
func (_WalletGuest *WalletGuestFilterer) WatchCreatedContract(opts *bind.WatchOpts, sink chan<- *WalletGuestCreatedContract) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "CreatedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestCreatedContract)
				if err := _WalletGuest.contract.UnpackLog(event, "CreatedContract", log); err != nil {
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

// ParseCreatedContract is a log parse operation binding the contract event 0xa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c.
//
// Solidity: event CreatedContract(address _contract)
func (_WalletGuest *WalletGuestFilterer) ParseCreatedContract(log types.Log) (*WalletGuestCreatedContract, error) {
	event := new(WalletGuestCreatedContract)
	if err := _WalletGuest.contract.UnpackLog(event, "CreatedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletGuest contract.
type WalletGuestImageHashUpdatedIterator struct {
	Event *WalletGuestImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletGuestImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestImageHashUpdated)
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
		it.Event = new(WalletGuestImageHashUpdated)
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
func (it *WalletGuestImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestImageHashUpdated represents a ImageHashUpdated event raised by the WalletGuest contract.
type WalletGuestImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletGuest *WalletGuestFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletGuestImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletGuestImageHashUpdatedIterator{contract: _WalletGuest.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletGuest *WalletGuestFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletGuestImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestImageHashUpdated)
				if err := _WalletGuest.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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

// ParseImageHashUpdated is a log parse operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletGuest *WalletGuestFilterer) ParseImageHashUpdated(log types.Log) (*WalletGuestImageHashUpdated, error) {
	event := new(WalletGuestImageHashUpdated)
	if err := _WalletGuest.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletGuest contract.
type WalletGuestNonceChangeIterator struct {
	Event *WalletGuestNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletGuestNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestNonceChange)
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
		it.Event = new(WalletGuestNonceChange)
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
func (it *WalletGuestNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestNonceChange represents a NonceChange event raised by the WalletGuest contract.
type WalletGuestNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletGuest *WalletGuestFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletGuestNonceChangeIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletGuestNonceChangeIterator{contract: _WalletGuest.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletGuest *WalletGuestFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletGuestNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestNonceChange)
				if err := _WalletGuest.contract.UnpackLog(event, "NonceChange", log); err != nil {
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

// ParseNonceChange is a log parse operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletGuest *WalletGuestFilterer) ParseNonceChange(log types.Log) (*WalletGuestNonceChange, error) {
	event := new(WalletGuestNonceChange)
	if err := _WalletGuest.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestTxExecutedIterator is returned from FilterTxExecuted and is used to iterate over the raw logs and unpacked data for TxExecuted events raised by the WalletGuest contract.
type WalletGuestTxExecutedIterator struct {
	Event *WalletGuestTxExecuted // Event containing the contract specifics and raw log

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
func (it *WalletGuestTxExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestTxExecuted)
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
		it.Event = new(WalletGuestTxExecuted)
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
func (it *WalletGuestTxExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestTxExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestTxExecuted represents a TxExecuted event raised by the WalletGuest contract.
type WalletGuestTxExecuted struct {
	Tx    [32]byte
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTxExecuted is a free log retrieval operation binding the contract event 0x5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae7.
//
// Solidity: event TxExecuted(bytes32 indexed _tx, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) FilterTxExecuted(opts *bind.FilterOpts, _tx [][32]byte) (*WalletGuestTxExecutedIterator, error) {

	var _txRule []interface{}
	for _, _txItem := range _tx {
		_txRule = append(_txRule, _txItem)
	}

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "TxExecuted", _txRule)
	if err != nil {
		return nil, err
	}
	return &WalletGuestTxExecutedIterator{contract: _WalletGuest.contract, event: "TxExecuted", logs: logs, sub: sub}, nil
}

// WatchTxExecuted is a free log subscription operation binding the contract event 0x5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae7.
//
// Solidity: event TxExecuted(bytes32 indexed _tx, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) WatchTxExecuted(opts *bind.WatchOpts, sink chan<- *WalletGuestTxExecuted, _tx [][32]byte) (event.Subscription, error) {

	var _txRule []interface{}
	for _, _txItem := range _tx {
		_txRule = append(_txRule, _txItem)
	}

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "TxExecuted", _txRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestTxExecuted)
				if err := _WalletGuest.contract.UnpackLog(event, "TxExecuted", log); err != nil {
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

// ParseTxExecuted is a log parse operation binding the contract event 0x5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae7.
//
// Solidity: event TxExecuted(bytes32 indexed _tx, uint256 _index)
func (_WalletGuest *WalletGuestFilterer) ParseTxExecuted(log types.Log) (*WalletGuestTxExecuted, error) {
	event := new(WalletGuestTxExecuted)
	if err := _WalletGuest.contract.UnpackLog(event, "TxExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletGuestTxFailedIterator is returned from FilterTxFailed and is used to iterate over the raw logs and unpacked data for TxFailed events raised by the WalletGuest contract.
type WalletGuestTxFailedIterator struct {
	Event *WalletGuestTxFailed // Event containing the contract specifics and raw log

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
func (it *WalletGuestTxFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletGuestTxFailed)
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
		it.Event = new(WalletGuestTxFailed)
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
func (it *WalletGuestTxFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletGuestTxFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletGuestTxFailed represents a TxFailed event raised by the WalletGuest contract.
type WalletGuestTxFailed struct {
	Tx     [32]byte
	Index  *big.Int
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxFailed is a free log retrieval operation binding the contract event 0xab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b419.
//
// Solidity: event TxFailed(bytes32 indexed _tx, uint256 _index, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) FilterTxFailed(opts *bind.FilterOpts, _tx [][32]byte) (*WalletGuestTxFailedIterator, error) {

	var _txRule []interface{}
	for _, _txItem := range _tx {
		_txRule = append(_txRule, _txItem)
	}

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "TxFailed", _txRule)
	if err != nil {
		return nil, err
	}
	return &WalletGuestTxFailedIterator{contract: _WalletGuest.contract, event: "TxFailed", logs: logs, sub: sub}, nil
}

// WatchTxFailed is a free log subscription operation binding the contract event 0xab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b419.
//
// Solidity: event TxFailed(bytes32 indexed _tx, uint256 _index, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) WatchTxFailed(opts *bind.WatchOpts, sink chan<- *WalletGuestTxFailed, _tx [][32]byte) (event.Subscription, error) {

	var _txRule []interface{}
	for _, _txItem := range _tx {
		_txRule = append(_txRule, _txItem)
	}

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "TxFailed", _txRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletGuestTxFailed)
				if err := _WalletGuest.contract.UnpackLog(event, "TxFailed", log); err != nil {
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

// ParseTxFailed is a log parse operation binding the contract event 0xab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b419.
//
// Solidity: event TxFailed(bytes32 indexed _tx, uint256 _index, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) ParseTxFailed(log types.Log) (*WalletGuestTxFailed, error) {
	event := new(WalletGuestTxFailed)
	if err := _WalletGuest.contract.UnpackLog(event, "TxFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
