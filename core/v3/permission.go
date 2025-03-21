package v3

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type ParameterOperation uint8

const (
	EQUAL                 ParameterOperation = 0
	NOT_EQUAL             ParameterOperation = 1
	GREATER_THAN_OR_EQUAL ParameterOperation = 2
	LESS_THAN_OR_EQUAL    ParameterOperation = 3
)

type ParameterRule struct {
	Cumulative bool               `json:"cumulative"`
	Operation  ParameterOperation `json:"operation"`
	Value      []byte             `json:"value"`
	Offset     *big.Int           `json:"offset"`
	Mask       []byte             `json:"mask"`
}

type Permission struct {
	Target common.Address  `json:"target"`
	Rules  []ParameterRule `json:"rules"`
}

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

// EncodeSessionPermissions encodes session permissions into bytes
func EncodeSessionPermissions(sp *SessionPermissions) ([]byte, error) {
	if len(sp.Permissions) > MAX_PERMISSIONS_COUNT {
		return nil, fmt.Errorf("too many permissions")
	}

	var result []byte
	result = append(result, common.LeftPadBytes(sp.Signer.Bytes(), 20)...)
	result = append(result, common.LeftPadBytes(sp.ValueLimit.Bytes(), 32)...)
	result = append(result, common.LeftPadBytes(sp.Deadline.Bytes(), 32)...)
	result = append(result, byte(len(sp.Permissions)))

	for _, perm := range sp.Permissions {
		encoded, err := EncodePermission(&perm)
		if err != nil {
			return nil, err
		}
		result = append(result, encoded...)
	}

	return result, nil
}

// EncodePermission encodes a permission into bytes
func EncodePermission(p *Permission) ([]byte, error) {
	if len(p.Rules) > MAX_RULES_COUNT {
		return nil, fmt.Errorf("too many rules")
	}

	var result []byte
	result = append(result, common.LeftPadBytes(p.Target.Bytes(), 20)...)
	result = append(result, byte(len(p.Rules)))

	for _, rule := range p.Rules {
		encoded := encodeParameterRule(&rule)
		result = append(result, encoded...)
	}

	return result, nil
}

// encodeParameterRule encodes a parameter rule into bytes
func encodeParameterRule(r *ParameterRule) []byte {
	// Combine operation and cumulative flag into a single byte
	// 0x[operationx3][cumulative]
	operationCumulative := byte(r.Operation<<1) | boolToByte(r.Cumulative)

	var result []byte
	result = append(result, operationCumulative)
	result = append(result, common.LeftPadBytes(r.Value, 32)...)
	result = append(result, common.LeftPadBytes(r.Offset.Bytes(), 32)...)
	result = append(result, common.LeftPadBytes(r.Mask, 32)...)

	return result
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

// SessionPermissionsFromJSON decodes session permissions from JSON
func SessionPermissionsFromJSON(data string) (*SessionPermissions, error) {
	var sp SessionPermissions
	if err := json.Unmarshal([]byte(data), &sp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session permissions: %w", err)
	}
	return &sp, nil
}

// PermissionFromJSON decodes a permission from JSON
func PermissionFromJSON(data string) (*Permission, error) {
	var p Permission
	if err := json.Unmarshal([]byte(data), &p); err != nil {
		return nil, fmt.Errorf("failed to unmarshal permission: %w", err)
	}
	return &p, nil
}

// UnmarshalJSON implements json.Unmarshaler for SessionPermissions
func (sp *SessionPermissions) UnmarshalJSON(data []byte) error {
	type Alias SessionPermissions
	aux := struct {
		ValueLimit json.RawMessage `json:"valueLimit"`
		Deadline   json.RawMessage `json:"deadline"`
		*Alias
	}{
		Alias: (*Alias)(sp),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse ValueLimit
	if len(aux.ValueLimit) > 0 {
		valueLimit := new(big.Int)
		var strVal string
		if err := json.Unmarshal(aux.ValueLimit, &strVal); err != nil {
			// If string unmarshal fails, try as number
			var numVal json.Number
			if err := json.Unmarshal(aux.ValueLimit, &numVal); err != nil {
				return fmt.Errorf("invalid valueLimit format: %w", err)
			}
			strVal = numVal.String()
		}
		if _, ok := valueLimit.SetString(strVal, 0); !ok {
			return fmt.Errorf("invalid valueLimit: %s", strVal)
		}
		sp.ValueLimit = valueLimit
	}

	// Parse Deadline
	if len(aux.Deadline) > 0 {
		deadline := new(big.Int)
		var strVal string
		if err := json.Unmarshal(aux.Deadline, &strVal); err != nil {
			// If string unmarshal fails, try as number
			var numVal json.Number
			if err := json.Unmarshal(aux.Deadline, &numVal); err != nil {
				return fmt.Errorf("invalid deadline format: %w", err)
			}
			strVal = numVal.String()
		}
		if _, ok := deadline.SetString(strVal, 0); !ok {
			return fmt.Errorf("invalid deadline: %s", strVal)
		}
		sp.Deadline = deadline
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler for ParameterRule
func (pr *ParameterRule) UnmarshalJSON(data []byte) error {
	type Alias ParameterRule
	aux := struct {
		Value  string          `json:"value"`
		Offset json.RawMessage `json:"offset"`
		Mask   string          `json:"mask"`
		*Alias
	}{
		Alias: (*Alias)(pr),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse Value
	if aux.Value != "" {
		value := common.FromHex(aux.Value)
		pr.Value = value
	}

	// Parse Offset
	if len(aux.Offset) > 0 {
		offset := new(big.Int)
		var strVal string
		if err := json.Unmarshal(aux.Offset, &strVal); err != nil {
			// If string unmarshal fails, try as number
			var numVal json.Number
			if err := json.Unmarshal(aux.Offset, &numVal); err != nil {
				return fmt.Errorf("invalid offset format: %w", err)
			}
			strVal = numVal.String()
		}
		if _, ok := offset.SetString(strVal, 0); !ok {
			return fmt.Errorf("invalid offset: %s", strVal)
		}
		pr.Offset = offset
	}

	// Parse Mask
	if aux.Mask != "" {
		mask := common.FromHex(aux.Mask)
		pr.Mask = mask
	}

	return nil
}

// MarshalJSON implements json.Marshaler for ParameterRule
func (pr *ParameterRule) MarshalJSON() ([]byte, error) {
	type parameterRuleType struct {
		Cumulative bool               `json:"cumulative"`
		Operation  ParameterOperation `json:"operation"`
		Value      string             `json:"value"`
		Offset     string             `json:"offset"`
		Mask       string             `json:"mask"`
	}

	return json.Marshal(parameterRuleType{
		Cumulative: pr.Cumulative,
		Operation:  pr.Operation,
		Value:      "0x" + hex.EncodeToString(pr.Value),
		Offset:     pr.Offset.String(),
		Mask:       "0x" + hex.EncodeToString(pr.Mask),
	})
}

// MarshalJSON implements json.Marshaler for Permission
func (p *Permission) MarshalJSON() ([]byte, error) {
	type permissionType struct {
		Target string          `json:"target"`
		Rules  []ParameterRule `json:"rules"`
	}

	return json.Marshal(permissionType{
		Target: p.Target.Hex(),
		Rules:  p.Rules,
	})
}
