# v0.60.0

- Adds support for Sequence Wallet v3 contracts
- Keeps backward compatibility with Sequence Wallet v1/v2 contracts

## API Breaking changes in go-sequence

### `sequence.DecodeExecdata()`

```diff
-func DecodeExecdata(data []byte) (Transactions, *big.Int, []byte, error)
+func DecodeRawTransactions(data []byte) (Transactions, error)
```

```diff
-sequenceTransactions, _, _, err := sequence.DecodeExecdata(data)
+ sequenceTransactions, err := sequence.DecodeRawTransactions(data)
```

### `sequence.AddressFromImageHash()`

```diff
-   address, err := sequence.AddressFromImageHash(imageHash.String(), seqContext)
+   address, err := sequence.AddressFromImageHash(imageHash, seqContext)
```

### `sequence.Relayer` interface

```diff
 package sequence

 import (
     "github.com/0xsequence/ethkit/go-ethereum/common"
 )
 
 type Relayer interface {
-	EstimateGasLimits(ctx context.Context, walletConfig core.WalletConfig, walletContext WalletContext, txns Transactions) (Transactions, error)
 
-	Simulate(ctx context.Context, txs *SignedTransactions) ([]*RelayerSimulateResult, error)
+	Simulate(ctx context.Context, wallet common.Address, transactions Transactions) ([]*SimulateResult, error)
 
	// ...
 }
```

`sequence.RelayerSimulateResult` type has been replaced by `sequence.SimulateResult`:

```diff
-type RelayerSimulateResult struct {
- 	Executed  bool
- 	Succeeded bool
- 	Result    *string
- 	Reason    *string
- 	GasUsed   uint
- 	GasLimit  uint
- }
 
+ type SimulateResult struct {
+ 	simulator.Result
+ 	GasLimit uint64
+ }
```

### `relayer.RelayerClient` interface

```diff
 type RelayerClient interface {
- 	SendMetaTxn(ctx context.Context, call *MetaTxn, quote *string, projectID *uint64) (bool, string, error)
+	SendMetaTxn(ctx context.Context, call *MetaTxn, quote *string, projectID *uint64, preconditions []*IntentPrecondition) (bool, string, error)
 
	// ...
 }
```

### `core`

```diff
 type SignerSignature struct {
 	Signer    common.Address
-	Subdigest Subdigest
 	Type      SignerSignatureType
 	Signature []byte
 	Error     error
 	Preimage ImageHashable
 }

- // A Subdigest is a hash signed by a Sequence wallet's signers.
- // Used for type safety and preimage recovery.
- type Subdigest struct {
- 	common.Hash
- 
- 	// Digest is the preimage of the subdigest, nil if unknown.
- 	Digest *Digest
- 
- 	// Wallet is the target wallet of the subdigest, nil if unknown.
- 	Wallet *common.Address
- 
- 	// ChainID is the target chain ID of the subdigest, nil if unknown.
- 	ChainID *big.Int
- 
- 	// EthSignPreimage is the preimage of the eth_sign subdigest, nil if unknown.
- 	EthSignPreimage *Subdigest
- }

- type Digest struct {
- 	common.Hash
- 
- 	// Preimage is the preimage of the digest, nil if unknown.
- 	Preimage []byte
- }
+type PayloadDigest struct {
+	common.Hash
+
+	Address_ common.Address
+	ChainID_ *big.Int
+
+	Payload Payload
+}

+type Payload interface {
+	Address() common.Address
+	ChainID() *big.Int
+
+	Digest() PayloadDigest
+}

-type SignerSignatures map[common.Address]map[common.Hash]SignerSignature
+type SignerSignatures map[common.Address]struct {
+	Signature         *SignerSignature
+	SapientSignatures map[common.Hash]SignerSignature
+}
```

`core.Digest` & `core.Subdigest` were replaced by `core.DigestPayload`:

```diff
-subdigest := core.Subdigest{
-	Hash:    common.BytesToHash(msgDigest),
-	Digest:  nil,
-	Wallet:  nil,
-	ChainID: nil,
-}
+subdigest := core.PayloadDigest{Hash: common.BytesToHash(msgDigest)}

-imageHash, weight, err := decoded.RecoverSubdigest(ctx, subdigest, nil)
+imageHash, weight, err := decoded.Recover(ctx, subdigest, nil)
```

```diff
-dig := core.Digest{
+dig := core.PayloadDigest{
	Hash: digest,
}
-subdig := dig.Subdigest(s.Address, chainID)

```

