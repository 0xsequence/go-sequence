package proto

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// TxnArgs preserves JSON number precision when decoding from the database.
type TxnArgs map[string]interface{}

func (t TxnArgs) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return json.Marshal(map[string]interface{}(t))
}

func (t *TxnArgs) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*t = nil
		return nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()

	var out map[string]interface{}
	if err := dec.Decode(&out); err != nil {
		return err
	}
	*t = TxnArgs(out)
	return nil
}

func (t TxnArgs) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}

	payload, err := json.Marshal(map[string]interface{}(t))
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (t *TxnArgs) Scan(src interface{}) error {
	if src == nil {
		*t = nil
		return nil
	}

	switch value := src.(type) {
	case []byte:
		return t.UnmarshalJSON(value)
	case string:
		return t.UnmarshalJSON([]byte(value))
	case json.RawMessage:
		return t.UnmarshalJSON(value)
	case map[string]interface{}:
		*t = TxnArgs(value)
		return nil
	default:
		return fmt.Errorf("TxnArgs.Scan: unsupported type %T", src)
	}
}
