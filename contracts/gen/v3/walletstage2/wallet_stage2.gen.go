// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletstage2

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

// WalletStage2MetaData contains all meta data concerning the WalletStage2 contract.
var WalletStage2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_subdigest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidNestedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Aborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Failed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Skipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Success\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50613e258061001f6000396000f3fe60806040526004361061012d5760003560e01c80638943ec02116100a5578063b93ea7ad11610074578063ca70785011610059578063ca707850146104ab578063f23a6e61146104cb578063f727ef1c1461051157610134565b8063b93ea7ad14610443578063bc197c811461046357610134565b80638943ec021461039a5780638c3f5563146103bb578063aaf10f42146103db578063ad55366b146103f057610134565b80631f6a1eb9116100fc5780634fcf3eca116100e15780634fcf3eca1461033757806351605d80146103575780636ea445771461037a57610134565b80631f6a1eb9146102f7578063295614261461031757610134565b8063025b22bc146101f7578063150b7a02146102175780631626ba7e146102925780631a9b2337146102b257610134565b3661013457005b600436106101f557600061015061014b3683612f1d565b610531565b905073ffffffffffffffffffffffffffffffffffffffff8116156101f3576000808273ffffffffffffffffffffffffffffffffffffffff16600036604051610199929190612f83565b600060405180830381855af49150503d80600081146101d4576040519150601f19603f3d011682016040523d82523d6000602084013e6101d9565b606091505b5091509150816101eb57805160208201fd5b805160208201f35b505b005b34801561020357600080fd5b506101f5610212366004612fbc565b610585565b34801561022357600080fd5b5061025c610232366004613020565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020015b60405180910390f35b34801561029e57600080fd5b5061025c6102ad36600461308f565b6105d1565b3480156102be57600080fd5b506102d26102cd366004613109565b610669565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610289565b34801561030357600080fd5b506101f5610312366004613126565b610674565b34801561032357600080fd5b506101f5610332366004613197565b6106f6565b34801561034357600080fd5b506101f5610352366004613109565b61073a565b34801561036357600080fd5b5061036c6107fc565b604051908152602001610289565b34801561038657600080fd5b506101f56103953660046131b0565b61082b565b3480156103a657600080fd5b506101f56103b53660046131f2565b50505050565b3480156103c757600080fd5b5061036c6103d6366004613197565b61088b565b3480156103e757600080fd5b506102d26108b7565b3480156103fc57600080fd5b5061041061040b366004613574565b6108c1565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610289565b34801561044f57600080fd5b506101f561045e3660046136a2565b6108fb565b34801561046f57600080fd5b5061025c61047e36600461371c565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104b757600080fd5b5061036c6104c6366004613574565b6109c0565b3480156104d757600080fd5b5061025c6104e63660046137e3565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561051d57600080fd5b506101f561052c36600461385b565b610b29565b600061057f7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610be4565b92915050565b3330146105c5576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105ce81610c42565b50565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905281610628828686610c97565b5090508061063c5750600091506106629050565b507f20c13b0b000000000000000000000000000000000000000000000000000000009150505b9392505050565b600061057f82610531565b60006106808585610de5565b905061069481606001518260800151611203565b6000806106a2838686610c97565b91509150816106e3578285856040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105bc93929190613b2f565b6106ed81846112eb565b50505050505050565b333014610731576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105bc565b6105ce8161157d565b333014610775576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105bc565b600061078082610531565b73ffffffffffffffffffffffffffffffffffffffff16036107f1576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105bc565b6105ce81600061160d565b60006108267fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b333014610866576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105bc565b60006108728383610de5565b9050600061087f826116cd565b90506103b581836112eb565b600061057f7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610be4565b6000610826305490565b6000806000806000806108d8898989600080611748565b9399509197509450925090506108ed83611a6d565b935093975093979195509350565b333014610936576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105bc565b600061094183610531565b73ffffffffffffffffffffffffffffffffffffffff16146109b2576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105bc565b6109bc828261160d565b5050565b6000808461010001515160016109d69190613b5f565b67ffffffffffffffff8111156109ee576109ee613234565b604051908082528060200260200182016040528015610a17578160200160208202803683370190505b50905060005b85610100015151811015610a89578561010001518181518110610a4257610a42613b99565b6020026020010151828281518110610a5c57610a5c613b99565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610a1d565b5033818661010001515181518110610aa357610aa3613b99565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610add868686610c97565b50905080610b1d578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105bc93929190613b2f565b50600195945050505050565b333014610b64576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105bc565b610b7d8383836bffffffffffffffffffffffff16611a78565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c03929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c4a813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610caf57610caf613b99565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610dae57610d0c866116cd565b9150600080610d1a84611b07565b9150915042811015610d62576040517f9fa4fe5400000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105bc565b73ffffffffffffffffffffffffffffffffffffffff82161580610d9a575073ffffffffffffffffffffffffffffffffffffffff821633145b15610dab5760019450505050610ddd565b50505b6000806000610dc1898989600080611748565b98509295509093509150610dd6905081611a6d565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610e495760006060840152610e5a565b84810135606090811c908401526014015b6007600183901c168015610ead5760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610ec257506001610eeb565b83602016602003610ede5750600282019186013560f01c610eeb565b50600182019186013560f81c5b8067ffffffffffffffff811115610f0457610f04613234565b604051908082528060200260200182016040528015610f8f57816020015b610f7c6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b815260200190600190039081610f225790505b50604086015260005b818110156111f85760018085019489013560f81c908082169003610ff7573087604001518381518110610fcd57610fcd613b99565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff9091169052611041565b8489013560601c601486018860400151848151811061101857611018613b99565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b6002808216900361107f5784890135602086018860400151848151811061106a5761106a613b99565b60200260200101516020018197508281525050505b6004808216900361111757600385019489013560e81c89868a6110a28483613b5f565b926110af93929190613bc8565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106110fa576110fa613b99565b6020908102919091010151604001526111138187613b5f565b9550505b600880821690036111555784890135602086018860400151848151811061114057611140613b99565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061117557611175613b99565b602002602001015160800190151590811515815250508060201660ff16602014876040015183815181106111ab576111ab613b99565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106111e0576111e0613b99565b602090810291909101015160c0015250600101610f98565b505050505092915050565b600061120e8361088b565b905081811461125a576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105bc565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b818110156115765760008460400151828151811061131657611316613b99565b602002602001015190508060a00151801561132f575083155b15611377576040805187815260208101849052600095507fa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7910160405180910390a15061156e565b6060810151801580159061138a5750805a105b156113c75785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105bc93929190613bf2565b60008260800151156113ed576113e68360000151838560400151611b53565b9050611408565b61140583600001518460200151848660400151611b69565b90505b806115315760c083015161145b576040805189815260208101869052600197507fe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643910160405180910390a150505061156e565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114c5578684611490611b81565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105bc93929190613c17565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016115315760408051898152602081018690527f6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8910160405180910390a1505050611576565b60408051898152602081018690527f2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3910160405180910390a15050505b6001016112f6565b5050505050565b806115b4576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6115dd7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610c8c565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b6000806116de836020015130611ba0565b905060006116eb84611c6d565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611793575073ffffffffffffffffffffffffffffffffffffffff8916155b156118af578b82013560601c9850601490910190896118af5760038201918c013560e81c60008d848e6117c68583613b5f565b926117d393929190613bc8565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc8915061185e9030908590600401613c42565b6040805180830381865afa15801561187a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061189e9190613c79565b92506118aa8285613b5f565b935050505b826001166001036118e9576118d78d8a838f8f879080926118d293929190613bc8565b611ec3565b97509750975097509750505050611a60565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c168382019096509250600090506119546001600586901c811690613b5f565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c169950909201915061199f8d6116cd565b93506119bd8d858e8e869080926119b893929190613bc8565b61210f565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611a08575080518614155b8015611a18575080602001518511155b15611a5c576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105bc565b5050505b9550955095509550959050565b600061057f82612af3565b611b027fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008080611b357fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610be4565b606081901c956bffffffffffffffffffffffff909116945092505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611c105746611c13565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080611c7e836101000151612b26565b835190915060ff16611cf1576000611c998460400151612ba9565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c001611729565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611d815760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611df15760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001611d63565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd01611e615760e0830151604080517f402e923b91e918306019e73f589362164a6a059499a504699c66feabbb3e262460208201529081019190915260608101829052608001611d63565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f556e737570706f72746564206b696e640000000000000000000000000000000060448201526064016105bc565b6000806000806000611f25604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b8882101561209c5760038201916000908b013560e81c611f6d8482613b5f565b9150600090508a8214611f81576000611f83565b8d5b905060007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8414611fb45785611fb6565b8f5b9050611fd6818e8e88908792611fce93929190613bc8565b600186611748565b50929d50909b50995097508a8a1015612033578c8c86908592611ffb93929190613bc8565b8c8c6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105bc9493929190613cca565b829450888e60000151036120465760008e525b838810612089576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101899052602481018590526044016105bc565b505060c084018790529150849050611f4d565b8a51158015906120b057508a602001518511155b156120f4576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105bc565b6120fd8d6116cd565b93505050509550955095509550959050565b60008060005b83811015612ae957600181019085013560f881901c9060fc1c8061223257600f8216600081900361214d5750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156121e2573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b506000612205828960ff16612c2b565b90508b6122125780612221565b60008c81526020829052604090205b9b5050505050505050505050612115565b6001810361229657600f821660008190036122545750600183019287013560f81c5b601484019388013560601c600061226e8260ff8516612c2b565b90508661227b578061228a565b60008781526020829052604090205b96505050505050612115565b6002810361247d576003821660008190036122b85750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261237693929190613bc8565b6040518463ffffffff1660e01b815260040161239493929190613cf1565b602060405180830381865afa1580156123b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123d59190613d0b565b7fffffffff000000000000000000000000000000000000000000000000000000001614612438578d8d858e8e6040517ff734863a0000000000000000000000000000000000000000000000000000000081526004016105bc959493929190613d28565b8097508460ff168a0199506000612452858760ff16612c2b565b90508961245f578061246e565b60008a81526020829052604090205b99505050505050505050612115565b600381036124b15760208301928701358461249857806124a7565b60008581526020829052604090205b9450505050612115565b6004810361255757600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806125268e8e8e8e8c9088926119b893929190613bc8565b91509150829750818a019950612546898260009182526020526040902090565b985082975050505050505050612115565b60068103612621576003600283901c16600081900361257d5750600183019287013560f81c5b6003831660008190036125975750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806125d58f8f8f8f8d9088926119b893929190613bc8565b915091508298508482106125e857998501995b60006125f5828789612c92565b90508a6126025780612611565b60008b81526020829052604090205b9a50505050505050505050612115565b6005810361268e57602083019287013588810361265c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061266782612cf6565b9050856126745780612683565b60008681526020829052604090205b955050505050612115565b6007810361279457600f821660008190036126b05750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a0016121c0565b600881036127e857602083019287013560006127b08b82612d4a565b90508082036127dd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061226e82612dc5565b6009810361294e5760038216600081900361280a5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff1663ca7078508f8e8e8c9087926128a493929190613bc8565b6040518463ffffffff1660e01b81526004016128c293929190613b2f565b602060405180830381865afa1580156128df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129039190613d7c565b90508197508460ff168a0199506000612920858760ff1684612e00565b90508961292d578061293c565b60008a81526020829052604090205b99508298505050505050505050612115565b600a8103612ab4576003821660008190036129705750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663957d2b238f8f8f8d908792612a0993929190613bc8565b6040518463ffffffff1660e01b8152600401612a2793929190613cf1565b602060405180830381865afa158015612a44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a689190613d7c565b90508198508560ff168b019a506000612a85868860ff1684612e00565b90508a612a925780612aa1565b60008b81526020829052604090205b9a50829950505050505050505050612115565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105bc565b5094509492505050565b6000811580159061057f5750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b6000606060005b8351811015612b9a5781848281518110612b4957612b49613b99565b6020026020010151604051602001612b62929190613d95565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529150600101612b2d565b50805160209091012092915050565b6000606060005b8351811015612b9a576000612bdd858381518110612bd057612bd0613b99565b6020026020010151612e6e565b90508281604051602001612bf2929190613dcd565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612bb0565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611c4f565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612d5b846020015184611ba0565b90506000612d6885611c6d565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612d2d565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612cd7565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612d2d98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff000000000000000000000000000000000000000000000000000000008116906004841015612f7c577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612fb757600080fd5b919050565b600060208284031215612fce57600080fd5b61066282612f93565b60008083601f840112612fe957600080fd5b50813567ffffffffffffffff81111561300157600080fd5b60208301915083602082850101111561301957600080fd5b9250929050565b60008060008060006080868803121561303857600080fd5b61304186612f93565b945061304f60208701612f93565b935060408601359250606086013567ffffffffffffffff81111561307257600080fd5b61307e88828901612fd7565b969995985093965092949392505050565b6000806000604084860312156130a457600080fd5b83359250602084013567ffffffffffffffff8111156130c257600080fd5b6130ce86828701612fd7565b9497909650939450505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105ce57600080fd5b60006020828403121561311b57600080fd5b8135610662816130db565b6000806000806040858703121561313c57600080fd5b843567ffffffffffffffff81111561315357600080fd5b61315f87828801612fd7565b909550935050602085013567ffffffffffffffff81111561317f57600080fd5b61318b87828801612fd7565b95989497509550505050565b6000602082840312156131a957600080fd5b5035919050565b600080602083850312156131c357600080fd5b823567ffffffffffffffff8111156131da57600080fd5b6131e685828601612fd7565b90969095509350505050565b6000806000806060858703121561320857600080fd5b61321185612f93565b935060208501359250604085013567ffffffffffffffff81111561317f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561328657613286613234565b60405290565b604051610120810167ffffffffffffffff8111828210171561328657613286613234565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156132f7576132f7613234565b604052919050565b803560ff81168114612fb757600080fd5b80358015158114612fb757600080fd5b600067ffffffffffffffff82111561333a5761333a613234565b5060051b60200190565b600082601f83011261335557600080fd5b813567ffffffffffffffff81111561336f5761336f613234565b6133a060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016132b0565b8181528460208386010111156133b557600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126133e357600080fd5b81356133f66133f182613320565b6132b0565b8082825260208201915060208360051b86010192508583111561341857600080fd5b602085015b8381101561350557803567ffffffffffffffff81111561343c57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561347057600080fd5b613478613263565b61348460208301612f93565b815260408201356020820152606082013567ffffffffffffffff8111156134aa57600080fd5b6134b98a602083860101613344565b604083015250608082013560608201526134d560a08301613310565b60808201526134e660c08301613310565b60a082015260e0919091013560c082015283526020928301920161341d565b5095945050505050565b600082601f83011261352057600080fd5b813561352e6133f182613320565b8082825260208201915060208360051b86010192508583111561355057600080fd5b602085015b838110156135055761356681612f93565b835260209283019201613555565b60008060006040848603121561358957600080fd5b833567ffffffffffffffff8111156135a057600080fd5b840161012081870312156135b357600080fd5b6135bb61328c565b6135c4826132ff565b81526135d260208301613310565b6020820152604082013567ffffffffffffffff8111156135f157600080fd5b6135fd888285016133d2565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561363157600080fd5b61363d88828501613344565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561367257600080fd5b61367e8882850161350f565b61010083015250935050602084013567ffffffffffffffff8111156130c257600080fd5b600080604083850312156136b557600080fd5b82356136c0816130db565b91506136ce60208401612f93565b90509250929050565b60008083601f8401126136e957600080fd5b50813567ffffffffffffffff81111561370157600080fd5b6020830191508360208260051b850101111561301957600080fd5b60008060008060008060008060a0898b03121561373857600080fd5b61374189612f93565b975061374f60208a01612f93565b9650604089013567ffffffffffffffff81111561376b57600080fd5b6137778b828c016136d7565b909750955050606089013567ffffffffffffffff81111561379757600080fd5b6137a38b828c016136d7565b909550935050608089013567ffffffffffffffff8111156137c357600080fd5b6137cf8b828c01612fd7565b999c989b5096995094979396929594505050565b60008060008060008060a087890312156137fc57600080fd5b61380587612f93565b955061381360208801612f93565b94506040870135935060608701359250608087013567ffffffffffffffff81111561383d57600080fd5b61384989828a01612fd7565b979a9699509497509295939492505050565b60008060006060848603121561387057600080fd5b8335925061388060208501612f93565b915060408401356bffffffffffffffffffffffff811681146138a157600080fd5b809150509250925092565b60005b838110156138c75781810151838201526020016138af565b50506000910152565b600081518084526138e88160208601602086016138ac565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b838110156139ea577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e060408601526139a660e08601826138d0565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613938565b50909695505050505050565b600081518084526020840193506020830160005b82811015613a3e57815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613a0a565b5093949350505050565b805160ff16825260006020820151613a64602085018215159052565b5060408201516101206040850152613a8061012085018261391a565b9050606083015160608501526080830151608085015260a083015184820360a0860152613aad82826138d0565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613add82826139f6565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613b426040830186613a48565b8281036020840152613b55818587613ae6565b9695505050505050565b8082018082111561057f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008085851115613bd857600080fd5b83861115613be557600080fd5b5050820193919092039150565b606081526000613c056060830186613a48565b60208301949094525060400152919050565b606081526000613c2a6060830186613a48565b8460208401528281036040840152613b5581856138d0565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613c7160408301846138d0565b949350505050565b60006040828403128015613c8c57600080fd5b506040805190810167ffffffffffffffff81118282101715613cb057613cb0613234565b604052825181526020928301519281019290925250919050565b606081526000613cde606083018688613ae6565b6020830194909452506040015292915050565b838152604060208201526000613add604083018486613ae6565b600060208284031215613d1d57600080fd5b8151610662816130db565b608081526000613d3b6080830188613a48565b86602084015273ffffffffffffffffffffffffffffffffffffffff861660408401528281036060840152613d70818587613ae6565b98975050505050505050565b600060208284031215613d8e57600080fd5b5051919050565b604081526000613da860408301856138d0565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b60008351613ddf8184602088016138ac565b919091019182525060200191905056fea2646970667358221220b1632644d2e6889a38551818e2a1d1663ff4967c7a1d37e0acf447d231481e2e64736f6c634300081b0033",
}

