package v3

import (
	"context"
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
