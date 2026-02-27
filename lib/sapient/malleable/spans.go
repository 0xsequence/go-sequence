package malleable

import (
	"fmt"

	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type Span struct {
	Start int
	Len   int
}

type ByteRange struct {
	CallIndex int
	Offset    int
	Size      int
}

func (r ByteRange) Slice(payload *v3.CallsPayload) ([]byte, error) {
	if payload == nil {
		return nil, fmt.Errorf("payload is nil")
	}
	if r.CallIndex < 0 || r.CallIndex >= len(payload.Calls) {
		return nil, fmt.Errorf("call index out of range: %d", r.CallIndex)
	}
	data := payload.Calls[r.CallIndex].Data
	if r.Offset < 0 || r.Size < 0 || r.Offset+r.Size > len(data) {
		return nil, fmt.Errorf("range out of bounds: [%d,%d) with len %d", r.Offset, r.Offset+r.Size, len(data))
	}
	return data[r.Offset : r.Offset+r.Size], nil
}

type RangeSelector struct {
	Range ByteRange
}

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
