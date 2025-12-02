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

// WalletStage1MetaData contains all meta data concerning the WalletStage1 contract.
var WalletStage1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_CODE_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAGE_2_IMPLEMENTATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainedSignatureNestedInChainedSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidPackedLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x61010080604052346101ae576040816182bf803803809161002082856101b3565b8339810103126101ae5761003f6020610038836101d6565b92016101d6565b604051909190613fcf808201906001600160401b0382118383101761018c5760209183916142f083396001600160a01b03861681520301906000f09081156101a257604051606081016001600160401b0381118282101761018c576040908152602f82527f6041600e3d396021805130553df33d3d36153402601f57363d3d373d363d305460208301526e5af43d82803e903d91601f57fd5bf360881b82820152519060005b602f811061017757505030604f820152604f8152610104606f826101b3565b805160209091012060805260a0526001600160a01b031660c05260e05260405161410590816101eb8239608051818181610c510152613d29015260a051818181610d9f0152613cfa015260c051818181610d060152611244015260e0518181816110dd015281816112b30152611b1a0152f35b806020809284010151828286010152016100e5565b634e487b7160e01b600052604160045260246000fd5b6040513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761018c57604052565b51906001600160a01b03821682036101ae5756fe6080604052600436101561001e575b361561001c5761001c611da2565b005b60003560e01c806223de29146101ad578063025b22bc146101a857806313792a4a146101a3578063150b7a021461019e5780631626ba7e1461019957806319822f7c146101945780631a9b23371461018f5780631f6a1eb91461018a578063257671f51461018557806329561426146101805780632dd310001461017b5780634fcf3eca146101765780636ea44577146101715780638943ec021461016c5780638c3f55631461016757806392dcb3fc146101625780639c145aed1461015d5780639f69ef5414610158578063a65d69d414610153578063aaf10f421461014e578063ad55366b14610149578063b93ea7ad14610144578063bc197c811461013f578063f23a6e611461013a5763f727ef1c0361000e57611632565b6115a5565b6114d3565b611374565b611328565b6112d7565b611268565b6111f9565b61106a565b61100c565b610fd0565b610f4c565b610f1d565b610dc3565b610d54565b610c74565b610c1b565b610aff565b610a9c565b6109e7565b61095f565b6108d2565b6107d5565b6102fb565b61026f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b9181601f840112156101d55782359167ffffffffffffffff83116101d557602083818601950101116101d557565b346101d55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576102a66101b2565b506102af6101da565b506102b86101fd565b5060843567ffffffffffffffff81116101d5576102d9903690600401610241565b505060a43567ffffffffffffffff81116101d55761001c903690600401610241565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55761032d6101b2565b30330361033d5761001c90611e6b565b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103b657604052565b61036b565b6040810190811067ffffffffffffffff8211176103b657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103b657604052565b6040519061042760e0836103d7565b565b60405190610427610120836103d7565b359060ff821682036101d557565b359081151582036101d557565b67ffffffffffffffff81116103b65760051b60200190565b67ffffffffffffffff81116103b657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104b28261046c565b916104c060405193846103d7565b8294818452818301116101d5578281602093846000960137010152565b9080601f830112156101d5578160206104f8933591016104a6565b90565b81601f820112156101d55780359061051282610454565b9261052060405194856103d7565b82845260208085019360051b830101918183116101d55760208101935b83851061054c57505050505090565b843567ffffffffffffffff81116101d557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101d557610593610418565b916105a060208301610220565b83526040820135602084015260608201359267ffffffffffffffff84116101d55760e0836105d58860208098819801016104dd565b6040840152608081013560608401526105f060a08201610447565b608084015261060160c08201610447565b60a0840152013560c082015281520194019361053d565b9080601f830112156101d557813561062f81610454565b9261063d60405194856103d7565b81845260208085019260051b8201019283116101d557602001905b8282106106655750505090565b6020809161067284610220565b815201910190610658565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101d55760043567ffffffffffffffff81116101d5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101d5576106ef610429565b906106fc81600401610439565b825261070a60248201610447565b6020830152604481013567ffffffffffffffff81116101d557846004610732928401016104fb565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101d55784600461076e928401016104dd565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101d557600485916107ab930101610618565b610100820152916024359067ffffffffffffffff82116101d5576107d191600401610241565b9091565b346101d5576107e33661067d565b90916101008101926107fe6107f985515161173a565b61175a565b9160005b85518051821015610865579061085f61083a610820836001956117d8565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61084483886117d8565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610802565b505083838661087a33610844835151856117d8565b52610886818484611eba565b50156108985760405160018152602090f35b6108ce906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611a33565b0390fd5b346101d55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576109096101b2565b506109126101da565b5060643567ffffffffffffffff81116101d557610933903690600401610241565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101d55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043560243567ffffffffffffffff81116101d5576020916109b76109bd923690600401610241565b91611a58565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043567ffffffffffffffff81116101d5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101d557610a6a60209160243560443591600401611b01565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101d557565b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020610ae1600435610adc81610a72565b6120ca565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043567ffffffffffffffff81116101d557610b49903690600401610241565b60243567ffffffffffffffff81116101d557610b69903690600401610241565b9160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610be657610bc09360027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55611c6b565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b60009103126101d557565b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043530330361033d578015610d2a576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a161001c7f0000000000000000000000000000000000000000000000000000000000000000611e6b565b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557600435610df981610a72565b30330361033d5773ffffffffffffffffffffffffffffffffffffffff610e1e826120ca565b1615610e845760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610e75600082612ea4565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101d5576004359067ffffffffffffffff82116101d5576107d191600401610241565b610f2636610ed4565b9030330361033d57610f3c61001c925a926121ff565b90610f46826128a3565b9061257e565b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557610f836101b2565b5060443567ffffffffffffffff81116101d557610fa4903690600401610241565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020610a6a600435612937565b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557604061104860043561297d565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101d55761107836610ed4565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610be65760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156111cf5733036111a157303b156101d55761114b9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611d61565b038183305af1801561119c576111815760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b806111906000611196936103d7565b80610c10565b38610bc0565b611aa4565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101d55760c061134460008061133e3661067d565b91612a7e565b92611350839293613c96565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576004356113aa81610a72565b6113b26101da565b9030330361033d5773ffffffffffffffffffffffffffffffffffffffff6113d8826120ca565b16611452577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166114458183612ea4565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101d55782359167ffffffffffffffff83116101d5576020808501948460051b0101116101d557565b346101d55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55761150a6101b2565b506115136101da565b5060443567ffffffffffffffff81116101d5576115349036906004016114a2565b505060643567ffffffffffffffff81116101d5576115569036906004016114a2565b505060843567ffffffffffffffff81116101d557611578903690600401610241565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101d55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576115dc6101b2565b506115e56101da565b5060843567ffffffffffffffff81116101d557611606903690600401610241565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043561166c6101da565b604435916bffffffffffffffffffffffff83168093036101d55730330361033d578273ffffffffffffffffffffffffffffffffffffffff836116f77febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612e20565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161174857565b61170b565b9190820180921161174857565b9061176482610454565b61177160405191826103d7565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061179f8294610454565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117ec5760209160051b010190565b6117a9565b919082519283825260005b84811061183b5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b806020809284010151828286010152016117fc565b9080602083519182815201916020808360051b8301019401926000915b83831061187c57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c0806118ef604085015160e0604086015260e08501906117f1565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061186d565b906020808351928381520192019060005b8181106119445750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611937565b805160ff1682526104f8916020828101511515908201526101006119ce6119a860408501516101206040860152610120850190611850565b606085015160608501526080850151608085015260a085015184820360a08601526117f1565b9260c081015160c084015260e081015160e0840152015190610100818403910152611926565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611a4a6104f89492604085526040850190611970565b9260208185039101526119f4565b90611a759291611a6661206f565b906003825260e0820152611eba565b5015611a9f577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101d5570180359067ffffffffffffffff82116101d5576020019181360383136101d557565b909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169081156111cf578133036111a15780611bb5575b5050611ba57f1626ba7e00000000000000000000000000000000000000000000000000000000926109b7836101007fffffffff00000000000000000000000000000000000000000000000000000000950190611ab0565b1603611bb057600090565b600190565b813b156101d5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af192831561119c576109b77fffffffff0000000000000000000000000000000000000000000000000000000093611ba5937f1626ba7e0000000000000000000000000000000000000000000000000000000096611c56575b5093505092611b4e565b806111906000611c65936103d7565b38611c4c565b91939290611c7a905a936121ff565b9160608301516080840151611c8e82612937565b818103611d2d57509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881920190611cc78282612e65565b604080519182526020820192909252a1611ce2828685611eba565b929015611cf5575061042793945061257e565b836108ce87926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611a33565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b9160206104f89381815201916119f4565b3d15611d9d573d90611d838261046c565b91611d9160405193846103d7565b82523d6000602084013e565b606090565b600436108015611daf5750565b611de5906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611e37575b506120ca565b73ffffffffffffffffffffffffffffffffffffffff8116611e035750565b60008091604051368382378036810184815203915af4611e21611d72565b9015611e2f57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611ddf565b60207f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039180305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1565b90156117ec5790565b611eed611ec78484611eb1565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f800000000000000000000000000000000000000000000000000000000000000080821614611f705750611f25926000928392612a7e565b905091909192808210611f40575050611f3d90613c96565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0200000000000000000000000000000000000000000000000000000000000000908116146020820152611fa6925090506128a3565b90611fb08261297d565b4281111561203d575073ffffffffffffffffffffffffffffffffffffffff81168015159081612032575b50611fe6575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611fda565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103b6576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208452166040820152604081526121416060826103d7565b519020541690565b9061215382610454565b61216060405191826103d7565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061218e8294610454565b019060005b82811061219f57505050565b6020906040516121ae8161039a565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201612193565b909392938483116101d55784116101d5578101920390565b9061220861206f565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016124d75750600060608601525b600761225760ff831660011c90565b1680612486575b50601081811603612458575060015b61227681612149565b604086019081526000925b8284106122be5750505050036122945790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c906001019490856001808316036124365750612308306122ec8487516117d8565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280821614612416575b6004808216146123c8575b600880821614612393575b9061237a61237460c0846123546010806001981614608061234b888c516117d8565b51019015159052565b61236a60208083161460a061234b888c516117d8565b1660061c60031690565b60ff1690565b60c06123878387516117d8565b51015201929190612281565b946001919061237a906123749060c0908681013590602001999060606123ba878b516117d8565b510152939450505050612329565b94612410908381013560e81c906003016124096123f06123e8848461174d565b838c896121e7565b919060406123ff888b516117d8565b51019236916104a6565b905261174d565b9461231e565b9482810135906020019590602061242e8487516117d8565b510152612313565b61245396508381013560601c9060140196906122ec8487516117d8565b612308565b60209081160361247557600282019181013560f01c905b9061226d565b600182019181013560f81c9061246f565b6124ca919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b929060808601523861225e565b80830135606090811c908701526014019250612248565b61250660409295949395606083526060830190611970565b9460208201520152565b926104f896959260c09592855260208501526040840152606083015260808201528160a082015201906117f1565b6040906104f89392815281602082015201906117f1565b61256b6104f89492606083526060830190611970565b92602082015260408184039101526117f1565b916000604082019384515190825b82811061259d575b50505050505050565b6125a88188516117d8565b51936125b760a0860151151590565b8061289b575b612861575060009360608101518015801580612858575b6128205784906125e76080850151151590565b156127da576126929261260e855173ffffffffffffffffffffffffffffffffffffffff1690565b91156127d457505a905b61268d8b61266160608d01516040890151908c8b604051998a967f4c4e814c00000000000000000000000000000000000000000000000000000000602089015260248801612510565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018552846103d7565b612df3565b156126db575b50600190867f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a604051806126d185829190602083019252565b0390a25b0161258c565b60c00180511561278b57600181511461274c57516002146126fc5738612698565b93505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b91925061272d612e05565b9061273d6040519283928361253e565b0390a238808080808080612594565b50846108ce612759612e05565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612555565b50925060018093867f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6127cc6127bf612e05565b604051918291868361253e565b0390a26126d5565b90612618565b835161281593925073ffffffffffffffffffffffffffffffffffffffff169160208501519160001461281a57505a905b604085015192612de1565b612692565b9061280a565b83886108ce5a6040519384937f21395274000000000000000000000000000000000000000000000000000000008552600485016124ee565b50815a106125d4565b9350600190867f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b604051806127cc85829190602083019252565b5080156125bd565b6129056129316128c36128bd602085015115153090612ee3565b93612fde565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103d7565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83526040820152604081526129766060826103d7565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526129bc6060826103d7565b51902054906bffffffffffffffffffffffff8260601c921690565b604051906129e4826103bb565b60006020838281520152565b600311156129fa57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160409103126101d557602060405191612a43836103bb565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff6104f8949316815281602082015201906117f1565b909491939291853560f81c600190938190612a976129d7565b92612aa1826129f0565b60018203612cab575b50600180871614612c4a575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612b199060051c90565b612b229061173a565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612b6e846128a3565b988993612b7a936121e7565b91612b8493613420565b9098612b9891600052602052604060002090565b90612bab91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612bd491600052602052604060002090565b94815190868215159283612c3f575b505081612c30575b50612bf35750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612beb565b141591508638612be3565b909691939450612c5c819893986129f0565b612c8157612c769581612c6e936121e7565b9390926131d3565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612aaa578981013560601c9750601401915087878a84612cf4856129f0565b60028503612d05575b505050612aaa565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c92604092612d8a929091612d5591612d4e918a90612d48898361174d565b926121e7565b36916104a6565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612a51565b0392165afa801561119c57612da892600091612db2575b509361174d565b9087388a81612cfd565b612dd4915060403d604011612dda575b612dcc81836103d7565b810190612a29565b38612da1565b503d612dc2565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612e5f6060826103d7565b51902055565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612e5f6060826103d7565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612e5f6060826103d7565b15612f93576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261293160c0826103d7565b4690612eec565b805160209091019060005b818110612fb25750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612fa5565b610100810151604051612ff981612905602082018095612f9a565b51902090613008815160ff1690565b60ff811680613081575050906129316130246040840151613d93565b9261290560806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b600181036130df57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466938101938452908101919091526060810192909252906129318160808101612905565b6002810361313557505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082019081529181019290925260608201929092526129318160808101612905565b600303613189575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082019081529181019290925260608201929092526129318160808101612905565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b9061250690604093969594966060845260608401916119f4565b91949290926000956000956000956000956000956131ef61206f565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061324657505050505050508051151580613238575b612bf35750565b506020810151841115613231565b600381019d50959b509399509197509290919061326a908b9085013560e81c61174d565b958287036133aa578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361337b57506132b7916132b0898c9387896121e7565b908b612a7e565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106133355750928b88511461332c575b808b10156132fa57508a60c08501528992959295949194939093613219565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526132db565b8d8f6108ce61334685858c8e6121e7565b9390926040519485947fb006aba0000000000000000000000000000000000000000000000000000000008652600486016131b9565b979899809b926133908b61339794888a6121e7565b9086612a7e565b50929d919c909b929a9092918e8e6132cb565b8a600291613277565b908160209103126101d557516104f881610a72565b6040906104f89492815281602082015201916119f4565b73ffffffffffffffffffffffffffffffffffffffff6104f89593606093835216602082015281604082015201916119f4565b908160209103126101d5575190565b9391909360009460009460005b81811061343b575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613c2f575060018914613bef5760028914613a2257600389146139f3576004891461397257600689146138d2576005891461388457600789146137bd5760088914613767576009891461363e57600a89146134db577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b9091929394959697506003891697881561362d575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816135539187876121e7565b6040517f898bd9210000000000000000000000000000000000000000000000000000000081529391849161358b918a600485016133c8565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561119c576135ce936000936135fa575b5060ff909a168091019a61405f565b9080156135f457906135e891600052602052604060002090565b955b939291909361342d565b506135e8565b60ff91935061361f9060203d8111613626575b61361781836103d7565b810190613411565b92906135bf565b503d61360d565b8084013560f81c98506001016134f0565b90919293949596975060038916978815613756575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816136b69187876121e7565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916136ee918b60048501611a33565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561119c57613730936000936135fa575060ff909a168091019a61405f565b908015613750579061374a91600052602052604060002090565b956135ea565b5061374a565b8084013560f81c9850600101613653565b98506020870197509495939492939192909182013561378586614006565b8114613795575b61373090614020565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff985061378c565b975090919293949597600f16968715613872575b602060006137e36138509a9b86613ece565b9c9092918a604051613826816129058a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561119c576137309060ff6000519a1680910199613f11565b600189019883013560f81c97506137d1565b9850602087019750949593949293919290918201358085146138aa575b61373090613fc7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506138a1565b989091929394959662ffffff98506138f4612374600c8416603f9060021c1690565b91821561395e575b600316801561394d575b90819061393190613929908781013560e81c906003019c168c01809c89896121e7565b90898b613420565b911115613944575b906137309291613f7c565b99820199613939565b50600281019084013560f01c613906565b8482013560f81c92506001909101906138fc565b9750976139c86139d5929394959697600f6139dd93169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92908301809386866121e7565b908688613420565b9061374a92980198600052602052604060002090565b985096509394929391929091908082013590602001968015613750579061374a91600052602052604060002090565b90919293949596975060038916978815613bde575b8084013560601c99613a969160140190613a569060021c600316612374565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613afa60208c613aac85858b8b6121e7565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e000000000000000000000000000000000000000000000000000000008552600485016133c8565b0392165afa90811561119c577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613bb0575b501603613b6c57509060ff61373092991680910199613f11565b6108ce613b7d8c93899389896121e7565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016133df565b613bd1915060203d8111613bd7575b613bc981836103d7565b8101906133b3565b38613b52565b503d613bbf565b8381013560f81c9850600101613a37565b98600f91929394959697985016968715613c1e575b60148101976137309160ff9091169084013560601c613f11565b8281013560f81c9750600101613c04565b98509091929394959698600f16978815613c81575b5060206000613c576138509a9b86613ece565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613c44565b73ffffffffffffffffffffffffffffffffffffffff9060405160208101917fff0000000000000000000000000000000000000000000000000000000000000083527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060601b16602183015260358201527f0000000000000000000000000000000000000000000000000000000000000000605582015260558152613d5c6075826103d7565b51902016301490565b805160209091019060005b818110613d7d5750505090565b8251845260209384019390920191600101613d70565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613dd9613dc383610454565b92613dd160405194856103d7565b808452610454565b0136602083013760005b8351811015613eb55780613df9600192866117d8565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613ea1610120826103d7565b519020613eae82856117d8565b5201613de3565b5090915060405161293181612905602082018095613d65565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116117485791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526129316065826103d7565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526129316098826103d7565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526129316058826103d7565b6129056129316128c36128bd600060208601511515612ee3565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526129316060826103d7565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612931608d826103d756fea26469706673582212204aded9bebadf620767a8456d364bea5dce9e2213033e4904099efbbfaacb1aca64736f6c634300081c003360a034607457601f613fcf38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b0381168103607457608052604051613f3f90816100908239608051818181611066015281816111cd0152611a340152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c611cbc565b005b60003560e01c806223de291461018d578063025b22bc1461018857806313792a4a14610183578063150b7a021461017e5780631626ba7e1461017957806319822f7c146101745780631a9b23371461016f5780631f6a1eb91461016a57806329561426146101655780634fcf3eca1461016057806351605d801461015b5780636ea44577146101565780638943ec02146101515780638c3f55631461014c57806392dcb3fc146101475780639c145aed14610142578063a65d69d41461013d578063aaf10f4214610138578063ad55366b14610133578063b93ea7ad1461012e578063bc197c8114610129578063f23a6e61146101245763f727ef1c0361000e5761154c565b6114bf565b6113ed565b61128e565b611242565b6111f1565b611182565b610ff3565b610f95565b610f59565b610ed5565b610ea6565b610e02565b610ce6565b610c2d565b610b1c565b610ab9565b610a04565b61097c565b6108ef565b6107f2565b6102db565b61024f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b9181601f840112156101b55782359167ffffffffffffffff83116101b557602083818601950101116101b557565b346101b55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610286610192565b5061028f6101ba565b506102986101dd565b5060843567ffffffffffffffff81116101b5576102b9903690600401610221565b505060a43567ffffffffffffffff81116101b55761001c903690600401610221565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761030d610192565b30330361035a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103d357604052565b610388565b6040810190811067ffffffffffffffff8211176103d357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103d357604052565b6040519061044460e0836103f4565b565b60405190610444610120836103f4565b359060ff821682036101b557565b359081151582036101b557565b67ffffffffffffffff81116103d35760051b60200190565b67ffffffffffffffff81116103d357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104cf82610489565b916104dd60405193846103f4565b8294818452818301116101b5578281602093846000960137010152565b9080601f830112156101b557816020610515933591016104c3565b90565b81601f820112156101b55780359061052f82610471565b9261053d60405194856103f4565b82845260208085019360051b830101918183116101b55760208101935b83851061056957505050505090565b843567ffffffffffffffff81116101b557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101b5576105b0610435565b916105bd60208301610200565b83526040820135602084015260608201359267ffffffffffffffff84116101b55760e0836105f28860208098819801016104fa565b60408401526080810135606084015261060d60a08201610464565b608084015261061e60c08201610464565b60a0840152013560c082015281520194019361055a565b9080601f830112156101b557813561064c81610471565b9261065a60405194856103f4565b81845260208085019260051b8201019283116101b557602001905b8282106106825750505090565b6020809161068f84610200565b815201910190610675565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101b55761070c610446565b9061071981600401610456565b825261072760248201610464565b6020830152604481013567ffffffffffffffff81116101b55784600461074f92840101610518565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101b55784600461078b928401016104fa565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101b557600485916107c8930101610635565b610100820152916024359067ffffffffffffffff82116101b5576107ee91600401610221565b9091565b346101b5576108003661069a565b909161010081019261081b610816855151611654565b611674565b9160005b85518051821015610882579061087c61085761083d836001956116f2565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61086183886116f2565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161081f565b505083838661089733610861835151856116f2565b526108a3818484611d8e565b50156108b55760405160018152602090f35b6108eb906040519384937ff58cc8b50000000000000000000000000000000000000000000000000000000085526004850161194d565b0390fd5b346101b55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610926610192565b5061092f6101ba565b5060643567ffffffffffffffff81116101b557610950903690600401610221565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101b55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043560243567ffffffffffffffff81116101b5576020916109d46109da923690600401610221565b91611972565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b557610a8760209160243560443591600401611a1b565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101b557565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020610afe600435610af981610a8f565b611f9e565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b557610b66903690600401610221565b60243567ffffffffffffffff81116101b557610b86903690600401610221565b9160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c0357610bdd9360027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55611b85565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043530330361035a578015610cbc576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557600435610d1c81610a8f565b30330361035a5773ffffffffffffffffffffffffffffffffffffffff610d4182611f9e565b1615610da75760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d98600082612d78565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101b557565b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101b5576004359067ffffffffffffffff82116101b5576107ee91600401610221565b610eaf36610e5d565b9030330361035a57610ec561001c925a926120d3565b90610ecf82612777565b90612452565b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610f0c610192565b5060443567ffffffffffffffff81116101b557610f2d903690600401610221565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020610a8760043561280b565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576040610fd1600435612851565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101b55761100136610e5d565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c035760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561115857330361112a57303b156101b5576110d49160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611c7b565b038183305af180156111255761110a5760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b80611119600061111f936103f4565b80610df7565b38610bdd565b6119be565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101b55760c061125e6000806112583661069a565b91612952565b9261126a839293613b6a565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576004356112c481610a8f565b6112cc6101ba565b9030330361035a5773ffffffffffffffffffffffffffffffffffffffff6112f282611f9e565b1661136c577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000060409316911661135f8183612d78565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101b55782359167ffffffffffffffff83116101b5576020808501948460051b0101116101b557565b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557611424610192565b5061142d6101ba565b5060443567ffffffffffffffff81116101b55761144e9036906004016113bc565b505060643567ffffffffffffffff81116101b5576114709036906004016113bc565b505060843567ffffffffffffffff81116101b557611492903690600401610221565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576114f6610192565b506114ff6101ba565b5060843567ffffffffffffffff81116101b557611520903690600401610221565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576004356115866101ba565b604435916bffffffffffffffffffffffff83168093036101b55730330361035a578273ffffffffffffffffffffffffffffffffffffffff836116117febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612cf4565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161166257565b611625565b9190820180921161166257565b9061167e82610471565b61168b60405191826103f4565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116b98294610471565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117065760209160051b010190565b6116c3565b919082519283825260005b8481106117555750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611716565b9080602083519182815201916020808360051b8301019401926000915b83831061179657505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611809604085015160e0604086015260e085019061170b565b936060810151606085015260808101511515608085015260a0810151151560a0850152015191015297019301930191939290611787565b906020808351928381520192019060005b81811061185e5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611851565b805160ff168252610515916020828101511515908201526101006118e86118c26040850151610120604086015261012085019061176a565b606085015160608501526080850151608085015260a085015184820360a086015261170b565b9260c081015160c084015260e081015160e0840152015190610100818403910152611840565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611964610515949260408552604085019061188a565b92602081850391015261190e565b9061198f9291611980611f43565b906003825260e0820152611d8e565b50156119b9577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101b5570180359067ffffffffffffffff82116101b5576020019181360383136101b557565b909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169081156111585781330361112a5780611acf575b5050611abf7f1626ba7e00000000000000000000000000000000000000000000000000000000926109d4836101007fffffffff000000000000000000000000000000000000000000000000000000009501906119ca565b1603611aca57600090565b600190565b813b156101b5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af1928315611125576109d47fffffffff0000000000000000000000000000000000000000000000000000000093611abf937f1626ba7e0000000000000000000000000000000000000000000000000000000096611b70575b5093505092611a68565b806111196000611b7f936103f4565b38611b66565b91939290611b94905a936120d3565b9160608301516080840151611ba88261280b565b818103611c4757509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881920190611be18282612d39565b604080519182526020820192909252a1611bfc828685611d8e565b929015611c0f5750610444939450612452565b836108eb87926040519384937fa2b6d61b0000000000000000000000000000000000000000000000000000000085526004850161194d565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b91602061051593818152019161190e565b3d15611cb7573d90611c9d82610489565b91611cab60405193846103f4565b82523d6000602084013e565b606090565b600436108015611cc95750565b611cff906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611d51575b50611f9e565b73ffffffffffffffffffffffffffffffffffffffff8116611d1d5750565b60008091604051368382378036810184815203915af4611d3b611c8c565b9015611d4957602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611cf9565b90156117065790565b611dc1611d9b8484611d85565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f800000000000000000000000000000000000000000000000000000000000000080821614611e445750611df9926000928392612952565b905091909192808210611e14575050611e1190613b6a565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0200000000000000000000000000000000000000000000000000000000000000908116146020820152611e7a92509050612777565b90611e8482612851565b42811115611f11575073ffffffffffffffffffffffffffffffffffffffff81168015159081611f06575b50611eba575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611eae565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103d3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208452166040820152604081526120156060826103f4565b519020541690565b9061202782610471565b61203460405191826103f4565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120628294610471565b019060005b82811061207357505050565b602090604051612082816103b7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201612067565b909392938483116101b55784116101b5578101920390565b906120dc611f43565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016123ab5750600060608601525b600761212b60ff831660011c90565b168061235a575b5060108181160361232c575060015b61214a8161201d565b604086019081526000925b8284106121925750505050036121685790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361230a57506121dc306121c08487516116f2565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808216146122ea575b60048082161461229c575b600880821614612267575b9061224e61224860c0846122286010806001981614608061221f888c516116f2565b51019015159052565b61223e60208083161460a061221f888c516116f2565b1660061c60031690565b60ff1690565b60c061225b8387516116f2565b51015201929190612155565b946001919061224e906122489060c09086810135906020019990606061228e878b516116f2565b5101529394505050506121fd565b946122e4908381013560e81c906003016122dd6122c46122bc8484611667565b838c896120bb565b919060406122d3888b516116f2565b51019236916104c3565b9052611667565b946121f2565b948281013590602001959060206123028487516116f2565b5101526121e7565b61232796508381013560601c9060140196906121c08487516116f2565b6121dc565b60209081160361234957600282019181013560f01c905b90612141565b600182019181013560f81c90612343565b61239e919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608086015238612132565b80830135606090811c90870152601401925061211c565b6123da6040929594939560608352606083019061188a565b9460208201520152565b9261051596959260c09592855260208501526040840152606083015260808201528160a0820152019061170b565b60409061051593928152816020820152019061170b565b61243f610515949260608352606083019061188a565b926020820152604081840391015261170b565b916000604082019384515190825b828110612471575b50505050505050565b61247c8188516116f2565b519361248b60a0860151151590565b8061276f575b61273557506000936060810151801580158061272c575b6126f45784906124bb6080850151151590565b156126ae57612566926124e2855173ffffffffffffffffffffffffffffffffffffffff1690565b91156126a857505a905b6125618b61253560608d01516040890151908c8b604051998a967f4c4e814c000000000000000000000000000000000000000000000000000000006020890152602488016123e4565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018552846103f4565b612cc7565b156125af575b50600190867f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a604051806125a585829190602083019252565b0390a25b01612460565b60c00180511561265f57600181511461262057516002146125d0573861256c565b93505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b919250612601612cd9565b9061261160405192839283612412565b0390a238808080808080612468565b50846108eb61262d612cd9565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612429565b50925060018093867f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6126a0612693612cd9565b6040519182918683612412565b0390a26125a9565b906124ec565b83516126e993925073ffffffffffffffffffffffffffffffffffffffff16916020850151916000146126ee57505a905b604085015192612cb5565b612566565b906126de565b83886108eb5a6040519384937f21395274000000000000000000000000000000000000000000000000000000008552600485016123c2565b50815a106124a8565b9350600190867f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b604051806126a085829190602083019252565b508015612491565b6127d9612805612797612791602085015115153090612db7565b93612eb2565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103f4565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e835260408201526040815261284a6060826103f4565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526128906060826103f4565b51902054906bffffffffffffffffffffffff8260601c921690565b604051906128b8826103d8565b60006020838281520152565b600311156128ce57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160409103126101b557602060405191612917836103d8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff6105159493168152816020820152019061170b565b909491939291853560f81c60019093819061296b6128ab565b92612975826128c4565b60018203612b7f575b50600180871614612b1e575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019080969181966020166129ed9060051c90565b6129f690611654565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612a4284612777565b988993612a4e936120bb565b91612a58936132f4565b9098612a6c91600052602052604060002090565b90612a7f91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612aa891600052602052604060002090565b94815190868215159283612b13575b505081612b04575b50612ac75750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612abf565b141591508638612ab7565b909691939450612b30819893986128c4565b612b5557612b4a9581612b42936120bb565b9390926130a7565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0604088160161297e578981013560601c9750601401915087878a84612bc8856128c4565b60028503612bd9575b50505061297e565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c92604092612c5e929091612c2991612c22918a90612c1c8983611667565b926120bb565b36916104c3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612925565b0392165afa801561112557612c7c92600091612c86575b5093611667565b9087388a81612bd1565b612ca8915060403d604011612cae575b612ca081836103f4565b8101906128fd565b38612c75565b503d612c96565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612d336060826103f4565b51902055565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612d336060826103f4565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612d336060826103f4565b15612e67576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261280560c0826103f4565b4690612dc0565b805160209091019060005b818110612e865750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612e79565b610100810151604051612ecd816127d9602082018095612e6e565b51902090612edc815160ff1690565b60ff811680612f5557505090612805612ef86040840151613bcd565b926127d960806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612fb357505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4669381019384529081019190915260608101929092529061280581608081016127d9565b6002810361300957505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e46020820190815291810192909252606082019290925261280581608081016127d9565b60030361305d575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4666020820190815291810192909252606082019290925261280581608081016127d9565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906123da906040939695949660608452606084019161190e565b91949290926000956000956000956000956000956130c3611f43565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061311a5750505050505050805115158061310c575b612ac75750565b506020810151841115613105565b600381019d50959b509399509197509290919061313e908b9085013560e81c611667565b9582870361327e578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361324f575061318b91613184898c9387896120bb565b908b612952565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106132095750928b885114613200575b808b10156131ce57508a60c085015289929592959491949390936130ed565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526131af565b8d8f6108eb61321a85858c8e6120bb565b9390926040519485947fb006aba00000000000000000000000000000000000000000000000000000000086526004860161308d565b979899809b926132648b61326b94888a6120bb565b9086612952565b50929d919c909b929a9092918e8e61319f565b8a60029161314b565b908160209103126101b5575161051581610a8f565b60409061051594928152816020820152019161190e565b73ffffffffffffffffffffffffffffffffffffffff61051595936060938352166020820152816040820152019161190e565b908160209103126101b5575190565b9391909360009460009460005b81811061330f575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613b03575060018914613ac357600289146138f657600389146138c7576004891461384657600689146137a657600589146137585760078914613691576008891461363b576009891461351257600a89146133af577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613501575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816134279187876120bb565b6040517f898bd9210000000000000000000000000000000000000000000000000000000081529391849161345f918a6004850161329c565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611125576134a2936000936134ce575b5060ff909a168091019a613e99565b9080156134c857906134bc91600052602052604060002090565b955b9392919093613301565b506134bc565b60ff9193506134f39060203d81116134fa575b6134eb81836103f4565b8101906132e5565b9290613493565b503d6134e1565b8084013560f81c98506001016133c4565b9091929394959697506003891697881561362a575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161358a9187876120bb565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916135c2918b6004850161194d565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561112557613604936000936134ce575060ff909a168091019a613e99565b908015613624579061361e91600052602052604060002090565b956134be565b5061361e565b8084013560f81c9850600101613527565b98506020870197509495939492939192909182013561365986613e40565b8114613669575b61360490613e5a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613660565b975090919293949597600f16968715613746575b602060006136b76137249a9b86613d08565b9c9092918a6040516136fa816127d98a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15611125576136049060ff6000519a1680910199613d4b565b600189019883013560f81c97506136a5565b98506020870197509495939492939192909182013580851461377e575b61360490613e01565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613775565b989091929394959662ffffff98506137c8612248600c8416603f9060021c1690565b918215613832575b6003168015613821575b908190613805906137fd908781013560e81c906003019c168c01809c89896120bb565b90898b6132f4565b911115613818575b906136049291613db6565b9982019961380d565b50600281019084013560f01c6137da565b8482013560f81c92506001909101906137d0565b97509761389c6138a9929394959697600f6138b193169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92908301809386866120bb565b9086886132f4565b9061361e92980198600052602052604060002090565b985096509394929391929091908082013590602001968015613624579061361e91600052602052604060002090565b90919293949596975060038916978815613ab2575b8084013560601c9961396a916014019061392a9060021c600316612248565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b908101906139ce60208c61398085858b8b6120bb565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e0000000000000000000000000000000000000000000000000000000085526004850161329c565b0392165afa908115611125577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613a84575b501603613a4057509060ff61360492991680910199613d4b565b6108eb613a518c93899389896120bb565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016132b3565b613aa5915060203d8111613aab575b613a9d81836103f4565b810190613287565b38613a26565b503d613a93565b8381013560f81c985060010161390b565b98600f91929394959697985016968715613af2575b60148101976136049160ff9091169084013560601c613d4b565b8281013560f81c9750600101613ad8565b98509091929394959698600f16978815613b55575b5060206000613b2b6137249a9b86613d08565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613b18565b8015159081613b77575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b818110613bb75750505090565b8251845260209384019390920191600101613baa565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613c13613bfd83610471565b92613c0b60405194856103f4565b808452610471565b0136602083013760005b8351811015613cef5780613c33600192866116f2565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613cdb610120826103f4565b519020613ce882856116f2565b5201613c1d565b50909150604051612805816127d9602082018095613b9f565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116116625791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526128056065826103f4565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526128056098826103f4565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526128056058826103f4565b6127d9612805612797612791600060208601511515612db7565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526128056060826103f4565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612805608d826103f456fea2646970667358221220308ea34d0bfd77d22053db540aaaa0e7d8db926aa22e88cb2203a3329f56415764736f6c634300081c0033",
}

