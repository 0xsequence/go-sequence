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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_current\",\"type\":\"uint256\"}],\"name\":\"BadNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"HookDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ImageHashIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidERC1271Signature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"InvalidKind\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSapientSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_flag\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureFlag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"InvalidSignatureWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"InvalidStaticSignatureExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_expectedCaller\",\"type\":\"address\"}],\"name\":\"InvalidStaticSignatureWrongCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"LowWeightChainedSignature\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLeft\",\"type\":\"uint256\"}],\"name\":\"NotEnoughGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"OnlySelf\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"Reverted\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"}],\"internalType\":\"structSnapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"}],\"name\":\"UnusedSnapshot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nextCheckpoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_checkpoint\",\"type\":\"uint256\"}],\"name\":\"WrongChainedCheckpointOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallAborted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_returnData\",\"type\":\"bytes\"}],\"name\":\"CallFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_opHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CallSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"DefinedHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newImageHash\",\"type\":\"bytes32\"}],\"name\":\"ImageHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ImplementationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"StaticSignatureSet\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getStaticSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"readHook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_space\",\"type\":\"uint256\"}],\"name\":\"readNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverPartialSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isValidImage\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"checkpoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"noChainId\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"space\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"imageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parentWallets\",\"type\":\"address[]\"}],\"internalType\":\"structPayload.Decoded\",\"name\":\"_payload\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"recoverSapientSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"signature\",\"type\":\"bytes4\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"selfExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_timestamp\",\"type\":\"uint96\"}],\"name\":\"setStaticSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokenReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_imageHash\",\"type\":\"bytes32\"}],\"name\":\"updateImageHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"updateImplementation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405260043610610117575f3560e01c80636ea445771161009f578063ad55366b11610063578063ad55366b14610475578063b93ea7ad146104b6578063bc197c81146104d2578063f23a6e611461050e578063f727ef1c1461054a5761011e565b80636ea445771461038e5780638943ec02146103aa5780638c3f5563146103d257806392dcb3fc1461040e578063aaf10f421461044b5761011e565b80631a9b2337116100e65780631a9b2337146102c85780631f6a1eb91461030457806329561426146103205780634fcf3eca1461034857806351605d80146103645761011e565b8063025b22bc146101f857806313792a4a14610214578063150b7a02146102505780631626ba7e1461028c5761011e565b3661011e57005b60045f369050106101f6575f61013f5f369061013a9190613487565b610572565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146101f4575f5f8273ffffffffffffffffffffffffffffffffffffffff165f3660405161019d929190613521565b5f60405180830381855af49150503d805f81146101d5576040519150601f19603f3d011682016040523d82523d5f602084013e6101da565b606091505b5091509150816101ec57805160208201fd5b805160208201f35b505b005b610212600480360381019061020d91906135a4565b6105c7565b005b34801561021f575f5ffd5b5061023a60048036038101906102359190613bcf565b610643565b6040516102479190613c57565b60405180910390f35b34801561025b575f5ffd5b5061027660048036038101906102719190613c70565b6107f2565b6040516102839190613d03565b60405180910390f35b348015610297575f5ffd5b506102b260048036038101906102ad9190613d1c565b610806565b6040516102bf9190613d03565b60405180910390f35b3480156102d3575f5ffd5b506102ee60048036038101906102e99190613da3565b610848565b6040516102fb9190613ddd565b60405180910390f35b61031e60048036038101906103199190613df6565b610859565b005b34801561032b575f5ffd5b5061034660048036038101906103419190613e74565b6108e8565b005b610362600480360381019061035d9190613da3565b610964565b005b34801561036f575f5ffd5b50610378610a59565b6040516103859190613c57565b60405180910390f35b6103a860048036038101906103a39190613e9f565b610a8a565b005b3480156103b5575f5ffd5b506103d060048036038101906103cb9190613eea565b610b29565b005b3480156103dd575f5ffd5b506103f860048036038101906103f39190613f5b565b610b2f565b6040516104059190613f95565b60405180910390f35b348015610419575f5ffd5b50610434600480360381019061042f9190613e74565b610b67565b604051610442929190613fae565b60405180910390f35b348015610456575f5ffd5b5061045f610b7b565b60405161046c9190613ddd565b60405180910390f35b348015610480575f5ffd5b5061049b60048036038101906104969190613bcf565b610b89565b6040516104ad96959493929190613fe4565b60405180910390f35b6104d060048036038101906104cb9190614043565b610bc7565b005b3480156104dd575f5ffd5b506104f860048036038101906104f391906140d6565b610cbd565b6040516105059190613d03565b60405180910390f35b348015610519575f5ffd5b50610534600480360381019061052f91906141ad565b610cd4565b6040516105419190613d03565b60405180910390f35b348015610555575f5ffd5b50610570600480360381019061056b9190614284565b610ce9565b005b5f6105be7fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916610db2565b5f1c9050919050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461063757336040517fa19dbf0000000000000000000000000000000000000000000000000000000000815260040161062e9190613ddd565b60405180910390fd5b61064081610dea565b50565b5f5f6001856101000151516106589190614301565b67ffffffffffffffff811115610671576106706135e3565b5b60405190808252806020026020018201604052801561069f5781602001602082028036833780820191505090505b5090505f5f90505b8561010001515181101561072f5785610100015181815181106106cd576106cc614334565b5b60200260200101518282815181106106e8576106e7614334565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506106a7565b503381866101000151518151811061074a57610749614334565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808561010001819052505f61079a868686610e2d565b509050806107e3578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016107da93929190614714565b60405180910390fd5b60015f1b925050509392505050565b5f63150b7a0260e01b905095945050505050565b5f5f6108118561101d565b90505f61081f828686610e2d565b50905080610834575f60e01b92505050610841565b6320c13b0b60e01b925050505b9392505050565b5f61085282610572565b9050919050565b5f5a90505f6108688686611046565b905061087c816060015182608001516114e8565b5f5f610889838787610e2d565b91509150816108d3578286866040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016108ca93929190614714565b60405180910390fd5b6108de84828561158c565b5050505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461095857336040517fa19dbf0000000000000000000000000000000000000000000000000000000000815260040161094f9190613ddd565b60405180910390fd5b610961816118bc565b50565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d457336040517fa19dbf000000000000000000000000000000000000000000000000000000000081526004016109cb9190613ddd565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff166109f482610572565b73ffffffffffffffffffffffffffffffffffffffff1603610a4c57806040517f1c3812cc000000000000000000000000000000000000000000000000000000008152600401610a439190613d03565b60405180910390fd5b610a56815f61195d565b50565b5f610a857fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b6119fe565b905090565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610afa57336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610af19190613ddd565b60405180910390fd5b5f5a90505f610b098484611046565b90505f610b1582611a08565b9050610b2283828461158c565b5050505050565b50505050565b5f610b5e7f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b610db2565b5f1c9050919050565b5f5f610b7283611a58565b91509150915091565b5f610b84611aa9565b905090565b5f5f5f5f5f5f610b9c8989895f5f611ab1565b809550819650829750839950849a505050505050610bb983611dee565b935093975093979195509350565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c3757336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610c2e9190613ddd565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff16610c5783610572565b73ffffffffffffffffffffffffffffffffffffffff1614610caf57816040517f5b4d6d6a000000000000000000000000000000000000000000000000000000008152600401610ca69190613d03565b60405180910390fd5b610cb9828261195d565b5050565b5f63bc197c8160e01b905098975050505050505050565b5f63f23a6e6160e01b90509695505050505050565b3073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d5957336040517fa19dbf00000000000000000000000000000000000000000000000000000000008152600401610d509190613ddd565b60405180910390fd5b610d728383836bffffffffffffffffffffffff16611dff565b7febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b1838383604051610da59392919061475a565b60405180910390a1505050565b5f5f8383604051602001610dc792919061478f565b604051602081830303815290604052805190602001209050805491505092915050565b610df381611e5d565b7f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca0381604051610e229190613ddd565b60405180910390a150565b5f5f5f84845f818110610e4357610e42614334565b5b9050013560f81c60f81b9050608060f81b608060f81b82167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191603610f9c57610e8b86611a08565b91505f5f610e9884611a58565b91509150428111610ee25783816040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600401610ed99291906147b6565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610f4a57503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15610f90578333836040517f8945c313000000000000000000000000000000000000000000000000000000008152600401610f87939291906147dd565b60405180910390fd5b60019450505050611015565b5f5f5f610fac8989895f5f611ab1565b905080985081945082955083965050505050828210156110055782826040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600401610ffc929190614812565b60405180910390fd5b61100e81611dee565b9550505050505b935093915050565b611025613392565b6003815f019060ff16908160ff1681525050818160e0018181525050919050565b61104e613392565b5f815f019060ff16908160ff16815250505f5f61106b8585611e63565b915060ff16915060018083160361108b575f8360600181815250506110c7565b6110a0818686611e799290919263ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff169150846060018193508281525050505b5f6007600184901c1690505f811115611104576110f682828888611eaa9190939291909392919063ffffffff16565b856080018194508281525050505b5f6010808516036111185760019050611170565b60208085160361114b57611137838888611ed79290919263ffffffff16565b8161ffff169150809450819250505061116f565b611160838888611ef69290919263ffffffff16565b8160ff16915080945081925050505b5b8067ffffffffffffffff81111561118a576111896135e3565b5b6040519080825280602002602001820160405280156111c357816020015b6111b06133dd565b8152602001906001900390816111a85790505b5085604001819052505f5f90505b818110156114dd575f6111ef858a8a611ef69290919263ffffffff16565b8096508192505050600180821660ff160361125e57308760400151838151811061121c5761121b614334565b5b60200260200101515f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506112ca565b611273858a8a611f119290919263ffffffff16565b8860400151848151811061128a57611289614334565b5b60200260200101515f018197508273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050505b600280821660ff1603611318576112ec858a8a611f429290919263ffffffff16565b8860400151848151811061130357611302614334565b5b60200260200101516020018197508281525050505b600480821660ff16036113e0575f61133b868b8b611f589290919263ffffffff16565b8162ffffff169150809750819250505089898790838961135b9190614301565b9261136893929190614841565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050886040015184815181106113c1576113c0614334565b5b60200260200101516040018190525080866113dc9190614301565b9550505b600880821660ff160361142e57611402858a8a611f429290919263ffffffff16565b8860400151848151811061141957611418614334565b5b60200260200101516060018197508281525050505b601080821660ff16148760400151838151811061144e5761144d614334565b5b60200260200101516080019015159081151581525050602080821660ff16148760400151838151811061148457611483614334565b5b602002602001015160a0019015159081151581525050600660c0821660ff16901c60ff16876040015183815181106114bf576114be614334565b5b602002602001015160c00181815250505080806001019150506111d1565b505050505092915050565b5f6114f283610b2f565b905081811461153c578282826040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004016115339392919061487b565b60405180910390fd5b5f60018301905061154d8482611f78565b7f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881848260405161157e929190614812565b60405180910390a150505050565b5f5f90505f82604001515190505f5f90505b818110156118b4575f846040015182815181106115be576115bd614334565b5b602002602001015190508060a0015180156115d7575083155b1561161b577f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b868360405161160d9291906147b6565b60405180910390a1506118a7565b5f93505f816060015190505f81141580156116355750805a105b1561167b5785835a6040517f21395274000000000000000000000000000000000000000000000000000000008152600401611672939291906148b0565b60405180910390fd5b5f82608001511561173057611729835f01515f841461169a578361169c565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016116c796959493929190614924565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611fad565b9050611758565b611755835f015184602001515f8514611749578461174b565b5a5b8660400151611fc2565b90505b8061186a575f60ff168360c00151036117b957600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d888561179a611fd9565b6040516117a99392919061498a565b60405180910390a15050506118a7565b600160ff168360c00151036118105786846117d2611fd9565b6040517f7f6b0bb1000000000000000000000000000000000000000000000000000000008152600401611807939291906149c6565b60405180910390fd5b600260ff168360c0015103611869577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b888561184a611fd9565b6040516118599392919061498a565b60405180910390a15050506118b4565b5b7f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a888560405161189b9291906147b6565b60405180910390a15050505b808060010191505061159e565b505050505050565b5f5f1b81036118f7576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119237fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b82611ff7565b7f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa816040516119529190613c57565b60405180910390a150565b6119c17fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1205f1b837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168373ffffffffffffffffffffffffffffffffffffffff165f1b611ffe565b7f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed182826040516119f2929190614a09565b60405180910390a15050565b5f81549050919050565b5f5f611a18836020015130612033565b90505f611a24846120d7565b90508181604051602001611a39929190614aa4565b6040516020818303038152906040528051906020012092505050919050565b5f5f5f611a877fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b85610db2565b5f1c9050606081901c816bffffffffffffffffffffffff169250925050915091565b5f3054905090565b5f5f5f5f5f5f5f611ac28b8b611e63565b915060ff169150611ad161342e565b6040808416148015611b0e57505f73ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff16145b15611c4a57611b28828d8d611f119290919263ffffffff16565b809350819a50505089611c49575f611b4b838e8e611f589290919263ffffffff16565b8162ffffff16915080945081925050505f8d8d85908487611b6c9190614301565b92611b7993929190614841565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090508a73ffffffffffffffffffffffffffffffffffffffff1663ccce3bc830836040518363ffffffff1660e01b8152600401611bf8929190614ada565b6040805180830381865afa158015611c12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c369190614b7d565b92508184611c449190614301565b935050505b5b600180841603611c8357611c718d8a838f8f87908092611c6c93929190614841565b61230b565b97509750975097509750505050611de1565b6002808416148d60200190151590811515815250505f6002601c8516901c9050611cbf83828f8f611eaa9190939291909392919063ffffffff16565b8094508197505050505f6001600560208616901c611cdd9190614301565b9050611cfb83828f8f611eaa9190939291909392919063ffffffff16565b809450819a50505050611d0d8d611a08565b9350611d2b8d858e8e86908092611d2693929190614841565b61256d565b8097508198505050611d3f86895f1b6130e1565b9550611d4d86865f1b6130e1565b9550611d71868a73ffffffffffffffffffffffffffffffffffffffff165f1b6130e1565b95505f5f1b815f015114158015611d8b575085815f015114155b8015611d9b575080602001518511155b15611ddd57806040517fccbb534f000000000000000000000000000000000000000000000000000000008152600401611dd49190614bd5565b60405180910390fd5b5050505b9550955095509550959050565b5f611df8826130f5565b9050919050565b611e587fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e865f1b846bffffffffffffffffffffffff841660608673ffffffffffffffffffffffffffffffffffffffff16901b175f1b611ffe565b505050565b80305550565b5f5f83358060f81c925060019150509250929050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f858401356008840261010003600180866008021b0382821c1693508486019250505094509492505050565b5f5f8483013561ffff8160f01c16925060028401915050935093915050565b5f5f848301358060f81c925060018401915050935093915050565b5f5f8483013573ffffffffffffffffffffffffffffffffffffffff8160601c16925060148401915050935093915050565b5f5f848301359150602083019050935093915050565b5f5f8483013562ffffff8160e81c16925060038401915050935093915050565b611fa97f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e5f1b835f1b835f1b611ffe565b5050565b5f5f5f8351602085018787f490509392505050565b5f5f5f835160208501878988f19050949350505050565b60603d604051915060208201818101604052818352815f823e505090565b8082555050565b5f838360405160200161201292919061478f565b60405160208183030381529060405280519060200120905081815550505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de856120a257466120a4565b5f5b856040516020016120b9959493929190614bee565b60405160208183030381529060405280519060200120905092915050565b5f5f8261010001516040516020016120ef9190614ccb565b6040516020818303038152906040528051906020012090505f60ff16835f015160ff1603612187575f6121258460400151613138565b90507f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a2818560600151866080015185604051602001612168959493929190614ce1565b6040516020818303038152906040528051906020012092505050612306565b600160ff16835f015160ff16036121f6577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360a0015180519060200120826040516020016121d893929190614d32565b60405160208183030381529060405280519060200120915050612306565b600260ff16835f015160ff160361225e577f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e48360c001518260405160200161224093929190614d32565b60405160208183030381529060405280519060200120915050612306565b600360ff16835f015160ff16036122c6577fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d4668360e00151826040516020016122a893929190614d32565b60405160208183030381529060405280519060200120915050612306565b825f01516040517f048183200000000000000000000000000000000000000000000000000000000081526004016122fd9190614d76565b60405180910390fd5b919050565b5f5f5f5f5f612318613392565b6002815f019060ff16908160ff16815250505f5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90505b89899050821015612502575f5f612372848d8d611f589290919263ffffffff16565b8162ffffff1691508095508192505050838161238e9190614301565b9150505f8b8b905082146123a2575f6123a4565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8303612404576123eb8f8d8d879086926123e393929190614841565b600185611ab1565b809a50819b50829c50839d50849e505050505050612434565b612422858d8d8790869261241a93929190614841565b600185611ab1565b50809a50819b50829c50839d50505050505b8989101561248f578b8b8590849261244e93929190614841565b8b8b6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016124869493929190614d8f565b60405180910390fd5b819350878d5f0151036124a8575f5f1b8d5f0181815250505b8287106124ee5786836040517f37daf62b0000000000000000000000000000000000000000000000000000000081526004016124e5929190614812565b60405180910390fd5b878560c00181815250508692505050612350565b5f5f1b8b5f01511415801561251b57508a602001518511155b1561255d578a6040517fccbb534f0000000000000000000000000000000000000000000000000000000081526004016125549190614bd5565b60405180910390fd5b5050509550955095509550959050565b5f5f5f5b848490508110156130d7575f612592828787611ef69290919263ffffffff16565b8160ff16915080935081925050505f600460f08316901c90505f81036126ea575f600f831690505f8160ff16036125e1576125d8848989611ef69290919263ffffffff16565b80955081925050505b5f5f6125f8868b8b6131b29290919263ffffffff16565b8097508193505050612615868b8b6131b29290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f8388866040515f815260200160405260405161267b9493929190614dcd565b6020604051602081039080840390855afa15801561269b573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f6126bd828960ff166131c8565b90505f5f1b8c036126ce57806126d9565b6126d88c826130e1565b5b9b5050505050505050505050612571565b60018103612775575f600f831690505f8160ff160361272157612718848989611ef69290919263ffffffff16565b80955081925050505b5f612737858a8a611f119290919263ffffffff16565b80965081925050505f61274d828460ff166131c8565b90505f5f1b870361275e5780612769565b61276887826130e1565b5b96505050505050612571565b60028103612974575f6003831690505f8160ff16036127ac576127a3848989611ef69290919263ffffffff16565b80955081925050505b5f6127c2858a8a611f119290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f6127f887838d8d611eaa9190939291909392919063ffffffff16565b80985081925050505f81880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261285c93929190614841565b6040518463ffffffff1660e01b815260040161287a93929190614e10565b602060405180830381865afa158015612895573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128b99190614e54565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614612930578c848d8d8b9085926128f193929190614841565b6040517fb2fed7ae0000000000000000000000000000000000000000000000000000000081526004016129279493929190614e7f565b60405180910390fd5b8097508460ff168a0199505f612949858760ff166131c8565b90505f5f1b8a0361295a5780612965565b6129648a826130e1565b5b99505050505050505050612571565b600381036129be575f6129928489896131b29290919263ffffffff16565b80955081925050505f5f1b85036129a957806129b4565b6129b385826130e1565b5b9450505050612571565b60048103612a3d575f600f831660ff1690505f6129ed85838b8b611eaa9190939291909392919063ffffffff16565b80965081925050505f81860190505f5f612a198e8e8e8e8c908892612a1493929190614841565b61256d565b91509150829750818a019950612a2f89826130e1565b985050505050505050612571565b60068103612b4d575f6002600c841660ff16901c60ff1690505f8103612a8157612a72848989611ef69290919263ffffffff16565b8160ff16915080955081925050505b5f6003841660ff1690505f8103612ab757612aa7858a8a611ed79290919263ffffffff16565b8161ffff16915080965081925050505b5f612acd868b8b611f589290919263ffffffff16565b8162ffffff16915080975081925050505f81870190505f5f612b018f8f8f8f8d908892612afc93929190614841565b61256d565b91509150829850848210612b1557858b019a505b5f612b218287896131fa565b90505f5f1b8b03612b325780612b3d565b612b3c8b826130e1565b5b9a50505050505050505050612571565b60058103612bcf575f612b6b8489896131b29290919263ffffffff16565b8095508192505050888103612b9e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b5f612ba88261322f565b90505f5f1b8603612bb95780612bc4565b612bc386826130e1565b5b955050505050612571565b60078103612d35575f600f831690505f8160ff1603612c0657612bfd848989611ef69290919263ffffffff16565b80955081925050505b5f5f612c1d868b8b6131b29290919263ffffffff16565b8097508193505050612c3a868b8b6131b29290919263ffffffff16565b80975081925050505f60ff82901c5f1c90505f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff835f1c165f1b90505f601b830190505f60018f604051602001612c919190614f07565b604051602081830303815290604052805190602001208388866040515f8152602001604052604051612cc69493929190614dcd565b6020604051602081039080840390855afa158015612ce6573d5f5f3e3d5ffd5b5050506020604051035190508660ff168c019b505f612d08828960ff166131c8565b90505f5f1b8c03612d195780612d24565b612d238c826130e1565b5b9b5050505050505050505050612571565b60088103612dce575f612d538489896131b29290919263ffffffff16565b80955081925050505f612d6f5f8c61325e90919063ffffffff16565b9050808203612d9c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b5f612da6836132af565b90505f5f1b8703612db75780612dc2565b612dc187826130e1565b5b96505050505050612571565b60098103612f34575f6003831690505f8160ff1603612e0557612dfc848989611ef69290919263ffffffff16565b80955081925050505b5f612e1b858a8a611f119290919263ffffffff16565b80965081925050505f5f6002600c871660ff16901c60ff169050612e5187828d8d611eaa9190939291909392919063ffffffff16565b8098508193505050505f81870190505f8373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612e9093929190614841565b6040518463ffffffff1660e01b8152600401612eae93929190614f2c565b602060405180830381865afa158015612ec9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612eed9190614f63565b90508197508460ff168a0199505f612f09858760ff16846132de565b90505f5f1b8a03612f1a5780612f25565b612f248a826130e1565b5b99505050505050505050612571565b600a810361309a575f6003831690505f8160ff1603612f6b57612f62848989611ef69290919263ffffffff16565b80955081925050505b5f612f81858a8a611f119290919263ffffffff16565b80965081925050505f6002600c861660ff16901c60ff1690505f612fb787838d8d611eaa9190939291909392919063ffffffff16565b80985081925050505f81880190505f8473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612ff593929190614841565b6040518463ffffffff1660e01b815260040161301393929190614e10565b602060405180830381865afa15801561302e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130529190614f63565b90508198508560ff168b019a505f61306e868860ff16846132de565b90505f5f1b8b0361307f578061308a565b6130898b826130e1565b5b9a50505050505050505050612571565b806040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004016130ce9190613f95565b60405180910390fd5b5094509492505050565b5f825f528160205260405f20905092915050565b5f5f5f1b8214158015613131575061312e7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85f1b6119fe565b82145b9050919050565b5f60605f5f90505b83518110156131a1575f61316d8583815181106131605761315f614334565b5b6020026020010151613313565b90508281604051602001613182929190614fbe565b6040516020818303038152906040529250508080600101915050613140565b508080519060200120915050919050565b5f5f848301359150602083019050935093915050565b5f82826040516020016131dc929190615094565b60405160208183030381529060405280519060200120905092915050565b5f83838360405160200161321093929190615114565b6040516020818303038152906040528051906020012090509392505050565b5f8160405160200161324191906151a5565b604051602081830303815290604052805190602001209050919050565b5f5f61326e846020015184612033565b90505f61327a856120d7565b9050818160405160200161328f929190614aa4565b604051602081830303815290604052805190602001209250505092915050565b5f816040516020016132c19190615214565b604051602081830303815290604052805190602001209050919050565b5f8383836040516020016132f493929190615283565b6040516020818303038152906040528051906020012090509392505050565b5f7f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437825f01518360200151846040015180519060200120856060015186608001518760a001518860c001516040516020016133759897969594939291906152ca565b604051602081830303815290604052805190602001209050919050565b6040518061012001604052805f60ff1681526020015f15158152602001606081526020015f81526020015f8152602001606081526020015f81526020015f8152602001606081525090565b6040518060e001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f8152602001606081526020015f81526020015f151581526020015f151581526020015f81525090565b60405180604001604052805f81526020015f81525090565b5f82905092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b5f82821b905092915050565b5f6134928383613446565b8261349d8135613450565b925060048210156134dd576134d87fffffffff000000000000000000000000000000000000000000000000000000008360040360080261347b565b831692505b505092915050565b5f81905092915050565b828183375f83830152505050565b5f61350883856134e5565b93506135158385846134ef565b82840190509392505050565b5f61352d8284866134fd565b91508190509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6135738261354a565b9050919050565b61358381613569565b811461358d575f5ffd5b50565b5f8135905061359e8161357a565b92915050565b5f602082840312156135b9576135b8613542565b5b5f6135c684828501613590565b91505092915050565b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b613619826135d3565b810181811067ffffffffffffffff82111715613638576136376135e3565b5b80604052505050565b5f61364a613539565b90506136568282613610565b919050565b5f5ffd5b5f60ff82169050919050565b6136748161365f565b811461367e575f5ffd5b50565b5f8135905061368f8161366b565b92915050565b5f8115159050919050565b6136a981613695565b81146136b3575f5ffd5b50565b5f813590506136c4816136a0565b92915050565b5f5ffd5b5f67ffffffffffffffff8211156136e8576136e76135e3565b5b602082029050602081019050919050565b5f5ffd5b5f819050919050565b61370f816136fd565b8114613719575f5ffd5b50565b5f8135905061372a81613706565b92915050565b5f5ffd5b5f67ffffffffffffffff82111561374e5761374d6135e3565b5b613757826135d3565b9050602081019050919050565b5f61377661377184613734565b613641565b90508281526020810184848401111561379257613791613730565b5b61379d8482856134ef565b509392505050565b5f82601f8301126137b9576137b86136ca565b5b81356137c9848260208601613764565b91505092915050565b5f60e082840312156137e7576137e66135cf565b5b6137f160e0613641565b90505f61380084828501613590565b5f8301525060206138138482850161371c565b602083015250604082013567ffffffffffffffff8111156138375761383661365b565b5b613843848285016137a5565b60408301525060606138578482850161371c565b606083015250608061386b848285016136b6565b60808301525060a061387f848285016136b6565b60a08301525060c06138938482850161371c565b60c08301525092915050565b5f6138b16138ac846136ce565b613641565b905080838252602082019050602084028301858111156138d4576138d36136f9565b5b835b8181101561391b57803567ffffffffffffffff8111156138f9576138f86136ca565b5b80860161390689826137d2565b855260208501945050506020810190506138d6565b5050509392505050565b5f82601f830112613939576139386136ca565b5b813561394984826020860161389f565b91505092915050565b5f819050919050565b61396481613952565b811461396e575f5ffd5b50565b5f8135905061397f8161395b565b92915050565b5f67ffffffffffffffff82111561399f5761399e6135e3565b5b602082029050602081019050919050565b5f6139c26139bd84613985565b613641565b905080838252602082019050602084028301858111156139e5576139e46136f9565b5b835b81811015613a0e57806139fa8882613590565b8452602084019350506020810190506139e7565b5050509392505050565b5f82601f830112613a2c57613a2b6136ca565b5b8135613a3c8482602086016139b0565b91505092915050565b5f6101208284031215613a5b57613a5a6135cf565b5b613a66610120613641565b90505f613a7584828501613681565b5f830152506020613a88848285016136b6565b602083015250604082013567ffffffffffffffff811115613aac57613aab61365b565b5b613ab884828501613925565b6040830152506060613acc8482850161371c565b6060830152506080613ae08482850161371c565b60808301525060a082013567ffffffffffffffff811115613b0457613b0361365b565b5b613b10848285016137a5565b60a08301525060c0613b2484828501613971565b60c08301525060e0613b3884828501613971565b60e08301525061010082013567ffffffffffffffff811115613b5d57613b5c61365b565b5b613b6984828501613a18565b6101008301525092915050565b5f5ffd5b5f5f83601f840112613b8f57613b8e6136ca565b5b8235905067ffffffffffffffff811115613bac57613bab613b76565b5b602083019150836001820283011115613bc857613bc76136f9565b5b9250929050565b5f5f5f60408486031215613be657613be5613542565b5b5f84013567ffffffffffffffff811115613c0357613c02613546565b5b613c0f86828701613a45565b935050602084013567ffffffffffffffff811115613c3057613c2f613546565b5b613c3c86828701613b7a565b92509250509250925092565b613c5181613952565b82525050565b5f602082019050613c6a5f830184613c48565b92915050565b5f5f5f5f5f60808688031215613c8957613c88613542565b5b5f613c9688828901613590565b9550506020613ca788828901613590565b9450506040613cb88882890161371c565b935050606086013567ffffffffffffffff811115613cd957613cd8613546565b5b613ce588828901613b7a565b92509250509295509295909350565b613cfd81613450565b82525050565b5f602082019050613d165f830184613cf4565b92915050565b5f5f5f60408486031215613d3357613d32613542565b5b5f613d4086828701613971565b935050602084013567ffffffffffffffff811115613d6157613d60613546565b5b613d6d86828701613b7a565b92509250509250925092565b613d8281613450565b8114613d8c575f5ffd5b50565b5f81359050613d9d81613d79565b92915050565b5f60208284031215613db857613db7613542565b5b5f613dc584828501613d8f565b91505092915050565b613dd781613569565b82525050565b5f602082019050613df05f830184613dce565b92915050565b5f5f5f5f60408587031215613e0e57613e0d613542565b5b5f85013567ffffffffffffffff811115613e2b57613e2a613546565b5b613e3787828801613b7a565b9450945050602085013567ffffffffffffffff811115613e5a57613e59613546565b5b613e6687828801613b7a565b925092505092959194509250565b5f60208284031215613e8957613e88613542565b5b5f613e9684828501613971565b91505092915050565b5f5f60208385031215613eb557613eb4613542565b5b5f83013567ffffffffffffffff811115613ed257613ed1613546565b5b613ede85828601613b7a565b92509250509250929050565b5f5f5f5f60608587031215613f0257613f01613542565b5b5f613f0f87828801613590565b9450506020613f208782880161371c565b935050604085013567ffffffffffffffff811115613f4157613f40613546565b5b613f4d87828801613b7a565b925092505092959194509250565b5f60208284031215613f7057613f6f613542565b5b5f613f7d8482850161371c565b91505092915050565b613f8f816136fd565b82525050565b5f602082019050613fa85f830184613f86565b92915050565b5f604082019050613fc15f830185613dce565b613fce6020830184613f86565b9392505050565b613fde81613695565b82525050565b5f60c082019050613ff75f830189613f86565b6140046020830188613f86565b6140116040830187613fd5565b61401e6060830186613c48565b61402b6080830185613f86565b61403860a0830184613c48565b979650505050505050565b5f5f6040838503121561405957614058613542565b5b5f61406685828601613d8f565b925050602061407785828601613590565b9150509250929050565b5f5f83601f840112614096576140956136ca565b5b8235905067ffffffffffffffff8111156140b3576140b2613b76565b5b6020830191508360208202830111156140cf576140ce6136f9565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b0312156140f2576140f1613542565b5b5f6140ff8b828c01613590565b98505060206141108b828c01613590565b975050604089013567ffffffffffffffff81111561413157614130613546565b5b61413d8b828c01614081565b9650965050606089013567ffffffffffffffff8111156141605761415f613546565b5b61416c8b828c01614081565b9450945050608089013567ffffffffffffffff81111561418f5761418e613546565b5b61419b8b828c01613b7a565b92509250509295985092959890939650565b5f5f5f5f5f5f60a087890312156141c7576141c6613542565b5b5f6141d489828a01613590565b96505060206141e589828a01613590565b95505060406141f689828a0161371c565b945050606061420789828a0161371c565b935050608087013567ffffffffffffffff81111561422857614227613546565b5b61423489828a01613b7a565b92509250509295509295509295565b5f6bffffffffffffffffffffffff82169050919050565b61426381614243565b811461426d575f5ffd5b50565b5f8135905061427e8161425a565b92915050565b5f5f5f6060848603121561429b5761429a613542565b5b5f6142a886828701613971565b93505060206142b986828701613590565b92505060406142ca86828701614270565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61430b826136fd565b9150614316836136fd565b925082820190508082111561432e5761432d6142d4565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61436a8161365f565b82525050565b61437981613695565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6143b181613569565b82525050565b6143c0816136fd565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6143f8826143c6565b61440281856143d0565b93506144128185602086016143e0565b61441b816135d3565b840191505092915050565b5f60e083015f83015161443b5f8601826143a8565b50602083015161444e60208601826143b7565b506040830151848203604086015261446682826143ee565b915050606083015161447b60608601826143b7565b50608083015161448e6080860182614370565b5060a08301516144a160a0860182614370565b5060c08301516144b460c08601826143b7565b508091505092915050565b5f6144ca8383614426565b905092915050565b5f602082019050919050565b5f6144e88261437f565b6144f28185614389565b93508360208202850161450485614399565b805f5b8581101561453f578484038952815161452085826144bf565b945061452b836144d2565b925060208a01995050600181019050614507565b50829750879550505050505092915050565b61455a81613952565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f61459483836143a8565b60208301905092915050565b5f602082019050919050565b5f6145b682614560565b6145c0818561456a565b93506145cb8361457a565b805f5b838110156145fb5781516145e28882614589565b97506145ed836145a0565b9250506001810190506145ce565b5085935050505092915050565b5f61012083015f83015161461e5f860182614361565b5060208301516146316020860182614370565b506040830151848203604086015261464982826144de565b915050606083015161465e60608601826143b7565b50608083015161467160808601826143b7565b5060a083015184820360a086015261468982826143ee565b91505060c083015161469e60c0860182614551565b5060e08301516146b160e0860182614551565b506101008301518482036101008601526146cb82826145ac565b9150508091505092915050565b5f82825260208201905092915050565b5f6146f383856146d8565b93506147008385846134ef565b614709836135d3565b840190509392505050565b5f6040820190508181035f83015261472c8186614608565b905081810360208301526147418184866146e8565b9050949350505050565b61475481614243565b82525050565b5f60608201905061476d5f830186613c48565b61477a6020830185613dce565b614787604083018461474b565b949350505050565b5f6040820190506147a25f830185613c48565b6147af6020830184613c48565b9392505050565b5f6040820190506147c95f830185613c48565b6147d66020830184613f86565b9392505050565b5f6060820190506147f05f830186613c48565b6147fd6020830185613dce565b61480a6040830184613dce565b949350505050565b5f6040820190506148255f830185613f86565b6148326020830184613f86565b9392505050565b5f5ffd5b5f5ffd5b5f5f8585111561485457614853614839565b5b838611156148655761486461483d565b5b6001850283019150848603905094509492505050565b5f60608201905061488e5f830186613f86565b61489b6020830185613f86565b6148a86040830184613f86565b949350505050565b5f6060820190508181035f8301526148c88186614608565b90506148d76020830185613f86565b6148e46040830184613f86565b949350505050565b5f6148f6826143c6565b61490081856146d8565b93506149108185602086016143e0565b614919816135d3565b840191505092915050565b5f60c0820190506149375f830189613c48565b6149446020830188613f86565b6149516040830187613f86565b61495e6060830186613f86565b61496b6080830185613f86565b81810360a083015261497d81846148ec565b9050979650505050505050565b5f60608201905061499d5f830186613c48565b6149aa6020830185613f86565b81810360408301526149bc81846148ec565b9050949350505050565b5f6060820190508181035f8301526149de8186614608565b90506149ed6020830185613f86565b81810360408301526149ff81846148ec565b9050949350505050565b5f604082019050614a1c5f830185613cf4565b614a296020830184613dce565b9392505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f614a6e600283614a30565b9150614a7982614a3a565b600282019050919050565b5f819050919050565b614a9e614a9982613952565b614a84565b82525050565b5f614aae82614a62565b9150614aba8285614a8d565b602082019150614aca8284614a8d565b6020820191508190509392505050565b5f604082019050614aed5f830185613dce565b8181036020830152614aff81846148ec565b90509392505050565b5f81519050614b168161395b565b92915050565b5f81519050614b2a81613706565b92915050565b5f60408284031215614b4557614b446135cf565b5b614b4f6040613641565b90505f614b5e84828501614b08565b5f830152506020614b7184828501614b1c565b60208301525092915050565b5f60408284031215614b9257614b91613542565b5b5f614b9f84828501614b30565b91505092915050565b604082015f820151614bbc5f850182614551565b506020820151614bcf60208501826143b7565b50505050565b5f604082019050614be85f830184614ba8565b92915050565b5f60a082019050614c015f830188613c48565b614c0e6020830187613c48565b614c1b6040830186613c48565b614c286060830185613f86565b614c356080830184613dce565b9695505050505050565b5f81905092915050565b614c5281613569565b82525050565b5f614c638383614c49565b60208301905092915050565b5f614c7982614560565b614c838185614c3f565b9350614c8e8361457a565b805f5b83811015614cbe578151614ca58882614c58565b9750614cb0836145a0565b925050600181019050614c91565b5085935050505092915050565b5f614cd68284614c6f565b915081905092915050565b5f60a082019050614cf45f830188613c48565b614d016020830187613c48565b614d0e6040830186613f86565b614d1b6060830185613f86565b614d286080830184613c48565b9695505050505050565b5f606082019050614d455f830186613c48565b614d526020830185613c48565b614d5f6040830184613c48565b949350505050565b614d708161365f565b82525050565b5f602082019050614d895f830184614d67565b92915050565b5f6060820190508181035f830152614da88186886146e8565b9050614db76020830185613f86565b614dc46040830184613f86565b95945050505050565b5f608082019050614de05f830187613c48565b614ded6020830186614d67565b614dfa6040830185613c48565b614e076060830184613c48565b95945050505050565b5f604082019050614e235f830186613c48565b8181036020830152614e368184866146e8565b9050949350505050565b5f81519050614e4e81613d79565b92915050565b5f60208284031215614e6957614e68613542565b5b5f614e7684828501614e40565b91505092915050565b5f606082019050614e925f830187613c48565b614e9f6020830186613dce565b8181036040830152614eb28184866146e8565b905095945050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f614ef1601c83614a30565b9150614efc82614ebd565b601c82019050919050565b5f614f1182614ee5565b9150614f1d8284614a8d565b60208201915081905092915050565b5f6040820190508181035f830152614f448186614608565b90508181036020830152614f598184866146e8565b9050949350505050565b5f60208284031215614f7857614f77613542565b5b5f614f8584828501614b08565b91505092915050565b5f614f98826143c6565b614fa281856134e5565b9350614fb28185602086016143e0565b80840191505092915050565b5f614fc98285614f8e565b9150614fd58284614a8d565b6020820191508190509392505050565b7f53657175656e6365207369676e65723a0a0000000000000000000000000000005f82015250565b5f615019601183614a30565b915061502482614fe5565b601182019050919050565b5f8160601b9050919050565b5f6150458261502f565b9050919050565b5f6150568261503b565b9050919050565b61506e61506982613569565b61504c565b82525050565b5f819050919050565b61508e615089826136fd565b615074565b82525050565b5f61509e8261500d565b91506150aa828561505d565b6014820191506150ba828461507d565b6020820191508190509392505050565b7f53657175656e6365206e657374656420636f6e6669673a0a00000000000000005f82015250565b5f6150fe601883614a30565b9150615109826150ca565b601882019050919050565b5f61511e826150f2565b915061512a8286614a8d565b60208201915061513a828561507d565b60208201915061514a828461507d565b602082019150819050949350505050565b7f53657175656e636520737461746963206469676573743a0a00000000000000005f82015250565b5f61518f601883614a30565b915061519a8261515b565b601882019050919050565b5f6151af82615183565b91506151bb8284614a8d565b60208201915081905092915050565b7f53657175656e636520616e792061646472657373207375626469676573743a0a5f82015250565b5f6151fe602083614a30565b9150615209826151ca565b602082019050919050565b5f61521e826151f2565b915061522a8284614a8d565b60208201915081905092915050565b7f53657175656e63652073617069656e7420636f6e6669673a0a000000000000005f82015250565b5f61526d601983614a30565b915061527882615239565b601982019050919050565b5f61528d82615261565b9150615299828661505d565b6014820191506152a9828561507d565b6020820191506152b98284614a8d565b602082019150819050949350505050565b5f610100820190506152de5f83018b613c48565b6152eb602083018a613dce565b6152f86040830189613f86565b6153056060830188613c48565b6153126080830187613f86565b61531f60a0830186613fd5565b61532c60c0830185613fd5565b61533960e0830184613f86565b999850505050505050505056fea264697066735822122019b5533fc4cbee24a73438e1a10df64732fe25b8872932ba010c1d5c7babb5eb64736f6c634300081c0033",
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

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage2 *WalletStage2Caller) GetStaticSignature(opts *bind.CallOpts, _hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "getStaticSignature", _hash)

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
func (_WalletStage2 *WalletStage2Session) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage2.Contract.GetStaticSignature(&_WalletStage2.CallOpts, _hash)
}

