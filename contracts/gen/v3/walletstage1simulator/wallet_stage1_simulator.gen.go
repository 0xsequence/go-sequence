// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletstage1simulator

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

// SimulatorResult is an auto generated low-level Go binding around an user-defined struct.
type SimulatorResult struct {
	Status  uint8
	Result  []byte
	GasUsed *big.Int
}

// WalletStage1SimulatorMetaData contains all meta data concerning the WalletStage1Simulator contract.
var WalletStage1SimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAGE_2_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSimulator.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"internalType\":\"structSimulator.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60e060405234801561001057600080fd5b5060405161840638038061840683398101604081905261002f916100ca565b808060405161003d906100bd565b604051809103906000f080158015610059573d6000803e3d6000fd5b5060006040518060600160405280602c81526020016183da602c9139604051610087919030906020016100fa565b60408051601f198184030181529190528051602090910120608052506001600160a01b0391821660a0521660c0525061012c9050565b613e1a806145c083390190565b6000602082840312156100dc57600080fd5b81516001600160a01b03811681146100f357600080fd5b9392505050565b6000835160005b8181101561011b5760208187018101518583015201610101565b509190910191825250602001919050565b60805160a05160c05161445e6101626000396000818161046d01526115c8015260006103b201526000610350015261445e6000f3fe60806040526004361061016e5760003560e01c80636ea44577116100cb578063ad55366b1161007f578063ca70785011610059578063ca70785014610552578063f23a6e6114610572578063f727ef1c146105b857610175565b8063ad55366b146104a4578063b93ea7ad146104f7578063bc197c811461050a57610175565b80638c3f5563116100b05780638c3f55631461043b5780639f69ef541461045b578063aaf10f421461048f57610175565b80636ea44577146104075780638943ec021461041a57610175565b8063257671f5116101225780632dd31000116101075780632dd31000146103a05780634fcf3eca146103d45780636406c314146103e757610175565b8063257671f51461033e578063295614261461038057610175565b80631626ba7e116101535780631626ba7e146102c65780631a9b2337146102e65780631f6a1eb91461032b57610175565b8063025b22bc14610238578063150b7a021461024b57610175565b3661017557005b6004361061023657600061019161018c368361342e565b6105d8565b905073ffffffffffffffffffffffffffffffffffffffff811615610234576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101da929190613494565b600060405180830381855af49150503d8060008114610215576040519150601f19603f3d011682016040523d82523d6000602084013e61021a565b606091505b50915091508161022c57805160208201fd5b805160208201f35b505b005b6102366102463660046134cd565b61062c565b34801561025757600080fd5b50610290610266366004613531565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102d257600080fd5b506102906102e13660046135a0565b610678565b3480156102f257600080fd5b5061030661030136600461361a565b610710565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102bd565b610236610339366004613637565b61071b565b34801561034a57600080fd5b506103727f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102bd565b34801561038c57600080fd5b5061023661039b3660046136a8565b61078b565b3480156103ac57600080fd5b506103067f000000000000000000000000000000000000000000000000000000000000000081565b6102366103e236600461361a565b6107cf565b6103fa6103f5366004613637565b610891565b6040516102bd919061375e565b61023661041536600461383c565b6108c3565b34801561042657600080fd5b5061023661043536600461387e565b50505050565b34801561044757600080fd5b506103726104563660046136a8565b610923565b34801561046757600080fd5b506103067f000000000000000000000000000000000000000000000000000000000000000081565b34801561049b57600080fd5b5061030661094f565b3480156104b057600080fd5b506104c46104bf366004613c00565b61095e565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c0016102bd565b610236610505366004613d2e565b61098d565b34801561051657600080fd5b50610290610525366004613da8565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561055e57600080fd5b5061037261056d366004613c00565b610a52565b34801561057e57600080fd5b5061029061058d366004613e6f565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b3480156105c457600080fd5b506102366105d3366004613ee7565b610bbb565b60006106267fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c76565b92915050565b33301461066c576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61067581610cd4565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e08101859052816106cf828686610d28565b509050806106e35750600091506107099050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b6000610626826105d8565b60006107278585610e69565b9050600080610737838686610d28565b9150915081610778578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016106639392919061414d565b6107828184611287565b50505050505050565b3330146107c6576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b61067581611530565b33301461080a576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b6000610815826105d8565b73ffffffffffffffffffffffffffffffffffffffff1603610886576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000082166004820152602401610663565b6106758160006115ec565b6060600061089f8686610e69565b905060006108ac826116ac565b90506108b88183611727565b979650505050505050565b3330146108fe576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b600061090a8383610e69565b90506000610917826116ac565b90506104358183611287565b60006106267f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c76565b6000610959305490565b905090565b600080600080600080610975898989600080611c97565b939d929c5060019b5090995097509095509350505050565b3330146109c8576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b60006109d3836105d8565b73ffffffffffffffffffffffffffffffffffffffff1614610a44576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610663565b610a4e82826115ec565b5050565b600080846101000151516001610a6891906141ac565b67ffffffffffffffff811115610a8057610a806138c0565b604051908082528060200260200182016040528015610aa9578160200160208202803683370190505b50905060005b85610100015151811015610b1b578561010001518181518110610ad457610ad46141bf565b6020026020010151828281518110610aee57610aee6141bf565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610aaf565b5033818661010001515181518110610b3557610b356141bf565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610b6f868686610d28565b50905080610baf578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016106639392919061414d565b50600195945050505050565b333014610bf6576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b610c0f8383836bffffffffffffffffffffffff16611fbc565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c95929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610cdc813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039060200160405180910390a150565b600080600084846000818110610d4057610d406141bf565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e3f57610d9d866116ac565b9150600080610dab8461204b565b9150915042811015610df3576040517f9fa4fe540000000000000000000000000000000000000000000000000000000081526004810185905260248101829052604401610663565b73ffffffffffffffffffffffffffffffffffffffff82161580610e2b575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e3c5760019450505050610e61565b50505b6000806000610e52898989600080611c97565b60019a50985050505050505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610ecd5760006060840152610ede565b84810135606090811c908401526014015b6007600183901c168015610f315760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610f4657506001610f6f565b83602016602003610f625750600282019186013560f01c610f6f565b50600182019186013560f81c5b8067ffffffffffffffff811115610f8857610f886138c0565b60405190808252806020026020018201604052801561101357816020015b6110006040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610fa65790505b50604086015260005b8181101561127c5760018085019489013560f81c90808216900361107b573087604001518381518110611051576110516141bf565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526110c5565b8489013560601c601486018860400151848151811061109c5761109c6141bf565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611103578489013560208601886040015184815181106110ee576110ee6141bf565b60200260200101516020018197508281525050505b6004808216900361119b57600385019489013560e81c89868a61112684836141ac565b92611133939291906141ee565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061117e5761117e6141bf565b60209081029190910101516040015261119781876141ac565b9550505b600880821690036111d9578489013560208601886040015184815181106111c4576111c46141bf565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106111f9576111f96141bf565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061122f5761122f6141bf565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110611264576112646141bf565b602090810291909101015160c001525060010161101c565b505050505092915050565b604081015151600090815b81811015611529576000846040015182815181106112b2576112b26141bf565b602002602001015190508060a0015180156112cb575083155b15611313576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611521565b606081015180158015906113265750805a105b156113635785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161066393929190614218565b600082608001511561139457825161138d9083156113815783611383565b5a5b8560400151612097565b90506113bb565b825160208401516113b8919084156113ac57846113ae565b5a5b86604001516120ad565b90505b806114e45760c083015161140e576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611521565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114785786846114436120c5565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016106639392919061423d565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016114e45760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a1505050611529565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b600101611292565b5050505050565b80611567576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115907fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200160405180910390a16106757f0000000000000000000000000000000000000000000000000000000000000000610cd4565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806116bd8360200151306120e4565b905060006116ca846121b1565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040810151516060906000908067ffffffffffffffff81111561174c5761174c6138c0565b6040519080825280602002602001820160405280156117a257816020015b61178f6040805160608101909152806000815260200160608152602001600081525090565b81526020019060019003908161176a5790505b50925060005b81811015611c8e576000856040015182815181106117c8576117c86141bf565b602002602001015190508060a0015180156117e1575083155b15611829576040805188815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611c86565b6060810151801580159061183c5750805a105b156118d5576005868481518110611855576118556141bf565b6020026020010151600001906005811115611872576118726136c1565b90816005811115611885576118856136c1565b9052505a60405160200161189b91815260200190565b6040516020818303038152906040528684815181106118bc576118bc6141bf565b6020026020010151602001819052505050505050610626565b600082608001511561193a5760005a84519091506119059084156118f957846118fb565b5a5b8660400151612097565b91505a6119129082614268565b888681518110611924576119246141bf565b6020026020010151604001818152505050611994565b60005a845160208601519192506119639185156119575785611959565b5a5b87604001516120ad565b91505a6119709082614268565b888681518110611982576119826141bf565b60200260200101516040018181525050505b80611bd95760c0830151611a5757604080518a815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a160028785815181106119f3576119f36141bf565b6020026020010151600001906005811115611a1057611a106136c1565b90816005811115611a2357611a236136c1565b905250611a2e6120c5565b878581518110611a4057611a406141bf565b602002602001015160200181905250505050611c86565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611afd576003878581518110611a9657611a966141bf565b6020026020010151600001906005811115611ab357611ab36136c1565b90816005811115611ac657611ac66136c1565b905250611ad16120c5565b878581518110611ae357611ae36141bf565b602002602001015160200181905250505050505050610626565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611bd957604080518a8152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a16004878581518110611b7557611b756141bf565b6020026020010151600001906005811115611b9257611b926136c1565b90816005811115611ba557611ba56136c1565b905250611bb06120c5565b878581518110611bc257611bc26141bf565b602002602001015160200181905250505050611c8e565b604080518a8152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a16001878581518110611c2657611c266141bf565b6020026020010151600001906005811115611c4357611c436136c1565b90816005811115611c5657611c566136c1565b905250611c616120c5565b878581518110611c7357611c736141bf565b6020026020010151602001819052505050505b6001016117a8565b50505092915050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611ce2575073ffffffffffffffffffffffffffffffffffffffff8916155b15611dfe578b82013560601c985060149091019089611dfe5760038201918c013560e81c60008d848e611d1585836141ac565b92611d22939291906141ee565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611dad903090859060040161427b565b6040805180830381865afa158015611dc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ded91906142b2565b9250611df982856141ac565b935050505b82600116600103611e3857611e268d8a838f8f87908092611e21939291906141ee565b612407565b97509750975097509750505050611faf565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611ea36001600586901c8116906141ac565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611eee8d6116ac565b9350611f0c8d858e8e86908092611f07939291906141ee565b612653565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611f57575080518614155b8015611f67575080602001518511155b15611fab576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528151600482015260208201516024820152604401610663565b5050505b9550955095509550959050565b6120467fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b600080806120797fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c76565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121545746612157565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000806121c2836101000151613037565b835190915060ff166122355760006121dd84604001516130ba565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611708565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016122c55760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016123355760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016122a7565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016123a55760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e2624602082015290810191909152606081018290526080016122a7565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e64000000000000000000000000000000006044820152606401610663565b6000806000806000612469604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156125e05760038201916000908b013560e81c6124b184826141ac565b9150600090508a82146124c55760006124c7565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146124f857856124fa565b8f5b905061251a818e8e88908792612512939291906141ee565b600186611c97565b50929d50909b50995097508a8a1015612577578c8c8690859261253f939291906141ee565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016106639493929190614303565b829450888e600001510361258a5760008e525b8388106125cd576040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004810189905260248101859052604401610663565b505060c084018790529150849050612491565b8a51158015906125f457508a602001518511155b15612638576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c01516024820152604401610663565b6126418d6116ac565b93505050509550955095509550959050565b60008060005b8381101561302d57600181019085013560f881901c9060fc1c8061277657600f821660008190036126915750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015612726573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000612749828960ff1661313c565b90508b6127565780612765565b60008c81526020829052604090205b9b5050505050505050505050612659565b600181036127da57600f821660008190036127985750600183019287013560f81c5b601484019388013560601c60006127b28260ff851661313c565b9050866127bf57806127ce565b60008781526020829052604090205b96505050505050612659565b600281036129c1576003821660008190036127fc5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d9087926128ba939291906141ee565b6040518463ffffffff1660e01b81526004016128d89392919061432a565b602060405180830381865afa1580156128f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129199190614344565b7fffffffff00000000000000000000000000000000000000000000000000000000161461297c578d8d858e8e6040517ff734863a000000000000000000000000000000000000000000000000000000008152600401610663959493929190614361565b8097508460ff168a0199506000612996858760ff1661313c565b9050896129a357806129b2565b60008a81526020829052604090205b99505050505050505050612659565b600381036129f5576020830192870135846129dc57806129eb565b60008581526020829052604090205b9450505050612659565b60048103612a9b57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c16858301809650819250505060008186019050600080612a6a8e8e8e8e8c908892611f07939291906141ee565b91509150829750818a019950612a8a898260009182526020526040902090565b985082975050505050505050612659565b60068103612b65576003600283901c166000819003612ac15750600183019287013560f81c5b600383166000819003612adb5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff169150809750819250505060008187019050600080612b198f8f8f8f8d908892611f07939291906141ee565b91509150829850848210612b2c57998501995b6000612b398287896131a3565b90508a612b465780612b55565b60008b81526020829052604090205b9a50505050505050505050612659565b60058103612bd2576020830192870135888103612ba0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612bab82613207565b905085612bb85780612bc7565b60008681526020829052604090205b955050505050612659565b60078103612cd857600f82166000819003612bf45750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612704565b60088103612d2c5760208301928701356000612cf48b8261325b565b9050808203612d21577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006127b2826132d6565b60098103612e9257600382166000819003612d4e5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612de8939291906141ee565b6040518463ffffffff1660e01b8152600401612e069392919061414d565b602060405180830381865afa158015612e23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e4791906143b5565b90508197508460ff168a0199506000612e64858760ff1684613311565b905089612e715780612e80565b60008a81526020829052604090205b99508298505050505050505050612659565b600a8103612ff857600382166000819003612eb45750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612f4d939291906141ee565b6040518463ffffffff1660e01b8152600401612f6b9392919061432a565b602060405180830381865afa158015612f88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fac91906143b5565b90508198508560ff168b019a506000612fc9868860ff1684613311565b90508a612fd65780612fe5565b60008b81526020829052604090205b9a50829950505050505050505050612659565b6040517fb2505f7c00000000000000000000000000000000000000000000000000000000815260048101829052602401610663565b5094509492505050565b6000606060005b83518110156130ab578184828151811061305a5761305a6141bf565b60200260200101516040516020016130739291906143ce565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052915060010161303e565b50805160209091012092915050565b6000606060005b83518110156130ab5760006130ee8583815181106130e1576130e16141bf565b602002602001015161337f565b90508281604051602001613103929190614406565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529250506001016130c1565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612193565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b60008061326c8460200151846120e4565b90506000613279856121b1565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a602082015290810182905260009060600161323e565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d016131e8565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161323e98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561348d577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146134c857600080fd5b919050565b6000602082840312156134df57600080fd5b610709826134a4565b60008083601f8401126134fa57600080fd5b50813567ffffffffffffffff81111561351257600080fd5b60208301915083602082850101111561352a57600080fd5b9250929050565b60008060008060006080868803121561354957600080fd5b613552866134a4565b9450613560602087016134a4565b935060408601359250606086013567ffffffffffffffff81111561358357600080fd5b61358f888289016134e8565b969995985093965092949392505050565b6000806000604084860312156135b557600080fd5b83359250602084013567ffffffffffffffff8111156135d357600080fd5b6135df868287016134e8565b9497909650939450505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461067557600080fd5b60006020828403121561362c57600080fd5b8135610709816135ec565b6000806000806040858703121561364d57600080fd5b843567ffffffffffffffff81111561366457600080fd5b613670878288016134e8565b909550935050602085013567ffffffffffffffff81111561369057600080fd5b61369c878288016134e8565b95989497509550505050565b6000602082840312156136ba57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b8381101561370b5781810151838201526020016136f3565b50506000910152565b6000815180845261372c8160208601602086016136f0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613830577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051600681106137f0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261380d6060880182613714565b604092830151979092019690965294506020938401939190910190600101613786565b50929695505050505050565b6000806020838503121561384f57600080fd5b823567ffffffffffffffff81111561386657600080fd5b613872858286016134e8565b90969095509350505050565b6000806000806060858703121561389457600080fd5b61389d856134a4565b935060208501359250604085013567ffffffffffffffff81111561369057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613912576139126138c0565b60405290565b604051610120810167ffffffffffffffff81118282101715613912576139126138c0565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613983576139836138c0565b604052919050565b803560ff811681146134c857600080fd5b803580151581146134c857600080fd5b600067ffffffffffffffff8211156139c6576139c66138c0565b5060051b60200190565b600082601f8301126139e157600080fd5b813567ffffffffffffffff8111156139fb576139fb6138c0565b613a2c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161393c565b818152846020838601011115613a4157600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112613a6f57600080fd5b8135613a82613a7d826139ac565b61393c565b8082825260208201915060208360051b860101925085831115613aa457600080fd5b602085015b83811015613b9157803567ffffffffffffffff811115613ac857600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613afc57600080fd5b613b046138ef565b613b10602083016134a4565b815260408201356020820152606082013567ffffffffffffffff811115613b3657600080fd5b613b458a6020838601016139d0565b60408301525060808201356060820152613b6160a0830161399c565b6080820152613b7260c0830161399c565b60a082015260e0919091013560c0820152835260209283019201613aa9565b5095945050505050565b600082601f830112613bac57600080fd5b8135613bba613a7d826139ac565b8082825260208201915060208360051b860101925085831115613bdc57600080fd5b602085015b83811015613b9157613bf2816134a4565b835260209283019201613be1565b600080600060408486031215613c1557600080fd5b833567ffffffffffffffff811115613c2c57600080fd5b84016101208187031215613c3f57600080fd5b613c47613918565b613c508261398b565b8152613c5e6020830161399c565b6020820152604082013567ffffffffffffffff811115613c7d57600080fd5b613c8988828501613a5e565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff811115613cbd57600080fd5b613cc9888285016139d0565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff811115613cfe57600080fd5b613d0a88828501613b9b565b61010083015250935050602084013567ffffffffffffffff8111156135d357600080fd5b60008060408385031215613d4157600080fd5b8235613d4c816135ec565b9150613d5a602084016134a4565b90509250929050565b60008083601f840112613d7557600080fd5b50813567ffffffffffffffff811115613d8d57600080fd5b6020830191508360208260051b850101111561352a57600080fd5b60008060008060008060008060a0898b031215613dc457600080fd5b613dcd896134a4565b9750613ddb60208a016134a4565b9650604089013567ffffffffffffffff811115613df757600080fd5b613e038b828c01613d63565b909750955050606089013567ffffffffffffffff811115613e2357600080fd5b613e2f8b828c01613d63565b909550935050608089013567ffffffffffffffff811115613e4f57600080fd5b613e5b8b828c016134e8565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613e8857600080fd5b613e91876134a4565b9550613e9f602088016134a4565b94506040870135935060608701359250608087013567ffffffffffffffff811115613ec957600080fd5b613ed589828a016134e8565b979a9699509497509295939492505050565b600080600060608486031215613efc57600080fd5b83359250613f0c602085016134a4565b915060408401356bffffffffffffffffffffffff81168114613f2d57600080fd5b809150509250925092565b600082825180855260208501945060208160051b8301016020850160005b83811015614008577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613fc460e0860182613714565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613f56565b50909695505050505050565b600081518084526020840193506020830160005b8281101561405c57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101614028565b5093949350505050565b805160ff16825260006020820151614082602085018215159052565b506040820151610120604085015261409e610120850182613f38565b9050606083015160608501526080830151608085015260a083015184820360a08601526140cb8282613714565b91505060c083015160c085015260e083015160e08501526101008301518482036101008601526140fb8282614014565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006141606040830186614066565b8281036020840152614173818587614104565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156106265761062661417d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080858511156141fe57600080fd5b8386111561420b57600080fd5b5050820193919092039150565b60608152600061422b6060830186614066565b60208301949094525060400152919050565b6060815260006142506060830186614066565b84602084015282810360408401526141738185613714565b818103818111156106265761062661417d565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006142aa6040830184613714565b949350505050565b600060408284031280156142c557600080fd5b506040805190810167ffffffffffffffff811182821017156142e9576142e96138c0565b604052825181526020928301519281019290925250919050565b606081526000614317606083018688614104565b6020830194909452506040015292915050565b8381526040602082015260006140fb604083018486614104565b60006020828403121561435657600080fd5b8151610709816135ec565b6080815260006143746080830188614066565b86602084015273ffffffffffffffffffffffffffffffffffffffff8616604084015282810360608401526143a9818587614104565b98975050505050505050565b6000602082840312156143c757600080fd5b5051919050565b6040815260006143e16040830185613714565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600083516144188184602088016136f0565b919091019182525060200191905056fea264697066735822122073b4cf763b0e24d195bae515a4551760beb0d5880f80c1795b09761490be6bdb64736f6c634300081b00336080604052348015600f57600080fd5b50613dfb8061001f6000396000f3fe60806040526004361061012d5760003560e01c80638943ec02116100a5578063b93ea7ad11610074578063ca70785011610059578063ca7078501461046a578063f23a6e611461048a578063f727ef1c146104d057610134565b8063b93ea7ad1461040f578063bc197c811461042257610134565b80638943ec02146103665780638c3f556314610387578063aaf10f42146103a7578063ad55366b146103bc57610134565b80631f6a1eb9116100fc5780634fcf3eca116100e15780634fcf3eca1461031d57806351605d80146103305780636ea445771461035357610134565b80631f6a1eb9146102ea57806329561426146102fd57610134565b8063025b22bc146101f7578063150b7a021461020a5780631626ba7e146102855780631a9b2337146102a557610134565b3661013457005b600436106101f557600061015061014b3683612ef3565b6104f0565b905073ffffffffffffffffffffffffffffffffffffffff8116156101f3576000808273ffffffffffffffffffffffffffffffffffffffff16600036604051610199929190612f59565b600060405180830381855af49150503d80600081146101d4576040519150601f19603f3d011682016040523d82523d6000602084013e6101d9565b606091505b5091509150816101eb57805160208201fd5b805160208201f35b505b005b6101f5610205366004612f92565b610544565b34801561021657600080fd5b5061024f610225366004612ff6565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b34801561029157600080fd5b5061024f6102a0366004613065565b610590565b3480156102b157600080fd5b506102c56102c03660046130df565b610628565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161027c565b6101f56102f83660046130fc565b610633565b34801561030957600080fd5b506101f561031836600461316d565b6106b5565b6101f561032b3660046130df565b6106f9565b34801561033c57600080fd5b506103456107bb565b60405190815260200161027c565b6101f5610361366004613186565b6107ea565b34801561037257600080fd5b506101f56103813660046131c8565b50505050565b34801561039357600080fd5b506103456103a236600461316d565b61084a565b3480156103b357600080fd5b506102c5610876565b3480156103c857600080fd5b506103dc6103d736600461354a565b610880565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c00161027c565b6101f561041d366004613678565b6108ba565b34801561042e57600080fd5b5061024f61043d3660046136f2565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561047657600080fd5b5061034561048536600461354a565b61097f565b34801561049657600080fd5b5061024f6104a53660046137b9565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b3480156104dc57600080fd5b506101f56104eb366004613831565b610ae8565b600061053e7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610ba3565b92915050565b333014610584576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61058d81610c01565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e08101859052816105e7828686610c56565b509050806105fb5750600091506106219050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061053e826104f0565b600061063f8585610da4565b9050610653816060015182608001516111c2565b600080610661838686610c56565b91509150816106a2578285856040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161057b93929190613b05565b6106ac81846112aa565b50505050505050565b3330146106f0576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b61058d81611553565b333014610734576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b600061073f826104f0565b73ffffffffffffffffffffffffffffffffffffffff16036107b0576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008216600482015260240161057b565b61058d8160006115e3565b60006107e57fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610825576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b60006108318383610da4565b9050600061083e826116a3565b905061038181836112aa565b600061053e7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610ba3565b60006107e5305490565b60008060008060008061089789898960008061171e565b9399509197509450925090506108ac83611a43565b935093975093979195509350565b3330146108f5576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b6000610900836104f0565b73ffffffffffffffffffffffffffffffffffffffff1614610971576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161057b565b61097b82826115e3565b5050565b6000808461010001515160016109959190613b35565b67ffffffffffffffff8111156109ad576109ad61320a565b6040519080825280602002602001820160405280156109d6578160200160208202803683370190505b50905060005b85610100015151811015610a48578561010001518181518110610a0157610a01613b6f565b6020026020010151828281518110610a1b57610a1b613b6f565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101526001016109dc565b5033818661010001515181518110610a6257610a62613b6f565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610a9c868686610c56565b50905080610adc578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161057b93929190613b05565b50600195945050505050565b333014610b23576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161057b565b610b3c8383836bffffffffffffffffffffffff16611a4e565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610bc2929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c09813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610c6e57610c6e613b6f565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610d6d57610ccb866116a3565b9150600080610cd984611add565b9150915042811015610d21576040517f9fa4fe54000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161057b565b73ffffffffffffffffffffffffffffffffffffffff82161580610d59575073ffffffffffffffffffffffffffffffffffffffff821633145b15610d6a5760019450505050610d9c565b50505b6000806000610d8089898960008061171e565b98509295509093509150610d95905081611a43565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610e085760006060840152610e19565b84810135606090811c908401526014015b6007600183901c168015610e6c5760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610e8157506001610eaa565b83602016602003610e9d5750600282019186013560f01c610eaa565b50600182019186013560f81c5b8067ffffffffffffffff811115610ec357610ec361320a565b604051908082528060200260200182016040528015610f4e57816020015b610f3b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610ee15790505b50604086015260005b818110156111b75760018085019489013560f81c908082169003610fb6573087604001518381518110610f8c57610f8c613b6f565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611000565b8489013560601c6014860188604001518481518110610fd757610fd7613b6f565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b6002808216900361103e5784890135602086018860400151848151811061102957611029613b6f565b60200260200101516020018197508281525050505b600480821690036110d657600385019489013560e81c89868a6110618483613b35565b9261106e93929190613b9e565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106110b9576110b9613b6f565b6020908102919091010151604001526110d28187613b35565b9550505b60088082169003611114578489013560208601886040015184815181106110ff576110ff613b6f565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061113457611134613b6f565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061116a5761116a613b6f565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061119f5761119f613b6f565b602090810291909101015160c0015250600101610f57565b505050505092915050565b60006111cd8361084a565b9050818114611219576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161057b565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561154c576000846040015182815181106112d5576112d5613b6f565b602002602001015190508060a0015180156112ee575083155b15611336576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611544565b606081015180158015906113495750805a105b156113865785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161057b93929190613bc8565b60008260800151156113b75782516113b09083156113a457836113a6565b5a5b8560400151611b29565b90506113de565b825160208401516113db919084156113cf57846113d1565b5a5b8660400151611b3f565b90505b806115075760c0830151611431576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611544565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161149b578684611466611b57565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161057b93929190613bed565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016115075760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a150505061154c565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b6001016112b5565b5050505050565b8061158a576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115b37fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610c4b565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806116b4836020015130611b76565b905060006116c184611c43565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611769575073ffffffffffffffffffffffffffffffffffffffff8916155b15611885578b82013560601c9850601490910190896118855760038201918c013560e81c60008d848e61179c8583613b35565b926117a993929190613b9e565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc891506118349030908590600401613c18565b6040805180830381865afa158015611850573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118749190613c4f565b92506118808285613b35565b935050505b826001166001036118bf576118ad8d8a838f8f879080926118a893929190613b9e565b611e99565b97509750975097509750505050611a36565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c1683820190965092506000905061192a6001600586901c811690613b35565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c16995090920191506119758d6116a3565b93506119938d858e8e8690809261198e93929190613b9e565b6120e5565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d16909152902082519198509650158015906119de575080518614155b80156119ee575080602001518511155b15611a32576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161057b565b5050505b9550955095509550959050565b600061053e82612ac9565b611ad87fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611b0b7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610ba3565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611be65746611be9565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611c54836101000151612afc565b835190915060ff16611cc7576000611c6f8460400151612b7f565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016116ff565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611d575760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611dc75760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001611d39565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01611e375760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001611d39565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e6400000000000000000000000000000000604482015260640161057b565b6000806000806000611efb604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156120725760038201916000908b013560e81c611f438482613b35565b9150600090508a8214611f57576000611f59565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8414611f8a5785611f8c565b8f5b9050611fac818e8e88908792611fa493929190613b9e565b60018661171e565b50929d50909b50995097508a8a1015612009578c8c86908592611fd193929190613b9e565b8c8c6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161057b9493929190613ca0565b829450888e600001510361201c5760008e525b83881061205f576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018990526024810185905260440161057b565b505060c084018790529150849050611f23565b8a511580159061208657508a602001518511155b156120ca576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161057b565b6120d38d6116a3565b93505050509550955095509550959050565b60008060005b83811015612abf57600181019085013560f881901c9060fc1c8061220857600f821660008190036121235750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156121b8573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006121db828960ff16612c01565b90508b6121e857806121f7565b60008c81526020829052604090205b9b50505050505050505050506120eb565b6001810361226c57600f8216600081900361222a5750600183019287013560f81c5b601484019388013560601c60006122448260ff8516612c01565b9050866122515780612260565b60008781526020829052604090205b965050505050506120eb565b600281036124535760038216600081900361228e5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261234c93929190613b9e565b6040518463ffffffff1660e01b815260040161236a93929190613cc7565b602060405180830381865afa158015612387573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123ab9190613ce1565b7fffffffff00000000000000000000000000000000000000000000000000000000161461240e578d8d858e8e6040517ff734863a00000000000000000000000000000000000000000000000000000000815260040161057b959493929190613cfe565b8097508460ff168a0199506000612428858760ff16612c01565b9050896124355780612444565b60008a81526020829052604090205b995050505050505050506120eb565b600381036124875760208301928701358461246e578061247d565b60008581526020829052604090205b94505050506120eb565b6004810361252d57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806124fc8e8e8e8e8c90889261198e93929190613b9e565b91509150829750818a01995061251c898260009182526020526040902090565b9850829750505050505050506120eb565b600681036125f7576003600283901c1660008190036125535750600183019287013560f81c5b60038316600081900361256d5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806125ab8f8f8f8f8d90889261198e93929190613b9e565b915091508298508482106125be57998501995b60006125cb828789612c68565b90508a6125d857806125e7565b60008b81526020829052604090205b9a505050505050505050506120eb565b60058103612664576020830192870135888103612632577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061263d82612ccc565b90508561264a5780612659565b60008681526020829052604090205b9550505050506120eb565b6007810361276a57600f821660008190036126865750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612196565b600881036127be57602083019287013560006127868b82612d20565b90508082036127b3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061224482612d9b565b60098103612924576003821660008190036127e05750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c90879261287a93929190613b9e565b6040518463ffffffff1660e01b815260040161289893929190613b05565b602060405180830381865afa1580156128b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128d99190613d52565b90508197508460ff168a01995060006128f6858760ff1684612dd6565b9050896129035780612912565b60008a81526020829052604090205b995082985050505050505050506120eb565b600a8103612a8a576003821660008190036129465750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d9087926129df93929190613b9e565b6040518463ffffffff1660e01b81526004016129fd93929190613cc7565b602060405180830381865afa158015612a1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a3e9190613d52565b90508198508560ff168b019a506000612a5b868860ff1684612dd6565b90508a612a685780612a77565b60008b81526020829052604090205b9a508299505050505050505050506120eb565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161057b565b5094509492505050565b6000811580159061053e5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b6000606060005b8351811015612b705781848281518110612b1f57612b1f613b6f565b6020026020010151604051602001612b38929190613d6b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612b03565b50805160209091012092915050565b6000606060005b8351811015612b70576000612bb3858381518110612ba657612ba6613b6f565b6020026020010151612e44565b90508281604051602001612bc8929190613da3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612b86565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611c25565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612d31846020015184611b76565b90506000612d3e85611c43565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612d03565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612cad565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612d0398979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015612f52577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612f8d57600080fd5b919050565b600060208284031215612fa457600080fd5b61062182612f69565b60008083601f840112612fbf57600080fd5b50813567ffffffffffffffff811115612fd757600080fd5b602083019150836020828501011115612fef57600080fd5b9250929050565b60008060008060006080868803121561300e57600080fd5b61301786612f69565b945061302560208701612f69565b935060408601359250606086013567ffffffffffffffff81111561304857600080fd5b61305488828901612fad565b969995985093965092949392505050565b60008060006040848603121561307a57600080fd5b83359250602084013567ffffffffffffffff81111561309857600080fd5b6130a486828701612fad565b9497909650939450505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461058d57600080fd5b6000602082840312156130f157600080fd5b8135610621816130b1565b6000806000806040858703121561311257600080fd5b843567ffffffffffffffff81111561312957600080fd5b61313587828801612fad565b909550935050602085013567ffffffffffffffff81111561315557600080fd5b61316187828801612fad565b95989497509550505050565b60006020828403121561317f57600080fd5b5035919050565b6000806020838503121561319957600080fd5b823567ffffffffffffffff8111156131b057600080fd5b6131bc85828601612fad565b90969095509350505050565b600080600080606085870312156131de57600080fd5b6131e785612f69565b935060208501359250604085013567ffffffffffffffff81111561315557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561325c5761325c61320a565b60405290565b604051610120810167ffffffffffffffff8111828210171561325c5761325c61320a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156132cd576132cd61320a565b604052919050565b803560ff81168114612f8d57600080fd5b80358015158114612f8d57600080fd5b600067ffffffffffffffff8211156133105761331061320a565b5060051b60200190565b600082601f83011261332b57600080fd5b813567ffffffffffffffff8111156133455761334561320a565b61337660207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613286565b81815284602083860101111561338b57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126133b957600080fd5b81356133cc6133c7826132f6565b613286565b8082825260208201915060208360051b8601019250858311156133ee57600080fd5b602085015b838110156134db57803567ffffffffffffffff81111561341257600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561344657600080fd5b61344e613239565b61345a60208301612f69565b815260408201356020820152606082013567ffffffffffffffff81111561348057600080fd5b61348f8a60208386010161331a565b604083015250608082013560608201526134ab60a083016132e6565b60808201526134bc60c083016132e6565b60a082015260e0919091013560c08201528352602092830192016133f3565b5095945050505050565b600082601f8301126134f657600080fd5b81356135046133c7826132f6565b8082825260208201915060208360051b86010192508583111561352657600080fd5b602085015b838110156134db5761353c81612f69565b83526020928301920161352b565b60008060006040848603121561355f57600080fd5b833567ffffffffffffffff81111561357657600080fd5b8401610120818703121561358957600080fd5b613591613262565b61359a826132d5565b81526135a8602083016132e6565b6020820152604082013567ffffffffffffffff8111156135c757600080fd5b6135d3888285016133a8565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561360757600080fd5b6136138882850161331a565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561364857600080fd5b613654888285016134e5565b61010083015250935050602084013567ffffffffffffffff81111561309857600080fd5b6000806040838503121561368b57600080fd5b8235613696816130b1565b91506136a460208401612f69565b90509250929050565b60008083601f8401126136bf57600080fd5b50813567ffffffffffffffff8111156136d757600080fd5b6020830191508360208260051b8501011115612fef57600080fd5b60008060008060008060008060a0898b03121561370e57600080fd5b61371789612f69565b975061372560208a01612f69565b9650604089013567ffffffffffffffff81111561374157600080fd5b61374d8b828c016136ad565b909750955050606089013567ffffffffffffffff81111561376d57600080fd5b6137798b828c016136ad565b909550935050608089013567ffffffffffffffff81111561379957600080fd5b6137a58b828c01612fad565b999c989b5096995094979396929594505050565b60008060008060008060a087890312156137d257600080fd5b6137db87612f69565b95506137e960208801612f69565b94506040870135935060608701359250608087013567ffffffffffffffff81111561381357600080fd5b61381f89828a01612fad565b979a9699509497509295939492505050565b60008060006060848603121561384657600080fd5b8335925061385660208501612f69565b915060408401356bffffffffffffffffffffffff8116811461387757600080fd5b809150509250925092565b60005b8381101561389d578181015183820152602001613885565b50506000910152565b600081518084526138be816020860160208601613882565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b838110156139c0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e0604086015261397c60e08601826138a6565b6060838101519087015260808084015115159087015260a08084015115159087015260c092830151929095019190915250602097880197919091019060010161390e565b50909695505050505050565b600081518084526020840193506020830160005b82811015613a1457815173ffffffffffffffffffffffffffffffffffffffff168652602095860195909101906001016139e0565b5093949350505050565b805160ff16825260006020820151613a3a602085018215159052565b5060408201516101206040850152613a566101208501826138f0565b9050606083015160608501526080830151608085015260a083015184820360a0860152613a8382826138a6565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613ab382826139cc565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613b186040830186613a1e565b8281036020840152613b2b818587613abc565b9695505050505050565b8082018082111561053e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008085851115613bae57600080fd5b83861115613bbb57600080fd5b5050820193919092039150565b606081526000613bdb6060830186613a1e565b60208301949094525060400152919050565b606081526000613c006060830186613a1e565b8460208401528281036040840152613b2b81856138a6565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613c4760408301846138a6565b949350505050565b60006040828403128015613c6257600080fd5b506040805190810167ffffffffffffffff81118282101715613c8657613c8661320a565b604052825181526020928301519281019290925250919050565b606081526000613cb4606083018688613abc565b6020830194909452506040015292915050565b838152604060208201526000613ab3604083018486613abc565b600060208284031215613cf357600080fd5b8151610621816130b1565b608081526000613d116080830188613a1e565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152613d46818587613abc565b98975050505050505050565b600060208284031215613d6457600080fd5b5051919050565b604081526000613d7e60408301856138a6565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351613db5818460208801613882565b919091019182525060200191905056fea2646970667358221220354b84e1c13eaed7c517db5658d888dec7d6c0c5da07910274b810754286780864736f6c634300081b0033603e600e3d39601e805130553df33d3d34601c57363d3d373d363d30545af43d82803e903d91601c57fd5bf3",
}

