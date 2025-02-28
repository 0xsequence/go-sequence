package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/spf13/cobra"
)

type RandomOptions struct {
	SeededRandom         *SeededRandom
	MinThresholdOnNested uint64
	MaxPermissions       uint64
	MaxRules             uint64
	CheckpointerMode     string // 'no', 'random', or 'yes'
}

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

func randomBytes(length int, options *RandomOptions) []byte {
	bytes := make([]byte, length)
	if options != nil && options.SeededRandom != nil {
		for i := 0; i < length; i++ {
			bytes[i] = byte(options.SeededRandom.Float64() * 256)
		}
		return bytes
	}
	// Fallback to crypto/rand in a real implementation
	for i := 0; i < length; i++ {
		bytes[i] = byte(math.Floor(rand.Float64() * 256))
	}
	return bytes
}

func randomBigInt(max *big.Int, options *RandomOptions) *big.Int {
	if options != nil && options.SeededRandom != nil {
		val := new(big.Int).SetUint64(uint64(options.SeededRandom.Float64() * float64(max.Uint64())))
		return val
	}
	// Fallback to math/rand in a real implementation
	val := new(big.Int).SetUint64(uint64(rand.Float64() * float64(max.Uint64())))
	return val
}

func randomAddress(options *RandomOptions) common.Address {
	return common.BytesToAddress(randomBytes(20, options))
}

func generateRandomTopology(depth int, options *RandomOptions) interface{} {
	if depth <= 0 {
		var leafType int
		if options != nil && options.SeededRandom != nil {
			leafType = int(options.SeededRandom.Float64() * 5)
		} else {
			leafType = int(rand.Float64() * 5)
		}

		switch leafType {
		case 0: // SignerLeaf
			return map[string]interface{}{
				"type":    "signer",
				"address": randomAddress(options).Hex(),
				"weight":  randomBigInt(big.NewInt(256), options).Uint64(),
			}

		case 1: // SapientSigner
			return map[string]interface{}{
				"type":      "sapient-signer",
				"address":   randomAddress(options).Hex(),
				"weight":    randomBigInt(big.NewInt(256), options).Uint64(),
				"imageHash": fmt.Sprintf("0x%x", randomBytes(32, options)),
			}

		case 2: // SubdigestLeaf
			return map[string]interface{}{
				"type":   "subdigest",
				"digest": fmt.Sprintf("0x%x", randomBytes(32, options)),
			}

		case 3: // NodeLeaf
			return fmt.Sprintf("0x%x", randomBytes(32, options))

		case 4: // NestedLeaf
			minThreshold := big.NewInt(0)
			if options != nil && options.MinThresholdOnNested > 0 {
				minThreshold = big.NewInt(int64(options.MinThresholdOnNested))
			}
			maxThreshold := new(big.Int).Sub(big.NewInt(65535), minThreshold)
			threshold := new(big.Int).Add(minThreshold, randomBigInt(maxThreshold, options))

			return map[string]interface{}{
				"type":      "nested",
				"tree":      generateRandomTopology(0, options),
				"weight":    randomBigInt(big.NewInt(256), options).Uint64(),
				"threshold": threshold.Uint64(),
			}
		}
	}

	// Generate a node with two random subtrees
	return []interface{}{
		generateRandomTopology(depth-1, options),
		generateRandomTopology(depth-1, options),
	}
}

func generateRandomRule(options *RandomOptions) map[string]interface{} {
	var cumulative bool
	if options != nil && options.SeededRandom != nil {
		cumulative = options.SeededRandom.Float64()*2 > 1
	} else {
		cumulative = rand.Float64()*2 > 1
	}

	return map[string]interface{}{
		"cumulative": cumulative,
		"operation":  uint8(rand.Float64() * 4),
		"value":      fmt.Sprintf("0x%x", randomBytes(32, options)),
		"offset":     randomBigInt(big.NewInt(100), options).Uint64(),
		"mask":       fmt.Sprintf("0x%x", randomBytes(32, options)),
	}
}

func generateRandomPermission(options *RandomOptions) map[string]interface{} {
	var rulesCount uint64
	if options != nil && options.MaxRules > 0 {
		if options.SeededRandom != nil {
			rulesCount = uint64(options.SeededRandom.Float64()*float64(options.MaxRules)) + 1
		} else {
			rulesCount = uint64(rand.Float64()*float64(options.MaxRules)) + 1
		}
	} else {
		rulesCount = uint64(rand.Float64()*5) + 1
	}

	rules := make([]interface{}, rulesCount)
	for i := uint64(0); i < rulesCount; i++ {
		rules[i] = generateRandomRule(options)
	}

	return map[string]interface{}{
		"target": randomAddress(options).Hex(),
		"rules":  rules,
	}
}

