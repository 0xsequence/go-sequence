package main

import (
	"encoding/json"
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

func handleToPackedSession(p *PermissionToPackedSessionParams) (string, error) {
	var permission v3.SessionPermissions
	err := json.Unmarshal([]byte(p.SessionPermission), &permission)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal session permission: %w", err)
	}

	packed, err := v3.EncodeSessionPermissions(&permission)
	if err != nil {
		return "", fmt.Errorf("failed to encode session permission: %w", err)
	}

	return hexutil.Encode(packed), nil
}

func handleToPackedPermission(p *PermissionToPackedParams) (string, error) {
	var permission v3.Permission
	err := json.Unmarshal([]byte(p.Permission), &permission)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal permission: %w", err)
	}

	packed, err := v3.EncodePermission(&permission)
	if err != nil {
		return "", fmt.Errorf("failed to encode permission: %w", err)
	}

	return hexutil.Encode(packed), nil
}

// newPermissionCmd defines the command-line interface.
func newPermissionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "permission",
		Short: "Permission conversion utilities",
	}

	toPackedSessionCmd := &cobra.Command{
		Use:   "to-packed-session [session-permission]",
		Short: "Convert session permission to packed format",
		RunE: func(cmd *cobra.Command, args []string) error {
			permission, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleToPackedSession(&PermissionToPackedSessionParams{SessionPermission: permission})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	toPackedCmd := &cobra.Command{
		Use:   "to-packed [permission]",
		Short: "Convert permission to packed format",
		RunE: func(cmd *cobra.Command, args []string) error {
			permission, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleToPackedPermission(&PermissionToPackedParams{Permission: permission})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(toPackedSessionCmd, toPackedCmd)
	return cmd
}
