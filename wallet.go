package sequence

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethtxn"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v2 "github.com/0xsequence/go-sequence/core/v2"
	v3 "github.com/0xsequence/go-sequence/core/v3"
	"github.com/0xsequence/go-sequence/sessions/proto"
)

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

	return &Wallet[C]{
		context:         seqContext,
		config:          walletOptions.Config,
		signers:         signers,
		address:         address,
		estimator:       NewEstimator(),
		skipSortSigners: walletOptions.SkipSortSigners,
	}, nil
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

var _ Signer = (*Wallet[core.WalletConfig])(nil)

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
	})
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
	ww.signers = signers
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

func (w *Wallet[C]) ImageHash() (core.ImageHash, error) {
	return w.config.ImageHash(), nil
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

func (w *Wallet[C]) Nonce(ctx context.Context, space *big.Int) (*big.Int, error) {
	if w.provider == nil {
		return nil, ErrProviderNotSet
	}
	return Nonce(ctx, w.address, space, w.provider)
}

func (w *Wallet[C]) GetTransactionCount(ctx context.Context) (*big.Int, error) {
	return w.Nonce(ctx, nil)
}

func (w *Wallet[C]) SignMessage(ctx context.Context, message []byte) ([]byte, error) {
	_, signature, err := w.Sign(ctx, v3.ConstructMessagePayload(w.address, w.chainID, message))
	return signature, err
}

func (w *Wallet[C]) SignTypedData(typedData *ethcoder.TypedData) ([]byte, []byte, error) {
	digest, encodedTypedData, err := typedData.Encode()
	if err != nil {
		return nil, nil, err
	}
	signature, err := w.SignDigest(context.Background(), common.Hash(digest))
	if err != nil {
		return nil, nil, err
	}
	return signature, encodedTypedData, nil
}

func (w *Wallet[C]) SignDigest(ctx context.Context, digest common.Hash) ([]byte, error) {
	_, signature, err := w.Sign(ctx, v3.ConstructDigestPayload(w.address, w.chainID, digest))
	return signature, err
}

func (w *Wallet[C]) Sign(ctx context.Context, payload v3.Payload) (core.SignerSignatureType, []byte, error) {
	sign := func(ctx context.Context, signerAddress common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		signer, _ := w.GetSigner(signerAddress)
		if signer == nil {
			// signer isn't available, just include the config value of address
			// without it's signature
			return 0, nil, core.ErrSigningNoSigner
		}

		return signer.Sign(ctx, payload)
	}

	signature, _, err := w.buildSignature(ctx, sign, payload.ChainID())
	return core.SignerSignatureTypeEIP1271, signature, err
}

func (w *Wallet[C]) SignTransaction(ctx context.Context, txn v3.Call, space ...*big.Int) (*SignedTransactions, error) {
	if len(space) == 0 {
		space = append(space, nil)
	}
	return w.SignTransactions(ctx, Transactions{Calls: []v3.Call{txn}, Space: space[0]})
}

func (w *Wallet[C]) SignTransactions(ctx context.Context, txns Transactions) (*SignedTransactions, error) {
	if len(txns.Calls) == 0 {
		return nil, fmt.Errorf("cannot sign an empty set of transactions")
	}

	var err error

	// If a transaction has 0 gasLimit and not revertOnError
	// compute all new gas limits
	estimateGas := false
	for _, txn := range txns.Calls {
		if txn.BehaviorOnError != v3.BehaviorOnErrorRevert && (txn.GasLimit == nil || txn.GasLimit.Cmp(big.NewInt(0)) == 0) {
			estimateGas = true
			break
		}
	}
	if estimateGas {
		txns, err = w.relayer.EstimateGasLimits(ctx, w.config, w.context, txns)
		if err != nil {
			return nil, fmt.Errorf("estimateGas failed for sequence transactions: %w", err)
		}
	}

	var space, nonce *big.Int
	if txns.Nonce != nil {
		if txns.Space != nil {
			space, nonce = txns.Space, txns.Nonce
		} else {
			space, nonce = common.Big0, txns.Nonce
		}
	} else {
		if txns.Space != nil {
			space = txns.Space
			nonce, err = w.Nonce(ctx, space)
			if err != nil {
				return nil, fmt.Errorf("unable to get nonce for space %v: %w", txns.Space, err)
			}
		} else {
			var data [20]byte
			rand.Read(data[:])
			space, nonce = new(big.Int).SetBytes(data[:]), common.Big0
		}
	}

	switch core.WalletConfig(w.config).(type) {
	case *v3.WalletConfig:
		payload := Transactions{Calls: txns.Calls, Space: space, Nonce: nonce}.Payload(w.address, w.chainID)

		// Sign the transactions
		_, signature, err := w.Sign(ctx, payload)
		if err != nil {
			return nil, err
		}

		return &SignedTransactions{CallsPayload: payload, Signature: signature}, nil
	}

	return nil, fmt.Errorf("unknown wallet config type")
}

func (w *Wallet[C]) SendTransaction(ctx context.Context, signedTxns *SignedTransactions, feeQuote ...*RelayerFeeQuote) (*types.Transaction, ethtxn.WaitReceipt, error) {
	return w.SendTransactions(ctx, signedTxns, feeQuote...)
}

func (w *Wallet[C]) SendTransactions(ctx context.Context, signedTxns *SignedTransactions, feeQuote ...*RelayerFeeQuote) (*types.Transaction, ethtxn.WaitReceipt, error) {
	if w.relayer == nil {
		return nil, nil, ErrRelayerNotSet
	}

	return w.relayer.Relay(ctx, signedTxns, w.config, feeQuote...)
}

func (w *Wallet[C]) FeeOptions(ctx context.Context, txs Transactions) ([]*RelayerFeeOption, *RelayerFeeQuote, error) {
	if w.relayer == nil {
		return []*RelayerFeeOption{}, nil, ErrRelayerNotSet
	}

	var space, nonce *big.Int
	if txs.Nonce != nil {
		if txs.Space != nil {
			space, nonce = txs.Space, txs.Nonce
		} else {
			space, nonce = common.Big0, txs.Nonce
		}
	} else {
		if txs.Space != nil {
			var err error
			space = txs.Space
			nonce, err = w.Nonce(ctx, space)
			if err != nil {
				return nil, nil, fmt.Errorf("unable to get nonce for space %v: %w", txs.Space, err)
			}
		} else {
			var data [20]byte
			rand.Read(data[:])
			space, nonce = new(big.Int).SetBytes(data[:]), common.Big0
		}
	}

	payload := v3.ConstructCallsPayload(w.address, w.chainID, txs.Calls, space, nonce)

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
	signedTxs := &SignedTransactions{CallsPayload: payload, Signature: sig}

	// get fee options
	return w.relayer.FeeOptions(ctx, signedTxs)
}

func (w *Wallet[C]) IsDeployed() (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}
	return IsWalletDeployed(w.provider, w.Address())
}

