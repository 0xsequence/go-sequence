package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/deployer"
)

type TestChain struct {
	options TestChainOptions

	chainID        *big.Int         // chainID determined by the test chain
	walletMnemonic string           // test wallet mnemonic parsed from package.json
	Provider       *ethrpc.Provider // provider rpc to the test chain
}

type TestChainOptions struct {
	NodeURL string
}

var DefaultTestChainOptions = TestChainOptions{
	NodeURL: "http://localhost:8545",
}

func NewTestChain(opts ...TestChainOptions) (*TestChain, error) {
	var err error
	tc := &TestChain{}

	// set options
	if opts != nil && len(opts) > 0 {
		tc.options = opts[0]
	} else {
		tc.options = DefaultTestChainOptions
	}

	// provider
	tc.Provider, err = ethrpc.NewProvider(tc.options.NodeURL)
	if err != nil {
		return nil, err
	}

	// connect to the test-chain or error out if fail to communicate
	if err := tc.connect(); err != nil {
		return nil, err
	}

	return tc, nil
}

func (c *TestChain) connect() error {
	numAttempts := 6
	for i := 0; i < numAttempts; i++ {
		chainID, err := c.Provider.ChainID(context.Background())
		if err != nil || chainID == nil {
			time.Sleep(1 * time.Second)
			continue
		}
		c.chainID = chainID
	}
	if c.chainID == nil {
		return fmt.Errorf("testutil: unable to connect to test-chain")
	}
	return nil
}

func (c *TestChain) ChainID() *big.Int {
	return c.chainID
}

func (c *TestChain) Wallet() (*ethwallet.Wallet, error) {
	var err error

	if c.walletMnemonic == "" {
		c.walletMnemonic, err = ParseTestWalletMnemonic()
		if err != nil {
			return nil, err
		}
	}

	// we create a new instance each time so we don't change the account indexes
	// on the wallet across consumers
	wallet, err := ethwallet.NewWalletFromMnemonic(c.walletMnemonic)
	if err != nil {
		return nil, err
	}
	wallet.SetProvider(c.Provider)

	return wallet, nil
}

func (c *TestChain) MustWallet(optAccountIndex ...uint32) *ethwallet.Wallet {
	wallet, err := c.Wallet()
	if err != nil {
		panic(err)
	}
	if len(optAccountIndex) > 0 {
		_, err = wallet.SelfDeriveAccountIndex(optAccountIndex[0])
		if err != nil {
			panic(err)
		}
	}
	return wallet
}

// TODO: add Deploy() and DeployWith() commands ....

func (c *TestChain) DeploySequenceContext() (sequence.WalletContext, error) {
	ud, err := deployer.NewUniversalDeployer(c.MustWallet(5))
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	ctx := context.Background()

	walletFactoryAddress, err := ud.Deploy(ctx, contracts.ABI_WALLET_FACTORY, common.FromHex(contracts.FactoryBin), 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	mainModuleAddress, err := ud.Deploy(ctx, contracts.ABI_WALLET_MAIN_MODULE, common.FromHex(contracts.MainModuleBin), 0, nil, walletFactoryAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	mainModuleUpgradableAddress, err := ud.Deploy(ctx, contracts.ABI_WALLET_MAIN_MODULE_UPGRADABLE, common.FromHex(contracts.MainModuleUpgradableBin), 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	guestModuleAddress, err := ud.Deploy(ctx, contracts.ABI_WALLET_GUEST_MODULE, common.FromHex(contracts.GuestModuleBin), 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	utilsAddress, err := ud.Deploy(ctx, contracts.ABI_WALLET_UTILS, common.FromHex(contracts.SequenceUtilsBin), 0, nil, walletFactoryAddress, mainModuleAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	return sequence.WalletContext{
		FactoryAddress:              walletFactoryAddress,
		MainModuleAddress:           mainModuleAddress,
		MainModuleUpgradableAddress: mainModuleUpgradableAddress,
		GuestModuleAddress:          guestModuleAddress,
		UtilsAddress:                utilsAddress,
	}, nil
}

// TODO: .DeployContract() etc..

// ParseTestWalletMnemonic parses the wallet mnemonic from ./package.json, the same
// key used to start the test chain server.
func ParseTestWalletMnemonic() (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	cwd := filepath.Dir(filename)

	packageJSONFile := filepath.Join(cwd, "./chain/package.json")
	data, err := ioutil.ReadFile(packageJSONFile)
	if err != nil {
		return "", fmt.Errorf("ParseTestWalletMnemonic, read: %w", err)
	}

	var dict struct {
		Config struct {
			Mnemonic string `json:"mnemonic"`
		} `json:"config"`
	}
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return "", fmt.Errorf("ParseTestWalletMnemonic, unmarshal: %w", err)
	}

	return dict.Config.Mnemonic, nil
}

var sequenceContext = sequence.WalletContext{
	FactoryAddress:              common.HexToAddress("0xf9D09D634Fb818b05149329C1dcCFAeA53639d96"),
	MainModuleAddress:           common.HexToAddress("0xd01F11855bCcb95f88D7A48492F66410d4637313"),
	MainModuleUpgradableAddress: common.HexToAddress("0x7EFE6cE415956c5f80C6530cC6cc81b4808F6118"),
	GuestModuleAddress:          common.HexToAddress("0x02390F3E6E5FD1C6786CB78FD3027C117a9955A7"),
	UtilsAddress:                common.HexToAddress("0xd130B43062D875a4B7aF3f8fc036Bc6e9D3E1B3E"),
}

func SequenceContext() sequence.WalletContext {
	return sequenceContext
}
