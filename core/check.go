package core

import (
	log "github.com/sirupsen/logrus"

	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Check representation
type Check struct {
	isRunning bool
	title     string
	processor p.Processor
	state     p.State
	notifier  n.Notifier
}

// NewCheck returns a new check instance
func NewCheck(
	title string,
	processor p.Processor,
	notifier n.Notifier,
) *Check {
	return &Check{
		title:     title,
		processor: processor,
		notifier:  notifier,
		state:     p.Ok,
		isRunning: false,
	}
}

// Run the processor
func (c *Check) Run() {
	c.isRunning = true

	log.WithField("title", c.title).Debug("running check")
	newState, message := c.processor.Process()
	if newState != c.state {
		c.notifier.Notify(c, newState, message)
		c.state = newState
	}

	c.isRunning = false
}

// GetNotifier returns the notifier from the current check
func (c *Check) GetNotifier() n.Notifier {
	return c.notifier
}

// GetTitle returns the title of the current check
func (c *Check) GetTitle() string {
	return c.title
}

// IsRunning returns if the check is running right now
func (c *Check) IsRunning() bool {
	return c.isRunning
}
