package intents

import (
	"encoding/json"
	"fmt"
)

func IntentResponseTypeToCode[T any](t *T) IntentResponseCode {
	var data any = t
	switch data.(type) {
	case *IntentResponseAuthInitiated:
		return IntentResponseCode_authInitiated
	case *IntentResponseSessionOpened:
		return IntentResponseCode_sessionOpened
	case *IntentResponseValidationStarted:
		return IntentResponseCode_validationStarted
	case *IntentResponseValidationFinished:
		return IntentResponseCode_validationFinished
	case *IntentResponseSessionAuthProof:
		return IntentResponseCode_sessionAuthProof
	case *IntentResponseSignedMessage:
		return IntentResponseCode_signedMessage
	case *IntentResponseFeeOptions:
		return IntentResponseCode_feeOptions
	case *IntentResponseTransactionReceipt:
		return IntentResponseCode_transactionReceipt
	case *IntentResponseTransactionFailed:
		return IntentResponseCode_transactionFailed
	case *IntentResponseGetSession:
		return IntentResponseCode_getSessionResponse
	case *IntentResponseAccountList:
		return IntentResponseCode_accountList
	case *IntentResponseAccountFederated:
		return IntentResponseCode_accountFederated
	case *IntentResponseAccountRemoved:
		return IntentResponseCode_accountRemoved
	case *IntentResponseIdToken:
		return IntentResponseCode_idToken
	default:
		return ""
	}
}

type IntentResponseTyped[T any] struct {
	IntentResponse
	Data T `json:"data"`
}

func NewIntentResponseTyped[T any](data T) *IntentResponseTyped[T] {
	return &IntentResponseTyped[T]{
		IntentResponse: IntentResponse{
			Code: IntentResponseTypeToCode(&data),
			Data: data,
		},
		Data: data,
	}
}

func NewIntentResponseTypedFromIntentResponse[T any](res *IntentResponse) (*IntentResponseTyped[T], error) {
	switch res.Data.(type) {
	case T:
		return &IntentResponseTyped[T]{
			IntentResponse: *res,
			Data:           res.Data.(T),
		}, nil
	case map[string]any:
		data := res.Data.(map[string]any)

		// convert to json
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		// convert to typed data
		var typedData T
		err = json.Unmarshal(dataJSON, &typedData)
		if err != nil {
			return nil, err
		}

		// check if intent response code and data type match
		if IntentResponseTypeToCode(&typedData) != res.Code {
			return nil, fmt.Errorf("intent response code and data type mismatch")
		}

		return &IntentResponseTyped[T]{
			IntentResponse: *res,
			Data:           typedData,
		}, nil
	default:
		return nil, fmt.Errorf("invalid intent data type")
	}
}
