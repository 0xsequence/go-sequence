package sequence

import (
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type WalletOptions struct {
	// Config is the wallet multi-sig configuration. Note: the first config of any wallet
	// before it is deployed is used to derive it's the account address of the wallet.
	Config WalletConfig

	// Context is the WalletContext of deployed wallet-contract modules for the Smart Wallet.
	// NOTE: if a WalletContext is not provided, then `SequenceContext()` value is used.
	Context *WalletContext
}

func NewWallet(walletOptions WalletOptions, signers ...*ethwallet.Wallet) (*Wallet, error) {
	context := sequenceContext
	if walletOptions.Context != nil {
		context = *walletOptions.Context
	}

	walletConfig := walletOptions.Config
	err := SortWalletConfig(walletConfig)
	if err != nil {
		return nil, fmt.Errorf("sequence.NewWallet: %w", err)
	}

	_, err = IsWalletConfigUsable(walletConfig)
	if err != nil {
		return nil, fmt.Errorf("sequence.NewWallet: %w", err)
	}

	w := &Wallet{
		config: walletConfig, context: context,
	}
	w.signers = signers

	return w, nil
}

func NewWalletSingleOwner(owner *ethwallet.Wallet, optContext ...WalletContext) (*Wallet, error) {
	context := sequenceContext
	if len(optContext) > 0 {
		context = optContext[0]
	}

	return NewWallet(WalletOptions{
		Config: WalletConfig{
			Threshold: 1, //big.NewInt(1),
			Signers: WalletConfigSigners{
				{Weight: 1, Address: owner.Address()},
			},
		},
		Context: &context,
	}, owner)
}

type Wallet struct {
	context WalletContext
	config  WalletConfig
	signers []*ethwallet.Wallet // NOTE: only supports EOA signers at this time

	provider *ethrpc.Provider
	relayer  Relayer

	chainID *big.Int
}

var (
	ErrUnknownChainID = fmt.Errorf("chainID is unknown")
	ErrProviderNotSet = fmt.Errorf("provider is not set")
	ErrRelayerNotSet  = fmt.Errorf("relayer is not set")
)

func (w *Wallet) UseConfig(config WalletConfig) (*Wallet, error) {
	ww, err := NewWallet(WalletOptions{
		Config: config, Context: &w.context,
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

func (w *Wallet) UseSigners(signers ...*ethwallet.Wallet) (*Wallet, error) {
	ww, err := NewWallet(WalletOptions{
		Config: w.config, Context: &w.context,
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

func (w *Wallet) Connect(provider *ethrpc.Provider, relayer Relayer) error {
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

func (w *Wallet) SetProvider(provider *ethrpc.Provider) error {
	chainID, err := provider.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("sequence.Wallet#SetProvider: %w", err)
	}
	w.chainID = chainID
	w.provider = provider
	return nil
}

func (w *Wallet) SetRelayer(relayer Relayer) error {
	w.relayer = relayer
	return nil
}

// SetChainID will set the wallet's associated chainID. However, for most part, this will automatically
// be set by the provider rpc.
func (w *Wallet) SetChainID(chainID *big.Int) {
	w.chainID = chainID
}

func (w *Wallet) GetProvider() *ethrpc.Provider {
	return w.provider
}

func (w *Wallet) GetRelayer() Relayer {
	return w.relayer
}

func (w *Wallet) GetChainID() *big.Int {
	return w.chainID
}

func (w *Wallet) GetWalletContext() WalletContext {
	return w.context
}

func (w *Wallet) GetWalletConfig() WalletConfig {
	return w.config
}

func (w *Wallet) Address() common.Address {
	hs, err := ImageHashOfWalletConfig(w.config)
	if err != nil {
		return common.Address{}
	}
	as, err := AddressFromImageHash(hs, w.context)
	if err != nil {
		return common.Address{}
	}
	return as
}

func (w *Wallet) ImageHash() (common.Hash, error) {
	hs, err := ImageHashOfWalletConfig(w.config)
	if err != nil {
		return common.Hash{}, err
	}
	return common.HexToHash(hs), nil
}

func (w *Wallet) GetSignerAddresses() []common.Address {
	as := []common.Address{}
	for _, s := range w.signers {
		as = append(as, s.Address())
	}
	return as
}

func (w *Wallet) IsSignerAvailable(address common.Address) bool {
	_, ok := w.GetSigner(address)
	return ok
}

func (w *Wallet) GetSigner(address common.Address) (*ethwallet.Wallet, bool) {
	for _, s := range w.signers {
		if s.Address() == address {
			return s, true
		}
	}
	return nil, false
}

func (w *Wallet) GetSignerWeight() *big.Int {
	signerAddresses := w.GetSignerAddresses()
	totalWeight := big.NewInt(0)
	for _, sa := range signerAddresses {
		for _, sc := range w.config.Signers {
			if sc.Address == sa {
				totalWeight = big.NewInt(0).Add(totalWeight, big.NewInt(0).SetUint64(uint64(sc.Weight)))
			}
		}
	}
	return totalWeight
}

func (w *Wallet) GetNonce(blockNum ...big.Int) *big.Int {
	// TODO: ask teh relayer..? not sure why were asking the relayer, but okay
	return nil
}

func (w *Wallet) GetTransactionCount(blockNum ...big.Int) *big.Int {
	return w.GetNonce(blockNum...)
}

func (w *Wallet) SignMessage(msg []byte) ([]byte, error) {
	return w.SignDigest(MessageDigest(msg))
}

func (w *Wallet) SignDigest(digest []byte) ([]byte, error) {
	if w.chainID == nil {
		return nil, fmt.Errorf("sequence.Wallet#SignDigest: %w", ErrUnknownChainID)
	}

	subDigest, err := SubDigest(w.Address(), w.chainID, digest)
	if err != nil {
		return nil, fmt.Errorf("SignDigest, subDigestOf: %w", err)
	}

	sig := &Signature{
		Threshold: w.config.Threshold,
		Signers:   SignatureParts{},
	}

	for _, signerInfo := range w.config.Signers {
		signer, _ := w.GetSigner(signerInfo.Address)

		if signer == nil {
			// signer isn't available, just include the config value of address
			// without it's signature
			sig.Signers = append(sig.Signers, &SignaturePart{
				Type: sigPartTypeAddress, Weight: signerInfo.Weight, Address: signerInfo.Address,
			})
			continue
		}

		// TODO: in future add support for Sequence Signers, right now we only support EOA
		// signers in go-sequence

		sigValue, err := signer.SignMessage(subDigest)
		if err != nil {
			return nil, fmt.Errorf("signer.SignMessage subDigest: %w", err)
		}
		sigValue = append(sigValue, sigTypeEthSign)

		sig.Signers = append(sig.Signers, &SignaturePart{
			Type: sigPartTypeEOA, Weight: signerInfo.Weight, Address: signer.Address(), Value: sigValue,
		})
	}

	encodedSig, err := sig.Encode()
	if err != nil {
		return nil, err
	}

	return encodedSig, nil
}

// func (w *Wallet) SignTypedData() // TODO

func (w *Wallet) SignTransaction(txn *Transaction) (interface{}, error) {
	return nil, nil
}

func (w *Wallet) SignTransactionBatch(txns Transactions) (interface{}, error) {
	return nil, nil
}

// TODO: SendTransaction()

// TODO: SendTransactionBatch()

func (w *Wallet) Deploy() error {
	if w.relayer == nil {
		return ErrRelayerNotSet
	}
	// TODO: .. via a relayer..
	return nil
}

func (w *Wallet) IsDeployed() (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}
	code, err := w.provider.CodeAt(context.Background(), w.Address(), nil)
	if err != nil {
		return false, err
	}
	if len(code) > 2 {
		return true, nil
	}
	return false, nil
}

// func (w *Wallet) UpdateConfig() // TODO in future

// func (w *Wallet) PublishConfig() // TODO in future

func (w *Wallet) IsValidSignature(digest, signature []byte) (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}
	return IsValidSignature(w.Address(), digest, signature, w.context, w.chainID, w.provider)
}

var (
	// ImageHashUpdatedEventSig is emitted anytime wallet config is updated.
	ImageHashUpdatedEventSig = MustEncodeSig("ImageHashUpdated(bytes32)")

	// ImplementationUpdatedEventSig is emitted anytime a wallet's mainModule is changed,
	// this is a rare occurence.
	ImplementationUpdatedEventSig = MustEncodeSig("ImplementationUpdated(address)")
)