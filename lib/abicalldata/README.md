# ABI calldata

Helpers for **Ethereum ABI-encoded calldata**: head layout, static argument words, and dynamic `bytes`/`string` tails. Also defines **`ByteRange`** and **`Selector`** for locating byte spans inside a v3 **`CallsPayload`**.

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

**`Selector`** is implemented by types in other packages that walk `CallsPayload` and nested calldata; they typically use the functions above to turn `abi.Method` + argument index into `ByteRange` updates. **`NewRangeSelector`** is the built-in implementation for a fixed call index, offset, and length.
