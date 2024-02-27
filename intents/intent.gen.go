// sequence-waas-intents v0.1.0 f671856a3104a7a8f1ab8957e5542b7f08a20fb8
// --
// Code generated by webrpc-gen@v0.14.0-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=intent.ridl -target=golang -pkg=intents -client -out=./intent.gen.go
package intents

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.1.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "f671856a3104a7a8f1ab8957e5542b7f08a20fb8"
}

//
// Types
//

type FeeTokenType uint32

const (
	FeeTokenType_UNKNOWN       FeeTokenType = 0
	FeeTokenType_ERC20_TOKEN   FeeTokenType = 1
	FeeTokenType_ERC1155_TOKEN FeeTokenType = 2
)

var FeeTokenType_name = map[uint32]string{
	0: "UNKNOWN",
	1: "ERC20_TOKEN",
	2: "ERC1155_TOKEN",
}

var FeeTokenType_value = map[string]uint32{
	"UNKNOWN":       0,
	"ERC20_TOKEN":   1,
	"ERC1155_TOKEN": 2,
}

func (x FeeTokenType) String() string {
	return FeeTokenType_name[uint32(x)]
}

func (x FeeTokenType) MarshalText() ([]byte, error) {
	return []byte(FeeTokenType_name[uint32(x)]), nil
}

func (x *FeeTokenType) UnmarshalText(b []byte) error {
	*x = FeeTokenType(FeeTokenType_value[string(b)])
	return nil
}

func (x *FeeTokenType) Is(values ...FeeTokenType) bool {
	if x == nil {
		return false
	}
	for _, v := range values {
		if *x == v {
			return true
		}
	}
	return false
}

type Intent struct {
	Version    string       `json:"version"`
	Name       string       `json:"name"`
	ExpiresAt  uint64       `json:"expiresAt"`
	IssuedAt   uint64       `json:"issuedAt"`
	Data       interface{}  `json:"data"`
	Signatures []*Signature `json:"signatures,omitempty"`
}

type Signature struct {
	SessionID string `json:"sessionId"`
	Signature string `json:"signature"`
}

type IntentDataOpenSession struct {
	SessionID string  `json:"sessionId"`
	Email     *string `json:"email,omitempty"`
	IdToken   *string `json:"idToken,omitempty"`
}

type IntentDataCloseSession struct {
	SessionID string `json:"sessionId"`
}

type IntentDataValidateSession struct {
	SessionID      string `json:"sessionId"`
	Wallet         string `json:"wallet"`
	DeviceMetadata string `json:"deviceMetadata"`
}

type IntentDataFinishValidateSession struct {
	SessionID string `json:"sessionId"`
	Wallet    string `json:"wallet"`
	Salt      string `json:"salt"`
	Challenge string `json:"challenge"`
}

type IntentDataListSessions struct {
	Wallet string `json:"wallet"`
}

type IntentDataGetSession struct {
	SessionID string `json:"sessionId"`
	Wallet    string `json:"wallet"`
}

type IntentDataSignMessage struct {
	Network string `json:"network"`
	Wallet  string `json:"wallet"`
	Message string `json:"message"`
}

type IntentDataFeeOptions struct {
	Network      string            `json:"network"`
	Wallet       string            `json:"wallet"`
	Identifier   string            `json:"identifier"`
	Transactions []json.RawMessage `json:"transactions"`
}

type IntentDataSendTransaction struct {
	Network              string            `json:"network"`
	Wallet               string            `json:"wallet"`
	Identifier           string            `json:"identifier"`
	Transactions         []json.RawMessage `json:"transactions"`
	TransactionsFeeQuote *string           `json:"transactionsFeeQuote,omitempty"`
}

type TransactionRaw struct {
	Type  string `json:"type"`
	To    string `json:"to"`
	Value string `json:"value"`
	Data  string `json:"data"`
}

type TransactionERC20 struct {
	Type         string `json:"type"`
	TokenAddress string `json:"tokenAddress"`
	To           string `json:"to"`
	Value        string `json:"value"`
}

type TransactionERC721 struct {
	Type         string  `json:"type"`
	TokenAddress string  `json:"tokenAddress"`
	To           string  `json:"to"`
	Id           string  `json:"id"`
	Safe         *bool   `json:"safe"`
	Data         *string `json:"data"`
}

type TransactionERC1155Value struct {
	ID     string `json:"id"`
	Amount string `json:"amount"`
}

type TransactionDelayedEncode struct {
	Type  string          `json:"type"`
	To    string          `json:"to"`
	Value string          `json:"value"`
	Data  json.RawMessage `json:"data"`
}

