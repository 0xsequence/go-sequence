package core

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type Payload interface {
	Digest() PayloadDigest
}

type PayloadDigest struct {
	common.Hash

	Address common.Address
	ChainID *big.Int

	Payload Payload
}

func (d PayloadDigest) Digest() PayloadDigest {
	return d
}
