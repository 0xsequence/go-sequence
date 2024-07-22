package sequence_test

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"
	"github.com/0xsequence/go-sequence/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	rpcURLEthereum = "https://nodes.sequence.app/mainnet"
	rpcURLArbitrum = "https://nodes.sequence.app/arbitrum"
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

	isValid, err := sequence.IsValidMessageSignature(eoa.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValidMessageSignatureSequence(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	assert.NoError(t, err)
	wallet, err := sequence.NewWalletSingleOwner(eoa)
	assert.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	_, eip191Message := accounts.TextAndHash([]byte(message))

	signature, err := wallet.SignMessage([]byte(eip191Message))
	assert.NoError(t, err)

	signature, err = sequence.EIP6492Signature(signature, wallet.GetWalletConfig())
	assert.NoError(t, err)

	isValid, err := sequence.IsValidMessageSignature(wallet.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
	assert.NoError(t, err)
	assert.True(t, isValid)
}

func TestIsValideMessageSignatureSequence_EIP6492SignatureWithMultipleDeployments(t *testing.T) {
	message := "hello world!"

	eoa, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)
	wallet, err := sequence.NewWalletSingleOwner(eoa)
	require.NoError(t, err)

	walletChild, err := sequence.NewWalletSingleOwner(wallet)
	require.NoError(t, err)

	provider, err := ethrpc.NewProvider(rpcURLEthereum)
	assert.NoError(t, err)

	err = wallet.Connect(provider, nil)
	assert.NoError(t, err)

	_, eip191Message := accounts.TextAndHash([]byte(message))

	signature, err := wallet.SignMessage([]byte(eip191Message))
	assert.NoError(t, err)

	signature, err = sequence.EIP6492SignatureWithMultipleDeployments(signature, []core.WalletConfig{wallet.GetWalletConfig(), walletChild.GetWalletConfig()})
	assert.NoError(t, err)

	isValid, err := sequence.IsValidMessageSignature(wallet.Address(), []byte(message), signature, big.NewInt(1), provider, nil)
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
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
		},
	})

	// missing seventh signer, optimal is 3+1+1 >= 5
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
		},
	}, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 3,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
		},
	})

	// nested signature is optimal after reduction, only four ecrecovers needed
	testSignatureReduction(t, &signature{
		Threshold: 5,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 4,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
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
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeAddress,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 4,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeAddress,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
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
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Signature: &signature{
					Threshold: 6,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
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
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 5,
				Type:   sequence.SignaturePartTypeAddress,
			},
		},
	})

	// all things being equal, prefer less nesting
	testSignatureReduction(t, &signature{
		Threshold: 2,
		Signers: []*signaturePart{
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 2,
				Signature: &signature{
					Threshold: 2,
					Signers: []*signaturePart{
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
						},
						{
							Weight: 1,
							Type:   sequence.SignaturePartTypeEOA,
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
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 1,
				Type:   sequence.SignaturePartTypeEOA,
			},
			{
				Weight: 2,
				Type:   sequence.SignaturePartTypeAddress,
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

func toSequenceSignature(s *signature) *sequence.Signature {
	var signers sequence.SignatureParts
	for _, signer := range s.Signers {
		part := sequence.SignaturePart{
			Weight: signer.Weight,
			Type:   signer.Type,
		}

		if signer.Signature != nil {
			part.Type = sequence.SignaturePartTypeDynamic
			part.Value, _ = toSequenceSignature(signer.Signature).Encode()
			part.Value = append(part.Value, sequence.SignatureTypeEip1271)
		} else if signer.Type == sequence.SignaturePartTypeEOA {
			wallet, _ := ethwallet.NewWalletFromRandomEntropy()
			part.Value, _ = wallet.SignMessage([]byte(message))
			part.Value = append(part.Value, sequence.SignatureTypeEthSign)
		}

		signers = append(signers, &part)
	}

	return &sequence.Signature{
		Threshold: s.Threshold,
		Signers:   signers,
	}
}

func fromSequenceSignature(s *sequence.Signature) *signature {
	var signers []*signaturePart
	for _, signer := range s.Signers {
		part := signaturePart{
			Weight: signer.Weight,
			Type:   signer.Type,
		}

		if signer.Type == sequence.SignaturePartTypeDynamic {
			sig, _ := sequence.GenericDecodeSignature(signer.Value[:len(signer.Value)-1])
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
