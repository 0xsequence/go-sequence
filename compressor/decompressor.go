package compressor

import (
	"encoding/binary"
	"fmt"
	"math/big"
)

type Decompressor struct {
	calldata []byte
	rindex   uint
	buffer   []byte

	indexToAddress map[uint][]byte
	indexToBytes32 map[uint][]byte
	indexToBytes4  map[uint][]byte
}

func (d *Decompressor) Buffer() []byte {
	return d.buffer
}

func NewDecompressor(calldata []byte) *Decompressor {
	return &Decompressor{
		calldata: calldata,
		rindex:   1,
		buffer:   make([]byte, 0),

		indexToAddress: make(map[uint][]byte),
		indexToBytes32: make(map[uint][]byte),
		indexToBytes4:  make(map[uint][]byte),
	}
}

func (d *Decompressor) LogFlag(flag string) {
	// fmt.Println("flag:", flag)
}

func (d *Decompressor) ReadFlag() error {
	flag := uint(d.calldata[d.rindex])
	d.rindex++

	if flag == FLAG_READ_POWER_OF_10_MISC {
		return d.ReadPow10Misc()
	}

	if flag >= FLAG_READ_BYTES32_1_BYTES && flag <= FLAG_READ_BYTES32_32_BYTES {
		return d.ReadBytes32(flag)
	}

	if flag >= FLAG_READ_ADDRESS_2 && flag <= FLAG_READ_ADDRESS_4 {
		return d.ReadAddressStorage(flag)
	}

	if flag >= FLAG_READ_BYTES32_2 && flag <= FLAG_READ_BYTES32_4 {
		return d.ReadBytes32Storage(flag)
	}

	if flag == FLAG_SAVE_ADDRESS {
		return d.ReadSaveAddress()
	}

	if flag == FLAG_SAVE_BYTES32 {
		return d.ReadSaveBytes32()
	}

	if flag == FLAG_READ_N_BYTES {
		return d.ReadNBytes(flag)
	}

	if flag == FLAG_READ_POWER_OF_2 {
		return d.ReadPow2()
	}

	if flag >= FLAG_ABI_0_PARAM && flag <= FLAG_ABI_6_PARAMS {
		return d.ReadAbiStatic(flag)
	}

	if flag >= FLAG_NESTED_N_FLAGS_8 && flag <= FLAG_NESTED_N_FLAGS_16 {
		return d.ReadNestedFlags(flag)
	}

	if flag >= FLAG_SIGNATURE_W0 && flag <= FLAG_SIGNATURE_W4 {
		return d.ReadSignature(flag)
	}

	if flag >= FLAG_ADDRESS_W0 && flag <= FLAG_ADDRESS_W4 {
		return d.ReadAddress(flag)
	}

	if flag == FLAG_NODE {
		return d.FlagNode()
	}

	if flag == FLAG_BRANCH {
		return d.ReadBranch()
	}

	if flag == FLAG_NESTED {
		return d.ReadNested()
	}

	if flag == FLAG_DYNAMIC_SIGNATURE {
		return d.ReadDynamicSignature()
	}

	if flag >= FLAG_S_SIG_NO_CHAIN && flag <= FLAG_S_L_SIG {
		return d.ReadSequenceSignatureV2(flag)
	}

	if flag == FLAG_READ_CHAINED || flag == FLAG_READ_CHAINED_L {
		return fmt.Errorf("read chained not implemented")
	}

	if flag == FLAG_READ_DYNAMIC_ABI {
		return d.ReadAbiDynamic()
	}

	if flag == FLAG_NO_OP {
		return nil
	}

	if flag == FLAG_MIRROR_FLAG || flag == FLAG_READ_STORE_FLAG {
		return d.ReadMirrorFlag()
	}

	if flag == FLAG_COPY_CALLDATA {
		return d.ReadCopyCalldata()
	}

	return d.ReadLiteral(flag)
}

func (d *Decompressor) ReadAndLoad32Bytes() ([]byte, error) {
	d.LogFlag("read_and_load_flag")

	d.ReadFlag()

	// Load the last 32 bytes and return them to the caller
	// they must be removed from the buffer
	if len(d.buffer) < 32 {
		return nil, fmt.Errorf("not enough bytes in buffer")
	}

	// Read the last 32 bytes
	word := d.buffer[len(d.buffer)-32:]
	d.buffer = d.buffer[:len(d.buffer)-32]
	return word, nil
}

