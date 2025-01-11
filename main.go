package main

import (
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/utils/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.InitLogger(logrus.DebugLevel)
	core.InitGorm()
}
