package core

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"

	_ "github.com/aerialls/meereen/processor"
)

func getTestFolder() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Clean(fmt.Sprintf("%s/../tests", filepath.Dir(filename)))
}

func TestLoadConfig(t *testing.T) {
	logger, _ := test.NewNullLogger()
	container := NewContainer(logger)
	folder := getTestFolder()

	err := container.LoadConfig("/foo/bar")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to load the config file")

	err = container.LoadConfig(fmt.Sprintf("%s/empty.yml", folder))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "checks parameter cannot be empty")
}

func TestLoadChecks(t *testing.T) {
	logger, _ := test.NewNullLogger()
	container := NewContainer(logger)
	folder := getTestFolder()

	container.notifiers["empty"] = &EmptyNotifier{}

	err := container.loadChecks(fmt.Sprintf("%s/foobarfolder", folder))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not exist")

	err = container.loadChecks(fmt.Sprintf("%s/good", folder))
	assert.Nil(t, err)
	assert.Len(t, container.checks, 2)
}

func TestLoadChecksWithWrongNotifier(t *testing.T) {
	logger, hook := test.NewNullLogger()
	container := NewContainer(logger)
	folder := getTestFolder()

	container.notifiers["empty"] = &EmptyNotifier{}

	err := container.loadChecks(fmt.Sprintf("%s/notifier", folder))
	assert.Nil(t, err)
	assert.Contains(t, hook.LastEntry().Message, "unable to load check")
	assert.Len(t, container.checks, 0)
}
