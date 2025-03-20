package mock

//go:generate go run go.uber.org/mock/mockgen -destination api.mock.go -package mock  -mock_names API=API ../api API
//go:generate go run go.uber.org/mock/mockgen -destination indexer.mock.go -package mock  -mock_names Indexer=Indexer ../indexer Indexer
//go:generate go run go.uber.org/mock/mockgen -destination marketplace.mock.go -package mock  -mock_names Marketplace=Marketplace ../marketplace Marketplace
//go:generate go run go.uber.org/mock/mockgen -destination metadata.mock.go -package mock  -mock_names Metadata=Metadata ../metadata Metadata
//go:generate go run go.uber.org/mock/mockgen -destination relayer.mock.go -package mock  -mock_names Relayer=Relayer ../relayer/proto Relayer
