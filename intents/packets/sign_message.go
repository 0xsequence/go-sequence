package packets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/intents"
)

type SignMessagePacket struct {
	Code    string `json:"code"`
	Wallet  string `json:"wallet"`
	Network string `json:"network"`
	Message string `json:"message"`
}

const SignMessagePacketCode = "signMessage"

func (p *SignMessagePacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
	if err != nil {
		return err
	}

	if p.Code != SignMessagePacketCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", SignMessagePacketCode, p.Code)
	}

	return nil
}

func (p *SignMessagePacket) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}

	return n, nil
}

func (p *SignMessagePacket) message() []byte {
	return common.FromHex(p.Message)
}

func (p *SignMessagePacket) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

func (p *SignMessagePacket) subdigest() ([]byte, error) {
	chainID, err := p.chainID()
	if err != nil {
		return nil, err
	}

	return sequence.SubDigest(chainID, p.wallet(), sequence.MessageDigest(p.message()))
}

// A SignMessagePacket (intent) *MUST* be mapped to a regular "SignMessage" Sequence action, this means that
// it must adhere to the following rules:
// - the subdigest must match `SubDigest(chainID, Wallet, Digest(Message))`
func (p *SignMessagePacket) IsValidInterpretation(subdigest common.Hash) bool {
	selfSubDigest, err := p.subdigest()
	if err != nil {
		return false
	}

	return bytes.Equal(selfSubDigest, subdigest[:])
}
