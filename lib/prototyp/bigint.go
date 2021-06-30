package prototyp

import (
	"database/sql/driver"
	"fmt"
	"math/big"
)

// BigInt is a type alias for big.Int used for JSON/Database marshalling.
//
// For JSON values we encoded BigInt's as strings.
//
// For Database values we encoded BigInt's as NUMERIC(78).
type BigInt big.Int

func NewBigInt(n int64) BigInt {
	b := big.NewInt(n)
	return BigInt(*b)
}

func ToBigInt(b *big.Int) BigInt {
	if b == nil {
		return BigInt{}
	}
	return BigInt(*b)
}

func ToBigIntArray(bs []*big.Int) []BigInt {
	var pbs []BigInt
	for _, b := range bs {
		pbs = append(pbs, ToBigInt(b))
	}
	return pbs
}

func ToBigIntArrayFromStringArray(s []string, base int) ([]BigInt, error) {
	var pbs []BigInt
	for _, v := range s {
		b, ok := (&big.Int{}).SetString(v, base)
		if !ok {
			return nil, fmt.Errorf("invalid number %s", s)
		}
		pbs = append(pbs, ToBigInt(b))
	}
	return pbs, nil
}

func ToBigIntFromInt64(n int64) BigInt {
	return ToBigInt(big.NewInt(n))
}

func (b BigInt) String() string {
	return b.Int().String()
}

func (b BigInt) Int() *big.Int {
	v := big.Int(b)
	return &v
}

func (b BigInt) Uint64() uint64 {
	return b.Int().Uint64()
}

func (b BigInt) Int64() int64 {
	return b.Int().Int64()
}

func (b *BigInt) Add(n *big.Int) {
	z := b.Int().Add(b.Int(), n)
	*b = BigInt(*z)
}

func (b *BigInt) Sub(n *big.Int) {
	z := b.Int().Sub(b.Int(), n)
	*b = BigInt(*z)
}

func (b BigInt) Equals(n *big.Int) bool {
	return b.Int().Cmp(n) == 0
}

func (b BigInt) Gt(n *big.Int) bool {
	return b.Int().Cmp(n) == 1
}

func (b BigInt) Gte(n *big.Int) bool {
	return b.Int().Cmp(n) == 0 || b.Int().Cmp(n) == 1
}

func (b BigInt) Lt(n *big.Int) bool {
	return b.Int().Cmp(n) == -1
}

func (b BigInt) Lte(n *big.Int) bool {
	return b.Int().Cmp(n) == 0 || b.Int().Cmp(n) == -1
}

// MarshalText implements encoding.TextMarshaler.
func (b BigInt) MarshalText() ([]byte, error) {
	v := fmt.Sprintf("\"%s\"", b.String())
	return []byte(v), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (b *BigInt) UnmarshalText(text []byte) error {
	t := string(text)
	if len(text) <= 2 || t == "null" || t == "" {
		return nil
	}
	i, _ := big.NewInt(0).SetString(string(text[1:len(text)-1]), 10)
	*b = BigInt(*i)
	return nil
}

// MarshalJSON implements json.Marshaler
func (b BigInt) MarshalJSON() ([]byte, error) {
	return b.MarshalText()
}

// UnmarshalJSON implements json.Unmarshaler
func (b *BigInt) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		return nil
	}
	return b.UnmarshalText(text)
}

func (b BigInt) Value() (driver.Value, error) {
	return b.String(), nil
}

func (b *BigInt) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	i := &big.Int{}
	i, _ = i.SetString(string(src.([]byte)), 10)
	*b = BigInt(*i)
	return nil
}