package v1v2_test

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/core/v1v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	rpcURLEthereum = "https://nodes.v1v2.app/mainnet"
	rpcURLArbitrum = "https://nodes.v1v2.app/arbitrum"
)

func TestIsValidMessageSignatureEOA(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)

	signature, err := eoa.SignMessage([]byte(message))
	assert.NoError(t, err)
	assert.Len(t, signature, crypto.SignatureLength)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	isValid, err := v1v2.IsValidMessageSignature(eoa.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValidMessageSignatureSequence(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)
	wallet, err := v1v2.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	_, eip191Message := accounts.TextAndHash([]byte(message))

	signature, err := wallet.SignMessage([]byte(eip191Message))
	assert.NoError(t, err)

	signature, err = v1v2.EIP6492Signature(signature, wallet.GetWalletConfig())
	assert.NoError(t, err)

	isValid, err := v1v2.IsValidMessageSignature(wallet.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValidSignatureEIP712Sequence(t *testing.T) {
	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)
	wallet, err := v1v2.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	typedDataJson := `{
		"types": {
			"EIP712Domain": [
				{ "name": "name", "type": "string" },
				{ "name": "version", "type": "string" },
				{ "name": "chainId", "type": "uint256" },
				{ "name": "verifyingContract", "type": "address" },
				{ "name": "salt", "type": "bytes32" }
			],
			"ExampleMessage": [
				{ "name": "message", "type": "string" },
				{ "name": "value", "type": "uint256" },
				{ "name": "from", "type": "address" },
				{ "name": "to", "type": "address" }
			]
		},
		"domain": {
			"name": "EIP712Example",
			"version": "1",
			"chainId": 5,
			"verifyingContract": "0xc0ffee254729296a45a3885639AC7E10F9d54979",
			"salt": "0x70736575646f2d72616e646f6d2076616c756500000000000000000000000000"
		},
		"message": {
			"message": "Test message",
			"value": 10000,
			"from": "0xc0ffee254729296a45a3885639AC7E10F9d54979",
			"to": "0xc0ffee254729296a45a3885639AC7E10F9d54979"
		}
	}`

	typedData, err := ethcoder.TypedDataFromJSON(typedDataJson)
	require.NoError(t, err)

	// domainHash, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	// require.NoError(t, err)
	// require.Equal(t, "0xe073fe030277efdf89d39322ac9321f17774fce4f686e14ff161942bbae5fdcd", ethcoder.HexEncode(domainHash))

	sig, encodedTypedData, err := wallet.SignTypedData(typedData)
	require.NoError(t, err)
	require.NotNil(t, sig)
	require.NotNil(t, encodedTypedData)

	signature, err := v1v2.EIP6492Signature(sig, wallet.GetWalletConfig())
	assert.NoError(t, err)

	isValid, err := v1v2.IsValidTypedDataSignature(wallet.Address(), encodedTypedData, signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValideMessageSignatureSequence_EIP6492SignatureWithMultipleDeployments(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)
	wallet, err := v1v2.NewWalletSingleOwner(eoa)
	require.NoError(t, err)

	walletChild, err := v1v2.NewWalletSingleOwner(wallet)
	require.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	err = walletChild.Connect(provider, nil)
	assert.NoError(t, err)

	_, eip191Message := accounts.TextAndHash([]byte(message))

	signature, err := walletChild.SignMessage([]byte(eip191Message))
	assert.NoError(t, err)

	signature, err = v1v2.EIP6492SignatureWithMultipleDeployments(signature, []core.WalletConfig{wallet.GetWalletConfig(), walletChild.GetWalletConfig()})
	assert.NoError(t, err)

	isValid, err := v1v2.IsValidMessageSignature(walletChild.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

/*
const message = "hi!"

func TestSignatureReduction(t *testing.T) {
	// second and seventh signers are optimal, 3+3 >= 5
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
		},
	})

	// missing seventh signer, optimal is 3+1+1 >= 5
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
		},
	})

	// nested signature is optimal after reduction, only four ecrecovers needed
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 4,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
					},
				},
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeAddress,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 4,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
					},
				},
			},
		},
	})

	// nested signature has higher threshold now and is suboptimal
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 6,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
					},
				},
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Type:   v1v2.SignaturePartTypeAddress,
			},
		},
	})

	// all things being equal, prefer less nesting
	testSignatureReduction(t, &signature{
		Threshold: 2,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 2,
				Signature: &signature{
					Threshold: 2,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   v1v2.SignaturePartTypeEOA,
						},
					},
				},
			},
		},
	}, &signature{
		Threshold: 2,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   v1v2.SignaturePartTypeEOA,
			},
			{
				Weight: 2,
				Type:   v1v2.SignaturePartTypeAddress,
			},
		},
	})
}

func testSignatureReduction(t *testing.T, input *signature, expected *signature) {
	output := toSequenceSignature(input)
	err := output.Reduce([]byte(message))
	assert.NoError(t, err)
	assert.True(t, areSignaturesIsomorphic(fromSequenceSignature(output), expected))
}

type signature struct {
	Threshold uint16
	Signers   []*signaturePart
}

type signaturePart struct {
	Weight    uint8
	Type      uint8
	Signature *signature
}

func toSequenceSignature(s *signature) *v1v2.Signature {
	var signers v1v2.SignatureParts
	for _, signer := range s.Signers {
		part := v1v2.SignaturePart{
			Weight: signer.Weight,
			Type:   signer.Type,
		}

		if signer.Signature != nil {
			part.Type = v1v2.SignaturePartTypeDynamic
			part.Value, _ = toSequenceSignature(signer.Signature).Encode()
			part.Value = append(part.Value, v1v2.SignatureTypeEip1271)
		} else if signer.Type == v1v2.SignaturePartTypeEOA {
			wallet, _ := ethwallet.NewWalletFromRandomEntropy()
			part.Value, _ = wallet.SignMessage([]byte(message))
			part.Value = append(part.Value, v1v2.SignatureTypeEthSign)
		}

		signers = append(signers, &part)
	}

	return &v1v2.Signature{
		Threshold: s.Threshold,
		Signers:   signers,
	}
}

func fromSequenceSignature(s *v1v2.Signature) *signature {
	var signers []*signaturePart
	for _, signer := range s.Signers {
		part := signaturePart{
			Weight: signer.Weight,
			Type:   signer.Type,
		}

		if signer.Type == v1v2.SignaturePartTypeDynamic {
			sig, _ := v1v2.GenericDecodeSignature(signer.Value[:len(signer.Value)-1])
			part.Signature = fromSequenceSignature(sig)
		}

		signers = append(signers, &part)
	}

	return &signature{
		Threshold: s.Threshold,
		Signers:   signers,
	}
}

func areSignaturesIsomorphic(a *signature, b *signature) bool {
	if a.Threshold != b.Threshold {
		return false
	}

	if len(a.Signers) != len(b.Signers) {
		return false
	}

	for i := range a.Signers {
		if a.Signers[i].Weight != b.Signers[i].Weight {
			return false
		}

		if (a.Signers[i].Signature != nil) != (b.Signers[i].Signature != nil) {
			return false
		} else if a.Signers[i].Signature != nil {
			if !areSignaturesIsomorphic(a.Signers[i].Signature, b.Signers[i].Signature) {
				return false
			}
		} else if a.Signers[i].Type != b.Signers[i].Type {
			return false
		}
	}

	return true
}
*/
