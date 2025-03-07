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
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/go-sequence/contracts"
	walletgasestimator "github.com/0xsequence/go-sequence/contracts/gen/v1/walletgasestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/v3/walletestimator"
	"github.com/0xsequence/go-sequence/contracts/gen/v3/walletsimulator"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
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
	Code      string           `json:"code"`
	Balance   *big.Int         `json:"balance"`
	Nonce     *big.Int         `json:"nonce"`
	StateDiff []*StateOverride `json:"stateDiff"`
	State     []*StateOverride `json:"state"`
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
type SimulatorResult walletsimulator.SimulatorResult

var defaultEstimator = &Estimator{
	BaseCost:     21000,
	DataOneCost:  16,
	DataZeroCost: 4,
}

var gasEstimatorCode = hexutil.Encode(contracts.GasEstimator.DeployedBin)
var walletGasEstimatorCode = hexutil.Encode(contracts.V1.WalletGasEstimator.DeployedBin)
var walletGasEstimatorCodeV2 = hexutil.Encode(contracts.V2.WalletGasEstimator.DeployedBin)
var walletGasEstimatorCodeV3 = hexutil.Encode(contracts.V3.WalletEstimator.DeployedBin)
var walletGasSimulatorCodeV3 = hexutil.Encode(contracts.V3.WalletSimulator.DeployedBin)

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
	if from == (common.Address{}) {
		from = stubAddress()
	}

	finalOverrides := map[common.Address]*CallOverride{
		from: {Code: gasEstimatorCode},
	}

	for key, value := range overrides {
		if key == from {
			return nil, fmt.Errorf("can't override address from")
		}

		finalOverrides[key] = value
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
	rpcCall := ethrpc.NewCallBuilder[string]("eth_call", nil, estimateCall, blockTag, finalOverrides)
	_, err = provider.Do(context.Background(), rpcCall.Into(&res))
	if err != nil {
		return nil, err
	}

	resBytes := common.Hex2Bytes(strings.Replace(res, "0x", "", 1))

	var success bool
	var result []byte
	var gas *big.Int

	// Uses `GasEstimator.sol` contract to estimate the gas usage
	// Params `estimate(address _to, bytes calldata _data)`
	// Returns `(bool success, bytes memory result, uint256 gas)`
	if err := ethcoder.ABIUnpackArgumentsByRef([]string{"bool", "bytes", "uint256"}, resBytes, []interface{}{&success, &result, &gas}); err != nil {
		return nil, err
	}

	gas.Add(gas, big.NewInt(int64(e.CalldataCost(call.Data))))

	if !success {
		reason, err := abi.UnpackRevert(result)
		if err == nil {
			return gas, fmt.Errorf("gas usage simulation failed: %v", reason)
		} else {
			return gas, fmt.Errorf("gas usage simulation failed: %v", hexutil.Encode(result))
		}
	}

	return gas, nil
}

func (e *Estimator) AreEOAs(ctx context.Context, provider *ethrpc.Provider, walletConfig core.WalletConfig) (map[common.Address]bool, error) {
	res := make(map[common.Address]bool, len(walletConfig.Signers()))

	// Get non-eoa signers
	// required for computing worse case scenario
	var wg sync.WaitGroup

	errCh := make(chan error)
	workersCh := make(chan struct{}, areEOAsMaxConcurrentTasks)

	var mutex sync.Mutex
	for address := range walletConfig.Signers() {
		wg.Add(1)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case err := <-errCh:
			return nil, err
		case workersCh <- struct{}{}: // wait until a worker slot becomes available to continue
		}

		go func(ctx context.Context, address common.Address) {
			defer func() {
				wg.Done()
				<-workersCh // release the worker
			}()

			var err error
			isEOA, err := e.isEOA(ctx, provider, address)
			if err != nil {
				errCh <- err
				return
			}

			mutex.Lock()
			res[address] = isEOA
			mutex.Unlock()
		}(ctx, address)
	}
	wg.Wait()

	return res, nil
}

func (e *Estimator) isEOA(ctx context.Context, provider *ethrpc.Provider, address common.Address) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	chainID, err := provider.ChainID(ctx)
	if err != nil {
		return false, err
	}

	key := fmt.Sprintf("isEOA::%d::%v", chainID, address)

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