// WalletStage1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage1MetaData.ABI instead.
var WalletStage1ABI = WalletStage1MetaData.ABI

// WalletStage1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage1MetaData.Bin instead.
var WalletStage1Bin = WalletStage1MetaData.Bin

// DeployWalletStage1 deploys a new Ethereum contract, binding an instance of WalletStage1 to it.
func DeployWalletStage1(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _entryPoint common.Address) (common.Address, *types.Transaction, *WalletStage1, error) {
	parsed, err := WalletStage1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage1Bin), backend, _factory, _entryPoint)
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

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage1 *WalletStage1Caller) Entrypoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "entrypoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage1 *WalletStage1Session) Entrypoint() (common.Address, error) {
	return _WalletStage1.Contract.Entrypoint(&_WalletStage1.CallOpts)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) Entrypoint() (common.Address, error) {
	return _WalletStage1.Contract.Entrypoint(&_WalletStage1.CallOpts)
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
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage1 *WalletStage1Caller) ReadHook(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "readHook", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage1 *WalletStage1Session) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, selector)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, selector)
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

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Caller) TokenReceived(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "tokenReceived", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1Session) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.CallOpts, arg0, arg1, arg2)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1 *WalletStage1CallerSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.CallOpts, arg0, arg1, arg2)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Transactor) AddHook(opts *bind.TransactOpts, selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "addHook", selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Session) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, selector, implementation)
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

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage1 *WalletStage1Transactor) ExecuteUserOp(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "executeUserOp", _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage1 *WalletStage1Session) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.ExecuteUserOp(&_WalletStage1.TransactOpts, _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletStage1 *WalletStage1TransactorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.ExecuteUserOp(&_WalletStage1.TransactOpts, _payload)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage1 *WalletStage1Transactor) RemoveHook(opts *bind.TransactOpts, selector [4]byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "removeHook", selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage1 *WalletStage1Session) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, selector)
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

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage1 *WalletStage1Transactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "tokensReceived", operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage1 *WalletStage1Session) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokensReceived(&_WalletStage1.TransactOpts, operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletStage1 *WalletStage1TransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokensReceived(&_WalletStage1.TransactOpts, operator, from, to, amount, data, operatorData)
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

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage1 *WalletStage1Transactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage1 *WalletStage1Session) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage1.Contract.ValidateUserOp(&_WalletStage1.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletStage1 *WalletStage1TransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletStage1.Contract.ValidateUserOp(&_WalletStage1.TransactOpts, userOp, userOpHash, missingAccountFunds)
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
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) FilterCallAborted(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage1CallAbortedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallAborted", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallAbortedIterator{contract: _WalletStage1.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage1CallAborted, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallAborted", _opHashRule)
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
// Solidity: event CallAborted(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
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
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) FilterCallFailed(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage1CallFailedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallFailed", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallFailedIterator{contract: _WalletStage1.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
func (_WalletStage1 *WalletStage1Filterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage1CallFailed, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallFailed", _opHashRule)
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
// Solidity: event CallFailed(bytes32 indexed _opHash, uint256 _index, bytes _returnData)
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
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) FilterCallSkipped(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage1CallSkippedIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallSkipped", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallSkippedIterator{contract: _WalletStage1.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage1CallSkipped, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallSkipped", _opHashRule)
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
// Solidity: event CallSkipped(bytes32 indexed _opHash, uint256 _index)
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
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) FilterCallSucceeded(opts *bind.FilterOpts, _opHash [][32]byte) (*WalletStage1CallSucceededIterator, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "CallSucceeded", _opHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletStage1CallSucceededIterator{contract: _WalletStage1.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
func (_WalletStage1 *WalletStage1Filterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletStage1CallSucceeded, _opHash [][32]byte) (event.Subscription, error) {

	var _opHashRule []interface{}
	for _, _opHashItem := range _opHash {
		_opHashRule = append(_opHashRule, _opHashItem)
	}

	logs, sub, err := _WalletStage1.contract.WatchLogs(opts, "CallSucceeded", _opHashRule)
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
// Solidity: event CallSucceeded(bytes32 indexed _opHash, uint256 _index)
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
	Selector       [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletStage1 *WalletStage1Filterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage1DefinedHookIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage1DefinedHookIterator{contract: _WalletStage1.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
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
// Solidity: event DefinedHook(bytes4 selector, address implementation)
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
