package packets

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/go-sequence/intents"
)

type ValidateSessionPacket struct {
	Code    string `json:"code"`
	Wallet  string `json:"wallet"`
	Session string `json:"session"`
}

const ValidateSessionPacketCode = "validateSession"

func (p *ValidateSessionPacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
	if err != nil {
		return err
	}

	if p.Code != ValidateSessionPacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", SignMessagePacketCode, p.Code)
	}

	return nil
}

const ValidateSessionPacketResponseCode = "validateSessionStarted"

type ValidateSessionPacketResponse struct {
	Code string `json:"code"`

	Data struct {
		EmailSent bool `json:"emailSent"`
		SmsSent   bool `json:"smsSent"`
	}
}