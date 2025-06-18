package v3

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"strconv"

	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/core"
	"github.com/0xsequence/go-sequence/eip6492"
)

const (
	addressLeafImageHashPrefix             = "Sequence signer:\n"
	nestedLeafImageHashPrefix              = "Sequence nested config:\n"
	sapientLeafImageHashPrefix             = "Sequence sapient config:\n"
	subdigestLeafImageHashPrefix           = "Sequence static digest:\n"
	anyAddressSubdigestLeafImageHashPrefix = "Sequence any address subdigest:\n"
)

var Core core.Core[*WalletConfig, core.Signature[*WalletConfig]] = v3Core{}

func init() {
	core.RegisterCore(Core)
}

type v3Core struct{}

func (v3Core) DecodeSignature(data []byte) (core.Signature[*WalletConfig], error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("missing signature flag")
	}

	signatureFlag := data[0]
	data = data[1:]

	signatureType := signatureFlag & 0x03
	if signatureType == 0x01 {
		return decodeChainedSignature(data)
	}

	noChainId := (signatureType & 0x02) == 0x02

	var checkpointer common.Address
	var checkpointerData []byte
	if signatureFlag&0x40 == 0x40 {
		if len(data) < 20 {
			return nil, fmt.Errorf("insufficient data for checkpointer address")
		}
		checkpointer = common.BytesToAddress(data[:20])
		data = data[20:]

		if len(data) < 3 {
			return nil, fmt.Errorf("insufficient data for checkpointer data size")
		}
		size := binary.BigEndian.Uint32([]byte{0, data[0], data[1], data[2]})
		data = data[3:]

		if len(data) < int(size) {
			return nil, fmt.Errorf("insufficient data for checkpointer data")
		}
		checkpointerData = data[:size]
		data = data[size:]
	}

	checkpointSize := (signatureFlag & 0x1c) >> 2
	checkpoint, err := readUintX(checkpointSize, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to read checkpoint: %w", err)
	}

	thresholdSize := ((signatureFlag & 0x20) >> 5) + 1
	threshold, err := readUintX(thresholdSize, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to read threshold: %w", err)
	}

	tree, err := decodeSignatureTree(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode signature tree: %w", err)
	}

	sig := &Signature{
		NoChainId:        noChainId,
		Threshold:        uint16(threshold),
		Checkpoint:       checkpoint,
		Tree:             tree,
		Checkpointer:     checkpointer,
		CheckpointerData: checkpointerData,
	}

	if noChainId {
		return &NoChainIDSignature{Signature: sig}, nil
	}
	return &RegularSignature{Signature: sig}, nil
}

func (v3Core) DecodeWalletConfig(object any) (*WalletConfig, error) {
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

	checkpoint, ok := object_["checkpoint"]
	if !ok {
		return nil, fmt.Errorf(`missing required "checkpoint" property`)
	}
	checkpoint_, err := toUint64(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to convert checkpoint: %w", err)
	}

	tree, ok := object_["tree"]
	if !ok {
		return nil, fmt.Errorf(`missing required "tree" property`)
	}
	tree_, err := DecodeWalletConfigTree(tree)
	if err != nil {
		return nil, fmt.Errorf("unable to decode wallet config tree: %w", err)
	}

	var checkpointer_ common.Address
	if checkpointerVal, ok := object_["checkpointer"]; ok && checkpointerVal != nil {
		checkpointerStr, ok := checkpointerVal.(string)
		if !ok {
			return nil, fmt.Errorf("checkpointer must be a string")
		}
		if checkpointerStr != "" {
			if !common.IsHexAddress(checkpointerStr) {
				return nil, fmt.Errorf(`invalid checkpointer address: "%v"`, checkpointerStr)
			}
			checkpointer_ = common.HexToAddress(checkpointerStr)
		}
	}

	return &WalletConfig{
		Threshold_:   threshold_,
		Checkpoint_:  checkpoint_,
		Tree:         tree_,
		Checkpointer: checkpointer_,
	}, nil
}

type Signature struct {
	NoChainId        bool
	Threshold        uint16
	Checkpoint       uint64
	Tree             signatureTree
	Checkpointer     common.Address
	CheckpointerData []byte
}

type RegularSignature struct {
	*Signature
}

func (s *RegularSignature) Threshold() uint16 {
	return s.Signature.Threshold
}

func (s *RegularSignature) Checkpoint() uint64 {
	return s.Signature.Checkpoint
}

func (s *RegularSignature) Recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	tree, weight, err := s.Tree.recover(ctx, payload, provider, signerSignatures[0])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover wallet config: %w", err)
	}

	return &WalletConfig{
		Threshold_:  s.Threshold(),
		Checkpoint_: s.Signature.Checkpoint,
		Tree:        tree,
	}, weight, nil
}

func (s *RegularSignature) Join(payload core.Payload, other core.Signature[*WalletConfig]) (core.Signature[*WalletConfig], error) {
	other_, ok := other.(*RegularSignature)
	if !ok {
		return nil, fmt.Errorf("expected regular signature, got %T", other)
	}

	if s.Threshold() != other_.Threshold() {
		return nil, fmt.Errorf("threshold mismatch: %v != %v", s.Threshold(), other_.Threshold())
	}

	if s.Checkpoint() != other_.Signature.Checkpoint {
		return nil, fmt.Errorf("checkpoint mismatch: %v != %v", s.Checkpoint(), other_.Checkpoint())
	}

	tree, err := s.Tree.join(other_.Tree)
	if err != nil {
		return nil, fmt.Errorf("unable to join signature trees: %w", err)
	}

	return &RegularSignature{&Signature{
		NoChainId:        s.NoChainId,
		Threshold:        s.Threshold(),
		Checkpoint:       s.Checkpoint(),
		Tree:             tree,
		Checkpointer:     s.Checkpointer,
		CheckpointerData: s.CheckpointerData,
	}}, nil
}

func (s *RegularSignature) Reduce(payload core.Payload) core.Signature[*WalletConfig] {
	return &RegularSignature{&Signature{
		NoChainId:        s.NoChainId,
		Threshold:        s.Threshold(),
		Checkpoint:       s.Checkpoint(),
		Tree:             s.Tree.reduce(),
		Checkpointer:     s.Checkpointer,
		CheckpointerData: s.CheckpointerData,
	}}
}

func (s *RegularSignature) Data() ([]byte, error) {
	return s.data(false, false)
}

func (s *RegularSignature) data(ignoreCheckpointer, ignoreCheckpointerData bool) ([]byte, error) {
	var buffer bytes.Buffer
	err := s.write(&buffer, ignoreCheckpointer, ignoreCheckpointerData)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (s *RegularSignature) Write(writer io.Writer) error {
	return s.write(writer, false, false)
}

func (s *RegularSignature) write(writer io.Writer, ignoreCheckpointer, ignoreCheckpointerData bool) error {
	checkpointSize := minBytesForUint64(s.Checkpoint())
	thresholdSize := minBytesForUint16(s.Threshold())

	if checkpointSize == 0 {
		checkpointSize = 1
	}
	if checkpointSize > 7 {
		return fmt.Errorf("checkpoint size %d exceeds maximum of 7 bytes", checkpointSize)
	}

	if thresholdSize == 0 {
		thresholdSize = 1
	}
	if thresholdSize > 2 {
		return fmt.Errorf("threshold size %d exceeds maximum of 2 bytes", thresholdSize)
	}

	flag := byte(0x00)
	flag |= (checkpointSize << 2)
	if thresholdSize == 2 {
		flag |= 0x20
	}
	if !ignoreCheckpointer && (s.Checkpointer != (common.Address{}) || len(s.CheckpointerData) > 0) {
		flag |= 0x40
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write signature flag: %w", err)
	}

	if !ignoreCheckpointer && (s.Checkpointer != (common.Address{}) || len(s.CheckpointerData) > 0) {
		_, err = writer.Write(s.Checkpointer.Bytes())
		if err != nil {
			return fmt.Errorf("unable to write checkpointer address: %w", err)
		}
		if !ignoreCheckpointerData {
			err = writeUint24(writer, uint32(len(s.CheckpointerData)))
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data length: %w", err)
			}
			_, err = writer.Write(s.CheckpointerData)
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data: %w", err)
			}
		}
	}

	err = writeUintX(writer, s.Checkpoint(), checkpointSize)
	if err != nil {
		return fmt.Errorf("unable to write checkpoint: %w", err)
	}

	err = writeUintX(writer, uint64(s.Threshold()), thresholdSize)
	if err != nil {
		return fmt.Errorf("unable to write threshold: %w", err)
	}

	err = s.Tree.write(writer)
	if err != nil {
		return fmt.Errorf("unable to write signature tree: %w", err)
	}

	return nil
}

func (s *RegularSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type NoChainIDSignature struct {
	*Signature
}

func (s *NoChainIDSignature) Threshold() uint16 {
	return s.Signature.Threshold
}

func (s *NoChainIDSignature) Checkpoint() uint64 {
	return s.Signature.Checkpoint
}

func (s *NoChainIDSignature) Recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	tree, weight, err := s.Tree.recover(ctx, payload, provider, signerSignatures[0])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover wallet config: %w", err)
	}

	return &WalletConfig{
		Threshold_:  s.Threshold(),
		Checkpoint_: s.Signature.Checkpoint,
		Tree:        tree,
	}, weight, nil
}

func (s *NoChainIDSignature) Join(payload core.Payload, other core.Signature[*WalletConfig]) (core.Signature[*WalletConfig], error) {
	other_, ok := other.(*NoChainIDSignature)
	if !ok {
		return nil, fmt.Errorf("expected no chain ID signature, got %T", other)
	}

	if s.Threshold() != other_.Threshold() {
		return nil, fmt.Errorf("threshold mismatch: %v != %v", s.Threshold(), other_.Threshold())
	}

	if s.Checkpoint() != other_.Signature.Checkpoint {
		return nil, fmt.Errorf("checkpoint mismatch: %v != %v", s.Checkpoint(), other_.Checkpoint())
	}

	tree, err := s.Tree.join(other_.Tree)
	if err != nil {
		return nil, fmt.Errorf("unable to join signature trees: %w", err)
	}

	return &NoChainIDSignature{&Signature{
		NoChainId:        s.NoChainId,
		Threshold:        s.Threshold(),
		Checkpoint:       s.Checkpoint(),
		Tree:             tree,
		Checkpointer:     s.Checkpointer,
		CheckpointerData: s.CheckpointerData,
	}}, nil
}

