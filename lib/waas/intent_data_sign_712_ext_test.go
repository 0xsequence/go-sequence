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

func TestRecoverTypedDataIntent(t *testing.T) {
	data := `{
		"version": "1",
		"name": "signTypedData",
		"issued": 0,
		"expires": 0,
		"data": {
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"typedData": {
				"types": {
					"EIP712Domain": [
						{"name": "name", "type": "string"},
						{"name": "version", "type": "string"},
						{"name": "chainId", "type": "uint256"},
						{"name": "verifyingContract", "type": "address"}
					],
					"Person": [
						{"name": "name", "type": "string"},
						{"name": "wallet", "type": "address"},
						{"name": "count", "type": "uint8"}
					]
				},
				"primaryType": "Person",
				"domain": {
					"name": "Ether Mail",
					"version": "1",
					"chainId": 1,
					"verifyingContract": "0xCcCCccccCCCCcCCCCCCcCcCccCcCCCcCcccccccC"
				},
				"message": {
					"name": "Bob",
					"wallet": "0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB",
					"count": 4
				}
			}
		},
		"signatures": []
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	require.Nil(t, err)

	require.Equal(t, "1", intent.Version)
	require.Equal(t, IntentName_signTypedData, intent.Name)

	hash, err := intent.Hash()
	require.Nil(t, err)
	require.NotNil(t, common.Bytes2Hex(hash))

	intent.IssuedAt = uint64(time.Now().Unix())
	intent.ExpiresAt = uint64(time.Now().Unix()) + 60

	wallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.Nil(t, err)

	session := NewSessionP256K1(wallet)

	err = session.Sign(intent)
	require.Nil(t, err)

	intentTyped, err := NewIntentTypedFromIntent[IntentDataSignTypedData](intent)
	require.NoError(t, err)

	signers := intent.Signers()
	require.Equal(t, 1, len(signers))
	require.Equal(t, "0x"+common.Bytes2Hex(append([]byte{0x00}, wallet.Address().Bytes()...)), signers[0])

	messageDigest, err := intentTyped.Data.TypedData.EncodeDigest()
	require.NoError(t, err)
	require.NotNil(t, messageDigest)

	subdigest, err := sequence.SubDigest(
		big.NewInt(1),
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		common.BytesToHash(messageDigest),
	)

	assert.Nil(t, err)
	assert.True(t, intentTyped.Data.IsValidInterpretation(common.BytesToHash(subdigest)))

}
