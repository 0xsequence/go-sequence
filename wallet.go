package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/sessions/proto"
)

var _ Signer = (*Wallet[core.WalletConfig])(nil)

type WalletOptions[C core.WalletConfig] struct {
	// Config is the wallet multi-sig configuration. Note: the first config of any wallet
	// before it is deployed is used to derive it's the account address of the wallet.
	Config C

	// Context is the WalletContext of deployed wallet-contract modules for the Smart Wallet.
	// NOTE: if a WalletContext is not provided, then `V1SequenceContext()` value is used.
	Context *WalletContext

	// Skips config sorting and keeps signers order as-is
	SkipSortSigners bool

	// Address used for the wallet
	// if this value is defined, the address derived from the sequence config is ignored
	Address common.Address
}

func GenericNewWallet[C core.WalletConfig](walletOptions WalletOptions[C], signers ...Signer) (*Wallet[C], error) {
	seqContext := SequenceContextForWalletConfig(walletOptions.Config)
	if walletOptions.Context != nil {
		seqContext = *walletOptions.Context
	}

	// Check if wallet config is usable
	if err := walletOptions.Config.IsUsable(); err != nil {
		return nil, fmt.Errorf("sequence.GenericNewWallet: %w", err)
	}

	// Generate address
	address := walletOptions.Address
	if address == (common.Address{}) {
		var err error
		address, err = AddressFromWalletConfig(walletOptions.Config, seqContext)
		if err != nil {
			return nil, fmt.Errorf("sequence.GenericNewWallet: %w", err)
		}
	}

	w := &Wallet[C]{
		config:          walletOptions.Config,
		context:         seqContext,
		address:         address,
		estimator:       NewEstimator(),
		skipSortSigners: walletOptions.SkipSortSigners,
	}
	w.signers = signers

	return w, nil
}

func V1NewWallet(walletOptions WalletOptions[*v1.WalletConfig], signers ...Signer) (*Wallet[*v1.WalletConfig], error) {
	return GenericNewWallet[*v1.WalletConfig](walletOptions, signers...)
}

func V2NewWallet(walletOptions WalletOptions[*v2.WalletConfig], signers ...Signer) (*Wallet[*v2.WalletConfig], error) {
	return GenericNewWallet[*v2.WalletConfig](walletOptions, signers...)
}

func V3NewWallet(walletOptions WalletOptions[*v3.WalletConfig], signers ...Signer) (*Wallet[*v3.WalletConfig], error) {
	return GenericNewWallet[*v3.WalletConfig](walletOptions, signers...)
}

func NewWallet(walletOptions WalletOptions[*v2.WalletConfig], signers ...Signer) (*Wallet[*v2.WalletConfig], error) {
	return V2NewWallet(walletOptions, signers...)
}

func GenericNewWalletSingleOwner[C core.WalletConfig](owner Signer, optContext ...WalletContext) (*Wallet[C], error) {
	var typeOfWallet C
	seqContext := SequenceContextForWalletConfig(typeOfWallet)
	if len(optContext) > 0 {
		seqContext = optContext[0]
	}

	if _, ok := core.WalletConfig(typeOfWallet).(*v1.WalletConfig); ok {
		// new wallet config v1
		var config core.WalletConfig = &v1.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Signers_: v1.WalletConfigSigners{
				{Weight: 1, Address: owner.Address()},
			},
		}

		// new sequence v1 wallet
		return GenericNewWallet[C](WalletOptions[C]{
			Config:  config.(C),
			Context: &seqContext,
		}, owner)
	} else if _, ok := core.WalletConfig(typeOfWallet).(*v2.WalletConfig); ok {
		// new wallet config v2
		var config core.WalletConfig = &v2.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: &v2.WalletConfigTreeAddressLeaf{
				Weight: 1, Address: owner.Address(),
			},
		}

		// new sequence v2 wallet
		return GenericNewWallet[C](WalletOptions[C]{
			Config:  config.(C),
			Context: &seqContext,
		}, owner)
	} else if _, ok := core.WalletConfig(typeOfWallet).(*v3.WalletConfig); ok {
		// new wallet config v3
		var config core.WalletConfig = &v3.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Tree: &v3.WalletConfigTreeAddressLeaf{
				Weight: 1, Address: owner.Address(),
			},
		}

		// new sequence v3 wallet
		return GenericNewWallet[C](WalletOptions[C]{
			Config:  config.(C),
			Context: &seqContext,
		}, owner)
	} else {
		return nil, fmt.Errorf("sequence.GenericNewWalletSingleOwner: unsupported wallet config type")
	}
}

