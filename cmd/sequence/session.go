package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"encoding/hex"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

func handleEmptyTopology(p *EmptyTopologyParams) (string, error) {
	signer := common.HexToAddress(p.IdentitySigner)
	topology := v3.EmptySessionsTopology(signer)
	return v3.SessionsTopologyToJSON(topology)
}

func handleEncodeTopology(p *EncodeTopologyParams) (string, error) {
	encoded, err := v3.EncodeSessionsTopology(p.SessionTopology)
	if err != nil {
		return "", fmt.Errorf("failed to encode session topology: %w", err)
	}
	return hexutil.Encode(encoded), nil
}

func handleEncodeSessionCallSignatures(p *EncodeSessionCallSignaturesParams) (string, error) {
	callSigs := make([]v3.SessionCallSignature, len(p.CallSignatures))
	for i, sig := range p.CallSignatures {
		sessionSig, err := rsyFromRsvStr(sig.SessionSignature)
		if err != nil {
			return "", fmt.Errorf("failed to decode session signature: %w", err)
		}
		if sig.Attestation.ApprovedSigner != "" && sig.IdentitySignature != "" {
			// Implicit call signature
			idSig, err := rsyFromRsvStr(sig.IdentitySignature)
			if err != nil {
				return "", fmt.Errorf("failed to decode identity signature: %w", err)
			}

			// Convert hex-encoded fields to bytes
			identityType, err := hex.DecodeString(sig.Attestation.IdentityType[2:]) // Remove "0x" prefix
			if err != nil {
				return "", fmt.Errorf("failed to decode identity type: %w", err)
			}

			issuerHash, err := hex.DecodeString(sig.Attestation.IssuerHash[2:])
			if err != nil {
				return "", fmt.Errorf("failed to decode issuer hash: %w", err)
			}

			audienceHash, err := hex.DecodeString(sig.Attestation.AudienceHash[2:])
			if err != nil {
				return "", fmt.Errorf("failed to decode audience hash: %w", err)
			}

			applicationData, err := hex.DecodeString(sig.Attestation.ApplicationData[2:])
			if err != nil {
				return "", fmt.Errorf("failed to decode application data: %w", err)
			}

			attestation := v3.Attestation{
				ApprovedSigner:  common.HexToAddress(sig.Attestation.ApprovedSigner),
				IdentityType:    identityType,
				IssuerHash:      issuerHash,
				AudienceHash:    audienceHash,
				ApplicationData: applicationData,
				AuthData: v3.AuthData{
					RedirectUrl: sig.Attestation.AuthData.RedirectUrl,
				},
			}

			callSig := v3.ImplicitSessionCallSignature{
				Attestation:       attestation,
				IdentitySignature: idSig,
				SessionSignature:  sessionSig,
			}
			callSigs[i] = callSig
		} else {
			// Explicit call signature
			permissionIndex, ok := new(big.Int).SetString(sig.PermissionIndex, 10)
			if !ok {
				return "", fmt.Errorf("failed to decode permission index: %w", err)
			}
			callSig := v3.ExplicitSessionCallSignature{
				PermissionIndex:  permissionIndex,
				SessionSignature: sessionSig,
			}
			callSigs[i] = callSig
		}
	}

	explicitAddrs := make([]common.Address, len(p.ExplicitSigners))
	for i, signer := range p.ExplicitSigners {
		explicitAddrs[i] = common.HexToAddress(signer)
	}

	implicitAddrs := make([]common.Address, len(p.ImplicitSigners))
	for i, signer := range p.ImplicitSigners {
		implicitAddrs[i] = common.HexToAddress(signer)
	}

	encoded, err := v3.EncodeSessionCallSignatures(callSigs, p.SessionTopology, explicitAddrs, implicitAddrs)
	if err != nil {
		return "", fmt.Errorf("failed to encode session call signatures: %w", err)
	}
	return hexutil.Encode(encoded), nil
}

func handleImageHash(p *ImageHashParams) (string, error) {
	configTree, err := v3.SessionsTopologyToGenericTree(p.SessionTopology)
	if err != nil {
		return "", fmt.Errorf("failed to convert session topology to generic tree: %w", err)
	}
	return configTree.ImageHash().String(), nil
}

func rsyFromRsvStr(sigStr string) (v3.RSY, error) {
	parts := strings.Split(sigStr, ":")
	if len(parts) != 3 {
		return v3.RSY{}, errors.New("signature must be in r:s:v format")
	}
	r, err := hex.DecodeString(parts[0][2:])
	if err != nil {
		return v3.RSY{}, fmt.Errorf("invalid r value: %w", err)
	}
	s, err := hex.DecodeString(parts[1][2:])
	if err != nil {
		return v3.RSY{}, fmt.Errorf("invalid s value: %w", err)
	}
	v, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		return v3.RSY{}, fmt.Errorf("invalid v value: %w", err)
	}
	return v3.RSY{R: r, S: s, YParity: uint8(v - 27)}, nil
}

// newSessionCmd defines the command-line interface.
func newSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session utilities",
	}

	emptyCmd := &cobra.Command{
		Use:   "empty [identity-signer]",
		Short: "Create an empty session topology with the given identity signer",
		RunE: func(cmd *cobra.Command, args []string) error {
			identitySigner, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleEmptyTopology(&EmptyTopologyParams{IdentitySigner: identitySigner})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeTopologyCmd := &cobra.Command{
		Use:   "encode-topology [session-topology]",
		Short: "Encode a session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			topology, err := v3.SessionsTopologyFromJSON(config)
			if err != nil {
				return err
			}
			result, err := handleEncodeTopology(&EncodeTopologyParams{SessionTopology: topology})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeCallsCmd := &cobra.Command{
		Use:   "encode-calls [session-topology] [call-signatures] [explicit-signers] [implicit-signers]",
		Short: "Encode call signatures for sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 4 {
				return fmt.Errorf("requires session topology, call signatures, explicit signers, and implicit signers")
			}

			config := args[0]
			callSigs := make([]CallSignaturesParams, len(args)-1)
			for i, sig := range args[1:] {
				callSigs[i] = CallSignaturesParams{
					SessionSignature: sig,
				}
			}
			explicitSigners, _ := cmd.Flags().GetStringSlice("explicit-signers")
			implicitSigners, _ := cmd.Flags().GetStringSlice("implicit-signers")

			topology, err := v3.SessionsTopologyFromJSON(config)
			if err != nil {
				return err
			}

			result, err := handleEncodeSessionCallSignatures(&EncodeSessionCallSignaturesParams{
				SessionTopology: topology,
				CallSignatures:  callSigs,
				ExplicitSigners: explicitSigners,
				ImplicitSigners: implicitSigners,
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
		Use:   "image-hash [session-topology]",
		Short: "Hash a session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			topology, err := v3.SessionsTopologyFromJSON(config)
			if err != nil {
				return err
			}
			result, err := handleImageHash(&ImageHashParams{SessionTopology: topology})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(emptyCmd, encodeTopologyCmd, encodeCallsCmd, imageHashCmd)
	cmd.AddCommand(newSessionExplicitCmd(), newSessionImplicitCmd())
	return cmd
}
