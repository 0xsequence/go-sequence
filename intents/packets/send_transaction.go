package packets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/intents"
)

type SendTransactionsPacket struct {
	Code         string            `json:"code"`
	Wallet       string            `json:"wallet"`
	Network      string            `json:"network"`
	Transactions []json.RawMessage `json:"transactions"`
}

const SendTransactionCode = "sendTransaction"

func (p *SendTransactionsPacket) Unmarshal(i *intents.Intent) error {
	err := json.Unmarshal(i.Packet, &p)
	if err != nil {
		return err
	}

	if p.Code != SendTransactionCode {
		return fmt.Errorf("packet code is not '%s', got '%s'", SendTransactionCode, p.Code)
	}

	return nil
}

func (p *SendTransactionsPacket) chainID() (*big.Int, error) {
	n, ok := sequence.ParseHexOrDec(p.Network)
	if !ok {
		return nil, fmt.Errorf("invalid network id '%s'", p.Network)
	}

	return n, nil
}

func (p *SendTransactionsPacket) wallet() common.Address {
	return common.HexToAddress(p.Wallet)
}

type expectedValuesForTransaction struct {
	To    *common.Address
	Value *big.Int
	Data  []byte
}

func (p *SendTransactionsPacket) expectedValuesFor(subpacket *json.RawMessage) (*expectedValuesForTransaction, error) {
	// Get the subpacket type
	var subpacketType struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(*subpacket, &subpacketType)
	if err != nil {
		return nil, err
	}

	switch subpacketType.Type {
	case "transaction":
		// This packet explicitly defines the transaction values
		var subpacketTransactionType struct {
			To    string `json:"to"`
			Value string `json:"value"`
			Data  string `json:"data"`
		}

		err := json.Unmarshal(*subpacket, &subpacketTransactionType)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(subpacketTransactionType.To)
		value, ok := sequence.ParseHexOrDec(subpacketTransactionType.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", subpacketTransactionType.Value)
		}

		data := common.Hex2Bytes(strings.TrimPrefix(subpacketTransactionType.Data, "0x"))

		return &expectedValuesForTransaction{
			To:    &to,
			Value: value,
			Data:  data,
		}, nil

	case "erc20send":
		// This packet defines the transaction values for an ERC20 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		var subpacketERC20SendType struct {
			Token string `json:"token"`
			To    string `json:"to"`
			Value string `json:"value"`
		}

		err := json.Unmarshal(*subpacket, &subpacketERC20SendType)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(subpacketERC20SendType.To)
		token := common.HexToAddress(subpacketERC20SendType.Token)
		value, ok := sequence.ParseHexOrDec(subpacketERC20SendType.Value)
		if !ok {
			return nil, fmt.Errorf("invalid value '%s'", subpacketERC20SendType.Value)
		}

		// Encode the transfer call
		data, err := ethcoder.AbiEncodeMethodCalldata("transfer(address,uint256)", []interface{}{to, value})
		if err != nil {
			return nil, err
		}

		return &expectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  data,
		}, nil

	case "erc721send":
		// This packet defines the transaction values for an ERC721 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		var subpacketERC721SendType struct {
			Token string `json:"token"`
			To    string `json:"to"`
			ID    string `json:"id"`
			Safe  bool   `json:"safe,omitempty"`
			Data  string `json:"data,omitempty"`
		}

		// Safe defaults to false
		if err := json.Unmarshal(*subpacket, &subpacketERC721SendType); err != nil {
			return nil, err
		}

		// If data is not empty, then safe *must* be true

		to := common.HexToAddress(subpacketERC721SendType.To)
		token := common.HexToAddress(subpacketERC721SendType.Token)
		id, ok := sequence.ParseHexOrDec(subpacketERC721SendType.ID)
		if !ok {
			return nil, fmt.Errorf("invalid id '%s'", subpacketERC721SendType.ID)
		}
		data := common.Hex2Bytes(strings.TrimPrefix(subpacketERC721SendType.Data, "0x"))

		// If data is not empty, then safe *must* be true
		if len(data) > 0 && !subpacketERC721SendType.Safe {
			return nil, fmt.Errorf("safe must be true if data is not empty")
		}

		var encodedData []byte
		if subpacketERC721SendType.Safe {
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

		return &expectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  encodedData,
		}, nil

	case "erc1155send":
		// This packet defines the transaction values for an ERC1155 transfer
		// so this should be an ABI encoded transfer call to `to`. The `value`
		// field must be 0.
		type subpacketERC1155SendValsType struct {
			ID     string `json:"id"`
			Amount string `json:"amount"`
		}

		var subpacketERC1155SendType struct {
			Token string                         `json:"token"`
			To    string                         `json:"to"`
			Vals  []subpacketERC1155SendValsType `json:"vals"`
			Data  string                         `json:"data,omitempty"`
		}

		err := json.Unmarshal(*subpacket, &subpacketERC1155SendType)
		if err != nil {
			return nil, err
		}

		to := common.HexToAddress(subpacketERC1155SendType.To)
		token := common.HexToAddress(subpacketERC1155SendType.Token)

		var parsedIDs []*big.Int
		var parsedAmounts []*big.Int

		for _, val := range subpacketERC1155SendType.Vals {
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

		data := common.Hex2Bytes(strings.TrimPrefix(subpacketERC1155SendType.Data, "0x"))

		encodedData, err := ethcoder.AbiEncodeMethodCalldata("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)", []interface{}{p.wallet(), to, parsedIDs, parsedAmounts, data})

		if err != nil {
			return nil, err
		}

		return &expectedValuesForTransaction{
			To:    &token,
			Value: big.NewInt(0),
			Data:  encodedData,
		}, nil
	default:
		return nil, fmt.Errorf("invalid subpacket type '%s'", subpacketType.Type)
	}
}

func (p *SendTransactionsPacket) IsValidInterpretation(subdigest common.Hash, txns sequence.Transactions, nonce *big.Int) bool {
	// Compare the digest with the provided transactions
	// otherwise we can't be sure that the subdigest belongs to the transactions
	bundle := sequence.Transaction{
		Transactions: txns,
		Nonce:        nonce,
	}

	calcDigest, err := bundle.Digest()
	if err != nil {
		return false
	}

	chainID, err := p.chainID()
	if err != nil {
		return false
	}

	calcSubdigest, err := sequence.SubDigest(chainID, p.wallet(), calcDigest)
	if err != nil || !bytes.Equal(calcSubdigest, subdigest[:]) {
		return false
	}

	// Now check that every transaction maps 1:1 to the transactions in the packet
	// meaning that they follow the intent signed by it
	if len(txns) != len(p.Transactions) {
		return false
	}

	for i, txn := range txns {
		if txn.DelegateCall != false {
			return false
		}

		expected, err := p.expectedValuesFor(&p.Transactions[i])
		if err != nil {
			return false
		}

		if !bytes.Equal(txn.To.Bytes(), expected.To.Bytes()) {
			return false
		}

		if txn.Value.Cmp(expected.Value) != 0 {
			return false
		}

		if !bytes.Equal(txn.Data, expected.Data) {
			return false
		}
	}

	return true
}
