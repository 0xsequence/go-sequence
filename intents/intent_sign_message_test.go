package intents

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/intents/packets"
	"github.com/stretchr/testify/assert"
)

func TestRecoverMessageIntent(t *testing.T) {
	data := `{
		"version": "1.0.0",
		"packet": {
			"code": "signMessage",
			"issued": 1600000000,
			"expires": 1600086400,
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"message": "0xdeadbeef"
		},
		"signatures": [{
			"session": "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B",
			"signature": "0x827b2a2afbf4a8a79e761fdb26e567b519a56a06e897dce5517b3ccfb408b55f20aaba276c1dade28112f51fe7262fbd0508da0019c0f8582c41b2be451ddede1b"
		}]
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1.0.0", intent.Version)
	assert.Equal(t, "signMessage", intent.PacketCode())

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.Equal(t, common.Bytes2Hex(hash), "5b15538a25716e951630dde1cf38ae056d764976145d1134576461203a621ddb")

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])

	packet := &packets.SignMessagePacket{}
	err = packet.Unmarshal(intent.Packet)
	assert.Nil(t, err)

	assert.Equal(t, "signMessage", packet.Code)

	subdigest, err := sequence.SubDigest(
		big.NewInt(1),
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		sequence.MessageDigest(common.Hex2Bytes("deadbeef")),
	)

	assert.Nil(t, err)
	assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest)))
}
