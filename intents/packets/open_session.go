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

const OpenSessionPacketResponseCode = "sessionOpened"

type OpenSessionPacketResponse struct {
	Code string                        `json:"code"`
	Data OpenSessionPacketResponseData `json:"data"`
}

type OpenSessionPacketResponseData struct {
	SessionId string `json:"sessionId"`
	Wallet    string `json:"wallet"`
}

func (p *OpenSessionPacketResponse) UnmarshalFromAny(a any) error {
	m, ok := a.(map[string]any)
	if !ok {
		return fmt.Errorf("packet is not a map")
	}

	if code, ok := m["code"].(string); ok && code == OpenSessionPacketResponseCode {
		p.Code = code
	} else {
		return fmt.Errorf("packet code is not '%s', got '%s'", OpenSessionPacketResponseCode, m["code"])
	}

	data, ok := m["data"].(map[string]any)
	if !ok {
		return fmt.Errorf("packet data is not a map")
	}

	if sessionId, ok := data["sessionId"].(string); ok {
		p.Data.SessionId = sessionId
	} else {
		return fmt.Errorf("packet data is missing sessionId")
	}

	if wallet, ok := data["wallet"].(string); ok {
		p.Data.Wallet = wallet
	} else {
		return fmt.Errorf("packet data is missing wallet")
	}

	return nil
}

func (p *OpenSessionPacketResponse) Unmarshal(packet json.RawMessage) error {
	err := json.Unmarshal(packet, &p)
	if err != nil {
		return err
	}

	if p.Code != OpenSessionPacketResponseCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", OpenSessionPacketResponseCode, p.Code)
	}

	return nil
}
