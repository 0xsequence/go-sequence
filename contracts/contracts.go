package contracts

import (
	_ "embed"

	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts/gen/walletfactory"
	"github.com/0xsequence/go-sequence/contracts/gen/walletguest"
	"github.com/0xsequence/go-sequence/contracts/gen/walletmain"
	"github.com/0xsequence/go-sequence/contracts/gen/walletupgradable"
	"github.com/0xsequence/go-sequence/contracts/gen/walletutils"
)

var (
	WalletFactory,
	WalletMainModule,
	WalletMainModuleUpgradable,
	WalletGuestModule,
	WalletUtils,
	IERC1271,
	_ ethartifact.Artifact
)

var (
	//go:embed artifacts/erc1271/ierc1271.json
	artifact_ierc1271 string
)

func init() {
	WalletFactory = artifact("WALLET_FACTORY", walletfactory.WalletFactoryABI, walletfactory.WalletFactoryBin)
	WalletMainModule = artifact("WALLET_MAIN", walletmain.WalletMainABI, walletmain.WalletMainBin)
	WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", walletupgradable.WalletUpgradableABI, walletupgradable.WalletUpgradableBin)
	WalletGuestModule = artifact("WALLET_GUEST", walletguest.WalletGuestABI, walletguest.WalletGuestBin)
	WalletUtils = artifact("WALLET_UTILS", walletutils.WalletUtilsABI, walletutils.WalletUtilsBin)

	IERC1271 = MustParseArtifactJSON(artifact_ierc1271)
}

func artifact(contractName, abiJSON, bytecodeHex string) ethartifact.Artifact {
	return ethartifact.Artifact{
		ContractName: contractName,
		ABI:          ethcontract.MustParseABI(abiJSON),
		Bin:          common.FromHex(bytecodeHex),
	}
}

// TODO: use ethkit function..
func MustParseArtifactJSON(artifactJSON string) ethartifact.Artifact {
	artifact, err := ethartifact.ParseArtifactJSON(artifactJSON)
	if err != nil {
		panic(err)
	}
	return artifact
}
