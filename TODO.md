TODO
====

- [x] write list of interfaces for current / old implementation..
- [x] create core interfaces for config, network, wallet, signer, provider, etc......

- [x] tests, 1. start ganache node, etc..

- [x] tests, 2. deploy the wallet-context with a universal deployer so we have deterministic addresses etc.

- [x] tests, 3. create a wallet (undeployed), sign a message and validate it, then recover

- [ ] tests, 4. create a wallet (deployed), sign a message and validate it, then recover

- [ ] tests, 5. part of testutil we can include some erc20 mock bytecode + abi to make mocking a bit easier

- [ ] tests, 6. deploy erc20 mock contract, mint some tokens, then do a transfer, then do another transfer with a batch

- [ ] tests, 7. test the local relayer

- [ ] tests, 8. transaction, batch transactions, etc.....

- [ ] tests, 9. transaction with parallel nonces.. we use this in skyweaver, lets review.. etc.

- [ ] add ethauth IsValidSequenceAccountProof

-----------------------------
-----------------------------


- [ ] review @0xsequence/wallet tests, and move to here under wallet_test.go


-----------------------------
-----------------------------


- [ ] relayer stub -- where we will hook up the client .. (webrpc)
- [ ] indexer stub -- where we will hook up the client .. (webrpc)
- [ ] api stub -- where we will hook up the client .. (webrpc)


-----------------------------
-----------------------------
-----------------------------




type Account struct {
	Weight  *big.Int
	Address common.Address
	Wallet  *ethwallet.Wallet
}

* func (a *Account) GetAddress() common.Address 




-----------------------------



type SmartMetaTxn struct {
	DelegateCall  bool
	RevertOnError bool
	GasLimit      *big.Int
	Target        common.Address
	Value         *big.Int
	Data          []byte
}


-----------------------------




type SmartMetaTxnBatch struct {
	Wallet    common.Address
	ChainID   *big.Int
	Nonce     *big.Int
	Txns      []SmartMetaTxn
	contracts map[string]*abi.ABI
	Signature []byte
}

* func (s *SmartMetaTxnBatch) RegisterJSON(name string, contract string)
* func (s *SmartMetaTxnBatch) Register(name string, contract *abi.ABI)
* func (s *SmartMetaTxnBatch) DecodeTxns(input []byte) error
* func (s *SmartMetaTxnBatch) MustEncodeTxns() []byte
* func (s *SmartMetaTxnBatch) EncodeTxns() ([]byte, error)
* func (s *SmartMetaTxnBatch) Reset()
* func (s *SmartMetaTxnBatch) MustEncode(contract string, name string, params ...interface{}) []byte
* func (s *SmartMetaTxnBatch) Encode(contract string, name string, params ...interface{}) ([]byte, error)
* func (s *SmartMetaTxnBatch) Add(txn SmartMetaTxn)
* func (s *SmartMetaTxnBatch) SignOneToOne(wallet *ethwallet.Wallet) ([]byte, error)
* func (s *SmartMetaTxnBatch) Sign(sender []*Account, threshold ...*big.Int) ([]byte, error)
* func (s *SmartMetaTxnBatch) AnonymousSign() ([]byte, error)
* func (s *SmartMetaTxnBatch) AnonymousHash() ([]byte, error)
* func (s *SmartMetaTxnBatch) SelfExecuteHash() ([]byte, error)
* func (s *SmartMetaTxnBatch) Hash() ([]byte, error)
* func (s *SmartMetaTxnBatch) signImpl(sender []*Account, fake bool, threshold ...*big.Int) ([]byte, error)
* func (s *SmartMetaTxnBatch) DecodeInput(input []byte) error
* func (s *SmartMetaTxnBatch) DecodeTopics(logs []*types.Log, isAnonymous bool) ([][]*types.Log, error)




-----------------------------




