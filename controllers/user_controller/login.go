package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/types/login_types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	logrus.Debug("登录服务调用")
	var resp response.Response = response.Response{}
	var req login_types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, "login服务 绑定json失败")
		return
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	db := svc.DB
	rdb := svc.RedisClient
	userService := services.NewUserService(db, rdb)
	loginResponse, err := userService.Login(req.Username, req.Password)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	resp.SuccessResponse(c, loginResponse.Token, "success")
}