func (s *NoChainIDSignature) Reduce(payload core.Payload) core.Signature[*WalletConfig] {
	return &NoChainIDSignature{&Signature{
		NoChainId:        s.NoChainId,
		Threshold:        s.Threshold(),
		Checkpoint:       s.Checkpoint(),
		Tree:             s.Tree.reduce(),
		Checkpointer:     s.Checkpointer,
		CheckpointerData: s.CheckpointerData,
	}}
}

func (s *NoChainIDSignature) Data() ([]byte, error) {
	return s.data(false, false)
}

func (s *NoChainIDSignature) data(ignoreCheckpointer, ignoreCheckpointerData bool) ([]byte, error) {
	var buffer bytes.Buffer
	err := s.write(&buffer, ignoreCheckpointer, ignoreCheckpointerData)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (s *NoChainIDSignature) Write(writer io.Writer) error {
	return s.write(writer, false, false)
}

func (s *NoChainIDSignature) write(writer io.Writer, ignoreCheckpointer, ignoreCheckpointerData bool) error {
	checkpointSize := minBytesForUint64(s.Checkpoint())
	thresholdSize := minBytesForUint16(s.Threshold())

	if checkpointSize == 0 {
		checkpointSize = 1
	}
	if checkpointSize > 7 {
		return fmt.Errorf("checkpoint size %d exceeds maximum of 7 bytes", checkpointSize)
	}

	if thresholdSize == 0 {
		thresholdSize = 1
	}
	if thresholdSize > 2 {
		return fmt.Errorf("threshold size %d exceeds maximum of 2 bytes", thresholdSize)
	}

	flag := byte(0x02)
	flag |= (checkpointSize << 2)
	if thresholdSize == 2 {
		flag |= 0x20
	}
	if !ignoreCheckpointer && (s.Checkpointer != (common.Address{}) || len(s.CheckpointerData) > 0) {
		flag |= 0x40
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write signature flag: %w", err)
	}

	if !ignoreCheckpointer && (s.Checkpointer != (common.Address{}) || len(s.CheckpointerData) > 0) {
		_, err = writer.Write(s.Checkpointer.Bytes())
		if err != nil {
			return fmt.Errorf("unable to write checkpointer address: %w", err)
		}
		if !ignoreCheckpointerData {
			err = writeUint24(writer, uint32(len(s.CheckpointerData)))
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data length: %w", err)
			}
			_, err = writer.Write(s.CheckpointerData)
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data: %w", err)
			}
		}
	}

	err = writeUintX(writer, s.Checkpoint(), checkpointSize)
	if err != nil {
		return fmt.Errorf("unable to write checkpoint: %w", err)
	}

	err = writeUintX(writer, uint64(s.Threshold()), thresholdSize)
	if err != nil {
		return fmt.Errorf("unable to write threshold: %w", err)
	}

	err = s.Tree.write(writer)
	if err != nil {
		return fmt.Errorf("unable to write signature tree: %w", err)
	}

	return nil
}

func (s *NoChainIDSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type ChainedSignature []core.Signature[*WalletConfig]

func decodeChainedSignature(data []byte) (ChainedSignature, error) {
	if len(data) < 1 {
		return nil, fmt.Errorf("insufficient data for chained signature")
	}
	if data[0]&0x03 != 0x01 {
		return nil, fmt.Errorf("expected chained signature type 0x01, got %v", data[0])
	}
	data = data[1:]

	var signature ChainedSignature
	for len(data) != 0 {
		length, err := readUint24(&data)
		if err != nil {
			return nil, fmt.Errorf("unable to read subsignature %v length: %w", len(signature), err)
		}

		if len(data) < int(length) {
			return nil, fmt.Errorf("insufficient data for subsignature %v", len(signature))
		}

		subsignature, err := Core.DecodeSignature(data[:length])
		if err != nil {
			return nil, fmt.Errorf("unable to decode subsignature %v: %w", len(signature), err)
		}
		data = data[length:]

		if len(signature) != 0 {
			currentCheckpoint := subsignature.Checkpoint()
			previousCheckpoint := signature[len(signature)-1].Checkpoint()
			if currentCheckpoint >= previousCheckpoint {
				return nil, fmt.Errorf(
					"subsignature %v's checkpoint %v is not less than subsignature %v's checkpoint %v",
					len(signature), currentCheckpoint, len(signature)-1, previousCheckpoint,
				)
			}
		}

		signature = append(signature, subsignature)
	}

	return signature, nil
}

func (s ChainedSignature) Threshold() uint16 {
	return s[len(s)-1].Threshold()
}

func (s ChainedSignature) Checkpoint() uint64 {
	return s[len(s)-1].Checkpoint()
}

func (s ChainedSignature) Checkpointer() common.Address {
	if len(s) == 0 {
		return common.Address{}
	}
	switch signature := s[len(s)-1].(type) {
	case *RegularSignature:
		return signature.Checkpointer
	case *NoChainIDSignature:
		return signature.Checkpointer
	case ChainedSignature:
		return signature.Checkpointer()
	case *ChainedSignature:
		return signature.Checkpointer()
	default:
		panic(fmt.Errorf("unknown signature type %T", signature))
	}
}

func (s ChainedSignature) CheckpointerData() []byte {
	if len(s) == 0 {
		return nil
	}
	switch signature := s[len(s)-1].(type) {
	case *RegularSignature:
		return signature.CheckpointerData
	case *NoChainIDSignature:
		return signature.CheckpointerData
	case ChainedSignature:
		return signature.CheckpointerData()
	case *ChainedSignature:
		return signature.CheckpointerData()
	default:
		panic(fmt.Errorf("unknown signature type %T", signature))
	}
}

func (s ChainedSignature) Recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	var config *WalletConfig
	var weight *big.Int

	for i, subsignature := range s {
		var err error
		config, weight, err = subsignature.Recover(ctx, payload, provider, signerSignatures...)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to recover subsignature %v: %w", i, err)
		}

		if weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) < 0 {
			return nil, nil, fmt.Errorf("recovered weight %v for subsignature %v does not meet required threshold %v", weight, i, config.Threshold())
		}

		payload = NewConfigUpdatePayload(payload.Digest().Address, nil, config.ImageHash().Hash)
	}

	return config, weight, nil
}

func (s ChainedSignature) Join(payload core.Payload, other core.Signature[*WalletConfig]) (core.Signature[*WalletConfig], error) {
	return nil, fmt.Errorf("chained signatures do not support joining")
}

func (s ChainedSignature) Reduce(payload core.Payload) core.Signature[*WalletConfig] {
	subsignatures := make(ChainedSignature, 0, len(s))
	for _, subsignature := range s {
		subsignatures = append(subsignatures, subsignature.Reduce(payload))
	}
	return subsignatures
}

func (s ChainedSignature) Data() ([]byte, error) {
	return s.data(false, false)
}

func (s ChainedSignature) data(ignoreCheckpointer, ignoreCheckpointerData bool) ([]byte, error) {
	var buffer bytes.Buffer
	err := s.write(&buffer, ignoreCheckpointer, ignoreCheckpointerData)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (s ChainedSignature) Write(writer io.Writer) error {
	return s.write(writer, false, false)
}

func (s ChainedSignature) write(writer io.Writer, ignoreCheckpointer, ignoreCheckpointerData bool) error {
	flag := byte(0x01)
	checkpointer := s.Checkpointer()
	checkpointerData := s.CheckpointerData()

	if !ignoreCheckpointer && (checkpointer != (common.Address{}) || len(checkpointerData) != 0) {
		flag |= 0x40
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write chained signature type: %w", err)
	}

	if !ignoreCheckpointer && (checkpointer != (common.Address{}) || len(checkpointerData) != 0) {
		_, err := writer.Write(checkpointer.Bytes())
		if err != nil {
			return fmt.Errorf("unable to write checkpointer address: %w", err)
		}
		if !ignoreCheckpointerData {
			err := writeUint24(writer, uint32(len(checkpointerData)))
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data size: %w", err)
			}
			_, err = writer.Write(checkpointerData)
			if err != nil {
				return fmt.Errorf("unable to write checkpointer data: %w", err)
			}
		}
	}

	for i, subsignature := range s {
		var data []byte
		switch subsignature := subsignature.(type) {
		case *RegularSignature:
			data, err = subsignature.data(i == len(s)-1, true)
		case *NoChainIDSignature:
			data, err = subsignature.data(i == len(s)-1, true)
		case ChainedSignature:
			data, err = subsignature.data(i == len(s)-1, true)
		case *ChainedSignature:
			data, err = subsignature.data(i == len(s)-1, true)
		}
		if err != nil {
			return fmt.Errorf("unable to encode subsignature %v: %w", i, err)
		}
		err = writeUint24(writer, uint32(len(data)))
		if err != nil {
			return fmt.Errorf("unable to write subsignature %v length: %w", i, err)
		}
		_, err = writer.Write(data)
		if err != nil {
			return fmt.Errorf("unable to write subsignature %v: %w", i, err)
		}
	}

	return nil
}

func (s ChainedSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type signatureTree interface {
	recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error)
	reduce() signatureTree
	join(other signatureTree) (signatureTree, error)
	reduceImageHash() (core.ImageHash, error)
	write(writer io.Writer) error
}

func decodeSignatureTree(data []byte) (signatureTree, error) {
	var tree signatureTree
	for len(data) != 0 {
		leaf, err := decodeSignatureTreeLeaf(&data)
		if err != nil {
			return nil, err
		}
		if tree == nil {
			tree = leaf
		} else {
			tree = &signatureTreeNode{left: tree, right: leaf}
		}
	}
	return tree, nil
}

