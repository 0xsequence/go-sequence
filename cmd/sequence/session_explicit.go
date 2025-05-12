package main

import (
	"encoding/json"
	"fmt"

	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

func handleAddExplicitSession(p *AddSessionParams) (string, error) {
	var session v3.SessionPermissions
	if err := json.Unmarshal(p.ExplicitSession, &session); err != nil {
		return "", fmt.Errorf("failed to decode explicit session: %w", err)
	}

	updated, err := v3.AddExplicitSession(p.SessionTopology, session)
	if err != nil {
		return "", fmt.Errorf("failed to add explicit session: %w", err)
	}

	return v3.SessionsTopologyToJSON(updated)
}

func handleRemoveExplicitSession(p *RemoveSessionParams) (string, error) {
	if p.ExplicitSessionAddress == "" || p.ExplicitSessionAddress[:2] != "0x" {
		return "", fmt.Errorf("explicit session address must be a valid address")
	}

	updated, err := v3.RemoveExplicitSession(p.SessionTopology, p.ExplicitSessionAddress)
	if err != nil {
		return "", fmt.Errorf("failed to remove explicit session: %w", err)
	}

	return v3.SessionsTopologyToJSON(updated)
}

// newSessionExplicitCmd defines the command-line interface.
func newSessionExplicitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "explicit",
		Short: "Explicit session utilities",
	}

	addCmd := &cobra.Command{
		Use:   "add [explicit-session] [session-topology]",
		Short: "Add an explicit session to a session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			explicitSession, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			sessionTopologyString, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}
			sessionTopology, err := v3.SessionsTopologyFromJSON(sessionTopologyString)
			if err != nil {
				return fmt.Errorf("failed to decode session topology: %w", err)
			}

			result, err := handleAddExplicitSession(&AddSessionParams{
				ExplicitSession: json.RawMessage(explicitSession),
				SessionTopology: sessionTopology,
			})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	removeCmd := &cobra.Command{
		Use:   "remove [explicit-session-address] [session-topology]",
		Short: "Remove a session from the session topology",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("explicit session address is required")
			}

			sessionTopologyString, err := fromPosOrStdin(args, 1)
			if err != nil {
				return err
			}
			sessionTopology, err := v3.SessionsTopologyFromJSON(sessionTopologyString)
			if err != nil {
				return fmt.Errorf("failed to decode session topology: %w", err)
			}

			result, err := handleRemoveExplicitSession(&RemoveSessionParams{
				ExplicitSessionAddress: args[0],
				SessionTopology:        sessionTopology,
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
