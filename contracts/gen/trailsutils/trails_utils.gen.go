// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trailsutils

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

// TrailsUtilsMetaData contains all meta data concerning the TrailsUtils contract.
var TrailsUtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"handleSequenceDelegateCall\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hydrateExecute\",\"inputs\":[{\"name\":\"packedPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hydratePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hydrateExecuteAndSweep\",\"inputs\":[{\"name\":\"packedPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hydratePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sweepTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokensToSweep\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"sweepNative\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC1155Approval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC1155ApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721Approval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721ApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721Owner\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721OwnerApproval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721OwnerApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721OwnerSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinBalance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinBalanceSelf\",\"inputs\":[{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155Balance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceApproval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceApprovalBatch\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceApprovalBatchSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceBatch\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceBatchSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20Allowance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20AllowanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20Balance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20BalanceAllowance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20BalanceAllowanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20BalanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireNonExpired\",\"inputs\":[{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"sweepTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokensToSweep\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"sweepNative\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Sweep\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DelegateCallFailed\",\"inputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"DelegateCallNotAllowed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155BalanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155BatchBalanceTooLow\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155NotApproved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20AllowanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20BalanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721NotApproved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NotOwner\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"requiredOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Expired\",\"inputs\":[{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidRepeatSection\",\"inputs\":[{\"name\":\"_tindex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_cindex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tindex2\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_cindex2\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NativeBalanceTooLow\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NativeSweepFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonTransactionPayload\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyDelegateCallAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownHydrateDataCommand\",\"inputs\":[{\"name\":\"flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnknownHydrateTypeCommand\",\"inputs\":[{\"name\":\"flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a08060405234602957306080526135af908161002e8239608051818181610c2d015261246b0152f35b5f80fdfe6080604052600436101561001a575b3615610018575f80fd5b005b5f3560e01c806284290d146101f7578062e79cd5146101f2578063024ea911146101ed57806303bbbd65146101e85780630cb17e5c146101e357806313792a4a146101de57806318c9003a146101d95780631c0bdfab146101d45780632ab793d7146101cf5780632f2227ba146101ca578063317df1e8146101c5578063333f790d146101c0578063372618cb146101bb57806342ab921e146101b6578063476a5f40146101b15780634c4e814c146101ac578063505c7ed3146101a75780636d42cf03146101a25780636ef1f0dd1461019d5780637f29d5381461019857806380df36a0146101935780638d0a9ea71461018e5780639c8d8ea114610189578063a7b7ec5a14610184578063ac38d0891461017f578063cda38f4c1461017a578063e246466a14610175578063f38bce1514610170578063fdc254281461016b5763fe08e7ca0361000e576110e9565b61105f565b61103b565b611022565b61100b565b610fb0565b610f3e565b610f19565b610f02565b610e51565b610ddc565b610dc3565b610d4d565b610cf6565b610bdb565b610b47565b610b2d565b610ae4565b610aa9565b610a42565b6109f9565b6109a2565b61097d565b61088e565b61047c565b6103fc565b6103b5565b610309565b6102b4565b61021e565b73ffffffffffffffffffffffffffffffffffffffff81160361021a57565b5f80fd5b3461021a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5761001860043561025c816101fc565b6024356044359133906113f7565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc606091011261021a576004356102a0816101fc565b906024356102ad816101fc565b9060443590565b3461021a576100186102c53661026a565b916102d183338361151b565b339061161c565b9181601f8401121561021a5782359167ffffffffffffffff831161021a576020808501948460051b01011161021a57565b3461021a5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57600435610344816101fc565b60243590610351826101fc565b60443567ffffffffffffffff811161021a576103719036906004016102d8565b92909160643567ffffffffffffffff811161021a576100189461039b6103b09236906004016102d8565b91608435966103a9886101fc565b86866118ac565b611a61565b3461021a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a576100186004356103f3816101fc565b60243590611b4e565b3461021a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5761001860043561043a816101fc565b60243590610447826101fc565b3390611a61565b9181601f8401121561021a5782359167ffffffffffffffff831161021a576020838186019501011161021a57565b3461021a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760043567ffffffffffffffff811161021a57806004016101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc833603011261021a5760243567ffffffffffffffff811161021a5761050e90369060040161044e565b909160ff61051b82611103565b166108665761053a60648501356084860135905f5260205260405f2090565b61054660248601611111565b156108525761055c905f525f60205260405f2090565b9193925b60445f9401925b610571848461111e565b905085101561069e57610696600191866105958161058f898961111e565b9061119f565b6106866105a1826111e4565b61065a6020840135936060810135906105bc60808201611111565b60c06105ca60a08401611111565b92013592604051978896602088019a8b93909796959273ffffffffffffffffffffffffffffffffffffffff60e09693610100875260046101008801527f63616c6c0000000000000000000000000000000000000000000000000000000061012088015261014087019a602088015216604086015260608501526080840152151560a0830152151560c08201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261123c565b519020905f5260205260405f2090565b940193610567565b93509391905f5b8381106106b757604051858152602090f35b9360808361070561ffff6106dd6106f2969798998b60019092919283013560f81c920190565b90968793918c81013560f01c91600290910190565b9216918b81013560f01c91600290910190565b9061073a836107306107268c61058f607f61ffff8998169a169a8b9261111e565b604081019061127d565b82880192916112ce565b909890961615610838575060038101966001918c019182013560f01c95913560f81c9491856107698c8c61111e565b610773929161119f565b604081016107809161127d565b888781019161078e936112ce565b929091369061079c92611320565b805190602001209136906107af92611320565b80519060200120036107e357916107cc93916107d9969593611c67565b905f5260205260405f2090565b935b9291906106a5565b6040517feacd10be000000000000000000000000000000000000000000000000000000008152600481019190915260248101919091526044810182905260648101839052608481018490528060a481015b0390fd5b9691905061084c959492506107cc93611ba5565b936107db565b5f9081524660205260409020919392610560565b7f3c28730f000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461021a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a576004356108c9816101fc565b60243567ffffffffffffffff811161021a576108e99036906004016102d8565b9160443567ffffffffffffffff811161021a57610018936109116104479236906004016102d8565b916064359561091f876101fc565b33866118ac565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc608091011261021a5760043561095c816101fc565b90602435610969816101fc565b90604435610976816101fc565b9060643590565b3461021a5761001861098e36610926565b9261099d848483959495611ce6565b611e06565b3461021a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a576100186004356109e0816101fc565b6024356109ec816101fc565b604435916103b0836101fc565b3461021a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57610018600435610a37816101fc565b602435903390611ce6565b3461021a5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57610018600435610a80816101fc565b602435610a8c816101fc565b608435916103b0606435604435610aa2866101fc565b84846113f7565b3461021a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5761001860043533611b4e565b3461021a5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57610018600435610b22816101fc565b60243590339061151b565b3461021a57610018610b3e36610926565b9291909161161c565b3461021a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57600435610b82816101fc565b602435610b8e816101fc565b60443567ffffffffffffffff811161021a57610bae9036906004016102d8565b606435939167ffffffffffffffff851161021a57610bd36100189536906004016102d8565b9490936118ac565b3461021a5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760a43567ffffffffffffffff811161021a57610c2a90369060040161044e565b907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff81163014610cce57825f939284936040519283928337810184815203915af4610c88611356565b9015610c9057005b610834906040519182917ff7a01e4d000000000000000000000000000000000000000000000000000000008352602060048401526024830190611385565b7f827e6c75000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461021a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57610018600435610d34816101fc565b602435610d40816101fc565b60443590606435926113f7565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760043567ffffffffffffffff811161021a57610d9790369060040161044e565b6024359167ffffffffffffffff831161021a57610dbb61001893369060040161044e565b9290916121af565b3461021a57610018610dd43661026a565b91339061161c565b3461021a5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760043580421015610e1857005b7faa2fd925000000000000000000000000000000000000000000000000000000005f526004524260245260445ffd5b8015150361021a57565b60a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760043567ffffffffffffffff811161021a57610e9b90369060040161044e565b9060243567ffffffffffffffff811161021a57610ebc90369060040161044e565b92604435610ec9816101fc565b6064359467ffffffffffffffff861161021a57610eed6100189636906004016102d8565b94909360843596610efd88610e47565b6113c8565b3461021a57610018610f133661026a565b9161151b565b3461021a57610018610f2a36610926565b92610f3984848395949561151b565b61161c565b3461021a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57600435610f79816101fc565b6024359067ffffffffffffffff821161021a57610f9d6100189236906004016102d8565b9060443592610fab84610e47565b612532565b3461021a5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a57610018600435610fee816101fc565b60643590610447604435602435611004856101fc565b33846113f7565b3461021a5761001861101c3661026a565b91611ce6565b3461021a576100186110333661026a565b913390611e06565b3461021a5761001861104c3661026a565b91611058833383611ce6565b3390611e06565b3461021a5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261021a5760043561109a816101fc565b60243567ffffffffffffffff811161021a576110ba9036906004016102d8565b916044359267ffffffffffffffff841161021a576110df6100189436906004016102d8565b93909233906118ac565b3461021a576100186110fa36610926565b92919091611e06565b3560ff8116810361021a5790565b3561111b81610e47565b90565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561021a570180359067ffffffffffffffff821161021a57602001918160051b3603831361021a57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b91908110156111df5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218136030182121561021a570190565b611172565b3561111b816101fc565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60e0810190811067ffffffffffffffff82111761123757604052565b6111ee565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761123757604052565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561021a570180359067ffffffffffffffff821161021a5760200191813603831361021a57565b9093929384831161021a57841161021a578101920390565b67ffffffffffffffff811161123757601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b92919261132c826112e6565b9161133a604051938461123c565b82948184528183011161021a578281602093845f960137010152565b3d15611380573d90611367826112e6565b91611375604051938461123c565b82523d5f602084013e565b606090565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b916113db979695949391610fab936121af565b565b9081602091031261021a575190565b6040513d5f823e3d90fd5b91929092604051907efdd58e00000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff8516600483015280602483015260208260448173ffffffffffffffffffffffffffffffffffffffff88165afa918215611516575f926114e5575b5082821061147f575050505050565b6040517fbf8ee56100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff948516600482015294909316602485015260448401929092526064830191909152608482015260a490fd5b61150891925060203d60201161150f575b611500818361123c565b8101906113dd565b905f611470565b503d6114f6565b6113ec565b919091604051917f70a0823100000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff8416600484015260208360248173ffffffffffffffffffffffffffffffffffffffff86165afa928315611516575f936115fb575b5080831061159c5750505050565b60606040519485947f7684050d00000000000000000000000000000000000000000000000000000000865273ffffffffffffffffffffffffffffffffffffffff6004870192816080850197168452166020830152604082015201520390fd5b61161591935060203d60201161150f57611500818361123c565b915f61158e565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015284811660248301529190911693929091602083604481885afa928315611516575f936116f5575b50838310611695575050505050565b60a495509073ffffffffffffffffffffffffffffffffffffffff8092604051967f73cf338a000000000000000000000000000000000000000000000000000000008852600488015216602486015216604484015260648301526084820152fd5b61170f91935060203d60201161150f57611500818361123c565b915f611686565b67ffffffffffffffff81116112375760051b60200190565b9061173882611716565b611745604051918261123c565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06117738294611716565b0190602036910137565b80518210156111df5760209160051b010190565b60208183031261021a5780519067ffffffffffffffff821161021a57019080601f8301121561021a5781516117c581611716565b926117d3604051948561123c565b81845260208085019260051b82010192831161021a57602001905b8282106117fb5750505090565b81518152602091820191016117ee565b604081016040825282518091526020606083019301905f5b8181106118705750505060208183039101528281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831161021a5760209260051b809284830137010190565b825173ffffffffffffffffffffffffffffffffffffffff16855260209485019490920191600101611823565b91908110156111df5760051b0190565b9190939293858503611a1a579084916118c48361172e565b905f5b8481106119df575050925f929173ffffffffffffffffffffffffffffffffffffffff61192395604051968795869485937f4e1273f40000000000000000000000000000000000000000000000000000000085526004850161180b565b0392165afa908115611516575f916119bd575b505f5b838110611947575050505050565b611951818361177d565b5161195d82878661189c565b351161196b57600101611939565b82611985828761197e826119ba9761177d565b519361189c565b7fe3758044000000000000000000000000000000000000000000000000000000005f5260049290925260245235604452606490565b5ffd5b6119d991503d805f833e6119d1818361123c565b810190611791565b5f611936565b82939450611a0f826119f4836001959661177d565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0190869392916118c7565b7fab8b67c6000000000000000000000000000000000000000000000000000000005f526004859052602486905260445ffd5b9081602091031261021a575161111b81610e47565b6040517fe985e9c500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015284811660248301529190911690602081604481855afa908115611516575f91611b1f575b5015611ad357505050565b73ffffffffffffffffffffffffffffffffffffffff929183917f628b17c6000000000000000000000000000000000000000000000000000000005f52600452166024521660445260645ffd5b611b41915060203d602011611b47575b611b39818361123c565b810190611a4c565b5f611ac8565b503d611b2f565b803190828210611b5d57505050565b73ffffffffffffffffffffffffffffffffffffffff907f1afe1f42000000000000000000000000000000000000000000000000000000005f521660045260245260445260645ffd5b92611c61917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6101009380604051978895602087019a60808c52600e60a08901527f7374617469632d73656374696f6e00000000000000000000000000000000000060c08901526040880152606087015260c060808701528160e0870152868601375f8582860101520116810103017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261123c565b51902090565b939192909260405193602085019560c08752600e60e08701527f7265706561742d73656374696f6e00000000000000000000000000000000000061010087015260408601526060850152608084015260a083015260c08201526101008152611c616101208261123c565b9081602091031261021a575161111b816101fc565b906040517f6352211e00000000000000000000000000000000000000000000000000000000815283600482015260208160248173ffffffffffffffffffffffffffffffffffffffff87165afa908115611516575f91611dd7575b5073ffffffffffffffffffffffffffffffffffffffff821673ffffffffffffffffffffffffffffffffffffffff821603611d7a5750505050565b6040517f9108eb1400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9384166004820152602481019490945282166044840152166064820152608490fd5b611df9915060203d602011611dff575b611df1818361123c565b810190611cd1565b5f611d40565b503d611de7565b919073ffffffffffffffffffffffffffffffffffffffff83166040517f081812fc00000000000000000000000000000000000000000000000000000000815260208180611e5b89600483019190602083019252565b0381855afa908115611516575f91611f87575b5073ffffffffffffffffffffffffffffffffffffffff808516911614159081611efa575b50611e9d5750505050565b6040517f6a7e2cc800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9384166004820152602481019490945282166044840152166064820152608490fd5b6040517fe985e9c500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152851660248201529150602090829060449082905afa908115611516575f91611f68575b50155f611e92565b611f81915060203d602011611b4757611b39818361123c565b5f611f60565b611fa0915060203d602011611dff57611df1818361123c565b5f611e6e565b9080602083519182815201916020808360051b8301019401925f915b838310611fd157505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080612044604085015160e0604086015260e0850190611385565b936060810151606085015260808101511515608085015260a0810151151560a0850152015191015297019301930191939290611fc2565b90602080835192838152019201905f5b8181106120985750505090565b825173ffffffffffffffffffffffffffffffffffffffff1684526020938401939092019160010161208b565b805160ff16825261111b916020828101511515908201526101006121226120fc60408501516101206040860152610120850190611fa6565b606085015160608501526080850151608085015260a085015184820360a0860152611385565b9260c081015160c084015260e081015160e084015201519061010081840391015261207b565b612160604092959493956060835260608301906120c4565b9460208201520152565b61111b9392606092825260208201528160408201520190611385565b61219c61111b94926060835260608301906120c4565b9260208201526040818403910152611385565b906121b991612864565b6121e46121c582612b25565b6121d0368686611320565b60208151910120905f5260205260405f2090565b916121ef8482612c31565b915f926040850192835151975f5b898110612211575b50505050505050505050565b80841461251b575b61222481875161177d565b519661223360a0890151151590565b80612513575b6124d357505f9660608101518015908115806124ca575b61249257608083018051151580612453575b6124265751156123dd57906122a39161228f845173ffffffffffffffffffffffffffffffffffffffff1690565b91156123d857505a5b604084015191613042565b156122e6575b50604080518a815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b016121fd565b60c00180511561239b57600181511461235c5751600214612307575f6122a9565b9796505050505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b925061234a61233e613053565b6040519384938461216a565b0390a15f808080808080808080612205565b5087610834612369613053565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501612186565b509550600180967f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d818b6123d061233e613053565b0390a16122e0565b612298565b61241b916123ff845173ffffffffffffffffffffffffffffffffffffffff1690565b916020850151915f1461242057505a905b604085015192613031565b6122a3565b90612410565b7f230d1ccc000000000000000000000000000000000000000000000000000000005f52600485905260245ffd5b5073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163014612262565b838b6108345a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501612148565b50805a10612250565b604080518b815260208101849052919850600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b91819081016123d0565b508015612239565b9192819361252a92868a612c96565b929091612219565b9092919073ffffffffffffffffffffffffffffffffffffffff8116612733575033925b5f5b8281106125f8575050506125685750565b479081612573575050565b5f80808085855af1612583611356565b50156125d05760405191825273ffffffffffffffffffffffffffffffffffffffff16905f907fed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab790602090a3565b7f16b452f7000000000000000000000000000000000000000000000000000000005f5260045ffd5b61262a61261161261161260c84878761189c565b6111e4565b73ffffffffffffffffffffffffffffffffffffffff1690565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529190602090839060249082905afa8015611516576001925f91612715575b5080612682575b5001612557565b61269d818861269861261161260c878b8b61189c565b61306d565b73ffffffffffffffffffffffffffffffffffffffff6126c061260c84888861189c565b167fed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab76040518061270c73ffffffffffffffffffffffffffffffffffffffff8c1695829190602083019252565b0390a35f61267b565b61272d915060203d811161150f57611500818361123c565b5f612674565b92612555565b60405190610120820182811067ffffffffffffffff821117611237576040526060610100835f81525f60208201528260408201525f838201525f60808201528260a08201525f60c08201525f60e08201520152565b9061279882611716565b6127a5604051918261123c565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06127d38294611716565b01905f5b8281106127e357505050565b6020906040516127f28161121b565b5f81525f83820152606060408201525f60608201525f60808201525f60a08201525f60c0820152828285010152016127d7565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9190820180921161285f57565b612825565b91909161286f612739565b5f815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81831601612b0e57505f60608701525b60076128bc60ff831660011c90565b1680612abd575b50601081811603612a8857506001925b6128dc8461278e565b92604087019384525f905b8582106128f657505050505050565b8281013560f81c90600101909181600180851603612a66575061293a3061291e83895161177d565b519073ffffffffffffffffffffffffffffffffffffffff169052565b600280841614612a46575b6004808416146129f8575b6008808416146129c5575b6129ae6129a860c0856129878a608061297e8860108060019c9d1614935161177d565b51019015159052565b61299e8a60a061297e88602080871614935161177d565b1660061c60031690565b60ff1690565b60c06129bb83895161177d565b51015201906128e7565b6001916129ae906129a89060c0908781013590602001969060606129ea878d5161177d565b51015295945050505061295b565b90612a40908481013560e81c90600301612a39612a20612a188484612852565b838a8a6112ce565b91906040612a2f888d5161177d565b5101923691611320565b9052612852565b90612950565b90838101359060200191906020612a5e83895161177d565b510152612945565b612a8392508481013560601c90601401929061291e83895161177d565b61293a565b602090811603612aaa5761ffff918381013560f01c906002015b9216926128d3565b60ff918381013560f81c90600101612aa2565b612b01919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b929060808701525f6128c3565b80850135606090811c9088015260140192506128ad565b602081015115612c165761065a611c61612bd45f5b60405160208101917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f83527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408301527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606083015260808201523060a082015260a08152612bcb60c08261123c565b51902093613179565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b61065a611c61612bd446612b3a565b908210156111df570190565b919015612c42576001913560f81c90565b5f91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461285f5760010190565b9492612cb190949291945b8581013560f81c91600190910190565b60ff8216818115612faa5750918785612cdb8694858b600f612cd4829a60041c90565b9716613352565b9095908060048611612f30575060028101959082013560f01c94869060018103612d275750505050916040612d1a612d229482612cb19895015161177d565b5101516133e1565b612ca1565b60028103612d555750505050916040612d4a612d229482612cb19895015161177d565b5101519131916133c8565b94959460038103612e3c57505050604080612da5602094612d9b73ffffffffffffffffffffffffffffffffffffffff9899612dfb9660149092919283013560601c920190565b999094015161177d565b510151966040519586809481937f70a082310000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b0392165afa90811561151657612cb194612d22935f93612e1c575b506133c8565b612e3591935060203d811161150f57611500818361123c565b915f612e16565b919291600414612e57575b505050505050612cb19150612ca1565b612f05959650612ea7602094612e9d60ff60409686612e868997612e969960149092919283013560601c920190565b90988183013560f81c9160010190565b9316613352565b999095015161177d565b5101516040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff988916600482015292881660248401529691948592839182906044820190565b0392165afa90811561151657612cb194612f25935f93612e1c57506133c8565b865f85858983612e47565b96955050505050600581145f14612f58575090612d22612cb19261291e8660408b015161177d565b60068103612f7f575090612cb191316020612f778660408b015161177d565b510152612ca1565b7ff024a3e2000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b975050505091929050838281105f1461300a57612ffa612fd46129a8926130069561300095612c25565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b92612c69565b9190565b5050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90565b915f9391849360208451940192f190565b915f9291839260208351930191f490565b3d90604051916020818401016040528083525f602084013e565b9173ffffffffffffffffffffffffffffffffffffffff604051927fa9059cbb000000000000000000000000000000000000000000000000000000005f521660045260245260205f60448180865af160015f5114811615613117575b604091909152156130d65750565b7f5274afe7000000000000000000000000000000000000000000000000000000005f5273ffffffffffffffffffffffffffffffffffffffff1660045260245ffd5b600181151661312d573d15833b151516166130c8565b503d5f823e3d90fd5b80516020909101905f5b81811061314d5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101613140565b6101008101516040516131948161065a602082018095613136565b519020906131a3815160ff1690565b60ff81168061321c57505090611c616131bf604084015161343f565b9261065a60806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b6001810361327a57505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290611c61816080810161065a565b600281036132d057505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252611c61816080810161065a565b600303613324575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252611c61816080810161065a565b7f04818320000000000000000000000000000000000000000000000000000000005f5260ff1660045260245ffd5b929392600f16915081613366575050309190565b60018203613375575050339190565b60028203613384575050329190565b6003829492146133ba57837f324dd124000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b81013560601c925060140190565b908151602082019082821091111761021a570160200152565b908151601482019082821091111761021a576020910101906bffffffffffffffffffffffff8251169060601b179052565b80516020909101905f5b8181106134295750505090565b825184526020938401939092019160010161341c565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061348561346f83611716565b9261347d604051948561123c565b808452611716565b013660208301375f5b835181101561356057806134a46001928661177d565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e0830152610100820152610100815261354c6101208261123c565b519020613559828561177d565b520161348e565b50909150604051611c618161065a60208201809561341256fea2646970667358221220a71302c1229d67897efb0ef833a703c8223d89df73ebf6079a4df6272b8fc02c64736f6c634300081e0033",
}

// TrailsUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use TrailsUtilsMetaData.ABI instead.
var TrailsUtilsABI = TrailsUtilsMetaData.ABI

// TrailsUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrailsUtilsMetaData.Bin instead.
var TrailsUtilsBin = TrailsUtilsMetaData.Bin

// DeployTrailsUtils deploys a new Ethereum contract, binding an instance of TrailsUtils to it.
func DeployTrailsUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrailsUtils, error) {
	parsed, err := TrailsUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrailsUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrailsUtils{TrailsUtilsCaller: TrailsUtilsCaller{contract: contract}, TrailsUtilsTransactor: TrailsUtilsTransactor{contract: contract}, TrailsUtilsFilterer: TrailsUtilsFilterer{contract: contract}}, nil
}

// TrailsUtils is an auto generated Go binding around an Ethereum contract.
type TrailsUtils struct {
	TrailsUtilsCaller     // Read-only binding to the contract
	TrailsUtilsTransactor // Write-only binding to the contract
	TrailsUtilsFilterer   // Log filterer for contract events
}

// TrailsUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrailsUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrailsUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrailsUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrailsUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrailsUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrailsUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrailsUtilsSession struct {
	Contract     *TrailsUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrailsUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrailsUtilsCallerSession struct {
	Contract *TrailsUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TrailsUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrailsUtilsTransactorSession struct {
	Contract     *TrailsUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TrailsUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrailsUtilsRaw struct {
	Contract *TrailsUtils // Generic contract binding to access the raw methods on
}

// TrailsUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrailsUtilsCallerRaw struct {
	Contract *TrailsUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// TrailsUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrailsUtilsTransactorRaw struct {
	Contract *TrailsUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrailsUtils creates a new instance of TrailsUtils, bound to a specific deployed contract.
func NewTrailsUtils(address common.Address, backend bind.ContractBackend) (*TrailsUtils, error) {
	contract, err := bindTrailsUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrailsUtils{TrailsUtilsCaller: TrailsUtilsCaller{contract: contract}, TrailsUtilsTransactor: TrailsUtilsTransactor{contract: contract}, TrailsUtilsFilterer: TrailsUtilsFilterer{contract: contract}}, nil
}

// NewTrailsUtilsCaller creates a new read-only instance of TrailsUtils, bound to a specific deployed contract.
func NewTrailsUtilsCaller(address common.Address, caller bind.ContractCaller) (*TrailsUtilsCaller, error) {
	contract, err := bindTrailsUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsCaller{contract: contract}, nil
}

// NewTrailsUtilsTransactor creates a new write-only instance of TrailsUtils, bound to a specific deployed contract.
func NewTrailsUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*TrailsUtilsTransactor, error) {
	contract, err := bindTrailsUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsTransactor{contract: contract}, nil
}

// NewTrailsUtilsFilterer creates a new log filterer instance of TrailsUtils, bound to a specific deployed contract.
func NewTrailsUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*TrailsUtilsFilterer, error) {
	contract, err := bindTrailsUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsFilterer{contract: contract}, nil
}

