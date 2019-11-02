package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	logger, _ := test.NewNullLogger()
	container := NewContainer(logger)

	err := container.LoadConfig("/foo/bar")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to load the config file")

	_, filename, _, _ := runtime.Caller(0)
	directory := filepath.Clean(fmt.Sprintf("%s/../tests", filepath.Dir(filename)))

	err = container.LoadConfig(fmt.Sprintf("%s/empty.yml", directory))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "checks parameter cannot be empty")
}
