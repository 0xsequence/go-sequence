package intents

import (
	"encoding/json"
	"fmt"
	"time"
)

func IntentDataTypeToName[T any](t *T) string {
	var data any = t
	switch data.(type) {
	case *IntentDataOpenSession:
		return IntentNameOpenSession
	case *IntentDataCloseSession:
		return IntentNameCloseSession
	case *IntentDataValidateSession:
		return IntentNameValidateSession
	case *IntentDataFinishValidateSession:
		return IntentNameFinishValidateSession
	case *IntentDataListSessions:
		return IntentNameListSessions
	case *IntentDataGetSession:
		return IntentNameGetSession
	case *IntentDataAuthProof:
		return IntentNameAuthProof
	case *IntentDataSignMessage:
		return IntentNameSignMessage
	case *IntentDataFeeOptions:
		return IntentNameFeeOptions
	case *IntentDataSendTransaction:
		return IntentNameSendTransaction
	case *IntentDataGetTransactionReceipt:
		return IntentNameGetTransactionReceipt
	default:
		return ""
	}
}

type IntentTyped[T any] struct {
	Intent
	Data T `json:"data"`
}

func NewIntentTyped[T any](data T) *IntentTyped[T] {
	return &IntentTyped[T]{
		Intent: Intent{
			Version:   "1",
			ExpiresAt: uint64(time.Now().Unix()) + IntentValidTimeInSec,
			IssuedAt:  uint64(time.Now().Unix()),
			Name:      IntentDataTypeToName(&data),
			Data:      data,
		},
		Data: data,
	}
}

func NewIntentTypedFromIntent[T any](intent *Intent) (*IntentTyped[T], error) {
	switch intent.Data.(type) {
	case T:
		return &IntentTyped[T]{
			Intent: *intent,
			Data:   intent.Data.(T),
		}, nil
	case map[string]any:
		data := intent.Data.(map[string]any)

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

		// check if intent name and data type match
		if IntentDataTypeToName(&typedData) != intent.Name {
			return nil, fmt.Errorf("intent name and data type mismatch")
		}

		return &IntentTyped[T]{
			Intent: *intent,
			Data:   typedData,
		}, nil
	default:
		return nil, fmt.Errorf("invalid intent data type")
	}
}

func (i *IntentTyped[T]) IsValid() error {
	// check if the intent is valid
	if err := i.Intent.IsValid(); err != nil {
		return err
	}

	// check if the intent data is valid
	var data any = &i.Data
	if validator, ok := data.(IntentDataValidator); ok {
		if err := validator.IsValid(); err != nil {
			return fmt.Errorf("invalid intent data: %w", err)
		}
	}
	// the intent is valid
	return nil
}

func (i *IntentTyped[T]) ToIntent() *Intent {
	i.Intent.Data = i.Data
	return &i.Intent
}
