package prototyp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONStringValuer(t *testing.T) {
	{
		s := NewJSONString("hello world")
		v, err := s.Value()

		assert.NoError(t, err)
		assert.Equal(t, []byte(`"hello world"`), v.([]byte))
	}
	{
		s := NewJSONString(map[string]interface{}{"foo": 1})
		v, err := s.Value()

		assert.NoError(t, err)
		assert.Equal(t, []byte(`{"foo":1}`), v.([]byte))
	}
	{
		s := NewJSONString(nil)
		v, err := s.Value()

		assert.NoError(t, err)
		assert.Equal(t, []byte(`null`), v.([]byte))
	}
}

func TestJSONStringScan(t *testing.T) {
	{
		var s JSONString
		err := s.Scan([]byte(`"hello world"`))
		assert.NoError(t, err)

		assert.Equal(t, "hello world", s.Data().(string))
	}
	{
		var s JSONString
		err := s.Scan(`123`)
		assert.NoError(t, err)

		assert.Equal(t, float64(123), s.Data().(float64))
	}
}
