// Sequence v1 core primitives
package v1

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
)

var isValidSignatureMagicValue = [4]byte{0x16, 0x26, 0xba, 0x7e}

var Core core.Core[*WalletConfig, core.Signature[*WalletConfig]] = v1Core{}

func init() {
	core.RegisterCore(Core)
}

type v1Core struct{}

func (v1Core) DecodeSignature(data []byte) (core.Signature[*WalletConfig], error) {
	threshold, err := readUint16(&data)
	if err != nil {
		return nil, fmt.Errorf("missing signature threshold: %w", err)
	}

	var leaves []signatureLeaf

	for len(data) != 0 {
		var leaf signatureLeaf

		switch data[0] {
		case byte(signatureLeafTypeECDSASignature):
			leaf, err = decodeSignatureTreeECDSASignatureLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode ecdsa signature leaf %v: %w", len(leaves), err)
			}

		case byte(signatureLeafTypeAddress):
			leaf, err = decodeSignatureTreeAddressLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode address leaf %v: %w", len(leaves), err)
			}

		case byte(signatureLeafTypeDynamicSignature):
			leaf, err = decodeSignatureTreeDynamicSignatureLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode dynamic signature leaf %v: %w", len(leaves), err)
			}

		default:
			return nil, fmt.Errorf("unknown signature leaf type %v for leaf %v", data[0], len(leaves))
		}

		leaves = append(leaves, leaf)
	}

	return &signature{threshold, leaves}, nil
}

func (v1Core) DecodeWalletConfig(object any) (*WalletConfig, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config must be an object")
	}

	threshold, ok := object_["threshold"]
	if !ok {
		return nil, fmt.Errorf(`missing required "threshold" property`)
	}
	threshold_, err := toUint16(threshold)
	if err != nil {
		return nil, fmt.Errorf("unable to convert threshold: %w", err)
	}

	signers, ok := object_["signers"]
	if !ok {
		return nil, fmt.Errorf(`missing required "signers" property`)
	}
	signers_, ok := signers.([]any)
	if !ok {
		signers__, ok := signers.([]map[string]any)
		if !ok {
			return nil, fmt.Errorf("signers must be an array")
		}

		signers_ = make([]any, 0, len(signers__))
		for _, signer := range signers__ {
			signers_ = append(signers_, signer)
		}
	}

	signers__ := make([]*WalletConfigSigner, 0, len(signers_))
	for i, signer := range signers_ {
		signer_, err := decodeWalletConfigSigner(signer)
		if err != nil {
			return nil, fmt.Errorf("signer %v is invalid: %w", i, err)
		}

		signers__ = append(signers__, signer_)
	}

	return &WalletConfig{
		Threshold_: threshold_,
		Signers_:   signers__,
	}, nil
}

type signerSignature struct {
	signer    common.Address
	type_     core.SignerSignatureType
	signature []byte
}

type signature struct {
	threshold uint16
	leaves    []signatureLeaf
}

func (s *signature) Threshold() uint16 {
	return s.threshold
}

func (s *signature) Checkpoint() uint32 {
	return 0
}

func (s *signature) Recover(ctx context.Context, digest core.Digest, wallet common.Address, chainID *big.Int, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	if chainID == nil {
		if provider == nil {
			return nil, nil, fmt.Errorf("provider is required if chain ID is not specified")
		}

		var err error
		chainID, err = provider.ChainID(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to get chain ID: %w", err)
		}
	}

	subdigest := digest.Subdigest(wallet, chainID)

	var signers []*WalletConfigSigner
	var total big.Int

	for i, leaf := range s.leaves {
		signer, weight, err := leaf.recover(ctx, subdigest, provider, signerSignatures[0])
		if err != nil {
			return nil, nil, fmt.Errorf("unable to recover signer for leaf %v: %w", i, err)
		}

		signers = append(signers, signer)
		total.Add(&total, weight)
	}

	return &WalletConfig{
		Threshold_: s.threshold,
		Signers_:   signers,
	}, &total, nil
}

