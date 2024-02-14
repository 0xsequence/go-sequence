package intents

import "fmt"

func (id *IntentDataOpenSession) IsValid() error {
	if id.SessionID == "" {
		return fmt.Errorf("session id is empty")
	}

	if id.IdToken == nil && id.Email == nil {
		return fmt.Errorf("idToken and email are both nil")
	}
	return nil
}
