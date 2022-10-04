package sequence

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/walletgasestimator"
	"github.com/goware/cachestore"
	"github.com/goware/cachestore/memlru"
)

const (
	defaultEstimatorCacheSize = 500

	areEOAsMaxConcurrentTasks = 10
)

var (
	// byte values that represent booleans in the cache.
	cachedTrue  = byte('t')
	cachedFalse = byte('f')
)

type CallOverride struct {
	Code      string
	Balance   *big.Int
	Nonce     *big.Int
	StateDiff []*StateOverride
	State     []*StateOverride
}

type StateOverride struct {
	Key   string
	Value string
}

type EstimateTransaction struct {
	From common.Address
	To   common.Address
	Data []byte
}

type Estimator struct {
	BaseCost     uint64
	DataOneCost  uint64
	DataZeroCost uint64

	cache cachestore.Store[[]byte]
}

type SimulateResult walletgasestimator.MainModuleGasEstimationSimulateResult

var defaultEstimator = &Estimator{
	BaseCost:     21000,
	DataOneCost:  16,
	DataZeroCost: 4,
}

var gasEstimatorCode = hexutil.Encode(contracts.GasEstimator.DeployedBin)
var walletGasEstimatorCode = hexutil.Encode(contracts.WalletGasEstimator.DeployedBin)

func NewEstimator() *Estimator {
	defaultCache, _ := memlru.NewWithSize[[]byte](defaultEstimatorCacheSize)
	return &Estimator{
		BaseCost:     defaultEstimator.BaseCost,
		DataZeroCost: defaultEstimator.DataZeroCost,
		DataOneCost:  defaultEstimator.DataOneCost,
		cache:        defaultCache,
	}
}

func (e *Estimator) SetCache(cache cachestore.Store[[]byte]) *Estimator {
	e.cache = cache
	return e
}

func (e *Estimator) CalldataCost(data []byte) uint64 {
	cost := e.BaseCost

	for _, b := range data {
		if b == 0 {
			cost += e.DataZeroCost
		} else {
			cost += e.DataOneCost
		}
	}

	return cost
}

// BuildProxy for address based on https://eips.ethereum.org/EIPS/eip-1167
// the bytecode contains an aditional SLOAD to mimic the Sequence proxies
// bytecode:
//
// |           0x00000000      36             calldatasize          cds
// |           0x00000001      3d             returndatasize        0 cds
// |           0x00000002      3d             returndatasize        0 0 cds
// |           0x00000003      37             calldatacopy
// |           0x00000004      30             address               addr
// |           0x00000005      54             sload                 stub
// |           0x00000006      50             pop
// |           0x00000007      3d             returndatasize        0
// |           0x00000008      3d             returndatasize        0 0
// |           0x00000009      3d             returndatasize        0 0 0
// |           0x0000000a      36             calldatasize          cds 0 0 0
// |           0x0000000b      3d             returndatasize        0 cds 0 0 0
// |           0x0000000c      73bebebebebe.  push20 0xbebebebe     0xbebe 0 cds 0 0 0
// |           0x00000020      5a             gas                   gas 0xbebe 0 cds 0 0 0
// |           0x00000021      f4             delegatecall          suc 0
// |           0x00000022      3d             returndatasize        rds suc 0
// |           0x00000023      82             dup3                  0 rds suc 0
// |           0x00000024      80             dup1                  0 0 rds suc 0
// |           0x00000025      3e             returndatacopy        suc 0
// |           0x00000026      90             swap1                 0 suc
// |           0x00000027      3d             returndatasize        rds 0 suc
// |           0x00000028      91             swap2                 suc 0 rds
// |           0x00000029      602d           push1 0x2e            0x2e suc 0 rds
// |       ,=< 0x0000002b      57             jumpi                 0 rds
// |       |   0x0000002c      fd             revert
// |       `-> 0x0000002d      5b             jumpdest              0 rds
// \           0x0000002e      f3             return
func BuildProxy(addr common.Address) string {
	return "0x363d3d3730543d3d3d363d73" + strings.Replace(addr.String(), "0x", "", 1) + "5af43d82803e903d91602d57fd5bf3"
}