func (s *signature) Data() ([]byte, error) {
	var buffer bytes.Buffer

	err := s.Write(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (s *signature) Write(writer io.Writer) error {
	err := binary.Write(writer, binary.BigEndian, s.threshold)
	if err != nil {
		return fmt.Errorf("unable to write signature threshold: %w", err)
	}

	for i, leaf := range s.leaves {
		err = leaf.write(writer)
		if err != nil {
			return fmt.Errorf("unable to write signature leaf %v: %w", i, err)
		}
	}

	return nil
}

func (s *signature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type signatureLeaf interface {
	recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (*WalletConfigSigner, *big.Int, error)
	write(writer io.Writer) error
}

type signatureLeafType byte

const (
	signatureLeafTypeECDSASignature   signatureLeafType = 0
	signatureLeafTypeAddress          signatureLeafType = 1
	signatureLeafTypeDynamicSignature signatureLeafType = 2
)

type eCDSASignatureType byte

const (
	eCDSASignatureTypeEIP712  eCDSASignatureType = 1
	eCDSASignatureTypeEthSign eCDSASignatureType = 2
)

func (t eCDSASignatureType) signerSignatureType() core.SignerSignatureType {
	switch t {
	case eCDSASignatureTypeEIP712:
		return core.SignerSignatureTypeEIP712
	case eCDSASignatureTypeEthSign:
		return core.SignerSignatureTypeEthSign
	default:
		panic(fmt.Sprintf("invalid ecdsa signature type %v", t))
	}
}

type signatureTreeECDSASignatureLeaf struct {
	weight    uint8
	type_     eCDSASignatureType
	signature [crypto.SignatureLength]byte
}

func decodeSignatureTreeECDSASignatureLeaf(data *[]byte) (*signatureTreeECDSASignatureLeaf, error) {
	next := *data

	if len(next) < 1+1+crypto.SignatureLength+1 {
		return nil, fmt.Errorf("insufficient data for ecdsa signature leaf")
	}

	var leafType byte
	var weight byte
	var signature []byte
	var signatureType byte
	leafType, weight, signature, signatureType, next = next[0], next[1], next[1+1:1+1+crypto.SignatureLength], next[1+1+crypto.SignatureLength], next[1+1+crypto.SignatureLength+1:]

	if leafType != byte(signatureLeafTypeECDSASignature) {
		return nil, fmt.Errorf("expected ecdsa signature leaf type %v, got %v", signatureLeafTypeECDSASignature, leafType)
	}

	switch signatureType {
	case byte(eCDSASignatureTypeEIP712), byte(eCDSASignatureTypeEthSign):
	default:
		return nil, fmt.Errorf("unknown ecdsa signature type %v", signatureType)
	}

	*data = next

	leaf := signatureTreeECDSASignatureLeaf{
		weight: weight,
		type_:  eCDSASignatureType(signatureType),
	}
	copy(leaf.signature[:], signature)
	return &leaf, nil
}

func (l *signatureTreeECDSASignatureLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (*WalletConfigSigner, *big.Int, error) {
	var address common.Address
	var err error
	switch l.type_ {
	case eCDSASignatureTypeEthSign:
		address, err = ecrecover(subdigest.EthSignSubdigest(), l.signature[:])
	default:
		address, err = ecrecover(subdigest, l.signature[:])
	}
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover ecdsa signature leaf: %w", err)
	}

	signerSignatures.Insert(address, core.SignerSignature{
		Subdigest: subdigest,
		Type:      l.type_.signerSignatureType(),
		Signature: l.signature[:],
	})

	return &WalletConfigSigner{
		Weight:  l.weight,
		Address: address,
	}, new(big.Int).SetUint64(uint64(l.weight)), nil
}

func (l *signatureTreeECDSASignatureLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeECDSASignature)})
	if err != nil {
		return fmt.Errorf("unable to write ecdsa signature leaf type: %w", err)
	}

	_, err = writer.Write([]byte{l.weight})
	if err != nil {
		return fmt.Errorf("unable to write ecdsa signature leaf weight: %w", err)
	}

	_, err = writer.Write(l.signature[:])
	if err != nil {
		return fmt.Errorf("unable to write ecdsa signature leaf signature: %w", err)
	}

	_, err = writer.Write([]byte{byte(l.type_)})
	if err != nil {
		return fmt.Errorf("unable to write ecdsa signature leaf signature type: %w", err)
	}

	return nil
}

