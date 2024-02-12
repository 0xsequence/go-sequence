package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/gibson042/canonicaljson-go"
)

const IntentValidTimeInSec = 60
const IntentAllowedTimeDriftInSec = 5

type KeyType int

const (
	KeyTypeSECP256K1 KeyType = iota
	KeyTypeSECP256R1
	KeyTypeUnknown
)

const (
	IntentNameOpenSession           = "openSession"
	IntentNameCloseSession          = "closeSession"
	IntentNameValidateSession       = "validateSession"
	IntentNameFinishValidateSession = "finishValidateSession"
	IntentNameListSessions          = "listSessions"
	IntentNameGetSession            = "getSession"
	IntentNameSignMessage           = "signMessage"
	IntentNameSendTransaction       = "sendTransaction"
)

const (
	IntentResponseCodeSessionOpened      = "sessionOpened"
	IntentResponseCodeValidationRequired = "validationRequired"
	IntentResponseCodeValidationStarted  = "validationStarted"
	IntentResponseCodeValidationFinished = "validationFinished"
	IntentResponseCodeSignedMessage      = "signedMessage"
	IntentResponseCodeTransactionReceipt = "transactionReceipt"
	IntentResponseCodeTransactionFailed  = "transactionFailed"
	IntentResponseCodeGetSessionResponse = "getSessionResponse"
)

type IntentDataValidator interface {
	IsValid() error
}

func (intent *Intent) Hash() ([]byte, error) {
	// copy intent and remove signatures
	var intentCopy = *intent
	intentCopy.Signatures = nil

	// Convert packet to bytes
	packetBytes, err := canonicaljson.Marshal(intentCopy)
	if err != nil {
		return nil, err
	}

	// Calculate keccak256 hash
	return crypto.Keccak256(packetBytes), nil
}

func (intent *Intent) IsValid() error {
	if len(intent.Signatures) == 0 {
		return fmt.Errorf("no signatures")
	}

	// check if the intent is expired
	if intent.ExpiresAt+IntentAllowedTimeDriftInSec < uint64(time.Now().Unix()) {
		return fmt.Errorf("intent expired")
	}

	// check if the intent is issued in the future
	if intent.IssuedAt-IntentAllowedTimeDriftInSec > uint64(time.Now().Unix()) {
		return fmt.Errorf("intent issued in the future")
	}

	// check if at least one signature is valid
	if len(intent.Signers()) == 0 {
		return fmt.Errorf("invalid signature")
	}

	// the intent is valid
	return nil
}

func (intent *Intent) keyType(sessionId string) KeyType {
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

func (intent *Intent) isValidSignature(sessionId string, signature string) error {
	switch intent.keyType(sessionId) {
	case KeyTypeSECP256K1:
		return intent.isValidSignatureP256K1(sessionId, signature)
	case KeyTypeSECP256R1:
		return intent.isValidSignatureP256R1(sessionId, signature)
	default:
		return fmt.Errorf("unknown session key type")
	}
}

// isValidSignatureP256K1 checks if the signature is valid for the given secp256k1 session
func (intent *Intent) isValidSignatureP256K1(sessionId string, signature string) error {
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

// isValidSignatureP256R1 checks if the signature is valid for the given secp256r1 session
func (intent *Intent) isValidSignatureP256R1(sessionId string, signature string) error {
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

func (intent *Intent) Signers() []string {
	var signers []string
	for _, signature := range intent.Signatures {
		if err := intent.isValidSignature(signature.SessionId, signature.Signature); err == nil {
			signers = append(signers, signature.SessionId)
		}
	}
	return signers
}
