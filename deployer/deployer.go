package deployer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contract"
)

// Sequence Universal Deployer constant values, do not change these.
//
// See: https://github.com/0xsequence/sequence.js/blob/master/packages/deployer/src/constants.ts
// which are the identical values to the ones below.

var (
	EOA_UNIVERSAL_DEPLOYER_ADDRESS = common.HexToAddress("0x9c5a87452d4FAC0cbd53BDCA580b20A45526B3AB")
	UNIVERSAL_DEPLOYER_ADDRESS     = common.HexToAddress("0x1b926fbb24a9f78dcdd3272f2d86f5d0660e59c0")
	UNIVERSAL_DEPLOYER_2_ADDRESS   = common.HexToAddress("0x8a5bc19e22d6ad55a2c763b93a75d09f321fe764")
	UNIVERSAL_DEPLOYER_FUNDING     = big.NewInt(1).Mul(big.NewInt(300), big.NewInt(100_000_000_000_000))
	UNIVERSAL_DEPLOYER_TX          = "0xf9010880852416b84e01830222e08080b8b66080604052348015600f57600080fd5b50609980601d6000396000f3fe60a06020601f369081018290049091028201604052608081815260009260609284918190838280828437600092018290525084519495509392505060208401905034f5604080516001600160a01b0383168152905191935081900360200190a0505000fea26469706673582212205a310755225e3c740b2f013fb6343f4c205e7141fcdf15947f5f0e0e818727fb64736f6c634300060a00331ca01820182018201820182018201820182018201820182018201820182018201820a01820182018201820182018201820182018201820182018201820182018201820"

	// expected bytecode for the universal deployer 2. If this changes for whatever reason then the universal
	// deployer's addresses of contracts it's deployed will change, and so will UNIVERSAL_DEPLOYER_2_ADDRESS.
	//
	// do not change this value. it is here to integrity check within the UniversalDeployer. if you do change
	// it however, then make sure to also update UNIVERSAL_DEPLOYER_2_ADDRESS.
	//
	// this value was originally copied from typings/contracts/factories/UniversalDeployer2__factory.ts
	// from https://github.com/0xsequence/sequence.js/blob/master/packages/deployer
	UNIVERSAL_DEPLOYER_2_BYTECODE = common.FromHex("0x608060405234801561001057600080fd5b5061013d806100206000396000f3fe60806040526004361061001e5760003560e01c80639c4ae2d014610023575b600080fd5b6100cb6004803603604081101561003957600080fd5b81019060208101813564010000000081111561005457600080fd5b82018360208201111561006657600080fd5b8035906020019184600183028401116401000000008311171561008857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506100cd915050565b005b60008183516020850134f56040805173ffffffffffffffffffffffffffffffffffffffff83168152905191925081900360200190a050505056fea264697066735822122033609f614f03931b92d88c309d698449bb77efcd517328d341fa4f923c5d8c7964736f6c63430007060033")

	UNIVERSAL_DEPLOYER_2_ABI = contract.MustParseABI(`[
    {
      "anonymous": true,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "_addr",
          "type": "address"
        }
      ],
      "name": "Deploy",
      "type": "event"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "_creationCode",
          "type": "bytes"
        },
        {
          "internalType": "uint256",
          "name": "_instance",
          "type": "uint256"
        }
      ],
      "name": "deploy",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    }
  ]`)
)

type UniversalDeployer struct {
	Wallet   *ethwallet.Wallet
	provider *ethrpc.Provider
}

func NewUniversalDeployer(wallet *ethwallet.Wallet) (*UniversalDeployer, error) {
	ud := &UniversalDeployer{Wallet: wallet}

	if wallet.GetProvider() == nil {
		return nil, fmt.Errorf("deployer: wallet provider is not set")
	}
	ud.provider = wallet.GetProvider()

	return ud, nil
}