type signatureTreeAddressLeaf struct {
	weight  uint8
	address common.Address
}

func decodeSignatureTreeAddressLeaf(data *[]byte) (*signatureTreeAddressLeaf, error) {
	next := *data

	if len(next) < 1+1+common.AddressLength {
		return nil, fmt.Errorf("insufficient data for address leaf")
	}

	var leafType byte
	var weight byte
	var address []byte
	leafType, weight, address, next = next[0], next[1], next[1+1:1+1+common.AddressLength], next[1+1+common.AddressLength:]

	if leafType != byte(signatureLeafTypeAddress) {
		return nil, fmt.Errorf("expected address leaf type %v, got %v", signatureLeafTypeAddress, leafType)
	}

	*data = next

	return &signatureTreeAddressLeaf{
		weight:  weight,
		address: common.BytesToAddress(address),
	}, nil
}

func (l *signatureTreeAddressLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (*WalletConfigSigner, *big.Int, error) {
	return &WalletConfigSigner{
		Weight:  l.weight,
		Address: l.address,
	}, new(big.Int), nil
}

func (l *signatureTreeAddressLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeAddress)})
	if err != nil {
		return fmt.Errorf("unable to write address leaf type: %w", err)
	}

	_, err = writer.Write([]byte{l.weight})
	if err != nil {
		return fmt.Errorf("unable to write address leaf weight: %w", err)
	}

	_, err = writer.Write(l.address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write address leaf address: %w", err)
	}

	return nil
}

type dynamicSignatureType byte

const (
	dynamicSignatureTypeEIP712  dynamicSignatureType = 1
	dynamicSignatureTypeEthSign dynamicSignatureType = 2
	dynamicSignatureTypeEIP1271 dynamicSignatureType = 3
)

func (t dynamicSignatureType) signerSignatureType() core.SignerSignatureType {
	switch t {
	case dynamicSignatureTypeEIP712:
		return core.SignerSignatureTypeEIP712
	case dynamicSignatureTypeEthSign:
		return core.SignerSignatureTypeEthSign
	case dynamicSignatureTypeEIP1271:
		return core.SignerSignatureTypeEIP1271
	default:
		panic(fmt.Sprintf("invalid dynamic signature type %v", t))
	}
}

type signatureTreeDynamicSignatureLeaf struct {
	weight    uint8
	address   common.Address
	type_     dynamicSignatureType
	signature []byte
}

func decodeSignatureTreeDynamicSignatureLeaf(data *[]byte) (*signatureTreeDynamicSignatureLeaf, error) {
	next := *data

	if len(next) < 1+1+common.AddressLength+2 {
		return nil, fmt.Errorf("insufficient data for dynamic signature leaf")
	}

	var leafType byte
	var weight byte
	var address []byte
	var signatureLength []byte
	leafType, weight, address, signatureLength, next = next[0], next[1], next[1+1:1+1+common.AddressLength], next[1+1+common.AddressLength:1+1+common.AddressLength+2], next[1+1+common.AddressLength+2:]

	if leafType != byte(signatureLeafTypeDynamicSignature) {
		return nil, fmt.Errorf("expected dynamic signature leaf type %v, got %v", signatureLeafTypeDynamicSignature, leafType)
	}

	length, err := readUint16(&signatureLength)
	if err != nil {
		return nil, fmt.Errorf("unable to read signature length: %w", err)
	}

	if len(next) < int(length) {
		return nil, fmt.Errorf("insufficient data for dynamic signature leaf")
	}

	if length == 0 {
		return nil, fmt.Errorf("missing dynamic signature type")
	}

	var signature []byte
	var signatureType byte
	signature, signatureType, next = next[:length-1], next[length-1], next[length:]

	switch signatureType {
	case byte(dynamicSignatureTypeEIP712), byte(dynamicSignatureTypeEthSign), byte(dynamicSignatureTypeEIP1271):
	default:
		return nil, fmt.Errorf("unknown dynamic signature type %v", signatureType)
	}

	*data = next

	return &signatureTreeDynamicSignatureLeaf{
		weight:    weight,
		address:   common.BytesToAddress(address),
		type_:     dynamicSignatureType(signatureType),
		signature: signature,
	}, nil
}

func (l *signatureTreeDynamicSignatureLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (*WalletConfigSigner, *big.Int, error) {
	switch l.type_ {
	case dynamicSignatureTypeEIP712, dynamicSignatureTypeEthSign:
		var address common.Address
		var err error
		switch l.type_ {
		case dynamicSignatureTypeEthSign:
			address, err = ecrecover(subdigest.EthSignSubdigest(), l.signature)
		default:
			address, err = ecrecover(subdigest, l.signature)
		}
		if err != nil {
			return nil, nil, fmt.Errorf("unable to recover dynamic signature leaf: %w", err)
		}

		if address != l.address {
			return nil, nil, fmt.Errorf("expected dynamic signature by %v, got dynamic signature by %v", l.address, address)
		}

		signerSignatures.Insert(l.address, core.SignerSignature{
			Subdigest: subdigest,
			Type:      l.type_.signerSignatureType(),
			Signature: l.signature,
		})

		return &WalletConfigSigner{
			Weight:  l.weight,
			Address: l.address,
		}, new(big.Int).SetUint64(uint64(l.weight)), nil

	case dynamicSignatureTypeEIP1271:
		if provider != nil {
			contract := ethcontract.NewContractCaller(l.address, contracts.IERC1271.ABI, provider)

			var results []any
			err := contract.Call(&bind.CallOpts{Context: ctx}, &results, "isValidSignature", subdigest.Hash, l.signature)
			if err != nil {
				return nil, nil, fmt.Errorf("unable to call isValidSignature on %v: %w", l.address, err)
			}

			if len(results) != 1 {
				return nil, nil, fmt.Errorf("expected single return value from isValidSignature, got %v", len(results))
			}

			magicValue, ok := results[0].([4]byte)
			if !ok {
				return nil, nil, fmt.Errorf("expected [4]byte return value from isValidSignature, got %T", results[0])
			}

			if magicValue != isValidSignatureMagicValue {
				return nil, nil, fmt.Errorf("isValidSignature returned %v, expected %v", hexutil.Encode(magicValue[:]), hexutil.Encode(isValidSignatureMagicValue[:]))
			}
		}

		signerSignatures.Insert(l.address, core.SignerSignature{
			Subdigest: subdigest,
			Type:      l.type_.signerSignatureType(),
			Signature: l.signature,
		})

		return &WalletConfigSigner{
			Weight:  l.weight,
			Address: l.address,
		}, new(big.Int).SetUint64(uint64(l.weight)), nil

	default:
		return nil, nil, fmt.Errorf("unknown dynamic signature type %v", l.type_)
	}
}

func (l *signatureTreeDynamicSignatureLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeDynamicSignature)})
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf type: %w", err)
	}

	_, err = writer.Write([]byte{l.weight})
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf weight: %w", err)
	}

	_, err = writer.Write(l.address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf address: %w", err)
	}

	err = binary.Write(writer, binary.BigEndian, uint16(len(l.signature)+1))
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf signature length: %w", err)
	}

	_, err = writer.Write(l.signature)
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf signature: %w", err)
	}

	_, err = writer.Write([]byte{byte(l.type_)})
	if err != nil {
		return fmt.Errorf("unable to write dynamic signature leaf signature type: %w", err)
	}

	return nil
}

type WalletConfigSigners []*WalletConfigSigner