// bindTrailsUtils binds a generic wrapper to an already deployed contract.
func bindTrailsUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrailsUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrailsUtils *TrailsUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrailsUtils.Contract.TrailsUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrailsUtils *TrailsUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrailsUtils.Contract.TrailsUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrailsUtils *TrailsUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrailsUtils.Contract.TrailsUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrailsUtils *TrailsUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrailsUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrailsUtils *TrailsUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrailsUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrailsUtils *TrailsUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrailsUtils.Contract.contract.Transact(opts, method, params...)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_TrailsUtils *TrailsUtilsCaller) RecoverSapientSignature(opts *bind.CallOpts, payload PayloadDecoded, signature []byte) ([32]byte, error) {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "recoverSapientSignature", payload, signature)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_TrailsUtils *TrailsUtilsSession) RecoverSapientSignature(payload PayloadDecoded, signature []byte) ([32]byte, error) {
	return _TrailsUtils.Contract.RecoverSapientSignature(&_TrailsUtils.CallOpts, payload, signature)
}

// RecoverSapientSignature is a free data retrieval call binding the contract method 0x13792a4a.
//
// Solidity: function recoverSapientSignature((uint8,bool,(address,uint256,bytes,uint256,bool,bool,uint256)[],uint256,uint256,bytes,bytes32,bytes32,address[]) payload, bytes signature) view returns(bytes32 imageHash)
func (_TrailsUtils *TrailsUtilsCallerSession) RecoverSapientSignature(payload PayloadDecoded, signature []byte) ([32]byte, error) {
	return _TrailsUtils.Contract.RecoverSapientSignature(&_TrailsUtils.CallOpts, payload, signature)
}

