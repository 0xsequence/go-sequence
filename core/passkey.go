package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"math/big"

	"github.com/0xsequence/ethkit/go-ethereum/accounts/abi"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

// ────────────────────────────────────────────────────────────────────────────────
// Constants & flag‑bits (must match Passkeys.sol)
// ────────────────────────────────────────────────────────────────────────────────
const (
	flagRequireUV   = 0x01
	flagAuth16      = 0x02
	flagClient16    = 0x04
	flagChallenge16 = 0x08
	flagType16      = 0x10
	flagFallbackABI = 0x20
	flagHasMetadata = 0x40
)

// ────────────────────────────────────────────────────────────────────────────────
// ABI helper (for the fallback format)
// ────────────────────────────────────────────────────────────────────────────────
var passkeyTupleABI abi.Arguments

func init() {
	// Build at run‑time so we do not need a *.json file.
	bytesType, _ := abi.NewType("bytes", "", nil)
	stringType, _ := abi.NewType("string", "", nil)
	uintType, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	boolType, _ := abi.NewType("bool", "", nil)

	// Struct WebAuthnAuth ‑ components must follow Solidity declaration order.
	authStruct, _ := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Type: "bytes"},   // authenticatorData
		{Type: "string"},  // clientDataJSON
		{Type: "uint256"}, // challengeIndex
		{Type: "uint256"}, // typeIndex
		{Type: "bytes32"}, // r
		{Type: "bytes32"}, // s
	})

	passkeyTupleABI = abi.Arguments{
		{Type: authStruct},  // 0
		{Type: boolType},    // 1
		{Type: bytes32Type}, // 2
		{Type: bytes32Type}, // 3
		{Type: bytes32Type}, // 4
	}

	// Keep the primitives around so the compiler doesn’t eliminate them
	_ = []abi.Type{bytesType, stringType, uintType, bytes32Type}
}

// ────────────────────────────────────────────────────────────────────────────────
// Helpers
// ────────────────────────────────────────────────────────────────────────────────
func uintToBytes(x uint64, n int) []byte {
	out := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		out[i] = byte(x & 0xff)
		x >>= 8
	}
	return out
}
func bytesToUint(b []byte) uint64 {
	var v uint64
	for _, c := range b {
		v = v<<8 + uint64(c)
	}
	return v
}

// ────────────────────────────────────────────────────────────────────────────────
// PasskeySignature
// ────────────────────────────────────────────────────────────────────────────────
type PasskeySignature struct {
	// ---- core fields (affect image‑hash) ----
	requireUserVerification bool
	x, y                    common.Hash
	metadata                common.Hash

	// ---- verification data ----
	authenticatorData []byte
	clientDataJSON    []byte
	challengeIndex    uint64
	typeIndex         uint64
	r, s              common.Hash
}

// ────────────────── Decoding ──────────────────
func DecodePasskeySignature(signature []byte) (PasskeySignature, error) {
	if len(signature) == 0 {
		return PasskeySignature{}, errors.New("signature: empty")
	}
	flags := signature[0]

	// ── ABI fallback ──────────────────────────────────────────────────────────
	if flags&flagFallbackABI != 0 {
		return decodeFallbackABI(signature[1:])
	}

	// ── Compact format ───────────────────────────────────────────────────────
	readCursor := 1
	read := func(count int) ([]byte, error) {
		if readCursor+count > len(signature) {
			return nil, errors.New("signature: truncated")
		}
		out := signature[readCursor : readCursor+count]
		readCursor += count
		return out, nil
	}
	readLen := func(twoBytes bool) (uint64, error) {
		lb := 1
		if twoBytes {
			lb = 2
		}
		b, err := read(lb)
		if err != nil {
			return 0, err
		}
		return bytesToUint(b), nil
	}

	var ps PasskeySignature
	ps.requireUserVerification = flags&flagRequireUV != 0

	if flags&flagHasMetadata != 0 {
		if b, err := read(32); err == nil {
			ps.metadata = common.BytesToHash(b)
		} else {
			return PasskeySignature{}, err
		}
	}

	authLen, err := readLen(flags&flagAuth16 != 0)
	if err != nil {
		return PasskeySignature{}, err
	}
	ps.authenticatorData, err = read(int(authLen))
	if err != nil {
		return PasskeySignature{}, err
	}

	clientLen, err := readLen(flags&flagClient16 != 0)
	if err != nil {
		return PasskeySignature{}, err
	}
	ps.clientDataJSON, err = read(int(clientLen))
	if err != nil {
		return PasskeySignature{}, err
	}

	ps.challengeIndex, err = readLen(flags&flagChallenge16 != 0)
	if err != nil {
		return PasskeySignature{}, err
	}
	ps.typeIndex, err = readLen(flags&flagType16 != 0)
	if err != nil {
		return PasskeySignature{}, err
	}

	if b, err := read(32); err == nil {
		ps.r = common.BytesToHash(b)
	} else {
		return PasskeySignature{}, err
	}
	if b, err := read(32); err == nil {
		ps.s = common.BytesToHash(b)
	} else {
		return PasskeySignature{}, err
	}
	if b, err := read(32); err == nil {
		ps.x = common.BytesToHash(b)
	} else {
		return PasskeySignature{}, err
	}
	if b, err := read(32); err == nil {
		ps.y = common.BytesToHash(b)
	} else {
		return PasskeySignature{}, err
	}

	// Solidity silently ignores extra data – mirror that behaviour.
	return ps, nil
}

