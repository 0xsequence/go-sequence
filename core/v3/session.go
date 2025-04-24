package v3

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	"golang.org/x/crypto/sha3"
)

// -----------------------------------------------------------------------------
// Constants
// -----------------------------------------------------------------------------

const (
	SESSIONS_FLAG_PERMISSIONS     = 0
	SESSIONS_FLAG_NODE            = 1
	SESSIONS_FLAG_BRANCH          = 2
	SESSIONS_FLAG_BLACKLIST       = 3
	SESSIONS_FLAG_IDENTITY_SIGNER = 4
)

// -----------------------------------------------------------------------------
// Topology Leaf Types
// -----------------------------------------------------------------------------

// ImplicitBlacklistLeaf is a leaf that holds a blacklist.
type ImplicitBlacklistLeaf struct {
	Type      string           `json:"type"` // must be "implicit-blacklist"
	Blacklist []common.Address `json:"blacklist"`
}

// IdentitySignerLeaf is a leaf that holds the identity signer.
type IdentitySignerLeaf struct {
	Type           string         `json:"type"` // must be "identity-signer"
	IdentitySigner common.Address `json:"identitySigner"`
}

// SessionPermissionsLeaf embeds SessionPermissions and adds a type field.
type SessionPermissionsLeaf struct {
	SessionPermissions
	Type string `json:"type"` // must be "session-permissions"
}

// -----------------------------------------------------------------------------
// SessionsTopology Type
// -----------------------------------------------------------------------------

// SessionsTopology is a union type that can be either a branch (slice), a leaf, or a node (hash bytes).
type SessionsTopology struct {
	Branch []SessionsTopology // non-nil if this is a branch
	Leaf   interface{}        // one of *ImplicitBlacklistLeaf, *IdentitySignerLeaf, *SessionPermissionsLeaf; non-nil if leaf
	Node   []byte             // non-nil if a node (hashed value)
}

// UnmarshalJSON implements custom unmarshaling.
// It accepts three JSON forms:
//   - A JSON array → branch
//   - A JSON string starting with "0x" → node
//   - A JSON object that must include a "type" field → leaf
func (st *SessionsTopology) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	// Check for array
	if len(data) > 0 && data[0] == '[' {
		var branch []SessionsTopology
		if err := json.Unmarshal(data, &branch); err != nil {
			return fmt.Errorf("failed to unmarshal branch: %w", err)
		}
		st.Branch = branch
		return nil
	}
	// Check for string (node)
	if len(data) > 0 && data[0] == '"' {
		// Unquote the string
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("failed to unmarshal node string: %w", err)
		}
		// If it starts with "0x", treat as hex bytes.
		if strings.HasPrefix(s, "0x") {
			b, err := hex.DecodeString(s[2:])
			if err != nil {
				return fmt.Errorf("failed to decode hex node: %w", err)
			}
			st.Node = b
			return nil
		}
	}
	// Otherwise, assume object (leaf). We need to inspect the "type" field.
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return fmt.Errorf("failed to unmarshal leaf object: %w", err)
	}
	// The object must have a "type" key.
	var typ string
	if rawType, ok := m["type"]; ok {
		if err := json.Unmarshal(rawType, &typ); err != nil {
			return fmt.Errorf("failed to unmarshal type field: %w", err)
		}
	} else {
		return fmt.Errorf("leaf object missing type field")
	}
	switch typ {
	case "implicit-blacklist":
		var leaf ImplicitBlacklistLeaf
		if err := json.Unmarshal(data, &leaf); err != nil {
			return fmt.Errorf("failed to unmarshal implicit-blacklist leaf: %w", err)
		}
		st.Leaf = &leaf
	case "identity-signer":
		var leaf IdentitySignerLeaf
		if err := json.Unmarshal(data, &leaf); err != nil {
			return fmt.Errorf("failed to unmarshal identity-signer leaf: %w", err)
		}
		st.Leaf = &leaf
	case "session-permissions":
		var sp SessionPermissions
		if err := json.Unmarshal(data, &sp); err != nil {
			return fmt.Errorf("failed to unmarshal session permissions: %w", err)
		}
		leaf := SessionPermissionsLeaf{
			SessionPermissions: sp,
			Type:               "session-permissions",
		}
		st.Leaf = &leaf
	default:
		return fmt.Errorf("unknown leaf type: %s", typ)
	}
	return nil
}

// MarshalJSON implements custom marshaling.
func (st SessionsTopology) MarshalJSON() ([]byte, error) {
	encoded, err := SessionsTopologyToJSON(st)
	if err != nil {
		return nil, err
	}
	// Convert the JSON string to a []bytes
	return []byte(encoded), nil
}

// -----------------------------------------------------------------------------
// Type Guard Functions
// -----------------------------------------------------------------------------