* func mustNewArrayTypeTuple(components []abi.ArgumentMarshaling) abi.Type
* func mustNewType(str string) abi.Type
* func IsExecuteTxn(input []byte) bool
* func EncodeSmartWalletInput(txns []SmartMetaTxn, nonce *big.Int, signature []byte) ([]byte, error)
* func mustEncodeSig(str string) common.Hash
* func IsTxFailedLog(logs []*types.Log) (string, bool, error)
* func isNonceChangedEvent(logs []*types.Log) bool
* func EncodeMetaTxn(chainID *big.Int, owner common.Address, nonce *big.Int, txns []SmartMetaTxn) ([]byte, error)
* func hashData(chainID *big.Int, owner common.Address, message []byte) ([]byte, error)
* func Sign(account *Account, threshold *big.Int, data []byte, fake bool) ([]byte, error)
* func EncodeImageHashOneToOne(addr common.Address) ([32]byte, error)
* func EncodeImageHash(threshold *big.Int, accounts []*Account) ([32]byte, error)
* func hexToBigInt(str string) (*big.Int, error)
* func AddressOf(factory, mainModule common.Address, imageHash []byte) (common.Address, error)
* func EncodeNonce(space *big.Int, nonce *big.Int) *big.Int
* func DecodeNonce(raw *big.Int) (*big.Int, *big.Int)
* func GenerateRandomNonce() (*big.Int, error)
* func EncodeSelfExecute(txns []SmartMetaTxn) ([]byte, error)
* func IsValidSignature(address common.Address, message, sig []byte, chainID *big.Int, factory, mainModule *common.Address, provider *ethrpc.Provider) (bool, error)



-----------------------------



* var ArrayOfMetaTxnType = mustNewArrayTypeTuple([]abi.ArgumentMarshaling

* var arrayOfMetaTxns = abi.Arguments{
	abi.Argument{
		Type: ArrayOfMetaTxnType,
	},
}

* var (
	// TxExecutedEventSig is the signature of the event emitted in a successful transaction
	TxExecutedEventSig = mustEncodeSig("TxExecuted(bytes32)")

	// TxFailedEventSig is the signature event emitted in a failed smart-wallet meta-transaction batch
	// 0x3dbd1590ea96dd3253a91f24e64e3a502e1225d602a5731357bc12643070ccd7
	TxFailedEventSig = mustEncodeSig("TxFailed(bytes32,bytes)")

	// NonceChangeSig is the signature event emitted as the first event on the batch execution
	// 0x1f180c27086c7a39ea2a7b25239d1ab92348f07ca7bb59d1438fcf527568f881
	NonceChangeSig = mustEncodeSig("NonceChange(uint256,uint256)")
)

* const walletCode = "0x603a600e3d39601a805130553df3363d3d373d3d3d363d30545af43d82803e903d91601857fd5bf3"

* const isValidSignatureBytes32 = "0x1626ba7e"



--------------------------


// SigPart of a signature
type SigPart struct {
	Weight    int
	Address   *common.Address
	Signature []byte
}

* func (p *SigPart) Recover(msg []byte) (common.Address, error)
* func (p *SigPart) IsValid(digest [32]byte, provider *ethrpc.Provider) (bool, error)



// Signature for sequence message
type Signature struct {
	Threshold int
	Parts     []*SigPart
}

* func (s *Signature) Recover(msg []byte, provider *ethrpc.Provider) error
* func (s *Signature) Weight() (int, error)
* func (s *Signature) Copy() *Signature
* func (s *Signature) AggregateTwo(s2 *Signature) error
* func (s *Signature) ImageHash() ([32]byte, error)


var (
	partTypeSignature        = 0
	partTypeAddress          = 1
	partTypeDynamicSignature = 2
)

var (
	sigTypeEip712  = 1
	sigTypeEthSign = 2
	sigTypeEip1271 = 3
)


* func DecodeSignature(sig []byte) (*Signature, error)


* func Aggregate(sigs ...Signature) (*Signature, error) 

