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

func TestBigIntFromBaseString(t *testing.T) {
	testCases := []struct {
		fn       func(string) BigInt
		input    string
		expected int64
		testName string
	}{
		// Binary tests
		{NewBigIntFromBinaryString, "00101010", 42, "Binary 42"},
		{NewBigIntFromBinaryString, "-00101010", -42, "Binary -42"},
		{NewBigIntFromBinaryString, "0B00101010", 42, "Binary with 0B 42"},
		{NewBigIntFromBinaryString, "0b00101010", 42, "Binary with 0b 42"},

		// Octal tests
		{NewBigIntFromOctalString, "52", 42, "Octal 42"},
		{NewBigIntFromOctalString, "-52", -42, "Octal -42"},
		{NewBigIntFromOctalString, "0O52", 42, "Octal with 0O 42"},
		{NewBigIntFromOctalString, "0o52", 42, "Octal with 0o 42"},

		// Decimal tests
		{NewBigIntFromDecimalString, "42", 42, "Decimal 42"},
		{NewBigIntFromDecimalString, "-42", -42, "Decimal -42"},

		// Hexadecimal tests
		{NewBigIntFromHexString, "2A", 42, "Hex 2A"},
		{NewBigIntFromHexString, "2a", 42, "Hex 2a"},
		{NewBigIntFromHexString, "-2A", -42, "Hex -2A"},
		{NewBigIntFromHexString, "-2a", -42, "Hex -2a"},
		{NewBigIntFromHexString, "0x2a", 42, "Hex with 0x 42"},
		{NewBigIntFromHexString, "0X2a", 42, "Hex with 0X 42"},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			b := tc.fn(tc.input)
			assert.Equal(t, tc.expected, b.Int64())
		})
	}
}

func TestBigIntInvalidInput(t *testing.T) {
	b := NewBigIntFromNumberString("1234")
	assert.Equal(t, int64(1234), b.Int64())

	// if invalid, it will parse out the number
	b = NewBigIntFromNumberString("/1234")
	assert.Equal(t, int64(1234), b.Int64())

	// if invalid, it will parse out the number
	b = NewBigIntFromHexString("/0xaBc-$$2Efg")
	assert.Equal(t, int64(0xaBc2Ef), b.Int64())

	// should return 0, since base is not valid
	b = NewBigIntFromString("1234", 666)
	assert.Equal(t, int64(0), b.Int64())
}

func TestBigIntCopy(t *testing.T) {
	a := ToBigInt(big.NewInt(1))
	b := ToBigInt(a.Int())
	a.Sub(big.NewInt(1))

	assert.Equal(t, int64(0), a.Int64())
	assert.Equal(t, int64(1), b.Int64())
}
