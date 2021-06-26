package sequence

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
)

type Relayer interface {
	// ..
	GetProvider() *ethrpc.Provider

	// ..
	EstimateGasLimits(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, txns Transactions) (Transactions, error)

	// NOTE: nonce space is 160 bits wide
	GetNonce(ctx context.Context, walletConfig WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error)

	// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
	// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
	// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
	Relay(ctx context.Context, signedTxs *SignedTransactions) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error)

	// ..
	Wait(ctx context.Context, metaTxnID MetaTxnID, timeout time.Duration) (MetaTxnStatus, *types.Receipt, error)

	// TODO, in future when needed..
	// GasRefundOptions()
}

type MetaTxnID string

type MetaTxnStatus uint8

const (
	MetaTxnStatusUnknown MetaTxnStatus = iota
	MetaTxnExecuted
	MetaTxnFailed
)

func ComputeMetaTxnID(walletAddress common.Address, chainID *big.Int, txns Transactions, nonce *big.Int) (MetaTxnID, error) {
	bundle := Transaction{
		Transactions: txns,
		Nonce:        nonce,
	}

	txnsDigest, err := bundle.Digest()
	if err != nil {
		return "", err
	}

	metaSubDigest, err := SubDigest(walletAddress, chainID, txnsDigest)
	if err != nil {
		return "", err
	}

	metaTxnIDHex := ethcoder.HexEncode(metaSubDigest)
	if len(metaTxnIDHex) != 66 {
		return "", fmt.Errorf("computed meta txn id is invalid length")
	}

	return MetaTxnID(metaTxnIDHex[2:]), nil
}

// returns `to` address (either guest or wallet) and `data` of signed-metatx-calldata, aka execdata
func EncodeTransactionsForRelaying(relayer Relayer, walletConfig WalletConfig, walletContext WalletContext, txns Transactions, nonce *big.Int, seqSig []byte) (common.Address, []byte, error) {
	// TODO/NOTE: first version, we assume the wallet is deployed, then we can add bundlecreation after.
	// .....

	if len(txns) == 0 {
		return common.Address{}, nil, fmt.Errorf("cannot encode empty transactions")
	}

	// Encode transaction to be sent to a deployed wallet
	walletAddress, err := AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return common.Address{}, nil, err
	}

	execdata, err := contracts.WalletMainModule.Encode("execute", txns.AsValues(), nonce, seqSig)
	if err != nil {
		return common.Address{}, nil, err
	}

	return walletAddress, execdata, nil
}

func WaitForMetaTxn(ctx context.Context, provider *ethrpc.Provider, metaTxnID MetaTxnID, timeout time.Duration) (MetaTxnStatus, *types.Receipt, error) {
	timeoutTime := time.Now().Add(timeout)

	// Start listening logs from current block - 1024
	block, err := provider.BlockNumber(ctx)
	if err != nil {
		return 0, nil, err
	}

	// TODO: Move the - 1024 hardcoded value to an option
	lastBlockNumber := block - 1024

	metaTxIdBytes := common.Hex2Bytes(string(metaTxnID))

	// All transactions must change nonces
	// so load all nonce changes and search the logs
	nonceChangedTopics := [][]common.Hash{{NonceChangeEventSig}}

	// Load all logs until we found the receipt or we reach timeout
	for time.Now().Before(timeoutTime) {
		block, err := provider.BlockNumber(ctx)

		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		query := ethereum.FilterQuery{
			// TODO: Move the - 12 hardcoded value to an option
			FromBlock: big.NewInt(int64(lastBlockNumber) - 12),
			ToBlock:   big.NewInt(int64(block)),
			Topics:    nonceChangedTopics,
		}

		logs, err := provider.FilterLogs(ctx, query)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		for _, log := range logs {
			tx, err := provider.TransactionReceipt(ctx, log.TxHash)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			for _, txLog := range tx.Logs {
				status := MetaTxnStatusUnknown

				// Success transactions have no topics and the metaTxId is the data
				if len(txLog.Topics) == 0 && bytes.Equal(txLog.Data, metaTxIdBytes) {
					status = MetaTxnExecuted
				}

				// Failed transactions have the TxFailed topic and the data begins with the metaTxInd
				if status == 0 && (len(txLog.Topics) == 1 && bytes.Equal(txLog.Topics[0].Bytes(), TxFailedEventSig.Bytes()) && bytes.HasPrefix(txLog.Data, metaTxIdBytes)) {
					status = MetaTxnFailed
				}

				if status > 0 {
					return status, tx, nil
				}
			}
		}

		time.Sleep(time.Second)
		lastBlockNumber = block
	}

	return 0, nil, fmt.Errorf("Waiting for meta transaction timeout %v", metaTxnID)
}
