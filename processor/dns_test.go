package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"

	p "github.com/aerialls/meereen/pkg/processor"
)

func TestNewDNS(t *testing.T) {
	DNSWrapper := func(params map[string]string) error {
		_, err := NewDNS(params)
		return err
	}

	errors := map[error]bool{
		// empty
		DNSWrapper(map[string]string{}): false,
		// missing "domain"
		DNSWrapper(map[string]string{
			"foo": "bar",
		}): false,
		// good
		DNSWrapper(map[string]string{
			"domain": "google.com",
		}): true,
		// good
		DNSWrapper(map[string]string{
			"domain":   "google.com",
			"resolver": "10.10.10.253",
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

func TestDNSWithoutCustomResolver(t *testing.T) {
	processor, err := NewDNS(map[string]string{
		"domain": "google.com",
	})

	assert.Nil(t, err)

	state, message := processor.Process()

	assert.Equal(t, state, p.Ok)
	assert.Empty(t, message)

	processor, err = NewDNS(map[string]string{
		"domain": "nonexistingdomain.local",
	})

	assert.Nil(t, err)

	state, message = processor.Process()

	assert.Equal(t, state, p.Error)
	assert.Contains(t, message, "Unable to resolve nonexistingdomain.local")
}

func TestDNSWithCustomResolver(t *testing.T) {
	processor, err := NewDNS(map[string]string{
		"domain":   "google.com",
		"resolver": "8.8.8.8",
	})

	assert.Nil(t, err)

	state, message := processor.Process()

	assert.Equal(t, state, p.Ok)
	assert.Empty(t, message)

	processor, err = NewDNS(map[string]string{
		"domain":   "google.com",
		"resolver": "192.168.1.254",
	})

	assert.Nil(t, err)

	state, message = processor.Process()

	assert.Equal(t, state, p.Error)
	assert.Contains(t, message, "Unable to resolve google.com")
}
