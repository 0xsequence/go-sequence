package contract

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

func ParseArtifactJSON(artifactJSON string) (*ContractABI, error) {
	var artifact struct {
		ContractName string          `json:"contractName"`
		ABI          json.RawMessage `json:"abi"`
		Bytecode     string          `json:"bytecode"`
	}

	err := json.Unmarshal([]byte(artifactJSON), &artifact)
	if err != nil {
		return nil, err
	}

	contractABI := &ContractABI{}

	contractABI.Name = artifact.ContractName
	if contractABI.Name == "" {
		// TODO: if this comes up a lot, we could allow to set optional name on parse
		return nil, fmt.Errorf("contract name is empty")
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(artifact.ABI)))
	if err != nil {
		return nil, fmt.Errorf("unable to parse abi json in artifact: %w", err)
	}
	contractABI.ABI = &parsedABI

	if len(artifact.Bytecode) > 2 {
		contractABI.Bin = common.FromHex(artifact.Bytecode)
	}

	return contractABI, nil
}

func MustParseArtifactJSON(artifactJSON string) *ContractABI {
	contractABI, err := ParseArtifactJSON(artifactJSON)
	if err != nil {
		panic(err)
	}
	return contractABI
}

func ParseABI(abiJSON string) (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("unable to parse abi json: %w", err)
	}
	return &parsed, nil
}

func MustParseABI(abiJSON string) *abi.ABI {
	parsed, err := ParseABI(abiJSON)
	if err != nil {
		panic(err)
	}
	return parsed
}
