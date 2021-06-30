package prototyp

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigIntJSONMarshalling(t *testing.T) {
	b, _ := big.NewInt(1).SetString("222222222222222222222222222222222255555555555555222222", 10)
	b1 := ToBigInt(b)

	data, err := json.Marshal(b1)
	assert.NoError(t, err)
	assert.Equal(t, "\"222222222222222222222222222222222255555555555555222222\"", string(data))

	var b2 BigInt
	err = json.Unmarshal(data, &b2)
	assert.NoError(t, err)

	if b1.Int().Cmp(b2.Int()) != 0 {
		assert.Fail(t, "bigints are not the same after unmarshalling")
	}
}

func TestBigIntValue(t *testing.T) {
	var b big.Int

	n := big.NewInt(0).Cmp(&b)
	assert.Equal(t, 0, n)

	b2, ok := b.SetString("123", 10)
	assert.True(t, ok)

	assert.Equal(t, int64(123), b2.Int64())
}

func TestBigIntAdd(t *testing.T) {
	b1 := NewBigInt(42)
	b2 := NewBigInt(84)

	b1.Add(b2.Int())

	assert.Equal(t, int64(42+84), b1.Int().Int64())
}

func TestBigIntScan(t *testing.T) {
	b := BigInt{}

	assert.NoError(t, b.Scan("1"))
	assert.NoError(t, b.Scan([]byte("1")))
	assert.Error(t, b.Scan(1))
	assert.Error(t, b.Scan(1.0))
	assert.Error(t, b.Scan("1x"))
	assert.Error(t, b.Scan("1."))
}
