package testutil

import (
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contract"
	"github.com/0xsequence/go-sequence/contract/gen/walletcallmock"
)

var (
	Contracts = contract.NewContractRegistry()
)

func init() {
	Contracts.MustAdd(contract.WalletFactory)

	Contracts.MustRegisterJSON("WALLET_CALL_RECV_MOCK", walletcallmock.CallReceiverMockABI, common.FromHex(walletcallmock.CallReceiverMockBin))
}
