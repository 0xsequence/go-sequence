package intents

import (
	"testing"

	"github.com/0xsequence/ethkit"
	"github.com/stretchr/testify/assert"
)

func TestSessionAuthProofMessage(t *testing.T) {
	// no nonce provided
	assert.Equal(t, "SessionAuthProof r1:0x01 0x02", SessionAuthProofMessage("r1:0x01", "0x02", nil))

	// nonce provided
	assert.Equal(t, "SessionAuthProof r1:0x01 0x02 0x03", SessionAuthProofMessage("r1:0x01", "0x02", ethkit.ToPtr("0x03")))
}
