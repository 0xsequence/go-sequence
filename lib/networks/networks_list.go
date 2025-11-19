package networks

// List of supported chains ids
const (
	ChainID_MAINNET                 = 1
	ChainID_ROPSTEN                 = 3
	ChainID_RINKEBY                 = 4
	ChainID_GOERLI                  = 5
	ChainID_OPTIMISM                = 10
	ChainID_TELOS                   = 40
	ChainID_TELOS_TESTNET           = 41
	ChainID_KOVAN                   = 42
	ChainID_BSC                     = 56
	ChainID_OPTIMISM_KOVAN          = 69
	ChainID_BSC_TESTNET             = 97
	ChainID_GNOSIS                  = 100
	ChainID_POLYGON                 = 137
	ChainID_MONAD                   = 143
	ChainID_OPTIMISM_GOERLI         = 420
	ChainID_POLYGON_ZKEVM           = 1101
	ChainID_MOONBEAM                = 1284
	ChainID_MOONBASE_ALPHA          = 1287
	ChainID_SEI_TESTNET             = 1328
	ChainID_SEI                     = 1329
	ChainID_SONEIUM                 = 1868
	ChainID_SONEIUM_MINATO          = 1946
	ChainID_B3_SEPOLIA              = 1993
	ChainID_SOMNIA                  = 5031
	ChainID_SANDBOX_TESTNET         = 6252
	ChainID_B3                      = 8333
	ChainID_BASE                    = 8453
	ChainID_MONAD_TESTNET           = 10143
	ChainID_INCENTIV_TESTNET        = 11690
	ChainID_IMMUTABLE_ZKEVM         = 13371
	ChainID_IMMUTABLE_ZKEVM_TESTNET = 13473
	ChainID_HOMEVERSE               = 19011
	ChainID_INCENTIV_TESTNET_V2     = 28802
	ChainID_HARDHAT                 = 31337
	ChainID_HARDHAT2                = 31338
	ChainID_APECHAIN_TESTNET        = 33111
	ChainID_APECHAIN                = 33139
	ChainID_HOMEVERSE_TESTNET       = 40875
	ChainID_ARBITRUM                = 42161
	ChainID_ARBITRUM_NOVA           = 42170
	ChainID_ETHERLINK               = 42793
	ChainID_AVALANCHE_TESTNET       = 43113
	ChainID_AVALANCHE               = 43114
	ChainID_SOMNIA_TESTNET          = 50312
	ChainID_MUMBAI                  = 80001
	ChainID_AMOY                    = 80002
	ChainID_BLAST                   = 81457
	ChainID_BASE_GOERLI             = 84531
	ChainID_BASE_SEPOLIA            = 84532
	ChainID_BORNE_TESTNET           = 94984
	ChainID_ETHERLINK_TESTNET       = 128123
	ChainID_ARBITRUM_GOERLI         = 421613
	ChainID_ARBITRUM_SEPOLIA        = 421614
	ChainID_XAI                     = 660279
	ChainID_KATANA                  = 747474
	ChainID_ARC_TESTNET             = 5042002
	ChainID_SEPOLIA                 = 11155111
	ChainID_OPTIMISM_SEPOLIA        = 11155420
	ChainID_TOY_TESTNET             = 21000000
	ChainID_SKALE_NEBULA_TESTNET    = 37084624
	ChainID_BLAST_SEPOLIA           = 168587773
	ChainID_SKALE_NEBULA            = 1482601649
	ChainID_XAI_SEPOLIA             = 37714555429
)

