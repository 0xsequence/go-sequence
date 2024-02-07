package compressor

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/0xsequence/go-sequence"
)

type Compressor struct {
	AddressIndexes map[string]uint
	Bytes32Indexes map[string]uint
	Bytes4Indexes  map[string]uint
}

type EncodeType int

const (
	Stateless EncodeType = iota
	Mirror
	ReadStorage
	WriteStorage
)

func NewCompressor() *Compressor {
	return &Compressor{
		Bytes4Indexes:  LoadBytes4(),
		AddressIndexes: make(map[string]uint),
		Bytes32Indexes: make(map[string]uint),
	}
}

func HighestPriority(a EncodeType, b EncodeType) EncodeType {
	if a == WriteStorage || b == WriteStorage {
		return WriteStorage
	}

	if a == ReadStorage || b == ReadStorage {
		return ReadStorage
	}

	if a == Mirror || b == Mirror {
		return Mirror
	}

	return Stateless
}

// Encodes a 32 bytes word, trying to optimize it as much as possible
func (c *Compressor) EncodeWordOptimized(pastData *CBuffer, word []byte, saveWord bool) ([]byte, EncodeType, error) {
	if len(word) > 32 {
		return nil, Stateless, fmt.Errorf("word exceeds 32 bytes")
	}

	trimmed := bytes.TrimLeft(word, "\x00")

	// If empty then it can be encoded as literal zero
	if len(trimmed) == 0 {
		return []byte{byte(LITERAL_ZERO)}, Stateless, nil
	}

	// Literals are the cheapest encoding
	if len(trimmed) <= 1 && trimmed[0] <= byte(MAX_LITERAL) {
		return []byte{trimmed[0] + byte(LITERAL_ZERO)}, Stateless, nil
	}

	// If it only has 1 byte, then we encode it as a word
	// all other methods use 2 bytes anyway
	if len(trimmed) == 1 {
		return c.EncodeWordBytes32(trimmed)
	}

	// If the word is a power of 2 or 10, we can encode it using 1 byte
	pow2 := ispow2(trimmed)
	if pow2 != -1 {
		return []byte{byte(FLAG_READ_POWER_OF_2), byte(pow2)}, Stateless, nil
	}

	// Notice that marking the first bit of N as 1 denotes that we are going to do
	// 10 ** N and not 10 ** N * X, that's why we do `| 0x80`
	pow10 := pow10(trimmed)
	if pow10 != -1 && pow10 <= 77 {
		return []byte{byte(FLAG_READ_POWER_OF_10_MISC), byte(pow10 | 0x80)}, Stateless, nil
	}

	// 2 ** n - 1 can be represented by two bytes (and the first two are 00)
	// so it goes next. We do it before word encoding since word encoding won't
	// have the first 2 bytes as 00, in reality this only applies to 0xffff
	pow2minus1 := ispow2minus1(trimmed)
	if pow2minus1 != -1 {
		// The opcode adds an extra 1 to the value, so we need to subtract 1
		return []byte{byte(FLAG_READ_POWER_OF_2), 0x00, byte(pow2minus1 - 1)}, Stateless, nil
	}

	// Now we can store words of 2 bytes, we have exhausted all the 1 byte options
	if len(trimmed) <= 2 {
		return c.EncodeWordBytes32(trimmed)
	}

	// We can also use (10 ** N) * X, this uses 2 bytes
	pow10fn, pow10fm := pow10Mantissa(trimmed, 127, 255)
	if pow10fn != -1 && pow10fm != -1 {
		return []byte{byte(FLAG_READ_POWER_OF_10_MISC), byte(pow10fn), byte(pow10fm)}, Stateless, nil
	}

	// Mirror flag uses 2 bytes, it lets us point to another flag that we had already used before
	// but we need to find a flag that mirrors the data with the padding included!
	padded32 := make([]byte, 32)
	copy(padded32[32-len(trimmed):], trimmed)
	padded32str := string(padded32)

	usedFlag := pastData.Refs.usedFlags[padded32str]
	if usedFlag != 0 {
		usedFlag -= 1

		// We can only encode 16 bits for the mirror flag
		// if it exceeds this value, then we can't mirror it
		// NOR it can be this pointer itself
		if usedFlag <= 0xffff && usedFlag != pastData.Len() {
			return []byte{byte(FLAG_MIRROR_FLAG), byte(usedFlag >> 8), byte(usedFlag)}, Mirror, nil
		}
	}

	// Mirror storage flags are different because we don't want to just
	// re-read the storage flag, that would mean writting to the storage twice
	// apart from that, they work like normal mirror flags
	usedStorageFlag := pastData.Refs.usedStorageFlags[padded32str]
	if usedStorageFlag != 0 {
		usedStorageFlag -= 1

		// We can only encode 16 bits for the mirror flag
		// if it exceeds this value, then we can't mirror it
		if usedStorageFlag <= 0xffff {
			return []byte{byte(FLAG_READ_STORE_FLAG), byte(usedStorageFlag >> 8), byte(usedStorageFlag)}, Mirror, nil
		}
	}

	// Now any 3 byte word can be encoded as-is, all the other
	// methods use more than 3 bytes
	if len(trimmed) <= 3 {
		return c.EncodeWordBytes32(trimmed)
	}

	// With 3 bytes we can encode 10 ** N * X (with a mantissa of 18 bits and an exp of 6 bits)
	pow10fn, pow10fm = pow10Mantissa(trimmed, 63, 262143)
	if pow10fn != -1 && pow10fm != -1 {
		// The first byte is 6 bits of exp and 2 bits of mantissa
		b1 := byte(uint64(pow10fn)<<2) | byte(uint64(pow10fm)>>16)
		// The next 16 bits are the last 16 bits of the mantissa
		b2 := byte(pow10fm >> 8)
		b3 := byte(pow10fm)

		return []byte{byte(FLAG_READ_POW_10_MANTISSA), byte(b1), byte(b2), byte(b3)}, Stateless, nil
	}

	// With 3 bytes we can also copy any other word from the calldata
	// this can be anything but notice: we must copy the value already padded
	copyIndex := FindPastData(pastData, padded32)
	if copyIndex != -1 && copyIndex <= 0xffff {
		return []byte{byte(FLAG_COPY_CALLDATA), byte(copyIndex >> 8), byte(copyIndex), byte(0x20)}, Stateless, nil
	}

	// Contract storage is only enabled on some networks
	// in most of L1s is cheaper to provide the data from calldata
	// rather than reading it from storage, let alone writing it
	if pastData.Refs.useContractStorage {
		// If the data is already on storage, we can look it up
		// on the addresses or bytes32 repositories
		addressIndex := c.AddressIndexes[padded32str]
		if addressIndex != 0 {
			// There are 4 different flags for reading addresses,
			// depending if the index fits on 2, 3, or 4 bytes
			bytesNeeded := minBytesToRepresent(addressIndex)
			switch bytesNeeded {
			case 1, 2:
				return []byte{
					byte(FLAG_READ_ADDRESS_2),
					byte(addressIndex >> 8),
					byte(addressIndex),
				}, ReadStorage, nil
			case 3:
				return []byte{
					byte(FLAG_READ_ADDRESS_3),
					byte(addressIndex >> 16),
					byte(addressIndex >> 8),
					byte(addressIndex),
				}, ReadStorage, nil
			case 4:
				return []byte{
					byte(FLAG_READ_ADDRESS_4),
					byte(addressIndex >> 24),
					byte(addressIndex >> 16),
					byte(addressIndex >> 8),
					byte(addressIndex),
				}, ReadStorage, nil
			}
		}

		bytes32Index := c.Bytes32Indexes[padded32str]
		if bytes32Index != 0 {
			// There are 4 different flags for reading bytes32,
			// depending if the index fits on 2, 3, or 4 bytes
			bytesNeeded := minBytesToRepresent(bytes32Index)
			switch bytesNeeded {
			case 1, 2:
				return []byte{
					byte(FLAG_READ_BYTES32_2),
					byte(bytes32Index >> 8),
					byte(bytes32Index),
				}, ReadStorage, nil
			case 3:
				return []byte{
					byte(FLAG_READ_BYTES32_3),
					byte(bytes32Index >> 16),
					byte(bytes32Index >> 8),
					byte(bytes32Index),
				}, ReadStorage, nil
			case 4:
				return []byte{
					byte(FLAG_READ_BYTES32_4),
					byte(bytes32Index >> 24),
					byte(bytes32Index >> 16),
					byte(bytes32Index >> 8),
					byte(bytes32Index),
				}, ReadStorage, nil
			}
		}

		if saveWord {
			// Any value smaller than 20 bytes can be saved as an address
			// ALL saved values must be padded to either 20 bytes or 32 bytes
			// For both cases skip values that are too short already
			if len(trimmed) <= 20 && len(trimmed) >= 15 {
				padded20 := make([]byte, 20)
				copy(padded20[20-len(trimmed):], trimmed)
				encoded := []byte{byte(FLAG_SAVE_ADDRESS)}
				encoded = append(encoded, padded20...)
				return encoded, WriteStorage, nil

			} else if len(trimmed) >= 27 {
				encoded := []byte{byte(FLAG_SAVE_BYTES32)}
				encoded = append(encoded, padded32...)
				return encoded, WriteStorage, nil
			}
		}
	}

	// We are out of options now, we need to encode the word as-is
	return c.EncodeWordBytes32(trimmed)
}

