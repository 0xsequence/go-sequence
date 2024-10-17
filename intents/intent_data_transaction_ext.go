package intents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/core/types"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/relayer/proto"
	"github.com/davecgh/go-spew/spew"
)

func (p *IntentDataSendTransaction) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}

	return n, nil
}

func (p *IntentDataSendTransaction) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

type ExpectedValuesForTransaction struct {
	To    *common.Address
	Value *big.Int
	Data  []byte
}

func (p *IntentDataSendTransaction) ExpectedValuesFor(txRaw *json.RawMessage) (*ExpectedValuesForTransaction, error) {
	// Get the tx type
	var baseTransactionType struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(*txRaw, &baseTransactionType)
	if err != nil {
		return nil, err
	}

	switch baseTransactionType.Type {
	case "transaction":
		// This struct explicitly defines the transaction values
		var tx TransactionRaw

		err := json.Unmarshal(*txRaw, &tx)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(tx.To)
		if tx.Value == "" {
			tx.Value = "0"
		}
		value, ok := sequence.ParseHexOrDec(tx.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", tx.Value)
		}

		data := common.FromHex(tx.Data)

		return &ExpectedValuesForTransaction{
			To:    &to,
			Value: value,
			Data:  data,
		}, nil

	case "erc20send":
		// This struct defines the transaction values for an ERC20 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		var tx TransactionERC20

		err := json.Unmarshal(*txRaw, &tx)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(tx.To)
		token := common.HexToAddress(tx.TokenAddress)
		value, ok := sequence.ParseHexOrDec(tx.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", tx.Value)
		}

		// Encode the transfer call
		data, err := ethcoder.AbiEncodeMethodCalldata("transfer(address,uint256)", []interface{}{to, value})
		if err != nil {
			return nil, err
		}

		return &ExpectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  data,
		}, nil

	case "erc721send":
		// This struct defines the transaction values for an ERC721 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		var tx TransactionERC721

		// Safe defaults to false
		if err := json.Unmarshal(*txRaw, &tx); err != nil {
			return nil, err
		}

		// If data is not empty, then safe *must* be true

		to := common.HexToAddress(tx.To)
		token := common.HexToAddress(tx.TokenAddress)
		id, ok := sequence.ParseHexOrDec(tx.Id)
		if !ok {
			return nil, fmt.Errorf("invalid id '%s'", tx.Id)
		}
		var data []byte
		if tx.Data != nil {
			data = common.FromHex(*tx.Data)
		}

		// If data is not empty, then safe *must* be true
		if len(data) > 0 && tx.Safe != nil && !*tx.Safe {
			return nil, fmt.Errorf("safe must be true if data is not empty")
		}

		var encodedData []byte
		if tx.Safe != nil && *tx.Safe {
			// Encode the safe transfer call
			encodedData, err = ethcoder.AbiEncodeMethodCalldata("safeTransferFrom(address,address,uint256,bytes)", []interface{}{p.wallet(), to, id, data})
			if err != nil {
				return nil, err
			}
		} else {
			// Encode the transfer call
			encodedData, err = ethcoder.AbiEncodeMethodCalldata("transferFrom(address,address,uint256)", []interface{}{p.wallet(), to, id})
			if err != nil {
				return nil, err
			}
		}

		return &ExpectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  encodedData,
		}, nil

	case "erc1155send":
		// This struct defines the transaction values for an ERC1155 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		var tx TransactionERC1155

		err := json.Unmarshal(*txRaw, &tx)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(tx.To)
		token := common.HexToAddress(tx.TokenAddress)

		var parsedIDs []*big.Int
		var parsedAmounts []*big.Int

		for _, val := range tx.Vals {
			id, ok := sequence.ParseHexOrDec(val.ID)
			if !ok {
				return nil, fmt.Errorf("invalid id '%s'", val.ID)
			}

			amount, ok := sequence.ParseHexOrDec(val.Amount)
			if !ok {
				return nil, fmt.Errorf("invalid amount '%s'", val.Amount)
			}

			parsedIDs = append(parsedIDs, id)
			parsedAmounts = append(parsedAmounts, amount)
		}

		var data []byte
		if tx.Data != nil {
			data = common.FromHex(*tx.Data)
		}

		encodedData, err := ethcoder.AbiEncodeMethodCalldata("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)", []interface{}{p.wallet(), to, parsedIDs, parsedAmounts, data})

		if err != nil {
			return nil, err
		}

		return &ExpectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  encodedData,
		}, nil

	case "contractCall":
		var tx TransactionContractCall

		err := json.Unmarshal(*txRaw, &tx)
		if err != nil {
			return nil, err
		}

		nst := &contractCallType{}
		nst.Abi = tx.Data.Abi
		if tx.Data.Func != nil {
			nst.Func = *tx.Data.Func
		}
		nst.Args = tx.Data.Args

		spew.Dump("$$$WEE$$", nst)

		encoded, err := EncodeContractCall(nst)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(tx.To)
		if tx.Value == "" {
			tx.Value = "0"
		}
		value, ok := sequence.ParseHexOrDec(tx.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", tx.Value)
		}

		return &ExpectedValuesForTransaction{
			To:    &to,
			Value: value,
			Data:  common.FromHex(encoded),
		}, nil

	// NOTE: 'delayedEncode' is deprecated, should use 'contractCall' instead, we're leaving
	// it here for backwards compatibility.
	case "delayedEncode":
		var tx TransactionDelayedEncode

		err := json.Unmarshal(*txRaw, &tx)
		if err != nil {
			return nil, err
		}

		nst := &delayedEncodeType{}
		err = json.Unmarshal(tx.Data, nst)
		if err != nil {
			return nil, err
		}

		encoded, err := EncodeDelayedABI(nst)
		if err != nil {
			return nil, err
		}
		to := common.HexToAddress(tx.To)
		value, ok := sequence.ParseHexOrDec(tx.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", tx.Value)
		}

		return &ExpectedValuesForTransaction{
			To:    &to,
			Value: value,
			Data:  common.FromHex(encoded),
		}, nil
	default:
		return nil, fmt.Errorf("invalid transaction type '%s'", baseTransactionType.Type)
	}
}

