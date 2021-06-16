package contract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestABIParseArtifact(t *testing.T) {
	artifactJSON := `{
		"_format": "hh-sol-artifact-1",
		"contractName": "IERC1271",
		"abi": [
			{
				"inputs": [
					{
						"internalType": "bytes32",
						"name": "_hash",
						"type": "bytes32"
					},
					{
						"internalType": "bytes",
						"name": "_signature",
						"type": "bytes"
					}
				],
				"name": "isValidSignature",
				"outputs": [
					{
						"internalType": "bytes4",
						"name": "magicValue",
						"type": "bytes4"
					}
				],
				"stateMutability": "view",
				"type": "function"
			},
			{
				"inputs": [
					{
						"internalType": "bytes",
						"name": "_data",
						"type": "bytes"
					},
					{
						"internalType": "bytes",
						"name": "_signature",
						"type": "bytes"
					}
				],
				"name": "isValidSignature",
				"outputs": [
					{
						"internalType": "bytes4",
						"name": "magicValue",
						"type": "bytes4"
					}
				],
				"stateMutability": "view",
				"type": "function"
			}
		],
		"bytecode": "0xabcd",
		"deployedBytecode": "0x",
		"linkReferences": {},
		"deployedLinkReferences": {}
	}`

	contractABI, err := ParseArtifactJSON(artifactJSON)
	assert.NoError(t, err)

	assert.Equal(t, "IERC1271", contractABI.Name)
	assert.Len(t, contractABI.Methods, 2)
	assert.Equal(t, "isValidSignature(bytes32,bytes)", contractABI.Methods["isValidSignature"].Sig)
	assert.Len(t, contractABI.Bin, 2) // mock bin
}
