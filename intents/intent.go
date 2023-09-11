package intents

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

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

func (intent *Intent) Hash() ([]byte, error) {
	packet := intent.Packet

	// Convert packet to bytes
	packetBytes, err := json.Marshal(packet)
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

func (intent *Intent) isValidSignature(session string, signature string) bool {
	// Get hash of the packet
	hash, err := intent.Hash()
	if err != nil {
		return false
	}

	// Add Ethereum prefix to the hash
	prefixedHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)))

	// Convert the signature to bytes
	sigBytes, err := hex.DecodeString(strings.TrimPrefix(signature, "0x"))
	if err != nil {
		return false
	}

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
