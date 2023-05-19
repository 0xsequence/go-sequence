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
)

type WalletOptions struct {
	// Config is the wallet multi-sig configuration. Note: the first config of any wallet
	// before it is deployed is used to derive it's the account address of the wallet.
	Config core.WalletConfig

	// Context is the WalletContext of deployed wallet-contract modules for the Smart Wallet.
	// NOTE: if a WalletContext is not provided, then `SequenceContext()` value is used.
	Context *WalletContext

	// Skips config sorting and keeps signers order as-is
	SkipSortSigners bool

	// Address used for the wallet
	// if this value is defined, the address derived from the sequence config is ignored
	Address common.Address
}

func NewWallet(walletOptions WalletOptions, signers ...Signer) (*Wallet, error) {
	if walletConfigV1 := walletOptions.Config.(*v1.WalletConfig); walletConfigV1 != nil {
		seqContext := sequenceContext
		if walletOptions.Context != nil {
			seqContext = *walletOptions.Context
		}

		// Clone wallet config to avoid mutating the original
		walletConfigV1 = walletConfigV1.Clone()

		// Sort signers
		if !walletOptions.SkipSortSigners {
			err := SortWalletConfig(walletConfigV1)
			if err != nil {
				return nil, fmt.Errorf("sequence.NewWallet: %w", err)
			}
		}

		// Check if wallet config is usable
		if err := walletConfigV1.IsUsable(); err != nil {
			return nil, fmt.Errorf("sequence.NewWallet: %w", err)
		}

		// Generate address
		address := walletOptions.Address
		if address == (common.Address{}) {
			var err error
			address, err = AddressFromImageHash(walletConfigV1.ImageHash().Hex(), seqContext)
			if err != nil {
				return nil, fmt.Errorf("sequence.NewWallet: %w", err)
			}
		}

		// Check signers
		for _, signer := range signers {
			_, canSignMessage := signer.(MessageSigner)
			_, canSignDigest := signer.(DigestSigner)
			if !canSignMessage && !canSignDigest {
				return nil, fmt.Errorf("sequence.Wallet#UseSigners: signer is not a valid signer")
			}
		}

		w := &Wallet{
			config:          walletConfigV1,
			context:         seqContext,
			address:         address,
			skipSortSigners: walletOptions.SkipSortSigners,
		}
		w.signers = signers

		return w, nil
	} else if walletConfigV2 := walletOptions.Config.(*v2.WalletConfig); walletConfigV2 != nil {
		return nil, nil
	}

	return nil, fmt.Errorf("sequence.NewWallet: unsupported wallet config version")
}

func NewWalletSingleOwner(owner Signer, optContext ...WalletContext) (*Wallet, error) {
	seqContext := sequenceContext
	if len(optContext) > 0 {
		seqContext = optContext[0]
	}

	_, canSignMessage := owner.(MessageSigner)
	_, canSignDigest := owner.(DigestSigner)
	if !canSignMessage && !canSignDigest {
		return nil, fmt.Errorf("sequence.Wallet#UseSigners: signer is not a valid signer")
	}

	// new sequence v1 wallet
	return NewWallet(WalletOptions{
		Config: &v1.WalletConfig{
			Threshold_: 1, //big.NewInt(1),
			Signers_: v1.WalletConfigSigners{
				{Weight: 1, Address: owner.Address()},
			},
		},
		Context: &seqContext,
	}, owner)
}

type Wallet struct {
	context WalletContext
	config  core.WalletConfig
	signers []Signer

	provider *ethrpc.Provider
	relayer  Relayer
	address  common.Address

	skipSortSigners bool

	chainID *big.Int
}

var (
	ErrUnknownChainID = fmt.Errorf("chainID is unknown")
	ErrProviderNotSet = fmt.Errorf("provider is not set")
	ErrRelayerNotSet  = fmt.Errorf("relayer is not set")
)

