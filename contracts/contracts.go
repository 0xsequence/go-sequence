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
	seqmarketplace "github.com/0xsequence/go-sequence/contracts/gen/seq_marketplace"
	seqsale1155v0 "github.com/0xsequence/go-sequence/contracts/gen/seq_sale/erc1155v0"
	seqsale1155v1 "github.com/0xsequence/go-sequence/contracts/gen/seq_sale/erc1155v1"
	seqsale721v0 "github.com/0xsequence/go-sequence/contracts/gen/seq_sale/erc721v0"
	"github.com/0xsequence/go-sequence/contracts/gen/supply"
	"github.com/0xsequence/go-sequence/contracts/gen/tokens"
	v1Factory "github.com/0xsequence/go-sequence/contracts/gen/v1/walletfactory"
	v1Estimator "github.com/0xsequence/go-sequence/contracts/gen/v1/walletgasestimator"
	v1Guest "github.com/0xsequence/go-sequence/contracts/gen/v1/walletguest"
	v1Main "github.com/0xsequence/go-sequence/contracts/gen/v1/walletmain"
	v1Upgradable "github.com/0xsequence/go-sequence/contracts/gen/v1/walletupgradable"
	v1Utils "github.com/0xsequence/go-sequence/contracts/gen/v1/walletutils"
	v2Factory "github.com/0xsequence/go-sequence/contracts/gen/v2/walletfactory"
	v2Estimator "github.com/0xsequence/go-sequence/contracts/gen/v2/walletgasestimator"
	v2Guest "github.com/0xsequence/go-sequence/contracts/gen/v2/walletguest"
	v2Main "github.com/0xsequence/go-sequence/contracts/gen/v2/walletmain"
	v2Upgradable "github.com/0xsequence/go-sequence/contracts/gen/v2/walletupgradable"
	v2Utils "github.com/0xsequence/go-sequence/contracts/gen/v2/walletutils"
	v3Estimator "github.com/0xsequence/go-sequence/contracts/gen/v3/walletestimator"
	v3Factory "github.com/0xsequence/go-sequence/contracts/gen/v3/walletfactory"
	v3Guest "github.com/0xsequence/go-sequence/contracts/gen/v3/walletguest"
	v3Simulator "github.com/0xsequence/go-sequence/contracts/gen/v3/walletsimulator"
	v3Stage1 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletstage1"
	v3Stage2 "github.com/0xsequence/go-sequence/contracts/gen/v3/walletstage2"
	"github.com/0xsequence/go-sequence/contracts/gen/valueforwarder"
)

var GasEstimator,
	IERC1271,
	ERC20Mock,
	IERC20,
	IERC721,
	IERC1155,
	IERC1155Supply,
	IERC20Wrapper,
	INiftyswapExchange,
	INiftyswapExchange20,
	NiftyswapExchange,
	NiftyswapExchange20,
	NiftyswapFactory,
	WrapAndNiftyswap,
	SeqMarketplace,
	SeqSale721V0,
	SeqSale1155V0,
	SeqSale1155V1,
	ValueForwarder,
	_ ethartifact.Artifact

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
	WalletEstimator    ethartifact.Artifact
	WalletSimulator    ethartifact.Artifact
	CreationCode       []byte
}

var (
	//go:embed artifacts/erc1155/mocks/ERC20Mock.sol/ERC20Mock.json
	artifact_erc20mock string
)

