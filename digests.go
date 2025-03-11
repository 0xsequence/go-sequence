package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/go-sequence/core"
)

/*
- MainModule.execute digest
  = keccak256(nonce, transactions)

- MetaTxnID
  = SubDigest(walletAddress, chainID, MainModule.execute digest)
  = keccak256("\x19\x01", chainID, walletAddress, MainModule.execute digest)

- MainModule.selfExecute digest
  = SubDigest(walletAddress, chainID, keccak256("self:", transactions))
  = keccak256("\x19\x01", chainID, walletAddress, keccak256("self:", transactions))

- GuestModule.execute digest
  = SubDigest(guestModuleAddress, chainID, keccak256("guest:", transactions))
  = keccak256("\x19\x01", chainID, guestModuleAddress, keccak256("guest:", transactions))

- the digest that is actually signed (eth_sign)
  = keccak256("\x19Ethereum Signed Message:\n32", MetaTxnID)
*/

func ExecuteDigest(transactions Transactions, nonce *big.Int) (core.Digest, error) {
	if nonce == nil {
		return core.Digest{}, fmt.Errorf("no nonce")
	}
	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return core.Digest{}, fmt.Errorf("unable to encode transactions: %w", err)
	}
	message, err := abiTransactionsDigestType.Pack(nonce, encoded)
	if err != nil {
		return core.Digest{}, fmt.Errorf("unable to pack nonce and transactions: %w", err)
	}
	return core.NewDigest(message), nil
}

func SelfExecuteDigest(transactions Transactions) (core.Digest, error) {
	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return core.Digest{}, fmt.Errorf("unable to encode transactions: %w", err)
	}
	message, err := abiTransactionsStringDigestType.Pack("self:", encoded)
	if err != nil {
		return core.Digest{}, fmt.Errorf(`unable to pack "self:" and transactions: %w`, err)
	}
	return core.NewDigest(message), nil
}

func GuestExecuteDigest(transactions Transactions) (core.Digest, error) {
	encoded, err := transactions.EncodedTransactions()
	if err != nil {
		return core.Digest{}, fmt.Errorf("unable to encode transactions: %w", err)
	}
	message, err := abiTransactionsStringDigestType.Pack("guest:", encoded)
	if err != nil {
		return core.Digest{}, fmt.Errorf(`unable to pack "guest:" and transactions: %w`, err)
	}
	return core.NewDigest(message), nil
}
