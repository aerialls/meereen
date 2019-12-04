package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	c "github.com/aerialls/meereen/pkg/check"
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Loader loads elements into the container from the config
type Loader struct {
	logger *log.Logger
}

// NewLoader returns a new loader instance
func NewLoader(logger *log.Logger) *Loader {
	return &Loader{
		logger: logger,
	}
}

// LoadNotifiers loads notifiers from the config into the container
func (l *Loader) LoadNotifiers(container *Container, configs []NotifierConfig) error {
	notifiers := map[string]n.Notifier{}
	for _, config := range configs {
		notifier, err := n.GetNotifier(config.Kind, config.Data)
		if err != nil {
			return err
		}

		if _, ok := notifiers[config.Name]; ok {
			return fmt.Errorf("notifier %s already loaded", config.Name)
		}

		notifiers[config.Name] = notifier
	}

	// Store the map inside the container
	container.notifiers = notifiers

	return nil
}

// LoadChecks loads checks from the config into the container
func (l *Loader) LoadChecks(container *Container, configs []CheckConfig) error {
	checks := []c.Check{}
	for _, config := range configs {
		notifier, err := container.GetNotifier(config.Notifier)
		if err != nil {
			return err
		}

		processor, err := p.GetProcessor(
			config.Processor.Kind,
			config.Processor.Data,
		)

		if err != nil {
			return err
		}

		check := NewCheck(
			config.Title,
			processor,
			notifier,
			container.GetRetries(),
		)

		checks = append(checks, check)
	}

	// Store the list inside the container
	container.checks = checks

	return nil
}
