package intents

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
)

func (p *IntentDataSignMessage) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}

	return n, nil
}

func (p *IntentDataSignMessage) message() []byte {
	return common.FromHex(p.Message)
}

func (p *IntentDataSignMessage) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

func (p *IntentDataSignMessage) subdigest() ([]byte, error) {
	chainID, err := p.chainID()
	if err != nil {
		return nil, err
	}

	return sequence.SubDigest(chainID, p.wallet(), sequence.MessageDigest(p.message()))
}

// A SignMessagePacket (intent) *MUST* be mapped to a regular "SignMessage" Sequence action, this means that
// it must adhere to the following rules:
// - the subdigest must match `SubDigest(chainID, Wallet, Digest(Message))`
func (p *IntentDataSignMessage) IsValidInterpretation(subdigest common.Hash) bool {
	selfSubDigest, err := p.subdigest()
	if err != nil {
		return false
	}

	return bytes.Equal(selfSubDigest, subdigest[:])
}
