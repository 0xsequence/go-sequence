go-sequence
===========

Sequence Wallet client written in Go.


## Usage

For documentation on sequence, please see our [docs](https://docs.sequence.xyz) page.

## Developing the go-sequence library

1. `make init-startchain` -- will install node modules of ./testutil/testchain
2. `make start-testchain` -- starts the test ethereum chain (id 1337)
3. (in a separate terminal) `make test` -- runs test suite


## Testing

Testing is super important, to run the tests just call `make test`. As well, you can
run the testchain separately with `make start-testchain` then in another terminal run `make test`.

**NOTE:** Go by default will execute tests in parallel if you run `go test -v ./...`, so ensure to pass `-p 1`
to set parallelization to just 1 (so it runs serially). The `make test` command is already set to do this.


## Other Go dev related tips

A. If you'd like to use a local version of a dependency/module, you can use the `replace` directive in go.mod,
for example, lets say you want to use a local version of "ethkit" that hasn't been released with go-sequence,
you can add `replace github.com/0xsequence/ethkit => /home/peter/Dev/0xsequence/ethkit` to your go.mod


## LICENSE

Apache 2.0
