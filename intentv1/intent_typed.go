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
		return "openSession"
	case *IntentDataCloseSession:
		return "closeSession"
	case *IntentDataValidateSession:
		return "validateSession"
	case *IntentDataFinishValidateSession:
		return "finishValidateSession"
	case *IntentDataListSessions:
		return "listSessions"
	case *IntentDataGetSession:
		return "getSession"
	case *IntentDataSign:
		return "sign"
	case *IntentDataTransaction:
		return "transaction"
	default:
		return ""
	}
}

type IntentTyped[T any] struct {
	Intent
	Data T
}

func NewIntentTyped[T any](data T) *IntentTyped[T] {
	return &IntentTyped[T]{
		Intent: Intent{
			Version: "1",
			Expires: uint64(time.Now().Unix()) + IntentValidTimeInSec,
			Issued:  uint64(time.Now().Unix()),
			Name:    IntentDataTypeToName(&data),
			Data:    data,
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

func (i *IntentTyped[T]) ToIntent() *Intent {
	var intentCopy = i.Intent
	intentCopy.Data = i.Data
	return &intentCopy
}
