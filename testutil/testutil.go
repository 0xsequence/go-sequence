package testutil

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethmonitor"
	"github.com/0xsequence/ethkit/ethreceipts"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	"github.com/0xsequence/go-sequence/deployer"
	"github.com/0xsequence/go-sequence/relayer"
)

type TestChain struct {
	options TestChainOptions

	chainID        *big.Int         // chainID determined by the test chain
	walletMnemonic string           // test wallet mnemonic parsed from package.json
	Provider       *ethrpc.Provider // provider rpc to the test chain

	Monitor          *ethmonitor.Monitor
	ReceiptsListener *ethreceipts.ReceiptsListener

	RpcRelayer *relayer.RpcRelayer // helper to track RpcRelayer client

	runCtx     context.Context
	runCtxStop context.CancelFunc
}

type TestChainOptions struct {
	NodeURL string
}

var DefaultTestChainOptions = TestChainOptions{
	NodeURL: "http://localhost:8545",
}

func NewTestChain(logger *slog.Logger, opts ...TestChainOptions) (*TestChain, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}

	var err error
	tc := &TestChain{}

	// set options
	if len(opts) > 0 {
		tc.options = opts[0]
	} else {
		tc.options = DefaultTestChainOptions
	}

	// provider
	tc.Provider, err = ethrpc.NewProvider(tc.options.NodeURL)
	if err != nil {
		return nil, err
	}

	// monitor
	var monitor *ethmonitor.Monitor
	monitorOptions := ethmonitor.DefaultOptions
	monitorOptions.Logger = logger
	monitorOptions.StartBlockNumber = nil                                   // track the head
	monitorOptions.PollingInterval = time.Duration(1000 * time.Millisecond) // default poll for new block once per second
	monitorOptions.BlockRetentionLimit = 400                                // keep high number of blocks to query history
	monitorOptions.WithLogs = true                                          // receipt listener needs logs from monitor

	monitor, err = ethmonitor.NewMonitor(tc.Provider, monitorOptions)
	if err != nil {
		return nil, err
	}
	tc.Monitor = monitor

	// receipts listener
	receiptsOptions := ethreceipts.DefaultOptions
	receiptsOptions.NumBlocksToFinality = 10
	receiptsOptions.FilterMaxWaitNumBlocks = 15

	receipts, err := ethreceipts.NewReceiptsListener(logger, tc.Provider, monitor, receiptsOptions)
	if err != nil {
		return nil, err
	}
	tc.ReceiptsListener = receipts

	return tc, nil
}

