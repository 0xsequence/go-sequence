//
// sequence wallet-contracts v1
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletfactory --type=WalletFactory --outFile=./v1/walletfactory/wallet_factory.gen.go --artifactsFile=../artifacts/wallet-contracts/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletmain --type=WalletMain --outFile=./v1/walletmain/wallet_main_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModule.sol/MainModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletupgradable --type=WalletUpgradable --outFile=./v1/walletupgradable/wallet_main_module_upgradable.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModuleUpgradable.sol/MainModuleUpgradable.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletguest --type=WalletGuest --outFile=./v1/walletguest/wallet_guest_module.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/GuestModule.sol/GuestModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletutils --type=WalletUtils --outFile=./v1/walletutils/wallet_utils.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/utils/SequenceUtils.sol/SequenceUtils.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletutils --type=WalletRequireFreshSigner --outFile=./v1/walletutils/wallet_require_fresh_signer.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/utils/libs/RequireFreshSigner.sol/RequireFreshSigner.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletgasestimator --type=WalletGasEstimator --outFile=./v1/walletgasestimator/wallet_gas_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts/modules/MainModuleGasEstimation.sol/MainModuleGasEstimation.json --includeDeployed=true

//
// sequence wallet-contracts v2
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletfactory --type=WalletFactory --outFile=./v2/walletfactory/wallet_factory.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletmain --type=WalletMain --outFile=./v2/walletmain/wallet_main_module.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/MainModule.sol/MainModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletupgradable --type=WalletUpgradable --outFile=./v2/walletupgradable/wallet_main_module_upgradable.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/MainModuleUpgradable.sol/MainModuleUpgradable.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletguest --type=WalletGuest --outFile=./v2/walletguest/wallet_guest_module.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/GuestModule.sol/GuestModule.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletutils --type=WalletUtils --outFile=./v2/walletutils/wallet_utils.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/utils/SequenceUtils.sol/SequenceUtils.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletcallmock --type=CallReceiverMock --outFile=./walletcallmock/wallet_call_mock.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/mocks/CallReceiverMock.sol/CallReceiverMock.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletgasestimator --type=WalletGasEstimator --outFile=./v2/walletgasestimator/wallet_gas_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/MainModuleGasEstimation.sol/MainModuleGasEstimation.json --includeDeployed=true
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=gasestimator --type=GasEstimator --outFile=./gasestimator/gas_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts-v2/modules/utils/GasEstimator.sol/GasEstimator.json --includeDeployed=true

//
// sequence wallet-contracts v3
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletfactory --type=WalletFactory --outFile=./v3/walletfactory/wallet_factory.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Factory.sol/Factory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletstage1 --type=WalletStage1 --outFile=./v3/walletstage1/wallet_stage1.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Stage1Module.sol/Stage1Module.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletstage2 --type=WalletStage2 --outFile=./v3/walletstage2/wallet_stage2.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Stage2Module.sol/Stage2Module.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletguest --type=WalletGuest --outFile=./v3/walletguest/wallet_guest.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Guest.sol/Guest.json

//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletestimator --type=WalletEstimator --outFile=./v3/walletestimator/wallet_estimator.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Estimator.sol/Estimator.json --includeDeployed=true
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=walletsimulator --type=WalletSimulator --outFile=./v3/walletsimulator/wallet_simulator.gen.go --artifactsFile=../artifacts/wallet-contracts-v3/Simulator.sol/Simulator.json --includeDeployed=true

//
// tokens
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC20 --outFile=./tokens/ierc20.gen.go --artifactsFile=../artifacts/erc20/IERC20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC721 --outFile=./tokens/ierc721.gen.go --artifactsFile=../artifacts/erc721/IERC721.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC1155 --outFile=./tokens/ierc1155.gen.go --artifactsFile=../artifacts/erc1155/interfaces/IERC1155.sol/IERC1155.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=tokens --type=IERC20Wrapper --outFile=./tokens/ierc20wrapper.gen.go --artifactsFile=../artifacts/erc20-meta-token/interfaces/IERC20Wrapper.sol/IERC20Wrapper.json

//
// niftyswap
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=INiftyswapExchange --outFile=niftyswap/iniftyswap_exchange.gen.go --artifactsFile=../artifacts/niftyswap/interfaces/INiftyswapExchange.sol/INiftyswapExchange.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=INiftyswapExchange20 --outFile=niftyswap/iniftyswap_exchange_20.gen.go --artifactsFile=../artifacts/niftyswap/interfaces/INiftyswapExchange20.sol/INiftyswapExchange20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=INiftyswapFactory --outFile=niftyswap/iniftyswap_factory.gen.go --artifactsFile=../artifacts/niftyswap/interfaces/INiftyswapFactory.sol/INiftyswapFactory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=INiftyswapFactory20 --outFile=niftyswap/iniftyswap_factory_20.gen.go --artifactsFile=../artifacts/niftyswap/interfaces/INiftyswapFactory20.sol/INiftyswapFactory20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=IWrapAndNiftyswap --outFile=niftyswap/iwrap_and_niftyswap.gen.go --artifactsFile=../artifacts/niftyswap/interfaces/IWrapAndNiftyswap.sol/IWrapAndNiftyswap.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=NiftyswapExchange --outFile=niftyswap/niftyswap_exchange.gen.go --artifactsFile=../artifacts/niftyswap/exchange/NiftyswapExchange.sol/NiftyswapExchange.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=NiftyswapExchange20 --outFile=niftyswap/niftyswap_exchange_20.gen.go --artifactsFile=../artifacts/niftyswap/exchange/NiftyswapExchange20.sol/NiftyswapExchange20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=NiftyswapFactory --outFile=niftyswap/niftyswap_factory.gen.go --artifactsFile=../artifacts/niftyswap/exchange/NiftyswapFactory.sol/NiftyswapFactory.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=NiftyswapFactory20 --outFile=niftyswap/niftyswap_factory_20.gen.go --artifactsFile=../artifacts/niftyswap/exchange/NiftyswapFactory20.sol/NiftyswapFactory20.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=niftyswap --type=WrapAndNiftyswap --outFile=niftyswap/wrap_and_niftyswap.gen.go --artifactsFile=../artifacts/niftyswap/utils/WrapAndNiftyswap.sol/WrapAndNiftyswap.json

//
// erc1271
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=ierc1271 --type=IERC1271 --outFile=./ierc1271/ierc1271.gen.go --artifactsFile=../artifacts/erc1271/ierc1271.json

//
// sequence marketplace
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=seq_marketplace --type=SequenceMarketplace --outFile=./seq_marketplace/seq_marketplace.gen.go --artifactsFile=../artifacts/seq_marketplace/ISequenceMarket.sol/ISequenceMarket.json

//
// sequence sale
//
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=seq_sale_erc721  --type=Sale --outFile=./seq_sale/erc721/sale_erc721.gen.go   --artifactsFile=../artifacts/seq_sale/IERC721Sale.sol/IERC721Sale.json
//go:generate go run github.com/0xsequence/ethkit/cmd/ethkit abigen --pkg=seq_sale_erc1155 --type=Sale --outFile=./seq_sale/erc1155/sale_erc1155.gen.go --artifactsFile=../artifacts/seq_sale/IERC1155Sale.sol/IERC1155Sale.json

package gen
