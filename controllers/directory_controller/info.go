package directory_controller

import (
	"MikuMikuCloudDrive/common/response"
	"MikuMikuCloudDrive/services"
	"MikuMikuCloudDrive/services/directory_service"
	"MikuMikuCloudDrive/types/directory_types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDirectoryInfo 获取目录信息
// @Summary 获取目录信息
// @Description 获取指定目录的信息和内容列表
// @Tags 目录管理
// @Accept json
// @Produce json
// @Param directory_id query string true "目录ID"
// @Success 200 {object} response.Response{data=directory_types.GetDirectoryInfoResponse} "成功返回目录信息"
// @Failure 400 {object} response.Response "请求参数错误"
// @Failure 404 {object} response.Response "目录不存在"
// @Failure 500 {object} response.Response "服务器内部错误"
// @Router /directory/info [get]
func GetDirectoryInfo(c *gin.Context) {
	// 绑定请求参数
	resp := response.NewResponse()
	var req directory_types.GetDirectoryInfoRequest
	directoryID := c.Query("directory_id")
	// 设置请求参数
	req.DirectoryID = directoryID
	// 获取服务上下文
	svc := services.GetServiceContextFromContext(c)

	// 创建目录信息服务
	dirService := directory_service.NewDirectoryService(svc.DB)

	// 调用服务层获取目录信息
	claims := services.GetClaimsFromContext(c)
	info, err := dirService.GetDirectoryInfo(req, claims)
	if err != nil {
		resp.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp.SuccessResponse(c, info, "操作成功")
}
