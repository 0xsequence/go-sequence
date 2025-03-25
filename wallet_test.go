package sequence_test

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/relayer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWalletAddress(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
		assert.NoError(t, err)
		assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

		w, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
			MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
			CreationCode:      hexutil.Encode(contracts.V1.CreationCode),
		})

		assert.NoError(t, err)
		assert.Equal(t, "0xcBf920895f0A101b85f94903775E61d1B652b6Ca", w.Address().Hex())
	})

	t.Run("v2", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
		assert.NoError(t, err)
		assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

		w, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
			MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
			CreationCode:      hexutil.Encode(contracts.V2.CreationCode),
		})

		assert.NoError(t, err)
		assert.Equal(t, "0x686F888A26CD92E859E504d14Db1c19c58b28b7f", w.Address().Hex())
	})

	t.Run("v3", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("2bf2dfccb8c9fb4bb4d46ac9e2b537c373b44ae4c2ee66de92e02f132f7c2237")
		assert.NoError(t, err)
		assert.Equal(t, "0x1d76701Ba8B8B87Eb36C4cB30B17aea32c22846c", eoa.Address().Hex())

		w, err := sequence.GenericNewWalletSingleOwner[*v3.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
			MainModuleAddress: common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"),
			CreationCode:      hexutil.Encode(contracts.V3.CreationCode),
		})

		assert.NoError(t, err)
		assert.Equal(t, "0xEE9BD133D9EEc8EAf91F37121fcB4DefE15aD449", w.Address().Hex())
	})
}

func TestWalletSignMessage(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V1.CreationCode),
		})
		assert.NoError(t, err)
		assert.Equal(t, "0xF0BA65550F2d1DCCf4B131B774844DC3d801D886", wallet.Address().Hex())

		// we set chainID here for debugging purposes, but in general should not set it manually
		wallet.SetChainID(big.NewInt(1))

		message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

		sig, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
		assert.NoError(t, err)

		expectedSig := "0x00010001a0fb306480bc3027c04d33a16370f4618b29f2d5b89464f526045c94802bc9d1525389c364b75daf58e859ed0d6105aac6b3718e4659814c7793c626653edb871b02"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))
	})

	t.Run("v2", func(t *testing.T) {
		// base on test in https://github.com/0xsequence/sequence.js/blob/master/packages/wallet/tests/wallet.unit.spec.ts
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V2.CreationCode),
		})
		assert.NoError(t, err)
		assert.Equal(t, "0x8AD6cbc016f72971d317f281497aa080DF87a013", wallet.Address().Hex())

		// we set chainID here for debugging purposes, but in general should not set it manually
		wallet.SetChainID(big.NewInt(1))

		message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

		sig, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
		assert.NoError(t, err)

		expectedSig := "0x0100010000000000012d1d8f90e8ae7cb4dff2b63643633bb0480ca8ce4c993f3934587e0f2041b4970832b96512cfc5418a125d4110c1e3c8293c515f144d0119b4ab57b2db0c95741c02"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))
	})

	t.Run("v3", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v3.WalletConfig](sequence.EOA(eoa), sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V3.CreationCode),
		})
		assert.NoError(t, err)
		assert.Equal(t, "0xD1E0034206797BB3B58A3Ab3FAc97044D01E4Ee7", wallet.Address().Hex())

		// we set chainID here for debugging purposes, but in general should not set it manually
		wallet.SetChainID(big.NewInt(1))

		message := "0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5"

		sig, err := wallet.SignMessage(context.Background(), ethcoder.MustHexDecode(message))
		assert.NoError(t, err)

		expectedSig := "0x040001714a0c0d23d0022c93449f212c05596b4e77cdf43750bcf9ebd84f0c3beb366150e81bc833236f0053d162f5edbf3fb318d0414399250a3304ec899b7abba573ae"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))
	})
}

