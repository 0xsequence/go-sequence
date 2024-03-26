package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntentNewIntentTyped(t *testing.T) {
	t.Run(string(IntentName_openSession), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234"})
		assert.Equal(t, IntentName_openSession, intent.Name)
	})

	t.Run(string(IntentName_closeSession), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataCloseSession{SessionID: "0x1234"})
		assert.Equal(t, IntentName_closeSession, intent.Name)
	})

	t.Run(string(IntentName_validateSession), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataValidateSession{SessionID: "0x1234"})
		assert.Equal(t, IntentName_validateSession, intent.Name)
	})

	t.Run(string(IntentName_finishValidateSession), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataFinishValidateSession{SessionID: "0x1234"})
		assert.Equal(t, IntentName_finishValidateSession, intent.Name)
	})

	t.Run(string(IntentName_listSessions), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataListSessions{})
		assert.Equal(t, IntentName_listSessions, intent.Name)
	})

	t.Run(string(IntentName_getSession), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataGetSession{SessionID: "0x1234"})
		assert.Equal(t, IntentName_getSession, intent.Name)
	})

	t.Run(string(IntentName_signMessage), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataSignMessage{Network: "ethereum", Message: "0x1234"})
		assert.Equal(t, IntentName_signMessage, intent.Name)
	})

	t.Run(string(IntentName_sendTransaction), func(t *testing.T) {
		intent := NewIntentTyped(IntentDataSendTransaction{})
		assert.Equal(t, IntentName_sendTransaction, intent.Name)
	})

	t.Run("unknown", func(t *testing.T) {
		intent := NewIntentTyped(map[string]interface{}{})
		assert.Equal(t, IntentName(""), intent.Name)
	})
}

func TestIntentNewIntentTypedFromIntent(t *testing.T) {
	t.Run("openSession", func(t *testing.T) {
		intent := Intent{
			Version: "1",
			Name:    IntentName_openSession,
			Data:    map[string]interface{}{"sessionId": "0x1234"},
		}

		intentTyped, err := NewIntentTypedFromIntent[IntentDataOpenSession](&intent)
		assert.NoError(t, err)
		assert.Equal(t, IntentName_openSession, intentTyped.Name)
	})

	t.Run("openSessionNameMismatch", func(t *testing.T) {
		intent := Intent{
			Version: "1",
			Name:    "openSession2",
			Data:    map[string]interface{}{"sessionId": "0x1234"},
		}

		_, err := NewIntentTypedFromIntent[IntentDataCloseSession](&intent)
		assert.ErrorContains(t, err, "mismatch")
	})
}

func TestIntentIsValid(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_p256k1Signature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_p256r1Signature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)

		session := NewSessionP256R1(privateKey)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_legacySignature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		err = SignIntentWithWalletLegacy(wallet, intent.ToIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("expired", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})
		intent.ExpiresAt = 0

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "expired")
	})

	t.Run("issuedInFuture", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})
		intent.IssuedAt = uint64(1 << 63)

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "issued in the future")
	})

	t.Run("noSignatures", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		assert.ErrorContains(t, intent.IsValid(), "no signatures")
	})

	t.Run("invalidSignature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		intent.Signatures[0].Signature = "0x1234"

		assert.ErrorContains(t, intent.IsValid(), "invalid signature")
	})
}

func TestIntentDataValidator(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("invalid", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.ToIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "invalid intent data")
	})
}

func TestIntent_Hash(t *testing.T) {
	wallet, _ := ethwallet.NewWalletFromRandomEntropy()

	session := NewSessionP256K1(wallet)

	intentTyped := NewIntentTyped(IntentDataOpenSession{SessionID: "0x1234", Email: ethkit.ToPtr("test@test.com")})

	intent := Intent{
		Version:   intentTyped.Version,
		Name:      intentTyped.Name,
		ExpiresAt: intentTyped.ExpiresAt,
		IssuedAt:  intentTyped.IssuedAt,
		Data:      map[string]interface{}{"sessionId": "0x1234", "email": "test@test.com"},
	}

	err := session.Sign(intentTyped.ToIntent())
	require.NoError(t, err)

	err = session.Sign(&intent)
	require.NoError(t, err)

	hashIntentTyped, _ := intentTyped.Hash()
	hashIntent, _ := intent.Hash()

	assert.Equal(t, hashIntentTyped, hashIntent)
}

func TestIntentTyped_Serialization_DoubleData(t *testing.T) {
	intent := Intent{
		Version:   "1",
		Name:      "openSession",
		ExpiresAt: 0,
		IssuedAt:  0,
		Data:      map[string]interface{}{"sessionId": "0x1234", "email": "test@test.com"},
	}

	intentTyped, err := NewIntentTypedFromIntent[IntentDataOpenSession](&intent)
	require.NoError(t, err)

	intentTypedJSON, err := json.Marshal(intentTyped)
	require.NoError(t, err)

	var intentMap map[string]any
	err = json.Unmarshal(intentTypedJSON, &intentMap)
	require.NoError(t, err)

	assert.Equal(t, "1", intentMap["version"])
	assert.Equal(t, "openSession", intentMap["name"])
	assert.Equal(t, "0x1234", intentMap["data"].(map[string]any)["sessionId"])
	assert.Equal(t, "test@test.com", intentMap["data"].(map[string]any)["email"])
	assert.Equal(t, float64(0), intentMap["expiresAt"])
	assert.Equal(t, float64(0), intentMap["issuedAt"])
	assert.Equal(t, nil, intentMap["Data"])
}
