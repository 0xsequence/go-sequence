package compressor

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
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

func pow10factor(b []byte) (int, int) {
	num := big.NewInt(0).SetBytes(b)

	var n int

	ten := big.NewInt(10)
	maxByteValue := big.NewInt(255)
	zero := big.NewInt(0)

	for n = 0; n < 256; n++ {
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

func minBytesToRepresent(num int) int {
	if num == 0 {
		return 1
	}

	var bitsNeeded uint
	for num > 0 {
		bitsNeeded++
		num >>= 1
	}

	// Convert bits to bytes
	return int((bitsNeeded + 7) / 8)
}

func intToBytes(n int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, n)

	if err != nil {
		log.Fatalf("binary.Write failed: %v", err)
	}

	return buf.Bytes()
}
