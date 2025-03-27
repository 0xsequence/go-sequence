package sequence

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
	v1 "github.com/0xsequence/go-sequence/core/v1"
	v3 "github.com/0xsequence/go-sequence/core/v3"
)

// Transaction type for Sequence meta-transaction, with encoded calldata.
//
// The fields with abi struct tags match the `Transaction` type as defined in the IModuleCalls interface.
//
// https://github.com/0xsequence/wallet-contracts/blob/master/src/contracts/modules/commons/interfaces/IModuleCalls.sol#L13
type Transaction struct {
	DelegateCall  bool           `abi:"delegateCall"`  // Performs delegatecall
	RevertOnError bool           `abi:"revertOnError"` // Reverts transaction bundle if tx fails
	GasLimit      *big.Int       `abi:"gasLimit"`      // Maximum gas to be forwarded
	To            common.Address `abi:"target"`        // Address to send transaction, aka target
	Value         *big.Int       `abi:"value"`         // Amount of ETH to pass with the call
	Data          []byte         `abi:"data"`          // Calldata to pass

	Transactions Transactions // Child transactions
	Nonce        *big.Int     // Meta-Transaction nonce, with encoded space
	Signature    []byte       // Meta-Transaction signature

	// Expiration *big.Int // optional.. TODO
	// AfterNonce .. // optional.. TODO

	encoded bool
}

func (t *Transaction) IsEncoded() bool {
	return t.encoded
}

func (t *Transaction) Clone() *Transaction {
	clone := Transaction{
		DelegateCall:  t.DelegateCall,
		RevertOnError: t.RevertOnError,
		To:            t.To,
	}
	if t.GasLimit != nil {
		clone.GasLimit = new(big.Int).Set(t.GasLimit)
	}
	if t.Value != nil {
		clone.Value = new(big.Int).Set(t.Value)
	}
	if t.Data != nil {
		clone.Data = make([]byte, len(t.Data))
		copy(clone.Data, t.Data)
	}
	if t.Transactions != nil {
		clone.Transactions = t.Transactions.Clone()
	}
	if t.Nonce != nil {
		clone.Nonce = new(big.Int).Set(t.Nonce)
	}
	if t.Signature != nil {
		clone.Signature = make([]byte, len(t.Signature))
		copy(clone.Signature, t.Signature)
	}
	return &clone
}

func (t *Transaction) Bundle() Transactions {
	return Transactions{t}
}

func (t Transactions) Nonce() (*big.Int, error) {
	var cand *big.Int

	for _, transaction := range t {
		if transaction.Nonce != nil {
			if cand == nil {
				cand = new(big.Int).Set(transaction.Nonce)
			} else {
				if cand.Cmp(transaction.Nonce) != 0 {
					return nil, fmt.Errorf("multiple nonces found")
				}
			}
		}
	}

	return cand, nil
}

func (t *Transaction) IsValid() error {
	// invariant 1: calldata and execdata are mutually exclusive
	if t.Data != nil && (t.Transactions != nil || t.Nonce != nil || t.Signature != nil) {
		return fmt.Errorf("transactions cannot have both calldata and execdata")
	}

	// invariant 2: if either nonce or signature are set, the other must also be set
	if (t.Nonce != nil) != (t.Signature != nil) {
		return fmt.Errorf("transactions must have both nonce and signature or neither")
	}

	return nil
}

func (t *Transaction) IsBundle() bool {
	return len(t.Transactions) != 0
}

