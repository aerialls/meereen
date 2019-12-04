package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	c "github.com/aerialls/meereen/pkg/check"
	p "github.com/aerialls/meereen/pkg/processor"
)

type ManualProcessor struct {
	state p.State
}

type EmptyNotifier struct {
}

func (m *ManualProcessor) Process() (p.State, string) {
	return m.state, ""
}

func (e *EmptyNotifier) Notify(check c.Check, state p.State, message string) error {
	return nil
}

func TestCheckRun(t *testing.T) {
	processor := &ManualProcessor{state: p.Error}
	notifier := &EmptyNotifier{}

	check := NewCheck("Awesome check", processor, notifier, 0)

	assert.False(t, check.IsRunning())
	assert.Equal(t, p.Ok, check.GetState())

	err := check.Run()

	assert.Nil(t, err)
	assert.Equal(t, p.Error, check.GetState())
}

func TestCheckRetries(t *testing.T) {
	processor := &ManualProcessor{state: p.Error}
	notifier := &EmptyNotifier{}

	check := NewCheck("Awesome check", processor, notifier, 3)

	error := check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(1), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(2), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(3), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)

	processor.state = p.Ok

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)

	processor.state = p.Error

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(1), check.retries)

	processor.state = p.Ok

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)

	processor.state = p.Error

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(1), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(2), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(3), check.retries)

	error = check.Run()
	assert.Nil(t, error)
	assert.Equal(t, uint(0), check.retries)
}
