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
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"handleSequenceDelegateCall\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hydrateExecute\",\"inputs\":[{\"name\":\"packedPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"hydratePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hydrateExecuteAndSweep\",\"inputs\":[{\"name\":\"packedPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sweepTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokensToSweep\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"sweepNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hydratePayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"recoverSapientSignature\",\"inputs\":[{\"name\":\"payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC1155Approval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC1155ApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721Approval\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireERC721ApprovalSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinBalance\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinBalanceSelf\",\"inputs\":[{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155Balance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceBatch\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceBatchSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"minBalances\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC1155BalanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20Allowance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20AllowanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20Balance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireMinERC20BalanceSelf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requireNonExpired\",\"inputs\":[{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sweep\",\"inputs\":[{\"name\":\"sweepTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokensToSweep\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"sweepNative\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CallAborted\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSkipped\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallSucceeded\",\"inputs\":[{\"name\":\"_opHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceSweepFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegateCallFailed\",\"inputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"DelegateCallNotAllowed\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155BalanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155BatchBalanceTooLow\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC1155NotApproved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20AllowanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minAllowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20BalanceTooLow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721NotApproved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Expired\",\"inputs\":[{\"name\":\"expiration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidKind\",\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"LengthMismatch\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NativeBalanceTooLow\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NonTransactionPayload\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotEnoughGas\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLeft\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyDelegateCallAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"_payload\",\"type\":\"tuple\",\"internalType\":\"structPayload.Decoded\",\"components\":[{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"noChainId\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structPayload.Call[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onlyFallback\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"behaviorOnError\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"space\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"imageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parentWallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnknownHydrateDataCommand\",\"inputs\":[{\"name\":\"flag\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a0806040523460295730608052613013908161002e82396080518181816107e80152611c790152f35b5f80fdfe6080604052600436101561001a575b3615610018575f80fd5b005b5f3560e01c806284290d1461015857806303bbbd65146101535780630cb17e5c1461014e57806313792a4a146101495780632ab793d714610144578063333f790d1461013f578063372618cb1461013a57806342ab921e14610135578063476a5f40146101305780634c4e814c1461012b578063505c7ed3146101265780636d42cf03146101215780636ef1f0dd1461011c5780637f29d538146101175780638d0a9ea714610112578063a7b7ec5a1461010d578063bd1e4acb14610108578063e246466a14610103578063fdc25428146100fe5763fe08e7ca0361000e57610c36565b610bac565b610b93565b610ae3565b610a71565b610a50565b6109e5565b6109cc565b61090c565b6108b5565b610796565b610702565b6106b7565b610617565b6105dc565b610580565b610292565b610212565b6101cb565b61017f565b73ffffffffffffffffffffffffffffffffffffffff81160361017b57565b5f80fd5b3461017b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356101bd8161015d565b602435604435913390610f0e565b3461017b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356102098161015d565b60243590611032565b3461017b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356102508161015d565b6024359061025d8261015d565b339061109e565b9181601f8401121561017b5782359167ffffffffffffffff831161017b576020838186019501011161017b57565b3461017b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760043567ffffffffffffffff811161017b5780600401906101207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc823603011261017b5760243567ffffffffffffffff811161017b57610325903690600401610264565b60ff61033085610c50565b1661055857919061036261035460648401356084850135905f5260205260405f2090565b46905f5260205260405f2090565b9060445f9301915b6103748387610c5e565b90508410156104a1576104996001918561039881610392888c610c5e565b90610cdf565b6104896103a482610d24565b61045d6020840135936060810135906103bf60808201610d31565b60c06103cd60a08401610d31565b92013592604051978896602088019a8b93909796959273ffffffffffffffffffffffffffffffffffffffff60e09693610100875260046101008801527f63616c6c0000000000000000000000000000000000000000000000000000000061012088015261014087019a602088015216604086015260608501526080840152151560a0830152151560c08201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610d89565b519020905f5260205260405f2090565b93019261036a565b939250905f5b8381106104b957604051858152602090f35b82939461054261ffff61050a61053a8a8861053061052660ff6103928c61051e8a9e9f61054f9e6104f59160019092919283013560f81c920190565b9416998a948881013560f01c91600290910190565b9b169a8b9781013560f01c91600290910190565b9e1696610c5e565b6040810190610dca565b9092820192610e1b565b92909161118b565b905f5260205260405f2090565b939291906104a7565b7f3c28730f000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461017b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356105be8161015d565b6024356105ca8161015d565b604435916105d78361015d565b61109e565b3461017b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5761001860043533611032565b3461017b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356106558161015d565b60243590339061124d565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc608091011261017b576004356106968161015d565b906024356106a38161015d565b906044356106b08161015d565b9060643590565b3461017b576100186106c836610660565b9291909161134e565b9181601f8401121561017b5782359167ffffffffffffffff831161017b576020808501948460051b01011161017b57565b3461017b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760043561073d8161015d565b6024356107498161015d565b60443567ffffffffffffffff811161017b576107699036906004016106d1565b606435939167ffffffffffffffff851161017b5761078e6100189536906004016106d1565b9490936115de565b3461017b5760c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760a43567ffffffffffffffff811161017b576107e5903690600401610264565b907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8116301461088d57825f939284936040519283928337810184815203915af4610843610e6d565b901561084b57005b610889906040519182917ff7a01e4d000000000000000000000000000000000000000000000000000000008352602060048401526024830190610e9c565b0390fd5b7f827e6c75000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461017b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b576100186004356108f38161015d565b6024356108ff8161015d565b6044359060643592610f0e565b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760043567ffffffffffffffff811161017b57610956903690600401610264565b6024359167ffffffffffffffff831161017b5761097a610018933690600401610264565b9290916119bd565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc606091011261017b576004356109b88161015d565b906024356109c58161015d565b9060443590565b3461017b576100186109dd36610982565b91339061134e565b3461017b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760043580421015610a2157005b7faa2fd925000000000000000000000000000000000000000000000000000000005f526004524260245260445ffd5b3461017b57610018610a6136610982565b9161124d565b8015150361017b57565b3461017b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b57600435610aac8161015d565b6024359067ffffffffffffffff821161017b57610ad06100189236906004016106d1565b9060443592610ade84610a67565b611d40565b60a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b5760043567ffffffffffffffff811161017b57610b2d903690600401610264565b60243591610b3a8361015d565b60443567ffffffffffffffff811161017b57610b5a9036906004016106d1565b60643591610b6783610a67565b6084359567ffffffffffffffff871161017b57610b8b610018973690600401610264565b969095610edf565b3461017b57610018610ba436610982565b913390611e9f565b3461017b5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261017b57600435610be78161015d565b60243567ffffffffffffffff811161017b57610c079036906004016106d1565b916044359267ffffffffffffffff841161017b57610c2c6100189436906004016106d1565b93909233906115de565b3461017b57610018610c4736610660565b92919091611e9f565b3560ff8116810361017b5790565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561017b570180359067ffffffffffffffff821161017b57602001918160051b3603831361017b57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b9190811015610d1f5760051b810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff218136030182121561017b570190565b610cb2565b35610d2e8161015d565b90565b35610d2e81610a67565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b60e0810190811067ffffffffffffffff821117610d8457604052565b610d3b565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d8457604052565b9035907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561017b570180359067ffffffffffffffff821161017b5760200191813603831361017b57565b9093929384831161017b57841161017b578101920390565b67ffffffffffffffff8111610d8457601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b3d15610e97573d90610e7e82610e33565b91610e8c6040519384610d89565b82523d5f602084013e565b606090565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b95610ef29791610ade93949596976119bd565b565b9081602091031261017b575190565b6040513d5f823e3d90fd5b91929092604051907efdd58e00000000000000000000000000000000000000000000000000000000825273ffffffffffffffffffffffffffffffffffffffff8516600483015280602483015260208260448173ffffffffffffffffffffffffffffffffffffffff88165afa91821561102d575f92610ffc575b50828210610f96575050505050565b6040517fbf8ee56100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff948516600482015294909316602485015260448401929092526064830191909152608482015260a490fd5b61101f91925060203d602011611026575b6110178183610d89565b810190610ef4565b905f610f87565b503d61100d565b610f03565b80319082821061104157505050565b73ffffffffffffffffffffffffffffffffffffffff907f1afe1f42000000000000000000000000000000000000000000000000000000005f521660045260245260445260645ffd5b9081602091031261017b5751610d2e81610a67565b6040517fe985e9c500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015284811660248301529190911690602081604481855afa90811561102d575f9161115c575b501561111057505050565b73ffffffffffffffffffffffffffffffffffffffff929183917f628b17c6000000000000000000000000000000000000000000000000000000005f52600452166024521660445260645ffd5b61117e915060203d602011611184575b6111768183610d89565b810190611089565b5f611105565b503d61116c565b92611247917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6101009380604051978895602087019a60808c52600e60a08901527f7374617469632d73656374696f6e00000000000000000000000000000000000060c08901526040880152606087015260c060808701528160e0870152868601375f8582860101520116810103017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282610d89565b51902090565b919091604051917f70a0823100000000000000000000000000000000000000000000000000000000835273ffffffffffffffffffffffffffffffffffffffff8416600484015260208360248173ffffffffffffffffffffffffffffffffffffffff86165afa92831561102d575f9361132d575b508083106112ce5750505050565b60606040519485947f7684050d00000000000000000000000000000000000000000000000000000000865273ffffffffffffffffffffffffffffffffffffffff6004870192816080850197168452166020830152604082015201520390fd5b61134791935060203d602011611026576110178183610d89565b915f6112c0565b6040517fdd62ed3e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015284811660248301529190911693929091602083604481885afa92831561102d575f93611427575b508383106113c7575050505050565b60a495509073ffffffffffffffffffffffffffffffffffffffff8092604051967f73cf338a000000000000000000000000000000000000000000000000000000008852600488015216602486015216604484015260648301526084820152fd5b61144191935060203d602011611026576110178183610d89565b915f6113b8565b67ffffffffffffffff8111610d845760051b60200190565b9061146a82611448565b6114776040519182610d89565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06114a58294611448565b0190602036910137565b8051821015610d1f5760209160051b010190565b60208183031261017b5780519067ffffffffffffffff821161017b57019080601f8301121561017b5781516114f781611448565b926115056040519485610d89565b81845260208085019260051b82010192831161017b57602001905b82821061152d5750505090565b8151815260209182019101611520565b604081016040825282518091526020606083019301905f5b8181106115a25750505060208183039101528281527f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831161017b5760209260051b809284830137010190565b825173ffffffffffffffffffffffffffffffffffffffff16855260209485019490920191600101611555565b9190811015610d1f5760051b0190565b919093929385850361174c579084916115f683611460565b905f5b848110611711575050925f929173ffffffffffffffffffffffffffffffffffffffff61165595604051968795869485937f4e1273f40000000000000000000000000000000000000000000000000000000085526004850161153d565b0392165afa90811561102d575f916116ef575b505f5b838110611679575050505050565b61168381836114af565b5161168f8287866115ce565b351161169d5760010161166b565b826116b782876116b0826116ec976114af565b51936115ce565b7fe3758044000000000000000000000000000000000000000000000000000000005f5260049290925260245235604452606490565b5ffd5b61170b91503d805f833e6117038183610d89565b8101906114c3565b5f611668565b829394506117418261172683600195966114af565b9073ffffffffffffffffffffffffffffffffffffffff169052565b0190869392916115f9565b7fab8b67c6000000000000000000000000000000000000000000000000000000005f526004859052602486905260445ffd5b92919261178a82610e33565b916117986040519384610d89565b82948184528183011161017b578281602093845f960137010152565b9080602083519182815201916020808360051b8301019401925f915b8383106117df57505050505090565b9091929394602080827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0856001950301865288519073ffffffffffffffffffffffffffffffffffffffff8251168152828201518382015260c080611852604085015160e0604086015260e0850190610e9c565b936060810151606085015260808101511515608085015260a0810151151560a08501520151910152970193019301919392906117d0565b90602080835192838152019201905f5b8181106118a65750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611899565b805160ff168252610d2e9160208281015115159082015261010061193061190a604085015161012060408601526101208501906117b4565b606085015160608501526080850151608085015260a085015184820360a0860152610e9c565b9260c081015160c084015260e081015160e0840152015190610100818403910152611889565b61196e604092959493956060835260608301906118d2565b9460208201520152565b610d2e9392606092825260208201528160408201520190610e9c565b6119aa610d2e94926060835260608301906118d2565b9260208201526040818403910152610e9c565b906119c79161217a565b6119f26119d38261243b565b6119de36868661177e565b60208151910120905f5260205260405f2090565b916119fd8482612547565b915f926040850192835151975f5b898110611a1f575b50505050505050505050565b808414611d29575b611a328187516114af565b5196611a4160a0890151151590565b80611d21575b611ce157505f966060810151801590811580611cd8575b611ca057608083018051151580611c61575b611c34575115611beb5790611ab191611a9d845173ffffffffffffffffffffffffffffffffffffffff1690565b9115611be657505a5b604084015191612b1c565b15611af4575b50604080518a815260208101839052600192917f5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a91a15b01611a0b565b60c001805115611ba9576001815114611b6a5751600214611b15575f611ab7565b9796505050505050507fc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b9250611b58611b4c612b2d565b60405193849384611978565b0390a15f808080808080808080611a13565b5087610889611b77612b2d565b6040519384937f7f6b0bb100000000000000000000000000000000000000000000000000000000855260048501611994565b509550600180967f115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d818b611bde611b4c612b2d565b0390a1611aee565b611aa6565b611c2991611c0d845173ffffffffffffffffffffffffffffffffffffffff1690565b916020850151915f14611c2e57505a905b604085015192612b0b565b611ab1565b90611c1e565b7f230d1ccc000000000000000000000000000000000000000000000000000000005f52600485905260245ffd5b5073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163014611a70565b838b6108895a6040519384937f2139527400000000000000000000000000000000000000000000000000000000855260048501611956565b50805a10611a5e565b604080518b815260208101849052919850600192917f9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b9181908101611bde565b508015611a47565b91928193611d3892868a6125ac565b929091611a27565b91929173ffffffffffffffffffffffffffffffffffffffff8116611e84575033925b5f5b818110611dc55750505080611dbc575b611d7b5750565b5f8080809347905af1611d8c610e6d565b5015611d9457565b7fd92fb848000000000000000000000000000000000000000000000000000000005f5260045ffd5b50471515611d74565b611df4611ddb611dd68385876115ce565b610d24565b73ffffffffffffffffffffffffffffffffffffffff1690565b90611e09611ddb611ddb611dd68487896115ce565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529290602090849060249082905afa801561102d5787611e5e926001955f93611e64575b50612b47565b01611d64565b611e7d91935060203d8111611026576110178183610d89565b915f611e58565b92611d62565b9081602091031261017b5751610d2e8161015d565b919073ffffffffffffffffffffffffffffffffffffffff83166040517f081812fc00000000000000000000000000000000000000000000000000000000815260208180611ef489600483019190602083019252565b0381855afa90811561102d575f91612020575b5073ffffffffffffffffffffffffffffffffffffffff808516911614159081611f93575b50611f365750505050565b6040517f6a7e2cc800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9384166004820152602481019490945282166044840152166064820152608490fd5b6040517fe985e9c500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152851660248201529150602090829060449082905afa90811561102d575f91612001575b50155f611f2b565b61201a915060203d602011611184576111768183610d89565b5f611ff9565b612042915060203d602011612048575b61203a8183610d89565b810190611e8a565b5f611f07565b503d612030565b60405190610120820182811067ffffffffffffffff821117610d84576040526060610100835f81525f60208201528260408201525f838201525f60808201528260a08201525f60c08201525f60e08201520152565b906120ae82611448565b6120bb6040519182610d89565b8281527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06120e98294611448565b01905f5b8281106120f957505050565b60209060405161210881610d68565b5f81525f83820152606060408201525f60608201525f60808201525f60a08201525f60c0820152828285010152016120ed565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b9190820180921161217557565b61213b565b91909161218561204f565b5f815292600190823560f81c827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8183160161242457505f60608701525b60076121d260ff831660011c90565b16806123d3575b5060108181160361239e57506001925b6121f2846120a4565b92604087019384525f905b85821061220c57505050505050565b8281013560f81c9060010190918160018085160361237c5750612250306122348389516114af565b519073ffffffffffffffffffffffffffffffffffffffff169052565b60028084161461235c575b60048084161461230e575b6008808416146122db575b6122c46122be60c08561229d8a60806122948860108060019c9d161493516114af565b51019015159052565b6122b48a60a06122948860208087161493516114af565b1660061c60031690565b60ff1690565b60c06122d18389516114af565b51015201906121fd565b6001916122c4906122be9060c090878101359060200196906060612300878d516114af565b510152959450505050612271565b90612356908481013560e81c9060030161234f61233661232e8484612168565b838a8a610e1b565b91906040612345888d516114af565b510192369161177e565b9052612168565b90612266565b908381013590602001919060206123748389516114af565b51015261225b565b61239992508481013560601c9060140192906122348389516114af565b612250565b6020908116036123c05761ffff918381013560f01c906002015b9216926121e9565b60ff918381013560f81c906001016123b8565b612417919385929190928160031b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001821b019185013590610100031c16920190565b929060808701525f6121d9565b80850135606090811c9088015260140192506121c3565b60208101511561252c5761045d6112476124ea5f5b60405160208101917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f83527f4aa45ca7ad825ceb1bf35643f0a58c295239df563b1b565c2485f96477c5631860408301527f2a80e1ef1d7842f27f2e6be0972bb708b9a135c38860dbe73c27c3486c34f4de606083015260808201523060a082015260a081526124e160c082610d89565b51902093612c53565b60405192839160208301958690916042927f19010000000000000000000000000000000000000000000000000000000000008352600283015260228201520190565b61045d6112476124ea46612450565b90821015610d1f570190565b919015612558576001913560f81c90565b5f91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146121755760010190565b94909193926125c6905b8381013560f81c91600190910190565b60ff819216918215612a865785918891600b85116129f257508086013560f01c93600290910192908390600181036126205750509260406126106125c6958261261b9501516114af565b510151903091612e45565b6125b6565b6002810361264b5750509260406126406125c6958261261b9501516114af565b510151903391612e45565b6003810361267657505092604061266b6125c6958261261b9501516114af565b510151903291612e45565b600481036126a15750509260406126966125c6958261261b9501516114af565b510151904791612e2c565b600581036126cd5750509260406126c16125c6958261261b9501516114af565b51015190333191612e2c565b600681036126f95750509260406126ed6125c6958261261b9501516114af565b51015190323191612e2c565b6007810361273557506125c6949361261b93508781013560601c9250601401905092604061272a89828d01516114af565b510151913191612e2c565b600881036127f157505050508381013560601c9060140190604061275c87828b01516114af565b5101516040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529093909160209083908173ffffffffffffffffffffffffffffffffffffffff81602481015b0392165afa90811561102d576125c69461261b935f936127d1575b50612e2c565b6127ea91935060203d8111611026576110178183610d89565b915f6127cb565b6009810361287057505050508381013560601c9060140190604061281887828b01516114af565b5101516040517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529093909160209083908173ffffffffffffffffffffffffffffffffffffffff81602481016127b0565b600a810361290757505050508381013560601c9060140190604061289787828b01516114af565b5101516040517f70a08231000000000000000000000000000000000000000000000000000000008152326004820152909390916020908390602490829073ffffffffffffffffffffffffffffffffffffffff165afa90811561102d576125c69461261b935f936127d15750612e2c565b93949293600b14612920575b505050506125c6906125b6565b6040945073ffffffffffffffffffffffffffffffffffffffff92916020916129c791612967918891612971918c81013560601c906014019094818e013560601c9160140190565b99909a01516114af565b510151966040519586809481937f70a082310000000000000000000000000000000000000000000000000000000083526004830191909173ffffffffffffffffffffffffffffffffffffffff6020820193169052565b0392165afa90811561102d576125c6946129e7935f936127d15750612e2c565b9050855f8581612913565b9392505050600c8103612a1657506125c69061261b336122348760408b01516114af565b600d8103612a3557506125c69061261b326122348760408b01516114af565b600e8103612a5b57506125c690476020612a538760408b01516114af565b5101526125b6565b7f324dd124000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b96505050929150838281105f14612ae457612ad4612aae6122be92612ae095612ada9561253b565b357fff000000000000000000000000000000000000000000000000000000000000001690565b60f81c90565b9261257f565b9190565b5050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90565b915f9391849360208451940192f190565b915f9291839260208351930191f490565b3d90604051916020818401016040528083525f602084013e565b9173ffffffffffffffffffffffffffffffffffffffff604051927fa9059cbb000000000000000000000000000000000000000000000000000000005f521660045260245260205f60448180865af160015f5114811615612bf1575b60409190915215612bb05750565b7f5274afe7000000000000000000000000000000000000000000000000000000005f5273ffffffffffffffffffffffffffffffffffffffff1660045260245ffd5b6001811516612c07573d15833b15151616612ba2565b503d5f823e3d90fd5b80516020909101905f5b818110612c275750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101612c1a565b610100810151604051612c6e8161045d602082018095612c10565b51902090612c7d815160ff1690565b60ff811680612cf657505090611247612c996040840151612ea3565b9261045d60806060830151920151936040519485936020850197889094939260809260a08301967f11e1e4079a79a66e4ade50033cfe2678cdd5341d2dfe5ef9513edb1a0be147a284526020840152604083015260608201520152565b60018103612d5457505060a001518051602091820120604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46693810193845290810191909152606081019290925290611247816080810161045d565b60028103612daa57505060c00151604080517f11fdeb7e8373a1aa96bfac8d0ea91526b2c5d15e5cee20e0543e780258f3e8e460208201908152918101929092526060820192909252611247816080810161045d565b600303612dfe575060e00151604080517fe19a3b94fc3c7ece3f890d98a99bc422615537a08dea0603fa8425867d87d46660208201908152918101929092526060820192909252611247816080810161045d565b7f04818320000000000000000000000000000000000000000000000000000000005f5260ff1660045260245ffd5b908151602082019082821091111761017b570160200152565b908151601482019082821091111761017b576020910101906bffffffffffffffffffffffff8251169060601b179052565b80516020909101905f5b818110612e8d5750505090565b8251845260209384019390920191600101612e80565b9081517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0612ee9612ed383611448565b92612ee16040519485610d89565b808452611448565b013660208301375f5b8351811015612fc45780612f08600192866114af565b5173ffffffffffffffffffffffffffffffffffffffff81511690602081015190604081015160208151910120906060810151608082015115159060c060a08401511515930151936040519560208701977f0603985259a953da1f65a522f589c17bd1d0117ec1d3abb7c0788aef251ef437895260408801526060870152608086015260a085015260c084015260e08301526101008201526101008152612fb061012082610d89565b519020612fbd82856114af565b5201612ef2565b509091506040516112478161045d602082018095612e7656fea264697066735822122001031968cb5ed5075c2647347ed676838da85e531feafbedf6b1c523eefc070864736f6c634300081e0033",
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

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinBalance(opts *bind.CallOpts, wallet common.Address, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinBalance", wallet, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinBalance(wallet common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalance(&_TrailsUtils.CallOpts, wallet, minBalance)
}

// RequireMinBalance is a free data retrieval call binding the contract method 0x03bbbd65.
//
// Solidity: function requireMinBalance(address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinBalance(wallet common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinBalance(&_TrailsUtils.CallOpts, wallet, minBalance)
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
// Solidity: function requireMinERC1155Balance(address token, address wallet, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155Balance(opts *bind.CallOpts, token common.Address, wallet common.Address, tokenId *big.Int, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155Balance", token, wallet, tokenId, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155Balance is a free data retrieval call binding the contract method 0x505c7ed3.
//
// Solidity: function requireMinERC1155Balance(address token, address wallet, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155Balance(token common.Address, wallet common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155Balance(&_TrailsUtils.CallOpts, token, wallet, tokenId, minBalance)
}

