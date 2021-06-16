package testutil

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
)

var (
	Artifacts = contracts.ArtifactMap{
		"WALLET_FACTORY": {
			ABI: contracts.ABI_WALLET_FACTORY,
			Bin: common.FromHex(contracts.FactoryBin),
		},

		// TODO ......

	}
)

// TODO: include a bunch of other abi + bins here.. for ERC20_MOCK, etc....

// hmmpf..?
// TODO: lets have a ContractCall(transactor, contractAddress, abi.ABI, method, args ...interface{})
