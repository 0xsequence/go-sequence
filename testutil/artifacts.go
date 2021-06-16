package testutil

import (
	"github.com/0xsequence/go-sequence/contract"
)

var (
	Artifacts = contract.NewContractRegistry()
)

func init() {
	Artifacts.Register(contract.WalletFactory)

	// Artifacts.Register(&contract.ContractABI{
	// 	Name: "WALLET_FACTORY",
	// 	ABI:  contract.ABI_WALLET_FACTORY,
	// 	Bin:  common.FromHex(contracts.FactoryBin),
	// })

	// TODO: include a bunch of other abi + bins here.. for ERC20_MOCK, etc....
}

// hmmpf..?
// TODO: lets have a ContractCall(transactor, contractAddress, abi.ABI, method, args ...interface{})