func (u *UniversalDeployer) Deploy(ctx context.Context, contractABI *abi.ABI, contractBin []byte, contractInstance uint, txParams interface{}, contractArgs ...interface{}) (common.Address, error) {
	// Deploy universal deployer 2 if not yet deployed
	code, err := u.provider.CodeAt(ctx, UNIVERSAL_DEPLOYER_2_ADDRESS, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: %w", err)
	}
	if len(code) == 0 {
		err = u.deployUniversalDeployer2(ctx, txParams)
		if err != nil {
			return common.Address{}, err
		}
	}

	// Deploying contract

	// first compute the deterministic address of the contract to-be
	contractAddress, deployData, err := ComputeCreate2Address(contractABI, contractBin, contractInstance, contractArgs...)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: %w", err)
	}

	code, err = u.provider.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer getCode: %w", err)
	}
	if len(code) > 0 {
		// contract is already deployed, we done
		return contractAddress, nil
	}

	// Deploy the contract via UniversalDeployer2 by calling the deploy contract method
	input, err := UNIVERSAL_DEPLOYER_2_ABI.Pack("deploy", deployData, big.NewInt(int64(contractInstance)))
	if err != nil {
		return common.Address{}, fmt.Errorf("deployer: deploy pack: %w", err)
	}

	tx, _, err := u.Wallet.NewTransaction(ctx, &ethwallet.TransactionRequest{
		To:   &UNIVERSAL_DEPLOYER_2_ADDRESS,
		Data: input,
	})

	_, waitTx, err := u.Wallet.SendTransaction(ctx, tx)
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

	return contractAddress, nil
}

// TODO: perhaps update txParams to *ethwallet.TransactionRequest or *ethtx.TransactionRequest,
// and intersect those values, or clone it etc. Presently txParams is ignored.

// Deploy universal deployer via universal deployer 1
func (u *UniversalDeployer) deployUniversalDeployer2(ctx context.Context, txParams interface{}) error {
	code, err := u.provider.CodeAt(ctx, UNIVERSAL_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return fmt.Errorf("deloyer: %w", err)
	}
	if len(code) == 0 {
		err = u.deployUniversalDeployer1(ctx, txParams)
		if err != nil {
			return err
		}
	}

	// NOTE: in case the getCode below fails, double check the UNIVERSAL_DEPLOYER_2_ADDRESS address
	// which is emitted from the deployer 1 contract creation logs. This address may change if
	// the UNIVERSAL_DEPLOYER_2_BYTECODE changes of the deployer -- which should never really happen.

	// deploy universal deployer 2 contract
	tx, _, err := u.Wallet.NewTransaction(ctx, &ethwallet.TransactionRequest{
		To:   &UNIVERSAL_DEPLOYER_ADDRESS,
		Data: UNIVERSAL_DEPLOYER_2_BYTECODE,
	})
	if err != nil {
		return err
	}

	_, waitTx, err := u.Wallet.SendTransaction(ctx, tx)
	if err != nil {
		return err
	}
	receipt, err := waitTx(ctx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("deployer (deployUniversalDeployer2): txn failed")
	}

	return nil
}

func (u *UniversalDeployer) deployUniversalDeployer1(ctx context.Context, txParams interface{}) error {
	deployerBalance, err := u.provider.BalanceAt(ctx, EOA_UNIVERSAL_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return fmt.Errorf("deployer (deployUniversalDeployer1): %w", err)
	}

	// ensure deployer's wallet balance has necessary funding
	if deployerBalance.Cmp(UNIVERSAL_DEPLOYER_FUNDING) < 0 {
		signedTxn, _, err := u.Wallet.NewTransaction(ctx, &ethwallet.TransactionRequest{
			To:       &EOA_UNIVERSAL_DEPLOYER_ADDRESS,
			ETHValue: UNIVERSAL_DEPLOYER_FUNDING,
		})
		if err != nil {
			return err
		}
		_, waitTx, err := u.Wallet.SendTransaction(ctx, signedTxn)
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
	provider := u.Wallet.GetProvider()

	txHash, err := provider.SendRawTransaction(ctx, UNIVERSAL_DEPLOYER_TX)
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
	universalDeployerCodeCheck, err := provider.CodeAt(ctx, UNIVERSAL_DEPLOYER_ADDRESS, nil)
	if err != nil {
		return fmt.Errorf("deployer: %w", err)
	}
	if universalDeployerCodeCheck == nil || len(universalDeployerCodeCheck) == 0 {
		return fmt.Errorf("deployer (deployUniversalDeployer1): failed to deploy stage 1 deployer")
	}

	return nil
}

func ComputeCreate2Address(contractABI *abi.ABI, contractBin []byte, contractInstance uint, contractArgs ...interface{}) (common.Address, []byte, error) {
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
		[]interface{}{[1]byte{255}, UNIVERSAL_DEPLOYER_2_ADDRESS, salt, codeHash},
	)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("ComputeCreate2Address hashPack: %w", err)
	}
	hash := crypto.Keccak256(hashPack)[12:]

	return common.BytesToAddress(hash), deployData, nil
}
