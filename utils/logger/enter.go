package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger(level logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(level)
}