// ────────────────── 1) ABI‑fallback decode helper  ──────────────────
func decodeFallbackABI(payload []byte) (PasskeySignature, error) {
	var ps PasskeySignature

	values, err := passkeyTupleABI.Unpack(payload)
	if err != nil {
		return PasskeySignature{}, err
	}

	// 0 → struct WebAuthnAuth
	auth := values[0].([]any)
	ps.authenticatorData = append([]byte(nil), auth[0].([]byte)...)
	ps.clientDataJSON = append([]byte(nil), []byte(auth[1].(string))...)
	ps.challengeIndex = auth[2].(*big.Int).Uint64()
	ps.typeIndex = auth[3].(*big.Int).Uint64()
	if v, ok := auth[4].([32]byte); ok {
		ps.r = common.BytesToHash(v[:])
	} else {
		return PasskeySignature{}, errors.New("ABI decode: r is not bytes32")
	}
	if v, ok := auth[5].([32]byte); ok {
		ps.s = common.BytesToHash(v[:])
	} else {
		return PasskeySignature{}, errors.New("ABI decode: s is not bytes32")
	}

	// 1 → bool  (requireUserVerification)
	ps.requireUserVerification = values[1].(bool)

	// 2,3,4 → x, y, metadata
	if v, ok := values[2].([32]byte); ok {
		ps.x = common.BytesToHash(v[:])
	} else {
		return PasskeySignature{}, errors.New("ABI decode: x is not bytes32")
	}
	if v, ok := values[3].([32]byte); ok {
		ps.y = common.BytesToHash(v[:])
	} else {
		return PasskeySignature{}, errors.New("ABI decode: y is not bytes32")
	}
	if v, ok := values[4].([32]byte); ok {
		ps.metadata = common.BytesToHash(v[:])
	} else {
		return PasskeySignature{}, errors.New("ABI decode: metadata is not bytes32")
	}

	// ⚠️  **No extra consistency check** –
	// Solidity’s `_decodeSignature` accepts any (flags, payload) pair as long
	// as bit 0x20 is set, so we mirror that laissez‑faire behaviour.
	return ps, nil
}

// ────────────────── Encoding ──────────────────
func (s PasskeySignature) Encode() []byte {
	// Decide whether we can use the compact form:
	// ‑ every variable must fit in ≤ 65 535 (2 bytes). If any exceed that
	//   we must fall back to ABI.
	tooLarge := func(v uint64) bool { return v > 0xffff }
	if tooLarge(uint64(len(s.authenticatorData))) ||
		tooLarge(uint64(len(s.clientDataJSON))) ||
		tooLarge(s.challengeIndex) ||
		tooLarge(s.typeIndex) {
		return s.encodeABI()
	}
	return s.encodeCompact()
}

// --- compact ---------------------------------------------------------------
func (s PasskeySignature) encodeCompact() []byte {
	flags := byte(0)
	if s.requireUserVerification {
		flags |= flagRequireUV
	}

	oneByte := func(v uint64) bool { return v <= 0xff }

	// Decide for each field whether we need 1 or 2 bytes and set the flags.
	if !oneByte(uint64(len(s.authenticatorData))) {
		flags |= flagAuth16
	}
	if !oneByte(uint64(len(s.clientDataJSON))) {
		flags |= flagClient16
	}
	if !oneByte(s.challengeIndex) {
		flags |= flagChallenge16
	}
	if !oneByte(s.typeIndex) {
		flags |= flagType16
	}
	if s.metadata != (common.Hash{}) {
		flags |= flagHasMetadata
	}

	var out bytes.Buffer
	out.WriteByte(flags)

	if flags&flagHasMetadata != 0 {
		out.Write(s.metadata.Bytes())
	}

	writeLen := func(v uint64, twoBytes bool) {
		if twoBytes {
			out.Write(uintToBytes(v, 2))
		} else {
			out.Write(uintToBytes(v, 1))
		}
	}

	writeLen(uint64(len(s.authenticatorData)), flags&flagAuth16 != 0)
	out.Write(s.authenticatorData)

	writeLen(uint64(len(s.clientDataJSON)), flags&flagClient16 != 0)
	out.Write(s.clientDataJSON)

	writeLen(s.challengeIndex, flags&flagChallenge16 != 0)
	writeLen(s.typeIndex, flags&flagType16 != 0)

	out.Write(s.r.Bytes())
	out.Write(s.s.Bytes())
	out.Write(s.x.Bytes())
	out.Write(s.y.Bytes())

	return out.Bytes()
}

