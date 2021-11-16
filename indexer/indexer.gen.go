// sequence-indexer v0.4.0 d44c148ecd843665eba71e66908382e569d8821a
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/golang
// Do not edit by hand. Update your webrpc schema and re-generate.
package indexer

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

	"github.com/0xsequence/go-sequence/lib/prototyp"
)

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
	return "d44c148ecd843665eba71e66908382e569d8821a"
}

//
// Types
//

type ContractType uint8

const (
	ContractType_UNKNOWN         ContractType = 0
	ContractType_ERC20           ContractType = 1
	ContractType_ERC721          ContractType = 2
	ContractType_ERC1155         ContractType = 3
	ContractType_SEQUENCE_WALLET ContractType = 4
	ContractType_ERC20_BRIDGE    ContractType = 5
	ContractType_ERC721_BRIDGE   ContractType = 6
	ContractType_ERC1155_BRIDGE  ContractType = 7
)

var ContractType_name = map[uint8]string{
	0: "UNKNOWN",
	1: "ERC20",
	2: "ERC721",
	3: "ERC1155",
	4: "SEQUENCE_WALLET",
	5: "ERC20_BRIDGE",
	6: "ERC721_BRIDGE",
	7: "ERC1155_BRIDGE",
}

var ContractType_value = map[string]uint8{
	"UNKNOWN":         0,
	"ERC20":           1,
	"ERC721":          2,
	"ERC1155":         3,
	"SEQUENCE_WALLET": 4,
	"ERC20_BRIDGE":    5,
	"ERC721_BRIDGE":   6,
	"ERC1155_BRIDGE":  7,
}

func (x ContractType) String() string {
	return ContractType_name[uint8(x)]
}

