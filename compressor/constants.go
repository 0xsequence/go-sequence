package compressor

import "github.com/0xsequence/ethkit/go-ethereum/common"

const (
	METHOD_EXECUTE_1 uint = iota
	METHOD_EXECUTE_N
	METHOD_READ_ADDRESS
	METHOD_READ_BYTES32
	METHOD_READ_SIZES
	METHOD_READ_STORAGE_SLOTS
	METHOD_DECOMPRESS_1
	METHOD_DECOMPRESS_N
)

const (
	FLAG_READ_POWER_OF_10_MISC uint = iota
	FLAG_READ_BYTES32_1_BYTES
	FLAG_READ_BYTES32_2_BYTES
	FLAG_READ_BYTES32_3_BYTES
	FLAG_READ_BYTES32_4_BYTES
	FLAG_READ_BYTES32_5_BYTES
	FLAG_READ_BYTES32_6_BYTES
	FLAG_READ_BYTES32_7_BYTES
	FLAG_READ_BYTES32_8_BYTES
	FLAG_READ_BYTES32_9_BYTES
	FLAG_READ_BYTES32_10_BYTES
	FLAG_READ_BYTES32_11_BYTES
	FLAG_READ_BYTES32_12_BYTES
	FLAG_READ_BYTES32_13_BYTES
	FLAG_READ_BYTES32_14_BYTES
	FLAG_READ_BYTES32_15_BYTES
	FLAG_READ_BYTES32_16_BYTES
	FLAG_READ_BYTES32_17_BYTES
	FLAG_READ_BYTES32_18_BYTES
	FLAG_READ_BYTES32_19_BYTES
	FLAG_READ_BYTES32_20_BYTES
	FLAG_READ_BYTES32_21_BYTES
	FLAG_READ_BYTES32_22_BYTES
	FLAG_READ_BYTES32_23_BYTES
	FLAG_READ_BYTES32_24_BYTES
	FLAG_READ_BYTES32_25_BYTES
	FLAG_READ_BYTES32_26_BYTES
	FLAG_READ_BYTES32_27_BYTES
	FLAG_READ_BYTES32_28_BYTES
	FLAG_READ_BYTES32_29_BYTES
	FLAG_READ_BYTES32_30_BYTES
	FLAG_READ_BYTES32_31_BYTES
	FLAG_READ_BYTES32_32_BYTES
	FLAG_SAVE_ADDRESS
	FLAG_SAVE_BYTES32
	FLAG_READ_ADDRESS_2
	FLAG_READ_ADDRESS_3
	FLAG_READ_ADDRESS_4
	FLAG_READ_EXECUTE
	FLAG_READ_BYTES32_2
	FLAG_READ_BYTES32_3
	FLAG_READ_BYTES32_4
	FLAG_READ_POW_10_MANTISSA
	FLAG_READ_N_BYTES
	FLAG_READ_POWER_OF_2
	FLAG_ABI_0_PARAM
	FLAG_ABI_1_PARAM
	FLAG_ABI_2_PARAMS
	FLAG_ABI_3_PARAMS
	FLAG_ABI_4_PARAMS
	FLAG_ABI_5_PARAMS
	FLAG_ABI_6_PARAMS
	FLAG_NESTED_N_FLAGS_8
	FLAG_NESTED_N_FLAGS_16
	FLAG_SIGNATURE_W0
	FLAG_SIGNATURE_W1
	FLAG_SIGNATURE_W2
	FLAG_SIGNATURE_W3
	FLAG_SIGNATURE_W4
	FLAG_ADDRESS_W0
	FLAG_ADDRESS_W1
	FLAG_ADDRESS_W2
	FLAG_ADDRESS_W3
	FLAG_ADDRESS_W4
	FLAG_NODE
	FLAG_BRANCH
	FLAG_SUBDIGEST
	FLAG_NESTED
	FLAG_DYNAMIC_SIGNATURE
	FLAG_S_SIG_NO_CHAIN
	FLAG_S_SIG
	FLAG_S_L_SIG_NO_CHAIN
	FLAG_S_L_SIG
	FLAG_READ_CHAINED
	FLAG_READ_CHAINED_L
	FLAG_READ_DYNAMIC_ABI
	FLAG_NO_OP
	FLAG_MIRROR_FLAG
	FLAG_COPY_CALLDATA
	FLAG_READ_STORE_FLAG
)

