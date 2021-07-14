package testutil

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
)

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

func SignAndSend(t *testing.T, wallet *sequence.Wallet, to common.Address, data []byte) error {
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

func SignAndSendRawTransaction(t *testing.T, wallet *sequence.Wallet, stx *sequence.Transaction) error {
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
	// NOTE: if you start test chain with `make start-test-chain-verbose`, you will see the metaTxnID above
	// correctly logged..

	return err
}

func BatchSignAndSend(t *testing.T, wallet *sequence.Wallet, to common.Address, data [][]byte) error {
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
	// NOTE: if you start test chain with `make start-test-chain-verbose`, you will see the metaTxnID above
	// correctly logged..

	return err
}

// RandomSeed will generate a random seed
func RandomSeed() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}
