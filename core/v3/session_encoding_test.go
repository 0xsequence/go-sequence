package v3_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// The expected byte vectors below are reference vectors verified against
// 0xsequence/wallet-contracts-v3: the SessionPermissions encoding is decoded by
// SessionSig.recoverConfiguration and the attestation hash matches
// LibAttestation.toHash (via forge tests constructing the identical structs).
// They guard the explicit-session and implicit-attestation encodings against
// silently drifting from the contracts (chainId, uint64 deadline, issuedAt).

func hexStr(b []byte) string { return "0x" + hex.EncodeToString(b) }

func sampleSessionPermissions() v3.SessionPermissions {
	return v3.SessionPermissions{
		Signer:     common.HexToAddress("0x1111111111111111111111111111111111111111"),
		ChainID:    big.NewInt(42161),
		ValueLimit: big.NewInt(1000000),
		Deadline:   big.NewInt(1678886400),
		Permissions: []v3.Permission{
			{
				Target: common.HexToAddress("0x2222222222222222222222222222222222222222"),
				Rules: []v3.ParameterRule{
					{
						Cumulative: true,
						Operation:  v3.GREATER_THAN_OR_EQUAL,
						Value:      []byte{0x01},
						Offset:     big.NewInt(0),
						Mask:       []byte{0xff},
					},
				},
			},
		},
	}
}

func TestEncodeSessionPermissionsParity(t *testing.T) {
	const expected = "0x" +
		"1111111111111111111111111111111111111111" + // signer (20)
		"000000000000000000000000000000000000000000000000000000000000a4b1" + // chainId (32) = 42161
		"00000000000000000000000000000000000000000000000000000000000f4240" + // valueLimit (32) = 1000000
		"000000006411c600" + // deadline (8, uint64) = 1678886400
		"01" + // permissions count
		"2222222222222222222222222222222222222222" + // target (20)
		"01" + // rules count
		"05" + // (GREATER_THAN_OR_EQUAL << 1) | cumulative
		"0000000000000000000000000000000000000000000000000000000000000001" + // value (32)
		"0000000000000000000000000000000000000000000000000000000000000000" + // offset (32)
		"00000000000000000000000000000000000000000000000000000000000000ff" // mask (32)

	sp := sampleSessionPermissions()
	encoded, err := v3.EncodeSessionPermissions(&sp)
	if err != nil {
		t.Fatalf("EncodeSessionPermissions: %v", err)
	}
	if got := hexStr(encoded); got != expected {
		t.Errorf("encoding mismatch:\n got %v\nwant %v", got, expected)
	}
}

func TestDecodeSessionPermissionsRoundTrip(t *testing.T) {
	sp := sampleSessionPermissions()
	encoded, err := v3.EncodeSessionPermissions(&sp)
	if err != nil {
		t.Fatalf("EncodeSessionPermissions: %v", err)
	}

	decoded, err := v3.DecodeSessionPermissions(encoded)
	if err != nil {
		t.Fatalf("DecodeSessionPermissions: %v", err)
	}

	if decoded.Signer != sp.Signer {
		t.Errorf("signer: got %v, want %v", decoded.Signer, sp.Signer)
	}
	if decoded.ChainID.Cmp(sp.ChainID) != 0 {
		t.Errorf("chainId: got %v, want %v", decoded.ChainID, sp.ChainID)
	}
	if decoded.ValueLimit.Cmp(sp.ValueLimit) != 0 {
		t.Errorf("valueLimit: got %v, want %v", decoded.ValueLimit, sp.ValueLimit)
	}
	if decoded.Deadline.Cmp(sp.Deadline) != 0 {
		t.Errorf("deadline: got %v, want %v", decoded.Deadline, sp.Deadline)
	}
	if len(decoded.Permissions) != 1 || decoded.Permissions[0].Target != sp.Permissions[0].Target {
		t.Fatalf("permissions did not round-trip: %+v", decoded.Permissions)
	}

	// Re-encoding the decoded value must be byte-stable.
	reencoded, err := v3.EncodeSessionPermissions(&decoded)
	if err != nil {
		t.Fatalf("re-EncodeSessionPermissions: %v", err)
	}
	if hexStr(reencoded) != hexStr(encoded) {
		t.Errorf("re-encoding not stable:\n got %v\nwant %v", hexStr(reencoded), hexStr(encoded))
	}
}

