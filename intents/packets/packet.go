package packets

import "time"

type BasePacket struct {
	Code    string `json:"code"`
	Issued  uint64 `json:"issued"`
	Expires uint64 `json:"expires"`
}

func (p *BasePacket) IsValid() bool {
	now := uint64(time.Now().Unix())
	return p.Code != "" && p.Issued <= now && p.Expires > now
}

type BasePacketForWallet struct {
	BasePacket
	Wallet string `json:"wallet"`
}
