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

func SignIntentWithWallet[T any](wallet *ethwallet.Wallet, intent *IntentTyped[T]) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	signature, err := wallet.SignMessage(hash)
	if err != nil {
		return err
	}

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionId: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(append([]byte{byte(KeyTypeSECP256K1)}, wallet.Address().Bytes()...)))),
		Signature: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(signature))),
	})
	return nil
}

func SignIntentWithP256K1[T any](wallet *ethwallet.Wallet, intent *IntentTyped[T]) error {
	return SignIntentWithWallet(wallet, intent)
}

func SignIntentWithP256R1[T any](privateKey *ecdsa.PrivateKey, intent *IntentTyped[T]) error {
	hash, err := intent.Hash()
	if err != nil {
		return err
	}

	sha256Hash := sha256.Sum256(hash)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, sha256Hash[:])
	if err != nil {
		return err
	}

	signature := append(r.Bytes(), s.Bytes()...)

	pubKey := elliptic.Marshal(privateKey.Curve, privateKey.PublicKey.X, privateKey.PublicKey.Y)

	intent.Signatures = append(intent.Signatures, &Signature{
		SessionId: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(append([]byte{byte(KeyTypeSECP256R1)}, pubKey...)))),
		Signature: strings.ToLower(fmt.Sprintf("0x%s", common.Bytes2Hex(signature))),
	})
	return nil
}
