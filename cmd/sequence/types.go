package main

import (
	"encoding/json"

	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// Address params
type AddressCalculateParams struct {
	ImageHash    string `json:"imageHash"`
	Factory      string `json:"factory"`
	Module       string `json:"module"`
	CreationCode string `json:"creationCode"`
}

// Config params
type ConfigNewParams struct {
	Threshold    string `json:"threshold"`
	Checkpoint   string `json:"checkpoint"`
	From         string `json:"from"`
	Content      string `json:"content"`
	Checkpointer string `json:"checkpointer,omitempty"`
}

// ConfigImageHashParams represents parameters for calculating image hash
type ConfigImageHashParams struct {
	Input json.RawMessage `json:"input"`
}

// ConfigEncodeParams represents parameters for encoding configuration
type ConfigEncodeParams struct {
	Input json.RawMessage `json:"input"`
}

// Signature params
type SignatureEncodeParams struct {
	Input      json.RawMessage `json:"input"`
	Signatures string          `json:"signatures"`
	ChainId    bool            `json:"chainId"`
}

type SignatureConcatParams struct {
	Signatures []string `json:"signatures"`
}

type SignatureDecodeParams struct {
	Signature string `json:"signature"`
}

// DevTools params
type RandomConfigParams struct {
	MaxDepth             uint64 `json:"maxDepth"`
	Seed                 string `json:"seed"`
	MinThresholdOnNested uint64 `json:"minThresholdOnNested"`
	Checkpointer         string `json:"checkpointer"`
}

type RandomSessionTopologyParams struct {
	MaxDepth       uint64 `json:"maxDepth"`
	MaxPermissions uint64 `json:"maxPermissions"`
	MaxRules       uint64 `json:"maxRules"`
	Seed           string `json:"seed"`
}

// Payload params
type PayloadToAbiParams struct {
	Payload string `json:"payload"`
}

type PayloadToPackedParams struct {
	Payload string `json:"payload"`
}

type PayloadToJSONParams struct {
	Payload string `json:"payload"`
}

type PayloadHashParams struct {
	Wallet  string `json:"wallet"`
	ChainId string `json:"chainId"`
	Payload string `json:"payload"`
}

// Session params
type EmptyTopologyParams struct {
	IdentitySigner string `json:"identitySigner"`
}

type EncodeTopologyParams struct {
	SessionTopology v3.SessionsTopology `json:"sessionTopology"`
}

type EncodeCallSignatureParams struct {
	// Always
	SessionSignature string `json:"sessionSignature"`
	// Implicit
	Attestation       string `json:"attestation"`
	IdentitySignature string `json:"identitySignature"`
	// Explicit
	PermissionIndex string `json:"permissionIndex"`
}

type EncodeSessionCallSignaturesParams struct {
	SessionTopology v3.SessionsTopology    `json:"sessionTopology"`
	CallSignatures  []CallSignaturesParams `json:"callSignatures"`
	ExplicitSigners []string               `json:"explicitSigners"`
	ImplicitSigners []string               `json:"implicitSigners"`
}

type CallSignaturesParams struct {
	Attestation struct {
		ApprovedSigner  string `json:"approvedSigner"`
		IdentityType    string `json:"identityType"`
		IssuerHash      string `json:"issuerHash"`
		AudienceHash    string `json:"audienceHash"`
		ApplicationData string `json:"applicationData"`
		AuthData        struct {
			RedirectUrl string `json:"redirectUrl"`
		} `json:"authData"`
	} `json:"attestation"`
	IdentitySignature string `json:"identitySignature"` // RSV string
	SessionSignature  string `json:"sessionSignature"`  // RSV string
	PermissionIndex   string `json:"permissionIndex"`
}

type ImageHashParams struct {
	SessionTopology v3.SessionsTopology `json:"sessionTopology"`
}

type AddSessionParams struct {
	ExplicitSession json.RawMessage     `json:"explicitSession"`
	SessionTopology v3.SessionsTopology `json:"sessionTopology"`
}

type RemoveSessionParams struct {
	ExplicitSessionAddress string              `json:"explicitSessionAddress"`
	SessionTopology        v3.SessionsTopology `json:"sessionTopology"`
}

type AddBlacklistParams struct {
	BlacklistAddress string              `json:"blacklistAddress"`
	SessionTopology  v3.SessionsTopology `json:"sessionTopology"`
}

type RemoveBlacklistParams struct {
	BlacklistAddress string              `json:"blacklistAddress"`
	SessionTopology  v3.SessionsTopology `json:"sessionTopology"`
}