func (t *Transaction) ReduceSignatures(chainID *big.Int) error {
	if err := t.IsValid(); err != nil {
		return err
	}

	if len(t.Signature) != 0 {
		signature, err := GenericDecodeSignature[*v1.WalletConfig](t.Signature)
		if err != nil {
			return err
		}

		_, subdigest, err := ComputeMetaTxnID(chainID, t.To, t.Transactions, t.Nonce, MetaTxnWalletExec)
		if err != nil {
			return err
		}
		signature = signature.Reduce(core.Subdigest{Hash: subdigest})

		encoded, err := signature.Data()
		if err != nil {
			return err
		}

		t.Signature = encoded
	}

	for _, transaction := range t.Transactions {
		err := transaction.ReduceSignatures(chainID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Transaction) Execdata() ([]byte, error) {
	if err := t.IsValid(); err != nil {
		return nil, err
	}

	if !t.IsBundle() {
		return nil, fmt.Errorf("transaction is not a bundle: only bundles have execdata")
	}

	encodedTxns, err := t.Transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}

	if t.Signature != nil {
		return contracts.V1.WalletMainModule.Encode("execute", encodedTxns, t.Nonce, t.Signature)
	} else {
		return contracts.V1.WalletMainModule.Encode("selfExecute", encodedTxns)
	}
}

func (t *Transaction) Digest() (common.Hash, error) {
	// precondition: the digest only exists for transaction bundles
	if t.Data != nil {
		return common.Hash{}, fmt.Errorf("transaction bundles cannot not have calldata")
	}

	if t.Nonce != nil {
		return ComputeWalletExecDigest(t.Nonce, t.Transactions)
	} else {
		return ComputeSelfExecDigest(t.Transactions)
	}
}

func (t *Transaction) GuestDigest() (common.Hash, error) {
	// precondition: the digest only exists for transaction bundles
	if t.Data != nil {
		return common.Hash{}, fmt.Errorf("transaction bundles cannot not have calldata")
	}

	// TODO: should we check t.Nonce and return error, as invarient check..?

	return ComputeGuestExecDigest(t.Transactions)
}

// AddToBundle will create a bundle from the passed txns and add it to current transaction
func (t *Transaction) AddToBundle(txns Transactions) {
	if t.IsBundle() {
		// append to existing bundle
		t.Transactions.Append(txns)
	} else {
		// create a new bundle
		t.Transactions = txns
	}
}

type Transactions []*Transaction

func NewTransactionsFromValues(values []Transaction) Transactions {
	transactions := Transactions{}
	for i := 0; i < len(values); i++ {
		transactions = append(transactions, &values[i])
	}
	return transactions
}

func (t Transactions) Payload(to common.Address, chainID *big.Int, space, nonce *big.Int) (v3.CallsPayload, error) {
	calls := make([]v3.Call, len(t))
	for i, tx := range t {
		// Convert from Transaction.RevertOnError to v3.BehaviorOnError
		var behaviorOnError v3.BehaviorOnError
		if tx.RevertOnError {
			behaviorOnError = v3.BehaviorOnErrorRevert
		} else {
			behaviorOnError = v3.BehaviorOnErrorIgnore
		}

		// Handle nested transactions by recursively encoding them
		data := tx.Data
		if tx.IsBundle() {
			subPayload, err := tx.Transactions.Payload(tx.To, chainID, big.NewInt(0), tx.Nonce)
			if err != nil {
				return v3.CallsPayload{}, fmt.Errorf("unable to convert nested transactions to v3 payload: %w", err)
			}

			data, err = contracts.V3.WalletStage1Module.Encode("selfExecute", subPayload.Encode(tx.To))
			if err != nil {
				return v3.CallsPayload{}, fmt.Errorf("unable to encode nested v3 payload selfExecute call: %w", err)
			}
		}

		calls[i] = v3.Call{
			To:              tx.To,
			Value:           tx.Value,
			Data:            data,
			GasLimit:        tx.GasLimit,
			DelegateCall:    tx.DelegateCall,
			BehaviorOnError: behaviorOnError,
		}
	}

	return v3.NewCallsPayload(to, chainID, calls, space, nonce), nil
}

func (t Transactions) EncodeRaw() ([]byte, error) {
	encoded, err := t.EncodedTransactions()
	if err != nil {
		return nil, err
	}
	encoder := abi.Arguments{abi.Argument{Type: abiTransactionsType}}
	return encoder.Pack(encoded)
}

func DecodeExecdata(data []byte, walletAddress common.Address, chainID *big.Int) (Transactions, *big.Int, []byte, error) {
	if len(data) < 4 {
		return nil, nil, nil, fmt.Errorf("not an execute or selfExecute call")
	}

	var transactions []Transaction
	var nonce *big.Int
	var signature []byte
	var err error

	// Check for V1 methods
	executeMethod := contracts.V1.WalletMainModule.ABI.Methods["execute"]
	selfExecuteMethod := contracts.V1.WalletMainModule.ABI.Methods["selfExecute"]

	// Check for V3 methods
	v3ExecuteMethod := contracts.V3.WalletStage1Module.ABI.Methods["execute"]
	v3SelfExecuteMethod := contracts.V3.WalletStage1Module.ABI.Methods["selfExecute"]

	if bytes.Equal(data[:4], executeMethod.ID) {
		// V1 execute method
		var values []interface{}
		values, err = executeMethod.Inputs.Unpack(data[4:])
		if err == nil {
			err = executeMethod.Inputs.Copy(&[]interface{}{&transactions, &nonce, &signature}, values)
		}
	} else if bytes.Equal(data[:4], selfExecuteMethod.ID) {
		// V1 selfExecute method
		var values []interface{}
		values, err = selfExecuteMethod.Inputs.Unpack(data[4:])
		if err == nil {
			err = selfExecuteMethod.Inputs.Copy(&transactions, values)
		}
	} else if bytes.Equal(data[:4], v3ExecuteMethod.ID) {
		// V3 execute method takes (bytes calldata _payload, bytes calldata _signature)
		var values []interface{}
		values, err = v3ExecuteMethod.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, nil, nil, err
		}

		var payload, sig []byte
		err = v3ExecuteMethod.Inputs.Copy(&[]interface{}{&payload, &sig}, values)
		if err != nil {
			return nil, nil, nil, err
		}

		// Decode the V3 payload
		decoded, err := v3.DecodeCalls(walletAddress, chainID, payload)
		if err != nil {
			return nil, nil, nil, err
		}

		// Convert V3 payload to Transactions
		transactions = make([]Transaction, len(decoded.Calls))
		for i, call := range decoded.Calls {
			transactions[i] = Transaction{
				To:            call.To,
				Value:         call.Value,
				Data:          call.Data,
				GasLimit:      call.GasLimit,
				DelegateCall:  call.DelegateCall,
				RevertOnError: call.BehaviorOnError == v3.BehaviorOnErrorRevert,
			}
		}
		nonce = decoded.Nonce
	} else if bytes.Equal(data[:4], v3SelfExecuteMethod.ID) {
		// V3 selfExecute method takes (bytes calldata _payload)
		var values []interface{}
		values, err = v3SelfExecuteMethod.Inputs.Unpack(data[4:])
		if err != nil {
			return nil, nil, nil, err
		}

		var payload []byte
		err = v3SelfExecuteMethod.Inputs.Copy(&[]interface{}{&payload}, values)
		if err != nil {
			return nil, nil, nil, err
		}

		// Decode the V3 payload
		decoded, err := v3.DecodeCalls(walletAddress, chainID, payload)
		if err != nil {
			return nil, nil, nil, err
		}

		transactions = make([]Transaction, len(decoded.Calls))
		for i, call := range decoded.Calls {
			transactions[i] = Transaction{
				To:            call.To,
				Value:         call.Value,
				Data:          call.Data,
				GasLimit:      call.GasLimit,
				DelegateCall:  call.DelegateCall,
				RevertOnError: call.BehaviorOnError == v3.BehaviorOnErrorRevert,
			}
		}
		nonce = decoded.Nonce
	} else {
		// V3 guest module allows sending just the packed calls payload without solidity abi encoding
		decoded, err := v3.DecodeCalls(walletAddress, chainID, data)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("not an execute or selfExecute call")
		}

		transactions = make([]Transaction, len(decoded.Calls))
		for i, call := range decoded.Calls {
			transactions[i] = Transaction{
				To:            call.To,
				Value:         call.Value,
				Data:          call.Data,
				GasLimit:      call.GasLimit,
				DelegateCall:  call.DelegateCall,
				RevertOnError: call.BehaviorOnError == v3.BehaviorOnErrorRevert,
			}
		}
		nonce = decoded.Nonce
	}

	if err != nil {
		return nil, nil, nil, err
	}

	// Recursively decode nested transactions
	for i := 0; i < len(transactions); i++ {
		if len(transactions[i].Data) > 0 {
			decodedTransactions, decodedNonce, decodedSignature, err := DecodeExecdata(transactions[i].Data, walletAddress, chainID)
			if err == nil {
				transactions[i].Data = nil
				transactions[i].Transactions = decodedTransactions
				transactions[i].Nonce = decodedNonce
				transactions[i].Signature = decodedSignature
			}
		}
	}

	return NewTransactionsFromValues(transactions), nonce, signature, nil
}

