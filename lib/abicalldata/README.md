# ABI calldata

Helpers for **Ethereum ABI-encoded calldata**: head layout, static argument words, and dynamic `bytes`/`string` tails. Also defines **`ByteRange`**, **`Selector`**, and a fluent **`Path`** for locating byte spans inside a v3 **`CallsPayload`**.

## Calldata layout

Offsets are **from the start of calldata** (bytes `0..3` are the 4-byte function selector).

- **`AbiHeadWords(t)`** — number of 32-byte words an ABI type occupies in the **head** (dynamic types count as one offset word).
- **`CalldataStaticWord(method, argIndex)`** — `start` and `length` (multiple of 32) for a **static** argument’s encoded words in `calldata`.
- **`CalldataBytesContent(calldata, method, argIndex)`** — `start`/`length` of the **raw** `bytes`/`string` payload (no length word, no tail padding).
- **`CalldataBytesEncoded(calldata, method, argIndex)`** — `start`/`length` of the full **tail slice**: 32-byte length word + content + padding to 32-byte boundary.

```go
method := contractABI.Methods["transfer"]
start, length, err := abicalldata.CalldataStaticWord(method, 1) // e.g. uint256 amount
if err != nil {
    return err
}
argBytes := calldata[start : start+length]
```

For nested calldata (e.g. a `bytes` argument whose body is another contract call), slice to the inner calldata and call these helpers again with the inner method’s `abi.Method`.

## Calls payload selectors

- **`ByteRange`** — `{CallIndex, Offset, Size}` into `payload.Calls[CallIndex].Data` (that field is the call’s calldata, typically including the inner 4-byte selector).
- **`Selector`** — `Resolve(*v3.CallsPayload) ([]ByteRange, error)` plus `String()` for errors.
- **`NewRangeSelector(callIndex, offset, size)`** — fixed range, validated on resolve.

```go
sel := abicalldata.NewRangeSelector(0, 0x24, 32)
ranges, err := sel.Resolve(payload)
```

## Path builder (`NewPath`)

**`*Path`** is a **`Selector`**: chain steps to walk from a top-level call into nested packed calls and ABI argument slots, then call **`.AsSelector()`** for APIs that take a **`Selector`**.

Typical steps:

- **`.CallData(i)`** — start from `payload.Calls[i].Data`.
- **`.ABI(contractABI, method)`** — bind the current range to that method (checks the 4-byte selector).
- **`.ArgSlot(name)`** / **`.ArgSlotIndex(i)`** — static argument word(s) in the current frame.
- **`.ArgBytesData(name)`** / **`.ArgBytesDataIndex(i)`** — inner payload of a `bytes`/`string` argument (clears ABI context; rebind with `.ABI` before further arg steps).
- **`.ArgBytesEncoded(name)`** — full ABI-encoded tail for that dynamic argument.
- **`.EncodedCallsPayload()`** — treat the active range as v3 packed calls; clear ABI context.
- **`.EncodedCallData(j)`** — select packed call `j`’s calldata within that layout.
- **`.Slice(offset, size)`** — byte slice within the active range.

```go
sel := abicalldata.NewPath().
    CallData(0).
    ABI(&outerABI, "hydrateExecute").
    ArgBytesData("payload").
    EncodedCallsPayload().
    EncodedCallData(0).
    ABI(&tokenABI, "permit").
    ArgSlot("value").
    AsSelector()
```

After any step that narrows the range or switches to a new byte frame (including **`.Slice`**, **`.ArgBytesData`**, **`.ArgBytesDataIndex`**, **`.ArgBytesEncoded`**, **`.EncodedCallsPayload`**, **`.EncodedCallData`**), ABI context is cleared. Call **`.ABI(...)`** again before **`.ArgSlot`**, **`.ArgSlotIndex`**, **`.ArgBytesData`**, **`.ArgBytesDataIndex`**, or **`.ArgBytesEncoded`** on the new frame.

## Packed calls layout

**`ParsePackedCalls(packed []byte)`** returns a **`PackedCallsLayout`** describing where each call’s calldata sits inside the encoded packed-calls blob (as produced by `CallsPayload.Encode`). **`CallData`** entries use **`Span`** (`Start`, `Len`; `Start == -1` when that call has no calldata). The path step **`.EncodedCallData(i)`** uses this parser internally.