func init() {
	V1.WalletFactory = artifact("WALLET_FACTORY", v1Factory.WalletFactoryABI, v1Factory.WalletFactoryBin)
	V1.WalletMainModule = artifact("WALLET_MAIN", v1Main.WalletMainABI, v1Main.WalletMainBin)
	V1.WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", v1Upgradable.WalletUpgradableABI, v1Upgradable.WalletUpgradableBin)
	V1.WalletGuestModule = artifact("WALLET_GUEST", v1Guest.WalletGuestABI, v1Guest.WalletGuestBin)
	V1.WalletUtils = artifact("WALLET_UTILS", v1Utils.WalletUtilsABI, v1Utils.WalletUtilsBin)
	V1.WalletRequireFreshSigner = artifact("WALLET_REQUIRE_FRESH_SIGNER", v1Utils.WalletRequireFreshSignerABI, v1Utils.WalletRequireFreshSignerBin)
	V1.WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", v1Estimator.WalletGasEstimatorABI, v1Estimator.WalletGasEstimatorBin, v1Estimator.WalletGasEstimatorDeployedBin)
	V1.CreationCode = hexutil.MustDecode("0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3")

	V2.WalletFactory = artifact("WALLET_FACTORY", v2Factory.WalletFactoryABI, v2Factory.WalletFactoryBin)
	V2.WalletMainModule = artifact("WALLET_MAIN", v2Main.WalletMainABI, v2Main.WalletMainBin)
	V2.WalletMainModuleUpgradable = artifact("WALLET_UPGRADABLE", v2Upgradable.WalletUpgradableABI, v2Upgradable.WalletUpgradableBin)
	V2.WalletGuestModule = artifact("WALLET_GUEST", v2Guest.WalletGuestABI, v2Guest.WalletGuestBin)
	V2.WalletUtils = artifact("WALLET_UTILS", v2Utils.WalletUtilsABI, v2Utils.WalletUtilsBin)
	V2.WalletGasEstimator = artifact("WALLET_GAS_ESTIMATOR", v2Estimator.WalletGasEstimatorABI, v2Estimator.WalletGasEstimatorBin, v2Estimator.WalletGasEstimatorDeployedBin)
	V2.CreationCode = hexutil.MustDecode("0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3")

	V3.WalletFactory = artifact("WALLET_FACTORY", v3Factory.WalletFactoryABI, v3Factory.WalletFactoryBin)
	V3.WalletStage1Module = artifact("WALLET_STAGE_1", v3Stage1.WalletStage1ABI, v3Stage1.WalletStage1Bin)
	V3.WalletStage2Module = artifact("WALLET_STAGE_2", v3Stage2.WalletStage2ABI, v3Stage2.WalletStage2Bin)
	V3.WalletGuestModule = artifact("WALLET_GUEST", v3Guest.WalletGuestABI, v3Guest.WalletGuestBin)
	V3.WalletEstimator = artifact("WALLET_ESTIMATOR", v3Estimator.WalletEstimatorABI, v3Estimator.WalletEstimatorBin, v3Estimator.WalletEstimatorDeployedBin)
	V3.WalletSimulator = artifact("WALLET_SIMULATOR", v3Simulator.WalletSimulatorABI, v3Simulator.WalletSimulatorBin, v3Simulator.WalletSimulatorDeployedBin)
	V3.CreationCode = hexutil.MustDecode("0x6041600e3d396021805130553df33d3d36153402601f57363d3d373d363d30545af43d82803e903d91601f57fd5bf3")

	GasEstimator = artifact("GAS_ESTIMATOR", gasestimator.GasEstimatorABI, gasestimator.GasEstimatorBin, gasestimator.GasEstimatorDeployedBin)

	ValueForwarder = artifact("VALUE_FORWARDER", valueforwarder.ValueForwarderABI, valueforwarder.ValueForwarderBin)

	IERC1271 = artifact("IERC1271", ierc1271.IERC1271ABI, "")

	IERC20 = artifact("IERC20", tokens.IERC20ABI, "")
	IERC721 = artifact("IERC721", tokens.IERC721ABI, "")
	IERC1155 = artifact("IERC1155", tokens.IERC1155ABI, "")
	IERC20Wrapper = artifact("IERC20Wrapper", tokens.IERC20WrapperABI, "")

	IERC1155Supply = artifact("IERC1155Supply", supply.IERC1155SupplyABI, "")

	INiftyswapExchange = artifact("INIFTYSWAP_EXCHANGE", niftyswap.INiftyswapExchangeABI, "")
	INiftyswapExchange20 = artifact("INIFTYSWAP_EXCHANGE_20", niftyswap.INiftyswapExchange20ABI, "")
	NiftyswapExchange = artifact("NIFTYSWAP_EXCHANGE", niftyswap.NiftyswapExchangeABI, niftyswap.NiftyswapExchangeBin)
	NiftyswapExchange20 = artifact("NIFTYSWAP_EXCHANGE_20", niftyswap.NiftyswapExchange20ABI, niftyswap.NiftyswapExchange20Bin)
	NiftyswapFactory = artifact("NIFTYSWAP_FACTORY", niftyswap.NiftyswapFactoryABI, niftyswap.NiftyswapFactoryBin)
	WrapAndNiftyswap = artifact("WRAP_AND_NIFTYSWAP", niftyswap.WrapAndNiftyswapABI, niftyswap.WrapAndNiftyswapBin)

	SeqMarketplace = artifact("SEQ_MARKETPLACE", seqmarketplace.SequenceMarketplaceABI, seqmarketplace.SequenceMarketplaceMetaData.Bin)
	SeqSale721V0 = artifact("SEQ_SALE_ERC721_v0", seqsale721v0.SaleABI, seqsale721v0.SaleMetaData.Bin)
	SeqSale1155V0 = artifact("SEQ_SALE_ERC1155_V0", seqsale1155v0.SaleABI, seqsale1155v0.SaleMetaData.Bin)
	SeqSale1155V1 = artifact("SEQ_SALE_ERC1155_V1", seqsale1155v1.SaleABI, seqsale1155v1.SaleMetaData.Bin)

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