// WalletStage2ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage2MetaData.ABI instead.
var WalletStage2ABI = WalletStage2MetaData.ABI

// WalletStage2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage2MetaData.Bin instead.
var WalletStage2Bin = WalletStage2MetaData.Bin

// DeployWalletStage2 deploys a new Ethereum contract, binding an instance of WalletStage2 to it.
func DeployWalletStage2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletStage2, error) {
	parsed, err := WalletStage2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletStage2{WalletStage2Caller: WalletStage2Caller{contract: contract}, WalletStage2Transactor: WalletStage2Transactor{contract: contract}, WalletStage2Filterer: WalletStage2Filterer{contract: contract}}, nil
}

// WalletStage2 is an auto generated Go binding around an Ethereum contract.
type WalletStage2 struct {
	WalletStage2Caller     // Read-only binding to the contract
	WalletStage2Transactor // Write-only binding to the contract
	WalletStage2Filterer   // Log filterer for contract events
}

// WalletStage2Caller is an auto generated read-only Go binding around an Ethereum contract.
type WalletStage2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletStage2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletStage2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletStage2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletStage2Session struct {
	Contract     *WalletStage2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletStage2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletStage2CallerSession struct {
	Contract *WalletStage2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WalletStage2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletStage2TransactorSession struct {
	Contract     *WalletStage2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WalletStage2Raw is an auto generated low-level Go binding around an Ethereum contract.
type WalletStage2Raw struct {
	Contract *WalletStage2 // Generic contract binding to access the raw methods on
}

// WalletStage2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletStage2CallerRaw struct {
	Contract *WalletStage2Caller // Generic read-only contract binding to access the raw methods on
}

// WalletStage2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletStage2TransactorRaw struct {
	Contract *WalletStage2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletStage2 creates a new instance of WalletStage2, bound to a specific deployed contract.
func NewWalletStage2(address common.Address, backend bind.ContractBackend) (*WalletStage2, error) {
	contract, err := bindWalletStage2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletStage2{WalletStage2Caller: WalletStage2Caller{contract: contract}, WalletStage2Transactor: WalletStage2Transactor{contract: contract}, WalletStage2Filterer: WalletStage2Filterer{contract: contract}}, nil
}

// NewWalletStage2Caller creates a new read-only instance of WalletStage2, bound to a specific deployed contract.
func NewWalletStage2Caller(address common.Address, caller bind.ContractCaller) (*WalletStage2Caller, error) {
	contract, err := bindWalletStage2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage2Caller{contract: contract}, nil
}

// NewWalletStage2Transactor creates a new write-only instance of WalletStage2, bound to a specific deployed contract.
func NewWalletStage2Transactor(address common.Address, transactor bind.ContractTransactor) (*WalletStage2Transactor, error) {
	contract, err := bindWalletStage2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletStage2Transactor{contract: contract}, nil
}

// NewWalletStage2Filterer creates a new log filterer instance of WalletStage2, bound to a specific deployed contract.
func NewWalletStage2Filterer(address common.Address, filterer bind.ContractFilterer) (*WalletStage2Filterer, error) {
	contract, err := bindWalletStage2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletStage2Filterer{contract: contract}, nil
}

// bindWalletStage2 binds a generic wrapper to an already deployed contract.
func bindWalletStage2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletStage2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage2 *WalletStage2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage2.Contract.WalletStage2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage2 *WalletStage2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2.Contract.WalletStage2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage2 *WalletStage2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage2.Contract.WalletStage2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletStage2 *WalletStage2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletStage2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletStage2 *WalletStage2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletStage2 *WalletStage2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletStage2.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2 *WalletStage2Caller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2 *WalletStage2Session) GetImplementation() (common.Address, error) {
	return _WalletStage2.Contract.GetImplementation(&_WalletStage2.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_WalletStage2 *WalletStage2CallerSession) GetImplementation() (common.Address, error) {
	return _WalletStage2.Contract.GetImplementation(&_WalletStage2.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2 *WalletStage2Caller) ImageHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2 *WalletStage2Session) ImageHash() ([32]byte, error) {
	return _WalletStage2.Contract.ImageHash(&_WalletStage2.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(bytes32)
func (_WalletStage2 *WalletStage2CallerSession) ImageHash() ([32]byte, error) {
	return _WalletStage2.Contract.ImageHash(&_WalletStage2.CallOpts)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2Caller) IsValidSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "isValidSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2Session) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2.Contract.IsValidSapientSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// IsValidSapientSignature is a free data retrieval call binding the contract method 0xca707850.
//
// Solidity: function isValidSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2CallerSession) IsValidSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2.Contract.IsValidSapientSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2 *WalletStage2Caller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2 *WalletStage2Session) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage2.Contract.IsValidSignature(&_WalletStage2.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_WalletStage2 *WalletStage2CallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _WalletStage2.Contract.IsValidSignature(&_WalletStage2.CallOpts, _hash, _signature)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Caller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Session) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC1155BatchReceived(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2CallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC1155BatchReceived(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Caller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Session) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC1155Received(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2CallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC1155Received(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC721Received(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_WalletStage2 *WalletStage2CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _WalletStage2.Contract.OnERC721Received(&_WalletStage2.CallOpts, arg0, arg1, arg2, arg3)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2 *WalletStage2Caller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2 *WalletStage2Session) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage2.Contract.ReadHook(&_WalletStage2.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage2 *WalletStage2CallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage2.Contract.ReadHook(&_WalletStage2.CallOpts, signature)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2 *WalletStage2Caller) ReadNonce(opts *bind.CallOpts, _space *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "readNonce", _space)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2 *WalletStage2Session) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage2.Contract.ReadNonce(&_WalletStage2.CallOpts, _space)
}

// ReadNonce is a free data retrieval call binding the contract method 0x8c3f5563.
//
// Solidity: function readNonce(uint256 _space) view returns(uint256)
func (_WalletStage2 *WalletStage2CallerSession) ReadNonce(_space *big.Int) (*big.Int, error) {
	return _WalletStage2.Contract.ReadNonce(&_WalletStage2.CallOpts, _space)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage2 *WalletStage2Caller) RecoverPartialSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "recoverPartialSignature", _payload, _signature)

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
func (_WalletStage2 *WalletStage2Session) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage2.Contract.RecoverPartialSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// RecoverPartialSignature is a free data retrieval call binding the contract method 0xad55366b.
//
// Solidity: function recoverPartialSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(uint256 threshold, uint256 weight, bool isValidImage, bytes32 imageHash, uint256 checkpoint, bytes32 opHash)
func (_WalletStage2 *WalletStage2CallerSession) RecoverPartialSignature(_payload PayloadDecoded, _signature []byte) (struct {
	Threshold    *big.Int
	Weight       *big.Int
	IsValidImage bool
	ImageHash    [32]byte
	Checkpoint   *big.Int
	OpHash       [32]byte
}, error) {
	return _WalletStage2.Contract.RecoverPartialSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) returns()
func (_WalletStage2 *WalletStage2Transactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) returns()
func (_WalletStage2 *WalletStage2Session) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.AddHook(&_WalletStage2.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) returns()
func (_WalletStage2 *WalletStage2TransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.AddHook(&_WalletStage2.TransactOpts, signature, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletStage2 *WalletStage2Transactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletStage2 *WalletStage2Session) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Execute(&_WalletStage2.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) returns()
func (_WalletStage2 *WalletStage2TransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Execute(&_WalletStage2.TransactOpts, _payload, _signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) returns()
func (_WalletStage2 *WalletStage2Transactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) returns()
func (_WalletStage2 *WalletStage2Session) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.RemoveHook(&_WalletStage2.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) returns()
func (_WalletStage2 *WalletStage2TransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.RemoveHook(&_WalletStage2.TransactOpts, signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletStage2 *WalletStage2Transactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletStage2 *WalletStage2Session) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.SelfExecute(&_WalletStage2.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) returns()
func (_WalletStage2 *WalletStage2TransactorSession) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.SelfExecute(&_WalletStage2.TransactOpts, _payload)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2 *WalletStage2Transactor) SetStaticSignature(opts *bind.TransactOpts, _hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "setStaticSignature", _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2 *WalletStage2Session) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2.Contract.SetStaticSignature(&_WalletStage2.TransactOpts, _hash, _address, _timestamp)
}

// SetStaticSignature is a paid mutator transaction binding the contract method 0xf727ef1c.
//
// Solidity: function setStaticSignature(bytes32 _hash, address _address, uint96 _timestamp) returns()
func (_WalletStage2 *WalletStage2TransactorSession) SetStaticSignature(_hash [32]byte, _address common.Address, _timestamp *big.Int) (*types.Transaction, error) {
	return _WalletStage2.Contract.SetStaticSignature(&_WalletStage2.TransactOpts, _hash, _address, _timestamp)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2 *WalletStage2Transactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2 *WalletStage2Session) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.TokenReceived(&_WalletStage2.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage2 *WalletStage2TransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.TokenReceived(&_WalletStage2.TransactOpts, arg0, arg1, arg2)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2 *WalletStage2Transactor) UpdateImageHash(opts *bind.TransactOpts, _imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "updateImageHash", _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2 *WalletStage2Session) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.UpdateImageHash(&_WalletStage2.TransactOpts, _imageHash)
}

// UpdateImageHash is a paid mutator transaction binding the contract method 0x29561426.
//
// Solidity: function updateImageHash(bytes32 _imageHash) returns()
func (_WalletStage2 *WalletStage2TransactorSession) UpdateImageHash(_imageHash [32]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.UpdateImageHash(&_WalletStage2.TransactOpts, _imageHash)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletStage2 *WalletStage2Transactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletStage2 *WalletStage2Session) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.UpdateImplementation(&_WalletStage2.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) returns()
func (_WalletStage2 *WalletStage2TransactorSession) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.UpdateImplementation(&_WalletStage2.TransactOpts, _implementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2 *WalletStage2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2 *WalletStage2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Fallback(&_WalletStage2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WalletStage2 *WalletStage2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Fallback(&_WalletStage2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2 *WalletStage2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletStage2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2 *WalletStage2Session) Receive() (*types.Transaction, error) {
	return _WalletStage2.Contract.Receive(&_WalletStage2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletStage2 *WalletStage2TransactorSession) Receive() (*types.Transaction, error) {
	return _WalletStage2.Contract.Receive(&_WalletStage2.TransactOpts)
}

// WalletStage2AbortedIterator is returned from FilterAborted and is used to iterate over the raw logs and unpacked data for Aborted events raised by the WalletStage2 contract.
type WalletStage2AbortedIterator struct {
	Event *WalletStage2Aborted // Event containing the contract specifics and raw log

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
func (it *WalletStage2AbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2Aborted)
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
		it.Event = new(WalletStage2Aborted)
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
func (it *WalletStage2AbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2AbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2Aborted represents a Aborted event raised by the WalletStage2 contract.
type WalletStage2Aborted struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAborted is a free log retrieval operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterAborted(opts *bind.FilterOpts) (*WalletStage2AbortedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return &WalletStage2AbortedIterator{contract: _WalletStage2.contract, event: "Aborted", logs: logs, sub: sub}, nil
}

// WatchAborted is a free log subscription operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchAborted(opts *bind.WatchOpts, sink chan<- *WalletStage2Aborted) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "Aborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2Aborted)
				if err := _WalletStage2.contract.UnpackLog(event, "Aborted", log); err != nil {
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

// ParseAborted is a log parse operation binding the contract event 0x6cd433d189cb0ff58b321e23f3e510c9d0f019f2230a2066e50962d4f867c0a8.
//
// Solidity: event Aborted(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) ParseAborted(log types.Log) (*WalletStage2Aborted, error) {
	event := new(WalletStage2Aborted)
	if err := _WalletStage2.contract.UnpackLog(event, "Aborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2DefinedHookIterator is returned from FilterDefinedHook and is used to iterate over the raw logs and unpacked data for DefinedHook events raised by the WalletStage2 contract.
type WalletStage2DefinedHookIterator struct {
	Event *WalletStage2DefinedHook // Event containing the contract specifics and raw log

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
func (it *WalletStage2DefinedHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2DefinedHook)
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
		it.Event = new(WalletStage2DefinedHook)
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
func (it *WalletStage2DefinedHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2DefinedHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2DefinedHook represents a DefinedHook event raised by the WalletStage2 contract.
type WalletStage2DefinedHook struct {
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage2 *WalletStage2Filterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage2DefinedHookIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage2DefinedHookIterator{contract: _WalletStage2.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage2 *WalletStage2Filterer) WatchDefinedHook(opts *bind.WatchOpts, sink chan<- *WalletStage2DefinedHook) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2DefinedHook)
				if err := _WalletStage2.contract.UnpackLog(event, "DefinedHook", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseDefinedHook(log types.Log) (*WalletStage2DefinedHook, error) {
	event := new(WalletStage2DefinedHook)
	if err := _WalletStage2.contract.UnpackLog(event, "DefinedHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2FailedIterator is returned from FilterFailed and is used to iterate over the raw logs and unpacked data for Failed events raised by the WalletStage2 contract.
type WalletStage2FailedIterator struct {
	Event *WalletStage2Failed // Event containing the contract specifics and raw log

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
func (it *WalletStage2FailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2Failed)
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
		it.Event = new(WalletStage2Failed)
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
func (it *WalletStage2FailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2FailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2Failed represents a Failed event raised by the WalletStage2 contract.
type WalletStage2Failed struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailed is a free log retrieval operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterFailed(opts *bind.FilterOpts) (*WalletStage2FailedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return &WalletStage2FailedIterator{contract: _WalletStage2.contract, event: "Failed", logs: logs, sub: sub}, nil
}

// WatchFailed is a free log subscription operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchFailed(opts *bind.WatchOpts, sink chan<- *WalletStage2Failed) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "Failed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2Failed)
				if err := _WalletStage2.contract.UnpackLog(event, "Failed", log); err != nil {
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

// ParseFailed is a log parse operation binding the contract event 0xe64040c2a394fc50904b208b60495abbcf56a8eff89806cada4162c27dd5f643.
//
// Solidity: event Failed(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) ParseFailed(log types.Log) (*WalletStage2Failed, error) {
	event := new(WalletStage2Failed)
	if err := _WalletStage2.contract.UnpackLog(event, "Failed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2ImageHashUpdatedIterator is returned from FilterImageHashUpdated and is used to iterate over the raw logs and unpacked data for ImageHashUpdated events raised by the WalletStage2 contract.
type WalletStage2ImageHashUpdatedIterator struct {
	Event *WalletStage2ImageHashUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage2ImageHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2ImageHashUpdated)
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
		it.Event = new(WalletStage2ImageHashUpdated)
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
func (it *WalletStage2ImageHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2ImageHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2ImageHashUpdated represents a ImageHashUpdated event raised by the WalletStage2 contract.
type WalletStage2ImageHashUpdated struct {
	NewImageHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImageHashUpdated is a free log retrieval operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage2 *WalletStage2Filterer) FilterImageHashUpdated(opts *bind.FilterOpts) (*WalletStage2ImageHashUpdatedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage2ImageHashUpdatedIterator{contract: _WalletStage2.contract, event: "ImageHashUpdated", logs: logs, sub: sub}, nil
}

// WatchImageHashUpdated is a free log subscription operation binding the contract event 0x307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa.
//
// Solidity: event ImageHashUpdated(bytes32 newImageHash)
func (_WalletStage2 *WalletStage2Filterer) WatchImageHashUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage2ImageHashUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "ImageHashUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2ImageHashUpdated)
				if err := _WalletStage2.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseImageHashUpdated(log types.Log) (*WalletStage2ImageHashUpdated, error) {
	event := new(WalletStage2ImageHashUpdated)
	if err := _WalletStage2.contract.UnpackLog(event, "ImageHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2ImplementationUpdatedIterator is returned from FilterImplementationUpdated and is used to iterate over the raw logs and unpacked data for ImplementationUpdated events raised by the WalletStage2 contract.
type WalletStage2ImplementationUpdatedIterator struct {
	Event *WalletStage2ImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *WalletStage2ImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2ImplementationUpdated)
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
		it.Event = new(WalletStage2ImplementationUpdated)
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
func (it *WalletStage2ImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2ImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2ImplementationUpdated represents a ImplementationUpdated event raised by the WalletStage2 contract.
type WalletStage2ImplementationUpdated struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterImplementationUpdated is a free log retrieval operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage2 *WalletStage2Filterer) FilterImplementationUpdated(opts *bind.FilterOpts) (*WalletStage2ImplementationUpdatedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return &WalletStage2ImplementationUpdatedIterator{contract: _WalletStage2.contract, event: "ImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchImplementationUpdated is a free log subscription operation binding the contract event 0x310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03.
//
// Solidity: event ImplementationUpdated(address newImplementation)
func (_WalletStage2 *WalletStage2Filterer) WatchImplementationUpdated(opts *bind.WatchOpts, sink chan<- *WalletStage2ImplementationUpdated) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "ImplementationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2ImplementationUpdated)
				if err := _WalletStage2.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseImplementationUpdated(log types.Log) (*WalletStage2ImplementationUpdated, error) {
	event := new(WalletStage2ImplementationUpdated)
	if err := _WalletStage2.contract.UnpackLog(event, "ImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2NonceChangeIterator is returned from FilterNonceChange and is used to iterate over the raw logs and unpacked data for NonceChange events raised by the WalletStage2 contract.
type WalletStage2NonceChangeIterator struct {
	Event *WalletStage2NonceChange // Event containing the contract specifics and raw log

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
func (it *WalletStage2NonceChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2NonceChange)
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
		it.Event = new(WalletStage2NonceChange)
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
func (it *WalletStage2NonceChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2NonceChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2NonceChange represents a NonceChange event raised by the WalletStage2 contract.
type WalletStage2NonceChange struct {
	Space    *big.Int
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceChange is a free log retrieval operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage2 *WalletStage2Filterer) FilterNonceChange(opts *bind.FilterOpts) (*WalletStage2NonceChangeIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return &WalletStage2NonceChangeIterator{contract: _WalletStage2.contract, event: "NonceChange", logs: logs, sub: sub}, nil
}

// WatchNonceChange is a free log subscription operation binding the contract event 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881.
//
// Solidity: event NonceChange(uint256 _space, uint256 _newNonce)
func (_WalletStage2 *WalletStage2Filterer) WatchNonceChange(opts *bind.WatchOpts, sink chan<- *WalletStage2NonceChange) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "NonceChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2NonceChange)
				if err := _WalletStage2.contract.UnpackLog(event, "NonceChange", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseNonceChange(log types.Log) (*WalletStage2NonceChange, error) {
	event := new(WalletStage2NonceChange)
	if err := _WalletStage2.contract.UnpackLog(event, "NonceChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SkippedIterator is returned from FilterSkipped and is used to iterate over the raw logs and unpacked data for Skipped events raised by the WalletStage2 contract.
type WalletStage2SkippedIterator struct {
	Event *WalletStage2Skipped // Event containing the contract specifics and raw log

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
func (it *WalletStage2SkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2Skipped)
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
		it.Event = new(WalletStage2Skipped)
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
func (it *WalletStage2SkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2Skipped represents a Skipped event raised by the WalletStage2 contract.
type WalletStage2Skipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkipped is a free log retrieval operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterSkipped(opts *bind.FilterOpts) (*WalletStage2SkippedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SkippedIterator{contract: _WalletStage2.contract, event: "Skipped", logs: logs, sub: sub}, nil
}

// WatchSkipped is a free log subscription operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage2Skipped) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "Skipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2Skipped)
				if err := _WalletStage2.contract.UnpackLog(event, "Skipped", log); err != nil {
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

// ParseSkipped is a log parse operation binding the contract event 0xa7df37e35254f22900087bd61c5b68001c8f034f7e924ec565af11317d7ee0f7.
//
// Solidity: event Skipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) ParseSkipped(log types.Log) (*WalletStage2Skipped, error) {
	event := new(WalletStage2Skipped)
	if err := _WalletStage2.contract.UnpackLog(event, "Skipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2StaticSignatureSetIterator is returned from FilterStaticSignatureSet and is used to iterate over the raw logs and unpacked data for StaticSignatureSet events raised by the WalletStage2 contract.
type WalletStage2StaticSignatureSetIterator struct {
	Event *WalletStage2StaticSignatureSet // Event containing the contract specifics and raw log

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
func (it *WalletStage2StaticSignatureSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2StaticSignatureSet)
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
		it.Event = new(WalletStage2StaticSignatureSet)
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
func (it *WalletStage2StaticSignatureSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2StaticSignatureSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2StaticSignatureSet represents a StaticSignatureSet event raised by the WalletStage2 contract.
type WalletStage2StaticSignatureSet struct {
	Hash      [32]byte
	Address   common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaticSignatureSet is a free log retrieval operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage2 *WalletStage2Filterer) FilterStaticSignatureSet(opts *bind.FilterOpts) (*WalletStage2StaticSignatureSetIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return &WalletStage2StaticSignatureSetIterator{contract: _WalletStage2.contract, event: "StaticSignatureSet", logs: logs, sub: sub}, nil
}

// WatchStaticSignatureSet is a free log subscription operation binding the contract event 0xebf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1.
//
// Solidity: event StaticSignatureSet(bytes32 _hash, address _address, uint96 _timestamp)
func (_WalletStage2 *WalletStage2Filterer) WatchStaticSignatureSet(opts *bind.WatchOpts, sink chan<- *WalletStage2StaticSignatureSet) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "StaticSignatureSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2StaticSignatureSet)
				if err := _WalletStage2.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseStaticSignatureSet(log types.Log) (*WalletStage2StaticSignatureSet, error) {
	event := new(WalletStage2StaticSignatureSet)
	if err := _WalletStage2.contract.UnpackLog(event, "StaticSignatureSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2SuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the WalletStage2 contract.
type WalletStage2SuccessIterator struct {
	Event *WalletStage2Success // Event containing the contract specifics and raw log

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
func (it *WalletStage2SuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2Success)
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
		it.Event = new(WalletStage2Success)
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
func (it *WalletStage2SuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2SuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2Success represents a Success event raised by the WalletStage2 contract.
type WalletStage2Success struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterSuccess(opts *bind.FilterOpts) (*WalletStage2SuccessIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return &WalletStage2SuccessIterator{contract: _WalletStage2.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *WalletStage2Success) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "Success")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2Success)
				if err := _WalletStage2.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x2fd98f16e3e0ef7b9373ea49ea6b76b871c7f2aa1e2c222747ef5bfb26de18b3.
//
// Solidity: event Success(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) ParseSuccess(log types.Log) (*WalletStage2Success, error) {
	event := new(WalletStage2Success)
	if err := _WalletStage2.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
