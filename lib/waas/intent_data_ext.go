package intents

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
)

func (id *IntentDataOpenSession) IsValid() error {
	if id.SessionID == "" {
		return fmt.Errorf("session id is empty")
	}

	return nil
}

func (id *IntentDataSessionAuthProof) IsValidInterpretation(sessionID string, message string) error {
	expected := hexutil.Encode([]byte(SessionAuthProofMessage(sessionID, id.Wallet, id.Nonce)))
	if message != expected {
		return fmt.Errorf("proof message does not match: '%s' != '%s'", message, expected)
	}
	return nil
}