type TransactionERC1155 struct {
	Type         string                     `json:"type"`
	TokenAddress string                     `json:"tokenAddress"`
	To           string                     `json:"to"`
	Vals         []*TransactionERC1155Value `json:"vals"`
	Data         *string                    `json:"data"`
}

type IntentResponse struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type IntentResponseSessionOpened struct {
	SessionID string `json:"sessionId"`
	Wallet    string `json:"wallet"`
}

type IntentResponseSessionClosed struct {
}

type IntentResponseValidateSession struct {
}

type IntentResponseValidationRequired struct {
	SessionID string `json:"sessionId"`
}

type IntentResponseValidationStarted struct {
	Salt string `json:"salt"`
}

type IntentResponseValidationFinished struct {
	IsValid bool `json:"isValid"`
}

type IntentResponseListSessions struct {
	Sessions []string `json:"sessions"`
}

type IntentResponseGetSession struct {
	SessionID string `json:"sessionId"`
	Wallet    string `json:"wallet"`
	Validated bool   `json:"validated"`
}

type IntentResponseSignedMessage struct {
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

type FeeOption struct {
	Token    *FeeToken `json:"token"`
	To       string    `json:"to"`
	Value    string    `json:"value"`
	GasLimit uint      `json:"gasLimit"`
}

type FeeToken struct {
	ChainId         uint64        `json:"chainId"`
	Name            string        `json:"name"`
	Symbol          string        `json:"symbol"`
	Type            *FeeTokenType `json:"type"`
	Decimals        *uint32       `json:"decimals"`
	LogoURL         string        `json:"logoURL"`
	ContractAddress *string       `json:"contractAddress"`
	TokenID         *string       `json:"tokenID"`
}

type IntentResponseFeeOptions struct {
	FeeOptions []*FeeOption `json:"feeOptions"`
	FeeQuote   *string      `json:"feeQuote,omitempty"`
}

type IntentResponseTransactionReceipt struct {
	Request       interface{} `json:"request"`
	TxHash        string      `json:"txHash"`
	MetaTxHash    string      `json:"metaTxHash"`
	Receipt       interface{} `json:"receipt"`
	NativeReceipt interface{} `json:"nativeReceipt"`
	Simulations   interface{} `json:"simulations"`
}

type IntentResponseTransactionFailed struct {
	Error       string      `json:"error"`
	Request     interface{} `json:"request"`
	Simulations interface{} `json:"simulations"`
}

//
// Helpers
//

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}

var (
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}
	HTTPRequestCtxKey              = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func ServiceNameFromContext(ctx context.Context) string {
	service, _ := ctx.Value(ServiceNameCtxKey).(string)
	return service
}

func MethodNameFromContext(ctx context.Context) string {
	method, _ := ctx.Value(MethodNameCtxKey).(string)
	return method
}

func RequestFromContext(ctx context.Context) *http.Request {
	r, _ := ctx.Value(HTTPRequestCtxKey).(*http.Request)
	return r
}

//
// Errors
//

type WebRPCError struct {
	Name       string `json:"error"`
	Code       int    `json:"code"`
	Message    string `json:"msg"`
	Cause      string `json:"cause,omitempty"`
	HTTPStatus int    `json:"status"`
	cause      error
}

var _ error = WebRPCError{}

func (e WebRPCError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s %d: %s: %v", e.Name, e.Code, e.Message, e.cause)
	}
	return fmt.Sprintf("%s %d: %s", e.Name, e.Code, e.Message)
}

func (e WebRPCError) Is(target error) bool {
	if rpcErr, ok := target.(WebRPCError); ok {
		return rpcErr.Code == e.Code
	}
	return errors.Is(e.cause, target)
}

func (e WebRPCError) Unwrap() error {
	return e.cause
}

func (e WebRPCError) WithCause(cause error) WebRPCError {
	err := e
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Deprecated: Use .WithCause() method on WebRPCError.
func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	return rpcErr.WithCause(cause)
}

// Webrpc errors
var (
	ErrWebrpcEndpoint           = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed      = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400}
	ErrWebrpcBadRoute           = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod          = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest         = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse        = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic        = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
	ErrWebrpcInternalError      = WebRPCError{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500}
	ErrWebrpcClientDisconnected = WebRPCError{Code: -8, Name: "WebrpcClientDisconnected", Message: "client disconnected", HTTPStatus: 400}
	ErrWebrpcStreamLost         = WebRPCError{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400}
	ErrWebrpcStreamFinished     = WebRPCError{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200}
)