// RequireERC1155Approval is a free data retrieval call binding the contract method 0x2ab793d7.
//
// Solidity: function requireERC1155Approval(address token, address owner, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC1155Approval(opts *bind.CallOpts, token common.Address, owner common.Address, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC1155Approval", token, owner, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireERC1155Approval is a free data retrieval call binding the contract method 0x2ab793d7.
//
// Solidity: function requireERC1155Approval(address token, address owner, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC1155Approval(token common.Address, owner common.Address, operator common.Address) error {
	return _TrailsUtils.Contract.RequireERC1155Approval(&_TrailsUtils.CallOpts, token, owner, operator)
}

// RequireERC1155Approval is a free data retrieval call binding the contract method 0x2ab793d7.
//
// Solidity: function requireERC1155Approval(address token, address owner, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC1155Approval(token common.Address, owner common.Address, operator common.Address) error {
	return _TrailsUtils.Contract.RequireERC1155Approval(&_TrailsUtils.CallOpts, token, owner, operator)
}

// RequireERC1155ApprovalSelf is a free data retrieval call binding the contract method 0x0cb17e5c.
//
// Solidity: function requireERC1155ApprovalSelf(address token, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC1155ApprovalSelf(opts *bind.CallOpts, token common.Address, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC1155ApprovalSelf", token, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireERC1155ApprovalSelf is a free data retrieval call binding the contract method 0x0cb17e5c.
//
// Solidity: function requireERC1155ApprovalSelf(address token, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC1155ApprovalSelf(token common.Address, operator common.Address) error {
	return _TrailsUtils.Contract.RequireERC1155ApprovalSelf(&_TrailsUtils.CallOpts, token, operator)
}

// RequireERC1155ApprovalSelf is a free data retrieval call binding the contract method 0x0cb17e5c.
//
// Solidity: function requireERC1155ApprovalSelf(address token, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC1155ApprovalSelf(token common.Address, operator common.Address) error {
	return _TrailsUtils.Contract.RequireERC1155ApprovalSelf(&_TrailsUtils.CallOpts, token, operator)
}

// RequireERC721Approval is a free data retrieval call binding the contract method 0xfe08e7ca.
//
// Solidity: function requireERC721Approval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721Approval(opts *bind.CallOpts, token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721Approval", token, owner, spender, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721Approval is a free data retrieval call binding the contract method 0xfe08e7ca.
//
// Solidity: function requireERC721Approval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721Approval(token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721Approval(&_TrailsUtils.CallOpts, token, owner, spender, tokenId)
}

// RequireERC721Approval is a free data retrieval call binding the contract method 0xfe08e7ca.
//
// Solidity: function requireERC721Approval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721Approval(token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721Approval(&_TrailsUtils.CallOpts, token, owner, spender, tokenId)
}

// RequireERC721ApprovalSelf is a free data retrieval call binding the contract method 0xe246466a.
//
// Solidity: function requireERC721ApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721ApprovalSelf(opts *bind.CallOpts, token common.Address, spender common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721ApprovalSelf", token, spender, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721ApprovalSelf is a free data retrieval call binding the contract method 0xe246466a.
//
// Solidity: function requireERC721ApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721ApprovalSelf(token common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721ApprovalSelf(&_TrailsUtils.CallOpts, token, spender, tokenId)
}

// RequireERC721ApprovalSelf is a free data retrieval call binding the contract method 0xe246466a.
//
// Solidity: function requireERC721ApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721ApprovalSelf(token common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721ApprovalSelf(&_TrailsUtils.CallOpts, token, spender, tokenId)
}

// RequireERC721Owner is a free data retrieval call binding the contract method 0xcda38f4c.
//
// Solidity: function requireERC721Owner(address token, address owner, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721Owner(opts *bind.CallOpts, token common.Address, owner common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721Owner", token, owner, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721Owner is a free data retrieval call binding the contract method 0xcda38f4c.
//
// Solidity: function requireERC721Owner(address token, address owner, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721Owner(token common.Address, owner common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721Owner(&_TrailsUtils.CallOpts, token, owner, tokenId)
}

