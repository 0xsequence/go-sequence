package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	v3 "github.com/0xsequence/go-sequence/core/v3"

	"github.com/spf13/cobra"
)

func handleToJson(p *PayloadToJsonParams) (string, error) {
	data, err := hexutil.Decode(p.Payload)
	if err != nil {
		return "", fmt.Errorf("unable to hex decode calls: %w", err)
	}
	payload, err := v3.DecodeCalls(common.Address{}, nil, data)
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}
	jsonBytes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func handleHash(p *PayloadHashParams) (string, error) {
	walletAddr := common.HexToAddress(p.Wallet)
	chainID := new(big.Int)
	if _, ok := chainID.SetString(p.ChainId, 10); !ok {
		return "", fmt.Errorf("invalid chain ID: %s", p.ChainId)
	}

	data, err := hexutil.Decode(p.Payload)
	if err != nil {
		return "", fmt.Errorf("unable to hex decode calls: %w", err)
	}
	payload, err := v3.DecodeCalls(walletAddr, chainID, data)
	if err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}

	return payload.Digest().String(), nil
}

// newPayloadCmd defines the command-line interface.
func newPayloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payload",
		Short: "Payload conversion utilities",
	}

	toJsonCmd := &cobra.Command{
		Use:   "to-json [payload]",
		Short: "Convert payload to JSON format",
		RunE: func(cmd *cobra.Command, args []string) error {
			payload, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleToJson(&PayloadToJsonParams{Payload: payload})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	hashCmd := &cobra.Command{
		Use:   "hash [payload]",
		Short: "Hash the payload",
		RunE: func(cmd *cobra.Command, args []string) error {
			payload, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			wallet, _ := cmd.Flags().GetString("wallet")
			chainId, _ := cmd.Flags().GetString("chain-id")
			result, err := handleHash(&PayloadHashParams{Payload: payload, Wallet: wallet, ChainId: chainId})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	hashCmd.Flags().String("wallet", "", "Wallet address")
	hashCmd.Flags().String("chain-id", "", "Chain ID")
	hashCmd.MarkFlagRequired("wallet")
	hashCmd.MarkFlagRequired("chain-id")

	cmd.AddCommand(toJsonCmd, hashCmd)
	return cmd
}
