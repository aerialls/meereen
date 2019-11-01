package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredParameters(t *testing.T) {
	data := map[string]string{
		"foo": "bar",
	}

	result, err := GetRequiredParameter(data, "foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", result)

	result, err = GetRequiredParameter(data, "key")
	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestParameters(t *testing.T) {
	data := map[string]string{
		"foo": "bar",
	}

	result := GetParameter(data, "foo", "baar")
	assert.Equal(t, "bar", result)

	result = GetParameter(data, "key", "value")
	assert.Equal(t, "value", result)
}
