package packets

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/intents"
	"github.com/stretchr/testify/assert"
)

func TestRecoverTransactionIntent(t *testing.T) {
	data := `{
		"version": "1",
		"packet": {
		  "code": "sendTransaction",
		  "wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
		  "network": "10",
		  "transactions": [
			{
			  "type": "transaction",
			  "to": "0x479F6a5b0C1728947318714963a583C56A78366A",
			  "value": "39381",
			  "data": "0x3251ba32"
			},
			{
			  "type": "erc20send",
			  "token": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			  "to": "0x7b1Bd3474D789e18e2E329E2c53F819B6E687b4A",
			  "value": "1000"
			},
			{
			  "type": "erc721send",
			  "token": "0xF87E31492Faf9A91B02Ee0dEAAd50d51d56D5d4d",
			  "to": "0x17fFA2d95b58228e1ECb0C6Ac25A6EfD20BA08E4",
			  "id": "7",
			  "safe": true,
			  "data": "0x112233"
			},
			{
			  "type": "erc1155send",
			  "token": "0x631998e91476da5b870d741192fc5cbc55f5a52e",
			  "to": "0x91E8aC543C5fEDf9F3Ef8b9dA1500dB84305681F",
			  "vals": [
				{
				  "id": "2",
				  "amount": "5"
				},
				{
				  "id": "500",
				  "amount": "1"
				}
			  ],
			  "data": "0x223344"
			}
		  ]
		},
		"signatures": [
		  {
			"session": "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B",
			"signature": "0x71a0d20a974e8a7cc3643c9ab0cf6f52bbeda20bf486449a6287c09c5f6c7d0a478a5ed8d0747f83c9a8851d9b87c05c1c74e28a0a3a428aef7b067c4bcbb1bb1b"
		  }
		]
	}`

	intent := &intents.Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1", intent.Version)
	assert.Equal(t, "sendTransaction", intent.PacketCode())

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.Equal(t, common.Bytes2Hex(hash), "f7bbc2ad0d74cccdc1d6a8eae95f39f6f3f7d29f571e8ced744f626168913a00")

	signers := intent.Signers()
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])

	packet := &SendTransactionsPacket{}
	err = packet.Unmarshal(intent)
	assert.Nil(t, err)

	assert.Equal(t, "sendTransaction", packet.Code)

	// Generate transactions as sequence.Wallet would
	nonce := big.NewInt(33151233)
	chainID := big.NewInt(10)

	transactions := make(sequence.Transactions, 0)
	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0x479F6a5b0C1728947318714963a583C56A78366A"),
		Value:         big.NewInt(39381),
		Data:          common.FromHex("0x3251ba32"),
	})

	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		Value:         big.NewInt(0),
		Data:          common.FromHex("0xa9059cbb0000000000000000000000007b1bd3474d789e18e2e329e2c53f819b6e687b4a00000000000000000000000000000000000000000000000000000000000003e8"),
	})

	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0xF87E31492Faf9A91B02Ee0dEAAd50d51d56D5d4d"),
		Value:         big.NewInt(0),
		Data:          common.FromHex("0xb88d4fde000000000000000000000000d67fc48b298b09ed3d03403d930769c527186c4e00000000000000000000000017ffa2d95b58228e1ecb0c6ac25a6efd20ba08e40000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000031122330000000000000000000000000000000000000000000000000000000000"),
	})

	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0x631998e91476da5b870d741192fc5cbc55f5a52e"),
		Value:         big.NewInt(0),
		Data:          common.FromHex("0x2eb2c2d6000000000000000000000000d67fc48b298b09ed3d03403d930769c527186c4e00000000000000000000000091e8ac543c5fedf9f3ef8b9da1500db84305681f00000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000001f400000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000032233440000000000000000000000000000000000000000000000000000000000"),
	})

	bundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        nonce,
	}

	digest, err := bundle.Digest()
	assert.Nil(t, err)

	subdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		digest,
	)
	assert.Nil(t, err)

	assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))

	// changing any transaction value should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Value
		transactions[i].Value = big.NewInt(123)
		assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
		transactions[i].Value = prev
		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
	}

	// changing any transaction data should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Data
		transactions[i].Data = common.Hex2Bytes("0x1234")
		assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
		transactions[i].Data = prev
		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
	}

	// changing any to address should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].To
		transactions[i].To = common.HexToAddress("0xd1333D70A344c26041a869077381209462e586F8")
		assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
		transactions[i].To = prev
		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
	}

	// setting any delegate call should invalidate the interpretation
	for i := range transactions {
		transactions[i].DelegateCall = true
		assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
		transactions[i].DelegateCall = false
		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
	}

	// changing revert on error should NOT invalidate the interpretation
	for i := range transactions {
		transactions[i].RevertOnError = false

		nxtBundle := sequence.Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		nxtdigest, err := nxtBundle.Digest()
		assert.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		assert.Nil(t, err)

		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce))
		transactions[i].RevertOnError = true
	}

	// changing any gas limit should NOT invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].GasLimit
		transactions[i].GasLimit = big.NewInt(123)

		nxtBundle := sequence.Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		nxtdigest, err := nxtBundle.Digest()
		assert.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		assert.Nil(t, err)

		assert.True(t, packet.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce))

		transactions[i].GasLimit = prev
	}

	// changing the nonce should NOT invalidate the interpretation
	nxtBundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        big.NewInt(123),
	}
	nxtdigest, err := nxtBundle.Digest()
	assert.Nil(t, err)
	nxtsubdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		nxtdigest,
	)
	assert.Nil(t, err)
	assert.True(t, packet.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, big.NewInt(123)))

	// removing a transaction should invalidate the interpretation
	assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions[1:], nonce))

	// adding an extra transaction should invalidate the interpretation
	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0x479F6a5b0C1728947318714963a583C56A78366A"),
		Value:         big.NewInt(39381),
		Data:          common.FromHex("0x3251ba32"),
	})

	assert.False(t, packet.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce))
}
