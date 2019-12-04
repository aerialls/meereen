package processor

import (
	p "github.com/aerialls/meereen/pkg/processor"
)

// Empty processor
type Empty struct {
}

// NewEmpty returns an empty processor (for testing)
func NewEmpty(data map[string]string) (p.Processor, error) {
	return &Empty{}, nil
}

// Process the empty processor by doing nothing
func (e *Empty) Process() (p.State, string) {
	return p.Ok, ""
}

func init() {
	p.RegisterProcessor("empty", NewEmpty)
}
