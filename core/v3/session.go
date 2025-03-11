package v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// SessionsTopology represents the topology of sessions
type SessionsTopology struct {
	GlobalSigner common.Address   `json:"globalSigner"`
	Sessions     []SessionNode    `json:"sessions"`
	Blacklist    []common.Address `json:"blacklist"`
}

// SessionNode represents a node in the session topology
type SessionNode struct {
	Signer      common.Address      `json:"signer"`
	Permissions *SessionPermissions `json:"permissions"`
	Children    []SessionNode       `json:"children"`
}

// SessionCallSignature represents a call signature for a session
type SessionCallSignature struct {
	Signer    common.Address `json:"signer"`
	Signature []byte         `json:"signature"`
}

// ConfigurationTree represents the configuration tree for sessions
type ConfigurationTree struct {
	ImageHash [32]byte      `json:"imageHash"`
	Nodes     []SessionNode `json:"nodes"`
}

// EmptySessionsTopology creates an empty session topology with the given global signer
func EmptySessionsTopology(signer common.Address) *SessionsTopology {
	return &SessionsTopology{
		GlobalSigner: signer,
		Sessions:     []SessionNode{},
		Blacklist:    []common.Address{},
	}
}

// SessionsTopologyToJSON converts a session topology to JSON
func SessionsTopologyToJSON(topology *SessionsTopology) (string, error) {
	result := []interface{}{
		map[string]interface{}{"blacklist": topology.Blacklist},
		[]interface{}{
			map[string]interface{}{"identitySigner": topology.GlobalSigner.Hex()},
		},
	}

	secondArray := result[1].([]interface{})
	for _, session := range topology.Sessions {
		sessionMap := map[string]interface{}{
			"signer": session.Signer,
		}

		if session.Permissions != nil {
			sessionMap["valueLimit"] = session.Permissions.ValueLimit
			sessionMap["deadline"] = session.Permissions.Deadline

			if len(session.Permissions.Permissions) > 0 {
				sessionMap["permissions"] = session.Permissions.Permissions
			}
		}

		if len(session.Children) > 0 {
			sessionMap["children"] = session.Children
		}

		secondArray = append(secondArray, sessionMap)
	}

	// Marshal the structured representation
	bytes, err := json.Marshal(result)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session topology: %w", err)
	}
	return string(bytes), nil
}

// SessionsTopologyFromJSON converts JSON to a session topology
func SessionsTopologyFromJSON(data string) (*SessionsTopology, error) {
	if len(data) > 0 && data[0] == '[' {
		var topArray []json.RawMessage
		if err := json.Unmarshal([]byte(data), &topArray); err == nil && len(topArray) >= 2 {
			var blacklistObj struct {
				Blacklist []common.Address `json:"blacklist"`
			}
			if err := json.Unmarshal(topArray[0], &blacklistObj); err == nil {
				topology := &SessionsTopology{
					Blacklist: blacklistObj.Blacklist,
					Sessions:  []SessionNode{},
				}

				var secondArray []json.RawMessage
				if err := json.Unmarshal(topArray[1], &secondArray); err == nil && len(secondArray) > 0 {
					var identitySignerStrObj struct {
						IdentitySigner string `json:"identitySigner"`
						GlobalSigner   string `json:"globalSigner"`
					}

					if err := json.Unmarshal(secondArray[0], &identitySignerStrObj); err == nil {
						if identitySignerStrObj.IdentitySigner != "" {
							topology.GlobalSigner = common.HexToAddress(identitySignerStrObj.IdentitySigner)
						} else if identitySignerStrObj.GlobalSigner != "" {
							topology.GlobalSigner = common.HexToAddress(identitySignerStrObj.GlobalSigner)
						}
					} else {
						var identitySignerObj struct {
							IdentitySigner common.Address `json:"identitySigner"`
							GlobalSigner   common.Address `json:"globalSigner"`
						}

						if err := json.Unmarshal(secondArray[0], &identitySignerObj); err == nil {
							if identitySignerObj.IdentitySigner != (common.Address{}) {
								topology.GlobalSigner = identitySignerObj.IdentitySigner
							} else {
								topology.GlobalSigner = identitySignerObj.GlobalSigner
							}
						}
					}

					for i := 1; i < len(secondArray); i++ {
						var sessionInfo struct {
							Signer      common.Address `json:"signer"`
							ValueLimit  string         `json:"valueLimit"`
							Deadline    string         `json:"deadline"`
							Permissions []Permission   `json:"permissions"`
						}

						if err := json.Unmarshal(secondArray[i], &sessionInfo); err == nil {
							session := SessionNode{
								Signer: sessionInfo.Signer,
							}

							if sessionInfo.ValueLimit != "" && sessionInfo.Deadline != "" {
								valueLimit, _ := new(big.Int).SetString(sessionInfo.ValueLimit, 10)
								deadline, _ := new(big.Int).SetString(sessionInfo.Deadline, 10)

								session.Permissions = &SessionPermissions{
									ValueLimit:  valueLimit,
									Deadline:    deadline,
									Permissions: sessionInfo.Permissions,
								}
							}

							topology.Sessions = append(topology.Sessions, session)
						}
					}

					return topology, nil
				}
			}
		}
	}

	var topology SessionsTopology
	if err := json.Unmarshal([]byte(data), &topology); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session topology: %w", err)
	}
	return &topology, nil
}

