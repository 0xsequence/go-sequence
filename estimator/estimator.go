package estimator

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/ethclient/gethclient"
	"github.com/0xsequence/go-sequence/contracts"
)

type AddressPreconditions struct {
	Balance *big.Int
	ERC20   map[common.Address]ERC20Preconditions
	ERC721  map[common.Address]ERC721Preconditions
	ERC1155 map[common.Address]ERC1155Preconditions
}

type ERC20Preconditions struct {
	Balance   *big.Int
	Allowance map[common.Address]*big.Int
}

type ERC721Preconditions struct {
	IsOwner    map[string]bool
	IsApproved map[string]bool
}

type ERC1155Preconditions struct {
	Balance          map[string]*big.Int
	IsApprovedForAll map[common.Address]bool
}

func Estimate(
	ctx context.Context,
	to common.Address,
	data []byte,
	preconditions map[common.Address]AddressPreconditions,
	provider *ethrpc.Provider,
) (uint64, error) {
	overrides := map[common.Address]gethclient.OverrideAccount{}

	for address, addressPreconditions := range preconditions {
		if addressPreconditions.Balance != nil && addressPreconditions.Balance.Sign() > 0 {
			overrides[address] = gethclient.OverrideAccount{Balance: addressPreconditions.Balance}
		}
	}

	for {
		list, err := accessList(ctx, to, data, overrides, provider)
		if err != nil {
			return 0, err
		}

		if !override(overrides, preconditions, list) {
			break
		}
	}

	var from common.Address
	rand.Read(from[:])

	overrides[from] = gethclient.OverrideAccount{Code: contracts.GasEstimator.DeployedBin}

	input, err := contracts.GasEstimator.Encode("estimate", to, data)
	if err != nil {
		return 0, fmt.Errorf("unable to encode estimate call: %w", err)
	}

	var encoded string
	_, err = provider.Do(ctx, ethrpc.NewCallBuilder[string]("eth_call", nil, map[string]any{"to": from, "input": hexutil.Encode(input)}, "latest", overrides).Into(&encoded))
	if err != nil {
		return 0, fmt.Errorf("unable to estimate: %w", err)
	}

	decoded, err := hexutil.Decode(encoded)
	if err != nil {
		return 0, fmt.Errorf("unable to decode estimate result: %w", err)
	}

	var result struct {
		Success bool     `abi:"success"`
		Result  []byte   `abi:"result"`
		Gas     *big.Int `abi:"gas"`
	}

	err = contracts.GasEstimator.Decode(&result, "estimate", decoded)
	if err != nil {
		return 0, fmt.Errorf("unable to decode estimate result: %w", err)
	}
	if !result.Success {
		message, err := abi.UnpackRevert(result.Result)
		if err == nil {
			return 0, fmt.Errorf("simulation with preconditions failed: %v (gas: %v)", message, result.Gas)
		} else {
			return 0, fmt.Errorf("simulation with preconditions failed: %v (gas: %v)", hexutil.Encode(result.Result), result.Gas)
		}
	}

	return result.Gas.Uint64(), nil
}

func override(
	overrides map[common.Address]gethclient.OverrideAccount,
	preconditions map[common.Address]AddressPreconditions,
	list types.AccessList,
) bool {
	var updated bool

	for _, slots := range list {
		for _, addressPreconditions := range preconditions {
			if erc20Preconditions, ok := addressPreconditions.ERC20[slots.Address]; ok {
				if update(overrides, slots, erc20Preconditions.Balance) {
					updated = true
				}
				for _, allowance := range erc20Preconditions.Allowance {
					if update(overrides, slots, allowance) {
						updated = true
					}
				}
			}

			if erc721Preconditions, ok := addressPreconditions.ERC721[slots.Address]; ok {
				for _, isOwner := range erc721Preconditions.IsOwner {
					if isOwner {
						if update(overrides, slots, common.Big1) {
							updated = true
						}
						break
					}
				}
				for _, isApproved := range erc721Preconditions.IsApproved {
					if isApproved {
						if update(overrides, slots, common.Big1) {
							updated = true
						}
						break
					}
				}
			}

			if erc1155Preconditions, ok := addressPreconditions.ERC1155[slots.Address]; ok {
				for _, balance := range erc1155Preconditions.Balance {
					if update(overrides, slots, balance) {
						updated = true
					}
				}
				for _, isApprovedForAll := range erc1155Preconditions.IsApprovedForAll {
					if isApprovedForAll {
						if update(overrides, slots, common.Big1) {
							updated = true
						}
						break
					}
				}
			}
		}
	}

	return updated
}

func update(overrides map[common.Address]gethclient.OverrideAccount, slots types.AccessTuple, value *big.Int) bool {
	if value == nil || value.Sign() <= 0 {
		return false
	}
	value_ := common.BigToHash(value)

	var updated bool
	addressOverrides := overrides[slots.Address]

	for slot, existing := range addressOverrides.StateDiff {
		if existing.Big().Cmp(value) < 0 {
			if addressOverrides.StateDiff == nil {
				addressOverrides.StateDiff = map[common.Hash]common.Hash{}
			}
			addressOverrides.StateDiff[slot] = value_
			updated = true
		}
	}

	for _, slot := range slots.StorageKeys {
		if addressOverrides.StateDiff[slot].Big().Cmp(value) < 0 {
			if addressOverrides.StateDiff == nil {
				addressOverrides.StateDiff = map[common.Hash]common.Hash{}
			}
			addressOverrides.StateDiff[slot] = value_
			updated = true
		}
	}

	overrides[slots.Address] = addressOverrides
	return updated
}

func accessList(
	ctx context.Context,
	to common.Address,
	data []byte,
	overrides map[common.Address]gethclient.OverrideAccount,
	provider *ethrpc.Provider,
) (types.AccessList, error) {
	into := func(raw json.RawMessage, ret *types.AccessList, strictness ethrpc.StrictnessLevel) error {
		var response struct {
			AccessList types.AccessList `json:"accessList,omitempty"`
		}

		err := json.Unmarshal(raw, &response)
		if err != nil {
			return fmt.Errorf("unable to decode eth_createAccessList response: %w", err)
		}

		*ret = response.AccessList
		return nil
	}

	var list types.AccessList
	_, err := provider.Do(ctx, ethrpc.NewCallBuilder(
		"eth_createAccessList",
		into,
		map[string]any{"to": to, "input": hexutil.Encode(data)},
		"latest",
		overrides,
	).Into(&list))
	if err != nil {
		_, err := provider.Do(ctx, ethrpc.NewCallBuilder(
			"eth_createAccessList",
			into,
			map[string]any{"to": to, "input": hexutil.Encode(data)},
			"latest",
		).Into(&list))
		if err != nil {
			return nil, fmt.Errorf("unable to create access list: %w", err)
		}
	}

	return list, nil
}