func (e *Estimator) PickSigners(ctx context.Context, walletConfig core.WalletConfig, isEoa map[common.Address]bool) (map[common.Address]bool, error) {
	// type SortedSigner struct {
	// 	s *v1.WalletConfigSigner
	// 	i int
	// }

	// Create a copy of the signers array
	// this will be sorted and used to pick the worst case scenario for the signers
	signersMap := walletConfig.Signers()
	sortedSigners := make([]*v1.WalletConfigSigner, 0, len(signersMap))
	for signer, weight := range signersMap {
		sortedSigners = append(sortedSigners, &v1.WalletConfigSigner{
			Weight:  uint8(weight),
			Address: signer,
		})
	}

	sort.SliceStable(sortedSigners, func(a, b int) bool {
		if !isEoa[sortedSigners[a].Address] && isEoa[sortedSigners[b].Address] {
			return true
		} else if isEoa[sortedSigners[a].Address] && !isEoa[sortedSigners[b].Address] {
			return false
		}

		return sortedSigners[a].Weight < sortedSigners[b].Weight
	})

	weightSum := 0

	// Define what signers are goint go be signing
	// it should construct a worse case scenario for the signature
	willSign := make(map[common.Address]bool, len(walletConfig.Signers()))
	threshold := int(walletConfig.Threshold())

	// Pick signers until we reach the threshold we stop
	// We use the sorted signers to get the ones with the non EOA and with lowest weight first
	for _, s := range sortedSigners {
		if weightSum >= threshold {
			willSign[s.Address] = false
		} else {
			weightSum += int(s.Weight)
			willSign[s.Address] = true
		}
	}

	return willSign, nil
}

func stubAddress() common.Address {
	raw := make([]byte, 20)
	rand.Read(raw)
	return common.BytesToAddress(raw)
}

func (e *Estimator) BuildStubSignature(walletConfig core.WalletConfig, willSign, isEoa map[common.Address]bool) []byte {

	// pre-determined signature, tailored for worse-case scenario in gas costs
	// TODO: Compute average siganture and present a more likely scenario for a more close estimation
	sig := common.Hex2Bytes("1fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a01b02")

	stubSigner := func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		if willSign[signer] {
			if isEoa[signer] {
				return core.SignerSignatureTypeEthSign, sig, nil
			} else {
				address1 := stubAddress()
				address2 := stubAddress()
				address3 := stubAddress()

				var sig []byte
				if _, ok := walletConfig.(*v1.WalletConfig); ok {
					sig = e.BuildStubSignature(&v1.WalletConfig{
						Threshold_: 2,
						Signers_: v1.WalletConfigSigners{
							{
								Address: address1,
								Weight:  1,
							},
							{
								Address: address2,
								Weight:  1,
							},
							{
								Address: address3,
								Weight:  1,
							},
						},
					},
						map[common.Address]bool{address1: false, address2: true, address3: true},
						map[common.Address]bool{address1: true, address2: true, address3: true})
				} else if _, ok := walletConfig.(*v2.WalletConfig); ok {
					sig = e.BuildStubSignature(&v2.WalletConfig{
						Threshold_: 2,
						Tree: v2.WalletConfigTreeNodes(
							&v2.WalletConfigTreeAddressLeaf{
								Address: address1,
								Weight:  1,
							},
							&v2.WalletConfigTreeAddressLeaf{
								Address: address2,
								Weight:  1,
							},
							&v2.WalletConfigTreeAddressLeaf{
								Address: address3,
								Weight:  1,
							},
						),
					},
						map[common.Address]bool{address1: false, address2: true, address3: true},
						map[common.Address]bool{address1: true, address2: true, address3: true})
				} else if _, ok := walletConfig.(*v3.WalletConfig); ok {
					sig = e.BuildStubSignature(&v3.WalletConfig{
						Threshold_: 2,
						Tree: v3.WalletConfigTreeNodes(
							&v3.WalletConfigTreeAddressLeaf{
								Address: address1,
								Weight:  1,
							},
							&v3.WalletConfigTreeAddressLeaf{
								Address: address2,
								Weight:  1,
							},
							&v3.WalletConfigTreeAddressLeaf{
								Address: address3,
								Weight:  1,
							},
						),
					},
						map[common.Address]bool{address1: false, address2: true, address3: true},
						map[common.Address]bool{address1: true, address2: true, address3: true})
				}

				return core.SignerSignatureTypeEIP1271, sig, nil
			}
		} else {
			return 0, nil, core.ErrSigningNoSigner
		}
	}

	if confV1, ok := walletConfig.(*v1.WalletConfig); ok {
		sigV1, err := confV1.BuildSignature(context.Background(), stubSigner, false)
		if err != nil {
			return nil
		}

		encoded, err := sigV1.Data()
		if err != nil {
			return nil
		}
		return encoded
	} else if confV2, ok := walletConfig.(*v2.WalletConfig); ok {
		sigV2, err := confV2.BuildRegularSignature(context.Background(), stubSigner, false)
		if err != nil {
			return nil
		}

		encoded, err := sigV2.Data()
		if err != nil {
			return nil
		}
		return encoded
	} else if confV3, ok := walletConfig.(*v3.WalletConfig); ok {
		sigV3, err := confV3.BuildRegularSignature(context.Background(), stubSigner, false)
		if err != nil {
			return nil
		}

		encoded, err := sigV3.Data()
		if err != nil {
			return nil
		}
		return encoded
	} else {
		return nil
	}
}

