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

type SapientSignerElement struct {
	Address   common.Address
	Weight    uint8
	ImageHash core.ImageHash
}

func (s *SapientSignerElement) ToTree() v3.WalletConfigTree {
	return &v3.WalletConfigTreeSapientSignerLeaf{
		Weight:     s.Weight,
		Address:    s.Address,
		ImageHash_: s.ImageHash,
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

type JSONOutput struct {
	Threshold  string      `json:"threshold"`
	Checkpoint string      `json:"checkpoint"`
	Topology   interface{} `json:"topology"`
}

type JSONLeaf struct {
	Type      string `json:"type"`
	Address   string `json:"address,omitempty"`
	Weight    string `json:"weight,omitempty"`
	Hash      string `json:"hash,omitempty"`
	Digest    string `json:"digest,omitempty"`
	ImageHash string `json:"imageHash,omitempty"`
}

type JSONNode struct {
	Left  interface{} `json:"left"`
	Right interface{} `json:"right"`
}

type ConfigOutput struct {
	Threshold  string      `json:"threshold"`
	Checkpoint string      `json:"checkpoint"`
	Topology   interface{} `json:"topology"`
}

func (c ConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(JSONOutput(c))
}

type TopologyOutput struct {
	Type    string          `json:"type,omitempty"`
	Address string          `json:"address,omitempty"`
	Weight  string          `json:"weight,omitempty"`
	Left    *TopologyOutput `json:"left,omitempty"`
	Right   *TopologyOutput `json:"right,omitempty"`
}

func (t TopologyOutput) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	if t.Type != "" {
		m["type"] = t.Type
	}
	if t.Address != "" {
		m["address"] = t.Address
	}
	if t.Weight != "" {
		m["weight"] = t.Weight
	}
	if t.Left != nil {
		m["left"] = t.Left
	}
	if t.Right != nil {
		m["right"] = t.Right
	}

	return json.Marshal(m)
}

func treeToMap(tree v3.WalletConfigTree) interface{} {
	switch t := tree.(type) {
	case *v3.WalletConfigTreeAddressLeaf:
		return JSONLeaf{
			Type:    "signer",
			Address: t.Address.Hex(),
			Weight:  fmt.Sprintf("%d", t.Weight),
		}
	case *v3.WalletConfigTreeSapientSignerLeaf:
		return JSONLeaf{
			Type:      "sapient-signer",
			Address:   t.Address.Hex(),
			Weight:    fmt.Sprintf("%d", t.Weight),
			ImageHash: t.ImageHash_.Hex(),
		}
	case *v3.WalletConfigTreeSubdigestLeaf:
		return JSONLeaf{
			Type:   "subdigest",
			Digest: t.Subdigest.Hex(),
		}
	case *v3.WalletConfigTreeNestedLeaf:
		return treeToMap(t.Tree)
	case *v3.WalletConfigTreeNodeLeaf:
		return JSONLeaf{
			Type: "node",
			Hash: t.Node.Hex(),
		}
	case *v3.WalletConfigTreeNode:
		return []interface{}{
			treeToMap(t.Left),
			treeToMap(t.Right),
		}
	default:
		log.Printf("Unsupported tree type: %T", tree)
		return nil
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

	case "sapient":
		if len(parts) != 4 {
			return nil, fmt.Errorf("invalid sapient-signer format: %s", element)
		}
		weight := new(big.Int)
		if _, ok := weight.SetString(parts[3], 10); !ok {
			return nil, fmt.Errorf("invalid weight: %s", parts[3])
		}
		return &SapientSignerElement{
			Address:   common.HexToAddress(parts[2]),
			ImageHash: core.ImageHash{Hash: common.HexToHash(parts[1])},
			Weight:    uint8(weight.Uint64()),
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
		log.Printf("Parsed element: %v", configElement)
		if err != nil {
			return ConfigOutput{}, err
		}
		configElements = append(configElements, configElement)
	}

	if len(configElements) == 0 {
		return ConfigOutput{}, fmt.Errorf("no valid config elements provided")
	}

	var tree v3.WalletConfigTree
	if len(configElements) == 1 {
		tree = configElements[0].ToTree()
	} else {
		var trees []v3.WalletConfigTree
		for _, element := range configElements {
			trees = append(trees, element.ToTree())
		}
		// Wrap the balanced tree in a nested leaf with the threshold
		tree = &v3.WalletConfigTreeNestedLeaf{
			Threshold: threshold,
			Weight:    0, // Weight is typically 0 for the root
			Tree:      buildBalancedTree(trees),
		}
	}

	topology := treeToMap(tree)

	return ConfigOutput{
		Threshold:  fmt.Sprintf("%d", threshold),
		Checkpoint: fmt.Sprintf("%d", checkpoint),
		Topology:   topology,
	}, nil
}

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

	elements := strings.Fields(params.Content)

	log.Printf("Creating config with threshold=%s checkpoint=%s content=%q", params.Threshold, params.Checkpoint, params.Content)
	config, err := createConfig(uint16(threshold.Uint64()), uint64(checkpoint.Uint64()), elements)
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

func buildBalancedTree(elements []v3.WalletConfigTree) v3.WalletConfigTree {
	if len(elements) == 0 {
		return nil
	}
	if len(elements) == 1 {
		return elements[0]
	}
	mid := len(elements) / 2
	left := buildBalancedTree(elements[:mid])
	right := buildBalancedTree(elements[mid:])
	return &v3.WalletConfigTreeNode{
		Left:  left,
		Right: right,
	}
}

func buildTree(element interface{}) (interface{}, error) {
	switch v := element.(type) {
	case []interface{}:
		if len(v) != 2 {
			return nil, fmt.Errorf("array must have exactly two elements")
		}
		left, err := buildTree(v[0])
		if err != nil {
			return nil, err
		}
		right, err := buildTree(v[1])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"left":  left,
			"right": right,
		}, nil
	case map[string]interface{}:
		if tree, ok := v["tree"]; ok && v["type"] == "nested" {
			processedTree, err := buildTree(tree)
			if err != nil {
				return nil, fmt.Errorf("failed to process nested tree: %w", err)
			}
			v["tree"] = processedTree
		}
		return v, nil
	case string:
		return map[string]interface{}{
			"type": "node",
			"hash": v,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported element type: %T", v)
	}
}

func handleRawConfig(input []byte) (map[string]interface{}, error) {
	var rawConfig map[string]interface{}
	if err := json.Unmarshal(input, &rawConfig); err != nil {
		return nil, fmt.Errorf("failed to parse raw config: %w", err)
	}

	if input, ok := rawConfig["input"].(map[string]interface{}); ok {
		rawConfig = input
	}

	if topology := rawConfig["topology"]; topology != nil {
		tree, err := buildTree(topology)
		if err != nil {
			return nil, fmt.Errorf("failed to build tree from topology: %w", err)
		}
		rawConfig["tree"] = tree
		delete(rawConfig, "topology")
	}

	return rawConfig, nil
}

// calculateImageHash calculates the image hash for a given configuration
func calculateImageHash(params *ConfigImageHashParams) (string, error) {
	rawConfig, err := handleRawConfig(params.Input)
	if err != nil {
		return "", fmt.Errorf("failed to handle raw config: %w", err)
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
	rawConfig, err := handleRawConfig(params.Input)
	if err != nil {
		return "", fmt.Errorf("failed to handle raw config: %w", err)
	}

	config, err := v3.Core.DecodeWalletConfig(rawConfig)
	if err != nil {
		return "", fmt.Errorf("failed to decode wallet config: %w", err)
	}

	spew.Dump(config)

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
