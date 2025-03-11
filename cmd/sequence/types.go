package main

import "encoding/json"

// Address params
type AddressCalculateParams struct {
	ImageHash    string `json:"imageHash"`
	Factory      string `json:"factory"`
	Module       string `json:"module"`
	CreationCode string `json:"creationCode"`
}

// Config params
type ConfigNewParams struct {
	Threshold  string `json:"threshold"`
	Checkpoint string `json:"checkpoint"`
	From       string `json:"from"`
	Content    string `json:"content"`
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

type PayloadToJsonParams struct {
	Payload string `json:"payload"`
}

type PayloadHashParams struct {
	Wallet  string `json:"wallet"`
	ChainId string `json:"chainId"`
	Payload string `json:"payload"`
}

// Permission params
type PermissionToPackedSessionParams struct {
	SessionPermission string `json:"sessionPermission"`
}

type PermissionToPackedParams struct {
	Permission string `json:"permission"`
}

// Session params
type EmptyTopologyParams struct {
	GlobalSigner string `json:"globalSigner"`
}

type EncodeConfigurationParams struct {
	SessionConfiguration string `json:"sessionConfiguration"`
}

type EncodeCallSignatureParams struct {
	Signature       string `json:"signature"`
	PermissionIndex int    `json:"permissionIndex"`
}

type EncodeSessionCallSignaturesParams struct {
	SessionConfiguration string          `json:"-"`
	SessionTopology      json.RawMessage `json:"-"`

	CallSignatures []string `json:"-"`

	ExplicitSigners []string `json:"-"`
	ImplicitSigners []string `json:"-"`
}

// UnmarshalJSON implements a custom JSON unmarshaler for EncodeSessionCallSignaturesParams
func (p *EncodeSessionCallSignaturesParams) UnmarshalJSON(data []byte) error {
	type rawParams struct {
		SessionConfiguration json.RawMessage `json:"sessionConfiguration"`
		SessionTopology      json.RawMessage `json:"sessionTopology"`
		CallSignatures       json.RawMessage `json:"callSignatures"`
		ExplicitSigners      []string        `json:"explicitSigners"`
		ImplicitSigners      []string        `json:"implicitSigners"`
	}

	var raw rawParams
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw.SessionConfiguration) > 0 {
		p.SessionConfiguration = string(raw.SessionConfiguration)
	}
	if len(raw.SessionTopology) > 0 {
		p.SessionTopology = raw.SessionTopology
	}

	p.ExplicitSigners = raw.ExplicitSigners
	p.ImplicitSigners = raw.ImplicitSigners

	if len(raw.CallSignatures) > 0 {
		var sigArray []json.RawMessage
		if err := json.Unmarshal(raw.CallSignatures, &sigArray); err == nil {
			callSigs := make([]string, len(sigArray))
			for i, sig := range sigArray {
				callSigs[i] = string(sig)
			}
			p.CallSignatures = callSigs
		} else {
			var singleSig string
			if jsonErr := json.Unmarshal(raw.CallSignatures, &singleSig); jsonErr == nil {
				p.CallSignatures = []string{singleSig}
			} else {
				p.CallSignatures = []string{string(raw.CallSignatures)}
			}
		}
	} else {
		p.CallSignatures = []string{}
	}

	return nil
}

type ImageHashParams struct {
	SessionConfiguration string `json:"sessionConfiguration"`
}

type AddSessionParams struct {
	ExplicitSession json.RawMessage `json:"explicitSession"`
	SessionTopology json.RawMessage `json:"sessionTopology"`
	IdentitySigner  string          `json:"-"` // Used internally, not from JSON
}

type RemoveSessionParams struct {
	ExplicitSessionAddress string `json:"explicitSessionAddress"`
	SessionTopology        string `json:"sessionTopology"`
}

type AddBlacklistParams struct {
	BlacklistAddress     string `json:"blacklistAddress"`
	SessionConfiguration string `json:"sessionConfiguration"`
}

type RemoveBlacklistParams struct {
	BlacklistAddress     string `json:"blacklistAddress"`
	SessionConfiguration string `json:"sessionConfiguration"`
}
