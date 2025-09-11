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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_entryPoint\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"entrypoint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"estimate\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeUserOp\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"imageHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"readHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverPartialSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isValidImage\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"selfExecute\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setStaticSignature\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tokensReceived\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"operatorData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImageHash\",\"inputs\":[{\"name\":\"_imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateImplementation\",\"inputs\":[{\"name\":\"_implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"validateUserOp\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structPackedUserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"accountGasLimits\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"preVerificationGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasFees\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"paymasterAndData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"missingAccountFunds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"validationData\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefinedHook\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"},{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImageHashUpdated\",\"inputs\":[{\"name\":\"newImageHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImplementationUpdated\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NonceChange\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaticSignatureSet\",\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_timestamp\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadNonce\",\"inputs\":[{\"name\":\"_space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_current\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4337Disabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"HookAlreadyExists\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"HookDoesNotExist\",\"inputs\":[{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"ImageHashIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidERC1271Signature\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidEntryPoint\",\"inputs\":[{\"name\":\"_entrypoint\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidSapientSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureFlag\",\"inputs\":[{\"name\":\"_flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidSignatureWeight\",\"inputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureExpired\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expires\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStaticSignatureWrongCaller\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_expectedCaller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LowWeightChainedSignature\",\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlySelf\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"UnusedSnapshot\",\"inputs\":[{\"name\":\"_snapshot\",\"type\":\"tuple\",\"internalType\":\"structSnapshot\",\"components\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"type\":\"error\",\"name\":\"WrongChainedCheckpointOrder\",\"inputs\":[{\"name\":\"_nextCheckpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_checkpoint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a034607457601f613e6d38819003918201601f19168301916001600160401b03831184841017607957808492602094604052833981010312607457516001600160a01b0381168103607457608052604051613ddd908161009082396080518181816110a3015281816111e00152611a660152f35b600080fd5b634e487b7160e01b600052604160045260246000fdfe6080604052600436101561001e575b361561001c5761001c611c2e565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc14610157578063975befdb146101525780639c145aed1461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e57611556565b6114c9565b6113f7565b611298565b611255565b611204565b611195565b61107d565b610fde565b610f80565b610f44565b610ec0565b610e91565b610ded565b610cd1565b610c18565b610b98565b610ac9565b610a14565b61098c565b6108ff565b610802565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b81601f820112156101c55780359061053f82610481565b9261054d6040519485610404565b82845260208085019360051b830101918183116101c55760208101935b83851061057957505050505090565b843567ffffffffffffffff81116101c557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101c5576105c0610445565b916105cd60208301610210565b83526040820135602084015260608201359267ffffffffffffffff84116101c55760e08361060288602080988198010161050a565b60408401526080810135606084015261061d60a08201610474565b608084015261062e60c08201610474565b60a0840152013560c082015281520194019361056a565b9080601f830112156101c557813561065c81610481565b9261066a6040519485610404565b81845260208085019260051b8201019283116101c557602001905b8282106106925750505090565b6020809161069f84610210565b815201910190610685565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c55761071c610456565b9061072981600401610466565b825261073760248201610474565b6020830152604481013567ffffffffffffffff81116101c55784600461075f92840101610528565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c55784600461079b9284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107d8930101610645565b610100820152916024359067ffffffffffffffff82116101c5576107fe91600401610231565b9091565b346101c557610810366106aa565b909161010081019261082b61082685515161165e565b611679565b9160005b85518051821015610892579061088c61086761084d836001956116f7565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61087183886116f7565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161082f565b50508383866108a733610871835151856116f7565b526108b3818484611d00565b50156108c55760405160018152602090f35b6108fb906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611952565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576109366101a2565b5061093f6101ca565b5060643567ffffffffffffffff81116101c557610960903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109e46109ea923690600401610231565b91611977565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a9760209160243560443591600401611a4c565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610b0e600435610b0981610a9f565b611edc565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c55760043567ffffffffffffffff81116101c55781610b7591600401610231565b929092916024359067ffffffffffffffff82116101c5576107fe91600401610231565b610baf610ba436610b2c565b919390935a93612011565b91610bc360608401516080850151906122d5565b610bce828585611d00565b929015610be0575061001c935061242e565b836108fb86926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611952565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610ca7576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610d0781610a9f565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d2c82611edc565b1615610d925760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d83600082612c86565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107fe91600401610231565b610e9a36610e48565b9030330361036a57610eb061001c925a92612011565b90610eba82612752565b9061242e565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610ef76101a2565b5060443567ffffffffffffffff81116101c557610f18903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a976004356127e6565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610fbc60043561282c565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b610ff6610fea36610b2c565b91939092935a94612011565b9161100e6060840151611008816127e6565b906122d5565b611019828285611d00565b929015611046578461102c85858361242e565b5a810390811161104157602090604051908152f35b61162f565b6108fb84916040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611952565b346101c55761108b36610e48565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561116b57330361113d57303b156101c5576111119160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611bed565b038183305af180156111385761112357005b80611132600061001c93610404565b80610de2565b6119c3565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061127160008061126b366106aa565b916128f4565b926040929192519485526020850152600160408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356112ce81610a9f565b6112d66101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff6112fc82611edc565b16611376577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166113698183612c86565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761142e6101a2565b506114376101ca565b5060443567ffffffffffffffff81116101c5576114589036906004016113c6565b505060643567ffffffffffffffff81116101c55761147a9036906004016113c6565b505060843567ffffffffffffffff81116101c55761149c903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115006101a2565b506115096101ca565b5060843567ffffffffffffffff81116101c55761152a903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356115906101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361161b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612c41565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161104157565b9190820180921161104157565b9061168382610481565b6116906040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116be8294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561170b5760209160051b010190565b6116c8565b919082519283825260005b84811061175a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161171b565b9080602083519182815201916020808360051b8301019401926000915b83831061179b57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c08061180e604085015160e0604086015260e0850190611710565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061178c565b906020808351928381520192019060005b8181106118635750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611856565b805160ff168252610525916020828101511515908201526101006118ed6118c76040850151610120604086015261012085019061176f565b606085015160608501526080850151608085015260a085015184820360a0860152611710565b9260c081015160c084015260e081015160e0840152015190610100818403910152611845565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611969610525949260408552604085019061188f565b926020818503910152611913565b906119949291611985611e81565b906003825260e0820152611d00565b50156119be577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b908160209103126101c5575161052581610a9f565b604090610525949281528160208201520191611913565b91909173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691821561116b5782330361113d5780611b7e575b506020915080610100611aae9201906119cf565b92611ae660405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a35565b0381305afa908115611138577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611b4f575b501603611b4a57600090565b600190565b611b71915060203d602011611b77575b611b698183610404565b810190611a20565b38611b3e565b503d611b5f565b823b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af190811561113857602092611aae92611bd8575b5090611a9a565b806111326000611be793610404565b38611bd1565b916020610525938181520191611913565b3d15611c29573d90611c0f82610499565b91611c1d6040519384610404565b82523d6000602084013e565b606090565b600436108015611c3b5750565b611c71906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611cc3575b50611edc565b73ffffffffffffffffffffffffffffffffffffffff8116611c8f5750565b60008091604051368382378036810184815203915af4611cad611bfe565b9015611cbb57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611c6b565b901561170b5790565b91907f800000000000000000000000000000000000000000000000000000000000000080611d57611d318585611cf7565b357fff000000000000000000000000000000000000000000000000000000000000001690565b1614611dad576000918291611d6b946128f4565b91505091808210611d7d575050600191565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050611db890612752565b90611dc28261282c565b42811115611e4f575073ffffffffffffffffffffffffffffffffffffffff81168015159081611e44575b50611df8575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611dec565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152611f53606082610404565b519020541690565b90611f6582610481565b611f726040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611fa08294610481565b019060005b828110611fb157505050565b602090604051611fc0816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201611fa5565b909392938483116101c55784116101c5578101920390565b91909161201c611e81565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016122be5750600060608701525b600761206b60ff831660011c90565b168061226d575b5060108181160361223857506001925b61208b84611f5b565b92604087019384526000905b8582106120a657505050505050565b8281013560f81c9060010190918160018085160361221657506120ea306120ce8389516116f7565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808416146121f6575b6004808416146121a8575b600880841614612175575b61215e61215860c0856121378a608061212e8860108060019c9d161493516116f7565b51019015159052565b61214e8a60a061212e8860208087161493516116f7565b1660061c60031690565b60ff1690565b60c061216b8389516116f7565b5101520190612097565b60019161215e906121589060c09087810135906020019690606061219a878d516116f7565b51015295945050505061210b565b906121f0908481013560e81c906003016121e96121d06121c8848461166c565b838a8a611ff9565b919060406121df888d516116f7565b51019236916104d3565b905261166c565b90612100565b9083810135906020019190602061220e8389516116f7565b5101526120f5565b61223392508481013560601c9060140192906120ce8389516116f7565b6120ea565b60209081160361225a5761ffff918381013560f01c906002015b921692612082565b60ff918381013560f81c90600101612252565b6122b1919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608087015238612072565b80850135606090811c90880152601401925061205c565b906122df826127e6565b81810361236557509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e82528360408201526040815261234f606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b6123b16040929594939560608352606083019061188f565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611710565b6105259392606092825260208201528160408201520190611710565b61241b610525949260608352606083019061188f565b9260208201526040818403910152611710565b916000604082019384515190825b82811061244d575b50505050505050565b6124588188516116f7565b519361246760a0860151151590565b8061274a575b61270a575060009360608101518015801580612701575b6126c95784906124976080850151151590565b1561268357612542926124be855173ffffffffffffffffffffffffffffffffffffffff1690565b911561267d57505a905b61253d8b61251160608d01516040890151908c8b604051998a967f4c4e814c000000000000000000000000000000000000000000000000000000006020890152602488016123bb565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612c14565b15612585575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b0161243c565b60c0018051156126345760018151146125f557516002146125a65738612548565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b92506125e66125da612c26565b604051938493846123e9565b0390a138808080808080612444565b50846108fb612602612c26565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612405565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612675612667612c26565b604051918291858c846123e9565b0390a161257f565b906124c8565b83516126be93925073ffffffffffffffffffffffffffffffffffffffff16916020850151916000146126c357505a905b604085015192612c02565b612542565b906126b3565b83886108fb5a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612399565b50815a10612484565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612675565b50801561246d565b6127b46127e061277261276c602085015115153090612cc5565b93612dc0565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612825606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261286b606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b60405190612893826103e8565b60006020838281520152565b908160409103126101c5576020604051916128b9836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611710565b909491939291853560f81c6001909381908761290e612886565b9360408089161480612be4575b612ae7575b505050600180861614612ac15760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019080969181966020166129909060051c90565b6129999061165e565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a816129e584612752565b9889936129f193611ff9565b916129fb936131d7565b9098612a0f91600052602052604060002090565b90612a2291600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612a4b91600052602052604060002090565b94815190868215159283612ab6575b505081612aa7575b50612a6a5750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612a62565b141591508638612a5a565b86612adc9792949550612ad49398611ff9565b939092612fb5565b919394909293929190565b8a81013560601c985060140192509087908a90849015612b0b575b50829150612920565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c92604092612b8e929091612b5991612b52918990612b4c898361166c565b92611ff9565b36916104d3565b83519586809481937fccce3bc800000000000000000000000000000000000000000000000000000000835230600484016128c7565b0392165afa801561113857612bac92600091612bb5575b509261166c565b86388981612b02565b612bd7915060403d604011612bdd575b612bcf8183610404565b81019061289f565b38612ba5565b503d612bc5565b5073ffffffffffffffffffffffffffffffffffffffff89161561291b565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612c80606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612c80606082610404565b15612d75576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a081526127e060c082610404565b4690612cce565b805160209091019060005b818110612d945750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612d87565b610100810151604051612ddb816127b4602082018095612d7c565b51902090612dea815160ff1690565b60ff811680612e63575050906127e0612e066040840151613a6b565b926127b460806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612ec157505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466938101938452908101919091526060810192909252906127e081608081016127b4565b60028103612f1757505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082019081529181019290925260608201929092526127e081608081016127b4565b600303612f6b575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082019081529181019290925260608201929092526127e081608081016127b4565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906123b19060409396959496606084526060840191611913565b9194929092600095600095600095600095600095612fd1611e81565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b8281106130285750505050505050805115158061301a575b612a6a5750565b506020810151841115613013565b600381019d50959b509399509197509290919061304c908b9085013560e81c61166c565b958a6000848903613190575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361316057509060016130958961309c948789611ff9565b908b6128f4565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b1061311a5750928b885114613111575b808b10156130df57508a60c08501528992959295949194939093612ffb565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526130c0565b8d8f6108fb61312b85858c8e611ff9565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601612f9b565b9798999a9160016131768b61317d94888a611ff9565b90866128f4565b50929d919c909b929a9092918e8e6130b0565b9161305b565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611913565b908160209103126101c5575190565b9391909360009460009460005b8181106131f2575050505050565b8481013560f881901c9860019092019788979692909160fc1c9889156139d657506001891461399657600289146137d957600389146137aa57600489146137295760068914613689576005891461363b5760078914613574576008891461351e57600989146133f557600a8914613292577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b909192939495969750600389169788156133e4575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161330a918787611ff9565b6040517f898bd92100000000000000000000000000000000000000000000000000000000815293918491613342918a60048501611a35565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561113857613385936000936133b1575b5060ff909a168091019a613d37565b9080156133ab579061339f91600052602052604060002090565b955b93929190936131e4565b5061339f565b60ff9193506133d69060203d81116133dd575b6133ce8183610404565b8101906131c8565b9290613376565b503d6133c4565b8084013560f81c98506001016132a7565b9091929394959697506003891697881561350d575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161346d918787611ff9565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916134a5918b60048501611952565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611138576134e7936000936133b1575060ff909a168091019a613d37565b908015613507579061350191600052602052604060002090565b956133a1565b50613501565b8084013560f81c985060010161340a565b98506020870197509495939492939192909182013561353c86613cde565b811461354c575b6134e790613cf8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613543565b975090919293949597600f16968715613629575b6020600061359a6136079a9b86613ba6565b9c9092918a6040516135dd816127b48a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15611138576134e79060ff6000519a1680910199613be9565b600189019883013560f81c9750613588565b985060208701975094959394929391929091820135808514613661575b6134e790613c9f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613658565b989091929394959662ffffff98506136ab612158600c8416603f9060021c1690565b918215613715575b6003168015613704575b9081906136e8906136e0908781013560e81c906003019c168c01809c8989611ff9565b90898b6131d7565b9111156136fb575b906134e79291613c54565b998201996136f0565b50600281019084013560f01c6136bd565b8482013560f81c92506001909101906136b3565b97509761377f61378c929394959697600f61379493169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686611ff9565b9086886131d7565b9061350192980198600052602052604060002090565b985096509394929391929091908082013590602001968015613507579061350191600052602052604060002090565b90919293949596975060038916978815613985575b8084013560601c9961384d916014019061380d9060021c600316612158565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b908101906138b160208c61386385858b8b611ff9565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a35565b0392165afa908115611138577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613967575b50160361392357509060ff6134e792991680910199613be9565b6108fb6139348c9389938989611ff9565b906040519485947fb2fed7ae00000000000000000000000000000000000000000000000000000000865260048601613196565b61397f915060203d8111611b7757611b698183610404565b38613909565b8381013560f81c98506001016137ee565b98600f919293949596979850169687156139c5575b60148101976134e79160ff9091169084013560601c613be9565b8281013560f81c97506001016139ab565b98509091929394959698600f16978815613a28575b50602060006139fe6136079a9b86613ba6565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c975060206139eb565b805160209091019060005b818110613a555750505090565b8251845260209384019390920191600101613a48565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613ab1613a9b83610481565b92613aa96040519485610404565b808452610481565b0136602083013760005b8351811015613b8d5780613ad1600192866116f7565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613b7961012082610404565b519020613b8682856116f7565b5201613abb565b509091506040516127e0816127b4602082018095613a3d565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116110415791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526127e0606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526127e0609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526127e0605882610404565b6127b46127e061277261276c600060208601511515612cc5565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526127e0606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d81526127e0608d8261040456fea2646970667358221220fd510fff57502b123cb8cca23f77a513f5d8001f96adc68f13bd2a44f529e4ad64736f6c634300081c0033",
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
const WalletEstimatorDeployedBin = "0x6080604052600436101561001e575b361561001c5761001c611c2e565b005b60003560e01c806223de291461019d578063025b22bc1461019857806313792a4a14610193578063150b7a021461018e5780631626ba7e1461018957806319822f7c146101845780631a9b23371461017f5780631f6a1eb91461017a57806329561426146101755780634fcf3eca1461017057806351605d801461016b5780636ea44577146101665780638943ec02146101615780638c3f55631461015c57806392dcb3fc14610157578063975befdb146101525780639c145aed1461014d578063a65d69d414610148578063aaf10f4214610143578063ad55366b1461013e578063b93ea7ad14610139578063bc197c8114610134578063f23a6e611461012f5763f727ef1c0361000e57611556565b6114c9565b6113f7565b611298565b611255565b611204565b611195565b61107d565b610fde565b610f80565b610f44565b610ec0565b610e91565b610ded565b610cd1565b610c18565b610b98565b610ac9565b610a14565b61098c565b6108ff565b610802565b6102eb565b61025f565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b359073ffffffffffffffffffffffffffffffffffffffff821682036101c557565b9181601f840112156101c55782359167ffffffffffffffff83116101c557602083818601950101116101c557565b346101c55760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576102966101a2565b5061029f6101ca565b506102a86101ed565b5060843567ffffffffffffffff81116101c5576102c9903690600401610231565b505060a43567ffffffffffffffff81116101c55761001c903690600401610231565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761031d6101a2565b30330361036a576020817f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0392305573ffffffffffffffffffffffffffffffffffffffff60405191168152a1005b7fa19dbf00000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60e0810190811067ffffffffffffffff8211176103e357604052565b610398565b6040810190811067ffffffffffffffff8211176103e357604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176103e357604052565b6040519061045460e083610404565b565b6040519061045461012083610404565b359060ff821682036101c557565b359081151582036101c557565b67ffffffffffffffff81116103e35760051b60200190565b67ffffffffffffffff81116103e357601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9291926104df82610499565b916104ed6040519384610404565b8294818452818301116101c5578281602093846000960137010152565b9080601f830112156101c557816020610525933591016104d3565b90565b81601f820112156101c55780359061053f82610481565b9261054d6040519485610404565b82845260208085019360051b830101918183116101c55760208101935b83851061057957505050505090565b843567ffffffffffffffff81116101c557820160e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe082860301126101c5576105c0610445565b916105cd60208301610210565b83526040820135602084015260608201359267ffffffffffffffff84116101c55760e08361060288602080988198010161050a565b60408401526080810135606084015261061d60a08201610474565b608084015261062e60c08201610474565b60a0840152013560c082015281520194019361056a565b9080601f830112156101c557813561065c81610481565b9261066a6040519485610404565b81845260208085019260051b8201019283116101c557602001905b8282106106925750505090565b6020809161069f84610210565b815201910190610685565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82850301126101c55761071c610456565b9061072981600401610466565b825261073760248201610474565b6020830152604481013567ffffffffffffffff81116101c55784600461075f92840101610528565b6040830152606481013560608301526084810135608083015260a481013567ffffffffffffffff81116101c55784600461079b9284010161050a565b60a083015260c481013560c083015260e481013560e083015261010481013567ffffffffffffffff81116101c557600485916107d8930101610645565b610100820152916024359067ffffffffffffffff82116101c5576107fe91600401610231565b9091565b346101c557610810366106aa565b909161010081019261082b61082685515161165e565b611679565b9160005b85518051821015610892579061088c61086761084d836001956116f7565b5173ffffffffffffffffffffffffffffffffffffffff1690565b61087183886116f7565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0161082f565b50508383866108a733610871835151856116f7565b526108b3818484611d00565b50156108c55760405160018152602090f35b6108fb906040519384937ff58cc8b500000000000000000000000000000000000000000000000000000000855260048501611952565b0390fd5b346101c55760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576109366101a2565b5061093f6101ca565b5060643567ffffffffffffffff81116101c557610960903690600401610231565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b346101c55760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043560243567ffffffffffffffff81116101c5576020916109e46109ea923690600401610231565b91611977565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043567ffffffffffffffff81116101c5576101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc82360301126101c557610a9760209160243560443591600401611a4c565b604051908152f35b7fffffffff000000000000000000000000000000000000000000000000000000008116036101c557565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610b0e600435610b0981610a9f565b611edc565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c55760043567ffffffffffffffff81116101c55781610b7591600401610231565b929092916024359067ffffffffffffffff82116101c5576107fe91600401610231565b610baf610ba436610b2c565b919390935a93612011565b91610bc360608401516080850151906122d5565b610bce828585611d00565b929015610be0575061001c935061242e565b836108fb86926040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611952565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760043530330361036a578015610ca7576020817f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa927fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf855604051908152a1005b7f4294d1270000000000000000000000000000000000000000000000000000000060005260046000fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557600435610d0781610a9f565b30330361036a5773ffffffffffffffffffffffffffffffffffffffff610d2c82611edc565b1615610d925760407fffffffff000000000000000000000000000000000000000000000000000000007f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19216610d83600082612c86565b815190815260006020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f1c3812cc000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60009103126101c557565b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55760207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf854604051908152f35b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101c5576004359067ffffffffffffffff82116101c5576107fe91600401610231565b610e9a36610e48565b9030330361036a57610eb061001c925a92612011565b90610eba82612752565b9061242e565b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557610ef76101a2565b5060443567ffffffffffffffff81116101c557610f18903690600401610231565b505060206040517f8943ec02000000000000000000000000000000000000000000000000000000008152f35b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020610a976004356127e6565b346101c55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576040610fbc60043561282c565b73ffffffffffffffffffffffffffffffffffffffff8351921682526020820152f35b610ff6610fea36610b2c565b91939092935a94612011565b9161100e6060840151611008816127e6565b906122d5565b611019828285611d00565b929015611046578461102c85858361242e565b5a810390811161104157602090604051908152f35b61162f565b6108fb84916040519384937fa2b6d61b00000000000000000000000000000000000000000000000000000000855260048501611952565b346101c55761108b36610e48565b9073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016801561116b57330361113d57303b156101c5576111119160009160405193849283927f6ea4457700000000000000000000000000000000000000000000000000000000845260048401611bed565b038183305af180156111385761112357005b80611132600061001c93610404565b80610de2565b6119c3565b7f1d6ddbf4000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b7fd13d78350000000000000000000000000000000000000000000000000000000060005260046000fd5b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c557602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b346101c55760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576020305473ffffffffffffffffffffffffffffffffffffffff60405191168152f35b346101c55760c061127160008061126b366106aa565b916128f4565b926040929192519485526020850152600160408501526060840152608083015260a0820152f35b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356112ce81610a9f565b6112d66101ca565b9030330361036a5773ffffffffffffffffffffffffffffffffffffffff6112fc82611edc565b16611376577f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19173ffffffffffffffffffffffffffffffffffffffff7fffffffff000000000000000000000000000000000000000000000000000000006040931691166113698183612c86565b82519182526020820152a1005b7fffffffff00000000000000000000000000000000000000000000000000000000907f5b4d6d6a000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b9181601f840112156101c55782359167ffffffffffffffff83116101c5576020808501948460051b0101116101c557565b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c55761142e6101a2565b506114376101ca565b5060443567ffffffffffffffff81116101c5576114589036906004016113c6565b505060643567ffffffffffffffff81116101c55761147a9036906004016113c6565b505060843567ffffffffffffffff81116101c55761149c903690600401610231565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b346101c55760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576115006101a2565b506115096101ca565b5060843567ffffffffffffffff81116101c55761152a903690600401610231565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b346101c55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c5576004356115906101ca565b604435916bffffffffffffffffffffffff83168093036101c55730330361036a578273ffffffffffffffffffffffffffffffffffffffff8361161b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1967fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606097881b161785612c41565b6040519384521660208301526040820152a1005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906001820180921161104157565b9190820180921161104157565b9061168382610481565b6116906040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06116be8294610481565b0190602036910137565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b805182101561170b5760209160051b010190565b6116c8565b919082519283825260005b84811061175a5750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b8060208092840101518282860101520161171b565b9080602083519182815201916020808360051b8301019401926000915b83831061179b57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c08061180e604085015160e0604086015260e0850190611710565b936060810151606085015260808101511515608085015260a0810151151560a085015201519101529701930193019193929061178c565b906020808351928381520192019060005b8181106118635750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611856565b805160ff168252610525916020828101511515908201526101006118ed6118c76040850151610120604086015261012085019061176f565b606085015160608501526080850151608085015260a085015184820360a0860152611710565b9260c081015160c084015260e081015160e0840152015190610100818403910152611845565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b91611969610525949260408552604085019061188f565b926020818503910152611913565b906119949291611985611e81565b906003825260e0820152611d00565b50156119be577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b600090565b6040513d6000823e3d90fd5b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1813603018212156101c5570180359067ffffffffffffffff82116101c5576020019181360383136101c557565b908160209103126101c5575161052581610a9f565b604090610525949281528160208201520191611913565b91909173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691821561116b5782330361113d5780611b7e575b506020915080610100611aae9201906119cf565b92611ae660405194859384937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a35565b0381305afa908115611138577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091611b4f575b501603611b4a57600090565b600190565b611b71915060203d602011611b77575b611b698183610404565b810190611a20565b38611b3e565b503d611b5f565b823b156101c5576040517fb760faf900000000000000000000000000000000000000000000000000000000815230600482015292600091849160249183915af190811561113857602092611aae92611bd8575b5090611a9a565b806111326000611be793610404565b38611bd1565b916020610525938181520191611913565b3d15611c29573d90611c0f82610499565b91611c1d6040519384610404565b82523d6000602084013e565b606090565b600436108015611c3b5750565b611c71906000357fffffffff00000000000000000000000000000000000000000000000000000000811691611cc3575b50611edc565b73ffffffffffffffffffffffffffffffffffffffff8116611c8f5750565b60008091604051368382378036810184815203915af4611cad611bfe565b9015611cbb57602081519101f35b602081519101fd5b7fffffffff000000000000000000000000000000000000000000000000000000008092503660040360031b1b161638611c6b565b901561170b5790565b91907f800000000000000000000000000000000000000000000000000000000000000080611d57611d318585611cf7565b357fff000000000000000000000000000000000000000000000000000000000000001690565b1614611dad576000918291611d6b946128f4565b91505091808210611d7d575050600191565b7ffd41fcba0000000000000000000000000000000000000000000000000000000060005260045260245260446000fd5b5050611db890612752565b90611dc28261282c565b42811115611e4f575073ffffffffffffffffffffffffffffffffffffffff81168015159081611e44575b50611df8575060019190565b7f8945c3130000000000000000000000000000000000000000000000000000000060005260048390523360245273ffffffffffffffffffffffffffffffffffffffff1660445260646000fd5b905033141538611dec565b7ff95b6ab700000000000000000000000000000000000000000000000000000000600052600484905260245260446000fd5b60405190610120820182811067ffffffffffffffff8211176103e3576040526060610100836000815260006020820152826040820152600083820152600060808201528260a0820152600060c0820152600060e08201520152565b73ffffffffffffffffffffffffffffffffffffffff906040517fffffffff0000000000000000000000000000000000000000000000000000000060208201927fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a120845216604082015260408152611f53606082610404565b519020541690565b90611f6582610481565b611f726040519182610404565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0611fa08294610481565b019060005b828110611fb157505050565b602090604051611fc0816103c7565b60008152600083820152606060408201526000606082015260006080820152600060a0820152600060c082015282828501015201611fa5565b909392938483116101c55784116101c5578101920390565b91909161201c611e81565b6000815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff818316016122be5750600060608701525b600761206b60ff831660011c90565b168061226d575b5060108181160361223857506001925b61208b84611f5b565b92604087019384526000905b8582106120a657505050505050565b8281013560f81c9060010190918160018085160361221657506120ea306120ce8389516116f7565b519073ffffffffffffffffffffffffffffffffffffffff169052565b6002808416146121f6575b6004808416146121a8575b600880841614612175575b61215e61215860c0856121378a608061212e8860108060019c9d161493516116f7565b51019015159052565b61214e8a60a061212e8860208087161493516116f7565b1660061c60031690565b60ff1690565b60c061216b8389516116f7565b5101520190612097565b60019161215e906121589060c09087810135906020019690606061219a878d516116f7565b51015295945050505061210b565b906121f0908481013560e81c906003016121e96121d06121c8848461166c565b838a8a611ff9565b919060406121df888d516116f7565b51019236916104d3565b905261166c565b90612100565b9083810135906020019190602061220e8389516116f7565b5101526120f5565b61223392508481013560601c9060140192906120ce8389516116f7565b6120ea565b60209081160361225a5761ffff918381013560f01c906002015b921692612082565b60ff918381013560f81c90600101612252565b6122b1919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290608087015238612072565b80850135606090811c90880152601401925061205c565b906122df826127e6565b81810361236557509060017f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8819201908160405160208101907f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e82528360408201526040815261234f606082610404565b51902055604080519182526020820192909252a1565b917f9b6514f40000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b6123b16040929594939560608352606083019061188f565b9460208201520152565b9261052596959260c09592855260208501526040840152606083015260808201528160a08201520190611710565b6105259392606092825260208201528160408201520190611710565b61241b610525949260608352606083019061188f565b9260208201526040818403910152611710565b916000604082019384515190825b82811061244d575b50505050505050565b6124588188516116f7565b519361246760a0860151151590565b8061274a575b61270a575060009360608101518015801580612701575b6126c95784906124976080850151151590565b1561268357612542926124be855173ffffffffffffffffffffffffffffffffffffffff1690565b911561267d57505a905b61253d8b61251160608d01516040890151908c8b604051998a967f4c4e814c000000000000000000000000000000000000000000000000000000006020890152602488016123bb565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101855284610404565b612c14565b15612585575b506040805187815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b0161243c565b60c0018051156126345760018151146125f557516002146125a65738612548565b9493505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b92506125e66125da612c26565b604051938493846123e9565b0390a138808080808080612444565b50846108fb612602612c26565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612405565b509250600180937f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d612675612667612c26565b604051918291858c846123e9565b0390a161257f565b906124c8565b83516126be93925073ffffffffffffffffffffffffffffffffffffffff16916020850151916000146126c357505a905b604085015192612c02565b612542565b906126b3565b83886108fb5a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612399565b50815a10612484565b6040805188815260208101849052919550600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101612675565b50801561246d565b6127b46127e061277261276c602085015115153090612cc5565b93612dc0565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610404565b51902090565b60405160208101917f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e8352604082015260408152612825606082610404565b5190205490565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86835260408201526040815261286b606082610404565b51902054906bffffffffffffffffffffffff8260601c921690565b60405190612893826103e8565b60006020838281520152565b908160409103126101c5576020604051916128b9836103e8565b805183520151602082015290565b60409073ffffffffffffffffffffffffffffffffffffffff61052594931681528160208201520190611710565b909491939291853560f81c6001909381908761290e612886565b9360408089161480612be4575b612ae7575b505050600180861614612ac15760028581161460208501526007600286901c1688820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011691019080969181966020166129909060051c90565b6129999061165e565b8a820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101959098899a816129e584612752565b9889936129f193611ff9565b916129fb936131d7565b9098612a0f91600052602052604060002090565b90612a2291600052602052604060002090565b9073ffffffffffffffffffffffffffffffffffffffff16612a4b91600052602052604060002090565b94815190868215159283612ab6575b505081612aa7575b50612a6a5750565b6040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020909101516024820152604490fd5b90506020820151101538612a62565b141591508638612a5a565b86612adc9792949550612ad49398611ff9565b939092612fb5565b919394909293929190565b8a81013560601c985060140192509087908a90849015612b0b575b50829150612920565b60038101955073ffffffffffffffffffffffffffffffffffffffff945081013560e81c92604092612b8e929091612b5991612b52918990612b4c898361166c565b92611ff9565b36916104d3565b83519586809481937fccce3bc800000000000000000000000000000000000000000000000000000000835230600484016128c7565b0392165afa801561113857612bac92600091612bb5575b509261166c565b86388981612b02565b612bd7915060403d604011612bdd575b612bcf8183610404565b81019061289f565b38612ba5565b503d612bc5565b5073ffffffffffffffffffffffffffffffffffffffff89161561291b565b9160009391849360208451940192f190565b9160009291839260208351930191f490565b3d90604051916020818401016040528083526000602084013e565b60405160208101917fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e868352604082015260408152612c80606082610404565b51902055565b60405160208101917fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1208352604082015260408152612c80606082610404565b15612d75576000905b73ffffffffffffffffffffffffffffffffffffffff6040519160208301937f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f85527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408501527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606085015260808401521660a082015260a081526127e060c082610404565b4690612cce565b805160209091019060005b818110612d945750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612d87565b610100810151604051612ddb816127b4602082018095612d7c565b51902090612dea815160ff1690565b60ff811680612e63575050906127e0612e066040840151613a6b565b926127b460806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612ec157505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466938101938452908101919091526060810192909252906127e081608081016127b4565b60028103612f1757505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082019081529181019290925260608201929092526127e081608081016127b4565b600303612f6b575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082019081529181019290925260608201929092526127e081608081016127b4565b7f048183200000000000000000000000000000000000000000000000000000000060005260ff1660045260246000fd5b906123b19060409396959496606084526060840191611913565b9194929092600095600095600095600095600095612fd1611e81565b60028152937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9460005b8281106130285750505050505050805115158061301a575b612a6a5750565b506020810151841115613013565b600381019d50959b509399509197509290919061304c908b9085013560e81c61166c565b958a6000848903613190575089915b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8c0361316057509060016130958961309c948789611ff9565b908b6128f4565b9c939c9b929b9a919a99909a9b9d8e9d9e8f905b1061311a5750928b885114613111575b808b10156130df57508a60c08501528992959295949194939093612ffb565b7f37daf62b0000000000000000000000000000000000000000000000000000000060005260048b905260245260446000fd5b600088526130c0565b8d8f6108fb61312b85858c8e611ff9565b9390926040519485947fb006aba000000000000000000000000000000000000000000000000000000000865260048601612f9b565b9798999a9160016131768b61317d94888a611ff9565b90866128f4565b50929d919c909b929a9092918e8e6130b0565b9161305b565b73ffffffffffffffffffffffffffffffffffffffff610525959360609383521660208201528160408201520191611913565b908160209103126101c5575190565b9391909360009460009460005b8181106131f2575050505050565b8481013560f881901c9860019092019788979692909160fc1c9889156139d657506001891461399657600289146137d957600389146137aa57600489146137295760068914613689576005891461363b5760078914613574576008891461351e57600989146133f557600a8914613292577fb2505f7c00000000000000000000000000000000000000000000000000000000600052600489905260246000fd5b909192939495969750600389169788156133e4575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161330a918787611ff9565b6040517f898bd92100000000000000000000000000000000000000000000000000000000815293918491613342918a60048501611a35565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa91821561113857613385936000936133b1575b5060ff909a168091019a613d37565b9080156133ab579061339f91600052602052604060002090565b955b93929190936131e4565b5061339f565b60ff9193506133d69060203d81116133dd575b6133ce8183610404565b8101906131c8565b9290613376565b503d6133c4565b8084013560f81c98506001016132a7565b9091929394959697506003891697881561350d575b8381013560601c90601401909960021c60031660ff1684820135600382901b6101008190039190911c600190911b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01169101908101908161346d918787611ff9565b6040517f13792a4a000000000000000000000000000000000000000000000000000000008152939184916134a5918b60048501611952565b038373ffffffffffffffffffffffffffffffffffffffff8d1691815a93602094fa918215611138576134e7936000936133b1575060ff909a168091019a613d37565b908015613507579061350191600052602052604060002090565b956133a1565b50613501565b8084013560f81c985060010161340a565b98506020870197509495939492939192909182013561353c86613cde565b811461354c575b6134e790613cf8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613543565b975090919293949597600f16968715613629575b6020600061359a6136079a9b86613ba6565b9c9092918a6040516135dd816127b48a82019485603c917f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c8201520190565b51902092604051948594859094939260ff6060936080840197845216602083015260408201520152565b838052039060015afa15611138576134e79060ff6000519a1680910199613be9565b600189019883013560f81c9750613588565b985060208701975094959394929391929091820135808514613661575b6134e790613c9f565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9850613658565b989091929394959662ffffff98506136ab612158600c8416603f9060021c1690565b918215613715575b6003168015613704575b9081906136e8906136e0908781013560e81c906003019c168c01809c8989611ff9565b90898b6131d7565b9111156136fb575b906134e79291613c54565b998201996136f0565b50600281019084013560f01c6136bd565b8482013560f81c92506001909101906136b3565b97509761377f61378c929394959697600f61379493169085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b9290830180938686611ff9565b9086886131d7565b9061350192980198600052602052604060002090565b985096509394929391929091908082013590602001968015613507579061350191600052602052604060002090565b90919293949596975060038916978815613985575b8084013560601c9961384d916014019061380d9060021c600316612158565b9085929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b908101906138b160208c61386385858b8b611ff9565b919073ffffffffffffffffffffffffffffffffffffffff8c604051968795869485937f1626ba7e00000000000000000000000000000000000000000000000000000000855260048501611a35565b0392165afa908115611138577f1626ba7e00000000000000000000000000000000000000000000000000000000917fffffffff0000000000000000000000000000000000000000000000000000000091600091613967575b50160361392357509060ff6134e792991680910199613be9565b6108fb6139348c9389938989611ff9565b906040519485947fb2fed7ae00000000000000000000000000000000000000000000000000000000865260048601613196565b61397f915060203d8111611b7757611b698183610404565b38613909565b8381013560f81c98506001016137ee565b98600f919293949596979850169687156139c5575b60148101976134e79160ff9091169084013560601c613be9565b8281013560f81c97506001016139ab565b98509091929394959698600f16978815613a28575b50602060006139fe6136079a9b86613ba6565b9c90916040519384938c859094939260ff6060936080840197845216602083015260408201520152565b60018101995083013560f81c975060206139eb565b805160209091019060005b818110613a555750505090565b8251845260209384019390920191600101613a48565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613ab1613a9b83610481565b92613aa96040519485610404565b808452610481565b0136602083013760005b8351811015613b8d5780613ad1600192866116f7565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152613b7961012082610404565b519020613b8682856116f7565b5201613abb565b509091506040516127e0816127b4602082018095613a3d565b8101916040602084359401359201601b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84169360ff1c0160ff81116110415791565b90604051907fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208301937f53657175656e6365207369676e65723a0a000000000000000000000000000000855260601b1660318301526045820152604581526127e0606582610404565b916040519160208301937f53657175656e6365206e657374656420636f6e6669673a0a00000000000000008552603884015260588301526078820152607881526127e0609882610404565b60405160208101917f53657175656e636520737461746963206469676573743a0a000000000000000083526038820152603881526127e0605882610404565b6127b46127e061277261276c600060208601511515612cc5565b60405160208101917f53657175656e636520616e792061646472657373207375626469676573743a0a83526040820152604081526127e0606082610404565b91604051917fffffffffffffffffffffffffffffffffffffffff00000000000000000000000060208401947f53657175656e63652073617069656e7420636f6e6669673a0a00000000000000865260601b166039840152604d830152606d820152606d81526127e0608d8261040456fea2646970667358221220fd510fff57502b123cb8cca23f77a513f5d8001f96adc68f13bd2a44f529e4ad64736f6c634300081c0033"
