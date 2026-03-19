package malleable

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

func ComputeImageHash(payload *v3.CallsPayload, signature []byte, chainID *big.Int) (common.Hash, error) {
	if payload == nil {
		return common.Hash{}, fmt.Errorf("payload is nil")
	}
	if len(payload.Calls) > 128 {
		return common.Hash{}, fmt.Errorf("too many calls (%d)", len(payload.Calls))
	}

	space := nz(payload.Space)
	nonce := nz(payload.Nonce)
	root := fkeccak256(u256Bytes32(space), u256Bytes32(nonce))

	payloadChainID := payload.ChainID()
	noChainID := payloadChainID == nil || payloadChainID.Sign() == 0
	if noChainID {
		root = fkeccak256(root, common.Hash{})
	} else {
		if chainID == nil {
			chainID = payloadChainID
		}
		root = fkeccak256(root, u256Bytes32(nz(chainID)))
	}

	stringTy, _ := abi.NewType("string", "", nil)
	u256Ty, _ := abi.NewType("uint256", "", nil)
	addrTy, _ := abi.NewType("address", "", nil)
	boolTy, _ := abi.NewType("bool", "", nil)
	bytesTy, _ := abi.NewType("bytes", "", nil)

	callMetaArgs := abi.Arguments{
		{Type: stringTy},
		{Type: u256Ty},
		{Type: addrTy},
		{Type: u256Ty},
		{Type: u256Ty},
		{Type: boolTy},
		{Type: boolTy},
		{Type: u256Ty},
	}

	staticArgs := abi.Arguments{
		{Type: stringTy},
		{Type: u256Ty},
		{Type: u256Ty},
		{Type: bytesTy},
	}

	repeatArgs := abi.Arguments{
		{Type: stringTy},
		{Type: u256Ty},
		{Type: u256Ty},
		{Type: u256Ty},
		{Type: u256Ty},
		{Type: u256Ty},
	}

	for i := 0; i < len(payload.Calls); i++ {
		c := payload.Calls[i]
		b, err := callMetaArgs.Pack(
			"call",
			big.NewInt(int64(i)),
			c.To,
			nz(c.Value),
			nz(c.GasLimit),
			c.DelegateCall,
			c.OnlyFallback,
			big.NewInt(int64(c.BehaviorOnError)),
		)
		if err != nil {
			return common.Hash{}, err
		}
		root = fkeccak256(root, crypto.Keccak256Hash(b))
	}

	r := 0
	for r < len(signature) {
		if r+5 > len(signature) {
			return common.Hash{}, fmt.Errorf("signature truncated at %d", r)
		}
		tRaw := signature[r]
		r++
		cindex := binary.BigEndian.Uint16(signature[r : r+2])
		r += 2
		size := binary.BigEndian.Uint16(signature[r : r+2])
		r += 2

		repeat := (tRaw & 0x80) != 0
		t := int(tRaw & 0x7F)
		if t >= len(payload.Calls) {
			return common.Hash{}, fmt.Errorf("tindex out of range: %d", t)
		}
		if int(cindex)+int(size) > len(payload.Calls[t].Data) {
			return common.Hash{}, fmt.Errorf("section out of bounds (t=%d)", t)
		}
		section := payload.Calls[t].Data[cindex : cindex+size]

		if repeat {
			if r+3 > len(signature) {
				return common.Hash{}, fmt.Errorf("repeat truncated at %d", r)
			}
			t2 := int(signature[r])
			r++
			c2 := binary.BigEndian.Uint16(signature[r : r+2])
			r += 2

			if t2 >= len(payload.Calls) {
				return common.Hash{}, fmt.Errorf("tindex2 out of range: %d", t2)
			}
			if int(c2)+int(size) > len(payload.Calls[t2].Data) {
				return common.Hash{}, fmt.Errorf("repeat section2 out of bounds")
			}
			section2 := payload.Calls[t2].Data[c2 : c2+size]
			if crypto.Keccak256Hash(section) != crypto.Keccak256Hash(section2) {
				return common.Hash{}, fmt.Errorf("repeat section mismatch")
			}

			b, err := repeatArgs.Pack(
				"repeat-section",
				big.NewInt(int64(t)),
				new(big.Int).SetUint64(uint64(cindex)),
				new(big.Int).SetUint64(uint64(size)),
				big.NewInt(int64(t2)),
				new(big.Int).SetUint64(uint64(c2)),
			)
			if err != nil {
				return common.Hash{}, err
			}
			root = fkeccak256(root, crypto.Keccak256Hash(b))
		} else {
			b, err := staticArgs.Pack(
				"static-section",
				big.NewInt(int64(t)),
				new(big.Int).SetUint64(uint64(cindex)),
				section,
			)
			if err != nil {
				return common.Hash{}, err
			}
			root = fkeccak256(root, crypto.Keccak256Hash(b))
		}
	}

	return root, nil
}

func ValidateSignature(payload *v3.CallsPayload, signature []byte) error {
	if payload == nil {
		return fmt.Errorf("payload is nil")
	}
	r := 0
	for r < len(signature) {
		if r+5 > len(signature) {
			return fmt.Errorf("signature truncated at %d", r)
		}
		tRaw := signature[r]
		r++
		cindex := binary.BigEndian.Uint16(signature[r : r+2])
		r += 2
		size := binary.BigEndian.Uint16(signature[r : r+2])
		r += 2

		t := int(tRaw & 0x7F)
		if t >= len(payload.Calls) {
			return fmt.Errorf("tindex out of range: %d", t)
		}
		if int(cindex)+int(size) > len(payload.Calls[t].Data) {
			return fmt.Errorf("section out of bounds (t=%d)", t)
		}

		if tRaw&0x80 != 0 {
			if r+3 > len(signature) {
				return fmt.Errorf("repeat truncated at %d", r)
			}
			t2 := int(signature[r])
			r++
			c2 := binary.BigEndian.Uint16(signature[r : r+2])
			r += 2

			if t2 >= len(payload.Calls) {
				return fmt.Errorf("tindex2 out of range: %d", t2)
			}
			if int(c2)+int(size) > len(payload.Calls[t2].Data) {
				return fmt.Errorf("repeat section2 out of bounds")
			}
			section := payload.Calls[t].Data[cindex : cindex+size]
			section2 := payload.Calls[t2].Data[c2 : c2+size]
			if crypto.Keccak256Hash(section) != crypto.Keccak256Hash(section2) {
				return fmt.Errorf("repeat section mismatch")
			}
		}
	}
	return nil
}

func fkeccak256(a, b common.Hash) common.Hash {
	return crypto.Keccak256Hash(a[:], b[:])
}

func u256Bytes32(x *big.Int) common.Hash {
	var out common.Hash
	if x == nil {
		return out
	}
	b := x.Bytes()
	copy(out[32-len(b):], b)
	return out
}

func nz(x *big.Int) *big.Int {
	if x == nil {
		return big.NewInt(0)
	}
	return x
}
