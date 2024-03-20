package intents

import (
	"encoding/json"
	"fmt"
)

func IntentResponseTypeToCode[T any](t *T) string {
	var data any = t
	switch data.(type) {
	case *IntentResponseSessionOpened:
		return IntentResponseCodeSessionOpened
	case *IntentResponseValidationStarted:
		return IntentResponseCodeValidationStarted
	case *IntentResponseValidationFinished:
		return IntentResponseCodeValidationFinished
	case *IntentResponseSessionAuthProof:
		return IntentResponseCodeSessionAuthProof
	case *IntentResponseSignedMessage:
		return IntentResponseCodeSignedMessage
	case *IntentResponseFeeOptions:
		return IntentResponseCodeFeeOptions
	case *IntentResponseTransactionReceipt:
		return IntentResponseCodeTransactionReceipt
	case *IntentResponseTransactionFailed:
		return IntentResponseCodeTransactionFailed
	case *IntentResponseGetSession:
		return IntentResponseCodeGetSessionResponse
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
