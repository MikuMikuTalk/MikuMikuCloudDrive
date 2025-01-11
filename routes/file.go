package routes

import (
	"MikuMikuCloudDrive/controllers/file_controller"
	"github.com/gin-gonic/gin"
)

func FileRouter(r *gin.Engine) {
	fileGroup := r.Group("/file")
	{
		fileGroup.POST("/upload", file_controller.Upload)
		fileGroup.POST("/download/:filename", file_controller.Download)
	}

}
