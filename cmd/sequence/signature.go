package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

type SignatureElement struct {
	Address string   `json:"address"`
	Type    string   `json:"type"`
	Values  []string `json:"values"`
}

func (s *SignatureElement) ToTree() v3.WalletConfigTree {
	return &v3.WalletConfigTreeAddressLeaf{
		Address: common.HexToAddress(s.Address),
	}
}

type SignatureConcatParams struct {
	Signatures []string `json:"signatures"`
}

type WalletConfigInput struct {
	Threshold  json.RawMessage `json:"threshold"`
	Checkpoint json.RawMessage `json:"checkpoint"`
	Tree       json.RawMessage `json:"tree,omitempty"`
	Topology   json.RawMessage `json:"topology,omitempty"`
}

func convertTopologyToTree(topology interface{}) (interface{}, error) {
	switch v := topology.(type) {
	case []interface{}:
		if len(v) != 2 {
			return nil, fmt.Errorf("array must have exactly two elements")
		}
		left, err := convertTopologyToTree(v[0])
		if err != nil {
			return nil, err
		}
		right, err := convertTopologyToTree(v[1])
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"left":  left,
			"right": right,
		}, nil
	case map[string]interface{}:
		if nodeType, ok := v["type"].(string); ok {
			switch nodeType {
			case "signer":
				return map[string]interface{}{
					"weight":  v["weight"],
					"address": v["address"],
				}, nil
			case "nested":
				if tree, ok := v["tree"]; ok {
					processedTree, err := convertTopologyToTree(tree)
					if err != nil {
						return nil, err
					}
					return map[string]interface{}{
						"threshold": v["threshold"],
						"weight":    v["weight"],
						"tree":      processedTree,
					}, nil
				}
			case "node":
				if hash, ok := v["hash"].(string); ok {
					return map[string]interface{}{
						"hash": hash,
					}, nil
				}
			}
		}

		if left, hasLeft := v["left"]; hasLeft {
			if right, hasRight := v["right"]; hasRight {
				leftProcessed, err := convertTopologyToTree(left)
				if err != nil {
					return nil, err
				}
				rightProcessed, err := convertTopologyToTree(right)
				if err != nil {
					return nil, err
				}
				return map[string]interface{}{
					"left":  leftProcessed,
					"right": rightProcessed,
				}, nil
			}
		}
		return v, nil
	default:
		return nil, fmt.Errorf("unsupported topology type: %T", v)
	}
}

func signatureConcat(params *SignatureConcatParams) (string, error) {
	var combined []byte
	for _, sig := range params.Signatures {
		sig = strings.TrimPrefix(sig, "0x")

		sigBytes := common.FromHex(sig)
		combined = append(combined, sigBytes...)
	}

	return "0x" + common.Bytes2Hex(combined), nil
}

