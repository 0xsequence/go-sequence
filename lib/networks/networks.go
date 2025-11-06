package networks



var All = []Network{
	{
		ChainID: ChainID_MAINNET,
		Type:    NetworkTypeMainnet,
		Name:    "mainnet",
		Title:   "Ethereum",
		NodeURL: getRpcURL("mainnet"),
		LogoURL: getLogoURL(ChainID_MAINNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan",
			RootURL: "https://etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	},
	{
		ChainID: ChainID_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "sepolia",
		Title:   "Sepolia",
		NodeURL: getRpcURL("sepolia"),
		LogoURL: getLogoURL(ChainID_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Sepolia)",
			RootURL: "https://sepolia.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_POLYGON,
		Type:    NetworkTypeMainnet,
		Name:    "polygon",
		Title:   "Polygon",
		NodeURL: getRpcURL("polygon"),
		LogoURL: getLogoURL(ChainID_POLYGON),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Polygonscan",
			RootURL: "https://polygonscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "POL",
			Name:     "POL",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_POLYGON_AMOY,
		Type:    NetworkTypeTestnet,
		Name:    "amoy",
		Title:   "Polygon Amoy",
		NodeURL: getRpcURL("amoy"),
		LogoURL: getLogoURL(ChainID_POLYGON_AMOY),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "OKLink (Amoy)",
			RootURL: "https://www.oklink.com/amoy/",
		},
		NativeToken: &Currency{
			Symbol:   "aPOL",
			Name:     "Amoy POL",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_POLYGON_ZKEVM,
		Type:    NetworkTypeMainnet,
		Name:    "polygon-zkevm",
		Title:   "Polygon zkEVM",
		NodeURL: getRpcURL("polygon-zkevm"),
		LogoURL: getLogoURL(ChainID_POLYGON_ZKEVM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Polygonscan (zkEVM)",
			RootURL: "https://zkevm.polygonscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BSC,
		Type:    NetworkTypeMainnet,
		Name:    "bsc",
		Title:   "BNB Smart Chain",
		NodeURL: getRpcURL("bsc"),
		LogoURL: getLogoURL(ChainID_BSC),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "BSCScan",
			RootURL: "https://bscscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "BNB",
			Name:     "BNB",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BSC_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "bsc-testnet",
		Title:   "BNB Smart Chain Testnet",
		NodeURL: getRpcURL("bsc-testnet"),
		LogoURL: getLogoURL(ChainID_BSC_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "BSCScan (Testnet)",
			RootURL: "https://testnet.bscscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "tBNB",
			Name:     "Testnet BNB",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_OPTIMISM,
		Type:    NetworkTypeMainnet,
		Name:    "optimism",
		Title:   "Optimism",
		NodeURL: getRpcURL("optimism"),
		LogoURL: getLogoURL(ChainID_OPTIMISM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism)",
			RootURL: "https://optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_OPTIMISM_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "optimism-sepolia",
		Title:   "Optimism Sepolia",
		NodeURL: getRpcURL("optimism-sepolia"),
		LogoURL: getLogoURL(ChainID_OPTIMISM_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism Sepolia)",
			RootURL: "https://sepolia-optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ARBITRUM,
		Type:    NetworkTypeMainnet,
		Name:    "arbitrum",
		Title:   "Arbitrum One",
		NodeURL: getRpcURL("arbitrum"),
		LogoURL: getLogoURL(ChainID_ARBITRUM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan",
			RootURL: "https://arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ARBITRUM_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "arbitrum-sepolia",
		Title:   "Arbitrum Sepolia",
		NodeURL: getRpcURL("arbitrum-sepolia"),
		LogoURL: getLogoURL(ChainID_ARBITRUM_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan (Sepolia Testnet)",
			RootURL: "https://sepolia.arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ARBITRUM_NOVA,
		Type:    NetworkTypeMainnet,
		Name:    "arbitrum-nova",
		Title:   "Arbitrum Nova",
		NodeURL: getRpcURL("arbitrum-nova"),
		LogoURL: getLogoURL(ChainID_ARBITRUM_NOVA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan Nova",
			RootURL: "https://nova.arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_AVALANCHE,
		Type:    NetworkTypeMainnet,
		Name:    "avalanche",
		Title:   "Avalanche",
		NodeURL: getRpcURL("avalanche"),
		LogoURL: getLogoURL(ChainID_AVALANCHE),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Snowtrace",
			RootURL: "https://subnets.avax.network/c-chain/",
		},
		NativeToken: &Currency{
			Symbol:   "AVAX",
			Name:     "AVAX",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_AVALANCHE_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "avalanche-testnet",
		Title:   "Avalanche Testnet",
		NodeURL: getRpcURL("avalanche-testnet"),
		LogoURL: getLogoURL(ChainID_AVALANCHE_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Snowtrace (Testnet)",
			RootURL: "https://subnets-test.avax.network/c-chain/",
		},
		NativeToken: &Currency{
			Symbol:   "tAVAX",
			Name:     "Testnet AVAX",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_GNOSIS,
		Type:    NetworkTypeMainnet,
		Name:    "gnosis",
		Title:   "Gnosis Chain",
		NodeURL: getRpcURL("gnosis"),
		LogoURL: getLogoURL(ChainID_GNOSIS),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Gnosis Chain Explorer",
			RootURL: "https://blockscout.com/xdai/mainnet/",
		},
		NativeToken: &Currency{
			Symbol:   "XDAI",
			Name:     "XDAI",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BASE,
		Type:    NetworkTypeMainnet,
		Name:    "base",
		Title:   "Base (Coinbase)",
		NodeURL: getRpcURL("base"),
		LogoURL: getLogoURL(ChainID_BASE),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Base Explorer",
			RootURL: "https://basescan.org/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BASE_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "base-sepolia",
		Title:   "Base Sepolia",
		NodeURL: getRpcURL("base-sepolia"),
		LogoURL: getLogoURL(ChainID_BASE_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Base Sepolia Explorer",
			RootURL: "https://base-sepolia.blockscout.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_HOMEVERSE,
		Type:    NetworkTypeMainnet,
		Name:    "homeverse",
		Title:   "Oasys Homeverse",
		NodeURL: getRpcURL("homeverse"),
		LogoURL: getLogoURL(ChainID_HOMEVERSE),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Oasys Homeverse Explorer",
			RootURL: "https://explorer.oasys.homeverse.games/",
		},
		NativeToken: &Currency{
			Symbol:   "OAS",
			Name:     "OAS",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_HOMEVERSE_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "homeverse-testnet",
		Title:   "Oasys Homeverse Testnet",
		NodeURL: getRpcURL("homeverse-testnet"),
		LogoURL: getLogoURL(ChainID_HOMEVERSE_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Oasys Homeverse Explorer (Testnet)",
			RootURL: "https://explorer.testnet.oasys.homeverse.games/",
		},
		NativeToken: &Currency{
			Symbol:   "tOAS",
			Name:     "Testnet OAS",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_XAI,
		Type:    NetworkTypeMainnet,
		Name:    "xai",
		Title:   "Xai",
		NodeURL: getRpcURL("xai"),
		LogoURL: getLogoURL(ChainID_XAI),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Xai Explorer",
			RootURL: "https://explorer.xai-chain.net/",
		},
		NativeToken: &Currency{
			Symbol:   "XAI",
			Name:     "XAI",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_XAI_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "xai-sepolia",
		Title:   "Xai Sepolia",
		NodeURL: getRpcURL("xai-sepolia"),
		LogoURL: getLogoURL(ChainID_XAI_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Xai Sepolia Explorer",
			RootURL: "https://testnet-explorer-v2.xai-chain.net/",
		},
		NativeToken: &Currency{
			Symbol:   "sXAI",
			Name:     "Sepolia XAI",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_B3,
		Type:    NetworkTypeMainnet,
		Name:    "b3",
		Title:   "B3",
		NodeURL: getRpcURL("b3"),
		LogoURL: getLogoURL(ChainID_B3),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "B3 Explorer",
			RootURL: "https://explorer.b3.fun/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_B3_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "b3-sepolia",
		Title:   "B3 Sepolia",
		NodeURL: getRpcURL("b3-sepolia"),
		LogoURL: getLogoURL(ChainID_B3_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "B3 Sepolia Explorer",
			RootURL: "https://sepolia.explorer.b3.fun/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_APECHAIN,
		Type:    NetworkTypeMainnet,
		Name:    "apechain",
		Title:   "APE Chain",
		NodeURL: getRpcURL("apechain"),
		LogoURL: getLogoURL(ChainID_APECHAIN),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "APE Chain Explorer",
			RootURL: "https://apechain.calderaexplorer.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "APE",
			Name:     "ApeCoin",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_APECHAIN_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "apechain-testnet",
		Title:   "APE Chain Testnet",
		NodeURL: getRpcURL("apechain-testnet"),
		LogoURL: getLogoURL(ChainID_APECHAIN_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "APE Chain Explorer",
			RootURL: "https://curtis.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "APE",
			Name:     "ApeCoin",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BLAST,
		Type:    NetworkTypeMainnet,
		Name:    "blast",
		Title:   "Blast",
		NodeURL: getRpcURL("blast"),
		LogoURL: getLogoURL(ChainID_BLAST),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Blast Explorer",
			RootURL: "https://blastscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_BLAST_SEPOLIA,
		Type:    NetworkTypeTestnet,
		Name:    "blast-sepolia",
		Title:   "Blast Sepolia",
		NodeURL: getRpcURL("blast-sepolia"),
		LogoURL: getLogoURL(ChainID_BLAST_SEPOLIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Blast Sepolia Explorer",
			RootURL: "https://sepolia.blastexplorer.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_TELOS,
		Type:    NetworkTypeMainnet,
		Name:    "telos",
		Title:   "Telos",
		NodeURL: getRpcURL("telos"),
		LogoURL: getLogoURL(ChainID_TELOS),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Telos Explorer",
			RootURL: "https://explorer.telos.net/network/",
		},
		NativeToken: &Currency{
			Symbol:   "TLOS",
			Name:     "TLOS",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_TELOS_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "telos-testnet",
		Title:   "Telos Testnet",
		NodeURL: getRpcURL("telos-testnet"),
		LogoURL: getLogoURL(ChainID_TELOS_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Telos Testnet Explorer",
			RootURL: "https://explorer-test.telos.net/network",
		},
		NativeToken: &Currency{
			Symbol:   "TLOS",
			Name:     "TLOS",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SKALE_NEBULA,
		Type:    NetworkTypeMainnet,
		Name:    "skale-nebula",
		Title:   "SKALE Nebula Gaming Hub",
		NodeURL: getRpcURL("skale-nebula"),
		LogoURL: getLogoURL(ChainID_SKALE_NEBULA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "SKALE Nebula Gaming Hub Explorer",
			RootURL: "https://green-giddy-denebola.explorer.mainnet.skalenodes.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sFUEL",
			Name:     "SKALE Fuel",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SKALE_NEBULA_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "skale-nebula-testnet",
		Title:   "SKALE Nebula Gaming Hub Testnet",
		NodeURL: getRpcURL("skale-nebula-testnet"),
		LogoURL: getLogoURL(ChainID_SKALE_NEBULA_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "SKALE Nebula Gaming Hub Testnet Explorer",
			RootURL: "https://lanky-ill-funny-testnet.explorer.testnet.skalenodes.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sFUEL",
			Name:     "SKALE Fuel",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SONEIUM,
		Type:    NetworkTypeMainnet,
		Name:    "soneium",
		Title:   "Soneium",
		NodeURL: getRpcURL("soneium"),
		LogoURL: getLogoURL(ChainID_SONEIUM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Soneium Explorer",
			RootURL: "https://soneium.blockscout.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SONEIUM_MINATO,
		Type:    NetworkTypeTestnet,
		Name:    "soneium-minato",
		Title:   "Soneium Minato (Testnet)",
		NodeURL: getRpcURL("soneium-minato"),
		LogoURL: getLogoURL(ChainID_SONEIUM_MINATO),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Soneium Minato Explorer",
			RootURL: "https://explorer-testnet.soneium.org/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_TOY_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "toy-testnet",
		Title:   "TOY (Testnet)",
		NodeURL: getRpcURL("toy-testnet"),
		LogoURL: getLogoURL(ChainID_TOY_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "TOY Testnet Explorer",
			RootURL: "https://toy-chain-testnet.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "TOY",
			Name:     "TOY",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_IMMUTABLE_ZKEVM,
		Type:    NetworkTypeMainnet,
		Name:    "immutable-zkevm",
		Title:   "Immutable zkEVM",
		NodeURL: getRpcURL("immutable-zkevm"),
		LogoURL: getLogoURL(ChainID_IMMUTABLE_ZKEVM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Immutable zkEVM Explorer",
			RootURL: "https://explorer.immutable.com/",
		},
		NativeToken: &Currency{
			Symbol:   "IMX",
			Name:     "IMX",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_IMMUTABLE_ZKEVM_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "immutable-zkevm-testnet",
		Title:   "Immutable zkEVM Testnet",
		NodeURL: getRpcURL("immutable-zkevm-testnet"),
		LogoURL: getLogoURL(ChainID_IMMUTABLE_ZKEVM_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Immutable zkEVM Testnet Explorer",
			RootURL: "https://explorer.testnet.immutable.com/",
		},
		NativeToken: &Currency{
			Symbol:   "IMX",
			Name:     "IMX",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_MOONBEAM,
		Type:    NetworkTypeMainnet,
		Name:    "moonbeam",
		Title:   "Moonbeam",
		NodeURL: getRpcURL("moonbeam"),
		LogoURL: getLogoURL(ChainID_MOONBEAM),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Moonscan",
			RootURL: "https://moonscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "GLMR",
			Name:     "GLMR",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_MOONBASE_ALPHA,
		Type:    NetworkTypeTestnet,
		Name:    "moonbase-alpha",
		Title:   "Moonbase Alpha",
		NodeURL: getRpcURL("moonbase-alpha"),
		LogoURL: getLogoURL(ChainID_MOONBASE_ALPHA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Moonscan (Moonbase Alpha)",
			RootURL: "https://moonbase.moonscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "GLMR",
			Name:     "GLMR",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ETHERLINK,
		Type:    NetworkTypeMainnet,
		Name:    "etherlink",
		Title:   "Etherlink",
		NodeURL: getRpcURL("etherlink"),
		LogoURL: getLogoURL(ChainID_ETHERLINK),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherlink Explorer",
			RootURL: "https://explorer.etherlink.com/",
		},
		NativeToken: &Currency{
			Symbol:   "XTZ",
			Name:     "Tez",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ETHERLINK_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "etherlink-testnet",
		Title:   "Etherlink Testnet",
		NodeURL: getRpcURL("etherlink-testnet"),
		LogoURL: getLogoURL(ChainID_ETHERLINK_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherlink Testnet Explorer",
			RootURL: "https://testnet.explorer.etherlink.com/",
		},
		NativeToken: &Currency{
			Symbol:   "XTZ",
			Name:     "Tez",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_MONAD,
		Type:    NetworkTypeMainnet,
		Name:    "monad",
		Title:   "Monad",
		NodeURL: getRpcURL("monad"),
		LogoURL: getLogoURL(ChainID_MONAD),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Monad Explorer",
			RootURL: "https://mainnet-beta.monvision.io/",
		},
		NativeToken: &Currency{
			Symbol:   "MON",
			Name:     "MON",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_MONAD_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "monad-testnet",
		Title:   "Monad Testnet",
		NodeURL: getRpcURL("monad-testnet"),
		LogoURL: getLogoURL(ChainID_MONAD_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Monad Testnet Explorer",
			RootURL: "https://testnet.monadexplorer.com/",
		},
		NativeToken: &Currency{
			Symbol:   "MON",
			Name:     "MON",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SOMNIA,
		Type:    NetworkTypeMainnet,
		Name:    "somnia",
		Title:   "Somnia",
		NodeURL: getRpcURL("somnia"),
		LogoURL: getLogoURL(ChainID_SOMNIA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Somnia Explorer",
			RootURL: "https://mainnet.somnia.w3us.site/",
		},
		NativeToken: &Currency{
			Symbol:   "SOMI",
			Name:     "SOMI",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SOMNIA_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "somnia-testnet",
		Title:   "Somnia Testnet",
		NodeURL: getRpcURL("somnia-testnet"),
		LogoURL: getLogoURL(ChainID_SOMNIA_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Somnia Testnet Explorer",
			RootURL: "https://somnia-testnet.socialscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "STT",
			Name:     "STT",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_INCENTIV_TESTNET_V2,
		Type:    NetworkTypeTestnet,
		Name:    "incentiv-testnet-v2",
		Title:   "Incentiv Testnet",
		NodeURL: getRpcURL("incentiv-testnet-v2"),
		LogoURL: getLogoURL(ChainID_INCENTIV_TESTNET_V2),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Incentiv Testnet Explorer",
			RootURL: "https://explorer.testnet.incentiv.net/",
		},
		NativeToken: &Currency{
			Symbol:   "TCENT",
			Name:     "TCENT",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_KATANA,
		Type:    NetworkTypeMainnet,
		Name:    "katana",
		Title:   "Katana",
		NodeURL: getRpcURL("katana"),
		LogoURL: getLogoURL(ChainID_KATANA),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Katana Explorer",
			RootURL: "https://katanascan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "ETH",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_SANDBOX_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "sandbox-testnet",
		Title:   "Sandbox Testnet",
		NodeURL: getRpcURL("sandbox-testnet"),
		LogoURL: getLogoURL(ChainID_SANDBOX_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Sandbox Testnet Explorer",
			RootURL: "https://sandbox-testnet.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "SAND",
			Name:     "SAND",
			Decimals: 18,
		},
	},
	{
		ChainID: ChainID_ARC_TESTNET,
		Type:    NetworkTypeTestnet,
		Name:    "arc-testnet",
		Title:   "Arc Testnet",
		NodeURL: getRpcURL("arc-testnet"),
		LogoURL: getLogoURL(ChainID_ARC_TESTNET),
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arc Testnet Explorer",
			RootURL: "https://1jr2dw1zdqvyes8u.blockscout.com/",
		},
		NativeToken: &Currency{
			Symbol:   "USDC",
			Name:     "USDC",
			Decimals: 6,
		},
	},
}

func getRpcURL(networkName string) string {
	return "https://nodes.sequence.app/" + networkName
}

func getLogoURL(chainID ChainID) string {
	return "https://assets.sequence.info/images/networks/medium/" + string(rune(chainID)) + ".webp"
}

func GetNetworkFromName(networkName string) *Network {
	for i := range All {
		if All[i].Name == networkName {
			return &All[i]
		}
	}
	return nil
}

func GetNetworkFromChainID(chainID ChainID) *Network {
	for i := range All {
		if All[i].ChainID == chainID {
			return &All[i]
		}
	}
	return nil
}
