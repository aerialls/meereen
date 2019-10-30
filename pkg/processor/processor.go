package processor

import (
	"fmt"
)

var processors = map[string]BuilderFunction{}

// Processor interface
type Processor interface {
	Process() (State, string)
}

// BuilderFunction represents the function prototype to implement
// to register a new builder
type BuilderFunction = func(map[string]string) (Processor, error)

// State of the processor
type State int

// List of all values for the processor after the check
const (
	Ok State = iota
	Error
)

// RegisterProcessor registers a new processor inside the global state
func RegisterProcessor(kind string, builder BuilderFunction) error {
	if _, ok := processors[kind]; ok {
		return fmt.Errorf("processor of kind %s already exists", kind)
	}

	processors[kind] = builder
	return nil
}

// GetProcessor returns a processor from its kind
func GetProcessor(kind string, data map[string]string) (Processor, error) {
	builder, ok := processors[kind]
	if !ok {
		return nil, fmt.Errorf("unable to find processor with kind %s", kind)
	}

	processor, err := builder(data)
	if err != nil {
		return nil, err
	}

	return processor, nil
}
