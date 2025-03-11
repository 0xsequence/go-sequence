package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

var sessionMgr = NewSessionManager()

type SessionManager struct {
	identitySigners map[string]string
	mu              sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		identitySigners: make(map[string]string),
	}
}

func (sm *SessionManager) StoreIdentitySigner(topologyData json.RawMessage, identitySigner string) {
	if identitySigner == "" {
		return
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	key := string(topologyData)
	sm.identitySigners[key] = identitySigner
	log.Printf("Stored identity signer %s for topology hash", identitySigner)
}

func (sm *SessionManager) GetIdentitySigner(topologyData json.RawMessage) string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	key := string(topologyData)
	return sm.identitySigners[key]
}

func handleAddExplicitSession(p *AddSessionParams) (string, error) {
	session, err := v3.SessionPermissionsFromJSON(string(p.ExplicitSession))
	if err != nil {
		return "", fmt.Errorf("failed to decode explicit session: %w", err)
	}

	topology, err := v3.SessionsTopologyFromJSON(string(p.SessionTopology))
	if err != nil {
		return "", fmt.Errorf("failed to decode session topology: %w", err)
	}

	if !v3.IsSessionsTopology(topology) {
		return "", fmt.Errorf("session topology must be a valid session topology")
	}

	var identitySigner string

	if p.IdentitySigner != "" {
		identitySigner = p.IdentitySigner
	}

	if identitySigner == "" {
		identitySigner = sessionMgr.GetIdentitySigner(p.SessionTopology)
	}

	if identitySigner == "" {
		var topArray []json.RawMessage
		if err := json.Unmarshal(p.SessionTopology, &topArray); err == nil && len(topArray) >= 2 {
			var secondElem json.RawMessage = topArray[1]
			if len(secondElem) > 0 {
				if secondElem[0] == '[' {
					var secondArray []json.RawMessage
					if err := json.Unmarshal(secondElem, &secondArray); err == nil && len(secondArray) > 0 {
						var signerObj struct {
							IdentitySigner string `json:"identitySigner"`
						}
						if err := json.Unmarshal(secondArray[0], &signerObj); err == nil && signerObj.IdentitySigner != "" {
							identitySigner = signerObj.IdentitySigner
							sessionMgr.StoreIdentitySigner(p.SessionTopology, identitySigner)
						}
					}
				} else if secondElem[0] == '{' {
					var signerObj struct {
						IdentitySigner string `json:"identitySigner"`
					}
					if err := json.Unmarshal(secondElem, &signerObj); err == nil && signerObj.IdentitySigner != "" {
						identitySigner = signerObj.IdentitySigner
						sessionMgr.StoreIdentitySigner(p.SessionTopology, identitySigner)
					}
				}
			}
		}
	}

	if identitySigner == "" {
		identitySigner = topology.GlobalSigner.Hex()
	}

	blacklistObj := map[string][]common.Address{
		"blacklist": topology.Blacklist,
	}

	type sessionNodeType struct {
		Signer      string          `json:"signer"`
		ValueLimit  string          `json:"valueLimit"`
		Deadline    string          `json:"deadline"`
		Permissions []v3.Permission `json:"permissions"`
	}

	sessionNode := sessionNodeType{
		Signer:      session.Signer.Hex(),
		ValueLimit:  session.ValueLimit.String(),
		Deadline:    session.Deadline.String(),
		Permissions: session.Permissions,
	}

	secondElement := []interface{}{
		map[string]string{
			"identitySigner": identitySigner,
		},
		sessionNode,
	}

	for _, s := range topology.Sessions {
		node := sessionNodeType{
			Signer:      s.Signer.Hex(),
			ValueLimit:  s.Permissions.ValueLimit.String(),
			Deadline:    s.Permissions.Deadline.String(),
			Permissions: s.Permissions.Permissions,
		}
		secondElement = append(secondElement, node)
	}

	finalResult := []interface{}{
		blacklistObj,
		secondElement,
	}

	bytes, err := json.Marshal(finalResult)
	if err != nil {
		return "", fmt.Errorf("failed to marshal result: %w", err)
	}

	return string(bytes), nil
}

