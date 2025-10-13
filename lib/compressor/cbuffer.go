package compressor

type CBuffer struct {
	SignatureLevel uint

	Commited []byte
	Pending  []byte

	Refs *References
}

type References struct {
	useContractStorage bool

	usedFlags        map[string]int
	usedStorageFlags map[string]int
}

func NewCBuffer(useStorage bool) *CBuffer {
	return &CBuffer{
		// Start with an empty byte, this
		// will be used as the method when calling the compressor
		// contract.
		Commited: make([]byte, 1),
		Pending:  make([]byte, 0),

		Refs: &References{
			useContractStorage: useStorage,
			usedFlags:          make(map[string]int),
			usedStorageFlags:   make(map[string]int),
		},
	}
}

func (r *References) Copy() *References {
	usedFlags := make(map[string]int, len(r.usedFlags))
	for k, v := range r.usedFlags {
		usedFlags[k] = v
	}

	usedStorageFlags := make(map[string]int, len(r.usedStorageFlags))
	for k, v := range r.usedStorageFlags {
		usedStorageFlags[k] = v
	}

	return &References{
		useContractStorage: r.useContractStorage,

		usedFlags:        usedFlags,
		usedStorageFlags: usedStorageFlags,
	}
}

func (cb *CBuffer) Data() []byte {
	return cb.Commited
}

func (cb *CBuffer) Len() int {
	return len(cb.Commited)
}

func (cb *CBuffer) WriteByte(b byte) {
	cb.Pending = append(cb.Pending, b)
}

func (cb *CBuffer) WriteBytes(b []byte) {
	cb.Pending = append(cb.Pending, b...)
}

func (cb *CBuffer) WriteInt(i uint) {
	cb.WriteByte(byte(i))
}

func (cb *CBuffer) End(uncompressed []byte, t EncodeType) {
	// We need 2 bytes to point to a flag, so any uncompressed value
	// that is 2 bytes or less is not worth saving.
	if len(uncompressed) > 2 {
		rindex := cb.Len()

		switch t {
		case ReadStorage:
		case Stateless:
			cb.Refs.usedFlags[string(uncompressed)] = rindex + 1
		case WriteStorage:
			cb.Refs.usedStorageFlags[string(uncompressed)] = rindex + 1
		default:
		}
	}

	cb.Commited = append(cb.Commited, cb.Pending...)
	cb.Pending = nil
}

type Snapshot struct {
	Commited []byte

	SignatureLevel uint

	Refs *References
}

func (cb *CBuffer) Snapshot() *Snapshot {
	// Create a copy of the commited buffer
	// and of the references.
	com := make([]byte, len(cb.Commited))
	copy(com, cb.Commited)

	refs := cb.Refs.Copy()

	return &Snapshot{
		Commited:       com,
		SignatureLevel: cb.SignatureLevel,
		Refs:           refs,
	}
}

func (cb *CBuffer) Restore(snap *Snapshot) {
	cb.Commited = snap.Commited
	cb.Refs = snap.Refs
	cb.SignatureLevel = snap.SignatureLevel
}
