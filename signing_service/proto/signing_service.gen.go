// sequence-dapp-signing-server v0.1.0 eba943dd810a3108374b471782bdad9e2f1ba29f
// --
// Code generated by webrpc-gen@v0.12.x-dev with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=dapp_signing_server.ridl -target=golang -pkg=proto -client -out=./clients/dapp_signing_server.gen.go
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
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
	return "eba943dd810a3108374b471782bdad9e2f1ba29f"
}

//
// Types
//

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	AppVersion    string `json:"appVersion"`
}

type RuntimeStatus struct {
	HealthOK   bool      `json:"healthOK"`
	StartTime  time.Time `json:"startTime"`
	Uptime     uint64    `json:"uptime"`
	Ver        string    `json:"ver"`
	Branch     string    `json:"branch"`
	CommitHash string    `json:"commitHash"`
}

type WalletConfig struct {
	Address string `json:"address"`
	Content string `json:"content"`
}

type SignRequest struct {
	ChainId     uint64       `json:"chainId"`
	Msg         string       `json:"msg"`
	AuxData     string       `json:"auxData"`
	SignContext *SignContext `json:"signContext"`
}

type SignContext struct {
	ChainId             uint64  `json:"chainId"`
	PreImage            string  `json:"preImage"`
	WalletAddress       *string `json:"walletAddress"`
	SigneeWalletAddress string  `json:"signeeWalletAddress"`
	Signature           string  `json:"signature"`
	Message             *string `json:"message"`
	Transactions        *string `json:"transactions"`
	Nonce               *uint64 `json:"nonce"`
	PartnerId           *uint64 `json:"partnerId"`
}

type SigningService interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetSignerConfig(ctx context.Context, signer string) (*WalletConfig, error)
	Sign(ctx context.Context, request *SignRequest) (string, error)
	SignWith(ctx context.Context, signer string, request *SignRequest) (string, error)
}

var WebRPCServices = map[string][]string{
	"SigningService": {
		"Ping",
		"Version",
		"RuntimeStatus",
		"GetSignerConfig",
		"Sign",
		"SignWith",
	},
}

//
// Client
//

const SigningServicePathPrefix = "/rpc/SigningService/"

type signingServiceClient struct {
	client HTTPClient
	urls   [6]string
}

func NewSigningServiceClient(addr string, client HTTPClient) SigningService {
	prefix := urlBase(addr) + SigningServicePathPrefix
	urls := [6]string{
		prefix + "Ping",
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "GetSignerConfig",
		prefix + "Sign",
		prefix + "SignWith",
	}
	return &signingServiceClient{
		client: client,
		urls:   urls,
	}
}

func (c *signingServiceClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	return out.Ret0, err
}

func (c *signingServiceClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *signingServiceClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, err
}

func (c *signingServiceClient) GetSignerConfig(ctx context.Context, signer string) (*WalletConfig, error) {
	in := struct {
		Arg0 string `json:"signer"`
	}{signer}
	out := struct {
		Ret0 *WalletConfig `json:"signerConfig"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], in, &out)
	return out.Ret0, err
}

func (c *signingServiceClient) Sign(ctx context.Context, request *SignRequest) (string, error) {
	in := struct {
		Arg0 *SignRequest `json:"request"`
	}{request}
	out := struct {
		Ret0 string `json:"sig"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], in, &out)
	return out.Ret0, err
}

func (c *signingServiceClient) SignWith(ctx context.Context, signer string, request *SignRequest) (string, error) {
	in := struct {
		Arg0 string       `json:"signer"`
		Arg1 *SignRequest `json:"request"`
	}{signer, request}
	out := struct {
		Ret0 string `json:"sig"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], in, &out)
	return out.Ret0, err
}

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doJSONRequest is common code to make a request to the remote service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to marshal JSON body: %w", err))
	}
	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("could not build request: %w", err))
	}
	resp, err := client.Do(req)
	if err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("failed to close response body: %w", cerr))
		}
	}()

	if err = ctx.Err(); err != nil {
		return ErrorWithCause(ErrWebrpcRequestFailed, fmt.Errorf("aborted because context was done: %w", err))
	}

	if resp.StatusCode != 200 {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read server error response body: %w", err))
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal server error: %w", err))
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return rpcErr
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to read response body: %w", err))
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return ErrorWithCause(ErrWebrpcBadResponse, fmt.Errorf("failed to unmarshal JSON response body: %w", err))
		}
	}

	return nil
}

func WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {
	if _, ok := h["Accept"]; ok {
		return nil, errors.New("provided header cannot set Accept")
	}
	if _, ok := h["Content-Type"]; ok {
		return nil, errors.New("provided header cannot set Content-Type")
	}

	copied := make(http.Header, len(h))
	for k, vv := range h {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}

	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil
}

func HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {
	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)
	return h, ok
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
	// For Client
	HTTPClientRequestHeadersCtxKey = &contextKey{"HTTPClientRequestHeaders"}

	// For Server
	HTTPResponseWriterCtxKey = &contextKey{"HTTPResponseWriter"}

	HTTPRequestCtxKey = &contextKey{"HTTPRequest"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

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

func ErrorWithCause(rpcErr WebRPCError, cause error) WebRPCError {
	err := rpcErr
	err.cause = cause
	err.Cause = cause.Error()
	return err
}

// Webrpc errors
var (
	ErrWebrpcEndpoint      = WebRPCError{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400}
	ErrWebrpcRequestFailed = WebRPCError{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 0}
	ErrWebrpcBadRoute      = WebRPCError{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404}
	ErrWebrpcBadMethod     = WebRPCError{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405}
	ErrWebrpcBadRequest    = WebRPCError{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400}
	ErrWebrpcBadResponse   = WebRPCError{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500}
	ErrWebrpcServerPanic   = WebRPCError{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500}
)