// WalletStage1SimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage1SimulatorMetaData.ABI instead.
var WalletStage1SimulatorABI = WalletStage1SimulatorMetaData.ABI

// WalletStage1SimulatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage1SimulatorMetaData.Bin instead.
var WalletStage1SimulatorBin = WalletStage1SimulatorMetaData.Bin

// DeployWalletStage1Simulator deploys a new Ethereum contract, binding an instance of WalletStage1Simulator to it.
func DeployWalletStage1Simulator(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *WalletStage1Simulator, error) {
	parsed, err := WalletStage1SimulatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage1SimulatorBin), backend, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletStage1Simulator{WalletStage1SimulatorCaller: WalletStage1SimulatorCaller{contract: contract}, WalletStage1SimulatorTransactor: WalletStage1SimulatorTransactor{contract: contract}, WalletStage1SimulatorFilterer: WalletStage1SimulatorFilterer{contract: contract}}, nil
}

// WalletStage1SimulatorDeployedBin is the resulting bytecode of the created contract
const WalletStage1SimulatorDeployedBin = "0x60806040526004361061016e5760003560e01c80636ea44577116100cb578063ad55366b1161007f578063ca70785011610059578063ca70785014610552578063f23a6e6114610572578063f727ef1c146105b857610175565b8063ad55366b146104a4578063b93ea7ad146104f7578063bc197c811461050a57610175565b80638c3f5563116100b05780638c3f55631461043b5780639f69ef541461045b578063aaf10f421461048f57610175565b80636ea44577146104075780638943ec021461041a57610175565b8063257671f5116101225780632dd31000116101075780632dd31000146103a05780634fcf3eca146103d45780636406c314146103e757610175565b8063257671f51461033e578063295614261461038057610175565b80631626ba7e116101535780631626ba7e146102c65780631a9b2337146102e65780631f6a1eb91461032b57610175565b8063025b22bc14610238578063150b7a021461024b57610175565b3661017557005b6004361061023657600061019161018c368361342e565b6105d8565b905073ffffffffffffffffffffffffffffffffffffffff811615610234576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101da929190613494565b600060405180830381855af49150503d8060008114610215576040519150601f19603f3d011682016040523d82523d6000602084013e61021a565b606091505b50915091508161022c57805160208201fd5b805160208201f35b505b005b6102366102463660046134cd565b61062c565b34801561025757600080fd5b50610290610266366004613531565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102d257600080fd5b506102906102e13660046135a0565b610678565b3480156102f257600080fd5b5061030661030136600461361a565b610710565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102bd565b610236610339366004613637565b61071b565b34801561034a57600080fd5b506103727f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016102bd565b34801561038c57600080fd5b5061023661039b3660046136a8565b61078b565b3480156103ac57600080fd5b506103067f000000000000000000000000000000000000000000000000000000000000000081565b6102366103e236600461361a565b6107cf565b6103fa6103f5366004613637565b610891565b6040516102bd919061375e565b61023661041536600461383c565b6108c3565b34801561042657600080fd5b5061023661043536600461387e565b50505050565b34801561044757600080fd5b506103726104563660046136a8565b610923565b34801561046757600080fd5b506103067f000000000000000000000000000000000000000000000000000000000000000081565b34801561049b57600080fd5b5061030661094f565b3480156104b057600080fd5b506104c46104bf366004613c00565b61095e565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c0016102bd565b610236610505366004613d2e565b61098d565b34801561051657600080fd5b50610290610525366004613da8565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561055e57600080fd5b5061037261056d366004613c00565b610a52565b34801561057e57600080fd5b5061029061058d366004613e6f565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b3480156105c457600080fd5b506102366105d3366004613ee7565b610bbb565b60006106267fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c76565b92915050565b33301461066c576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b61067581610cd4565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e08101859052816106cf828686610d28565b509050806106e35750600091506107099050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b6000610626826105d8565b60006107278585610e69565b9050600080610737838686610d28565b9150915081610778578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016106639392919061414d565b6107828184611287565b50505050505050565b3330146107c6576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b61067581611530565b33301461080a576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b6000610815826105d8565b73ffffffffffffffffffffffffffffffffffffffff1603610886576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000082166004820152602401610663565b6106758160006115ec565b6060600061089f8686610e69565b905060006108ac826116ac565b90506108b88183611727565b979650505050505050565b3330146108fe576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b600061090a8383610e69565b90506000610917826116ac565b90506104358183611287565b60006106267f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c76565b6000610959305490565b905090565b600080600080600080610975898989600080611c97565b939d929c5060019b5090995097509095509350505050565b3330146109c8576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b60006109d3836105d8565b73ffffffffffffffffffffffffffffffffffffffff1614610a44576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000083166004820152602401610663565b610a4e82826115ec565b5050565b600080846101000151516001610a6891906141ac565b67ffffffffffffffff811115610a8057610a806138c0565b604051908082528060200260200182016040528015610aa9578160200160208202803683370190505b50905060005b85610100015151811015610b1b578561010001518181518110610ad457610ad46141bf565b6020026020010151828281518110610aee57610aee6141bf565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610aaf565b5033818661010001515181518110610b3557610b356141bf565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610b6f868686610d28565b50905080610baf578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016106639392919061414d565b50600195945050505050565b333014610bf6576040517fa19dbf00000000000000000000000000000000000000000000000000000000008152336004820152602401610663565b610c0f8383836bffffffffffffffffffffffff16611fbc565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c95929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610cdc813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039060200160405180910390a150565b600080600084846000818110610d4057610d406141bf565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e3f57610d9d866116ac565b9150600080610dab8461204b565b9150915042811015610df3576040517f9fa4fe540000000000000000000000000000000000000000000000000000000081526004810185905260248101829052604401610663565b73ffffffffffffffffffffffffffffffffffffffff82161580610e2b575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e3c5760019450505050610e61565b50505b6000806000610e52898989600080611c97565b60019a50985050505050505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610ecd5760006060840152610ede565b84810135606090811c908401526014015b6007600183901c168015610f315760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610f4657506001610f6f565b83602016602003610f625750600282019186013560f01c610f6f565b50600182019186013560f81c5b8067ffffffffffffffff811115610f8857610f886138c0565b60405190808252806020026020018201604052801561101357816020015b6110006040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610fa65790505b50604086015260005b8181101561127c5760018085019489013560f81c90808216900361107b573087604001518381518110611051576110516141bf565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526110c5565b8489013560601c601486018860400151848151811061109c5761109c6141bf565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611103578489013560208601886040015184815181106110ee576110ee6141bf565b60200260200101516020018197508281525050505b6004808216900361119b57600385019489013560e81c89868a61112684836141ac565b92611133939291906141ee565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061117e5761117e6141bf565b60209081029190910101516040015261119781876141ac565b9550505b600880821690036111d9578489013560208601886040015184815181106111c4576111c46141bf565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106111f9576111f96141bf565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061122f5761122f6141bf565b602090810291909101015190151560a090910152604087015180516003600684901c16919084908110611264576112646141bf565b602090810291909101015160c001525060010161101c565b505050505092915050565b604081015151600090815b81811015611529576000846040015182815181106112b2576112b26141bf565b602002602001015190508060a0015180156112cb575083155b15611313576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611521565b606081015180158015906113265750805a105b156113635785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161066393929190614218565b600082608001511561139457825161138d9083156113815783611383565b5a5b8560400151612097565b90506113bb565b825160208401516113b8919084156113ac57846113ae565b5a5b86604001516120ad565b90505b806114e45760c083015161140e576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611521565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114785786846114436120c5565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016106639392919061423d565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016114e45760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a1505050611529565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b600101611292565b5050505050565b80611567576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115907fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200160405180910390a16106757f0000000000000000000000000000000000000000000000000000000000000000610cd4565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806116bd8360200151306120e4565b905060006116ca846121b1565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040810151516060906000908067ffffffffffffffff81111561174c5761174c6138c0565b6040519080825280602002602001820160405280156117a257816020015b61178f6040805160608101909152806000815260200160608152602001600081525090565b81526020019060019003908161176a5790505b50925060005b81811015611c8e576000856040015182815181106117c8576117c86141bf565b602002602001015190508060a0015180156117e1575083155b15611829576040805188815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611c86565b6060810151801580159061183c5750805a105b156118d5576005868481518110611855576118556141bf565b6020026020010151600001906005811115611872576118726136c1565b90816005811115611885576118856136c1565b9052505a60405160200161189b91815260200190565b6040516020818303038152906040528684815181106118bc576118bc6141bf565b6020026020010151602001819052505050505050610626565b600082608001511561193a5760005a84519091506119059084156118f957846118fb565b5a5b8660400151612097565b91505a6119129082614268565b888681518110611924576119246141bf565b6020026020010151604001818152505050611994565b60005a845160208601519192506119639185156119575785611959565b5a5b87604001516120ad565b91505a6119709082614268565b888681518110611982576119826141bf565b60200260200101516040018181525050505b80611bd95760c0830151611a5757604080518a815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a160028785815181106119f3576119f36141bf565b6020026020010151600001906005811115611a1057611a106136c1565b90816005811115611a2357611a236136c1565b905250611a2e6120c5565b878581518110611a4057611a406141bf565b602002602001015160200181905250505050611c86565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611afd576003878581518110611a9657611a966141bf565b6020026020010151600001906005811115611ab357611ab36136c1565b90816005811115611ac657611ac66136c1565b905250611ad16120c5565b878581518110611ae357611ae36141bf565b602002602001015160200181905250505050505050610626565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611bd957604080518a8152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a16004878581518110611b7557611b756141bf565b6020026020010151600001906005811115611b9257611b926136c1565b90816005811115611ba557611ba56136c1565b905250611bb06120c5565b878581518110611bc257611bc26141bf565b602002602001015160200181905250505050611c8e565b604080518a8152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a16001878581518110611c2657611c266141bf565b6020026020010151600001906005811115611c4357611c436136c1565b90816005811115611c5657611c566136c1565b905250611c616120c5565b878581518110611c7357611c736141bf565b6020026020010151602001819052505050505b6001016117a8565b50505092915050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611ce2575073ffffffffffffffffffffffffffffffffffffffff8916155b15611dfe578b82013560601c985060149091019089611dfe5760038201918c013560e81c60008d848e611d1585836141ac565b92611d22939291906141ee565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611dad903090859060040161427b565b6040805180830381865afa158015611dc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ded91906142b2565b9250611df982856141ac565b935050505b82600116600103611e3857611e268d8a838f8f87908092611e21939291906141ee565b612407565b97509750975097509750505050611faf565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611ea36001600586901c8116906141ac565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611eee8d6116ac565b9350611f0c8d858e8e86908092611f07939291906141ee565b612653565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611f57575080518614155b8015611f67575080602001518511155b15611fab576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528151600482015260208201516024820152604401610663565b5050505b9550955095509550959050565b6120467fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b600080806120797fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c76565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121545746612157565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b6000806121c2836101000151613037565b835190915060ff166122355760006121dd84604001516130ba565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611708565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016122c55760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016123355760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016122a7565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016123a55760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e2624602082015290810191909152606081018290526080016122a7565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e64000000000000000000000000000000006044820152606401610663565b6000806000806000612469604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156125e05760038201916000908b013560e81c6124b184826141ac565b9150600090508a82146124c55760006124c7565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146124f857856124fa565b8f5b905061251a818e8e88908792612512939291906141ee565b600186611c97565b50929d50909b50995097508a8a1015612577578c8c8690859261253f939291906141ee565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016106639493929190614303565b829450888e600001510361258a5760008e525b8388106125cd576040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004810189905260248101859052604401610663565b505060c084018790529150849050612491565b8a51158015906125f457508a602001518511155b15612638576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c01516024820152604401610663565b6126418d6116ac565b93505050509550955095509550959050565b60008060005b8381101561302d57600181019085013560f881901c9060fc1c8061277657600f821660008190036126915750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015612726573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000612749828960ff1661313c565b90508b6127565780612765565b60008c81526020829052604090205b9b5050505050505050505050612659565b600181036127da57600f821660008190036127985750600183019287013560f81c5b601484019388013560601c60006127b28260ff851661313c565b9050866127bf57806127ce565b60008781526020829052604090205b96505050505050612659565b600281036129c1576003821660008190036127fc5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d9087926128ba939291906141ee565b6040518463ffffffff1660e01b81526004016128d89392919061432a565b602060405180830381865afa1580156128f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129199190614344565b7fffffffff00000000000000000000000000000000000000000000000000000000161461297c578d8d858e8e6040517ff734863a000000000000000000000000000000000000000000000000000000008152600401610663959493929190614361565b8097508460ff168a0199506000612996858760ff1661313c565b9050896129a357806129b2565b60008a81526020829052604090205b99505050505050505050612659565b600381036129f5576020830192870135846129dc57806129eb565b60008581526020829052604090205b9450505050612659565b60048103612a9b57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c16858301809650819250505060008186019050600080612a6a8e8e8e8e8c908892611f07939291906141ee565b91509150829750818a019950612a8a898260009182526020526040902090565b985082975050505050505050612659565b60068103612b65576003600283901c166000819003612ac15750600183019287013560f81c5b600383166000819003612adb5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff169150809750819250505060008187019050600080612b198f8f8f8f8d908892611f07939291906141ee565b91509150829850848210612b2c57998501995b6000612b398287896131a3565b90508a612b465780612b55565b60008b81526020829052604090205b9a50505050505050505050612659565b60058103612bd2576020830192870135888103612ba0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612bab82613207565b905085612bb85780612bc7565b60008681526020829052604090205b955050505050612659565b60078103612cd857600f82166000819003612bf45750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612704565b60088103612d2c5760208301928701356000612cf48b8261325b565b9050808203612d21577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006127b2826132d6565b60098103612e9257600382166000819003612d4e5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612de8939291906141ee565b6040518463ffffffff1660e01b8152600401612e069392919061414d565b602060405180830381865afa158015612e23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e4791906143b5565b90508197508460ff168a0199506000612e64858760ff1684613311565b905089612e715780612e80565b60008a81526020829052604090205b99508298505050505050505050612659565b600a8103612ff857600382166000819003612eb45750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612f4d939291906141ee565b6040518463ffffffff1660e01b8152600401612f6b9392919061432a565b602060405180830381865afa158015612f88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fac91906143b5565b90508198508560ff168b019a506000612fc9868860ff1684613311565b90508a612fd65780612fe5565b60008b81526020829052604090205b9a50829950505050505050505050612659565b6040517fb2505f7c00000000000000000000000000000000000000000000000000000000815260048101829052602401610663565b5094509492505050565b6000606060005b83518110156130ab578184828151811061305a5761305a6141bf565b60200260200101516040516020016130739291906143ce565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052915060010161303e565b50805160209091012092915050565b6000606060005b83518110156130ab5760006130ee8583815181106130e1576130e16141bf565b602002602001015161337f565b90508281604051602001613103929190614406565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529250506001016130c1565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612193565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b60008061326c8460200151846120e4565b90506000613279856121b1565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a602082015290810182905260009060600161323e565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d016131e8565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161323e98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561348d577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146134c857600080fd5b919050565b6000602082840312156134df57600080fd5b610709826134a4565b60008083601f8401126134fa57600080fd5b50813567ffffffffffffffff81111561351257600080fd5b60208301915083602082850101111561352a57600080fd5b9250929050565b60008060008060006080868803121561354957600080fd5b613552866134a4565b9450613560602087016134a4565b935060408601359250606086013567ffffffffffffffff81111561358357600080fd5b61358f888289016134e8565b969995985093965092949392505050565b6000806000604084860312156135b557600080fd5b83359250602084013567ffffffffffffffff8111156135d357600080fd5b6135df868287016134e8565b9497909650939450505050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461067557600080fd5b60006020828403121561362c57600080fd5b8135610709816135ec565b6000806000806040858703121561364d57600080fd5b843567ffffffffffffffff81111561366457600080fd5b613670878288016134e8565b909550935050602085013567ffffffffffffffff81111561369057600080fd5b61369c878288016134e8565b95989497509550505050565b6000602082840312156136ba57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b8381101561370b5781810151838201526020016136f3565b50506000910152565b6000815180845261372c8160208601602086016136f0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613830577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281518051600681106137f0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261380d6060880182613714565b604092830151979092019690965294506020938401939190910190600101613786565b50929695505050505050565b6000806020838503121561384f57600080fd5b823567ffffffffffffffff81111561386657600080fd5b613872858286016134e8565b90969095509350505050565b6000806000806060858703121561389457600080fd5b61389d856134a4565b935060208501359250604085013567ffffffffffffffff81111561369057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613912576139126138c0565b60405290565b604051610120810167ffffffffffffffff81118282101715613912576139126138c0565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613983576139836138c0565b604052919050565b803560ff811681146134c857600080fd5b803580151581146134c857600080fd5b600067ffffffffffffffff8211156139c6576139c66138c0565b5060051b60200190565b600082601f8301126139e157600080fd5b813567ffffffffffffffff8111156139fb576139fb6138c0565b613a2c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161393c565b818152846020838601011115613a4157600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112613a6f57600080fd5b8135613a82613a7d826139ac565b61393c565b8082825260208201915060208360051b860101925085831115613aa457600080fd5b602085015b83811015613b9157803567ffffffffffffffff811115613ac857600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613afc57600080fd5b613b046138ef565b613b10602083016134a4565b815260408201356020820152606082013567ffffffffffffffff811115613b3657600080fd5b613b458a6020838601016139d0565b60408301525060808201356060820152613b6160a0830161399c565b6080820152613b7260c0830161399c565b60a082015260e0919091013560c0820152835260209283019201613aa9565b5095945050505050565b600082601f830112613bac57600080fd5b8135613bba613a7d826139ac565b8082825260208201915060208360051b860101925085831115613bdc57600080fd5b602085015b83811015613b9157613bf2816134a4565b835260209283019201613be1565b600080600060408486031215613c1557600080fd5b833567ffffffffffffffff811115613c2c57600080fd5b84016101208187031215613c3f57600080fd5b613c47613918565b613c508261398b565b8152613c5e6020830161399c565b6020820152604082013567ffffffffffffffff811115613c7d57600080fd5b613c8988828501613a5e565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff811115613cbd57600080fd5b613cc9888285016139d0565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff811115613cfe57600080fd5b613d0a88828501613b9b565b61010083015250935050602084013567ffffffffffffffff8111156135d357600080fd5b60008060408385031215613d4157600080fd5b8235613d4c816135ec565b9150613d5a602084016134a4565b90509250929050565b60008083601f840112613d7557600080fd5b50813567ffffffffffffffff811115613d8d57600080fd5b6020830191508360208260051b850101111561352a57600080fd5b60008060008060008060008060a0898b031215613dc457600080fd5b613dcd896134a4565b9750613ddb60208a016134a4565b9650604089013567ffffffffffffffff811115613df757600080fd5b613e038b828c01613d63565b909750955050606089013567ffffffffffffffff811115613e2357600080fd5b613e2f8b828c01613d63565b909550935050608089013567ffffffffffffffff811115613e4f57600080fd5b613e5b8b828c016134e8565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613e8857600080fd5b613e91876134a4565b9550613e9f602088016134a4565b94506040870135935060608701359250608087013567ffffffffffffffff811115613ec957600080fd5b613ed589828a016134e8565b979a9699509497509295939492505050565b600080600060608486031215613efc57600080fd5b83359250613f0c602085016134a4565b915060408401356bffffffffffffffffffffffff81168114613f2d57600080fd5b809150509250925092565b600082825180855260208501945060208160051b8301016020850160005b83811015614008577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613fc460e0860182613714565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613f56565b50909695505050505050565b600081518084526020840193506020830160005b8281101561405c57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101614028565b5093949350505050565b805160ff16825260006020820151614082602085018215159052565b506040820151610120604085015261409e610120850182613f38565b9050606083015160608501526080830151608085015260a083015184820360a08601526140cb8282613714565b91505060c083015160c085015260e083015160e08501526101008301518482036101008601526140fb8282614014565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006141606040830186614066565b8281036020840152614173818587614104565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156106265761062661417d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080858511156141fe57600080fd5b8386111561420b57600080fd5b5050820193919092039150565b60608152600061422b6060830186614066565b60208301949094525060400152919050565b6060815260006142506060830186614066565b84602084015282810360408401526141738185613714565b818103818111156106265761062661417d565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006142aa6040830184613714565b949350505050565b600060408284031280156142c557600080fd5b506040805190810167ffffffffffffffff811182821017156142e9576142e96138c0565b604052825181526020928301519281019290925250919050565b606081526000614317606083018688614104565b6020830194909452506040015292915050565b8381526040602082015260006140fb604083018486614104565b60006020828403121561435657600080fd5b8151610709816135ec565b6080815260006143746080830188614066565b86602084015273ffffffffffffffffffffffffffffffffffffffff8616604084015282810360608401526143a9818587614104565b98975050505050505050565b6000602082840312156143c757600080fd5b5051919050565b6040815260006143e16040830185613714565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600083516144188184602088016136f0565b919091019182525060200191905056fea264697066735822122073b4cf763b0e24d195bae515a4551760beb0d5880f80c1795b09761490be6bdb64736f6c634300081b0033"