func (w *Wallet[C]) Deploy(ctx context.Context, space ...*big.Int) (*SignedTransactions, *types.Transaction, ethtxn.WaitReceipt, error) {
	if w.relayer == nil {
		return nil, nil, nil, ErrRelayerNotSet
	}

	walletAddress, walletFactoryAddress, deploymentData, err := EncodeWalletDeployment(w.config, w.context)
	if err != nil {
		return nil, nil, nil, err
	}

	if w.address != (common.Address{}) && w.address != walletAddress {
		return nil, nil, nil, fmt.Errorf("wallet address %s does not match the address derived from the config %s", w.address, walletAddress)
	}

	var txn = v3.Call{
		To:              walletFactoryAddress,
		Data:            deploymentData,
		BehaviorOnError: v3.BehaviorOnErrorRevert,
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

	signerTxn, err := w.SignTransaction(ctx, txn, space...)
	if err != nil {
		return nil, nil, nil, err
	}

	transaction, waitReceipt, err := w.relayer.Relay(ctx, signerTxn, w.config)
	return signerTxn, transaction, waitReceipt, err
}

// func (w *Wallet) UpdateConfig() // TODO in future

// func (w *Wallet) PublishConfig() // TODO in future

func (w *Wallet[C]) IsValidSignature(payload v3.PayloadDigestable, signature []byte) (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}

	digest := core.Digest{Hash: payload.Digest().Hash}

	// todo: this is a hack to get around the fact that the signature verification is not available in WalletConfig
	var generalWalletConfig core.WalletConfig = w.config
	if _, ok := generalWalletConfig.(*v3.WalletConfig); ok {
		sig, err := v3.Core.DecodeSignature(signature)
		if err != nil {
			return false, err
		}

		config, weight, err := sig.Recover(context.Background(), digest, w.address, w.chainID, w.provider)
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

		config, weight, err := sig.Recover(context.Background(), digest, w.address, w.chainID, w.provider)
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

		config, weight, err := sig.Recover(context.Background(), digest, w.address, w.chainID, w.provider)
		if err != nil {
			return false, err
		} else {
			return weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) >= 0, nil
		}
	} else {
		return false, fmt.Errorf("unknown wallet config type")
	}
}

