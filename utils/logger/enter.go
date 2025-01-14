package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(level logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		DisableColors:   false,                       // 不禁用颜色
		TimestampFormat: "2006-01-02T15:04:05+08:00", // 时间格式
		ForceColors:     true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(level)
}