func TestEncodeSessionPermissionsDeadlineOutOfRange(t *testing.T) {
	sp := sampleSessionPermissions()
	// A "no expiry" deadline larger than uint64 must error, not panic.
	sp.Deadline = new(big.Int).Lsh(big.NewInt(1), 64) // 2^64

	if _, err := v3.EncodeSessionPermissions(&sp); err == nil {
		t.Fatal("expected error for deadline out of uint64 range, got nil")
	}
}

func TestAttestationFromJsonIssuedAtNumeric(t *testing.T) {
	const tmpl = `{
		"approvedSigner": "0x3333333333333333333333333333333333333333",
		"identityType": "0xaabbccdd",
		"issuerHash": "0x00000000000000000000000000000000000000000000000000000000000000a1",
		"audienceHash": "0x00000000000000000000000000000000000000000000000000000000000000b2",
		"applicationData": "0xdead",
		"authData": {"redirectUrl": "https://x.example", "issuedAt": %s}
	}`

	// A valid integral number is accepted.
	att, err := v3.AttestationFromJson(fmt.Sprintf(tmpl, "1678886400"))
	if err != nil {
		t.Fatalf("valid numeric issuedAt rejected: %v", err)
	}
	if att.AuthData.IssuedAt != 1678886400 {
		t.Errorf("issuedAt: got %v, want 1678886400", att.AuthData.IssuedAt)
	}

	// Negative, fractional and out-of-range numbers must be rejected, not
	// silently coerced into a valid-looking uint64.
	for _, bad := range []string{"-1", "1.5", "1e20"} {
		if _, err := v3.AttestationFromJson(fmt.Sprintf(tmpl, bad)); err == nil {
			t.Errorf("issuedAt %q: expected error, got nil", bad)
		}
	}
}

func sampleAttestation() v3.Attestation {
	return v3.Attestation{
		ApprovedSigner:  common.HexToAddress("0x3333333333333333333333333333333333333333"),
		IdentityType:    []byte{0xaa, 0xbb, 0xcc, 0xdd},
		IssuerHash:      common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000a1").Bytes(),
		AudienceHash:    common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000b2").Bytes(),
		ApplicationData: []byte{0xde, 0xad},
		AuthData:        v3.AuthData{RedirectUrl: "https://x.example", IssuedAt: 1678886400},
	}
}

func TestAttestationEncodingParity(t *testing.T) {
	const expectedEnc = "0x" +
		"3333333333333333333333333333333333333333" + // approvedSigner (20)
		"aabbccdd" + // identityType (4)
		"00000000000000000000000000000000000000000000000000000000000000a1" + // issuerHash (32)
		"00000000000000000000000000000000000000000000000000000000000000b2" + // audienceHash (32)
		"000002" + "dead" + // applicationData: uint24 length + data
		"000011" + "68747470733a2f2f782e6578616d706c65" + // authData redirectUrl: uint24 length + "https://x.example"
		"000000006411c600" // authData issuedAt (8, uint64)

	// keccak256(expectedEnc), verified equal to LibAttestation.toHash on-chain.
	const expectedHash = "0x94b5432a124f92ecfdd6da50f16d8f254f584c00a01682dafa4668bc72dbce06"

	att := sampleAttestation()
	if got := hexStr(att.Encode()); got != expectedEnc {
		t.Errorf("attestation encoding mismatch:\n got %v\nwant %v", got, expectedEnc)
	}
	if got := hexStr(att.Hash()); got != expectedHash {
		t.Errorf("attestation hash mismatch:\n got %v\nwant %v", got, expectedHash)
	}
}

func TestAttestationJSONRoundTrip(t *testing.T) {
	att := sampleAttestation()
	jsonStr, err := att.ToJson()
	if err != nil {
		t.Fatalf("ToJson: %v", err)
	}

	// issuedAt must serialize as a decimal string to match the sequence.js
	// primitives (attestation authData.issuedAt.toString()).
	if !strings.Contains(jsonStr, `"issuedAt":"1678886400"`) {
		t.Errorf("issuedAt not serialized as a string: %s", jsonStr)
	}

	parsed, err := v3.AttestationFromJson(jsonStr)
	if err != nil {
		t.Fatalf("AttestationFromJson: %v", err)
	}

	if parsed.AuthData.IssuedAt != att.AuthData.IssuedAt {
		t.Errorf("issuedAt: got %v, want %v", parsed.AuthData.IssuedAt, att.AuthData.IssuedAt)
	}
	if hexStr(parsed.Hash()) != hexStr(att.Hash()) {
		t.Errorf("hash after JSON round-trip: got %v, want %v", hexStr(parsed.Hash()), hexStr(att.Hash()))
	}
}
