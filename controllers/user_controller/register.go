package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/types/resgister_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	logrus.Debug("注册服务调用")
	resp := response.Response{}
	var req resgister_types.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, "register 绑定json失败")
		return
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	userService := services.NewUserService(svc.DB, svc.RedisClient)

	_, err = userService.Register(req.Username, req.Password)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	resp.SuccessResponse(c, "用户创建成功", "success")
}
