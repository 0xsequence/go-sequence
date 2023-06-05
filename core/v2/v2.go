// Sequence v2 core primitives
package v2

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	"github.com/0xsequence/ethkit/ethcontract"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts"
	"github.com/0xsequence/go-sequence/core"
)

const (
	nestedLeafImageHashPrefix    = "Sequence nested config:\n"
	subdigestLeafImageHashPrefix = "Sequence static digest:\n"
)

var isValidSignatureMagicValue = [4]byte{0x16, 0x26, 0xba, 0x7e}

var Core core.Core[*WalletConfig, core.Signature[*WalletConfig]] = v2Core{}

func init() {
	core.RegisterCore(Core)
}

type v2Core struct{}

func (v2Core) DecodeSignature(data []byte) (core.Signature[*WalletConfig], error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("missing signature type")
	}

	switch data[0] {
	case byte(signatureTypeLegacy), byte(signatureTypeRegular):
		return decodeRegularSignature(data)

	case byte(signatureTypeNoChainID):
		return decodeNoChainIDSignature(data)

	case byte(signatureTypeChained):
		return decodeChainedSignature(data)

	default:
		return nil, fmt.Errorf("unknown signature type %v", data[0])
	}
}

func (v2Core) DecodeWalletConfig(object any) (*WalletConfig, error) {
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
	checkpoint_, err := toUint32(checkpoint)
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

	return &WalletConfig{
		Threshold_:  threshold_,
		Checkpoint_: checkpoint_,
		Tree:        tree_,
	}, nil
}

type signatureType byte

const (
	signatureTypeLegacy    signatureType = 0
	signatureTypeRegular   signatureType = 1
	signatureTypeNoChainID signatureType = 2
	signatureTypeChained   signatureType = 3
)

type regularSignature struct {
	isRegular  bool
	threshold  uint16
	checkpoint uint32
	tree       signatureTree
}

func decodeRegularSignature(data []byte) (*regularSignature, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("missing regular signature type")
	}

	var isRegular bool
	switch data[0] {
	case byte(signatureTypeLegacy):
		isRegular = false
	case byte(signatureTypeRegular):
		isRegular = true
		data = data[1:]
	default:
		return nil, fmt.Errorf("expected regular signature type %v or %v, got %v", signatureTypeLegacy, signatureTypeRegular, data[0])
	}

	threshold, err := readUint16(&data)
	if err != nil {
		return nil, fmt.Errorf("unable to read regular signature threshold: %w", err)
	}

	checkpoint, err := readUint32(&data)
	if err != nil {
		return nil, fmt.Errorf("unable to read regular signature checkpoint: %w", err)
	}

	tree, err := decodeSignatureTree(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode regular signature: %w", err)
	}

	return &regularSignature{isRegular, threshold, checkpoint, tree}, nil
}

type signerSignature struct {
	signer    common.Address
	type_     core.SignerSignatureType
	signature []byte
}

func (s *regularSignature) Threshold() uint16 {
	return s.threshold
}

func (s *regularSignature) Checkpoint() uint32 {
	return s.checkpoint
}

func (s *regularSignature) Recover(ctx context.Context, digest core.Digest, wallet common.Address, chainID *big.Int, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
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

	tree, weight, err := s.tree.recover(ctx, digest.Subdigest(wallet, chainID), provider, signerSignatures[0])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover wallet config: %w", err)
	}

	return &WalletConfig{
		Threshold_:  s.threshold,
		Checkpoint_: s.checkpoint,
		Tree:        tree,
	}, weight, nil
}

func (s *regularSignature) Join(subdigest core.Subdigest, other core.Signature[*WalletConfig]) core.Signature[*WalletConfig] {
	//TODO implement me
	panic("implement me")
}

func (s *regularSignature) Reduce(subdigest core.Subdigest) core.Signature[*WalletConfig] {
	return &regularSignature{
		isRegular:  s.isRegular,
		threshold:  s.threshold,
		checkpoint: s.checkpoint,
		tree:       s.tree.reduce(),
	}
}