// SessionCallSignatureFromJSON converts JSON to a session call signature
func SessionCallSignatureFromJSON(data string) (*SessionCallSignature, error) {
	data = strings.TrimSpace(data)
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var inner string
		if err := json.Unmarshal([]byte(data), &inner); err == nil {
			data = inner
		} else {
			return nil, fmt.Errorf("failed to unquote data: %w", err)
		}
	}
	if len(data) > 0 && data[0] == '{' {
		type ExplicitSig struct {
			PermissionIndex  string `json:"permissionIndex"`
			SessionSignature string `json:"sessionSignature"`
		}
		var explicitSig ExplicitSig
		if err := json.Unmarshal([]byte(data), &explicitSig); err != nil {
			return nil, fmt.Errorf("failed to parse explicit signature: %w", err)
		}
		if explicitSig.PermissionIndex == "" || explicitSig.SessionSignature == "" {
			return nil, fmt.Errorf("missing permissionIndex or sessionSignature")
		}
		return processExplicitSignature(explicitSig.PermissionIndex, explicitSig.SessionSignature)
	}

	return nil, fmt.Errorf("invalid signature format: %s", data)
}

// Helper function to process explicit signatures
func processExplicitSignature(permissionIndexStr, sessionSignatureStr string) (*SessionCallSignature, error) {
	if !strings.HasPrefix(sessionSignatureStr, "0x") && !strings.Contains(sessionSignatureStr, ":") {
		sessionSignatureStr = "0x" + sessionSignatureStr
	}

	var signature []byte

	if strings.Contains(sessionSignatureStr, ":") {
		parts := strings.Split(sessionSignatureStr, ":")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid signature format, expected r:s:v, got: %s", sessionSignatureStr)
		}

		r := common.FromHex(ensureHexPrefix(parts[0]))
		s := common.FromHex(ensureHexPrefix(parts[1]))

		if len(r) != 32 {
			return nil, fmt.Errorf("invalid r size, expected 32 bytes, got %d bytes", len(r))
		}
		if len(s) != 32 {
			return nil, fmt.Errorf("invalid s size, expected 32 bytes, got %d bytes", len(s))
		}

		v, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, fmt.Errorf("invalid v value: %w", err)
		}

		signature = append(append(r, s...), byte(v))
	} else {
		signature = common.FromHex(sessionSignatureStr)
		if len(signature) != 65 {
			return nil, fmt.Errorf("invalid signature length, expected 65 bytes, got %d bytes", len(signature))
		}
	}

	baseSig := &SessionCallSignature{
		Signature: signature,
	}

	return ProcessExplicitSessionCallSignature(baseSig, permissionIndexStr)
}