// WalletStage1Simulator is an auto generated Go binding around an Ethereum contract.
type WalletStage1Simulator struct {
	WalletStage1SimulatorCaller     // Read-only binding to the contract
	WalletStage1SimulatorTransactor // Write-only binding to the contract
	WalletStage1SimulatorFilterer   // Log filterer for contract events
}

// WalletStage1SimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletStage1SimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1SimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletStage1SimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1SimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletStage1SimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage1SimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletStage1SimulatorSession struct {
	Contract     *WalletStage1Simulator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletStage1SimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletStage1SimulatorCallerSession struct {
	Contract *WalletStage1SimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WalletStage1SimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletStage1SimulatorTransactorSession struct {
	Contract     *WalletStage1SimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WalletStage1SimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletStage1SimulatorRaw struct {
	Contract *WalletStage1Simulator // Generic contract binding to access the raw methods on
}

// WalletStage1SimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletStage1SimulatorCallerRaw struct {
	Contract *WalletStage1SimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// WalletStage1SimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletStage1SimulatorTransactorRaw struct {
	Contract *WalletStage1SimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletStage1Simulator creates a new instance of WalletStage1Simulator, bound to a specific deployed contract.
func NewWalletStage1Simulator(address common.Address, backend bind.ContractBackend) (*WalletStage1Simulator, error) {
	contract, err := bindWalletStage1Simulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletStage1Simulator{WalletStage1SimulatorCaller: WalletStage1SimulatorCaller{contract: contract}, WalletStage1SimulatorTransactor: WalletStage1SimulatorTransactor{contract: contract}, WalletStage1SimulatorFilterer: WalletStage1SimulatorFilterer{contract: contract}}, nil
}

// NewWalletStage1SimulatorCaller creates a new read-only instance of WalletStage1Simulator, bound to a specific deployed contract.
func NewWalletStage1SimulatorCaller(address common.Address, caller bind.ContractCaller) (*WalletStage1SimulatorCaller, error) {
	contract, err := bindWalletStage1Simulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorCaller{contract: contract}, nil
}

// NewWalletStage1SimulatorTransactor creates a new write-only instance of WalletStage1Simulator, bound to a specific deployed contract.
func NewWalletStage1SimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletStage1SimulatorTransactor, error) {
	contract, err := bindWalletStage1Simulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorTransactor{contract: contract}, nil
}

// NewWalletStage1SimulatorFilterer creates a new log filterer instance of WalletStage1Simulator, bound to a specific deployed contract.
func NewWalletStage1SimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletStage1SimulatorFilterer, error) {
	contract, err := bindWalletStage1Simulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorFilterer{contract: contract}, nil
}

