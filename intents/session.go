package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type Session interface {
	SessionId() string
	Sign(intent *Intent) error
}

type session256K1 struct {
	wallet *ethwallet.Wallet
}

func NewSessionP256K1(wallet *ethwallet.Wallet) Session {
	return &session256K1{wallet: wallet}
}

func (s session256K1) SessionId() string {
	return strings.ToLower(
		fmt.Sprintf(
			"0x%s",
			common.Bytes2Hex(
				append([]byte{byte(KeyTypeSECP256K1)}, s.wallet.Address().Bytes()...),
			),
		),
	)
}

func (s session256K1) Sign(intent *Intent) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	signature, err := s.wallet.SignMessage(hash)
	if err != nil {
		return err
	}

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionId: s.SessionId(),
		Signature: bytesToSignature(signature),
	})
	return nil
}

type session256R1 struct {
	privateKey *ecdsa.PrivateKey
}

func NewSessionP256R1(privateKey *ecdsa.PrivateKey) Session {
	return &session256R1{privateKey: privateKey}
}

func (s session256R1) SessionId() string {
	pubKey := elliptic.Marshal(s.privateKey.Curve, s.privateKey.PublicKey.X, s.privateKey.PublicKey.Y)
	return strings.ToLower(
		fmt.Sprintf(
			"0x%s",
			common.Bytes2Hex(
				append([]byte{byte(KeyTypeSECP256R1)}, pubKey...),
			),
		),
	)
}

func (s session256R1) Sign(intent *Intent) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	sha256Hash := sha256.Sum256(hash)

	sr, ss, err := ecdsa.Sign(rand.Reader, s.privateKey, sha256Hash[:])
	if err != nil {
		return err
	}

	signature := append(sr.Bytes(), ss.Bytes()...)

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionId: s.SessionId(),
		Signature: bytesToSignature(signature),
	})
	return nil
}

func SignIntentWithWalletLegacy[T any](wallet *ethwallet.Wallet, intent *IntentTyped[T]) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	signature, err := wallet.SignMessage(hash)
	if err != nil {
		return err
	}

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionId: strings.ToLower(wallet.Address().String()),
		Signature: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(signature))),
	})
	return nil
}

func bytesToSignature(sig []byte) string {
	return strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(sig)))
}
