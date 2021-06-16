package sequence_test

import (
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
)

func TestWalletConfig(t *testing.T) {
	wc := sequence.WalletConfig{
		Threshold: 4, //big.NewInt(4),
		Signers: sequence.WalletConfigSigners{
			{Weight: 2, Address: common.HexToAddress("def")},
			{Weight: 2, Address: common.HexToAddress("abc")},
			{Weight: 2, Address: common.HexToAddress("456")},
		},
	}

	err := sequence.SortWalletConfig(wc)
	assert.NoError(t, err)

	assert.Equal(t, common.HexToAddress("456"), wc.Signers[0].Address)
	assert.Equal(t, common.HexToAddress("abc"), wc.Signers[1].Address)
	assert.Equal(t, common.HexToAddress("def"), wc.Signers[2].Address)
}

func TestWalletConfigNoDupes(t *testing.T) {
	wc := sequence.WalletConfig{
		Threshold: 4, //big.NewInt(4),
		Signers: sequence.WalletConfigSigners{
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

func TestWalletConfigImageHash(t *testing.T) {
	wc := sequence.WalletConfig{
		Threshold: 1, //big.NewInt(1),
		Signers: sequence.WalletConfigSigners{
			{Weight: 1, Address: common.Address{}},
		},
	}

	expected := "0xd5eb26a4673c3bf5bb325d407fe1544f0325b97d4b68afa6a28851b6dbbbd29f"

	imageHash, err := sequence.ImageHashOfWalletConfig(wc)
	assert.NoError(t, err)
	assert.Equal(t, expected, imageHash)
}

func TestWalletAddressFromWalletConfig(t *testing.T) {
	wc := sequence.WalletConfig{
		Threshold: 1, //big.NewInt(1),
		Signers: sequence.WalletConfigSigners{
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

func TestWalletIsWalletConfigUsable(t *testing.T) {
	{
		wcGood := sequence.WalletConfig{
			Threshold: 1, //big.NewInt(1),
			Signers: sequence.WalletConfigSigners{
				{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
			},
		}

		ok, err := sequence.IsWalletConfigUsable(wcGood)
		assert.NoError(t, err)
		assert.True(t, ok)
	}
	{
		wcBad := sequence.WalletConfig{
			Threshold: 0, //big.NewInt(0),
			Signers: sequence.WalletConfigSigners{
				{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
			},
		}

		ok, err := sequence.IsWalletConfigUsable(wcBad)
		assert.Error(t, err)
		assert.False(t, ok)
	}
	{
		wcBad := sequence.WalletConfig{
			Threshold: 2, //big.NewInt(2),
			Signers: sequence.WalletConfigSigners{
				{Weight: 1, Address: common.HexToAddress("0xd63A09C47FDc03e2Cff620446b37f205A7D0679D")},
			},
		}

		ok, err := sequence.IsWalletConfigUsable(wcBad)
		assert.Error(t, err)
		assert.False(t, ok)
	}
}
