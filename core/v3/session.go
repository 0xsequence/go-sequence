package v3

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// `SessionsTopology` represents the topology of sessions
type SessionsTopology struct {
	GlobalSigner common.Address   `json:"globalSigner"`
	Sessions     []SessionNode    `json:"sessions"`
	Blacklist    []common.Address `json:"blacklist"`
}

// `SessionNode` represents a node in the session topology
type SessionNode struct {
	Signer      common.Address      `json:"signer"`
	Permissions *SessionPermissions `json:"permissions"`
	Children    []SessionNode       `json:"children"`
}

// `SessionCallSignature` represents a call signature for a session
type SessionCallSignature struct {
	Signer    common.Address `json:"signer"`
	Signature []byte         `json:"signature"`
}

// ConfigurationTree represents the configuration tree for sessions
type ConfigurationTree struct {
	ImageHash [32]byte      `json:"imageHash"`
	Nodes     []SessionNode `json:"nodes"`
}

// `EmptySessionsTopology` creates an empty session topology with the given global signer
func EmptySessionsTopology(signer common.Address) *SessionsTopology {
	return &SessionsTopology{
		GlobalSigner: signer,
		Sessions:     []SessionNode{},
		Blacklist:    []common.Address{},
	}
}

// `SessionsTopologyToJSON` converts a session topology to JSON
func SessionsTopologyToJSON(topology *SessionsTopology) (string, error) {
	bytes, err := json.Marshal(topology)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session topology: %w", err)
	}
	return string(bytes), nil
}

// `SessionsTopologyFromJSON` converts JSON to a session topology
func SessionsTopologyFromJSON(data string) (*SessionsTopology, error) {
	var topology SessionsTopology
	if err := json.Unmarshal([]byte(data), &topology); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session topology: %w", err)
	}
	return &topology, nil
}

// `SessionCallSignatureFromJSON` converts JSON to a session call signature
func SessionCallSignatureFromJSON(data string) (*SessionCallSignature, error) {
	var sig SessionCallSignature
	if err := json.Unmarshal([]byte(data), &sig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session call signature: %w", err)
	}
	return &sig, nil
}

// `IsSessionsTopology` checks if the given interface is a valid session topology
func IsSessionsTopology(v interface{}) bool {
	_, ok := v.(*SessionsTopology)
	return ok
}

// `GetSessionPermissions` gets the permissions for a given signer in the topology
func GetSessionPermissions(topology *SessionsTopology, signer common.Address) *SessionPermissions {
	for _, session := range topology.Sessions {
		if session.Signer == signer {
			return session.Permissions
		}
	}
	return nil
}

// `MergeSessionsTopologies` merges two session topologies
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

// `BalanceSessionsTopology` balances a session topology
func BalanceSessionsTopology(topology *SessionsTopology) *SessionsTopology {
	// For now, just return the topology as is
	// In a real implementation, this would rebalance the tree structure
	return topology
}

// `SessionsTopologyToConfigurationTree` converts a session topology to a configuration tree
func SessionsTopologyToConfigurationTree(topology *SessionsTopology) *ConfigurationTree {
	return &ConfigurationTree{
		Nodes: topology.Sessions,
	}
}

// `HashConfigurationTree` computes the hash of a configuration tree
func HashConfigurationTree(tree *ConfigurationTree) [32]byte {
	// For now, return empty hash
	// In a real implementation, this would compute a proper hash
	return [32]byte{}
}

// `EncodeSessionCallSignatures` encodes session call signatures
func EncodeSessionCallSignatures(sigs []SessionCallSignature, topology *SessionsTopology, explicitSigners, implicitSigners []common.Address) ([]byte, error) {
	// For now, return empty bytes
	// In a real implementation, this would encode the signatures properly
	return []byte{}, nil
}

// `AddToImplicitBlacklist` adds an address to the implicit blacklist
func AddToImplicitBlacklist(topology *SessionsTopology, addr common.Address) *SessionsTopology {
	result := &SessionsTopology{
		GlobalSigner: topology.GlobalSigner,
		Sessions:     append([]SessionNode{}, topology.Sessions...),
		Blacklist:    append([]common.Address{}, topology.Blacklist...),
	}
	result.Blacklist = append(result.Blacklist, addr)
	return result
}

// `RemoveFromImplicitBlacklist` removes an address from the implicit blacklist
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

// `RemoveExplicitSession` removes an explicit session from the topology
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
