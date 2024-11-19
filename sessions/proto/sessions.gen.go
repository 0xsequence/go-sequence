// sessions v0.0.1 67f782e8acfe452f905078a7423ed5d27c6639a8
// --
// Code generated by webrpc-gen@v0.21.3 with golang generator. DO NOT EDIT.
//
// webrpc-gen -schema=sessions.ridl -target=golang -pkg=proto -client -out=./clients/sessions.gen.go
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

const WebrpcHeaderValue = "webrpc@v0.21.3;gen-golang@v0.16.2;sessions@v0.0.1"

// WebRPC description and code-gen version
func WebRPCVersion() string {
	return "v1"
}

// Schema version of your RIDL schema
func WebRPCSchemaVersion() string {
	return "v0.0.1"
}

// Schema hash generated from your RIDL schema
func WebRPCSchemaHash() string {
	return "67f782e8acfe452f905078a7423ed5d27c6639a8"
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

type SignatureType int32

const (
	SignatureType_EIP712  SignatureType = 0
	SignatureType_EthSign SignatureType = 1
	SignatureType_EIP1271 SignatureType = 2
)

var SignatureType_name = map[int32]string{
	0: "EIP712",
	1: "EthSign",
	2: "EIP1271",
}

var SignatureType_value = map[string]int32{
	"EIP712":  0,
	"EthSign": 1,
	"EIP1271": 2,
}

func (x SignatureType) String() string {
	return SignatureType_name[int32(x)]
}

func (x SignatureType) MarshalText() ([]byte, error) {
	return []byte(SignatureType_name[int32(x)]), nil
}

func (x *SignatureType) UnmarshalText(b []byte) error {
	*x = SignatureType(SignatureType_value[string(b)])
	return nil
}

func (x *SignatureType) Is(values ...SignatureType) bool {
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

type RuntimeStatus struct {
	Healthy bool           `json:"healthy"`
	Started time.Time      `json:"started"`
	Uptime  uint64         `json:"uptime"`
	Version string         `json:"version"`
	Branch  string         `json:"branch"`
	Commit  string         `json:"commit"`
	Arweave *ArweaveStatus `json:"arweave"`
}

type ArweaveStatus struct {
	NodeURL          string     `json:"nodeURL"`
	Namespace        string     `json:"namespace"`
	Sender           string     `json:"sender"`
	Signer           string     `json:"signer"`
	FlushInterval    string     `json:"flushInterval"`
	BundleDelay      string     `json:"bundleDelay"`
	BundleLimit      int        `json:"bundleLimit"`
	Confirmations    int        `json:"confirmations"`
	LockTTL          string     `json:"lockTTL"`
	Healthy          bool       `json:"healthy"`
	LastFlush        *time.Time `json:"lastFlush"`
	LastFlushSeconds *uint64    `json:"lastFlushSeconds"`
}

type Info struct {
	Wallets     map[string]uint64 `json:"wallets"`
	Configs     map[string]uint64 `json:"configs"`
	ConfigTrees uint64            `json:"configTrees"`
	Migrations  map[string]uint64 `json:"migrations"`
	Signatures  uint64            `json:"signatures"`
	Digests     uint64            `json:"digests"`
	Recorder    *RecorderInfo     `json:"recorder,omitempty"`
	Arweave     *ArweaveInfo      `json:"arweave,omitempty"`
}

type RecorderInfo struct {
	Requests         uint64            `json:"requests"`
	Buffer           uint64            `json:"buffer"`
	LastFlush        *time.Time        `json:"lastFlush"`
	LastFlushSeconds *uint64           `json:"lastFlushSeconds"`
	Endpoints        map[string]uint64 `json:"endpoints"`
}

type ArweaveInfo struct {
	NodeURL          string              `json:"nodeURL"`
	Namespace        string              `json:"namespace"`
	Sender           *ArweaveSenderInfo  `json:"sender"`
	Signer           string              `json:"signer"`
	FlushInterval    string              `json:"flushInterval"`
	BundleDelay      string              `json:"bundleDelay"`
	BundleLimit      int                 `json:"bundleLimit"`
	Confirmations    int                 `json:"confirmations"`
	LockTTL          string              `json:"lockTTL"`
	Healthy          bool                `json:"healthy"`
	LastFlush        *time.Time          `json:"lastFlush"`
	LastFlushSeconds *uint64             `json:"lastFlushSeconds"`
	Bundles          uint64              `json:"bundles"`
	Pending          *ArweavePendingInfo `json:"pending"`
}

type ArweaveSenderInfo struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

type ArweavePendingInfo struct {
	Wallets    uint64               `json:"wallets"`
	Configs    uint64               `json:"configs"`
	Migrations uint64               `json:"migrations"`
	Signatures uint64               `json:"signatures"`
	Bundles    []*ArweaveBundleInfo `json:"bundles"`
}

type ArweaveBundleInfo struct {
	Transaction   string    `json:"transaction"`
	Block         uint64    `json:"block"`
	Items         uint64    `json:"items"`
	SentAt        time.Time `json:"sentAt"`
	Confirmations uint64    `json:"confirmations"`
}

type Context struct {
	Version              int           `json:"version"`
	Factory              prototyp.Hash `json:"factory"`
	MainModule           prototyp.Hash `json:"mainModule"`
	MainModuleUpgradable prototyp.Hash `json:"mainModuleUpgradable"`
	GuestModule          prototyp.Hash `json:"guestModule"`
	WalletCreationCode   prototyp.Hash `json:"walletCreationCode"`
}

type Signature struct {
	Digest           prototyp.Hash      `json:"digest"`
	ToImageHash      prototyp.HashMaybe `json:"toImageHash,omitempty"`
	ChainID          prototyp.BigInt    `json:"chainID"`
	Type             SignatureType      `json:"type"`
	Signature        prototyp.Hash      `json:"signature"`
	ValidOnChain     *prototyp.BigInt   `json:"validOnChain,omitempty"`
	ValidOnBlock     *prototyp.BigInt   `json:"validOnBlock,omitempty"`
	ValidOnBlockHash prototyp.HashMaybe `json:"validOnBlockHash,omitempty"`
}

type SignerSignature struct {
	ReferenceChainID *string `json:"referenceChainID"`
	Signer           *string `json:"signer"`
	Signature        string  `json:"signature"`
}

type ConfigUpdate struct {
	ToImageHash prototyp.Hash `json:"toImageHash"`
	Signature   prototyp.Hash `json:"signature"`
}

type Transaction struct {
	To            prototyp.Hash      `json:"to"`
	Value         prototyp.BigInt    `json:"value,omitempty"`
	Data          prototyp.HashMaybe `json:"data,omitempty"`
	GasLimit      prototyp.BigInt    `json:"gasLimit,omitempty"`
	DelegateCall  *bool              `json:"delegateCall,omitempty"`
	RevertOnError *bool              `json:"revertOnError,omitempty"`
}

type TransactionBundle struct {
	Executor     prototyp.Hash   `json:"executor"`
	Transactions []*Transaction  `json:"transactions"`
	Nonce        prototyp.BigInt `json:"nonce"`
	Signature    prototyp.Hash   `json:"signature"`
}

var methods = map[string]method{
	"/rpc/Sessions/Ping": {
		Name:        "Ping",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/Config": {
		Name:        "Config",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/Wallets": {
		Name:        "Wallets",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/DeployHash": {
		Name:        "DeployHash",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/ConfigUpdates": {
		Name:        "ConfigUpdates",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/Migrations": {
		Name:        "Migrations",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveConfig": {
		Name:        "SaveConfig",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveWallet": {
		Name:        "SaveWallet",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveSignature": {
		Name:        "SaveSignature",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveSignerSignatures": {
		Name:        "SaveSignerSignatures",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveSignerSignatures2": {
		Name:        "SaveSignerSignatures2",
		Service:     "Sessions",
		Annotations: map[string]string{},
	},
	"/rpc/Sessions/SaveMigration": {
		Name:        "SaveMigration",
		Service:     "Sessions",
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
	"Sessions": {
		"Ping",
		"Config",
		"Wallets",
		"DeployHash",
		"ConfigUpdates",
		"Migrations",
		"SaveConfig",
		"SaveWallet",
		"SaveSignature",
		"SaveSignerSignatures",
		"SaveSignerSignatures2",
		"SaveMigration",
	},
}

//
// Server types
//

type Sessions interface {
	Ping(ctx context.Context) error
	Config(ctx context.Context, imageHash string) (int, interface{}, error)
	Wallets(ctx context.Context, signer string, cursor *uint64, limit *uint64) (map[string]*Signature, uint64, error)
	DeployHash(ctx context.Context, wallet string) (string, *Context, error)
	ConfigUpdates(ctx context.Context, wallet string, fromImageHash string, allUpdates *bool) ([]*ConfigUpdate, error)
	Migrations(ctx context.Context, wallet string, fromVersion int, fromImageHash string, chainID *string) (map[string]map[int]map[string]*TransactionBundle, error)
	SaveConfig(ctx context.Context, version int, config interface{}) error
	SaveWallet(ctx context.Context, version int, deployConfig interface{}) error
	SaveSignature(ctx context.Context, wallet string, digest string, chainID string, signature string, toConfig *interface{}, referenceChainID *string) error
	SaveSignerSignatures(ctx context.Context, wallet string, digest string, chainID string, signatures []string, toConfig *interface{}) error
	SaveSignerSignatures2(ctx context.Context, wallet string, digest string, chainID string, signatures []*SignerSignature, toConfig *interface{}) error
	SaveMigration(ctx context.Context, wallet string, fromVersion int, toVersion int, toConfig interface{}, executor string, transactions []*Transaction, nonce string, signature string, chainID *string) error
}

//
// Client types
//

type SessionsClient interface {
	Ping(ctx context.Context) error
	Config(ctx context.Context, imageHash string) (int, interface{}, error)
	Wallets(ctx context.Context, signer string, cursor *uint64, limit *uint64) (map[string]*Signature, uint64, error)
	DeployHash(ctx context.Context, wallet string) (string, *Context, error)
	ConfigUpdates(ctx context.Context, wallet string, fromImageHash string, allUpdates *bool) ([]*ConfigUpdate, error)
	Migrations(ctx context.Context, wallet string, fromVersion int, fromImageHash string, chainID *string) (map[string]map[int]map[string]*TransactionBundle, error)
	SaveConfig(ctx context.Context, version int, config interface{}) error
	SaveWallet(ctx context.Context, version int, deployConfig interface{}) error
	SaveSignature(ctx context.Context, wallet string, digest string, chainID string, signature string, toConfig *interface{}, referenceChainID *string) error
	SaveSignerSignatures(ctx context.Context, wallet string, digest string, chainID string, signatures []string, toConfig *interface{}) error
	SaveSignerSignatures2(ctx context.Context, wallet string, digest string, chainID string, signatures []*SignerSignature, toConfig *interface{}) error
	SaveMigration(ctx context.Context, wallet string, fromVersion int, toVersion int, toConfig interface{}, executor string, transactions []*Transaction, nonce string, signature string, chainID *string) error
}

//
// Client
//

const SessionsPathPrefix = "/rpc/Sessions/"

type sessionsClient struct {
	client HTTPClient
	urls   [12]string
}

func NewSessionsClient(addr string, client HTTPClient) SessionsClient {
	prefix := urlBase(addr) + SessionsPathPrefix
	urls := [12]string{
		prefix + "Ping",
		prefix + "Config",
		prefix + "Wallets",
		prefix + "DeployHash",
		prefix + "ConfigUpdates",
		prefix + "Migrations",
		prefix + "SaveConfig",
		prefix + "SaveWallet",
		prefix + "SaveSignature",
		prefix + "SaveSignerSignatures",
		prefix + "SaveSignerSignatures2",
		prefix + "SaveMigration",
	}
	return &sessionsClient{
		client: client,
		urls:   urls,
	}
}

func (c *sessionsClient) Ping(ctx context.Context) error {

	resp, err := doHTTPRequest(ctx, c.client, c.urls[0], nil, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) Config(ctx context.Context, imageHash string) (int, interface{}, error) {
	in := struct {
		Arg0 string `json:"imageHash"`
	}{imageHash}
	out := struct {
		Ret0 int         `json:"version"`
		Ret1 interface{} `json:"config"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[1], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *sessionsClient) Wallets(ctx context.Context, signer string, cursor *uint64, limit *uint64) (map[string]*Signature, uint64, error) {
	in := struct {
		Arg0 string  `json:"signer"`
		Arg1 *uint64 `json:"cursor"`
		Arg2 *uint64 `json:"limit"`
	}{signer, cursor, limit}
	out := struct {
		Ret0 map[string]*Signature `json:"wallets"`
		Ret1 uint64                `json:"cursor"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[2], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *sessionsClient) DeployHash(ctx context.Context, wallet string) (string, *Context, error) {
	in := struct {
		Arg0 string `json:"wallet"`
	}{wallet}
	out := struct {
		Ret0 string   `json:"deployHash"`
		Ret1 *Context `json:"context"`
	}{}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[3], in, &out)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return out.Ret0, out.Ret1, err
}

func (c *sessionsClient) ConfigUpdates(ctx context.Context, wallet string, fromImageHash string, allUpdates *bool) ([]*ConfigUpdate, error) {
	in := struct {
		Arg0 string `json:"wallet"`
		Arg1 string `json:"fromImageHash"`
		Arg2 *bool  `json:"allUpdates"`
	}{wallet, fromImageHash, allUpdates}
	out := struct {
		Ret0 []*ConfigUpdate `json:"updates"`
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

func (c *sessionsClient) Migrations(ctx context.Context, wallet string, fromVersion int, fromImageHash string, chainID *string) (map[string]map[int]map[string]*TransactionBundle, error) {
	in := struct {
		Arg0 string  `json:"wallet"`
		Arg1 int     `json:"fromVersion"`
		Arg2 string  `json:"fromImageHash"`
		Arg3 *string `json:"chainID"`
	}{wallet, fromVersion, fromImageHash, chainID}
	out := struct {
		Ret0 map[string]map[int]map[string]*TransactionBundle `json:"migrations"`
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

func (c *sessionsClient) SaveConfig(ctx context.Context, version int, config interface{}) error {
	in := struct {
		Arg0 int         `json:"version"`
		Arg1 interface{} `json:"config"`
	}{version, config}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[6], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) SaveWallet(ctx context.Context, version int, deployConfig interface{}) error {
	in := struct {
		Arg0 int         `json:"version"`
		Arg1 interface{} `json:"deployConfig"`
	}{version, deployConfig}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[7], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) SaveSignature(ctx context.Context, wallet string, digest string, chainID string, signature string, toConfig *interface{}, referenceChainID *string) error {
	in := struct {
		Arg0 string       `json:"wallet"`
		Arg1 string       `json:"digest"`
		Arg2 string       `json:"chainID"`
		Arg3 string       `json:"signature"`
		Arg4 *interface{} `json:"toConfig"`
		Arg5 *string      `json:"referenceChainID"`
	}{wallet, digest, chainID, signature, toConfig, referenceChainID}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[8], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) SaveSignerSignatures(ctx context.Context, wallet string, digest string, chainID string, signatures []string, toConfig *interface{}) error {
	in := struct {
		Arg0 string       `json:"wallet"`
		Arg1 string       `json:"digest"`
		Arg2 string       `json:"chainID"`
		Arg3 []string     `json:"signatures"`
		Arg4 *interface{} `json:"toConfig"`
	}{wallet, digest, chainID, signatures, toConfig}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[9], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) SaveSignerSignatures2(ctx context.Context, wallet string, digest string, chainID string, signatures []*SignerSignature, toConfig *interface{}) error {
	in := struct {
		Arg0 string             `json:"wallet"`
		Arg1 string             `json:"digest"`
		Arg2 string             `json:"chainID"`
		Arg3 []*SignerSignature `json:"signatures"`
		Arg4 *interface{}       `json:"toConfig"`
	}{wallet, digest, chainID, signatures, toConfig}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[10], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
}

func (c *sessionsClient) SaveMigration(ctx context.Context, wallet string, fromVersion int, toVersion int, toConfig interface{}, executor string, transactions []*Transaction, nonce string, signature string, chainID *string) error {
	in := struct {
		Arg0 string         `json:"wallet"`
		Arg1 int            `json:"fromVersion"`
		Arg2 int            `json:"toVersion"`
		Arg3 interface{}    `json:"toConfig"`
		Arg4 string         `json:"executor"`
		Arg5 []*Transaction `json:"transactions"`
		Arg6 string         `json:"nonce"`
		Arg7 string         `json:"signature"`
		Arg8 *string        `json:"chainID"`
	}{wallet, fromVersion, toVersion, toConfig, executor, transactions, nonce, signature, chainID}

	resp, err := doHTTPRequest(ctx, c.client, c.urls[11], in, nil)
	if resp != nil {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = ErrWebrpcRequestFailed.WithCausef("failed to close response body: %w", cerr)
		}
	}

	return err
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
	ErrInvalidArgument = WebRPCError{Code: 1, Name: "InvalidArgument", Message: "invalid argument", HTTPStatus: 400}
	ErrNotFound        = WebRPCError{Code: 2, Name: "NotFound", Message: "not found", HTTPStatus: 400}
)
