//
// sequence wallet-contracts
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletfactory --type=WalletFactory --outFile=./walletfactory/wallet_factory.gen.go --artifactsFile=../artifacts/wallet-contracts/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletmain --type=WalletMain --outFile=./walletmain/wallet_main_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModule.sol/MainModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletupgradable --type=WalletUpgradable --outFile=./walletupgradable/wallet_main_module_upgradable.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModuleUpgradable.sol/MainModuleUpgradable.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletguest --type=WalletGuest --outFile=./walletguest/wallet_guest_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/GuestModule.sol/GuestModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletutils --type=WalletUtils --outFile=./walletutils/wallet_utils.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/utils/SequenceUtils.sol/SequenceUtils.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletcallmock --type=CallReceiverMock --outFile=./walletcallmock/wallet_call_mock.gen.go --artifactsFile=../artifacts/wallet-contracts/mocks/CallReceiverMock.sol/CallReceiverMock.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletgasestimator --type=WalletGasEstimator --outFile=./walletgasestimator/wallet_gas_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModuleGasEstimation.sol/MainModuleGasEstimation.json --includeDeployed=true
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=gasestimator --type=GasEstimator --outFile=./gasestimator/gas_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/utils/GasEstimator.sol/GasEstimator.json --includeDeployed=true

//
// tokens
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC20 --outFile=./tokens/ierc20.gen.go --artifactsFile=../artifacts/erc-1155/interfaces/IERC20.sol/IERC20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC1155 --outFile=./tokens/ierc1155.gen.go --artifactsFile=../artifacts/erc-1155/interfaces/IERC1155.sol/IERC1155.json

//
// erc1271
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=ierc1271 --type=IERC1271 --outFile=./ierc1271/ierc1271.gen.go --artifactsFile=../artifacts/erc1271/ierc1271.json

package gen
