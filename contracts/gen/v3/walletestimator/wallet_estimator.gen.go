// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletestimator

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

// WalletEstimatorMetaData contains all meta data concerning the WalletEstimator contract.
var WalletEstimatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimate\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChainedSignatureNestedInChainedSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidPackedLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a034607457601f6140ac38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b038116810360745760805260405161401c908161009082396080518181816111fc015281816113630152611bbc0152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c611d4e565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc14610157578063975befdb146101525780639c145aed1461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e576116d9565b61164c565b61157a565b61141b565b6113d8565b611387565b611318565b611189565b61107b565b61101d565b610fe1565b610f5d565b610f2e565b610e8a565b610d6e565b610cb5565b610b98565b610ac9565b610a14565b61098c565b6108ff565b610802565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b81601f820112156101c55780359061053f82610481565b9261054d6040519485610404565b82845260208085019360051b830101918183116101c55760208101935b83851061057957505050505090565b843567ffffffffffffffff81116101c557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101c5576105c0610445565b916105cd60208301610210565b83526040820135602084015260608201359267ffffffffffffffff84116101c55760e08361060288602080988198010161050a565b60408401526080810135606084015261061d60a08201610474565b608084015261062e60c08201610474565b60a0840152013560c082015281520194019361056a565b9080601f830112156101c557813561065c81610481565b9261066a6040519485610404565b81845260208085019260051b8201019283116101c557602001905b8282106106925750505090565b6020809161069f84610210565b815201910190610685565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c55761071c610456565b9061072981600401610466565b825261073760248201610474565b6020830152604481013567ffffffffffffffff81116101c55784600461075f92840101610528565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c55784600461079b9284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107d8930101610645565b610100820152916024359067ffffffffffffffff82116101c5576107fe91600401610231565b9091565b346101c557610810366106aa565b909161010081019261082b6108268551516117e1565b6117fc565b9160005b85518051821015610892579061088c61086761084d8360019561187a565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610871838861187a565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161082f565b50508383866108a7336108718351518561187a565b526108b3818484611e20565b50156108c55760405160018152602090f35b6108fb906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611ad5565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576109366101a2565b5061093f6101ca565b5060643567ffffffffffffffff81116101c557610960903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109e46109ea923690600401610231565b91611afa565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a9760209160243560443591600401611ba3565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610b0e600435610b0981610a9f565b612027565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c55760043567ffffffffffffffff81116101c55781610b7591600401610231565b929092916024359067ffffffffffffffff82116101c5576107fe91600401610231565b610ba136610b2c565b90929160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b57610bfc9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde555a9361215c565b91610c10606084015160808501519061244b565b610c1b828585611e20565b929015610c535750610c2d93506125a4565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b836108fb86926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611ad5565b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610d44576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610da481610a9f565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610dc982612027565b1615610e2f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610e20600082612e8a565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107fe91600401610231565b610f3736610ee5565b9030330361036a57610f4d61001c925a9261215c565b90610f57826128c8565b906125a4565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610f946101a2565b5060443567ffffffffffffffff81116101c557610fb5903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a9760043561295c565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760406110596004356129a2565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b61108436610b2c565b909160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b576110de9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde555a9461215c565b916110f660608401516110f08161295c565b9061244b565b611101828285611e20565b92901561115257846111148585836125a4565b5a810390811161114d5760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55604051908152602090f35b6117b2565b6108fb84916040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611ad5565b346101c55761119736610ee5565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b5760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156112ee5733036112c057303b156101c55761126a9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611d0d565b038183305af180156112bb576112a05760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b806112af60006112b593610404565b80610e7f565b38610c2d565b611b46565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c06113f46000806113ee366106aa565b91612aa3565b926040929192519485526020850152600160408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561145181610a9f565b6114596101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff61147f82612027565b166114f9577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166114ec8183612e8a565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115b16101a2565b506115ba6101ca565b5060443567ffffffffffffffff81116101c5576115db903690600401611549565b505060643567ffffffffffffffff81116101c5576115fd903690600401611549565b505060843567ffffffffffffffff81116101c55761161f903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576116836101a2565b5061168c6101ca565b5060843567ffffffffffffffff81116101c5576116ad903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356117136101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361179e7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612e45565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161114d57565b9190820180921161114d57565b9061180682610481565b6118136040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118418294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561188e5760209160051b010190565b61184b565b919082519283825260005b8481106118dd5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161189e565b9080602083519182815201916020808360051b8301019401926000915b83831061191e57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611991604085015160e0604086015260e0850190611893565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061190f565b906020808351928381520192019060005b8181106119e65750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016119d9565b805160ff16825261052591602082810151151590820152610100611a70611a4a604085015161012060408601526101208501906118f2565b606085015160608501526080850151608085015260a085015184820360a0860152611893565b9260c081015160c084015260e081015160e08401520151906101008184039101526119c8565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611aec6105259492604085526040850190611a12565b926020818503910152611a96565b90611b179291611b08611fcc565b906003825260e0820152611e20565b5015611b41577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169081156112ee578133036112c05780611c57575b5050611c477f1626ba7e00000000000000000000000000000000000000000000000000000000926109e4836101007fffffffff00000000000000000000000000000000000000000000000000000000950190611b52565b1603611c5257600090565b600190565b813b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af19283156112bb576109e47fffffffff0000000000000000000000000000000000000000000000000000000093611c47937f1626ba7e0000000000000000000000000000000000000000000000000000000096611cf8575b5093505092611bf0565b806112af6000611d0793610404565b38611cee565b916020610525938181520191611a96565b3d15611d49573d90611d2f82610499565b91611d3d6040519384610404565b82523d6000602084013e565b606090565b600436108015611d5b5750565b611d91906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611de3575b50612027565b73ffffffffffffffffffffffffffffffffffffffff8116611daf5750565b60008091604051368382378036810184815203915af4611dcd611d1e565b9015611ddb57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611d8b565b901561188e5790565b611e53611e2d8484611e17565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f800000000000000000000000000000000000000000000000000000000000000080821614611ecd5750611e8b926000928392612aa3565b91505091808210611e9d575050600191565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0200000000000000000000000000000000000000000000000000000000000000908116146020820152611f03925090506128c8565b90611f0d826129a2565b42811115611f9a575073ffffffffffffffffffffffffffffffffffffffff81168015159081611f8f575b50611f43575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611f37565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a12084521660408201526040815261209e606082610404565b519020541690565b906120b082610481565b6120bd6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120eb8294610481565b019060005b8281106120fc57505050565b60209060405161210b816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c0820152828285010152016120f0565b909392938483116101c55784116101c5578101920390565b90612165611fcc565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016124345750600060608601525b60076121b460ff831660011c90565b16806123e3575b506010818116036123b5575060015b6121d3816120a6565b604086019081526000925b82841061221b5750505050036121f15790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361239357506122653061224984875161187a565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280821614612373575b600480821614612325575b6008808216146122f0575b906122d76122d160c0846122b1601080600198161460806122a8888c5161187a565b51019015159052565b6122c760208083161460a06122a8888c5161187a565b1660061c60031690565b60ff1690565b60c06122e483875161187a565b510152019291906121de565b94600191906122d7906122d19060c090868101359060200199906060612317878b5161187a565b510152939450505050612286565b9461236d908381013560e81c9060030161236661234d61234584846117ef565b838c89612144565b9190604061235c888b5161187a565b51019236916104d3565b90526117ef565b9461227b565b9482810135906020019590602061238b84875161187a565b510152612270565b6123b096508381013560601c90601401969061224984875161187a565b612265565b6020908116036123d257600282019181013560f01c905b906121ca565b600182019181013560f81c906123cc565b612427919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92906080860152386121bb565b80830135606090811c9087015260140192506121a5565b906124558261295c565b8181036124db57509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8252836040820152604081526124c5606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61252760409295949395606083526060830190611a12565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611893565b6105259392606092825260208201528160408201520190611893565b6125916105259492606083526060830190611a12565b9260208201526040818403910152611893565b916000604082019384515190825b8281106125c3575b50505050505050565b6125ce81885161187a565b51936125dd60a0860151151590565b806128c0575b612880575060009360608101518015801580612877575b61283f57849061260d6080850151151590565b156127f9576126b892612634855173ffffffffffffffffffffffffffffffffffffffff1690565b91156127f357505a905b6126b38b61268760608d01516040890151908c8b604051998a967f4c4e814c00000000000000000000000000000000000000000000000000000000602089015260248801612531565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612e18565b156126fb575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016125b2565b60c0018051156127aa57600181511461276b575160021461271c57386126be565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b925061275c612750612e2a565b6040519384938461255f565b0390a1388080808080806125ba565b50846108fb612778612e2a565b6040519384937f7f6b0bb10000000000000000000000000000000000000000000000000000000085526004850161257b565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6127eb6127dd612e2a565b604051918291858c8461255f565b0390a16126f5565b9061263e565b835161283493925073ffffffffffffffffffffffffffffffffffffffff169160208501519160001461283957505a905b604085015192612e06565b6126b8565b90612829565b83886108fb5a6040519384937f213952740000000000000000000000000000000000000000000000000000000085526004850161250f565b50815a106125fa565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b91819081016127eb565b5080156125e3565b61292a6129566128e86128e2602085015115153090612ec9565b93612fc4565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e835260408201526040815261299b606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526129e1606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b60405190612a09826103e8565b60006020838281520152565b60031115612a1f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160409103126101c557602060405191612a68836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611893565b909491939291853560f81c600190938190612abc6129fc565b92612ac682612a15565b60018203612cd0575b50600180871614612c6f575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612b3e9060051c90565b612b47906117e1565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612b93846128c8565b988993612b9f93612144565b91612ba993613406565b9098612bbd91600052602052604060002090565b90612bd091600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612bf991600052602052604060002090565b94815190868215159283612c64575b505081612c55575b50612c185750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612c10565b141591508638612c08565b909691939450612c8181989398612a15565b612ca657612c9b9581612c9393612144565b9390926131b9565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612acf578981013560601c9750601401915087878a84612d1985612a15565b60028503612d2a575b505050612acf565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c92604092612daf929091612d7a91612d73918a90612d6d89836117ef565b92612144565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612a76565b0392165afa80156112bb57612dcd92600091612dd7575b50936117ef565b9087388a81612d22565b612df9915060403d604011612dff575b612df18183610404565b810190612a4e565b38612dc6565b503d612de7565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612e84606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612e84606082610404565b15612f79576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261295660c082610404565b4690612ed2565b805160209091019060005b818110612f985750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612f8b565b610100810151604051612fdf8161292a602082018095612f80565b51902090612fee815160ff1690565b60ff8116806130675750509061295661300a6040840151613caa565b9261292a60806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b600181036130c557505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612956816080810161292a565b6002810361311b57505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612956816080810161292a565b60030361316f575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612956816080810161292a565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906125279060409396959496606084526060840191611a96565b91949290926000956000956000956000956000956131d5611fcc565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061322c5750505050505050805115158061321e575b612c185750565b506020810151841115613217565b600381019d50959b5093995091975092909190613250908b9085013560e81c6117ef565b95828703613390578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c03613361575061329d91613296898c938789612144565b908b612aa3565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b1061331b5750928b885114613312575b808b10156132e057508a60c085015289929592959491949390936131ff565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526132c1565b8d8f6108fb61332c85858c8e612144565b9390926040519485947fb006aba00000000000000000000000000000000000000000000000000000000086526004860161319f565b979899809b926133768b61337d94888a612144565b9086612aa3565b50929d919c909b929a9092918e8e6132b1565b8a60029161325d565b908160209103126101c5575161052581610a9f565b604090610525949281528160208201520191611a96565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611a96565b908160209103126101c5575190565b9391909360009460009460005b818110613421575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613c15575060018914613bd55760028914613a0857600389146139d9576004891461395857600689146138b8576005891461386a57600789146137a3576008891461374d576009891461362457600a89146134c1577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613613575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613539918787612144565b6040517f898bd92100000000000000000000000000000000000000000000000000000000815293918491613571918a600485016133ae565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156112bb576135b4936000936135e0575b5060ff909a168091019a613f76565b9080156135da57906135ce91600052602052604060002090565b955b9392919093613413565b506135ce565b60ff9193506136059060203d811161360c575b6135fd8183610404565b8101906133f7565b92906135a5565b503d6135f3565b8084013560f81c98506001016134d6565b9091929394959697506003891697881561373c575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161369c918787612144565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916136d4918b60048501611ad5565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156112bb57613716936000936135e0575060ff909a168091019a613f76565b908015613736579061373091600052602052604060002090565b956135d0565b50613730565b8084013560f81c9850600101613639565b98506020870197509495939492939192909182013561376b86613f1d565b811461377b575b61371690613f37565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613772565b975090919293949597600f16968715613858575b602060006137c96138369a9b86613de5565b9c9092918a60405161380c8161292a8a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa156112bb576137169060ff6000519a1680910199613e28565b600189019883013560f81c97506137b7565b985060208701975094959394929391929091820135808514613890575b61371690613ede565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613887565b989091929394959662ffffff98506138da6122d1600c8416603f9060021c1690565b918215613944575b6003168015613933575b9081906139179061390f908781013560e81c906003019c168c01809c8989612144565b90898b613406565b91111561392a575b906137169291613e93565b9982019961391f565b50600281019084013560f01c6138ec565b8482013560f81c92506001909101906138e2565b9750976139ae6139bb929394959697600f6139c393169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686612144565b908688613406565b9061373092980198600052602052604060002090565b985096509394929391929091908082013590602001968015613736579061373091600052602052604060002090565b90919293949596975060038916978815613bc4575b8084013560601c99613a7c9160140190613a3c9060021c6003166122d1565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613ae060208c613a9285858b8b612144565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e000000000000000000000000000000000000000000000000000000008552600485016133ae565b0392165afa9081156112bb577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613b96575b501603613b5257509060ff61371692991680910199613e28565b6108fb613b638c9389938989612144565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016133c5565b613bb7915060203d8111613bbd575b613baf8183610404565b810190613399565b38613b38565b503d613ba5565b8381013560f81c9850600101613a1d565b98600f91929394959697985016968715613c04575b60148101976137169160ff9091169084013560601c613e28565b8281013560f81c9750600101613bea565b98509091929394959698600f16978815613c67575b5060206000613c3d6138369a9b86613de5565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613c2a565b805160209091019060005b818110613c945750505090565b8251845260209384019390920191600101613c87565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613cf0613cda83610481565b92613ce86040519485610404565b808452610481565b0136602083013760005b8351811015613dcc5780613d106001928661187a565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613db861012082610404565b519020613dc5828561187a565b5201613cfa565b509091506040516129568161292a602082018095613c7c565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff811161114d5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612956606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612956609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612956605882610404565b61292a6129566128e86128e2600060208601511515612ec9565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612956606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612956608d8261040456fea2646970667358221220f32f6ee7ce70c5af164bab9106b12e6b2b298f8aa56010b00e260d8ea76b907464736f6c634300081c0033",
}

// WalletEstimatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletEstimatorMetaData.ABI instead.
var WalletEstimatorABI = WalletEstimatorMetaData.ABI

// WalletEstimatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletEstimatorMetaData.Bin instead.
var WalletEstimatorBin = WalletEstimatorMetaData.Bin

// DeployWalletEstimator deploys a new Ethereum contract, binding an instance of WalletEstimator to it.
func DeployWalletEstimator(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *WalletEstimator, error) {
	parsed, err := WalletEstimatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletEstimatorBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletEstimator{WalletEstimatorCaller: WalletEstimatorCaller{contract: contract}, WalletEstimatorTransactor: WalletEstimatorTransactor{contract: contract}, WalletEstimatorFilterer: WalletEstimatorFilterer{contract: contract}}, nil
}

// WalletEstimator is an auto generated Go binding around an Ethereum contract.
type WalletEstimator struct {
	WalletEstimatorCaller     // Read-only binding to the contract
	WalletEstimatorTransactor // Write-only binding to the contract
	WalletEstimatorFilterer   // Log filterer for contract events
}

// WalletEstimatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletEstimatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletEstimatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletEstimatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletEstimatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletEstimatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletEstimatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletEstimatorSession struct {
	Contract     *WalletEstimator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletEstimatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletEstimatorCallerSession struct {
	Contract *WalletEstimatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WalletEstimatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletEstimatorTransactorSession struct {
	Contract     *WalletEstimatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WalletEstimatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletEstimatorRaw struct {
	Contract *WalletEstimator // Generic contract binding to access the raw methods on
}

// WalletEstimatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletEstimatorCallerRaw struct {
	Contract *WalletEstimatorCaller // Generic read-only contract binding to access the raw methods on
}

// WalletEstimatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletEstimatorTransactorRaw struct {
	Contract *WalletEstimatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletEstimator creates a new instance of WalletEstimator, bound to a specific deployed contract.
func NewWalletEstimator(address common.Address, backend bind.ContractBackend) (*WalletEstimator, error) {
	contract, err := bindWalletEstimator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletEstimator{WalletEstimatorCaller: WalletEstimatorCaller{contract: contract}, WalletEstimatorTransactor: WalletEstimatorTransactor{contract: contract}, WalletEstimatorFilterer: WalletEstimatorFilterer{contract: contract}}, nil
}

// NewWalletEstimatorCaller creates a new read-only instance of WalletEstimator, bound to a specific deployed contract.
func NewWalletEstimatorCaller(address common.Address, caller bind.ContractCaller) (*WalletEstimatorCaller, error) {
	contract, err := bindWalletEstimator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCaller{contract: contract}, nil
}

// NewWalletEstimatorTransactor creates a new write-only instance of WalletEstimator, bound to a specific deployed contract.
func NewWalletEstimatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletEstimatorTransactor, error) {
	contract, err := bindWalletEstimator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorTransactor{contract: contract}, nil
}

