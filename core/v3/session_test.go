package v3

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

func TestMarshalJSON(t *testing.T) {
	globalSigner := common.HexToAddress("0x74214e263d5669885065f215Fa21BC940588f7C1")
	topology := EmptySessionsTopology(globalSigner)
	json, err := SessionsTopologyToJSON(topology)
	if err != nil {
		t.Errorf("Failed to marshal SessionsTopology: %v", err)
		return
	}

	expected := `[{"blacklist":[]},[{"globalSigner":"0x74214e263d5669885065f215Fa21BC940588f7C1"}]]`
	if json != expected {
		t.Errorf("Unexpected JSON output: %s", json)
	}
}

func TestSessionExplicitAdd(t *testing.T) {
	signer := common.HexToAddress("0x1111111111111111111111111111111111111111")
	target := common.HexToAddress("0x2222222222222222222222222222222222222222")

	valueLimit := big.NewInt(100)
	deadline := big.NewInt(1678886400)

	rule := ParameterRule{
		Cumulative: true,
		Operation:  EQUAL,
		Value:      []byte{0x01},
		Offset:     big.NewInt(0),
		Mask:       []byte{0xff},
	}

	permission := Permission{
		Target: target,
		Rules:  []ParameterRule{rule},
	}

	permissions := SessionPermissions{
		Signer:      signer,
		ValueLimit:  valueLimit,
		Deadline:    deadline,
		Permissions: []Permission{permission},
	}

	// Marshal to JSON
	permissionsJSON, err := json.Marshal(permissions)
	if err != nil {
		t.Fatalf("Failed to marshal SessionPermissions: %v", err)
	}

	// Unmarshal from JSON
	var unmarshaledPermissions SessionPermissions
	err = json.Unmarshal(permissionsJSON, &unmarshaledPermissions)
	if err != nil {
		t.Fatalf("Failed to unmarshal SessionPermissions: %v", err)
	}

	// Check if the original and unmarshaled permissions are the same
	if permissions.Signer != unmarshaledPermissions.Signer {
		t.Errorf("Signer mismatch: original=%v, unmarshaled=%v", permissions.Signer, unmarshaledPermissions.Signer)
	}

	if permissions.ValueLimit.Cmp(unmarshaledPermissions.ValueLimit) != 0 {
		t.Errorf("ValueLimit mismatch: original=%v, unmarshaled=%v", permissions.ValueLimit, unmarshaledPermissions.ValueLimit)
	}

	if permissions.Deadline.Cmp(unmarshaledPermissions.Deadline) != 0 {
		t.Errorf("Deadline mismatch: original=%v, unmarshaled=%v", permissions.Deadline, unmarshaledPermissions.Deadline)
	}

	if len(permissions.Permissions) != len(unmarshaledPermissions.Permissions) {
		t.Errorf("Permissions length mismatch: original=%d, unmarshaled=%d", len(permissions.Permissions), len(unmarshaledPermissions.Permissions))
	}

	if permissions.Permissions[0].Target != unmarshaledPermissions.Permissions[0].Target {
		t.Errorf("Permission Target mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Target, unmarshaledPermissions.Permissions[0].Target)
	}

	// Check ParameterRule
	if permissions.Permissions[0].Rules[0].Cumulative != unmarshaledPermissions.Permissions[0].Rules[0].Cumulative {
		t.Errorf("ParameterRule Cumulative mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Rules[0].Cumulative, unmarshaledPermissions.Permissions[0].Rules[0].Cumulative)
	}

	if permissions.Permissions[0].Rules[0].Operation != unmarshaledPermissions.Permissions[0].Rules[0].Operation {
		t.Errorf("ParameterRule Operation mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Rules[0].Operation, unmarshaledPermissions.Permissions[0].Rules[0].Operation)
	}

	if string(permissions.Permissions[0].Rules[0].Value) != string(unmarshaledPermissions.Permissions[0].Rules[0].Value) {
		t.Errorf("ParameterRule Value mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Rules[0].Value, unmarshaledPermissions.Permissions[0].Rules[0].Value)
	}

	if permissions.Permissions[0].Rules[0].Offset.Cmp(unmarshaledPermissions.Permissions[0].Rules[0].Offset) != 0 {
		t.Errorf("ParameterRule Offset mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Rules[0].Offset, unmarshaledPermissions.Permissions[0].Rules[0].Offset)
	}

	if string(permissions.Permissions[0].Rules[0].Mask) != string(unmarshaledPermissions.Permissions[0].Rules[0].Mask) {
		t.Errorf("ParameterRule Mask mismatch: original=%v, unmarshaled=%v", permissions.Permissions[0].Rules[0].Mask, unmarshaledPermissions.Permissions[0].Rules[0].Mask)
	}
}

func TestSessionPermissionsFromJSON(t *testing.T) {
	validJSON := `{
		"signer": "0x9ed233eCAE5E093CAff8Ff8E147DdAfc704EC619",
		"valueLimit": 1000,
		"deadline": 2000,
		"permissions": [
			{
				"target": "0x000000000000000000000000000000000000bEEF",
				"rules": [
					{
						"cumulative": false,
						"operation": 0,
						"value": "0x0000000000000000000000000000000000000000000000000000000000000000",
						"offset": "0",
						"mask": "0x0000000000000000000000000000000000000000000000000000000000000000"
					}
				]
			}
		]
	}`

	var session SessionPermissions
	err := json.Unmarshal([]byte(validJSON), &session)
	if err != nil {
		t.Fatalf("Failed to unmarshal SessionPermissions: %v", err)
	}

	invalidJSON := `{
		"signer": "0x9ed233eCAE5E093CAff8Ff8E147DdAfc704EC619",
		"valueLimit": 1000,
		"deadline": 2000,
		"permissions": "invalid data"
	}`

	err = json.Unmarshal([]byte(invalidJSON), &session)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
}
