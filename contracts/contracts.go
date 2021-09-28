package contracts

import (
	_ "embed"

	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts/gen/gasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/ierc1271"
	"github.com/0xsequence/go-sequence/contracts/gen/niftyswap"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
	"github.com/0xsequence/go-sequence/contracts/gen/walletfactory"
	"github.com/0xsequence/go-sequence/contracts/gen/walletgasestimator"
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
	WalletRequireFreshSigner,
	WalletGasEstimator,
	GasEstimator,
	IERC1271,
	ERC20Mock,
	IERC20,
	IERC721,
	IERC1155,
	IERC20Wrapper,
	NiftyswapExchange,
	NiftyswapFactory,
	WrapAndNiftyswap,
	_ ethartifact.Artifact
)

var (
	//go:embed artifacts/erc1155/mocks/ERC20Mock.sol/ERC20Mock.json
	artifact_erc20mock string
)

func init() {
	WalletFactory = artifact("WALLET_FACTORY", walletfactory.WalletFactoryABI, walletfactory.WalletFactoryBin)
	WalletMainModule = artifact("WALLET_MAIN", walletmain.WalletMainABI, walletmain.WalletMainBin)
	WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", walletupgradable.WalletUpgradableABI, walletupgradable.WalletUpgradableBin)
	WalletGuestModule = artifact("WALLET_GUEST", walletguest.WalletGuestABI, walletguest.WalletGuestBin)
	WalletUtils = artifact("WALLET_UTILS", walletutils.WalletUtilsABI, walletutils.WalletUtilsBin)
	WalletRequireFreshSigner = artifact("WALLET_REQUIRE_FRESH_SIGNER", walletutils.WalletRequireFreshSignerABI, walletutils.WalletRequireFreshSignerBin)
	WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", walletgasestimator.WalletGasEstimatorABI, walletgasestimator.WalletGasEstimatorBin, walletgasestimator.WalletGasEstimatorDeployedBin)
	GasEstimator = artifact("GAS_ESTIMATOR", gasestimator.GasEstimatorABI, gasestimator.GasEstimatorBin, gasestimator.GasEstimatorDeployedBin)

	IERC1271 = artifact("IERC1271", ierc1271.IERC1271ABI, "")

	IERC20 = artifact("IERC20", tokens.IERC20ABI, "")
	IERC721 = artifact("IERC721", tokens.IERC721ABI, "")
	IERC1155 = artifact("IERC1155", tokens.IERC1155ABI, "")
	IERC20Wrapper = artifact("IERC20Wrapper", tokens.IERC20WrapperABI, "")

	NiftyswapExchange = artifact("NIFTYSWAP_EXCHANGE", niftyswap.NiftyswapExchangeABI, niftyswap.NiftyswapFactoryBin)
	NiftyswapFactory = artifact("NIFTYSWAP_FACTORY", niftyswap.NiftyswapFactoryABI, niftyswap.NiftyswapFactoryBin)
	WrapAndNiftyswap = artifact("WRAP_AND_NIFTYSWAP", niftyswap.WrapAndNiftyswapABI, niftyswap.WrapAndNiftyswapBin)

	ERC20Mock = ethartifact.MustParseArtifactJSON(artifact_erc20mock)
}

func artifact(contractName, abiJSON, bytecodeHex string, deployedBytecodeHex ...string) ethartifact.Artifact {
	var deployedBin []byte
	if len(deployedBytecodeHex) > 0 {
		deployedBin = common.FromHex(deployedBytecodeHex[0])
	}
	return ethartifact.Artifact{
		ContractName: contractName,
		ABI:          ethcontract.MustParseABI(abiJSON),
		Bin:          common.FromHex(bytecodeHex),
		DeployedBin:  deployedBin,
	}
}
