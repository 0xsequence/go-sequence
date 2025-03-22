package sequence

import (
	"context"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/core"
	v3 "github.com/0xsequence/go-sequence/core/v3"
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
	TokenID         *big.Int
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
	Relay(ctx context.Context, signedTxs *SignedTransactions, quote ...*RelayerFeeQuote) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error)

	//
	FeeOptions(ctx context.Context, signedTxs *SignedTransactions) ([]*RelayerFeeOption, *RelayerFeeQuote, error)

	// ..
	Wait(ctx context.Context, payload v3.PayloadDigestable) (MetaTxnStatus, *types.Receipt, error)

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
