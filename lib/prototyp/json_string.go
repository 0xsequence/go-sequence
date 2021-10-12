package prototyp

import (
	"encoding/json"

	"database/sql"
	"database/sql/driver"
)

// JSONString is a custom database type that gets serialized into a text value.
type JSONString struct {
	v interface{}
}

func NewJSONString(v interface{}) *JSONString {
	return &JSONString{v}
}

func (s JSONString) Data() interface{} {
	return s.v
}

func (s JSONString) Value() (driver.Value, error) {
	return s.MarshalJSON()
}

func (s JSONString) MarshalJSON() ([]byte, error) {
	buf, err := json.Marshal(s.v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (s *JSONString) Scan(in interface{}) error {
	var buf []byte
	switch t := in.(type) {
	case string:
		buf = []byte(t)
	case []byte:
		buf = t
	}
	return s.UnmarshalJSON(buf)
}

func (s *JSONString) UnmarshalJSON(buf []byte) error {
	return json.Unmarshal(buf, &s.v)
}

var _ = interface {
	sql.Scanner
	driver.Valuer
}(&JSONString{})
