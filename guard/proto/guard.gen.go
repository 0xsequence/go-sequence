// sequence-guard v0.4.0 d6b4a3c89539b494875af543fff459df65bb7b9e
// --
// Code generated by webrpc-gen@v0.24.0 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=guard.ridl -target=golang -pkg=proto -client -out=./clients/guard.gen.go
package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/0xsequence/go-sequence/lib/prototyp"
)

const WebrpcHeader = "Webrpc"

const WebrpcHeaderValue = "webrpc@v0.24.0;gen-golang@v0.18.2;sequence-guard@v0.4.0"

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.4.0"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "d6b4a3c89539b494875af543fff459df65bb7b9e"
}

type WebrpcGenVersions struct {
	WebrpcGenVersion string
	CodeGenName      string
	CodeGenVersion   string
	SchemaName       string
	SchemaVersion    string
}

func VersionFromHeader(h http.Header) (*WebrpcGenVersions, error) {
	if h.Get(WebrpcHeader) == "" {
		return nil, fmt.Errorf("header is empty or missing")
	}

	versions, err := parseWebrpcGenVersions(h.Get(WebrpcHeader))
	if err != nil {
		return nil, fmt.Errorf("webrpc header is invalid: %w", err)
	}

	return versions, nil
}

func parseWebrpcGenVersions(header string) (*WebrpcGenVersions, error) {
	versions := strings.Split(header, ";")
	if len(versions) < 3 {
		return nil, fmt.Errorf("expected at least 3 parts while parsing webrpc header: %v", header)
	}

	_, webrpcGenVersion, ok := strings.Cut(versions[0], "@")
	if !ok {
		return nil, fmt.Errorf("webrpc gen version could not be parsed from: %s", versions[0])
	}

	tmplTarget, tmplVersion, ok := strings.Cut(versions[1], "@")
	if !ok {
		return nil, fmt.Errorf("tmplTarget and tmplVersion could not be parsed from: %s", versions[1])
	}

	schemaName, schemaVersion, ok := strings.Cut(versions[2], "@")
	if !ok {
		return nil, fmt.Errorf("schema name and schema version could not be parsed from: %s", versions[2])
	}

	return &WebrpcGenVersions{
		WebrpcGenVersion: webrpcGenVersion,
		CodeGenName:      tmplTarget,
		CodeGenVersion:   tmplVersion,
		SchemaName:       schemaName,
		SchemaVersion:    schemaVersion,
	}, nil
}

//
// Common types
//

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	AppVersion    string `json:"appVersion"`
}