// --- ABI fallback ----------------------------------------------------------
func (s PasskeySignature) encodeABI() []byte {
	flags := byte(flagFallbackABI)
	if s.requireUserVerification {
		flags |= flagRequireUV
	}
	if s.metadata != (common.Hash{}) {
		flags |= flagHasMetadata
	}

	// Build the WebAuthnAuth tuple (array of interfaces in the right order),
	// then pack the five ABI‑level arguments.
	authTuple := []any{
		[]byte(s.authenticatorData),
		string(s.clientDataJSON),
		new(big.Int).SetUint64(s.challengeIndex),
		new(big.Int).SetUint64(s.typeIndex),
		s.r,
		s.s,
	}

	tuple, err := passkeyTupleABI.Pack(
		authTuple,
		s.requireUserVerification,
		s.x,
		s.y,
		s.metadata,
	)
	if err != nil {
		panic(err) // should never happen with correct argument types
	}

	return append([]byte{flags}, tuple...)
}

// ────────────────── 2) Signature validity ──────────────────
func (s PasskeySignature) IsValid(digest common.Hash) bool {
	// ---- 1) authenticator flags (UP / UV)  ----
	if len(s.authenticatorData) < 33 {
		return false
	}
	if s.authenticatorData[32]&0x01 == 0 {
		return false
	} // UP
	if s.requireUserVerification && s.authenticatorData[32]&0x04 == 0 {
		return false
	} // UV

	// ---- 2) clientDataJSON "type":"webauthn.get" guard ----
	typeLiteral := `"type":"webauthn.get"`
	if int(s.typeIndex)+len(typeLiteral) > len(s.clientDataJSON) ||
		!bytes.Equal([]byte(typeLiteral),
			s.clientDataJSON[s.typeIndex:int(s.typeIndex)+len(typeLiteral)]) {
		return false
	}

	// ---- 3) challenge guard (match Solidity byte‑for‑byte) ----
	encoded := base64.RawURLEncoding.EncodeToString(digest.Bytes())
	prefix := `"challenge":"` + encoded // WITHOUT closing quote
	off := int(s.challengeIndex)
	if off+len(prefix)+1 > len(s.clientDataJSON) { // +1 for the '"' we test next
		return false
	}
	if !bytes.Equal([]byte(prefix), s.clientDataJSON[off:off+len(prefix)]) {
		return false
	}
	if s.clientDataJSON[off+len(prefix)] != '"' { // closing quote (0x22)
		return false
	}

	// ---- 4) P‑256 signature verification (+ low‑S) ----
	hashClient := sha256.Sum256(s.clientDataJSON)
	msgHash := sha256.Sum256(bytes.Join([][]byte{s.authenticatorData, hashClient[:]}, nil))

	r := new(big.Int).SetBytes(s.r.Bytes())
	ss := new(big.Int).SetBytes(s.s.Bytes())
	if ss.Cmp(new(big.Int).Rsh(elliptic.P256().Params().N, 1)) > 0 { // s > N/2
		return false
	}

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(s.x.Bytes()),
		Y:     new(big.Int).SetBytes(s.y.Bytes()),
	}
	return ecdsa.Verify(&pub, msgHash[:], r, ss)
}

// ────────────────── Image‑hash (unchanged) ──────────────────
func (s PasskeySignature) ImageHash() ImageHash {
	var requireUserVerification common.Hash
	if s.requireUserVerification {
		requireUserVerification[len(requireUserVerification)-1] = 1
	}

	return ImageHash{
		Hash: crypto.Keccak256Hash(
			crypto.Keccak256(s.x.Bytes(), s.y.Bytes()),
			crypto.Keccak256(requireUserVerification.Bytes(), s.metadata.Bytes()),
		),
		Preimage: s,
	}
}
