// Client
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=intent.ridl -target=golang -pkg=intents -out=./intent.gen.go
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=intent.ridl -target=typescript -out=./intent.gen.ts

package intents
