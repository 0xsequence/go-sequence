package sequence_test

import (
	"context"
	"errors"
	"math/big"
	"sort"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	"github.com/0xsequence/go-sequence/relayer"
	"github.com/0xsequence/go-sequence/signing_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWalletAddress(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
		assert.NoError(t, err)
		assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

		w, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa, sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
			MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
		})

		assert.NoError(t, err)
		assert.Equal(t, "0xcBf920895f0A101b85f94903775E61d1B652b6Ca", w.Address().Hex())
	})

	t.Run("v2", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
		assert.NoError(t, err)
		assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

		w, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa, sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
			MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
		})

		assert.NoError(t, err)
		assert.Equal(t, "0x686F888A26CD92E859E504d14Db1c19c58b28b7f", w.Address().Hex())
	})
}

func TestWalletSignMessage(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa, sequence.WalletContext{
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
	})

	t.Run("v2", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa, sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
		})
		assert.NoError(t, err)
		assert.Equal(t, "0x8AD6cbc016f72971d317f281497aa080DF87a013", wallet.Address().Hex())

		// we set chainID here for debugging purposes, but in general should not set it manually
		wallet.SetChainID(big.NewInt(1))

		message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

		sig, _, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
		assert.NoError(t, err)

		expectedSig := "0x0100010000000000012d1d8f90e8ae7cb4dff2b63643633bb0480ca8ce4c993f3934587e0f2041b4970832b96512cfc5418a125d4110c1e3c8293c515f144d0119b4ab57b2db0c95741c02"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))
	})
}

func TestWalletSignMessageAndValidate(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa)
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
	})

	t.Run("v2", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa)
		assert.NoError(t, err)
		assert.Equal(t, "0xA26bE1afd2517c4Da0FA0464953EC0D43aA0552b", wallet.Address().Hex())

		wallet.SetProvider(testChain.Provider)

		message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

		sig, _, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
		assert.NoError(t, err)

		expectedSig := "0x010001000000000001a91667dc105882d23a555c49e0e76a135c70bf7e4493cc9550a1bff6ffb7525372c7f67e8d6e18ccba35bdabe9264c8241c6a7fec88ccd8b74fcf680f83009541b02"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))

		isValidSig, err := wallet.IsValidSignature(sequence.MessageDigest(ethcoder.MustHexDecode(message)), sig)
		assert.NoError(t, err)
		assert.True(t, isValidSig)
	})
}

func TestWalletDeploy(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		rel, err := relayer.NewLocalRelayer(testChain.MustWallet(0), nil)
		require.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa)
		assert.NoError(t, err)

		err = wallet.SetProvider(testChain.Provider)
		require.NoError(t, err)

		err = wallet.SetRelayer(rel)
		require.NoError(t, err)

		_, _, wait, err := wallet.Deploy(context.Background())
		assert.NoError(t, err)

		_, err = wait(context.Background())
		require.NoError(t, err)

		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})

	t.Run("v2", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		rel, err := relayer.NewLocalRelayer(testChain.MustWallet(0), nil)
		require.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa)
		assert.NoError(t, err)

		err = wallet.SetProvider(testChain.Provider)
		require.NoError(t, err)

		err = wallet.SetRelayer(rel)
		require.NoError(t, err)

		_, _, wait, err := wallet.Deploy(context.Background())
		assert.NoError(t, err)

		_, err = wait(context.Background())
		require.NoError(t, err)

		isDeployed, err := wallet.IsDeployed()
		assert.NoError(t, err)
		assert.True(t, isDeployed)
	})
}

func TestWalletSignAndRecoverConfig(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](eoa)
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		s, err := v1.Core.DecodeSignature(sig)
		assert.NoError(t, err)

		recoveredWalletConfig, weight, err := s.Recover(context.Background(), core.Digest{Hash: common.BytesToHash(ethcoder.Keccak256([]byte(message)))}, wallet.Address(), wallet.GetChainID(), testChain.Provider)
		assert.NoError(t, err)

		assert.Equal(t, uint16(1), recoveredWalletConfig.Threshold_)
		assert.GreaterOrEqual(t, 0, weight.Cmp(big.NewInt(0).SetUint64(uint64(recoveredWalletConfig.Threshold_))))
		assert.Len(t, recoveredWalletConfig.Signers_, 1)

		address, err := sequence.AddressFromWalletConfig(recoveredWalletConfig, wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.Equal(t, wallet.Address(), address)
	})

	t.Run("v2", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](eoa)
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		s, err := v2.Core.DecodeSignature(sig)
		assert.NoError(t, err)

		recoveredWalletConfig, weight, err := s.Recover(context.Background(), core.Digest{Hash: common.BytesToHash(ethcoder.Keccak256([]byte(message)))}, wallet.Address(), wallet.GetChainID(), testChain.Provider)
		assert.NoError(t, err)

		assert.Equal(t, uint16(1), recoveredWalletConfig.Threshold_)
		assert.GreaterOrEqual(t, 0, weight.Cmp(big.NewInt(0).SetUint64(uint64(recoveredWalletConfig.Threshold_))))
		assert.Equal(t, recoveredWalletConfig.Tree.ImageHash(), wallet.GetWalletConfig().Tree.ImageHash())

		address, err := sequence.AddressFromWalletConfig(recoveredWalletConfig, wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.Equal(t, wallet.Address(), address)
	})
}

