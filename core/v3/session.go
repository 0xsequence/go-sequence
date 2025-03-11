package v3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
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

// RSY represents an ECDSA signature with r, s, and yParity components
type RSY struct {
	R       [32]byte `json:"r"`
	S       [32]byte `json:"s"`
	YParity byte     `json:"yParity"`
}

// Attestation represents an attestation (adjust fields based on actual definition)
type Attestation struct {
	Data string `json:"data"`
}

// SessionCallSignature is an interface for both implicit and explicit signatures
type SessionCallSignature struct {
	IsImplicit        bool
	Attestation       *Attestation
	IdentitySignature *RSY   // Used if implicit
	PermissionIndex   string // Used if explicit
	SessionSignature  RSY
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

	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(data), &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if _, hasAttestation := parsed["attestation"]; hasAttestation {
		// Implicit signature
		attestationJSON, err := json.Marshal(parsed["attestation"])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal attestation: %w", err)
		}
		attestation, err := AttestationFromJSON(string(attestationJSON))
		if err != nil {
			return nil, fmt.Errorf("failed to parse attestation: %w", err)
		}

		identitySigStr, ok := parsed["identitySignature"].(string)
		if !ok {
			return nil, fmt.Errorf("identitySignature must be a string")
		}
		identitySignature, err := rsyFromRsvStr(identitySigStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse identitySignature: %w", err)
		}

		sessionSigStr, ok := parsed["sessionSignature"].(string)
		if !ok {
			return nil, fmt.Errorf("sessionSignature must be a string")
		}
		sessionSignature, err := rsyFromRsvStr(sessionSigStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sessionSignature: %w", err)
		}

		return &SessionCallSignature{
			IsImplicit:        true,
			Attestation:       &attestation,
			IdentitySignature: &identitySignature,
			SessionSignature:  sessionSignature,
		}, nil
	}

	// Explicit signature
	permissionIndex, ok := parsed["permissionIndex"].(string)
	if !ok {
		permissionIndexFloat, ok := parsed["permissionIndex"].(float64)
		if !ok {
			return nil, fmt.Errorf("permissionIndex must be a string or number")
		}
		permissionIndex = fmt.Sprintf("%d", int(permissionIndexFloat))
	}

	sessionSigStr, ok := parsed["sessionSignature"].(string)
	if !ok {
		return nil, fmt.Errorf("sessionSignature must be a string")
	}

	return processExplicitSignature(permissionIndex, sessionSigStr)
}

// processExplicitSignature processes an explicit session call signature
func processExplicitSignature(permissionIndexStr, sessionSignatureStr string) (*SessionCallSignature, error) {
	if !strings.HasPrefix(sessionSignatureStr, "0x") && !strings.Contains(sessionSignatureStr, ":") {
		sessionSignatureStr = "0x" + sessionSignatureStr
	}

	var rsy RSY
	if strings.Contains(sessionSignatureStr, ":") {
		var err error
		rsy, err = rsyFromRsvStr(sessionSignatureStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sessionSignature: %w", err)
		}
	} else {
		sigBytes := common.FromHex(sessionSignatureStr)
		if len(sigBytes) != 65 {
			return nil, fmt.Errorf("invalid signature length, expected 65 bytes, got %d bytes", len(sigBytes))
		}

		r := new(big.Int).SetBytes(sigBytes[:32])
		s := new(big.Int).SetBytes(sigBytes[32:64])

		var rBytes, sBytes [32]byte
		r.FillBytes(rBytes[:])
		s.FillBytes(sBytes[:])

		rsy = RSY{
			R:       rBytes,
			S:       sBytes,
			YParity: sigBytes[64] - 27, // Adjust v from 27/28 to 0/1
		}
	}

	return &SessionCallSignature{
		IsImplicit:       false,
		PermissionIndex:  permissionIndexStr,
		SessionSignature: rsy,
	}, nil
}