// Encodes a 32 word, without any optimizations
func (c *Compressor) EncodeWordBytes32(word []byte) ([]byte, EncodeType, error) {
	if len(word) > 32 {
		return nil, Stateless, fmt.Errorf("word exceeds 32 bytes")
	}

	if len(word) == 0 {
		return nil, Stateless, fmt.Errorf("word is empty")
	}

	encodedWord := []byte{byte(FLAG_READ_BYTES32_1_BYTES + uint(len(word)) - 1)}
	encodedWord = append(encodedWord, word...)
	return encodedWord, Stateless, nil
}

func (c *Compressor) WriteWord(word []byte, dest *CBuffer, useStorage bool) (EncodeType, error) {
	encoded, t, err := c.EncodeWordOptimized(dest, word, useStorage)
	if err != nil {
		return Stateless, err
	}

	paddedWord := make([]byte, 32)
	copy(paddedWord[32-len(word):], word)

	dest.WriteBytes(encoded)
	dest.End(paddedWord, t)

	return t, nil
}

// Encode N bytes, as optimized as possible
func (c *Compressor) WriteBytesOptimized(dest *CBuffer, bytes []byte, saveWord bool) (EncodeType, error) {
	// Empty bytes can be represented with a no-op
	// cost: 0
	if len(bytes) == 0 {
		dest.WriteInt(FLAG_NO_OP)
		dest.End(bytes, Stateless)
		return Stateless, nil
	}

	// 32 bytes long can be encoded as a word, it has its own set of optimizations.
	// cost: word
	if len(bytes) == 32 {
		return c.WriteWord(bytes, dest, saveWord)
	}

	// Now we can try to find a mirror flag for the bytes
	// cost: 2 bytes
	bytesStr := string(bytes)
	usedFlag := dest.Refs.usedFlags[bytesStr]
	if usedFlag != 0 {
		usedFlag -= 1

		// We can only encode 16 bits for the mirror flag
		// if it exceeds this value, then we can't mirror it
		if usedFlag <= 0xffff {
			dest.WriteInt(FLAG_MIRROR_FLAG)
			dest.WriteBytes([]byte{byte(usedFlag >> 8), byte(usedFlag)})
			// End without creating a second pointer
			// otherwise we will be creating a pointer to a pointer
			dest.End([]byte{}, Mirror)
			return Mirror, nil
		}
	}

	// Another optimization is to copy the bytes from the calldata
	// cost: 3 bytes
	copyIndex := FindPastData(dest, bytes)
	if copyIndex != -1 && copyIndex <= 0xffff {
		dest.WriteInt(FLAG_COPY_CALLDATA)
		dest.WriteBytes([]byte{byte(copyIndex >> 8), byte(copyIndex), byte(len(bytes))})
		// End without creating a second pointer
		// data can be copied from the calldata directly
		dest.End([]byte{}, Stateless)
		return Mirror, nil
	}

	// If the bytes are 33 bytes long, and the first byte is 0x03 it can be represented as a "node"
	// cost: 0 bytes + word
	if len(bytes) == 33 && bytes[0] == 0x03 {
		dest.WriteInt(FLAG_NODE)
		dest.End(bytes, Stateless)

		t, err := c.WriteWord(bytes[1:], dest, saveWord)
		if err != nil {
			return Stateless, err
		}

		return t, nil
	}

	// If the bytes are 33 bytes long and starts with 0x05 it can be represented as a "subdigest"
	// cost: 0 bytes + word
	if len(bytes) == 33 && bytes[0] == 0x05 {
		dest.WriteInt(FLAG_SUBDIGEST)
		dest.End(bytes, Stateless)

		t, err := c.WriteWord(bytes[1:], dest, saveWord)
		if err != nil {
			return Stateless, err
		}

		return t, nil
	}

	// If bytes has 22 bytes and starts with 0x01, then it is probably an address on a signature
	// cost: 1 / 0 bytes + address word
	if len(bytes) == 22 && bytes[0] == 0x01 {
		// If the firt byte (weight) is between 1 and 4, then there is a special flag
		if bytes[1] >= 1 && bytes[1] <= 4 {
			dest.WriteInt(FLAG_ADDRESS_W0 + uint(bytes[1]))
		} else {
			// We need to use FLAG_ADDRES_W0 and 1 extra byte for the weight
			dest.WriteInt(FLAG_ADDRESS_W0)
			dest.WriteByte(bytes[1])
		}

		dest.End(bytes, Stateless)

		t, err := c.WriteWord(bytes[2:], dest, saveWord)
		if err != nil {
			return Stateless, err
		}

		return t, nil
	}

	// If the bytes are 68 bytes long and starts with 0x00, the it is probably a signature for a Sequence wallet
	// cost: 66/67 bytes
	if len(bytes) == 68 && bytes[0] == 0x00 {
		// If the first byte (weight) is between 1 and 4, then there is a special flag
		if bytes[1] >= 1 && bytes[1] <= 4 {
			dest.WriteInt(FLAG_SIGNATURE_W0 + uint(bytes[1]))
		} else {
			// We need to use FLAG_SIGNATURE_W0 and 1 extra byte for the weight
			dest.WriteInt(FLAG_SIGNATURE_W0)
			dest.WriteByte(bytes[1])
		}

		dest.WriteBytes(bytes[2:])
		dest.End(bytes, Stateless)
		return Stateless, nil
	}

	// We can try encoding this as a signature, we don't know if it is a Sequence signature
	// we generate a snapshot of the dest buffer and try to encode it as a signature, if it fails
	// or if the result is bigger than the original bytes, we restore the snapshot and continue.
	// Notice: pass `false` to `mayUseBytes` or else this will be an infinite loop
	snapshot := dest.Snapshot()
	t, err := c.WriteSignature(dest, bytes, false)
	if err == nil && dest.Len() < len(bytes)+3+len(snapshot.Commited) {
		return t, nil
	}
	dest.Restore(snapshot)

	// If the bytes are a multiple of 32 + 4 bytes (max 6 * 32 + 4) then it
	// can be encoded as an ABI call with 0 to 6 parameters
	if len(bytes) <= 6*32+4 && (len(bytes)-4)%32 == 0 {
		dest.WriteInt(FLAG_ABI_0_PARAM + uint((len(bytes)-4)/32))
		dest.WriteBytes(c.Encode4Bytes(bytes[:4]))
		dest.End(bytes, Stateless)

		encodeType := Stateless

		for i := 4; i < len(bytes); i += 32 {
			t, err := c.WriteWord(bytes[i:i+32], dest, saveWord)
			if err != nil {
				return Stateless, err
			}

			encodeType = HighestPriority(encodeType, t)
		}

		return encodeType, nil
	}

	// If the bytes are a multiple of 32 + 4 bytes (max 256 * 32 + 4) then it
	// can be represented using dynamic encoded ABI
	if len(bytes) <= 256*32+4 && (len(bytes)-4)%32 == 0 {
		dest.WriteInt(FLAG_READ_DYNAMIC_ABI)
		dest.WriteBytes(c.Encode4Bytes(bytes[:4]))
		dest.WriteInt(uint((len(bytes) - 4) / 32)) // The number of ARGs
		// This flag can be used to compress dynamic size arguments too
		// but in this case, we just leave it as 0s so all arguments are 32 bytes
		dest.WriteInt(0)
		dest.End(bytes, Stateless)

		encodeType := Stateless

		for i := 4; i < len(bytes); i += 32 {
			t, err := c.WriteWord(bytes[i:i+32], dest, saveWord)
			if err != nil {
				return Stateless, err
			}

			encodeType = HighestPriority(encodeType, t)
		}

		return encodeType, nil
	}

	// If there are no other options, then we encode the bytes as-is
	// we can try two different methods: bytes_n or splitting it in many words + an extra bytes
	// for now we leave it as bytes_n for simplicity
	return c.WriteNBytesRaw(dest, bytes)
}

