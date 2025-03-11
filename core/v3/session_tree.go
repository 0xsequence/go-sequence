package v3

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
)

// Constants for session node types
const (
	SESSIONS_FLAG_PERMISSIONS     = 0
	SESSIONS_FLAG_NODE            = 1
	SESSIONS_FLAG_BRANCH          = 2
	SESSIONS_FLAG_BLACKLIST       = 3
	SESSIONS_FLAG_IDENTITY_SIGNER = 4
)

// Convert a SessionsTopology to a generic Tree representation
func SessionsTopologyToTree(topology *SessionsTopology) Tree {
	if topology == nil {
		return nil
	}

	branch := make(Branch, 0)

	identitySignerLeaf := createIdentitySignerLeaf(topology.GlobalSigner)
	branch = append(branch, identitySignerLeaf)

	blacklistLeaf := createBlacklistLeaf(topology.Blacklist)
	branch = append(branch, blacklistLeaf)

	for _, session := range topology.Sessions {
		sessionTree := sessionNodeToTree(&session)
		branch = append(branch, sessionTree)
	}

	return branch
}

// Convert a generic Tree back to a SessionsTopology
func TreeToSessionsTopology(tree Tree) (*SessionsTopology, error) {
	if !IsBranch(tree) {
		return nil, fmt.Errorf("tree must be a branch")
	}

	branch := tree.(Branch)
	if len(branch) < 2 {
		return nil, fmt.Errorf("tree branch must have at least identity signer and blacklist")
	}

	topology := &SessionsTopology{
		Sessions:  []SessionNode{},
		Blacklist: []common.Address{},
	}

	for _, elem := range branch {
		if IsLeaf(elem) {
			leaf := elem.(Leaf)
			if len(leaf.Value) == 0 {
				continue
			}

			flag := leaf.Value[0]

			if flag == SESSIONS_FLAG_IDENTITY_SIGNER && len(leaf.Value) >= 21 {
				topology.GlobalSigner = common.BytesToAddress(leaf.Value[1:21])
			} else if flag == SESSIONS_FLAG_BLACKLIST {
				topology.Blacklist = decodeBlacklist(leaf.Value[1:])
			}
		} else if IsBranch(elem) {
			sessionNode, err := treeToSessionNode(elem)
			if err == nil {
				topology.Sessions = append(topology.Sessions, *sessionNode)
			}
		}
	}

	return topology, nil
}

func createIdentitySignerLeaf(signer common.Address) Leaf {
	value := make([]byte, 21)
	value[0] = SESSIONS_FLAG_IDENTITY_SIGNER
	copy(value[1:], signer.Bytes())

	return Leaf{
		Type:  "leaf",
		Value: value,
	}
}

func createBlacklistLeaf(blacklist []common.Address) Leaf {
	value := make([]byte, 1+len(blacklist)*20)
	value[0] = SESSIONS_FLAG_BLACKLIST

	for i, addr := range blacklist {
		copy(value[1+i*20:], addr.Bytes())
	}

	return Leaf{
		Type:  "leaf",
		Value: value,
	}
}

func decodeBlacklist(data []byte) []common.Address {
	if len(data)%20 != 0 {
		return []common.Address{}
	}

	addressCount := len(data) / 20
	result := make([]common.Address, addressCount)

	for i := 0; i < addressCount; i++ {
		result[i] = common.BytesToAddress(data[i*20 : (i+1)*20])
	}

	return result
}

func sessionNodeToTree(node *SessionNode) Tree {
	if node == nil {
		return nil
	}

	branch := make(Branch, 0)

	signerLeaf := createSignerLeaf(node.Signer)
	branch = append(branch, signerLeaf)

	if node.Permissions != nil {
		permissionsLeaf := createPermissionsLeaf(node.Permissions)
		branch = append(branch, permissionsLeaf)
	}

	for _, child := range node.Children {
		childCopy := child
		childTree := sessionNodeToTree(&childCopy)
		branch = append(branch, childTree)
	}

	return branch
}

func createSignerLeaf(signer common.Address) Leaf {
	value := make([]byte, 21)
	value[0] = SESSIONS_FLAG_PERMISSIONS
	copy(value[1:], signer.Bytes())

	return Leaf{
		Type:  "leaf",
		Value: value,
	}
}

func createPermissionsLeaf(permissions *SessionPermissions) Leaf {
	// TODO: Implement
	value := []byte{SESSIONS_FLAG_PERMISSIONS}

	return Leaf{
		Type:  "leaf",
		Value: value,
	}
}

func treeToSessionNode(tree Tree) (*SessionNode, error) {
	if !IsBranch(tree) {
		return nil, fmt.Errorf("tree must be a branch")
	}

	branch := tree.(Branch)
	if len(branch) == 0 {
		return nil, fmt.Errorf("empty branch")
	}

	node := &SessionNode{
		Children: []SessionNode{},
	}

	for _, elem := range branch {
		if IsLeaf(elem) {
			leaf := elem.(Leaf)
			if len(leaf.Value) == 0 {
				continue
			}

			flag := leaf.Value[0]

			if flag == SESSIONS_FLAG_PERMISSIONS && len(leaf.Value) >= 21 {
				node.Signer = common.BytesToAddress(leaf.Value[1:21])

				// TODO: Decode permissions
			}
		} else if IsBranch(elem) {
			childNode, err := treeToSessionNode(elem)
			if err == nil {
				node.Children = append(node.Children, *childNode)
			}
		}
	}

	return node, nil
}
