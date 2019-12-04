package processor

import (
	"testing"

	p "github.com/aerialls/meereen/pkg/processor"
	"github.com/stretchr/testify/assert"
)

func TestNewEmpty(t *testing.T) {
	processor, err := NewEmpty(map[string]string{})
	assert.Nil(t, err)

	state, _ := processor.Process()
	assert.Equal(t, p.Ok, state)
}
