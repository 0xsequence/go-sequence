package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/spf13/cobra"
)

type JsonRpcRequest struct {
	JsonRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      interface{}     `json:"id,omitempty"`
}

type JsonRpcSuccessResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      interface{} `json:"id,omitempty"`
}

type JsonRpcErrorResponse struct {
	JsonRPC string       `json:"jsonrpc"`
	Error   JsonRpcError `json:"error"`
	ID      interface{}  `json:"id,omitempty"`
}

type JsonRpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type AddressCalculateParams struct {
	ImageHash    string `json:"imageHash"`
	Factory      string `json:"factory"`
	Module       string `json:"module"`
	CreationCode string `json:"creationCode,omitempty"`
}

type ConfigNewParams struct {
	Threshold  string `json:"threshold"`
	Checkpoint string `json:"checkpoint"`
	From       string `json:"from"`
	Content    string `json:"content"`
}

type ConfigImageHashParams struct {
	Input json.RawMessage `json:"input"`
}

func successResponse(id interface{}, result interface{}) JsonRpcSuccessResponse {
	return JsonRpcSuccessResponse{
		JsonRPC: "2.0",
		ID:      id,
		Result:  result,
	}
}

func errorResponse(id interface{}, code int, message string, data interface{}) JsonRpcErrorResponse {
	return JsonRpcErrorResponse{
		JsonRPC: "2.0",
		ID:      id,
		Error: JsonRpcError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
}

func handleAddressCalculate(params json.RawMessage) (interface{}, error) {
	var p AddressCalculateParams
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	imageHash := common.HexToHash(p.ImageHash)
	factory := common.HexToAddress(p.Factory)
	module := common.HexToAddress(p.Module)

	creationCode := p.CreationCode
	if creationCode == "" {
		creationCode = contracts.DEFAULT_CREATION_CODE
	}

	address, err := contracts.GetCounterfactualAddress(factory, module, imageHash, creationCode)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate address: %w", err)
	}

	return address.Hex(), nil
}

func handleConfigNew(params json.RawMessage) (interface{}, error) {
	var p ConfigNewParams
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	threshold, ok := new(big.Int).SetString(p.Threshold, 10)
	if !ok {
		return nil, fmt.Errorf("invalid threshold: %s", p.Threshold)
	}
	if !threshold.IsUint64() || threshold.Uint64() > uint64(^uint16(0)) {
		return nil, fmt.Errorf("threshold too large: %s", p.Threshold)
	}

	checkpoint, ok := new(big.Int).SetString(p.Checkpoint, 10)
	if !ok {
		return nil, fmt.Errorf("invalid checkpoint: %s", p.Checkpoint)
	}
	if !checkpoint.IsUint64() {
		return nil, fmt.Errorf("checkpoint too large: %s", p.Checkpoint)
	}

	if p.From != "flat" {
		return nil, fmt.Errorf("unsupported 'from' format: %s", p.From)
	}

	log.Printf("Creating config with threshold=%s checkpoint=%s content=%q", p.Threshold, p.Checkpoint, p.Content)
	config, err := createConfig(uint16(threshold.Uint64()), uint64(checkpoint.Uint64()), []string{p.Content})
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %w", err)
	}

	log.Printf("Created config: %s", string(jsonConfig))

	return string(jsonConfig), nil
}

func handleConfigImageHash(params json.RawMessage) (interface{}, error) {
	var p ConfigImageHashParams
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	var rawConfig map[string]interface{}
	if err := json.Unmarshal(p.Input, &rawConfig); err != nil {
		return nil, fmt.Errorf("failed to parse raw config: %w", err)
	}

	config, err := v3.Core.DecodeWalletConfig(rawConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to decode wallet config: %w", err)
	}

	imageHash := config.ImageHash()

	return "0x" + common.Bytes2Hex(imageHash.Bytes()), nil
}

func handleSignatureEncode(params json.RawMessage) (interface{}, error) {
	var p SignatureEncodeParams
	if err := json.Unmarshal(params, &p); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}

	return signatureEncode(params)
}

func handleRPCRequest(w http.ResponseWriter, r *http.Request, debug bool, silent bool) {
	if !silent {
		log.Printf("[%s] %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
	}

	if r.Method != http.MethodPost || r.URL.Path != "/rpc" {
		if !silent {
			log.Printf("404 Not Found")
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var req JsonRpcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if !silent {
			log.Printf("JSON parse error: %v", err)
		}
		json.NewEncoder(w).Encode(errorResponse(nil, -32700, "Parse error", err.Error()))
		return
	}

	if debug && !silent {
		log.Printf("Request details: %+v", req)
	}

	if req.JsonRPC != "2.0" {
		json.NewEncoder(w).Encode(errorResponse(req.ID, -32600, "Invalid JSON-RPC version", nil))
		return
	}

	var result interface{}
	var err error

	switch req.Method {
	case "address_calculate":
		result, err = handleAddressCalculate(req.Params)
	case "config_new":
		result, err = handleConfigNew(req.Params)
	case "config_imageHash":
		result, err = handleConfigImageHash(req.Params)
	case "signature_encode":
		result, err = handleSignatureEncode(req.Params)
	default:
		json.NewEncoder(w).Encode(errorResponse(req.ID, -32601, fmt.Sprintf("Method not found: %s", req.Method), nil))
		return
	}

	if err != nil {
		json.NewEncoder(w).Encode(errorResponse(req.ID, -32000, err.Error(), nil))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successResponse(req.ID, result))
}

func startRPCServer(host string, port int, debug bool, silent bool) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	if !silent {
		log.Printf("RPC server running at http://%s/rpc", addr)
		if debug {
			log.Printf("Debug mode enabled - detailed logging active")
		}
	}

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		handleRPCRequest(w, r, debug, silent)
	})

	return http.ListenAndServe(addr, nil)
}

func newServerCmd() *cobra.Command {
	var (
		host   string
		port   int
		debug  bool
		silent bool
	)

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run a JSON-RPC server exposing CLI functionality",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startRPCServer(host, port, debug, silent)
		},
	}

	cmd.Flags().StringVar(&host, "host", "127.0.0.1", "Hostname to listen on")
	cmd.Flags().IntVar(&port, "port", 9998, "Port to listen on")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug logging")
	cmd.Flags().BoolVar(&silent, "silent", false, "Disable all logging output")

	return cmd
}