// RequireERC721Owner is a free data retrieval call binding the contract method 0xcda38f4c.
//
// Solidity: function requireERC721Owner(address token, address owner, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721Owner(token common.Address, owner common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721Owner(&_TrailsUtils.CallOpts, token, owner, tokenId)
}

// RequireERC721OwnerApproval is a free data retrieval call binding the contract method 0x1c0bdfab.
//
// Solidity: function requireERC721OwnerApproval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721OwnerApproval(opts *bind.CallOpts, token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721OwnerApproval", token, owner, spender, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721OwnerApproval is a free data retrieval call binding the contract method 0x1c0bdfab.
//
// Solidity: function requireERC721OwnerApproval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721OwnerApproval(token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerApproval(&_TrailsUtils.CallOpts, token, owner, spender, tokenId)
}

// RequireERC721OwnerApproval is a free data retrieval call binding the contract method 0x1c0bdfab.
//
// Solidity: function requireERC721OwnerApproval(address token, address owner, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721OwnerApproval(token common.Address, owner common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerApproval(&_TrailsUtils.CallOpts, token, owner, spender, tokenId)
}

// RequireERC721OwnerApprovalSelf is a free data retrieval call binding the contract method 0xf38bce15.
//
// Solidity: function requireERC721OwnerApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721OwnerApprovalSelf(opts *bind.CallOpts, token common.Address, spender common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721OwnerApprovalSelf", token, spender, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721OwnerApprovalSelf is a free data retrieval call binding the contract method 0xf38bce15.
//
// Solidity: function requireERC721OwnerApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721OwnerApprovalSelf(token common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerApprovalSelf(&_TrailsUtils.CallOpts, token, spender, tokenId)
}

// RequireERC721OwnerApprovalSelf is a free data retrieval call binding the contract method 0xf38bce15.
//
// Solidity: function requireERC721OwnerApprovalSelf(address token, address spender, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721OwnerApprovalSelf(token common.Address, spender common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerApprovalSelf(&_TrailsUtils.CallOpts, token, spender, tokenId)
}

// RequireERC721OwnerSelf is a free data retrieval call binding the contract method 0x2f2227ba.
//
// Solidity: function requireERC721OwnerSelf(address token, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireERC721OwnerSelf(opts *bind.CallOpts, token common.Address, tokenId *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireERC721OwnerSelf", token, tokenId)

	if err != nil {
		return err
	}

	return err

}

// RequireERC721OwnerSelf is a free data retrieval call binding the contract method 0x2f2227ba.
//
// Solidity: function requireERC721OwnerSelf(address token, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireERC721OwnerSelf(token common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerSelf(&_TrailsUtils.CallOpts, token, tokenId)
}

// RequireERC721OwnerSelf is a free data retrieval call binding the contract method 0x2f2227ba.
//
// Solidity: function requireERC721OwnerSelf(address token, uint256 tokenId) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireERC721OwnerSelf(token common.Address, tokenId *big.Int) error {
	return _TrailsUtils.Contract.RequireERC721OwnerSelf(&_TrailsUtils.CallOpts, token, tokenId)
}

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinBalance(opts *bind.CallOpts, owner common.Address, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinBalance", owner, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinBalance(owner common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalance(&_TrailsUtils.CallOpts, owner, minBalance)
}

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinBalance(owner common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalance(&_TrailsUtils.CallOpts, owner, minBalance)
}

// RequireMinBalanceSelf is a free data retrieval call binding the contract method 0x333f790d.
//
// Solidity: function requireMinBalanceSelf(uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinBalanceSelf(opts *bind.CallOpts, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinBalanceSelf", minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinBalanceSelf is a free data retrieval call binding the contract method 0x333f790d.
//
// Solidity: function requireMinBalanceSelf(uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinBalanceSelf(minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalanceSelf(&_TrailsUtils.CallOpts, minBalance)
}

// RequireMinBalanceSelf is a free data retrieval call binding the contract method 0x333f790d.
//
// Solidity: function requireMinBalanceSelf(uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinBalanceSelf(minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalanceSelf(&_TrailsUtils.CallOpts, minBalance)
}

// RequireMinERC1155Balance is a free data retrieval call binding the contract method 0x505c7ed3.
//
// Solidity: function requireMinERC1155Balance(address token, address owner, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155Balance(opts *bind.CallOpts, token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155Balance", token, owner, tokenId, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155Balance is a free data retrieval call binding the contract method 0x505c7ed3.
//
// Solidity: function requireMinERC1155Balance(address token, address owner, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155Balance(token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155Balance(&_TrailsUtils.CallOpts, token, owner, tokenId, minBalance)
}

// RequireMinERC1155Balance is a free data retrieval call binding the contract method 0x505c7ed3.
//
// Solidity: function requireMinERC1155Balance(address token, address owner, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155Balance(token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155Balance(&_TrailsUtils.CallOpts, token, owner, tokenId, minBalance)
}

// RequireMinERC1155BalanceApproval is a free data retrieval call binding the contract method 0x317df1e8.
//
// Solidity: function requireMinERC1155BalanceApproval(address token, address owner, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceApproval(opts *bind.CallOpts, token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceApproval", token, owner, tokenId, minBalance, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceApproval is a free data retrieval call binding the contract method 0x317df1e8.
//
// Solidity: function requireMinERC1155BalanceApproval(address token, address owner, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceApproval(token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApproval(&_TrailsUtils.CallOpts, token, owner, tokenId, minBalance, operator)
}

// RequireMinERC1155BalanceApproval is a free data retrieval call binding the contract method 0x317df1e8.
//
// Solidity: function requireMinERC1155BalanceApproval(address token, address owner, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceApproval(token common.Address, owner common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApproval(&_TrailsUtils.CallOpts, token, owner, tokenId, minBalance, operator)
}

// RequireMinERC1155BalanceApprovalBatch is a free data retrieval call binding the contract method 0x024ea911.
//
// Solidity: function requireMinERC1155BalanceApprovalBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceApprovalBatch(opts *bind.CallOpts, token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceApprovalBatch", token, owner, tokenIds, minBalances, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceApprovalBatch is a free data retrieval call binding the contract method 0x024ea911.
//
// Solidity: function requireMinERC1155BalanceApprovalBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceApprovalBatch(token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalBatch(&_TrailsUtils.CallOpts, token, owner, tokenIds, minBalances, operator)
}

// RequireMinERC1155BalanceApprovalBatch is a free data retrieval call binding the contract method 0x024ea911.
//
// Solidity: function requireMinERC1155BalanceApprovalBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceApprovalBatch(token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalBatch(&_TrailsUtils.CallOpts, token, owner, tokenIds, minBalances, operator)
}

// RequireMinERC1155BalanceApprovalBatchSelf is a free data retrieval call binding the contract method 0x18c9003a.
//
// Solidity: function requireMinERC1155BalanceApprovalBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceApprovalBatchSelf(opts *bind.CallOpts, token common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceApprovalBatchSelf", token, tokenIds, minBalances, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceApprovalBatchSelf is a free data retrieval call binding the contract method 0x18c9003a.
//
// Solidity: function requireMinERC1155BalanceApprovalBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceApprovalBatchSelf(token common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalBatchSelf(&_TrailsUtils.CallOpts, token, tokenIds, minBalances, operator)
}

// RequireMinERC1155BalanceApprovalBatchSelf is a free data retrieval call binding the contract method 0x18c9003a.
//
// Solidity: function requireMinERC1155BalanceApprovalBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceApprovalBatchSelf(token common.Address, tokenIds []*big.Int, minBalances []*big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalBatchSelf(&_TrailsUtils.CallOpts, token, tokenIds, minBalances, operator)
}

// RequireMinERC1155BalanceApprovalSelf is a free data retrieval call binding the contract method 0xac38d089.
//
// Solidity: function requireMinERC1155BalanceApprovalSelf(address token, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceApprovalSelf(opts *bind.CallOpts, token common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceApprovalSelf", token, tokenId, minBalance, operator)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceApprovalSelf is a free data retrieval call binding the contract method 0xac38d089.
//
// Solidity: function requireMinERC1155BalanceApprovalSelf(address token, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceApprovalSelf(token common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalSelf(&_TrailsUtils.CallOpts, token, tokenId, minBalance, operator)
}

