package main

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

func handleAddBlacklist(p *AddBlacklistParams) (string, error) {
	var topology v3.SessionsTopology
	if err := json.Unmarshal([]byte(p.SessionConfiguration), &topology); err != nil {
		return "", fmt.Errorf("failed to decode session configuration: %w", err)
	}

	blacklistAddr := common.HexToAddress(p.BlacklistAddress)
	updated := v3.AddToImplicitBlacklist(&topology, blacklistAddr)

	jsonBytes, err := json.Marshal(updated)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated topology: %w", err)
	}

	return string(jsonBytes), nil
}

func handleRemoveBlacklist(p *RemoveBlacklistParams) (string, error) {
	var topology v3.SessionsTopology
	if err := json.Unmarshal([]byte(p.SessionConfiguration), &topology); err != nil {
		return "", fmt.Errorf("failed to decode session configuration: %w", err)
	}

	blacklistAddr := common.HexToAddress(p.BlacklistAddress)
	updated := v3.RemoveFromImplicitBlacklist(&topology, blacklistAddr)

	jsonBytes, err := json.Marshal(updated)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated topology: %w", err)
	}

	return string(jsonBytes), nil
}

// newSessionImplicitCmd defines the command-line interface.
func newSessionImplicitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "implicit",
		Short: "Implicit session utilities",
	}

	addCmd := &cobra.Command{
		Use:   "blacklist-add [blacklist-address] [session-configuration]",
		Short: "Add an address to the implicit session blacklist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("blacklist address is required")
			}

			config, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}

			result, err := handleAddBlacklist(&AddBlacklistParams{
				BlacklistAddress:     args[0],
				SessionConfiguration: config,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	removeCmd := &cobra.Command{
		Use:   "blacklist-remove [blacklist-address] [session-configuration]",
		Short: "Remove an address from the implicit session blacklist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("blacklist address is required")
			}

			config, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}

			result, err := handleRemoveBlacklist(&RemoveBlacklistParams{
				BlacklistAddress:     args[0],
				SessionConfiguration: config,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(addCmd, removeCmd)
	return cmd
}