// RequireMinERC1155Balance is a free data retrieval call binding the contract method 0x505c7ed3.
//
// Solidity: function requireMinERC1155Balance(address token, address wallet, uint256 tokenId, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155Balance(token common.Address, wallet common.Address, tokenId *big.Int, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155Balance(&_TrailsUtils.CallOpts, token, wallet, tokenId, minBalance)
}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address wallet, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC1155BalanceBatch(opts *bind.CallOpts, token common.Address, wallet common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC1155BalanceBatch", token, wallet, tokenIds, minBalances)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address wallet, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC1155BalanceBatch(token common.Address, wallet common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatch(&_TrailsUtils.CallOpts, token, wallet, tokenIds, minBalances)
}

// RequireMinERC1155BalanceBatch is a free data retrieval call binding the contract method 0x476a5f40.
//
// Solidity: function requireMinERC1155BalanceBatch(address token, address wallet, uint256[] tokenIds, uint256[] minBalances) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC1155BalanceBatch(token common.Address, wallet common.Address, tokenIds []*big.Int, minBalances []*big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC1155BalanceBatch(&_TrailsUtils.CallOpts, token, wallet, tokenIds, minBalances)
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
// Solidity: function requireMinERC20Balance(address token, address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCaller) RequireMinERC20Balance(opts *bind.CallOpts, token common.Address, wallet common.Address, minBalance *big.Int) error {
	var out []interface{}
	err := _TrailsUtils.contract.Call(opts, &out, "requireMinERC20Balance", token, wallet, minBalance)

	if err != nil {
		return err
	}

	return err

}

