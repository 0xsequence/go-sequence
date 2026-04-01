# Malleable Sapient

Build MalleableSapient signatures by locating byte ranges in call data, and optionally compute the image hash.

**ABI calldata paths** use **`abicalldata.NewPath()`** (see [`lib/abicalldata`](https://github.com/0xsequence/go-sequence/tree/master/lib/abicalldata)); this package supplies **`Builder`**, repeat constraints, and **`ComputeImageHash`**.

## Usage

```go
payload := v3.NewCallsPayload(...)

permitValue := abicalldata.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(0).
    ABI(erc2612ABI, "permit").
    ArgSlot("value").
    AsSelector()

transferValue := abicalldata.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(1).
    ABI(erc20ABI, "transferFrom").
    ArgSlot("_value").
    AsSelector()

b := malleable.NewBuilder(&payload, &malleable.BuilderOptions{
    ValidateRepeats:     true,
    MergeAdjacentStatic: true,
})

b.Repeat(permitValue, transferValue) // repeat constraint

// mark other malleable fields
b.Malleable(abicalldata.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(0).
    ABI(erc2612ABI, "permit").
    ArgSlot("deadline").
    AsSelector(),
)

sig, _, err := b.Build()
```

If ABI params are unnamed, use index-based selectors:

```go
value := abicalldata.NewPath().
    CallData(0).
    ABI(erc20ABI, "transferFrom").
    ArgSlotIndex(2).
    AsSelector()
```

Path semantics (when to rebind **`.ABI`** after nested steps) are documented in the **abicalldata** README.

Compute the image hash:

```go
hash, err := malleable.ComputeImageHash(&payload, sig, chainID)
```
