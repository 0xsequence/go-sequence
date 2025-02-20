package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sequence",
		Short: "Sequence CLI tool",
		Long:  `A command line interface for interacting with Sequence services`,
	}

	rootCmd.AddCommand(
		// newPayloadCmd(),
		// newConfigCmd(),
		// newDevToolsCmd(),
		// newSignatureCmd(),
		// newPermissionCmd(),
		// newSessionCmd(),
		newServerCmd(),
		newAddressCmd(),
		newConfigCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
