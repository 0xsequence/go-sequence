package prototyp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashIsValidAddress(t *testing.T) {
	// IsValidAddress accepts exactly the values ToAddress decodes losslessly.
	cases := []struct {
		in   Hash
		want bool
	}{
		{Hash("0x2222222222222222222222222222222222222222"), true},
		{Hash("0xAbCdEf" + strings.Repeat("0", 32) + "01"), true},   // mixed-case hex decodes fine
		{Hash("0x" + strings.Repeat("z", 40)), false},               // 42 chars, but not hex
		{Hash("0x22222222222222222222222222222222222222zz"), false}, // trailing non-hex would be silently truncated by ToAddress
		{Hash("0x" + strings.Repeat("2", 41)), false},               // 21 bytes
		{Hash("0x22"), false},
		{Hash("0x"), false},
		{Hash(""), false},
		{Hash(strings.Repeat("2", 40)), false},                      // missing 0x prefix
		{Hash("0X2222222222222222222222222222222222222222"), false}, // uppercase prefix stays rejected
	}
	for _, c := range cases {
		require.Equal(t, c.want, c.in.IsValidAddress(), c.in)
	}
}

func TestHashIsValidTxnHash(t *testing.T) {
	// IsValidTxnHash accepts exactly the values ToHash decodes losslessly.
	cases := []struct {
		in   Hash
		want bool
	}{
		{Hash("0x" + strings.Repeat("ab", 32)), true},
		{Hash("0x" + strings.Repeat("Ab", 32)), true},        // mixed-case hex decodes fine
		{Hash("0x" + strings.Repeat("z", 64)), false},        // 66 chars, but not hex
		{Hash("0x" + strings.Repeat("2", 63) + "zz"), false}, // trailing non-hex would be silently truncated by ToHash
		{Hash("0x" + strings.Repeat("2", 66)), false},        // 33 bytes
		{Hash("0x22"), false},
		{Hash("0x"), false},
		{Hash(""), false},
		{Hash("0"), false}, // used to panic on h[0:2]
		{Hash("0X" + strings.Repeat("2", 64)), false},
	}
	for _, c := range cases {
		require.Equal(t, c.want, c.in.IsValidTxnHash(), c.in)
	}
}

func TestHexStringToBytes(t *testing.T) {
	hexString := "0x1234567890abcdef"
	bytes := HexStringToBytes(hexString)
	require.Equal(t, []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}, bytes)
	require.Len(t, bytes, 8)

	hexString2 := HexBytesToString(bytes)
	require.Equal(t, hexString, hexString2)
}

func TestHash(t *testing.T) {
	hash := Hash("0x1234567890abcdef")
	require.Equal(t, []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}, hash.Bytes())
	require.Len(t, hash.Bytes(), 8)
	require.Equal(t, 8, hash.ByteSize())
	require.True(t, hash.IsValidHex())
}

func TestHashBinaryMarshaler(t *testing.T) {
	hash := Hash("0x1234567890abcdef")
	bytes, err := hash.MarshalBinary()
	require.NoError(t, err)
	require.Equal(t, []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}, bytes)

	hash2 := Hash("")
	err = hash2.UnmarshalBinary(bytes)
	require.NoError(t, err)
	require.Equal(t, hash, hash2)

	require.Equal(t, hash.String(), "0x1234567890abcdef")
}
