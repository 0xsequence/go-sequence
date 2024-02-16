package sequence

import (
	"bytes"
	"fmt"
)

func IsEIP191Message(msg []byte) bool {
	return bytes.HasPrefix(msg, []byte("\x19Ethereum Signed Message:\n"))
}

func MessageToEIP191(msg []byte) []byte {
	if !IsEIP191Message(msg) {
		return bytes.Join([][]byte{
			[]byte("\x19Ethereum Signed Message:\n"),
			[]byte(fmt.Sprintf("%v", len(msg))),
			msg}, nil)
	}
	return msg
}