func TestWalletSignMessageAndValidate(t *testing.T) {
	t.Run("v3", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromPrivateKey("87306d4b9fe56c2af23c7cc3bc69914eba8f7c8fc1d35b4c9a7dd7ea198a428b")
		assert.NoError(t, err)
		assert.Equal(t, "0xd63A09C47FDc03e2Cff620446b37f205A7D0679D", eoa.Address().Hex())

		wallet, err := sequence.GenericNewWalletSingleOwner[*v3.WalletConfig](sequence.EOA(eoa))
		assert.NoError(t, err)
		assert.Equal(t, "0x7AFC40947041350252C4E86F7581fA3a4fE6eb07", wallet.Address().Hex())

		wallet.SetProvider(testChain.Provider)

		message := hexutil.MustDecode("0x1901f0ba65550f2d1dccf4b131b774844dc3d801d886bbd4edcf660f395f21fe94792f7c1da94638270a049646e541004312b3ec1ac5")

		sig, err := wallet.SignMessage(context.Background(), message)
		assert.NoError(t, err)

		expectedSig := "0x04000171a25c939935e4913e0f2fee8af8318b521903bac0b3c81d2ff09cfaf76593c508d7ee20a50cc8e1a245d6fb80edbc66f6a25f70822b69005b4d9e458d909e3992"
		assert.Equal(t, expectedSig, ethcoder.HexEncode(sig))

		isValid, err := sequence.IsValidMessageSignature(context.Background(), sig, wallet.Address(), message, testChain.Provider)
		assert.NoError(t, err)
		assert.True(t, isValid)
	})
}

