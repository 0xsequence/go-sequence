package sequence

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type SignedTransactions struct {
	v3.CallsPayload
	Signature []byte
}

func (t SignedTransactions) Data() ([]byte, error) {
	return contracts.V3.WalletStage1Module.Encode("execute", t.Encode(t.Address()), t.Signature)
}

type Transactions struct {
	Calls        []v3.Call
	Space, Nonce *big.Int
}

func (t Transactions) Payload(address common.Address, chainID *big.Int) v3.CallsPayload {
	return v3.ConstructCallsPayload(address, chainID, t.Calls, t.Space, t.Nonce)
}

type Transaction struct {
	v3.Call
}
