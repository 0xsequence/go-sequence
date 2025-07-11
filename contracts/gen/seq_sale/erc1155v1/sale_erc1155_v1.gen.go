// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package seq_sale_erc1155

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

// Attestation is an auto generated low-level Go binding around an user-defined struct.
type Attestation struct {
	ApprovedSigner  common.Address
	IdentityType    [4]byte
	IssuerHash      [32]byte
	AudienceHash    [32]byte
	ApplicationData []byte
	AuthData        AuthData
}

// AuthData is an auto generated low-level Go binding around an user-defined struct.
type AuthData struct {
	RedirectUrl string
	IssuedAt    uint64
}

// IERC1155SaleFunctionsGlobalSaleDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC1155SaleFunctionsGlobalSaleDetails struct {
	MinTokenId      *big.Int
	MaxTokenId      *big.Int
	Cost            *big.Int
	RemainingSupply *big.Int
	StartTime       uint64
	EndTime         uint64
	MerkleRoot      [32]byte
}

// IERC1155SaleFunctionsSaleDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC1155SaleFunctionsSaleDetails struct {
	Cost            *big.Int
	RemainingSupply *big.Int
	StartTime       uint64
	EndTime         uint64
	MerkleRoot      [32]byte
}

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

// SaleMetaData contains all meta data concerning the Sale contract.
var SaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GlobalSaleInactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InsufficientSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSaleDetails\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenIds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"MerkleProofInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SaleInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"GlobalSaleDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ItemsMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"TokenSaleDetailsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"approvedSigner\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"identityType\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"issuerHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"audienceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"applicationData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"redirectUrl\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issuedAt\",\"type\":\"uint64\"}],\"internalType\":\"structAuthData\",\"name\":\"authData\",\"type\":\"tuple\"}],\"internalType\":\"structAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"delegateCall\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"onlyFallback\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"behaviorOnError\",\"type\":\"uint256\"}],\"internalType\":\"structPayload.Call\",\"name\":\"call\",\"type\":\"tuple\"}],\"name\":\"acceptImplicitRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"checkMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalSaleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIERC1155SaleFunctions.GlobalSaleDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"items\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implicitModeValidator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"implicitModeProjectId\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotal\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setGlobalSaleDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"setImplicitModeProjectId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"setImplicitModeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentTokenAddr\",\"type\":\"address\"}],\"name\":\"setPaymentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setTokenSaleDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"costs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"remainingSupplies\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64[]\",\"name\":\"startTimes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"endTimes\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleRoots\",\"type\":\"bytes32[]\"}],\"name\":\"setTokenSaleDetailsBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenSaleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIERC1155SaleFunctions.SaleDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"tokenSaleDetailsBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIERC1155SaleFunctions.SaleDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SaleABI is the input ABI used to generate the binding from.
// Deprecated: Use SaleMetaData.ABI instead.
var SaleABI = SaleMetaData.ABI

// Sale is an auto generated Go binding around an Ethereum contract.
type Sale struct {
	SaleCaller     // Read-only binding to the contract
	SaleTransactor // Write-only binding to the contract
	SaleFilterer   // Log filterer for contract events
}

// SaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleSession struct {
	Contract     *Sale             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleCallerSession struct {
	Contract *SaleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleTransactorSession struct {
	Contract     *SaleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleRaw struct {
	Contract *Sale // Generic contract binding to access the raw methods on
}

// SaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleCallerRaw struct {
	Contract *SaleCaller // Generic read-only contract binding to access the raw methods on
}

// SaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleTransactorRaw struct {
	Contract *SaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSale creates a new instance of Sale, bound to a specific deployed contract.
func NewSale(address common.Address, backend bind.ContractBackend) (*Sale, error) {
	contract, err := bindSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sale{SaleCaller: SaleCaller{contract: contract}, SaleTransactor: SaleTransactor{contract: contract}, SaleFilterer: SaleFilterer{contract: contract}}, nil
}

// NewSaleCaller creates a new read-only instance of Sale, bound to a specific deployed contract.
func NewSaleCaller(address common.Address, caller bind.ContractCaller) (*SaleCaller, error) {
	contract, err := bindSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleCaller{contract: contract}, nil
}

// NewSaleTransactor creates a new write-only instance of Sale, bound to a specific deployed contract.
func NewSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleTransactor, error) {
	contract, err := bindSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleTransactor{contract: contract}, nil
}

