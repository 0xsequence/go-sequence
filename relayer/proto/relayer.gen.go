// sequence-relayer v0.4.0 e7d1a1c505a5bcc9189296c17a160c72e1f6ef8a
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/golang
// Do not edit by hand. Update your webrpc schema and re-generate.
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
	return "e7d1a1c505a5bcc9189296c17a160c72e1f6ef8a"
}

//
// Types
//

type ETHTxnStatus uint

const (
	ETHTxnStatus_UNKNOWN          ETHTxnStatus = 0
	ETHTxnStatus_DROPPED          ETHTxnStatus = 1
	ETHTxnStatus_QUEUED           ETHTxnStatus = 2
	ETHTxnStatus_SENT             ETHTxnStatus = 3
	ETHTxnStatus_SUCCEEDED        ETHTxnStatus = 4
	ETHTxnStatus_PARTIALLY_FAILED ETHTxnStatus = 5
	ETHTxnStatus_FAILED           ETHTxnStatus = 6
)

var ETHTxnStatus_name = map[uint]string{
	0: "UNKNOWN",
	1: "DROPPED",
	2: "QUEUED",
	3: "SENT",
	4: "SUCCEEDED",
	5: "PARTIALLY_FAILED",
	6: "FAILED",
}

var ETHTxnStatus_value = map[string]uint{
	"UNKNOWN":          0,
	"DROPPED":          1,
	"QUEUED":           2,
	"SENT":             3,
	"SUCCEEDED":        4,
	"PARTIALLY_FAILED": 5,
	"FAILED":           6,
}

func (x ETHTxnStatus) String() string {
	return ETHTxnStatus_name[uint(x)]
}

