package prototyp

import (
	"database/sql/driver"
	"encoding/hex"
)

// HashMaybe is a nullable Hash value useful for database fields which accept NULL type.
type HashMaybe struct {
	Hash

	// IsAssigned=false means value is nil. IsAssigned=true means value in .Hash is the value.
	IsAssigned bool
}

func HashMaybeFromString(s string) HashMaybe {
	if s == "" {
		return HashMaybe{}
	} else {
		return HashMaybe{
			Hash:       HashFromString(s),
			IsAssigned: true,
		}
	}
}

// UnmarshalText implements encoding.TextMarshaler.
func (h HashMaybe) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (h *HashMaybe) UnmarshalText(src []byte) error {
	t := string(src)
	if t == "null" || t == "" {
		return nil
	}
	*h = HashMaybeFromString(t)
	return nil
}

func (h HashMaybe) Value() (driver.Value, error) {
	if h.IsAssigned {
		s := h.String()
		if len(s) < 2 {
			return []byte{}, nil
		}
		return hex.DecodeString(s[2:])
	}
	return nil, nil
}

func (h *HashMaybe) Scan(src interface{}) error {
	h.IsAssigned = false
	if src == nil {
		return nil
	}

	h.Hash = BytesToHash(src.([]byte))
	h.IsAssigned = true
	return nil
}

func ToHashMaybe(h Hexer) HashMaybe {
	return HashMaybeFromString(h.Hex())
}

func BytesToHashMaybe(src []byte) HashMaybe {
	return HashMaybeFromString("0x" + hex.EncodeToString(src))
}
