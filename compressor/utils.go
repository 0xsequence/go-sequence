package compressor

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	"github.com/0xsequence/go-sequence"
)

func ispow2minus1(b []byte) int {
	seen1 := false
	val := 1

	for _, byteVal := range b {
		for ii := 7; ii >= 0; ii-- {
			bit := (byteVal >> ii) & 1
			if bit == 1 {
				if !seen1 {
					seen1 = true
					continue
				}
				val += 1
			} else {
				if seen1 {
					return -1
				}
			}
		}
	}

	return val
}

func ispow2(b []byte) int {
	seen1 := false
	val := 0

	for _, byteVal := range b {
		for ii := 7; ii >= 0; ii-- {
			bit := (byteVal >> ii) & 1
			if bit == 1 {
				if !seen1 {
					seen1 = true
					continue
				} else {
					return -1
				}
			} else {
				if seen1 {
					val++
				}
			}
		}
	}

	return val
}

func pow10(b []byte) int {
	num := big.NewInt(0).SetBytes(b)

	ten := big.NewInt(10)
	zero := big.NewInt(0)
	count := 0

	for num.Cmp(ten) >= 0 {
		remainder := big.NewInt(0)
		remainder.Mod(num, ten)

		if remainder.Cmp(zero) != 0 {
			return -1
		}

		num.Div(num, ten)
		count++
	}

	if num.Cmp(big.NewInt(1)) != 0 {
		return -1
	}

	return count
}

func pow10Mantissa(b []byte, maxExp int, maxMantissa int) (int, int) {
	num := big.NewInt(0).SetBytes(b)

	var n int

	ten := big.NewInt(10)
	maxByteValue := big.NewInt(int64(maxMantissa))
	zero := big.NewInt(0)

	for n = 0; n < maxExp; n++ {
		powerOfTen := new(big.Int).Exp(ten, big.NewInt(int64(n)), nil)

		if powerOfTen.Cmp(num) > 0 {
			break
		}

		quotient := new(big.Int)
		remainder := new(big.Int)
		quotient.DivMod(num, powerOfTen, remainder)

		if remainder.Cmp(zero) == 0 && quotient.Cmp(maxByteValue) <= 0 {
			return n, int(quotient.Int64())
		}
	}

	return -1, -1
}

func minBytesToRepresent(num uint) uint {
	if num == 0 {
		return 1
	}

	var bitsNeeded uint
	for num > 0 {
		bitsNeeded++
		num >>= 1
	}

	// Convert bits to bytes
	return uint((bitsNeeded + 7) / 8)
}

func uintToBytes(n uint64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, n)

	if err != nil {
		log.Fatalf("binary.Write failed: %v", err)
	}

	return buf.Bytes()
}

func padToX(b []byte, n uint) ([]byte, error) {
	if len(b) > int(n) {
		// Try trimming leading zeros, if that doesn't work, return an error
		b = bytes.TrimLeft(b, "\x00")
		if len(b) > int(n) {
			return nil, fmt.Errorf("value too large to pad")
		}
	}

	if len(b) == int(n) {
		return b, nil
	}

	pad := make([]byte, int(n)-len(b))
	return append(pad, b...), nil
}

func uintPadToX(u uint, n uint) ([]byte, error) {
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, uint64(u))
	return padToX(buff, n)
}

func bytesToUint64(b []byte) uint {
	// Read byte by byte, as it may be shorter
	// this is big endian
	var val uint
	for _, byteVal := range b {
		val <<= 8
		val |= uint(byteVal)
	}
	return val
}

func FindPastData(pastData *CBuffer, data []byte) int {
	for i := 0; i+len(data) < len(pastData.Data()); i++ {
		if bytes.Equal(pastData.Data()[i:i+len(data)], data) {
			return i
		}
	}

	return -1
}

func NormalizeSignature(signature []byte) ([]byte, error) {
	if len(signature) == 0 {
		return signature, nil
	}
	// One thing that happens is that there are two ways of representing a Sequence signature in v2
	// lagacy and dynamic, decompressor will ALWAYS decompress to the dynamic format, so if
	// the input is in the legacy format, we need to convert it to the dynamic format.
	// This is easy, because the legacy format starts with 0x00, if that's the case we just
	// add 0x01 at the beginning of the signature
	// Notice that guest executes have no signature, so we don't need to fix those
	if signature[0] == 0 {
		return append([]byte{1}, signature...), nil
	}

	// Chained signatures are a bit harder, as they need to be dissected and normalized
	// then-reassembled again
	if signature[0] == 3 {
		normalized := []byte{0x03}
		var rindex uint = 1

		for rindex < uint(len(signature)) {
			// The first 3 bytes are the length of the part
			length := bytesToUint64(signature[rindex : rindex+3])
			rindex += 3

			part := signature[rindex : rindex+uint(length)]
			rindex += uint(length)

			npart, err := NormalizeSignature(part)
			if err != nil {
				return nil, err
			}

			// Notice that we need to write it with the new length
			plen, err := uintPadToX(uint(len(npart)), 3)
			if err != nil {
				return nil, err
			}

			normalized = append(normalized, plen...)
			normalized = append(normalized, npart...)
		}

		return normalized, nil
	}

	return signature, nil
}

func NormalizeTransactionSignature(
	transaction *sequence.Transaction,
) error {
	sig, err := NormalizeSignature(transaction.Signature)
	if err != nil {
		return err
	}

	transaction.Signature = sig

	// If the transaction is a nested sequence transaction, we need to normalize its calldata too
	for _, tx := range transaction.Transactions {
		if tx.Signature != nil && len(tx.Signature) != 0 {
			err := NormalizeTransactionSignature(tx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
