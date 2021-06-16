package sequence

import (
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// WalletContext is the module addresses deployed on a network, aka the context / environment
// of the Sequence Smart Wallet system on Ethereum.
type WalletContext struct {
	FactoryAddress              common.Address `json:"factory"`
	MainModuleAddress           common.Address `json:"mainModule"`
	MainModuleUpgradableAddress common.Address `json:"mainModuleUpgradable"`
	GuestModuleAddress          common.Address `json:"guestModule"`
	UtilsAddress                common.Address `json:"utils"`
}

type NetworkConfig struct {
	Name       string
	ChainID    big.Int
	ENSAddress string // TODO: address?

	RpcURL   string // TODO: rename to NodeURL ? ..
	Provider *ethrpc.Provider

	Relayer interface{} // TODO .. or add RelayerURL
	// TODO: indexer..

	IsDefaultChain bool
	IsAuthChain    bool

	SequenceAPIURL string
}

// type ChainID interface{} // TOOD: hmm.. do we want this..? and add diff methods..?
type ChainID big.Int

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

// SequenceContext returns copy of the package-level internal variable, to prevent change
// by other packages.
func SequenceContext() WalletContext {
	return sequenceContext
}
