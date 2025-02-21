package main

import "encoding/json"

// AddressCalculateParams represents parameters for address calculation
type AddressCalculateParams struct {
	ImageHash    string `json:"imageHash"`
	Factory      string `json:"factory"`
	Module       string `json:"module"`
	CreationCode string `json:"creationCode,omitempty"`
}

// ConfigNewParams represents parameters for creating a new configuration
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

// SignatureEncodeParams represents parameters for encoding signatures
type SignatureEncodeParams struct {
	Input      json.RawMessage `json:"input"`
	Signatures string          `json:"signatures"`
	ChainId    bool            `json:"chainId"`
}
