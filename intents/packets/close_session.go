package packets

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/go-sequence/intents"
)

type CloseSessionPacket struct {
	Code    string `json:"code"`
	Wallet  string `json:"wallet"`
	Session string `json:"session"`
}

const CloseSessionPacketCode = "closeSession"

func (p *CloseSessionPacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
	if err != nil {
		return err
	}

	if p.Code != SignMessagePacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", SignMessagePacketCode, p.Code)
	}

	return nil
}
