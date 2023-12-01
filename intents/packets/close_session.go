package packets

import (
	"encoding/json"
	"fmt"
)

type CloseSessionPacket struct {
	BasePacketForWallet
	Session string `json:"session"`
}

const CloseSessionPacketCode = "closeSession"

func (p *CloseSessionPacket) Unmarshal(packet json.RawMessage) error {
	err := json.Unmarshal(packet, &p)
	if err != nil {
		return err
	}

	if p.Code != SignMessagePacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", SignMessagePacketCode, p.Code)
	}

	return nil
}
