package sequence

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/relayer/proto"
)

type RelayerSimulateResult struct {
	Executed  bool
	Succeeded bool
	Result    *string
	Reason    *string
	GasUsed   uint
	GasLimit  uint
}

type RelayerFeeTokenType uint32

const (
	UNKNOWN RelayerFeeTokenType = iota
	ERC20_TOKEN
	ERC1155_TOKEN
)

type RelayerFeeToken struct {
	ChainID         *big.Int
	Name            string
	Symbol          string
	Type            RelayerFeeTokenType
	Decimals        *uint32
	LogoURL         string
	ContractAddress *common.Address
	tokenID         *big.Int
}

type RelayerFeeOption struct {
	Token    RelayerFeeToken
	To       common.Address
	Value    *big.Int
	GasLimit *big.Int
}

type RelayerFeeQuote string

type Relayer interface {
	// ..
	GetProvider() *ethrpc.Provider

	// ..
	EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext WalletContext, txns Transactions) (Transactions, error)

	// ..
	Simulate(ctx context.Context, txs *SignedTransactions) ([]*RelayerSimulateResult, error)

	// NOTE: nonce space is 160 bits wide
	GetNonce(ctx context.Context, walletConfig core.WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error)

	// Relay will submit the Sequence signed meta transaction to the relayer. The method will block until the relayer
	// responds with the native transaction hash (*types.Transaction), which means the relayer has submitted the transaction
	// request to the network. Clients can use WaitReceipt to wait until the metaTxnID has been mined.
	Relay(ctx context.Context, signedTxs *SignedTransactions, quote *RelayerFeeQuote) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error)

	//
	FeeOptions(ctx context.Context, signedTxs *SignedTransactions) ([]*RelayerFeeOption, *RelayerFeeQuote, error)

	// ..
	Wait(ctx context.Context, metaTxnID MetaTxnID, optTimeout ...time.Duration) (MetaTxnStatus, *types.Receipt, error)

	// ..
	Client() proto.Relayer

	// TODO, in future when needed..
	// GasRefundOptions()
}

type MetaTxnID string

func (id MetaTxnID) String() string {
	return string(id)
}

type MetaTxnStatus uint8

const (
	MetaTxnStatusUnknown MetaTxnStatus = iota
	MetaTxnExecuted
	MetaTxnFailed
	MetaTxnReverted
)

// returns `to` address (either guest or wallet) and `data` of signed-metatx-calldata, aka execdata
func EncodeTransactionsForRelaying(relayer Relayer, walletAddress common.Address, walletConfig core.WalletConfig, walletContext WalletContext, txns Transactions, nonce *big.Int, seqSig []byte) (common.Address, []byte, error) {
	// TODO/NOTE: first version, we assume the wallet is deployed, then we can add bundlecreation after.
	// .....

	if len(txns) == 0 {
		return common.Address{}, nil, fmt.Errorf("cannot encode empty transactions")
	}

	// Encode transaction to be sent to a deployed wallet
	var err error
	if walletAddress == (common.Address{}) {
		walletAddress, err = AddressFromWalletConfig(walletConfig, walletContext)
		if err != nil {
			return common.Address{}, nil, err
		}
	}

	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return common.Address{}, nil, err
	}

	execdata, err := contracts.WalletMainModule.Encode("execute", encodedTxns, nonce, seqSig)
	if err != nil {
		return common.Address{}, nil, err
	}

	return walletAddress, execdata, nil
}

// DEPRECATED
// this method is horribly inefficient and we now have the new receipt_fetcher.go impl.
func LegacyWaitForMetaTxn(ctx context.Context, provider *ethrpc.Provider, metaTxnID MetaTxnID, optTimeout ...time.Duration) (MetaTxnStatus, *types.Receipt, error) {
	// Use optional timeout if passed, otherwise use deadline on the provided ctx, or finally,
	// set a default timeout of 120 seconds.
	var cancel context.CancelFunc
	if len(optTimeout) > 0 {
		ctx, cancel = context.WithTimeout(ctx, optTimeout[0])
		defer cancel()
	} else {
		if _, ok := ctx.Deadline(); !ok {
			ctx, cancel = context.WithTimeout(ctx, 120*time.Second)
			defer cancel()
		}
	}

	metaTxIdBytes := common.Hex2Bytes(string(metaTxnID))

	// All transactions must change nonces
	// so load all nonce changes and search the logs
	nonceChangedTopics := [][]common.Hash{{NonceChangeEventSig}}

	var fromBlock uint64

	// Load all logs until we found the receipt or we reach timeout
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.DeadlineExceeded) {
				return 0, nil, fmt.Errorf("waiting for meta transaction timeout for %v: %w", metaTxnID, err)
			} else if err != nil {
				return 0, nil, fmt.Errorf("failed waiting for meta transaction for %v: %w", metaTxnID, err)
			} else {
				return 0, nil, nil
			}
		default:
		}

		latestBlock, err := provider.BlockNumber(ctx)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		if fromBlock == 0 {
			del := uint64(200)      // TODO: make it an option..
			if latestBlock >= del { // clamp
				fromBlock = latestBlock - del
			} else {
				fromBlock = latestBlock
			}
		}

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(0).SetUint64(fromBlock),
			ToBlock:   big.NewInt(0).SetUint64(latestBlock),
			Topics:    nonceChangedTopics,
		}

		logs, err := provider.FilterLogs(ctx, query)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		for _, log := range logs {
			// NOTE: unfortunately we have to fetch the entire receipt of this NonceChanged receipt between
			// these blocks, so that we may then check the log data of that respective txn to find the
			// MetaTxnExecuted or MetaTxnFailed status.
			//
			// TODO: this method would HUGELY benefit to use goware/cachestore so we don't asks nodes
			// for the same transaction receipts we just fetched.
			tx, err := provider.TransactionReceipt(ctx, log.TxHash)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					break
				}
				time.Sleep(1500 * time.Millisecond)
				continue
			}

			for _, txLog := range tx.Logs {
				status := MetaTxnStatusUnknown

				// Success transactions have no topics and the metaTxId is the data
				if len(txLog.Topics) == 0 && bytes.Equal(txLog.Data, metaTxIdBytes) {
					status = MetaTxnExecuted
				}

				// Failed transactions have the TxFailed topic and the data begins with the metaTxInd
				if status == 0 && (len(txLog.Topics) == 1 && bytes.Equal(txLog.Topics[0].Bytes(), V1TxFailedEventSig.Bytes()) && bytes.HasPrefix(txLog.Data, metaTxIdBytes)) {
					status = MetaTxnFailed
				}

				if status > 0 {
					return status, tx, nil
				}
			}
		}

		// advance the cursor
		time.Sleep(time.Second)

		del := uint64(12)       // NOTE: we go back in case of reorgs, etc.
		if latestBlock >= del { // clamp
			fromBlock = latestBlock - del
		} else {
			fromBlock = latestBlock
		}
	}
}
