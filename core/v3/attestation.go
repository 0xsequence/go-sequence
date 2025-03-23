package v3

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// Attestation implements the Attestation interface
type Attestation struct {
	ApprovedSigner  common.Address `json:"approvedSigner"`
	IdentityType    []byte         `json:"identityType"`    // bytes4
	IssuerHash      []byte         `json:"issuerHash"`      // bytes32
	AudienceHash    []byte         `json:"audienceHash"`    // bytes32
	ApplicationData []byte         `json:"applicationData"` // bytes
	AuthData        AuthData       `json:"authData"`
}

// AuthData represents authentication data with a redirect URL
type AuthData struct {
	RedirectUrl string `json:"redirectUrl"`
}

// Encode converts an Attestation to its binary representation
func (a *Attestation) Encode() []byte {
	// Pad identity type to 4 bytes if needed
	identityType := a.IdentityType
	if len(identityType) > 4 {
		identityType = identityType[:4]
	} else if len(identityType) < 4 {
		padded := make([]byte, 4)
		copy(padded[4-len(identityType):], identityType)
		identityType = padded
	}

	// Encode auth data
	authDataBytes := encodeAuthData(a.AuthData)

	// Concatenate all parts
	parts := [][]byte{
		a.ApprovedSigner.Bytes(),              // 20 bytes
		identityType,                          // 4 bytes
		PadLeft(a.IssuerHash, 32),             // 32 bytes
		PadLeft(a.AudienceHash, 32),           // 32 bytes
		intToBytes(len(a.ApplicationData), 3), // 3 bytes for length
		a.ApplicationData,                     // variable length
		authDataBytes,                         // variable length
	}

	return Concat(parts)
}

// encodeAuthData converts AuthData to its binary representation
func encodeAuthData(authData AuthData) []byte {
	redirectUrlBytes := []byte(authData.RedirectUrl)
	return Concat([][]byte{
		intToBytes(len(redirectUrlBytes), 3), // 3 bytes for length
		redirectUrlBytes,                     // variable length
	})
}

// Hash computes the Keccak256 hash of the encoded attestation
func (a *Attestation) Hash() []byte {
	return keccak256(a.Encode())
}

// EncodeForJson returns a JSON-friendly representation of the attestation
func (a *Attestation) EncodeForJson() interface{} {
	return map[string]interface{}{
		"approvedSigner":  a.ApprovedSigner.Hex(),
		"identityType":    "0x" + hex.EncodeToString(a.IdentityType),
		"issuerHash":      "0x" + hex.EncodeToString(a.IssuerHash),
		"audienceHash":    "0x" + hex.EncodeToString(a.AudienceHash),
		"applicationData": "0x" + hex.EncodeToString(a.ApplicationData),
		"authData":        a.AuthData,
	}
}

// ToJson converts the attestation to a JSON string
func (a *Attestation) ToJson() (string, error) {
	jsonData, err := json.Marshal(a.EncodeForJson())
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// FromJson parses a JSON string into an Attestation
func AttestationFromJson(jsonStr string) (*Attestation, error) {
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		return nil, err
	}
	return AttestationFromParsed(parsed)
}

// FromParsed reconstructs an Attestation from parsed JSON
func AttestationFromParsed(parsed map[string]interface{}) (*Attestation, error) {
	approvedSigner, ok := parsed["approvedSigner"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid approvedSigner")
	}

	identityType, ok := parsed["identityType"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid identityType")
	}
	identityTypeBytes, err := hex.DecodeString(identityType[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("invalid identityType hex: %w", err)
	}

	issuerHash, ok := parsed["issuerHash"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid issuerHash")
	}
	issuerHashBytes, err := hex.DecodeString(issuerHash[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("invalid issuerHash hex: %w", err)
	}

	audienceHash, ok := parsed["audienceHash"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid audienceHash")
	}
	audienceHashBytes, err := hex.DecodeString(audienceHash[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("invalid audienceHash hex: %w", err)
	}

	applicationData, ok := parsed["applicationData"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid applicationData")
	}
	applicationDataBytes, err := hex.DecodeString(applicationData[2:]) // Remove "0x" prefix
	if err != nil {
		return nil, fmt.Errorf("invalid applicationData hex: %w", err)
	}

	authData, ok := parsed["authData"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid authData")
	}

	redirectUrl, ok := authData["redirectUrl"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid redirectUrl")
	}

	return &Attestation{
		ApprovedSigner:  common.HexToAddress(approvedSigner),
		IdentityType:    identityTypeBytes,
		IssuerHash:      issuerHashBytes,
		AudienceHash:    audienceHashBytes,
		ApplicationData: applicationDataBytes,
		AuthData: AuthData{
			RedirectUrl: redirectUrl,
		},
	}, nil
}

// ACCEPT_IMPLICIT_REQUEST_MAGIC_PREFIX is the magic prefix for acceptImplicitRequest
var ACCEPT_IMPLICIT_REQUEST_MAGIC_PREFIX = keccak256([]byte("acceptImplicitRequest"))

// GenerateImplicitRequestMagic generates the magic value for an implicit request
func GenerateImplicitRequestMagic(attestation *Attestation, wallet common.Address) []byte {
	parts := [][]byte{
		ACCEPT_IMPLICIT_REQUEST_MAGIC_PREFIX,
		wallet.Bytes(),
		PadLeft(attestation.AudienceHash, 32),
		PadLeft(attestation.IssuerHash, 32),
	}
	return keccak256(Concat(parts))
}
