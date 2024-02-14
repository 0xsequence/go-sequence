package prototyp

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.Equal(t, b.Int64(), int64(1))
	assert.NoError(t, b.Scan([]byte("1")))
	require.Equal(t, b.Int64(), int64(1))
	assert.Error(t, b.Scan(1))
	require.Equal(t, b.Int64(), int64(1))
	assert.Error(t, b.Scan(1.0))
	require.Equal(t, b.Int64(), int64(1))
	assert.Error(t, b.Scan("1x"))
	require.Equal(t, b.Int64(), int64(1))
	assert.Error(t, b.Scan("1."))
	require.Equal(t, b.Int64(), int64(1))
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

	b = NewBigIntFromString("1234", 10)
	assert.Equal(t, int64(1234), b.Int64())
}

func TestToBigInt(t *testing.T) {
	a := big.NewInt(1)

	// copies  a
	b := ToBigInt(a)
	c := ToBigInt(a)

	// not mutating b or a
	c.Sub(b.Int())

	require.Equal(t, int64(1), b.Int64())
	require.Equal(t, int64(0), c.Int64())
}

func TestSetString(t *testing.T) {
	a := NewBigInt(1)

	ok := a.SetString("12", 10)
	require.True(t, ok)
	require.Equal(t, int64(12), a.Int64())
}

func TestEqual(t *testing.T) {
	a := NewBigInt(10)
	b := NewBigInt(10)
	require.True(t, b.Equals(a.Int()))
}

func TestBigIntArray(t *testing.T) {
	input := []int64{1, 2, 3, 4}

	a := []*big.Int{}
	for _, i := range input {
		a = append(a, big.NewInt(i))
	}

	b := ToBigIntArray(a)
	for i := range input {
		require.Equal(t, a[i].Int64(), b[i].Int64())
	}
}

func TestToBigIntArrayFromStringArray(t *testing.T) {
	input := []string{"10", "12", "123", "562"}

	expected := []BigInt{}
	for _, i := range input {
		expected = append(expected, NewBigIntFromNumberString(i))
	}

	got, err := ToBigIntArrayFromStringArray(input, 10)
	require.NoError(t, err)
	require.Equal(t, expected, got)
}

func TestToBigIntFromInt64(t *testing.T) {
	got := ToBigIntFromInt64(10)
	expected := NewBigInt(10)

	require.Equal(t, expected, got)
}

func TestString(t *testing.T) {
	a := NewBigInt(10)
	require.Equal(t, "10", a.String())
}

func TestAdd(t *testing.T) {
	AddArgs := [][]*big.Int{
		// x,y,z
		{big.NewInt(1), big.NewInt(2), big.NewInt(3)},
		{big.NewInt(3), big.NewInt(2), big.NewInt(5)},
		{big.NewInt(123), big.NewInt(5), big.NewInt(128)},
		{big.NewInt(1), big.NewInt(200), big.NewInt(201)},
		{big.NewInt(-100), big.NewInt(200), big.NewInt(100)},
		{big.NewInt(100), big.NewInt(-200), big.NewInt(-100)},
	}

	for _, arg := range AddArgs {
		x := ToBigInt(arg[0])
		y := ToBigInt(arg[1])
		x.Add(y.Int())
		require.Equal(t, x.Int(), arg[2])
	}
}

func TestComparison(t *testing.T) {
	type Cases struct {
		A        BigInt
		B        *big.Int
		Expected bool
	}

	greater := []Cases{{
		A:        NewBigInt(10),
		B:        big.NewInt(1),
		Expected: true,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(20),
		Expected: false,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(10),
		Expected: false,
	}}

	for _, test := range greater {
		require.Equal(t, test.Expected, test.A.Gt(test.B))
	}

	greaterEq := []Cases{{
		A:        NewBigInt(10),
		B:        big.NewInt(1),
		Expected: true,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(20),
		Expected: false,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(10),
		Expected: true,
	}}

	for _, test := range greaterEq {
		require.Equal(t, test.Expected, test.A.Gte(test.B))
	}

	lesser := []Cases{{
		A:        NewBigInt(10),
		B:        big.NewInt(1),
		Expected: false,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(20),
		Expected: true,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(10),
		Expected: false,
	}}

	for _, test := range lesser {
		require.Equal(t, test.Expected, test.A.Lt(test.B))
	}

	lesserEq := []Cases{{
		A:        NewBigInt(10),
		B:        big.NewInt(1),
		Expected: false,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(20),
		Expected: true,
	}, {
		A:        NewBigInt(10),
		B:        big.NewInt(10),
		Expected: true,
	}}

	for _, test := range lesserEq {
		require.Equal(t, test.Expected, test.A.Lte(test.B))
	}
}
