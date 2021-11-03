package sequence

import (
	"fmt"
	"sort"

	"github.com/0xsequence/ethkit/ethcoder"
	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/0xsequence/ethkit/go-ethereum/crypto"
)

type WalletConfig struct {
	Threshold uint16              `json:"threshold"`
	Signers   WalletConfigSigners `json:"signers"`
}

type WalletConfigSigner struct {
	Weight  uint8          `json:"weight"`
	Address common.Address `json:"address"`
}

type WalletConfigSigners []WalletConfigSigner

func (s WalletConfigSigners) Len() int { return len(s) }
func (s WalletConfigSigners) Less(i, j int) bool {
	return s[i].Address.Hash().Big().Cmp(s[j].Address.Hash().Big()) < 0
}
func (s WalletConfigSigners) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s WalletConfigSigners) GetWeightByAddress(address common.Address) (uint8, bool) {
	for _, signer := range s {
		if signer.Address == address {
			return signer.Weight, true
		}
	}
	return 0, false
}

func AddressFromWalletConfig(walletConfig WalletConfig, context WalletContext) (common.Address, error) {
	imageHash, err := ImageHashOfWalletConfig(walletConfig)
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromWalletConfig: %w", err)
	}
	return AddressFromImageHash(imageHash, context)
}

func AddressFromImageHash(imageHash string, context WalletContext) (common.Address, error) {
	mainModule32 := [32]byte{}
	copy(mainModule32[12:], context.MainModuleAddress.Bytes())

	codePack, err := ethcoder.SolidityPack([]string{"bytes", "bytes32"}, []interface{}{walletContractBytecode, mainModule32})
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	codeHash := crypto.Keccak256(codePack)

	hashPack, err := ethcoder.SolidityPack(
		[]string{"bytes1", "address", "bytes32", "bytes32"},
		[]interface{}{[]byte{0xff}, context.FactoryAddress, common.FromHex(imageHash), codeHash},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("sequence, AddressFromImageHash: %w", err)
	}
	hash := crypto.Keccak256(hashPack)[12:]

	return common.BytesToAddress(hash), nil
}

func ImageHashOfWalletConfig(walletConfig WalletConfig) (string, error) {
	imageHash, err := ImageHashOfWalletConfigBytes(walletConfig)
	if err != nil {
		return "", err
	}
	return ethcoder.HexEncode(imageHash), nil
}

func ImageHashOfWalletConfigBytes(walletConfig WalletConfig) ([]byte, error) {
	imageHash, err := ethcoder.SolidityPack([]string{"uint256"}, []interface{}{walletConfig.Threshold})
	if err != nil {
		return nil, fmt.Errorf("sequence, WalletConfigImageHash: %w", err)
	}

	for _, signer := range walletConfig.Signers {
		mm := [32]byte{}
		copy(mm[:], imageHash)

		weight := signer.Weight
		address := signer.Address

		aux, err := ethcoder.AbiCoder([]string{"bytes32", "uint8", "address"}, []interface{}{mm, weight, address})
		if err != nil {
			return nil, err
		}
		imageHash = ethcoder.Keccak256(aux)
	}

	return imageHash, nil
}

func ImageHashOfWalletConfigBytes32(walletConfig WalletConfig) ([32]byte, error) {
	imageHash, err := ImageHashOfWalletConfigBytes(walletConfig)
	if err != nil {
		return [32]byte{}, err
	}
	ih := [32]byte{}
	copy(ih[:], imageHash)
	return ih, nil
}

func SortWalletConfig(walletConfig WalletConfig) error {
	signers := walletConfig.Signers
	sort.Sort(signers) // Sort the signers

	// Ensure no duplicates
	for i := 0; i < len(signers)-1; i++ {
		if signers[i].Address == signers[i+1].Address {
			return fmt.Errorf("signer duplicate detected in the wallet config")
		}
	}

	return nil
}

func IsWalletConfigUsable(walletConfig WalletConfig) (bool, error) {
	// if walletConfig.Threshold.Cmp(big.NewInt(0)) == 0 {
	if walletConfig.Threshold == 0 {
		return false, fmt.Errorf("invalid wallet config - wallet threshold cannot be 0")
	}
	totalWeight := uint64(0)
	for _, s := range walletConfig.Signers {
		totalWeight += uint64(s.Weight)
	}
	// if walletConfig.Threshold.Cmp(big.NewInt(0).SetUint64(totalWeight)) > 0 {
	if uint64(walletConfig.Threshold) > totalWeight {
		return false, fmt.Errorf("invalid wallet config - total weight of the wallet config is less than the threshold required")
	}
	return true, nil
}

func IsWalletConfigEqual(walletConfigA, walletConfigB WalletConfig) bool {
	imageHashA, err := ImageHashOfWalletConfig(walletConfigA)
	if err != nil {
		return false
	}
	imageHashB, err := ImageHashOfWalletConfig(walletConfigB)
	if err != nil {
		return false
	}

	return imageHashA == imageHashB
}

// TODO: can leave out for now
/*type WalletState struct {
	Context WalletContext
	Config  *WalletConfig

	// the wallet address
	Address common.Address

	// the chainID of the network
	ChainID *big.Int

	// whether the wallet has been ever deployed
	Deployed bool

	// the imageHash of the `config` WalletConfig
	ImageHash string

	// the last imageHash of a WalletConfig, stored on-chain
	LastImageHash string

	// whether the WalletConfig object itself has been published to logs
	Published bool
}*/