// Append will append the passed txns to the `t` array (as separate Transaction elements).
func (t *Transactions) Append(txns Transactions) {
	*t = append(*t, txns...)
}

// Prepend will prepend the passed txns to the `t` array (as separate Transaction elements).
func (t *Transactions) Prepend(txns Transactions) {
	*t = append(txns, *t...)
}

// AppendBundle will append the passed txns as *a bundle of txns*. This means, it will be included as a single
// element at this level of the `t` array.
func (t Transactions) AppendBundle(txns Transactions) {
	bundleTxn := &Transaction{}
	bundleTxn.AddToBundle(txns)
	t.Append(Transactions{bundleTxn})
}

// PrependBundle will prepend the passed txns as *a bundle of txns*. This means, it will be included as a single
// element at this level of the `t` array.
func (t Transactions) PrependBundle(txns Transactions) {
	bundleTxn := &Transaction{}
	bundleTxn.AddToBundle(txns)
	t.Prepend(Transactions{bundleTxn})
}

func (t Transactions) EncodedTransactions() ([]Transaction, error) {
	stxns := []Transaction{}
	for _, txn := range t {
		if txn == nil {
			return nil, fmt.Errorf("cannot sign a nil transaction")
		}

		txn := txn.Clone()

		if !txn.encoded {
			// encode the transaction is still unencoded
			if txn.Value == nil {
				txn.Value = big.NewInt(0) // default of 0 is expected by abi coder
			}
			if txn.GasLimit == nil {
				txn.GasLimit = big.NewInt(0) // default of 0 is expected by abi coder
			}

			// flatten the bundle into a single transaction as expected by `execute` contract method
			if txn.IsBundle() {
				var err error
				txn.Data, err = txn.Execdata()
				if err != nil {
					return nil, err
				}
			}
		}

		txn.encoded = true
		stxns = append(stxns, *txn)
	}

	return stxns, nil
}