// ConvertTransactionsToV3Payload converts sequence.Transactions to v3.DecodedPayload
func ConvertTransactionsToV3Payload(txs Transactions, space, nonce *big.Int) (v3.DecodedPayload, error) {
	calls := make([]v3.Call, len(txs))
	for i, tx := range txs {
		// Convert from Transaction.RevertOnError to v3.BehaviorOnError
		var behaviorOnError v3.BehaviorOnError
		if tx.RevertOnError {
			behaviorOnError = v3.Revert
		} else {
			behaviorOnError = v3.Ignore
		}

		// Handle nested transactions by recursively encoding them
		data := tx.Data
		if tx.IsBundle() {
			subPayload, err := ConvertTransactionsToV3Payload(tx.Transactions, big.NewInt(0), tx.Nonce)
			if err != nil {
				return v3.DecodedPayload{}, fmt.Errorf("unable to convert nested transactions to v3 payload: %w", err)
			}

			encodedSubPayload, err := v3.Encode(subPayload, &tx.To)
			if err != nil {
				return v3.DecodedPayload{}, fmt.Errorf("unable to encode nested v3 payload: %w", err)
			}

			data, err = contracts.V3.WalletStage1Module.Encode("selfExecute", encodedSubPayload)
			if err != nil {
				return v3.DecodedPayload{}, fmt.Errorf("unable to encode nested v3 payload selfExecute call: %w", err)
			}
		}

		calls[i] = v3.Call{
			To:              tx.To,
			Value:           tx.Value,
			Data:            data,
			GasLimit:        tx.GasLimit,
			DelegateCall:    tx.DelegateCall,
			OnlyFallback:    false,
			BehaviorOnError: behaviorOnError,
		}
	}

	return v3.DecodedPayload{
		Kind:      v3.KindTransactions,
		NoChainId: false,
		Calls:     calls,
		Space:     space,
		Nonce:     nonce,
	}, nil
}

