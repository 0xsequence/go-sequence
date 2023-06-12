package deployer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

var (
	EIP2740_EOA_DEPLOYER_ADDRESS = common.HexToAddress("0xBb6e024b9cFFACB947A71991E386681B1Cd1477D")
	EIP2740_DEPLOYER_ADDRESS     = common.HexToAddress("0xce0042B868300000d44A59004Da54A005ffdcf9f")
	EIP2740_DEPLOYER_TX          = "0xf9016c8085174876e8008303c4d88080b90154608060405234801561001057600080fd5b50610134806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80634af63f0214602d575b600080fd5b60cf60048036036040811015604157600080fd5b810190602081018135640100000000811115605b57600080fd5b820183602082011115606c57600080fd5b80359060200191846001830284011164010000000083111715608d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550509135925060eb915050565b604080516001600160a01b039092168252519081900360200190f35b6000818351602085016000f5939250505056fea26469706673582212206b44f8a82cb6b156bfcc3dc6aadd6df4eefd204bc928a4397fd15dacf6d5320564736f6c634300060200331b83247000822470"
	EIP2740_DEPLOYER_FUNDING     = big.NewInt(1).Mul(big.NewInt(300), big.NewInt(100_000_000_000_000))

	EIP2740_DEPLOYER_ABI = ethcontract.MustParseABI(`[
    {
        "constant": false,
        "inputs": [
            {
                "internalType": "bytes",
                "name": "_initCode",
                "type": "bytes"
            },
            {
                "internalType": "bytes32",
                "name": "_salt",
                "type": "bytes32"
            }
        ],
        "name": "deploy",
        "outputs": [
            {
                "internalType": "address payable",
                "name": "createdContract",
                "type": "address"
            }
        ],
        "payable": false,
        "stateMutability": "nonpayable",
        "type": "function"
    }
  ]`)
)

type EIP2740Deployer struct {
	Wallet   *ethwallet.Wallet
	provider *ethrpc.Provider
}

func NewEIP2740Deployer(wallet *ethwallet.Wallet) (*EIP2740Deployer, error) {
	e := &EIP2740Deployer{
		Wallet: wallet,
	}

	if wallet.GetProvider() == nil {
		return nil, fmt.Errorf("deployer: wallet provider is not set")
	}
	e.provider = wallet.GetProvider()

	return e, nil
}

func (e *EIP2740Deployer) Deploy(ctx context.Context, contractABI abi.ABI, contractBin []byte, contractInstance uint, txParams interface{}, gasLimit uint, contractArgs ...interface{}) (common.Address, error) {
	// Deploy universal deployer 2 if not yet deployed
	code, err := e.provider.CodeAt(ctx, EIP2740_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: %w", err)
	}
	if len(code) == 0 {
		err = e.deployEIP2740Deployer(ctx, txParams)
		if err != nil {
			return common.Address{}, err
		}

		code, err = e.provider.CodeAt(ctx, UNIVERSAL_DEPLOYER_2_ADDRESS, nil)
		if len(code) == 0 {
			return common.Address{}, fmt.Errorf("can't deploy eip2740 deployer")
		}
	}

	// Deploying contract

	// first compute the deterministic address of the contract to-be
	contractAddress, deployData, err := e.ComputeCreate2Address(contractABI, contractBin, contractInstance, contractArgs...)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: %w", err)
	}

	code, err = e.provider.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer getCode: %w", err)
	}
	if len(code) > 0 {
		// contract is already deployed, we done
		return contractAddress, nil
	}

	// Deploy the contract via EIP2740Deployer by calling the deploy contract method
	var salt [32]byte
	big.NewInt(int64(contractInstance)).SetBytes(salt[:])
	input, err := EIP2740_DEPLOYER_ABI.Pack("deploy", deployData, salt)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: deploy pack: %w", err)
	}

	tx, err := e.Wallet.NewTransaction(ctx, &ethtxn.TransactionRequest{
		To:       &EIP2740_DEPLOYER_ADDRESS,
		Data:     input,
		GasLimit: uint64(gasLimit),
	})

	_, waitTx, err := e.Wallet.SendTransaction(ctx, tx)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: deploy call: %w", err)
	}

	receipt, err := waitTx(ctx)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: deploy tx receipt wait: %w", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, fmt.Errorf("deployer: tx failed: %w", err)
	}

	code, err = e.provider.CodeAt(ctx, contractAddress, nil)
	if len(code) == 0 {
		return common.Address{}, fmt.Errorf("can't deploy contract")
	}

	return contractAddress, nil
}

func (e *EIP2740Deployer) deployEIP2740Deployer(ctx context.Context, txParams interface{}) error {
	deployerBalance, err := e.provider.BalanceAt(ctx, EIP2740_EOA_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return fmt.Errorf("deployer (deployUniversalDeployer1): %w", err)
	}

	// ensure deployer's wallet balance has necessary funding
	if deployerBalance.Cmp(EIP2740_DEPLOYER_FUNDING) < 0 {
		signedTxn, err := e.Wallet.NewTransaction(ctx, &ethtxn.TransactionRequest{
			To:       &EIP2740_EOA_DEPLOYER_ADDRESS,
			ETHValue: EIP2740_DEPLOYER_FUNDING,
		})
		if err != nil {
			return err
		}
		_, waitTx, err := e.Wallet.SendTransaction(ctx, signedTxn)
		if err != nil {
			return err
		}
		receipt, err := waitTx(ctx)
		if err != nil {
			return err
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("deployer: txn failed")
		}
	}

	// deploying universal deployer contract
	provider := e.Wallet.GetProvider()

	txHash, err := provider.SendRawTransaction(ctx, EIP2740_DEPLOYER_TX)
	if err != nil {
		return err
	}

	receipt, err := ethrpc.WaitForTxnReceipt(ctx, provider, txHash)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("deployer: txn failed")
	}

	// lets check the universal deployer address is now deployed, after txn above
	universalDeployerCodeCheck, err := provider.CodeAt(ctx, EIP2740_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return fmt.Errorf("deployer: %w", err)
	}
	if universalDeployerCodeCheck == nil || len(universalDeployerCodeCheck) == 0 {
		return fmt.Errorf("deployer (deployUniversalDeployer1): failed to deploy stage 1 deployer")
	}

	return nil
}

func (e *EIP2740Deployer) ComputeCreate2Address(contractABI abi.ABI, contractBin []byte, contractInstance uint, contractArgs ...interface{}) (common.Address, []byte, error) {
	var input []byte
	var err error

	if len(contractArgs) > 0 && len(contractABI.Constructor.Inputs) > 0 {
		input, err = contractABI.Pack("", contractArgs...)
	} else {
		input, err = contractABI.Pack("")
	}
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("ComputeCreate2Address pack: %w", err)
	}

	deployData := append(contractBin, input...)
	codeHash := crypto.Keccak256(deployData)

	salt, err := ethcoder.SolidityPack([]string{"uint256"}, []interface{}{big.NewInt(int64(contractInstance))})
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("ComputeCreate2Address salt: %w", err)
	}

	hashPack, err := ethcoder.SolidityPack(
		[]string{"bytes1", "address", "bytes32", "bytes32"},
		[]interface{}{[1]byte{255}, EIP2740_DEPLOYER_ADDRESS, salt, codeHash},
	)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("ComputeCreate2Address hashPack: %w", err)
	}
	hash := crypto.Keccak256(hashPack)[12:]

	return common.BytesToAddress(hash), deployData, nil
}
