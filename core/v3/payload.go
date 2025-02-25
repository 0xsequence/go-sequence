package v3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
)

type BehaviorOnError uint8

const (
	Ignore BehaviorOnError = 0
	Revert BehaviorOnError = 1
	Abort  BehaviorOnError = 2
)

type Call struct {
	To              common.Address
	Value           *big.Int
	Data            []byte
	GasLimit        *big.Int
	DelegateCall    bool
	OnlyFallback    bool
	BehaviorOnError BehaviorOnError
}

type DecodedPayload struct {
	Kind          uint8
	NoChainId     bool
	Calls         []Call
	Space         *big.Int
	Nonce         *big.Int
	Message       []byte
	ImageHash     common.Hash
	Digest        common.Hash
	ParentWallets []common.Address
}

const (
	KindTransactions uint8 = 0x00
	KindMessage      uint8 = 0x01
	KindConfigUpdate uint8 = 0x02
	KindDigest       uint8 = 0x03
)

func minIntBytesFor(val *big.Int) int {
	if val == nil {
		return 0
	}
	// Always encode at least 1 byte, even for zero
	return max(1, (val.BitLen()+7)/8)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func encodeBehaviorOnError(behavior BehaviorOnError) byte {
	switch behavior {
	case Ignore:
		return 0
	case Revert:
		return 1
	case Abort:
		return 2
	default:
		return 0 // Default to ignore if unknown
	}
}

// Encode encodes a transaction (call) payload into a compact byte array.
func Encode(payload DecodedPayload, self *common.Address) ([]byte, error) {
	if payload.Kind != KindTransactions {
		return nil, fmt.Errorf("encode only supports transaction payloads")
	}

	if payload.Space == nil {
		payload.Space = big.NewInt(0)
	}
	if payload.Nonce == nil {
		payload.Nonce = big.NewInt(0)
	}

	callsLen := len(payload.Calls)
	nonceBytesNeeded := minIntBytesFor(payload.Nonce)
	if nonceBytesNeeded > 15 {
		return nil, fmt.Errorf("nonce is too large (requires %d bytes, max 15)", nonceBytesNeeded)
	}

	/*
		globalFlag layout:
		  bit 0: spaceZeroFlag => 1 if space == 0, else 0
		  bits [1..3]: how many bytes we use to encode nonce
		  bit 4: singleCallFlag => 1 if there's exactly one call
		  bit 5: callsCountSizeFlag => 1 if #calls stored in 2 bytes, 0 if in 1 byte
		  (bits [6..7] are unused/free)
	*/
	var globalFlag byte
	if payload.Space.Sign() == 0 {
		globalFlag |= 0x01
	}
	globalFlag |= byte(nonceBytesNeeded << 1)
	if callsLen == 1 {
		globalFlag |= 0x10
	}

	// Determine calls count size (1 or 2 bytes) if not a single call
	var callsCountSize int
	if callsLen != 1 {
		if callsLen < 256 {
			callsCountSize = 1
		} else if callsLen < 65536 {
			callsCountSize = 2
			globalFlag |= 0x20
		} else {
			return nil, fmt.Errorf("too many calls: %d (max 65535)", callsLen)
		}
	}

	// Start building the output with a bytes.Buffer
	var out bytes.Buffer
	out.WriteByte(globalFlag)

	// If space isn't 0, store it as exactly 20 bytes (like uint160)
	if payload.Space.Sign() != 0 {
		spaceBytes := make([]byte, 20)
		payload.Space.FillBytes(spaceBytes)
		out.Write(spaceBytes)
	}

	// Encode nonce in nonceBytesNeeded bytes
	if nonceBytesNeeded > 0 {
		nonceBytes := make([]byte, nonceBytesNeeded)
		payload.Nonce.FillBytes(nonceBytes)
		out.Write(nonceBytes)
	}

	// Store callsLen if not single-call
	if callsLen != 1 {
		if callsCountSize == 1 {
			out.WriteByte(byte(callsLen))
		} else {
			out.Write([]byte{byte(callsLen >> 8), byte(callsLen)})
		}
	}

	// Encode each call
	for _, call := range payload.Calls {
		var flags byte
		if self != nil && call.To == *self {
			flags |= 0x01
		}
		if call.Value != nil && call.Value.Sign() != 0 {
			flags |= 0x02
		}
		if len(call.Data) > 0 {
			flags |= 0x04
		}
		if call.GasLimit != nil && call.GasLimit.Sign() != 0 {
			flags |= 0x08
		}
		if call.DelegateCall {
			flags |= 0x10
		}
		if call.OnlyFallback {
			flags |= 0x20
		}
		behavior := encodeBehaviorOnError(call.BehaviorOnError)
		flags |= behavior << 6

		// Write the flags byte
		out.WriteByte(flags)

		// If toSelf bit not set, store 20-byte address
		if flags&0x01 == 0 {
			addrBytes := call.To.Bytes()
			if len(addrBytes) != 20 {
				return nil, fmt.Errorf("invalid 'to' address length: %d bytes", len(addrBytes))
			}
			out.Write(addrBytes)
		}

		// If hasValue, store 32 bytes of value
		if flags&0x02 != 0 {
			valueBytes := make([]byte, 32)
			if call.Value == nil {
				call.Value = big.NewInt(0)
			}
			call.Value.FillBytes(valueBytes)
			out.Write(valueBytes)
		}

		// If hasData, store 3 bytes of data length + data
		if flags&0x04 != 0 {
			dataLen := len(call.Data)
			if dataLen > 0xffffff { // Max value for 3 bytes
				return nil, fmt.Errorf("call data too large: %d bytes (max 16777215)", dataLen)
			}
			dataLenBytes := []byte{
				byte(dataLen >> 16),
				byte(dataLen >> 8),
				byte(dataLen),
			}
			out.Write(dataLenBytes)
			out.Write(call.Data)
		}

		// If hasGasLimit, store 32 bytes of gasLimit
		if flags&0x08 != 0 {
			gasBytes := make([]byte, 32)
			if call.GasLimit == nil {
				call.GasLimit = big.NewInt(0)
			}
			call.GasLimit.FillBytes(gasBytes)
			out.Write(gasBytes)
		}
	}

	// Return the encoded bytes
	return out.Bytes(), nil
}

type TypedDataField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// hashPayload computes the EIP-712 hash of a payload.
func HashPayload(wallet common.Address, chainId *big.Int, payload DecodedPayload) ([32]byte, error) {
	domain := ethcoder.TypedDataDomain{
		Name:              "Sequence Wallet",
		Version:           "3",
		ChainID:           chainId,
		VerifyingContract: &wallet,
	}

	var types map[string][]TypedDataField
	var message map[string]interface{}
	var primaryType string

	switch payload.Kind {
	case KindTransactions:
		types = map[string][]TypedDataField{
			"Calls": {
				{Name: "calls", Type: "Call[]"},
				{Name: "space", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "wallets", Type: "address[]"},
			},
			"Call": {
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "gasLimit", Type: "uint256"},
				{Name: "delegateCall", Type: "bool"},
				{Name: "onlyFallback", Type: "bool"},
				{Name: "behaviorOnError", Type: "uint256"},
			},
		}
		primaryType = "Calls"
		calls := make([]map[string]interface{}, len(payload.Calls))
		for i, call := range payload.Calls {
			calls[i] = map[string]interface{}{
				"to":              call.To,
				"value":           call.Value,
				"data":            call.Data,
				"gasLimit":        call.GasLimit,
				"delegateCall":    call.DelegateCall,
				"onlyFallback":    call.OnlyFallback,
				"behaviorOnError": big.NewInt(int64(call.BehaviorOnError)),
			}
		}
		message = map[string]interface{}{
			"calls":   calls,
			"space":   payload.Space,
			"nonce":   payload.Nonce,
			"wallets": payload.ParentWallets,
		}

	case KindMessage:
		types = map[string][]TypedDataField{
			"Message": {
				{Name: "message", Type: "bytes"},
				{Name: "wallets", Type: "address[]"},
			},
		}
		primaryType = "Message"
		message = map[string]interface{}{
			"message": payload.Message,
			"wallets": payload.ParentWallets,
		}

	case KindConfigUpdate:
		types = map[string][]TypedDataField{
			"ConfigUpdate": {
				{Name: "imageHash", Type: "bytes32"},
				{Name: "wallets", Type: "address[]"},
			},
		}
		primaryType = "ConfigUpdate"
		message = map[string]interface{}{
			"imageHash": payload.ImageHash,
			"wallets":   payload.ParentWallets,
		}

	case KindDigest:
		types = map[string][]TypedDataField{
			"Digest": {
				{Name: "digest", Type: "bytes32"},
				{Name: "wallets", Type: "address[]"},
			},
		}
		primaryType = "Digest"
		message = map[string]interface{}{
			"digest":  payload.Digest,
			"wallets": payload.ParentWallets,
		}

	default:
		return [32]byte{}, fmt.Errorf("unknown payload kind: %d", payload.Kind)
	}

	typedDataJSON := map[string]interface{}{
		"types":       types,
		"primaryType": primaryType,
		"domain":      domain,
		"message":     message,
	}

	typedDataJSONStr, err := json.Marshal(typedDataJSON)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to marshal typed data to JSON: %w", err)
	}

	typedData, err := ethcoder.TypedDataFromJSON(string(typedDataJSONStr))
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to parse typed data from JSON: %w", err)
	}

	digest, err := typedData.EncodeDigest()
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to encode digest: %w", err)
	}

	var result [32]byte
	copy(result[:], digest)
	return result, nil
}