// V1NewWalletSingleOwner creates a new Sequence v1 wallet with a single owner.
//
// Deprecated: use NewWalletSingleOwner instead. V1NewWalletSingleOwner is kept for historical
// compatibility and should not be used.
func V1NewWalletSingleOwner(owner Signer, optContext ...WalletContext) (*Wallet[*v1.WalletConfig], error) {
	return GenericNewWalletSingleOwner[*v1.WalletConfig](owner, optContext...)
}

// V2NewWalletSingleOwner creates a new Sequence v2 wallet with a single owner.
func V2NewWalletSingleOwner(owner Signer, optContext ...WalletContext) (*Wallet[*v2.WalletConfig], error) {
	return GenericNewWalletSingleOwner[*v2.WalletConfig](owner, optContext...)
}

// V3NewWalletSingleOwner creates a new Sequence v3 wallet with a single owner.
func V3NewWalletSingleOwner(owner Signer, optContext ...WalletContext) (*Wallet[*v3.WalletConfig], error) {
	return GenericNewWalletSingleOwner[*v3.WalletConfig](owner, optContext...)
}

// NewWalletSingleOwner creates a new Sequence v2 wallet with a single owner.
func NewWalletSingleOwner(owner Signer, optContext ...WalletContext) (*Wallet[*v2.WalletConfig], error) {
	return V2NewWalletSingleOwner(owner, optContext...)
}

func GenericNewWalletWithCoreWalletConfig[C core.WalletConfig](wallet *Wallet[C]) *Wallet[core.WalletConfig] {
	return &Wallet[core.WalletConfig]{
		context:         wallet.context,
		config:          wallet.config,
		signers:         wallet.signers,
		provider:        wallet.provider,
		relayer:         wallet.relayer,
		address:         wallet.address,
		skipSortSigners: wallet.skipSortSigners,
		chainID:         wallet.chainID,
	}
}

func V1NewWalletWithCoreWalletConfig(wallet *Wallet[*v1.WalletConfig]) *Wallet[core.WalletConfig] {
	return &Wallet[core.WalletConfig]{
		context:         wallet.context,
		config:          wallet.config,
		signers:         wallet.signers,
		provider:        wallet.provider,
		relayer:         wallet.relayer,
		address:         wallet.address,
		skipSortSigners: wallet.skipSortSigners,
		chainID:         wallet.chainID,
	}
}

func V2NewWalletWithCoreWalletConfig(wallet *Wallet[*v2.WalletConfig]) *Wallet[core.WalletConfig] {
	return &Wallet[core.WalletConfig]{
		context:         wallet.context,
		config:          wallet.config,
		signers:         wallet.signers,
		provider:        wallet.provider,
		relayer:         wallet.relayer,
		address:         wallet.address,
		skipSortSigners: wallet.skipSortSigners,
		chainID:         wallet.chainID,
	}
}

func V3NewWalletWithCoreWalletConfig(wallet *Wallet[*v3.WalletConfig]) *Wallet[core.WalletConfig] {
	return &Wallet[core.WalletConfig]{
		context:         wallet.context,
		config:          wallet.config,
		signers:         wallet.signers,
		provider:        wallet.provider,
		relayer:         wallet.relayer,
		address:         wallet.address,
		skipSortSigners: wallet.skipSortSigners,
		chainID:         wallet.chainID,
	}
}

func NewWalletWithCoreWalletConfig(wallet *Wallet[*v2.WalletConfig]) *Wallet[core.WalletConfig] {
	return V2NewWalletWithCoreWalletConfig(wallet)
}

type Wallet[C core.WalletConfig] struct {
	context WalletContext
	config  C
	signers []Signer

	provider  *ethrpc.Provider
	estimator *Estimator
	relayer   Relayer
	sessions  proto.Sessions
	address   common.Address

	skipSortSigners bool

	chainID *big.Int
}

var (
	ErrUnknownChainID = fmt.Errorf("chainID is unknown")
	ErrProviderNotSet = fmt.Errorf("provider is not set")
	ErrRelayerNotSet  = fmt.Errorf("relayer is not set")
)

