package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/user_service"
	"MikuMikuCloudDrive/types/login_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Login API(登录接口)
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body login_types.LoginRequest true "登录请求参数"
// @Success 200 {object} response.Response{data=login_types.LoginResponse} "成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /user/login [post]
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
	userService := user_service.NewUserService(db, rdb)
	loginResponse, err := userService.Login(req.Username, req.Password)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}
	resp.SuccessResponse(c, loginResponse.Token, "success")
}
