package prototyp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashMaybe(t *testing.T) {
	hm := HashMaybeFromString("")

	hm.SetValue(HashFromString("0xf0245f6251bef9447a08766b9da2b07b28ad80b0"))
	assert.True(t, hm.IsValidAddress())
	assert.True(t, hm.IsAssigned)

	hm.SetNil()
	assert.False(t, hm.IsValidAddress())
	assert.False(t, hm.IsAssigned)
}