func (d *Decompressor) ReadBytes32(flag uint) error {
	d.LogFlag("bytes32")

	// FLAG_READ_BYTES32_1_BYTES reads 1 byte
	// FLAG_READ_BYTES32_1_BYTES + 1 reads 2 bytes
	// ... etc
	readb := flag - FLAG_READ_BYTES32_1_BYTES + 1
	word := d.calldata[d.rindex : d.rindex+readb]

	// Word is always read padded with 0s to 32 bytes
	padded, err := padToX(word, 32)
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, padded...)
	d.rindex += readb

	return nil
}

func (d *Decompressor) ReadSaveAddress() error {
	d.LogFlag("save_address")

	// Read 20 bytes, pad it to 32 bytes
	addr := d.calldata[d.rindex : d.rindex+20]
	padded, err := padToX(addr, 32)
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, padded...)
	d.rindex += 20

	return nil
}

func (d *Decompressor) ReadSaveBytes32() error {
	d.LogFlag("save_bytes32")

	// Read 32 bytes
	word := d.calldata[d.rindex : d.rindex+32]
	d.buffer = append(d.buffer, word...)
	d.rindex += 32

	return nil
}

func (d *Decompressor) ReadAddressStorage(flag uint) error {
	d.LogFlag("address_storage")

	// Number of bytes used for the index:
	// FLAG_READ_ADDRESS_2 -> 2 bytes
	// FLAG_READ_ADDRESS_2 + 1 -> 3 bytes
	// ... etc
	readb := flag - FLAG_READ_ADDRESS_2 + 2
	if readb < 2 {
		return fmt.Errorf("reading less than 2 bytes on address %d", flag)
	}

	if readb > 4 {
		return fmt.Errorf("reading more than 5 bytes on address %d", flag)
	}

	ibs := d.calldata[d.rindex : d.rindex+readb]
	d.rindex += readb

	index := bytesToUint64(ibs)

	// Read the address from storage
	addr := d.indexToAddress[index]
	if addr == nil {
		return fmt.Errorf("address %d not found", index)
	}

	d.buffer = append(d.buffer, addr...)
	return nil
}

func (d *Decompressor) ReadBytes32Storage(flag uint) error {
	d.LogFlag("bytes32_storage")

	// Number of bytes used for the index:
	// FLAG_READ_BYTES32_2 -> 2 bytes
	// FLAG_READ_BYTES32_2 + 1 -> 3 bytes
	// ... etc
	readb := flag - FLAG_READ_BYTES32_2 + 2

	if readb < 2 {
		return fmt.Errorf("reading less than 2 bytes on bytes32 %d", flag)
	}

	if readb > 4 {
		return fmt.Errorf("reading more than 5 bytes on bytes32 %d", flag)
	}

	ibs := d.calldata[d.rindex : d.rindex+readb]
	d.rindex += readb

	index := bytesToUint64(ibs)

	// Read the bytes32 from storage
	bytes32 := d.indexToBytes32[index]
	if bytes32 == nil {
		return fmt.Errorf("bytes32 %d not found", index)
	}

	d.buffer = append(d.buffer, bytes32...)
	return nil
}

func (d *Decompressor) ReadNBytes(flag uint) error {
	d.LogFlag("n_bytes")

	// Read a nested flag, this gives us the number of bytes to read
	nbytes, err := d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	n := uint(binary.BigEndian.Uint64(nbytes[len(nbytes)-8:]))
	d.LogFlag("n_bytes " + fmt.Sprintf("%d", n))

	// Read the number of bytes specified by the nested flag
	d.buffer = append(d.buffer, d.calldata[d.rindex:d.rindex+n]...)
	d.rindex += n

	return nil
}

func (d *Decompressor) ReadPow2() error {
	d.LogFlag("pow2")

	exp := int(d.calldata[d.rindex])
	d.rindex++

	var num *big.Int
	if exp == 0 {
		// We need to read another exp, and this time do 2 ** exp - 1
		exp = int(d.calldata[d.rindex])
		d.rindex++

		base := big.NewInt(2)
		pow := big.NewInt(int64(exp))
		num = new(big.Int).Exp(base, pow, nil)
	} else {
		base := big.NewInt(2)
		pow := big.NewInt(int64(exp))
		num = new(big.Int).Exp(base, pow, nil)
	}

	// Write the number padded to 32 bytes
	padded, err := padToX(num.Bytes(), 32)
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, padded...)

	return nil
}

