package malleable

import (
	"fmt"
	"math"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	v3 "github.com/0xsequence/go-sequence/core/v3"

	"github.com/0xsequence/go-sequence/lib/abicalldata"
)

type Path struct {
	steps []pathStep
	desc  []string
}

func NewPath() *Path {
	return &Path{}
}

func (p *Path) CallData(i int) *Path {
	p.steps = append(p.steps, callDataStep{index: i})
	p.desc = append(p.desc, fmt.Sprintf("callData(%d)", i))
	return p
}

func (p *Path) Slice(offset uint32, size uint32) *Path {
	p.steps = append(p.steps, sliceStep{offset: int(offset), size: int(size)})
	p.desc = append(p.desc, fmt.Sprintf("slice(%d,%d)", offset, size))
	return p
}

func (p *Path) ABI(contractABI *abi.ABI, method string) *Path {
	p.steps = append(p.steps, abiStep{contractABI: contractABI, method: method})
	p.desc = append(p.desc, fmt.Sprintf("abi(%s)", method))
	return p
}

func (p *Path) ArgSlot(argName string) *Path {
	p.steps = append(p.steps, argSlotStep{name: argName})
	p.desc = append(p.desc, fmt.Sprintf("argSlot(%s)", argName))
	return p
}

func (p *Path) ArgSlotIndex(argIndex int) *Path {
	p.steps = append(p.steps, argSlotIndexStep{index: argIndex})
	p.desc = append(p.desc, fmt.Sprintf("argSlotIndex(%d)", argIndex))
	return p
}

func (p *Path) ArgBytesData(argName string) *Path {
	p.steps = append(p.steps, argBytesDataStep{name: argName})
	p.desc = append(p.desc, fmt.Sprintf("argBytesData(%s)", argName))
	return p
}

func (p *Path) ArgBytesDataIndex(argIndex int) *Path {
	p.steps = append(p.steps, argBytesDataIndexStep{index: argIndex})
	p.desc = append(p.desc, fmt.Sprintf("argBytesDataIndex(%d)", argIndex))
	return p
}

func (p *Path) ArgBytesEncoded(argName string) *Path {
	p.steps = append(p.steps, argBytesEncodedStep{name: argName})
	p.desc = append(p.desc, fmt.Sprintf("argBytesEncoded(%s)", argName))
	return p
}

func (p *Path) EncodedCallsPayload() *Path {
	p.steps = append(p.steps, encodedCallsPayloadStep{})
	p.desc = append(p.desc, "encodedCallsPayload()")
	return p
}

func (p *Path) EncodedCallData(i int) *Path {
	p.steps = append(p.steps, encodedCallDataStep{index: i})
	p.desc = append(p.desc, fmt.Sprintf("encodedCallData(%d)", i))
	return p
}

func (p *Path) AsSelector() abicalldata.Selector {
	return p
}

func (p *Path) Resolve(payload *v3.CallsPayload) ([]abicalldata.ByteRange, error) {
	state := pathState{}
	for _, step := range p.steps {
		if err := step.apply(payload, &state); err != nil {
			return nil, err
		}
	}
	if len(state.ranges) == 0 {
		return nil, fmt.Errorf("path resolved to empty ranges")
	}
	return state.ranges, nil
}

func (p *Path) String() string {
	if len(p.desc) == 0 {
		return "path()"
	}
	return "path(" + strings.Join(p.desc, " -> ") + ")"
}

type pathState struct {
	ranges []abicalldata.ByteRange
	// method only applies to the current full calldata frame. Any step that
	// narrows or reinterprets the active range must clear it.
	method *abi.Method
}

func (s *pathState) replaceRanges(ranges []abicalldata.ByteRange) {
	s.ranges = ranges
	s.method = nil
}

func (s *pathState) clearMethod() {
	s.method = nil
}

func (s *pathState) bindMethod(method abi.Method) {
	s.method = &method
}

func (s *pathState) requireMethod(step string) (abi.Method, error) {
	if s.method == nil {
		return abi.Method{}, fmt.Errorf("%s step requires ABI context", step)
	}
	return *s.method, nil
}

type pathStep interface {
	apply(payload *v3.CallsPayload, state *pathState) error
}

