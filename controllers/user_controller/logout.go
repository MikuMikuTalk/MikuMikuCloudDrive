package user_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/types/logout_types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Logout(c *gin.Context) {
	resp := response.NewResponse()

	token := c.Request.Header.Get("Authorization")
	var logoutReq logout_types.LogoutRequest = logout_types.LogoutRequest{
		Token: token,
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	userService := services.NewUserService(svc.DB, svc.RedisClient)
	logoutResp, err := userService.Logout(logoutReq)
	if err != nil {
		logrus.Error("Logout err:", err)
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessResponse(c, logoutResp, "success")
}
