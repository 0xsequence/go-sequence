package relayer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
	"sort"
	"time"

	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/common"
)

type FeeQuote struct {
	TransactionDigest common.Hash
	IsWhitelisted     bool
	GasSponsor        *uint64
	GasTank           *uint64
	GasUsage          *uint64
	GasPrice          *big.Int
	NativePrice       *big.Int
	TokenPrices       map[string]*big.Int
	ExpiresAt         *time.Time
	Signature         []byte
}

func (q *FeeQuote) Sign(wallet *ethwallet.Wallet, validFor time.Duration) error {
	expiresAt := time.Now().Add(validFor)
	q.ExpiresAt = &expiresAt

	message, err := q.message()
	if err != nil {
		return err
	}

	q.Signature, err = wallet.SignMessage(message)
	return err
}

func (q *FeeQuote) Verify(wallet *ethwallet.Wallet) error {
	return q.VerifySignedBy(wallet.Address())
}

// VerifySignedBy verifies the quote against one or more expected signer addresses
func (q *FeeQuote) VerifySignedBy(expectedAddresses ...common.Address) error {
	if q.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("quote expired")
	}

	message, err := q.message()
	if err != nil {
		return err
	}

	address, err := ethwallet.RecoverAddress(message, q.Signature)
	if err != nil {
		return err
	}

	for _, expected := range expectedAddresses {
		if address == expected {
			return nil
		}
	}

	return fmt.Errorf("quote signed by unexpected address %s", address.Hex())
}

func (q *FeeQuote) message() ([]byte, error) {
	if q.ExpiresAt == nil {
		return nil, fmt.Errorf("missing expiration date")
	}

	var message bytes.Buffer

	_, err := message.Write(q.TransactionDigest.Bytes())
	if err != nil {
		return nil, err
	}

	err = binary.Write(&message, binary.LittleEndian, q.IsWhitelisted)
	if err != nil {
		return nil, err
	}

	err = writeOptionalUint64(&message, q.GasSponsor)
	if err != nil {
		return nil, err
	}

	err = writeOptionalUint64(&message, q.GasTank)
	if err != nil {
		return nil, err
	}

	err = writeOptionalUint64(&message, q.GasUsage)
	if err != nil {
		return nil, err
	}

	err = writeOptionalBigInt(&message, q.GasPrice)
	if err != nil {
		return nil, err
	}

	err = writeOptionalBigInt(&message, q.NativePrice)
	if err != nil {
		return nil, err
	}

	tokens := make([]string, 0, len(q.TokenPrices))
	for token := range q.TokenPrices {
		tokens = append(tokens, token)
	}
	sort.Strings(tokens)
	for _, token := range tokens {
		_, err = message.WriteString(token)
		if err != nil {
			return nil, err
		}

		err = writeBigInt(&message, q.TokenPrices[token])
		if err != nil {
			return nil, err
		}
	}

	err = binary.Write(&message, binary.LittleEndian, q.ExpiresAt.Unix())
	if err != nil {
		return nil, err
	}

	return message.Bytes(), nil
}

func writeOptionalUint64(w io.Writer, n *uint64) error {
	err := binary.Write(w, binary.LittleEndian, n != nil)
	if err != nil {
		return err
	}

	if n != nil {
		err = binary.Write(w, binary.LittleEndian, *n)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeOptionalBigInt(w io.Writer, n *big.Int) error {
	err := binary.Write(w, binary.LittleEndian, n != nil)
	if err != nil {
		return err
	}

	if n != nil {
		err = writeBigInt(w, n)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeBigInt(w io.Writer, n *big.Int) error {
	err := binary.Write(w, binary.LittleEndian, int8(n.Sign()))
	if err != nil {
		return err
	}

	_, err = w.Write(n.Bytes())
	if err != nil {
		return err
	}

	return nil
}
