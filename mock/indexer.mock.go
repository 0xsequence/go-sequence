// Code generated by MockGen. DO NOT EDIT.
// Source: ../indexer (interfaces: Indexer)
//
// Generated by this command:
//
//	mockgen -destination indexer.mock.go -package mock -mock_names Indexer=Indexer ../indexer Indexer
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	indexer "github.com/0xsequence/go-sequence/indexer"
	gomock "go.uber.org/mock/gomock"
)

// Indexer is a mock of Indexer interface.
type Indexer struct {
	ctrl     *gomock.Controller
	recorder *IndexerMockRecorder
	isgomock struct{}
}

// IndexerMockRecorder is the mock recorder for Indexer.
type IndexerMockRecorder struct {
	mock *Indexer
}

// NewIndexer creates a new mock instance.
func NewIndexer(ctrl *gomock.Controller) *Indexer {
	mock := &Indexer{ctrl: ctrl}
	mock.recorder = &IndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Indexer) EXPECT() *IndexerMockRecorder {
	return m.recorder
}

// AddWebhookListener mocks base method.
func (m *Indexer) AddWebhookListener(ctx context.Context, url string, filters *indexer.EventFilter, projectId *uint64) (bool, *indexer.WebhookListener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWebhookListener", ctx, url, filters, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*indexer.WebhookListener)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddWebhookListener indicates an expected call of AddWebhookListener.
func (mr *IndexerMockRecorder) AddWebhookListener(ctx, url, filters, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWebhookListener", reflect.TypeOf((*Indexer)(nil).AddWebhookListener), ctx, url, filters, projectId)
}

// FetchTransactionReceipt mocks base method.
func (m *Indexer) FetchTransactionReceipt(ctx context.Context, txnHash string, maxBlockWait *int) (*indexer.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTransactionReceipt", ctx, txnHash, maxBlockWait)
	ret0, _ := ret[0].(*indexer.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTransactionReceipt indicates an expected call of FetchTransactionReceipt.
func (mr *IndexerMockRecorder) FetchTransactionReceipt(ctx, txnHash, maxBlockWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTransactionReceipt", reflect.TypeOf((*Indexer)(nil).FetchTransactionReceipt), ctx, txnHash, maxBlockWait)
}

// FetchTransactionReceiptWithFilter mocks base method.
func (m *Indexer) FetchTransactionReceiptWithFilter(ctx context.Context, filter *indexer.TransactionFilter, maxBlockWait *int) (*indexer.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTransactionReceiptWithFilter", ctx, filter, maxBlockWait)
	ret0, _ := ret[0].(*indexer.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTransactionReceiptWithFilter indicates an expected call of FetchTransactionReceiptWithFilter.
func (mr *IndexerMockRecorder) FetchTransactionReceiptWithFilter(ctx, filter, maxBlockWait any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTransactionReceiptWithFilter", reflect.TypeOf((*Indexer)(nil).FetchTransactionReceiptWithFilter), ctx, filter, maxBlockWait)
}

// GetAllWebhookListeners mocks base method.
func (m *Indexer) GetAllWebhookListeners(ctx context.Context, projectId *uint64) ([]*indexer.WebhookListener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWebhookListeners", ctx, projectId)
	ret0, _ := ret[0].([]*indexer.WebhookListener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWebhookListeners indicates an expected call of GetAllWebhookListeners.
func (mr *IndexerMockRecorder) GetAllWebhookListeners(ctx, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWebhookListeners", reflect.TypeOf((*Indexer)(nil).GetAllWebhookListeners), ctx, projectId)
}

// GetBalanceUpdates mocks base method.
func (m *Indexer) GetBalanceUpdates(ctx context.Context, contractAddress string, lastBlockNumber uint64, lastBlockHash *string, page *indexer.Page) (*indexer.Page, []*indexer.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceUpdates", ctx, contractAddress, lastBlockNumber, lastBlockHash, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.TokenBalance)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBalanceUpdates indicates an expected call of GetBalanceUpdates.
func (mr *IndexerMockRecorder) GetBalanceUpdates(ctx, contractAddress, lastBlockNumber, lastBlockHash, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceUpdates", reflect.TypeOf((*Indexer)(nil).GetBalanceUpdates), ctx, contractAddress, lastBlockNumber, lastBlockHash, page)
}

// GetChainID mocks base method.
func (m *Indexer) GetChainID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainID indicates an expected call of GetChainID.
func (mr *IndexerMockRecorder) GetChainID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*Indexer)(nil).GetChainID), ctx)
}

// GetEtherBalance mocks base method.
func (m *Indexer) GetEtherBalance(ctx context.Context, accountAddress *string) (*indexer.EtherBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEtherBalance", ctx, accountAddress)
	ret0, _ := ret[0].(*indexer.EtherBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEtherBalance indicates an expected call of GetEtherBalance.
func (mr *IndexerMockRecorder) GetEtherBalance(ctx, accountAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEtherBalance", reflect.TypeOf((*Indexer)(nil).GetEtherBalance), ctx, accountAddress)
}

// GetNativeTokenBalance mocks base method.
func (m *Indexer) GetNativeTokenBalance(ctx context.Context, accountAddress *string) (*indexer.NativeTokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNativeTokenBalance", ctx, accountAddress)
	ret0, _ := ret[0].(*indexer.NativeTokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNativeTokenBalance indicates an expected call of GetNativeTokenBalance.
func (mr *IndexerMockRecorder) GetNativeTokenBalance(ctx, accountAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNativeTokenBalance", reflect.TypeOf((*Indexer)(nil).GetNativeTokenBalance), ctx, accountAddress)
}

// GetOrderbookOrders mocks base method.
func (m *Indexer) GetOrderbookOrders(ctx context.Context, page *indexer.Page, orderbookContractAddress, collectionAddress string, currencyAddresses []string, filter *indexer.OrderbookOrderFilter, orderStatuses []indexer.OrderStatus, filters []*indexer.OrderbookOrderFilter, beforeExpiryTimestamp, blockNumberAfter, createdAtAfter int64) (*indexer.Page, []*indexer.OrderbookOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderbookOrders", ctx, page, orderbookContractAddress, collectionAddress, currencyAddresses, filter, orderStatuses, filters, beforeExpiryTimestamp, blockNumberAfter, createdAtAfter)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.OrderbookOrder)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrderbookOrders indicates an expected call of GetOrderbookOrders.
func (mr *IndexerMockRecorder) GetOrderbookOrders(ctx, page, orderbookContractAddress, collectionAddress, currencyAddresses, filter, orderStatuses, filters, beforeExpiryTimestamp, blockNumberAfter, createdAtAfter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderbookOrders", reflect.TypeOf((*Indexer)(nil).GetOrderbookOrders), ctx, page, orderbookContractAddress, collectionAddress, currencyAddresses, filter, orderStatuses, filters, beforeExpiryTimestamp, blockNumberAfter, createdAtAfter)
}

// GetTokenBalances mocks base method.
func (m *Indexer) GetTokenBalances(ctx context.Context, accountAddress, contractAddress, tokenID *string, includeMetadata *bool, metadataOptions *indexer.MetadataOptions, includeCollectionTokens *bool, page *indexer.Page) (*indexer.Page, []*indexer.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalances", ctx, accountAddress, contractAddress, tokenID, includeMetadata, metadataOptions, includeCollectionTokens, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.TokenBalance)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokenBalances indicates an expected call of GetTokenBalances.
func (mr *IndexerMockRecorder) GetTokenBalances(ctx, accountAddress, contractAddress, tokenID, includeMetadata, metadataOptions, includeCollectionTokens, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalances", reflect.TypeOf((*Indexer)(nil).GetTokenBalances), ctx, accountAddress, contractAddress, tokenID, includeMetadata, metadataOptions, includeCollectionTokens, page)
}

// GetTokenBalancesByContract mocks base method.
func (m *Indexer) GetTokenBalancesByContract(ctx context.Context, filter *indexer.TokenBalancesByContractFilter, omitMetadata *bool, page *indexer.Page) (*indexer.Page, []*indexer.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalancesByContract", ctx, filter, omitMetadata, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.TokenBalance)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokenBalancesByContract indicates an expected call of GetTokenBalancesByContract.
func (mr *IndexerMockRecorder) GetTokenBalancesByContract(ctx, filter, omitMetadata, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalancesByContract", reflect.TypeOf((*Indexer)(nil).GetTokenBalancesByContract), ctx, filter, omitMetadata, page)
}

// GetTokenBalancesDetails mocks base method.
func (m *Indexer) GetTokenBalancesDetails(ctx context.Context, filter *indexer.TokenBalancesFilter, omitMetadata *bool, page *indexer.Page) (*indexer.Page, []*indexer.NativeTokenBalance, []*indexer.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalancesDetails", ctx, filter, omitMetadata, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.NativeTokenBalance)
	ret2, _ := ret[2].([]*indexer.TokenBalance)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTokenBalancesDetails indicates an expected call of GetTokenBalancesDetails.
func (mr *IndexerMockRecorder) GetTokenBalancesDetails(ctx, filter, omitMetadata, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalancesDetails", reflect.TypeOf((*Indexer)(nil).GetTokenBalancesDetails), ctx, filter, omitMetadata, page)
}

// GetTokenBalancesSummary mocks base method.
func (m *Indexer) GetTokenBalancesSummary(ctx context.Context, filter *indexer.TokenBalancesFilter, omitMetadata *bool, page *indexer.Page) (*indexer.Page, []*indexer.NativeTokenBalance, []*indexer.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenBalancesSummary", ctx, filter, omitMetadata, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.NativeTokenBalance)
	ret2, _ := ret[2].([]*indexer.TokenBalance)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTokenBalancesSummary indicates an expected call of GetTokenBalancesSummary.
func (mr *IndexerMockRecorder) GetTokenBalancesSummary(ctx, filter, omitMetadata, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalancesSummary", reflect.TypeOf((*Indexer)(nil).GetTokenBalancesSummary), ctx, filter, omitMetadata, page)
}

// GetTokenIDRanges mocks base method.
func (m *Indexer) GetTokenIDRanges(ctx context.Context, contractAddress string) (indexer.ContractType, []*indexer.TokenIDRange, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenIDRanges", ctx, contractAddress)
	ret0, _ := ret[0].(indexer.ContractType)
	ret1, _ := ret[1].([]*indexer.TokenIDRange)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTokenIDRanges indicates an expected call of GetTokenIDRanges.
func (mr *IndexerMockRecorder) GetTokenIDRanges(ctx, contractAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenIDRanges", reflect.TypeOf((*Indexer)(nil).GetTokenIDRanges), ctx, contractAddress)
}

// GetTokenIDs mocks base method.
func (m *Indexer) GetTokenIDs(ctx context.Context, contractAddress string, page *indexer.Page) (*indexer.Page, indexer.ContractType, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenIDs", ctx, contractAddress, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].(indexer.ContractType)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTokenIDs indicates an expected call of GetTokenIDs.
func (mr *IndexerMockRecorder) GetTokenIDs(ctx, contractAddress, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenIDs", reflect.TypeOf((*Indexer)(nil).GetTokenIDs), ctx, contractAddress, page)
}

// GetTokenSupplies mocks base method.
func (m *Indexer) GetTokenSupplies(ctx context.Context, contractAddress string, includeMetadata *bool, metadataOptions *indexer.MetadataOptions, page *indexer.Page) (*indexer.Page, indexer.ContractType, []*indexer.TokenSupply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenSupplies", ctx, contractAddress, includeMetadata, metadataOptions, page)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].(indexer.ContractType)
	ret2, _ := ret[2].([]*indexer.TokenSupply)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetTokenSupplies indicates an expected call of GetTokenSupplies.
func (mr *IndexerMockRecorder) GetTokenSupplies(ctx, contractAddress, includeMetadata, metadataOptions, page any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenSupplies", reflect.TypeOf((*Indexer)(nil).GetTokenSupplies), ctx, contractAddress, includeMetadata, metadataOptions, page)
}

// GetTokenSuppliesMap mocks base method.
func (m *Indexer) GetTokenSuppliesMap(ctx context.Context, tokenMap map[string][]string, includeMetadata *bool, metadataOptions *indexer.MetadataOptions) (map[string][]*indexer.TokenSupply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenSuppliesMap", ctx, tokenMap, includeMetadata, metadataOptions)
	ret0, _ := ret[0].(map[string][]*indexer.TokenSupply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenSuppliesMap indicates an expected call of GetTokenSuppliesMap.
func (mr *IndexerMockRecorder) GetTokenSuppliesMap(ctx, tokenMap, includeMetadata, metadataOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenSuppliesMap", reflect.TypeOf((*Indexer)(nil).GetTokenSuppliesMap), ctx, tokenMap, includeMetadata, metadataOptions)
}

// GetTopOrders mocks base method.
func (m *Indexer) GetTopOrders(ctx context.Context, orderbookContractAddress, collectionAddress string, currencyAddresses, tokenIDs []string, isListing bool, priceSort indexer.SortOrder, excludeUser *string) ([]*indexer.OrderbookOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopOrders", ctx, orderbookContractAddress, collectionAddress, currencyAddresses, tokenIDs, isListing, priceSort, excludeUser)
	ret0, _ := ret[0].([]*indexer.OrderbookOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTopOrders indicates an expected call of GetTopOrders.
func (mr *IndexerMockRecorder) GetTopOrders(ctx, orderbookContractAddress, collectionAddress, currencyAddresses, tokenIDs, isListing, priceSort, excludeUser any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopOrders", reflect.TypeOf((*Indexer)(nil).GetTopOrders), ctx, orderbookContractAddress, collectionAddress, currencyAddresses, tokenIDs, isListing, priceSort, excludeUser)
}

// GetTransactionHistory mocks base method.
func (m *Indexer) GetTransactionHistory(ctx context.Context, filter *indexer.TransactionHistoryFilter, page *indexer.Page, includeMetadata *bool, metadataOptions *indexer.MetadataOptions) (*indexer.Page, []*indexer.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionHistory", ctx, filter, page, includeMetadata, metadataOptions)
	ret0, _ := ret[0].(*indexer.Page)
	ret1, _ := ret[1].([]*indexer.Transaction)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionHistory indicates an expected call of GetTransactionHistory.
func (mr *IndexerMockRecorder) GetTransactionHistory(ctx, filter, page, includeMetadata, metadataOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionHistory", reflect.TypeOf((*Indexer)(nil).GetTransactionHistory), ctx, filter, page, includeMetadata, metadataOptions)
}

// GetWebhookListener mocks base method.
func (m *Indexer) GetWebhookListener(ctx context.Context, id uint64, projectId *uint64) (*indexer.WebhookListener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWebhookListener", ctx, id, projectId)
	ret0, _ := ret[0].(*indexer.WebhookListener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWebhookListener indicates an expected call of GetWebhookListener.
func (mr *IndexerMockRecorder) GetWebhookListener(ctx, id, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebhookListener", reflect.TypeOf((*Indexer)(nil).GetWebhookListener), ctx, id, projectId)
}

// PauseAllWebhookListeners mocks base method.
func (m *Indexer) PauseAllWebhookListeners(ctx context.Context, projectId *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PauseAllWebhookListeners", ctx, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PauseAllWebhookListeners indicates an expected call of PauseAllWebhookListeners.
func (mr *IndexerMockRecorder) PauseAllWebhookListeners(ctx, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PauseAllWebhookListeners", reflect.TypeOf((*Indexer)(nil).PauseAllWebhookListeners), ctx, projectId)
}

// Ping mocks base method.
func (m *Indexer) Ping(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *IndexerMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*Indexer)(nil).Ping), ctx)
}

// RemoveAllWebhookListeners mocks base method.
func (m *Indexer) RemoveAllWebhookListeners(ctx context.Context, projectId *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllWebhookListeners", ctx, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAllWebhookListeners indicates an expected call of RemoveAllWebhookListeners.
func (mr *IndexerMockRecorder) RemoveAllWebhookListeners(ctx, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllWebhookListeners", reflect.TypeOf((*Indexer)(nil).RemoveAllWebhookListeners), ctx, projectId)
}

// RemoveWebhookListener mocks base method.
func (m *Indexer) RemoveWebhookListener(ctx context.Context, id uint64, projectId *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWebhookListener", ctx, id, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveWebhookListener indicates an expected call of RemoveWebhookListener.
func (mr *IndexerMockRecorder) RemoveWebhookListener(ctx, id, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWebhookListener", reflect.TypeOf((*Indexer)(nil).RemoveWebhookListener), ctx, id, projectId)
}

// ResumeAllWebhookListeners mocks base method.
func (m *Indexer) ResumeAllWebhookListeners(ctx context.Context, projectId *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeAllWebhookListeners", ctx, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResumeAllWebhookListeners indicates an expected call of ResumeAllWebhookListeners.
func (mr *IndexerMockRecorder) ResumeAllWebhookListeners(ctx, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeAllWebhookListeners", reflect.TypeOf((*Indexer)(nil).ResumeAllWebhookListeners), ctx, projectId)
}

// RuntimeStatus mocks base method.
func (m *Indexer) RuntimeStatus(ctx context.Context) (*indexer.RuntimeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RuntimeStatus", ctx)
	ret0, _ := ret[0].(*indexer.RuntimeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RuntimeStatus indicates an expected call of RuntimeStatus.
func (mr *IndexerMockRecorder) RuntimeStatus(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RuntimeStatus", reflect.TypeOf((*Indexer)(nil).RuntimeStatus), ctx)
}

// SubscribeBalanceUpdates mocks base method.
func (m *Indexer) SubscribeBalanceUpdates(ctx context.Context, contractAddress string, stream indexer.SubscribeBalanceUpdatesStreamWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeBalanceUpdates", ctx, contractAddress, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeBalanceUpdates indicates an expected call of SubscribeBalanceUpdates.
func (mr *IndexerMockRecorder) SubscribeBalanceUpdates(ctx, contractAddress, stream any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeBalanceUpdates", reflect.TypeOf((*Indexer)(nil).SubscribeBalanceUpdates), ctx, contractAddress, stream)
}

// SubscribeEvents mocks base method.
func (m *Indexer) SubscribeEvents(ctx context.Context, filter *indexer.EventFilter, stream indexer.SubscribeEventsStreamWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeEvents", ctx, filter, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeEvents indicates an expected call of SubscribeEvents.
func (mr *IndexerMockRecorder) SubscribeEvents(ctx, filter, stream any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeEvents", reflect.TypeOf((*Indexer)(nil).SubscribeEvents), ctx, filter, stream)
}

// SubscribeReceipts mocks base method.
func (m *Indexer) SubscribeReceipts(ctx context.Context, filter *indexer.TransactionFilter, stream indexer.SubscribeReceiptsStreamWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeReceipts", ctx, filter, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeReceipts indicates an expected call of SubscribeReceipts.
func (mr *IndexerMockRecorder) SubscribeReceipts(ctx, filter, stream any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeReceipts", reflect.TypeOf((*Indexer)(nil).SubscribeReceipts), ctx, filter, stream)
}

// SyncBalance mocks base method.
func (m *Indexer) SyncBalance(ctx context.Context, accountAddress, contractAddress string, tokenID *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncBalance", ctx, accountAddress, contractAddress, tokenID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncBalance indicates an expected call of SyncBalance.
func (mr *IndexerMockRecorder) SyncBalance(ctx, accountAddress, contractAddress, tokenID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncBalance", reflect.TypeOf((*Indexer)(nil).SyncBalance), ctx, accountAddress, contractAddress, tokenID)
}

// ToggleWebhookListener mocks base method.
func (m *Indexer) ToggleWebhookListener(ctx context.Context, id uint64, projectId *uint64) (*indexer.WebhookListener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleWebhookListener", ctx, id, projectId)
	ret0, _ := ret[0].(*indexer.WebhookListener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToggleWebhookListener indicates an expected call of ToggleWebhookListener.
func (mr *IndexerMockRecorder) ToggleWebhookListener(ctx, id, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleWebhookListener", reflect.TypeOf((*Indexer)(nil).ToggleWebhookListener), ctx, id, projectId)
}

// UpdateWebhookListener mocks base method.
func (m *Indexer) UpdateWebhookListener(ctx context.Context, listener *indexer.WebhookListener, projectId *uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWebhookListener", ctx, listener, projectId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWebhookListener indicates an expected call of UpdateWebhookListener.
func (mr *IndexerMockRecorder) UpdateWebhookListener(ctx, listener, projectId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWebhookListener", reflect.TypeOf((*Indexer)(nil).UpdateWebhookListener), ctx, listener, projectId)
}

// Version mocks base method.
func (m *Indexer) Version(ctx context.Context) (*indexer.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", ctx)
	ret0, _ := ret[0].(*indexer.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *IndexerMockRecorder) Version(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*Indexer)(nil).Version), ctx)
}
