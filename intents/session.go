package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

type KeyType int

const (
	KeyTypeSECP256K1 KeyType = iota
	KeyTypeSECP256R1
	KeyTypeUnknown
)

type Session interface {
	SessionID() string
	Sign(intent *Intent) error
}

type session256K1 struct {
	wallet *ethwallet.Wallet
}

func NewSessionP256K1(wallet *ethwallet.Wallet) Session {
	return &session256K1{wallet: wallet}
}

func (s session256K1) SessionID() string {
	return strings.ToLower(
		fmt.Sprintf(
			"0x%s",
			common.Bytes2Hex(
				append([]byte{byte(KeyTypeSECP256K1)}, s.wallet.Address().Bytes()...),
			),
		),
	)
}

func (s session256K1) Sign(intent *Intent) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	signature, err := s.wallet.SignMessage(hash)
	if err != nil {
		return err
	}

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionID: s.SessionID(),
		Signature: bytesToSignature(signature),
	})
	return nil
}

type session256R1 struct {
	privateKey *ecdsa.PrivateKey
}

func NewSessionP256R1(privateKey *ecdsa.PrivateKey) Session {
	return &session256R1{privateKey: privateKey}
}

func (s session256R1) SessionID() string {
	pubKey := elliptic.Marshal(s.privateKey.Curve, s.privateKey.PublicKey.X, s.privateKey.PublicKey.Y)
	return strings.ToLower(
		fmt.Sprintf(
			"0x%s",
			common.Bytes2Hex(
				append([]byte{byte(KeyTypeSECP256R1)}, pubKey...),
			),
		),
	)
}

func (s session256R1) Sign(intent *Intent) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	sha256Hash := sha256.Sum256(hash)

	sr, ss, err := ecdsa.Sign(rand.Reader, s.privateKey, sha256Hash[:])
	if err != nil {
		return err
	}

	signature := append(sr.Bytes(), ss.Bytes()...)

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionID: s.SessionID(),
		Signature: bytesToSignature(signature),
	})
	return nil
}

func SignIntentWithWalletLegacy(wallet *ethwallet.Wallet, intent *Intent) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	signature, err := wallet.SignMessage(hash)
	if err != nil {
		return err
	}

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionID: strings.ToLower(wallet.Address().String()),
		Signature: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(signature))),
	})
	return nil
}

func SignIntentP256K1(wallet *ethwallet.Wallet, intent *Intent) error {
	return NewSessionP256K1(wallet).Sign(intent)
}

func SignIntentP256R1(privateKey *ecdsa.PrivateKey, intent *Intent) error {
	return NewSessionP256R1(privateKey).Sign(intent)
}

func IsValidSessionSignature(sessionId string, signature string, intent *Intent) error {
	switch KeyTypeFromSessionId(sessionId) {
	case KeyTypeSECP256K1:
		return IsValidSessionSignatureP256K1(sessionId, signature, intent)
	case KeyTypeSECP256R1:
		return IsValidSessionSignatureP256R1(sessionId, signature, intent)
	default:
		return fmt.Errorf("unknown session key type")
	}
}

// IsValidSessionSignatureP256K1 checks if the signature is valid for the given secp256k1 session
func IsValidSessionSignatureP256K1(sessionId string, signature string, intent *Intent) error {
	// validate session address and signature
	if !strings.HasPrefix(signature, "0x") || !strings.HasPrefix(sessionId, "0x") {
		// invalid params
		return fmt.Errorf("invalid sessionId or signature format")
	}

	// validate session address
	sessionAddressBytes := common.FromHex(sessionId)
	if len(sessionAddressBytes) != 21 && len(sessionAddressBytes) != 20 {
		// invalid session address
		return fmt.Errorf("invalid sessionId length")
	}

	if len(sessionAddressBytes) == 21 && sessionAddressBytes[0] != byte(KeyTypeSECP256K1) {
		return fmt.Errorf("invalid sessionId format")
	}

	// validate signature
	sigBytes := common.FromHex(signature)
	if len(sigBytes) != 65 {
		return fmt.Errorf("invalid signature length")
	}

	// handle typed session address
	if len(sessionAddressBytes) == 21 {
		sessionAddressBytes = sessionAddressBytes[1:]
	}

	sessionAddress := fmt.Sprintf("0x%s", common.Bytes2Hex(sessionAddressBytes))

	// Get hash of the packet
	hash, err := intent.Hash()
	if err != nil {
		return fmt.Errorf("failed to hash intent: %w", err)
	}

	// Add Ethereum prefix to the hash
	prefixedHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)))

	// handle recovery byte
	if sigBytes[64] == 27 || sigBytes[64] == 28 {
		sigBytes[64] -= 27
	}

	// Recover the public key from the signature
	pubKey, err := crypto.Ecrecover(prefixedHash.Bytes(), sigBytes)
	if err != nil {
		return fmt.Errorf("failed to recover public key: %w", err)
	}

	addr := common.BytesToAddress(crypto.Keccak256(pubKey[1:])[12:])

	// Check if the recovered address matches the session address
	if strings.ToLower(addr.Hex()) != strings.ToLower(sessionAddress) {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

// IsValidSessionSignatureP256R1 checks if the signature is valid for the given secp256r1 session
func IsValidSessionSignatureP256R1(sessionId string, signature string, intent *Intent) error {
	// validate session address and signature
	if !strings.HasPrefix(signature, "0x") || !strings.HasPrefix(sessionId, "0x") {
		// invalid params
		return fmt.Errorf("invalid sessionId or signature format")
	}

	// validate session id
	sessionIdBytes := common.FromHex(sessionId)
	if len(sessionIdBytes) != 66 {
		return fmt.Errorf("invalid sessionId length")
	}

	if sessionIdBytes[0] != byte(KeyTypeSECP256R1) {
		return fmt.Errorf("invalid sessionId format")
	}

	// validate signature
	signatureBytes := common.FromHex(signature)
	if len(signatureBytes) != 64 {
		return fmt.Errorf("invalid signature length")
	}

	// message hash
	messageHash, _ := intent.Hash()
	messageHash2 := sha256.Sum256(messageHash)

	// public key
	publicKeyBuff := common.FromHex(sessionId)[1:]
	x, y := elliptic.Unmarshal(elliptic.P256(), publicKeyBuff)
	if x == nil || y == nil {
		return fmt.Errorf("invalid public key")
	}

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// signature
	r := new(big.Int).SetBytes(signatureBytes[:32])
	s := new(big.Int).SetBytes(signatureBytes[32:64])
	if !ecdsa.Verify(&pub, messageHash2[:], r, s) {
		return fmt.Errorf("invalid signature")
	}
	return nil
}

func KeyTypeFromSessionId(sessionId string) KeyType {
	// handle empty session ids
	if len(sessionId) == 0 {
		return KeyTypeUnknown
	}

	// handle old session ids
	if len(sessionId) == 42 {
		return KeyTypeSECP256K1
	}

	// handle key typed session ids
	sessionIdBytes := common.FromHex(sessionId)
	return KeyType(sessionIdBytes[0])
}

func bytesToSignature(sig []byte) string {
	return strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(sig)))
}