func (d *Decompressor) ReadAbi4Bytes() error {
	d.LogFlag("abi_4_bytes")

	// The first value is always the bytes4, it may be an index
	// or the bytes4 itself (if it is prefixed with 00)
	ib4 := d.calldata[d.rindex]
	d.rindex++

	var selector []byte

	if ib4 == 0 {
		// Read the next 4 bytes
		selector = d.calldata[d.rindex : d.rindex+4]
		d.rindex += 4
	} else {
		selector = d.indexToBytes4[uint(ib4)]
	}

	d.buffer = append(d.buffer, selector...)
	return nil
}

func (d *Decompressor) ReadAbiStatic(flag uint) error {
	d.LogFlag("abi_static")

	err := d.ReadAbi4Bytes()
	if err != nil {
		return err
	}

	// The number of args is determined by the flag
	// FLAG_ABI_0_PARAM -> 0 args
	// FLAG_ABI_0_PARAM + 1 -> 1 arg
	// ... etc

	nargs := flag - FLAG_ABI_0_PARAM
	if nargs > 6 {
		return fmt.Errorf("reading more than 6 args on abi static %d", flag)
	}

	return d.ReadNFlags(nargs)
}

func (d *Decompressor) ReadAbiDynamic() error {
	d.LogFlag("abi_dynamic")

	err := d.ReadAbi4Bytes()
	if err != nil {
		return err
	}

	// There are two flags, the second one marks which args are dynamic
	// the first one marks how many args there are, both are 1 byte
	// (it is only possible to make the first 8 args as dynamic)
	fs := uint(d.calldata[d.rindex])
	d.rindex++

	fd := d.calldata[d.rindex]
	d.rindex++

	windex2 := uint(len(d.buffer))

	// Reserve 32 bytes for each arg
	d.buffer = append(d.buffer, make([]byte, fs*32)...)

	for i := uint(0); i < fs; i++ {
		argSpot := windex2 + (i * 32)

		if (fd & (1 << i)) != 0 {
			// This arg is dynamic. we need to write a relative pointer
			// to the end of the buffer (from the start of the args)
			rpointer := uint(len(d.buffer)) - windex2
			padded, err := uintPadToX(rpointer, 32)
			if err != nil {
				return err
			}

			copy(d.buffer[argSpot:], padded)

			// Reserve 32 bytes for the size
			sizeSpot := uint(len(d.buffer))
			d.buffer = append(d.buffer, make([]byte, 32)...)

			// Now we track the size and we write another flag
			// to read the actual bytes
			windex3 := uint(len(d.buffer))
			err = d.ReadFlag()
			if err != nil {
				return err
			}

			size := uint(len(d.buffer)) - windex3

			// Whatever we wrote, it has to be a multiple of 32
			// so we pad the buffer to 32 bytes
			pdamount := uint(32 - (size % 32))
			if pdamount != 32 {
				d.buffer = append(d.buffer, make([]byte, pdamount)...)
			}

			padded = make([]byte, 32)
			binary.BigEndian.PutUint64(padded, uint64(size))
			copy(d.buffer[sizeSpot:], padded)
		} else {
			// This arg is static, we just read a nested ON argSpot
			val, err := d.ReadAndLoad32Bytes()
			if err != nil {
				return err
			}

			copy(d.buffer[argSpot:], val)
		}
	}

	return nil
}

func (d *Decompressor) ReadNestedFlags(flag uint) error {
	d.LogFlag("nested_flags")

	// FLAG_NESTED_N_FLAGS_8 -> 1 byte for n of flags
	// FLAG_NESTED_N_FLAGS_8 + 1 -> 2 bytes for n of flags
	nb := flag - FLAG_NESTED_N_FLAGS_8 + 1

	if nb < 1 || nb > 2 {
		return fmt.Errorf("reading more than 2 bytes on nested flags %d", flag)
	}

	var nflags uint

	// Read the number of flags
	if nb == 1 {
		nflags = uint(d.calldata[d.rindex])
	} else {
		nflags = uint(binary.BigEndian.Uint16(d.calldata[d.rindex : d.rindex+nb]))
	}
	d.rindex += nb

	return d.ReadNFlags(nflags)
}

