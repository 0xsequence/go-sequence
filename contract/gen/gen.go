//
// sequence wallet-contracts
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletfactory --type=WalletFactory --outFile=./walletfactory/wallet_factory.gen.go --artifactsFile=../artifacts/wallet-contracts/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletmain --type=WalletMain --outFile=./walletmain/wallet_main_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModule.sol/MainModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletupgradable --type=WalletUpgradable --outFile=./walletupgradable/wallet_main_module_upgradable.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModuleUpgradable.sol/MainModuleUpgradable.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletguest --type=WalletGuest --outFile=./walletguest/wallet_guest_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/GuestModule.sol/GuestModule.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletcallmock --type=CallReceiverMock --outFile=./walletcallmock/wallet_call_mock.gen.go --artifactsFile=../artifacts/wallet-contracts/mocks/CallReceiverMock.sol/CallReceiverMock.json

//
// erc1271
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=ierc1271 --type=IERC1271 --outFile=./ierc1271/ierc1271.gen.go --artifactsFile=../artifacts/erc1271/ierc1271.json

package gen
