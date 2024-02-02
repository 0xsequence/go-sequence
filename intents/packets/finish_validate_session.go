package packets

import (
	"encoding/json"
	"fmt"
)

type FinishValidateSessionPacket struct {
	Code      string `json:"code"`
	SessionId string `json:"sessionId"`
	Salt      string `json:"salt"`
	Challenge string `json:"challenge"`
}

const FinishValidateSessionPacketCode = "finishValidateSession"

func (p *FinishValidateSessionPacket) Unmarshal(packet json.RawMessage) error {
	err := json.Unmarshal(packet, &p)
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
