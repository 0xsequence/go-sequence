package binaryutil

import (
	"bytes"
	"testing"
)

func TestWriteOptionalUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    *uint64
		expected []byte
	}{
		{
			name:     "nil",
			input:    nil,
			expected: []byte{0x00},
		},
		{
			name:     "present",
			input:    func() *uint64 { v := uint64(12345); return &v }(),
			expected: []byte{0x01, 0x39, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := WriteOptionalUint64(&buf, tt.input)
			if err != nil {
				t.Fatalf("WriteOptionalUint64() error = %v", err)
			}
			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("WriteOptionalUint64() = %v, want %v", buf.Bytes(), tt.expected)
			}
		})
	}
}