func TestWalletSignAndRecoverConfigOfMultipleSigners(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
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

		sequence.V1SortWalletConfig(walletConfig)

		wallet, err := sequence.GenericNewWallet[*v1.WalletConfig](sequence.WalletOptions[*v1.WalletConfig]{
			Config: walletConfig,
		}, eoa1)
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		subDigest, err := sequence.SubDigest(wallet.GetChainID(), wallet.Address(), common.BytesToHash(ethcoder.Keccak256([]byte(message))))
		assert.NoError(t, err)

		recoveredWalletConfig, weight, err := sequence.GenericRecoverWalletConfigFromDigest[*v1.WalletConfig](subDigest, sig, wallet.Address(), wallet.GetWalletContext(), wallet.GetChainID(), testChain.Provider)
		assert.NoError(t, err)

		assert.Equal(t, uint16(3), recoveredWalletConfig.Threshold())
		assert.Equal(t, weight.Cmp(big.NewInt(int64(2))), 0)
		assert.Len(t, recoveredWalletConfig.Signers(), 2)

		address, err := sequence.AddressFromWalletConfig(walletConfig, wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.Equal(t, wallet.Address(), address)
	})

	t.Run("v2", func(t *testing.T) {
		eoa1, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa2, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		walletConfig := &v2.WalletConfig{
			Threshold_: 3,
			Tree: &v2.WalletConfigTreeNode{
				Left: &v2.WalletConfigTreeAddressLeaf{
					Weight: 2, Address: eoa1.Address(),
				},
				Right: &v2.WalletConfigTreeAddressLeaf{
					Weight: 5, Address: eoa2.Address(),
				},
			},
		}

		wallet, err := sequence.GenericNewWallet[*v2.WalletConfig](sequence.WalletOptions[*v2.WalletConfig]{
			Config: walletConfig,
		}, eoa1)
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, _, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		recoveredWalletConfig, weight, err := sequence.GenericRecoverWalletConfigFromDigest[*v2.WalletConfig](ethcoder.Keccak256([]byte(message)), sig, wallet.Address(), wallet.GetWalletContext(), wallet.GetChainID(), testChain.Provider)
		assert.NoError(t, err)

		assert.Equal(t, uint16(3), recoveredWalletConfig.Threshold())
		assert.Equal(t, weight.Cmp(big.NewInt(int64(2))), 0)
		assert.Len(t, recoveredWalletConfig.Signers(), 2)

		address, err := sequence.AddressFromWalletConfig(walletConfig, wallet.GetWalletContext())
		assert.NoError(t, err)
		assert.Equal(t, wallet.Address(), address)
	})
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

	wallet, err := sequence.GenericNewWallet[*v1.WalletConfig](sequence.WalletOptions[*v1.WalletConfig]{
		Config:          walletConfig,
		SkipSortSigners: true,
	}, eoa1)
	assert.NoError(t, err)

	// Wallet config should remain the same
	walletFromGet := wallet.GetWalletConfig()
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

	sequence.V1SortWalletConfig(walletConfig)

	wallet, err := sequence.GenericNewWallet[*v1.WalletConfig](sequence.WalletOptions[*v1.WalletConfig]{
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

type NotFirstTestSigner struct {
	Wallet *ethwallet.Wallet
}

func (n *NotFirstTestSigner) Address() common.Address {
	return n.Wallet.Address()
}

func (n *NotFirstTestSigner) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, core.Signature[core.WalletConfig], error) {
	signContext := signing_service.SignContextFromContext(ctx)
	if signContext == nil {
		return nil, nil, errors.New("signing context not found")
	}

	if len(signContext.Signature) <= 2 {
		return nil, nil, core.ErrSigningFunctionNotReady
	}

	sig, err := n.Wallet.SignMessage(digest.Bytes())
	if err != nil {
		return nil, nil, err
	}

	return append(sig, 1), nil, nil
}

var _ sequence.SignerDigestSigner = &NotFirstTestSigner{}

func TestWalletSignDigestWithNotFirstSigner(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := &v2.WalletConfig{
		Threshold_: 3,
		Tree: &v2.WalletConfigTreeNode{
			Left: &v2.WalletConfigTreeAddressLeaf{
				Weight: 2, Address: eoa1.Address(),
			},
			Right: &v2.WalletConfigTreeAddressLeaf{
				Weight: 2, Address: eoa2.Address(),
			},
		},
	}

	wallet, err := sequence.NewWallet(sequence.WalletOptions[*v2.WalletConfig]{
		Config: walletConfig,
	}, &NotFirstTestSigner{Wallet: eoa2}, eoa1)
	require.NoError(t, err)

	wallet.SetChainID(big.NewInt(1))

	sig, _, err := wallet.SignMessage(context.Background(), []byte("hello world"))
	require.NoError(t, err)
	require.NotNil(t, sig)

	wallet, err = sequence.NewWallet(sequence.WalletOptions[*v2.WalletConfig]{
		Config: walletConfig,
	}, eoa1, &NotFirstTestSigner{Wallet: eoa2})
	require.NoError(t, err)

	wallet.SetChainID(big.NewInt(1))

	sig, _, err = wallet.SignMessage(context.Background(), []byte("hello world"))
	require.NoError(t, err)
	require.NotNil(t, sig)
}
