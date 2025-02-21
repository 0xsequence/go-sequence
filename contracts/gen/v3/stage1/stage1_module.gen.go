// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletupgradable

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

// PayloadDecoded is an auto generated low-level Go binding around an user-defined struct.
type PayloadDecoded struct {
	Kind          uint8
	NoChainId     bool
	Calls         []PayloadCall
	Space         *big.Int
	Nonce         *big.Int
	Message       []byte
	ImageHash     [32]byte
	Digest        [32]byte
	ParentWallets []common.Address
}

// WalletUpgradableMetaData contains all meta data concerning the WalletUpgradable contract.
var WalletUpgradableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Failed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Skipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Success\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAGE_2_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561001057600080fd5b50604051616df2380380616df283398101604081905261002f916100c7565b8060405161003c906100ba565b604051809103906000f080158015610058573d6000803e3d6000fd5b506000604051806060016040528060288152602001616dca60289139604051610086919030906020016100f7565b60408051601f198184030181529190528051602090910120608052506001600160a01b0391821660a0521660c05250610129565b61359a8061383083390190565b6000602082840312156100d957600080fd5b81516001600160a01b03811681146100f057600080fd5b9392505050565b6000835160005b8181101561011857602081870181015185830152016100fe565b509190910191825250602001919050565b60805160a05160c0516136c361016d60003960008181610214015261112b0152600081816101a2015261159f01526000818161015a01526115d001526136c36000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80636ea445771161008c578063aaf10f4211610066578063aaf10f4214610236578063ad55366b1461023e578063ca70785014610284578063f727ef1c1461029757600080fd5b80636ea44577146101e95780638c3f5563146101fc5780639f69ef541461020f57600080fd5b8063257671f5116100bd578063257671f514610155578063295614261461018a5780632dd310001461019d57600080fd5b8063025b22bc146100e45780631626ba7e146100f95780631f6a1eb914610142575b600080fd5b6100f76100f2366004612aea565b6102aa565b005b61010c610107366004612b4e565b6102f6565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b6100f7610150366004612b9a565b61038e565b61017c7f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001610139565b6100f7610198366004612c0b565b610410565b6101c47f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610139565b6100f76101f7366004612c24565b610454565b61017c61020a366004612c0b565b6104ba565b6101c47f000000000000000000000000000000000000000000000000000000000000000081565b6101c46104ec565b61025161024c366004612fa6565b6104fb565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610139565b61017c610292366004612fa6565b610535565b6100f76102a53660046130d4565b61069e565b3330146102ea576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6102f381610759565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e081018590528161034d8286866107ad565b509050806103615750600091506103879050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061039a85856108fb565b90506103ae81606001518260800151610d19565b6000806103bc8386866107ad565b91509150816103fd578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016102e1939291906133a8565b6104078184610e01565b50505050505050565b33301461044b576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016102e1565b6102f381611093565b33301461048f576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016102e1565b600061049b83836108fb565b905060006104a88261114f565b90506104b48183610e01565b50505050565b60006104e67f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e836111ca565b92915050565b60006104f6305490565b905090565b600080600080600080610512898989600080611228565b9399509197509450925090506105278361154d565b935093975093979195509350565b60008084610100015151600161054b91906133d8565b67ffffffffffffffff81111561056357610563612c66565b60405190808252806020026020018201604052801561058c578160200160208202803683370190505b50905060005b856101000151518110156105fe5785610100015181815181106105b7576105b7613412565b60200260200101518282815181106105d1576105d1613412565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610592565b503381866101000151518151811061061857610618613412565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006106528686866107ad565b50905080610692578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016102e1939291906133a8565b50600195945050505050565b3330146106d9576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016102e1565b6106f28383836bffffffffffffffffffffffff1661164f565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b610761813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039060200160405180910390a150565b6000806000848460008181106107c5576107c5613412565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f800000000000000000000000000000000000000000000000000000000000000090811690036108c4576108228661114f565b9150600080610830846116de565b9150915042811015610878576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016102e1565b73ffffffffffffffffffffffffffffffffffffffff821615806108b0575073ffffffffffffffffffffffffffffffffffffffff821633145b156108c157600194505050506108f3565b50505b60008060006108d7898989600080611228565b985092955090935091506108ec90508161154d565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c6001808216810361095f5760006060840152610970565b84810135606090811c908401526014015b6007600183901c1680156109c35760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b6000836010166010036109d857506001610a01565b836020166020036109f45750600282019186013560f01c610a01565b50600182019186013560f81c5b8067ffffffffffffffff811115610a1a57610a1a612c66565b604051908082528060200260200182016040528015610aa557816020015b610a926040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610a385790505b50604086015260005b81811015610d0e5760018085019489013560f81c908082169003610b0d573087604001518381518110610ae357610ae3613412565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052610b57565b8489013560601c6014860188604001518481518110610b2e57610b2e613412565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003610b9557848901356020860188604001518481518110610b8057610b80613412565b60200260200101516020018197508281525050505b60048082169003610c2d57600385019489013560e81c89868a610bb884836133d8565b92610bc593929190613441565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506040890151805185908110610c1057610c10613412565b602090810291909101015160400152610c2981876133d8565b9550505b60088082169003610c6b57848901356020860188604001518481518110610c5657610c56613412565b60200260200101516060018197508281525050505b8060101660ff1660101487604001518381518110610c8b57610c8b613412565b602002602001015160800190151590811515815250508060201660ff1660201487604001518381518110610cc157610cc1613412565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110610cf657610cf6613412565b602090810291909101015160c0015250600101610aae565b505050505092915050565b6000610d24836104ba565b9050818114610d70576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016102e1565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561108c57600084604001518281518110610e2c57610e2c613412565b602002602001015190508060a001518015610e45575083155b15610e8d576040805187815260208101849052600095507fa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7910160405180910390a150611084565b60608101518015801590610ea05750805a105b15610edd5785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016102e19392919061346b565b6000826080015115610f0357610efc836000015183856040015161172a565b9050610f1e565b610f1b83600001518460200151848660400151611740565b90505b806110475760c0830151610f71576040805189815260208101869052600197507fe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643910160405180910390a1505050611084565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610fdb578684610fa6611758565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016102e193929190613490565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016110475760408051898152602081018690527f6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8910160405180910390a150505061108c565b60408051898152602081018690527f2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3910160405180910390a15050505b600101610e0c565b5050505050565b806110ca576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110f37fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200160405180910390a16102f37f0000000000000000000000000000000000000000000000000000000000000000610759565b600080611160836020015130611777565b9050600061116d84611844565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b60008083836040516020016111e9929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611273575073ffffffffffffffffffffffffffffffffffffffff8916155b1561138f578b82013560601c98506014909101908961138f5760038201918c013560e81c60008d848e6112a685836133d8565b926112b393929190613441565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc8915061133e90309085906004016134bb565b6040805180830381865afa15801561135a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137e91906134f2565b925061138a82856133d8565b935050505b826001166001036113c9576113b78d8a838f8f879080926113b293929190613441565b611a9a565b97509750975097509750505050611540565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c168382019096509250600090506114346001600586901c8116906133d8565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c169950909201915061147f8d61114f565b935061149d8d858e8e8690809261149893929190613441565b611ce6565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d16909152902082519198509650158015906114e8575080518614155b80156114f8575080602001518511155b1561153c576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016102e1565b5050505b9550955095509550959050565b60006104e6826040517fff0000000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060601b166021820152603581018290527f000000000000000000000000000000000000000000000000000000000000000060558201526000903090607501604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012073ffffffffffffffffffffffffffffffffffffffff161492915050565b6116d97fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b6000808061170c7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86856111ca565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856117e757466117ea565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000806118558361010001516126ca565b835190915060ff166118c8576000611870846040015161274d565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016111ab565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016119585760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016119c85760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e46020820152908101919091526060810182905260800161193a565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01611a385760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e26246020820152908101919091526060810182905260800161193a565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016102e1565b6000806000806000611afc604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b88821015611c735760038201916000908b013560e81c611b4484826133d8565b9150600090508a8214611b58576000611b5a565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8414611b8b5785611b8d565b8f5b9050611bad818e8e88908792611ba593929190613441565b600186611228565b50929d50909b50995097508a8a1015611c0a578c8c86908592611bd293929190613441565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016102e19493929190613543565b829450888e6000015103611c1d5760008e525b838810611c60576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016102e1565b505060c084018790529150849050611b24565b8a5115801590611c8757508a602001518511155b15611ccb576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016102e1565b611cd48d61114f565b93505050509550955095509550959050565b60008060005b838110156126c057600181019085013560f881901c9060fc1c80611e0957600f82166000819003611d245750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015611db9573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000611ddc828960ff166127cf565b90508b611de95780611df8565b60008c81526020829052604090205b9b5050505050505050505050611cec565b60018103611e6d57600f82166000819003611e2b5750600183019287013560f81c5b601484019388013560601c6000611e458260ff85166127cf565b905086611e525780611e61565b60008781526020829052604090205b96505050505050611cec565b6002810361205457600382166000819003611e8f5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d908792611f4d93929190613441565b6040518463ffffffff1660e01b8152600401611f6b9392919061356a565b602060405180830381865afa158015611f88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fac9190613584565b7fffffffff00000000000000000000000000000000000000000000000000000000161461200f578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016102e19594939291906135c6565b8097508460ff168a0199506000612029858760ff166127cf565b9050896120365780612045565b60008a81526020829052604090205b99505050505050505050611cec565b600381036120885760208301928701358461206f578061207e565b60008581526020829052604090205b9450505050611cec565b6004810361212e57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806120fd8e8e8e8e8c90889261149893929190613441565b91509150829750818a01995061211d898260009182526020526040902090565b985082975050505050505050611cec565b600681036121f8576003600283901c1660008190036121545750600183019287013560f81c5b60038316600081900361216e5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806121ac8f8f8f8f8d90889261149893929190613441565b915091508298508482106121bf57998501995b60006121cc828789612836565b90508a6121d957806121e8565b60008b81526020829052604090205b9a50505050505050505050611cec565b60058103612265576020830192870135888103612233577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061223e8261289a565b90508561224b578061225a565b60008681526020829052604090205b955050505050611cec565b6007810361236b57600f821660008190036122875750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001611d97565b600881036123bf57602083019287013560006123878b826128ee565b90508082036123b4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b6000611e4582612969565b60098103612525576003821660008190036123e15750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c90879261247b93929190613441565b6040518463ffffffff1660e01b8152600401612499939291906133a8565b602060405180830381865afa1580156124b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124da919061361a565b90508197508460ff168a01995060006124f7858760ff16846129a4565b9050896125045780612513565b60008a81526020829052604090205b99508298505050505050505050611cec565b600a810361268b576003821660008190036125475750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d9087926125e093929190613441565b6040518463ffffffff1660e01b81526004016125fe9392919061356a565b602060405180830381865afa15801561261b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061263f919061361a565b90508198508560ff168b019a50600061265c868860ff16846129a4565b90508a6126695780612678565b60008b81526020829052604090205b9a50829950505050505050505050611cec565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016102e1565b5094509492505050565b6000606060005b835181101561273e57818482815181106126ed576126ed613412565b6020026020010151604051602001612706929190613633565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905291506001016126d1565b50805160209091012092915050565b6000606060005b835181101561273e57600061278185838151811061277457612774613412565b6020026020010151612a12565b9050828160405160200161279692919061366b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612754565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611826565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806128ff846020015184611777565b9050600061290c85611844565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a60208201529081018290526000906060016128d1565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d0161287b565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016128d198979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b803573ffffffffffffffffffffffffffffffffffffffff81168114612ae557600080fd5b919050565b600060208284031215612afc57600080fd5b61038782612ac1565b60008083601f840112612b1757600080fd5b50813567ffffffffffffffff811115612b2f57600080fd5b602083019150836020828501011115612b4757600080fd5b9250929050565b600080600060408486031215612b6357600080fd5b83359250602084013567ffffffffffffffff811115612b8157600080fd5b612b8d86828701612b05565b9497909650939450505050565b60008060008060408587031215612bb057600080fd5b843567ffffffffffffffff811115612bc757600080fd5b612bd387828801612b05565b909550935050602085013567ffffffffffffffff811115612bf357600080fd5b612bff87828801612b05565b95989497509550505050565b600060208284031215612c1d57600080fd5b5035919050565b60008060208385031215612c3757600080fd5b823567ffffffffffffffff811115612c4e57600080fd5b612c5a85828601612b05565b90969095509350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612cb857612cb8612c66565b60405290565b604051610120810167ffffffffffffffff81118282101715612cb857612cb8612c66565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612d2957612d29612c66565b604052919050565b803560ff81168114612ae557600080fd5b80358015158114612ae557600080fd5b600067ffffffffffffffff821115612d6c57612d6c612c66565b5060051b60200190565b600082601f830112612d8757600080fd5b813567ffffffffffffffff811115612da157612da1612c66565b612dd260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612ce2565b818152846020838601011115612de757600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112612e1557600080fd5b8135612e28612e2382612d52565b612ce2565b8082825260208201915060208360051b860101925085831115612e4a57600080fd5b602085015b83811015612f3757803567ffffffffffffffff811115612e6e57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215612ea257600080fd5b612eaa612c95565b612eb660208301612ac1565b815260408201356020820152606082013567ffffffffffffffff811115612edc57600080fd5b612eeb8a602083860101612d76565b60408301525060808201356060820152612f0760a08301612d42565b6080820152612f1860c08301612d42565b60a082015260e0919091013560c0820152835260209283019201612e4f565b5095945050505050565b600082601f830112612f5257600080fd5b8135612f60612e2382612d52565b8082825260208201915060208360051b860101925085831115612f8257600080fd5b602085015b83811015612f3757612f9881612ac1565b835260209283019201612f87565b600080600060408486031215612fbb57600080fd5b833567ffffffffffffffff811115612fd257600080fd5b84016101208187031215612fe557600080fd5b612fed612cbe565b612ff682612d31565b815261300460208301612d42565b6020820152604082013567ffffffffffffffff81111561302357600080fd5b61302f88828501612e04565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561306357600080fd5b61306f88828501612d76565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff8111156130a457600080fd5b6130b088828501612f41565b61010083015250935050602084013567ffffffffffffffff811115612b8157600080fd5b6000806000606084860312156130e957600080fd5b833592506130f960208501612ac1565b915060408401356bffffffffffffffffffffffff8116811461311a57600080fd5b809150509250925092565b60005b83811015613140578181015183820152602001613128565b50506000910152565b60008151808452613161816020860160208601613125565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613263577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e0604086015261321f60e0860182613149565b6060838101519087015260808084015115159087015260a08084015115159087015260c09283015192909501919091525060209788019791909101906001016131b1565b50909695505050505050565b600081518084526020840193506020830160005b828110156132b757815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613283565b5093949350505050565b805160ff168252600060208201516132dd602085018215159052565b50604082015161012060408501526132f9610120850182613193565b9050606083015160608501526080830151608085015260a083015184820360a08601526133268282613149565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613356828261326f565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006133bb60408301866132c1565b82810360208401526133ce81858761335f565b9695505050505050565b808201808211156104e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561345157600080fd5b8386111561345e57600080fd5b5050820193919092039150565b60608152600061347e60608301866132c1565b60208301949094525060400152919050565b6060815260006134a360608301866132c1565b84602084015282810360408401526133ce8185613149565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006134ea6040830184613149565b949350505050565b6000604082840312801561350557600080fd5b506040805190810167ffffffffffffffff8111828210171561352957613529612c66565b604052825181526020928301519281019290925250919050565b60608152600061355760608301868861335f565b6020830194909452506040015292915050565b83815260406020820152600061335660408301848661335f565b60006020828403121561359657600080fd5b81517fffffffff000000000000000000000000000000000000000000000000000000008116811461038757600080fd5b6080815260006135d960808301886132c1565b86602084015273ffffffffffffffffffffffffffffffffffffffff86166040840152828103606084015261360e81858761335f565b98975050505050505050565b60006020828403121561362c57600080fd5b5051919050565b6040815260006136466040830185613149565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b6000835161367d818460208801613125565b919091019182525060200191905056fea26469706673582212203da640fc25d2fe0c4d645a2b22667cf1cf660a75777f9e61480b16db7c2dbba764736f6c634300081b00336080604052348015600f57600080fd5b5061357b8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80636ea4457711610081578063ad55366b1161005b578063ad55366b146101bb578063ca70785014610201578063f727ef1c1461021457600080fd5b80636ea44577146101685780638c3f55631461017b578063aaf10f421461018e57600080fd5b80631f6a1eb9116100b25780631f6a1eb91461012c578063295614261461013f57806351605d801461015257600080fd5b8063025b22bc146100ce5780631626ba7e146100e3575b600080fd5b6100e16100dc3660046129a2565b610227565b005b6100f66100f1366004612a06565b610273565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b6100e161013a366004612a52565b61030b565b6100e161014d366004612ac3565b61038d565b61015a6103d1565b604051908152602001610123565b6100e1610176366004612adc565b610400565b61015a610189366004612ac3565b610466565b610196610498565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610123565b6101ce6101c9366004612e5e565b6104a2565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610123565b61015a61020f366004612e5e565b6104dc565b6100e1610222366004612f8c565b610645565b333014610267576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61027081610700565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e08101859052816102ca828686610755565b509050806102de5750600091506103049050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061031785856108a3565b905061032b81606001518260800151610cc1565b600080610339838686610755565b915091508161037a578285856040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161025e93929190613260565b6103848184610da9565b50505050505050565b3330146103c8576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161025e565b6102708161103b565b60006103fb7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b33301461043b576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161025e565b600061044783836108a3565b90506000610454826110cb565b90506104608183610da9565b50505050565b60006104927f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83611146565b92915050565b60006103fb305490565b6000806000806000806104b98989896000806111a4565b9399509197509450925090506104ce836114c9565b935093975093979195509350565b6000808461010001515160016104f29190613290565b67ffffffffffffffff81111561050a5761050a612b1e565b604051908082528060200260200182016040528015610533578160200160208202803683370190505b50905060005b856101000151518110156105a557856101000151818151811061055e5761055e6132ca565b6020026020010151828281518110610578576105786132ca565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610539565b50338186610100015151815181106105bf576105bf6132ca565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006105f9868686610755565b50905080610639578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161025e93929190613260565b50600195945050505050565b333014610680576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161025e565b6106998383836bffffffffffffffffffffffff166114d4565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b610708813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b60008060008484600081811061076d5761076d6132ca565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f8000000000000000000000000000000000000000000000000000000000000000908116900361086c576107ca866110cb565b91506000806107d884611563565b9150915042811015610820576040517f9fa4fe54000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161025e565b73ffffffffffffffffffffffffffffffffffffffff82161580610858575073ffffffffffffffffffffffffffffffffffffffff821633145b15610869576001945050505061089b565b50505b600080600061087f8989896000806111a4565b985092955090935091506108949050816114c9565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c600180821681036109075760006060840152610918565b84810135606090811c908401526014015b6007600183901c16801561096b5760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610980575060016109a9565b8360201660200361099c5750600282019186013560f01c6109a9565b50600182019186013560f81c5b8067ffffffffffffffff8111156109c2576109c2612b1e565b604051908082528060200260200182016040528015610a4d57816020015b610a3a6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b8152602001906001900390816109e05790505b50604086015260005b81811015610cb65760018085019489013560f81c908082169003610ab5573087604001518381518110610a8b57610a8b6132ca565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052610aff565b8489013560601c6014860188604001518481518110610ad657610ad66132ca565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003610b3d57848901356020860188604001518481518110610b2857610b286132ca565b60200260200101516020018197508281525050505b60048082169003610bd557600385019489013560e81c89868a610b608483613290565b92610b6d939291906132f9565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506040890151805185908110610bb857610bb86132ca565b602090810291909101015160400152610bd18187613290565b9550505b60088082169003610c1357848901356020860188604001518481518110610bfe57610bfe6132ca565b60200260200101516060018197508281525050505b8060101660ff1660101487604001518381518110610c3357610c336132ca565b602002602001015160800190151590811515815250508060201660ff1660201487604001518381518110610c6957610c696132ca565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110610c9e57610c9e6132ca565b602090810291909101015160c0015250600101610a56565b505050505092915050565b6000610ccc83610466565b9050818114610d18576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161025e565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561103457600084604001518281518110610dd457610dd46132ca565b602002602001015190508060a001518015610ded575083155b15610e35576040805187815260208101849052600095507fa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7910160405180910390a15061102c565b60608101518015801590610e485750805a105b15610e855785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161025e93929190613323565b6000826080015115610eab57610ea483600001518385604001516115af565b9050610ec6565b610ec3836000015184602001518486604001516115c5565b90505b80610fef5760c0830151610f19576040805189815260208101869052600197507fe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643910160405180910390a150505061102c565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610f83578684610f4e6115dd565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161025e93929190613348565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01610fef5760408051898152602081018690527f6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8910160405180910390a1505050611034565b60408051898152602081018690527f2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3910160405180910390a15050505b600101610db4565b5050505050565b80611072576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61109b7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200161074a565b6000806110dc8360200151306115fc565b905060006110e9846116c9565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6000808383604051602001611165929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c916001918084161480156111ef575073ffffffffffffffffffffffffffffffffffffffff8916155b1561130b578b82013560601c98506014909101908961130b5760038201918c013560e81c60008d848e6112228583613290565b9261122f939291906132f9565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc891506112ba9030908590600401613373565b6040805180830381865afa1580156112d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112fa91906133aa565b92506113068285613290565b935050505b82600116600103611345576113338d8a838f8f8790809261132e939291906132f9565b61191f565b975097509750975097505050506114bc565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c168382019096509250600090506113b06001600586901c811690613290565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c16995090920191506113fb8d6110cb565b93506114198d858e8e86908092611414939291906132f9565b611b6b565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611464575080518614155b8015611474575080602001518511155b156114b8576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161025e565b5050505b9550955095509550959050565b60006104928261254f565b61155e7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b600080806115917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685611146565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de8561166c574661166f565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000806116da836101000151612582565b835190915060ff1661174d5760006116f58460400151612605565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611127565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016117dd5760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161184d5760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016117bf565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016118bd5760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e2624602082015290810191909152606081018290526080016117bf565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e6400000000000000000000000000000000604482015260640161025e565b6000806000806000611981604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b88821015611af85760038201916000908b013560e81c6119c98482613290565b9150600090508a82146119dd5760006119df565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8414611a105785611a12565b8f5b9050611a32818e8e88908792611a2a939291906132f9565b6001866111a4565b50929d50909b50995097508a8a1015611a8f578c8c86908592611a57939291906132f9565b8c8c6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161025e94939291906133fb565b829450888e6000015103611aa25760008e525b838810611ae5576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018990526024810185905260440161025e565b505060c0840187905291508490506119a9565b8a5115801590611b0c57508a602001518511155b15611b50576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161025e565b611b598d6110cb565b93505050509550955095509550959050565b60008060005b8381101561254557600181019085013560f881901c9060fc1c80611c8e57600f82166000819003611ba95750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015611c3e573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000611c61828960ff16612687565b90508b611c6e5780611c7d565b60008c81526020829052604090205b9b5050505050505050505050611b71565b60018103611cf257600f82166000819003611cb05750600183019287013560f81c5b601484019388013560601c6000611cca8260ff8516612687565b905086611cd75780611ce6565b60008781526020829052604090205b96505050505050611b71565b60028103611ed957600382166000819003611d145750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d908792611dd2939291906132f9565b6040518463ffffffff1660e01b8152600401611df093929190613422565b602060405180830381865afa158015611e0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e31919061343c565b7fffffffff000000000000000000000000000000000000000000000000000000001614611e94578d8d858e8e6040517ff734863a00000000000000000000000000000000000000000000000000000000815260040161025e95949392919061347e565b8097508460ff168a0199506000611eae858760ff16612687565b905089611ebb5780611eca565b60008a81526020829052604090205b99505050505050505050611b71565b60038103611f0d57602083019287013584611ef45780611f03565b60008581526020829052604090205b9450505050611b71565b60048103611fb357600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c16858301809650819250505060008186019050600080611f828e8e8e8e8c908892611414939291906132f9565b91509150829750818a019950611fa2898260009182526020526040902090565b985082975050505050505050611b71565b6006810361207d576003600283901c166000819003611fd95750600183019287013560f81c5b600383166000819003611ff35750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806120318f8f8f8f8d908892611414939291906132f9565b9150915082985084821061204457998501995b60006120518287896126ee565b90508a61205e578061206d565b60008b81526020829052604090205b9a50505050505050505050611b71565b600581036120ea5760208301928701358881036120b8577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b60006120c382612752565b9050856120d057806120df565b60008681526020829052604090205b955050505050611b71565b600781036121f057600f8216600081900361210c5750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001611c1c565b60088103612244576020830192870135600061220c8b826127a6565b9050808203612239577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b6000611cca82612821565b600981036123aa576003821660008190036122665750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612300939291906132f9565b6040518463ffffffff1660e01b815260040161231e93929190613260565b602060405180830381865afa15801561233b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061235f91906134d2565b90508197508460ff168a019950600061237c858760ff168461285c565b9050896123895780612398565b60008a81526020829052604090205b99508298505050505050505050611b71565b600a8103612510576003821660008190036123cc5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612465939291906132f9565b6040518463ffffffff1660e01b815260040161248393929190613422565b602060405180830381865afa1580156124a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124c491906134d2565b90508198508560ff168b019a5060006124e1868860ff168461285c565b90508a6124ee57806124fd565b60008b81526020829052604090205b9a50829950505050505050505050611b71565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161025e565b5094509492505050565b600081158015906104925750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b6000606060005b83518110156125f657818482815181106125a5576125a56132ca565b60200260200101516040516020016125be9291906134eb565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612589565b50805160209091012092915050565b6000606060005b83518110156125f657600061263985838151811061262c5761262c6132ca565b60200260200101516128ca565b9050828160405160200161264e929190613523565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905292505060010161260c565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b166031820152604581018290526000906065016116ab565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806127b78460200151846115fc565b905060006127c4856116c9565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612789565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612733565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161278998979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b803573ffffffffffffffffffffffffffffffffffffffff8116811461299d57600080fd5b919050565b6000602082840312156129b457600080fd5b61030482612979565b60008083601f8401126129cf57600080fd5b50813567ffffffffffffffff8111156129e757600080fd5b6020830191508360208285010111156129ff57600080fd5b9250929050565b600080600060408486031215612a1b57600080fd5b83359250602084013567ffffffffffffffff811115612a3957600080fd5b612a45868287016129bd565b9497909650939450505050565b60008060008060408587031215612a6857600080fd5b843567ffffffffffffffff811115612a7f57600080fd5b612a8b878288016129bd565b909550935050602085013567ffffffffffffffff811115612aab57600080fd5b612ab7878288016129bd565b95989497509550505050565b600060208284031215612ad557600080fd5b5035919050565b60008060208385031215612aef57600080fd5b823567ffffffffffffffff811115612b0657600080fd5b612b12858286016129bd565b90969095509350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715612b7057612b70612b1e565b60405290565b604051610120810167ffffffffffffffff81118282101715612b7057612b70612b1e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612be157612be1612b1e565b604052919050565b803560ff8116811461299d57600080fd5b8035801515811461299d57600080fd5b600067ffffffffffffffff821115612c2457612c24612b1e565b5060051b60200190565b600082601f830112612c3f57600080fd5b813567ffffffffffffffff811115612c5957612c59612b1e565b612c8a60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601612b9a565b818152846020838601011115612c9f57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112612ccd57600080fd5b8135612ce0612cdb82612c0a565b612b9a565b8082825260208201915060208360051b860101925085831115612d0257600080fd5b602085015b83811015612def57803567ffffffffffffffff811115612d2657600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215612d5a57600080fd5b612d62612b4d565b612d6e60208301612979565b815260408201356020820152606082013567ffffffffffffffff811115612d9457600080fd5b612da38a602083860101612c2e565b60408301525060808201356060820152612dbf60a08301612bfa565b6080820152612dd060c08301612bfa565b60a082015260e0919091013560c0820152835260209283019201612d07565b5095945050505050565b600082601f830112612e0a57600080fd5b8135612e18612cdb82612c0a565b8082825260208201915060208360051b860101925085831115612e3a57600080fd5b602085015b83811015612def57612e5081612979565b835260209283019201612e3f565b600080600060408486031215612e7357600080fd5b833567ffffffffffffffff811115612e8a57600080fd5b84016101208187031215612e9d57600080fd5b612ea5612b76565b612eae82612be9565b8152612ebc60208301612bfa565b6020820152604082013567ffffffffffffffff811115612edb57600080fd5b612ee788828501612cbc565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff811115612f1b57600080fd5b612f2788828501612c2e565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff811115612f5c57600080fd5b612f6888828501612df9565b61010083015250935050602084013567ffffffffffffffff811115612a3957600080fd5b600080600060608486031215612fa157600080fd5b83359250612fb160208501612979565b915060408401356bffffffffffffffffffffffff81168114612fd257600080fd5b809150509250925092565b60005b83811015612ff8578181015183820152602001612fe0565b50506000910152565b60008151808452613019816020860160208601612fdd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b8381101561311b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e060408601526130d760e0860182613001565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613069565b50909695505050505050565b600081518084526020840193506020830160005b8281101561316f57815173ffffffffffffffffffffffffffffffffffffffff1686526020958601959091019060010161313b565b5093949350505050565b805160ff16825260006020820151613195602085018215159052565b50604082015161012060408501526131b161012085018261304b565b9050606083015160608501526080830151608085015260a083015184820360a08601526131de8282613001565b91505060c083015160c085015260e083015160e085015261010083015184820361010086015261320e8282613127565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006132736040830186613179565b8281036020840152613286818587613217565b9695505050505050565b80820180821115610492577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561330957600080fd5b8386111561331657600080fd5b5050820193919092039150565b6060815260006133366060830186613179565b60208301949094525060400152919050565b60608152600061335b6060830186613179565b84602084015282810360408401526132868185613001565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006133a26040830184613001565b949350505050565b600060408284031280156133bd57600080fd5b506040805190810167ffffffffffffffff811182821017156133e1576133e1612b1e565b604052825181526020928301519281019290925250919050565b60608152600061340f606083018688613217565b6020830194909452506040015292915050565b83815260406020820152600061320e604083018486613217565b60006020828403121561344e57600080fd5b81517fffffffff000000000000000000000000000000000000000000000000000000008116811461030457600080fd5b6080815260006134916080830188613179565b86602084015273ffffffffffffffffffffffffffffffffffffffff8616604084015282810360608401526134c6818587613217565b98975050505050505050565b6000602082840312156134e457600080fd5b5051919050565b6040815260006134fe6040830185613001565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351613535818460208801612fdd565b919091019182525060200191905056fea26469706673582212204ff581b28abced44877e9daf41810be6ae2a2049d64a6efc732528bc7a9900af64736f6c634300081b0033603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3",
}

// WalletUpgradableABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletUpgradableMetaData.ABI instead.
var WalletUpgradableABI = WalletUpgradableMetaData.ABI

// WalletUpgradableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletUpgradableMetaData.Bin instead.
var WalletUpgradableBin = WalletUpgradableMetaData.Bin

// DeployWalletUpgradable deploys a new Ethereum contract, binding an instance of WalletUpgradable to it.
func DeployWalletUpgradable(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *WalletUpgradable, error) {
	parsed, err := WalletUpgradableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletUpgradableBin), backend, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletUpgradable{WalletUpgradableCaller: WalletUpgradableCaller{contract: contract}, WalletUpgradableTransactor: WalletUpgradableTransactor{contract: contract}, WalletUpgradableFilterer: WalletUpgradableFilterer{contract: contract}}, nil
}

// WalletUpgradable is an auto generated Go binding around an Ethereum contract.
type WalletUpgradable struct {
	WalletUpgradableCaller     // Read-only binding to the contract
	WalletUpgradableTransactor // Write-only binding to the contract
	WalletUpgradableFilterer   // Log filterer for contract events
}

// WalletUpgradableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletUpgradableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUpgradableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletUpgradableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUpgradableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletUpgradableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletUpgradableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletUpgradableSession struct {
	Contract     *WalletUpgradable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletUpgradableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletUpgradableCallerSession struct {
	Contract *WalletUpgradableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// WalletUpgradableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletUpgradableTransactorSession struct {
	Contract     *WalletUpgradableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// WalletUpgradableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletUpgradableRaw struct {
	Contract *WalletUpgradable // Generic contract binding to access the raw methods on
}

// WalletUpgradableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletUpgradableCallerRaw struct {
	Contract *WalletUpgradableCaller // Generic read-only contract binding to access the raw methods on
}

// WalletUpgradableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletUpgradableTransactorRaw struct {
	Contract *WalletUpgradableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletUpgradable creates a new instance of WalletUpgradable, bound to a specific deployed contract.
func NewWalletUpgradable(address common.Address, backend bind.ContractBackend) (*WalletUpgradable, error) {
	contract, err := bindWalletUpgradable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletUpgradable{WalletUpgradableCaller: WalletUpgradableCaller{contract: contract}, WalletUpgradableTransactor: WalletUpgradableTransactor{contract: contract}, WalletUpgradableFilterer: WalletUpgradableFilterer{contract: contract}}, nil
}

// NewWalletUpgradableCaller creates a new read-only instance of WalletUpgradable, bound to a specific deployed contract.
func NewWalletUpgradableCaller(address common.Address, caller bind.ContractCaller) (*WalletUpgradableCaller, error) {
	contract, err := bindWalletUpgradable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableCaller{contract: contract}, nil
}

// NewWalletUpgradableTransactor creates a new write-only instance of WalletUpgradable, bound to a specific deployed contract.
func NewWalletUpgradableTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletUpgradableTransactor, error) {
	contract, err := bindWalletUpgradable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableTransactor{contract: contract}, nil
}

