TEST_FLAGS       ?= -p 1 -v

all:
	@echo "make <cmd>:"
	@echo ""
	@echo "commands:"
	@awk -F'[ :]' '/^#+/ {comment=$$0; gsub(/^#+[ ]*/, "", comment)} !/^(_|all:)/ && /^([A-Za-z_-]+):/ && !seen[$$1]++ {printf "  %-24s %s\n", $$1, (comment ? "- " comment : ""); comment=""} !/^#+/ {comment=""}' Makefile

bootstrap:
	cd ./testutil/chain && pnpm install

build:
	go build ./...

test: wait-on-chain check-testchain-running go-test

go-test:
	go clean -testcache && go test $(TEST_FLAGS) -run=$(TEST) ./...

clean:
	@go clean -testcache

.PHONY: mock
mock:
	go generate ./mock

test-concurrently:
	cd ./testutil/chain && pnpm test


#
# Testchain
#
start-testchain:
	cd ./testutil/chain && pnpm start:hardhat

start-testchain-verbose:
	cd ./testutil/chain && pnpm start:hardhat:verbose

start-testchain-geth:
	cd ./testutil/chain && pnpm start:geth

start-testchain-geth-verbose:
	cd ./testutil/chain && pnpm start:geth:verbose

start-testchain-anvil:
	cd ./testutil/chain && pnpm start:anvil

start-testchain-anvil-verbose:
	cd ./testutil/chain && pnpm start:anvil:verbose

check-testchain-running:
	@curl http://localhost:8545 -H"Content-type: application/json" -X POST -d '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' --write-out '%{http_code}' --silent --output /dev/null | grep 200 > /dev/null \
	|| { echo "*****"; echo "Oops! testchain is not running. Please run 'make start-testchain' in another terminal or use 'test-concurrently'."; echo "*****"; exit 1; }

wait-on-chain:
	cd ./testutil/chain && pnpm wait:server
