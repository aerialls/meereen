package config

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"

	"github.com/aerialls/meereen/notifier"
	n "github.com/aerialls/meereen/pkg/notifier"
	_ "github.com/aerialls/meereen/processor"
)

func GetContainerAndLoader() (*Container, *Loader) {
	logger, _ := test.NewNullLogger()
	loader := NewLoader(logger)
	container := NewContainer(logger)

	container.notifiers = map[string]n.Notifier{
		"empty": &notifier.Empty{},
	}

	return container, loader
}

func TestLoadNotifiers(t *testing.T) {
	logger, _ := test.NewNullLogger()
	loader := NewLoader(logger)
	container := NewContainer(logger)

	// Empty
	notifiers := []NotifierConfig{}
	err := loader.LoadNotifiers(container, notifiers)

	assert.Nil(t, err)
	assert.Len(t, container.notifiers, 0)

	// Non-existing kind
	notifiers = []NotifierConfig{
		NotifierConfig{
			Name: "default",
			Kind: "notexisting",
		},
	}

	err = loader.LoadNotifiers(container, notifiers)

	assert.NotNil(t, err)
	assert.Equal(t, "unable to find notifier with kind notexisting", err.Error())
	assert.Len(t, container.notifiers, 0)

	// Empty notifier
	notifiers = []NotifierConfig{
		NotifierConfig{
			Name: "default",
			Kind: "empty",
		},
	}

	err = loader.LoadNotifiers(container, notifiers)

	assert.Nil(t, err)
	assert.Len(t, container.notifiers, 1)

	empty, err := container.GetNotifier("default")
	assert.Nil(t, err)
	assert.NotNil(t, empty)
}

func TestLoadChecks(t *testing.T) {
	container, loader := GetContainerAndLoader()

	// Empty
	checks := []CheckConfig{}
	err := loader.LoadChecks(container, checks)

	assert.Nil(t, err)
	assert.Len(t, container.checks, 0)

	// Good check
	checks = []CheckConfig{
		CheckConfig{
			Title:    "My check",
			Notifier: "empty",
			Processor: ProcessorConfig{
				Kind: "empty",
			},
		},
	}
}

func TestLoadChecksError(t *testing.T) {
	container, loader := GetContainerAndLoader()

	checks := map[string]CheckConfig{
		"notifier foobar does not exist": CheckConfig{
			Title:    "My check",
			Notifier: "foobar",
		},
		"unable to find processor with kind foobar": CheckConfig{
			Title:    "My check",
			Notifier: "empty",
			Processor: ProcessorConfig{
				Kind: "foobar",
			},
		},
		"required parameter url is missing": CheckConfig{
			Title:    "My check",
			Notifier: "empty",
			Processor: ProcessorConfig{
				Kind: "http",
			},
		},
	}

	for message, check := range checks {
		err := loader.LoadChecks(container, []CheckConfig{check})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), message)
	}
}
