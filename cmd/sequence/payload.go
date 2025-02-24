package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	v3 "github.com/0xsequence/go-sequence/core/v3"

	"github.com/spf13/cobra"
)

func handleToAbi(p *PayloadToAbiParams) (string, error) {
	log.Printf("payload: %s", p.Payload)
	return "", fmt.Errorf("not implemented")
}

func handleToPacked(p *PayloadToPackedParams) (string, error) {
	log.Printf("payload: %s", p.Payload)
	decoded, err := v3.DecodeABIPayload(p.Payload)
	if err != nil {
		return "", fmt.Errorf("failed to decode ABI payload: %w", err)
	}
	if decoded.Kind != v3.KindTransactions {
		return "", fmt.Errorf("conversion to packed only implemented for call payloads")
	}
	packed, err := v3.Encode(decoded, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encode to packed: %w", err)
	}
	return hexutil.Encode(packed), nil
}

func handleToJson(p *PayloadToJsonParams) (string, error) {
	decoded, err := v3.DecodeABIPayload(p.Payload)
	if err != nil {
		return "", fmt.Errorf("failed to decode ABI payload: %w", err)
	}
	jsonBytes, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func handleHash(p *PayloadHashParams) (string, error) {
	decoded, err := v3.DecodeABIPayload(p.Payload)
	if err != nil {
		return "", fmt.Errorf("failed to decode ABI payload: %w", err)
	}
	chainIDInt := new(big.Int)
	if _, ok := chainIDInt.SetString(p.ChainId, 10); !ok {
		return "", fmt.Errorf("invalid chain ID: %s", p.ChainId)
	}
	walletAddr := common.HexToAddress(p.Wallet)
	digest, err := v3.HashPayload(walletAddr, chainIDInt, decoded)
	if err != nil {
		return "", fmt.Errorf("failed to hash payload: %w", err)
	}
	return hexutil.Encode(digest[:]), nil
}

// newPayloadCmd defines the command-line interface.
func newPayloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payload",
		Short: "Payload conversion utilities",
	}

	toAbiCmd := &cobra.Command{
		Use:   "to-abi [payload]",
		Short: "Convert payload to ABI format",
		RunE: func(cmd *cobra.Command, args []string) error {
			payload, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleToAbi(&PayloadToAbiParams{Payload: payload})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	toPackedCmd := &cobra.Command{
		Use:   "to-packed [payload]",
		Short: "Convert payload to packed format",
		RunE: func(cmd *cobra.Command, args []string) error {
			payload, err := fromPosOrStdin(args, 0)
			if err != nil {
				return err
			}
			result, err := handleToPacked(&PayloadToPackedParams{Payload: payload})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
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

	cmd.AddCommand(toAbiCmd, toPackedCmd, toJsonCmd, hashCmd)
	return cmd
}
