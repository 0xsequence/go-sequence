package intents

import (
	"testing"

	"github.com/0xsequence/ethkit"
	"github.com/stretchr/testify/assert"
)

func TestOpenSessionProofMessage(t *testing.T) {
	// no nonce provided
	assert.Equal(t, "OpenedSession r1:0x01 0x02", OpenSessionProofMessage("r1:0x01", "0x02", nil))

	// nonce provided
	assert.Equal(t, "OpenedSession r1:0x01 0x02 0x03", OpenSessionProofMessage("r1:0x01", "0x02", ethkit.ToPtr("0x03")))
}
