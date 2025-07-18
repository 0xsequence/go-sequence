package sequence

import (
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
)

func AddressFromWalletConfig(deployConfig core.ImageHashable, context WalletContext) (common.Address, error) {
	return AddressFromImageHash(deployConfig.ImageHash(), context)
}

func AddressFromImageHash(deployHash core.ImageHash, context WalletContext) (common.Address, error) {
	creationCode, err := hexutil.Decode(context.CreationCode)
	if err != nil {
		return common.Address{}, fmt.Errorf(`invalid creation code "%v": %w`, context.CreationCode, err)
	}

	var mainModule common.Hash
	mainModule.SetBytes(context.MainModuleAddress.Bytes())
	return crypto.CreateAddress2(context.FactoryAddress, deployHash.Hash, crypto.Keccak256(creationCode, mainModule.Bytes())), nil
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