func (c *TestChain) Connect() error {
	c.runCtx, c.runCtxStop = context.WithCancel(context.Background())

	// connect to the testchain or error out if fail to communicate
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
		return fmt.Errorf("testutil: unable to connect to testchain")
	}

	// start the monitor
	go func() {
		err := c.Monitor.Run(c.runCtx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// start the receipts listener
	go func() {
		err := c.ReceiptsListener.Run(c.runCtx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (c *TestChain) Disconnect() {
	if c.runCtxStop == nil {
		return
	}
	c.runCtxStop()
	c.runCtx = nil
}

func (c *TestChain) ChainID() *big.Int {
	if c.chainID == nil {
		return big.NewInt(0)
	} else {
		return c.chainID
	}
}

func (c *TestChain) SequenceContext() sequence.WalletContext {
	return V2SequenceContext()
}

func (c *TestChain) V1SequenceContext() sequence.WalletContext {
	return V1SequenceContext()
}

func (c *TestChain) V2SequenceContext() sequence.WalletContext {
	return V2SequenceContext()
}

func (c *TestChain) SetRpcRelayer(relayerURL string) error {
	rpcRelayer, err := relayer.NewRpcRelayer(relayerURL, "", c.Provider, c.ReceiptsListener)
	if err != nil {
		return err
	}
	c.RpcRelayer = rpcRelayer
	return nil
}

func (c *TestChain) Wallet() (*ethwallet.Wallet, error) {
	var err error

	if c.walletMnemonic == "" {
		c.walletMnemonic, err = parseTestWalletMnemonic()
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

	err = c.MustFundAddress(wallet.Address())
	if err != nil {
		return nil, err
	}

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

	err = c.MustFundAddress(wallet.Address())
	if err != nil {
		panic(err)
	}

	return wallet
}

func (c *TestChain) MustFundAddress(addr common.Address) error {
	min := fromEther(big.NewInt(1))
	target := fromEther(big.NewInt(100))

	balance, err := c.Provider.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if balance.Cmp(min) != -1 {
		return nil
	}

	var accounts []common.Address
	rpcCall := ethrpc.NewCallBuilder[[]common.Address]("eth_accounts", nil)
	_, err = c.Provider.Do(context.Background(), rpcCall.Into(&accounts))
	if err != nil {
		return err
	}

	type SendTx struct {
		From  *common.Address `json:"from"`
		To    *common.Address `json:"to"`
		Gas   string          `json:"gas"`
		Value string          `json:"value"`
	}

	diff := big.NewInt(0)
	diff.Sub(target, balance)

	tx := &SendTx{
		To:   &addr,
		From: &accounts[0],
		// default not always sufficient to send ETH to a contract
		Gas:   "0x" + big.NewInt(1000000).Text(16),
		Value: "0x" + diff.Text(16),
	}

	// err = c.Provider.RPC.CallContext(context.Background(), nil, "eth_sendTransaction", tx)
	_, err = c.Provider.Do(context.Background(), ethrpc.NewCall("eth_sendTransaction", tx))
	if err != nil {
		return err
	}

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		balance, err = c.Provider.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			return err
		}
		if balance.Cmp(min) != -1 {
			return nil
		}
	}

	return fmt.Errorf("test wallet has no balance")
}

func (c *TestChain) GetDeployWallet() *ethwallet.Wallet {
	return c.MustWallet(5)
}

// GetDeployTransactor returns a account transactor typically used for deploying contracts
func (c *TestChain) GetDeployTransactor() (*bind.TransactOpts, error) {
	return c.GetDeployWallet().Transactor(context.Background())
}

// GetRelayerWallet is the wallet dedicated EOA wallet to relaying transactions
func (c *TestChain) GetRelayerWallet() *ethwallet.Wallet {
	return c.MustWallet(6)
}

func (c *TestChain) V1DeploySequenceContext() (sequence.WalletContext, error) {
	ud, err := deployer.NewUniversalDeployer(c.GetDeployWallet())
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	ctx := context.Background()

	walletFactoryAddress, err := ud.Deploy(ctx, contracts.V1.WalletFactory.ABI, contracts.V1.WalletFactory.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	mainModuleAddress, err := ud.Deploy(ctx, contracts.V1.WalletMainModule.ABI, contracts.V1.WalletMainModule.Bin, 0, nil, 7000000, walletFactoryAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	mainModuleUpgradableAddress, err := ud.Deploy(ctx, contracts.V1.WalletMainModuleUpgradable.ABI, contracts.V1.WalletMainModuleUpgradable.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	guestModuleAddress, err := ud.Deploy(ctx, contracts.V1.WalletGuestModule.ABI, contracts.V1.WalletGuestModule.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	utilsAddress, err := ud.Deploy(ctx, contracts.V1.WalletUtils.ABI, contracts.V1.WalletUtils.Bin, 0, nil, 7000000, walletFactoryAddress, mainModuleAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V1DeploySequenceContext: %w", err)
	}

	return sequence.WalletContext{
		FactoryAddress:              walletFactoryAddress,
		MainModuleAddress:           mainModuleAddress,
		MainModuleUpgradableAddress: mainModuleUpgradableAddress,
		GuestModuleAddress:          guestModuleAddress,
		UtilsAddress:                utilsAddress,
	}, nil
}

func (c *TestChain) V2DeploySequenceContext() (sequence.WalletContext, error) {
	ud, err := deployer.NewEIP2740Deployer(c.GetDeployWallet())
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	ctx := context.Background()

	walletFactoryAddress, err := ud.Deploy(ctx, contracts.V2.WalletFactory.ABI, contracts.V2.WalletFactory.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	mainModuleUpgradableAddress, err := ud.Deploy(ctx, contracts.V2.WalletMainModuleUpgradable.ABI, contracts.V2.WalletMainModuleUpgradable.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	mainModuleAddress, err := ud.Deploy(ctx, contracts.V2.WalletMainModule.ABI, contracts.V2.WalletMainModule.Bin, 0, nil, 7000000, walletFactoryAddress, mainModuleUpgradableAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	guestModuleAddress, err := ud.Deploy(ctx, contracts.V2.WalletGuestModule.ABI, contracts.V2.WalletGuestModule.Bin, 0, nil, 7000000)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	utilsAddress, err := ud.Deploy(ctx, contracts.V2.WalletUtils.ABI, contracts.V2.WalletUtils.Bin, 0, nil, 7000000, walletFactoryAddress, mainModuleAddress)
	if err != nil {
		return sequence.WalletContext{}, fmt.Errorf("testutil, V2DeploySequenceContext: %w", err)
	}

	return sequence.WalletContext{
		FactoryAddress:              walletFactoryAddress,
		MainModuleAddress:           mainModuleAddress,
		MainModuleUpgradableAddress: mainModuleUpgradableAddress,
		GuestModuleAddress:          guestModuleAddress,
		UtilsAddress:                utilsAddress,
	}, nil
}

func (c *TestChain) MustDeploySequenceContext() sequence.WalletContext {
	return c.MustDeploySequenceContexts()[1] // return the v1 context
}

func (c *TestChain) MustDeploySequenceContexts() map[uint8]sequence.WalletContext {
	sc, err := c.V1DeploySequenceContext()
	if err != nil {
		panic(err)
	}

	sc2, err := c.V2DeploySequenceContext()
	if err != nil {
		panic(err)
	}

	if sc != sequenceContext {
		panic("MustDeploySequenceContext failed, deployed context does not match testutil.sequenceContext")
	}

	if sc2 != sequenceContextV2 {
		panic("MustDeploySequenceContext failed, deployed context does not match testutil.sequenceContextV2")
	}

	return map[uint8]sequence.WalletContext{
		1: sc,
		2: sc2,
	}
}

// UniDeploy will deploy a contract registered in `Contracts` registry using the universal deployer. Multiple calls to UniDeploy
// will instantiate just a single instance for the same contract with the same `contractInstanceNum`.
func (c *TestChain) UniDeploy(t *testing.T, contractName string, contractInstanceNum uint, contractConstructorArgs ...interface{}) *ethcontract.Contract {
	artifact, ok := Contracts.Get(contractName)
	if !ok {
		t.Fatal(fmt.Errorf("contract abi not found for name %s", contractName))
	}
	ud, err := deployer.NewUniversalDeployer(c.GetDeployWallet())
	if err != nil {
		t.Fatal(err)
	}
	contractAddress, err := ud.Deploy(context.Background(), artifact.ABI, artifact.Bin, contractInstanceNum, nil, 5000000, contractConstructorArgs...)
	if err != nil {
		t.Fatal(err)
	}
	return ethcontract.NewContractCaller(contractAddress, artifact.ABI, c.Provider)
}

// Deploy will deploy a contract registered in `Contracts` registry using the standard deployment method. Each Deploy call
// will instanitate a new contract on the test chain.
func (c *TestChain) Deploy(t *testing.T, contractName string, contractConstructorArgs ...interface{}) (*ethcontract.Contract, *types.Receipt) {
	artifact, ok := Contracts.Get(contractName)
	if !ok {
		t.Fatal(fmt.Errorf("contract abi not found for name %s", contractName))
	}

	data := make([]byte, len(artifact.Bin))
	copy(data, artifact.Bin)

	var input []byte
	var err error

	// encode constructor call
	if len(contractConstructorArgs) > 0 && len(artifact.ABI.Constructor.Inputs) > 0 {
		input, err = artifact.ABI.Pack("", contractConstructorArgs...)
	} else {
		input, err = artifact.ABI.Pack("")
	}
	if err != nil {
		t.Fatal(fmt.Errorf("contract constructor pack failed: %w", err))
	}

	// append constructor calldata at end of the contract bin
	data = append(data, input...)

	wallet := c.GetDeployWallet()
	signedTxn, err := wallet.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
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

	return ethcontract.NewContractCaller(receipt.ContractAddress, artifact.ABI, c.Provider), receipt
}

func (c *TestChain) WaitMined(txn common.Hash) error {
	_, err := ethrpc.WaitForTxnReceipt(context.Background(), c.Provider, txn)
	return err
}

func (c *TestChain) V1DummySequenceWallets(nWallets uint64, startingSeed uint64) ([]*sequence.Wallet[core.WalletConfig], error) {
	var wallets []*sequence.Wallet[core.WalletConfig]

	for i := uint64(0); i < nWallets; i++ {
		wallet, err := c.V1DummySequenceWallet(startingSeed + i*1000)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func (c *TestChain) V2DummySequenceWallets(nWallets uint64, startingSeed uint64) ([]*sequence.Wallet[core.WalletConfig], error) {
	var wallets []*sequence.Wallet[core.WalletConfig]

	for i := uint64(0); i < nWallets; i++ {
		wallet, err := c.V2DummySequenceWallet(startingSeed + i*1000)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

func (c *TestChain) RandomNonce() *big.Int {
	space := big.NewInt(int64(time.Now().Nanosecond()))

	encoded, err := sequence.EncodeNonce(space, big.NewInt(0))
	if err != nil {
		panic(err)
	}

	return encoded
}

func (c *TestChain) V1DummySequenceWallet(seed uint64, optSkipDeploy ...bool) (*sequence.Wallet[core.WalletConfig], error) {
	// Generate a single-owner sequence wallet based on a private key generated from seed above
	owner, err := ethwallet.NewWalletFromPrivateKey(DummyPrivateKey(seed))
	if err != nil {
		return nil, err
	}
	wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](owner, V1SequenceContext())
	if err != nil {
		return nil, err
	}

	// Set provider on sequence wallet
	err = wallet.SetProvider(c.Provider)
	if err != nil {
		return nil, err
	}

	// Set relayer on sequence wallet, which is used when the wallet sends transactions
	localRelayer, err := relayer.NewLocalRelayer(c.GetRelayerWallet(), c.ReceiptsListener)
	if err != nil {
		return nil, err
	}
	err = wallet.SetRelayer(localRelayer)
	if err != nil {
		return nil, err
	}

	genericWallet := sequence.GenericNewWalletWithCoreWalletConfig[*v1.WalletConfig](wallet)

	// Skip deploying the dummy wallet if specified
	if len(optSkipDeploy) > 0 && optSkipDeploy[0] {
		return genericWallet, nil
	}

	err = c.DeploySequenceWallet(genericWallet)
	if err != nil {
		return nil, err
	}

	return genericWallet, nil
}

func (c *TestChain) V2DummySequenceWallet(seed uint64, optSkipDeploy ...bool) (*sequence.Wallet[core.WalletConfig], error) {
	// Generate a single-owner sequence wallet based on a private key generated from seed above
	owner, err := ethwallet.NewWalletFromPrivateKey(DummyPrivateKey(seed))
	if err != nil {
		return nil, err
	}
	wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](owner, V2SequenceContext())
	if err != nil {
		return nil, err
	}

	// Set provider on sequence wallet
	err = wallet.SetProvider(c.Provider)
	if err != nil {
		return nil, err
	}

	// Set relayer on sequence wallet, which is used when the wallet sends transactions
	localRelayer, err := relayer.NewLocalRelayer(c.GetRelayerWallet(), c.ReceiptsListener)
	if err != nil {
		return nil, err
	}
	err = wallet.SetRelayer(localRelayer)
	if err != nil {
		return nil, err
	}

	genericWallet := sequence.GenericNewWalletWithCoreWalletConfig[*v2.WalletConfig](wallet)

	// Skip deploying the dummy wallet if specified
	if len(optSkipDeploy) > 0 && optSkipDeploy[0] {
		return genericWallet, nil
	}

	err = c.DeploySequenceWallet(genericWallet)
	if err != nil {
		return nil, err
	}

	return genericWallet, nil
}

func (c *TestChain) DeploySequenceWallet(wallet *sequence.Wallet[core.WalletConfig]) error {
	// Check if wallet is already deployed, in which case we will return right away
	ok, _ := wallet.IsDeployed()
	if ok {
		return nil
	}

	// Deploy the wallet, via our account designated for relaying txs, but it could be any with some ETH
	sender := c.GetRelayerWallet()
	_, _, waitReceipt, err := sequence.DeploySequenceWallet(sender, wallet.GetWalletConfig(), wallet.GetWalletContext())
	if err != nil {
		return err
	}
	_, err = waitReceipt(context.Background())
	if err != nil {
		return err
	}

	// Ensure deployment worked
	ok, err = wallet.IsDeployed()
	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("dummy sequence wallet failed to deploy")
	}

	return nil
}
