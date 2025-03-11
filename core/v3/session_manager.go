package v3

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
)

// SessionManager manages cached session data, particularly identity signers
type SessionManager struct {
	identitySigners map[string]string
	mu              sync.RWMutex
}

// NewSessionManager creates a new session manager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		identitySigners: make(map[string]string),
	}
}

// extractStableKeyFromTopology tries to extract a stable key from the topology data
func (sm *SessionManager) extractStableKeyFromTopology(topologyData json.RawMessage) string {
	var topArray []json.RawMessage
	if err := json.Unmarshal(topologyData, &topArray); err == nil && len(topArray) >= 2 {

		var secondArray []json.RawMessage
		if err := json.Unmarshal(topArray[1], &secondArray); err == nil && len(secondArray) > 0 {
			var signerObj struct {
				IdentitySigner string `json:"identitySigner"`
			}

			if err := json.Unmarshal(secondArray[0], &signerObj); err == nil && signerObj.IdentitySigner != "" {
				return "identitySigner:" + strings.ToLower(signerObj.IdentitySigner)
			}
		}
	}

	return string(topologyData)
}

// StoreIdentitySigner stores an identity signer for a given topology hash
func (sm *SessionManager) StoreIdentitySigner(topologyData json.RawMessage, identitySigner string) {
	if identitySigner == "" {
		return
	}

	sm.mu.Lock()
	defer sm.mu.Unlock()

	key := sm.extractStableKeyFromTopology(topologyData)
	sm.identitySigners[key] = identitySigner

	fallbackKey := string(topologyData)
	if fallbackKey != key {
		sm.identitySigners[fallbackKey] = identitySigner
		log.Printf("Also stored identity signer with fallback key")
	}
}

// GetIdentitySigner retrieves an identity signer for a given topology hash
func (sm *SessionManager) GetIdentitySigner(topologyData json.RawMessage) string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	stableKey := sm.extractStableKeyFromTopology(topologyData)
	if value, ok := sm.identitySigners[stableKey]; ok {
		log.Printf("Retrieved identity signer %s using stable key", value)
		return value
	}

	directKey := string(topologyData)
	value := sm.identitySigners[directKey]
	if value != "" {
		log.Printf("Retrieved identity signer %s using fallback key", value)
	} else {
		log.Printf("No identity signer found for topology")
	}
	return value
}

// ListStoredSigners lists all stored identity signers for debugging
func (sm *SessionManager) ListStoredSigners() []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	result := []string{}
	for k, v := range sm.identitySigners {
		if strings.HasPrefix(k, "identitySigner:") {
			result = append(result, fmt.Sprintf("%s -> %s", k, v))
		}
	}
	return result
}

// DefaultSessionManager provides a default session manager instance
var DefaultSessionManager = NewSessionManager()
