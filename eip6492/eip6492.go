package eip6492

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/ethrpc"
	"github.com/0xsequence/ethkit/go-ethereum"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/common/hexutil"
)

/* Source of Offchain EIP-6492 validation:

// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.18;


// As per ERC-1271
interface IERC1271Wallet {
  function isValidSignature(bytes32 hash, bytes calldata signature) external view returns (bytes4 magicValue);
}

error ERC1271Revert(bytes error);
error ERC6492DeployFailed(bytes error);

contract UniversalSigValidator {
  bytes32 private constant ERC6492_DETECTION_SUFFIX = 0x6492649264926492649264926492649264926492649264926492649264926492;
  bytes4 private constant ERC1271_SUCCESS = 0x1626ba7e;

  function isValidSigImpl(
    address _signer,
    bytes32 _hash,
    bytes calldata _signature,
    bool allowSideEffects,
    bool deployAlreadyDeployed
  ) public returns (bool) {
    uint contractCodeLen = address(_signer).code.length;
    bytes memory sigToValidate;
    // The order here is striclty defined in https://eips.ethereum.org/EIPS/eip-6492
    // - ERC-6492 suffix check and verification first, while being permissive in case the contract is already deployed; if the contract is deployed we will check the sig against the deployed version, this allows 6492 signatures to still be validated while taking into account potential key rotation
    // - ERC-1271 verification if there's contract code
    // - finally, ecrecover
    bool isCounterfactual = bytes32(_signature[_signature.length-32:_signature.length]) == ERC6492_DETECTION_SUFFIX;
    if (isCounterfactual) {
      address create2Factory;
      bytes memory factoryCalldata;
      (create2Factory, factoryCalldata, sigToValidate) = abi.decode(_signature[0:_signature.length-32], (address, bytes, bytes));

      if (contractCodeLen == 0 || deployAlreadyDeployed) {
        (bool success, bytes memory err) = create2Factory.call(factoryCalldata);
        if (!success) revert ERC6492DeployFailed(err);
      }
    } else {
      sigToValidate = _signature;
    }

    // Try ERC-1271 verification
    if (isCounterfactual || contractCodeLen > 0) {
      try IERC1271Wallet(_signer).isValidSignature(_hash, sigToValidate) returns (bytes4 magicValue) {
        bool isValid = magicValue == ERC1271_SUCCESS;

        // EXPERIMENTAL: This is not part of the EIP-6492 spec *yet*
        // but it may be useful to retry the call making the factory call
        // even if the wallet is already deployed, in case the wallet
        // needs to perform some sort of migration or onchain key rotation
        if (!isValid && !deployAlreadyDeployed && contractCodeLen > 0) {
          return isValidSigImpl(_signer, _hash, _signature, allowSideEffects, true);
        }

        if (contractCodeLen == 0 && isCounterfactual && !allowSideEffects) {
          // if the call had side effects we need to return the
          // result using a `revert` (to undo the state changes)
          assembly {
            mstore(0, isValid)
            revert(31, 1)
          }
        }

        return isValid;
      } catch (bytes memory err) {
        // EXPERIMENTAL: This is not part of the EIP-6492 spec *yet*
        // but it may be useful to retry the call making the factory call
        // even if the wallet is already deployed, in case the wallet
        // needs to perform some sort of migration or onchain key rotation
        if (!deployAlreadyDeployed && contractCodeLen > 0) {
          return isValidSigImpl(_signer, _hash, _signature, allowSideEffects, true);
        }

        revert ERC1271Revert(err);
      }
    }

    // ecrecover verification
    require(_signature.length == 65, 'SignatureValidator#recoverSigner: invalid signature length');
    bytes32 r = bytes32(_signature[0:32]);
    bytes32 s = bytes32(_signature[32:64]);
    uint8 v = uint8(_signature[64]);

    if (v != 27 && v != 28) {
      revert('SignatureValidator: invalid signature v value');
    }

    return ecrecover(_hash, v, r, s) == _signer;
  }

  function isValidSigWithSideEffects(
    address _signer,
    bytes32 _hash,
    bytes calldata _signature
  ) external returns (bool) {
    return this.isValidSigImpl(_signer, _hash, _signature, true, false);
  }

  function isValidSig(
    address _signer,
    bytes32 _hash,
    bytes calldata _signature
  ) external returns (bool) {
    try this.isValidSigImpl(_signer, _hash, _signature, false, false) returns (bool isValid) {
      return isValid;
    } catch (bytes memory error) {
      // in order to avoid side effects from the contract getting deployed, the entire call will revert with a single byte result
      uint len = error.length;
      if (len == 1) {
        return error[0] == 0x01;
        // all other errors are simply forwarded, but in custom formats so that nothing else can revert with a single byte in the call
      } else {
        assembly { revert(error, len) }
      }
    }
  }

  // NOTICE: These functions aren't part of the standard
  // they are helpers that behave like the above functions
  // but they don't revert on failure, instead they return false

  function isValidSigNoThrow(
    address _signer,
    bytes32 _hash,
    bytes calldata _signature
  ) external returns (bool) {
    try this.isValidSigImpl(_signer, _hash, _signature, false, false) returns (bool isValid) {
      return isValid;
    } catch (bytes memory error) {
      // in order to avoid side effects from the contract getting deployed, the entire call will revert with a single byte result
      uint len = error.length;
      if (len == 1) {
        return error[0] == 0x01;
        // all other errors are simply forwarded, but in custom formats so that nothing else can revert with a single byte in the call
      } else {
        // Ignore all other errors and return false
        return false;
      }
    }
  }

  function isValidSigWithSideEffectsNoThrow(
    address _signer,
    bytes32 _hash,
    bytes calldata _signature
  ) external returns (bool) {
    try this.isValidSigImpl(_signer, _hash, _signature, true, false) returns (bool isValid) {
      return isValid;
    } catch (bytes memory error) {
      // Ignore all errors and return false
      return false;
    }
  }
}

// this is a helper so we can perform validation in a single eth_call without pre-deploying a singleton
contract ValidateSigOffchain {
  constructor (address _signer, bytes32 _hash, bytes memory _signature) {
    UniversalSigValidator validator = new UniversalSigValidator();
    bool isValidSig = validator.isValidSigWithSideEffects(_signer, _hash, _signature);
    assembly {
      mstore(0, isValidSig)
      return(31, 1)
    }
  }
}
*/

