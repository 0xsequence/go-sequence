package malleable

import (
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

type BuilderOptions struct {
	ValidateRepeats     bool
	MergeAdjacentStatic bool
	MaxOffset           uint32
	MaxSize             uint32
}

type Builder struct {
	payload   *v3.CallsPayload
	options   BuilderOptions
	malleable []Selector
	repeats   []repeatSelector
}

type repeatSelector struct {
	a Selector
	b Selector
}

type Plan struct {
	Static []StaticSection
	Repeat []RepeatSection
}

func (p *Plan) DebugString() string {
	out := "static:"
	for _, s := range p.Static {
		out += fmt.Sprintf(" [t=%d c=%d s=%d]", s.TIndex, s.CIndex, s.Size)
	}
	out += " repeat:"
	for _, r := range p.Repeat {
		out += fmt.Sprintf(" [t=%d c=%d s=%d t2=%d c2=%d]", r.TIndex, r.CIndex, r.Size, r.TIndex2, r.CIndex2)
	}
	return out
}

func NewBuilder(payload *v3.CallsPayload, opts *BuilderOptions) *Builder {
	options := BuilderOptions{
		MaxOffset:           0xFFFF,
		MaxSize:             0xFFFF,
		ValidateRepeats:     false,
		MergeAdjacentStatic: true,
	}
	if opts != nil {
		if opts.MaxOffset != 0 {
			options.MaxOffset = opts.MaxOffset
		}
		if opts.MaxSize != 0 {
			options.MaxSize = opts.MaxSize
		}
		options.ValidateRepeats = opts.ValidateRepeats
		options.MergeAdjacentStatic = opts.MergeAdjacentStatic
	}

	return &Builder{
		payload: payload,
		options: options,
	}
}

func (b *Builder) Malleable(sel Selector) *Builder {
	b.malleable = append(b.malleable, sel)
	return b
}

func (b *Builder) Repeat(a Selector, b2 Selector) *Builder {
	b.repeats = append(b.repeats, repeatSelector{a: a, b: b2})
	return b
}

func (b *Builder) Build() ([]byte, *Plan, error) {
	if b.payload == nil {
		return nil, nil, fmt.Errorf("payload is nil")
	}
	if len(b.payload.Calls) > 128 {
		return nil, nil, fmt.Errorf("too many calls (%d): tindex is 7-bit", len(b.payload.Calls))
	}

	exByCall := make([][]Span, len(b.payload.Calls))

	addExclude := func(r ByteRange) error {
		if r.CallIndex < 0 || r.CallIndex >= len(b.payload.Calls) {
			return fmt.Errorf("tindex out of range: %d", r.CallIndex)
		}
		dataLen := len(b.payload.Calls[r.CallIndex].Data)
		if r.Offset < 0 || r.Size < 0 || r.Offset+r.Size > dataLen {
			return fmt.Errorf("span out of bounds: [%d,%d) > %d", r.Offset, r.Offset+r.Size, dataLen)
		}
		exByCall[r.CallIndex] = append(exByCall[r.CallIndex], Span{Start: r.Offset, Len: r.Size})
		return nil
	}

	for _, sel := range b.malleable {
		ranges, err := sel.Resolve(b.payload)
		if err != nil {
			return nil, nil, fmt.Errorf("malleable %s: %w", sel.String(), err)
		}
		for _, r := range ranges {
			if err := addExclude(r); err != nil {
				return nil, nil, fmt.Errorf("malleable %s: %w", sel.String(), err)
			}
		}
	}

	var repeats []RepeatSection
	for _, rp := range b.repeats {
		rangesA, err := rp.a.Resolve(b.payload)
		if err != nil {
			return nil, nil, fmt.Errorf("repeat a %s: %w", rp.a.String(), err)
		}
		rangesB, err := rp.b.Resolve(b.payload)
		if err != nil {
			return nil, nil, fmt.Errorf("repeat b %s: %w", rp.b.String(), err)
		}
		if len(rangesA) != 1 || len(rangesB) != 1 {
			return nil, nil, fmt.Errorf("repeat expects single range per selector")
		}
		a := rangesA[0]
		bb := rangesB[0]
		if a.Size != bb.Size {
			return nil, nil, fmt.Errorf("repeat size mismatch: %d vs %d", a.Size, bb.Size)
		}
		if err := addExclude(a); err != nil {
			return nil, nil, fmt.Errorf("repeat a: %w", err)
		}
		if err := addExclude(bb); err != nil {
			return nil, nil, fmt.Errorf("repeat b: %w", err)
		}
		if b.options.ValidateRepeats {
			sectionA, err := a.Slice(b.payload)
			if err != nil {
				return nil, nil, err
			}
			sectionB, err := bb.Slice(b.payload)
			if err != nil {
				return nil, nil, err
			}
			if crypto.Keccak256Hash(sectionA) != crypto.Keccak256Hash(sectionB) {
				return nil, nil, fmt.Errorf("repeat section mismatch")
			}
		}
		repeats = append(repeats, RepeatSection{
			TIndex:  uint8(a.CallIndex),
			CIndex:  uint16(a.Offset),
			Size:    uint16(a.Size),
			TIndex2: uint8(bb.CallIndex),
			CIndex2: uint16(bb.Offset),
		})
	}

	var statics []SpanWithCall
	for t := range b.payload.Calls {
		length := len(b.payload.Calls[t].Data)
		ex := mergeSpans(exByCall[t])
		cursor := 0
		for _, s := range ex {
			if cursor < s.Start {
				statics = append(statics, SpanWithCall{CallIndex: t, Span: Span{Start: cursor, Len: s.Start - cursor}})
			}
			cursor = max(cursor, s.Start+s.Len)
		}
		if cursor < length {
			statics = append(statics, SpanWithCall{CallIndex: t, Span: Span{Start: cursor, Len: length - cursor}})
		}
	}

	sort.Slice(statics, func(i, j int) bool {
		if statics[i].CallIndex != statics[j].CallIndex {
			return statics[i].CallIndex < statics[j].CallIndex
		}
		return statics[i].Start < statics[j].Start
	})

	if b.options.MergeAdjacentStatic {
		statics = mergeAdjacentStatics(statics)
	}

	staticSections, err := b.encodeStaticSections(statics)
	if err != nil {
		return nil, nil, err
	}

	signature, err := EncodeSignature(staticSections, repeats)
	if err != nil {
		return nil, nil, err
	}

	return signature, &Plan{Static: staticSections, Repeat: repeats}, nil
}

type SpanWithCall struct {
	CallIndex int
	Span
}

func (b *Builder) encodeStaticSections(statics []SpanWithCall) ([]StaticSection, error) {
	var sections []StaticSection
	for _, s := range statics {
		if s.Len == 0 {
			continue
		}
		if s.CallIndex < 0 || s.CallIndex >= len(b.payload.Calls) {
			return nil, fmt.Errorf("tindex out of range: %d", s.CallIndex)
		}
		offset := s.Start
		length := s.Len
		for length > 0 {
			chunk := length
			if uint32(chunk) > b.options.MaxSize {
				chunk = int(b.options.MaxSize)
			}
			if uint32(offset) > b.options.MaxOffset {
				return nil, fmt.Errorf("cindex too large: %d", offset)
			}
			sections = append(sections, StaticSection{
				TIndex: uint8(s.CallIndex),
				CIndex: uint16(offset),
				Size:   uint16(chunk),
			})
			offset += chunk
			length -= chunk
		}
	}
	return sections, nil
}

func mergeSpans(spans []Span) []Span {
	if len(spans) == 0 {
		return nil
	}
	sort.Slice(spans, func(i, j int) bool { return spans[i].Start < spans[j].Start })
	out := []Span{spans[0]}
	for _, s := range spans[1:] {
		last := &out[len(out)-1]
		if s.Start <= last.Start+last.Len {
			end := max(last.Start+last.Len, s.Start+s.Len)
			last.Len = end - last.Start
		} else {
			out = append(out, s)
		}
	}
	return out
}

func mergeAdjacentStatics(statics []SpanWithCall) []SpanWithCall {
	if len(statics) == 0 {
		return statics
	}
	out := []SpanWithCall{statics[0]}
	for _, s := range statics[1:] {
		last := &out[len(out)-1]
		if s.CallIndex == last.CallIndex && s.Start == last.Start+last.Len {
			last.Len += s.Len
		} else {
			out = append(out, s)
		}
	}
	return out
}