func IsSessionsNode(st SessionsTopology) bool {
	return st.Node != nil
}

func IsSessionsBranch(st SessionsTopology) bool {
	return st.Branch != nil
}

func IsSessionsLeaf(st SessionsTopology) bool {
	return st.Leaf != nil
}

func IsImplicitBlacklistLeaf(st SessionsTopology) bool {
	if !IsSessionsLeaf(st) {
		return false
	}
	_, ok := st.Leaf.(*ImplicitBlacklistLeaf)
	return ok
}

func IsIdentitySignerLeaf(st SessionsTopology) bool {
	if !IsSessionsLeaf(st) {
		return false
	}
	_, ok := st.Leaf.(*IdentitySignerLeaf)
	return ok
}

func IsSessionPermissionsLeaf(st SessionsTopology) bool {
	if !IsSessionsLeaf(st) {
		return false
	}
	_, ok := st.Leaf.(*SessionPermissionsLeaf)
	return ok
}

// IsCompleteSessionsTopology checks that there is exactly one identity signer leaf and one implicit-blacklist leaf.
func IsCompleteSessionsTopology(topology SessionsTopology) bool {
	ids, blacks := checkIsCompleteSessionsBranch(topology)
	return ids == 1 && blacks == 1
}

func checkIsCompleteSessionsBranch(topology SessionsTopology) (identitySignerCount, blacklistCount int) {
	if IsSessionsBranch(topology) {
		for _, child := range topology.Branch {
			i, b := checkIsCompleteSessionsBranch(child)
			identitySignerCount += i
			blacklistCount += b
		}
	}
	if IsIdentitySignerLeaf(topology) {
		identitySignerCount++
	}
	if IsImplicitBlacklistLeaf(topology) {
		blacklistCount++
	}
	return
}

// getIdentitySigner returns the identity signer address from the topology, or nil if not found.
func getIdentitySigner(topology SessionsTopology) *common.Address {
	if IsIdentitySignerLeaf(topology) {
		leaf := topology.Leaf.(*IdentitySignerLeaf)
		return &leaf.IdentitySigner
	}
	if IsSessionsBranch(topology) {
		var results []common.Address
		for _, child := range topology.Branch {
			if addr := getIdentitySigner(child); addr != nil {
				results = append(results, *addr)
			}
		}
		if len(results) > 1 {
			panic("multiple identity signers found")
		}
		if len(results) == 1 {
			return &results[0]
		}
	}
	return nil
}

// getImplicitBlacklist returns the blacklist array from the topology or nil.
func getImplicitBlacklist(topology SessionsTopology) []common.Address {
	leaf := getImplicitBlacklistLeaf(topology)
	if leaf == nil {
		return nil
	}
	return leaf.Blacklist
}

// getImplicitBlacklistLeaf returns the implicit-blacklist leaf from the topology or nil.
func getImplicitBlacklistLeaf(topology SessionsTopology) *ImplicitBlacklistLeaf {
	if IsImplicitBlacklistLeaf(topology) {
		return topology.Leaf.(*ImplicitBlacklistLeaf)
	}
	if IsSessionsBranch(topology) {
		var results []*ImplicitBlacklistLeaf
		for _, child := range topology.Branch {
			if leaf := getImplicitBlacklistLeaf(child); leaf != nil {
				results = append(results, leaf)
			}
		}
		if len(results) > 1 {
			panic("multiple implicit blacklist leaves found")
		}
		if len(results) == 1 {
			return results[0]
		}
	}
	return nil
}

// getSessionPermissions searches for an explicit session permission leaf with the given signer.
func GetSessionPermissions(topology SessionsTopology, addr common.Address) *SessionPermissions {
	if IsSessionPermissionsLeaf(topology) {
		leaf := topology.Leaf.(*SessionPermissionsLeaf)
		if leaf.SessionPermissions.Signer == addr {
			return &leaf.SessionPermissions
		}
	}
	if IsSessionsBranch(topology) {
		for _, child := range topology.Branch {
			if sp := GetSessionPermissions(child, addr); sp != nil {
				return sp
			}
		}
	}
	return nil
}

// GetExplicitSigners returns all explicit session permission signers from the topology.
func GetExplicitSigners(topology SessionsTopology) []common.Address {
	var current []common.Address
	if IsSessionPermissionsLeaf(topology) {
		leaf := topology.Leaf.(*SessionPermissionsLeaf)
		current = append(current, leaf.SessionPermissions.Signer)
	}
	if IsSessionsBranch(topology) {
		for _, child := range topology.Branch {
			current = append(current, GetExplicitSigners(child)...)
		}
	}
	return current
}

// -----------------------------------------------------------------------------
// Encoding / Decoding Functions for Contract Validation and JSON
// -----------------------------------------------------------------------------