// Encodes N bytes, without any optimizations
func (c *Compressor) WriteNBytesRaw(dest *CBuffer, bytes []byte) (EncodeType, error) {
	dest.WriteInt(FLAG_READ_N_BYTES)
	dest.End(bytes, Stateless)

	t, err := c.WriteWord(uintToBytes(uint64(len(bytes))), dest, false)
	if err != nil {
		return Stateless, err
	}

	dest.WriteBytes(bytes)

	// End this last write without creating a flag pointer
	// this is a data blob, not a flag
	dest.End([]byte{}, t)

	return t, nil
}

func (c *Compressor) Encode4Bytes(bytes []byte) []byte {
	// If Bytes4Indexes has this value, then we can just use the index
	index := c.Bytes4Indexes[string(bytes)]
	if index != 0 {
		return []byte{byte(index)}
	}

	// If don't then we need to provide it as-is, but it has to be prefixed with 0x00
	return append([]byte{0x00}, bytes...)
}

func (c *Compressor) WriteNonce(dest *CBuffer, nonce *big.Int, randomNonce bool) (EncodeType, error) {
	paddedNonce := make([]byte, 32)
	copy(paddedNonce[32-len(nonce.Bytes()):], nonce.Bytes())

	// The nonce is encoded in two parts, the first 160 bits are the space
	// the last 96 bits are the nonce itself, the nonce space can be saved to storage
	// unless it is a random nonce
	spaceBytes := paddedNonce[:20]
	nonceBytes := paddedNonce[20:]

	t1, err := c.WriteWord(spaceBytes, dest, !randomNonce)
	if err != nil {
		return Stateless, err
	}

	t2, err := c.WriteWord(nonceBytes, dest, false)
	if err != nil {
		return Stateless, err
	}

	return HighestPriority(t1, t2), nil
}

