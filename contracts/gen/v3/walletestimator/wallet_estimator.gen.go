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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"estimate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061429d8061001f6000396000f3fe6080604052600436106101485760003560e01c80636ea44577116100c0578063ad55366b11610074578063bc197c8111610059578063bc197c8114610470578063f23a6e61146104b8578063f727ef1c146104fe5761014f565b8063ad55366b1461040a578063b93ea7ad1461045d5761014f565b80638c3f5563116100a55780638c3f5563146103c2578063975befdb146103e2578063aaf10f42146103f55761014f565b80636ea445771461038e5780638943ec02146103a15761014f565b80631a9b23371161011757806329561426116100fc57806329561426146103465780634fcf3eca1461036657806351605d80146103795761014f565b80631a9b2337146102ee5780631f6a1eb9146103335761014f565b8063025b22bc1461021257806313792a4a14610225578063150b7a02146102585780631626ba7e146102ce5761014f565b3661014f57005b6004361061021057600061016b610166368361332a565b61051e565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b4929190613390565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b6102106102203660046133c9565b610572565b34801561023157600080fd5b5061024561024036600461376d565b6105be565b6040519081526020015b60405180910390f35b34801561026457600080fd5b5061029d6102733660046138b4565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161024f565b3480156102da57600080fd5b5061029d6102e9366004613923565b610729565b3480156102fa57600080fd5b5061030e610309366004613984565b6107bf565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161024f565b6102106103413660046139a1565b6107ca565b34801561035257600080fd5b50610210610361366004613a12565b610853565b610210610374366004613984565b610897565b34801561038557600080fd5b50610245610959565b61021061039c366004613a2b565b610988565b3480156103ad57600080fd5b506102106103bc366004613a6d565b50505050565b3480156103ce57600080fd5b506102456103dd366004613a12565b6109f5565b6102456103f03660046139a1565b610a21565b34801561040157600080fd5b5061030e610abf565b34801561041657600080fd5b5061042a61042536600461376d565b610ac9565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c00161024f565b61021061046b366004613aaf565b610b03565b34801561047c57600080fd5b5061029d61048b366004613b29565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104c457600080fd5b5061029d6104d3366004613bf0565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561050a57600080fd5b50610210610519366004613c68565b610bc8565b600061056c7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c83565b92915050565b3330146105b2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105bb81610ce1565b50565b6000808461010001515160016105d49190613ce8565b67ffffffffffffffff8111156105ec576105ec6133e4565b604051908082528060200260200182016040528015610615578160200160208202803683370190505b50905060005b8561010001515181101561068757856101000151818151811061064057610640613cfb565b602002602001015182828151811061065a5761065a613cfb565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015260010161061b565b50338186610100015151815181106106a1576106a1613cfb565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006106db868686610d36565b5090508061071b578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610780828686610d36565b509050806107945750600091506107229050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b600061056c8261051e565b60005a905060006107db8686610ec7565b90506107ef816060015182608001516112e5565b6000806107fd838787610d36565b915091508161083e578286866040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b6108498482856113cd565b5050505050505050565b33301461088e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6105bb8161172d565b3330146108d2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60006108dd8261051e565b73ffffffffffffffffffffffffffffffffffffffff160361094e576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105a9565b6105bb8160006117bd565b60006109837fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b3330146109c3576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60005a905060006109d48484610ec7565b905060006109e18261187d565b90506109ee8382846113cd565b5050505050565b600061056c7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c83565b6000805a90506000610a338787610ec7565b6060810151909150610a4d90610a48816109f5565b6112e5565b600080610a5b838888610d36565b9150915081610a9c578287876040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b610aa78482856118f8565b5a610ab29085613fdd565b9998505050505050505050565b6000610983305490565b600080600080600080610ae0898989600080611b41565b939950919750945092509050610af583611e66565b935093975093979195509350565b333014610b3e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6000610b498361051e565b73ffffffffffffffffffffffffffffffffffffffff1614610bba576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105a9565b610bc482826117bd565b5050565b333014610c03576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b610c1c8383836bffffffffffffffffffffffff16611e7a565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610ca2929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610ce9813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610d4e57610d4e613cfb565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e4d57610dab8661187d565b9150600080610db984611f09565b9150915042811015610e01576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105a9565b73ffffffffffffffffffffffffffffffffffffffff82161580610e39575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e4a5760019450505050610ebf565b50505b6000806000610e60898989600080611b41565b985092955090935091505082821015610eaf576040517ffd41fcba00000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016105a9565b610eb881611e66565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610f2b5760006060840152610f3c565b84810135606090811c908401526014015b6007600183901c168015610f8f5760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610fa457506001610fcd565b83602016602003610fc05750600282019186013560f01c610fcd565b50600182019186013560f81c5b8067ffffffffffffffff811115610fe657610fe66133e4565b60405190808252806020026020018201604052801561107157816020015b61105e6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b8152602001906001900390816110045790505b50604086015260005b818110156112da5760018085019489013560f81c9080821690036110d95730876040015183815181106110af576110af613cfb565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611123565b8489013560601c60148601886040015184815181106110fa576110fa613cfb565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036111615784890135602086018860400151848151811061114c5761114c613cfb565b60200260200101516020018197508281525050505b600480821690036111f957600385019489013560e81c89868a6111848483613ce8565b9261119193929190613ff0565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106111dc576111dc613cfb565b6020908102919091010151604001526111f58187613ce8565b9550505b600880821690036112375784890135602086018860400151848151811061122257611222613cfb565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061125757611257613cfb565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061128d5761128d613cfb565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106112c2576112c2613cfb565b602090810291909101015160c001525060010161107a565b505050505092915050565b60006112f0836109f5565b905081811461133c576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105a9565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b81811015611725576000846040015182815181106113f8576113f8613cfb565b602002602001015190508315611411576000935061145a565b8060a001511561145a5760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a15061171d565b6060810151801580159061146d5750805a105b156114aa5785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a99392919061401a565b600082608001511561157e5782516115779083156114c857836114ca565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016114f59695949392919061403f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611f55565b90506115a5565b825160208401516115a2919084156115965784611598565b5a5b8660400151611f6b565b90505b806116e05760c083015161160157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856115e2611f83565b6040516115f19392919061407c565b60405180910390a150505061171d565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161166b578684611636611f83565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105a99392919061409b565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016116e0577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856116c1611f83565b6040516116d09392919061407c565b60405180910390a1505050611725565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b6001016113d8565b505050505050565b80611764576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61178d7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610d2b565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b60008061188e836020015130611fa2565b9050600061189b8461206f565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604081015151600090815b818110156117255760008460400151828151811061192357611923613cfb565b60200260200101519050831561193c5760009350611985565b8060a00151156119855760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611b39565b606081015180158015906119985750805a105b156119d55785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a99392919061401a565b60008260800151156119fa5782516119f39083156114c857836114ca565b9050611a15565b82516020840151611a12919084156115965784611598565b90505b80611afc5760c0830151611a7157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d8885611a52611f83565b604051611a619392919061407c565b60405180910390a1505050611b39565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611aa6578684611636611f83565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611afc577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856116c1611f83565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b600101611903565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611b8c575073ffffffffffffffffffffffffffffffffffffffff8916155b15611ca8578b82013560601c985060149091019089611ca85760038201918c013560e81c60008d848e611bbf8583613ce8565b92611bcc93929190613ff0565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611c5790309085906004016140c6565b6040805180830381865afa158015611c73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c9791906140fd565b9250611ca38285613ce8565b935050505b82600116600103611ce257611cd08d8a838f8f87908092611ccb93929190613ff0565b6122c5565b97509750975097509750505050611e59565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611d4d6001600586901c811690613ce8565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611d988d61187d565b9350611db68d858e8e86908092611db193929190613ff0565b612511565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611e01575080518614155b8015611e11575080602001518511155b15611e55576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105a9565b5050505b9550955095509550959050565b6000611e7182612ef5565b50600192915050565b611f047fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611f377fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c83565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856120125746612015565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080612080836101000151612f00565b835190915060ff166120f357600061209b8460400151612f83565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016118d9565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016121835760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016121f35760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612165565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016122635760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201529081019190915260608101829052608001612165565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105a9565b6000806000806000612327604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b8882101561249e5760038201916000908b013560e81c61236f8482613ce8565b9150600090508a8214612383576000612385565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146123b657856123b8565b8f5b90506123d8818e8e889087926123d093929190613ff0565b600186611b41565b50929d50909b50995097508a8a1015612435578c8c869085926123fd93929190613ff0565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105a9949392919061414e565b829450888e60000151036124485760008e525b83881061248b576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105a9565b505060c08401879052915084905061234f565b8a51158015906124b257508a602001518511155b156124f6576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105a9565b6124ff8d61187d565b93505050509550955095509550959050565b60008060005b83811015612eeb57600181019085013560f881901c9060fc1c8061263457600f8216600081900361254f5750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156125e4573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000612607828960ff16613005565b90508b6126145780612623565b60008c81526020829052604090205b9b5050505050505050505050612517565b6001810361269857600f821660008190036126565750600183019287013560f81c5b601484019388013560601c60006126708260ff8516613005565b90508661267d578061268c565b60008781526020829052604090205b96505050505050612517565b6002810361287f576003821660008190036126ba5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261277893929190613ff0565b6040518463ffffffff1660e01b815260040161279693929190614175565b602060405180830381865afa1580156127b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127d7919061418f565b7fffffffff00000000000000000000000000000000000000000000000000000000161461283a578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105a99594939291906141ac565b8097508460ff168a0199506000612854858760ff16613005565b9050896128615780612870565b60008a81526020829052604090205b99505050505050505050612517565b600381036128b35760208301928701358461289a57806128a9565b60008581526020829052604090205b9450505050612517565b6004810361295957600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806129288e8e8e8e8c908892611db193929190613ff0565b91509150829750818a019950612948898260009182526020526040902090565b985082975050505050505050612517565b60068103612a23576003600283901c16600081900361297f5750600183019287013560f81c5b6003831660008190036129995750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806129d78f8f8f8f8d908892611db193929190613ff0565b915091508298508482106129ea57998501995b60006129f782878961306c565b90508a612a045780612a13565b60008b81526020829052604090205b9a50505050505050505050612517565b60058103612a90576020830192870135888103612a5e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612a69826130d0565b905085612a765780612a85565b60008681526020829052604090205b955050505050612517565b60078103612b9657600f82166000819003612ab25750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016125c2565b60088103612bea5760208301928701356000612bb28b82613124565b9050808203612bdf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006126708261319f565b60098103612d5057600382166000819003612c0c5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612ca693929190613ff0565b6040518463ffffffff1660e01b8152600401612cc493929190613fad565b602060405180830381865afa158015612ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d0591906141f4565b90508197508460ff168a0199506000612d22858760ff16846131da565b905089612d2f5780612d3e565b60008a81526020829052604090205b99508298505050505050505050612517565b600a8103612eb657600382166000819003612d725750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612e0b93929190613ff0565b6040518463ffffffff1660e01b8152600401612e2993929190614175565b602060405180830381865afa158015612e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6a91906141f4565b90508198508560ff168b019a506000612e87868860ff16846131da565b90508a612e945780612ea3565b60008b81526020829052604090205b9a50829950505050505050505050612517565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105a9565b5094509492505050565b600061056c82613248565b6000606060005b8351811015612f745781848281518110612f2357612f23613cfb565b6020026020010151604051602001612f3c92919061420d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612f07565b50805160209091012092915050565b6000606060005b8351811015612f74576000612fb7858381518110612faa57612faa613cfb565b602002602001015161327b565b90508281604051602001612fcc929190614245565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612f8a565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612051565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080613135846020015184611fa2565b905060006131428561206f565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001613107565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d016130b1565b6000811580159061056c5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161310798979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015613389577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146133c457600080fd5b919050565b6000602082840312156133db57600080fd5b610722826133a0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613436576134366133e4565b60405290565b604051610120810167ffffffffffffffff81118282101715613436576134366133e4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156134a7576134a76133e4565b604052919050565b803560ff811681146133c457600080fd5b803580151581146133c457600080fd5b600067ffffffffffffffff8211156134ea576134ea6133e4565b5060051b60200190565b600082601f83011261350557600080fd5b813567ffffffffffffffff81111561351f5761351f6133e4565b61355060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613460565b81815284602083860101111561356557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261359357600080fd5b81356135a66135a1826134d0565b613460565b8082825260208201915060208360051b8601019250858311156135c857600080fd5b602085015b838110156136b557803567ffffffffffffffff8111156135ec57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561362057600080fd5b613628613413565b613634602083016133a0565b815260408201356020820152606082013567ffffffffffffffff81111561365a57600080fd5b6136698a6020838601016134f4565b6040830152506080820135606082015261368560a083016134c0565b608082015261369660c083016134c0565b60a082015260e0919091013560c08201528352602092830192016135cd565b5095945050505050565b600082601f8301126136d057600080fd5b81356136de6135a1826134d0565b8082825260208201915060208360051b86010192508583111561370057600080fd5b602085015b838110156136b557613716816133a0565b835260209283019201613705565b60008083601f84011261373657600080fd5b50813567ffffffffffffffff81111561374e57600080fd5b60208301915083602082850101111561376657600080fd5b9250929050565b60008060006040848603121561378257600080fd5b833567ffffffffffffffff81111561379957600080fd5b840161012081870312156137ac57600080fd5b6137b461343c565b6137bd826134af565b81526137cb602083016134c0565b6020820152604082013567ffffffffffffffff8111156137ea57600080fd5b6137f688828501613582565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561382a57600080fd5b613836888285016134f4565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561386b57600080fd5b613877888285016136bf565b61010083015250935050602084013567ffffffffffffffff81111561389b57600080fd5b6138a786828701613724565b9497909650939450505050565b6000806000806000608086880312156138cc57600080fd5b6138d5866133a0565b94506138e3602087016133a0565b935060408601359250606086013567ffffffffffffffff81111561390657600080fd5b61391288828901613724565b969995985093965092949392505050565b60008060006040848603121561393857600080fd5b83359250602084013567ffffffffffffffff81111561389b57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105bb57600080fd5b60006020828403121561399657600080fd5b813561072281613956565b600080600080604085870312156139b757600080fd5b843567ffffffffffffffff8111156139ce57600080fd5b6139da87828801613724565b909550935050602085013567ffffffffffffffff8111156139fa57600080fd5b613a0687828801613724565b95989497509550505050565b600060208284031215613a2457600080fd5b5035919050565b60008060208385031215613a3e57600080fd5b823567ffffffffffffffff811115613a5557600080fd5b613a6185828601613724565b90969095509350505050565b60008060008060608587031215613a8357600080fd5b613a8c856133a0565b935060208501359250604085013567ffffffffffffffff8111156139fa57600080fd5b60008060408385031215613ac257600080fd5b8235613acd81613956565b9150613adb602084016133a0565b90509250929050565b60008083601f840112613af657600080fd5b50813567ffffffffffffffff811115613b0e57600080fd5b6020830191508360208260051b850101111561376657600080fd5b60008060008060008060008060a0898b031215613b4557600080fd5b613b4e896133a0565b9750613b5c60208a016133a0565b9650604089013567ffffffffffffffff811115613b7857600080fd5b613b848b828c01613ae4565b909750955050606089013567ffffffffffffffff811115613ba457600080fd5b613bb08b828c01613ae4565b909550935050608089013567ffffffffffffffff811115613bd057600080fd5b613bdc8b828c01613724565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613c0957600080fd5b613c12876133a0565b9550613c20602088016133a0565b94506040870135935060608701359250608087013567ffffffffffffffff811115613c4a57600080fd5b613c5689828a01613724565b979a9699509497509295939492505050565b600080600060608486031215613c7d57600080fd5b83359250613c8d602085016133a0565b915060408401356bffffffffffffffffffffffff81168114613cae57600080fd5b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561056c5761056c613cb9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613d45578181015183820152602001613d2d565b50506000910152565b60008151808452613d66816020860160208601613d2a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613e68577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613e2460e0860182613d4e565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613db6565b50909695505050505050565b600081518084526020840193506020830160005b82811015613ebc57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613e88565b5093949350505050565b805160ff16825260006020820151613ee2602085018215159052565b5060408201516101206040850152613efe610120850182613d98565b9050606083015160608501526080830151608085015260a083015184820360a0860152613f2b8282613d4e565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613f5b8282613e74565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613fc06040830186613ec6565b8281036020840152613fd3818587613f64565b9695505050505050565b8181038181111561056c5761056c613cb9565b6000808585111561400057600080fd5b8386111561400d57600080fd5b5050820193919092039150565b60608152600061402d6060830186613ec6565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a0820152600061407060c0830184613d4e565b98975050505050505050565b838152826020820152606060408201526000613f5b6060830184613d4e565b6060815260006140ae6060830186613ec6565b8460208401528281036040840152613fd38185613d4e565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006140f56040830184613d4e565b949350505050565b6000604082840312801561411057600080fd5b506040805190810167ffffffffffffffff81118282101715614134576141346133e4565b604052825181526020928301519281019290925250919050565b606081526000614162606083018688613f64565b6020830194909452506040015292915050565b838152604060208201526000613f5b604083018486613f64565b6000602082840312156141a157600080fd5b815161072281613956565b6080815260006141bf6080830188613ec6565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152614070818587613f64565b60006020828403121561420657600080fd5b5051919050565b6040815260006142206040830185613d4e565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351614257818460208801613d2a565b919091019182525060200191905056fea26469706673582212208bdc3ae308f296ddf5c8a0d9120f3c8ef76f3cfa7a579471a5d5a515becbefdd64736f6c634300081b0033",
}

// WalletEstimatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletEstimatorMetaData.ABI instead.
var WalletEstimatorABI = WalletEstimatorMetaData.ABI

// WalletEstimatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletEstimatorMetaData.Bin instead.
var WalletEstimatorBin = WalletEstimatorMetaData.Bin

// DeployWalletEstimator deploys a new Ethereum contract, binding an instance of WalletEstimator to it.
func DeployWalletEstimator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletEstimator, error) {
	parsed, err := WalletEstimatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletEstimatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletEstimator{WalletEstimatorCaller: WalletEstimatorCaller{contract: contract}, WalletEstimatorTransactor: WalletEstimatorTransactor{contract: contract}, WalletEstimatorFilterer: WalletEstimatorFilterer{contract: contract}}, nil
}

// WalletEstimatorDeployedBin is the resulting bytecode of the created contract
const WalletEstimatorDeployedBin = "0x6080604052600436106101485760003560e01c80636ea44577116100c0578063ad55366b11610074578063bc197c8111610059578063bc197c8114610470578063f23a6e61146104b8578063f727ef1c146104fe5761014f565b8063ad55366b1461040a578063b93ea7ad1461045d5761014f565b80638c3f5563116100a55780638c3f5563146103c2578063975befdb146103e2578063aaf10f42146103f55761014f565b80636ea445771461038e5780638943ec02146103a15761014f565b80631a9b23371161011757806329561426116100fc57806329561426146103465780634fcf3eca1461036657806351605d80146103795761014f565b80631a9b2337146102ee5780631f6a1eb9146103335761014f565b8063025b22bc1461021257806313792a4a14610225578063150b7a02146102585780631626ba7e146102ce5761014f565b3661014f57005b6004361061021057600061016b610166368361332a565b61051e565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b4929190613390565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b6102106102203660046133c9565b610572565b34801561023157600080fd5b5061024561024036600461376d565b6105be565b6040519081526020015b60405180910390f35b34801561026457600080fd5b5061029d6102733660046138b4565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161024f565b3480156102da57600080fd5b5061029d6102e9366004613923565b610729565b3480156102fa57600080fd5b5061030e610309366004613984565b6107bf565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161024f565b6102106103413660046139a1565b6107ca565b34801561035257600080fd5b50610210610361366004613a12565b610853565b610210610374366004613984565b610897565b34801561038557600080fd5b50610245610959565b61021061039c366004613a2b565b610988565b3480156103ad57600080fd5b506102106103bc366004613a6d565b50505050565b3480156103ce57600080fd5b506102456103dd366004613a12565b6109f5565b6102456103f03660046139a1565b610a21565b34801561040157600080fd5b5061030e610abf565b34801561041657600080fd5b5061042a61042536600461376d565b610ac9565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c00161024f565b61021061046b366004613aaf565b610b03565b34801561047c57600080fd5b5061029d61048b366004613b29565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104c457600080fd5b5061029d6104d3366004613bf0565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561050a57600080fd5b50610210610519366004613c68565b610bc8565b600061056c7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c83565b92915050565b3330146105b2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105bb81610ce1565b50565b6000808461010001515160016105d49190613ce8565b67ffffffffffffffff8111156105ec576105ec6133e4565b604051908082528060200260200182016040528015610615578160200160208202803683370190505b50905060005b8561010001515181101561068757856101000151818151811061064057610640613cfb565b602002602001015182828151811061065a5761065a613cfb565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015260010161061b565b50338186610100015151815181106106a1576106a1613cfb565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006106db868686610d36565b5090508061071b578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610780828686610d36565b509050806107945750600091506107229050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b600061056c8261051e565b60005a905060006107db8686610ec7565b90506107ef816060015182608001516112e5565b6000806107fd838787610d36565b915091508161083e578286866040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b6108498482856113cd565b5050505050505050565b33301461088e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6105bb8161172d565b3330146108d2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60006108dd8261051e565b73ffffffffffffffffffffffffffffffffffffffff160361094e576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105a9565b6105bb8160006117bd565b60006109837fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b3330146109c3576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60005a905060006109d48484610ec7565b905060006109e18261187d565b90506109ee8382846113cd565b5050505050565b600061056c7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c83565b6000805a90506000610a338787610ec7565b6060810151909150610a4d90610a48816109f5565b6112e5565b600080610a5b838888610d36565b9150915081610a9c578287876040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613fad565b610aa78482856118f8565b5a610ab29085613fdd565b9998505050505050505050565b6000610983305490565b600080600080600080610ae0898989600080611b41565b939950919750945092509050610af583611e66565b935093975093979195509350565b333014610b3e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6000610b498361051e565b73ffffffffffffffffffffffffffffffffffffffff1614610bba576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105a9565b610bc482826117bd565b5050565b333014610c03576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b610c1c8383836bffffffffffffffffffffffff16611e7a565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610ca2929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610ce9813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610d4e57610d4e613cfb565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e4d57610dab8661187d565b9150600080610db984611f09565b9150915042811015610e01576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105a9565b73ffffffffffffffffffffffffffffffffffffffff82161580610e39575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e4a5760019450505050610ebf565b50505b6000806000610e60898989600080611b41565b985092955090935091505082821015610eaf576040517ffd41fcba00000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016105a9565b610eb881611e66565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610f2b5760006060840152610f3c565b84810135606090811c908401526014015b6007600183901c168015610f8f5760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610fa457506001610fcd565b83602016602003610fc05750600282019186013560f01c610fcd565b50600182019186013560f81c5b8067ffffffffffffffff811115610fe657610fe66133e4565b60405190808252806020026020018201604052801561107157816020015b61105e6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b8152602001906001900390816110045790505b50604086015260005b818110156112da5760018085019489013560f81c9080821690036110d95730876040015183815181106110af576110af613cfb565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611123565b8489013560601c60148601886040015184815181106110fa576110fa613cfb565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036111615784890135602086018860400151848151811061114c5761114c613cfb565b60200260200101516020018197508281525050505b600480821690036111f957600385019489013560e81c89868a6111848483613ce8565b9261119193929190613ff0565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106111dc576111dc613cfb565b6020908102919091010151604001526111f58187613ce8565b9550505b600880821690036112375784890135602086018860400151848151811061122257611222613cfb565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061125757611257613cfb565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061128d5761128d613cfb565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106112c2576112c2613cfb565b602090810291909101015160c001525060010161107a565b505050505092915050565b60006112f0836109f5565b905081811461133c576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105a9565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b81811015611725576000846040015182815181106113f8576113f8613cfb565b602002602001015190508315611411576000935061145a565b8060a001511561145a5760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a15061171d565b6060810151801580159061146d5750805a105b156114aa5785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a99392919061401a565b600082608001511561157e5782516115779083156114c857836114ca565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016114f59695949392919061403f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611f55565b90506115a5565b825160208401516115a2919084156115965784611598565b5a5b8660400151611f6b565b90505b806116e05760c083015161160157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856115e2611f83565b6040516115f19392919061407c565b60405180910390a150505061171d565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161166b578684611636611f83565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105a99392919061409b565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016116e0577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856116c1611f83565b6040516116d09392919061407c565b60405180910390a1505050611725565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b6001016113d8565b505050505050565b80611764576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61178d7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610d2b565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b60008061188e836020015130611fa2565b9050600061189b8461206f565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604081015151600090815b818110156117255760008460400151828151811061192357611923613cfb565b60200260200101519050831561193c5760009350611985565b8060a00151156119855760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611b39565b606081015180158015906119985750805a105b156119d55785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a99392919061401a565b60008260800151156119fa5782516119f39083156114c857836114ca565b9050611a15565b82516020840151611a12919084156115965784611598565b90505b80611afc5760c0830151611a7157600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d8885611a52611f83565b604051611a619392919061407c565b60405180910390a1505050611b39565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611aa6578684611636611f83565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611afc577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856116c1611f83565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b600101611903565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611b8c575073ffffffffffffffffffffffffffffffffffffffff8916155b15611ca8578b82013560601c985060149091019089611ca85760038201918c013560e81c60008d848e611bbf8583613ce8565b92611bcc93929190613ff0565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611c5790309085906004016140c6565b6040805180830381865afa158015611c73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c9791906140fd565b9250611ca38285613ce8565b935050505b82600116600103611ce257611cd08d8a838f8f87908092611ccb93929190613ff0565b6122c5565b97509750975097509750505050611e59565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611d4d6001600586901c811690613ce8565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611d988d61187d565b9350611db68d858e8e86908092611db193929190613ff0565b612511565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611e01575080518614155b8015611e11575080602001518511155b15611e55576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105a9565b5050505b9550955095509550959050565b6000611e7182612ef5565b50600192915050565b611f047fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611f377fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c83565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856120125746612015565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080612080836101000151612f00565b835190915060ff166120f357600061209b8460400151612f83565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016118d9565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016121835760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016121f35760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612165565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016122635760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201529081019190915260608101829052608001612165565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105a9565b6000806000806000612327604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b8882101561249e5760038201916000908b013560e81c61236f8482613ce8565b9150600090508a8214612383576000612385565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146123b657856123b8565b8f5b90506123d8818e8e889087926123d093929190613ff0565b600186611b41565b50929d50909b50995097508a8a1015612435578c8c869085926123fd93929190613ff0565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105a9949392919061414e565b829450888e60000151036124485760008e525b83881061248b576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105a9565b505060c08401879052915084905061234f565b8a51158015906124b257508a602001518511155b156124f6576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105a9565b6124ff8d61187d565b93505050509550955095509550959050565b60008060005b83811015612eeb57600181019085013560f881901c9060fc1c8061263457600f8216600081900361254f5750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156125e4573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000612607828960ff16613005565b90508b6126145780612623565b60008c81526020829052604090205b9b5050505050505050505050612517565b6001810361269857600f821660008190036126565750600183019287013560f81c5b601484019388013560601c60006126708260ff8516613005565b90508661267d578061268c565b60008781526020829052604090205b96505050505050612517565b6002810361287f576003821660008190036126ba5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261277893929190613ff0565b6040518463ffffffff1660e01b815260040161279693929190614175565b602060405180830381865afa1580156127b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127d7919061418f565b7fffffffff00000000000000000000000000000000000000000000000000000000161461283a578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105a99594939291906141ac565b8097508460ff168a0199506000612854858760ff16613005565b9050896128615780612870565b60008a81526020829052604090205b99505050505050505050612517565b600381036128b35760208301928701358461289a57806128a9565b60008581526020829052604090205b9450505050612517565b6004810361295957600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806129288e8e8e8e8c908892611db193929190613ff0565b91509150829750818a019950612948898260009182526020526040902090565b985082975050505050505050612517565b60068103612a23576003600283901c16600081900361297f5750600183019287013560f81c5b6003831660008190036129995750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806129d78f8f8f8f8d908892611db193929190613ff0565b915091508298508482106129ea57998501995b60006129f782878961306c565b90508a612a045780612a13565b60008b81526020829052604090205b9a50505050505050505050612517565b60058103612a90576020830192870135888103612a5e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612a69826130d0565b905085612a765780612a85565b60008681526020829052604090205b955050505050612517565b60078103612b9657600f82166000819003612ab25750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016125c2565b60088103612bea5760208301928701356000612bb28b82613124565b9050808203612bdf577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006126708261319f565b60098103612d5057600382166000819003612c0c5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612ca693929190613ff0565b6040518463ffffffff1660e01b8152600401612cc493929190613fad565b602060405180830381865afa158015612ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d0591906141f4565b90508197508460ff168a0199506000612d22858760ff16846131da565b905089612d2f5780612d3e565b60008a81526020829052604090205b99508298505050505050505050612517565b600a8103612eb657600382166000819003612d725750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612e0b93929190613ff0565b6040518463ffffffff1660e01b8152600401612e2993929190614175565b602060405180830381865afa158015612e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6a91906141f4565b90508198508560ff168b019a506000612e87868860ff16846131da565b90508a612e945780612ea3565b60008b81526020829052604090205b9a50829950505050505050505050612517565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105a9565b5094509492505050565b600061056c82613248565b6000606060005b8351811015612f745781848281518110612f2357612f23613cfb565b6020026020010151604051602001612f3c92919061420d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612f07565b50805160209091012092915050565b6000606060005b8351811015612f74576000612fb7858381518110612faa57612faa613cfb565b602002602001015161327b565b90508281604051602001612fcc929190614245565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612f8a565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501612051565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080613135846020015184611fa2565b905060006131428561206f565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001613107565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d016130b1565b6000811580159061056c5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161310798979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015613389577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146133c457600080fd5b919050565b6000602082840312156133db57600080fd5b610722826133a0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613436576134366133e4565b60405290565b604051610120810167ffffffffffffffff81118282101715613436576134366133e4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156134a7576134a76133e4565b604052919050565b803560ff811681146133c457600080fd5b803580151581146133c457600080fd5b600067ffffffffffffffff8211156134ea576134ea6133e4565b5060051b60200190565b600082601f83011261350557600080fd5b813567ffffffffffffffff81111561351f5761351f6133e4565b61355060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613460565b81815284602083860101111561356557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261359357600080fd5b81356135a66135a1826134d0565b613460565b8082825260208201915060208360051b8601019250858311156135c857600080fd5b602085015b838110156136b557803567ffffffffffffffff8111156135ec57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561362057600080fd5b613628613413565b613634602083016133a0565b815260408201356020820152606082013567ffffffffffffffff81111561365a57600080fd5b6136698a6020838601016134f4565b6040830152506080820135606082015261368560a083016134c0565b608082015261369660c083016134c0565b60a082015260e0919091013560c08201528352602092830192016135cd565b5095945050505050565b600082601f8301126136d057600080fd5b81356136de6135a1826134d0565b8082825260208201915060208360051b86010192508583111561370057600080fd5b602085015b838110156136b557613716816133a0565b835260209283019201613705565b60008083601f84011261373657600080fd5b50813567ffffffffffffffff81111561374e57600080fd5b60208301915083602082850101111561376657600080fd5b9250929050565b60008060006040848603121561378257600080fd5b833567ffffffffffffffff81111561379957600080fd5b840161012081870312156137ac57600080fd5b6137b461343c565b6137bd826134af565b81526137cb602083016134c0565b6020820152604082013567ffffffffffffffff8111156137ea57600080fd5b6137f688828501613582565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561382a57600080fd5b613836888285016134f4565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561386b57600080fd5b613877888285016136bf565b61010083015250935050602084013567ffffffffffffffff81111561389b57600080fd5b6138a786828701613724565b9497909650939450505050565b6000806000806000608086880312156138cc57600080fd5b6138d5866133a0565b94506138e3602087016133a0565b935060408601359250606086013567ffffffffffffffff81111561390657600080fd5b61391288828901613724565b969995985093965092949392505050565b60008060006040848603121561393857600080fd5b83359250602084013567ffffffffffffffff81111561389b57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105bb57600080fd5b60006020828403121561399657600080fd5b813561072281613956565b600080600080604085870312156139b757600080fd5b843567ffffffffffffffff8111156139ce57600080fd5b6139da87828801613724565b909550935050602085013567ffffffffffffffff8111156139fa57600080fd5b613a0687828801613724565b95989497509550505050565b600060208284031215613a2457600080fd5b5035919050565b60008060208385031215613a3e57600080fd5b823567ffffffffffffffff811115613a5557600080fd5b613a6185828601613724565b90969095509350505050565b60008060008060608587031215613a8357600080fd5b613a8c856133a0565b935060208501359250604085013567ffffffffffffffff8111156139fa57600080fd5b60008060408385031215613ac257600080fd5b8235613acd81613956565b9150613adb602084016133a0565b90509250929050565b60008083601f840112613af657600080fd5b50813567ffffffffffffffff811115613b0e57600080fd5b6020830191508360208260051b850101111561376657600080fd5b60008060008060008060008060a0898b031215613b4557600080fd5b613b4e896133a0565b9750613b5c60208a016133a0565b9650604089013567ffffffffffffffff811115613b7857600080fd5b613b848b828c01613ae4565b909750955050606089013567ffffffffffffffff811115613ba457600080fd5b613bb08b828c01613ae4565b909550935050608089013567ffffffffffffffff811115613bd057600080fd5b613bdc8b828c01613724565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613c0957600080fd5b613c12876133a0565b9550613c20602088016133a0565b94506040870135935060608701359250608087013567ffffffffffffffff811115613c4a57600080fd5b613c5689828a01613724565b979a9699509497509295939492505050565b600080600060608486031215613c7d57600080fd5b83359250613c8d602085016133a0565b915060408401356bffffffffffffffffffffffff81168114613cae57600080fd5b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561056c5761056c613cb9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613d45578181015183820152602001613d2d565b50506000910152565b60008151808452613d66816020860160208601613d2a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613e68577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613e2460e0860182613d4e565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613db6565b50909695505050505050565b600081518084526020840193506020830160005b82811015613ebc57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613e88565b5093949350505050565b805160ff16825260006020820151613ee2602085018215159052565b5060408201516101206040850152613efe610120850182613d98565b9050606083015160608501526080830151608085015260a083015184820360a0860152613f2b8282613d4e565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613f5b8282613e74565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613fc06040830186613ec6565b8281036020840152613fd3818587613f64565b9695505050505050565b8181038181111561056c5761056c613cb9565b6000808585111561400057600080fd5b8386111561400d57600080fd5b5050820193919092039150565b60608152600061402d6060830186613ec6565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a0820152600061407060c0830184613d4e565b98975050505050505050565b838152826020820152606060408201526000613f5b6060830184613d4e565b6060815260006140ae6060830186613ec6565b8460208401528281036040840152613fd38185613d4e565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006140f56040830184613d4e565b949350505050565b6000604082840312801561411057600080fd5b506040805190810167ffffffffffffffff81118282101715614134576141346133e4565b604052825181526020928301519281019290925250919050565b606081526000614162606083018688613f64565b6020830194909452506040015292915050565b838152604060208201526000613f5b604083018486613f64565b6000602082840312156141a157600080fd5b815161072281613956565b6080815260006141bf6080830188613ec6565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152614070818587613f64565b60006020828403121561420657600080fd5b5051919050565b6040815260006142206040830185613d4e565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351614257818460208801613d2a565b919091019182525060200191905056fea26469706673582212208bdc3ae308f296ddf5c8a0d9120f3c8ef76f3cfa7a579471a5d5a515becbefdd64736f6c634300081b0033"

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
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletEstimator *WalletEstimatorCaller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletEstimator *WalletEstimatorSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletEstimator.Contract.ReadHook(&_WalletEstimator.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletEstimator *WalletEstimatorCallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletEstimator.Contract.ReadHook(&_WalletEstimator.CallOpts, signature)
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

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.AddHook(&_WalletEstimator.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletEstimator.Contract.AddHook(&_WalletEstimator.TransactOpts, signature, implementation)
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

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletEstimator *WalletEstimatorTransactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletEstimator *WalletEstimatorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.RemoveHook(&_WalletEstimator.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.RemoveHook(&_WalletEstimator.TransactOpts, signature)
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

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletEstimator *WalletEstimatorTransactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletEstimator.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletEstimator *WalletEstimatorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.TokenReceived(&_WalletEstimator.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletEstimator *WalletEstimatorTransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletEstimator.Contract.TokenReceived(&_WalletEstimator.TransactOpts, arg0, arg1, arg2)
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
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletEstimator *WalletEstimatorFilterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletEstimatorDefinedHookIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorDefinedHookIterator{contract: _WalletEstimator.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
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
// Solidity: event DefinedHook(bytes4 signature, address implementation)
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
