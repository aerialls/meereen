package core

import (
	"fmt"
	"io/ioutil"
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

func getContainer() (*Container, *test.Hook) {
	logger, hook := test.NewNullLogger()
	return NewContainer(logger), hook
}

func TestLoadChecksCorrectFolder(t *testing.T) {
	container, _ := getContainer()
	folder := getTestFolder()

	container.notifiers["empty"] = &EmptyNotifier{}

	err := container.loadChecks(fmt.Sprintf("%s/correct", folder))
	assert.Nil(t, err)
	assert.Len(t, container.checks, 4)
}

func TestLoadConfigCorrectFile(t *testing.T) {
	container, _ := getContainer()

	tmpFile, err := ioutil.TempFile("", "")
	assert.Nil(t, err)

	tmpFolder, err := ioutil.TempDir("", "")
	assert.Nil(t, err)

	tmpFile.Write([]byte(fmt.Sprintf(`checks: %s
notifiers: []`, tmpFolder)))

	err = container.LoadConfig(tmpFile.Name())
	assert.Nil(t, err)
}

func TestLoadConfigFileNotExists(t *testing.T) {
	container, _ := getContainer()
	err := container.LoadConfig("")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to load the config file")
}

func TestLoadConfigEmptyFile(t *testing.T) {
	container, _ := getContainer()

	tmpFile, err := ioutil.TempFile("", "")
	assert.Nil(t, err)

	err = container.LoadConfig(tmpFile.Name())
	assert.Contains(t, err.Error(), "checks parameter cannot be empty")
}

func TestLoadChecksFolderNotExists(t *testing.T) {
	container, _ := getContainer()

	err := container.loadChecks("")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func TestLoadChecksWithWrongNotifier(t *testing.T) {
	container, hook := getContainer()
	folder := getTestFolder()

	container.notifiers["empty"] = &EmptyNotifier{}

	err := container.loadChecks(fmt.Sprintf("%s/notifier", folder))

	assert.Nil(t, err)
	assert.Contains(t, hook.LastEntry().Message, "unable to load check")
	assert.Len(t, container.checks, 0)
}
