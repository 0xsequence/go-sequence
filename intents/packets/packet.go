package packets

type BasePacket struct {
	Code    string `json:"code"`
	Issued  uint64 `json:"issued"`
	Expires uint64 `json:"expires"`
}

type BasePacketForWallet struct {
	BasePacket
	Wallet string `json:"wallet"`
}
