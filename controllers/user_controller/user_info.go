package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/user_service"
	userinfo_types "MikuMikuCloudDrive/types/user_info_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary 获取用户信息 API(获取用户信息)
// @Description 用户信息获取接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body userinfo_types.UserInfoRequest true "注册请求参数"
// @Success 200 {object} response.Response{data=userinfo_types.UserInfoResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	logrus.Info("用户信息接口调用")
	resp := response.NewResponse()
	userInfoReq := userinfo_types.UserInfoRequest{
		Token: c.Request.Header.Get("Authorization"),
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	userService := user_service.NewUserService(svc.DB, svc.RedisClient)
	userinfoResp, err := userService.GetUserInfo(userInfoReq)
	if err != nil {
		resp.ErrorResponse(c, http.StatusBadGateway, "获取用户信息失败")
		return
	}
	resp.SuccessResponse(c, userinfoResp, "success")
}

// @Summary 更新用户信息 API(更新用户信息)
// @Description 更新用户信息接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body userinfo_types.UpdateUserInfoRequest true "注册请求参数"
// @Success 200 {object} response.Response{data=userinfo_types.UpdateUserInfoResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /user/info [put]
func UpdateUserInfo(c *gin.Context) {
	logrus.Info("用户信息更新接口调用")
	resp := response.NewResponse()
	userInfoUpdateReq := userinfo_types.UpdateUserInfoRequest{}
	if err := c.ShouldBindHeader(&userInfoUpdateReq); err != nil {
		resp.ErrorResponse(c, http.StatusBadRequest, "error")
		return
	}
	if err := c.ShouldBindJSON(&userInfoUpdateReq); err != nil {
		resp.ErrorResponse(c, http.StatusBadRequest, "error")
		return
	}
	svc := c.MustGet("svc").(*services.ServiceContext)
	userService := user_service.NewUserService(svc.DB, svc.RedisClient)
	userInfoUpdateResp, err := userService.UpdateUserInfo(userInfoUpdateReq)
	if err != nil {
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	resp.SuccessResponse(c, userInfoUpdateResp, "success")
}
