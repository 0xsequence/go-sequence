package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func handleEmptyTopology(p *EmptyTopologyParams) (string, error) {
	signer := common.HexToAddress(p.GlobalSigner)
	topology := v3.EmptySessionsTopology(signer)
	return v3.SessionsTopologyToJSON(topology)
}

func handleEncodeConfiguration(p *EncodeConfigurationParams) (string, error) {
	topology, err := v3.SessionsTopologyFromJSON(p.SessionConfiguration)
	if err != nil {
		return "", fmt.Errorf("failed to decode session configuration: %w", err)
	}

	configTree := v3.SessionsTopologyToConfigurationTree(topology)
	jsonBytes, err := json.Marshal(configTree)
	if err != nil {
		return "", fmt.Errorf("failed to marshal configuration tree: %w", err)
	}

	return string(jsonBytes), nil
}

func handleEncodeSessionCallSignatures(p *EncodeSessionCallSignaturesParams) (string, error) {
	var topologyStr string
	if p.SessionConfiguration != "" {
		topologyStr = p.SessionConfiguration
	} else if len(p.SessionTopology) > 0 {
		topologyStr = string(p.SessionTopology)
	} else {
		return "", fmt.Errorf("missing session topology configuration")
	}

	topology, err := v3.SessionsTopologyFromJSON(topologyStr)
	if err != nil {
		return "", fmt.Errorf("failed to decode session configuration: %w", err)
	}

	log.Println("topology")
	spew.Dump(topology)

	callSigs := make([]v3.SessionCallSignature, len(p.CallSignatures))
	for i, sigStr := range p.CallSignatures {
		callSig, err := v3.SessionCallSignatureFromJSON(sigStr)
		if err != nil {
			return "", fmt.Errorf("failed to decode call signature %d: %w", i, err)
		}
		callSigs[i] = *callSig
	}

	log.Println("callSigs")
	spew.Dump(callSigs)

	explicitAddrs := make([]common.Address, len(p.ExplicitSigners))
	for i, signer := range p.ExplicitSigners {
		explicitAddrs[i] = common.HexToAddress(signer)
	}

	log.Println("explicitAddrs")
	spew.Dump(explicitAddrs)

	implicitAddrs := make([]common.Address, len(p.ImplicitSigners))
	for i, signer := range p.ImplicitSigners {
		implicitAddrs[i] = common.HexToAddress(signer)
	}

	log.Println("implicitAddrs")
	spew.Dump(implicitAddrs)

	encoded, err := v3.EncodeSessionCallSignatures(topology, callSigs, explicitAddrs, implicitAddrs)
	if err != nil {
		return "", fmt.Errorf("failed to encode session call signatures: %w", err)
	}

	return string(encoded), nil
}

func handleImageHash(p *ImageHashParams) (string, error) {
	topology, err := v3.SessionsTopologyFromJSON(p.SessionConfiguration)
	if err != nil {
		return "", fmt.Errorf("failed to decode session configuration: %w", err)
	}

	configTree := v3.SessionsTopologyToConfigurationTree(topology)
	hash := v3.HashConfigurationTree(configTree)
	return common.Bytes2Hex(hash[:]), nil
}

// newSessionCmd defines the command-line interface.
func newSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session utilities",
	}

	emptyCmd := &cobra.Command{
		Use:   "empty [global-signer]",
		Short: "Create an empty session topology with the given global signer",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalSigner, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleEmptyTopology(&EmptyTopologyParams{GlobalSigner: globalSigner})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeConfigurationCmd := &cobra.Command{
		Use:   "encode-configuration [session-configuration]",
		Short: "Encode a session configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleEncodeConfiguration(&EncodeConfigurationParams{SessionConfiguration: config})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeCallsCmd := &cobra.Command{
		Use:   "encode-calls [session-configuration] [call-signatures...]",
		Short: "Encode a call signature for an implicit session",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("requires session configuration and at least one call signature")
			}

			config := args[0]
			callSigs := args[1:]
			explicitSigners, _ := cmd.Flags().GetStringSlice("explicit-signers")
			implicitSigners, _ := cmd.Flags().GetStringSlice("implicit-signers")

			result, err := handleEncodeSessionCallSignatures(&EncodeSessionCallSignaturesParams{
				SessionConfiguration: config,
				CallSignatures:       callSigs,
				ExplicitSigners:      explicitSigners,
				ImplicitSigners:      implicitSigners,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	encodeCallsCmd.Flags().StringSlice("explicit-signers", []string{}, "Explicit signers")
	encodeCallsCmd.Flags().StringSlice("implicit-signers", []string{}, "Implicit signers")

	imageHashCmd := &cobra.Command{
		Use:   "image-hash [session-configuration]",
		Short: "Hash a session configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleImageHash(&ImageHashParams{SessionConfiguration: config})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(emptyCmd, encodeConfigurationCmd, encodeCallsCmd, imageHashCmd)
	cmd.AddCommand(newSessionExplicitCmd(), newSessionImplicitCmd())
	return cmd
}