type RuntimeStatus struct {
	// overall status, true/false
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

type WalletSigner struct {
	Address string `json:"address"`
	Weight  uint8  `json:"weight"`
}

type SignRequest struct {
	// TODO: make this a string/BigInt TODO: with webrpc v2, we can import BigInt type, etc.
	ChainId uint64 `json:"chainId"`
	Msg     string `json:"msg"`
	AuxData string `json:"auxData"`
}

type OwnershipProof struct {
	Wallet    prototyp.Hash `json:"wallet"`
	Timestamp uint64        `json:"timestamp"`
	Signer    prototyp.Hash `json:"signer"`
	Signature prototyp.Hash `json:"signature"`
}

type AuthToken struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type RecoveryCode struct {
	Code string `json:"code"`
	Used bool   `json:"used"`
}

var methods = map[string]method{
	"/rpc/Guard/Ping": {
		Name:        "Ping",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/Version": {
		Name:        "Version",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/RuntimeStatus": {
		Name:        "RuntimeStatus",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/GetSignerConfig": {
		Name:        "GetSignerConfig",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/Sign": {
		Name:        "Sign",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/SignWith": {
		Name:        "SignWith",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/Patch": {
		Name:        "Patch",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/AuthMethods": {
		Name:        "AuthMethods",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/SetPIN": {
		Name:        "SetPIN",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/ResetPIN": {
		Name:        "ResetPIN",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/CreateTOTP": {
		Name:        "CreateTOTP",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/CommitTOTP": {
		Name:        "CommitTOTP",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/ResetTOTP": {
		Name:        "ResetTOTP",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/Reset2FA": {
		Name:        "Reset2FA",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/RecoveryCodes": {
		Name:        "RecoveryCodes",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
	"/rpc/Guard/ResetRecoveryCodes": {
		Name:        "ResetRecoveryCodes",
		Service:     "Guard",
		Annotations: map[string]string{},
	},
}

func WebrpcMethods() map[string]method {
	res := make(map[string]method, len(methods))
	for k, v := range methods {
		res[k] = v
	}

	return res
}

var WebRPCServices = map[string][]string{
	"Guard": {
		"Ping",
		"Version",
		"RuntimeStatus",
		"GetSignerConfig",
		"Sign",
		"SignWith",
		"Patch",
		"AuthMethods",
		"SetPIN",
		"ResetPIN",
		"CreateTOTP",
		"CommitTOTP",
		"ResetTOTP",
		"Reset2FA",
		"RecoveryCodes",
		"ResetRecoveryCodes",
	},
}

//
// Server types
//

type Guard interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetSignerConfig(ctx context.Context, signer string) (*WalletConfig, error)
	// Called by sequence.app when the user signs in, and signs messages/transactions/migrations.
	// Requires a valid 2FA token if enabled.
	//
	Sign(ctx context.Context, request *SignRequest, token *AuthToken) (string, error)
	SignWith(ctx context.Context, signer string, request *SignRequest, token *AuthToken) (string, error)
	// Internal use only.
	// Only ever needs to be called once per chain.
	// Signs a preconfigured payload that the caller has no control over.
	//
	Patch(ctx context.Context, signer string, chainId uint64, secret string) (interface{}, error)
	// Called by sequence.app when it needs to check the user's 2FA.
	// This happens during sign in, before signing messages and transactions, and when configuring 2FA.
	// Requires either a valid JWT or a signature by one of the wallet's signers.
	//
	AuthMethods(ctx context.Context, proof *OwnershipProof) ([]string, bool, error)
	// Not currently called.
	// Requires both a JWT and a wallet signature.
	//
	SetPIN(ctx context.Context, pin string, timestamp uint64, signature string) error
	ResetPIN(ctx context.Context, timestamp uint64, signature string) error
	// Called by sequence.app when the user configures their 2FA.
	// Requires both a JWT and a wallet signature.
	//
	CreateTOTP(ctx context.Context, timestamp uint64, signature string) (string, error)
	CommitTOTP(ctx context.Context, token string) ([]*RecoveryCode, error)
	ResetTOTP(ctx context.Context, timestamp uint64, signature string) error
	// Called by sequence.app when the user uses a recovery code.
	// Requires either a valid JWT or a signature by one of the wallet's signers.
	//
	Reset2FA(ctx context.Context, code string, proof *OwnershipProof) error
	// Called by sequence.app when the user is viewing their recovery codes.
	// Requires both a JWT and a wallet signature.
	//
	RecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error)
	ResetRecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error)
}

//
// Client types
//

type GuardClient interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetSignerConfig(ctx context.Context, signer string) (*WalletConfig, error)
	// Called by sequence.app when the user signs in, and signs messages/transactions/migrations.
	// Requires a valid 2FA token if enabled.
	//
	Sign(ctx context.Context, request *SignRequest, token *AuthToken) (string, error)
	SignWith(ctx context.Context, signer string, request *SignRequest, token *AuthToken) (string, error)
	// Internal use only.
	// Only ever needs to be called once per chain.
	// Signs a preconfigured payload that the caller has no control over.
	//
	Patch(ctx context.Context, signer string, chainId uint64, secret string) (interface{}, error)
	// Called by sequence.app when it needs to check the user's 2FA.
	// This happens during sign in, before signing messages and transactions, and when configuring 2FA.
	// Requires either a valid JWT or a signature by one of the wallet's signers.
	//
	AuthMethods(ctx context.Context, proof *OwnershipProof) ([]string, bool, error)
	// Not currently called.
	// Requires both a JWT and a wallet signature.
	//
	SetPIN(ctx context.Context, pin string, timestamp uint64, signature string) error
	ResetPIN(ctx context.Context, timestamp uint64, signature string) error
	// Called by sequence.app when the user configures their 2FA.
	// Requires both a JWT and a wallet signature.
	//
	CreateTOTP(ctx context.Context, timestamp uint64, signature string) (string, error)
	CommitTOTP(ctx context.Context, token string) ([]*RecoveryCode, error)
	ResetTOTP(ctx context.Context, timestamp uint64, signature string) error
	// Called by sequence.app when the user uses a recovery code.
	// Requires either a valid JWT or a signature by one of the wallet's signers.
	//
	Reset2FA(ctx context.Context, code string, proof *OwnershipProof) error
	// Called by sequence.app when the user is viewing their recovery codes.
	// Requires both a JWT and a wallet signature.
	//
	RecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error)
	ResetRecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error)
}