func (t Transactions) AsValues() []Transaction {
	v := []Transaction{}
	for _, o := range t {
		v = append(v, *o)
	}
	return v
}

func (t Transactions) Clone() Transactions {
	txns := make(Transactions, len(t))
	for i, txn := range t {
		txns[i] = txn.Clone()
	}
	return txns
}

// SignedTransactions includes a signed meta-transaction payload intended for the relayer.
type SignedTransactions struct {
	ChainID       *big.Int
	WalletAddress common.Address
	WalletConfig  core.WalletConfig
	WalletContext WalletContext

	Transactions Transactions // The meta-transactions
	Space        *big.Int     // Nonce space of the transactions
	Nonce        *big.Int     // Nonce of the transactions
	Digest       common.Hash  // Digest of the transactions
	Signature    []byte       // Signature (encoded as bytes from *Signature) of the txn digest
}

func (t *SignedTransactions) Payload() (v3.CallsPayload, error) {
	return t.Transactions.Payload(t.WalletAddress, t.ChainID, t.Space, t.Nonce)
}

func (t *SignedTransactions) ExecuteV3() ([]byte, error) {
	payload, err := t.Payload()
	if err != nil {
		return nil, err
	}
	return contracts.V3.WalletStage1Module.Encode("execute", payload.Encode(t.WalletAddress), t.Signature)
}

func (t *SignedTransactions) Execdata() ([]byte, error) {
	encodedTxns, err := t.Transactions.EncodedTransactions()
	if err != nil {
		return nil, err
	}
	return contracts.V1.WalletMainModule.Encode("execute", encodedTxns, t.Nonce, t.Signature)
}

// Transaction events as defined in wallet-contracts IModuleCalls.sol
var (
	// NonceChangeEventSig is the signature event emitted as the first event on the batch execution
	// 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881
	NonceChangeEventSig = MustEncodeSig("NonceChange(uint256,uint256)")

	// TxFailedEventSig is the signature event emitted in a failed smart-wallet meta-transaction batch
	// 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7
	V1TxFailedEventSig = MustEncodeSig("TxFailed(bytes32,bytes)")

	// TxExecutedEventSig is the signature event emitted in a successful smart-wallet meta-transaction batch (for v2)
	// 0x5c4eeb02dabf8976016ab414d617f9a162936dcace3cdef8c69ef6e262ad5ae7
	// TxExecuted(bytes32 indexed _tx, uint256 _index)
	V2TxExecutedEventSig = MustEncodeSig("TxExecuted(bytes32,uint256)")

	// TxFailedEventSig is the signature event emitted in a failed smart-wallet meta-transaction batch (for v2)
	// 0xab46c69f7f32e1bf09b0725853da82a211e5402a0600296ab499a2fb5ea3b419
	// TxFailed(bytes32 indexed _tx, uint256 _index, bytes _reason)
	V2TxFailedEventSig = MustEncodeSig("TxFailed(bytes32,uint256,bytes)")

	// 0x5a589b1d8062f33451d29cae3dabd9b2e36c62aee644178c600977ca8dda661a
	V3CallSucceeded = MustEncodeSig("CallSucceeded(bytes32,uint256)")
	// 0x115f347c00e69f252cd6b63c4f81022a9564c6befe8aa719cb74640a4a306f0d
	V3CallFailed = MustEncodeSig("CallFailed(bytes32,uint256,bytes)")
	// 0xc2c704302430fe0dc8d95f272e2f4e54bbbc51a3327fd5d75ab41f9fc8fd129b
	V3CallAborted = MustEncodeSig("CallAborted(bytes32,uint256,bytes)")
	// 0x9ae934bf8a986157c889a24c3b3fa85e74b7e4ee4b1f8fc6e7362cb4c1d19d8b
	V3CallSkipped = MustEncodeSig("CallSkipped(bytes32,uint256)")
)