// NewWalletUpgradableFilterer creates a new log filterer instance of WalletUpgradable, bound to a specific deployed contract.
func NewWalletUpgradableFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletUpgradableFilterer, error) {
	contract, err := bindWalletUpgradable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableFilterer{contract: contract}, nil
}

// bindWalletUpgradable binds a generic wrapper to an already deployed contract.
func bindWalletUpgradable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletUpgradableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletUpgradable *WalletUpgradableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletUpgradable.Contract.WalletUpgradableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletUpgradable *WalletUpgradableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.WalletUpgradableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletUpgradable *WalletUpgradableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.WalletUpgradableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletUpgradable *WalletUpgradableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletUpgradable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletUpgradable *WalletUpgradableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletUpgradable *WalletUpgradableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.contract.Transact(opts, method, params...)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletUpgradable *WalletUpgradableCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletUpgradable *WalletUpgradableSession) FACTORY() (common.Address, error) {
	return _WalletUpgradable.Contract.FACTORY(&_WalletUpgradable.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletUpgradable *WalletUpgradableCallerSession) FACTORY() (common.Address, error) {
	return _WalletUpgradable.Contract.FACTORY(&_WalletUpgradable.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableCaller) INITCODEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "INIT_CODE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableSession) INITCODEHASH() ([32]byte, error) {
	return _WalletUpgradable.Contract.INITCODEHASH(&_WalletUpgradable.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableCallerSession) INITCODEHASH() ([32]byte, error) {
	return _WalletUpgradable.Contract.INITCODEHASH(&_WalletUpgradable.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletUpgradable *WalletUpgradableCaller) STAGE2IMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "STAGE_2_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletUpgradable *WalletUpgradableSession) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletUpgradable.Contract.STAGE2IMPLEMENTATION(&_WalletUpgradable.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletUpgradable *WalletUpgradableCallerSession) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletUpgradable.Contract.STAGE2IMPLEMENTATION(&_WalletUpgradable.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletUpgradable *WalletUpgradableCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletUpgradable *WalletUpgradableSession) GetImplementation() (common.Address, error) {
	return _WalletUpgradable.Contract.GetImplementation(&_WalletUpgradable.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletUpgradable *WalletUpgradableCallerSession) GetImplementation() (common.Address, error) {
	return _WalletUpgradable.Contract.GetImplementation(&_WalletUpgradable.CallOpts)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableCaller) IsValidSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "isValidSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletUpgradable.Contract.IsValidSapientSignature(&_WalletUpgradable.CallOpts, _payload, _signature)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletUpgradable *WalletUpgradableCallerSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletUpgradable.Contract.IsValidSapientSignature(&_WalletUpgradable.CallOpts, _payload, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletUpgradable *WalletUpgradableCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletUpgradable *WalletUpgradableSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletUpgradable.Contract.IsValidSignature(&_WalletUpgradable.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletUpgradable *WalletUpgradableCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletUpgradable.Contract.IsValidSignature(&_WalletUpgradable.CallOpts, _hash, _signature)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletUpgradable *WalletUpgradableCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletUpgradable *WalletUpgradableSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletUpgradable.Contract.ReadNonce(&_WalletUpgradable.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletUpgradable *WalletUpgradableCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletUpgradable.Contract.ReadNonce(&_WalletUpgradable.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletUpgradable *WalletUpgradableCaller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletUpgradable.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

	outstruct := new(struct {
		Threshold    *big.Int
		Weight       *big.Int
		IsValidImage bool
		ImageHash    [32]byte
		Checkpoint   *big.Int
		OpHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Threshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Weight = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsValidImage = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.ImageHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Checkpoint = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.OpHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletUpgradable *WalletUpgradableSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletUpgradable.Contract.RecoverPartialSignature(&_WalletUpgradable.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletUpgradable *WalletUpgradableCallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletUpgradable.Contract.RecoverPartialSignature(&_WalletUpgradable.CallOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletUpgradable *WalletUpgradableTransactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletUpgradable.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletUpgradable *WalletUpgradableSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.Execute(&_WalletUpgradable.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletUpgradable *WalletUpgradableTransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.Execute(&_WalletUpgradable.TransactOpts, _payload, _signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletUpgradable *WalletUpgradableTransactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletUpgradable.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletUpgradable *WalletUpgradableSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.SelfExecute(&_WalletUpgradable.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletUpgradable *WalletUpgradableTransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.SelfExecute(&_WalletUpgradable.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletUpgradable *WalletUpgradableTransactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletUpgradable.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletUpgradable *WalletUpgradableSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.SetStaticSignature(&_WalletUpgradable.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletUpgradable *WalletUpgradableTransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.SetStaticSignature(&_WalletUpgradable.TransactOpts, _hash, _address, _timestamp)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletUpgradable *WalletUpgradableTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletUpgradable.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletUpgradable *WalletUpgradableSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.UpdateImageHash(&_WalletUpgradable.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletUpgradable *WalletUpgradableTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.UpdateImageHash(&_WalletUpgradable.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletUpgradable *WalletUpgradableTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletUpgradable.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletUpgradable *WalletUpgradableSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.UpdateImplementation(&_WalletUpgradable.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletUpgradable *WalletUpgradableTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletUpgradable.Contract.UpdateImplementation(&_WalletUpgradable.TransactOpts, _implementation)
}

// WalletUpgradableAbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the WalletUpgradable contract.
type WalletUpgradableAbortedIterator struct {
	Event *WalletUpgradableAborted // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableAborted)
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
		it.Event = new(WalletUpgradableAborted)
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
func (it *WalletUpgradableAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableAborted represents a Aborted event raised by the WalletUpgradable contract.
type WalletUpgradableAborted struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterAborted(opts *bind.FilterOpts) (*WalletUpgradableAbortedIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableAbortedIterator{contract: _WalletUpgradable.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *WalletUpgradableAborted) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableAborted)
				if err := _WalletUpgradable.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// ParseAborted is a log parse operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseAborted(log types.Log) (*WalletUpgradableAborted, error) {
	event := new(WalletUpgradableAborted)
	if err := _WalletUpgradable.contract.UnpackLog(event, "Aborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableFailedIterator is returned from FilterFailed and is used to iterate over the raw logs and unpacked data for Failed events raised by the WalletUpgradable contract.
type WalletUpgradableFailedIterator struct {
	Event *WalletUpgradableFailed // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableFailed)
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
		it.Event = new(WalletUpgradableFailed)
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
func (it *WalletUpgradableFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableFailed represents a Failed event raised by the WalletUpgradable contract.
type WalletUpgradableFailed struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailed is a free log retrieval operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterFailed(opts *bind.FilterOpts) (*WalletUpgradableFailedIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableFailedIterator{contract: _WalletUpgradable.contract, event: "Failed", logs: logs, sub: sub}, nil
}

// WatchFailed is a free log subscription operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchFailed(opts *bind.WatchOpts, sink chan<- *WalletUpgradableFailed) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableFailed)
				if err := _WalletUpgradable.contract.UnpackLog(event, "Failed", log); err != nil {
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

// ParseFailed is a log parse operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseFailed(log types.Log) (*WalletUpgradableFailed, error) {
	event := new(WalletUpgradableFailed)
	if err := _WalletUpgradable.contract.UnpackLog(event, "Failed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletUpgradable contract.
type WalletUpgradableImageHashUpdatedIterator struct {
	Event *WalletUpgradableImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableImageHashUpdated)
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
		it.Event = new(WalletUpgradableImageHashUpdated)
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
func (it *WalletUpgradableImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableImageHashUpdated represents a ImageHashUpdated event raised by the WalletUpgradable contract.
type WalletUpgradableImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletUpgradableImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableImageHashUpdatedIterator{contract: _WalletUpgradable.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletUpgradableImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableImageHashUpdated)
				if err := _WalletUpgradable.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletUpgradable *WalletUpgradableFilterer) ParseImageHashUpdated(log types.Log) (*WalletUpgradableImageHashUpdated, error) {
	event := new(WalletUpgradableImageHashUpdated)
	if err := _WalletUpgradable.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletUpgradable contract.
type WalletUpgradableImplementationUpdatedIterator struct {
	Event *WalletUpgradableImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableImplementationUpdated)
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
		it.Event = new(WalletUpgradableImplementationUpdated)
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
func (it *WalletUpgradableImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableImplementationUpdated represents a ImplementationUpdated event raised by the WalletUpgradable contract.
type WalletUpgradableImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletUpgradableImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableImplementationUpdatedIterator{contract: _WalletUpgradable.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletUpgradableImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableImplementationUpdated)
				if err := _WalletUpgradable.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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

// ParseImplementationUpdated is a log parse operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseImplementationUpdated(log types.Log) (*WalletUpgradableImplementationUpdated, error) {
	event := new(WalletUpgradableImplementationUpdated)
	if err := _WalletUpgradable.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletUpgradable contract.
type WalletUpgradableNonceChangeIterator struct {
	Event *WalletUpgradableNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableNonceChange)
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
		it.Event = new(WalletUpgradableNonceChange)
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
func (it *WalletUpgradableNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableNonceChange represents a NonceChange event raised by the WalletUpgradable contract.
type WalletUpgradableNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletUpgradableNonceChangeIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableNonceChangeIterator{contract: _WalletUpgradable.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletUpgradableNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableNonceChange)
				if err := _WalletUpgradable.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletUpgradable *WalletUpgradableFilterer) ParseNonceChange(log types.Log) (*WalletUpgradableNonceChange, error) {
	event := new(WalletUpgradableNonceChange)
	if err := _WalletUpgradable.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableSkippedIterator is returned from FilterSkipped and is used to iterate over the raw logs and unpacked data for Skipped events raised by the WalletUpgradable contract.
type WalletUpgradableSkippedIterator struct {
	Event *WalletUpgradableSkipped // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableSkipped)
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
		it.Event = new(WalletUpgradableSkipped)
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
func (it *WalletUpgradableSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableSkipped represents a Skipped event raised by the WalletUpgradable contract.
type WalletUpgradableSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkipped is a free log retrieval operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterSkipped(opts *bind.FilterOpts) (*WalletUpgradableSkippedIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableSkippedIterator{contract: _WalletUpgradable.contract, event: "Skipped", logs: logs, sub: sub}, nil
}

// WatchSkipped is a free log subscription operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchSkipped(opts *bind.WatchOpts, sink chan<- *WalletUpgradableSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableSkipped)
				if err := _WalletUpgradable.contract.UnpackLog(event, "Skipped", log); err != nil {
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

// ParseSkipped is a log parse operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseSkipped(log types.Log) (*WalletUpgradableSkipped, error) {
	event := new(WalletUpgradableSkipped)
	if err := _WalletUpgradable.contract.UnpackLog(event, "Skipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableStaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletUpgradable contract.
type WalletUpgradableStaticSignatureSetIterator struct {
	Event *WalletUpgradableStaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableStaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableStaticSignatureSet)
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
		it.Event = new(WalletUpgradableStaticSignatureSet)
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
func (it *WalletUpgradableStaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableStaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableStaticSignatureSet represents a StaticSignatureSet event raised by the WalletUpgradable contract.
type WalletUpgradableStaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletUpgradableStaticSignatureSetIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableStaticSignatureSetIterator{contract: _WalletUpgradable.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletUpgradableStaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableStaticSignatureSet)
				if err := _WalletUpgradable.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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

// ParseStaticSignatureSet is a log parse operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseStaticSignatureSet(log types.Log) (*WalletUpgradableStaticSignatureSet, error) {
	event := new(WalletUpgradableStaticSignatureSet)
	if err := _WalletUpgradable.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletUpgradableSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the WalletUpgradable contract.
type WalletUpgradableSuccessIterator struct {
	Event *WalletUpgradableSuccess // Event containing the contract specifics and raw log

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
func (it *WalletUpgradableSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletUpgradableSuccess)
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
		it.Event = new(WalletUpgradableSuccess)
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
func (it *WalletUpgradableSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletUpgradableSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletUpgradableSuccess represents a Success event raised by the WalletUpgradable contract.
type WalletUpgradableSuccess struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) FilterSuccess(opts *bind.FilterOpts) (*WalletUpgradableSuccessIterator, error) {

	logs, sub, err := _WalletUpgradable.contract.FilterLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return &WalletUpgradableSuccessIterator{contract: _WalletUpgradable.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *WalletUpgradableSuccess) (event.Subscription, error) {

	logs, sub, err := _WalletUpgradable.contract.WatchLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletUpgradableSuccess)
				if err := _WalletUpgradable.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletUpgradable *WalletUpgradableFilterer) ParseSuccess(log types.Log) (*WalletUpgradableSuccess, error) {
	event := new(WalletUpgradableSuccess)
	if err := _WalletUpgradable.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
