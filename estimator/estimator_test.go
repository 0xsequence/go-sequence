package estimator_test

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/estimator"
	"github.com/stretchr/testify/assert"
)

func TestEstimator(t *testing.T) {
	ctx := context.Background()

	provider, err := ethrpc.NewProvider("https://dev-nodes.sequence.app/mainnet")
	assert.NoError(t, err)

	var sender, receiver common.Address
	rand.Read(sender[:])
	rand.Read(receiver[:])

	token := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	value := big.NewInt(1000000)

	transfer, err := contracts.ERC20Mock.Encode("transfer", receiver, value)
	assert.NoError(t, err)

	gas, err := estimator.Estimate(ctx, token, transfer, nil, provider)
	assert.Error(t, err)
	assert.Zero(t, gas)

	fmt.Println("error without preconditions:", err)

	preconditions := map[common.Address]estimator.AddressPreconditions{
		sender: {
			ERC20: map[common.Address]estimator.ERC20Preconditions{
				token: {Balance: value},
			},
		},
	}
	gas, err = estimator.Estimate(ctx, token, transfer, preconditions, provider)
	assert.NoError(t, err)
	assert.Positive(t, gas)

	fmt.Println("gas assuming preconditions:", gas)
}