func (s *regularSignature) Data() ([]byte, error) {
	var buffer bytes.Buffer

	err := s.Write(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (s *regularSignature) Write(writer io.Writer) error {
	if s.isRegular {
		_, err := writer.Write([]byte{byte(signatureTypeRegular)})
		if err != nil {
			return fmt.Errorf("unable to write regular signature type: %w", err)
		}
	}

	err := binary.Write(writer, binary.BigEndian, s.threshold)
	if err != nil {
		return fmt.Errorf("unable to write regular signature threshold: %w", err)
	}

	err = binary.Write(writer, binary.BigEndian, s.checkpoint)
	if err != nil {
		return fmt.Errorf("unable to write regular signature checkpoint: %w", err)
	}

	err = s.tree.write(writer)
	if err != nil {
		return fmt.Errorf("unable to write regular signature: %w", err)
	}

	return nil
}

func (s *regularSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type noChainIDSignature struct {
	threshold  uint16
	checkpoint uint32
	tree       signatureTree
}

func decodeNoChainIDSignature(data []byte) (*noChainIDSignature, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("missing no chain ID signature type")
	}

	switch data[0] {
	case byte(signatureTypeNoChainID):
		data = data[1:]
	default:
		return nil, fmt.Errorf("expected no chain ID signature type %v, got %v", signatureTypeNoChainID, data[0])
	}

	threshold, err := readUint16(&data)
	if err != nil {
		return nil, fmt.Errorf("unable to read no chain ID signature threshold: %w", err)
	}

	checkpoint, err := readUint32(&data)
	if err != nil {
		return nil, fmt.Errorf("unable to read no chain ID signature checkpoint: %w", err)
	}

	tree, err := decodeSignatureTree(data)
	if err != nil {
		return nil, fmt.Errorf("unable to decode no chain ID signature: %w", err)
	}

	return &noChainIDSignature{threshold, checkpoint, tree}, nil
}

func (s *noChainIDSignature) Threshold() uint16 {
	return s.threshold
}

func (s *noChainIDSignature) Checkpoint() uint32 {
	return s.checkpoint
}

func (s *noChainIDSignature) Recover(ctx context.Context, digest core.Digest, wallet common.Address, chainID *big.Int, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	tree, weight, err := s.tree.recover(ctx, digest.Subdigest(wallet), provider, signerSignatures[0])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover wallet config: %w", err)
	}

	return &WalletConfig{
		Threshold_:  s.threshold,
		Checkpoint_: s.checkpoint,
		Tree:        tree,
	}, weight, nil
}

func (s *noChainIDSignature) Join(subdigest core.Subdigest, other core.Signature[*WalletConfig]) core.Signature[*WalletConfig] {
	//TODO implement me
	panic("implement me")
}

func (s *noChainIDSignature) Reduce(subdigest core.Subdigest) core.Signature[*WalletConfig] {
	return &noChainIDSignature{
		threshold:  s.threshold,
		checkpoint: s.checkpoint,
		tree:       s.tree.reduce(),
	}
}

func (s *noChainIDSignature) Data() ([]byte, error) {
	var buffer bytes.Buffer

	err := s.Write(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (s *noChainIDSignature) Write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureTypeNoChainID)})
	if err != nil {
		return fmt.Errorf("unable to write no chain ID signature type: %w", err)
	}

	err = binary.Write(writer, binary.BigEndian, s.threshold)
	if err != nil {
		return fmt.Errorf("unable to write no chain ID signature threshold: %w", err)
	}

	err = binary.Write(writer, binary.BigEndian, s.checkpoint)
	if err != nil {
		return fmt.Errorf("unable to write no chain ID signature checkpoint: %w", err)
	}

	err = s.tree.write(writer)
	if err != nil {
		return fmt.Errorf("unable to write no chain ID signature: %w", err)
	}

	return nil
}

func (s *noChainIDSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type chainedSignature []core.Signature[*WalletConfig]

func decodeChainedSignature(data []byte) (chainedSignature, error) {
	switch len(data) {
	case 0:
		return nil, fmt.Errorf("missing signature type")
	case 1:
		return nil, fmt.Errorf("insufficient data for chained signature")
	}

	switch data[0] {
	case byte(signatureTypeChained):
		data = data[1:]
	default:
		return nil, fmt.Errorf("expected chained signature type %v, got %v", signatureTypeChained, data[0])
	}

	var signature chainedSignature

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
					len(signature),
					currentCheckpoint,
					len(signature)-1,
					previousCheckpoint,
				)
			}
		}

		signature = append(signature, subsignature)
	}

	return signature, nil
}

func (s chainedSignature) Threshold() uint16 {
	return s[len(s)-1].Threshold()
}

func (s chainedSignature) Checkpoint() uint32 {
	return s[len(s)-1].Checkpoint()
}

func (s chainedSignature) Recover(ctx context.Context, digest core.Digest, wallet common.Address, chainID *big.Int, provider *ethrpc.Provider, signerSignatures ...core.SignerSignatures) (*WalletConfig, *big.Int, error) {
	if len(signerSignatures) == 0 {
		signerSignatures = []core.SignerSignatures{nil}
	}

	var config *WalletConfig
	var weight *big.Int

	for i, subsignature := range s {
		var err error
		config, weight, err = subsignature.Recover(ctx, digest, wallet, chainID, provider, signerSignatures...)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to recover subsignature %v: %w", i, err)
		}

		if weight.Cmp(new(big.Int).SetUint64(uint64(config.Threshold()))) < 0 {
			return nil, nil, fmt.Errorf("recovered weight %v for subsignature %v does not meet required threshold %v", weight, i, config.Threshold())
		}

		digest = config.ImageHash().Approval()
	}

	return config, weight, nil
}

