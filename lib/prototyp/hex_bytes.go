package prototyp

import (
	"encoding/hex"
	"fmt"
)

type HexBytes []byte

// MarshalText implements encoding.TextMarshaler
func (b HexBytes) MarshalText() ([]byte, error) {
	result := make([]byte, len(b)*2+2)
	copy(result, `0x`)
	hex.Encode(result[2:], b)
	return result, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *HexBytes) UnmarshalJSON(input []byte) error {
	if !isString(input) {
		return fmt.Errorf("dbtype: HexBytes UnmarshalJSON received non-string input")
	}
	err := b.UnmarshalText(input[1 : len(input)-1])
	if err != nil {
		return fmt.Errorf("dbtype: HexBytes UnmarshalJSON: %w", err)
	}
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *HexBytes) UnmarshalText(input []byte) error {
	raw, err := checkText(input, true)
	if err != nil {
		return err
	}
	dec := make([]byte, len(raw)/2)
	if _, err = hex.Decode(dec, raw); err != nil {
		err = fmt.Errorf("dbtype: HexBytes UnmarshalText failed: %w", err)
	} else {
		*b = dec
	}
	return err
}

// String returns encoded b as a hex string with 0x prefix.
func (b HexBytes) String() string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

// ScanBytes implements pgx/pgtype.BytesScanner
func (b *HexBytes) ScanBytes(v []byte) error {
	*b = v
	return nil
}

// BytesValue implements pgx/pgtype.BytesValuer
func (b HexBytes) BytesValue() ([]byte, error) {
	return b, nil
}

func isString(input []byte) bool {
	return len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"'
}

func bytesHave0xPrefix(input []byte) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func checkText(input []byte, wantPrefix bool) ([]byte, error) {
	if len(input) == 0 {
		return nil, nil // empty strings are allowed
	}
	if bytesHave0xPrefix(input) {
		input = input[2:]
	} else if wantPrefix {
		return nil, fmt.Errorf("dbtype: HexBytes, hex string without 0x prefix")
	}
	if len(input)%2 != 0 {
		return nil, fmt.Errorf("dbtype: HexBytes, hex string of odd length")
	}
	return input, nil
}
