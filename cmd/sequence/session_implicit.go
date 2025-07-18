package main

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

func handleAddBlacklist(p *AddBlacklistParams) (string, error) {
	blacklistAddr := common.HexToAddress(p.BlacklistAddress)
	updated, err := v3.AddToImplicitBlacklist(p.SessionTopology, blacklistAddr)
	if err != nil {
		return "", fmt.Errorf("failed to add to implicit blacklist: %w", err)
	}

	jsonBytes, err := json.Marshal(updated)
	if err != nil {
		return "", fmt.Errorf("failed to marshal updated topology: %w", err)
	}

	return string(jsonBytes), nil
}

func handleRemoveBlacklist(p *RemoveBlacklistParams) (string, error) {
	blacklistAddr := common.HexToAddress(p.BlacklistAddress)
	updated, err := v3.RemoveFromImplicitBlacklist(p.SessionTopology, blacklistAddr)
	if err != nil {
		return "", fmt.Errorf("failed to remove from implicit blacklist: %w", err)
	}

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
		Use:   "blacklist-add [blacklist-address] [session-topology]",
		Short: "Add an address to the implicit session blacklist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("blacklist address is required")
			}

			sessionTopologyString, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}
			sessionTopology, err := v3.SessionsTopologyFromJSON(sessionTopologyString)
			if err != nil {
				return fmt.Errorf("failed to decode session topology: %w", err)
			}

			result, err := handleAddBlacklist(&AddBlacklistParams{
				BlacklistAddress: args[0],
				SessionTopology:  sessionTopology,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	removeCmd := &cobra.Command{
		Use:   "blacklist-remove [blacklist-address] [session-topology]",
		Short: "Remove an address from the implicit session blacklist",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("blacklist address is required")
			}

			sessionTopologyString, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}
			sessionTopology, err := v3.SessionsTopologyFromJSON(sessionTopologyString)
			if err != nil {
				return fmt.Errorf("failed to decode session topology: %w", err)
			}

			result, err := handleRemoveBlacklist(&RemoveBlacklistParams{
				BlacklistAddress: args[0],
				SessionTopology:  sessionTopology,
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