func (d *Decompressor) ReadNFlags(n uint) error {
	d.LogFlag("n_flags")

	for i := uint(0); i < n; i++ {
		err := d.ReadFlag()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Decompressor) ReadSignature(flag uint) error {
	d.LogFlag("signature_part")

	// FLAG_SIGNATURE_W0 -> for 1 byte for the weight (has to read more)
	// FLAG_SIGNATURE_W1 -> for weight == 1
	// FLAG_SIGNATURE_W2 -> for weight == 2
	// ... etc (max 4)

	weight := flag - FLAG_SIGNATURE_W0

	if weight > 4 {
		return fmt.Errorf("signature static weight too high %d", flag)
	}

	if weight == 0 {
		// Read another byte for the weight
		weight = uint(d.calldata[d.rindex])
		d.rindex++
	}

	// Write the "signature" sequence flag (0x00)
	d.buffer = append(d.buffer, 0x00)

	// Write the weight
	d.buffer = append(d.buffer, byte(weight))

	// Now just copy 66 bytes
	d.buffer = append(d.buffer, d.calldata[d.rindex:d.rindex+66]...)
	d.rindex += 66

	return nil
}

func (d *Decompressor) ReadAddress(flag uint) error {
	d.LogFlag("address_part")

	// FLAG_ADDRESS_W0 -> for 1 byte for the weight (has to read more)
	// FLAG_ADDRESS_W1 + 1 -> for weight == 1
	// ... etc (max 4)
	weight := flag - FLAG_ADDRESS_W0

	if weight > 4 {
		return fmt.Errorf("address static weight too high %d", flag)
	}

	if weight == 0 {
		// Read another byte for the weight
		weight = uint(d.calldata[d.rindex])
		d.rindex++
	}

	// Write the "address" sequence flag (0x01)
	d.buffer = append(d.buffer, 0x01)

	// Write the weight
	d.buffer = append(d.buffer, byte(weight))

	// Now read a nested flag, the difference is that we will shift the value
	// to the left by 12 bytes
	val, err := d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	// Pad the value to 20 bytes
	padded, err := padToX(val, 20)
	d.buffer = append(d.buffer, padded...)

	return nil
}

func (d *Decompressor) ReadDynamicSignature() error {
	d.LogFlag("dynamic_signature_part")

	// Write the "dynamic signature" sequence flag (0x02)
	d.buffer = append(d.buffer, 0x02)

	// Read 1 byte as the weight
	weight := uint(d.calldata[d.rindex])
	d.rindex++
	d.buffer = append(d.buffer, byte(weight))

	// Read the address as a nested flag
	val, err := d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	// Write only the last 20 bytes
	d.buffer = append(d.buffer, val[12:]...)

	// Reserve 3 bytes for the size, we are going to read
	// the dynamic signature as a nested flag
	d.buffer = append(d.buffer, make([]byte, 3)...)
	windex := uint(len(d.buffer))

	err = d.ReadFlag()
	if err != nil {
		return err
	}

	// Write an extra "03" to the end, this is the 1271 flag
	d.buffer = append(d.buffer, 0x03)

	size := uint(len(d.buffer)) - windex
	if size > 0xffffff {
		return fmt.Errorf("dynamic signature size too big %d", size)
	}

	padded, err := uintPadToX(size, 3)
	if err != nil {
		return err
	}

	copy(d.buffer[windex-3:], padded)

	return nil
}

func (d *Decompressor) FlagNode() error {
	d.LogFlag("node")

	// Write the node flag (0x03) and just read a nested flag
	d.buffer = append(d.buffer, 0x03)
	return d.ReadFlag()
}

func (d *Decompressor) ReadBranch() error {
	d.LogFlag("branch")

	// Write the branch flag (0x04) and just read a nested flag
	d.buffer = append(d.buffer, 0x04)

	// Reserve 3 bytes for the size
	sizeSpot := uint(len(d.buffer))
	d.buffer = append(d.buffer, make([]byte, 3)...)

	// Now track the size and read another flag
	windex := uint(len(d.buffer))
	err := d.ReadFlag()
	if err != nil {
		return err
	}

	size := uint(len(d.buffer)) - windex

	// Write the size on the reserved spot
	// notice that it must not be bigger than 2^24
	if size > 0xffffff {
		return fmt.Errorf("branch size too big %d", size)
	}

	padded, err := uintPadToX(size, 3)
	if err != nil {
		return err
	}

	copy(d.buffer[sizeSpot:], padded)

	return nil
}

func (d *Decompressor) FlagSubdigest() error {
	d.LogFlag("subdigest")

	// Write the subdigest flag (0x05) and just read a nested flag
	d.buffer = append(d.buffer, 0x05)
	return d.ReadFlag()
}

func (d *Decompressor) ReadNested() error {
	d.LogFlag("nested")

	// Write the nested flag (0x06) and just read a nested flag
	d.buffer = append(d.buffer, 0x06)

	// We use 1 byte for the weight
	weight := uint(d.calldata[d.rindex])
	d.rindex++

	// Write the weight
	d.buffer = append(d.buffer, byte(weight))

	// Another byte represent the threshold, but we write it on 2 bytes
	threshold := uint(d.calldata[d.rindex])
	d.rindex++

	padded, err := uintPadToX(threshold, 2)
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, padded...)

	// Now read a nested flag, keeping track of the size
	// and reserving 3 bytes for the size
	d.buffer = append(d.buffer, make([]byte, 3)...)

	windex := uint(len(d.buffer))
	err = d.ReadFlag()
	if err != nil {
		return err
	}

	size := uint(len(d.buffer)) - windex
	if size > 0xffffff {
		return fmt.Errorf("nested size too big %d", size)
	}

	padded = make([]byte, 3)
	binary.BigEndian.PutUint64(padded, uint64(size))
	copy(d.buffer[windex-3:], padded)

	return nil
}

