package routes

import (
	"MikuMikuCloudDrive/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user_controller.Register)
		userGroup.POST("/login", user_controller.Login)
	}
}
