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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"estimate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506141408061001f6000396000f3fe6080604052600436106101485760003560e01c80638943ec02116100c0578063b93ea7ad11610074578063ca70785011610059578063ca70785014610498578063f23a6e61146104b8578063f727ef1c146104fe5761014f565b8063b93ea7ad1461043d578063bc197c81146104505761014f565b8063975befdb116100a5578063975befdb146103c2578063aaf10f42146103d5578063ad55366b146103ea5761014f565b80638943ec02146103815780638c3f5563146103a25761014f565b80631f6a1eb9116101175780634fcf3eca116100fc5780634fcf3eca1461033857806351605d801461034b5780636ea445771461036e5761014f565b80631f6a1eb91461030557806329561426146103185761014f565b8063025b22bc14610212578063150b7a02146102255780631626ba7e146102a05780631a9b2337146102c05761014f565b3661014f57005b6004361061021057600061016b610166368361321d565b61051e565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b4929190613283565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b6102106102203660046132bc565b610572565b34801561023157600080fd5b5061026a610240366004613320565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102ac57600080fd5b5061026a6102bb36600461338f565b6105be565b3480156102cc57600080fd5b506102e06102db366004613409565b610656565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610297565b610210610313366004613426565b610661565b34801561032457600080fd5b50610210610333366004613497565b6106e3565b610210610346366004613409565b610727565b34801561035757600080fd5b506103606107e9565b604051908152602001610297565b61021061037c3660046134b0565b610818565b34801561038d57600080fd5b5061021061039c3660046134f2565b50505050565b3480156103ae57600080fd5b506103606103bd366004613497565b610878565b6103606103d0366004613426565b6108a4565b3480156103e157600080fd5b506102e0610941565b3480156103f657600080fd5b5061040a610405366004613874565b61094b565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610297565b61021061044b3660046139a2565b610985565b34801561045c57600080fd5b5061026a61046b366004613a1c565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104a457600080fd5b506103606104b3366004613874565b610a4a565b3480156104c457600080fd5b5061026a6104d3366004613ae3565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561050a57600080fd5b50610210610519366004613b5b565b610bb3565b600061056c7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c6e565b92915050565b3330146105b2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105bb81610ccc565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610615828686610d21565b5090508061062957506000915061064f9050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061056c8261051e565b600061066d8585610e6f565b90506106818160600151826080015161128d565b60008061068f838686610d21565b91509150816106d0578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b6106da8184611375565b50505050505050565b33301461071e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6105bb8161161f565b333014610762576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b600061076d8261051e565b73ffffffffffffffffffffffffffffffffffffffff16036107de576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105a9565b6105bb8160006116af565b60006108137fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610853576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b600061085f8383610e6f565b9050600061086c8261176f565b905061039c8183611375565b600061056c7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c6e565b6000805a905060006108b68787610e6f565b60608101519091506108d0906108cb81610878565b61128d565b6000806108de838888610d21565b915091508161091f578287876040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b61092981846117ea565b5a6109349085613e8e565b9998505050505050505050565b6000610813305490565b600080600080600080610962898989600080611a34565b93995091975094509250905061097783611d59565b935093975093979195509350565b3330146109c0576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60006109cb8361051e565b73ffffffffffffffffffffffffffffffffffffffff1614610a3c576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105a9565b610a4682826116af565b5050565b600080846101000151516001610a609190613ea1565b67ffffffffffffffff811115610a7857610a78613534565b604051908082528060200260200182016040528015610aa1578160200160208202803683370190505b50905060005b85610100015151811015610b13578561010001518181518110610acc57610acc613eb4565b6020026020010151828281518110610ae657610ae6613eb4565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610aa7565b5033818661010001515181518110610b2d57610b2d613eb4565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610b67868686610d21565b50905080610ba7578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b50600195945050505050565b333014610bee576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b610c078383836bffffffffffffffffffffffff16611d6d565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c8d929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610cd4813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610d3957610d39613eb4565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e3857610d968661176f565b9150600080610da484611dfc565b9150915042811015610dec576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105a9565b73ffffffffffffffffffffffffffffffffffffffff82161580610e24575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e355760019450505050610e67565b50505b6000806000610e4b898989600080611a34565b98509295509093509150610e60905081611d59565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610ed35760006060840152610ee4565b84810135606090811c908401526014015b6007600183901c168015610f375760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610f4c57506001610f75565b83602016602003610f685750600282019186013560f01c610f75565b50600182019186013560f81c5b8067ffffffffffffffff811115610f8e57610f8e613534565b60405190808252806020026020018201604052801561101957816020015b6110066040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610fac5790505b50604086015260005b818110156112825760018085019489013560f81c90808216900361108157308760400151838151811061105757611057613eb4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526110cb565b8489013560601c60148601886040015184815181106110a2576110a2613eb4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611109578489013560208601886040015184815181106110f4576110f4613eb4565b60200260200101516020018197508281525050505b600480821690036111a157600385019489013560e81c89868a61112c8483613ea1565b9261113993929190613ee3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061118457611184613eb4565b60209081029190910101516040015261119d8187613ea1565b9550505b600880821690036111df578489013560208601886040015184815181106111ca576111ca613eb4565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106111ff576111ff613eb4565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061123557611235613eb4565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061126a5761126a613eb4565b602090810291909101015160c0015250600101611022565b505050505092915050565b600061129883610878565b90508181146112e4576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105a9565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b81811015611618576000846040015182815181106113a0576113a0613eb4565b602002602001015190508060a0015180156113b9575083155b15611401576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611610565b606081015180158015906114145750805a105b156114515785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a993929190613f0d565b600082608001511561148257825161147b90831561146f5783611471565b5a5b8560400151611e48565b90506114a9565b825160208401516114a69190841561149a578461149c565b5a5b8660400151611e5e565b90505b806115d35760c08301516114fc576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611610565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611566578684611531611e76565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105a993929190613f32565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016115d35760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c391015b60405180910390a1505050611618565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b600101611380565b5050505050565b80611656576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61167f7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610d16565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b600080611780836020015130611e95565b9050600061178d84611f62565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604081015151600090815b818110156116185760008460400151828151811061181557611815613eb4565b602002602001015190508060a00151801561182e575083155b15611876576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611a2c565b606081015180158015906118895750805a105b156118c65785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a993929190613f0d565b60008260800151156118eb5782516118e490831561146f5783611471565b9050611906565b825160208401516119039190841561149a578461149c565b90505b806119ef5760c0830151611959576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611a2c565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161198e578684611531611e76565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016119ef5760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c391016115c3565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b6001016117f5565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611a7f575073ffffffffffffffffffffffffffffffffffffffff8916155b15611b9b578b82013560601c985060149091019089611b9b5760038201918c013560e81c60008d848e611ab28583613ea1565b92611abf93929190613ee3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611b4a9030908590600401613f5d565b6040805180830381865afa158015611b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8a9190613f94565b9250611b968285613ea1565b935050505b82600116600103611bd557611bc38d8a838f8f87908092611bbe93929190613ee3565b6121b8565b97509750975097509750505050611d4c565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611c406001600586901c811690613ea1565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611c8b8d61176f565b9350611ca98d858e8e86908092611ca493929190613ee3565b612404565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611cf4575080518614155b8015611d04575080602001518511155b15611d48576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105a9565b5050505b9550955095509550959050565b6000611d6482612de8565b50600192915050565b611df77fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611e2a7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c6e565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611f055746611f08565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611f73836101000151612df3565b835190915060ff16611fe6576000611f8e8460400151612e76565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016117cb565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016120765760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016120e65760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612058565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016121565760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001612058565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105a9565b600080600080600061221a604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156123915760038201916000908b013560e81c6122628482613ea1565b9150600090508a8214612276576000612278565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146122a957856122ab565b8f5b90506122cb818e8e889087926122c393929190613ee3565b600186611a34565b50929d50909b50995097508a8a1015612328578c8c869085926122f093929190613ee3565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105a99493929190613fe5565b829450888e600001510361233b5760008e525b83881061237e576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105a9565b505060c084018790529150849050612242565b8a51158015906123a557508a602001518511155b156123e9576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105a9565b6123f28d61176f565b93505050509550955095509550959050565b60008060005b83811015612dde57600181019085013560f881901c9060fc1c8061252757600f821660008190036124425750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156124d7573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006124fa828960ff16612ef8565b90508b6125075780612516565b60008c81526020829052604090205b9b505050505050505050505061240a565b6001810361258b57600f821660008190036125495750600183019287013560f81c5b601484019388013560601c60006125638260ff8516612ef8565b905086612570578061257f565b60008781526020829052604090205b9650505050505061240a565b60028103612772576003821660008190036125ad5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261266b93929190613ee3565b6040518463ffffffff1660e01b81526004016126899392919061400c565b602060405180830381865afa1580156126a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ca9190614026565b7fffffffff00000000000000000000000000000000000000000000000000000000161461272d578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105a9959493929190614043565b8097508460ff168a0199506000612747858760ff16612ef8565b9050896127545780612763565b60008a81526020829052604090205b9950505050505050505061240a565b600381036127a65760208301928701358461278d578061279c565b60008581526020829052604090205b945050505061240a565b6004810361284c57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c1685830180965081925050506000818601905060008061281b8e8e8e8e8c908892611ca493929190613ee3565b91509150829750818a01995061283b898260009182526020526040902090565b98508297505050505050505061240a565b60068103612916576003600283901c1660008190036128725750600183019287013560f81c5b60038316600081900361288c5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806128ca8f8f8f8f8d908892611ca493929190613ee3565b915091508298508482106128dd57998501995b60006128ea828789612f5f565b90508a6128f75780612906565b60008b81526020829052604090205b9a5050505050505050505061240a565b60058103612983576020830192870135888103612951577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061295c82612fc3565b9050856129695780612978565b60008681526020829052604090205b95505050505061240a565b60078103612a8957600f821660008190036129a55750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016124b5565b60088103612add5760208301928701356000612aa58b82613017565b9050808203612ad2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061256382613092565b60098103612c4357600382166000819003612aff5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612b9993929190613ee3565b6040518463ffffffff1660e01b8152600401612bb793929190613e2f565b602060405180830381865afa158015612bd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf89190614097565b90508197508460ff168a0199506000612c15858760ff16846130cd565b905089612c225780612c31565b60008a81526020829052604090205b9950829850505050505050505061240a565b600a8103612da957600382166000819003612c655750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612cfe93929190613ee3565b6040518463ffffffff1660e01b8152600401612d1c9392919061400c565b602060405180830381865afa158015612d39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d5d9190614097565b90508198508560ff168b019a506000612d7a868860ff16846130cd565b90508a612d875780612d96565b60008b81526020829052604090205b9a5082995050505050505050505061240a565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105a9565b5094509492505050565b600061056c8261313b565b6000606060005b8351811015612e675781848281518110612e1657612e16613eb4565b6020026020010151604051602001612e2f9291906140b0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612dfa565b50805160209091012092915050565b6000606060005b8351811015612e67576000612eaa858381518110612e9d57612e9d613eb4565b602002602001015161316e565b90508281604051602001612ebf9291906140e8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612e7d565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611f44565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080613028846020015184611e95565b9050600061303585611f62565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612ffa565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612fa4565b6000811580159061056c5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612ffa98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561327c577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146132b757600080fd5b919050565b6000602082840312156132ce57600080fd5b61064f82613293565b60008083601f8401126132e957600080fd5b50813567ffffffffffffffff81111561330157600080fd5b60208301915083602082850101111561331957600080fd5b9250929050565b60008060008060006080868803121561333857600080fd5b61334186613293565b945061334f60208701613293565b935060408601359250606086013567ffffffffffffffff81111561337257600080fd5b61337e888289016132d7565b969995985093965092949392505050565b6000806000604084860312156133a457600080fd5b83359250602084013567ffffffffffffffff8111156133c257600080fd5b6133ce868287016132d7565b9497909650939450505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105bb57600080fd5b60006020828403121561341b57600080fd5b813561064f816133db565b6000806000806040858703121561343c57600080fd5b843567ffffffffffffffff81111561345357600080fd5b61345f878288016132d7565b909550935050602085013567ffffffffffffffff81111561347f57600080fd5b61348b878288016132d7565b95989497509550505050565b6000602082840312156134a957600080fd5b5035919050565b600080602083850312156134c357600080fd5b823567ffffffffffffffff8111156134da57600080fd5b6134e6858286016132d7565b90969095509350505050565b6000806000806060858703121561350857600080fd5b61351185613293565b935060208501359250604085013567ffffffffffffffff81111561347f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561358657613586613534565b60405290565b604051610120810167ffffffffffffffff8111828210171561358657613586613534565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156135f7576135f7613534565b604052919050565b803560ff811681146132b757600080fd5b803580151581146132b757600080fd5b600067ffffffffffffffff82111561363a5761363a613534565b5060051b60200190565b600082601f83011261365557600080fd5b813567ffffffffffffffff81111561366f5761366f613534565b6136a060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016135b0565b8181528460208386010111156136b557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126136e357600080fd5b81356136f66136f182613620565b6135b0565b8082825260208201915060208360051b86010192508583111561371857600080fd5b602085015b8381101561380557803567ffffffffffffffff81111561373c57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561377057600080fd5b613778613563565b61378460208301613293565b815260408201356020820152606082013567ffffffffffffffff8111156137aa57600080fd5b6137b98a602083860101613644565b604083015250608082013560608201526137d560a08301613610565b60808201526137e660c08301613610565b60a082015260e0919091013560c082015283526020928301920161371d565b5095945050505050565b600082601f83011261382057600080fd5b813561382e6136f182613620565b8082825260208201915060208360051b86010192508583111561385057600080fd5b602085015b838110156138055761386681613293565b835260209283019201613855565b60008060006040848603121561388957600080fd5b833567ffffffffffffffff8111156138a057600080fd5b840161012081870312156138b357600080fd5b6138bb61358c565b6138c4826135ff565b81526138d260208301613610565b6020820152604082013567ffffffffffffffff8111156138f157600080fd5b6138fd888285016136d2565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561393157600080fd5b61393d88828501613644565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561397257600080fd5b61397e8882850161380f565b61010083015250935050602084013567ffffffffffffffff8111156133c257600080fd5b600080604083850312156139b557600080fd5b82356139c0816133db565b91506139ce60208401613293565b90509250929050565b60008083601f8401126139e957600080fd5b50813567ffffffffffffffff811115613a0157600080fd5b6020830191508360208260051b850101111561331957600080fd5b60008060008060008060008060a0898b031215613a3857600080fd5b613a4189613293565b9750613a4f60208a01613293565b9650604089013567ffffffffffffffff811115613a6b57600080fd5b613a778b828c016139d7565b909750955050606089013567ffffffffffffffff811115613a9757600080fd5b613aa38b828c016139d7565b909550935050608089013567ffffffffffffffff811115613ac357600080fd5b613acf8b828c016132d7565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613afc57600080fd5b613b0587613293565b9550613b1360208801613293565b94506040870135935060608701359250608087013567ffffffffffffffff811115613b3d57600080fd5b613b4989828a016132d7565b979a9699509497509295939492505050565b600080600060608486031215613b7057600080fd5b83359250613b8060208501613293565b915060408401356bffffffffffffffffffffffff81168114613ba157600080fd5b809150509250925092565b60005b83811015613bc7578181015183820152602001613baf565b50506000910152565b60008151808452613be8816020860160208601613bac565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613cea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613ca660e0860182613bd0565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613c38565b50909695505050505050565b600081518084526020840193506020830160005b82811015613d3e57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613d0a565b5093949350505050565b805160ff16825260006020820151613d64602085018215159052565b5060408201516101206040850152613d80610120850182613c1a565b9050606083015160608501526080830151608085015260a083015184820360a0860152613dad8282613bd0565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613ddd8282613cf6565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613e426040830186613d48565b8281036020840152613e55818587613de6565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561056c5761056c613e5f565b8082018082111561056c5761056c613e5f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008085851115613ef357600080fd5b83861115613f0057600080fd5b5050820193919092039150565b606081526000613f206060830186613d48565b60208301949094525060400152919050565b606081526000613f456060830186613d48565b8460208401528281036040840152613e558185613bd0565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613f8c6040830184613bd0565b949350505050565b60006040828403128015613fa757600080fd5b506040805190810167ffffffffffffffff81118282101715613fcb57613fcb613534565b604052825181526020928301519281019290925250919050565b606081526000613ff9606083018688613de6565b6020830194909452506040015292915050565b838152604060208201526000613ddd604083018486613de6565b60006020828403121561403857600080fd5b815161064f816133db565b6080815260006140566080830188613d48565b86602084015273ffffffffffffffffffffffffffffffffffffffff86166040840152828103606084015261408b818587613de6565b98975050505050505050565b6000602082840312156140a957600080fd5b5051919050565b6040815260006140c36040830185613bd0565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600083516140fa818460208801613bac565b919091019182525060200191905056fea264697066735822122059a6b1f52cdc641f54d3803cdfc29582ce0a1b3e39122a3f85ff43bc4a9d5c6f64736f6c634300081b0033",
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
const WalletEstimatorDeployedBin = "0x6080604052600436106101485760003560e01c80638943ec02116100c0578063b93ea7ad11610074578063ca70785011610059578063ca70785014610498578063f23a6e61146104b8578063f727ef1c146104fe5761014f565b8063b93ea7ad1461043d578063bc197c81146104505761014f565b8063975befdb116100a5578063975befdb146103c2578063aaf10f42146103d5578063ad55366b146103ea5761014f565b80638943ec02146103815780638c3f5563146103a25761014f565b80631f6a1eb9116101175780634fcf3eca116100fc5780634fcf3eca1461033857806351605d801461034b5780636ea445771461036e5761014f565b80631f6a1eb91461030557806329561426146103185761014f565b8063025b22bc14610212578063150b7a02146102255780631626ba7e146102a05780631a9b2337146102c05761014f565b3661014f57005b6004361061021057600061016b610166368361321d565b61051e565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b4929190613283565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b6102106102203660046132bc565b610572565b34801561023157600080fd5b5061026a610240366004613320565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b3480156102ac57600080fd5b5061026a6102bb36600461338f565b6105be565b3480156102cc57600080fd5b506102e06102db366004613409565b610656565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610297565b610210610313366004613426565b610661565b34801561032457600080fd5b50610210610333366004613497565b6106e3565b610210610346366004613409565b610727565b34801561035757600080fd5b506103606107e9565b604051908152602001610297565b61021061037c3660046134b0565b610818565b34801561038d57600080fd5b5061021061039c3660046134f2565b50505050565b3480156103ae57600080fd5b506103606103bd366004613497565b610878565b6103606103d0366004613426565b6108a4565b3480156103e157600080fd5b506102e0610941565b3480156103f657600080fd5b5061040a610405366004613874565b61094b565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610297565b61021061044b3660046139a2565b610985565b34801561045c57600080fd5b5061026a61046b366004613a1c565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104a457600080fd5b506103606104b3366004613874565b610a4a565b3480156104c457600080fd5b5061026a6104d3366004613ae3565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561050a57600080fd5b50610210610519366004613b5b565b610bb3565b600061056c7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c6e565b92915050565b3330146105b2576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105bb81610ccc565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610615828686610d21565b5090508061062957506000915061064f9050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061056c8261051e565b600061066d8585610e6f565b90506106818160600151826080015161128d565b60008061068f838686610d21565b91509150816106d0578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b6106da8184611375565b50505050505050565b33301461071e576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b6105bb8161161f565b333014610762576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b600061076d8261051e565b73ffffffffffffffffffffffffffffffffffffffff16036107de576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105a9565b6105bb8160006116af565b60006108137fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610853576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b600061085f8383610e6f565b9050600061086c8261176f565b905061039c8183611375565b600061056c7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c6e565b6000805a905060006108b68787610e6f565b60608101519091506108d0906108cb81610878565b61128d565b6000806108de838888610d21565b915091508161091f578287876040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b61092981846117ea565b5a6109349085613e8e565b9998505050505050505050565b6000610813305490565b600080600080600080610962898989600080611a34565b93995091975094509250905061097783611d59565b935093975093979195509350565b3330146109c0576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b60006109cb8361051e565b73ffffffffffffffffffffffffffffffffffffffff1614610a3c576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105a9565b610a4682826116af565b5050565b600080846101000151516001610a609190613ea1565b67ffffffffffffffff811115610a7857610a78613534565b604051908082528060200260200182016040528015610aa1578160200160208202803683370190505b50905060005b85610100015151811015610b13578561010001518181518110610acc57610acc613eb4565b6020026020010151828281518110610ae657610ae6613eb4565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610aa7565b5033818661010001515181518110610b2d57610b2d613eb4565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610b67868686610d21565b50905080610ba7578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105a993929190613e2f565b50600195945050505050565b333014610bee576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105a9565b610c078383836bffffffffffffffffffffffff16611d6d565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c8d929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610cd4813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610d3957610d39613eb4565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e3857610d968661176f565b9150600080610da484611dfc565b9150915042811015610dec576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105a9565b73ffffffffffffffffffffffffffffffffffffffff82161580610e24575073ffffffffffffffffffffffffffffffffffffffff821633145b15610e355760019450505050610e67565b50505b6000806000610e4b898989600080611a34565b98509295509093509150610e60905081611d59565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610ed35760006060840152610ee4565b84810135606090811c908401526014015b6007600183901c168015610f375760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610f4c57506001610f75565b83602016602003610f685750600282019186013560f01c610f75565b50600182019186013560f81c5b8067ffffffffffffffff811115610f8e57610f8e613534565b60405190808252806020026020018201604052801561101957816020015b6110066040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610fac5790505b50604086015260005b818110156112825760018085019489013560f81c90808216900361108157308760400151838151811061105757611057613eb4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526110cb565b8489013560601c60148601886040015184815181106110a2576110a2613eb4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b60028082169003611109578489013560208601886040015184815181106110f4576110f4613eb4565b60200260200101516020018197508281525050505b600480821690036111a157600385019489013560e81c89868a61112c8483613ea1565b9261113993929190613ee3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061118457611184613eb4565b60209081029190910101516040015261119d8187613ea1565b9550505b600880821690036111df578489013560208601886040015184815181106111ca576111ca613eb4565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106111ff576111ff613eb4565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061123557611235613eb4565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061126a5761126a613eb4565b602090810291909101015160c0015250600101611022565b505050505092915050565b600061129883610878565b90508181146112e4576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105a9565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b81811015611618576000846040015182815181106113a0576113a0613eb4565b602002602001015190508060a0015180156113b9575083155b15611401576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611610565b606081015180158015906114145750805a105b156114515785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a993929190613f0d565b600082608001511561148257825161147b90831561146f5783611471565b5a5b8560400151611e48565b90506114a9565b825160208401516114a69190841561149a578461149c565b5a5b8660400151611e5e565b90505b806115d35760c08301516114fc576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611610565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611566578684611531611e76565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105a993929190613f32565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016115d35760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c391015b60405180910390a1505050611618565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b600101611380565b5050505050565b80611656576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61167f7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610d16565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b600080611780836020015130611e95565b9050600061178d84611f62565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b604081015151600090815b818110156116185760008460400151828151811061181557611815613eb4565b602002602001015190508060a00151801561182e575083155b15611876576040805187815260208101849052600095507f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611a2c565b606081015180158015906118895750805a105b156118c65785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105a993929190613f0d565b60008260800151156118eb5782516118e490831561146f5783611471565b9050611906565b825160208401516119039190841561149a578461149c565b90505b806119ef5760c0830151611959576040805189815260208101869052600197507f20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb910160405180910390a1505050611a2c565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161198e578684611531611e76565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016119ef5760408051898152602081018690527f5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c391016115c3565b60408051898152602081018690527fec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd910160405180910390a15050505b6001016117f5565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611a7f575073ffffffffffffffffffffffffffffffffffffffff8916155b15611b9b578b82013560601c985060149091019089611b9b5760038201918c013560e81c60008d848e611ab28583613ea1565b92611abf93929190613ee3565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611b4a9030908590600401613f5d565b6040805180830381865afa158015611b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8a9190613f94565b9250611b968285613ea1565b935050505b82600116600103611bd557611bc38d8a838f8f87908092611bbe93929190613ee3565b6121b8565b97509750975097509750505050611d4c565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611c406001600586901c811690613ea1565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611c8b8d61176f565b9350611ca98d858e8e86908092611ca493929190613ee3565b612404565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611cf4575080518614155b8015611d04575080602001518511155b15611d48576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105a9565b5050505b9550955095509550959050565b6000611d6482612de8565b50600192915050565b611df77fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611e2a7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c6e565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611f055746611f08565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611f73836101000151612df3565b835190915060ff16611fe6576000611f8e8460400151612e76565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016117cb565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016120765760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016120e65760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001612058565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016121565760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001612058565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105a9565b600080600080600061221a604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156123915760038201916000908b013560e81c6122628482613ea1565b9150600090508a8214612276576000612278565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84146122a957856122ab565b8f5b90506122cb818e8e889087926122c393929190613ee3565b600186611a34565b50929d50909b50995097508a8a1015612328578c8c869085926122f093929190613ee3565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105a99493929190613fe5565b829450888e600001510361233b5760008e525b83881061237e576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105a9565b505060c084018790529150849050612242565b8a51158015906123a557508a602001518511155b156123e9576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105a9565b6123f28d61176f565b93505050509550955095509550959050565b60008060005b83811015612dde57600181019085013560f881901c9060fc1c8061252757600f821660008190036124425750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156124d7573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006124fa828960ff16612ef8565b90508b6125075780612516565b60008c81526020829052604090205b9b505050505050505050505061240a565b6001810361258b57600f821660008190036125495750600183019287013560f81c5b601484019388013560601c60006125638260ff8516612ef8565b905086612570578061257f565b60008781526020829052604090205b9650505050505061240a565b60028103612772576003821660008190036125ad5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261266b93929190613ee3565b6040518463ffffffff1660e01b81526004016126899392919061400c565b602060405180830381865afa1580156126a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126ca9190614026565b7fffffffff00000000000000000000000000000000000000000000000000000000161461272d578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105a9959493929190614043565b8097508460ff168a0199506000612747858760ff16612ef8565b9050896127545780612763565b60008a81526020829052604090205b9950505050505050505061240a565b600381036127a65760208301928701358461278d578061279c565b60008581526020829052604090205b945050505061240a565b6004810361284c57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c1685830180965081925050506000818601905060008061281b8e8e8e8e8c908892611ca493929190613ee3565b91509150829750818a01995061283b898260009182526020526040902090565b98508297505050505050505061240a565b60068103612916576003600283901c1660008190036128725750600183019287013560f81c5b60038316600081900361288c5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806128ca8f8f8f8f8d908892611ca493929190613ee3565b915091508298508482106128dd57998501995b60006128ea828789612f5f565b90508a6128f75780612906565b60008b81526020829052604090205b9a5050505050505050505061240a565b60058103612983576020830192870135888103612951577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061295c82612fc3565b9050856129695780612978565b60008681526020829052604090205b95505050505061240a565b60078103612a8957600f821660008190036129a55750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016124b5565b60088103612add5760208301928701356000612aa58b82613017565b9050808203612ad2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061256382613092565b60098103612c4357600382166000819003612aff5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c908792612b9993929190613ee3565b6040518463ffffffff1660e01b8152600401612bb793929190613e2f565b602060405180830381865afa158015612bd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bf89190614097565b90508197508460ff168a0199506000612c15858760ff16846130cd565b905089612c225780612c31565b60008a81526020829052604090205b9950829850505050505050505061240a565b600a8103612da957600382166000819003612c655750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612cfe93929190613ee3565b6040518463ffffffff1660e01b8152600401612d1c9392919061400c565b602060405180830381865afa158015612d39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d5d9190614097565b90508198508560ff168b019a506000612d7a868860ff16846130cd565b90508a612d875780612d96565b60008b81526020829052604090205b9a5082995050505050505050505061240a565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105a9565b5094509492505050565b600061056c8261313b565b6000606060005b8351811015612e675781848281518110612e1657612e16613eb4565b6020026020010151604051602001612e2f9291906140b0565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612dfa565b50805160209091012092915050565b6000606060005b8351811015612e67576000612eaa858381518110612e9d57612e9d613eb4565b602002602001015161316e565b90508281604051602001612ebf9291906140e8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612e7d565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611f44565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080613028846020015184611e95565b9050600061303585611f62565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612ffa565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612fa4565b6000811580159061056c5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612ffa98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561327c577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146132b757600080fd5b919050565b6000602082840312156132ce57600080fd5b61064f82613293565b60008083601f8401126132e957600080fd5b50813567ffffffffffffffff81111561330157600080fd5b60208301915083602082850101111561331957600080fd5b9250929050565b60008060008060006080868803121561333857600080fd5b61334186613293565b945061334f60208701613293565b935060408601359250606086013567ffffffffffffffff81111561337257600080fd5b61337e888289016132d7565b969995985093965092949392505050565b6000806000604084860312156133a457600080fd5b83359250602084013567ffffffffffffffff8111156133c257600080fd5b6133ce868287016132d7565b9497909650939450505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105bb57600080fd5b60006020828403121561341b57600080fd5b813561064f816133db565b6000806000806040858703121561343c57600080fd5b843567ffffffffffffffff81111561345357600080fd5b61345f878288016132d7565b909550935050602085013567ffffffffffffffff81111561347f57600080fd5b61348b878288016132d7565b95989497509550505050565b6000602082840312156134a957600080fd5b5035919050565b600080602083850312156134c357600080fd5b823567ffffffffffffffff8111156134da57600080fd5b6134e6858286016132d7565b90969095509350505050565b6000806000806060858703121561350857600080fd5b61351185613293565b935060208501359250604085013567ffffffffffffffff81111561347f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561358657613586613534565b60405290565b604051610120810167ffffffffffffffff8111828210171561358657613586613534565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156135f7576135f7613534565b604052919050565b803560ff811681146132b757600080fd5b803580151581146132b757600080fd5b600067ffffffffffffffff82111561363a5761363a613534565b5060051b60200190565b600082601f83011261365557600080fd5b813567ffffffffffffffff81111561366f5761366f613534565b6136a060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016135b0565b8181528460208386010111156136b557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126136e357600080fd5b81356136f66136f182613620565b6135b0565b8082825260208201915060208360051b86010192508583111561371857600080fd5b602085015b8381101561380557803567ffffffffffffffff81111561373c57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561377057600080fd5b613778613563565b61378460208301613293565b815260408201356020820152606082013567ffffffffffffffff8111156137aa57600080fd5b6137b98a602083860101613644565b604083015250608082013560608201526137d560a08301613610565b60808201526137e660c08301613610565b60a082015260e0919091013560c082015283526020928301920161371d565b5095945050505050565b600082601f83011261382057600080fd5b813561382e6136f182613620565b8082825260208201915060208360051b86010192508583111561385057600080fd5b602085015b838110156138055761386681613293565b835260209283019201613855565b60008060006040848603121561388957600080fd5b833567ffffffffffffffff8111156138a057600080fd5b840161012081870312156138b357600080fd5b6138bb61358c565b6138c4826135ff565b81526138d260208301613610565b6020820152604082013567ffffffffffffffff8111156138f157600080fd5b6138fd888285016136d2565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561393157600080fd5b61393d88828501613644565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561397257600080fd5b61397e8882850161380f565b61010083015250935050602084013567ffffffffffffffff8111156133c257600080fd5b600080604083850312156139b557600080fd5b82356139c0816133db565b91506139ce60208401613293565b90509250929050565b60008083601f8401126139e957600080fd5b50813567ffffffffffffffff811115613a0157600080fd5b6020830191508360208260051b850101111561331957600080fd5b60008060008060008060008060a0898b031215613a3857600080fd5b613a4189613293565b9750613a4f60208a01613293565b9650604089013567ffffffffffffffff811115613a6b57600080fd5b613a778b828c016139d7565b909750955050606089013567ffffffffffffffff811115613a9757600080fd5b613aa38b828c016139d7565b909550935050608089013567ffffffffffffffff811115613ac357600080fd5b613acf8b828c016132d7565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613afc57600080fd5b613b0587613293565b9550613b1360208801613293565b94506040870135935060608701359250608087013567ffffffffffffffff811115613b3d57600080fd5b613b4989828a016132d7565b979a9699509497509295939492505050565b600080600060608486031215613b7057600080fd5b83359250613b8060208501613293565b915060408401356bffffffffffffffffffffffff81168114613ba157600080fd5b809150509250925092565b60005b83811015613bc7578181015183820152602001613baf565b50506000910152565b60008151808452613be8816020860160208601613bac565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613cea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613ca660e0860182613bd0565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613c38565b50909695505050505050565b600081518084526020840193506020830160005b82811015613d3e57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613d0a565b5093949350505050565b805160ff16825260006020820151613d64602085018215159052565b5060408201516101206040850152613d80610120850182613c1a565b9050606083015160608501526080830151608085015260a083015184820360a0860152613dad8282613bd0565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613ddd8282613cf6565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613e426040830186613d48565b8281036020840152613e55818587613de6565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561056c5761056c613e5f565b8082018082111561056c5761056c613e5f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008085851115613ef357600080fd5b83861115613f0057600080fd5b5050820193919092039150565b606081526000613f206060830186613d48565b60208301949094525060400152919050565b606081526000613f456060830186613d48565b8460208401528281036040840152613e558185613bd0565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613f8c6040830184613bd0565b949350505050565b60006040828403128015613fa757600080fd5b506040805190810167ffffffffffffffff81118282101715613fcb57613fcb613534565b604052825181526020928301519281019290925250919050565b606081526000613ff9606083018688613de6565b6020830194909452506040015292915050565b838152604060208201526000613ddd604083018486613de6565b60006020828403121561403857600080fd5b815161064f816133db565b6080815260006140566080830188613d48565b86602084015273ffffffffffffffffffffffffffffffffffffffff86166040840152828103606084015261408b818587613de6565b98975050505050505050565b6000602082840312156140a957600080fd5b5051919050565b6040815260006140c36040830185613bd0565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b600083516140fa818460208801613bac565b919091019182525060200191905056fea264697066735822122059a6b1f52cdc641f54d3803cdfc29582ce0a1b3e39122a3f85ff43bc4a9d5c6f64736f6c634300081b0033"

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

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCaller) IsValidSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletEstimator.contract.Call(opts, &out, "isValidSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletEstimator.Contract.IsValidSapientSignature(&_WalletEstimator.CallOpts, _payload, _signature)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletEstimator *WalletEstimatorCallerSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletEstimator.Contract.IsValidSapientSignature(&_WalletEstimator.CallOpts, _payload, _signature)
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
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletEstimatorCallAbortedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallAbortedIterator{contract: _WalletEstimator.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
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

// ParseCallAborted is a log parse operation binding the contract event 0x5b5cb72c79981de49f1f950d4d8d62397e2fc2b772e1b788a640025075ab47c3.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index)
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
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletEstimatorCallFailedIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallFailedIterator{contract: _WalletEstimator.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
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

// ParseCallFailed is a log parse operation binding the contract event 0x20832642214d5218c6428e71d8d2ddd9ad15a81ad2be8154d8c2e3ab08293fcb.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index)
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

// WalletEstimatorCallSuccessIterator is returned from FilterCallSuccess and is used to iterate over the raw logs and unpacked data for CallSuccess events raised by the WalletEstimator contract.
type WalletEstimatorCallSuccessIterator struct {
	Event *WalletEstimatorCallSuccess // Event containing the contract specifics and raw log

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
func (it *WalletEstimatorCallSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEstimatorCallSuccess)
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
		it.Event = new(WalletEstimatorCallSuccess)
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
func (it *WalletEstimatorCallSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEstimatorCallSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEstimatorCallSuccess represents a CallSuccess event raised by the WalletEstimator contract.
type WalletEstimatorCallSuccess struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSuccess is a free log retrieval operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) FilterCallSuccess(opts *bind.FilterOpts) (*WalletEstimatorCallSuccessIterator, error) {

	logs, sub, err := _WalletEstimator.contract.FilterLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return &WalletEstimatorCallSuccessIterator{contract: _WalletEstimator.contract, event: "CallSuccess", logs: logs, sub: sub}, nil
}

// WatchCallSuccess is a free log subscription operation binding the contract event 0xec670aed5ee1e72eb3eb601271be4b3f312e71f17eebdf10c1a0ab5a3af30ffd.
//
// Solidity: event CallSuccess(bytes32 _opHash, uint256 _index)
func (_WalletEstimator *WalletEstimatorFilterer) WatchCallSuccess(opts *bind.WatchOpts, sink chan<- *WalletEstimatorCallSuccess) (event.Subscription, error) {

	logs, sub, err := _WalletEstimator.contract.WatchLogs(opts, "CallSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEstimatorCallSuccess)
				if err := _WalletEstimator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
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
func (_WalletEstimator *WalletEstimatorFilterer) ParseCallSuccess(log types.Log) (*WalletEstimatorCallSuccess, error) {
	event := new(WalletEstimatorCallSuccess)
	if err := _WalletEstimator.contract.UnpackLog(event, "CallSuccess", log); err != nil {
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