func decodeSignatureTreeLeaf(data *[]byte) (signatureTree, error) {
	if len(*data) == 0 {
		return nil, fmt.Errorf("insufficient data for signature tree leaf")
	}

	firstByte := (*data)[0]
	*data = (*data)[1:]
	flag := (firstByte & 0xf0) >> 4

	switch flag {
	case FLAG_SIGNATURE_HASH:
		return decodeSignatureHashLeaf(firstByte, data)
	case FLAG_ADDRESS:
		return decodeAddressLeaf(firstByte, data)
	case FLAG_SIGNATURE_ERC1271:
		return decodeSignatureERC1271Leaf(firstByte, data)
	case FLAG_NODE:
		return decodeNodeLeaf(data)
	case FLAG_BRANCH:
		return decodeBranchLeaf(firstByte, data)
	case FLAG_SUBDIGEST:
		return decodeSubdigestLeaf(data)
	case FLAG_NESTED:
		return decodeNestedLeaf(firstByte, data)
	case FLAG_SIGNATURE_ETH_SIGN:
		return decodeSignatureEthSignLeaf(firstByte, data)
	case FLAG_SIGNATURE_ANY_ADDRESS_SUBDIGEST:
		return decodeAnyAddressSubdigestLeaf(data)
	case FLAG_SIGNATURE_SAPIENT:
		return decodeSignatureSapientLeaf(firstByte, data)
	case FLAG_SIGNATURE_SAPIENT_COMPACT:
		return decodeSignatureSapientCompactLeaf(firstByte, data)
	default:
		return nil, fmt.Errorf("unknown signature leaf flag %v", flag)
	}
}

const (
	FLAG_SIGNATURE_HASH                  = 0x00
	FLAG_ADDRESS                         = 0x01
	FLAG_SIGNATURE_ERC1271               = 0x02
	FLAG_NODE                            = 0x03
	FLAG_BRANCH                          = 0x04
	FLAG_SUBDIGEST                       = 0x05
	FLAG_NESTED                          = 0x06
	FLAG_SIGNATURE_ETH_SIGN              = 0x07
	FLAG_SIGNATURE_ANY_ADDRESS_SUBDIGEST = 0x08
	FLAG_SIGNATURE_SAPIENT               = 0x09
	FLAG_SIGNATURE_SAPIENT_COMPACT       = 0x0A
)

type signatureTreeNode struct {
	left, right signatureTree
}

func (n *signatureTreeNode) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	left, leftWeight, err := n.left.recover(ctx, payload, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover left subtree: %w", err)
	}

	right, rightWeight, err := n.right.recover(ctx, payload, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover right subtree: %w", err)
	}

	return &WalletConfigTreeNode{left, right}, new(big.Int).Add(leftWeight, rightWeight), nil
}

func (n *signatureTreeNode) reduce() signatureTree {
	if imageHash, err := n.reduceImageHash(); err == nil {
		return signatureTreeNodeLeaf{imageHash}
	}
	return &signatureTreeNode{
		left:  n.left.reduce(),
		right: n.right.reduce(),
	}
}

func (n *signatureTreeNode) join(other signatureTree) (signatureTree, error) {
	switch other := other.(type) {
	case *signatureTreeNode:
		var err error
		var left, right signatureTree
		if other.left != nil {
			left, err = n.left.join(other.left)
			if err != nil {
				return nil, fmt.Errorf("unable to join left subtree: %w", err)
			}
		}
		if other.right != nil {
			right, err = n.right.join(other.right)
			if err != nil {
				return nil, fmt.Errorf("unable to join right subtree: %w", err)
			}
		}
		return &signatureTreeNode{left, right}, nil
	case signatureTreeNodeLeaf:
		return n, nil
	default:
		return nil, fmt.Errorf("unable to join signature tree node with %T", other)
	}
}

func (n *signatureTreeNode) reduceImageHash() (core.ImageHash, error) {
	if leftImageHash, err := n.left.reduceImageHash(); err == nil {
		if rightImageHash, err := n.right.reduceImageHash(); err == nil {
			node := WalletConfigTreeNode{
				Left:  WalletConfigTreeNodeLeaf{leftImageHash},
				Right: WalletConfigTreeNodeLeaf{rightImageHash},
			}
			return node.ImageHash(), nil
		}
	}
	return core.ImageHash{}, fmt.Errorf("node might have signing power")
}

func (n *signatureTreeNode) write(writer io.Writer) error {
	err := n.left.write(writer)
	if err != nil {
		return fmt.Errorf("unable to write left subtree: %w", err)
	}

	_, isNode := n.right.(*signatureTreeNode)
	if isNode {
		var buffer bytes.Buffer
		err = n.right.write(&buffer)
		if err != nil {
			return fmt.Errorf("unable to encode right subtree: %w", err)
		}
		size := uint64(buffer.Len())
		sizeBytes := minBytesFor(size)
		if sizeBytes > 15 {
			return fmt.Errorf("right subtree too large: %d bytes", size)
		}
		flag := byte(FLAG_BRANCH<<4) | byte(sizeBytes)
		_, err = writer.Write([]byte{flag})
		if err != nil {
			return fmt.Errorf("unable to write branch flag: %w", err)
		}
		err = writeUintX(writer, size, sizeBytes)
		if err != nil {
			return fmt.Errorf("unable to write right subtree length: %w", err)
		}
		_, err = writer.Write(buffer.Bytes())
		if err != nil {
			return fmt.Errorf("unable to write right subtree: %w", err)
		}
	} else {
		err = n.right.write(writer)
		if err != nil {
			return fmt.Errorf("unable to write right subtree: %w", err)
		}
	}
	return nil
}

type signatureTreeSignatureHashLeaf struct {
	Weight  uint8
	R       [32]byte
	S       [32]byte
	YParity bool
}

func decodeSignatureHashLeaf(firstByte byte, data *[]byte) (*signatureTreeSignatureHashLeaf, error) {
	weight := firstByte & 0x0f
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 64 {
		return nil, fmt.Errorf("insufficient data for signature")
	}

	r := [32]byte(*data)
	s := [32]byte((*data)[32:])
	yParity := s[0]&0x80 != 0
	s[0] &^= 0x80
	*data = (*data)[64:]

	return &signatureTreeSignatureHashLeaf{
		Weight:  weight,
		R:       r,
		S:       s,
		YParity: yParity,
	}, nil
}

func (l *signatureTreeSignatureHashLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	var v byte
	if l.YParity {
		v = 28
	} else {
		v = 27
	}
	signature := bytes.Join([][]byte{l.R[:], l.S[:], {v}}, nil)

	address, err := core.Ecrecover(payload.Digest().Hash, signature)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover signature: %w", err)
	}

	signerSignatures.Insert2(address, core.SignerSignature{
		Type:      core.SignerSignatureTypeEIP712,
		Signature: signature,
	})

	return &WalletConfigTreeAddressLeaf{
		Weight:  l.Weight,
		Address: address,
	}, new(big.Int).SetUint64(uint64(l.Weight)), nil
}

func (l *signatureTreeSignatureHashLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeSignatureHashLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeSignatureHashLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("signature hash leaf has signing power")
}

func (l *signatureTreeSignatureHashLeaf) write(writer io.Writer) error {
	var flag byte
	var weightBytes []byte
	flag = FLAG_SIGNATURE_HASH << 4

	if l.Weight > 0 && l.Weight <= 15 {
		flag |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write signature hash leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write weight byte: %w", err)
		}
	}

	_, err = writer.Write(l.R[:])
	if err != nil {
		return fmt.Errorf("unable to write R: %w", err)
	}

	s := l.S
	if l.YParity {
		s[0] |= 0x80
	}
	_, err = writer.Write(s[:])
	if err != nil {
		return fmt.Errorf("unable to write S: %w", err)
	}

	return nil
}

type signatureTreeAddressLeaf struct {
	Weight  uint8
	Address common.Address
}

func decodeAddressLeaf(firstByte byte, data *[]byte) (*signatureTreeAddressLeaf, error) {
	weight := firstByte & 0x0f
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 20 {
		return nil, fmt.Errorf("insufficient data for address")
	}

	address := common.BytesToAddress((*data)[:20])
	*data = (*data)[20:]

	return &signatureTreeAddressLeaf{
		Weight:  weight,
		Address: address,
	}, nil
}

func (l *signatureTreeAddressLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	return &WalletConfigTreeAddressLeaf{
		Weight:  l.Weight,
		Address: l.Address,
	}, new(big.Int), nil
}

func (l *signatureTreeAddressLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeAddressLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeAddressLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("address leaf has signing power")
}

