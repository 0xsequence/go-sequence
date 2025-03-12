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
// This implementation matches the TypeScript implementation in sequence-primitives
func EncodeSessionCallSignatures(
	callSignatures []SessionCallSignature,
	topology *SessionsTopology,
	explicitSigners []common.Address,
	implicitSigners []common.Address,
) ([]byte, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}

	// Optimize the configuration tree by rolling unused signers into nodes
	minimisedTopology := minimiseSessionsTopology(topology, explicitSigners, implicitSigners)

	// Encode the topology
	encodedTopology, err := encodeSessionsTopology(minimisedTopology)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topology: %w", err)
	}

	// Check if topology is too large
	if len(encodedTopology) > 0xFFFFFF { // 3 bytes max
		return nil, fmt.Errorf("session topology is too large")
	}

	var parts [][]byte

	topologyLengthBytes := []byte{
		byte((len(encodedTopology) >> 16) & 0xFF),
		byte((len(encodedTopology) >> 8) & 0xFF),
		byte(len(encodedTopology) & 0xFF),
	}
	parts = append(parts, topologyLengthBytes, encodedTopology)

	type attestationKey struct {
		data string
	}
	attestationMap := make(map[attestationKey]int)
	var encodedAttestations [][]byte

	// Process all implicit signatures to build the attestation list
	for _, sig := range callSignatures {
		if sig.IsImplicit && sig.Attestation != nil {
			key := attestationKey{data: sig.Attestation.Data}
			if _, exists := attestationMap[key]; !exists {
				// Add to attestation map
				attestationMap[key] = len(encodedAttestations)

				// Encode attestation + identity signature
				var attestationBytes []byte
				attestationBytes = append(attestationBytes, []byte(sig.Attestation.Data)...)

				if sig.IdentitySignature != nil {
					identitySigBytes, err := packRSY(*sig.IdentitySignature)
					if err != nil {
						return nil, fmt.Errorf("failed to pack identity signature: %w", err)
					}
					attestationBytes = append(attestationBytes, identitySigBytes...)
				}

				encodedAttestations = append(encodedAttestations, attestationBytes)
			}
		}
	}

	// Check attestation count
	if len(encodedAttestations) >= 128 {
		return nil, fmt.Errorf("too many attestations")
	}

	// Add attestation count as 1 byte
	parts = append(parts, []byte{byte(len(encodedAttestations))})

	// Add all encoded attestations
	for _, attestation := range encodedAttestations {
		parts = append(parts, attestation)
	}

	// Process call signatures
	for _, sig := range callSignatures {
		if sig.IsImplicit {
			// Implicit signature
			if sig.Attestation == nil {
				return nil, fmt.Errorf("implicit signature missing attestation")
			}

			key := attestationKey{data: sig.Attestation.Data}
			attestationIndex, exists := attestationMap[key]
			if !exists {
				return nil, fmt.Errorf("failed to find attestation index for signature")
			}

			if attestationIndex >= 128 {
				return nil, fmt.Errorf("attestation index too large: %d", attestationIndex)
			}

			flagByte := byte(0x80 | attestationIndex)
			parts = append(parts, []byte{flagByte})

			// Add session signature
			sessionSigBytes, err := packRSY(sig.SessionSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to pack session signature: %w", err)
			}
			parts = append(parts, sessionSigBytes)

		} else {
			// Explicit signature
			permIndex, err := strconv.Atoi(sig.PermissionIndex)
			if err != nil {
				return nil, fmt.Errorf("invalid permission index: %s", sig.PermissionIndex)
			}

			if permIndex > 127 {
				return nil, fmt.Errorf("permission index too large: %d", permIndex)
			}

			flagByte := byte(permIndex)
			parts = append(parts, []byte{flagByte})

			sessionSigBytes, err := packRSY(sig.SessionSignature)
			if err != nil {
				return nil, fmt.Errorf("failed to pack session signature: %w", err)
			}
			parts = append(parts, sessionSigBytes)
		}
	}

	return bytes.Join(parts, nil), nil
}

