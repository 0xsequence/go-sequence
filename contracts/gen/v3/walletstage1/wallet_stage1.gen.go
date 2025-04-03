// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletstage1

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

// WalletStage1MetaData contains all meta data concerning the WalletStage1 contract.
var WalletStage1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAGE_2_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561001057600080fd5b506040516181ec3803806181ec83398101604081905261002f916100c7565b8060405161003c906100ba565b604051809103906000f080158015610058573d6000803e3d6000fd5b5060006040518060600160405280602c81526020016181c0602c9139604051610086919030906020016100f7565b60408051601f198184030181529190528051602090910120608052506001600160a01b0391821660a0521660c05250610129565b613f788061424883390190565b6000602082840312156100d957600080fd5b81516001600160a01b03811681146100f057600080fd5b9392505050565b6000835160005b8181101561011857602081870181015185830152016100fe565b509190910191825250602001919050565b60805160a05160c0516140db61016d60003960008181610462015261178b0152600081816103c70152611c610152600081816103730152611c9201526140db6000f3fe6080604052600436106101635760003560e01c80636ea44577116100c0578063ad55366b11610074578063bc197c8111610059578063bc197c81146104ff578063f23a6e6114610547578063f727ef1c1461058d5761016a565b8063ad55366b14610499578063b93ea7ad146104ec5761016a565b80638c3f5563116100a55780638c3f5563146104305780639f69ef5414610450578063aaf10f42146104845761016a565b80636ea44577146103fc5780638943ec021461040f5761016a565b80631f6a1eb91161011757806329561426116100fc57806329561426146103955780632dd31000146103b55780634fcf3eca146103e95761016a565b80631f6a1eb91461034e578063257671f5146103615761016a565b8063150b7a0211610148578063150b7a02146102735780631626ba7e146102e95780631a9b2337146103095761016a565b8063025b22bc1461022d57806313792a4a146102405761016a565b3661016a57005b6004361061022b5760006101866101813683613183565b6105ad565b905073ffffffffffffffffffffffffffffffffffffffff811615610229576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101cf9291906131e9565b600060405180830381855af49150503d806000811461020a576040519150601f19603f3d011682016040523d82523d6000602084013e61020f565b606091505b50915091508161022157805160208201fd5b805160208201f35b505b005b61022b61023b366004613222565b610601565b34801561024c57600080fd5b5061026061025b3660046135c6565b61064d565b6040519081526020015b60405180910390f35b34801561027f57600080fd5b506102b861028e36600461370d565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161026a565b3480156102f557600080fd5b506102b861030436600461377c565b6107b8565b34801561031557600080fd5b506103296103243660046137dd565b61084e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161026a565b61022b61035c3660046137fa565b610859565b34801561036d57600080fd5b506102607f000000000000000000000000000000000000000000000000000000000000000081565b3480156103a157600080fd5b5061022b6103b036600461386b565b6108e2565b3480156103c157600080fd5b506103297f000000000000000000000000000000000000000000000000000000000000000081565b61022b6103f73660046137dd565b610926565b61022b61040a366004613884565b6109e8565b34801561041b57600080fd5b5061022b61042a3660046138c6565b50505050565b34801561043c57600080fd5b5061026061044b36600461386b565b610a55565b34801561045c57600080fd5b506103297f000000000000000000000000000000000000000000000000000000000000000081565b34801561049057600080fd5b50610329610a81565b3480156104a557600080fd5b506104b96104b43660046135c6565b610a90565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c00161026a565b61022b6104fa366004613908565b610aca565b34801561050b57600080fd5b506102b861051a366004613982565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561055357600080fd5b506102b8610562366004613a49565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561059957600080fd5b5061022b6105a8366004613ac1565b610b8f565b60006105fb7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c4a565b92915050565b333014610641576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61064a81610ca8565b50565b6000808461010001515160016106639190613b12565b67ffffffffffffffff81111561067b5761067b61323d565b6040519080825280602002602001820160405280156106a4578160200160208202803683370190505b50905060005b856101000151518110156107165785610100015181815181106106cf576106cf613b4c565b60200260200101518282815181106106e9576106e9613b4c565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016106aa565b503381866101000151518151811061073057610730613b4c565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526101008501819052600061076a868686610cfc565b509050806107aa578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161063893929190613dfe565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e081018590528161080f828686610cfc565b509050806108235750600091506107b19050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b60006105fb826105ad565b60005a9050600061086a8686610e8d565b905061087e816060015182608001516112ab565b60008061088c838787610cfc565b91509150816108cd578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161063893929190613dfe565b6108d8848285611393565b5050505050505050565b33301461091d576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610638565b61064a816116f3565b333014610961576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610638565b600061096c826105ad565b73ffffffffffffffffffffffffffffffffffffffff16036109dd576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000082166004820152602401610638565b61064a8160006117af565b333014610a23576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610638565b60005a90506000610a348484610e8d565b90506000610a418261186f565b9050610a4e838284611393565b5050505050565b60006105fb7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c4a565b6000610a8b305490565b905090565b600080600080600080610aa78989896000806118ea565b939950919750945092509050610abc83611c0f565b935093975093979195509350565b333014610b05576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610638565b6000610b10836105ad565b73ffffffffffffffffffffffffffffffffffffffff1614610b81576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610638565b610b8b82826117af565b5050565b333014610bca576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610638565b610be38383836bffffffffffffffffffffffff16611d11565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c69929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610cb0813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039060200160405180910390a150565b600080600084846000818110610d1457610d14613b4c565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e1357610d718661186f565b9150600080610d7f84611da0565b9150915042811015610dc7576040517f9fa4fe540000000000000000000000000000000000000000000000000000000081526004810185905260248101829052604401610638565b73ffffffffffffffffffffffffffffffffffffffff82161580610dff575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e105760019450505050610e85565b50505b6000806000610e268989896000806118ea565b985092955090935091505082821015610e75576040517ffd41fcba0000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604401610638565b610e7e81611c0f565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610ef15760006060840152610f02565b84810135606090811c908401526014015b6007600183901c168015610f555760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610f6a57506001610f93565b83602016602003610f865750600282019186013560f01c610f93565b50600182019186013560f81c5b8067ffffffffffffffff811115610fac57610fac61323d565b60405190808252806020026020018201604052801561103757816020015b6110246040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610fca5790505b50604086015260005b818110156112a05760018085019489013560f81c90808216900361109f57308760400151838151811061107557611075613b4c565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526110e9565b8489013560601c60148601886040015184815181106110c0576110c0613b4c565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036111275784890135602086018860400151848151811061111257611112613b4c565b60200260200101516020018197508281525050505b600480821690036111bf57600385019489013560e81c89868a61114a8483613b12565b9261115793929190613e2e565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106111a2576111a2613b4c565b6020908102919091010151604001526111bb8187613b12565b9550505b600880821690036111fd578489013560208601886040015184815181106111e8576111e8613b4c565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061121d5761121d613b4c565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061125357611253613b4c565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061128857611288613b4c565b602090810291909101015160c0015250600101611040565b505050505092915050565b60006112b683610a55565b9050818114611302576040517f9b6514f4000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260448101829052606401610638565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b818110156116eb576000846040015182815181106113be576113be613b4c565b6020026020010151905083156113d75760009350611420565b8060a00151156114205760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506116e3565b606081015180158015906114335750805a105b156114705785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161063893929190613e58565b600082608001511561154457825161153d90831561148e5783611490565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016114bb96959493929190613e7d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611dec565b905061156b565b825160208401516115689190841561155c578461155e565b5a5b8660400151611e02565b90505b806116a65760c08301516115c757600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856115a8611e1a565b6040516115b793929190613eba565b60405180910390a15050506116e3565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016116315786846115fc611e1a565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161063893929190613ed9565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016116a6577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b8885611687611e1a565b60405161169693929190613eba565b60405180910390a15050506116eb565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b60010161139e565b505050505050565b8061172a576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117537fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200160405180910390a161064a7f0000000000000000000000000000000000000000000000000000000000000000610ca8565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b600080611880836020015130611e39565b9050600061188d84611f06565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611935575073ffffffffffffffffffffffffffffffffffffffff8916155b15611a51578b82013560601c985060149091019089611a515760038201918c013560e81c60008d848e6119688583613b12565b9261197593929190613e2e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611a009030908590600401613f04565b6040805180830381865afa158015611a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a409190613f3b565b9250611a4c8285613b12565b935050505b82600116600103611a8b57611a798d8a838f8f87908092611a7493929190613e2e565b61215c565b97509750975097509750505050611c02565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611af66001600586901c811690613b12565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611b418d61186f565b9350611b5f8d858e8e86908092611b5a93929190613e2e565b6123a8565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611baa575080518614155b8015611bba575080602001518511155b15611bfe576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528151600482015260208201516024820152604401610638565b5050505b9550955095509550959050565b60006105fb826040517fff0000000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060601b166021820152603581018290527f000000000000000000000000000000000000000000000000000000000000000060558201526000903090607501604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012073ffffffffffffffffffffffffffffffffffffffff161492915050565b611d9b7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611dce7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c4a565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611ea95746611eac565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611f17836101000151612d8c565b835190915060ff16611f8a576000611f328460400151612e0f565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016118cb565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161201a5760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161208a5760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001611ffc565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016120fa5760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201529081019190915260608101829052608001611ffc565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e64000000000000000000000000000000006044820152606401610638565b60008060008060006121be604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156123355760038201916000908b013560e81c6122068482613b12565b9150600090508a821461221a57600061221c565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841461224d578561224f565b8f5b905061226f818e8e8890879261226793929190613e2e565b6001866118ea565b50929d50909b50995097508a8a10156122cc578c8c8690859261229493929190613e2e565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016106389493929190613f8c565b829450888e60000151036122df5760008e525b838810612322576040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004810189905260248101859052604401610638565b505060c0840187905291508490506121e6565b8a511580159061234957508a602001518511155b1561238d576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c01516024820152604401610638565b6123968d61186f565b93505050509550955095509550959050565b60008060005b83811015612d8257600181019085013560f881901c9060fc1c806124cb57600f821660008190036123e65750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa15801561247b573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b50600061249e828960ff16612e91565b90508b6124ab57806124ba565b60008c81526020829052604090205b9b50505050505050505050506123ae565b6001810361252f57600f821660008190036124ed5750600183019287013560f81c5b601484019388013560601c60006125078260ff8516612e91565b9050866125145780612523565b60008781526020829052604090205b965050505050506123ae565b60028103612716576003821660008190036125515750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261260f93929190613e2e565b6040518463ffffffff1660e01b815260040161262d93929190613fb3565b602060405180830381865afa15801561264a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061266e9190613fcd565b7fffffffff0000000000000000000000000000000000000000000000000000000016146126d1578d8d858e8e6040517ff734863a000000000000000000000000000000000000000000000000000000008152600401610638959493929190613fea565b8097508460ff168a01995060006126eb858760ff16612e91565b9050896126f85780612707565b60008a81526020829052604090205b995050505050505050506123ae565b6003810361274a576020830192870135846127315780612740565b60008581526020829052604090205b94505050506123ae565b600481036127f057600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806127bf8e8e8e8e8c908892611b5a93929190613e2e565b91509150829750818a0199506127df898260009182526020526040902090565b9850829750505050505050506123ae565b600681036128ba576003600283901c1660008190036128165750600183019287013560f81c5b6003831660008190036128305750600284019388013560f01c5b6000858a013560e81c600387018162ffffff16915080975081925050506000818701905060008061286e8f8f8f8f8d908892611b5a93929190613e2e565b9150915082985084821061288157998501995b600061288e828789612ef8565b90508a61289b57806128aa565b60008b81526020829052604090205b9a505050505050505050506123ae565b600581036129275760208301928701358881036128f5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061290082612f5c565b90508561290d578061291c565b60008681526020829052604090205b9550505050506123ae565b60078103612a2d57600f821660008190036129495750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612459565b60088103612a815760208301928701356000612a498b82612fb0565b9050808203612a76577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006125078261302b565b60098103612be757600382166000819003612aa35750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612b3d93929190613e2e565b6040518463ffffffff1660e01b8152600401612b5b93929190613dfe565b602060405180830381865afa158015612b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b9c9190614032565b90508197508460ff168a0199506000612bb9858760ff1684613066565b905089612bc65780612bd5565b60008a81526020829052604090205b995082985050505050505050506123ae565b600a8103612d4d57600382166000819003612c095750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612ca293929190613e2e565b6040518463ffffffff1660e01b8152600401612cc093929190613fb3565b602060405180830381865afa158015612cdd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d019190614032565b90508198508560ff168b019a506000612d1e868860ff1684613066565b90508a612d2b5780612d3a565b60008b81526020829052604090205b9a508299505050505050505050506123ae565b6040517fb2505f7c00000000000000000000000000000000000000000000000000000000815260048101829052602401610638565b5094509492505050565b6000606060005b8351811015612e005781848281518110612daf57612daf613b4c565b6020026020010151604051602001612dc892919061404b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612d93565b50805160209091012092915050565b6000606060005b8351811015612e00576000612e43858381518110612e3657612e36613b4c565b60200260200101516130d4565b90508281604051602001612e58929190614083565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612e16565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611ee8565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612fc1846020015184611e39565b90506000612fce85611f06565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612f93565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612f3d565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612f9398979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff0000000000000000000000000000000000000000000000000000000081169060048410156131e2577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461321d57600080fd5b919050565b60006020828403121561323457600080fd5b6107b1826131f9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561328f5761328f61323d565b60405290565b604051610120810167ffffffffffffffff8111828210171561328f5761328f61323d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156133005761330061323d565b604052919050565b803560ff8116811461321d57600080fd5b8035801515811461321d57600080fd5b600067ffffffffffffffff8211156133435761334361323d565b5060051b60200190565b600082601f83011261335e57600080fd5b813567ffffffffffffffff8111156133785761337861323d565b6133a960207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016132b9565b8181528460208386010111156133be57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126133ec57600080fd5b81356133ff6133fa82613329565b6132b9565b8082825260208201915060208360051b86010192508583111561342157600080fd5b602085015b8381101561350e57803567ffffffffffffffff81111561344557600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561347957600080fd5b61348161326c565b61348d602083016131f9565b815260408201356020820152606082013567ffffffffffffffff8111156134b357600080fd5b6134c28a60208386010161334d565b604083015250608082013560608201526134de60a08301613319565b60808201526134ef60c08301613319565b60a082015260e0919091013560c0820152835260209283019201613426565b5095945050505050565b600082601f83011261352957600080fd5b81356135376133fa82613329565b8082825260208201915060208360051b86010192508583111561355957600080fd5b602085015b8381101561350e5761356f816131f9565b83526020928301920161355e565b60008083601f84011261358f57600080fd5b50813567ffffffffffffffff8111156135a757600080fd5b6020830191508360208285010111156135bf57600080fd5b9250929050565b6000806000604084860312156135db57600080fd5b833567ffffffffffffffff8111156135f257600080fd5b8401610120818703121561360557600080fd5b61360d613295565b61361682613308565b815261362460208301613319565b6020820152604082013567ffffffffffffffff81111561364357600080fd5b61364f888285016133db565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561368357600080fd5b61368f8882850161334d565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff8111156136c457600080fd5b6136d088828501613518565b61010083015250935050602084013567ffffffffffffffff8111156136f457600080fd5b6137008682870161357d565b9497909650939450505050565b60008060008060006080868803121561372557600080fd5b61372e866131f9565b945061373c602087016131f9565b935060408601359250606086013567ffffffffffffffff81111561375f57600080fd5b61376b8882890161357d565b969995985093965092949392505050565b60008060006040848603121561379157600080fd5b83359250602084013567ffffffffffffffff8111156136f457600080fd5b7fffffffff000000000000000000000000000000000000000000000000000000008116811461064a57600080fd5b6000602082840312156137ef57600080fd5b81356107b1816137af565b6000806000806040858703121561381057600080fd5b843567ffffffffffffffff81111561382757600080fd5b6138338782880161357d565b909550935050602085013567ffffffffffffffff81111561385357600080fd5b61385f8782880161357d565b95989497509550505050565b60006020828403121561387d57600080fd5b5035919050565b6000806020838503121561389757600080fd5b823567ffffffffffffffff8111156138ae57600080fd5b6138ba8582860161357d565b90969095509350505050565b600080600080606085870312156138dc57600080fd5b6138e5856131f9565b935060208501359250604085013567ffffffffffffffff81111561385357600080fd5b6000806040838503121561391b57600080fd5b8235613926816137af565b9150613934602084016131f9565b90509250929050565b60008083601f84011261394f57600080fd5b50813567ffffffffffffffff81111561396757600080fd5b6020830191508360208260051b85010111156135bf57600080fd5b60008060008060008060008060a0898b03121561399e57600080fd5b6139a7896131f9565b97506139b560208a016131f9565b9650604089013567ffffffffffffffff8111156139d157600080fd5b6139dd8b828c0161393d565b909750955050606089013567ffffffffffffffff8111156139fd57600080fd5b613a098b828c0161393d565b909550935050608089013567ffffffffffffffff811115613a2957600080fd5b613a358b828c0161357d565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613a6257600080fd5b613a6b876131f9565b9550613a79602088016131f9565b94506040870135935060608701359250608087013567ffffffffffffffff811115613aa357600080fd5b613aaf89828a0161357d565b979a9699509497509295939492505050565b600080600060608486031215613ad657600080fd5b83359250613ae6602085016131f9565b915060408401356bffffffffffffffffffffffff81168114613b0757600080fd5b809150509250925092565b808201808211156105fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613b96578181015183820152602001613b7e565b50506000910152565b60008151808452613bb7816020860160208601613b7b565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613cb9577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613c7560e0860182613b9f565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613c07565b50909695505050505050565b600081518084526020840193506020830160005b82811015613d0d57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613cd9565b5093949350505050565b805160ff16825260006020820151613d33602085018215159052565b5060408201516101206040850152613d4f610120850182613be9565b9050606083015160608501526080830151608085015260a083015184820360a0860152613d7c8282613b9f565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613dac8282613cc5565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613e116040830186613d17565b8281036020840152613e24818587613db5565b9695505050505050565b60008085851115613e3e57600080fd5b83861115613e4b57600080fd5b5050820193919092039150565b606081526000613e6b6060830186613d17565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000613eae60c0830184613b9f565b98975050505050505050565b838152826020820152606060408201526000613dac6060830184613b9f565b606081526000613eec6060830186613d17565b8460208401528281036040840152613e248185613b9f565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613f336040830184613b9f565b949350505050565b60006040828403128015613f4e57600080fd5b506040805190810167ffffffffffffffff81118282101715613f7257613f7261323d565b604052825181526020928301519281019290925250919050565b606081526000613fa0606083018688613db5565b6020830194909452506040015292915050565b838152604060208201526000613dac604083018486613db5565b600060208284031215613fdf57600080fd5b81516107b1816137af565b608081526000613ffd6080830188613d17565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152613eae818587613db5565b60006020828403121561404457600080fd5b5051919050565b60408152600061405e6040830185613b9f565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351614095818460208801613b7b565b919091019182525060200191905056fea2646970667358221220fb02bbd969bad3ebcf0e326a5f5858fa52e8d1770b545ee95c17d76337a0a2f164736f6c634300081b00336080604052348015600f57600080fd5b50613f598061001f6000396000f3fe60806040526004361061012d5760003560e01c80636ea44577116100a5578063ad55366b11610074578063bc197c8111610059578063bc197c8114610442578063f23a6e611461048a578063f727ef1c146104d057610134565b8063ad55366b146103dc578063b93ea7ad1461042f57610134565b80636ea44577146103735780638943ec02146103865780638c3f5563146103a7578063aaf10f42146103c757610134565b80631a9b2337116100fc57806329561426116100e1578063295614261461032b5780634fcf3eca1461034b57806351605d801461035e57610134565b80631a9b2337146102d35780631f6a1eb91461031857610134565b8063025b22bc146101f757806313792a4a1461020a578063150b7a021461023d5780631626ba7e146102b357610134565b3661013457005b600436106101f557600061015061014b3683613001565b6104f0565b905073ffffffffffffffffffffffffffffffffffffffff8116156101f3576000808273ffffffffffffffffffffffffffffffffffffffff16600036604051610199929190613067565b600060405180830381855af49150503d80600081146101d4576040519150601f19603f3d011682016040523d82523d6000602084013e6101d9565b606091505b5091509150816101eb57805160208201fd5b805160208201f35b505b005b6101f56102053660046130a0565b610544565b34801561021657600080fd5b5061022a610225366004613444565b610590565b6040519081526020015b60405180910390f35b34801561024957600080fd5b5061028261025836600461358b565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610234565b3480156102bf57600080fd5b506102826102ce3660046135fa565b6106fb565b3480156102df57600080fd5b506102f36102ee36600461365b565b610791565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610234565b6101f5610326366004613678565b61079c565b34801561033757600080fd5b506101f56103463660046136e9565b610825565b6101f561035936600461365b565b610869565b34801561036a57600080fd5b5061022a61092b565b6101f5610381366004613702565b61095a565b34801561039257600080fd5b506101f56103a1366004613744565b50505050565b3480156103b357600080fd5b5061022a6103c23660046136e9565b6109c7565b3480156103d357600080fd5b506102f36109f3565b3480156103e857600080fd5b506103fc6103f7366004613444565b6109fd565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610234565b6101f561043d366004613786565b610a37565b34801561044e57600080fd5b5061028261045d366004613800565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561049657600080fd5b506102826104a53660046138c7565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b3480156104dc57600080fd5b506101f56104eb36600461393f565b610afc565b600061053e7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610bb7565b92915050565b333014610584576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61058d81610c15565b50565b6000808461010001515160016105a69190613990565b67ffffffffffffffff8111156105be576105be6130bb565b6040519080825280602002602001820160405280156105e7578160200160208202803683370190505b50905060005b85610100015151811015610659578561010001518181518110610612576106126139ca565b602002602001015182828151811061062c5761062c6139ca565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016105ed565b5033818661010001515181518110610673576106736139ca565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006106ad868686610c6a565b509050806106ed578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161057b93929190613c7c565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610752828686610c6a565b509050806107665750600091506106f49050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b600061053e826104f0565b60005a905060006107ad8686610dfb565b90506107c181606001518260800151611219565b6000806107cf838787610c6a565b9150915081610810578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161057b93929190613c7c565b61081b848285611301565b5050505050505050565b333014610860576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b61058d81611661565b3330146108a4576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b60006108af826104f0565b73ffffffffffffffffffffffffffffffffffffffff1603610920576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008216600482015260240161057b565b61058d8160006116f1565b60006109557fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610995576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b60005a905060006109a68484610dfb565b905060006109b3826117b1565b90506109c0838284611301565b5050505050565b600061053e7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610bb7565b6000610955305490565b600080600080600080610a1489898960008061182c565b939950919750945092509050610a2983611b51565b935093975093979195509350565b333014610a72576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b6000610a7d836104f0565b73ffffffffffffffffffffffffffffffffffffffff1614610aee576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161057b565b610af882826116f1565b5050565b333014610b37576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b610b508383836bffffffffffffffffffffffff16611b5c565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610bd6929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c1d813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610c8257610c826139ca565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610d8157610cdf866117b1565b9150600080610ced84611beb565b9150915042811015610d35576040517f9fa4fe54000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161057b565b73ffffffffffffffffffffffffffffffffffffffff82161580610d6d575073ffffffffffffffffffffffffffffffffffffffff821633145b15610d7e5760019450505050610df3565b50505b6000806000610d9489898960008061182c565b985092955090935091505082821015610de3576040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161057b565b610dec81611b51565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610e5f5760006060840152610e70565b84810135606090811c908401526014015b6007600183901c168015610ec35760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610ed857506001610f01565b83602016602003610ef45750600282019186013560f01c610f01565b50600182019186013560f81c5b8067ffffffffffffffff811115610f1a57610f1a6130bb565b604051908082528060200260200182016040528015610fa557816020015b610f926040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610f385790505b50604086015260005b8181101561120e5760018085019489013560f81c90808216900361100d573087604001518381518110610fe357610fe36139ca565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611057565b8489013560601c601486018860400151848151811061102e5761102e6139ca565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b6002808216900361109557848901356020860188604001518481518110611080576110806139ca565b60200260200101516020018197508281525050505b6004808216900361112d57600385019489013560e81c89868a6110b88483613990565b926110c593929190613cac565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506040890151805185908110611110576111106139ca565b6020908102919091010151604001526111298187613990565b9550505b6008808216900361116b57848901356020860188604001518481518110611156576111566139ca565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061118b5761118b6139ca565b602002602001015160800190151590811515815250508060201660ff16602014876040015183815181106111c1576111c16139ca565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106111f6576111f66139ca565b602090810291909101015160c0015250600101610fae565b505050505092915050565b6000611224836109c7565b9050818114611270576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161057b565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b818110156116595760008460400151828151811061132c5761132c6139ca565b602002602001015190508315611345576000935061138e565b8060a001511561138e5760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611651565b606081015180158015906113a15750805a105b156113de5785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161057b93929190613cd6565b60008260800151156114b25782516114ab9083156113fc57836113fe565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161142996959493929190613cfb565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611c37565b90506114d9565b825160208401516114d6919084156114ca57846114cc565b5a5b8660400151611c4d565b90505b806116145760c083015161153557600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d8885611516611c65565b60405161152593929190613d38565b60405180910390a1505050611651565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161159f57868461156a611c65565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161057b93929190613d57565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611614577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856115f5611c65565b60405161160493929190613d38565b60405180910390a1505050611659565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b60010161130c565b505050505050565b80611698576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116c17fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610c5f565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806117c2836020015130611c84565b905060006117cf84611d51565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611877575073ffffffffffffffffffffffffffffffffffffffff8916155b15611993578b82013560601c9850601490910190896119935760038201918c013560e81c60008d848e6118aa8583613990565b926118b793929190613cac565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc891506119429030908590600401613d82565b6040805180830381865afa15801561195e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119829190613db9565b925061198e8285613990565b935050505b826001166001036119cd576119bb8d8a838f8f879080926119b693929190613cac565b611fa7565b97509750975097509750505050611b44565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611a386001600586901c811690613990565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611a838d6117b1565b9350611aa18d858e8e86908092611a9c93929190613cac565b6121f3565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611aec575080518614155b8015611afc575080602001518511155b15611b40576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161057b565b5050505b9550955095509550959050565b600061053e82612bd7565b611be67fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611c197fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610bb7565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611cf45746611cf7565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611d62836101000151612c0a565b835190915060ff16611dd5576000611d7d8460400151612c8d565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c00161180d565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611e655760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611ed55760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001611e47565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01611f455760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201529081019190915260608101829052608001611e47565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e6400000000000000000000000000000000604482015260640161057b565b6000806000806000612009604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156121805760038201916000908b013560e81c6120518482613990565b9150600090508a8214612065576000612067565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8414612098578561209a565b8f5b90506120ba818e8e889087926120b293929190613cac565b60018661182c565b50929d50909b50995097508a8a1015612117578c8c869085926120df93929190613cac565b8c8c6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161057b9493929190613e0a565b829450888e600001510361212a5760008e525b83881061216d576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018990526024810185905260440161057b565b505060c084018790529150849050612031565b8a511580159061219457508a602001518511155b156121d8576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161057b565b6121e18d6117b1565b93505050509550955095509550959050565b60008060005b83811015612bcd57600181019085013560f881901c9060fc1c8061231657600f821660008190036122315750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156122c6573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006122e9828960ff16612d0f565b90508b6122f65780612305565b60008c81526020829052604090205b9b50505050505050505050506121f9565b6001810361237a57600f821660008190036123385750600183019287013560f81c5b601484019388013560601c60006123528260ff8516612d0f565b90508661235f578061236e565b60008781526020829052604090205b965050505050506121f9565b600281036125615760038216600081900361239c5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261245a93929190613cac565b6040518463ffffffff1660e01b815260040161247893929190613e31565b602060405180830381865afa158015612495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124b99190613e4b565b7fffffffff00000000000000000000000000000000000000000000000000000000161461251c578d8d858e8e6040517ff734863a00000000000000000000000000000000000000000000000000000000815260040161057b959493929190613e68565b8097508460ff168a0199506000612536858760ff16612d0f565b9050896125435780612552565b60008a81526020829052604090205b995050505050505050506121f9565b600381036125955760208301928701358461257c578061258b565b60008581526020829052604090205b94505050506121f9565b6004810361263b57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c1685830180965081925050506000818601905060008061260a8e8e8e8e8c908892611a9c93929190613cac565b91509150829750818a01995061262a898260009182526020526040902090565b9850829750505050505050506121f9565b60068103612705576003600283901c1660008190036126615750600183019287013560f81c5b60038316600081900361267b5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806126b98f8f8f8f8d908892611a9c93929190613cac565b915091508298508482106126cc57998501995b60006126d9828789612d76565b90508a6126e657806126f5565b60008b81526020829052604090205b9a505050505050505050506121f9565b60058103612772576020830192870135888103612740577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061274b82612dda565b9050856127585780612767565b60008681526020829052604090205b9550505050506121f9565b6007810361287857600f821660008190036127945750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016122a4565b600881036128cc57602083019287013560006128948b82612e2e565b90508082036128c1577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061235282612ea9565b60098103612a32576003821660008190036128ee5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c90879261298893929190613cac565b6040518463ffffffff1660e01b81526004016129a693929190613c7c565b602060405180830381865afa1580156129c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129e79190613eb0565b90508197508460ff168a0199506000612a04858760ff1684612ee4565b905089612a115780612a20565b60008a81526020829052604090205b995082985050505050505050506121f9565b600a8103612b9857600382166000819003612a545750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612aed93929190613cac565b6040518463ffffffff1660e01b8152600401612b0b93929190613e31565b602060405180830381865afa158015612b28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b4c9190613eb0565b90508198508560ff168b019a506000612b69868860ff1684612ee4565b90508a612b765780612b85565b60008b81526020829052604090205b9a508299505050505050505050506121f9565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161057b565b5094509492505050565b6000811580159061053e5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b6000606060005b8351811015612c7e5781848281518110612c2d57612c2d6139ca565b6020026020010151604051602001612c46929190613ec9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612c11565b50805160209091012092915050565b6000606060005b8351811015612c7e576000612cc1858381518110612cb457612cb46139ca565b6020026020010151612f52565b90508281604051602001612cd6929190613f01565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612c94565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611d33565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612e3f846020015184611c84565b90506000612e4c85611d51565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612e11565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612dbb565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612e1198979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015613060577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461309b57600080fd5b919050565b6000602082840312156130b257600080fd5b6106f482613077565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561310d5761310d6130bb565b60405290565b604051610120810167ffffffffffffffff8111828210171561310d5761310d6130bb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561317e5761317e6130bb565b604052919050565b803560ff8116811461309b57600080fd5b8035801515811461309b57600080fd5b600067ffffffffffffffff8211156131c1576131c16130bb565b5060051b60200190565b600082601f8301126131dc57600080fd5b813567ffffffffffffffff8111156131f6576131f66130bb565b61322760207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613137565b81815284602083860101111561323c57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261326a57600080fd5b813561327d613278826131a7565b613137565b8082825260208201915060208360051b86010192508583111561329f57600080fd5b602085015b8381101561338c57803567ffffffffffffffff8111156132c357600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe00112156132f757600080fd5b6132ff6130ea565b61330b60208301613077565b815260408201356020820152606082013567ffffffffffffffff81111561333157600080fd5b6133408a6020838601016131cb565b6040830152506080820135606082015261335c60a08301613197565b608082015261336d60c08301613197565b60a082015260e0919091013560c08201528352602092830192016132a4565b5095945050505050565b600082601f8301126133a757600080fd5b81356133b5613278826131a7565b8082825260208201915060208360051b8601019250858311156133d757600080fd5b602085015b8381101561338c576133ed81613077565b8352602092830192016133dc565b60008083601f84011261340d57600080fd5b50813567ffffffffffffffff81111561342557600080fd5b60208301915083602082850101111561343d57600080fd5b9250929050565b60008060006040848603121561345957600080fd5b833567ffffffffffffffff81111561347057600080fd5b8401610120818703121561348357600080fd5b61348b613113565b61349482613186565b81526134a260208301613197565b6020820152604082013567ffffffffffffffff8111156134c157600080fd5b6134cd88828501613259565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561350157600080fd5b61350d888285016131cb565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561354257600080fd5b61354e88828501613396565b61010083015250935050602084013567ffffffffffffffff81111561357257600080fd5b61357e868287016133fb565b9497909650939450505050565b6000806000806000608086880312156135a357600080fd5b6135ac86613077565b94506135ba60208701613077565b935060408601359250606086013567ffffffffffffffff8111156135dd57600080fd5b6135e9888289016133fb565b969995985093965092949392505050565b60008060006040848603121561360f57600080fd5b83359250602084013567ffffffffffffffff81111561357257600080fd5b7fffffffff000000000000000000000000000000000000000000000000000000008116811461058d57600080fd5b60006020828403121561366d57600080fd5b81356106f48161362d565b6000806000806040858703121561368e57600080fd5b843567ffffffffffffffff8111156136a557600080fd5b6136b1878288016133fb565b909550935050602085013567ffffffffffffffff8111156136d157600080fd5b6136dd878288016133fb565b95989497509550505050565b6000602082840312156136fb57600080fd5b5035919050565b6000806020838503121561371557600080fd5b823567ffffffffffffffff81111561372c57600080fd5b613738858286016133fb565b90969095509350505050565b6000806000806060858703121561375a57600080fd5b61376385613077565b935060208501359250604085013567ffffffffffffffff8111156136d157600080fd5b6000806040838503121561379957600080fd5b82356137a48161362d565b91506137b260208401613077565b90509250929050565b60008083601f8401126137cd57600080fd5b50813567ffffffffffffffff8111156137e557600080fd5b6020830191508360208260051b850101111561343d57600080fd5b60008060008060008060008060a0898b03121561381c57600080fd5b61382589613077565b975061383360208a01613077565b9650604089013567ffffffffffffffff81111561384f57600080fd5b61385b8b828c016137bb565b909750955050606089013567ffffffffffffffff81111561387b57600080fd5b6138878b828c016137bb565b909550935050608089013567ffffffffffffffff8111156138a757600080fd5b6138b38b828c016133fb565b999c989b5096995094979396929594505050565b60008060008060008060a087890312156138e057600080fd5b6138e987613077565b95506138f760208801613077565b94506040870135935060608701359250608087013567ffffffffffffffff81111561392157600080fd5b61392d89828a016133fb565b979a9699509497509295939492505050565b60008060006060848603121561395457600080fd5b8335925061396460208501613077565b915060408401356bffffffffffffffffffffffff8116811461398557600080fd5b809150509250925092565b8082018082111561053e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613a145781810151838201526020016139fc565b50506000910152565b60008151808452613a358160208601602086016139f9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613b37577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613af360e0860182613a1d565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613a85565b50909695505050505050565b600081518084526020840193506020830160005b82811015613b8b57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613b57565b5093949350505050565b805160ff16825260006020820151613bb1602085018215159052565b5060408201516101206040850152613bcd610120850182613a67565b9050606083015160608501526080830151608085015260a083015184820360a0860152613bfa8282613a1d565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613c2a8282613b43565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613c8f6040830186613b95565b8281036020840152613ca2818587613c33565b9695505050505050565b60008085851115613cbc57600080fd5b83861115613cc957600080fd5b5050820193919092039150565b606081526000613ce96060830186613b95565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000613d2c60c0830184613a1d565b98975050505050505050565b838152826020820152606060408201526000613c2a6060830184613a1d565b606081526000613d6a6060830186613b95565b8460208401528281036040840152613ca28185613a1d565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613db16040830184613a1d565b949350505050565b60006040828403128015613dcc57600080fd5b506040805190810167ffffffffffffffff81118282101715613df057613df06130bb565b604052825181526020928301519281019290925250919050565b606081526000613e1e606083018688613c33565b6020830194909452506040015292915050565b838152604060208201526000613c2a604083018486613c33565b600060208284031215613e5d57600080fd5b81516106f48161362d565b608081526000613e7b6080830188613b95565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152613d2c818587613c33565b600060208284031215613ec257600080fd5b5051919050565b604081526000613edc6040830185613a1d565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351613f138184602088016139f9565b919091019182525060200191905056fea26469706673582212209259e64891bb73bed6a64e6f258360a4ee8aa14e6ef72690b4f3e4e96b2cf43f64736f6c634300081b0033603e600e3d39601e805130553df33d3d34601c57363d3d373d363d30545af43d82803e903d91601c57fd5bf3",
}