// bindWalletStage1Simulator binds a generic wrapper to an already deployed contract.
func bindWalletStage1Simulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletStage1SimulatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage1Simulator *WalletStage1SimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage1Simulator.Contract.WalletStage1SimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage1Simulator *WalletStage1SimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.WalletStage1SimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage1Simulator *WalletStage1SimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.WalletStage1SimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage1Simulator *WalletStage1SimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage1Simulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.contract.Transact(opts, method, params...)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) FACTORY() (common.Address, error) {
	return _WalletStage1Simulator.Contract.FACTORY(&_WalletStage1Simulator.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) FACTORY() (common.Address, error) {
	return _WalletStage1Simulator.Contract.FACTORY(&_WalletStage1Simulator.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) INITCODEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "INIT_CODE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) INITCODEHASH() ([32]byte, error) {
	return _WalletStage1Simulator.Contract.INITCODEHASH(&_WalletStage1Simulator.CallOpts)
}

// INITCODEHASH is a free data retrieval call binding the contract method 0x257671f5.
//
// Solidity: function INIT_CODE_HASH() view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) INITCODEHASH() ([32]byte, error) {
	return _WalletStage1Simulator.Contract.INITCODEHASH(&_WalletStage1Simulator.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) STAGE2IMPLEMENTATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "STAGE_2_IMPLEMENTATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletStage1Simulator.Contract.STAGE2IMPLEMENTATION(&_WalletStage1Simulator.CallOpts)
}