func (x ETHTxnStatus) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(ETHTxnStatus_name[uint(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *ETHTxnStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = ETHTxnStatus(ETHTxnStatus_value[j])
	return nil
}

type TransferType uint32

const (
	TransferType_SEND            TransferType = 0
	TransferType_RECEIVE         TransferType = 1
	TransferType_BRIDGE_DEPOSIT  TransferType = 2
	TransferType_BRIDGE_WITHDRAW TransferType = 3
	TransferType_BURN            TransferType = 4
	TransferType_UNKNOWN         TransferType = 5
)

var TransferType_name = map[uint32]string{
	0: "SEND",
	1: "RECEIVE",
	2: "BRIDGE_DEPOSIT",
	3: "BRIDGE_WITHDRAW",
	4: "BURN",
	5: "UNKNOWN",
}

var TransferType_value = map[string]uint32{
	"SEND":            0,
	"RECEIVE":         1,
	"BRIDGE_DEPOSIT":  2,
	"BRIDGE_WITHDRAW": 3,
	"BURN":            4,
	"UNKNOWN":         5,
}

func (x TransferType) String() string {
	return TransferType_name[uint32(x)]
}

func (x TransferType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(TransferType_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *TransferType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = TransferType(TransferType_value[j])
	return nil
}

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

func (x FeeTokenType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(FeeTokenType_name[uint32(x)])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (x *FeeTokenType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*x = FeeTokenType(FeeTokenType_value[j])
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
	HealthOK   bool            `json:"healthOK"`
	StartTime  time.Time       `json:"startTime"`
	Uptime     uint64          `json:"uptime"`
	Ver        string          `json:"ver"`
	Branch     string          `json:"branch"`
	CommitHash string          `json:"commitHash"`
	Senders    []*SenderStatus `json:"senders"`
	Checks     *RuntimeChecks  `json:"checks"`
}

type SenderStatus struct {
	Index   uint32 `json:"index"`
	Address string `json:"address"`
	Active  bool   `json:"active"`
}

type RuntimeChecks struct {
}

type SequenceContext struct {
	Factory              string `json:"factory"`
	MainModule           string `json:"mainModule"`
	MainModuleUpgradable string `json:"mainModuleUpgradable"`
	GuestModule          string `json:"guestModule"`
	Utils                string `json:"utils"`
}

type WalletConfig struct {
	Address   string          `json:"address"`
	Signers   []*WalletSigner `json:"signers"`
	Threshold uint16          `json:"threshold"`
	ChainId   *uint64         `json:"chainId"`
}

type WalletSigner struct {
	Address string `json:"address"`
	Weight  uint8  `json:"weight"`
}

type GasSponsor struct {
	ID        uint64        `json:"id" db:"id,omitempty"`
	Address   prototyp.Hash `json:"address" db:"address"`
	Name      string        `json:"name" db:"name"`
	Active    bool          `json:"active" db:"active"`
	UpdatedAt *time.Time    `json:"updatedAt" db:"updated_at,omitempty"`
	CreatedAt *time.Time    `json:"createdAt" db:"created_at,omitempty"`
}

type GasSponsorUsage struct {
	Name         string     `json:"name" db:"name"`
	ID           uint64     `json:"id" db:"gas_sponsor_id,omitempty"`
	TotalGasUsed int64      `json:"totalGasUsed" db:"total_gas_used"`
	TotalTxnFees float64    `json:"totalTxnFees" db:"total_txn_fees"`
	AvgGasPrice  float64    `json:"avgGasPrice" db:"avg_gas_price"`
	TotalTxns    int64      `json:"totalTxns" db:"total_txns"`
	StartTime    *time.Time `json:"startTime"`
	EndTime      *time.Time `json:"endTime"`
}

type MetaTxn struct {
	WalletAddress string `json:"walletAddress" db:"wallet_address"`
	Contract      string `json:"contract" db:"to_address"`
	Input         string `json:"input" db:"tx_data"`
}

type MetaTxnLog struct {
	ID              uint64                 `json:"id" db:"id,omitempty"`
	TxnHash         prototyp.HashMaybe     `json:"txnHash" db:"txn_hash"`
	TxnNonce        prototyp.BigInt        `json:"txnNonce" db:"txn_nonce"`
	MetaTxnID       *string                `json:"metaTxnID" db:"meta_txn_id"`
	TxnStatus       ETHTxnStatus           `json:"txnStatus" db:"txn_status"`
	TxnRevertReason string                 `json:"txnRevertReason" db:"txn_revert_reason"`
	Requeues        uint                   `json:"requeues" db:"requeues"`
	QueuedAt        *time.Time             `json:"queuedAt" db:"queued_at,omitempty"`
	SentAt          *time.Time             `json:"sentAt" db:"sent_at,omitempty"`
	MinedAt         *time.Time             `json:"minedAt" db:"mined_at,omitempty"`
	Target          prototyp.Hash          `json:"target" db:"target"`
	Input           prototyp.Hash          `json:"input" db:"input"`
	TxnArgs         map[string]interface{} `json:"txnArgs" db:"txn_args"`
	TxnReceipt      map[string]interface{} `json:"txnReceipt" db:"txn_receipt,omitempty"`
	WalletAddress   prototyp.Hash          `json:"walletAddress" db:"wallet_address"`
	MetaTxnNonce    prototyp.BigInt        `json:"metaTxnNonce" db:"metatx_nonce"`
	GasLimit        uint64                 `json:"gasLimit" db:"gas_limit"`
	GasPrice        prototyp.BigInt        `json:"gasPrice" db:"gas_price"`
	GasUsed         uint64                 `json:"gasUsed" db:"gas_used"`
	IsWhitelisted   bool                   `json:"isWhitelisted" db:"is_whitelisted,omitempty"`
	GasSponsor      *uint64                `json:"gasSponsor" db:"gas_sponsor_id,omitempty"`
	UpdatedAt       *time.Time             `json:"updatedAt" db:"updated_at,omitempty"`
	CreatedAt       *time.Time             `json:"createdAt" db:"created_at,omitempty"`
}

type MetaTxnEntry struct {
	ID              uint64        `json:"id" db:"id,omitempty"`
	MetaTxnID       string        `json:"metaTxnID" db:"metatx_logs_id"`
	TxnStatus       ETHTxnStatus  `json:"txnStatus" db:"status"`
	TxnRevertReason string        `json:"txnRevertReason" db:"revert_reason"`
	Index           uint64        `json:"index" db:"index"`
	Logs            []interface{} `json:"logs" db:"logs,omitempty"`
	UpdatedAt       *time.Time    `json:"updatedAt" db:"updated_at,omitempty"`
	CreatedAt       *time.Time    `json:"createdAt" db:"created_at,omitempty"`
}

type MetaTxnReceipt struct {
	ID           string               `json:"id"`
	Status       string               `json:"status"`
	RevertReason *string              `json:"revertReason"`
	Index        uint                 `json:"index"`
	Logs         []*MetaTxnReceiptLog `json:"logs"`
	Receipts     []*MetaTxnReceipt    `json:"receipts"`
	TxnReceipt   string               `json:"txnReceipt"`
}

type MetaTxnReceiptLog struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type Transaction struct {
	TxnHash     *string                `json:"txnHash"`
	BlockNumber uint64                 `json:"blockNumber"`
	ChainId     uint64                 `json:"chainId"`
	MetaTxnID   *string                `json:"metaTxnID"`
	Transfers   []*TxnLogTransfer      `json:"transfers"`
	Users       map[string]*TxnLogUser `json:"users"`
	Timestamp   *time.Time             `json:"timestamp" db:"ts,omitempty"`
}

type TxnLogUser struct {
	Username string `json:"username"`
}

type TxnLogTransfer struct {
	TransferType    *TransferType     `json:"transferType"`
	ContractAddress string            `json:"contractAddress"`
	From            string            `json:"from"`
	To              string            `json:"to"`
	Ids             []prototyp.BigInt `json:"ids"`
	Amounts         []prototyp.BigInt `json:"amounts"`
}

type SentTransactionsFilter struct {
	Pending *bool `json:"pending"`
	Failed  *bool `json:"failed"`
}

type SimulateResult struct {
	Executed  bool    `json:"executed"`
	Succeeded bool    `json:"succeeded"`
	Result    *string `json:"result"`
	Reason    *string `json:"reason"`
	GasUsed   uint    `json:"gasUsed"`
	GasLimit  uint    `json:"gasLimit"`
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
	OriginAddress   *string       `json:"originAddress"`
	TokenID         *string       `json:"tokenID"`
}

type Page struct {
	PageSize     *uint32      `json:"pageSize"`
	Page         *uint32      `json:"page"`
	TotalRecords *uint64      `json:"totalRecords"`
	Column       *string      `json:"column"`
	Before       *interface{} `json:"before"`
	After        *interface{} `json:"after"`
	Sort         []*SortBy    `json:"sort"`
}

type SortBy struct {
	Column string     `json:"column"`
	Order  *SortOrder `json:"order"`
}

type Relayer interface {
	Ping(ctx context.Context) (bool, error)
	Version(ctx context.Context) (*Version, error)
	RuntimeStatus(ctx context.Context) (*RuntimeStatus, error)
	GetSequenceContext(ctx context.Context) (*SequenceContext, error)
	GetChainID(ctx context.Context) (uint64, error)
	SendMetaTxn(ctx context.Context, call *MetaTxn, quote *string) (bool, string, error)
	GetMetaTxnNonce(ctx context.Context, walletContractAddress string, space *string) (string, error)
	GetMetaTxnReceipt(ctx context.Context, metaTxID string) (*MetaTxnReceipt, error)
	Simulate(ctx context.Context, wallet string, transactions string) ([]*SimulateResult, error)
	FeeTokens(ctx context.Context) (bool, []*FeeToken, error)
	FeeOptions(ctx context.Context, wallet string, to string, data string) ([]*FeeOption, *string, error)
	SentTransactions(ctx context.Context, filter *SentTransactionsFilter, page *Page) (*Page, []*Transaction, error)
	PendingTransactions(ctx context.Context, page *Page) (*Page, []*Transaction, error)
	ListGasSponsors(ctx context.Context, page *Page) (*Page, []*GasSponsor, error)
	AddGasSponsor(ctx context.Context, address string, name *string, active *bool) (bool, *GasSponsor, error)
	UpdateGasSponsor(ctx context.Context, address string, name *string, active *bool) (bool, *GasSponsor, error)
	ReportGasSponsorUsage(ctx context.Context, startTime *time.Time, endTime *time.Time) ([]*GasSponsorUsage, error)
}

var WebRPCServices = map[string][]string{
	"Relayer": {
		"Ping",
		"Version",
		"RuntimeStatus",
		"GetSequenceContext",
		"GetChainID",
		"SendMetaTxn",
		"GetMetaTxnNonce",
		"GetMetaTxnReceipt",
		"Simulate",
		"FeeTokens",
		"FeeOptions",
		"SentTransactions",
		"PendingTransactions",
		"ListGasSponsors",
		"AddGasSponsor",
		"UpdateGasSponsor",
		"ReportGasSponsorUsage",
	},
}

//
// Client
//

const RelayerPathPrefix = "/rpc/Relayer/"

type relayerClient struct {
	client HTTPClient
	urls   [17]string
}

func NewRelayerClient(addr string, client HTTPClient) Relayer {
	prefix := urlBase(addr) + RelayerPathPrefix
	urls := [17]string{
		prefix + "Ping",
		prefix + "Version",
		prefix + "RuntimeStatus",
		prefix + "GetSequenceContext",
		prefix + "GetChainID",
		prefix + "SendMetaTxn",
		prefix + "GetMetaTxnNonce",
		prefix + "GetMetaTxnReceipt",
		prefix + "Simulate",
		prefix + "FeeTokens",
		prefix + "FeeOptions",
		prefix + "SentTransactions",
		prefix + "PendingTransactions",
		prefix + "ListGasSponsors",
		prefix + "AddGasSponsor",
		prefix + "UpdateGasSponsor",
		prefix + "ReportGasSponsorUsage",
	}
	return &relayerClient{
		client: client,
		urls:   urls,
	}
}

func (c *relayerClient) Ping(ctx context.Context) (bool, error) {
	out := struct {
		Ret0 bool `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	return out.Ret0, err
}

func (c *relayerClient) Version(ctx context.Context) (*Version, error) {
	out := struct {
		Ret0 *Version `json:"version"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[1], nil, &out)
	return out.Ret0, err
}

func (c *relayerClient) RuntimeStatus(ctx context.Context) (*RuntimeStatus, error) {
	out := struct {
		Ret0 *RuntimeStatus `json:"status"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[2], nil, &out)
	return out.Ret0, err
}

func (c *relayerClient) GetSequenceContext(ctx context.Context) (*SequenceContext, error) {
	out := struct {
		Ret0 *SequenceContext `json:"data"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[3], nil, &out)
	return out.Ret0, err
}

func (c *relayerClient) GetChainID(ctx context.Context) (uint64, error) {
	out := struct {
		Ret0 uint64 `json:"chainID"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[4], nil, &out)
	return out.Ret0, err
}

func (c *relayerClient) SendMetaTxn(ctx context.Context, call *MetaTxn, quote *string) (bool, string, error) {
	in := struct {
		Arg0 *MetaTxn `json:"call"`
		Arg1 *string  `json:"quote"`
	}{call, quote}
	out := struct {
		Ret0 bool   `json:"status"`
		Ret1 string `json:"txnHash"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[5], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) GetMetaTxnNonce(ctx context.Context, walletContractAddress string, space *string) (string, error) {
	in := struct {
		Arg0 string  `json:"walletContractAddress"`
		Arg1 *string `json:"space"`
	}{walletContractAddress, space}
	out := struct {
		Ret0 string `json:"nonce"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[6], in, &out)
	return out.Ret0, err
}

func (c *relayerClient) GetMetaTxnReceipt(ctx context.Context, metaTxID string) (*MetaTxnReceipt, error) {
	in := struct {
		Arg0 string `json:"metaTxID"`
	}{metaTxID}
	out := struct {
		Ret0 *MetaTxnReceipt `json:"receipt"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[7], in, &out)
	return out.Ret0, err
}

func (c *relayerClient) Simulate(ctx context.Context, wallet string, transactions string) ([]*SimulateResult, error) {
	in := struct {
		Arg0 string `json:"wallet"`
		Arg1 string `json:"transactions"`
	}{wallet, transactions}
	out := struct {
		Ret0 []*SimulateResult `json:"results"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[8], in, &out)
	return out.Ret0, err
}

func (c *relayerClient) FeeTokens(ctx context.Context) (bool, []*FeeToken, error) {
	out := struct {
		Ret0 bool        `json:"isFeeRequired"`
		Ret1 []*FeeToken `json:"tokens"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[9], nil, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) FeeOptions(ctx context.Context, wallet string, to string, data string) ([]*FeeOption, *string, error) {
	in := struct {
		Arg0 string `json:"wallet"`
		Arg1 string `json:"to"`
		Arg2 string `json:"data"`
	}{wallet, to, data}
	out := struct {
		Ret0 []*FeeOption `json:"options"`
		Ret1 *string      `json:"quote"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[10], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) SentTransactions(ctx context.Context, filter *SentTransactionsFilter, page *Page) (*Page, []*Transaction, error) {
	in := struct {
		Arg0 *SentTransactionsFilter `json:"filter"`
		Arg1 *Page                   `json:"page"`
	}{filter, page}
	out := struct {
		Ret0 *Page          `json:"page"`
		Ret1 []*Transaction `json:"transactions"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[11], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) PendingTransactions(ctx context.Context, page *Page) (*Page, []*Transaction, error) {
	in := struct {
		Arg0 *Page `json:"page"`
	}{page}
	out := struct {
		Ret0 *Page          `json:"page"`
		Ret1 []*Transaction `json:"transactions"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[12], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) ListGasSponsors(ctx context.Context, page *Page) (*Page, []*GasSponsor, error) {
	in := struct {
		Arg0 *Page `json:"page"`
	}{page}
	out := struct {
		Ret0 *Page         `json:"page"`
		Ret1 []*GasSponsor `json:"gasSponsors"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[13], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) AddGasSponsor(ctx context.Context, address string, name *string, active *bool) (bool, *GasSponsor, error) {
	in := struct {
		Arg0 string  `json:"address"`
		Arg1 *string `json:"name"`
		Arg2 *bool   `json:"active"`
	}{address, name, active}
	out := struct {
		Ret0 bool        `json:"status"`
		Ret1 *GasSponsor `json:"gasSponsor"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[14], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) UpdateGasSponsor(ctx context.Context, address string, name *string, active *bool) (bool, *GasSponsor, error) {
	in := struct {
		Arg0 string  `json:"address"`
		Arg1 *string `json:"name"`
		Arg2 *bool   `json:"active"`
	}{address, name, active}
	out := struct {
		Ret0 bool        `json:"status"`
		Ret1 *GasSponsor `json:"gasSponsor"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[15], in, &out)
	return out.Ret0, out.Ret1, err
}

func (c *relayerClient) ReportGasSponsorUsage(ctx context.Context, startTime *time.Time, endTime *time.Time) ([]*GasSponsorUsage, error) {
	in := struct {
		Arg0 *time.Time `json:"startTime"`
		Arg1 *time.Time `json:"endTime"`
	}{startTime, endTime}
	out := struct {
		Ret0 []*GasSponsorUsage `json:"gasSponsorUsage"`
	}{}

	err := doJSONRequest(ctx, c.client, c.urls[16], in, &out)
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
