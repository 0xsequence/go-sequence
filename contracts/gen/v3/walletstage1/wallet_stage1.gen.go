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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidERC1271Signature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_type\",\"type\":\"bytes1\"}],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_expectedCaller\",\"type\":\"address\"}],\"name\":\"InvalidStaticSignatureWrongCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CODE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAGE_2_IMPLEMENTATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getStaticSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526004361061012d575f3560e01c80636ea44577116100aa578063aaf10f421161006e578063aaf10f42146104b5578063ad55366b146104df578063b93ea7ad14610520578063bc197c811461053c578063f23a6e6114610578578063f727ef1c146105b457610134565b80636ea44577146103ce5780638943ec02146103ea5780638c3f55631461041257806392dcb3fc1461044e5780639f69ef541461048b57610134565b80631f6a1eb9116100f15780631f6a1eb91461031a578063257671f51461033657806329561426146103605780632dd31000146103885780634fcf3eca146103b257610134565b8063025b22bc1461020e57806313792a4a1461022a578063150b7a02146102665780631626ba7e146102a25780631a9b2337146102de57610134565b3661013457005b60045f3690501061020c575f6101555f369061015091906135a8565b6105dc565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461020a575f5f8273ffffffffffffffffffffffffffffffffffffffff165f366040516101b3929190613642565b5f60405180830381855af49150503d805f81146101eb576040519150601f19603f3d011682016040523d82523d5f602084013e6101f0565b606091505b50915091508161020257805160208201fd5b805160208201f35b505b005b610228600480360381019061022391906136c5565b610631565b005b348015610235575f5ffd5b50610250600480360381019061024b9190613cf0565b6106ad565b60405161025d9190613d78565b60405180910390f35b348015610271575f5ffd5b5061028c60048036038101906102879190613d91565b61085c565b6040516102999190613e24565b60405180910390f35b3480156102ad575f5ffd5b506102c860048036038101906102c39190613e3d565b610870565b6040516102d59190613e24565b60405180910390f35b3480156102e9575f5ffd5b5061030460048036038101906102ff9190613ec4565b6108ca565b6040516103119190613efe565b60405180910390f35b610334600480360381019061032f9190613f17565b6108db565b005b348015610341575f5ffd5b5061034a61096a565b6040516103579190613d78565b60405180910390f35b34801561036b575f5ffd5b5061038660048036038101906103819190613f95565b61098e565b005b348015610393575f5ffd5b5061039c610a0a565b6040516103a99190613efe565b60405180910390f35b6103cc60048036038101906103c79190613ec4565b610a2e565b005b6103e860048036038101906103e39190613fc0565b610b23565b005b3480156103f5575f5ffd5b50610410600480360381019061040b919061400b565b610bc2565b005b34801561041d575f5ffd5b506104386004803603810190610433919061407c565b610bc8565b60405161044591906140b6565b60405180910390f35b348015610459575f5ffd5b50610474600480360381019061046f9190613f95565b610c00565b6040516104829291906140cf565b60405180910390f35b348015610496575f5ffd5b5061049f610c14565b6040516104ac9190613efe565b60405180910390f35b3480156104c0575f5ffd5b506104c9610c38565b6040516104d69190613efe565b60405180910390f35b3480156104ea575f5ffd5b5061050560048036038101906105009190613cf0565b610c46565b60405161051796959493929190614105565b60405180910390f35b61053a60048036038101906105359190614164565b610c84565b005b348015610547575f5ffd5b50610562600480360381019061055d91906141f7565b610d7a565b60405161056f9190613e24565b60405180910390f35b348015610583575f5ffd5b5061059e600480360381019061059991906142ce565b610d91565b6040516105ab9190613e24565b60405180910390f35b3480156105bf575f5ffd5b506105da60048036038101906105d591906143a5565b610da6565b005b5f6106287fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610e6f565b5f1c9050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106a157336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016106989190613efe565b60405180910390fd5b6106aa81610ea7565b50565b5f5f6001856101000151516106c29190614422565b67ffffffffffffffff8111156106db576106da613704565b5b6040519080825280602002602001820160405280156107095781602001602082028036833780820191505090505b5090505f5f90505b8561010001515181101561079957856101000151818151811061073757610736614455565b5b602002602001015182828151811061075257610751614455565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610711565b50338186610100015151815181106107b4576107b3614455565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808561010001819052505f610804868686610eea565b5090508061084d578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161084493929190614835565b60405180910390fd5b60015f1b925050509392505050565b5f63150b7a0260e01b905095945050505050565b5f6108796134b3565b6003815f019060ff16908160ff1681525050848160e00181815250505f6108a1828686610eea565b509050806108b6575f60e01b925050506108c3565b6320c13b0b60e01b925050505b9392505050565b5f6108d4826105dc565b9050919050565b5f5a90505f6108ea86866110da565b90506108fe8160600151826080015161157c565b5f5f61090b838787610eea565b9150915081610955578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161094c93929190614835565b60405180910390fd5b610960848285611620565b5050505050505050565b7f4ef049b07ac1b33f2f9562706ecb68a55b978a125ed5d9a4177ccc9eba1ee1d381565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109fe57336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016109f59190613efe565b60405180910390fd5b610a0781611952565b50565b7f000000000000000000000000bd0f8abd58b4449b39c57ac9d5c67433239ac44781565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a9e57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610a959190613efe565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610abe826105dc565b73ffffffffffffffffffffffffffffffffffffffff1603610b1657806040517f1c3812cc000000000000000000000000000000000000000000000000000000008152600401610b0d9190613e24565b60405180910390fd5b610b20815f611a1c565b50565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b9357336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610b8a9190613efe565b60405180910390fd5b5f5a90505f610ba284846110da565b90505f610bae82611abd565b9050610bbb838284611620565b5050505050565b50505050565b5f610bf77f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b610e6f565b5f1c9050919050565b5f5f610c0b83611b0d565b91509150915091565b7f000000000000000000000000dad4fb19ec2044f8fece8e85ccba0aff59ad7bfd81565b5f610c41611b5e565b905090565b5f5f5f5f5f5f610c598989895f5f611b66565b809550819650829750839950849a505050505050610c7683611ea3565b935093975093979195509350565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cf457336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610ceb9190613efe565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610d14836105dc565b73ffffffffffffffffffffffffffffffffffffffff1614610d6c57816040517f5b4d6d6a000000000000000000000000000000000000000000000000000000008152600401610d639190613e24565b60405180910390fd5b610d768282611a1c565b5050565b5f63bc197c8160e01b905098975050505050505050565b5f63f23a6e6160e01b90509695505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e1657336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610e0d9190613efe565b60405180910390fd5b610e2f8383836bffffffffffffffffffffffff16611eb4565b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1838383604051610e629392919061487b565b60405180910390a1505050565b5f5f8383604051602001610e849291906148b0565b604051602081830303815290604052805190602001209050805491505092915050565b610eb081611f12565b7f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0381604051610edf9190613efe565b60405180910390a150565b5f5f5f84845f818110610f0057610eff614455565b5b9050013560f81c60f81b9050608060f81b608060f81b82167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19160361105957610f4886611abd565b91505f5f610f5584611b0d565b91509150428111610f9f5783816040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600401610f969291906148d7565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415801561100757503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561104d578333836040517f8945c313000000000000000000000000000000000000000000000000000000008152600401611044939291906148fe565b60405180910390fd5b600194505050506110d2565b5f5f5f6110698989895f5f611b66565b905080985081945082955083965050505050828210156110c25782826040517ffd41fcba0000000000000000000000000000000000000000000000000000000081526004016110b9929190614933565b60405180910390fd5b6110cb81611ea3565b9550505050505b935093915050565b6110e26134b3565b5f815f019060ff16908160ff16815250505f5f6110ff8585611f18565b915060ff16915060018083160361111f575f83606001818152505061115b565b611134818686611f2e9290919263ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff169150846060018193508281525050505b5f6007600184901c1690505f8111156111985761118a82828888611f5f9190939291909392919063ffffffff16565b856080018194508281525050505b5f6010808516036111ac5760019050611204565b6020808516036111df576111cb838888611f8c9290919263ffffffff16565b8161ffff1691508094508192505050611203565b6111f4838888611fab9290919263ffffffff16565b8160ff16915080945081925050505b5b8067ffffffffffffffff81111561121e5761121d613704565b5b60405190808252806020026020018201604052801561125757816020015b6112446134fe565b81526020019060019003908161123c5790505b5085604001819052505f5f90505b81811015611571575f611283858a8a611fab9290919263ffffffff16565b8096508192505050600180821660ff16036112f25730876040015183815181106112b0576112af614455565b5b60200260200101515f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061135e565b611307858a8a611fc69290919263ffffffff16565b8860400151848151811061131e5761131d614455565b5b60200260200101515f018197508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b600280821660ff16036113ac57611380858a8a611ff79290919263ffffffff16565b8860400151848151811061139757611396614455565b5b60200260200101516020018197508281525050505b600480821660ff1603611474575f6113cf868b8b61200d9290919263ffffffff16565b8162ffffff16915080975081925050508989879083896113ef9190614422565b926113fc93929190614962565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508860400151848151811061145557611454614455565b5b60200260200101516040018190525080866114709190614422565b9550505b600880821660ff16036114c257611496858a8a611ff79290919263ffffffff16565b886040015184815181106114ad576114ac614455565b5b60200260200101516060018197508281525050505b601080821660ff1614876040015183815181106114e2576114e1614455565b5b60200260200101516080019015159081151581525050602080821660ff16148760400151838151811061151857611517614455565b5b602002602001015160a0019015159081151581525050600660c0821660ff16901c60ff168760400151838151811061155357611552614455565b5b602002602001015160c0018181525050508080600101915050611265565b505050505092915050565b5f61158683610bc8565b90508181146115d0578282826040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004016115c79392919061499c565b60405180910390fd5b5f6001830190506115e1848261202d565b7f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f8818482604051611612929190614933565b60405180910390a150505050565b5f5f90505f82604001515190505f5f90505b8181101561194a575f8460400151828151811061165257611651614455565b5b60200260200101519050831561166a575f93506116b4565b8060a00151156116b3577f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b86836040516116a59291906148d7565b60405180910390a15061193d565b5b5f816060015190505f81141580156116cb5750805a105b156117115785835a6040517f21395274000000000000000000000000000000000000000000000000000000008152600401611708939291906149d1565b60405180910390fd5b5f8260800151156117c6576117bf835f01515f84146117305783611732565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161175d96959493929190614a45565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050612062565b90506117ee565b6117eb835f015184602001515f85146117df57846117e1565b5a5b8660400151612077565b90505b80611900575f60ff168360c001510361184f57600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d888561183061208e565b60405161183f93929190614aab565b60405180910390a150505061193d565b600160ff168360c00151036118a657868461186861208e565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161189d93929190614ae7565b60405180910390fd5b600260ff168360c00151036118ff577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856118e061208e565b6040516118ef93929190614aab565b60405180910390a150505061194a565b5b7f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a88856040516119319291906148d7565b60405180910390a15050505b8080600101915050611632565b505050505050565b5f5f1b810361198d576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119b97fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b826120ac565b7f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa816040516119e89190613d78565b60405180910390a1611a197f000000000000000000000000dad4fb19ec2044f8fece8e85ccba0aff59ad7bfd610ea7565b50565b611a807fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff165f1b6120b3565b7f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed18282604051611ab1929190614b2a565b60405180910390a15050565b5f5f611acd8360200151306120e8565b90505f611ad98461218c565b90508181604051602001611aee929190614bc5565b6040516020818303038152906040528051906020012092505050919050565b5f5f5f611b3c7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b85610e6f565b5f1c9050606081901c816bffffffffffffffffffffffff169250925050915091565b5f3054905090565b5f5f5f5f5f5f5f611b778b8b611f18565b915060ff169150611b8661354f565b6040808416148015611bc357505f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff16145b15611cff57611bdd828d8d611fc69290919263ffffffff16565b809350819a50505089611cfe575f611c00838e8e61200d9290919263ffffffff16565b8162ffffff16915080945081925050505f8d8d85908487611c219190614422565b92611c2e93929190614962565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090508a73ffffffffffffffffffffffffffffffffffffffff1663ccce3bc830836040518363ffffffff1660e01b8152600401611cad929190614bfb565b6040805180830381865afa158015611cc7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ceb9190614c9e565b92508184611cf99190614422565b935050505b5b600180841603611d3857611d268d8a838f8f87908092611d2193929190614962565b6123c0565b97509750975097509750505050611e96565b6002808416148d60200190151590811515815250505f6002601c8516901c9050611d7483828f8f611f5f9190939291909392919063ffffffff16565b8094508197505050505f6001600560208616901c611d929190614422565b9050611db083828f8f611f5f9190939291909392919063ffffffff16565b809450819a50505050611dc28d611abd565b9350611de08d858e8e86908092611ddb93929190614962565b612625565b8097508198505050611df486895f1b6131a2565b9550611e0286865f1b6131a2565b9550611e26868a73ffffffffffffffffffffffffffffffffffffffff165f1b6131a2565b95505f5f1b815f015114158015611e40575085815f015114155b8015611e50575080602001518511155b15611e9257806040517fccbb534f000000000000000000000000000000000000000000000000000000008152600401611e899190614cf6565b60405180910390fd5b5050505b9550955095509550959050565b5f611ead826131b6565b9050919050565b611f0d7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b846bffffffffffffffffffffffff841660608673ffffffffffffffffffffffffffffffffffffffff16901b175f1b6120b3565b505050565b80305550565b5f5f83358060f81c925060019150509250929050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f858401356008840261010003600180866008021b0382821c1693508486019250505094509492505050565b5f5f8483013561ffff8160f01c16925060028401915050935093915050565b5f5f848301358060f81c925060018401915050935093915050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f848301359150602083019050935093915050565b5f5f8483013562ffffff8160e81c16925060038401915050935093915050565b61205e7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b835f1b6120b3565b5050565b5f5f5f8351602085018787f490509392505050565b5f5f5f835160208501878988f19050949350505050565b60603d604051915060208201818101604052818352815f823e505090565b8082555050565b5f83836040516020016120c79291906148b0565b60405160208183030381529060405280519060200120905081815550505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856121575746612159565b5f5b8560405160200161216e959493929190614d0f565b60405160208183030381529060405280519060200120905092915050565b5f5f8261010001516040516020016121a49190614dec565b6040516020818303038152906040528051906020012090505f60ff16835f015160ff160361223c575f6121da8460400151613259565b90507f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a281856060015186608001518560405160200161221d959493929190614e02565b60405160208183030381529060405280519060200120925050506123bb565b600160ff16835f015160ff16036122ab577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360a00151805190602001208260405160200161228d93929190614e53565b604051602081830303815290604052805190602001209150506123bb565b600260ff16835f015160ff1603612313577f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e48360c00151826040516020016122f593929190614e53565b604051602081830303815290604052805190602001209150506123bb565b600360ff16835f015160ff160361237b577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360e001518260405160200161235d93929190614e53565b604051602081830303815290604052805190602001209150506123bb565b825f01516040517f048183200000000000000000000000000000000000000000000000000000000081526004016123b29190614e97565b60405180910390fd5b919050565b5f5f5f5f5f6123cd6134b3565b6002815f019060ff16908160ff16815250505f5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b898990508210156125ba575f5f612427848d8d61200d9290919263ffffffff16565b8162ffffff169150809550819250505083816124439190614422565b9150505f8b8b90508214612457575f612459565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83036124b9576124a08f8d8d8790869261249893929190614962565b600185611b66565b809a50819b50829c50839d50849e5050505050506124e9565b6124d7858d8d879086926124cf93929190614962565b600185611b66565b50809a50819b50829c50839d50505050505b89891015612544578b8b8590849261250393929190614962565b8b8b6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161253b9493929190614eb0565b60405180910390fd5b819350878d5f01510361255d575f5f1b8d5f0181815250505b8287106125a35786836040517f37daf62b00000000000000000000000000000000000000000000000000000000815260040161259a929190614933565b60405180910390fd5b878560c00181815250508692508193505050612405565b5f5f1b8b5f0151141580156125d357508a602001518511155b15612615578a6040517fccbb534f00000000000000000000000000000000000000000000000000000000815260040161260c9190614cf6565b60405180910390fd5b5050509550955095509550959050565b5f5f5f5b84849050811015613198575f61264a828787611fab9290919263ffffffff16565b8160ff16915080935081925050505f600460f08316901c90505f81036127a2575f600f831690505f8160ff160361269957612690848989611fab9290919263ffffffff16565b80955081925050505b5f5f6126b0868b8b6132d39290919263ffffffff16565b80975081935050506126cd868b8b6132d39290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f8388866040515f81526020016040526040516127339493929190614eee565b6020604051602081039080840390855afa158015612753573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612775828960ff166132e9565b90505f5f1b8c036127865780612791565b6127908c826131a2565b5b9b5050505050505050505050612629565b6001810361282d575f600f831690505f8160ff16036127d9576127d0848989611fab9290919263ffffffff16565b80955081925050505b5f6127ef858a8a611fc69290919263ffffffff16565b80965081925050505f612805828460ff166132e9565b90505f5f1b87036128165780612821565b61282087826131a2565b5b96505050505050612629565b60028103612a2c575f6003831690505f8160ff16036128645761285b848989611fab9290919263ffffffff16565b80955081925050505b5f61287a858a8a611fc69290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f6128b087838d8d611f5f9190939291909392919063ffffffff16565b80985081925050505f81880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261291493929190614962565b6040518463ffffffff1660e01b815260040161293293929190614f31565b602060405180830381865afa15801561294d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129719190614f75565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146129e8578c848d8d8b9085926129a993929190614962565b6040517fb2fed7ae0000000000000000000000000000000000000000000000000000000081526004016129df9493929190614fa0565b60405180910390fd5b8097508460ff168a0199505f612a01858760ff166132e9565b90505f5f1b8a03612a125780612a1d565b612a1c8a826131a2565b5b99505050505050505050612629565b60038103612a76575f612a4a8489896132d39290919263ffffffff16565b80955081925050505f5f1b8503612a615780612a6c565b612a6b85826131a2565b5b9450505050612629565b60048103612af8575f600f831660ff1690505f612aa585838b8b611f5f9190939291909392919063ffffffff16565b80965081925050505f81860190505f5f612ad18e8e8e8e8c908892612acc93929190614962565b612625565b91509150829750818a019950612ae789826131a2565b985082975050505050505050612629565b60068103612c08575f6002600c841660ff16901c60ff1690505f8103612b3c57612b2d848989611fab9290919263ffffffff16565b8160ff16915080955081925050505b5f6003841660ff1690505f8103612b7257612b62858a8a611f8c9290919263ffffffff16565b8161ffff16915080965081925050505b5f612b88868b8b61200d9290919263ffffffff16565b8162ffffff16915080975081925050505f81870190505f5f612bbc8f8f8f8f8d908892612bb793929190614962565b612625565b91509150829850848210612bd057858b019a505b5f612bdc82878961331b565b90505f5f1b8b03612bed5780612bf8565b612bf78b826131a2565b5b9a50505050505050505050612629565b60058103612c8a575f612c268489896132d39290919263ffffffff16565b8095508192505050888103612c59577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b5f612c6382613350565b90505f5f1b8603612c745780612c7f565b612c7e86826131a2565b5b955050505050612629565b60078103612df0575f600f831690505f8160ff1603612cc157612cb8848989611fab9290919263ffffffff16565b80955081925050505b5f5f612cd8868b8b6132d39290919263ffffffff16565b8097508193505050612cf5868b8b6132d39290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f604051602001612d4c9190615028565b604051602081830303815290604052805190602001208388866040515f8152602001604052604051612d819493929190614eee565b6020604051602081039080840390855afa158015612da1573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612dc3828960ff166132e9565b90505f5f1b8c03612dd45780612ddf565b612dde8c826131a2565b5b9b5050505050505050505050612629565b60088103612e89575f612e0e8489896132d39290919263ffffffff16565b80955081925050505f612e2a5f8c61337f90919063ffffffff16565b9050808203612e57577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b5f612e61826133d0565b90505f5f1b8703612e725780612e7d565b612e7c87826131a2565b5b96505050505050612629565b60098103612ff2575f6003831690505f8160ff1603612ec057612eb7848989611fab9290919263ffffffff16565b80955081925050505b5f612ed6858a8a611fc69290919263ffffffff16565b80965081925050505f5f6002600c871660ff16901c60ff169050612f0c87828d8d611f5f9190939291909392919063ffffffff16565b8098508193505050505f81870190505f8373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612f4b93929190614962565b6040518463ffffffff1660e01b8152600401612f699392919061504d565b602060405180830381865afa158015612f84573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612fa89190615084565b90508197508460ff168a0199505f612fc4858760ff16846133ff565b90505f5f1b8a03612fd55780612fe0565b612fdf8a826131a2565b5b99508298505050505050505050612629565b600a810361315b575f6003831690505f8160ff160361302957613020848989611fab9290919263ffffffff16565b80955081925050505b5f61303f858a8a611fc69290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f61307587838d8d611f5f9190939291909392919063ffffffff16565b80985081925050505f81880190505f8473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d9087926130b393929190614962565b6040518463ffffffff1660e01b81526004016130d193929190614f31565b602060405180830381865afa1580156130ec573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131109190615084565b90508198508560ff168b019a505f61312c868860ff16846133ff565b90505f5f1b8b0361313d5780613148565b6131478b826131a2565b5b9a50829950505050505050505050612629565b806040517fb2505f7c00000000000000000000000000000000000000000000000000000000815260040161318f91906140b6565b60405180910390fd5b5094509492505050565b5f825f528160205260405f20905092915050565b5f3073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000bd0f8abd58b4449b39c57ac9d5c67433239ac447837f4ef049b07ac1b33f2f9562706ecb68a55b978a125ed5d9a4177ccc9eba1ee1d36040516020016132239392919061513e565b604051602081830303815290604052805190602001205f1c73ffffffffffffffffffffffffffffffffffffffff16149050919050565b5f60605f5f90505b83518110156132c2575f61328e85838151811061328157613280614455565b5b6020026020010151613434565b905082816040516020016132a39291906151b5565b6040516020818303038152906040529250508080600101915050613261565b508080519060200120915050919050565b5f5f848301359150602083019050935093915050565b5f82826040516020016132fd929190615246565b60405160208183030381529060405280519060200120905092915050565b5f838383604051602001613331939291906152c6565b6040516020818303038152906040528051906020012090509392505050565b5f816040516020016133629190615357565b604051602081830303815290604052805190602001209050919050565b5f5f61338f8460200151846120e8565b90505f61339b8561218c565b905081816040516020016133b0929190614bc5565b604051602081830303815290604052805190602001209250505092915050565b5f816040516020016133e291906153c6565b604051602081830303815290604052805190602001209050919050565b5f83838360405160200161341593929190615435565b6040516020818303038152906040528051906020012090509392505050565b5f7f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437825f01518360200151846040015180519060200120856060015186608001518760a001518860c0015160405160200161349698979695949392919061547c565b604051602081830303815290604052805190602001209050919050565b6040518061012001604052805f60ff1681526020015f15158152602001606081526020015f81526020015f8152602001606081526020015f81526020015f8152602001606081525090565b6040518060e001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020015f81526020015f151581526020015f151581526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f82905092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f82821b905092915050565b5f6135b38383613567565b826135be8135613571565b925060048210156135fe576135f97fffffffff000000000000000000000000000000000000000000000000000000008360040360080261359c565b831692505b505092915050565b5f81905092915050565b828183375f83830152505050565b5f6136298385613606565b9350613636838584613610565b82840190509392505050565b5f61364e82848661361e565b91508190509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6136948261366b565b9050919050565b6136a48161368a565b81146136ae575f5ffd5b50565b5f813590506136bf8161369b565b92915050565b5f602082840312156136da576136d9613663565b5b5f6136e7848285016136b1565b91505092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61373a826136f4565b810181811067ffffffffffffffff8211171561375957613758613704565b5b80604052505050565b5f61376b61365a565b90506137778282613731565b919050565b5f5ffd5b5f60ff82169050919050565b61379581613780565b811461379f575f5ffd5b50565b5f813590506137b08161378c565b92915050565b5f8115159050919050565b6137ca816137b6565b81146137d4575f5ffd5b50565b5f813590506137e5816137c1565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561380957613808613704565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b6138308161381e565b811461383a575f5ffd5b50565b5f8135905061384b81613827565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561386f5761386e613704565b5b613878826136f4565b9050602081019050919050565b5f61389761389284613855565b613762565b9050828152602081018484840111156138b3576138b2613851565b5b6138be848285613610565b509392505050565b5f82601f8301126138da576138d96137eb565b5b81356138ea848260208601613885565b91505092915050565b5f60e08284031215613908576139076136f0565b5b61391260e0613762565b90505f613921848285016136b1565b5f8301525060206139348482850161383d565b602083015250604082013567ffffffffffffffff8111156139585761395761377c565b5b613964848285016138c6565b60408301525060606139788482850161383d565b606083015250608061398c848285016137d7565b60808301525060a06139a0848285016137d7565b60a08301525060c06139b48482850161383d565b60c08301525092915050565b5f6139d26139cd846137ef565b613762565b905080838252602082019050602084028301858111156139f5576139f461381a565b5b835b81811015613a3c57803567ffffffffffffffff811115613a1a57613a196137eb565b5b808601613a2789826138f3565b855260208501945050506020810190506139f7565b5050509392505050565b5f82601f830112613a5a57613a596137eb565b5b8135613a6a8482602086016139c0565b91505092915050565b5f819050919050565b613a8581613a73565b8114613a8f575f5ffd5b50565b5f81359050613aa081613a7c565b92915050565b5f67ffffffffffffffff821115613ac057613abf613704565b5b602082029050602081019050919050565b5f613ae3613ade84613aa6565b613762565b90508083825260208201905060208402830185811115613b0657613b0561381a565b5b835b81811015613b2f5780613b1b88826136b1565b845260208401935050602081019050613b08565b5050509392505050565b5f82601f830112613b4d57613b4c6137eb565b5b8135613b5d848260208601613ad1565b91505092915050565b5f6101208284031215613b7c57613b7b6136f0565b5b613b87610120613762565b90505f613b96848285016137a2565b5f830152506020613ba9848285016137d7565b602083015250604082013567ffffffffffffffff811115613bcd57613bcc61377c565b5b613bd984828501613a46565b6040830152506060613bed8482850161383d565b6060830152506080613c018482850161383d565b60808301525060a082013567ffffffffffffffff811115613c2557613c2461377c565b5b613c31848285016138c6565b60a08301525060c0613c4584828501613a92565b60c08301525060e0613c5984828501613a92565b60e08301525061010082013567ffffffffffffffff811115613c7e57613c7d61377c565b5b613c8a84828501613b39565b6101008301525092915050565b5f5ffd5b5f5f83601f840112613cb057613caf6137eb565b5b8235905067ffffffffffffffff811115613ccd57613ccc613c97565b5b602083019150836001820283011115613ce957613ce861381a565b5b9250929050565b5f5f5f60408486031215613d0757613d06613663565b5b5f84013567ffffffffffffffff811115613d2457613d23613667565b5b613d3086828701613b66565b935050602084013567ffffffffffffffff811115613d5157613d50613667565b5b613d5d86828701613c9b565b92509250509250925092565b613d7281613a73565b82525050565b5f602082019050613d8b5f830184613d69565b92915050565b5f5f5f5f5f60808688031215613daa57613da9613663565b5b5f613db7888289016136b1565b9550506020613dc8888289016136b1565b9450506040613dd98882890161383d565b935050606086013567ffffffffffffffff811115613dfa57613df9613667565b5b613e0688828901613c9b565b92509250509295509295909350565b613e1e81613571565b82525050565b5f602082019050613e375f830184613e15565b92915050565b5f5f5f60408486031215613e5457613e53613663565b5b5f613e6186828701613a92565b935050602084013567ffffffffffffffff811115613e8257613e81613667565b5b613e8e86828701613c9b565b92509250509250925092565b613ea381613571565b8114613ead575f5ffd5b50565b5f81359050613ebe81613e9a565b92915050565b5f60208284031215613ed957613ed8613663565b5b5f613ee684828501613eb0565b91505092915050565b613ef88161368a565b82525050565b5f602082019050613f115f830184613eef565b92915050565b5f5f5f5f60408587031215613f2f57613f2e613663565b5b5f85013567ffffffffffffffff811115613f4c57613f4b613667565b5b613f5887828801613c9b565b9450945050602085013567ffffffffffffffff811115613f7b57613f7a613667565b5b613f8787828801613c9b565b925092505092959194509250565b5f60208284031215613faa57613fa9613663565b5b5f613fb784828501613a92565b91505092915050565b5f5f60208385031215613fd657613fd5613663565b5b5f83013567ffffffffffffffff811115613ff357613ff2613667565b5b613fff85828601613c9b565b92509250509250929050565b5f5f5f5f6060858703121561402357614022613663565b5b5f614030878288016136b1565b94505060206140418782880161383d565b935050604085013567ffffffffffffffff81111561406257614061613667565b5b61406e87828801613c9b565b925092505092959194509250565b5f6020828403121561409157614090613663565b5b5f61409e8482850161383d565b91505092915050565b6140b08161381e565b82525050565b5f6020820190506140c95f8301846140a7565b92915050565b5f6040820190506140e25f830185613eef565b6140ef60208301846140a7565b9392505050565b6140ff816137b6565b82525050565b5f60c0820190506141185f8301896140a7565b61412560208301886140a7565b61413260408301876140f6565b61413f6060830186613d69565b61414c60808301856140a7565b61415960a0830184613d69565b979650505050505050565b5f5f6040838503121561417a57614179613663565b5b5f61418785828601613eb0565b9250506020614198858286016136b1565b9150509250929050565b5f5f83601f8401126141b7576141b66137eb565b5b8235905067ffffffffffffffff8111156141d4576141d3613c97565b5b6020830191508360208202830111156141f0576141ef61381a565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b03121561421357614212613663565b5b5f6142208b828c016136b1565b98505060206142318b828c016136b1565b975050604089013567ffffffffffffffff81111561425257614251613667565b5b61425e8b828c016141a2565b9650965050606089013567ffffffffffffffff81111561428157614280613667565b5b61428d8b828c016141a2565b9450945050608089013567ffffffffffffffff8111156142b0576142af613667565b5b6142bc8b828c01613c9b565b92509250509295985092959890939650565b5f5f5f5f5f5f60a087890312156142e8576142e7613663565b5b5f6142f589828a016136b1565b965050602061430689828a016136b1565b955050604061431789828a0161383d565b945050606061432889828a0161383d565b935050608087013567ffffffffffffffff81111561434957614348613667565b5b61435589828a01613c9b565b92509250509295509295509295565b5f6bffffffffffffffffffffffff82169050919050565b61438481614364565b811461438e575f5ffd5b50565b5f8135905061439f8161437b565b92915050565b5f5f5f606084860312156143bc576143bb613663565b5b5f6143c986828701613a92565b93505060206143da868287016136b1565b92505060406143eb86828701614391565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61442c8261381e565b91506144378361381e565b925082820190508082111561444f5761444e6143f5565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61448b81613780565b82525050565b61449a816137b6565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6144d28161368a565b82525050565b6144e18161381e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f614519826144e7565b61452381856144f1565b9350614533818560208601614501565b61453c816136f4565b840191505092915050565b5f60e083015f83015161455c5f8601826144c9565b50602083015161456f60208601826144d8565b5060408301518482036040860152614587828261450f565b915050606083015161459c60608601826144d8565b5060808301516145af6080860182614491565b5060a08301516145c260a0860182614491565b5060c08301516145d560c08601826144d8565b508091505092915050565b5f6145eb8383614547565b905092915050565b5f602082019050919050565b5f614609826144a0565b61461381856144aa565b935083602082028501614625856144ba565b805f5b85811015614660578484038952815161464185826145e0565b945061464c836145f3565b925060208a01995050600181019050614628565b50829750879550505050505092915050565b61467b81613a73565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f6146b583836144c9565b60208301905092915050565b5f602082019050919050565b5f6146d782614681565b6146e1818561468b565b93506146ec8361469b565b805f5b8381101561471c57815161470388826146aa565b975061470e836146c1565b9250506001810190506146ef565b5085935050505092915050565b5f61012083015f83015161473f5f860182614482565b5060208301516147526020860182614491565b506040830151848203604086015261476a82826145ff565b915050606083015161477f60608601826144d8565b50608083015161479260808601826144d8565b5060a083015184820360a08601526147aa828261450f565b91505060c08301516147bf60c0860182614672565b5060e08301516147d260e0860182614672565b506101008301518482036101008601526147ec82826146cd565b9150508091505092915050565b5f82825260208201905092915050565b5f61481483856147f9565b9350614821838584613610565b61482a836136f4565b840190509392505050565b5f6040820190508181035f83015261484d8186614729565b90508181036020830152614862818486614809565b9050949350505050565b61487581614364565b82525050565b5f60608201905061488e5f830186613d69565b61489b6020830185613eef565b6148a8604083018461486c565b949350505050565b5f6040820190506148c35f830185613d69565b6148d06020830184613d69565b9392505050565b5f6040820190506148ea5f830185613d69565b6148f760208301846140a7565b9392505050565b5f6060820190506149115f830186613d69565b61491e6020830185613eef565b61492b6040830184613eef565b949350505050565b5f6040820190506149465f8301856140a7565b61495360208301846140a7565b9392505050565b5f5ffd5b5f5ffd5b5f5f858511156149755761497461495a565b5b838611156149865761498561495e565b5b6001850283019150848603905094509492505050565b5f6060820190506149af5f8301866140a7565b6149bc60208301856140a7565b6149c960408301846140a7565b949350505050565b5f6060820190508181035f8301526149e98186614729565b90506149f860208301856140a7565b614a0560408301846140a7565b949350505050565b5f614a17826144e7565b614a2181856147f9565b9350614a31818560208601614501565b614a3a816136f4565b840191505092915050565b5f60c082019050614a585f830189613d69565b614a6560208301886140a7565b614a7260408301876140a7565b614a7f60608301866140a7565b614a8c60808301856140a7565b81810360a0830152614a9e8184614a0d565b9050979650505050505050565b5f606082019050614abe5f830186613d69565b614acb60208301856140a7565b8181036040830152614add8184614a0d565b9050949350505050565b5f6060820190508181035f830152614aff8186614729565b9050614b0e60208301856140a7565b8181036040830152614b208184614a0d565b9050949350505050565b5f604082019050614b3d5f830185613e15565b614b4a6020830184613eef565b9392505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f614b8f600283614b51565b9150614b9a82614b5b565b600282019050919050565b5f819050919050565b614bbf614bba82613a73565b614ba5565b82525050565b5f614bcf82614b83565b9150614bdb8285614bae565b602082019150614beb8284614bae565b6020820191508190509392505050565b5f604082019050614c0e5f830185613eef565b8181036020830152614c208184614a0d565b90509392505050565b5f81519050614c3781613a7c565b92915050565b5f81519050614c4b81613827565b92915050565b5f60408284031215614c6657614c656136f0565b5b614c706040613762565b90505f614c7f84828501614c29565b5f830152506020614c9284828501614c3d565b60208301525092915050565b5f60408284031215614cb357614cb2613663565b5b5f614cc084828501614c51565b91505092915050565b604082015f820151614cdd5f850182614672565b506020820151614cf060208501826144d8565b50505050565b5f604082019050614d095f830184614cc9565b92915050565b5f60a082019050614d225f830188613d69565b614d2f6020830187613d69565b614d3c6040830186613d69565b614d4960608301856140a7565b614d566080830184613eef565b9695505050505050565b5f81905092915050565b614d738161368a565b82525050565b5f614d848383614d6a565b60208301905092915050565b5f614d9a82614681565b614da48185614d60565b9350614daf8361469b565b805f5b83811015614ddf578151614dc68882614d79565b9750614dd1836146c1565b925050600181019050614db2565b5085935050505092915050565b5f614df78284614d90565b915081905092915050565b5f60a082019050614e155f830188613d69565b614e226020830187613d69565b614e2f60408301866140a7565b614e3c60608301856140a7565b614e496080830184613d69565b9695505050505050565b5f606082019050614e665f830186613d69565b614e736020830185613d69565b614e806040830184613d69565b949350505050565b614e9181613780565b82525050565b5f602082019050614eaa5f830184614e88565b92915050565b5f6060820190508181035f830152614ec9818688614809565b9050614ed860208301856140a7565b614ee560408301846140a7565b95945050505050565b5f608082019050614f015f830187613d69565b614f0e6020830186614e88565b614f1b6040830185613d69565b614f286060830184613d69565b95945050505050565b5f604082019050614f445f830186613d69565b8181036020830152614f57818486614809565b9050949350505050565b5f81519050614f6f81613e9a565b92915050565b5f60208284031215614f8a57614f89613663565b5b5f614f9784828501614f61565b91505092915050565b5f606082019050614fb35f830187613d69565b614fc06020830186613eef565b8181036040830152614fd3818486614809565b905095945050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f615012601c83614b51565b915061501d82614fde565b601c82019050919050565b5f61503282615006565b915061503e8284614bae565b60208201915081905092915050565b5f6040820190508181035f8301526150658186614729565b9050818103602083015261507a818486614809565b9050949350505050565b5f6020828403121561509957615098613663565b5b5f6150a684828501614c29565b91505092915050565b7fff000000000000000000000000000000000000000000000000000000000000005f82015250565b5f6150e3600183614b51565b91506150ee826150af565b600182019050919050565b5f8160601b9050919050565b5f61510f826150f9565b9050919050565b5f61512082615105565b9050919050565b6151386151338261368a565b615116565b82525050565b5f615148826150d7565b91506151548286615127565b6014820191506151648285614bae565b6020820191506151748284614bae565b602082019150819050949350505050565b5f61518f826144e7565b6151998185613606565b93506151a9818560208601614501565b80840191505092915050565b5f6151c08285615185565b91506151cc8284614bae565b6020820191508190509392505050565b7f53657175656e6365207369676e65723a0a0000000000000000000000000000005f82015250565b5f615210601183614b51565b915061521b826151dc565b601182019050919050565b5f819050919050565b61524061523b8261381e565b615226565b82525050565b5f61525082615204565b915061525c8285615127565b60148201915061526c828461522f565b6020820191508190509392505050565b7f53657175656e6365206e657374656420636f6e6669673a0a00000000000000005f82015250565b5f6152b0601883614b51565b91506152bb8261527c565b601882019050919050565b5f6152d0826152a4565b91506152dc8286614bae565b6020820191506152ec828561522f565b6020820191506152fc828461522f565b602082019150819050949350505050565b7f53657175656e636520737461746963206469676573743a0a00000000000000005f82015250565b5f615341601883614b51565b915061534c8261530d565b601882019050919050565b5f61536182615335565b915061536d8284614bae565b60208201915081905092915050565b7f53657175656e636520616e792061646472657373207375626469676573743a0a5f82015250565b5f6153b0602083614b51565b91506153bb8261537c565b602082019050919050565b5f6153d0826153a4565b91506153dc8284614bae565b60208201915081905092915050565b7f53657175656e63652073617069656e7420636f6e6669673a0a000000000000005f82015250565b5f61541f601983614b51565b915061542a826153eb565b601982019050919050565b5f61543f82615413565b915061544b8286615127565b60148201915061545b828561522f565b60208201915061546b8284614bae565b602082019150819050949350505050565b5f610100820190506154905f83018b613d69565b61549d602083018a613eef565b6154aa60408301896140a7565b6154b76060830188613d69565b6154c460808301876140a7565b6154d160a08301866140f6565b6154de60c08301856140f6565b6154eb60e08301846140a7565b999850505050505050505056fea2646970667358221220612dd00679c3299f77a9296bc1c3648da9d4855701ab79906c1b3b8db37bfe6f64736f6c634300081c0033",
}