// NewSaleFilterer creates a new log filterer instance of Sale, bound to a specific deployed contract.
func NewSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleFilterer, error) {
	contract, err := bindSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleFilterer{contract: contract}, nil
}

// bindSale binds a generic wrapper to an already deployed contract.
func bindSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SaleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.SaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.SaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sale *SaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sale *SaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sale *SaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sale.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sale *SaleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sale *SaleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sale.Contract.DEFAULTADMINROLE(&_Sale.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Sale *SaleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Sale.Contract.DEFAULTADMINROLE(&_Sale.CallOpts)
}

// AcceptImplicitRequest is a free data retrieval call binding the contract method 0x9d043a66.
//
// Solidity: function acceptImplicitRequest(address wallet, (address,bytes4,bytes32,bytes32,bytes,(string,uint64)) attestation, (address,uint256,bytes,uint256,bool,bool,uint256) call) view returns(bytes32)
func (_Sale *SaleCaller) AcceptImplicitRequest(opts *bind.CallOpts, wallet common.Address, attestation Attestation, call PayloadCall) ([32]byte, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "acceptImplicitRequest", wallet, attestation, call)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AcceptImplicitRequest is a free data retrieval call binding the contract method 0x9d043a66.
//
// Solidity: function acceptImplicitRequest(address wallet, (address,bytes4,bytes32,bytes32,bytes,(string,uint64)) attestation, (address,uint256,bytes,uint256,bool,bool,uint256) call) view returns(bytes32)
func (_Sale *SaleSession) AcceptImplicitRequest(wallet common.Address, attestation Attestation, call PayloadCall) ([32]byte, error) {
	return _Sale.Contract.AcceptImplicitRequest(&_Sale.CallOpts, wallet, attestation, call)
}

// AcceptImplicitRequest is a free data retrieval call binding the contract method 0x9d043a66.
//
// Solidity: function acceptImplicitRequest(address wallet, (address,bytes4,bytes32,bytes32,bytes,(string,uint64)) attestation, (address,uint256,bytes,uint256,bool,bool,uint256) call) view returns(bytes32)
func (_Sale *SaleCallerSession) AcceptImplicitRequest(wallet common.Address, attestation Attestation, call PayloadCall) ([32]byte, error) {
	return _Sale.Contract.AcceptImplicitRequest(&_Sale.CallOpts, wallet, attestation, call)
}