// RequireMinERC1155BalanceApprovalSelf is a free data retrieval call binding the contract method 0xac38d089.
//
// Solidity: function requireMinERC1155BalanceApprovalSelf(address token, uint256 tokenId, uint256 minBalance, address operator) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceApprovalSelf(token common.Address, tokenId *big.Int, minBalance *big.Int, operator common.Address) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceApprovalSelf(&_TrailsUtils.CallOpts, token, tokenId, minBalance, operator)
}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceBatch(opts *bind.CallOpts, token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceBatch", token, owner, tokenIds, minBalances)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceBatch(token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatch(&_TrailsUtils.CallOpts, token, owner, tokenIds, minBalances)
}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address owner, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceBatch(token common.Address, owner common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatch(&_TrailsUtils.CallOpts, token, owner, tokenIds, minBalances)
}

// RequireMinERC1155BalanceBatchSelf is a free data retrieval call binding the contract method 0xfdc25428.
//
// Solidity: function requireMinERC1155BalanceBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceBatchSelf(opts *bind.CallOpts, token common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceBatchSelf", token, tokenIds, minBalances)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceBatchSelf is a free data retrieval call binding the contract method 0xfdc25428.
//
// Solidity: function requireMinERC1155BalanceBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceBatchSelf(token common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatchSelf(&_TrailsUtils.CallOpts, token, tokenIds, minBalances)
}

// RequireMinERC1155BalanceBatchSelf is a free data retrieval call binding the contract method 0xfdc25428.
//
// Solidity: function requireMinERC1155BalanceBatchSelf(address token, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceBatchSelf(token common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatchSelf(&_TrailsUtils.CallOpts, token, tokenIds, minBalances)
}

// RequireMinERC1155BalanceSelf is a free data retrieval call binding the contract method 0x0084290d.
//
// Solidity: function requireMinERC1155BalanceSelf(address token, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceSelf(opts *bind.CallOpts, token common.Address, tokenId *big.Int, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceSelf", token, tokenId, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceSelf is a free data retrieval call binding the contract method 0x0084290d.
//
// Solidity: function requireMinERC1155BalanceSelf(address token, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceSelf(token common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceSelf(&_TrailsUtils.CallOpts, token, tokenId, minBalance)
}

// RequireMinERC1155BalanceSelf is a free data retrieval call binding the contract method 0x0084290d.
//
// Solidity: function requireMinERC1155BalanceSelf(address token, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceSelf(token common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceSelf(&_TrailsUtils.CallOpts, token, tokenId, minBalance)
}

// RequireMinERC20Allowance is a free data retrieval call binding the contract method 0x42ab921e.
//
// Solidity: function requireMinERC20Allowance(address token, address owner, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20Allowance(opts *bind.CallOpts, token common.Address, owner common.Address, spender common.Address, minAllowance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20Allowance", token, owner, spender, minAllowance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20Allowance is a free data retrieval call binding the contract method 0x42ab921e.
//
// Solidity: function requireMinERC20Allowance(address token, address owner, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20Allowance(token common.Address, owner common.Address, spender common.Address, minAllowance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Allowance(&_TrailsUtils.CallOpts, token, owner, spender, minAllowance)
}

// RequireMinERC20Allowance is a free data retrieval call binding the contract method 0x42ab921e.
//
// Solidity: function requireMinERC20Allowance(address token, address owner, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20Allowance(token common.Address, owner common.Address, spender common.Address, minAllowance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Allowance(&_TrailsUtils.CallOpts, token, owner, spender, minAllowance)
}

// RequireMinERC20AllowanceSelf is a free data retrieval call binding the contract method 0x6ef1f0dd.
//
// Solidity: function requireMinERC20AllowanceSelf(address token, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20AllowanceSelf(opts *bind.CallOpts, token common.Address, spender common.Address, minAllowance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20AllowanceSelf", token, spender, minAllowance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20AllowanceSelf is a free data retrieval call binding the contract method 0x6ef1f0dd.
//
// Solidity: function requireMinERC20AllowanceSelf(address token, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20AllowanceSelf(token common.Address, spender common.Address, minAllowance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20AllowanceSelf(&_TrailsUtils.CallOpts, token, spender, minAllowance)
}

// RequireMinERC20AllowanceSelf is a free data retrieval call binding the contract method 0x6ef1f0dd.
//
// Solidity: function requireMinERC20AllowanceSelf(address token, address spender, uint256 minAllowance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20AllowanceSelf(token common.Address, spender common.Address, minAllowance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20AllowanceSelf(&_TrailsUtils.CallOpts, token, spender, minAllowance)
}

// RequireMinERC20Balance is a free data retrieval call binding the contract method 0x8d0a9ea7.
//
// Solidity: function requireMinERC20Balance(address token, address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20Balance(opts *bind.CallOpts, token common.Address, owner common.Address, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20Balance", token, owner, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20Balance is a free data retrieval call binding the contract method 0x8d0a9ea7.
//
// Solidity: function requireMinERC20Balance(address token, address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20Balance(token common.Address, owner common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Balance(&_TrailsUtils.CallOpts, token, owner, minBalance)
}

// RequireMinERC20Balance is a free data retrieval call binding the contract method 0x8d0a9ea7.
//
// Solidity: function requireMinERC20Balance(address token, address owner, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20Balance(token common.Address, owner common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Balance(&_TrailsUtils.CallOpts, token, owner, minBalance)
}

// RequireMinERC20BalanceAllowance is a free data retrieval call binding the contract method 0x9c8d8ea1.
//
// Solidity: function requireMinERC20BalanceAllowance(address token, address owner, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20BalanceAllowance(opts *bind.CallOpts, token common.Address, owner common.Address, spender common.Address, minAmount *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20BalanceAllowance", token, owner, spender, minAmount)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20BalanceAllowance is a free data retrieval call binding the contract method 0x9c8d8ea1.
//
// Solidity: function requireMinERC20BalanceAllowance(address token, address owner, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20BalanceAllowance(token common.Address, owner common.Address, spender common.Address, minAmount *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceAllowance(&_TrailsUtils.CallOpts, token, owner, spender, minAmount)
}

// RequireMinERC20BalanceAllowance is a free data retrieval call binding the contract method 0x9c8d8ea1.
//
// Solidity: function requireMinERC20BalanceAllowance(address token, address owner, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20BalanceAllowance(token common.Address, owner common.Address, spender common.Address, minAmount *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceAllowance(&_TrailsUtils.CallOpts, token, owner, spender, minAmount)
}

// RequireMinERC20BalanceAllowanceSelf is a free data retrieval call binding the contract method 0x00e79cd5.
//
// Solidity: function requireMinERC20BalanceAllowanceSelf(address token, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20BalanceAllowanceSelf(opts *bind.CallOpts, token common.Address, spender common.Address, minAmount *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20BalanceAllowanceSelf", token, spender, minAmount)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20BalanceAllowanceSelf is a free data retrieval call binding the contract method 0x00e79cd5.
//
// Solidity: function requireMinERC20BalanceAllowanceSelf(address token, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20BalanceAllowanceSelf(token common.Address, spender common.Address, minAmount *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceAllowanceSelf(&_TrailsUtils.CallOpts, token, spender, minAmount)
}

// RequireMinERC20BalanceAllowanceSelf is a free data retrieval call binding the contract method 0x00e79cd5.
//
// Solidity: function requireMinERC20BalanceAllowanceSelf(address token, address spender, uint256 minAmount) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20BalanceAllowanceSelf(token common.Address, spender common.Address, minAmount *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceAllowanceSelf(&_TrailsUtils.CallOpts, token, spender, minAmount)
}