func (s WalletConfigSigners) Len() int { return len(s) }
func (s WalletConfigSigners) Less(i, j int) bool {
	return s[i].Address.Hash().Big().Cmp(s[j].Address.Hash().Big()) < 0
}
func (s WalletConfigSigners) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type WalletConfig struct {
	Threshold_ uint16              `json:"threshold" toml:"threshold"`
	Signers_   WalletConfigSigners `json:"signers" toml:"signers"`
}

func (c *WalletConfig) Threshold() uint16 {
	return c.Threshold_
}

func (c *WalletConfig) Checkpoint() uint32 {
	return 0
}

func (c *WalletConfig) Signers() map[common.Address]uint16 {
	signers := map[common.Address]uint16{}

	for _, signer := range c.Signers_ {
		signers[signer.Address] = uint16(signer.Weight)
	}

	return signers
}

func (c *WalletConfig) IsUsable() error {
	if c.Threshold_ == 0 {
		return fmt.Errorf("threshold is 0")
	}

	var weight uint64
	for _, signer := range c.Signers_ {
		weight += uint64(signer.Weight)
		if weight >= uint64(c.Threshold_) {
			return nil
		}
	}

	return fmt.Errorf("threshold %v exceeds maximum weight %v", c.Threshold_, weight)
}

func (c *WalletConfig) ImageHash() core.ImageHash {
	imageHash := common.BigToHash(new(big.Int).SetUint64(uint64(c.Threshold_)))

	for _, signer := range c.Signers_ {
		preimage, err := ethcoder.AbiCoder(
			[]string{"bytes32", "uint8", "address"},
			[]any{imageHash, signer.Weight, signer.Address},
		)
		if err != nil {
			return core.ImageHash{}
		}

		imageHash = crypto.Keccak256Hash(preimage)
	}

	return core.ImageHash{Hash: imageHash, Preimage: c}
}

func (c *WalletConfig) BuildSignature(ctx context.Context, sign core.SigningFunction, validateSigningPower ...bool) (core.Signature[*WalletConfig], error) {
	isValid := false
	if len(validateSigningPower) > 0 {
		isValid = !validateSigningPower[0]
	}

	signerWeights := map[common.Address]*big.Int{}

	for _, signer := range c.Signers_ {
		weight := signerWeights[signer.Address]
		if weight == nil {
			weight = new(big.Int)
			signerWeights[signer.Address] = weight
		}

		weight.Add(weight, new(big.Int).SetUint64(uint64(signer.Weight)))
	}

	signerSignatureCh := make(chan signerSignature)

	signCtx, signCancel := context.WithCancel(ctx)
	defer signCancel()

	for signer := range signerWeights {
		go func(signer common.Address) {
			type_, signature, _ := sign(signCtx, signer)
			signerSignatureCh <- signerSignature{signer, type_, signature}
		}(signer)
	}

	signerSignatures := map[common.Address]signerSignature{}
	weight := new(big.Int)

	for range signerWeights {
		signerSignature := <-signerSignatureCh

		if signerSignature.signature != nil {
			signerSignatures[signerSignature.signer] = signerSignature

			weight.Add(weight, signerWeights[signerSignature.signer])
			if weight.Cmp(new(big.Int).SetUint64(uint64(c.Threshold_))) >= 0 {
				signCancel()
				isValid = true
			}
		}
	}

	if isValid {
		var leaves []signatureLeaf

		for _, signer := range c.Signers_ {
			leaves = append(leaves, signer.buildSignatureLeaf(signerSignatures))
		}

		return &signature{
			threshold: c.Threshold_,
			leaves:    leaves,
		}, nil
	} else {
		return nil, fmt.Errorf("not enough signers to build signature")
	}
}

func (c *WalletConfig) Clone() *WalletConfig {
	signers := make(WalletConfigSigners, len(c.Signers_))
	copy(signers, c.Signers_)
	return &WalletConfig{
		Threshold_: c.Threshold_,
		Signers_:   signers,
	}
}

type WalletConfigSigner struct {
	Weight  uint8          `json:"weight" toml:"weight"`
	Address common.Address `json:"address" toml:"address"`
}

