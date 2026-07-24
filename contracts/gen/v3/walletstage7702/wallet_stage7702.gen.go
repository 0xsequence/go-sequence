// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletstage7702

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

// WalletStage7702MetaData contains all meta data concerning the WalletStage7702 contract.
var WalletStage7702MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_defaultCheckpointer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_CHECKPOINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainedSignatureNestedInChainedSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidPackedLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c03461008e57601f6140e038819003918201601f19168301916001600160401b0383118484101761009357808492604094855283398101031261008e57610052602061004b836100a9565b92016100a9565b60805260a05260405161402290816100be8239608051818181610e070152611d5c015260a0518181816110340152818161119b0152611a4a0152f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b038216820361008e5756fe6080604052600436101561001e575b361561001c5761001c611dc8565b005b60003560e01c806223de291461018d57806313792a4a14610188578063150b7a02146101835780631626ba7e1461017e57806319822f7c146101795780631a9b2337146101745780631f6a1eb91461016f578063295614261461016a5780634fcf3eca1461016557806351605d8014610160578063540e3ad91461015b5780636ea44577146101565780638943ec02146101515780638c3f55631461014c57806392dcb3fc146101475780639c145aed14610142578063a65d69d41461013d578063aaf10f4214610138578063ad55366b14610133578063b93ea7ad1461012e578063bc197c8114610129578063f23a6e61146101245763f727ef1c0361000e57611562565b6114d5565b611403565b6112a4565b611258565b6111bf565b611150565b610fc1565b610f63565b610f27565b610ea3565b610e74565b610dbc565b610d83565b610c67565b610b80565b610a6f565b610a0c565b610957565b6108cf565b610842565b610745565b61024f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b9181601f840112156101b55782359167ffffffffffffffff83116101b557602083818601950101116101b557565b346101b55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610286610192565b5061028f6101ba565b506102986101dd565b5060843567ffffffffffffffff81116101b5576102b9903690600401610221565b505060a43567ffffffffffffffff81116101b55761001c903690600401610221565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff82111761032657604052565b6102db565b6040810190811067ffffffffffffffff82111761032657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761032657604052565b6040519061039760e083610347565b565b6040519061039761012083610347565b359060ff821682036101b557565b359081151582036101b557565b67ffffffffffffffff81116103265760051b60200190565b67ffffffffffffffff811161032657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b929192610422826103dc565b916104306040519384610347565b8294818452818301116101b5578281602093846000960137010152565b9080601f830112156101b55781602061046893359101610416565b90565b81601f820112156101b557803590610482826103c4565b926104906040519485610347565b82845260208085019360051b830101918183116101b55760208101935b8385106104bc57505050505090565b843567ffffffffffffffff81116101b557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101b557610503610388565b9161051060208301610200565b83526040820135602084015260608201359267ffffffffffffffff84116101b55760e08361054588602080988198010161044d565b60408401526080810135606084015261056060a082016103b7565b608084015261057160c082016103b7565b60a0840152013560c08201528152019401936104ad565b9080601f830112156101b557813561059f816103c4565b926105ad6040519485610347565b81845260208085019260051b8201019283116101b557602001905b8282106105d55750505090565b602080916105e284610200565b8152019101906105c8565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101b55761065f610399565b9061066c816004016103a9565b825261067a602482016103b7565b6020830152604481013567ffffffffffffffff81116101b5578460046106a29284010161046b565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101b5578460046106de9284010161044d565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101b5576004859161071b930101610588565b610100820152916024359067ffffffffffffffff82116101b55761074191600401610221565b9091565b346101b557610753366105ed565b909161010081019261076e61076985515161166a565b61168a565b9160005b855180518210156107d557906107cf6107aa61079083600195611708565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6107b48388611708565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610772565b50508383866107ea336107b483515185611708565b526107f6818484611e9a565b50156108085760405160018152602090f35b61083e906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611963565b0390fd5b346101b55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610879610192565b506108826101ba565b5060643567ffffffffffffffff81116101b5576108a3903690600401610221565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101b55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043560243567ffffffffffffffff81116101b55760209161092761092d923690600401610221565b91611988565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b5576109da60209160243560443591600401611a31565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101b557565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020610a51600435610a4c816109e2565b6120aa565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b557610ab9903690600401610221565b60243567ffffffffffffffff81116101b557610ad9903690600401610221565b9160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610b5657610b309360027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55611b9b565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557600435303303610c39578015610c0f576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557600435610c9d816109e2565b303303610c395773ffffffffffffffffffffffffffffffffffffffff610cc2826120aa565b1615610d285760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d19600082612eef565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101b557565b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760206109da611c91565b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101b5576004359067ffffffffffffffff82116101b55761074191600401610221565b610e7d36610e2b565b90303303610c3957610e9361001c925a926121df565b90610e9d826128f4565b9061255e565b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610eda610192565b5060443567ffffffffffffffff81116101b557610efb903690600401610221565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760206109da600435612982565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576040610f9f6004356129c8565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101b557610fcf36610e2b565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610b565760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156111265733036110f857303b156101b5576110a29160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611d87565b038183305af180156110f3576110d85760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b806110e760006110ed93610347565b80610d78565b38610b30565b6119d4565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557602060006017303b146112195773ffffffffffffffffffffffffffffffffffffffff60405191168152f35b6017600080303c60005162ef01008160e81c14611237575b50610a51565b73ffffffffffffffffffffffffffffffffffffffff915060481c1638611231565b346101b55760c061127460008061126e366105ed565b91612ac9565b92909190611280611c91565b821490604051958652602086015260408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576004356112da816109e2565b6112e26101ba565b90303303610c395773ffffffffffffffffffffffffffffffffffffffff611308826120aa565b16611382577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166113758183612eef565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101b55782359167ffffffffffffffff83116101b5576020808501948460051b0101116101b557565b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761143a610192565b506114436101ba565b5060443567ffffffffffffffff81116101b5576114649036906004016113d2565b505060643567ffffffffffffffff81116101b5576114869036906004016113d2565b505060843567ffffffffffffffff81116101b5576114a8903690600401610221565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761150c610192565b506115156101ba565b5060843567ffffffffffffffff81116101b557611536903690600401610221565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043561159c6101ba565b604435916bffffffffffffffffffffffff83168093036101b557303303610c39578273ffffffffffffffffffffffffffffffffffffffff836116277febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612e6b565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161167857565b61163b565b9190820180921161167857565b90611694826103c4565b6116a16040519182610347565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116cf82946103c4565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561171c5760209160051b010190565b6116d9565b919082519283825260005b84811061176b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161172c565b9080602083519182815201916020808360051b8301019401926000915b8383106117ac57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c08061181f604085015160e0604086015260e0850190611721565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061179d565b906020808351928381520192019060005b8181106118745750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611867565b805160ff168252610468916020828101511515908201526101006118fe6118d860408501516101206040860152610120850190611780565b606085015160608501526080850151608085015260a085015184820360a0860152611721565b9260c081015160c084015260e081015160e0840152015190610100818403910152611856565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9161197a61046894926040855260408501906118a0565b926020818503910152611924565b906119a5929161199661204f565b906003825260e0820152611e9a565b50156119cf577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101b5570180359067ffffffffffffffff82116101b5576020019181360383136101b557565b909173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016908115611126578133036110f85780611ae5575b5050611ad57f1626ba7e0000000000000000000000000000000000000000000000000000000092610927836101007fffffffff000000000000000000000000000000000000000000000000000000009501906119e0565b1603611ae057600090565b600190565b813b156101b5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af19283156110f3576109277fffffffff0000000000000000000000000000000000000000000000000000000093611ad5937f1626ba7e0000000000000000000000000000000000000000000000000000000096611b86575b5093505092611a7e565b806110e76000611b9593610347565b38611b7c565b91939290611baa905a936121df565b9160608301516080840151611bbe82612982565b818103611c5d57509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881920190611bf78282612eb0565b604080519182526020820192909252a1611c12828685611e9a565b929015611c25575061039793945061255e565b8361083e87926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611963565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85480610468575060405160208101907f53657175656e6365207369676e65723a0a00000000000000000000000000000082527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000003060601b1660318201526001604582015260458152611d25606582610347565b519020600052600160205260406000206000526000602052604060002060005273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016602052604060002090565b916020610468938181520191611924565b3d15611dc3573d90611da9826103dc565b91611db76040519384610347565b82523d6000602084013e565b606090565b600436108015611dd55750565b611e0b906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611e5d575b506120aa565b73ffffffffffffffffffffffffffffffffffffffff8116611e295750565b60008091604051368382378036810184815203915af4611e47611d98565b9015611e5557602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611e05565b901561171c5790565b611ecd611ea78484611e91565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f800000000000000000000000000000000000000000000000000000000000000080821614611f505750611f05926000928392612ac9565b905091909192808210611f20575050611f1d90613ce1565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0200000000000000000000000000000000000000000000000000000000000000908116146020820152611f86925090506128f4565b90611f90826129c8565b4281111561201d575073ffffffffffffffffffffffffffffffffffffffff81168015159081612012575b50611fc6575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611fba565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff821117610326576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152612121606082610347565b519020541690565b90612133826103c4565b6121406040519182610347565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061216e82946103c4565b019060005b82811061217f57505050565b60209060405161218e8161030a565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201612173565b909392938483116101b55784116101b5578101920390565b906121e861204f565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016124b75750600060608601525b600761223760ff831660011c90565b1680612466575b50601081811603612438575060015b61225681612129565b604086019081526000925b82841061229e5750505050036122745790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361241657506122e8306122cc848751611708565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808216146123f6575b6004808216146123a8575b600880821614612373575b9061235a61235460c0846123346010806001981614608061232b888c51611708565b51019015159052565b61234a60208083161460a061232b888c51611708565b1660061c60031690565b60ff1690565b60c0612367838751611708565b51015201929190612261565b946001919061235a906123549060c09086810135906020019990606061239a878b51611708565b510152939450505050612309565b946123f0908381013560e81c906003016123e96123d06123c8848461167d565b838c896121c7565b919060406123df888b51611708565b5101923691610416565b905261167d565b946122fe565b9482810135906020019590602061240e848751611708565b5101526122f3565b61243396508381013560601c9060140196906122cc848751611708565b6122e8565b60209081160361245557600282019181013560f01c905b9061224d565b600182019181013560f81c9061244f565b6124aa919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b929060808601523861223e565b80830135606090811c908701526014019250612228565b6124e6604092959493956060835260608301906118a0565b9460208201520152565b9261046896959260c09592855260208501526040840152606083015260808201528160a08201520190611721565b604090610468939281528160208201520190611721565b61254b61046894926060835260608301906118a0565b9260208201526040818403910152611721565b916000604082019384515190825b82811061257d575b50505050505050565b612588818851611708565b519361259760a0860151151590565b8061287b575b612841575060009360608101518015801580612838575b6128005784906125c76080850151151590565b156127ba57612672926125ee855173ffffffffffffffffffffffffffffffffffffffff1690565b91156127b457505a905b61266d8b61264160608d01516040890151908c8b604051998a967f4c4e814c000000000000000000000000000000000000000000000000000000006020890152602488016124f0565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610347565b612e3e565b156126bb575b50600190867f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a604051806126b185829190602083019252565b0390a25b0161256c565b60c00180511561276b57600181511461272c57516002146126dc5738612678565b93505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b91925061270d612e50565b9061271d6040519283928361251e565b0390a238808080808080612574565b508461083e612739612e50565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612535565b50925060018093867f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6127ac61279f612e50565b604051918291868361251e565b0390a26126b5565b906125f8565b83516127f593925073ffffffffffffffffffffffffffffffffffffffff16916020850151916000146127fa57505a905b604085015192612e2c565b612672565b906127ea565b838861083e5a6040519384937f21395274000000000000000000000000000000000000000000000000000000008552600485016124ce565b50815a106125b4565b9350600190867f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b604051806127ac85829190602083019252565b50801561259d565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526128ee606582610347565b51902090565b6129566128ee61291461290e602085015115153090612f2e565b93613029565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610347565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83526040820152604081526129c1606082610347565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612a07606082610347565b51902054906bffffffffffffffffffffffff8260601c921690565b60405190612a2f8261032b565b60006020838281520152565b60031115612a4557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160409103126101b557602060405191612a8e8361032b565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61046894931681528160208201520190611721565b909491939291853560f81c600190938190612ae2612a22565b92612aec82612a3b565b60018203612cf6575b50600180871614612c95575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612b649060051c90565b612b6d9061166a565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612bb9846128f4565b988993612bc5936121c7565b91612bcf9361346b565b9098612be391600052602052604060002090565b90612bf691600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612c1f91600052602052604060002090565b94815190868215159283612c8a575b505081612c7b575b50612c3e5750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612c36565b141591508638612c2e565b909691939450612ca781989398612a3b565b612ccc57612cc19581612cb9936121c7565b93909261321e565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612af5578981013560601c9750601401915087878a84612d3f85612a3b565b60028503612d50575b505050612af5565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c92604092612dd5929091612da091612d99918a90612d93898361167d565b926121c7565b3691610416565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612a9c565b0392165afa80156110f357612df392600091612dfd575b509361167d565b9087388a81612d48565b612e1f915060403d604011612e25575b612e178183610347565b810190612a74565b38612dec565b503d612e0d565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612eaa606082610347565b51902055565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612eaa606082610347565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612eaa606082610347565b15612fde576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a081526128ee60c082610347565b4690612f37565b805160209091019060005b818110612ffd5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612ff0565b61010081015160405161304481612956602082018095612fe5565b51902090613053815160ff1690565b60ff8116806130cc575050906128ee61306f6040840151613d1b565b9261295660806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b6001810361312a57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466938101938452908101919091526060810192909252906128ee8160808101612956565b6002810361318057505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082019081529181019290925260608201929092526128ee8160808101612956565b6003036131d4575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082019081529181019290925260608201929092526128ee8160808101612956565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906124e69060409396959496606084526060840191611924565b919492909260009560009560009560009560009561323a61204f565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061329157505050505050508051151580613283575b612c3e5750565b50602081015184111561327c565b600381019d50959b50939950919750929091906132b5908b9085013560e81c61167d565b958287036133f5578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c036133c65750613302916132fb898c9387896121c7565b908b612ac9565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106133805750928b885114613377575b808b101561334557508a60c08501528992959295949194939093613264565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b60008852613326565b8d8f61083e61339185858c8e6121c7565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601613204565b979899809b926133db8b6133e294888a6121c7565b9086612ac9565b50929d919c909b929a9092918e8e613316565b8a6002916132c2565b908160209103126101b55751610468816109e2565b604090610468949281528160208201520191611924565b73ffffffffffffffffffffffffffffffffffffffff610468959360609383521660208201528160408201520191611924565b908160209103126101b5575190565b9391909360009460009460005b818110613486575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613c7a575060018914613c3a5760028914613a6d5760038914613a3e57600489146139bd576006891461391d57600589146138cf576007891461380857600889146137b2576009891461368957600a8914613526577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613678575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161359e9187876121c7565b6040517f898bd921000000000000000000000000000000000000000000000000000000008152939184916135d6918a60048501613413565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156110f35761361993600093613645575b5060ff909a168091019a613f7c565b90801561363f579061363391600052602052604060002090565b955b9392919093613478565b50613633565b60ff91935061366a9060203d8111613671575b6136628183610347565b81019061345c565b929061360a565b503d613658565b8084013560f81c985060010161353b565b909192939495969750600389169788156137a1575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816137019187876121c7565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613739918b60048501611963565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156110f35761377b93600093613645575060ff909a168091019a613f7c565b90801561379b579061379591600052602052604060002090565b95613635565b50613795565b8084013560f81c985060010161369e565b9850602087019750949593949293919290918201356137d086613f23565b81146137e0575b61377b90613f3d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506137d7565b975090919293949597600f169687156138bd575b6020600061382e61389b9a9b86613e56565b9c9092918a604051613871816129568a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa156110f35761377b9060ff6000519a1680910199612883565b600189019883013560f81c975061381c565b9850602087019750949593949293919290918201358085146138f5575b61377b90613ee4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506138ec565b989091929394959662ffffff985061393f612354600c8416603f9060021c1690565b9182156139a9575b6003168015613998575b90819061397c90613974908781013560e81c906003019c168c01809c89896121c7565b90898b61346b565b91111561398f575b9061377b9291613e99565b99820199613984565b50600281019084013560f01c613951565b8482013560f81c9250600190910190613947565b975097613a13613a20929394959697600f613a2893169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92908301809386866121c7565b90868861346b565b9061379592980198600052602052604060002090565b98509650939492939192909190808201359060200196801561379b579061379591600052602052604060002090565b90919293949596975060038916978815613c29575b8084013560601c99613ae19160140190613aa19060021c600316612354565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613b4560208c613af785858b8b6121c7565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501613413565b0392165afa9081156110f3577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613bfb575b501603613bb757509060ff61377b92991680910199612883565b61083e613bc88c93899389896121c7565b906040519485947fb2fed7ae0000000000000000000000000000000000000000000000000000000086526004860161342a565b613c1c915060203d8111613c22575b613c148183610347565b8101906133fe565b38613b9d565b503d613c0a565b8381013560f81c9850600101613a82565b98600f91929394959697985016968715613c69575b601481019761377b9160ff9091169084013560601c612883565b8281013560f81c9750600101613c4f565b98509091929394959698600f16978815613ccc575b5060206000613ca261389b9a9b86613e56565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613c8f565b613ce9611c91565b1490565b805160209091019060005b818110613d055750505090565b8251845260209384019390920191600101613cf8565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613d61613d4b836103c4565b92613d596040519485610347565b8084526103c4565b0136602083013760005b8351811015613e3d5780613d8160019286611708565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613e2961012082610347565b519020613e368285611708565b5201613d6b565b509091506040516128ee81612956602082018095613ced565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116116785791565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526128ee609882610347565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526128ee605882610347565b6129566128ee61291461290e600060208601511515612f2e565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526128ee606082610347565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d81526128ee608d8261034756fea264697066735822122070d2ccb5d2e8d1051d4fbeee4718ed4044611e6d09516015fc4e660973e187a464736f6c634300081c0033",
}

