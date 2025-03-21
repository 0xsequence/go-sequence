package sequence

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type SignedTransactions struct {
	Transactions
	Signature []byte
}

type Transactions []v3.Call

func (t Transactions) Payload(address common.Address, chainID *big.Int, space, nonce *big.Int) v3.CallsPayload {
	return v3.ConstructCallsPayload(address, chainID, t, space, nonce)
}

type Transaction struct {
	v3.Call
}
