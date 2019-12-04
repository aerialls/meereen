package config

import (
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Check representation
type Check struct {
	isRunning  bool
	title      string
	processor  p.Processor
	state      p.State
	notifier   n.Notifier
	maxRetries uint
	retries    uint
}

// NewCheck returns a new check instance
func NewCheck(
	title string,
	processor p.Processor,
	notifier n.Notifier,
	retries uint,
) *Check {
	return &Check{
		title:      title,
		processor:  processor,
		notifier:   notifier,
		state:      p.Ok,
		isRunning:  false,
		maxRetries: retries,
		retries:    0,
	}
}

// Run the processor
func (c *Check) Run() error {
	c.isRunning = true
	defer func() { c.isRunning = false }()

	newState, message := c.processor.Process()
	if newState != c.state || c.retries != 0 {
		// Send the notification if the new status is "OK" (no retries for this status)
		// Or if we are already above the max retries limit
		if newState == p.Ok || c.retries >= c.maxRetries {
			defer func() { c.retries = 0 }()

			err := c.notifier.Notify(c, newState, message)
			c.state = newState

			if err != nil {
				return nil
			}
		} else {
			c.retries++
			return nil
		}
	}

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
