package estimator

import (
	"encoding/json"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeAccessListResponseAllowsMissingStorageKeys(t *testing.T) {
	address := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	list, err := decodeAccessListResponse(json.RawMessage(`{
		"accessList": [
			{
				"address": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
			}
		]
	}`))
	require.NoError(t, err)
	require.Len(t, list, 1)
	assert.Equal(t, address, list[0].Address)
	assert.Nil(t, list[0].StorageKeys)
}

func TestDecodeAccessListResponseRequiresAddress(t *testing.T) {
	_, err := decodeAccessListResponse(json.RawMessage(`{
		"accessList": [
			{
				"storageKeys": []
			}
		]
	}`))
	require.EqualError(t, err, "unable to decode eth_createAccessList response: missing required field 'address' for AccessTuple")
}
