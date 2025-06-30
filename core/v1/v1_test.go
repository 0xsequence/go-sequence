package v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var signatures = []string{
	"0x0005010200a17865a33ed1fc9285b67d3b73e273b92bd3f1010206b54d1321bc871d74a215791b028a76123f9a6101031112240f18e5e8ea52c47450ee8ad879c05291680102131e002341ca11920a17fde6a6835490b7c651b7010215e5e938279efe474ab0ad814c01cc30b7841fda010220ce19ce77a9982ffe075614c14e09d1c683698d010221ac797918b75473d132d6c0421b08e02bee2f0e0102244093813960805e9787260e7d8dbd24c9f35f7c010227a380baf51044343efa393c69e1dbac9822a04f010229bfea22d2e0c8aa17af83ee655b8faef26f68d7010229d71969d8563147ba9974f433485a742f19498001022ad9786132fe6cdf3ff2ee26c46113a5b5c6e87f010233ac06c6d3516253db712625f25ec9c202da55e2010234e9aee8ef165addffa9b37cfd11bc026e92e7f3010238aecc0b56cab4968a19475a1eb523a18d07bf0801023b732ad9fa4e96ef217a78c5ebaccf625d95689601024be2a9f87433a85c6c63c86ed451768d5e89a6e601024c4cb775073763e9d8b9051ce5bb71965a13757d01024c59aafa0bb9558287f76a1dbb52cc0638107d5901024c6a8e473830aed93dbefc605073cbd6edccb65201024fe546d5d9498af8b6d9d4a84a116347a4e54467010256b170097c80f25f987ce72beb233d72ca6fa418010256f6d684ddded3cc2d33ba498580f1e088d7aedc010257b2a2a2fb419f3f8eafd1e22afc1cea1e53436c0203596af90cecdbf9a768886e771178fd5561dd27ab005d000100015863eef3815df42cce090e7c062a045d1124758c0cd62846542aa2ba6e20aac74c56f93ef273f26272c778a79fb46279e4dd3099878b8143716d035eb3951e231c020101c50adeadb7fe15bee45dcb820610cdedcd314eb003010264f3e857c3046f8f1f88b7eb0ebbd6da69ed745401026ac2c330f5be7fc9d46b91aea76cea31d74b2243010275da5b8d534cd97fb2a1d4b6dd3b36459b6e7739010279413baa9ba9543091bc1dc03849be540e6334bc01027d205afe563d7c594b852eb2ce336552fd786ce4010284608758c3ee874a67e68e4a5384108198ccf2ea010285b15d992f00d21b03250e627da716ae18964c5c01028f4b5f0f89c627b0361c6dd781a04eaa00b8230d01028fe27c58e4c579c650e6f6a3bf4c7160eb3567b501029dc63ce95131ed3a6de0d04129573828de0744c20102a0105042841f10c8031debebeae4cd5b87a753fd0102aadc029069dfb9ea33003facdba3bf7d3c2123000102ac99de33845b9b5ae2c9af03cfbf2e758576e7120102ae00aed3b19cb743860b22b3fc3d10e6c107706b0102ae1f61f1a946dd50fedf330b225d3d77911cd4170102b21b24335940ce804f0d00bd2b777a34fb4365db0002d08650c4a9b751f15158af9a0af9d08c4c680125ddb809ea6214991d07fa702923034df116d2111d981411c8f14ac0a1a091cbf4fbff59b583734ac267a51a141b020102b5ac88b4f3ec999ab281ca91d3dfb13cbf558ef10102bfa91a56381159308f49b274bcf49b6065f3d0aa0102c29b254e88aff672e2b4bb59af015db138b8ddfa0102c527cd197cc397a049599ad7ca0ab568368631c00102c714e9c14dec5d7bbc91370ee5706e0d5fdc91d80102c85208375c5760b7e8336d2db9e85e0486e1b2560102c912dd7282dd3b3bb6821e49d365c44118140c910102db845c8e084a313f0f0e3dfcaefd95bc547cd9920102dd388b197311c0db798d399dc61a9f8e504d7d560102dee82b6d76f4c1f75dd6db933a789dc02813a9130102e2e1eadd2eb7dcde83c5b3cd95f1d9ab0925c7480102e376498fe3590c20540e4c9c9846aa78aacdb6e80102e599e408c776c7284f0d201b8366405497ad3fa20102e5c3c58ee6478df2dfdb99b5b3840242d5b193150102f007e656dacf90f4920a2c9806ffe421d2f5587d0102f2fa090bd240450e909f98ba4168052f81059c18",
}

