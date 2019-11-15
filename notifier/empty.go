package notifier

import (
	c "github.com/aerialls/meereen/pkg/check"
	n "github.com/aerialls/meereen/pkg/notifier"
	p "github.com/aerialls/meereen/pkg/processor"
)

// Empty notifier
type Empty struct {
}

// NewEmpty returns a new empty notifier
func NewEmpty(data map[string]string) (n.Notifier, error) {
	return &Empty{}, nil
}

// Notify empty
func (e *Empty) Notify(check c.Check, state p.State, message string) error {
	return nil
}

func init() {
	n.RegisterNotifier("empty", NewEmpty)
}