func (l *signatureTreeAddressLeaf) write(writer io.Writer) error {
	flag := byte(FLAG_ADDRESS << 4)
	var weightBytes []byte

	if l.Weight > 0 && l.Weight <= 15 {
		flag |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write address leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	_, err = writer.Write(l.Address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write address: %w", err)
	}

	return nil
}

type signatureTreeSignatureERC1271Leaf struct {
	Weight    uint8
	Address   common.Address
	Signature []byte
}

func decodeSignatureERC1271Leaf(firstByte byte, data *[]byte) (*signatureTreeSignatureERC1271Leaf, error) {
	weight := firstByte & 0x03
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 20 {
		return nil, fmt.Errorf("insufficient data for address")
	}
	addr := common.BytesToAddress((*data)[:20])
	*data = (*data)[20:]

	sizeSize := (firstByte & 0x0c) >> 2
	size, err := readUintX(uint8(sizeSize), data)
	if err != nil {
		return nil, fmt.Errorf("unable to read signature size: %w", err)
	}

	if len(*data) < int(size) {
		return nil, fmt.Errorf("insufficient data for signature")
	}
	signature := make([]byte, size)
	copy(signature, (*data)[:size])
	*data = (*data)[size:]

	return &signatureTreeSignatureERC1271Leaf{
		Weight:    weight,
		Address:   addr,
		Signature: signature,
	}, nil
}

func (l *signatureTreeSignatureERC1271Leaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	effectiveWeight := l.Weight
	signature := l.Signature

	if provider != nil {
		isValid, err := eip6492.ValidateEIP6492Offchain(ctx, provider, l.Address, payload.Digest().Hash, signature, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to validate ERC-1271 signature: %w", err)
		}
		if !isValid {
			return nil, nil, fmt.Errorf("invalid ERC-1271 signature for %v", l.Address)
		}
		if eip6492.IsEIP6492Signature(signature) {
			_, _, signature, err = eip6492.DecodeEIP6492Signature(signature)
			if err != nil {
				return nil, nil, fmt.Errorf("unable to decode EIP-6492 signature: %w", err)
			}
		}
	} else {
		effectiveWeight = 0
	}

	signerSignatures.Insert2(l.Address, core.SignerSignature{
		Type:      core.SignerSignatureTypeEIP1271,
		Signature: signature,
	})

	return &WalletConfigTreeAddressLeaf{
		Weight:  l.Weight,
		Address: l.Address,
	}, new(big.Int).SetUint64(uint64(effectiveWeight)), nil
}

func (l *signatureTreeSignatureERC1271Leaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeSignatureERC1271Leaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeSignatureERC1271Leaf) reduceImageHash() (core.ImageHash, error) {
	if l.Weight == 0 {
		return (&WalletConfigTreeAddressLeaf{
			Weight:  l.Weight,
			Address: l.Address,
		}).ImageHash(), nil
	}
	return core.ImageHash{}, fmt.Errorf("ERC-1271 signature might have signing power")
}

func (l *signatureTreeSignatureERC1271Leaf) write(writer io.Writer) error {
	sigLen := uint64(len(l.Signature))

	sizeSize := minBytesFor(sigLen)
	if sizeSize == 0 {
		// Ensure at least 1 byte for the length, even if 0
		sizeSize = 1
	}
	if sizeSize > 3 {
		return fmt.Errorf("signature length %d requires %d bytes, exceeds maximum of 3", sigLen, sizeSize)
	}

	var firstByte byte = FLAG_SIGNATURE_ERC1271<<4 | (byte(sizeSize) << 2)
	var weightBytes []byte

	if l.Weight >= 1 && l.Weight <= 3 {
		firstByte |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{firstByte})
	if err != nil {
		return fmt.Errorf("unable to write ERC-1271 leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	_, err = writer.Write(l.Address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write address: %w", err)
	}

	err = writeUintX(writer, sigLen, sizeSize)
	if err != nil {
		return fmt.Errorf("unable to write signature length: %w", err)
	}

	_, err = writer.Write(l.Signature)
	if err != nil {
		return fmt.Errorf("unable to write signature: %w", err)
	}

	return nil
}

type signatureTreeNodeLeaf struct{ core.ImageHash }

func decodeNodeLeaf(data *[]byte) (signatureTreeNodeLeaf, error) {
	if len(*data) < 32 {
		return signatureTreeNodeLeaf{}, fmt.Errorf("insufficient data for node leaf")
	}
	hash := common.BytesToHash((*data)[:32])
	*data = (*data)[32:]
	return signatureTreeNodeLeaf{core.ImageHash{Hash: hash}}, nil
}

func (l signatureTreeNodeLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	leaf := WalletConfigTreeNodeLeaf{l.ImageHash}
	leaf.Node.Preimage = &leaf
	return leaf, new(big.Int), nil
}

func (l signatureTreeNodeLeaf) reduce() signatureTree {
	return l
}

func (l signatureTreeNodeLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l signatureTreeNodeLeaf) reduceImageHash() (core.ImageHash, error) {
	return l.ImageHash, nil
}

func (l signatureTreeNodeLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{FLAG_NODE << 4})
	if err != nil {
		return fmt.Errorf("unable to write node leaf type: %w", err)
	}

	_, err = writer.Write(l.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write node hash: %w", err)
	}

	return nil
}

func decodeBranchLeaf(firstByte byte, data *[]byte) (signatureTree, error) {
	lengthSize := firstByte & 0x0f

	length, err := readUintX(lengthSize, data)
	if err != nil {
		return nil, fmt.Errorf("unable to read branch length: %w", err)
	}

	if len(*data) < int(length) {
		return nil, fmt.Errorf("insufficient data for branch")
	}

	branchData := (*data)[:length]
	*data = (*data)[length:]
	return decodeSignatureTree(branchData)
}

type signatureTreeSubdigestLeaf struct{ common.Hash }

func decodeSubdigestLeaf(data *[]byte) (signatureTreeSubdigestLeaf, error) {
	if len(*data) < 32 {
		return signatureTreeSubdigestLeaf{}, fmt.Errorf("insufficient data for subdigest leaf")
	}
	hash := common.BytesToHash((*data)[:32])
	*data = (*data)[32:]
	return signatureTreeSubdigestLeaf{Hash: hash}, nil
}

func (l signatureTreeSubdigestLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	if payload.Digest().Hash == l.Hash {
		return WalletConfigTreeSubdigestLeaf{l.Hash}, new(big.Int).Set(maxUint256), nil
	} else {
		return WalletConfigTreeSubdigestLeaf{l.Hash}, new(big.Int), nil
	}
}

func (l signatureTreeSubdigestLeaf) reduce() signatureTree {
	return l
}

func (l signatureTreeSubdigestLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l signatureTreeSubdigestLeaf) reduceImageHash() (core.ImageHash, error) {
	return WalletConfigTreeSubdigestLeaf{l.Hash}.ImageHash(), nil
}

func (l signatureTreeSubdigestLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{FLAG_SUBDIGEST << 4})

	if err != nil {
		return fmt.Errorf("unable to write subdigest leaf type: %w", err)
	}

	_, err = writer.Write(l.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write subdigest: %w", err)
	}

	return nil
}

type signatureTreeNestedLeaf struct {
	Weight    uint8
	Threshold uint16
	Tree      signatureTree
}

func decodeNestedLeaf(firstByte byte, data *[]byte) (*signatureTreeNestedLeaf, error) {
	weight := firstByte & 0x0f
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	threshold, err := readUint16(data)
	if err != nil {
		return nil, fmt.Errorf("unable to read nested threshold: %w", err)
	}

	length, err := readUint24(data)
	if err != nil {
		return nil, fmt.Errorf("unable to read nested tree length: %w", err)
	}

	if len(*data) < int(length) {
		return nil, fmt.Errorf("insufficient data for nested tree")
	}

	treeData := (*data)[:length]
	*data = (*data)[length:]
	tree, err := decodeSignatureTree(treeData)
	if err != nil {
		return nil, fmt.Errorf("unable to decode nested tree: %w", err)
	}

	return &signatureTreeNestedLeaf{weight, threshold, tree}, nil
}

func (l *signatureTreeNestedLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	tree, weight, err := l.Tree.recover(ctx, payload, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover nested tree: %w", err)
	}

	if weight.Cmp(new(big.Int).SetUint64(uint64(l.Threshold))) >= 0 {
		return &WalletConfigTreeNestedLeaf{
			Weight:    l.Weight,
			Threshold: l.Threshold,
			Tree:      tree,
		}, new(big.Int).SetUint64(uint64(l.Weight)), nil
	}
	return &WalletConfigTreeNestedLeaf{
		Weight:    l.Weight,
		Threshold: l.Threshold,
		Tree:      tree,
	}, new(big.Int), nil
}

func (l *signatureTreeNestedLeaf) reduce() signatureTree {
	return &signatureTreeNestedLeaf{
		Weight:    l.Weight,
		Threshold: l.Threshold,
		Tree:      l.Tree.reduce(),
	}
}

func (l *signatureTreeNestedLeaf) join(other signatureTree) (signatureTree, error) {
	switch other := other.(type) {
	case *signatureTreeNestedLeaf:
		tree, err := l.Tree.join(other.Tree)
		if err != nil {
			return nil, fmt.Errorf("unable to join nested trees: %w", err)
		}
		return &signatureTreeNestedLeaf{
			Weight:    l.Weight,
			Threshold: l.Threshold,
			Tree:      tree,
		}, nil
	default:
		return nil, fmt.Errorf("unable to join nested leaf with %T", other)
	}
}

func (l *signatureTreeNestedLeaf) reduceImageHash() (core.ImageHash, error) {
	if l.Weight == 0 {
		if imageHash, err := l.Tree.reduceImageHash(); err == nil {
			return (&WalletConfigTreeNestedLeaf{
				Weight:    l.Weight,
				Threshold: l.Threshold,
				Tree:      WalletConfigTreeNodeLeaf{imageHash},
			}).ImageHash(), nil
		}
	}
	return core.ImageHash{}, fmt.Errorf("nested leaf might have signing power")
}

func (l *signatureTreeNestedLeaf) write(writer io.Writer) error {
	flag := byte(FLAG_NESTED << 4)

	var weightBytes []byte
	if l.Weight <= 3 && l.Weight > 0 {
		flag |= (l.Weight << 2)
	} else {
		weightBytes = []byte{l.Weight}
	}

	var thresholdBytes []byte
	if l.Threshold <= 3 && l.Threshold > 0 {
		flag |= byte(l.Threshold)
	} else {
		thresholdBytes = make([]byte, 2)
		binary.BigEndian.PutUint16(thresholdBytes, l.Threshold)
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write nested leaf flag: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	if len(thresholdBytes) > 0 {
		_, err = writer.Write(thresholdBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic threshold: %w", err)
		}
	}

	var buffer bytes.Buffer
	err = l.Tree.write(&buffer)
	if err != nil {
		return fmt.Errorf("unable to encode nested tree: %w", err)
	}

	err = writeUint24(writer, uint32(buffer.Len()))
	if err != nil {
		return fmt.Errorf("unable to write nested tree length: %w", err)
	}

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write nested tree: %w", err)
	}

	return nil
}

type signatureTreeSignatureEthSignLeaf struct {
	Weight  uint8
	R       [32]byte
	S       [32]byte
	YParity bool
}

func decodeSignatureEthSignLeaf(firstByte byte, data *[]byte) (*signatureTreeSignatureEthSignLeaf, error) {
	weight := firstByte & 0x0f
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 64 {
		return nil, fmt.Errorf("insufficient data for signature")
	}

	r := [32]byte(*data)
	s := [32]byte((*data)[32:])
	yParity := s[0]&0x80 != 0
	s[0] &^= 0x80
	*data = (*data)[64:]

	return &signatureTreeSignatureEthSignLeaf{
		Weight:  weight,
		R:       r,
		S:       s,
		YParity: yParity,
	}, nil
}

func (l *signatureTreeSignatureEthSignLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	var v byte
	if l.YParity {
		v = 28
	} else {
		v = 27
	}
	signature := bytes.Join([][]byte{l.R[:], l.S[:], {v}}, nil)

	address, err := core.Ecrecover(core.EthereumSignedMessage(payload.Digest().Bytes()), signature)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover eth sign signature: %w", err)
	}

	signerSignatures.Insert2(address, core.SignerSignature{
		Type:      core.SignerSignatureTypeEthSign,
		Signature: signature,
	})

	return &WalletConfigTreeAddressLeaf{
		Weight:  l.Weight,
		Address: address,
	}, new(big.Int).SetUint64(uint64(l.Weight)), nil
}

func (l *signatureTreeSignatureEthSignLeaf) reduce() signatureTree { return l }
func (l *signatureTreeSignatureEthSignLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}
func (l *signatureTreeSignatureEthSignLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("eth sign signature has signing power")
}

func (l *signatureTreeSignatureEthSignLeaf) write(writer io.Writer) error {
	flag := byte(FLAG_SIGNATURE_ETH_SIGN << 4)
	var weightBytes []byte

	if l.Weight > 0 && l.Weight <= 15 {
		flag |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{flag})
	if err != nil {
		return fmt.Errorf("unable to write eth sign leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	_, err = writer.Write(l.R[:])
	if err != nil {
		return fmt.Errorf("unable to write R: %w", err)
	}

	s := l.S
	if l.YParity {
		s[0] |= 0x80
	}
	_, err = writer.Write(s[:])
	if err != nil {
		return fmt.Errorf("unable to write S: %w", err)
	}
	return nil
}

type signatureTreeAnyAddressSubdigestLeaf struct{ common.Hash }

func decodeAnyAddressSubdigestLeaf(data *[]byte) (*signatureTreeAnyAddressSubdigestLeaf, error) {
	if len(*data) < 32 {
		return nil, fmt.Errorf("insufficient data for any address subdigest leaf")
	}
	hash := common.BytesToHash((*data)[:32])
	*data = (*data)[32:]
	return &signatureTreeAnyAddressSubdigestLeaf{Hash: hash}, nil
}

func (l *signatureTreeAnyAddressSubdigestLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	payload_, err := PayloadWithAddress(payload, common.Address{})
	if err != nil {
		return nil, nil, fmt.Errorf("unable to override payload address: %w", err)
	}

	if payload_.Digest().Hash == l.Hash {
		return WalletConfigTreeAnyAddressSubdigestLeaf{l.Hash}, new(big.Int).Set(maxUint256), nil
	} else {
		return WalletConfigTreeAnyAddressSubdigestLeaf{l.Hash}, new(big.Int), nil
	}
}

func (l *signatureTreeAnyAddressSubdigestLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeAnyAddressSubdigestLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeAnyAddressSubdigestLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{Hash: crypto.Keccak256Hash([]byte(anyAddressSubdigestLeafImageHashPrefix), l.Bytes())}, nil
}

func (l *signatureTreeAnyAddressSubdigestLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{FLAG_SIGNATURE_ANY_ADDRESS_SUBDIGEST << 4})
	if err != nil {
		return fmt.Errorf("unable to write any address subdigest leaf type: %w", err)
	}

	_, err = writer.Write(l.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write subdigest: %w", err)
	}

	return nil
}

type signatureTreeSapientLeaf struct {
	Weight    uint8
	Address   common.Address
	Signature []byte
}

func decodeSignatureSapientLeaf(firstByte byte, data *[]byte) (*signatureTreeSapientLeaf, error) {
	weight := firstByte & 0x03
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 20 {
		return nil, fmt.Errorf("insufficient data for address")
	}
	addr := common.BytesToAddress((*data)[:20])
	*data = (*data)[20:]

	sizeSize := (firstByte & 0x0c) >> 2
	size, err := readUintX(uint8(sizeSize), data)
	if err != nil {
		return nil, fmt.Errorf("unable to read signature size: %w", err)
	}

	if len(*data) < int(size) {
		return nil, fmt.Errorf("insufficient data for signature")
	}
	signature := make([]byte, size)
	copy(signature, (*data)[:size])
	*data = (*data)[size:]

	return &signatureTreeSapientLeaf{
		Weight:    weight,
		Address:   addr,
		Signature: signature,
	}, nil
}

func (l *signatureTreeSapientLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	imageHash, err := RecoverSapientSignature(ctx, l.Address, payload, l.Signature, provider)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover sapient signature: %w", err)
	}

	signerSignatures.InsertSapient(l.Address, imageHash.Hash, core.SignerSignature{
		Signer:    l.Address,
		Type:      core.SignerSignatureTypeSapient,
		Signature: l.Signature,
	})

	return &WalletConfigTreeSapientSignerLeaf{
		Weight:     l.Weight,
		Address:    l.Address,
		ImageHash_: imageHash,
	}, new(big.Int).SetUint64(uint64(l.Weight)), nil
}

func (l *signatureTreeSapientLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeSapientLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeSapientLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("sapient leaf might have signing power")
}

func (l *signatureTreeSapientLeaf) write(writer io.Writer) error {
	sigLen := uint64(len(l.Signature))
	sizeSize := minBytesFor(sigLen)
	if sizeSize == 0 {
		// Ensure at least 1 byte for the length, even if 0
		sizeSize = 1
	}
	if sizeSize > 3 {
		return fmt.Errorf("signature length %d requires %d bytes, exceeds maximum of 3", sigLen, sizeSize)
	}

	var firstByte byte = FLAG_SIGNATURE_SAPIENT<<4 | (byte(sizeSize) << 2)
	var weightBytes []byte

	if l.Weight >= 1 && l.Weight <= 3 {
		firstByte |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{firstByte})
	if err != nil {
		return fmt.Errorf("unable to write Sapient leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	_, err = writer.Write(l.Address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write address: %w", err)
	}

	err = writeUintX(writer, sigLen, sizeSize)
	if err != nil {
		return fmt.Errorf("unable to write signature length: %w", err)
	}

	_, err = writer.Write(l.Signature)
	if err != nil {
		return fmt.Errorf("unable to write signature: %w", err)
	}

	return nil
}

type signatureTreeSapientCompactLeaf struct {
	Weight    uint8
	Address   common.Address
	Signature []byte
}

func decodeSignatureSapientCompactLeaf(firstByte byte, data *[]byte) (*signatureTreeSapientCompactLeaf, error) {
	weight := firstByte & 0x03
	if weight == 0 {
		if len(*data) < 1 {
			return nil, fmt.Errorf("insufficient data for dynamic weight")
		}
		weight = (*data)[0]
		*data = (*data)[1:]
	}

	if len(*data) < 20 {
		return nil, fmt.Errorf("insufficient data for address")
	}
	addr := common.BytesToAddress((*data)[:20])
	*data = (*data)[20:]

	sizeSize := (firstByte & 0x0c) >> 2
	size, err := readUintX(uint8(sizeSize), data)
	if err != nil {
		return nil, fmt.Errorf("unable to read signature size: %w", err)
	}

	if len(*data) < int(size) {
		return nil, fmt.Errorf("insufficient data for signature")
	}
	signature := make([]byte, size)
	copy(signature, (*data)[:size])
	*data = (*data)[size:]

	return &signatureTreeSapientCompactLeaf{
		Weight:    weight,
		Address:   addr,
		Signature: signature,
	}, nil
}

func (l *signatureTreeSapientCompactLeaf) recover(ctx context.Context, payload core.Payload, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	imageHash, err := RecoverSapientSignatureCompact(ctx, l.Address, payload, l.Signature, provider)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover compact sapient signature: %w", err)
	}

	signerSignatures.InsertSapient(l.Address, imageHash.Hash, core.SignerSignature{
		Signer:    l.Address,
		Type:      core.SignerSignatureTypeSapientCompact,
		Signature: l.Signature,
	})

	return &WalletConfigTreeSapientSignerLeaf{
		Weight:     l.Weight,
		Address:    l.Address,
		ImageHash_: imageHash,
	}, new(big.Int).SetUint64(uint64(l.Weight)), nil
}

func (l *signatureTreeSapientCompactLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeSapientCompactLeaf) join(other signatureTree) (signatureTree, error) {
	return l, nil
}

func (l *signatureTreeSapientCompactLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("sapient compact leaf might have signing power")
}

func (l *signatureTreeSapientCompactLeaf) write(writer io.Writer) error {
	sigLen := uint64(len(l.Signature))
	sizeSize := minBytesFor(sigLen)
	if sizeSize == 0 {
		// Ensure at least 1 byte for the length, even if 0
		sizeSize = 1
	}
	if sizeSize > 3 {
		return fmt.Errorf("signature length %d requires %d bytes, exceeds maximum of 3", sigLen, sizeSize)
	}

	var firstByte byte = FLAG_SIGNATURE_SAPIENT_COMPACT<<4 | (byte(sizeSize) << 2)
	var weightBytes []byte

	if l.Weight >= 1 && l.Weight <= 3 {
		firstByte |= l.Weight
	} else {
		weightBytes = []byte{l.Weight}
	}

	_, err := writer.Write([]byte{firstByte})
	if err != nil {
		return fmt.Errorf("unable to write SapientCompact leaf type: %w", err)
	}

	if len(weightBytes) > 0 {
		_, err = writer.Write(weightBytes)
		if err != nil {
			return fmt.Errorf("unable to write dynamic weight: %w", err)
		}
	}

	_, err = writer.Write(l.Address.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write address: %w", err)
	}

	err = writeUintX(writer, sigLen, sizeSize)
	if err != nil {
		return fmt.Errorf("unable to write signature length: %w", err)
	}

	_, err = writer.Write(l.Signature)
	if err != nil {
		return fmt.Errorf("unable to write signature: %w", err)
	}

	return nil
}

type WalletConfig struct {
	Threshold_ uint16 `json:"threshold" toml:"threshold"`
	// WARNING: contract code is uint64
	Checkpoint_  uint64           `json:"checkpoint" toml:"checkpoint"`
	Tree         WalletConfigTree `json:"tree" toml:"tree"`
	Checkpointer common.Address   `json:"checkpointer,omitempty" toml:"checkpointer,omitempty"`
}

