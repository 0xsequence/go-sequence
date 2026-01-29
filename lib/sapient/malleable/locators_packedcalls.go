package malleable

import (
	"encoding/binary"
	"fmt"
)

type PackedCallsLayout struct {
	GlobalFlag byte
	NumCalls   int
	CallData   []Span
}

// ParsePackedCalls parses the packed calls layout from the packed payload.calls data.
func ParsePackedCalls(packed []byte) (*PackedCallsLayout, error) {
	if len(packed) < 1 {
		return nil, fmt.Errorf("packed calls too short")
	}
	p := 0
	globalFlag := packed[p]
	p++

	if globalFlag&0x01 == 0x00 {
		if p+20 > len(packed) {
			return nil, fmt.Errorf("packed calls truncated reading space")
		}
		p += 20
	}

	nonceSize := int((globalFlag >> 1) & 0x07)
	if nonceSize > 0 {
		if p+nonceSize > len(packed) {
			return nil, fmt.Errorf("packed calls truncated reading nonce (%d bytes)", nonceSize)
		}
		p += nonceSize
	}

	var numCalls int
	if globalFlag&0x10 == 0x10 {
		numCalls = 1
	} else {
		if globalFlag&0x20 == 0x20 {
			if p+2 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading numCalls (u16)")
			}
			numCalls = int(binary.BigEndian.Uint16(packed[p : p+2]))
			p += 2
		} else {
			if p+1 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading numCalls (u8)")
			}
			numCalls = int(packed[p])
			p++
		}
	}

	callData := make([]Span, numCalls)
	for i := 0; i < numCalls; i++ {
		if p+1 > len(packed) {
			return nil, fmt.Errorf("packed calls truncated reading call flags (i=%d)", i)
		}
		flags := packed[p]
		p++

		if flags&0x01 == 0x00 {
			if p+20 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading to (i=%d)", i)
			}
			p += 20
		}

		if flags&0x02 == 0x02 {
			if p+32 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading value (i=%d)", i)
			}
			p += 32
		}

		if flags&0x04 == 0x04 {
			if p+3 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading calldataSize (i=%d)", i)
			}
			calldataSize := int(packed[p])<<16 | int(packed[p+1])<<8 | int(packed[p+2])
			p += 3

			if p+calldataSize > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading calldata bytes (i=%d size=%d)", i, calldataSize)
			}
			callData[i] = Span{Start: p, Len: calldataSize}
			p += calldataSize
		} else {
			callData[i] = Span{Start: -1, Len: 0}
		}

		if flags&0x08 == 0x08 {
			if p+32 > len(packed) {
				return nil, fmt.Errorf("packed calls truncated reading gasLimit (i=%d)", i)
			}
			p += 32
		}
	}

	return &PackedCallsLayout{
		GlobalFlag: globalFlag,
		NumCalls:   numCalls,
		CallData:   callData,
	}, nil
}