// GetStaticSignature is a free data retrieval call binding the contract method 0x92dcb3fc.
//
// Solidity: function getStaticSignature(bytes32 _hash) view returns(address, uint256)
func (_WalletStage2 *WalletStage2CallerSession) GetStaticSignature(_hash [32]byte) (common.Address, *big.Int, error) {
	return _WalletStage2.Contract.GetStaticSignature(&_WalletStage2.CallOpts, _hash)
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

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2Caller) RecoverSapientSignature(opts *bind.CallOpts, _payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletStage2.contract.Call(opts, &out, "recoverSapientSignature", _payload, _signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2Session) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2.Contract.RecoverSapientSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) _payload, bytes _signature) view returns(bytes32)
func (_WalletStage2 *WalletStage2CallerSession) RecoverSapientSignature(_payload PayloadDecoded, _signature []byte) ([32]byte, error) {
	return _WalletStage2.Contract.RecoverSapientSignature(&_WalletStage2.CallOpts, _payload, _signature)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2 *WalletStage2Transactor) AddHook(opts *bind.TransactOpts, signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "addHook", signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2 *WalletStage2Session) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.AddHook(&_WalletStage2.TransactOpts, signature, implementation)
}

// AddHook is a paid mutator transaction binding the contract method 0xb93ea7ad.
//
// Solidity: function addHook(bytes4 signature, address implementation) payable returns()
func (_WalletStage2 *WalletStage2TransactorSession) AddHook(signature [4]byte, implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.AddHook(&_WalletStage2.TransactOpts, signature, implementation)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2 *WalletStage2Transactor) Execute(opts *bind.TransactOpts, _payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "execute", _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2 *WalletStage2Session) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Execute(&_WalletStage2.TransactOpts, _payload, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0x1f6a1eb9.
//
// Solidity: function execute(bytes _payload, bytes _signature) payable returns()
func (_WalletStage2 *WalletStage2TransactorSession) Execute(_payload []byte, _signature []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.Execute(&_WalletStage2.TransactOpts, _payload, _signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2 *WalletStage2Transactor) RemoveHook(opts *bind.TransactOpts, signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "removeHook", signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2 *WalletStage2Session) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.RemoveHook(&_WalletStage2.TransactOpts, signature)
}

// RemoveHook is a paid mutator transaction binding the contract method 0x4fcf3eca.
//
// Solidity: function removeHook(bytes4 signature) payable returns()
func (_WalletStage2 *WalletStage2TransactorSession) RemoveHook(signature [4]byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.RemoveHook(&_WalletStage2.TransactOpts, signature)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage2 *WalletStage2Transactor) SelfExecute(opts *bind.TransactOpts, _payload []byte) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "selfExecute", _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
func (_WalletStage2 *WalletStage2Session) SelfExecute(_payload []byte) (*types.Transaction, error) {
	return _WalletStage2.Contract.SelfExecute(&_WalletStage2.TransactOpts, _payload)
}

// SelfExecute is a paid mutator transaction binding the contract method 0x6ea44577.
//
// Solidity: function selfExecute(bytes _payload) payable returns()
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
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage2 *WalletStage2Transactor) UpdateImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.contract.Transact(opts, "updateImplementation", _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
func (_WalletStage2 *WalletStage2Session) UpdateImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _WalletStage2.Contract.UpdateImplementation(&_WalletStage2.TransactOpts, _implementation)
}

