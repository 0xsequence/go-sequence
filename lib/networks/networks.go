package networks

import (
	"github.com/0xsequence/ethkit/ethproviders"
)

type Networks map[string]*Network

type Network struct {
	ChainID    uint64      `toml:"chain_id"    json:"chainId,omitempty"`
	Name       string      `toml:"name"        json:"name,omitempty"`
	Title      string      `toml:"title"       json:"title,omitempty"`
	Type       NetworkType `toml:"type"        json:"type,omitempty"`
	LogoURL    string      `toml:"logo_url"    json:"logoUrl,omitempty"`
	ENSAddress string      `toml:"ens_address" json:"ensAddress,omitempty"`
	Deprecated bool        `toml:"deprecated"  json:"deprecated,omitempty"`
	Disabled   bool        `toml:"disabled"    json:"disabled,omitempty"`
	WSEnabled  bool        `toml:"ws_enabled"  json:"wsEnabled,omitempty"`
	WSURL      string      `toml:"ws_url"      json:"wsUrl,omitempty"`

	BlockExplorer *BlockExplorerConfig `toml:"block_explorer" json:"blockExplorer,omitempty"`

	NodeURL            string `toml:"node_url"             json:"nodeUrl,omitempty"`
	IndexerURL         string `toml:"indexer_url"          json:"indexerUrl,omitempty"`
	RelayerURL         string `toml:"relayer_url"          json:"relayerUrl,omitempty"`
	InternalNodeURL    string `toml:"internal_node_url"    json:"-"`
	InternalIndexerURL string `toml:"internal_indexer_url" json:"-"`
	InternalRelayerURL string `toml:"internal_relayer_url" json:"-"`

	NativeToken *Currency   `toml:"native_token" json:"nativeToken,omitempty"`
	Currencies  []*Currency `toml:"currencies"   json:"currencies,omitempty"`
}

type BlockExplorerConfig struct {
	Name       string `toml:"name"         json:"name,omitempty"`
	RootURL    string `toml:"root_url"     json:"rootUrl,omitempty"`
	AddressURL string `toml:"address_url"  json:"addressUrl,omitempty"`
	TxnHashURL string `toml:"txn_hash_url" json:"txnHashUrl,omitempty"`
}

type Currency struct {
	ContractAddress *string       `toml:"contract_address" json:"contractAddress,omitempty"`
	Name            string        `toml:"name"             json:"name,omitempty"`
	Symbol          string        `toml:"symbol"           json:"symbol,omitempty"`
	Decimals        uint64        `toml:"decimals"         json:"decimals,omitempty"`
	ImageURL        string        `toml:"image_url"        json:"imageUrl,omitempty"`
	Native          bool          `toml:"native"           json:"native,omitempty"`
	Default         bool          `toml:"default"          json:"default,omitempty"`
	Group           CurrencyGroup `toml:"group"            json:"group,omitempty"`
}

type NetworkType uint8

const (
	NetworkTypeUnknown NetworkType = 0
	NetworkTypeMainnet NetworkType = 1
	NetworkTypeTestnet NetworkType = 2
)

var OrderSideName = map[uint8]string{
	0: "unknown",
	1: "mainnet",
	2: "testnet",
}

var OrderSideValue = map[string]uint8{
	"unknown": 0,
	"mainnet": 1,
	"testnet": 2,
}

func (x NetworkType) String() string {
	return OrderSideName[uint8(x)]
}

func (x NetworkType) MarshalText() ([]byte, error) {
	return []byte(OrderSideName[uint8(x)]), nil
}

func (x *NetworkType) UnmarshalText(b []byte) error {
	*x = NetworkType(OrderSideValue[string(b)])
	return nil
}

// Grouping currencies for cross chain purchases
type CurrencyGroup uint8

const (
	CurrencyGroupNone CurrencyGroup = iota
	CurrencyGroupUSDC
	CurrencyGroupUSDCTestnet
	CurrencyGroupWETH
)

var CurrencyGroupName = map[uint8]string{
	0: "none",
	1: "usdc",
	2: "usdc_testnet",
	3: "weth",
}

var CurrencyGroupValue = map[string]uint8{
	"none":         0,
	"usdc":         1,
	"usdc_testnet": 2,
	"weth":         3,
}

func (x CurrencyGroup) String() string {
	return CurrencyGroupName[uint8(x)]
}

func (x CurrencyGroup) MarshalText() ([]byte, error) {
	return []byte(CurrencyGroupName[uint8(x)]), nil
}

func (x *CurrencyGroup) UnmarshalText(b []byte) error {
	*x = CurrencyGroup(CurrencyGroupValue[string(b)])
	return nil
}

func (n *Network) IsMainnet() bool {
	return n.Type == NetworkTypeMainnet
}

func (n *Network) IsTesnet() bool {
	return n.Type == NetworkTypeTestnet
}

func (n Networks) Active() Networks {
	res := Networks{}
	for name, network := range n {
		if !network.Disabled {
			res[name] = network
		}
	}

	return res
}

func (n Networks) EthProvidersConfig() ethproviders.Config {
	res := ethproviders.Config{}
	for _, network := range n {
		res[network.Name] = ethproviders.NetworkConfig{
			ID:        network.ChainID,
			URL:       network.NodeURL,
			WSEnabled: network.WSEnabled,
			WSURL:     network.WSURL,
			Testnet:   network.IsTesnet(),
			Disabled:  network.Disabled,
		}
	}

	return res
}