func (c *WalletConfig) Threshold() uint16 {
	return c.Threshold_
}

func (c *WalletConfig) Checkpoint() uint64 {
	return c.Checkpoint_
}

func (c *WalletConfig) Signers() map[common.Address]uint16 {
	signers := map[common.Address]uint16{}
	c.Tree.readSignersIntoMap(signers)
	return signers
}

func (c *WalletConfig) SignersWeight(signers []common.Address) uint16 {
	signersMap := make(map[common.Address]uint16)
	for _, signer := range signers {
		signersMap[signer] = 0
	}
	return uint16(c.Tree.unverifiedWeight(signersMap).Uint64())
}

func (c *WalletConfig) IsComplete() bool {
	return c.Tree.isComplete()
}

func (c *WalletConfig) IsUsable() error {
	if c.Threshold_ == 0 {
		return fmt.Errorf("threshold is 0")
	}

	threshold := new(big.Int).SetUint64(uint64(c.Threshold_))
	weight := c.Tree.maxWeight()
	if threshold.Cmp(weight) > 0 {
		return fmt.Errorf("threshold %v exceeds maximum weight %v", threshold, weight)
	}

	return nil
}

func (c *WalletConfig) ImageHash() core.ImageHash {
	treeHash := c.Tree.ImageHash().Bytes()

	thresholdBytes := common.BigToHash(new(big.Int).SetUint64(uint64(c.Threshold_))).Bytes()
	checkpointBytes := common.BigToHash(new(big.Int).SetUint64(c.Checkpoint_)).Bytes()

	checkpointerBytes := make([]byte, 32)
	copy(checkpointerBytes[12:], c.Checkpointer.Bytes())

	root := crypto.Keccak256Hash(treeHash, thresholdBytes)

	root = crypto.Keccak256Hash(root.Bytes(), checkpointBytes)

	root = crypto.Keccak256Hash(root.Bytes(), checkpointerBytes)

	return core.ImageHash{
		Hash:     root,
		Preimage: c,
	}
}

func (c *WalletConfig) BuildSubdigestSignature(noChainID bool) (core.Signature[*WalletConfig], error) {
	signerSignatures := make(map[common.Address]signerSignature)
	tree := c.Tree.buildSignatureTree(signerSignatures)

	// Construct a NoChainIDSignature with the wallet config's parameters
	if noChainID {
		return &NoChainIDSignature{
			Signature: &Signature{
				NoChainId:        noChainID,
				Threshold:        c.Threshold_,
				Checkpoint:       c.Checkpoint_,
				Tree:             tree,
				Checkpointer:     common.Address{},
				CheckpointerData: nil,
			},
		}, nil
	}

	return &RegularSignature{
		Signature: &Signature{
			NoChainId:        noChainID,
			Threshold:        c.Threshold_,
			Checkpoint:       c.Checkpoint_,
			Tree:             tree,
			Checkpointer:     common.Address{},
			CheckpointerData: nil,
		},
	}, nil
}

func (c *WalletConfig) BuildRegularSignature(ctx context.Context, sign core.SigningFunction, validateSigningPower bool, checkpointerData ...[]byte) (core.Signature[*WalletConfig], error) {
	var isValid bool
	configSigners := c.Signers()

	signCtx, signCancel := context.WithCancel(ctx)
	defer signCancel()

	signerSignatureCh := core.SigningOrchestrator(signCtx, configSigners, sign)
	signerSignatures := map[common.Address]signerSignature{}
	signedSigners := map[common.Address]uint16{}

	for range configSigners {
		signerSig := <-signerSignatureCh
		if signerSig.Error != nil {
			signCancel()
			return nil, fmt.Errorf("signer %s signing error: %w", signerSig.Signer.Hex(), signerSig.Error)
		}

		if signerSig.Signature != nil {
			signerSignatures[signerSig.Signer] = signerSignature{signerSig.Signer, signerSig.Type, signerSig.Signature}
			signedSigners[signerSig.Signer] = 0

			weight := c.Tree.unverifiedWeight(signedSigners)
			if weight.Cmp(new(big.Int).SetUint64(uint64(c.Threshold_))) >= 0 {
				signCancel()
				isValid = true
			}
		}
	}

	if !isValid && validateSigningPower {
		return nil, fmt.Errorf("not enough signers to build regular signature")
	}

	var cpData []byte
	if len(checkpointerData) > 0 {
		cpData = checkpointerData[0]
	}

	return &RegularSignature{&Signature{
		NoChainId:        false,
		Threshold:        c.Threshold_,
		Checkpoint:       c.Checkpoint_,
		Tree:             c.Tree.buildSignatureTree(signerSignatures),
		Checkpointer:     c.Checkpointer,
		CheckpointerData: cpData,
	}}, nil
}

func (c *WalletConfig) BuildNoChainIDSignature(ctx context.Context, sign core.SigningFunction, validateSigningPower bool, checkpointerData ...[]byte) (core.Signature[*WalletConfig], error) {
	var isValid bool
	configSigners := c.Signers()

	signCtx, signCancel := context.WithCancel(ctx)
	defer signCancel()

	signerSignatureCh := core.SigningOrchestrator(signCtx, configSigners, sign)
	signerSignatures := map[common.Address]signerSignature{}
	signedSigners := map[common.Address]uint16{}

	for range configSigners {
		signerSig := <-signerSignatureCh
		if signerSig.Error != nil {
			signCancel()
			return nil, fmt.Errorf("signer %s signing error: %w", signerSig.Signer.Hex(), signerSig.Error)
		}

		if signerSig.Signature != nil {
			signerSignatures[signerSig.Signer] = signerSignature{signerSig.Signer, signerSig.Type, signerSig.Signature}
			signedSigners[signerSig.Signer] = 0

			weight := c.Tree.unverifiedWeight(signedSigners)
			if weight.Cmp(new(big.Int).SetUint64(uint64(c.Threshold_))) >= 0 {
				signCancel()
				isValid = true
			}
		}
	}

	if !isValid && validateSigningPower {
		return nil, fmt.Errorf("not enough signers to build no chain ID signature")
	}

	var cpData []byte
	if len(checkpointerData) > 0 {
		cpData = checkpointerData[0]
	}

	return &NoChainIDSignature{&Signature{
		NoChainId:        true,
		Threshold:        c.Threshold_,
		Checkpoint:       c.Checkpoint_,
		Tree:             c.Tree.buildSignatureTree(signerSignatures),
		Checkpointer:     c.Checkpointer,
		CheckpointerData: cpData,
	}}, nil
}

type signerSignature struct {
	signer    common.Address
	type_     core.SignerSignatureType
	signature []byte
}

type WalletConfigTree interface {
	core.ImageHashable

	isComplete() bool
	maxWeight() *big.Int
	readSignersIntoMap(signers map[common.Address]uint16)
	unverifiedWeight(signers map[common.Address]uint16) *big.Int
	buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree
}

func DecodeWalletConfigTree(object any) (WalletConfigTree, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config tree must be an object")
	}

	if typ, ok := object_["type"].(string); ok {
		switch typ {
		case "signer":
			return decodeWalletConfigTreeAddressLeaf(object_)
		case "sapient-signer":
			return decodeWalletConfigTreeSapientSignerLeaf(object_)
		case "subdigest":
			return decodeWalletConfigTreeSubdigestLeaf(object_)
		case "any-address-subdigest":
			return decodeWalletConfigTreeAnyAddressSubdigestLeaf(object_)
		case "nested":
			return decodeWalletConfigTreeNestedLeaf(object_)
		case "node":
			hash, ok := object_["hash"].(string)
			if !ok {
				return nil, fmt.Errorf("missing or invalid 'hash' for node type")
			}
			hashBytes, err := hexutil.Decode(hash)
			if err != nil || len(hashBytes) != 32 {
				return nil, fmt.Errorf("invalid hash for node type: %v", hash)
			}
			return &WalletConfigTreeNodeLeaf{Node: core.ImageHash{Hash: common.BytesToHash(hashBytes)}}, nil
		default:
			return nil, fmt.Errorf("unknown wallet config tree type: %s", typ)
		}
	}

	if hasKeys(object_, []string{"left", "right"}) {
		return decodeWalletConfigTreeNode(object_)
	} else if hasKeys(object_, []string{"weight", "address"}) {
		return decodeWalletConfigTreeAddressLeaf(object_)
	} else if hasKeys(object_, []string{"node"}) {
		return decodeWalletConfigTreeNodeLeaf(object_)
	} else if hasKeys(object_, []string{"weight", "threshold", "tree"}) {
		return decodeWalletConfigTreeNestedLeaf(object_)
	} else if hasKeys(object_, []string{"subdigest"}) {
		return decodeWalletConfigTreeSubdigestLeaf(object_)
	} else if hasKeys(object_, []string{"weight", "address", "imageHash"}) {
		return decodeWalletConfigTreeSapientSignerLeaf(object_)
	} else if hasKeys(object_, []string{"digest"}) {
		return decodeWalletConfigTreeAnyAddressSubdigestLeaf(object_)
	}

	return nil, fmt.Errorf("unknown wallet config tree type: %v", object_)
}

func WalletConfigTreeNodes(nodes ...WalletConfigTree) WalletConfigTree {
	if len(nodes) == 0 {
		return nil
	} else if len(nodes) == 1 {
		return nodes[0]
	} else {
		numberOfNodes := len(nodes)/2 + len(nodes)%2
		for numberOfNodes > 1 {
			newNodes := make([]WalletConfigTree, numberOfNodes)
			for i := range numberOfNodes {
				var left, right WalletConfigTree
				if i*2 < len(nodes) {
					left = nodes[i*2]
				}
				if i*2+1 < len(nodes) {
					right = nodes[i*2+1]
				}

				if right == nil {
					newNodes[i] = left
				} else {
					newNodes[i] = &WalletConfigTreeNode{
						Left:  left,
						Right: right,
					}
				}
			}

			nodes = newNodes
			numberOfNodes = len(nodes)/2 + len(nodes)%2
		}

		return &WalletConfigTreeNode{
			Left:  nodes[0],
			Right: nodes[1],
		}
	}
}

