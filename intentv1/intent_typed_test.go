package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/0xsequence/ethkit"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntentNewIntentTyped(t *testing.T) {
	t.Run("openSession", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})
		assert.Equal(t, "openSession", intent.Name)
	})

	t.Run("closeSession", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataCloseSession{SessionId: "0x1234"})
		assert.Equal(t, "closeSession", intent.Name)
	})

	t.Run("validateSession", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataValidateSession{SessionId: "0x1234"})
		assert.Equal(t, "validateSession", intent.Name)
	})

	t.Run("finishValidateSession", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataFinishValidateSession{SessionId: "0x1234"})
		assert.Equal(t, "finishValidateSession", intent.Name)
	})

	t.Run("listSessions", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataListSessions{})
		assert.Equal(t, "listSessions", intent.Name)
	})

	t.Run("getSession", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataGetSession{SessionId: "0x1234"})
		assert.Equal(t, "getSession", intent.Name)
	})

	t.Run("sign", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataSign{Network: "ethereum", Message: "0x1234"})
		assert.Equal(t, "sign", intent.Name)
	})

	t.Run("transaction", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataSendTransaction{})
		assert.Equal(t, "sendTransaction", intent.Name)
	})

	t.Run("unknown", func(t *testing.T) {
		intent := NewIntentTyped(map[string]interface{}{})
		assert.Equal(t, "", intent.Name)
	})
}

func TestIntentNewIntentTypedFromIntent(t *testing.T) {
	t.Run("openSession", func(t *testing.T) {
		intent := Intent{
			Version: "1",
			Name:    "openSession",
			Data:    map[string]interface{}{"sessionId": "0x1234"},
		}

		intentTyped, err := NewIntentTypedFromIntent[IntentDataOpenSession](&intent)
		assert.NoError(t, err)
		assert.Equal(t, "openSession", intentTyped.Name)
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
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_p256k1Signature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_p256r1Signature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)

		session := NewSessionP256R1(privateKey)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("valid_legacySignature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		err = SignIntentWithWalletLegacy(wallet, intent)
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("expired", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})
		intent.ExpiresAt = 0

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "expired")
	})

	t.Run("issuedInFuture", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})
		intent.IssuedAt = uint64(1 << 63)

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "issued in the future")
	})

	t.Run("noSignatures", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		assert.ErrorContains(t, intent.IsValid(), "no signatures")
	})

	t.Run("invalidSignature", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		intent.Signatures[0].Signature = "0x1234"

		assert.ErrorContains(t, intent.IsValid(), "invalid signature")
	})
}

func TestIntentDataValidator(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234", Email: ethkit.ToPtr("test@test.com")})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.NoError(t, intent.IsValid())
	})

	t.Run("invalid", func(t *testing.T) {
		intent := NewIntentTyped(IntentDataOpenSession{SessionId: "0x1234"})

		wallet, err := ethwallet.NewWalletFromRandomEntropy()
		require.NoError(t, err)

		session := NewSessionP256K1(wallet)
		err = session.Sign(intent.AsIntent())
		require.NoError(t, err)

		assert.ErrorContains(t, intent.IsValid(), "invalid intent data")
	})
}