// STAGE2IMPLEMENTATION is a free data retrieval call binding the contract method 0x9f69ef54.
//
// Solidity: function STAGE_2_IMPLEMENTATION() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) STAGE2IMPLEMENTATION() (common.Address, error) {
	return _WalletStage1Simulator.Contract.STAGE2IMPLEMENTATION(&_WalletStage1Simulator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) GetImplementation() (common.Address, error) {
	return _WalletStage1Simulator.Contract.GetImplementation(&_WalletStage1Simulator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) GetImplementation() (common.Address, error) {
	return _WalletStage1Simulator.Contract.GetImplementation(&_WalletStage1Simulator.CallOpts)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) IsValidSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "isValidSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage1Simulator.Contract.IsValidSapientSignature(&_WalletStage1Simulator.CallOpts, _payload, _signature)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage1Simulator.Contract.IsValidSapientSignature(&_WalletStage1Simulator.CallOpts, _payload, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.IsValidSignature(&_WalletStage1Simulator.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.IsValidSignature(&_WalletStage1Simulator.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC1155BatchReceived(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC1155BatchReceived(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC1155Received(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC1155Received(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC721Received(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage1Simulator.Contract.OnERC721Received(&_WalletStage1Simulator.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1Simulator.Contract.ReadHook(&_WalletStage1Simulator.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1Simulator.Contract.ReadHook(&_WalletStage1Simulator.CallOpts, signature)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage1Simulator.Contract.ReadNonce(&_WalletStage1Simulator.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage1Simulator.Contract.ReadNonce(&_WalletStage1Simulator.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage1Simulator *WalletStage1SimulatorCaller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletStage1Simulator.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletStage1Simulator *WalletStage1SimulatorSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage1Simulator.Contract.RecoverPartialSignature(&_WalletStage1Simulator.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage1Simulator *WalletStage1SimulatorCallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage1Simulator.Contract.RecoverPartialSignature(&_WalletStage1Simulator.CallOpts, _payload, _signature)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.AddHook(&_WalletStage1Simulator.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.AddHook(&_WalletStage1Simulator.TransactOpts, signature, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Execute(&_WalletStage1Simulator.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Execute(&_WalletStage1Simulator.TransactOpts, _payload, _signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.RemoveHook(&_WalletStage1Simulator.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.RemoveHook(&_WalletStage1Simulator.TransactOpts, signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.SelfExecute(&_WalletStage1Simulator.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.SelfExecute(&_WalletStage1Simulator.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.SetStaticSignature(&_WalletStage1Simulator.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.SetStaticSignature(&_WalletStage1Simulator.TransactOpts, _hash, _address, _timestamp)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) Simulate(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "simulate", _payload, _signature)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage1Simulator *WalletStage1SimulatorSession) Simulate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Simulate(&_WalletStage1Simulator.TransactOpts, _payload, _signature)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) Simulate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Simulate(&_WalletStage1Simulator.TransactOpts, _payload, _signature)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.TokenReceived(&_WalletStage1Simulator.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.TokenReceived(&_WalletStage1Simulator.TransactOpts, arg0, arg1, arg2)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.UpdateImageHash(&_WalletStage1Simulator.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.UpdateImageHash(&_WalletStage1Simulator.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.UpdateImplementation(&_WalletStage1Simulator.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.UpdateImplementation(&_WalletStage1Simulator.TransactOpts, _implementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Fallback(&_WalletStage1Simulator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Fallback(&_WalletStage1Simulator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage1Simulator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorSession) Receive() (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Receive(&_WalletStage1Simulator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage1Simulator *WalletStage1SimulatorTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletStage1Simulator.Contract.Receive(&_WalletStage1Simulator.TransactOpts)
}