// encodeSessionsTopology encodes the session topology into bytes
// This matches the TypeScript implementation of encodeSessionsTopology
func encodeSessionsTopology(topology *SessionsTopology) ([]byte, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology cannot be nil")
	}

	if len(topology.Sessions) > 1 {
		var encodedBranches [][]byte

		identitySignerBytes := encodeIdentitySignerLeaf(topology.GlobalSigner)
		encodedBranches = append(encodedBranches, identitySignerBytes)

		blacklistBytes := encodeBlacklistLeaf(topology.Blacklist)
		encodedBranches = append(encodedBranches, blacklistBytes)

		for _, session := range topology.Sessions {
			sessionBytes, err := encodeSessionPermissions(session)
			if err != nil {
				return nil, fmt.Errorf("failed to encode session: %w", err)
			}
			encodedBranches = append(encodedBranches, sessionBytes)
		}

		encoded := bytes.Join(encodedBranches, nil)

		encodedSize := minBytesFor(uint64(len(encoded)))
		if encodedSize > 15 {
			return nil, fmt.Errorf("branch too large")
		}

		flagByte := byte((SESSIONS_FLAG_BRANCH << 4) | encodedSize)

		lengthBytes := make([]byte, encodedSize)
		encodeLength(len(encoded), lengthBytes)

		result := []byte{flagByte}
		result = append(result, lengthBytes...)
		result = append(result, encoded...)
		return result, nil
	}

	if len(topology.Sessions) == 1 {
		return encodeSessionPermissions(topology.Sessions[0])
	}

	return encodeIdentitySignerLeaf(topology.GlobalSigner), nil
}

func encodeIdentitySignerLeaf(signer common.Address) []byte {
	flagByte := byte(SESSIONS_FLAG_IDENTITY_SIGNER << 4)
	result := []byte{flagByte}
	result = append(result, signer.Bytes()...)
	return result
}

// encodeBlacklistLeaf encodes a blacklist leaf
func encodeBlacklistLeaf(blacklist []common.Address) []byte {
	if len(blacklist) > 14 {
		flagByte := byte((SESSIONS_FLAG_BLACKLIST << 4) | 0x0F)
		result := []byte{flagByte, byte(len(blacklist))}
		for _, addr := range blacklist {
			result = append(result, addr.Bytes()...)
		}
		return result
	}

	flagByte := byte((SESSIONS_FLAG_BLACKLIST << 4) | len(blacklist))
	result := []byte{flagByte}
	for _, addr := range blacklist {
		result = append(result, addr.Bytes()...)
	}
	return result
}

// encodeSessionPermissions encodes session permissions
func encodeSessionPermissions(session SessionNode) ([]byte, error) {
	if session.Permissions == nil {
		return nil, fmt.Errorf("session permissions cannot be nil")
	}

	flagByte := byte(SESSIONS_FLAG_PERMISSIONS << 4)
	result := []byte{flagByte}

	permBytes, err := encodePermissionsData(session)
	if err != nil {
		return nil, err
	}

	result = append(result, permBytes...)
	return result, nil
}

// encodePermissionsData encodes the internal permission data
func encodePermissionsData(session SessionNode) ([]byte, error) {
	if session.Permissions == nil {
		return nil, fmt.Errorf("permissions cannot be nil")
	}

	var result []byte

	result = append(result, session.Signer.Bytes()...)

	valueLimitBytes := make([]byte, 32)
	if session.Permissions.ValueLimit != nil {
		session.Permissions.ValueLimit.FillBytes(valueLimitBytes)
	}
	result = append(result, valueLimitBytes...)

	deadlineBytes := make([]byte, 32)
	if session.Permissions.Deadline != nil {
		session.Permissions.Deadline.FillBytes(deadlineBytes)
	}
	result = append(result, deadlineBytes...)

	if len(session.Permissions.Permissions) > 255 {
		return nil, fmt.Errorf("too many permissions: %d", len(session.Permissions.Permissions))
	}
	result = append(result, byte(len(session.Permissions.Permissions)))

	for _, perm := range session.Permissions.Permissions {
		result = append(result, perm.Target.Bytes()...)

		if len(perm.Rules) > 255 {
			return nil, fmt.Errorf("too many rules: %d", len(perm.Rules))
		}
		result = append(result, byte(len(perm.Rules)))

		for _, rule := range perm.Rules {
			cumByte := byte(0)
			if rule.Cumulative {
				cumByte = 1
			}
			result = append(result, cumByte)

			result = append(result, byte(rule.Operation))

			result = append(result, rule.Value...)

			offsetBytes := make([]byte, 32)
			rule.Offset.FillBytes(offsetBytes)
			result = append(result, offsetBytes...)

			result = append(result, rule.Mask...)
		}
	}

	return result, nil
}

