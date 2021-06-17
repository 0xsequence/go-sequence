package testutil

import (
	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/contracts/gen/walletcallmock"
)

var (
	Contracts = ethartifact.NewContractRegistry()
)

func init() {
	Contracts.MustAdd(contracts.WalletFactory)

	Contracts.MustRegisterJSON("WALLET_CALL_RECV_MOCK", walletcallmock.CallReceiverMockABI, common.FromHex(walletcallmock.CallReceiverMockBin))
}

func ContractCall(provider *ethrpc.Provider, contractAddress common.Address, contractABI abi.ABI, method string, args ...interface{}) ([]string, []interface{}, error) {
	// weeeeee........... yeaaaa..........
	return nil, nil, nil
}

// .. keep..? maybe.. doesnt hurt..
func ContractQuery(provider *ethrpc.Provider, contractAddress common.Address, contractABI abi.ABI, inputExpr, outputExpr string, args []string) ([]string, []interface{}, error) {
	// weeeeee........... yeaaaa..........
	return nil, nil, nil
}

func ContractTransact(signer *ethwallet.Wallet, contractAddress common.Address, contractABI abi.ABI, method string, args ...interface{}) (*types.Transaction, error) {
	return nil, nil
}
