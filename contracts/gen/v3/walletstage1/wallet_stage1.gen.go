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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"INIT_CODE_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAGE_2_IMPLEMENTATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureType\",\"inputs\":[{\"name\":\"_type\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x61010080604052346101ae57604081617f1f803803809161002082856101b3565b8339810103126101ae5761003f6020610038836101d6565b92016101d6565b604051909190613dff808201906001600160401b0382118383101761018c57602091839161412083396001600160a01b03861681520301906000f09081156101a257604051606081016001600160401b0381118282101761018c576040908152602f82527f6041600e3d396021805130553df33d3d36153402601f57363d3d373d363d305460208301526e5af43d82803e903d91601f57fd5bf360881b82820152519060005b602f811061017757505030604f820152604f8152610104606f826101b3565b805160209091012060805260a0526001600160a01b031660c05260e052604051613f3590816101eb8239608051818181610c210152613b59015260a051818181610d6f0152613b2a015260c051818181610cd6015261119d015260e0518181816110600152818161120c0152611aa00152f35b806020809284010151828286010152016100e5565b634e487b7160e01b600052604160045260246000fd5b6040513d6000823e3d90fd5b600080fd5b601f909101601f19168101906001600160401b0382119082101761018c57604052565b51906001600160a01b03821682036101ae5756fe6080604052600436101561001e575b361561001c5761001c611c68565b005b60003560e01c806223de29146101ad578063025b22bc146101a857806313792a4a146101a3578063150b7a021461019e5780631626ba7e1461019957806319822f7c146101945780631a9b23371461018f5780631f6a1eb91461018a578063257671f51461018557806329561426146101805780632dd310001461017b5780634fcf3eca146101765780636ea44577146101715780638943ec021461016c5780638c3f55631461016757806392dcb3fc146101625780639c145aed1461015d5780639f69ef5414610158578063a65d69d414610153578063aaf10f421461014e578063ad55366b14610149578063b93ea7ad14610144578063bc197c811461013f578063f23a6e611461013a5763f727ef1c0361000e5761158b565b6114fe565b61142c565b6112cd565b611281565b611230565b6111c1565b611152565b61103a565b610fdc565b610fa0565b610f1c565b610eed565b610d93565b610d24565b610c44565b610beb565b610aff565b610a9c565b6109e7565b61095f565b6108d2565b6107d5565b6102fb565b61026f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101d557565b9181601f840112156101d55782359167ffffffffffffffff83116101d557602083818601950101116101d557565b346101d55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576102a66101b2565b506102af6101da565b506102b86101fd565b5060843567ffffffffffffffff81116101d5576102d9903690600401610241565b505060a43567ffffffffffffffff81116101d55761001c903690600401610241565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55761032d6101b2565b30330361033d5761001c90611d31565b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103b657604052565b61036b565b6040810190811067ffffffffffffffff8211176103b657604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103b657604052565b6040519061042760e0836103d7565b565b60405190610427610120836103d7565b359060ff821682036101d557565b359081151582036101d557565b67ffffffffffffffff81116103b65760051b60200190565b67ffffffffffffffff81116103b657601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104b28261046c565b916104c060405193846103d7565b8294818452818301116101d5578281602093846000960137010152565b9080601f830112156101d5578160206104f8933591016104a6565b90565b81601f820112156101d55780359061051282610454565b9261052060405194856103d7565b82845260208085019360051b830101918183116101d55760208101935b83851061054c57505050505090565b843567ffffffffffffffff81116101d557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101d557610593610418565b916105a060208301610220565b83526040820135602084015260608201359267ffffffffffffffff84116101d55760e0836105d58860208098819801016104dd565b6040840152608081013560608401526105f060a08201610447565b608084015261060160c08201610447565b60a0840152013560c082015281520194019361053d565b9080601f830112156101d557813561062f81610454565b9261063d60405194856103d7565b81845260208085019260051b8201019283116101d557602001905b8282106106655750505090565b6020809161067284610220565b815201910190610658565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101d55760043567ffffffffffffffff81116101d5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101d5576106ef610429565b906106fc81600401610439565b825261070a60248201610447565b6020830152604481013567ffffffffffffffff81116101d557846004610732928401016104fb565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101d55784600461076e928401016104dd565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101d557600485916107ab930101610618565b610100820152916024359067ffffffffffffffff82116101d5576107d191600401610241565b9091565b346101d5576107e33661067d565b90916101008101926107fe6107f9855151611693565b6116b3565b9160005b85518051821015610865579061085f61083a61082083600195611731565b5173ffffffffffffffffffffffffffffffffffffffff1690565b6108448388611731565b9073ffffffffffffffffffffffffffffffffffffffff169052565b01610802565b505083838661087a3361084483515185611731565b52610886818484611d80565b50156108985760405160018152602090f35b6108ce906040519384937ff58cc8b50000000000000000000000000000000000000000000000000000000085526004850161198c565b0390fd5b346101d55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576109096101b2565b506109126101da565b5060643567ffffffffffffffff81116101d557610933903690600401610241565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101d55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043560243567ffffffffffffffff81116101d5576020916109b76109bd923690600401610241565b916119b1565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043567ffffffffffffffff81116101d5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101d557610a6a60209160243560443591600401611a86565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101d557565b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020610ae1600435610adc81610a72565b611f65565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043567ffffffffffffffff81116101d557610b49903690600401610241565b9060243567ffffffffffffffff81116101d557610b6d610b77913690600401610241565b9190935a9361209a565b91610b8b606084015160808501519061235e565b610b96828585611d80565b929015610ba8575061001c93506124b7565b836108ce86926040519384937fa2b6d61b0000000000000000000000000000000000000000000000000000000085526004850161198c565b60009103126101d557565b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043530330361033d578015610cfa576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a161001c7f0000000000000000000000000000000000000000000000000000000000000000611d31565b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557600435610dc981610a72565b30330361033d5773ffffffffffffffffffffffffffffffffffffffff610dee82611f65565b1615610e545760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610e45600082612d0f565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101d5576004359067ffffffffffffffff82116101d5576107d191600401610241565b610ef636610ea4565b9030330361033d57610f0c61001c925a9261209a565b90610f16826127db565b906124b7565b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557610f536101b2565b5060443567ffffffffffffffff81116101d557610f74903690600401610241565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020610a6a60043561286f565b346101d55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760406110186004356128b5565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101d55761104836610ea4565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156111285733036110fa57303b156101d5576110ce9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611c27565b038183305af180156110f5576110e057005b806110ef600061001c936103d7565b80610be0565b6119fd565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101d55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101d55760c061129d6000806112973661067d565b9161297d565b926112a9839293613ac6565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d55760043561130381610a72565b61130b6101da565b9030330361033d5773ffffffffffffffffffffffffffffffffffffffff61133182611f65565b166113ab577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff0000000000000000000000000000000000000000000000000000000060409316911661139e8183612d0f565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101d55782359167ffffffffffffffff83116101d5576020808501948460051b0101116101d557565b346101d55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576114636101b2565b5061146c6101da565b5060443567ffffffffffffffff81116101d55761148d9036906004016113fb565b505060643567ffffffffffffffff81116101d5576114af9036906004016113fb565b505060843567ffffffffffffffff81116101d5576114d1903690600401610241565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101d55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576115356101b2565b5061153e6101da565b5060843567ffffffffffffffff81116101d55761155f903690600401610241565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101d55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101d5576004356115c56101da565b604435916bffffffffffffffffffffffff83168093036101d55730330361033d578273ffffffffffffffffffffffffffffffffffffffff836116507febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612cca565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90600182018092116116a157565b611664565b919082018092116116a157565b906116bd82610454565b6116ca60405191826103d7565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116f88294610454565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b80518210156117455760209160051b010190565b611702565b919082519283825260005b8481106117945750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611755565b9080602083519182815201916020808360051b8301019401926000915b8383106117d557505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611848604085015160e0604086015260e085019061174a565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906117c6565b906020808351928381520192019060005b81811061189d5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611890565b805160ff1682526104f891602082810151151590820152610100611927611901604085015161012060408601526101208501906117a9565b606085015160608501526080850151608085015260a085015184820360a086015261174a565b9260c081015160c084015260e081015160e084015201519061010081840391015261187f565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916119a36104f894926040855260408501906118c9565b92602081850391015261194d565b906119ce92916119bf611f0a565b906003825260e0820152611d80565b50156119f8577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101d5570180359067ffffffffffffffff82116101d5576020019181360383136101d557565b908160209103126101d557516104f881610a72565b6040906104f894928152816020820152019161194d565b91909173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016918215611128578233036110fa5780611bb8575b506020915080610100611ae8920190611a09565b92611b2060405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a6f565b0381305afa9081156110f5577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611b89575b501603611b8457600090565b600190565b611bab915060203d602011611bb1575b611ba381836103d7565b810190611a5a565b38611b78565b503d611b99565b823b156101d5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af19081156110f557602092611ae892611c12575b5090611ad4565b806110ef6000611c21936103d7565b38611c0b565b9160206104f893818152019161194d565b3d15611c63573d90611c498261046c565b91611c5760405193846103d7565b82523d6000602084013e565b606090565b600436108015611c755750565b611cab906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611cfd575b50611f65565b73ffffffffffffffffffffffffffffffffffffffff8116611cc95750565b60008091604051368382378036810184815203915af4611ce7611c38565b9015611cf557602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611ca5565b60207f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039180305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1565b90156117455790565b91907f800000000000000000000000000000000000000000000000000000000000000080611dd7611db18585611d77565b357fff000000000000000000000000000000000000000000000000000000000000001690565b1614611e36576000918291611deb9461297d565b905091909192808210611e06575050611e0390613ac6565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050611e41906127db565b90611e4b826128b5565b42811115611ed8575073ffffffffffffffffffffffffffffffffffffffff81168015159081611ecd575b50611e81575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611e75565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103b6576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152611fdc6060826103d7565b519020541690565b90611fee82610454565b611ffb60405191826103d7565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120298294610454565b019060005b82811061203a57505050565b6020906040516120498161039a565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c08201528282850101520161202e565b909392938483116101d55784116101d5578101920390565b9190916120a5611f0a565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016123475750600060608701525b60076120f460ff831660011c90565b16806122f6575b506010818116036122c157506001925b61211484611fe4565b92604087019384526000905b85821061212f57505050505050565b8281013560f81c9060010190918160018085160361229f575061217330612157838951611731565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60028084161461227f575b600480841614612231575b6008808416146121fe575b6121e76121e160c0856121c08a60806121b78860108060019c9d16149351611731565b51019015159052565b6121d78a60a06121b7886020808716149351611731565b1660061c60031690565b60ff1690565b60c06121f4838951611731565b5101520190612120565b6001916121e7906121e19060c090878101359060200196906060612223878d51611731565b510152959450505050612194565b90612279908481013560e81c9060030161227261225961225184846116a6565b838a8a612082565b91906040612268888d51611731565b51019236916104a6565b90526116a6565b90612189565b90838101359060200191906020612297838951611731565b51015261217e565b6122bc92508481013560601c906014019290612157838951611731565b612173565b6020908116036122e35761ffff918381013560f01c906002015b92169261210b565b60ff918381013560f81c906001016122db565b61233a919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92906080870152386120fb565b80850135606090811c9088015260140192506120e5565b906123688261286f565b8181036123ee57509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8252836040820152604081526123d86060826103d7565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61243a604092959493956060835260608301906118c9565b9460208201520152565b926104f896959260c09592855260208501526040840152606083015260808201528160a0820152019061174a565b6104f8939260609282526020820152816040820152019061174a565b6124a46104f894926060835260608301906118c9565b926020820152604081840391015261174a565b916000604082019384515190825b8281106124d6575b50505050505050565b6124e1818851611731565b51936124f060a0860151151590565b806127d3575b61279357506000936060810151801580158061278a575b6127525784906125206080850151151590565b1561270c576125cb92612547855173ffffffffffffffffffffffffffffffffffffffff1690565b911561270657505a905b6125c68b61259a60608d01516040890151908c8b604051998a967f4c4e814c00000000000000000000000000000000000000000000000000000000602089015260248801612444565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018552846103d7565b612c9d565b1561260e575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016124c5565b60c0018051156126bd57600181511461267e575160021461262f57386125d1565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b925061266f612663612caf565b60405193849384612472565b0390a1388080808080806124cd565b50846108ce61268b612caf565b6040519384937f7f6b0bb10000000000000000000000000000000000000000000000000000000085526004850161248e565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6126fe6126f0612caf565b604051918291858c84612472565b0390a1612608565b90612551565b835161274793925073ffffffffffffffffffffffffffffffffffffffff169160208501519160001461274c57505a905b604085015192612c8b565b6125cb565b9061273c565b83886108ce5a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612422565b50815a1061250d565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b91819081016126fe565b5080156124f6565b61283d6128696127fb6127f5602085015115153090612d4e565b93612e49565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103d7565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83526040820152604081526128ae6060826103d7565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526128f46060826103d7565b51902054906bffffffffffffffffffffffff8260601c921690565b6040519061291c826103bb565b60006020838281520152565b908160409103126101d557602060405191612942836103bb565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff6104f89493168152816020820152019061174a565b909491939291853560f81c6001909381908761299761290f565b9360408089161480612c6d575b612b70575b505050600180861614612b4a5760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612a199060051c90565b612a2290611693565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612a6e846127db565b988993612a7a93612082565b91612a8493613260565b9098612a9891600052602052604060002090565b90612aab91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612ad491600052602052604060002090565b94815190868215159283612b3f575b505081612b30575b50612af35750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612aeb565b141591508638612ae3565b86612b659792949550612b5d9398612082565b93909261303e565b919394909293929190565b8a81013560601c985060140192509087908a90849015612b94575b508291506129a9565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c92604092612c17929091612be291612bdb918990612bd589836116a6565b92612082565b36916104a6565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612950565b0392165afa80156110f557612c3592600091612c3e575b50926116a6565b86388981612b8b565b612c60915060403d604011612c66575b612c5881836103d7565b810190612928565b38612c2e565b503d612c4e565b5073ffffffffffffffffffffffffffffffffffffffff8916156129a4565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612d096060826103d7565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612d096060826103d7565b15612dfe576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261286960c0826103d7565b4690612d57565b805160209091019060005b818110612e1d5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612e10565b610100810151604051612e648161283d602082018095612e05565b51902090612e73815160ff1690565b60ff811680612eec57505090612869612e8f6040840151613bc3565b9261283d60806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612f4a57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612869816080810161283d565b60028103612fa057505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612869816080810161283d565b600303612ff4575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612869816080810161283d565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b9061243a906040939695949660608452606084019161194d565b919492909260009560009560009560009560009561305a611f0a565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b8281106130b1575050505050505080511515806130a3575b612af35750565b50602081015184111561309c565b600381019d50959b50939950919750929091906130d5908b9085013560e81c6116a6565b958a6000848903613219575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c036131e9575090600161311e89613125948789612082565b908b61297d565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106131a35750928b88511461319a575b808b101561316857508a60c08501528992959295949194939093613084565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b60008852613149565b8d8f6108ce6131b485858c8e612082565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601613024565b9798999a9160016131ff8b61320694888a612082565b908661297d565b50929d919c909b929a9092918e8e613139565b916130e4565b73ffffffffffffffffffffffffffffffffffffffff6104f895936060938352166020820152816040820152019161194d565b908160209103126101d5575190565b9391909360009460009460005b81811061327b575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613a5f575060018914613a1f5760028914613862576003891461383357600489146137b2576006891461371257600589146136c457600789146135fd57600889146135a7576009891461347e57600a891461331b577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b9091929394959697506003891697881561346d575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613393918787612082565b6040517f898bd921000000000000000000000000000000000000000000000000000000008152939184916133cb918a60048501611a6f565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156110f55761340e9360009361343a575b5060ff909a168091019a613e8f565b908015613434579061342891600052602052604060002090565b955b939291909361326d565b50613428565b60ff91935061345f9060203d8111613466575b61345781836103d7565b810190613251565b92906133ff565b503d61344d565b8084013560f81c9850600101613330565b90919293949596975060038916978815613596575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816134f6918787612082565b6040517f13792a4a0000000000000000000000000000000000000000000000000000000081529391849161352e918b6004850161198c565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156110f5576135709360009361343a575060ff909a168091019a613e8f565b908015613590579061358a91600052602052604060002090565b9561342a565b5061358a565b8084013560f81c9850600101613493565b9850602087019750949593949293919290918201356135c586613e36565b81146135d5575b61357090613e50565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506135cc565b975090919293949597600f169687156136b2575b602060006136236136909a9b86613cfe565b9c9092918a6040516136668161283d8a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa156110f5576135709060ff6000519a1680910199613d41565b600189019883013560f81c9750613611565b9850602087019750949593949293919290918201358085146136ea575b61357090613df7565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506136e1565b989091929394959662ffffff98506137346121e1600c8416603f9060021c1690565b91821561379e575b600316801561378d575b90819061377190613769908781013560e81c906003019c168c01809c8989612082565b90898b613260565b911115613784575b906135709291613dac565b99820199613779565b50600281019084013560f01c613746565b8482013560f81c925060019091019061373c565b975097613808613815929394959697600f61381d93169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686612082565b908688613260565b9061358a92980198600052602052604060002090565b985096509394929391929091908082013590602001968015613590579061358a91600052602052604060002090565b90919293949596975060038916978815613a0e575b8084013560601c996138d691601401906138969060021c6003166121e1565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9081019061393a60208c6138ec85858b8b612082565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a6f565b0392165afa9081156110f5577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff00000000000000000000000000000000000000000000000000000000916000916139f0575b5016036139ac57509060ff61357092991680910199613d41565b6108ce6139bd8c9389938989612082565b906040519485947fb2fed7ae0000000000000000000000000000000000000000000000000000000086526004860161321f565b613a08915060203d8111611bb157611ba381836103d7565b38613992565b8381013560f81c9850600101613877565b98600f91929394959697985016968715613a4e575b60148101976135709160ff9091169084013560601c613d41565b8281013560f81c9750600101613a34565b98509091929394959698600f16978815613ab1575b5060206000613a876136909a9b86613cfe565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613a74565b73ffffffffffffffffffffffffffffffffffffffff9060405160208101917fff0000000000000000000000000000000000000000000000000000000000000083527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060601b16602183015260358201527f0000000000000000000000000000000000000000000000000000000000000000605582015260558152613b8c6075826103d7565b51902016301490565b805160209091019060005b818110613bad5750505090565b8251845260209384019390920191600101613ba0565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613c09613bf383610454565b92613c0160405194856103d7565b808452610454565b0136602083013760005b8351811015613ce55780613c2960019286611731565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613cd1610120826103d7565b519020613cde8285611731565b5201613c13565b509091506040516128698161283d602082018095613b95565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116116a15791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526128696065826103d7565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526128696098826103d7565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526128696058826103d7565b61283d6128696127fb6127f5600060208601511515612d4e565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526128696060826103d7565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612869608d826103d756fea2646970667358221220825b940dd691f79f4782ba8dca1eaa7ff741f07679c4faf28637d8117499069a64736f6c634300081c003360a034607457601f613dff38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b0381168103607457608052604051613d6f90816100908239608051818181610fe90152818161112601526119ba0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c611b82565b005b60003560e01c806223de291461018d578063025b22bc1461018857806313792a4a14610183578063150b7a021461017e5780631626ba7e1461017957806319822f7c146101745780631a9b23371461016f5780631f6a1eb91461016a57806329561426146101655780634fcf3eca1461016057806351605d801461015b5780636ea44577146101565780638943ec02146101515780638c3f55631461014c57806392dcb3fc146101475780639c145aed14610142578063a65d69d41461013d578063aaf10f4214610138578063ad55366b14610133578063b93ea7ad1461012e578063bc197c8114610129578063f23a6e61146101245763f727ef1c0361000e576114a5565b611418565b611346565b6111e7565b61119b565b61114a565b6110db565b610fc3565b610f65565b610f29565b610ea5565b610e76565b610dd2565b610cb6565b610bfd565b610b1c565b610ab9565b610a04565b61097c565b6108ef565b6107f2565b6102db565b61024f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101b557565b9181601f840112156101b55782359167ffffffffffffffff83116101b557602083818601950101116101b557565b346101b55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610286610192565b5061028f6101ba565b506102986101dd565b5060843567ffffffffffffffff81116101b5576102b9903690600401610221565b505060a43567ffffffffffffffff81116101b55761001c903690600401610221565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761030d610192565b30330361035a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103d357604052565b610388565b6040810190811067ffffffffffffffff8211176103d357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103d357604052565b6040519061044460e0836103f4565b565b60405190610444610120836103f4565b359060ff821682036101b557565b359081151582036101b557565b67ffffffffffffffff81116103d35760051b60200190565b67ffffffffffffffff81116103d357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104cf82610489565b916104dd60405193846103f4565b8294818452818301116101b5578281602093846000960137010152565b9080601f830112156101b557816020610515933591016104c3565b90565b81601f820112156101b55780359061052f82610471565b9261053d60405194856103f4565b82845260208085019360051b830101918183116101b55760208101935b83851061056957505050505090565b843567ffffffffffffffff81116101b557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101b5576105b0610435565b916105bd60208301610200565b83526040820135602084015260608201359267ffffffffffffffff84116101b55760e0836105f28860208098819801016104fa565b60408401526080810135606084015261060d60a08201610464565b608084015261061e60c08201610464565b60a0840152013560c082015281520194019361055a565b9080601f830112156101b557813561064c81610471565b9261065a60405194856103f4565b81845260208085019260051b8201019283116101b557602001905b8282106106825750505090565b6020809161068f84610200565b815201910190610675565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101b55761070c610446565b9061071981600401610456565b825261072760248201610464565b6020830152604481013567ffffffffffffffff81116101b55784600461074f92840101610518565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101b55784600461078b928401016104fa565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101b557600485916107c8930101610635565b610100820152916024359067ffffffffffffffff82116101b5576107ee91600401610221565b9091565b346101b5576108003661069a565b909161010081019261081b6108168551516115ad565b6115cd565b9160005b85518051821015610882579061087c61085761083d8360019561164b565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610861838861164b565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161081f565b5050838386610897336108618351518561164b565b526108a3818484611c54565b50156108b55760405160018152602090f35b6108eb906040519384937ff58cc8b5000000000000000000000000000000000000000000000000000000008552600485016118a6565b0390fd5b346101b55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610926610192565b5061092f6101ba565b5060643567ffffffffffffffff81116101b557610950903690600401610221565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101b55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043560243567ffffffffffffffff81116101b5576020916109d46109da923690600401610221565b916118cb565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101b557610a87602091602435604435916004016119a0565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101b557565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020610afe600435610af981610a8f565b611e39565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043567ffffffffffffffff81116101b557610b66903690600401610221565b9060243567ffffffffffffffff81116101b557610b8a610b94913690600401610221565b9190935a93611f6e565b91610ba86060840151608085015190612232565b610bb3828585611c54565b929015610bc5575061001c935061238b565b836108eb86926040519384937fa2b6d61b000000000000000000000000000000000000000000000000000000008552600485016118a6565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043530330361035a578015610c8c576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557600435610cec81610a8f565b30330361035a5773ffffffffffffffffffffffffffffffffffffffff610d1182611e39565b1615610d775760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d68600082612be3565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101b557565b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101b5576004359067ffffffffffffffff82116101b5576107ee91600401610221565b610e7f36610e2d565b9030330361035a57610e9561001c925a92611f6e565b90610e9f826126af565b9061238b565b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557610edc610192565b5060443567ffffffffffffffff81116101b557610efd903690600401610221565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020610a87600435612743565b346101b55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576040610fa1600435612789565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b346101b557610fd136610e2d565b9073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156110b157330361108357303b156101b5576110579160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611b41565b038183305af1801561107e5761106957005b80611078600061001c936103f4565b80610dc7565b611917565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101b55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101b55760c06111b76000806111b13661069a565b91612851565b926111c383929361399a565b906040519586526020860152151560408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55760043561121d81610a8f565b6112256101ba565b9030330361035a5773ffffffffffffffffffffffffffffffffffffffff61124b82611e39565b166112c5577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166112b88183612be3565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101b55782359167ffffffffffffffff83116101b5576020808501948460051b0101116101b557565b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761137d610192565b506113866101ba565b5060443567ffffffffffffffff81116101b5576113a7903690600401611315565b505060643567ffffffffffffffff81116101b5576113c9903690600401611315565b505060843567ffffffffffffffff81116101b5576113eb903690600401610221565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101b55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b55761144f610192565b506114586101ba565b5060843567ffffffffffffffff81116101b557611479903690600401610221565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101b55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101b5576004356114df6101ba565b604435916bffffffffffffffffffffffff83168093036101b55730330361035a578273ffffffffffffffffffffffffffffffffffffffff8361156a7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612b9e565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90600182018092116115bb57565b61157e565b919082018092116115bb57565b906115d782610471565b6115e460405191826103f4565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116128294610471565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561165f5760209160051b010190565b61161c565b919082519283825260005b8481106116ae5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161166f565b9080602083519182815201916020808360051b8301019401926000915b8383106116ef57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611762604085015160e0604086015260e0850190611664565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906116e0565b906020808351928381520192019060005b8181106117b75750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016117aa565b805160ff1682526105159160208281015115159082015261010061184161181b604085015161012060408601526101208501906116c3565b606085015160608501526080850151608085015260a085015184820360a0860152611664565b9260c081015160c084015260e081015160e0840152015190610100818403910152611799565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b916118bd61051594926040855260408501906117e3565b926020818503910152611867565b906118e892916118d9611dde565b906003825260e0820152611c54565b5015611912577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101b5570180359067ffffffffffffffff82116101b5576020019181360383136101b557565b908160209103126101b5575161051581610a8f565b604090610515949281528160208201520191611867565b91909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169182156110b1578233036110835780611ad2575b506020915080610100611a02920190611923565b92611a3a60405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611989565b0381305afa90811561107e577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611aa3575b501603611a9e57600090565b600190565b611ac5915060203d602011611acb575b611abd81836103f4565b810190611974565b38611a92565b503d611ab3565b823b156101b5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af190811561107e57602092611a0292611b2c575b50906119ee565b806110786000611b3b936103f4565b38611b25565b916020610515938181520191611867565b3d15611b7d573d90611b6382610489565b91611b7160405193846103f4565b82523d6000602084013e565b606090565b600436108015611b8f5750565b611bc5906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611c17575b50611e39565b73ffffffffffffffffffffffffffffffffffffffff8116611be35750565b60008091604051368382378036810184815203915af4611c01611b52565b9015611c0f57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611bbf565b901561165f5790565b91907f800000000000000000000000000000000000000000000000000000000000000080611cab611c858585611c4b565b357fff000000000000000000000000000000000000000000000000000000000000001690565b1614611d0a576000918291611cbf94612851565b905091909192808210611cda575050611cd79061399a565b91565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050611d15906126af565b90611d1f82612789565b42811115611dac575073ffffffffffffffffffffffffffffffffffffffff81168015159081611da1575b50611d55575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611d49565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103d3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152611eb06060826103f4565b519020541690565b90611ec282610471565b611ecf60405191826103f4565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611efd8294610471565b019060005b828110611f0e57505050565b602090604051611f1d816103b7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201611f02565b909392938483116101b55784116101b5578101920390565b919091611f79611dde565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183160161221b5750600060608701525b6007611fc860ff831660011c90565b16806121ca575b5060108181160361219557506001925b611fe884611eb8565b92604087019384526000905b85821061200357505050505050565b8281013560f81c9060010190918160018085160361217357506120473061202b83895161164b565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280841614612153575b600480841614612105575b6008808416146120d2575b6120bb6120b560c0856120948a608061208b8860108060019c9d1614935161164b565b51019015159052565b6120ab8a60a061208b88602080871614935161164b565b1660061c60031690565b60ff1690565b60c06120c883895161164b565b5101520190611ff4565b6001916120bb906120b59060c0908781013590602001969060606120f7878d5161164b565b510152959450505050612068565b9061214d908481013560e81c9060030161214661212d61212584846115c0565b838a8a611f56565b9190604061213c888d5161164b565b51019236916104c3565b90526115c0565b9061205d565b9083810135906020019190602061216b83895161164b565b510152612052565b61219092508481013560601c90601401929061202b83895161164b565b612047565b6020908116036121b75761ffff918381013560f01c906002015b921692611fdf565b60ff918381013560f81c906001016121af565b61220e919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608087015238611fcf565b80850135606090811c908801526014019250611fb9565b9061223c82612743565b8181036122c257509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8252836040820152604081526122ac6060826103f4565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61230e604092959493956060835260608301906117e3565b9460208201520152565b9261051596959260c09592855260208501526040840152606083015260808201528160a08201520190611664565b6105159392606092825260208201528160408201520190611664565b61237861051594926060835260608301906117e3565b9260208201526040818403910152611664565b916000604082019384515190825b8281106123aa575b50505050505050565b6123b581885161164b565b51936123c460a0860151151590565b806126a7575b61266757506000936060810151801580158061265e575b6126265784906123f46080850151151590565b156125e05761249f9261241b855173ffffffffffffffffffffffffffffffffffffffff1690565b91156125da57505a905b61249a8b61246e60608d01516040890151908c8b604051998a967f4c4e814c00000000000000000000000000000000000000000000000000000000602089015260248801612318565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018552846103f4565b612b71565b156124e2575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b01612399565b60c001805115612591576001815114612552575160021461250357386124a5565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250612543612537612b83565b60405193849384612346565b0390a1388080808080806123a1565b50846108eb61255f612b83565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612362565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6125d26125c4612b83565b604051918291858c84612346565b0390a16124dc565b90612425565b835161261b93925073ffffffffffffffffffffffffffffffffffffffff169160208501519160001461262057505a905b604085015192612b5f565b61249f565b90612610565b83886108eb5a6040519384937f21395274000000000000000000000000000000000000000000000000000000008552600485016122f6565b50815a106123e1565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b91819081016125d2565b5080156123ca565b61271161273d6126cf6126c9602085015115153090612c22565b93612d1d565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826103f4565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83526040820152604081526127826060826103f4565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526127c86060826103f4565b51902054906bffffffffffffffffffffffff8260601c921690565b604051906127f0826103d8565b60006020838281520152565b908160409103126101b557602060405191612816836103d8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61051594931681528160208201520190611664565b909491939291853560f81c6001909381908761286b6127e3565b9360408089161480612b41575b612a44575b505050600180861614612a1e5760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019080969181966020166128ed9060051c90565b6128f6906115ad565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612942846126af565b98899361294e93611f56565b9161295893613134565b909861296c91600052602052604060002090565b9061297f91600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff166129a891600052602052604060002090565b94815190868215159283612a13575b505081612a04575b506129c75750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b905060208201511015386129bf565b1415915086386129b7565b86612a399792949550612a319398611f56565b939092612f12565b919394909293929190565b8a81013560601c985060140192509087908a90849015612a68575b5082915061287d565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c92604092612aeb929091612ab691612aaf918990612aa989836115c0565b92611f56565b36916104c3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612824565b0392165afa801561107e57612b0992600091612b12575b50926115c0565b86388981612a5f565b612b34915060403d604011612b3a575b612b2c81836103f4565b8101906127fc565b38612b02565b503d612b22565b5073ffffffffffffffffffffffffffffffffffffffff891615612878565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612bdd6060826103f4565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612bdd6060826103f4565b15612cd2576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261273d60c0826103f4565b4690612c2b565b805160209091019060005b818110612cf15750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612ce4565b610100810151604051612d3881612711602082018095612cd9565b51902090612d47815160ff1690565b60ff811680612dc05750509061273d612d6360408401516139fd565b9261271160806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612e1e57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4669381019384529081019190915260608101929092529061273d8160808101612711565b60028103612e7457505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e46020820190815291810192909252606082019290925261273d8160808101612711565b600303612ec8575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4666020820190815291810192909252606082019290925261273d8160808101612711565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b9061230e9060409396959496606084526060840191611867565b9194929092600095600095600095600095600095612f2e611dde565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b828110612f8557505050505050508051151580612f77575b6129c75750565b506020810151841115612f70565b600381019d50959b5093995091975092909190612fa9908b9085013560e81c6115c0565b958a60008489036130ed575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c036130bd5750906001612ff289612ff9948789611f56565b908b612851565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b106130775750928b88511461306e575b808b101561303c57508a60c08501528992959295949194939093612f58565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b6000885261301d565b8d8f6108eb61308885858c8e611f56565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601612ef8565b9798999a9160016130d38b6130da94888a611f56565b9086612851565b50929d919c909b929a9092918e8e61300d565b91612fb8565b73ffffffffffffffffffffffffffffffffffffffff610515959360609383521660208201528160408201520191611867565b908160209103126101b5575190565b9391909360009460009460005b81811061314f575050505050565b8481013560f881901c9860019092019788979692909160fc1c9889156139335750600189146138f357600289146137365760038914613707576004891461368657600689146135e6576005891461359857600789146134d1576008891461347b576009891461335257600a89146131ef577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613341575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613267918787611f56565b6040517f898bd9210000000000000000000000000000000000000000000000000000000081529391849161329f918a60048501611989565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561107e576132e29360009361330e575b5060ff909a168091019a613cc9565b90801561330857906132fc91600052602052604060002090565b955b9392919093613141565b506132fc565b60ff9193506133339060203d811161333a575b61332b81836103f4565b810190613125565b92906132d3565b503d613321565b8084013560f81c9850600101613204565b9091929394959697506003891697881561346a575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0116910190810190816133ca918787611f56565b6040517f13792a4a00000000000000000000000000000000000000000000000000000000815293918491613402918b600485016118a6565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561107e576134449360009361330e575060ff909a168091019a613cc9565b908015613464579061345e91600052602052604060002090565b956132fe565b5061345e565b8084013560f81c9850600101613367565b98506020870197509495939492939192909182013561349986613c70565b81146134a9575b61344490613c8a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506134a0565b975090919293949597600f16968715613586575b602060006134f76135649a9b86613b38565b9c9092918a60405161353a816127118a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa1561107e576134449060ff6000519a1680910199613b7b565b600189019883013560f81c97506134e5565b9850602087019750949593949293919290918201358085146135be575b61344490613c31565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff98506135b5565b989091929394959662ffffff98506136086120b5600c8416603f9060021c1690565b918215613672575b6003168015613661575b9081906136459061363d908781013560e81c906003019c168c01809c8989611f56565b90898b613134565b911115613658575b906134449291613be6565b9982019961364d565b50600281019084013560f01c61361a565b8482013560f81c9250600190910190613610565b9750976136dc6136e9929394959697600f6136f193169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686611f56565b908688613134565b9061345e92980198600052602052604060002090565b985096509394929391929091908082013590602001968015613464579061345e91600052602052604060002090565b909192939495969750600389169788156138e2575b8084013560601c996137aa916014019061376a9060021c6003166120b5565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9081019061380e60208c6137c085858b8b611f56565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611989565b0392165afa90811561107e577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff00000000000000000000000000000000000000000000000000000000916000916138c4575b50160361388057509060ff61344492991680910199613b7b565b6108eb6138918c9389938989611f56565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016130f3565b6138dc915060203d8111611acb57611abd81836103f4565b38613866565b8381013560f81c985060010161374b565b98600f91929394959697985016968715613922575b60148101976134449160ff9091169084013560601c613b7b565b8281013560f81c9750600101613908565b98509091929394959698600f16978815613985575b506020600061395b6135649a9b86613b38565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613948565b80151590816139a7575090565b90507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b805160209091019060005b8181106139e75750505090565b82518452602093840193909201916001016139da565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613a43613a2d83610471565b92613a3b60405194856103f4565b808452610471565b0136602083013760005b8351811015613b1f5780613a636001928661164b565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613b0b610120826103f4565b519020613b18828561164b565b5201613a4d565b5090915060405161273d816127116020820180956139cf565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116115bb5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b16603183015260458201526045815261273d6065826103f4565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a000000000000000085526038840152605883015260788201526078815261273d6098826103f4565b60405160208101917f53657175656e636520737461746963206469676573743a0a0000000000000000835260388201526038815261273d6058826103f4565b61271161273d6126cf6126c9600060208601511515612c22565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a835260408201526040815261273d6060826103f4565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d815261273d608d826103f456fea264697066735822122071c846bfb3a07ed1d03bb0cecee955d8b864be4c4d91b0f9a960528ecc060e8964736f6c634300081c0033",
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
