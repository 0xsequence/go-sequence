package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/spf13/cobra"
)

type SeededRandom struct {
	currentSeed string
	hash        []byte
	index       int
}

func createSeededRandom(seed string) *SeededRandom {
	hash := sha256.Sum256([]byte(seed))
	return &SeededRandom{
		currentSeed: seed,
		hash:        hash[:],
		index:       0,
	}
}

func (r *SeededRandom) Float64() float64 {
	if r.index >= len(r.hash)-4 {
		r.currentSeed = r.currentSeed + "1"
		hash := sha256.Sum256([]byte(r.currentSeed))
		r.hash = hash[:]
		r.index = 0
	}

	value := float64(binary.LittleEndian.Uint32(r.hash[r.index:r.index+4])) / float64(math.MaxUint32)
	r.index += 4
	return value
}

func randomBytes(length int, random *SeededRandom) []byte {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(random.Float64() * 256)
	}
	return bytes
}

func randomBigInt(max *big.Int, random *SeededRandom) *big.Int {
	val := new(big.Int).SetUint64(uint64(random.Float64() * float64(max.Uint64())))
	return val
}

func randomAddress(random *SeededRandom) common.Address {
	return common.BytesToAddress(randomBytes(20, random))
}

func generateRandomTopology(depth int, random *SeededRandom, minThresholdOnNested uint64) interface{} {
	if depth <= 0 {
		leafType := int(random.Float64() * 5)

		switch leafType {
		case 0: // SignerLeaf
			return map[string]interface{}{
				"type":    "signer",
				"address": randomAddress(random).Hex(),
				"weight":  randomBigInt(big.NewInt(256), random).Uint64(),
			}

		case 1: // SapientSigner
			return map[string]interface{}{
				"type":      "sapient-signer",
				"address":   randomAddress(random).Hex(),
				"weight":    randomBigInt(big.NewInt(256), random).Uint64(),
				"imageHash": fmt.Sprintf("0x%x", randomBytes(32, random)),
			}

		case 2: // SubdigestLeaf
			return map[string]interface{}{
				"type":   "subdigest",
				"digest": fmt.Sprintf("0x%x", randomBytes(32, random)),
			}

		case 3: // NodeLeaf
			return map[string]interface{}{
				"node": fmt.Sprintf("0x%x", randomBytes(32, random)),
			}

		case 4: // NestedLeaf
			minThreshold := big.NewInt(int64(minThresholdOnNested))
			maxThreshold := new(big.Int).Sub(big.NewInt(65535), minThreshold)
			threshold := new(big.Int).Add(minThreshold, randomBigInt(maxThreshold, random))

			return map[string]interface{}{
				"type":      "nested",
				"tree":      generateRandomTopology(0, random, minThresholdOnNested),
				"weight":    randomBigInt(big.NewInt(256), random).Uint64(),
				"threshold": threshold.Uint64(),
			}
		}
	}

	// Generate a node with two random subtrees
	return []interface{}{
		generateRandomTopology(depth-1, random, minThresholdOnNested),
		generateRandomTopology(depth-1, random, minThresholdOnNested),
	}
}

func generateRandomPermissionRule(random *SeededRandom) map[string]interface{} {
	return map[string]interface{}{
		"cumulative": random.Float64()*2 > 1,
		"operation":  uint8(random.Float64() * 4),
		"value":      fmt.Sprintf("0x%x", randomBytes(32, random)),
		"offset":     randomBigInt(big.NewInt(100), random).Uint64(),
		"mask":       fmt.Sprintf("0x%x", randomBytes(32, random)),
	}
}

func generateRandomPermission(random *SeededRandom, maxRules uint64) map[string]interface{} {
	rulesCount := uint64(random.Float64()*float64(maxRules)) + 1
	rules := make([]interface{}, rulesCount)
	for i := uint64(0); i < rulesCount; i++ {
		rules[i] = generateRandomPermissionRule(random)
	}

	return map[string]interface{}{
		"target": randomAddress(random).Hex(),
		"rules":  rules,
	}
}

