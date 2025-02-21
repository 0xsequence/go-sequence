package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

type ConfigElement interface {
	ToTree() v3.WalletConfigTree
}

type SignerElement struct {
	Address common.Address
	Weight  uint8
}

func (s *SignerElement) ToTree() v3.WalletConfigTree {
	return &v3.WalletConfigTreeAddressLeaf{
		Weight:  s.Weight,
		Address: s.Address,
	}
}

type SubdigestElement struct {
	Digest core.Subdigest
}

func (s *SubdigestElement) ToTree() v3.WalletConfigTree {
	return &v3.WalletConfigTreeSubdigestLeaf{
		Subdigest: s.Digest,
	}
}

type NestedElement struct {
	Threshold uint16
	Weight    uint8
	Elements  []ConfigElement
}

func (n *NestedElement) ToTree() v3.WalletConfigTree {
	var trees []v3.WalletConfigTree
	for _, element := range n.Elements {
		trees = append(trees, element.ToTree())
	}
	return &v3.WalletConfigTreeNestedLeaf{
		Weight:    n.Weight,
		Threshold: n.Threshold,
		Tree:      v3.WalletConfigTreeNodes(trees...),
	}
}

type NodeElement struct {
	Hash core.ImageHash
}

func (n *NodeElement) ToTree() v3.WalletConfigTree {
	return &v3.WalletConfigTreeNodeLeaf{
		Node: n.Hash,
	}
}

type ConfigOutput struct {
	Threshold  string         `json:"threshold"`
	Checkpoint string         `json:"checkpoint"`
	Topology   TopologyOutput `json:"topology"`
}

type TopologyOutput struct {
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
	Weight  string `json:"weight,omitempty"`
}

func treeToMap(tree v3.WalletConfigTree) TopologyOutput {
	switch t := tree.(type) {
	case *v3.WalletConfigTreeAddressLeaf:
		return TopologyOutput{
			Type:    "signer",
			Address: t.Address.Hex(),
			Weight:  fmt.Sprintf("%d", t.Weight),
		}
	case *v3.WalletConfigTreeSubdigestLeaf:
		return TopologyOutput{
			Type:    "subdigest",
			Address: t.Subdigest.Hex(),
		}
	case *v3.WalletConfigTreeNestedLeaf:
		return TopologyOutput{
			Type: "nested",
		}
	case *v3.WalletConfigTreeNodeLeaf:
		return TopologyOutput{
			Type:    "node",
			Address: t.Node.Hex(),
		}
	default:
		log.Printf("Unsupported tree type: %T", tree)
		return TopologyOutput{}
	}
}

func parseNestedElement(element string) (*NestedElement, error) {
	log.Printf("Parsing nested element: %q", element)

	// Extract the header part (before the opening parenthesis)
	headerEnd := strings.Index(element, "(")
	if headerEnd == -1 {
		return nil, fmt.Errorf("invalid nested format: missing opening parenthesis")
	}

	header := strings.TrimSpace(element[:headerEnd])
	header = strings.TrimSuffix(header, ":")
	log.Printf("Header: %q", header)

	parts := strings.Split(header, ":")
	log.Printf("Header parts: %v", parts)

	if len(parts) != 3 || parts[0] != "nested" {
		return nil, fmt.Errorf("invalid nested format: header should be 'nested:threshold:weight'")
	}

	// Parse threshold and weight
	threshold := new(big.Int)
	if _, ok := threshold.SetString(parts[1], 10); !ok {
		return nil, fmt.Errorf("invalid threshold: %s", parts[1])
	}

	weight := new(big.Int)
	if _, ok := weight.SetString(parts[2], 10); !ok {
		return nil, fmt.Errorf("invalid weight: %s", parts[2])
	}

	// Extract and parse the nested elements
	contentStart := headerEnd + 1
	contentEnd := strings.LastIndex(element, ")")
	if contentEnd == -1 || contentEnd <= contentStart {
		return nil, fmt.Errorf("invalid nested format: missing or mismatched closing parenthesis")
	}

	content := element[contentStart:contentEnd]
	log.Printf("Content: %q", content)

	var elements []ConfigElement

	// Split the content by spaces, respecting nested parentheses
	var currentElement strings.Builder
	parenthesesCount := 0

	for i := 0; i < len(content); i++ {
		char := content[i]
		switch char {
		case '(':
			parenthesesCount++
			currentElement.WriteByte(char)
		case ')':
			parenthesesCount--
			currentElement.WriteByte(char)
		case ' ':
			if parenthesesCount == 0 && currentElement.Len() > 0 {
				elementStr := currentElement.String()
				log.Printf("Parsing sub-element: %q", elementStr)
				element, err := parseElement(elementStr)
				if err != nil {
					return nil, fmt.Errorf("in nested element: %w", err)
				}
				elements = append(elements, element)
				currentElement.Reset()
			} else if parenthesesCount > 0 {
				currentElement.WriteByte(char)
			}
		default:
			currentElement.WriteByte(char)
		}
	}

	if currentElement.Len() > 0 {
		elementStr := currentElement.String()
		log.Printf("Parsing final sub-element: %q", elementStr)
		element, err := parseElement(elementStr)
		if err != nil {
			return nil, fmt.Errorf("in nested element: %w", err)
		}
		elements = append(elements, element)
	}

	if len(elements) == 0 {
		return nil, fmt.Errorf("invalid nested format: no elements found")
	}

	return &NestedElement{
		Threshold: uint16(threshold.Uint64()),
		Weight:    uint8(weight.Uint64()),
		Elements:  elements,
	}, nil
}