func (d *Decompressor) ReadPow10Misc() error {
	d.LogFlag("pow10misc")

	exp := uint(d.calldata[d.rindex])
	d.rindex++

	// If the exp is 0, then we are actually pointing to an extension of the flagset
	// in this case we only have one method "READ_SELF_EXECUTE"
	if exp == 0 {
		return fmt.Errorf("read self execute not implemented")
	}

	// The first bit determines if we have a mantissa or not
	hasm := (exp & 0x80) == 0
	// Print exp in binary
	exp &= 0x7f

	var num *big.Int
	num = big.NewInt(10)
	num = num.Exp(num, big.NewInt(int64(exp)), nil)

	if hasm {
		// Read another byte for the mantissa
		mantissa := uint(d.calldata[d.rindex])
		d.rindex++

		num = num.Mul(num, big.NewInt(int64(mantissa)))
	}

	// It should not exceed 2 ** 256 - 1
	if len(num.Bytes()) > 32 {
		return fmt.Errorf("pow10misc number too big %d", exp)
	}

	// Write the number padded to 32 bytes
	padded, err := padToX(num.Bytes(), 32)
	if err != nil {
		return err
	}
	d.buffer = append(d.buffer, padded...)

	return nil
}

func (d *Decompressor) ReadMirrorFlag() error {
	// The next 2 bytes determine the temporal rindex
	trindex := uint(binary.BigEndian.Uint16(d.calldata[d.rindex : d.rindex+2]))
	d.rindex += 2

	prevrindex := d.rindex
	d.LogFlag("mirror " + fmt.Sprintf("%d %d", trindex, prevrindex))

	if prevrindex-3 == trindex {
		return fmt.Errorf("mirror flag pointing to the same rindex %d", trindex)
	}

	// Replace the rindex with the temporal one
	// and read another flag
	d.rindex = trindex
	err := d.ReadFlag()
	if err != nil {
		return err
	}

	// Restore the rindex
	d.rindex = prevrindex

	return nil
}

func (d *Decompressor) ReadCopyCalldata() error {
	d.LogFlag("copy_calldata")

	// The next 2 bytes determine from where to copy, the next byte determines the size
	// this is a simple copy of the calldata
	from := uint(binary.BigEndian.Uint16(d.calldata[d.rindex : d.rindex+2]))
	d.rindex += 2

	size := uint(d.calldata[d.rindex])
	d.rindex++

	d.buffer = append(d.buffer, d.calldata[from:from+size]...)
	return nil
}

