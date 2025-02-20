package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
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

func createConfig(threshold uint16, checkpoint uint32, elements []string) (*v3.WalletConfig, error) {
	log.Printf("Creating config with threshold: %d, checkpoint: %d, elements: %v", threshold, checkpoint, elements)

	var configElements []ConfigElement
	for _, element := range elements {
		if element == "" {
			continue
		}
		configElement, err := parseElement(element)
		if err != nil {
			return nil, err
		}
		configElements = append(configElements, configElement)
	}

	var trees []v3.WalletConfigTree
	for _, element := range configElements {
		trees = append(trees, element.ToTree())
	}

	config := &v3.WalletConfig{
		Threshold_:  threshold,
		Checkpoint_: checkpoint,
		Tree:        v3.WalletConfigTreeNodes(trees...),
	}

	return config, nil
}

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration utilities",
	}

	var threshold uint16
	var checkpoint uint32

	newCmd := &cobra.Command{
		Use:   "new [elements...]",
		Short: "Create a new configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := createConfig(threshold, checkpoint, args)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal config: %w", err)
			}

			fmt.Println(string(output))
			return nil
		},
	}

	newCmd.Flags().Uint16VarP(&threshold, "threshold", "t", 1, "Threshold value for the configuration")
	newCmd.Flags().Uint32VarP(&checkpoint, "checkpoint", "c", 0, "Checkpoint value for the configuration")

	cmd.AddCommand(newCmd)

	return cmd
}
