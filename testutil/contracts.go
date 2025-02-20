package testutil

import (
	"context"

	"github.com/0xsequence/ethkit/ethartifact"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	walletcallmock "github.com/0xsequence/go-sequence/contracts/gen/v1/walletcallmock"
)

var (
	Contracts = ethartifact.NewContractRegistry()
)

func init() {
	Contracts.MustAdd(contracts.ERC20Mock)
	Contracts.MustAdd(contracts.WalletFactory)

	Contracts.MustRegisterJSON("WALLET_CALL_RECV_MOCK", walletcallmock.CallReceiverMockABI, common.FromHex(walletcallmock.CallReceiverMockBin))
}

func ContractCall(provider *ethrpc.Provider, contractAddress common.Address, contractABI abi.ABI, result interface{}, method string, args ...interface{}) ([]byte, error) {
	calldata, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: calldata,
	}

	output, err := provider.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return output, nil
	}

	err = contractABI.UnpackIntoInterface(result, method, output)
	if err != nil {
		return output, err
	}
	return output, nil
}

func ContractQuery(provider *ethrpc.Provider, contractAddress common.Address, inputExpr, outputExpr string, args []string) ([]string, error) {
	return provider.ContractQuery(context.Background(), contractAddress.Hex(), inputExpr, outputExpr, args)
}

func ContractTransact(signer *ethwallet.Wallet, contractAddress common.Address, contractABI abi.ABI, method string, args ...interface{}) (*types.Receipt, error) {
	calldata, err := contractABI.Pack(method, args...)
	if err != nil {
		return nil, err
	}

	signedTxn, err := signer.NewTransaction(context.Background(), &ethtxn.TransactionRequest{
		To:   &contractAddress,
		Data: calldata,
	})
	if err != nil {
		return nil, err
	}

	_, waitReceipt, err := signer.SendTransaction(context.Background(), signedTxn)
	if err != nil {
		return nil, err
	}

	receipt, err := waitReceipt(context.Background())
	if err != nil {
		return nil, err
	}

	return receipt, nil
}
