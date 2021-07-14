package sequence

import (
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

/*
- MainModule.execute digest
  = keccak256(nonce, transactions)

- MetaTxnID
  = SubDigest(walletAddress, chainID, MainModule.execute digest)
  = keccak256(chainID, walletAddress, MainModule.execute digest)

- MainModule.selfExecute digest
  = SubDigest(walletAddress, chainID, keccak256("self:", transactions))
  = keccak256("\x19\x01", chainID, walletAddress, keccak256("self:", transactions))

- GuestModule.execute digest
  = SubDigest(guestModuleAddress, chainID, keccak256("guest:", transactions))
  = keccak256("\x19\x01", chainID, guestModuleAddress, keccak256("guest:", transactions))


- the digest that is actually signed (eth_sign)
  = keccak256("\x19Ethereum Signed Message:\n32", MetaTxnID)
*/

// MetaTxnExecType represents the kind of execution call for the meta-transaction
type MetaTxnExecType uint32

const (
	MetaTxnWalletExec MetaTxnExecType = iota // MainModule.execute
	MetaTxnSelfExec                          // MainModule.selfExecute
	MetaTxnGuestExec                         // GuestModule.execute
)

func ComputeWalletExecDigest(nonce *big.Int, txns Transactions) (common.Hash, error) {
	if nonce == nil {
		return common.Hash{}, fmt.Errorf("nonce is required for wallet execute")
	}
	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return common.Hash{}, err
	}
	message, err := abiTransactionsDigestType.Pack(nonce, encodedTxns)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack nonce and transactions: %w", err)
	}
	return common.BytesToHash(ethcoder.Keccak256(message)), nil
}

func ComputeSelfExecDigest(txns Transactions) (common.Hash, error) {
	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return common.Hash{}, err
	}
	message, err := abiTransactionsStringDigestType.Pack("self:", encodedTxns)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack \"self:\" and transactions: %w", err)
	}
	return common.BytesToHash(ethcoder.Keccak256(message)), nil
}

func ComputeGuestExecDigest(txns Transactions) (common.Hash, error) {
	encodedTxns, err := txns.EncodedTransactions()
	if err != nil {
		return common.Hash{}, err
	}
	message, err := abiTransactionsStringDigestType.Pack("guest:", encodedTxns)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to pack \"guest:\" and transactions: %w", err)
	}
	return common.BytesToHash(ethcoder.Keccak256(message)), nil
}

func ComputeMetaTxnID(chainID *big.Int, address common.Address, txns Transactions, nonce *big.Int, execType MetaTxnExecType) (MetaTxnID, common.Hash, error) {
	var digest common.Hash
	var err error

	switch execType {
	case MetaTxnWalletExec:
		if nonce == nil {
			return "", common.Hash{}, fmt.Errorf("wallet exec type requires nonce but is nil")
		}
		digest, err = ComputeWalletExecDigest(nonce, txns)

	case MetaTxnSelfExec:
		digest, err = ComputeSelfExecDigest(txns)

	case MetaTxnGuestExec:
		digest, err = ComputeGuestExecDigest(txns)

	default:
		return "", common.Hash{}, fmt.Errorf("unknown exec type")
	}
	if err != nil {
		return "", common.Hash{}, err
	}

	return ComputeMetaTxnIDFromDigest(chainID, address, digest)
}

func ComputeMetaTxnIDFromDigest(chainID *big.Int, address common.Address, digest common.Hash) (MetaTxnID, common.Hash, error) {
	subDigest, err := SubDigest(chainID, address, digest)
	if err != nil {
		return "", common.Hash{}, nil
	}

	metaTxnIDHex := ethcoder.HexEncode(subDigest)
	if len(metaTxnIDHex) != 66 {
		return "", common.Hash{}, fmt.Errorf("computed meta txn id is invalid length")
	}
	return MetaTxnID(metaTxnIDHex[2:]), common.BytesToHash(subDigest), nil
}

func SubDigest(chainID *big.Int, address common.Address, digest common.Hash) ([]byte, error) {
	if chainID == nil {
		return nil, ErrUnknownChainID
	}

	// sequence smart wallet uses additional encoding of the digest in IsValidSignature()
	packedData, err := PackMessageData(chainID, address, digest)
	if err != nil {
		return nil, fmt.Errorf("subDigest, packageMessageData failed: %w", err)
	}

	// returns subdigest
	return ethcoder.Keccak256(packedData), nil
}

// PackMessageData encodes a Sequence contract "message"
func PackMessageData(chainID *big.Int, address common.Address, digest common.Hash) ([]byte, error) {
	if chainID == nil {
		return nil, ErrUnknownChainID
	}
	output, err := ethcoder.SolidityPack([]string{"string", "uint256", "address", "bytes"}, []interface{}{
		"\x19\x01", chainID, address, digest[:],
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}
