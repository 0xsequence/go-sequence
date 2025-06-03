package v3_test

import (
	"testing"

	v3 "github.com/0xsequence/go-sequence/core/v3"
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

	config_, err := v3.Core.DecodeWalletConfig(config)
	assert.NoError(t, err)

	spew.Dump(config_)
}
