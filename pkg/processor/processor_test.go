package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EmptyProcessor struct {
}

func (p *EmptyProcessor) Process() (State, string) {
	return Ok, ""
}

func NewEmptyProcessor(data map[string]string) (Processor, error) {
	return &EmptyProcessor{}, nil
}

func TestRegisterProcessor(t *testing.T) {
	err := RegisterProcessor("empty", NewEmptyProcessor)
	assert.Nil(t, err)

	err = RegisterProcessor("empty", NewEmptyProcessor)
	assert.NotNil(t, err)
}
