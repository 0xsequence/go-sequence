package sequence

import (
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// WalletContext is the module addresses deployed on a network, aka the context / environment
// of the Sequence Smart Wallet system on Ethereum.
type WalletContext struct {
	FactoryAddress              common.Address `json:"factory" toml:"factory_address"`
	MainModuleAddress           common.Address `json:"mainModule" toml:"main_module_address"`
	MainModuleUpgradableAddress common.Address `json:"mainModuleUpgradable" toml:"main_module_upgradable_address"`
	GuestModuleAddress          common.Address `json:"guestModule" toml:"guest_module_address"`
	UtilsAddress                common.Address `json:"utils" toml:"utils_address"`
}

// A map of a wallet context for each version
type WalletContexts map[uint16]WalletContext

type NetworkConfig struct {
	Name       string
	ChainID    big.Int
	ENSAddress *common.Address

	RpcURL   string
	Provider *ethrpc.Provider

	RelayerURL *string // optional, one of the these should be set
	Relayer    Relayer

	IndexerURL *string // optional, one of these should be set
	// Indexer Indexer

	IsDefaultChain bool
	IsAuthChain    bool

	SequenceAPIURL string
}

type Networks []NetworkConfig

// TODO, etc.....
// var MainnetNetworks = createNetworkConfig()
//
// var TestnetNetworks = createNetworkConfig()

// sequenceContext are the deployed addresses of modules available on public networks.
var sequenceContext = WalletContext{
	FactoryAddress:              common.HexToAddress("0xf9D09D634Fb818b05149329C1dcCFAeA53639d96"),
	MainModuleAddress:           common.HexToAddress("0xd01F11855bCcb95f88D7A48492F66410d4637313"),
	MainModuleUpgradableAddress: common.HexToAddress("0x7EFE6cE415956c5f80C6530cC6cc81b4808F6118"),
	GuestModuleAddress:          common.HexToAddress("0x02390F3E6E5FD1C6786CB78FD3027C117a9955A7"),
	UtilsAddress:                common.HexToAddress("0xd130B43062D875a4B7aF3f8fc036Bc6e9D3E1B3E"),
}

var sequenceContextV2 = WalletContext{
	FactoryAddress:              common.HexToAddress("0x0D7604Bdf2cAcc2943b6388e1c26c3C33213f673"),
	MainModuleAddress:           common.HexToAddress("0xA507eF52f3fd34dd54566bf3055fA66bdabE2ef3"),
	MainModuleUpgradableAddress: common.HexToAddress("0x13Cc7b579e1acfDc8aD1F9996dd38ff744818a34"),
	GuestModuleAddress:          common.HexToAddress("0xCcB6cA914c20fAde6F2be5827eE40d899076ac2A"),
}

// SequenceContext returns copy of the package-level internal variable, to prevent change
// by other packages.
func SequenceContext() WalletContext {
	return sequenceContext
}

func SequenceContextV2() WalletContext {
	return sequenceContextV2
}

func SequenceContexts() WalletContexts {
	return WalletContexts{
		1: sequenceContext,
		2: sequenceContextV2,
	}
}
