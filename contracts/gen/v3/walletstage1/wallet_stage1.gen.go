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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidERC1271Signature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_expectedCaller\",\"type\":\"address\"}],\"name\":\"InvalidStaticSignatureWrongCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAGE_2_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getStaticSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161abaf38038061abaf83398181016040528101906100319190610196565b8060405161003e9061012b565b604051809103905ff080158015610057573d5f5f3e3d5ffd5b505f6040518060600160405280602c815260200161ab83602c91393073ffffffffffffffffffffffffffffffffffffffff1660405160200161009a92919061023c565b60405160208183030381529060405280519060200120905080608081815250508273ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff168152505050505050610263565b6153a4806157df83390190565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101658261013c565b9050919050565b6101758161015b565b811461017f575f5ffd5b50565b5f815190506101908161016c565b92915050565b5f602082840312156101ab576101aa610138565b5b5f6101b884828501610182565b91505092915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6101ed826101c1565b6101f781856101cb565b93506102078185602086016101d5565b80840191505092915050565b5f819050919050565b5f819050919050565b61023661023182610213565b61021c565b82525050565b5f61024782856101e3565b91506102538284610225565b6020820191508190509392505050565b60805160a05160c05161553d6102a25f395f8181610bfe0152611a0401525f81816109f401526131df01525f81816109540152613201015261553d5ff3fe60806040526004361061012d575f3560e01c80636ea44577116100aa578063aaf10f421161006e578063aaf10f42146104b5578063ad55366b146104df578063b93ea7ad14610520578063bc197c811461053c578063f23a6e6114610578578063f727ef1c146105b457610134565b80636ea44577146103ce5780638943ec02146103ea5780638c3f55631461041257806392dcb3fc1461044e5780639f69ef541461048b57610134565b80631f6a1eb9116100f15780631f6a1eb91461031a578063257671f51461033657806329561426146103605780632dd31000146103885780634fcf3eca146103b257610134565b8063025b22bc1461020e57806313792a4a1461022a578063150b7a02146102665780631626ba7e146102a25780631a9b2337146102de57610134565b3661013457005b60045f3690501061020c575f6101555f369061015091906135b7565b6105dc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461020a575f5f8273ffffffffffffffffffffffffffffffffffffffff165f366040516101b3929190613651565b5f60405180830381855af49150503d805f81146101eb576040519150601f19603f3d011682016040523d82523d5f602084013e6101f0565b606091505b50915091508161020257805160208201fd5b805160208201f35b505b005b610228600480360381019061022391906136d4565b610631565b005b348015610235575f5ffd5b50610250600480360381019061024b9190613cff565b6106ad565b60405161025d9190613d87565b60405180910390f35b348015610271575f5ffd5b5061028c60048036038101906102879190613da0565b61085c565b6040516102999190613e33565b60405180910390f35b3480156102ad575f5ffd5b506102c860048036038101906102c39190613e4c565b610870565b6040516102d59190613e33565b60405180910390f35b3480156102e9575f5ffd5b5061030460048036038101906102ff9190613ed3565b6108b2565b6040516103119190613f0d565b60405180910390f35b610334600480360381019061032f9190613f26565b6108c3565b005b348015610341575f5ffd5b5061034a610952565b6040516103579190613d87565b60405180910390f35b34801561036b575f5ffd5b5061038660048036038101906103819190613fa4565b610976565b005b348015610393575f5ffd5b5061039c6109f2565b6040516103a99190613f0d565b60405180910390f35b6103cc60048036038101906103c79190613ed3565b610a16565b005b6103e860048036038101906103e39190613fcf565b610b0b565b005b3480156103f5575f5ffd5b50610410600480360381019061040b919061401a565b610baa565b005b34801561041d575f5ffd5b506104386004803603810190610433919061408b565b610bb0565b60405161044591906140c5565b60405180910390f35b348015610459575f5ffd5b50610474600480360381019061046f9190613fa4565b610be8565b6040516104829291906140de565b60405180910390f35b348015610496575f5ffd5b5061049f610bfc565b6040516104ac9190613f0d565b60405180910390f35b3480156104c0575f5ffd5b506104c9610c20565b6040516104d69190613f0d565b60405180910390f35b3480156104ea575f5ffd5b5061050560048036038101906105009190613cff565b610c2e565b60405161051796959493929190614114565b60405180910390f35b61053a60048036038101906105359190614173565b610c6c565b005b348015610547575f5ffd5b50610562600480360381019061055d9190614206565b610d62565b60405161056f9190613e33565b60405180910390f35b348015610583575f5ffd5b5061059e600480360381019061059991906142dd565b610d79565b6040516105ab9190613e33565b60405180910390f35b3480156105bf575f5ffd5b506105da60048036038101906105d591906143b4565b610d8e565b005b5f6106287fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610e57565b5f1c9050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106a157336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016106989190613f0d565b60405180910390fd5b6106aa81610e8f565b50565b5f5f6001856101000151516106c29190614431565b67ffffffffffffffff8111156106db576106da613713565b5b6040519080825280602002602001820160405280156107095781602001602082028036833780820191505090505b5090505f5f90505b8561010001515181101561079957856101000151818151811061073757610736614464565b5b602002602001015182828151811061075257610751614464565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610711565b50338186610100015151815181106107b4576107b3614464565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808561010001819052505f610804868686610ed2565b5090508061084d578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161084493929190614844565b60405180910390fd5b60015f1b925050509392505050565b5f63150b7a0260e01b905095945050505050565b5f5f61087b856110c2565b90505f610889828686610ed2565b5090508061089e575f60e01b925050506108ab565b6320c13b0b60e01b925050505b9392505050565b5f6108bc826105dc565b9050919050565b5f5a90505f6108d286866110eb565b90506108e68160600151826080015161158d565b5f5f6108f3838787610ed2565b915091508161093d578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161093493929190614844565b60405180910390fd5b610948848285611631565b5050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109e657336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016109dd9190613f0d565b60405180910390fd5b6109ef81611961565b50565b7f000000000000000000000000000000000000000000000000000000000000000081565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a8657336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610a7d9190613f0d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610aa6826105dc565b73ffffffffffffffffffffffffffffffffffffffff1603610afe57806040517f1c3812cc000000000000000000000000000000000000000000000000000000008152600401610af59190613e33565b60405180910390fd5b610b08815f611a2b565b50565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b7b57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610b729190613f0d565b60405180910390fd5b5f5a90505f610b8a84846110eb565b90505f610b9682611acc565b9050610ba3838284611631565b5050505050565b50505050565b5f610bdf7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b610e57565b5f1c9050919050565b5f5f610bf383611b1c565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f610c29611b6d565b905090565b5f5f5f5f5f5f610c418989895f5f611b75565b809550819650829750839950849a505050505050610c5e83611eb2565b935093975093979195509350565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cdc57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610cd39190613f0d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610cfc836105dc565b73ffffffffffffffffffffffffffffffffffffffff1614610d5457816040517f5b4d6d6a000000000000000000000000000000000000000000000000000000008152600401610d4b9190613e33565b60405180910390fd5b610d5e8282611a2b565b5050565b5f63bc197c8160e01b905098975050505050505050565b5f63f23a6e6160e01b90509695505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610dfe57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610df59190613f0d565b60405180910390fd5b610e178383836bffffffffffffffffffffffff16611ec3565b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1838383604051610e4a9392919061488a565b60405180910390a1505050565b5f5f8383604051602001610e6c9291906148bf565b604051602081830303815290604052805190602001209050805491505092915050565b610e9881611f21565b7f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0381604051610ec79190613f0d565b60405180910390a150565b5f5f5f84845f818110610ee857610ee7614464565b5b9050013560f81c60f81b9050608060f81b608060f81b82167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19160361104157610f3086611acc565b91505f5f610f3d84611b1c565b91509150428111610f875783816040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600401610f7e9291906148e6565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610fef57503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611035578333836040517f8945c31300000000000000000000000000000000000000000000000000000000815260040161102c9392919061490d565b60405180910390fd5b600194505050506110ba565b5f5f5f6110518989895f5f611b75565b905080985081945082955083965050505050828210156110aa5782826040517ffd41fcba0000000000000000000000000000000000000000000000000000000081526004016110a1929190614942565b60405180910390fd5b6110b381611eb2565b9550505050505b935093915050565b6110ca6134c2565b6003815f019060ff16908160ff1681525050818160e0018181525050919050565b6110f36134c2565b5f815f019060ff16908160ff16815250505f5f6111108585611f27565b915060ff169150600180831603611130575f83606001818152505061116c565b611145818686611f3d9290919263ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff169150846060018193508281525050505b5f6007600184901c1690505f8111156111a95761119b82828888611f6e9190939291909392919063ffffffff16565b856080018194508281525050505b5f6010808516036111bd5760019050611215565b6020808516036111f0576111dc838888611f9b9290919263ffffffff16565b8161ffff1691508094508192505050611214565b611205838888611fba9290919263ffffffff16565b8160ff16915080945081925050505b5b8067ffffffffffffffff81111561122f5761122e613713565b5b60405190808252806020026020018201604052801561126857816020015b61125561350d565b81526020019060019003908161124d5790505b5085604001819052505f5f90505b81811015611582575f611294858a8a611fba9290919263ffffffff16565b8096508192505050600180821660ff16036113035730876040015183815181106112c1576112c0614464565b5b60200260200101515f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061136f565b611318858a8a611fd59290919263ffffffff16565b8860400151848151811061132f5761132e614464565b5b60200260200101515f018197508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b600280821660ff16036113bd57611391858a8a6120069290919263ffffffff16565b886040015184815181106113a8576113a7614464565b5b60200260200101516020018197508281525050505b600480821660ff1603611485575f6113e0868b8b61201c9290919263ffffffff16565b8162ffffff16915080975081925050508989879083896114009190614431565b9261140d93929190614971565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508860400151848151811061146657611465614464565b5b60200260200101516040018190525080866114819190614431565b9550505b600880821660ff16036114d3576114a7858a8a6120069290919263ffffffff16565b886040015184815181106114be576114bd614464565b5b60200260200101516060018197508281525050505b601080821660ff1614876040015183815181106114f3576114f2614464565b5b60200260200101516080019015159081151581525050602080821660ff16148760400151838151811061152957611528614464565b5b602002602001015160a0019015159081151581525050600660c0821660ff16901c60ff168760400151838151811061156457611563614464565b5b602002602001015160c0018181525050508080600101915050611276565b505050505092915050565b5f61159783610bb0565b90508181146115e1578282826040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004016115d8939291906149ab565b60405180910390fd5b5f6001830190506115f2848261203c565b7f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8818482604051611623929190614942565b60405180910390a150505050565b5f5f90505f82604001515190505f5f90505b81811015611959575f8460400151828151811061166357611662614464565b5b602002602001015190508060a00151801561167c575083155b156116c0577f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b86836040516116b29291906148e6565b60405180910390a15061194c565b5f93505f816060015190505f81141580156116da5750805a105b156117205785835a6040517f21395274000000000000000000000000000000000000000000000000000000008152600401611717939291906149e0565b60405180910390fd5b5f8260800151156117d5576117ce835f01515f841461173f5783611741565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161176c96959493929190614a54565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612071565b90506117fd565b6117fa835f015184602001515f85146117ee57846117f0565b5a5b8660400151612086565b90505b8061190f575f60ff168360c001510361185e57600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d888561183f61209d565b60405161184e93929190614aba565b60405180910390a150505061194c565b600160ff168360c00151036118b557868461187761209d565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016118ac93929190614af6565b60405180910390fd5b600260ff168360c001510361190e577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856118ef61209d565b6040516118fe93929190614aba565b60405180910390a1505050611959565b5b7f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a88856040516119409291906148e6565b60405180910390a15050505b8080600101915050611643565b505050505050565b5f5f1b810361199c576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119c87fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b826120bb565b7f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa816040516119f79190613d87565b60405180910390a1611a287f0000000000000000000000000000000000000000000000000000000000000000610e8f565b50565b611a8f7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff165f1b6120c2565b7f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed18282604051611ac0929190614b39565b60405180910390a15050565b5f5f611adc8360200151306120f7565b90505f611ae88461219b565b90508181604051602001611afd929190614bd4565b6040516020818303038152906040528051906020012092505050919050565b5f5f5f611b4b7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b85610e57565b5f1c9050606081901c816bffffffffffffffffffffffff169250925050915091565b5f3054905090565b5f5f5f5f5f5f5f611b868b8b611f27565b915060ff169150611b9561355e565b6040808416148015611bd257505f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff16145b15611d0e57611bec828d8d611fd59290919263ffffffff16565b809350819a50505089611d0d575f611c0f838e8e61201c9290919263ffffffff16565b8162ffffff16915080945081925050505f8d8d85908487611c309190614431565b92611c3d93929190614971565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090508a73ffffffffffffffffffffffffffffffffffffffff1663ccce3bc830836040518363ffffffff1660e01b8152600401611cbc929190614c0a565b6040805180830381865afa158015611cd6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611cfa9190614cad565b92508184611d089190614431565b935050505b5b600180841603611d4757611d358d8a838f8f87908092611d3093929190614971565b6123cf565b97509750975097509750505050611ea5565b6002808416148d60200190151590811515815250505f6002601c8516901c9050611d8383828f8f611f6e9190939291909392919063ffffffff16565b8094508197505050505f6001600560208616901c611da19190614431565b9050611dbf83828f8f611f6e9190939291909392919063ffffffff16565b809450819a50505050611dd18d611acc565b9350611def8d858e8e86908092611dea93929190614971565b612634565b8097508198505050611e0386895f1b6131b1565b9550611e1186865f1b6131b1565b9550611e35868a73ffffffffffffffffffffffffffffffffffffffff165f1b6131b1565b95505f5f1b815f015114158015611e4f575085815f015114155b8015611e5f575080602001518511155b15611ea157806040517fccbb534f000000000000000000000000000000000000000000000000000000008152600401611e989190614d05565b60405180910390fd5b5050505b9550955095509550959050565b5f611ebc826131c5565b9050919050565b611f1c7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b846bffffffffffffffffffffffff841660608673ffffffffffffffffffffffffffffffffffffffff16901b175f1b6120c2565b505050565b80305550565b5f5f83358060f81c925060019150509250929050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f858401356008840261010003600180866008021b0382821c1693508486019250505094509492505050565b5f5f8483013561ffff8160f01c16925060028401915050935093915050565b5f5f848301358060f81c925060018401915050935093915050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f848301359150602083019050935093915050565b5f5f8483013562ffffff8160e81c16925060038401915050935093915050565b61206d7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b835f1b6120c2565b5050565b5f5f5f8351602085018787f490509392505050565b5f5f5f835160208501878988f19050949350505050565b60603d604051915060208201818101604052818352815f823e505090565b8082555050565b5f83836040516020016120d69291906148bf565b60405160208183030381529060405280519060200120905081815550505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121665746612168565b5f5b8560405160200161217d959493929190614d1e565b60405160208183030381529060405280519060200120905092915050565b5f5f8261010001516040516020016121b39190614dfb565b6040516020818303038152906040528051906020012090505f60ff16835f015160ff160361224b575f6121e98460400151613268565b90507f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a281856060015186608001518560405160200161222c959493929190614e11565b60405160208183030381529060405280519060200120925050506123ca565b600160ff16835f015160ff16036122ba577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360a00151805190602001208260405160200161229c93929190614e62565b604051602081830303815290604052805190602001209150506123ca565b600260ff16835f015160ff1603612322577f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e48360c001518260405160200161230493929190614e62565b604051602081830303815290604052805190602001209150506123ca565b600360ff16835f015160ff160361238a577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360e001518260405160200161236c93929190614e62565b604051602081830303815290604052805190602001209150506123ca565b825f01516040517f048183200000000000000000000000000000000000000000000000000000000081526004016123c19190614ea6565b60405180910390fd5b919050565b5f5f5f5f5f6123dc6134c2565b6002815f019060ff16908160ff16815250505f5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b898990508210156125c9575f5f612436848d8d61201c9290919263ffffffff16565b8162ffffff169150809550819250505083816124529190614431565b9150505f8b8b90508214612466575f612468565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83036124c8576124af8f8d8d879086926124a793929190614971565b600185611b75565b809a50819b50829c50839d50849e5050505050506124f8565b6124e6858d8d879086926124de93929190614971565b600185611b75565b50809a50819b50829c50839d50505050505b89891015612553578b8b8590849261251293929190614971565b8b8b6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161254a9493929190614ebf565b60405180910390fd5b819350878d5f01510361256c575f5f1b8d5f0181815250505b8287106125b25786836040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004016125a9929190614942565b60405180910390fd5b878560c00181815250508692508193505050612414565b5f5f1b8b5f0151141580156125e257508a602001518511155b15612624578a6040517fccbb534f00000000000000000000000000000000000000000000000000000000815260040161261b9190614d05565b60405180910390fd5b5050509550955095509550959050565b5f5f5f5b848490508110156131a7575f612659828787611fba9290919263ffffffff16565b8160ff16915080935081925050505f600460f08316901c90505f81036127b1575f600f831690505f8160ff16036126a85761269f848989611fba9290919263ffffffff16565b80955081925050505b5f5f6126bf868b8b6132e29290919263ffffffff16565b80975081935050506126dc868b8b6132e29290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f8388866040515f81526020016040526040516127429493929190614efd565b6020604051602081039080840390855afa158015612762573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612784828960ff166132f8565b90505f5f1b8c0361279557806127a0565b61279f8c826131b1565b5b9b5050505050505050505050612638565b6001810361283c575f600f831690505f8160ff16036127e8576127df848989611fba9290919263ffffffff16565b80955081925050505b5f6127fe858a8a611fd59290919263ffffffff16565b80965081925050505f612814828460ff166132f8565b90505f5f1b87036128255780612830565b61282f87826131b1565b5b96505050505050612638565b60028103612a3b575f6003831690505f8160ff16036128735761286a848989611fba9290919263ffffffff16565b80955081925050505b5f612889858a8a611fd59290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f6128bf87838d8d611f6e9190939291909392919063ffffffff16565b80985081925050505f81880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261292393929190614971565b6040518463ffffffff1660e01b815260040161294193929190614f40565b602060405180830381865afa15801561295c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129809190614f84565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146129f7578c848d8d8b9085926129b893929190614971565b6040517fb2fed7ae0000000000000000000000000000000000000000000000000000000081526004016129ee9493929190614faf565b60405180910390fd5b8097508460ff168a0199505f612a10858760ff166132f8565b90505f5f1b8a03612a215780612a2c565b612a2b8a826131b1565b5b99505050505050505050612638565b60038103612a85575f612a598489896132e29290919263ffffffff16565b80955081925050505f5f1b8503612a705780612a7b565b612a7a85826131b1565b5b9450505050612638565b60048103612b07575f600f831660ff1690505f612ab485838b8b611f6e9190939291909392919063ffffffff16565b80965081925050505f81860190505f5f612ae08e8e8e8e8c908892612adb93929190614971565b612634565b91509150829750818a019950612af689826131b1565b985082975050505050505050612638565b60068103612c17575f6002600c841660ff16901c60ff1690505f8103612b4b57612b3c848989611fba9290919263ffffffff16565b8160ff16915080955081925050505b5f6003841660ff1690505f8103612b8157612b71858a8a611f9b9290919263ffffffff16565b8161ffff16915080965081925050505b5f612b97868b8b61201c9290919263ffffffff16565b8162ffffff16915080975081925050505f81870190505f5f612bcb8f8f8f8f8d908892612bc693929190614971565b612634565b91509150829850848210612bdf57858b019a505b5f612beb82878961332a565b90505f5f1b8b03612bfc5780612c07565b612c068b826131b1565b5b9a50505050505050505050612638565b60058103612c99575f612c358489896132e29290919263ffffffff16565b8095508192505050888103612c68577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b5f612c728261335f565b90505f5f1b8603612c835780612c8e565b612c8d86826131b1565b5b955050505050612638565b60078103612dff575f600f831690505f8160ff1603612cd057612cc7848989611fba9290919263ffffffff16565b80955081925050505b5f5f612ce7868b8b6132e29290919263ffffffff16565b8097508193505050612d04868b8b6132e29290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f604051602001612d5b9190615037565b604051602081830303815290604052805190602001208388866040515f8152602001604052604051612d909493929190614efd565b6020604051602081039080840390855afa158015612db0573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612dd2828960ff166132f8565b90505f5f1b8c03612de35780612dee565b612ded8c826131b1565b5b9b5050505050505050505050612638565b60088103612e98575f612e1d8489896132e29290919263ffffffff16565b80955081925050505f612e395f8c61338e90919063ffffffff16565b9050808203612e66577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b5f612e70826133df565b90505f5f1b8703612e815780612e8c565b612e8b87826131b1565b5b96505050505050612638565b60098103613001575f6003831690505f8160ff1603612ecf57612ec6848989611fba9290919263ffffffff16565b80955081925050505b5f612ee5858a8a611fd59290919263ffffffff16565b80965081925050505f5f6002600c871660ff16901c60ff169050612f1b87828d8d611f6e9190939291909392919063ffffffff16565b8098508193505050505f81870190505f8373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612f5a93929190614971565b6040518463ffffffff1660e01b8152600401612f789392919061505c565b602060405180830381865afa158015612f93573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fb79190615093565b90508197508460ff168a0199505f612fd3858760ff168461340e565b90505f5f1b8a03612fe45780612fef565b612fee8a826131b1565b5b99508298505050505050505050612638565b600a810361316a575f6003831690505f8160ff16036130385761302f848989611fba9290919263ffffffff16565b80955081925050505b5f61304e858a8a611fd59290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f61308487838d8d611f6e9190939291909392919063ffffffff16565b80985081925050505f81880190505f8473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d9087926130c293929190614971565b6040518463ffffffff1660e01b81526004016130e093929190614f40565b602060405180830381865afa1580156130fb573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061311f9190615093565b90508198508560ff168b019a505f61313b868860ff168461340e565b90505f5f1b8b0361314c5780613157565b6131568b826131b1565b5b9a50829950505050505050505050612638565b806040517fb2505f7c00000000000000000000000000000000000000000000000000000000815260040161319e91906140c5565b60405180910390fd5b5094509492505050565b5f825f528160205260405f20905092915050565b5f3073ffffffffffffffffffffffffffffffffffffffff167f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006040516020016132329392919061514d565b604051602081830303815290604052805190602001205f1c73ffffffffffffffffffffffffffffffffffffffff16149050919050565b5f60605f5f90505b83518110156132d1575f61329d8583815181106132905761328f614464565b5b6020026020010151613443565b905082816040516020016132b29291906151c4565b6040516020818303038152906040529250508080600101915050613270565b508080519060200120915050919050565b5f5f848301359150602083019050935093915050565b5f828260405160200161330c929190615255565b60405160208183030381529060405280519060200120905092915050565b5f838383604051602001613340939291906152d5565b6040516020818303038152906040528051906020012090509392505050565b5f816040516020016133719190615366565b604051602081830303815290604052805190602001209050919050565b5f5f61339e8460200151846120f7565b90505f6133aa8561219b565b905081816040516020016133bf929190614bd4565b604051602081830303815290604052805190602001209250505092915050565b5f816040516020016133f191906153d5565b604051602081830303815290604052805190602001209050919050565b5f83838360405160200161342493929190615444565b6040516020818303038152906040528051906020012090509392505050565b5f7f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437825f01518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016134a598979695949392919061548b565b604051602081830303815290604052805190602001209050919050565b6040518061012001604052805f60ff1681526020015f15158152602001606081526020015f81526020015f8152602001606081526020015f81526020015f8152602001606081525090565b6040518060e001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020015f81526020015f151581526020015f151581526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f82905092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f82821b905092915050565b5f6135c28383613576565b826135cd8135613580565b9250600482101561360d576136087fffffffff00000000000000000000000000000000000000000000000000000000836004036008026135ab565b831692505b505092915050565b5f81905092915050565b828183375f83830152505050565b5f6136388385613615565b935061364583858461361f565b82840190509392505050565b5f61365d82848661362d565b91508190509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6136a38261367a565b9050919050565b6136b381613699565b81146136bd575f5ffd5b50565b5f813590506136ce816136aa565b92915050565b5f602082840312156136e9576136e8613672565b5b5f6136f6848285016136c0565b91505092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61374982613703565b810181811067ffffffffffffffff8211171561376857613767613713565b5b80604052505050565b5f61377a613669565b90506137868282613740565b919050565b5f5ffd5b5f60ff82169050919050565b6137a48161378f565b81146137ae575f5ffd5b50565b5f813590506137bf8161379b565b92915050565b5f8115159050919050565b6137d9816137c5565b81146137e3575f5ffd5b50565b5f813590506137f4816137d0565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561381857613817613713565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b61383f8161382d565b8114613849575f5ffd5b50565b5f8135905061385a81613836565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561387e5761387d613713565b5b61388782613703565b9050602081019050919050565b5f6138a66138a184613864565b613771565b9050828152602081018484840111156138c2576138c1613860565b5b6138cd84828561361f565b509392505050565b5f82601f8301126138e9576138e86137fa565b5b81356138f9848260208601613894565b91505092915050565b5f60e08284031215613917576139166136ff565b5b61392160e0613771565b90505f613930848285016136c0565b5f8301525060206139438482850161384c565b602083015250604082013567ffffffffffffffff8111156139675761396661378b565b5b613973848285016138d5565b60408301525060606139878482850161384c565b606083015250608061399b848285016137e6565b60808301525060a06139af848285016137e6565b60a08301525060c06139c38482850161384c565b60c08301525092915050565b5f6139e16139dc846137fe565b613771565b90508083825260208201905060208402830185811115613a0457613a03613829565b5b835b81811015613a4b57803567ffffffffffffffff811115613a2957613a286137fa565b5b808601613a368982613902565b85526020850194505050602081019050613a06565b5050509392505050565b5f82601f830112613a6957613a686137fa565b5b8135613a798482602086016139cf565b91505092915050565b5f819050919050565b613a9481613a82565b8114613a9e575f5ffd5b50565b5f81359050613aaf81613a8b565b92915050565b5f67ffffffffffffffff821115613acf57613ace613713565b5b602082029050602081019050919050565b5f613af2613aed84613ab5565b613771565b90508083825260208201905060208402830185811115613b1557613b14613829565b5b835b81811015613b3e5780613b2a88826136c0565b845260208401935050602081019050613b17565b5050509392505050565b5f82601f830112613b5c57613b5b6137fa565b5b8135613b6c848260208601613ae0565b91505092915050565b5f6101208284031215613b8b57613b8a6136ff565b5b613b96610120613771565b90505f613ba5848285016137b1565b5f830152506020613bb8848285016137e6565b602083015250604082013567ffffffffffffffff811115613bdc57613bdb61378b565b5b613be884828501613a55565b6040830152506060613bfc8482850161384c565b6060830152506080613c108482850161384c565b60808301525060a082013567ffffffffffffffff811115613c3457613c3361378b565b5b613c40848285016138d5565b60a08301525060c0613c5484828501613aa1565b60c08301525060e0613c6884828501613aa1565b60e08301525061010082013567ffffffffffffffff811115613c8d57613c8c61378b565b5b613c9984828501613b48565b6101008301525092915050565b5f5ffd5b5f5f83601f840112613cbf57613cbe6137fa565b5b8235905067ffffffffffffffff811115613cdc57613cdb613ca6565b5b602083019150836001820283011115613cf857613cf7613829565b5b9250929050565b5f5f5f60408486031215613d1657613d15613672565b5b5f84013567ffffffffffffffff811115613d3357613d32613676565b5b613d3f86828701613b75565b935050602084013567ffffffffffffffff811115613d6057613d5f613676565b5b613d6c86828701613caa565b92509250509250925092565b613d8181613a82565b82525050565b5f602082019050613d9a5f830184613d78565b92915050565b5f5f5f5f5f60808688031215613db957613db8613672565b5b5f613dc6888289016136c0565b9550506020613dd7888289016136c0565b9450506040613de88882890161384c565b935050606086013567ffffffffffffffff811115613e0957613e08613676565b5b613e1588828901613caa565b92509250509295509295909350565b613e2d81613580565b82525050565b5f602082019050613e465f830184613e24565b92915050565b5f5f5f60408486031215613e6357613e62613672565b5b5f613e7086828701613aa1565b935050602084013567ffffffffffffffff811115613e9157613e90613676565b5b613e9d86828701613caa565b92509250509250925092565b613eb281613580565b8114613ebc575f5ffd5b50565b5f81359050613ecd81613ea9565b92915050565b5f60208284031215613ee857613ee7613672565b5b5f613ef584828501613ebf565b91505092915050565b613f0781613699565b82525050565b5f602082019050613f205f830184613efe565b92915050565b5f5f5f5f60408587031215613f3e57613f3d613672565b5b5f85013567ffffffffffffffff811115613f5b57613f5a613676565b5b613f6787828801613caa565b9450945050602085013567ffffffffffffffff811115613f8a57613f89613676565b5b613f9687828801613caa565b925092505092959194509250565b5f60208284031215613fb957613fb8613672565b5b5f613fc684828501613aa1565b91505092915050565b5f5f60208385031215613fe557613fe4613672565b5b5f83013567ffffffffffffffff81111561400257614001613676565b5b61400e85828601613caa565b92509250509250929050565b5f5f5f5f6060858703121561403257614031613672565b5b5f61403f878288016136c0565b94505060206140508782880161384c565b935050604085013567ffffffffffffffff81111561407157614070613676565b5b61407d87828801613caa565b925092505092959194509250565b5f602082840312156140a05761409f613672565b5b5f6140ad8482850161384c565b91505092915050565b6140bf8161382d565b82525050565b5f6020820190506140d85f8301846140b6565b92915050565b5f6040820190506140f15f830185613efe565b6140fe60208301846140b6565b9392505050565b61410e816137c5565b82525050565b5f60c0820190506141275f8301896140b6565b61413460208301886140b6565b6141416040830187614105565b61414e6060830186613d78565b61415b60808301856140b6565b61416860a0830184613d78565b979650505050505050565b5f5f6040838503121561418957614188613672565b5b5f61419685828601613ebf565b92505060206141a7858286016136c0565b9150509250929050565b5f5f83601f8401126141c6576141c56137fa565b5b8235905067ffffffffffffffff8111156141e3576141e2613ca6565b5b6020830191508360208202830111156141ff576141fe613829565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b03121561422257614221613672565b5b5f61422f8b828c016136c0565b98505060206142408b828c016136c0565b975050604089013567ffffffffffffffff81111561426157614260613676565b5b61426d8b828c016141b1565b9650965050606089013567ffffffffffffffff8111156142905761428f613676565b5b61429c8b828c016141b1565b9450945050608089013567ffffffffffffffff8111156142bf576142be613676565b5b6142cb8b828c01613caa565b92509250509295985092959890939650565b5f5f5f5f5f5f60a087890312156142f7576142f6613672565b5b5f61430489828a016136c0565b965050602061431589828a016136c0565b955050604061432689828a0161384c565b945050606061433789828a0161384c565b935050608087013567ffffffffffffffff81111561435857614357613676565b5b61436489828a01613caa565b92509250509295509295509295565b5f6bffffffffffffffffffffffff82169050919050565b61439381614373565b811461439d575f5ffd5b50565b5f813590506143ae8161438a565b92915050565b5f5f5f606084860312156143cb576143ca613672565b5b5f6143d886828701613aa1565b93505060206143e9868287016136c0565b92505060406143fa868287016143a0565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61443b8261382d565b91506144468361382d565b925082820190508082111561445e5761445d614404565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61449a8161378f565b82525050565b6144a9816137c5565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6144e181613699565b82525050565b6144f08161382d565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614528826144f6565b6145328185614500565b9350614542818560208601614510565b61454b81613703565b840191505092915050565b5f60e083015f83015161456b5f8601826144d8565b50602083015161457e60208601826144e7565b5060408301518482036040860152614596828261451e565b91505060608301516145ab60608601826144e7565b5060808301516145be60808601826144a0565b5060a08301516145d160a08601826144a0565b5060c08301516145e460c08601826144e7565b508091505092915050565b5f6145fa8383614556565b905092915050565b5f602082019050919050565b5f614618826144af565b61462281856144b9565b935083602082028501614634856144c9565b805f5b8581101561466f578484038952815161465085826145ef565b945061465b83614602565b925060208a01995050600181019050614637565b50829750879550505050505092915050565b61468a81613a82565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6146c483836144d8565b60208301905092915050565b5f602082019050919050565b5f6146e682614690565b6146f0818561469a565b93506146fb836146aa565b805f5b8381101561472b57815161471288826146b9565b975061471d836146d0565b9250506001810190506146fe565b5085935050505092915050565b5f61012083015f83015161474e5f860182614491565b50602083015161476160208601826144a0565b5060408301518482036040860152614779828261460e565b915050606083015161478e60608601826144e7565b5060808301516147a160808601826144e7565b5060a083015184820360a08601526147b9828261451e565b91505060c08301516147ce60c0860182614681565b5060e08301516147e160e0860182614681565b506101008301518482036101008601526147fb82826146dc565b9150508091505092915050565b5f82825260208201905092915050565b5f6148238385614808565b935061483083858461361f565b61483983613703565b840190509392505050565b5f6040820190508181035f83015261485c8186614738565b90508181036020830152614871818486614818565b9050949350505050565b61488481614373565b82525050565b5f60608201905061489d5f830186613d78565b6148aa6020830185613efe565b6148b7604083018461487b565b949350505050565b5f6040820190506148d25f830185613d78565b6148df6020830184613d78565b9392505050565b5f6040820190506148f95f830185613d78565b61490660208301846140b6565b9392505050565b5f6060820190506149205f830186613d78565b61492d6020830185613efe565b61493a6040830184613efe565b949350505050565b5f6040820190506149555f8301856140b6565b61496260208301846140b6565b9392505050565b5f5ffd5b5f5ffd5b5f5f8585111561498457614983614969565b5b838611156149955761499461496d565b5b6001850283019150848603905094509492505050565b5f6060820190506149be5f8301866140b6565b6149cb60208301856140b6565b6149d860408301846140b6565b949350505050565b5f6060820190508181035f8301526149f88186614738565b9050614a0760208301856140b6565b614a1460408301846140b6565b949350505050565b5f614a26826144f6565b614a308185614808565b9350614a40818560208601614510565b614a4981613703565b840191505092915050565b5f60c082019050614a675f830189613d78565b614a7460208301886140b6565b614a8160408301876140b6565b614a8e60608301866140b6565b614a9b60808301856140b6565b81810360a0830152614aad8184614a1c565b9050979650505050505050565b5f606082019050614acd5f830186613d78565b614ada60208301856140b6565b8181036040830152614aec8184614a1c565b9050949350505050565b5f6060820190508181035f830152614b0e8186614738565b9050614b1d60208301856140b6565b8181036040830152614b2f8184614a1c565b9050949350505050565b5f604082019050614b4c5f830185613e24565b614b596020830184613efe565b9392505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f614b9e600283614b60565b9150614ba982614b6a565b600282019050919050565b5f819050919050565b614bce614bc982613a82565b614bb4565b82525050565b5f614bde82614b92565b9150614bea8285614bbd565b602082019150614bfa8284614bbd565b6020820191508190509392505050565b5f604082019050614c1d5f830185613efe565b8181036020830152614c2f8184614a1c565b90509392505050565b5f81519050614c4681613a8b565b92915050565b5f81519050614c5a81613836565b92915050565b5f60408284031215614c7557614c746136ff565b5b614c7f6040613771565b90505f614c8e84828501614c38565b5f830152506020614ca184828501614c4c565b60208301525092915050565b5f60408284031215614cc257614cc1613672565b5b5f614ccf84828501614c60565b91505092915050565b604082015f820151614cec5f850182614681565b506020820151614cff60208501826144e7565b50505050565b5f604082019050614d185f830184614cd8565b92915050565b5f60a082019050614d315f830188613d78565b614d3e6020830187613d78565b614d4b6040830186613d78565b614d5860608301856140b6565b614d656080830184613efe565b9695505050505050565b5f81905092915050565b614d8281613699565b82525050565b5f614d938383614d79565b60208301905092915050565b5f614da982614690565b614db38185614d6f565b9350614dbe836146aa565b805f5b83811015614dee578151614dd58882614d88565b9750614de0836146d0565b925050600181019050614dc1565b5085935050505092915050565b5f614e068284614d9f565b915081905092915050565b5f60a082019050614e245f830188613d78565b614e316020830187613d78565b614e3e60408301866140b6565b614e4b60608301856140b6565b614e586080830184613d78565b9695505050505050565b5f606082019050614e755f830186613d78565b614e826020830185613d78565b614e8f6040830184613d78565b949350505050565b614ea08161378f565b82525050565b5f602082019050614eb95f830184614e97565b92915050565b5f6060820190508181035f830152614ed8818688614818565b9050614ee760208301856140b6565b614ef460408301846140b6565b95945050505050565b5f608082019050614f105f830187613d78565b614f1d6020830186614e97565b614f2a6040830185613d78565b614f376060830184613d78565b95945050505050565b5f604082019050614f535f830186613d78565b8181036020830152614f66818486614818565b9050949350505050565b5f81519050614f7e81613ea9565b92915050565b5f60208284031215614f9957614f98613672565b5b5f614fa684828501614f70565b91505092915050565b5f606082019050614fc25f830187613d78565b614fcf6020830186613efe565b8181036040830152614fe2818486614818565b905095945050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f615021601c83614b60565b915061502c82614fed565b601c82019050919050565b5f61504182615015565b915061504d8284614bbd565b60208201915081905092915050565b5f6040820190508181035f8301526150748186614738565b90508181036020830152615089818486614818565b9050949350505050565b5f602082840312156150a8576150a7613672565b5b5f6150b584828501614c38565b91505092915050565b7fff000000000000000000000000000000000000000000000000000000000000005f82015250565b5f6150f2600183614b60565b91506150fd826150be565b600182019050919050565b5f8160601b9050919050565b5f61511e82615108565b9050919050565b5f61512f82615114565b9050919050565b61514761514282613699565b615125565b82525050565b5f615157826150e6565b91506151638286615136565b6014820191506151738285614bbd565b6020820191506151838284614bbd565b602082019150819050949350505050565b5f61519e826144f6565b6151a88185613615565b93506151b8818560208601614510565b80840191505092915050565b5f6151cf8285615194565b91506151db8284614bbd565b6020820191508190509392505050565b7f53657175656e6365207369676e65723a0a0000000000000000000000000000005f82015250565b5f61521f601183614b60565b915061522a826151eb565b601182019050919050565b5f819050919050565b61524f61524a8261382d565b615235565b82525050565b5f61525f82615213565b915061526b8285615136565b60148201915061527b828461523e565b6020820191508190509392505050565b7f53657175656e6365206e657374656420636f6e6669673a0a00000000000000005f82015250565b5f6152bf601883614b60565b91506152ca8261528b565b601882019050919050565b5f6152df826152b3565b91506152eb8286614bbd565b6020820191506152fb828561523e565b60208201915061530b828461523e565b602082019150819050949350505050565b7f53657175656e636520737461746963206469676573743a0a00000000000000005f82015250565b5f615350601883614b60565b915061535b8261531c565b601882019050919050565b5f61537082615344565b915061537c8284614bbd565b60208201915081905092915050565b7f53657175656e636520616e792061646472657373207375626469676573743a0a5f82015250565b5f6153bf602083614b60565b91506153ca8261538b565b602082019050919050565b5f6153df826153b3565b91506153eb8284614bbd565b60208201915081905092915050565b7f53657175656e63652073617069656e7420636f6e6669673a0a000000000000005f82015250565b5f61542e601983614b60565b9150615439826153fa565b601982019050919050565b5f61544e82615422565b915061545a8286615136565b60148201915061546a828561523e565b60208201915061547a8284614bbd565b602082019150819050949350505050565b5f6101008201905061549f5f83018b613d78565b6154ac602083018a613efe565b6154b960408301896140b6565b6154c66060830188613d78565b6154d360808301876140b6565b6154e060a0830186614105565b6154ed60c0830185614105565b6154fa60e08301846140b6565b999850505050505050505056fea2646970667358221220c9d74b60173f281cd82bd7e3fc3c73996ce296df881f7a829c762bff92fe78f964736f6c634300081c00336080604052348015600e575f5ffd5b506153888061001c5f395ff3fe608060405260043610610117575f3560e01c80636ea445771161009f578063ad55366b11610063578063ad55366b14610475578063b93ea7ad146104b6578063bc197c81146104d2578063f23a6e611461050e578063f727ef1c1461054a5761011e565b80636ea445771461038e5780638943ec02146103aa5780638c3f5563146103d257806392dcb3fc1461040e578063aaf10f421461044b5761011e565b80631a9b2337116100e65780631a9b2337146102c85780631f6a1eb91461030457806329561426146103205780634fcf3eca1461034857806351605d80146103645761011e565b8063025b22bc146101f857806313792a4a14610214578063150b7a02146102505780631626ba7e1461028c5761011e565b3661011e57005b60045f369050106101f6575f61013f5f369061013a9190613493565b610572565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146101f4575f5f8273ffffffffffffffffffffffffffffffffffffffff165f3660405161019d92919061352d565b5f60405180830381855af49150503d805f81146101d5576040519150601f19603f3d011682016040523d82523d5f602084013e6101da565b606091505b5091509150816101ec57805160208201fd5b805160208201f35b505b005b610212600480360381019061020d91906135b0565b6105c7565b005b34801561021f575f5ffd5b5061023a60048036038101906102359190613bdb565b610643565b6040516102479190613c63565b60405180910390f35b34801561025b575f5ffd5b5061027660048036038101906102719190613c7c565b6107f2565b6040516102839190613d0f565b60405180910390f35b348015610297575f5ffd5b506102b260048036038101906102ad9190613d28565b610806565b6040516102bf9190613d0f565b60405180910390f35b3480156102d3575f5ffd5b506102ee60048036038101906102e99190613daf565b610848565b6040516102fb9190613de9565b60405180910390f35b61031e60048036038101906103199190613e02565b610859565b005b34801561032b575f5ffd5b5061034660048036038101906103419190613e80565b6108e8565b005b610362600480360381019061035d9190613daf565b610964565b005b34801561036f575f5ffd5b50610378610a59565b6040516103859190613c63565b60405180910390f35b6103a860048036038101906103a39190613eab565b610a8a565b005b3480156103b5575f5ffd5b506103d060048036038101906103cb9190613ef6565b610b29565b005b3480156103dd575f5ffd5b506103f860048036038101906103f39190613f67565b610b2f565b6040516104059190613fa1565b60405180910390f35b348015610419575f5ffd5b50610434600480360381019061042f9190613e80565b610b67565b604051610442929190613fba565b60405180910390f35b348015610456575f5ffd5b5061045f610b7b565b60405161046c9190613de9565b60405180910390f35b348015610480575f5ffd5b5061049b60048036038101906104969190613bdb565b610b89565b6040516104ad96959493929190613ff0565b60405180910390f35b6104d060048036038101906104cb919061404f565b610bc7565b005b3480156104dd575f5ffd5b506104f860048036038101906104f391906140e2565b610cbd565b6040516105059190613d0f565b60405180910390f35b348015610519575f5ffd5b50610534600480360381019061052f91906141b9565b610cd4565b6040516105419190613d0f565b60405180910390f35b348015610555575f5ffd5b50610570600480360381019061056b9190614290565b610ce9565b005b5f6105be7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610db2565b5f1c9050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063757336040517fa19dbf0000000000000000000000000000000000000000000000000000000000815260040161062e9190613de9565b60405180910390fd5b61064081610dea565b50565b5f5f600185610100015151610658919061430d565b67ffffffffffffffff811115610671576106706135ef565b5b60405190808252806020026020018201604052801561069f5781602001602082028036833780820191505090505b5090505f5f90505b8561010001515181101561072f5785610100015181815181106106cd576106cc614340565b5b60200260200101518282815181106106e8576106e7614340565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506106a7565b503381866101000151518151811061074a57610749614340565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808561010001819052505f61079a868686610e2d565b509050806107e3578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016107da93929190614720565b60405180910390fd5b60015f1b925050509392505050565b5f63150b7a0260e01b905095945050505050565b5f5f6108118561101d565b90505f61081f828686610e2d565b50905080610834575f60e01b92505050610841565b6320c13b0b60e01b925050505b9392505050565b5f61085282610572565b9050919050565b5f5a90505f6108688686611046565b905061087c816060015182608001516114e8565b5f5f610889838787610e2d565b91509150816108d3578286866040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016108ca93929190614720565b60405180910390fd5b6108de84828561158c565b5050505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095857336040517fa19dbf0000000000000000000000000000000000000000000000000000000000815260040161094f9190613de9565b60405180910390fd5b610961816118bc565b50565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d457336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016109cb9190613de9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166109f482610572565b73ffffffffffffffffffffffffffffffffffffffff1603610a4c57806040517f1c3812cc000000000000000000000000000000000000000000000000000000008152600401610a439190613d0f565b60405180910390fd5b610a56815f61195d565b50565b5f610a857fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b6119fe565b905090565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610afa57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610af19190613de9565b60405180910390fd5b5f5a90505f610b098484611046565b90505f610b1582611a08565b9050610b2283828461158c565b5050505050565b50505050565b5f610b5e7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b610db2565b5f1c9050919050565b5f5f610b7283611a58565b91509150915091565b5f610b84611aa9565b905090565b5f5f5f5f5f5f610b9c8989895f5f611ab1565b809550819650829750839950849a505050505050610bb983611dee565b935093975093979195509350565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c3757336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610c2e9190613de9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610c5783610572565b73ffffffffffffffffffffffffffffffffffffffff1614610caf57816040517f5b4d6d6a000000000000000000000000000000000000000000000000000000008152600401610ca69190613d0f565b60405180910390fd5b610cb9828261195d565b5050565b5f63bc197c8160e01b905098975050505050505050565b5f63f23a6e6160e01b90509695505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d5957336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610d509190613de9565b60405180910390fd5b610d728383836bffffffffffffffffffffffff16611dff565b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1838383604051610da593929190614766565b60405180910390a1505050565b5f5f8383604051602001610dc792919061479b565b604051602081830303815290604052805190602001209050805491505092915050565b610df381611e5d565b7f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0381604051610e229190613de9565b60405180910390a150565b5f5f5f84845f818110610e4357610e42614340565b5b9050013560f81c60f81b9050608060f81b608060f81b82167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610f9c57610e8b86611a08565b91505f5f610e9884611a58565b91509150428111610ee25783816040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600401610ed99291906147c2565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610f4a57503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15610f90578333836040517f8945c313000000000000000000000000000000000000000000000000000000008152600401610f87939291906147e9565b60405180910390fd5b60019450505050611015565b5f5f5f610fac8989895f5f611ab1565b905080985081945082955083965050505050828210156110055782826040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600401610ffc92919061481e565b60405180910390fd5b61100e81611dee565b9550505050505b935093915050565b61102561339e565b6003815f019060ff16908160ff1681525050818160e0018181525050919050565b61104e61339e565b5f815f019060ff16908160ff16815250505f5f61106b8585611e63565b915060ff16915060018083160361108b575f8360600181815250506110c7565b6110a0818686611e799290919263ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff169150846060018193508281525050505b5f6007600184901c1690505f811115611104576110f682828888611eaa9190939291909392919063ffffffff16565b856080018194508281525050505b5f6010808516036111185760019050611170565b60208085160361114b57611137838888611ed79290919263ffffffff16565b8161ffff169150809450819250505061116f565b611160838888611ef69290919263ffffffff16565b8160ff16915080945081925050505b5b8067ffffffffffffffff81111561118a576111896135ef565b5b6040519080825280602002602001820160405280156111c357816020015b6111b06133e9565b8152602001906001900390816111a85790505b5085604001819052505f5f90505b818110156114dd575f6111ef858a8a611ef69290919263ffffffff16565b8096508192505050600180821660ff160361125e57308760400151838151811061121c5761121b614340565b5b60200260200101515f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506112ca565b611273858a8a611f119290919263ffffffff16565b8860400151848151811061128a57611289614340565b5b60200260200101515f018197508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b600280821660ff1603611318576112ec858a8a611f429290919263ffffffff16565b8860400151848151811061130357611302614340565b5b60200260200101516020018197508281525050505b600480821660ff16036113e0575f61133b868b8b611f589290919263ffffffff16565b8162ffffff169150809750819250505089898790838961135b919061430d565b926113689392919061484d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050886040015184815181106113c1576113c0614340565b5b60200260200101516040018190525080866113dc919061430d565b9550505b600880821660ff160361142e57611402858a8a611f429290919263ffffffff16565b8860400151848151811061141957611418614340565b5b60200260200101516060018197508281525050505b601080821660ff16148760400151838151811061144e5761144d614340565b5b60200260200101516080019015159081151581525050602080821660ff16148760400151838151811061148457611483614340565b5b602002602001015160a0019015159081151581525050600660c0821660ff16901c60ff16876040015183815181106114bf576114be614340565b5b602002602001015160c00181815250505080806001019150506111d1565b505050505092915050565b5f6114f283610b2f565b905081811461153c578282826040517f9b6514f400000000000000000000000000000000000000000000000000000000815260040161153393929190614887565b60405180910390fd5b5f60018301905061154d8482611f78565b7f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881848260405161157e92919061481e565b60405180910390a150505050565b5f5f90505f82604001515190505f5f90505b818110156118b4575f846040015182815181106115be576115bd614340565b5b602002602001015190508060a0015180156115d7575083155b1561161b577f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b868360405161160d9291906147c2565b60405180910390a1506118a7565b5f93505f816060015190505f81141580156116355750805a105b1561167b5785835a6040517f21395274000000000000000000000000000000000000000000000000000000008152600401611672939291906148bc565b60405180910390fd5b5f82608001511561173057611729835f01515f841461169a578361169c565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016116c796959493929190614930565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611fad565b9050611758565b611755835f015184602001515f8514611749578461174b565b5a5b8660400151611fc2565b90505b8061186a575f60ff168360c00151036117b957600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d888561179a611fd9565b6040516117a993929190614996565b60405180910390a15050506118a7565b600160ff168360c00151036118105786846117d2611fd9565b6040517f7f6b0bb1000000000000000000000000000000000000000000000000000000008152600401611807939291906149d2565b60405180910390fd5b600260ff168360c0015103611869577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b888561184a611fd9565b60405161185993929190614996565b60405180910390a15050506118b4565b5b7f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a888560405161189b9291906147c2565b60405180910390a15050505b808060010191505061159e565b505050505050565b5f5f1b81036118f7576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119237fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b82611ff7565b7f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa816040516119529190613c63565b60405180910390a150565b6119c17fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff165f1b611ffe565b7f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed182826040516119f2929190614a15565b60405180910390a15050565b5f81549050919050565b5f5f611a18836020015130612033565b90505f611a24846120d7565b90508181604051602001611a39929190614ab0565b6040516020818303038152906040528051906020012092505050919050565b5f5f5f611a877fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b85610db2565b5f1c9050606081901c816bffffffffffffffffffffffff169250925050915091565b5f3054905090565b5f5f5f5f5f5f5f611ac28b8b611e63565b915060ff169150611ad161343a565b6040808416148015611b0e57505f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff16145b15611c4a57611b28828d8d611f119290919263ffffffff16565b809350819a50505089611c49575f611b4b838e8e611f589290919263ffffffff16565b8162ffffff16915080945081925050505f8d8d85908487611b6c919061430d565b92611b799392919061484d565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090508a73ffffffffffffffffffffffffffffffffffffffff1663ccce3bc830836040518363ffffffff1660e01b8152600401611bf8929190614ae6565b6040805180830381865afa158015611c12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c369190614b89565b92508184611c44919061430d565b935050505b5b600180841603611c8357611c718d8a838f8f87908092611c6c9392919061484d565b61230b565b97509750975097509750505050611de1565b6002808416148d60200190151590811515815250505f6002601c8516901c9050611cbf83828f8f611eaa9190939291909392919063ffffffff16565b8094508197505050505f6001600560208616901c611cdd919061430d565b9050611cfb83828f8f611eaa9190939291909392919063ffffffff16565b809450819a50505050611d0d8d611a08565b9350611d2b8d858e8e86908092611d269392919061484d565b612570565b8097508198505050611d3f86895f1b6130ed565b9550611d4d86865f1b6130ed565b9550611d71868a73ffffffffffffffffffffffffffffffffffffffff165f1b6130ed565b95505f5f1b815f015114158015611d8b575085815f015114155b8015611d9b575080602001518511155b15611ddd57806040517fccbb534f000000000000000000000000000000000000000000000000000000008152600401611dd49190614be1565b60405180910390fd5b5050505b9550955095509550959050565b5f611df882613101565b9050919050565b611e587fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b846bffffffffffffffffffffffff841660608673ffffffffffffffffffffffffffffffffffffffff16901b175f1b611ffe565b505050565b80305550565b5f5f83358060f81c925060019150509250929050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f858401356008840261010003600180866008021b0382821c1693508486019250505094509492505050565b5f5f8483013561ffff8160f01c16925060028401915050935093915050565b5f5f848301358060f81c925060018401915050935093915050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f848301359150602083019050935093915050565b5f5f8483013562ffffff8160e81c16925060038401915050935093915050565b611fa97f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b835f1b611ffe565b5050565b5f5f5f8351602085018787f490509392505050565b5f5f5f835160208501878988f19050949350505050565b60603d604051915060208201818101604052818352815f823e505090565b8082555050565b5f838360405160200161201292919061479b565b60405160208183030381529060405280519060200120905081815550505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856120a257466120a4565b5f5b856040516020016120b9959493929190614bfa565b60405160208183030381529060405280519060200120905092915050565b5f5f8261010001516040516020016120ef9190614cd7565b6040516020818303038152906040528051906020012090505f60ff16835f015160ff1603612187575f6121258460400151613144565b90507f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a2818560600151866080015185604051602001612168959493929190614ced565b6040516020818303038152906040528051906020012092505050612306565b600160ff16835f015160ff16036121f6577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360a0015180519060200120826040516020016121d893929190614d3e565b60405160208183030381529060405280519060200120915050612306565b600260ff16835f015160ff160361225e577f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e48360c001518260405160200161224093929190614d3e565b60405160208183030381529060405280519060200120915050612306565b600360ff16835f015160ff16036122c6577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360e00151826040516020016122a893929190614d3e565b60405160208183030381529060405280519060200120915050612306565b825f01516040517f048183200000000000000000000000000000000000000000000000000000000081526004016122fd9190614d82565b60405180910390fd5b919050565b5f5f5f5f5f61231861339e565b6002815f019060ff16908160ff16815250505f5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b89899050821015612505575f5f612372848d8d611f589290919263ffffffff16565b8162ffffff1691508095508192505050838161238e919061430d565b9150505f8b8b905082146123a2575f6123a4565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303612404576123eb8f8d8d879086926123e39392919061484d565b600185611ab1565b809a50819b50829c50839d50849e505050505050612434565b612422858d8d8790869261241a9392919061484d565b600185611ab1565b50809a50819b50829c50839d50505050505b8989101561248f578b8b8590849261244e9392919061484d565b8b8b6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016124869493929190614d9b565b60405180910390fd5b819350878d5f0151036124a8575f5f1b8d5f0181815250505b8287106124ee5786836040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004016124e592919061481e565b60405180910390fd5b878560c00181815250508692508193505050612350565b5f5f1b8b5f01511415801561251e57508a602001518511155b15612560578a6040517fccbb534f0000000000000000000000000000000000000000000000000000000081526004016125579190614be1565b60405180910390fd5b5050509550955095509550959050565b5f5f5f5b848490508110156130e3575f612595828787611ef69290919263ffffffff16565b8160ff16915080935081925050505f600460f08316901c90505f81036126ed575f600f831690505f8160ff16036125e4576125db848989611ef69290919263ffffffff16565b80955081925050505b5f5f6125fb868b8b6131be9290919263ffffffff16565b8097508193505050612618868b8b6131be9290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f8388866040515f815260200160405260405161267e9493929190614dd9565b6020604051602081039080840390855afa15801561269e573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f6126c0828960ff166131d4565b90505f5f1b8c036126d157806126dc565b6126db8c826130ed565b5b9b5050505050505050505050612574565b60018103612778575f600f831690505f8160ff16036127245761271b848989611ef69290919263ffffffff16565b80955081925050505b5f61273a858a8a611f119290919263ffffffff16565b80965081925050505f612750828460ff166131d4565b90505f5f1b8703612761578061276c565b61276b87826130ed565b5b96505050505050612574565b60028103612977575f6003831690505f8160ff16036127af576127a6848989611ef69290919263ffffffff16565b80955081925050505b5f6127c5858a8a611f119290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f6127fb87838d8d611eaa9190939291909392919063ffffffff16565b80985081925050505f81880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261285f9392919061484d565b6040518463ffffffff1660e01b815260040161287d93929190614e1c565b602060405180830381865afa158015612898573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128bc9190614e60565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614612933578c848d8d8b9085926128f49392919061484d565b6040517fb2fed7ae00000000000000000000000000000000000000000000000000000000815260040161292a9493929190614e8b565b60405180910390fd5b8097508460ff168a0199505f61294c858760ff166131d4565b90505f5f1b8a0361295d5780612968565b6129678a826130ed565b5b99505050505050505050612574565b600381036129c1575f6129958489896131be9290919263ffffffff16565b80955081925050505f5f1b85036129ac57806129b7565b6129b685826130ed565b5b9450505050612574565b60048103612a43575f600f831660ff1690505f6129f085838b8b611eaa9190939291909392919063ffffffff16565b80965081925050505f81860190505f5f612a1c8e8e8e8e8c908892612a179392919061484d565b612570565b91509150829750818a019950612a3289826130ed565b985082975050505050505050612574565b60068103612b53575f6002600c841660ff16901c60ff1690505f8103612a8757612a78848989611ef69290919263ffffffff16565b8160ff16915080955081925050505b5f6003841660ff1690505f8103612abd57612aad858a8a611ed79290919263ffffffff16565b8161ffff16915080965081925050505b5f612ad3868b8b611f589290919263ffffffff16565b8162ffffff16915080975081925050505f81870190505f5f612b078f8f8f8f8d908892612b029392919061484d565b612570565b91509150829850848210612b1b57858b019a505b5f612b27828789613206565b90505f5f1b8b03612b385780612b43565b612b428b826130ed565b5b9a50505050505050505050612574565b60058103612bd5575f612b718489896131be9290919263ffffffff16565b8095508192505050888103612ba4577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b5f612bae8261323b565b90505f5f1b8603612bbf5780612bca565b612bc986826130ed565b5b955050505050612574565b60078103612d3b575f600f831690505f8160ff1603612c0c57612c03848989611ef69290919263ffffffff16565b80955081925050505b5f5f612c23868b8b6131be9290919263ffffffff16565b8097508193505050612c40868b8b6131be9290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f604051602001612c979190614f13565b604051602081830303815290604052805190602001208388866040515f8152602001604052604051612ccc9493929190614dd9565b6020604051602081039080840390855afa158015612cec573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612d0e828960ff166131d4565b90505f5f1b8c03612d1f5780612d2a565b612d298c826130ed565b5b9b5050505050505050505050612574565b60088103612dd4575f612d598489896131be9290919263ffffffff16565b80955081925050505f612d755f8c61326a90919063ffffffff16565b9050808203612da2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b5f612dac826132bb565b90505f5f1b8703612dbd5780612dc8565b612dc787826130ed565b5b96505050505050612574565b60098103612f3d575f6003831690505f8160ff1603612e0b57612e02848989611ef69290919263ffffffff16565b80955081925050505b5f612e21858a8a611f119290919263ffffffff16565b80965081925050505f5f6002600c871660ff16901c60ff169050612e5787828d8d611eaa9190939291909392919063ffffffff16565b8098508193505050505f81870190505f8373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612e969392919061484d565b6040518463ffffffff1660e01b8152600401612eb493929190614f38565b602060405180830381865afa158015612ecf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ef39190614f6f565b90508197508460ff168a0199505f612f0f858760ff16846132ea565b90505f5f1b8a03612f205780612f2b565b612f2a8a826130ed565b5b99508298505050505050505050612574565b600a81036130a6575f6003831690505f8160ff1603612f7457612f6b848989611ef69290919263ffffffff16565b80955081925050505b5f612f8a858a8a611f119290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f612fc087838d8d611eaa9190939291909392919063ffffffff16565b80985081925050505f81880190505f8473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612ffe9392919061484d565b6040518463ffffffff1660e01b815260040161301c93929190614e1c565b602060405180830381865afa158015613037573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061305b9190614f6f565b90508198508560ff168b019a505f613077868860ff16846132ea565b90505f5f1b8b036130885780613093565b6130928b826130ed565b5b9a50829950505050505050505050612574565b806040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004016130da9190613fa1565b60405180910390fd5b5094509492505050565b5f825f528160205260405f20905092915050565b5f5f5f1b821415801561313d575061313a7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b6119fe565b82145b9050919050565b5f60605f5f90505b83518110156131ad575f61317985838151811061316c5761316b614340565b5b602002602001015161331f565b9050828160405160200161318e929190614fca565b604051602081830303815290604052925050808060010191505061314c565b508080519060200120915050919050565b5f5f848301359150602083019050935093915050565b5f82826040516020016131e89291906150a0565b60405160208183030381529060405280519060200120905092915050565b5f83838360405160200161321c93929190615120565b6040516020818303038152906040528051906020012090509392505050565b5f8160405160200161324d91906151b1565b604051602081830303815290604052805190602001209050919050565b5f5f61327a846020015184612033565b90505f613286856120d7565b9050818160405160200161329b929190614ab0565b604051602081830303815290604052805190602001209250505092915050565b5f816040516020016132cd9190615220565b604051602081830303815290604052805190602001209050919050565b5f8383836040516020016133009392919061528f565b6040516020818303038152906040528051906020012090509392505050565b5f7f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437825f01518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016133819897969594939291906152d6565b604051602081830303815290604052805190602001209050919050565b6040518061012001604052805f60ff1681526020015f15158152602001606081526020015f81526020015f8152602001606081526020015f81526020015f8152602001606081525090565b6040518060e001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020015f81526020015f151581526020015f151581526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f82905092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f82821b905092915050565b5f61349e8383613452565b826134a9813561345c565b925060048210156134e9576134e47fffffffff0000000000000000000000000000000000000000000000000000000083600403600802613487565b831692505b505092915050565b5f81905092915050565b828183375f83830152505050565b5f61351483856134f1565b93506135218385846134fb565b82840190509392505050565b5f613539828486613509565b91508190509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61357f82613556565b9050919050565b61358f81613575565b8114613599575f5ffd5b50565b5f813590506135aa81613586565b92915050565b5f602082840312156135c5576135c461354e565b5b5f6135d28482850161359c565b91505092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613625826135df565b810181811067ffffffffffffffff82111715613644576136436135ef565b5b80604052505050565b5f613656613545565b9050613662828261361c565b919050565b5f5ffd5b5f60ff82169050919050565b6136808161366b565b811461368a575f5ffd5b50565b5f8135905061369b81613677565b92915050565b5f8115159050919050565b6136b5816136a1565b81146136bf575f5ffd5b50565b5f813590506136d0816136ac565b92915050565b5f5ffd5b5f67ffffffffffffffff8211156136f4576136f36135ef565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b61371b81613709565b8114613725575f5ffd5b50565b5f8135905061373681613712565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561375a576137596135ef565b5b613763826135df565b9050602081019050919050565b5f61378261377d84613740565b61364d565b90508281526020810184848401111561379e5761379d61373c565b5b6137a98482856134fb565b509392505050565b5f82601f8301126137c5576137c46136d6565b5b81356137d5848260208601613770565b91505092915050565b5f60e082840312156137f3576137f26135db565b5b6137fd60e061364d565b90505f61380c8482850161359c565b5f83015250602061381f84828501613728565b602083015250604082013567ffffffffffffffff81111561384357613842613667565b5b61384f848285016137b1565b604083015250606061386384828501613728565b6060830152506080613877848285016136c2565b60808301525060a061388b848285016136c2565b60a08301525060c061389f84828501613728565b60c08301525092915050565b5f6138bd6138b8846136da565b61364d565b905080838252602082019050602084028301858111156138e0576138df613705565b5b835b8181101561392757803567ffffffffffffffff811115613905576139046136d6565b5b80860161391289826137de565b855260208501945050506020810190506138e2565b5050509392505050565b5f82601f830112613945576139446136d6565b5b81356139558482602086016138ab565b91505092915050565b5f819050919050565b6139708161395e565b811461397a575f5ffd5b50565b5f8135905061398b81613967565b92915050565b5f67ffffffffffffffff8211156139ab576139aa6135ef565b5b602082029050602081019050919050565b5f6139ce6139c984613991565b61364d565b905080838252602082019050602084028301858111156139f1576139f0613705565b5b835b81811015613a1a5780613a06888261359c565b8452602084019350506020810190506139f3565b5050509392505050565b5f82601f830112613a3857613a376136d6565b5b8135613a488482602086016139bc565b91505092915050565b5f6101208284031215613a6757613a666135db565b5b613a7261012061364d565b90505f613a818482850161368d565b5f830152506020613a94848285016136c2565b602083015250604082013567ffffffffffffffff811115613ab857613ab7613667565b5b613ac484828501613931565b6040830152506060613ad884828501613728565b6060830152506080613aec84828501613728565b60808301525060a082013567ffffffffffffffff811115613b1057613b0f613667565b5b613b1c848285016137b1565b60a08301525060c0613b308482850161397d565b60c08301525060e0613b448482850161397d565b60e08301525061010082013567ffffffffffffffff811115613b6957613b68613667565b5b613b7584828501613a24565b6101008301525092915050565b5f5ffd5b5f5f83601f840112613b9b57613b9a6136d6565b5b8235905067ffffffffffffffff811115613bb857613bb7613b82565b5b602083019150836001820283011115613bd457613bd3613705565b5b9250929050565b5f5f5f60408486031215613bf257613bf161354e565b5b5f84013567ffffffffffffffff811115613c0f57613c0e613552565b5b613c1b86828701613a51565b935050602084013567ffffffffffffffff811115613c3c57613c3b613552565b5b613c4886828701613b86565b92509250509250925092565b613c5d8161395e565b82525050565b5f602082019050613c765f830184613c54565b92915050565b5f5f5f5f5f60808688031215613c9557613c9461354e565b5b5f613ca28882890161359c565b9550506020613cb38882890161359c565b9450506040613cc488828901613728565b935050606086013567ffffffffffffffff811115613ce557613ce4613552565b5b613cf188828901613b86565b92509250509295509295909350565b613d098161345c565b82525050565b5f602082019050613d225f830184613d00565b92915050565b5f5f5f60408486031215613d3f57613d3e61354e565b5b5f613d4c8682870161397d565b935050602084013567ffffffffffffffff811115613d6d57613d6c613552565b5b613d7986828701613b86565b92509250509250925092565b613d8e8161345c565b8114613d98575f5ffd5b50565b5f81359050613da981613d85565b92915050565b5f60208284031215613dc457613dc361354e565b5b5f613dd184828501613d9b565b91505092915050565b613de381613575565b82525050565b5f602082019050613dfc5f830184613dda565b92915050565b5f5f5f5f60408587031215613e1a57613e1961354e565b5b5f85013567ffffffffffffffff811115613e3757613e36613552565b5b613e4387828801613b86565b9450945050602085013567ffffffffffffffff811115613e6657613e65613552565b5b613e7287828801613b86565b925092505092959194509250565b5f60208284031215613e9557613e9461354e565b5b5f613ea28482850161397d565b91505092915050565b5f5f60208385031215613ec157613ec061354e565b5b5f83013567ffffffffffffffff811115613ede57613edd613552565b5b613eea85828601613b86565b92509250509250929050565b5f5f5f5f60608587031215613f0e57613f0d61354e565b5b5f613f1b8782880161359c565b9450506020613f2c87828801613728565b935050604085013567ffffffffffffffff811115613f4d57613f4c613552565b5b613f5987828801613b86565b925092505092959194509250565b5f60208284031215613f7c57613f7b61354e565b5b5f613f8984828501613728565b91505092915050565b613f9b81613709565b82525050565b5f602082019050613fb45f830184613f92565b92915050565b5f604082019050613fcd5f830185613dda565b613fda6020830184613f92565b9392505050565b613fea816136a1565b82525050565b5f60c0820190506140035f830189613f92565b6140106020830188613f92565b61401d6040830187613fe1565b61402a6060830186613c54565b6140376080830185613f92565b61404460a0830184613c54565b979650505050505050565b5f5f604083850312156140655761406461354e565b5b5f61407285828601613d9b565b92505060206140838582860161359c565b9150509250929050565b5f5f83601f8401126140a2576140a16136d6565b5b8235905067ffffffffffffffff8111156140bf576140be613b82565b5b6020830191508360208202830111156140db576140da613705565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b0312156140fe576140fd61354e565b5b5f61410b8b828c0161359c565b985050602061411c8b828c0161359c565b975050604089013567ffffffffffffffff81111561413d5761413c613552565b5b6141498b828c0161408d565b9650965050606089013567ffffffffffffffff81111561416c5761416b613552565b5b6141788b828c0161408d565b9450945050608089013567ffffffffffffffff81111561419b5761419a613552565b5b6141a78b828c01613b86565b92509250509295985092959890939650565b5f5f5f5f5f5f60a087890312156141d3576141d261354e565b5b5f6141e089828a0161359c565b96505060206141f189828a0161359c565b955050604061420289828a01613728565b945050606061421389828a01613728565b935050608087013567ffffffffffffffff81111561423457614233613552565b5b61424089828a01613b86565b92509250509295509295509295565b5f6bffffffffffffffffffffffff82169050919050565b61426f8161424f565b8114614279575f5ffd5b50565b5f8135905061428a81614266565b92915050565b5f5f5f606084860312156142a7576142a661354e565b5b5f6142b48682870161397d565b93505060206142c58682870161359c565b92505060406142d68682870161427c565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61431782613709565b915061432283613709565b925082820190508082111561433a576143396142e0565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b6143768161366b565b82525050565b614385816136a1565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6143bd81613575565b82525050565b6143cc81613709565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614404826143d2565b61440e81856143dc565b935061441e8185602086016143ec565b614427816135df565b840191505092915050565b5f60e083015f8301516144475f8601826143b4565b50602083015161445a60208601826143c3565b506040830151848203604086015261447282826143fa565b915050606083015161448760608601826143c3565b50608083015161449a608086018261437c565b5060a08301516144ad60a086018261437c565b5060c08301516144c060c08601826143c3565b508091505092915050565b5f6144d68383614432565b905092915050565b5f602082019050919050565b5f6144f48261438b565b6144fe8185614395565b935083602082028501614510856143a5565b805f5b8581101561454b578484038952815161452c85826144cb565b9450614537836144de565b925060208a01995050600181019050614513565b50829750879550505050505092915050565b6145668161395e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6145a083836143b4565b60208301905092915050565b5f602082019050919050565b5f6145c28261456c565b6145cc8185614576565b93506145d783614586565b805f5b838110156146075781516145ee8882614595565b97506145f9836145ac565b9250506001810190506145da565b5085935050505092915050565b5f61012083015f83015161462a5f86018261436d565b50602083015161463d602086018261437c565b506040830151848203604086015261465582826144ea565b915050606083015161466a60608601826143c3565b50608083015161467d60808601826143c3565b5060a083015184820360a086015261469582826143fa565b91505060c08301516146aa60c086018261455d565b5060e08301516146bd60e086018261455d565b506101008301518482036101008601526146d782826145b8565b9150508091505092915050565b5f82825260208201905092915050565b5f6146ff83856146e4565b935061470c8385846134fb565b614715836135df565b840190509392505050565b5f6040820190508181035f8301526147388186614614565b9050818103602083015261474d8184866146f4565b9050949350505050565b6147608161424f565b82525050565b5f6060820190506147795f830186613c54565b6147866020830185613dda565b6147936040830184614757565b949350505050565b5f6040820190506147ae5f830185613c54565b6147bb6020830184613c54565b9392505050565b5f6040820190506147d55f830185613c54565b6147e26020830184613f92565b9392505050565b5f6060820190506147fc5f830186613c54565b6148096020830185613dda565b6148166040830184613dda565b949350505050565b5f6040820190506148315f830185613f92565b61483e6020830184613f92565b9392505050565b5f5ffd5b5f5ffd5b5f5f858511156148605761485f614845565b5b8386111561487157614870614849565b5b6001850283019150848603905094509492505050565b5f60608201905061489a5f830186613f92565b6148a76020830185613f92565b6148b46040830184613f92565b949350505050565b5f6060820190508181035f8301526148d48186614614565b90506148e36020830185613f92565b6148f06040830184613f92565b949350505050565b5f614902826143d2565b61490c81856146e4565b935061491c8185602086016143ec565b614925816135df565b840191505092915050565b5f60c0820190506149435f830189613c54565b6149506020830188613f92565b61495d6040830187613f92565b61496a6060830186613f92565b6149776080830185613f92565b81810360a083015261498981846148f8565b9050979650505050505050565b5f6060820190506149a95f830186613c54565b6149b66020830185613f92565b81810360408301526149c881846148f8565b9050949350505050565b5f6060820190508181035f8301526149ea8186614614565b90506149f96020830185613f92565b8181036040830152614a0b81846148f8565b9050949350505050565b5f604082019050614a285f830185613d00565b614a356020830184613dda565b9392505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f614a7a600283614a3c565b9150614a8582614a46565b600282019050919050565b5f819050919050565b614aaa614aa58261395e565b614a90565b82525050565b5f614aba82614a6e565b9150614ac68285614a99565b602082019150614ad68284614a99565b6020820191508190509392505050565b5f604082019050614af95f830185613dda565b8181036020830152614b0b81846148f8565b90509392505050565b5f81519050614b2281613967565b92915050565b5f81519050614b3681613712565b92915050565b5f60408284031215614b5157614b506135db565b5b614b5b604061364d565b90505f614b6a84828501614b14565b5f830152506020614b7d84828501614b28565b60208301525092915050565b5f60408284031215614b9e57614b9d61354e565b5b5f614bab84828501614b3c565b91505092915050565b604082015f820151614bc85f85018261455d565b506020820151614bdb60208501826143c3565b50505050565b5f604082019050614bf45f830184614bb4565b92915050565b5f60a082019050614c0d5f830188613c54565b614c1a6020830187613c54565b614c276040830186613c54565b614c346060830185613f92565b614c416080830184613dda565b9695505050505050565b5f81905092915050565b614c5e81613575565b82525050565b5f614c6f8383614c55565b60208301905092915050565b5f614c858261456c565b614c8f8185614c4b565b9350614c9a83614586565b805f5b83811015614cca578151614cb18882614c64565b9750614cbc836145ac565b925050600181019050614c9d565b5085935050505092915050565b5f614ce28284614c7b565b915081905092915050565b5f60a082019050614d005f830188613c54565b614d0d6020830187613c54565b614d1a6040830186613f92565b614d276060830185613f92565b614d346080830184613c54565b9695505050505050565b5f606082019050614d515f830186613c54565b614d5e6020830185613c54565b614d6b6040830184613c54565b949350505050565b614d7c8161366b565b82525050565b5f602082019050614d955f830184614d73565b92915050565b5f6060820190508181035f830152614db48186886146f4565b9050614dc36020830185613f92565b614dd06040830184613f92565b95945050505050565b5f608082019050614dec5f830187613c54565b614df96020830186614d73565b614e066040830185613c54565b614e136060830184613c54565b95945050505050565b5f604082019050614e2f5f830186613c54565b8181036020830152614e428184866146f4565b9050949350505050565b5f81519050614e5a81613d85565b92915050565b5f60208284031215614e7557614e7461354e565b5b5f614e8284828501614e4c565b91505092915050565b5f606082019050614e9e5f830187613c54565b614eab6020830186613dda565b8181036040830152614ebe8184866146f4565b905095945050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f614efd601c83614a3c565b9150614f0882614ec9565b601c82019050919050565b5f614f1d82614ef1565b9150614f298284614a99565b60208201915081905092915050565b5f6040820190508181035f830152614f508186614614565b90508181036020830152614f658184866146f4565b9050949350505050565b5f60208284031215614f8457614f8361354e565b5b5f614f9184828501614b14565b91505092915050565b5f614fa4826143d2565b614fae81856134f1565b9350614fbe8185602086016143ec565b80840191505092915050565b5f614fd58285614f9a565b9150614fe18284614a99565b6020820191508190509392505050565b7f53657175656e6365207369676e65723a0a0000000000000000000000000000005f82015250565b5f615025601183614a3c565b915061503082614ff1565b601182019050919050565b5f8160601b9050919050565b5f6150518261503b565b9050919050565b5f61506282615047565b9050919050565b61507a61507582613575565b615058565b82525050565b5f819050919050565b61509a61509582613709565b615080565b82525050565b5f6150aa82615019565b91506150b68285615069565b6014820191506150c68284615089565b6020820191508190509392505050565b7f53657175656e6365206e657374656420636f6e6669673a0a00000000000000005f82015250565b5f61510a601883614a3c565b9150615115826150d6565b601882019050919050565b5f61512a826150fe565b91506151368286614a99565b6020820191506151468285615089565b6020820191506151568284615089565b602082019150819050949350505050565b7f53657175656e636520737461746963206469676573743a0a00000000000000005f82015250565b5f61519b601883614a3c565b91506151a682615167565b601882019050919050565b5f6151bb8261518f565b91506151c78284614a99565b60208201915081905092915050565b7f53657175656e636520616e792061646472657373207375626469676573743a0a5f82015250565b5f61520a602083614a3c565b9150615215826151d6565b602082019050919050565b5f61522a826151fe565b91506152368284614a99565b60208201915081905092915050565b7f53657175656e63652073617069656e7420636f6e6669673a0a000000000000005f82015250565b5f615279601983614a3c565b915061528482615245565b601982019050919050565b5f6152998261526d565b91506152a58286615069565b6014820191506152b58285615089565b6020820191506152c58284614a99565b602082019150819050949350505050565b5f610100820190506152ea5f83018b613c54565b6152f7602083018a613dda565b6153046040830189613f92565b6153116060830188613c54565b61531e6080830187613f92565b61532b60a0830186613fe1565b61533860c0830185613fe1565b61534560e0830184613f92565b999850505050505050505056fea26469706673582212201a0a84c32c1e19e877ba9e1007435d7c5f2f6c129905addd2ceff9fc24b9016164736f6c634300081c0033603e600e3d39601e805130553df33d3d34601c57363d3d373d363d30545af43d82803e903d91601c57fd5bf3",
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

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage1 *WalletStage1Caller) GetStaticSignature(opts *bind.CallOpts, _hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "getStaticSignature", _hash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage1 *WalletStage1Session) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage1.Contract.GetStaticSignature(&_WalletStage1.CallOpts, _hash)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage1 *WalletStage1CallerSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage1.Contract.GetStaticSignature(&_WalletStage1.CallOpts, _hash)
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