//
// Client
//

const GuardPathPrefix = "/rpc/Guard/"

type guardClient struct {
	client HTTPClient
	urls   [16]string
}

func NewGuardClient(addr string, client HTTPClient) GuardClient {
	prefix := urlBase(addr) + GuardPathPrefix
	urls := [16]string{
		prefix + "Ping",
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "GetSignerConfig",
		prefix + "Sign",
		prefix + "SignWith",
		prefix + "Patch",
		prefix + "AuthMethods",
		prefix + "SetPIN",
		prefix + "ResetPIN",
		prefix + "CreateTOTP",
		prefix + "CommitTOTP",
		prefix + "ResetTOTP",
		prefix + "Reset2FA",
		prefix + "RecoveryCodes",
		prefix + "ResetRecoveryCodes",
	}
	return &guardClient{
		client: client,
		urls:   urls,
	}
}

func (c *guardClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], nil, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) GetSignerConfig(ctx context.Context, signer string) (*WalletConfig, error) {
	in := struct {
		Arg0 string `json:"signer"`
	}{signer}
	out := struct {
		Ret0 *WalletConfig `json:"signerConfig"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) Sign(ctx context.Context, request *SignRequest, token *AuthToken) (string, error) {
	in := struct {
		Arg0 *SignRequest `json:"request"`
		Arg1 *AuthToken   `json:"token"`
	}{request, token}
	out := struct {
		Ret0 string `json:"sig"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[4], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) SignWith(ctx context.Context, signer string, request *SignRequest, token *AuthToken) (string, error) {
	in := struct {
		Arg0 string       `json:"signer"`
		Arg1 *SignRequest `json:"request"`
		Arg2 *AuthToken   `json:"token"`
	}{signer, request, token}
	out := struct {
		Ret0 string `json:"sig"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[5], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) Patch(ctx context.Context, signer string, chainId uint64, secret string) (interface{}, error) {
	in := struct {
		Arg0 string `json:"signer"`
		Arg1 uint64 `json:"chainId"`
		Arg2 string `json:"secret"`
	}{signer, chainId, secret}
	out := struct {
		Ret0 interface{} `json:"txs"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[6], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) AuthMethods(ctx context.Context, proof *OwnershipProof) ([]string, bool, error) {
	in := struct {
		Arg0 *OwnershipProof `json:"proof"`
	}{proof}
	out := struct {
		Ret0 []string `json:"methods"`
		Ret1 bool     `json:"active"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[7], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *guardClient) SetPIN(ctx context.Context, pin string, timestamp uint64, signature string) error {
	in := struct {
		Arg0 string `json:"pin"`
		Arg1 uint64 `json:"timestamp"`
		Arg2 string `json:"signature"`
	}{pin, timestamp, signature}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[8], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *guardClient) ResetPIN(ctx context.Context, timestamp uint64, signature string) error {
	in := struct {
		Arg0 uint64 `json:"timestamp"`
		Arg1 string `json:"signature"`
	}{timestamp, signature}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[9], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *guardClient) CreateTOTP(ctx context.Context, timestamp uint64, signature string) (string, error) {
	in := struct {
		Arg0 uint64 `json:"timestamp"`
		Arg1 string `json:"signature"`
	}{timestamp, signature}
	out := struct {
		Ret0 string `json:"uri"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[10], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) CommitTOTP(ctx context.Context, token string) ([]*RecoveryCode, error) {
	in := struct {
		Arg0 string `json:"token"`
	}{token}
	out := struct {
		Ret0 []*RecoveryCode `json:"codes"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[11], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) ResetTOTP(ctx context.Context, timestamp uint64, signature string) error {
	in := struct {
		Arg0 uint64 `json:"timestamp"`
		Arg1 string `json:"signature"`
	}{timestamp, signature}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[12], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *guardClient) Reset2FA(ctx context.Context, code string, proof *OwnershipProof) error {
	in := struct {
		Arg0 string          `json:"code"`
		Arg1 *OwnershipProof `json:"proof"`
	}{code, proof}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[13], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *guardClient) RecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error) {
	in := struct {
		Arg0 uint64 `json:"timestamp"`
		Arg1 string `json:"signature"`
	}{timestamp, signature}
	out := struct {
		Ret0 []*RecoveryCode `json:"codes"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[14], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, err
}

func (c *guardClient) ResetRecoveryCodes(ctx context.Context, timestamp uint64, signature string) ([]*RecoveryCode, error) {
	in := struct {
		Arg0 uint64 `json:"timestamp"`
		Arg1 string `json:"signature"`
	}{timestamp, signature}
	out := struct {
		Ret0 []*RecoveryCode `json:"codes"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[15], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

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
	req, err := http.NewRequestWithContext(ctx, "POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set(WebrpcHeader, WebrpcHeaderValue)
	if headers, ok := HTTPRequestHeaders(ctx); ok {
		for k := range headers {
			for _, v := range headers[k] {
				req.Header.Add(k, v)
			}
		}
	}
	return req, nil
}

// doHTTPRequest is common code to make a request to the remote service.
func doHTTPRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) (*http.Response, error) {
	reqBody, err := json.Marshal(in)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("failed to marshal JSON body: %w", err)
	}
	if err = ctx.Err(); err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("aborted because context was done: %w", err)
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCausef("could not build request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrWebrpcRequestFailed.WithCause(err)
	}

	if resp.StatusCode != 200 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to read server error response body: %w", err)
		}

		var rpcErr WebRPCError
		if err := json.Unmarshal(respBody, &rpcErr); err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to unmarshal server error: %w", err)
		}
		if rpcErr.Cause != "" {
			rpcErr.cause = errors.New(rpcErr.Cause)
		}
		return nil, rpcErr
	}

	if out != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to read response body: %w", err)
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return nil, ErrWebrpcBadResponse.WithCausef("failed to unmarshal JSON response body: %w", err)
		}
	}

	return resp, nil
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

type method struct {
	Name        string
	Service     string
	Annotations map[string]string
}

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

func MethodCtx(ctx context.Context) (method, bool) {
	req := RequestFromContext(ctx)
	if req == nil {
		return method{}, false
	}

	m, ok := methods[req.URL.Path]
	if !ok {
		return method{}, false
	}

	return m, true
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

// Schema errors
var (
	ErrUnauthorized     = WebRPCError{Code: 1000, Name: "Unauthorized", Message: "Unauthorized access", HTTPStatus: 401}
	ErrPermissionDenied = WebRPCError{Code: 1001, Name: "PermissionDenied", Message: "Permission denied", HTTPStatus: 403}
	ErrSessionExpired   = WebRPCError{Code: 1002, Name: "SessionExpired", Message: "Session expired", HTTPStatus: 403}
	ErrMethodNotFound   = WebRPCError{Code: 1003, Name: "MethodNotFound", Message: "Method not found", HTTPStatus: 404}
	ErrRequestConflict  = WebRPCError{Code: 1004, Name: "RequestConflict", Message: "Conflict with target resource", HTTPStatus: 409}
	ErrAborted          = WebRPCError{Code: 1005, Name: "Aborted", Message: "Request aborted", HTTPStatus: 400}
	ErrGeoblocked       = WebRPCError{Code: 1006, Name: "Geoblocked", Message: "Geoblocked region", HTTPStatus: 451}
	ErrRateLimited      = WebRPCError{Code: 1007, Name: "RateLimited", Message: "Rate-limited. Please slow down.", HTTPStatus: 429}
	ErrInvalidArgument  = WebRPCError{Code: 2001, Name: "InvalidArgument", Message: "Invalid argument", HTTPStatus: 400}
	ErrUnavailable      = WebRPCError{Code: 2002, Name: "Unavailable", Message: "Unavailable resource", HTTPStatus: 400}
	ErrQueryFailed      = WebRPCError{Code: 2003, Name: "QueryFailed", Message: "Query failed", HTTPStatus: 400}
	ErrValidationFailed = WebRPCError{Code: 2004, Name: "ValidationFailed", Message: "Validation Failed", HTTPStatus: 422}
	ErrNotFound         = WebRPCError{Code: 3000, Name: "NotFound", Message: "Resource not found", HTTPStatus: 400}
)