// RequireMinERC20BalanceSelf is a free data retrieval call binding the contract method 0x372618cb.
//
// Solidity: function requireMinERC20BalanceSelf(address token, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20BalanceSelf(opts *bind.CallOpts, token common.Address, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20BalanceSelf", token, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20BalanceSelf is a free data retrieval call binding the contract method 0x372618cb.
//
// Solidity: function requireMinERC20BalanceSelf(address token, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20BalanceSelf(token common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceSelf(&_TrailsUtils.CallOpts, token, minBalance)
}

// RequireMinERC20BalanceSelf is a free data retrieval call binding the contract method 0x372618cb.
//
// Solidity: function requireMinERC20BalanceSelf(address token, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20BalanceSelf(token common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20BalanceSelf(&_TrailsUtils.CallOpts, token, minBalance)
}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 expiration) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireNonExpired(opts *bind.CallOpts, expiration *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireNonExpired", expiration)

	if err != nil {
		return err
	}

	return err

}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 expiration) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireNonExpired(expiration *big.Int) error {
	return _TrailsUtils.Contract.RequireNonExpired(&_TrailsUtils.CallOpts, expiration)
}

// RequireNonExpired is a free data retrieval call binding the contract method 0x7f29d538.
//
// Solidity: function requireNonExpired(uint256 expiration) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireNonExpired(expiration *big.Int) error {
	return _TrailsUtils.Contract.RequireNonExpired(&_TrailsUtils.CallOpts, expiration)
}

// HandleSequenceDelegateCall is a paid mutator transaction binding the contract method 0x4c4e814c.
//
// Solidity: function handleSequenceDelegateCall(bytes32 , uint256 , uint256 , uint256 , uint256 , bytes data) returns()
func (_TrailsUtils *TrailsUtilsTransactor) HandleSequenceDelegateCall(opts *bind.TransactOpts, arg0 [32]byte, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, data []byte) (*types.Transaction, error) {
	return _TrailsUtils.contract.Transact(opts, "handleSequenceDelegateCall", arg0, arg1, arg2, arg3, arg4, data)
}

// HandleSequenceDelegateCall is a paid mutator transaction binding the contract method 0x4c4e814c.
//
// Solidity: function handleSequenceDelegateCall(bytes32 , uint256 , uint256 , uint256 , uint256 , bytes data) returns()
func (_TrailsUtils *TrailsUtilsSession) HandleSequenceDelegateCall(arg0 [32]byte, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, data []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HandleSequenceDelegateCall(&_TrailsUtils.TransactOpts, arg0, arg1, arg2, arg3, arg4, data)
}

// HandleSequenceDelegateCall is a paid mutator transaction binding the contract method 0x4c4e814c.
//
// Solidity: function handleSequenceDelegateCall(bytes32 , uint256 , uint256 , uint256 , uint256 , bytes data) returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) HandleSequenceDelegateCall(arg0 [32]byte, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 *big.Int, data []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HandleSequenceDelegateCall(&_TrailsUtils.TransactOpts, arg0, arg1, arg2, arg3, arg4, data)
}

// HydrateExecute is a paid mutator transaction binding the contract method 0x6d42cf03.
//
// Solidity: function hydrateExecute(bytes packedPayload, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsTransactor) HydrateExecute(opts *bind.TransactOpts, packedPayload []byte, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.contract.Transact(opts, "hydrateExecute", packedPayload, hydratePayload)
}

// HydrateExecute is a paid mutator transaction binding the contract method 0x6d42cf03.
//
// Solidity: function hydrateExecute(bytes packedPayload, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsSession) HydrateExecute(packedPayload []byte, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecute(&_TrailsUtils.TransactOpts, packedPayload, hydratePayload)
}

// HydrateExecute is a paid mutator transaction binding the contract method 0x6d42cf03.
//
// Solidity: function hydrateExecute(bytes packedPayload, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) HydrateExecute(packedPayload []byte, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecute(&_TrailsUtils.TransactOpts, packedPayload, hydratePayload)
}

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0x80df36a0.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, bytes hydratePayload, address sweepTarget, address[] tokensToSweep, bool sweepNative) payable returns()
func (_TrailsUtils *TrailsUtilsTransactor) HydrateExecuteAndSweep(opts *bind.TransactOpts, packedPayload []byte, hydratePayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.contract.Transact(opts, "hydrateExecuteAndSweep", packedPayload, hydratePayload, sweepTarget, tokensToSweep, sweepNative)
}

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0x80df36a0.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, bytes hydratePayload, address sweepTarget, address[] tokensToSweep, bool sweepNative) payable returns()
func (_TrailsUtils *TrailsUtilsSession) HydrateExecuteAndSweep(packedPayload []byte, hydratePayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecuteAndSweep(&_TrailsUtils.TransactOpts, packedPayload, hydratePayload, sweepTarget, tokensToSweep, sweepNative)
}

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0x80df36a0.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, bytes hydratePayload, address sweepTarget, address[] tokensToSweep, bool sweepNative) payable returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) HydrateExecuteAndSweep(packedPayload []byte, hydratePayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecuteAndSweep(&_TrailsUtils.TransactOpts, packedPayload, hydratePayload, sweepTarget, tokensToSweep, sweepNative)
}

// Sweep is a paid mutator transaction binding the contract method 0xa7b7ec5a.
//
// Solidity: function sweep(address sweepTarget, address[] tokensToSweep, bool sweepNative) returns()
func (_TrailsUtils *TrailsUtilsTransactor) Sweep(opts *bind.TransactOpts, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.contract.Transact(opts, "sweep", sweepTarget, tokensToSweep, sweepNative)
}

// Sweep is a paid mutator transaction binding the contract method 0xa7b7ec5a.
//
// Solidity: function sweep(address sweepTarget, address[] tokensToSweep, bool sweepNative) returns()
func (_TrailsUtils *TrailsUtilsSession) Sweep(sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.Contract.Sweep(&_TrailsUtils.TransactOpts, sweepTarget, tokensToSweep, sweepNative)
}

