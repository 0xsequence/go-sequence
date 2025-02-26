package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	"github.com/stretchr/testify/assert"
)

var sequenceContext = sequence.WalletContext{
	FactoryAddress:              common.HexToAddress("0xf9D09D634Fb818b05149329C1dcCFAeA53639d96"),
	MainModuleAddress:           common.HexToAddress("0xd01F11855bCcb95f88D7A48492F66410d4637313"),
	MainModuleUpgradableAddress: common.HexToAddress("0x7EFE6cE415956c5f80C6530cC6cc81b4808F6118"),
	GuestModuleAddress:          common.HexToAddress("0x02390F3E6E5FD1C6786CB78FD3027C117a9955A7"),
	UtilsAddress:                common.HexToAddress("0xd130B43062D875a4B7aF3f8fc036Bc6e9D3E1B3E"),
	CreationCode:                hexutil.Encode(contracts.V1.CreationCode),
}

var sequenceContextV2 = sequence.WalletContext{
	FactoryAddress:              common.HexToAddress("0xFaA5c0b14d1bED5C888Ca655B9a8A5911F78eF4A"),
	MainModuleAddress:           common.HexToAddress("0xfBf8f1A5E00034762D928f46d438B947f5d4065d"),
	MainModuleUpgradableAddress: common.HexToAddress("0x4222dcA3974E39A8b41c411FeDDE9b09Ae14b911"),
	GuestModuleAddress:          common.HexToAddress("0xfea230Ee243f88BC698dD8f1aE93F8301B6cdfaE"),
	UtilsAddress:                common.HexToAddress("0xdbbFa3cB3B087B64F4ef5E3D20Dda2488AA244e6"),
	CreationCode:                hexutil.Encode(contracts.V2.CreationCode),
}

var sequenceContextV3 = sequence.WalletContext{
	FactoryAddress:              common.HexToAddress("0x6843d600C5fF98E75DF4e7b236D9513eD54A5344"),
	MainModuleAddress:           common.HexToAddress("0x429937a2aCB7789c5D29Cd3325cC77f02E21539d"),
	MainModuleUpgradableAddress: common.HexToAddress("0x429937a2aCB7789c5D29Cd3325cC77f02E21539d"),
	GuestModuleAddress:          common.HexToAddress("0xC79f7f217D753222EFb5c6CA698A1823c3F15EA5"),
	UtilsAddress:                common.HexToAddress("0x429937a2aCB7789c5D29Cd3325cC77f02E21539d"),
	CreationCode:                hexutil.Encode(contracts.V3.CreationCode),
}

func SequenceContext() sequence.WalletContext {
	return sequenceContextV2
}

func V1SequenceContext() sequence.WalletContext {
	return sequenceContext
}

func V2SequenceContext() sequence.WalletContext {
	return sequenceContextV2
}

func V3SequenceContext() sequence.WalletContext {
	return sequenceContextV3
}

func SequenceContexts() map[uint8]sequence.WalletContext {
	return map[uint8]sequence.WalletContext{
		1: sequenceContext,
		2: sequenceContextV2,
		3: sequenceContextV3,
	}
}

// parseTestWalletMnemonic parses the wallet mnemonic from ./package.json, the same
// key used to start the test chain server.
func parseTestWalletMnemonic() (string, error) {
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

// DummyAddr returns a dummy address
func DummyAddr() common.Address {
	addr, _ := ethwallet.NewWalletFromRandomEntropy()
	return addr.Address()
}

// DummyPrivateKey returns random private key in hex used with ethwallet
func DummyPrivateKey(seed uint64) string {
	return fmt.Sprintf("%064x", seed)
}

func SignAndSend[C core.WalletConfig](t *testing.T, wallet *sequence.Wallet[C], to common.Address, data []byte) error {
	stx := &sequence.Transaction{
		// DelegateCall:  false,
		// RevertOnError: false,
		// GasLimit: big.NewInt(800000),
		// Value:         big.NewInt(0),
		To:   to,
		Data: data,
	}

	return SignAndSendRawTransaction(t, wallet, stx)
}

func SignAndSendRawTransaction[C core.WalletConfig](t *testing.T, wallet *sequence.Wallet[C], stx *sequence.Transaction) error {
	// Now, we must sign the meta txn
	signedTx, err := wallet.SignTransaction(context.Background(), stx)
	assert.NoError(t, err)

	metaTxnID, _, waitReceipt, err := wallet.SendTransaction(context.Background(), signedTx)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)

	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)
	assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

	// TODO: decode the receipt, and lets confirm we have the metaTxnID event in there..
	// NOTE: if you start test chain with `make start-testchain-verbose`, you will see the metaTxnID above
	// correctly logged..

	return err
}

func BatchSignAndSend[C core.WalletConfig](t *testing.T, wallet *sequence.Wallet[C], to common.Address, data [][]byte) error {
	var stxs []*sequence.Transaction
	for i := 0; i < len(data); i++ {
		stxs = append(stxs, &sequence.Transaction{
			// DelegateCall:  false,
			// RevertOnError: false,
			// GasLimit: big.NewInt(800000),
			// Value:         big.NewInt(0),
			To:   to,
			Data: data[i],
		})
	}

	// Now, we must sign the meta txn
	signedTx, err := wallet.SignTransactions(context.Background(), stxs)
	assert.NoError(t, err)

	metaTxnID, tx, waitReceipt, err := wallet.SendTransactions(context.Background(), signedTx)
	assert.NoError(t, err)
	assert.NotEmpty(t, metaTxnID)
	assert.NotNil(t, tx)

	receipt, err := waitReceipt(context.Background())
	assert.NoError(t, err)
	assert.True(t, receipt.Status == types.ReceiptStatusSuccessful)

	// TODO: decode the receipt, and lets confirm we have the metaTxnID event in there..
	// NOTE: if you start test chain with `make start-testchain-verbose`, you will see the metaTxnID above
	// correctly logged..

	return err
}

// RandomSeed will generate a random seed
func RandomSeed() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}

func fromEther(ether *big.Int) *big.Int {
	oneEth := big.NewInt(10)
	oneEth.Exp(oneEth, big.NewInt(18), nil)

	return ether.Mul(ether, oneEth)
}
