package mock

//go:generate go run go.uber.org/mock/mockgen -destination api.mock.go -package mock  -mock_names API=API ../../services/api API
//go:generate go run go.uber.org/mock/mockgen -destination indexer.mock.go -package mock -mock_names Indexer=Indexer,IndexerClient=IndexerClient,IndexerGateway=IndexerGateway,IndexerGatewayClient=IndexerGatewayClient ../../services/indexer Indexer,IndexerClient,IndexerGateway,IndexerGatewayClient
//go:generate go run go.uber.org/mock/mockgen -destination marketplace.mock.go -package mock  -mock_names Marketplace=Marketplace ../../services/marketplace Marketplace
//go:generate go run go.uber.org/mock/mockgen -destination metadata.mock.go -package mock  -mock_names Metadata=Metadata ../../services/metadata Metadata
//go:generate go run go.uber.org/mock/mockgen -destination relayer.mock.go -package mock  -mock_names Relayer=Relayer ../../relayer/proto Relayer
