package sequence_test

import (
	"context"
	"math/big"
	"sort"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	"github.com/stretchr/testify/assert"
)

func TestWalletAddress(t *testing.T) {
	eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
	assert.NoError(t, err)
	assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

	w, err := sequence.NewWalletSingleOwner(eoa, sequence.WalletContext{
		FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
	})

	assert.NoError(t, err)
	assert.Equal(t, "0xcBf920895f0A101b85f94903775E61d1B652b6Ca", w.Address().Hex())
}

func TestWalletSignMessage(t *testing.T) {
	// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
	eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
	assert.NoError(t, err)
	assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

	wallet, err := sequence.NewWalletSingleOwner(eoa, sequence.WalletContext{
		FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
		MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "0xF0BA65550F2d1DCCf4B131B774844DC3d801D886", wallet.Address().Hex())

	// we set chainID here for debugging purposes, but in general should not set it manually
	wallet.SetChainID(big.NewInt(1))

	message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

	sig, _, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
	assert.NoError(t, err)

	expectedSig := "0x00010001a0fb306480bc3027c04d33a16370f4618b29f2d5b89464f526045c94802bc9d1525389c364b75daf58e859ed0d6105aac6b3718e4659814c7793c626653edb871b02"
	assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))
}

func TestWalletSignMessageAndValidate(t *testing.T) {
	// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
	eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
	assert.NoError(t, err)
	assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

	wallet, err := sequence.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)
	assert.Equal(t, "0xdb6a2e3368C302723B71A1Cd154510a11952998E", wallet.Address().Hex())

	wallet.SetProvider(testChain.Provider)

	message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

	sig, _, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
	assert.NoError(t, err)

	expectedSig := "0x00010001aea46cf46662768aa0287f04788afeaf84ed12eda6dd759bdb500db113b289b20cbc85a8dbb040e14082b401b081e33f16a9685993b168212a1a0f4e686cb8b31c02"
	assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))

	isValidSig, err := wallet.IsValidSignature(sequence.MessageDigest(ethcoder.MustHexDecode(message)), sig)
	assert.NoError(t, err)
	assert.True(t, isValidSig)
}

func TestWalletSignAndRecoverConfig(t *testing.T) {
	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	wallet, err := sequence.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)

	wallet.SetChainID(big.NewInt(3))

	message := "Hi! this is a test message"
	sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
	assert.NoError(t, err)

	subDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), common.BytesToHash(ethcoder.Keccak256([]byte(message))))
	assert.NoError(t, err)

	recoveredWalletConfig, weight, err := sequence.RecoverWalletConfigFromDigest(subDigest, sig, wallet.GetWalletContext(), wallet.GetChainID(), testChain.Provider)
	assert.NoError(t, err)

	assert.Equal(t, uint16(1), recoveredWalletConfig.Threshold_)
	assert.GreaterOrEqual(t, weight, uint(recoveredWalletConfig.Threshold_))
	assert.Len(t, recoveredWalletConfig.Signers_, 1)

	address, err := sequence.AddressFromWalletConfig(recoveredWalletConfig, wallet.GetWalletContext())
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), address)
}

func TestWalletSignAndRecoverConfigOfMultipleSignersV1(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := &v1.WalletConfig{
		Threshold_: 3,
		Signers_: v1.WalletConfigSigners{
			{Weight: 2, Address: eoa1.Address()},
			{Weight: 5, Address: eoa2.Address()},
		},
	}

	sequence.SortWalletConfig(walletConfig)

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config: walletConfig,
	}, eoa1)
	assert.NoError(t, err)

	wallet.SetChainID(big.NewInt(3))

	message := "Hi! this is a test message"
	sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
	assert.NoError(t, err)

	subDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), common.BytesToHash(ethcoder.Keccak256([]byte(message))))
	assert.NoError(t, err)

	recoveredWalletConfig, weight, err := sequence.RecoverWalletConfigFromDigest(subDigest, sig, wallet.GetWalletContext(), wallet.GetChainID(), testChain.Provider)
	assert.NoError(t, err)

	assert.Equal(t, uint16(3), recoveredWalletConfig.Threshold_)
	assert.Equal(t, uint(2), weight)
	assert.Len(t, recoveredWalletConfig.Signers_, 2)

	address, err := sequence.AddressFromWalletConfig(walletConfig, wallet.GetWalletContext())
	assert.NoError(t, err)
	assert.Equal(t, wallet.Address(), address)
}

