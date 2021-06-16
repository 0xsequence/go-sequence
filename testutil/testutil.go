package testutil

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contract"
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

func (c *TestChain) GetDeployWallet() *ethwallet.Wallet {
	return c.MustWallet(5)
}

// GetDeployTransactor returns a account transactor typically used for deploying contracts
func (c *TestChain) GetDeployTransactor() *bind.TransactOpts {
	return c.GetDeployWallet().Transactor()
}

func (c *TestChain) DeploySequenceContext() (sequence.WalletContext, error) {
	ud, err := deployer.NewUniversalDeployer(c.GetDeployWallet())
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	ctx := context.Background()

	walletFactoryAddress, err := ud.Deploy(ctx, contract.WalletFactory.ABI, contract.WalletFactory.Bin, 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	mainModuleAddress, err := ud.Deploy(ctx, contract.WalletMainModule.ABI, contract.WalletMainModule.Bin, 0, nil, walletFactoryAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	mainModuleUpgradableAddress, err := ud.Deploy(ctx, contract.WalletMainModuleUpgradable.ABI, contract.WalletMainModuleUpgradable.Bin, 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	guestModuleAddress, err := ud.Deploy(ctx, contract.WalletGuestModule.ABI, contract.WalletGuestModule.Bin, 0, nil)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, DeploySequenceContext: %w", err)
	}

	utilsAddress, err := ud.Deploy(ctx, contract.WalletUtils.ABI, contract.WalletUtils.Bin, 0, nil, walletFactoryAddress, mainModuleAddress)
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

// UniDeploy will deploy a contract registered in `Contracts` registry using the universal deployer. Multiple calls to UniDeploy
// will instantiate just a single instance for the same contract with the same `contractInstanceNum`.
func (c *TestChain) UniDeploy(t *testing.T, contractName string, contractInstanceNum uint, contractConstructorArgs ...interface{}) (common.Address, *contract.ContractABI) {
	contractABI := Contracts.Get(contractName)
	if contractABI == nil {
		t.Fatal(fmt.Errorf("contract abi not found for name %s", contractName))
	}
	ud, err := deployer.NewUniversalDeployer(c.GetDeployWallet())
	if err != nil {
		t.Fatal(err)
	}
	contractAddress, err := ud.Deploy(context.Background(), contractABI.ABI, contractABI.Bin, contractInstanceNum, nil, contractConstructorArgs...)
	if err != nil {
		t.Fatal(err)
	}
	return contractAddress, contractABI
}

// Deploy will deploy a contract registered in `Contracts` registry using the standard deployment method. Each Deploy call
// will instanitate a new contract on the test chain.
func (c *TestChain) Deploy(t *testing.T, contractName string, contractConstructorArgs ...interface{}) (*types.Receipt, *contract.ContractABI) {
	contractABI := Contracts.Get(contractName)
	if contractABI == nil {
		t.Fatal(fmt.Errorf("contract abi not found for name %s", contractName))
	}

	data := make([]byte, len(contractABI.Bin))
	copy(data, contractABI.Bin)

	var input []byte
	var err error

	// encode constructor call
	if len(contractConstructorArgs) > 0 && len(contractABI.Constructor.Inputs) > 0 {
		input, err = contractABI.Pack("", contractConstructorArgs...)
	} else {
		input, err = contractABI.Pack("")
	}
	if err != nil {
		t.Fatal(fmt.Errorf("contract constructor pack failed: %w", err))
	}

	// append constructor calldata at end of the contract bin
	data = append(data, input...)

	wallet := c.GetDeployWallet()
	signedTxn, _, err := wallet.NewTransaction(context.Background(), &ethwallet.TransactionRequest{
		Data: data,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, waitTx, err := wallet.SendTransaction(context.Background(), signedTxn)
	if err != nil {
		t.Fatal(err)
	}
	receipt, err := waitTx(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		t.Fatal(fmt.Errorf("txn failed: %w", err))
	}

	return receipt, contractABI
}

func (c *TestChain) AssertCodeAt(t *testing.T, address common.Address) {
	code, err := c.Provider.CodeAt(context.Background(), address, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(code) == 0 {
		t.Fatal(fmt.Errorf("no contract code at %s", address.Hex()))
	}
}

func (c *TestChain) WaitMined(txn common.Hash) error {
	_, err := ethrpc.WaitForTxnReceipt(context.Background(), c.Provider, txn)
	return err
}