const LITERAL_ZERO = FLAG_READ_STORE_FLAG + 1
const MAX_LITERAL = 0xff - LITERAL_ZERO

const BYTES4_TABLE = "00000000a9059cbb095ea7b37ff36ab538ed173918cbafe5202ee0edfb3bdb41e2bbb158ab834bab6ea056a923b872dda68a76cc5f5755298803dbeea22cb465c89e43612da0340990411a321cff79cd223da1ba2e1a7d4df305d71939125215d0e30db0f7654176a694fc3a1a695230b6b55f25791ac94764887334c658695c441a3e704f1d48324a25d94aa454dfa9c18a84bce2b39746178979aec9807539ddd81f82a8a41c70cf557fe33d18b9129cec63924ab0d1900dcd7a6cd9627aa49149bafe672a9400dfbe4a31c6bf3262f242432ae5ab4da240c10f192e7ba6efc23e1a211aa3a008d6b347f7ded9382a00000003e9fad8eefaebafa8ae169a50e8e3370041fe00a0fa558b712e95b6c8c48fdfca000000006a80c33f627dd56a5c11d7954946e2065e83b463ca722cdcfb90b32000000008f7c1e582a32fe0a1db006a75000000010002191ce6d66ac8a0712d685d5d442296aa7368d3392ddf0ea5812fa5d754d1d29dff129979ef457901451ca64f797659d667a500032587865a6b4f379607f57c02520082d2697fc11695488758a5f34e71d92d9ddd67ba183d4e0b8f69c188e3dec8fbedc9af952d2da8066a761202a9b1d507ca120b1ff14fcbc8961c9ae442842e0e2195995c94b918de608060405174e8532505c3d98568523a0e89439bd149d05cefef39a14ce6931a000225879bfcb236415565b0454a2ab3ce558087f7a1696342966c688b4cb0ec4faa8a26e4a76726e8eda9df1519cdeb356282bfe17376b5009952eb3d7989fe34b0793b38bcdfc0f053566e02751cecc01a8c84f463e18e3cd18ca029ada03907d6b3483805550fa59f3e0c89bbb8b2c5ebeaec4997adb6f5e54063761610fcb88a802f3ccfd60b2e2d726ca4202615b44848f51610ca95bcf64e0579b177ec22895118ed436a474d474898c0f4ed31b967cb0ca6e158f8db7fd4089120491ca415bcad8201aa3f6e5110ae5312ea8e3df02124b77d239ba67a6a45156e29f6241735bbd017e8c73f7658fd86b2ecc4c44193c39bc12042d96a094a13d98d135d4c66a3ad4451a32e17de789ec9b36be47d166cbfff3b87f884e54a0b020003ad58bdd147e7ef24bad42590c8fd6ed002032587c6427474f6162b01baa2abde1ff013f11846eac55915d806f6aa658b00024a9c564a515869328dec4454b20df5298aca853828b6f06427e5b6b4af05f3fef3a352a438b81249c58bfeab2e5af9d83bb568c2c5fb02022587d586d8e0db254e5005eec2890e7527028f4af52f6a627842508c1dbd0f694584a6417ed63049105d1e9a6950d9caed120103258748d5c7e3be389d577430e0c649b780f00af49149d508e6238e1e280cae47bea8683fa88d5db3b4df1e83409a852a12e3c2998238343009a2daa6d5560f0439589c1298a06aa1e6d24d559317"

func LoadBytes4() map[string]uint {
	btable := common.Hex2Bytes(BYTES4_TABLE)
	table := make(map[string]uint)

	for i := uint(0); i < uint(len(btable)); i += 4 {
		table[string(btable[i:i+4])] = i / 4
	}

	return table
}
