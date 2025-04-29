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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506117588061001c5f395ff3fe608060405234801561000f575f5ffd5b505f61001b5f36610035565b90505f610027826104d7565b90506100338282610527565b005b61003d610d03565b5f815f019060ff16908160ff16815250505f5f61005a85856107e8565b915060ff16915060018083160361007a575f8360600181815250506100b6565b61008f8186866107fe9290919263ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff169150846060018193508281525050505b5f6007600184901c1690505f8111156100f3576100e58282888861082f9190939291909392919063ffffffff16565b856080018194508281525050505b5f601080851603610107576001905061015f565b60208085160361013a5761012683888861085c9290919263ffffffff16565b8161ffff169150809450819250505061015e565b61014f83888861087b9290919263ffffffff16565b8160ff16915080945081925050505b5b8067ffffffffffffffff81111561017957610178610d9f565b5b6040519080825280602002602001820160405280156101b257816020015b61019f610d4e565b8152602001906001900390816101975790505b5085604001819052505f5f90505b818110156104cc575f6101de858a8a61087b9290919263ffffffff16565b8096508192505050600180821660ff160361024d57308760400151838151811061020b5761020a610dcc565b5b60200260200101515f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506102b9565b610262858a8a6108969290919263ffffffff16565b8860400151848151811061027957610278610dcc565b5b60200260200101515f018197508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b600280821660ff1603610307576102db858a8a6108c79290919263ffffffff16565b886040015184815181106102f2576102f1610dcc565b5b60200260200101516020018197508281525050505b600480821660ff16036103cf575f61032a868b8b6108dd9290919263ffffffff16565b8162ffffff169150809750819250505089898790838961034a9190610e2f565b9261035793929190610e6a565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050886040015184815181106103b0576103af610dcc565b5b60200260200101516040018190525080866103cb9190610e2f565b9550505b600880821660ff160361041d576103f1858a8a6108c79290919263ffffffff16565b8860400151848151811061040857610407610dcc565b5b60200260200101516060018197508281525050505b601080821660ff16148760400151838151811061043d5761043c610dcc565b5b60200260200101516080019015159081151581525050602080821660ff16148760400151838151811061047357610472610dcc565b5b602002602001015160a0019015159081151581525050600660c0821660ff16901c60ff16876040015183815181106104ae576104ad610dcc565b5b602002602001015160c00181815250505080806001019150506101c0565b505050505092915050565b5f5f6104e78360200151306108fd565b90505f6104f3846109a1565b90508181604051602001610508929190610f21565b6040516020818303038152906040528051906020012092505050919050565b5f5f90505f83604001515190505f5f90505b818110156107e1575f8560400151828151811061055957610558610dcc565b5b602002602001015190508060a001518015610572575083155b156105b6577f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b85836040516105a8929190610f75565b60405180910390a1506107d4565b5f93505f816060015190505f81141580156105d05750805a105b156106165786835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161060d9392919061136a565b60405180910390fd5b81608001511561065d57826040517f230d1ccc00000000000000000000000000000000000000000000000000000000815260040161065491906113a6565b60405180910390fd5b5f610683835f015184602001515f85146106775784610679565b5a5b8660400151610bd5565b905080610797575f60ff168360c00151036106e657600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d87856106c7610bec565b6040516106d693929190611407565b60405180910390a15050506107d4565b600160ff168360c001510361073d5787846106ff610bec565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161073493929190611443565b60405180910390fd5b600260ff168360c0015103610796577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b8785610777610bec565b60405161078693929190611407565b60405180910390a15050506107e1565b5b7f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a87856040516107c8929190610f75565b60405180910390a15050505b8080600101915050610539565b5050505050565b5f5f83358060f81c925060019150509250929050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f858401356008840261010003600180866008021b0382821c1693508486019250505094509492505050565b5f5f8483013561ffff8160f01c16925060028401915050935093915050565b5f5f848301358060f81c925060018401915050935093915050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f848301359150602083019050935093915050565b5f5f8483013562ffffff8160e81c16925060038401915050935093915050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de8561096c574661096e565b5f5b85604051602001610983959493929190611495565b60405160208183030381529060405280519060200120905092915050565b5f5f8261010001516040516020016109b99190611572565b6040516020818303038152906040528051906020012090505f60ff16835f015160ff1603610a51575f6109ef8460400151610c0a565b90507f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a2818560600151866080015185604051602001610a32959493929190611588565b6040516020818303038152906040528051906020012092505050610bd0565b600160ff16835f015160ff1603610ac0577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360a001518051906020012082604051602001610aa2939291906115d9565b60405160208183030381529060405280519060200120915050610bd0565b600260ff16835f015160ff1603610b28577f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e48360c0015182604051602001610b0a939291906115d9565b60405160208183030381529060405280519060200120915050610bd0565b600360ff16835f015160ff1603610b90577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360e0015182604051602001610b72939291906115d9565b60405160208183030381529060405280519060200120915050610bd0565b825f01516040517f04818320000000000000000000000000000000000000000000000000000000008152600401610bc7919061161d565b60405180910390fd5b919050565b5f5f5f835160208501878988f19050949350505050565b60603d604051915060208201818101604052818352815f823e505090565b5f60605f5f90505b8351811015610c73575f610c3f858381518110610c3257610c31610dcc565b5b6020026020010151610c84565b90508281604051602001610c54929190611670565b6040516020818303038152906040529250508080600101915050610c12565b508080519060200120915050919050565b5f7f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437825f01518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001610ce69897969594939291906116a6565b604051602081830303815290604052805190602001209050919050565b6040518061012001604052805f60ff1681526020015f15158152602001606081526020015f81526020015f8152602001606081526020015f81526020015f8152602001606081525090565b6040518060e001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020015f81526020015f151581526020015f151581526020015f81525090565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e3982610df9565b9150610e4483610df9565b9250828201905080821115610e5c57610e5b610e02565b5b92915050565b5f5ffd5b5f5ffd5b5f5f85851115610e7d57610e7c610e62565b5b83861115610e8e57610e8d610e66565b5b6001850283019150848603905094509492505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f610ee2600283610ea4565b9150610eed82610eae565b600282019050919050565b5f819050919050565b5f819050919050565b610f1b610f1682610ef8565b610f01565b82525050565b5f610f2b82610ed6565b9150610f378285610f0a565b602082019150610f478284610f0a565b6020820191508190509392505050565b610f6081610ef8565b82525050565b610f6f81610df9565b82525050565b5f604082019050610f885f830185610f57565b610f956020830184610f66565b9392505050565b5f60ff82169050919050565b610fb181610f9c565b82525050565b5f8115159050919050565b610fcb81610fb7565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61102382610ffa565b9050919050565b61103381611019565b82525050565b61104281610df9565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61108a82611048565b6110948185611052565b93506110a4818560208601611062565b6110ad81611070565b840191505092915050565b5f60e083015f8301516110cd5f86018261102a565b5060208301516110e06020860182611039565b50604083015184820360408601526110f88282611080565b915050606083015161110d6060860182611039565b5060808301516111206080860182610fc2565b5060a083015161113360a0860182610fc2565b5060c083015161114660c0860182611039565b508091505092915050565b5f61115c83836110b8565b905092915050565b5f602082019050919050565b5f61117a82610fd1565b6111848185610fdb565b93508360208202850161119685610feb565b805f5b858110156111d157848403895281516111b28582611151565b94506111bd83611164565b925060208a01995050600181019050611199565b50829750879550505050505092915050565b6111ec81610ef8565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f611226838361102a565b60208301905092915050565b5f602082019050919050565b5f611248826111f2565b61125281856111fc565b935061125d8361120c565b805f5b8381101561128d578151611274888261121b565b975061127f83611232565b925050600181019050611260565b5085935050505092915050565b5f61012083015f8301516112b05f860182610fa8565b5060208301516112c36020860182610fc2565b50604083015184820360408601526112db8282611170565b91505060608301516112f06060860182611039565b5060808301516113036080860182611039565b5060a083015184820360a086015261131b8282611080565b91505060c083015161133060c08601826111e3565b5060e083015161134360e08601826111e3565b5061010083015184820361010086015261135d828261123e565b9150508091505092915050565b5f6060820190508181035f830152611382818661129a565b90506113916020830185610f66565b61139e6040830184610f66565b949350505050565b5f6020820190506113b95f830184610f66565b92915050565b5f82825260208201905092915050565b5f6113d982611048565b6113e381856113bf565b93506113f3818560208601611062565b6113fc81611070565b840191505092915050565b5f60608201905061141a5f830186610f57565b6114276020830185610f66565b818103604083015261143981846113cf565b9050949350505050565b5f6060820190508181035f83015261145b818661129a565b905061146a6020830185610f66565b818103604083015261147c81846113cf565b9050949350505050565b61148f81611019565b82525050565b5f60a0820190506114a85f830188610f57565b6114b56020830187610f57565b6114c26040830186610f57565b6114cf6060830185610f66565b6114dc6080830184611486565b9695505050505050565b5f81905092915050565b6114f981611019565b82525050565b5f61150a83836114f0565b60208301905092915050565b5f611520826111f2565b61152a81856114e6565b93506115358361120c565b805f5b8381101561156557815161154c88826114ff565b975061155783611232565b925050600181019050611538565b5085935050505092915050565b5f61157d8284611516565b915081905092915050565b5f60a08201905061159b5f830188610f57565b6115a86020830187610f57565b6115b56040830186610f66565b6115c26060830185610f66565b6115cf6080830184610f57565b9695505050505050565b5f6060820190506115ec5f830186610f57565b6115f96020830185610f57565b6116066040830184610f57565b949350505050565b61161781610f9c565b82525050565b5f6020820190506116305f83018461160e565b92915050565b5f81905092915050565b5f61164a82611048565b6116548185611636565b9350611664818560208601611062565b80840191505092915050565b5f61167b8285611640565b91506116878284610f0a565b6020820191508190509392505050565b6116a081610fb7565b82525050565b5f610100820190506116ba5f83018b610f57565b6116c7602083018a611486565b6116d46040830189610f66565b6116e16060830188610f57565b6116ee6080830187610f66565b6116fb60a0830186611697565b61170860c0830185611697565b61171560e0830184610f66565b999850505050505050505056fea2646970667358221220280d359ada9ef72db08ca1154f48a98bf6e0a9a66087c02899dc98c60189caf464736f6c634300081c0033",
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
// Solidity: fallback() returns()
func (_WalletGuest *WalletGuestTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_WalletGuest *WalletGuestSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletGuest.Contract.Fallback(&_WalletGuest.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
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
