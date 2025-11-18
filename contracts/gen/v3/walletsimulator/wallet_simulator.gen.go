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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulate\",\"inputs\":[{\"name\":\"_calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"structSimulator.Result[]\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSimulator.Status\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainedSignatureNestedInChainedSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidPackedLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a034607457601f6144ea38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b038116810360745760805260405161445a9081610090823960805181818161105e015281816113870152611b5e0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c612233565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc146101575780639c145aed14610152578063a27c39221461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e576116d5565b611648565b611576565b611448565b6113fc565b6113ab565b61133c565b611239565b610feb565b610f8d565b610f51565b610ecd565b610e9e565b610dfa565b610cde565b610c25565b610b14565b610ab1565b6109fc565b610974565b6108e7565b6107ea565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b919060e0838203126101c55761053c610445565b9261054681610210565b845260208101356020850152604081013567ffffffffffffffff81116101c55760c09261057491830161050a565b60408501526060810135606085015261058f60808201610474565b60808501526105a060a08201610474565b60a0850152013560c0830152565b9080601f830112156101c55781356105c581610481565b926105d36040519485610404565b81845260208085019260051b820101918383116101c55760208201905b8382106105ff57505050505090565b813567ffffffffffffffff81116101c55760209161062287848094880101610528565b8152019101906105f0565b9080601f830112156101c557813561064481610481565b926106526040519485610404565b81845260208085019260051b8201019283116101c557602001905b82821061067a5750505090565b6020809161068784610210565b81520191019061066d565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c557610704610456565b9061071181600401610466565b825261071f60248201610474565b6020830152604481013567ffffffffffffffff81116101c557846004610747928401016105ae565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c5578460046107839284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107c093010161062d565b610100820152916024359067ffffffffffffffff82116101c5576107e691600401610231565b9091565b346101c5576107f836610692565b909161010081019261081361080e8551516117dd565b6117fd565b9160005b8551805182101561087a579061087461084f6108358360019561187b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610859838861187b565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610817565b505083838661088f336108598351518561187b565b5261089b818484612305565b50156108ad5760405160018152602090f35b6108e3906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611a77565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761091e6101a2565b506109276101ca565b5060643567ffffffffffffffff81116101c557610948903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109cc6109d2923690600401610231565b91611a9c565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a7f60209160243560443591600401611b45565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610af6600435610af181610a87565b612515565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c557610b5e903690600401610231565b60243567ffffffffffffffff81116101c557610b7e903690600401610231565b9160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610bfb57610bd59360027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55611caf565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610cb4576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610d1481610a87565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d3982612515565b1615610d9f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d90600082613293565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107e691600401610231565b610ea736610e55565b9030330361036a57610ebd61001c925a9261264a565b90610ec782612cc1565b906129ce565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610f046101a2565b5060443567ffffffffffffffff81116101c557610f25903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a7f600435612d55565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610fc9600435612d9b565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101c557610ff936610e55565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610bfb5760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561115057330361112257303b156101c5576110cc9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611da5565b038183305af1801561111d576111025760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b80611111600061111793610404565b80610def565b38610bd5565b611ae8565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919082519283825260005b8481106112245750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016111e5565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c55761128b61129191369060040161117a565b90611ed3565b60405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106112c55784840385f35b9091927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0858203018652835190815160068110156113375760019282602093928493526040806113228585015160608786015260608501906111da565b930151910152950196019101949190946112b8565b6111ab565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061141860008061141236610692565b91612eac565b92611424839293614085565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561147e81610a87565b6114866101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff6114ac82612515565b16611526577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166115198183613293565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115ad6101a2565b506115b66101ca565b5060443567ffffffffffffffff81116101c5576115d790369060040161117a565b505060643567ffffffffffffffff81116101c5576115f990369060040161117a565b505060843567ffffffffffffffff81116101c55761161b903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761167f6101a2565b506116886101ca565b5060843567ffffffffffffffff81116101c5576116a9903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561170f6101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361179a7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b16178561320f565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90600182018092116117eb57565b6117ae565b919082018092116117eb57565b9061180782610481565b6118146040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118428294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561188f5760209160051b010190565b61184c565b9080602083519182815201916020808360051b8301019401926000915b8383106118c057505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611933604085015160e0604086015260e08501906111da565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906118b1565b906020808351928381520192019060005b8181106119885750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161197b565b805160ff16825261052591602082810151151590820152610100611a126119ec60408501516101206040860152610120850190611894565b606085015160608501526080850151608085015260a085015184820360a08601526111da565b9260c081015160c084015260e081015160e084015201519061010081840391015261196a565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611a8e61052594926040855260408501906119b4565b926020818503910152611a38565b90611ab99291611aaa6124ba565b906003825260e0820152612305565b5015611ae3577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b909173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908115611150578133036111225780611bf9575b5050611be97f1626ba7e00000000000000000000000000000000000000000000000000000000926109cc836101007fffffffff00000000000000000000000000000000000000000000000000000000950190611af4565b1603611bf457600090565b600190565b813b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af192831561111d576109cc7fffffffff0000000000000000000000000000000000000000000000000000000093611be9937f1626ba7e0000000000000000000000000000000000000000000000000000000096611c9a575b5093505092611b92565b806111116000611ca993610404565b38611c90565b91939290611cbe905a9361264a565b9160608301516080840151611cd282612d55565b818103611d7157509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881920190611d0b8282613254565b604080519182526020820192909252a1611d26828685612305565b929015611d3957506104549394506129ce565b836108e387926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611a77565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b916020610525938181520191611a38565b90611dc082610481565b611dcd6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611dfb8294610481565b0160005b818110611e0b57505050565b60405190606082019180831067ffffffffffffffff8411176103e357602092604052600081526060838201526000604082015282828601015201611dff565b919081101561188f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21813603018212156101c5570190565b610525903690610528565b919082039182116117eb57565b909260c092610525959460008452602084015260408301526060820152600060808201528160a082015201906111da565b90915a92600090611ee381611db6565b94825b828110611ef6575b505050509050565b611f09611f04828589611e4a565b611e8a565b93611f1760a0860151151590565b806121fb575b6121f1575060009360608101518015908115806121e8575b6121775783611f476080850151151590565b1561211857611ff290611ffa92935a94611f75875173ffffffffffffffffffffffffffffffffffffffff1690565b911561211257505a905b611fed89611fc160408a01518d6040519788947f4c4e814c00000000000000000000000000000000000000000000000000000000602087015260248601611ea2565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612e07565b915a90611e95565b6040612006858c61187b565b5101525b15612046575b50806120286120216001938a61187b565b5160019052565b612030612e19565b602061203c838b61187b565b5101525b01611ee6565b60c0018051156120db5760018151146120a257516002146120675738612010565b93945050505061208161207a828561187b565b5160039052565b602061209561208e612e19565b928561187b565b5101528038808080611eee565b5093958095506120c193506120ba925084915061187b565b5160049052565b60206120d56120ce612e19565b928461187b565b51015290565b509250600180936120f66120ef828a61187b565b5160029052565b6120fe612e19565b602061210a838b61187b565b510152612040565b90611f7f565b50611ff261215d915a93612140865173ffffffffffffffffffffffffffffffffffffffff1690565b9160208701519160001461217157505a905b604087015192612df5565b6040612169858c61187b565b51015261200a565b90612152565b5050509050859450602092506120d591506121a0612199826121b6969861187b565b5160059052565b6121e25a60405185810191825295869160200190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610404565b8461187b565b50805a10611f35565b9350600190612040565b508015611f1d565b3d1561222e573d9061221482610499565b916122226040519384610404565b82523d6000602084013e565b606090565b6004361080156122405750565b612276906000357fffffffff000000000000000000000000000000000000000000000000000000008116916122c8575b50612515565b73ffffffffffffffffffffffffffffffffffffffff81166122945750565b60008091604051368382378036810184815203915af46122b2612203565b90156122c057602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638612270565b901561188f5790565b61233861231284846122fc565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f8000000000000000000000000000000000000000000000000000000000000000808216146123bb5750612370926000928392612eac565b90509190919280821061238b57505061238890614085565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f02000000000000000000000000000000000000000000000000000000000000009081161460208201526123f192509050612cc1565b906123fb82612d9b565b42811115612488575073ffffffffffffffffffffffffffffffffffffffff8116801515908161247d575b50612431575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538612425565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a12084521660408201526040815261258c606082610404565b519020541690565b9061259e82610481565b6125ab6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125d98294610481565b019060005b8281106125ea57505050565b6020906040516125f9816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c0820152828285010152016125de565b909392938483116101c55784116101c5578101920390565b906126536124ba565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016129225750600060608601525b60076126a260ff831660011c90565b16806128d1575b506010818116036128a3575060015b6126c181612594565b604086019081526000925b8284106127095750505050036126df5790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361288157506127533061273784875161187b565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280821614612861575b600480821614612813575b6008808216146127de575b906127c56127bf60c08461279f60108060019816146080612796888c5161187b565b51019015159052565b6127b560208083161460a0612796888c5161187b565b1660061c60031690565b60ff1690565b60c06127d283875161187b565b510152019291906126cc565b94600191906127c5906127bf9060c090868101359060200199906060612805878b5161187b565b510152939450505050612774565b9461285b908381013560e81c9060030161285461283b61283384846117f0565b838c89612632565b9190604061284a888b5161187b565b51019236916104d3565b90526117f0565b94612769565b9482810135906020019590602061287984875161187b565b51015261275e565b61289e96508381013560601c90601401969061273784875161187b565b612753565b6020908116036128c057600282019181013560f01c905b906126b8565b600182019181013560f81c906128ba565b612915919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92906080860152386126a9565b80830135606090811c908701526014019250612693565b612951604092959493956060835260608301906119b4565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a082015201906111da565b61052593926060928252602082015281604082015201906111da565b6129bb61052594926060835260608301906119b4565b92602082015260408184039101526111da565b916000604082019384515190825b8281106129ed575b50505050505050565b6129f881885161187b565b5193612a0760a0860151151590565b80612cb9575b612c79575060009360608101518015801580612c70575b612c38578490612a376080850151151590565b15612bf257612ab192612a5e855173ffffffffffffffffffffffffffffffffffffffff1690565b9115612bec57505a905b611fed8b611fc160608d01516040890151908c8b604051998a967f4c4e814c0000000000000000000000000000000000000000000000000000000060208901526024880161295b565b15612af4575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016129dc565b60c001805115612ba3576001815114612b645751600214612b155738612ab7565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250612b55612b49612e19565b60405193849384612989565b0390a1388080808080806129e4565b50846108e3612b71612e19565b6040519384937f7f6b0bb1000000000000000000000000000000000000000000000000000000008552600485016129a5565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612be4612bd6612e19565b604051918291858c84612989565b0390a1612aee565b90612a68565b8351612c2d93925073ffffffffffffffffffffffffffffffffffffffff1691602085015191600014612c3257505a905b604085015192612df5565b612ab1565b90612c22565b83886108e35a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612939565b50815a10612a24565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612be4565b508015612a0d565b612d23612d4f612ce1612cdb6020850151151530906132d2565b936133cd565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612d94606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612dda606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405190612e41826103e8565b60006020838281520152565b6003111561133757565b908160409103126101c557602060405191612e71836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff610525949316815281602082015201906111da565b909491939291853560f81c600190938190612ec5612e34565b92612ecf82612e4d565b600182036130d9575b50600180871614613078575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612f479060051c90565b612f50906117dd565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612f9c84612cc1565b988993612fa893612632565b91612fb29361380f565b9098612fc691600052602052604060002090565b90612fd991600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff1661300291600052602052604060002090565b9481519086821515928361306d575b50508161305e575b506130215750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538613019565b141591508638613011565b90969193945061308a81989398612e4d565b6130af576130a4958161309c93612632565b9390926135c2565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612ed8578981013560601c9750601401915087878a8461312285612e4d565b60028503613133575b505050612ed8565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c926040926131b89290916131839161317c918a9061317689836117f0565b92612632565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612e7f565b0392165afa801561111d576131d6926000916131e0575b50936117f0565b9087388a8161312b565b613202915060403d604011613208575b6131fa8183610404565b810190612e57565b386131cf565b503d6131f0565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261324e606082610404565b51902055565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e835260408201526040815261324e606082610404565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120835260408201526040815261324e606082610404565b15613382576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a08152612d4f60c082610404565b46906132db565b805160209091019060005b8181106133a15750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613394565b6101008101516040516133e881612d23602082018095613389565b519020906133f7815160ff1690565b60ff81168061347057505090612d4f61341360408401516140e8565b92612d2360806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b600181036134ce57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612d4f8160808101612d23565b6002810361352457505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612d4f8160808101612d23565b600303613578575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612d4f8160808101612d23565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906129519060409396959496606084526060840191611a38565b91949290926000956000956000956000956000956135de6124ba565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061363557505050505050508051151580613627575b6130215750565b506020810151841115613620565b600381019d50959b5093995091975092909190613659908b9085013560e81c6117f0565b95828703613799578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361376a57506136a69161369f898c938789612632565b908b612eac565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106137245750928b88511461371b575b808b10156136e957508a60c08501528992959295949194939093613608565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526136ca565b8d8f6108e361373585858c8e612632565b9390926040519485947fb006aba0000000000000000000000000000000000000000000000000000000008652600486016135a8565b979899809b9261377f8b61378694888a612632565b9086612eac565b50929d919c909b929a9092918e8e6136ba565b8a600291613666565b908160209103126101c5575161052581610a87565b604090610525949281528160208201520191611a38565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611a38565b908160209103126101c5575190565b9391909360009460009460005b81811061382a575050505050565b8481013560f881901c9860019092019788979692909160fc1c98891561401e575060018914613fde5760028914613e115760038914613de25760048914613d615760068914613cc15760058914613c735760078914613bac5760088914613b565760098914613a2d57600a89146138ca577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613a1c575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613942918787612632565b6040517f898bd9210000000000000000000000000000000000000000000000000000000081529391849161397a918a600485016137b7565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561111d576139bd936000936139e9575b5060ff909a168091019a6143b4565b9080156139e357906139d791600052602052604060002090565b955b939291909361381c565b506139d7565b60ff919350613a0e9060203d8111613a15575b613a068183610404565b810190613800565b92906139ae565b503d6139fc565b8084013560f81c98506001016138df565b90919293949596975060038916978815613b45575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613aa5918787612632565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613add918b60048501611a77565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561111d57613b1f936000936139e9575060ff909a168091019a6143b4565b908015613b3f5790613b3991600052602052604060002090565b956139d9565b50613b39565b8084013560f81c9850600101613a42565b985060208701975094959394929391929091820135613b748661435b565b8114613b84575b613b1f90614375565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613b7b565b975090919293949597600f16968715613c61575b60206000613bd2613c3f9a9b86614223565b9c9092918a604051613c1581612d238a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561111d57613b1f9060ff6000519a1680910199614266565b600189019883013560f81c9750613bc0565b985060208701975094959394929391929091820135808514613c99575b613b1f9061431c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613c90565b989091929394959662ffffff9850613ce36127bf600c8416603f9060021c1690565b918215613d4d575b6003168015613d3c575b908190613d2090613d18908781013560e81c906003019c168c01809c8989612632565b90898b61380f565b911115613d33575b90613b1f92916142d1565b99820199613d28565b50600281019084013560f01c613cf5565b8482013560f81c9250600190910190613ceb565b975097613db7613dc4929394959697600f613dcc93169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686612632565b90868861380f565b90613b3992980198600052602052604060002090565b985096509394929391929091908082013590602001968015613b3f5790613b3991600052602052604060002090565b90919293949596975060038916978815613fcd575b8084013560601c99613e859160140190613e459060021c6003166127bf565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613ee960208c613e9b85858b8b612632565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e000000000000000000000000000000000000000000000000000000008552600485016137b7565b0392165afa90811561111d577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613f9f575b501603613f5b57509060ff613b1f92991680910199614266565b6108e3613f6c8c9389938989612632565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016137ce565b613fc0915060203d8111613fc6575b613fb88183610404565b8101906137a2565b38613f41565b503d613fae565b8381013560f81c9850600101613e26565b98600f9192939495969798501696871561400d575b6014810197613b1f9160ff9091169084013560601c614266565b8281013560f81c9750600101613ff3565b98509091929394959698600f16978815614070575b5060206000614046613c3f9a9b86614223565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020614033565b8015159081614092575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b8181106140d25750505090565b82518452602093840193909201916001016140c5565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061412e61411883610481565b926141266040519485610404565b808452610481565b0136602083013760005b835181101561420a578061414e6001928661187b565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e083015261010082015261010081526141f661012082610404565b519020614203828561187b565b5201614138565b50909150604051612d4f81612d236020820180956140ba565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116117eb5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612d4f606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612d4f609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612d4f605882610404565b612d23612d4f612ce1612cdb6000602086015115156132d2565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612d4f606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612d4f608d8261040456fea26469706673582212205ab89f8cc4ae4fd38cd73527005144490aa053823dceb239295470e37d85cf6264736f6c634300081c0033",
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
const WalletSimulatorDeployedBin = "0x6080604052600436101561001e575b361561001c5761001c612233565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc146101575780639c145aed14610152578063a27c39221461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e576116d5565b611648565b611576565b611448565b6113fc565b6113ab565b61133c565b611239565b610feb565b610f8d565b610f51565b610ecd565b610e9e565b610dfa565b610cde565b610c25565b610b14565b610ab1565b6109fc565b610974565b6108e7565b6107ea565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b919060e0838203126101c55761053c610445565b9261054681610210565b845260208101356020850152604081013567ffffffffffffffff81116101c55760c09261057491830161050a565b60408501526060810135606085015261058f60808201610474565b60808501526105a060a08201610474565b60a0850152013560c0830152565b9080601f830112156101c55781356105c581610481565b926105d36040519485610404565b81845260208085019260051b820101918383116101c55760208201905b8382106105ff57505050505090565b813567ffffffffffffffff81116101c55760209161062287848094880101610528565b8152019101906105f0565b9080601f830112156101c557813561064481610481565b926106526040519485610404565b81845260208085019260051b8201019283116101c557602001905b82821061067a5750505090565b6020809161068784610210565b81520191019061066d565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c557610704610456565b9061071181600401610466565b825261071f60248201610474565b6020830152604481013567ffffffffffffffff81116101c557846004610747928401016105ae565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c5578460046107839284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107c093010161062d565b610100820152916024359067ffffffffffffffff82116101c5576107e691600401610231565b9091565b346101c5576107f836610692565b909161010081019261081361080e8551516117dd565b6117fd565b9160005b8551805182101561087a579061087461084f6108358360019561187b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610859838861187b565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610817565b505083838661088f336108598351518561187b565b5261089b818484612305565b50156108ad5760405160018152602090f35b6108e3906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611a77565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761091e6101a2565b506109276101ca565b5060643567ffffffffffffffff81116101c557610948903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109cc6109d2923690600401610231565b91611a9c565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a7f60209160243560443591600401611b45565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610af6600435610af181610a87565b612515565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c557610b5e903690600401610231565b60243567ffffffffffffffff81116101c557610b7e903690600401610231565b9160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610bfb57610bd59360027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55611caf565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610cb4576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610d1481610a87565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d3982612515565b1615610d9f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d90600082613293565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107e691600401610231565b610ea736610e55565b9030330361036a57610ebd61001c925a9261264a565b90610ec782612cc1565b906129ce565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610f046101a2565b5060443567ffffffffffffffff81116101c557610f25903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a7f600435612d55565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610fc9600435612d9b565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101c557610ff936610e55565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610bfb5760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561115057330361112257303b156101c5576110cc9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611da5565b038183305af1801561111d576111025760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b80611111600061111793610404565b80610def565b38610bd5565b611ae8565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b919082519283825260005b8481106112245750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016111e5565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c55761128b61129191369060040161117a565b90611ed3565b60405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106112c55784840385f35b9091927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0858203018652835190815160068110156113375760019282602093928493526040806113228585015160608786015260608501906111da565b930151910152950196019101949190946112b8565b6111ab565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061141860008061141236610692565b91612eac565b92611424839293614085565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561147e81610a87565b6114866101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff6114ac82612515565b16611526577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166115198183613293565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115ad6101a2565b506115b66101ca565b5060443567ffffffffffffffff81116101c5576115d790369060040161117a565b505060643567ffffffffffffffff81116101c5576115f990369060040161117a565b505060843567ffffffffffffffff81116101c55761161b903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761167f6101a2565b506116886101ca565b5060843567ffffffffffffffff81116101c5576116a9903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561170f6101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361179a7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b16178561320f565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90600182018092116117eb57565b6117ae565b919082018092116117eb57565b9061180782610481565b6118146040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118428294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561188f5760209160051b010190565b61184c565b9080602083519182815201916020808360051b8301019401926000915b8383106118c057505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611933604085015160e0604086015260e08501906111da565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906118b1565b906020808351928381520192019060005b8181106119885750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161197b565b805160ff16825261052591602082810151151590820152610100611a126119ec60408501516101206040860152610120850190611894565b606085015160608501526080850151608085015260a085015184820360a08601526111da565b9260c081015160c084015260e081015160e084015201519061010081840391015261196a565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611a8e61052594926040855260408501906119b4565b926020818503910152611a38565b90611ab99291611aaa6124ba565b906003825260e0820152612305565b5015611ae3577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b909173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908115611150578133036111225780611bf9575b5050611be97f1626ba7e00000000000000000000000000000000000000000000000000000000926109cc836101007fffffffff00000000000000000000000000000000000000000000000000000000950190611af4565b1603611bf457600090565b600190565b813b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af192831561111d576109cc7fffffffff0000000000000000000000000000000000000000000000000000000093611be9937f1626ba7e0000000000000000000000000000000000000000000000000000000096611c9a575b5093505092611b92565b806111116000611ca993610404565b38611c90565b91939290611cbe905a9361264a565b9160608301516080840151611cd282612d55565b818103611d7157509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881920190611d0b8282613254565b604080519182526020820192909252a1611d26828685612305565b929015611d3957506104549394506129ce565b836108e387926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611a77565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b916020610525938181520191611a38565b90611dc082610481565b611dcd6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611dfb8294610481565b0160005b818110611e0b57505050565b60405190606082019180831067ffffffffffffffff8411176103e357602092604052600081526060838201526000604082015282828601015201611dff565b919081101561188f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21813603018212156101c5570190565b610525903690610528565b919082039182116117eb57565b909260c092610525959460008452602084015260408301526060820152600060808201528160a082015201906111da565b90915a92600090611ee381611db6565b94825b828110611ef6575b505050509050565b611f09611f04828589611e4a565b611e8a565b93611f1760a0860151151590565b806121fb575b6121f1575060009360608101518015908115806121e8575b6121775783611f476080850151151590565b1561211857611ff290611ffa92935a94611f75875173ffffffffffffffffffffffffffffffffffffffff1690565b911561211257505a905b611fed89611fc160408a01518d6040519788947f4c4e814c00000000000000000000000000000000000000000000000000000000602087015260248601611ea2565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612e07565b915a90611e95565b6040612006858c61187b565b5101525b15612046575b50806120286120216001938a61187b565b5160019052565b612030612e19565b602061203c838b61187b565b5101525b01611ee6565b60c0018051156120db5760018151146120a257516002146120675738612010565b93945050505061208161207a828561187b565b5160039052565b602061209561208e612e19565b928561187b565b5101528038808080611eee565b5093958095506120c193506120ba925084915061187b565b5160049052565b60206120d56120ce612e19565b928461187b565b51015290565b509250600180936120f66120ef828a61187b565b5160029052565b6120fe612e19565b602061210a838b61187b565b510152612040565b90611f7f565b50611ff261215d915a93612140865173ffffffffffffffffffffffffffffffffffffffff1690565b9160208701519160001461217157505a905b604087015192612df5565b6040612169858c61187b565b51015261200a565b90612152565b5050509050859450602092506120d591506121a0612199826121b6969861187b565b5160059052565b6121e25a60405185810191825295869160200190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610404565b8461187b565b50805a10611f35565b9350600190612040565b508015611f1d565b3d1561222e573d9061221482610499565b916122226040519384610404565b82523d6000602084013e565b606090565b6004361080156122405750565b612276906000357fffffffff000000000000000000000000000000000000000000000000000000008116916122c8575b50612515565b73ffffffffffffffffffffffffffffffffffffffff81166122945750565b60008091604051368382378036810184815203915af46122b2612203565b90156122c057602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638612270565b901561188f5790565b61233861231284846122fc565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f8000000000000000000000000000000000000000000000000000000000000000808216146123bb5750612370926000928392612eac565b90509190919280821061238b57505061238890614085565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f02000000000000000000000000000000000000000000000000000000000000009081161460208201526123f192509050612cc1565b906123fb82612d9b565b42811115612488575073ffffffffffffffffffffffffffffffffffffffff8116801515908161247d575b50612431575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538612425565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a12084521660408201526040815261258c606082610404565b519020541690565b9061259e82610481565b6125ab6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06125d98294610481565b019060005b8281106125ea57505050565b6020906040516125f9816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c0820152828285010152016125de565b909392938483116101c55784116101c5578101920390565b906126536124ba565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016129225750600060608601525b60076126a260ff831660011c90565b16806128d1575b506010818116036128a3575060015b6126c181612594565b604086019081526000925b8284106127095750505050036126df5790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361288157506127533061273784875161187b565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280821614612861575b600480821614612813575b6008808216146127de575b906127c56127bf60c08461279f60108060019816146080612796888c5161187b565b51019015159052565b6127b560208083161460a0612796888c5161187b565b1660061c60031690565b60ff1690565b60c06127d283875161187b565b510152019291906126cc565b94600191906127c5906127bf9060c090868101359060200199906060612805878b5161187b565b510152939450505050612774565b9461285b908381013560e81c9060030161285461283b61283384846117f0565b838c89612632565b9190604061284a888b5161187b565b51019236916104d3565b90526117f0565b94612769565b9482810135906020019590602061287984875161187b565b51015261275e565b61289e96508381013560601c90601401969061273784875161187b565b612753565b6020908116036128c057600282019181013560f01c905b906126b8565b600182019181013560f81c906128ba565b612915919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92906080860152386126a9565b80830135606090811c908701526014019250612693565b612951604092959493956060835260608301906119b4565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a082015201906111da565b61052593926060928252602082015281604082015201906111da565b6129bb61052594926060835260608301906119b4565b92602082015260408184039101526111da565b916000604082019384515190825b8281106129ed575b50505050505050565b6129f881885161187b565b5193612a0760a0860151151590565b80612cb9575b612c79575060009360608101518015801580612c70575b612c38578490612a376080850151151590565b15612bf257612ab192612a5e855173ffffffffffffffffffffffffffffffffffffffff1690565b9115612bec57505a905b611fed8b611fc160608d01516040890151908c8b604051998a967f4c4e814c0000000000000000000000000000000000000000000000000000000060208901526024880161295b565b15612af4575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016129dc565b60c001805115612ba3576001815114612b645751600214612b155738612ab7565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250612b55612b49612e19565b60405193849384612989565b0390a1388080808080806129e4565b50846108e3612b71612e19565b6040519384937f7f6b0bb1000000000000000000000000000000000000000000000000000000008552600485016129a5565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612be4612bd6612e19565b604051918291858c84612989565b0390a1612aee565b90612a68565b8351612c2d93925073ffffffffffffffffffffffffffffffffffffffff1691602085015191600014612c3257505a905b604085015192612df5565b612ab1565b90612c22565b83886108e35a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612939565b50815a10612a24565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612be4565b508015612a0d565b612d23612d4f612ce1612cdb6020850151151530906132d2565b936133cd565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612d94606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612dda606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405190612e41826103e8565b60006020838281520152565b6003111561133757565b908160409103126101c557602060405191612e71836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff610525949316815281602082015201906111da565b909491939291853560f81c600190938190612ec5612e34565b92612ecf82612e4d565b600182036130d9575b50600180871614613078575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612f479060051c90565b612f50906117dd565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612f9c84612cc1565b988993612fa893612632565b91612fb29361380f565b9098612fc691600052602052604060002090565b90612fd991600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff1661300291600052602052604060002090565b9481519086821515928361306d575b50508161305e575b506130215750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538613019565b141591508638613011565b90969193945061308a81989398612e4d565b6130af576130a4958161309c93612632565b9390926135c2565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612ed8578981013560601c9750601401915087878a8461312285612e4d565b60028503613133575b505050612ed8565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c926040926131b89290916131839161317c918a9061317689836117f0565b92612632565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612e7f565b0392165afa801561111d576131d6926000916131e0575b50936117f0565b9087388a8161312b565b613202915060403d604011613208575b6131fa8183610404565b810190612e57565b386131cf565b503d6131f0565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261324e606082610404565b51902055565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e835260408201526040815261324e606082610404565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120835260408201526040815261324e606082610404565b15613382576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a08152612d4f60c082610404565b46906132db565b805160209091019060005b8181106133a15750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613394565b6101008101516040516133e881612d23602082018095613389565b519020906133f7815160ff1690565b60ff81168061347057505090612d4f61341360408401516140e8565b92612d2360806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b600181036134ce57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612d4f8160808101612d23565b6002810361352457505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612d4f8160808101612d23565b600303613578575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612d4f8160808101612d23565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906129519060409396959496606084526060840191611a38565b91949290926000956000956000956000956000956135de6124ba565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061363557505050505050508051151580613627575b6130215750565b506020810151841115613620565b600381019d50959b5093995091975092909190613659908b9085013560e81c6117f0565b95828703613799578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361376a57506136a69161369f898c938789612632565b908b612eac565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106137245750928b88511461371b575b808b10156136e957508a60c08501528992959295949194939093613608565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526136ca565b8d8f6108e361373585858c8e612632565b9390926040519485947fb006aba0000000000000000000000000000000000000000000000000000000008652600486016135a8565b979899809b9261377f8b61378694888a612632565b9086612eac565b50929d919c909b929a9092918e8e6136ba565b8a600291613666565b908160209103126101c5575161052581610a87565b604090610525949281528160208201520191611a38565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611a38565b908160209103126101c5575190565b9391909360009460009460005b81811061382a575050505050565b8481013560f881901c9860019092019788979692909160fc1c98891561401e575060018914613fde5760028914613e115760038914613de25760048914613d615760068914613cc15760058914613c735760078914613bac5760088914613b565760098914613a2d57600a89146138ca577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613a1c575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613942918787612632565b6040517f898bd9210000000000000000000000000000000000000000000000000000000081529391849161397a918a600485016137b7565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561111d576139bd936000936139e9575b5060ff909a168091019a6143b4565b9080156139e357906139d791600052602052604060002090565b955b939291909361381c565b506139d7565b60ff919350613a0e9060203d8111613a15575b613a068183610404565b810190613800565b92906139ae565b503d6139fc565b8084013560f81c98506001016138df565b90919293949596975060038916978815613b45575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613aa5918787612632565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613add918b60048501611a77565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561111d57613b1f936000936139e9575060ff909a168091019a6143b4565b908015613b3f5790613b3991600052602052604060002090565b956139d9565b50613b39565b8084013560f81c9850600101613a42565b985060208701975094959394929391929091820135613b748661435b565b8114613b84575b613b1f90614375565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613b7b565b975090919293949597600f16968715613c61575b60206000613bd2613c3f9a9b86614223565b9c9092918a604051613c1581612d238a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561111d57613b1f9060ff6000519a1680910199614266565b600189019883013560f81c9750613bc0565b985060208701975094959394929391929091820135808514613c99575b613b1f9061431c565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613c90565b989091929394959662ffffff9850613ce36127bf600c8416603f9060021c1690565b918215613d4d575b6003168015613d3c575b908190613d2090613d18908781013560e81c906003019c168c01809c8989612632565b90898b61380f565b911115613d33575b90613b1f92916142d1565b99820199613d28565b50600281019084013560f01c613cf5565b8482013560f81c9250600190910190613ceb565b975097613db7613dc4929394959697600f613dcc93169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686612632565b90868861380f565b90613b3992980198600052602052604060002090565b985096509394929391929091908082013590602001968015613b3f5790613b3991600052602052604060002090565b90919293949596975060038916978815613fcd575b8084013560601c99613e859160140190613e459060021c6003166127bf565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613ee960208c613e9b85858b8b612632565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e000000000000000000000000000000000000000000000000000000008552600485016137b7565b0392165afa90811561111d577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613f9f575b501603613f5b57509060ff613b1f92991680910199614266565b6108e3613f6c8c9389938989612632565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016137ce565b613fc0915060203d8111613fc6575b613fb88183610404565b8101906137a2565b38613f41565b503d613fae565b8381013560f81c9850600101613e26565b98600f9192939495969798501696871561400d575b6014810197613b1f9160ff9091169084013560601c614266565b8281013560f81c9750600101613ff3565b98509091929394959698600f16978815614070575b5060206000614046613c3f9a9b86614223565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020614033565b8015159081614092575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b8181106140d25750505090565b82518452602093840193909201916001016140c5565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061412e61411883610481565b926141266040519485610404565b808452610481565b0136602083013760005b835181101561420a578061414e6001928661187b565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e083015261010082015261010081526141f661012082610404565b519020614203828561187b565b5201614138565b50909150604051612d4f81612d236020820180956140ba565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116117eb5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612d4f606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612d4f609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612d4f605882610404565b612d23612d4f612ce1612cdb6000602086015115156132d2565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612d4f606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612d4f608d8261040456fea26469706673582212205ab89f8cc4ae4fd38cd73527005144490aa053823dceb239295470e37d85cf6264736f6c634300081c0033"
