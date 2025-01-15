package user_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	userinfo_types "MikuMikuCloudDrive/types/user_info_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UserInfo(c *gin.Context) {
	logrus.Info("用户信息接口调用")
	resp := response.NewResponse()
	userInfoReq := userinfo_types.UserInfoRequest{
		Token: c.Request.Header.Get("Authorization"),
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	userService := services.NewUserService(svc.DB, svc.RedisClient)
	userinfoResp, err := userService.GetUserInfo(userInfoReq)
	if err != nil {
		resp.ErrorResponse(c, http.StatusBadGateway, "获取用户信息失败")
		return
	}
	resp.SuccessResponse(c, userinfoResp, "success")
}
