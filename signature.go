package sequence

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/ethwallet"
	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi/bind"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
	"github.com/0xsequence/go-sequence/contracts"
)

// Signature for sequence message
type Signature struct {
	Threshold uint16         `json:"threshold"`
	Signers   SignatureParts `json:"signers"`
}

// SignaturePart of an overall sequence signature message
type SignaturePart struct {
	Weight  uint8          `json:"weight"`
	Address common.Address `json:"address"`
	Type    uint8          `json:"type"`  // signature type, ie. EOA, Address, Dynamic
	Value   []byte         `json:"value"` // signature value for this part
}

type SignatureParts []*SignaturePart

func (a SignatureParts) Len() int { return len(a) }
func (a SignatureParts) Less(i, j int) bool {
	return a[i].Address.Hash().Big().Cmp(a[j].Address.Hash().Big()) < 0
}
func (a SignatureParts) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

const (
	sigPartTypeEOA     uint8 = iota // 0
	sigPartTypeAddress              // 1
	sigPartTypeDynamic              // 2
)

const (
	sigTypeEip712  uint8 = iota + 1 // 1
	sigTypeEthSign                  // 2
	sigTypeEip1271                  // 3
)

// DecodeSignature sequence into individual parts
func DecodeSignature(sequenceSignature []byte) (*Signature, error) {
	sig := sequenceSignature

	s := len(sig)
	if s < 2 {
		return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds threshold")
	}

	threshold := binary.BigEndian.Uint16(sig[0:2])
	parts := []*SignaturePart{}

	for i := 2; i < s; {
		// Read type of entry
		ni := i + 1
		if s < ni {
			return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds signature type")
		}
		t := uint8(sig[i])
		i = ni

		// Read weight
		ni = i + 1
		if s < ni {
			return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds weight")
		}
		w := uint8(sig[i])
		i = ni

		var p *SignaturePart

		// Process corresponding type
		switch {

		case t == sigPartTypeEOA:
			// Read signature
			ni := i + 66
			if s < ni {
				return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds eoa signature")
			}
			sp := sig[i:ni]
			i = ni

			// Create signature part
			p = &SignaturePart{
				Weight: w,
				Value:  sp,
				Type:   t,
			}

		case t == sigPartTypeAddress:
			// Read address
			ni := i + 20
			if s < ni {
				return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds address part")
			}
			addr := common.BytesToAddress(sig[i:ni])
			i = ni

			// Create address part
			p = &SignaturePart{
				Weight:  w,
				Address: addr,
				Type:    t,
			}

		case t == sigPartTypeDynamic:
			// Read address
			ni := i + 20
			if s < ni {
				return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds dynamic address")
			}
			addr := common.BytesToAddress(sig[i:ni])
			i = ni

			// Read signature size
			ni = i + 2
			if s < ni {
				return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds signature size")
			}
			sigsize := int(binary.BigEndian.Uint16(sig[i:ni]))
			i = ni

			// Read signature
			ni = i + sigsize
			if s < ni {
				return nil, fmt.Errorf("sequence: invalid signature, out-of-bounds dynamic signature")
			}
			sp := sig[i:ni]
			i = ni

			// Create dynamic signature part
			p = &SignaturePart{
				Weight:  w,
				Address: addr,
				Value:   sp,
				Type:    t,
			}

		default:
			return nil, fmt.Errorf("sequence: invalid signature, unknown part")
		}

		parts = append(parts, p)
	}

	return &Signature{
		Threshold: threshold,
		Signers:   parts,
	}, nil
}

// Encode returns encoded sequence signature bytes of all signatures in the Signers set
func (s *Signature) Encode() ([]byte, error) {
	sig := []byte{}

	for _, signer := range s.Signers {
		var pack []byte
		var err error

		switch signer.Type {
		case sigPartTypeAddress:
			pack, err = ethcoder.SolidityPack(
				[]string{"uint8", "uint8", "address"},
				[]interface{}{sigPartTypeAddress, signer.Weight, signer.Address},
			)
		case sigPartTypeEOA:
			pack, err = ethcoder.SolidityPack(
				[]string{"uint8", "uint8", "bytes"},
				[]interface{}{signer.Type, signer.Weight, signer.Value},
			)
		case sigPartTypeDynamic:
			pack, err = ethcoder.SolidityPack(
				[]string{"uint8", "uint8", "address", "uint16", "bytes"},
				[]interface{}{signer.Type, signer.Weight, signer.Address, len(signer.Value), signer.Value},
			)

		default:
			return nil, fmt.Errorf("known signer type")
		}

		if err != nil {
			return nil, err
		}

		sig = append(sig, pack...)
	}

	encodedSig, err := ethcoder.SolidityPack([]string{"uint16", "bytes"}, []interface{}{s.Threshold, sig})
	if err != nil {
		return nil, err
	}

	return encodedSig, nil
}

