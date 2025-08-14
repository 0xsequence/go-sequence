// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletsimulator

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

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

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

// SimulatorResult is an auto generated low-level Go binding around an user-defined struct.
type SimulatorResult struct {
	Status  uint8
	Result  []byte
	GasUsed *big.Int
}

// WalletSimulatorMetaData contains all meta data concerning the WalletSimulator contract.
var WalletSimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"CreateFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC4337Disabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidERC1271Signature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_entrypoint\",\"type\":\"address\"}],\"name\":\"InvalidEntryPoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_expectedCaller\",\"type\":\"address\"}],\"name\":\"InvalidStaticSignatureWrongCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"CreatedContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_code\",\"type\":\"bytes\"}],\"name\":\"createContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entrypoint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"executeUserOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getStaticSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"simulate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSimulator.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"internalType\":\"structSimulator.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"operatorData\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051614e33380380614e3383398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051614d856100ae600039600081816105590152818161094a015281816109cd01528181610a5a01528181610eef0152610f720152614d856000f3fe6080604052600436106101995760003560e01c80638943ec02116100e1578063a65d69d41161008a578063b93ea7ad11610064578063b93ea7ad146105e3578063bc197c81146105f6578063f23a6e611461063e578063f727ef1c14610684576101a0565b8063a65d69d414610547578063aaf10f421461057b578063ad55366b14610590576101a0565b806392dcb3fc116100bb57806392dcb3fc146104ae5780639c145aed146104fa578063a27c39221461051a576101a0565b80638943ec02146104375780638c3f55631461047b57806390042baf1461049b576101a0565b80631a9b2337116101435780634fcf3eca1161011d5780634fcf3eca146103fc57806351605d801461040f5780636ea4457714610424576101a0565b80631a9b2337146103845780631f6a1eb9146103c957806329561426146103dc576101a0565b8063150b7a0211610174578063150b7a02146102ce5780631626ba7e1461034457806319822f7c14610364576101a0565b806223de2914610263578063025b22bc1461028857806313792a4a1461029b576101a0565b366101a057005b600436106102615760006101bc6101b73683613a9f565b6106a4565b905073ffffffffffffffffffffffffffffffffffffffff81161561025f576000808273ffffffffffffffffffffffffffffffffffffffff16600036604051610205929190613b05565b600060405180830381855af49150503d8060008114610240576040519150601f19603f3d011682016040523d82523d6000602084013e610245565b606091505b50915091508161025757805160208201fd5b805160208201f35b505b005b34801561026f57600080fd5b5061026161027e366004613b87565b5050505050505050565b610261610296366004613c37565b6106f8565b3480156102a757600080fd5b506102bb6102b6366004613f82565b610744565b6040519081526020015b60405180910390f35b3480156102da57600080fd5b506103136102e93660046140c9565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016102c5565b34801561035057600080fd5b5061031361035f366004614138565b6108af565b34801561037057600080fd5b506102bb61037f36600461416b565b610946565b34801561039057600080fd5b506103a461039f3660046141ed565b610b9c565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c5565b6102616103d736600461420a565b610ba7565b3480156103e857600080fd5b506102616103f736600461427b565b610c26565b61026161040a3660046141ed565b610c6a565b34801561041b57600080fd5b506102bb610d2c565b610261610432366004614294565b610d5b565b34801561044357600080fd5b506103136104523660046142d6565b7f8943ec0200000000000000000000000000000000000000000000000000000000949350505050565b34801561048757600080fd5b506102bb61049636600461427b565b610dc8565b6103a46104a9366004614318565b610df4565b3480156104ba57600080fd5b506104ce6104c936600461427b565b610ed8565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016102c5565b34801561050657600080fd5b50610261610515366004614294565b610eed565b34801561052657600080fd5b5061053a61053536600461439a565b61103f565b6040516102c59190614463565b34801561055357600080fd5b506103a47f000000000000000000000000000000000000000000000000000000000000000081565b34801561058757600080fd5b506103a461157c565b34801561059c57600080fd5b506105b06105ab366004613f82565b611586565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c0016102c5565b6102616105f1366004614541565b6115c0565b34801561060257600080fd5b50610313610611366004614576565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561064a57600080fd5b5061031361065936600461461d565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561069057600080fd5b5061026161069f366004614695565b611685565b60006106f27fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416611740565b92915050565b333014610738576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6107418161179e565b50565b60008084610100015151600161075a9190614715565b67ffffffffffffffff81111561077257610772613c52565b60405190808252806020026020018201604052801561079b578160200160208202803683370190505b50905060005b8561010001515181101561080d5785610100015181815181106107c6576107c6614728565b60200260200101518282815181106107e0576107e0614728565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016107a1565b503381866101000151518151811061082757610827614728565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006108618686866117f3565b509050806108a1578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161072f9392919061496c565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905260006109078286866117f3565b5090508061091b5750600091506108a89050565b507f1626ba7e0000000000000000000000000000000000000000000000000000000095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166109b5576040517fd13d783500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a26576040517f1d6ddbf400000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b8115610ace576040517fb760faf90000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b760faf99084906024016000604051808303818588803b158015610ab457600080fd5b505af1158015610ac8573d6000803e3d6000fd5b50505050505b7f1626ba7e0000000000000000000000000000000000000000000000000000000030631626ba7e85610b0461010089018961499c565b6040518463ffffffff1660e01b8152600401610b2293929190614a01565b602060405180830381865afa158015610b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b639190614a1b565b7fffffffff000000000000000000000000000000000000000000000000000000001614610b92575060016108a8565b5060009392505050565b60006106f2826106a4565b60005a90506000610bb886866119da565b9050610bcc81606001518260800151611df8565b600080610bda8387876117f3565b9150915081610c1b578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161072f9392919061496c565b61027e848285611ee0565b333014610c61576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b610741816121b5565b333014610ca5576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6000610cb0826106a4565b73ffffffffffffffffffffffffffffffffffffffff1603610d21576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008216600482015260240161072f565b610741816000612245565b6000610d567fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610d96576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b60005a90506000610da784846119da565b90506000610db482612305565b9050610dc1838284611ee0565b5050505050565b60006106f27f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83611740565b6000333014610e31576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b81516020830134f0905073ffffffffffffffffffffffffffffffffffffffff8116610e8a57816040517f0d25719100000000000000000000000000000000000000000000000000000000815260040161072f9190614a38565b60405173ffffffffffffffffffffffffffffffffffffffff821681527fa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c9060200160405180910390a1919050565b600080610ee483612380565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f5a576040517fd13d783500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610fcb576040517f1d6ddbf400000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6040517f6ea445770000000000000000000000000000000000000000000000000000000081523090636ea44577906110099085908590600401614a4b565b600060405180830381600087803b15801561102357600080fd5b505af1158015611037573d6000803e3d6000fd5b505050505050565b606060005a90506000838067ffffffffffffffff81111561106257611062613c52565b6040519080825280602002602001820160405280156110b857816020015b6110a56040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816110805790505b50935060005b818110156115725760008787838181106110da576110da614728565b90506020028101906110ec9190614a5f565b6110f590614a9d565b90508060a001518015611106575083155b15611111575061156a565b60608101516000945080158015906111285750805a105b156111c257600587848151811061114157611141614728565b602002602001015160000190600581111561115e5761115e6143d0565b90816005811115611171576111716143d0565b9052505a60405160200161118791815260200190565b6040516020818303038152906040528784815181106111a8576111a8614728565b6020026020010151602001819052505050505050506106f2565b60008260800151156112c85760005a84519091506112939084156111e657846111e8565b5a5b634c4e814c60e01b60008c8a8c60008c6040015160405160240161121196959493929190614aa9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526123cc565b91505a6112a09082614aec565b8986815181106112b2576112b2614728565b6020026020010151604001818152505050611322565b60005a845160208601519192506112f19185156112e557856112e7565b5a5b87604001516123e2565b91505a6112fe9082614aec565b89868151811061131057611310614728565b60200260200101516040018181525050505b806114f65760c08301516113ac5760019550600288858151811061134857611348614728565b6020026020010151600001906005811115611365576113656143d0565b90816005811115611378576113786143d0565b9052506113836123fa565b88858151811061139557611395614728565b60200260200101516020018190525050505061156a565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114535760048885815181106113eb576113eb614728565b6020026020010151600001906005811115611408576114086143d0565b9081600581111561141b5761141b6143d0565b9052506114266123fa565b88858151811061143857611438614728565b602002602001015160200181905250505050505050506106f2565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016114f657600388858151811061149257611492614728565b60200260200101516000019060058111156114af576114af6143d0565b908160058111156114c2576114c26143d0565b9052506114cd6123fa565b8885815181106114df576114df614728565b602002602001015160200181905250505050611572565b600188858151811061150a5761150a614728565b6020026020010151600001906005811115611527576115276143d0565b9081600581111561153a5761153a6143d0565b9052506115456123fa565b88858151811061155757611557614728565b6020026020010151602001819052505050505b6001016110be565b5050505092915050565b6000610d56305490565b60008060008060008061159d898989600080612419565b9399509197509450925090506115b28361273e565b935093975093979195509350565b3330146115fb576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6000611606836106a4565b73ffffffffffffffffffffffffffffffffffffffff1614611677576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161072f565b6116818282612245565b5050565b3330146116c0576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6116d98383836bffffffffffffffffffffffff16612749565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b600080838360405160200161175f929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b6117a6813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b60008060008484600081811061180b5761180b614728565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f800000000000000000000000000000000000000000000000000000000000000090811690036119605761186886612305565b915060008061187684612380565b915091504281116118bd576040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161072f565b73ffffffffffffffffffffffffffffffffffffffff8216158015906118f8575073ffffffffffffffffffffffffffffffffffffffff82163314155b15611954576040517f8945c3130000000000000000000000000000000000000000000000000000000081526004810185905233602482015273ffffffffffffffffffffffffffffffffffffffff8316604482015260640161072f565b600194505050506119d2565b6000806000611973898989600080612419565b9850929550909350915050828210156119c2576040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161072f565b6119cb8161273e565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103611a3e5760006060840152611a4f565b84810135606090811c908401526014015b6007600183901c168015611aa25760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003611ab757506001611ae0565b83602016602003611ad35750600282019186013560f01c611ae0565b50600182019186013560f81c5b8067ffffffffffffffff811115611af957611af9613c52565b604051908082528060200260200182016040528015611b8457816020015b611b716040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081611b175790505b50604086015260005b81811015611ded5760018085019489013560f81c908082169003611bec573087604001518381518110611bc257611bc2614728565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611c36565b8489013560601c6014860188604001518481518110611c0d57611c0d614728565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611c7457848901356020860188604001518481518110611c5f57611c5f614728565b60200260200101516020018197508281525050505b60048082169003611d0c57600385019489013560e81c89868a611c978483614715565b92611ca493929190614aff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506040890151805185908110611cef57611cef614728565b602090810291909101015160400152611d088187614715565b9550505b60088082169003611d4a57848901356020860188604001518481518110611d3557611d35614728565b60200260200101516060018197508281525050505b8060101660ff1660101487604001518381518110611d6a57611d6a614728565b602002602001015160800190151590811515815250508060201660ff1660201487604001518381518110611da057611da0614728565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110611dd557611dd5614728565b602090810291909101015160c0015250600101611b8d565b505050505092915050565b6000611e0383610dc8565b9050818114611e4f576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161072f565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561103757600084604001518281518110611f0b57611f0b614728565b602002602001015190508060a001518015611f24575083155b15611f685760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506121ad565b6060810151600094508015801590611f7f5750805a105b15611fbc5785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161072f93929190614b29565b600082608001511561200e578251612007908315611fda5783611fdc565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161121196959493929190614b4e565b9050612035565b82516020840151612032919084156120265784612028565b5a5b86604001516123e2565b90505b806121705760c083015161209157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856120726123fa565b60405161208193929190614b7f565b60405180910390a15050506121ad565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016120fb5786846120c66123fa565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161072f93929190614b9e565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612170577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856121516123fa565b60405161216093929190614b7f565b60405180910390a1505050611037565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b600101611eeb565b806121ec576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6122157fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa906020016117e8565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806123168360200151306127d8565b90506000612323846128a5565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b600080806123ae7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685611740565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015612464575073ffffffffffffffffffffffffffffffffffffffff8916155b15612580578b82013560601c9850601490910190896125805760038201918c013560e81c60008d848e6124978583614715565b926124a493929190614aff565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc8915061252f9030908590600401614bc9565b6040805180830381865afa15801561254b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256f9190614bf8565b925061257b8285614715565b935050505b826001166001036125ba576125a88d8a838f8f879080926125a393929190614aff565b612b11565b97509750975097509750505050612731565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c168382019096509250600090506126256001600586901c811690614715565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c16995090920191506126708d612305565b935061268e8d858e8e8690809261268993929190614aff565b612d55565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d16909152902082519198509650158015906126d9575080518614155b80156126e9575080602001518511155b1561272d576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161072f565b5050505b9550955095509550959050565b60006106f28261367c565b6127d37fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85612848574661284b565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000808261010001516040516020016128be9190614c49565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120835190915060ff1661296757600061290f84604001516136af565b606085810151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001612361565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016129f75760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612a675760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016129d9565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01612ad75760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082015290810191909152606081018290526080016129d9565b82516040517f0481832000000000000000000000000000000000000000000000000000000000815260ff909116600482015260240161072f565b6000806000806000612b73604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b88821015612cfd5760038201916000908b013560e81c612bbb8482614715565b9150600090508a8214612bcf576000612bd1565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303612c2957612c188f8d8d87908692612c1093929190614aff565b600185612419565b939d50919b50995097509550612c4b565b612c3f858d8d87908692612c1093929190614aff565b50929c50909a50985096505b89891015612c9757612c5f82858d8f614aff565b8b8b6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161072f9493929190614c95565b819350878d6000015103612caa5760008d525b828710612ced576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018890526024810184905260440161072f565b50505060c0820185905283612b9b565b8a5115801590612d1157508a602001518511155b1561272d576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161072f565b60008060005b8381101561367257600181019085013560f881901c9060fc1c80612e6057600f82166000819003612d935750600183019287013560f81c5b60008080612da28b8b8961375d565b809a5081945082955083965050505050600060018d83868660405160008152602001604052604051612df0949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612e12573d6000803e3d6000fd5b5050506020604051035190508460ff168a0199506000612e35828760ff166137b0565b905089612e425780612e51565b60008a81526020829052604090205b99505050505050505050612d5b565b60018103612ec457600f82166000819003612e825750600183019287013560f81c5b601484019388013560601c6000612e9c8260ff85166137b0565b905086612ea95780612eb8565b60008781526020829052604090205b96505050505050612d5b565b6002810361308e57600382166000819003612ee65750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c16828801809850819250505060008188019050631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d908792612fa493929190614aff565b6040518463ffffffff1660e01b8152600401612fc293929190614a01565b602060405180830381865afa158015612fdf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130039190614a1b565b7fffffffff000000000000000000000000000000000000000000000000000000001614613074578c848d8d8b90859261303e93929190614aff565b6040517fb2fed7ae00000000000000000000000000000000000000000000000000000000815260040161072f9493929190614cbc565b8097508460ff168a0199506000612e35858760ff166137b0565b600381036130c2576020830192870135846130a957806130b8565b60008581526020829052604090205b9450505050612d5b565b6004810361316557600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806131378e8e8e8e8c90889261268993929190614aff565b91509150829750818a019950613157898260009182526020526040902090565b985050505050505050612d5b565b6006810361322f576003600283901c16600081900361318b5750600183019287013560f81c5b6003831660008190036131a55750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806131e38f8f8f8f8d90889261268993929190614aff565b915091508298508482106131f657998501995b6000613203828789613817565b90508a613210578061321f565b60008b81526020829052604090205b9a50505050505050505050612d5b565b6005810361329c57602083019287013588810361326a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b60006132758261387b565b9050856132825780613291565b60008681526020829052604090205b955050505050612d5b565b6007810361337a57600f821660008190036132be5750600183019287013560f81c5b600080806132cd8b8b8961375d565b604051909a50929550909350915060009060019061331b908f906020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018690526080810185905260a001612df0565b600881036133ce57602083019287013560006133968b826138cf565b90508082036133c3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b6000612e9c8361394a565b60098103613506576003821660008190036133f05750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c90879261348a93929190614aff565b6040518463ffffffff1660e01b81526004016134a89392919061496c565b602060405180830381865afa1580156134c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134e99190614cf2565b90508197508460ff168a0199506000612e35858760ff1684613985565b600a810361363d576003821660008190036135285750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d9087926135c193929190614aff565b6040518463ffffffff1660e01b81526004016135df93929190614a01565b602060405180830381865afa1580156135fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136209190614cf2565b90508198508560ff168b019a506000613203868860ff1684613985565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161072f565b5094509492505050565b600081158015906106f25750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b600080825167ffffffffffffffff8111156136cc576136cc613c52565b6040519080825280602002602001820160405280156136f5578160200160208202803683370190505b50905060005b835181101561374b5761372684828151811061371957613719614728565b60200260200101516139f3565b82828151811061373857613738614728565b60209081029190910101526001016136fb565b50806040516020016129d99190614d0b565b828101803590602001357f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811690600090604085019060ff81901c6137a381601b614d36565b9350505093509350935093565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612887565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806138e08460200151846127d8565b905060006138ed856128a5565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a60208201529081018290526000906060016138b2565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d0161385c565b80516020808301516040808501518051908401206060860151608087015160a088015160c089015194516000986138b2987f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef4379891979196959493920197885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015613afe577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613b3957600080fd5b919050565b60008083601f840112613b5057600080fd5b50813567ffffffffffffffff811115613b6857600080fd5b602083019150836020828501011115613b8057600080fd5b9250929050565b60008060008060008060008060c0898b031215613ba357600080fd5b613bac89613b15565b9750613bba60208a01613b15565b9650613bc860408a01613b15565b955060608901359450608089013567ffffffffffffffff811115613beb57600080fd5b613bf78b828c01613b3e565b90955093505060a089013567ffffffffffffffff811115613c1757600080fd5b613c238b828c01613b3e565b999c989b5096995094979396929594505050565b600060208284031215613c4957600080fd5b6108a882613b15565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613ca457613ca4613c52565b60405290565b604051610120810167ffffffffffffffff81118282101715613ca457613ca4613c52565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613d1557613d15613c52565b604052919050565b803560ff81168114613b3957600080fd5b80358015158114613b3957600080fd5b600067ffffffffffffffff821115613d5857613d58613c52565b5060051b60200190565b600082601f830112613d7357600080fd5b813567ffffffffffffffff811115613d8d57613d8d613c52565b613dbe60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613cce565b818152846020838601011115613dd357600080fd5b816020850160208301376000918101602001919091529392505050565b600060e08284031215613e0257600080fd5b613e0a613c81565b9050613e1582613b15565b815260208281013590820152604082013567ffffffffffffffff811115613e3b57600080fd5b613e4784828501613d62565b60408301525060608281013590820152613e6360808301613d2e565b6080820152613e7460a08301613d2e565b60a082015260c09182013591810191909152919050565b600082601f830112613e9c57600080fd5b8135613eaf613eaa82613d3e565b613cce565b8082825260208201915060208360051b860101925085831115613ed157600080fd5b602085015b83811015613f1357803567ffffffffffffffff811115613ef557600080fd5b613f04886020838a0101613df0565b84525060209283019201613ed6565b5095945050505050565b600082601f830112613f2e57600080fd5b8135613f3c613eaa82613d3e565b8082825260208201915060208360051b860101925085831115613f5e57600080fd5b602085015b83811015613f1357613f7481613b15565b835260209283019201613f63565b600080600060408486031215613f9757600080fd5b833567ffffffffffffffff811115613fae57600080fd5b84016101208187031215613fc157600080fd5b613fc9613caa565b613fd282613d1d565b8152613fe060208301613d2e565b6020820152604082013567ffffffffffffffff811115613fff57600080fd5b61400b88828501613e8b565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561403f57600080fd5b61404b88828501613d62565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561408057600080fd5b61408c88828501613f1d565b61010083015250935050602084013567ffffffffffffffff8111156140b057600080fd5b6140bc86828701613b3e565b9497909650939450505050565b6000806000806000608086880312156140e157600080fd5b6140ea86613b15565b94506140f860208701613b15565b935060408601359250606086013567ffffffffffffffff81111561411b57600080fd5b61412788828901613b3e565b969995985093965092949392505050565b60008060006040848603121561414d57600080fd5b83359250602084013567ffffffffffffffff8111156140b057600080fd5b60008060006060848603121561418057600080fd5b833567ffffffffffffffff81111561419757600080fd5b840161012081870312156141aa57600080fd5b95602085013595506040909401359392505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461074157600080fd5b6000602082840312156141ff57600080fd5b81356108a8816141bf565b6000806000806040858703121561422057600080fd5b843567ffffffffffffffff81111561423757600080fd5b61424387828801613b3e565b909550935050602085013567ffffffffffffffff81111561426357600080fd5b61426f87828801613b3e565b95989497509550505050565b60006020828403121561428d57600080fd5b5035919050565b600080602083850312156142a757600080fd5b823567ffffffffffffffff8111156142be57600080fd5b6142ca85828601613b3e565b90969095509350505050565b600080600080606085870312156142ec57600080fd5b6142f585613b15565b935060208501359250604085013567ffffffffffffffff81111561426357600080fd5b60006020828403121561432a57600080fd5b813567ffffffffffffffff81111561434157600080fd5b61434d84828501613d62565b949350505050565b60008083601f84011261436757600080fd5b50813567ffffffffffffffff81111561437f57600080fd5b6020830191508360208260051b8501011115613b8057600080fd5b600080602083850312156143ad57600080fd5b823567ffffffffffffffff8111156143c457600080fd5b6142ca85828601614355565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000815180845260005b8181101561442557602081850181015186830182015201614409565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015614535577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051600681106144f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261451260608801826143ff565b60409283015197909201969096529450602093840193919091019060010161448b565b50929695505050505050565b6000806040838503121561455457600080fd5b823561455f816141bf565b915061456d60208401613b15565b90509250929050565b60008060008060008060008060a0898b03121561459257600080fd5b61459b89613b15565b97506145a960208a01613b15565b9650604089013567ffffffffffffffff8111156145c557600080fd5b6145d18b828c01614355565b909750955050606089013567ffffffffffffffff8111156145f157600080fd5b6145fd8b828c01614355565b909550935050608089013567ffffffffffffffff811115613c1757600080fd5b60008060008060008060a0878903121561463657600080fd5b61463f87613b15565b955061464d60208801613b15565b94506040870135935060608701359250608087013567ffffffffffffffff81111561467757600080fd5b61468389828a01613b3e565b979a9699509497509295939492505050565b6000806000606084860312156146aa57600080fd5b833592506146ba60208501613b15565b915060408401356bffffffffffffffffffffffff811681146146db57600080fd5b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156106f2576106f26146e6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825180855260208501945060208160051b8301016020850160005b83811015614827577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e060408601526147e360e08601826143ff565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101614775565b50909695505050505050565b600081518084526020840193506020830160005b8281101561487b57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101614847565b5093949350505050565b805160ff168252600060208201516148a1602085018215159052565b50604082015161012060408501526148bd610120850182614757565b9050606083015160608501526080830151608085015260a083015184820360a08601526148ea82826143ff565b91505060c083015160c085015260e083015160e085015261010083015184820361010086015261491a8282614833565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60408152600061497f6040830186614885565b8281036020840152614992818587614923565b9695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126149d157600080fd5b83018035915067ffffffffffffffff8211156149ec57600080fd5b602001915036819003821315613b8057600080fd5b83815260406020820152600061491a604083018486614923565b600060208284031215614a2d57600080fd5b81516108a8816141bf565b6020815260006108a860208301846143ff565b60208152600061434d602083018486614923565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21833603018112614a9357600080fd5b9190910192915050565b60006106f23683613df0565b60ff8716815285602082015284604082015283606082015260ff8316608082015260c060a08201526000614ae060c08301846143ff565b98975050505050505050565b818103818111156106f2576106f26146e6565b60008085851115614b0f57600080fd5b83861115614b1c57600080fd5b5050820193919092039150565b606081526000614b3c6060830186614885565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000614ae060c08301846143ff565b83815282602082015260606040820152600061491a60608301846143ff565b606081526000614bb16060830186614885565b846020840152828103604084015261499281856143ff565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061434d60408301846143ff565b60006040828403128015614c0b57600080fd5b506040805190810167ffffffffffffffff81118282101715614c2f57614c2f613c52565b604052825181526020928301519281019290925250919050565b8151600090829060208501835b82811015614c8a57815173ffffffffffffffffffffffffffffffffffffffff16845260209384019390910190600101614c56565b509195945050505050565b606081526000614ca9606083018688614923565b6020830194909452506040015292915050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000614992606083018486614923565b600060208284031215614d0457600080fd5b5051919050565b8151600090829060208501835b82811015614c8a578151845260209384019390910190600101614d18565b60ff81811683821601908111156106f2576106f26146e656fea2646970667358221220cd4d3ff6f35d824793bb81ca683112170d9fbcc498509bbd60f84905ffb3b13964736f6c634300081c0033",
}

// WalletSimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletSimulatorMetaData.ABI instead.
var WalletSimulatorABI = WalletSimulatorMetaData.ABI

// WalletSimulatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletSimulatorMetaData.Bin instead.
var WalletSimulatorBin = WalletSimulatorMetaData.Bin

// DeployWalletSimulator deploys a new Ethereum contract, binding an instance of WalletSimulator to it.
func DeployWalletSimulator(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *WalletSimulator, error) {
	parsed, err := WalletSimulatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletSimulatorBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletSimulator{WalletSimulatorCaller: WalletSimulatorCaller{contract: contract}, WalletSimulatorTransactor: WalletSimulatorTransactor{contract: contract}, WalletSimulatorFilterer: WalletSimulatorFilterer{contract: contract}}, nil
}

// WalletSimulator is an auto generated Go binding around an Ethereum contract.
type WalletSimulator struct {
	WalletSimulatorCaller     // Read-only binding to the contract
	WalletSimulatorTransactor // Write-only binding to the contract
	WalletSimulatorFilterer   // Log filterer for contract events
}

// WalletSimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletSimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletSimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletSimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSimulatorSession struct {
	Contract     *WalletSimulator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletSimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletSimulatorCallerSession struct {
	Contract *WalletSimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WalletSimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletSimulatorTransactorSession struct {
	Contract     *WalletSimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WalletSimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletSimulatorRaw struct {
	Contract *WalletSimulator // Generic contract binding to access the raw methods on
}

// WalletSimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletSimulatorCallerRaw struct {
	Contract *WalletSimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// WalletSimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletSimulatorTransactorRaw struct {
	Contract *WalletSimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletSimulator creates a new instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulator(address common.Address, backend bind.ContractBackend) (*WalletSimulator, error) {
	contract, err := bindWalletSimulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletSimulator{WalletSimulatorCaller: WalletSimulatorCaller{contract: contract}, WalletSimulatorTransactor: WalletSimulatorTransactor{contract: contract}, WalletSimulatorFilterer: WalletSimulatorFilterer{contract: contract}}, nil
}

// NewWalletSimulatorCaller creates a new read-only instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorCaller(address common.Address, caller bind.ContractCaller) (*WalletSimulatorCaller, error) {
	contract, err := bindWalletSimulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCaller{contract: contract}, nil
}

// NewWalletSimulatorTransactor creates a new write-only instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletSimulatorTransactor, error) {
	contract, err := bindWalletSimulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorTransactor{contract: contract}, nil
}