func (e *Estimator) Estimate(ctx context.Context, provider *ethrpc.Provider, address common.Address, walletConfig core.WalletConfig, walletContext WalletContext, txs Transactions) (uint64, error) {
	isEOA, err := e.AreEOAs(ctx, provider, walletConfig)
	if err != nil {
		return 0, err
	}

	willSign, err := e.PickSigners(ctx, walletConfig, isEOA)
	if err != nil {
		return 0, err
	}

	signature := e.BuildStubSignature(walletConfig, willSign, isEOA)

	var overrides map[common.Address]*CallOverride
	if _, ok := walletConfig.(*v1.WalletConfig); ok {
		overrides = map[common.Address]*CallOverride{
			walletContext.MainModuleAddress:           {Code: walletGasEstimatorCode},
			walletContext.MainModuleUpgradableAddress: {Code: walletGasEstimatorCode},
		}
	} else if _, ok := walletConfig.(*v2.WalletConfig); ok {
		overrides = map[common.Address]*CallOverride{
			walletContext.MainModuleAddress:           {Code: walletGasEstimatorCodeV2},
			walletContext.MainModuleUpgradableAddress: {Code: walletGasEstimatorCodeV2},
		}
	} else if _, ok := walletConfig.(*v3.WalletConfig); ok {
		overrides = map[common.Address]*CallOverride{
			walletContext.MainModuleAddress:           {Code: walletGasEstimatorCodeV3},
			walletContext.MainModuleUpgradableAddress: {Code: walletGasEstimatorCodeV3},
		}
	} else {
		return 0, fmt.Errorf("unknown wallet config type")
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

		var execData []byte
		if _, ok := walletConfig.(*v1.WalletConfig); ok {
			execData, err = contracts.V1.WalletMainModule.Encode("execute", encTxs, nonce, signature)
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
		} else if _, ok := walletConfig.(*v2.WalletConfig); ok {
			execData, err = contracts.V2.WalletMainModule.Encode("execute", encTxs, nonce, signature)
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
		} else if _, ok := walletConfig.(*v3.WalletConfig); ok {
			payload, err := ConvertTransactionsToV3Payload(subTxs, big.NewInt(0), nonce)
			if err != nil {
				return 0, err
			}

			encoded, err := v3.Encode(payload, &address)
			if err != nil {
				return 0, err
			}

			execData, err := contracts.V3.WalletEstimator.Encode("estimate", encoded, signature)
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
	}

	// Apply gas limits to all transactions
	for i := range txs {
		txs[i].GasLimit = big.NewInt(0).Sub(estimates[i+1], estimates[i])
	}

	return estimates[len(estimates)-1].Uint64(), nil
}

func V1Simulate(provider *ethrpc.Provider, wallet common.Address, transactions Transactions, block string, overrides map[common.Address]*CallOverride) ([]SimulateResult, error) {
	if block == "" {
		block = "latest"
	}

	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	callData, err := contracts.V1.WalletGasEstimator.Encode("simulateExecute", encoded)
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
	rpcCall := ethrpc.NewCallBuilder[string]("eth_call", nil, params, block, allOverrides)
	_, err = provider.Do(context.Background(), rpcCall.Into(&response))
	if err != nil {
		return nil, err
	}

	resultsData, err := hexutil.Decode(response)
	if err != nil {
		return nil, err
	}

	var results []SimulateResult
	err = contracts.V1.WalletGasEstimator.Decode(&results, "simulateExecute", resultsData)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func V2Simulate(provider *ethrpc.Provider, wallet common.Address, transactions Transactions, block string, overrides map[common.Address]*CallOverride) ([]SimulateResult, error) {
	if block == "" {
		block = "latest"
	}

	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	callData, err := contracts.V2.WalletGasEstimator.Encode("simulateExecute", encoded)
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
		wallet: {Code: walletGasEstimatorCodeV2},
	}
	for address, override := range overrides {
		if address == wallet {
			return nil, fmt.Errorf("cannot override wallet address %v", wallet.Hex())
		}

		allOverrides[address] = override
	}

	var response string
	rpcCall := ethrpc.NewCallBuilder[string]("eth_call", nil, params, block, allOverrides)
	_, err = provider.Do(context.Background(), rpcCall.Into(&response))
	if err != nil {
		return nil, err
	}

	resultsData, err := hexutil.Decode(response)
	if err != nil {
		return nil, err
	}

	var results []SimulateResult
	err = contracts.V2.WalletGasEstimator.Decode(&results, "simulateExecute", resultsData)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func V3Simulate(provider *ethrpc.Provider, wallet common.Address, transactions Transactions, block string, overrides map[common.Address]*CallOverride) ([]SimulatorResult, error) {
	if block == "" {
		block = "latest"
	}

	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	calls := make([]walletestimator.PayloadCall, 0, len(encoded))
	for _, transaction := range encoded {
		behaviorOnError := big.NewInt(0)
		if transaction.RevertOnError {
			behaviorOnError.SetInt64(1)
		}
		calls = append(calls, walletestimator.PayloadCall{
			To:              transaction.To,
			Value:           transaction.Value,
			Data:            transaction.Data,
			GasLimit:        transaction.GasLimit,
			DelegateCall:    transaction.DelegateCall,
			OnlyFallback:    false,
			BehaviorOnError: behaviorOnError,
		})
	}

	data, err := contracts.V3.WalletSimulator.Encode("simulate", calls)
	if err != nil {
		return nil, err
	}

	type ethCallParams struct {
		To   common.Address `json:"to"`
		Data string         `json:"data"`
	}

	params := ethCallParams{
		To:   wallet,
		Data: hexutil.Encode(data),
	}

	allOverrides := map[common.Address]*CallOverride{
		wallet: {Code: walletGasSimulatorCodeV3},
	}
	for address, override := range overrides {
		if address == wallet {
			return nil, fmt.Errorf("cannot override wallet address %v", wallet.Hex())
		}

		allOverrides[address] = override
	}

	var responseHex string
	call := ethrpc.NewCallBuilder[string]("eth_call", nil, params, block, allOverrides)
	_, err = provider.Do(context.Background(), call.Into(&responseHex))
	if err != nil {
		return nil, err
	}

	response, err := hexutil.Decode(responseHex)
	if err != nil {
		return nil, err
	}

	var results []SimulatorResult
	err = contracts.V3.WalletSimulator.Decode(&results, "simulate", response)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func Simulate(provider *ethrpc.Provider, wallet common.Address, transactions Transactions, block string, overrides map[common.Address]*CallOverride) ([]SimulatorResult, error) {
	return V3Simulate(provider, wallet, transactions, block, overrides)
}