func TestDecodeSignature(t *testing.T) {
	for i, signature := range signatures {
		decodedSignature, err := Core.DecodeSignature(hexutil.MustDecode(signature))
		assert.NoErrorf(t, err, "unable to decode signature %v", i)

		spew.Dump(decodedSignature)
		fmt.Println()

		reEncodedSignature, err := decodedSignature.Data()
		assert.NoErrorf(t, err, "unable to re-encode signature %v", i)

		fmt.Printf("original signature:   %v\n", signature)
		fmt.Printf("re-encoded signature: %v\n", hexutil.Encode(reEncodedSignature))
		fmt.Println()

		assert.Equalf(t, signature, hexutil.Encode(reEncodedSignature), "re-encoded signature does not match original signature")
	}
}

func TestSignatureJoin(t *testing.T) {
	eoa1, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	eoa2, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	eoa3, err := ethwallet.NewWalletFromRandomEntropy()
	require.NoError(t, err)

	wc := &WalletConfig{
		Threshold_: 2,
		Signers_: WalletConfigSigners{
			&WalletConfigSigner{
				Weight:  1,
				Address: eoa1.Address(),
			},
			&WalletConfigSigner{
				Weight:  1,
				Address: eoa2.Address(),
			},
			&WalletConfigSigner{
				Weight:  1,
				Address: eoa3.Address(),
			},
		},
	}

	msg := []byte("hello")
	payload := Digest(crypto.Keccak256Hash(msg), common.Address{})

	sig1, err := wc.BuildSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		if signer == eoa1.Address() {
			sig, _ := eoa1.SignMessage(payload.Digest().Bytes())
			return core.SignerSignatureTypeEthSign, sig, nil
		} else {
			return 0, nil, nil
		}
	}, false)
	require.NoError(t, err)

	sig2, err := wc.BuildSignature(context.Background(), func(ctx context.Context, signer common.Address, signatures []core.SignerSignature) (core.SignerSignatureType, []byte, error) {
		if signer == eoa2.Address() {
			sig, _ := eoa2.SignMessage(payload.Digest().Bytes())
			return core.SignerSignatureTypeEthSign, sig, nil
		} else {
			return 0, nil, nil
		}
	}, false)
	require.NoError(t, err)

	sig1Leaves := sig1.(*signature).leaves
	sig2Leaves := sig2.(*signature).leaves

	sigJoined, err := sig1.Join(payload, sig2)
	require.NoError(t, err)

	sigJoinedLeaves := sigJoined.(*signature).leaves

	require.Equal(t, 3, len(sig1Leaves))
	require.Equal(t, 3, len(sig2Leaves))
	require.Equal(t, 3, len(sigJoinedLeaves))

	assert.Equal(t, sig1Leaves[0], sigJoinedLeaves[0])
	assert.Equal(t, sig2Leaves[1], sigJoinedLeaves[1])
	assert.Equal(t, sig1Leaves[2], sigJoinedLeaves[2])
	assert.Equal(t, sig2Leaves[2], sigJoinedLeaves[2])
}

const configTOML = `
	threshold = 5

	[[signers]]
	weight = 3
	address = "0x1111111111111111111111111111111111111111"

	[[signers]]
	weight = 3
	address = "0x2222222222222222222222222222222222222222"

	[[signers]]
	weight = 2
	address = "0x3333333333333333333333333333333333333333"
`

func TestWalletConfigTOML(t *testing.T) {
	var config map[string]any
	_, err := toml.Decode(configTOML, &config)
	assert.NoError(t, err)

	config_, err := Core.DecodeWalletConfig(config)
	assert.NoError(t, err)

	spew.Dump(config_)

	var config__ WalletConfig
	_, err = toml.Decode(configTOML, &config__)
	assert.NoError(t, err)

	spew.Dump(config__)
}