func (w *Wallet[C]) UseConfig(config C) (*Wallet[C], error) {
	ww, err := GenericNewWallet(WalletOptions[C]{
		Config:          config,
		Context:         &w.context,
		SkipSortSigners: w.skipSortSigners,
		Address:         w.address,
	})

	if err != nil {
		return nil, fmt.Errorf("sequence.Wallet#UseConfig: %w", err)
	}
	if w.provider != nil {
		err = ww.SetProvider(w.provider)
		if err != nil {
			return nil, fmt.Errorf("sequence.Wallet#UseConfig setProvider: %w", err)
		}
	}
	if w.relayer != nil {
		err = ww.SetRelayer(w.relayer)
		if err != nil {
			return nil, fmt.Errorf("sequence.Wallet#UseConfig setRelayer: %w", err)
		}
	}
	return ww, nil
}

func (w *Wallet[C]) UseSigners(signers ...Signer) (*Wallet[C], error) {
	ww, err := GenericNewWallet(WalletOptions[C]{
		Config:          w.config,
		Context:         &w.context,
		SkipSortSigners: w.skipSortSigners,
		Address:         w.address,
	}, signers...)
	if err != nil {
		return nil, fmt.Errorf("sequence.Wallet#UseSigners: %w", err)
	}

	if w.provider != nil {
		err = ww.SetProvider(w.provider)
		if err != nil {
			return nil, fmt.Errorf("sequence.Wallet#UseSigners connect: %w", err)
		}
	}

	if w.relayer != nil {
		err = ww.SetRelayer(w.relayer)
		if err != nil {
			return nil, fmt.Errorf("sequence.Wallet#UseSigners connect: %w", err)
		}
	}

	return ww, nil
}

func (w *Wallet[C]) Connect(provider *ethrpc.Provider, relayer Relayer) error {
	err := w.SetProvider(provider)
	if err != nil {
		return err
	}
	err = w.SetRelayer(relayer)
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet[C]) SetProvider(provider *ethrpc.Provider) error {
	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("sequence.Wallet#SetProvider: %w", err)
	}
	w.chainID = chainID
	w.provider = provider
	return nil
}

func (w *Wallet[C]) SetRelayer(relayer Relayer) error {
	w.relayer = relayer
	return nil
}

func (w *Wallet[C]) SetSessions(sessions proto.Sessions) error {
	w.sessions = sessions
	return nil
}

func (w *Wallet[C]) UpdateSessionsWallet(ctx context.Context) error {
	if w.sessions == nil {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: sessions are not set")
	}

	var version int
	if _, ok := core.WalletConfig(w.config).(*v1.WalletConfig); ok {
		version = 1
	} else if _, ok := core.WalletConfig(w.config).(*v2.WalletConfig); ok {
		version = 2
	} else if _, ok := core.WalletConfig(w.config).(*v3.WalletConfig); ok {
		version = 3
	} else {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: unknown wallet config version")
	}

	err := w.sessions.SaveWallet(ctx, version, w.config)
	if err != nil {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: %w", err)
	}

	return nil
}

func (w *Wallet[C]) UpdateSessionsConfig(ctx context.Context) error {
	if w.sessions == nil {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: sessions are not set")
	}

	var version int
	if _, ok := core.WalletConfig(w.config).(*v1.WalletConfig); ok {
		version = 1
	} else if _, ok := core.WalletConfig(w.config).(*v2.WalletConfig); ok {
		version = 2
	} else if _, ok := core.WalletConfig(w.config).(*v3.WalletConfig); ok {
		version = 3
	} else {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: unknown wallet config version")
	}

	err := w.sessions.SaveConfig(ctx, version, w.config)
	if err != nil {
		return fmt.Errorf("sequence.Wallet#UpdateSessions: %w", err)
	}

	return nil
}

// SetChainID will set the wallet's associated chainID. However, for most part, this will automatically
// be set by the provider rpc.
func (w *Wallet[C]) SetChainID(chainID *big.Int) {
	w.chainID = chainID
}

func (w *Wallet[C]) GetProvider() *ethrpc.Provider {
	return w.provider
}

func (w *Wallet[C]) GetRelayer() Relayer {
	return w.relayer
}

func (w *Wallet[C]) GetChainID() *big.Int {
	return w.chainID
}

func (w *Wallet[C]) GetWalletContext() WalletContext {
	return w.context
}

func (w *Wallet[C]) GetWalletConfig() C {
	return w.config
}