func (s chainedSignature) Join(subdigest core.Subdigest, other core.Signature[*WalletConfig]) core.Signature[*WalletConfig] {
	//TODO implement me
	panic("implement me")
}

func (s chainedSignature) Reduce(subdigest core.Subdigest) core.Signature[*WalletConfig] {
	subsignatures := make(chainedSignature, 0, len(s))
	for _, subsignature := range s {
		subsignatures = append(subsignatures, subsignature.Reduce(subdigest))
	}
	return subsignatures
}

func (s chainedSignature) Data() ([]byte, error) {
	var buffer bytes.Buffer

	err := s.Write(&buffer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (s chainedSignature) Write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureTypeChained)})
	if err != nil {
		return fmt.Errorf("unable to write chained signature type: %w", err)
	}

	for i, subsignature := range s {
		data, err := subsignature.Data()
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

func (s chainedSignature) String() string {
	data, _ := s.Data()
	return hexutil.Encode(data)
}

type signatureTree interface {
	recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error)
	reduce() signatureTree
	reduceImageHash() (core.ImageHash, error)
	write(writer io.Writer) error
}

func decodeSignatureTree(data []byte) (signatureTree, error) {
	var tree signatureTree

	for len(data) != 0 {
		var leaf signatureTree
		var err error

		switch data[0] {
		case byte(signatureLeafTypeECDSASignature):
			leaf, err = decodeSignatureTreeECDSASignatureLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode ecdsa signature leaf: %w", err)
			}

		case byte(signatureLeafTypeAddress):
			leaf, err = decodeSignatureTreeAddressLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode address leaf: %w", err)
			}

		case byte(signatureLeafTypeDynamicSignature):
			leaf, err = decodeSignatureTreeDynamicSignatureLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode dynamic signature leaf: %w", err)
			}

		case byte(signatureLeafTypeNode):
			leaf, err = decodeSignatureTreeNodeLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode node leaf: %w", err)
			}

		case byte(signatureLeafTypeBranch):
			leaf, err = decodeSignatureTreeBranchLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode branch leaf: %w", err)
			}

		case byte(signatureLeafTypeSubdigest):
			leaf, err = decodeSignatureTreeSubdigestLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode subdigest leaf: %w", err)
			}

		case byte(signatureLeafTypeNested):
			leaf, err = decodeSignatureTreeNestedLeaf(&data)
			if err != nil {
				return nil, fmt.Errorf("unable to decode nested leaf: %w", err)
			}

		default:
			return nil, fmt.Errorf("unknown signature leaf type %v", data[0])
		}

		if tree == nil {
			tree = leaf
		} else {
			tree = &signatureTreeNode{left: tree, right: leaf}
		}
	}

	return tree, nil
}

type signatureTreeNode struct {
	left, right signatureTree
}

func (n *signatureTreeNode) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	left, leftWeight, err := n.left.recover(ctx, subdigest, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover left subtree: %w", err)
	}

	right, rightWeight, err := n.right.recover(ctx, subdigest, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover right subtree: %w", err)
	}

	return &WalletConfigTreeNode{left, right}, new(big.Int).Add(leftWeight, rightWeight), nil
}

