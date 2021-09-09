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
// NOTE: when used with a db like postgres, the column type must be a `varchar`,
// `text` or string equivalent.
type HashText string

func HashTextFromString(s string) HashText {
	return HashText(strings.ToLower(s))
}

func HashTextFromBytes(src []byte) HashText {
	return HashText("0x" + hex.EncodeToString(src))
}

func ToHashText(h Hexer) HashText {
	return HashTextFromString(h.Hex())
}

func (h HashText) ToAddress() common.Address {
	return common.HexToAddress(string(h))
}

func (h HashText) ToHash() common.Hash {
	return common.HexToHash(string(h))
}

// UnmarshalText implements encoding.TextMarshaler.
func (h *HashText) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (h *HashText) UnmarshalText(src []byte) error {
	*h = HashTextFromString(string(src))
	return nil
}

func (h HashText) String() string {
	return string(h)
}

func (h HashText) IsZeroValue() bool {
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

func (h HashText) IsValidAddress() bool {
	if h[0:2] != "0x" {
		return false
	}
	if len(h) != 42 {
		return false
	}
	return true
}

func (h HashText) IsValidTxnHash() bool {
	if h[0:2] != "0x" {
		return false
	}
	if len(h) != 66 {
		return false
	}
	return true
}

func (h *HashText) Hash() common.Hash {
	return common.HexToHash(h.String())
}

func (h HashText) Value() (driver.Value, error) {
	return h.String(), nil
}

func (h *HashText) Scan(src interface{}) error {
	// NOTE: the 'scany' package we use is unable to scan values of
	// *string, aka *prototyp.Hash, when needing to have a nullable Hash
	// please use the HashMaybe type instead.
	*h = HashText(src.(string))
	return nil
}