var all = Networks{
	"mainnet": &Network{
		ChainID: ChainID_MAINNET,
		Name:    "Ethereum",
		Title:   "Ethereum",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1.webp",
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
	"ropsten": &Network{
		ChainID: ChainID_ROPSTEN,
		Name:    "Ropsten",
		Title:   "Ropsten",
		LogoURL: "https://assets.sequence.info/images/networks/medium/3.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Ropsten)",
			RootURL: "https://ropsten.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "roETH",
			Name:     "Ropsten Ether",
			Decimals: 18,
		},
		ENSAddress: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	},
	"rinkeby": &Network{
		ChainID: ChainID_RINKEBY,
		Name:    "Rinkeby",
		Title:   "Rinkeby",
		LogoURL: "https://assets.sequence.info/images/networks/medium/4.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Rinkeby)",
			RootURL: "https://rinkeby.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "rETH",
			Name:     "Rinkeby Ether",
			Decimals: 18,
		},
		ENSAddress: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	},
	"goerli": &Network{
		ChainID: ChainID_GOERLI,
		Name:    "Goerli",
		Title:   "Goerli",
		LogoURL: "https://assets.sequence.info/images/networks/medium/5.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Goerli)",
			RootURL: "https://goerli.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "gETH",
			Name:     "Goerli Ether",
			Decimals: 18,
		},
		ENSAddress: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	},
	"optimism": &Network{
		ChainID: ChainID_OPTIMISM,
		Name:    "Optimism",
		Title:   "Optimism",
		LogoURL: "https://assets.sequence.info/images/networks/medium/10.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism)",
			RootURL: "https://optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"telos": &Network{
		ChainID: ChainID_TELOS,
		Name:    "Telos",
		Title:   "Telos",
		LogoURL: "https://assets.sequence.info/images/networks/medium/40.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Telos Explorer",
			RootURL: "https://explorer.telos.net/network/",
		},
		NativeToken: &Currency{
			Symbol:   "TLOS",
			Name:     "TLOS",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"telos-testnet": &Network{
		ChainID: ChainID_TELOS_TESTNET,
		Name:    "Telos Testnet",
		Title:   "Telos Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/41.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Telos Testnet Explorer",
			RootURL: "https://explorer-test.telos.net/network",
		},
		NativeToken: &Currency{
			Symbol:   "TLOS",
			Name:     "TLOS",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"kovan": &Network{
		ChainID: ChainID_KOVAN,
		Name:    "Kovan",
		Title:   "Kovan",
		LogoURL: "https://assets.sequence.info/images/networks/medium/42.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Kovan)",
			RootURL: "https://kovan.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "kETH",
			Name:     "Kovan Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"bsc": &Network{
		ChainID: ChainID_BSC,
		Name:    "BNB Smart Chain",
		Title:   "BNB Smart Chain",
		LogoURL: "https://assets.sequence.info/images/networks/medium/56.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "BSCScan",
			RootURL: "https://bscscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "BNB",
			Name:     "BNB",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"optimism-kovan": &Network{
		ChainID: ChainID_OPTIMISM_KOVAN,
		Name:    "Optimism Kovan",
		Title:   "Optimism Kovan",
		LogoURL: "https://assets.sequence.info/images/networks/medium/69.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism Kovan)",
			RootURL: "https://kovan-optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "kETH",
			Name:     "Kovan Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"bsc-testnet": &Network{
		ChainID: ChainID_BSC_TESTNET,
		Name:    "BNB Smart Chain Testnet",
		Title:   "BNB Smart Chain Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/97.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "BSCScan (Testnet)",
			RootURL: "https://testnet.bscscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "tBNB",
			Name:     "Testnet BNB",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"gnosis": &Network{
		ChainID: ChainID_GNOSIS,
		Name:    "Gnosis Chain",
		Title:   "Gnosis Chain",
		LogoURL: "https://assets.sequence.info/images/networks/medium/100.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Gnosis Chain Explorer",
			RootURL: "https://blockscout.com/xdai/mainnet/",
		},
		NativeToken: &Currency{
			Symbol:   "XDAI",
			Name:     "XDAI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"polygon": &Network{
		ChainID: ChainID_POLYGON,
		Name:    "Polygon",
		Title:   "Polygon",
		LogoURL: "https://assets.sequence.info/images/networks/medium/137.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Polygonscan",
			RootURL: "https://polygonscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "POL",
			Name:     "POL",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"monad": &Network{
		ChainID: ChainID_MONAD,
		Name:    "Monad",
		Title:   "Monad",
		LogoURL: "https://assets.sequence.info/images/networks/medium/143.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Monad Explorer",
			RootURL: "https://mainnet-beta.monvision.io/",
		},
		NativeToken: &Currency{
			Symbol:   "MON",
			Name:     "MON",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"optimism-goerli": &Network{
		ChainID: ChainID_OPTIMISM_GOERLI,
		Name:    "Optimism Goerli",
		Title:   "Optimism Goerli",
		LogoURL: "https://assets.sequence.info/images/networks/medium/420.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism Goerli)",
			RootURL: "https://goerli-optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "gETH",
			Name:     "Goerli Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"polygon-zkevm": &Network{
		ChainID: ChainID_POLYGON_ZKEVM,
		Name:    "Polygon zkEVM",
		Title:   "Polygon zkEVM",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1101.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Polygonscan (zkEVM)",
			RootURL: "https://zkevm.polygonscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"moonbeam": &Network{
		ChainID: ChainID_MOONBEAM,
		Name:    "Moonbeam",
		Title:   "Moonbeam",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1284.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Moonscan",
			RootURL: "https://moonscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "GLMR",
			Name:     "GLMR",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"moonbase-alpha": &Network{
		ChainID: ChainID_MOONBASE_ALPHA,
		Name:    "Moonbase Alpha",
		Title:   "Moonbase Alpha",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1287.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Moonscan (Moonbase Alpha)",
			RootURL: "https://moonbase.moonscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "GLMR",
			Name:     "GLMR",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"sei-testnet": &Network{
		ChainID: ChainID_SEI_TESTNET,
		Name:    "Sei Testnet",
		Title:   "Sei Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1328.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Sei Testnet Explorer",
			RootURL: "https://seitrace.com/?chain=atlantic-2",
		},
		NativeToken: &Currency{
			Symbol:   "SEI",
			Name:     "SEI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"sei": &Network{
		ChainID: ChainID_SEI,
		Name:    "Sei",
		Title:   "Sei",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1329.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "SEI Explorer",
			RootURL: "https://seitrace.com/?chain=pacific-1",
		},
		NativeToken: &Currency{
			Symbol:   "SEI",
			Name:     "SEI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"soneium": &Network{
		ChainID: ChainID_SONEIUM,
		Name:    "Soneium",
		Title:   "Soneium",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1868.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Soneium Explorer",
			RootURL: "https://soneium.blockscout.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"soneium-minato": &Network{
		ChainID: ChainID_SONEIUM_MINATO,
		Name:    "Soneium Minato (Testnet)",
		Title:   "Soneium Minato (Testnet)",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1946.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Soneium Minato Explorer",
			RootURL: "https://explorer-testnet.soneium.org/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"b3-sepolia": &Network{
		ChainID: ChainID_B3_SEPOLIA,
		Name:    "B3 Sepolia",
		Title:   "B3 Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1993.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "B3 Sepolia Explorer",
			RootURL: "https://sepolia.explorer.b3.fun/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"somnia": &Network{
		ChainID: ChainID_SOMNIA,
		Name:    "Somnia",
		Title:   "Somnia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/5031.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Somnia Explorer",
			RootURL: "https://mainnet.somnia.w3us.site/",
		},
		NativeToken: &Currency{
			Symbol:   "SOMI",
			Name:     "SOMI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"sandbox-testnet": &Network{
		ChainID: ChainID_SANDBOX_TESTNET,
		Name:    "Sandbox Testnet",
		Title:   "Sandbox Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/6252.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Sandbox Testnet Explorer",
			RootURL: "https://sandbox-testnet.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "SAND",
			Name:     "SAND",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"b3": &Network{
		ChainID: ChainID_B3,
		Name:    "B3",
		Title:   "B3",
		LogoURL: "https://assets.sequence.info/images/networks/medium/8333.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "B3 Explorer",
			RootURL: "https://explorer.b3.fun/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"base": &Network{
		ChainID: ChainID_BASE,
		Name:    "Base (Coinbase)",
		Title:   "Base (Coinbase)",
		LogoURL: "https://assets.sequence.info/images/networks/medium/8453.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Base Explorer",
			RootURL: "https://basescan.org/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"monad-testnet": &Network{
		ChainID: ChainID_MONAD_TESTNET,
		Name:    "Monad Testnet",
		Title:   "Monad Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/10143.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Monad Testnet Explorer",
			RootURL: "https://testnet.monadexplorer.com/",
		},
		NativeToken: &Currency{
			Symbol:   "MON",
			Name:     "MON",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"incentiv-testnet": &Network{
		ChainID: ChainID_INCENTIV_TESTNET,
		Name:    "Incentiv Testnet",
		Title:   "Incentiv Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/11690.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Incentiv Testnet Explorer",
			RootURL: "https://explorer.testnet.incentiv.net/",
		},
		NativeToken: &Currency{
			Symbol:   "CENT",
			Name:     "CENT",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"immutable-zkevm": &Network{
		ChainID: ChainID_IMMUTABLE_ZKEVM,
		Name:    "Immutable zkEVM",
		Title:   "Immutable zkEVM",
		LogoURL: "https://assets.sequence.info/images/networks/medium/13371.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Immutable zkEVM Explorer",
			RootURL: "https://explorer.immutable.com/",
		},
		NativeToken: &Currency{
			Symbol:   "IMX",
			Name:     "IMX",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"immutable-zkevm-testnet": &Network{
		ChainID: ChainID_IMMUTABLE_ZKEVM_TESTNET,
		Name:    "Immutable zkEVM Testnet",
		Title:   "Immutable zkEVM Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/13473.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Immutable zkEVM Testnet Explorer",
			RootURL: "https://explorer.testnet.immutable.com/",
		},
		NativeToken: &Currency{
			Symbol:   "IMX",
			Name:     "IMX",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"homeverse": &Network{
		ChainID: ChainID_HOMEVERSE,
		Name:    "Oasys Homeverse",
		Title:   "Oasys Homeverse",
		LogoURL: "https://assets.sequence.info/images/networks/medium/19011.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Oasys Homeverse Explorer",
			RootURL: "https://explorer.oasys.homeverse.games/",
		},
		NativeToken: &Currency{
			Symbol:   "OAS",
			Name:     "OAS",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"incentiv-testnet-v2": &Network{
		ChainID: ChainID_INCENTIV_TESTNET_V2,
		Name:    "Incentiv Testnet v2",
		Title:   "Incentiv Testnet v2",
		LogoURL: "https://assets.sequence.info/images/networks/medium/28802.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Incentiv Testnet Explorer v2",
			RootURL: "https://explorer.testnet.incentiv.net/",
		},
		NativeToken: &Currency{
			Symbol:   "TCENT",
			Name:     "TCENT",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"hardhat": &Network{
		ChainID: ChainID_HARDHAT,
		Name:    "Hardhat (local testnet)",
		Title:   "Hardhat (local testnet)",
		LogoURL: "",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Explorer",
			RootURL: "",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"hardhat2": &Network{
		ChainID: ChainID_HARDHAT2,
		Name:    "Hardhat (local testnet)",
		Title:   "Hardhat (local testnet)",
		LogoURL: "",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Explorer",
			RootURL: "",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"apechain-testnet": &Network{
		ChainID: ChainID_APECHAIN_TESTNET,
		Name:    "APE Chain Testnet",
		Title:   "APE Chain Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/33111.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "APE Chain Explorer",
			RootURL: "https://curtis.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "APE",
			Name:     "ApeCoin",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"apechain": &Network{
		ChainID: ChainID_APECHAIN,
		Name:    "APE Chain",
		Title:   "APE Chain",
		LogoURL: "https://assets.sequence.info/images/networks/medium/33139.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "APE Chain Explorer",
			RootURL: "https://apechain.calderaexplorer.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "APE",
			Name:     "ApeCoin",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"homeverse-testnet": &Network{
		ChainID: ChainID_HOMEVERSE_TESTNET,
		Name:    "Oasys Homeverse Testnet",
		Title:   "Oasys Homeverse Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/40875.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Oasys Homeverse Explorer (Testnet)",
			RootURL: "https://explorer.testnet.oasys.homeverse.games/",
		},
		NativeToken: &Currency{
			Symbol:   "tOAS",
			Name:     "Testnet OAS",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"arbitrum": &Network{
		ChainID: ChainID_ARBITRUM,
		Name:    "Arbitrum One",
		Title:   "Arbitrum One",
		LogoURL: "https://assets.sequence.info/images/networks/medium/42161.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan",
			RootURL: "https://arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"arbitrum-nova": &Network{
		ChainID: ChainID_ARBITRUM_NOVA,
		Name:    "Arbitrum Nova",
		Title:   "Arbitrum Nova",
		LogoURL: "https://assets.sequence.info/images/networks/medium/42170.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan Nova",
			RootURL: "https://nova.arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"etherlink": &Network{
		ChainID: ChainID_ETHERLINK,
		Name:    "Etherlink",
		Title:   "Etherlink",
		LogoURL: "https://assets.sequence.info/images/networks/medium/42793.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherlink Explorer",
			RootURL: "https://explorer.etherlink.com/",
		},
		NativeToken: &Currency{
			Symbol:   "XTZ",
			Name:     "Tez",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"avalanche-testnet": &Network{
		ChainID: ChainID_AVALANCHE_TESTNET,
		Name:    "Avalanche Testnet",
		Title:   "Avalanche Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/43113.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Snowtrace (Testnet)",
			RootURL: "https://subnets-test.avax.network/c-chain/",
		},
		NativeToken: &Currency{
			Symbol:   "tAVAX",
			Name:     "Testnet AVAX",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"avalanche": &Network{
		ChainID: ChainID_AVALANCHE,
		Name:    "Avalanche",
		Title:   "Avalanche",
		LogoURL: "https://assets.sequence.info/images/networks/medium/43114.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Snowtrace",
			RootURL: "https://subnets.avax.network/c-chain/",
		},
		NativeToken: &Currency{
			Symbol:   "AVAX",
			Name:     "AVAX",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"somnia-testnet": &Network{
		ChainID: ChainID_SOMNIA_TESTNET,
		Name:    "Somnia Testnet",
		Title:   "Somnia Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/50312.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Somnia Testnet Explorer",
			RootURL: "https://shannon-explorer.somnia.network/",
		},
		NativeToken: &Currency{
			Symbol:   "STT",
			Name:     "STT",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"mumbai": &Network{
		ChainID: ChainID_MUMBAI,
		Name:    "Polygon Mumbai",
		Title:   "Polygon Mumbai",
		LogoURL: "https://assets.sequence.info/images/networks/medium/80001.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Polygonscan (Mumbai)",
			RootURL: "https://mumbai.polygonscan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "mMATIC",
			Name:     "Mumbai Polygon",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"amoy": &Network{
		ChainID: ChainID_AMOY,
		Name:    "Polygon Amoy",
		Title:   "Polygon Amoy",
		LogoURL: "https://assets.sequence.info/images/networks/medium/80002.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "OKLink (Amoy)",
			RootURL: "https://www.oklink.com/amoy/",
		},
		NativeToken: &Currency{
			Symbol:   "aPOL",
			Name:     "Amoy POL",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"blast": &Network{
		ChainID: ChainID_BLAST,
		Name:    "Blast",
		Title:   "Blast",
		LogoURL: "https://assets.sequence.info/images/networks/medium/81457.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Blast Explorer",
			RootURL: "https://blastscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"base-goerli": &Network{
		ChainID: ChainID_BASE_GOERLI,
		Name:    "Base Goerli",
		Title:   "Base Goerli",
		LogoURL: "https://assets.sequence.info/images/networks/medium/84531.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Base Goerli Explorer",
			RootURL: "https://goerli.basescan.org/",
		},
		NativeToken: &Currency{
			Symbol:   "gETH",
			Name:     "Goerli Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"base-sepolia": &Network{
		ChainID: ChainID_BASE_SEPOLIA,
		Name:    "Base Sepolia",
		Title:   "Base Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/84532.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Base Sepolia Explorer",
			RootURL: "https://base-sepolia.blockscout.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"borne-testnet": &Network{
		ChainID: ChainID_BORNE_TESTNET,
		Name:    "Borne Testnet",
		Title:   "Borne Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/94984.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Borne Testnet Explorer",
			RootURL: "https://subnets-test.avax.network/bornegfdn",
		},
		NativeToken: &Currency{
			Symbol:   "BORNE",
			Name:     "BORNE",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"etherlink-testnet": &Network{
		ChainID: ChainID_ETHERLINK_TESTNET,
		Name:    "Etherlink Testnet",
		Title:   "Etherlink Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/128123.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherlink Testnet Explorer",
			RootURL: "https://testnet.explorer.etherlink.com/",
		},
		NativeToken: &Currency{
			Symbol:   "XTZ",
			Name:     "Tez",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"arbitrum-goerli": &Network{
		ChainID: ChainID_ARBITRUM_GOERLI,
		Name:    "Arbitrum Goerli",
		Title:   "Arbitrum Goerli",
		LogoURL: "https://assets.sequence.info/images/networks/medium/421613.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan (Goerli Testnet)",
			RootURL: "https://testnet.arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "gETH",
			Name:     "Goerli Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"arbitrum-sepolia": &Network{
		ChainID: ChainID_ARBITRUM_SEPOLIA,
		Name:    "Arbitrum Sepolia",
		Title:   "Arbitrum Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/421614.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arbiscan (Sepolia Testnet)",
			RootURL: "https://sepolia.arbiscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"xai": &Network{
		ChainID: ChainID_XAI,
		Name:    "Xai",
		Title:   "Xai",
		LogoURL: "https://assets.sequence.info/images/networks/medium/660279.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Xai Explorer",
			RootURL: "https://explorer.xai-chain.net/",
		},
		NativeToken: &Currency{
			Symbol:   "XAI",
			Name:     "XAI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"katana": &Network{
		ChainID: ChainID_KATANA,
		Name:    "Katana",
		Title:   "Katana",
		LogoURL: "https://assets.sequence.info/images/networks/medium/747474.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Katana",
			RootURL: "https://katanascan.com/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "ETH",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"arc-testnet": &Network{
		ChainID: ChainID_ARC_TESTNET,
		Name:    "Arc Testnet",
		Title:   "Arc Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/5042002.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Arc Testnet Explorer",
			RootURL: "https://testnet.arcscan.app/",
		},
		NativeToken: &Currency{
			Symbol:   "USDC",
			Name:     "USDC",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"sepolia": &Network{
		ChainID: ChainID_SEPOLIA,
		Name:    "Sepolia",
		Title:   "Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/11155111.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Sepolia)",
			RootURL: "https://sepolia.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"optimism-sepolia": &Network{
		ChainID: ChainID_OPTIMISM_SEPOLIA,
		Name:    "Optimism Sepolia",
		Title:   "Optimism Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/11155420.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Etherscan (Optimism Sepolia)",
			RootURL: "https://sepolia-optimistic.etherscan.io/",
		},
		NativeToken: &Currency{
			Symbol:   "sETH",
			Name:     "Sepolia Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"toy-testnet": &Network{
		ChainID: ChainID_TOY_TESTNET,
		Name:    "TOY (Testnet)",
		Title:   "TOY (Testnet)",
		LogoURL: "https://assets.sequence.info/images/networks/medium/21000000.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "TOY Testnet Explorer",
			RootURL: "https://toy-chain-testnet.explorer.caldera.xyz/",
		},
		NativeToken: &Currency{
			Symbol:   "TOY",
			Name:     "TOY",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"skale-nebula-testnet": &Network{
		ChainID: ChainID_SKALE_NEBULA_TESTNET,
		Name:    "SKALE Nebula Gaming Hub Testnet",
		Title:   "SKALE Nebula Gaming Hub Testnet",
		LogoURL: "https://assets.sequence.info/images/networks/medium/37084624.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "SKALE Nebula Gaming Hub Testnet Explorer",
			RootURL: "https://lanky-ill-funny-testnet.explorer.testnet.skalenodes.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sFUEL",
			Name:     "SKALE Fuel",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"blast-sepolia": &Network{
		ChainID: ChainID_BLAST_SEPOLIA,
		Name:    "Blast Sepolia",
		Title:   "Blast Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/168587773.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Blast Sepolia Explorer",
			RootURL: "https://sepolia.blastexplorer.io/",
		},
		NativeToken: &Currency{
			Symbol:   "ETH",
			Name:     "Ether",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"skale-nebula": &Network{
		ChainID: ChainID_SKALE_NEBULA,
		Name:    "SKALE Nebula Gaming Hub",
		Title:   "SKALE Nebula Gaming Hub",
		LogoURL: "https://assets.sequence.info/images/networks/medium/1482601649.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "SKALE Nebula Gaming Hub Explorer",
			RootURL: "https://green-giddy-denebola.explorer.mainnet.skalenodes.com/",
		},
		NativeToken: &Currency{
			Symbol:   "sFUEL",
			Name:     "SKALE Fuel",
			Decimals: 18,
		},
		ENSAddress: "",
	},
	"xai-sepolia": &Network{
		ChainID: ChainID_XAI_SEPOLIA,
		Name:    "Xai Sepolia",
		Title:   "Xai Sepolia",
		LogoURL: "https://assets.sequence.info/images/networks/medium/37714555429.webp",
		BlockExplorer: &BlockExplorerConfig{
			Name:    "Xai Sepolia Explorer",
			RootURL: "https://testnet-explorer-v2.xai-chain.net/",
		},
		NativeToken: &Currency{
			Symbol:   "sXAI",
			Name:     "Sepolia XAI",
			Decimals: 18,
		},
		ENSAddress: "",
	},
}

func GetAllNetworks() Networks {
	return all
}

func GetByChainID(chainID int) (Network, bool) {
	network, ok := all.GetByChainID(chainID)
	if !ok {
		return Network{}, false
	}
	return network, true
}

func GetByName(name string) (Network, bool) {
	network, ok := all.GetByName(name)
	if !ok {
		return Network{}, false
	}
	return network, true
}
