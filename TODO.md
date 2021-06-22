TODO
====

- [x] write list of interfaces for current / old implementation..
- [x] create core interfaces for config, network, wallet, signer, provider, etc......

- [x] tests, 1. start ganache node, etc..

- [x] tests, 2. deploy the wallet-context with a universal deployer so we have deterministic addresses etc.

- [x] tests, 3. create a wallet (undeployed), sign a message and validate it, then recover

- [x] tests, 4. create a wallet (deployed), sign a message and validate it, then recover

- [x] implement `Transaction` and `Transactions` to send basic meta-transaction and write test with CallReceiverMock

- [x] implement ComputeMetaTxID method

- [ ] implement support for `Transaction.Nested` to add support for nested txns, and add `Transactions.Encode()` method
      which will return a new `Transactions` array, where each child of depths >=2 will be reduced to single node,
			and calls `selfExecute` method all the way down the tree

- [ ] implement decoder of calldata for a meta-txn exec nested calldata, back to a `Transactions` structure

- [ ] implement helper functions to send parallel txns to the relayer to make it simpler for the developer
			when using go-sequence and wants to send native parallel txns which are not dependent on eachother (ie. skyweaver
			uses this for sending rewards and conquest tickets)

- [ ] implement LocalRelayer.Wait method -- listening on events for the metaTxID to show up, but we should also have a timeout. See WaitReceipt() impl 				in ethkit

- [ ] implement relayer.RpcRelayer which adheres to sequence.Relayer interface

- [ ] add erc20 mock contract to `contracts`, then write test deploy erc20 mock contract, mint some tokens,
			then do a transfer, then do another transfer with a batch

- [ ] part of testutil, include erc20 and erc1155 mock tokens so we can use it in other tests easily

- [x] add ethauth IsValidSequenceAccountProof

