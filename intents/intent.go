package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xsequence/go-sequence/intents/packets"
	"github.com/gibson042/canonicaljson-go"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

var (
	ErrInvalidPacket    = fmt.Errorf("invalid packet")
	ErrNoSignatures     = fmt.Errorf("no signatures")
	ErrInvalidSignature = fmt.Errorf("invalid signature")
)

type KeyType int

const (
	KeyTypeSECP256K1 KeyType = iota
	KeyTypeSECP256R1
	KeyTypeUnknown
)

type Intent struct {
	Version string          `json:"version"`
	Packet  json.RawMessage `json:"packet"`

	// signatures is made private to prevent it from being accidentally used raw
	// by the user. The Signers() method should be used instead, it will return
	// *only* the valid signatures.
	signatures []Signature
}

type JSONIntent struct {
	Version    string          `json:"version"`
	Packet     json.RawMessage `json:"packet"`
	Signatures []Signature     `json:"signatures"`
}

type Signature struct {
	SessionId string `json:"sessionId"`
	Signature string `json:"signature"`
}

func (i *Intent) UnmarshalJSON(data []byte) error {
	var intent JSONIntent
	err := json.Unmarshal(data, &intent)
	if err != nil {
		return err
	}
	i.Version = intent.Version
	i.Packet = intent.Packet
	i.signatures = intent.Signatures
	return nil
}

func (i *Intent) MarshalJSON() ([]byte, error) {
	return json.Marshal(JSONIntent{
		Version:    i.Version,
		Packet:     i.Packet,
		Signatures: i.signatures,
	})
}

func (intent *Intent) Hash() ([]byte, error) {
	packet := intent.Packet

	// Convert packet to bytes
	packetBytes, err := canonicaljson.Marshal(packet)
	if err != nil {
		return nil, err
	}

	// Calculate keccak256 hash
	return crypto.Keccak256(packetBytes), nil
}

func (intent *Intent) Signers() []string {
	var signers []string

	for _, signature := range intent.signatures {
		if intent.isValidSignature(signature.SessionId, signature.Signature) {
			signers = append(signers, signature.SessionId)
		}
	}

	return signers
}

func (intent *Intent) IsValid() (bool, error) {
	// Check if the packet is valid
	var packet packets.BasePacket
	err := json.Unmarshal(intent.Packet, &packet)
	if err != nil {
		return false, fmt.Errorf("intent: %w", ErrInvalidPacket)
	}

	// OpenSession packets do not require signatures
	if packet.Code == packets.OpenSessionPacketCode {
		if ok, err := packet.IsValid(); !ok {
			return false, fmt.Errorf("intent: %w", err)
		}
		return true, nil
	}

	// Check if there are any signatures
	if len(intent.signatures) == 0 {
		return false, fmt.Errorf("intent: %w", ErrNoSignatures)
	}

	// Check if all signatures are valid
	for _, signature := range intent.signatures {
		if !intent.isValidSignature(signature.SessionId, signature.Signature) {
			return false, fmt.Errorf("intent: %w", ErrInvalidSignature)
		}
	}

	// Check if the packet is valid
	if ok, err := packet.IsValid(); !ok {
		return false, fmt.Errorf("intent: %w", err)
	}
	return true, nil
}

func (intent *Intent) keyType(sessionId string) KeyType {
	// handle empty session ids
	if len(sessionId) == 0 {
		return KeyTypeUnknown
	}

	// handle old session ids
	if len(sessionId) <= 42 {
		return KeyTypeSECP256K1
	}

	// handle key typed session ids
	sessionIdBytes := common.FromHex(sessionId)
	return KeyType(sessionIdBytes[0])
}

func (intent *Intent) isValidSignature(sessionId string, signature string) bool {
	switch intent.keyType(sessionId) {
	case KeyTypeSECP256K1:
		return intent.isValidSignatureP256K1(sessionId, signature)
	case KeyTypeSECP256R1:
		return intent.isValidSignatureP256R1(sessionId, signature)
	default:
		return false
	}
}

// isValidSignatureP256K1 checks if the signature is valid for the given secp256k1 session
func (intent *Intent) isValidSignatureP256K1(sessionAddress string, signature string) bool {
	// handle typed session address
	if len(sessionAddress) > 42 {
		sessionAddressBytes := common.FromHex(sessionAddress)
		sessionAddress = fmt.Sprintf("0x%s", common.Bytes2Hex(sessionAddressBytes[1:]))
	}

	// Get hash of the packet
	hash, err := intent.Hash()
	if err != nil {
		return false
	}

	// Add Ethereum prefix to the hash
	prefixedHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)))

	// Convert the signature to bytes
	sigBytes := common.FromHex(signature)
	if sigBytes[64] == 27 || sigBytes[64] == 28 {
		sigBytes[64] -= 27
	}

	// Recover the public key from the signature
	pubKey, err := crypto.Ecrecover(prefixedHash.Bytes(), sigBytes)
	if err != nil {
		return false
	}

	addr := common.BytesToAddress(crypto.Keccak256(pubKey[1:])[12:])

	// Check if the recovered address matches the session address
	return strings.ToLower(addr.Hex()) == strings.ToLower(sessionAddress)
}

// isValidSignatureP256R1 checks if the signature is valid for the given secp256r1 session
func (intent *Intent) isValidSignatureP256R1(publicKey string, signature string) bool {
	publicKeyBuff := common.FromHex(publicKey)[1:]

	// public key
	// TODO: check if can use ecdh instead of unmarshal
	// NOTE: no way to convert ecdh pub key into elliptic pub key?
	x, y := elliptic.Unmarshal(elliptic.P256(), publicKeyBuff)
	if x == nil || y == nil {
		return false
	}

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// message hash
	messageHash, _ := intent.Hash()
	messageHash2 := sha256.Sum256(messageHash)

	// signature
	signatureBytes := common.FromHex(signature)
	if len(signatureBytes) != 64 {
		return false
	}

	r := new(big.Int).SetBytes(signatureBytes[:32])
	s := new(big.Int).SetBytes(signatureBytes[32:64])
	return ecdsa.Verify(&pub, messageHash2[:], r, s)
}

func (intent *Intent) PacketCode() string {
	var packetCode struct {
		Code string `json:"code"`
	}
	json.Unmarshal(intent.Packet, &packetCode)
	return packetCode.Code
}
