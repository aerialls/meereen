package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/aerialls/meereen/pkg/check"
	p "github.com/aerialls/meereen/pkg/processor"
)

type EmptyNotifier struct {
}

func (n *EmptyNotifier) Notify(check c.Check, state p.State, message string) error {
	return nil
}

func NewEmptyNotifier(data map[string]string) (Notifier, error) {
	return &EmptyNotifier{}, nil
}

func TestRegisterNotifier(t *testing.T) {
	err := RegisterNotifier("empty", NewEmptyNotifier)
	assert.Nil(t, err)

	err = RegisterNotifier("empty", NewEmptyNotifier)
	assert.NotNil(t, err)
}
