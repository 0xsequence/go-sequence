// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletstage2simulator

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

// WalletStage2SimulatorMetaData contains all meta data concerning the WalletStage2Simulator contract.
var WalletStage2SimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSimulator.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"internalType\":\"structSimulator.Result[]\",\"name\":\"results\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506144c08061001f6000396000f3fe6080604052600436106101485760003560e01c80636ea44577116100c0578063b93ea7ad11610074578063ca70785011610059578063ca707850146104a5578063f23a6e61146104c5578063f727ef1c1461050b5761014f565b8063b93ea7ad1461044a578063bc197c811461045d5761014f565b80638c3f5563116100a55780638c3f5563146103c2578063aaf10f42146103e2578063ad55366b146103f75761014f565b80636ea445771461038e5780638943ec02146103a15761014f565b80631f6a1eb9116101175780634fcf3eca116100fc5780634fcf3eca1461033857806351605d801461034b5780636406c3141461036e5761014f565b80631f6a1eb91461030557806329561426146103185761014f565b8063025b22bc14610212578063150b7a02146102255780631626ba7e146102a05780631a9b2337146102c05761014f565b3661014f57005b6004361061021057600061016b6101663683613490565b61052b565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b49291906134f6565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b61021061022036600461352f565b61057f565b34801561023157600080fd5b5061026a610240366004613593565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102ac57600080fd5b5061026a6102bb366004613602565b6105cb565b3480156102cc57600080fd5b506102e06102db36600461367c565b610663565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610297565b610210610313366004613699565b61066e565b34801561032457600080fd5b5061021061033336600461370a565b6106f0565b61021061034636600461367c565b610734565b34801561035757600080fd5b506103606107f6565b604051908152602001610297565b61038161037c366004613699565b610825565b60405161029791906137c0565b61021061039c36600461389e565b610857565b3480156103ad57600080fd5b506102106103bc3660046138e0565b50505050565b3480156103ce57600080fd5b506103606103dd36600461370a565b6108b7565b3480156103ee57600080fd5b506102e06108e3565b34801561040357600080fd5b50610417610412366004613c62565b6108ed565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610297565b610210610458366004613d90565b61091c565b34801561046957600080fd5b5061026a610478366004613e0a565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104b157600080fd5b506103606104c0366004613c62565b6109e1565b3480156104d157600080fd5b5061026a6104e0366004613ed1565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561051757600080fd5b50610210610526366004613f49565b610b4a565b60006105797fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c05565b92915050565b3330146105bf576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105c881610c63565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610622828686610cb8565b5090508061063657506000915061065c9050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b60006105798261052b565b600061067a8585610df9565b905061068e81606001518260800151611217565b60008061069c838686610cb8565b91509150816106dd578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105b6939291906141af565b6106e78184611221565b50505050505050565b33301461072b576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b6105c8816114ca565b33301461076f576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b600061077a8261052b565b73ffffffffffffffffffffffffffffffffffffffff16036107eb576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105b6565b6105c881600061155a565b60006108207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b606060006108338686610df9565b905060006108408261161a565b905061084c8183611695565b979650505050505050565b333014610892576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b600061089e8383610df9565b905060006108ab8261161a565b90506103bc8183611221565b60006105797f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c05565b6000610820305490565b600080600080600080610904898989600080611c05565b939d929c5060019b5090995097509095509350505050565b333014610957576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b60006109628361052b565b73ffffffffffffffffffffffffffffffffffffffff16146109d3576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105b6565b6109dd828261155a565b5050565b6000808461010001515160016109f7919061420e565b67ffffffffffffffff811115610a0f57610a0f613922565b604051908082528060200260200182016040528015610a38578160200160208202803683370190505b50905060005b85610100015151811015610aaa578561010001518181518110610a6357610a63614221565b6020026020010151828281518110610a7d57610a7d614221565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610a3e565b5033818661010001515181518110610ac457610ac4614221565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610afe868686610cb8565b50905080610b3e578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105b6939291906141af565b50600195945050505050565b333014610b85576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b610b9e8383836bffffffffffffffffffffffff16611f2a565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c24929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c6b813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610cd057610cd0614221565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610dcf57610d2d8661161a565b9150600080610d3b84611fb9565b9150915042811015610d83576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105b6565b73ffffffffffffffffffffffffffffffffffffffff82161580610dbb575073ffffffffffffffffffffffffffffffffffffffff821633145b15610dcc5760019450505050610df1565b50505b6000806000610de2898989600080611c05565b60019a50985050505050505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610e5d5760006060840152610e6e565b84810135606090811c908401526014015b6007600183901c168015610ec15760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610ed657506001610eff565b83602016602003610ef25750600282019186013560f01c610eff565b50600182019186013560f81c5b8067ffffffffffffffff811115610f1857610f18613922565b604051908082528060200260200182016040528015610fa357816020015b610f906040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610f365790505b50604086015260005b8181101561120c5760018085019489013560f81c90808216900361100b573087604001518381518110610fe157610fe1614221565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611055565b8489013560601c601486018860400151848151811061102c5761102c614221565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036110935784890135602086018860400151848151811061107e5761107e614221565b60200260200101516020018197508281525050505b6004808216900361112b57600385019489013560e81c89868a6110b6848361420e565b926110c393929190614250565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061110e5761110e614221565b602090810291909101015160400152611127818761420e565b9550505b600880821690036111695784890135602086018860400151848151811061115457611154614221565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061118957611189614221565b602002602001015160800190151590811515815250508060201660ff16602014876040015183815181106111bf576111bf614221565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106111f4576111f4614221565b602090810291909101015160c0015250600101610fac565b505050505092915050565b6109dd8282612005565b604081015151600090815b818110156114c35760008460400151828151811061124c5761124c614221565b602002602001015190508060a001518015611265575083155b156112ad576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506114bb565b606081015180158015906112c05750805a105b156112fd5785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105b69392919061427a565b600082608001511561132e57825161132790831561131b578361131d565b5a5b85604001516120f9565b9050611355565b82516020840151611352919084156113465784611348565b5a5b866040015161210f565b90505b8061147e5760c08301516113a8576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a15050506114bb565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114125786846113dd612127565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105b69392919061429f565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161147e5760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a15050506114c3565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b60010161122c565b5050505050565b80611501576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61152a7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610cad565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b60008061162b836020015130612146565b9050600061163884612213565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040810151516060906000908067ffffffffffffffff8111156116ba576116ba613922565b60405190808252806020026020018201604052801561171057816020015b6116fd6040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816116d85790505b50925060005b81811015611bfc5760008560400151828151811061173657611736614221565b602002602001015190508060a00151801561174f575083155b15611797576040805188815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611bf4565b606081015180158015906117aa5750805a105b156118435760058684815181106117c3576117c3614221565b60200260200101516000019060058111156117e0576117e0613723565b908160058111156117f3576117f3613723565b9052505a60405160200161180991815260200190565b60405160208183030381529060405286848151811061182a5761182a614221565b6020026020010151602001819052505050505050610579565b60008260800151156118a85760005a84519091506118739084156118675784611869565b5a5b86604001516120f9565b91505a61188090826142ca565b88868151811061189257611892614221565b6020026020010151604001818152505050611902565b60005a845160208601519192506118d19185156118c557856118c7565b5a5b876040015161210f565b91505a6118de90826142ca565b8886815181106118f0576118f0614221565b60200260200101516040018181525050505b80611b475760c08301516119c557604080518a815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1600287858151811061196157611961614221565b602002602001015160000190600581111561197e5761197e613723565b9081600581111561199157611991613723565b90525061199c612127565b8785815181106119ae576119ae614221565b602002602001015160200181905250505050611bf4565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611a6b576003878581518110611a0457611a04614221565b6020026020010151600001906005811115611a2157611a21613723565b90816005811115611a3457611a34613723565b905250611a3f612127565b878581518110611a5157611a51614221565b602002602001015160200181905250505050505050610579565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611b4757604080518a8152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a16004878581518110611ae357611ae3614221565b6020026020010151600001906005811115611b0057611b00613723565b90816005811115611b1357611b13613723565b905250611b1e612127565b878581518110611b3057611b30614221565b602002602001015160200181905250505050611bfc565b604080518a8152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a16001878581518110611b9457611b94614221565b6020026020010151600001906005811115611bb157611bb1613723565b90816005811115611bc457611bc4613723565b905250611bcf612127565b878581518110611be157611be1614221565b6020026020010151602001819052505050505b600101611716565b50505092915050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611c50575073ffffffffffffffffffffffffffffffffffffffff8916155b15611d6c578b82013560601c985060149091019089611d6c5760038201918c013560e81c60008d848e611c83858361420e565b92611c9093929190614250565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611d1b90309085906004016142dd565b6040805180830381865afa158015611d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d5b9190614314565b9250611d67828561420e565b935050505b82600116600103611da657611d948d8a838f8f87908092611d8f93929190614250565b612469565b97509750975097509750505050611f1d565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611e116001600586901c81169061420e565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611e5c8d61161a565b9350611e7a8d858e8e86908092611e7593929190614250565b6126b5565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611ec5575080518614155b8015611ed5575080602001518511155b15611f19576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105b6565b5050505b9550955095509550959050565b611fb47fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611fe77fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c05565b606081901c956bffffffffffffffffffffffff909116945092505050565b6000612010836108b7565b9050818114158015612020575060005b15612068576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105b6565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121b657466121b9565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080612224836101000151613099565b835190915060ff1661229757600061223f846040015161311c565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611676565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016123275760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016123975760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612309565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016124075760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001612309565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105b6565b60008060008060006124cb604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156126425760038201916000908b013560e81c612513848261420e565b9150600090508a8214612527576000612529565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841461255a578561255c565b8f5b905061257c818e8e8890879261257493929190614250565b600186611c05565b50929d50909b50995097508a8a10156125d9578c8c869085926125a193929190614250565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105b69493929190614365565b829450888e60000151036125ec5760008e525b83881061262f576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105b6565b505060c0840187905291508490506124f3565b8a511580159061265657508a602001518511155b1561269a576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105b6565b6126a38d61161a565b93505050509550955095509550959050565b60008060005b8381101561308f57600181019085013560f881901c9060fc1c806127d857600f821660008190036126f35750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015612788573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006127ab828960ff1661319e565b90508b6127b857806127c7565b60008c81526020829052604090205b9b50505050505050505050506126bb565b6001810361283c57600f821660008190036127fa5750600183019287013560f81c5b601484019388013560601c60006128148260ff851661319e565b9050866128215780612830565b60008781526020829052604090205b965050505050506126bb565b60028103612a235760038216600081900361285e5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261291c93929190614250565b6040518463ffffffff1660e01b815260040161293a9392919061438c565b602060405180830381865afa158015612957573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061297b91906143a6565b7fffffffff0000000000000000000000000000000000000000000000000000000016146129de578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105b69594939291906143c3565b8097508460ff168a01995060006129f8858760ff1661319e565b905089612a055780612a14565b60008a81526020829052604090205b995050505050505050506126bb565b60038103612a5757602083019287013584612a3e5780612a4d565b60008581526020829052604090205b94505050506126bb565b60048103612afd57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c16858301809650819250505060008186019050600080612acc8e8e8e8e8c908892611e7593929190614250565b91509150829750818a019950612aec898260009182526020526040902090565b9850829750505050505050506126bb565b60068103612bc7576003600283901c166000819003612b235750600183019287013560f81c5b600383166000819003612b3d5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff169150809750819250505060008187019050600080612b7b8f8f8f8f8d908892611e7593929190614250565b91509150829850848210612b8e57998501995b6000612b9b828789613205565b90508a612ba85780612bb7565b60008b81526020829052604090205b9a505050505050505050506126bb565b60058103612c34576020830192870135888103612c02577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612c0d82613269565b905085612c1a5780612c29565b60008681526020829052604090205b9550505050506126bb565b60078103612d3a57600f82166000819003612c565750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612766565b60088103612d8e5760208301928701356000612d568b826132bd565b9050808203612d83577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061281482613338565b60098103612ef457600382166000819003612db05750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612e4a93929190614250565b6040518463ffffffff1660e01b8152600401612e68939291906141af565b602060405180830381865afa158015612e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ea99190614417565b90508197508460ff168a0199506000612ec6858760ff1684613373565b905089612ed35780612ee2565b60008a81526020829052604090205b995082985050505050505050506126bb565b600a810361305a57600382166000819003612f165750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612faf93929190614250565b6040518463ffffffff1660e01b8152600401612fcd9392919061438c565b602060405180830381865afa158015612fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300e9190614417565b90508198508560ff168b019a50600061302b868860ff1684613373565b90508a6130385780613047565b60008b81526020829052604090205b9a508299505050505050505050506126bb565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105b6565b5094509492505050565b6000606060005b835181101561310d57818482815181106130bc576130bc614221565b60200260200101516040516020016130d5929190614430565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905291506001016130a0565b50805160209091012092915050565b6000606060005b835181101561310d57600061315085838151811061314357613143614221565b60200260200101516133e1565b90508281604051602001613165929190614468565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101613123565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b166031820152604581018290526000906065016121f5565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806132ce846020015184612146565b905060006132db85612213565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a60208201529081018290526000906060016132a0565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d0161324a565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016132a098979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff0000000000000000000000000000000000000000000000000000000081169060048410156134ef577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461352a57600080fd5b919050565b60006020828403121561354157600080fd5b61065c82613506565b60008083601f84011261355c57600080fd5b50813567ffffffffffffffff81111561357457600080fd5b60208301915083602082850101111561358c57600080fd5b9250929050565b6000806000806000608086880312156135ab57600080fd5b6135b486613506565b94506135c260208701613506565b935060408601359250606086013567ffffffffffffffff8111156135e557600080fd5b6135f18882890161354a565b969995985093965092949392505050565b60008060006040848603121561361757600080fd5b83359250602084013567ffffffffffffffff81111561363557600080fd5b6136418682870161354a565b9497909650939450505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105c857600080fd5b60006020828403121561368e57600080fd5b813561065c8161364e565b600080600080604085870312156136af57600080fd5b843567ffffffffffffffff8111156136c657600080fd5b6136d28782880161354a565b909550935050602085013567ffffffffffffffff8111156136f257600080fd5b6136fe8782880161354a565b95989497509550505050565b60006020828403121561371c57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b8381101561376d578181015183820152602001613755565b50506000910152565b6000815180845261378e816020860160208601613752565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613892577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110613852577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261386f6060880182613776565b6040928301519790920196909652945060209384019391909101906001016137e8565b50929695505050505050565b600080602083850312156138b157600080fd5b823567ffffffffffffffff8111156138c857600080fd5b6138d48582860161354a565b90969095509350505050565b600080600080606085870312156138f657600080fd5b6138ff85613506565b935060208501359250604085013567ffffffffffffffff8111156136f257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561397457613974613922565b60405290565b604051610120810167ffffffffffffffff8111828210171561397457613974613922565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156139e5576139e5613922565b604052919050565b803560ff8116811461352a57600080fd5b8035801515811461352a57600080fd5b600067ffffffffffffffff821115613a2857613a28613922565b5060051b60200190565b600082601f830112613a4357600080fd5b813567ffffffffffffffff811115613a5d57613a5d613922565b613a8e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161399e565b818152846020838601011115613aa357600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112613ad157600080fd5b8135613ae4613adf82613a0e565b61399e565b8082825260208201915060208360051b860101925085831115613b0657600080fd5b602085015b83811015613bf357803567ffffffffffffffff811115613b2a57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613b5e57600080fd5b613b66613951565b613b7260208301613506565b815260408201356020820152606082013567ffffffffffffffff811115613b9857600080fd5b613ba78a602083860101613a32565b60408301525060808201356060820152613bc360a083016139fe565b6080820152613bd460c083016139fe565b60a082015260e0919091013560c0820152835260209283019201613b0b565b5095945050505050565b600082601f830112613c0e57600080fd5b8135613c1c613adf82613a0e565b8082825260208201915060208360051b860101925085831115613c3e57600080fd5b602085015b83811015613bf357613c5481613506565b835260209283019201613c43565b600080600060408486031215613c7757600080fd5b833567ffffffffffffffff811115613c8e57600080fd5b84016101208187031215613ca157600080fd5b613ca961397a565b613cb2826139ed565b8152613cc0602083016139fe565b6020820152604082013567ffffffffffffffff811115613cdf57600080fd5b613ceb88828501613ac0565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff811115613d1f57600080fd5b613d2b88828501613a32565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff811115613d6057600080fd5b613d6c88828501613bfd565b61010083015250935050602084013567ffffffffffffffff81111561363557600080fd5b60008060408385031215613da357600080fd5b8235613dae8161364e565b9150613dbc60208401613506565b90509250929050565b60008083601f840112613dd757600080fd5b50813567ffffffffffffffff811115613def57600080fd5b6020830191508360208260051b850101111561358c57600080fd5b60008060008060008060008060a0898b031215613e2657600080fd5b613e2f89613506565b9750613e3d60208a01613506565b9650604089013567ffffffffffffffff811115613e5957600080fd5b613e658b828c01613dc5565b909750955050606089013567ffffffffffffffff811115613e8557600080fd5b613e918b828c01613dc5565b909550935050608089013567ffffffffffffffff811115613eb157600080fd5b613ebd8b828c0161354a565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613eea57600080fd5b613ef387613506565b9550613f0160208801613506565b94506040870135935060608701359250608087013567ffffffffffffffff811115613f2b57600080fd5b613f3789828a0161354a565b979a9699509497509295939492505050565b600080600060608486031215613f5e57600080fd5b83359250613f6e60208501613506565b915060408401356bffffffffffffffffffffffff81168114613f8f57600080fd5b809150509250925092565b600082825180855260208501945060208160051b8301016020850160005b8381101561406a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e0604086015261402660e0860182613776565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613fb8565b50909695505050505050565b600081518084526020840193506020830160005b828110156140be57815173ffffffffffffffffffffffffffffffffffffffff1686526020958601959091019060010161408a565b5093949350505050565b805160ff168252600060208201516140e4602085018215159052565b5060408201516101206040850152614100610120850182613f9a565b9050606083015160608501526080830151608085015260a083015184820360a086015261412d8282613776565b91505060c083015160c085015260e083015160e085015261010083015184820361010086015261415d8282614076565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006141c260408301866140c8565b82810360208401526141d5818587614166565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610579576105796141df565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561426057600080fd5b8386111561426d57600080fd5b5050820193919092039150565b60608152600061428d60608301866140c8565b60208301949094525060400152919050565b6060815260006142b260608301866140c8565b84602084015282810360408401526141d58185613776565b81810381811115610579576105796141df565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061430c6040830184613776565b949350505050565b6000604082840312801561432757600080fd5b506040805190810167ffffffffffffffff8111828210171561434b5761434b613922565b604052825181526020928301519281019290925250919050565b606081526000614379606083018688614166565b6020830194909452506040015292915050565b83815260406020820152600061415d604083018486614166565b6000602082840312156143b857600080fd5b815161065c8161364e565b6080815260006143d660808301886140c8565b86602084015273ffffffffffffffffffffffffffffffffffffffff86166040840152828103606084015261440b818587614166565b98975050505050505050565b60006020828403121561442957600080fd5b5051919050565b6040815260006144436040830185613776565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b6000835161447a818460208801613752565b919091019182525060200191905056fea2646970667358221220563c18f240859c4cde3bd5a555565b3053b825ff8d04364a7098fee06a3a4b1d64736f6c634300081b0033",
}

