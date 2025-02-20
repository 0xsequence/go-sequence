package sequence

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/spf13/cobra"
)

func newAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Address related commands",
	}

	cmd.AddCommand(newAddressCalculateCmd())

	return cmd
}

func newAddressCalculateCmd() *cobra.Command {
	var creationCode string

	cmd := &cobra.Command{
		Use:   "calculate <imageHash> <factory> <module>",
		Short: "Calculate counterfactual wallet address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			imageHash := common.HexToHash(args[0])
			factory := common.HexToAddress(args[1])
			module := common.HexToAddress(args[2])

			if creationCode == "" {
				creationCode = contracts.DEFAULT_CREATION_CODE
			}

			address, err := contracts.GetCounterfactualAddress(factory, module, imageHash, creationCode)
			if err != nil {
				return fmt.Errorf("failed to calculate address: %w", err)
			}

			fmt.Println(address.Hex())
			return nil
		},
	}

	cmd.Flags().StringVar(&creationCode, "creationCode", contracts.DEFAULT_CREATION_CODE, "Creation code for the wallet")

	return cmd
}
