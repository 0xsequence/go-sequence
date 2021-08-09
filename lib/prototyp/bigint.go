package prototyp

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"strings"
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

func NewBigIntFromString(s string, base int) BigInt {
	b := big.NewInt(0)
	b, _ = b.SetString(s, base)
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

func (b *BigInt) SetString(s string, base int) bool {
	v := big.Int(*b)
	n, ok := v.SetString(s, base)
	if !ok {
		return false
	}
	*b = BigInt(*n)
	return true
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

	var svalue string
	switch v := src.(type) {
	case string:
		svalue = v
	case []byte:
		svalue = string(v)
	default:
		return fmt.Errorf("BigInt.Scan: unexpected type %T", src)
	}

	// pgx driver returns NeX where N is digits and X is exponent
	parts := strings.SplitN(svalue, "e", 2)

	var ok bool
	i := &big.Int{}
	i, ok = i.SetString(parts[0], 10)
	if !ok {
		return fmt.Errorf("BigInt.Scan: failed to scan value %q", svalue)
	}

	if len(parts) >= 2 {
		exp := big.NewInt(0)
		exp, ok = exp.SetString(parts[1], 10)
		if !ok {
			return fmt.Errorf("BigInt.Scan failed to scan exp component %q", svalue)
		}
		i = i.Mul(i, big.NewInt(1).Exp(big.NewInt(10), exp, nil))
	}

	*b = BigInt(*i)

	return nil
}
