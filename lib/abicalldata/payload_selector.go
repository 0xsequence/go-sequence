package abicalldata

import (
	"fmt"

	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// ByteRange identifies a contiguous slice of one call's calldata (data field) in a CallsPayload.
type ByteRange struct {
	CallIndex int
	Offset    int
	Size      int
}

// Slice returns the referenced bytes from payload.Calls[CallIndex].Data.
func (r ByteRange) Slice(payload *v3.CallsPayload) ([]byte, error) {
	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}
	if r.CallIndex < 0 || r.CallIndex >= len(payload.Calls) {
		return nil, fmt.Errorf("call index out of range: %d", r.CallIndex)
	}
	data := payload.Calls[r.CallIndex].Data
	if r.Offset < 0 || r.Size < 0 {
		return nil, fmt.Errorf("range offset/size negative")
	}
	if r.Size > len(data) || r.Offset > len(data)-r.Size {
		return nil, fmt.Errorf("range out of bounds: offset=%d size=%d with len %d", r.Offset, r.Size, len(data))
	}
	return data[r.Offset : r.Offset+r.Size], nil
}

// Selector resolves one or more byte ranges in a calls payload (e.g. ABI paths or fixed ranges).
type Selector interface {
	Resolve(payload *v3.CallsPayload) ([]ByteRange, error)
	String() string
}

// RangeSelector is a Selector backed by a fixed ByteRange (validated on Resolve).
type RangeSelector struct {
	Range ByteRange
}

// NewRangeSelector returns a Selector for an explicit call index, offset, and size within that call's data.
func NewRangeSelector(callIndex, offset, size int) RangeSelector {
	return RangeSelector{
		Range: ByteRange{
			CallIndex: callIndex,
			Offset:    offset,
			Size:      size,
		},
	}
}

func (r RangeSelector) Resolve(payload *v3.CallsPayload) ([]ByteRange, error) {
	if _, err := r.Range.Slice(payload); err != nil {
		return nil, err
	}
	return []ByteRange{r.Range}, nil
}

func (r RangeSelector) String() string {
	return fmt.Sprintf("range(call=%d,offset=%d,size=%d)", r.Range.CallIndex, r.Range.Offset, r.Range.Size)
}
