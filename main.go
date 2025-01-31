package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"MikuMikuCloudDrive/config"
	"MikuMikuCloudDrive/core"
	"MikuMikuCloudDrive/ioc"
	"MikuMikuCloudDrive/middleware"
	"MikuMikuCloudDrive/models"
	"MikuMikuCloudDrive/routes"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/utils/logger"

	_ "MikuMikuCloudDrive/docs" // 导入生成的 docs 包

	"github.com/fatih/color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initDatabase() error {
	db := core.InitGorm()
	err := db.AutoMigrate(
		&models.UserModel{},
		&models.FileModel{},
		&models.DirectoryModel{},
	)
	if err != nil {
		logrus.Error("创建表结构失败:", err)
		return err
	}
	logrus.Info("数据库表结构初始化成功")
	return nil
}

func initCorsMiddleware(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                           // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},                                // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length", "X-Custom-Header"},                           // 允许暴露的响应头
		AllowCredentials: true,                                                                    // 允许凭证
	}))
}

func outputRunningInfo(r *gin.Engine) {
	app := config.ReadAppConfig()

	logrus.Infof("%s[Ver %s] is running on %s:%d", color.GreenString(app.Title), color.BlackString(app.Version), app.Server, app.Port)
	if err := r.Run(fmt.Sprintf("%s:%d", app.Server, app.Port)); err != nil {
		logrus.Error("服务器启动失败:", err)
		os.Exit(1)
	}
}

func injectSvcIntoContext(r *gin.Engine, svc *services.ServiceContext) {
	// 将服务上下文注入到 Gin 的上下文中
	r.Use(func(c *gin.Context) {
		c.Set("svc", svc)
		c.Next()
	})
}

// @title MikuMikuCloudDrive
// @version 1.0
// @description Gin实现的云盘后端文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT LICENSE
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8888

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization
// @Security JWTAuth
func main() {
	// 解析命令行参数
	initDB := flag.Bool("initdb", false, "Initialize database tables")
	flag.Parse()

	// 初始化日志
	logger.InitLogger(logrus.DebugLevel)

	// 如果指定了--initdb参数，只初始化数据库
	// main.go 文件已修改完成，现在支持通过 --initdb 参数初始化数据库表结构
	if *initDB {
		if err := initDatabase(); err != nil {
			logrus.Error("数据库初始化失败:", err)
			os.Exit(1)
		}
		return
	}
	// 自动迁移数据库
	if err := initDatabase(); err != nil {
		logrus.Error("数据库初始化失败:", err)
		os.Exit(1)
	}

	r := gin.Default()

	//初始化跨域中间件
	initCorsMiddleware(r)

	// 初始化服务上下文
	svc, err := ioc.InitializeApp()
	if err != nil {
		logrus.Error("初始化上下文失败")
		return
	}
	logrus.Debug("Gin 上下文创建成功！")

	//注入svc到gin的上下文
	injectSvcIntoContext(r, svc)

	// 注册路由
	routes.UserRouter(r)
	routes.FileRouter(r)
	routes.DirectoryRoute(r)

	// 注册静态文件
	r.StaticFS("/web", http.Dir("./web"))

	r.GET("/web", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/web/index.html")
	})

	// 注册swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 获取配置文件中的上传目录
	uploadDir := svc.Configuration.App.UploadDir
	r.Use(middleware.CheckDirectoryAccess)
	r.StaticFS("/uploads", http.Dir("./"+uploadDir))

	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	//	输出运行信息
	outputRunningInfo(r)

}
