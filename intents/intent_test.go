package intents

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/0xsequence/ethkit/go-ethereum/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParseAndRecoverIntent(t *testing.T) {
	fmt.Println(uuid.New().String())

	data := `{
		"version": "1.0.0",
		"packet": {
			"code": "sendTransactions",
			"identifier": "test-identifier",
			"issued": 1600000000,
			"expires": 1600086400,
			"wallet": "0xD67FC48b298B09Ed3D03403d930769C527186c4e",
			"network": "1",
			"transactions": [{
				"type": "erc20send",
				"token": "0x0000000000000000000000000000000000000000",
				"to": "0x0dc9603d4da53841C1C83f3B550C6143e60e0425",
				"value": "0"
			}]
		},
		"signatures": [{
			"sessionId": "afaf60c0-67ba-4c9b-89ae-b115c78026a4",
			"signature": "0xcca6253c4fd281247ddd0fa487252ef91932eaec8d68b61f0901ccaa70345bf66fdbbd98ed3e3c9752f9e35ef2a7bc88dd9c8ae23c594241b476fe988824ab881c"
		}]
	}`

	intent := &Intent{}
	err := json.Unmarshal([]byte(data), intent)
	assert.Nil(t, err)

	assert.Equal(t, "1.0.0", intent.Version)

	hash, err := intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")

	getSessionVerifier := func(sessionId string) (string, error) {
		if sessionId == "afaf60c0-67ba-4c9b-89ae-b115c78026a4" {
			return "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", nil
		} else {
			return "", fmt.Errorf("invalid session id")
		}
	}

	signers := intent.Signers(getSessionVerifier)
	assert.Equal(t, 1, len(signers))
	assert.Equal(t, "0x1111BD4F3233e7a7f552AdAf32C910fD30de598B", signers[0])
	assert.Equal(t, intent.PacketCode(), "sendTransactions")

	// Changing the version should not affect the hash
	intent.Version = "2.0.0"
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")

	// Changing the packet code SHOULD affect the hash (and make Signers() return empty)
	intent.Packet = json.RawMessage(`{"code": "sendTransactions2"}`)
	hash, err = intent.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	assert.NotEqual(t, common.Bytes2Hex(hash), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")
	assert.Equal(t, intent.PacketCode(), "sendTransactions2")

	signers = intent.Signers(getSessionVerifier)
	assert.Equal(t, 0, len(signers))

	// Parsing the JSON without tabs, spaces, newlines, etc. should still work
	// and produce the same hash
	data2 := `{"signatures":[{"signature":"0xcca6253c4fd281247ddd0fa487252ef91932eaec8d68b61f0901ccaa70345bf66fdbbd98ed3e3c9752f9e35ef2a7bc88dd9c8ae23c594241b476fe988824ab881c","session":"0x1111BD4F3233e7a7f552AdAf32C910fD30de598B"}],"version":"1.0.0","packet":{"transactions":[{"token":"0x0000000000000000000000000000000000000000","value":"0","type":"erc20send","to":"0x0dc9603d4da53841C1C83f3B550C6143e60e0425"}],"wallet":"0xD67FC48b298B09Ed3D03403d930769C527186c4e","expires":1600086400,"code":"sendTransactions","network":"1","identifier":"test-identifier","issued":1600000000}}`
	intent2 := &Intent{}
	err = json.Unmarshal([]byte(data2), intent2)
	assert.Nil(t, err)

	hash2, err := intent2.Hash()
	assert.Nil(t, err)
	assert.NotNil(t, hash2)
	assert.Equal(t, common.Bytes2Hex(hash2), "893060f818437f8e3d9b4d8e103c5eb3c325fa25dd0221fb7b61cca6dd03a79e")
}

func TestBase64url(t *testing.T) {
	var keyMap = map[string]uint8{
		"0":  4,
		"1":  43,
		"2":  148,
		"3":  238,
		"4":  57,
		"5":  214,
		"6":  157,
		"7":  138,
		"8":  175,
		"9":  239,
		"10": 172,
		"11": 236,
		"12": 18,
		"13": 198,
		"14": 251,
		"15": 186,
		"16": 224,
		"17": 178,
		"18": 102,
		"19": 131,
		"20": 137,
		"21": 22,
		"22": 40,
		"23": 59,
		"24": 251,
		"25": 55,
		"26": 142,
		"27": 51,
		"28": 153,
		"29": 45,
		"30": 136,
		"31": 255,
		"32": 123,
		"33": 66,
		"34": 109,
		"35": 89,
		"36": 17,
		"37": 22,
		"38": 231,
		"39": 240,
		"40": 189,
		"41": 65,
		"42": 40,
		"43": 132,
		"44": 53,
		"45": 67,
		"46": 104,
		"47": 191,
		"48": 26,
		"49": 103,
		"50": 133,
		"51": 83,
		"52": 188,
		"53": 159,
		"54": 192,
		"55": 23,
		"56": 222,
		"57": 157,
		"58": 239,
		"59": 41,
		"60": 138,
		"61": 94,
		"62": 183,
		"63": 101,
		"64": 189,
	}

	var rawKey []byte
	for i := 0; i < len(keyMap); i++ {
		rawKey = append(rawKey, keyMap[fmt.Sprintf("%d", i)])
	}

	x, y := elliptic.Unmarshal(elliptic.P256(), rawKey)
	assert.NotNil(t, x)

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	spew.Dump(pub)

	sigMap := map[string]uint8{
		"0":  154,
		"1":  100,
		"2":  255,
		"3":  225,
		"4":  83,
		"5":  191,
		"6":  91,
		"7":  32,
		"8":  90,
		"9":  78,
		"10": 73,
		"11": 248,
		"12": 86,
		"13": 165,
		"14": 61,
		"15": 93,
		"16": 203,
		"17": 253,
		"18": 24,
		"19": 12,
		"20": 1,
		"21": 243,
		"22": 58,
		"23": 222,
		"24": 231,
		"25": 65,
		"26": 133,
		"27": 250,
		"28": 203,
		"29": 39,
		"30": 6,
		"31": 175,
		"32": 100,
		"33": 171,
		"34": 230,
		"35": 122,
		"36": 84,
		"37": 247,
		"38": 37,
		"39": 201,
		"40": 31,
		"41": 12,
		"42": 212,
		"43": 235,
		"44": 42,
		"45": 162,
		"46": 86,
		"47": 140,
		"48": 185,
		"49": 82,
		"50": 99,
		"51": 137,
		"52": 0,
		"53": 161,
		"54": 237,
		"55": 143,
		"56": 225,
		"57": 238,
		"58": 22,
		"59": 179,
		"60": 95,
		"61": 246,
		"62": 154,
		"63": 93,
	}

	var sigBytes []byte
	for i := 0; i < len(sigMap); i++ {
		sigBytes = append(sigBytes, sigMap[fmt.Sprintf("%d", i)])
	}

	hash := sha256.Sum256([]byte("test"))

	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:64])
	verified := ecdsa.Verify(&pub, hash[:], r, s)
	assert.True(t, verified)
}

