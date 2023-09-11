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
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"chainId": 1,
			"transactions": [{
				"type": "erc20send",
				"token": "0x0000000000000000000000000000000000000000",
				"to": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425",
				"value": "0"
			}]
		},
		"signatures": [{
			"session": "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B",
			"signature": "0xd22e5853e09ea29812774fd658811aa6007781d3a7fa23bf5ab69d58cd18cd6467cf1a7969b982695e35db5514b90313c53a9eb941a986d572b35c5bd1d0a46f1b"
		}]
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1.0.0", intent.Version)

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "e54e41eca96c1c047ad6a31f80dd3e61ba5f8a6e0a8a83b95e58515afe1780da")

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])

	// Changing the version should not affect the hash
	intent.Version = "2.0.0"
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "e54e41eca96c1c047ad6a31f80dd3e61ba5f8a6e0a8a83b95e58515afe1780da")

	// Changing the packet code SHOULD affect the hash (and make Signers() return empty)
	intent.Packet = json.RawMessage(`{"code": "sendTransactions2"}`)
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.NotEqual(t, common.Bytes2Hex(hash), "e54e41eca96c1c047ad6a31f80dd3e61ba5f8a6e0a8a83b95e58515afe1780da")

	signers = intent.Signers()
	assert.Equal(t, 0, len(signers))

	// Parsing the JSON without tabs, spaces, newlines, etc. should still work
	// and produce the same hash
	data2 := `{"version":"1.0.0","packet":{"code":"sendTransactions","wallet":"0xD67FC48b298B09Ed3D03403d930769C527186c4e","chainId":1,"transactions":[{"type":"erc20send","token":"0x0000000000000000000000000000000000000000","to":"0x0dc9603d4da53841C1C83f3B550C6143e60e0425","value":"0"}]},"signatures":[{"session":"0x1111BD4F3233e7a7f552AdAf32C910fD30de598B","signature":"0xd22e5853e09ea29812774fd658811aa6007781d3a7fa23bf5ab69d58cd18cd6467cf1a7969b982695e35db5514b90313c53a9eb941a986d572b35c5bd1d0a46f1b"}]}`

	intent2 := &Intent{}
	err = json.Unmarshal([]byte(data2), intent2)
	assert.Nil(t, err)

	hash2, err := intent2.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash2)
	assert.Equal(t, common.Bytes2Hex(hash2), "e54e41eca96c1c047ad6a31f80dd3e61ba5f8a6e0a8a83b95e58515afe1780da")
}
