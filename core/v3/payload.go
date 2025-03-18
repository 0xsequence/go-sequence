package v3

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"reflect"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/go-sequence/core"
)

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
}

type EncodablePayload interface {
	Payload

	Encode(address common.Address) []byte
}

var (
	_ EncodablePayload = calls{}
	_ Payload          = message{}
	_ Payload          = configUpdate{}
	_ Payload          = digest{}
)

type PayloadDigest struct {
	common.Hash

	Preimage Payload
}

func (d PayloadDigest) Digest() PayloadDigest {
	return d
}

type payload struct {
	address       common.Address
	chainID       *big.Int
	parentWallets []common.Address
}

func (p payload) Address() common.Address {
	return p.address
}

func (p payload) ChainID() *big.Int {
	return p.chainID
}

var eip712Domain = []ethcoder.TypedDataArgument{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}

func (p payload) domain() ethcoder.TypedDataDomain {
	return ethcoder.TypedDataDomain{
		Name:              "Sequence Wallet",
		Version:           "3",
		ChainID:           p.chainID,
		VerifyingContract: &p.address,
	}
}

func Calls(address common.Address, chainID *big.Int, calls_ []Call, space, nonce *big.Int, parentWallets ...[]common.Address) EncodablePayload {
	if chainID == nil {
		chainID = new(big.Int)
	}
	if len(calls_) >= 0x10000 {
		panic(fmt.Errorf("number of calls %v >= 2^16", len(calls_)))
	}
	if space == nil {
		space = new(big.Int)
	}
	if space.Sign() < 0 || space.Cmp(new(big.Int).Lsh(common.Big1, 160)) >= 0 {
		panic(fmt.Errorf("space %v is out of bounds [0, 2^160)", space))
	}
	if nonce == nil {
		nonce = new(big.Int)
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

	return calls{
		payload: payload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		calls: calls_,
		space: space,
		nonce: nonce,
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
		value = new(big.Int)
	}

	data := c.Data
	if data == nil {
		data = []byte{}
	}

	gasLimit := c.GasLimit
	if gasLimit == nil {
		gasLimit = new(big.Int)
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
	var call Call

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
		call.To, data = common.BytesToAddress(data), data[20:]
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
		dataSize := int(data[0])<<16 + int(data[1])<<8 + int(data[2])
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
	BehaviorOnErrorRevert                 = 1
	BehaviorOnErrorAbort                  = 2
)

type calls struct {
	payload

	calls        []Call
	space, nonce *big.Int
}

func (c calls) Digest() PayloadDigest {
	calls := make([]any, 0, len(c.calls))
	for _, call := range c.calls {
		calls = append(calls, any(call.asMap()))
	}

	wallets := make([]any, 0, len(c.parentWallets))
	for _, wallet := range c.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      c.domain(),
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
			"space":   c.space,
			"nonce":   c.nonce,
			"wallets": wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: c}
}

func (c calls) Encode(address common.Address) []byte {
	var flags byte
	var buffer bytes.Buffer

	if c.space.Sign() == 0 {
		flags |= 0x01
	} else {
		var space [20]byte
		mustWrite(&buffer, c.space.FillBytes(space[:]))
	}

	nonce := c.nonce.Bytes()
	flags |= byte(len(nonce)) << 1
	mustWrite(&buffer, nonce)

	calls := len(c.calls)
	if calls == 1 {
		flags |= 0x10
	} else if calls < 0x100 {
		mustWrite(&buffer, []byte{byte(calls)})
	} else {
		flags |= 0x20
		mustWrite(&buffer, []byte{byte(calls >> 8), byte(calls)})
	}

	for _, call := range c.calls {
		mustWrite(&buffer, call.encode(address))
	}

	data := make([]byte, 1+buffer.Len())
	data[0] = flags
	copy(data[1:], buffer.Bytes())
	return data
}

func DecodeCalls(address common.Address, chainID *big.Int, data []byte) (EncodablePayload, error) {
	var calls []Call
	var space, nonce *big.Int

	if len(data) < 1 {
		return nil, fmt.Errorf("no flags")
	}
	var flags byte
	flags, data = data[0], data[1:]

	if flags&0x01 == 0 {
		if len(data) < 20 {
			return nil, fmt.Errorf("no space")
		}
		space.SetBytes(data[:20])
		data = data[20:]
	}

	nonceSize := int(flags >> 1 & 0x07)
	if len(data) < nonceSize {
		return nil, fmt.Errorf("no nonce")
	}
	nonce.SetBytes(data[:nonceSize])
	data = data[nonceSize:]

	var calls_ int
	if flags&0x10 != 0 {
		calls_ = 1
	} else if flags&0x20 == 0 {
		if len(data) < 1 {
			return nil, fmt.Errorf("no number of calls")
		}
		calls_, data = int(data[0]), data[1:]
	} else {
		if len(data) < 2 {
			return nil, fmt.Errorf("no number of calls")
		}
		calls_, data = int(data[0])<<8+int(data[1]), data[2:]
	}

	for len(calls) < calls_ {
		var call Call
		var err error
		call, data, err = decodeCall(data, address)
		if err != nil {
			return nil, fmt.Errorf("unable to decode call %v: %w", len(calls), err)
		}

		calls = append(calls, call)
	}

	return Calls(address, chainID, calls, space, nonce), nil
}

func Message(address common.Address, chainID *big.Int, message_ []byte, parentWallets ...[]common.Address) Payload {
	if chainID == nil {
		chainID = new(big.Int)
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

	return message{
		payload: payload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		message: message_,
	}
}

type message struct {
	payload

	message []byte
}

func (m message) Digest() PayloadDigest {
	wallets := make([]any, 0, len(m.parentWallets))
	for _, wallet := range m.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      m.domain(),
		PrimaryType: "Message",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"Message": {
				{Name: "message", Type: "bytes"},
				{Name: "wallets", Type: "address[]"},
			},
		},
		Message: map[string]any{
			"message": m.message,
			"wallets": wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: m}
}

func ConfigUpdate(address common.Address, imageHash core.ImageHashable, parentWallets ...[]common.Address) Payload {
	if isNil(imageHash) {
		panic(fmt.Errorf("no config"))
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return configUpdate{
		payload: payload{
			address:       address,
			chainID:       new(big.Int),
			parentWallets: parentWallets[0],
		},

		imageHash: imageHash,
	}
}

type configUpdate struct {
	payload

	imageHash core.ImageHashable
}

func (u configUpdate) Digest() PayloadDigest {
	wallets := make([]any, 0, len(u.parentWallets))
	for _, wallet := range u.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      u.domain(),
		PrimaryType: "ConfigUpdate",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"ConfigUpdate": {
				{Name: "imageHash", Type: "bytes32"},
				{Name: "wallets", Type: "address[]"},
			},
		},
		Message: map[string]any{
			"imageHash": u.imageHash.ImageHash().Hash,
			"wallets":   wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: u}
}

func Digest(address common.Address, chainID *big.Int, digest_ common.Hash, parentWallets ...[]common.Address) Payload {
	if chainID == nil {
		chainID = new(big.Int)
	}
	if parentWallets == nil {
		parentWallets = append(parentWallets, nil)
	}
	if parentWallets[0] == nil {
		parentWallets[0] = []common.Address{}
	}

	return digest{
		payload: payload{
			address:       address,
			chainID:       chainID,
			parentWallets: parentWallets[0],
		},

		digest: digest_,
	}
}

type digest struct {
	payload

	digest common.Hash
}

func (d digest) Digest() PayloadDigest {
	wallets := make([]any, 0, len(d.parentWallets))
	for _, wallet := range d.parentWallets {
		wallets = append(wallets, any(wallet))
	}

	data := ethcoder.TypedData{
		Domain:      d.domain(),
		PrimaryType: "Digest",
		Types: map[string][]ethcoder.TypedDataArgument{
			"EIP712Domain": eip712Domain,
			"Digest": {
				{Name: "digest", Type: "bytes32"},
				{Name: "wallets", Type: "address[]"},
			},
		},
		Message: map[string]any{
			"digest":  d.digest,
			"wallets": wallets,
		},
	}

	digest, err := data.EncodeDigest()
	if err != nil {
		panic(err)
	}

	return PayloadDigest{Hash: common.Hash(digest), Preimage: d}
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
