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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"CreatedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"}],\"name\":\"TxExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_tx\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"TxFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revertOnError\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIModuleCalls.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ddc806100206000396000f3fe60806040526004361061007b5760003560e01c80637a9a16281161004e5780637a9a1628146101255780638c3f55631461014557806390042baf14610172578063affed0e0146101925761007b565b806301ffc9a7146100805780631626ba7e146100b657806320c13b0b146100e357806361c2926c14610103575b600080fd5b34801561008c57600080fd5b506100a061009b366004611677565b6101a7565b6040516100ad91906118be565b60405180910390f35b3480156100c257600080fd5b506100d66100d136600461162d565b6101ba565b6040516100ad91906118eb565b3480156100ef57600080fd5b506100d66100fe3660046116b7565b610233565b34801561010f57600080fd5b5061012361011e366004611590565b61028d565b005b34801561013157600080fd5b506101236101403660046115c3565b6102ce565b34801561015157600080fd5b50610165610160366004611753565b6102f6565b6040516100ad91906118c9565b610185610180366004611720565b610322565b6040516100ad919061189d565b34801561019e57600080fd5b506101656103d6565b60006101b2826103e7565b90505b919050565b60006102046101c885610444565b84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104a492505050565b1561022c57507f1626ba7e000000000000000000000000000000000000000000000000000000005b9392505050565b600061025d6101c88686604051808383808284376040519201829003909120935061044492505050565b1561028557507f20c13b0b000000000000000000000000000000000000000000000000000000005b949350505050565b60006102be826040516020016102a39190611a19565b60405160208183030381529060405280519060200120610444565b90506102ca818361069c565b5050565b60006102e4846040516020016102a39190611975565b90506102f0818561069c565b50505050565b60006101b27f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610817565b600033301461037c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611d806027913960400191505060405180910390fd5b81516020830134f06040805173ffffffffffffffffffffffffffffffffffffffff8316815290519192507fa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c919081900360200190a1919050565b60006103e260006102f6565b905090565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f90042baf00000000000000000000000000000000000000000000000000000000141561043b575060016101b5565b6101b282610844565b604080517f19010000000000000000000000000000000000000000000000000000000000006020808301919091524660228301523060601b6042830152605680830194909452825180830390940184526076909101909152815191012090565b60008060006104b2846108a1565b909250905061ffff821660005b855183101561067957600080806104d6898761090f565b975060ff918216945016915060018314156104fe576104f58987610990565b96509050610622565b8261052a57606061050f8a88610a08565b9750905061051d8b82610ab9565b9150828501945050610622565b60028314156105d15761053d8987610990565b96509050600061054d8a88610e43565b975061ffff16905060606105628b8984610eb4565b985090506105718c8483610fa3565b6105c6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526032815260200180611c0b6032913960400191505060405180910390fd5b505092810192610622565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c815260200180611b28602c913960400191505060405180910390fd5b848282604051602001808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff16815260200193505050506040516020818303038152906040528051906020012094505050506104bf565b8361ffff1681101580156106915750610691826111eb565b979650505050505050565b60005b81518110156108125760008282815181106106b657fe5b6020026020010151905060006060826000015115610709576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610700906119bc565b60405180910390fd5b82604001515a1015610747576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070090611918565b826060015173ffffffffffffffffffffffffffffffffffffffff168360800151846040015160001461077d57846040015161077f565b5a5b908560a001516040516107929190611881565b600060405180830381858888f193505050503d80600081146107d0576040519150601f19603f3d011682016040523d82523d6000602084013e6107d5565b606091505b50909250905081156107fc57856040516107ef91906118c9565b60405180910390a0610807565b6108078387836111f1565b50505060010161069f565b505050565b60408051602080820194909452808201929092528051808303820181526060909201905280519101205490565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f389901c7000000000000000000000000000000000000000000000000000000001415610898575060016101b5565b6101b282611241565b6020810151815160f09190911c9060029081111561090a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180611b776027913960400191505060405180910390fd5b915091565b8082016020015160f881901c9060f01c60ff166002830183811161092f57fe5b8451811115610989576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611cdb6026913960400191505060405180910390fd5b9250925092565b8082016020015160601c601482018281116109a757fe5b8351811115610a01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611b546023913960400191505060405180910390fd5b9250929050565b604080516042808252608082019092526060916000919060208201818036833701905050915082840160200180516020840152602081015160408401526022810151604284015250604283019050828111610a5f57fe5b8351811115610a01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611c7c6023913960400191505060405180910390fd5b60008151604214610b15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611aee603a913960400191505060405180910390fd5b600082600184510381518110610b2757fe5b602001015160f81c60f81b60f81c60ff169050600083604081518110610b4957fe5b016020015160f81c90506000610b5f85826112c9565b90506000610b6e8660206112c9565b90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0811115610be9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603d815260200180611ab1603d913960400191505060405180910390fd5b8260ff16601b14158015610c0157508260ff16601c14155b15610c57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603d815260200180611b9e603d913960400191505060405180910390fd5b6001841415610ccb5760018784848460405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610cba573d6000803e3d6000fd5b505050602060405103519450610dcd565b6002841415610d7c5760018760405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250601c018281526020019150506040516020818303038152906040528051906020012084848460405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610cba573d6000803e3d6000fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c815260200180611c9f603c913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8516610e39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180611bdb6030913960400191505060405180910390fd5b5050505092915050565b8082016020015160f01c60028201828111610e5a57fe5b8351811115610a01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180611d226022913960400191505060405180910390fd5b606060008267ffffffffffffffff81118015610ecf57600080fd5b506040519080825280601f01601f191660200182016040528015610efa576020820181803683370190505b509150838501602001600060205b85811015610f2157908201518482015260208101610f08565b8486016020018051939092015190850152525082820183811015610f4157fe5b8451811115610f9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611d016021913960400191505060405180910390fd5b935093915050565b60008082600184510381518110610fb657fe5b016020015160f81c90506001811480610fcf5750600281145b15611013578373ffffffffffffffffffffffffffffffffffffffff16610ff58685610ab9565b73ffffffffffffffffffffffffffffffffffffffff161491506111e3565b60038114156111925782517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81018452604080517f1626ba7e000000000000000000000000000000000000000000000000000000008152600481018881526024820192835286516044830152865173ffffffffffffffffffffffffffffffffffffffff891693631626ba7e938b938a9390929160640190602085019080838360005b838110156110cd5781810151838201526020016110b5565b50505050905090810190601f1680156110fa5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15801561111857600080fd5b505afa15801561112c573d6000803e3d6000fd5b505050506040513d602081101561114257600080fd5b50519084527fffffffff00000000000000000000000000000000000000000000000000000000167f1626ba7e000000000000000000000000000000000000000000000000000000001491506111e3565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603f815260200180611c3d603f913960400191505060405180910390fd5b509392505050565b50600190565b82602001511561120357805160208201fd5b7f3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd782826040516112349291906118d2565b60405180910390a1505050565b60007fffffffff00000000000000000000000000000000000000000000000000000000821615806112b357507fffffffff0000000000000000000000000000000000000000000000000000000082167f36e7817500000000000000000000000000000000000000000000000000000000145b156112c0575060016101b5565b6101b282611331565b60008160200183511015611328576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c815260200180611d44603c913960400191505060405180910390fd5b50016020015190565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f01ffc9a70000000000000000000000000000000000000000000000000000000014919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101b557600080fd5b600082601f8301126113af578081fd5b8135602067ffffffffffffffff808311156113c657fe5b6113d38283850201611a60565b83815282810190868401865b868110156114af578135890160c0807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0838e0301121561141d57898afd5b604080518281018181108a8211171561143257fe5b825261143f848b016114bd565b815261144c8285016114bd565b8a820152606080850135838301526080925061146983860161137b565b9082015260a08481013583830152928401359289841115611488578c8dfd5b6114968f8c8688010161150d565b90820152875250505092850192908501906001016113df565b509098975050505050505050565b803580151581146101b557600080fd5b60008083601f8401126114de578182fd5b50813567ffffffffffffffff8111156114f5578182fd5b602083019150836020828501011115610a0157600080fd5b600082601f83011261151d578081fd5b813567ffffffffffffffff81111561153157fe5b61156260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611a60565b818152846020838601011115611576578283fd5b816020850160208301379081016020019190915292915050565b6000602082840312156115a1578081fd5b813567ffffffffffffffff8111156115b7578182fd5b6102858482850161139f565b6000806000606084860312156115d7578182fd5b833567ffffffffffffffff808211156115ee578384fd5b6115fa8783880161139f565b9450602086013593506040860135915080821115611616578283fd5b506116238682870161150d565b9150509250925092565b600080600060408486031215611641578283fd5b83359250602084013567ffffffffffffffff81111561165e578283fd5b61166a868287016114cd565b9497909650939450505050565b600060208284031215611688578081fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461022c578182fd5b600080600080604085870312156116cc578081fd5b843567ffffffffffffffff808211156116e3578283fd5b6116ef888389016114cd565b90965094506020870135915080821115611707578283fd5b50611714878288016114cd565b95989497509550505050565b600060208284031215611731578081fd5b813567ffffffffffffffff811115611747578182fd5b6102858482850161150d565b600060208284031215611764578081fd5b5035919050565b60008282518085526020808601955080818302840101818601855b8481101561182a578583037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00189528151805115158452848101511515858501526040808201519085015260608082015173ffffffffffffffffffffffffffffffffffffffff16908501526080808201519085015260a09081015160c09185018290529061181681860183611837565b9a86019a9450505090830190600101611786565b5090979650505050505050565b6000815180845261184f816020860160208601611a84565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60008251611893818460208701611a84565b9190910192915050565b73ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b901515815260200190565b90815260200190565b6000838252604060208301526102856040830184611837565b7fffffffff0000000000000000000000000000000000000000000000000000000091909116815260200190565b60208082526029908201527f47756573744d6f64756c65235f6578656375746547756573743a204e4f545f4560408201527f4e4f5547485f4741530000000000000000000000000000000000000000000000606082015260800190565b600060408252600660408301527f67756573743a000000000000000000000000000000000000000000000000000060608301526080602083015261022c608083018461176b565b60208082526033908201527f47756573744d6f64756c65235f6578656375746547756573743a2064656c656760408201527f61746543616c6c206e6f7420616c6c6f77656400000000000000000000000000606082015260800190565b600060408252600560408301527f73656c663a00000000000000000000000000000000000000000000000000000060608301526080602083015261022c608083018461176b565b60405181810167ffffffffffffffff81118282101715611a7c57fe5b604052919050565b60005b83811015611a9f578181015183820152602001611a87565b838111156102f0575050600091015256fe5369676e617475726556616c696461746f72237265636f7665725369676e65723a20696e76616c6964207369676e6174757265202773272076616c75655369676e617475726556616c696461746f72237265636f7665725369676e65723a20696e76616c6964207369676e6174757265206c656e6774684d6f64756c6541757468235f7369676e617475726556616c69646174696f6e20494e56414c49445f464c41474c696242797465732372656164416464726573733a204f55545f4f465f424f554e44534c696242797465732372656164466972737455696e7431363a204f55545f4f465f424f554e44535369676e617475726556616c696461746f72237265636f7665725369676e65723a20696e76616c6964207369676e6174757265202776272076616c75655369676e617475726556616c696461746f72237265636f7665725369676e65723a20494e56414c49445f5349474e45524d6f64756c6541757468235f7369676e617475726556616c69646174696f6e3a20494e56414c49445f5349474e41545552455369676e617475726556616c696461746f7223697356616c69645369676e61747572653a20554e535550504f525445445f5349474e41545552455f545950454c696242797465732372656164427974657336363a204f55545f4f465f424f554e44535369676e617475726556616c696461746f72237265636f7665725369676e65723a20554e535550504f525445445f5349474e41545552455f545950454c69624279746573237265616455696e743855696e74383a204f55545f4f465f424f554e44534c69624279746573237265616442797465733a204f55545f4f465f424f554e44534c69624279746573237265616455696e7431363a204f55545f4f465f424f554e44534c696242797465732372656164427974657333323a20475245415445525f4f525f455155414c5f544f5f33325f4c454e4754485f52455155495245444d6f64756c6553656c6641757468236f6e6c7953656c663a204e4f545f415554484f52495a4544a2646970667358221220f5a1de0b650baa2ee828e8766bc6dbd0c74da0cc4735a143852d24f868e4b62464736f6c63430007060033",
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
	Reason []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxFailed is a free log retrieval operation binding the contract event 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7.
//
// Solidity: event TxFailed(bytes32 _tx, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) FilterTxFailed(opts *bind.FilterOpts) (*WalletGuestTxFailedIterator, error) {

	logs, sub, err := _WalletGuest.contract.FilterLogs(opts, "TxFailed")
	if err != nil {
		return nil, err
	}
	return &WalletGuestTxFailedIterator{contract: _WalletGuest.contract, event: "TxFailed", logs: logs, sub: sub}, nil
}

// WatchTxFailed is a free log subscription operation binding the contract event 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7.
//
// Solidity: event TxFailed(bytes32 _tx, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) WatchTxFailed(opts *bind.WatchOpts, sink chan<- *WalletGuestTxFailed) (event.Subscription, error) {

	logs, sub, err := _WalletGuest.contract.WatchLogs(opts, "TxFailed")
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

// ParseTxFailed is a log parse operation binding the contract event 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7.
//
// Solidity: event TxFailed(bytes32 _tx, bytes _reason)
func (_WalletGuest *WalletGuestFilterer) ParseTxFailed(log types.Log) (*WalletGuestTxFailed, error) {
	event := new(WalletGuestTxFailed)
	if err := _WalletGuest.contract.UnpackLog(event, "TxFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
