package hydrate

// Hydration op-type nibble (high bits of command flag byte). Must match HydrateProxy.sol.
const (
	TypeDataAddress        = 0x01
	TypeDataBalance        = 0x02
	TypeDataERC20Balance   = 0x03
	TypeDataERC20Allowance = 0x04
	TypeTo                 = 0x05
	TypeValue              = 0x06
)

// Hydration address-source nibble (low bits of command flag byte).
const (
	DataSelf       = 0x00
	DataMsgSender  = 0x01
	DataTxOrigin   = 0x02
	DataAnyAddress = 0x03
)

func encodeCommandFlag(typeNibble, dataNibble byte) byte {
	return (typeNibble << 4) | (dataNibble & 0x0f)
}
