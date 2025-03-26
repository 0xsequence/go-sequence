package prototyp

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

var (
	EmptyKey      = [16]byte{}
	EmptyKeySlice = EmptyKey[:]
)

type Key [16]byte

func (k Key) String() string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(k[:]))
}

func (k Key) IsValid() bool {
	return k != EmptyKey
}

func (k Key) IsZeroValue() bool {
	return k == EmptyKey
}

func (k Key) Bytes() []byte {
	return k[:]
}

func (k *Key) Scan(src interface{}) error {
	copy(k[:16], src.([]byte))
	return nil
}

func (k Key) Value() (driver.Value, error) {
	return k[:], nil
}

func (k *Key) Len() int {
	return 16
}

func (k *Key) ByteSize() int {
	return 16
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (k Key) MarshalBinary() ([]byte, error) {
	return k[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (k *Key) UnmarshalBinary(b []byte) error {
	copy(k[:], b)
	return nil
}