func (n *signatureTreeNode) reduce() signatureTree {
	if imageHash, err := n.reduceImageHash(); err == nil {
		return signatureTreeNodeLeaf{imageHash}
	} else {
		return &signatureTreeNode{
			left:  n.left.reduce(),
			right: n.right.reduce(),
		}
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

	_, ok := n.right.(*signatureTreeNode)
	if ok {
		_, err = writer.Write([]byte{byte(signatureLeafTypeBranch)})
		if err != nil {
			return fmt.Errorf("unable to write branch leaf type: %w", err)
		}

		var buffer bytes.Buffer
		err = n.right.write(&buffer)
		if err != nil {
			return fmt.Errorf("unable to encode right subtree: %w", err)
		}

		err = writeUint24(writer, uint32(buffer.Len()))
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

type signatureLeafType byte

const (
	signatureLeafTypeECDSASignature   signatureLeafType = 0
	signatureLeafTypeAddress          signatureLeafType = 1
	signatureLeafTypeDynamicSignature signatureLeafType = 2
	signatureLeafTypeNode             signatureLeafType = 3
	signatureLeafTypeBranch           signatureLeafType = 4
	signatureLeafTypeSubdigest        signatureLeafType = 5
	signatureLeafTypeNested           signatureLeafType = 6
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

func (l *signatureTreeECDSASignatureLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
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

	return &WalletConfigTreeAddressLeaf{
		Weight:  l.weight,
		Address: address,
	}, new(big.Int).SetUint64(uint64(l.weight)), nil
}

func (l *signatureTreeECDSASignatureLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeECDSASignatureLeaf) reduceImageHash() (core.ImageHash, error) {
	return core.ImageHash{}, fmt.Errorf("ecdsa signature leaf has signing power")
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

func (l *signatureTreeAddressLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	return &WalletConfigTreeAddressLeaf{
		Weight:  l.weight,
		Address: l.address,
	}, new(big.Int), nil
}

func (l *signatureTreeAddressLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeAddressLeaf) reduceImageHash() (core.ImageHash, error) {
	leaf := WalletConfigTreeAddressLeaf{
		Weight:  l.weight,
		Address: l.address,
	}
	return leaf.ImageHash(), nil
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

	if len(next) < 1+1+common.AddressLength+3 {
		return nil, fmt.Errorf("insufficient data for dynamic signature leaf")
	}

	var leafType byte
	var weight byte
	var address []byte
	var signatureLength []byte
	leafType, weight, address, signatureLength, next = next[0], next[1], next[1+1:1+1+common.AddressLength], next[1+1+common.AddressLength:1+1+common.AddressLength+3], next[1+1+common.AddressLength+3:]

	if leafType != byte(signatureLeafTypeDynamicSignature) {
		return nil, fmt.Errorf("expected dynamic signature leaf type %v, got %v", signatureLeafTypeDynamicSignature, leafType)
	}

	length, err := readUint24(&signatureLength)
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

func (l *signatureTreeDynamicSignatureLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
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

		return &WalletConfigTreeAddressLeaf{
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

		return &WalletConfigTreeAddressLeaf{
			Weight:  l.weight,
			Address: l.address,
		}, new(big.Int).SetUint64(uint64(l.weight)), nil

	default:
		return nil, nil, fmt.Errorf("unknown dynamic signature type %v", l.type_)
	}
}

func (l *signatureTreeDynamicSignatureLeaf) reduce() signatureTree {
	return l
}

func (l *signatureTreeDynamicSignatureLeaf) reduceImageHash() (core.ImageHash, error) {
	if l.weight == 0 {
		leaf := WalletConfigTreeAddressLeaf{
			Weight:  l.weight,
			Address: l.address,
		}
		return leaf.ImageHash(), nil
	}

	return core.ImageHash{}, fmt.Errorf("dynamic signature leaf might have signing power")
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

	err = writeUint24(writer, uint32(len(l.signature)+1))
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

type signatureTreeNodeLeaf struct{ core.ImageHash }

func decodeSignatureTreeNodeLeaf(data *[]byte) (signatureTreeNodeLeaf, error) {
	next := *data

	if len(next) < 1+common.HashLength {
		return signatureTreeNodeLeaf{}, fmt.Errorf("insufficient data for node leaf")
	}

	var leafType byte
	var hash []byte
	leafType, hash, next = next[0], next[1:1+common.HashLength], next[1+common.HashLength:]

	if leafType != byte(signatureLeafTypeNode) {
		return signatureTreeNodeLeaf{}, fmt.Errorf("expected node leaf type %v, got %v", signatureLeafTypeNode, leafType)
	}

	*data = next

	return signatureTreeNodeLeaf{core.ImageHash{Hash: common.BytesToHash(hash)}}, nil
}

func (l signatureTreeNodeLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	leaf := WalletConfigTreeNodeLeaf{core.ImageHash{Hash: l.ImageHash.Hash}}
	leaf.Node.Preimage = &leaf
	return leaf, new(big.Int), nil
}

func (l signatureTreeNodeLeaf) reduce() signatureTree {
	return l
}

func (l signatureTreeNodeLeaf) reduceImageHash() (core.ImageHash, error) {
	return WalletConfigTreeNodeLeaf{l.ImageHash}.ImageHash(), nil
}

func (l signatureTreeNodeLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeNode)})
	if err != nil {
		return fmt.Errorf("unable to write node leaf type: %w", err)
	}

	_, err = writer.Write(l.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write node leaf hash: %w", err)
	}

	return nil
}

func decodeSignatureTreeBranchLeaf(data *[]byte) (signatureTree, error) {
	next := *data

	if len(next) < 1+3 {
		return signatureTreeNodeLeaf{}, fmt.Errorf("insufficient data for branch leaf")
	}

	var leafType byte
	var signatureLength []byte
	leafType, signatureLength, next = next[0], next[1:1+3], next[1+3:]

	if leafType != byte(signatureLeafTypeBranch) {
		return nil, fmt.Errorf("expected branch leaf type %v, got %v", signatureLeafTypeBranch, leafType)
	}

	length, err := readUint24(&signatureLength)
	if err != nil {
		return nil, fmt.Errorf("unable to read signature length: %w", err)
	}

	if len(next) < int(length) {
		return nil, fmt.Errorf("insufficient data for branch leaf")
	}

	var signature []byte
	signature, next = next[:length], next[length:]

	branch, err := decodeSignatureTree(signature)
	if err != nil {
		return nil, fmt.Errorf("unable to decode branch leaf: %w", err)
	}

	*data = next

	return branch, nil
}

type signatureTreeSubdigestLeaf struct{ core.Subdigest }

func decodeSignatureTreeSubdigestLeaf(data *[]byte) (signatureTreeSubdigestLeaf, error) {
	next := *data

	if len(next) < 1+common.HashLength {
		return signatureTreeSubdigestLeaf{}, fmt.Errorf("insufficient data for subdigest leaf")
	}

	var leafType byte
	var hash []byte
	leafType, hash, next = next[0], next[1:1+common.HashLength], next[1+common.HashLength:]

	if leafType != byte(signatureLeafTypeSubdigest) {
		return signatureTreeSubdigestLeaf{}, fmt.Errorf("expected subdigest leaf type %v, got %v", signatureLeafTypeSubdigest, leafType)
	}

	*data = next

	return signatureTreeSubdigestLeaf{core.Subdigest{Hash: common.BytesToHash(hash)}}, nil
}

func (l signatureTreeSubdigestLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	if subdigest.Hash == l.Subdigest.Hash {
		return WalletConfigTreeSubdigestLeaf{l.Subdigest}, new(big.Int).Set(maxUint256), nil
	} else {
		return WalletConfigTreeSubdigestLeaf{l.Subdigest}, new(big.Int), nil
	}
}

func (l signatureTreeSubdigestLeaf) reduce() signatureTree {
	return l
}

func (l signatureTreeSubdigestLeaf) reduceImageHash() (core.ImageHash, error) {
	return WalletConfigTreeSubdigestLeaf{l.Subdigest}.ImageHash(), nil
}

func (l signatureTreeSubdigestLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeSubdigest)})
	if err != nil {
		return fmt.Errorf("unable to write subdigest leaf type: %w", err)
	}

	_, err = writer.Write(l.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write subdigest leaf hash: %w", err)
	}

	return nil
}

type signatureTreeNestedLeaf struct {
	weight    uint8
	threshold uint16
	tree      signatureTree
}

func decodeSignatureTreeNestedLeaf(data *[]byte) (*signatureTreeNestedLeaf, error) {
	next := *data

	if len(next) < 1+1+2+3 {
		return nil, fmt.Errorf("insufficient data for nested leaf")
	}

	var leafType byte
	var weight byte
	var internalThreshold []byte
	var signatureLength []byte
	leafType, weight, internalThreshold, signatureLength, next = next[0], next[1], next[1+1:1+1+2], next[1+1+2:1+1+2+3], next[1+1+2+3:]

	if leafType != byte(signatureLeafTypeNested) {
		return nil, fmt.Errorf("expected nested leaf type %v, got %v", signatureLeafTypeNested, leafType)
	}

	threshold, err := readUint16(&internalThreshold)
	if err != nil {
		return nil, fmt.Errorf("unable to read nested leaf threshold: %w", err)
	}

	length, err := readUint24(&signatureLength)
	if err != nil {
		return nil, fmt.Errorf("unable to read nested leaf signature length: %w", err)
	}

	if len(next) < int(length) {
		return nil, fmt.Errorf("insufficient data for nested leaf")
	}

	var signature []byte
	signature, next = next[:length], next[length:]

	tree, err := decodeSignatureTree(signature)
	if err != nil {
		return nil, fmt.Errorf("unable to decode nested leaf: %w", err)
	}

	*data = next

	return &signatureTreeNestedLeaf{weight, threshold, tree}, nil
}

func (l *signatureTreeNestedLeaf) recover(ctx context.Context, subdigest core.Subdigest, provider *ethrpc.Provider, signerSignatures core.SignerSignatures) (WalletConfigTree, *big.Int, error) {
	tree, weight, err := l.tree.recover(ctx, subdigest, provider, signerSignatures)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to recover nested leaf: %w", err)
	}

	if weight.Cmp(new(big.Int).SetUint64(uint64(l.threshold))) >= 0 {
		return tree, new(big.Int).SetUint64(uint64(l.weight)), nil
	} else {
		return tree, new(big.Int), nil
	}
}