// WalletStage1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletStage1MetaData.ABI instead.
var WalletStage1ABI = WalletStage1MetaData.ABI

// WalletStage1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletStage1MetaData.Bin instead.
var WalletStage1Bin = WalletStage1MetaData.Bin

// DeployWalletStage1 deploys a new Ethereum contract, binding an instance of WalletStage1 to it.
func DeployWalletStage1(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *WalletStage1, error) {
	parsed, err := WalletStage1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletStage1Bin), backend, _factory)
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
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1Caller) ReadHook(opts *bind.CallOpts, signature [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletStage1.contract.Call(opts, &out, "readHook", signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1Session) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, signature)
}

// ReadHook is a free data retrieval call binding the contract method 0x1a9b2337.
//
// Solidity: function readHook(bytes4 signature) view returns(address)
func (_WalletStage1 *WalletStage1CallerSession) ReadHook(signature [4]byte) (common.Address, error) {
	return _WalletStage1.Contract.ReadHook(&_WalletStage1.CallOpts, signature)
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

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Transactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1Session) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage1.Contract.AddHook(&_WalletStage1.TransactOpts, signature, implementation)
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

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1Transactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1Session) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage1 *WalletStage1TransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.RemoveHook(&_WalletStage1.TransactOpts, signature)
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

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1Transactor) TokenReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.contract.Transact(opts, "tokenReceived", arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1Session) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.TransactOpts, arg0, arg1, arg2)
}

// TokenReceived is a paid mutator transaction binding the contract method 0x8943ec02.
//
// Solidity: function tokenReceived(address , uint256 , bytes ) returns()
func (_WalletStage1 *WalletStage1TransactorSession) TokenReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _WalletStage1.Contract.TokenReceived(&_WalletStage1.TransactOpts, arg0, arg1, arg2)
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
	Signature      [4]byte
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDefinedHook is a free log retrieval operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
func (_WalletStage1 *WalletStage1Filterer) FilterDefinedHook(opts *bind.FilterOpts) (*WalletStage1DefinedHookIterator, error) {

	logs, sub, err := _WalletStage1.contract.FilterLogs(opts, "DefinedHook")
	if err != nil {
		return nil, err
	}
	return &WalletStage1DefinedHookIterator{contract: _WalletStage1.contract, event: "DefinedHook", logs: logs, sub: sub}, nil
}

// WatchDefinedHook is a free log subscription operation binding the contract event 0x0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed1.
//
// Solidity: event DefinedHook(bytes4 signature, address implementation)
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
// Solidity: event DefinedHook(bytes4 signature, address implementation)
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
