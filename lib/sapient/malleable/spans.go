package malleable

import "github.com/0xsequence/go-sequence/lib/abicalldata"

type Span struct {
	Start int
	Len   int
}

// Re-exports for callers that imported these types from malleable; definitions live in abicalldata.
type (
	ByteRange     = abicalldata.ByteRange
	Selector      = abicalldata.Selector
	RangeSelector = abicalldata.RangeSelector
)

// NewRangeSelector delegates to abicalldata.NewRangeSelector.
func NewRangeSelector(callIndex, offset, size int) RangeSelector {
	return abicalldata.NewRangeSelector(callIndex, offset, size)
}
