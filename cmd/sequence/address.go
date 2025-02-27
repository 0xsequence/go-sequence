package main

import (
	"fmt"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	"github.com/spf13/cobra"
)

func calculateAddress(params *AddressCalculateParams) (string, error) {
	imageHash := core.ImageHash{Hash: common.HexToHash(params.ImageHash)}
	factory := common.HexToAddress(params.Factory)
	module := common.HexToAddress(params.Module)

	creationCode := params.CreationCode
	if creationCode == "" {
		creationCode = hexutil.Encode(contracts.V3.CreationCode)
	}

	address, err := sequence.AddressFromImageHash(imageHash.Hex(), sequence.WalletContext{
		FactoryAddress:    factory,
		MainModuleAddress: module,
		CreationCode:      creationCode,
	})
	if err != nil {
		return "", fmt.Errorf("failed to calculate address: %w", err)
	}

	return address.Hex(), nil
}

func newAddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address",
		Short: "Address utilities",
	}

	var params AddressCalculateParams

	calculateCmd := &cobra.Command{
		Use:   "calculate",
		Short: "Calculate counterfactual address",
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := calculateAddress(&params)
			if err != nil {
				return err
			}
			fmt.Println(address)
			return nil
		},
	}

	calculateCmd.Flags().StringVar(&params.ImageHash, "image-hash", "", "Image hash")
	calculateCmd.Flags().StringVar(&params.Factory, "factory", "", "Factory address")
	calculateCmd.Flags().StringVar(&params.Module, "module", "", "Module address")
	calculateCmd.Flags().StringVar(&params.CreationCode, "creation-code", "", "Creation code (optional)")

	calculateCmd.MarkFlagRequired("image-hash")
	calculateCmd.MarkFlagRequired("factory")
	calculateCmd.MarkFlagRequired("module")

	cmd.AddCommand(calculateCmd)

	return cmd
}