// WalletStage2SimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage2SimulatorMetaData.ABI instead.
var WalletStage2SimulatorABI = WalletStage2SimulatorMetaData.ABI

// WalletStage2SimulatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage2SimulatorMetaData.Bin instead.
var WalletStage2SimulatorBin = WalletStage2SimulatorMetaData.Bin

// DeployWalletStage2Simulator deploys a new Ethereum contract, binding an instance of WalletStage2Simulator to it.
func DeployWalletStage2Simulator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletStage2Simulator, error) {
	parsed, err := WalletStage2SimulatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage2SimulatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletStage2Simulator{WalletStage2SimulatorCaller: WalletStage2SimulatorCaller{contract: contract}, WalletStage2SimulatorTransactor: WalletStage2SimulatorTransactor{contract: contract}, WalletStage2SimulatorFilterer: WalletStage2SimulatorFilterer{contract: contract}}, nil
}

// WalletStage2SimulatorDeployedBin is the resulting bytecode of the created contract
const WalletStage2SimulatorDeployedBin = "0x6080604052600436106101485760003560e01c80636ea44577116100c0578063b93ea7ad11610074578063ca70785011610059578063ca707850146104a5578063f23a6e61146104c5578063f727ef1c1461050b5761014f565b8063b93ea7ad1461044a578063bc197c811461045d5761014f565b80638c3f5563116100a55780638c3f5563146103c2578063aaf10f42146103e2578063ad55366b146103f75761014f565b80636ea445771461038e5780638943ec02146103a15761014f565b80631f6a1eb9116101175780634fcf3eca116100fc5780634fcf3eca1461033857806351605d801461034b5780636406c3141461036e5761014f565b80631f6a1eb91461030557806329561426146103185761014f565b8063025b22bc14610212578063150b7a02146102255780631626ba7e146102a05780631a9b2337146102c05761014f565b3661014f57005b6004361061021057600061016b6101663683613490565b61052b565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b49291906134f6565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b61021061022036600461352f565b61057f565b34801561023157600080fd5b5061026a610240366004613593565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102ac57600080fd5b5061026a6102bb366004613602565b6105cb565b3480156102cc57600080fd5b506102e06102db36600461367c565b610663565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610297565b610210610313366004613699565b61066e565b34801561032457600080fd5b5061021061033336600461370a565b6106f0565b61021061034636600461367c565b610734565b34801561035757600080fd5b506103606107f6565b604051908152602001610297565b61038161037c366004613699565b610825565b60405161029791906137c0565b61021061039c36600461389e565b610857565b3480156103ad57600080fd5b506102106103bc3660046138e0565b50505050565b3480156103ce57600080fd5b506103606103dd36600461370a565b6108b7565b3480156103ee57600080fd5b506102e06108e3565b34801561040357600080fd5b50610417610412366004613c62565b6108ed565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610297565b610210610458366004613d90565b61091c565b34801561046957600080fd5b5061026a610478366004613e0a565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104b157600080fd5b506103606104c0366004613c62565b6109e1565b3480156104d157600080fd5b5061026a6104e0366004613ed1565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561051757600080fd5b50610210610526366004613f49565b610b4a565b60006105797fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c05565b92915050565b3330146105bf576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105c881610c63565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610622828686610cb8565b5090508061063657506000915061065c9050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b60006105798261052b565b600061067a8585610df9565b905061068e81606001518260800151611217565b60008061069c838686610cb8565b91509150816106dd578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105b6939291906141af565b6106e78184611221565b50505050505050565b33301461072b576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b6105c8816114ca565b33301461076f576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b600061077a8261052b565b73ffffffffffffffffffffffffffffffffffffffff16036107eb576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105b6565b6105c881600061155a565b60006108207fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b606060006108338686610df9565b905060006108408261161a565b905061084c8183611695565b979650505050505050565b333014610892576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b600061089e8383610df9565b905060006108ab8261161a565b90506103bc8183611221565b60006105797f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c05565b6000610820305490565b600080600080600080610904898989600080611c05565b939d929c5060019b5090995097509095509350505050565b333014610957576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b60006109628361052b565b73ffffffffffffffffffffffffffffffffffffffff16146109d3576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105b6565b6109dd828261155a565b5050565b6000808461010001515160016109f7919061420e565b67ffffffffffffffff811115610a0f57610a0f613922565b604051908082528060200260200182016040528015610a38578160200160208202803683370190505b50905060005b85610100015151811015610aaa578561010001518181518110610a6357610a63614221565b6020026020010151828281518110610a7d57610a7d614221565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610a3e565b5033818661010001515181518110610ac457610ac4614221565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610afe868686610cb8565b50905080610b3e578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105b6939291906141af565b50600195945050505050565b333014610b85576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105b6565b610b9e8383836bffffffffffffffffffffffff16611f2a565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c24929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c6b813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610cd057610cd0614221565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610dcf57610d2d8661161a565b9150600080610d3b84611fb9565b9150915042811015610d83576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105b6565b73ffffffffffffffffffffffffffffffffffffffff82161580610dbb575073ffffffffffffffffffffffffffffffffffffffff821633145b15610dcc5760019450505050610df1565b50505b6000806000610de2898989600080611c05565b60019a50985050505050505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610e5d5760006060840152610e6e565b84810135606090811c908401526014015b6007600183901c168015610ec15760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610ed657506001610eff565b83602016602003610ef25750600282019186013560f01c610eff565b50600182019186013560f81c5b8067ffffffffffffffff811115610f1857610f18613922565b604051908082528060200260200182016040528015610fa357816020015b610f906040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610f365790505b50604086015260005b8181101561120c5760018085019489013560f81c90808216900361100b573087604001518381518110610fe157610fe1614221565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611055565b8489013560601c601486018860400151848151811061102c5761102c614221565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036110935784890135602086018860400151848151811061107e5761107e614221565b60200260200101516020018197508281525050505b6004808216900361112b57600385019489013560e81c89868a6110b6848361420e565b926110c393929190614250565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061110e5761110e614221565b602090810291909101015160400152611127818761420e565b9550505b600880821690036111695784890135602086018860400151848151811061115457611154614221565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061118957611189614221565b602002602001015160800190151590811515815250508060201660ff16602014876040015183815181106111bf576111bf614221565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106111f4576111f4614221565b602090810291909101015160c0015250600101610fac565b505050505092915050565b6109dd8282612005565b604081015151600090815b818110156114c35760008460400151828151811061124c5761124c614221565b602002602001015190508060a001518015611265575083155b156112ad576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506114bb565b606081015180158015906112c05750805a105b156112fd5785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105b69392919061427a565b600082608001511561132e57825161132790831561131b578361131d565b5a5b85604001516120f9565b9050611355565b82516020840151611352919084156113465784611348565b5a5b866040015161210f565b90505b8061147e5760c08301516113a8576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a15050506114bb565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114125786846113dd612127565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105b69392919061429f565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161147e5760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a15050506114c3565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b60010161122c565b5050505050565b80611501576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61152a7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610cad565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b60008061162b836020015130612146565b9050600061163884612213565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040810151516060906000908067ffffffffffffffff8111156116ba576116ba613922565b60405190808252806020026020018201604052801561171057816020015b6116fd6040805160608101909152806000815260200160608152602001600081525090565b8152602001906001900390816116d85790505b50925060005b81811015611bfc5760008560400151828151811061173657611736614221565b602002602001015190508060a00151801561174f575083155b15611797576040805188815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611bf4565b606081015180158015906117aa5750805a105b156118435760058684815181106117c3576117c3614221565b60200260200101516000019060058111156117e0576117e0613723565b908160058111156117f3576117f3613723565b9052505a60405160200161180991815260200190565b60405160208183030381529060405286848151811061182a5761182a614221565b6020026020010151602001819052505050505050610579565b60008260800151156118a85760005a84519091506118739084156118675784611869565b5a5b86604001516120f9565b91505a61188090826142ca565b88868151811061189257611892614221565b6020026020010151604001818152505050611902565b60005a845160208601519192506118d19185156118c557856118c7565b5a5b876040015161210f565b91505a6118de90826142ca565b8886815181106118f0576118f0614221565b60200260200101516040018181525050505b80611b475760c08301516119c557604080518a815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1600287858151811061196157611961614221565b602002602001015160000190600581111561197e5761197e613723565b9081600581111561199157611991613723565b90525061199c612127565b8785815181106119ae576119ae614221565b602002602001015160200181905250505050611bf4565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611a6b576003878581518110611a0457611a04614221565b6020026020010151600001906005811115611a2157611a21613723565b90816005811115611a3457611a34613723565b905250611a3f612127565b878581518110611a5157611a51614221565b602002602001015160200181905250505050505050610579565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611b4757604080518a8152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3910160405180910390a16004878581518110611ae357611ae3614221565b6020026020010151600001906005811115611b0057611b00613723565b90816005811115611b1357611b13613723565b905250611b1e612127565b878581518110611b3057611b30614221565b602002602001015160200181905250505050611bfc565b604080518a8152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a16001878581518110611b9457611b94614221565b6020026020010151600001906005811115611bb157611bb1613723565b90816005811115611bc457611bc4613723565b905250611bcf612127565b878581518110611be157611be1614221565b6020026020010151602001819052505050505b600101611716565b50505092915050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611c50575073ffffffffffffffffffffffffffffffffffffffff8916155b15611d6c578b82013560601c985060149091019089611d6c5760038201918c013560e81c60008d848e611c83858361420e565b92611c9093929190614250565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611d1b90309085906004016142dd565b6040805180830381865afa158015611d37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d5b9190614314565b9250611d67828561420e565b935050505b82600116600103611da657611d948d8a838f8f87908092611d8f93929190614250565b612469565b97509750975097509750505050611f1d565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611e116001600586901c81169061420e565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611e5c8d61161a565b9350611e7a8d858e8e86908092611e7593929190614250565b6126b5565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611ec5575080518614155b8015611ed5575080602001518511155b15611f19576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105b6565b5050505b9550955095509550959050565b611fb47fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611fe77fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c05565b606081901c956bffffffffffffffffffffffff909116945092505050565b6000612010836108b7565b9050818114158015612020575060005b15612068576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105b6565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121b657466121b9565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080612224836101000151613099565b835190915060ff1661229757600061223f846040015161311c565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611676565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016123275760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016123975760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612309565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016124075760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001612309565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105b6565b60008060008060006124cb604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156126425760038201916000908b013560e81c612513848261420e565b9150600090508a8214612527576000612529565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841461255a578561255c565b8f5b905061257c818e8e8890879261257493929190614250565b600186611c05565b50929d50909b50995097508a8a10156125d9578c8c869085926125a193929190614250565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105b69493929190614365565b829450888e60000151036125ec5760008e525b83881061262f576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105b6565b505060c0840187905291508490506124f3565b8a511580159061265657508a602001518511155b1561269a576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105b6565b6126a38d61161a565b93505050509550955095509550959050565b60008060005b8381101561308f57600181019085013560f881901c9060fc1c806127d857600f821660008190036126f35750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa158015612788573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006127ab828960ff1661319e565b90508b6127b857806127c7565b60008c81526020829052604090205b9b50505050505050505050506126bb565b6001810361283c57600f821660008190036127fa5750600183019287013560f81c5b601484019388013560601c60006128148260ff851661319e565b9050866128215780612830565b60008781526020829052604090205b965050505050506126bb565b60028103612a235760038216600081900361285e5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261291c93929190614250565b6040518463ffffffff1660e01b815260040161293a9392919061438c565b602060405180830381865afa158015612957573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061297b91906143a6565b7fffffffff0000000000000000000000000000000000000000000000000000000016146129de578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105b69594939291906143c3565b8097508460ff168a01995060006129f8858760ff1661319e565b905089612a055780612a14565b60008a81526020829052604090205b995050505050505050506126bb565b60038103612a5757602083019287013584612a3e5780612a4d565b60008581526020829052604090205b94505050506126bb565b60048103612afd57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c16858301809650819250505060008186019050600080612acc8e8e8e8e8c908892611e7593929190614250565b91509150829750818a019950612aec898260009182526020526040902090565b9850829750505050505050506126bb565b60068103612bc7576003600283901c166000819003612b235750600183019287013560f81c5b600383166000819003612b3d5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff169150809750819250505060008187019050600080612b7b8f8f8f8f8d908892611e7593929190614250565b91509150829850848210612b8e57998501995b6000612b9b828789613205565b90508a612ba85780612bb7565b60008b81526020829052604090205b9a505050505050505050506126bb565b60058103612c34576020830192870135888103612c02577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b6000612c0d82613269565b905085612c1a5780612c29565b60008681526020829052604090205b9550505050506126bb565b60078103612d3a57600f82166000819003612c565750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612766565b60088103612d8e5760208301928701356000612d568b826132bd565b9050808203612d83577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061281482613338565b60098103612ef457600382166000819003612db05750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612e4a93929190614250565b6040518463ffffffff1660e01b8152600401612e68939291906141af565b602060405180830381865afa158015612e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ea99190614417565b90508197508460ff168a0199506000612ec6858760ff1684613373565b905089612ed35780612ee2565b60008a81526020829052604090205b995082985050505050505050506126bb565b600a810361305a57600382166000819003612f165750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612faf93929190614250565b6040518463ffffffff1660e01b8152600401612fcd9392919061438c565b602060405180830381865afa158015612fea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061300e9190614417565b90508198508560ff168b019a50600061302b868860ff1684613373565b90508a6130385780613047565b60008b81526020829052604090205b9a508299505050505050505050506126bb565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105b6565b5094509492505050565b6000606060005b835181101561310d57818482815181106130bc576130bc614221565b60200260200101516040516020016130d5929190614430565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905291506001016130a0565b50805160209091012092915050565b6000606060005b835181101561310d57600061315085838151811061314357613143614221565b60200260200101516133e1565b90508281604051602001613165929190614468565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101613123565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b166031820152604581018290526000906065016121f5565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b6000806132ce846020015184612146565b905060006132db85612213565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a60208201529081018290526000906060016132a0565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d0161324a565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016132a098979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff0000000000000000000000000000000000000000000000000000000081169060048410156134ef577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461352a57600080fd5b919050565b60006020828403121561354157600080fd5b61065c82613506565b60008083601f84011261355c57600080fd5b50813567ffffffffffffffff81111561357457600080fd5b60208301915083602082850101111561358c57600080fd5b9250929050565b6000806000806000608086880312156135ab57600080fd5b6135b486613506565b94506135c260208701613506565b935060408601359250606086013567ffffffffffffffff8111156135e557600080fd5b6135f18882890161354a565b969995985093965092949392505050565b60008060006040848603121561361757600080fd5b83359250602084013567ffffffffffffffff81111561363557600080fd5b6136418682870161354a565b9497909650939450505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105c857600080fd5b60006020828403121561368e57600080fd5b813561065c8161364e565b600080600080604085870312156136af57600080fd5b843567ffffffffffffffff8111156136c657600080fd5b6136d28782880161354a565b909550935050602085013567ffffffffffffffff8111156136f257600080fd5b6136fe8782880161354a565b95989497509550505050565b60006020828403121561371c57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b8381101561376d578181015183820152602001613755565b50506000910152565b6000815180845261378e816020860160208601613752565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613892577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08786030184528151805160068110613852577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8087525060208101516060602088015261386f6060880182613776565b6040928301519790920196909652945060209384019391909101906001016137e8565b50929695505050505050565b600080602083850312156138b157600080fd5b823567ffffffffffffffff8111156138c857600080fd5b6138d48582860161354a565b90969095509350505050565b600080600080606085870312156138f657600080fd5b6138ff85613506565b935060208501359250604085013567ffffffffffffffff8111156136f257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561397457613974613922565b60405290565b604051610120810167ffffffffffffffff8111828210171561397457613974613922565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156139e5576139e5613922565b604052919050565b803560ff8116811461352a57600080fd5b8035801515811461352a57600080fd5b600067ffffffffffffffff821115613a2857613a28613922565b5060051b60200190565b600082601f830112613a4357600080fd5b813567ffffffffffffffff811115613a5d57613a5d613922565b613a8e60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8401160161399e565b818152846020838601011115613aa357600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112613ad157600080fd5b8135613ae4613adf82613a0e565b61399e565b8082825260208201915060208360051b860101925085831115613b0657600080fd5b602085015b83811015613bf357803567ffffffffffffffff811115613b2a57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0011215613b5e57600080fd5b613b66613951565b613b7260208301613506565b815260408201356020820152606082013567ffffffffffffffff811115613b9857600080fd5b613ba78a602083860101613a32565b60408301525060808201356060820152613bc360a083016139fe565b6080820152613bd460c083016139fe565b60a082015260e0919091013560c0820152835260209283019201613b0b565b5095945050505050565b600082601f830112613c0e57600080fd5b8135613c1c613adf82613a0e565b8082825260208201915060208360051b860101925085831115613c3e57600080fd5b602085015b83811015613bf357613c5481613506565b835260209283019201613c43565b600080600060408486031215613c7757600080fd5b833567ffffffffffffffff811115613c8e57600080fd5b84016101208187031215613ca157600080fd5b613ca961397a565b613cb2826139ed565b8152613cc0602083016139fe565b6020820152604082013567ffffffffffffffff811115613cdf57600080fd5b613ceb88828501613ac0565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff811115613d1f57600080fd5b613d2b88828501613a32565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff811115613d6057600080fd5b613d6c88828501613bfd565b61010083015250935050602084013567ffffffffffffffff81111561363557600080fd5b60008060408385031215613da357600080fd5b8235613dae8161364e565b9150613dbc60208401613506565b90509250929050565b60008083601f840112613dd757600080fd5b50813567ffffffffffffffff811115613def57600080fd5b6020830191508360208260051b850101111561358c57600080fd5b60008060008060008060008060a0898b031215613e2657600080fd5b613e2f89613506565b9750613e3d60208a01613506565b9650604089013567ffffffffffffffff811115613e5957600080fd5b613e658b828c01613dc5565b909750955050606089013567ffffffffffffffff811115613e8557600080fd5b613e918b828c01613dc5565b909550935050608089013567ffffffffffffffff811115613eb157600080fd5b613ebd8b828c0161354a565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613eea57600080fd5b613ef387613506565b9550613f0160208801613506565b94506040870135935060608701359250608087013567ffffffffffffffff811115613f2b57600080fd5b613f3789828a0161354a565b979a9699509497509295939492505050565b600080600060608486031215613f5e57600080fd5b83359250613f6e60208501613506565b915060408401356bffffffffffffffffffffffff81168114613f8f57600080fd5b809150509250925092565b600082825180855260208501945060208160051b8301016020850160005b8381101561406a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e0604086015261402660e0860182613776565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613fb8565b50909695505050505050565b600081518084526020840193506020830160005b828110156140be57815173ffffffffffffffffffffffffffffffffffffffff1686526020958601959091019060010161408a565b5093949350505050565b805160ff168252600060208201516140e4602085018215159052565b5060408201516101206040850152614100610120850182613f9a565b9050606083015160608501526080830151608085015260a083015184820360a086015261412d8282613776565b91505060c083015160c085015260e083015160e085015261010083015184820361010086015261415d8282614076565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6040815260006141c260408301866140c8565b82810360208401526141d5818587614166565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610579576105796141df565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000808585111561426057600080fd5b8386111561426d57600080fd5b5050820193919092039150565b60608152600061428d60608301866140c8565b60208301949094525060400152919050565b6060815260006142b260608301866140c8565b84602084015282810360408401526141d58185613776565b81810381811115610579576105796141df565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061430c6040830184613776565b949350505050565b6000604082840312801561432757600080fd5b506040805190810167ffffffffffffffff8111828210171561434b5761434b613922565b604052825181526020928301519281019290925250919050565b606081526000614379606083018688614166565b6020830194909452506040015292915050565b83815260406020820152600061415d604083018486614166565b6000602082840312156143b857600080fd5b815161065c8161364e565b6080815260006143d660808301886140c8565b86602084015273ffffffffffffffffffffffffffffffffffffffff86166040840152828103606084015261440b818587614166565b98975050505050505050565b60006020828403121561442957600080fd5b5051919050565b6040815260006144436040830185613776565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b6000835161447a818460208801613752565b919091019182525060200191905056fea2646970667358221220563c18f240859c4cde3bd5a555565b3053b825ff8d04364a7098fee06a3a4b1d64736f6c634300081b0033"

