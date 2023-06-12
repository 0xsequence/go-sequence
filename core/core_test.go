package core

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSigningOrchestrator(t *testing.T) {
	signers := map[common.Address]uint16{
		common.HexToAddress("0x00"): 1,
		common.HexToAddress("0x01"): 1,
		common.HexToAddress("0x02"): 1,
	}

	mockSignatures := map[common.Address]struct {
		sigType SignerSignatureType
		sig     []byte
	}{
		common.HexToAddress("0x00"): {
			sigType: SignerSignatureTypeEthSign,
			sig:     []byte("signature"),
		},
		common.HexToAddress("0x01"): {
			sigType: SignerSignatureTypeEthSign,
			sig:     []byte("signature2"),
		},
		common.HexToAddress("0x02"): {
			sigType: SignerSignatureTypeEIP1271,
			sig:     []byte("signature3"),
		},
	}

	signer0Retried := false
	signingFunction := func(ctx context.Context, signer common.Address, signatures []SignerSignature) (SignerSignatureType, []byte, error) {
		if signer == common.HexToAddress("0x00") && !signer0Retried {
			signer0Retried = true
			return 0, nil, ErrSigningFunctionNotReady
		}

		return mockSignatures[signer].sigType, mockSignatures[signer].sig, nil
	}

	signaturesChan := SigningOrchestrator(context.Background(), signers, signingFunction)

	signatures := make([]SignerSignature, 0, len(signers))
	for i := 0; i < len(signers); i++ {
		signature := <-signaturesChan
		signatures = append(signatures, signature)
	}

	assert.True(t, signer0Retried)
	for _, signature := range signatures {
		assert.Equal(t, mockSignatures[signature.Signer].sigType, signature.Type)
		assert.Equal(t, mockSignatures[signature.Signer].sig, signature.Signature)
	}
}
