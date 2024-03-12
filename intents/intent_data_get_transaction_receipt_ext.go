package intents

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"

	"github.com/0xsequence/go-sequence"
)

func (p *IntentDataGetTransactionReceipt) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}

	return n, nil
}

func (p *IntentDataGetTransactionReceipt) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}
