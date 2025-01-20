package routes

import (
	"MikuMikuCloudDrive/controllers/directory_controller"
	"MikuMikuCloudDrive/middleware"

	"github.com/gin-gonic/gin"
)

func DirectoryRoute(r *gin.Engine) {
	directoryGroup := r.Group("/directory")
	{
		// JWT 认证
		directoryGroup.Use(middleware.AuthMiddleware())
		// 获取用户所有文件夹
		directoryGroup.GET("/list", directory_controller.GetDirectoryList)
		//创建文件夹
		directoryGroup.POST("/create", directory_controller.CreateDirectory)
		//删除文件夹
		directoryGroup.POST("/delete", directory_controller.DeleteDirectory)
		// 文件夹信息
		directoryGroup.GET("/info", directory_controller.GetDirectoryInfo)
		// 重命名
		directoryGroup.PUT("/rename", directory_controller.RenameDirectory)
		//移动
		directoryGroup.PUT("/move")
		// 复制
		directoryGroup.POST("/copy")
		// 分享
		directoryGroup.POST("/share")

		// 压缩文件夹
		directoryGroup.POST("/compress")
		//解压文件夹
		directoryGroup.POST("/extract")

	}
}
