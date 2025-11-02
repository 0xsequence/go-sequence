package metadata

import (
	"context"
	"strings"
	"time"

	cachestore "github.com/goware/cachestore2"
)

// TokenMetadataClient defines a minimal interface for interacting with the Sequence Metadata service
// just to query contract and token metadata information.
type TokenMetadataClient interface {
	Ping(ctx context.Context) (bool, error)

	GetContractInfo(ctx context.Context, chainID string, contractAddress string) (*ContractInfo, *uint64, error)

	GetContractInfoBatch(ctx context.Context, chainID string, contractAddresses []string) (map[string]*ContractInfo, *uint64, error)

	GetTokenMetadata(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) ([]*TokenMetadata, *uint64, error)

	// GetTokenMetadataBatch allows you to query the token metadata of a batch of contracts and respective tokenIDs
	// where map is contractAddress::[]tokenID => contractAddress::[]TokenMetadata
	//
	// Note, we limit each request to 50 contracts max and 50 tokens max per contract.
	GetTokenMetadataBatch(ctx context.Context, chainID string, contractTokenMap map[string][]string) (map[string][]*TokenMetadata, *uint64, error)
}

// Ensure TokenMetadataClient interface is a subset of the full MetadataClient interface
var _ TokenMetadataClient = (MetadataClient)(nil)

// NewTokenMetadataClient creates a new TokenMetadataClient instance
// with optional caching support. Please pass nil for cache argument if you don't
// want to use caching.
func NewTokenMetadataClient(clientOptions Options, cache cachestore.Backend, optCacheExpiry ...time.Duration) TokenMetadataClient {
	// If cache is disabled, return a standard client
	if cache == nil {
		return NewClient("", clientOptions)
	}

	cacheExpiry := 4 * time.Minute // 4 minute default
	if len(optCacheExpiry) > 0 {
		cacheExpiry = optCacheExpiry[0]
	}

	contractInfoCache := cachestore.OpenStore[ContractInfo](cache, cachestore.WithDefaultKeyExpiry(cacheExpiry))
	tokenMetadataCache := cachestore.OpenStore[TokenMetadata](cache, cachestore.WithDefaultKeyExpiry(cacheExpiry))

	return &cachedClient{
		client:             NewClient("", clientOptions),
		contractInfoCache:  contractInfoCache,
		tokenMetadataCache: tokenMetadataCache,
	}
}

type cachedClient struct {
	client             TokenMetadataClient
	contractInfoCache  cachestore.Store[ContractInfo]
	tokenMetadataCache cachestore.Store[TokenMetadata]
}

