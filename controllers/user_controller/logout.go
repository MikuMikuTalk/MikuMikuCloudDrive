package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/user_service"
	"MikuMikuCloudDrive/types/logout_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Logout API(注销接口)
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body logout_types.LogoutRequest true "登录请求参数"
// @Success 200 {object} response.Response{data=logout_types.LogoutResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /user/logout [post]
func Logout(c *gin.Context) {
	resp := response.NewResponse()
	var logoutReq logout_types.LogoutRequest = logout_types.LogoutRequest{}
	svc := services.GetServiceContextFromContext(c)
	userService := user_service.NewUserService(svc.DB, svc.RedisClient)
	claims := services.GetClaimsFromContext(c)
	logoutResp, err := userService.Logout(logoutReq, claims)
	if err != nil {
		logrus.Error("Logout err:", err)
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessResponse(c, logoutResp, "success")
}
