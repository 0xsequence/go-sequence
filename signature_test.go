package sequence_test

import (
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence"
	"github.com/stretchr/testify/assert"
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

func TestIsValidMessageSignatureUbisoft(t *testing.T) {
	provider, err := ethrpc.NewProvider(rpcURLArbitrum)
	assert.NoError(t, err)
	assert.NotNil(t, provider)

	isValid, err := sequence.IsValidMessageSignature(
		common.HexToAddress("0xB9D79174FA54a3161482102d167205b60564329f"),
		[]byte("edenonline.ubisoft.com wants you to sign in with your Ethereum account:\n0xB9D79174FA54a3161482102d167205b60564329f\n\nAssociate my wallet to my Ubisoft account.\n\nURI: https://edenonline.ubisoft.com\nVersion: 1\nChain ID: 42161\nNonce: 2Cd3JLESd6k2i27r7MXPUe\nIssued At: 2024-05-28T14:42:11.974Z"),
		hexutil.MustDecode("0x000000000000000000000000faa5c0b14d1bed5c888ca655b9a8a5911f78ef4a000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000004432c02a14000000000000000000000000fbf8f1a5e00034762d928f46d438b947f5d4065dc9ad3b83d2673efd9f3ddfd59cf2b9d6e76338dd0c92f4964bea3872697e381d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001340100010000000002015ee37038bd4368b8fb13756429c1f69120be7f510000fe01000300000000060100010000740001b3f98519d4dd3e57119c8d696d3c37df4ac07d1ef73ca0b0bd3aa9b76e60179b785b93bc9d4766328649aab791e9e62ecc0156a3e5b05dfd0aaddb8896c818901c010400002c01014d6d738c7e26faf0bfdbe12601fe9f798aa8816201010b87dfa13d00ab4f335dd83fa21529240ef556ed060200010000740001999ff5dfc3f57df741b2e7f552f58e5dac73fc4783c5ba69ce8591ce415eb17815fdad7be8aec9a0c99ebe1e37b932b1ce23dac3295b054dc68f46f3526b3ea21b010400002c0101951448847a03ad1005a0e463dff0da093690ff2401016dd76dd81c47c5d20605ec72bf4c2257cabcad0b0301005e8ddc2afade4d69c4329345653f27ca25e62c7c0000000000000000000000006492649264926492649264926492649264926492649264926492649264926492"),
		big.NewInt(42161),
		provider,
		nil,
	)
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