// WalletStage2Simulator is an auto generated Go binding around an Ethereum contract.
type WalletStage2Simulator struct {
	WalletStage2SimulatorCaller     // Read-only binding to the contract
	WalletStage2SimulatorTransactor // Write-only binding to the contract
	WalletStage2SimulatorFilterer   // Log filterer for contract events
}

// WalletStage2SimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletStage2SimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2SimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletStage2SimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2SimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletStage2SimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2SimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletStage2SimulatorSession struct {
	Contract     *WalletStage2Simulator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletStage2SimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletStage2SimulatorCallerSession struct {
	Contract *WalletStage2SimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WalletStage2SimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletStage2SimulatorTransactorSession struct {
	Contract     *WalletStage2SimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WalletStage2SimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletStage2SimulatorRaw struct {
	Contract *WalletStage2Simulator // Generic contract binding to access the raw methods on
}

// WalletStage2SimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletStage2SimulatorCallerRaw struct {
	Contract *WalletStage2SimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// WalletStage2SimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletStage2SimulatorTransactorRaw struct {
	Contract *WalletStage2SimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletStage2Simulator creates a new instance of WalletStage2Simulator, bound to a specific deployed contract.
func NewWalletStage2Simulator(address common.Address, backend bind.ContractBackend) (*WalletStage2Simulator, error) {
	contract, err := bindWalletStage2Simulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletStage2Simulator{WalletStage2SimulatorCaller: WalletStage2SimulatorCaller{contract: contract}, WalletStage2SimulatorTransactor: WalletStage2SimulatorTransactor{contract: contract}, WalletStage2SimulatorFilterer: WalletStage2SimulatorFilterer{contract: contract}}, nil
}

// NewWalletStage2SimulatorCaller creates a new read-only instance of WalletStage2Simulator, bound to a specific deployed contract.
func NewWalletStage2SimulatorCaller(address common.Address, caller bind.ContractCaller) (*WalletStage2SimulatorCaller, error) {
	contract, err := bindWalletStage2Simulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorCaller{contract: contract}, nil
}

// NewWalletStage2SimulatorTransactor creates a new write-only instance of WalletStage2Simulator, bound to a specific deployed contract.
func NewWalletStage2SimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletStage2SimulatorTransactor, error) {
	contract, err := bindWalletStage2Simulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorTransactor{contract: contract}, nil
}

// NewWalletStage2SimulatorFilterer creates a new log filterer instance of WalletStage2Simulator, bound to a specific deployed contract.
func NewWalletStage2SimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletStage2SimulatorFilterer, error) {
	contract, err := bindWalletStage2Simulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorFilterer{contract: contract}, nil
}

