package log

import (
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func GetLogger() *logrus.Logger {
	return log
}

func SetLogger(logger *logrus.Logger) {
	log = logger
}