// CheckMerkleProof is a free data retrieval call binding the contract method 0xbad43661.
//
// Solidity: function checkMerkleProof(bytes32 root, bytes32[] proof, address addr, bytes32 salt) view returns(bool)
func (_Sale *SaleCaller) CheckMerkleProof(opts *bind.CallOpts, root [32]byte, proof [][32]byte, addr common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "checkMerkleProof", root, proof, addr, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMerkleProof is a free data retrieval call binding the contract method 0xbad43661.
//
// Solidity: function checkMerkleProof(bytes32 root, bytes32[] proof, address addr, bytes32 salt) view returns(bool)
func (_Sale *SaleSession) CheckMerkleProof(root [32]byte, proof [][32]byte, addr common.Address, salt [32]byte) (bool, error) {
	return _Sale.Contract.CheckMerkleProof(&_Sale.CallOpts, root, proof, addr, salt)
}

// CheckMerkleProof is a free data retrieval call binding the contract method 0xbad43661.
//
// Solidity: function checkMerkleProof(bytes32 root, bytes32[] proof, address addr, bytes32 salt) view returns(bool)
func (_Sale *SaleCallerSession) CheckMerkleProof(root [32]byte, proof [][32]byte, addr common.Address, salt [32]byte) (bool, error) {
	return _Sale.Contract.CheckMerkleProof(&_Sale.CallOpts, root, proof, addr, salt)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sale *SaleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sale *SaleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sale.Contract.GetRoleAdmin(&_Sale.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Sale *SaleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Sale.Contract.GetRoleAdmin(&_Sale.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Sale *SaleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Sale *SaleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Sale.Contract.GetRoleMember(&_Sale.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Sale *SaleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Sale.Contract.GetRoleMember(&_Sale.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Sale *SaleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Sale *SaleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Sale.Contract.GetRoleMemberCount(&_Sale.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Sale *SaleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Sale.Contract.GetRoleMemberCount(&_Sale.CallOpts, role)
}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCaller) GlobalSaleDetails(opts *bind.CallOpts) (IERC1155SaleFunctionsGlobalSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "globalSaleDetails")

	if err != nil {
		return *new(IERC1155SaleFunctionsGlobalSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC1155SaleFunctionsGlobalSaleDetails)).(*IERC1155SaleFunctionsGlobalSaleDetails)

	return out0, err

}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleSession) GlobalSaleDetails() (IERC1155SaleFunctionsGlobalSaleDetails, error) {
	return _Sale.Contract.GlobalSaleDetails(&_Sale.CallOpts)
}

// GlobalSaleDetails is a free data retrieval call binding the contract method 0x119cd50c.
//
// Solidity: function globalSaleDetails() view returns((uint256,uint256,uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCallerSession) GlobalSaleDetails() (IERC1155SaleFunctionsGlobalSaleDetails, error) {
	return _Sale.Contract.GlobalSaleDetails(&_Sale.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sale *SaleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sale *SaleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sale.Contract.HasRole(&_Sale.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Sale *SaleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Sale.Contract.HasRole(&_Sale.CallOpts, role, account)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleSession) PaymentToken() (common.Address, error) {
	return _Sale.Contract.PaymentToken(&_Sale.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Sale *SaleCallerSession) PaymentToken() (common.Address, error) {
	return _Sale.Contract.PaymentToken(&_Sale.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sale.Contract.SupportsInterface(&_Sale.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Sale *SaleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Sale.Contract.SupportsInterface(&_Sale.CallOpts, interfaceId)
}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCaller) TokenSaleDetails(opts *bind.CallOpts, tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenSaleDetails", tokenId)

	if err != nil {
		return *new(IERC1155SaleFunctionsSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC1155SaleFunctionsSaleDetails)).(*IERC1155SaleFunctionsSaleDetails)

	return out0, err

}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleSession) TokenSaleDetails(tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetails(&_Sale.CallOpts, tokenId)
}

// TokenSaleDetails is a free data retrieval call binding the contract method 0x0869678c.
//
// Solidity: function tokenSaleDetails(uint256 tokenId) view returns((uint256,uint256,uint64,uint64,bytes32))
func (_Sale *SaleCallerSession) TokenSaleDetails(tokenId *big.Int) (IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetails(&_Sale.CallOpts, tokenId)
}

// TokenSaleDetailsBatch is a free data retrieval call binding the contract method 0xff81434e.
//
// Solidity: function tokenSaleDetailsBatch(uint256[] tokenIds) view returns((uint256,uint256,uint64,uint64,bytes32)[])
func (_Sale *SaleCaller) TokenSaleDetailsBatch(opts *bind.CallOpts, tokenIds []*big.Int) ([]IERC1155SaleFunctionsSaleDetails, error) {
	var out []interface{}
	err := _Sale.contract.Call(opts, &out, "tokenSaleDetailsBatch", tokenIds)

	if err != nil {
		return *new([]IERC1155SaleFunctionsSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]IERC1155SaleFunctionsSaleDetails)).(*[]IERC1155SaleFunctionsSaleDetails)

	return out0, err

}

// TokenSaleDetailsBatch is a free data retrieval call binding the contract method 0xff81434e.
//
// Solidity: function tokenSaleDetailsBatch(uint256[] tokenIds) view returns((uint256,uint256,uint64,uint64,bytes32)[])
func (_Sale *SaleSession) TokenSaleDetailsBatch(tokenIds []*big.Int) ([]IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetailsBatch(&_Sale.CallOpts, tokenIds)
}

// TokenSaleDetailsBatch is a free data retrieval call binding the contract method 0xff81434e.
//
// Solidity: function tokenSaleDetailsBatch(uint256[] tokenIds) view returns((uint256,uint256,uint64,uint64,bytes32)[])
func (_Sale *SaleCallerSession) TokenSaleDetailsBatch(tokenIds []*big.Int) ([]IERC1155SaleFunctionsSaleDetails, error) {
	return _Sale.Contract.TokenSaleDetailsBatch(&_Sale.CallOpts, tokenIds)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sale *SaleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.GrantRole(&_Sale.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.GrantRole(&_Sale.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address owner, address items, address implicitModeValidator, bytes32 implicitModeProjectId) returns()
func (_Sale *SaleTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, items common.Address, implicitModeValidator common.Address, implicitModeProjectId [32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "initialize", owner, items, implicitModeValidator, implicitModeProjectId)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address owner, address items, address implicitModeValidator, bytes32 implicitModeProjectId) returns()
func (_Sale *SaleSession) Initialize(owner common.Address, items common.Address, implicitModeValidator common.Address, implicitModeProjectId [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Initialize(&_Sale.TransactOpts, owner, items, implicitModeValidator, implicitModeProjectId)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address owner, address items, address implicitModeValidator, bytes32 implicitModeProjectId) returns()
func (_Sale *SaleTransactorSession) Initialize(owner common.Address, items common.Address, implicitModeValidator common.Address, implicitModeProjectId [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Initialize(&_Sale.TransactOpts, owner, items, implicitModeValidator, implicitModeProjectId)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address expectedPaymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, expectedPaymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "mint", to, tokenIds, amounts, data, expectedPaymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address expectedPaymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleSession) Mint(to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, expectedPaymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, tokenIds, amounts, data, expectedPaymentToken, maxTotal, proof)
}

// Mint is a paid mutator transaction binding the contract method 0x60e606f6.
//
// Solidity: function mint(address to, uint256[] tokenIds, uint256[] amounts, bytes data, address expectedPaymentToken, uint256 maxTotal, bytes32[] proof) payable returns()
func (_Sale *SaleTransactorSession) Mint(to common.Address, tokenIds []*big.Int, amounts []*big.Int, data []byte, expectedPaymentToken common.Address, maxTotal *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.Mint(&_Sale.TransactOpts, to, tokenIds, amounts, data, expectedPaymentToken, maxTotal, proof)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Sale *SaleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.RenounceRole(&_Sale.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.RenounceRole(&_Sale.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sale *SaleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.RevokeRole(&_Sale.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Sale *SaleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Sale.Contract.RevokeRole(&_Sale.TransactOpts, role, account)
}

// SetGlobalSaleDetails is a paid mutator transaction binding the contract method 0xc81ee650.
//
// Solidity: function setGlobalSaleDetails(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactor) SetGlobalSaleDetails(opts *bind.TransactOpts, minTokenId *big.Int, maxTokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setGlobalSaleDetails", minTokenId, maxTokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetGlobalSaleDetails is a paid mutator transaction binding the contract method 0xc81ee650.
//
// Solidity: function setGlobalSaleDetails(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleSession) SetGlobalSaleDetails(minTokenId *big.Int, maxTokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetGlobalSaleDetails(&_Sale.TransactOpts, minTokenId, maxTokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetGlobalSaleDetails is a paid mutator transaction binding the contract method 0xc81ee650.
//
// Solidity: function setGlobalSaleDetails(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactorSession) SetGlobalSaleDetails(minTokenId *big.Int, maxTokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetGlobalSaleDetails(&_Sale.TransactOpts, minTokenId, maxTokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetImplicitModeProjectId is a paid mutator transaction binding the contract method 0xed4c2ac7.
//
// Solidity: function setImplicitModeProjectId(bytes32 projectId) returns()
func (_Sale *SaleTransactor) SetImplicitModeProjectId(opts *bind.TransactOpts, projectId [32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setImplicitModeProjectId", projectId)
}

// SetImplicitModeProjectId is a paid mutator transaction binding the contract method 0xed4c2ac7.
//
// Solidity: function setImplicitModeProjectId(bytes32 projectId) returns()
func (_Sale *SaleSession) SetImplicitModeProjectId(projectId [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetImplicitModeProjectId(&_Sale.TransactOpts, projectId)
}

// SetImplicitModeProjectId is a paid mutator transaction binding the contract method 0xed4c2ac7.
//
// Solidity: function setImplicitModeProjectId(bytes32 projectId) returns()
func (_Sale *SaleTransactorSession) SetImplicitModeProjectId(projectId [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetImplicitModeProjectId(&_Sale.TransactOpts, projectId)
}

// SetImplicitModeValidator is a paid mutator transaction binding the contract method 0x0bb310de.
//
// Solidity: function setImplicitModeValidator(address validator) returns()
func (_Sale *SaleTransactor) SetImplicitModeValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setImplicitModeValidator", validator)
}

// SetImplicitModeValidator is a paid mutator transaction binding the contract method 0x0bb310de.
//
// Solidity: function setImplicitModeValidator(address validator) returns()
func (_Sale *SaleSession) SetImplicitModeValidator(validator common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetImplicitModeValidator(&_Sale.TransactOpts, validator)
}

// SetImplicitModeValidator is a paid mutator transaction binding the contract method 0x0bb310de.
//
// Solidity: function setImplicitModeValidator(address validator) returns()
func (_Sale *SaleTransactorSession) SetImplicitModeValidator(validator common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetImplicitModeValidator(&_Sale.TransactOpts, validator)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentTokenAddr) returns()
func (_Sale *SaleTransactor) SetPaymentToken(opts *bind.TransactOpts, paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setPaymentToken", paymentTokenAddr)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentTokenAddr) returns()
func (_Sale *SaleSession) SetPaymentToken(paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetPaymentToken(&_Sale.TransactOpts, paymentTokenAddr)
}

// SetPaymentToken is a paid mutator transaction binding the contract method 0x6a326ab1.
//
// Solidity: function setPaymentToken(address paymentTokenAddr) returns()
func (_Sale *SaleTransactorSession) SetPaymentToken(paymentTokenAddr common.Address) (*types.Transaction, error) {
	return _Sale.Contract.SetPaymentToken(&_Sale.TransactOpts, paymentTokenAddr)
}

// SetTokenSaleDetails is a paid mutator transaction binding the contract method 0x4f651ccd.
//
// Solidity: function setTokenSaleDetails(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactor) SetTokenSaleDetails(opts *bind.TransactOpts, tokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setTokenSaleDetails", tokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetTokenSaleDetails is a paid mutator transaction binding the contract method 0x4f651ccd.
//
// Solidity: function setTokenSaleDetails(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleSession) SetTokenSaleDetails(tokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetTokenSaleDetails(&_Sale.TransactOpts, tokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetTokenSaleDetails is a paid mutator transaction binding the contract method 0x4f651ccd.
//
// Solidity: function setTokenSaleDetails(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot) returns()
func (_Sale *SaleTransactorSession) SetTokenSaleDetails(tokenId *big.Int, cost *big.Int, remainingSupply *big.Int, startTime uint64, endTime uint64, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetTokenSaleDetails(&_Sale.TransactOpts, tokenId, cost, remainingSupply, startTime, endTime, merkleRoot)
}

// SetTokenSaleDetailsBatch is a paid mutator transaction binding the contract method 0xf07f04ff.
//
// Solidity: function setTokenSaleDetailsBatch(uint256[] tokenIds, uint256[] costs, uint256[] remainingSupplies, uint64[] startTimes, uint64[] endTimes, bytes32[] merkleRoots) returns()
func (_Sale *SaleTransactor) SetTokenSaleDetailsBatch(opts *bind.TransactOpts, tokenIds []*big.Int, costs []*big.Int, remainingSupplies []*big.Int, startTimes []uint64, endTimes []uint64, merkleRoots [][32]byte) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "setTokenSaleDetailsBatch", tokenIds, costs, remainingSupplies, startTimes, endTimes, merkleRoots)
}

// SetTokenSaleDetailsBatch is a paid mutator transaction binding the contract method 0xf07f04ff.
//
// Solidity: function setTokenSaleDetailsBatch(uint256[] tokenIds, uint256[] costs, uint256[] remainingSupplies, uint64[] startTimes, uint64[] endTimes, bytes32[] merkleRoots) returns()
func (_Sale *SaleSession) SetTokenSaleDetailsBatch(tokenIds []*big.Int, costs []*big.Int, remainingSupplies []*big.Int, startTimes []uint64, endTimes []uint64, merkleRoots [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetTokenSaleDetailsBatch(&_Sale.TransactOpts, tokenIds, costs, remainingSupplies, startTimes, endTimes, merkleRoots)
}

// SetTokenSaleDetailsBatch is a paid mutator transaction binding the contract method 0xf07f04ff.
//
// Solidity: function setTokenSaleDetailsBatch(uint256[] tokenIds, uint256[] costs, uint256[] remainingSupplies, uint64[] startTimes, uint64[] endTimes, bytes32[] merkleRoots) returns()
func (_Sale *SaleTransactorSession) SetTokenSaleDetailsBatch(tokenIds []*big.Int, costs []*big.Int, remainingSupplies []*big.Int, startTimes []uint64, endTimes []uint64, merkleRoots [][32]byte) (*types.Transaction, error) {
	return _Sale.Contract.SetTokenSaleDetailsBatch(&_Sale.TransactOpts, tokenIds, costs, remainingSupplies, startTimes, endTimes, merkleRoots)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 value) returns()
func (_Sale *SaleTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "withdrawERC20", token, to, value)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 value) returns()
func (_Sale *SaleSession) WithdrawERC20(token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.WithdrawERC20(&_Sale.TransactOpts, token, to, value)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address token, address to, uint256 value) returns()
func (_Sale *SaleTransactorSession) WithdrawERC20(token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.WithdrawERC20(&_Sale.TransactOpts, token, to, value)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 value) returns()
func (_Sale *SaleTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.contract.Transact(opts, "withdrawETH", to, value)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 value) returns()
func (_Sale *SaleSession) WithdrawETH(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.WithdrawETH(&_Sale.TransactOpts, to, value)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 value) returns()
func (_Sale *SaleTransactorSession) WithdrawETH(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Sale.Contract.WithdrawETH(&_Sale.TransactOpts, to, value)
}

// SaleGlobalSaleDetailsUpdatedIterator is returned from FilterGlobalSaleDetailsUpdated and is used to iterate over the raw logs and unpacked data for GlobalSaleDetailsUpdated events raised by the Sale contract.
type SaleGlobalSaleDetailsUpdatedIterator struct {
	Event *SaleGlobalSaleDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *SaleGlobalSaleDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleGlobalSaleDetailsUpdated)
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
		it.Event = new(SaleGlobalSaleDetailsUpdated)
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
func (it *SaleGlobalSaleDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleGlobalSaleDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleGlobalSaleDetailsUpdated represents a GlobalSaleDetailsUpdated event raised by the Sale contract.
type SaleGlobalSaleDetailsUpdated struct {
	MinTokenId      *big.Int
	MaxTokenId      *big.Int
	Cost            *big.Int
	RemainingSupply *big.Int
	StartTime       uint64
	EndTime         uint64
	MerkleRoot      [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterGlobalSaleDetailsUpdated is a free log retrieval operation binding the contract event 0x099eb0c4bcb0e32c2243e81b424e1960f39d7ac2fe1598742f19aa8aec203acb.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) FilterGlobalSaleDetailsUpdated(opts *bind.FilterOpts) (*SaleGlobalSaleDetailsUpdatedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "GlobalSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SaleGlobalSaleDetailsUpdatedIterator{contract: _Sale.contract, event: "GlobalSaleDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalSaleDetailsUpdated is a free log subscription operation binding the contract event 0x099eb0c4bcb0e32c2243e81b424e1960f39d7ac2fe1598742f19aa8aec203acb.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) WatchGlobalSaleDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SaleGlobalSaleDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "GlobalSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleGlobalSaleDetailsUpdated)
				if err := _Sale.contract.UnpackLog(event, "GlobalSaleDetailsUpdated", log); err != nil {
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

// ParseGlobalSaleDetailsUpdated is a log parse operation binding the contract event 0x099eb0c4bcb0e32c2243e81b424e1960f39d7ac2fe1598742f19aa8aec203acb.
//
// Solidity: event GlobalSaleDetailsUpdated(uint256 minTokenId, uint256 maxTokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) ParseGlobalSaleDetailsUpdated(log types.Log) (*SaleGlobalSaleDetailsUpdated, error) {
	event := new(SaleGlobalSaleDetailsUpdated)
	if err := _Sale.contract.UnpackLog(event, "GlobalSaleDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleItemsMintedIterator is returned from FilterItemsMinted and is used to iterate over the raw logs and unpacked data for ItemsMinted events raised by the Sale contract.
type SaleItemsMintedIterator struct {
	Event *SaleItemsMinted // Event containing the contract specifics and raw log

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
func (it *SaleItemsMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleItemsMinted)
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
		it.Event = new(SaleItemsMinted)
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
func (it *SaleItemsMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleItemsMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleItemsMinted represents a ItemsMinted event raised by the Sale contract.
type SaleItemsMinted struct {
	To       common.Address
	TokenIds []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemsMinted is a free log retrieval operation binding the contract event 0x23519238b590c499a2abcf44a33b5f431ac6ca51c22fad26bc2c3d08b97eaa21.
//
// Solidity: event ItemsMinted(address to, uint256[] tokenIds, uint256[] amounts)
func (_Sale *SaleFilterer) FilterItemsMinted(opts *bind.FilterOpts) (*SaleItemsMintedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "ItemsMinted")
	if err != nil {
		return nil, err
	}
	return &SaleItemsMintedIterator{contract: _Sale.contract, event: "ItemsMinted", logs: logs, sub: sub}, nil
}

// WatchItemsMinted is a free log subscription operation binding the contract event 0x23519238b590c499a2abcf44a33b5f431ac6ca51c22fad26bc2c3d08b97eaa21.
//
// Solidity: event ItemsMinted(address to, uint256[] tokenIds, uint256[] amounts)
func (_Sale *SaleFilterer) WatchItemsMinted(opts *bind.WatchOpts, sink chan<- *SaleItemsMinted) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "ItemsMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleItemsMinted)
				if err := _Sale.contract.UnpackLog(event, "ItemsMinted", log); err != nil {
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

// ParseItemsMinted is a log parse operation binding the contract event 0x23519238b590c499a2abcf44a33b5f431ac6ca51c22fad26bc2c3d08b97eaa21.
//
// Solidity: event ItemsMinted(address to, uint256[] tokenIds, uint256[] amounts)
func (_Sale *SaleFilterer) ParseItemsMinted(log types.Log) (*SaleItemsMinted, error) {
	event := new(SaleItemsMinted)
	if err := _Sale.contract.UnpackLog(event, "ItemsMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Sale contract.
type SaleRoleAdminChangedIterator struct {
	Event *SaleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SaleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleRoleAdminChanged)
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
		it.Event = new(SaleRoleAdminChanged)
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
func (it *SaleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleRoleAdminChanged represents a RoleAdminChanged event raised by the Sale contract.
type SaleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sale *SaleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SaleRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SaleRoleAdminChangedIterator{contract: _Sale.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sale *SaleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SaleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleRoleAdminChanged)
				if err := _Sale.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Sale *SaleFilterer) ParseRoleAdminChanged(log types.Log) (*SaleRoleAdminChanged, error) {
	event := new(SaleRoleAdminChanged)
	if err := _Sale.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Sale contract.
type SaleRoleGrantedIterator struct {
	Event *SaleRoleGranted // Event containing the contract specifics and raw log

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
func (it *SaleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleRoleGranted)
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
		it.Event = new(SaleRoleGranted)
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
func (it *SaleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleRoleGranted represents a RoleGranted event raised by the Sale contract.
type SaleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SaleRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SaleRoleGrantedIterator{contract: _Sale.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SaleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleRoleGranted)
				if err := _Sale.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) ParseRoleGranted(log types.Log) (*SaleRoleGranted, error) {
	event := new(SaleRoleGranted)
	if err := _Sale.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Sale contract.
type SaleRoleRevokedIterator struct {
	Event *SaleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SaleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleRoleRevoked)
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
		it.Event = new(SaleRoleRevoked)
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
func (it *SaleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleRoleRevoked represents a RoleRevoked event raised by the Sale contract.
type SaleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SaleRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sale.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SaleRoleRevokedIterator{contract: _Sale.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SaleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Sale.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleRoleRevoked)
				if err := _Sale.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Sale *SaleFilterer) ParseRoleRevoked(log types.Log) (*SaleRoleRevoked, error) {
	event := new(SaleRoleRevoked)
	if err := _Sale.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleTokenSaleDetailsUpdatedIterator is returned from FilterTokenSaleDetailsUpdated and is used to iterate over the raw logs and unpacked data for TokenSaleDetailsUpdated events raised by the Sale contract.
type SaleTokenSaleDetailsUpdatedIterator struct {
	Event *SaleTokenSaleDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *SaleTokenSaleDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleTokenSaleDetailsUpdated)
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
		it.Event = new(SaleTokenSaleDetailsUpdated)
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
func (it *SaleTokenSaleDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleTokenSaleDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleTokenSaleDetailsUpdated represents a TokenSaleDetailsUpdated event raised by the Sale contract.
type SaleTokenSaleDetailsUpdated struct {
	TokenId         *big.Int
	Cost            *big.Int
	RemainingSupply *big.Int
	StartTime       uint64
	EndTime         uint64
	MerkleRoot      [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenSaleDetailsUpdated is a free log retrieval operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) FilterTokenSaleDetailsUpdated(opts *bind.FilterOpts) (*SaleTokenSaleDetailsUpdatedIterator, error) {

	logs, sub, err := _Sale.contract.FilterLogs(opts, "TokenSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SaleTokenSaleDetailsUpdatedIterator{contract: _Sale.contract, event: "TokenSaleDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenSaleDetailsUpdated is a free log subscription operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) WatchTokenSaleDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SaleTokenSaleDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Sale.contract.WatchLogs(opts, "TokenSaleDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleTokenSaleDetailsUpdated)
				if err := _Sale.contract.UnpackLog(event, "TokenSaleDetailsUpdated", log); err != nil {
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

// ParseTokenSaleDetailsUpdated is a log parse operation binding the contract event 0x8ced76aee4b96a1e218e7903610fc7d648023d9075677163a5b31396cb280f96.
//
// Solidity: event TokenSaleDetailsUpdated(uint256 tokenId, uint256 cost, uint256 remainingSupply, uint64 startTime, uint64 endTime, bytes32 merkleRoot)
func (_Sale *SaleFilterer) ParseTokenSaleDetailsUpdated(log types.Log) (*SaleTokenSaleDetailsUpdated, error) {
	event := new(SaleTokenSaleDetailsUpdated)
	if err := _Sale.contract.UnpackLog(event, "TokenSaleDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
