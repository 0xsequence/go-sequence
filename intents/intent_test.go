package intents

import (
	"encoding/json"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestParseAndRecoverIntent(t *testing.T) {
	data := `{
		"version": "1.0.0",
		"packet": {
			"code": "sendTransactions",
			"identifier": "test-identifier",
			"issued": 1600000000,
			"expires": 1600086400,
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"transactions": [{
				"type": "erc20send",
				"token": "0x0000000000000000000000000000000000000000",
				"to": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425",
				"value": "0"
			}]
		},
		"signatures": [{
			"session": "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B",
			"signature": "0xcca6253c4fd281247ddd0fa487252ef91932eaec8d68b61f0901ccaa70345bf66fdbbd98ed3e3c9752f9e35ef2a7bc88dd9c8ae23c594241b476fe988824ab881c"
		}]
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1.0.0", intent.Version)

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])
	assert.Equal(t, intent.PacketCode(), "sendTransactions")

	// Changing the version should not affect the hash
	intent.Version = "2.0.0"
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")

	// Changing the packet code SHOULD affect the hash (and make Signers() return empty)
	intent.Packet = json.RawMessage(`{"code": "sendTransactions2"}`)
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.NotEqual(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")
	assert.Equal(t, intent.PacketCode(), "sendTransactions2")

	signers = intent.Signers()
	assert.Equal(t, 0, len(signers))

	// Parsing the JSON without tabs, spaces, newlines, etc. should still work
	// and produce the same hash
	data2 := `{"signatures":[{"signature":"0xcca6253c4fd281247ddd0fa487252ef91932eaec8d68b61f0901ccaa70345bf66fdbbd98ed3e3c9752f9e35ef2a7bc88dd9c8ae23c594241b476fe988824ab881c","session":"0x1111BD4F3233e7a7f552AdAf32C910fD30de598B"}],"version":"1.0.0","packet":{"transactions":[{"token":"0x0000000000000000000000000000000000000000","value":"0","type":"erc20send","to":"0x0dc9603d4da53841C1C83f3B550C6143e60e0425"}],"wallet":"0xD67FC48b298B09Ed3D03403d930769C527186c4e","expires":1600086400,"code":"sendTransactions","network":"1","identifier":"test-identifier","issued":1600000000}}`
	intent2 := &Intent{}
	err = json.Unmarshal([]byte(data2), intent2)
	assert.Nil(t, err)

	hash2, err := intent2.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash2)
	assert.Equal(t, common.Bytes2Hex(hash2), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")
}
