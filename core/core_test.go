package core_test

import (
	"context"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
	"github.com/stretchr/testify/assert"
)

func TestSigningOrchestrator(t *testing.T) {
	signers := map[core.Signer]uint16{
		{Address: common.HexToAddress("0x00")}: 1,
		{Address: common.HexToAddress("0x01")}: 1,
		{Address: common.HexToAddress("0x02")}: 1,
	}

	mockSignatures := map[core.Signer]struct {
		sigType core.SignerSignatureType
		sig     []byte
	}{
		{Address: common.HexToAddress("0x00")}: {
			sigType: core.SignerSignatureTypeEthSign,
			sig:     []byte("signature"),
		},
		{Address: common.HexToAddress("0x01")}: {
			sigType: core.SignerSignatureTypeEthSign,
			sig:     []byte("signature2"),
		},
		{Address: common.HexToAddress("0x02")}: {
			sigType: core.SignerSignatureTypeEIP1271,
			sig:     []byte("signature3"),
		},
	}

	signer0Retried := false
	signingFunction := func(ctx context.Context, signer core.Signer, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		if signer == (core.Signer{Address: common.HexToAddress("0x00")}) && !signer0Retried {
			signer0Retried = true
			return 0, nil, core.ErrSigningFunctionNotReady
		}

		return mockSignatures[signer].sigType, mockSignatures[signer].sig, nil
	}

	signaturesChan := core.SigningOrchestrator(context.Background(), signers, signingFunction)

	signatures := make([]core.SignerSignature, 0, len(signers))
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