// WalletStage7702ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage7702MetaData.ABI instead.
var WalletStage7702ABI = WalletStage7702MetaData.ABI

// WalletStage7702Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage7702MetaData.Bin instead.
var WalletStage7702Bin = WalletStage7702MetaData.Bin

// DeployWalletStage7702 deploys a new Ethereum contract, binding an instance of WalletStage7702 to it.
func DeployWalletStage7702(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address, _defaultCheckpointer common.Address) (common.Address, *types.Transaction, *WalletStage7702, error) {
	parsed, err := WalletStage7702MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage7702Bin), backend, _entryPoint, _defaultCheckpointer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletStage7702{WalletStage7702Caller: WalletStage7702Caller{contract: contract}, WalletStage7702Transactor: WalletStage7702Transactor{contract: contract}, WalletStage7702Filterer: WalletStage7702Filterer{contract: contract}}, nil
}

// WalletStage7702 is an auto generated Go binding around an Ethereum contract.
type WalletStage7702 struct {
	WalletStage7702Caller     // Read-only binding to the contract
	WalletStage7702Transactor // Write-only binding to the contract
	WalletStage7702Filterer   // Log filterer for contract events
}

// WalletStage7702Caller is an auto generated read-only Go binding around an Ethereum contract.
type WalletStage7702Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage7702Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletStage7702Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage7702Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletStage7702Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage7702Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletStage7702Session struct {
	Contract     *WalletStage7702  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletStage7702CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletStage7702CallerSession struct {
	Contract *WalletStage7702Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WalletStage7702TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletStage7702TransactorSession struct {
	Contract     *WalletStage7702Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WalletStage7702Raw is an auto generated low-level Go binding around an Ethereum contract.
type WalletStage7702Raw struct {
	Contract *WalletStage7702 // Generic contract binding to access the raw methods on
}

// WalletStage7702CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletStage7702CallerRaw struct {
	Contract *WalletStage7702Caller // Generic read-only contract binding to access the raw methods on
}

// WalletStage7702TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletStage7702TransactorRaw struct {
	Contract *WalletStage7702Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletStage7702 creates a new instance of WalletStage7702, bound to a specific deployed contract.
func NewWalletStage7702(address common.Address, backend bind.ContractBackend) (*WalletStage7702, error) {
	contract, err := bindWalletStage7702(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702{WalletStage7702Caller: WalletStage7702Caller{contract: contract}, WalletStage7702Transactor: WalletStage7702Transactor{contract: contract}, WalletStage7702Filterer: WalletStage7702Filterer{contract: contract}}, nil
}

// NewWalletStage7702Caller creates a new read-only instance of WalletStage7702, bound to a specific deployed contract.
func NewWalletStage7702Caller(address common.Address, caller bind.ContractCaller) (*WalletStage7702Caller, error) {
	contract, err := bindWalletStage7702(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702Caller{contract: contract}, nil
}

// NewWalletStage7702Transactor creates a new write-only instance of WalletStage7702, bound to a specific deployed contract.
func NewWalletStage7702Transactor(address common.Address, transactor bind.ContractTransactor) (*WalletStage7702Transactor, error) {
	contract, err := bindWalletStage7702(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702Transactor{contract: contract}, nil
}

// NewWalletStage7702Filterer creates a new log filterer instance of WalletStage7702, bound to a specific deployed contract.
func NewWalletStage7702Filterer(address common.Address, filterer bind.ContractFilterer) (*WalletStage7702Filterer, error) {
	contract, err := bindWalletStage7702(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702Filterer{contract: contract}, nil
}

// bindWalletStage7702 binds a generic wrapper to an already deployed contract.
func bindWalletStage7702(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletStage7702MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage7702 *WalletStage7702Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage7702.Contract.WalletStage7702Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage7702 *WalletStage7702Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage7702.Contract.WalletStage7702Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage7702 *WalletStage7702Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage7702.Contract.WalletStage7702Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage7702 *WalletStage7702CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage7702.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage7702 *WalletStage7702TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage7702.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage7702 *WalletStage7702TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage7702.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTCHECKPOINTER is a free data retrieval call binding the contract method 0x540e3ad9.
//
// Solidity: function DEFAULT_CHECKPOINTER() view returns(address)
func (_WalletStage7702 *WalletStage7702Caller) DEFAULTCHECKPOINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "DEFAULT_CHECKPOINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEFAULTCHECKPOINTER is a free data retrieval call binding the contract method 0x540e3ad9.
//
// Solidity: function DEFAULT_CHECKPOINTER() view returns(address)
func (_WalletStage7702 *WalletStage7702Session) DEFAULTCHECKPOINTER() (common.Address, error) {
	return _WalletStage7702.Contract.DEFAULTCHECKPOINTER(&_WalletStage7702.CallOpts)
}

// DEFAULTCHECKPOINTER is a free data retrieval call binding the contract method 0x540e3ad9.
//
// Solidity: function DEFAULT_CHECKPOINTER() view returns(address)
func (_WalletStage7702 *WalletStage7702CallerSession) DEFAULTCHECKPOINTER() (common.Address, error) {
	return _WalletStage7702.Contract.DEFAULTCHECKPOINTER(&_WalletStage7702.CallOpts)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage7702 *WalletStage7702Caller) Entrypoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "entrypoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage7702 *WalletStage7702Session) Entrypoint() (common.Address, error) {
	return _WalletStage7702.Contract.Entrypoint(&_WalletStage7702.CallOpts)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage7702 *WalletStage7702CallerSession) Entrypoint() (common.Address, error) {
	return _WalletStage7702.Contract.Entrypoint(&_WalletStage7702.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage7702 *WalletStage7702Caller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage7702 *WalletStage7702Session) GetImplementation() (common.Address, error) {
	return _WalletStage7702.Contract.GetImplementation(&_WalletStage7702.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage7702 *WalletStage7702CallerSession) GetImplementation() (common.Address, error) {
	return _WalletStage7702.Contract.GetImplementation(&_WalletStage7702.CallOpts)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage7702 *WalletStage7702Caller) GetStaticSignature(opts *bind.CallOpts, _hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "getStaticSignature", _hash)

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
func (_WalletStage7702 *WalletStage7702Session) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage7702.Contract.GetStaticSignature(&_WalletStage7702.CallOpts, _hash)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage7702 *WalletStage7702CallerSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage7702.Contract.GetStaticSignature(&_WalletStage7702.CallOpts, _hash)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage7702 *WalletStage7702Caller) ImageHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage7702 *WalletStage7702Session) ImageHash() ([32]byte, error) {
	return _WalletStage7702.Contract.ImageHash(&_WalletStage7702.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage7702 *WalletStage7702CallerSession) ImageHash() ([32]byte, error) {
	return _WalletStage7702.Contract.ImageHash(&_WalletStage7702.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage7702 *WalletStage7702Caller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage7702 *WalletStage7702Session) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.IsValidSignature(&_WalletStage7702.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage7702 *WalletStage7702CallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.IsValidSignature(&_WalletStage7702.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Caller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC1155BatchReceived(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702CallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC1155BatchReceived(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Caller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC1155Received(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702CallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC1155Received(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC721Received(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.OnERC721Received(&_WalletStage7702.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage7702 *WalletStage7702Caller) ReadHook(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "readHook", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage7702 *WalletStage7702Session) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletStage7702.Contract.ReadHook(&_WalletStage7702.CallOpts, selector)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage7702 *WalletStage7702CallerSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletStage7702.Contract.ReadHook(&_WalletStage7702.CallOpts, selector)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage7702 *WalletStage7702Caller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage7702 *WalletStage7702Session) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage7702.Contract.ReadNonce(&_WalletStage7702.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage7702 *WalletStage7702CallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage7702.Contract.ReadNonce(&_WalletStage7702.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage7702 *WalletStage7702Caller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletStage7702 *WalletStage7702Session) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage7702.Contract.RecoverPartialSignature(&_WalletStage7702.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage7702 *WalletStage7702CallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage7702.Contract.RecoverPartialSignature(&_WalletStage7702.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage7702 *WalletStage7702Caller) RecoverSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "recoverSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage7702 *WalletStage7702Session) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage7702.Contract.RecoverSapientSignature(&_WalletStage7702.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage7702 *WalletStage7702CallerSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage7702.Contract.RecoverSapientSignature(&_WalletStage7702.CallOpts, _payload, _signature)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Caller) TokenReceived(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage7702.contract.Call(opts, &out, "tokenReceived", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702Session) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.TokenReceived(&_WalletStage7702.CallOpts, arg0, arg1, arg2)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage7702 *WalletStage7702CallerSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletStage7702.Contract.TokenReceived(&_WalletStage7702.CallOpts, arg0, arg1, arg2)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) AddHook(opts *bind.TransactOpts, selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "addHook", selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage7702 *WalletStage7702Session) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage7702.Contract.AddHook(&_WalletStage7702.TransactOpts, selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage7702.Contract.AddHook(&_WalletStage7702.TransactOpts, selector, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage7702 *WalletStage7702Session) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.Execute(&_WalletStage7702.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.Execute(&_WalletStage7702.TransactOpts, _payload, _signature)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage7702 *WalletStage7702Transactor) ExecuteUserOp(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "executeUserOp", _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage7702 *WalletStage7702Session) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.ExecuteUserOp(&_WalletStage7702.TransactOpts, _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.ExecuteUserOp(&_WalletStage7702.TransactOpts, _payload)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) RemoveHook(opts *bind.TransactOpts, selector [4]byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "removeHook", selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage7702 *WalletStage7702Session) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.RemoveHook(&_WalletStage7702.TransactOpts, selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.RemoveHook(&_WalletStage7702.TransactOpts, selector)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage7702 *WalletStage7702Session) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.SelfExecute(&_WalletStage7702.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.SelfExecute(&_WalletStage7702.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage7702 *WalletStage7702Transactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage7702 *WalletStage7702Session) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.Contract.SetStaticSignature(&_WalletStage7702.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.Contract.SetStaticSignature(&_WalletStage7702.TransactOpts, _hash, _address, _timestamp)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage7702 *WalletStage7702Transactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "tokensReceived", operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage7702 *WalletStage7702Session) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.TokensReceived(&_WalletStage7702.TransactOpts, operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.TokensReceived(&_WalletStage7702.TransactOpts, operator, from, to, amount, data, operatorData)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage7702 *WalletStage7702Transactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage7702 *WalletStage7702Session) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.UpdateImageHash(&_WalletStage7702.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.UpdateImageHash(&_WalletStage7702.TransactOpts, _imageHash)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage7702 *WalletStage7702Transactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage7702 *WalletStage7702Session) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.Contract.ValidateUserOp(&_WalletStage7702.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage7702 *WalletStage7702TransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage7702.Contract.ValidateUserOp(&_WalletStage7702.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletStage7702.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage7702 *WalletStage7702Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.Fallback(&_WalletStage7702.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage7702.Contract.Fallback(&_WalletStage7702.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage7702 *WalletStage7702Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage7702.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage7702 *WalletStage7702Session) Receive() (*types.Transaction, error) {
	return _WalletStage7702.Contract.Receive(&_WalletStage7702.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage7702 *WalletStage7702TransactorSession) Receive() (*types.Transaction, error) {
	return _WalletStage7702.Contract.Receive(&_WalletStage7702.TransactOpts)
}

// WalletStage7702CallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletStage7702 contract.
type WalletStage7702CallAbortedIterator struct {
	Event *WalletStage7702CallAborted // Event containing the contract specifics and raw log

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
func (it *WalletStage7702CallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702CallAborted)
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
		it.Event = new(WalletStage7702CallAborted)
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
func (it *WalletStage7702CallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702CallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702CallAborted represents a CallAborted event raised by the WalletStage7702 contract.
type WalletStage7702CallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) FilterCallAborted(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage7702CallAbortedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "CallAborted", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702CallAbortedIterator{contract: _WalletStage7702.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage7702CallAborted, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "CallAborted", _opHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702CallAborted)
				if err := _WalletStage7702.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) ParseCallAborted(log types.Log) (*WalletStage7702CallAborted, error) {
	event := new(WalletStage7702CallAborted)
	if err := _WalletStage7702.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702CallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletStage7702 contract.
type WalletStage7702CallFailedIterator struct {
	Event *WalletStage7702CallFailed // Event containing the contract specifics and raw log

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
func (it *WalletStage7702CallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702CallFailed)
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
		it.Event = new(WalletStage7702CallFailed)
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
func (it *WalletStage7702CallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702CallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702CallFailed represents a CallFailed event raised by the WalletStage7702 contract.
type WalletStage7702CallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) FilterCallFailed(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage7702CallFailedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "CallFailed", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702CallFailedIterator{contract: _WalletStage7702.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage7702CallFailed, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "CallFailed", _opHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702CallFailed)
				if err := _WalletStage7702.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage7702 *WalletStage7702Filterer) ParseCallFailed(log types.Log) (*WalletStage7702CallFailed, error) {
	event := new(WalletStage7702CallFailed)
	if err := _WalletStage7702.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702CallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletStage7702 contract.
type WalletStage7702CallSkippedIterator struct {
	Event *WalletStage7702CallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletStage7702CallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702CallSkipped)
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
		it.Event = new(WalletStage7702CallSkipped)
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
func (it *WalletStage7702CallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702CallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702CallSkipped represents a CallSkipped event raised by the WalletStage7702 contract.
type WalletStage7702CallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) FilterCallSkipped(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage7702CallSkippedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "CallSkipped", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702CallSkippedIterator{contract: _WalletStage7702.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage7702CallSkipped, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "CallSkipped", _opHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702CallSkipped)
				if err := _WalletStage7702.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) ParseCallSkipped(log types.Log) (*WalletStage7702CallSkipped, error) {
	event := new(WalletStage7702CallSkipped)
	if err := _WalletStage7702.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702CallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletStage7702 contract.
type WalletStage7702CallSucceededIterator struct {
	Event *WalletStage7702CallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletStage7702CallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702CallSucceeded)
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
		it.Event = new(WalletStage7702CallSucceeded)
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
func (it *WalletStage7702CallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702CallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702CallSucceeded represents a CallSucceeded event raised by the WalletStage7702 contract.
type WalletStage7702CallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) FilterCallSucceeded(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage7702CallSucceededIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "CallSucceeded", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage7702CallSucceededIterator{contract: _WalletStage7702.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletStage7702CallSucceeded, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "CallSucceeded", _opHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702CallSucceeded)
				if err := _WalletStage7702.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage7702 *WalletStage7702Filterer) ParseCallSucceeded(log types.Log) (*WalletStage7702CallSucceeded, error) {
	event := new(WalletStage7702CallSucceeded)
	if err := _WalletStage7702.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702DefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletStage7702 contract.
type WalletStage7702DefinedHookIterator struct {
	Event *WalletStage7702DefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletStage7702DefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702DefinedHook)
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
		it.Event = new(WalletStage7702DefinedHook)
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
func (it *WalletStage7702DefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702DefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702DefinedHook represents a DefinedHook event raised by the WalletStage7702 contract.
type WalletStage7702DefinedHook struct {
	Selector       [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletStage7702 *WalletStage7702Filterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage7702DefinedHookIterator, error) {

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage7702DefinedHookIterator{contract: _WalletStage7702.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletStage7702 *WalletStage7702Filterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletStage7702DefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702DefinedHook)
				if err := _WalletStage7702.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
func (_WalletStage7702 *WalletStage7702Filterer) ParseDefinedHook(log types.Log) (*WalletStage7702DefinedHook, error) {
	event := new(WalletStage7702DefinedHook)
	if err := _WalletStage7702.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702ImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletStage7702 contract.
type WalletStage7702ImageHashUpdatedIterator struct {
	Event *WalletStage7702ImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage7702ImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702ImageHashUpdated)
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
		it.Event = new(WalletStage7702ImageHashUpdated)
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
func (it *WalletStage7702ImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702ImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702ImageHashUpdated represents a ImageHashUpdated event raised by the WalletStage7702 contract.
type WalletStage7702ImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage7702 *WalletStage7702Filterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletStage7702ImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage7702ImageHashUpdatedIterator{contract: _WalletStage7702.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage7702 *WalletStage7702Filterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage7702ImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702ImageHashUpdated)
				if err := _WalletStage7702.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletStage7702 *WalletStage7702Filterer) ParseImageHashUpdated(log types.Log) (*WalletStage7702ImageHashUpdated, error) {
	event := new(WalletStage7702ImageHashUpdated)
	if err := _WalletStage7702.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702NonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletStage7702 contract.
type WalletStage7702NonceChangeIterator struct {
	Event *WalletStage7702NonceChange // Event containing the contract specifics and raw log

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
func (it *WalletStage7702NonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702NonceChange)
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
		it.Event = new(WalletStage7702NonceChange)
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
func (it *WalletStage7702NonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702NonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702NonceChange represents a NonceChange event raised by the WalletStage7702 contract.
type WalletStage7702NonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage7702 *WalletStage7702Filterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletStage7702NonceChangeIterator, error) {

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletStage7702NonceChangeIterator{contract: _WalletStage7702.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage7702 *WalletStage7702Filterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletStage7702NonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702NonceChange)
				if err := _WalletStage7702.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletStage7702 *WalletStage7702Filterer) ParseNonceChange(log types.Log) (*WalletStage7702NonceChange, error) {
	event := new(WalletStage7702NonceChange)
	if err := _WalletStage7702.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage7702StaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletStage7702 contract.
type WalletStage7702StaticSignatureSetIterator struct {
	Event *WalletStage7702StaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletStage7702StaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage7702StaticSignatureSet)
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
		it.Event = new(WalletStage7702StaticSignatureSet)
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
func (it *WalletStage7702StaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage7702StaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage7702StaticSignatureSet represents a StaticSignatureSet event raised by the WalletStage7702 contract.
type WalletStage7702StaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage7702 *WalletStage7702Filterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletStage7702StaticSignatureSetIterator, error) {

	logs, sub, err := _WalletStage7702.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletStage7702StaticSignatureSetIterator{contract: _WalletStage7702.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage7702 *WalletStage7702Filterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletStage7702StaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletStage7702.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage7702StaticSignatureSet)
				if err := _WalletStage7702.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletStage7702 *WalletStage7702Filterer) ParseStaticSignatureSet(log types.Log) (*WalletStage7702StaticSignatureSet, error) {
	event := new(WalletStage7702StaticSignatureSet)
	if err := _WalletStage7702.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
