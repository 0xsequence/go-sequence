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
	Bin: "0x60e060405234801561001057600080fd5b5060405161821438038061821483398101604081905261002f916100c7565b8060405161003c906100ba565b604051809103906000f080158015610058573d6000803e3d6000fd5b5060006040518060600160405280602c81526020016181e8602c9139604051610086919030906020016100f7565b60408051601f198184030181529190528051602090910120608052506001600160a01b0391821660a0521660c05250610129565b613f948061425483390190565b6000602082840312156100d957600080fd5b81516001600160a01b03811681146100f057600080fd5b9392505050565b6000835160005b8181101561011857602081870181015185830152016100fe565b509190910191825250602001919050565b60805160a05160c0516140e761016d600039600081816104b9015261184d0152600081816103d20152611d6f01526000818161037e0152611da001526140e76000f3fe60806040526004361061016e5760003560e01c80636ea44577116100cb578063aaf10f421161007f578063bc197c8111610059578063bc197c8114610556578063f23a6e611461059e578063f727ef1c146105e457610175565b8063aaf10f42146104db578063ad55366b146104f0578063b93ea7ad1461054357610175565b80638c3f5563116100b05780638c3f55631461043b57806392dcb3fc1461045b5780639f69ef54146104a757610175565b80636ea44577146104075780638943ec021461041a57610175565b80631f6a1eb911610122578063295614261161010757806329561426146103a05780632dd31000146103c05780634fcf3eca146103f457610175565b80631f6a1eb914610359578063257671f51461036c57610175565b8063150b7a0211610153578063150b7a021461027e5780631626ba7e146102f45780631a9b23371461031457610175565b8063025b22bc1461023857806313792a4a1461024b57610175565b3661017557005b6004361061023657600061019161018c368361318d565b610604565b905073ffffffffffffffffffffffffffffffffffffffff811615610234576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101da9291906131f3565b600060405180830381855af49150503d8060008114610215576040519150601f19603f3d011682016040523d82523d6000602084013e61021a565b606091505b50915091508161022c57805160208201fd5b805160208201f35b505b005b61023661024636600461322c565b610658565b34801561025757600080fd5b5061026b6102663660046135d0565b6106a4565b6040519081526020015b60405180910390f35b34801561028a57600080fd5b506102c3610299366004613717565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610275565b34801561030057600080fd5b506102c361030f366004613786565b61080f565b34801561032057600080fd5b5061033461032f3660046137e7565b6108a6565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610275565b610236610367366004613804565b6108b1565b34801561037857600080fd5b5061026b7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103ac57600080fd5b506102366103bb366004613875565b61093a565b3480156103cc57600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b6102366104023660046137e7565b61097e565b61023661041536600461388e565b610a40565b34801561042657600080fd5b506102366104353660046138d0565b50505050565b34801561044757600080fd5b5061026b610456366004613875565b610aad565b34801561046757600080fd5b5061047b610476366004613875565b610ad9565b6040805173ffffffffffffffffffffffffffffffffffffffff9093168352602083019190915201610275565b3480156104b357600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000081565b3480156104e757600080fd5b50610334610aee565b3480156104fc57600080fd5b5061051061050b3660046135d0565b610afd565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c001610275565b610236610551366004613912565b610b37565b34801561056257600080fd5b506102c361057136600461398c565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156105aa57600080fd5b506102c36105b9366004613a53565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b3480156105f057600080fd5b506102366105ff366004613acb565b610bfc565b60006106527fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610cb7565b92915050565b333014610698576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6106a181610d15565b50565b6000808461010001515160016106ba9190613b1c565b67ffffffffffffffff8111156106d2576106d2613247565b6040519080825280602002602001820160405280156106fb578160200160208202803683370190505b50905060005b8561010001515181101561076d57856101000151818151811061072657610726613b56565b602002602001015182828151811061074057610740613b56565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610701565b503381866101000151518151811061078757610787613b56565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152610100850181905260006107c1868686610d69565b50905080610801578585856040517ff58cc8b500000000000000000000000000000000000000000000000000000000815260040161068f93929190613e08565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e081018590526000610867828686610d69565b5090508061087b5750600091506108089050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b600061065282610604565b60005a905060006108c28686610f50565b90506108d68160600151826080015161136e565b6000806108e4838787610d69565b9150915081610925578286866040517fa2b6d61b00000000000000000000000000000000000000000000000000000000815260040161068f93929190613e08565b610930848285611456565b5050505050505050565b333014610975576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161068f565b6106a1816117b5565b3330146109b9576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161068f565b60006109c482610604565b73ffffffffffffffffffffffffffffffffffffffff1603610a35576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008216600482015260240161068f565b6106a1816000611871565b333014610a7b576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161068f565b60005a90506000610a8c8484610f50565b90506000610a9982611931565b9050610aa6838284611456565b5050505050565b60006106527f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610cb7565b600080610ae5836119ac565b91509150915091565b6000610af8305490565b905090565b600080600080600080610b148989896000806119f8565b939950919750945092509050610b2983611d1d565b935093975093979195509350565b333014610b72576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161068f565b6000610b7d83610604565b73ffffffffffffffffffffffffffffffffffffffff1614610bee576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161068f565b610bf88282611871565b5050565b333014610c37576040517fa19dbf0000000000000000000000000000000000000000000000000000000000815233600482015260240161068f565b610c508383836bffffffffffffffffffffffff16611e1f565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610cd6929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610d1d813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca039060200160405180910390a150565b600080600084846000818110610d8157610d81613b56565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610ed657610dde86611931565b9150600080610dec846119ac565b91509150428111610e33576040517ff95b6ab7000000000000000000000000000000000000000000000000000000008152600481018590526024810182905260440161068f565b73ffffffffffffffffffffffffffffffffffffffff821615801590610e6e575073ffffffffffffffffffffffffffffffffffffffff82163314155b15610eca576040517f8945c3130000000000000000000000000000000000000000000000000000000081526004810185905233602482015273ffffffffffffffffffffffffffffffffffffffff8316604482015260640161068f565b60019450505050610f48565b6000806000610ee98989896000806119f8565b985092955090935091505082821015610f38576040517ffd41fcba000000000000000000000000000000000000000000000000000000008152600481018490526024810183905260440161068f565b610f4181611d1d565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610fb45760006060840152610fc5565b84810135606090811c908401526014015b6007600183901c1680156110185760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b60008360101660100361102d57506001611056565b836020166020036110495750600282019186013560f01c611056565b50600182019186013560f81c5b8067ffffffffffffffff81111561106f5761106f613247565b6040519080825280602002602001820160405280156110fa57816020015b6110e76040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b81526020019060019003908161108d5790505b50604086015260005b818110156113635760018085019489013560f81c90808216900361116257308760400151838151811061113857611138613b56565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690526111ac565b8489013560601c601486018860400151848151811061118357611183613b56565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b600280821690036111ea578489013560208601886040015184815181106111d5576111d5613b56565b60200260200101516020018197508281525050505b6004808216900361128257600385019489013560e81c89868a61120d8483613b1c565b9261121a93929190613e38565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050604089015180518590811061126557611265613b56565b60209081029190910101516040015261127e8187613b1c565b9550505b600880821690036112c0578489013560208601886040015184815181106112ab576112ab613b56565b60200260200101516060018197508281525050505b8060101660ff16601014876040015183815181106112e0576112e0613b56565b602002602001015160800190151590811515815250508060201660ff166020148760400151838151811061131657611316613b56565b602090810291909101015190151560a090910152604087015180516003600684901c1691908490811061134b5761134b613b56565b602090810291909101015160c0015250600101611103565b505050505092915050565b600061137983610aad565b90508181146113c5576040517f9b6514f400000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810182905260640161068f565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b818110156117ad5760008460400151828151811061148157611481613b56565b602002602001015190508060a00151801561149a575083155b156114de5760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a1506117a5565b60608101516000945080158015906114f55750805a105b156115325785835a6040517f2139527400000000000000000000000000000000000000000000000000000000815260040161068f93929190613e62565b60008260800151156116065782516115ff9083156115505783611552565b5a5b634c4e814c60e01b8b8d898b8e606001518b6040015160405160240161157d96959493929190613e87565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611eae565b905061162d565b8251602084015161162a9190841561161e5784611620565b5a5b8660400151611ec4565b90505b806117685760c083015161168957600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d888561166a611edc565b60405161167993929190613ec4565b60405180910390a15050506117a5565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016116f35786846116be611edc565b6040517f7f6b0bb100000000000000000000000000000000000000000000000000000000815260040161068f93929190613ee3565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611768577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b8885611749611edc565b60405161175893929190613ec4565b60405180910390a15050506117ad565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b600101611461565b505050505050565b806117ec576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6118157fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa9060200160405180910390a16106a17f0000000000000000000000000000000000000000000000000000000000000000610d15565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b600080611942836020015130611efb565b9050600061194f84611fc8565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b600080806119da7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610cb7565b606081901c956bffffffffffffffffffffffff909116945092505050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611a43575073ffffffffffffffffffffffffffffffffffffffff8916155b15611b5f578b82013560601c985060149091019089611b5f5760038201918c013560e81c60008d848e611a768583613b1c565b92611a8393929190613e38565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611b0e9030908590600401613f0e565b6040805180830381865afa158015611b2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b4e9190613f45565b9250611b5a8285613b1c565b935050505b82600116600103611b9957611b878d8a838f8f87908092611b8293929190613e38565b612234565b97509750975097509750505050611d10565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611c046001600586901c811690613b1c565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611c4f8d611931565b9350611c6d8d858e8e86908092611c6893929190613e38565b612478565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611cb8575080518614155b8015611cc8575080602001518511155b15611d0c576040517fccbb534f000000000000000000000000000000000000000000000000000000008152815160048201526020820151602482015260440161068f565b5050505b9550955095509550959050565b6000610652826040517fff0000000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000060601b166021820152603581018290527f000000000000000000000000000000000000000000000000000000000000000060558201526000903090607501604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012073ffffffffffffffffffffffffffffffffffffffff161492915050565b611ea97fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611f6b5746611f6e565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080826101000151604051602001611fe19190613f96565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120835190915060ff1661208a5760006120328460400151612e0a565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c00161198d565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161211a5760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161218a5760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e4602082015290810191909152606081018290526080016120fc565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016121fa5760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d466602082015290810191909152606081018290526080016120fc565b82516040517f0481832000000000000000000000000000000000000000000000000000000000815260ff909116600482015260240161068f565b6000806000806000612296604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b888210156124205760038201916000908b013560e81c6122de8482613b1c565b9150600090508a82146122f25760006122f4565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff830361234c5761233b8f8d8d8790869261233393929190613e38565b6001856119f8565b939d50919b5099509750955061236e565b612362858d8d8790869261233393929190613e38565b50929c50909a50985096505b898910156123ba5761238282858d8f613e38565b8b8b6040517fb006aba000000000000000000000000000000000000000000000000000000000815260040161068f9493929190613fe2565b819350878d60000151036123cd5760008d525b828710612410576040517f37daf62b000000000000000000000000000000000000000000000000000000008152600481018890526024810184905260440161068f565b50505060c08201859052836122be565b8a511580159061243457508a602001518511155b15611d0c576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c0151602482015260440161068f565b60008060005b83811015612e0057600181019085013560f881901c9060fc1c8061259b57600f821660008190036124b65750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa15801561254b573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b50600061256e828960ff16612e9b565b90508b61257b578061258a565b60008c81526020829052604090205b9b505050505050505050505061247e565b600181036125ff57600f821660008190036125bd5750600183019287013560f81c5b601484019388013560601c60006125d78260ff8516612e9b565b9050866125e457806125f3565b60008781526020829052604090205b9650505050505061247e565b600281036127f4576003821660008190036126215750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d9087926126df93929190613e38565b6040518463ffffffff1660e01b81526004016126fd93929190614009565b602060405180830381865afa15801561271a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061273e9190614023565b7fffffffff0000000000000000000000000000000000000000000000000000000016146127af578c848d8d8b90859261277993929190613e38565b6040517fb2fed7ae00000000000000000000000000000000000000000000000000000000815260040161068f9493929190614040565b8097508460ff168a01995060006127c9858760ff16612e9b565b9050896127d657806127e5565b60008a81526020829052604090205b9950505050505050505061247e565b600381036128285760208301928701358461280f578061281e565b60008581526020829052604090205b945050505061247e565b600481036128cb57600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c1685830180965081925050506000818601905060008061289d8e8e8e8e8c908892611c6893929190613e38565b91509150829750818a0199506128bd898260009182526020526040902090565b98505050505050505061247e565b60068103612995576003600283901c1660008190036128f15750600183019287013560f81c5b60038316600081900361290b5750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806129498f8f8f8f8d908892611c6893929190613e38565b9150915082985084821061295c57998501995b6000612969828789612f02565b90508a6129765780612985565b60008b81526020829052604090205b9a5050505050505050505061247e565b60058103612a025760208301928701358881036129d0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b60006129db82612f66565b9050856129e857806129f7565b60008681526020829052604090205b95505050505061247e565b60078103612b0857600f82166000819003612a245750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612529565b60088103612b5c5760208301928701356000612b248b82612fba565b9050808203612b51577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b60006125d782613035565b60098103612c9457600382166000819003612b7e5750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612c1893929190613e38565b6040518463ffffffff1660e01b8152600401612c3693929190613e08565b602060405180830381865afa158015612c53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c779190614076565b90508197508460ff168a01995060006127c9858760ff1684613070565b600a8103612dcb57600382166000819003612cb65750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612d4f93929190613e38565b6040518463ffffffff1660e01b8152600401612d6d93929190614009565b602060405180830381865afa158015612d8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dae9190614076565b90508198508560ff168b019a506000612969868860ff1684613070565b6040517fb2505f7c0000000000000000000000000000000000000000000000000000000081526004810182905260240161068f565b5094509492505050565b6000606060005b8351811015612e8c576000612e3e858381518110612e3157612e31613b56565b60200260200101516130de565b90508281604051602001612e5392919061408f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612e11565b50805160209091012092915050565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611faa565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612fcb846020015184611efb565b90506000612fd885611fc8565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612f9d565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612f47565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612f9d98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff0000000000000000000000000000000000000000000000000000000081169060048410156131ec577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461322757600080fd5b919050565b60006020828403121561323e57600080fd5b61080882613203565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff8111828210171561329957613299613247565b60405290565b604051610120810167ffffffffffffffff8111828210171561329957613299613247565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561330a5761330a613247565b604052919050565b803560ff8116811461322757600080fd5b8035801515811461322757600080fd5b600067ffffffffffffffff82111561334d5761334d613247565b5060051b60200190565b600082601f83011261336857600080fd5b813567ffffffffffffffff81111561338257613382613247565b6133b360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116016132c3565b8181528460208386010111156133c857600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f8301126133f657600080fd5b813561340961340482613333565b6132c3565b8082825260208201915060208360051b86010192508583111561342b57600080fd5b602085015b8381101561351857803567ffffffffffffffff81111561344f57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561348357600080fd5b61348b613276565b61349760208301613203565b815260408201356020820152606082013567ffffffffffffffff8111156134bd57600080fd5b6134cc8a602083860101613357565b604083015250608082013560608201526134e860a08301613323565b60808201526134f960c08301613323565b60a082015260e0919091013560c0820152835260209283019201613430565b5095945050505050565b600082601f83011261353357600080fd5b813561354161340482613333565b8082825260208201915060208360051b86010192508583111561356357600080fd5b602085015b838110156135185761357981613203565b835260209283019201613568565b60008083601f84011261359957600080fd5b50813567ffffffffffffffff8111156135b157600080fd5b6020830191508360208285010111156135c957600080fd5b9250929050565b6000806000604084860312156135e557600080fd5b833567ffffffffffffffff8111156135fc57600080fd5b8401610120818703121561360f57600080fd5b61361761329f565b61362082613312565b815261362e60208301613323565b6020820152604082013567ffffffffffffffff81111561364d57600080fd5b613659888285016133e5565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561368d57600080fd5b61369988828501613357565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff8111156136ce57600080fd5b6136da88828501613522565b61010083015250935050602084013567ffffffffffffffff8111156136fe57600080fd5b61370a86828701613587565b9497909650939450505050565b60008060008060006080868803121561372f57600080fd5b61373886613203565b945061374660208701613203565b935060408601359250606086013567ffffffffffffffff81111561376957600080fd5b61377588828901613587565b969995985093965092949392505050565b60008060006040848603121561379b57600080fd5b83359250602084013567ffffffffffffffff8111156136fe57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000811681146106a157600080fd5b6000602082840312156137f957600080fd5b8135610808816137b9565b6000806000806040858703121561381a57600080fd5b843567ffffffffffffffff81111561383157600080fd5b61383d87828801613587565b909550935050602085013567ffffffffffffffff81111561385d57600080fd5b61386987828801613587565b95989497509550505050565b60006020828403121561388757600080fd5b5035919050565b600080602083850312156138a157600080fd5b823567ffffffffffffffff8111156138b857600080fd5b6138c485828601613587565b90969095509350505050565b600080600080606085870312156138e657600080fd5b6138ef85613203565b935060208501359250604085013567ffffffffffffffff81111561385d57600080fd5b6000806040838503121561392557600080fd5b8235613930816137b9565b915061393e60208401613203565b90509250929050565b60008083601f84011261395957600080fd5b50813567ffffffffffffffff81111561397157600080fd5b6020830191508360208260051b85010111156135c957600080fd5b60008060008060008060008060a0898b0312156139a857600080fd5b6139b189613203565b97506139bf60208a01613203565b9650604089013567ffffffffffffffff8111156139db57600080fd5b6139e78b828c01613947565b909750955050606089013567ffffffffffffffff811115613a0757600080fd5b613a138b828c01613947565b909550935050608089013567ffffffffffffffff811115613a3357600080fd5b613a3f8b828c01613587565b999c989b5096995094979396929594505050565b60008060008060008060a08789031215613a6c57600080fd5b613a7587613203565b9550613a8360208801613203565b94506040870135935060608701359250608087013567ffffffffffffffff811115613aad57600080fd5b613ab989828a01613587565b979a9699509497509295939492505050565b600080600060608486031215613ae057600080fd5b83359250613af060208501613203565b915060408401356bffffffffffffffffffffffff81168114613b1157600080fd5b809150509250925092565b80820180821115610652577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613ba0578181015183820152602001613b88565b50506000910152565b60008151808452613bc1816020860160208601613b85565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613cc3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613c7f60e0860182613ba9565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613c11565b50909695505050505050565b600081518084526020840193506020830160005b82811015613d1757815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613ce3565b5093949350505050565b805160ff16825260006020820151613d3d602085018215159052565b5060408201516101206040850152613d59610120850182613bf3565b9050606083015160608501526080830151608085015260a083015184820360a0860152613d868282613ba9565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613db68282613ccf565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613e1b6040830186613d21565b8281036020840152613e2e818587613dbf565b9695505050505050565b60008085851115613e4857600080fd5b83861115613e5557600080fd5b5050820193919092039150565b606081526000613e756060830186613d21565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000613eb860c0830184613ba9565b98975050505050505050565b838152826020820152606060408201526000613db66060830184613ba9565b606081526000613ef66060830186613d21565b8460208401528281036040840152613e2e8185613ba9565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613f3d6040830184613ba9565b949350505050565b60006040828403128015613f5857600080fd5b506040805190810167ffffffffffffffff81118282101715613f7c57613f7c613247565b604052825181526020928301519281019290925250919050565b8151600090829060208501835b82811015613fd757815173ffffffffffffffffffffffffffffffffffffffff16845260209384019390910190600101613fa3565b509195945050505050565b606081526000613ff6606083018688613dbf565b6020830194909452506040015292915050565b838152604060208201526000613db6604083018486613dbf565b60006020828403121561403557600080fd5b8151610808816137b9565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000613e2e606083018486613dbf565b60006020828403121561408857600080fd5b5051919050565b600083516140a1818460208801613b85565b919091019182525060200191905056fea26469706673582212205ebdcc1f641b736398850a460db2f367a6596107b28e4525fb7623f5eff8596764736f6c634300081b00336080604052348015600f57600080fd5b50613f758061001f6000396000f3fe6080604052600436106101485760003560e01c80636ea44577116100c0578063ad55366b11610074578063bc197c8111610059578063bc197c81146104a9578063f23a6e61146104f1578063f727ef1c146105375761014f565b8063ad55366b14610443578063b93ea7ad146104965761014f565b80638c3f5563116100a55780638c3f5563146103c257806392dcb3fc146103e2578063aaf10f421461042e5761014f565b80636ea445771461038e5780638943ec02146103a15761014f565b80631a9b23371161011757806329561426116100fc57806329561426146103465780634fcf3eca1461036657806351605d80146103795761014f565b80631a9b2337146102ee5780631f6a1eb9146103335761014f565b8063025b22bc1461021257806313792a4a14610225578063150b7a02146102585780631626ba7e146102ce5761014f565b3661014f57005b6004361061021057600061016b610166368361301b565b610557565b905073ffffffffffffffffffffffffffffffffffffffff81161561020e576000808273ffffffffffffffffffffffffffffffffffffffff166000366040516101b4929190613081565b600060405180830381855af49150503d80600081146101ef576040519150601f19603f3d011682016040523d82523d6000602084013e6101f4565b606091505b50915091508161020657805160208201fd5b805160208201f35b505b005b6102106102203660046130ba565b6105ab565b34801561023157600080fd5b5061024561024036600461345e565b6105f7565b6040519081526020015b60405180910390f35b34801561026457600080fd5b5061029d6102733660046135a5565b7f150b7a020000000000000000000000000000000000000000000000000000000095945050505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161024f565b3480156102da57600080fd5b5061029d6102e9366004613614565b610762565b3480156102fa57600080fd5b5061030e610309366004613675565b6107f9565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161024f565b610210610341366004613692565b610804565b34801561035257600080fd5b50610210610361366004613703565b61088d565b610210610374366004613675565b6108d1565b34801561038557600080fd5b50610245610993565b61021061039c36600461371c565b6109c2565b3480156103ad57600080fd5b506102106103bc36600461375e565b50505050565b3480156103ce57600080fd5b506102456103dd366004613703565b610a2f565b3480156103ee57600080fd5b506104026103fd366004613703565b610a5b565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835260208301919091520161024f565b34801561043a57600080fd5b5061030e610a70565b34801561044f57600080fd5b5061046361045e36600461345e565b610a7a565b604080519687526020870195909552921515938501939093526060840152608083019190915260a082015260c00161024f565b6102106104a43660046137a0565b610ab4565b3480156104b557600080fd5b5061029d6104c436600461381a565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b3480156104fd57600080fd5b5061029d61050c3660046138e1565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561054357600080fd5b50610210610552366004613959565b610b79565b60006105a57fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1207fffffffff000000000000000000000000000000000000000000000000000000008416610c34565b92915050565b3330146105eb576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b6105f481610c92565b50565b60008084610100015151600161060d91906139aa565b67ffffffffffffffff811115610625576106256130d5565b60405190808252806020026020018201604052801561064e578160200160208202803683370190505b50905060005b856101000151518110156106c0578561010001518181518110610679576106796139e4565b6020026020010151828281518110610693576106936139e4565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152600101610654565b50338186610100015151815181106106da576106da6139e4565b73ffffffffffffffffffffffffffffffffffffffff9092166020928302919091019091015261010085018190526000610714868686610ce7565b50905080610754578585856040517ff58cc8b50000000000000000000000000000000000000000000000000000000081526004016105e293929190613c96565b5060019150505b9392505050565b604080516101208101825260006020820181905260609282018390528282018190526080820181905260a0820183905260c082018190526101008201929092526003815260e0810185905260006107ba828686610ce7565b509050806107ce57506000915061075b9050565b507f20c13b0b0000000000000000000000000000000000000000000000000000000095945050505050565b60006105a582610557565b60005a905060006108158686610ece565b9050610829816060015182608001516112ec565b600080610837838787610ce7565b9150915081610878578286866040517fa2b6d61b0000000000000000000000000000000000000000000000000000000081526004016105e293929190613c96565b6108838482856113d4565b5050505050505050565b3330146108c8576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105e2565b6105f481611733565b33301461090c576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105e2565b600061091782610557565b73ffffffffffffffffffffffffffffffffffffffff1603610988576040517f1c3812cc0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000821660048201526024016105e2565b6105f48160006117c3565b60006109bd7fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf85490565b905090565b3330146109fd576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105e2565b60005a90506000610a0e8484610ece565b90506000610a1b82611883565b9050610a288382846113d4565b5050505050565b60006105a57f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e83610c34565b600080610a67836118fe565b91509150915091565b60006109bd305490565b600080600080600080610a9189898960008061194a565b939950919750945092509050610aa683611c6f565b935093975093979195509350565b333014610aef576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105e2565b6000610afa83610557565b73ffffffffffffffffffffffffffffffffffffffff1614610b6b576040517f5b4d6d6a0000000000000000000000000000000000000000000000000000000081527fffffffff00000000000000000000000000000000000000000000000000000000831660048201526024016105e2565b610b7582826117c3565b5050565b333014610bb4576040517fa19dbf000000000000000000000000000000000000000000000000000000000081523360048201526024016105e2565b610bcd8383836bffffffffffffffffffffffff16611c7a565b6040805184815273ffffffffffffffffffffffffffffffffffffffff841660208201526bffffffffffffffffffffffff83168183015290517febf265acfac1c01de588ed7ef49743b9c3ce8d6d1edeaf510a1f5453228515b19181900360600190a1505050565b6000808383604051602001610c53929190918252602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012054949350505050565b610c9a813055565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f310ba5f1d2ed074b51e2eccd052a47ae9ab7c6b800d1fca3db3999d6a592ca03906020015b60405180910390a150565b600080600084846000818110610cff57610cff6139e4565b7fff000000000000000000000000000000000000000000000000000000000000009201359182169250507f80000000000000000000000000000000000000000000000000000000000000009081169003610e5457610d5c86611883565b9150600080610d6a846118fe565b91509150428111610db1576040517ff95b6ab700000000000000000000000000000000000000000000000000000000815260048101859052602481018290526044016105e2565b73ffffffffffffffffffffffffffffffffffffffff821615801590610dec575073ffffffffffffffffffffffffffffffffffffffff82163314155b15610e48576040517f8945c3130000000000000000000000000000000000000000000000000000000081526004810185905233602482015273ffffffffffffffffffffffffffffffffffffffff831660448201526064016105e2565b60019450505050610ec6565b6000806000610e6789898960008061194a565b985092955090935091505082821015610eb6576040517ffd41fcba00000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044016105e2565b610ebf81611c6f565b9550505050505b935093915050565b604080516101208101825260008082526020820181905260609282018390528282018190526080820181905260a0820183905260c0820181905260e0820152610100810191909152823560f81c60018082168103610f325760006060840152610f43565b84810135606090811c908401526014015b6007600183901c168015610f965760016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0187840135610100929092039190911c166080850152908101905b600083601016601003610fab57506001610fd4565b83602016602003610fc75750600282019186013560f01c610fd4565b50600182019186013560f81c5b8067ffffffffffffffff811115610fed57610fed6130d5565b60405190808252806020026020018201604052801561107857816020015b6110656040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016060815260200160008152602001600015158152602001600015158152602001600081525090565b81526020019060019003908161100b5790505b50604086015260005b818110156112e15760018085019489013560f81c9080821690036110e05730876040015183815181106110b6576110b66139e4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff909116905261112a565b8489013560601c6014860188604001518481518110611101576111016139e4565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90921690915294505b6002808216900361116857848901356020860188604001518481518110611153576111536139e4565b60200260200101516020018197508281525050505b6004808216900361120057600385019489013560e81c89868a61118b84836139aa565b9261119893929190613cc6565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408901518051859081106111e3576111e36139e4565b6020908102919091010151604001526111fc81876139aa565b9550505b6008808216900361123e57848901356020860188604001518481518110611229576112296139e4565b60200260200101516060018197508281525050505b8060101660ff166010148760400151838151811061125e5761125e6139e4565b602002602001015160800190151590811515815250508060201660ff1660201487604001518381518110611294576112946139e4565b602090810291909101015190151560a090910152604087015180516003600684901c169190849081106112c9576112c96139e4565b602090810291909101015160c0015250600101611081565b505050505092915050565b60006112f783610a2f565b9050818114611343576040517f9b6514f40000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018290526064016105e2565b604080517f8d0bf1fd623d628c741362c1289948e57b3e2905218c676d3e69abee36d6ae2e60208083019190915281830186905282518083038401815260609092019092528051910120600183019081905560408051858152602081018390527f1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881910160405180910390a150505050565b604081015151600090815b8181101561172b576000846040015182815181106113ff576113ff6139e4565b602002602001015190508060a001518015611418575083155b1561145c5760408051878152602081018490527f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b910160405180910390a150611723565b60608101516000945080158015906114735750805a105b156114b05785835a6040517f213952740000000000000000000000000000000000000000000000000000000081526004016105e293929190613cf0565b600082608001511561158457825161157d9083156114ce57836114d0565b5a5b634c4e814c60e01b8b8d898b8e606001518b604001516040516024016114fb96959493929190613d15565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611d09565b90506115ab565b825160208401516115a89190841561159c578461159e565b5a5b8660400151611d1f565b90505b806116e65760c083015161160757600195507f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d88856115e8611d37565b6040516115f793929190613d52565b60405180910390a1505050611723565b60c08301517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161167157868461163c611d37565b6040517f7f6b0bb10000000000000000000000000000000000000000000000000000000081526004016105e293929190613d71565b60c08301517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016116e6577fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b88856116c7611d37565b6040516116d693929190613d52565b60405180910390a150505061172b565b60408051898152602081018690527f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a910160405180910390a15050505b6001016113df565b505050505050565b8061176a576040517f4294d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117937fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8829055565b6040518181527f307ed6bd941ee9fc80f369c94af5fa11e25bab5102a6140191756c5474a30bfa90602001610cdc565b604080517fbe27a319efc8734e89e26ba4bc95f5c788584163b959f03fa04e2d7ab4b9a1206020808301919091527fffffffff000000000000000000000000000000000000000000000000000000008516828401819052835180840385018152606084018086528151919093012073ffffffffffffffffffffffffffffffffffffffff8616908190559152608082015290517f0d7fc113eaf016db4681a1ba86d083ce3e0961f321062a75ac2b0aeb33deeed19181900360a00190a15050565b600080611894836020015130611d56565b905060006118a184611e23565b6040517f1901000000000000000000000000000000000000000000000000000000000000602082015260228101849052604281018290529091506062015b6040516020818303038152906040528051906020012092505050919050565b6000808061192c7fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e8685610c34565b606081901c956bffffffffffffffffffffffff909116945092505050565b6040805180820182526000808252602082018190529182918291829182918a3560f81c91600191808416148015611995575073ffffffffffffffffffffffffffffffffffffffff8916155b15611ab1578b82013560601c985060149091019089611ab15760038201918c013560e81c60008d848e6119c885836139aa565b926119d593929190613cc6565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250506040517fccce3bc80000000000000000000000000000000000000000000000000000000081529293505073ffffffffffffffffffffffffffffffffffffffff8d169163ccce3bc89150611a609030908590600401613d9c565b6040805180830381865afa158015611a7c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aa09190613dd3565b9250611aac82856139aa565b935050505b82600116600103611aeb57611ad98d8a838f8f87908092611ad493929190613cc6565b61208f565b97509750975097509750505050611c62565b6002838116811460208f015283901c60071660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018e850135610100929092039190911c16838201909650925060009050611b566001600586901c8116906139aa565b60016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018f860135610100929092039190911c1699509092019150611ba18d611883565b9350611bbf8d858e8e86908092611bba93929190613cc6565b6122d3565b600090815260208a815260408083208352888252808320835273ffffffffffffffffffffffffffffffffffffffff8d1690915290208251919850965015801590611c0a575080518614155b8015611c1a575080602001518511155b15611c5e576040517fccbb534f00000000000000000000000000000000000000000000000000000000815281516004820152602082015160248201526044016105e2565b5050505b9550955095509550959050565b60006105a582612c65565b611d047fc852adf5e97c2fc3b38f405671e91b7af1697ef0287577f227ef10494c2a8e86847fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606086901b166bffffffffffffffffffffffff8516176040805160208082019590955280820193909352805180840382018152606090930190528151919092012055565b505050565b60008060008351602085018787f4949350505050565b6000806000835160208501878988f195945050505050565b60603d604051915060208201818101604052818352816000823e505090565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c563187f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de85611dc65746611dc9565b60005b6040805160208101959095528401929092526060830152608082015273ffffffffffffffffffffffffffffffffffffffff831660a082015260c0015b60405160208183030381529060405280519060200120905092915050565b600080826101000151604051602001611e3c9190613e24565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120835190915060ff16611ee5576000611e8d8460400151612c98565b606080860151608080880151604080517f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a260208201529081018690529384019290925282015260a0810184905290915060c0016118df565b825160ff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611f755760a08301518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193909352820152606081018290526080015b60405160208183030381529060405280519060200120915050919050565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01611fe55760c0830151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201529081019190915260608101829052608001611f57565b825160ff167ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd016120555760e0830151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201529081019190915260608101829052608001611f57565b82516040517f0481832000000000000000000000000000000000000000000000000000000000815260ff90911660048201526024016105e2565b60008060008060006120f1604051806101200160405280600060ff168152602001600015158152602001606081526020016000815260200160008152602001606081526020016000801916815260200160008019168152602001606081525090565b6002815260007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5b8882101561227b5760038201916000908b013560e81c61213984826139aa565b9150600090508a821461214d57600061214f565b8d5b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83036121a7576121968f8d8d8790869261218e93929190613cc6565b60018561194a565b939d50919b509950975095506121c9565b6121bd858d8d8790869261218e93929190613cc6565b50929c50909a50985096505b89891015612215576121dd82858d8f613cc6565b8b8b6040517fb006aba00000000000000000000000000000000000000000000000000000000081526004016105e29493929190613e70565b819350878d60000151036122285760008d525b82871061226b576040517f37daf62b00000000000000000000000000000000000000000000000000000000815260048101889052602481018490526044016105e2565b50505060c0820185905283612119565b8a511580159061228f57508a602001518511155b15611c5e576040517fccbb534f0000000000000000000000000000000000000000000000000000000081528b51600482015260208c015160248201526044016105e2565b60008060005b83811015612c5b57600181019085013560f881901c9060fc1c806123f657600f821660008190036123115750600183019287013560f81c5b60408051600080825260208281018085528d9052601b8c89019182013560ff81811c928301908116868801529235606086018190527f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82166080870181905296909a0199959094919390929160019060a0015b6020604051602081039080840390855afa1580156123a6573d6000803e3d6000fd5b5050506020604051035190508660ff168c019b5060006123c9828960ff16612d29565b90508b6123d657806123e5565b60008c81526020829052604090205b9b50505050505050505050506122d9565b6001810361245a57600f821660008190036124185750600183019287013560f81c5b601484019388013560601c60006124328260ff8516612d29565b90508661243f578061244e565b60008781526020829052604090205b965050505050506122d9565b6002810361264f5760038216600081900361247c5750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168288018098508192505050600081880190506320c13b0b60e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168473ffffffffffffffffffffffffffffffffffffffff16631626ba7e8f8f8f8d90879261253a93929190613cc6565b6040518463ffffffff1660e01b815260040161255893929190613e97565b602060405180830381865afa158015612575573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125999190613eb1565b7fffffffff00000000000000000000000000000000000000000000000000000000161461260a578c848d8d8b9085926125d493929190613cc6565b6040517fb2fed7ae0000000000000000000000000000000000000000000000000000000081526004016105e29493929190613ece565b8097508460ff168a0199506000612624858760ff16612d29565b9050896126315780612640565b60008a81526020829052604090205b995050505050505050506122d9565b600381036126835760208301928701358461266a5780612679565b60008581526020829052604090205b94505050506122d9565b6004810361272657600f8216600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018a870135610100929092039190911c168583018096508192505050600081860190506000806126f88e8e8e8e8c908892611bba93929190613cc6565b91509150829750818a019950612718898260009182526020526040902090565b9850505050505050506122d9565b600681036127f0576003600283901c16600081900361274c5750600183019287013560f81c5b6003831660008190036127665750600284019388013560f01c5b6000858a013560e81c600387018162ffffff1691508097508192505050600081870190506000806127a48f8f8f8f8d908892611bba93929190613cc6565b915091508298508482106127b757998501995b60006127c4828789612d90565b90508a6127d157806127e0565b60008b81526020829052604090205b9a505050505050505050506122d9565b6005810361285d57602083019287013588810361282b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b600061283682612df4565b9050856128435780612852565b60008681526020829052604090205b9550505050506122d9565b6007810361296357600f8216600081900361287f5750600183019287013560f81c5b600080858a0135602087019650915089860135602087016040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018e905290975090915060ff82901c907f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831690601b830190600090600190605c01604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600084529083018083525260ff851690820152606081018890526080810185905260a001612384565b600881036129b7576020830192870135600061297f8b82612e48565b90508082036129ac577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff96505b600061243282612ec3565b60098103612aef576003821660008190036129d95750600183019287013560f81c5b60008489013560601c601486019550905060006003600286901c1660016008820290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c168188018098508193505050506000818701905060008373ffffffffffffffffffffffffffffffffffffffff166313792a4a8f8e8e8c908792612a7393929190613cc6565b6040518463ffffffff1660e01b8152600401612a9193929190613c96565b602060405180830381865afa158015612aae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ad29190613f04565b90508197508460ff168a0199506000612624858760ff1684612efe565b600a8103612c2657600382166000819003612b115750600183019287013560f81c5b60008489013560601c60148601955090506003600285901c16600060016008830290811b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff018c890135610100929092039190911c1682880180985081925050506000818801905060008473ffffffffffffffffffffffffffffffffffffffff1663898bd9218f8f8f8d908792612baa93929190613cc6565b6040518463ffffffff1660e01b8152600401612bc893929190613e97565b602060405180830381865afa158015612be5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c099190613f04565b90508198508560ff168b019a5060006127c4868860ff1684612efe565b6040517fb2505f7c000000000000000000000000000000000000000000000000000000008152600481018290526024016105e2565b5094509492505050565b600081158015906105a55750507fea7157fa25e3aa17d0ae2d5280fa4e24d421c61842aa85e45194e1145aa72bf8541490565b6000606060005b8351811015612d1a576000612ccc858381518110612cbf57612cbf6139e4565b6020026020010151612f6c565b90508281604051602001612ce1929190613f1d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052925050600101612c9f565b50805160209091012092915050565b6040517f53657175656e6365207369676e65723a0a00000000000000000000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084901b16603182015260458101829052600090606501611e05565b6040517f53657175656e6365206e657374656420636f6e6669673a0a000000000000000060208201526038810184905260588101839052607881018290526000906098015b6040516020818303038152906040528051906020012090509392505050565b6040517f53657175656e636520737461746963206469676573743a0a00000000000000006020820152603881018290526000906058015b604051602081830303815290604052805190602001209050919050565b600080612e59846020015184611d56565b90506000612e6685611e23565b6040517f190100000000000000000000000000000000000000000000000000000000000060208201526022810184905260428101829052909150606201604051602081830303815290604052805190602001209250505092915050565b604080517f53657175656e636520616e792061646472657373207375626469676573743a0a6020820152908101829052600090606001612e2b565b6040517f53657175656e63652073617069656e7420636f6e6669673a0a0000000000000060208201527fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166039820152604d8101839052606d8101829052600090608d01612dd5565b60007f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef43782600001518360200151846040015180519060200120856060015186608001518760a001518860c00151604051602001612e2b98979695949392919097885273ffffffffffffffffffffffffffffffffffffffff969096166020880152604087019490945260608601929092526080850152151560a0840152151560c083015260e08201526101000190565b80357fffffffff00000000000000000000000000000000000000000000000000000000811690600484101561307a577fffffffff00000000000000000000000000000000000000000000000000000000808560040360031b1b82161691505b5092915050565b8183823760009101908152919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146130b557600080fd5b919050565b6000602082840312156130cc57600080fd5b61075b82613091565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160e0810167ffffffffffffffff81118282101715613127576131276130d5565b60405290565b604051610120810167ffffffffffffffff81118282101715613127576131276130d5565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613198576131986130d5565b604052919050565b803560ff811681146130b557600080fd5b803580151581146130b557600080fd5b600067ffffffffffffffff8211156131db576131db6130d5565b5060051b60200190565b600082601f8301126131f657600080fd5b813567ffffffffffffffff811115613210576132106130d5565b61324160207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613151565b81815284602083860101111561325657600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261328457600080fd5b8135613297613292826131c1565b613151565b8082825260208201915060208360051b8601019250858311156132b957600080fd5b602085015b838110156133a657803567ffffffffffffffff8111156132dd57600080fd5b860160e08189037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001121561331157600080fd5b613319613104565b61332560208301613091565b815260408201356020820152606082013567ffffffffffffffff81111561334b57600080fd5b61335a8a6020838601016131e5565b6040830152506080820135606082015261337660a083016131b1565b608082015261338760c083016131b1565b60a082015260e0919091013560c08201528352602092830192016132be565b5095945050505050565b600082601f8301126133c157600080fd5b81356133cf613292826131c1565b8082825260208201915060208360051b8601019250858311156133f157600080fd5b602085015b838110156133a65761340781613091565b8352602092830192016133f6565b60008083601f84011261342757600080fd5b50813567ffffffffffffffff81111561343f57600080fd5b60208301915083602082850101111561345757600080fd5b9250929050565b60008060006040848603121561347357600080fd5b833567ffffffffffffffff81111561348a57600080fd5b8401610120818703121561349d57600080fd5b6134a561312d565b6134ae826131a0565b81526134bc602083016131b1565b6020820152604082013567ffffffffffffffff8111156134db57600080fd5b6134e788828501613273565b604083015250606082810135908201526080808301359082015260a082013567ffffffffffffffff81111561351b57600080fd5b613527888285016131e5565b60a08301525060c0828101359082015260e0808301359082015261010082013567ffffffffffffffff81111561355c57600080fd5b613568888285016133b0565b61010083015250935050602084013567ffffffffffffffff81111561358c57600080fd5b61359886828701613415565b9497909650939450505050565b6000806000806000608086880312156135bd57600080fd5b6135c686613091565b94506135d460208701613091565b935060408601359250606086013567ffffffffffffffff8111156135f757600080fd5b61360388828901613415565b969995985093965092949392505050565b60008060006040848603121561362957600080fd5b83359250602084013567ffffffffffffffff81111561358c57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000811681146105f457600080fd5b60006020828403121561368757600080fd5b813561075b81613647565b600080600080604085870312156136a857600080fd5b843567ffffffffffffffff8111156136bf57600080fd5b6136cb87828801613415565b909550935050602085013567ffffffffffffffff8111156136eb57600080fd5b6136f787828801613415565b95989497509550505050565b60006020828403121561371557600080fd5b5035919050565b6000806020838503121561372f57600080fd5b823567ffffffffffffffff81111561374657600080fd5b61375285828601613415565b90969095509350505050565b6000806000806060858703121561377457600080fd5b61377d85613091565b935060208501359250604085013567ffffffffffffffff8111156136eb57600080fd5b600080604083850312156137b357600080fd5b82356137be81613647565b91506137cc60208401613091565b90509250929050565b60008083601f8401126137e757600080fd5b50813567ffffffffffffffff8111156137ff57600080fd5b6020830191508360208260051b850101111561345757600080fd5b60008060008060008060008060a0898b03121561383657600080fd5b61383f89613091565b975061384d60208a01613091565b9650604089013567ffffffffffffffff81111561386957600080fd5b6138758b828c016137d5565b909750955050606089013567ffffffffffffffff81111561389557600080fd5b6138a18b828c016137d5565b909550935050608089013567ffffffffffffffff8111156138c157600080fd5b6138cd8b828c01613415565b999c989b5096995094979396929594505050565b60008060008060008060a087890312156138fa57600080fd5b61390387613091565b955061391160208801613091565b94506040870135935060608701359250608087013567ffffffffffffffff81111561393b57600080fd5b61394789828a01613415565b979a9699509497509295939492505050565b60008060006060848603121561396e57600080fd5b8335925061397e60208501613091565b915060408401356bffffffffffffffffffffffff8116811461399f57600080fd5b809150509250925092565b808201808211156105a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60005b83811015613a2e578181015183820152602001613a16565b50506000910152565b60008151808452613a4f816020860160208601613a13565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208501945060208160051b8301016020850160005b83811015613b51577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0858403018852815173ffffffffffffffffffffffffffffffffffffffff815116845260208101516020850152604081015160e06040860152613b0d60e0860182613a37565b6060838101519087015260808084015115159087015260a08084015115159087015260c0928301519290950191909152506020978801979190910190600101613a9f565b50909695505050505050565b600081518084526020840193506020830160005b82811015613ba557815173ffffffffffffffffffffffffffffffffffffffff16865260209586019590910190600101613b71565b5093949350505050565b805160ff16825260006020820151613bcb602085018215159052565b5060408201516101206040850152613be7610120850182613a81565b9050606083015160608501526080830151608085015260a083015184820360a0860152613c148282613a37565b91505060c083015160c085015260e083015160e0850152610100830151848203610100860152613c448282613b5d565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000613ca96040830186613baf565b8281036020840152613cbc818587613c4d565b9695505050505050565b60008085851115613cd657600080fd5b83861115613ce357600080fd5b5050820193919092039150565b606081526000613d036060830186613baf565b60208301949094525060400152919050565b86815285602082015284604082015283606082015282608082015260c060a08201526000613d4660c0830184613a37565b98975050505050505050565b838152826020820152606060408201526000613c446060830184613a37565b606081526000613d846060830186613baf565b8460208401528281036040840152613cbc8185613a37565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000613dcb6040830184613a37565b949350505050565b60006040828403128015613de657600080fd5b506040805190810167ffffffffffffffff81118282101715613e0a57613e0a6130d5565b604052825181526020928301519281019290925250919050565b8151600090829060208501835b82811015613e6557815173ffffffffffffffffffffffffffffffffffffffff16845260209384019390910190600101613e31565b509195945050505050565b606081526000613e84606083018688613c4d565b6020830194909452506040015292915050565b838152604060208201526000613c44604083018486613c4d565b600060208284031215613ec357600080fd5b815161075b81613647565b84815273ffffffffffffffffffffffffffffffffffffffff84166020820152606060408201526000613cbc606083018486613c4d565b600060208284031215613f1657600080fd5b5051919050565b60008351613f2f818460208801613a13565b919091019182525060200191905056fea264697066735822122006da30180defa0cbeac8b556cd633a8a6922d3219e091cc8ac780e6ef4e19c4764736f6c634300081b0033603e600e3d39601e805130553df33d3d34601c57363d3d373d363d30545af43d82803e903d91601c57fd5bf3",
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