func (l *signatureTreeNestedLeaf) reduce() signatureTree {
	return &signatureTreeNestedLeaf{
		weight:    l.weight,
		threshold: l.threshold,
		tree:      l.tree.reduce(),
	}
}

func (l *signatureTreeNestedLeaf) reduceImageHash() (core.ImageHash, error) {
	if l.weight == 0 {
		if imageHash, err := l.tree.reduceImageHash(); err == nil {
			leaf := WalletConfigTreeNestedLeaf{
				Weight:    l.weight,
				Threshold: l.threshold,
				Tree:      WalletConfigTreeNodeLeaf{imageHash},
			}
			return leaf.ImageHash(), nil
		}
	}

	return core.ImageHash{}, fmt.Errorf("nested leaf might have signing power")
}

func (l *signatureTreeNestedLeaf) write(writer io.Writer) error {
	_, err := writer.Write([]byte{byte(signatureLeafTypeNested)})
	if err != nil {
		return fmt.Errorf("unable to write nested leaf type: %w", err)
	}

	_, err = writer.Write([]byte{l.weight})
	if err != nil {
		return fmt.Errorf("unable to write nested leaf weight: %w", err)
	}

	err = binary.Write(writer, binary.BigEndian, l.threshold)
	if err != nil {
		return fmt.Errorf("unable to write nested leaf threshold: %w", err)
	}

	var buffer bytes.Buffer
	err = l.tree.write(&buffer)
	if err != nil {
		return fmt.Errorf("unable to encode nested leaf signature: %w", err)
	}

	err = writeUint24(writer, uint32(len(buffer.Bytes())))
	if err != nil {
		return fmt.Errorf("unable to write nested leaf signature length: %w", err)
	}

	_, err = writer.Write(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("unable to write nested leaf signature: %w", err)
	}

	return nil
}

