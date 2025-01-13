package routes

import (
	"MikuMikuCloudDrive/controllers/file_controller"
	"MikuMikuCloudDrive/middleware"

	"github.com/gin-gonic/gin"
)

func FileRouter(r *gin.Engine) {
	fileGroup := r.Group("/file")
	{
		fileGroup.GET("/download/:filename", file_controller.Download)
		fileGroup.Use(middleware.AuthMiddleware())
		fileGroup.POST("/upload", file_controller.Upload)
		fileGroup.POST("/getUploadedChunks", file_controller.GetUploadedChunks)
		fileGroup.POST("/merge", file_controller.MergeChunks)

	}

}