// WalletStage1SimulatorCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallAbortedIterator struct {
	Event *WalletStage1SimulatorCallAborted // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorCallAborted)
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
		it.Event = new(WalletStage1SimulatorCallAborted)
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
func (it *WalletStage1SimulatorCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorCallAborted represents a CallAborted event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallAborted struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletStage1SimulatorCallAbortedIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorCallAbortedIterator{contract: _WalletStage1Simulator.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorCallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorCallAborted)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
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

// ParseCallAborted is a log parse operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseCallAborted(log types.Log) (*WalletStage1SimulatorCallAborted, error) {
	event := new(WalletStage1SimulatorCallAborted)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallFailedIterator struct {
	Event *WalletStage1SimulatorCallFailed // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorCallFailed)
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
		it.Event = new(WalletStage1SimulatorCallFailed)
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
func (it *WalletStage1SimulatorCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorCallFailed represents a CallFailed event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallFailed struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletStage1SimulatorCallFailedIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorCallFailedIterator{contract: _WalletStage1Simulator.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorCallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorCallFailed)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
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

// ParseCallFailed is a log parse operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseCallFailed(log types.Log) (*WalletStage1SimulatorCallFailed, error) {
	event := new(WalletStage1SimulatorCallFailed)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallSkippedIterator struct {
	Event *WalletStage1SimulatorCallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorCallSkipped)
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
		it.Event = new(WalletStage1SimulatorCallSkipped)
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
func (it *WalletStage1SimulatorCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorCallSkipped represents a CallSkipped event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletStage1SimulatorCallSkippedIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorCallSkippedIterator{contract: _WalletStage1Simulator.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorCallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorCallSkipped)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseCallSkipped(log types.Log) (*WalletStage1SimulatorCallSkipped, error) {
	event := new(WalletStage1SimulatorCallSkipped)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorCallSuccessIterator is returned from FilterCallSuccess and is used to iterate over the raw logs and unpacked data for CallSuccess events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallSuccessIterator struct {
	Event *WalletStage1SimulatorCallSuccess // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorCallSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorCallSuccess)
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
		it.Event = new(WalletStage1SimulatorCallSuccess)
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
func (it *WalletStage1SimulatorCallSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorCallSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorCallSuccess represents a CallSuccess event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorCallSuccess struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSuccess is a free log retrieval operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterCallSuccess(opts *bind.FilterOpts) (*WalletStage1SimulatorCallSuccessIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorCallSuccessIterator{contract: _WalletStage1Simulator.contract, event: "CallSuccess", logs: logs, sub: sub}, nil
}

// WatchCallSuccess is a free log subscription operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchCallSuccess(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorCallSuccess) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorCallSuccess)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
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

// ParseCallSuccess is a log parse operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseCallSuccess(log types.Log) (*WalletStage1SimulatorCallSuccess, error) {
	event := new(WalletStage1SimulatorCallSuccess)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorDefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorDefinedHookIterator struct {
	Event *WalletStage1SimulatorDefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorDefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorDefinedHook)
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
		it.Event = new(WalletStage1SimulatorDefinedHook)
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
func (it *WalletStage1SimulatorDefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorDefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorDefinedHook represents a DefinedHook event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorDefinedHook struct {
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage1SimulatorDefinedHookIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorDefinedHookIterator{contract: _WalletStage1Simulator.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorDefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorDefinedHook)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseDefinedHook(log types.Log) (*WalletStage1SimulatorDefinedHook, error) {
	event := new(WalletStage1SimulatorDefinedHook)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorImageHashUpdatedIterator struct {
	Event *WalletStage1SimulatorImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorImageHashUpdated)
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
		it.Event = new(WalletStage1SimulatorImageHashUpdated)
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
func (it *WalletStage1SimulatorImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorImageHashUpdated represents a ImageHashUpdated event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletStage1SimulatorImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorImageHashUpdatedIterator{contract: _WalletStage1Simulator.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorImageHashUpdated)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseImageHashUpdated(log types.Log) (*WalletStage1SimulatorImageHashUpdated, error) {
	event := new(WalletStage1SimulatorImageHashUpdated)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorImplementationUpdatedIterator struct {
	Event *WalletStage1SimulatorImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorImplementationUpdated)
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
		it.Event = new(WalletStage1SimulatorImplementationUpdated)
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
func (it *WalletStage1SimulatorImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorImplementationUpdated represents a ImplementationUpdated event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletStage1SimulatorImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorImplementationUpdatedIterator{contract: _WalletStage1Simulator.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorImplementationUpdated)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseImplementationUpdated(log types.Log) (*WalletStage1SimulatorImplementationUpdated, error) {
	event := new(WalletStage1SimulatorImplementationUpdated)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorNonceChangeIterator struct {
	Event *WalletStage1SimulatorNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorNonceChange)
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
		it.Event = new(WalletStage1SimulatorNonceChange)
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
func (it *WalletStage1SimulatorNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorNonceChange represents a NonceChange event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletStage1SimulatorNonceChangeIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorNonceChangeIterator{contract: _WalletStage1Simulator.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorNonceChange)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseNonceChange(log types.Log) (*WalletStage1SimulatorNonceChange, error) {
	event := new(WalletStage1SimulatorNonceChange)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage1SimulatorStaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorStaticSignatureSetIterator struct {
	Event *WalletStage1SimulatorStaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletStage1SimulatorStaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage1SimulatorStaticSignatureSet)
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
		it.Event = new(WalletStage1SimulatorStaticSignatureSet)
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
func (it *WalletStage1SimulatorStaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage1SimulatorStaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage1SimulatorStaticSignatureSet represents a StaticSignatureSet event raised by the WalletStage1Simulator contract.
type WalletStage1SimulatorStaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletStage1SimulatorStaticSignatureSetIterator, error) {

	logs, sub, err := _WalletStage1Simulator.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletStage1SimulatorStaticSignatureSetIterator{contract: _WalletStage1Simulator.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletStage1SimulatorStaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletStage1Simulator.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage1SimulatorStaticSignatureSet)
				if err := _WalletStage1Simulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletStage1Simulator *WalletStage1SimulatorFilterer) ParseStaticSignatureSet(log types.Log) (*WalletStage1SimulatorStaticSignatureSet, error) {
	event := new(WalletStage1SimulatorStaticSignatureSet)
	if err := _WalletStage1Simulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