func (e *Estimator) EstimateCall(ctx context.Context, provider *ethrpc.Provider, call *EstimateTransaction, overrides map[common.Address]*CallOverride, blockTag string) (*big.Int, error) {
	if blockTag == "" {
		blockTag = "latest"
	}

	from := call.From
	if from.Hash().Big().Cmp(common.Big0) == 0 {
		from = stubAddress()
	}

	finalOverrides := map[common.Address]*CallOverride{
		from: {Code: gasEstimatorCode},
	}

	if overrides != nil {
		for key, value := range overrides {
			if key == from {
				return nil, fmt.Errorf("can't override address from")
			}

			finalOverrides[key] = value
		}
	}

	estimator := ethcontract.NewContractCaller(from, contracts.GasEstimator.ABI, provider)
	callData, err := estimator.Encode("estimate", call.To, call.Data)
	if err != nil {
		return nil, err
	}

	type Call struct {
		To   common.Address `json:"to"`
		Data string         `json:"data"`
	}

	estimateCall := &Call{
		To:   from,
		Data: "0x" + common.Bytes2Hex(callData),
	}

	var res string
	err = provider.RPC.Call(&res, "eth_call", estimateCall, blockTag, finalOverrides)
	if err != nil {
		return nil, err
	}

	resBytes := common.Hex2Bytes(strings.Replace(res, "0x", "", 1))

	var success bool
	var result []byte
	var gas *big.Int

	if err := ethcoder.AbiDecoder([]string{"bool", "bytes", "uint256"}, resBytes, []interface{}{&success, &result, &gas}); err != nil {
		return nil, err
	}

	gas.Add(gas, big.NewInt(int64(e.CalldataCost(call.Data))))

	if !success {
		if len(result) <= 68 {
			return gas, fmt.Errorf("error calling estimate: UNKNOWN_REASON")
		}

		reason := string(result[68 : len(result)-1])
		return gas, fmt.Errorf("error calling estimate: " + reason)
	}

	return gas, nil
}

func (e *Estimator) AreEOAs(ctx context.Context, provider *ethrpc.Provider, walletConfig WalletConfig) ([]bool, error) {
	res := make([]bool, walletConfig.Signers.Len())

	// Get non-eoa signers
	// required for computing worse case scenario
	var wg sync.WaitGroup

	errCh := make(chan error)
	workersCh := make(chan struct{}, areEOAsMaxConcurrentTasks)

	for i := range walletConfig.Signers {
		wg.Add(1)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case err := <-errCh:
			return nil, err
		case workersCh <- struct{}{}: // wait until a worker slot becomes available to continue
		}

		go func(ctx context.Context, i int, address common.Address) {
			defer func() {
				wg.Done()
				<-workersCh // release the worker
			}()

			var err error
			res[i], err = e.isEOA(ctx, provider, address)
			if err != nil {
				errCh <- err
				return
			}
		}(ctx, i, walletConfig.Signers[i].Address)
	}
	wg.Wait()

	return res, nil
}

func (e *Estimator) isEOA(ctx context.Context, provider *ethrpc.Provider, address common.Address) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	key := fmt.Sprintf("isEOA::%d::%v", provider.Config.ChaindID, address)

	if val, exists, _ := e.cache.Get(ctx, key); exists {
		// we have recorded data for this key, let's use it
		return val[0] == cachedTrue, nil
	}

	code, err := provider.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}

	if len(code) == 0 {
		// if the address does not contain a smart contract, then it's an EOA
		_ = e.cache.Set(ctx, key, []byte{cachedTrue})
		return true, nil
	}

	_ = e.cache.Set(ctx, key, []byte{cachedFalse})
	return false, nil
}

func (e *Estimator) PickSigners(ctx context.Context, walletConfig WalletConfig, isEoa []bool) ([]bool, error) {
	type SortedSigner struct {
		s WalletConfigSigner
		i int
	}

	// Create a copy of the signers array
	// this will be sorted and used to pick the worst case scenario for the signers
	sortedSigners := make([]SortedSigner, walletConfig.Signers.Len())
	for i, _ := range walletConfig.Signers {
		sortedSigners[i] = SortedSigner{
			s: walletConfig.Signers[i],
			i: i,
		}
	}

	sort.SliceStable(sortedSigners, func(a, b int) bool {
		if !isEoa[sortedSigners[a].i] && isEoa[sortedSigners[b].i] {
			return true
		}

		return sortedSigners[a].s.Weight < sortedSigners[b].s.Weight
	})

	weightSum := 0

	// Define what signers are goint go be signing
	// it should construct a worse case scenario for the signature
	willSign := make([]bool, walletConfig.Signers.Len())
	threshold := int(walletConfig.Threshold)

	// Pick signers until we reach the threshold we stop
	// We use the sorted signers to get the ones with the non EOA and with lowest weight first
	for _, s := range sortedSigners {
		weightSum += int(s.s.Weight)
		willSign[s.i] = true

		if weightSum >= threshold {
			break
		}
	}

	return willSign, nil
}

func stubAddress() common.Address {
	raw := make([]byte, 20)
	rand.Read(raw)
	return common.BytesToAddress(raw)
}