func generateSessionsTopology(depth int, options *RandomOptions) interface{} {
	var isLeaf bool
	if options != nil && options.SeededRandom != nil {
		isLeaf = options.SeededRandom.Float64()*2 > 1
	} else {
		isLeaf = rand.Float64()*2 > 1
	}

	if isLeaf || depth <= 1 {
		var permissionsCount uint64
		if options != nil && options.MaxPermissions > 0 {
			if options.SeededRandom != nil {
				permissionsCount = uint64(options.SeededRandom.Float64()*float64(options.MaxPermissions)) + 1
			} else {
				permissionsCount = uint64(rand.Float64()*float64(options.MaxPermissions)) + 1
			}
		} else {
			permissionsCount = uint64(rand.Float64()*5) + 1
		}

		permissions := make([]interface{}, permissionsCount)
		for i := uint64(0); i < permissionsCount; i++ {
			permissions[i] = generateRandomPermission(options)
		}

		return map[string]interface{}{
			"signer":      randomAddress(options).Hex(),
			"valueLimit":  randomBigInt(big.NewInt(100), options).Uint64(),
			"deadline":    randomBigInt(big.NewInt(1000), options).Uint64(),
			"permissions": permissions,
		}
	}

	return []interface{}{
		generateSessionsTopology(depth-1, options),
		generateSessionsTopology(depth-1, options),
	}
}

func doRandomConfig(maxDepth int, options *RandomOptions) (string, error) {
	config := map[string]interface{}{
		"threshold":  randomBigInt(big.NewInt(100), options).Uint64(),
		"checkpoint": randomBigInt(big.NewInt(1000), options).Uint64(),
		"topology":   generateRandomTopology(maxDepth, options),
	}

	// Handle checkpointer based on mode
	if options != nil {
		switch options.CheckpointerMode {
		case "yes":
			config["checkpointer"] = randomAddress(options).Hex()
		case "random":
			isAdded := false
			if options.SeededRandom != nil {
				isAdded = options.SeededRandom.Float64() > 0.5
			} else {
				isAdded = rand.Float64() > 0.5
			}
			if isAdded {
				config["checkpointer"] = randomAddress(options).Hex()
			}
		}
	}

	output, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func doRandomSessionTopology(maxDepth int, options *RandomOptions) (string, error) {
	topology := generateSessionsTopology(maxDepth, options)
	output, err := json.MarshalIndent(topology, "", "  ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// Handler functions for RPC calls
func handleRandomConfig(p *RandomConfigParams) (string, error) {
	options := &RandomOptions{
		MinThresholdOnNested: p.MinThresholdOnNested,
		CheckpointerMode:     p.Checkpointer,
	}

	if p.Seed != "" {
		options.SeededRandom = createSeededRandom(p.Seed)
	}

	return doRandomConfig(int(p.MaxDepth), options)
}

func handleRandomSessionTopology(p *RandomSessionTopologyParams) (string, error) {
	options := &RandomOptions{
		MaxPermissions: p.MaxPermissions,
		MaxRules:       p.MaxRules,
	}

	if p.Seed != "" {
		options.SeededRandom = createSeededRandom(p.Seed)
	}

	return doRandomSessionTopology(int(p.MaxDepth), options)
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
			maxDepth, _ := cmd.Flags().GetInt("max-depth")
			seed, _ := cmd.Flags().GetString("seed")
			minThreshold, _ := cmd.Flags().GetUint64("min-threshold-on-nested")
			checkpointer, _ := cmd.Flags().GetString("checkpointer")

			options := &RandomOptions{
				MinThresholdOnNested: minThreshold,
				CheckpointerMode:     checkpointer,
			}

			if seed != "" {
				options.SeededRandom = createSeededRandom(seed)
			}

			result, err := doRandomConfig(maxDepth, options)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	randomConfigCmd.Flags().Int("max-depth", 3, "Maximum depth of the configuration tree")
	randomConfigCmd.Flags().String("seed", "", "Seed for deterministic generation")
	randomConfigCmd.Flags().Uint64("min-threshold-on-nested", 0, "Minimum threshold value for nested leaves")
	randomConfigCmd.Flags().String("checkpointer", "no", "Checkpointer mode: no (never add), random (50% chance), yes (always add)")

	randomSessionCmd := &cobra.Command{
		Use:   "random-session-topology",
		Short: "Generate a random session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			maxDepth, _ := cmd.Flags().GetInt("max-depth")
			maxPermissions, _ := cmd.Flags().GetUint64("max-permissions")
			maxRules, _ := cmd.Flags().GetUint64("max-rules")
			seed, _ := cmd.Flags().GetString("seed")

			options := &RandomOptions{
				MaxPermissions: maxPermissions,
				MaxRules:       maxRules,
			}

			if seed != "" {
				options.SeededRandom = createSeededRandom(seed)
			}

			result, err := doRandomSessionTopology(maxDepth, options)
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	randomSessionCmd.Flags().Int("max-depth", 1, "Maximum depth of the session topology")
	randomSessionCmd.Flags().Uint64("max-permissions", 1, "Maximum number of permissions in each session")
	randomSessionCmd.Flags().Uint64("max-rules", 1, "Maximum number of rules in each permission")
	randomSessionCmd.Flags().String("seed", "", "Seed for deterministic generation")

	cmd.AddCommand(randomConfigCmd, randomSessionCmd)

	return cmd
}