// rsyFromRsvStr converts a string in r:s:v format to an RSY struct
func rsyFromRsvStr(sigStr string) (RSY, error) {
	parts := strings.Split(sigStr, ":")
	if len(parts) != 3 {
		return RSY{}, fmt.Errorf("signature must be in r:s:v format")
	}

	rStr, sStr, vStr := parts[0], parts[1], parts[2]
	if !strings.HasPrefix(rStr, "0x") {
		rStr = "0x" + rStr
	}
	if !strings.HasPrefix(sStr, "0x") {
		sStr = "0x" + sStr
	}

	rBytes := common.FromHex(rStr)
	if len(rBytes) != 32 {
		return RSY{}, fmt.Errorf("invalid r size, expected 32 bytes, got %d bytes", len(rBytes))
	}
	sBytes := common.FromHex(sStr)
	if len(sBytes) != 32 {
		return RSY{}, fmt.Errorf("invalid s size, expected 32 bytes, got %d bytes", len(sBytes))
	}

	v, err := strconv.Atoi(vStr)
	if err != nil {
		return RSY{}, fmt.Errorf("invalid v value: %w", err)
	}
	if v != 27 && v != 28 {
		return RSY{}, fmt.Errorf("invalid v value, expected 27 or 28, got %d", v)
	}
	yParity := byte(v - 27)

	r := new(big.Int).SetBytes(rBytes)
	s := new(big.Int).SetBytes(sBytes)

	var r32, s32 [32]byte
	r.FillBytes(r32[:])
	s.FillBytes(s32[:])

	return RSY{
		R:       r32,
		S:       s32,
		YParity: yParity,
	}, nil
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

// EncodeSessionCallSignatures encodes a slice of session call signatures
func EncodeSessionCallSignatures(
	topology *SessionsTopology,
	callSignatures []SessionCallSignature,
	explicitSigners []common.Address,
	implicitSigners []common.Address,
) ([]byte, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}

	minimisedTopology := minimiseSessionsTopology(topology, explicitSigners, implicitSigners)
	topologyBytes, err := encodeMinimizedTopology(minimisedTopology)
	if err != nil {
		return nil, fmt.Errorf("failed to encode minimized topology: %w", err)
	}

	var parts [][]byte
	parts = append(parts, topologyBytes)

	for _, sig := range callSignatures {
		if sig.IsImplicit {
			flagByte := byte(0x80)
			parts = append(parts, []byte{flagByte})

			if sig.Attestation != nil {
				attBytes := []byte(sig.Attestation.Data)
				parts = append(parts, attBytes)
			}

			if sig.IdentitySignature != nil {
				idSigBytes, err := packRSY(*sig.IdentitySignature)
				if err != nil {
					return nil, fmt.Errorf("failed to pack identitySignature: %w", err)
				}
				parts = append(parts, idSigBytes)
			}

			sessionSigBytes, err := packRSY(sig.SessionSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to pack sessionSignature: %w", err)
			}
			parts = append(parts, sessionSigBytes)
		} else {
			permIndex, err := strconv.Atoi(sig.PermissionIndex)
			if err != nil {
				return nil, fmt.Errorf("invalid permissionIndex: %s", sig.PermissionIndex)
			}
			if permIndex > 127 {
				return nil, fmt.Errorf("permissionIndex too large: %d", permIndex)
			}
			flagByte := byte(permIndex)
			parts = append(parts, []byte{flagByte})

			rsyBytes, err := packRSY(sig.SessionSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to pack sessionSignature: %w", err)
			}
			parts = append(parts, rsyBytes)
		}
	}

	return bytes.Join(parts, nil), nil
}