func mapRanges(ranges []abicalldata.ByteRange, fn func(abicalldata.ByteRange) (abicalldata.ByteRange, error)) ([]abicalldata.ByteRange, error) {
	out := make([]abicalldata.ByteRange, 0, len(ranges))
	for _, r := range ranges {
		next, err := fn(r)
		if err != nil {
			return nil, err
		}
		out = append(out, next)
	}
	return out, nil
}

type callDataStep struct {
	index int
}

func (s callDataStep) apply(payload *v3.CallsPayload, state *pathState) error {
	if payload == nil {
		return fmt.Errorf("payload is nil")
	}
	if s.index < 0 || s.index >= len(payload.Calls) {
		return fmt.Errorf("call index out of range: %d", s.index)
	}
	state.replaceRanges([]abicalldata.ByteRange{{
		CallIndex: s.index,
		Offset:    0,
		Size:      len(payload.Calls[s.index].Data),
	}})
	return nil
}

type sliceStep struct {
	offset int
	size   int
}

func (s sliceStep) apply(payload *v3.CallsPayload, state *pathState) error {
	if len(state.ranges) == 0 {
		return fmt.Errorf("slice step has no active ranges")
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		if s.offset < 0 || s.size < 0 {
			return abicalldata.ByteRange{}, fmt.Errorf("slice offset/size negative: %d, %d", s.offset, s.size)
		}
		if s.size > r.Size || s.offset > r.Size-s.size {
			return abicalldata.ByteRange{}, fmt.Errorf("slice out of bounds: offset=%d size=%d within %d", s.offset, s.size, r.Size)
		}
		if s.offset > math.MaxInt-r.Offset {
			return abicalldata.ByteRange{}, fmt.Errorf("slice offset overflow with range")
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + s.offset,
			Size:      s.size,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type abiStep struct {
	contractABI *abi.ABI
	method      string
}

func (s abiStep) apply(payload *v3.CallsPayload, state *pathState) error {
	if s.contractABI == nil {
		return fmt.Errorf("abi step missing contract ABI")
	}
	method, ok := s.contractABI.Methods[s.method]
	if !ok {
		return fmt.Errorf("method not found: %s", s.method)
	}
	if len(state.ranges) == 0 {
		return fmt.Errorf("abi step has no active ranges")
	}
	for _, r := range state.ranges {
		if r.Size < 4 {
			return fmt.Errorf("calldata too short for selector")
		}
		data, err := r.Slice(payload)
		if err != nil {
			return err
		}
		if len(data) < 4 {
			return fmt.Errorf("calldata too short for selector")
		}
		if !bytesEqual(data[:4], method.ID) {
			return fmt.Errorf("calldata selector mismatch for %s", s.method)
		}
	}
	state.bindMethod(method)
	return nil
}

type argSlotStep struct {
	name string
}

func (s argSlotStep) apply(payload *v3.CallsPayload, state *pathState) error {
	method, err := state.requireMethod("argSlot")
	if err != nil {
		return err
	}
	argIndex, err := argIndexByName(method, s.name)
	if err != nil {
		return err
	}
	argType := method.Inputs[argIndex].Type
	if isDynamicType(argType) {
		return fmt.Errorf("arg %s is dynamic (%s)", s.name, argType.String())
	}
	start, length, err := abicalldata.CalldataStaticWord(method, argIndex)
	if err != nil {
		return err
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		if start+length > r.Size {
			return abicalldata.ByteRange{}, fmt.Errorf("arg slot out of bounds for %s", s.name)
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + start,
			Size:      length,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type argSlotIndexStep struct {
	index int
}

func (s argSlotIndexStep) apply(payload *v3.CallsPayload, state *pathState) error {
	method, err := state.requireMethod("argSlotIndex")
	if err != nil {
		return err
	}
	argIndex := s.index
	if argIndex < 0 || argIndex >= len(method.Inputs) {
		return fmt.Errorf("arg index out of range: %d", argIndex)
	}
	argType := method.Inputs[argIndex].Type
	if isDynamicType(argType) {
		return fmt.Errorf("arg %d is dynamic (%s)", argIndex, argType.String())
	}
	start, length, err := abicalldata.CalldataStaticWord(method, argIndex)
	if err != nil {
		return err
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		if start+length > r.Size {
			return abicalldata.ByteRange{}, fmt.Errorf("arg slot out of bounds for %d", argIndex)
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + start,
			Size:      length,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type argBytesDataStep struct {
	name string
}

func (s argBytesDataStep) apply(payload *v3.CallsPayload, state *pathState) error {
	method, err := state.requireMethod("argBytesData")
	if err != nil {
		return err
	}
	argIndex, err := argIndexByName(method, s.name)
	if err != nil {
		return err
	}
	argType := method.Inputs[argIndex].Type
	if argType.T != abi.BytesTy && argType.T != abi.StringTy {
		return fmt.Errorf("arg %s is not bytes/string (%s)", s.name, argType.String())
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		data, err := r.Slice(payload)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		start, length, err := abicalldata.CalldataBytesContent(data, method, argIndex)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + start,
			Size:      length,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type argBytesDataIndexStep struct {
	index int
}

func (s argBytesDataIndexStep) apply(payload *v3.CallsPayload, state *pathState) error {
	method, err := state.requireMethod("argBytesDataIndex")
	if err != nil {
		return err
	}
	argIndex := s.index
	if argIndex < 0 || argIndex >= len(method.Inputs) {
		return fmt.Errorf("arg index out of range: %d", argIndex)
	}
	argType := method.Inputs[argIndex].Type
	if argType.T != abi.BytesTy && argType.T != abi.StringTy {
		return fmt.Errorf("arg %d is not bytes/string (%s)", argIndex, argType.String())
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		data, err := r.Slice(payload)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		start, length, err := abicalldata.CalldataBytesContent(data, method, argIndex)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + start,
			Size:      length,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type argBytesEncodedStep struct {
	name string
}

func (s argBytesEncodedStep) apply(payload *v3.CallsPayload, state *pathState) error {
	method, err := state.requireMethod("argBytesEncoded")
	if err != nil {
		return err
	}
	argIndex, err := argIndexByName(method, s.name)
	if err != nil {
		return err
	}
	argType := method.Inputs[argIndex].Type
	if argType.T != abi.BytesTy && argType.T != abi.StringTy {
		return fmt.Errorf("arg %s is not bytes/string (%s)", s.name, argType.String())
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		data, err := r.Slice(payload)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		start, length, err := abicalldata.CalldataBytesEncoded(data, method, argIndex)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + start,
			Size:      length,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

type encodedCallsPayloadStep struct{}

func (s encodedCallsPayloadStep) apply(payload *v3.CallsPayload, state *pathState) error {
	if len(state.ranges) == 0 {
		return fmt.Errorf("encodedCallsPayload step has no active ranges")
	}
	state.clearMethod()
	return nil
}

type encodedCallDataStep struct {
	index int
}

func (s encodedCallDataStep) apply(payload *v3.CallsPayload, state *pathState) error {
	if len(state.ranges) == 0 {
		return fmt.Errorf("encodedCallData step has no active ranges")
	}
	out, err := mapRanges(state.ranges, func(r abicalldata.ByteRange) (abicalldata.ByteRange, error) {
		data, err := r.Slice(payload)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		layout, err := ParsePackedCalls(data)
		if err != nil {
			return abicalldata.ByteRange{}, err
		}
		if s.index < 0 || s.index >= layout.NumCalls {
			return abicalldata.ByteRange{}, fmt.Errorf("packed call index out of range: %d", s.index)
		}
		span := layout.CallData[s.index]
		if span.Start < 0 || span.Len == 0 {
			return abicalldata.ByteRange{}, fmt.Errorf("packed call %d has no calldata", s.index)
		}
		if span.Start > math.MaxInt-r.Offset {
			return abicalldata.ByteRange{}, fmt.Errorf("packed call %d offset overflow", s.index)
		}
		return abicalldata.ByteRange{
			CallIndex: r.CallIndex,
			Offset:    r.Offset + span.Start,
			Size:      span.Len,
		}, nil
	})
	if err != nil {
		return err
	}
	state.replaceRanges(out)
	return nil
}

func argIndexByName(method abi.Method, name string) (int, error) {
	for i, input := range method.Inputs {
		if input.Name == name {
			return i, nil
		}
	}
	return -1, fmt.Errorf("arg not found: %s", name)
}

func isDynamicType(t abi.Type) bool {
	switch t.T {
	case abi.BytesTy, abi.StringTy, abi.SliceTy:
		return true
	case abi.ArrayTy:
		return t.Size == 0 || isDynamicType(*t.Elem)
	case abi.TupleTy:
		for _, elem := range t.TupleElems {
			if isDynamicType(*elem) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
