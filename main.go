package main

import (
	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/models/file_models"
	"MikuMikuCloudDrive/models/user_models"
	"MikuMikuCloudDrive/routes"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/utils/logger"
	"fmt"

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	app := config.ReadAppConfig()

	r := gin.Default()
	db := core.InitGorm()
	rdb := core.InitRedis()
	err := db.AutoMigrate(
		&user_models.UserModel{},
		&file_models.FileModel{},
	)
	if err != nil {
		panic("创建表结构失败")
	}
	logrus.Debug("表结构创建成功")

	// 初始化服务上下文
	svc := &services.ServiceContext{
		DB:          db,
		RedisClient: rdb,
	}
	r.Use(cors.Default())
	// 将服务上下文注入到 Gin 的上下文中
	r.Use(func(c *gin.Context) {
		c.Set("svc", svc)
		c.Next()
	})

	routes.UserRouter(r)
	routes.FileRouter(r)
	logger.InitLogger(logrus.DebugLevel)
	logrus.Infof("%s[Ver %s] is running on %s:%d", color.GreenString(app.Title), color.BlackString(app.Version), app.Server, app.Port)
	err = r.Run(fmt.Sprintf("%s:%d", app.Server, app.Port))
	if err != nil {
		panic(err)
	}
}
