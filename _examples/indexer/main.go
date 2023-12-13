package main

import (
	"context"
	"log"

	"github.com/0xsequence/go-sequence/indexer"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	seqIndexer := indexer.NewIndexer("https://polygon-indexer.sequence.app", "kzk9Hdk0EJRoj9qgHWncKtHAAAAAAAAA")

	accountAddress := "0x8e3E38fe7367dd3b52D1e281E4e8400447C8d8B9"
	includeMetadata := true
	metadataOptions := indexer.MetadataOptions{
		VerifiedOnly: true,
	}

	_, tokenBalances, err := seqIndexer.GetTokenBalances(context.Background(), &accountAddress, nil, nil, &includeMetadata, &metadataOptions, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, tokenBalance := range tokenBalances {
		spew.Dump(tokenBalance)
	}
}