// bindWalletStage2Simulator binds a generic wrapper to an already deployed contract.
func bindWalletStage2Simulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletStage2SimulatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage2Simulator *WalletStage2SimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage2Simulator.Contract.WalletStage2SimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage2Simulator *WalletStage2SimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.WalletStage2SimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage2Simulator *WalletStage2SimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.WalletStage2SimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage2Simulator *WalletStage2SimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage2Simulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) GetImplementation() (common.Address, error) {
	return _WalletStage2Simulator.Contract.GetImplementation(&_WalletStage2Simulator.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) GetImplementation() (common.Address, error) {
	return _WalletStage2Simulator.Contract.GetImplementation(&_WalletStage2Simulator.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) ImageHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) ImageHash() ([32]byte, error) {
	return _WalletStage2Simulator.Contract.ImageHash(&_WalletStage2Simulator.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) ImageHash() ([32]byte, error) {
	return _WalletStage2Simulator.Contract.ImageHash(&_WalletStage2Simulator.CallOpts)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) IsValidSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "isValidSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2Simulator.Contract.IsValidSapientSignature(&_WalletStage2Simulator.CallOpts, _payload, _signature)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2Simulator.Contract.IsValidSapientSignature(&_WalletStage2Simulator.CallOpts, _payload, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.IsValidSignature(&_WalletStage2Simulator.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.IsValidSignature(&_WalletStage2Simulator.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC1155BatchReceived(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC1155BatchReceived(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC1155Received(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC1155Received(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC721Received(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage2Simulator.Contract.OnERC721Received(&_WalletStage2Simulator.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage2Simulator.Contract.ReadHook(&_WalletStage2Simulator.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage2Simulator.Contract.ReadHook(&_WalletStage2Simulator.CallOpts, signature)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage2Simulator.Contract.ReadNonce(&_WalletStage2Simulator.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage2Simulator.Contract.ReadNonce(&_WalletStage2Simulator.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage2Simulator *WalletStage2SimulatorCaller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletStage2Simulator.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletStage2Simulator *WalletStage2SimulatorSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage2Simulator.Contract.RecoverPartialSignature(&_WalletStage2Simulator.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage2Simulator *WalletStage2SimulatorCallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage2Simulator.Contract.RecoverPartialSignature(&_WalletStage2Simulator.CallOpts, _payload, _signature)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.AddHook(&_WalletStage2Simulator.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.AddHook(&_WalletStage2Simulator.TransactOpts, signature, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Execute(&_WalletStage2Simulator.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Execute(&_WalletStage2Simulator.TransactOpts, _payload, _signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.RemoveHook(&_WalletStage2Simulator.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.RemoveHook(&_WalletStage2Simulator.TransactOpts, signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.SelfExecute(&_WalletStage2Simulator.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.SelfExecute(&_WalletStage2Simulator.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.SetStaticSignature(&_WalletStage2Simulator.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.SetStaticSignature(&_WalletStage2Simulator.TransactOpts, _hash, _address, _timestamp)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) Simulate(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "simulate", _payload, _signature)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage2Simulator *WalletStage2SimulatorSession) Simulate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Simulate(&_WalletStage2Simulator.TransactOpts, _payload, _signature)
}

// Simulate is a paid mutator transaction binding the contract method 0x6406c314.
//
// Solidity: function simulate(bytes _payload, bytes _signature) payable returns((uint8,bytes,uint256)[] results)
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) Simulate(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Simulate(&_WalletStage2Simulator.TransactOpts, _payload, _signature)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.TokenReceived(&_WalletStage2Simulator.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.TokenReceived(&_WalletStage2Simulator.TransactOpts, arg0, arg1, arg2)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.UpdateImageHash(&_WalletStage2Simulator.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.UpdateImageHash(&_WalletStage2Simulator.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.UpdateImplementation(&_WalletStage2Simulator.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.UpdateImplementation(&_WalletStage2Simulator.TransactOpts, _implementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Fallback(&_WalletStage2Simulator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Fallback(&_WalletStage2Simulator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2Simulator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorSession) Receive() (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Receive(&_WalletStage2Simulator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2Simulator *WalletStage2SimulatorTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletStage2Simulator.Contract.Receive(&_WalletStage2Simulator.TransactOpts)
}

// WalletStage2SimulatorCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallAbortedIterator struct {
	Event *WalletStage2SimulatorCallAborted // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorCallAborted)
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
		it.Event = new(WalletStage2SimulatorCallAborted)
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
func (it *WalletStage2SimulatorCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorCallAborted represents a CallAborted event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallAborted struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletStage2SimulatorCallAbortedIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorCallAbortedIterator{contract: _WalletStage2Simulator.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorCallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorCallAborted)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseCallAborted(log types.Log) (*WalletStage2SimulatorCallAborted, error) {
	event := new(WalletStage2SimulatorCallAborted)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallFailedIterator struct {
	Event *WalletStage2SimulatorCallFailed // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorCallFailed)
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
		it.Event = new(WalletStage2SimulatorCallFailed)
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
func (it *WalletStage2SimulatorCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorCallFailed represents a CallFailed event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallFailed struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletStage2SimulatorCallFailedIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorCallFailedIterator{contract: _WalletStage2Simulator.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorCallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorCallFailed)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseCallFailed(log types.Log) (*WalletStage2SimulatorCallFailed, error) {
	event := new(WalletStage2SimulatorCallFailed)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallSkippedIterator struct {
	Event *WalletStage2SimulatorCallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorCallSkipped)
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
		it.Event = new(WalletStage2SimulatorCallSkipped)
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
func (it *WalletStage2SimulatorCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorCallSkipped represents a CallSkipped event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletStage2SimulatorCallSkippedIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorCallSkippedIterator{contract: _WalletStage2Simulator.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorCallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorCallSkipped)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseCallSkipped(log types.Log) (*WalletStage2SimulatorCallSkipped, error) {
	event := new(WalletStage2SimulatorCallSkipped)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorCallSuccessIterator is returned from FilterCallSuccess and is used to iterate over the raw logs and unpacked data for CallSuccess events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallSuccessIterator struct {
	Event *WalletStage2SimulatorCallSuccess // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorCallSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorCallSuccess)
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
		it.Event = new(WalletStage2SimulatorCallSuccess)
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
func (it *WalletStage2SimulatorCallSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorCallSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorCallSuccess represents a CallSuccess event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorCallSuccess struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSuccess is a free log retrieval operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterCallSuccess(opts *bind.FilterOpts) (*WalletStage2SimulatorCallSuccessIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorCallSuccessIterator{contract: _WalletStage2Simulator.contract, event: "CallSuccess", logs: logs, sub: sub}, nil
}

// WatchCallSuccess is a free log subscription operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchCallSuccess(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorCallSuccess) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorCallSuccess)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseCallSuccess(log types.Log) (*WalletStage2SimulatorCallSuccess, error) {
	event := new(WalletStage2SimulatorCallSuccess)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorDefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorDefinedHookIterator struct {
	Event *WalletStage2SimulatorDefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorDefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorDefinedHook)
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
		it.Event = new(WalletStage2SimulatorDefinedHook)
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
func (it *WalletStage2SimulatorDefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorDefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorDefinedHook represents a DefinedHook event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorDefinedHook struct {
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage2SimulatorDefinedHookIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorDefinedHookIterator{contract: _WalletStage2Simulator.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorDefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorDefinedHook)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseDefinedHook(log types.Log) (*WalletStage2SimulatorDefinedHook, error) {
	event := new(WalletStage2SimulatorDefinedHook)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorImageHashUpdatedIterator struct {
	Event *WalletStage2SimulatorImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorImageHashUpdated)
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
		it.Event = new(WalletStage2SimulatorImageHashUpdated)
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
func (it *WalletStage2SimulatorImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorImageHashUpdated represents a ImageHashUpdated event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletStage2SimulatorImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorImageHashUpdatedIterator{contract: _WalletStage2Simulator.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorImageHashUpdated)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseImageHashUpdated(log types.Log) (*WalletStage2SimulatorImageHashUpdated, error) {
	event := new(WalletStage2SimulatorImageHashUpdated)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorImplementationUpdatedIterator struct {
	Event *WalletStage2SimulatorImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorImplementationUpdated)
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
		it.Event = new(WalletStage2SimulatorImplementationUpdated)
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
func (it *WalletStage2SimulatorImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorImplementationUpdated represents a ImplementationUpdated event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletStage2SimulatorImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorImplementationUpdatedIterator{contract: _WalletStage2Simulator.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorImplementationUpdated)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseImplementationUpdated(log types.Log) (*WalletStage2SimulatorImplementationUpdated, error) {
	event := new(WalletStage2SimulatorImplementationUpdated)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorNonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorNonceChangeIterator struct {
	Event *WalletStage2SimulatorNonceChange // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorNonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorNonceChange)
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
		it.Event = new(WalletStage2SimulatorNonceChange)
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
func (it *WalletStage2SimulatorNonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorNonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorNonceChange represents a NonceChange event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorNonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletStage2SimulatorNonceChangeIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorNonceChangeIterator{contract: _WalletStage2Simulator.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorNonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorNonceChange)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseNonceChange(log types.Log) (*WalletStage2SimulatorNonceChange, error) {
	event := new(WalletStage2SimulatorNonceChange)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SimulatorStaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorStaticSignatureSetIterator struct {
	Event *WalletStage2SimulatorStaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletStage2SimulatorStaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2SimulatorStaticSignatureSet)
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
		it.Event = new(WalletStage2SimulatorStaticSignatureSet)
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
func (it *WalletStage2SimulatorStaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SimulatorStaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2SimulatorStaticSignatureSet represents a StaticSignatureSet event raised by the WalletStage2Simulator contract.
type WalletStage2SimulatorStaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletStage2SimulatorStaticSignatureSetIterator, error) {

	logs, sub, err := _WalletStage2Simulator.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SimulatorStaticSignatureSetIterator{contract: _WalletStage2Simulator.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletStage2SimulatorStaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletStage2Simulator.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2SimulatorStaticSignatureSet)
				if err := _WalletStage2Simulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletStage2Simulator *WalletStage2SimulatorFilterer) ParseStaticSignatureSet(log types.Log) (*WalletStage2SimulatorStaticSignatureSet, error) {
	event := new(WalletStage2SimulatorStaticSignatureSet)
	if err := _WalletStage2Simulator.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