// Recover signers of signature
func (s *Signature) Recover(msg []byte, provider *ethrpc.Provider) error {
	for _, p := range s.Signers {
		if len(p.Value) != 0 {
			if p.Address != zeroAddress {
				// Dynamic signature
				// validate the signature
				var digest [32]byte
				copy(digest[:], msg[:32])

				// validate the dynamic signature if provider is passed (in most cases it should be.)
				if provider != nil {
					isValid, err := p.IsValid(digest, provider)
					if err != nil {
						return err
					}
					if !isValid {
						return fmt.Errorf("invalid signature for %v", p.Address)
					}
				}

			} else {
				// EOA Signature
				// recover signer
				rc, err := p.Recover(msg)
				if err != nil {
					return err
				}

				if p.Address == zeroAddress || len(p.Address.Bytes()) != 0 {
					p.Address = rc
				} else {
					if p.Address != rc {
						return fmt.Errorf("Recover signature conflict - prev: %v new: %v", p.Address, rc)
					}
				}
			}
		}
	}

	return nil
}

// Recover signature part
func (p *SignaturePart) Recover(msg []byte) (common.Address, error) {
	sigType := uint8(p.Value[len(p.Value)-1])

	if sigType != sigTypeEthSign {
		return common.Address{}, fmt.Errorf("signature type not implemented %d", sigType)
	}

	if len(p.Value) != 65 && len(p.Value) != 66 {
		return common.Address{}, fmt.Errorf("bad length, expected 65 or 66 but found %d", len(p.Value))
	}

	m := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)
	h := crypto.Keccak256([]byte(m))

	sigx := make([]byte, 65)
	copy(sigx, p.Value)
	if sigx[64] > 1 {
		sigx[64] -= 27
	}

	pubkey, err := crypto.Ecrecover(h, sigx)
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(crypto.Keccak256(pubkey[1:])[12:]), nil
}

// Weight of combined signatures, assumes all signatures to be valid
func (s *Signature) Weight() (uint16, error) {
	weight := uint16(0)
	for _, p := range s.Signers {
		if len(p.Value) != 0 {
			weight += uint16(p.Weight)
		}
	}
	return weight, nil
}

// Copy a signature
func (s *Signature) Copy() *Signature {
	parts := make([]*SignaturePart, len(s.Signers))

	for _, p := range s.Signers {
		ns := make([]byte, len(p.Value))
		copy(ns, p.Value)
		parts = append(parts, &SignaturePart{
			Weight:  p.Weight,
			Address: p.Address,
			Value:   ns,
		})
	}

	return &Signature{
		Threshold: s.Threshold,
		Signers:   parts,
	}
}

// JoinTwo signatures, saves the aggregated one into s1
func (s *Signature) JoinTwo(s2 *Signature) error {
	s.Signers = append(s.Signers, s2.Signers...)

	for _, p := range s2.Signers {
		if p.Address == zeroAddress || len(p.Address.Bytes()) == 0 {
			return fmt.Errorf("Can't agregate non-recovered signatures")
		}

		var px *SignaturePart

		for i := range s.Signers {
			if s.Signers[i].Address == p.Address {
				px = s.Signers[i]
				break
			}
		}

		switch {
		// s2 entry does not exists on s1, copy entry as-is
		case px == nil:
			s.Signers = append(s.Signers, p)
		// s1 entry exists, validate that we don't have different signatures for the same signer
		case len(px.Value) != 0:
			if bytes.Compare(px.Value, p.Value) != 0 {
				return fmt.Errorf("Can't join different signed messages")
			}
		// s1 entry exists, append signature
		case len(px.Value) != 1:
			px.Value = p.Value
		}
	}

	return nil
}

// ImageHash returns the imagehash for a given signature
func (s *Signature) ImageHash() ([32]byte, error) {
	wc := WalletConfig{
		Threshold: s.Threshold,
		Signers:   WalletConfigSigners{},
	}

	for i, p := range s.Signers {
		if p.Address == zeroAddress {
			return [32]byte{}, fmt.Errorf("address not found for index %d", i)
		}
		wc.Signers = append(wc.Signers, WalletConfigSigner{
			Weight: p.Weight, Address: p.Address,
		})
	}

	imageHash, err := ImageHashOfWalletConfigBytes(wc)
	if err != nil {
		return [32]byte{}, err
	}

	b := [32]byte{}
	copy(b[:], imageHash)

	return b, nil
}

