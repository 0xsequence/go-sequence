package sequence_test

import (
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/assert"
)

func TestWalletConfig(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wc := &v1.WalletConfig{
			Threshold_: 4, //big.NewInt(4),
			Signers_: v1.WalletConfigSigners{
				{Weight: 2, Address: common.HexToAddress("def")},
				{Weight: 2, Address: common.HexToAddress("abc")},
				{Weight: 2, Address: common.HexToAddress("456")},
			},
		}

		err := sequence.V1SortWalletConfig(wc)
		assert.NoError(t, err)

		assert.Equal(t, common.HexToAddress("456"), wc.Signers_[0].Address)
		assert.Equal(t, common.HexToAddress("abc"), wc.Signers_[1].Address)
		assert.Equal(t, common.HexToAddress("def"), wc.Signers_[2].Address)
	})

	t.Run("v2", func(t *testing.T) {
		wc := &v2.WalletConfig{
			Threshold_: 4, //big.NewInt(4),
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("def")},
				&v2.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("abc")},
				&v2.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("456")},
			),
		}

		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("def")]))
		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("abc")]))
		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("456")]))
	})

	t.Run("v3", func(t *testing.T) {
		wc := &v3.WalletConfig{
			Threshold_: 4, //big.NewInt(4),
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("def")},
				&v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("abc")},
				&v3.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("456")},
			),
		}

		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("def")]))
		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("abc")]))
		assert.Equal(t, 2, int(wc.Signers()[common.HexToAddress("456")]))
	})
}

func TestWalletConfigNoDupes(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wc := &v1.WalletConfig{
			Threshold_: 4, //big.NewInt(4),
			Signers_: v1.WalletConfigSigners{
				{Weight: 2, Address: common.HexToAddress("def")},
				{Weight: 2, Address: common.HexToAddress("abc")},
				{Weight: 2, Address: common.HexToAddress("456")},
				{Weight: 2, Address: common.HexToAddress("abc")},
			},
		}

		// expect dupe error
		err := sequence.V1SortWalletConfig(wc)
		assert.Error(t, err)
	})

	t.Run("v2", func(t *testing.T) {
		// todo: no way to verify that
		/*wc := &v2.WalletConfig{
			Threshold_: 4, //big.NewInt(4),
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("def")},
				&v2.WalletConfigTreeAddressLeaf{Weight: 2, Address: common.HexToAddress("abc")},
				&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("def")},
			),
		}

		err := wc.IsUsable()
		assert.Error(t, err)*/
	})
}

func TestWalletConfigImageHash(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wc := &v1.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Signers_: v1.WalletConfigSigners{
				{Weight: 1, Address: common.Address{}},
			},
		}

		expected := "0xd5eb26a4673c3bf5bb325d407fe1544f0325b97d4b68afa6a28851b6dbbbd29f"

		assert.Equal(t, expected, wc.ImageHash().Hex())
	})

	t.Run("v2", func(t *testing.T) {
		wc := &v2.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.Address{}},
			),
		}

		expected := "0xb7062abf352bdc2978b8de0748c9f08a53d9f84aa4fcc0af15384567419ccf58"

		assert.Equal(t, expected, wc.ImageHash().Hex())
	})

	t.Run("v3", func(t *testing.T) {
		wc := &v3.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.Address{}},
			),
		}

		expected := "0x2e933435748a0b51a395e77bfc6690e8d541c4b9b0b6b507e109d8db15344907"

		assert.Equal(t, expected, wc.ImageHash().Hex())
	})
}

func TestWalletAddressFromWalletConfig(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		wc := &v1.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Signers_: v1.WalletConfigSigners{
				{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
			},
		}

		context := sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V1.CreationCode),
		}

		expected := common.HexToAddress("0xF0BA65550F2d1DCCf4B131B774844DC3d801D886")

		address, err := sequence.AddressFromWalletConfig(wc, context)
		assert.NoError(t, err)
		assert.Equal(t, expected, address)
	})

	t.Run("v2", func(t *testing.T) {
		wc := &v2.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: v2.WalletConfigTreeNodes(
				&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
			),
		}

		context := sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V2.CreationCode),
		}

		expected := common.HexToAddress("0x8AD6cbc016f72971d317f281497aa080DF87a013")

		address, err := sequence.AddressFromWalletConfig(wc, context)
		assert.NoError(t, err)
		assert.Equal(t, expected, address)
	})

	t.Run("v3", func(t *testing.T) {
		wc := &v3.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: v3.WalletConfigTreeNodes(
				&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.Address{}},
			),
		}

		context := sequence.WalletContext{
			FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
			MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
			CreationCode:      hexutil.Encode(contracts.V3.CreationCode),
		}

		expected := common.HexToAddress("0xce06CE19FB2Fd57CA337Cf484BA5A37E22DF9f35")

		address, err := sequence.AddressFromWalletConfig(wc, context)
		assert.NoError(t, err)
		assert.Equal(t, expected, address)
	})
}

func TestWalletIsWalletConfigUsable(t *testing.T) {
	t.Run("v1", func(t *testing.T) {
		{
			wcGood := &v1.WalletConfig{
				Threshold_: 1, //big.NewInt(1),
				Signers_: v1.WalletConfigSigners{
					{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				},
			}

			assert.NoError(t, wcGood.IsUsable())
		}
		{
			wcBad := &v1.WalletConfig{
				Threshold_: 0, //big.NewInt(0),
				Signers_: v1.WalletConfigSigners{
					{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				},
			}

			assert.Error(t, wcBad.IsUsable())
		}
		{
			wcBad := &v1.WalletConfig{
				Threshold_: 2, //big.NewInt(2),
				Signers_: v1.WalletConfigSigners{
					{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				},
			}

			assert.Error(t, wcBad.IsUsable())
		}
	})

	t.Run("v2", func(t *testing.T) {
		{
			wcGood := &v2.WalletConfig{
				Threshold_: 1, //big.NewInt(1),
				Tree: v2.WalletConfigTreeNodes(
					&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.NoError(t, wcGood.IsUsable())
		}
		{
			wcBad := &v2.WalletConfig{
				Threshold_: 0, //big.NewInt(0),
				Tree: v2.WalletConfigTreeNodes(
					&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.Error(t, wcBad.IsUsable())
		}
		{
			wcBad := &v2.WalletConfig{
				Threshold_: 2, //big.NewInt(2),
				Tree: v2.WalletConfigTreeNodes(
					&v2.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.Error(t, wcBad.IsUsable())
		}
	})

	t.Run("v3", func(t *testing.T) {
		{
			wcGood := &v3.WalletConfig{
				Threshold_: 1, //big.NewInt(1),
				Tree: v3.WalletConfigTreeNodes(
					&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.NoError(t, wcGood.IsUsable())
		}
		{
			wcBad := &v3.WalletConfig{
				Threshold_: 0, //big.NewInt(0),
				Tree: v3.WalletConfigTreeNodes(
					&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.Error(t, wcBad.IsUsable())
		}
		{
			wcBad := &v3.WalletConfig{
				Threshold_: 2, //big.NewInt(2),
				Tree: v3.WalletConfigTreeNodes(
					&v3.WalletConfigTreeAddressLeaf{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
				),
			}

			assert.Error(t, wcBad.IsUsable())
		}
	})
}
