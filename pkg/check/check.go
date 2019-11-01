package check

import (
	p "github.com/aerialls/meereen/pkg/processor"
)

// Check represents a way to verify a component
type Check interface {
	Run() error
	GetTitle() string
	GetState() p.State
	IsRunning() bool
}
