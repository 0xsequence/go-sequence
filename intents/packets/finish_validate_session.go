package packets

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/go-sequence/intents"
)

type FinishValidateSessionPacket struct {
	Code      string `json:"code"`
	Session   string `json:"session"`
	Salt      string `json:"salt"`
	Challenge string `json:"challenge"`
}

const FinishValidateSessionPacketCode = "finishValidateSession"

func (p *FinishValidateSessionPacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
	if err != nil {
		return err
	}

	if p.Code != FinishValidateSessionPacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", FinishValidateSessionPacketCode, p.Code)
	}

	return nil
}

const FinishValidateSessionPacketResponseCode = "finishedSessionValidation"

type FinishedSessionValidationPacketResponse struct {
	Code string `json:"code"`

	Data struct {
		IsValid bool `json:"isValid"`
	}
}
