package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	p "github.com/aerialls/meereen/pkg/processor"
)

func TestNewHTTP(t *testing.T) {
	HTTPWrapper := func(params map[string]string) error {
		_, err := NewHTTP(params)
		return err
	}

	errors := map[error]bool{
		// empty
		HTTPWrapper(map[string]string{}): false,
		// missing "url"
		HTTPWrapper(map[string]string{
			"foo": "bar",
		}): false,
		// good
		HTTPWrapper(map[string]string{
			"url": "https://www.google.com/",
		}): true,
	}

	for err, ok := range errors {
		if ok {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestHTTPWithGoodSite(t *testing.T) {
	processor, err := NewHTTP(map[string]string{
		"url": "https://www.google.com/",
	})

	assert.Nil(t, err)

	state, message := processor.Process()

	assert.Equal(t, state, p.Ok)
	assert.Empty(t, message)
}

func TestHTTPWithWrongSite(t *testing.T) {
	processor, err := NewHTTP(map[string]string{
		"url": "https://www.mywebsite.local/",
	})

	assert.Nil(t, err)

	state, message := processor.Process()

	assert.Equal(t, state, p.Error)
	assert.Contains(t, message, "Unable to fetch the remote URL")
}
