package v3

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// ParameterOperation represents the comparison operation for a parameter rule.
type ParameterOperation uint8

const (
	EQUAL ParameterOperation = iota
	NOT_EQUAL
	GREATER_THAN_OR_EQUAL
	LESS_THAN_OR_EQUAL
)

// ParameterRule defines one rule for comparing a parameter.
type ParameterRule struct {
	Cumulative bool               `json:"cumulative"`
	Operation  ParameterOperation `json:"operation"`
	Value      []byte             `json:"value"`
	Offset     *big.Int           `json:"offset"`
	Mask       []byte             `json:"mask"`
}

// Permission represents a permission with a target address and a set of parameter rules.
type Permission struct {
	Target common.Address  `json:"target"`
	Rules  []ParameterRule `json:"rules"`
}

// SessionPermissions groups a signer with its associated permissions.
type SessionPermissions struct {
	Signer      common.Address `json:"signer"`
	ValueLimit  *big.Int       `json:"valueLimit"`
	Deadline    *big.Int       `json:"deadline"`
	Permissions []Permission   `json:"permissions"`
}

const (
	MAX_PERMISSIONS_COUNT = (1 << 7) - 1
	MAX_RULES_COUNT       = (1 << 8) - 1
)

// -----------------------------------------------------------------------------
// Binary Encoding
// -----------------------------------------------------------------------------

// EncodeSessionPermissions encodes a SessionPermissions structure into bytes.
func EncodeSessionPermissions(sp *SessionPermissions) ([]byte, error) {
	if len(sp.Permissions) > MAX_PERMISSIONS_COUNT {
		return nil, fmt.Errorf("too many permissions")
	}
	var result []byte
	// Append signer as 20-byte left‐padded value.
	result = append(result, LeftPad(sp.Signer.Bytes(), 20)...)
	// Append valueLimit (32 bytes) and deadline (32 bytes).
	result = append(result, LeftPad(sp.ValueLimit.Bytes(), 32)...)
	result = append(result, LeftPad(sp.Deadline.Bytes(), 32)...)
	// Append a single byte with the number of permissions.
	result = append(result, byte(len(sp.Permissions)))
	// Encode each permission.
	for _, perm := range sp.Permissions {
		encoded, err := EncodePermission(&perm)
		if err != nil {
			return nil, err
		}
		result = append(result, encoded...)
	}
	return result, nil
}

// EncodePermission encodes a Permission into bytes.
func EncodePermission(p *Permission) ([]byte, error) {
	if len(p.Rules) > MAX_RULES_COUNT {
		return nil, fmt.Errorf("too many rules")
	}
	var result []byte
	// Append target (20 bytes).
	result = append(result, LeftPad(p.Target.Bytes(), 20)...)
	// Append a single byte with the number of rules.
	result = append(result, byte(len(p.Rules)))
	// Append each encoded parameter rule.
	for _, rule := range p.Rules {
		result = append(result, encodeParameterRule(&rule)...)
	}
	return result, nil
}

