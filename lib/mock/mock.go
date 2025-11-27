package mock

//go:generate go run go.uber.org/mock/mockgen -destination api.mock.go -package mock  -mock_names API=API ../../services/api API
//go:generate go run go.uber.org/mock/mockgen -destination indexer.mock.go -package mock -mock_names IndexerClient=IndexerClient,IndexerGatewayClient=IndexerGatewayClient ../../services/indexer IndexerClient,IndexerGatewayClient
//go:generate go run go.uber.org/mock/mockgen -destination marketplace.mock.go -package mock  -mock_names Marketplace=Marketplace ../../services/marketplace Marketplace
//go:generate go run go.uber.org/mock/mockgen -destination metadata.mock.go -package mock  -mock_names MetadataClient=Metadata ../../services/metadata MetadataClient
//go:generate go run go.uber.org/mock/mockgen -destination relayer.mock.go -package mock  -mock_names RelayerClient=Relayer ../../relayer/proto RelayerClient
