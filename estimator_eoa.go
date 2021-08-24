package sequence

import (
	"context"
	"fmt"
	"log"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/hashicorp/golang-lru"
)

const (
	isEOAMaxConcurrentTasks = 10
	isEOACacheSize          = 100
)

var (
	isEOATicket = make(chan struct{}, isEOAMaxConcurrentTasks)

	isEOACache *lru.Cache
)

func init() {
	for i := 0; i < isEOAMaxConcurrentTasks; i++ {
		isEOATicket <- struct{}{}
	}
	var err error
	isEOACache, err = lru.New(isEOACacheSize)
	if err != nil {
		log.Fatalf("failed to initialize EOA cache: %v", err)
	}
}

func isEOA(ctx context.Context, provider *ethrpc.Provider, address common.Address) (bool, error) {
	ticket := <-isEOATicket
	defer func() {
		isEOATicket <- ticket
	}()

	key := fmt.Sprintf("%d::%v", provider.Config.ChaindID, address)

	if val, ok := isEOACache.Get(key); ok {
		// we have recorded data for this key, let's use it
		return val.(bool), nil
	}

	// TODO: will this function time-out at some point?
	code, err := provider.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}

	value := len(code) == 0

	// store value in the cache
	_ = isEOACache.Add(key, value)

	return value, nil
}