func (w *Wallet[C]) Address() common.Address {
	return w.address
}

func (w *Wallet[C]) ImageHash() (common.Hash, error) {
	return w.config.ImageHash().Hash, nil
}

func (w *Wallet[C]) GetSignerAddresses() []common.Address {
	as := []common.Address{}
	for _, s := range w.signers {
		as = append(as, s.Address())
	}
	return as
}

func (w *Wallet[C]) IsSignerAvailable(address common.Address) bool {
	_, ok := w.GetSigner(address)
	return ok
}

func (w *Wallet[C]) GetSigner(address common.Address) (Signer, bool) {
	for _, s := range w.signers {
		if s.Address() == address {
			return s, true
		}
	}
	return nil, false
}

func (w *Wallet[C]) GetSignerWeight() *big.Int {
	var signers = make([]common.Address, 0, len(w.signers))
	for _, s := range w.signers {
		signers = append(signers, s.Address())
	}
	return big.NewInt(0).SetUint64(uint64(w.config.SignersWeight(signers)))
}

func (w *Wallet[C]) GetNonce(optBlockNum ...*big.Int) (*big.Int, error) {
	if w.relayer == nil {
		return nil, ErrRelayerNotSet
	}
	var blockNum *big.Int
	if len(optBlockNum) > 0 {
		blockNum = optBlockNum[0]
	}
	return w.relayer.GetNonce(context.Background(), w.config, w.context, nil, blockNum)
}

func (w *Wallet[C]) GetTransactionCount(optBlockNum ...*big.Int) (*big.Int, error) {
	return w.GetNonce(optBlockNum...)
}

func (w *Wallet[C]) SignTransactions(ctx context.Context, transactions Transactions, chainID ...*big.Int) (*SignedTransactions, error) {
	if chainID == nil {
		chainID = append(chainID, nil)
	}
	if chainID[0] == nil {
		chainID[0] = w.chainID
	}
	if chainID[0] == nil {
		return nil, fmt.Errorf("no chain id")
	}

	for _, transaction := range transactions {
		if !transaction.RevertOnError && (transaction.GasLimit == nil || transaction.GasLimit.Sign() <= 0) {
			var err error
			transactions, err = w.relayer.EstimateGasLimits(ctx, w.config, w.context, transactions)
			if err != nil {
				return nil, fmt.Errorf("unable to estimate gas limits: %w", err)
			}
			break
		}
	}

	spaceNonce, err := transactions.Nonce()
	if err != nil {
		return nil, fmt.Errorf("unable to sign transactions: %w", err)
	}
	if spaceNonce == nil {
		spaceNonce, err = w.GetNonce()
		if err != nil {
			return nil, fmt.Errorf("unable to get nonce: %w", err)
		}
	}
	space, nonce := DecodeNonce(spaceNonce)

	var subdigest core.Subdigest
	switch core.WalletConfig(w.config).(type) {
	case *v1.WalletConfig, *v2.WalletConfig:
		digest, err := ExecuteDigest(transactions, spaceNonce)
		if err != nil {
			return nil, fmt.Errorf("unable to compute transaction digest: %w", err)
		}

		subdigest = digest.Subdigest(w.address, chainID[0])

	case *v3.WalletConfig:
		payload, err := ConvertTransactionsToV3Payload(transactions, space, nonce)
		if err != nil {
			return nil, fmt.Errorf("unable to convert transactions to payload: %w", err)
		}

		subdigest, err = payload.Hash(w.address, chainID[0])
		if err != nil {
			return nil, fmt.Errorf("unable to compute transaction subdigest: %w", err)
		}

	default:
		return nil, fmt.Errorf("unknown wallet config type %T", w.config)
	}

	_, signature, err := w.Sign(ctx, subdigest)
	if err != nil {
		return nil, fmt.Errorf("unable to sign transactions: %w", err)
	}

	return &SignedTransactions{
		ChainID:       chainID[0],
		WalletAddress: w.address,
		WalletConfig:  w.config,
		WalletContext: w.context,
		Transactions:  transactions,
		Space:         space,
		Nonce:         nonce,
		Subdigest:     subdigest,
		Signature:     signature,
	}, nil
}

