TEST_FLAGS       ?= -p 1 -v

all:
	@echo "See Makefile contents for details."

bootstrap:
	cd ./testutil/chain && pnpm install

build:
	go build ./...

test: wait-on-chain check-testchain-running go-test

go-test:
	go clean -testcache && go test $(TEST_FLAGS) -failfast=false -run=$(TEST) ./...

test-concurrently:
	cd ./testutil/chain && pnpm test

start-testchain:
	cd ./testutil/chain && pnpm start:geth

start-testchain-verbose:
	cd ./testutil/chain && pnpm start:geth:verbose

start-testchain-anvil:
	cd ./testutil/chain && pnpm start:anvil

start-testchain-anvil-verbose:
	cd ./testutil/chain && pnpm start:anvil:verbose

clean:
	@go clean -testcache

check-testchain-running:
	@curl http://localhost:8545 -H"Content-type: application/json" -X POST -d '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' --write-out '%{http_code}' --silent --output /dev/null | grep 200 > /dev/null \
	|| { echo "*****"; echo "Oops! testchain is not running. Please run 'make start-testchain' in another terminal or use 'test-concurrently'."; echo "*****"; exit 1; }

wait-on-chain:
	cd ./testutil/chain && pnpm wait:server