func TestShouldNotSideEffectWalletConfigV1(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	// Inverse-sort signers
	signers := v1.WalletConfigSigners{
		{Weight: 2, Address: eoa1.Address()},
		{Weight: 5, Address: eoa2.Address()},
	}

	sort.Sort(sort.Reverse(signers))

	walletConfig := &v1.WalletConfig{
		Threshold_: 3,
		Signers_:   signers,
	}

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config: walletConfig,
	}, eoa1)
	assert.NoError(t, err)

	// Wallet config should be sorted
	// but signers should not be affected
	walletFromGet := wallet.GetWalletConfig().(*v1.WalletConfig)
	assert.Equal(t, signers[0].Address, walletFromGet.Signers_[1].Address)
	assert.Equal(t, signers[1].Address, walletFromGet.Signers_[0].Address)
}

func TestShouldIgnoreConfigSortV1(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	// Inverse-sort signers
	signers := v1.WalletConfigSigners{
		{Weight: 2, Address: eoa1.Address()},
		{Weight: 5, Address: eoa2.Address()},
	}

	sort.Sort(sort.Reverse(signers))

	walletConfig := &v1.WalletConfig{
		Threshold_: 3,
		Signers_:   signers,
	}

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config:          walletConfig,
		SkipSortSigners: true,
	}, eoa1)
	assert.NoError(t, err)

	// Wallet config should remain the same
	walletFromGet := wallet.GetWalletConfig().(*v1.WalletConfig)
	assert.Equal(t, signers[0].Address, walletFromGet.Signers_[0].Address)
	assert.Equal(t, signers[1].Address, walletFromGet.Signers_[1].Address)
}

func TestWalletWithNonDeterministicConfigV1(t *testing.T) {
	rw, _ := ethwallet.NewWalletFromRandomEntropy()
	randomAddr := rw.Address()

	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := &v1.WalletConfig{
		Threshold_: 3,
		Signers_: v1.WalletConfigSigners{
			{Weight: 2, Address: eoa1.Address()},
			{Weight: 5, Address: eoa2.Address()},
		},
	}

	sequence.SortWalletConfig(walletConfig)

	wallet, err := sequence.NewWallet(sequence.WalletOptions{
		Config:          walletConfig,
		Address:         randomAddr,
		SkipSortSigners: true,
	}, eoa1)
	assert.NoError(t, err)

	assert.Equal(t, wallet.Address(), randomAddr)
	assert.Equal(t, wallet.GetWalletConfig(), walletConfig)

	// Should maintain properties after updating config
	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	newConfig := &v1.WalletConfig{
		Threshold_: 3,
		Signers_: v1.WalletConfigSigners{
			{Weight: 2, Address: eoa3.Address()},
			{Weight: 5, Address: eoa2.Address()},
		},
	}

	// Use config with wallet
	nwallet, err := wallet.UseConfig(newConfig)
	assert.NoError(t, err)

	assert.Equal(t, nwallet.Address(), randomAddr)
	assert.Equal(t, nwallet.GetWalletConfig(), newConfig)

	// Should maintain properties after updating signers
	eoa4, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	nwallet, err = nwallet.UseSigners(eoa4)
	assert.NoError(t, err)

	assert.Equal(t, nwallet.Address(), randomAddr)
	assert.Equal(t, nwallet.GetWalletConfig(), newConfig)
}
