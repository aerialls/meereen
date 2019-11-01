package notifier

import (
	"fmt"

	c "github.com/aerialls/meereen/pkg/check"
	p "github.com/aerialls/meereen/pkg/processor"
)

var notifiers = map[string]BuilderFunction{}

// Notifier represents the way to send notifications
type Notifier interface {
	Notify(check c.Check, state p.State, message string) error
}

// BuilderFunction represents the function prototype to implement
// to register a new builder
type BuilderFunction = func(map[string]string) (Notifier, error)

// RegisterNotifier registers a new notifier inside the system
// The kind should be unique and will raise an error if its already exists
func RegisterNotifier(kind string, builder BuilderFunction) error {
	if _, ok := notifiers[kind]; ok {
		return fmt.Errorf("notifier of kind %s already exists", kind)
	}

	notifiers[kind] = builder
	return nil
}

// GetNotifier returns the notifier from the specified kind
func GetNotifier(kind string, data map[string]string) (Notifier, error) {
	builder, ok := notifiers[kind]
	if !ok {
		return nil, fmt.Errorf("unable to find notifier with kind %s", kind)
	}

	notifier, err := builder(data)
	if err != nil {
		return nil, err
	}

	return notifier, nil
}