// EncodeNonce with space
func EncodeNonce(space *big.Int, nonce *big.Int) (*big.Int, error) {
	if space.Cmp(new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil)) >= 0 {
		return nil, fmt.Errorf("nonce space exceeds maximum of 2^160-1")
	}
	if nonce.Cmp(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)) >= 0 {
		return nil, fmt.Errorf("nonce exceeds maximum of 2^96-1")
	}

	shl := big.NewInt(2)
	shl.Exp(shl, big.NewInt(96), nil)

	res := new(big.Int).Mul(space, shl)
	res.Add(res, nonce)

	return res, nil
}

// DecodeNonce raw nonce, returns (space, nonce)
func DecodeNonce(raw *big.Int) (*big.Int, *big.Int) {
	shr := big.NewInt(2)
	shr.Exp(shr, big.NewInt(96), nil)

	nonce := new(big.Int)
	space := new(big.Int)

	nonce.Mod(raw, shr)
	space.Div(raw, shr)

	return space, nonce
}

// GenerateRandomNonce returns a random space for a nonce to ensure
// transactions can be executed in parallel.
func GenerateRandomNonce() (*big.Int, error) {
	// Max random value i.e 2^100 - 1, since nonce space is the first
	// 160 bits in a nonce. We skip the first 60 bits to keep
	// some of the nonce space clean.
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(100), nil).Sub(max, big.NewInt(1))

	// Generate random number between 0 and 2**160-1
	nonce, err := rand.Int(rand.Reader, max)
	if err != nil {
		return new(big.Int).Set(nil), err
	}

	// Right pad the nonce by 96 bits so it's a full uint256
	nonce = new(big.Int).Lsh(nonce, 96)
	return nonce, nil
}

func GetWalletNonce(provider *ethrpc.Provider, walletConfig core.WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
	walletAddress, err := AddressFromWalletConfig(walletConfig, walletContext)
	if err != nil {
		return nil, err
	}

	ok, err := IsWalletDeployed(provider, walletAddress)
	if err != nil {
		return nil, err
	}
	if !ok {
		return big.NewInt(0), nil
	}

	if space == nil {
		space = big.NewInt(0)
	}

	contract := ethcontract.NewContractCaller(walletAddress, contracts.V1.WalletMainModule.ABI, provider)

	var nonceResult *big.Int
	results := []interface{}{&nonceResult}
	err = contract.Call(nil, &results, "readNonce", space)
	if err != nil {
		return nil, err
	}

	return nonceResult, nil
}

func ReduceExecdataSignatures(chainID *big.Int, data []byte) ([]byte, error) {
	transactions, nonce, signature, err := DecodeExecdata(data, common.Address{}, chainID)
	if err != nil {
		return nil, fmt.Errorf("unable to decode execdata: %w", err)
	}

	transaction := Transaction{
		Transactions: transactions,
		Nonce:        nonce,
		Signature:    signature,
	}

	err = transaction.ReduceSignatures(chainID)
	if err != nil {
		return nil, fmt.Errorf("unable to reduce signatures: %w", err)
	}

	return transaction.Execdata()
}

func DecodeRawTransactions(data []byte) (Transactions, error) {
	decoder := abi.Arguments{abi.Argument{Type: abiTransactionsType}}
	values, err := decoder.Unpack(data)
	if err != nil {
		return nil, err
	}
	var transactions []Transaction
	if err = decoder.Copy(&transactions, values); err != nil {
		return nil, err
	}
	for i := range transactions {
		children, nonce, signature, err := DecodeExecdata(transactions[i].Data, common.Address{}, nil)
		if err == nil {
			transactions[i].Data = nil
			transactions[i].Transactions = children
			transactions[i].Nonce = nonce
			transactions[i].Signature = signature
		}
	}
	return NewTransactionsFromValues(transactions), nil
}