// RequireMinERC20Balance is a free data retrieval call binding the contract method 0x8d0a9ea7.
//
// Solidity: function requireMinERC20Balance(address token, address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsSession) RequireMinERC20Balance(token common.Address, wallet common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Balance(&_TrailsUtils.CallOpts, token, wallet, minBalance)
}

// RequireMinERC20Balance is a free data retrieval call binding the contract method 0x8d0a9ea7.
//
// Solidity: function requireMinERC20Balance(address token, address wallet, uint256 minBalance) view returns()
func (_TrailsUtils *TrailsUtilsCallerSession) RequireMinERC20Balance(token common.Address, wallet common.Address, minBalance *big.Int) error {
	return _TrailsUtils.Contract.RequireMinERC20Balance(&_TrailsUtils.CallOpts, token, wallet, minBalance)
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

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0xbd1e4acb.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, address sweepTarget, address[] tokensToSweep, bool sweepNative, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsTransactor) HydrateExecuteAndSweep(opts *bind.TransactOpts, packedPayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.contract.Transact(opts, "hydrateExecuteAndSweep", packedPayload, sweepTarget, tokensToSweep, sweepNative, hydratePayload)
}

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0xbd1e4acb.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, address sweepTarget, address[] tokensToSweep, bool sweepNative, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsSession) HydrateExecuteAndSweep(packedPayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecuteAndSweep(&_TrailsUtils.TransactOpts, packedPayload, sweepTarget, tokensToSweep, sweepNative, hydratePayload)
}

// HydrateExecuteAndSweep is a paid mutator transaction binding the contract method 0xbd1e4acb.
//
// Solidity: function hydrateExecuteAndSweep(bytes packedPayload, address sweepTarget, address[] tokensToSweep, bool sweepNative, bytes hydratePayload) payable returns()
func (_TrailsUtils *TrailsUtilsTransactorSession) HydrateExecuteAndSweep(packedPayload []byte, sweepTarget common.Address, tokensToSweep []common.Address, sweepNative bool, hydratePayload []byte) (*types.Transaction, error) {
	return _TrailsUtils.Contract.HydrateExecuteAndSweep(&_TrailsUtils.TransactOpts, packedPayload, sweepTarget, tokensToSweep, sweepNative, hydratePayload)
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