// encodeLength encodes a length value into the provided byte slice
func encodeLength(length int, dest []byte) {
	for i := 0; i < len(dest); i++ {
		dest[len(dest)-1-i] = byte(length & 0xFF)
		length >>= 8
	}
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
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) == 2 {
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

	type SessionsTopologyAlias SessionsTopology
	var topology SessionsTopologyAlias
	if err := json.Unmarshal(data, &topology); err != nil {
		return fmt.Errorf("failed to unmarshal session topology: %w", err)
	}

	*s = SessionsTopology(topology)
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

// UnmarshalJSON implements json.Unmarshaler for SessionNode
func (s *SessionNode) UnmarshalJSON(data []byte) error {
	type SessionNodeAlias SessionNode
	var node SessionNodeAlias

	if err := json.Unmarshal(data, &node); err == nil {
		*s = SessionNode(node)
		return nil
	}

	var objMap map[string]json.RawMessage
	if err := json.Unmarshal(data, &objMap); err != nil {
		return fmt.Errorf("failed to unmarshal session node: %w", err)
	}

	if signerBytes, ok := objMap["signer"]; ok {
		var signerStr string
		if err := json.Unmarshal(signerBytes, &signerStr); err == nil {
			s.Signer = common.HexToAddress(signerStr)
		} else {
			if err := json.Unmarshal(signerBytes, &s.Signer); err != nil {
				return fmt.Errorf("failed to unmarshal signer: %w", err)
			}
		}
	}

	hasPermissions := false

	_, hasValueLimit := objMap["valueLimit"]
	_, hasDeadline := objMap["deadline"]
	_, hasPermsList := objMap["permissions"]

	if hasValueLimit || hasDeadline || hasPermsList {
		hasPermissions = true
		s.Permissions = &SessionPermissions{}

		if valLimitBytes, ok := objMap["valueLimit"]; ok {
			var valLimitStr string
			if err := json.Unmarshal(valLimitBytes, &valLimitStr); err == nil {
				s.Permissions.ValueLimit, _ = new(big.Int).SetString(valLimitStr, 10)
			} else {
				var valLimit big.Int
				if err := json.Unmarshal(valLimitBytes, &valLimit); err != nil {
					return fmt.Errorf("failed to unmarshal valueLimit: %w", err)
				}
				s.Permissions.ValueLimit = &valLimit
			}
		}

		if deadlineBytes, ok := objMap["deadline"]; ok {
			var deadlineStr string
			if err := json.Unmarshal(deadlineBytes, &deadlineStr); err == nil {
				s.Permissions.Deadline, _ = new(big.Int).SetString(deadlineStr, 10)
			} else {
				var deadline big.Int
				if err := json.Unmarshal(deadlineBytes, &deadline); err != nil {
					return fmt.Errorf("failed to unmarshal deadline: %w", err)
				}
				s.Permissions.Deadline = &deadline
			}
		}

		if permsBytes, ok := objMap["permissions"]; ok {
			if err := json.Unmarshal(permsBytes, &s.Permissions.Permissions); err != nil {
				return fmt.Errorf("failed to unmarshal permissions list: %w", err)
			}
		}
	}

	if childrenBytes, ok := objMap["children"]; ok {
		if err := json.Unmarshal(childrenBytes, &s.Children); err != nil {
			return fmt.Errorf("failed to unmarshal children: %w", err)
		}
	}

	if hasPermissions {
		s.Permissions = &SessionPermissions{}
	}

	return nil
}

func (s *SessionCallSignature) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	dataStr = strings.TrimSpace(dataStr)
	if len(dataStr) >= 2 && dataStr[0] == '"' && dataStr[len(dataStr)-1] == '"' {
		var inner string
		if err := json.Unmarshal(data, &inner); err == nil {
			dataStr = inner
			data = []byte(inner)
		} else {
			return fmt.Errorf("failed to unquote data: %w", err)
		}
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	if _, hasAttestation := parsed["attestation"]; hasAttestation {
		attestationJSON, err := json.Marshal(parsed["attestation"])
		if err != nil {
			return fmt.Errorf("failed to marshal attestation: %w", err)
		}
		attestation, err := AttestationFromJSON(string(attestationJSON))
		if err != nil {
			return fmt.Errorf("failed to parse attestation: %w", err)
		}

		identitySigStr, ok := parsed["identitySignature"].(string)
		if !ok {
			return fmt.Errorf("identitySignature must be a string")
		}
		identitySignature, err := rsyFromRsvStr(identitySigStr)
		if err != nil {
			return fmt.Errorf("failed to parse identitySignature: %w", err)
		}

		sessionSigStr, ok := parsed["sessionSignature"].(string)
		if !ok {
			return fmt.Errorf("sessionSignature must be a string")
		}
		sessionSignature, err := rsyFromRsvStr(sessionSigStr)
		if err != nil {
			return fmt.Errorf("failed to parse sessionSignature: %w", err)
		}

		s.IsImplicit = true
		s.Attestation = &attestation
		s.IdentitySignature = &identitySignature
		s.SessionSignature = sessionSignature
		return nil
	}

	permissionIndex, ok := parsed["permissionIndex"].(string)
	if !ok {
		permissionIndexFloat, ok := parsed["permissionIndex"].(float64)
		if !ok {
			return fmt.Errorf("permissionIndex must be a string or number")
		}
		permissionIndex = fmt.Sprintf("%d", int(permissionIndexFloat))
	}

	sessionSigStr, ok := parsed["sessionSignature"].(string)
	if !ok {
		return fmt.Errorf("sessionSignature must be a string")
	}

	var sessionSignature RSY
	if strings.HasPrefix(sessionSigStr, "0x") && strings.Contains(sessionSigStr, ":") {
		var err error
		sessionSignature, err = rsyFromRsvStr(sessionSigStr)
		if err != nil {
			return fmt.Errorf("failed to parse sessionSignature: %w", err)
		}
	} else if strings.Contains(sessionSigStr, ":") {
		parts := strings.Split(sessionSigStr, ":")
		if len(parts) != 3 {
			return fmt.Errorf("signature must be in r:s:v format")
		}

		if !strings.HasPrefix(parts[0], "0x") {
			parts[0] = "0x" + parts[0]
		}
		if !strings.HasPrefix(parts[1], "0x") {
			parts[1] = "0x" + parts[1]
		}

		fullSig := strings.Join(parts, ":")
		var err error
		sessionSignature, err = rsyFromRsvStr(fullSig)
		if err != nil {
			return fmt.Errorf("failed to parse sessionSignature: %w", err)
		}
	} else {
		if !strings.HasPrefix(sessionSigStr, "0x") {
			sessionSigStr = "0x" + sessionSigStr
		}

		sigBytes := common.FromHex(sessionSigStr)
		if len(sigBytes) != 65 {
			return fmt.Errorf("invalid signature length, expected 65 bytes, got %d bytes", len(sigBytes))
		}

		r := new(big.Int).SetBytes(sigBytes[:32])
		s := new(big.Int).SetBytes(sigBytes[32:64])

		var rBytes, sBytes [32]byte
		r.FillBytes(rBytes[:])
		s.FillBytes(sBytes[:])

		sessionSignature = RSY{
			R:       rBytes,
			S:       sBytes,
			YParity: sigBytes[64] - 27, // Adjust v from 27/28 to 0/1
		}
	}

	s.IsImplicit = false
	s.PermissionIndex = permissionIndex
	s.SessionSignature = sessionSignature

	return nil
}

// AttestationFromJSON converts JSON to an Attestation (placeholder)
func AttestationFromJSON(jsonStr string) (Attestation, error) {
	var attestation Attestation
	if err := json.Unmarshal([]byte(jsonStr), &attestation); err != nil {
		return Attestation{}, fmt.Errorf("failed to parse attestation: %w", err)
	}
	return attestation, nil
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

// packRSY packs an RSY signature into a byte array
func packRSY(rsy RSY) ([]byte, error) {
	result := make([]byte, 64)

	copy(result[0:32], rsy.R[:])

	yParityBit := (rsy.YParity & 0x01) << 7

	sByte := (rsy.S[0] & 0x7F) | yParityBit
	result[32] = sByte

	copy(result[33:64], rsy.S[1:])

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

// encodeBlacklist encodes the blacklist into bytes
func encodeBlacklist(blacklist []common.Address) ([]byte, error) {
	result := []byte{}
	for _, addr := range blacklist {
		result = append(result, addr.Bytes()...)
	}
	return result, nil
}

// encodeSessionNode encodes a session node
func encodeSessionNode(node SessionNode) ([]byte, error) {
	result := []byte{}
	result = append(result, node.Signer.Bytes()...)

	if node.Permissions != nil {
		valueLimit := node.Permissions.ValueLimit
		if valueLimit == nil {
			valueLimit = big.NewInt(0)
		}

		valueLimitBytes := make([]byte, 32)
		valueLimit.FillBytes(valueLimitBytes)
		result = append(result, valueLimitBytes...)

		deadline := node.Permissions.Deadline
		if deadline == nil {
			deadline = big.NewInt(0)
		}

		deadlineBytes := make([]byte, 32)
		deadline.FillBytes(deadlineBytes)
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