func (w *Wallet[C]) SignMessage(ctx context.Context, message []byte, chainID ...*big.Int) ([]byte, error) {
	if chainID == nil {
		chainID = append(chainID, nil)
	}
	if chainID[0] == nil {
		chainID[0] = w.chainID
	}
	if chainID[0] == nil {
		return nil, fmt.Errorf("no chain id")
	}

	var subdigest core.Subdigest
	switch core.WalletConfig(w.config).(type) {
	case *v1.WalletConfig, *v2.WalletConfig:
		subdigest = core.NewDigest(message).Subdigest(w.address, chainID[0])

	case *v3.WalletConfig:
		var err error
		subdigest, err = v3.Payload{Kind: v3.KindMessage, Message: message}.Hash(w.address, chainID[0])
		if err != nil {
			return nil, fmt.Errorf("unable to compute message subdigest: %w", err)
		}

	default:
		return nil, fmt.Errorf("unknown wallet config type %T", w.config)
	}

	_, signature, err := w.Sign(ctx, subdigest)
	return signature, err
}

func (w *Wallet[C]) SignTypedData(ctx context.Context, data *ethcoder.TypedData, chainID ...*big.Int) ([]byte, error) {
	digest, err := data.EncodeDigest()
	if err != nil {
		return nil, fmt.Errorf("unable to compute typed data digest: %w", err)
	}

	return w.SignHash(ctx, common.Hash(digest), chainID...)
}

func (w *Wallet[C]) SignHash(ctx context.Context, hash common.Hash, chainID ...*big.Int) ([]byte, error) {
	if chainID == nil {
		chainID = append(chainID, nil)
	}
	if chainID[0] == nil {
		chainID[0] = w.chainID
	}
	if chainID[0] == nil {
		return nil, fmt.Errorf("no chain id")
	}

	var subdigest core.Subdigest
	switch core.WalletConfig(w.config).(type) {
	case *v1.WalletConfig, *v2.WalletConfig:
		subdigest = core.Digest{Hash: hash}.Subdigest(w.address, chainID[0])

	case *v3.WalletConfig:
		var err error
		subdigest, err = v3.Payload{Kind: v3.KindDigest, Digest: hash}.Hash(w.address, chainID[0])
		if err != nil {
			return nil, fmt.Errorf("unable to compute digest subdigest: %w", err)
		}

	default:
		return nil, fmt.Errorf("unknown wallet config type %T", w.config)
	}

	_, signature, err := w.Sign(ctx, subdigest)
	return signature, err
}

func (w *Wallet[C]) SignConfiguration(ctx context.Context, imageHash core.ImageHash) ([]byte, error) {
	var subdigest core.Subdigest
	switch core.WalletConfig(w.config).(type) {
	case *v2.WalletConfig:
		subdigest = imageHash.Approval().Subdigest(w.address)

	case *v3.WalletConfig:
		var err error
		subdigest, err = v3.Payload{Kind: v3.KindConfigUpdate, ImageHash: imageHash}.Hash(w.address)
		if err != nil {
			return nil, fmt.Errorf("unable to compute configuration subdigest: %w", err)
		}

	default:
		return nil, fmt.Errorf("unknown wallet config type %T", w.config)
	}

	_, signature, err := w.Sign(ctx, subdigest)
	return signature, err
}

func (w *Wallet[C]) SignPayload(ctx context.Context, payload *v3.Payload, chainID ...*big.Int) ([]byte, error) {
	if chainID == nil {
		chainID = append(chainID, nil)
	}
	if chainID[0] == nil {
		chainID[0] = w.chainID
	}
	if chainID[0] == nil {
		return nil, fmt.Errorf("no chain id")
	}

	switch core.WalletConfig(w.config).(type) {
	case *v1.WalletConfig, *v2.WalletConfig:
		return nil, fmt.Errorf("refusing to sign v3 payload with non-v3 wallet")

	case *v3.WalletConfig:
		subdigest, err := payload.Hash(w.address, chainID[0])
		if err != nil {
			return nil, fmt.Errorf("unable to compute payload subdigest: %w", err)
		}

		_, signature, err := w.Sign(ctx, subdigest)
		return signature, err

	default:
		return nil, fmt.Errorf("unknown wallet config type %T", w.config)
	}
}

