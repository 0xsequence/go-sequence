package packets

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/intents"
	"github.com/stretchr/testify/assert"
)

func TestRecoverMessageIntent(t *testing.T) {
	data := `{
		"version": "1.0.0",
		"packet": {
			"code": "signMessage",
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"message": "0xdeadbeef"
		},
		"signatures": [{
			"session": "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B",
			"signature": "0xfa8697812a75e3893b7d498695fe699bcf00cfc84a6ca8815bd0a23021e631176dcbfeb7f07b942386277ebc821dc15ac6520e80e4d6d3947e556376d5a74e051c"
		}]
	}`

	intent := &intents.Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1.0.0", intent.Version)
	assert.Equal(t, "signMessage", intent.PacketCode())

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.Equal(t, common.Bytes2Hex(hash), "26abd75b9e3f2f8f8a777f8e0ac628edb136dbe82a461c9c7ee4b102d654cb87")

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])

	packet := &SignMessagePacket{}
	err = packet.Unmarshal(intent)
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