type WalletConfigTreeNode struct {
	Left  WalletConfigTree `json:"left" toml:"left"`
	Right WalletConfigTree `json:"right" toml:"right"`
}

func decodeWalletConfigTreeNode(object any) (*WalletConfigTreeNode, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config tree node must be an object")
	}

	left, ok := object_["left"]
	if !ok {
		return nil, fmt.Errorf(`missing required "left" property`)
	}
	left_, err := DecodeWalletConfigTree(left)
	if err != nil {
		return nil, fmt.Errorf("unable to decode left wallet config tree: %w", err)
	}

	right, ok := object_["right"]
	if !ok {
		return nil, fmt.Errorf(`missing required "right" property`)
	}
	right_, err := DecodeWalletConfigTree(right)
	if err != nil {
		return nil, fmt.Errorf("unable to decode right wallet config tree: %w", err)
	}

	return &WalletConfigTreeNode{Left: left_, Right: right_}, nil
}

func (n *WalletConfigTreeNode) ImageHash() core.ImageHash {
	return core.ImageHash{
		Hash: crypto.Keccak256Hash(
			n.Left.ImageHash().Bytes(),
			n.Right.ImageHash().Bytes(),
		),
		Preimage: n,
	}
}

func (n *WalletConfigTreeNode) isComplete() bool {
	return n.Left.isComplete() && n.Right.isComplete()
}

func (n *WalletConfigTreeNode) maxWeight() *big.Int {
	return new(big.Int).Add(n.Left.maxWeight(), n.Right.maxWeight())
}

func (n *WalletConfigTreeNode) readSignersIntoMap(signers map[common.Address]uint16) {
	n.Left.readSignersIntoMap(signers)
	n.Right.readSignersIntoMap(signers)
}

func (n *WalletConfigTreeNode) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	return new(big.Int).Add(n.Left.unverifiedWeight(signers), n.Right.unverifiedWeight(signers))
}

func (n *WalletConfigTreeNode) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	leftTree := n.Left.buildSignatureTree(signerSignatures)
	rightTree := n.Right.buildSignatureTree(signerSignatures)
	return &signatureTreeNode{left: leftTree, right: rightTree}
}

type WalletConfigTreeAddressLeaf struct {
	Weight  uint8          `json:"weight" toml:"weight"`
	Address common.Address `json:"address" toml:"address"`
}

func decodeWalletConfigTreeAddressLeaf(object any) (*WalletConfigTreeAddressLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config tree address leaf must be an object")
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

	return &WalletConfigTreeAddressLeaf{
		Weight:  weight_,
		Address: common.HexToAddress(address_),
	}, nil
}

func (l *WalletConfigTreeAddressLeaf) ImageHash() core.ImageHash {
	weight := new(big.Int).SetUint64(uint64(l.Weight))
	weightBytes := make([]byte, 32)
	weight.FillBytes(weightBytes)

	hash := crypto.Keccak256Hash(
		[]byte(addressLeafImageHashPrefix),
		l.Address.Bytes(),
		weightBytes,
	)
	return core.ImageHash{Hash: hash, Preimage: l}
}

func (l *WalletConfigTreeAddressLeaf) isComplete() bool {
	return true
}

func (l *WalletConfigTreeAddressLeaf) maxWeight() *big.Int {
	return new(big.Int).SetUint64(uint64(l.Weight))
}

func (l *WalletConfigTreeAddressLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
	signers[l.Address] = uint16(l.Weight)
}

func (l *WalletConfigTreeAddressLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	if _, ok := signers[l.Address]; ok {
		return new(big.Int).SetUint64(uint64(l.Weight))
	}
	return new(big.Int)
}

func (l *WalletConfigTreeAddressLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	if signature, ok := signerSignatures[l.Address]; ok {
		switch signature.type_ {
		case core.SignerSignatureTypeEIP712, core.SignerSignatureTypeEthSign:
			r := [32]byte(signature.signature[0:32])
			s := [32]byte(signature.signature[32:64])
			v := signature.signature[64]

			var yParity bool
			if v == 0 || v == 27 {
				yParity = false
			} else if v == 1 || v == 28 {
				yParity = true
			} else if v > 35 {
				// EIP-155: V = 35 + 2*chainID + parity
				yParity = ((v - 35) % 2) == 1
			}

			if signature.type_ == core.SignerSignatureTypeEIP712 {
				return &signatureTreeSignatureHashLeaf{
					Weight:  l.Weight,
					R:       r,
					S:       s,
					YParity: yParity,
				}
			} else {
				return &signatureTreeSignatureEthSignLeaf{
					Weight:  l.Weight,
					R:       r,
					S:       s,
					YParity: yParity,
				}
			}
		case core.SignerSignatureTypeEIP1271:
			return &signatureTreeSignatureERC1271Leaf{
				Weight:    l.Weight,
				Address:   l.Address,
				Signature: signature.signature,
			}
		default:
			panic(fmt.Sprintf("unknown signer signature type %v", signature.type_))
		}
	}
	return &signatureTreeAddressLeaf{Weight: l.Weight, Address: l.Address}
}

type WalletConfigTreeNodeLeaf struct {
	Node core.ImageHash `json:"node" toml:"node"`
}

func decodeWalletConfigTreeNodeLeaf(object any) (WalletConfigTreeNodeLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf("wallet config tree node leaf must be an object")
	}

	node, ok := object_["node"]
	if !ok {
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf(`missing required "node" property`)
	}
	node_, ok := node.(string)
	if !ok {
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf("node must be a string")
	}
	node__, err := hexutil.Decode(node_)
	if err != nil {
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf(`"%v" is not valid hex`, node_)
	}
	if len(node__) != common.HashLength {
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf("expected hash of length %v, got %v", common.HashLength, len(node__))
	}

	leaf := WalletConfigTreeNodeLeaf{core.ImageHash{Hash: common.BytesToHash(node__)}}
	leaf.Node.Preimage = &leaf
	return leaf, nil
}

func (l WalletConfigTreeNodeLeaf) ImageHash() core.ImageHash {
	return l.Node
}

func (l WalletConfigTreeNodeLeaf) isComplete() bool {
	return false
}

func (l WalletConfigTreeNodeLeaf) maxWeight() *big.Int {
	return new(big.Int)
}

func (l WalletConfigTreeNodeLeaf) readSignersIntoMap(signers map[common.Address]uint16) {}

func (l WalletConfigTreeNodeLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	return new(big.Int)
}

func (l WalletConfigTreeNodeLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return signatureTreeNodeLeaf{l.Node}
}

type WalletConfigTreeNestedLeaf struct {
	Weight    uint8            `json:"weight" toml:"weight"`
	Threshold uint16           `json:"threshold" toml:"threshold"`
	Tree      WalletConfigTree `json:"tree" toml:"tree"`
}

func decodeWalletConfigTreeNestedLeaf(object any) (*WalletConfigTreeNestedLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config tree nested leaf must be an object")
	}

	weight, ok := object_["weight"]
	if !ok {
		return nil, fmt.Errorf(`missing required "weight" property`)
	}
	weight_, err := toUint8(weight)
	if err != nil {
		return nil, fmt.Errorf("unable to convert weight: %w", err)
	}

	threshold, ok := object_["threshold"]
	if !ok {
		return nil, fmt.Errorf(`missing required "threshold" property`)
	}
	threshold_, err := toUint16(threshold)
	if err != nil {
		return nil, fmt.Errorf("unable to convert threshold: %w", err)
	}

	tree, ok := object_["tree"]
	if !ok {
		return nil, fmt.Errorf(`missing required "tree" property`)
	}
	tree_, err := DecodeWalletConfigTree(tree)
	if err != nil {
		return nil, fmt.Errorf("unable to decode nested tree: %w", err)
	}

	return &WalletConfigTreeNestedLeaf{
		Weight:    weight_,
		Threshold: threshold_,
		Tree:      tree_,
	}, nil
}

func (l *WalletConfigTreeNestedLeaf) ImageHash() core.ImageHash {
	imageHash := l.Tree.ImageHash()
	threshold := common.BigToHash(new(big.Int).SetUint64(uint64(l.Threshold)))
	weight := common.BigToHash(new(big.Int).SetUint64(uint64(l.Weight)))

	return core.ImageHash{
		Hash: crypto.Keccak256Hash(
			[]byte(nestedLeafImageHashPrefix),
			imageHash.Bytes(),
			threshold.Bytes(),
			weight.Bytes(),
		),
		Preimage: l,
	}
}

func (l *WalletConfigTreeNestedLeaf) isComplete() bool {
	return l.Tree.isComplete()
}

func (l *WalletConfigTreeNestedLeaf) maxWeight() *big.Int {
	if l.Tree.maxWeight().Cmp(new(big.Int).SetUint64(uint64(l.Threshold))) >= 0 {
		return new(big.Int).SetUint64(uint64(l.Weight))
	}
	return new(big.Int)
}

func (l *WalletConfigTreeNestedLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
	l.Tree.readSignersIntoMap(signers)
}

func (l *WalletConfigTreeNestedLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	if l.Tree.unverifiedWeight(signers).Cmp(new(big.Int).SetUint64(uint64(l.Threshold))) >= 0 {
		return new(big.Int).SetUint64(uint64(l.Weight))
	}
	return new(big.Int)
}

func (l *WalletConfigTreeNestedLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return &signatureTreeNestedLeaf{
		Weight:    l.Weight,
		Threshold: l.Threshold,
		Tree:      l.Tree.buildSignatureTree(signerSignatures),
	}
}

type WalletConfigTreeSubdigestLeaf struct {
	Subdigest common.Hash `json:"subdigest" toml:"subdigest"`
}

func decodeWalletConfigTreeSubdigestLeaf(object any) (WalletConfigTreeSubdigestLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("wallet config tree subdigest leaf must be an object")
	}

	digest, ok := object_["digest"]
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf(`missing required "digest" property`)
	}

	digest_, ok := digest.(string)
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("digest must be a string")
	}

	digestBytes, err := hexutil.Decode(digest_)
	if err != nil {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf(`"%v" is not valid hex`, digest_)
	}

	if len(digestBytes) != common.HashLength {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("expected hash of length %v, got %v", common.HashLength, len(digestBytes))
	}

	return WalletConfigTreeSubdigestLeaf{common.BytesToHash(digestBytes)}, nil
}

