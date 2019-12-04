package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	c "github.com/aerialls/meereen/pkg/check"
	n "github.com/aerialls/meereen/pkg/notifier"
)

// Container contains checks and notifiers
type Container struct {
	logger    *log.Logger
	notifiers map[string]n.Notifier
	checks    []c.Check
	retries   uint
	delta     uint
}

// NewContainer creates a new container
func NewContainer(logger *log.Logger) *Container {
	return &Container{
		logger: logger,
	}
}

// Load loads all elements (notifiers, checks, ...) from the config into the container
func (ctn *Container) Load(path string) error {
	cfg, err := NewConfig(path)
	if err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	ctn.retries = cfg.Global.Retries
	ctn.delta = cfg.Global.Delta

	loader := NewLoader(ctn.logger)

	err = loader.LoadNotifiers(ctn, cfg.GetNotifiers())
	if err != nil {
		return err
	}

	checks, err := cfg.GetChecks()
	if err != nil {
		return err
	}

	err = loader.LoadChecks(ctn, checks)
	if err != nil {
		return err
	}

	ctn.logger.Infof("%d check(s) has been loaded from %s", len(ctn.checks), cfg.ChecksFolder)
	return nil
}

// GetNotifier returns a notifier with a dedicated name
func (ctn *Container) GetNotifier(name string) (n.Notifier, error) {
	notifier, ok := ctn.notifiers[name]
	if !ok {
		return nil, fmt.Errorf("notifier %s does not exist", name)
	}

	return notifier, nil
}

// GetChecks returns all checks registered
func (ctn *Container) GetChecks() []c.Check {
	return ctn.checks
}

// GetDelta returns the number of seconds between two checks
func (ctn *Container) GetDelta() uint {
	return ctn.delta
}

// GetRetries returns the number of retries before sending notifications
func (ctn *Container) GetRetries() uint {
	return ctn.retries
}
