module tools

go 1.24.0

toolchain go1.24.1

// replace github.com/0xsequence/ethkit => ../../ethkit

require (
	github.com/0xsequence/ethkit v1.38.5
	github.com/webrpc/webrpc v0.28.1
	go.uber.org/mock v0.6.0
)

require (
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
)
