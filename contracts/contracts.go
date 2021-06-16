//
// sequence wallet-contracts
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=Factory --outFile=wallet_factory.gen.go --artifactsFile=./wallet-contracts/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=MainModule --outFile=wallet_main_module.gen.go --artifactsFile=./wallet-contracts/modules/MainModule.sol/MainModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=MainModuleUpgradable --outFile=wallet_main_module_upgradable.gen.go --artifactsFile=./wallet-contracts/modules/MainModuleUpgradable.sol/MainModuleUpgradable.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=GuestModule --outFile=wallet_guest_module.gen.go --artifactsFile=./wallet-contracts/modules/GuestModule.sol/GuestModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=SequenceUtils --outFile=wallet_utils.gen.go --artifactsFile=./wallet-contracts/modules/utils/SequenceUtils.sol/SequenceUtils.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=CallReceiverMock --outFile=wallet_call_mock.gen.go --artifactsFile=./wallet-contracts/mocks/CallReceiverMock.sol/CallReceiverMock.json

//
// erc1271
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=contracts --type=IERC1271 --outFile=ierc1271.gen.go --artifactsFile=./erc1271/ierc1271.json

package contracts

import (
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
)

var (
	ABI_WALLET_FACTORY,
	ABI_WALLET_MAIN_MODULE,
	ABI_WALLET_MAIN_MODULE_UPGRADABLE,
	ABI_WALLET_GUEST_MODULE,
	ABI_WALLET_UTILS,
	_ *abi.ABI
)

var (
	ABI_ERC1271_isValidSignatureBytes32 = "0x1626ba7e"
)

func init() {
	ABI_WALLET_FACTORY = MustParseABI(FactoryABI)
	ABI_WALLET_MAIN_MODULE = MustParseABI(MainModuleABI)
	ABI_WALLET_MAIN_MODULE_UPGRADABLE = MustParseABI(MainModuleUpgradableABI)
	ABI_WALLET_GUEST_MODULE = MustParseABI(GuestModuleABI)
	ABI_WALLET_UTILS = MustParseABI(SequenceUtilsABI)
}

type Artifact struct {
	ABI *abi.ABI
	Bin []byte
}

type ArtifactMap map[string]*Artifact

func MustParseABI(abiJSON string) *abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic("contracts: unable to parse abi json")
	}
	return &parsed
}

// TODO: ....... is there any convention which we can do like, contract_wallet, contract_erc1155, contract_erc20, contract_erc721, contract_erc1271, ..

// TODO: convienient to include erc20, erc721, erc1155 in here too..

// or name packages by their erc20

// could also have abigen tool with flag --concise where we just generate the ABI and the Bin.. leaving out all other commands..

// walletcontractfactory.FactoryABI ..
// walletcontractmain
// walletcontractupgradeable
// walletcontractXX .. ugly. lol..

// TODO: include a bunch of others.. openzeppelin, erc-1155, ......
