package sequence

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence/contracts"
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

	Nonce *big.Int // Nonce of the transaction, will be set automatically if unset
	// Expiration *big.Int // optional.. TODO
	// AfterNonce .. // optional.. TODO

	// Children represents nested/bundled transactions
	Children TransactionBundle
}

func (t *Transaction) IsBundle() bool {
	return t.Children.Transactions != nil && len(t.Children.Transactions) > 0
}

// AddToBundle will create a bundle from the passed txns and add it to current transaction
func (t *Transaction) AddToBundle(txns Transactions) {
	if t.IsBundle() {
		// append to existing bundle
		t.Children.Transactions.Append(txns)
	} else {
		// create a new bundle
		t.Children = TransactionBundle{Transactions: txns}
	}
}

func (t *Transaction) Encode() error {
	// will encode the bundle if necessary..
	if t.IsBundle() {
		err := t.Children.Encode()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t Transaction) Bundle() Transactions {
	return []*Transaction{&t}
}

type Transactions []*Transaction

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

func (t Transactions) Nonce() (*big.Int, error) {
	var nonce *big.Int
	for _, txn := range t {
		if txn.Nonce != nil {
			nonce = txn.Nonce
			break
		}
	}
	// For a given bundle of transactions (of the same depth), all nonce must be the same.
	for _, txn := range t {
		if txn.Nonce != nil && txn.Nonce.Cmp(nonce) != 0 {
			return nil, fmt.Errorf("mixed nonces on sequence transactions are not allowed for a given bundle")
		}
	}
	return nonce, nil
}

func (t Transactions) Digest() ([]byte, error) {
	metaNonce, err := t.Nonce()
	if err != nil {
		return nil, err
	}

	data, err := abiTransactionsDigestType.PackValues([]interface{}{metaNonce, t.AsValues()})
	if err != nil {
		return nil, fmt.Errorf("transaction digest failed to pack values: %w", err)
	}

	return ethcoder.Keccak256(data), nil
}

func (t Transactions) Encode() error {
	return nil
}

func (t Transactions) AsValues() []Transaction {
	v := []Transaction{}
	for _, o := range t {
		v = append(v, *o)
	}
	return v
}

type TransactionBundle struct {
	// Transactions in a bundle
	Transactions

	// ExecData represents the encoding of Transactions above, in form of a selfExecute call
	ExecData []byte
}

func (b *TransactionBundle) Encode() error {
	return nil
}

// SignedTransactions includes a signed meta-transaction payload intended for the relayer.
type SignedTransactions struct {
	ChainID       *big.Int
	WalletConfig  WalletConfig
	WalletContext WalletContext

	Digest       []byte       // Digest of the transactions
	Signature    []byte       // Signature (encoded as bytes from *Signature) of the txn digest
	Transactions Transactions // The meta-transactions
}

// Transaction events as defined in wallet-contracts IModuleCalls.sol
var (
	// NonceChangeEventSig is the signature event emitted as the first event on the batch execution
	// 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881
	NonceChangeEventSig = MustEncodeSig("NonceChange(uint256,uint256)")

	// TxFailedEventSig is the signature event emitted in a failed smart-wallet meta-transaction batch
	// 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7
	TxFailedEventSig = MustEncodeSig("TxFailed(bytes32,bytes)")

	// TxExecutedEventSig is the signature of the event emitted in a successful transaction
	// ........
	TxExecutedEventSig = MustEncodeSig("TxExecuted(bytes32)")
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

func GetWalletNonce(provider *ethrpc.Provider, walletConfig WalletConfig, walletContext WalletContext, space *big.Int, blockNum *big.Int) (*big.Int, error) {
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

	contract := ethcontract.NewContractCaller(walletAddress, contracts.WalletMainModule.ABI, provider)

	var nonceResult *big.Int
	err = contract.Call(nil, &nonceResult, "readNonce", space)
	if err != nil {
		return nil, err
	}

	return nonceResult, nil
}

// IsTxFailedLog returns true if the log belongs to a failed smart meta-transaction
func IsTxFailedLog(logs []*types.Log) (string, bool, error) {
	log := logs[len(logs)-1] // check the last log in the set
	if ok := len(log.Topics) == 1 && log.Topics[0] == TxFailedEventSig; ok {

		var txhash [32]byte
		var reason []byte

		if err := ethcoder.AbiDecoder([]string{"bytes32", "bytes"}, log.Data, []interface{}{&txhash, &reason}); err != nil {
			return "", false, err
		}

		var inputType = abi.Arguments{
			abi.Argument{
				Type: ethcoder.MustNewType("string"),
			},
		}

		if len(reason) < 4 {
			// There is no reason
			return "", true, nil
		}
		errMsg := ""
		if err := inputType.Unpack(&errMsg, reason[4:]); err != nil {
			return "", false, err
		}
		return errMsg, true, nil
	}
	return "", false, nil
}

func IsNonceChangedEvent(logs []*types.Log) bool {
	if len(logs[0].Topics) != 1 {
		return false
	}
	if logs[0].Topics[0] != NonceChangeEventSig {
		return false
	}
	return true
}

// prepareTransactionsForSigning checks the transactions data structure with basic
// integrity checks
func prepareTransactionsForSigning(txns Transactions) (Transactions, error) {
	if len(txns) == 0 {
		return nil, fmt.Errorf("cannot sign an empty set of transactions")
	}

	stxns := Transactions{}
	for _, txn := range txns {
		if txn == nil {
			return nil, fmt.Errorf("cannot sign a nil transaction")
		}
		if txn.Value == nil {
			txn.Value = big.NewInt(0) // default of 0 is expected by abi coder
		}

		stxns = append(stxns, txn)
	}

	return stxns, nil
}

func DecodeTransactions(execdata []byte) (Transactions, error) {
	if len(execdata) < 4 {
		return nil, fmt.Errorf("not enough data to be an execute or selfExecute call")
	}

	var transactions []Transaction
	var nonce *big.Int
	var signature []byte

	executeMethod := contracts.WalletMainModule.ABI.Methods["execute"]
	selfExecuteMethod := contracts.WalletMainModule.ABI.Methods["selfExecute"]

	if bytes.Equal(execdata[:4], executeMethod.ID) {
		err := executeMethod.Inputs.Unpack(&[]interface{}{&transactions, &nonce, &signature}, execdata[4:])
		if err != nil {
			return nil, err
		}
	} else if bytes.Equal(execdata[:4], selfExecuteMethod.ID) {
		err := selfExecuteMethod.Inputs.Unpack(&transactions, execdata[4:])
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("not an execute or selfExecute call")
	}

	var pointers Transactions
	for i := 0; i < len(transactions); i++ {
		pointers = append(pointers, &transactions[i])
	}
	return pointers, nil
}
