package contract

import (
	_ "embed"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contract/gen/walletfactory"
	"github.com/0xsequence/go-sequence/contract/gen/walletguest"
	"github.com/0xsequence/go-sequence/contract/gen/walletmain"
	"github.com/0xsequence/go-sequence/contract/gen/walletupgradable"
	"github.com/0xsequence/go-sequence/contract/gen/walletutils"
)

var (
	WalletFactory,
	WalletMainModule,
	WalletMainModuleUpgradable,
	WalletGuestModule,
	WalletUtils,
	IERC1271,
	_ *ContractABI
)

var (
	//go:embed artifacts/erc1271/ierc1271.json
	artifact_ierc1271 string
)

func init() {
	WalletFactory = &ContractABI{Name: "WALLET_FACTORY", ABI: MustParseABI(walletfactory.WalletFactoryABI), Bin: common.FromHex(walletfactory.WalletFactoryBin)}
	WalletMainModule = &ContractABI{Name: "WALLET_MAIN", ABI: MustParseABI(walletmain.WalletMainABI), Bin: common.FromHex(walletmain.WalletMainBin)}
	WalletMainModuleUpgradable = &ContractABI{Name: "WALLET_UPGRADABLE", ABI: MustParseABI(walletupgradable.WalletUpgradableABI), Bin: common.FromHex(walletupgradable.WalletUpgradableBin)}
	WalletGuestModule = &ContractABI{Name: "WALLET_GUEST", ABI: MustParseABI(walletguest.WalletGuestABI), Bin: common.FromHex(walletguest.WalletGuestBin)}
	WalletUtils = &ContractABI{Name: "WALLET_UTILS", ABI: MustParseABI(walletutils.WalletUtilsABI), Bin: common.FromHex(walletutils.WalletUtilsBin)}

	IERC1271 = MustParseArtifactJSON(artifact_ierc1271)
}
