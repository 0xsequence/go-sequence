// sequence-waas-intents v0.1.0 a9a0421e25cc87ddc15682e9e30e2305371486d0
// --
// Code generated by webrpc-gen@v0.19.3 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=intent.ridl -target=golang -pkg=intents -out=./intent.gen.go
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
	return "a9a0421e25cc87ddc15682e9e30e2305371486d0"
}

//
// Common types
//

type IntentName string

const (
	IntentName_initiateAuth          IntentName = "initiateAuth"
	IntentName_openSession           IntentName = "openSession"
	IntentName_closeSession          IntentName = "closeSession"
	IntentName_validateSession       IntentName = "validateSession"
	IntentName_finishValidateSession IntentName = "finishValidateSession"
	IntentName_listSessions          IntentName = "listSessions"
	IntentName_getSession            IntentName = "getSession"
	IntentName_sessionAuthProof      IntentName = "sessionAuthProof"
	IntentName_feeOptions            IntentName = "feeOptions"
	IntentName_signMessage           IntentName = "signMessage"
	IntentName_sendTransaction       IntentName = "sendTransaction"
	IntentName_getTransactionReceipt IntentName = "getTransactionReceipt"
	IntentName_federateAccount       IntentName = "federateAccount"
	IntentName_removeAccount         IntentName = "removeAccount"
	IntentName_listAccounts          IntentName = "listAccounts"
	IntentName_getIdToken            IntentName = "getIdToken"
)

func (x IntentName) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *IntentName) UnmarshalText(b []byte) error {
	*x = IntentName(string(b))
	return nil
}

func (x *IntentName) Is(values ...IntentName) bool {
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

type TransactionType string

const (
	TransactionType_transaction   TransactionType = "transaction"
	TransactionType_erc20send     TransactionType = "erc20send"
	TransactionType_erc721send    TransactionType = "erc721send"
	TransactionType_erc1155send   TransactionType = "erc1155send"
	TransactionType_delayedEncode TransactionType = "delayedEncode"
	TransactionType_contractCall  TransactionType = "contractCall"
)

func (x TransactionType) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *TransactionType) UnmarshalText(b []byte) error {
	*x = TransactionType(string(b))
	return nil
}

func (x *TransactionType) Is(values ...TransactionType) bool {
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

type IntentResponseCode string

const (
	IntentResponseCode_authInitiated      IntentResponseCode = "authInitiated"
	IntentResponseCode_sessionOpened      IntentResponseCode = "sessionOpened"
	IntentResponseCode_sessionClosed      IntentResponseCode = "sessionClosed"
	IntentResponseCode_sessionList        IntentResponseCode = "sessionList"
	IntentResponseCode_validationRequired IntentResponseCode = "validationRequired"
	IntentResponseCode_validationStarted  IntentResponseCode = "validationStarted"
	IntentResponseCode_validationFinished IntentResponseCode = "validationFinished"
	IntentResponseCode_sessionAuthProof   IntentResponseCode = "sessionAuthProof"
	IntentResponseCode_signedMessage      IntentResponseCode = "signedMessage"
	IntentResponseCode_feeOptions         IntentResponseCode = "feeOptions"
	IntentResponseCode_transactionReceipt IntentResponseCode = "transactionReceipt"
	IntentResponseCode_transactionFailed  IntentResponseCode = "transactionFailed"
	IntentResponseCode_getSessionResponse IntentResponseCode = "getSessionResponse"
	IntentResponseCode_accountList        IntentResponseCode = "accountList"
	IntentResponseCode_accountFederated   IntentResponseCode = "accountFederated"
	IntentResponseCode_accountRemoved     IntentResponseCode = "accountRemoved"
	IntentResponseCode_idToken            IntentResponseCode = "idToken"
)

func (x IntentResponseCode) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *IntentResponseCode) UnmarshalText(b []byte) error {
	*x = IntentResponseCode(string(b))
	return nil
}