func ensureHexPrefix(s string) string {
	if !strings.HasPrefix(s, "0x") {
		return "0x" + s
	}
	return s
}

func IsSessionsTopology(v interface{}) bool {
	_, ok := v.(*SessionsTopology)
	return ok
}

func GetSessionPermissions(topology *SessionsTopology, signer common.Address) *SessionPermissions {
	for _, session := range topology.Sessions {
		if session.Signer == signer {
			return session.Permissions
		}
	}
	return nil
}

func MergeSessionsTopologies(a, b *SessionsTopology) *SessionsTopology {
	result := &SessionsTopology{
		GlobalSigner: a.GlobalSigner,
		Sessions:     append([]SessionNode{}, a.Sessions...),
		Blacklist:    append([]common.Address{}, a.Blacklist...),
	}
	result.Sessions = append(result.Sessions, b.Sessions...)
	result.Blacklist = append(result.Blacklist, b.Blacklist...)
	return result
}

// BalanceSessionsTopology balances a session topology
func BalanceSessionsTopology(topology *SessionsTopology) *SessionsTopology {
	return topology
}

// SessionsTopologyToConfigurationTree converts a session topology to a configuration tree
func SessionsTopologyToConfigurationTree(topology *SessionsTopology) *ConfigurationTree {
	tree := SessionsTopologyToTree(topology)

	hashBytes := Hash(tree)

	var imageHash [32]byte
	copy(imageHash[:], hashBytes)

	return &ConfigurationTree{
		ImageHash: imageHash,
		Nodes:     topology.Sessions,
	}
}

// HashConfigurationTree computes the hash of a configuration tree
func HashConfigurationTree(tree *ConfigurationTree) [32]byte {
	if tree == nil {
		return [32]byte{}
	}

	if tree.ImageHash != [32]byte{} {
		return tree.ImageHash
	}

	topology := &SessionsTopology{
		Sessions: tree.Nodes,
	}

	genericTree := SessionsTopologyToTree(topology)
	hashBytes := Hash(genericTree)

	var result [32]byte
	copy(result[:], hashBytes)

	return result
}

// EncodeSessionCallSignatures encodes session call signatures
func EncodeSessionCallSignatures(sigs []SessionCallSignature, topology *SessionsTopology, explicitSigners, implicitSigners []common.Address) ([]byte, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}

	result := []byte{}

	topologyBytes, err := encodeSessionsTopology(topology)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topology: %w", err)
	}

	topologyLength := len(topologyBytes)
	topologyLengthBytes := []byte{
		byte(topologyLength >> 16),
		byte(topologyLength >> 8),
		byte(topologyLength),
	}
	result = append(result, topologyLengthBytes...)
	result = append(result, topologyBytes...)

	result = append(result, 0)

	for _, sig := range sigs {
		result = append(result, sig.Signature...)
	}

	return result, nil
}

