package prototyp

import (
	"database/sql/driver"
	"encoding"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// Hash is a type alias for common.Hash used for data normalization
// with JSON/Database marshalling. Hash is expected to be a hex string.
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

func (h Hash) ToShortHash() Hash {
	// in case hash is shorter then we expect, return just its value
	if len(h) < 10 {
		return h
	}
	// return short-hash if already is len(0x12345678) == 10
	if len(h) == 10 {
		return h
	}
	// return short-hash as 0x plus first 8 bytes of the blockHash,
	// which assumes the blockHash will always have 0x prefix.
	return h[0:10]
}

var (
	_h                            = Hash("")
	_  encoding.BinaryMarshaler   = _h
	_  encoding.BinaryUnmarshaler = &_h
	_  encoding.TextMarshaler     = _h
	_  encoding.TextUnmarshaler   = &_h
)

// UnmarshalText implements encoding.TextMarshaler.
func (h Hash) MarshalText() ([]byte, error) {
	return []byte(h.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (h *Hash) UnmarshalText(src []byte) error {
	*h = HashFromString(string(src))
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (h Hash) MarshalBinary() ([]byte, error) {
	return HexStringToBytes(string(h)), nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (h *Hash) UnmarshalBinary(b []byte) error {
	*h = Hash(HexBytesToString(b))
	return nil
}

func (h Hash) String() string {
	return string(h)
}

func (h Hash) Hex() string {
	return string(h)
}

func (h Hash) ToLower() Hash {
	return Hash(strings.ToLower(string(h)))
}

func (h Hash) Len() int {
	return len(h.String())
}

func (h Hash) ByteSize() int {
	if len(h) == 0 {
		return 0
	}
	if has0xPrefix(string(h)) {
		return (len(h) - 2) / 2
	} else {
		return len(h) / 2
	}
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

func (h Hash) IsEmpty() bool {
	return h == ""
}

// IsNativeTokenAddress checks if the hash value is the zero address (0x000..000) or
// ERC-7528 (https://eips.ethereum.org/EIPS/eip-7528) native token address represented
// as a contract address.
//
// For some background, many use 0x000..000 or 0xeee..eee to represent a native token
// like ETH in the context of erc20 tokens.
func (h Hash) IsNativeTokenAddress() bool {
	return h == "0x0000000000000000000000000000000000000000" ||
		h == "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE" ||
		h == "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
}

func (h Hash) IsValidHex() bool {
	if len(h) == 0 {
		return false
	}

	s := string(h)
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s) == 0 {
		return true // its valid, but its empty
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}

	return true
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

func (h Hash) Bytes() []byte {
	return HexStringToBytes(string(h))
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
	if src == nil {
		return nil
	}
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}
	*h = HashFromBytes(b)
	return nil
}

func ToHashList[T Hexer](list []T) []Hash {
	result := make([]Hash, 0, len(list))
	for _, a := range list {
		result = append(result, ToHash(a))
	}
	return result
}

func HexBytesToString(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func HexStringToBytes(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	h, _ := hex.DecodeString(s)
	return h
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