func TestWalletDeploy(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		rel, err := relayer.NewLocalRelayer(testChain.MustWallet(0), nil)
		require.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](sequence.EOA(eoa))
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

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](sequence.EOA(eoa))
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

	t.Run("v3", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		rel, err := relayer.NewLocalRelayer(testChain.MustWallet(0), nil)
		require.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v3.WalletConfig](sequence.EOA(eoa))
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

		wallet, err := sequence.GenericNewWalletSingleOwner[*v1.WalletConfig](sequence.EOA(eoa))
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, err := wallet.SignMessage(context.Background(), []byte(message))
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

		wallet, err := sequence.GenericNewWalletSingleOwner[*v2.WalletConfig](sequence.EOA(eoa))
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, err := wallet.SignMessage(context.Background(), []byte(message))
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

	t.Run("v3", func(t *testing.T) {
		eoa, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		wallet, err := sequence.GenericNewWalletSingleOwner[*v3.WalletConfig](sequence.EOA(eoa))
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		s, err := v3.Core.DecodeSignature(sig)
		assert.NoError(t, err)

		payload := v3.ConstructMessagePayload(wallet.Address(), wallet.GetChainID(), []byte(message))
		recoveredWalletConfig, weight, err := s.Recover(context.Background(), core.Digest{Hash: payload.Digest().Hash}, wallet.Address(), wallet.GetChainID(), testChain.Provider)
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
	t.Run("v3", func(t *testing.T) {
		eoa1, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		eoa2, err := ethwallet.NewWalletFromRandomEntropy()
		assert.NoError(t, err)

		walletConfig := &v3.WalletConfig{
			Threshold_: 3,
			Tree: &v3.WalletConfigTreeNode{
				Left: &v3.WalletConfigTreeAddressLeaf{
					Weight: 2, Address: eoa1.Address(),
				},
				Right: &v3.WalletConfigTreeAddressLeaf{
					Weight: 5, Address: eoa2.Address(),
				},
			},
		}

		wallet, err := sequence.GenericNewWallet(
			sequence.WalletOptions[*v3.WalletConfig]{Config: walletConfig},
			sequence.EOA(eoa1),
		)
		assert.NoError(t, err)

		wallet.SetChainID(big.NewInt(3))

		message := "Hi! this is a test message"
		sig, err := wallet.SignMessage(context.Background(), []byte(message))
		assert.NoError(t, err)

		payload := v3.ConstructMessagePayload(wallet.Address(), wallet.GetChainID(), []byte(message))
		signature, err := v3.Core.DecodeSignature(sig)
		assert.NoError(t, err)
		recoveredWalletConfig, weight, err := signature.Recover(context.Background(), core.Digest{Hash: payload.Digest().Hash}, wallet.Address(), wallet.GetChainID(), testChain.Provider)
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

	wallet, err := sequence.GenericNewWallet(
		sequence.WalletOptions[*v1.WalletConfig]{
			Config:          walletConfig,
			SkipSortSigners: true,
		},
		sequence.EOA(eoa1),
	)
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

	wallet, err := sequence.GenericNewWallet(
		sequence.WalletOptions[*v1.WalletConfig]{
			Config:          walletConfig,
			Address:         randomAddr,
			SkipSortSigners: true,
		},
		sequence.EOA(eoa1),
	)
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

	nwallet, err = nwallet.UseSigners(sequence.EOA(eoa4))
	assert.NoError(t, err)

	assert.Equal(t, nwallet.Address(), randomAddr)
	assert.Equal(t, nwallet.GetWalletConfig(), newConfig)
}

type NotFirstTestSigner struct {
	Wallet *ethwallet.Wallet
}

var _ sequence.Signer = &NotFirstTestSigner{}

func (n *NotFirstTestSigner) Address() common.Address {
	return n.Wallet.Address()
}

func (n *NotFirstTestSigner) Sign(ctx context.Context, payload v3.Payload) (core.SignerSignatureType, []byte, error) {
	signature, err := n.Wallet.SignData(payload.Digest().Bytes())
	if err != nil {
		return core.SignerSignatureTypeEIP712, nil, fmt.Errorf("unable to sign: %w", err)
	}

	return core.SignerSignatureTypeEIP712, signature, nil
}

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

	wallet, err := sequence.NewWallet(
		sequence.WalletOptions[*v2.WalletConfig]{Config: walletConfig},
		&NotFirstTestSigner{Wallet: eoa2},
		sequence.EOA(eoa1),
	)
	require.NoError(t, err)

	wallet.SetChainID(big.NewInt(1))

	sig, err := wallet.SignMessage(context.Background(), []byte("hello world"))
	require.NoError(t, err)
	require.NotNil(t, sig)

	wallet, err = sequence.NewWallet(
		sequence.WalletOptions[*v2.WalletConfig]{Config: walletConfig},
		sequence.EOA(eoa1),
		&NotFirstTestSigner{Wallet: eoa2},
	)
	require.NoError(t, err)

	wallet.SetChainID(big.NewInt(1))

	sig, err = wallet.SignMessage(context.Background(), []byte("hello world"))
	require.NoError(t, err)
	require.NotNil(t, sig)
}

type MockSignerDigestSigner struct {
	mock.Mock
}

func (m *MockSignerDigestSigner) Address() common.Address {
	args := m.Called()
	return args.Get(0).(common.Address)
}

func (m *MockSignerDigestSigner) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, error) {
	args := m.Called(ctx, digest, optChainID)

	var sig []byte
	if args.Get(0) != nil {
		sig = args.Get(0).([]byte)
	}

	var err error
	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return sig, err
}

func TestWalletLazyDeployment(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	walletConfig := &v2.WalletConfig{
		Threshold_: 3,
		Tree: &v2.WalletConfigTreeAddressLeaf{
			Weight: 3, Address: eoa1.Address(),
		},
	}

	wallet, err := sequence.NewWallet(
		sequence.WalletOptions[*v2.WalletConfig]{Config: walletConfig},
		sequence.EOA(eoa1),
	)
	require.NoError(t, err)

	err = wallet.SetProvider(testChain.Provider)
	require.NoError(t, err)

	rel, err := relayer.NewLocalRelayer(testChain.MustWallet(1), nil)
	require.NoError(t, err)

	err = wallet.SetRelayer(rel)
	require.NoError(t, err)

	err = testChain.MustFundAddress(wallet.Address())
	require.NoError(t, err)

	isDeployed, err := wallet.IsDeployed()
	require.NoError(t, err)
	assert.False(t, isDeployed)

	sigTx, err := wallet.SignTransaction(context.Background(), v3.Call{
		To:    eoa1.Address(),
		Value: big.NewInt(1),
	})
	require.NoError(t, err)

	_, wait, err := wallet.SendTransaction(context.Background(), sigTx)
	require.NoError(t, err)

	receipt, err := wait(context.Background())
	require.NoError(t, err)

	assert.Equal(t, receipt.Status, uint64(1))

	isDeployed, err = wallet.IsDeployed()
	require.NoError(t, err)
	assert.True(t, isDeployed)

	bal, err := testChain.Provider.BalanceAt(context.Background(), eoa1.Address(), nil)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), bal)
}