func SessionsTopologyToGenericTree(topology SessionsTopology) (Tree, error) {
	if IsSessionsBranch(topology) {
		branch := make(TreeBranch, len(topology.Branch))
		for i, child := range topology.Branch {
			leaf, err := SessionsTopologyToGenericTree(child)
			if err != nil {
				return nil, err
			}
			branch[i] = leaf
		}
		return branch, nil
	}
	if IsSessionsNode(topology) {
		return core.ImageHash{Hash: common.Hash(PadLeft(topology.Node, 32))}, nil
	}
	return encodeLeafToGenericLeaf(topology)
}

// encodeLeafToGeneric encodes a session leaf to a GenericTree.Leaf.
func encodeLeafToGenericLeaf(leaf SessionsTopology) (TreeLeaf, error) {
	if IsSessionPermissionsLeaf(leaf) {
		// Call the permission package encoder.
		encoded, err := EncodeSessionPermissions(&leaf.Leaf.(*SessionPermissionsLeaf).SessionPermissions)
		if err != nil {
			return TreeLeaf{}, err
		}
		return TreeLeaf(Concat([][]byte{FromNumber(SESSIONS_FLAG_PERMISSIONS), encoded})), nil
	}
	if IsImplicitBlacklistLeaf(leaf) {
		ibl := leaf.Leaf.(*ImplicitBlacklistLeaf)
		var parts [][]byte
		parts = append(parts, FromNumber(SESSIONS_FLAG_BLACKLIST))
		for _, addr := range ibl.Blacklist {
			parts = append(parts, PadLeft(addr.Bytes(), 20))
		}
		return TreeLeaf(Concat(parts)), nil
	}
	if IsIdentitySignerLeaf(leaf) {
		isLeaf := leaf.Leaf.(*IdentitySignerLeaf)
		return TreeLeaf(Concat([][]byte{FromNumber(SESSIONS_FLAG_IDENTITY_SIGNER), PadLeft(isLeaf.IdentitySigner.Bytes(), 20)})), nil
	}
	return TreeLeaf{}, fmt.Errorf("invalid leaf")
}

// EncodeSessionCallSignatures encodes the call signatures for contract validation.
func EncodeSessionCallSignatures(callSignatures []SessionCallSignature, topology SessionsTopology, explicitSigners, implicitSigners []common.Address) ([]byte, error) {
	// Create a buffer to write to
	var buf bytes.Buffer

	// Validate the topology
	if !IsCompleteSessionsTopology(topology) {
		return nil, fmt.Errorf("incomplete topology")
	}

	// Optimize the topology
	optimizedTopology, err := MinimiseSessionsTopology(topology, explicitSigners, implicitSigners)
	if err != nil {
		return nil, err
	}

	// Encode topology
	encodedTopology, err := EncodeSessionsTopology(optimizedTopology)
	if err != nil {
		return nil, err
	}
	if minBytesFor(uint64(len(encodedTopology))) > 3 {
		return nil, fmt.Errorf("session topology is too large")
	}

	// Write topology length and data
	if err := writeUint24(&buf, uint32(len(encodedTopology))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(encodedTopology); err != nil {
		return nil, err
	}

	// Build attestation map and encoded attestations
	attestationMap := make(map[string]int)
	var encodedAttestations [][]byte

	// Map implicit signatures to attestation indices
	for _, cs := range callSignatures {
		if isImplicitSessionCallSignature(cs) {
			sig := cs.(ImplicitSessionCallSignature)
			attJSON, err := json.Marshal(sig.Attestation)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal attestation: %w", err)
			}
			if _, ok := attestationMap[string(attJSON)]; !ok {
				attestationMap[string(attJSON)] = len(encodedAttestations)
				// Concatenate attestation encoding and packed identity signature
				attEncoded := sig.Attestation.Encode()
				encodedAttestations = append(encodedAttestations, ConcatBytes(attEncoded, packRSY(sig.IdentitySignature)))
			}
		}
	}

	// Write attestation count and data
	if minBytesFor(uint64(len(encodedAttestations))) > 1 {
		return nil, fmt.Errorf("too many attestations")
	}
	if err := writeUint8(&buf, uint8(len(encodedAttestations))); err != nil {
		return nil, err
	}
	for _, att := range encodedAttestations {
		if _, err := buf.Write(att); err != nil {
			return nil, err
		}
	}

	// Write call signatures
	for _, cs := range callSignatures {
		if isImplicitSessionCallSignature(cs) {
			sig := cs.(ImplicitSessionCallSignature)
			attJSON, err := json.Marshal(sig.Attestation)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal attestation: %w", err)
			}
			attIndex, ok := attestationMap[string(attJSON)]
			if !ok {
				return nil, fmt.Errorf("failed to find attestation index")
			}
			// Set flag: highest bit 1 for implicit, lower 7 bits = attestation index
			flag := byte(0x80) | byte(attIndex)
			if err := writeUint8(&buf, flag); err != nil {
				return nil, err
			}
			if _, err := buf.Write(packRSY(sig.SessionSignature)); err != nil {
				return nil, err
			}
		} else if isExplicitSessionCallSignature(cs) {
			sig := cs.(ExplicitSessionCallSignature)
			if sig.PermissionIndex.Cmp(big.NewInt(MAX_PERMISSIONS_COUNT)) > 0 {
				return nil, fmt.Errorf("permission index is too large")
			}
			if err := writeUint8(&buf, uint8(sig.PermissionIndex.Uint64())); err != nil {
				return nil, err
			}
			if _, err := buf.Write(packRSY(sig.SessionSignature)); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid call signature")
		}
	}

	return buf.Bytes(), nil
}