func (p *IntentDataSendTransaction) Nonce() (*big.Int, error) {
	// Hash the identifier, it will be used as the nonce
	// space. The nonce number is always 0.
	hashed := ethcoder.Keccak256([]byte(p.Identifier))

	// The space contains only 160 bits
	return sequence.EncodeNonce(big.NewInt(0).SetBytes(hashed[:20]), common.Big0)
}

func (p *IntentDataSendTransaction) IsValidInterpretation(subdigest common.Hash, txns sequence.Transactions, nonce *big.Int) (bool, error) {
	// Nonce must be the expected one
	// (defined by the identifier)
	enonce, err := p.Nonce()
	if err != nil {
		return false, fmt.Errorf("invalid nonce: %w", err)
	}

	if enonce.Cmp(nonce) != 0 {
		return false, fmt.Errorf("invalid nonce")
	}

	// Compare the digest with the provided transactions
	// otherwise we can't be sure that the subdigest belongs to the transactions
	bundle := sequence.Transaction{
		Transactions: txns,
		Nonce:        nonce,
	}

	calcDigest, err := bundle.Digest()
	if err != nil {
		return false, fmt.Errorf("invalid bundle digest: %w", err)
	}

	chainID, err := p.chainID()
	if err != nil {
		return false, err
	}

	calcSubdigest, err := sequence.SubDigest(chainID, p.wallet(), calcDigest)
	if err != nil || !bytes.Equal(calcSubdigest, subdigest[:]) {
		return false, fmt.Errorf("invalid subdigest: %w", err)
	}

	// Now check that every transaction maps 1:1 to the transactions in the intent
	// meaning that they follow the intent signed by it
	if len(txns) != len(p.Transactions) {
		return false, fmt.Errorf("intent transaction count mismatch")
	}

	for i, txn := range txns {
		if txn.DelegateCall {
			return false, fmt.Errorf("delegate call not allowed")
		}

		expected, err := p.ExpectedValuesFor(&p.Transactions[i])
		if err != nil {
			return false, fmt.Errorf("invalid transaction: %w", err)
		}

		if !bytes.Equal(txn.To.Bytes(), expected.To.Bytes()) {
			return false, fmt.Errorf("invalid to address")
		}

		if txn.Value.Cmp(expected.Value) != 0 {
			return false, fmt.Errorf("invalid value")
		}

		if !bytes.Equal(txn.Data, expected.Data) {
			return false, fmt.Errorf("invalid data")
		}
	}

	return true, nil
}

type SendTransactionResponse struct {
	Code string `json:"code"`
	Data struct {
		Request       *IntentDataSendTransaction `json:"request"`
		TxHash        string                     `json:"txHash"`
		Receipt       *proto.MetaTxnReceipt      `json:"receipt"`
		NativeReceipt *types.Receipt             `json:"nativeReceipt"`
		Simulations   []*proto.SimulateResult    `json:"simulations"`
	}
}

const SendTransactionResponseCode = "transactionReceipt"

type SendTransactionFailed struct {
	Code string `json:"code"`
	Data struct {
		Request     *IntentDataSendTransaction `json:"request"`
		Simulations []*proto.SimulateResult    `json:"simulations"`
	}
}

const SendTransactionFailedCode = "transactionFailed"