func encodeSessionsTopology(topology *SessionsTopology) ([]byte, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}

	result := []byte{}

	result = append(result, topology.GlobalSigner.Bytes()...)

	sessionsCount := len(topology.Sessions)
	if sessionsCount > 255 {
		return nil, fmt.Errorf("too many sessions")
	}
	result = append(result, byte(sessionsCount))

	for _, session := range topology.Sessions {
		result = append(result, session.Signer.Bytes()...)

		if session.Permissions != nil {
			valueLimitBytes := make([]byte, 32)
			session.Permissions.ValueLimit.FillBytes(valueLimitBytes)
			result = append(result, valueLimitBytes...)

			deadlineBytes := make([]byte, 32)
			session.Permissions.Deadline.FillBytes(deadlineBytes)
			result = append(result, deadlineBytes...)

			permissionsCount := len(session.Permissions.Permissions)
			if permissionsCount > 255 {
				return nil, fmt.Errorf("too many permissions")
			}
			result = append(result, byte(permissionsCount))

			for _, perm := range session.Permissions.Permissions {
				result = append(result, perm.Target.Bytes()...)

				rulesCount := len(perm.Rules)
				if rulesCount > 255 {
					return nil, fmt.Errorf("too many rules")
				}
				result = append(result, byte(rulesCount))

				for _, rule := range perm.Rules {
					cumulativeByte := byte(0)
					if rule.Cumulative {
						cumulativeByte = 1
					}
					result = append(result, cumulativeByte)

					result = append(result, byte(rule.Operation))

					if len(rule.Value) != 32 {
						return nil, fmt.Errorf("invalid value length: expected 32 bytes, got %d", len(rule.Value))
					}
					result = append(result, rule.Value...)

					offsetBytes := make([]byte, 32)
					rule.Offset.FillBytes(offsetBytes)
					result = append(result, offsetBytes...)

					if len(rule.Mask) != 32 {
						return nil, fmt.Errorf("invalid mask length: expected 32 bytes, got %d", len(rule.Mask))
					}
					result = append(result, rule.Mask...)
				}
			}
		} else {
			result = append(result, 0)
		}
	}

	blacklistCount := len(topology.Blacklist)
	if blacklistCount > 255 {
		return nil, fmt.Errorf("too many blacklisted addresses")
	}
	result = append(result, byte(blacklistCount))

	for _, addr := range topology.Blacklist {
		result = append(result, addr.Bytes()...)
	}

	return result, nil
}

// AddToImplicitBlacklist adds an address to the implicit blacklist
func AddToImplicitBlacklist(topology *SessionsTopology, addr common.Address) *SessionsTopology {
	result := &SessionsTopology{
		GlobalSigner: topology.GlobalSigner,
		Sessions:     append([]SessionNode{}, topology.Sessions...),
		Blacklist:    append([]common.Address{}, topology.Blacklist...),
	}
	result.Blacklist = append(result.Blacklist, addr)
	return result
}

// RemoveFromImplicitBlacklist removes an address from the implicit blacklist
func RemoveFromImplicitBlacklist(topology *SessionsTopology, addr common.Address) *SessionsTopology {
	result := &SessionsTopology{
		GlobalSigner: topology.GlobalSigner,
		Sessions:     append([]SessionNode{}, topology.Sessions...),
		Blacklist:    make([]common.Address, 0, len(topology.Blacklist)),
	}
	for _, a := range topology.Blacklist {
		if a != addr {
			result.Blacklist = append(result.Blacklist, a)
		}
	}
	return result
}

// RemoveExplicitSession removes an explicit session from the topology
func RemoveExplicitSession(topology *SessionsTopology, addr string) *SessionsTopology {
	addrBytes := common.HexToAddress(addr)
	result := &SessionsTopology{
		GlobalSigner: topology.GlobalSigner,
		Sessions:     make([]SessionNode, 0, len(topology.Sessions)),
		Blacklist:    append([]common.Address{}, topology.Blacklist...),
	}
	for _, session := range topology.Sessions {
		if session.Signer != addrBytes {
			result.Sessions = append(result.Sessions, session)
		}
	}
	if len(result.Sessions) == 0 {
		return nil
	}
	return result
}

// ProcessExplicitSessionCallSignature processes a session call signature with a permission index
func ProcessExplicitSessionCallSignature(sig *SessionCallSignature, permissionIndexStr string) (*SessionCallSignature, error) {
	permissionIndex, err := strconv.Atoi(permissionIndexStr)
	if err != nil {
		return nil, fmt.Errorf("invalid permission index: %w", err)
	}
	if permissionIndex > 255 {
		return nil, fmt.Errorf("permission index out of range")
	}

	// Prepend permision index to the signature
	result := []byte{byte(permissionIndex)}
	result = append(result, sig.Signature...)

	return &SessionCallSignature{
		Signer:    sig.Signer,
		Signature: result,
	}, nil
}
