package intents

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

func (id *IntentDataOpenSession) IsValid() error {
	if id.SessionID == "" {
		return fmt.Errorf("session id is empty")
	}

	return nil
}

func (id *IntentDataSessionAuthProof) IsValidInterpretation(sessionID string, message string) error {
	message2 := "0x" + common.Bytes2Hex(
		[]byte(SessionAuthProofMessage(sessionID, id.Wallet, id.Nonce)),
	)
	if message != message2 {
		return fmt.Errorf("proof message does not match: '%s' != '%s'", message, message2)
	}
	return nil
}
