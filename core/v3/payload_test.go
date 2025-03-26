package v3_test

import (
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/assert"
)

func TestDigestPayloadMatchesMessagePayload(t *testing.T) {
	// dapp asks wallet to eth_sign 'hello world'
	// dapp expects message signature on '\x19Ethereum Signed Message:\n11hello world'
	// dapp calls isValidSignature(keccak256('\x19Ethereum Signed Message:\n11hello world'), signature)
	// contract validates signature on keccak256('\x19Ethereum Signed Message:\n11hello world')
	// so digest payload on keccak256(msg) must conform to message payload on msg
	// where msg is '\x19Ethereum Signed Message:\n11hello world'

	address := common.HexToAddress("0x468E8e29F6cfb0F6b7ff10ec6A1AB516ec849c04")
	chainID := common.Big1
	message := "hello world"
	digest := crypto.Keccak256Hash([]byte(message))

	messagePayload := v3.NewMessagePayload(address, chainID, []byte(message))
	digestPayload := v3.NewDigestPayload(address, chainID, digest)

	assert.Equal(t, digestPayload.Digest().Hash, messagePayload.Digest().Hash)
}