func (w *Wallet[C]) Sign(ctx context.Context, subdigest core.Subdigest) (core.SignerSignatureType, []byte, error) {
	if w.sessions != nil {
		err := w.UpdateSessionsConfig(ctx)
		if err != nil {
			return 0, nil, fmt.Errorf("unable to save config: %w", err)
		}
	}

	sign := func(ctx context.Context, signerAddress common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		signer, _ := w.GetSigner(signerAddress)
		if signer == nil {
			return 0, nil, core.ErrSigningNoSigner
		}
		return signer.Sign(ctx, subdigest)
	}

	switch config := core.WalletConfig(w.config).(type) {
	case *v1.WalletConfig:
		signature, err := config.BuildSignature(ctx, sign)
		if err != nil {
			return 0, nil, fmt.Errorf("unable to build signature: %w", err)
		}

		data, err := signature.Data()
		if err != nil {
			return 0, nil, fmt.Errorf("unable to encode signature: %w", err)
		}

		return core.SignerSignatureTypeEIP1271, data, nil

	case *v2.WalletConfig:
		var signature core.Signature[*v2.WalletConfig]
		var err error
		if subdigest.ChainID.Sign() == 0 {
			signature, err = config.BuildNoChainIDSignature(ctx, sign)
		} else {
			signature, err = config.BuildRegularSignature(ctx, sign)
		}
		if err != nil {
			return 0, nil, fmt.Errorf("unable to build signature: %w", err)
		}

		data, err := signature.Data()
		if err != nil {
			return 0, nil, fmt.Errorf("unable to encode signature: %w", err)
		}

		return core.SignerSignatureTypeEIP1271, data, nil

	case *v3.WalletConfig:
		var signature core.Signature[*v3.WalletConfig]
		var err error
		if subdigest.ChainID.Sign() == 0 {
			signature, err = config.BuildNoChainIDSignature(ctx, sign)
		} else {
			signature, err = config.BuildRegularSignature(ctx, sign)
		}
		if err != nil {
			return 0, nil, fmt.Errorf("unable to build signature: %w", err)
		}

		data, err := signature.Data()
		if err != nil {
			return 0, nil, fmt.Errorf("unable to encode signature: %w", err)
		}

		return core.SignerSignatureTypeEIP1271, data, nil

	default:
		return 0, nil, fmt.Errorf("unknown wallet config type %T", config)
	}
}

func (w *Wallet[C]) SendTransaction(ctx context.Context, signedTxns *SignedTransactions, feeQuote ...*RelayerFeeQuote) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	return w.SendTransactions(ctx, signedTxns, feeQuote...)
}

func (w *Wallet[C]) SendTransactions(ctx context.Context, signedTxns *SignedTransactions, feeQuote ...*RelayerFeeQuote) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	if w.relayer == nil {
		return "", nil, nil, ErrRelayerNotSet
	}

	return w.relayer.Relay(ctx, signedTxns, feeQuote...)
}

func (w *Wallet[C]) FeeOptions(ctx context.Context, txs Transactions) ([]*RelayerFeeOption, *RelayerFeeQuote, error) {
	if w.relayer == nil {
		return []*RelayerFeeOption{}, nil, ErrRelayerNotSet
	}

	// prepare for signed txs
	nonce, err := txs.Nonce()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot load nonce from transactions: %w", err)
	}

	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return nil, nil, err
		}
	}

	bundle := Transaction{
		Transactions: txs,
		Nonce:        nonce,
	}

	// get transactions digest
	digest, err := bundle.Digest()
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get digest from transactions: %w", err)
	}

	// prepare for fee estimation
	areEOAs, err := w.estimator.AreEOAs(ctx, w.provider, w.config)
	if err != nil {
		return nil, nil, fmt.Errorf("estimator areEOAs error: %w", err)
	}

	willSign, err := w.estimator.PickSigners(ctx, w.config, areEOAs)
	if err != nil {
		return nil, nil, fmt.Errorf("estimator pickSigners error: %w", err)
	}

	sig := w.estimator.BuildStubSignature(w.config, willSign, areEOAs)

	// signed txs
	signedTxs := &SignedTransactions{
		ChainID:       w.chainID,
		WalletAddress: w.address,
		WalletConfig:  w.config,
		WalletContext: w.context,
		Transactions:  txs,
		Nonce:         nonce,
		Subdigest:     digest.Subdigest(w.address, w.chainID),
		Signature:     sig,
	}

	// get fee options
	return w.relayer.FeeOptions(ctx, signedTxs)
}

func (w *Wallet[C]) IsDeployed() (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}
	return IsWalletDeployed(w.provider, w.Address())
}

