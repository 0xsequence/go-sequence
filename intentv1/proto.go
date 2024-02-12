// Server
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=intent.ridl -target=golang -pkg=intents -client -out=./intent.gen.go
//go:generate go run github.com/webrpc/webrpc/cmd/webrpc-gen -schema=intent.ridl -target=typescript -client -out=./intent.gen.ts

package intents
