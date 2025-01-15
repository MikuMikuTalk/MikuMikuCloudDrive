package user_controller

import (
	"net/http"

	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/user_service"
	"MikuMikuCloudDrive/types/resgister_types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Register API(注册接口)
// @Description 用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param body body resgister_types.RegisterRequest true "注册请求参数"
// @Success 200 {object} response.Response{data=resgister_types.RegisterResponse} "状态码为200 表示成功返回"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 502 {object} response.Response "服务内部错误"
// @Router /user/register [post]
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
	userService := user_service.NewUserService(svc.DB, svc.RedisClient)

	_, err = userService.Register(req.Username, req.Password)
	if err != nil {
		logrus.Error(err)
		resp.ErrorResponse(c, http.StatusBadGateway, err.Error())
		return
	}

	resp.SuccessResponse(c, "用户创建成功", "success")
}
