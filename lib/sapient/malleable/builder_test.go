package malleable

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/stretchr/testify/require"
)

func TestBuilder_StaticComplement(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{Data: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}, big.NewInt(0), big.NewInt(0))

	builder := NewBuilder(&payload, &BuilderOptions{MergeAdjacentStatic: true})
	builder.Malleable(NewRangeSelector(0, 2, 2))
	builder.Malleable(NewRangeSelector(0, 7, 2))

	sig, plan, err := builder.Build()
	require.NoError(t, err)
	require.NotEmpty(t, sig)

	require.Equal(t, []StaticSection{
		{TIndex: 0, CIndex: 0, Size: 2},
		{TIndex: 0, CIndex: 4, Size: 3},
		{TIndex: 0, CIndex: 9, Size: 1},
	}, plan.Static)
	require.Empty(t, plan.Repeat)
}

func TestBuilder_RepeatValidation(t *testing.T) {
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{Data: []byte{0xaa, 0xbb, 0xcc, 0xdd}},
		{Data: []byte{0xaa, 0xbb, 0xee, 0xff}},
	}, big.NewInt(0), big.NewInt(0))

	builder := NewBuilder(&payload, &BuilderOptions{ValidateRepeats: true})
	builder.Repeat(NewRangeSelector(0, 0, 2), NewRangeSelector(1, 0, 2))
	_, _, err := builder.Build()
	require.NoError(t, err)

	builder = NewBuilder(&payload, &BuilderOptions{ValidateRepeats: true})
	builder.Repeat(NewRangeSelector(0, 0, 2), NewRangeSelector(1, 2, 2))
	_, _, err = builder.Build()
	require.Error(t, err)
}

func TestBuilder_RepeatRejectsOffsetOrSizeOverUint16(t *testing.T) {
	// Calldata beyond 65535 bytes: repeat section serializes offset/size as uint16,
	// so Build() must fail instead of silently wrapping.
	largeSize := 70000
	largeData := make([]byte, largeSize)
	for i := range largeData {
		largeData[i] = byte(i & 0xff)
	}
	payload := v3.NewCallsPayload(common.Address{}, big.NewInt(1), []v3.Call{
		{Data: largeData},
	}, big.NewInt(0), big.NewInt(0))

	// Repeat with first range at offset > 65535
	builder := NewBuilder(&payload, nil)
	builder.Repeat(
		NewRangeSelector(0, 65536, 4),
		NewRangeSelector(0, 0, 4),
	)
	_, _, err := builder.Build()
	require.Error(t, err)
	require.Contains(t, err.Error(), "exceeds uint16 range")
	require.Contains(t, err.Error(), "65536")
}

func TestEncodeDecodeSignature_RoundTrip(t *testing.T) {
	statics := []StaticSection{
		{TIndex: 0, CIndex: 1, Size: 2},
		{TIndex: 1, CIndex: 3, Size: 4},
	}
	repeats := []RepeatSection{
		{TIndex: 0, CIndex: 5, Size: 2, TIndex2: 1, CIndex2: 6},
	}

	sig, err := EncodeSignature(statics, repeats)
	require.NoError(t, err)

	sections, err := DecodeSignature(sig)
	require.NoError(t, err)
	require.Len(t, sections, 3)
}