func parseElement(element string) (ConfigElement, error) {
	log.Printf("Parsing element: %q", element)

	if strings.HasPrefix(element, "nested:") {
		return parseNestedElement(element)
	}

	parts := strings.Split(element, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid element format: %s", element)
	}

	switch parts[0] {
	case "signer":
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid signer format: %s", element)
		}
		weight := new(big.Int)
		if _, ok := weight.SetString(parts[2], 10); !ok {
			return nil, fmt.Errorf("invalid weight: %s", parts[2])
		}
		return &SignerElement{
			Address: common.HexToAddress(parts[1]),
			Weight:  uint8(weight.Uint64()),
		}, nil

	case "subdigest":
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid subdigest format: %s", element)
		}
		hash := common.HexToHash(parts[1])
		return &SubdigestElement{
			Digest: core.Subdigest{Hash: hash},
		}, nil

	case "node":
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid node format: %s", element)
		}
		hash := common.HexToHash(parts[1])
		return &NodeElement{
			Hash: core.ImageHash{Hash: hash},
		}, nil

	default:
		return nil, fmt.Errorf("unknown element type: %s", parts[0])
	}
}

func createConfig(threshold uint16, checkpoint uint64, elements []string) (ConfigOutput, error) {
	log.Printf("Creating config with threshold: %d, checkpoint: %d, elements: %v", threshold, checkpoint, elements)

	var configElements []ConfigElement
	for _, element := range elements {
		if element == "" {
			continue
		}
		configElement, err := parseElement(element)
		if err != nil {
			return ConfigOutput{}, err
		}
		configElements = append(configElements, configElement)
	}

	var tree v3.WalletConfigTree
	if len(configElements) == 0 {
		return ConfigOutput{}, fmt.Errorf("no valid config elements provided")
	} else if len(configElements) == 1 {
		tree = configElements[0].ToTree()
	} else {
		var trees []v3.WalletConfigTree
		for _, element := range configElements {
			trees = append(trees, element.ToTree())
		}
		tree = &v3.WalletConfigTreeNestedLeaf{
			Threshold: threshold,
			Weight:    0,
			Tree:      v3.WalletConfigTreeNodes(trees...),
		}
	}

	return ConfigOutput{
		Threshold:  fmt.Sprintf("%d", threshold),
		Checkpoint: fmt.Sprintf("%d", checkpoint),
		Topology:   treeToMap(tree),
	}, nil
}

// createNewConfig creates a new configuration from the given parameters
func createNewConfig(params *ConfigNewParams) (string, error) {
	threshold, ok := new(big.Int).SetString(params.Threshold, 10)
	if !ok {
		return "", fmt.Errorf("invalid threshold: %s", params.Threshold)
	}
	if !threshold.IsUint64() || threshold.Uint64() > uint64(^uint16(0)) {
		return "", fmt.Errorf("threshold too large: %s", params.Threshold)
	}

	checkpoint, ok := new(big.Int).SetString(params.Checkpoint, 10)
	if !ok {
		return "", fmt.Errorf("invalid checkpoint: %s", params.Checkpoint)
	}
	if !checkpoint.IsUint64() {
		return "", fmt.Errorf("checkpoint too large: %s", params.Checkpoint)
	}

	if params.From != "flat" {
		return "", fmt.Errorf("unsupported 'from' format: %s", params.From)
	}

	log.Printf("Creating config with threshold=%s checkpoint=%s content=%q", params.Threshold, params.Checkpoint, params.Content)
	config, err := createConfig(uint16(threshold.Uint64()), uint64(checkpoint.Uint64()), []string{params.Content})
	if err != nil {
		return "", fmt.Errorf("failed to create config: %w", err)
	}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config: %w", err)
	}

	log.Printf("Created config: %s", string(jsonConfig))
	return string(jsonConfig), nil
}