func (c *cachedClient) Ping(ctx context.Context) (bool, error) {
	_, err := c.client.Ping(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *cachedClient) GetContractInfo(ctx context.Context, chainID string, contractAddress string) (*ContractInfo, *uint64, error) {
	key := contractInfoCacheKey(chainID, contractAddress)
	if info, ok, err := c.contractInfoCache.Get(ctx, key); err == nil && ok {
		// served from cache, no task id
		return &info, nil, nil
	} else if err != nil {
		// proceed but return error if cache failure (non-fatal?)
		// we choose to ignore cache error and fallback to network
	}

	ci, taskID, err := c.client.GetContractInfo(ctx, chainID, contractAddress)
	if err != nil || ci == nil {
		return ci, taskID, err
	}
	_ = c.contractInfoCache.Set(ctx, key, *ci)
	return ci, taskID, nil
}

func (c *cachedClient) GetContractInfoBatch(ctx context.Context, chainID string, contractAddresses []string) (map[string]*ContractInfo, *uint64, error) {
	// Prepare cache keys
	keys := make([]string, len(contractAddresses))
	for i, addr := range contractAddresses {
		keys[i] = contractInfoCacheKey(chainID, addr)
	}

	// Attempt batch get from cache
	cachedValues, exists, err := c.contractInfoCache.BatchGet(ctx, keys)
	if err != nil {
		// ignore cache error and treat all as missing
		exists = make([]bool, len(keys))
	}

	// Determine missing addresses
	missing := make([]string, 0)
	for i, ok := range exists {
		if !ok {
			missing = append(missing, contractAddresses[i])
		}
	}

	result := make(map[string]*ContractInfo, len(contractAddresses))
	// Add cached ones
	for i, ok := range exists {
		if ok {
			v := cachedValues[i]
			vv := v
			result[contractAddresses[i]] = &vv
		}
	}

	var taskID *uint64
	if len(missing) > 0 {
		fetched, tID, err := c.client.GetContractInfoBatch(ctx, chainID, missing)
		if err != nil {
			return result, tID, err
		}
		taskID = tID
		// cache fetched
		// gather keys + values
		setKeys := make([]string, 0, len(fetched))
		setVals := make([]ContractInfo, 0, len(fetched))
		for addr, info := range fetched {
			if info == nil {
				continue
			}
			result[addr] = info
			setKeys = append(setKeys, contractInfoCacheKey(chainID, addr))
			setVals = append(setVals, *info)
		}
		if len(setKeys) > 0 {
			_ = c.contractInfoCache.BatchSet(ctx, setKeys, setVals)
		}
	}

	return result, taskID, nil
}

func (c *cachedClient) GetTokenMetadata(ctx context.Context, chainID string, contractAddress string, tokenIDs []string) ([]*TokenMetadata, *uint64, error) {
	// Build cache keys
	keys := make([]string, len(tokenIDs))
	for i, id := range tokenIDs {
		keys[i] = tokenMetadataCacheKey(chainID, contractAddress, id)
	}
	cachedValues, exists, err := c.tokenMetadataCache.BatchGet(ctx, keys)
	if err != nil {
		// ignore cache err
		exists = make([]bool, len(keys))
	}

	missing := make([]string, 0)
	for i, ok := range exists {
		if !ok {
			missing = append(missing, tokenIDs[i])
		}
	}

	// Prepare ordered result slice length of tokenIDs
	result := make([]*TokenMetadata, len(tokenIDs))
	for i, ok := range exists {
		if ok {
			v := cachedValues[i]
			vCopy := v
			result[i] = &vCopy
		}
	}

	var taskID *uint64
	if len(missing) > 0 {
		fetched, tID, err := c.client.GetTokenMetadata(ctx, chainID, contractAddress, missing)
		if err != nil {
			return result, tID, err
		}
		taskID = tID
		// Map missing tokenID -> fetched metadata in order fetched slice corresponds to missing slice order
		// We don't assume order; we iterate fetched and insert into result at matching index of original tokenIDs
		setKeys := make([]string, 0, len(fetched))
		setVals := make([]TokenMetadata, 0, len(fetched))
		// Build lookup for quick index mapping
		idxByID := make(map[string]int, len(tokenIDs))
		for i, id := range tokenIDs {
			idxByID[id] = i
		}
		for _, tm := range fetched {
			if tm == nil {
				continue
			}
			id := tm.TokenID
			if pos, ok := idxByID[id]; ok && result[pos] == nil {
				result[pos] = tm
			}
			setKeys = append(setKeys, tokenMetadataCacheKey(chainID, contractAddress, id))
			setVals = append(setVals, *tm)
		}
		if len(setKeys) > 0 {
			_ = c.tokenMetadataCache.BatchSet(ctx, setKeys, setVals)
		}
	}

	return result, taskID, nil
}

func (c *cachedClient) GetTokenMetadataBatch(ctx context.Context, chainID string, contractTokenMap map[string][]string) (map[string][]*TokenMetadata, *uint64, error) {
	// Prepare aggregate of keys for batch get per contract
	// We'll process each contract separately since BatchGet expects homogeneous keys list.
	result := make(map[string][]*TokenMetadata, len(contractTokenMap))
	var finalTaskID *uint64

	missingPerContract := make(map[string][]string)

	for contract, tokenIDs := range contractTokenMap {
		keys := make([]string, len(tokenIDs))
		for i, id := range tokenIDs {
			keys[i] = tokenMetadataCacheKey(chainID, contract, id)
		}
		cachedVals, exists, err := c.tokenMetadataCache.BatchGet(ctx, keys)
		if err != nil {
			exists = make([]bool, len(keys))
		}
		ordered := make([]*TokenMetadata, len(tokenIDs))
		for i, ok := range exists {
			if ok {
				v := cachedVals[i]
				vCopy := v
				ordered[i] = &vCopy
			}
		}
		// determine missing
		for i, ok := range exists {
			if !ok {
				missingPerContract[contract] = append(missingPerContract[contract], tokenIDs[i])
			}
		}
		result[contract] = ordered
	}

	// Build request map for missing items only
	requestMap := make(map[string][]string, len(missingPerContract))
	for contract, list := range missingPerContract {
		if len(list) > 0 {
			requestMap[contract] = list
		}
	}

	if len(requestMap) > 0 {
		fetchedMap, taskID, err := c.client.GetTokenMetadataBatch(ctx, chainID, requestMap)
		if err != nil {
			return result, taskID, err
		}
		finalTaskID = taskID
		// cache & merge
		setKeys := []string{}
		setVals := []TokenMetadata{}
		for contract, fetchedList := range fetchedMap {
			// index mapping for this contract
			idxByID := make(map[string]int, len(contractTokenMap[contract]))
			for i, id := range contractTokenMap[contract] {
				idxByID[id] = i
			}
			ordered := result[contract]
			for _, tm := range fetchedList {
				if tm == nil {
					continue
				}
				id := tm.TokenID
				if pos, ok := idxByID[id]; ok && ordered[pos] == nil {
					ordered[pos] = tm
				}
				setKeys = append(setKeys, tokenMetadataCacheKey(chainID, contract, id))
				setVals = append(setVals, *tm)
			}
			result[contract] = ordered
		}
		if len(setKeys) > 0 {
			_ = c.tokenMetadataCache.BatchSet(ctx, setKeys, setVals)
		}
	}

	return result, finalTaskID, nil
}

func contractInfoCacheKey(chainID, contractAddress string) string {
	return "contractInfo:" + chainID + ":" + strings.ToLower(contractAddress)
}

func tokenMetadataCacheKey(chainID, contractAddress, tokenID string) string {
	return "tokenMetadata:" + chainID + ":" + strings.ToLower(contractAddress) + ":" + tokenID
}
