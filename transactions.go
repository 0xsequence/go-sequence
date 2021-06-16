package sequence

import (
	"crypto/rand"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
)

// Transaction is a meta-transaction request..
// TODO: we could do embedded thing..?
type Transaction struct {
	DelegateCall  bool           // Performs delegatecall
	RevertOnError bool           // Reverts transaction bundle if tx fails
	GasLimit      *big.Int       // Maximum gas to be forwarded
	To            common.Address // Address to send transaction, aka target
	Value         *big.Int       // Amount of ETH to pass with the call
	Data          []byte         // Calldata to pass

	Nonce *big.Int // Nonce of the transaction, will be set automatically if unset
	// Expiration *big.Int // optional, .. // TODO
	// AfterNonce ... // TODO

	Nested Transactions // Nested transaction
}

func (t *Transaction) Encode() error {
	return nil
}

// for this structure..
type Transactions []*Transaction

type SignedTransactions struct {
	ChainID *big.Int
	Config  WalletConfig
	Context WalletContext

	Digest       []byte       // Digest of the transactions
	Signature    []byte       // or *Signature ..? // Signature of the digest
	Transactions Transactions // The meta-transactions
}

// TransactionEncoded type for Sequence meta-transaction, with encoded calldata. This structure
// is a fully encoded transaction payload which contains all sub-transactions, however, the Sequence
// wallet execute function accepts an array of TransactionEncoded as well, but note, that a MetaTransaction
// itself may encode nested sub-transactions, aka trees of calls.
//
// The type matches the `Transaction` type as defined in the IModuleCalls interface.
//
// https://github.com/0xsequence/wallet-contracts/blob/master/src/contracts/modules/commons/interfaces/IModuleCalls.sol#L13
// type TransactionEncoded struct {
// 	DelegateCall  bool           // Performs delegatecall
// 	RevertOnError bool           // Reverts transaction bundle if tx fails
// 	GasLimit      *big.Int       // Maximum gas to be forwarded
// 	Target        common.Address // Address of the contract to call
// 	Value         *big.Int       // Amount of ETH to pass with the call
// 	Data          []byte         // Calldata to pass
// }

func NewTransaction() (*Transaction, error) {
	return nil, nil
}

func NewTransactionBatch() (Transactions, error) {
	return nil, nil
}

// SignTransactions(wallet *Wallet, ..etc..)

func SendTransaction(signedTxs SignedTransactions) (string, error) {
	// returns the metaTxID ......
	return "", nil
}

//-------

// Transaction type for Sequence meta-transaction, with encoded calldata.
//
// The type matches the `Transaction` type as defined in the IModuleCalls interface.
//
// https://github.com/0xsequence/wallet-contracts/blob/master/src/contracts/modules/commons/interfaces/IModuleCalls.sol#L13
// type Transaction struct {
// 	DelegateCall  bool           // Performs delegatecall
// 	RevertOnError bool           // Reverts transaction bundle if tx fails
// 	GasLimit      *big.Int       // Maximum gas to be forwarded
// 	Target        common.Address // Address of the contract to call
// 	Value         *big.Int       // Amount of ETH to pass with the call
// 	Data          []byte         // Calldata to pass
// }

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

// type TransactionBatch struct {
// }

// TransactionRequest etc..

// TODO: see sequence.js..

// TODO: review ethkit / TransactionRequest
// and lets update from there.. yep, easier..

// type TransactionRequest struct {
// 	From  *common.Address // NOTE: this is automatically set..
// 	To    common.Address
// 	Nonce *big.Int

// 	GasLimit *big.Int
// 	// GasPrice *big.Int // NOTE: this is automatically set/estimated

// 	Data    []byte
// 	Value   *big.Int
// 	ChainID *big.Int

// 	// expiration?: BigNumberish
// 	// afterNonce?: NonceDependency | BigNumberish
// }

// TODO:

// EncodeNonce with space
func EncodeNonce(space *big.Int, nonce *big.Int) *big.Int {
	shl := big.NewInt(2)
	shl.Exp(shl, big.NewInt(96), nil)

	res := new(big.Int).Mul(space, shl)
	res.Add(res, nonce)

	return res
}

// DecodeNonce with space
func DecodeNonce(raw *big.Int) (*big.Int, *big.Int) {
	shr := big.NewInt(2)
	shr.Exp(shr, big.NewInt(96), nil)

	nonce := new(big.Int)
	space := new(big.Int)

	nonce.Mod(raw, shr)
	space.Div(raw, shr)

	return space, nonce
}

// GenerateRandomNonce returns a ................................................
//
// This function can be used to get a random space for a nonce
// to ensure transactions can be executed in parallel
func GenerateRandomNonce() (*big.Int, error) {
	// Max random value i.e 2^100 - 1, since nonce space is the first
	// 160 bits in a nonce. We skip the first 60 bits to keep
	// some of the nonce space clean.
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(100), nil).Sub(max, big.NewInt(1))

	//Generate random number between 0 and 2**160-1
	nonce, err := rand.Int(rand.Reader, max)
	if err != nil {
		return new(big.Int).Set(nil), err
	}

	// Right pad the nonce by 96 bits so it's a full uint256
	nonce = new(big.Int).Lsh(nonce, 96)
	return nonce, nil
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
				Type: MustNewType("string"),
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
