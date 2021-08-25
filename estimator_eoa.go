package sequence

import (
	"context"
	"fmt"
	"log"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/goware/cachestore"
	"github.com/goware/cachestore/memlru"
)

const (
	eoaCacheSize = 100
)

var (
	eoaCache cachestore.Storage
)

var (
	cachedTrue  = byte('t')
	cachedFalse = byte('f')
)

func init() {
	var err error
	eoaCache, err = memlru.NewWithSize(eoaCacheSize)
	if err != nil {
		log.Fatalf("failed to initialize EOA cache: %v", err)
	}
}

func isEOA(ctx context.Context, provider *ethrpc.Provider, address common.Address) (bool, error) {
	key := fmt.Sprintf("isEOA::%d::%v", provider.Config.ChaindID, address)

	if val, _ := eoaCache.Get(ctx, key); val != nil {
		// we have recorded data for this key, let's use it
		return val[0] == cachedTrue, nil
	}

	// TODO: will this function time-out at some point?
	code, err := provider.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}

	if len(code) == 0 {
		// if the address does not contain a smart contract, then it's an EOA
		_ = eoaCache.Set(ctx, key, []byte{cachedTrue})
		return true, nil
	}

	_ = eoaCache.Set(ctx, key, []byte{cachedFalse})
	return false, nil
}
