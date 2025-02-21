package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

const (
	KindTransactions uint8 = 0x00
	KindMessage      uint8 = 0x01
	KindConfigUpdate uint8 = 0x02
	KindDigest       uint8 = 0x03

	BehaviorIgnoreError   uint8 = 0x00
	BehaviorRevertOnError uint8 = 0x01
	BehaviorAbortOnError  uint8 = 0x02
)

type Call struct {
	To              common.Address `json:"to"`
	Value           *big.Int       `json:"value"`
	Data            hexutil.Bytes  `json:"data"`
	GasLimit        *big.Int       `json:"gasLimit"`
	DelegateCall    bool           `json:"delegateCall"`
	OnlyFallback    bool           `json:"onlyFallback"`
	BehaviorOnError uint8          `json:"behaviorOnError"`
}

type DecodedPayload struct {
	Kind          uint8            `json:"kind"`
	NoChainId     bool             `json:"noChainId"`
	Calls         []Call           `json:"calls"`
	Space         *big.Int         `json:"space"`
	Nonce         *big.Int         `json:"nonce"`
	Message       hexutil.Bytes    `json:"message"`
	ImageHash     common.Hash      `json:"imageHash"`
	Digest        common.Hash      `json:"digest"`
	ParentWallets []common.Address `json:"parentWallets"`
}

// func behaviorToString(behavior uint8) string {
// 	switch behavior {
// 	case BehaviorIgnoreError:
// 		return "ignore"
// 	case BehaviorRevertOnError:
// 		return "revert"
// 	case BehaviorAbortOnError:
// 		return "abort"
// 	default:
// 		return fmt.Sprintf("unknown(%d)", behavior)
// 	}
// }

func convertToAbi(payload string) (string, error) {
	log.Printf("convertToAbi: %s", payload)
	// Not implemented yet, following the pattern
	return "", fmt.Errorf("not implemented")
}

func convertToPacked(payload string) (string, error) {
	var decoded DecodedPayload
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}

	if decoded.Kind == KindTransactions {
		// TODO: Implement the packing logic using sequence-primitives
		// This will require implementing the encoding logic in Go
		return "", fmt.Errorf("not implemented")
	}

	return "", fmt.Errorf("not implemented")
}

func convertToJson(payload string) (string, error) {
	// TODO: Implement ABI decoding using ethkit
	var decoded DecodedPayload
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		return "", fmt.Errorf("failed to decode payload: %w", err)
	}

	result, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}

	return string(result), nil
}

func newPayloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payload",
		Short: "Payload conversion utilities",
	}

	toAbiCmd := &cobra.Command{
		Use:   "to-abi [payload]",
		Short: "Convert payload to ABI format",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("payload argument is required")
			}

			result, err := convertToAbi(args[0])
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
			if len(args) != 1 {
				return fmt.Errorf("payload argument is required")
			}

			result, err := convertToPacked(args[0])
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
			if len(args) != 1 {
				return fmt.Errorf("payload argument is required")
			}

			result, err := convertToJson(args[0])
			if err != nil {
				return err
			}

			fmt.Println(result)
			return nil
		},
	}

	cmd.AddCommand(toAbiCmd, toPackedCmd, toJsonCmd)

	return cmd
}
