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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"simulate\",\"inputs\":[{\"name\":\"_calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"structSimulator.Result[]\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSimulator.Status\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a034607457601f61434038819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b03811681036074576080526040516142b090816100908239608051818181610fe1015281816112db0152611adf0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c6120f4565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc146101575780639c145aed14610152578063a27c39221461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e57611629565b61159c565b6114ca565b61139c565b611350565b6112ff565b611290565b611163565b610fbb565b610f5d565b610f21565b610e9d565b610e6e565b610dca565b610cae565b610bf5565b610b14565b610ab1565b6109fc565b610974565b6108e7565b6107ea565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b919060e0838203126101c55761053c610445565b9261054681610210565b845260208101356020850152604081013567ffffffffffffffff81116101c55760c09261057491830161050a565b60408501526060810135606085015261058f60808201610474565b60808501526105a060a08201610474565b60a0850152013560c0830152565b9080601f830112156101c55781356105c581610481565b926105d36040519485610404565b81845260208085019260051b820101918383116101c55760208201905b8382106105ff57505050505090565b813567ffffffffffffffff81116101c55760209161062287848094880101610528565b8152019101906105f0565b9080601f830112156101c557813561064481610481565b926106526040519485610404565b81845260208085019260051b8201019283116101c557602001905b82821061067a5750505090565b6020809161068784610210565b81520191019061066d565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c557610704610456565b9061071181600401610466565b825261071f60248201610474565b6020830152604481013567ffffffffffffffff81116101c557846004610747928401016105ae565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c5578460046107839284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107c093010161062d565b610100820152916024359067ffffffffffffffff82116101c5576107e691600401610231565b9091565b346101c5576107f836610692565b909161010081019261081361080e855151611731565b611751565b9160005b8551805182101561087a579061087461084f610835836001956117cf565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61085983886117cf565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610817565b505083838661088f33610859835151856117cf565b5261089b8184846121c6565b50156108ad5760405160018152602090f35b6108e3906040519384937ff58cc8b5000000000000000000000000000000000000000000000000000000008552600485016119cb565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761091e6101a2565b506109276101ca565b5060643567ffffffffffffffff81116101c557610948903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109cc6109d2923690600401610231565b916119f0565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a7f60209160243560443591600401611ac5565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610af6600435610af181610a87565b6123ab565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c557610b5e903690600401610231565b9060243567ffffffffffffffff81116101c557610b82610b8c913690600401610231565b9190935a936124e0565b91610ba060608401516080850151906127a4565b610bab8285856121c6565b929015610bbd575061001c93506128fd565b836108e386926040519384937fa2b6d61b000000000000000000000000000000000000000000000000000000008552600485016119cb565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610c84576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610ce481610a87565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d09826123ab565b1615610d6f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d60600082613124565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107e691600401610231565b610e7736610e25565b9030330361036a57610e8d61001c925a926124e0565b90610e9782612bf0565b906128fd565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610ed46101a2565b5060443567ffffffffffffffff81116101c557610ef5903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a7f600435612c84565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610f99600435612cca565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101c557610fc936610e25565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156110a957330361107b57303b156101c55761104f9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611c66565b038183305af180156110765761106157005b80611070600061001c93610404565b80610dbf565b611a3c565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b919082519283825260005b84811061114e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161110f565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576111b56111bb9136906004016110d3565b90611d94565b60405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106111ef5784840385f35b9091927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08582030186528351908151600681101561126157600192826020939284935260408061124c858501516060878601526060850190611104565b930151910152950196019101949190946111e2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061136c60008061136636610692565b91612dd1565b92611378839293613edb565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356113d281610a87565b6113da6101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff611400826123ab565b1661147a577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000060409316911661146d8183613124565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115016101a2565b5061150a6101ca565b5060443567ffffffffffffffff81116101c55761152b9036906004016110d3565b505060643567ffffffffffffffff81116101c55761154d9036906004016110d3565b505060843567ffffffffffffffff81116101c55761156f903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115d36101a2565b506115dc6101ca565b5060843567ffffffffffffffff81116101c5576115fd903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356116636101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff836116ee7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b1617856130df565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161173f57565b611702565b9190820180921161173f57565b9061175b82610481565b6117686040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06117968294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117e35760209160051b010190565b6117a0565b9080602083519182815201916020808360051b8301019401926000915b83831061181457505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611887604085015160e0604086015260e0850190611104565b936060810151606085015260808101511515608085015260a0810151151560a0850152015191015297019301930191939290611805565b906020808351928381520192019060005b8181106118dc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016118cf565b805160ff16825261052591602082810151151590820152610100611966611940604085015161012060408601526101208501906117e8565b606085015160608501526080850151608085015260a085015184820360a0860152611104565b9260c081015160c084015260e081015160e08401520151906101008184039101526118be565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916119e26105259492604085526040850190611908565b92602081850391015261198c565b90611a0d92916119fe612350565b906003825260e08201526121c6565b5015611a37577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b908160209103126101c5575161052581610a87565b60409061052594928152816020820152019161198c565b91909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169182156110a95782330361107b5780611bf7575b506020915080610100611b27920190611a48565b92611b5f60405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611aae565b0381305afa908115611076577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611bc8575b501603611bc357600090565b600190565b611bea915060203d602011611bf0575b611be28183610404565b810190611a99565b38611bb7565b503d611bd8565b823b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af190811561107657602092611b2792611c51575b5090611b13565b806110706000611c6093610404565b38611c4a565b91602061052593818152019161198c565b90611c8182610481565b611c8e6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611cbc8294610481565b0160005b818110611ccc57505050565b60405190606082019180831067ffffffffffffffff8411176103e357602092604052600081526060838201526000604082015282828601015201611cc0565b91908110156117e35760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21813603018212156101c5570190565b610525903690610528565b9190820391821161173f57565b909260c092610525959460008452602084015260408301526060820152600060808201528160a08201520190611104565b90915a92600090611da481611c77565b94825b828110611db7575b505050509050565b611dca611dc5828589611d0b565b611d4b565b93611dd860a0860151151590565b806120bc575b6120b2575060009360608101518015908115806120a9575b6120385783611e086080850151151590565b15611fd957611eb390611ebb92935a94611e36875173ffffffffffffffffffffffffffffffffffffffff1690565b9115611fd357505a905b611eae89611e8260408a01518d6040519788947f4c4e814c00000000000000000000000000000000000000000000000000000000602087015260248601611d63565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612d36565b915a90611d56565b6040611ec7858c6117cf565b5101525b15611f07575b5080611ee9611ee26001938a6117cf565b5160019052565b611ef1612d48565b6020611efd838b6117cf565b5101525b01611da7565b60c001805115611f9c576001815114611f635751600214611f285738611ed1565b939450505050611f42611f3b82856117cf565b5160039052565b6020611f56611f4f612d48565b92856117cf565b5101528038808080611daf565b509395809550611f829350611f7b92508491506117cf565b5160049052565b6020611f96611f8f612d48565b92846117cf565b51015290565b50925060018093611fb7611fb0828a6117cf565b5160029052565b611fbf612d48565b6020611fcb838b6117cf565b510152611f01565b90611e40565b50611eb361201e915a93612001865173ffffffffffffffffffffffffffffffffffffffff1690565b9160208701519160001461203257505a905b604087015192612d24565b604061202a858c6117cf565b510152611ecb565b90612013565b505050905085945060209250611f96915061206161205a8261207796986117cf565b5160059052565b6120a35a60405185810191825295869160200190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610404565b846117cf565b50805a10611df6565b9350600190611f01565b508015611dde565b3d156120ef573d906120d582610499565b916120e36040519384610404565b82523d6000602084013e565b606090565b6004361080156121015750565b612137906000357fffffffff00000000000000000000000000000000000000000000000000000000811691612189575b506123ab565b73ffffffffffffffffffffffffffffffffffffffff81166121555750565b60008091604051368382378036810184815203915af46121736120c4565b901561218157602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638612131565b90156117e35790565b91907f80000000000000000000000000000000000000000000000000000000000000008061221d6121f785856121bd565b357fff000000000000000000000000000000000000000000000000000000000000001690565b161461227c57600091829161223194612dd1565b90509190919280821061224c57505061224990613edb565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b505061228790612bf0565b9061229182612cca565b4281111561231e575073ffffffffffffffffffffffffffffffffffffffff81168015159081612313575b506122c7575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b9050331415386122bb565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152612422606082610404565b519020541690565b9061243482610481565b6124416040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061246f8294610481565b019060005b82811061248057505050565b60209060405161248f816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201612474565b909392938483116101c55784116101c5578101920390565b9190916124eb612350565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183160161278d5750600060608701525b600761253a60ff831660011c90565b168061273c575b5060108181160361270757506001925b61255a8461242a565b92604087019384526000905b85821061257557505050505050565b8281013560f81c906001019091816001808516036126e557506125b93061259d8389516117cf565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808416146126c5575b600480841614612677575b600880841614612644575b61262d61262760c0856126068a60806125fd8860108060019c9d161493516117cf565b51019015159052565b61261d8a60a06125fd8860208087161493516117cf565b1660061c60031690565b60ff1690565b60c061263a8389516117cf565b5101520190612566565b60019161262d906126279060c090878101359060200196906060612669878d516117cf565b5101529594505050506125da565b906126bf908481013560e81c906003016126b861269f6126978484611744565b838a8a6124c8565b919060406126ae888d516117cf565b51019236916104d3565b9052611744565b906125cf565b908381013590602001919060206126dd8389516117cf565b5101526125c4565b61270292508481013560601c90601401929061259d8389516117cf565b6125b9565b6020908116036127295761ffff918381013560f01c906002015b921692612551565b60ff918381013560f81c90600101612721565b612780919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608087015238612541565b80850135606090811c90880152601401925061252b565b906127ae82612c84565b81810361283457509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e82528360408201526040815261281e606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61288060409295949395606083526060830190611908565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611104565b6105259392606092825260208201528160408201520190611104565b6128ea6105259492606083526060830190611908565b9260208201526040818403910152611104565b916000604082019384515190825b82811061291c575b50505050505050565b6129278188516117cf565b519361293660a0860151151590565b80612be8575b612ba8575060009360608101518015801580612b9f575b612b675784906129666080850151151590565b15612b21576129e09261298d855173ffffffffffffffffffffffffffffffffffffffff1690565b9115612b1b57505a905b611eae8b611e8260608d01516040890151908c8b604051998a967f4c4e814c0000000000000000000000000000000000000000000000000000000060208901526024880161288a565b15612a23575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b0161290b565b60c001805115612ad2576001815114612a935751600214612a4457386129e6565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250612a84612a78612d48565b604051938493846128b8565b0390a138808080808080612913565b50846108e3612aa0612d48565b6040519384937f7f6b0bb1000000000000000000000000000000000000000000000000000000008552600485016128d4565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612b13612b05612d48565b604051918291858c846128b8565b0390a1612a1d565b90612997565b8351612b5c93925073ffffffffffffffffffffffffffffffffffffffff1691602085015191600014612b6157505a905b604085015192612d24565b6129e0565b90612b51565b83886108e35a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612868565b50815a10612953565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612b13565b50801561293c565b612c52612c7e612c10612c0a602085015115153090613163565b9361325e565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612cc3606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612d09606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405190612d70826103e8565b60006020838281520152565b908160409103126101c557602060405191612d96836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611104565b909491939291853560f81c60019093819087612deb612d63565b93604080891614806130c1575b612fc4575b505050600180861614612f9e5760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612e6d9060051c90565b612e7690611731565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612ec284612bf0565b988993612ece936124c8565b91612ed893613675565b9098612eec91600052602052604060002090565b90612eff91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612f2891600052602052604060002090565b94815190868215159283612f93575b505081612f84575b50612f475750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612f3f565b141591508638612f37565b86612fb99792949550612fb193986124c8565b939092613453565b919394909293929190565b8a81013560601c985060140192509087908a90849015612fe8575b50829150612dfd565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c9260409261306b9290916130369161302f9189906130298983611744565b926124c8565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612da4565b0392165afa80156110765761308992600091613092575b5092611744565b86388981612fdf565b6130b4915060403d6040116130ba575b6130ac8183610404565b810190612d7c565b38613082565b503d6130a2565b5073ffffffffffffffffffffffffffffffffffffffff891615612df8565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261311e606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120835260408201526040815261311e606082610404565b15613213576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a08152612c7e60c082610404565b469061316c565b805160209091019060005b8181106132325750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613225565b61010081015160405161327981612c5260208201809561321a565b51902090613288815160ff1690565b60ff81168061330157505090612c7e6132a46040840151613f3e565b92612c5260806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b6001810361335f57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612c7e8160808101612c52565b600281036133b557505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612c7e8160808101612c52565b600303613409575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612c7e8160808101612c52565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b90612880906040939695949660608452606084019161198c565b919492909260009560009560009560009560009561346f612350565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b8281106134c6575050505050505080511515806134b8575b612f475750565b5060208101518411156134b1565b600381019d50959b50939950919750929091906134ea908b9085013560e81c611744565b958a600084890361362e575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c036135fe57509060016135338961353a9487896124c8565b908b612dd1565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106135b85750928b8851146135af575b808b101561357d57508a60c08501528992959295949194939093613499565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b6000885261355e565b8d8f6108e36135c985858c8e6124c8565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601613439565b9798999a9160016136148b61361b94888a6124c8565b9086612dd1565b50929d919c909b929a9092918e8e61354e565b916134f9565b73ffffffffffffffffffffffffffffffffffffffff61052595936060938352166020820152816040820152019161198c565b908160209103126101c5575190565b9391909360009460009460005b818110613690575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613e74575060018914613e345760028914613c775760038914613c485760048914613bc75760068914613b275760058914613ad95760078914613a1257600889146139bc576009891461389357600a8914613730577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613882575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816137a89187876124c8565b6040517f898bd921000000000000000000000000000000000000000000000000000000008152939184916137e0918a60048501611aae565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611076576138239360009361384f575b5060ff909a168091019a61420a565b908015613849579061383d91600052602052604060002090565b955b9392919093613682565b5061383d565b60ff9193506138749060203d811161387b575b61386c8183610404565b810190613666565b9290613814565b503d613862565b8084013560f81c9850600101613745565b909192939495969750600389169788156139ab575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161390b9187876124c8565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613943918b600485016119cb565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611076576139859360009361384f575060ff909a168091019a61420a565b9080156139a5579061399f91600052602052604060002090565b9561383f565b5061399f565b8084013560f81c98506001016138a8565b9850602087019750949593949293919290918201356139da866141b1565b81146139ea575b613985906141cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506139e1565b975090919293949597600f16968715613ac7575b60206000613a38613aa59a9b86614079565b9c9092918a604051613a7b81612c528a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15611076576139859060ff6000519a16809101996140bc565b600189019883013560f81c9750613a26565b985060208701975094959394929391929091820135808514613aff575b61398590614172565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613af6565b989091929394959662ffffff9850613b49612627600c8416603f9060021c1690565b918215613bb3575b6003168015613ba2575b908190613b8690613b7e908781013560e81c906003019c168c01809c89896124c8565b90898b613675565b911115613b99575b906139859291614127565b99820199613b8e565b50600281019084013560f01c613b5b565b8482013560f81c9250600190910190613b51565b975097613c1d613c2a929394959697600f613c3293169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92908301809386866124c8565b908688613675565b9061399f92980198600052602052604060002090565b9850965093949293919290919080820135906020019680156139a5579061399f91600052602052604060002090565b90919293949596975060038916978815613e23575b8084013560601c99613ceb9160140190613cab9060021c600316612627565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613d4f60208c613d0185858b8b6124c8565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611aae565b0392165afa908115611076577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613e05575b501603613dc157509060ff613985929916809101996140bc565b6108e3613dd28c93899389896124c8565b906040519485947fb2fed7ae00000000000000000000000000000000000000000000000000000000865260048601613634565b613e1d915060203d8111611bf057611be28183610404565b38613da7565b8381013560f81c9850600101613c8c565b98600f91929394959697985016968715613e63575b60148101976139859160ff9091169084013560601c6140bc565b8281013560f81c9750600101613e49565b98509091929394959698600f16978815613ec6575b5060206000613e9c613aa59a9b86614079565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613e89565b8015159081613ee8575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b818110613f285750505090565b8251845260209384019390920191600101613f1b565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613f84613f6e83610481565b92613f7c6040519485610404565b808452610481565b0136602083013760005b83518110156140605780613fa4600192866117cf565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e0830152610100820152610100815261404c61012082610404565b51902061405982856117cf565b5201613f8e565b50909150604051612c7e81612c52602082018095613f10565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff811161173f5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612c7e606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612c7e609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612c7e605882610404565b612c52612c7e612c10612c0a600060208601511515613163565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612c7e606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612c7e608d8261040456fea2646970667358221220663e0e19a1f44876322f0807bffe0131225f2bf0bc3ca63984e33df77f73e01664736f6c634300081c0033",
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
const WalletSimulatorDeployedBin = "0x6080604052600436101561001e575b361561001c5761001c6120f4565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc146101575780639c145aed14610152578063a27c39221461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e57611629565b61159c565b6114ca565b61139c565b611350565b6112ff565b611290565b611163565b610fbb565b610f5d565b610f21565b610e9d565b610e6e565b610dca565b610cae565b610bf5565b610b14565b610ab1565b6109fc565b610974565b6108e7565b6107ea565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b919060e0838203126101c55761053c610445565b9261054681610210565b845260208101356020850152604081013567ffffffffffffffff81116101c55760c09261057491830161050a565b60408501526060810135606085015261058f60808201610474565b60808501526105a060a08201610474565b60a0850152013560c0830152565b9080601f830112156101c55781356105c581610481565b926105d36040519485610404565b81845260208085019260051b820101918383116101c55760208201905b8382106105ff57505050505090565b813567ffffffffffffffff81116101c55760209161062287848094880101610528565b8152019101906105f0565b9080601f830112156101c557813561064481610481565b926106526040519485610404565b81845260208085019260051b8201019283116101c557602001905b82821061067a5750505090565b6020809161068784610210565b81520191019061066d565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c557610704610456565b9061071181600401610466565b825261071f60248201610474565b6020830152604481013567ffffffffffffffff81116101c557846004610747928401016105ae565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c5578460046107839284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107c093010161062d565b610100820152916024359067ffffffffffffffff82116101c5576107e691600401610231565b9091565b346101c5576107f836610692565b909161010081019261081361080e855151611731565b611751565b9160005b8551805182101561087a579061087461084f610835836001956117cf565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61085983886117cf565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610817565b505083838661088f33610859835151856117cf565b5261089b8184846121c6565b50156108ad5760405160018152602090f35b6108e3906040519384937ff58cc8b5000000000000000000000000000000000000000000000000000000008552600485016119cb565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761091e6101a2565b506109276101ca565b5060643567ffffffffffffffff81116101c557610948903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109cc6109d2923690600401610231565b916119f0565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a7f60209160243560443591600401611ac5565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610af6600435610af181610a87565b6123ab565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c557610b5e903690600401610231565b9060243567ffffffffffffffff81116101c557610b82610b8c913690600401610231565b9190935a936124e0565b91610ba060608401516080850151906127a4565b610bab8285856121c6565b929015610bbd575061001c93506128fd565b836108e386926040519384937fa2b6d61b000000000000000000000000000000000000000000000000000000008552600485016119cb565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610c84576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610ce481610a87565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d09826123ab565b1615610d6f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d60600082613124565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107e691600401610231565b610e7736610e25565b9030330361036a57610e8d61001c925a926124e0565b90610e9782612bf0565b906128fd565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610ed46101a2565b5060443567ffffffffffffffff81116101c557610ef5903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a7f600435612c84565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610f99600435612cca565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101c557610fc936610e25565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156110a957330361107b57303b156101c55761104f9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611c66565b038183305af180156110765761106157005b80611070600061001c93610404565b80610dbf565b611a3c565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b919082519283825260005b84811061114e5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161110f565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576111b56111bb9136906004016110d3565b90611d94565b60405160208101916020825280518093526040820192602060408260051b85010192019060005b8181106111ef5784840385f35b9091927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08582030186528351908151600681101561126157600192826020939284935260408061124c858501516060878601526060850190611104565b930151910152950196019101949190946111e2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061136c60008061136636610692565b91612dd1565b92611378839293613edb565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356113d281610a87565b6113da6101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff611400826123ab565b1661147a577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000060409316911661146d8183613124565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115016101a2565b5061150a6101ca565b5060443567ffffffffffffffff81116101c55761152b9036906004016110d3565b505060643567ffffffffffffffff81116101c55761154d9036906004016110d3565b505060843567ffffffffffffffff81116101c55761156f903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115d36101a2565b506115dc6101ca565b5060843567ffffffffffffffff81116101c5576115fd903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356116636101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff836116ee7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b1617856130df565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161173f57565b611702565b9190820180921161173f57565b9061175b82610481565b6117686040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06117968294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117e35760209160051b010190565b6117a0565b9080602083519182815201916020808360051b8301019401926000915b83831061181457505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611887604085015160e0604086015260e0850190611104565b936060810151606085015260808101511515608085015260a0810151151560a0850152015191015297019301930191939290611805565b906020808351928381520192019060005b8181106118dc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016118cf565b805160ff16825261052591602082810151151590820152610100611966611940604085015161012060408601526101208501906117e8565b606085015160608501526080850151608085015260a085015184820360a0860152611104565b9260c081015160c084015260e081015160e08401520151906101008184039101526118be565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916119e26105259492604085526040850190611908565b92602081850391015261198c565b90611a0d92916119fe612350565b906003825260e08201526121c6565b5015611a37577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b908160209103126101c5575161052581610a87565b60409061052594928152816020820152019161198c565b91909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169182156110a95782330361107b5780611bf7575b506020915080610100611b27920190611a48565b92611b5f60405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611aae565b0381305afa908115611076577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611bc8575b501603611bc357600090565b600190565b611bea915060203d602011611bf0575b611be28183610404565b810190611a99565b38611bb7565b503d611bd8565b823b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af190811561107657602092611b2792611c51575b5090611b13565b806110706000611c6093610404565b38611c4a565b91602061052593818152019161198c565b90611c8182610481565b611c8e6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611cbc8294610481565b0160005b818110611ccc57505050565b60405190606082019180831067ffffffffffffffff8411176103e357602092604052600081526060838201526000604082015282828601015201611cc0565b91908110156117e35760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21813603018212156101c5570190565b610525903690610528565b9190820391821161173f57565b909260c092610525959460008452602084015260408301526060820152600060808201528160a08201520190611104565b90915a92600090611da481611c77565b94825b828110611db7575b505050509050565b611dca611dc5828589611d0b565b611d4b565b93611dd860a0860151151590565b806120bc575b6120b2575060009360608101518015908115806120a9575b6120385783611e086080850151151590565b15611fd957611eb390611ebb92935a94611e36875173ffffffffffffffffffffffffffffffffffffffff1690565b9115611fd357505a905b611eae89611e8260408a01518d6040519788947f4c4e814c00000000000000000000000000000000000000000000000000000000602087015260248601611d63565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612d36565b915a90611d56565b6040611ec7858c6117cf565b5101525b15611f07575b5080611ee9611ee26001938a6117cf565b5160019052565b611ef1612d48565b6020611efd838b6117cf565b5101525b01611da7565b60c001805115611f9c576001815114611f635751600214611f285738611ed1565b939450505050611f42611f3b82856117cf565b5160039052565b6020611f56611f4f612d48565b92856117cf565b5101528038808080611daf565b509395809550611f829350611f7b92508491506117cf565b5160049052565b6020611f96611f8f612d48565b92846117cf565b51015290565b50925060018093611fb7611fb0828a6117cf565b5160029052565b611fbf612d48565b6020611fcb838b6117cf565b510152611f01565b90611e40565b50611eb361201e915a93612001865173ffffffffffffffffffffffffffffffffffffffff1690565b9160208701519160001461203257505a905b604087015192612d24565b604061202a858c6117cf565b510152611ecb565b90612013565b505050905085945060209250611f96915061206161205a8261207796986117cf565b5160059052565b6120a35a60405185810191825295869160200190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101865285610404565b846117cf565b50805a10611df6565b9350600190611f01565b508015611dde565b3d156120ef573d906120d582610499565b916120e36040519384610404565b82523d6000602084013e565b606090565b6004361080156121015750565b612137906000357fffffffff00000000000000000000000000000000000000000000000000000000811691612189575b506123ab565b73ffffffffffffffffffffffffffffffffffffffff81166121555750565b60008091604051368382378036810184815203915af46121736120c4565b901561218157602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638612131565b90156117e35790565b91907f80000000000000000000000000000000000000000000000000000000000000008061221d6121f785856121bd565b357fff000000000000000000000000000000000000000000000000000000000000001690565b161461227c57600091829161223194612dd1565b90509190919280821061224c57505061224990613edb565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b505061228790612bf0565b9061229182612cca565b4281111561231e575073ffffffffffffffffffffffffffffffffffffffff81168015159081612313575b506122c7575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b9050331415386122bb565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152612422606082610404565b519020541690565b9061243482610481565b6124416040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061246f8294610481565b019060005b82811061248057505050565b60209060405161248f816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201612474565b909392938483116101c55784116101c5578101920390565b9190916124eb612350565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183160161278d5750600060608701525b600761253a60ff831660011c90565b168061273c575b5060108181160361270757506001925b61255a8461242a565b92604087019384526000905b85821061257557505050505050565b8281013560f81c906001019091816001808516036126e557506125b93061259d8389516117cf565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808416146126c5575b600480841614612677575b600880841614612644575b61262d61262760c0856126068a60806125fd8860108060019c9d161493516117cf565b51019015159052565b61261d8a60a06125fd8860208087161493516117cf565b1660061c60031690565b60ff1690565b60c061263a8389516117cf565b5101520190612566565b60019161262d906126279060c090878101359060200196906060612669878d516117cf565b5101529594505050506125da565b906126bf908481013560e81c906003016126b861269f6126978484611744565b838a8a6124c8565b919060406126ae888d516117cf565b51019236916104d3565b9052611744565b906125cf565b908381013590602001919060206126dd8389516117cf565b5101526125c4565b61270292508481013560601c90601401929061259d8389516117cf565b6125b9565b6020908116036127295761ffff918381013560f01c906002015b921692612551565b60ff918381013560f81c90600101612721565b612780919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608087015238612541565b80850135606090811c90880152601401925061252b565b906127ae82612c84565b81810361283457509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e82528360408201526040815261281e606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61288060409295949395606083526060830190611908565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611104565b6105259392606092825260208201528160408201520190611104565b6128ea6105259492606083526060830190611908565b9260208201526040818403910152611104565b916000604082019384515190825b82811061291c575b50505050505050565b6129278188516117cf565b519361293660a0860151151590565b80612be8575b612ba8575060009360608101518015801580612b9f575b612b675784906129666080850151151590565b15612b21576129e09261298d855173ffffffffffffffffffffffffffffffffffffffff1690565b9115612b1b57505a905b611eae8b611e8260608d01516040890151908c8b604051998a967f4c4e814c0000000000000000000000000000000000000000000000000000000060208901526024880161288a565b15612a23575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b0161290b565b60c001805115612ad2576001815114612a935751600214612a4457386129e6565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250612a84612a78612d48565b604051938493846128b8565b0390a138808080808080612913565b50846108e3612aa0612d48565b6040519384937f7f6b0bb1000000000000000000000000000000000000000000000000000000008552600485016128d4565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612b13612b05612d48565b604051918291858c846128b8565b0390a1612a1d565b90612997565b8351612b5c93925073ffffffffffffffffffffffffffffffffffffffff1691602085015191600014612b6157505a905b604085015192612d24565b6129e0565b90612b51565b83886108e35a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612868565b50815a10612953565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612b13565b50801561293c565b612c52612c7e612c10612c0a602085015115153090613163565b9361325e565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612cc3606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612d09606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405190612d70826103e8565b60006020838281520152565b908160409103126101c557602060405191612d96836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611104565b909491939291853560f81c60019093819087612deb612d63565b93604080891614806130c1575b612fc4575b505050600180861614612f9e5760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612e6d9060051c90565b612e7690611731565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612ec284612bf0565b988993612ece936124c8565b91612ed893613675565b9098612eec91600052602052604060002090565b90612eff91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612f2891600052602052604060002090565b94815190868215159283612f93575b505081612f84575b50612f475750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612f3f565b141591508638612f37565b86612fb99792949550612fb193986124c8565b939092613453565b919394909293929190565b8a81013560601c985060140192509087908a90849015612fe8575b50829150612dfd565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c9260409261306b9290916130369161302f9189906130298983611744565b926124c8565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612da4565b0392165afa80156110765761308992600091613092575b5092611744565b86388981612fdf565b6130b4915060403d6040116130ba575b6130ac8183610404565b810190612d7c565b38613082565b503d6130a2565b5073ffffffffffffffffffffffffffffffffffffffff891615612df8565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261311e606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120835260408201526040815261311e606082610404565b15613213576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a08152612c7e60c082610404565b469061316c565b805160209091019060005b8181106132325750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613225565b61010081015160405161327981612c5260208201809561321a565b51902090613288815160ff1690565b60ff81168061330157505090612c7e6132a46040840151613f3e565b92612c5260806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b6001810361335f57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612c7e8160808101612c52565b600281036133b557505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612c7e8160808101612c52565b600303613409575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612c7e8160808101612c52565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b90612880906040939695949660608452606084019161198c565b919492909260009560009560009560009560009561346f612350565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b8281106134c6575050505050505080511515806134b8575b612f475750565b5060208101518411156134b1565b600381019d50959b50939950919750929091906134ea908b9085013560e81c611744565b958a600084890361362e575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c036135fe57509060016135338961353a9487896124c8565b908b612dd1565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106135b85750928b8851146135af575b808b101561357d57508a60c08501528992959295949194939093613499565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b6000885261355e565b8d8f6108e36135c985858c8e6124c8565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601613439565b9798999a9160016136148b61361b94888a6124c8565b9086612dd1565b50929d919c909b929a9092918e8e61354e565b916134f9565b73ffffffffffffffffffffffffffffffffffffffff61052595936060938352166020820152816040820152019161198c565b908160209103126101c5575190565b9391909360009460009460005b818110613690575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613e74575060018914613e345760028914613c775760038914613c485760048914613bc75760068914613b275760058914613ad95760078914613a1257600889146139bc576009891461389357600a8914613730577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613882575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816137a89187876124c8565b6040517f898bd921000000000000000000000000000000000000000000000000000000008152939184916137e0918a60048501611aae565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611076576138239360009361384f575b5060ff909a168091019a61420a565b908015613849579061383d91600052602052604060002090565b955b9392919093613682565b5061383d565b60ff9193506138749060203d811161387b575b61386c8183610404565b810190613666565b9290613814565b503d613862565b8084013560f81c9850600101613745565b909192939495969750600389169788156139ab575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161390b9187876124c8565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613943918b600485016119cb565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611076576139859360009361384f575060ff909a168091019a61420a565b9080156139a5579061399f91600052602052604060002090565b9561383f565b5061399f565b8084013560f81c98506001016138a8565b9850602087019750949593949293919290918201356139da866141b1565b81146139ea575b613985906141cb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506139e1565b975090919293949597600f16968715613ac7575b60206000613a38613aa59a9b86614079565b9c9092918a604051613a7b81612c528a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15611076576139859060ff6000519a16809101996140bc565b600189019883013560f81c9750613a26565b985060208701975094959394929391929091820135808514613aff575b61398590614172565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613af6565b989091929394959662ffffff9850613b49612627600c8416603f9060021c1690565b918215613bb3575b6003168015613ba2575b908190613b8690613b7e908781013560e81c906003019c168c01809c89896124c8565b90898b613675565b911115613b99575b906139859291614127565b99820199613b8e565b50600281019084013560f01c613b5b565b8482013560f81c9250600190910190613b51565b975097613c1d613c2a929394959697600f613c3293169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92908301809386866124c8565b908688613675565b9061399f92980198600052602052604060002090565b9850965093949293919290919080820135906020019680156139a5579061399f91600052602052604060002090565b90919293949596975060038916978815613e23575b8084013560601c99613ceb9160140190613cab9060021c600316612627565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613d4f60208c613d0185858b8b6124c8565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611aae565b0392165afa908115611076577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613e05575b501603613dc157509060ff613985929916809101996140bc565b6108e3613dd28c93899389896124c8565b906040519485947fb2fed7ae00000000000000000000000000000000000000000000000000000000865260048601613634565b613e1d915060203d8111611bf057611be28183610404565b38613da7565b8381013560f81c9850600101613c8c565b98600f91929394959697985016968715613e63575b60148101976139859160ff9091169084013560601c6140bc565b8281013560f81c9750600101613e49565b98509091929394959698600f16978815613ec6575b5060206000613e9c613aa59a9b86614079565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613e89565b8015159081613ee8575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b818110613f285750505090565b8251845260209384019390920191600101613f1b565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613f84613f6e83610481565b92613f7c6040519485610404565b808452610481565b0136602083013760005b83518110156140605780613fa4600192866117cf565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e0830152610100820152610100815261404c61012082610404565b51902061405982856117cf565b5201613f8e565b50909150604051612c7e81612c52602082018095613f10565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff811161173f5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612c7e606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612c7e609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612c7e605882610404565b612c52612c7e612c10612c0a600060208601511515613163565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612c7e606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612c7e608d8261040456fea2646970667358221220663e0e19a1f44876322f0807bffe0131225f2bf0bc3ca63984e33df77f73e01664736f6c634300081c0033"