// encodeParameterRule encodes a ParameterRule into bytes.
// The first byte packs the 3‑bit operation (shifted left by 1) and 1‑bit cumulative flag.
func encodeParameterRule(r *ParameterRule) []byte {
	opCumulative := byte(r.Operation<<1) | boolToByte(r.Cumulative)
	var result []byte
	result = append(result, opCumulative)
	result = append(result, LeftPad(r.Value, 32)...)
	result = append(result, LeftPad(r.Offset.Bytes(), 32)...)
	result = append(result, LeftPad(r.Mask, 32)...)
	return result
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// DecodeSessionPermissions decodes a byte slice into a SessionPermissions structure.
func DecodeSessionPermissions(b []byte) (SessionPermissions, error) {
	if len(b) < 85 {
		return SessionPermissions{}, fmt.Errorf("insufficient bytes for session permissions")
	}
	var sp SessionPermissions
	sp.Signer = common.BytesToAddress(b[0:20])
	sp.ValueLimit = new(big.Int).SetBytes(b[20:52])
	sp.Deadline = new(big.Int).SetBytes(b[52:84])
	permCount := int(b[84])
	ptr := 85
	var perms []Permission
	for i := 0; i < permCount; i++ {
		perm, consumed, err := decodePermission(b[ptr:])
		if err != nil {
			return SessionPermissions{}, err
		}
		perms = append(perms, perm)
		ptr += consumed
	}
	sp.Permissions = perms
	return sp, nil
}

// decodePermission decodes a Permission from bytes. It returns the Permission, the number of bytes consumed, and an error if any.
func decodePermission(b []byte) (Permission, int, error) {
	if len(b) < 21 {
		return Permission{}, 0, fmt.Errorf("insufficient bytes for permission")
	}
	var perm Permission
	perm.Target = common.BytesToAddress(b[0:20])
	rulesCount := int(b[20])
	ptr := 21
	var rules []ParameterRule
	for i := 0; i < rulesCount; i++ {
		if len(b) < ptr+97 {
			return Permission{}, 0, fmt.Errorf("insufficient bytes for parameter rule")
		}
		rule := decodeParameterRule(b[ptr : ptr+97])
		rules = append(rules, rule)
		ptr += 97
	}
	perm.Rules = rules
	return perm, ptr, nil
}

// decodeParameterRule decodes a 97‑byte slice into a ParameterRule.
func decodeParameterRule(b []byte) ParameterRule {
	opCumulative := b[0]
	cumulative := (opCumulative & 1) == 1
	operation := opCumulative >> 1
	value := b[1:33]
	offset := new(big.Int).SetBytes(b[33:65])
	mask := b[65:97]
	return ParameterRule{
		Cumulative: cumulative,
		Operation:  ParameterOperation(operation),
		Value:      value,
		Offset:     offset,
		Mask:       mask,
	}
}

// -----------------------------------------------------------------------------
// ABI Encoding
// -----------------------------------------------------------------------------

// permissionStructAbi defines the ABI tuple for Permission.
var permissionStructAbi = []abi.ArgumentMarshaling{
	{
		Name: "target",
		Type: "address",
	},
	{
		Name: "rules",
		Type: "tuple[]",
		Components: []abi.ArgumentMarshaling{
			{Name: "cumulative", Type: "bool"},
			{Name: "operation", Type: "uint8"},
			{Name: "value", Type: "bytes32"},
			{Name: "offset", Type: "uint256"},
			{Name: "mask", Type: "bytes32"},
		},
	},
}

// AbiEncodePermission packs a Permission into an ABI-encoded hex string.
func AbiEncodePermission(p *Permission) (string, error) {
	// Define internal types for ABI encoding.
	type ruleABI struct {
		Cumulative bool
		Operation  uint8
		Value      [32]byte
		Offset     *big.Int
		Mask       [32]byte
	}
	type permissionABI struct {
		Target common.Address
		Rules  []ruleABI
	}
	var rules []ruleABI
	for _, r := range p.Rules {
		var val [32]byte
		copy(val[:], LeftPad(r.Value, 32))
		var mask [32]byte
		copy(mask[:], LeftPad(r.Mask, 32))
		rules = append(rules, ruleABI{
			Cumulative: r.Cumulative,
			Operation:  uint8(r.Operation),
			Value:      val,
			Offset:     r.Offset,
			Mask:       mask,
		})
	}
	perm := permissionABI{
		Target: p.Target,
		Rules:  rules,
	}
	// Build ABI arguments.
	args, err := argumentsFromABIArguments(permissionStructAbi)
	if err != nil {
		return "", fmt.Errorf("failed to build ABI arguments: %w", err)
	}
	packed, err := args.Pack(perm.Target, perm.Rules)
	if err != nil {
		return "", fmt.Errorf("failed to pack permission: %w", err)
	}
	return "0x" + hex.EncodeToString(packed), nil
}

// argumentsFromABIArguments converts a slice of ArgumentMarshaling into abi.Arguments.
func argumentsFromABIArguments(argMarsh []abi.ArgumentMarshaling) (abi.Arguments, error) {
	var args abi.Arguments
	for _, am := range argMarsh {
		t, err := abi.NewType(am.Type, "", am.Components)
		if err != nil {
			return nil, err
		}
		args = append(args, abi.Argument{
			Name: am.Name,
			Type: t,
		})
	}
	return args, nil
}

// -----------------------------------------------------------------------------
// JSON Conversion
// -----------------------------------------------------------------------------

// SessionPermissionsToJSON converts SessionPermissions to its JSON string representation.
func (sp *SessionPermissions) MarshalJSON() ([]byte, error) {
	enc := encodeSessionPermissionsForJson(sp)
	return json.Marshal(enc)
}

// encodeSessionPermissionsForJson returns a map representing SessionPermissions.
func encodeSessionPermissionsForJson(sp *SessionPermissions) map[string]interface{} {
	perms := make([]interface{}, 0, len(sp.Permissions))
	for _, p := range sp.Permissions {
		perms = append(perms, encodePermissionForJson(&p))
	}
	return map[string]interface{}{
		"signer":      sp.Signer.Hex(),
		"valueLimit":  sp.ValueLimit.String(),
		"deadline":    sp.Deadline.String(),
		"permissions": perms,
	}
}

// PermissionToJSON converts a Permission to its JSON string representation.
func (p *Permission) MarshalJSON() ([]byte, error) {
	enc := encodePermissionForJson(p)
	return json.Marshal(enc)
}

// encodePermissionForJson returns a map representing a Permission.
func encodePermissionForJson(p *Permission) map[string]interface{} {
	rules := make([]interface{}, 0, len(p.Rules))
	for _, r := range p.Rules {
		rules = append(rules, encodeParameterRuleForJson(r))
	}
	return map[string]interface{}{
		"target": p.Target.Hex(),
		"rules":  rules,
	}
}

// ParameterRuleToJSON converts a ParameterRule to a JSON string.
func (r *ParameterRule) MarshalJSON() ([]byte, error) {
	enc := encodeParameterRuleForJson(*r)
	return json.Marshal(enc)
}

// encodeParameterRuleForJson returns a map representing a ParameterRule.
func encodeParameterRuleForJson(r ParameterRule) map[string]interface{} {
	return map[string]interface{}{
		"cumulative": r.Cumulative,
		"operation":  r.Operation,
		"value":      "0x" + hex.EncodeToString(r.Value),
		"offset":     r.Offset.String(),
		"mask":       "0x" + hex.EncodeToString(r.Mask),
	}
}

// SessionPermissionsFromJSON parses a JSON string into a SessionPermissions struct.
func (sp *SessionPermissions) UnmarshalJSON(data []byte) error {
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return fmt.Errorf("failed to unmarshal session permissions JSON: %w", err)
	}
	*sp = sessionPermissionsFromParsed(parsed)
	return nil
}

