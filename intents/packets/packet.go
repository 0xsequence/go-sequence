package packets

import (
	"fmt"
	"time"
)

const allowedTimeDrift = 5

var (
	ErrInvalidPacketCode      = fmt.Errorf("invalid packet code")
	ErrorPackedIssuedInFuture = fmt.Errorf("packet issued time is in the future")
	ErrorPacketExpired        = fmt.Errorf("packet expired")
)

type BasePacket struct {
	Code    string `json:"code"`
	Issued  uint64 `json:"issued"`
	Expires uint64 `json:"expires"`
}

func (p *BasePacket) IsValid() (bool, error) {
	now := uint64(time.Now().Unix())
	if p.Code == "" {
		return false, ErrInvalidPacketCode
	} else if p.Issued <= now+allowedTimeDrift {
		return false, ErrorPackedIssuedInFuture
	} else if p.Expires > now-allowedTimeDrift {
		return false, ErrorPacketExpired
	}
	return true, nil
}

type BasePacketForWallet struct {
	BasePacket
	Wallet string `json:"wallet"`
}