func TestECDSAP256SessionSig(t *testing.T) {
	//sessionId := "r1:0x04d0ca179fec17ca13e0a05dfee4dd6e56c4f14975001ab093f6306f96f915c04d437047e1ca0244080fb2741d0518fe78c9e893c2a5b7d0549816198414c39335"
	//message := "0x7b2270726f6a6563744964223a31312c226964546f6b656e223a2265794a68624763694f694a53557a49314e694973496d74705a434936496a67315a5455314d5441334e445932596a646c4d6a6b344d7a59784f546c6a4e54686a4e7a55344d575931596a6b794d324a6c4e4451694c434a30655841694f694a4b5631516966512e65794a7063334d694f694a6f64485277637a6f764c32466a59323931626e527a4c6d6476623264735a53356a623230694c434a68656e41694f6949354e7a41354f4463334e5459324e6a41744d7a56684e6e526a4e44686f646d6b34593256324f574e75613235774d476c315a335935634739684d6a4d75595842776379356e6232396e624756316332567959323975644756756443356a623230694c434a68645751694f6949354e7a41354f4463334e5459324e6a41744d7a56684e6e526a4e44686f646d6b34593256324f574e75613235774d476c315a335935634739684d6a4d75595842776379356e6232396e624756316332567959323975644756756443356a623230694c434a7a645749694f6949784d5451324f4455344d5467784e4441784f44597a4e446b304d4451694c434a6c6257467062434936496d3168636d6c75627a4d355147647459576c734c6d4e7662534973496d567459576c7358335a6c636d6c6d6157566b496a7030636e566c4c434a75596d59694f6a45334d4459334f44677a4f546773496d3568625755694f694a4e59584a6a6157346752384f7a636e70357859527a61326b694c434a7761574e3064584a6c496a6f696148523063484d364c79397361444d755a3239765a32786c64584e6c636d4e76626e526c626e5175593239744c32457651554e6e4f47396a536b464c623031435956424e5a3064494e3246355a6e644e54485236596931786431706f526b524b527a645656485a70556c705854566f7850584d354e69316a496977695a326c325a573566626d46745a534936496b3168636d4e7062694973496d5a6862576c73655639755957316c496a6f6952384f7a636e70357859527a61326b694c434a7362324e68624755694f694a7762434973496d6c68644349364d5463774e6a63344f4459354f4377695a586877496a6f784e7a41324e7a6b794d6a6b344c434a7164476b694f6949314e4759315a5745334d6a646a4f546b314f544a684e446b784e6d497a59574d31593245344f4755334e7a637a596a677a4e546b35496e302e4d346c73304f434d6a30586d4152354c7358385852707074644f55724e4a34346845364d476e4e664770334d444e495642396e7043524567386d64415033674f427171784e4d4e3464566d50734d334e5a54385779575a546871616e736254484a6d586e556d355245767357675f456d7954525a436d486545414c4e6877636e57477a386d375636664a7a7050366543554474534942722d376832756f364f637a6b704c534b382d74324569796b533773706c54475952486f42427265364d787762685044377074574a767957466c6f35594a617a74574c347454376c795957504264504470334456555743687679622d707944654e2d5066756530716a333165326b7448756c6b6a576b7a35325a5a73723758314876734164767642436838614c5f3553354f2d314a3649475f455f4e497168373254596f4141683968374e7639474c706e746a6d5a566764687769755f6f485067222c2273657373696f6e41646472657373223a2272313a307830346430636131373966656331376361313365306130356466656534646436653536633466313439373530303161623039336636333036663936663931356330346434333730343765316361303234343038306662323734316430353138666537386339653839336332613562376430353439383136313938343134633339333335222c22667269656e646c794e616d65223a22f09f95b72072697475616c2076696c6c616765222c22696e74656e744a534f4e223a227b5c2276657273696f6e5c223a5c22302e302e302d646576315c222c5c227061636b65745c223a7b5c226973737565645c223a313730363738383639392c5c22657870697265735c223a313730363738383939392c5c22636f64655c223a5c226f70656e53657373696f6e5c222c5c2273657373696f6e5c223a5c2272313a3078303464306361313739666563313763613133653061303564666565346464366535366334663134393735303031616230393366363330366639366639313563303464343337303437653163613032343430383066623237343164303531386665373863396538393363326135623764303534393831363139383431346333393333355c222c5c2270726f6f665c223a7b5c226964546f6b656e5c223a5c2265794a68624763694f694a53557a49314e694973496d74705a434936496a67315a5455314d5441334e445932596a646c4d6a6b344d7a59784f546c6a4e54686a4e7a55344d575931596a6b794d324a6c4e4451694c434a30655841694f694a4b5631516966512e65794a7063334d694f694a6f64485277637a6f764c32466a59323931626e527a4c6d6476623264735a53356a623230694c434a68656e41694f6949354e7a41354f4463334e5459324e6a41744d7a56684e6e526a4e44686f646d6b34593256324f574e75613235774d476c315a335935634739684d6a4d75595842776379356e6232396e624756316332567959323975644756756443356a623230694c434a68645751694f6949354e7a41354f4463334e5459324e6a41744d7a56684e6e526a4e44686f646d6b34593256324f574e75613235774d476c315a335935634739684d6a4d75595842776379356e6232396e624756316332567959323975644756756443356a623230694c434a7a645749694f6949784d5451324f4455344d5467784e4441784f44597a4e446b304d4451694c434a6c6257467062434936496d3168636d6c75627a4d355147647459576c734c6d4e7662534973496d567459576c7358335a6c636d6c6d6157566b496a7030636e566c4c434a75596d59694f6a45334d4459334f44677a4f546773496d3568625755694f694a4e59584a6a6157346752384f7a636e70357859527a61326b694c434a7761574e3064584a6c496a6f696148523063484d364c79397361444d755a3239765a32786c64584e6c636d4e76626e526c626e5175593239744c32457651554e6e4f47396a536b464c623031435956424e5a3064494e3246355a6e644e54485236596931786431706f526b524b527a645656485a70556c705854566f7850584d354e69316a496977695a326c325a573566626d46745a534936496b3168636d4e7062694973496d5a6862576c73655639755957316c496a6f6952384f7a636e70357859527a61326b694c434a7362324e68624755694f694a7762434973496d6c68644349364d5463774e6a63344f4459354f4377695a586877496a6f784e7a41324e7a6b794d6a6b344c434a7164476b694f6949314e4759315a5745334d6a646a4f546b314f544a684e446b784e6d497a59574d31593245344f4755334e7a637a596a677a4e546b35496e302e4d346c73304f434d6a30586d4152354c7358385852707074644f55724e4a34346845364d476e4e664770334d444e495642396e7043524567386d64415033674f427171784e4d4e3464566d50734d334e5a54385779575a546871616e736254484a6d586e556d355245767357675f456d7954525a436d486545414c4e6877636e57477a386d375636664a7a7050366543554474534942722d376832756f364f637a6b704c534b382d74324569796b533773706c54475952486f42427265364d787762685044377074574a767957466c6f35594a617a74574c347454376c795957504264504470334456555743687679622d707944654e2d5066756530716a333165326b7448756c6b6a576b7a35325a5a73723758314876734164767642436838614c5f3553354f2d314a3649475f455f4e497168373254596f4141683968374e7639474c706e746a6d5a566764687769755f6f4850675c227d7d2c5c227369676e6174757265735c223a5b5d7d227d"
	//signature := "r1:0x5ca339d779449af10bc0dd74eb7e1ff6f9b584a52681bbb6dc27bb8923ec3ef629cdb9b17d4b6ed44fd125c4bb9376f869a2478513326613b03eb80a2bbfaeae"

	sessionId := "r1:0x040714f2ed82b5748ba30e3d81df81d481371b20c43cdbec81a89cbdb74e149e73ee083a1306328236c7de6d26b6f8d4494951d7423946422a04700ed182092a45"
	message := "0x7a7e5a0913e63cac5886afcafedba93b17baae3eb4066534ffdd5e3da3e8c714"
	signature := "r1:0x4038376385b045c19754bb69fa6cde925674778e6a1a78b8fa3135ec96b695aef5b8126c78dc17a2cc0be522a4e6154bf5152c908d763fb1c28e47cf419a3ea5"

	// get public key from sessionId
	sessionIdBuff := common.FromHex(sessionId[3:])

	x, y := elliptic.Unmarshal(elliptic.P256(), sessionIdBuff)
	assert.NotNil(t, x)

	pub := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	spew.Dump(pub)

	// get message hash
	messageBytes := common.FromHex(message)
	messageHash := sha256.Sum256(messageBytes)

	// get signature
	signatureBytes := common.FromHex(signature[3:])

	r := new(big.Int).SetBytes(signatureBytes[:32])
	s := new(big.Int).SetBytes(signatureBytes[32:64])
	verified := ecdsa.Verify(&pub, messageHash[:], r, s)
	assert.True(t, verified)
}
