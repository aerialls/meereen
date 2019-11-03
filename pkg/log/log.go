package log

import (
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

// GetLogger returns the global logger
func GetLogger() *logrus.Logger {
	return log
}

// SetLogger sets the global logger
func SetLogger(logger *logrus.Logger) {
	log = logger
}