func (x ContractType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(ContractType_name[uint8(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *ContractType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = ContractType(ContractType_value[j])
	return nil
}

type EventLogType uint8

const (
	EventLogType_UNKNOWN       EventLogType = 0
	EventLogType_BLOCK_ADDED   EventLogType = 1
	EventLogType_BLOCK_REMOVED EventLogType = 2
)

var EventLogType_name = map[uint8]string{
	0: "UNKNOWN",
	1: "BLOCK_ADDED",
	2: "BLOCK_REMOVED",
}

var EventLogType_value = map[string]uint8{
	"UNKNOWN":       0,
	"BLOCK_ADDED":   1,
	"BLOCK_REMOVED": 2,
}

func (x EventLogType) String() string {
	return EventLogType_name[uint8(x)]
}

func (x EventLogType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(EventLogType_name[uint8(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *EventLogType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = EventLogType(EventLogType_value[j])
	return nil
}

type EventLogDataType uint8

const (
	EventLogDataType_UNKNOWN        EventLogDataType = 0
	EventLogDataType_TOKEN_TRANSFER EventLogDataType = 1
	EventLogDataType_SEQUENCE_TXN   EventLogDataType = 2
)

var EventLogDataType_name = map[uint8]string{
	0: "UNKNOWN",
	1: "TOKEN_TRANSFER",
	2: "SEQUENCE_TXN",
}

var EventLogDataType_value = map[string]uint8{
	"UNKNOWN":        0,
	"TOKEN_TRANSFER": 1,
	"SEQUENCE_TXN":   2,
}

func (x EventLogDataType) String() string {
	return EventLogDataType_name[uint8(x)]
}

func (x EventLogDataType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(EventLogDataType_name[uint8(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *EventLogDataType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = EventLogDataType(EventLogDataType_value[j])
	return nil
}

type TxnTransferType uint32

const (
	TxnTransferType_UNKNOWN TxnTransferType = 0
	TxnTransferType_SEND    TxnTransferType = 1
	TxnTransferType_RECEIVE TxnTransferType = 2
)

var TxnTransferType_name = map[uint32]string{
	0: "UNKNOWN",
	1: "SEND",
	2: "RECEIVE",
}

var TxnTransferType_value = map[string]uint32{
	"UNKNOWN": 0,
	"SEND":    1,
	"RECEIVE": 2,
}

func (x TxnTransferType) String() string {
	return TxnTransferType_name[uint32(x)]
}

func (x TxnTransferType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(TxnTransferType_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *TxnTransferType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = TxnTransferType(TxnTransferType_value[j])
	return nil
}

type SortOrder uint32

const (
	SortOrder_DESC SortOrder = 0
	SortOrder_ASC  SortOrder = 1
)

var SortOrder_name = map[uint32]string{
	0: "DESC",
	1: "ASC",
}

var SortOrder_value = map[string]uint32{
	"DESC": 0,
	"ASC":  1,
}

func (x SortOrder) String() string {
	return SortOrder_name[uint32(x)]
}

func (x SortOrder) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(SortOrder_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *SortOrder) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = SortOrder(SortOrder_value[j])
	return nil
}

type Version struct {
	WebrpcVersion string `json:"webrpcVersion"`
	SchemaVersion string `json:"schemaVersion"`
	SchemaHash    string `json:"schemaHash"`
	AppVersion    string `json:"appVersion"`
}

type RuntimeStatus struct {
	HealthOK       bool           `json:"healthOK"`
	IndexerEnabled bool           `json:"indexerEnabled"`
	StartTime      time.Time      `json:"startTime"`
	Uptime         uint64         `json:"uptime"`
	Ver            string         `json:"ver"`
	Branch         string         `json:"branch"`
	CommitHash     string         `json:"commitHash"`
	ChainID        uint64         `json:"chainID"`
	Checks         *RuntimeChecks `json:"checks"`
}

type RuntimeChecks struct {
	Running      bool   `json:"running"`
	SyncMode     string `json:"syncMode"`
	LastBlockNum uint64 `json:"lastBlockNum"`
}

type EtherBalance struct {
	AccountAddress prototyp.Hash   `json:"accountAddress"`
	BalanceWei     prototyp.BigInt `json:"balanceWei"`
}

type IndexState struct {
	ChainID      prototyp.BigInt `json:"chainId" db:"chain_id"`
	LastBlockNum uint64          `json:"lastBlockNum" db:"last_block_num"`
}

type EventLog struct {
	ID              uint64           `json:"id" db:"id,omitempty"`
	Type            EventLogType     `json:"type" db:"type"`
	BlockNumber     uint64           `json:"blockNumber" db:"block_num"`
	BlockHash       prototyp.Hash    `json:"blockHash" db:"block_hash"`
	ContractAddress prototyp.Hash    `json:"contractAddress" db:"contract_address"`
	ContractType    ContractType     `json:"contractType" db:"contract_type"`
	TxnHash         prototyp.Hash    `json:"txnHash" db:"txn_hash"`
	TxnIndex        uint             `json:"txnIndex" db:"txn_index"`
	TxnLogIndex     uint             `json:"txnLogIndex" db:"txn_log_index"`
	LogDataType     EventLogDataType `json:"logDataType" db:"log_data_type"`
	TS              time.Time        `json:"ts" db:"ts"`
	LogData         string           `json:"logData" db:"log_data"`
}

type TokenBalance struct {
	ID              uint64             `json:"id,omitempty" db:"id,omitempty"`
	Key             prototyp.Key       `json:"-" db:"key"`
	ContractAddress prototyp.Hash      `json:"contractAddress" db:"contract_address"`
	ContractType    ContractType       `json:"contractType" db:"contract_type"`
	AccountAddress  prototyp.HashMaybe `json:"accountAddress,omitempty" db:"account_address,omitempty"`
	TokenID         prototyp.BigInt    `json:"tokenID,omitempty" db:"token_id"`
	Balance         prototyp.BigInt    `json:"balance" db:"balance"`
	BlockHash       prototyp.Hash      `json:"blockHash,omitempty" db:"block_hash"`
	BlockNumber     uint64             `json:"blockNumber,omitempty" db:"block_num"`
	UpdateID        uint64             `json:"updateId" db:"update_id"`
	ChainID         uint64             `json:"chainId" db:"-"`
}

type TokenHistory struct {
	ID              uint64        `json:"id" db:"id,omitempty"`
	BlockNumber     uint64        `json:"blockNumber" db:"block_num"`
	BlockHash       prototyp.Hash `json:"blockHash" db:"block_hash"`
	ContractAddress prototyp.Hash `json:"contractAddress" db:"contract_address"`
	ContractType    ContractType  `json:"contractType" db:"contract_type"`
	FromAddress     prototyp.Hash `json:"fromAddress" db:"from_address"`
	ToAddress       prototyp.Hash `json:"toAddress" db:"to_address"`
	TxnHash         prototyp.Hash `json:"txnHash" db:"txn_hash"`
	TxnIndex        uint          `json:"txnIndex" db:"txn_index"`
	TxnLogIndex     uint          `json:"txnLogIndex" db:"txn_log_index"`
	LogData         string        `json:"logData" db:"log_data"`
	TS              time.Time     `json:"ts" db:"ts"`
}

type TokenSupply struct {
	TokenID prototyp.BigInt `json:"tokenID,omitempty" db:"token_id"`
	Supply  prototyp.BigInt `json:"supply" db:"supply"`
	ChainID uint64          `json:"chainId" db:"-"`
}

type Transaction struct {
	TxnHash     prototyp.Hash  `json:"txnHash"`
	BlockNumber uint64         `json:"blockNumber"`
	BlockHash   prototyp.Hash  `json:"blockHash"`
	ChainID     uint64         `json:"chainId"`
	MetaTxnID   *string        `json:"metaTxnID"`
	Transfers   []*TxnTransfer `json:"transfers"`
	Timestamp   *time.Time     `json:"timestamp" db:"ts,omitempty"`
}

type TxnTransfer struct {
	TransferType    TxnTransferType   `json:"transferType"`
	ContractAddress prototyp.Hash     `json:"contractAddress"`
	ContractType    ContractType      `json:"contractType"`
	From            prototyp.Hash     `json:"from"`
	To              prototyp.Hash     `json:"to"`
	TokenIds        []prototyp.BigInt `json:"tokenIds,omitempty"`
	Amounts         []prototyp.BigInt `json:"amounts"`
}

type TransactionHistoryFilter struct {
	AccountAddress     *string  `json:"accountAddress"`
	ContractAddress    *string  `json:"contractAddress"`
	AccountAddresses   []string `json:"accountAddresses"`
	ContractAddresses  []string `json:"contractAddresses"`
	TransactionHashes  []string `json:"transactionHashes"`
	MetaTransactionIDs []string `json:"metaTransactionIDs"`
	FromBlock          *uint64  `json:"fromBlock"`
	ToBlock            *uint64  `json:"toBlock"`
}

type Page struct {
	PageSize *uint32      `json:"pageSize"`
	Page     *uint32      `json:"page"`
	More     *bool        `json:"more"`
	Column   *string      `json:"column"`
	Before   *interface{} `json:"before"`
	After    *interface{} `json:"after"`
	Sort     []*SortBy    `json:"sort"`
}

type SortBy struct {
	Column string     `json:"column"`
	Order  *SortOrder `json:"order"`
}

type Indexer interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetChainID(ctx context.Context) (uint64, error)
	GetEtherBalance(ctx context.Context, accountAddress *string) (*EtherBalance, error)
	GetTokenBalances(ctx context.Context, accountAddress *string, contractAddress *string) ([]*TokenBalance, error)
	GetTokenSupplies(ctx context.Context, contractAddress string) (*ContractType, []*TokenSupply, error)
	GetTokenSuppliesMap(ctx context.Context, tokenMap map[string][]string) (map[string][]*TokenSupply, error)
	GetBalanceUpdates(ctx context.Context, contractAddress string, lastUpdateID uint64, page *Page) (*Page, []*TokenBalance, error)
	GetTransactionHistory(ctx context.Context, filter *TransactionHistoryFilter, page *Page) (*Page, []*Transaction, error)
}

var WebRPCServices = map[string][]string{
	"Indexer": {
		"Ping",
		"Version",
		"RuntimeStatus",
		"GetChainID",
		"GetEtherBalance",
		"GetTokenBalances",
		"GetTokenSupplies",
		"GetTokenSuppliesMap",
		"GetBalanceUpdates",
		"GetTransactionHistory",
	},
}

//
// Client
//

const IndexerPathPrefix = "/rpc/Indexer/"

type indexerClient struct {
	client HTTPClient
	urls   [10]string
}

func NewIndexerClient(addr string, client HTTPClient) Indexer {
	prefix := urlBase(addr) + IndexerPathPrefix
	urls := [10]string{
		prefix + "Ping",
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "GetChainID",
		prefix + "GetEtherBalance",
		prefix + "GetTokenBalances",
		prefix + "GetTokenSupplies",
		prefix + "GetTokenSuppliesMap",
		prefix + "GetBalanceUpdates",
		prefix + "GetTransactionHistory",
	}
	return &indexerClient{
		client: client,
		urls:   urls,
	}
}

func (c *indexerClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	return out.Ret0, err
}

func (c *indexerClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *indexerClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, err
}

func (c *indexerClient) GetChainID(ctx context.Context) (uint64, error) {
	out := struct {
		Ret0 uint64 `json:"chainID"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], nil, &out)
	return out.Ret0, err
}

func (c *indexerClient) GetEtherBalance(ctx context.Context, accountAddress *string) (*EtherBalance, error) {
	in := struct {
		Arg0 *string `json:"accountAddress"`
	}{accountAddress}
	out := struct {
		Ret0 *EtherBalance `json:"balance"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], in, &out)
	return out.Ret0, err
}

func (c *indexerClient) GetTokenBalances(ctx context.Context, accountAddress *string, contractAddress *string) ([]*TokenBalance, error) {
	in := struct {
		Arg0 *string `json:"accountAddress"`
		Arg1 *string `json:"contractAddress"`
	}{accountAddress, contractAddress}
	out := struct {
		Ret0 []*TokenBalance `json:"balances"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], in, &out)
	return out.Ret0, err
}

func (c *indexerClient) GetTokenSupplies(ctx context.Context, contractAddress string) (*ContractType, []*TokenSupply, error) {
	in := struct {
		Arg0 string `json:"contractAddress"`
	}{contractAddress}
	out := struct {
		Ret0 *ContractType  `json:"contractType"`
		Ret1 []*TokenSupply `json:"tokenIDs"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[6], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *indexerClient) GetTokenSuppliesMap(ctx context.Context, tokenMap map[string][]string) (map[string][]*TokenSupply, error) {
	in := struct {
		Arg0 map[string][]string `json:"tokenMap"`
	}{tokenMap}
	out := struct {
		Ret0 map[string][]*TokenSupply `json:"supplies"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[7], in, &out)
	return out.Ret0, err
}

func (c *indexerClient) GetBalanceUpdates(ctx context.Context, contractAddress string, lastUpdateID uint64, page *Page) (*Page, []*TokenBalance, error) {
	in := struct {
		Arg0 string `json:"contractAddress"`
		Arg1 uint64 `json:"lastUpdateID"`
		Arg2 *Page  `json:"page"`
	}{contractAddress, lastUpdateID, page}
	out := struct {
		Ret0 *Page           `json:"page"`
		Ret1 []*TokenBalance `json:"balances"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[8], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *indexerClient) GetTransactionHistory(ctx context.Context, filter *TransactionHistoryFilter, page *Page) (*Page, []*Transaction, error) {
	in := struct {
		Arg0 *TransactionHistoryFilter `json:"filter"`
		Arg1 *Page                     `json:"page"`
	}{filter, page}
	out := struct {
		Ret0 *Page          `json:"page"`
		Ret1 []*Transaction `json:"transactions"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[9], in, &out)
	return out.Ret0, out.Ret1, err
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
		return clientError("failed to marshal json request", err)
	}
	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return clientError("request failed", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = clientError("failed to close response body", cerr)
		}
	}()

	if err = ctx.Err(); err != nil {
		return clientError("aborted because context was done", err)
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	if out != nil {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return clientError("failed to read response body", err)
		}

		err = json.Unmarshal(respBody, &out)
		if err != nil {
			return clientError("failed to unmarshal json response body", err)
		}
		if err = ctx.Err(); err != nil {
			return clientError("aborted because context was done", err)
		}
	}

	return nil
}

// errorFromResponse builds a webrpc Error from a non-200 HTTP response.
func errorFromResponse(resp *http.Response) Error {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clientError("failed to read server error response body", err)
	}

	var respErr ErrorPayload
	if err := json.Unmarshal(respBody, &respErr); err != nil {
		return clientError("failed unmarshal error response", err)
	}

	errCode := ErrorCode(respErr.Code)

	if HTTPStatusFromErrorCode(errCode) == 0 {
		return ErrorInternal("invalid code returned from server error response: %s", respErr.Code)
	}

	return &rpcErr{
		code:  errCode,
		msg:   respErr.Msg,
		cause: errors.New(respErr.Cause),
	}
}

func clientError(desc string, err error) Error {
	return WrapError(ErrInternal, err, desc)
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

type ErrorPayload struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Cause  string `json:"cause,omitempty"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

type Error interface {
	// Code is of the valid error codes
	Code() ErrorCode

	// Msg returns a human-readable, unstructured messages describing the error
	Msg() string

	// Cause is reason for the error
	Cause() error

	// Error returns a string of the form "webrpc error <Code>: <Msg>"
	Error() string

	// Error response payload
	Payload() ErrorPayload
}

func Errorf(code ErrorCode, msgf string, args ...interface{}) Error {
	msg := fmt.Sprintf(msgf, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code)}
}

func WrapError(code ErrorCode, cause error, format string, args ...interface{}) Error {
	msg := fmt.Sprintf(format, args...)
	if IsValidErrorCode(code) {
		return &rpcErr{code: code, msg: msg, cause: cause}
	}
	return &rpcErr{code: ErrInternal, msg: "invalid error type " + string(code), cause: cause}
}

func Failf(format string, args ...interface{}) Error {
	return Errorf(ErrFail, format, args...)
}

func WrapFailf(cause error, format string, args ...interface{}) Error {
	return WrapError(ErrFail, cause, format, args...)
}

func ErrorNotFound(format string, args ...interface{}) Error {
	return Errorf(ErrNotFound, format, args...)
}

func ErrorInvalidArgument(argument string, validationMsg string) Error {
	return Errorf(ErrInvalidArgument, argument+" "+validationMsg)
}

func ErrorRequiredArgument(argument string) Error {
	return ErrorInvalidArgument(argument, "is required")
}

func ErrorInternal(format string, args ...interface{}) Error {
	return Errorf(ErrInternal, format, args...)
}

type ErrorCode string

const (
	// Unknown error. For example when handling errors raised by APIs that do not
	// return enough error information.
	ErrUnknown ErrorCode = "unknown"

	// Fail error. General failure error type.
	ErrFail ErrorCode = "fail"

	// Canceled indicates the operation was cancelled (typically by the caller).
	ErrCanceled ErrorCode = "canceled"

	// InvalidArgument indicates client specified an invalid argument. It
	// indicates arguments that are problematic regardless of the state of the
	// system (i.e. a malformed file name, required argument, number out of range,
	// etc.).
	ErrInvalidArgument ErrorCode = "invalid argument"

	// DeadlineExceeded means operation expired before completion. For operations
	// that change the state of the system, this error may be returned even if the
	// operation has completed successfully (timeout).
	ErrDeadlineExceeded ErrorCode = "deadline exceeded"

	// NotFound means some requested entity was not found.
	ErrNotFound ErrorCode = "not found"

	// BadRoute means that the requested URL path wasn't routable to a webrpc
	// service and method. This is returned by the generated server, and usually
	// shouldn't be returned by applications. Instead, applications should use
	// NotFound or Unimplemented.
	ErrBadRoute ErrorCode = "bad route"

	// AlreadyExists means an attempt to create an entity failed because one
	// already exists.
	ErrAlreadyExists ErrorCode = "already exists"

	// PermissionDenied indicates the caller does not have permission to execute
	// the specified operation. It must not be used if the caller cannot be
	// identified (Unauthenticated).
	ErrPermissionDenied ErrorCode = "permission denied"

	// Unauthenticated indicates the request does not have valid authentication
	// credentials for the operation.
	ErrUnauthenticated ErrorCode = "unauthenticated"

	// ResourceExhausted indicates some resource has been exhausted, perhaps a
	// per-user quota, or perhaps the entire file system is out of space.
	ErrResourceExhausted ErrorCode = "resource exhausted"

	// FailedPrecondition indicates operation was rejected because the system is
	// not in a state required for the operation's execution. For example, doing
	// an rmdir operation on a directory that is non-empty, or on a non-directory
	// object, or when having conflicting read-modify-write on the same resource.
	ErrFailedPrecondition ErrorCode = "failed precondition"

	// Aborted indicates the operation was aborted, typically due to a concurrency
	// issue like sequencer check failures, transaction aborts, etc.
	ErrAborted ErrorCode = "aborted"

	// OutOfRange means operation was attempted past the valid range. For example,
	// seeking or reading past end of a paginated collection.
	//
	// Unlike InvalidArgument, this error indicates a problem that may be fixed if
	// the system state changes (i.e. adding more items to the collection).
	//
	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.
	// We recommend using OutOfRange (the more specific error) when it applies so
	// that callers who are iterating through a space can easily look for an
	// OutOfRange error to detect when they are done.
	ErrOutOfRange ErrorCode = "out of range"

	// Unimplemented indicates operation is not implemented or not
	// supported/enabled in this service.
	ErrUnimplemented ErrorCode = "unimplemented"

	// Internal errors. When some invariants expected by the underlying system
	// have been broken. In other words, something bad happened in the library or
	// backend service. Do not confuse with HTTP Internal Server Error; an
	// Internal error could also happen on the client code, i.e. when parsing a
	// server response.
	ErrInternal ErrorCode = "internal"

	// Unavailable indicates the service is currently unavailable. This is a most
	// likely a transient condition and may be corrected by retrying with a
	// backoff.
	ErrUnavailable ErrorCode = "unavailable"

	// DataLoss indicates unrecoverable data loss or corruption.
	ErrDataLoss ErrorCode = "data loss"

	// ErrNone is the zero-value, is considered an empty error and should not be
	// used.
	ErrNone ErrorCode = ""
)

func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case ErrCanceled:
		return 408 // RequestTimeout
	case ErrUnknown:
		return 400 // Bad Request
	case ErrFail:
		return 422 // Unprocessable Entity
	case ErrInvalidArgument:
		return 400 // BadRequest
	case ErrDeadlineExceeded:
		return 408 // RequestTimeout
	case ErrNotFound:
		return 404 // Not Found
	case ErrBadRoute:
		return 404 // Not Found
	case ErrAlreadyExists:
		return 409 // Conflict
	case ErrPermissionDenied:
		return 403 // Forbidden
	case ErrUnauthenticated:
		return 401 // Unauthorized
	case ErrResourceExhausted:
		return 403 // Forbidden
	case ErrFailedPrecondition:
		return 412 // Precondition Failed
	case ErrAborted:
		return 409 // Conflict
	case ErrOutOfRange:
		return 400 // Bad Request
	case ErrUnimplemented:
		return 501 // Not Implemented
	case ErrInternal:
		return 500 // Internal Server Error
	case ErrUnavailable:
		return 503 // Service Unavailable
	case ErrDataLoss:
		return 500 // Internal Server Error
	case ErrNone:
		return 200 // OK
	default:
		return 0 // Invalid!
	}
}

func IsErrorCode(err error, code ErrorCode) bool {
	if rpcErr, ok := err.(Error); ok {
		if rpcErr.Code() == code {
			return true
		}
	}
	return false
}

func IsValidErrorCode(code ErrorCode) bool {
	return HTTPStatusFromErrorCode(code) != 0
}

type rpcErr struct {
	code  ErrorCode
	msg   string
	cause error
}

func (e *rpcErr) Code() ErrorCode {
	return e.code
}

func (e *rpcErr) Msg() string {
	return e.msg
}

func (e *rpcErr) Cause() error {
	return e.cause
}

func (e *rpcErr) Error() string {
	if e.cause != nil && e.cause.Error() != "" {
		if e.msg != "" {
			return fmt.Sprintf("webrpc %s error: %s -- %s", e.code, e.cause.Error(), e.msg)
		} else {
			return fmt.Sprintf("webrpc %s error: %s", e.code, e.cause.Error())
		}
	} else {
		return fmt.Sprintf("webrpc %s error: %s", e.code, e.msg)
	}
}

func (e *rpcErr) Payload() ErrorPayload {
	statusCode := HTTPStatusFromErrorCode(e.Code())
	errPayload := ErrorPayload{
		Status: statusCode,
		Code:   string(e.Code()),
		Msg:    e.Msg(),
		Error:  e.Error(),
	}
	if e.Cause() != nil {
		errPayload.Cause = e.Cause().Error()
	}
	return errPayload
}

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