// IsValid signature part
func (p *SignaturePart) IsValid(digest [32]byte, provider *ethrpc.Provider) (bool, error) {
	sigType := uint8(p.Value[len(p.Value)-1])

	if p.Address == zeroAddress {
		return false, fmt.Errorf("can't validate signature without address")
	}

	switch {

	case sigType == sigTypeEip712:
		return false, fmt.Errorf("signature eip 712 not implemented")

	case sigType == sigTypeEthSign:
		// Recover address and compare with the one in the part
		recovered, err := p.Recover(digest[:])
		if err != nil {
			return false, err
		}

		return recovered == p.Address, nil

	case sigType == sigTypeEip1271:
		// Call EIP1271 contract to validate it's siganture
		// NOTE: the wallet must be deployed, it doesn't support counter-factual sigantures
		erc1271, err := contracts.NewIERC1271(p.Address, provider)
		if err != nil {
			return false, err
		}

		res, err := erc1271.IsValidSignature(&bind.CallOpts{From: common.Address{0x1}}, digest, p.Value[:len(p.Value)-2])
		if err != nil {
			return false, err
		}

		return contracts.ABI_ERC1271_isValidSignatureBytes32 != hexutil.Encode(res[:]), nil

	default:
		return false, fmt.Errorf("signature type not implemented %d", sigType)

	}
}

// Join signatures
func JoinSignatures(sigs ...*Signature) (*Signature, error) {
	switch {
	case len(sigs) == 0:
		return nil, fmt.Errorf("No signatures provided")
	case len(sigs) == 1:
		return sigs[0].Copy(), nil
	default:
		a := sigs[0].Copy()
		for _, x := range sigs[1:] {
			err := a.JoinTwo(x)
			if err != nil {
				return nil, err
			}
		}
		return a, nil
	}
}

func RecoverWalletConfigFromDigest(digest, seqSig []byte, context WalletContext, chainID *big.Int, provider *ethrpc.Provider) (WalletConfig, error) {
	decoded, err := DecodeSignature(seqSig)
	if err != nil {
		return WalletConfig{}, err
	}

	err = decoded.Recover(digest, provider)
	if err != nil {
		return WalletConfig{}, err
	}

	wc := WalletConfig{Threshold: decoded.Threshold}

	for _, signer := range decoded.Signers {
		wc.Signers = append(wc.Signers, WalletConfigSigner{
			Weight:  signer.Weight,
			Address: signer.Address,
		})
	}

	return wc, nil
}

func IsValidSignature(walletAddress common.Address, digest, seqSig []byte, walletContext WalletContext, chainID *big.Int, provider *ethrpc.Provider) (bool, error) {
	// Try to do it first with ethereum sign signature format
	// TODO: .... update.. use IsValidEOASignature() ..
	ok, err := ethwallet.ValidateEthereumSignature(walletAddress.String(), digest, hexutil.Encode(seqSig))
	if err == nil {
		return ok, nil
	}

	code, err := provider.CodeAt(context.Background(), walletAddress, nil)
	if err != nil {
		return false, err
	}

	if len(code) == 0 {
		// It may be a signature from a non-deployed sequence wallet, check and attempt to validate
		subDigest, err := SubDigest(walletAddress, chainID, digest)
		if err != nil {
			return false, err
		}
		recoveredWalletConfig, err := RecoverWalletConfigFromDigest(subDigest, seqSig, walletContext, chainID, nil)
		if err != nil {
			return false, err
		}
		recoveredAddress, err := AddressFromWalletConfig(recoveredWalletConfig, walletContext)
		if err != nil {
			return false, err
		}
		if recoveredAddress != walletAddress {
			return false, fmt.Errorf("failed to validate")
		}

	} else {
		// Smart wallet is deployed, query erc1271 method to check signature
		erc1271, err := contracts.NewIERC1271(walletAddress, provider)
		if err != nil {
			return false, err
		}

		// NOTE: we expect digest to be ready for ERC1271 call, so digest must be fully encoded as expected
		hash := [32]byte{}
		copy(hash[:], digest)

		res, err := erc1271.IsValidSignature(&bind.CallOpts{From: common.Address{0x1}}, hash, seqSig)
		if err != nil {
			return false, err
		}

		if contracts.ABI_ERC1271_isValidSignatureBytes32 != hexutil.Encode(res[:]) {
			return false, fmt.Errorf("failed to validate")
		}
	}

	return true, nil
}

func MessageDigest(message []byte) []byte {
	return ethcoder.Keccak256(message)
}

func SubDigest(address common.Address, chainID *big.Int, digest []byte) ([]byte, error) {
	if chainID == nil {
		return nil, ErrUnknownChainID
	}

	// sequence smart wallet uses additional encoding of the digest in IsValidSignature()
	packedData, err := PackMessageData(chainID, address, digest)
	if err != nil {
		return nil, fmt.Errorf("subDigest, packageMessageData failed: %w", err)
	}

	// returns subdigest
	return ethcoder.Keccak256(packedData), nil
}

// PackMessageData encodes a Sequence contract "message"
func PackMessageData(chainID *big.Int, owner common.Address, digest []byte) ([]byte, error) {
	if chainID == nil {
		return nil, ErrUnknownChainID
	}
	output, err := ethcoder.SolidityPack([]string{"string", "uint256", "address", "bytes"}, []interface{}{
		"\x19\x01", chainID, owner, digest,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}

func MustEncodeSig(str string) common.Hash {
	return crypto.Keccak256Hash([]byte(str))
}