func (w *Wallet) UseConfig(config core.WalletConfig) (*Wallet, error) {
	ww, err := NewWallet(WalletOptions{
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

func (w *Wallet) UseSigners(signers ...Signer) (*Wallet, error) {
	ww, err := NewWallet(WalletOptions{
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
	for _, signer := range signers {
		_, canSignMessage := signer.(MessageSigner)
		_, canSignDigest := signer.(DigestSigner)
		if !canSignMessage && !canSignDigest {
			return nil, fmt.Errorf("sequence.Wallet#UseSigners: signer is not a valid signer")
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

func (w *Wallet) GetWalletConfig() core.WalletConfig {
	return w.config
}

func (w *Wallet) Address() common.Address {
	return w.address
}

func (w *Wallet) ImageHash() (common.Hash, error) {
	return w.config.ImageHash().Hash, nil
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

func (w *Wallet) GetSigner(address common.Address) (Signer, bool) {
	for _, s := range w.signers {
		if s.Address() == address {
			return s, true
		}
	}
	return nil, false
}

func (w *Wallet) GetSignerWeight() *big.Int {
	signerAddresses := w.GetSignerAddresses()

	if walletConfigV1 := w.config.(*v1.WalletConfig); walletConfigV1 != nil {
		totalWeight := big.NewInt(0)
		for _, sa := range signerAddresses {
			for address, weight := range w.config.Signers() {
				if address == sa {
					totalWeight = big.NewInt(0).Add(totalWeight, big.NewInt(0).SetUint64(uint64(weight)))
				}
			}
		}
		return totalWeight
	} else if walletConfigV2 := w.config.(*v2.WalletConfig); walletConfigV2 != nil {
		// todo: implement weight for v2
	}

	return nil
}

func (w *Wallet) GetNonce(optBlockNum ...*big.Int) (*big.Int, error) {
	if w.relayer == nil {
		return nil, ErrRelayerNotSet
	}
	var blockNum *big.Int
	if len(optBlockNum) > 0 {
		blockNum = optBlockNum[0]
	}
	return w.relayer.GetNonce(context.Background(), w.config, w.context, nil, blockNum)
}

func (w *Wallet) GetTransactionCount(optBlockNum ...*big.Int) (*big.Int, error) {
	return w.GetNonce(optBlockNum...)
}

func (w *Wallet) SignMessage(ctx context.Context, msg []byte) ([]byte, *Signature, error) {
	sCtx := ContextWithAuxData(ctx, &AuxData{
		ChainID: w.chainID,
		Address: &w.address,
		Msg:     msg,
	})
	return w.SignDigest(sCtx, MessageDigest(msg))
}

// func (w *Wallet) SignTypedData() // TODO

func (w *Wallet) SignDigest(ctx context.Context, digest common.Hash, optChainID ...*big.Int) ([]byte, *Signature, error) {
	if (optChainID == nil && len(optChainID) == 0) && w.chainID == nil {
		return nil, nil, fmt.Errorf("sequence.Wallet#SignDigest: %w", ErrUnknownChainID)
	}

	var chainID *big.Int
	if optChainID != nil && len(optChainID) > 0 {
		chainID = optChainID[0]
	} else {
		chainID = w.chainID
	}

	subDigest, err := SubDigest(chainID, w.Address(), digest)
	if err != nil {
		return nil, nil, fmt.Errorf("SignDigest, subDigestOf: %w", err)
	}

	sign := func(ctx context.Context, signerAddress common.Address) (core.SignerSignatureType, []byte, error) {
		signer, _ := w.GetSigner(signerAddress)

		if signer == nil {
			// signer isn't available, just include the config value of address
			// without it's signature
			return 0, nil, fmt.Errorf("no signature for %v", signerAddress)
		}

		if eoaSigner, ok := signer.(MessageSigner); ok {
			sigValue, err := eoaSigner.SignMessage(subDigest)
			if err != nil {
				return 0, nil, fmt.Errorf("signer.SignMessage subDigest: %w", err)
			}

			return core.SignerSignatureTypeEthSign, sigValue, nil
		} else if seqSigner, ok := signer.(DigestSigner); ok {
			_, seqSign, err := seqSigner.SignDigest(ctx, common.BytesToHash(subDigest), chainID)
			if err != nil {
				return 0, nil, fmt.Errorf("signer.SignMessage subDigest: %w", err)
			}

			return core.SignerSignatureTypeEIP1271, seqSign.Signers[0].Value, nil
		} else if _, ok := signer.(SignerGuardSigner); ok {
			// todo: rework how guard signer works (take inspiration from sequence-js orchestrator)
		}
		return 0, nil, fmt.Errorf("signer is not a valid signer type")
	}

	if config, ok := w.config.(*v2.WalletConfig); ok {
		sig, err := config.BuildRegularSignature(ctx, sign)
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, BuildRegularSignature: %w", err)
		}

		sigEnc, err := sig.Data()
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, sig.Data: %w", err)
		}
		return sigEnc, nil, nil
	} else if config, ok := w.config.(*v1.WalletConfig); ok {
		sig, err := config.BuildSignature(ctx, sign)
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, BuildSignature: %w", err)
		}

		sigEnc, err := sig.Data()
		if err != nil {
			return nil, nil, fmt.Errorf("SignDigest, sig.Data: %w", err)
		}
		return sigEnc, nil, nil
	} else {
		return nil, nil, fmt.Errorf("SignDigest, unknown config type")
	}
}

func (w *Wallet) SignTransaction(ctx context.Context, txn *Transaction) (*SignedTransactions, error) {
	return w.SignTransactions(ctx, Transactions{txn})
}

func (w *Wallet) SignTransactions(ctx context.Context, txns Transactions) (*SignedTransactions, error) {
	if len(txns) == 0 {
		return nil, fmt.Errorf("cannot sign an empty set of transactions")
	}

	var err error

	// If a transaction has 0 gasLimit and not revertOnError
	// compute all new gas limits
	estimateGas := false
	for _, txn := range txns {
		if !txn.RevertOnError && (txn.GasLimit == nil || txn.GasLimit.Cmp(big.NewInt(0)) == 0) {
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

	// load nonce from transactions
	nonce, err := txns.Nonce()
	if err != nil {
		return nil, fmt.Errorf("cannot load nonce from transactions: %w", err)
	}

	// if nonce is undefined
	// load latest nonce from wallet
	if nonce == nil {
		nonce, err = w.GetNonce()
		if err != nil {
			return nil, err
		}
	}

	bundle := Transaction{
		Transactions: txns,
		Nonce:        nonce,
	}

	// Get transactions digest
	digest, err := bundle.Digest()
	if err != nil {
		return nil, err
	}

	// Sign AuxData (for signing services)
	sCtx := ctx
	if auxData := w.auxDataFromTransactionBundle(&bundle); auxData != nil {
		sCtx = ContextWithAuxData(ctx, auxData)
	}

	// Sign the transactions
	sig, _, err := w.SignDigest(sCtx, digest)
	if err != nil {
		return nil, err
	}

	return &SignedTransactions{
		ChainID:       w.chainID,
		WalletConfig:  w.config,
		WalletContext: w.context,
		Transactions:  txns,
		Nonce:         nonce,
		Digest:        digest,
		Signature:     sig,
	}, nil
}

func (w *Wallet) SendTransaction(ctx context.Context, signedTxns *SignedTransactions) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	return w.SendTransactions(ctx, signedTxns)
}

func (w *Wallet) SendTransactions(ctx context.Context, signedTxns *SignedTransactions) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
	if w.relayer == nil {
		return "", nil, nil, ErrRelayerNotSet
	}
	return w.relayer.Relay(ctx, signedTxns)
}

func (w *Wallet) IsDeployed() (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}
	return IsWalletDeployed(w.provider, w.Address())
}

func (w *Wallet) Deploy(ctx context.Context) (MetaTxnID, *types.Transaction, ethtxn.WaitReceipt, error) {
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

	signerTxn, err := w.SignTransaction(ctx, txn)
	if err != nil {
		return "", nil, nil, err
	}

	return w.relayer.Relay(ctx, signerTxn)
}

// func (w *Wallet) UpdateConfig() // TODO in future

// func (w *Wallet) PublishConfig() // TODO in future

func (w *Wallet) IsValidSignature(digest common.Hash, signature []byte) (bool, error) {
	if w.provider == nil {
		return false, ErrProviderNotSet
	}

	// Assume that this context if for a v1 wallet, so we bundle the context in a mapping
	// with a single key of (1)
	contexts := make(map[uint16]WalletContext)
	contexts[1] = w.context

	return IsValidSignature(w.Address(), digest, signature, contexts, w.chainID, w.provider)
}

func (w *Wallet) auxDataFromTransactionBundle(bundle *Transaction) *AuxData {
	bundle.Signature = []byte{}

	msg, err := Transactions{bundle}.EncodeRaw()
	if err != nil {
		return nil
	}

	return &AuxData{
		Msg:     msg,
		Sig:     nil,
		ChainID: w.chainID,
		Address: &w.address,
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

var (
	// ImageHashUpdatedEventSig is emitted anytime wallet config is updated.
	ImageHashUpdatedEventSig = MustEncodeSig("ImageHashUpdated(bytes32)")

	// ImplementationUpdatedEventSig is emitted anytime a wallet's mainModule is changed,
	// this is a rare occurence.
	ImplementationUpdatedEventSig = MustEncodeSig("ImplementationUpdated(address)")
)

// https://github.com/0xsequence/wallet-contracts/blob/master/src/contracts/Wallet.sol#L57-L59
const WalletContractBytecode = "0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3"

var walletContractBytecode = common.FromHex(WalletContractBytecode)