// Sweep is a paid mutator transaction binding the contract method 0xa7b7ec5a.
//
// Solidity: function sweep(address sweepTarget, address[] tokensToSweep, bool sweepNative) returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) Sweep(sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool) (*types.Transaction, error) {
	return _TrailsUtils.Contract.Sweep(&_TrailsUtils.TransactOpts, sweepTarget, tokensToSweep, sweepNative)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrailsUtils *TrailsUtilsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrailsUtils.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrailsUtils *TrailsUtilsSession) Receive() (*types.Transaction, error) {
	return _TrailsUtils.Contract.Receive(&_TrailsUtils.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) Receive() (*types.Transaction, error) {
	return _TrailsUtils.Contract.Receive(&_TrailsUtils.TransactOpts)
}

// TrailsUtilsCallAbortedIterator is returned from FilterCallAborted and is used to iterate over the raw logs and unpacked data for CallAborted events raised by the TrailsUtils contract.
type TrailsUtilsCallAbortedIterator struct {
	Event *TrailsUtilsCallAborted // Event containing the contract specifics and raw log

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
func (it *TrailsUtilsCallAbortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrailsUtilsCallAborted)
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
		it.Event = new(TrailsUtilsCallAborted)
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
func (it *TrailsUtilsCallAbortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrailsUtilsCallAbortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrailsUtilsCallAborted represents a CallAborted event raised by the TrailsUtils contract.
type TrailsUtilsCallAborted struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallAborted is a free log retrieval operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_TrailsUtils *TrailsUtilsFilterer) FilterCallAborted(opts *bind.FilterOpts) (*TrailsUtilsCallAbortedIterator, error) {

	logs, sub, err := _TrailsUtils.contract.FilterLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsCallAbortedIterator{contract: _TrailsUtils.contract, event: "CallAborted", logs: logs, sub: sub}, nil
}

// WatchCallAborted is a free log subscription operation binding the contract event 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b.
//
// Solidity: event CallAborted(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_TrailsUtils *TrailsUtilsFilterer) WatchCallAborted(opts *bind.WatchOpts, sink chan<- *TrailsUtilsCallAborted) (event.Subscription, error) {

	logs, sub, err := _TrailsUtils.contract.WatchLogs(opts, "CallAborted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrailsUtilsCallAborted)
				if err := _TrailsUtils.contract.UnpackLog(event, "CallAborted", log); err != nil {
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
func (_TrailsUtils *TrailsUtilsFilterer) ParseCallAborted(log types.Log) (*TrailsUtilsCallAborted, error) {
	event := new(TrailsUtilsCallAborted)
	if err := _TrailsUtils.contract.UnpackLog(event, "CallAborted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrailsUtilsCallFailedIterator is returned from FilterCallFailed and is used to iterate over the raw logs and unpacked data for CallFailed events raised by the TrailsUtils contract.
type TrailsUtilsCallFailedIterator struct {
	Event *TrailsUtilsCallFailed // Event containing the contract specifics and raw log

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
func (it *TrailsUtilsCallFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrailsUtilsCallFailed)
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
		it.Event = new(TrailsUtilsCallFailed)
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
func (it *TrailsUtilsCallFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrailsUtilsCallFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrailsUtilsCallFailed represents a CallFailed event raised by the TrailsUtils contract.
type TrailsUtilsCallFailed struct {
	OpHash     [32]byte
	Index      *big.Int
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallFailed is a free log retrieval operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_TrailsUtils *TrailsUtilsFilterer) FilterCallFailed(opts *bind.FilterOpts) (*TrailsUtilsCallFailedIterator, error) {

	logs, sub, err := _TrailsUtils.contract.FilterLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsCallFailedIterator{contract: _TrailsUtils.contract, event: "CallFailed", logs: logs, sub: sub}, nil
}

// WatchCallFailed is a free log subscription operation binding the contract event 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d.
//
// Solidity: event CallFailed(bytes32 _opHash, uint256 _index, bytes _returnData)
func (_TrailsUtils *TrailsUtilsFilterer) WatchCallFailed(opts *bind.WatchOpts, sink chan<- *TrailsUtilsCallFailed) (event.Subscription, error) {

	logs, sub, err := _TrailsUtils.contract.WatchLogs(opts, "CallFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrailsUtilsCallFailed)
				if err := _TrailsUtils.contract.UnpackLog(event, "CallFailed", log); err != nil {
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
func (_TrailsUtils *TrailsUtilsFilterer) ParseCallFailed(log types.Log) (*TrailsUtilsCallFailed, error) {
	event := new(TrailsUtilsCallFailed)
	if err := _TrailsUtils.contract.UnpackLog(event, "CallFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrailsUtilsCallSkippedIterator is returned from FilterCallSkipped and is used to iterate over the raw logs and unpacked data for CallSkipped events raised by the TrailsUtils contract.
type TrailsUtilsCallSkippedIterator struct {
	Event *TrailsUtilsCallSkipped // Event containing the contract specifics and raw log

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
func (it *TrailsUtilsCallSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrailsUtilsCallSkipped)
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
		it.Event = new(TrailsUtilsCallSkipped)
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
func (it *TrailsUtilsCallSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrailsUtilsCallSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrailsUtilsCallSkipped represents a CallSkipped event raised by the TrailsUtils contract.
type TrailsUtilsCallSkipped struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSkipped is a free log retrieval operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_TrailsUtils *TrailsUtilsFilterer) FilterCallSkipped(opts *bind.FilterOpts) (*TrailsUtilsCallSkippedIterator, error) {

	logs, sub, err := _TrailsUtils.contract.FilterLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsCallSkippedIterator{contract: _TrailsUtils.contract, event: "CallSkipped", logs: logs, sub: sub}, nil
}

// WatchCallSkipped is a free log subscription operation binding the contract event 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b.
//
// Solidity: event CallSkipped(bytes32 _opHash, uint256 _index)
func (_TrailsUtils *TrailsUtilsFilterer) WatchCallSkipped(opts *bind.WatchOpts, sink chan<- *TrailsUtilsCallSkipped) (event.Subscription, error) {

	logs, sub, err := _TrailsUtils.contract.WatchLogs(opts, "CallSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrailsUtilsCallSkipped)
				if err := _TrailsUtils.contract.UnpackLog(event, "CallSkipped", log); err != nil {
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
func (_TrailsUtils *TrailsUtilsFilterer) ParseCallSkipped(log types.Log) (*TrailsUtilsCallSkipped, error) {
	event := new(TrailsUtilsCallSkipped)
	if err := _TrailsUtils.contract.UnpackLog(event, "CallSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrailsUtilsCallSucceededIterator is returned from FilterCallSucceeded and is used to iterate over the raw logs and unpacked data for CallSucceeded events raised by the TrailsUtils contract.
type TrailsUtilsCallSucceededIterator struct {
	Event *TrailsUtilsCallSucceeded // Event containing the contract specifics and raw log

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
func (it *TrailsUtilsCallSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrailsUtilsCallSucceeded)
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
		it.Event = new(TrailsUtilsCallSucceeded)
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
func (it *TrailsUtilsCallSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrailsUtilsCallSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrailsUtilsCallSucceeded represents a CallSucceeded event raised by the TrailsUtils contract.
type TrailsUtilsCallSucceeded struct {
	OpHash [32]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallSucceeded is a free log retrieval operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_TrailsUtils *TrailsUtilsFilterer) FilterCallSucceeded(opts *bind.FilterOpts) (*TrailsUtilsCallSucceededIterator, error) {

	logs, sub, err := _TrailsUtils.contract.FilterLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsCallSucceededIterator{contract: _TrailsUtils.contract, event: "CallSucceeded", logs: logs, sub: sub}, nil
}

// WatchCallSucceeded is a free log subscription operation binding the contract event 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a.
//
// Solidity: event CallSucceeded(bytes32 _opHash, uint256 _index)
func (_TrailsUtils *TrailsUtilsFilterer) WatchCallSucceeded(opts *bind.WatchOpts, sink chan<- *TrailsUtilsCallSucceeded) (event.Subscription, error) {

	logs, sub, err := _TrailsUtils.contract.WatchLogs(opts, "CallSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrailsUtilsCallSucceeded)
				if err := _TrailsUtils.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
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
func (_TrailsUtils *TrailsUtilsFilterer) ParseCallSucceeded(log types.Log) (*TrailsUtilsCallSucceeded, error) {
	event := new(TrailsUtilsCallSucceeded)
	if err := _TrailsUtils.contract.UnpackLog(event, "CallSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrailsUtilsSweepIterator is returned from FilterSweep and is used to iterate over the raw logs and unpacked data for Sweep events raised by the TrailsUtils contract.
type TrailsUtilsSweepIterator struct {
	Event *TrailsUtilsSweep // Event containing the contract specifics and raw log

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
func (it *TrailsUtilsSweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrailsUtilsSweep)
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
		it.Event = new(TrailsUtilsSweep)
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
func (it *TrailsUtilsSweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrailsUtilsSweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrailsUtilsSweep represents a Sweep event raised by the TrailsUtils contract.
type TrailsUtilsSweep struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSweep is a free log retrieval operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address indexed token, address indexed recipient, uint256 amount)
func (_TrailsUtils *TrailsUtilsFilterer) FilterSweep(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*TrailsUtilsSweepIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TrailsUtils.contract.FilterLogs(opts, "Sweep", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TrailsUtilsSweepIterator{contract: _TrailsUtils.contract, event: "Sweep", logs: logs, sub: sub}, nil
}

// WatchSweep is a free log subscription operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address indexed token, address indexed recipient, uint256 amount)
func (_TrailsUtils *TrailsUtilsFilterer) WatchSweep(opts *bind.WatchOpts, sink chan<- *TrailsUtilsSweep, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TrailsUtils.contract.WatchLogs(opts, "Sweep", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrailsUtilsSweep)
				if err := _TrailsUtils.contract.UnpackLog(event, "Sweep", log); err != nil {
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

// ParseSweep is a log parse operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address indexed token, address indexed recipient, uint256 amount)
func (_TrailsUtils *TrailsUtilsFilterer) ParseSweep(log types.Log) (*TrailsUtilsSweep, error) {
	event := new(TrailsUtilsSweep)
	if err := _TrailsUtils.contract.UnpackLog(event, "Sweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
