package intents

import (
	"encoding/json"
	"fmt"
)

type IntentResponseTyped[T any] struct {
	IntentResponse
	Data T `json:"data"`
}

func NewIntentResponseTypedFromIntentResponse[T any](res *IntentResponse) (*IntentResponseTyped[T], error) {
	switch res.Code {
	case IntentResponseCodeSessionOpened,
		IntentResponseCodeValidationRequired,
		IntentResponseCodeValidationStarted,
		IntentResponseCodeValidationFinished,
		IntentResponseCodeSignedMessage,
		IntentResponseCodeFeeOptions,
		IntentResponseCodeTransactionReceipt,
		IntentResponseCodeTransactionFailed,
		IntentResponseCodeGetSessionResponse:
		data, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}

		var resTyped IntentResponseTyped[T]
		err = json.Unmarshal(data, &resTyped)
		if err != nil {
			return nil, err
		}
		return &resTyped, nil
	default:
		return nil, fmt.Errorf("unknown response code: %s", res.Code)
	}
}
