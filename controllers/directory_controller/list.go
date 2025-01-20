package directory_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/directory_service"
	"MikuMikuCloudDrive/types/directory_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetDirectoryList 获取目录列表
// @Summary 获取目录列表
// @Description 获取指定用户的目录列表
// @Tags 目录管理
// @Accept json
// @Produce json
// @Param body body directory_types.GetDirectoryListRequest true "目录列表请求参数"
// @Success 200 {object} response.Response{data=directory_types.GetDirectoryListResponse} "成功返回目录信息"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /directory/list [get]
func GetDirectoryList(c *gin.Context) {
	resp := response.NewResponse()
	svc := services.GetServiceContextFromContext(c)
	dirService := directory_service.NewDirectoryService(svc.DB)
	req := directory_types.GetDirectoryListRequest{}
	// 获取用户信息
	claims := services.GetClaimsFromContext(c)
	logrus.Info("claims:", claims)
	getDirectoryListResponse, err := dirService.GetDirectoryList(req, claims)
	if err != nil {
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	resp.SuccessResponse(c, getDirectoryListResponse, "success")
}