func (x *IntentResponseCode) Is(values ...IntentResponseCode) bool {
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

type FeeTokenType uint32

const (
	FeeTokenType_unknown      FeeTokenType = 0
	FeeTokenType_erc20Token   FeeTokenType = 1
	FeeTokenType_erc1155Token FeeTokenType = 2
)

var FeeTokenType_name = map[uint32]string{
	0: "unknown",
	1: "erc20Token",
	2: "erc1155Token",
}

var FeeTokenType_value = map[string]uint32{
	"unknown":      0,
	"erc20Token":   1,
	"erc1155Token": 2,
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

type IdentityType string

const (
	IdentityType_None    IdentityType = "None"
	IdentityType_Guest   IdentityType = "Guest"
	IdentityType_OIDC    IdentityType = "OIDC"
	IdentityType_Email   IdentityType = "Email"
	IdentityType_PlayFab IdentityType = "PlayFab"
	IdentityType_Stytch  IdentityType = "Stytch"
)

func (x IdentityType) MarshalText() ([]byte, error) {
	return []byte(x), nil
}

func (x *IdentityType) UnmarshalText(b []byte) error {
	*x = IdentityType(string(b))
	return nil
}

func (x *IdentityType) Is(values ...IdentityType) bool {
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
	Name       IntentName   `json:"name"`
	ExpiresAt  uint64       `json:"expiresAt"`
	IssuedAt   uint64       `json:"issuedAt"`
	Data       interface{}  `json:"data"`
	Signatures []*Signature `json:"signatures,omitempty"`
}

type Signature struct {
	SessionID string `json:"sessionId"`
	Signature string `json:"signature"`
}

type IntentDataInitiateAuth struct {
	SessionID    string       `json:"sessionId"`
	IdentityType IdentityType `json:"identityType"`
	Verifier     string       `json:"verifier"`
	Metadata     *string      `json:"metadata,omitempty"`
}

type IntentDataOpenSession struct {
	SessionID          string       `json:"sessionId"`
	IdentityType       IdentityType `json:"identityType,omitempty"`
	Verifier           string       `json:"verifier,omitempty"`
	Answer             string       `json:"answer,omitempty"`
	ForceCreateAccount bool         `json:"forceCreateAccount,omitempty"`
	// Deprecated
	Email *string `json:"email,omitempty"`
	// Deprecated
	IdToken *string `json:"idToken,omitempty"`
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

type IntentDataSessionAuthProof struct {
	Network string  `json:"network"`
	Wallet  string  `json:"wallet"`
	Nonce   *string `json:"nonce"`
}

type IntentDataSignMessage struct {
	Network string `json:"network"`
	Wallet  string `json:"wallet"`
	Message string `json:"message"`
}

type IntentDataFeeOptions struct {
	Network string `json:"network"`
	Wallet  string `json:"wallet"`
	// is used to generate nonce space
	Identifier   string            `json:"identifier"`
	Transactions []json.RawMessage `json:"transactions"`
}

type IntentDataSendTransaction struct {
	Network string `json:"network"`
	Wallet  string `json:"wallet"`
	// is used to generate nonce space
	Identifier           string            `json:"identifier"`
	Transactions         []json.RawMessage `json:"transactions"`
	TransactionsFeeQuote *string           `json:"transactionsFeeQuote,omitempty"`
}

type IntentDataGetTransactionReceipt struct {
	Network    string `json:"network"`
	Wallet     string `json:"wallet"`
	MetaTxHash string `json:"metaTxHash"`
}

type IntentDataFederateAccount struct {
	SessionID    string       `json:"sessionId"`
	Wallet       string       `json:"wallet"`
	IdentityType IdentityType `json:"identityType"`
	Verifier     string       `json:"verifier,omitempty"`
	Answer       string       `json:"answer,omitempty"`
}

type IntentDataListAccounts struct {
	Wallet string `json:"wallet"`
}

type IntentDataRemoveAccount struct {
	Wallet    string `json:"wallet"`
	AccountID string `json:"accountId"`
}

type IntentDataGetIdToken struct {
	SessionID string `json:"sessionId"`
	Wallet    string `json:"wallet"`
	Nonce     string `json:"nonce"`
}

type TransactionRaw struct {
	Type  string  `json:"type"`
	To    string  `json:"to"`
	Value *string `json:"value"`
	Data  string  `json:"data"`
}

type AbiData struct {
	Abi  string        `json:"abi"`
	Func *string       `json:"func"`
	Args []interface{} `json:"args"`
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

// NOTE: TransactionDelayedEncode is deprecated, please use TransactionContractCall instead
type TransactionDelayedEncode struct {
	Type  string          `json:"type"`
	To    string          `json:"to"`
	Value string          `json:"value"`
	Data  json.RawMessage `json:"data"`
}

type TransactionContractCall struct {
	Type  string  `json:"type"`
	To    string  `json:"to"`
	Value *string `json:"value"`
	Data  AbiData `json:"data"`
}

type TransactionERC1155 struct {
	Type         string                     `json:"type"`
	TokenAddress string                     `json:"tokenAddress"`
	To           string                     `json:"to"`
	Vals         []*TransactionERC1155Value `json:"vals"`
	Data         *string                    `json:"data"`
}

type IntentResponse struct {
	Code IntentResponseCode `json:"code"`
	Data interface{}        `json:"data"`
}

type IntentResponseAuthInitiated struct {
	SessionID    string       `json:"sessionId"`
	IdentityType IdentityType `json:"identityType"`
	ExpiresIn    int          `json:"expiresIn"`
	Challenge    *string      `json:"challenge"`
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

type IntentResponseSessionAuthProof struct {
	SessionID string `json:"sessionId"`
	Network   string `json:"network"`
	Wallet    string `json:"wallet"`
	// The message contents: “SessionAuthProof <sessionId> <wallet> <nonce?>” hex encoded
	Message   string `json:"message"`
	Signature string `json:"signature"`
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
	ChainId         uint64       `json:"chainId"`
	Name            string       `json:"name"`
	Symbol          string       `json:"symbol"`
	Type            FeeTokenType `json:"type"`
	Decimals        *uint32      `json:"decimals"`
	LogoURL         string       `json:"logoURL"`
	ContractAddress *string      `json:"contractAddress"`
	TokenID         *string      `json:"tokenID"`
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

type IntentResponseAccountList struct {
	Accounts         []*Account `json:"accounts"`
	CurrentAccountID string     `json:"currentAccountId"`
}

type IntentResponseAccountFederated struct {
	Account *Account `json:"account"`
}

type IntentResponseAccountRemoved struct {
}

type IntentResponseIdToken struct {
	IdToken   string `json:"idToken"`
	ExpiresIn int    `json:"expiresIn"`
}

type Account struct {
	ID     string       `json:"id"`
	Type   IdentityType `json:"type"`
	Issuer *string      `json:"issuer,omitempty"`
	Email  *string      `json:"email,omitempty"`
}

var WebRPCServices = map[string][]string{}

//
// Server types
//

//
// Client types
//

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
	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

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
	if target == nil {
		return false
	}
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

func (e WebRPCError) WithCausef(format string, args ...interface{}) WebRPCError {
	cause := fmt.Errorf(format, args...)
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