func handleRemoveExplicitSession(p *RemoveSessionParams) (string, error) {
	topology, err := v3.SessionsTopologyFromJSON(p.SessionTopology)
	if err != nil {
		return "", fmt.Errorf("failed to decode session topology: %w", err)
	}

	if !v3.IsSessionsTopology(topology) {
		return "", fmt.Errorf("session topology must be a valid session topology")
	}

	if p.ExplicitSessionAddress == "" || p.ExplicitSessionAddress[:2] != "0x" {
		return "", fmt.Errorf("explicit session address must be a valid address")
	}

	updated := v3.RemoveExplicitSession(topology, p.ExplicitSessionAddress)
	if updated == nil {
		return "", fmt.Errorf("session topology is empty")
	}

	return v3.SessionsTopologyToJSON(updated)
}

func handleEncodeCallSignature(p *EncodeCallSignatureParams) (string, error) {
	log.Printf("Encoding call signature with params: %+v", p)
	var sig v3.SessionCallSignature

	// Check if signature is in r:s:v format
	if strings.Contains(p.Signature, ":") {
		log.Printf("Processing signature in r:s:v format")
		// Split the signature into r, s, v components
		parts := strings.Split(p.Signature, ":")
		if len(parts) != 3 {
			return "", fmt.Errorf("invalid signature format, expected r:s:v")
		}

		r := common.FromHex(parts[0])
		s := common.FromHex(parts[1])
		v, err := strconv.Atoi(parts[2])
		if err != nil {
			return "", fmt.Errorf("invalid v value: %w", err)
		}

		sig.Signature = append(append(r, s...), byte(v))
		log.Printf("Created signature from r:s:v format: %x", sig.Signature)
	} else {
		log.Printf("Processing signature in hex format")
		// Handle hex-encoded signature
		if !strings.HasPrefix(p.Signature, "0x") {
			p.Signature = "0x" + p.Signature
		}
		sigBytes := common.FromHex(p.Signature)
		if len(sigBytes) != 65 {
			return "", fmt.Errorf("invalid signature length, expected 65 bytes, got %d", len(sigBytes))
		}
		sig.Signature = sigBytes
		log.Printf("Created signature from hex format: %x", sig.Signature)
	}

	// Encode the signature
	encoded, err := v3.EncodeSessionCallSignature(&sig, p.PermissionIndex)
	if err != nil {
		return "", fmt.Errorf("failed to encode call signature: %w", err)
	}

	result := "0x" + hex.EncodeToString(encoded)
	log.Printf("Encoded result: %s", result)
	return result, nil
}

// newSessionExplicitCmd defines the command-line interface.
func newSessionExplicitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "explicit",
		Short: "Explicit session utilities",
	}

	addCmd := &cobra.Command{
		Use:   "add [explicit-session] [session-topology]",
		Short: "Add an explicit session to a session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			explicitSession, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			sessionTopology, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}
			result, err := handleAddExplicitSession(&AddSessionParams{
				ExplicitSession: json.RawMessage(explicitSession),
				SessionTopology: json.RawMessage(sessionTopology),
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	removeCmd := &cobra.Command{
		Use:   "remove [explicit-session-address] [session-topology]",
		Short: "Remove a session from the session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("explicit session address is required")
			}

			topology, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}

			result, err := handleRemoveExplicitSession(&RemoveSessionParams{
				ExplicitSessionAddress: args[0],
				SessionTopology:        topology,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	encodeCallSignatureCmd := &cobra.Command{
		Use:   "encodeCallSignature [signature] [permission-index]",
		Short: "Encode a session call signature",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("signature and permission index are required")
			}

			permissionIndex, err := strconv.Atoi(args[1])
			if err != nil {
				return fmt.Errorf("invalid permission index: %w", err)
			}

			result, err := handleEncodeCallSignature(&EncodeCallSignatureParams{
				Signature:       args[0],
				PermissionIndex: permissionIndex,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(addCmd, removeCmd, encodeCallSignatureCmd)
	return cmd
}