func (l WalletConfigTreeSubdigestLeaf) ImageHash() core.ImageHash {
	return core.ImageHash{
		Hash: crypto.Keccak256Hash(
			[]byte(subdigestLeafImageHashPrefix),
			l.Subdigest.Bytes(),
		),
		Preimage: &l,
	}
}

func (l WalletConfigTreeSubdigestLeaf) isComplete() bool {
	return true
}

func (l WalletConfigTreeSubdigestLeaf) maxWeight() *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeSubdigestLeaf) readSignersIntoMap(signers map[common.Address]uint16) {}

func (l WalletConfigTreeSubdigestLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeSubdigestLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return signatureTreeSubdigestLeaf{l.Subdigest}
}

type WalletConfigTreeSapientSignerLeaf struct {
	Weight     uint8          `json:"weight" toml:"weight"`
	Address    common.Address `json:"address" toml:"address"`
	ImageHash_ core.ImageHash `json:"imageHash" toml:"imageHash"`
}

func decodeWalletConfigTreeSapientSignerLeaf(object any) (*WalletConfigTreeSapientSignerLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wallet config tree sapient signer leaf must be an object")
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

	imageHash, ok := object_["imageHash"]
	if !ok {
		return nil, fmt.Errorf(`missing required "imageHash" property`)
	}
	imageHash_, ok := imageHash.(string)
	if !ok {
		return nil, fmt.Errorf("imageHash must be a string")
	}
	imageHashBytes, err := hexutil.Decode(imageHash_)
	if err != nil || len(imageHashBytes) != 32 {
		return nil, fmt.Errorf(`"%v" is not a valid 32-byte hash`, imageHash_)
	}

	return &WalletConfigTreeSapientSignerLeaf{
		Weight:     weight_,
		Address:    common.HexToAddress(address_),
		ImageHash_: core.ImageHash{Hash: common.BytesToHash(imageHashBytes)},
	}, nil
}

func (l *WalletConfigTreeSapientSignerLeaf) ImageHash() core.ImageHash {
	weight := new(big.Int).SetUint64(uint64(l.Weight))
	weightBytes := make([]byte, 32)
	weight.FillBytes(weightBytes)

	hash := crypto.Keccak256Hash(
		[]byte(sapientLeafImageHashPrefix),
		l.Address.Bytes(),
		weightBytes,
		l.ImageHash_.Bytes(),
	)
	return core.ImageHash{Hash: hash, Preimage: l}
}

func (l *WalletConfigTreeSapientSignerLeaf) isComplete() bool {
	return true
}

func (l *WalletConfigTreeSapientSignerLeaf) maxWeight() *big.Int {
	return new(big.Int).SetUint64(uint64(l.Weight))
}

func (l *WalletConfigTreeSapientSignerLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
	signers[l.Address] = uint16(l.Weight)
}

func (l *WalletConfigTreeSapientSignerLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	if _, ok := signers[l.Address]; ok {
		return new(big.Int).SetUint64(uint64(l.Weight))
	}
	return new(big.Int)
}

func (l *WalletConfigTreeSapientSignerLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	if signature, ok := signerSignatures[l.Address]; ok {
		switch signature.type_ {
		case core.SignerSignatureTypeSapient:
			return &signatureTreeSapientLeaf{
				Weight:    l.Weight,
				Address:   l.Address,
				Signature: signature.signature,
			}
		case core.SignerSignatureTypeSapientCompact:
			return &signatureTreeSapientCompactLeaf{
				Weight:    l.Weight,
				Address:   l.Address,
				Signature: signature.signature,
			}
		default:
			panic(fmt.Sprintf("unexpected signature type for sapient signer: %v", signature.type_))
		}
	}
	hashedImage := l.ImageHash()
	return &signatureTreeNodeLeaf{
		ImageHash: hashedImage,
	}
}

type WalletConfigTreeAnyAddressSubdigestLeaf struct {
	Digest common.Hash `json:"digest" toml:"digest"`
}

func decodeWalletConfigTreeAnyAddressSubdigestLeaf(object any) (WalletConfigTreeAnyAddressSubdigestLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return WalletConfigTreeAnyAddressSubdigestLeaf{}, fmt.Errorf("wallet config tree any address subdigest leaf must be an object")
	}

	digest, ok := object_["digest"]
	if !ok {
		return WalletConfigTreeAnyAddressSubdigestLeaf{}, fmt.Errorf(`missing required "digest" property`)
	}
	digest_, ok := digest.(string)
	if !ok {
		return WalletConfigTreeAnyAddressSubdigestLeaf{}, fmt.Errorf("digest must be a string")
	}
	digestBytes, err := hexutil.Decode(digest_)
	if err != nil || len(digestBytes) != 32 {
		return WalletConfigTreeAnyAddressSubdigestLeaf{}, fmt.Errorf(`"%v" is not a valid 32-byte hash`, digest_)
	}

	return WalletConfigTreeAnyAddressSubdigestLeaf{Digest: common.BytesToHash(digestBytes)}, nil
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) ImageHash() core.ImageHash {
	hash := crypto.Keccak256Hash(
		[]byte(anyAddressSubdigestLeafImageHashPrefix),
		l.Digest.Bytes(),
	)
	return core.ImageHash{
		Hash:     hash,
		Preimage: &l,
	}
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) isComplete() bool {
	return true
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) maxWeight() *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeAnyAddressSubdigestLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return &signatureTreeAnyAddressSubdigestLeaf{l.Digest}
}

func minBytesFor(val uint64) uint8 {
	if val == 0 {
		return 0
	}
	return uint8(math.Ceil(float64(bitsFor(val)) / 8))
}

func minBytesForUint64(value uint64) byte {
	if value == 0 {
		return 0
	}
	size := byte(1)
	for value > 0xff {
		size++
		value >>= 8
	}
	return size
}

func minBytesForUint16(value uint16) byte {
	if value == 0 {
		return 0
	}
	if value <= 0xff {
		return 1
	}
	return 2
}

func bitsFor(val uint64) int {
	if val == 0 {
		return 0
	}
	return 64 - bits.LeadingZeros64(val)
}

func readUintX(size uint8, data *[]byte) (uint64, error) {
	if len(*data) < int(size) {
		return 0, fmt.Errorf("insufficient data for uint%d", size*8)
	}
	var value uint64
	for i := range size {
		value = (value << 8) | uint64((*data)[i])
	}
	*data = (*data)[size:]
	return value, nil
}

func writeUintX(writer io.Writer, value uint64, size byte) error {
	if size == 0 {
		_, err := writer.Write([]byte{0})
		if err != nil {
			return fmt.Errorf("unable to write uint0: %w", err)
		}
		return nil
	}
	buf := make([]byte, size)
	for i := int(size) - 1; i >= 0; i-- {
		buf[i] = byte(value & 0xff)
		value >>= 8
	}
	_, err := writer.Write(buf)
	return err
}

func readUint16(data *[]byte) (uint16, error) {
	if len(*data) < 2 {
		return 0, fmt.Errorf("insufficient data for uint16")
	}
	value := binary.BigEndian.Uint16(*data)
	*data = (*data)[2:]
	return value, nil
}

func readUint24(data *[]byte) (uint32, error) {
	if len(*data) < 3 {
		return 0, fmt.Errorf("insufficient data for uint24")
	}
	value := uint32((*data)[0])<<16 | uint32((*data)[1])<<8 | uint32((*data)[2])
	*data = (*data)[3:]
	return value, nil
}

func writeUint8(writer io.Writer, value uint8) error {
	_, err := writer.Write([]byte{value})
	return err
}

func writeUint16(writer io.Writer, value uint16) error {
	buf := []byte{
		byte(value >> 8),
		byte(value),
	}
	_, err := writer.Write(buf)
	return err
}

func writeUint24(writer io.Writer, value uint32) error {
	buf := []byte{
		byte(value >> 16),
		byte(value >> 8),
		byte(value),
	}
	_, err := writer.Write(buf)
	return err
}

func toUint8(number any) (uint8, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 || number > 0xff {
			return 0, fmt.Errorf("%v does not fit in uint8", number)
		}
		return uint8(number), nil
	case float64:
		if number < 0 || number > 0xff || math.Floor(number) != number {
			return 0, fmt.Errorf("%v does not fit in uint8", number)
		}
		return uint8(number), nil
	case string:
		n, err := strconv.ParseUint(number, 10, 8)
		if err != nil {
			return 0, fmt.Errorf("unable to parse string %s as uint8: %w", number, err)
		}
		return uint8(n), nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint8", number)
	}
}

func toUint16(number any) (uint16, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 || number > 0xffff {
			return 0, fmt.Errorf("%v does not fit in uint16", number)
		}
		return uint16(number), nil
	case float64:
		if number < 0 || number > 0xffff || math.Floor(number) != number {
			return 0, fmt.Errorf("%v does not fit in uint16", number)
		}
		return uint16(number), nil
	case string:
		n, err := strconv.ParseUint(number, 10, 16)
		if err != nil {
			return 0, fmt.Errorf("unable to parse string %s as uint16: %w", number, err)
		}
		return uint16(n), nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint16", number)
	}
}

func toUint64(number any) (uint64, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 {
			return 0, fmt.Errorf("%v does not fit in uint64", number)
		}
		return uint64(number), nil
	case float64:
		if number < 0 || number > math.MaxUint64 || math.Floor(number) != number {
			return 0, fmt.Errorf("%v does not fit in uint64", number)
		}
		return uint64(number), nil
	case string:
		n, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("unable to parse string %s as uint64: %w", number, err)
		}
		return n, nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint64", number)
	}
}

func hasKeys(object map[string]any, keys []string) bool {
	if len(object) != len(keys) {
		return false
	}
	for _, key := range keys {
		if _, ok := object[key]; !ok {
			return false
		}
	}
	return true
}

var maxUint256 = new(big.Int).SetBytes([]byte{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
})
