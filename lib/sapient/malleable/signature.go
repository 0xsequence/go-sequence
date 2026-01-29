package malleable

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type SectionKind uint8

const (
	SectionStatic SectionKind = iota
	SectionRepeat
)

type Section interface {
	Kind() SectionKind
}

type StaticSection struct {
	TIndex uint8
	CIndex uint16
	Size   uint16
}

func (s StaticSection) Kind() SectionKind { return SectionStatic }

type RepeatSection struct {
	TIndex  uint8
	CIndex  uint16
	Size    uint16
	TIndex2 uint8
	CIndex2 uint16
}

func (s RepeatSection) Kind() SectionKind { return SectionRepeat }

func EncodeSignature(statics []StaticSection, repeats []RepeatSection) ([]byte, error) {
	var out bytes.Buffer

	for _, s := range statics {
		if s.TIndex > 0x7F {
			return nil, fmt.Errorf("tindex out of range: %d", s.TIndex)
		}
		out.WriteByte(byte(s.TIndex & 0x7F))
		_ = binary.Write(&out, binary.BigEndian, s.CIndex)
		_ = binary.Write(&out, binary.BigEndian, s.Size)
	}

	for _, r := range repeats {
		if r.TIndex > 0x7F {
			return nil, fmt.Errorf("tindex out of range: %d", r.TIndex)
		}
		out.WriteByte(byte(r.TIndex&0x7F) | 0x80)
		_ = binary.Write(&out, binary.BigEndian, r.CIndex)
		_ = binary.Write(&out, binary.BigEndian, r.Size)
		out.WriteByte(r.TIndex2)
		_ = binary.Write(&out, binary.BigEndian, r.CIndex2)
	}

	return out.Bytes(), nil
}

func DecodeSignature(sig []byte) ([]Section, error) {
	var sections []Section
	i := 0
	for i < len(sig) {
		if i+5 > len(sig) {
			return nil, fmt.Errorf("signature truncated at %d", i)
		}
		tRaw := sig[i]
		i++
		cindex := binary.BigEndian.Uint16(sig[i : i+2])
		i += 2
		size := binary.BigEndian.Uint16(sig[i : i+2])
		i += 2

		if tRaw&0x80 != 0 {
			if i+3 > len(sig) {
				return nil, fmt.Errorf("repeat truncated at %d", i)
			}
			t2 := sig[i]
			i++
			c2 := binary.BigEndian.Uint16(sig[i : i+2])
			i += 2
			sections = append(sections, RepeatSection{
				TIndex:  tRaw & 0x7F,
				CIndex:  cindex,
				Size:    size,
				TIndex2: t2,
				CIndex2: c2,
			})
		} else {
			sections = append(sections, StaticSection{
				TIndex: tRaw & 0x7F,
				CIndex: cindex,
				Size:   size,
			})
		}
	}
	return sections, nil
}
