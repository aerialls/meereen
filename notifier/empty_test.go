package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"

	p "github.com/aerialls/meereen/pkg/processor"
)

func TestNewEmpty(t *testing.T) {
	notifier, err := NewEmpty(map[string]string{})
	assert.Nil(t, err)

	err = notifier.Notify(nil, p.Ok, "")
	assert.Nil(t, err)
}
