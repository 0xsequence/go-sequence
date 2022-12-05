package x

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/0xsequence/go-sequence/api"
)

func ValidateETHAuth() {
	seqAPI := api.NewAPIClient("https://api.sequence.app", http.DefaultClient)

	chainID := "137"
	walletAddress := "0x2fa0b551fdFa31a4471c1C52206fdb448ad997d1"
	ethAuthProofString := "eth.0x2fa0b551fdfa31a4471c1c52206fdb448ad997d1.eyJhcHAiOiJEZW1vIERhcHAiLCJpYXQiOjAsImV4cCI6MTY2MDIzMTAyOCwidiI6IjEiLCJvZ24iOiJodHRwOi8vbG9jYWxob3N0OjQwMDAifQ.0x000501032a44625bec3b842df681db00a92a74dda5e42bcf0203596af90cecdbf9a768886e771178fd5561dd27ab005d00010001f7dad5ade840bb961cbab889d731bbc080bb4c36fc090435e82fe78e3c152b671505ad544adb562cc25a5933cd06c9108e239a52a82ba797c3d3522645c69cd81b020101c50adeadb7fe15bee45dcb820610cdedcd314eb003000274164fb33c93b4384582c54c30d9a1e2ef219063d03084005edc1da853af2f1f2e67275dbb6ef945d04600b6dd83cfd997cc9ae4173ea61b0c5cc0808fb196681b02"

	isValid, err := seqAPI.IsValidETHAuthProof(context.Background(), chainID, walletAddress, ethAuthProofString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("isValid?", isValid)
}
