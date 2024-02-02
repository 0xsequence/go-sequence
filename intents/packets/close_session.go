package packets

import (
	"encoding/json"
	"fmt"
)

type CloseSessionPacket struct {
	BasePacketForWallet
	SessionId string `json:"sessionId"`
}

const CloseSessionPacketCode = "closeSession"

func (p *CloseSessionPacket) Unmarshal(packet json.RawMessage) error {
	err := json.Unmarshal(packet, &p)
	if err != nil {
		return err
	}

	if p.Code != CloseSessionPacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", CloseSessionPacketCode, p.Code)
	}

	return nil
}
