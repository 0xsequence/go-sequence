package intents

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xsequence/go-sequence/intents/packets"
	"github.com/gibson042/canonicaljson-go"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
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
	Session   string `json:"session"`
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
		if intent.isValidSignature(signature.Session, signature.Signature) {
			signers = append(signers, signature.Session)
		}
	}

	return signers
}

func (intent *Intent) IsValid() bool {
	// Check if the packet is valid
	var packet packets.BasePacket
	err := json.Unmarshal(intent.Packet, &packet)
	if err != nil {
		return false
	}

	// OpenSession packets do not require signatures
	if packet.Code == packets.OpenSessionPacketCode {
		return packet.IsValid()
	}

	// Check if there are any signatures
	if len(intent.signatures) == 0 {
		return false
	}

	// Check if all signatures are valid
	for _, signature := range intent.signatures {
		if !intent.isValidSignature(signature.Session, signature.Signature) {
			return false
		}
	}

	// Check if the packet is valid
	return packet.IsValid()
}

func (intent *Intent) isValidSignature(session string, signature string) bool {
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
	return strings.ToLower(addr.Hex()) == strings.ToLower(session)
}

func (intent *Intent) PacketCode() string {
	var packetCode struct {
		Code string `json:"code"`
	}
	json.Unmarshal(intent.Packet, &packetCode)
	return packetCode.Code
}
