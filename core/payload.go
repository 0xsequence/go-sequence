package core

import (
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type Payload interface {
	Address() common.Address
	ChainID() *big.Int

	Digest() PayloadDigest
}

type PayloadDigest struct {
	common.Hash

	Address_ common.Address
	ChainID_ *big.Int

	Payload Payload
}

func (d PayloadDigest) Address() common.Address {
	return d.Address_
}

func (d PayloadDigest) ChainID() *big.Int {
	if d.ChainID_ != nil {
		return new(big.Int).Set(d.ChainID_)
	} else {
		return new(big.Int)
	}
}

func (d PayloadDigest) Digest() PayloadDigest {
	return d
}
