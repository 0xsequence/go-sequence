# Hydrate

Build `hydratePayload` bytes for **HydrateProxy** (`hydrateExecute` / `hydrateExecuteAndSweep` in trails-contracts) by resolving patch targets with [`abicalldata`](https://github.com/0xsequence/go-sequence/tree/master/lib/abicalldata) **`Selector`** and **`ByteRange`** instead of hand-counting calldata offsets.

`hydrate` depends only on `abicalldata` for selector types. The example below builds an `abicalldata.Selector` using `Path` from `lib/sapient/malleable`; fixed offsets can use `abicalldata.NewRangeSelector` instead.

## Usage

```go
payload := v3.NewCallsPayload(...)

// Selector must resolve to exactly one range on the target call (same index as ForCall).
permitOwner := malleable.NewPath().
    CallData(0).
    ABI(trailsABI, "hydrateExecute").
    ArgBytesData("packedPayload").
    EncodedCallsPayload().
    EncodedCallData(0).
    ABI(erc2612ABI, "permit").
    ArgSlot("owner").
    AsSelector()

b := hydrate.NewBuilder(&payload)

if err := b.ForCall(0).DataAddress(permitOwner, hydrate.SourceSelf()); err != nil {
    return err
}

hydratePayload, err := b.Build()
if err != nil {
    return err
}
// nil/empty hydratePayload means "no hydration" (contract skips hydration).

calldata, err := hydrate.PackHydrateExecute(packedPayload, hydratePayload)
```

With sweep after execution:

```go
calldata, err := hydrate.PackHydrateExecuteAndSweep(
    packedPayload,
    hydratePayload,
    sweepTarget,
    tokensToSweep,
    sweepNative,
)
```

If ABI parameters are unnamed, use index-based steps such as `ArgSlotIndex` / `ArgBytesDataIndex` when your path builder provides them.

## Call sections and ordering

- Use `ForCall(tindex)` for each packed call you want to hydrate. Multiple methods on the same `ForCall` append commands to that call's section.
- `Build()` emits sections in **ascending** `tindex` order, each as: `[tindex byte][commands…][0x00]`. That order matches how the proxy walks the stream while executing calls; sections must not be reordered arbitrarily.

## Address sources

Data patches and `CallTo` / `CallValue` use a low nibble on the command byte; optional literals append 20 bytes when needed:

| Helper | Meaning at execution time |
|--------|---------------------------|
| `SourceSelf()` | `address(this)` (the proxy) |
| `SourceMsgSender()` | `msg.sender` |
| `SourceTxOrigin()` | `tx.origin` |
| `SourceAddress(a)` | fixed `a` (encoded after the flag) |

## Selectors

Each `Data*` method requires an `abicalldata.Selector` that resolves to **exactly one** `abicalldata.ByteRange`, and that range's `CallIndex` must equal the `ForCall` index. The range's offset within that call's `data` becomes the patch offset (`uint16`); the builder checks that 20-byte (address) or 32-byte (uint256) replacements fit.

For unit tests or fixed layouts, `abicalldata.NewRangeSelector(callIndex, offset, size)` returns a selector backed by an explicit range.

When using a multi-step path into nested calldata, ABI context is cleared after any step that narrows the active range or switches to a new byte frame (including `.Slice`, `.ArgBytesData`, `.ArgBytesDataIndex`, `.ArgBytesEncoded`, `.EncodedCallsPayload`, `.EncodedCallData`). Call `.ABI(contractABI, method)` again before `.ArgSlot`, `.ArgSlotIndex`, `.ArgBytesData`, `.ArgBytesDataIndex`, or `.ArgBytesEncoded` on that frame.