const EIP_6492_OFFCHAIN_DEPLOY_CODE = "0x608060405234801561001057600080fd5b5060405161124a38038061124a83398101604081905261002f91610124565b600060405161003d906100dd565b604051809103906000f080158015610059573d6000803e3d6000fd5b5090506000816001600160a01b0316638f0684308686866040518463ffffffff1660e01b815260040161008e939291906101fb565b6020604051808303816000875af11580156100ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100d19190610244565b9050806000526001601ff35b610fdc8061026e83390190565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561011b578181015183820152602001610103565b50506000910152565b60008060006060848603121561013957600080fd5b83516001600160a01b038116811461015057600080fd5b6020850151604086015191945092506001600160401b038082111561017457600080fd5b818601915086601f83011261018857600080fd5b81518181111561019a5761019a6100ea565b604051601f8201601f19908116603f011681019083821181831017156101c2576101c26100ea565b816040528281528960208487010111156101db57600080fd5b6101ec836020830160208801610100565b80955050505050509250925092565b60018060a01b0384168152826020820152606060408201526000825180606084015261022e816080850160208701610100565b601f01601f191691909101608001949350505050565b60006020828403121561025657600080fd5b8151801515811461026657600080fd5b939250505056fe608060405234801561001057600080fd5b50610fbc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c806376be4cea1161005057806376be4cea146100a65780638f068430146100b957806398ef1ed8146100cc57600080fd5b80631c6453271461006c5780633d787b6314610093575b600080fd5b61007f61007a366004610ad4565b6100df565b604051901515815260200160405180910390f35b61007f6100a1366004610ad4565b61023d565b61007f6100b4366004610b3e565b61031e565b61007f6100c7366004610ad4565b6108e1565b61007f6100da366004610ad4565b61096e565b6040517f76be4cea00000000000000000000000000000000000000000000000000000000815260009030906376be4cea9061012890889088908890889088908190600401610bc3565b6020604051808303816000875af1925050508015610181575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261017e91810190610c45565b60015b610232573d8080156101af576040519150601f19603f3d011682016040523d82523d6000602084013e6101b4565b606091505b508051600181900361022757816000815181106101d3576101d3610c69565b6020910101517fff00000000000000000000000000000000000000000000000000000000000000167f0100000000000000000000000000000000000000000000000000000000000000149250610235915050565b600092505050610235565b90505b949350505050565b6040517f76be4cea00000000000000000000000000000000000000000000000000000000815260009030906376be4cea906102879088908890889088906001908990600401610bc3565b6020604051808303816000875af19250505080156102e0575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526102dd91810190610c45565b60015b610232573d80801561030e576040519150601f19603f3d011682016040523d82523d6000602084013e610313565b606091505b506000915050610235565b600073ffffffffffffffffffffffffffffffffffffffff87163b6060827f64926492649264926492649264926492649264926492649264926492649264928888610369602082610c98565b610375928b9290610cd8565b61037e91610d02565b1490508015610484576000606089828a610399602082610c98565b926103a693929190610cd8565b8101906103b39190610e18565b955090925090508415806103c45750865b1561047d576000808373ffffffffffffffffffffffffffffffffffffffff16836040516103f19190610eb2565b6000604051808303816000865af19150503d806000811461042e576040519150601f19603f3d011682016040523d82523d6000602084013e610433565b606091505b50915091508161047a57806040517f9d0d6e2d0000000000000000000000000000000000000000000000000000000081526004016104719190610f18565b60405180910390fd5b50505b50506104be565b87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509294505050505b80806104ca5750600083115b156106bb576040517f1626ba7e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b1690631626ba7e90610523908c908690600401610f2b565b602060405180830381865afa92505050801561057a575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820190925261057791810190610f44565b60015b61060f573d8080156105a8576040519150601f19603f3d011682016040523d82523d6000602084013e6105ad565b606091505b50851580156105bc5750600084115b156105db576105d08b8b8b8b8b600161031e565b9450505050506108d7565b806040517f6f2a95990000000000000000000000000000000000000000000000000000000081526004016104719190610f18565b7fffffffff0000000000000000000000000000000000000000000000000000000081167f1626ba7e000000000000000000000000000000000000000000000000000000001480158161065f575086155b801561066b5750600085115b1561068b5761067f8c8c8c8c8c600161031e565b955050505050506108d7565b841580156106965750825b80156106a0575087155b156106af57806000526001601ffd5b94506108d79350505050565b6041871461074b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5369676e617475726556616c696461746f72237265636f7665725369676e657260448201527f3a20696e76616c6964207369676e6174757265206c656e6774680000000000006064820152608401610471565b600061075a6020828a8c610cd8565b61076391610d02565b90506000610775604060208b8d610cd8565b61077e91610d02565b905060008a8a604081811061079557610795610c69565b919091013560f81c915050601b81148015906107b557508060ff16601c14155b15610842576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f5369676e617475726556616c696461746f723a20696e76616c6964207369676e60448201527f617475726520762076616c7565000000000000000000000000000000000000006064820152608401610471565b6040805160008152602081018083528e905260ff831691810191909152606081018490526080810183905273ffffffffffffffffffffffffffffffffffffffff8e169060019060a0016020604051602081039080840390855afa1580156108ad573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff161496505050505050505b9695505050505050565b6040517f76be4cea00000000000000000000000000000000000000000000000000000000815260009030906376be4cea9061092b9088908890889088906001908990600401610bc3565b6020604051808303816000875af115801561094a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102329190610c45565b6040517f76be4cea00000000000000000000000000000000000000000000000000000000815260009030906376be4cea906109b790889088908890889088908190600401610bc3565b6020604051808303816000875af1925050508015610a10575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610a0d91810190610c45565b60015b610232573d808015610a3e576040519150601f19603f3d011682016040523d82523d6000602084013e610a43565b606091505b5080516001819003610a6257816000815181106101d3576101d3610c69565b8082fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610a8857600080fd5b50565b60008083601f840112610a9d57600080fd5b50813567ffffffffffffffff811115610ab557600080fd5b602083019150836020828501011115610acd57600080fd5b9250929050565b60008060008060608587031215610aea57600080fd5b8435610af581610a66565b935060208501359250604085013567ffffffffffffffff811115610b1857600080fd5b610b2487828801610a8b565b95989497509550505050565b8015158114610a8857600080fd5b60008060008060008060a08789031215610b5757600080fd5b8635610b6281610a66565b955060208701359450604087013567ffffffffffffffff811115610b8557600080fd5b610b9189828a01610a8b565b9095509350506060870135610ba581610b30565b91506080870135610bb581610b30565b809150509295509295509295565b73ffffffffffffffffffffffffffffffffffffffff8716815285602082015260a060408201528360a0820152838560c0830137600060c085830181019190915292151560608201529015156080820152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016909101019392505050565b600060208284031215610c5757600080fd5b8151610c6281610b30565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b81810381811115610cd2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60008085851115610ce857600080fd5b83861115610cf557600080fd5b5050820193919092039150565b80356020831015610cd2577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610d7e57600080fd5b813567ffffffffffffffff80821115610d9957610d99610d3e565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610ddf57610ddf610d3e565b81604052838152866020858801011115610df857600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215610e2d57600080fd5b8335610e3881610a66565b9250602084013567ffffffffffffffff80821115610e5557600080fd5b610e6187838801610d6d565b93506040860135915080821115610e7757600080fd5b50610e8486828701610d6d565b9150509250925092565b60005b83811015610ea9578181015183820152602001610e91565b50506000910152565b60008251610ec4818460208701610e8e565b9190910192915050565b60008151808452610ee6816020860160208601610e8e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610c626020830184610ece565b8281526040602082015260006102356040830184610ece565b600060208284031215610f5657600080fd5b81517fffffffff0000000000000000000000000000000000000000000000000000000081168114610c6257600080fdfea26469706673582212201a72aed4b15ffb05b6502997a9bb655992e06590bd26b336dfbb153d7ff6f34b64736f6c63430008120033"
const EIP_6492_SUFFIX = "0x6492649264926492649264926492649264926492649264926492649264926492"