func (d *Decompressor) ReadSequenceSignatureV2(flag uint) error {
	d.LogFlag("sequence_signature_v2")

	// The flag determines the type of signature
	var noChainId bool
	if flag == FLAG_S_L_SIG || flag == FLAG_S_SIG {
		noChainId = false
	} else {
		noChainId = true
	}

	// Write the sequence signature flag (0x02 for noChain, 0x01 for chain)
	if noChainId {
		d.buffer = append(d.buffer, 0x02)
	} else {
		d.buffer = append(d.buffer, 0x01)
	}

	// The threshold may be provided as 1 or 2 bytes (the flag also tells you)
	var threshold uint
	if flag == FLAG_S_L_SIG || flag == FLAG_S_L_SIG_NO_CHAIN {
		threshold = uint(binary.BigEndian.Uint16(d.calldata[d.rindex : d.rindex+2]))
		d.rindex += 2
	} else {
		threshold = uint(d.calldata[d.rindex])
		d.rindex++
	}

	// Write the threshold (padded to two bytes)
	padded, err := uintPadToX(threshold, 2)
	if err != nil {
		return err
	}
	d.buffer = append(d.buffer, padded...)

	// The checkpoint is always just 4 bytes, read it as a word
	val, err := d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	padded, err = padToX(val, 4)
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, padded...)

	// Now we read the signature tree, this is just a nested flag
	return d.ReadFlag()
}

func (d *Decompressor) ReadNonce() error {
	d.LogFlag("read_nonce")

	// Read a word, but use only the last 20 bytes
	// then read another word, and use only the last 12 bytes
	val, err := d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, val[12:]...)

	val, err = d.ReadAndLoad32Bytes()
	if err != nil {
		return err
	}

	d.buffer = append(d.buffer, val[20:]...)
	return nil
}

