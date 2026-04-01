package hydrate

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"

	"github.com/0xsequence/go-sequence/lib/abicalldata"
)

const signalNextHydrate = 0x00

// AddrSource selects how HydrateProxy resolves an address at execution time
// (low nibble of the command flag; literal addresses add 20 bytes after the flag).
type AddrSource struct {
	kind byte
	addr common.Address
}

func SourceSelf() AddrSource                    { return AddrSource{kind: DataSelf} }
func SourceMsgSender() AddrSource               { return AddrSource{kind: DataMsgSender} }
func SourceTxOrigin() AddrSource                { return AddrSource{kind: DataTxOrigin} }
func SourceAddress(a common.Address) AddrSource { return AddrSource{kind: DataAnyAddress, addr: a} }

// Builder constructs hydratePayload bytes for HydrateProxy using abicalldata.Selector
// (calldata selectors: ABI paths via abicalldata.NewPath().AsSelector() in app code, fixed
// ranges via abicalldata.NewRangeSelector, etc.) so call-data offsets are not hand-computed.
//
// Sections are emitted in ascending call index order, which matches how HydrateProxy
// consumes the stream during sequential execution.
type Builder struct {
	payload *v3.CallsPayload
	byCall  map[int]*bytes.Buffer
}

func NewBuilder(payload *v3.CallsPayload) *Builder {
	return &Builder{
		payload: payload,
		byCall:  make(map[int]*bytes.Buffer),
	}
}

// CallSection targets one packed call index; method calls append hydrate commands for that call.
type CallSection struct {
	b      *Builder
	tindex int
}

// ForCall begins or continues the hydrate section for packed call tindex.
func (b *Builder) ForCall(tindex int) *CallSection {
	return &CallSection{b: b, tindex: tindex}
}

func (s *CallSection) buf() *bytes.Buffer {
	buf := s.b.byCall[s.tindex]
	if buf == nil {
		buf = new(bytes.Buffer)
		s.b.byCall[s.tindex] = buf
	}
	return buf
}

func (s *CallSection) validate() error {
	if s.b.payload == nil {
		return fmt.Errorf("hydrate: nil payload")
	}
	if s.tindex < 0 || s.tindex >= len(s.b.payload.Calls) {
		return fmt.Errorf("hydrate: call index %d out of range (len=%d)", s.tindex, len(s.b.payload.Calls))
	}
	if s.tindex > 255 {
		return fmt.Errorf("hydrate: call index %d does not fit hydrate tindex byte", s.tindex)
	}
	return nil
}

func resolveOneRange(payload *v3.CallsPayload, wantCall int, sel abicalldata.Selector) (abicalldata.ByteRange, error) {
	if sel == nil {
		return abicalldata.ByteRange{}, fmt.Errorf("hydrate: nil selector")
	}
	rs, err := sel.Resolve(payload)
	if err != nil {
		return abicalldata.ByteRange{}, fmt.Errorf("hydrate: %s: %w", sel.String(), err)
	}
	if len(rs) != 1 {
		return abicalldata.ByteRange{}, fmt.Errorf("hydrate: %s: want exactly one byte range, got %d", sel.String(), len(rs))
	}
	r := rs[0]
	if r.CallIndex != wantCall {
		return abicalldata.ByteRange{}, fmt.Errorf("hydrate: %s resolves to call %d, section is call %d", sel.String(), r.CallIndex, wantCall)
	}
	return r, nil
}

func toCIndex(offset int) (uint16, error) {
	if offset < 0 || offset > 65535 {
		return 0, fmt.Errorf("hydrate: calldata offset %d out of uint16 range", offset)
	}
	return uint16(offset), nil
}

func calldataLen(p *v3.CallsPayload, callIndex int) int {
	if callIndex < 0 || callIndex >= len(p.Calls) {
		return 0
	}
	return len(p.Calls[callIndex].Data)
}

func checkReplaceAddress(p *v3.CallsPayload, r abicalldata.ByteRange) (int, error) {
	n := calldataLen(p, r.CallIndex)
	switch r.Size {
	case 20:
		if r.Offset+20 > n {
			return 0, fmt.Errorf("hydrate: replaceAddress needs 20 bytes at offset %d (calldata len %d)", r.Offset, n)
		}
		return r.Offset, nil
	case 32:
		if r.Offset+32 > n {
			return 0, fmt.Errorf("hydrate: replaceAddress needs 32-byte slot at offset %d (calldata len %d)", r.Offset, n)
		}
		// ABI-encoded address values are right-aligned in a 32-byte slot.
		return r.Offset + 12, nil
	default:
		return 0, fmt.Errorf("hydrate: replaceAddress selector must resolve to 20 or 32 bytes, got %d", r.Size)
	}
}

