package networks

import (
	"github.com/0xsequence/ethkit/ethproviders"
)

type Networks map[string]*Network

type Network struct {
	ID    uint64      `toml:"id"`
	Name  string      `toml:"name"`
	Title string      `toml:"title"`
	Type  NetworkType `toml:"type"`

	LogoURL string `toml:"logoUrl"`

	ENSAddress string `toml:"ensAddress"`

	AuthChain  bool `toml:"authChain"`
	Deprecated bool `toml:"deprecated"`
	Disabled   bool `toml:"disabled"`

	NodeURL    string `toml:"nodeUrl"`
	IndexerURL string `toml:"indexerUrl"`
	RelayerURL string `toml:"relayerUrl"`

	InternalNodeURL    string `toml:"internalNodeUrl"`
	InternalIndexerURL string `toml:"internalIndexerUrl"`
	InternalRelayerURL string `toml:"internalRelayerUrl"`

	WSEnabled bool   `toml:"wsEnabled"`
	WSURL     string `toml:"wsUrl"`

	BlockExplorer *BlockExplorerConfig `toml:"blockExplorer"`

	NativeToken *Currency   `toml:"nativeToken"`
	Currencies  []*Currency `toml:"currencies"`

	NodeGateway *NodeGatewayNetwork `toml:"nodeGateway"`
}

type Currency struct {
	ContractAddress *string       `toml:"contractAddress"`
	Name            string        `toml:"name"`
	Symbol          string        `toml:"symbol"`
	Decimals        uint64        `toml:"decimals"`
	ImageURL        string        `toml:"imageUrl"`
	Native          bool          `toml:"native"`
	Default         bool          `toml:"default"`
	Group           CurrencyGroup `toml:"group"`
}

type BlockExplorerConfig struct {
	Name       string `toml:"name"`
	RootURL    string `toml:"rootUrl"`
	AddressURL string `toml:"addressUrl"`
	TxnHashURL string `toml:"txnHashUrl"`
}

type NodeGatewayNetwork struct {
	Endpoints             []*NodeGatewayEndpoint `toml:"endpoints"`
	Finality              uint64                 `toml:"finality"`
	MonitorDisabled       bool                   `toml:"monitorDisabled"`
	MonitorPollInterval   string                 `toml:"monitorPollInterval"`
	MonitorBlockRetention uint64                 `toml:"monitorBlockRetention"`
	AlertsOff             bool                   `toml:"alertsOff"`
	Important             bool                   `toml:"important"`
}

type NodeGatewayEndpoint struct {
	Label       string   `toml:"label"`
	Priority    int      `toml:"priority"`
	URL         string   `toml:"url"`
	WSURL       string   `toml:"wsurl"`
	SkipMethods []string `toml:"skipMethods"`
	AuxMethodNs []string `toml:"auxMethodNs"`
	Disabled    bool     `toml:"disabled"`
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
			ID:        network.ID,
			URL:       network.NodeURL,
			WSEnabled: network.WSEnabled,
			WSURL:     network.WSURL,
			AuthChain: network.AuthChain,
			Testnet:   network.IsTesnet(),
			Disabled:  network.Disabled,
		}
	}

	return res
}
