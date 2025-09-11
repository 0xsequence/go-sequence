package binaryutil

import (
	"bytes"
	"math/big"
	"testing"
)

func TestWriteBigInt(t *testing.T) {
	tests := []struct {
		name     string
		input    *big.Int
		expected []byte
	}{
		{
			name:     "positive",
			input:    big.NewInt(12345),
			expected: append([]byte{0x01}, big.NewInt(12345).Bytes()...),
		},
		{
			name:     "negative",
			input:    big.NewInt(-12345),
			expected: append([]byte{0xFF}, big.NewInt(12345).Bytes()...),
		},
		{
			name:     "zero",
			input:    big.NewInt(0),
			expected: []byte{0x00},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteBigInt(&buf, tt.input)
			if err != nil {
				t.Fatalf("WriteBigInt() error = %v", err)
			}
			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("WriteBigInt() = %v, want %v", buf.Bytes(), tt.expected)
			}
		})
	}
}

func TestWriteOptionalBigInt(t *testing.T) {
	tests := []struct {
		name     string
		input    *big.Int
		expected []byte
	}{
		{
			name:     "nil",
			input:    nil,
			expected: []byte{0x00},
		},
		{
			name:     "positive",
			input:    big.NewInt(12345),
			expected: append([]byte{0x01}, append([]byte{0x01}, big.NewInt(12345).Bytes()...)...),
		},
		{
			name:     "negative",
			input:    big.NewInt(-12345),
			expected: append([]byte{0x01}, append([]byte{0xFF}, big.NewInt(12345).Bytes()...)...),
		},
		{
			name:     "zero",
			input:    big.NewInt(0),
			expected: append([]byte{0x01}, []byte{0x00}...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteOptionalBigInt(&buf, tt.input)
			if err != nil {
				t.Fatalf("WriteOptionalBigInt() error = %v", err)
			}
			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("WriteOptionalBigInt() = %v, want %v", buf.Bytes(), tt.expected)
			}
		})
	}
}