type WalletConfig struct {
	Threshold_  uint16           `json:"threshold" toml:"threshold"`
	Checkpoint_ uint32           `json:"checkpoint" toml:"checkpoint"`
	Tree        WalletConfigTree `json:"tree" toml:"tree"`
}

func (c *WalletConfig) Threshold() uint16 {
	return c.Threshold_
}

func (c *WalletConfig) Checkpoint() uint32 {
	return c.Checkpoint_
}

func (c *WalletConfig) Signers() map[common.Address]uint16 {
	signers := map[common.Address]uint16{}
	c.Tree.readSignersIntoMap(signers)
	return signers
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
	imageHash := c.Tree.ImageHash()
	threshold := common.BigToHash(new(big.Int).SetUint64(uint64(c.Threshold_)))
	checkpoint := common.BigToHash(new(big.Int).SetUint64(uint64(c.Checkpoint_)))

	return core.ImageHash{
		Hash: crypto.Keccak256Hash(
			crypto.Keccak256Hash(
				imageHash.Bytes(),
				threshold.Bytes(),
			).Bytes(),
			checkpoint.Bytes(),
		),
		Preimage: c,
	}
}

func (c *WalletConfig) BuildRegularSignature(ctx context.Context, sign core.SigningFunction, validateSigningPower ...bool) (core.Signature[*WalletConfig], error) {
	isValid := false
	if len(validateSigningPower) > 0 {
		isValid = !validateSigningPower[0]
	}

	configSigners := c.Signers()

	signCtx, signCancel := context.WithCancel(ctx)
	defer signCancel()

	signerSignatureCh := core.SigningOrchestrator(signCtx, configSigners, sign)

	signerSignatures := map[common.Address]signerSignature{}
	signedSigners := map[common.Address]uint16{}

	for range configSigners {
		signerSig := <-signerSignatureCh

		if signerSig.Signature != nil {
			signerSignatures[signerSig.Signer] = signerSignature{signerSig.Signer,
				signerSig.Type, signerSig.Signature}
			signedSigners[signerSig.Signer] = 0

			weight := c.Tree.unverifiedWeight(signedSigners)
			if weight.Cmp(new(big.Int).SetUint64(uint64(c.Threshold_))) >= 0 {
				signCancel()
				isValid = true
			}
		}
	}

	if isValid {
		return &regularSignature{
			isRegular:  true,
			threshold:  c.Threshold_,
			checkpoint: c.Checkpoint_,
			tree:       c.Tree.buildSignatureTree(signerSignatures),
		}, nil
	} else {
		return nil, fmt.Errorf("not enough signers to build regular signature")
	}
}