func (e *Estimator) BuildStubSignature(walletConfig WalletConfig, willSign []bool, isEoa []bool) []byte {
	parts := make(SignatureParts, walletConfig.Signers.Len())

	// pre-determined signature, tailored for worse-case scenario in gas costs
	// TODO: Compute average siganture and present a more likely scenario for a more close estimation
	sig := common.Hex2Bytes("1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a01b02")

	for i, s := range walletConfig.Signers {
		if willSign[i] {
			if isEoa[i] {
				parts[i] = &SignaturePart{
					Weight:  s.Weight,
					Address: s.Address,
					Type:    SignaturePartTypeEOA,
					Value:   sig,
				}
			} else {
				sig := e.BuildStubSignature(WalletConfig{
					Threshold: 2,
					Signers: WalletConfigSigners{
						WalletConfigSigner{
							Address: stubAddress(),
							Weight:  1,
						},
						WalletConfigSigner{
							Address: stubAddress(),
							Weight:  1,
						},
						WalletConfigSigner{
							Address: stubAddress(),
							Weight:  1,
						},
					},
				}, []bool{false, true, true}, []bool{true, true, true})

				sig = append(sig, 03)

				parts[i] = &SignaturePart{
					Weight:  s.Weight,
					Address: s.Address,
					Type:    SignaturePartTypeDynamic,
					Value:   sig,
				}
			}
		} else {
			parts[i] = &SignaturePart{
				Weight:  s.Weight,
				Address: s.Address,
				Type:    SignaturePartTypeAddress,
			}
		}
	}

	signature := Signature{
		Threshold: walletConfig.Threshold,
		Signers:   parts,
	}

	encoded, err := signature.Encode()

	// This method builds a signature from scratch
	// it should never faild to encode it
	if err != nil {
		panic(err)
	}

	return encoded
}

func (e *Estimator) Estimate(ctx context.Context, provider *ethrpc.Provider, address common.Address, walletConfig WalletConfig, walletContext WalletContext, txs Transactions) (uint64, error) {
	isEOA, err := e.AreEOAs(ctx, provider, walletConfig)
	if err != nil {
		return 0, err
	}

	willSign, err := e.PickSigners(ctx, walletConfig, isEOA)
	if err != nil {
		return 0, err
	}

	signature := e.BuildStubSignature(walletConfig, willSign, isEOA)

	overrides := map[common.Address]*CallOverride{
		walletContext.MainModuleAddress:           {Code: walletGasEstimatorCode},
		walletContext.MainModuleUpgradableAddress: {Code: walletGasEstimatorCode},
	}

	isDeployed, err := IsWalletDeployed(provider, address)
	if err != nil {
		return 0, err
	}

	if !isDeployed {
		overrides[address] = &CallOverride{
			Code: BuildProxy(walletContext.MainModuleAddress),
		}
	}

	estimates := make([]*big.Int, len(txs)+1)

	// The nonce is ignored by the MainModuleGasEstimator
	// so we use a stub nonce takes at least 4 bytes
	nonce := big.NewInt(4294967295)

	// Compute gas estimation for slices of all transactions
	// including no transaction execution and all transactions

	for i := range estimates {
		subTxs := txs[0:i]

		encTxs, err := subTxs.EncodedTransactions()
		if err != nil {
			return 0, err
		}

		execData, err := contracts.WalletMainModule.Encode("execute", encTxs, nonce, signature)
		if err != nil {
			return 0, err
		}

		estimated, err := e.EstimateCall(ctx, provider, &EstimateTransaction{
			To:   address,
			Data: execData,
		}, overrides, "")
		if err != nil {
			return 0, err
		}

		estimates[i] = estimated
	}

	// Apply gas limits to all transactions
	for i := range txs {
		txs[i].GasLimit = big.NewInt(0).Sub(estimates[i+1], estimates[i])
	}

	return estimates[len(estimates)-1].Uint64(), nil
}

func Simulate(provider *ethrpc.Provider, wallet common.Address, transactions Transactions, block string, overrides map[common.Address]*CallOverride) ([]SimulateResult, error) {
	if block == "" {
		block = "latest"
	}

	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	callData, err := contracts.WalletGasEstimator.Encode("simulateExecute", encoded)
	if err != nil {
		return nil, err
	}

	type ethCallParams struct {
		To   common.Address `json:"to"`
		Data string         `json:"data"`
	}

	params := ethCallParams{
		To:   wallet,
		Data: hexutil.Encode(callData),
	}

	allOverrides := map[common.Address]*CallOverride{
		wallet: {Code: walletGasEstimatorCode},
	}
	for address, override := range overrides {
		if address == wallet {
			return nil, fmt.Errorf("cannot override wallet address %v", wallet.Hex())
		}

		allOverrides[address] = override
	}

	var response string
	err = provider.RPC.Call(&response, "eth_call", params, block, allOverrides)
	if err != nil {
		return nil, err
	}

	resultsData, err := hexutil.Decode(response)
	if err != nil {
		return nil, err
	}

	var results []SimulateResult
	err = contracts.WalletGasEstimator.Decode(&results, "simulateExecute", resultsData)
	if err != nil {
		return nil, err
	}

	return results, nil
}
