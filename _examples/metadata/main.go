package main

import (
	"context"
	"log"

	"github.com/0xsequence/go-sequence/metadata"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	seqMetadata := metadata.NewMetadata("kzk9Hdk0EJRoj9qgHWncKtHAAAAAAAAA")

	contractInfo, err := seqMetadata.GetContractInfo(context.Background(), "polygon", "0x631998e91476DA5B870D741192fc5Cbc55F5a52E")
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(contractInfo)

	collectibleInfo, err := seqMetadata.GetTokenMetadata(context.Background(), "polygon", "0x631998e91476DA5B870D741192fc5Cbc55F5a52E", []string{"1", "2"})
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(collectibleInfo)
}
