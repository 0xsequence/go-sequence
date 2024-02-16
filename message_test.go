package sequence_test

import (
	"testing"

	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
)

func TestIsEIP191Message(t *testing.T) {
	t.Run("EIP191", func(t *testing.T) {
		// EIP191 message
		msg := []byte("\x19Ethereum Signed Message:\n5hello")
		ok := sequence.IsEIP191Message(msg)
		if !ok {
			t.Error("expected EIP191 message")
		}
	})

	t.Run("non-EIP191", func(t *testing.T) {
		// non-EIP191 message
		msg := []byte("hello")
		ok := sequence.IsEIP191Message(msg)
		if ok {
			t.Error("expected non-EIP191 message")
		}
	})
}

func TestMessageToEIP191(t *testing.T) {
	t.Run("EIP191_none", func(t *testing.T) {
		// EIP191 message
		msg := []byte("hello")
		expectedEIP191 := []byte("\x19Ethereum Signed Message:\n5hello")

		eip191 := sequence.MessageToEIP191(msg)
		assert.Equal(t, eip191, expectedEIP191)
	})

	t.Run("EIP191_already", func(t *testing.T) {
		// EIP191 message
		msg := []byte("\x19Ethereum Signed Message:\n5hello")
		eip191 := sequence.MessageToEIP191(msg)

		assert.Equal(t, msg, eip191)
	})
}
