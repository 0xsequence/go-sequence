# Malleable Sapient

Build MalleableSapient signatures by locating byte ranges in call data, and optionally compute the image hash.

`ByteRange`, `Selector`, `RangeSelector`, and `NewRangeSelector` are defined in [`github.com/0xsequence/go-sequence/lib/abicalldata`](https://github.com/0xsequence/go-sequence/tree/master/lib/abicalldata). `Path` implements `abicalldata.Selector`. This package re-exports those types as aliases (and delegates `NewRangeSelector`) for compatibility with code that imported them from `malleable`.

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

After any step that narrows the active range or descends into a new byte
frame, ABI context is cleared. This includes steps like `.Slice(...)`,
`.ArgBytesData(...)`, `.ArgBytesDataIndex(...)`, `.ArgBytesEncoded(...)`,
`.EncodedCallsPayload()`, and `.EncodedCallData(...)`.

Call `.ABI(...)` again before using `.ArgSlot(...)`, `.ArgSlotIndex(...)`,
`.ArgBytesData(...)`, `.ArgBytesDataIndex(...)`, or `.ArgBytesEncoded(...)`
against the new frame.

Compute the image hash:

```go
hash, err := malleable.ComputeImageHash(payload, sig, chainID)
```
