package main

import (
	"fmt"
	"math"
	"math/big"
)

func canBeCompressed(number *big.Int) (bool, int, int) {
	var n int

	ten := big.NewInt(10)
	maxByteValue := big.NewInt(255)
	zero := big.NewInt(0)

	for n = 0; n < 256; n++ {
		powerOfTen := new(big.Int).Exp(ten, big.NewInt(int64(n)), nil)

		if powerOfTen.Cmp(number) > 0 {
			break
		}

		quotient := new(big.Int)
		remainder := new(big.Int)
		quotient.DivMod(number, powerOfTen, remainder)

		if remainder.Cmp(zero) == 0 && quotient.Cmp(maxByteValue) <= 0 {
			return true, n, int(quotient.Int64())
		}
	}

	return false, 0, 0
}

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

func minBytesToRepresent(num int) int {
	if num == 0 {
		return 1 // Even 0 requires 1 byte to represent
	}

	var bitsNeeded uint
	for num > 0 {
		bitsNeeded++
		num >>= 1
	}

	// Convert bits to bytes
	return int((bitsNeeded + 7) / 8)
}

func minBytesToRepresent2(num int) int {
	if num == 0 {
		return 1 // Even 0 requires 1 byte to represent
	}
	// Calculate the number of bits needed
	bitsNeeded := math.Floor(math.Log2(float64(num))) + 1
	// Convert bits to bytes and return
	return int(math.Ceil(bitsNeeded / 8))
}

func main() {
	// Create a slice with length 2 and capacity 4
	originalSlice := make([]int, 2, 4)
	originalSlice[0] = 1
	originalSlice[1] = 2

	// Print the original slice and its capacity
	fmt.Println("Original Slice:", originalSlice, "Capacity:", cap(originalSlice))

	// Append an element to the slice
	newSlice := append(originalSlice, 3)

	// Print the new slice and its capacity
	fmt.Println("New Slice:", newSlice, "Capacity:", cap(newSlice))

	// Demonstrating that both slices share the same underlying array
	newSlice[0] = 100
	fmt.Println("After modification - Original Slice:", originalSlice)
	fmt.Println("After modification - New Slice:", newSlice)
}
