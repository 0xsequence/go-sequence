package packets

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/go-sequence/intents"
)

type OpenSessionPacket struct {
	BasePacket
	Session string                 `json:"session"`
	Proof   OpenSessionPacketProof `json:"proof"`
}

type OpenSessionPacketProof struct {
	IDToken string `json:"idToken"`
	Email   string `json:"email"`
}

const OpenSessionPacketCode = "openSession"

func (p *OpenSessionPacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
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
