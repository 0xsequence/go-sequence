package packets

import (
	"encoding/json"
	"fmt"
)

type OpenSessionPacket struct {
	BasePacket
	SessionId string                 `json:"sessionId"`
	Proof     OpenSessionPacketProof `json:"proof"`
}

type OpenSessionPacketProof struct {
	IDToken string `json:"idToken"`
	Email   string `json:"email"`
}

const OpenSessionPacketCode = "openSession"

func (p *OpenSessionPacket) Unmarshal(packet json.RawMessage) error {
	err := json.Unmarshal(packet, &p)
	if err != nil {
		return err
	}

	if p.Code != OpenSessionPacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", OpenSessionPacketCode, p.Code)
	}

	return nil
}

func (p *OpenSessionPacket) HasEmail() bool {
	return p.Proof.Email != ""
}

func (p *OpenSessionPacket) HasIDToken() bool {
	return p.Proof.IDToken != ""
}