func (c *WalletConfig) BuildNoChainIDSignature(ctx context.Context, sign core.SigningFunction) (core.Signature[*WalletConfig], error) {
	configSigners := c.Signers()

	signCtx, signCancel := context.WithCancel(ctx)
	defer signCancel()

	signerSignatureCh := core.SigningOrchestrator(signCtx, configSigners, sign)

	signerSignatures := map[common.Address]signerSignature{}
	signedSigners := map[common.Address]uint16{}
	isValid := false

	for range configSigners {
		signerSig := <-signerSignatureCh

		if signerSig.Signature != nil {
			signerSignatures[signerSig.Signer] = signerSignature{signerSig.Signer,
				signerSig.Type, signerSig.Signature}
			signedSigners[signerSig.Signer] = 0

			weight := c.Tree.unverifiedWeight(signedSigners)
			if weight.Cmp(new(big.Int).SetUint64(uint64(c.Threshold_))) >= 0 {
				signCancel()
				isValid = true
			}
		}
	}

	if isValid {
		return &noChainIDSignature{
			threshold:  c.Threshold_,
			checkpoint: c.Checkpoint_,
			tree:       c.Tree.buildSignatureTree(signerSignatures),
		}, nil
	} else {
		return nil, fmt.Errorf("not enough signers to build no chain ID signature")
	}
}

type WalletConfigTree interface {
	core.ImageHashable

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
	} else {
		return nil, fmt.Errorf("unknown wallet config tree type")
	}
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
			for i := 0; i < numberOfNodes; i++ {
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

func (n *WalletConfigTreeNode) maxWeight() *big.Int {
	left, right := n.Left.maxWeight(), n.Right.maxWeight()
	return new(big.Int).Add(left, right)
}

func (n *WalletConfigTreeNode) readSignersIntoMap(signers map[common.Address]uint16) {
	n.Left.readSignersIntoMap(signers)
	n.Right.readSignersIntoMap(signers)
}

func (n *WalletConfigTreeNode) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	left, right := n.Left.unverifiedWeight(signers), n.Right.unverifiedWeight(signers)
	return new(big.Int).Add(left, right)
}

func (n *WalletConfigTreeNode) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return &signatureTreeNode{
		left:  n.Left.buildSignatureTree(signerSignatures),
		right: n.Right.buildSignatureTree(signerSignatures),
	}
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
	hash := l.Address.Hash()
	hash[common.HashLength-common.AddressLength-1] = l.Weight
	return core.ImageHash{Hash: hash, Preimage: l}
}

func (l *WalletConfigTreeAddressLeaf) maxWeight() *big.Int {
	return new(big.Int).SetUint64(uint64(l.Weight))
}

func (l *WalletConfigTreeAddressLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
	signers[l.Address] = uint16(l.Weight)
}

func (l *WalletConfigTreeAddressLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	_, ok := signers[l.Address]
	if ok {
		return new(big.Int).SetUint64(uint64(l.Weight))
	} else {
		return new(big.Int)
	}
}

func (l *WalletConfigTreeAddressLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	signature, ok := signerSignatures[l.Address]
	if ok {
		switch signature.type_ {
		case core.SignerSignatureTypeEIP712:
			leaf := signatureTreeECDSASignatureLeaf{
				weight: l.Weight,
				type_:  eCDSASignatureTypeEIP712,
			}
			copy(leaf.signature[:], signature.signature)
			return &leaf

		case core.SignerSignatureTypeEthSign:
			leaf := signatureTreeECDSASignatureLeaf{
				weight: l.Weight,
				type_:  eCDSASignatureTypeEthSign,
			}
			copy(leaf.signature[:], signature.signature)
			return &leaf

		case core.SignerSignatureTypeEIP1271:
			return &signatureTreeDynamicSignatureLeaf{
				weight:    l.Weight,
				address:   l.Address,
				type_:     dynamicSignatureTypeEIP1271,
				signature: signature.signature,
			}

		default:
			panic(fmt.Sprintf("unknown signer signature type %v", signature.type_))
		}
	} else {
		return &signatureTreeAddressLeaf{
			weight:  l.Weight,
			address: l.Address,
		}
	}
}

func (l *WalletConfigTreeAddressLeaf) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"weight":  l.Weight,
		"address": l.Address.String(),
	})
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
		return WalletConfigTreeNodeLeaf{}, fmt.Errorf("expected hash of length %v, got hash of length %v", common.HashLength, len(node__))
	}

	leaf := WalletConfigTreeNodeLeaf{core.ImageHash{Hash: common.BytesToHash(node__)}}
	leaf.Node.Preimage = &leaf
	return leaf, nil
}

