package intents

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/0xsequence/go-sequence"
)

func TestRecoverTransactionIntent(t *testing.T) {
	data := `{
		"version": "1",
		"name": "sendTransaction",
		"issued": 0,
		"expires": 0,
		"data": {
		  "code": "sendTransaction",
			"identifier": "test-identifier",
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
			  "tokenAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			  "to": "0x7b1Bd3474D789e18e2E329E2c53F819B6E687b4A",
			  "value": "1000"
			},
			{
			  "type": "erc721send",
			  "tokenAddress": "0xF87E31492Faf9A91B02Ee0dEAAd50d51d56D5d4d",
			  "to": "0x17fFA2d95b58228e1ECb0C6Ac25A6EfD20BA08E4",
			  "id": "7",
			  "safe": true,
			  "data": "0x112233"
			},
			{
			  "type": "erc1155send",
			  "tokenAddress": "0x631998e91476da5b870d741192fc5cbc55f5a52e",
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
			},
			{
				"type": "contractCall",
        "to": "0x140d72763D1ce39Ad4E2e73EC6e8FC53E5b73B64",
				"data": {
					"abi": "fillOrKillOrder(uint256 orderId, uint256 maxCost, address[] fees, bytes data)",
					"args": [
						"48774435471364917511246724398022004900255301025912680232738918790354204737320",
						"1000000000000000000",
						["0x8541D65829f98f7D71A4655cCD7B2bB8494673bF"],
						{
							"abi": "notExpired(uint256,string)",
							"args": [
								"1600000000",
								"Nov 1st, 2020"
							]
						}
					]
				}
			}
		  ]
		},
		"signatures": []
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	require.Nil(t, err)

	require.Equal(t, "1", intent.Version)
	require.Equal(t, IntentName_sendTransaction, intent.Name)

	hash, err := intent.Hash()
	require.Nil(t, err)
	require.NotNil(t, common.Bytes2Hex(hash))

	intent.IssuedAt = uint64(time.Now().Unix())
	intent.ExpiresAt = uint64(time.Now().Unix()) + 60

	wallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.Nil(t, err)

	session := NewSessionP256K1(wallet)

	err = session.Sign(intent)
	require.Nil(t, err)

	intentTyped, err := NewIntentTypedFromIntent[IntentDataSendTransaction](intent)
	require.Nil(t, err)

	signers := intent.Signers()
	require.Equal(t, 1, len(signers))
	require.Equal(t, "0x"+common.Bytes2Hex(append([]byte{0x00}, wallet.Address().Bytes()...)), signers[0])

	require.NoError(t, intentTyped.IsValid())

	sendTransactionData := intentTyped.Data

	// Generate transactions as sequence.Wallet would
	nonce, err := sendTransactionData.Nonce()
	require.Nil(t, err)

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

	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(0),
		To:            common.HexToAddress("0x140d72763D1ce39Ad4E2e73EC6e8FC53E5b73B64"),
		Value:         big.NewInt(0),
		Data:          common.FromHex("0x6365f1646bd55a2877890bd58871eefe886770a7734077a74981910a75d7b1f044b5bf280000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000008541d65829f98f7d71a4655ccd7b2bb8494673bf000000000000000000000000000000000000000000000000000000000000008446c421fa000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000d4e6f76203173742c20323032300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	})

	bundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        nonce,
	}

	digest, err := bundle.Digest()
	require.Nil(t, err)

	subdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		digest,
	)
	require.Nil(t, err)

	isValid, err := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
	require.NoError(t, err)
	require.True(t, isValid)

	// changing any transaction value should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Value
		transactions[i].Value = big.NewInt(123)
		isValid, err = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.Error(t, err)
		require.False(t, isValid)
		transactions[i].Value = prev
		isValid, err = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.NoError(t, err)
		require.True(t, isValid)
	}

	// changing any transaction data should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Data
		transactions[i].Data = common.Hex2Bytes("0x1234")
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].Data = prev
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// changing any to address should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].To
		transactions[i].To = common.HexToAddress("0xd1333D70A344c26041a869077381209462e586F8")
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].To = prev
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// setting any delegate call should invalidate the interpretation
	for i := range transactions {
		transactions[i].DelegateCall = true
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].DelegateCall = false
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// changing revert on error should NOT invalidate the interpretation
	for i := range transactions {
		transactions[i].RevertOnError = false

		nxtBundle := sequence.Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		nxtdigest, err := nxtBundle.Digest()
		require.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		require.Nil(t, err)

		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce)
		require.True(t, ok)
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
		require.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		require.Nil(t, err)

		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce)
		require.True(t, ok)

		transactions[i].GasLimit = prev
	}

	// changing the nonce should invalidate the interpretation
	nxtBundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        big.NewInt(123),
	}
	nxtdigest, err := nxtBundle.Digest()
	require.Nil(t, err)
	nxtsubdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		nxtdigest,
	)
	require.Nil(t, err)
	ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, big.NewInt(123))
	require.False(t, ok)

	// removing a transaction should invalidate the interpretation
	ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions[1:], nonce)
	require.False(t, ok)

	// adding an extra transaction should invalidate the interpretation
	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0x479F6a5b0C1728947318714963a583C56A78366A"),
		Value:         big.NewInt(39381),
		Data:          common.FromHex("0x3251ba32"),
	})

	ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
	require.False(t, ok)
}

// NOTE: DelayedEncode is deprecated, see TestRecoverTransactionIntent below which uses contractCall type
func TestDelayedEncodeRecoverTransactionIntent(t *testing.T) {
	data := `{
		"version": "1",
		"name": "sendTransaction",
		"issued": 0,
		"expires": 0,
		"data": {
		  "code": "sendTransaction",
			"identifier": "test-identifier",
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
			  "tokenAddress": "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			  "to": "0x7b1Bd3474D789e18e2E329E2c53F819B6E687b4A",
			  "value": "1000"
			},
			{
			  "type": "erc721send",
			  "tokenAddress": "0xF87E31492Faf9A91B02Ee0dEAAd50d51d56D5d4d",
			  "to": "0x17fFA2d95b58228e1ECb0C6Ac25A6EfD20BA08E4",
			  "id": "7",
			  "safe": true,
			  "data": "0x112233"
			},
			{
			  "type": "erc1155send",
			  "tokenAddress": "0x631998e91476da5b870d741192fc5cbc55f5a52e",
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
			},
			{
				"type": "delayedEncode",
        "to": "0x140d72763D1ce39Ad4E2e73EC6e8FC53E5b73B64",
        "value": "0",
				"data": {
					"abi": "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCost\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_fees\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"notExpired\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherMethods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
					"func": "fillOrKillOrder",
					"args": [
						"48774435471364917511246724398022004900255301025912680232738918790354204737320",
						"1000000000000000000",
						"[\"0x8541D65829f98f7D71A4655cCD7B2bB8494673bF\"]",
						{
							"abi": "notExpired(uint256,string)",
							"func": "notExpired",
							"args": [
								"1600000000",
								"Nov 1st, 2020"
							]
						}
					]
				}
			}
		  ]
		},
		"signatures": []
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	require.Nil(t, err)

	require.Equal(t, "1", intent.Version)
	require.Equal(t, IntentName_sendTransaction, intent.Name)

	hash, err := intent.Hash()
	require.Nil(t, err)
	require.NotNil(t, common.Bytes2Hex(hash))

	intent.IssuedAt = uint64(time.Now().Unix())
	intent.ExpiresAt = uint64(time.Now().Unix()) + 60

	wallet, err := ethwallet.NewWalletFromRandomEntropy()
	require.Nil(t, err)

	session := NewSessionP256K1(wallet)

	err = session.Sign(intent)
	require.Nil(t, err)

	intentTyped, err := NewIntentTypedFromIntent[IntentDataSendTransaction](intent)
	require.Nil(t, err)

	signers := intent.Signers()
	require.Equal(t, 1, len(signers))
	require.Equal(t, "0x"+common.Bytes2Hex(append([]byte{0x00}, wallet.Address().Bytes()...)), signers[0])

	require.NoError(t, intentTyped.IsValid())

	sendTransactionData := intentTyped.Data

	// Generate transactions as sequence.Wallet would
	nonce, err := sendTransactionData.Nonce()
	require.Nil(t, err)

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

	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(0),
		To:            common.HexToAddress("0x140d72763D1ce39Ad4E2e73EC6e8FC53E5b73B64"),
		Value:         big.NewInt(0),
		Data:          common.FromHex("0x6365f1646bd55a2877890bd58871eefe886770a7734077a74981910a75d7b1f044b5bf280000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000008541d65829f98f7d71a4655ccd7b2bb8494673bf000000000000000000000000000000000000000000000000000000000000008446c421fa000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000d4e6f76203173742c20323032300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
	})

	bundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        nonce,
	}

	digest, err := bundle.Digest()
	require.Nil(t, err)

	subdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		digest,
	)
	require.Nil(t, err)

	ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
	require.True(t, ok)

	// changing any transaction value should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Value
		transactions[i].Value = big.NewInt(123)
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].Value = prev
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// changing any transaction data should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].Data
		transactions[i].Data = common.Hex2Bytes("0x1234")
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].Data = prev
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// changing any to address should invalidate the interpretation
	for i := range transactions {
		prev := transactions[i].To
		transactions[i].To = common.HexToAddress("0xd1333D70A344c26041a869077381209462e586F8")
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].To = prev
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// setting any delegate call should invalidate the interpretation
	for i := range transactions {
		transactions[i].DelegateCall = true
		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.False(t, ok)
		transactions[i].DelegateCall = false
		ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
		require.True(t, ok)
	}

	// changing revert on error should NOT invalidate the interpretation
	for i := range transactions {
		transactions[i].RevertOnError = false

		nxtBundle := sequence.Transaction{
			Transactions: transactions,
			Nonce:        nonce,
		}
		nxtdigest, err := nxtBundle.Digest()
		require.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		require.Nil(t, err)

		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce)
		require.True(t, ok)
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
		require.Nil(t, err)
		nxtsubdigest, err := sequence.SubDigest(
			chainID,
			common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
			nxtdigest,
		)
		require.Nil(t, err)

		ok, _ := sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, nonce)
		require.True(t, ok)

		transactions[i].GasLimit = prev
	}

	// changing the nonce should invalidate the interpretation
	nxtBundle := sequence.Transaction{
		Transactions: transactions,
		Nonce:        big.NewInt(123),
	}
	nxtdigest, err := nxtBundle.Digest()
	require.Nil(t, err)
	nxtsubdigest, err := sequence.SubDigest(
		chainID,
		common.HexToAddress("0xD67FC48b298B09Ed3D03403d930769C527186c4e"),
		nxtdigest,
	)
	require.Nil(t, err)
	ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(nxtsubdigest), transactions, big.NewInt(123))
	require.False(t, ok)

	// removing a transaction should invalidate the interpretation
	ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions[1:], nonce)
	require.False(t, ok)

	// adding an extra transaction should invalidate the interpretation
	transactions = append(transactions, &sequence.Transaction{
		DelegateCall:  false,
		RevertOnError: true,
		GasLimit:      big.NewInt(21000),
		To:            common.HexToAddress("0x479F6a5b0C1728947318714963a583C56A78366A"),
		Value:         big.NewInt(39381),
		Data:          common.FromHex("0x3251ba32"),
	})

	ok, _ = sendTransactionData.IsValidInterpretation(common.BytesToHash(subdigest), transactions, nonce)
	require.False(t, ok)
}
