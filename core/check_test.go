package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/aerialls/meereen/pkg/check"
	p "github.com/aerialls/meereen/pkg/processor"
)

type ErrorProcessor struct {
}

type EmptyNotifier struct {
}

func (ep *ErrorProcessor) Process() (p.State, string) {
	return p.Error, ""
}

func (en *EmptyNotifier) Notify(check c.Check, state p.State, message string) error {
	return nil
}

func TestCheckRun(t *testing.T) {
	processor := &ErrorProcessor{}
	notifier := &EmptyNotifier{}

	check := NewCheck("Awesome check", processor, notifier)

	assert.False(t, check.IsRunning())
	assert.Equal(t, p.Ok, check.GetState())

	err := check.Run()

	assert.Nil(t, err)
	assert.Equal(t, p.Error, check.GetState())
}
