package prototyp

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
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

func (h *HashMaybe) SetValue(hash Hash) {
	h.Hash = hash
	h.IsAssigned = true
}

func (h *HashMaybe) SetNil() {
	h.Hash = ""
	h.IsAssigned = false
}

// MarshalText implements encoding.TextMarshaler.
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

// MarshalBinary implements encoding.BinaryMarshaler.
func (h HashMaybe) MarshalBinary() ([]byte, error) {
	if h.IsAssigned {
		return h.Hash.MarshalBinary()
	} else {
		return []byte{}, nil
	}
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (h *HashMaybe) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		h.IsAssigned = false
		return nil
	} else {
		h.IsAssigned = true
		h.Hash = Hash(HexBytesToString(b))
		return nil
	}
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
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	h.Hash = HashFromBytes(b)
	h.IsAssigned = true
	return nil
}

func (h HashMaybe) Len() int {
	return len(h.Hash.String())
}

func (h HashMaybe) ByteSize() int {
	if h.IsAssigned {
		return h.Hash.ByteSize()
	}
	return 0
}

func ToHashMaybe(h Hexer) HashMaybe {
	return HashMaybeFromString(h.Hex())
}

func BytesToHashMaybe(src []byte) HashMaybe {
	return HashMaybeFromString("0x" + hex.EncodeToString(src))
}