func generateSessionsTopology(depth int, random *SeededRandom, maxPermissions, maxRules uint64) interface{} {
	isLeaf := random.Float64()*2 > 1

	if isLeaf || depth <= 1 {
		permissionsCount := uint64(random.Float64()*float64(maxPermissions)) + 1
		permissions := make([]interface{}, permissionsCount)
		for i := uint64(0); i < permissionsCount; i++ {
			permissions[i] = generateRandomPermission(random, maxRules)
		}

		return map[string]interface{}{
			"signer":      randomAddress(random).Hex(),
			"valueLimit":  randomBigInt(big.NewInt(100), random).Uint64(),
			"deadline":    randomBigInt(big.NewInt(1000), random).Uint64(),
			"permissions": permissions,
		}
	}

	return []interface{}{
		generateSessionsTopology(depth-1, random, maxPermissions, maxRules),
		generateSessionsTopology(depth-1, random, maxPermissions, maxRules),
	}
}

func handleRandomConfig(p *RandomConfigParams) (interface{}, error) {
	random := createSeededRandom(p.Seed)

	config := map[string]interface{}{
		"threshold":  randomBigInt(big.NewInt(100), random).Uint64(),
		"checkpoint": randomBigInt(big.NewInt(1000), random).Uint64(),
		"topology":   generateRandomTopology(int(p.MaxDepth), random, p.MinThresholdOnNested),
	}

	// Handle checkpointer based on mode
	switch p.Checkpointer {
	case "yes":
		config["checkpointer"] = randomAddress(random).Hex()
	case "random":
		if random.Float64() > 0.5 {
			config["checkpointer"] = randomAddress(random).Hex()
		}
	}

	return config, nil
}

func handleRandomSessionTopology(p *RandomSessionTopologyParams) (interface{}, error) {
	random := createSeededRandom(p.Seed)
	topology := generateSessionsTopology(int(p.MaxDepth), random, p.MaxPermissions, p.MaxRules)
	return topology, nil
}

func newDevToolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dev-tools",
		Short: "Development tools and utilities",
	}

	randomConfigCmd := &cobra.Command{
		Use:   "random-config",
		Short: "Generate a random configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			maxDepth, _ := cmd.Flags().GetUint64("max-depth")
			seed, _ := cmd.Flags().GetString("seed")
			minThreshold, _ := cmd.Flags().GetUint64("min-threshold-on-nested")
			checkpointer, _ := cmd.Flags().GetString("checkpointer")

			params := &RandomConfigParams{
				MaxDepth:             maxDepth,
				Seed:                 seed,
				MinThresholdOnNested: minThreshold,
				Checkpointer:         checkpointer,
			}

			result, err := handleRandomConfig(params)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	randomConfigCmd.Flags().Uint64("max-depth", 3, "Maximum depth of the configuration tree")
	randomConfigCmd.Flags().String("seed", "", "Seed for deterministic generation")
	randomConfigCmd.Flags().Uint64("min-threshold-on-nested", 0, "Minimum threshold value for nested leaves")
	randomConfigCmd.Flags().String("checkpointer", "no", "Checkpointer mode: no (never add), random (50% chance), yes (always add)")

	randomSessionCmd := &cobra.Command{
		Use:   "random-session-topology",
		Short: "Generate a random session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			maxDepth, _ := cmd.Flags().GetUint64("max-depth")
			maxPermissions, _ := cmd.Flags().GetUint64("max-permissions")
			maxRules, _ := cmd.Flags().GetUint64("max-rules")
			seed, _ := cmd.Flags().GetString("seed")

			params := &RandomSessionTopologyParams{
				MaxDepth:       maxDepth,
				MaxPermissions: maxPermissions,
				MaxRules:       maxRules,
				Seed:           seed,
			}

			result, err := handleRandomSessionTopology(params)
			if err != nil {
				return err
			}

			output, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	randomSessionCmd.Flags().Uint64("max-depth", 1, "Maximum depth of the session topology")
	randomSessionCmd.Flags().Uint64("max-permissions", 1, "Maximum number of permissions in each session")
	randomSessionCmd.Flags().Uint64("max-rules", 1, "Maximum number of rules in each permission")
	randomSessionCmd.Flags().String("seed", "", "Seed for deterministic generation")

	cmd.AddCommand(randomConfigCmd, randomSessionCmd)

	return cmd
}
