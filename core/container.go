package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	c "github.com/aerialls/meereen/pkg/check"
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Container contains checks and notifiers
type Container struct {
	logger    *log.Logger
	notifiers map[string]n.Notifier
	checks    []c.Check
}

// NewContainer creates a new container
func NewContainer(logger *log.Logger) *Container {
	return &Container{
		logger:    logger,
		notifiers: make(map[string]n.Notifier),
	}
}

// LoadConfig loads the YAML config file
func (c *Container) LoadConfig(path string) error {
	c.logger.Infof("loading configuration file %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to load the config file (%s)", err)
	}

	cfg := Config{}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return err
	}

	for _, notifier := range cfg.Notifiers {
		err = c.loadNotifier(notifier)
		if err != nil {
			c.logger.WithError(err).Warnf("ignoring notifier %s", notifier.Name)
		}
	}

	if err := c.validate(cfg); err != nil {
		return err
	}

	err = c.loadChecks(cfg.ChecksFolder)
	if err != nil {
		return err
	}

	c.logger.Infof("%d check(s) has been loaded from %s", len(c.checks), cfg.ChecksFolder)

	return nil
}

func (c *Container) validate(cfg Config) error {
	if cfg.ChecksFolder == "" {
		return fmt.Errorf("checks parameter cannot be empty")
	}

	return nil
}

func (c *Container) loadChecks(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return fmt.Errorf("checks folder %s does not exist", folder)
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		// Only check the number of file is the log level is the right one
		files, _ := ioutil.ReadDir(folder)
		c.logger.Debugf("found %d file(s) in the check folder", len(files))
	}

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			err := c.loadCheckFile(path)
			if err != nil {
				c.logger.WithError(err).Errorf("unable to load checks in file %s (%s)", path, err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Container) loadCheckFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	cfg := ConfigChecks{}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return err
	}

	if len(cfg.Checks) == 0 {
		log.Warnf("no checks in file %s", filepath.Base(path))
		return nil
	}

	for _, cfg := range cfg.Checks {
		notifier, err := c.GetNotifier(cfg.Notifier)
		if err != nil {
			return err
		}

		processor, err := p.GetProcessor(
			cfg.Processor.Kind,
			cfg.Processor.Data,
		)
		if err != nil {
			return err
		}

		check := NewCheck(
			cfg.Title,
			processor,
			notifier,
		)

		c.logger.WithFields(log.Fields{
			"title": cfg.Title,
			"kind":  cfg.Processor.Kind,
		}).Debugf("new check loaded")

		c.checks = append(c.checks, check)
	}

	return nil
}

func (c *Container) loadNotifier(cfg ConfigNotifier) error {
	notifier, err := n.GetNotifier(cfg.Kind, cfg.Data)
	if err != nil {
		return err
	}

	if _, ok := c.notifiers[cfg.Name]; ok {
		return fmt.Errorf("notifier %s already loaded", cfg.Name)
	}

	c.logger.WithField("name", cfg.Name).Debugf("new notifier of kind %s available", cfg.Kind)

	c.notifiers[cfg.Name] = notifier
	return nil
}

// GetNotifier returns a notifier with a dedicated name
func (c *Container) GetNotifier(name string) (n.Notifier, error) {
	notifier, ok := c.notifiers[name]
	if !ok {
		return nil, fmt.Errorf("notifier %s does not exist", name)
	}

	return notifier, nil
}

// GetChecks returns all checks registered
func (c *Container) GetChecks() []c.Check {
	return c.checks
}
