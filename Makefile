all:
	@echo "See Makefile contents for details."

bootstrap:
	cd ./testutil/chain && yarn install

test:
	$(MAKE) check-test-chain-running && go clean -testcache && go test -p 1 -v ./...

test-ci:
	cd ./testutil/chain && yarn test

start-test-chain:
	cd ./testutil/chain && yarn start:server

start-test-chain-verbose:
	cd ./testutil/chain && yarn start:server:verbose

clean:
	@go clean -testcache

check-test-chain-running:
	@curl http://localhost:8545 -H"Content-type: application/json" -X POST -d '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' --write-out '%{http_code}' --silent --output /dev/null | grep 200 > /dev/null \
	|| { echo "*****"; echo "Oops! test-chain is not running. Please run 'make start-test-chain' in another terminal or use 'test-ci' to run+test concurrently."; echo "*****"; exit 1; }
