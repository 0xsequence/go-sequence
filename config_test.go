package sequence_test

import (
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	"github.com/stretchr/testify/assert"
)

func TestWalletConfigV1(t *testing.T) {
	wc := &v1.WalletConfig{
		Threshold_: 4, //big.NewInt(4),
		Signers_: v1.WalletConfigSigners{
			{Weight: 2, Address: common.HexToAddress("def")},
			{Weight: 2, Address: common.HexToAddress("abc")},
			{Weight: 2, Address: common.HexToAddress("456")},
		},
	}

	err := sequence.SortWalletConfig(wc)
	assert.NoError(t, err)

	assert.Equal(t, common.HexToAddress("456"), wc.Signers_[0].Address)
	assert.Equal(t, common.HexToAddress("abc"), wc.Signers_[1].Address)
	assert.Equal(t, common.HexToAddress("def"), wc.Signers_[2].Address)
}

func TestWalletConfigNoDupesV1(t *testing.T) {
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
	err := sequence.SortWalletConfig(wc)
	assert.Error(t, err)
}

func TestWalletConfigImageHashV1(t *testing.T) {
	wc := &v1.WalletConfig{
		Threshold_: 1, //big.NewInt(1),
		Signers_: v1.WalletConfigSigners{
			{Weight: 1, Address: common.Address{}},
		},
	}

	expected := "0xd5eb26a4673c3bf5bb325d407fe1544f0325b97d4b68afa6a28851b6dbbbd29f"

	assert.Equal(t, expected, wc.ImageHash().Hex())
}

func TestWalletAddressFromWalletConfigV1(t *testing.T) {
	wc := &v1.WalletConfig{
		Threshold_: 1, //big.NewInt(1),
		Signers_: v1.WalletConfigSigners{
			{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
		},
	}

	context := sequence.WalletContext{
		FactoryAddress:    common.HexToAddress("0x7c2C195CD6D34B8F845992d380aADB2730bB9C6F"),
		MainModuleAddress: common.HexToAddress("0x8858eeB3DfffA017D4BCE9801D340D36Cf895CCf"),
	}

	expected := common.HexToAddress("0xF0BA65550F2d1DCCf4B131B774844DC3d801D886")

	address, err := sequence.AddressFromWalletConfig(wc, context)
	assert.NoError(t, err)
	assert.Equal(t, expected, address)
}

func TestWalletIsWalletConfigUsableV1(t *testing.T) {
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
}
