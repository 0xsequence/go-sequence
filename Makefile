all:
	@echo "See Makefile contents for details."

bootstrap:
	cd ./testutil/chain && yarn install

test:
	cd ./testutil/chain && yarn test

go-test:
	go clean -testcache && go test -p 1 -v ./...

start-test-chain:
	cd ./testutil/chain && yarn start:server

start-test-chain-verbose:
	cd ./testutil/chain && yarn start:server:verbose

clean:
	@go clean -testcache
