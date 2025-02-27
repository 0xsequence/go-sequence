package contracts

import (
	_ "embed"

	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence/contracts/gen/gasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/ierc1271"
	"github.com/0xsequence/go-sequence/contracts/gen/niftyswap"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
	walletfactory1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletfactory"
	walletgasestimator1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletgasestimator"
	walletguest1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletguest"
	walletmain1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletmain"
	walletupgradable1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletupgradable"
	walletutils1 "github.com/0xsequence/go-sequence/contracts/gen/v1/walletutils"
	walletfactory2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletfactory"
	walletgasestimator2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletgasestimator"
	walletguest2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletguest"
	walletmain2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletmain"
	walletupgradable2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletupgradable"
	walletutils2 "github.com/0xsequence/go-sequence/contracts/gen/v2/walletutils"
	walletfactory3 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletfactory"
	walletguest3 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletguest"
	walletstage13 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletstage1"
	walletstage23 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletstage2"
)

var (
	GasEstimator,
	IERC1271,
	ERC20Mock,
	IERC20,
	IERC721,
	IERC1155,
	IERC20Wrapper,
	INiftyswapExchange,
	INiftyswapExchange20,
	NiftyswapExchange,
	NiftyswapExchange20,
	NiftyswapFactory,
	WrapAndNiftyswap,
	_ ethartifact.Artifact
)

var V1 struct {
	WalletFactory              ethartifact.Artifact
	WalletMainModule           ethartifact.Artifact
	WalletMainModuleUpgradable ethartifact.Artifact
	WalletGuestModule          ethartifact.Artifact
	WalletUtils                ethartifact.Artifact
	WalletRequireFreshSigner   ethartifact.Artifact
	WalletGasEstimator         ethartifact.Artifact
	CreationCode               []byte
}

var V2 struct {
	WalletFactory              ethartifact.Artifact
	WalletMainModule           ethartifact.Artifact
	WalletMainModuleUpgradable ethartifact.Artifact
	WalletGuestModule          ethartifact.Artifact
	WalletUtils                ethartifact.Artifact
	WalletGasEstimator         ethartifact.Artifact
	CreationCode               []byte
}

var V3 struct {
	WalletFactory      ethartifact.Artifact
	WalletStage1Module ethartifact.Artifact
	WalletStage2Module ethartifact.Artifact
	WalletGuestModule  ethartifact.Artifact
	WalletGasEstimator ethartifact.Artifact
	CreationCode       []byte
}

var (
	//go:embed artifacts/erc1155/mocks/ERC20Mock.sol/ERC20Mock.json
	artifact_erc20mock string
)

func init() {
	V1.WalletFactory = artifact("WALLET_FACTORY", walletfactory1.WalletFactoryABI, walletfactory1.WalletFactoryBin)
	V1.WalletMainModule = artifact("WALLET_MAIN", walletmain1.WalletMainABI, walletmain1.WalletMainBin)
	V1.WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", walletupgradable1.WalletUpgradableABI, walletupgradable1.WalletUpgradableBin)
	V1.WalletGuestModule = artifact("WALLET_GUEST", walletguest1.WalletGuestABI, walletguest1.WalletGuestBin)
	V1.WalletUtils = artifact("WALLET_UTILS", walletutils1.WalletUtilsABI, walletutils1.WalletUtilsBin)
	V1.WalletRequireFreshSigner = artifact("WALLET_REQUIRE_FRESH_SIGNER", walletutils1.WalletRequireFreshSignerABI, walletutils1.WalletRequireFreshSignerBin)
	V1.WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", walletgasestimator1.WalletGasEstimatorABI, walletgasestimator1.WalletGasEstimatorBin, walletgasestimator1.WalletGasEstimatorDeployedBin)
	V1.CreationCode = hexutil.MustDecode("0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3")

	V2.WalletFactory = artifact("WALLET_FACTORY", walletfactory2.WalletFactoryABI, walletfactory2.WalletFactoryBin)
	V2.WalletMainModule = artifact("WALLET_MAIN", walletmain2.WalletMainABI, walletmain2.WalletMainBin)
	V2.WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", walletupgradable2.WalletUpgradableABI, walletupgradable2.WalletUpgradableBin)
	V2.WalletGuestModule = artifact("WALLET_GUEST", walletguest2.WalletGuestABI, walletguest2.WalletGuestBin)
	V2.WalletUtils = artifact("WALLET_UTILS", walletutils2.WalletUtilsABI, walletutils2.WalletUtilsBin)
	V2.WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", walletgasestimator2.WalletGasEstimatorABI, walletgasestimator2.WalletGasEstimatorBin, walletgasestimator2.WalletGasEstimatorDeployedBin)
	V2.CreationCode = hexutil.MustDecode("0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3")

	V3.WalletFactory = artifact("WALLET_FACTORY", walletfactory3.WalletFactoryABI, walletfactory3.WalletFactoryBin)
	V3.WalletStage1Module = artifact("WALLET_STAGE_1", walletstage13.WalletStage1ABI, walletstage13.WalletStage1Bin)
	V3.WalletStage2Module = artifact("WALLET_STAGE_2", walletstage23.WalletStage2ABI, walletstage23.WalletStage2Bin)
	V3.WalletGuestModule = artifact("WALLET_GUEST", walletguest3.WalletGuestABI, walletguest3.WalletGuestBin)
	V3.CreationCode = hexutil.MustDecode("0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3")

	// TODO: Update w/ v3 gas estimator
	V3.WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", walletgasestimator2.WalletGasEstimatorABI, walletgasestimator2.WalletGasEstimatorBin, walletgasestimator2.WalletGasEstimatorDeployedBin)

	GasEstimator = artifact("GAS_ESTIMATOR", gasestimator.GasEstimatorABI, gasestimator.GasEstimatorBin, gasestimator.GasEstimatorDeployedBin)

	IERC1271 = artifact("IERC1271", ierc1271.IERC1271ABI, "")

	IERC20 = artifact("IERC20", tokens.IERC20ABI, "")
	IERC721 = artifact("IERC721", tokens.IERC721ABI, "")
	IERC1155 = artifact("IERC1155", tokens.IERC1155ABI, "")
	IERC20Wrapper = artifact("IERC20Wrapper", tokens.IERC20WrapperABI, "")

	INiftyswapExchange = artifact("INIFTYSWAP_EXCHANGE", niftyswap.INiftyswapExchangeABI, "")
	INiftyswapExchange20 = artifact("INIFTYSWAP_EXCHANGE_20", niftyswap.INiftyswapExchange20ABI, "")
	NiftyswapExchange = artifact("NIFTYSWAP_EXCHANGE", niftyswap.NiftyswapExchangeABI, niftyswap.NiftyswapExchangeBin)
	NiftyswapExchange20 = artifact("NIFTYSWAP_EXCHANGE_20", niftyswap.NiftyswapExchange20ABI, niftyswap.NiftyswapExchange20Bin)
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
