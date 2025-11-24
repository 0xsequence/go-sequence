package networks

type Network struct {
	ChainID       uint64               `json:"chainId,omitempty"`
	Name          string               `json:"name,omitempty"`
	Title         string               `json:"title,omitempty"`
	Type          NetworkType          `json:"type,omitempty"`
	LogoURL       string               `json:"logoUrl,omitempty"`
	ENSAddress    string               `json:"ensAddress,omitempty"`
	Deprecated    bool                 `json:"deprecated,omitempty"`
	BlockExplorer *BlockExplorerConfig `json:"blockExplorer,omitempty"`
	NativeToken   *Currency            `json:"nativeToken,omitempty"`
	Currencies    []*Currency          `json:"currencies,omitempty"`
}

type BlockExplorerConfig struct {
	Name       string `json:"name,omitempty"`
	RootURL    string `json:"rootUrl,omitempty"`
	AddressURL string `json:"addressUrl,omitempty"`
	TxnHashURL string `json:"txnHashUrl,omitempty"`
}

type Currency struct {
	ContractAddress *string       `json:"contractAddress,omitempty"`
	Name            string        `json:"name,omitempty"`
	Symbol          string        `json:"symbol,omitempty"`
	Decimals        uint64        `json:"decimals,omitempty"`
	ImageURL        string        `json:"imageUrl,omitempty"`
	Native          bool          `json:"native,omitempty"`
	Default         bool          `json:"default,omitempty"`
	Group           CurrencyGroup `json:"group,omitempty"`
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

func (n *Network) IsTestnet() bool {
	return n.Type == NetworkTypeTestnet
}
