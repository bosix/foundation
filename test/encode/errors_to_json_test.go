package encode

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanNotConvertStringToJsonError(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble("Foo")
	assert.False(t, result)
}

func TestOneErrorCanConvertToJson(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble(errors.New("entity not found"))
	assert.True(t, result)
}

func TestNotCorrectErrorCanNotConvertToJson(t *testing.T) {
	encoders := []inter.Encoder{encoder.ErrorToJson{}}
	result, err := encoder.ErrorToJson{}.EncodeThrough("foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to json in error format")
}

func TestOneErrorToJson(t *testing.T) {
	result, err := encoder.ErrorToJson{}.EncodeThrough(errors.New("entity not found"), defaultEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func TestOneErrorWithLongErrorMessage(t *testing.T) {
	result, err := encoder.ErrorToJson{}.EncodeThrough(
		errors.New(
			"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message",
		),
		defaultEncoders,
	)

	assert.Nil(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"This is a long error message, "+
		"this is a long error message, this is a long error message, this is a long error message, this is a long error message, this is a long error message, this is a long error message\"}]}", result)
}
