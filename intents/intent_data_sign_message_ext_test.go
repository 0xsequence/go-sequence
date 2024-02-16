package intents

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoverMessageIntent(t *testing.T) {
	data := `{
		"version": "1",
		"name": "signMessage",
		"issued": 0,
		"expires": 0,
		"data": {
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"message": "0xdeadbeef"
		},
		"signatures": []
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1", intent.Version)
	assert.Equal(t, "signMessage", intent.Name)

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, common.Bytes2Hex(hash))

	intent.IssuedAt = uint64(time.Now().Unix())
	intent.ExpiresAt = uint64(time.Now().Unix()) + 60

	wallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.Nil(t, err)

	session := NewSessionP256K1(wallet)

	err = session.Sign(intent)
	require.Nil(t, err)

	intentTyped, err := NewIntentTypedFromIntent[IntentDataSignMessage](intent)
	require.NoError(t, err)

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x"+common.Bytes2Hex(append([]byte{0x00}, wallet.Address().Bytes()...)), signers[0])

	subdigest, err := sequence.SubDigest(
		big.NewInt(1),
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		sequence.MessageDigest(sequence.MessageToEIP191(common.Hex2Bytes("deadbeef"))),
	)

	assert.Nil(t, err)
	assert.True(t, intentTyped.Data.IsValidInterpretation(common.BytesToHash(subdigest)))
}