func (c *Compressor) WriteTransactions(dest *CBuffer, transactions sequence.Transactions) (EncodeType, error) {
	if len(transactions) == 0 {
		return Stateless, fmt.Errorf("transactions is empty")
	}

	if len(transactions) > 255 {
		return Stateless, fmt.Errorf("transactions exceeds 255")
	}

	// The first byte is the number of transactions
	dest.WriteInt(uint(len(transactions)))
	dest.End([]byte{}, Stateless)

	encodeType := Stateless

	for _, transaction := range transactions {
		t, err := c.WriteTransaction(dest, transaction)
		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	return encodeType, nil
}

func (c *Compressor) WriteTransaction(dest *CBuffer, transaction *sequence.Transaction) (EncodeType, error) {
	// The first byte is the flag of the transaction, the specification follows:
	// - 1000 0000 - 1 if it uses delegate call
	// - 0100 0000 - 1 if it uses revert on error
	// - 0010 0000 - 1 if it has a defined gas limit
	// - 0001 0000 - 1 if it has a defined value
	// - 0000 1000 - Unused
	// - 0000 0100 - Unused
	// - 0000 0010 - Unused
	// - 0000 0001 - 1 if it has a defined data

	flag := byte(0x00)

	if transaction.DelegateCall {
		flag |= 0x80
	}

	if transaction.RevertOnError {
		flag |= 0x40
	}

	hasGasLimit := transaction.GasLimit != nil && transaction.GasLimit.Cmp(big.NewInt(0)) != 0
	if hasGasLimit {
		flag |= 0x20
	}

	hasValue := transaction.Value != nil && transaction.Value.Cmp(big.NewInt(0)) != 0
	if hasValue {
		flag |= 0x10
	}

	hasData := len(transaction.Data) > 0 || len(transaction.Transactions) > 0
	if hasData {
		flag |= 0x01
	}

	dest.WriteByte(flag)
	dest.End([]byte{}, Stateless)

	encodeType := Stateless

	// Now we start writing the values that we do have
	if hasGasLimit {
		t, err := c.WriteWord(transaction.GasLimit.Bytes(), dest, false)
		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	// Encode the address as a word
	t, err := c.WriteWord(transaction.To.Bytes(), dest, dest.Refs.useContractStorage)
	if err != nil {
		return Stateless, err
	}

	encodeType = HighestPriority(encodeType, t)

	if hasValue {
		t, err := c.WriteWord(transaction.Value.Bytes(), dest, false)
		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	if hasData {
		// If the transaction has "transactions" then it is another sequence transaction
		// in that case, we must Write a transaction, since data is empty
		var t EncodeType

		if len(transaction.Transactions) > 0 {
			dest.WriteInt(FLAG_READ_EXECUTE)
			dest.End([]byte{}, Stateless)

			t, err = c.WriteExecute(dest, nil, transaction)
		} else {
			t, err = c.WriteBytesOptimized(dest, transaction.Data, dest.Refs.useContractStorage)
		}

		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	return encodeType, nil
}

func (c *Compressor) WriteSignature(dest *CBuffer, signature []byte, mayUseBytes bool) (EncodeType, error) {
	dest.SignatureLevel += 1

	defer func() {
		dest.SignatureLevel -= 1
	}()

	// First byte determines the signature type
	if mayUseBytes && (len(signature) == 0) {
		// Guestmodule signatures are empty
		return c.WriteBytesOptimized(dest, signature, false)
	}

	typeByte := signature[0]

	switch typeByte {
	case 0x00: // Legacy
		return c.WriteSignatureBody(dest, false, signature)
	case 0x01: // Dynamic
		return c.WriteSignatureBody(dest, false, signature[1:])
	case 0x02: // No chain ID
		return c.WriteSignatureBody(dest, true, signature[1:])
	case 0x03: // Chained
		return c.WriteChainedSignature(dest, signature[1:])

	default:
		return Stateless, fmt.Errorf("invalid signature type %d", typeByte)
	}
}

func (c *Compressor) WriteChainedSignature(dest *CBuffer, signature []byte) (EncodeType, error) {
	// First we need to count how many chained signatures are there
	// for this we read the first 3 bytes of each, as they contain the size
	var pointer uint
	var parts [][]byte

	for pointer < uint(len(signature)) {
		length := uint(signature[pointer])<<16 | uint(signature[pointer+1])<<8 | uint(signature[pointer+2])
		pointer += 3

		npointer := pointer + length
		parts = append(parts, signature[pointer:npointer])
		pointer = npointer
	}

	// We have two instructions for this, one for 8 bits and one for 16 bits
	// depending on the number of parts
	totalParts := uint(len(parts))
	if totalParts > 255 {
		dest.WriteInt(FLAG_READ_CHAINED_L)
		dest.WriteByte(byte(totalParts >> 8))
		dest.WriteByte(byte(totalParts))
	} else {
		dest.WriteInt(FLAG_READ_CHAINED)
		dest.WriteByte(byte(totalParts))
	}

	dest.End([]byte{}, Stateless)

	// Now we need to encode every nested part, one for each signature part
	encodeType := Stateless

	for _, part := range parts {
		t, err := c.WriteSignature(dest, part, false)
		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	return encodeType, nil
}

func (c *Compressor) WriteExecute(dest *CBuffer, wallet []byte, transaction *sequence.Transaction) (EncodeType, error) {
	var t EncodeType

	t, err := c.WriteNonce(dest, transaction.Nonce, false)
	if err != nil {
		return Stateless, err
	}

	tt, err := c.WriteTransactions(dest, transaction.Transactions)
	if err != nil {
		return Stateless, err
	}
	t = HighestPriority(t, tt)

	tt, err = c.WriteSignature(dest, transaction.Signature, true)
	if err != nil {
		return Stateless, err
	}
	t = HighestPriority(t, tt)

	if wallet != nil {
		tt, err = c.WriteWord(wallet, dest, true)
		if err != nil {
			return Stateless, err
		}
	}

	return HighestPriority(t, tt), nil
}

func (c *Compressor) WriteSignatureBody(dest *CBuffer, noChain bool, body []byte) (EncodeType, error) {
	// The first 2 bytes are the threshold
	// this (alongside the noChain flag) defines
	// the encoding flag
	if len(body) < 2 {
		return Stateless, fmt.Errorf("signature is too short")
	}

	threshold := uint(body[0])<<8 | uint(body[1])
	longThreshold := threshold > 0xff

	var tflag uint

	if !longThreshold && !noChain {
		tflag = FLAG_S_SIG
	} else if !longThreshold && noChain {
		tflag = FLAG_S_SIG_NO_CHAIN
	} else if longThreshold && !noChain {
		tflag = FLAG_S_L_SIG
	} else {
		tflag = FLAG_S_L_SIG_NO_CHAIN
	}

	dest.WriteInt(tflag)

	// On long threshold we use 2 bytes for the threshold
	if longThreshold {
		dest.WriteBytes(body[:2])
	} else {
		dest.WriteByte(body[1])
	}

	dest.End(body, Stateless)

	// Next 4 bytes is the checkpoint
	if len(body) < 6 {
		return Stateless, fmt.Errorf("signature is too short")
	}

	checkpoint := body[2:6]
	c.WriteWord(checkpoint, dest, false)

	return c.WriteSignatureTree(dest, body[6:])
}

func (c *Compressor) WriteSignatureTree(dest *CBuffer, tree []byte) (EncodeType, error) {
	// Signature trees need to be encoded as N nested bytes (one per part)
	// for this we first need to count the number of parts
	if len(tree) == 0 {
		return Stateless, fmt.Errorf("signature tree is empty")
	}

	totalParts := 0
	pointer := uint(0)

	for pointer < uint(len(tree)) {
		partType := tree[pointer]
		pointer += 1

		switch partType {
		case 0x00: // EOA Signature
			pointer += (1 + 66)
		case 0x01: // Address
			pointer += (1 + 20)
		case 0x02: // Dynamic
			pointer += (1 + 20)
			// 3 bytes after address and weight are the length
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += (3 + length)
		case 0x03: // Node
			pointer += 32
		case 0x04: // Branch
			// 3 bytes of length
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += (3 + length)
		case 0x05: // Subdigest
			pointer += 32
		case 0x06: // Nested
			pointer += 3
			// 3 bytes of length
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += (3 + length)
		default:
			return Stateless, fmt.Errorf("invalid signature part type %d", partType)
		}

		totalParts += 1
	}

	if totalParts > 1 {
		if totalParts > 255 {
			dest.WriteInt(FLAG_NESTED_N_FLAGS_16)
			dest.WriteByte(byte(totalParts >> 8))
			dest.WriteByte(byte(totalParts))
		} else {
			dest.WriteInt(FLAG_NESTED_N_FLAGS_8)
			dest.WriteByte(byte(totalParts))
		}
	}

	dest.End([]byte{}, Stateless)

	// Now we need to encode every nested part, one for each signature part
	encodeType := Stateless
	pointer = uint(0)
	for pointer < uint(len(tree)) {
		partType := tree[pointer]
		pointer += 1

		var t EncodeType
		var err error

		switch partType {
		case 0x00: // EOA Signature
			next := pointer + (1 + 66)
			t, err = c.WriteBytesOptimized(dest, tree[pointer-1:next], false)
			pointer = next
		case 0x01: // Address
			next := pointer + (1 + 20)
			t, err = c.WriteBytesOptimized(dest, tree[pointer-1:next], true)
			encodeType = HighestPriority(encodeType, t)
			pointer = next
		case 0x02: // Dynamic
			weight := uint(tree[pointer])
			pointer += 1

			addr := tree[pointer : pointer+20]
			pointer += 20
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += 3

			next := pointer + length
			t, err = c.WriteDynamicSignaturePart(dest, addr, weight, tree[pointer:next])
			pointer = next
		case 0x03: // Node
			next := pointer + 32
			t, err = c.WriteBytesOptimized(dest, tree[pointer-1:next], true)
			pointer = next
		case 0x04: // Branch
			// 3 bytes of length
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += 3

			next := pointer + length
			t, err = c.WriteBranchSignaturePart(dest, tree[pointer:next])
			pointer = next
		case 0x05: // Subdigest
			next := pointer + 32
			t, err = c.WriteBytesOptimized(dest, tree[pointer-1:next], false)
			pointer = next
		case 0x06: // Nested
			weight := uint(tree[pointer])
			pointer += 1
			threshold := uint(tree[pointer])<<8 | uint(tree[pointer+1])
			pointer += 2
			length := uint(tree[pointer])<<16 | uint(tree[pointer+1])<<8 | uint(tree[pointer+2])
			pointer += 3

			next := pointer + length
			t, err = c.WriteNestedSignaturePart(dest, weight, threshold, tree[pointer:next])
			pointer = next
		default:
			// This should never happen
			return Stateless, fmt.Errorf("invalid signature part type, wtf, unreachable code")
		}

		if err != nil {
			return Stateless, err
		}

		encodeType = HighestPriority(encodeType, t)
	}

	return encodeType, nil
}

func (c *Compressor) WriteNestedSignaturePart(dest *CBuffer, weight uint, threshold uint, branch []byte) (EncodeType, error) {
	if weight > 255 {
		return Stateless, fmt.Errorf("weight exceeds 255")
	}

	if threshold > 255 {
		return Stateless, fmt.Errorf("threshold exceeds 255")
	}

	dest.WriteInt(FLAG_NESTED)
	dest.WriteInt(weight)
	dest.WriteInt(threshold)
	dest.End([]byte{}, Stateless)

	return c.WriteSignatureTree(dest, branch)
}

func (c *Compressor) WriteBranchSignaturePart(dest *CBuffer, branch []byte) (EncodeType, error) {
	if len(branch) == 0 {
		return Stateless, fmt.Errorf("branch is empty")
	}

	dest.WriteInt(FLAG_BRANCH)
	dest.End([]byte{}, Stateless)

	return c.WriteSignatureTree(dest, branch)
}

func (c *Compressor) WriteDynamicSignaturePart(dest *CBuffer, address []byte, weight uint, signature []byte) (EncodeType, error) {
	if weight > 255 {
		return Stateless, fmt.Errorf("weight exceeds 255")
	}

	if signature[len(signature)-1] != 0x03 {
		return Stateless, fmt.Errorf("signature is not a dynamic signature")
	}

	unsuffixed := signature[:len(signature)-1]

	dest.WriteInt(FLAG_DYNAMIC_SIGNATURE)
	dest.WriteInt(weight)
	dest.End([]byte{}, Stateless)

	// The address must be 20 bytes long
	// write it as a word
	if len(address) != 20 {
		return Stateless, fmt.Errorf("address is not 20 bytes long")
	}

	t1, err := c.WriteWord(address, dest, true)
	if err != nil {
		return Stateless, err
	}

	// Encode the signature using bytes, as it may or may not be a Sequence signature
	// bytes is going to try to encode it as a Sequence signature, if it fails it will
	// encode it as a bunch of bytes
	t2, err := c.WriteBytesOptimized(dest, unsuffixed, true)
	if err != nil {
		return Stateless, err
	}

	return HighestPriority(t1, t2), nil
}
