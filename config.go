package sequence

import (
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
)

func AddressFromWalletConfig(walletConfig core.WalletConfig, context WalletContext) (common.Address, error) {
	return AddressFromImageHash(walletConfig.ImageHash().Hex(), context)
}

func AddressFromImageHash(imageHash string, context WalletContext) (common.Address, error) {
	mainModule32 := [32]byte{}
	copy(mainModule32[12:], context.MainModuleAddress.Bytes())

	codePack, err := ethcoder.SolidityPack([]string{"bytes", "bytes32"}, []interface{}{walletContractBytecode, mainModule32})
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	codeHash := crypto.Keccak256(codePack)

	hashPack, err := ethcoder.SolidityPack(
		[]string{"bytes1", "address", "bytes32", "bytes32"},
		[]interface{}{[]byte{0xff}, context.FactoryAddress, common.FromHex(imageHash), codeHash},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	hash := crypto.Keccak256(hashPack)[12:]

	return common.BytesToAddress(hash), nil
}

func IsWalletConfigEqual(walletConfigA, walletConfigB core.WalletConfig) bool {
	return walletConfigA.ImageHash().Hex() == walletConfigB.ImageHash().Hex()
}

func V1SortWalletConfig(walletConfig *v1.WalletConfig) error {
	signers := walletConfig.Signers_
	sort.Sort(signers) // Sort the signers

	// Ensure no duplicates
	for i := 0; i < len(signers)-1; i++ {
		if signers[i].Address == signers[i+1].Address {
			return fmt.Errorf("signer duplicate detected in the wallet config")
		}
	}

	return nil
}