// calculateImageHash calculates the image hash for a given configuration
func calculateImageHash(params *ConfigImageHashParams) (string, error) {
	var rawConfig map[string]interface{}
	if err := json.Unmarshal(params.Input, &rawConfig); err != nil {
		return "", fmt.Errorf("failed to parse raw config: %w", err)
	}

	// Check if the config is nested under an "input" field
	if input, ok := rawConfig["input"].(map[string]interface{}); ok {
		rawConfig = input
	}

	// Convert topology to tree if present
	if topology, ok := rawConfig["topology"].(map[string]interface{}); ok {
		// Convert topology to the expected tree format
		treeConfig := make(map[string]interface{})
		if topology["type"] == "signer" {
			// Convert string weight to int64
			if weight, ok := topology["weight"].(string); ok {
				weightInt := new(big.Int)
				if _, ok := weightInt.SetString(weight, 10); !ok {
					return "", fmt.Errorf("invalid weight: %s", weight)
				}
				if !weightInt.IsInt64() {
					return "", fmt.Errorf("weight too large: %s", weight)
				}
				treeConfig["weight"] = weightInt.Int64()
			} else {
				treeConfig["weight"] = topology["weight"]
			}
			treeConfig["address"] = topology["address"]
		} else {
			// Handle other types if needed
			return "", fmt.Errorf("unsupported topology type: %v", topology["type"])
		}
		rawConfig["tree"] = treeConfig
		delete(rawConfig, "topology")
	}

	config, err := v3.Core.DecodeWalletConfig(rawConfig)
	if err != nil {
		return "", fmt.Errorf("failed to decode wallet config: %w", err)
	}

	log.Printf("Decoded config: %+v", config)
	spew.Dump(config)

	imageHash := config.ImageHash()
	return "0x" + common.Bytes2Hex(imageHash.Bytes()), nil
}

// encodeConfig encodes a configuration into a signature format
func encodeConfig(params *ConfigEncodeParams) (string, error) {
	var rawConfig map[string]interface{}
	if err := json.Unmarshal(params.Input, &rawConfig); err != nil {
		return "", fmt.Errorf("failed to parse raw config: %w", err)
	}

	config, err := v3.Core.DecodeWalletConfig(rawConfig)
	if err != nil {
		return "", fmt.Errorf("failed to decode wallet config: %w", err)
	}

	sig, err := config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, sigs []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		return 0, nil, core.ErrSigningNoSigner
	}, false)
	if err != nil {
		return "", fmt.Errorf("failed to build signature: %w", err)
	}

	data, err := sig.Data()
	if err != nil {
		return "", fmt.Errorf("failed to encode signature: %w", err)
	}

	return "0x" + common.Bytes2Hex(data), nil
}

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration utilities",
	}

	var newParams ConfigNewParams
	newCmd := &cobra.Command{
		Use:   "new [elements...]",
		Short: "Create a new configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			newParams.Content = strings.Join(args, " ")
			result, err := createNewConfig(&newParams)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	newCmd.Flags().StringVar(&newParams.Threshold, "threshold", "1", "Threshold value for the configuration")
	newCmd.Flags().StringVar(&newParams.Checkpoint, "checkpoint", "0", "Checkpoint value for the configuration")
	newCmd.Flags().StringVar(&newParams.From, "from", "flat", "The process to use to create the configuration")

	imageHashCmd := &cobra.Command{
		Use:   "image-hash [input]",
		Short: "Calculate image hash from hex input",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("input is required")
			}
			params := &ConfigImageHashParams{
				Input: json.RawMessage(args[0]),
			}
			result, err := calculateImageHash(params)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeCmd := &cobra.Command{
		Use:   "encode [input]",
		Short: "Encode configuration from hex input",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("input is required")
			}
			params := &ConfigEncodeParams{
				Input: json.RawMessage(args[0]),
			}
			result, err := encodeConfig(params)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(newCmd, imageHashCmd, encodeCmd)

	log.SetOutput(os.Stderr)

	return cmd
}
