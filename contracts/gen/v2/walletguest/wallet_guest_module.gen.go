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
	Bin: "0x60a060405234801561001057600080fd5b50306080526080516121de61002d600039600050506121de6000f3fe6080604052600436106100bc5760003560e01c806361c2926c116100745780638c3f55631161004e5780638c3f55631461025357806390042baf14610273578063affed0e0146102ab57600080fd5b806361c2926c146101cb5780637a9a1628146101eb578063853c50681461020b57600080fd5b806320c13b0b116100a557806320c13b0b14610147578063295614261461016757806357c56d6b1461018957600080fd5b806301ffc9a7146100c15780631626ba7e146100f6575b600080fd5b3480156100cd57600080fd5b506100e16100dc366004611880565b6102c0565b60405190151581526020015b60405180910390f35b34801561010257600080fd5b506101166101113660046118e6565b6102d1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016100ed565b34801561015357600080fd5b50610116610162366004611932565b61031e565b34801561017357600080fd5b5061018761018236600461199e565b610383565b005b34801561019557600080fd5b506101bd7f8713a7c4465f6fbee2b6e9d6646d1d9f83fec929edfc4baf661f3c865bdd04d181565b6040519081526020016100ed565b3480156101d757600080fd5b506101876101e63660046119fc565b6103d5565b3480156101f757600080fd5b50610187610206366004611a3e565b61041a565b34801561021757600080fd5b5061022b6102263660046118e6565b610447565b604080519586526020860194909452928401919091526060830152608082015260a0016100ed565b34801561025f57600080fd5b506101bd61026e36600461199e565b61060f565b610286610281366004611ae7565b61063b565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ed565b3480156102b757600080fd5b506101bd610725565b60006102cb82610736565b92915050565b6000806102df858585610792565b509050801561031157507f1626ba7e000000000000000000000000000000000000000000000000000000009050610317565b50600090505b9392505050565b6000806103438686604051610334929190611bb6565b60405180910390208585610792565b509050801561037557507f20c13b0b00000000000000000000000000000000000000000000000000000000905061037b565b50600090505b949350505050565b3330146103c9576040517fe12588940000000000000000000000000000000000000000000000000000000081523360048201523060248201526044015b60405180910390fd5b6103d2816107ca565b50565b600061040883836040516020016103ed929190611d97565b604051602081830303815290604052805190602001206107fc565b9050610415818484610881565b505050565b600061043286866040516020016103ed929190611ddf565b905061043f818787610881565b505050505050565b6000806000806000808787600081811061046357610463611e27565b909101357fff000000000000000000000000000000000000000000000000000000000000001691508190506104b95761049b896107fc565b92506104a8838989610a0e565b929850909650945091506106049050565b7fff00000000000000000000000000000000000000000000000000000000000000818116016104f8576104eb896107fc565b92506104a8838989610a5f565b7ffe000000000000000000000000000000000000000000000000000000000000007fff0000000000000000000000000000000000000000000000000000000000000082160161054a576104eb89610a8b565b7ffd000000000000000000000000000000000000000000000000000000000000007fff000000000000000000000000000000000000000000000000000000000000008216016105ae5761059e898989610af8565b9550955095509550955050610604565b6040517f6085cd820000000000000000000000000000000000000000000000000000000081527fff00000000000000000000000000000000000000000000000000000000000000821660048201526024016103c0565b939792965093509350565b60006102cb7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c75565b600033301461067e576040517fe12588940000000000000000000000000000000000000000000000000000000081523360048201523060248201526044016103c0565b81516020830134f0905073ffffffffffffffffffffffffffffffffffffffff81166106d757816040517f0d2571910000000000000000000000000000000000000000000000000000000081526004016103c09190611eba565b60405173ffffffffffffffffffffffffffffffffffffffff821681527fa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c9060200160405180910390a1919050565b6000610731600061060f565b905090565b60007f6ffbd451000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083160161078957506001919050565b6102cb82610cd3565b60008060008060006107a5888888610447565b509650919450925090508282108015906107bd575060015b9450505050935093915050565b6040517fa038794000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f190100000000000000000000000000000000000000000000000000000000000060208201524660228201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003060601b166042820152605681018290526000906076015b604051602081830303815290604052805190602001209050919050565b8060005b81811015610a0757368484838181106108a0576108a0611e27565b90506020028101906108b29190611ecd565b90506108c16020820182611f0b565b156108fb576040517f230d1ccc000000000000000000000000000000000000000000000000000000008152600481018390526024016103c0565b6040810135805a101561094e5782815a6040517f2bb3e3ba0000000000000000000000000000000000000000000000000000000081526004810193909352602483019190915260448201526064016103c0565b60006109886109636080850160608601611f26565b608085013584156109745784610976565b5a5b61098360a0880188611f41565b610d2f565b905080156109cf57877f5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae7856040516109c291815260200190565b60405180910390a26109f1565b6109f16109e26040850160208601611f0b565b89866109ec610d4c565b610d6b565b50505080806109ff90611fd5565b915050610885565b5050505050565b6000808080610a2987610a24876006818b61200d565b610db9565b6000908152873560f01c6020818152604080842084526002909a013560e01c908190529890912090999198509695509350505050565b6000808080610a7a87610a75876001818b61200d565b610a0e565b935093509350935093509350935093565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526000602282018190527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003060601b1660428301526056820183905290607601610864565b6000808080806004600188013560e81c82610b138383612037565b9050610b258b61022683868d8f61200d565b939b5091995097509550935087871015610b7d57610b4581848b8d61200d565b89896040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016103c0949392919061204a565b8092505b88831015610c675760038301928a013560e81c9150610ba08383612037565b90506000610bc2610bb08861124f565b8c8c879086926102269392919061200d565b939c50919a5098509091505088881015610c1a57610be282858c8e61200d565b8a8a6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016103c0949392919061204a565b848110610c5d576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101829052602481018690526044016103c0565b9350915081610b81565b505050939792965093509350565b6000808383604051602001610c94929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b60007fe4a77bbc000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831601610d2657506001919050565b6102cb82611283565b6000604051828482376000808483898b8af1979650505050505050565b60603d604051915060208201818101604052818352816000823e505090565b8315610d7957805160208201fd5b827fab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b4198383604051610dab929190612071565b60405180910390a250505050565b60008060005b8381101561124657600181019085013560f81c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101610e6057601582019186013560f881901c9060581c73ffffffffffffffffffffffffffffffffffffffff81169074ff000000000000000000000000000000000000000016811785610e465780610e55565b60008681526020829052604090205b955050505050610dbf565b80610ef65760018201918681013560f81c906043016000610e8c8a610e8784888c8e61200d565b61136d565b60ff841697909701969194508491905060a083901b74ff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff82161786610edb5780610eea565b60008781526020829052604090205b96505050505050610dbf565b6002810361101e576000808784013560f881901c9060581c73ffffffffffffffffffffffffffffffffffffffff16601586019550909250905060008885013560e81c600386018162ffffff169150809650819250505060008186019050610f6f8b848c8c8a908692610f6a9392919061200d565b611630565b610fb7578a83610f8183898d8f61200d565b6040517f9a9462320000000000000000000000000000000000000000000000000000000081526004016103c0949392919061208a565b60ff8416979097019694508460a084901b74ff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841617876110025780611011565b60008881526020829052604090205b9750505050505050610dbf565b60038103611051576020820191860135836110395780611048565b60008481526020829052604090205b93505050610dbf565b6004810361109d576003808301928781013560e81c919082010160008061107e8b610a2485898d8f61200d565b60009889526020526040909720969097019650909350610dbf92505050565b600681036111a55760008287013560f81c60018401935060ff16905060008784013560f01c60028501945061ffff16905060008885013560e81c600386018162ffffff16915080965081925050506000818601905060008061110b8d8d8d8b908792610a249392919061200d565b9398508893909250905084821061112157988501985b604080517f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000602080830191909152603882018490526058820188905260788083018a90528351808403909101815260989092019092528051910120896111875780611196565b60008a81526020829052604090205b99505050505050505050610dbf565b600581036112115760208201918601358781036111e0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff94505b60006111eb82611817565b9050846111f85780611207565b60008581526020829052604090205b9450505050610dbf565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016103c0565b50935093915050565b7f8713a7c4465f6fbee2b6e9d6646d1d9f83fec929edfc4baf661f3c865bdd04d160009081526020829052604081206102cb565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fac6a444e00000000000000000000000000000000000000000000000000000000148061131657507fffffffff0000000000000000000000000000000000000000000000000000000082167f36e7817500000000000000000000000000000000000000000000000000000000145b1561132357506001919050565b7f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316146102cb565b6000604282146113ad5782826040517f2ee17a3d0000000000000000000000000000000000000000000000000000000081526004016103c09291906120ca565b60006113c66113bd6001856120de565b85013560f81c90565b60ff169050604084013560f81c843560208601357f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a081111561143a578686826040517fad4aac760000000000000000000000000000000000000000000000000000000081526004016103c0939291906120f1565b8260ff16601b1415801561145257508260ff16601c14155b1561148f578686846040517fe578897e0000000000000000000000000000000000000000000000000000000081526004016103c093929190612115565b600184036114fc576040805160008152602081018083528a905260ff851691810191909152606081018390526080810182905260019060a0015b6020604051602081039080840390855afa1580156114eb573d6000803e3d6000fd5b5050506020604051035194506115d4565b60028403611599576040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101899052600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff861690820152606081018490526080810183905260a0016114c9565b86868560016040517f9dfba8520000000000000000000000000000000000000000000000000000000081526004016103c0949392919061213c565b73ffffffffffffffffffffffffffffffffffffffff85166116255786866040517f6c1719d20000000000000000000000000000000000000000000000000000000081526004016103c09291906120ca565b505050509392505050565b600081810361166b576040517fac241e1100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838361167a6001826120de565b81811061168957611689611e27565b919091013560f81c91505060018114806116a35750600281145b156116e8578473ffffffffffffffffffffffffffffffffffffffff166116ca87868661136d565b73ffffffffffffffffffffffffffffffffffffffff1614915061180e565b600381036117d35773ffffffffffffffffffffffffffffffffffffffff8516631626ba7e878660008761171c6001826120de565b926117299392919061200d565b6040518463ffffffff1660e01b815260040161174793929190612168565b602060405180830381865afa158015611764573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611788919061218b565b7fffffffff00000000000000000000000000000000000000000000000000000000167f1626ba7e0000000000000000000000000000000000000000000000000000000014915061180e565b83838260006040517f9dfba8520000000000000000000000000000000000000000000000000000000081526004016103c0949392919061213c565b50949350505050565b6040517f53657175656e636520737461746963206469676573743a0a0000000000000000602082015260388101829052600090605801610864565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146103d257600080fd5b60006020828403121561189257600080fd5b813561031781611852565b60008083601f8401126118af57600080fd5b50813567ffffffffffffffff8111156118c757600080fd5b6020830191508360208285010111156118df57600080fd5b9250929050565b6000806000604084860312156118fb57600080fd5b83359250602084013567ffffffffffffffff81111561191957600080fd5b6119258682870161189d565b9497909650939450505050565b6000806000806040858703121561194857600080fd5b843567ffffffffffffffff8082111561196057600080fd5b61196c8883890161189d565b9096509450602087013591508082111561198557600080fd5b506119928782880161189d565b95989497509550505050565b6000602082840312156119b057600080fd5b5035919050565b60008083601f8401126119c957600080fd5b50813567ffffffffffffffff8111156119e157600080fd5b6020830191508360208260051b85010111156118df57600080fd5b60008060208385031215611a0f57600080fd5b823567ffffffffffffffff811115611a2657600080fd5b611a32858286016119b7565b90969095509350505050565b600080600080600060608688031215611a5657600080fd5b853567ffffffffffffffff80821115611a6e57600080fd5b611a7a89838a016119b7565b9097509550602088013594506040880135915080821115611a9a57600080fd5b50611aa78882890161189d565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215611af957600080fd5b813567ffffffffffffffff80821115611b1157600080fd5b818401915084601f830112611b2557600080fd5b813581811115611b3757611b37611ab8565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611b7d57611b7d611ab8565b81604052828152876020848701011115611b9657600080fd5b826020860160208301376000928101602001929092525095945050505050565b8183823760009101908152919050565b80358015158114611bd657600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611bd657600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b81835260006020808501808196508560051b810191508460005b87811015611d8a57828403895281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41883603018112611ca157600080fd5b870160c0611cae82611bc6565b15158652611cbd878301611bc6565b15158688015260408281013590870152606073ffffffffffffffffffffffffffffffffffffffff611cef828501611bdb565b16908701526080828101359087015260a080830135368490037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1018112611d3557600080fd5b90920187810192903567ffffffffffffffff811115611d5357600080fd5b803603841315611d6257600080fd5b8282890152611d748389018286611bff565b9c89019c97505050928601925050600101611c62565b5091979650505050505050565b60408152600560408201527f73656c663a000000000000000000000000000000000000000000000000000000606082015260806020820152600061037b608083018486611c48565b60408152600660408201527f67756573743a0000000000000000000000000000000000000000000000000000606082015260806020820152600061037b608083018486611c48565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000815180845260005b81811015611e7c57602081850181015186830182015201611e60565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006103176020830184611e56565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff41833603018112611f0157600080fd5b9190910192915050565b600060208284031215611f1d57600080fd5b61031782611bc6565b600060208284031215611f3857600080fd5b61031782611bdb565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611f7657600080fd5b83018035915067ffffffffffffffff821115611f9157600080fd5b6020019150368190038213156118df57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361200657612006611fa6565b5060010190565b6000808585111561201d57600080fd5b8386111561202a57600080fd5b5050820193919092039150565b808201808211156102cb576102cb611fa6565b60608152600061205e606083018688611bff565b6020830194909452506040015292915050565b82815260406020820152600061037b6040830184611e56565b84815273ffffffffffffffffffffffffffffffffffffffff841660208201526060604082015260006120c0606083018486611bff565b9695505050505050565b60208152600061037b602083018486611bff565b818103818111156102cb576102cb611fa6565b604081526000612105604083018587611bff565b9050826020830152949350505050565b604081526000612129604083018587611bff565b905060ff83166020830152949350505050565b606081526000612150606083018688611bff565b60208301949094525090151560409091015292915050565b838152604060208201526000612182604083018486611bff565b95945050505050565b60006020828403121561219d57600080fd5b81516103178161185256fea26469706673582212200896636ab1dae9ad33c5080d1044c3c12105a6d4bc196fd0009bf12ed3b0f85364736f6c63430008120033",
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"InvalidSValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_v\",\"type\":\"uint256\"}],\"name\":\"InvalidVValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_available\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDelegatecall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"OnlySelfAuth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"SignerIsAddress0\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_type\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_recoverMode\",\"type\":\"bool\"}],\"name\":\"UnsupportedSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_prev\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"CreatedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"TxExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"TxFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SET_IMAGE_HASH_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"signatureRecovery\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