func decodeWalletConfigSigner(object any) (*WalletConfigSigner, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config signer must be an object")
	}

	weight, ok := object_["weight"]
	if !ok {
		return nil, fmt.Errorf(`missing required "weight" property`)
	}
	weight_, err := toUint8(weight)
	if err != nil {
		return nil, fmt.Errorf("unable to convert weight: %w", err)
	}

	address, ok := object_["address"]
	if !ok {
		return nil, fmt.Errorf(`missing required "address" property`)
	}
	address_, ok := address.(string)
	if !ok {
		return nil, fmt.Errorf("address must be a string")
	}
	if !common.IsHexAddress(address_) {
		return nil, fmt.Errorf(`"%v" is not a valid address`, address_)
	}

	return &WalletConfigSigner{
		Weight:  weight_,
		Address: common.HexToAddress(address_),
	}, nil
}

func (s *WalletConfigSigner) buildSignatureLeaf(signerSignatures map[common.Address]signerSignature) signatureLeaf {
	signature, ok := signerSignatures[s.Address]
	if ok {
		switch signature.type_ {
		case core.SignerSignatureTypeEIP712:
			leaf := signatureTreeECDSASignatureLeaf{
				weight: s.Weight,
				type_:  eCDSASignatureTypeEIP712,
			}
			copy(leaf.signature[:], signature.signature)
			return &leaf

		case core.SignerSignatureTypeEthSign:
			leaf := signatureTreeECDSASignatureLeaf{
				weight: s.Weight,
				type_:  eCDSASignatureTypeEthSign,
			}
			copy(leaf.signature[:], signature.signature)
			return &leaf

		case core.SignerSignatureTypeEIP1271:
			return &signatureTreeDynamicSignatureLeaf{
				weight:    s.Weight,
				address:   s.Address,
				type_:     dynamicSignatureTypeEIP1271,
				signature: signature.signature,
			}

		default:
			panic(fmt.Sprintf("unknown signer signature type %v", signature.type_))
		}
	} else {
		return &signatureTreeAddressLeaf{
			weight:  s.Weight,
			address: s.Address,
		}
	}
}

func (s *WalletConfigSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"weight":  s.Weight,
		"address": s.Address.String(),
	})
}

func ecrecover(subdigest core.Subdigest, signature []byte) (common.Address, error) {
	if length := len(signature); length != crypto.SignatureLength {
		return common.Address{}, fmt.Errorf("invalid ecdsa signature length %v, expected length %v", length, crypto.SignatureLength)
	}

	var fixedSignature [crypto.SignatureLength]byte
	copy(fixedSignature[:], signature)
	fixedSignature[len(fixedSignature)-1] -= 27

	pubkey, err := crypto.SigToPub(subdigest.Bytes(), fixedSignature[:])
	if err != nil {
		return common.Address{}, fmt.Errorf("unable to recover ecdsa signature: %w", err)
	}

	return crypto.PubkeyToAddress(*pubkey), nil
}

func toUint8(number any) (uint8, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 || number > 0xff {
			return 0, fmt.Errorf("%v does not fit in a uint8", number)
		}
		return uint8(number), nil
	case float64:
		if number < 0 || number > 0xff {
			return 0, fmt.Errorf("%v does not fit in a uint8", number)
		}
		return uint8(number), nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint8", number)
	}
}

func toUint16(number any) (uint16, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 || number > 0xffff {
			return 0, fmt.Errorf("%v does not fit in a uint16", number)
		}
		return uint16(number), nil
	case float64:
		if number < 0 || number > 0xffff {
			return 0, fmt.Errorf("%v does not fit in a uint16", number)
		}
		return uint16(number), nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint16", number)
	}
}

func readUint16(data *[]byte) (uint16, error) {
	if len(*data) < 2 {
		return 0, fmt.Errorf("insufficient data for uint16")
	}

	value := binary.BigEndian.Uint16(*data)
	*data = (*data)[2:]
	return value, nil
}