// NewWalletEstimatorFilterer creates a new log filterer instance of WalletEstimator, bound to a specific deployed contract.
func NewWalletEstimatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletEstimatorFilterer, error) {
	contract, err := bindWalletEstimator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorFilterer{contract: contract}, nil
}

// bindWalletEstimator binds a generic wrapper to an already deployed contract.
func bindWalletEstimator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletEstimatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletEstimator *WalletEstimatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletEstimator.Contract.WalletEstimatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletEstimator *WalletEstimatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletEstimator.Contract.WalletEstimatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletEstimator *WalletEstimatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletEstimator.Contract.WalletEstimatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletEstimator *WalletEstimatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletEstimator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletEstimator *WalletEstimatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletEstimator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletEstimator *WalletEstimatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletEstimator.Contract.contract.Transact(opts, method, params...)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletEstimator *WalletEstimatorCaller) Entrypoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "entrypoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletEstimator *WalletEstimatorSession) Entrypoint() (common.Address, error) {
	return _WalletEstimator.Contract.Entrypoint(&_WalletEstimator.CallOpts)
}

// Entrypoint is a free data retrieval call binding the contract method 0xa65d69d4.
//
// Solidity: function entrypoint() view returns(address)
func (_WalletEstimator *WalletEstimatorCallerSession) Entrypoint() (common.Address, error) {
	return _WalletEstimator.Contract.Entrypoint(&_WalletEstimator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletEstimator *WalletEstimatorCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletEstimator *WalletEstimatorSession) GetImplementation() (common.Address, error) {
	return _WalletEstimator.Contract.GetImplementation(&_WalletEstimator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletEstimator *WalletEstimatorCallerSession) GetImplementation() (common.Address, error) {
	return _WalletEstimator.Contract.GetImplementation(&_WalletEstimator.CallOpts)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletEstimator *WalletEstimatorCaller) GetStaticSignature(opts *bind.CallOpts, _hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "getStaticSignature", _hash)

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
func (_WalletEstimator *WalletEstimatorSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletEstimator.Contract.GetStaticSignature(&_WalletEstimator.CallOpts, _hash)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletEstimator *WalletEstimatorCallerSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletEstimator.Contract.GetStaticSignature(&_WalletEstimator.CallOpts, _hash)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCaller) ImageHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletEstimator *WalletEstimatorSession) ImageHash() ([32]byte, error) {
	return _WalletEstimator.Contract.ImageHash(&_WalletEstimator.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCallerSession) ImageHash() ([32]byte, error) {
	return _WalletEstimator.Contract.ImageHash(&_WalletEstimator.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletEstimator *WalletEstimatorCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletEstimator *WalletEstimatorSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.IsValidSignature(&_WalletEstimator.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletEstimator *WalletEstimatorCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.IsValidSignature(&_WalletEstimator.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC1155BatchReceived(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC1155BatchReceived(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC1155Received(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC1155Received(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC721Received(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.OnERC721Received(&_WalletEstimator.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletEstimator *WalletEstimatorCaller) ReadHook(opts *bind.CallOpts, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "readHook", selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletEstimator *WalletEstimatorSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletEstimator.Contract.ReadHook(&_WalletEstimator.CallOpts, selector)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 selector) view returns(address)
func (_WalletEstimator *WalletEstimatorCallerSession) ReadHook(selector [4]byte) (common.Address, error) {
	return _WalletEstimator.Contract.ReadHook(&_WalletEstimator.CallOpts, selector)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletEstimator *WalletEstimatorCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletEstimator *WalletEstimatorSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletEstimator.Contract.ReadNonce(&_WalletEstimator.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletEstimator *WalletEstimatorCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletEstimator.Contract.ReadNonce(&_WalletEstimator.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletEstimator *WalletEstimatorCaller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletEstimator *WalletEstimatorSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletEstimator.Contract.RecoverPartialSignature(&_WalletEstimator.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletEstimator *WalletEstimatorCallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletEstimator.Contract.RecoverPartialSignature(&_WalletEstimator.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCaller) RecoverSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "recoverSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletEstimator.Contract.RecoverSapientSignature(&_WalletEstimator.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCallerSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletEstimator.Contract.RecoverSapientSignature(&_WalletEstimator.CallOpts, _payload, _signature)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCaller) TokenReceived(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "tokenReceived", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.TokenReceived(&_WalletEstimator.CallOpts, arg0, arg1, arg2)
}

// TokenReceived is a free data retrieval call binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) pure returns(bytes4)
func (_WalletEstimator *WalletEstimatorCallerSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) ([4]byte, error) {
	return _WalletEstimator.Contract.TokenReceived(&_WalletEstimator.CallOpts, arg0, arg1, arg2)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) AddHook(opts *bind.TransactOpts, selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "addHook", selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.AddHook(&_WalletEstimator.TransactOpts, selector, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 selector, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) AddHook(selector [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.AddHook(&_WalletEstimator.TransactOpts, selector, implementation)
}

// Estimate is a paid mutator transaction binding the contract method 0x975befdb.
//
// Solidity: function estimate(bytes _payload, bytes _signature) payable returns(uint256 gasUsed)
func (_WalletEstimator *WalletEstimatorTransactor) Estimate(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "estimate", _payload, _signature)
}

// Estimate is a paid mutator transaction binding the contract method 0x975befdb.
//
// Solidity: function estimate(bytes _payload, bytes _signature) payable returns(uint256 gasUsed)
func (_WalletEstimator *WalletEstimatorSession) Estimate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Estimate(&_WalletEstimator.TransactOpts, _payload, _signature)
}

// Estimate is a paid mutator transaction binding the contract method 0x975befdb.
//
// Solidity: function estimate(bytes _payload, bytes _signature) payable returns(uint256 gasUsed)
func (_WalletEstimator *WalletEstimatorTransactorSession) Estimate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Estimate(&_WalletEstimator.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletEstimator *WalletEstimatorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Execute(&_WalletEstimator.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Execute(&_WalletEstimator.TransactOpts, _payload, _signature)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletEstimator *WalletEstimatorTransactor) ExecuteUserOp(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "executeUserOp", _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletEstimator *WalletEstimatorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.ExecuteUserOp(&_WalletEstimator.TransactOpts, _payload)
}

// ExecuteUserOp is a paid mutator transaction binding the contract method 0x9c145aed.
//
// Solidity: function executeUserOp(bytes _payload) returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) ExecuteUserOp(_payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.ExecuteUserOp(&_WalletEstimator.TransactOpts, _payload)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) RemoveHook(opts *bind.TransactOpts, selector [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "removeHook", selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletEstimator *WalletEstimatorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.RemoveHook(&_WalletEstimator.TransactOpts, selector)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 selector) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) RemoveHook(selector [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.RemoveHook(&_WalletEstimator.TransactOpts, selector)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletEstimator *WalletEstimatorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.SelfExecute(&_WalletEstimator.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.SelfExecute(&_WalletEstimator.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletEstimator *WalletEstimatorTransactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletEstimator *WalletEstimatorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.Contract.SetStaticSignature(&_WalletEstimator.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.Contract.SetStaticSignature(&_WalletEstimator.TransactOpts, _hash, _address, _timestamp)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletEstimator *WalletEstimatorTransactor) TokensReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "tokensReceived", operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletEstimator *WalletEstimatorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.TokensReceived(&_WalletEstimator.TransactOpts, operator, from, to, amount, data, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address operator, address from, address to, uint256 amount, bytes data, bytes operatorData) returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) TokensReceived(operator common.Address, from common.Address, to common.Address, amount *big.Int, data []byte, operatorData []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.TokensReceived(&_WalletEstimator.TransactOpts, operator, from, to, amount, data, operatorData)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletEstimator *WalletEstimatorTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletEstimator *WalletEstimatorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.UpdateImageHash(&_WalletEstimator.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.UpdateImageHash(&_WalletEstimator.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletEstimator *WalletEstimatorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.UpdateImplementation(&_WalletEstimator.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.UpdateImplementation(&_WalletEstimator.TransactOpts, _implementation)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletEstimator *WalletEstimatorTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletEstimator *WalletEstimatorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.Contract.ValidateUserOp(&_WalletEstimator.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x19822f7c.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_WalletEstimator *WalletEstimatorTransactorSession) ValidateUserOp(userOp PackedUserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _WalletEstimator.Contract.ValidateUserOp(&_WalletEstimator.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletEstimator *WalletEstimatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Fallback(&_WalletEstimator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.Fallback(&_WalletEstimator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletEstimator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletEstimator *WalletEstimatorSession) Receive() (*types.Transaction, error) {
	return _WalletEstimator.Contract.Receive(&_WalletEstimator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletEstimator.Contract.Receive(&_WalletEstimator.TransactOpts)
}

// WalletEstimatorCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletEstimator contract.
type WalletEstimatorCallAbortedIterator struct {
	Event *WalletEstimatorCallAborted // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorCallAborted)
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
		it.Event = new(WalletEstimatorCallAborted)
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
func (it *WalletEstimatorCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorCallAborted represents a CallAborted event raised by the WalletEstimator contract.
type WalletEstimatorCallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletEstimatorCallAbortedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallAbortedIterator{contract: _WalletEstimator.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletEstimator *WalletEstimatorFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletEstimatorCallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorCallAborted)
				if err := _WalletEstimator.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseCallAborted(log types.Log) (*WalletEstimatorCallAborted, error) {
	event := new(WalletEstimatorCallAborted)
	if err := _WalletEstimator.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletEstimator contract.
type WalletEstimatorCallFailedIterator struct {
	Event *WalletEstimatorCallFailed // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorCallFailed)
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
		it.Event = new(WalletEstimatorCallFailed)
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
func (it *WalletEstimatorCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorCallFailed represents a CallFailed event raised by the WalletEstimator contract.
type WalletEstimatorCallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletEstimatorCallFailedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallFailedIterator{contract: _WalletEstimator.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletEstimator *WalletEstimatorFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletEstimatorCallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorCallFailed)
				if err := _WalletEstimator.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseCallFailed(log types.Log) (*WalletEstimatorCallFailed, error) {
	event := new(WalletEstimatorCallFailed)
	if err := _WalletEstimator.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletEstimator contract.
type WalletEstimatorCallSkippedIterator struct {
	Event *WalletEstimatorCallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorCallSkipped)
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
		it.Event = new(WalletEstimatorCallSkipped)
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
func (it *WalletEstimatorCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorCallSkipped represents a CallSkipped event raised by the WalletEstimator contract.
type WalletEstimatorCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletEstimatorCallSkippedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallSkippedIterator{contract: _WalletEstimator.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletEstimatorCallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorCallSkipped)
				if err := _WalletEstimator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseCallSkipped(log types.Log) (*WalletEstimatorCallSkipped, error) {
	event := new(WalletEstimatorCallSkipped)
	if err := _WalletEstimator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletEstimator contract.
type WalletEstimatorCallSucceededIterator struct {
	Event *WalletEstimatorCallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorCallSucceeded)
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
		it.Event = new(WalletEstimatorCallSucceeded)
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
func (it *WalletEstimatorCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorCallSucceeded represents a CallSucceeded event raised by the WalletEstimator contract.
type WalletEstimatorCallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallSucceeded(opts *bind.FilterOpts) (*WalletEstimatorCallSucceededIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallSucceededIterator{contract: _WalletEstimator.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletEstimatorCallSucceeded) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorCallSucceeded)
				if err := _WalletEstimator.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseCallSucceeded(log types.Log) (*WalletEstimatorCallSucceeded, error) {
	event := new(WalletEstimatorCallSucceeded)
	if err := _WalletEstimator.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorDefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletEstimator contract.
type WalletEstimatorDefinedHookIterator struct {
	Event *WalletEstimatorDefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorDefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorDefinedHook)
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
		it.Event = new(WalletEstimatorDefinedHook)
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
func (it *WalletEstimatorDefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorDefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorDefinedHook represents a DefinedHook event raised by the WalletEstimator contract.
type WalletEstimatorDefinedHook struct {
	Selector       [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletEstimator *WalletEstimatorFilterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletEstimatorDefinedHookIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorDefinedHookIterator{contract: _WalletEstimator.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 selector, address implementation)
func (_WalletEstimator *WalletEstimatorFilterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletEstimatorDefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorDefinedHook)
				if err := _WalletEstimator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseDefinedHook(log types.Log) (*WalletEstimatorDefinedHook, error) {
	event := new(WalletEstimatorDefinedHook)
	if err := _WalletEstimator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletEstimator contract.
type WalletEstimatorImageHashUpdatedIterator struct {
	Event *WalletEstimatorImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorImageHashUpdated)
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
		it.Event = new(WalletEstimatorImageHashUpdated)
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
func (it *WalletEstimatorImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorImageHashUpdated represents a ImageHashUpdated event raised by the WalletEstimator contract.
type WalletEstimatorImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletEstimator *WalletEstimatorFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletEstimatorImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorImageHashUpdatedIterator{contract: _WalletEstimator.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletEstimator *WalletEstimatorFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletEstimatorImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorImageHashUpdated)
				if err := _WalletEstimator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseImageHashUpdated(log types.Log) (*WalletEstimatorImageHashUpdated, error) {
	event := new(WalletEstimatorImageHashUpdated)
	if err := _WalletEstimator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletEstimator contract.
type WalletEstimatorImplementationUpdatedIterator struct {
	Event *WalletEstimatorImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorImplementationUpdated)
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
		it.Event = new(WalletEstimatorImplementationUpdated)
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
func (it *WalletEstimatorImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorImplementationUpdated represents a ImplementationUpdated event raised by the WalletEstimator contract.
type WalletEstimatorImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletEstimator *WalletEstimatorFilterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletEstimatorImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorImplementationUpdatedIterator{contract: _WalletEstimator.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletEstimator *WalletEstimatorFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletEstimatorImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorImplementationUpdated)
				if err := _WalletEstimator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseImplementationUpdated(log types.Log) (*WalletEstimatorImplementationUpdated, error) {
	event := new(WalletEstimatorImplementationUpdated)
	if err := _WalletEstimator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletEstimator contract.
type WalletEstimatorNonceChangeIterator struct {
	Event *WalletEstimatorNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorNonceChange)
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
		it.Event = new(WalletEstimatorNonceChange)
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
func (it *WalletEstimatorNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorNonceChange represents a NonceChange event raised by the WalletEstimator contract.
type WalletEstimatorNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletEstimator *WalletEstimatorFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletEstimatorNonceChangeIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorNonceChangeIterator{contract: _WalletEstimator.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletEstimator *WalletEstimatorFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletEstimatorNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorNonceChange)
				if err := _WalletEstimator.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseNonceChange(log types.Log) (*WalletEstimatorNonceChange, error) {
	event := new(WalletEstimatorNonceChange)
	if err := _WalletEstimator.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorStaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletEstimator contract.
type WalletEstimatorStaticSignatureSetIterator struct {
	Event *WalletEstimatorStaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorStaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorStaticSignatureSet)
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
		it.Event = new(WalletEstimatorStaticSignatureSet)
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
func (it *WalletEstimatorStaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorStaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorStaticSignatureSet represents a StaticSignatureSet event raised by the WalletEstimator contract.
type WalletEstimatorStaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletEstimator *WalletEstimatorFilterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletEstimatorStaticSignatureSetIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorStaticSignatureSetIterator{contract: _WalletEstimator.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletEstimator *WalletEstimatorFilterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletEstimatorStaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorStaticSignatureSet)
				if err := _WalletEstimator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseStaticSignatureSet(log types.Log) (*WalletEstimatorStaticSignatureSet, error) {
	event := new(WalletEstimatorStaticSignatureSet)
	if err := _WalletEstimator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEstimatorDeployedBin is the resulting bytecode of the created contract
const WalletEstimatorDeployedBin = "0x6080604052600436101561001e575b361561001c5761001c611d4e565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc14610157578063975befdb146101525780639c145aed1461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e576116d9565b61164c565b61157a565b61141b565b6113d8565b611387565b611318565b611189565b61107b565b61101d565b610fe1565b610f5d565b610f2e565b610e8a565b610d6e565b610cb5565b610b98565b610ac9565b610a14565b61098c565b6108ff565b610802565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b81601f820112156101c55780359061053f82610481565b9261054d6040519485610404565b82845260208085019360051b830101918183116101c55760208101935b83851061057957505050505090565b843567ffffffffffffffff81116101c557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101c5576105c0610445565b916105cd60208301610210565b83526040820135602084015260608201359267ffffffffffffffff84116101c55760e08361060288602080988198010161050a565b60408401526080810135606084015261061d60a08201610474565b608084015261062e60c08201610474565b60a0840152013560c082015281520194019361056a565b9080601f830112156101c557813561065c81610481565b9261066a6040519485610404565b81845260208085019260051b8201019283116101c557602001905b8282106106925750505090565b6020809161069f84610210565b815201910190610685565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c55761071c610456565b9061072981600401610466565b825261073760248201610474565b6020830152604481013567ffffffffffffffff81116101c55784600461075f92840101610528565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c55784600461079b9284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107d8930101610645565b610100820152916024359067ffffffffffffffff82116101c5576107fe91600401610231565b9091565b346101c557610810366106aa565b909161010081019261082b6108268551516117e1565b6117fc565b9160005b85518051821015610892579061088c61086761084d8360019561187a565b5173ffffffffffffffffffffffffffffffffffffffff1690565b610871838861187a565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161082f565b50508383866108a7336108718351518561187a565b526108b3818484611e20565b50156108c55760405160018152602090f35b6108fb906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611ad5565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576109366101a2565b5061093f6101ca565b5060643567ffffffffffffffff81116101c557610960903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109e46109ea923690600401610231565b91611afa565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a9760209160243560443591600401611ba3565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610b0e600435610b0981610a9f565b612027565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c55760043567ffffffffffffffff81116101c55781610b7591600401610231565b929092916024359067ffffffffffffffff82116101c5576107fe91600401610231565b610ba136610b2c565b90929160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b57610bfc9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde555a9361215c565b91610c10606084015160808501519061244b565b610c1b828585611e20565b929015610c535750610c2d93506125a4565b60017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b836108fb86926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611ad5565b7f37ed32e80000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610d44576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610da481610a9f565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610dc982612027565b1615610e2f5760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610e20600082612e8a565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107fe91600401610231565b610f3736610ee5565b9030330361036a57610f4d61001c925a9261215c565b90610f57826128c8565b906125a4565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610f946101a2565b5060443567ffffffffffffffff81116101c557610fb5903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a9760043561295c565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760406110596004356129a2565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b61108436610b2c565b909160027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b576110de9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde555a9461215c565b916110f660608401516110f08161295c565b9061244b565b611101828285611e20565b92901561115257846111148585836125a4565b5a810390811161114d5760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55604051908152602090f35b6117b2565b6108fb84916040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611ad5565b346101c55761119736610ee5565b9060027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5414610c8b5760027ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde5573ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680156112ee5733036112c057303b156101c55761126a9160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611d0d565b038183305af180156112bb576112a05760017ffc6e07e3992c7c3694a921dc9e412b6cfe475380556756a19805a9e3ddfe2fde55005b806112af60006112b593610404565b80610e7f565b38610c2d565b611b46565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c06113f46000806113ee366106aa565b91612aa3565b926040929192519485526020850152600160408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043561145181610a9f565b6114596101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff61147f82612027565b166114f9577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166114ec8183612e8a565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115b16101a2565b506115ba6101ca565b5060443567ffffffffffffffff81116101c5576115db903690600401611549565b505060643567ffffffffffffffff81116101c5576115fd903690600401611549565b505060843567ffffffffffffffff81116101c55761161f903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576116836101a2565b5061168c6101ca565b5060843567ffffffffffffffff81116101c5576116ad903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356117136101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361179e7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612e45565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161114d57565b9190820180921161114d57565b9061180682610481565b6118136040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06118418294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561188e5760209160051b010190565b61184b565b919082519283825260005b8481106118dd5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161189e565b9080602083519182815201916020808360051b8301019401926000915b83831061191e57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611991604085015160e0604086015260e0850190611893565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061190f565b906020808351928381520192019060005b8181106119e65750505090565b825173ffffffffffffffffffffffffffffffffffffffff168452602093840193909201916001016119d9565b805160ff16825261052591602082810151151590820152610100611a70611a4a604085015161012060408601526101208501906118f2565b606085015160608501526080850151608085015260a085015184820360a0860152611893565b9260c081015160c084015260e081015160e08401520151906101008184039101526119c8565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611aec6105259492604085526040850190611a12565b926020818503910152611a96565b90611b179291611b08611fcc565b906003825260e0820152611e20565b5015611b41577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b909173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169081156112ee578133036112c05780611c57575b5050611c477f1626ba7e00000000000000000000000000000000000000000000000000000000926109e4836101007fffffffff00000000000000000000000000000000000000000000000000000000950190611b52565b1603611c5257600090565b600190565b813b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015291600091839160249183915af19283156112bb576109e47fffffffff0000000000000000000000000000000000000000000000000000000093611c47937f1626ba7e0000000000000000000000000000000000000000000000000000000096611cf8575b5093505092611bf0565b806112af6000611d0793610404565b38611cee565b916020610525938181520191611a96565b3d15611d49573d90611d2f82610499565b91611d3d6040519384610404565b82523d6000602084013e565b606090565b600436108015611d5b5750565b611d91906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611de3575b50612027565b73ffffffffffffffffffffffffffffffffffffffff8116611daf5750565b60008091604051368382378036810184815203915af4611dcd611d1e565b9015611ddb57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611d8b565b901561188e5790565b611e53611e2d8484611e17565b357fff000000000000000000000000000000000000000000000000000000000000001690565b7f800000000000000000000000000000000000000000000000000000000000000080821614611ecd5750611e8b926000928392612aa3565b91505091808210611e9d575050600191565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b7f0200000000000000000000000000000000000000000000000000000000000000908116146020820152611f03925090506128c8565b90611f0d826129a2565b42811115611f9a575073ffffffffffffffffffffffffffffffffffffffff81168015159081611f8f575b50611f43575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611f37565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a12084521660408201526040815261209e606082610404565b519020541690565b906120b082610481565b6120bd6040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120eb8294610481565b019060005b8281106120fc57505050565b60209060405161210b816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c0820152828285010152016120f0565b909392938483116101c55784116101c5578101920390565b90612165611fcc565b6000815291600190803560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016124345750600060608601525b60076121b460ff831660011c90565b16806123e3575b506010818116036123b5575060015b6121d3816120a6565b604086019081526000925b82841061221b5750505050036121f15790565b7f0bdf80380000000000000000000000000000000000000000000000000000000060005260046000fd5b9293919290918082013560f81c9060010194908560018083160361239357506122653061224984875161187a565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280821614612373575b600480821614612325575b6008808216146122f0575b906122d76122d160c0846122b1601080600198161460806122a8888c5161187a565b51019015159052565b6122c760208083161460a06122a8888c5161187a565b1660061c60031690565b60ff1690565b60c06122e483875161187a565b510152019291906121de565b94600191906122d7906122d19060c090868101359060200199906060612317878b5161187a565b510152939450505050612286565b9461236d908381013560e81c9060030161236661234d61234584846117ef565b838c89612144565b9190604061235c888b5161187a565b51019236916104d3565b90526117ef565b9461227b565b9482810135906020019590602061238b84875161187a565b510152612270565b6123b096508381013560601c90601401969061224984875161187a565b612265565b6020908116036123d257600282019181013560f01c905b906121ca565b600182019181013560f81c906123cc565b612427919383929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b92906080860152386121bb565b80830135606090811c9087015260140192506121a5565b906124558261295c565b8181036124db57509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8252836040820152604081526124c5606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b61252760409295949395606083526060830190611a12565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611893565b6105259392606092825260208201528160408201520190611893565b6125916105259492606083526060830190611a12565b9260208201526040818403910152611893565b916000604082019384515190825b8281106125c3575b50505050505050565b6125ce81885161187a565b51936125dd60a0860151151590565b806128c0575b612880575060009360608101518015801580612877575b61283f57849061260d6080850151151590565b156127f9576126b892612634855173ffffffffffffffffffffffffffffffffffffffff1690565b91156127f357505a905b6126b38b61268760608d01516040890151908c8b604051998a967f4c4e814c00000000000000000000000000000000000000000000000000000000602089015260248801612531565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612e18565b156126fb575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016125b2565b60c0018051156127aa57600181511461276b575160021461271c57386126be565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b925061275c612750612e2a565b6040519384938461255f565b0390a1388080808080806125ba565b50846108fb612778612e2a565b6040519384937f7f6b0bb10000000000000000000000000000000000000000000000000000000085526004850161257b565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d6127eb6127dd612e2a565b604051918291858c8461255f565b0390a16126f5565b9061263e565b835161283493925073ffffffffffffffffffffffffffffffffffffffff169160208501519160001461283957505a905b604085015192612e06565b6126b8565b90612829565b83886108fb5a6040519384937f213952740000000000000000000000000000000000000000000000000000000085526004850161250f565b50815a106125fa565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b91819081016127eb565b5080156125e3565b61292a6129566128e86128e2602085015115153090612ec9565b93612fc4565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e835260408201526040815261299b606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8683526040820152604081526129e1606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b60405190612a09826103e8565b60006020838281520152565b60031115612a1f57565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b908160409103126101c557602060405191612a68836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611893565b909491939291853560f81c600190938190612abc6129fc565b92612ac682612a15565b60018203612cd0575b50600180871614612c6f575060028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908096918196602016612b3e9060051c90565b612b47906117e1565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a81612b93846128c8565b988993612b9f93612144565b91612ba993613406565b9098612bbd91600052602052604060002090565b90612bd091600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612bf991600052602052604060002090565b94815190868215159283612c64575b505081612c55575b50612c185750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612c10565b141591508638612c08565b909691939450612c8181989398612a15565b612ca657612c9b9581612c9393612144565b9390926131b9565b919394909293929190565b7ffdf132ad0000000000000000000000000000000000000000000000000000000060005260046000fd5b600097507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc06040881601612acf578981013560601c9750601401915087878a84612d1985612a15565b60028503612d2a575b505050612acf565b60038101965093945073ffffffffffffffffffffffffffffffffffffffff9381013560e81c92604092612daf929091612d7a91612d73918a90612d6d89836117ef565b92612144565b36916104d3565b83519586809481937fccce3bc80000000000000000000000000000000000000000000000000000000083523060048401612a76565b0392165afa80156112bb57612dcd92600091612dd7575b50936117ef565b9087388a81612d22565b612df9915060403d604011612dff575b612df18183610404565b810190612a4e565b38612dc6565b503d612de7565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612e84606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612e84606082610404565b15612f79576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a0815261295660c082610404565b4690612ed2565b805160209091019060005b818110612f985750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612f8b565b610100810151604051612fdf8161292a602082018095612f80565b51902090612fee815160ff1690565b60ff8116806130675750509061295661300a6040840151613caa565b9261292a60806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b600181036130c557505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290612956816080810161292a565b6002810361311b57505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252612956816080810161292a565b60030361316f575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252612956816080810161292a565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906125279060409396959496606084526060840191611a96565b91949290926000956000956000956000956000956131d5611fcc565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b82811061322c5750505050505050805115158061321e575b612c185750565b506020810151841115613217565b600381019d50959b5093995091975092909190613250908b9085013560e81c6117ef565b95828703613390578a6001915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c03613361575061329d91613296898c938789612144565b908b612aa3565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b1061331b5750928b885114613312575b808b10156132e057508a60c085015289929592959491949390936131ff565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526132c1565b8d8f6108fb61332c85858c8e612144565b9390926040519485947fb006aba00000000000000000000000000000000000000000000000000000000086526004860161319f565b979899809b926133768b61337d94888a612144565b9086612aa3565b50929d919c909b929a9092918e8e6132b1565b8a60029161325d565b908160209103126101c5575161052581610a9f565b604090610525949281528160208201520191611a96565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611a96565b908160209103126101c5575190565b9391909360009460009460005b818110613421575050505050565b8481013560f881901c9860019092019788979692909160fc1c988915613c15575060018914613bd55760028914613a0857600389146139d9576004891461395857600689146138b8576005891461386a57600789146137a3576008891461374d576009891461362457600a89146134c1577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b90919293949596975060038916978815613613575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019081019081613539918787612144565b6040517f898bd92100000000000000000000000000000000000000000000000000000000815293918491613571918a600485016133ae565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156112bb576135b4936000936135e0575b5060ff909a168091019a613f76565b9080156135da57906135ce91600052602052604060002090565b955b9392919093613413565b506135ce565b60ff9193506136059060203d811161360c575b6135fd8183610404565b8101906133f7565b92906135a5565b503d6135f3565b8084013560f81c98506001016134d6565b9091929394959697506003891697881561373c575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161369c918787612144565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916136d4918b60048501611ad5565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa9182156112bb57613716936000936135e0575060ff909a168091019a613f76565b908015613736579061373091600052602052604060002090565b956135d0565b50613730565b8084013560f81c9850600101613639565b98506020870197509495939492939192909182013561376b86613f1d565b811461377b575b61371690613f37565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613772565b975090919293949597600f16968715613858575b602060006137c96138369a9b86613de5565b9c9092918a60405161380c8161292a8a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa156112bb576137169060ff6000519a1680910199613e28565b600189019883013560f81c97506137b7565b985060208701975094959394929391929091820135808514613890575b61371690613ede565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613887565b989091929394959662ffffff98506138da6122d1600c8416603f9060021c1690565b918215613944575b6003168015613933575b9081906139179061390f908781013560e81c906003019c168c01809c8989612144565b90898b613406565b91111561392a575b906137169291613e93565b9982019961391f565b50600281019084013560f01c6138ec565b8482013560f81c92506001909101906138e2565b9750976139ae6139bb929394959697600f6139c393169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686612144565b908688613406565b9061373092980198600052602052604060002090565b985096509394929391929091908082013590602001968015613736579061373091600052602052604060002090565b90919293949596975060038916978815613bc4575b8084013560601c99613a7c9160140190613a3c9060021c6003166122d1565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b90810190613ae060208c613a9285858b8b612144565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e000000000000000000000000000000000000000000000000000000008552600485016133ae565b0392165afa9081156112bb577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613b96575b501603613b5257509060ff61371692991680910199613e28565b6108fb613b638c9389938989612144565b906040519485947fb2fed7ae000000000000000000000000000000000000000000000000000000008652600486016133c5565b613bb7915060203d8111613bbd575b613baf8183610404565b810190613399565b38613b38565b503d613ba5565b8381013560f81c9850600101613a1d565b98600f91929394959697985016968715613c04575b60148101976137169160ff9091169084013560601c613e28565b8281013560f81c9750600101613bea565b98509091929394959698600f16978815613c67575b5060206000613c3d6138369a9b86613de5565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c97506020613c2a565b805160209091019060005b818110613c945750505090565b8251845260209384019390920191600101613c87565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613cf0613cda83610481565b92613ce86040519485610404565b808452610481565b0136602083013760005b8351811015613dcc5780613d106001928661187a565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613db861012082610404565b519020613dc5828561187a565b5201613cfa565b509091506040516129568161292a602082018095613c7c565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff811161114d5791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b166031830152604582015260458152612956606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a0000000000000000855260388401526058830152607882015260788152612956609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a00000000000000008352603882015260388152612956605882610404565b61292a6129566128e86128e2600060208601511515612ec9565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a8352604082015260408152612956606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d8152612956608d8261040456fea2646970667358221220f32f6ee7ce70c5af164bab9106b12e6b2b298f8aa56010b00e260d8ea76b907464736f6c634300081c0033"