// UpdateImplementation is a paid mutator transaction binding the contract method 0x025b22bc.
//
// Solidity: function updateImplementation(address _implementation) payable returns()
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

// WalletStage2CallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the WalletStage2 contract.
type WalletStage2CallAbortedIterator struct {
	Event *WalletStage2CallAborted // Event containing the contract specifics and raw log

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
func (it *WalletStage2CallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2CallAborted)
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
		it.Event = new(WalletStage2CallAborted)
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
func (it *WalletStage2CallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2CallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2CallAborted represents a CallAborted event raised by the WalletStage2 contract.
type WalletStage2CallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage2 *WalletStage2Filterer) FilterCallAborted(opts *bind.FilterOpts) (*WalletStage2CallAbortedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &WalletStage2CallAbortedIterator{contract: _WalletStage2.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage2 *WalletStage2Filterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *WalletStage2CallAborted) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2CallAborted)
				if err := _WalletStage2.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseCallAborted(log types.Log) (*WalletStage2CallAborted, error) {
	event := new(WalletStage2CallAborted)
	if err := _WalletStage2.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2CallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the WalletStage2 contract.
type WalletStage2CallFailedIterator struct {
	Event *WalletStage2CallFailed // Event containing the contract specifics and raw log

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
func (it *WalletStage2CallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2CallFailed)
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
		it.Event = new(WalletStage2CallFailed)
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
func (it *WalletStage2CallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2CallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2CallFailed represents a CallFailed event raised by the WalletStage2 contract.
type WalletStage2CallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage2 *WalletStage2Filterer) FilterCallFailed(opts *bind.FilterOpts) (*WalletStage2CallFailedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &WalletStage2CallFailedIterator{contract: _WalletStage2.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_WalletStage2 *WalletStage2Filterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *WalletStage2CallFailed) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2CallFailed)
				if err := _WalletStage2.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseCallFailed(log types.Log) (*WalletStage2CallFailed, error) {
	event := new(WalletStage2CallFailed)
	if err := _WalletStage2.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2CallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the WalletStage2 contract.
type WalletStage2CallSkippedIterator struct {
	Event *WalletStage2CallSkipped // Event containing the contract specifics and raw log

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
func (it *WalletStage2CallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2CallSkipped)
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
		it.Event = new(WalletStage2CallSkipped)
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
func (it *WalletStage2CallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2CallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2CallSkipped represents a CallSkipped event raised by the WalletStage2 contract.
type WalletStage2CallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterCallSkipped(opts *bind.FilterOpts) (*WalletStage2CallSkippedIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &WalletStage2CallSkippedIterator{contract: _WalletStage2.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *WalletStage2CallSkipped) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2CallSkipped)
				if err := _WalletStage2.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseCallSkipped(log types.Log) (*WalletStage2CallSkipped, error) {
	event := new(WalletStage2CallSkipped)
	if err := _WalletStage2.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletStage2CallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the WalletStage2 contract.
type WalletStage2CallSucceededIterator struct {
	Event *WalletStage2CallSucceeded // Event containing the contract specifics and raw log

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
func (it *WalletStage2CallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletStage2CallSucceeded)
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
		it.Event = new(WalletStage2CallSucceeded)
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
func (it *WalletStage2CallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletStage2CallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletStage2CallSucceeded represents a CallSucceeded event raised by the WalletStage2 contract.
type WalletStage2CallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) FilterCallSucceeded(opts *bind.FilterOpts) (*WalletStage2CallSucceededIterator, error) {

	logs, sub, err := _WalletStage2.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &WalletStage2CallSucceededIterator{contract: _WalletStage2.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_WalletStage2 *WalletStage2Filterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *WalletStage2CallSucceeded) (event.Subscription, error) {

	logs, sub, err := _WalletStage2.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletStage2CallSucceeded)
				if err := _WalletStage2.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_WalletStage2 *WalletStage2Filterer) ParseCallSucceeded(log types.Log) (*WalletStage2CallSucceeded, error) {
	event := new(WalletStage2CallSucceeded)
	if err := _WalletStage2.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