// NewWalletSimulatorFilterer creates a new log filterer instance of WalletSimulator, bound to a specific deployed contract.
func NewWalletSimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletSimulatorFilterer, error) {
	contract, err := bindWalletSimulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorFilterer{contract: contract}, nil
}

// bindWalletSimulator binds a generic wrapper to an already deployed contract.
func bindWalletSimulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletSimulatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletSimulator *WalletSimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletSimulator.Contract.WalletSimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletSimulator *WalletSimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletSimulator.Contract.WalletSimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletSimulator *WalletSimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletSimulator.Contract.WalletSimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletSimulator *WalletSimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletSimulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletSimulator *WalletSimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletSimulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletSimulator *WalletSimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletSimulator.Contract.contract.Transact(opts, method, params...)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletSimulator *WalletSimulatorCaller) Entrypoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "entrypoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletSimulator *WalletSimulatorSession) Entrypoint() (common.Address, error) {
	return _WalletSimulator.Contract.Entrypoint(&_WalletSimulator.CallOpts)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletSimulator *WalletSimulatorCallerSession) Entrypoint() (common.Address, error) {
	return _WalletSimulator.Contract.Entrypoint(&_WalletSimulator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletSimulator *WalletSimulatorCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletSimulator *WalletSimulatorSession) GetImplementation() (common.Address, error) {
	return _WalletSimulator.Contract.GetImplementation(&_WalletSimulator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletSimulator *WalletSimulatorCallerSession) GetImplementation() (common.Address, error) {
	return _WalletSimulator.Contract.GetImplementation(&_WalletSimulator.CallOpts)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletSimulator *WalletSimulatorCaller) GetStaticSignature(opts *bind.CallOpts, _hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "getStaticSignature", _hash)

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
func (_WalletSimulator *WalletSimulatorSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletSimulator.Contract.GetStaticSignature(&_WalletSimulator.CallOpts, _hash)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletSimulator *WalletSimulatorCallerSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletSimulator.Contract.GetStaticSignature(&_WalletSimulator.CallOpts, _hash)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletSimulator *WalletSimulatorCaller) ImageHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletSimulator *WalletSimulatorSession) ImageHash() ([32]byte, error) {
	return _WalletSimulator.Contract.ImageHash(&_WalletSimulator.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletSimulator *WalletSimulatorCallerSession) ImageHash() ([32]byte, error) {
	return _WalletSimulator.Contract.ImageHash(&_WalletSimulator.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletSimulator *WalletSimulatorCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletSimulator *WalletSimulatorSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.IsValidSignature(&_WalletSimulator.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletSimulator *WalletSimulatorCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.IsValidSignature(&_WalletSimulator.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC1155BatchReceived(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC1155BatchReceived(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC1155Received(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC1155Received(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC721Received(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.OnERC721Received(&_WalletSimulator.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletSimulator *WalletSimulatorCaller) ReadHook(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "readHook", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletSimulator *WalletSimulatorSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletSimulator.Contract.ReadHook(&_WalletSimulator.CallOpts, selector)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletSimulator *WalletSimulatorCallerSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletSimulator.Contract.ReadHook(&_WalletSimulator.CallOpts, selector)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletSimulator *WalletSimulatorCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletSimulator *WalletSimulatorSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletSimulator.Contract.ReadNonce(&_WalletSimulator.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletSimulator *WalletSimulatorCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletSimulator.Contract.ReadNonce(&_WalletSimulator.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletSimulator *WalletSimulatorCaller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletSimulator *WalletSimulatorSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletSimulator.Contract.RecoverPartialSignature(&_WalletSimulator.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletSimulator *WalletSimulatorCallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletSimulator.Contract.RecoverPartialSignature(&_WalletSimulator.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletSimulator *WalletSimulatorCaller) RecoverSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "recoverSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletSimulator *WalletSimulatorSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletSimulator.Contract.RecoverSapientSignature(&_WalletSimulator.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletSimulator *WalletSimulatorCallerSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletSimulator.Contract.RecoverSapientSignature(&_WalletSimulator.CallOpts, _payload, _signature)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCaller) TokenReceived(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletSimulator.contract.Call(opts, &out, "tokenReceived", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.TokenReceived(&_WalletSimulator.CallOpts, arg0, arg1, arg2)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletSimulator *WalletSimulatorCallerSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletSimulator.Contract.TokenReceived(&_WalletSimulator.CallOpts, arg0, arg1, arg2)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) AddHook(opts *bind.TransactOpts, selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "addHook", selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletSimulator *WalletSimulatorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.Contract.AddHook(&_WalletSimulator.TransactOpts, selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.Contract.AddHook(&_WalletSimulator.TransactOpts, selector, implementation)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletSimulator *WalletSimulatorTransactor) CreateContract(opts *bind.TransactOpts, _code []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "createContract", _code)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletSimulator *WalletSimulatorSession) CreateContract(_code []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.CreateContract(&_WalletSimulator.TransactOpts, _code)
}

// CreateContract is a paid mutator transaction binding the contract method 0x90042baf.
//
// Solidity: function createContract(bytes _code) payable returns(address addr)
func (_WalletSimulator *WalletSimulatorTransactorSession) CreateContract(_code []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.CreateContract(&_WalletSimulator.TransactOpts, _code)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletSimulator *WalletSimulatorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Execute(&_WalletSimulator.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Execute(&_WalletSimulator.TransactOpts, _payload, _signature)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletSimulator *WalletSimulatorTransactor) ExecuteUserOp(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "executeUserOp", _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletSimulator *WalletSimulatorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.ExecuteUserOp(&_WalletSimulator.TransactOpts, _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.ExecuteUserOp(&_WalletSimulator.TransactOpts, _payload)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) RemoveHook(opts *bind.TransactOpts, selector [4]byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "removeHook", selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletSimulator *WalletSimulatorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.RemoveHook(&_WalletSimulator.TransactOpts, selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.RemoveHook(&_WalletSimulator.TransactOpts, selector)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletSimulator *WalletSimulatorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.SelfExecute(&_WalletSimulator.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.SelfExecute(&_WalletSimulator.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletSimulator *WalletSimulatorTransactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletSimulator *WalletSimulatorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.Contract.SetStaticSignature(&_WalletSimulator.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.Contract.SetStaticSignature(&_WalletSimulator.TransactOpts, _hash, _address, _timestamp)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorTransactor) Simulate(opts *bind.TransactOpts, _calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "simulate", _calls)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorSession) Simulate(_calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Simulate(&_WalletSimulator.TransactOpts, _calls)
}

// Simulate is a paid mutator transaction binding the contract method 0xa27c3922.
//
// Solidity: function simulate((address,uint256,bytes,uint256,bool,bool,uint256)[] _calls) returns((uint8,bytes,uint256)[] results)
func (_WalletSimulator *WalletSimulatorTransactorSession) Simulate(_calls []PayloadCall) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Simulate(&_WalletSimulator.TransactOpts, _calls)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletSimulator *WalletSimulatorTransactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "tokensReceived", operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletSimulator *WalletSimulatorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.TokensReceived(&_WalletSimulator.TransactOpts, operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.TokensReceived(&_WalletSimulator.TransactOpts, operator, from, to, amount, data, operatorData)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletSimulator *WalletSimulatorTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletSimulator *WalletSimulatorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.UpdateImageHash(&_WalletSimulator.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.UpdateImageHash(&_WalletSimulator.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletSimulator *WalletSimulatorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.Contract.UpdateImplementation(&_WalletSimulator.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletSimulator.Contract.UpdateImplementation(&_WalletSimulator.TransactOpts, _implementation)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletSimulator *WalletSimulatorTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletSimulator *WalletSimulatorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.Contract.ValidateUserOp(&_WalletSimulator.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletSimulator *WalletSimulatorTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletSimulator.Contract.ValidateUserOp(&_WalletSimulator.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletSimulator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletSimulator *WalletSimulatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Fallback(&_WalletSimulator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletSimulator.Contract.Fallback(&_WalletSimulator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletSimulator *WalletSimulatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletSimulator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletSimulator *WalletSimulatorSession) Receive() (*types.Transaction, error) {
	return _WalletSimulator.Contract.Receive(&_WalletSimulator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletSimulator *WalletSimulatorTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletSimulator.Contract.Receive(&_WalletSimulator.TransactOpts)
}

// WalletSimulatorCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletSimulator contract.
type WalletSimulatorCallAbortedIterator struct {
	Event *WalletSimulatorCallAborted // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorCallAborted)
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
		it.Event = new(WalletSimulatorCallAborted)
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
func (it *WalletSimulatorCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorCallAborted represents a CallAborted event raised by the WalletSimulator contract.
type WalletSimulatorCallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletSimulator *WalletSimulatorFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletSimulatorCallAbortedIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCallAbortedIterator{contract: _WalletSimulator.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletSimulator *WalletSimulatorFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletSimulatorCallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorCallAborted)
				if err := _WalletSimulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseCallAborted(log types.Log) (*WalletSimulatorCallAborted, error) {
	event := new(WalletSimulatorCallAborted)
	if err := _WalletSimulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletSimulator contract.
type WalletSimulatorCallFailedIterator struct {
	Event *WalletSimulatorCallFailed // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorCallFailed)
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
		it.Event = new(WalletSimulatorCallFailed)
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
func (it *WalletSimulatorCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorCallFailed represents a CallFailed event raised by the WalletSimulator contract.
type WalletSimulatorCallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletSimulator *WalletSimulatorFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletSimulatorCallFailedIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCallFailedIterator{contract: _WalletSimulator.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletSimulator *WalletSimulatorFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletSimulatorCallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorCallFailed)
				if err := _WalletSimulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseCallFailed(log types.Log) (*WalletSimulatorCallFailed, error) {
	event := new(WalletSimulatorCallFailed)
	if err := _WalletSimulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletSimulator contract.
type WalletSimulatorCallSkippedIterator struct {
	Event *WalletSimulatorCallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorCallSkipped)
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
		it.Event = new(WalletSimulatorCallSkipped)
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
func (it *WalletSimulatorCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorCallSkipped represents a CallSkipped event raised by the WalletSimulator contract.
type WalletSimulatorCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletSimulator *WalletSimulatorFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletSimulatorCallSkippedIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCallSkippedIterator{contract: _WalletSimulator.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletSimulator *WalletSimulatorFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletSimulatorCallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorCallSkipped)
				if err := _WalletSimulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseCallSkipped(log types.Log) (*WalletSimulatorCallSkipped, error) {
	event := new(WalletSimulatorCallSkipped)
	if err := _WalletSimulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletSimulator contract.
type WalletSimulatorCallSucceededIterator struct {
	Event *WalletSimulatorCallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorCallSucceeded)
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
		it.Event = new(WalletSimulatorCallSucceeded)
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
func (it *WalletSimulatorCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorCallSucceeded represents a CallSucceeded event raised by the WalletSimulator contract.
type WalletSimulatorCallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletSimulator *WalletSimulatorFilterer) FilterCallSucceeded(opts *bind.FilterOpts) (*WalletSimulatorCallSucceededIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCallSucceededIterator{contract: _WalletSimulator.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletSimulator *WalletSimulatorFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletSimulatorCallSucceeded) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorCallSucceeded)
				if err := _WalletSimulator.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseCallSucceeded(log types.Log) (*WalletSimulatorCallSucceeded, error) {
	event := new(WalletSimulatorCallSucceeded)
	if err := _WalletSimulator.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorCreatedContractIterator is returned from FilterCreatedContract and is used to iterate over the raw logs and unpacked data for CreatedContract events raised by the WalletSimulator contract.
type WalletSimulatorCreatedContractIterator struct {
	Event *WalletSimulatorCreatedContract // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorCreatedContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorCreatedContract)
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
		it.Event = new(WalletSimulatorCreatedContract)
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
func (it *WalletSimulatorCreatedContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorCreatedContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorCreatedContract represents a CreatedContract event raised by the WalletSimulator contract.
type WalletSimulatorCreatedContract struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatedContract is a free log retrieval operation binding the contract event 0xa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c.
//
// Solidity: event CreatedContract(address _contract)
func (_WalletSimulator *WalletSimulatorFilterer) FilterCreatedContract(opts *bind.FilterOpts) (*WalletSimulatorCreatedContractIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "CreatedContract")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorCreatedContractIterator{contract: _WalletSimulator.contract, event: "CreatedContract", logs: logs, sub: sub}, nil
}

// WatchCreatedContract is a free log subscription operation binding the contract event 0xa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c.
//
// Solidity: event CreatedContract(address _contract)
func (_WalletSimulator *WalletSimulatorFilterer) WatchCreatedContract(opts *bind.WatchOpts, sink chan<- *WalletSimulatorCreatedContract) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "CreatedContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorCreatedContract)
				if err := _WalletSimulator.contract.UnpackLog(event, "CreatedContract", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseCreatedContract(log types.Log) (*WalletSimulatorCreatedContract, error) {
	event := new(WalletSimulatorCreatedContract)
	if err := _WalletSimulator.contract.UnpackLog(event, "CreatedContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorDefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletSimulator contract.
type WalletSimulatorDefinedHookIterator struct {
	Event *WalletSimulatorDefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorDefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorDefinedHook)
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
		it.Event = new(WalletSimulatorDefinedHook)
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
func (it *WalletSimulatorDefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorDefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorDefinedHook represents a DefinedHook event raised by the WalletSimulator contract.
type WalletSimulatorDefinedHook struct {
	Selector       [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletSimulator *WalletSimulatorFilterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletSimulatorDefinedHookIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorDefinedHookIterator{contract: _WalletSimulator.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletSimulator *WalletSimulatorFilterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletSimulatorDefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorDefinedHook)
				if err := _WalletSimulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletSimulator *WalletSimulatorFilterer) ParseDefinedHook(log types.Log) (*WalletSimulatorDefinedHook, error) {
	event := new(WalletSimulatorDefinedHook)
	if err := _WalletSimulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletSimulator contract.
type WalletSimulatorImageHashUpdatedIterator struct {
	Event *WalletSimulatorImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorImageHashUpdated)
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
		it.Event = new(WalletSimulatorImageHashUpdated)
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
func (it *WalletSimulatorImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorImageHashUpdated represents a ImageHashUpdated event raised by the WalletSimulator contract.
type WalletSimulatorImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletSimulator *WalletSimulatorFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletSimulatorImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorImageHashUpdatedIterator{contract: _WalletSimulator.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletSimulator *WalletSimulatorFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletSimulatorImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorImageHashUpdated)
				if err := _WalletSimulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseImageHashUpdated(log types.Log) (*WalletSimulatorImageHashUpdated, error) {
	event := new(WalletSimulatorImageHashUpdated)
	if err := _WalletSimulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletSimulator contract.
type WalletSimulatorImplementationUpdatedIterator struct {
	Event *WalletSimulatorImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorImplementationUpdated)
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
		it.Event = new(WalletSimulatorImplementationUpdated)
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
func (it *WalletSimulatorImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorImplementationUpdated represents a ImplementationUpdated event raised by the WalletSimulator contract.
type WalletSimulatorImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletSimulator *WalletSimulatorFilterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletSimulatorImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorImplementationUpdatedIterator{contract: _WalletSimulator.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletSimulator *WalletSimulatorFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletSimulatorImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorImplementationUpdated)
				if err := _WalletSimulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseImplementationUpdated(log types.Log) (*WalletSimulatorImplementationUpdated, error) {
	event := new(WalletSimulatorImplementationUpdated)
	if err := _WalletSimulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletSimulator contract.
type WalletSimulatorNonceChangeIterator struct {
	Event *WalletSimulatorNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorNonceChange)
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
		it.Event = new(WalletSimulatorNonceChange)
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
func (it *WalletSimulatorNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorNonceChange represents a NonceChange event raised by the WalletSimulator contract.
type WalletSimulatorNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletSimulator *WalletSimulatorFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletSimulatorNonceChangeIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorNonceChangeIterator{contract: _WalletSimulator.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletSimulator *WalletSimulatorFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletSimulatorNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorNonceChange)
				if err := _WalletSimulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseNonceChange(log types.Log) (*WalletSimulatorNonceChange, error) {
	event := new(WalletSimulatorNonceChange)
	if err := _WalletSimulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorStaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletSimulator contract.
type WalletSimulatorStaticSignatureSetIterator struct {
	Event *WalletSimulatorStaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletSimulatorStaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletSimulatorStaticSignatureSet)
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
		it.Event = new(WalletSimulatorStaticSignatureSet)
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
func (it *WalletSimulatorStaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletSimulatorStaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletSimulatorStaticSignatureSet represents a StaticSignatureSet event raised by the WalletSimulator contract.
type WalletSimulatorStaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletSimulator *WalletSimulatorFilterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletSimulatorStaticSignatureSetIterator, error) {

	logs, sub, err := _WalletSimulator.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletSimulatorStaticSignatureSetIterator{contract: _WalletSimulator.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletSimulator *WalletSimulatorFilterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletSimulatorStaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletSimulator.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletSimulatorStaticSignatureSet)
				if err := _WalletSimulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletSimulator *WalletSimulatorFilterer) ParseStaticSignatureSet(log types.Log) (*WalletSimulatorStaticSignatureSet, error) {
	event := new(WalletSimulatorStaticSignatureSet)
	if err := _WalletSimulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletSimulatorDeployedBin is the resulting bytecode of the created contract
const WalletSimulatorDeployedBin = "0x6080604052600436106101995760003560e01c80638943ec02116100e1578063a65d69d41161008a578063b93ea7ad11610064578063b93ea7ad146105e3578063bc197c81146105f6578063f23a6e611461063e578063f727ef1c14610684576101a0565b8063a65d69d414610547578063aaf10f421461057b578063ad55366b14610590576101a0565b806392dcb3fc116100bb57806392dcb3fc146104ae5780639c145aed146104fa578063a27c39221461051a576101a0565b80638943ec02146104375780638c3f55631461047b57806390042baf1461049b576101a0565b80631a9b2337116101435780634fcf3eca1161011d5780634fcf3eca146103fc57806351605d801461040f5780636ea4457714610424576101a0565b80631a9b2337146103845780631f6a1eb9146103c957806329561426146103dc576101a0565b8063150b7a0211610174578063150b7a02146102ce5780631626ba7e1461034457806319822f7c14610364576101a0565b806223de2914610263578063025b22bc1461028857806313792a4a1461029b576101a0565b366101a057005b600436106102615760006101bc6101b73683613a9f565b6106a4565b905073ffffffffffffffffffffffffffffffffffffffff81161561025f576000808273ffffffffffffffffffffffffffffffffffffffff16600036604051610205929190613b05565b600060405180830381855af49150503d8060008114610240576040519150601f19603f3d011682016040523d82523d6000602084013e610245565b606091505b50915091508161025757805160208201fd5b805160208201f35b505b005b34801561026f57600080fd5b5061026161027e366004613b87565b5050505050505050565b610261610296366004613c37565b6106f8565b3480156102a757600080fd5b506102bb6102b6366004613f82565b610744565b6040519081526020015b60405180910390f35b3480156102da57600080fd5b506103136102e93660046140c9565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016102c5565b34801561035057600080fd5b5061031361035f366004614138565b6108af565b34801561037057600080fd5b506102bb61037f36600461416b565b610946565b34801561039057600080fd5b506103a461039f3660046141ed565b610b9c565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102c5565b6102616103d736600461420a565b610ba7565b3480156103e857600080fd5b506102616103f736600461427b565b610c26565b61026161040a3660046141ed565b610c6a565b34801561041b57600080fd5b506102bb610d2c565b610261610432366004614294565b610d5b565b34801561044357600080fd5b506103136104523660046142d6565b7f8943ec0200000000000000000000000000000000000000000000000000000000949350505050565b34801561048757600080fd5b506102bb61049636600461427b565b610dc8565b6103a46104a9366004614318565b610df4565b3480156104ba57600080fd5b506104ce6104c936600461427b565b610ed8565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016102c5565b34801561050657600080fd5b50610261610515366004614294565b610eed565b34801561052657600080fd5b5061053a61053536600461439a565b61103f565b6040516102c59190614463565b34801561055357600080fd5b506103a47f000000000000000000000000000000000000000000000000000000000000000081565b34801561058757600080fd5b506103a461157c565b34801561059c57600080fd5b506105b06105ab366004613f82565b611586565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c0016102c5565b6102616105f1366004614541565b6115c0565b34801561060257600080fd5b50610313610611366004614576565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561064a57600080fd5b5061031361065936600461461d565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561069057600080fd5b5061026161069f366004614695565b611685565b60006106f27fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416611740565b92915050565b333014610738576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6107418161179e565b50565b60008084610100015151600161075a9190614715565b67ffffffffffffffff81111561077257610772613c52565b60405190808252806020026020018201604052801561079b578160200160208202803683370190505b50905060005b8561010001515181101561080d5785610100015181815181106107c6576107c6614728565b60200260200101518282815181106107e0576107e0614728565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016107a1565b503381866101000151518151811061082757610827614728565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006108618686866117f3565b509050806108a1578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161072f9392919061496c565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905260006109078286866117f3565b5090508061091b5750600091506108a89050565b507f1626ba7e0000000000000000000000000000000000000000000000000000000095945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166109b5576040517fd13d783500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a26576040517f1d6ddbf400000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b8115610ace576040517fb760faf90000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b760faf99084906024016000604051808303818588803b158015610ab457600080fd5b505af1158015610ac8573d6000803e3d6000fd5b50505050505b7f1626ba7e0000000000000000000000000000000000000000000000000000000030631626ba7e85610b0461010089018961499c565b6040518463ffffffff1660e01b8152600401610b2293929190614a01565b602060405180830381865afa158015610b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b639190614a1b565b7fffffffff000000000000000000000000000000000000000000000000000000001614610b92575060016108a8565b5060009392505050565b60006106f2826106a4565b60005a90506000610bb886866119da565b9050610bcc81606001518260800151611df8565b600080610bda8387876117f3565b9150915081610c1b578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161072f9392919061496c565b61027e848285611ee0565b333014610c61576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b610741816121b5565b333014610ca5576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6000610cb0826106a4565b73ffffffffffffffffffffffffffffffffffffffff1603610d21576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008216600482015260240161072f565b610741816000612245565b6000610d567fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610d96576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b60005a90506000610da784846119da565b90506000610db482612305565b9050610dc1838284611ee0565b5050505050565b60006106f27f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83611740565b6000333014610e31576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b81516020830134f0905073ffffffffffffffffffffffffffffffffffffffff8116610e8a57816040517f0d25719100000000000000000000000000000000000000000000000000000000815260040161072f9190614a38565b60405173ffffffffffffffffffffffffffffffffffffffff821681527fa506ad4e7f05eceba62a023c3219e5bd98a615f4fa87e2afb08a2da5cf62bf0c9060200160405180910390a1919050565b600080610ee483612380565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f5a576040517fd13d783500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610fcb576040517f1d6ddbf400000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6040517f6ea445770000000000000000000000000000000000000000000000000000000081523090636ea44577906110099085908590600401614a4b565b600060405180830381600087803b15801561102357600080fd5b505af1158015611037573d6000803e3d6000fd5b505050505050565b606060005a90506000838067ffffffffffffffff81111561106257611062613c52565b6040519080825280602002602001820160405280156110b857816020015b6110a56040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816110805790505b50935060005b818110156115725760008787838181106110da576110da614728565b90506020028101906110ec9190614a5f565b6110f590614a9d565b90508060a001518015611106575083155b15611111575061156a565b60608101516000945080158015906111285750805a105b156111c257600587848151811061114157611141614728565b602002602001015160000190600581111561115e5761115e6143d0565b90816005811115611171576111716143d0565b9052505a60405160200161118791815260200190565b6040516020818303038152906040528784815181106111a8576111a8614728565b6020026020010151602001819052505050505050506106f2565b60008260800151156112c85760005a84519091506112939084156111e657846111e8565b5a5b634c4e814c60e01b60008c8a8c60008c6040015160405160240161121196959493929190614aa9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526123cc565b91505a6112a09082614aec565b8986815181106112b2576112b2614728565b6020026020010151604001818152505050611322565b60005a845160208601519192506112f19185156112e557856112e7565b5a5b87604001516123e2565b91505a6112fe9082614aec565b89868151811061131057611310614728565b60200260200101516040018181525050505b806114f65760c08301516113ac5760019550600288858151811061134857611348614728565b6020026020010151600001906005811115611365576113656143d0565b90816005811115611378576113786143d0565b9052506113836123fa565b88858151811061139557611395614728565b60200260200101516020018190525050505061156a565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114535760048885815181106113eb576113eb614728565b6020026020010151600001906005811115611408576114086143d0565b9081600581111561141b5761141b6143d0565b9052506114266123fa565b88858151811061143857611438614728565b602002602001015160200181905250505050505050506106f2565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016114f657600388858151811061149257611492614728565b60200260200101516000019060058111156114af576114af6143d0565b908160058111156114c2576114c26143d0565b9052506114cd6123fa565b8885815181106114df576114df614728565b602002602001015160200181905250505050611572565b600188858151811061150a5761150a614728565b6020026020010151600001906005811115611527576115276143d0565b9081600581111561153a5761153a6143d0565b9052506115456123fa565b88858151811061155757611557614728565b6020026020010151602001819052505050505b6001016110be565b5050505092915050565b6000610d56305490565b60008060008060008061159d898989600080612419565b9399509197509450925090506115b28361273e565b935093975093979195509350565b3330146115fb576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6000611606836106a4565b73ffffffffffffffffffffffffffffffffffffffff1614611677576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161072f565b6116818282612245565b5050565b3330146116c0576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161072f565b6116d98383836bffffffffffffffffffffffff16612749565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b600080838360405160200161175f929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b6117a6813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b60008060008484600081811061180b5761180b614728565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f800000000000000000000000000000000000000000000000000000000000000090811690036119605761186886612305565b915060008061187684612380565b915091504281116118bd576040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161072f565b73ffffffffffffffffffffffffffffffffffffffff8216158015906118f8575073ffffffffffffffffffffffffffffffffffffffff82163314155b15611954576040517f8945c3130000000000000000000000000000000000000000000000000000000081526004810185905233602482015273ffffffffffffffffffffffffffffffffffffffff8316604482015260640161072f565b600194505050506119d2565b6000806000611973898989600080612419565b9850929550909350915050828210156119c2576040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161072f565b6119cb8161273e565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103611a3e5760006060840152611a4f565b84810135606090811c908401526014015b6007600183901c168015611aa25760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003611ab757506001611ae0565b83602016602003611ad35750600282019186013560f01c611ae0565b50600182019186013560f81c5b8067ffffffffffffffff811115611af957611af9613c52565b604051908082528060200260200182016040528015611b8457816020015b611b716040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081611b175790505b50604086015260005b81811015611ded5760018085019489013560f81c908082169003611bec573087604001518381518110611bc257611bc2614728565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611c36565b8489013560601c6014860188604001518481518110611c0d57611c0d614728565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611c7457848901356020860188604001518481518110611c5f57611c5f614728565b60200260200101516020018197508281525050505b60048082169003611d0c57600385019489013560e81c89868a611c978483614715565b92611ca493929190614aff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506040890151805185908110611cef57611cef614728565b602090810291909101015160400152611d088187614715565b9550505b60088082169003611d4a57848901356020860188604001518481518110611d3557611d35614728565b60200260200101516060018197508281525050505b8060101660ff1660101487604001518381518110611d6a57611d6a614728565b602002602001015160800190151590811515815250508060201660ff1660201487604001518381518110611da057611da0614728565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110611dd557611dd5614728565b602090810291909101015160c0015250600101611b8d565b505050505092915050565b6000611e0383610dc8565b9050818114611e4f576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161072f565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561103757600084604001518281518110611f0b57611f0b614728565b602002602001015190508060a001518015611f24575083155b15611f685760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506121ad565b6060810151600094508015801590611f7f5750805a105b15611fbc5785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161072f93929190614b29565b600082608001511561200e578251612007908315611fda5783611fdc565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161121196959493929190614b4e565b9050612035565b82516020840151612032919084156120265784612028565b5a5b86604001516123e2565b90505b806121705760c083015161209157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856120726123fa565b60405161208193929190614b7f565b60405180910390a15050506121ad565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016120fb5786846120c66123fa565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161072f93929190614b9e565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612170577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856121516123fa565b60405161216093929190614b7f565b60405180910390a1505050611037565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b600101611eeb565b806121ec576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6122157fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa906020016117e8565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806123168360200151306127d8565b90506000612323846128a5565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b600080806123ae7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685611740565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015612464575073ffffffffffffffffffffffffffffffffffffffff8916155b15612580578b82013560601c9850601490910190896125805760038201918c013560e81c60008d848e6124978583614715565b926124a493929190614aff565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc8915061252f9030908590600401614bc9565b6040805180830381865afa15801561254b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061256f9190614bf8565b925061257b8285614715565b935050505b826001166001036125ba576125a88d8a838f8f879080926125a393929190614aff565b612b11565b97509750975097509750505050612731565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c168382019096509250600090506126256001600586901c811690614715565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c16995090920191506126708d612305565b935061268e8d858e8e8690809261268993929190614aff565b612d55565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d16909152902082519198509650158015906126d9575080518614155b80156126e9575080602001518511155b1561272d576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161072f565b5050505b9550955095509550959050565b60006106f28261367c565b6127d37fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85612848574661284b565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000808261010001516040516020016128be9190614c49565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120835190915060ff1661296757600061290f84604001516136af565b606085810151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001612361565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016129f75760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01612a675760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016129d9565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01612ad75760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082015290810191909152606081018290526080016129d9565b82516040517f0481832000000000000000000000000000000000000000000000000000000000815260ff909116600482015260240161072f565b6000806000806000612b73604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b88821015612cfd5760038201916000908b013560e81c612bbb8482614715565b9150600090508a8214612bcf576000612bd1565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303612c2957612c188f8d8d87908692612c1093929190614aff565b600185612419565b939d50919b50995097509550612c4b565b612c3f858d8d87908692612c1093929190614aff565b50929c50909a50985096505b89891015612c9757612c5f82858d8f614aff565b8b8b6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161072f9493929190614c95565b819350878d6000015103612caa5760008d525b828710612ced576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018890526024810184905260440161072f565b50505060c0820185905283612b9b565b8a5115801590612d1157508a602001518511155b1561272d576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161072f565b60008060005b8381101561367257600181019085013560f881901c9060fc1c80612e6057600f82166000819003612d935750600183019287013560f81c5b60008080612da28b8b8961375d565b809a5081945082955083965050505050600060018d83868660405160008152602001604052604051612df0949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015612e12573d6000803e3d6000fd5b5050506020604051035190508460ff168a0199506000612e35828760ff166137b0565b905089612e425780612e51565b60008a81526020829052604090205b99505050505050505050612d5b565b60018103612ec457600f82166000819003612e825750600183019287013560f81c5b601484019388013560601c6000612e9c8260ff85166137b0565b905086612ea95780612eb8565b60008781526020829052604090205b96505050505050612d5b565b6002810361308e57600382166000819003612ee65750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c16828801809850819250505060008188019050631626ba7e60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d908792612fa493929190614aff565b6040518463ffffffff1660e01b8152600401612fc293929190614a01565b602060405180830381865afa158015612fdf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130039190614a1b565b7fffffffff000000000000000000000000000000000000000000000000000000001614613074578c848d8d8b90859261303e93929190614aff565b6040517fb2fed7ae00000000000000000000000000000000000000000000000000000000815260040161072f9493929190614cbc565b8097508460ff168a0199506000612e35858760ff166137b0565b600381036130c2576020830192870135846130a957806130b8565b60008581526020829052604090205b9450505050612d5b565b6004810361316557600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806131378e8e8e8e8c90889261268993929190614aff565b91509150829750818a019950613157898260009182526020526040902090565b985050505050505050612d5b565b6006810361322f576003600283901c16600081900361318b5750600183019287013560f81c5b6003831660008190036131a55750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806131e38f8f8f8f8d90889261268993929190614aff565b915091508298508482106131f657998501995b6000613203828789613817565b90508a613210578061321f565b60008b81526020829052604090205b9a50505050505050505050612d5b565b6005810361329c57602083019287013588810361326a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b60006132758261387b565b9050856132825780613291565b60008681526020829052604090205b955050505050612d5b565b6007810361337a57600f821660008190036132be5750600183019287013560f81c5b600080806132cd8b8b8961375d565b604051909a50929550909350915060009060019061331b908f906020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018690526080810185905260a001612df0565b600881036133ce57602083019287013560006133968b826138cf565b90508082036133c3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b6000612e9c8361394a565b60098103613506576003821660008190036133f05750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c90879261348a93929190614aff565b6040518463ffffffff1660e01b81526004016134a89392919061496c565b602060405180830381865afa1580156134c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134e99190614cf2565b90508197508460ff168a0199506000612e35858760ff1684613985565b600a810361363d576003821660008190036135285750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d9087926135c193929190614aff565b6040518463ffffffff1660e01b81526004016135df93929190614a01565b602060405180830381865afa1580156135fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136209190614cf2565b90508198508560ff168b019a506000613203868860ff1684613985565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161072f565b5094509492505050565b600081158015906106f25750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b600080825167ffffffffffffffff8111156136cc576136cc613c52565b6040519080825280602002602001820160405280156136f5578160200160208202803683370190505b50905060005b835181101561374b5761372684828151811061371957613719614728565b60200260200101516139f3565b82828151811061373857613738614728565b60209081029190910101526001016136fb565b50806040516020016129d99190614d0b565b828101803590602001357f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811690600090604085019060ff81901c6137a381601b614d36565b9350505093509350935093565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612887565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806138e08460200151846127d8565b905060006138ed856128a5565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a60208201529081018290526000906060016138b2565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d0161385c565b80516020808301516040808501518051908401206060860151608087015160a088015160c089015194516000986138b2987f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef4379891979196959493920197885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015613afe577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613b3957600080fd5b919050565b60008083601f840112613b5057600080fd5b50813567ffffffffffffffff811115613b6857600080fd5b602083019150836020828501011115613b8057600080fd5b9250929050565b60008060008060008060008060c0898b031215613ba357600080fd5b613bac89613b15565b9750613bba60208a01613b15565b9650613bc860408a01613b15565b955060608901359450608089013567ffffffffffffffff811115613beb57600080fd5b613bf78b828c01613b3e565b90955093505060a089013567ffffffffffffffff811115613c1757600080fd5b613c238b828c01613b3e565b999c989b5096995094979396929594505050565b600060208284031215613c4957600080fd5b6108a882613b15565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613ca457613ca4613c52565b60405290565b604051610120810167ffffffffffffffff81118282101715613ca457613ca4613c52565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613d1557613d15613c52565b604052919050565b803560ff81168114613b3957600080fd5b80358015158114613b3957600080fd5b600067ffffffffffffffff821115613d5857613d58613c52565b5060051b60200190565b600082601f830112613d7357600080fd5b813567ffffffffffffffff811115613d8d57613d8d613c52565b613dbe60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613cce565b818152846020838601011115613dd357600080fd5b816020850160208301376000918101602001919091529392505050565b600060e08284031215613e0257600080fd5b613e0a613c81565b9050613e1582613b15565b815260208281013590820152604082013567ffffffffffffffff811115613e3b57600080fd5b613e4784828501613d62565b60408301525060608281013590820152613e6360808301613d2e565b6080820152613e7460a08301613d2e565b60a082015260c09182013591810191909152919050565b600082601f830112613e9c57600080fd5b8135613eaf613eaa82613d3e565b613cce565b8082825260208201915060208360051b860101925085831115613ed157600080fd5b602085015b83811015613f1357803567ffffffffffffffff811115613ef557600080fd5b613f04886020838a0101613df0565b84525060209283019201613ed6565b5095945050505050565b600082601f830112613f2e57600080fd5b8135613f3c613eaa82613d3e565b8082825260208201915060208360051b860101925085831115613f5e57600080fd5b602085015b83811015613f1357613f7481613b15565b835260209283019201613f63565b600080600060408486031215613f9757600080fd5b833567ffffffffffffffff811115613fae57600080fd5b84016101208187031215613fc157600080fd5b613fc9613caa565b613fd282613d1d565b8152613fe060208301613d2e565b6020820152604082013567ffffffffffffffff811115613fff57600080fd5b61400b88828501613e8b565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561403f57600080fd5b61404b88828501613d62565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561408057600080fd5b61408c88828501613f1d565b61010083015250935050602084013567ffffffffffffffff8111156140b057600080fd5b6140bc86828701613b3e565b9497909650939450505050565b6000806000806000608086880312156140e157600080fd5b6140ea86613b15565b94506140f860208701613b15565b935060408601359250606086013567ffffffffffffffff81111561411b57600080fd5b61412788828901613b3e565b969995985093965092949392505050565b60008060006040848603121561414d57600080fd5b83359250602084013567ffffffffffffffff8111156140b057600080fd5b60008060006060848603121561418057600080fd5b833567ffffffffffffffff81111561419757600080fd5b840161012081870312156141aa57600080fd5b95602085013595506040909401359392505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461074157600080fd5b6000602082840312156141ff57600080fd5b81356108a8816141bf565b6000806000806040858703121561422057600080fd5b843567ffffffffffffffff81111561423757600080fd5b61424387828801613b3e565b909550935050602085013567ffffffffffffffff81111561426357600080fd5b61426f87828801613b3e565b95989497509550505050565b60006020828403121561428d57600080fd5b5035919050565b600080602083850312156142a757600080fd5b823567ffffffffffffffff8111156142be57600080fd5b6142ca85828601613b3e565b90969095509350505050565b600080600080606085870312156142ec57600080fd5b6142f585613b15565b935060208501359250604085013567ffffffffffffffff81111561426357600080fd5b60006020828403121561432a57600080fd5b813567ffffffffffffffff81111561434157600080fd5b61434d84828501613d62565b949350505050565b60008083601f84011261436757600080fd5b50813567ffffffffffffffff81111561437f57600080fd5b6020830191508360208260051b8501011115613b8057600080fd5b600080602083850312156143ad57600080fd5b823567ffffffffffffffff8111156143c457600080fd5b6142ca85828601614355565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000815180845260005b8181101561442557602081850181015186830182015201614409565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015614535577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051600681106144f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261451260608801826143ff565b60409283015197909201969096529450602093840193919091019060010161448b565b50929695505050505050565b6000806040838503121561455457600080fd5b823561455f816141bf565b915061456d60208401613b15565b90509250929050565b60008060008060008060008060a0898b03121561459257600080fd5b61459b89613b15565b97506145a960208a01613b15565b9650604089013567ffffffffffffffff8111156145c557600080fd5b6145d18b828c01614355565b909750955050606089013567ffffffffffffffff8111156145f157600080fd5b6145fd8b828c01614355565b909550935050608089013567ffffffffffffffff811115613c1757600080fd5b60008060008060008060a0878903121561463657600080fd5b61463f87613b15565b955061464d60208801613b15565b94506040870135935060608701359250608087013567ffffffffffffffff81111561467757600080fd5b61468389828a01613b3e565b979a9699509497509295939492505050565b6000806000606084860312156146aa57600080fd5b833592506146ba60208501613b15565b915060408401356bffffffffffffffffffffffff811681146146db57600080fd5b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156106f2576106f26146e6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825180855260208501945060208160051b8301016020850160005b83811015614827577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e060408601526147e360e08601826143ff565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101614775565b50909695505050505050565b600081518084526020840193506020830160005b8281101561487b57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101614847565b5093949350505050565b805160ff168252600060208201516148a1602085018215159052565b50604082015161012060408501526148bd610120850182614757565b9050606083015160608501526080830151608085015260a083015184820360a08601526148ea82826143ff565b91505060c083015160c085015260e083015160e085015261010083015184820361010086015261491a8282614833565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60408152600061497f6040830186614885565b8281036020840152614992818587614923565b9695505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126149d157600080fd5b83018035915067ffffffffffffffff8211156149ec57600080fd5b602001915036819003821315613b8057600080fd5b83815260406020820152600061491a604083018486614923565b600060208284031215614a2d57600080fd5b81516108a8816141bf565b6020815260006108a860208301846143ff565b60208152600061434d602083018486614923565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21833603018112614a9357600080fd5b9190910192915050565b60006106f23683613df0565b60ff8716815285602082015284604082015283606082015260ff8316608082015260c060a08201526000614ae060c08301846143ff565b98975050505050505050565b818103818111156106f2576106f26146e6565b60008085851115614b0f57600080fd5b83861115614b1c57600080fd5b5050820193919092039150565b606081526000614b3c6060830186614885565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000614ae060c08301846143ff565b83815282602082015260606040820152600061491a60608301846143ff565b606081526000614bb16060830186614885565b846020840152828103604084015261499281856143ff565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061434d60408301846143ff565b60006040828403128015614c0b57600080fd5b506040805190810167ffffffffffffffff81118282101715614c2f57614c2f613c52565b604052825181526020928301519281019290925250919050565b8151600090829060208501835b82811015614c8a57815173ffffffffffffffffffffffffffffffffffffffff16845260209384019390910190600101614c56565b509195945050505050565b606081526000614ca9606083018688614923565b6020830194909452506040015292915050565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000614992606083018486614923565b600060208284031215614d0457600080fd5b5051919050565b8151600090829060208501835b82811015614c8a578151845260209384019390910190600101614d18565b60ff81811683821601908111156106f2576106f26146e656fea2646970667358221220cd4d3ff6f35d824793bb81ca683112170d9fbcc498509bbd60f84905ffb3b13964736f6c634300081c0033"