func (l WalletConfigTreeNodeLeaf) ImageHash() core.ImageHash {
	return l.Node
}

func (l WalletConfigTreeNodeLeaf) maxWeight() *big.Int {
	return new(big.Int)
}

func (l WalletConfigTreeNodeLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
}

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
		return nil, fmt.Errorf("unable to decode wallet config tree: %w", err)
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

func (l *WalletConfigTreeNestedLeaf) maxWeight() *big.Int {
	if l.Tree.maxWeight().Cmp(new(big.Int).SetUint64(uint64(l.Threshold))) >= 0 {
		return new(big.Int).SetUint64(uint64(l.Weight))
	} else {
		return new(big.Int)
	}
}

func (l *WalletConfigTreeNestedLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
	l.Tree.readSignersIntoMap(signers)
}

func (l *WalletConfigTreeNestedLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	if l.Tree.unverifiedWeight(signers).Cmp(new(big.Int).SetUint64(uint64(l.Threshold))) >= 0 {
		return new(big.Int).SetUint64(uint64(l.Weight))
	} else {
		return new(big.Int)
	}
}

func (l *WalletConfigTreeNestedLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return &signatureTreeNestedLeaf{
		weight:    l.Weight,
		threshold: l.Threshold,
		tree:      l.Tree.buildSignatureTree(signerSignatures),
	}
}

type WalletConfigTreeSubdigestLeaf struct {
	Subdigest core.Subdigest `json:"subdigest" toml:"subdigest"`
}

func decodeWalletConfigTreeSubdigestLeaf(object any) (WalletConfigTreeSubdigestLeaf, error) {
	object_, ok := object.(map[string]any)
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("wallet config tree subdigest leaf must be an object")
	}

	subdigest, ok := object_["subdigest"]
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf(`missing required "subdigest" property`)
	}
	subdigest_, ok := subdigest.(string)
	if !ok {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("subdigest must be a string")
	}
	subdigest__, err := hexutil.Decode(subdigest_)
	if err != nil {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf(`"%v" is not valid hex`, subdigest_)
	}
	if len(subdigest__) != common.HashLength {
		return WalletConfigTreeSubdigestLeaf{}, fmt.Errorf("expected hash of length %v, got hash of length %v", common.HashLength, len(subdigest__))
	}

	return WalletConfigTreeSubdigestLeaf{core.Subdigest{Hash: common.BytesToHash(subdigest__)}}, nil
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

func (l WalletConfigTreeSubdigestLeaf) maxWeight() *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeSubdigestLeaf) readSignersIntoMap(signers map[common.Address]uint16) {
}

func (l WalletConfigTreeSubdigestLeaf) unverifiedWeight(signers map[common.Address]uint16) *big.Int {
	return new(big.Int).Set(maxUint256)
}

func (l WalletConfigTreeSubdigestLeaf) buildSignatureTree(signerSignatures map[common.Address]signerSignature) signatureTree {
	return signatureTreeSubdigestLeaf{l.Subdigest}
}

func hasKeys(object map[string]any, keys []string) bool {
	if len(object) != len(keys) {
		return false
	}

	for _, key := range keys {
		_, ok := object[key]
		if !ok {
			return false
		}
	}

	return true
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

func toUint32(number any) (uint32, error) {
	switch number := number.(type) {
	case int64:
		if number < 0 || number > 0xffffffff {
			return 0, fmt.Errorf("%v does not fit in a uint32", number)
		}
		return uint32(number), nil
	case float64:
		if number < 0 || number > 0xffffffff {
			return 0, fmt.Errorf("%v does not fit in a uint32", number)
		}
		return uint32(number), nil
	default:
		return 0, fmt.Errorf("unable to convert %v to uint32", number)
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

func readUint24(data *[]byte) (uint32, error) {
	if len(*data) < 3 {
		return 0, fmt.Errorf("insufficient data for uint24")
	}

	value := uint32((*data)[0])<<16 | uint32((*data)[1])<<8 | uint32((*data)[2])<<0
	*data = (*data)[3:]
	return value, nil
}

func readUint32(data *[]byte) (uint32, error) {
	if len(*data) < 4 {
		return 0, fmt.Errorf("insufficient data for uint32")
	}

	value := binary.BigEndian.Uint32(*data)
	*data = (*data)[4:]
	return value, nil
}

func writeUint24(writer io.Writer, value uint32) error {
	_, err := writer.Write([]byte{byte(value >> 16), byte(value >> 8), byte(value >> 0)})
	if err != nil {
		return fmt.Errorf("unable to write uint24: %w", err)
	}

	return nil
}

var maxUint256 = new(big.Int).SetBytes([]byte{
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
})
