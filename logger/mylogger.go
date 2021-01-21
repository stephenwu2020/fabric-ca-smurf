package logger

import (
	"github.com/sirupsen/logrus"
)

var MyLogger *logrus.Logger

func init() {
	MyLogger = logrus.New()

	MyLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	MyLogger.SetLevel(logrus.DebugLevel)
}