// WalletStage1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage1MetaData.ABI instead.
var WalletStage1ABI = WalletStage1MetaData.ABI

// WalletStage1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage1MetaData.Bin instead.
var WalletStage1Bin = WalletStage1MetaData.Bin

// DeployWalletStage1 deploys a new Ethereum contract, binding an instance of WalletStage1 to it.
func DeployWalletStage1(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *WalletStage1, error) {
	parsed, err := WalletStage1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage1Bin), backend, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletStage1{WalletStage1Caller: WalletStage1Caller{contract: contract}, WalletStage1Transactor: WalletStage1Transactor{contract: contract}, WalletStage1Filterer: WalletStage1Filterer{contract: contract}}, nil
}

// WalletStage1 is an auto generated Go binding around an Ethereum contract.
type WalletStage1 struct {
	WalletStage1Caller     // Read-only binding to the contract
	WalletStage1Transactor // Write-only binding to the contract
	WalletStage1Filterer   // Log filterer for contract events
}

// WalletStage1Caller is an auto generated read-only Go binding around an Ethereum contract.
type WalletStage1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletStage1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletStage1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletStage1Session struct {
	Contract     *WalletStage1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletStage1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletStage1CallerSession struct {
	Contract *WalletStage1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WalletStage1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletStage1TransactorSession struct {
	Contract     *WalletStage1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WalletStage1Raw is an auto generated low-level Go binding around an Ethereum contract.
type WalletStage1Raw struct {
	Contract *WalletStage1 // Generic contract binding to access the raw methods on
}

// WalletStage1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletStage1CallerRaw struct {
	Contract *WalletStage1Caller // Generic read-only contract binding to access the raw methods on
}

// WalletStage1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletStage1TransactorRaw struct {
	Contract *WalletStage1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletStage1 creates a new instance of WalletStage1, bound to a specific deployed contract.
func NewWalletStage1(address common.Address, backend bind.ContractBackend) (*WalletStage1, error) {
	contract, err := bindWalletStage1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletStage1{WalletStage1Caller: WalletStage1Caller{contract: contract}, WalletStage1Transactor: WalletStage1Transactor{contract: contract}, WalletStage1Filterer: WalletStage1Filterer{contract: contract}}, nil
}

// NewWalletStage1Caller creates a new read-only instance of WalletStage1, bound to a specific deployed contract.
func NewWalletStage1Caller(address common.Address, caller bind.ContractCaller) (*WalletStage1Caller, error) {
	contract, err := bindWalletStage1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage1Caller{contract: contract}, nil
}

// NewWalletStage1Transactor creates a new write-only instance of WalletStage1, bound to a specific deployed contract.
func NewWalletStage1Transactor(address common.Address, transactor bind.ContractTransactor) (*WalletStage1Transactor, error) {
	contract, err := bindWalletStage1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage1Transactor{contract: contract}, nil
}

// NewWalletStage1Filterer creates a new log filterer instance of WalletStage1, bound to a specific deployed contract.
func NewWalletStage1Filterer(address common.Address, filterer bind.ContractFilterer) (*WalletStage1Filterer, error) {
	contract, err := bindWalletStage1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletStage1Filterer{contract: contract}, nil
}

// bindWalletStage1 binds a generic wrapper to an already deployed contract.
func bindWalletStage1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletStage1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage1 *WalletStage1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage1.Contract.WalletStage1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage1 *WalletStage1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1.Contract.WalletStage1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage1 *WalletStage1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage1.Contract.WalletStage1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage1 *WalletStage1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage1 *WalletStage1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage1 *WalletStage1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage1.Contract.contract.Transact(opts, method, params...)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1 *WalletStage1Caller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1 *WalletStage1Session) FACTORY() (common.Address, error) {
	return _WalletStage1.Contract.FACTORY(&_WalletStage1.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) FACTORY() (common.Address, error) {
	return _WalletStage1.Contract.FACTORY(&_WalletStage1.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1 *WalletStage1Caller) INITCODEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "INIT_CODE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1 *WalletStage1Session) INITCODEHASH() ([32]byte, error) {
	return _WalletStage1.Contract.INITCODEHASH(&_WalletStage1.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1 *WalletStage1CallerSession) INITCODEHASH() ([32]byte, error) {
	return _WalletStage1.Contract.INITCODEHASH(&_WalletStage1.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1 *WalletStage1Caller) STAGE2IMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "STAGE_2_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1 *WalletStage1Session) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletStage1.Contract.STAGE2IMPLEMENTATION(&_WalletStage1.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletStage1.Contract.STAGE2IMPLEMENTATION(&_WalletStage1.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1 *WalletStage1Caller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1 *WalletStage1Session) GetImplementation() (common.Address, error) {
	return _WalletStage1.Contract.GetImplementation(&_WalletStage1.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) GetImplementation() (common.Address, error) {
	return _WalletStage1.Contract.GetImplementation(&_WalletStage1.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1 *WalletStage1Caller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1 *WalletStage1Session) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage1.Contract.IsValidSignature(&_WalletStage1.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1 *WalletStage1CallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage1.Contract.IsValidSignature(&_WalletStage1.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Caller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC1155BatchReceived(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1CallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC1155BatchReceived(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Caller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC1155Received(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1CallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC1155Received(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC721Received(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.OnERC721Received(&_WalletStage1.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1Caller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1Session) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, signature)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1 *WalletStage1Caller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1 *WalletStage1Session) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage1.Contract.ReadNonce(&_WalletStage1.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1 *WalletStage1CallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage1.Contract.ReadNonce(&_WalletStage1.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage1 *WalletStage1Caller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletStage1 *WalletStage1Session) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage1.Contract.RecoverPartialSignature(&_WalletStage1.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage1 *WalletStage1CallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage1.Contract.RecoverPartialSignature(&_WalletStage1.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1 *WalletStage1Caller) RecoverSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "recoverSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1 *WalletStage1Session) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage1.Contract.RecoverSapientSignature(&_WalletStage1.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1 *WalletStage1CallerSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage1.Contract.RecoverSapientSignature(&_WalletStage1.CallOpts, _payload, _signature)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Transactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Session) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, signature, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1 *WalletStage1Transactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1 *WalletStage1Session) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.Execute(&_WalletStage1.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.Execute(&_WalletStage1.TransactOpts, _payload, _signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1Transactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1Session) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1 *WalletStage1Transactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1 *WalletStage1Session) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.SelfExecute(&_WalletStage1.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.SelfExecute(&_WalletStage1.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1 *WalletStage1Transactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1 *WalletStage1Session) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1.Contract.SetStaticSignature(&_WalletStage1.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1 *WalletStage1TransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1.Contract.SetStaticSignature(&_WalletStage1.TransactOpts, _hash, _address, _timestamp)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1Transactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1Session) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1TransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.TransactOpts, arg0, arg1, arg2)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1 *WalletStage1Transactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1 *WalletStage1Session) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.UpdateImageHash(&_WalletStage1.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1 *WalletStage1TransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.UpdateImageHash(&_WalletStage1.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1 *WalletStage1Transactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1 *WalletStage1Session) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.UpdateImplementation(&_WalletStage1.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.UpdateImplementation(&_WalletStage1.TransactOpts, _implementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1 *WalletStage1Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1 *WalletStage1Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.Fallback(&_WalletStage1.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.Fallback(&_WalletStage1.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1 *WalletStage1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1 *WalletStage1Session) Receive() (*types.Transaction, error) {
	return _WalletStage1.Contract.Receive(&_WalletStage1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) Receive() (*types.Transaction, error) {
	return _WalletStage1.Contract.Receive(&_WalletStage1.TransactOpts)
}

// WalletStage1CallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletStage1 contract.
type WalletStage1CallAbortedIterator struct {
	Event *WalletStage1CallAborted // Event containing the contract specifics and raw log

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
func (it *WalletStage1CallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1CallAborted)
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
		it.Event = new(WalletStage1CallAborted)
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
func (it *WalletStage1CallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1CallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1CallAborted represents a CallAborted event raised by the WalletStage1 contract.
type WalletStage1CallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletStage1CallAbortedIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallAbortedIterator{contract: _WalletStage1.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage1CallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1CallAborted)
				if err := _WalletStage1.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseCallAborted(log types.Log) (*WalletStage1CallAborted, error) {
	event := new(WalletStage1CallAborted)
	if err := _WalletStage1.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1CallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletStage1 contract.
type WalletStage1CallFailedIterator struct {
	Event *WalletStage1CallFailed // Event containing the contract specifics and raw log

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
func (it *WalletStage1CallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1CallFailed)
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
		it.Event = new(WalletStage1CallFailed)
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
func (it *WalletStage1CallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1CallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1CallFailed represents a CallFailed event raised by the WalletStage1 contract.
type WalletStage1CallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletStage1CallFailedIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallFailedIterator{contract: _WalletStage1.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage1CallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1CallFailed)
				if err := _WalletStage1.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseCallFailed(log types.Log) (*WalletStage1CallFailed, error) {
	event := new(WalletStage1CallFailed)
	if err := _WalletStage1.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1CallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletStage1 contract.
type WalletStage1CallSkippedIterator struct {
	Event *WalletStage1CallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletStage1CallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1CallSkipped)
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
		it.Event = new(WalletStage1CallSkipped)
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
func (it *WalletStage1CallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1CallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1CallSkipped represents a CallSkipped event raised by the WalletStage1 contract.
type WalletStage1CallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletStage1CallSkippedIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallSkippedIterator{contract: _WalletStage1.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage1CallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1CallSkipped)
				if err := _WalletStage1.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseCallSkipped(log types.Log) (*WalletStage1CallSkipped, error) {
	event := new(WalletStage1CallSkipped)
	if err := _WalletStage1.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1CallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletStage1 contract.
type WalletStage1CallSucceededIterator struct {
	Event *WalletStage1CallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletStage1CallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1CallSucceeded)
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
		it.Event = new(WalletStage1CallSucceeded)
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
func (it *WalletStage1CallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1CallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1CallSucceeded represents a CallSucceeded event raised by the WalletStage1 contract.
type WalletStage1CallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) FilterCallSucceeded(opts *bind.FilterOpts) (*WalletStage1CallSucceededIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallSucceededIterator{contract: _WalletStage1.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletStage1CallSucceeded) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1CallSucceeded)
				if err := _WalletStage1.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseCallSucceeded(log types.Log) (*WalletStage1CallSucceeded, error) {
	event := new(WalletStage1CallSucceeded)
	if err := _WalletStage1.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1DefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletStage1 contract.
type WalletStage1DefinedHookIterator struct {
	Event *WalletStage1DefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletStage1DefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1DefinedHook)
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
		it.Event = new(WalletStage1DefinedHook)
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
func (it *WalletStage1DefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1DefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1DefinedHook represents a DefinedHook event raised by the WalletStage1 contract.
type WalletStage1DefinedHook struct {
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1 *WalletStage1Filterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage1DefinedHookIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage1DefinedHookIterator{contract: _WalletStage1.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1 *WalletStage1Filterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletStage1DefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1DefinedHook)
				if err := _WalletStage1.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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

// ParseDefinedHook is a log parse operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1 *WalletStage1Filterer) ParseDefinedHook(log types.Log) (*WalletStage1DefinedHook, error) {
	event := new(WalletStage1DefinedHook)
	if err := _WalletStage1.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1ImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletStage1 contract.
type WalletStage1ImageHashUpdatedIterator struct {
	Event *WalletStage1ImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage1ImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1ImageHashUpdated)
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
		it.Event = new(WalletStage1ImageHashUpdated)
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
func (it *WalletStage1ImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1ImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1ImageHashUpdated represents a ImageHashUpdated event raised by the WalletStage1 contract.
type WalletStage1ImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage1 *WalletStage1Filterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletStage1ImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage1ImageHashUpdatedIterator{contract: _WalletStage1.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage1 *WalletStage1Filterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage1ImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1ImageHashUpdated)
				if err := _WalletStage1.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseImageHashUpdated(log types.Log) (*WalletStage1ImageHashUpdated, error) {
	event := new(WalletStage1ImageHashUpdated)
	if err := _WalletStage1.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1ImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletStage1 contract.
type WalletStage1ImplementationUpdatedIterator struct {
	Event *WalletStage1ImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage1ImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1ImplementationUpdated)
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
		it.Event = new(WalletStage1ImplementationUpdated)
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
func (it *WalletStage1ImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1ImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1ImplementationUpdated represents a ImplementationUpdated event raised by the WalletStage1 contract.
type WalletStage1ImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage1 *WalletStage1Filterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletStage1ImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage1ImplementationUpdatedIterator{contract: _WalletStage1.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage1 *WalletStage1Filterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage1ImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1ImplementationUpdated)
				if err := _WalletStage1.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseImplementationUpdated(log types.Log) (*WalletStage1ImplementationUpdated, error) {
	event := new(WalletStage1ImplementationUpdated)
	if err := _WalletStage1.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1NonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletStage1 contract.
type WalletStage1NonceChangeIterator struct {
	Event *WalletStage1NonceChange // Event containing the contract specifics and raw log

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
func (it *WalletStage1NonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1NonceChange)
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
		it.Event = new(WalletStage1NonceChange)
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
func (it *WalletStage1NonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1NonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1NonceChange represents a NonceChange event raised by the WalletStage1 contract.
type WalletStage1NonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage1 *WalletStage1Filterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletStage1NonceChangeIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletStage1NonceChangeIterator{contract: _WalletStage1.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage1 *WalletStage1Filterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletStage1NonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1NonceChange)
				if err := _WalletStage1.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseNonceChange(log types.Log) (*WalletStage1NonceChange, error) {
	event := new(WalletStage1NonceChange)
	if err := _WalletStage1.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1StaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletStage1 contract.
type WalletStage1StaticSignatureSetIterator struct {
	Event *WalletStage1StaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletStage1StaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1StaticSignatureSet)
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
		it.Event = new(WalletStage1StaticSignatureSet)
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
func (it *WalletStage1StaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1StaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1StaticSignatureSet represents a StaticSignatureSet event raised by the WalletStage1 contract.
type WalletStage1StaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage1 *WalletStage1Filterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletStage1StaticSignatureSetIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletStage1StaticSignatureSetIterator{contract: _WalletStage1.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage1 *WalletStage1Filterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletStage1StaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1StaticSignatureSet)
				if err := _WalletStage1.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletStage1 *WalletStage1Filterer) ParseStaticSignatureSet(log types.Log) (*WalletStage1StaticSignatureSet, error) {
	event := new(WalletStage1StaticSignatureSet)
	if err := _WalletStage1.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
