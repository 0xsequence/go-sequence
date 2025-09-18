# v0.60.0

- Adds support for Sequence Wallet v3 contracts
- Keeps backward compatibility with Sequence Wallet v1/v2 contracts

## API Breaking changes in go-sequence

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
