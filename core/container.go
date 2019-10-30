package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/aerialls/meereen/pkg/check"
	"github.com/aerialls/meereen/pkg/notifier"
	"github.com/aerialls/meereen/pkg/processor"
)

// Container contains checks and notifiers
type Container struct {
	notifiers map[string]notifier.Notifier
	checks    []check.Check
}

// NewContainer creates a new container
func NewContainer() *Container {
	return &Container{
		notifiers: make(map[string]notifier.Notifier),
	}
}

// LoadConfig loads the YAML config file
func (c *Container) LoadConfig(path string) error {
	log.Infof("loading configuration file %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	cfg := Config{}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return err
	}

	for _, notifier := range cfg.Notifiers {
		err = c.loadNotifier(notifier)
		if err != nil {
			log.WithError(err).Warnf("ignoring notifier %s", notifier.Name)
		}
	}

	err = c.loadChecks(cfg.ChecksFolder)
	if err != nil {
		log.WithError(err).Fatal("unable to load checks")
	}

	log.Infof("%d check(s) has been loaded from %s", len(c.checks), cfg.ChecksFolder)

	return nil
}

func (c *Container) loadChecks(folder string) error {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		return err
	}

	if log.GetLevel() == log.DebugLevel {
		// Only check the number of file is the log level is the right one
		files, _ := ioutil.ReadDir(folder)
		log.Debugf("found %d file(s) in the check folder", len(files))
	}

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			err := c.loadCheck(path)
			if err != nil {
				log.WithError(err).Errorf("unable to load check %s", path)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Container) loadCheck(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	cfg := ConfigCheck{}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return err
	}

	notifier, err := c.GetNotifier(cfg.Notifier)
	if err != nil {
		return err
	}

	processor, err := processor.GetProcessor(
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

	log.WithFields(log.Fields{
		"title": cfg.Title,
		"kind":  cfg.Processor.Kind,
	}).Debugf("new check loaded")

	c.checks = append(c.checks, check)
	return nil
}

func (c *Container) loadNotifier(cfg ConfigNotifier) error {
	notifier, err := notifier.GetNotifier(cfg.Kind, cfg.Data)
	if err != nil {
		return err
	}

	if _, ok := c.notifiers[cfg.Name]; ok {
		return fmt.Errorf("notifier %s already loaded", cfg.Name)
	}

	log.WithField("name", cfg.Name).Debugf("new notifier of kind %s available", cfg.Kind)

	c.notifiers[cfg.Name] = notifier
	return nil
}

// GetNotifier returns a notifier with a dedicated name
func (c *Container) GetNotifier(name string) (notifier.Notifier, error) {
	notifier, ok := c.notifiers[name]
	if !ok {
		return nil, fmt.Errorf("notifier %s does not exist", name)
	}

	return notifier, nil
}

// GetChecks returns all checks registered
func (c *Container) GetChecks() []check.Check {
	return c.checks
}