func (d *Decompressor) ReadTransactions() error {
	// The first byte determines the number of transactions
	n := uint(d.calldata[d.rindex])
	d.rindex++

	d.LogFlag("read_transactions " + fmt.Sprintf("%d", n))

	// Write the number of transactions, padded to 32 bytes
	padded, err := uintPadToX(n, 32)
	if err != nil {
		return err
	}
	d.buffer = append(d.buffer, padded...)

	// Reserve 32 bytes for each transaction's pointer
	posPointer := uint(len(d.buffer))
	d.buffer = append(d.buffer, make([]byte, n*32)...)

	for i := uint(0); i < n; i++ {
		// Write the pointer to the transaction (this is to the end of the buffer)
		pointer := uint(len(d.buffer)) - posPointer + (i * 32)
		padded, err := uintPadToX(pointer, 32)
		if err != nil {
			return err
		}
		copy(d.buffer[posPointer:], padded)

		// Now read the transaction
		err = d.ReadTransaction()
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Decompressor) ReadTransaction() error {
	d.LogFlag("read_transaction")

	// The first byte is a bitmap, it contains information about what values are defined
	// - 1000 0000 - 1 if it uses delegate call
	// - 0100 0000 - 1 if it uses revert on error
	// - 0010 0000 - 1 if it has a defined gas limit
	// - 0001 0000 - 1 if it has a defined value
	// - 0000 1000 - Unused
	// - 0000 0100 - Unused
	// - 0000 0010 - Unused
	// - 0000 0001 - 1 if it has a defined data
	bitmap := d.calldata[d.rindex]
	d.rindex++

	// Write 0x01 or 0x00 (delegatecall) padded to 32 bytes
	d.buffer = append(d.buffer, make([]byte, 31)...)
	if (bitmap & 0x80) != 0 {
		d.buffer = append(d.buffer, 0x01)
	} else {
		d.buffer = append(d.buffer, 0x00)
	}

	// Write 0x01 or 0x00 (revertOnError) padded to 32 bytes
	d.buffer = append(d.buffer, make([]byte, 31)...)
	if (bitmap & 0x40) != 0 {
		d.buffer = append(d.buffer, 0x01)
	} else {
		d.buffer = append(d.buffer, 0x00)
	}

	// Now gas limit may need to be read as a flag
	if (bitmap & 0x20) != 0 {
		err := d.ReadFlag()
		if err != nil {
			return err
		}
	} else {
		// Write 0x00 padded to 32 bytes
		d.buffer = append(d.buffer, make([]byte, 32)...)
	}

	// All transactions have a defined address, this is just a flag
	err := d.ReadFlag()
	if err != nil {
		return err
	}

	// Same for value, read as a flag
	if (bitmap & 0x10) != 0 {
		err := d.ReadFlag()
		if err != nil {
			return err
		}
	} else {
		// Write 0x00 padded to 32 bytes
		d.buffer = append(d.buffer, make([]byte, 32)...)
	}

	// If the transaction has no data, we need to write 0x00 padded to 32 bytes twice
	if (bitmap & 0x01) == 0 {
		d.buffer = append(d.buffer, make([]byte, 64)...)
	} else {
		// Write the pointer of the data, it is always padded 0xc0
		d.buffer = append(d.buffer, make([]byte, 31)...)
		d.buffer = append(d.buffer, 0xc0)

		// Reserve 32 bytes for the size
		sizeSpot := uint(len(d.buffer))
		d.buffer = append(d.buffer, make([]byte, 32)...)
		windex := uint(len(d.buffer))

		err := d.ReadFlag()
		if err != nil {
			return err
		}

		size := uint(len(d.buffer)) - windex

		// Write the size on the reserved spot, padded to 32 bytes
		padded, err := uintPadToX(size, 32)
		if err != nil {
			return err
		}

		copy(d.buffer[sizeSpot:], padded)

		// Pad the buffer to 32 bytes
		pdamount := uint(32 - (size % 32))
		if pdamount != 32 {
			d.buffer = append(d.buffer, make([]byte, pdamount)...)
		}
	}

	return nil
}

func (d *Decompressor) ReadExecute() error {
	d.LogFlag("execute")

	// Write the execute method (0x7a9a1628)
	d.buffer = append(d.buffer, []byte{0x7a, 0x9a, 0x16, 0x28}...)

	// Write the start of the transaction's list (always 0x60 padded to 32 bytes)
	d.buffer = append(d.buffer, make([]byte, 31)...)
	d.buffer = append(d.buffer, 0x60)

	// Read the nonce
	err := d.ReadNonce()
	if err != nil {
		return err
	}

	// We can't know where the signatures will start (we need to read the transactions)
	// so we leave 32 bytes and a pointer to the start of the signatures
	// (we will write the pointer later)
	sigsPointer := uint(len(d.buffer))
	d.buffer = append(d.buffer, make([]byte, 32)...)

	// Read the transactions
	err = d.ReadTransactions()
	if err != nil {
		return err
	}

	// Write the pointer to the start of the signatures
	pointer := uint(len(d.buffer)) - 4
	padded, err := uintPadToX(pointer, 32)
	if err != nil {
		return err
	}
	copy(d.buffer[sigsPointer:], padded)

	// Read the signatures, handle the size and the padding
	sizePointer := uint(len(d.buffer))
	d.buffer = append(d.buffer, make([]byte, 32)...)
	windex := uint(len(d.buffer))

	err = d.ReadFlag()
	if err != nil {
		return err
	}

	// Write the size padded to 32 bytes on the reserved spot
	size := uint(len(d.buffer)) - windex
	padded, err = uintPadToX(size, 32)
	if err != nil {
		return err
	}
	copy(d.buffer[sizePointer:], padded)

	// Padd the buffer to 32 bytes
	pdamount := uint(32 - (size % 32))
	if pdamount != 32 {
		d.buffer = append(d.buffer, make([]byte, pdamount)...)
	}

	return nil
}

func (d *Decompressor) ReadLiteral(flag uint) error {
	d.LogFlag("literal " + fmt.Sprintf("%d", flag))

	if flag < LITERAL_ZERO {
		return fmt.Errorf("literal flag too low %d", flag)
	}

	// Write the literal number (LITERAL_ZERO = 0, LITERAL_ZERO + 1 = 1, etc)
	// padded to 32 bytes
	padded, err := uintPadToX(flag-LITERAL_ZERO, 32)
	if err != nil {
		return err
	}
	d.buffer = append(d.buffer, padded...)

	return nil
}
