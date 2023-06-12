package sequence

import (
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
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

// sequenceContextV1 are the deployed addresses of modules available on public networks.
var sequenceContextV1 = WalletContext{
	FactoryAddress:              common.HexToAddress("0xf9D09D634Fb818b05149329C1dcCFAeA53639d96"),
	MainModuleAddress:           common.HexToAddress("0xd01F11855bCcb95f88D7A48492F66410d4637313"),
	MainModuleUpgradableAddress: common.HexToAddress("0x7EFE6cE415956c5f80C6530cC6cc81b4808F6118"),
	GuestModuleAddress:          common.HexToAddress("0x02390F3E6E5FD1C6786CB78FD3027C117a9955A7"),
	UtilsAddress:                common.HexToAddress("0xd130B43062D875a4B7aF3f8fc036Bc6e9D3E1B3E"),
}

var sequenceContextV2 = WalletContext{
	FactoryAddress:              common.HexToAddress("0xFaA5c0b14d1bED5C888Ca655B9a8A5911F78eF4A"),
	MainModuleAddress:           common.HexToAddress("0xfBf8f1A5E00034762D928f46d438B947f5d4065d"),
	MainModuleUpgradableAddress: common.HexToAddress("0x4222dcA3974E39A8b41c411FeDDE9b09Ae14b911"),
	GuestModuleAddress:          common.HexToAddress("0xfea230Ee243f88BC698dD8f1aE93F8301B6cdfaE"),
}

// V1SequenceContext returns copy of the package-level internal variable, to prevent change
// by other packages.
func V1SequenceContext() WalletContext {
	return sequenceContextV1
}

func V2SequenceContext() WalletContext {
	return sequenceContextV2
}

func SequenceContext() WalletContext {
	return V2SequenceContext()
}

func SequenceContexts() WalletContexts {
	return WalletContexts{
		1: sequenceContextV1,
		2: sequenceContextV2,
	}
}

func SequenceContextForWalletConfig(walletConfig core.WalletConfig) WalletContext {
	if _, ok := walletConfig.(*v1.WalletConfig); ok {
		return sequenceContextV1
	} else if _, ok := walletConfig.(*v2.WalletConfig); ok {
		return sequenceContextV2
	}
	return WalletContext{}
}
