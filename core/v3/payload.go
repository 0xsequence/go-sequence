package v3

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

const (
	KindTransactions uint8 = 0x00
	KindMessage      uint8 = 0x01
	KindConfigUpdate uint8 = 0x02
	KindDigest       uint8 = 0x03
)

type DecodedPayload struct {
	Kind          uint8
	NoChainId     bool
	Calls         []Call
	Space         *big.Int
	Nonce         *big.Int
	Message       []byte
	ImageHash     common.Hash
	Digest        common.Hash
	ParentWallets []common.Address
}

func DecodeABIPayload(input string) (DecodedPayload, error) {
	data, err := hexutil.Decode(input)

	if err != nil {
		return DecodedPayload{}, fmt.Errorf("invalid hex input: %w", err)
	}

	const abiDef = `[{
		"name": "f",
		"type": "function",
		"inputs": [{
			"type": "tuple",
			"components": [
				{"name": "kind", "type": "uint8"},
				{"name": "noChainId", "type": "bool"},
				{"name": "calls", "type": "tuple[]", "components": [
					{"name": "to", "type": "address"},
					{"name": "value", "type": "uint256"},
					{"name": "data", "type": "bytes"},
					{"name": "gasLimit", "type": "uint256"},
					{"name": "delegateCall", "type": "bool"},
					{"name": "onlyFallback", "type": "bool"},
					{"name": "behaviorOnError", "type": "uint256"}
				]},
				{"name": "space", "type": "uint256"},
				{"name": "nonce", "type": "uint256"},
				{"name": "message", "type": "bytes"},
				{"name": "imageHash", "type": "bytes32"},
				{"name": "digest", "type": "bytes32"},
				{"name": "parentWallets", "type": "address[]"}
			]
		}]
	}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiDef))
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to parse ABI: %w", err)
	}

	decodedData, err := parsedABI.Methods["f"].Inputs.Unpack(data)
	if err != nil {
		return DecodedPayload{}, fmt.Errorf("failed to decode ABI: %w", err)
	}

	if len(decodedData) != 1 {
		return DecodedPayload{}, fmt.Errorf("invalid decoded data length")
	}

	v := decodedData[0]
	if v == nil {
		return DecodedPayload{}, fmt.Errorf("decoded data is nil")
	}

	if s, ok := v.(struct {
		Kind      uint8 `json:"kind"`
		NoChainId bool  `json:"noChainId"`
		Calls     []struct {
			To              common.Address `json:"to"`
			Value           *big.Int       `json:"value"`
			Data            []uint8        `json:"data"`
			GasLimit        *big.Int       `json:"gasLimit"`
			DelegateCall    bool           `json:"delegateCall"`
			OnlyFallback    bool           `json:"onlyFallback"`
			BehaviorOnError *big.Int       `json:"behaviorOnError"`
		} `json:"calls"`
		Space         *big.Int         `json:"space"`
		Nonce         *big.Int         `json:"nonce"`
		Message       []uint8          `json:"message"`
		ImageHash     [32]uint8        `json:"imageHash"`
		Digest        [32]uint8        `json:"digest"`
		ParentWallets []common.Address `json:"parentWallets"`
	}); ok {
		calls := make([]Call, len(s.Calls))
		for i, call := range s.Calls {
			calls[i] = Call{
				To:              call.To,
				Value:           call.Value,
				Data:            call.Data,
				GasLimit:        call.GasLimit,
				DelegateCall:    call.DelegateCall,
				OnlyFallback:    call.OnlyFallback,
				BehaviorOnError: BehaviorOnError(call.BehaviorOnError.Uint64()),
			}
		}

		var imageHash common.Hash
		copy(imageHash[:], s.ImageHash[:])

		var digest common.Hash
		copy(digest[:], s.Digest[:])

		return DecodedPayload{
			Kind:          s.Kind,
			NoChainId:     s.NoChainId,
			Calls:         calls,
			Space:         s.Space,
			Nonce:         s.Nonce,
			Message:       s.Message,
			ImageHash:     imageHash,
			Digest:        digest,
			ParentWallets: s.ParentWallets,
		}, nil
	}

	return DecodedPayload{}, fmt.Errorf("failed to decode payload")
}

func minIntBytesFor(val *big.Int) int {
	if val == nil {
		return 0
	}
	return max(1, (val.BitLen()+7)/8)
}

func encodeBehaviorOnError(behavior BehaviorOnError) byte {
	switch behavior {
	case BehaviorOnErrorIgnore:
		return 0
	case BehaviorOnErrorRevert:
		return 1
	case BehaviorOnErrorAbort:
		return 2
	default:
		return 0
	}
}

func EncodeDecodedPayload(payload DecodedPayload, self *common.Address) ([]byte, error) {
	if payload.Kind != KindTransactions {
		return nil, fmt.Errorf("encode only supports transaction payloads")
	}

	if payload.Space == nil {
		payload.Space = big.NewInt(0)
	}
	if payload.Nonce == nil {
		payload.Nonce = big.NewInt(0)
	}

	callsLen := len(payload.Calls)
	nonceBytesNeeded := minIntBytesFor(payload.Nonce)
	if nonceBytesNeeded > 15 {
		return nil, fmt.Errorf("nonce is too large (requires %d bytes, max 15)", nonceBytesNeeded)
	}

	/*
		globalFlag layout:
		  bit 0: spaceZeroFlag => 1 if space == 0, else 0
		  bits [1..3]: how many bytes we use to encode nonce
		  bit 4: singleCallFlag => 1 if there's exactly one call
		  bit 5: callsCountSizeFlag => 1 if #calls stored in 2 bytes, 0 if in 1 byte
		  (bits [6..7] are unused/free)
	*/
	var globalFlag byte
	if payload.Space.Sign() == 0 {
		globalFlag |= 0x01
	}
	globalFlag |= byte(nonceBytesNeeded << 1)
	if callsLen == 1 {
		globalFlag |= 0x10
	}

	// Determine calls count size (1 or 2 bytes) if not a single call
	var callsCountSize int
	if callsLen != 1 {
		if callsLen < 256 {
			callsCountSize = 1
		} else if callsLen < 65536 {
			callsCountSize = 2
			globalFlag |= 0x20
		} else {
			return nil, fmt.Errorf("too many calls: %d (max 65535)", callsLen)
		}
	}

	// Start building the output with a bytes.Buffer
	var out bytes.Buffer
	out.WriteByte(globalFlag)

	// If space isn't 0, store it as exactly 20 bytes (like uint160)
	if payload.Space.Sign() != 0 {
		spaceBytes := make([]byte, 20)
		payload.Space.FillBytes(spaceBytes)
		out.Write(spaceBytes)
	}

	// Encode nonce in nonceBytesNeeded bytes
	if nonceBytesNeeded > 0 {
		nonceBytes := make([]byte, nonceBytesNeeded)
		payload.Nonce.FillBytes(nonceBytes)
		out.Write(nonceBytes)
	}

	// Store callsLen if not single-call
	if callsLen != 1 {
		if callsCountSize == 1 {
			out.WriteByte(byte(callsLen))
		} else {
			out.Write([]byte{byte(callsLen >> 8), byte(callsLen)})
		}
	}

	// Encode each call
	for _, call := range payload.Calls {
		var flags byte
		if self != nil && call.To == *self {
			flags |= 0x01
		}
		if call.Value != nil && call.Value.Sign() != 0 {
			flags |= 0x02
		}
		if len(call.Data) > 0 {
			flags |= 0x04
		}
		if call.GasLimit != nil && call.GasLimit.Sign() != 0 {
			flags |= 0x08
		}
		if call.DelegateCall {
			flags |= 0x10
		}
		if call.OnlyFallback {
			flags |= 0x20
		}
		behavior := encodeBehaviorOnError(call.BehaviorOnError)
		flags |= behavior << 6

		// Write the flags byte
		out.WriteByte(flags)

		// If toSelf bit not set, store 20-byte address
		if flags&0x01 == 0 {
			addrBytes := call.To.Bytes()
			if len(addrBytes) != 20 {
				return nil, fmt.Errorf("invalid 'to' address length: %d bytes", len(addrBytes))
			}
			out.Write(addrBytes)
		}

		// If hasValue, store 32 bytes of value
		if flags&0x02 != 0 {
			valueBytes := make([]byte, 32)
			if call.Value == nil {
				call.Value = big.NewInt(0)
			}
			call.Value.FillBytes(valueBytes)
			out.Write(valueBytes)
		}

		// If hasData, store 3 bytes of data length + data
		if flags&0x04 != 0 {
			dataLen := len(call.Data)
			if dataLen > 0xffffff { // Max value for 3 bytes
				return nil, fmt.Errorf("call data too large: %d bytes (max 16777215)", dataLen)
			}
			dataLenBytes := []byte{
				byte(dataLen >> 16),
				byte(dataLen >> 8),
				byte(dataLen),
			}
			out.Write(dataLenBytes)
			out.Write(call.Data)
		}

		// If hasGasLimit, store 32 bytes of gasLimit
		if flags&0x08 != 0 {
			gasBytes := make([]byte, 32)
			if call.GasLimit == nil {
				call.GasLimit = big.NewInt(0)
			}
			call.GasLimit.FillBytes(gasBytes)
			out.Write(gasBytes)
		}
	}

	// Return the encoded bytes
	return out.Bytes(), nil
}

type PayloadDigestable interface {
	Digest() PayloadDigest
}

var (
	_ PayloadDigestable = PayloadDigest{}
)

type Payload interface {
	PayloadDigestable

	Address() common.Address
	ChainID() *big.Int
	Encode(address common.Address) []byte
}

var (
	_ Payload = CallsPayload{}
	_ Payload = MessagePayload{}
	_ Payload = ConfigUpdatePayload{}
	_ Payload = DigestPayload{}
)

type PayloadDigest struct {
	common.Hash

	Preimage Payload
}

func (d PayloadDigest) Digest() PayloadDigest {
	return d
}

type basePayload struct {
	address       common.Address
	chainID       *big.Int
	parentWallets []common.Address
}

func (p basePayload) Address() common.Address {
	return p.address
}

func (p basePayload) ChainID() *big.Int {
	return p.chainID
}

var eip712Domain = []ethcoder.TypedDataArgument{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}

func (p basePayload) domain() ethcoder.TypedDataDomain {
	return ethcoder.TypedDataDomain{
		Name:              "Sequence Wallet",
		Version:           "3",
		ChainID:           p.chainID,
		VerifyingContract: &p.address,
	}
}

func NewCallsPayload(address common.Address, chainID *big.Int, calls_ []Call, space, nonce *big.Int, parentWallets ...[]common.Address) CallsPayload {
	if chainID == nil {
		chainID = big.NewInt(0)
	}
	if len(calls_) >= 0x10000 {
		panic(fmt.Errorf("number of calls %v >= 2^16", len(calls_)))
	}
	if space == nil {
		space = big.NewInt(0)
	}
	if space.Sign() < 0 || space.Cmp(new(big.Int).Lsh(common.Big1, 160)) >= 0 {
		panic(fmt.Errorf("space %v is out of bounds [0, 2^160)", space))
	}
	if nonce == nil {
		nonce = big.NewInt(0)
	}
	if nonce.Sign() < 0 || nonce.Cmp(new(big.Int).Lsh(common.Big1, 56)) >= 0 {
		panic(fmt.Errorf("nonce %v is out of bounds [0, 2^56)", nonce))
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return CallsPayload{
		basePayload: basePayload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		Calls: calls_,
		Space: space,
		Nonce: nonce,
	}
}

type Call struct {
	To              common.Address
	Value           *big.Int
	Data            []byte
	GasLimit        *big.Int
	DelegateCall    bool
	OnlyFallback    bool
	BehaviorOnError BehaviorOnError
}

func (c Call) asMap() map[string]any {
	value := c.Value
	if value == nil {
		value = big.NewInt(0)
	}

	data := c.Data
	if data == nil {
		data = []byte{}
	}

	gasLimit := c.GasLimit
	if gasLimit == nil {
		gasLimit = big.NewInt(0)
	}

	return map[string]any{
		"to":              c.To,
		"value":           value,
		"data":            data,
		"gasLimit":        gasLimit,
		"delegateCall":    c.DelegateCall,
		"onlyFallback":    c.OnlyFallback,
		"behaviorOnError": new(big.Int).SetInt64(int64(c.BehaviorOnError)),
	}
}

func (c Call) encode(address common.Address) []byte {
	var flags byte
	var buffer bytes.Buffer

	if c.To == address {
		flags |= 0x01
	} else {
		mustWrite(&buffer, c.To.Bytes())
	}

	if c.Value != nil && c.Value.Sign() != 0 {
		flags |= 0x02
		mustWrite(&buffer, common.BigToHash(c.Value).Bytes())
	}

	if len(c.Data) != 0 {
		flags |= 0x04
		size := len(c.Data)
		mustWrite(&buffer, []byte{byte(size >> 16), byte(size >> 8), byte(size)})
		mustWrite(&buffer, c.Data)
	}

	if c.GasLimit != nil && c.GasLimit.Sign() != 0 {
		flags |= 0x08
		mustWrite(&buffer, common.BigToHash(c.GasLimit).Bytes())
	}

	if c.DelegateCall {
		flags |= 0x10
	}

	if c.OnlyFallback {
		flags |= 0x20
	}

	flags |= byte(c.BehaviorOnError << 6)

	data := make([]byte, 1+buffer.Len())
	data[0] = flags
	copy(data[1:], buffer.Bytes())
	return data
}

func decodeCall(data []byte, address common.Address) (Call, []byte, error) {
	call := Call{
		Value:    new(big.Int),
		GasLimit: new(big.Int),
	}

	if len(data) < 1 {
		return Call{}, nil, fmt.Errorf("no flags")
	}
	var flags byte
	flags, data = data[0], data[1:]

	if flags&0x01 != 0 {
		call.To = address
	} else {
		if len(data) < 20 {
			return Call{}, nil, fmt.Errorf("no to")
		}
		call.To, data = common.BytesToAddress(data[:20]), data[20:]
	}

	if flags&0x02 != 0 {
		if len(data) < 32 {
			return Call{}, nil, fmt.Errorf("no value")
		}
		call.Value, data = new(big.Int).SetBytes(data[:32]), data[32:]
	}

	if flags&0x04 != 0 {
		if len(data) < 3 {
			return Call{}, nil, fmt.Errorf("no data size")
		}
		var dataSize int
		dataSize, data = int(data[0])<<16+int(data[1])<<8+int(data[2]), data[3:]
		if len(data) < dataSize {
			return Call{}, nil, fmt.Errorf("no data")
		}
		call.Data, data = data[:dataSize], data[dataSize:]
	}

	if flags&0x08 != 0 {
		if len(data) < 32 {
			return Call{}, nil, fmt.Errorf("no gas limit")
		}
		call.GasLimit, data = new(big.Int).SetBytes(data[:32]), data[32:]
	}

	call.DelegateCall = flags&0x10 != 0
	call.OnlyFallback = flags&0x20 != 0
	call.BehaviorOnError = BehaviorOnError(flags & 0xc0 >> 6)

	return call, data, nil
}

type BehaviorOnError int

const (
	BehaviorOnErrorIgnore BehaviorOnError = 0
	BehaviorOnErrorRevert BehaviorOnError = 1
	BehaviorOnErrorAbort  BehaviorOnError = 2
)

type CallsPayload struct {
	basePayload

	Calls []Call
	Space *big.Int
	Nonce *big.Int
}

func (p CallsPayload) Digest() PayloadDigest {
	calls := make([]any, 0, len(p.Calls))
	for _, call := range p.Calls {
		calls = append(calls, any(call.asMap()))
	}

	wallets := make([]any, 0, len(p.parentWallets))
	for _, wallet := range p.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      p.domain(),
		PrimaryType: "Calls",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"Calls": {
				{Name: "calls", Type: "Call[]"},
				{Name: "space", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "wallets", Type: "address[]"},
			},
			"Call": {
				{Name: "to", Type: "address"},
				{Name: "value", Type: "uint256"},
				{Name: "data", Type: "bytes"},
				{Name: "gasLimit", Type: "uint256"},
				{Name: "delegateCall", Type: "bool"},
				{Name: "onlyFallback", Type: "bool"},
				{Name: "behaviorOnError", Type: "uint256"},
			},
		},
		Message: map[string]any{
			"calls":   calls,
			"space":   p.Space,
			"nonce":   p.Nonce,
			"wallets": wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: p}
}

func (p CallsPayload) Encode(address common.Address) []byte {
	var flags byte
	var buffer bytes.Buffer

	if p.Space.Sign() == 0 {
		flags |= 0x01
	} else {
		var space [20]byte
		mustWrite(&buffer, p.Space.FillBytes(space[:]))
	}

	nonce := p.Nonce.Bytes()
	flags |= byte(len(nonce)) << 1
	mustWrite(&buffer, nonce)

	calls := len(p.Calls)
	if calls == 1 {
		flags |= 0x10
	} else if calls < 0x100 {
		mustWrite(&buffer, []byte{byte(calls)})
	} else {
		flags |= 0x20
		mustWrite(&buffer, []byte{byte(calls >> 8), byte(calls)})
	}

	for _, call := range p.Calls {
		mustWrite(&buffer, call.encode(address))
	}

	data := make([]byte, 1+buffer.Len())
	data[0] = flags
	copy(data[1:], buffer.Bytes())
	return data
}

func DecodeCalls(address common.Address, chainID *big.Int, data []byte) (CallsPayload, error) {
	var space, nonce *big.Int
	space = big.NewInt(0)
	nonce = big.NewInt(0)

	if len(data) < 1 {
		return CallsPayload{}, fmt.Errorf("no flags")
	}
	var flags byte
	flags, data = data[0], data[1:]

	if flags&0x01 == 0 {
		if len(data) < 20 {
			return CallsPayload{}, fmt.Errorf("no space")
		}
		space.SetBytes(data[:20])
		data = data[20:]
	}

	nonceSize := int(flags >> 1 & 0x07)
	if len(data) < nonceSize {
		return CallsPayload{}, fmt.Errorf("no nonce")
	}
	nonce.SetBytes(data[:nonceSize])
	data = data[nonceSize:]

	var calls_ int
	if flags&0x10 != 0 {
		calls_ = 1
	} else if flags&0x20 == 0 {
		if len(data) < 1 {
			return CallsPayload{}, fmt.Errorf("no number of calls")
		}
		calls_, data = int(data[0]), data[1:]
	} else {
		if len(data) < 2 {
			return CallsPayload{}, fmt.Errorf("no number of calls")
		}
		calls_, data = int(data[0])<<8+int(data[1]), data[2:]
	}

	var calls []Call
	for len(calls) < calls_ {
		var call Call
		var err error
		call, data, err = decodeCall(data, address)
		if err != nil {
			return CallsPayload{}, fmt.Errorf("unable to decode call %v: %w", len(calls), err)
		}

		calls = append(calls, call)
	}

	if len(data) != 0 {
		return CallsPayload{}, fmt.Errorf("%v trailing bytes in payload", len(data))
	}

	return NewCallsPayload(address, chainID, calls, space, nonce), nil
}

func NewMessagePayload(address common.Address, chainID *big.Int, message_ []byte, parentWallets ...[]common.Address) Payload {
	if chainID == nil {
		chainID = big.NewInt(0)
	}
	if message_ == nil {
		message_ = []byte{}
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return MessagePayload{
		basePayload: basePayload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		message: message_,
	}
}

type MessagePayload struct {
	basePayload

	message []byte
}

func (p MessagePayload) Digest() PayloadDigest {
	wallets := make([]any, 0, len(p.parentWallets))
	for _, wallet := range p.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      p.domain(),
		PrimaryType: "Message",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"Message": {
				{Name: "message", Type: "bytes"},
				{Name: "wallets", Type: "address[]"},
			},
		},
		Message: map[string]any{
			"message": p.message,
			"wallets": wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: p}
}

func (p MessagePayload) Encode(address common.Address) []byte {
	var buffer bytes.Buffer

	// Write kind byte for Message (0x01)
	mustWrite(&buffer, []byte{KindMessage})

	// Write message length as 3 bytes
	size := len(p.message)
	mustWrite(&buffer, []byte{byte(size >> 16), byte(size >> 8), byte(size)})

	// Write message data
	if len(p.message) > 0 {
		mustWrite(&buffer, p.message)
	}

	return buffer.Bytes()
}

func DecodeMessage(address common.Address, chainID *big.Int, data []byte) (Payload, error) {
	if len(data) < 4 { // kind byte + 3 bytes length
		return nil, fmt.Errorf("message payload too short")
	}
	data = data[1:] // Skip kind byte
	msgLen := int(data[0])<<16 + int(data[1])<<8 + int(data[2])
	data = data[3:]
	if len(data) < msgLen {
		return nil, fmt.Errorf("message data too short")
	}
	message := data[:msgLen]
	return NewMessagePayload(address, chainID, message), nil
}

func NewConfigUpdatePayload(address common.Address, chainID *big.Int, imageHash common.Hash, parentWallets ...[]common.Address) Payload {
	if chainID == nil {
		chainID = big.NewInt(0)
	}
	if isNil(imageHash) {
		panic(fmt.Errorf("no config"))
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return ConfigUpdatePayload{
		basePayload: basePayload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		imageHash: imageHash,
	}
}

type ConfigUpdatePayload struct {
	basePayload

	imageHash common.Hash
}

func (p ConfigUpdatePayload) Digest() PayloadDigest {
	wallets := make([]any, 0, len(p.parentWallets))
	for _, wallet := range p.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      p.domain(),
		PrimaryType: "ConfigUpdate",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"ConfigUpdate": {
				{Name: "imageHash", Type: "bytes32"},
				{Name: "wallets", Type: "address[]"},
			},
		},
		Message: map[string]any{
			"imageHash": p.imageHash,
			"wallets":   wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: p}
}

func (p ConfigUpdatePayload) Encode(address common.Address) []byte {
	var buffer bytes.Buffer

	// Write kind byte for ConfigUpdate (0x02)
	mustWrite(&buffer, []byte{KindConfigUpdate})

	// Write image hash
	mustWrite(&buffer, p.imageHash.Bytes())

	return buffer.Bytes()
}

func DecodeConfigUpdate(address common.Address, chainID *big.Int, data []byte) (Payload, error) {
	if len(data) < 33 {
		return nil, fmt.Errorf("config update payload too short")
	}
	data = data[1:]
	var hash common.Hash
	copy(hash[:], data[:32])
	return NewConfigUpdatePayload(address, chainID, hash), nil
}

func NewDigestPayload(address common.Address, chainID *big.Int, digest_ common.Hash, parentWallets ...[]common.Address) Payload {
	if chainID == nil {
		chainID = big.NewInt(0)
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return DigestPayload{
		basePayload: basePayload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		digest: digest_,
	}
}

type DigestPayload struct {
	basePayload

	digest common.Hash
}

func (p DigestPayload) Digest() PayloadDigest {
	domain, _ := ethcoder.ABIPackArguments(
		[]string{
			"bytes32",
			"bytes32",
			"bytes32",
			"uint256",
			"address",
		},
		[]any{
			crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)")),
			crypto.Keccak256Hash([]byte("Sequence Wallet")),
			crypto.Keccak256Hash([]byte("3")),
			p.chainID,
			p.address,
		},
	)

	state := crypto.NewKeccakState()
	for _, wallet := range p.parentWallets {
		var word common.Hash
		word.SetBytes(wallet.Bytes())
		mustWrite(state, word.Bytes())
	}
	var wallets common.Hash
	_, err := state.Read(wallets[:])
	if err != nil {
		panic(err)
	}

	message, _ := ethcoder.ABIPackArguments(
		[]string{
			"bytes32",
			"bytes32",
			"bytes32",
		},
		[]any{
			crypto.Keccak256Hash([]byte("Message(bytes message,address[] wallets)")),
			p.digest,
			wallets,
		},
	)

	return PayloadDigest{Hash: crypto.Keccak256Hash(
		[]byte("\x19\x01"),
		crypto.Keccak256(domain),
		crypto.Keccak256(message),
	), Preimage: p}
}

func (p DigestPayload) Encode(address common.Address) []byte {
	var buffer bytes.Buffer

	// Write kind byte for Digest (0x03)
	mustWrite(&buffer, []byte{KindDigest})

	// Write digest
	mustWrite(&buffer, p.digest.Bytes())

	return buffer.Bytes()
}

func mustWrite(writer io.Writer, data []byte) {
	_, err := writer.Write(data)
	if err != nil {
		panic(err)
	}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	value_ := reflect.ValueOf(value)
	switch value_.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Interface:
		return value_.IsNil()
	}
	return false
}

func DecodeRawPayload(address common.Address, chainID *big.Int, data []byte) (Payload, error) {
	decoded, err := DecodeABIPayload(hexutil.Encode(data))
	if err != nil {
		return nil, err
	}

	switch decoded.Kind {
	case KindTransactions:
		return NewCallsPayload(address, chainID, decoded.Calls, decoded.Space, decoded.Nonce, []common.Address(decoded.ParentWallets)), nil
	case KindMessage:
		return NewMessagePayload(address, chainID, decoded.Message, []common.Address(decoded.ParentWallets)), nil
	case KindConfigUpdate:
		return NewConfigUpdatePayload(address, chainID, decoded.ImageHash, []common.Address(decoded.ParentWallets)), nil
	case KindDigest:
		return NewDigestPayload(address, chainID, decoded.Digest, []common.Address(decoded.ParentWallets)), nil
	default:
		return nil, fmt.Errorf("unknown payload kind: %d", decoded.Kind)
	}
}
