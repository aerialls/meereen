package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getConfig(content string) (*Config, error) {
	file, _ := ioutil.TempFile("", "")
	file.Write([]byte(content))

	return NewConfig(file.Name())
}

func getConfigWithFolder(content string) (*Config, error) {
	emptyFolder, _ := ioutil.TempDir("", "")
	return getConfig(fmt.Sprintf(content, emptyFolder))
}

func getTestFolder() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Clean(fmt.Sprintf("%s/../tests", filepath.Dir(filename)))
}

func TestNewConfig(t *testing.T) {
	_, err := NewConfig("")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to load the config from file")

	emptyFile, _ := ioutil.TempFile("", "")
	container, err := NewConfig(emptyFile.Name())

	assert.Nil(t, err)
	assert.NotNil(t, container)
}

func TestConfigDefaultValues(t *testing.T) {
	// No values
	emptyFile, _ := ioutil.TempFile("", "")
	config, _ := NewConfig(emptyFile.Name())

	assert.Equal(t, DefaultGlobalConfig.Delta, config.Global.Delta)
	assert.Equal(t, DefaultGlobalConfig.Retries, config.Global.Retries)

	// Good values
	config, _ = getConfig(`
global:
  delta: 55
  retries: 32
`)

	assert.Equal(t, uint(55), config.Global.Delta)
	assert.Equal(t, uint(32), config.Global.Retries)

	// Wrong values
	config, err := getConfig(`
global:
  delta: -5
`)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot unmarshal !!int")
}

func TestConfigValidate(t *testing.T) {
	// Good validation
	config, _ := getConfigWithFolder(`checks_folder: %s`)

	err := config.Validate()
	assert.Nil(t, err)

	// Checks folder not existing
	config, _ = getConfig(`checks_folder: /foobar`)

	err = config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "checks folder /foobar does not exist", err.Error())

	// Wrong delta
	config, _ = getConfigWithFolder(`
checks_folder: %s
global:
    delta: 0
`)

	err = config.Validate()
	assert.Nil(t, err)
	assert.Equal(t, DefaultGlobalConfig.Delta, config.Global.Delta)
}
