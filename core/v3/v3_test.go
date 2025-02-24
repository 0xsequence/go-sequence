package v3

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

const configTOML = `
	threshold = 5
	checkpoint = 1

	[tree.left.left]
	weight = 3
	address = "0x1111111111111111111111111111111111111111"

	[tree.left.right]
	weight = 3
	address = "0x2222222222222222222222222222222222222222"

	[tree.right.left]
	weight = 2
	address = "0x3333333333333333333333333333333333333333"

	[tree.right.right]
	weight = 2
	threshold = 1

	[tree.right.right.tree.left]
	weight = 1
	address = "0x4444444444444444444444444444444444444444"

	[tree.right.right.tree.right]
	weight = 1
	address = "0x5555555555555555555555555555555555555555"
`

func TestWalletConfigTOML(t *testing.T) {
	var config map[string]any
	_, err := toml.Decode(configTOML, &config)
	assert.NoError(t, err)

	config_, err := Core.DecodeWalletConfig(config)
	assert.NoError(t, err)

	spew.Dump(config_)
}

var noSigner = func(ctx context.Context, signer common.Address, sigs []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
	return 0, nil, core.ErrSigningNoSigner
}

func TestEncodeAddressLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeAddressLeaf{
			Weight:  1,
			Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000111" + "1234567890abcdef1234567890abcdef12345678"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}

func TestEncodeNestedLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeNestedLeaf{
			Weight:    1,
			Threshold: 1,
			Tree: &WalletConfigTreeAddressLeaf{
				Weight:  1,
				Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			},
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000165000015111234567890abcdef1234567890abcdef12345678"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}

func TestEncodeSubdigestLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeSubdigestLeaf{
			Subdigest: core.Subdigest{Hash: common.HexToHash("0x" + strings.Repeat("abcdef", 10) + "abcd")},
		},
	}

	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000150" + "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}

func TestEncodeSapientSignerLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeNodeLeaf{
			Node: core.ImageHash{Hash: common.HexToHash("0x51a314ac5ecb4db7c39e7b5f50df7fa964ed7c2cd55d23c8ac3bea4e4693e879")},
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000130" + "51a314ac5ecb4db7c39e7b5f50df7fa964ed7c2cd55d23c8ac3bea4e4693e879"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}

func TestEncodeSapientCompactLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeAddressLeaf{
			Weight:  1,
			Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000111" + "1234567890abcdef1234567890abcdef12345678"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}

func TestEncodeAnyAddressSubdigestLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeAnyAddressSubdigestLeaf{
			Digest: core.Subdigest{Hash: common.HexToHash("0x" + strings.Repeat("abcdef", 10) + "abcd")},
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000180" + "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	t.Logf("Actual encoding: %s", actual)
	assert.Equal(t, expected, actual)
}

func TestEncodeNodeLeaf(t *testing.T) {
	config := &WalletConfig{
		Threshold_:  1,
		Checkpoint_: 0,
		Tree: &WalletConfigTreeNode{
			Left: &WalletConfigTreeAddressLeaf{
				Weight:  1,
				Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			},
			Right: &WalletConfigTreeAddressLeaf{
				Weight:  1,
				Address: common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"),
			},
		},
	}
	sig, err := config.BuildNoChainIDSignature(context.Background(), noSigner, false)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = sig.Write(&buf)
	assert.NoError(t, err)
	expected := "0x06000111" + "1234567890abcdef1234567890abcdef12345678" + "11" + "abcdefabcdefabcdefabcdefabcdefabcdefabcd"
	actual := "0x" + common.Bytes2Hex(buf.Bytes())
	assert.Equal(t, expected, actual)
}