var EIP6492MagicBytes = hexutil.MustDecode(EIP_6492_SUFFIX)

func IsEIP6492Signature(signature []byte) bool {
	return bytes.HasSuffix(signature, EIP6492MagicBytes)
}

func DecodeEIP6492Signature(signature []byte) (common.Address, []byte, []byte, error) {
	if !IsEIP6492Signature(signature) {
		return common.Address{}, nil, nil, fmt.Errorf("not an eip6492 signature")
	}

	// Strip the magic bytes
	signature = signature[:len(signature)-32]

	var create2Factory common.Address
	var factoryCalldata []byte
	var sigToValidate []byte

	if err := ethcoder.AbiDecoder([]string{"address", "bytes", "bytes"}, signature, []interface{}{&create2Factory, &factoryCalldata, &sigToValidate}); err != nil {
		return common.Address{}, nil, nil, err
	}

	return create2Factory, factoryCalldata, sigToValidate, nil
}

func ValidateEIP6492Offchain(
	ctx context.Context,
	provider *ethrpc.Provider,
	signer common.Address,
	hash common.Hash,
	signature []byte,
	block *big.Int,
) (bool, error) {
	base, err := ethcoder.AbiCoder(
		[]string{"address", "bytes32", "bytes"}, []interface{}{
			signer, hash, signature,
		},
	)

	if err != nil {
		return false, err
	}

	packed, err := ethcoder.SolidityPack([]string{"bytes", "bytes"}, []interface{}{
		hexutil.MustDecode(EIP_6492_OFFCHAIN_DEPLOY_CODE), base,
	})

	if err != nil {
		return false, err
	}

	msg := ethereum.CallMsg{
		Data: packed,
	}

	result, err := provider.CallContract(ctx, msg, block)
	if err != nil {
		return false, err
	}

	// Check if the result is '0x01'
	expectedResult := hexutil.MustDecode("0x01")
	return bytes.Equal(result, expectedResult), nil
}
