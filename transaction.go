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

func (t Transactions) Execute(address common.Address, signature []byte) ([]byte, error) {
	return contracts.V3.WalletStage1Module.Encode("execute", t.Encode(address), signature)
}

func (t Transactions) MustExecute(address common.Address, signature []byte) []byte {
	data, err := t.Execute(address, signature)
	if err != nil {
		panic(err)
	}
	return data
}

func (t Transactions) SelfExecute(address common.Address) ([]byte, error) {
	return contracts.V3.WalletStage1Module.Encode("selfExecute", t.Encode(address))
}

func (t Transactions) MustSelfExecute(address common.Address) []byte {
	data, err := t.SelfExecute(address)
	if err != nil {
		panic(err)
	}
	return data
}

func (t Transactions) Encode(address common.Address) []byte {
	return v3.EncodeCalls(address, t.Calls, t.Space, t.Nonce)
}
