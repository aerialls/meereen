package core

import (
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
func (c *Check) Run() error {
	c.isRunning = true
	newState, message := c.processor.Process()
	if newState != c.state {
		err := c.notifier.Notify(c, newState, message)
		c.state = newState

		if err != nil {
			return nil
		}
	}

	c.isRunning = false
	return nil
}

// GetNotifier returns the notifier from the current check
func (c *Check) GetNotifier() n.Notifier {
	return c.notifier
}

// GetTitle returns the title of the current check
func (c *Check) GetTitle() string {
	return c.title
}

// GetState returns the current state from the latest processor
func (c *Check) GetState() p.State {
	return c.state
}

// IsRunning returns if the check is running right now
func (c *Check) IsRunning() bool {
	return c.isRunning
}