// EncodeSessionsTopology encodes the topology for contract validation.
func EncodeSessionsTopology(topology SessionsTopology) ([]byte, error) {
	var buf bytes.Buffer

	// Branch: if this is a branch then encode each child and then add a branch header.
	if IsSessionsBranch(topology) {
		var encodedChildren [][]byte
		for _, child := range topology.Branch {
			enc, err := EncodeSessionsTopology(child)
			if err != nil {
				return nil, err
			}
			encodedChildren = append(encodedChildren, enc)
		}
		encoded := Concat(encodedChildren)
		encodedSize := minBytesFor(uint64(len(encoded)))
		if encodedSize > 15 {
			return nil, fmt.Errorf("branch too large")
		}
		flagByte := byte((SESSIONS_FLAG_BRANCH << 4) | encodedSize)
		if err := writeUint8(&buf, flagByte); err != nil {
			return nil, err
		}
		if err := writeUintX(&buf, uint64(len(encoded)), encodedSize); err != nil {
			return nil, err
		}
		if _, err := buf.Write(encoded); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	// Leaf: determine which leaf type.
	if IsSessionPermissionsLeaf(topology) {
		flagByte := byte(SESSIONS_FLAG_PERMISSIONS << 4)
		if err := writeUint8(&buf, flagByte); err != nil {
			return nil, err
		}
		encodedLeaf, err := EncodeSessionPermissions(&topology.Leaf.(*SessionPermissionsLeaf).SessionPermissions)
		if err != nil {
			return nil, err
		}
		if _, err := buf.Write(encodedLeaf); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	if IsSessionsNode(topology) {
		flagByte := byte(SESSIONS_FLAG_NODE << 4)
		if err := writeUint8(&buf, flagByte); err != nil {
			return nil, err
		}
		if _, err := buf.Write(topology.Node); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	if IsImplicitBlacklistLeaf(topology) {
		ibl := topology.Leaf.(*ImplicitBlacklistLeaf)
		if len(ibl.Blacklist) >= 0x0f {
			if len(ibl.Blacklist) > 0xffff {
				return nil, fmt.Errorf("blacklist too large")
			}
			if err := writeUint8(&buf, byte((SESSIONS_FLAG_BLACKLIST<<4)|0x0f)); err != nil {
				return nil, err
			}
			if err := writeUint16(&buf, uint16(len(ibl.Blacklist))); err != nil {
				return nil, err
			}
		} else {
			if err := writeUintX(&buf, uint64((SESSIONS_FLAG_BLACKLIST<<4)|uint8(len(ibl.Blacklist))), 1); err != nil {
				return nil, err
			}
		}
		for _, addr := range ibl.Blacklist {
			if _, err := buf.Write(PadLeft(addr.Bytes(), 20)); err != nil {
				return nil, err
			}
		}
		return buf.Bytes(), nil
	}

	if IsIdentitySignerLeaf(topology) {
		flagByte := byte(SESSIONS_FLAG_IDENTITY_SIGNER << 4)
		if err := writeUint8(&buf, flagByte); err != nil {
			return nil, err
		}
		leaf := topology.Leaf.(*IdentitySignerLeaf)
		if _, err := buf.Write(PadLeft(leaf.IdentitySigner.Bytes(), 20)); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("invalid topology")
}

// sessionsTopologyToJSON returns the JSON representation of the topology.
// For node types, it returns a hex string; for session-permissions leaves, it calls the permission package encoder for JSON; for other leaves it returns the object as is;
// for branches, it maps recursively.
func SessionsTopologyToJSON(topology SessionsTopology) (string, error) {
	enc, err := encodeSessionsTopologyForJson(topology)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(enc)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func encodeSessionsTopologyForJson(topology SessionsTopology) (any, error) {
	if IsSessionsBranch(topology) {
		var result []any
		for _, child := range topology.Branch {
			enc, err := encodeSessionsTopologyForJson(child)
			if err != nil {
				return nil, err
			}
			result = append(result, enc)
		}
		return result, nil
	}
	if IsSessionsNode(topology) {
		return "0x" + hex.EncodeToString(topology.Node), nil
	}
	if IsSessionPermissionsLeaf(topology) {
		// Use the permission package JSON encoder.
		encoded := encodeSessionPermissionsForJson(&topology.Leaf.(*SessionPermissionsLeaf).SessionPermissions)
		// Add the type field.
		encoded["type"] = "session-permissions"
		return encoded, nil
	}
	if IsImplicitBlacklistLeaf(topology) || IsIdentitySignerLeaf(topology) {
		return topology.Leaf, nil
	}
	return nil, fmt.Errorf("invalid topology")
}

// sessionsTopologyFromJSON parses the JSON string into a SessionsTopology.
func SessionsTopologyFromJSON(data string) (SessionsTopology, error) {
	var parsed any
	if err := json.Unmarshal([]byte(data), &parsed); err != nil {
		return SessionsTopology{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return sessionsTopologyFromParsed(parsed)
}

func sessionsTopologyFromParsed(parsed any) (SessionsTopology, error) {
	// If parsed is an array, treat as branch.
	switch v := parsed.(type) {
	case []any:
		var branch []SessionsTopology
		for _, item := range v {
			child, err := sessionsTopologyFromParsed(item)
			if err != nil {
				return SessionsTopology{}, err
			}
			branch = append(branch, child)
		}
		return SessionsTopology{Branch: branch}, nil
	case string:
		if strings.HasPrefix(v, "0x") {
			b, err := hex.DecodeString(v[2:])
			if err != nil {
				return SessionsTopology{}, fmt.Errorf("invalid hex in node: %w", err)
			}
			return SessionsTopology{Node: b}, nil
		}
		return SessionsTopology{}, fmt.Errorf("invalid string value for topology")
	case map[string]any:
		// Check if object has "type" field.
		typVal, ok := v["type"]
		if !ok {
			return SessionsTopology{}, fmt.Errorf("object missing type field")
		}
		typ, ok := typVal.(string)
		if !ok {
			return SessionsTopology{}, fmt.Errorf("type field must be a string")
		}
		switch typ {
		case "session-permissions":
			// Unmarshal into SessionPermissions then wrap
			b, err := json.Marshal(v)
			if err != nil {
				return SessionsTopology{}, err
			}
			var sp SessionPermissions
			if err := json.Unmarshal(b, &sp); err != nil {
				return SessionsTopology{}, fmt.Errorf("failed to unmarshal session permissions: %w", err)
			}
			leaf := SessionPermissionsLeaf{
				SessionPermissions: sp,
				Type:               "session-permissions",
			}
			return SessionsTopology{Leaf: &leaf}, nil
		case "identity-signer":
			b, err := json.Marshal(v)
			if err != nil {
				return SessionsTopology{}, err
			}
			var leaf IdentitySignerLeaf
			if err := json.Unmarshal(b, &leaf); err != nil {
				return SessionsTopology{}, fmt.Errorf("failed to unmarshal identity signer leaf: %w", err)
			}
			return SessionsTopology{Leaf: &leaf}, nil
		case "implicit-blacklist":
			b, err := json.Marshal(v)
			if err != nil {
				return SessionsTopology{}, err
			}
			var leaf ImplicitBlacklistLeaf
			if err := json.Unmarshal(b, &leaf); err != nil {
				return SessionsTopology{}, fmt.Errorf("failed to unmarshal implicit blacklist leaf: %w", err)
			}
			return SessionsTopology{Leaf: &leaf}, nil
		default:
			return SessionsTopology{}, fmt.Errorf("unknown leaf type: %s", typ)
		}
	default:
		return SessionsTopology{}, fmt.Errorf("invalid topology type")
	}
}

// -----------------------------------------------------------------------------
// Topology Operations
// -----------------------------------------------------------------------------

// removeExplicitSession removes any explicit session-permissions leaf that has a signer matching the given address.
// If the removal causes a branch to become empty, it returns nil; if exactly one child remains, it collapses it.
func RemoveExplicitSession(topology SessionsTopology, signerAddr string) (SessionsTopology, error) {
	target := common.HexToAddress(signerAddr)
	if IsSessionPermissionsLeaf(topology) {
		leaf := topology.Leaf.(*SessionPermissionsLeaf)
		if leaf.SessionPermissions.Signer == target {
			return SessionsTopology{}, nil
		}
		return topology, nil
	}
	if IsSessionsBranch(topology) {
		var newChildren []SessionsTopology
		for _, child := range topology.Branch {
			updated, err := RemoveExplicitSession(child, signerAddr)
			if err != nil {
				return SessionsTopology{}, err
			}
			// Skip nil children
			if !IsEmptyTopology(updated) {
				newChildren = append(newChildren, updated)
			}
		}
		if len(newChildren) == 0 {
			return SessionsTopology{}, nil
		}
		if len(newChildren) == 1 {
			return newChildren[0], nil
		}
		return SessionsTopology{Branch: newChildren}, nil
	}
	// For nodes and other leaves, return unchanged.
	return topology, nil
}

func IsEmptyTopology(st SessionsTopology) bool {
	return !IsSessionsBranch(st) && !IsSessionsLeaf(st) && !IsSessionsNode(st)
}

// addExplicitSession adds a new session-permissions leaf.
// If a session for the given signer already exists, it returns an error.
// Otherwise it merges the new leaf into the topology and then balances it.
func AddExplicitSession(topology SessionsTopology, sp SessionPermissions) (SessionsTopology, error) {
	if GetSessionPermissions(topology, sp.Signer) != nil {
		return SessionsTopology{}, fmt.Errorf("session already exists")
	}
	newLeaf := SessionsTopology{
		Leaf: &SessionPermissionsLeaf{
			SessionPermissions: sp,
			Type:               "session-permissions",
		},
	}
	merged := MergeSessionsTopologies(topology, newLeaf)
	return balanceSessionsTopology(merged), nil
}

// MergeSessionsTopologies merges two topologies into a branch.
func MergeSessionsTopologies(a, b SessionsTopology) SessionsTopology {
	return SessionsTopology{Branch: []SessionsTopology{a, b}}
}

// flattenSessionsTopology returns a flat slice of all leaves and nodes in the topology.
func flattenSessionsTopology(topology SessionsTopology) []SessionsTopology {
	if IsSessionsLeaf(topology) || IsSessionsNode(topology) {
		return []SessionsTopology{topology}
	}
	if IsSessionsBranch(topology) {
		var result []SessionsTopology
		for _, child := range topology.Branch {
			result = append(result, flattenSessionsTopology(child)...)
		}
		return result
	}
	return nil
}

// buildBalancedSessionsTopology builds a balanced branch from a slice of topologies.
func buildBalancedSessionsTopology(items []SessionsTopology) SessionsTopology {
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 0 {
		panic("cannot build topology from empty list")
	}
	mid := len(items) / 2
	left := buildBalancedSessionsTopology(items[:mid])
	right := buildBalancedSessionsTopology(items[mid:])
	return SessionsTopology{Branch: []SessionsTopology{left, right}}
}

// balanceSessionsTopology flattens the topology and rebuilds it as a balanced branch,
// ensuring that the implicit blacklist and identity signer leaves remain present.
func balanceSessionsTopology(topology SessionsTopology) SessionsTopology {
	flattened := flattenSessionsTopology(topology)
	var ibl, idSig, others []SessionsTopology
	for _, t := range flattened {
		if IsImplicitBlacklistLeaf(t) {
			ibl = append(ibl, t)
		} else if IsIdentitySignerLeaf(t) {
			idSig = append(idSig, t)
		} else {
			others = append(others, t)
		}
	}
	if len(ibl) != 1 || len(idSig) != 1 {
		return topology
	}
	all := append([]SessionsTopology{}, ibl[0], idSig[0])
	all = append(all, others...)
	return buildBalancedSessionsTopology(all)
}

// cleanSessionsTopology removes session-permissions leaves whose deadline has expired.
// It recursively cleans branches; if a branch becomes empty it returns an empty topology,
// and if a branch has a single child, it collapses it upward.
func cleanSessionsTopology(topology SessionsTopology, currentTime *big.Int) (SessionsTopology, error) {
	if IsSessionsNode(topology) {
		return topology, nil
	}
	if IsSessionPermissionsLeaf(topology) {
		leaf := topology.Leaf.(*SessionPermissionsLeaf)
		if leaf.Deadline.Cmp(currentTime) < 0 {
			return SessionsTopology{}, nil
		}
		return topology, nil
	}
	if IsIdentitySignerLeaf(topology) || IsImplicitBlacklistLeaf(topology) {
		return topology, nil
	}
	if IsSessionsBranch(topology) {
		var newChildren []SessionsTopology
		for _, child := range topology.Branch {
			cleaned, err := cleanSessionsTopology(child, currentTime)
			if err != nil {
				return SessionsTopology{}, err
			}
			if !IsEmptyTopology(cleaned) {
				newChildren = append(newChildren, cleaned)
			}
		}
		if len(newChildren) == 0 {
			return SessionsTopology{}, nil
		}
		if len(newChildren) == 1 {
			return newChildren[0], nil
		}
		return SessionsTopology{Branch: newChildren}, nil
	}
	return SessionsTopology{}, fmt.Errorf("invalid topology in cleaning")
}

// minimiseSessionsTopology rolls up unused signers into a hashed node.
// If a session-permissions leaf is for a signer in explicitSigners, it is not rolled up.
func MinimiseSessionsTopology(topology SessionsTopology, explicitSigners []common.Address, implicitSigners []common.Address) (SessionsTopology, error) {
	if IsSessionsBranch(topology) {
		var branches []SessionsTopology
		for _, child := range topology.Branch {
			encodedBranch, err := MinimiseSessionsTopology(child, explicitSigners, implicitSigners)
			if err != nil {
				// Do not re-encode the error when recursing
				return SessionsTopology{}, err
			}
			branches = append(branches, encodedBranch)
		}
		// If all branches are nodes, roll them up
		allNodes := true
		for _, b := range branches {
			if !IsSessionsNode(b) {
				allNodes = false
				break
			}
		}
		if allNodes {
			nodeBranch := make(TreeBranch, len(branches))
			for i, b := range branches {
				nodeBranch[i] = core.ImageHash{Hash: common.Hash(PadLeft(b.Node, 32))}
			}
			return SessionsTopology{Node: nodeBranch.ImageHash().Bytes()}, nil
		}
		return SessionsTopology{Branch: branches}, nil
	}
	if IsSessionPermissionsLeaf(topology) {
		leaf := topology.Leaf.(*SessionPermissionsLeaf)
		for _, addr := range explicitSigners {
			if leaf.Signer == addr {
				// Used, return unchanged
				return topology, nil
			}
		}
		// Unused. Optimise by hashing this node
		encoded, err := encodeLeafToGenericLeaf(topology)
		if err != nil {
			return SessionsTopology{}, fmt.Errorf("failed to encode leaf: %w", err)
		}
		return SessionsTopology{Node: encoded.ImageHash().Bytes()}, nil
	}
	if IsImplicitBlacklistLeaf(topology) {
		if len(implicitSigners) == 0 {
			// Unused. Optimise by hashing this node
			encoded, err := encodeLeafToGenericLeaf(topology)
			if err != nil {
				return SessionsTopology{}, fmt.Errorf("failed to encode leaf: %w", err)
			}
			return SessionsTopology{Node: encoded.ImageHash().Bytes()}, nil
		}
		return topology, nil
	}
	if IsIdentitySignerLeaf(topology) || IsSessionsNode(topology) {
		// Return unchanged
		return topology, nil
	}
	return SessionsTopology{}, fmt.Errorf("invalid topology in minimisation")
}

// addToImplicitBlacklist adds an address to the blacklist in the implicit-blacklist leaf.
func AddToImplicitBlacklist(topology SessionsTopology, addr common.Address) (SessionsTopology, error) {
	ibl := getImplicitBlacklistLeaf(topology)
	if ibl == nil {
		return SessionsTopology{}, fmt.Errorf("no blacklist found")
	}
	// Check if already present.
	for _, a := range ibl.Blacklist {
		if a == addr {
			return topology, nil
		}
	}
	ibl.Blacklist = append(ibl.Blacklist, addr)
	// Sort to ensure deterministic order.
	sort.Slice(ibl.Blacklist, func(i, j int) bool {
		return strings.ToLower(ibl.Blacklist[i].Hex()) < strings.ToLower(ibl.Blacklist[j].Hex())
	})
	return topology, nil
}

// removeFromImplicitBlacklist removes an address from the implicit blacklist.
func RemoveFromImplicitBlacklist(topology SessionsTopology, addr common.Address) (SessionsTopology, error) {
	ibl := getImplicitBlacklistLeaf(topology)
	if ibl == nil {
		return SessionsTopology{}, fmt.Errorf("no blacklist found")
	}
	var newList []common.Address
	for _, a := range ibl.Blacklist {
		if a != addr {
			newList = append(newList, a)
		}
	}
	ibl.Blacklist = newList
	return topology, nil
}

// EmptySessionsTopology generates an empty topology given the identity signer.
func EmptySessionsTopology(identitySigner common.Address) SessionsTopology {
	return SessionsTopology{
		Branch: []SessionsTopology{
			{Leaf: &IdentitySignerLeaf{Type: "identity-signer", IdentitySigner: identitySigner}},
			{Leaf: &ImplicitBlacklistLeaf{Type: "implicit-blacklist", Blacklist: []common.Address{}}},
		},
	}
}

// -----------------------------------------------------------------------------
// Signature encoding / decoding
// -----------------------------------------------------------------------------

// RSY represents a signature in r, s, yParity format.
type RSY struct {
	R       []byte
	S       []byte
	YParity uint8 // 0 or 1
}

// packRSY converts an RSY into its 64-byte representation using ERC-2098
func packRSY(sig RSY) []byte {
	var buf bytes.Buffer
	r := PadLeft(sig.R, 32)
	buf.Write(r)

	s := PadLeft(sig.S, 32)
	if sig.YParity == 1 {
		s[0] |= 0x80 // Set high bit if yParity is 1
	}
	buf.Write(s)

	return buf.Bytes()
}

// --- Session Call Signature types and helper methods ---

// ImplicitSessionCallSignature represents a call signature that carries an attestation, an identity signature and a session signature.
type ImplicitSessionCallSignature struct {
	Attestation       Attestation
	IdentitySignature RSY
	SessionSignature  RSY
}

func (ImplicitSessionCallSignature) isSessionCallSignature() {}

// ExplicitSessionCallSignature represents a call signature that carries a permission index and a session signature.
type ExplicitSessionCallSignature struct {
	PermissionIndex  *big.Int // expected to be within byte-range
	SessionSignature RSY
}

func (ExplicitSessionCallSignature) isSessionCallSignature() {}

// SessionCallSignature is a union interface for both kinds.
type SessionCallSignature interface {
	isSessionCallSignature()
}

// isImplicitSessionCallSignature checks if the interface holds an implicit signature.
func isImplicitSessionCallSignature(sig SessionCallSignature) bool {
	_, ok := sig.(ImplicitSessionCallSignature)
	return ok
}

// isExplicitSessionCallSignature checks if the interface holds an explicit signature.
func isExplicitSessionCallSignature(sig SessionCallSignature) bool {
	_, ok := sig.(ExplicitSessionCallSignature)
	return ok
}

// --- Hashing call with replay protection ---
//
// hashCallWithReplayProtection computes the hash of a call with chain id, space and nonce.
// It concatenates 32-byte representations of chainId, space and nonce with the call hash,
// then computes the Keccak256 hash, returning it as a hex string.
func hashCallWithReplayProtection(call PayloadCall, chainId, space, nonce *big.Int) (string, error) {
	chainIdB := intToBytesBig(chainId, 32)
	spaceB := intToBytesBig(space, 32)
	nonceB := intToBytesBig(nonce, 32)
	callHash := call.HashCall() // Assume call.HashCall() returns []byte in hex format? Adjust as needed.
	data := ConcatBytes(chainIdB, spaceB, nonceB, callHash)
	hash := keccak256(data)
	return "0x" + hex.EncodeToString(hash), nil
}

// --- Helper functions ---

// ConcatBytes concatenates multiple byte slices.
func ConcatBytes(slices ...[]byte) []byte {
	var total int
	for _, s := range slices {
		total += len(s)
	}
	buf := make([]byte, 0, total)
	for _, s := range slices {
		buf = append(buf, s...)
	}
	return buf
}

// intToBytes converts an integer to a big-endian byte slice of exactly the given size.
func intToBytes(n int, size int) []byte {
	b := make([]byte, size)
	for i := size - 1; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n >>= 8
	}
	return b
}

// intToBytesBig converts a *big.Int to a byte slice of a given size.
// It left-pads the number with zeros.
func intToBytesBig(n *big.Int, size int) []byte {
	b := n.Bytes()
	if len(b) > size {
		return b[len(b)-size:]
	}
	padded := make([]byte, size)
	copy(padded[size-len(b):], b)
	return padded
}

// keccak256 computes the Keccak256 hash of the given data.
func keccak256(data []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	return h.Sum(nil)
}

// --- PayloadCall interface ---
//
// For hashing calls we assume a minimal interface.
type PayloadCall interface {
	// HashCall returns a []byte representing the call data to be hashed.
	HashCall() []byte
}

// -----------------------------------------------------------------------------
// Helper Functions
// -----------------------------------------------------------------------------

// Concat concatenates a slice of byte slices.
func Concat(slices [][]byte) []byte {
	return bytes.Join(slices, nil)
}

// FromNumber converts a uint64 to a 1-byte slice (assumes value fits in 1 byte).
func FromNumber(n uint64) []byte {
	return []byte{byte(n)}
}

// PadLeft pads a byte slice on the left to the given length.
func PadLeft(b []byte, length int) []byte {
	if len(b) >= length {
		return b
	}
	pad := make([]byte, length-len(b))
	return append(pad, b...)
}

// PadLeftNumber converts a number to a big-endian byte slice of fixed size.
func PadLeftNumber(n uint64, size int) []byte {
	b := make([]byte, size)
	for i := size - 1; i >= 0; i-- {
		b[i] = byte(n & 0xff)
		n >>= 8
	}
	return b
}

// ConcatMap applies f to each element of ts and concatenates the results.
func ConcatMap(ts []SessionsTopology, f func(SessionsTopology) []byte) []byte {
	var parts [][]byte
	for _, t := range ts {
		parts = append(parts, f(t))
	}
	return Concat(parts)
}