func (w *Wallet[C]) Deploy(ctx context.Context) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	if w.relayer == nil {
		return "", nil, nil, ErrRelayerNotSet
	}

	walletAddress, walletFactoryAddress, deploymentData, err := EncodeWalletDeployment(w.config, w.context)
	if err != nil {
		return "", nil, nil, err
	}

	if w.address != (common.Address{}) && w.address != walletAddress {
		return "", nil, nil, fmt.Errorf("wallet address %s does not match the address derived from the config %s", w.address, walletAddress)
	}

	var txn = &Transaction{
		RevertOnError: true,
		To:            walletFactoryAddress,
		Data:          deploymentData,
	}

	// for v1 sequence wallet default does not work
	if _, ok := core.WalletConfig(w.config).(*v1.WalletConfig); ok {
		// TODO: Move this hardcoded gas limit to a configuration
		// or fix it with a contract patch
		txn.GasLimit = big.NewInt(3_000_000)
	} else if _, ok := core.WalletConfig(w.config).(*v3.WalletConfig); ok {
		// TODO: Move this hardcoded gas limit to a configuration
		// or fix it with a contract patch
		txn.GasLimit = big.NewInt(3_000_000)
	}

	signerTxn, err := w.SignTransactions(ctx, Transactions{txn})
	if err != nil {
		return "", nil, nil, err
	}

	return w.relayer.Relay(ctx, signerTxn)
}

// func (w *Wallet) UpdateConfig() // TODO in future

// func (w *Wallet) PublishConfig() // TODO in future

func (w *Wallet[C]) IsValidSignature(digest common.Hash, signature []byte) (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}

	// todo: this is a hack to get around the fact that the signature verification is not available in WalletConfig
	var generalWalletConfig core.WalletConfig = w.config
	if _, ok := generalWalletConfig.(*v3.WalletConfig); ok {
		sig, err := v3.Core.DecodeSignature(signature)
		if err != nil {
			return false, err
		}

		subdigest, err := v3.Payload{Kind: v3.KindDigest, Digest: digest}.Hash(w.address, w.chainID)
		if err != nil {
			return false, err
		}

		config, weight, err := sig.Recover(context.Background(), subdigest, w.provider)
		if err != nil {
			return false, err
		} else {
			return weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) >= 0, nil
		}
	} else if _, ok := generalWalletConfig.(*v2.WalletConfig); ok {
		sig, err := v2.Core.DecodeSignature(signature)
		if err != nil {
			return false, err
		}

		config, weight, err := sig.Recover(context.Background(), core.Digest{Hash: digest}.Subdigest(w.address, w.chainID), w.provider)
		if err != nil {
			return false, err
		} else {
			return weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) >= 0, nil
		}
	} else if _, ok := generalWalletConfig.(*v1.WalletConfig); ok {
		sig, err := v1.Core.DecodeSignature(signature)
		if err != nil {
			return false, err
		}

		config, weight, err := sig.Recover(context.Background(), core.Digest{Hash: digest}.Subdigest(w.address, w.chainID), w.provider)
		if err != nil {
			return false, err
		} else {
			return weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) >= 0, nil
		}
	} else {
		return false, fmt.Errorf("unknown wallet config type")
	}
}

func IsWalletDeployed(provider *ethrpc.Provider, walletAddress common.Address) (bool, error) {
	code, err := provider.CodeAt(context.Background(), walletAddress, nil)
	if err != nil {
		return false, err
	}
	if len(code) > 2 {
		return true, nil
	}
	return false, nil
}

// AuxData is the data that is signed by the wallet
type AuxData struct {
	Msg     []byte
	Sig     []byte
	ChainID *big.Int
	Address *common.Address
}

func (a *AuxData) Pack() ([]byte, error) {
	return ethcoder.AbiCoder(
		[]string{"address", "uint256", "bytes", "bytes"},
		[]interface{}{a.Address, a.ChainID, &a.Msg, &a.Sig},
	)
}

func ContextWithAuxData(ctx context.Context, auxData *AuxData) context.Context {
	return context.WithValue(ctx, "auxData", auxData)
}

func AuxDataFromContext(ctx context.Context) (*AuxData, error) {
	auxData, ok := ctx.Value("auxData").(*AuxData)
	if !ok {
		return nil, fmt.Errorf("auxData not found in context")
	}
	return auxData, nil
}
