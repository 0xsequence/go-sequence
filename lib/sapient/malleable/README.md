# Malleable Sapient (Go)

Build MalleableSapient signatures by locating byte ranges in call data, and optionally compute the image hash.

## Usage

```go
payload := v3.NewCallsPayload(...)

permitValue := malleable.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(0).
    ABI(erc2612ABI, "permit").
    ArgSlot("value").
    AsSelector()

transferValue := malleable.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(1).
    ABI(erc20ABI, "transferFrom").
    ArgSlot("_value").
    AsSelector()

b := malleable.NewBuilder(payload, &malleable.BuilderOptions{
    ValidateRepeats:     true,
    MergeAdjacentStatic: true,
})

b.Repeat(permitValue, transferValue) // repeat constraint

// mark other malleable fields
b.Malleable(malleable.NewPath().
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
value := malleable.NewPath().
    CallData(0).
    ABI(erc20ABI, "transferFrom").
    ArgSlotIndex(2).
    AsSelector()
```

Compute the image hash:

```go
hash, err := malleable.ComputeImageHash(payload, sig, chainID)
```
