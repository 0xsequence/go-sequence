package sequence

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
)

type StateOverride struct {
	Key   string
	Value string
}

type Override struct {
	Code      string
	Balance   *big.Int
	Nonce     *big.Int
	StateDiff []*StateOverride
	State     []*StateOverride
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
}

var defaultEstimator = &Estimator{
	BaseCost:     21000,
	DataOneCost:  16,
	DataZeroCost: 4,
}

func NewEstimator() *Estimator {
	return &Estimator{
		BaseCost:     defaultEstimator.BaseCost,
		DataZeroCost: defaultEstimator.DataZeroCost,
		DataOneCost:  defaultEstimator.DataOneCost,
	}
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

func (e *Estimator) EstimateCall(ctx context.Context, provider *ethrpc.Provider, call *EstimateTransaction, overrides map[common.Address]*Override, blockTag string) (*big.Int, error) {
	if blockTag == "" {
		blockTag = "latest"
	}

	from := call.From
	if from.Hash().Big().Cmp(common.Big0) == 0 {
		from = stubAddress()
	}

	finalOverrides := map[common.Address]*Override{
		from: {
			Code: contracts.GasEstimatorDeployedBytecode,
		},
	}

	if overrides != nil {
		for key, value := range overrides {
			if key == call.From {
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
		reason := string(result[68 : len(result)-1])
		return gas, fmt.Errorf("error calling restimate: " + reason)
	}

	return gas, nil
}

func (e *Estimator) AreEOAs(ctx context.Context, provider *ethrpc.Provider, walletConfig WalletConfig) ([]bool, error) {
	// Get non-eoa signers
	// required for computing worse case scenario
	// TODO: Do these calls in parallel
	isEoa := make([]bool, walletConfig.Signers.Len())
	for i, s := range walletConfig.Signers {
		code, err := provider.CodeAt(ctx, s.Address, nil)
		if err != nil {
			return nil, err
		}

		isEoa[i] = len(code) == 0
	}

	return isEoa, nil
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
					Type:    sigPartTypeEOA,
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
					Type:    sigPartTypeDynamic,
					Value:   sig,
				}
			}
		} else {
			parts[i] = &SignaturePart{
				Weight:  s.Weight,
				Address: s.Address,
				Type:    sigPartTypeAddress,
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

func (e *Estimator) Estimate(ctx context.Context, provider *ethrpc.Provider, walletConfig WalletConfig, walletContext WalletContext, txs Transactions) (uint64, error) {
	address, err := AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return 0, err
	}

	isEOA, err := e.AreEOAs(ctx, provider, walletConfig)
	if err != nil {
		return 0, err
	}

	willSign, err := e.PickSigners(ctx, walletConfig, isEOA)
	if err != nil {
		return 0, err
	}

	signature := e.BuildStubSignature(walletConfig, willSign, isEOA)

	overrides := map[common.Address]*Override{
		walletContext.MainModuleAddress: {
			Code: contracts.WalletGasEstimatoreDeployedBytecode,
		},
		walletContext.MainModuleUpgradableAddress: {
			Code: contracts.WalletGasEstimatoreDeployedBytecode,
		},
	}

	estimates := make([]*big.Int, len(txs)+1)

	// Compute gas estimation for slices of all transactions
	// including no transaction execution and all transactions

	nonce, err := txs.Nonce()
	if err != nil {
		return 0, err
	}

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