func signatureEncode(p *SignatureEncodeParams) (interface{}, error) {
	var input WalletConfigInput
	if err := json.Unmarshal(p.Input, &input); err != nil {
		return nil, fmt.Errorf("failed to parse input config: %w", err)
	}

	var threshold, checkpoint interface{}
	if err := json.Unmarshal(input.Threshold, &threshold); err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}
	if err := json.Unmarshal(input.Checkpoint, &checkpoint); err != nil {
		return nil, fmt.Errorf("failed to parse checkpoint: %w", err)
	}

	rawConfig := map[string]interface{}{
		"threshold":  convertNumbers(threshold),
		"checkpoint": convertNumbers(checkpoint),
	}

	var treeData json.RawMessage
	if input.Tree != nil {
		treeData = input.Tree
	} else if input.Topology != nil {
		treeData = input.Topology
	} else {
		return nil, fmt.Errorf("either tree or topology must be provided")
	}

	var tree interface{}
	if err := json.Unmarshal(treeData, &tree); err != nil {
		return nil, fmt.Errorf("failed to parse tree/topology: %w", err)
	}

	var err error
	tree, err = convertTopologyToTree(tree)
	if err != nil {
		return nil, fmt.Errorf("failed to convert topology to tree: %w", err)
	}

	rawConfig["tree"] = convertNumbers(tree)

	config, err := v3.Core.DecodeWalletConfig(rawConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to decode wallet config: %w", err)
	}

	var signatures []SignatureElement
	if p.Signatures != "" {
		for _, sig := range strings.Fields(p.Signatures) {
			parts := strings.Split(sig, ":")
			if len(parts) < 2 {
				continue
			}

			if len(parts) == 5 && parts[1] == "hash" {
				signatures = append(signatures, SignatureElement{
					Address: parts[0],
					Type:    "hash",
					Values:  []string{parts[2], parts[3], parts[4]},
				})
				continue
			}

			switch parts[0] {
			case "signer":
				if len(parts) != 3 {
					continue
				}
				signatures = append(signatures, SignatureElement{
					Address: parts[1],
					Type:    "hash",
					Values:  []string{},
				})
			case "sapient":
				if len(parts) != 4 {
					continue
				}
				signatures = append(signatures, SignatureElement{
					Address: parts[2],
					Type:    "sapient",
					Values:  []string{parts[1], parts[2], parts[3]},
				})
			case "subdigest":
				if len(parts) != 2 {
					continue
				}
				signatures = append(signatures, SignatureElement{
					Type:   "subdigest",
					Values: []string{parts[1]},
				})
			case "node":
				if len(parts) != 2 {
					continue
				}
				signatures = append(signatures, SignatureElement{
					Type:   "node",
					Values: []string{parts[1]},
				})
			case "any-address-subdigest":
				if len(parts) != 2 {
					continue
				}
				signatures = append(signatures, SignatureElement{
					Type:   "any-address-subdigest",
					Values: []string{parts[1]},
				})
			}
		}
	}

	var signature core.Signature[*v3.WalletConfig]
	if p.ChainId {
		signature, err = config.BuildRegularSignature(context.Background(), func(ctx context.Context, signer common.Address, sigs []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
			for _, sig := range signatures {
				if strings.EqualFold(sig.Address, signer.Hex()) {
					switch sig.Type {
					case "eth_sign":
						if len(sig.Values) != 3 {
							continue
						}
						r := common.HexToHash(sig.Values[0]).Bytes()
						s := common.HexToHash(sig.Values[1]).Bytes()
						v := byte(common.HexToHash(sig.Values[2]).Big().Uint64())

						return core.SignerSignatureTypeEthSign, append(append(r, s...), v), nil
					case "hash":
						if len(sig.Values) != 3 {
							continue
						}
						r := common.HexToHash(sig.Values[0]).Bytes()
						s := common.HexToHash(sig.Values[1]).Bytes()
						v := byte(common.HexToHash(sig.Values[2]).Big().Uint64())

						return core.SignerSignatureTypeEIP712, append(append(r, s...), v), nil
					case "erc1271":
						if len(sig.Values) != 1 {
							continue
						}

						return core.SignerSignatureTypeEIP1271, common.FromHex(sig.Values[0]), nil
					case "sapient", "sapient_compact":
						if len(sig.Values) != 1 {
							continue
						}

						if sig.Type == "sapient" {
							return core.SignerSignatureTypeSapient, common.FromHex(sig.Values[0]), nil
						} else {
							return core.SignerSignatureTypeSapientCompact, common.FromHex(sig.Values[0]), nil
						}
					default:
						log.Printf("Unsupported signature type: %s", sig.Type)
						continue
					}
				}
			}

			return 0, nil, core.ErrSigningNoSigner
		})
	} else {
		signature, err = config.BuildNoChainIDSignature(context.Background(), func(ctx context.Context, signer common.Address, sigs []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
			for _, sig := range signatures {
				if strings.EqualFold(sig.Address, signer.Hex()) {
					switch sig.Type {
					case "eth_sign":
						if len(sig.Values) != 3 {
							continue
						}
						r := common.HexToHash(sig.Values[0]).Bytes()
						s := common.HexToHash(sig.Values[1]).Bytes()
						v := byte(common.HexToHash(sig.Values[2]).Big().Uint64())

						return core.SignerSignatureTypeEthSign, append(append(r, s...), v), nil
					case "hash":
						if len(sig.Values) != 3 {
							continue
						}
						r := common.HexToHash(sig.Values[0]).Bytes()
						s := common.HexToHash(sig.Values[1]).Bytes()
						v := byte(common.HexToHash(sig.Values[2]).Big().Uint64())

						return core.SignerSignatureTypeEIP712, append(append(r, s...), v), nil
					case "erc1271":
						if len(sig.Values) != 1 {
							continue
						}

						return core.SignerSignatureTypeEIP1271, common.FromHex(sig.Values[0]), nil
					case "sapient", "sapient_compact":
						if len(sig.Values) != 1 {
							continue
						}

						if sig.Type == "sapient" {
							return core.SignerSignatureTypeSapient, common.FromHex(sig.Values[0]), nil
						} else {
							return core.SignerSignatureTypeSapientCompact, common.FromHex(sig.Values[0]), nil
						}
					default:
						log.Printf("Unsupported signature type: %s", sig.Type)
						continue
					}
				}
			}

			return 0, nil, core.ErrSigningNoSigner
		})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to build signature: %w", err)
	}

	data, err := signature.Data()
	if err != nil {
		return nil, fmt.Errorf("failed to encode signature: %w", err)
	}

	return "0x" + common.Bytes2Hex(data), nil
}

func newSignatureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signature",
		Short: "Signature utilities",
	}

	var input string
	var signatures []string
	var chainId bool

	encodeCmd := &cobra.Command{
		Use:   "encode [input]",
		Short: "Encode signature from hex input",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				input = args[0]
			}
			if input == "" {
				return fmt.Errorf("input is required")
			}
			if len(signatures) == 0 {
				return fmt.Errorf("at least one signature is required")
			}

			allSignatures := strings.Join(signatures, " ")

			params := &SignatureEncodeParams{
				Input:      json.RawMessage(input),
				Signatures: allSignatures,
				ChainId:    chainId,
			}

			result, err := signatureEncode(params)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	concatCmd := &cobra.Command{
		Use:   "concat [signatures...]",
		Short: "Concatenate multiple signatures",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("at least one signature is required")
			}

			params := &SignatureConcatParams{
				Signatures: args,
			}

			result, err := signatureConcat(params)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	encodeCmd.Flags().StringArrayVarP(&signatures, "signature", "s", []string{}, "Signatures")
	encodeCmd.Flags().BoolVar(&chainId, "chain-id", true, "Use chainId of recovered chain on signature")

	cmd.AddCommand(encodeCmd, concatCmd)

	return cmd
}
