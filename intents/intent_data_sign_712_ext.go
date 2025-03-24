package intents

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

func (p *IntentDataSignTypedData) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}
	return n, nil
}

func (p *IntentDataSignTypedData) digest() (common.Hash, error) {
	// typedData, ok := p.TypedData.(*ethcoder.TypedData)
	// if !ok {
	// 	return nil, fmt.Errorf("typedData field is not a valid typed data object")
	// }

	digest, _, err := p.TypedData.Encode()
	if err != nil {
		return common.Hash{}, err
	}

	return common.Hash(digest), nil
}

func (p *IntentDataSignTypedData) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

func (p *IntentDataSignTypedData) payload() (v3.Payload, error) {
	chainID, err := p.chainID()
	if err != nil {
		return nil, err
	}

	digest, err := p.digest()
	if err != nil {
		return nil, err
	}

	return v3.ConstructDigestPayload(p.wallet(), chainID, digest), nil
}

func (p *IntentDataSignTypedData) IsValidInterpretation(digest common.Hash) bool {
	payload, err := p.payload()
	if err != nil {
		return false
	}

	return digest == payload.Digest().Hash
}

// func (p *IntentDataSignTypedData) UnmarshalJSON(data []byte) error {
// 	type Raw struct {
// 		Network   string          `json:"network"`
// 		Wallet    string          `json:"wallet"`
// 		TypedData json.RawMessage `json:"typedData"`
// 	}

// 	dec := json.NewDecoder(bytes.NewReader(data))
// 	var raw Raw
// 	if err := dec.Decode(&raw); err != nil {
// 		return err
// 	}

// 	typedData, err := ethcoder.TypedDataFromJSON(string(raw.TypedData))
// 	if err != nil {
// 		return err
// 	}

// 	p.Network = raw.Network
// 	p.Wallet = raw.Wallet
// 	p.TypedData = typedData
// 	return nil
// }