// sessionPermissionsFromParsed converts a parsed JSON object (map) into SessionPermissions.
func sessionPermissionsFromParsed(parsed interface{}) SessionPermissions {
	m, ok := parsed.(map[string]interface{})
	if !ok {
		panic("invalid type for session permissions")
	}
	return SessionPermissions{
		Signer:     common.HexToAddress(m["signer"].(string)),
		ValueLimit: valueToBigInt(m["valueLimit"]),
		Deadline:   valueToBigInt(m["deadline"]),
		Permissions: func() []Permission {
			raw := m["permissions"].([]interface{})
			perms := make([]Permission, len(raw))
			for i, v := range raw {
				perms[i] = permissionFromParsed(v)
			}
			return perms
		}(),
	}
}

// PermissionFromJSON parses a JSON string into a Permission.
func (p *Permission) UnmarshalJSON(data []byte) error {
	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return fmt.Errorf("failed to unmarshal permission JSON: %w", err)
	}
	*p = permissionFromParsed(parsed)
	return nil
}

// permissionFromParsed converts a parsed JSON object (map) into a Permission.
func permissionFromParsed(parsed interface{}) Permission {
	m, ok := parsed.(map[string]interface{})
	if !ok {
		panic("invalid type for permission")
	}
	return Permission{
		Target: common.HexToAddress(m["target"].(string)),
		Rules: func() []ParameterRule {
			raw := m["rules"].([]interface{})
			rules := make([]ParameterRule, len(raw))
			for i, v := range raw {
				rules[i] = parameterRuleFromParsed(v)
			}
			return rules
		}(),
	}
}

// parameterRuleFromParsed converts a parsed JSON object into a ParameterRule.
func parameterRuleFromParsed(parsed interface{}) ParameterRule {
	m, ok := parsed.(map[string]interface{})
	if !ok {
		panic("invalid type for parameter rule")
	}
	valueStr := strings.TrimPrefix(m["value"].(string), "0x")
	maskStr := strings.TrimPrefix(m["mask"].(string), "0x")

	return ParameterRule{
		Cumulative: m["cumulative"].(bool),
		Operation:  ParameterOperation(uint8(m["operation"].(float64))),
		Value:      mustDecodeHex(valueStr),
		Offset:     valueToBigInt(m["offset"]),
		Mask:       mustDecodeHex(maskStr),
	}
}

// mustDecodeHex decodes a hex string and panics if it fails.
func mustDecodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// valueToBigInt converts a string (or number) to *big.Int.
func valueToBigInt(v interface{}) *big.Int {
	switch val := v.(type) {
	case string:
		i, ok := new(big.Int).SetString(val, 10)
		if !ok {
			panic("cannot convert string to big.Int: " + val)
		}
		return i
	case float64:
		return big.NewInt(int64(val))
	case int64:
		return big.NewInt(val)
	case int:
		return big.NewInt(int64(val))
	default:
		s := fmt.Sprintf("%v", v)
		i, ok := new(big.Int).SetString(s, 10)
		if !ok {
			panic("cannot convert to big.Int: " + s)
		}
		return i
	}
}

// -----------------------------------------------------------------------------
// Helper Functions
// -----------------------------------------------------------------------------

// LeftPad returns a new byte slice padded on the left with zeros to the given length.
func LeftPad(b []byte, length int) []byte {
	if len(b) >= length {
		return b
	}
	pad := make([]byte, length-len(b))
	return append(pad, b...)
}
