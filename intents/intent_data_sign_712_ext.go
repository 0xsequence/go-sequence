package intents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
)

func (p *IntentDataSignTypedData) UnmarshalJSON(data []byte) error {
	type Raw struct {
		Network   string          `json:"network"`
		Wallet    string          `json:"wallet"`
		TypedData json.RawMessage `json:"typedData"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	var raw Raw
	if err := dec.Decode(&raw); err != nil {
		return err
	}

	typedData, err := ethcoder.TypedDataFromJSON(string(raw.TypedData))
	if err != nil {
		return err
	}

	p.Network = raw.Network
	p.Wallet = raw.Wallet
	p.TypedData = typedData
	return nil
}

func (p *IntentDataSignTypedData) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}
	return n, nil
}

func (p *IntentDataSignTypedData) message() ([]byte, error) {
	typedData, ok := p.TypedData.(*ethcoder.TypedData)
	if !ok {
		return nil, fmt.Errorf("typedData field is not a valid typed data object")
	}

	_, encodedMessage, err := typedData.Encode()
	if err != nil {
		return nil, err
	}

	return encodedMessage, nil
}

func (p *IntentDataSignTypedData) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

func (p *IntentDataSignTypedData) subdigest() ([]byte, error) {
	chainID, err := p.chainID()
	if err != nil {
		return nil, err
	}
	msgData, err := p.message()
	if err != nil {
		return nil, err
	}
	return sequence.SubDigest(chainID, p.wallet(), sequence.MessageDigest(msgData))
}

// A SignMessagePacket (intent) *MUST* be mapped to a regular "SignMessage" Sequence action, this means that
// it must adhere to the following rules:
// - the subdigest must match `SubDigest(chainID, Wallet, Digest(Message))`
func (p *IntentDataSignTypedData) IsValidInterpretation(subdigest common.Hash) bool {
	selfSubDigest, err := p.subdigest()
	if err != nil {
		return false
	}

	return bytes.Equal(selfSubDigest, subdigest[:])
}
