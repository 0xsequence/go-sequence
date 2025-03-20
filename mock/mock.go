package mock

//go:generate go run go.uber.org/mock/mockgen -imports proto=../api -destination api.mock.go -package mock  -mock_names API=API ../api API
//go:generate go run go.uber.org/mock/mockgen -imports proto=../indexer -destination indexer.mock.go -package mock  -mock_names Indexer=Indexer ../indexer Indexer
//go:generate go run go.uber.org/mock/mockgen -imports proto=../marketplace -destination marketplace.mock.go -package mock  -mock_names Marketplace=Marketplace ../marketplace Marketplace
//go:generate go run go.uber.org/mock/mockgen -imports proto=../metadata -destination metadata.mock.go -package mock  -mock_names Metadata=Metadata ../metadata Metadata
//go:generate go run go.uber.org/mock/mockgen -imports proto=../relayer/proto -destination relayer.mock.go -package mock  -mock_names Relayer=Relayer ../relayer/proto Relayer