func (w *Wallet[C]) buildSignature(ctx context.Context, sign core.SigningFunction, chainID *big.Int) ([]byte, core.Signature[C], error) {
	var coreWalletConfig core.WalletConfig = w.config

	if config, ok := coreWalletConfig.(*v1.WalletConfig); ok {
		sig, err := config.BuildSignature(ctx, sign, false)
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, BuildSignature: %w", err)
		}

		sigEnc, err := sig.Data()
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, sig.Data: %w", err)
		}

		sigTyped, _ := sig.(core.Signature[C])
		// todo: implement core.Signature[core.WalletConfig] wrapper
		return sigEnc, sigTyped, nil
	} else if config, ok := coreWalletConfig.(*v2.WalletConfig); ok {
		var (
			sig core.Signature[*v2.WalletConfig]
			err error
		)
		if chainID.Sign() == 0 {
			sig, err = config.BuildNoChainIDSignature(ctx, sign, false)
			if err != nil {
				return nil, nil, fmt.Errorf("SignDigest, BuildNoChainIDSignature: %w", err)
			}
		} else {
			sig, err = config.BuildRegularSignature(ctx, sign, false)
			if err != nil {
				return nil, nil, fmt.Errorf("SignDigest, BuildRegularSignature: %w", err)
			}
		}

		sigEnc, err := sig.Data()
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, sig.Data: %w", err)
		}

		sigTyped, _ := sig.(core.Signature[C])
		// todo: implement core.Signature[core.WalletConfig] wrapper
		return sigEnc, sigTyped, nil
	} else if config, ok := coreWalletConfig.(*v3.WalletConfig); ok {
		var (
			sig core.Signature[*v3.WalletConfig]
			err error
		)
		if chainID.Sign() == 0 {
			sig, err = config.BuildNoChainIDSignature(ctx, sign, false)
			if err != nil {
				return nil, nil, fmt.Errorf("SignDigest, BuildNoChainIDSignature: %w", err)
			}
		} else {
			sig, err = config.BuildRegularSignature(ctx, sign, false)
			if err != nil {
				return nil, nil, fmt.Errorf("SignDigest, BuildRegularSignature: %w", err)
			}
		}

		sigEnc, err := sig.Data()
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, sig.Data: %w", err)
		}

		sigTyped, _ := sig.(core.Signature[C])
		// todo: implement core.Signature[core.WalletConfig] wrapper
		return sigEnc, sigTyped, nil
	} else {
		return nil, nil, fmt.Errorf("SignDigest, unknown config type")
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

func Nonce(ctx context.Context, wallet common.Address, space *big.Int, provider *ethrpc.Provider) (*big.Int, error) {
	if space == nil {
		space = common.Big0
	}

	code, err := provider.CodeAt(ctx, wallet, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to get code at %v: %w", wallet, err)
	}
	if len(code) == 0 {
		return new(big.Int), nil
	}

	calldata, err := contracts.V3.WalletStage1Module.Encode("readNonce", space)
	if err != nil {
		return nil, fmt.Errorf("unable to encode calldata: %w", err)
	}

	result, err := provider.CallContract(ctx, ethereum.CallMsg{To: &wallet, Data: calldata}, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to read nonce: %w", err)
	}

	var nonce *big.Int
	err = contracts.V3.WalletStage1Module.Decode(&nonce, "readNonce", result)
	if err != nil {
		return nil, fmt.Errorf("unable to decode nonce: %w", err)
	}

	return nonce, nil
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

var (
	// ImageHashUpdatedEventSig is emitted anytime wallet config is updated.
	ImageHashUpdatedEventSig = MustEncodeSig("ImageHashUpdated(bytes32)")

	// ImplementationUpdatedEventSig is emitted anytime a wallet's mainModule is changed,
	// this is a rare occurence.
	ImplementationUpdatedEventSig = MustEncodeSig("ImplementationUpdated(address)")
)
