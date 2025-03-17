package prototyp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
