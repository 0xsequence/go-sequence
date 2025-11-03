package prototyp

// EtherNativeTokenAddress represents ETH as an erc20 token address
// according to https://eips.ethereum.org/EIPS/eip-7528
//
// NOTE: throughout sequence we actually use 0x0000000000000000000000000000000000000000
// as the address to represent a native token, so just be mindful when using
// both. We will evolve to treat both as native ETH.
const EtherNativeTokenAddress = Hash("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")

// ZeroAddress is the canonical zero address.
const ZeroAddress = Hash("0x0000000000000000000000000000000000000000")