// encodeMinimizedTopology encodes the minimized topology
func encodeMinimizedTopology(topology *SessionsTopology) ([]byte, error) {
	result := []byte{}

	result = append(result, topology.GlobalSigner.Bytes()...)

	sessionsCount := len(topology.Sessions)
	if sessionsCount > 255 {
		return nil, fmt.Errorf("too many sessions")
	}
	result = append(result, byte(sessionsCount))

	for _, session := range topology.Sessions {
		encodedNode, err := encodeSessionNode(session)
		if err != nil {
			return nil, err
		}
		result = append(result, encodedNode...)
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

// encodeSessionNode encodes a session node
func encodeSessionNode(node SessionNode) ([]byte, error) {
	result := []byte{}
	result = append(result, node.Signer.Bytes()...)

	if node.Permissions != nil {
		valueLimitBytes := make([]byte, 32)
		node.Permissions.ValueLimit.FillBytes(valueLimitBytes)
		result = append(result, valueLimitBytes...)

		deadlineBytes := make([]byte, 32)
		node.Permissions.Deadline.FillBytes(deadlineBytes)
		result = append(result, deadlineBytes...)

		permissionsCount := len(node.Permissions.Permissions)
		if permissionsCount > 255 {
			return nil, fmt.Errorf("too many permissions")
		}
		result = append(result, byte(permissionsCount))

		for _, perm := range node.Permissions.Permissions {
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
				result = append(result, rule.Value...)
				offsetBytes := make([]byte, 32)
				rule.Offset.FillBytes(offsetBytes)
				result = append(result, offsetBytes...)
				result = append(result, rule.Mask...)
			}
		}
	} else {
		result = append(result, 0)
	}

	// Encode children recursively
	childrenCount := len(node.Children)
	if childrenCount > 255 {
		return nil, fmt.Errorf("too many children")
	}
	result = append(result, byte(childrenCount))

	for _, child := range node.Children {
		childBytes, err := encodeSessionNode(child)
		if err != nil {
			return nil, err
		}
		result = append(result, childBytes...)
	}

	return result, nil
}

// encodeBlacklist encodes the blacklist into bytes
func encodeBlacklist(blacklist []common.Address) ([]byte, error) {
	result := []byte{}
	for _, addr := range blacklist {
		result = append(result, addr.Bytes()...)
	}
	return result, nil
}

// containsAddress checks if an address is in a list
func containsAddress(addrs []common.Address, addr common.Address) bool {
	for _, a := range addrs {
		if a == addr {
			return true
		}
	}
	return false
}

// Existing functions like packRSY remain unchanged unless needed
func packRSY(rsy RSY) ([]byte, error) {
	result := make([]byte, 64)
	copy(result[0:32], rsy.R[:])
	// Combine yParity with S[0] (adjust based on your encoding scheme)
	yParityAndS := rsy.S[0] | (rsy.YParity & 1)
	result[32] = yParityAndS
	copy(result[33:], rsy.S[1:])
	return result, nil
}

// minimiseSessionsTopology reduces the topology based on signers
func minimiseSessionsTopology(
	topology *SessionsTopology,
	explicitSigners []common.Address,
	implicitSigners []common.Address,
) *SessionsTopology {
	if topology == nil {
		return nil
	}

	result := &SessionsTopology{
		GlobalSigner: topology.GlobalSigner,
		Blacklist:    topology.Blacklist,
		Sessions:     make([]SessionNode, 0, len(topology.Sessions)),
	}

	for _, session := range topology.Sessions {
		minimisedSession := minimiseSessionNode(session, explicitSigners, implicitSigners)
		result.Sessions = append(result.Sessions, minimisedSession)
	}

	if len(implicitSigners) == 0 && len(result.Blacklist) > 0 {
		blacklistBytes, _ := encodeBlacklist(result.Blacklist)
		log.Println("blacklistBytes", blacklistBytes)

		result.Blacklist = []common.Address{}

		result.Sessions = append(result.Sessions, SessionNode{
			Signer: common.Address{},
			Permissions: &SessionPermissions{
				ValueLimit:  big.NewInt(0),
				Deadline:    big.NewInt(0),
				Permissions: nil,
			},
			Children: []SessionNode{},
		})
	}

	return result
}

// minimiseSessionNode minimizes a single session node
func minimiseSessionNode(
	node SessionNode,
	explicitSigners []common.Address,
	implicitSigners []common.Address,
) SessionNode {
	if containsAddress(explicitSigners, node.Signer) {
		return node
	}

	if node.Permissions != nil || len(node.Children) > 0 {
		encoded, err := encodeSessionNode(node)
		if err != nil {
			return node
		}
		hashed := crypto.Keccak256(encoded)
		log.Println("hashed", hashed)
		return SessionNode{
			Signer:      common.Address{},
			Permissions: nil,
			Children:    nil,
		}
	}

	log.Println("node", node)
	log.Println("explicitSigners", explicitSigners)
	log.Println("implicitSigners", implicitSigners)

	return node
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

// AttestationFromJSON converts JSON to an Attestation (placeholder)
func AttestationFromJSON(jsonStr string) (Attestation, error) {
	var attestation Attestation
	if err := json.Unmarshal([]byte(jsonStr), &attestation); err != nil {
		return Attestation{}, fmt.Errorf("failed to parse attestation: %w", err)
	}
	return attestation, nil
}

// MarshalJSON implements json.Marshaler for SessionsTopology
func (s *SessionsTopology) MarshalJSON() ([]byte, error) {
	blacklistObj := map[string][]common.Address{
		"blacklist": s.Blacklist,
	}

	secondElement := make([]interface{}, 0, len(s.Sessions)+1)
	secondElement = append(secondElement, map[string]string{
		"globalSigner": s.GlobalSigner.String(),
	})
	for _, session := range s.Sessions {
		secondElement = append(secondElement, session)
	}

	result := []interface{}{
		blacklistObj,
		secondElement,
	}
	return json.Marshal(result)
}

// UnmarshalJSON implements json.Unmarshaler for SessionsTopology
func (s *SessionsTopology) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return fmt.Errorf("failed to unmarshal sessions topology array: %w", err)
	}

	if len(arr) != 2 {
		return fmt.Errorf("invalid sessions topology array length: expected 2, got %d", len(arr))
	}

	var blacklistObj struct {
		Blacklist []common.Address `json:"blacklist"`
	}
	if err := json.Unmarshal(arr[0], &blacklistObj); err != nil {
		return fmt.Errorf("failed to unmarshal blacklist: %w", err)
	}
	s.Blacklist = blacklistObj.Blacklist

	var second json.RawMessage
	if err := json.Unmarshal(arr[1], &second); err != nil {
		return fmt.Errorf("failed to unmarshal second element: %w", err)
	}

	var secondArr []json.RawMessage
	if err := json.Unmarshal(second, &secondArr); err == nil && len(secondArr) > 0 {
		var globalSignerObj struct {
			GlobalSigner string `json:"globalSigner"`
		}
		if err := json.Unmarshal(secondArr[0], &globalSignerObj); err != nil {
			return fmt.Errorf("failed to unmarshal globalSigner: %w", err)
		}
		s.GlobalSigner = common.HexToAddress(globalSignerObj.GlobalSigner)

		s.Sessions = make([]SessionNode, 0, len(secondArr)-1)
		for i := 1; i < len(secondArr); i++ {
			var session SessionNode
			if err := json.Unmarshal(secondArr[i], &session); err != nil {
				return fmt.Errorf("failed to unmarshal session at index %d: %w", i-1, err)
			}
			s.Sessions = append(s.Sessions, session)
		}
	} else {
		var identitySignerObj struct {
			IdentitySigner string `json:"identitySigner"`
		}
		if err := json.Unmarshal(second, &identitySignerObj); err != nil {
			return fmt.Errorf("failed to unmarshal identitySigner: %w", err)
		}
		if identitySignerObj.IdentitySigner == "" {
			return fmt.Errorf("missing identitySigner in second element")
		}
		s.GlobalSigner = common.HexToAddress(identitySignerObj.IdentitySigner)
		s.Sessions = make([]SessionNode, 0)
	}

	return nil
}

// EncodeSessionCallSignature encodes a session call signature with a permission index
func EncodeSessionCallSignature(sig *SessionCallSignature, permissionIndex int) ([]byte, error) {
	if permissionIndex < 0 {
		return nil, fmt.Errorf("permission index must be non-negative")
	}

	result := []byte{byte(permissionIndex)}

	sigBytes, err := packRSY(sig.SessionSignature)
	if err != nil {
		return nil, fmt.Errorf("failed to pack signature: %w", err)
	}

	result = append(result, sigBytes...)

	return result, nil
}
