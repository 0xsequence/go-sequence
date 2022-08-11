package prototyp

import (
	"database/sql/driver"
	"encoding/hex"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// Hash is a type alias for common.Hash used for data normalization
// with JSON/Database marshalling.
//
// NOTE: when used with a db like postgres, the column type must be a `bytea`
type Hash string

func HashFromString(s string) Hash {
	return Hash(strings.ToLower(s))
}

func HashFromBytes(src []byte) Hash {
	return Hash("0x" + hex.EncodeToString(src))
}

type Hexer interface {
	Hex() string
}

func ToHash(h Hexer) Hash {
	return HashFromString(h.Hex())
}

func (h Hash) ToAddress() common.Address {
	return common.HexToAddress(string(h))
}

func (h Hash) ToHash() common.Hash {
	return common.HexToHash(string(h))
}

// UnmarshalText implements encoding.TextMarshaler.
func (h *Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (h *Hash) UnmarshalText(src []byte) error {
	*h = HashFromString(string(src))
	return nil
}

func (h *Hash) ExtensionType() int8 {
	return 10
}

func (h *Hash) Len() int {
	return len(h.String())
}

func (h *Hash) MarshalBinaryTo(b []byte) error {
	copy(b[:], h.String())
	return nil
}

func (h *Hash) UnmarshalBinary(b []byte) error {
	*h = HashFromString(string(b))
	return nil
}

func (h Hash) String() string {
	return string(h)
}

func (h Hash) IsZeroValue() bool {
	if h.String() == "" {
		return true
	}
	if h.String() == "0x" {
		return true
	}
	if h.String() == "0x00000000" {
		return true
	}
	if h.String() == "0x0000000000000000000000000000000000000000" {
		return true
	}
	if h.String() == "0x0000000000000000000000000000000000000000000000000000000000000000" {
		return true
	}
	return false
}

func (h Hash) IsValidAddress() bool {
	if len(h) <= 2 || h[0:2] != "0x" {
		return false
	}
	if len(h) != 42 {
		return false
	}
	return true
}

func (h Hash) IsValidTxnHash() bool {
	if h[0:2] != "0x" {
		return false
	}
	if len(h) != 66 {
		return false
	}
	return true
}

func (h *Hash) Hash() common.Hash {
	return common.HexToHash(h.String())
}

func (h Hash) Value() (driver.Value, error) {
	s := h.String()
	if len(s) < 2 {
		return []byte{}, nil
	}
	return hex.DecodeString(s[2:])
}

func (h *Hash) Scan(src interface{}) error {
	// NOTE: the 'scany' package we use is unable to scan values of
	// *string, aka *prototyp.Hash, when needing to have a nullable Hash
	// please use the HashMaybe type instead.
	*h = HashFromBytes(src.([]byte))
	return nil
}
