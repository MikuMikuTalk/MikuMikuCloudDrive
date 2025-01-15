package routes

import (
	"MikuMikuCloudDrive/controllers/user_controller"
	"MikuMikuCloudDrive/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/logout", user_controller.Logout)
		userGroup.POST("/login", user_controller.Login)
		userGroup.POST("/register", user_controller.Register)
		userGroup.Use(middleware.AuthMiddleware())
		userGroup.GET("/info", user_controller.UserInfo)
		userGroup.PUT("info", user_controller.UpdateUserInfo)
	}
}