func checkReplaceUint256(p *v3.CallsPayload, r abicalldata.ByteRange) error {
	if r.Size != 32 {
		return fmt.Errorf("hydrate: replaceUint256 selector must resolve to 32 bytes, got %d", r.Size)
	}
	n := calldataLen(p, r.CallIndex)
	if r.Offset+32 > n {
		return fmt.Errorf("hydrate: replaceUint256 needs 32 bytes at offset %d (calldata len %d)", r.Offset, n)
	}
	return nil
}

func appendAddrTail(buf *bytes.Buffer, src AddrSource) {
	if src.kind == DataAnyAddress {
		buf.Write(src.addr.Bytes())
	}
}

// DataAddress patches calldata at sel with replaceAddress (20 bytes) using src as the hydrated address.
func (s *CallSection) DataAddress(sel abicalldata.Selector, src AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	r, err := resolveOneRange(s.b.payload, s.tindex, sel)
	if err != nil {
		return err
	}
	patchOffset, err := checkReplaceAddress(s.b.payload, r)
	if err != nil {
		return err
	}
	cidx, err := toCIndex(patchOffset)
	if err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeDataAddress, src.kind))
	appendAddrTail(buf, src)
	_ = binary.Write(buf, binary.BigEndian, cidx)
	return nil
}

// DataNativeBalance patches calldata at sel with replaceUint256 using the native balance of
// the address resolved from src at execution time.
func (s *CallSection) DataNativeBalance(sel abicalldata.Selector, src AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	r, err := resolveOneRange(s.b.payload, s.tindex, sel)
	if err != nil {
		return err
	}
	if err := checkReplaceUint256(s.b.payload, r); err != nil {
		return err
	}
	cidx, err := toCIndex(r.Offset)
	if err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeDataBalance, src.kind))
	appendAddrTail(buf, src)
	_ = binary.Write(buf, binary.BigEndian, cidx)
	return nil
}

// DataERC20Balance patches calldata at sel with IERC20(token).balanceOf(holder) at execution time.
func (s *CallSection) DataERC20Balance(sel abicalldata.Selector, token common.Address, holder AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	r, err := resolveOneRange(s.b.payload, s.tindex, sel)
	if err != nil {
		return err
	}
	if err := checkReplaceUint256(s.b.payload, r); err != nil {
		return err
	}
	cidx, err := toCIndex(r.Offset)
	if err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeDataERC20Balance, holder.kind))
	appendAddrTail(buf, holder)
	_ = binary.Write(buf, binary.BigEndian, cidx)
	buf.Write(token.Bytes())
	return nil
}

// DataERC20Allowance patches calldata at sel with IERC20(token).allowance(owner, spender) at execution time.
func (s *CallSection) DataERC20Allowance(sel abicalldata.Selector, owner AddrSource, token common.Address, spender AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	r, err := resolveOneRange(s.b.payload, s.tindex, sel)
	if err != nil {
		return err
	}
	if err := checkReplaceUint256(s.b.payload, r); err != nil {
		return err
	}
	cidx, err := toCIndex(r.Offset)
	if err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeDataERC20Allowance, owner.kind))
	appendAddrTail(buf, owner)
	_ = binary.Write(buf, binary.BigEndian, cidx)
	buf.Write(token.Bytes())
	buf.WriteByte(spender.kind)
	if spender.kind == DataAnyAddress {
		buf.Write(spender.addr.Bytes())
	}
	return nil
}

// CallTo sets decoded.calls[tindex].to to the address resolved from src at execution time.
func (s *CallSection) CallTo(src AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeTo, src.kind))
	appendAddrTail(buf, src)
	return nil
}

// CallValue sets decoded.calls[tindex].value to the native balance of the address resolved from src.
func (s *CallSection) CallValue(src AddrSource) error {
	if err := s.validate(); err != nil {
		return err
	}
	buf := s.buf()
	buf.WriteByte(encodeCommandFlag(TypeValue, src.kind))
	appendAddrTail(buf, src)
	return nil
}

// Build returns hydratePayload. An empty builder yields nil, which HydrateProxy treats as "no hydration".
func (b *Builder) Build() ([]byte, error) {
	if b.payload == nil {
		return nil, fmt.Errorf("hydrate: nil payload")
	}
	if len(b.byCall) == 0 {
		return nil, nil
	}
	tindices := make([]int, 0, len(b.byCall))
	for t, buf := range b.byCall {
		if buf != nil && buf.Len() > 0 {
			tindices = append(tindices, t)
		}
	}
	if len(tindices) == 0 {
		return nil, nil
	}
	sort.Ints(tindices)
	var out bytes.Buffer
	for _, t := range tindices {
		buf := b.byCall[t]
		out.WriteByte(byte(t))
		out.Write(buf.Bytes())
		out.WriteByte(signalNextHydrate)
	}
	return out.Bytes(), nil
}