// DecodeABIPayload decodes an ABI-encoded payload into DecodedPayload.
func DecodeABIPayload(input string) (DecodedPayload, error) {
	// Define the ABI with proper tuple array formatting
	const abiDef = `[{
		"name": "f",
		"type": "function",
		"inputs": [{
			"type": "tuple",
			"components": [
				{"name": "kind", "type": "uint8"},
				{"name": "noChainId", "type": "bool"},
				{"name": "calls", "type": "tuple[]", "components": [
					{"name": "to", "type": "address"},
					{"name": "value", "type": "uint256"},
					{"name": "data", "type": "bytes"},
					{"name": "gasLimit", "type": "uint256"},
					{"name": "delegateCall", "type": "bool"},
					{"name": "onlyFallback", "type": "bool"},
					{"name": "behaviorOnError", "type": "uint256"}
				]},
				{"name": "space", "type": "uint256"},
				{"name": "nonce", "type": "uint256"},
				{"name": "message", "type": "bytes"},
				{"name": "imageHash", "type": "bytes32"},
				{"name": "digest", "type": "bytes32"},
				{"name": "parentWallets", "type": "address[]"}
			]
		}]
	}]`

	data, err := hexutil.Decode(input)
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("invalid hex input: %w", err)
	}

	// Parse the ABI definition
	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Decode the input data
	decodedData, err := parsedABI.Methods["f"].Inputs.Unpack(data)
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to decode ABI: %w", err)
	}

	if len(decodedData) != 1 {
		return DecodedPayload{}, fmt.Errorf("invalid decoded data length")
	}

	v := decodedData[0]
	if v == nil {
		return DecodedPayload{}, fmt.Errorf("decoded data is nil")
	}

	// Try to decode as a struct with JSON tags
	if s, ok := v.(struct {
		Kind      uint8 `json:"kind"`
		NoChainId bool  `json:"noChainId"`
		Calls     []struct {
			To              common.Address `json:"to"`
			Value           *big.Int       `json:"value"`
			Data            []uint8        `json:"data"`
			GasLimit        *big.Int       `json:"gasLimit"`
			DelegateCall    bool           `json:"delegateCall"`
			OnlyFallback    bool           `json:"onlyFallback"`
			BehaviorOnError *big.Int       `json:"behaviorOnError"`
		} `json:"calls"`
		Space         *big.Int         `json:"space"`
		Nonce         *big.Int         `json:"nonce"`
		Message       []uint8          `json:"message"`
		ImageHash     [32]uint8        `json:"imageHash"`
		Digest        [32]uint8        `json:"digest"`
		ParentWallets []common.Address `json:"parentWallets"`
	}); ok {
		calls := make([]Call, len(s.Calls))
		for i, call := range s.Calls {
			calls[i] = Call{
				To:              call.To,
				Value:           call.Value,
				Data:            call.Data,
				GasLimit:        call.GasLimit,
				DelegateCall:    call.DelegateCall,
				OnlyFallback:    call.OnlyFallback,
				BehaviorOnError: BehaviorOnError(call.BehaviorOnError.Uint64()),
			}
		}

		return DecodedPayload{
			Kind:          s.Kind,
			NoChainId:     s.NoChainId,
			Calls:         calls,
			Space:         s.Space,
			Nonce:         s.Nonce,
			Message:       s.Message,
			ImageHash:     common.BytesToHash(s.ImageHash[:]),
			Digest:        common.BytesToHash(s.Digest[:]),
			ParentWallets: s.ParentWallets,
		}, nil
	}

	// If struct decoding fails, try JSON marshaling and unmarshaling
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to marshal decoded data: %w", err)
	}

	var payload DecodedPayload
	if err := json.Unmarshal(jsonBytes, &payload); err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to unmarshal decoded data: %w", err)
	}

	return payload, nil
}

func ConvertToAbi(payload string) (string, error) {
	log.Printf("payload: %s", payload)
	return "", fmt.Errorf("not implemented")
}

func ConvertToPacked(payload string) (string, error) {
	decoded, err := DecodeABIPayload(payload)
	if err != nil {
		return "", fmt.Errorf("failed to decode ABI payload: %w", err)
	}
	if decoded.Kind != KindTransactions {
		return "", fmt.Errorf("conversion to packed only implemented for call payloads")
	}
	packed, err := Encode(decoded, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encode to packed: %w", err)
	}
	return hexutil.Encode(packed), nil
}

func ConvertToJson(payload string) (string, error) {
	decoded, err := DecodeABIPayload(payload)
	if err != nil {
		return "", fmt.Errorf("failed to decode ABI payload: %w", err)
	}
	jsonBytes, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal to JSON: %w", err)
	}
	return string(jsonBytes), nil
}
